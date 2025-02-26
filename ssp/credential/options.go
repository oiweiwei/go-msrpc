package credential

type Option interface {
	is_CredentialOption()
}

type wkstOpt string

func (wkstOpt) is_CredentialOption() {}

// Workstation option.
func Workstation(s string) Option {
	return wkstOpt(s)
}

type allowEmptyPassword struct{}

func (allowEmptyPassword) is_CredentialOption() {}

// Allow empty password option.
func AllowEmptyPassword() Option {
	return allowEmptyPassword{}
}

type domainOpt string

func (domainOpt) is_CredentialOption() {}

func Domain(s string) Option {
	return domainOpt(s)
}
