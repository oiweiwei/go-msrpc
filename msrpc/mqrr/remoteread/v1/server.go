package remoteread

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
)

// RemoteRead server interface.
type RemoteReadServer interface {

	// The R_GetServerPort method returns the RPC endpoint port for the client to use in
	// subsequent method calls on the RemoteRead interface.
	//
	// Return Values: On success, this method MUST return a nonzero TCP port value for the
	// RPC interface. If an error occurs, the server MUST return 0x00000000.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	//
	// As specified in section 3.1.3, this protocol configures a fixed listening endpoint
	// at an RPC port number that can vary. This method returns the RPC port number determined
	// at server initialization time.
	GetServerPort(context.Context, *GetServerPortRequest) (*GetServerPortResponse, error)

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// The R_OpenQueue method opens a queue in preparation for subsequent operations against
	// it. This method MUST be called prior to calling any of the following operations:
	//
	// * R_CreateCursor (Opnum 4) (section 3.1.4.4) ( 142d17a9-7fb3-481a-8d07-2170fd22e5ec
	// )
	//
	// * R_CloseCursor (Opnum 5) (section 3.1.4.5) ( dee84ebb-3d4a-4c26-ab11-3fa3b6e1acbe
	// )
	//
	// * R_PurgeQueue (Opnum 6) (section 3.1.4.6) ( 9807ff2d-be58-456c-ac96-be23be8561f4
	// )
	//
	// * R_StartReceive (Opnum 7) (section 3.1.4.7) ( 1bd20827-f774-4279-ae8d-4898d6faf6e9
	// )
	//
	// * R_CancelReceive (Opnum 8) (section 3.1.4.8) ( 471d411a-2757-478e-b121-ef24ad0543a5
	// )
	//
	// * R_EndReceive (Opnum 9) (section 3.1.4.9) ( 7e10ff5b-2991-4d73-b712-f3916bf4e85d
	// )
	//
	// * R_MoveMessage (Opnum 10) (section 3.1.4.10) ( 81c28eb5-5658-4089-b8ac-7edc5d629e28
	// ) for the source queue only.
	//
	// * R_StartTransactionalReceive (Opnum 13) (section 3.1.4.13) ( c5296e0c-9e12-4fe8-a25d-56924982ac3e
	// )
	//
	// * R_SetUserAcknowledgementClass (Opnum 14) (section 3.1.4.14) ( add5b3ed-7b02-48df-af90-a304ed0b9eea
	// )
	//
	// * R_EndTransactionalReceive (Opnum 15) (section 3.1.4.15) ( f15f7fba-9259-4641-ab05-08a0485f5a8a
	// )
	//
	// This method returns a QUEUE_CONTEXT_HANDLE_SERIALIZE (section 2.2.4.2) handle value,
	// which is required as input in the operations listed preceding.
	//
	// Return Values: The method has no return values. If the method fails, an RPC exception
	// is thrown.
	//
	// Exceptions Thrown:
	//
	// In addition to the exceptions thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE], the method throws HRESULT failure codes as RPC exceptions. The client
	// MUST treat all thrown HRESULT codes identically. The client MUST disregard all output
	// parameter values when any failure HRESULT is thrown.
	OpenQueue(context.Context, *OpenQueueRequest) (*OpenQueueResponse, error)

	// The R_CloseQueue method closes a QUEUE_CONTEXT_HANDLE_SERIALIZE (section 2.2.4.2)
	// handle that was previously opened by using a call to the R_OpenQueue (Opnum 2) (section
	// 3.1.4.2) method or the R_OpenQueueForMove (Opnum 11) (section 3.1.4.11) method.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	CloseQueue(context.Context, *CloseQueueRequest) (*CloseQueueResponse, error)

	// The R_CreateCursor method creates a cursor and returns a handle to it. The handle
	// can be used in subsequent calls to the R_StartReceive (Opnum 7) (section 3.1.4.7)
	// method or the R_StartTransactionalReceive (Opnum 13) (section 3.1.4.13) method to
	// specify a relative location in the queue from which to receive a message.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure, and the client MUST treat all
	// failure HRESULTs identically.
	//
	// The client MUST disregard all out-parameter values when any failure HRESULT is returned.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	CreateCursor(context.Context, *CreateCursorRequest) (*CreateCursorResponse, error)

	// The R_CloseCursor method closes the handle for a previously created cursor. The client
	// MUST call this method to reclaim resources on the server allocated by the R_CreateCursor
	// (Opnum 4) (section 3.1.4.4) method.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT and the client MUST
	// treat all failure HRESULTs identically.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol,
	// as specified in [MS-RPCE].
	CloseCursor(context.Context, *CloseCursorRequest) (*CloseCursorResponse, error)

	// The R_PurgeQueue method removes all messages from the queue.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	PurgeQueue(context.Context, *PurgeQueueRequest) (*PurgeQueueResponse, error)

	// The R_StartReceive method peeks or receives a message from an open queue.
	//
	// If the R_StartReceive method is invoked with a Peek action type, as specified in
	// the ulAction parameter, the operation completes when the R_StartReceive method returns.
	//
	// If the R_StartReceive method is invoked with a Receive action type, as specified
	// in the ulAction parameter, the client MUST pair each call to the R_StartReceive method
	// with a call to the R_EndReceive (Opnum 9) (section 3.1.4.9) method to complete the
	// operation, or to the R_CancelReceive (Opnum 8) (section 3.1.4.8) method to cancel
	// the operation. The call to the R_EndReceive method or the R_CancelReceive method
	// is correlated to a call to the R_StartReceive method through matching dwRequestId
	// parameters.
	//
	// If the client specifies a nonzero ulTimeout parameter, and a message is not available
	// in the queue at the time of the call, the server waits up to the specified time-out
	// for a message to become available in the queue before responding to the call. The
	// client can call the R_CancelReceive method with a matching dwRequestId parameter
	// to cancel the pending R_StartReceive method request.
	//
	// The message to be returned can be specified in one of three ways:
	//
	// * LookupId: A nonzero LookupId parameter value specifies the unique identifier for
	// the message to be returned. The ulAction parameter further specifies whether the
	// message to be returned is the one identified by the LookupId parameter or the first
	// unlocked message immediately preceding or following it. For more details, see the
	// description of the ulAction parameter.
	//
	// * Cursor: A nonzero cursor ( b23ea276-dfad-4083-bc48-d0f8a40255fd#gt_aa5e9c2d-16c1-4301-8bfe-18a0913ed275
	// ) handle specifies the cursor to be used to identify the message to be returned.
	// The cursor specifies a location in the queue. The ulAction parameter further specifies
	// whether the message to be returned is the one identified by the cursor or the first
	// unlocked message immediately following it. For more details, see the description
	// of the ulAction parameter.
	//
	// * First: if the LookupId parameter is set to zero and the hCursor parameter is set
	// to zero, the first unlocked message in the queue can be returned. The ulAction parameter
	// further specifies whether the first message is to be received or peeked.
	//
	// The ppPacketSections parameter is the address of one or more pointers to one or more
	// SectionBuffer (section 2.2.6) structures. The pSectionBuffer member of the first
	// SectionBuffer structure points to the beginning of the message packet. If more than
	// one SectionBuffer structure is present, the packet sections are concatenated in the
	// order in which they appear in the array to form the entire packet. The size of each
	// section is stored in the SectionSizeAlloc member of the SectionBuffer structure.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically. The client MUST disregard all output parameter
	// values when any failure HRESULT is returned. For descriptions of the following error
	// codes, see [MS-MQMQ] section 2.4. For error codes not described in [MS-MQMQ], refer
	// to [MSDN-MQEIC].
	//
	//	+----------------------------------------------+-------------+
	//	|                    RETURN                    |             |
	//	|                  VALUE/CODE                  | DESCRIPTION |
	//	|                                              |             |
	//	+----------------------------------------------+-------------+
	//	+----------------------------------------------+-------------+
	//	| 0x00000000 MQ_OK                             |             |
	//	+----------------------------------------------+-------------+
	//	| 0xC00E0007 MQ_ERROR_INVALID_HANDLE           |             |
	//	+----------------------------------------------+-------------+
	//	| 0xC00E001B MQ_ERROR_IO_TIMEOUT               |             |
	//	+----------------------------------------------+-------------+
	//	| 0xC00E0088 MQ_ERROR_MESSAGE_NOT_FOUND        |             |
	//	+----------------------------------------------+-------------+
	//	| 0xC00E001D MQ_ERROR_MESSAGE_ALREADY_RECEIVED |             |
	//	+----------------------------------------------+-------------+
	//	| 0xC00E0008 MQ_ERROR_OPERATION_CANCELLED      |             |
	//	+----------------------------------------------+-------------+
	//	| 0xC00E0006 MQ_ERROR_INVALID_PARAMETER        |             |
	//	+----------------------------------------------+-------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	StartReceive(context.Context, *StartReceiveRequest) (*StartReceiveResponse, error)

	// The R_CancelReceive method cancels a pending call to the R_StartReceive (Opnum 7)
	// (section 3.1.4.7) method or the R_StartTransactionalReceive (Opnum 13) (section 3.1.4.13)
	// method. Each of those methods takes a time-out parameter that can cause the server
	// to not return a response until a message becomes available or the time-out expires.
	// The R_CancelReceive method provides a way for the client to cancel a blocked request.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	CancelReceive(context.Context, *CancelReceiveRequest) (*CancelReceiveResponse, error)

	// The client MUST invoke the R_EndReceive method to advise the server that the message
	// packet returned by the R_StartReceive (Opnum 7) (section 3.1.4.7) method has been
	// received.
	//
	// The combination of the R_StartReceive method and the positive acknowledgment of the
	// R_EndReceive method ensures that a message packet is not lost in transit from the
	// server to the client due to a network outage during the call sequence.
	//
	// Note that a call to the R_StartTransactionalReceive (Opnum 13) (section 3.1.4.13)
	// method is ended through a corresponding call to the R_EndTransactionalReceive (Opnum
	// 15) (section 3.1.4.15) method, not through a call to this method.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol [MS-RPCE].
	EndReceive(context.Context, *EndReceiveRequest) (*EndReceiveResponse, error)

	// The R_MoveMessage method moves a message from one queue to another.<20> The source
	// and destination queues MUST be related as follows:
	//
	// * The source is a queue, and the destination is a subqueue ( b23ea276-dfad-4083-bc48-d0f8a40255fd#gt_328fe8aa-d006-44dd-8cc8-dba7c862aaf8
	// ) of the source queue, or
	//
	// * The destination is a queue, and the source is a subqueue of the destination queue,
	// or
	//
	// * The source and destination are two subqueues of the same queue.
	//
	// HRESULT R_MoveMessage(
	//
	// [in] handle_t hBind,
	//
	// [in] QUEUE_CONTEXT_HANDLE_NOSERIALIZE phContextFrom,
	//
	// [in] ULONGLONG ullContextTo,
	//
	// [in] ULONGLONG LookupId,
	//
	// * );
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client treat
	// all failure HRESULTs identically.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	//
	// The R_MoveMessage method provides both transactional and non-transactional operations.
	// When using a transaction identifier, this method provisionally moves a message from
	// the source queue to the destination queue, pending notification of the transaction
	// outcome. See section 3.1.6. The non-transactional operation moves a message from
	// the source queue to the destination queue without enlisting in a transaction.
	MoveMessage(context.Context, *MoveMessageRequest) (*MoveMessageResponse, error)

	// The R_OpenQueueForMove method opens the queue and returns a QUEUE_CONTEXT_HANDLE_SERIALIZE
	// (section 2.2.4.2) handle that can subsequently be used as the ullContextTo (destination
	// queue) parameter of a call to the R_MoveMessage (Opnum 10) (section 3.1.4.10) method.
	// This method MUST be called before the R_MoveMessage method.<21>
	//
	// Return Values: The method has no return values. If the method fails, an RPC exception
	// is thrown.
	//
	// Exceptions Thrown:
	//
	// In addition to the exceptions thrown by the underlying RPC protocol [MS-RPCE], the
	// method throws HRESULT failure codes as RPC exceptions. The client MUST treat all
	// thrown HRESULT codes identically.
	//
	// The client MUST disregard all out-parameter values when any failure HRESULT is thrown.
	OpenQueueForMove(context.Context, *OpenQueueForMoveRequest) (*OpenQueueForMoveResponse, error)

	// The R_QMEnlistRemoteTransaction method propagates a distributed atomic transaction
	// context to the server. The server MUST enlist in the transaction context. The client
	// MUST call this method prior to the R_StartTransactionalReceive (Opnum 13) (section
	// 3.1.4.13) method or the R_MoveMessage (Opnum 10) (section 3.1.4.10) method calls.<27>
	// Subsequent calls to the R_StartTransactionalReceive method and the R_MoveMessage
	// method that use the same transaction identifier are coordinated such that either
	// all occur or none occurs, depending on whether the transaction outcome is Commit
	// or Rollback.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol,
	// as specified in [MS-RPCE].
	EnlistRemoteTransaction(context.Context, *EnlistRemoteTransactionRequest) (*EnlistRemoteTransactionResponse, error)

	// The R_StartTransactionalReceive method peeks or receives a message from the opened
	// queue.<29> If a transaction identifier is provided, a message is received inside
	// the specified transaction.
	//
	// If the R_StartTransactionalReceive method is invoked with a Peek action type, as
	// specified in the ulAction parameter, the operation completes when the R_StartTransactionalReceive
	// method returns.
	//
	// If the R_StartTransactionalReceive method is invoked with a Receive action type,
	// as specified in the ulAction parameter, the client MUST pair each call to the R_StartTransactionalReceive
	// method with a call to the R_EndTransactionalReceive (Opnum 15) (section 3.1.4.15)
	// method to complete the operation or to the R_CancelReceive (Opnum 8) (section 3.1.4.8)
	// method to cancel the operation. The call to the R_EndTransactionalReceive method
	// or the R_CancelReceive method is correlated to a call to the R_StartTransactionalReceive
	// method through matching dwRequestId parameters.
	//
	// If the client specifies a nonzero ulTimeout parameter and a message is not available
	// in the queue at the time of the call, the server waits up to the specified time-out
	// for a message to become available in the queue before responding to the call. The
	// client can call the R_CancelReceive method with a matching dwRequestId parameter
	// to cancel the pending R_StartTransactionalReceive method request.
	//
	// The message to be returned can be specified in one of three ways:
	//
	// * LookupId: A nonzero LookupId parameter value that specifies the unique identifier
	// for the message to be returned. The ulAction parameter further specifies whether
	// the message to be returned is the one identified by the LookupId parameter or the
	// first unlocked message immediately preceding or following it. For more details, see
	// the description of the ulAction parameter.
	//
	// * Cursor: A nonzero cursor ( b23ea276-dfad-4083-bc48-d0f8a40255fd#gt_aa5e9c2d-16c1-4301-8bfe-18a0913ed275
	// ) handle that specifies the cursor to be used to identify the message to be returned.
	// The cursor specifies a location in the queue. The ulAction parameter further specifies
	// whether the message to be returned is the one identified by the cursor or the first
	// unlocked message immediately following it. For more details, see the description
	// of the ulAction parameter.
	//
	// * First: If the LookupId parameter is set to 0x0000000000000000 and hCursor is set
	// to 0x00000000, the first unlocked message in the queue can be returned. For more
	// details, see the description of the ulAction parameter.
	//
	// The ppPacketSections parameter is the address of one or more pointers to one or more
	// SectionBuffer (section 2.2.6) structures. The pSectionBuffer member of the first
	// SectionBuffer structure points to the beginning of the message packet. If more than
	// one SectionBuffer structure is present, the packet sections are concatenated in the
	// order in which they appear in the array to form the entire packet. The size of each
	// section is stored in the SectionSizeAlloc member of the SectionBuffer structure.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically. The client MUST disregard all output parameter
	// values when any failure HRESULT is returned.
	//
	//	+-----------------------------------------+-------------+
	//	|                 RETURN                  |             |
	//	|               VALUE/CODE                | DESCRIPTION |
	//	|                                         |             |
	//	+-----------------------------------------+-------------+
	//	+-----------------------------------------+-------------+
	//	| 0x00000000 MQ_OK                        |             |
	//	+-----------------------------------------+-------------+
	//	| 0xC00E0007 MQ_ERROR_INVALID_HANDLE      |             |
	//	+-----------------------------------------+-------------+
	//	| 0xC00E0088 MQ_ERROR_MESSAGE_NOT_FOUND   |             |
	//	+-----------------------------------------+-------------+
	//	| 0xC00E001B MQ_ERROR_IO_TIMEOUT          |             |
	//	+-----------------------------------------+-------------+
	//	| 0xC00E0050 MQ_ERROR_TRANSACTION_USAGE   |             |
	//	+-----------------------------------------+-------------+
	//	| 0xC00E0008 MQ_ERROR_OPERATION_CANCELLED |             |
	//	+-----------------------------------------+-------------+
	//	| 0xC00E0006 MQ_ERROR_INVALID_PARAMETER   |             |
	//	+-----------------------------------------+-------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	StartTransactionalReceive(context.Context, *StartTransactionalReceiveRequest) (*StartTransactionalReceiveResponse, error)

	// The R_SetUserAcknowledgementClass method sets the acknowledgment class property of
	// a message in the queue. This allows marking the message as rejected.<30> This method
	// MUST be called subsequent to calls to R_StartTransactionalReceive and R_EndTransactionalReceive
	// (Opnum 15) (section 3.1.4.15) and before the transaction is committed or aborted.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_INVALID_HANDLE (0xC00E0007)
	//
	// MQ_ERROR_TRANSACTION_USAGE (0xC00E0050)
	//
	// MQ_ERROR_MESSAGE_NOT_FOUND (0xC00E0088)
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	SetUserAcknowledgementClass(context.Context, *SetUserAcknowledgementClassRequest) (*SetUserAcknowledgementClassResponse, error)

	// The client MUST invoke the R_EndTransactionalReceive method to advise the server
	// that the message packet returned by the R_StartTransactionalReceive (Opnum 13) (section
	// 3.1.4.13) method has been received by the client.<31>
	//
	// The combination of the R_StartTransactionalReceive method and the positive acknowledgment
	// of the R_EndTransactionalReceive method ensures that a message packet is not lost
	// in transit from the server to the client due to a network outage during the call
	// sequence.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If an error occurs, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown except those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	EndTransactionalReceive(context.Context, *EndTransactionalReceiveRequest) (*EndTransactionalReceiveResponse, error)
}

