package ieventsubscription

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	comev "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
	_ = comev.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comev"
)

var (
	// IEventSubscription interface identifier 4a6b0e15-2e38-11d1-9965-00c04fbbb345
	EventSubscriptionIID = &dcom.IID{Data1: 0x4a6b0e15, Data2: 0x2e38, Data3: 0x11d1, Data4: []byte{0x99, 0x65, 0x00, 0xc0, 0x4f, 0xbb, 0xb3, 0x45}}
	// Syntax UUID
	EventSubscriptionSyntaxUUID = &uuid.UUID{TimeLow: 0x4a6b0e15, TimeMid: 0x2e38, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x65, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xbb, 0xb3, 0x45}}
	// Syntax ID
	EventSubscriptionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EventSubscriptionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IEventSubscription interface.
type EventSubscriptionClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// SubscriptionID operation.
	GetSubscriptionID(context.Context, *GetSubscriptionIDRequest, ...dcerpc.CallOption) (*GetSubscriptionIDResponse, error)

	// SubscriptionID operation.
	SetSubscriptionID(context.Context, *SetSubscriptionIDRequest, ...dcerpc.CallOption) (*SetSubscriptionIDResponse, error)

	// SubscriptionName operation.
	GetSubscriptionName(context.Context, *GetSubscriptionNameRequest, ...dcerpc.CallOption) (*GetSubscriptionNameResponse, error)

	// SubscriptionName operation.
	SetSubscriptionName(context.Context, *SetSubscriptionNameRequest, ...dcerpc.CallOption) (*SetSubscriptionNameResponse, error)

	// PublisherID operation.
	GetPublisherID(context.Context, *GetPublisherIDRequest, ...dcerpc.CallOption) (*GetPublisherIDResponse, error)

	// PublisherID operation.
	SetPublisherID(context.Context, *SetPublisherIDRequest, ...dcerpc.CallOption) (*SetPublisherIDResponse, error)

	// EventClassID operation.
	GetEventClassID(context.Context, *GetEventClassIDRequest, ...dcerpc.CallOption) (*GetEventClassIDResponse, error)

	// EventClassID operation.
	SetEventClassID(context.Context, *SetEventClassIDRequest, ...dcerpc.CallOption) (*SetEventClassIDResponse, error)

	// MethodName operation.
	GetMethodName(context.Context, *GetMethodNameRequest, ...dcerpc.CallOption) (*GetMethodNameResponse, error)

	// MethodName operation.
	SetMethodName(context.Context, *SetMethodNameRequest, ...dcerpc.CallOption) (*SetMethodNameResponse, error)

	// SubscriberCLSID operation.
	GetSubscriberClassID(context.Context, *GetSubscriberClassIDRequest, ...dcerpc.CallOption) (*GetSubscriberClassIDResponse, error)

	// SubscriberCLSID operation.
	SetSubscriberClassID(context.Context, *SetSubscriberClassIDRequest, ...dcerpc.CallOption) (*SetSubscriberClassIDResponse, error)

	// SubscriberInterface operation.
	GetSubscriberInterface(context.Context, *GetSubscriberInterfaceRequest, ...dcerpc.CallOption) (*GetSubscriberInterfaceResponse, error)

	// SubscriberInterface operation.
	SetSubscriberInterface(context.Context, *SetSubscriberInterfaceRequest, ...dcerpc.CallOption) (*SetSubscriberInterfaceResponse, error)

	// PerUser operation.
	GetPerUser(context.Context, *GetPerUserRequest, ...dcerpc.CallOption) (*GetPerUserResponse, error)

	// PerUser operation.
	SetPerUser(context.Context, *SetPerUserRequest, ...dcerpc.CallOption) (*SetPerUserResponse, error)

	// OwnerSID operation.
	GetOwnerSID(context.Context, *GetOwnerSIDRequest, ...dcerpc.CallOption) (*GetOwnerSIDResponse, error)

	// OwnerSID operation.
	SetOwnerSID(context.Context, *SetOwnerSIDRequest, ...dcerpc.CallOption) (*SetOwnerSIDResponse, error)

	// Enabled operation.
	GetEnabled(context.Context, *GetEnabledRequest, ...dcerpc.CallOption) (*GetEnabledResponse, error)

	// Enabled operation.
	SetEnabled(context.Context, *SetEnabledRequest, ...dcerpc.CallOption) (*SetEnabledResponse, error)

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest, ...dcerpc.CallOption) (*GetDescriptionResponse, error)

	// Description operation.
	SetDescription(context.Context, *SetDescriptionRequest, ...dcerpc.CallOption) (*SetDescriptionResponse, error)

	// MachineName operation.
	GetMachineName(context.Context, *GetMachineNameRequest, ...dcerpc.CallOption) (*GetMachineNameResponse, error)

	// MachineName operation.
	SetMachineName(context.Context, *SetMachineNameRequest, ...dcerpc.CallOption) (*SetMachineNameResponse, error)

	// The GetPublisherProperty method gets the application-specific publisher property
	// of the subscription. See publisher properties in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetPublisherProperty(context.Context, *GetPublisherPropertyRequest, ...dcerpc.CallOption) (*GetPublisherPropertyResponse, error)

	// The PutPublisherProperty method sets the application-specific publisher property
	// of the subscription. If the subscription does not already have a publisher property,
	// this method will add it to the publisher property collection. If the same name property
	// exists, it would be overwritten by the new value provided as part of this method.
	// See publisher properties in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	PutPublisherProperty(context.Context, *PutPublisherPropertyRequest, ...dcerpc.CallOption) (*PutPublisherPropertyResponse, error)

	// The RemovePublisherProperty method removes the specified application-specific publisher
	// property for the subscription. See publisher properties in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	RemovePublisherProperty(context.Context, *RemovePublisherPropertyRequest, ...dcerpc.CallOption) (*RemovePublisherPropertyResponse, error)

	// The GetPublisherPropertyCollection method gets all the application-specific publisher
	// properties as a collection of the subscription. See publisher properties in section
	// 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetPublisherPropertyCollection(context.Context, *GetPublisherPropertyCollectionRequest, ...dcerpc.CallOption) (*GetPublisherPropertyCollectionResponse, error)

	// The GetSubscriberProperty method gets the value of an application-specific subscriber
	// property of the subscription, as specified in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetSubscriberProperty(context.Context, *GetSubscriberPropertyRequest, ...dcerpc.CallOption) (*GetSubscriberPropertyResponse, error)

	// The PutSubscriberProperty method sets the value of an application-specific subscriber
	// property of the subscription, as specified in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	PutSubscriberProperty(context.Context, *PutSubscriberPropertyRequest, ...dcerpc.CallOption) (*PutSubscriberPropertyResponse, error)

	// The RemoveSubscriberProperty method removes the specified application-specific subscriber
	// property for the subscription, as specified in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	RemoveSubscriberProperty(context.Context, *RemoveSubscriberPropertyRequest, ...dcerpc.CallOption) (*RemoveSubscriberPropertyResponse, error)

	// The GetSubscriberPropertyCollection method gets the collection of all the application-specific
	// subscriber properties for the subscription, as specified in section 3.1.1.2.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetSubscriberPropertyCollection(context.Context, *GetSubscriberPropertyCollectionRequest, ...dcerpc.CallOption) (*GetSubscriberPropertyCollectionResponse, error)

	// InterfaceID operation.
	GetInterfaceID(context.Context, *GetInterfaceIDRequest, ...dcerpc.CallOption) (*GetInterfaceIDResponse, error)

	// InterfaceID operation.
	SetInterfaceID(context.Context, *SetInterfaceIDRequest, ...dcerpc.CallOption) (*SetInterfaceIDResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) EventSubscriptionClient
}

type xxx_DefaultEventSubscriptionClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultEventSubscriptionClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultEventSubscriptionClient) GetSubscriptionID(ctx context.Context, in *GetSubscriptionIDRequest, opts ...dcerpc.CallOption) (*GetSubscriptionIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSubscriptionIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetSubscriptionID(ctx context.Context, in *SetSubscriptionIDRequest, opts ...dcerpc.CallOption) (*SetSubscriptionIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSubscriptionIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetSubscriptionName(ctx context.Context, in *GetSubscriptionNameRequest, opts ...dcerpc.CallOption) (*GetSubscriptionNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSubscriptionNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetSubscriptionName(ctx context.Context, in *SetSubscriptionNameRequest, opts ...dcerpc.CallOption) (*SetSubscriptionNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSubscriptionNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetPublisherID(ctx context.Context, in *GetPublisherIDRequest, opts ...dcerpc.CallOption) (*GetPublisherIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPublisherIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetPublisherID(ctx context.Context, in *SetPublisherIDRequest, opts ...dcerpc.CallOption) (*SetPublisherIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetPublisherIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetEventClassID(ctx context.Context, in *GetEventClassIDRequest, opts ...dcerpc.CallOption) (*GetEventClassIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetEventClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetEventClassID(ctx context.Context, in *SetEventClassIDRequest, opts ...dcerpc.CallOption) (*SetEventClassIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetEventClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetMethodName(ctx context.Context, in *GetMethodNameRequest, opts ...dcerpc.CallOption) (*GetMethodNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetMethodNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetMethodName(ctx context.Context, in *SetMethodNameRequest, opts ...dcerpc.CallOption) (*SetMethodNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetMethodNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetSubscriberClassID(ctx context.Context, in *GetSubscriberClassIDRequest, opts ...dcerpc.CallOption) (*GetSubscriberClassIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSubscriberClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetSubscriberClassID(ctx context.Context, in *SetSubscriberClassIDRequest, opts ...dcerpc.CallOption) (*SetSubscriberClassIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSubscriberClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetSubscriberInterface(ctx context.Context, in *GetSubscriberInterfaceRequest, opts ...dcerpc.CallOption) (*GetSubscriberInterfaceResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSubscriberInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetSubscriberInterface(ctx context.Context, in *SetSubscriberInterfaceRequest, opts ...dcerpc.CallOption) (*SetSubscriberInterfaceResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSubscriberInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetPerUser(ctx context.Context, in *GetPerUserRequest, opts ...dcerpc.CallOption) (*GetPerUserResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPerUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetPerUser(ctx context.Context, in *SetPerUserRequest, opts ...dcerpc.CallOption) (*SetPerUserResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetPerUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetOwnerSID(ctx context.Context, in *GetOwnerSIDRequest, opts ...dcerpc.CallOption) (*GetOwnerSIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetOwnerSIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetOwnerSID(ctx context.Context, in *SetOwnerSIDRequest, opts ...dcerpc.CallOption) (*SetOwnerSIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetOwnerSIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetEnabled(ctx context.Context, in *GetEnabledRequest, opts ...dcerpc.CallOption) (*GetEnabledResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetEnabled(ctx context.Context, in *SetEnabledRequest, opts ...dcerpc.CallOption) (*SetEnabledResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetDescription(ctx context.Context, in *GetDescriptionRequest, opts ...dcerpc.CallOption) (*GetDescriptionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDescriptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetDescription(ctx context.Context, in *SetDescriptionRequest, opts ...dcerpc.CallOption) (*SetDescriptionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetDescriptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetMachineName(ctx context.Context, in *GetMachineNameRequest, opts ...dcerpc.CallOption) (*GetMachineNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetMachineNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetMachineName(ctx context.Context, in *SetMachineNameRequest, opts ...dcerpc.CallOption) (*SetMachineNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetMachineNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetPublisherProperty(ctx context.Context, in *GetPublisherPropertyRequest, opts ...dcerpc.CallOption) (*GetPublisherPropertyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPublisherPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) PutPublisherProperty(ctx context.Context, in *PutPublisherPropertyRequest, opts ...dcerpc.CallOption) (*PutPublisherPropertyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PutPublisherPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) RemovePublisherProperty(ctx context.Context, in *RemovePublisherPropertyRequest, opts ...dcerpc.CallOption) (*RemovePublisherPropertyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemovePublisherPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetPublisherPropertyCollection(ctx context.Context, in *GetPublisherPropertyCollectionRequest, opts ...dcerpc.CallOption) (*GetPublisherPropertyCollectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPublisherPropertyCollectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetSubscriberProperty(ctx context.Context, in *GetSubscriberPropertyRequest, opts ...dcerpc.CallOption) (*GetSubscriberPropertyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSubscriberPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) PutSubscriberProperty(ctx context.Context, in *PutSubscriberPropertyRequest, opts ...dcerpc.CallOption) (*PutSubscriberPropertyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PutSubscriberPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) RemoveSubscriberProperty(ctx context.Context, in *RemoveSubscriberPropertyRequest, opts ...dcerpc.CallOption) (*RemoveSubscriberPropertyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveSubscriberPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetSubscriberPropertyCollection(ctx context.Context, in *GetSubscriberPropertyCollectionRequest, opts ...dcerpc.CallOption) (*GetSubscriberPropertyCollectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSubscriberPropertyCollectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) GetInterfaceID(ctx context.Context, in *GetInterfaceIDRequest, opts ...dcerpc.CallOption) (*GetInterfaceIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetInterfaceIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) SetInterfaceID(ctx context.Context, in *SetInterfaceIDRequest, opts ...dcerpc.CallOption) (*SetInterfaceIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInterfaceIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscriptionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventSubscriptionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultEventSubscriptionClient) IPID(ctx context.Context, ipid *dcom.IPID) EventSubscriptionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultEventSubscriptionClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewEventSubscriptionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventSubscriptionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventSubscriptionSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultEventSubscriptionClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetSubscriptionIDOperation structure represents the SubscriptionID operation
type xxx_GetSubscriptionIDOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriptionID *oaut.String   `idl:"name:pbstrSubscriptionID" json:"subscription_id"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubscriptionIDOperation) OpNum() int { return 7 }

func (o *xxx_GetSubscriptionIDOperation) OpName() string {
	return "/IEventSubscription/v0/SubscriptionID"
}

func (o *xxx_GetSubscriptionIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriptionIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriptionIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriptionIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriptionIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrSubscriptionID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SubscriptionID != nil {
			_ptr_pbstrSubscriptionID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriptionID != nil {
					if err := o.SubscriptionID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriptionID, _ptr_pbstrSubscriptionID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriptionIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrSubscriptionID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrSubscriptionID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriptionID == nil {
				o.SubscriptionID = &oaut.String{}
			}
			if err := o.SubscriptionID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrSubscriptionID := func(ptr interface{}) { o.SubscriptionID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SubscriptionID, _s_pbstrSubscriptionID, _ptr_pbstrSubscriptionID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSubscriptionIDRequest structure represents the SubscriptionID operation request
type GetSubscriptionIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubscriptionIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriptionIDOperation) *xxx_GetSubscriptionIDOperation {
	if op == nil {
		op = &xxx_GetSubscriptionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSubscriptionIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriptionIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubscriptionIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSubscriptionIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriptionIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubscriptionIDResponse structure represents the SubscriptionID operation response
type GetSubscriptionIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriptionID *oaut.String   `idl:"name:pbstrSubscriptionID" json:"subscription_id"`
	// Return: The SubscriptionID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubscriptionIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriptionIDOperation) *xxx_GetSubscriptionIDOperation {
	if op == nil {
		op = &xxx_GetSubscriptionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SubscriptionID = o.SubscriptionID
	op.Return = o.Return
	return op
}

func (o *GetSubscriptionIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriptionIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SubscriptionID = op.SubscriptionID
	o.Return = op.Return
}
func (o *GetSubscriptionIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSubscriptionIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriptionIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubscriptionIDOperation structure represents the SubscriptionID operation
type xxx_SetSubscriptionIDOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriptionID *oaut.String   `idl:"name:bstrSubscriptionID" json:"subscription_id"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubscriptionIDOperation) OpNum() int { return 8 }

func (o *xxx_SetSubscriptionIDOperation) OpName() string {
	return "/IEventSubscription/v0/SubscriptionID"
}

func (o *xxx_SetSubscriptionIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriptionIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrSubscriptionID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SubscriptionID != nil {
			_ptr_bstrSubscriptionID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriptionID != nil {
					if err := o.SubscriptionID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriptionID, _ptr_bstrSubscriptionID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriptionIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrSubscriptionID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSubscriptionID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriptionID == nil {
				o.SubscriptionID = &oaut.String{}
			}
			if err := o.SubscriptionID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSubscriptionID := func(ptr interface{}) { o.SubscriptionID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SubscriptionID, _s_bstrSubscriptionID, _ptr_bstrSubscriptionID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriptionIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriptionIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriptionIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSubscriptionIDRequest structure represents the SubscriptionID operation request
type SetSubscriptionIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	SubscriptionID *oaut.String   `idl:"name:bstrSubscriptionID" json:"subscription_id"`
}

func (o *SetSubscriptionIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSubscriptionIDOperation) *xxx_SetSubscriptionIDOperation {
	if op == nil {
		op = &xxx_SetSubscriptionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SubscriptionID = o.SubscriptionID
	return op
}

func (o *SetSubscriptionIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriptionIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SubscriptionID = op.SubscriptionID
}
func (o *SetSubscriptionIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSubscriptionIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriptionIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubscriptionIDResponse structure represents the SubscriptionID operation response
type SetSubscriptionIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SubscriptionID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSubscriptionIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSubscriptionIDOperation) *xxx_SetSubscriptionIDOperation {
	if op == nil {
		op = &xxx_SetSubscriptionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSubscriptionIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriptionIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSubscriptionIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSubscriptionIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriptionIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubscriptionNameOperation structure represents the SubscriptionName operation
type xxx_GetSubscriptionNameOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriptionName *oaut.String   `idl:"name:pbstrSubscriptionName" json:"subscription_name"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubscriptionNameOperation) OpNum() int { return 9 }

func (o *xxx_GetSubscriptionNameOperation) OpName() string {
	return "/IEventSubscription/v0/SubscriptionName"
}

func (o *xxx_GetSubscriptionNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriptionNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriptionNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriptionNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriptionNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrSubscriptionName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SubscriptionName != nil {
			_ptr_pbstrSubscriptionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriptionName != nil {
					if err := o.SubscriptionName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriptionName, _ptr_pbstrSubscriptionName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriptionNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrSubscriptionName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrSubscriptionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriptionName == nil {
				o.SubscriptionName = &oaut.String{}
			}
			if err := o.SubscriptionName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrSubscriptionName := func(ptr interface{}) { o.SubscriptionName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SubscriptionName, _s_pbstrSubscriptionName, _ptr_pbstrSubscriptionName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSubscriptionNameRequest structure represents the SubscriptionName operation request
type GetSubscriptionNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubscriptionNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriptionNameOperation) *xxx_GetSubscriptionNameOperation {
	if op == nil {
		op = &xxx_GetSubscriptionNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSubscriptionNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriptionNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubscriptionNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSubscriptionNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriptionNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubscriptionNameResponse structure represents the SubscriptionName operation response
type GetSubscriptionNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriptionName *oaut.String   `idl:"name:pbstrSubscriptionName" json:"subscription_name"`
	// Return: The SubscriptionName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubscriptionNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriptionNameOperation) *xxx_GetSubscriptionNameOperation {
	if op == nil {
		op = &xxx_GetSubscriptionNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SubscriptionName = o.SubscriptionName
	op.Return = o.Return
	return op
}

func (o *GetSubscriptionNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriptionNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SubscriptionName = op.SubscriptionName
	o.Return = op.Return
}
func (o *GetSubscriptionNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSubscriptionNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriptionNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubscriptionNameOperation structure represents the SubscriptionName operation
type xxx_SetSubscriptionNameOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriptionName *oaut.String   `idl:"name:bstrSubscriptionName" json:"subscription_name"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubscriptionNameOperation) OpNum() int { return 10 }

func (o *xxx_SetSubscriptionNameOperation) OpName() string {
	return "/IEventSubscription/v0/SubscriptionName"
}

func (o *xxx_SetSubscriptionNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriptionNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrSubscriptionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SubscriptionName != nil {
			_ptr_bstrSubscriptionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriptionName != nil {
					if err := o.SubscriptionName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriptionName, _ptr_bstrSubscriptionName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriptionNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrSubscriptionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSubscriptionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriptionName == nil {
				o.SubscriptionName = &oaut.String{}
			}
			if err := o.SubscriptionName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSubscriptionName := func(ptr interface{}) { o.SubscriptionName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SubscriptionName, _s_bstrSubscriptionName, _ptr_bstrSubscriptionName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriptionNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriptionNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriptionNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSubscriptionNameRequest structure represents the SubscriptionName operation request
type SetSubscriptionNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	SubscriptionName *oaut.String   `idl:"name:bstrSubscriptionName" json:"subscription_name"`
}

func (o *SetSubscriptionNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSubscriptionNameOperation) *xxx_SetSubscriptionNameOperation {
	if op == nil {
		op = &xxx_SetSubscriptionNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SubscriptionName = o.SubscriptionName
	return op
}

func (o *SetSubscriptionNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriptionNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SubscriptionName = op.SubscriptionName
}
func (o *SetSubscriptionNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSubscriptionNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriptionNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubscriptionNameResponse structure represents the SubscriptionName operation response
type SetSubscriptionNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SubscriptionName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSubscriptionNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSubscriptionNameOperation) *xxx_SetSubscriptionNameOperation {
	if op == nil {
		op = &xxx_SetSubscriptionNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSubscriptionNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriptionNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSubscriptionNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSubscriptionNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriptionNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPublisherIDOperation structure represents the PublisherID operation
type xxx_GetPublisherIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	PublisherID *oaut.String   `idl:"name:pbstrPublisherID" json:"publisher_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPublisherIDOperation) OpNum() int { return 11 }

func (o *xxx_GetPublisherIDOperation) OpName() string { return "/IEventSubscription/v0/PublisherID" }

func (o *xxx_GetPublisherIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrPublisherID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PublisherID != nil {
			_ptr_pbstrPublisherID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PublisherID != nil {
					if err := o.PublisherID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PublisherID, _ptr_pbstrPublisherID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrPublisherID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPublisherID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PublisherID == nil {
				o.PublisherID = &oaut.String{}
			}
			if err := o.PublisherID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPublisherID := func(ptr interface{}) { o.PublisherID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PublisherID, _s_pbstrPublisherID, _ptr_pbstrPublisherID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPublisherIDRequest structure represents the PublisherID operation request
type GetPublisherIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPublisherIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherIDOperation) *xxx_GetPublisherIDOperation {
	if op == nil {
		op = &xxx_GetPublisherIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPublisherIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPublisherIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPublisherIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPublisherIDResponse structure represents the PublisherID operation response
type GetPublisherIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	PublisherID *oaut.String   `idl:"name:pbstrPublisherID" json:"publisher_id"`
	// Return: The PublisherID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPublisherIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherIDOperation) *xxx_GetPublisherIDOperation {
	if op == nil {
		op = &xxx_GetPublisherIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PublisherID = o.PublisherID
	op.Return = o.Return
	return op
}

func (o *GetPublisherIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PublisherID = op.PublisherID
	o.Return = op.Return
}
func (o *GetPublisherIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPublisherIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPublisherIDOperation structure represents the PublisherID operation
type xxx_SetPublisherIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	PublisherID *oaut.String   `idl:"name:bstrPublisherID" json:"publisher_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPublisherIDOperation) OpNum() int { return 12 }

func (o *xxx_SetPublisherIDOperation) OpName() string { return "/IEventSubscription/v0/PublisherID" }

func (o *xxx_SetPublisherIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPublisherIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrPublisherID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PublisherID != nil {
			_ptr_bstrPublisherID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PublisherID != nil {
					if err := o.PublisherID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PublisherID, _ptr_bstrPublisherID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPublisherIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPublisherID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPublisherID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PublisherID == nil {
				o.PublisherID = &oaut.String{}
			}
			if err := o.PublisherID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPublisherID := func(ptr interface{}) { o.PublisherID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PublisherID, _s_bstrPublisherID, _ptr_bstrPublisherID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPublisherIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPublisherIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPublisherIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetPublisherIDRequest structure represents the PublisherID operation request
type SetPublisherIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	PublisherID *oaut.String   `idl:"name:bstrPublisherID" json:"publisher_id"`
}

func (o *SetPublisherIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPublisherIDOperation) *xxx_SetPublisherIDOperation {
	if op == nil {
		op = &xxx_SetPublisherIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PublisherID = o.PublisherID
	return op
}

func (o *SetPublisherIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPublisherIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PublisherID = op.PublisherID
}
func (o *SetPublisherIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPublisherIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPublisherIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPublisherIDResponse structure represents the PublisherID operation response
type SetPublisherIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PublisherID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPublisherIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPublisherIDOperation) *xxx_SetPublisherIDOperation {
	if op == nil {
		op = &xxx_SetPublisherIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPublisherIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPublisherIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPublisherIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPublisherIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPublisherIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEventClassIDOperation structure represents the EventClassID operation
type xxx_GetEventClassIDOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassID *oaut.String   `idl:"name:pbstrEventClassID" json:"event_class_id"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventClassIDOperation) OpNum() int { return 13 }

func (o *xxx_GetEventClassIDOperation) OpName() string { return "/IEventSubscription/v0/EventClassID" }

func (o *xxx_GetEventClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrEventClassID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassID != nil {
			_ptr_pbstrEventClassID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassID != nil {
					if err := o.EventClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassID, _ptr_pbstrEventClassID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrEventClassID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrEventClassID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassID == nil {
				o.EventClassID = &oaut.String{}
			}
			if err := o.EventClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrEventClassID := func(ptr interface{}) { o.EventClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassID, _s_pbstrEventClassID, _ptr_pbstrEventClassID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetEventClassIDRequest structure represents the EventClassID operation request
type GetEventClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEventClassIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEventClassIDOperation) *xxx_GetEventClassIDOperation {
	if op == nil {
		op = &xxx_GetEventClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetEventClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEventClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventClassIDResponse structure represents the EventClassID operation response
type GetEventClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassID *oaut.String   `idl:"name:pbstrEventClassID" json:"event_class_id"`
	// Return: The EventClassID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventClassIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEventClassIDOperation) *xxx_GetEventClassIDOperation {
	if op == nil {
		op = &xxx_GetEventClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.EventClassID = o.EventClassID
	op.Return = o.Return
	return op
}

func (o *GetEventClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EventClassID = op.EventClassID
	o.Return = op.Return
}
func (o *GetEventClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEventClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEventClassIDOperation structure represents the EventClassID operation
type xxx_SetEventClassIDOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassID *oaut.String   `idl:"name:bstrEventClassID" json:"event_class_id"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEventClassIDOperation) OpNum() int { return 14 }

func (o *xxx_SetEventClassIDOperation) OpName() string { return "/IEventSubscription/v0/EventClassID" }

func (o *xxx_SetEventClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrEventClassID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassID != nil {
			_ptr_bstrEventClassID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassID != nil {
					if err := o.EventClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassID, _ptr_bstrEventClassID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrEventClassID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrEventClassID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassID == nil {
				o.EventClassID = &oaut.String{}
			}
			if err := o.EventClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrEventClassID := func(ptr interface{}) { o.EventClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassID, _s_bstrEventClassID, _ptr_bstrEventClassID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetEventClassIDRequest structure represents the EventClassID operation request
type SetEventClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	EventClassID *oaut.String   `idl:"name:bstrEventClassID" json:"event_class_id"`
}

func (o *SetEventClassIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetEventClassIDOperation) *xxx_SetEventClassIDOperation {
	if op == nil {
		op = &xxx_SetEventClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.EventClassID = o.EventClassID
	return op
}

func (o *SetEventClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EventClassID = op.EventClassID
}
func (o *SetEventClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetEventClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEventClassIDResponse structure represents the EventClassID operation response
type SetEventClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EventClassID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEventClassIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetEventClassIDOperation) *xxx_SetEventClassIDOperation {
	if op == nil {
		op = &xxx_SetEventClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetEventClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEventClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetEventClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMethodNameOperation structure represents the MethodName operation
type xxx_GetMethodNameOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	MethodName *oaut.String   `idl:"name:pbstrMethodName" json:"method_name"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMethodNameOperation) OpNum() int { return 15 }

func (o *xxx_GetMethodNameOperation) OpName() string { return "/IEventSubscription/v0/MethodName" }

func (o *xxx_GetMethodNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMethodNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMethodNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMethodNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMethodNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrMethodName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MethodName != nil {
			_ptr_pbstrMethodName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MethodName != nil {
					if err := o.MethodName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MethodName, _ptr_pbstrMethodName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMethodNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrMethodName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrMethodName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MethodName == nil {
				o.MethodName = &oaut.String{}
			}
			if err := o.MethodName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrMethodName := func(ptr interface{}) { o.MethodName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MethodName, _s_pbstrMethodName, _ptr_pbstrMethodName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetMethodNameRequest structure represents the MethodName operation request
type GetMethodNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMethodNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMethodNameOperation) *xxx_GetMethodNameOperation {
	if op == nil {
		op = &xxx_GetMethodNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMethodNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMethodNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMethodNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMethodNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMethodNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMethodNameResponse structure represents the MethodName operation response
type GetMethodNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	MethodName *oaut.String   `idl:"name:pbstrMethodName" json:"method_name"`
	// Return: The MethodName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMethodNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMethodNameOperation) *xxx_GetMethodNameOperation {
	if op == nil {
		op = &xxx_GetMethodNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MethodName = o.MethodName
	op.Return = o.Return
	return op
}

func (o *GetMethodNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMethodNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MethodName = op.MethodName
	o.Return = op.Return
}
func (o *GetMethodNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMethodNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMethodNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMethodNameOperation structure represents the MethodName operation
type xxx_SetMethodNameOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	MethodName *oaut.String   `idl:"name:bstrMethodName" json:"method_name"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMethodNameOperation) OpNum() int { return 16 }

func (o *xxx_SetMethodNameOperation) OpName() string { return "/IEventSubscription/v0/MethodName" }

func (o *xxx_SetMethodNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMethodNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrMethodName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MethodName != nil {
			_ptr_bstrMethodName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MethodName != nil {
					if err := o.MethodName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MethodName, _ptr_bstrMethodName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMethodNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMethodName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMethodName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MethodName == nil {
				o.MethodName = &oaut.String{}
			}
			if err := o.MethodName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMethodName := func(ptr interface{}) { o.MethodName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MethodName, _s_bstrMethodName, _ptr_bstrMethodName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMethodNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMethodNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMethodNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetMethodNameRequest structure represents the MethodName operation request
type SetMethodNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	MethodName *oaut.String   `idl:"name:bstrMethodName" json:"method_name"`
}

func (o *SetMethodNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMethodNameOperation) *xxx_SetMethodNameOperation {
	if op == nil {
		op = &xxx_SetMethodNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MethodName = o.MethodName
	return op
}

func (o *SetMethodNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMethodNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MethodName = op.MethodName
}
func (o *SetMethodNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMethodNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMethodNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMethodNameResponse structure represents the MethodName operation response
type SetMethodNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MethodName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMethodNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMethodNameOperation) *xxx_SetMethodNameOperation {
	if op == nil {
		op = &xxx_SetMethodNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMethodNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMethodNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMethodNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMethodNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMethodNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubscriberClassIDOperation structure represents the SubscriberCLSID operation
type xxx_GetSubscriberClassIDOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberClassID *oaut.String   `idl:"name:pbstrSubscriberCLSID" json:"subscriber_class_id"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubscriberClassIDOperation) OpNum() int { return 17 }

func (o *xxx_GetSubscriberClassIDOperation) OpName() string {
	return "/IEventSubscription/v0/SubscriberCLSID"
}

func (o *xxx_GetSubscriberClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrSubscriberCLSID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SubscriberClassID != nil {
			_ptr_pbstrSubscriberCLSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriberClassID != nil {
					if err := o.SubscriberClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriberClassID, _ptr_pbstrSubscriberCLSID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrSubscriberCLSID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrSubscriberCLSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriberClassID == nil {
				o.SubscriberClassID = &oaut.String{}
			}
			if err := o.SubscriberClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrSubscriberCLSID := func(ptr interface{}) { o.SubscriberClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SubscriberClassID, _s_pbstrSubscriberCLSID, _ptr_pbstrSubscriberCLSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSubscriberClassIDRequest structure represents the SubscriberCLSID operation request
type GetSubscriberClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubscriberClassIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriberClassIDOperation) *xxx_GetSubscriberClassIDOperation {
	if op == nil {
		op = &xxx_GetSubscriberClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSubscriberClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubscriberClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSubscriberClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubscriberClassIDResponse structure represents the SubscriberCLSID operation response
type GetSubscriberClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberClassID *oaut.String   `idl:"name:pbstrSubscriberCLSID" json:"subscriber_class_id"`
	// Return: The SubscriberCLSID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubscriberClassIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriberClassIDOperation) *xxx_GetSubscriberClassIDOperation {
	if op == nil {
		op = &xxx_GetSubscriberClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SubscriberClassID = o.SubscriberClassID
	op.Return = o.Return
	return op
}

func (o *GetSubscriberClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SubscriberClassID = op.SubscriberClassID
	o.Return = op.Return
}
func (o *GetSubscriberClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSubscriberClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubscriberClassIDOperation structure represents the SubscriberCLSID operation
type xxx_SetSubscriberClassIDOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberClassID *oaut.String   `idl:"name:bstrSubscriberCLSID" json:"subscriber_class_id"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubscriberClassIDOperation) OpNum() int { return 18 }

func (o *xxx_SetSubscriberClassIDOperation) OpName() string {
	return "/IEventSubscription/v0/SubscriberCLSID"
}

func (o *xxx_SetSubscriberClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrSubscriberCLSID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SubscriberClassID != nil {
			_ptr_bstrSubscriberCLSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriberClassID != nil {
					if err := o.SubscriberClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriberClassID, _ptr_bstrSubscriberCLSID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrSubscriberCLSID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSubscriberCLSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriberClassID == nil {
				o.SubscriberClassID = &oaut.String{}
			}
			if err := o.SubscriberClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSubscriberCLSID := func(ptr interface{}) { o.SubscriberClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SubscriberClassID, _s_bstrSubscriberCLSID, _ptr_bstrSubscriberCLSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSubscriberClassIDRequest structure represents the SubscriberCLSID operation request
type SetSubscriberClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	SubscriberClassID *oaut.String   `idl:"name:bstrSubscriberCLSID" json:"subscriber_class_id"`
}

func (o *SetSubscriberClassIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSubscriberClassIDOperation) *xxx_SetSubscriberClassIDOperation {
	if op == nil {
		op = &xxx_SetSubscriberClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SubscriberClassID = o.SubscriberClassID
	return op
}

func (o *SetSubscriberClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriberClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SubscriberClassID = op.SubscriberClassID
}
func (o *SetSubscriberClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSubscriberClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriberClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubscriberClassIDResponse structure represents the SubscriberCLSID operation response
type SetSubscriberClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SubscriberCLSID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSubscriberClassIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSubscriberClassIDOperation) *xxx_SetSubscriberClassIDOperation {
	if op == nil {
		op = &xxx_SetSubscriberClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSubscriberClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriberClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSubscriberClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSubscriberClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriberClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubscriberInterfaceOperation structure represents the SubscriberInterface operation
type xxx_GetSubscriberInterfaceOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberInterface *dcom.Unknown  `idl:"name:ppSubscriberInterface" json:"subscriber_interface"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubscriberInterfaceOperation) OpNum() int { return 19 }

func (o *xxx_GetSubscriberInterfaceOperation) OpName() string {
	return "/IEventSubscription/v0/SubscriberInterface"
}

func (o *xxx_GetSubscriberInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppSubscriberInterface {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.SubscriberInterface != nil {
			_ptr_ppSubscriberInterface := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriberInterface != nil {
					if err := o.SubscriberInterface.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriberInterface, _ptr_ppSubscriberInterface); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppSubscriberInterface {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_ppSubscriberInterface := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriberInterface == nil {
				o.SubscriberInterface = &dcom.Unknown{}
			}
			if err := o.SubscriberInterface.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSubscriberInterface := func(ptr interface{}) { o.SubscriberInterface = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.SubscriberInterface, _s_ppSubscriberInterface, _ptr_ppSubscriberInterface); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSubscriberInterfaceRequest structure represents the SubscriberInterface operation request
type GetSubscriberInterfaceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubscriberInterfaceRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriberInterfaceOperation) *xxx_GetSubscriberInterfaceOperation {
	if op == nil {
		op = &xxx_GetSubscriberInterfaceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSubscriberInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberInterfaceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubscriberInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSubscriberInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubscriberInterfaceResponse structure represents the SubscriberInterface operation response
type GetSubscriberInterfaceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberInterface *dcom.Unknown  `idl:"name:ppSubscriberInterface" json:"subscriber_interface"`
	// Return: The SubscriberInterface return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubscriberInterfaceResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriberInterfaceOperation) *xxx_GetSubscriberInterfaceOperation {
	if op == nil {
		op = &xxx_GetSubscriberInterfaceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SubscriberInterface = o.SubscriberInterface
	op.Return = o.Return
	return op
}

func (o *GetSubscriberInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberInterfaceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SubscriberInterface = op.SubscriberInterface
	o.Return = op.Return
}
func (o *GetSubscriberInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSubscriberInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubscriberInterfaceOperation structure represents the SubscriberInterface operation
type xxx_SetSubscriberInterfaceOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberInterface *dcom.Unknown  `idl:"name:pSubscriberInterface" json:"subscriber_interface"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubscriberInterfaceOperation) OpNum() int { return 20 }

func (o *xxx_SetSubscriberInterfaceOperation) OpName() string {
	return "/IEventSubscription/v0/SubscriberInterface"
}

func (o *xxx_SetSubscriberInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pSubscriberInterface {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.SubscriberInterface != nil {
			_ptr_pSubscriberInterface := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriberInterface != nil {
					if err := o.SubscriberInterface.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriberInterface, _ptr_pSubscriberInterface); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pSubscriberInterface {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_pSubscriberInterface := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriberInterface == nil {
				o.SubscriberInterface = &dcom.Unknown{}
			}
			if err := o.SubscriberInterface.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pSubscriberInterface := func(ptr interface{}) { o.SubscriberInterface = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.SubscriberInterface, _s_pSubscriberInterface, _ptr_pSubscriberInterface); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSubscriberInterfaceRequest structure represents the SubscriberInterface operation request
type SetSubscriberInterfaceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	SubscriberInterface *dcom.Unknown  `idl:"name:pSubscriberInterface" json:"subscriber_interface"`
}

func (o *SetSubscriberInterfaceRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSubscriberInterfaceOperation) *xxx_SetSubscriberInterfaceOperation {
	if op == nil {
		op = &xxx_SetSubscriberInterfaceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SubscriberInterface = o.SubscriberInterface
	return op
}

func (o *SetSubscriberInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriberInterfaceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SubscriberInterface = op.SubscriberInterface
}
func (o *SetSubscriberInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSubscriberInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriberInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubscriberInterfaceResponse structure represents the SubscriberInterface operation response
type SetSubscriberInterfaceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SubscriberInterface return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSubscriberInterfaceResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSubscriberInterfaceOperation) *xxx_SetSubscriberInterfaceOperation {
	if op == nil {
		op = &xxx_SetSubscriberInterfaceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSubscriberInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriberInterfaceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSubscriberInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSubscriberInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriberInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPerUserOperation structure represents the PerUser operation
type xxx_GetPerUserOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	PerUser bool           `idl:"name:pfPerUser" json:"per_user"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPerUserOperation) OpNum() int { return 21 }

func (o *xxx_GetPerUserOperation) OpName() string { return "/IEventSubscription/v0/PerUser" }

func (o *xxx_GetPerUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pfPerUser {out, retval} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.PerUser {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pfPerUser {out, retval} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bPerUser int32
		if err := w.ReadData(&_bPerUser); err != nil {
			return err
		}
		o.PerUser = _bPerUser != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPerUserRequest structure represents the PerUser operation request
type GetPerUserRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPerUserRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPerUserOperation) *xxx_GetPerUserOperation {
	if op == nil {
		op = &xxx_GetPerUserOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPerUserRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPerUserOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPerUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPerUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPerUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPerUserResponse structure represents the PerUser operation response
type GetPerUserResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	PerUser bool           `idl:"name:pfPerUser" json:"per_user"`
	// Return: The PerUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPerUserResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPerUserOperation) *xxx_GetPerUserOperation {
	if op == nil {
		op = &xxx_GetPerUserOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PerUser = o.PerUser
	op.Return = o.Return
	return op
}

func (o *GetPerUserResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPerUserOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PerUser = op.PerUser
	o.Return = op.Return
}
func (o *GetPerUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPerUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPerUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPerUserOperation structure represents the PerUser operation
type xxx_SetPerUserOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	PerUser bool           `idl:"name:fPerUser" json:"per_user"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPerUserOperation) OpNum() int { return 22 }

func (o *xxx_SetPerUserOperation) OpName() string { return "/IEventSubscription/v0/PerUser" }

func (o *xxx_SetPerUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPerUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// fPerUser {in} (1:{alias=BOOL}(int32))
	{
		if !o.PerUser {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetPerUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fPerUser {in} (1:{alias=BOOL}(int32))
	{
		var _bPerUser int32
		if err := w.ReadData(&_bPerUser); err != nil {
			return err
		}
		o.PerUser = _bPerUser != 0
	}
	return nil
}

func (o *xxx_SetPerUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPerUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPerUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetPerUserRequest structure represents the PerUser operation request
type SetPerUserRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	PerUser bool           `idl:"name:fPerUser" json:"per_user"`
}

func (o *SetPerUserRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPerUserOperation) *xxx_SetPerUserOperation {
	if op == nil {
		op = &xxx_SetPerUserOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PerUser = o.PerUser
	return op
}

func (o *SetPerUserRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPerUserOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PerUser = op.PerUser
}
func (o *SetPerUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPerUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPerUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPerUserResponse structure represents the PerUser operation response
type SetPerUserResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PerUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPerUserResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPerUserOperation) *xxx_SetPerUserOperation {
	if op == nil {
		op = &xxx_SetPerUserOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPerUserResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPerUserOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPerUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPerUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPerUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetOwnerSIDOperation structure represents the OwnerSID operation
type xxx_GetOwnerSIDOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	OwnerSID *oaut.String   `idl:"name:pbstrOwnerSID" json:"owner_sid"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOwnerSIDOperation) OpNum() int { return 23 }

func (o *xxx_GetOwnerSIDOperation) OpName() string { return "/IEventSubscription/v0/OwnerSID" }

func (o *xxx_GetOwnerSIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOwnerSIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOwnerSIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOwnerSIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOwnerSIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrOwnerSID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OwnerSID != nil {
			_ptr_pbstrOwnerSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OwnerSID != nil {
					if err := o.OwnerSID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OwnerSID, _ptr_pbstrOwnerSID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOwnerSIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrOwnerSID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrOwnerSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OwnerSID == nil {
				o.OwnerSID = &oaut.String{}
			}
			if err := o.OwnerSID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrOwnerSID := func(ptr interface{}) { o.OwnerSID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OwnerSID, _s_pbstrOwnerSID, _ptr_pbstrOwnerSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetOwnerSIDRequest structure represents the OwnerSID operation request
type GetOwnerSIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetOwnerSIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetOwnerSIDOperation) *xxx_GetOwnerSIDOperation {
	if op == nil {
		op = &xxx_GetOwnerSIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetOwnerSIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOwnerSIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetOwnerSIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetOwnerSIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOwnerSIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOwnerSIDResponse structure represents the OwnerSID operation response
type GetOwnerSIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	OwnerSID *oaut.String   `idl:"name:pbstrOwnerSID" json:"owner_sid"`
	// Return: The OwnerSID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetOwnerSIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetOwnerSIDOperation) *xxx_GetOwnerSIDOperation {
	if op == nil {
		op = &xxx_GetOwnerSIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.OwnerSID = o.OwnerSID
	op.Return = o.Return
	return op
}

func (o *GetOwnerSIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOwnerSIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.OwnerSID = op.OwnerSID
	o.Return = op.Return
}
func (o *GetOwnerSIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetOwnerSIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOwnerSIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOwnerSIDOperation structure represents the OwnerSID operation
type xxx_SetOwnerSIDOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	OwnerSID *oaut.String   `idl:"name:bstrOwnerSID" json:"owner_sid"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOwnerSIDOperation) OpNum() int { return 24 }

func (o *xxx_SetOwnerSIDOperation) OpName() string { return "/IEventSubscription/v0/OwnerSID" }

func (o *xxx_SetOwnerSIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOwnerSIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrOwnerSID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OwnerSID != nil {
			_ptr_bstrOwnerSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OwnerSID != nil {
					if err := o.OwnerSID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OwnerSID, _ptr_bstrOwnerSID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOwnerSIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrOwnerSID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrOwnerSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OwnerSID == nil {
				o.OwnerSID = &oaut.String{}
			}
			if err := o.OwnerSID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrOwnerSID := func(ptr interface{}) { o.OwnerSID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OwnerSID, _s_bstrOwnerSID, _ptr_bstrOwnerSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOwnerSIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOwnerSIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOwnerSIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetOwnerSIDRequest structure represents the OwnerSID operation request
type SetOwnerSIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	OwnerSID *oaut.String   `idl:"name:bstrOwnerSID" json:"owner_sid"`
}

func (o *SetOwnerSIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetOwnerSIDOperation) *xxx_SetOwnerSIDOperation {
	if op == nil {
		op = &xxx_SetOwnerSIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.OwnerSID = o.OwnerSID
	return op
}

func (o *SetOwnerSIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOwnerSIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.OwnerSID = op.OwnerSID
}
func (o *SetOwnerSIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetOwnerSIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOwnerSIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOwnerSIDResponse structure represents the OwnerSID operation response
type SetOwnerSIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OwnerSID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetOwnerSIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetOwnerSIDOperation) *xxx_SetOwnerSIDOperation {
	if op == nil {
		op = &xxx_SetOwnerSIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetOwnerSIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOwnerSIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetOwnerSIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetOwnerSIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOwnerSIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEnabledOperation structure represents the Enabled operation
type xxx_GetEnabledOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled bool           `idl:"name:pfEnabled" json:"enabled"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEnabledOperation) OpNum() int { return 25 }

func (o *xxx_GetEnabledOperation) OpName() string { return "/IEventSubscription/v0/Enabled" }

func (o *xxx_GetEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pfEnabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.Enabled {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pfEnabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bEnabled int32
		if err := w.ReadData(&_bEnabled); err != nil {
			return err
		}
		o.Enabled = _bEnabled != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetEnabledRequest structure represents the Enabled operation request
type GetEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEnabledOperation) *xxx_GetEnabledOperation {
	if op == nil {
		op = &xxx_GetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEnabledResponse structure represents the Enabled operation response
type GetEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled bool           `idl:"name:pfEnabled" json:"enabled"`
	// Return: The Enabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEnabledOperation) *xxx_GetEnabledOperation {
	if op == nil {
		op = &xxx_GetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enabled = o.Enabled
	op.Return = o.Return
	return op
}

func (o *GetEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enabled = op.Enabled
	o.Return = op.Return
}
func (o *GetEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEnabledOperation structure represents the Enabled operation
type xxx_SetEnabledOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled bool           `idl:"name:fEnabled" json:"enabled"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEnabledOperation) OpNum() int { return 26 }

func (o *xxx_SetEnabledOperation) OpName() string { return "/IEventSubscription/v0/Enabled" }

func (o *xxx_SetEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// fEnabled {in} (1:{alias=BOOL}(int32))
	{
		if !o.Enabled {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fEnabled {in} (1:{alias=BOOL}(int32))
	{
		var _bEnabled int32
		if err := w.ReadData(&_bEnabled); err != nil {
			return err
		}
		o.Enabled = _bEnabled != 0
	}
	return nil
}

func (o *xxx_SetEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetEnabledRequest structure represents the Enabled operation request
type SetEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Enabled bool           `idl:"name:fEnabled" json:"enabled"`
}

func (o *SetEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_SetEnabledOperation) *xxx_SetEnabledOperation {
	if op == nil {
		op = &xxx_SetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Enabled = o.Enabled
	return op
}

func (o *SetEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Enabled = op.Enabled
}
func (o *SetEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEnabledResponse structure represents the Enabled operation response
type SetEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Enabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_SetEnabledOperation) *xxx_SetEnabledOperation {
	if op == nil {
		op = &xxx_SetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDescriptionOperation structure represents the Description operation
type xxx_GetDescriptionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Description *oaut.String   `idl:"name:pbstrDescription" json:"description"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDescriptionOperation) OpNum() int { return 27 }

func (o *xxx_GetDescriptionOperation) OpName() string { return "/IEventSubscription/v0/Description" }

func (o *xxx_GetDescriptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrDescription {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Description != nil {
			_ptr_pbstrDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Description != nil {
					if err := o.Description.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Description, _ptr_pbstrDescription); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrDescription {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Description == nil {
				o.Description = &oaut.String{}
			}
			if err := o.Description.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrDescription := func(ptr interface{}) { o.Description = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Description, _s_pbstrDescription, _ptr_pbstrDescription); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDescriptionRequest structure represents the Description operation request
type GetDescriptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDescriptionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDescriptionOperation) *xxx_GetDescriptionOperation {
	if op == nil {
		op = &xxx_GetDescriptionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetDescriptionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDescriptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDescriptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDescriptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDescriptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDescriptionResponse structure represents the Description operation response
type GetDescriptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Description *oaut.String   `idl:"name:pbstrDescription" json:"description"`
	// Return: The Description return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDescriptionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDescriptionOperation) *xxx_GetDescriptionOperation {
	if op == nil {
		op = &xxx_GetDescriptionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Description = o.Description
	op.Return = o.Return
	return op
}

func (o *GetDescriptionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDescriptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Description = op.Description
	o.Return = op.Return
}
func (o *GetDescriptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDescriptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDescriptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDescriptionOperation structure represents the Description operation
type xxx_SetDescriptionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Description *oaut.String   `idl:"name:bstrDescription" json:"description"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDescriptionOperation) OpNum() int { return 28 }

func (o *xxx_SetDescriptionOperation) OpName() string { return "/IEventSubscription/v0/Description" }

func (o *xxx_SetDescriptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrDescription {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Description != nil {
			_ptr_bstrDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Description != nil {
					if err := o.Description.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Description, _ptr_bstrDescription); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrDescription {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Description == nil {
				o.Description = &oaut.String{}
			}
			if err := o.Description.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrDescription := func(ptr interface{}) { o.Description = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Description, _s_bstrDescription, _ptr_bstrDescription); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetDescriptionRequest structure represents the Description operation request
type SetDescriptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	Description *oaut.String   `idl:"name:bstrDescription" json:"description"`
}

func (o *SetDescriptionRequest) xxx_ToOp(ctx context.Context, op *xxx_SetDescriptionOperation) *xxx_SetDescriptionOperation {
	if op == nil {
		op = &xxx_SetDescriptionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Description = o.Description
	return op
}

func (o *SetDescriptionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDescriptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Description = op.Description
}
func (o *SetDescriptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetDescriptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDescriptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDescriptionResponse structure represents the Description operation response
type SetDescriptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Description return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDescriptionResponse) xxx_ToOp(ctx context.Context, op *xxx_SetDescriptionOperation) *xxx_SetDescriptionOperation {
	if op == nil {
		op = &xxx_SetDescriptionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetDescriptionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDescriptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDescriptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetDescriptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDescriptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMachineNameOperation structure represents the MachineName operation
type xxx_GetMachineNameOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineName *oaut.String   `idl:"name:pbstrMachineName" json:"machine_name"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMachineNameOperation) OpNum() int { return 29 }

func (o *xxx_GetMachineNameOperation) OpName() string { return "/IEventSubscription/v0/MachineName" }

func (o *xxx_GetMachineNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrMachineName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineName != nil {
			_ptr_pbstrMachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineName != nil {
					if err := o.MachineName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineName, _ptr_pbstrMachineName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrMachineName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrMachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineName == nil {
				o.MachineName = &oaut.String{}
			}
			if err := o.MachineName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrMachineName := func(ptr interface{}) { o.MachineName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineName, _s_pbstrMachineName, _ptr_pbstrMachineName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetMachineNameRequest structure represents the MachineName operation request
type GetMachineNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMachineNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMachineNameOperation) *xxx_GetMachineNameOperation {
	if op == nil {
		op = &xxx_GetMachineNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMachineNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMachineNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMachineNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMachineNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMachineNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMachineNameResponse structure represents the MachineName operation response
type GetMachineNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineName *oaut.String   `idl:"name:pbstrMachineName" json:"machine_name"`
	// Return: The MachineName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMachineNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMachineNameOperation) *xxx_GetMachineNameOperation {
	if op == nil {
		op = &xxx_GetMachineNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MachineName = o.MachineName
	op.Return = o.Return
	return op
}

func (o *GetMachineNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMachineNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MachineName = op.MachineName
	o.Return = op.Return
}
func (o *GetMachineNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMachineNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMachineNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMachineNameOperation structure represents the MachineName operation
type xxx_SetMachineNameOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineName *oaut.String   `idl:"name:bstrMachineName" json:"machine_name"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMachineNameOperation) OpNum() int { return 30 }

func (o *xxx_SetMachineNameOperation) OpName() string { return "/IEventSubscription/v0/MachineName" }

func (o *xxx_SetMachineNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMachineNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineName != nil {
			_ptr_bstrMachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineName != nil {
					if err := o.MachineName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineName, _ptr_bstrMachineName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMachineNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineName == nil {
				o.MachineName = &oaut.String{}
			}
			if err := o.MachineName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineName := func(ptr interface{}) { o.MachineName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineName, _s_bstrMachineName, _ptr_bstrMachineName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMachineNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMachineNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMachineNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetMachineNameRequest structure represents the MachineName operation request
type SetMachineNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	MachineName *oaut.String   `idl:"name:bstrMachineName" json:"machine_name"`
}

func (o *SetMachineNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMachineNameOperation) *xxx_SetMachineNameOperation {
	if op == nil {
		op = &xxx_SetMachineNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MachineName = o.MachineName
	return op
}

func (o *SetMachineNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMachineNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MachineName = op.MachineName
}
func (o *SetMachineNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMachineNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMachineNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMachineNameResponse structure represents the MachineName operation response
type SetMachineNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MachineName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMachineNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMachineNameOperation) *xxx_SetMachineNameOperation {
	if op == nil {
		op = &xxx_SetMachineNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMachineNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMachineNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMachineNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMachineNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMachineNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPublisherPropertyOperation structure represents the GetPublisherProperty operation
type xxx_GetPublisherPropertyOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	PropertyName  *oaut.String   `idl:"name:bstrPropertyName" json:"property_name"`
	PropertyValue *oaut.Variant  `idl:"name:propertyValue" json:"property_value"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPublisherPropertyOperation) OpNum() int { return 31 }

func (o *xxx_GetPublisherPropertyOperation) OpName() string {
	return "/IEventSubscription/v0/GetPublisherProperty"
}

func (o *xxx_GetPublisherPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PropertyName != nil {
			_ptr_bstrPropertyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyName != nil {
					if err := o.PropertyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyName, _ptr_bstrPropertyName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPropertyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyName == nil {
				o.PropertyName = &oaut.String{}
			}
			if err := o.PropertyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPropertyName := func(ptr interface{}) { o.PropertyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PropertyName, _s_bstrPropertyName, _ptr_bstrPropertyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// propertyValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.PropertyValue != nil {
			_ptr_propertyValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyValue != nil {
					if err := o.PropertyValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyValue, _ptr_propertyValue); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_propertyValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyValue == nil {
				o.PropertyValue = &oaut.Variant{}
			}
			if err := o.PropertyValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyValue := func(ptr interface{}) { o.PropertyValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.PropertyValue, _s_propertyValue, _ptr_propertyValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPublisherPropertyRequest structure represents the GetPublisherProperty operation request
type GetPublisherPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrPropertyName: The application-specific name for publisher property. The format
	// for the publisher property name MUST adhere to the format specified in section 2.2.2.1.
	PropertyName *oaut.String `idl:"name:bstrPropertyName" json:"property_name"`
}

func (o *GetPublisherPropertyRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherPropertyOperation) *xxx_GetPublisherPropertyOperation {
	if op == nil {
		op = &xxx_GetPublisherPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PropertyName = o.PropertyName
	return op
}

func (o *GetPublisherPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PropertyName = op.PropertyName
}
func (o *GetPublisherPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPublisherPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPublisherPropertyResponse structure represents the GetPublisherProperty operation response
type GetPublisherPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// propertyValue: If the function returns a success HRESULT, this MUST contain the application-specific
	// publisher property value which MUST be of the type specified in 2.2.2.2.
	PropertyValue *oaut.Variant `idl:"name:propertyValue" json:"property_value"`
	// Return: The GetPublisherProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPublisherPropertyResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherPropertyOperation) *xxx_GetPublisherPropertyOperation {
	if op == nil {
		op = &xxx_GetPublisherPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PropertyValue = o.PropertyValue
	op.Return = o.Return
	return op
}

func (o *GetPublisherPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyValue = op.PropertyValue
	o.Return = op.Return
}
func (o *GetPublisherPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPublisherPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PutPublisherPropertyOperation structure represents the PutPublisherProperty operation
type xxx_PutPublisherPropertyOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	PropertyName  *oaut.String   `idl:"name:bstrPropertyName" json:"property_name"`
	PropertyValue *oaut.Variant  `idl:"name:propertyValue" json:"property_value"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PutPublisherPropertyOperation) OpNum() int { return 32 }

func (o *xxx_PutPublisherPropertyOperation) OpName() string {
	return "/IEventSubscription/v0/PutPublisherProperty"
}

func (o *xxx_PutPublisherPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutPublisherPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PropertyName != nil {
			_ptr_bstrPropertyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyName != nil {
					if err := o.PropertyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyName, _ptr_bstrPropertyName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// propertyValue {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.PropertyValue != nil {
			_ptr_propertyValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyValue != nil {
					if err := o.PropertyValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyValue, _ptr_propertyValue); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutPublisherPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPropertyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyName == nil {
				o.PropertyName = &oaut.String{}
			}
			if err := o.PropertyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPropertyName := func(ptr interface{}) { o.PropertyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PropertyName, _s_bstrPropertyName, _ptr_bstrPropertyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyValue {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_propertyValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyValue == nil {
				o.PropertyValue = &oaut.Variant{}
			}
			if err := o.PropertyValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyValue := func(ptr interface{}) { o.PropertyValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.PropertyValue, _s_propertyValue, _ptr_propertyValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutPublisherPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutPublisherPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutPublisherPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PutPublisherPropertyRequest structure represents the PutPublisherProperty operation request
type PutPublisherPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrPropertyName: The application-specific name for publisher property. The format
	// for the publisher property name MUST adhere to the format specified in section 2.2.2.1.
	PropertyName *oaut.String `idl:"name:bstrPropertyName" json:"property_name"`
	// propertyValue: The application-specific publisher property value which MUST be of
	// the type specified in 2.2.2.2.
	PropertyValue *oaut.Variant `idl:"name:propertyValue" json:"property_value"`
}

func (o *PutPublisherPropertyRequest) xxx_ToOp(ctx context.Context, op *xxx_PutPublisherPropertyOperation) *xxx_PutPublisherPropertyOperation {
	if op == nil {
		op = &xxx_PutPublisherPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PropertyName = o.PropertyName
	op.PropertyValue = o.PropertyValue
	return op
}

func (o *PutPublisherPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_PutPublisherPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PropertyName = op.PropertyName
	o.PropertyValue = op.PropertyValue
}
func (o *PutPublisherPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PutPublisherPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutPublisherPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PutPublisherPropertyResponse structure represents the PutPublisherProperty operation response
type PutPublisherPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PutPublisherProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PutPublisherPropertyResponse) xxx_ToOp(ctx context.Context, op *xxx_PutPublisherPropertyOperation) *xxx_PutPublisherPropertyOperation {
	if op == nil {
		op = &xxx_PutPublisherPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *PutPublisherPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_PutPublisherPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PutPublisherPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PutPublisherPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutPublisherPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemovePublisherPropertyOperation structure represents the RemovePublisherProperty operation
type xxx_RemovePublisherPropertyOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	PropertyName *oaut.String   `idl:"name:bstrPropertyName" json:"property_name"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemovePublisherPropertyOperation) OpNum() int { return 33 }

func (o *xxx_RemovePublisherPropertyOperation) OpName() string {
	return "/IEventSubscription/v0/RemovePublisherProperty"
}

func (o *xxx_RemovePublisherPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePublisherPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PropertyName != nil {
			_ptr_bstrPropertyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyName != nil {
					if err := o.PropertyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyName, _ptr_bstrPropertyName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePublisherPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPropertyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyName == nil {
				o.PropertyName = &oaut.String{}
			}
			if err := o.PropertyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPropertyName := func(ptr interface{}) { o.PropertyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PropertyName, _s_bstrPropertyName, _ptr_bstrPropertyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePublisherPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePublisherPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePublisherPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemovePublisherPropertyRequest structure represents the RemovePublisherProperty operation request
type RemovePublisherPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrPropertyName: The application-specific name for the publisher property. The format
	// for the publisher property name MUST adhere to the format specified in section 2.2.2.1.
	PropertyName *oaut.String `idl:"name:bstrPropertyName" json:"property_name"`
}

func (o *RemovePublisherPropertyRequest) xxx_ToOp(ctx context.Context, op *xxx_RemovePublisherPropertyOperation) *xxx_RemovePublisherPropertyOperation {
	if op == nil {
		op = &xxx_RemovePublisherPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PropertyName = o.PropertyName
	return op
}

func (o *RemovePublisherPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_RemovePublisherPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PropertyName = op.PropertyName
}
func (o *RemovePublisherPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemovePublisherPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemovePublisherPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemovePublisherPropertyResponse structure represents the RemovePublisherProperty operation response
type RemovePublisherPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RemovePublisherProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemovePublisherPropertyResponse) xxx_ToOp(ctx context.Context, op *xxx_RemovePublisherPropertyOperation) *xxx_RemovePublisherPropertyOperation {
	if op == nil {
		op = &xxx_RemovePublisherPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RemovePublisherPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_RemovePublisherPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RemovePublisherPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemovePublisherPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemovePublisherPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPublisherPropertyCollectionOperation structure represents the GetPublisherPropertyCollection operation
type xxx_GetPublisherPropertyCollectionOperation struct {
	This       *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat               `idl:"name:That" json:"that"`
	Collection *comev.EventObjectCollection `idl:"name:collection" json:"collection"`
	Return     int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPublisherPropertyCollectionOperation) OpNum() int { return 34 }

func (o *xxx_GetPublisherPropertyCollectionOperation) OpName() string {
	return "/IEventSubscription/v0/GetPublisherPropertyCollection"
}

func (o *xxx_GetPublisherPropertyCollectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherPropertyCollectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherPropertyCollectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherPropertyCollectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherPropertyCollectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// collection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IEventObjectCollection}(interface))
	{
		if o.Collection != nil {
			_ptr_collection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collection != nil {
					if err := o.Collection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&comev.EventObjectCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collection, _ptr_collection); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherPropertyCollectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// collection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IEventObjectCollection}(interface))
	{
		_ptr_collection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collection == nil {
				o.Collection = &comev.EventObjectCollection{}
			}
			if err := o.Collection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_collection := func(ptr interface{}) { o.Collection = *ptr.(**comev.EventObjectCollection) }
		if err := w.ReadPointer(&o.Collection, _s_collection, _ptr_collection); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPublisherPropertyCollectionRequest structure represents the GetPublisherPropertyCollection operation request
type GetPublisherPropertyCollectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPublisherPropertyCollectionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherPropertyCollectionOperation) *xxx_GetPublisherPropertyCollectionOperation {
	if op == nil {
		op = &xxx_GetPublisherPropertyCollectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPublisherPropertyCollectionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherPropertyCollectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPublisherPropertyCollectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPublisherPropertyCollectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherPropertyCollectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPublisherPropertyCollectionResponse structure represents the GetPublisherPropertyCollection operation response
type GetPublisherPropertyCollectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// collection: If the function returns a success HRESULT, this MUST return an instance
	// of DCOM object supporting the IEventObjectCollection which MUST contain a collection
	// of application-specific publisher properties. These properties MUST conform to the
	// specification given in section 2.2.2.
	Collection *comev.EventObjectCollection `idl:"name:collection" json:"collection"`
	// Return: The GetPublisherPropertyCollection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPublisherPropertyCollectionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherPropertyCollectionOperation) *xxx_GetPublisherPropertyCollectionOperation {
	if op == nil {
		op = &xxx_GetPublisherPropertyCollectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Collection = o.Collection
	op.Return = o.Return
	return op
}

func (o *GetPublisherPropertyCollectionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherPropertyCollectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collection = op.Collection
	o.Return = op.Return
}
func (o *GetPublisherPropertyCollectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPublisherPropertyCollectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherPropertyCollectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubscriberPropertyOperation structure represents the GetSubscriberProperty operation
type xxx_GetSubscriberPropertyOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	PropertyName  *oaut.String   `idl:"name:bstrPropertyName" json:"property_name"`
	PropertyValue *oaut.Variant  `idl:"name:propertyValue" json:"property_value"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubscriberPropertyOperation) OpNum() int { return 35 }

func (o *xxx_GetSubscriberPropertyOperation) OpName() string {
	return "/IEventSubscription/v0/GetSubscriberProperty"
}

func (o *xxx_GetSubscriberPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PropertyName != nil {
			_ptr_bstrPropertyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyName != nil {
					if err := o.PropertyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyName, _ptr_bstrPropertyName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPropertyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyName == nil {
				o.PropertyName = &oaut.String{}
			}
			if err := o.PropertyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPropertyName := func(ptr interface{}) { o.PropertyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PropertyName, _s_bstrPropertyName, _ptr_bstrPropertyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// propertyValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.PropertyValue != nil {
			_ptr_propertyValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyValue != nil {
					if err := o.PropertyValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyValue, _ptr_propertyValue); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_propertyValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyValue == nil {
				o.PropertyValue = &oaut.Variant{}
			}
			if err := o.PropertyValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyValue := func(ptr interface{}) { o.PropertyValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.PropertyValue, _s_propertyValue, _ptr_propertyValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSubscriberPropertyRequest structure represents the GetSubscriberProperty operation request
type GetSubscriberPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrPropertyName: The application-specific name for the subscriber property. The
	// format for the subscriber property name MUST adhere to the format specified in section
	// 2.2.2.1.
	PropertyName *oaut.String `idl:"name:bstrPropertyName" json:"property_name"`
}

func (o *GetSubscriberPropertyRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriberPropertyOperation) *xxx_GetSubscriberPropertyOperation {
	if op == nil {
		op = &xxx_GetSubscriberPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PropertyName = o.PropertyName
	return op
}

func (o *GetSubscriberPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PropertyName = op.PropertyName
}
func (o *GetSubscriberPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSubscriberPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubscriberPropertyResponse structure represents the GetSubscriberProperty operation response
type GetSubscriberPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// propertyValue: If the function returns a success HRESULT, this MUST contain the application-specific
	// subscriber property value which MUST be of the type specified in 2.2.2.2.
	PropertyValue *oaut.Variant `idl:"name:propertyValue" json:"property_value"`
	// Return: The GetSubscriberProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubscriberPropertyResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriberPropertyOperation) *xxx_GetSubscriberPropertyOperation {
	if op == nil {
		op = &xxx_GetSubscriberPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PropertyValue = o.PropertyValue
	op.Return = o.Return
	return op
}

func (o *GetSubscriberPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyValue = op.PropertyValue
	o.Return = op.Return
}
func (o *GetSubscriberPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSubscriberPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PutSubscriberPropertyOperation structure represents the PutSubscriberProperty operation
type xxx_PutSubscriberPropertyOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	PropertyName  *oaut.String   `idl:"name:bstrPropertyName" json:"property_name"`
	PropertyValue *oaut.Variant  `idl:"name:propertyValue" json:"property_value"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PutSubscriberPropertyOperation) OpNum() int { return 36 }

func (o *xxx_PutSubscriberPropertyOperation) OpName() string {
	return "/IEventSubscription/v0/PutSubscriberProperty"
}

func (o *xxx_PutSubscriberPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutSubscriberPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PropertyName != nil {
			_ptr_bstrPropertyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyName != nil {
					if err := o.PropertyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyName, _ptr_bstrPropertyName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// propertyValue {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.PropertyValue != nil {
			_ptr_propertyValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyValue != nil {
					if err := o.PropertyValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyValue, _ptr_propertyValue); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutSubscriberPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPropertyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyName == nil {
				o.PropertyName = &oaut.String{}
			}
			if err := o.PropertyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPropertyName := func(ptr interface{}) { o.PropertyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PropertyName, _s_bstrPropertyName, _ptr_bstrPropertyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyValue {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_propertyValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyValue == nil {
				o.PropertyValue = &oaut.Variant{}
			}
			if err := o.PropertyValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyValue := func(ptr interface{}) { o.PropertyValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.PropertyValue, _s_propertyValue, _ptr_propertyValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutSubscriberPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutSubscriberPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutSubscriberPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PutSubscriberPropertyRequest structure represents the PutSubscriberProperty operation request
type PutSubscriberPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrPropertyName: The application-specific name for the subscriber property. The
	// format for the subscriber property name MUST adhere to the format specified in section
	// 2.2.2.1.
	PropertyName *oaut.String `idl:"name:bstrPropertyName" json:"property_name"`
	// propertyValue: The application-specific subscriber property value which MUST be of
	// the type specified in 2.2.2.2.
	PropertyValue *oaut.Variant `idl:"name:propertyValue" json:"property_value"`
}

func (o *PutSubscriberPropertyRequest) xxx_ToOp(ctx context.Context, op *xxx_PutSubscriberPropertyOperation) *xxx_PutSubscriberPropertyOperation {
	if op == nil {
		op = &xxx_PutSubscriberPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PropertyName = o.PropertyName
	op.PropertyValue = o.PropertyValue
	return op
}

func (o *PutSubscriberPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_PutSubscriberPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PropertyName = op.PropertyName
	o.PropertyValue = op.PropertyValue
}
func (o *PutSubscriberPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PutSubscriberPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutSubscriberPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PutSubscriberPropertyResponse structure represents the PutSubscriberProperty operation response
type PutSubscriberPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PutSubscriberProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PutSubscriberPropertyResponse) xxx_ToOp(ctx context.Context, op *xxx_PutSubscriberPropertyOperation) *xxx_PutSubscriberPropertyOperation {
	if op == nil {
		op = &xxx_PutSubscriberPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *PutSubscriberPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_PutSubscriberPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PutSubscriberPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PutSubscriberPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutSubscriberPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveSubscriberPropertyOperation structure represents the RemoveSubscriberProperty operation
type xxx_RemoveSubscriberPropertyOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	PropertyName *oaut.String   `idl:"name:bstrPropertyName" json:"property_name"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveSubscriberPropertyOperation) OpNum() int { return 37 }

func (o *xxx_RemoveSubscriberPropertyOperation) OpName() string {
	return "/IEventSubscription/v0/RemoveSubscriberProperty"
}

func (o *xxx_RemoveSubscriberPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubscriberPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PropertyName != nil {
			_ptr_bstrPropertyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyName != nil {
					if err := o.PropertyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyName, _ptr_bstrPropertyName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubscriberPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPropertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPropertyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyName == nil {
				o.PropertyName = &oaut.String{}
			}
			if err := o.PropertyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPropertyName := func(ptr interface{}) { o.PropertyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PropertyName, _s_bstrPropertyName, _ptr_bstrPropertyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubscriberPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubscriberPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubscriberPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveSubscriberPropertyRequest structure represents the RemoveSubscriberProperty operation request
type RemoveSubscriberPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrPropertyName: The application-specific name for the subscriber property. The
	// format for the subscriber property name MUST adhere to the format specified in section
	// 2.2.2.1.
	PropertyName *oaut.String `idl:"name:bstrPropertyName" json:"property_name"`
}

func (o *RemoveSubscriberPropertyRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoveSubscriberPropertyOperation) *xxx_RemoveSubscriberPropertyOperation {
	if op == nil {
		op = &xxx_RemoveSubscriberPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PropertyName = o.PropertyName
	return op
}

func (o *RemoveSubscriberPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveSubscriberPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PropertyName = op.PropertyName
}
func (o *RemoveSubscriberPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoveSubscriberPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveSubscriberPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveSubscriberPropertyResponse structure represents the RemoveSubscriberProperty operation response
type RemoveSubscriberPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RemoveSubscriberProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveSubscriberPropertyResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoveSubscriberPropertyOperation) *xxx_RemoveSubscriberPropertyOperation {
	if op == nil {
		op = &xxx_RemoveSubscriberPropertyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RemoveSubscriberPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveSubscriberPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RemoveSubscriberPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoveSubscriberPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveSubscriberPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubscriberPropertyCollectionOperation structure represents the GetSubscriberPropertyCollection operation
type xxx_GetSubscriberPropertyCollectionOperation struct {
	This       *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat               `idl:"name:That" json:"that"`
	Collection *comev.EventObjectCollection `idl:"name:collection" json:"collection"`
	Return     int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubscriberPropertyCollectionOperation) OpNum() int { return 38 }

func (o *xxx_GetSubscriberPropertyCollectionOperation) OpName() string {
	return "/IEventSubscription/v0/GetSubscriberPropertyCollection"
}

func (o *xxx_GetSubscriberPropertyCollectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPropertyCollectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPropertyCollectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPropertyCollectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPropertyCollectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// collection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IEventObjectCollection}(interface))
	{
		if o.Collection != nil {
			_ptr_collection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collection != nil {
					if err := o.Collection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&comev.EventObjectCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collection, _ptr_collection); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPropertyCollectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// collection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IEventObjectCollection}(interface))
	{
		_ptr_collection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collection == nil {
				o.Collection = &comev.EventObjectCollection{}
			}
			if err := o.Collection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_collection := func(ptr interface{}) { o.Collection = *ptr.(**comev.EventObjectCollection) }
		if err := w.ReadPointer(&o.Collection, _s_collection, _ptr_collection); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSubscriberPropertyCollectionRequest structure represents the GetSubscriberPropertyCollection operation request
type GetSubscriberPropertyCollectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubscriberPropertyCollectionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriberPropertyCollectionOperation) *xxx_GetSubscriberPropertyCollectionOperation {
	if op == nil {
		op = &xxx_GetSubscriberPropertyCollectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSubscriberPropertyCollectionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberPropertyCollectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubscriberPropertyCollectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSubscriberPropertyCollectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberPropertyCollectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubscriberPropertyCollectionResponse structure represents the GetSubscriberPropertyCollection operation response
type GetSubscriberPropertyCollectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// collection: If the function returns a success HRESULT, this MUST return an instance
	// of a DCOM object supporting the IEventObjectCollection which MUST contain a collection
	// of application-specific subscriber properties. These properties MUST conform to the
	// specification given in section 2.2.2.
	Collection *comev.EventObjectCollection `idl:"name:collection" json:"collection"`
	// Return: The GetSubscriberPropertyCollection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubscriberPropertyCollectionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSubscriberPropertyCollectionOperation) *xxx_GetSubscriberPropertyCollectionOperation {
	if op == nil {
		op = &xxx_GetSubscriberPropertyCollectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Collection = o.Collection
	op.Return = o.Return
	return op
}

func (o *GetSubscriberPropertyCollectionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberPropertyCollectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collection = op.Collection
	o.Return = op.Return
}
func (o *GetSubscriberPropertyCollectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSubscriberPropertyCollectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberPropertyCollectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetInterfaceIDOperation structure represents the InterfaceID operation
type xxx_GetInterfaceIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	InterfaceID *oaut.String   `idl:"name:pbstrInterfaceID" json:"interface_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInterfaceIDOperation) OpNum() int { return 39 }

func (o *xxx_GetInterfaceIDOperation) OpName() string { return "/IEventSubscription/v0/InterfaceID" }

func (o *xxx_GetInterfaceIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInterfaceIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInterfaceIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInterfaceIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInterfaceIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbstrInterfaceID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.InterfaceID != nil {
			_ptr_pbstrInterfaceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InterfaceID != nil {
					if err := o.InterfaceID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InterfaceID, _ptr_pbstrInterfaceID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInterfaceIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbstrInterfaceID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrInterfaceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InterfaceID == nil {
				o.InterfaceID = &oaut.String{}
			}
			if err := o.InterfaceID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrInterfaceID := func(ptr interface{}) { o.InterfaceID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.InterfaceID, _s_pbstrInterfaceID, _ptr_pbstrInterfaceID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetInterfaceIDRequest structure represents the InterfaceID operation request
type GetInterfaceIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetInterfaceIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetInterfaceIDOperation) *xxx_GetInterfaceIDOperation {
	if op == nil {
		op = &xxx_GetInterfaceIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetInterfaceIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInterfaceIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetInterfaceIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetInterfaceIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInterfaceIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInterfaceIDResponse structure represents the InterfaceID operation response
type GetInterfaceIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	InterfaceID *oaut.String   `idl:"name:pbstrInterfaceID" json:"interface_id"`
	// Return: The InterfaceID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetInterfaceIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetInterfaceIDOperation) *xxx_GetInterfaceIDOperation {
	if op == nil {
		op = &xxx_GetInterfaceIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.InterfaceID = o.InterfaceID
	op.Return = o.Return
	return op
}

func (o *GetInterfaceIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInterfaceIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.InterfaceID = op.InterfaceID
	o.Return = op.Return
}
func (o *GetInterfaceIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetInterfaceIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInterfaceIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInterfaceIDOperation structure represents the InterfaceID operation
type xxx_SetInterfaceIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	InterfaceID *oaut.String   `idl:"name:bstrInterfaceID" json:"interface_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInterfaceIDOperation) OpNum() int { return 40 }

func (o *xxx_SetInterfaceIDOperation) OpName() string { return "/IEventSubscription/v0/InterfaceID" }

func (o *xxx_SetInterfaceIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInterfaceIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// bstrInterfaceID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.InterfaceID != nil {
			_ptr_bstrInterfaceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InterfaceID != nil {
					if err := o.InterfaceID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InterfaceID, _ptr_bstrInterfaceID); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInterfaceIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrInterfaceID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrInterfaceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InterfaceID == nil {
				o.InterfaceID = &oaut.String{}
			}
			if err := o.InterfaceID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrInterfaceID := func(ptr interface{}) { o.InterfaceID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.InterfaceID, _s_bstrInterfaceID, _ptr_bstrInterfaceID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInterfaceIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInterfaceIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInterfaceIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInterfaceIDRequest structure represents the InterfaceID operation request
type SetInterfaceIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	InterfaceID *oaut.String   `idl:"name:bstrInterfaceID" json:"interface_id"`
}

func (o *SetInterfaceIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetInterfaceIDOperation) *xxx_SetInterfaceIDOperation {
	if op == nil {
		op = &xxx_SetInterfaceIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.InterfaceID = o.InterfaceID
	return op
}

func (o *SetInterfaceIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInterfaceIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InterfaceID = op.InterfaceID
}
func (o *SetInterfaceIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetInterfaceIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInterfaceIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInterfaceIDResponse structure represents the InterfaceID operation response
type SetInterfaceIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The InterfaceID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInterfaceIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetInterfaceIDOperation) *xxx_SetInterfaceIDOperation {
	if op == nil {
		op = &xxx_SetInterfaceIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetInterfaceIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInterfaceIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetInterfaceIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetInterfaceIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInterfaceIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
