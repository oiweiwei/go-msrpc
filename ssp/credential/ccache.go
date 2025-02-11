package credential

import (
	"github.com/oiweiwei/gokrb5.fork/v9/credentials"
)

type CCache interface {
	Credential
	// CCache.
	CCache() *credentials.CCache
}

type ccache struct {
	user
	ccache *credentials.CCache
}

func (cc *ccache) CCache() *credentials.CCache {
	if cc != nil && cc.ccache != nil {
		ccache := *cc.ccache
		return &ccache
	}
	return nil
}

func NewFromCCache(un string, cc *credentials.CCache, opts ...Option) CCache {
	return &ccache{
		user:   parseUser(un, opts...),
		ccache: cc,
	}
}
