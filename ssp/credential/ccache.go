package credential

import (
	"github.com/oiweiwei/gokrb5.fork/v9/credentials"
	"github.com/oiweiwei/gokrb5.fork/v9/types"

	v8_credentials "github.com/jcmturner/gokrb5/v8/credentials"
)

type CCache interface {
	Credential
	// CCache.
	CCache() *credentials.CCache
}

type ccacheCred struct {
	userCred
	ccache *credentials.CCache
}

func (cred *ccacheCred) CCache() *credentials.CCache {
	if cred != nil {
		return cred.ccache
	}
	return nil
}

func (cred *ccacheCred) IsEmpty() bool {
	return cred == nil || cred.ccache == nil || len(cred.ccache.Credentials) == 0
}

func (cred *ccacheCred) Validate() error {
	return nil
}

func NewFromCCache(un string, ccache *credentials.CCache, opts ...Option) CCache {
	return &ccacheCred{
		userCred: parseUser(un, opts...),
		ccache:   ccache,
	}
}

// NewFromCCacheV8 function creates a new CCache from jcmturner/gokrb5/v8/credentials.CCache.
func NewFromCCacheV8(un string, ccV8 *v8_credentials.CCache, opts ...Option) CCache {

	var cc = &credentials.CCache{
		Version: ccV8.Version,
		DefaultPrincipal: credentials.Principal{
			Realm:         ccV8.DefaultPrincipal.Realm,
			PrincipalName: types.PrincipalName(ccV8.DefaultPrincipal.PrincipalName),
		},
		Credentials: make([]*credentials.Credential, len(ccV8.Credentials)),
		Path:        ccV8.Path,
	}

	for i := range ccV8.Credentials {

		var addresses []types.HostAddress

		if ccV8.Credentials[i].Addresses != nil {
			addresses = make([]types.HostAddress, len(ccV8.Credentials[i].Addresses))
			for j := range ccV8.Credentials[i].Addresses {
				addresses[j] = types.HostAddress(ccV8.Credentials[i].Addresses[j])
			}
		}

		var authData []types.AuthorizationDataEntry

		if ccV8.Credentials[i].AuthData != nil {
			authData = make([]types.AuthorizationDataEntry, len(ccV8.Credentials[i].AuthData))
			for j := range ccV8.Credentials[i].AuthData {
				authData[j] = (types.AuthorizationDataEntry)(ccV8.Credentials[i].AuthData[j])
			}
		}

		cc.Credentials[i] = &credentials.Credential{
			Client: credentials.Principal{
				Realm:         ccV8.Credentials[i].Client.Realm,
				PrincipalName: types.PrincipalName(ccV8.Credentials[i].Client.PrincipalName),
			},
			Server: credentials.Principal{
				Realm:         ccV8.Credentials[i].Server.Realm,
				PrincipalName: types.PrincipalName(ccV8.Credentials[i].Server.PrincipalName),
			},
			Key:          (types.EncryptionKey)(ccV8.Credentials[i].Key),
			AuthTime:     ccV8.Credentials[i].AuthTime,
			StartTime:    ccV8.Credentials[i].StartTime,
			EndTime:      ccV8.Credentials[i].EndTime,
			RenewTill:    ccV8.Credentials[i].RenewTill,
			IsSKey:       ccV8.Credentials[i].IsSKey,
			TicketFlags:  ccV8.Credentials[i].TicketFlags,
			Addresses:    addresses,
			AuthData:     authData,
			Ticket:       ccV8.Credentials[i].Ticket,
			SecondTicket: ccV8.Credentials[i].SecondTicket,
		}
	}

	return NewFromCCache(un, cc, opts...)
}
