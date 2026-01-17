package ieventclass2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ieventclass "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventclass/v0"
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
	_ = ieventclass.GoPackage
)

// IEventClass2 server interface.
type EventClass2Server interface {

	// IEventClass base class.
	ieventclass.EventClassServer

	// PublisherID operation.
	GetPublisherID(context.Context, *GetPublisherIDRequest) (*GetPublisherIDResponse, error)

	// PublisherID operation.
	SetPublisherID(context.Context, *SetPublisherIDRequest) (*SetPublisherIDResponse, error)

	// MultiInterfacePublisherFilterCLSID operation.
	GetMultiInterfacePublisherFilterClassID(context.Context, *GetMultiInterfacePublisherFilterClassIDRequest) (*GetMultiInterfacePublisherFilterClassIDResponse, error)

	// MultiInterfacePublisherFilterCLSID operation.
	SetMultiInterfacePublisherFilterClassID(context.Context, *SetMultiInterfacePublisherFilterClassIDRequest) (*SetMultiInterfacePublisherFilterClassIDResponse, error)

	// AllowInprocActivation operation.
	GetAllowInProcessActivation(context.Context, *GetAllowInProcessActivationRequest) (*GetAllowInProcessActivationResponse, error)

	// AllowInprocActivation operation.
	SetAllowInProcessActivation(context.Context, *SetAllowInProcessActivationRequest) (*SetAllowInProcessActivationResponse, error)

	// FireInParallel operation.
	GetFireInParallel(context.Context, *GetFireInParallelRequest) (*GetFireInParallelResponse, error)

	// FireInParallel operation.
	SetFireInParallel(context.Context, *SetFireInParallelRequest) (*SetFireInParallelResponse, error)
}

func RegisterEventClass2Server(conn dcerpc.Conn, o EventClass2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventClass2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventClass2SyntaxV0_0))...)
}

func NewEventClass2ServerHandle(o EventClass2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventClass2ServerHandle(ctx, o, opNum, r)
	}
}

func EventClass2ServerHandle(ctx context.Context, o EventClass2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 21 {
		// IEventClass base method.
		return ieventclass.EventClassServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 21: // PublisherID
		op := &xxx_GetPublisherIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPublisherIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPublisherID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // PublisherID
		op := &xxx_SetPublisherIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPublisherIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPublisherID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // MultiInterfacePublisherFilterCLSID
		op := &xxx_GetMultiInterfacePublisherFilterClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMultiInterfacePublisherFilterClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMultiInterfacePublisherFilterClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // MultiInterfacePublisherFilterCLSID
		op := &xxx_SetMultiInterfacePublisherFilterClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMultiInterfacePublisherFilterClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMultiInterfacePublisherFilterClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // AllowInprocActivation
		op := &xxx_GetAllowInProcessActivationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllowInProcessActivationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllowInProcessActivation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // AllowInprocActivation
		op := &xxx_SetAllowInProcessActivationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAllowInProcessActivationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAllowInProcessActivation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // FireInParallel
		op := &xxx_GetFireInParallelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFireInParallelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFireInParallel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // FireInParallel
		op := &xxx_SetFireInParallelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFireInParallelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFireInParallel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IEventClass2
type UnimplementedEventClass2Server struct {
	ieventclass.UnimplementedEventClassServer
}

func (UnimplementedEventClass2Server) GetPublisherID(context.Context, *GetPublisherIDRequest) (*GetPublisherIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClass2Server) SetPublisherID(context.Context, *SetPublisherIDRequest) (*SetPublisherIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClass2Server) GetMultiInterfacePublisherFilterClassID(context.Context, *GetMultiInterfacePublisherFilterClassIDRequest) (*GetMultiInterfacePublisherFilterClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClass2Server) SetMultiInterfacePublisherFilterClassID(context.Context, *SetMultiInterfacePublisherFilterClassIDRequest) (*SetMultiInterfacePublisherFilterClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClass2Server) GetAllowInProcessActivation(context.Context, *GetAllowInProcessActivationRequest) (*GetAllowInProcessActivationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClass2Server) SetAllowInProcessActivation(context.Context, *SetAllowInProcessActivationRequest) (*SetAllowInProcessActivationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClass2Server) GetFireInParallel(context.Context, *GetFireInParallelRequest) (*GetFireInParallelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClass2Server) SetFireInParallel(context.Context, *SetFireInParallelRequest) (*SetFireInParallelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EventClass2Server = (*UnimplementedEventClass2Server)(nil)
