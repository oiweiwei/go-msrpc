package urlmon

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
	iremunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown/v0"
	iremunknown2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown2/v0"
	ibindctx "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon/ibindctx/v0"
	ienummoniker "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon/ienummoniker/v0"
	ienumstring "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon/ienumstring/v0"
	imoniker "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon/imoniker/v0"
	ipersist "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon/ipersist/v0"
	ipersistmoniker "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon/ipersistmoniker/v0"
	irunningobjecttable "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon/irunningobjecttable/v0"
	istream "github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon/istream/v0"
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
	_ = dcom_client.GoPackage
	_ = dcom.GoPackage
	_ = iremunknown.GoPackage
	_ = iremunknown2.GoPackage
	_ = istream.GoPackage
	_ = irunningobjecttable.GoPackage
	_ = ienumstring.GoPackage
	_ = ibindctx.GoPackage
	_ = imoniker.GoPackage
	_ = ienummoniker.GoPackage
	_ = ipersist.GoPackage
	_ = ipersistmoniker.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/urlmon"
)

// dcom/urlmon client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	Stream() istream.StreamClient
	RunningObjectTable() irunningobjecttable.RunningObjectTableClient
	EnumString() ienumstring.EnumStringClient
	BindContext() ibindctx.BindContextClient
	Moniker() imoniker.MonikerClient
	EnumMoniker() ienummoniker.EnumMonikerClient
	Persist() ipersist.PersistClient
	PersistMoniker() ipersistmoniker.PersistMonikerClient
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

	stream             istream.StreamClient
	runningObjectTable irunningobjecttable.RunningObjectTableClient
	enumString         ienumstring.EnumStringClient
	bindContext        ibindctx.BindContextClient
	moniker            imoniker.MonikerClient
	enumMoniker        ienummoniker.EnumMonikerClient
	persist            ipersist.PersistClient
	persistMoniker     ipersistmoniker.PersistMonikerClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) Stream() istream.StreamClient {
	return o.stream
}

func (o *xxx_DefaultClient) RunningObjectTable() irunningobjecttable.RunningObjectTableClient {
	return o.runningObjectTable
}

func (o *xxx_DefaultClient) EnumString() ienumstring.EnumStringClient {
	return o.enumString
}

func (o *xxx_DefaultClient) BindContext() ibindctx.BindContextClient {
	return o.bindContext
}

func (o *xxx_DefaultClient) Moniker() imoniker.MonikerClient {
	return o.moniker
}

func (o *xxx_DefaultClient) EnumMoniker() ienummoniker.EnumMonikerClient {
	return o.enumMoniker
}

func (o *xxx_DefaultClient) Persist() ipersist.PersistClient {
	return o.persist
}

func (o *xxx_DefaultClient) PersistMoniker() ipersistmoniker.PersistMonikerClient {
	return o.persistMoniker
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(istream.StreamSyntaxV0_0),
		dcerpc.WithAbstractSyntax(irunningobjecttable.RunningObjectTableSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumstring.EnumStringSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ibindctx.BindContextSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imoniker.MonikerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienummoniker.EnumMonikerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ipersist.PersistSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ipersistmoniker.PersistMonikerSyntaxV0_0),
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

	streamSubConn, err := sub.SubConn(ctx, istream.StreamSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		streamSubConn = sub
	}

	o.stream, err = istream.NewStreamClient(ctx, streamSubConn, append(opts, dcerpc.WithNoBind(streamSubConn))...)

	runningObjectTableSubConn, err := sub.SubConn(ctx, irunningobjecttable.RunningObjectTableSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		runningObjectTableSubConn = sub
	}

	o.runningObjectTable, err = irunningobjecttable.NewRunningObjectTableClient(ctx, runningObjectTableSubConn, append(opts, dcerpc.WithNoBind(runningObjectTableSubConn))...)

	enumStringSubConn, err := sub.SubConn(ctx, ienumstring.EnumStringSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumStringSubConn = sub
	}

	o.enumString, err = ienumstring.NewEnumStringClient(ctx, enumStringSubConn, append(opts, dcerpc.WithNoBind(enumStringSubConn))...)

	bindContextSubConn, err := sub.SubConn(ctx, ibindctx.BindContextSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		bindContextSubConn = sub
	}

	o.bindContext, err = ibindctx.NewBindContextClient(ctx, bindContextSubConn, append(opts, dcerpc.WithNoBind(bindContextSubConn))...)

	monikerSubConn, err := sub.SubConn(ctx, imoniker.MonikerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		monikerSubConn = sub
	}

	o.moniker, err = imoniker.NewMonikerClient(ctx, monikerSubConn, append(opts, dcerpc.WithNoBind(monikerSubConn))...)

	enumMonikerSubConn, err := sub.SubConn(ctx, ienummoniker.EnumMonikerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumMonikerSubConn = sub
	}

	o.enumMoniker, err = ienummoniker.NewEnumMonikerClient(ctx, enumMonikerSubConn, append(opts, dcerpc.WithNoBind(enumMonikerSubConn))...)

	persistSubConn, err := sub.SubConn(ctx, ipersist.PersistSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		persistSubConn = sub
	}

	o.persist, err = ipersist.NewPersistClient(ctx, persistSubConn, append(opts, dcerpc.WithNoBind(persistSubConn))...)

	persistMonikerSubConn, err := sub.SubConn(ctx, ipersistmoniker.PersistMonikerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		persistMonikerSubConn = sub
	}

	o.persistMoniker, err = ipersistmoniker.NewPersistMonikerClient(ctx, persistMonikerSubConn, append(opts, dcerpc.WithNoBind(persistMonikerSubConn))...)
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
		dcomClient:         o.dcomClient.IPID(ctx, ipid),
		stream:             o.stream.IPID(ctx, ipid),
		runningObjectTable: o.runningObjectTable.IPID(ctx, ipid),
		enumString:         o.enumString.IPID(ctx, ipid),
		bindContext:        o.bindContext.IPID(ctx, ipid),
		moniker:            o.moniker.IPID(ctx, ipid),
		enumMoniker:        o.enumMoniker.IPID(ctx, ipid),
		persist:            o.persist.IPID(ctx, ipid),
		persistMoniker:     o.persistMoniker.IPID(ctx, ipid),
		cc:                 o.cc,
	}
}
