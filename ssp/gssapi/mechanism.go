package gssapi

import (
	"context"
	"errors"
)

type MechanismConfig interface {
	// The mechanism type object identifier.
	Type() OID
	// Copy must return copy of the configuration.
	Copy() MechanismConfig
}

type MechanismFactory interface {
	// New returns the mechanism instance.
	New(context.Context) (Mechanism, error)
	// DefaultConfig returns the default config for the instance.
	DefaultConfig(context.Context) (MechanismConfig, error)
	// The mechanism type object identifier.
	Type() OID
}

// MechanismFactoryWithConfig represents the mechanism factory with default
// configuration attached.
type MechanismFactoryWithConfig struct {
	MechanismFactory
	defaultConfig MechanismConfig
}

var ErrInvalidConfig = errors.New("mechanism factory with config: config type mismatch")

// DefaultConfig function returns the default configuration associated with mechanism factory.
func (f MechanismFactoryWithConfig) DefaultConfig(ctx context.Context) (MechanismConfig, error) {

	if f.defaultConfig == nil {
		// use default.
		return f.MechanismFactory.DefaultConfig(ctx)
	}

	if !f.defaultConfig.Type().Equal(f.Type()) {
		// check type.
		return nil, ErrInvalidConfig
	}

	return f.defaultConfig.Copy(), nil
}

// WithDefaultConfig function returns the mechanism factory with default configuration attached.
func WithDefaultConfig(factory MechanismFactory, config MechanismConfig) MechanismFactory {
	return MechanismFactoryWithConfig{
		MechanismFactory: factory,
		defaultConfig:    config,
	}
}

type Mechanism interface {
	// The mechanism type object identifier.
	Type() OID
	// The security context init call.
	Init(ctx context.Context, token *Token) (*Token, error)
	// The security context accept call.
	Accept(ctx context.Context, token *Token) (*Token, error)
	// The maximum message size for the given limit. (and flag determining if
	// conf is required).
	WrapSizeLimit(context.Context, int, bool) int
	// Wrap token.
	Wrap(context.Context, *MessageToken) (*MessageToken, error)
	// Unwrap token.
	Unwrap(context.Context, *MessageToken) (*MessageToken, error)
	// MakeSignature token.
	MakeSignature(context.Context, *MessageToken) (*MessageToken, error)
	// VerifySignature token.
	VerifySignature(context.Context, *MessageToken) error
}

type MechanismEx interface {
	// Wrap token.
	WrapEx(context.Context, *MessageTokenEx) (*MessageTokenEx, error)
	// Unwrap token.
	UnwrapEx(context.Context, *MessageTokenEx) (*MessageTokenEx, error)
	// MakeSignature token.
	MakeSignatureEx(context.Context, *MessageTokenEx) (*MessageTokenEx, error)
	// VerifySignature token.
	VerifySignatureEx(context.Context, *MessageTokenEx) error
}
