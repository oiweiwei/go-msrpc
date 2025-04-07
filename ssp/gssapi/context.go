package gssapi

import (
	"context"
	"fmt"
)

type ChannelBindings interface {
	Marshal() ([]byte, error)
}

// The security context represents the GSS-API context entitiy that
// contains the selected mechanism, credentials and settings for the
// security services.
type SecurityContext struct {
	// Status.
	Status Status
	// Error.
	Error error
	// The credential handle.
	Credential Credential
	// The security compatibility parameter. (NTLM).
	Compatibility int
	// The Quality-of-Protection.
	QoP int
	// Target name.
	TargetName string
	// The flag that indicates whether the target name
	// was retrieved from the untrusted source.
	TargetNameFromUntrustedSource bool
	// The selected mechanism.
	Mechanism Mechanism
	// The negotiated capabilities.
	Capabilities Cap
	// The lifetime of the security context.
	ContextTTL int
	// Channel binding.
	ChannelBindings ChannelBindings
	// Mechanism-specific configuration.
	MechanismConfigs []MechanismConfig
	// IsServer.
	IsServer bool
	// attributes.
	Attributes map[string]interface{}
	// Local context mechanism storage.
	MechanismStore *MechanismStore
	// Local context credentials storage.
	CredentialStore *CredentialStore
}

// The security context key.
type ctxKey struct{}

// fromContext.
func fromContext(ctx context.Context) *SecurityContext {
	cc, _ := ctx.Value(ctxKey{}).(*SecurityContext)
	return cc
}

// withContextStatus.
func withContextStatus(ctx context.Context, status Status, err error) error {

	cc := fromContext(ctx)
	if cc == nil {
		return ErrNoContext
	}

	if cc.Error != nil {
		return fmt.Errorf("security provider: %w", cc.Error)
	}

	cc.Status, cc.Error = status, err

	if cc.Error != nil {
		return fmt.Errorf("security provider: %w", err)
	}

	return nil
}

type ContextOption any

// WithCredential function returns the credential option for the
// local context credential store.
func WithCredential(value any) ContextOption {
	return &credential{value: value}
}

// WithMechanismFactory function returns the mechanism factory option for the
// local context mechanism store.
func WithMechanismFactory(value MechanismFactory, defaultConfig ...MechanismConfig) ContextOption {
	if len(defaultConfig) > 0 {
		return WithDefaultConfig(value, defaultConfig[0].Copy())
	}
	return value
}

// NewSecurityContext initializes the Security Context.
// The function must be called before InitSecurityContext.
func NewSecurityContext(ctx context.Context, opts ...ContextOption) context.Context {

	var (
		creds *CredentialStore
		mechs *MechanismStore
	)

	// fill-in the credentials and mechanisms local to the current
	// context.
	for _, o := range opts {
		switch o := o.(type) {
		case *credential:
			if creds == nil {
				creds = new(CredentialStore)
			}
			creds.AddCredential(ctx, o)
		case MechanismFactory:
			if mechs == nil {
				mechs = new(MechanismStore)
			}
			mechs.AddMechanism(o)
		}
	}

	return context.WithValue(ctx, ctxKey{}, &SecurityContext{MechanismStore: mechs, CredentialStore: creds})
}

// ResetSecurityContext to it's initial state.
func ResetSecurityContext(ctx context.Context) context.Context {
	cc := fromContext(ctx)
	if cc == nil {
		return ctx
	}
	return context.WithValue(ctx, ctxKey{}, &SecurityContext{MechanismStore: cc.MechanismStore, CredentialStore: cc.CredentialStore})
}

// FromContext retrieves the Security Context.
func FromContext(ctx context.Context) SecurityContext {
	cc := fromContext(ctx)
	if cc != nil {
		// prevent altering context by mechanism.
		c := *cc
		c.Mechanism = nil                              // guard from changing the state.
		c.MechanismConfigs = nil                       // guard from changing the state.
		c.Attributes = nil                             // guard from changing the state.
		c.MechanismStore, c.CredentialStore = nil, nil // guard from changing the state.
		return c
	}
	return SecurityContext{}
}

// GetMechanismConfig returns the default config for the mechanism.
func GetMechanismConfig(ctx context.Context, oid OID) any {

	cc := fromContext(ctx)
	if cc == nil {
		return nil
	}

	for _, c := range cc.MechanismConfigs {
		if c.Type().Equal(oid) {
			return c.Copy()
		}
	}

	for _, m := range ListMechanisms(ctx) {
		if m.Type().Equal(oid) {
			c, err := m.DefaultConfig(ctx)
			if err != nil {
				return nil
			}
			return c.Copy()
		}
	}

	return nil
}

// Clear the security context.
func DeleteSecurityContext(ctx context.Context, _ ...Option) error {

	cc := fromContext(ctx)
	if cc == nil {
		return ErrNoContext
	}

	*cc = SecurityContext{}
	return nil
}

