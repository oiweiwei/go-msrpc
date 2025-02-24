package credential

import (
	"errors"
)

// Password credential.
type Password interface {
	// Credential. (UserName / DomainName).
	Credential
	// Password.
	Password() string
}

// Password implementation.
type passwordCred struct {
	userCred
	password   string
	allowEmpty bool
}

// Password.
func (p *passwordCred) Password() string {
	if p != nil {
		return p.password
	}
	return ""
}

// IsEmpty returns true if the password is empty.
func (p *passwordCred) IsEmpty() bool {
	return p == nil || (p.password == "" && !p.allowEmpty)
}

// Validate the password credential.
func (p *passwordCred) Validate() error {
	if p != nil && p.password == "" && !p.allowEmpty {
		return errors.New("password is empty, use credentials.Anonymous() or AllowEmptyPassword() options to allow empty password")
	}

	return nil
}

// NewFromPassword function returns the username/password credential.
func NewFromPassword(un, passwd string, opts ...Option) Password {
	cred := &passwordCred{
		userCred: parseUser(un, opts...),
		password: passwd,
	}

	for _, opt := range opts {
		if _, ok := opt.(allowEmptyPassword); ok {
			cred.allowEmpty = true
		}
	}

	return cred
}
