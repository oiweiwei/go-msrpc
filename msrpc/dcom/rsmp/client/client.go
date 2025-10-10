package rsmp

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
	iclientsink "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/iclientsink/v0"
	imessenger "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/imessenger/v0"
	intmslibrarycontrol1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmslibrarycontrol1/v0"
	intmslibrarycontrol2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmslibrarycontrol2/v0"
	intmsmediaservices1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsmediaservices1/v0"
	intmsnotifysink "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsnotifysink/v0"
	intmsobjectinfo1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsobjectinfo1/v0"
	intmsobjectmanagement1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsobjectmanagement1/v0"
	intmsobjectmanagement2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsobjectmanagement2/v0"
	intmsobjectmanagement3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsobjectmanagement3/v0"
	intmssession1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmssession1/v0"
	irobustntmsmediaservices1 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/irobustntmsmediaservices1/v0"
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
	_ = intmslibrarycontrol1.GoPackage
	_ = intmsmediaservices1.GoPackage
	_ = intmsobjectinfo1.GoPackage
	_ = intmsobjectmanagement1.GoPackage
	_ = intmssession1.GoPackage
	_ = iclientsink.GoPackage
	_ = intmslibrarycontrol2.GoPackage
	_ = intmsobjectmanagement2.GoPackage
	_ = intmsobjectmanagement3.GoPackage
	_ = irobustntmsmediaservices1.GoPackage
	_ = imessenger.GoPackage
	_ = intmsnotifysink.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

// dcom/rsmp client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	LibraryControl1() intmslibrarycontrol1.LibraryControl1Client
	MediaServices1() intmsmediaservices1.MediaServices1Client
	ObjectInfo1() intmsobjectinfo1.ObjectInfo1Client
	ObjectManagement1() intmsobjectmanagement1.ObjectManagement1Client
	Session1() intmssession1.Session1Client
	ClientSink() iclientsink.ClientSinkClient
	LibraryControl2() intmslibrarycontrol2.LibraryControl2Client
	ObjectManagement2() intmsobjectmanagement2.ObjectManagement2Client
	ObjectManagement3() intmsobjectmanagement3.ObjectManagement3Client
	RobustNTMSMediaServices1() irobustntmsmediaservices1.RobustNTMSMediaServices1Client
	Messenger() imessenger.MessengerClient
	NotifySink() intmsnotifysink.NotifySinkClient
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

	libraryControl1          intmslibrarycontrol1.LibraryControl1Client
	mediaServices1           intmsmediaservices1.MediaServices1Client
	objectInfo1              intmsobjectinfo1.ObjectInfo1Client
	objectManagement1        intmsobjectmanagement1.ObjectManagement1Client
	session1                 intmssession1.Session1Client
	clientSink               iclientsink.ClientSinkClient
	libraryControl2          intmslibrarycontrol2.LibraryControl2Client
	objectManagement2        intmsobjectmanagement2.ObjectManagement2Client
	objectManagement3        intmsobjectmanagement3.ObjectManagement3Client
	robustNTMSMediaServices1 irobustntmsmediaservices1.RobustNTMSMediaServices1Client
	messenger                imessenger.MessengerClient
	notifySink               intmsnotifysink.NotifySinkClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) LibraryControl1() intmslibrarycontrol1.LibraryControl1Client {
	return o.libraryControl1
}

func (o *xxx_DefaultClient) MediaServices1() intmsmediaservices1.MediaServices1Client {
	return o.mediaServices1
}

func (o *xxx_DefaultClient) ObjectInfo1() intmsobjectinfo1.ObjectInfo1Client {
	return o.objectInfo1
}

func (o *xxx_DefaultClient) ObjectManagement1() intmsobjectmanagement1.ObjectManagement1Client {
	return o.objectManagement1
}

func (o *xxx_DefaultClient) Session1() intmssession1.Session1Client {
	return o.session1
}

func (o *xxx_DefaultClient) ClientSink() iclientsink.ClientSinkClient {
	return o.clientSink
}

func (o *xxx_DefaultClient) LibraryControl2() intmslibrarycontrol2.LibraryControl2Client {
	return o.libraryControl2
}

func (o *xxx_DefaultClient) ObjectManagement2() intmsobjectmanagement2.ObjectManagement2Client {
	return o.objectManagement2
}

func (o *xxx_DefaultClient) ObjectManagement3() intmsobjectmanagement3.ObjectManagement3Client {
	return o.objectManagement3
}

func (o *xxx_DefaultClient) RobustNTMSMediaServices1() irobustntmsmediaservices1.RobustNTMSMediaServices1Client {
	return o.robustNTMSMediaServices1
}

func (o *xxx_DefaultClient) Messenger() imessenger.MessengerClient {
	return o.messenger
}

