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
