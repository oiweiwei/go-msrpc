package dfsrh

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
	iadproxy "github.com/oiweiwei/go-msrpc/msrpc/dcom/dfsrh/iadproxy/v0"
	iadproxy2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/dfsrh/iadproxy2/v0"
	iserverhealthreport "github.com/oiweiwei/go-msrpc/msrpc/dcom/dfsrh/iserverhealthreport/v0"
	iserverhealthreport2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/dfsrh/iserverhealthreport2/v0"
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
	_ = iadproxy.GoPackage
	_ = iadproxy2.GoPackage
	_ = iserverhealthreport.GoPackage
	_ = iserverhealthreport2.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/dfsrh"
)

// dcom/dfsrh client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	IADProxy() iadproxy.IADProxyClient
	IADProxy2() iadproxy2.IADProxy2Client
	ServerHealthReport() iserverhealthreport.ServerHealthReportClient
	ServerHealthReport2() iserverhealthreport2.ServerHealthReport2Client
	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Client
}
type xxx_DefaultClient struct {
	cc dcerpc.Conn

	dcomClient dcom_client.Client

	iadProxy            iadproxy.IADProxyClient
	iadProxy2           iadproxy2.IADProxy2Client
	serverHealthReport  iserverhealthreport.ServerHealthReportClient
	serverHealthReport2 iserverhealthreport2.ServerHealthReport2Client
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) IADProxy() iadproxy.IADProxyClient {
	return o.iadProxy
}

func (o *xxx_DefaultClient) IADProxy2() iadproxy2.IADProxy2Client {
	return o.iadProxy2
}

func (o *xxx_DefaultClient) ServerHealthReport() iserverhealthreport.ServerHealthReportClient {
	return o.serverHealthReport
}

func (o *xxx_DefaultClient) ServerHealthReport2() iserverhealthreport2.ServerHealthReport2Client {
	return o.serverHealthReport2
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(iadproxy.IADProxySyntaxV0_0),
		dcerpc.WithAbstractSyntax(iadproxy2.IADProxy2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iserverhealthreport.ServerHealthReportSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iserverhealthreport2.ServerHealthReport2SyntaxV0_0),
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

	iadProxySubConn, err := sub.SubConn(ctx, iadproxy.IADProxySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iadProxySubConn = sub
	}

	o.iadProxy, err = iadproxy.NewIADProxyClient(ctx, iadProxySubConn, append(opts, dcerpc.WithNoBind(iadProxySubConn))...)

	iadProxy2SubConn, err := sub.SubConn(ctx, iadproxy2.IADProxy2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iadProxy2SubConn = sub
	}

	o.iadProxy2, err = iadproxy2.NewIADProxy2Client(ctx, iadProxy2SubConn, append(opts, dcerpc.WithNoBind(iadProxy2SubConn))...)

	serverHealthReportSubConn, err := sub.SubConn(ctx, iserverhealthreport.ServerHealthReportSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		serverHealthReportSubConn = sub
	}

	o.serverHealthReport, err = iserverhealthreport.NewServerHealthReportClient(ctx, serverHealthReportSubConn, append(opts, dcerpc.WithNoBind(serverHealthReportSubConn))...)

	serverHealthReport2SubConn, err := sub.SubConn(ctx, iserverhealthreport2.ServerHealthReport2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		serverHealthReport2SubConn = sub
	}

	o.serverHealthReport2, err = iserverhealthreport2.NewServerHealthReport2Client(ctx, serverHealthReport2SubConn, append(opts, dcerpc.WithNoBind(serverHealthReport2SubConn))...)
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
		dcomClient:          o.dcomClient.IPID(ctx, ipid),
		iadProxy:            o.iadProxy.IPID(ctx, ipid),
		iadProxy2:           o.iadProxy2.IPID(ctx, ipid),
		serverHealthReport:  o.serverHealthReport.IPID(ctx, ipid),
		serverHealthReport2: o.serverHealthReport2.IPID(ctx, ipid),
		cc:                  o.cc,
	}
}
