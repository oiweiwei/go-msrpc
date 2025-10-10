package imsmqprivateevent

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

// IMSMQPrivateEvent server interface.
type ImsmqPrivateEventServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Hwnd operation.
	GetHwnd(context.Context, *GetHwndRequest) (*GetHwndResponse, error)

	// FireArrivedEvent operation.
	FireArrivedEvent(context.Context, *FireArrivedEventRequest) (*FireArrivedEventResponse, error)

	// FireArrivedErrorEvent operation.
	FireArrivedErrorEvent(context.Context, *FireArrivedErrorEventRequest) (*FireArrivedErrorEventResponse, error)
}

func RegisterImsmqPrivateEventServer(conn dcerpc.Conn, o ImsmqPrivateEventServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqPrivateEventServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqPrivateEventSyntaxV0_0))...)
}

func NewImsmqPrivateEventServerHandle(o ImsmqPrivateEventServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqPrivateEventServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqPrivateEventServerHandle(ctx context.Context, o ImsmqPrivateEventServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Hwnd
		op := &xxx_GetHwndOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHwndRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHwnd(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // FireArrivedEvent
		op := &xxx_FireArrivedEventOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FireArrivedEventRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FireArrivedEvent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // FireArrivedErrorEvent
		op := &xxx_FireArrivedErrorEventOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FireArrivedErrorEventRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FireArrivedErrorEvent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQPrivateEvent
type UnimplementedImsmqPrivateEventServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqPrivateEventServer) GetHwnd(context.Context, *GetHwndRequest) (*GetHwndResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqPrivateEventServer) FireArrivedEvent(context.Context, *FireArrivedEventRequest) (*FireArrivedEventResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqPrivateEventServer) FireArrivedErrorEvent(context.Context, *FireArrivedErrorEventRequest) (*FireArrivedErrorEventResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqPrivateEventServer = (*UnimplementedImsmqPrivateEventServer)(nil)
