package ntlm

import (
	"bytes"
	"context"
	"fmt"
	"hash"

	"github.com/oiweiwei/go-msrpc/ssp/crypto"
)

// The NTLM Authentifier.
type Authentifier struct {
	// The authentifier configuration.
	Config *Config
	// The list of messages for MIC computation.
	mic bytes.Buffer
	// The cryptographic state of the NTLM.
	state *SecurityService
	// The negotiated session parameters.
	session *SecurityParameters
}

func (a *Authentifier) SessionKey() []byte {
	if a.state != nil {
		return a.state.ExportedSessionKey
	}
	return nil
}

func (a *Authentifier) TargetName() string {
	if a.session != nil {
		if a.session.DomainName != "" {
			return a.session.DomainName
		} else if a.session.ServerName != "" {
			return a.session.ServerName
		}
	}
	return ""
}

func (a *Authentifier) Reset() {
	a.state, a.session = nil, nil
	if a.mic.Reset(); a.Config == nil {
		a.Config = &Config{}
	}
}

// The authentifier state structure represents the negotiated
// authentication state.
type SecurityService struct {
	// The exported session key.
	ExportedSessionKey []byte
	// The client signing key.
	ClientSignKey []byte
	// The client sealing key.
	ClientSealKey []byte
	// The server signing key.
	ServerSignKey []byte
	// The server sealing key.
	ServerSealKey []byte
	// The client cipher.
	OutboundCipher *Cipher
	// The server cipher.
	InboundCipher *Cipher
	// The sequence number to send.
	OutboundSequenceNumber uint32
	// The sequence number to receive.
	InboundSequenceNumber uint32
}

// Negotiate function returns the initial token.
func (a *Authentifier) Negotiate(ctx context.Context) ([]byte, error) {

	a.Reset()

	nm := &NegotiateMessage{
		Negotiate: a.Config.Negotiate(),
		Version:   a.Config.Version,
	}

	if a.Config.Credential != nil {
		nm.DomainName = a.Config.Credential.DomainName()
		nm.Workstation = a.Config.Credential.Workstation()
	}

	if nm.Negotiate.IsSet(NegotiateVersion) {
		if nm.Version.IsZero() {
			nm.Version = DefaultVersion
		}
		// clear domain/workstation if version is set.
		nm.DomainName = ""
		nm.Workstation = ""
	}

	if nm.DomainName != "" {
		nm.Negotiate = nm.Negotiate.Set(NegotiateOEMDomainSupplied)
	}
	if nm.Workstation != "" {
		nm.Negotiate = nm.Negotiate.Set(NegotiateOEMWorkstationSupplied)
	}

	b, err := nm.Marshal(ctx)
	if err != nil {
		return nil, fmt.Errorf("ntlm: init: negotiate: %w", err)
	}

	a.mic.Write(b)

	return b, nil
}

