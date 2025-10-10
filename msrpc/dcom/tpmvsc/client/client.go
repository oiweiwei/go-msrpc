package tpmvsc

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
	itpmvirtualsmartcardmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc/itpmvirtualsmartcardmanager/v0"
	itpmvirtualsmartcardmanager2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc/itpmvirtualsmartcardmanager2/v0"
	itpmvirtualsmartcardmanager3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc/itpmvirtualsmartcardmanager3/v0"
	itpmvirtualsmartcardmanagerstatuscallback "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc/itpmvirtualsmartcardmanagerstatuscallback/v0"
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
	_ = itpmvirtualsmartcardmanagerstatuscallback.GoPackage
	_ = itpmvirtualsmartcardmanager.GoPackage
	_ = itpmvirtualsmartcardmanager2.GoPackage
	_ = itpmvirtualsmartcardmanager3.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/tpmvsc"
)

// dcom/tpmvsc client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	VirtualSmartCardManagerStatusCallback() itpmvirtualsmartcardmanagerstatuscallback.VirtualSmartCardManagerStatusCallbackClient
	VirtualSmartCardManager() itpmvirtualsmartcardmanager.VirtualSmartCardManagerClient
	VirtualSmartCardManager2() itpmvirtualsmartcardmanager2.VirtualSmartCardManager2Client
	VirtualSmartCardManager3() itpmvirtualsmartcardmanager3.VirtualSmartCardManager3Client
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

	virtualSmartCardManagerStatusCallback itpmvirtualsmartcardmanagerstatuscallback.VirtualSmartCardManagerStatusCallbackClient
	virtualSmartCardManager               itpmvirtualsmartcardmanager.VirtualSmartCardManagerClient
	virtualSmartCardManager2              itpmvirtualsmartcardmanager2.VirtualSmartCardManager2Client
	virtualSmartCardManager3              itpmvirtualsmartcardmanager3.VirtualSmartCardManager3Client
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) VirtualSmartCardManagerStatusCallback() itpmvirtualsmartcardmanagerstatuscallback.VirtualSmartCardManagerStatusCallbackClient {
	return o.virtualSmartCardManagerStatusCallback
}

func (o *xxx_DefaultClient) VirtualSmartCardManager() itpmvirtualsmartcardmanager.VirtualSmartCardManagerClient {
	return o.virtualSmartCardManager
}

func (o *xxx_DefaultClient) VirtualSmartCardManager2() itpmvirtualsmartcardmanager2.VirtualSmartCardManager2Client {
	return o.virtualSmartCardManager2
}

func (o *xxx_DefaultClient) VirtualSmartCardManager3() itpmvirtualsmartcardmanager3.VirtualSmartCardManager3Client {
	return o.virtualSmartCardManager3
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(itpmvirtualsmartcardmanagerstatuscallback.VirtualSmartCardManagerStatusCallbackSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itpmvirtualsmartcardmanager.VirtualSmartCardManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itpmvirtualsmartcardmanager2.VirtualSmartCardManager2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(itpmvirtualsmartcardmanager3.VirtualSmartCardManager3SyntaxV0_0),
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

	virtualSmartCardManagerStatusCallbackSubConn, err := sub.SubConn(ctx, itpmvirtualsmartcardmanagerstatuscallback.VirtualSmartCardManagerStatusCallbackSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		virtualSmartCardManagerStatusCallbackSubConn = sub
	}

	o.virtualSmartCardManagerStatusCallback, err = itpmvirtualsmartcardmanagerstatuscallback.NewVirtualSmartCardManagerStatusCallbackClient(ctx, virtualSmartCardManagerStatusCallbackSubConn, append(opts, dcerpc.WithNoBind(virtualSmartCardManagerStatusCallbackSubConn))...)

	virtualSmartCardManagerSubConn, err := sub.SubConn(ctx, itpmvirtualsmartcardmanager.VirtualSmartCardManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		virtualSmartCardManagerSubConn = sub
	}

	o.virtualSmartCardManager, err = itpmvirtualsmartcardmanager.NewVirtualSmartCardManagerClient(ctx, virtualSmartCardManagerSubConn, append(opts, dcerpc.WithNoBind(virtualSmartCardManagerSubConn))...)

	virtualSmartCardManager2SubConn, err := sub.SubConn(ctx, itpmvirtualsmartcardmanager2.VirtualSmartCardManager2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		virtualSmartCardManager2SubConn = sub
	}

	o.virtualSmartCardManager2, err = itpmvirtualsmartcardmanager2.NewVirtualSmartCardManager2Client(ctx, virtualSmartCardManager2SubConn, append(opts, dcerpc.WithNoBind(virtualSmartCardManager2SubConn))...)

	virtualSmartCardManager3SubConn, err := sub.SubConn(ctx, itpmvirtualsmartcardmanager3.VirtualSmartCardManager3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		virtualSmartCardManager3SubConn = sub
	}

	o.virtualSmartCardManager3, err = itpmvirtualsmartcardmanager3.NewVirtualSmartCardManager3Client(ctx, virtualSmartCardManager3SubConn, append(opts, dcerpc.WithNoBind(virtualSmartCardManager3SubConn))...)
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
		dcomClient:                            o.dcomClient.IPID(ctx, ipid),
		virtualSmartCardManagerStatusCallback: o.virtualSmartCardManagerStatusCallback.IPID(ctx, ipid),
		virtualSmartCardManager:               o.virtualSmartCardManager.IPID(ctx, ipid),
		virtualSmartCardManager2:              o.virtualSmartCardManager2.IPID(ctx, ipid),
		virtualSmartCardManager3:              o.virtualSmartCardManager3.IPID(ctx, ipid),
		cc:                                    o.cc,
	}
}
