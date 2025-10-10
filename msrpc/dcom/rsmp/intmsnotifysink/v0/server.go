package intmsnotifysink

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = iunknown.GoPackage
)

// INtmsNotifySink server interface.
type NTMSNotifySinkServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	ConnectCallback(context.Context, *ConnectCallbackRequest) (*ConnectCallbackResponse, error)

	// OnNotify operation.
	OnNotify(context.Context, *OnNotifyRequest) (*OnNotifyResponse, error)

	ReleaseCallback(context.Context, *ReleaseCallbackRequest) (*ReleaseCallbackResponse, error)
}

func RegisterNTMSNotifySinkServer(conn dcerpc.Conn, o NTMSNotifySinkServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNTMSNotifySinkServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NTMSNotifySinkSyntaxV0_0))...)
}

func NewNTMSNotifySinkServerHandle(o NTMSNotifySinkServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NTMSNotifySinkServerHandle(ctx, o, opNum, r)
	}
}

func NTMSNotifySinkServerHandle(ctx context.Context, o NTMSNotifySinkServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
type UnimplementedNTMSNotifySinkServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedNTMSNotifySinkServer) ConnectCallback(context.Context, *ConnectCallbackRequest) (*ConnectCallbackResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSNotifySinkServer) OnNotify(context.Context, *OnNotifyRequest) (*OnNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNTMSNotifySinkServer) ReleaseCallback(context.Context, *ReleaseCallbackRequest) (*ReleaseCallbackResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NTMSNotifySinkServer = (*UnimplementedNTMSNotifySinkServer)(nil)
