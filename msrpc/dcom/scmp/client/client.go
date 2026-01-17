package scmp

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	SnapshotManagement() ivsssnapshotmgmt.SnapshotManagementClient
	DifferentialSoftwareSnapshotManagement() ivssdifferentialsoftwaresnapshotmgmt.DifferentialSoftwareSnapshotManagementClient
	EnumObject() ivssenumobject.EnumObjectClient
	EnumManagementObject() ivssenummgmtobject.EnumManagementObjectClient
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

	snapshotManagement                     ivsssnapshotmgmt.SnapshotManagementClient
	differentialSoftwareSnapshotManagement ivssdifferentialsoftwaresnapshotmgmt.DifferentialSoftwareSnapshotManagementClient
	enumObject                             ivssenumobject.EnumObjectClient
	enumManagementObject                   ivssenummgmtobject.EnumManagementObjectClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) SnapshotManagement() ivsssnapshotmgmt.SnapshotManagementClient {
	return o.snapshotManagement
}

func (o *xxx_DefaultClient) DifferentialSoftwareSnapshotManagement() ivssdifferentialsoftwaresnapshotmgmt.DifferentialSoftwareSnapshotManagementClient {
	return o.differentialSoftwareSnapshotManagement
}

func (o *xxx_DefaultClient) EnumObject() ivssenumobject.EnumObjectClient {
	return o.enumObject
}

func (o *xxx_DefaultClient) EnumManagementObject() ivssenummgmtobject.EnumManagementObjectClient {
	return o.enumManagementObject
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(ivsssnapshotmgmt.SnapshotManagementSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivssdifferentialsoftwaresnapshotmgmt.DifferentialSoftwareSnapshotManagementSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivssenumobject.EnumObjectSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivssenummgmtobject.EnumManagementObjectSyntaxV0_0),
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

	snapshotManagementSubConn, err := sub.SubConn(ctx, ivsssnapshotmgmt.SnapshotManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		snapshotManagementSubConn = sub
	}

	o.snapshotManagement, err = ivsssnapshotmgmt.NewSnapshotManagementClient(ctx, snapshotManagementSubConn, append(opts, dcerpc.WithNoBind(snapshotManagementSubConn))...)

	differentialSoftwareSnapshotManagementSubConn, err := sub.SubConn(ctx, ivssdifferentialsoftwaresnapshotmgmt.DifferentialSoftwareSnapshotManagementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		differentialSoftwareSnapshotManagementSubConn = sub
	}

	o.differentialSoftwareSnapshotManagement, err = ivssdifferentialsoftwaresnapshotmgmt.NewDifferentialSoftwareSnapshotManagementClient(ctx, differentialSoftwareSnapshotManagementSubConn, append(opts, dcerpc.WithNoBind(differentialSoftwareSnapshotManagementSubConn))...)

	enumObjectSubConn, err := sub.SubConn(ctx, ivssenumobject.EnumObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumObjectSubConn = sub
	}

	o.enumObject, err = ivssenumobject.NewEnumObjectClient(ctx, enumObjectSubConn, append(opts, dcerpc.WithNoBind(enumObjectSubConn))...)

	enumManagementObjectSubConn, err := sub.SubConn(ctx, ivssenummgmtobject.EnumManagementObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumManagementObjectSubConn = sub
	}

	o.enumManagementObject, err = ivssenummgmtobject.NewEnumManagementObjectClient(ctx, enumManagementObjectSubConn, append(opts, dcerpc.WithNoBind(enumManagementObjectSubConn))...)
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
		dcomClient:                             o.dcomClient.IPID(ctx, ipid),
		snapshotManagement:                     o.snapshotManagement.IPID(ctx, ipid),
		differentialSoftwareSnapshotManagement: o.differentialSoftwareSnapshotManagement.IPID(ctx, ipid),
		enumObject:                             o.enumObject.IPID(ctx, ipid),
		enumManagementObject:                   o.enumManagementObject.IPID(ctx, ipid),
		cc:                                     o.cc,
	}
}
