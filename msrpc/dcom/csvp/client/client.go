package csvp

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
	iclustercleanup "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp/iclustercleanup/v0"
	iclusterfirewall "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp/iclusterfirewall/v0"
	iclusterlog "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp/iclusterlog/v0"
	iclusterlogex "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp/iclusterlogex/v0"
	iclusternetwork2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp/iclusternetwork2/v0"
	iclustersetup "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp/iclustersetup/v0"
	iclusterstorage2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp/iclusterstorage2/v0"
	iclusterstorage3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp/iclusterstorage3/v0"
	iclusterupdate "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp/iclusterupdate/v0"
	iremunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown/v0"
	iremunknown2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown2/v0"
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
	_ = iclusterstorage2.GoPackage
	_ = iclusterstorage3.GoPackage
	_ = iclusternetwork2.GoPackage
	_ = iclustercleanup.GoPackage
	_ = iclustersetup.GoPackage
	_ = iclusterlog.GoPackage
	_ = iclusterlogex.GoPackage
	_ = iclusterfirewall.GoPackage
	_ = iclusterupdate.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csvp"
)

// dcom/csvp client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	ClusterStorage2() iclusterstorage2.ClusterStorage2Client
	ClusterStorage3() iclusterstorage3.ClusterStorage3Client
	ClusterNetwork2() iclusternetwork2.ClusterNetwork2Client
	ClusterCleanup() iclustercleanup.ClusterCleanupClient
	ClusterSetup() iclustersetup.ClusterSetupClient
	ClusterLog() iclusterlog.ClusterLogClient
	ClusterLogEx() iclusterlogex.ClusterLogExClient
	ClusterFirewall() iclusterfirewall.ClusterFirewallClient
	ClusterUpdate() iclusterupdate.ClusterUpdateClient
	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Client
}
type xxx_DefaultClient struct {
	cc dcerpc.Conn

	dcomClient dcom_client.Client

	clusterStorage2 iclusterstorage2.ClusterStorage2Client
	clusterStorage3 iclusterstorage3.ClusterStorage3Client
	clusterNetwork2 iclusternetwork2.ClusterNetwork2Client
	clusterCleanup  iclustercleanup.ClusterCleanupClient
	clusterSetup    iclustersetup.ClusterSetupClient
	clusterLog      iclusterlog.ClusterLogClient
	clusterLogEx    iclusterlogex.ClusterLogExClient
	clusterFirewall iclusterfirewall.ClusterFirewallClient
	clusterUpdate   iclusterupdate.ClusterUpdateClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) ClusterStorage2() iclusterstorage2.ClusterStorage2Client {
	return o.clusterStorage2
}

func (o *xxx_DefaultClient) ClusterStorage3() iclusterstorage3.ClusterStorage3Client {
	return o.clusterStorage3
}

func (o *xxx_DefaultClient) ClusterNetwork2() iclusternetwork2.ClusterNetwork2Client {
	return o.clusterNetwork2
}

func (o *xxx_DefaultClient) ClusterCleanup() iclustercleanup.ClusterCleanupClient {
	return o.clusterCleanup
}

func (o *xxx_DefaultClient) ClusterSetup() iclustersetup.ClusterSetupClient {
	return o.clusterSetup
}

func (o *xxx_DefaultClient) ClusterLog() iclusterlog.ClusterLogClient {
	return o.clusterLog
}

func (o *xxx_DefaultClient) ClusterLogEx() iclusterlogex.ClusterLogExClient {
	return o.clusterLogEx
}

func (o *xxx_DefaultClient) ClusterFirewall() iclusterfirewall.ClusterFirewallClient {
	return o.clusterFirewall
}