// Initialize outbound security context.
func InitSecurityContext(ctx context.Context, tok *Token, opts ...Option) (*Token, error) {

	cc := fromContext(ctx)
	if cc == nil {
		return nil, ErrNoContext
	}

	var err error

	if cc.Status == NoContext {

		// initalize context with parameters.

		cfg := MakeOptions(opts...)

		cc.Compatibility = cfg.Compatibility
		cc.QoP = cfg.QoP
		cc.Capabilities = cfg.Capabilities
		cc.ContextTTL = cfg.ContextTTL
		cc.TargetName = cfg.TargetName
		cc.TargetNameFromUntrustedSource = cfg.TargetNameFromUntrustedSource
		cc.MechanismConfigs = cfg.MechanismConfigs

		f := GetMechanism(ctx, cfg.MechanismType)
		if f == nil {
			return nil, ContextError(ctx, BadMech, ErrBadMech)
		}

		// get stored credentials.
		cc.Credential = GetCredential(ctx, cfg.TargetName, f.Type(), InitiateOnly)

		// optionally validate credential.
		if cc.Credential != nil {
			if validator, ok := (any)(cc.Credential.Value()).(interface{ Validate() error }); ok {
				if err = validator.Validate(); err != nil {
					return nil, ContextError(ctx, Failure, err)
				}
			}
		}

		// initiator is a client.
		cc.IsServer = false

		if cc.Mechanism, err = f.New(ctx); err != nil {
			return nil, ContextError(ctx, Failure, err)
		}
	}

	// handle error.
	if tok, err = cc.Mechanism.Init(ctx, tok); err != nil {
		// set proper error if not set.
		if cc.Status == NoContext || cc.Error == nil {
			return nil, withContextStatus(ctx, Failure, err)
		}
		return nil, err
	}

	// set proper status. (if nil is returned.)
	if cc.Status == NoContext {
		cc.Status = Complete
	}

	if tok != nil {
		return tok, nil
	}

	return &Token{}, nil
}

// The maximum message size for the given limit.
func WrapSizeLimit(ctx context.Context, sz int, opts ...Option) int {

	cc := fromContext(ctx)
	if cc == nil {
		return sz
	}

	if cc.Status != Complete {
		return sz
	}

	cfg := MakeOptions(opts...)

	return cc.Mechanism.WrapSizeLimit(ctx, sz, cfg.Capabilities.IsSet(Confidentiality))
}

func Wrap(ctx context.Context, tok *MessageToken, opts ...Option) (*MessageToken, error) {

	cc := fromContext(ctx)
	if cc == nil {
		return nil, ErrNoContext
	}

	if cc.Status != Complete {
		return nil, ErrUnavailable
	}

	cfg := MakeOptions(opts...)

	if cfg.QoP >= 0 {
		tok.QoP = cfg.QoP
	}

	tok.Capabilities |= cfg.Capabilities

	return cc.Mechanism.Wrap(ctx, tok)
}

// WrapEx function accepts the list of unencrypted payloads and returns the
// encrypted payload and signature.
func WrapEx(ctx context.Context, tokEx *MessageTokenEx, opts ...Option) (*MessageTokenEx, error) {

	cc := fromContext(ctx)
	if cc == nil {
		return nil, ErrNoContext
	}

	if cc.Status != Complete {
		return nil, ErrUnavailable
	}

	cfg := MakeOptions(opts...)

	if cfg.QoP >= 0 {
		tokEx.QoP = cfg.QoP
	}

	for _, tok := range tokEx.Payloads {
		tok.Capabilities |= cfg.Capabilities
	}

	mechEx, ok := (interface{})(cc.Mechanism).(MechanismEx)
	if !ok {
		return nil, ContextError(ctx, Unavailable, ErrUnavailable)
	}

	return mechEx.WrapEx(ctx, tokEx)
}

// UnwrapEx function accepts the list of encrypted payloads and signature and
// returns the unencrypted paylaod.
func Unwrap(ctx context.Context, tok *MessageToken, opts ...Option) (*MessageToken, error) {

	cc := fromContext(ctx)
	if cc == nil {
		return nil, ErrNoContext
	}

	if cc.Status != Complete {
		return nil, ErrUnavailable
	}

	cfg := MakeOptions(opts...)

	if cfg.QoP >= 0 {
		tok.QoP = cfg.QoP
	}

	tok.Capabilities |= cfg.Capabilities

	return cc.Mechanism.Unwrap(ctx, tok)
}

