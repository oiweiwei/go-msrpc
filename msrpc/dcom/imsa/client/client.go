package imsa

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
	iiisapplicationadmin "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/iiisapplicationadmin/v0"
	iiiscertobj "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/iiiscertobj/v0"
	imsadminbase2w "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/imsadminbase2w/v0"
	imsadminbase3w "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/imsadminbase3w/v0"
	imsadminbasew "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/imsadminbasew/v0"
	iwamadmin "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/iwamadmin/v0"
	iwamadmin2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/imsa/iwamadmin2/v0"
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
	_ = imsadminbasew.GoPackage
	_ = imsadminbase2w.GoPackage
	_ = imsadminbase3w.GoPackage
	_ = iwamadmin.GoPackage
	_ = iwamadmin2.GoPackage
	_ = iiisapplicationadmin.GoPackage
	_ = iiiscertobj.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/imsa"
)

// dcom/imsa client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	IMSAdminBaseW() imsadminbasew.IMSAdminBaseWClient
	IMSAdminBase2W() imsadminbase2w.IMSAdminBase2WClient
	IMSAdminBase3W() imsadminbase3w.IMSAdminBase3WClient
	WAMAdmin() iwamadmin.WAMAdminClient
	WAMAdmin2() iwamadmin2.WAMAdmin2Client
	IISApplicationAdmin() iiisapplicationadmin.IISApplicationAdminClient
	IISCertObject() iiiscertobj.IISCertObjectClient
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

	imsAdminBaseW       imsadminbasew.IMSAdminBaseWClient
	imsAdminBase2W      imsadminbase2w.IMSAdminBase2WClient
	imsAdminBase3W      imsadminbase3w.IMSAdminBase3WClient
	wamAdmin            iwamadmin.WAMAdminClient
	wamAdmin2           iwamadmin2.WAMAdmin2Client
	iisApplicationAdmin iiisapplicationadmin.IISApplicationAdminClient
	iisCertObject       iiiscertobj.IISCertObjectClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) IMSAdminBaseW() imsadminbasew.IMSAdminBaseWClient {
	return o.imsAdminBaseW
}

func (o *xxx_DefaultClient) IMSAdminBase2W() imsadminbase2w.IMSAdminBase2WClient {
	return o.imsAdminBase2W
}

func (o *xxx_DefaultClient) IMSAdminBase3W() imsadminbase3w.IMSAdminBase3WClient {
	return o.imsAdminBase3W
}

func (o *xxx_DefaultClient) WAMAdmin() iwamadmin.WAMAdminClient {
	return o.wamAdmin
}

func (o *xxx_DefaultClient) WAMAdmin2() iwamadmin2.WAMAdmin2Client {
	return o.wamAdmin2
}

func (o *xxx_DefaultClient) IISApplicationAdmin() iiisapplicationadmin.IISApplicationAdminClient {
	return o.iisApplicationAdmin
}

func (o *xxx_DefaultClient) IISCertObject() iiiscertobj.IISCertObjectClient {
	return o.iisCertObject
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(imsadminbasew.IMSAdminBaseWSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsadminbase2w.IMSAdminBase2WSyntaxV0_0),
		dcerpc.WithAbstractSyntax(imsadminbase3w.IMSAdminBase3WSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwamadmin.WAMAdminSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iwamadmin2.WAMAdmin2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iiisapplicationadmin.IISApplicationAdminSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iiiscertobj.IISCertObjectSyntaxV0_0),
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

	imsAdminBaseWSubConn, err := sub.SubConn(ctx, imsadminbasew.IMSAdminBaseWSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsAdminBaseWSubConn = sub
	}

	o.imsAdminBaseW, err = imsadminbasew.NewIMSAdminBaseWClient(ctx, imsAdminBaseWSubConn, append(opts, dcerpc.WithNoBind(imsAdminBaseWSubConn))...)

	imsAdminBase2WSubConn, err := sub.SubConn(ctx, imsadminbase2w.IMSAdminBase2WSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsAdminBase2WSubConn = sub
	}

	o.imsAdminBase2W, err = imsadminbase2w.NewIMSAdminBase2WClient(ctx, imsAdminBase2WSubConn, append(opts, dcerpc.WithNoBind(imsAdminBase2WSubConn))...)

	imsAdminBase3WSubConn, err := sub.SubConn(ctx, imsadminbase3w.IMSAdminBase3WSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		imsAdminBase3WSubConn = sub
	}

	o.imsAdminBase3W, err = imsadminbase3w.NewIMSAdminBase3WClient(ctx, imsAdminBase3WSubConn, append(opts, dcerpc.WithNoBind(imsAdminBase3WSubConn))...)

	wamAdminSubConn, err := sub.SubConn(ctx, iwamadmin.WAMAdminSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		wamAdminSubConn = sub
	}

	o.wamAdmin, err = iwamadmin.NewWAMAdminClient(ctx, wamAdminSubConn, append(opts, dcerpc.WithNoBind(wamAdminSubConn))...)

	wamAdmin2SubConn, err := sub.SubConn(ctx, iwamadmin2.WAMAdmin2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		wamAdmin2SubConn = sub
	}

	o.wamAdmin2, err = iwamadmin2.NewWAMAdmin2Client(ctx, wamAdmin2SubConn, append(opts, dcerpc.WithNoBind(wamAdmin2SubConn))...)

	iisApplicationAdminSubConn, err := sub.SubConn(ctx, iiisapplicationadmin.IISApplicationAdminSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iisApplicationAdminSubConn = sub
	}

	o.iisApplicationAdmin, err = iiisapplicationadmin.NewIISApplicationAdminClient(ctx, iisApplicationAdminSubConn, append(opts, dcerpc.WithNoBind(iisApplicationAdminSubConn))...)

	iisCertObjectSubConn, err := sub.SubConn(ctx, iiiscertobj.IISCertObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		iisCertObjectSubConn = sub
	}

	o.iisCertObject, err = iiiscertobj.NewIISCertObjectClient(ctx, iisCertObjectSubConn, append(opts, dcerpc.WithNoBind(iisCertObjectSubConn))...)
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
		dcomClient:          o.dcomClient.IPID(ctx, ipid),
		imsAdminBaseW:       o.imsAdminBaseW.IPID(ctx, ipid),
		imsAdminBase2W:      o.imsAdminBase2W.IPID(ctx, ipid),
		imsAdminBase3W:      o.imsAdminBase3W.IPID(ctx, ipid),
		wamAdmin:            o.wamAdmin.IPID(ctx, ipid),
		wamAdmin2:           o.wamAdmin2.IPID(ctx, ipid),
		iisApplicationAdmin: o.iisApplicationAdmin.IPID(ctx, ipid),
		iisCertObject:       o.iisCertObject.IPID(ctx, ipid),
		cc:                  o.cc,
	}
}
