package qm2qm

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// qm2qm server interface.
type Qm2qmServer interface {

	// The RemoteQMStartReceive method peeks or receives a message from an open queue.
	//
	// If RemoteQMStartReceive is invoked with a Peek action type, as specified in the ulAction
	// member of the lpRemoteReadDesc parameter, the operation completes when RemoteQMStartReceive
	// returns.
	//
	// If RemoteQMStartReceive is invoked with a Receive action type, as specified in the
	// ulAction member of the lpRemoteReadDesc parameter, the client MUST pair each call
	// to RemoteQMStartReceive with a call to RemoteQMEndReceive to complete the operation,
	// or to RemoteQMCancelReceive to cancel the operation.
	//
	// For each call to RemoteQMCancelReceive, the dwRequestID parameter MUST match the
	// dwRequestID member of the lpRemoteReadDesc parameter in a previous call to RemoteQMStartReceive.
	//
	// If the client specifies a nonzero value for the ulTimeout member of the lpRemoteReadDesc
	// parameter, and a message is not available in the queue at the time of the call, the
	// server waits up to the specified time-out for a message to become available in the
	// queue before responding to the call. The client can call RemoteQMCancelReceive with
	// a matching REMOTEREADDESC.dwRequestID to cancel the pending RemoteQMStartReceive
	// request.
	//
	// Before calling this method, the client MUST have already called RemoteQMOpenQueue.
	//
	// Return Values: The method MUST return MQ_OK (0x00000000) on success; otherwise, it
	// MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs identically.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_INVALID_HANDLE (0xC00E0007)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// STATUS_INVALID_PARAMETER (0xC000000D)
	//
	// Exceptions Thrown: None except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	StartReceive(context.Context, *StartReceiveRequest) (*StartReceiveResponse, error)

	// The client MUST invoke the RemoteQMEndReceive method to advise the server that the
	// message packet returned by the RemoteQMStartReceive, RemoteQMStartReceive2, or RemoteQMStartReceiveByLookupId
	// method has been received.
	//
	// The combination of the RemoteQMStartReceive, RemoteQMStartReceive2, or RemoteQMStartReceiveByLookupId
	// method and the positive acknowledgment of the RemoteQMEndReceive method ensures that
	// a message packet is not lost in transit from the server to the client due to a network
	// outage during the call sequence.
	//
	// Before calling this method, the following methods MUST be called:
	//
	// * RemoteQMOpenQueue
	//
	// * RemoteQMStartReceive, RemoteQMStartReceive2, or RemoteQMStartReceiveByLookupId
	//
	// HRESULT RemoteQMEndReceive(
	//
	// [in] handle_t hBind,
	//
	// * [in, range(1, 2)] DWORD dwAck
	//
	// );
	//
	// Return Values: The method MUST return MQ_OK (0x00000000) on success; otherwise, it
	// MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs identically.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_INVALID_HANDLE (0xC00E0007)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// MQ_ERROR_TRANSACTION_SEQUENCE (0xC00E0051)
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EndReceive(context.Context, *EndReceiveRequest) (*EndReceiveResponse, error)

	// The RemoteQMOpenQueue method opens a queue in preparation for subsequent operations
	// against it. This method assumes that the client has called qmcomm:R_QMOpenRemoteQueue
	// to obtain a queue handle; for more information, see [MS-MQMP] section 3.1.4.2. This
	// method is called as part of the sequence of events involved in opening a remote queue
	// by an MQMP application as described in [MS-MQMP] section 4.2. This method MUST be
	// called prior to calling any of the following operations:
	//
	// * RemoteQMStartReceive ( aab4d8fa-3a56-483c-ac61-4172a09fc199 )
	//
	// * RemoteQMEndReceive ( 4c0b91ba-37fa-441c-974a-91bbea58a62d )
	//
	// * RemoteQMCloseQueue ( f8b37cf7-cdf4-4e3a-8497-cf2c65074db5 )
	//
	// * RemoteQMCloseCursor ( 0ec761b0-e700-4e3d-b735-54bdf458509d )
	//
	// * RemoteQMCancelReceive ( 76f56b96-7c8d-432e-a169-a1b750e20527 )
	//
	// * RemoteQMPurgeQueue ( 119644cf-f6ca-4f74-bb4a-c2624ccab22c )
	//
	// * RemoteQMStartReceive2 ( a0df225e-80ea-4242-90e5-9891bfa9ebfb )
	//
	// * RemoteQMStartReceiveByLookupId ( b8a82b34-6feb-4057-8f7f-48275fa65068 )
	//
	// HRESULT RemoteQMOpenQueue(
	//
	// [in] handle_t hBind,
	//
	// * [in, range(0, 16)] DWORD dwMQS,
	//
	// [in] DWORD hQueue,
	//
	// [in] DWORD pQueue,
	//
	// [in] DWORD dwpContext
	//
	// );
	//
	// Return Values: The method MUST return MQ_OK (0x00000000) on success; otherwise, it
	// MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs identically.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol (see
	// [MS-RPCE]).
	OpenQueue(context.Context, *OpenQueueRequest) (*OpenQueueResponse, error)

	// The RemoteQMCloseQueue method closes a PCTX_RRSESSION_HANDLE_TYPE that was previously
	// opened by using a call to the RemoteQMOpenQueue method. The client MUST call this
	// method to reclaim resources on the server allocated by the RemoteQMOpenQueue method.
	//
	// Return Values: The method MUST return MQ_OK (0x00000000) on success; otherwise, it
	// MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs identically.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_INVALID_HANDLE (0xC00E0007)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	CloseQueue(context.Context, *CloseQueueRequest) (*CloseQueueResponse, error)

	// The RemoteQMCloseCursor method closes the handle for a previously created cursor.
	// The client MUST call this method to reclaim resources on the server allocated by
	// the qmcomm:R_QMCreateRemoteCursor method, as specified in [MS-MQMP] section 3.1.4.4.
	//
	// Return Values:  The method MUST return MQ_OK (0x00000000) on success; otherwise,
	// it MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs
	// identically.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_INVALID_HANDLE (0xC00E0007)
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	CloseCursor(context.Context, *CloseCursorRequest) (*CloseCursorResponse, error)

	// The RemoteQMCancelReceive method cancels a pending call to RemoteQMStartReceive and
	// provides a way for the client to cancel a blocked request.
	//
	// Before calling this method, the following methods MUST be called:
	//
	// * RemoteQMOpenQueue ( 0dbcb6fd-ed56-44c3-9c02-e9fa2d7798b3 )
	//
	// * RemoteQMStartReceive or RemoteQMStartReceive2 ( a0df225e-80ea-4242-90e5-9891bfa9ebfb
	// )
	//
	// HRESULT RemoteQMCancelReceive(
	//
	// [in] handle_t hBind,
	//
	// [in] DWORD hQueue,
	//
	// [in] DWORD pQueue,
	//
	// [in] DWORD dwRequestID
	//
	// );
	//
	// Return Values:  The method MUST return MQ_OK (0x00000000) on success; otherwise,
	// it MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs
	// identically.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR (0xC00E0001)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// MQ_ERROR_INVALID_HANDLE (0xC00E0007)
	//
	// STATUS_NOT_FOUND (0xC0000225)
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	CancelReceive(context.Context, *CancelReceiveRequest) (*CancelReceiveResponse, error)

	// The RemoteQMPurgeQueue method removes all messages from the queue.
	//
	// Before calling this method, the RemoteQMOpenQueue method MUST be called.
	//
	// Return Values: The method MUST return MQ_OK (0x00000000) on success; otherwise, it
	// MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs identically.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_INVALID_HANDLE (0xC00E0007)
	//
	// Exceptions Thrown: Failure HRESULTs returned by this method MAY<18> be thrown as
	// exceptions as well as those thrown by the underlying RPC protocol, as specified in
	// [MS-RPCE].
	PurgeQueue(context.Context, *PurgeQueueRequest) (*PurgeQueueResponse, error)

	// The RemoteQMGetQMQMServerPort method returns an RPC port number (see [MS-RPCE]) for
	// the requested combination of interface and protocol.
	//
	// Return Values:  On success, this method returns a nonzero IP port value for the
	// RPC interface specified by the dwPortType parameter. If an invalid value is specified
	// for dwPortType, or if the requested interface is otherwise unavailable, or if any
	// other error is encountered, this method MUST return 0x00000000.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol (see [MS-RPCE]).
	//
	// As specified in section 3.1.3, this protocol configures a fixed listening endpoint
	// at an RPC port number, which can vary. For the interface and protocol specified by
	// the dwPortType parameter, this method returns the RPC port number determined at server
	// initialization time. If the default port is already in use, the server SHOULD increment
	// the port number by 11 until an unused port is found.
	GetQMQMServerPort(context.Context, *GetQMQMServerPortRequest) (*GetQMQMServerPortResponse, error)

	// The RemoteQmGetVersion method retrieves the Message queuing version of the server;
	// this method is called before the RemoteQMOpenQueue method.<21>
	//
	// Return Values: This method has no return values.
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)

	// The RemoteQMStartReceive2 method functions in the same way as RemoteQMStartReceive
	// (section 3.1.4.1), except that it returns a structure that contains the SequentialId
	// of the message.<25>
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically.
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	StartReceive2(context.Context, *StartReceive2Request) (*StartReceive2Response, error)

	// The RemoteQMStartReceiveByLookupId method reads a message from the opened remote
	// queue by using the lookup identifier.<27>
	//
	// Return Values: The method MUST return MQ_OK (0x00000000) on success; otherwise, it
	// MUST return a failure HRESULT and the client MUST treat all failure HRESULTs identically.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_INVALID_HANDLE (0xC00E0007)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// MQ_ERROR_IO_TIMEOUT ((0xC00E001B))
	//
	// MQ_ERROR_MESSAGE_ALREADY_RECEIVED ((0xC00E001D))
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	StartReceiveByLookupID(context.Context, *StartReceiveByLookupIDRequest) (*StartReceiveByLookupIDResponse, error)
}

