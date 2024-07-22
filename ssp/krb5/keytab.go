package krb5

import (
	"time"

	"github.com/jcmturner/gokrb5/v8/credentials"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/types"
)

// Entry represents a keytab entry.
type Entry struct {
	// Principal is the principal name.
	Principal Principal
	// Timestamp is the timestamp of the key.
	Timestamp time.Time
	// KVNO8 is the key version number.
	KVNO8 uint8
	// Key is the encryption key.
	Key types.EncryptionKey
	// KVNO is the key version number.
	KVNO uint32
}

// Principal represents a keytab principal.
type Principal struct {
	// NumComponents.
	NumComponents int16 `json:"-"`
	// Realm.
	Realm string
	// Components.
	Components []string
	// NameType.
	NameType int32
}

// FIXME: this is a workaround to add an entry to the keytab.
func AddEntry(kt *keytab.Keytab, entry *Entry) error {

	raw := []byte{
		// header.
		0x05,                   // first-byte
		0x02,                   // version
		0x00, 0x00, 0x00, 0x11, // entry-length
		// principal.
		0x00, 0x00, // num components.
		0x00, 0x00, // realm length.
		0x00, 0x00, 0x00, 0x00, // name type.
		// key.
		0x00, 0x00, 0x00, 0x00, // timestamp.
		0x00,       // kvno8.
		0x00, 0x00, // key type.
		0x00, 0x00, // key lenght.
	}

	tmp := &keytab.Keytab{}
	tmp.Unmarshal(raw)

	e := tmp.Entries[0]

	e.Principal.NumComponents = entry.Principal.NumComponents
	e.Principal.Components = entry.Principal.Components
	e.Principal.Realm = entry.Principal.Realm
	e.Principal.NameType = entry.Principal.NameType

	e.Timestamp = entry.Timestamp
	e.KVNO8 = entry.KVNO8
	e.Key.KeyType = entry.Key.KeyType
	e.Key.KeyValue = entry.Key.KeyValue
	e.KVNO = entry.KVNO

	kt.Entries = append(kt.Entries, e)
	return nil
}

func WithNTHash(creds *credentials.Credentials, ntHash []byte) *credentials.Credentials {

	kt := new(keytab.Keytab)

	AddEntry(kt, &Entry{
		Principal: Principal{
			Realm:         creds.Realm(),
			Components:    creds.CName().NameString,
			NameType:      creds.CName().NameType,
			NumComponents: int16(len(creds.CName().NameString)),
		},
		Timestamp: time.Now(),
		KVNO8:     0,
		Key: types.EncryptionKey{
			KeyType:  etypeID.RC4_HMAC,
			KeyValue: ntHash,
		},
		KVNO: 1, // TODO: seems to be a good choice.
	})

	return creds.WithKeytab(kt)
}
