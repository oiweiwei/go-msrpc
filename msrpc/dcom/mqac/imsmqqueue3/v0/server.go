package imsmqqueue3

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

// IMSMQQueue3 server interface.
type Queue3Server interface {

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

	// The Handle2 method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the handle of the open queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetHandle2(context.Context, *GetHandle2Request) (*GetHandle2Response, error)

	// The ReceiveByLookupId method is received by the server in an RPC_REQUEST packet.
	// In response, the server retrieves the Message from the referenced queue's MessagePositionList
	// for which the Message.LookupIdentifier property equals the lookup identifier and
	// removes it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	ReceiveByLookupID(context.Context, *ReceiveByLookupIDRequest) (*ReceiveByLookupIDResponse, error)

	// The ReceiveNextByLookupId method is received by the server in an RPC_REQUEST packet.
	// In response, the server retrieves the Message following a Message from the referenced
	// queue's MessagePositionList for which the Message.LookupIdentifier property equals
	// the lookup identifier and removes it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	ReceiveNextByLookupID(context.Context, *ReceiveNextByLookupIDRequest) (*ReceiveNextByLookupIDResponse, error)

	// The ReceivePreviousByLookupId method is received by the server in an RPC_REQUEST
	// packet. In response, the server retrieves the Message preceding a Message from the
	// referenced queue's MessagePositionList for which the Message.LookupIdentifier property
	// equals the lookup identifier and removes it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	ReceivePreviousByLookupID(context.Context, *ReceivePreviousByLookupIDRequest) (*ReceivePreviousByLookupIDResponse, error)

	// The ReceiveFirstByLookupId method is received by the server in an RPC_REQUEST packet.
	// In response, the server retrieves the Message at the head of the referenced queue's
	// MessagePositionList and removes it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	ReceiveFirstByLookupID(context.Context, *ReceiveFirstByLookupIDRequest) (*ReceiveFirstByLookupIDResponse, error)

	// The ReceiveLastByLookupId method is received by the server in an RPC_REQUEST packet.
	// In response, the server retrieves the Message at the tail of the referenced queue's
	// MessagePositionList and removes it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	ReceiveLastByLookupID(context.Context, *ReceiveLastByLookupIDRequest) (*ReceiveLastByLookupIDResponse, error)

	// The PeekByLookupId method is received by the server in an RPC_REQUEST packet. In
	// response, the server retrieves the Message from the referenced queue's MessagePositionList
	// for which the Message.LookupIdentifier property equals the lookup ID without removing
	// it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	PeekByLookupID(context.Context, *PeekByLookupIDRequest) (*PeekByLookupIDResponse, error)

	// The PeekNextByLookupId method is received by the server in an RPC_REQUEST packet.
	// In response, the server retrieves the Message following a Message from the referenced
	// queue's MessagePositionList for which the Message.LookupIdentifier property equals
	// the lookup identifier without removing it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	PeekNextByLookupID(context.Context, *PeekNextByLookupIDRequest) (*PeekNextByLookupIDResponse, error)

	// The PeekPreviousByLookupId method is received by the server in an RPC_REQUEST packet.
	// In response, the server retrieves the Message preceding a Message from the referenced
	// queue's MessagePositionList for which the Message.LookupIdentifier property equals
	// the lookup identifier without removing it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	PeekPreviousByLookupID(context.Context, *PeekPreviousByLookupIDRequest) (*PeekPreviousByLookupIDResponse, error)

	// The PeekFirstByLookupId method is received by the server in an RPC_REQUEST packet.
	// In response, the server retrieves the Message at the head of the referenced queue's
	// MessagePositionList, without removing it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	PeekFirstByLookupID(context.Context, *PeekFirstByLookupIDRequest) (*PeekFirstByLookupIDResponse, error)

	// The PeekLastByLookupId method is received by the server in an RPC_REQUEST packet.
	// In response, the server retrieves the Message at the tail of the referenced queue's
	// MessagePositionList without removing it.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	PeekLastByLookupID(context.Context, *PeekLastByLookupIDRequest) (*PeekLastByLookupIDResponse, error)

	// The Purge method is received by the server in an RPC_REQUEST packet. In response,
	// the server purges all Messages in the MessagePositionList of the referenced queue.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Purge(context.Context, *PurgeRequest) (*PurgeResponse, error)

	// The IsOpen2 method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns a value indicating whether the queue is open.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	GetIsOpen2(context.Context, *GetIsOpen2Request) (*GetIsOpen2Response, error)
}