// Authenticate function accepts the token containing the ChallengeMessage and returns
// the AuthenticateMessage token.
func (a *Authentifier) Authenticate(ctx context.Context, b []byte) ([]byte, error) {

	cm := &ChallengeMessage{}

	// unmarshal challenge message.
	if err := cm.Unmarshal(ctx, b); err != nil {
		return nil, fmt.Errorf("ntlm: init: authenticate: unmarshal challenge message: %w", err)
	}

	// check negotiated parameters.
	if err := a.Config.Verify(cm.Negotiate); err != nil {
		return nil, fmt.Errorf("ntlm: init: authenticate: %w", err)
	}

	// set session parameters.
	a.makeSecurityParameters(ctx, cm)

	ntlm := a.Config.NewNTLMVersion(ctx, a.Config, a.session)

	// generate nonce.
	clientChallenge, err := crypto.Nonce(8)
	if err != nil {
		return nil, fmt.Errorf("ntlm: init: authenticate: generate nonce: %w", err)
	}

	// compute challenge response.
	resp, err := ntlm.ChallengeResponse(ctx, a.Config.Credential, cm, clientChallenge)
	if err != nil {
		return nil, fmt.Errorf("ntlm: init: authenticate: compute challenge response: %w", err)
	}

	// compute key-exchange-key.
	if resp.KeyExchangeKey, err = ntlm.KeyExchangeKey(ctx, cm, resp); err != nil {
		return nil, fmt.Errorf("ntlm: init: authenticate: compute key-exchange-key: %w", err)
	}

	// append challenge message to buffer.
	a.mic.Write(b)

	am := &AuthenticateMessage{
		Negotiate:           cm.Negotiate,
		LMChallengeResponse: resp.LM,
		NTChallengeResponse: resp.NT,
		MIC:                 make([]byte, 16),
	}

	if a.Config.Credential != nil {
		am.DomainName = a.Config.Credential.DomainName()
		am.UserName = a.Config.Credential.UserName()
		am.Workstation = a.Config.Credential.Workstation()
	}

	var exportedKey []byte

	if a.session.KeyExchange {
		if exportedKey, err = crypto.Nonce(16); err != nil {
			return nil, fmt.Errorf("ntlm: init: authenticate: compute exported session key: %w", err)
		}
		if am.EncryptedRandomSessionKey, err = RC4K(ctx, resp.KeyExchangeKey, exportedKey); err != nil {
			return nil, fmt.Errorf("ntlm: init: authenticate: compute encrypted random session key: %w", err)
		}
	} else {
		exportedKey = resp.KeyExchangeKey
	}

	// initialize sealing and signing.
	if err = a.makeSecurityService(ctx, exportedKey); err != nil {
		return nil, fmt.Errorf("ntlm: init: authenticate: initialize authenticator state: %w", err)
	}

	b, err = am.Marshal(ctx)
	if err != nil {
		return nil, fmt.Errorf("ntlm: init: authenticate: marshal response: %w", err)
	}

	// append message to the buffer.
	a.mic.Write(b)

	if resp.RequestMIC {
		// calculate mic.
		if am.MIC, err = crypto.HMAC_MD5(exportedKey, a.mic.Bytes()); err != nil {
			return nil, fmt.Errorf("ntlm: init: authenticate: compute mic: %w", err)
		}
		// copy mic value.
		copy(b[micidx:], am.MIC)
		// reset the mic buffer.
		a.mic.Reset()
	}

	return b, nil
}

// The MIC index in the output buffer.
const micidx = 72

var (
	// client-to-server seal.
	magicSealC2S = []byte("session key to client-to-server sealing key magic constant\x00")
	// server-to-client seal.
	magicSealS2C = []byte("session key to server-to-client sealing key magic constant\x00")
	// client-to-server sign.
	magicSignC2S = []byte("session key to client-to-server signing key magic constant\x00")
	// server-to-client sign.
	magicSignS2C = []byte("session key to server-to-client signing key magic constant\x00")
)

func (a *Authentifier) ResetSecurityService(ctx context.Context) error {

	if err := a.makeSecurityService(ctx, a.state.ExportedSessionKey); err != nil {
		return err
	}

	return nil
}

// makeSecurityService function initializes the signing and sealing data structures.
func (a *Authentifier) makeSecurityService(ctx context.Context, key []byte) error {

	var err error

	if a.state == nil {
		a.state = &SecurityService{ExportedSessionKey: key}
	}

	// client signing key.
	if a.state.ClientSignKey, err = a.makeSignKey(ctx, key, magicSignC2S); err != nil {
		return err
	}
	// server signing key.
	if a.state.ServerSignKey, err = a.makeSignKey(ctx, key, magicSignS2C); err != nil {
		return err
	}
	// client seal key.
	if a.state.ClientSealKey, err = a.makeSealKey(ctx, key, magicSealC2S); err != nil {
		return err
	}
	// server seal key.
	if a.state.ServerSealKey, err = a.makeSealKey(ctx, key, magicSealS2C); err != nil {
		return err
	}

	// client hmac-md5/crc32 checksum function.
	clientHash := a.makeHashFunc(ctx, a.state.ClientSignKey)

	// client rc4 cipher.
	clientCipher, err := NewCipher(ctx, a.state.ClientSealKey, clientHash)
	if err != nil {
		return err
	}

	// server hmac-md5/crc32 checksum function.
	serverHash := a.makeHashFunc(ctx, a.state.ServerSignKey)

	// server rc4 cipher.
	serverCipher, err := NewCipher(ctx, a.state.ServerSealKey, serverHash)
	if err != nil {
		return err
	}

	// assign proper inbound/outbound ciphers.
	if !a.Config.IsServer {
		a.state.OutboundCipher, a.state.InboundCipher = clientCipher, serverCipher
	} else {
		a.state.OutboundCipher, a.state.InboundCipher = serverCipher, clientCipher
	}

	return nil

}

