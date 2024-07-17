package netlogon

import "github.com/oiweiwei/go-msrpc/ssp/credential"

// The generic credential.
type Credential = credential.Credential

type Config struct {
	ServerName      string
	Capabilities    Cap
	Credential      Credential
	ClientChallenge []byte
	ServerChallenge []byte
	IsServer        bool
}

func NewConfig() *Config {
	return &Config{Capabilities: CapAES_SHA2 | CapStrongKey | CapSecureRPC | CapRC4}
}
