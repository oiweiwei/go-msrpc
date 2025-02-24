package credential

import (
	"encoding/hex"
)

type NTHash interface {
	// Credential. (UserName / DomainName).
	Credential
	// NT Hash.
	NTHash() []byte
}

type ntHashCred struct {
	userCred
	hash    []byte
	hashErr error
}

// NT Hash.
func (cred *ntHashCred) NTHash() []byte {
	if cred != nil {
		return cred.hash
	}
	return nil
}

func (cred *ntHashCred) Validate() error {
	if cred != nil && cred.hashErr != nil {
		return cred.hashErr
	}
	return nil
}

// IsEmpty function returns true if the NT Hash credential is empty.
func (cred *ntHashCred) IsEmpty() bool {
	return cred == nil || len(cred.hash) == 0
}

// NewFromNTHash function returns the NT Hash credentials using the NT hash string
// (hex-encoded MD4 of the password).
func NewFromNTHash(un, hash string, opts ...Option) NTHash {
	b, err := hex.DecodeString(hash)
	if err != nil {
		return &ntHashCred{hashErr: err}
	}
	return NewFromNTHashBytes(un, b, opts...)
}

// NewFromNTHashBytes function returns the username/hash credential.
func NewFromNTHashBytes(un string, hash []byte, opts ...Option) NTHash {
	return &ntHashCred{
		userCred: parseUser(un, opts...),
		hash:     hash,
	}
}
