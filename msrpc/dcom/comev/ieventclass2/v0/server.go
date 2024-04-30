package ieventclass2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
		in := &GetPublisherIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPublisherID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // PublisherID
		in := &SetPublisherIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPublisherID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // MultiInterfacePublisherFilterCLSID
		in := &GetMultiInterfacePublisherFilterClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMultiInterfacePublisherFilterClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // MultiInterfacePublisherFilterCLSID
		in := &SetMultiInterfacePublisherFilterClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMultiInterfacePublisherFilterClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // AllowInprocActivation
		in := &GetAllowInProcessActivationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllowInProcessActivation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // AllowInprocActivation
		in := &SetAllowInProcessActivationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAllowInProcessActivation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // FireInParallel
		in := &GetFireInParallelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFireInParallel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // FireInParallel
		in := &SetFireInParallelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFireInParallel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
