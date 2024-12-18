package adtg

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	idatafactory "github.com/oiweiwei/go-msrpc/msrpc/dcom/adtg/idatafactory/v0"
	idatafactory2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/adtg/idatafactory2/v0"
	idatafactory3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/adtg/idatafactory3/v0"
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
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
	_ = idatafactory.GoPackage
	_ = idatafactory2.GoPackage
	_ = idatafactory3.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/adtg"
)

// dcom/adtg client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	DataFactory() idatafactory.DataFactoryClient
	DataFactory2() idatafactory2.DataFactory2Client
	DataFactory3() idatafactory3.DataFactory3Client
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

	dataFactory  idatafactory.DataFactoryClient
	dataFactory2 idatafactory2.DataFactory2Client
	dataFactory3 idatafactory3.DataFactory3Client
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) DataFactory() idatafactory.DataFactoryClient {
	return o.dataFactory
}

func (o *xxx_DefaultClient) DataFactory2() idatafactory2.DataFactory2Client {
	return o.dataFactory2
}

func (o *xxx_DefaultClient) DataFactory3() idatafactory3.DataFactory3Client {
	return o.dataFactory3
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(idatafactory.DataFactorySyntaxV0_0),
		dcerpc.WithAbstractSyntax(idatafactory2.DataFactory2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(idatafactory3.DataFactory3SyntaxV0_0),
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

	dataFactorySubConn, err := sub.SubConn(ctx, idatafactory.DataFactorySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		dataFactorySubConn = sub
	}

	o.dataFactory, err = idatafactory.NewDataFactoryClient(ctx, dataFactorySubConn, append(opts, dcerpc.WithNoBind(dataFactorySubConn))...)

	dataFactory2SubConn, err := sub.SubConn(ctx, idatafactory2.DataFactory2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		dataFactory2SubConn = sub
	}

	o.dataFactory2, err = idatafactory2.NewDataFactory2Client(ctx, dataFactory2SubConn, append(opts, dcerpc.WithNoBind(dataFactory2SubConn))...)

	dataFactory3SubConn, err := sub.SubConn(ctx, idatafactory3.DataFactory3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		dataFactory3SubConn = sub
	}

	o.dataFactory3, err = idatafactory3.NewDataFactory3Client(ctx, dataFactory3SubConn, append(opts, dcerpc.WithNoBind(dataFactory3SubConn))...)
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
		dcomClient:   o.dcomClient.IPID(ctx, ipid),
		dataFactory:  o.dataFactory.IPID(ctx, ipid),
		dataFactory2: o.dataFactory2.IPID(ctx, ipid),
		dataFactory3: o.dataFactory3.IPID(ctx, ipid),
		cc:           o.cc,
	}
}
