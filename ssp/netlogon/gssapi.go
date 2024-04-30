package netlogon

import (
	"bytes"
	"context"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

var (
	MechanismType = gssapi.OID{1, 2, 752, 43, 14, 2}
)

type Mechanism struct {
	*Authentifier
}

func (Config) Type() gssapi.OID {
	return MechanismType
}

// The mechanism type object identifier.
func (Mechanism) Type() gssapi.OID {
	return MechanismType
}

// DefaultConfig function returns the default config.
func (Mechanism) DefaultConfig(ctx context.Context) (gssapi.MechanismConfig, error) {
	return NewConfig(), nil
}

// New function returns the new mechanism instance from the GSSAPI configuration.
func (Mechanism) New(ctx context.Context) (gssapi.Mechanism, error) {

	var (
		ok bool
	)

	// extract the context.
	cc := gssapi.FromContext(ctx)

	// try get the mechanism config base.
	c, ok := gssapi.GetMechanismConfig(ctx, MechanismType).(*Config)
	if !ok || c == nil {
		// config should have been populated.
		return nil, gssapi.ContextError(ctx, gssapi.NoContext, gssapi.ErrNoContext)
	}

	c.IsServer = cc.IsServer

	if cc.Credential != nil {
		if c.Credential, ok = cc.Credential.Value().(Credential); !ok {
			return nil, gssapi.ContextError(ctx, gssapi.DefectiveCredential, gssapi.ErrDefectiveCredential)
		}
	} else {
		return nil, gssapi.ContextError(ctx, gssapi.DefectiveCredential, gssapi.ErrDefectiveCredential)
	}

	if len(c.ClientChallenge) < 8 {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, fmt.Errorf("client challenge is missing"))
	} else {
		c.ClientChallenge = c.ClientChallenge[:8]
	}

	if len(c.ServerChallenge) < 8 {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, fmt.Errorf("server challenge is missing"))
	} else {
		c.ServerChallenge = c.ServerChallenge[:8]
	}

	return &Mechanism{
		Authentifier: &Authentifier{Config: c},
	}, nil

}

func (m *Mechanism) Init(ctx context.Context, tok *gssapi.Token) (*gssapi.Token, error) {

	if tok.Payload != nil {
		b, err := m.AuthMessageReply(ctx, tok.Payload)
		if err != nil {
			return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
		}

		gssapi.SetAttribute(ctx, gssapi.AttributeSessionKey, m.SessionKey())
		gssapi.SetAttribute(ctx, gssapi.AttributeTarget, m.Config.ServerName)

		return &gssapi.Token{Payload: b}, gssapi.ContextComplete(ctx)
	}

	b, err := m.AuthMessageInit(ctx)
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
	return sz - m.Authentifier.OutboundSignatureSize(ctx, conf)
}

// Wrap token.
func (m *Mechanism) Wrap(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {

	sgn, err := m.WrapOutboundPayload(ctx, [][]byte{tok.Payload}, [][]byte{tok.Payload})
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	return &gssapi.MessageToken{
		QoP:          tok.QoP,
		Capabilities: tok.Capabilities,
		Payload:      tok.Payload,
		Signature:    sgn,
	}, nil
}

// WrapEx function accepts the list of unencrypted payloads and returns the
// encrypted payload and signature.
func (m *Mechanism) WrapEx(ctx context.Context, tokEx *gssapi.MessageTokenEx) (*gssapi.MessageTokenEx, error) {

	forSign, forSeal := [][]byte{}, [][]byte{}
	for _, tok := range tokEx.Payloads {
		if tok.Capabilities.IsSet(gssapi.Confidentiality) {
			forSeal = append(forSeal, tok.Payload)
		}
		if tok.Capabilities.IsSet(gssapi.Integrity) {
			forSign = append(forSign, tok.Payload)
		}
	}

	var err error

	tokEx.Signature, err = m.WrapOutboundPayload(ctx, forSign, forSeal)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	return tokEx, nil
}

// Unwrap token.
func (m *Mechanism) Unwrap(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {

	expSgn, err := m.UnwrapInboundPayload(ctx, [][]byte{tok.Payload}, [][]byte{tok.Payload}, tok.Signature)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if !bytes.Equal(expSgn[:24], expSgn[:24]) { /* compare first 24 bytes only */
		return nil, gssapi.ContextError(ctx, gssapi.BadMIC, gssapi.ErrBadMIC)
	}

	return &gssapi.MessageToken{
		QoP:          tok.QoP,
		Capabilities: tok.Capabilities,
		Payload:      tok.Payload,
		Signature:    tok.Signature,
	}, nil
}

// UnwrapEx function accepts the list of encrypted payloads and signature and
// returns the unencrypted paylaod.
func (m *Mechanism) UnwrapEx(ctx context.Context, tokEx *gssapi.MessageTokenEx) (*gssapi.MessageTokenEx, error) {

	forSign, forSeal := [][]byte{}, [][]byte{}
	for _, tok := range tokEx.Payloads {
		if tok.Capabilities.IsSet(gssapi.Confidentiality) {
			forSeal = append(forSeal, tok.Payload)
		}
		if tok.Capabilities.IsSet(gssapi.Integrity) {
			forSign = append(forSign, tok.Payload)
		}
	}

	expSgn, err := m.UnwrapInboundPayload(ctx, forSign, forSeal, tokEx.Signature)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if !bytes.Equal(expSgn[:24], tokEx.Signature[:24]) { /* compare first 24 bytes only */
		return nil, gssapi.ContextError(ctx, gssapi.BadMIC, gssapi.ErrBadMIC)
	}

	return tokEx, nil
}

// MakeSignature token.
func (m *Mechanism) MakeSignature(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {

	sgn, err := m.MakeOutboundSignature(ctx, [][]byte{tok.Payload})
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
		if tok.Capabilities.IsSet(gssapi.Integrity) || tok.Capabilities.IsSet(gssapi.Confidentiality) {
			forSign = append(forSign, tok.Payload)
		}
	}

	sgn, err := m.MakeOutboundSignature(ctx, forSign)
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

	expSgn, err := m.MakeInboundSignature(ctx, [][]byte{tok.Payload})
	if err != nil {
		return gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if !bytes.Equal(tok.Signature[:12], expSgn[:12]) { /* compare first 12 bytes */
		return gssapi.ContextError(ctx, gssapi.BadMIC, gssapi.ErrBadMIC)
	}

	return gssapi.ContextComplete(ctx)
}

// VerifySignatureEx token.
func (m *Mechanism) VerifySignatureEx(ctx context.Context, tokEx *gssapi.MessageTokenEx) error {

	forSign := [][]byte{}

	for _, tok := range tokEx.Payloads {
		if tok.Capabilities.IsSet(gssapi.Integrity) || tok.Capabilities.IsSet(gssapi.Confidentiality) {
			forSign = append(forSign, tok.Payload)
		}
	}

	expSgn, err := m.MakeInboundSignature(ctx, forSign)
	if err != nil {
		return gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if !bytes.Equal(tokEx.Signature[:12], expSgn[:12]) { /* compare first 12 bytes */
		return gssapi.ContextError(ctx, gssapi.BadMIC, gssapi.ErrBadMIC)
	}

	return gssapi.ContextComplete(ctx)
}

var (
	_ gssapi.Mechanism   = (*Mechanism)(nil)
	_ gssapi.MechanismEx = (*Mechanism)(nil)
)
