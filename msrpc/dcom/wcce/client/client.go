package wcce

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
	icertrequestd "github.com/oiweiwei/go-msrpc/msrpc/dcom/wcce/icertrequestd/v0"
	icertrequestd2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/wcce/icertrequestd2/v0"
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
	_ = icertrequestd.GoPackage
	_ = icertrequestd2.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wcce"
)

// dcom/wcce client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	CertRequestD() icertrequestd.CertRequestDClient
	CertRequestD2() icertrequestd2.CertRequestD2Client
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

	certRequestD  icertrequestd.CertRequestDClient
	certRequestD2 icertrequestd2.CertRequestD2Client
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) CertRequestD() icertrequestd.CertRequestDClient {
	return o.certRequestD
}

func (o *xxx_DefaultClient) CertRequestD2() icertrequestd2.CertRequestD2Client {
	return o.certRequestD2
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(icertrequestd.CertRequestDSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icertrequestd2.CertRequestD2SyntaxV0_0),
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

	certRequestDSubConn, err := sub.SubConn(ctx, icertrequestd.CertRequestDSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		certRequestDSubConn = sub
	}

	o.certRequestD, err = icertrequestd.NewCertRequestDClient(ctx, certRequestDSubConn, append(opts, dcerpc.WithNoBind(certRequestDSubConn))...)

	certRequestD2SubConn, err := sub.SubConn(ctx, icertrequestd2.CertRequestD2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		certRequestD2SubConn = sub
	}

	o.certRequestD2, err = icertrequestd2.NewCertRequestD2Client(ctx, certRequestD2SubConn, append(opts, dcerpc.WithNoBind(certRequestD2SubConn))...)
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
		dcomClient:    o.dcomClient.IPID(ctx, ipid),
		certRequestD:  o.certRequestD.IPID(ctx, ipid),
		certRequestD2: o.certRequestD2.IPID(ctx, ipid),
		cc:            o.cc,
	}
}
