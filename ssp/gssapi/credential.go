package gssapi

import (
	"context"
	"sync"
)

type CredentialUsage int

const (
	InitiateAndAccept CredentialUsage = 0
	InitiateOnly      CredentialUsage = 1
	AcceptOnly        CredentialUsage = 3
)

type Credential interface {
	// The target name for the credential.
	TargetName() string
	// The list of supported mechanisms.
	MechanismTypes() []OID
	// The credential usage.
	Usage() CredentialUsage
	// The actual credentials value (protocol specific).
	Value() any
}

// The credential represents the GSS API credential.
type credential struct {
	targetName     string
	mechanismTypes []OID
	usage          CredentialUsage
	value          any
}

// The credential target name.
func (c *credential) TargetName() string {
	return c.targetName
}

// The allowed mechanism types.
func (c *credential) MechanismTypes() []OID {
	mechanismTypes := make([]OID, len(c.mechanismTypes))
	copy(mechanismTypes, c.mechanismTypes)
	return mechanismTypes
}

// The credential usage.
func (c *credential) Usage() CredentialUsage {
	return c.usage
}

// The credential value.
func (c *credential) Value() any {
	return c.value
}

type CredentialStore struct {
	mu    sync.Mutex
	creds []Credential
}

func NewCredential(targetName string, mechanismTypes []OID, usage CredentialUsage, value any) *credential {
	return &credential{targetName: targetName,
		mechanismTypes: mechanismTypes,
		usage:          usage,
		value:          value,
	}
}

// AddCredential function adds the credential to the storage.
func (c *CredentialStore) AddCredential(ctx context.Context, value any) {

	c.mu.Lock()
	defer c.mu.Unlock()

	if cred, ok := value.(*credential); ok {
		c.creds = append(c.creds, cred)
		return
	}

	c.creds = append(c.creds, &credential{
		value: value,
	})
}

// GetCredential function retrieves the matching credential from the storage.
func (c *CredentialStore) GetCredential(ctx context.Context, name string, mechanismType OID, usage CredentialUsage) Credential {

	c.mu.Lock()
	defer c.mu.Unlock()

	for _, cred := range c.creds {

		// check target name.
		if cred.TargetName() != "" && name != "" && name != cred.TargetName() {
			continue
		}

		// check usage.
		if cred.Usage() != InitiateAndAccept && cred.Usage() != usage {
			continue
		}

		// check mechanism types.
		if mechs := cred.MechanismTypes(); len(mechs) != 0 && mechanismType != nil {

			found := false

			for _, mech := range mechs {
				if mech.Equal(mechanismType) {
					found = true
					break
				}
			}

			if !found {
				continue
			}
		}

		// found credential.
		return cred
	}

	return nil
}

var (
	defaultCredentialStore = new(CredentialStore)
)

func AddCredential(value any) {
	defaultCredentialStore.AddCredential(context.Background(), value)
}

func GetCredential(ctx context.Context, name string, mechanismType OID, usage CredentialUsage) Credential {
	cc := fromContext(ctx)
	if cc != nil && cc.CredentialStore != nil {
		// use local credential store if defined.
		return cc.CredentialStore.GetCredential(ctx, name, mechanismType, usage)
	}
	return defaultCredentialStore.GetCredential(ctx, name, mechanismType, usage)
}

func GetCredentialValue(ctx context.Context, name string, mechanismType OID, usage CredentialUsage) any {
	cred := GetCredential(ctx, name, mechanismType, usage)
	if cred != nil {
		return cred.Value()
	}
	return nil
}