// makeSignKey creates the key for the signature.
func (a *Authentifier) makeSignKey(ctx context.Context, key []byte, magic []byte) ([]byte, error) {
	if !a.session.ExtendedSessionSecurity {
		return nil, nil
	}
	return crypto.MD5(key, magic)
}

type SecurityParameters struct {
	ExtendedSessionSecurity bool
	UseLMKey                bool
	NonNTSessionKey         bool
	Datagram                bool
	KeyExchange             bool
	KeySize                 int
	ServerName              string
	DomainName              string
}

// makeHashFunc creates the hash function.
func (a *Authentifier) makeHashFunc(ctx context.Context, key []byte) func(uint32) hash.Hash {
	if a.session.ExtendedSessionSecurity {
		return HMACMD5(key)
	}
	return CRC32(key)
}

// makeSealKey creates the key for the encryption.
func (a *Authentifier) makeSealKey(ctx context.Context, key []byte, magic []byte) ([]byte, error) {

	if a.session.ExtendedSessionSecurity {
		switch a.session.KeySize {
		case 128:
			// no-op.
		case 56:
			key = key[:7]
		default:
			key = key[:5]
		}
		return crypto.MD5(key, magic)
	}

	if a.session.UseLMKey || a.session.Datagram {
		if a.session.KeySize == 56 {
			return append(key[:7], 0xa0), nil
		}
		return append(key[:5], 0xe5, 0x38, 0xb0), nil
	}

	return key, nil
}

const (
	// The signature version.
	SignatureVersion = 0x01
)

func (a *Authentifier) makeSecurityParameters(ctx context.Context, cm *ChallengeMessage) {
	a.session = &SecurityParameters{
		ExtendedSessionSecurity: cm.Negotiate.IsSet(NegotiateExtendedSessionSecurity),
		UseLMKey:                cm.Negotiate.IsSet(NegotiateLMKey),
		NonNTSessionKey:         cm.Negotiate.IsSet(RequestNonNTSessionKey),
		KeyExchange:             cm.Negotiate.IsSet(NegotiateKeyExchange),
		Datagram:                cm.Negotiate.IsSet(NegotiateDatagram),
		KeySize:                 40,
	}

	if cm.Negotiate.IsSet(TargetTypeDomain) {
		a.session.DomainName = cm.TargetName
	} else if cm.Negotiate.IsSet(TargetTypeServer) {
		a.session.ServerName = cm.TargetName
	}

	if cm.Negotiate.IsSet(Negotiate128) {
		a.session.KeySize = 128
	} else if cm.Negotiate.IsSet(Negotiate56) {
		a.session.KeySize = 56
	}
}

func (a *Authentifier) ApplyInboundCipher(ctx context.Context, b []byte) error {
	if err := a.state.InboundCipher.XORKeyStream(b); err != nil {
		return fmt.Errorf("ntlm: apply inbound cipher: %w", err)
	}
	return nil
}

func (a *Authentifier) ApplyOutboundCipher(ctx context.Context, b []byte) error {
	if err := a.state.OutboundCipher.XORKeyStream(b); err != nil {
		return fmt.Errorf("ntlm: apply outbound cipher: %w", err)
	}
	return nil
}

