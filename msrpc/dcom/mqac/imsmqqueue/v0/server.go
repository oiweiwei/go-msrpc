package imsmqqueue

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

// IMSMQQueue server interface.
type QueueServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Access operation.
	GetAccess(context.Context, *GetAccessRequest) (*GetAccessResponse, error)

	// ShareMode operation.
	GetShareMode(context.Context, *GetShareModeRequest) (*GetShareModeResponse, error)

	// QueueInfo operation.
	GetQueueInfo(context.Context, *GetQueueInfoRequest) (*GetQueueInfoResponse, error)

	// Handle operation.
	GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error)

	// IsOpen operation.
	GetIsOpen(context.Context, *GetIsOpenRequest) (*GetIsOpenResponse, error)

	// Close operation.
	Close(context.Context, *CloseRequest) (*CloseResponse, error)

	// Receive operation.
	Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error)

	// Peek operation.
	Peek(context.Context, *PeekRequest) (*PeekResponse, error)

	// EnableNotification operation.
	EnableNotification(context.Context, *EnableNotificationRequest) (*EnableNotificationResponse, error)

	// Reset operation.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// ReceiveCurrent operation.
	ReceiveCurrent(context.Context, *ReceiveCurrentRequest) (*ReceiveCurrentResponse, error)

	// PeekNext operation.
	PeekNext(context.Context, *PeekNextRequest) (*PeekNextResponse, error)

	// PeekCurrent operation.
	PeekCurrent(context.Context, *PeekCurrentRequest) (*PeekCurrentResponse, error)
}

func RegisterQueueServer(conn dcerpc.Conn, o QueueServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQueueServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QueueSyntaxV0_0))...)
}

func NewQueueServerHandle(o QueueServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QueueServerHandle(ctx, o, opNum, r)
	}
}

func QueueServerHandle(ctx context.Context, o QueueServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Access
		op := &xxx_GetAccessOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAccessRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAccess(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ShareMode
		op := &xxx_GetShareModeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetShareModeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetShareMode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // QueueInfo
		op := &xxx_GetQueueInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQueueInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQueueInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Handle
		op := &xxx_GetHandleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHandleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHandle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // IsOpen
		op := &xxx_GetIsOpenOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsOpenRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsOpen(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Close
		op := &xxx_CloseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Close(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Receive
		op := &xxx_ReceiveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Receive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Peek
		op := &xxx_PeekOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Peek(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // EnableNotification
		op := &xxx_EnableNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnableNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnableNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Reset
		op := &xxx_ResetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Reset(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // ReceiveCurrent
		op := &xxx_ReceiveCurrentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveCurrentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReceiveCurrent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // PeekNext
		op := &xxx_PeekNextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekNextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekNext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // PeekCurrent
		op := &xxx_PeekCurrentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekCurrentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekCurrent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQueue
type UnimplementedQueueServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedQueueServer) GetAccess(context.Context, *GetAccessRequest) (*GetAccessResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) GetShareMode(context.Context, *GetShareModeRequest) (*GetShareModeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) GetQueueInfo(context.Context, *GetQueueInfoRequest) (*GetQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) GetIsOpen(context.Context, *GetIsOpenRequest) (*GetIsOpenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) Peek(context.Context, *PeekRequest) (*PeekResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) EnableNotification(context.Context, *EnableNotificationRequest) (*EnableNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) ReceiveCurrent(context.Context, *ReceiveCurrentRequest) (*ReceiveCurrentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) PeekNext(context.Context, *PeekNextRequest) (*PeekNextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueServer) PeekCurrent(context.Context, *PeekCurrentRequest) (*PeekCurrentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QueueServer = (*UnimplementedQueueServer)(nil)
