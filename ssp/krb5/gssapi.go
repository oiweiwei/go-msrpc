package krb5

import (
	"bytes"
	"context"

	"github.com/jcmturner/gokrb5/v8/iana/flags"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

var (
	MechanismType = gssapi.OID{1, 2, 840, 113554, 1, 2, 2}
)

type Mechanism struct {
	*Authentifier
}

func (Config) Type() gssapi.OID {
	return MechanismType
}

func (c *Config) Copy() gssapi.MechanismConfig {

	cp := *c

	if c.KRB5Config != nil {
		// shallow copy.
		cfg := *c.KRB5Config
		cp.KRB5Config = &cfg
	}

	cp.Flags = make([]int, len(c.Flags))
	copy(cp.Flags, c.Flags)

	cp.APOptions = make([]int, len(c.APOptions))
	copy(cp.APOptions, c.APOptions)

	return &cp
}

// DefaultConfig function returns the default config.
func (Mechanism) DefaultConfig(ctx context.Context) (gssapi.MechanismConfig, error) {
	return NewConfig(), nil
}

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

	if c.KRB5Config == nil {
		// load kerberos config.
		if c.KRB5Config, err = LoadKRB5Conf(c.KRB5ConfigPath); err != nil {
			return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
		}
	}

	c.IsServer = cc.IsServer

	if cc.Credential != nil {
		if c.Credential, ok = cc.Credential.Value().(Credential); !ok || !IsValidCredential(c.Credential) {
			return nil, gssapi.ContextError(ctx, gssapi.DefectiveCredential, gssapi.ErrDefectiveCredential)
		}
	}

	if cc.TargetName != "" {
		c.SName = cc.TargetName
	}

	if cc.Capabilities.IsSet(gssapi.Anonymity) {
		c.Flags = append(c.Flags, int(gssapi.Anonymity))
	}

	if cc.Capabilities.IsSet(gssapi.MutualAuthn) || c.FlagIsSet(gssapi.MutualAuthn) {
		c.APOptions = append(c.APOptions, flags.APOptionMutualRequired)
	}

	if cc.Capabilities.IsSet(gssapi.ReplayDetection) {
		c.Flags = append(c.Flags, int(gssapi.ReplayDetection))
	}

	if cc.Capabilities.IsSet(gssapi.Sequencing) {
		c.Flags = append(c.Flags, int(gssapi.Sequencing))
	}

	if cc.Capabilities.IsSet(gssapi.Integrity) {
		c.Flags = append(c.Flags, int(gssapi.Integrity))
	}

	if cc.Capabilities.IsSet(gssapi.Confidentiality) {
		c.Flags = append(c.Flags, int(gssapi.Confidentiality))
	}

	if cc.Capabilities.IsSet(gssapi.DCEStyle) || c.DCEStyle {
		c.Flags, c.DCEStyle = append(c.Flags, int(gssapi.DCEStyle)), true
	}

	// XXX: this option controls whether the key from TGT must be
	// used to decrypt the contents.
	// c.APOptions = append(c.APOptions, flags.APOptionUseSessionKey)

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

	if tok.Payload != nil || m.APReq != nil /* non-dce style auth */ {

		b, err := m.APReply(ctx, tok.Payload)
		if err != nil {
			return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
		}

		gssapi.SetAttribute(ctx, gssapi.AttributeSessionKey, m.ExportedSessionKey)
		gssapi.SetAttribute(ctx, gssapi.AttributeTarget, m.Config.SName)

		if !m.Config.DCEStyle && m.Config.FlagIsSet(gssapi.MutualAuthn) {
			// return empty apreply for non-dce style mutual authentication.
			return &gssapi.Token{}, gssapi.ContextComplete(ctx)
		}

		return &gssapi.Token{Payload: b}, gssapi.ContextComplete(ctx)
	}

	b, err := m.APRequest(ctx)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if !m.Config.DCEStyle && !m.Config.FlagIsSet(gssapi.MutualAuthn) /* for non-dce style, there will be no APReply */ {

		// compute session keys.
		_, err := m.APReply(ctx, nil)
		if err != nil {
			return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
		}

		gssapi.SetAttribute(ctx, gssapi.AttributeSessionKey, m.ExportedSessionKey)
		gssapi.SetAttribute(ctx, gssapi.AttributeTarget, m.Config.SName)

		return &gssapi.Token{Payload: b}, gssapi.ContextComplete(ctx)
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

	ok, err := m.UnwrapInboundPayload(ctx, forSign, forSeal, tokEx.Signature)
	if err != nil {
		return nil, gssapi.ContextError(ctx, gssapi.Failure, err)
	}

	if !ok {
		return nil, gssapi.ContextError(ctx, gssapi.BadMIC, gssapi.ErrBadMIC)
	}

	return tokEx, nil

}

// Wrap token.
func (m *Mechanism) Wrap(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {

	tokEx, err := m.WrapEx(ctx, &gssapi.MessageTokenEx{
		QoP: tok.QoP,
		Payloads: []*gssapi.PayloadEx{
			{Capabilities: tok.Capabilities, Payload: tok.Payload},
		},
	})
	if err != nil {
		return nil, err
	}

	return &gssapi.MessageToken{
		QoP:          tokEx.QoP,
		Capabilities: tokEx.Payloads[0].Capabilities,
		Payload:      tokEx.Payloads[0].Payload,
		Signature:    tokEx.Signature,
	}, nil

}

// Unwrap token.
func (m *Mechanism) Unwrap(ctx context.Context, tok *gssapi.MessageToken) (*gssapi.MessageToken, error) {

	tokEx, err := m.UnwrapEx(ctx, &gssapi.MessageTokenEx{
		QoP: tok.QoP,
		Payloads: []*gssapi.PayloadEx{
			{Capabilities: tok.Capabilities, Payload: tok.Payload},
		},
	})
	if err != nil {
		return nil, err
	}

	return &gssapi.MessageToken{
		QoP:          tokEx.QoP,
		Capabilities: tokEx.Payloads[0].Capabilities,
		Payload:      tokEx.Payloads[0].Payload,
		Signature:    tokEx.Signature,
	}, nil
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

	if !bytes.Equal(tok.Signature, expSgn) {
		return gssapi.ContextError(ctx, gssapi.BadMIC, gssapi.ErrBadMIC)
	}

	return gssapi.ContextComplete(ctx)
}

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

	if !bytes.Equal(tokEx.Signature, expSgn) {
		return gssapi.ContextError(ctx, gssapi.BadMIC, gssapi.ErrBadMIC)
	}

	return gssapi.ContextComplete(ctx)
}

var (
	_ gssapi.Mechanism   = (*Mechanism)(nil)
	_ gssapi.MechanismEx = (*Mechanism)(nil)
)
