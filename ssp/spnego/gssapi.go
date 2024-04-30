package spnego

import (
	"context"
	"encoding/asn1"
	"errors"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

type Mechanism struct {
	*Authentifier
}

var (
	MechanismTypeSPNEGO = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 2}
)

func (Mechanism) Type() gssapi.OID {
	return (gssapi.OID)(MechanismTypeSPNEGO)
}

// The configuration type.
func (Config) Type() gssapi.OID {
	return (gssapi.OID)(MechanismTypeSPNEGO)
}

func (Mechanism) DefaultConfig(ctx context.Context) (gssapi.MechanismConfig, error) {

	config := &Config{}

	for _, m := range gssapi.ListMechanisms(ctx) {
		if m.Type().Equal((gssapi.OID)(MechanismTypeSPNEGO)) {
			continue
		}
		config.MechanismsList = append(config.MechanismsList, m)
	}

	return config, nil
}

func (Mechanism) New(ctx context.Context) (gssapi.Mechanism, error) {

	// extract the context.
	cc := gssapi.FromContext(ctx)

	// try get the mechanism config base.
	c, ok := gssapi.GetMechanismConfig(ctx, (gssapi.OID)(MechanismTypeSPNEGO)).(*Config)
	if !ok {
		// config should have been populated.
		return nil, gssapi.ContextError(ctx, gssapi.NoContext, gssapi.ErrNoContext)
	}

	// set capabilities.
	c.Capabilities = cc.Capabilities

	// check that mechanism list is not empty.
	if len(c.MechanismsList) == 0 {
		return nil, gssapi.ContextError(ctx, gssapi.Unavailable, gssapi.ErrUnavailable)
	}
	return &Mechanism{
		Authentifier: &Authentifier{
			Config: c,
		},
	}, nil
}

var (
	ErrReject = errors.New("spnego: rejected")
)

func (m *Mechanism) Init(ctx context.Context, tok *gssapi.Token) (*gssapi.Token, error) {

	var err error

	if tok.Payload == nil {

		b, err := m.Negotiate(ctx)
		if err != nil {
			return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
		}

		return &gssapi.Token{
			Payload: b,
		}, gssapi.ContextContinueNeeded(ctx)
	}

	b, err := m.Respond(ctx, tok.Payload)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if b != nil {
		return &gssapi.Token{
			Payload: b,
		}, gssapi.ContextContinueNeeded(ctx)
	}

	return nil, gssapi.ContextComplete(ctx)

}

func (m *Mechanism) Accept(ctx context.Context, tok *gssapi.Token) (*gssapi.Token, error) {
	return nil, nil
}

func (m *Mechanism) WrapSizeLimit(ctx context.Context, sz int, conf bool) int {
	return m.Mechanism.WrapSizeLimit(ctx, sz, conf)
}

// Wrap function.
func (m *Mechanism) Wrap(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {
	return m.Mechanism.Wrap(ctx, tok)
}

// Unwrap function.
func (m *Mechanism) Unwrap(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {
	return m.Mechanism.Unwrap(ctx, tok)
}

// MakeSignature function.
func (m *Mechanism) MakeSignature(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {
	return m.Mechanism.MakeSignature(ctx, tok)
}

// VerifySignature function.
func (m *Mechanism) VerifySignature(ctx context.Context, tok *gssapi.MessageToken) error {
	return m.Mechanism.VerifySignature(ctx, tok)
}

// WrapEx function.
func (m *Mechanism) WrapEx(ctx context.Context, tok *gssapi.MessageTokenEx) (*gssapi.MessageTokenEx, error) {
	mechEx, ok := (interface{})(m.Mechanism).(gssapi.MechanismEx)
	if !ok {
		return nil, gssapi.ContextError(ctx, gssapi.Unavailable, gssapi.ErrUnavailable)
	}
	return mechEx.WrapEx(ctx, tok)
}

// UnwrapEx function.
func (m *Mechanism) UnwrapEx(ctx context.Context, tok *gssapi.MessageTokenEx) (*gssapi.MessageTokenEx, error) {
	mechEx, ok := (interface{})(m.Mechanism).(gssapi.MechanismEx)
	if !ok {
		return nil, gssapi.ContextError(ctx, gssapi.Unavailable, gssapi.ErrUnavailable)
	}
	return mechEx.UnwrapEx(ctx, tok)
}

// MakeSignatureEx function.
func (m *Mechanism) MakeSignatureEx(ctx context.Context, tok *gssapi.MessageTokenEx) (*gssapi.MessageTokenEx, error) {
	mechEx, ok := (interface{})(m.Mechanism).(gssapi.MechanismEx)
	if !ok {
		return nil, gssapi.ContextError(ctx, gssapi.Unavailable, gssapi.ErrUnavailable)
	}
	return mechEx.MakeSignatureEx(ctx, tok)
}

// VerifySignatureEx function.
func (m *Mechanism) VerifySignatureEx(ctx context.Context, tok *gssapi.MessageTokenEx) error {
	mechEx, ok := (interface{})(m.Mechanism).(gssapi.MechanismEx)
	if !ok {
		return gssapi.ContextError(ctx, gssapi.Unavailable, gssapi.ErrUnavailable)
	}
	return mechEx.VerifySignatureEx(ctx, tok)
}

var (
	_ gssapi.Mechanism   = (*Mechanism)(nil)
	_ gssapi.MechanismEx = (*Mechanism)(nil)
)
