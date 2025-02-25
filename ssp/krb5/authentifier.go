package krb5

import (
	"context"
	"fmt"
	"os"

	"github.com/oiweiwei/gokrb5.fork/v9/client"
	"github.com/oiweiwei/gokrb5.fork/v9/credentials"
	krb_crypto "github.com/oiweiwei/gokrb5.fork/v9/crypto"
	"github.com/oiweiwei/gokrb5.fork/v9/iana/etypeID"
	"github.com/oiweiwei/gokrb5.fork/v9/messages"
	"github.com/oiweiwei/gokrb5.fork/v9/service"
	"github.com/oiweiwei/gokrb5.fork/v9/spnego"
	"github.com/oiweiwei/gokrb5.fork/v9/types"

	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	"github.com/oiweiwei/go-msrpc/ssp/krb5/crypto"
)

type Authentifier struct {
	// The authentifier configuration.
	Config *Config
	// The kerberos client.
	client *client.Client
	// The service settings.
	service *service.Settings
	// The AP Req message.
	APReq *messages.APReq
	// The AP Rep message.
	APRep *messages.APRep
	// The session key.
	SessionKey types.EncryptionKey
	// Exported session key.
	ExportedSessionKey []byte
	// key.
	state *SecurityService
}

type SecurityService struct {
	Key                    types.EncryptionKey
	IsSubKey               bool
	OutboundSequenceNumber uint64
	InboundSequenceNumber  uint64
	OutboundCipher         crypto.Cipher
	InboundCipher          crypto.Cipher
}

func (a *Authentifier) makeService(ctx context.Context) (*service.Settings, error) {
	kt, _ := a.Config.Credential.(credential.Keytab)
	return service.NewSettings(kt.Keytab(), a.Config.ServiceSettings()...), nil
}

func (a *Authentifier) tryLoadCCache(ctx context.Context) (*credentials.CCache, bool, error) {
	p := a.Config.CCachePath
	if p != "" {
		cc, err := credentials.LoadCCache(p)
		if err != nil {
			return nil, false, fmt.Errorf("load ccache: %w", err)
		}
		return cc, true, nil
	}

	if p = os.Getenv("KRB5_CCACHE"); p != "" {
		if cc, _ := credentials.LoadCCache(p); cc != nil {
			return cc, true, nil
		}
	}

	return nil, false, nil
}

// makeClient function initializes the kerberos client.
func (a *Authentifier) makeClient(ctx context.Context) (*client.Client, error) {

	var cli *client.Client
	var err error

	if a.Config.GetKRB5Config() == nil {
		// load kerberos config.
		if a.Config.KRB5Config, err = LoadKRB5Conf(a.Config.KRB5ConfigPath); err != nil {
			return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
		}
	}

	// try load credentials cache.
	cc, ok, err := a.tryLoadCCache(ctx)
	if err != nil {
		// ccache usage was required.
		return nil, err
	}

	creds := credentials.New(a.Config.Credential.UserName(), a.Config.Credential.DomainName())
	if ok {
		// ccache is present.
		cli, err = client.NewFromCCacheOptionalTGT(cc, a.Config.GetKRB5Config(), a.Config.ClientSettings()...)
		if err != nil {
			return nil, fmt.Errorf("client from ccache: %w", err)
		}
	}

	if cli == nil {
		cli = client.NewWithPassword(creds.UserName(), creds.Realm(), "",
			a.Config.GetKRB5Config(), a.Config.ClientSettings()...)
	}

	switch cred := a.Config.Credential.(type) {
	case credential.Password:
		cli.Credentials = creds.WithPassword(cred.Password())
	case credential.Keytab:
		cli.Credentials = creds.WithKeytab(cred.Keytab())
	case credential.NTHash:
		cli.Credentials = creds.WithEncryptionKey(types.EncryptionKey{
			KeyType:  etypeID.RC4_HMAC,
			KeyValue: cred.NTHash(),
		})
		setEnctype(cli, etypeID.RC4_HMAC)
	case credential.EncryptionKey:
		cli.Credentials = creds.WithEncryptionKey(types.EncryptionKey{
			KeyType:  int32(cred.KeyType()),
			KeyValue: cred.KeyValue(),
		})
		setEnctype(cli, int32(cred.KeyType()))
	case credential.CCache:
		cli, err = client.NewFromCCacheOptionalTGT(cred.CCache(), a.Config.GetKRB5Config(), a.Config.ClientSettings()...)
		if err != nil {
			return nil, fmt.Errorf("client from ccache credential: %w", err)
		}
	}

	_, err = cli.IsConfigured()
	if err != nil {
		return nil, fmt.Errorf("client configuration: %w", err)
	}

	return cli, nil
}

