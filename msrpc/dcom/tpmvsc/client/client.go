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
	TpmVirtualSmartCardManagerStatusCallback() itpmvirtualsmartcardmanagerstatuscallback.TpmVirtualSmartCardManagerStatusCallbackClient
	TpmVirtualSmartCardManager() itpmvirtualsmartcardmanager.TpmVirtualSmartCardManagerClient
	TpmVirtualSmartCardManager2() itpmvirtualsmartcardmanager2.TpmVirtualSmartCardManager2Client
	TpmVirtualSmartCardManager3() itpmvirtualsmartcardmanager3.TpmVirtualSmartCardManager3Client
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

	tpmVirtualSmartCardManagerStatusCallback itpmvirtualsmartcardmanagerstatuscallback.TpmVirtualSmartCardManagerStatusCallbackClient
	tpmVirtualSmartCardManager               itpmvirtualsmartcardmanager.TpmVirtualSmartCardManagerClient
	tpmVirtualSmartCardManager2              itpmvirtualsmartcardmanager2.TpmVirtualSmartCardManager2Client
	tpmVirtualSmartCardManager3              itpmvirtualsmartcardmanager3.TpmVirtualSmartCardManager3Client
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) TpmVirtualSmartCardManagerStatusCallback() itpmvirtualsmartcardmanagerstatuscallback.TpmVirtualSmartCardManagerStatusCallbackClient {
	return o.tpmVirtualSmartCardManagerStatusCallback
}

func (o *xxx_DefaultClient) TpmVirtualSmartCardManager() itpmvirtualsmartcardmanager.TpmVirtualSmartCardManagerClient {
	return o.tpmVirtualSmartCardManager
}

func (o *xxx_DefaultClient) TpmVirtualSmartCardManager2() itpmvirtualsmartcardmanager2.TpmVirtualSmartCardManager2Client {
	return o.tpmVirtualSmartCardManager2
}

func (o *xxx_DefaultClient) TpmVirtualSmartCardManager3() itpmvirtualsmartcardmanager3.TpmVirtualSmartCardManager3Client {
	return o.tpmVirtualSmartCardManager3
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(itpmvirtualsmartcardmanagerstatuscallback.TpmVirtualSmartCardManagerStatusCallbackSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itpmvirtualsmartcardmanager.TpmVirtualSmartCardManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itpmvirtualsmartcardmanager2.TpmVirtualSmartCardManager2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(itpmvirtualsmartcardmanager3.TpmVirtualSmartCardManager3SyntaxV0_0),
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

	tpmVirtualSmartCardManagerStatusCallbackSubConn, err := sub.SubConn(ctx, itpmvirtualsmartcardmanagerstatuscallback.TpmVirtualSmartCardManagerStatusCallbackSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		tpmVirtualSmartCardManagerStatusCallbackSubConn = sub
	}

	o.tpmVirtualSmartCardManagerStatusCallback, err = itpmvirtualsmartcardmanagerstatuscallback.NewTpmVirtualSmartCardManagerStatusCallbackClient(ctx, tpmVirtualSmartCardManagerStatusCallbackSubConn, append(opts, dcerpc.WithNoBind(tpmVirtualSmartCardManagerStatusCallbackSubConn))...)

	tpmVirtualSmartCardManagerSubConn, err := sub.SubConn(ctx, itpmvirtualsmartcardmanager.TpmVirtualSmartCardManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		tpmVirtualSmartCardManagerSubConn = sub
	}

	o.tpmVirtualSmartCardManager, err = itpmvirtualsmartcardmanager.NewTpmVirtualSmartCardManagerClient(ctx, tpmVirtualSmartCardManagerSubConn, append(opts, dcerpc.WithNoBind(tpmVirtualSmartCardManagerSubConn))...)

	tpmVirtualSmartCardManager2SubConn, err := sub.SubConn(ctx, itpmvirtualsmartcardmanager2.TpmVirtualSmartCardManager2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		tpmVirtualSmartCardManager2SubConn = sub
	}

	o.tpmVirtualSmartCardManager2, err = itpmvirtualsmartcardmanager2.NewTpmVirtualSmartCardManager2Client(ctx, tpmVirtualSmartCardManager2SubConn, append(opts, dcerpc.WithNoBind(tpmVirtualSmartCardManager2SubConn))...)

	tpmVirtualSmartCardManager3SubConn, err := sub.SubConn(ctx, itpmvirtualsmartcardmanager3.TpmVirtualSmartCardManager3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		tpmVirtualSmartCardManager3SubConn = sub
	}

	o.tpmVirtualSmartCardManager3, err = itpmvirtualsmartcardmanager3.NewTpmVirtualSmartCardManager3Client(ctx, tpmVirtualSmartCardManager3SubConn, append(opts, dcerpc.WithNoBind(tpmVirtualSmartCardManager3SubConn))...)
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
		dcomClient:                               o.dcomClient.IPID(ctx, ipid),
		tpmVirtualSmartCardManagerStatusCallback: o.tpmVirtualSmartCardManagerStatusCallback.IPID(ctx, ipid),
		tpmVirtualSmartCardManager:               o.tpmVirtualSmartCardManager.IPID(ctx, ipid),
		tpmVirtualSmartCardManager2:              o.tpmVirtualSmartCardManager2.IPID(ctx, ipid),
		tpmVirtualSmartCardManager3:              o.tpmVirtualSmartCardManager3.IPID(ctx, ipid),
		cc:                                       o.cc,
	}
}
