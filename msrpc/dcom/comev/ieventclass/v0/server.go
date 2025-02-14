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
		op := &xxx_GetEventClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // EventClassID
		op := &xxx_SetEventClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEventClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // EventClassName
		op := &xxx_GetEventClassNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventClassNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventClassName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // EventClassName
		op := &xxx_SetEventClassNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventClassNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEventClassName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // OwnerSID
		op := &xxx_GetOwnerSIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOwnerSIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOwnerSID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // OwnerSID
		op := &xxx_SetOwnerSIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOwnerSIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOwnerSID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // FiringInterfaceID
		op := &xxx_GetFiringInterfaceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFiringInterfaceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFiringInterfaceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // FiringInterfaceID
		op := &xxx_SetFiringInterfaceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFiringInterfaceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFiringInterfaceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Description
		op := &xxx_GetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Description
		op := &xxx_SetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Opnum17NotUsedOnWire
		// Opnum17NotUsedOnWire
		return nil, nil
	case 18: // Opnum18NotUsedOnWire
		// Opnum18NotUsedOnWire
		return nil, nil
	case 19: // TypeLib
		op := &xxx_GetTypeLibOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeLibRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTypeLib(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // TypeLib
		op := &xxx_SetTypeLibOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTypeLibRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTypeLib(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IEventClass
type UnimplementedEventClassServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedEventClassServer) GetEventClassID(context.Context, *GetEventClassIDRequest) (*GetEventClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) SetEventClassID(context.Context, *SetEventClassIDRequest) (*SetEventClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) GetEventClassName(context.Context, *GetEventClassNameRequest) (*GetEventClassNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) SetEventClassName(context.Context, *SetEventClassNameRequest) (*SetEventClassNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) GetOwnerSID(context.Context, *GetOwnerSIDRequest) (*GetOwnerSIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) SetOwnerSID(context.Context, *SetOwnerSIDRequest) (*SetOwnerSIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) GetFiringInterfaceID(context.Context, *GetFiringInterfaceIDRequest) (*GetFiringInterfaceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) SetFiringInterfaceID(context.Context, *SetFiringInterfaceIDRequest) (*SetFiringInterfaceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) GetTypeLib(context.Context, *GetTypeLibRequest) (*GetTypeLibResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClassServer) SetTypeLib(context.Context, *SetTypeLibRequest) (*SetTypeLibResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EventClassServer = (*UnimplementedEventClassServer)(nil)
