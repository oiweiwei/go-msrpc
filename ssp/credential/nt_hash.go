package credential

import (
	"encoding/hex"
)

type NTHash interface {
	// Credential. (UserName / DomainName).
	Credential
	// NT Hash.
	NTHash() []byte
	// KVNO. (optional for KRB5).
	KVNO() int
}

type ntHash struct {
	userName    string
	domainName  string
	ntHash      []byte
	kvno        int
	workstation string
}

// User name.
func (n *ntHash) UserName() string {
	if n != nil {
		return n.userName
	}
	return ""
}

// Domain name.
func (n *ntHash) DomainName() string {
	if n != nil {
		return n.domainName
	}
	return ""
}

// NT Hash.
func (n *ntHash) NTHash() []byte {
	if n != nil {
		return n.ntHash
	}
	return nil
}

// Workstation.
func (p *ntHash) Workstation() string {
	if p != nil {
		return p.workstation
	}
	return ""
}

// KVNO.
func (p *ntHash) KVNO() int {
	if p != nil {
		return p.kvno
	}
	return 0
}

// NewFromNTHash function returns the NT Hash credentials using the NT hash string
// (hex-encoded MD4 of the password).
func NewFromNTHash(un, hash string, opts ...Option) NTHash {
	hashB, _ := hex.DecodeString(hash)
	return NewFromNTHashBytes(un, hashB, opts...)
}

// NewFromNTHashBytes function returns the username/hash credential.
func NewFromNTHashBytes(un string, hash []byte, opts ...Option) NTHash {

	dn, un, wkst := parseDomainUserWorkstation(un, opts...)
	kvno := 1 // default is 1.

	for _, opt := range opts {
		switch v := opt.(type) {
		case kvnoOpt:
			kvno = int(v)
		}
	}
	return &ntHash{
		domainName:  dn,
		userName:    un,
		ntHash:      hash,
		workstation: wkst,
		kvno:        kvno,
	}
}
