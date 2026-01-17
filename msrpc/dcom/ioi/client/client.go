package ioi

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
	imanagedobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/ioi/imanagedobject/v0"
	iremotedispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/ioi/iremotedispatch/v0"
	iservicedcomponentinfo "github.com/oiweiwei/go-msrpc/msrpc/dcom/ioi/iservicedcomponentinfo/v0"
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
	_ = dcom_client.GoPackage
	_ = dcom.GoPackage
	_ = iremunknown.GoPackage
	_ = iremunknown2.GoPackage
	_ = imanagedobject.GoPackage
	_ = iremotedispatch.GoPackage
	_ = iservicedcomponentinfo.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/ioi"
)

// dcom/ioi client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	ManagedObject() imanagedobject.ManagedObjectClient
	RemoteDispatch() iremotedispatch.RemoteDispatchClient
	ServicedComponentInfo() iservicedcomponentinfo.ServicedComponentInfoClient
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

	managedObject         imanagedobject.ManagedObjectClient
	remoteDispatch        iremotedispatch.RemoteDispatchClient
	servicedComponentInfo iservicedcomponentinfo.ServicedComponentInfoClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) ManagedObject() imanagedobject.ManagedObjectClient {
	return o.managedObject
}

func (o *xxx_DefaultClient) RemoteDispatch() iremotedispatch.RemoteDispatchClient {
	return o.remoteDispatch
}

func (o *xxx_DefaultClient) ServicedComponentInfo() iservicedcomponentinfo.ServicedComponentInfoClient {
	return o.servicedComponentInfo
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(imanagedobject.ManagedObjectSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremotedispatch.RemoteDispatchSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iservicedcomponentinfo.ServicedComponentInfoSyntaxV0_0),
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

	managedObjectSubConn, err := sub.SubConn(ctx, imanagedobject.ManagedObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		managedObjectSubConn = sub
	}

	o.managedObject, err = imanagedobject.NewManagedObjectClient(ctx, managedObjectSubConn, append(opts, dcerpc.WithNoBind(managedObjectSubConn))...)

	remoteDispatchSubConn, err := sub.SubConn(ctx, iremotedispatch.RemoteDispatchSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteDispatchSubConn = sub
	}

	o.remoteDispatch, err = iremotedispatch.NewRemoteDispatchClient(ctx, remoteDispatchSubConn, append(opts, dcerpc.WithNoBind(remoteDispatchSubConn))...)

	servicedComponentInfoSubConn, err := sub.SubConn(ctx, iservicedcomponentinfo.ServicedComponentInfoSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		servicedComponentInfoSubConn = sub
	}

	o.servicedComponentInfo, err = iservicedcomponentinfo.NewServicedComponentInfoClient(ctx, servicedComponentInfoSubConn, append(opts, dcerpc.WithNoBind(servicedComponentInfoSubConn))...)
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
		dcomClient:            o.dcomClient.IPID(ctx, ipid),
		managedObject:         o.managedObject.IPID(ctx, ipid),
		remoteDispatch:        o.remoteDispatch.IPID(ctx, ipid),
		servicedComponentInfo: o.servicedComponentInfo.IPID(ctx, ipid),
		cc:                    o.cc,
	}
}
