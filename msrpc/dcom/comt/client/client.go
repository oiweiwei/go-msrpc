package comt

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
	icomtrackinginfoevents "github.com/oiweiwei/go-msrpc/msrpc/dcom/comt/icomtrackinginfoevents/v0"
	igettrackingdata "github.com/oiweiwei/go-msrpc/msrpc/dcom/comt/igettrackingdata/v0"
	iprocessdump "github.com/oiweiwei/go-msrpc/msrpc/dcom/comt/iprocessdump/v0"
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
	_ = igettrackingdata.GoPackage
	_ = icomtrackinginfoevents.GoPackage
	_ = iprocessdump.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comt"
)

// dcom/comt client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	GetTrackingData() igettrackingdata.GetTrackingDataClient
	COMTrackingInfoEvents() icomtrackinginfoevents.COMTrackingInfoEventsClient
	ProcessDump() iprocessdump.ProcessDumpClient
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

	getTrackingData       igettrackingdata.GetTrackingDataClient
	comTrackingInfoEvents icomtrackinginfoevents.COMTrackingInfoEventsClient
	processDump           iprocessdump.ProcessDumpClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) GetTrackingData() igettrackingdata.GetTrackingDataClient {
	return o.getTrackingData
}

func (o *xxx_DefaultClient) COMTrackingInfoEvents() icomtrackinginfoevents.COMTrackingInfoEventsClient {
	return o.comTrackingInfoEvents
}

func (o *xxx_DefaultClient) ProcessDump() iprocessdump.ProcessDumpClient {
	return o.processDump
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(igettrackingdata.GetTrackingDataSyntaxV0_0),
		dcerpc.WithAbstractSyntax(icomtrackinginfoevents.COMTrackingInfoEventsSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iprocessdump.ProcessDumpSyntaxV0_0),
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

	getTrackingDataSubConn, err := sub.SubConn(ctx, igettrackingdata.GetTrackingDataSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		getTrackingDataSubConn = sub
	}

	o.getTrackingData, err = igettrackingdata.NewGetTrackingDataClient(ctx, getTrackingDataSubConn, append(opts, dcerpc.WithNoBind(getTrackingDataSubConn))...)

	comTrackingInfoEventsSubConn, err := sub.SubConn(ctx, icomtrackinginfoevents.COMTrackingInfoEventsSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		comTrackingInfoEventsSubConn = sub
	}

	o.comTrackingInfoEvents, err = icomtrackinginfoevents.NewCOMTrackingInfoEventsClient(ctx, comTrackingInfoEventsSubConn, append(opts, dcerpc.WithNoBind(comTrackingInfoEventsSubConn))...)

	processDumpSubConn, err := sub.SubConn(ctx, iprocessdump.ProcessDumpSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		processDumpSubConn = sub
	}

	o.processDump, err = iprocessdump.NewProcessDumpClient(ctx, processDumpSubConn, append(opts, dcerpc.WithNoBind(processDumpSubConn))...)
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
		getTrackingData:       o.getTrackingData.IPID(ctx, ipid),
		comTrackingInfoEvents: o.comTrackingInfoEvents.IPID(ctx, ipid),
		processDump:           o.processDump.IPID(ctx, ipid),
		cc:                    o.cc,
	}
}
