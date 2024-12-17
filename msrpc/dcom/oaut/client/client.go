package oaut

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
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
	ienumvariant "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/ienumvariant/v0"
	itypecomp "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/itypecomp/v0"
	itypeinfo "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/itypeinfo/v0"
	itypeinfo2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/itypeinfo2/v0"
	itypelib "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/itypelib/v0"
	itypelib2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/itypelib2/v0"
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
	_ = idispatch.GoPackage
	_ = ienumvariant.GoPackage
	_ = itypecomp.GoPackage
	_ = itypeinfo.GoPackage
	_ = itypeinfo2.GoPackage
	_ = itypelib.GoPackage
	_ = itypelib2.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/oaut"
)

// dcom/oaut client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	Dispatch() idispatch.DispatchClient
	EnumVariant() ienumvariant.EnumVariantClient
	TypeComp() itypecomp.TypeCompClient
	TypeInfo() itypeinfo.TypeInfoClient
	TypeInfo2() itypeinfo2.TypeInfo2Client
	TypeLib() itypelib.TypeLibClient
	TypeLib2() itypelib2.TypeLib2Client
	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Client
}
type xxx_DefaultClient struct {
	cc dcerpc.Conn

	dcomClient dcom_client.Client

	dispatch    idispatch.DispatchClient
	enumVariant ienumvariant.EnumVariantClient
	typeComp    itypecomp.TypeCompClient
	typeInfo    itypeinfo.TypeInfoClient
	typeInfo2   itypeinfo2.TypeInfo2Client
	typeLib     itypelib.TypeLibClient
	typeLib2    itypelib2.TypeLib2Client
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) Dispatch() idispatch.DispatchClient {
	return o.dispatch
}

func (o *xxx_DefaultClient) EnumVariant() ienumvariant.EnumVariantClient {
	return o.enumVariant
}

func (o *xxx_DefaultClient) TypeComp() itypecomp.TypeCompClient {
	return o.typeComp
}

func (o *xxx_DefaultClient) TypeInfo() itypeinfo.TypeInfoClient {
	return o.typeInfo
}

func (o *xxx_DefaultClient) TypeInfo2() itypeinfo2.TypeInfo2Client {
	return o.typeInfo2
}

func (o *xxx_DefaultClient) TypeLib() itypelib.TypeLibClient {
	return o.typeLib
}

func (o *xxx_DefaultClient) TypeLib2() itypelib2.TypeLib2Client {
	return o.typeLib2
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(idispatch.DispatchSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumvariant.EnumVariantSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itypecomp.TypeCompSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itypeinfo.TypeInfoSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itypeinfo2.TypeInfo2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(itypelib.TypeLibSyntaxV0_0),
		dcerpc.WithAbstractSyntax(itypelib2.TypeLib2SyntaxV0_0),
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

	dispatchSubConn, err := sub.SubConn(ctx, idispatch.DispatchSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		dispatchSubConn = sub
	}

	o.dispatch, err = idispatch.NewDispatchClient(ctx, dispatchSubConn, append(opts, dcerpc.WithNoBind(dispatchSubConn))...)

	enumVariantSubConn, err := sub.SubConn(ctx, ienumvariant.EnumVariantSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumVariantSubConn = sub
	}

	o.enumVariant, err = ienumvariant.NewEnumVariantClient(ctx, enumVariantSubConn, append(opts, dcerpc.WithNoBind(enumVariantSubConn))...)

	typeCompSubConn, err := sub.SubConn(ctx, itypecomp.TypeCompSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		typeCompSubConn = sub
	}

	o.typeComp, err = itypecomp.NewTypeCompClient(ctx, typeCompSubConn, append(opts, dcerpc.WithNoBind(typeCompSubConn))...)

	typeInfoSubConn, err := sub.SubConn(ctx, itypeinfo.TypeInfoSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		typeInfoSubConn = sub
	}

	o.typeInfo, err = itypeinfo.NewTypeInfoClient(ctx, typeInfoSubConn, append(opts, dcerpc.WithNoBind(typeInfoSubConn))...)

	typeInfo2SubConn, err := sub.SubConn(ctx, itypeinfo2.TypeInfo2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		typeInfo2SubConn = sub
	}

	o.typeInfo2, err = itypeinfo2.NewTypeInfo2Client(ctx, typeInfo2SubConn, append(opts, dcerpc.WithNoBind(typeInfo2SubConn))...)

	typeLibSubConn, err := sub.SubConn(ctx, itypelib.TypeLibSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		typeLibSubConn = sub
	}

	o.typeLib, err = itypelib.NewTypeLibClient(ctx, typeLibSubConn, append(opts, dcerpc.WithNoBind(typeLibSubConn))...)

	typeLib2SubConn, err := sub.SubConn(ctx, itypelib2.TypeLib2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		typeLib2SubConn = sub
	}

	o.typeLib2, err = itypelib2.NewTypeLib2Client(ctx, typeLib2SubConn, append(opts, dcerpc.WithNoBind(typeLib2SubConn))...)
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
		dcomClient:  o.dcomClient.IPID(ctx, ipid),
		dispatch:    o.dispatch.IPID(ctx, ipid),
		enumVariant: o.enumVariant.IPID(ctx, ipid),
		typeComp:    o.typeComp.IPID(ctx, ipid),
		typeInfo:    o.typeInfo.IPID(ctx, ipid),
		typeInfo2:   o.typeInfo2.IPID(ctx, ipid),
		typeLib:     o.typeLib.IPID(ctx, ipid),
		typeLib2:    o.typeLib2.IPID(ctx, ipid),
		cc:          o.cc,
	}
}