func (o *xxx_DefaultClient) NotifySink() intmsnotifysink.NotifySinkClient {
	return o.notifySink
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(intmslibrarycontrol1.LibraryControl1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsmediaservices1.MediaServices1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsobjectinfo1.ObjectInfo1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsobjectmanagement1.ObjectManagement1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmssession1.Session1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iclientsink.ClientSinkSyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmslibrarycontrol2.LibraryControl2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsobjectmanagement2.ObjectManagement2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsobjectmanagement3.ObjectManagement3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(irobustntmsmediaservices1.RobustNTMSMediaServices1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imessenger.MessengerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsnotifysink.NotifySinkSyntaxV0_0),
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

	libraryControl1SubConn, err := sub.SubConn(ctx, intmslibrarycontrol1.LibraryControl1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		libraryControl1SubConn = sub
	}

	o.libraryControl1, err = intmslibrarycontrol1.NewLibraryControl1Client(ctx, libraryControl1SubConn, append(opts, dcerpc.WithNoBind(libraryControl1SubConn))...)

	mediaServices1SubConn, err := sub.SubConn(ctx, intmsmediaservices1.MediaServices1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		mediaServices1SubConn = sub
	}

	o.mediaServices1, err = intmsmediaservices1.NewMediaServices1Client(ctx, mediaServices1SubConn, append(opts, dcerpc.WithNoBind(mediaServices1SubConn))...)

	objectInfo1SubConn, err := sub.SubConn(ctx, intmsobjectinfo1.ObjectInfo1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		objectInfo1SubConn = sub
	}

	o.objectInfo1, err = intmsobjectinfo1.NewObjectInfo1Client(ctx, objectInfo1SubConn, append(opts, dcerpc.WithNoBind(objectInfo1SubConn))...)

	objectManagement1SubConn, err := sub.SubConn(ctx, intmsobjectmanagement1.ObjectManagement1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		objectManagement1SubConn = sub
	}

	o.objectManagement1, err = intmsobjectmanagement1.NewObjectManagement1Client(ctx, objectManagement1SubConn, append(opts, dcerpc.WithNoBind(objectManagement1SubConn))...)

	session1SubConn, err := sub.SubConn(ctx, intmssession1.Session1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		session1SubConn = sub
	}

	o.session1, err = intmssession1.NewSession1Client(ctx, session1SubConn, append(opts, dcerpc.WithNoBind(session1SubConn))...)

	clientSinkSubConn, err := sub.SubConn(ctx, iclientsink.ClientSinkSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clientSinkSubConn = sub
	}

	o.clientSink, err = iclientsink.NewClientSinkClient(ctx, clientSinkSubConn, append(opts, dcerpc.WithNoBind(clientSinkSubConn))...)

	libraryControl2SubConn, err := sub.SubConn(ctx, intmslibrarycontrol2.LibraryControl2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		libraryControl2SubConn = sub
	}

	o.libraryControl2, err = intmslibrarycontrol2.NewLibraryControl2Client(ctx, libraryControl2SubConn, append(opts, dcerpc.WithNoBind(libraryControl2SubConn))...)

	objectManagement2SubConn, err := sub.SubConn(ctx, intmsobjectmanagement2.ObjectManagement2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		objectManagement2SubConn = sub
	}

	o.objectManagement2, err = intmsobjectmanagement2.NewObjectManagement2Client(ctx, objectManagement2SubConn, append(opts, dcerpc.WithNoBind(objectManagement2SubConn))...)

	objectManagement3SubConn, err := sub.SubConn(ctx, intmsobjectmanagement3.ObjectManagement3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		objectManagement3SubConn = sub
	}

	o.objectManagement3, err = intmsobjectmanagement3.NewObjectManagement3Client(ctx, objectManagement3SubConn, append(opts, dcerpc.WithNoBind(objectManagement3SubConn))...)

	robustNTMSMediaServices1SubConn, err := sub.SubConn(ctx, irobustntmsmediaservices1.RobustNTMSMediaServices1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		robustNTMSMediaServices1SubConn = sub
	}

	o.robustNTMSMediaServices1, err = irobustntmsmediaservices1.NewRobustNTMSMediaServices1Client(ctx, robustNTMSMediaServices1SubConn, append(opts, dcerpc.WithNoBind(robustNTMSMediaServices1SubConn))...)

	messengerSubConn, err := sub.SubConn(ctx, imessenger.MessengerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		messengerSubConn = sub
	}

	o.messenger, err = imessenger.NewMessengerClient(ctx, messengerSubConn, append(opts, dcerpc.WithNoBind(messengerSubConn))...)

	notifySinkSubConn, err := sub.SubConn(ctx, intmsnotifysink.NotifySinkSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		notifySinkSubConn = sub
	}

	o.notifySink, err = intmsnotifysink.NewNotifySinkClient(ctx, notifySinkSubConn, append(opts, dcerpc.WithNoBind(notifySinkSubConn))...)
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
		dcomClient:               o.dcomClient.IPID(ctx, ipid),
		libraryControl1:          o.libraryControl1.IPID(ctx, ipid),
		mediaServices1:           o.mediaServices1.IPID(ctx, ipid),
		objectInfo1:              o.objectInfo1.IPID(ctx, ipid),
		objectManagement1:        o.objectManagement1.IPID(ctx, ipid),
		session1:                 o.session1.IPID(ctx, ipid),
		clientSink:               o.clientSink.IPID(ctx, ipid),
		libraryControl2:          o.libraryControl2.IPID(ctx, ipid),
		objectManagement2:        o.objectManagement2.IPID(ctx, ipid),
		objectManagement3:        o.objectManagement3.IPID(ctx, ipid),
		robustNTMSMediaServices1: o.robustNTMSMediaServices1.IPID(ctx, ipid),
		messenger:                o.messenger.IPID(ctx, ipid),
		notifySink:               o.notifySink.IPID(ctx, ipid),
		cc:                       o.cc,
	}
}
