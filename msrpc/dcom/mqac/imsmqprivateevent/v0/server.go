package imsmqprivateevent

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IMSMQPrivateEvent server interface.
type PrivateEventServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The Hwnd method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a LONG value, which the client MUST ignore. Because the returned
	// LONG value serves no purpose, the server MAY<90> return 0x00000000.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	//
	// This interface is optional for communication with the client. If implemented, this
	// interface MUST reside on the server.
	GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error)

	// The FireArrivedEvent method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST provide notification of the availability of a Message.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	FireArrivedEvent(context.Context, *FireArrivedEventRequest) (*FireArrivedEventResponse, error)

	// The FireArrivedErrorEvent method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST provide notification of an error relating to the arrival
	// of a message.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	FireArrivedErrorEvent(context.Context, *FireArrivedErrorEventRequest) (*FireArrivedErrorEventResponse, error)
}

func RegisterPrivateEventServer(conn dcerpc.Conn, o PrivateEventServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPrivateEventServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PrivateEventSyntaxV0_0))...)
}

func NewPrivateEventServerHandle(o PrivateEventServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PrivateEventServerHandle(ctx, o, opNum, r)
	}
}

func PrivateEventServerHandle(ctx context.Context, o PrivateEventServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Hwnd
		op := &xxx_GetHandleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHandleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHandle(ctx, req)
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
type UnimplementedPrivateEventServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedPrivateEventServer) GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPrivateEventServer) FireArrivedEvent(context.Context, *FireArrivedEventRequest) (*FireArrivedEventResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPrivateEventServer) FireArrivedErrorEvent(context.Context, *FireArrivedErrorEventRequest) (*FireArrivedErrorEventResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PrivateEventServer = (*UnimplementedPrivateEventServer)(nil)