func setEnctype(cli *client.Client, etype int32) {
	cli.Config.LibDefaults.DefaultTGSEnctypeIDs = []int32{etype}
	cli.Config.LibDefaults.DefaultTktEnctypeIDs = []int32{etype}
	cli.Config.LibDefaults.DefaultTktEnctypeIDs = []int32{etype}
}

// makeSecurityService function sets up the security service for
// signing/encryption.
func (a *Authentifier) makeSecurityService(ctx context.Context) error {
	if a.APRep != nil {
		// derive service from mutual-authentication. (dce-style or mutual-authn
		// GSS flags are set.)
		a.state = &SecurityService{
			Key:                    a.APRep.DecryptedEncPart.Subkey,
			IsSubKey:               true,
			OutboundSequenceNumber: uint64(a.APReq.Authenticator.SeqNumber),
			InboundSequenceNumber:  uint64(a.APRep.DecryptedEncPart.SequenceNumber),
		}
	} else {
		a.state = &SecurityService{
			Key:                    a.SessionKey,
			OutboundSequenceNumber: uint64(a.APReq.Authenticator.SeqNumber),
			InboundSequenceNumber:  uint64(a.APReq.Authenticator.SeqNumber),
		}
	}

	eType, err := krb_crypto.GetEtype(a.state.Key.KeyType)
	if err != nil {
		return fmt.Errorf("get etype: %w", err)
	}

	clientCipher, err := crypto.NewCipher(ctx, crypto.CipherSetting{
		Type:     eType,
		Key:      a.state.Key,
		IsSubKey: a.state.IsSubKey,
		IsLocal:  true,
		DCEStyle: a.Config.DCEStyle,
	})
	if err != nil {
		return fmt.Errorf("client cipher: %w", err)
	}

	serverCipher, err := crypto.NewCipher(ctx, crypto.CipherSetting{
		Type:     eType,
		Key:      a.state.Key,
		IsSubKey: a.state.IsSubKey,
		IsLocal:  false,
		DCEStyle: a.Config.DCEStyle,
	})
	if err != nil {
		return fmt.Errorf("server cipher: %w", err)
	}

	if !a.Config.IsServer {
		a.state.InboundCipher, a.state.OutboundCipher = serverCipher, clientCipher
	} else {
		a.state.InboundCipher, a.state.OutboundCipher = clientCipher, serverCipher
	}

	a.ExportedSessionKey = a.state.Key.KeyValue

	return nil
}

func (a *Authentifier) APRequest(ctx context.Context) ([]byte, error) {

	var err error

	if a.client, err = a.makeClient(ctx); err != nil {
		return nil, fmt.Errorf("krb5: init: apreq: make client: %w", err)
	}

	// affirm only for password and key (keytab) credentials. if ccache is used,
	// it should fail on getting service ticket.
	if a.client.Credentials.HasPassword() || a.client.Credentials.HasKeyProvider() {
		if err := a.client.AffirmLogin(); err != nil {
			return nil, fmt.Errorf("krb5: init: apreq: affirm login: %w", err)
		}
	}

	tkt, key, err := a.client.GetServiceTicket(a.Config.SName)
	if err != nil {
		return nil, fmt.Errorf("krb5: init: apreq: get service ticket: %w", err)
	}

	tok, err := spnego.NewKRB5TokenAPREQ(a.client, tkt, key, a.Config.Flags, a.Config.APOptions)
	if err != nil {
		return nil, fmt.Errorf("krb5: init: apreq: call new_krb5_token_apreq: %w", err)
	}

	a.APReq, a.SessionKey = &tok.APReq, key

	if err := a.APReq.DecryptAuthenticator(a.SessionKey); err != nil {
		return nil, fmt.Errorf("krb5: init: apreq: decrypt authenticator: %w", err)
	}

	if a.Config.DCEStyle {
		b, err := tok.APReq.Marshal()
		if err != nil {
			return nil, fmt.Errorf("krb5: init: apreq: marshal: dcestyle: %w", err)
		}

		return b, nil
	}

	b, err := tok.Marshal()
	if err != nil {
		return nil, fmt.Errorf("krb5: init: apreq: %w", err)
	}

	return b, nil
}

