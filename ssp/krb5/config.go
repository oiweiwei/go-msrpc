package krb5

import (
	"os"
	"strings"
	"time"

	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
	"github.com/jcmturner/gokrb5/v8/service"
	"github.com/jcmturner/gokrb5/v8/types"

	"github.com/oiweiwei/go-msrpc/ssp/credential"
)

// The generic credential.
type Credential = credential.Credential

// The Kerberos Version 5 Configuration.
type Config struct {
	// IsServer.
	IsServer bool
	// The client credential.
	Credential Credential
	// The kerberos config file.
	KRB5Config *config.Config
	// The kerberos config file path.
	KRB5ConfigPath string
	// The credentials cache file path.
	CCachePath string
	// The GSSAPI flags.
	Flags []int
	// The Kerberos Options.
	APOptions []int

	// common.
	DCEStyle bool

	// service settings.

	// KeytabPrincipal used to override the principal name
	// used to find the key in the keytab.
	KeytabPrincipal string
	// SName used provide a specific service name to the
	// service settings.
	SName string
	// RequireHostAddr indicates if the service should require
	// the host address to be included in the ticket.
	RequireHostAddr bool
	// DisablePACDecoding used to configure service side to
	// enable/disable PAC decoding if the PAC is present.
	// Defaults to enabled if not specified.
	DisablePACDecoding bool
	// ClientAddress used to configure service side with the
	// clients host address to be used during validation.
	ClientAddress *types.HostAddress
	// MaxClockSkew returns the maximum acceptable clock skew
	// between the service and the issue time of kerberos tickets.
	// If none is defined a duration of 5 minutes is returned.
	MaxClockSkew time.Duration

	// client settings.

	// DisablePAFXFAST used to configure the client to not use
	// PA_FX_FAST.
	DisablePAFXFAST bool
	// AssumePreAuthentication used to configure the client to
	// assume pre-authentication is required.
	AssumePreAuthentication bool
}

var DefaultKRB5ConfPath = "/etc/krb5.conf"

// NewConfig function returns the default configuration. (or
// configuration under KRB5_CONFIG environment variable).
func NewConfig() *Config {

	var p string

	if p = os.Getenv("KRB5_CONFIG"); p == "" {
		p = DefaultKRB5ConfPath
	}

	if c, _ := config.Load(p); c != nil {
		return &Config{KRB5Config: ParsedLibDefaults(c), DCEStyle: true, DisablePAFXFAST: true}
	}

	c := config.New()

	return &Config{KRB5Config: ParsedLibDefaults(c), DCEStyle: true, DisablePAFXFAST: true}
}

func LoadKRB5Conf(p string) (*config.Config, error) {

	if p != "" {
		c, err := config.Load(p)
		if err != nil {
			return nil, err
		}
		return ParsedLibDefaults(c), nil
	}

	if p = os.Getenv("KRB5_CONFIG"); p == "" {
		p = DefaultKRB5ConfPath
	}

	c, _ := config.Load(p)
	return ParsedLibDefaults(c), nil
}

func AddEnctypeID(c *config.Config, etype int32) *config.Config {
	l := c.LibDefaults
	c.LibDefaults = l
	return c
}

func ParsedLibDefaults(c *config.Config) *config.Config {
	l := c.LibDefaults
	l.DefaultTGSEnctypeIDs = parseETypes(l.DefaultTGSEnctypes, l.AllowWeakCrypto)
	l.DefaultTktEnctypeIDs = parseETypes(l.DefaultTktEnctypes, l.AllowWeakCrypto)
	l.PermittedEnctypeIDs = parseETypes(l.PermittedEnctypes, l.AllowWeakCrypto)
	c.LibDefaults = l
	return c
}

func parseETypes(s []string, w bool) []int32 {
	var eti []int32
	for _, et := range s {
		if !w {
			var weak bool
			for _, wet := range strings.Fields(config.WeakETypeList) {
				if et == wet {
					weak = true
					break
				}
			}
			if weak {
				continue
			}
		}
		i := etypeID.EtypeSupported(et)
		if i != 0 {
			eti = append(eti, i)
		}
	}
	return eti
}

func IsValidCredential(cred any) bool {
	if _, ok := cred.(credential.Keytab); ok {
		return true
	}

	if _, ok := cred.(credential.Password); ok {
		return true
	}

	if _, ok := cred.(credential.NTHash); ok {
		return true
	}

	return false
}

// ServiceSettings function returns the set of options for the
// service settings.
func (c *Config) ServiceSettings() []func(*service.Settings) {

	o := make([]func(*service.Settings), 0, 6)

	if c.KeytabPrincipal != "" {
		o = append(o, service.KeytabPrincipal(c.KeytabPrincipal))
	}

	if c.SName != "" {
		o = append(o, service.SName(c.SName))
	}

	if c.RequireHostAddr {
		o = append(o, service.RequireHostAddr(c.RequireHostAddr))
	}

	if c.DisablePACDecoding {
		o = append(o, service.DecodePAC(false))
	}

	if c.ClientAddress != nil {
		o = append(o, service.ClientAddress(*c.ClientAddress))
	}

	if c.MaxClockSkew != 0 {
		o = append(o, service.MaxClockSkew(c.MaxClockSkew))
	}

	return o
}

// ClientSettings function returns the set of options for the
// kerberos client.
func (c *Config) ClientSettings() []func(*client.Settings) {

	o := make([]func(*client.Settings), 0, 2)

	if c.DisablePAFXFAST {
		o = append(o, client.DisablePAFXFAST(c.DisablePAFXFAST))
	}

	if c.AssumePreAuthentication {
		o = append(o, client.AssumePreAuthentication(c.AssumePreAuthentication))
	}

	return o
}
