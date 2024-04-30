package ieventsubscription

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IEventSubscription server interface.
type EventSubscriptionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// SubscriptionID operation.
	GetSubscriptionID(context.Context, *GetSubscriptionIDRequest) (*GetSubscriptionIDResponse, error)

	// SubscriptionID operation.
	SetSubscriptionID(context.Context, *SetSubscriptionIDRequest) (*SetSubscriptionIDResponse, error)

	// SubscriptionName operation.
	GetSubscriptionName(context.Context, *GetSubscriptionNameRequest) (*GetSubscriptionNameResponse, error)

	// SubscriptionName operation.
	SetSubscriptionName(context.Context, *SetSubscriptionNameRequest) (*SetSubscriptionNameResponse, error)

	// PublisherID operation.
	GetPublisherID(context.Context, *GetPublisherIDRequest) (*GetPublisherIDResponse, error)

	// PublisherID operation.
	SetPublisherID(context.Context, *SetPublisherIDRequest) (*SetPublisherIDResponse, error)

	// EventClassID operation.
	GetEventClassID(context.Context, *GetEventClassIDRequest) (*GetEventClassIDResponse, error)

	// EventClassID operation.
	SetEventClassID(context.Context, *SetEventClassIDRequest) (*SetEventClassIDResponse, error)

	// MethodName operation.
	GetMethodName(context.Context, *GetMethodNameRequest) (*GetMethodNameResponse, error)

	// MethodName operation.
	SetMethodName(context.Context, *SetMethodNameRequest) (*SetMethodNameResponse, error)

	// SubscriberCLSID operation.
	GetSubscriberClassID(context.Context, *GetSubscriberClassIDRequest) (*GetSubscriberClassIDResponse, error)

	// SubscriberCLSID operation.
	SetSubscriberClassID(context.Context, *SetSubscriberClassIDRequest) (*SetSubscriberClassIDResponse, error)

	// SubscriberInterface operation.
	GetSubscriberInterface(context.Context, *GetSubscriberInterfaceRequest) (*GetSubscriberInterfaceResponse, error)

	// SubscriberInterface operation.
	SetSubscriberInterface(context.Context, *SetSubscriberInterfaceRequest) (*SetSubscriberInterfaceResponse, error)

	// PerUser operation.
	GetPerUser(context.Context, *GetPerUserRequest) (*GetPerUserResponse, error)

	// PerUser operation.
	SetPerUser(context.Context, *SetPerUserRequest) (*SetPerUserResponse, error)

	// OwnerSID operation.
	GetOwnerSID(context.Context, *GetOwnerSIDRequest) (*GetOwnerSIDResponse, error)

	// OwnerSID operation.
	SetOwnerSID(context.Context, *SetOwnerSIDRequest) (*SetOwnerSIDResponse, error)

	// Enabled operation.
	GetEnabled(context.Context, *GetEnabledRequest) (*GetEnabledResponse, error)

	// Enabled operation.
	SetEnabled(context.Context, *SetEnabledRequest) (*SetEnabledResponse, error)

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	// Description operation.
	SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error)

	// MachineName operation.
	GetMachineName(context.Context, *GetMachineNameRequest) (*GetMachineNameResponse, error)

	// MachineName operation.
	SetMachineName(context.Context, *SetMachineNameRequest) (*SetMachineNameResponse, error)

	// The GetPublisherProperty method gets the application-specific publisher property
	// of the subscription. See publisher properties in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetPublisherProperty(context.Context, *GetPublisherPropertyRequest) (*GetPublisherPropertyResponse, error)

	// The PutPublisherProperty method sets the application-specific publisher property
	// of the subscription. If the subscription does not already have a publisher property,
	// this method will add it to the publisher property collection. If the same name property
	// exists, it would be overwritten by the new value provided as part of this method.
	// See publisher properties in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	PutPublisherProperty(context.Context, *PutPublisherPropertyRequest) (*PutPublisherPropertyResponse, error)

	// The RemovePublisherProperty method removes the specified application-specific publisher
	// property for the subscription. See publisher properties in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	RemovePublisherProperty(context.Context, *RemovePublisherPropertyRequest) (*RemovePublisherPropertyResponse, error)

	// The GetPublisherPropertyCollection method gets all the application-specific publisher
	// properties as a collection of the subscription. See publisher properties in section
	// 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetPublisherPropertyCollection(context.Context, *GetPublisherPropertyCollectionRequest) (*GetPublisherPropertyCollectionResponse, error)

	// The GetSubscriberProperty method gets the value of an application-specific subscriber
	// property of the subscription, as specified in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetSubscriberProperty(context.Context, *GetSubscriberPropertyRequest) (*GetSubscriberPropertyResponse, error)

	// The PutSubscriberProperty method sets the value of an application-specific subscriber
	// property of the subscription, as specified in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	PutSubscriberProperty(context.Context, *PutSubscriberPropertyRequest) (*PutSubscriberPropertyResponse, error)

	// The RemoveSubscriberProperty method removes the specified application-specific subscriber
	// property for the subscription, as specified in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	RemoveSubscriberProperty(context.Context, *RemoveSubscriberPropertyRequest) (*RemoveSubscriberPropertyResponse, error)

	// The GetSubscriberPropertyCollection method gets the collection of all the application-specific
	// subscriber properties for the subscription, as specified in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetSubscriberPropertyCollection(context.Context, *GetSubscriberPropertyCollectionRequest) (*GetSubscriberPropertyCollectionResponse, error)

	// InterfaceID operation.
	GetInterfaceID(context.Context, *GetInterfaceIDRequest) (*GetInterfaceIDResponse, error)

	// InterfaceID operation.
	SetInterfaceID(context.Context, *SetInterfaceIDRequest) (*SetInterfaceIDResponse, error)
}

