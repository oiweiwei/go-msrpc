package ieventsubscription2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
		in := &GetFilterCriteriaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFilterCriteria(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // FilterCriteria
		in := &SetFilterCriteriaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFilterCriteria(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // SubscriberMoniker
		in := &GetSubscriberMonikerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubscriberMoniker(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // SubscriberMoniker
		in := &SetSubscriberMonikerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubscriberMoniker(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
