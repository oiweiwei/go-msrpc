package comev

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
	ienumeventobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ienumeventobject/v0"
	ieventclass "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventclass/v0"
	ieventclass2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventclass2/v0"
	ieventclass3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventclass3/v0"
	ieventobjectcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventobjectcollection/v0"
	ieventsubscription "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsubscription/v0"
	ieventsubscription2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsubscription2/v0"
	ieventsubscription3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsubscription3/v0"
	ieventsystem "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsystem/v0"
	ieventsystem2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsystem2/v0"
	ieventsysteminitialize "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsysteminitialize/v0"
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
	_ = ieventsystem.GoPackage
	_ = ieventclass.GoPackage
	_ = ieventclass2.GoPackage
	_ = ienumeventobject.GoPackage
	_ = ieventobjectcollection.GoPackage
	_ = ieventsubscription.GoPackage
	_ = ieventsubscription2.GoPackage
	_ = ieventclass3.GoPackage
	_ = ieventsubscription3.GoPackage
	_ = ieventsystem2.GoPackage
	_ = ieventsysteminitialize.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comev"
)

// dcom/comev client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	EventSystem() ieventsystem.EventSystemClient
	EventClass() ieventclass.EventClassClient
	EventClass2() ieventclass2.EventClass2Client
	EnumEventObject() ienumeventobject.EnumEventObjectClient
	EventObjectCollection() ieventobjectcollection.EventObjectCollectionClient
	EventSubscription() ieventsubscription.EventSubscriptionClient
	EventSubscription2() ieventsubscription2.EventSubscription2Client
	EventClass3() ieventclass3.EventClass3Client
	EventSubscription3() ieventsubscription3.EventSubscription3Client
	EventSystem2() ieventsystem2.EventSystem2Client
	EventSystemInitialize() ieventsysteminitialize.EventSystemInitializeClient
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

	eventSystem           ieventsystem.EventSystemClient
	eventClass            ieventclass.EventClassClient
	eventClass2           ieventclass2.EventClass2Client
	enumEventObject       ienumeventobject.EnumEventObjectClient
	eventObjectCollection ieventobjectcollection.EventObjectCollectionClient
	eventSubscription     ieventsubscription.EventSubscriptionClient
	eventSubscription2    ieventsubscription2.EventSubscription2Client
	eventClass3           ieventclass3.EventClass3Client
	eventSubscription3    ieventsubscription3.EventSubscription3Client
	eventSystem2          ieventsystem2.EventSystem2Client
	eventSystemInitialize ieventsysteminitialize.EventSystemInitializeClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) EventSystem() ieventsystem.EventSystemClient {
	return o.eventSystem
}

func (o *xxx_DefaultClient) EventClass() ieventclass.EventClassClient {
	return o.eventClass
}

func (o *xxx_DefaultClient) EventClass2() ieventclass2.EventClass2Client {
	return o.eventClass2
}

func (o *xxx_DefaultClient) EnumEventObject() ienumeventobject.EnumEventObjectClient {
	return o.enumEventObject
}

func (o *xxx_DefaultClient) EventObjectCollection() ieventobjectcollection.EventObjectCollectionClient {
	return o.eventObjectCollection
}

func (o *xxx_DefaultClient) EventSubscription() ieventsubscription.EventSubscriptionClient {
	return o.eventSubscription
}

func (o *xxx_DefaultClient) EventSubscription2() ieventsubscription2.EventSubscription2Client {
	return o.eventSubscription2
}

func (o *xxx_DefaultClient) EventClass3() ieventclass3.EventClass3Client {
	return o.eventClass3
}

func (o *xxx_DefaultClient) EventSubscription3() ieventsubscription3.EventSubscription3Client {
	return o.eventSubscription3
}

func (o *xxx_DefaultClient) EventSystem2() ieventsystem2.EventSystem2Client {
	return o.eventSystem2
}

