package credential

import "github.com/oiweiwei/gokrb5.fork/v9/keytab"

type Keytab interface {
	Credential
	// Keytab.
	Keytab() *keytab.Keytab
	// Keytab error.
	Error() error
}

type kt struct {
	user
	kt  *keytab.Keytab
	err error
}

func (kt *kt) Keytab() *keytab.Keytab {
	if kt != nil && kt.kt != nil {
		kkt := *kt.kt
		return &kkt
	}
	return nil
}

func (kt *kt) Error() error {
	if kt != nil {
		return kt.err
	}
	return nil
}

// NewFromKeytabFile ...
func NewFromKeytabFile(un string, keytabFile string, opts ...Option) Keytab {
	tab, err := keytab.Load(keytabFile)
	kkt := NewFromKeytab(un, tab)
	kkt.(*kt).err = err
	return kkt
}

// NewFromKeytab ...
func NewFromKeytab(un string, keytab *keytab.Keytab, opts ...Option) Keytab {
	return &kt{
		user: parseUser(un, opts...),
		kt:   keytab,
	}
}
