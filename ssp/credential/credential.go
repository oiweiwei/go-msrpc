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

func parseUser(un string, opts ...Option) user {
	dn, un, wkst := parseDomainUserWorkstation(un, opts...)
	return user{un, dn, wkst}
}

// parseDomainUserWorkstation parses the username and variadic workstation argument and outputs
// the domain name, user name, and workstation.
func parseDomainUserWorkstation(un string, opts ...Option) (string, string, string) {

	wkst := ""

	for _, opt := range opts {
		switch v := opt.(type) {
		case wkstOpt:
			wkst = string(v)
		}
	}

	// down-level logon name.
	if strings.Contains(un, "\\") {
		un := strings.SplitN(un, "\\", 2)
		return un[0], un[1], wkst
	}

	if strings.Contains(un, "@") {
		un := strings.SplitN(un, "@", 2)
		return un[1], un[0], wkst
	}

	return "", un, wkst
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
