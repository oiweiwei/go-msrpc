package ieventclass

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

// IEventClass server interface.
type EventClassServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// EventClassID operation.
	GetEventClassID(context.Context, *GetEventClassIDRequest) (*GetEventClassIDResponse, error)

	// EventClassID operation.
	SetEventClassID(context.Context, *SetEventClassIDRequest) (*SetEventClassIDResponse, error)

	// EventClassName operation.
	GetEventClassName(context.Context, *GetEventClassNameRequest) (*GetEventClassNameResponse, error)

	// EventClassName operation.
	SetEventClassName(context.Context, *SetEventClassNameRequest) (*SetEventClassNameResponse, error)

	// OwnerSID operation.
	GetOwnerSID(context.Context, *GetOwnerSIDRequest) (*GetOwnerSIDResponse, error)

	// OwnerSID operation.
	SetOwnerSID(context.Context, *SetOwnerSIDRequest) (*SetOwnerSIDResponse, error)

	// FiringInterfaceID operation.
	GetFiringInterfaceID(context.Context, *GetFiringInterfaceIDRequest) (*GetFiringInterfaceIDResponse, error)

	// FiringInterfaceID operation.
	SetFiringInterfaceID(context.Context, *SetFiringInterfaceIDRequest) (*SetFiringInterfaceIDResponse, error)

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	// Description operation.
	SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error)

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	// Opnum18NotUsedOnWire operation.
	// Opnum18NotUsedOnWire

	// TypeLib operation.
	GetTypeLib(context.Context, *GetTypeLibRequest) (*GetTypeLibResponse, error)

	// TypeLib operation.
	SetTypeLib(context.Context, *SetTypeLibRequest) (*SetTypeLibResponse, error)
}

func RegisterEventClassServer(conn dcerpc.Conn, o EventClassServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventClassServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventClassSyntaxV0_0))...)
}

func NewEventClassServerHandle(o EventClassServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventClassServerHandle(ctx, o, opNum, r)
	}
}

func EventClassServerHandle(ctx context.Context, o EventClassServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // EventClassID
		in := &GetEventClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // EventClassID
		in := &SetEventClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEventClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // EventClassName
		in := &GetEventClassNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventClassName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // EventClassName
		in := &SetEventClassNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEventClassName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // OwnerSID
		in := &GetOwnerSIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOwnerSID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // OwnerSID
		in := &SetOwnerSIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOwnerSID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // FiringInterfaceID
		in := &GetFiringInterfaceIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFiringInterfaceID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // FiringInterfaceID
		in := &SetFiringInterfaceIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFiringInterfaceID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Description
		in := &GetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // Description
		in := &SetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // Opnum17NotUsedOnWire
		// Opnum17NotUsedOnWire
		return nil, nil
	case 18: // Opnum18NotUsedOnWire
		// Opnum18NotUsedOnWire
		return nil, nil
	case 19: // TypeLib
		in := &GetTypeLibRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTypeLib(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // TypeLib
		in := &SetTypeLibRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTypeLib(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
