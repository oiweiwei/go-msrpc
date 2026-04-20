package credential

import (
	"context"
	"strings"
)

// Database is an interface for storing and retrieving credentials.
type Database interface {
	// Add adds a credential to the database.
	Add(Credential)
	// Lookup looks up a credential in the database.
	Lookup(context.Context, Credential) (Credential, bool)
}

type localDatabase struct {
	credentials map[string]Credential
}

// NewLocalDatabase creates a new in-memory database for credentials.
func NewLocalDatabase() Database {
	return &localDatabase{
		credentials: make(map[string]Credential),
	}
}

// Add adds a credential to the local database.
func (db *localDatabase) Add(cred Credential) {
	if cred != nil {
		db.credentials[strings.ToLower(cred.UserName())] = cred
		if cred.DomainName() != "" {
			db.credentials[strings.ToLower(cred.DomainName()+"\\"+cred.UserName())] = cred
		}
	}
}

// Lookup looks up a credential in the local database.
func (db *localDatabase) Lookup(ctx context.Context, cred Credential) (Credential, bool) {
	if cred == nil {
		return nil, false
	}
	if cred.DomainName() != "" {
		name := strings.ToLower(cred.DomainName() + "\\" + cred.UserName())
		if cred, ok := db.credentials[name]; ok {
			return cred, true
		}
	}
	if cred, ok := db.credentials[strings.ToLower(cred.UserName())]; ok {
		return cred, true
	}
	return nil, false
}