func RegisterEventSubscriptionServer(conn dcerpc.Conn, o EventSubscriptionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventSubscriptionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventSubscriptionSyntaxV0_0))...)
}

func NewEventSubscriptionServerHandle(o EventSubscriptionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventSubscriptionServerHandle(ctx, o, opNum, r)
	}
}

func EventSubscriptionServerHandle(ctx context.Context, o EventSubscriptionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // SubscriptionID
		in := &GetSubscriptionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubscriptionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // SubscriptionID
		in := &SetSubscriptionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubscriptionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // SubscriptionName
		in := &GetSubscriptionNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubscriptionName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // SubscriptionName
		in := &SetSubscriptionNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubscriptionName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // PublisherID
		in := &GetPublisherIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPublisherID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // PublisherID
		in := &SetPublisherIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPublisherID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // EventClassID
		in := &GetEventClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // EventClassID
		in := &SetEventClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEventClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // MethodName
		in := &GetMethodNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMethodName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // MethodName
		in := &SetMethodNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMethodName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // SubscriberCLSID
		in := &GetSubscriberClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubscriberClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // SubscriberCLSID
		in := &SetSubscriberClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubscriberClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // SubscriberInterface
		in := &GetSubscriberInterfaceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubscriberInterface(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // SubscriberInterface
		in := &SetSubscriberInterfaceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubscriberInterface(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // PerUser
		in := &GetPerUserRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPerUser(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // PerUser
		in := &SetPerUserRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPerUser(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // OwnerSID
		in := &GetOwnerSIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOwnerSID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // OwnerSID
		in := &SetOwnerSIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOwnerSID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // Enabled
		in := &GetEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // Enabled
		in := &SetEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // Description
		in := &GetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // Description
		in := &SetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // MachineName
		in := &GetMachineNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMachineName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // MachineName
		in := &SetMachineNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMachineName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // GetPublisherProperty
		in := &GetPublisherPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPublisherProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // PutPublisherProperty
		in := &PutPublisherPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PutPublisherProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // RemovePublisherProperty
		in := &RemovePublisherPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemovePublisherProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // GetPublisherPropertyCollection
		in := &GetPublisherPropertyCollectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPublisherPropertyCollection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // GetSubscriberProperty
		in := &GetSubscriberPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubscriberProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // PutSubscriberProperty
		in := &PutSubscriberPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PutSubscriberProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // RemoveSubscriberProperty
		in := &RemoveSubscriberPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveSubscriberProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // GetSubscriberPropertyCollection
		in := &GetSubscriberPropertyCollectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubscriberPropertyCollection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // InterfaceID
		in := &GetInterfaceIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetInterfaceID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // InterfaceID
		in := &SetInterfaceIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetInterfaceID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
