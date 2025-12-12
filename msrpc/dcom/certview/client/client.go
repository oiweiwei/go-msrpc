package certview

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	icertview "github.com/oiweiwei/go-msrpc/msrpc/dcom/certview/icertview/v0"
	icertview2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/certview/icertview2/v0"
	ienumcertviewattribute "github.com/oiweiwei/go-msrpc/msrpc/dcom/certview/ienumcertviewattribute/v0"
	ienumcertviewcolumn "github.com/oiweiwei/go-msrpc/msrpc/dcom/certview/ienumcertviewcolumn/v0"
	ienumcertviewextension "github.com/oiweiwei/go-msrpc/msrpc/dcom/certview/ienumcertviewextension/v0"
	ienumcertviewrow "github.com/oiweiwei/go-msrpc/msrpc/dcom/certview/ienumcertviewrow/v0"
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
	_ = ienumcertviewcolumn.GoPackage
	_ = ienumcertviewattribute.GoPackage
	_ = ienumcertviewextension.GoPackage
	_ = ienumcertviewrow.GoPackage
	_ = icertview.GoPackage
	_ = icertview2.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/certview"
)

// dcom/certview client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	EnumCertViewColumn() ienumcertviewcolumn.EnumCertViewColumnClient
	EnumCertViewAttribute() ienumcertviewattribute.EnumCertViewAttributeClient
	EnumCertViewExtension() ienumcertviewextension.EnumCertViewExtensionClient
	EnumCertViewRow() ienumcertviewrow.EnumCertViewRowClient
	CertView() icertview.CertViewClient
	CertView2() icertview2.CertView2Client
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

	enumCertViewColumn    ienumcertviewcolumn.EnumCertViewColumnClient
	enumCertViewAttribute ienumcertviewattribute.EnumCertViewAttributeClient
	enumCertViewExtension ienumcertviewextension.EnumCertViewExtensionClient
	enumCertViewRow       ienumcertviewrow.EnumCertViewRowClient
	certView              icertview.CertViewClient
	certView2             icertview2.CertView2Client
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) EnumCertViewColumn() ienumcertviewcolumn.EnumCertViewColumnClient {
	return o.enumCertViewColumn
}

func (o *xxx_DefaultClient) EnumCertViewAttribute() ienumcertviewattribute.EnumCertViewAttributeClient {
	return o.enumCertViewAttribute
}

func (o *xxx_DefaultClient) EnumCertViewExtension() ienumcertviewextension.EnumCertViewExtensionClient {
	return o.enumCertViewExtension
}

func (o *xxx_DefaultClient) EnumCertViewRow() ienumcertviewrow.EnumCertViewRowClient {
	return o.enumCertViewRow
}

func (o *xxx_DefaultClient) CertView() icertview.CertViewClient {
	return o.certView
}

func (o *xxx_DefaultClient) CertView2() icertview2.CertView2Client {
	return o.certView2
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(ienumcertviewcolumn.EnumCertViewColumnSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumcertviewattribute.EnumCertViewAttributeSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumcertviewextension.EnumCertViewExtensionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumcertviewrow.EnumCertViewRowSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icertview.CertViewSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icertview2.CertView2SyntaxV0_0),
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

	enumCertViewColumnSubConn, err := sub.SubConn(ctx, ienumcertviewcolumn.EnumCertViewColumnSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumCertViewColumnSubConn = sub
	}

	o.enumCertViewColumn, err = ienumcertviewcolumn.NewEnumCertViewColumnClient(ctx, enumCertViewColumnSubConn, append(opts, dcerpc.WithNoBind(enumCertViewColumnSubConn))...)

	enumCertViewAttributeSubConn, err := sub.SubConn(ctx, ienumcertviewattribute.EnumCertViewAttributeSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumCertViewAttributeSubConn = sub
	}

	o.enumCertViewAttribute, err = ienumcertviewattribute.NewEnumCertViewAttributeClient(ctx, enumCertViewAttributeSubConn, append(opts, dcerpc.WithNoBind(enumCertViewAttributeSubConn))...)

	enumCertViewExtensionSubConn, err := sub.SubConn(ctx, ienumcertviewextension.EnumCertViewExtensionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumCertViewExtensionSubConn = sub
	}

	o.enumCertViewExtension, err = ienumcertviewextension.NewEnumCertViewExtensionClient(ctx, enumCertViewExtensionSubConn, append(opts, dcerpc.WithNoBind(enumCertViewExtensionSubConn))...)

	enumCertViewRowSubConn, err := sub.SubConn(ctx, ienumcertviewrow.EnumCertViewRowSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumCertViewRowSubConn = sub
	}

	o.enumCertViewRow, err = ienumcertviewrow.NewEnumCertViewRowClient(ctx, enumCertViewRowSubConn, append(opts, dcerpc.WithNoBind(enumCertViewRowSubConn))...)

	certViewSubConn, err := sub.SubConn(ctx, icertview.CertViewSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		certViewSubConn = sub
	}

	o.certView, err = icertview.NewCertViewClient(ctx, certViewSubConn, append(opts, dcerpc.WithNoBind(certViewSubConn))...)

	certView2SubConn, err := sub.SubConn(ctx, icertview2.CertView2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		certView2SubConn = sub
	}

	o.certView2, err = icertview2.NewCertView2Client(ctx, certView2SubConn, append(opts, dcerpc.WithNoBind(certView2SubConn))...)
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
		enumCertViewColumn:    o.enumCertViewColumn.IPID(ctx, ipid),
		enumCertViewAttribute: o.enumCertViewAttribute.IPID(ctx, ipid),
		enumCertViewExtension: o.enumCertViewExtension.IPID(ctx, ipid),
		enumCertViewRow:       o.enumCertViewRow.IPID(ctx, ipid),
		certView:              o.certView.IPID(ctx, ipid),
		certView2:             o.certView2.IPID(ctx, ipid),
		cc:                    o.cc,
	}
}
