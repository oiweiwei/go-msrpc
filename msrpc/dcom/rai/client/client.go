package rai

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
	ipchcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/rai/ipchcollection/v0"
	ipchservice "github.com/oiweiwei/go-msrpc/msrpc/dcom/rai/ipchservice/v0"
	irasrv "github.com/oiweiwei/go-msrpc/msrpc/dcom/rai/irasrv/v0"
	isafsession "github.com/oiweiwei/go-msrpc/msrpc/dcom/rai/isafsession/v0"
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
	_ = ipchcollection.GoPackage
	_ = ipchservice.GoPackage
	_ = isafsession.GoPackage
	_ = irasrv.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rai"
)

// dcom/rai client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	PCHCollection() ipchcollection.PCHCollectionClient
	PCHService() ipchservice.PCHServiceClient
	SAFSession() isafsession.SAFSessionClient
	RemoteAssistanceServer() irasrv.RemoteAssistanceServerClient
	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Client
}
type xxx_DefaultClient struct {
	cc dcerpc.Conn

	dcomClient dcom_client.Client

	pchCollection          ipchcollection.PCHCollectionClient
	pchService             ipchservice.PCHServiceClient
	safSession             isafsession.SAFSessionClient
	remoteAssistanceServer irasrv.RemoteAssistanceServerClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) PCHCollection() ipchcollection.PCHCollectionClient {
	return o.pchCollection
}

func (o *xxx_DefaultClient) PCHService() ipchservice.PCHServiceClient {
	return o.pchService
}

func (o *xxx_DefaultClient) SAFSession() isafsession.SAFSessionClient {
	return o.safSession
}

func (o *xxx_DefaultClient) RemoteAssistanceServer() irasrv.RemoteAssistanceServerClient {
	return o.remoteAssistanceServer
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(ipchcollection.PCHCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ipchservice.PCHServiceSyntaxV0_0),
		dcerpc.WithAbstractSyntax(isafsession.SAFSessionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(irasrv.RemoteAssistanceServerSyntaxV0_0),
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

	pchCollectionSubConn, err := sub.SubConn(ctx, ipchcollection.PCHCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		pchCollectionSubConn = sub
	}

	o.pchCollection, err = ipchcollection.NewPCHCollectionClient(ctx, pchCollectionSubConn, append(opts, dcerpc.WithNoBind(pchCollectionSubConn))...)

	pchServiceSubConn, err := sub.SubConn(ctx, ipchservice.PCHServiceSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		pchServiceSubConn = sub
	}

	o.pchService, err = ipchservice.NewPCHServiceClient(ctx, pchServiceSubConn, append(opts, dcerpc.WithNoBind(pchServiceSubConn))...)

	safSessionSubConn, err := sub.SubConn(ctx, isafsession.SAFSessionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		safSessionSubConn = sub
	}

	o.safSession, err = isafsession.NewSAFSessionClient(ctx, safSessionSubConn, append(opts, dcerpc.WithNoBind(safSessionSubConn))...)

	remoteAssistanceServerSubConn, err := sub.SubConn(ctx, irasrv.RemoteAssistanceServerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		remoteAssistanceServerSubConn = sub
	}

	o.remoteAssistanceServer, err = irasrv.NewRemoteAssistanceServerClient(ctx, remoteAssistanceServerSubConn, append(opts, dcerpc.WithNoBind(remoteAssistanceServerSubConn))...)
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
		dcomClient:             o.dcomClient.IPID(ctx, ipid),
		pchCollection:          o.pchCollection.IPID(ctx, ipid),
		pchService:             o.pchService.IPID(ctx, ipid),
		safSession:             o.safSession.IPID(ctx, ipid),
		remoteAssistanceServer: o.remoteAssistanceServer.IPID(ctx, ipid),
		cc:                     o.cc,
	}
}
