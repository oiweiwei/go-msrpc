package credential

import "strings"

// Generic credential.
type Credential interface {
	// User name.
	UserName() string
	// Domain name.
	DomainName() string
	// Workstation.
	Workstation() string
}

func parseUser(un string, opts ...Option) userCred {
	dn, un, wkst := parseDomainUserWorkstation(un, opts...)
	return userCred{
		userName:    un,
		domainName:  dn,
		workstation: wkst,
	}
}

// parseDomainUserWorkstation parses the username and variadic workstation argument and outputs
// the domain name, user name, and workstation.
func parseDomainUserWorkstation(un string, opts ...Option) (string, string, string) {

	wkst, dn := "", ""

	for _, opt := range opts {
		switch v := opt.(type) {
		case wkstOpt:
			wkst = string(v)
		case domainOpt:
			dn = string(v)
		}
	}

	// down-level logon name.
	if strings.Contains(un, "\\") {
		un := strings.SplitN(un, "\\", 2)
		if dn != "" {
			return dn, un[1], wkst
		}
		return un[0], un[1], wkst
	}

	if strings.Contains(un, "@") {
		un := strings.SplitN(un, "@", 2)
		if dn != "" {
			return dn, un[0], wkst
		}
		return un[1], un[0], wkst
	}

	return dn, un, wkst
}

// DomainName function returns the domain name from the user name.
func DomainName(un string) string {
	dn, _, _ := parseDomainUserWorkstation(un)
	return dn
}

// Anonymous function returns the anonymous password credentials.
func Anonymous() Password {
	return &passwordCred{allowEmpty: true}
}

func NewFromString(s string, opts ...Option) Password {
	ss := strings.Split(s, "%")
	if len(ss) > 1 {
		return NewFromPassword(ss[0], ss[1], opts...)
	}
	return NewFromPassword(s, "", append(opts, AllowEmptyPassword())...)
}

func V8ToV9(cred Credential) Credential {

	if cred, ok := (any)(cred).(KeytabV8); ok {
		return NewFromKeytabV8(cred.UserName(), cred.Keytab(),
			Workstation(cred.Workstation()), Domain(cred.DomainName()))
	}

	if cred, ok := (any)(cred).(CCacheV8); ok {
		return NewFromCCacheV8(cred.UserName(), cred.CCache(),
			Workstation(cred.Workstation()), Domain(cred.DomainName()))
	}

	return cred
}