func RegisterQm2qmServer(conn dcerpc.Conn, o Qm2qmServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQm2qmServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Qm2qmSyntaxV1_0))...)
}

func NewQm2qmServerHandle(o Qm2qmServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Qm2qmServerHandle(ctx, o, opNum, r)
	}
}

func Qm2qmServerHandle(ctx context.Context, o Qm2qmServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RemoteQMStartReceive
		op := &xxx_StartReceiveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartReceiveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartReceive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RemoteQMEndReceive
		op := &xxx_EndReceiveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EndReceiveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EndReceive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // RemoteQMOpenQueue
		op := &xxx_OpenQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // RemoteQMCloseQueue
		op := &xxx_CloseQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RemoteQMCloseCursor
		op := &xxx_CloseCursorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseCursorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseCursor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // RemoteQMCancelReceive
		op := &xxx_CancelReceiveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CancelReceiveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CancelReceive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // RemoteQMPurgeQueue
		op := &xxx_PurgeQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PurgeQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PurgeQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // RemoteQMGetQMQMServerPort
		op := &xxx_GetQMQMServerPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQMQMServerPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQMQMServerPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RemoteQmGetVersion
		op := &xxx_GetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RemoteQMStartReceive2
		op := &xxx_StartReceive2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartReceive2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartReceive2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RemoteQMStartReceiveByLookupId
		op := &xxx_StartReceiveByLookupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartReceiveByLookupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartReceiveByLookupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented qm2qm
type UnimplementedQm2qmServer struct {
}

func (UnimplementedQm2qmServer) StartReceive(context.Context, *StartReceiveRequest) (*StartReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQm2qmServer) EndReceive(context.Context, *EndReceiveRequest) (*EndReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQm2qmServer) OpenQueue(context.Context, *OpenQueueRequest) (*OpenQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQm2qmServer) CloseQueue(context.Context, *CloseQueueRequest) (*CloseQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQm2qmServer) CloseCursor(context.Context, *CloseCursorRequest) (*CloseCursorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQm2qmServer) CancelReceive(context.Context, *CancelReceiveRequest) (*CancelReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQm2qmServer) PurgeQueue(context.Context, *PurgeQueueRequest) (*PurgeQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQm2qmServer) GetQMQMServerPort(context.Context, *GetQMQMServerPortRequest) (*GetQMQMServerPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQm2qmServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQm2qmServer) StartReceive2(context.Context, *StartReceive2Request) (*StartReceive2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQm2qmServer) StartReceiveByLookupID(context.Context, *StartReceiveByLookupIDRequest) (*StartReceiveByLookupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Qm2qmServer = (*UnimplementedQm2qmServer)(nil)
