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
