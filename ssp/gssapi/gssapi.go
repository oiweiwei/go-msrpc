package gssapi

import (
	"context"
	"encoding/asn1"
)

// Cap is a capability used both for request and response.
type Cap int

// IsSet function returns true if provided capability is set.
func (c Cap) IsSet(cc Cap) bool {
	return c&cc != 0
}

const (
	// The deleg_req_flag requests delegation of access rights.
	Delegation Cap = 1 << 0
	// The mutual_req_flag requests mutual authentication.
	MutualAuthn Cap = 1 << 1
	// The replay_det_req_flag requests that replay detection features
	// be applied to messages transferred on the established context.
	ReplayDetection Cap = 1 << 2
	// The sequence_req_flag requests that sequencing be enforced.
	Sequencing Cap = 1 << 3
	// The conf_req_flag provide informatory input to
	// the GSS-API implementation as to per-message
	// confidentiality services will be required on the context.
	Confidentiality Cap = 1 << 4
	// The integ_req_flag provide informatory input to the GSS-API
	// implementation as to per-message integrity services will be
	// required on the context.
	Integrity Cap = 1 << 5
	// The anon_req_flag requests that the initiator's identity
	// not be transferred within tokens to be sent to the acceptor.
	Anonymity Cap = 1 << 6
	// This flag allows the client to indicate to the server
	// that datagram service should be used.
	Datagram Cap = 1 << 8
	// This flag was added for use with Microsoft's implementation of
	// Distributed Computing Environment Remote Procedure Call (DCE RPC),
	// which initially expected three legs of authentication.
	// Setting this flag causes an extra AP reply to be sent from the
	// client back to the server after receiving the server's AP reply.
	DCEStyle Cap = 1 << 12
	// This flag allows the client to indicate to the
	// server that it should only allow the server application to identify
	// the client by name and ID, but not to impersonate the client.
	Identify Cap = 1 << 13
	// Setting this flag indicates that the client wants to be informed of
	// extended error information. In particular, Windows 2000 status codes
	// may be returned in the data field of a Kerberos error message.
	// This allows the client to understand a server failure more precisely.
	ExtendedError Cap = 1 << 14
)

// The object identifier.
type OID asn1.ObjectIdentifier

// Equal.
func (o OID) Equal(other OID) bool {
	return (asn1.ObjectIdentifier)(o).Equal((asn1.ObjectIdentifier)(other))
}

// String.
func (o OID) String() string {
	return (asn1.ObjectIdentifier)(o).String()
}

// CredentialsManager manages credentials.
type CredentialsManager interface {
	// Add the credentials to the context.
	AddCredentials(context.Context, Credential) error
	// Obtains the user's identity proof, often a secret cryptographic key.
	AcquireCredentials(context.Context) error
}

// The Security Context token.
type Token struct {
	// The token payload.
	Payload []byte
}

// ContextManager establishes and manages security contexts
// between peers.
type ContextManager interface {
	// Initialize outbound security context.
	InitSecurityContext(context.Context, *Token, ...Option) (*Token, error)
	// Accepts inbound security context.
	AcceptSecurityContext(context.Context, *Token, ...Option) (*Token, error)
	// Deletes the security context.
	DeleteSecurityContext(context.Context, ...Option) error
	// The maximum message size for the given limit.
	WrapSizeLimit(context.Context, int, ...Option) int
	// Sets the context attribute.
	SetAttribute(context.Context, string, any, ...Option)
	// Extracts the context attribute.
	GetAttribute(context.Context, string, ...Option) (any, error)
}

var (
	AttributeSessionKey = "session_key"
	AttributeTarget     = "target"
	AttributeRPCContext = "rpc_security_context"
)

// The GSSAPI call option.
type Config struct {
	// The security compatibility parameter. (NTLM).
	Compatibility int
	// The Quality-of-Protection.
	QoP int
	// The request flag.
	Capabilities Cap
	// The liftime of the context.
	ContextTTL int
	// The target name.
	TargetName string
	// The flag that indicates whether the target name
	// was retrieved from the untrusted source.
	TargetNameFromUntrustedSource bool
	// The context mechanism.
	MechanismType OID
	// The list of mechanism configs.
	MechanismConfigs []MechanismConfig
	// The flag that indicates whether it's a server
	// handle.
	IsServer bool
}

