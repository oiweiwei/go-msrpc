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
	err      error
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

func (kt *kt) Workstation() string {
	return ""
}

func (kt *kt) Keytab() *keytab.Keytab {
	if kt != nil && kt.kt != nil {
		kkt := *kt.kt
		return &kkt
	}
	return nil
}

// NewFromKeytabFile ...
func NewFromKeytabFile(un string, keytabFile string, opts ...Option) Keytab {
	kt, _ := keytab.Load(keytabFile)
	return NewFromKeytab(un, kt)
}

// NewFromKeytab ...
func NewFromKeytab(un string, keytab *keytab.Keytab, opts ...Option) Keytab {
	realm, un, _ := parseDomainUserWorkstation(un, opts...)
	return &kt{
		userName: un,
		realm:    realm,
		kt:       keytab,
	}
}