func (o *xxx_DefaultClient) ClusterUpdate() iclusterupdate.ClusterUpdateClient {
	return o.clusterUpdate
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(iclusterstorage2.ClusterStorage2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iclusterstorage3.ClusterStorage3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iclusternetwork2.ClusterNetwork2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iclustercleanup.ClusterCleanupSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iclustersetup.ClusterSetupSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iclusterlog.ClusterLogSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iclusterlogex.ClusterLogExSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iclusterfirewall.ClusterFirewallSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iclusterupdate.ClusterUpdateSyntaxV0_0),
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

	clusterStorage2SubConn, err := sub.SubConn(ctx, iclusterstorage2.ClusterStorage2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clusterStorage2SubConn = sub
	}

	o.clusterStorage2, err = iclusterstorage2.NewClusterStorage2Client(ctx, clusterStorage2SubConn, append(opts, dcerpc.WithNoBind(clusterStorage2SubConn))...)

	clusterStorage3SubConn, err := sub.SubConn(ctx, iclusterstorage3.ClusterStorage3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clusterStorage3SubConn = sub
	}

	o.clusterStorage3, err = iclusterstorage3.NewClusterStorage3Client(ctx, clusterStorage3SubConn, append(opts, dcerpc.WithNoBind(clusterStorage3SubConn))...)

	clusterNetwork2SubConn, err := sub.SubConn(ctx, iclusternetwork2.ClusterNetwork2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clusterNetwork2SubConn = sub
	}

	o.clusterNetwork2, err = iclusternetwork2.NewClusterNetwork2Client(ctx, clusterNetwork2SubConn, append(opts, dcerpc.WithNoBind(clusterNetwork2SubConn))...)

	clusterCleanupSubConn, err := sub.SubConn(ctx, iclustercleanup.ClusterCleanupSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clusterCleanupSubConn = sub
	}

	o.clusterCleanup, err = iclustercleanup.NewClusterCleanupClient(ctx, clusterCleanupSubConn, append(opts, dcerpc.WithNoBind(clusterCleanupSubConn))...)

	clusterSetupSubConn, err := sub.SubConn(ctx, iclustersetup.ClusterSetupSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clusterSetupSubConn = sub
	}

	o.clusterSetup, err = iclustersetup.NewClusterSetupClient(ctx, clusterSetupSubConn, append(opts, dcerpc.WithNoBind(clusterSetupSubConn))...)

	clusterLogSubConn, err := sub.SubConn(ctx, iclusterlog.ClusterLogSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clusterLogSubConn = sub
	}

	o.clusterLog, err = iclusterlog.NewClusterLogClient(ctx, clusterLogSubConn, append(opts, dcerpc.WithNoBind(clusterLogSubConn))...)

	clusterLogExSubConn, err := sub.SubConn(ctx, iclusterlogex.ClusterLogExSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clusterLogExSubConn = sub
	}

	o.clusterLogEx, err = iclusterlogex.NewClusterLogExClient(ctx, clusterLogExSubConn, append(opts, dcerpc.WithNoBind(clusterLogExSubConn))...)

	clusterFirewallSubConn, err := sub.SubConn(ctx, iclusterfirewall.ClusterFirewallSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clusterFirewallSubConn = sub
	}

	o.clusterFirewall, err = iclusterfirewall.NewClusterFirewallClient(ctx, clusterFirewallSubConn, append(opts, dcerpc.WithNoBind(clusterFirewallSubConn))...)

	clusterUpdateSubConn, err := sub.SubConn(ctx, iclusterupdate.ClusterUpdateSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clusterUpdateSubConn = sub
	}

	o.clusterUpdate, err = iclusterupdate.NewClusterUpdateClient(ctx, clusterUpdateSubConn, append(opts, dcerpc.WithNoBind(clusterUpdateSubConn))...)
	return o, nil
}

func (o *xxx_DefaultClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClient) IPID(ctx context.Context, ipid *dcom.IPID) Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClient{
		dcomClient:      o.dcomClient.IPID(ctx, ipid),
		clusterStorage2: o.clusterStorage2.IPID(ctx, ipid),
		clusterStorage3: o.clusterStorage3.IPID(ctx, ipid),
		clusterNetwork2: o.clusterNetwork2.IPID(ctx, ipid),
		clusterCleanup:  o.clusterCleanup.IPID(ctx, ipid),
		clusterSetup:    o.clusterSetup.IPID(ctx, ipid),
		clusterLog:      o.clusterLog.IPID(ctx, ipid),
		clusterLogEx:    o.clusterLogEx.IPID(ctx, ipid),
		clusterFirewall: o.clusterFirewall.IPID(ctx, ipid),
		clusterUpdate:   o.clusterUpdate.IPID(ctx, ipid),
		cc:              o.cc,
	}
}
