package csra

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
	icertadmind "github.com/oiweiwei/go-msrpc/msrpc/dcom/csra/icertadmind/v0"
	icertadmind2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/csra/icertadmind2/v0"
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
	_ = icertadmind.GoPackage
	_ = icertadmind2.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csra"
)

// dcom/csra client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	CertAdminD() icertadmind.CertAdminDClient
	CertAdminD2() icertadmind2.CertAdminD2Client
	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Client
}
type xxx_DefaultClient struct {
	cc dcerpc.Conn

	dcomClient dcom_client.Client

	certAdminD  icertadmind.CertAdminDClient
	certAdminD2 icertadmind2.CertAdminD2Client
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) CertAdminD() icertadmind.CertAdminDClient {
	return o.certAdminD
}

func (o *xxx_DefaultClient) CertAdminD2() icertadmind2.CertAdminD2Client {
	return o.certAdminD2
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(icertadmind.CertAdminDSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icertadmind2.CertAdminD2SyntaxV0_0),
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

	certAdminDSubConn, err := sub.SubConn(ctx, icertadmind.CertAdminDSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		certAdminDSubConn = sub
	}

	o.certAdminD, err = icertadmind.NewCertAdminDClient(ctx, certAdminDSubConn, append(opts, dcerpc.WithNoBind(certAdminDSubConn))...)

	certAdminD2SubConn, err := sub.SubConn(ctx, icertadmind2.CertAdminD2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		certAdminD2SubConn = sub
	}

	o.certAdminD2, err = icertadmind2.NewCertAdminD2Client(ctx, certAdminD2SubConn, append(opts, dcerpc.WithNoBind(certAdminD2SubConn))...)
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
		certAdminD:  o.certAdminD.IPID(ctx, ipid),
		certAdminD2: o.certAdminD2.IPID(ctx, ipid),
		cc:          o.cc,
	}
}