func (a *Authentifier) MakeInboundSignature(ctx context.Context, checksum []byte) ([]byte, error) {
	sgn, err := a.makeSignature(ctx, a.state.InboundCipher, checksum, &a.state.InboundSequenceNumber)
	if err != nil {
		return nil, fmt.Errorf("ntlm: make inbound signature: %w", err)
	}
	return sgn, nil
}

func (a *Authentifier) MakeOutboundSignature(ctx context.Context, checksum []byte) ([]byte, error) {
	sgn, err := a.makeSignature(ctx, a.state.OutboundCipher, checksum, &a.state.OutboundSequenceNumber)
	if err != nil {
		return nil, fmt.Errorf("ntlm: make outbound signature: %w", err)
	}
	return sgn, nil
}

func (a *Authentifier) MakeInboundChecksum(ctx context.Context, b [][]byte) ([]byte, error) {
	chk, err := a.makeChecksum(ctx, a.state.InboundCipher, a.state.InboundSequenceNumber, b...)
	if err != nil {
		return nil, fmt.Errorf("ntlm: make inbound checksum: %w", err)
	}
	return chk, nil
}

func (a *Authentifier) MakeOutboundChecksum(ctx context.Context, b [][]byte) ([]byte, error) {
	chk, err := a.makeChecksum(ctx, a.state.OutboundCipher, a.state.OutboundSequenceNumber, b...)
	if err != nil {
		return nil, fmt.Errorf("ntlm: make outbound checksum: %w", err)
	}
	return chk, nil
}

func (a *Authentifier) makeChecksum(ctx context.Context, cipher *Cipher, seqNum uint32, b ...[]byte) ([]byte, error) {
	if !a.Config.Integrity {
		return nil, nil
	}
	return cipher.Checksum(seqNum, b...)
}

// makeSignature function creates the NTLM signature.
func (a *Authentifier) makeSignature(ctx context.Context, cipher *Cipher, checksum []byte, seqNum *uint32) ([]byte, error) {

	var err error

	if !a.Config.Integrity {
		// signing is not negotiated.
		return (&Signature{Version: SignatureVersion}).Marshal(ctx)
	}

	if !a.session.ExtendedSessionSecurity {

		// extended session security is not negotiated.

		// basically this should not work because, a:
		//
		// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/524cdccb-563e-4793-92b0-7bc321fce096:
		// If extended session security is not negotiated (section 2.2.2.5),
		// then no signing keys are available and message signing is not supported.
		//
		// but at the same time, there is an algorithm to produce a signature:
		// (possibly this is applicable for NTLMv1).
		//
		// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/0b1fb6a6-7224-4d5b-af35-fdd45c0791e5

		s := &Signature{
			Version:  SignatureVersion,
			Checksum: checksum,
			SeqNum:   *seqNum,
		}

		// random_pad + checksum.
		if err = cipher.XORKeyStream(s.Checksum); err != nil {
			return nil, fmt.Errorf("random_pad: xor_key_stream: %w", err)
		}
		if err = cipher.XORKeyStream(&s.SeqNum); err != nil {
			return nil, fmt.Errorf("seq_num: xor_key_stream: %w", err)
		}

		// clear random pad.
		copy(s.Checksum[:4], make([]byte, 4))

		if !a.Config.Datagram {
			*seqNum++
		}

		return s.Marshal(ctx)
	}

	s := Signature{
		Version:  SignatureVersion,
		Checksum: checksum,
		SeqNum:   *seqNum,
	}

	if a.session.KeyExchange {
		// key-exchange negotiated.
		if err = cipher.XORKeyStream(s.Checksum); err != nil {
			return nil, fmt.Errorf("checksum: key_exchange: xor_key_stream: %w", err)
		}
	}

	if !a.Config.Datagram {
		*seqNum++
	}

	return s.Marshal(ctx)
}
