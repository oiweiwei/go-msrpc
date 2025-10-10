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
	NTMSLibraryControl1() intmslibrarycontrol1.NTMSLibraryControl1Client
	NTMSMediaServices1() intmsmediaservices1.NTMSMediaServices1Client
	NTMSObjectInfo1() intmsobjectinfo1.NTMSObjectInfo1Client
	NTMSObjectManagement1() intmsobjectmanagement1.NTMSObjectManagement1Client
	NTMSSession1() intmssession1.NTMSSession1Client
	ClientSink() iclientsink.ClientSinkClient
	NTMSLibraryControl2() intmslibrarycontrol2.NTMSLibraryControl2Client
	NTMSObjectManagement2() intmsobjectmanagement2.NTMSObjectManagement2Client
	NTMSObjectManagement3() intmsobjectmanagement3.NTMSObjectManagement3Client
	RobustNTMSMediaServices1() irobustntmsmediaservices1.RobustNTMSMediaServices1Client
	Messenger() imessenger.MessengerClient
	NTMSNotifySink() intmsnotifysink.NTMSNotifySinkClient
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

	ntmsLibraryControl1      intmslibrarycontrol1.NTMSLibraryControl1Client
	ntmsMediaServices1       intmsmediaservices1.NTMSMediaServices1Client
	ntmsObjectInfo1          intmsobjectinfo1.NTMSObjectInfo1Client
	ntmsObjectManagement1    intmsobjectmanagement1.NTMSObjectManagement1Client
	ntmsSession1             intmssession1.NTMSSession1Client
	clientSink               iclientsink.ClientSinkClient
	ntmsLibraryControl2      intmslibrarycontrol2.NTMSLibraryControl2Client
	ntmsObjectManagement2    intmsobjectmanagement2.NTMSObjectManagement2Client
	ntmsObjectManagement3    intmsobjectmanagement3.NTMSObjectManagement3Client
	robustNTMSMediaServices1 irobustntmsmediaservices1.RobustNTMSMediaServices1Client
	messenger                imessenger.MessengerClient
	ntmsNotifySink           intmsnotifysink.NTMSNotifySinkClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) NTMSLibraryControl1() intmslibrarycontrol1.NTMSLibraryControl1Client {
	return o.ntmsLibraryControl1
}

func (o *xxx_DefaultClient) NTMSMediaServices1() intmsmediaservices1.NTMSMediaServices1Client {
	return o.ntmsMediaServices1
}

func (o *xxx_DefaultClient) NTMSObjectInfo1() intmsobjectinfo1.NTMSObjectInfo1Client {
	return o.ntmsObjectInfo1
}

func (o *xxx_DefaultClient) NTMSObjectManagement1() intmsobjectmanagement1.NTMSObjectManagement1Client {
	return o.ntmsObjectManagement1
}

func (o *xxx_DefaultClient) NTMSSession1() intmssession1.NTMSSession1Client {
	return o.ntmsSession1
}

func (o *xxx_DefaultClient) ClientSink() iclientsink.ClientSinkClient {
	return o.clientSink
}

func (o *xxx_DefaultClient) NTMSLibraryControl2() intmslibrarycontrol2.NTMSLibraryControl2Client {
	return o.ntmsLibraryControl2
}

func (o *xxx_DefaultClient) NTMSObjectManagement2() intmsobjectmanagement2.NTMSObjectManagement2Client {
	return o.ntmsObjectManagement2
}

func (o *xxx_DefaultClient) NTMSObjectManagement3() intmsobjectmanagement3.NTMSObjectManagement3Client {
	return o.ntmsObjectManagement3
}

func (o *xxx_DefaultClient) RobustNTMSMediaServices1() irobustntmsmediaservices1.RobustNTMSMediaServices1Client {
	return o.robustNTMSMediaServices1
}

func (o *xxx_DefaultClient) Messenger() imessenger.MessengerClient {
	return o.messenger
}

