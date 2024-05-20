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
	userName    string
	domainName  string
	ntHash      []byte
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

// Password.
func (n *ntHash) NTHash() []byte {
	if n != nil {
		return n.ntHash
	}
	return nil
}

func (p *ntHash) Workstation() string {
	if p != nil {
		return p.workstation
	}
	return ""
}

// NewFromNTHash function returns the NT Hash credentials using the NT hash string
// (hex-encoded MD4 of the password).
func NewFromNTHash(un, hash string, workstation ...string) NTHash {
	hashB, _ := hex.DecodeString(hash)
	return NewFromNTHashBytes(un, hashB, workstation...)
}

// NewFromNTHashBytes function returns the username/hash credential.
func NewFromNTHashBytes(un string, hash []byte, workstation ...string) NTHash {
	dn, un, wkst := parseDomainUserWorkstation(un, workstation...)
	return &ntHash{
		domainName:  dn,
		userName:    un,
		ntHash:      hash,
		workstation: wkst,
	}
}
