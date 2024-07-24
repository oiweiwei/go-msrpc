package krb5

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/credentials"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
	"github.com/jcmturner/gokrb5/v8/service"
	"github.com/jcmturner/gokrb5/v8/spnego"
	"github.com/jcmturner/gokrb5/v8/types"

	"github.com/oiweiwei/go-msrpc/ssp/credential"

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
	APReq *APReq
	// The AP Rep message.
	APRep *APRep
	// The session key.
	SessionKey types.EncryptionKey
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

	// try load credentials cache.
	cc, ok, err := a.tryLoadCCache(ctx)
	if err != nil {
		// ccache usage was required.
		return nil, err
	}

	if ok {
		// ccache is present.
		if cli, err = client.NewFromCCache(cc, a.Config.KRB5Config,
			a.Config.ClientSettings()...); err != nil {
			return nil, fmt.Errorf("client from ccache: %w", err)
		}
	}

	creds := credentials.New(a.Config.Credential.UserName(), a.Config.Credential.DomainName())

	if cli == nil {
		cli = client.NewWithPassword(creds.UserName(), creds.Realm(), "",
			a.Config.KRB5Config, a.Config.ClientSettings()...)
	}

	if pwd, ok := a.Config.Credential.(credential.Password); ok {
		cli.Credentials = creds.WithPassword(pwd.Password())
	} else if kt, ok := a.Config.Credential.(credential.Keytab); ok {
		cli.Credentials = creds.WithKeytab(kt.Keytab())
	} else if ntHash, ok := a.Config.Credential.(credential.NTHash); ok {
		cli.Credentials = WithNTHash(creds, ntHash.NTHash(), ntHash.KVNO())
		// XXX: add rc4-hmac to allowed etypes.
		cli.Config.LibDefaults.DefaultTGSEnctypeIDs = []int32{etypeID.RC4_HMAC}
		cli.Config.LibDefaults.DefaultTktEnctypeIDs = []int32{etypeID.RC4_HMAC}
		cli.Config.LibDefaults.PermittedEnctypeIDs = []int32{etypeID.RC4_HMAC}
	}

	if _, err := cli.IsConfigured(); err != nil {
		return nil, fmt.Errorf("client configuration: %w", err)
	}

	return cli, nil
}

// makeSecurityService function sets up the security service for
// signing/encryption.
func (a *Authentifier) makeSecurityService(ctx context.Context) error {

	if a.APRep != nil {
		// derive service from mutual-authentication. (dce-style or mutual-authn
		// GSS flags are set.)
		a.state = &SecurityService{
			Key:                    a.APRep.Subkey,
			IsSubKey:               true,
			OutboundSequenceNumber: uint64(a.APReq.SequenceNumber()),
			InboundSequenceNumber:  uint64(a.APRep.SequenceNumber()),
		}
	} else {
		a.state = &SecurityService{
			Key:                    a.SessionKey,
			OutboundSequenceNumber: uint64(a.APReq.SequenceNumber()),
			InboundSequenceNumber:  uint64(a.APReq.SequenceNumber()),
		}
	}

	isServer := true

	clientCipher, err := crypto.NewCipher(ctx, a.state.Key, a.state.IsSubKey, !isServer)
	if err != nil {
		return fmt.Errorf("client cipher: %w", err)
	}
	serverCipher, err := crypto.NewCipher(ctx, a.state.Key, a.state.IsSubKey, isServer)
	if err != nil {
		return fmt.Errorf("server cipher: %w", err)
	}

	if !a.Config.IsServer {
		a.state.InboundCipher, a.state.OutboundCipher = serverCipher, clientCipher
	} else {
		a.state.InboundCipher, a.state.OutboundCipher = clientCipher, serverCipher
	}

	return nil
}

var kvnoErrRe = regexp.MustCompile(`kvno: (\d+)`)

func (a *Authentifier) affirmLogin(ctx context.Context) error {
	if err := a.client.AffirmLogin(); err != nil {
		if _, ok := a.Config.Credential.(credential.NTHash); !ok {
			return err
		}
		// FIXME: retry using kvno from error message.
		if m := kvnoErrRe.FindStringSubmatch(err.Error()); len(m) > 1 {
			kvno, err1 := strconv.ParseInt(m[1], 10, 32)
			if err1 != nil {
				return err
			}
			a.client.Credentials.Keytab().Entries[0].KVNO = uint32(kvno)
			return a.client.AffirmLogin()
		}
	}
	return nil
}

func (a *Authentifier) APRequest(ctx context.Context) ([]byte, error) {

	cli, err := a.makeClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("krb5: init: apreq: make client: %w", err)
	}

	a.client = cli

	if err := a.affirmLogin(ctx); err != nil {
		return nil, fmt.Errorf("krb5: init: apreq: affirm login: %w", err)
	}

	tkt, key, err := a.client.GetServiceTicket(a.Config.SName)
	if err != nil {
		return nil, fmt.Errorf("krb5: init: apreq: get service ticket: %w", err)
	}

	tok, err := spnego.NewKRB5TokenAPREQ(a.client, tkt, key, a.Config.Flags, a.Config.APOptions)
	if err != nil {
		return nil, fmt.Errorf("krb5: init: apreq: call new_krb5_token_apreq: %w", err)
	}

	a.APReq, a.SessionKey = (*APReq)(&tok.APReq), key

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

	a.APRep = &APRep{}

	if !a.Config.DCEStyle {
		tok := &spnego.KRB5Token{}
		if err := tok.Unmarshal(b); err != nil {
			return nil, fmt.Errorf("krb5: init: aprep: unmarshal: dcestyle: %w", err)
		}
		a.APRep.APRep = tok.APRep
	} else {
		if err := a.APRep.Unmarshal(b); err != nil {
			return nil, fmt.Errorf("krb5: init: aprep: unmarshal: %w", err)
		}
	}

	if err := a.APRep.DecryptEncPart(a.SessionKey); err != nil {
		return nil, fmt.Errorf("krb5: init: aprep: decrypt enc part: %w", err)
	}

	b, err := a.APRep.MarshalWithSeqNumber(a.APRep.EncAPRepPart.SequenceNumber, a.SessionKey)
	if err != nil {
		return nil, fmt.Errorf("krb5: init: aprep: marshal enc part: %w", err)
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
