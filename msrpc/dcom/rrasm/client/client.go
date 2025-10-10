package rrasm

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
	iremunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown/v0"
	iremunknown2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown2/v0"
	iremoteicficsconfig "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm/iremoteicficsconfig/v0"
	iremoteipv6config "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm/iremoteipv6config/v0"
	iremotenetworkconfig "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm/iremotenetworkconfig/v0"
	iremoterouterrestart "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm/iremoterouterrestart/v0"
	iremotesetdnsconfig "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm/iremotesetdnsconfig/v0"
	iremotesstpcertcheck "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm/iremotesstpcertcheck/v0"
	iremotestringidconfig "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm/iremotestringidconfig/v0"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

var (
	_ = context.Background
	_ = fmt.Errorf
	_ = utf16.Encode
	_ = strings.TrimPrefix
	_ = ndr.ZeroString
	_ = (*uuid.UUID)(nil)
	_ = (*dcerpc.SyntaxID)(nil)
	_ = (*errors.Error)(nil)
	_ = dcom_client.GoPackage
	_ = dcom.GoPackage
	_ = iremunknown.GoPackage
	_ = iremunknown2.GoPackage
	_ = iremotenetworkconfig.GoPackage
	_ = iremoterouterrestart.GoPackage
	_ = iremotesetdnsconfig.GoPackage
	_ = iremoteicficsconfig.GoPackage
	_ = iremotestringidconfig.GoPackage
	_ = iremoteipv6config.GoPackage
	_ = iremotesstpcertcheck.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

// dcom/rrasm client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	RemoteNetworkConfig() iremotenetworkconfig.RemoteNetworkConfigClient
	RemoteRouterRestart() iremoterouterrestart.RemoteRouterRestartClient
	RemoteSetDNSConfig() iremotesetdnsconfig.RemoteSetDNSConfigClient
	RemoteIcficsConfig() iremoteicficsconfig.RemoteIcficsConfigClient
	RemoteStringIDConfig() iremotestringidconfig.RemoteStringIDConfigClient
	RemoteIPv6Config() iremoteipv6config.RemoteIPv6ConfigClient
	RemoteSstpCertCheck() iremotesstpcertcheck.RemoteSstpCertCheckClient
	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Client
}
type xxx_DefaultClient struct {
	cc dcerpc.Conn

	dcomClient dcom_client.Client

	remoteNetworkConfig  iremotenetworkconfig.RemoteNetworkConfigClient
	remoteRouterRestart  iremoterouterrestart.RemoteRouterRestartClient
	remoteSetDNSConfig   iremotesetdnsconfig.RemoteSetDNSConfigClient
	remoteIcficsConfig   iremoteicficsconfig.RemoteIcficsConfigClient
	remoteStringIDConfig iremotestringidconfig.RemoteStringIDConfigClient
	remoteIPv6Config     iremoteipv6config.RemoteIPv6ConfigClient
	remoteSstpCertCheck  iremotesstpcertcheck.RemoteSstpCertCheckClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) RemoteNetworkConfig() iremotenetworkconfig.RemoteNetworkConfigClient {
	return o.remoteNetworkConfig
}

func (o *xxx_DefaultClient) RemoteRouterRestart() iremoterouterrestart.RemoteRouterRestartClient {
	return o.remoteRouterRestart
}

func (o *xxx_DefaultClient) RemoteSetDNSConfig() iremotesetdnsconfig.RemoteSetDNSConfigClient {
	return o.remoteSetDNSConfig
}

func (o *xxx_DefaultClient) RemoteIcficsConfig() iremoteicficsconfig.RemoteIcficsConfigClient {
	return o.remoteIcficsConfig
}

func (o *xxx_DefaultClient) RemoteStringIDConfig() iremotestringidconfig.RemoteStringIDConfigClient {
	return o.remoteStringIDConfig
}

func (o *xxx_DefaultClient) RemoteIPv6Config() iremoteipv6config.RemoteIPv6ConfigClient {
	return o.remoteIPv6Config
}

