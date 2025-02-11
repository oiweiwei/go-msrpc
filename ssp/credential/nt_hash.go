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

type ntHash struct {
	user
	ntHash []byte
}

// NT Hash.
func (n *ntHash) NTHash() []byte {
	if n != nil {
		return n.ntHash
	}
	return nil
}

// NewFromNTHash function returns the NT Hash credentials using the NT hash string
// (hex-encoded MD4 of the password).
func NewFromNTHash(un, hash string, opts ...Option) NTHash {
	hashB, _ := hex.DecodeString(hash)
	return NewFromNTHashBytes(un, hashB, opts...)
}

// NewFromNTHashBytes function returns the username/hash credential.
func NewFromNTHashBytes(un string, hash []byte, opts ...Option) NTHash {
	return &ntHash{
		user:   parseUser(un, opts...),
		ntHash: hash,
	}
}
