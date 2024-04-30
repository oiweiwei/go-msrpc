package credential

// Generic credential.
type Credential interface {
	// User name.
	UserName() string
	// Domain name.
	DomainName() string
}
