package ieventsubscription2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ieventsubscription "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsubscription/v0"
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
	_ = ieventsubscription.GoPackage
)

// IEventSubscription2 server interface.
type EventSubscription2Server interface {

	// IEventSubscription base class.
	ieventsubscription.EventSubscriptionServer

	// FilterCriteria operation.
	GetFilterCriteria(context.Context, *GetFilterCriteriaRequest) (*GetFilterCriteriaResponse, error)

	// FilterCriteria operation.
	SetFilterCriteria(context.Context, *SetFilterCriteriaRequest) (*SetFilterCriteriaResponse, error)

	// SubscriberMoniker operation.
	GetSubscriberMoniker(context.Context, *GetSubscriberMonikerRequest) (*GetSubscriberMonikerResponse, error)

	// SubscriberMoniker operation.
	SetSubscriberMoniker(context.Context, *SetSubscriberMonikerRequest) (*SetSubscriberMonikerResponse, error)
}

func RegisterEventSubscription2Server(conn dcerpc.Conn, o EventSubscription2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventSubscription2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventSubscription2SyntaxV0_0))...)
}

func NewEventSubscription2ServerHandle(o EventSubscription2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventSubscription2ServerHandle(ctx, o, opNum, r)
	}
}

func EventSubscription2ServerHandle(ctx context.Context, o EventSubscription2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 41 {
		// IEventSubscription base method.
		return ieventsubscription.EventSubscriptionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 41: // FilterCriteria
		op := &xxx_GetFilterCriteriaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFilterCriteriaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFilterCriteria(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // FilterCriteria
		op := &xxx_SetFilterCriteriaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFilterCriteriaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFilterCriteria(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // SubscriberMoniker
		op := &xxx_GetSubscriberMonikerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubscriberMonikerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubscriberMoniker(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // SubscriberMoniker
		op := &xxx_SetSubscriberMonikerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubscriberMonikerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubscriberMoniker(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IEventSubscription2
type UnimplementedEventSubscription2Server struct {
	ieventsubscription.UnimplementedEventSubscriptionServer
}

func (UnimplementedEventSubscription2Server) GetFilterCriteria(context.Context, *GetFilterCriteriaRequest) (*GetFilterCriteriaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscription2Server) SetFilterCriteria(context.Context, *SetFilterCriteriaRequest) (*SetFilterCriteriaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscription2Server) GetSubscriberMoniker(context.Context, *GetSubscriberMonikerRequest) (*GetSubscriberMonikerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventSubscription2Server) SetSubscriberMoniker(context.Context, *SetSubscriberMonikerRequest) (*SetSubscriberMonikerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EventSubscription2Server = (*UnimplementedEventSubscription2Server)(nil)
