package ieventsubscription3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
		in := &GetEventClassPartitionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventClassPartitionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // EventClassPartitionID
		in := &SetEventClassPartitionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEventClassPartitionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // EventClassApplicationID
		in := &GetEventClassApplicationIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventClassApplicationID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // EventClassApplicationID
		in := &SetEventClassApplicationIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEventClassApplicationID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // SubscriberPartitionID
		in := &GetSubscriberPartitionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubscriberPartitionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // SubscriberPartitionID
		in := &SetSubscriberPartitionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubscriberPartitionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // SubscriberApplicationID
		in := &GetSubscriberApplicationIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubscriberApplicationID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // SubscriberApplicationID
		in := &SetSubscriberApplicationIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubscriberApplicationID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