// MakeOption function is used to build the option structure.
func MakeOptions(opts ...Option) *Config {
	opt := &Config{}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

// The option.
type Option func(*Config)

// WithConfig function sets the configuration.
func WithConfig(cfg *Config) Option {
	return func(o *Config) {
		*o = *cfg
	}
}

// WithMechanismConfig function appends the mechanism-specific
// configuration.
func WithMechanismConfig(cfg MechanismConfig) Option {
	return func(o *Config) {
		if o.MechanismType == nil {
			// set the mechanism type for the config.
			o.MechanismType = cfg.Type()
		}
		o.MechanismConfigs = append(o.MechanismConfigs, cfg)
	}
}

// WithQoP returns the option for quality-of-protection.
func WithQoP(qop int) Option {
	return func(o *Config) {
		o.QoP = qop
	}
}

// WithRequest returns the option for the various request
// flags.
func WithRequest(req Cap) Option {
	return func(o *Config) {
		o.Capabilities |= req
	}
}

// WithCompatibility returns the option for the compatibility
// parameter.
func WithCompatibility(compat int) Option {
	return func(o *Config) {
		o.Compatibility = compat
	}
}

// Option indicates that target name SPN was retrieved from
// the untrusted source.
type TargetNameSource struct {
	Trusted bool
}

// WithTargetName returns the option of the target name.
func WithTargetName(name string, source ...TargetNameSource) Option {
	return func(o *Config) {
		o.TargetName = name
		if len(source) > 0 {
			o.TargetNameFromUntrustedSource = !source[0].Trusted
		}
	}
}

// WithMechanismType returns the option of the mechanism type.
func WithMechanismType(oid OID) Option {
	return func(o *Config) {
		o.MechanismType = oid
	}
}

type MessageToken struct {
	// The quality-of-protection.
	QoP int
	// The request/response capabilities of the token.
	Capabilities Cap
	// The input/output payload for signature calculation
	// and/or encryption (or verification and decryption).
	Payload []byte
	// The input/output signature for verification or
	// generation.
	Signature []byte
}

// Per-Message Security Service Availability.
type SecurityService interface {
	// MakeSignature function accepts the payload and returns the
	// signature for the payload.
	MakeSignature(context.Context, *MessageToken, ...Option) (*MessageToken, error)
	// VerifySignature function accepts the payload and signature
	// and returns nil if signature is valid.
	VerifySignature(context.Context, *MessageToken, ...Option) (*MessageToken, error)
	// Wrap function accepts the unencrypted payload and returns the
	// encrypted payload and signature.
	Wrap(context.Context, *MessageToken, ...Option) (*MessageToken, error)
	// Unwrap function accepts the encrypted payload and signature and
	// returns the unencrypted paylaod.
	Unwrap(context.Context, *MessageToken, ...Option) (*MessageToken, error)
}

// Security context attributes storage.
type SecurityAttributes interface {
	// GetAttribute function retrieves the security context attribute, ie session key,
	// target name.
	GetAttribute(context.Context, string, ...Option) (any, bool)
	// SetAttribute function sets the security context attribute.
	SetAttribute(context.Context, string, any, ...Option)
}

// MessageTokenEx represents the extended message token structure.
type MessageTokenEx struct {
	// The quality-of-protection.
	QoP int
	// The list of Payloads.
	Payloads []*PayloadEx
	// The signature.
	Signature []byte
}

// PayloadEx represents the list of payloads with capabilities.
type PayloadEx struct {
	// The request/response capabilities of the token. Must be only
	// Confidentiality and/or Integrity.
	Capabilities Cap
	// The payload.
	Payload []byte
}

// Microsoft per-Message Security Service Availability.
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nlmp/a06bfc2b-30fc-4483-b876-a9386f4808ed
type SecurityServiceEx interface {
	// MakeSignatureEx function accepts the list of payloads and returns the
	// signature for the payload.
	MakeSignatureEx(context.Context, *MessageTokenEx, ...Option) (*MessageTokenEx, error)
	// VerifySignatureEx function accepts the list of payloads and signature
	// and returns nil if signature is valid.
	VerifySignatureEx(context.Context, *MessageTokenEx, ...Option) (*MessageTokenEx, error)
	// WrapEx function accepts the list of unencrypted payloads and returns the
	// encrypted payload and signature.
	WrapEx(context.Context, *MessageTokenEx, ...Option) (*MessageTokenEx, error)
	// UnwrapEx function accepts the list of encrypted payloads and signature and
	// returns the unencrypted paylaod.
	UnwrapEx(context.Context, *MessageTokenEx, ...Option) (*MessageTokenEx, error)
}
