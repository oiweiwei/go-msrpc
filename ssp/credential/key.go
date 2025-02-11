package credential

import "encoding/hex"

type EncryptionKey interface {
	// Credential. (UserName / DomainName).
	Credential
	// Encryption key value.
	KeyValue() []byte
	// Key type.
	KeyType() int
}

type encryptionKey struct {
	user
	keyValue []byte
	keyType  int
}

func (e *encryptionKey) KeyValue() []byte {
	if e != nil {
		return e.keyValue
	}
	return nil
}

func (e *encryptionKey) KeyType() int {
	if e != nil {
		return e.keyType
	}
	return 0
}

// NewFromEncryptionKey function returns the encryption key credentials using the encryption key string.
func NewFromEncryptionKey(un string, keyType int, key string, opts ...Option) EncryptionKey {
	keyb, _ := hex.DecodeString(key)
	return NewFromEncryptionKeyBytes(un, keyType, keyb, opts...)
}

// NewFromEncryptionKeyBytes function returns the username/key credential.
func NewFromEncryptionKeyBytes(un string, keyType int, key []byte, opts ...Option) EncryptionKey {
	return &encryptionKey{
		user:     parseUser(un, opts...),
		keyValue: key,
		keyType:  keyType,
	}
}
