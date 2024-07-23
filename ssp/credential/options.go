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

type kvnoOpt int

func (kvnoOpt) is_CredentialOption() {}

// KVNO option.
func KVNO(i int) Option {
	return kvnoOpt(i)
}
