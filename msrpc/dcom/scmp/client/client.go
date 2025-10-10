package scmp

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
	ivssdifferentialsoftwaresnapshotmgmt "github.com/oiweiwei/go-msrpc/msrpc/dcom/scmp/ivssdifferentialsoftwaresnapshotmgmt/v0"
	ivssenummgmtobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/scmp/ivssenummgmtobject/v0"
	ivssenumobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/scmp/ivssenumobject/v0"
	ivsssnapshotmgmt "github.com/oiweiwei/go-msrpc/msrpc/dcom/scmp/ivsssnapshotmgmt/v0"
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
	_ = ivsssnapshotmgmt.GoPackage
	_ = ivssdifferentialsoftwaresnapshotmgmt.GoPackage
	_ = ivssenumobject.GoPackage
	_ = ivssenummgmtobject.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/scmp"
)

// dcom/scmp client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	VSSSnapshotManagement() ivsssnapshotmgmt.VSSSnapshotManagementClient
	VSSDifferentialSoftwareSnapshotManagement() ivssdifferentialsoftwaresnapshotmgmt.VSSDifferentialSoftwareSnapshotManagementClient
	VSSEnumObject() ivssenumobject.VSSEnumObjectClient
	VSSEnumManagementObject() ivssenummgmtobject.VSSEnumManagementObjectClient
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

	vssSnapshotManagement                     ivsssnapshotmgmt.VSSSnapshotManagementClient
	vssDifferentialSoftwareSnapshotManagement ivssdifferentialsoftwaresnapshotmgmt.VSSDifferentialSoftwareSnapshotManagementClient
	vssEnumObject                             ivssenumobject.VSSEnumObjectClient
	vssEnumManagementObject                   ivssenummgmtobject.VSSEnumManagementObjectClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) VSSSnapshotManagement() ivsssnapshotmgmt.VSSSnapshotManagementClient {
	return o.vssSnapshotManagement
}

func (o *xxx_DefaultClient) VSSDifferentialSoftwareSnapshotManagement() ivssdifferentialsoftwaresnapshotmgmt.VSSDifferentialSoftwareSnapshotManagementClient {
	return o.vssDifferentialSoftwareSnapshotManagement
}

func (o *xxx_DefaultClient) VSSEnumObject() ivssenumobject.VSSEnumObjectClient {
	return o.vssEnumObject
}

func (o *xxx_DefaultClient) VSSEnumManagementObject() ivssenummgmtobject.VSSEnumManagementObjectClient {
	return o.vssEnumManagementObject
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(ivsssnapshotmgmt.VSSSnapshotManagementSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivssdifferentialsoftwaresnapshotmgmt.VSSDifferentialSoftwareSnapshotManagementSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivssenumobject.VSSEnumObjectSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivssenummgmtobject.VSSEnumManagementObjectSyntaxV0_0),
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

	vssSnapshotManagementSubConn, err := sub.SubConn(ctx, ivsssnapshotmgmt.VSSSnapshotManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		vssSnapshotManagementSubConn = sub
	}

	o.vssSnapshotManagement, err = ivsssnapshotmgmt.NewVSSSnapshotManagementClient(ctx, vssSnapshotManagementSubConn, append(opts, dcerpc.WithNoBind(vssSnapshotManagementSubConn))...)

	vssDifferentialSoftwareSnapshotManagementSubConn, err := sub.SubConn(ctx, ivssdifferentialsoftwaresnapshotmgmt.VSSDifferentialSoftwareSnapshotManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		vssDifferentialSoftwareSnapshotManagementSubConn = sub
	}

	o.vssDifferentialSoftwareSnapshotManagement, err = ivssdifferentialsoftwaresnapshotmgmt.NewVSSDifferentialSoftwareSnapshotManagementClient(ctx, vssDifferentialSoftwareSnapshotManagementSubConn, append(opts, dcerpc.WithNoBind(vssDifferentialSoftwareSnapshotManagementSubConn))...)

	vssEnumObjectSubConn, err := sub.SubConn(ctx, ivssenumobject.VSSEnumObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		vssEnumObjectSubConn = sub
	}

	o.vssEnumObject, err = ivssenumobject.NewVSSEnumObjectClient(ctx, vssEnumObjectSubConn, append(opts, dcerpc.WithNoBind(vssEnumObjectSubConn))...)

	vssEnumManagementObjectSubConn, err := sub.SubConn(ctx, ivssenummgmtobject.VSSEnumManagementObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		vssEnumManagementObjectSubConn = sub
	}

	o.vssEnumManagementObject, err = ivssenummgmtobject.NewVSSEnumManagementObjectClient(ctx, vssEnumManagementObjectSubConn, append(opts, dcerpc.WithNoBind(vssEnumManagementObjectSubConn))...)
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
		vssSnapshotManagement: o.vssSnapshotManagement.IPID(ctx, ipid),
		vssDifferentialSoftwareSnapshotManagement: o.vssDifferentialSoftwareSnapshotManagement.IPID(ctx, ipid),
		vssEnumObject:           o.vssEnumObject.IPID(ctx, ipid),
		vssEnumManagementObject: o.vssEnumManagementObject.IPID(ctx, ipid),
		cc:                      o.cc,
	}
}
