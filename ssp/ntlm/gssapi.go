package ntlm

import (
	"bytes"
	"context"
	"errors"
	"time"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

var (
	MechanismType = gssapi.OID{1, 3, 6, 1, 4, 1, 311, 2, 2, 10}
)

// The NTLMSSP GSS API Mechanism.
type Mechanism struct {
	*Authentifier
}

// Config binding for GSS.
func (Config) Type() gssapi.OID {
	return MechanismType
}

// DefaultConfig function returns the default config.
func (Mechanism) DefaultConfig(ctx context.Context) (gssapi.MechanismConfig, error) {
	return NewConfig(), nil
}

var (
	ErrMutualAuthnNotSupported = errors.New("mutual authn is not supported")
	ErrDelegationNotSupported  = errors.New("delegation is not supported")
)

// New function returns the new mechanism instance from the GSSAPI configuration.
func (Mechanism) New(ctx context.Context) (gssapi.Mechanism, error) {

	var (
		ok  bool
		err error
	)

	// extract the context.
	cc := gssapi.FromContext(ctx)

	// try get the mechanism config base.
	c, ok := gssapi.GetMechanismConfig(ctx, MechanismType).(*Config)
	if !ok || c == nil {
		// config should have been populated.
		return nil, gssapi.ContextError(ctx, gssapi.NoContext, gssapi.ErrNoContext)
	}
	if cc.Credential != nil {
		if c.Credential, ok = cc.Credential.Value().(Credential); !ok || !IsValidCredential(c.Credential) {
			return nil, gssapi.ContextError(ctx, gssapi.DefectiveCredential, gssapi.ErrDefectiveCredential)
		}
	}
	if cc.Capabilities.IsSet(gssapi.MutualAuthn) {
		return nil, gssapi.ContextError(ctx, gssapi.Unavailable, ErrMutualAuthnNotSupported)
	}
	if cc.Capabilities.IsSet(gssapi.Delegation) {
		return nil, gssapi.ContextError(ctx, gssapi.Unavailable, ErrDelegationNotSupported)
	}
	if cc.Capabilities.IsSet(gssapi.Integrity | gssapi.ReplayDetection | gssapi.Sequencing) {
		c.Integrity = true
	}
	if cc.Capabilities.IsSet(gssapi.Confidentiality) {
		c.Confidentiality = true
	}
	if cc.Capabilities.IsSet(gssapi.Identify) {
		c.Identify = true
	}
	if cc.Capabilities.IsSet(gssapi.Anonymity) {
		c.Anonymity = true
	}
	if cc.Capabilities.IsSet(gssapi.Datagram) {
		c.Datagram = true
	}
	if cc.ContextTTL != 0 {
		c.MaxLifetime = time.Duration(cc.ContextTTL)
	}
	if cc.TargetName != "" {
		c.TargetName = cc.TargetName
	}
	if cc.TargetNameFromUntrustedSource {
		c.UnverifiedTargetName = cc.TargetNameFromUntrustedSource
	}
	if cc.ChannelBindings != nil {
		if c.ChannelBindings, err = cc.ChannelBindings.Marshal(); err != nil {
			return nil, gssapi.ContextError(ctx, gssapi.BadBindings, gssapi.ErrBadBindings)
		}
	}
	return &Mechanism{
		Authentifier: &Authentifier{Config: c},
	}, nil
}

var (
	_ gssapi.MechanismFactory = (*Mechanism)(nil)
	_ gssapi.Mechanism        = (*Mechanism)(nil)
)

// The mechanism type object identifier.
func (Mechanism) Type() gssapi.OID {
	return MechanismType
}

// The security context init call.
func (m *Mechanism) Init(ctx context.Context, tok *gssapi.Token) (*gssapi.Token, error) {

	if tok.Payload != nil {
		b, err := m.Authenticate(ctx, tok.Payload)
		if err != nil {
			return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
		}

		gssapi.SetAttribute(ctx, gssapi.AttributeSessionKey, m.SessionKey())
		gssapi.SetAttribute(ctx, gssapi.AttributeTarget, m.TargetName())

		return &gssapi.Token{Payload: b}, gssapi.ContextComplete(ctx)
	}

	b, err := m.Negotiate(ctx)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	return &gssapi.Token{Payload: b}, gssapi.ContextContinueNeeded(ctx)
}

// The security context accept call.
func (m *Mechanism) Accept(ctx context.Context, tok *gssapi.Token) (*gssapi.Token, error) {
	return nil, gssapi.ContextError(ctx, gssapi.Unavailable, gssapi.ErrUnavailable)
}

// The maximum message size for the given limit. (and flag determining if
// conf is required).
func (m *Mechanism) WrapSizeLimit(ctx context.Context, sz int, conf bool) int {
	return sz - 16
}

// WrapEx function accepts the list of unencrypted payloads and returns the
// encrypted payload and signature.
func (m *Mechanism) WrapEx(ctx context.Context, tokEx *gssapi.MessageTokenEx) (*gssapi.MessageTokenEx, error) {

	forChk := [][]byte{}

	for _, tok := range tokEx.Payloads {
		// XXX: NTLM uses all input buffers disregarding of flags.
		if !m.Config.NoSignAllBuffers || tok.Capabilities.IsSet(gssapi.Integrity) {
			forChk = append(forChk, tok.Payload)
		}
	}

	// compute checksum.
	chk, err := m.MakeOutboundChecksum(ctx, forChk)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	// encrypt data.
	for _, tok := range tokEx.Payloads {
		if !m.Config.Confidentiality {
			// unset conf_req_flag.
			tok.Capabilities &= ^gssapi.Confidentiality
		}
		if tok.Capabilities.IsSet(gssapi.Confidentiality) {
			// encrypt the payload.
			if err = m.ApplyOutboundCipher(ctx, tok.Payload); err != nil {
				return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
			}
		}
	}

	// make signature.
	sgn, err := m.MakeOutboundSignature(ctx, chk)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	return &gssapi.MessageTokenEx{
		Payloads:  tokEx.Payloads,
		QoP:       tokEx.QoP,
		Signature: sgn,
	}, nil
}

// UnwrapEx function accepts the list of encrypted payloads and signature and
// returns the unencrypted paylaod.
func (m *Mechanism) UnwrapEx(ctx context.Context, tokEx *gssapi.MessageTokenEx) (*gssapi.MessageTokenEx, error) {

	// decrypt data.
	for _, tok := range tokEx.Payloads {
		if !m.Config.Confidentiality {
			// unset conf_req_flag.
			tok.Capabilities &= ^gssapi.Confidentiality
		}
		if tok.Capabilities.IsSet(gssapi.Confidentiality) {
			// encrypt the payload.
			if err := m.ApplyInboundCipher(ctx, tok.Payload); err != nil {
				return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
			}
		}
	}

	forChk := [][]byte{}

	for _, tok := range tokEx.Payloads {
		// XXX: NTLM uses all input buffers disregarding of flags.
		if !m.Config.NoSignAllBuffers || tok.Capabilities.IsSet(gssapi.Integrity) {
			forChk = append(forChk, tok.Payload)
		}
	}

	// compute checksum.
	chk, err := m.MakeInboundChecksum(ctx, forChk)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	// make signature.
	sgn, err := m.MakeInboundSignature(ctx, chk)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	return &gssapi.MessageTokenEx{
		Payloads:  tokEx.Payloads,
		QoP:       tokEx.QoP,
		Signature: sgn,
	}, nil
}

// Wrap token.
func (m *Mechanism) Wrap(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {

	if !m.Config.Confidentiality {
		// clear conf_req_flag.
		tok.Capabilities &= ^gssapi.Confidentiality
	}

	// compute checksum.
	chk, err := m.MakeOutboundChecksum(ctx, [][]byte{tok.Payload})
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if tok.Capabilities.IsSet(gssapi.Confidentiality) {
		// encrypt the payload.
		if err = m.ApplyOutboundCipher(ctx, tok.Payload); err != nil {
			return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
		}
	}

	// make signature.
	sgn, err := m.MakeOutboundSignature(ctx, chk)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	return &gssapi.MessageToken{
		QoP:          tok.QoP,
		Capabilities: tok.Capabilities,
		Payload:      tok.Payload,
		Signature:    sgn,
	}, gssapi.ContextComplete(ctx)
}

// Unwrap token.
func (m *Mechanism) Unwrap(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {

	if !m.Config.Confidentiality {
		// clear conf_req_flag.
		tok.Capabilities &= ^gssapi.Confidentiality
	}

	// decrypt the payload.
	if tok.Capabilities.IsSet(gssapi.Confidentiality) {
		if err := m.ApplyInboundCipher(ctx, tok.Payload); err != nil {
			return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
		}
	}

	// compute checksum.
	expChk, err := m.MakeInboundChecksum(ctx, [][]byte{tok.Payload})
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	// make signature.
	expSgn, err := m.MakeInboundSignature(ctx, expChk)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if !bytes.Equal(tok.Signature, expSgn) {
		return nil, gssapi.ContextError(ctx, gssapi.BadMIC, gssapi.ErrBadMIC)
	}

	return &gssapi.MessageToken{
		QoP:          tok.QoP,
		Capabilities: tok.Capabilities,
		Payload:      tok.Payload,
		Signature:    tok.Signature,
	}, gssapi.ContextComplete(ctx)
}

// MakeSignature token.
func (m *Mechanism) MakeSignature(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {

	chk, err := m.MakeOutboundChecksum(ctx, [][]byte{tok.Payload})
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	sgn, err := m.MakeOutboundSignature(ctx, chk)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	return &gssapi.MessageToken{
		QoP:          tok.QoP,
		Capabilities: tok.Capabilities,
		Payload:      tok.Payload,
		Signature:    sgn,
	}, gssapi.ContextComplete(ctx)
}

// MakeSignature token.
func (m *Mechanism) MakeSignatureEx(ctx context.Context, tokEx *gssapi.MessageTokenEx) (*gssapi.MessageTokenEx, error) {

	forSign := [][]byte{}

	for _, tok := range tokEx.Payloads {
		// XXX: NTLM uses all input buffers disregarding of flags.
		if !m.Config.NoSignAllBuffers || tok.Capabilities.IsSet(gssapi.Integrity) || tok.Capabilities.IsSet(gssapi.Confidentiality) {
			forSign = append(forSign, tok.Payload)
		}
	}

	chk, err := m.MakeOutboundChecksum(ctx, forSign)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	sgn, err := m.MakeOutboundSignature(ctx, chk)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	return &gssapi.MessageTokenEx{
		QoP:       tokEx.QoP,
		Payloads:  tokEx.Payloads,
		Signature: sgn,
	}, gssapi.ContextComplete(ctx)
}

// VerifySignature token.
func (m *Mechanism) VerifySignature(ctx context.Context, tok *gssapi.MessageToken) error {

	expChk, err := m.MakeInboundChecksum(ctx, [][]byte{tok.Payload})
	if err != nil {
		return gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	expSgn, err := m.MakeInboundSignature(ctx, expChk)
	if err != nil {
		return gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if !bytes.Equal(tok.Signature, expSgn) {
		return gssapi.ContextError(ctx, gssapi.BadMIC, gssapi.ErrBadMIC)
	}

	return gssapi.ContextComplete(ctx)
}

// VerifySignatureEx token.
func (m *Mechanism) VerifySignatureEx(ctx context.Context, tokEx *gssapi.MessageTokenEx) error {

	forSign := [][]byte{}

	for _, tok := range tokEx.Payloads {
		// XXX: NTLM uses all input buffers disregarding of flags.
		if !m.Config.NoSignAllBuffers || tok.Capabilities.IsSet(gssapi.Integrity) || tok.Capabilities.IsSet(gssapi.Confidentiality) {
			forSign = append(forSign, tok.Payload)
		}
	}

	expChk, err := m.MakeInboundChecksum(ctx, forSign)
	if err != nil {
		return gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	expSgn, err := m.MakeInboundSignature(ctx, expChk)
	if err != nil {
		return gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if !bytes.Equal(tokEx.Signature, expSgn) {
		return gssapi.ContextError(ctx, gssapi.BadMIC, gssapi.ErrBadMIC)
	}

	return gssapi.ContextComplete(ctx)
}

var (
	_ gssapi.Mechanism   = (*Mechanism)(nil)
	_ gssapi.MechanismEx = (*Mechanism)(nil)
)