// UnwrapEx function accepts the list of encrypted payloads and signature and
// returns the unencrypted paylaod.
func UnwrapEx(ctx context.Context, tokEx *MessageTokenEx, opts ...Option) (*MessageTokenEx, error) {

	cc := fromContext(ctx)
	if cc == nil {
		return nil, ErrNoContext
	}

	if cc.Status != Complete {
		return nil, ErrUnavailable
	}

	cfg := MakeOptions(opts...)

	if cfg.QoP >= 0 {
		tokEx.QoP = cfg.QoP
	}

	for _, tok := range tokEx.Payloads {
		tok.Capabilities |= cfg.Capabilities
	}

	mechEx, ok := (interface{})(cc.Mechanism).(MechanismEx)
	if !ok {
		return nil, ContextError(ctx, Unavailable, ErrUnavailable)
	}

	return mechEx.UnwrapEx(ctx, tokEx)
}

// MakeSignature function accepts the payload and returns the
// signature for the payload.
func MakeSignature(ctx context.Context, tok *MessageToken, opts ...Option) (*MessageToken, error) {

	cc := fromContext(ctx)
	if cc == nil {
		return nil, ErrNoContext
	}

	if cc.Status != Complete {
		return nil, ErrUnavailable
	}

	cfg := MakeOptions(opts...)

	if cfg.QoP >= 0 {
		tok.QoP = cfg.QoP
	}

	tok.Capabilities |= cfg.Capabilities

	return cc.Mechanism.MakeSignature(ctx, tok)
}

// MakeSignatureEx function accepts the list of payloads and returns the
// payload signature.
func MakeSignatureEx(ctx context.Context, tokEx *MessageTokenEx, opts ...Option) (*MessageTokenEx, error) {

	cc := fromContext(ctx)
	if cc == nil {
		return nil, ErrNoContext
	}

	if cc.Status != Complete {
		return nil, ErrUnavailable
	}

	cfg := MakeOptions(opts...)

	if cfg.QoP >= 0 {
		tokEx.QoP = cfg.QoP
	}

	for _, tok := range tokEx.Payloads {
		tok.Capabilities |= cfg.Capabilities
	}

	mechEx, ok := (interface{})(cc.Mechanism).(MechanismEx)
	if !ok {
		return nil, ContextError(ctx, Unavailable, ErrUnavailable)
	}

	return mechEx.MakeSignatureEx(ctx, tokEx)
}

// VerifySignature function accepts the payload and signature
// and returns nil if signature is valid.
func VerifySignature(ctx context.Context, tok *MessageToken, opts ...Option) error {

	cc := fromContext(ctx)
	if cc == nil {
		return ErrNoContext
	}

	if cc.Status != Complete {
		return ErrUnavailable
	}

	cfg := MakeOptions(opts...)

	if cfg.QoP >= 0 {
		tok.QoP = cfg.QoP
	}

	tok.Capabilities |= cfg.Capabilities

	return cc.Mechanism.VerifySignature(ctx, tok)
}

func VerifySignatureEx(ctx context.Context, tokEx *MessageTokenEx, opts ...Option) error {

	cc := fromContext(ctx)
	if cc == nil {
		return ErrNoContext
	}

	if cc.Status != Complete {
		return ErrUnavailable
	}

	cfg := MakeOptions(opts...)

	if cfg.QoP >= 0 {
		tokEx.QoP = cfg.QoP
	}

	for _, tok := range tokEx.Payloads {
		tok.Capabilities |= cfg.Capabilities
	}

	mechEx, ok := (interface{})(cc.Mechanism).(MechanismEx)
	if !ok {
		return ContextError(ctx, Unavailable, ErrUnavailable)
	}

	return mechEx.VerifySignatureEx(ctx, tokEx)
}

// GetAttribute function retrieves the attribute from the security context.
func GetAttribute(ctx context.Context, attrName string, _ ...Option) (any, bool) {

	cc := fromContext(ctx)

	if cc == nil || cc.Attributes == nil {
		return nil, false
	}

	attr, ok := cc.Attributes[attrName]
	return attr, ok
}

// SetAttribute function sets the attribute to the current security context.
func SetAttribute(ctx context.Context, attrName string, attrValue any, _ ...Option) {

	cc := fromContext(ctx)
	if cc == nil {
		return
	}

	if cc.Attributes == nil {
		cc.Attributes = make(map[string]any)
	}

	cc.Attributes[attrName] = attrValue
	return
}

// ContextComplete function informs on successful operation complete
// or context establishment.
func ContextComplete(ctx context.Context) error {
	return withContextStatus(ctx, Complete, nil)
}

func IsComplete(ctx context.Context) bool {
	return FromContext(ctx).Status == Complete
}

// ContextContinueNeeded function sets the context status to CONTINUE_NEEDED.
func ContextContinueNeeded(ctx context.Context) error {
	return withContextStatus(ctx, ContinueNeeded, nil)
}

// ContextError function sets the context error.
func ContextError(ctx context.Context, status Status, err error) error {
	return withContextStatus(ctx, status, err)
}
