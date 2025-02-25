package credential

import (
	"fmt"

	"github.com/oiweiwei/gokrb5.fork/v9/keytab"
	"github.com/oiweiwei/gokrb5.fork/v9/types"

	v8_keytab "github.com/jcmturner/gokrb5/v8/keytab"
)

// Keytab interface defines the Kerberos 5 Keytab credential.
type Keytab interface {
	Credential
	// Keytab.
	Keytab() *keytab.Keytab
}

// keytabCred is the implementation of the Keytab interface.
type keytabCred struct {
	userCred
	// keytab.
	keytab *keytab.Keytab
	// keytab error.
	keytabErr error
}

// Keytab function returns the Keytab.
func (cred *keytabCred) Keytab() *keytab.Keytab {
	if cred != nil {
		return cred.keytab
	}
	return &keytab.Keytab{}
}

// Validate function validates the Keytab credential.
func (cred *keytabCred) Validate() error {
	if cred != nil && cred.keytabErr != nil {
		return cred.keytabErr
	}
	return nil
}

// IsEmpty function returns true if the Keytab credential is empty.
func (cred *keytabCred) IsEmpty() bool {
	return cred == nil || cred.keytab == nil || len(cred.keytab.Entries) == 0
}

// NewFromKeytabFile function creates a new Keytab credential from a keytab file.
func NewFromKeytabFile(un string, keytabFile string, opts ...Option) Keytab {
	keytab, err := keytab.Load(keytabFile)
	if err != nil {
		return &keytabCred{keytabErr: err}
	}
	return NewFromKeytabV9(un, keytab)
}

func NewFromKeytab(un string, ktab any, opts ...Option) Keytab {
	switch ktab := ktab.(type) {
	case *keytab.Keytab:
		return NewFromKeytabV9(un, ktab, opts...)
	case *v8_keytab.Keytab:
		return NewFromKeytabV8(un, ktab, opts...)
	}
	return &keytabCred{
		keytabErr: fmt.Errorf("invalid type %T for keytab", ktab),
	}
}

// NewFromKeytab function creates a new Keytab credential from a keytab.
func NewFromKeytabV9(un string, keytab *keytab.Keytab, opts ...Option) Keytab {
	return &keytabCred{
		userCred: parseUser(un, opts...),
		keytab:   keytab,
	}
}

// NewFromKeytabV8 function creates a new Keytab credential from a keytab from
// github.com/jcmtruner/gokrb5/v8 library.
func NewFromKeytabV8(un string, ktv8 *v8_keytab.Keytab, opts ...Option) Keytab {
	kt := &keytab.Keytab{}
	for i := range ktv8.Entries {
		kt.Entries = append(kt.Entries, keytab.Entry{
			Principal: (keytab.Principal)(ktv8.Entries[i].Principal),
			Timestamp: ktv8.Entries[i].Timestamp,
			KVNO8:     ktv8.Entries[i].KVNO8,
			Key:       (types.EncryptionKey)(ktv8.Entries[i].Key),
			KVNO:      ktv8.Entries[i].KVNO,
		})
	}
	return NewFromKeytab(un, kt, opts...)
}
