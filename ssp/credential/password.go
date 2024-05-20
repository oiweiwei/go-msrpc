package credential

import (
	"strings"
)

// Password credential.
type Password interface {
	// Credential. (UserName / DomainName).
	Credential
	// Password.
	Password() string
}

// Password implementation.
type password struct {
	userName    string
	domainName  string
	password    string
	workstation string
}

// User name.
func (p *password) UserName() string {
	if p != nil {
		return p.userName
	}
	return ""
}

// Domain name.
func (p *password) DomainName() string {
	if p != nil {
		return p.domainName
	}
	return ""
}

// Password.
func (p *password) Password() string {
	if p != nil {
		return p.password
	}
	return ""
}

func (p *password) Workstation() string {
	if p != nil {
		return p.workstation
	}
	return ""
}

func NewFromString(s string) Password {
	ss := strings.Split(s, "%")
	if len(ss) > 1 {
		return NewFromPassword(ss[0], ss[1], "")
	}
	return NewFromPassword(s, "", "")
}

// NewFromPassword function returns the username/password credential.
func NewFromPassword(un, passwd string, workstation ...string) Password {
	dn, un, wkst := parseDomainUserWorkstation(un, workstation...)
	return &password{
		domainName:  dn,
		userName:    un,
		password:    passwd,
		workstation: wkst,
	}
}