func (o *xxx_DefaultClient) EventSystemInitialize() ieventsysteminitialize.EventSystemInitializeClient {
	return o.eventSystemInitialize
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(ieventsystem.EventSystemSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ieventclass.EventClassSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ieventclass2.EventClass2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ienumeventobject.EnumEventObjectSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ieventobjectcollection.EventObjectCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ieventsubscription.EventSubscriptionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ieventsubscription2.EventSubscription2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ieventclass3.EventClass3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ieventsubscription3.EventSubscription3SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ieventsystem2.EventSystem2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ieventsysteminitialize.EventSystemInitializeSyntaxV0_0),
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

	eventSystemSubConn, err := sub.SubConn(ctx, ieventsystem.EventSystemSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		eventSystemSubConn = sub
	}

	o.eventSystem, err = ieventsystem.NewEventSystemClient(ctx, eventSystemSubConn, append(opts, dcerpc.WithNoBind(eventSystemSubConn))...)

	eventClassSubConn, err := sub.SubConn(ctx, ieventclass.EventClassSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		eventClassSubConn = sub
	}

	o.eventClass, err = ieventclass.NewEventClassClient(ctx, eventClassSubConn, append(opts, dcerpc.WithNoBind(eventClassSubConn))...)

	eventClass2SubConn, err := sub.SubConn(ctx, ieventclass2.EventClass2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		eventClass2SubConn = sub
	}

	o.eventClass2, err = ieventclass2.NewEventClass2Client(ctx, eventClass2SubConn, append(opts, dcerpc.WithNoBind(eventClass2SubConn))...)

	enumEventObjectSubConn, err := sub.SubConn(ctx, ienumeventobject.EnumEventObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		enumEventObjectSubConn = sub
	}

	o.enumEventObject, err = ienumeventobject.NewEnumEventObjectClient(ctx, enumEventObjectSubConn, append(opts, dcerpc.WithNoBind(enumEventObjectSubConn))...)

	eventObjectCollectionSubConn, err := sub.SubConn(ctx, ieventobjectcollection.EventObjectCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		eventObjectCollectionSubConn = sub
	}

	o.eventObjectCollection, err = ieventobjectcollection.NewEventObjectCollectionClient(ctx, eventObjectCollectionSubConn, append(opts, dcerpc.WithNoBind(eventObjectCollectionSubConn))...)

	eventSubscriptionSubConn, err := sub.SubConn(ctx, ieventsubscription.EventSubscriptionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		eventSubscriptionSubConn = sub
	}

	o.eventSubscription, err = ieventsubscription.NewEventSubscriptionClient(ctx, eventSubscriptionSubConn, append(opts, dcerpc.WithNoBind(eventSubscriptionSubConn))...)

	eventSubscription2SubConn, err := sub.SubConn(ctx, ieventsubscription2.EventSubscription2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		eventSubscription2SubConn = sub
	}

	o.eventSubscription2, err = ieventsubscription2.NewEventSubscription2Client(ctx, eventSubscription2SubConn, append(opts, dcerpc.WithNoBind(eventSubscription2SubConn))...)

	eventClass3SubConn, err := sub.SubConn(ctx, ieventclass3.EventClass3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		eventClass3SubConn = sub
	}

	o.eventClass3, err = ieventclass3.NewEventClass3Client(ctx, eventClass3SubConn, append(opts, dcerpc.WithNoBind(eventClass3SubConn))...)

	eventSubscription3SubConn, err := sub.SubConn(ctx, ieventsubscription3.EventSubscription3SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		eventSubscription3SubConn = sub
	}

	o.eventSubscription3, err = ieventsubscription3.NewEventSubscription3Client(ctx, eventSubscription3SubConn, append(opts, dcerpc.WithNoBind(eventSubscription3SubConn))...)

	eventSystem2SubConn, err := sub.SubConn(ctx, ieventsystem2.EventSystem2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		eventSystem2SubConn = sub
	}

	o.eventSystem2, err = ieventsystem2.NewEventSystem2Client(ctx, eventSystem2SubConn, append(opts, dcerpc.WithNoBind(eventSystem2SubConn))...)

	eventSystemInitializeSubConn, err := sub.SubConn(ctx, ieventsysteminitialize.EventSystemInitializeSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		eventSystemInitializeSubConn = sub
	}

	o.eventSystemInitialize, err = ieventsysteminitialize.NewEventSystemInitializeClient(ctx, eventSystemInitializeSubConn, append(opts, dcerpc.WithNoBind(eventSystemInitializeSubConn))...)
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
		eventSystem:           o.eventSystem.IPID(ctx, ipid),
		eventClass:            o.eventClass.IPID(ctx, ipid),
		eventClass2:           o.eventClass2.IPID(ctx, ipid),
		enumEventObject:       o.enumEventObject.IPID(ctx, ipid),
		eventObjectCollection: o.eventObjectCollection.IPID(ctx, ipid),
		eventSubscription:     o.eventSubscription.IPID(ctx, ipid),
		eventSubscription2:    o.eventSubscription2.IPID(ctx, ipid),
		eventClass3:           o.eventClass3.IPID(ctx, ipid),
		eventSubscription3:    o.eventSubscription3.IPID(ctx, ipid),
		eventSystem2:          o.eventSystem2.IPID(ctx, ipid),
		eventSystemInitialize: o.eventSystemInitialize.IPID(ctx, ipid),
		cc:                    o.cc,
	}
}
