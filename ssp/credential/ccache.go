package credential

import (
	"github.com/jcmturner/gokrb5/v8/credentials"
)

type CCache interface {
	Credential
	// CCache.
	CCache() *credentials.CCache
}

type ccacheCred struct {
	userName string
	realm    string
	ccache   *credentials.CCache
}

func (cc *ccacheCred) UserName() string {
	if cc != nil {
		return cc.userName
	}
	return ""
}

func (cc *ccacheCred) DomainName() string {
	if cc != nil {
		return cc.realm
	}
	return ""
}

func (cc *ccacheCred) Workstation() string {
	return ""
}

func (cc *ccacheCred) CCache() *credentials.CCache {
	if cc != nil && cc.ccache != nil {
		ccache := *cc.ccache
		return &ccache
	}
	return nil
}

func NewFromCCache(un string, ccache *credentials.CCache, opts ...Option) CCache {
	realm, un, _ := parseDomainUserWorkstation(un, opts...)
	return &ccacheCred{
		userName: un,
		realm:    realm,
		ccache:   ccache,
	}
}
