package intmsnotifysink

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// INtmsNotifySink server interface.
type NotifySinkServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The ConnectCallback method connects a connection point to the sink.
	ConnectCallback(context.Context, *ConnectCallbackRequest) (*ConnectCallbackResponse, error)

	// OnNotify operation.
	OnNotify(context.Context, *OnNotifyRequest) (*OnNotifyResponse, error)

	// The ReleaseCallback method removes a connection point from the sink.
	//
	// This method has no parameters.
	//
	//	+-------------------+------------------------------------+
	//	|      RETURN       |                                    |
	//	|    VALUE/CODE     |            DESCRIPTION             |
	//	|                   |                                    |
	//	+-------------------+------------------------------------+
	//	+-------------------+------------------------------------+
	//	| 0x00000000 S_OK   | The method completed successfully. |
	//	+-------------------+------------------------------------+
	ReleaseCallback(context.Context, *ReleaseCallbackRequest) (*ReleaseCallbackResponse, error)
}

func RegisterNotifySinkServer(conn dcerpc.Conn, o NotifySinkServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNotifySinkServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NotifySinkSyntaxV0_0))...)
}

func NewNotifySinkServerHandle(o NotifySinkServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NotifySinkServerHandle(ctx, o, opNum, r)
	}
}

func NotifySinkServerHandle(ctx context.Context, o NotifySinkServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ConnectCallback
		op := &xxx_ConnectCallbackOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectCallbackRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectCallback(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // OnNotify
		op := &xxx_OnNotifyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OnNotifyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OnNotify(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ReleaseCallback
		op := &xxx_ReleaseCallbackOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReleaseCallbackRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReleaseCallback(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented INtmsNotifySink
type UnimplementedNotifySinkServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedNotifySinkServer) ConnectCallback(context.Context, *ConnectCallbackRequest) (*ConnectCallbackResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNotifySinkServer) OnNotify(context.Context, *OnNotifyRequest) (*OnNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNotifySinkServer) ReleaseCallback(context.Context, *ReleaseCallbackRequest) (*ReleaseCallbackResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NotifySinkServer = (*UnimplementedNotifySinkServer)(nil)