func (a *Authentifier) APReply(ctx context.Context, b []byte) ([]byte, error) {

	if len(b) == 0 {
		if err := a.makeSecurityService(ctx); err != nil {
			return nil, fmt.Errorf("krb5: init: aprep: make security service: %w", err)
		}
		return nil, nil
	}

	a.APRep = &messages.APRep{}

	if !a.Config.DCEStyle {
		tok := &spnego.KRB5Token{}
		if err := tok.Unmarshal(b); err != nil {
			return nil, fmt.Errorf("krb5: init: aprep: unmarshal: dcestyle: %w", err)
		}
		a.APRep = &tok.APRep
	} else {
		if err := a.APRep.Unmarshal(b); err != nil {
			return nil, fmt.Errorf("krb5: init: aprep: unmarshal: %w", err)
		}
	}

	if err := a.APRep.DecryptEncPart(a.SessionKey); err != nil {
		return nil, fmt.Errorf("krb5: init: aprep: decrypt enc part: %w", err)
	}

	ap, err := messages.NewAPRep(a.SessionKey, messages.NewEncAPRepPart(a.APRep.DecryptedEncPart.SequenceNumber))
	if err != nil {
		return nil, fmt.Errorf("krb5: init: aprep: new aprep: %w", err)
	}

	b, err = ap.Marshal()
	if err != nil {
		return nil, fmt.Errorf("krb5: init: aprep: marshal: %w", err)
	}

	if err := a.makeSecurityService(ctx); err != nil {
		return nil, fmt.Errorf("krb5: aprep: make security service: %w", err)
	}

	return b, nil
}

func (a *Authentifier) MakeOutboundSignature(ctx context.Context, forSign [][]byte) ([]byte, error) {
	sgn, err := a.state.OutboundCipher.MakeSignature(ctx, a.state.OutboundSequenceNumber, forSign)
	if err != nil {
		return nil, fmt.Errorf("krb5: make outbound signature: %w", err)
	}
	a.state.OutboundSequenceNumber++
	return sgn, nil
}

func (a *Authentifier) MakeInboundSignature(ctx context.Context, forSign [][]byte) ([]byte, error) {
	sgn, err := a.state.InboundCipher.MakeSignature(ctx, a.state.InboundSequenceNumber, forSign)
	if err != nil {
		return nil, fmt.Errorf("krb5: make inbound signature: %w", err)
	}
	a.state.InboundSequenceNumber++
	return sgn, nil
}

func (a *Authentifier) WrapOutboundPayload(ctx context.Context, forSign, forSeal [][]byte) ([]byte, error) {
	sgn, err := a.state.OutboundCipher.Wrap(ctx, a.state.OutboundSequenceNumber, forSign, forSeal)
	if err != nil {
		return nil, fmt.Errorf("krb5: wrap outbound payload: %w", err)
	}

	a.state.OutboundSequenceNumber++
	return sgn, nil
}

func (a *Authentifier) UnwrapInboundPayload(ctx context.Context, forSign, forSeal [][]byte, sgn []byte) (bool, error) {
	ok, err := a.state.InboundCipher.Unwrap(ctx, a.state.InboundSequenceNumber, forSign, forSeal, sgn)
	if err != nil {
		return ok, fmt.Errorf("krb5: unwrap inbound payload: %w", err)
	}
	if ok {
		a.state.InboundSequenceNumber++
	}
	return ok, nil
}

func (a *Authentifier) OutboundSignatureSize(ctx context.Context, conf bool) int {
	// use outbound cipher as we wrap the outbound data.
	return a.state.OutboundCipher.Size(ctx, conf)
}
