package imsmqqueue2

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

// IMSMQQueue2 server interface.
type Queue2Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The Access method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns a flag that indicates the access mode in which the queue was opened.
	// The access mode specifies whether peek, receive, send, and/or administration operations
	// can be performed.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetAccess(context.Context, *GetAccessRequest) (*GetAccessResponse, error)

	// The ShareMode method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns a flag that indicates the share mode in which the queue was opened.
	// The share mode specifies whether this instance of the open queue has exclusive access
	// to the queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetShareMode(context.Context, *GetShareModeRequest) (*GetShareModeResponse, error)

	// The QueueInfo method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns an MSMQQueueInfo object that represents the referenced queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetQueueInfo(context.Context, *GetQueueInfoRequest) (*GetQueueInfoResponse, error)

	// Handle operation.
	GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error)

	// IsOpen operation.
	GetIsOpen(context.Context, *GetIsOpenRequest) (*GetIsOpenResponse, error)

	// Close operation.
	Close(context.Context, *CloseRequest) (*CloseResponse, error)

	// The Receive_v1 method is received by the server in an RPC_REQUEST packet. In response,
	// the server retrieves the Message at the head of the referenced queue's MessagePositionList.Head
	// and removes it.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	ReceiveV1(context.Context, *ReceiveV1Request) (*ReceiveV1Response, error)

	// The Peek_v1 method is received by the server in an RPC_REQUEST packet. In response,
	// the server retrieves the Message at the head of the referenced queue's MessagePositionList
	// without removing it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	PeekV1(context.Context, *PeekV1Request) (*PeekV1Response, error)

	// The EnableNotification method is received by the server in an RPC_REQUEST packet.
	// In response, the server starts event notification for asynchronously receiving or
	// peeking messages.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	EnableNotification(context.Context, *EnableNotificationRequest) (*EnableNotificationResponse, error)

	// Reset operation.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// The ReceiveCurrent_v1 method is received by the server in an RPC_REQUEST packet.
	// In response, the server retrieves the Message at the current cursor position in the
	// referenced queue's MessagePositionList, removes it, and advances the cursor.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	ReceiveCurrentV1(context.Context, *ReceiveCurrentV1Request) (*ReceiveCurrentV1Response, error)

	// The PeekNext_v1 method is received by the server in an RPC_REQUEST packet. In response,
	// the server retrieves the Message that follows the Message that is identified by the
	// cursor represented by the Cursor instance variable in the referenced queue's MessagePositionList
	// without removing it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	PeekNextV1(context.Context, *PeekNextV1Request) (*PeekNextV1Response, error)

	// The PeekCurrent_v1 method is received by the server in an RPC_REQUEST packet. In
	// response, the server retrieves the Message that is identified by the Cursor instance
	// variable in the referenced queue's MessagePositionList, without removing it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	PeekCurrentV1(context.Context, *PeekCurrentV1Request) (*PeekCurrentV1Response, error)

	// The Receive method is received by the server in an RPC_REQUEST packet. In response,
	// the server retrieves the Message at the head of the referenced queue's MessagePositionList
	// and removes it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error)

	// The Peek method is received by the server in an RPC_REQUEST packet. In response,
	// the server retrieves the Message at the head of the referenced queue's MessagePositionList
	// without removing it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Peek(context.Context, *PeekRequest) (*PeekResponse, error)

	// The ReceiveCurrent method is received by the server in an RPC_REQUEST packet. In
	// response, the server retrieves the Message that is identified by the Cursor instance
	// variable in the referenced queue's MessagePositionList and removes it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	ReceiveCurrent(context.Context, *ReceiveCurrentRequest) (*ReceiveCurrentResponse, error)

	// The PeekNext method is received by the server in an RPC_REQUEST packet. In response,
	// the server retrieves the Message that follows the Message that is identified by the
	// Cursor instance variable in the referenced queue's MessagePositionList without removing
	// it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	PeekNext(context.Context, *PeekNextRequest) (*PeekNextResponse, error)

	// The PeekCurrent method is received by the server in an RPC_REQUEST packet. In response,
	// the server retrieves the Message that is identified by the Cursor instance variable
	// in the referenced queue's MessagePositionList, without removing it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	PeekCurrent(context.Context, *PeekCurrentRequest) (*PeekCurrentResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterQueue2Server(conn dcerpc.Conn, o Queue2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQueue2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Queue2SyntaxV0_0))...)
}

func NewQueue2ServerHandle(o Queue2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Queue2ServerHandle(ctx, o, opNum, r)
	}
}

func Queue2ServerHandle(ctx context.Context, o Queue2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 13: // Receive_v1
		op := &xxx_ReceiveV1Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveV1Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReceiveV1(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Peek_v1
		op := &xxx_PeekV1Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekV1Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekV1(ctx, req)
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
	case 17: // ReceiveCurrent_v1
		op := &xxx_ReceiveCurrentV1Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveCurrentV1Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReceiveCurrentV1(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // PeekNext_v1
		op := &xxx_PeekNextV1Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekNextV1Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekNextV1(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // PeekCurrent_v1
		op := &xxx_PeekCurrentV1Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekCurrentV1Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekCurrentV1(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Receive
		op := &xxx_ReceiveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Receive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Peek
		op := &xxx_PeekOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Peek(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // ReceiveCurrent
		op := &xxx_ReceiveCurrentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveCurrentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReceiveCurrent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // PeekNext
		op := &xxx_PeekNextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekNextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekNext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // PeekCurrent
		op := &xxx_PeekCurrentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekCurrentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekCurrent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQueue2
type UnimplementedQueue2Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedQueue2Server) GetAccess(context.Context, *GetAccessRequest) (*GetAccessResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) GetShareMode(context.Context, *GetShareModeRequest) (*GetShareModeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) GetQueueInfo(context.Context, *GetQueueInfoRequest) (*GetQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) GetIsOpen(context.Context, *GetIsOpenRequest) (*GetIsOpenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) ReceiveV1(context.Context, *ReceiveV1Request) (*ReceiveV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) PeekV1(context.Context, *PeekV1Request) (*PeekV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) EnableNotification(context.Context, *EnableNotificationRequest) (*EnableNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) ReceiveCurrentV1(context.Context, *ReceiveCurrentV1Request) (*ReceiveCurrentV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) PeekNextV1(context.Context, *PeekNextV1Request) (*PeekNextV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) PeekCurrentV1(context.Context, *PeekCurrentV1Request) (*PeekCurrentV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) Peek(context.Context, *PeekRequest) (*PeekResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) ReceiveCurrent(context.Context, *ReceiveCurrentRequest) (*ReceiveCurrentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) PeekNext(context.Context, *PeekNextRequest) (*PeekNextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) PeekCurrent(context.Context, *PeekCurrentRequest) (*PeekCurrentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue2Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Queue2Server = (*UnimplementedQueue2Server)(nil)