func RegisterRemoteReadServer(conn dcerpc.Conn, o RemoteReadServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteReadServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteReadSyntaxV1_0))...)
}

func NewRemoteReadServerHandle(o RemoteReadServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteReadServerHandle(ctx, o, opNum, r)
	}
}

func RemoteReadServerHandle(ctx context.Context, o RemoteReadServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // R_GetServerPort
		in := &GetServerPortRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServerPort(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // Opnum1NotUsedOnWire
		// Opnum1NotUsedOnWire
		return nil, nil
	case 2: // R_OpenQueue
		in := &OpenQueueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenQueue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // R_CloseQueue
		in := &CloseQueueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseQueue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // R_CreateCursor
		in := &CreateCursorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateCursor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // R_CloseCursor
		in := &CloseCursorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseCursor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // R_PurgeQueue
		in := &PurgeQueueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PurgeQueue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // R_StartReceive
		in := &StartReceiveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StartReceive(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // R_CancelReceive
		in := &CancelReceiveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CancelReceive(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // R_EndReceive
		in := &EndReceiveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EndReceive(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // R_MoveMessage
		in := &MoveMessageRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MoveMessage(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // R_OpenQueueForMove
		in := &OpenQueueForMoveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenQueueForMove(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // R_QMEnlistRemoteTransaction
		in := &EnlistRemoteTransactionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnlistRemoteTransaction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // R_StartTransactionalReceive
		in := &StartTransactionalReceiveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StartTransactionalReceive(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // R_SetUserAcknowledgementClass
		in := &SetUserAcknowledgementClassRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetUserAcknowledgementClass(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // R_EndTransactionalReceive
		in := &EndTransactionalReceiveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EndTransactionalReceive(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
