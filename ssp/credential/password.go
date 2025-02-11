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
	user
	password string
}

// Password.
func (p *password) Password() string {
	if p != nil {
		return p.password
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
	return &password{
		user:     parseUser(un, opts...),
		password: passwd,
	}
}
