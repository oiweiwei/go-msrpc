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
		op := &xxx_GetSubscriptionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubscriptionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubscriptionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SubscriptionID
		op := &xxx_SetSubscriptionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubscriptionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubscriptionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // SubscriptionName
		op := &xxx_GetSubscriptionNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubscriptionNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubscriptionName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // SubscriptionName
		op := &xxx_SetSubscriptionNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubscriptionNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubscriptionName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // PublisherID
		op := &xxx_GetPublisherIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPublisherIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPublisherID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // PublisherID
		op := &xxx_SetPublisherIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPublisherIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPublisherID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // EventClassID
		op := &xxx_GetEventClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // EventClassID
		op := &xxx_SetEventClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEventClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // MethodName
		op := &xxx_GetMethodNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMethodNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMethodName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // MethodName
		op := &xxx_SetMethodNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMethodNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMethodName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // SubscriberCLSID
		op := &xxx_GetSubscriberClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubscriberClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubscriberClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // SubscriberCLSID
		op := &xxx_SetSubscriberClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubscriberClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubscriberClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // SubscriberInterface
		op := &xxx_GetSubscriberInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubscriberInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubscriberInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // SubscriberInterface
		op := &xxx_SetSubscriberInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubscriberInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubscriberInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // PerUser
		op := &xxx_GetPerUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPerUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPerUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // PerUser
		op := &xxx_SetPerUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPerUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPerUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // OwnerSID
		op := &xxx_GetOwnerSIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOwnerSIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOwnerSID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // OwnerSID
		op := &xxx_SetOwnerSIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOwnerSIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOwnerSID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // Enabled
		op := &xxx_GetEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // Enabled
		op := &xxx_SetEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // Description
		op := &xxx_GetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // Description
		op := &xxx_SetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // MachineName
		op := &xxx_GetMachineNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMachineNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMachineName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // MachineName
		op := &xxx_SetMachineNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMachineNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMachineName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // GetPublisherProperty
		op := &xxx_GetPublisherPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPublisherPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPublisherProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // PutPublisherProperty
		op := &xxx_PutPublisherPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PutPublisherPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PutPublisherProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // RemovePublisherProperty
		op := &xxx_RemovePublisherPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemovePublisherPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemovePublisherProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // GetPublisherPropertyCollection
		op := &xxx_GetPublisherPropertyCollectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPublisherPropertyCollectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPublisherPropertyCollection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // GetSubscriberProperty
		op := &xxx_GetSubscriberPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubscriberPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubscriberProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // PutSubscriberProperty
		op := &xxx_PutSubscriberPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PutSubscriberPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PutSubscriberProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // RemoveSubscriberProperty
		op := &xxx_RemoveSubscriberPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveSubscriberPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveSubscriberProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // GetSubscriberPropertyCollection
		op := &xxx_GetSubscriberPropertyCollectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubscriberPropertyCollectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubscriberPropertyCollection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // InterfaceID
		op := &xxx_GetInterfaceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInterfaceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInterfaceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // InterfaceID
		op := &xxx_SetInterfaceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInterfaceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInterfaceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IEventSubscription
type UnimplementedEventSubscriptionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedEventSubscriptionServer) GetSubscriptionID(context.Context, *GetSubscriptionIDRequest) (*GetSubscriptionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetSubscriptionID(context.Context, *SetSubscriptionIDRequest) (*SetSubscriptionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetSubscriptionName(context.Context, *GetSubscriptionNameRequest) (*GetSubscriptionNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetSubscriptionName(context.Context, *SetSubscriptionNameRequest) (*SetSubscriptionNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetPublisherID(context.Context, *GetPublisherIDRequest) (*GetPublisherIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetPublisherID(context.Context, *SetPublisherIDRequest) (*SetPublisherIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetEventClassID(context.Context, *GetEventClassIDRequest) (*GetEventClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetEventClassID(context.Context, *SetEventClassIDRequest) (*SetEventClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetMethodName(context.Context, *GetMethodNameRequest) (*GetMethodNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetMethodName(context.Context, *SetMethodNameRequest) (*SetMethodNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetSubscriberClassID(context.Context, *GetSubscriberClassIDRequest) (*GetSubscriberClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetSubscriberClassID(context.Context, *SetSubscriberClassIDRequest) (*SetSubscriberClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetSubscriberInterface(context.Context, *GetSubscriberInterfaceRequest) (*GetSubscriberInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetSubscriberInterface(context.Context, *SetSubscriberInterfaceRequest) (*SetSubscriberInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetPerUser(context.Context, *GetPerUserRequest) (*GetPerUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetPerUser(context.Context, *SetPerUserRequest) (*SetPerUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetOwnerSID(context.Context, *GetOwnerSIDRequest) (*GetOwnerSIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetOwnerSID(context.Context, *SetOwnerSIDRequest) (*SetOwnerSIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetEnabled(context.Context, *GetEnabledRequest) (*GetEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetEnabled(context.Context, *SetEnabledRequest) (*SetEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetMachineName(context.Context, *GetMachineNameRequest) (*GetMachineNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetMachineName(context.Context, *SetMachineNameRequest) (*SetMachineNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetPublisherProperty(context.Context, *GetPublisherPropertyRequest) (*GetPublisherPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) PutPublisherProperty(context.Context, *PutPublisherPropertyRequest) (*PutPublisherPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) RemovePublisherProperty(context.Context, *RemovePublisherPropertyRequest) (*RemovePublisherPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetPublisherPropertyCollection(context.Context, *GetPublisherPropertyCollectionRequest) (*GetPublisherPropertyCollectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetSubscriberProperty(context.Context, *GetSubscriberPropertyRequest) (*GetSubscriberPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) PutSubscriberProperty(context.Context, *PutSubscriberPropertyRequest) (*PutSubscriberPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) RemoveSubscriberProperty(context.Context, *RemoveSubscriberPropertyRequest) (*RemoveSubscriberPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetSubscriberPropertyCollection(context.Context, *GetSubscriberPropertyCollectionRequest) (*GetSubscriberPropertyCollectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) GetInterfaceID(context.Context, *GetInterfaceIDRequest) (*GetInterfaceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscriptionServer) SetInterfaceID(context.Context, *SetInterfaceIDRequest) (*SetInterfaceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EventSubscriptionServer = (*UnimplementedEventSubscriptionServer)(nil)
