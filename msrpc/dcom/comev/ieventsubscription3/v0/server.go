package ieventsubscription3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ieventsubscription2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsubscription2/v0"
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
	_ = ieventsubscription2.GoPackage
)

// IEventSubscription3 server interface.
type EventSubscription3Server interface {

	// IEventSubscription2 base class.
	ieventsubscription2.EventSubscription2Server

	// EventClassPartitionID operation.
	GetEventClassPartitionID(context.Context, *GetEventClassPartitionIDRequest) (*GetEventClassPartitionIDResponse, error)

	// EventClassPartitionID operation.
	SetEventClassPartitionID(context.Context, *SetEventClassPartitionIDRequest) (*SetEventClassPartitionIDResponse, error)

	// EventClassApplicationID operation.
	GetEventClassApplicationID(context.Context, *GetEventClassApplicationIDRequest) (*GetEventClassApplicationIDResponse, error)

	// EventClassApplicationID operation.
	SetEventClassApplicationID(context.Context, *SetEventClassApplicationIDRequest) (*SetEventClassApplicationIDResponse, error)

	// SubscriberPartitionID operation.
	GetSubscriberPartitionID(context.Context, *GetSubscriberPartitionIDRequest) (*GetSubscriberPartitionIDResponse, error)

	// SubscriberPartitionID operation.
	SetSubscriberPartitionID(context.Context, *SetSubscriberPartitionIDRequest) (*SetSubscriberPartitionIDResponse, error)

	// SubscriberApplicationID operation.
	GetSubscriberApplicationID(context.Context, *GetSubscriberApplicationIDRequest) (*GetSubscriberApplicationIDResponse, error)

	// SubscriberApplicationID operation.
	SetSubscriberApplicationID(context.Context, *SetSubscriberApplicationIDRequest) (*SetSubscriberApplicationIDResponse, error)
}

func RegisterEventSubscription3Server(conn dcerpc.Conn, o EventSubscription3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventSubscription3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventSubscription3SyntaxV0_0))...)
}

func NewEventSubscription3ServerHandle(o EventSubscription3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventSubscription3ServerHandle(ctx, o, opNum, r)
	}
}

func EventSubscription3ServerHandle(ctx context.Context, o EventSubscription3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 45 {
		// IEventSubscription2 base method.
		return ieventsubscription2.EventSubscription2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 45: // EventClassPartitionID
		op := &xxx_GetEventClassPartitionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventClassPartitionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventClassPartitionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // EventClassPartitionID
		op := &xxx_SetEventClassPartitionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventClassPartitionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEventClassPartitionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // EventClassApplicationID
		op := &xxx_GetEventClassApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventClassApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventClassApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // EventClassApplicationID
		op := &xxx_SetEventClassApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventClassApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEventClassApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // SubscriberPartitionID
		op := &xxx_GetSubscriberPartitionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubscriberPartitionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubscriberPartitionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // SubscriberPartitionID
		op := &xxx_SetSubscriberPartitionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubscriberPartitionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubscriberPartitionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // SubscriberApplicationID
		op := &xxx_GetSubscriberApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubscriberApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubscriberApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // SubscriberApplicationID
		op := &xxx_SetSubscriberApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubscriberApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubscriberApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IEventSubscription3
type UnimplementedEventSubscription3Server struct {
	ieventsubscription2.UnimplementedEventSubscription2Server
}

func (UnimplementedEventSubscription3Server) GetEventClassPartitionID(context.Context, *GetEventClassPartitionIDRequest) (*GetEventClassPartitionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscription3Server) SetEventClassPartitionID(context.Context, *SetEventClassPartitionIDRequest) (*SetEventClassPartitionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscription3Server) GetEventClassApplicationID(context.Context, *GetEventClassApplicationIDRequest) (*GetEventClassApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscription3Server) SetEventClassApplicationID(context.Context, *SetEventClassApplicationIDRequest) (*SetEventClassApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscription3Server) GetSubscriberPartitionID(context.Context, *GetSubscriberPartitionIDRequest) (*GetSubscriberPartitionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscription3Server) SetSubscriberPartitionID(context.Context, *SetSubscriberPartitionIDRequest) (*SetSubscriberPartitionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscription3Server) GetSubscriberApplicationID(context.Context, *GetSubscriberApplicationIDRequest) (*GetSubscriberApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscription3Server) SetSubscriberApplicationID(context.Context, *SetSubscriberApplicationIDRequest) (*SetSubscriberApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EventSubscription3Server = (*UnimplementedEventSubscription3Server)(nil)