func RegisterQueue3Server(conn dcerpc.Conn, o Queue3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQueue3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Queue3SyntaxV0_0))...)
}

func NewQueue3ServerHandle(o Queue3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Queue3ServerHandle(ctx, o, opNum, r)
	}
}

func Queue3ServerHandle(ctx context.Context, o Queue3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 26: // Handle2
		op := &xxx_GetHandle2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHandle2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHandle2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // ReceiveByLookupId
		op := &xxx_ReceiveByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReceiveByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // ReceiveNextByLookupId
		op := &xxx_ReceiveNextByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveNextByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReceiveNextByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // ReceivePreviousByLookupId
		op := &xxx_ReceivePreviousByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceivePreviousByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReceivePreviousByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // ReceiveFirstByLookupId
		op := &xxx_ReceiveFirstByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveFirstByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReceiveFirstByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // ReceiveLastByLookupId
		op := &xxx_ReceiveLastByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveLastByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReceiveLastByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // PeekByLookupId
		op := &xxx_PeekByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // PeekNextByLookupId
		op := &xxx_PeekNextByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekNextByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekNextByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // PeekPreviousByLookupId
		op := &xxx_PeekPreviousByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekPreviousByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekPreviousByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // PeekFirstByLookupId
		op := &xxx_PeekFirstByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekFirstByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekFirstByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // PeekLastByLookupId
		op := &xxx_PeekLastByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PeekLastByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PeekLastByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // Purge
		op := &xxx_PurgeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PurgeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Purge(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // IsOpen2
		op := &xxx_GetIsOpen2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsOpen2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsOpen2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQueue3
type UnimplementedQueue3Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedQueue3Server) GetAccess(context.Context, *GetAccessRequest) (*GetAccessResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) GetShareMode(context.Context, *GetShareModeRequest) (*GetShareModeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) GetQueueInfo(context.Context, *GetQueueInfoRequest) (*GetQueueInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) GetIsOpen(context.Context, *GetIsOpenRequest) (*GetIsOpenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) ReceiveV1(context.Context, *ReceiveV1Request) (*ReceiveV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) PeekV1(context.Context, *PeekV1Request) (*PeekV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) EnableNotification(context.Context, *EnableNotificationRequest) (*EnableNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) ReceiveCurrentV1(context.Context, *ReceiveCurrentV1Request) (*ReceiveCurrentV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) PeekNextV1(context.Context, *PeekNextV1Request) (*PeekNextV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) PeekCurrentV1(context.Context, *PeekCurrentV1Request) (*PeekCurrentV1Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) Peek(context.Context, *PeekRequest) (*PeekResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) ReceiveCurrent(context.Context, *ReceiveCurrentRequest) (*ReceiveCurrentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) PeekNext(context.Context, *PeekNextRequest) (*PeekNextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) PeekCurrent(context.Context, *PeekCurrentRequest) (*PeekCurrentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) GetHandle2(context.Context, *GetHandle2Request) (*GetHandle2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) ReceiveByLookupID(context.Context, *ReceiveByLookupIDRequest) (*ReceiveByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) ReceiveNextByLookupID(context.Context, *ReceiveNextByLookupIDRequest) (*ReceiveNextByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) ReceivePreviousByLookupID(context.Context, *ReceivePreviousByLookupIDRequest) (*ReceivePreviousByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) ReceiveFirstByLookupID(context.Context, *ReceiveFirstByLookupIDRequest) (*ReceiveFirstByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) ReceiveLastByLookupID(context.Context, *ReceiveLastByLookupIDRequest) (*ReceiveLastByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) PeekByLookupID(context.Context, *PeekByLookupIDRequest) (*PeekByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) PeekNextByLookupID(context.Context, *PeekNextByLookupIDRequest) (*PeekNextByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) PeekPreviousByLookupID(context.Context, *PeekPreviousByLookupIDRequest) (*PeekPreviousByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) PeekFirstByLookupID(context.Context, *PeekFirstByLookupIDRequest) (*PeekFirstByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) PeekLastByLookupID(context.Context, *PeekLastByLookupIDRequest) (*PeekLastByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) Purge(context.Context, *PurgeRequest) (*PurgeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueue3Server) GetIsOpen2(context.Context, *GetIsOpen2Request) (*GetIsOpen2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Queue3Server = (*UnimplementedQueue3Server)(nil)
