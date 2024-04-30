package gssapi

import (
	"context"
)

type MechanismConfig interface {
	// The mechanism type object identifier.
	Type() OID
}

type MechanismFactory interface {
	// New returns the mechanism instance.
	New(context.Context) (Mechanism, error)
	// DefaultConfig returns the default config for the instance.
	DefaultConfig(context.Context) (MechanismConfig, error)
	// The mechanism type object identifier.
	Type() OID
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
