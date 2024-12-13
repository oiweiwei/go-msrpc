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
		return NewFromPassword(ss[0], ss[1])
	}
	return NewFromPassword(s, "")
}

// NewFromPassword function returns the username/password credential.
func NewFromPassword(un, passwd string, opts ...Option) Password {
	dn, un, wkst := parseDomainUserWorkstation(un, opts...)
	return &password{
		domainName:  dn,
		userName:    un,
		password:    passwd,
		workstation: wkst,
	}
}

// DomainName function returns the domain name from the user name.
func DomainName(un string) string {
	dn, _, _ := parseDomainUserWorkstation(un)
	return dn
}

// Anonymous function returns the anonymous password credentials.
func Anonymous() Password {
	return &password{}
}
