package credential

import (
	"encoding/hex"
)

// EncryptionKey interface represents the encryption key credentials.
// The KeyType represents the Kerberos 5 IANA assigned number for the
// key.
type EncryptionKey interface {
	Credential
	// Encryption key value.
	KeyValue() []byte
	// Key type.
	KeyType() int
}

// encryptionKeyCred struct represents the encryption key
// credentials.
type encryptionKeyCred struct {
	userCred
	keyValue []byte
	keyType  int
	keyErr   error
}

// KeyValue function returns the encryption key value.
func (cred *encryptionKeyCred) KeyValue() []byte {
	if cred != nil {
		return cred.keyValue
	}
	return nil
}

// KeyType function returns the encryption key type.
func (cred *encryptionKeyCred) KeyType() int {
	if cred != nil {
		return cred.keyType
	}
	return 0
}

func (cred *encryptionKeyCred) Validate() error {
	if cred != nil && cred.keyErr != nil {
		return cred.keyErr
	}
	return nil
}

func (cred *encryptionKeyCred) IsEmpty() bool {
	return cred == nil || (len(cred.keyValue) == 0 && cred.keyType == 0)
}

// NewFromEncryptionKey function returns the encryption key credentials using the encryption key string.
func NewFromEncryptionKey(un string, keyType int, key string, opts ...Option) EncryptionKey {
	b, err := hex.DecodeString(key)
	if err != nil {
		return &encryptionKeyCred{keyErr: err}
	}
	return NewFromEncryptionKeyBytes(un, keyType, b, opts...)
}

// NewFromEncryptionKeyBytes function returns the username/key credential.
func NewFromEncryptionKeyBytes(un string, keyType int, key []byte, opts ...Option) EncryptionKey {
	return &encryptionKeyCred{
		userCred: parseUser(un, opts...),
		keyValue: key,
		keyType:  keyType,
	}
}
