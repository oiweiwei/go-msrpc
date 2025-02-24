package ntlm

import (
	"context"
	"errors"
	"time"

	"github.com/oiweiwei/go-msrpc/ssp/credential"
)

const (
	NTLMv1 = 1
	NTLMv2 = 2
)

// The authentication level.
type AuthLevel int

const (
	// Client devices use LM and NTLM authentication, and they
	// never use NTLMv2 session security.
	LMv1AndNTLMv1 AuthLevel = iota
	// Client devices use LM and NTLM authentication, and they use
	// NTLMv2 session security if the server supports it.
	LMv1AndNTLMv1WithESS
	// Client devices use NTLMv1 authentication, and they use NTLMv2
	// session security if the server supports it.
	NTLMv1WithESS
	// Client devices use NTLMv2 authentication, and they use NTLMv2
	// session security if the server supports it.
	NTLMv2WithESS
)

// NewConfigFromAuthLevel function returns new configuration for the
// NTLM based on the authentication levels.
func NewConfigFromAuthLevel(lvl AuthLevel) *Config {

	switch lvl {
	case LMv1AndNTLMv1:
		return &Config{
			NTLMVersion: NTLMv1,
			NoESS:       true,
		}
	case LMv1AndNTLMv1WithESS:
		return &Config{
			NTLMVersion:       NTLMv1,
			ClientConfigFlags: NegotiateExtendedSessionSecurity,
		}
	case NTLMv1WithESS:
		return &Config{
			NTLMVersion:       NTLMv1,
			NoLMResponse:      true,
			ClientConfigFlags: NegotiateExtendedSessionSecurity,
		}
	default:
		return &Config{
			NTLMVersion:       NTLMv2,
			ClientConfigFlags: NegotiateExtendedSessionSecurity,
		}
	}
}

// NewConfig function returns the configuration for the NTLM SSP.
func NewConfig() *Config {
	return &Config{
		// By default, reqest MIC.
		RequestMIC: true,
		// By default ask for strongest encryption.
		Require128BitEncryption: true,
		// By default, don't compute LM response for
		// NTLMv1.
		NoLMResponse: true,
		// By default, extended session security must
		// be enabled.
		NoESS: false,
		// By default, use NTLMv2.
		NTLMVersion: 2,
	}
}

// The NTLM configuration.
type Config struct {
	// The services available for the NTLM.
	Integrity, Confidentiality, Identify, Anonymity, Datagram bool
	// The flag that indicates whether the MIC must be computed.
	RequestMIC bool
	// The set of client configuration flags that specify the
	// full set of capabilities of the client.
	ClientConfigFlags Flag
	// The client credential.
	Credential Credential
	// The version.
	Version Version
	// A flag that should control using the NTLM response
	// for the LM response to the server challenge when NTLMv1
	// authentication is used.
	NoLMResponse bool
	// Explicitly disable extended session security (for compatibility
	// level 0). Cannot be used with NTLMv2.
	NoESS bool
	// A flag that requires the client to use 128-bit encryption.
	Require128BitEncryption bool
	// The maximum lifetime for challenge/response pairs.
	MaxLifetime time.Duration
	// Service principal name (SPN) of the service to which the client
	// wishes to authenticate.
	TargetName string
	// An octet string provided by the application used for channel binding.
	ChannelBindings []byte
	// A Boolean setting that indicates that the caller generated the target's
	// SPN from an untrusted source.
	UnverifiedTargetName bool
	// The NTLM version.
	NTLMVersion int
	// IsServer.
	IsServer bool
	// The flag that indicates whether all the input buffers must be used to
	// build a signature. (DO NOT USE IT).
	NoSignAllBuffers bool
}

func IsCredentialEmpty(cred any) bool {

	if cred, ok := cred.(credential.Password); ok {
		return cred.Password() == ""
	}

	if cred, ok := cred.(credential.NTHash); ok {
		return len(cred.NTHash()) == 0
	}

	if cred, ok := cred.(credential.EncryptionKey); ok {
		return cred.KeyType() == 0 && len(cred.KeyValue()) == 0
	}

	return true
}

func IsValidCredential(cred any) bool {
	switch cred := cred.(type) {
	case credential.Password:
		return true
	case credential.NTHash:
		return true
	case credential.EncryptionKey:
		return cred.KeyType() == 23
	case credential.CCache:
		for _, key := range cred.CCache().GetEntries() {
			if key.Key.KeyType == 23 {
				return true
			}
		}
	}
	return false
}

func GetCredentialNTHash(cred any) ([]byte, bool) {
	var k []byte
	switch cred := cred.(type) {
	case credential.NTHash:
		k = cred.NTHash()
	case credential.EncryptionKey:
		k = cred.KeyValue()
	case credential.CCache:
		for _, key := range cred.CCache().GetEntries() {
			if key.Key.KeyType == 23 {
				k = key.Key.KeyValue
				break
			}
		}
	}
	return k, len(k) > 0
}

func (c *Config) NewNTLMVersion(ctx context.Context, cfg *Config, sess *SecurityParameters) NTLMVersion {
	switch c.NTLMVersion {
	case 1:
		return &V1{cfg, sess}
	}
	return &V2{cfg, sess}
}

func (c *Config) clone() *Config {
	newC := *c
	return &newC
}

var (
	// ErrIntegrityNotAvailable
	ErrIntegrityNotAvailable = errors.New("integrity, replay_detection, or sequencing were requested but not available")
	// ErrConfidentialityNotAvailable
	ErrConfidentialityNotAvailable = errors.New("confidentiality was requested but not available")
	// Err128BitEncryptionNotAvailable
	Err128BitEncryptionNotSupported = errors.New("128-bit encryption was requested but not supported")
)

// Verify function verifies the negotiated flags.
func (c *Config) Verify(f Flag) error {

	if c.Integrity && !f.IsSet(NegotiateSign) {
		return ErrIntegrityNotAvailable
	}

	if c.Confidentiality && !f.IsSet(NegotiateSeal) {
		return ErrConfidentialityNotAvailable
	}

	if c.Require128BitEncryption && !f.IsSet(Negotiate128) {
		return Err128BitEncryptionNotSupported
	}

	return nil
}

// Negotiate function returns the set of negotiated flags.
func (c *Config) Negotiate() Flag {

	f := c.ClientConfigFlags |
		RequestTarget |
		NegotiateNTLM |
		NegotiateAlwaysSign |
		NegotiateUnicode |
		Negotiate56

	if !c.NoESS {
		f = f.Set(NegotiateExtendedSessionSecurity | Negotiate128)
	}

	// always sign.
	c.Integrity = true

	if c.Integrity {
		f = f.Set(NegotiateSign | NegotiateKeyExchange)
	}

	if c.Confidentiality {
		f = f.Set(NegotiateSeal | NegotiateKeyExchange | NegotiateLMKey)
	}

	if c.Anonymity {
		f = f.Set(Anonymous)
	}

	if c.Identify {
		f = f.Set(NegotiateIdentify)
	}

	if c.Datagram {
		f = f.Set(NegotiateDatagram | NegotiateKeyExchange)
	}

	if !c.Version.IsZero() {
		f = f.Set(NegotiateVersion)
	}

	return f
}