func (o *xxx_DefaultClient) RemoteSstpCertCheck() iremotesstpcertcheck.RemoteSstpCertCheckClient {
	return o.remoteSstpCertCheck
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(iremotenetworkconfig.RemoteNetworkConfigSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremoterouterrestart.RemoteRouterRestartSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremotesetdnsconfig.RemoteSetDNSConfigSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremoteicficsconfig.RemoteIcficsConfigSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremotestringidconfig.RemoteStringIDConfigSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremoteipv6config.RemoteIPv6ConfigSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremotesstpcertcheck.RemoteSstpCertCheckSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremunknown.RemoteUnknownSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremunknown2.RemoteUnknown2SyntaxV0_0),
	)

	cc, err := cc.Bind(ctx, opts...)
	if err != nil {
		return nil, err
	}

	o := &xxx_DefaultClient{cc: cc}

	dcomClient, err := dcom_client.NewClient(ctx, cc, append(opts, dcerpc.WithNoBind(cc))...)
	if err != nil {
		return nil, err
	}
	o.dcomClient = dcomClient

	sub, ok := cc.(dcerpc.SubConn)
	if !ok {
		return nil, fmt.Errorf("sub-conn is not supported")
	}

	remoteNetworkConfigSubConn, err := sub.SubConn(ctx, iremotenetworkconfig.RemoteNetworkConfigSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteNetworkConfigSubConn = sub
	}

	o.remoteNetworkConfig, err = iremotenetworkconfig.NewRemoteNetworkConfigClient(ctx, remoteNetworkConfigSubConn, append(opts, dcerpc.WithNoBind(remoteNetworkConfigSubConn))...)

	remoteRouterRestartSubConn, err := sub.SubConn(ctx, iremoterouterrestart.RemoteRouterRestartSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteRouterRestartSubConn = sub
	}

	o.remoteRouterRestart, err = iremoterouterrestart.NewRemoteRouterRestartClient(ctx, remoteRouterRestartSubConn, append(opts, dcerpc.WithNoBind(remoteRouterRestartSubConn))...)

	remoteSetDNSConfigSubConn, err := sub.SubConn(ctx, iremotesetdnsconfig.RemoteSetDNSConfigSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteSetDNSConfigSubConn = sub
	}

	o.remoteSetDNSConfig, err = iremotesetdnsconfig.NewRemoteSetDNSConfigClient(ctx, remoteSetDNSConfigSubConn, append(opts, dcerpc.WithNoBind(remoteSetDNSConfigSubConn))...)

	remoteIcficsConfigSubConn, err := sub.SubConn(ctx, iremoteicficsconfig.RemoteIcficsConfigSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteIcficsConfigSubConn = sub
	}

	o.remoteIcficsConfig, err = iremoteicficsconfig.NewRemoteIcficsConfigClient(ctx, remoteIcficsConfigSubConn, append(opts, dcerpc.WithNoBind(remoteIcficsConfigSubConn))...)

	remoteStringIDConfigSubConn, err := sub.SubConn(ctx, iremotestringidconfig.RemoteStringIDConfigSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteStringIDConfigSubConn = sub
	}

	o.remoteStringIDConfig, err = iremotestringidconfig.NewRemoteStringIDConfigClient(ctx, remoteStringIDConfigSubConn, append(opts, dcerpc.WithNoBind(remoteStringIDConfigSubConn))...)

	remoteIPv6ConfigSubConn, err := sub.SubConn(ctx, iremoteipv6config.RemoteIPv6ConfigSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteIPv6ConfigSubConn = sub
	}

	o.remoteIPv6Config, err = iremoteipv6config.NewRemoteIPv6ConfigClient(ctx, remoteIPv6ConfigSubConn, append(opts, dcerpc.WithNoBind(remoteIPv6ConfigSubConn))...)

	remoteSstpCertCheckSubConn, err := sub.SubConn(ctx, iremotesstpcertcheck.RemoteSstpCertCheckSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteSstpCertCheckSubConn = sub
	}

	o.remoteSstpCertCheck, err = iremotesstpcertcheck.NewRemoteSstpCertCheckClient(ctx, remoteSstpCertCheckSubConn, append(opts, dcerpc.WithNoBind(remoteSstpCertCheckSubConn))...)
	return o, nil
}

func (o *xxx_DefaultClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClient) IPID(ctx context.Context, ipid *dcom.IPID) Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClient{
		dcomClient:           o.dcomClient.IPID(ctx, ipid),
		remoteNetworkConfig:  o.remoteNetworkConfig.IPID(ctx, ipid),
		remoteRouterRestart:  o.remoteRouterRestart.IPID(ctx, ipid),
		remoteSetDNSConfig:   o.remoteSetDNSConfig.IPID(ctx, ipid),
		remoteIcficsConfig:   o.remoteIcficsConfig.IPID(ctx, ipid),
		remoteStringIDConfig: o.remoteStringIDConfig.IPID(ctx, ipid),
		remoteIPv6Config:     o.remoteIPv6Config.IPID(ctx, ipid),
		remoteSstpCertCheck:  o.remoteSstpCertCheck.IPID(ctx, ipid),
		cc:                   o.cc,
	}
}
