// package ssp contains definitions for security service providers.
package ssp

import (
	"context"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	"github.com/oiweiwei/go-msrpc/ssp/krb5"
	"github.com/oiweiwei/go-msrpc/ssp/netlogon"
	"github.com/oiweiwei/go-msrpc/ssp/ntlm"
	"github.com/oiweiwei/go-msrpc/ssp/spnego"
)

var (
	// The NTLM authentication mechanism.
	NTLM = ntlm.Mechanism{}
	// The SPNEGO authentication mechanism.
	SPNEGO = spnego.Mechanism{}
	// The Kerberos Version 5 authentication mechanism.
	KRB5 = krb5.Mechanism{}
	// The Netlogon SSP Secure Channel mechanism.
	Netlogon = netlogon.Mechanism{}

	// The SPNEGO mechanism type.
	MechanismTypeSPNEGO = SPNEGO.Type()
	// The NTLM mechanism type.
	MechanismTypeNTLM = NTLM.Type()
	// The Kerberos Version 5 mechanism type.
	MechanismTypeKRB5 = KRB5.Type()
	// The Netlogon SSP Secure Channel mechanism type.
	MechanismTypeNetlogon = Netlogon.Type()
)

// MechanismTypeDefault function returns the default mechanism.
func MechanismTypeDefault(ctx context.Context) gssapi.OID {
	mech := gssapi.GetMechanism(ctx, nil)
	if mech != nil {
		return mech.Type()
	}
	return nil
}

func WithNTLM(cfg *ntlm.Config) gssapi.Option {
	return gssapi.WithMechanismConfig(cfg)
}

func WithSPNEGO(cfg *spnego.Config) gssapi.Option {
	return gssapi.WithMechanismConfig(cfg)
}

func WithKRB5(cfg *krb5.Config) gssapi.Option {
	return gssapi.WithMechanismConfig(cfg)
}

func WithNetlogon(cfg *netlogon.Config) gssapi.Option {
	return gssapi.WithMechanismConfig(cfg)
}