func (o *xxx_DefaultClient) NTMSNotifySink() intmsnotifysink.NTMSNotifySinkClient {
	return o.ntmsNotifySink
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(intmslibrarycontrol1.NTMSLibraryControl1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsmediaservices1.NTMSMediaServices1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsobjectinfo1.NTMSObjectInfo1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsobjectmanagement1.NTMSObjectManagement1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmssession1.NTMSSession1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(iclientsink.ClientSinkSyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmslibrarycontrol2.NTMSLibraryControl2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsobjectmanagement2.NTMSObjectManagement2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsobjectmanagement3.NTMSObjectManagement3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(irobustntmsmediaservices1.RobustNTMSMediaServices1SyntaxV0_0),
		dcerpc.WithAbstractSyntax(imessenger.MessengerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(intmsnotifysink.NTMSNotifySinkSyntaxV0_0),
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

	ntmsLibraryControl1SubConn, err := sub.SubConn(ctx, intmslibrarycontrol1.NTMSLibraryControl1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		ntmsLibraryControl1SubConn = sub
	}

	o.ntmsLibraryControl1, err = intmslibrarycontrol1.NewNTMSLibraryControl1Client(ctx, ntmsLibraryControl1SubConn, append(opts, dcerpc.WithNoBind(ntmsLibraryControl1SubConn))...)

	ntmsMediaServices1SubConn, err := sub.SubConn(ctx, intmsmediaservices1.NTMSMediaServices1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		ntmsMediaServices1SubConn = sub
	}

	o.ntmsMediaServices1, err = intmsmediaservices1.NewNTMSMediaServices1Client(ctx, ntmsMediaServices1SubConn, append(opts, dcerpc.WithNoBind(ntmsMediaServices1SubConn))...)

	ntmsObjectInfo1SubConn, err := sub.SubConn(ctx, intmsobjectinfo1.NTMSObjectInfo1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		ntmsObjectInfo1SubConn = sub
	}

	o.ntmsObjectInfo1, err = intmsobjectinfo1.NewNTMSObjectInfo1Client(ctx, ntmsObjectInfo1SubConn, append(opts, dcerpc.WithNoBind(ntmsObjectInfo1SubConn))...)

	ntmsObjectManagement1SubConn, err := sub.SubConn(ctx, intmsobjectmanagement1.NTMSObjectManagement1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		ntmsObjectManagement1SubConn = sub
	}

	o.ntmsObjectManagement1, err = intmsobjectmanagement1.NewNTMSObjectManagement1Client(ctx, ntmsObjectManagement1SubConn, append(opts, dcerpc.WithNoBind(ntmsObjectManagement1SubConn))...)

	ntmsSession1SubConn, err := sub.SubConn(ctx, intmssession1.NTMSSession1SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		ntmsSession1SubConn = sub
	}

	o.ntmsSession1, err = intmssession1.NewNTMSSession1Client(ctx, ntmsSession1SubConn, append(opts, dcerpc.WithNoBind(ntmsSession1SubConn))...)

	clientSinkSubConn, err := sub.SubConn(ctx, iclientsink.ClientSinkSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		clientSinkSubConn = sub
	}

	o.clientSink, err = iclientsink.NewClientSinkClient(ctx, clientSinkSubConn, append(opts, dcerpc.WithNoBind(clientSinkSubConn))...)

	ntmsLibraryControl2SubConn, err := sub.SubConn(ctx, intmslibrarycontrol2.NTMSLibraryControl2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		ntmsLibraryControl2SubConn = sub
	}

	o.ntmsLibraryControl2, err = intmslibrarycontrol2.NewNTMSLibraryControl2Client(ctx, ntmsLibraryControl2SubConn, append(opts, dcerpc.WithNoBind(ntmsLibraryControl2SubConn))...)

	ntmsObjectManagement2SubConn, err := sub.SubConn(ctx, intmsobjectmanagement2.NTMSObjectManagement2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		ntmsObjectManagement2SubConn = sub
	}

	o.ntmsObjectManagement2, err = intmsobjectmanagement2.NewNTMSObjectManagement2Client(ctx, ntmsObjectManagement2SubConn, append(opts, dcerpc.WithNoBind(ntmsObjectManagement2SubConn))...)

	ntmsObjectManagement3SubConn, err := sub.SubConn(ctx, intmsobjectmanagement3.NTMSObjectManagement3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		ntmsObjectManagement3SubConn = sub
	}

	o.ntmsObjectManagement3, err = intmsobjectmanagement3.NewNTMSObjectManagement3Client(ctx, ntmsObjectManagement3SubConn, append(opts, dcerpc.WithNoBind(ntmsObjectManagement3SubConn))...)

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

	ntmsNotifySinkSubConn, err := sub.SubConn(ctx, intmsnotifysink.NTMSNotifySinkSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		ntmsNotifySinkSubConn = sub
	}

	o.ntmsNotifySink, err = intmsnotifysink.NewNTMSNotifySinkClient(ctx, ntmsNotifySinkSubConn, append(opts, dcerpc.WithNoBind(ntmsNotifySinkSubConn))...)
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
		ntmsLibraryControl1:      o.ntmsLibraryControl1.IPID(ctx, ipid),
		ntmsMediaServices1:       o.ntmsMediaServices1.IPID(ctx, ipid),
		ntmsObjectInfo1:          o.ntmsObjectInfo1.IPID(ctx, ipid),
		ntmsObjectManagement1:    o.ntmsObjectManagement1.IPID(ctx, ipid),
		ntmsSession1:             o.ntmsSession1.IPID(ctx, ipid),
		clientSink:               o.clientSink.IPID(ctx, ipid),
		ntmsLibraryControl2:      o.ntmsLibraryControl2.IPID(ctx, ipid),
		ntmsObjectManagement2:    o.ntmsObjectManagement2.IPID(ctx, ipid),
		ntmsObjectManagement3:    o.ntmsObjectManagement3.IPID(ctx, ipid),
		robustNTMSMediaServices1: o.robustNTMSMediaServices1.IPID(ctx, ipid),
		messenger:                o.messenger.IPID(ctx, ipid),
		ntmsNotifySink:           o.ntmsNotifySink.IPID(ctx, ipid),
		cc:                       o.cc,
	}
}
