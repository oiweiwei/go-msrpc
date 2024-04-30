package credential

import "github.com/jcmturner/gokrb5/v8/keytab"

type Keytab interface {
	Credential
	// Keytab.
	Keytab() *keytab.Keytab
}

type kt struct {
	userName string
	realm    string
	kt       *keytab.Keytab
}

func (kt *kt) UserName() string {
	if kt != nil {
		return kt.userName
	}
	return ""
}

func (kt *kt) DomainName() string {
	if kt != nil {
		return kt.realm
	}
	return ""
}

func (kt *kt) Keytab() *keytab.Keytab {
	if kt != nil && kt.kt != nil {
		kkt := *kt.kt
		return &kkt
	}
	return nil
}
