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
	// Workstation.
	Workstation() string
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

	wkst := ""
	if len(workstation) != 0 {
		wkst = workstation[0]
	}

	// down-level logon name.
	if strings.Contains(un, "\\") {
		un := strings.SplitN(un, "\\", 2)
		return &password{
			// domainName:  strings.ToUpper(un[0]),
			domainName:  un[0],
			userName:    un[1],
			password:    passwd,
			workstation: wkst,
		}
	}

	// UPN.
	if strings.Contains(un, "@") {
		un := strings.SplitN(un, "@", 2)
		return &password{
			domainName:  un[1], //strings.ToLower(un[1]),
			userName:    un[0],
			password:    passwd,
			workstation: wkst,
		}
	}

	return &password{
		userName:    un,
		password:    passwd,
		workstation: wkst,
	}
}
