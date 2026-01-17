package dmrp

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
	idmnotify "github.com/oiweiwei/go-msrpc/msrpc/dcom/dmrp/idmnotify/v0"
	idmremoteserver "github.com/oiweiwei/go-msrpc/msrpc/dcom/dmrp/idmremoteserver/v0"
	ivolumeclient "github.com/oiweiwei/go-msrpc/msrpc/dcom/dmrp/ivolumeclient/v0"
	ivolumeclient2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/dmrp/ivolumeclient2/v0"
	ivolumeclient3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/dmrp/ivolumeclient3/v0"
	ivolumeclient4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/dmrp/ivolumeclient4/v0"
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
	_ = ivolumeclient.GoPackage
	_ = idmnotify.GoPackage
	_ = idmremoteserver.GoPackage
	_ = ivolumeclient2.GoPackage
	_ = ivolumeclient3.GoPackage
	_ = ivolumeclient4.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/dmrp"
)

// dcom/dmrp client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	VolumeClient() ivolumeclient.VolumeClientClient
	IDMNotify() idmnotify.IDMNotifyClient
	IDMRemoteServer() idmremoteserver.IDMRemoteServerClient
	VolumeClient2() ivolumeclient2.VolumeClient2Client
	VolumeClient3() ivolumeclient3.VolumeClient3Client
	VolumeClient4() ivolumeclient4.VolumeClient4Client
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

	volumeClient    ivolumeclient.VolumeClientClient
	idmNotify       idmnotify.IDMNotifyClient
	idmRemoteServer idmremoteserver.IDMRemoteServerClient
	volumeClient2   ivolumeclient2.VolumeClient2Client
	volumeClient3   ivolumeclient3.VolumeClient3Client
	volumeClient4   ivolumeclient4.VolumeClient4Client
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) VolumeClient() ivolumeclient.VolumeClientClient {
	return o.volumeClient
}

func (o *xxx_DefaultClient) IDMNotify() idmnotify.IDMNotifyClient {
	return o.idmNotify
}

func (o *xxx_DefaultClient) IDMRemoteServer() idmremoteserver.IDMRemoteServerClient {
	return o.idmRemoteServer
}

func (o *xxx_DefaultClient) VolumeClient2() ivolumeclient2.VolumeClient2Client {
	return o.volumeClient2
}

func (o *xxx_DefaultClient) VolumeClient3() ivolumeclient3.VolumeClient3Client {
	return o.volumeClient3
}

func (o *xxx_DefaultClient) VolumeClient4() ivolumeclient4.VolumeClient4Client {
	return o.volumeClient4
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(ivolumeclient.VolumeClientSyntaxV0_0),
		dcerpc.WithAbstractSyntax(idmnotify.IDMNotifySyntaxV0_0),
		dcerpc.WithAbstractSyntax(idmremoteserver.IDMRemoteServerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivolumeclient2.VolumeClient2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivolumeclient3.VolumeClient3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ivolumeclient4.VolumeClient4SyntaxV0_0),
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

	volumeClientSubConn, err := sub.SubConn(ctx, ivolumeclient.VolumeClientSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumeClientSubConn = sub
	}

	o.volumeClient, err = ivolumeclient.NewVolumeClientClient(ctx, volumeClientSubConn, append(opts, dcerpc.WithNoBind(volumeClientSubConn))...)

	idmNotifySubConn, err := sub.SubConn(ctx, idmnotify.IDMNotifySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		idmNotifySubConn = sub
	}

	o.idmNotify, err = idmnotify.NewIDMNotifyClient(ctx, idmNotifySubConn, append(opts, dcerpc.WithNoBind(idmNotifySubConn))...)

	idmRemoteServerSubConn, err := sub.SubConn(ctx, idmremoteserver.IDMRemoteServerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		idmRemoteServerSubConn = sub
	}

	o.idmRemoteServer, err = idmremoteserver.NewIDMRemoteServerClient(ctx, idmRemoteServerSubConn, append(opts, dcerpc.WithNoBind(idmRemoteServerSubConn))...)

	volumeClient2SubConn, err := sub.SubConn(ctx, ivolumeclient2.VolumeClient2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumeClient2SubConn = sub
	}

	o.volumeClient2, err = ivolumeclient2.NewVolumeClient2Client(ctx, volumeClient2SubConn, append(opts, dcerpc.WithNoBind(volumeClient2SubConn))...)

	volumeClient3SubConn, err := sub.SubConn(ctx, ivolumeclient3.VolumeClient3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumeClient3SubConn = sub
	}

	o.volumeClient3, err = ivolumeclient3.NewVolumeClient3Client(ctx, volumeClient3SubConn, append(opts, dcerpc.WithNoBind(volumeClient3SubConn))...)

	volumeClient4SubConn, err := sub.SubConn(ctx, ivolumeclient4.VolumeClient4SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		volumeClient4SubConn = sub
	}

	o.volumeClient4, err = ivolumeclient4.NewVolumeClient4Client(ctx, volumeClient4SubConn, append(opts, dcerpc.WithNoBind(volumeClient4SubConn))...)
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
		dcomClient:      o.dcomClient.IPID(ctx, ipid),
		volumeClient:    o.volumeClient.IPID(ctx, ipid),
		idmNotify:       o.idmNotify.IPID(ctx, ipid),
		idmRemoteServer: o.idmRemoteServer.IPID(ctx, ipid),
		volumeClient2:   o.volumeClient2.IPID(ctx, ipid),
		volumeClient3:   o.volumeClient3.IPID(ctx, ipid),
		volumeClient4:   o.volumeClient4.IPID(ctx, ipid),
		cc:              o.cc,
	}
}
