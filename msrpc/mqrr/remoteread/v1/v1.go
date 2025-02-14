package remoteread

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	mqmq "github.com/oiweiwei/go-msrpc/msrpc/mqmq"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
	_ = mqmq.GoPackage
)

var (
	// import guard
	GoPackage = "mqrr"
)

var (
	// Syntax UUID
	RemoteReadSyntaxUUID = &uuid.UUID{TimeLow: 0x1a9134dd, TimeMid: 0x7b39, TimeHiAndVersion: 0x45ba, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0x88, Node: [6]uint8{0x44, 0xd0, 0x1c, 0xa4, 0x7f, 0x28}}
	// Syntax ID
	RemoteReadSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: RemoteReadSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// RemoteRead interface.
type RemoteReadClient interface {

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
	GetServerPort(context.Context, *GetServerPortRequest, ...dcerpc.CallOption) (*GetServerPortResponse, error)

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
	OpenQueue(context.Context, *OpenQueueRequest, ...dcerpc.CallOption) (*OpenQueueResponse, error)

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
	CloseQueue(context.Context, *CloseQueueRequest, ...dcerpc.CallOption) (*CloseQueueResponse, error)

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
	CreateCursor(context.Context, *CreateCursorRequest, ...dcerpc.CallOption) (*CreateCursorResponse, error)

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
	CloseCursor(context.Context, *CloseCursorRequest, ...dcerpc.CallOption) (*CloseCursorResponse, error)

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
	PurgeQueue(context.Context, *PurgeQueueRequest, ...dcerpc.CallOption) (*PurgeQueueResponse, error)

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
	StartReceive(context.Context, *StartReceiveRequest, ...dcerpc.CallOption) (*StartReceiveResponse, error)

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
	CancelReceive(context.Context, *CancelReceiveRequest, ...dcerpc.CallOption) (*CancelReceiveResponse, error)

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
	EndReceive(context.Context, *EndReceiveRequest, ...dcerpc.CallOption) (*EndReceiveResponse, error)

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
	MoveMessage(context.Context, *MoveMessageRequest, ...dcerpc.CallOption) (*MoveMessageResponse, error)

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
	OpenQueueForMove(context.Context, *OpenQueueForMoveRequest, ...dcerpc.CallOption) (*OpenQueueForMoveResponse, error)

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
	EnlistRemoteTransaction(context.Context, *EnlistRemoteTransactionRequest, ...dcerpc.CallOption) (*EnlistRemoteTransactionResponse, error)

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
	StartTransactionalReceive(context.Context, *StartTransactionalReceiveRequest, ...dcerpc.CallOption) (*StartTransactionalReceiveResponse, error)

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
	SetUserAcknowledgementClass(context.Context, *SetUserAcknowledgementClassRequest, ...dcerpc.CallOption) (*SetUserAcknowledgementClassResponse, error)

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
	EndTransactionalReceive(context.Context, *EndTransactionalReceiveRequest, ...dcerpc.CallOption) (*EndTransactionalReceiveResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// QueueNoSerialize structure represents QUEUE_CONTEXT_HANDLE_NOSERIALIZE RPC structure.
type QueueNoSerialize dcetypes.ContextHandle

func (o *QueueNoSerialize) ContextHandle() *dcetypes.ContextHandle {
	return (*dcetypes.ContextHandle)(o)
}

func (o *QueueNoSerialize) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueNoSerialize) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueNoSerialize) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// QueueSerialize structure represents QUEUE_CONTEXT_HANDLE_SERIALIZE RPC structure.
type QueueSerialize dcetypes.ContextHandle

func (o *QueueSerialize) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *QueueSerialize) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueSerialize) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueSerialize) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SectionType type represents SectionType RPC enumeration.
//
// The SectionType enumeration defines the available SectionBuffer (section 2.2.6) types.
type SectionType uint16

var (
	// stFullPacket:  The pSectionBuffer member of the SectionBuffer structure contains
	// a complete Message Packet Structure (section 2.2.5). The UserMessage (section 2.2.5.1)
	// is either that specified in section 2.2.5.1.1 or in section 2.2.5.1.2.
	SectionTypeFullPacket SectionType = 0
	// stBinaryFirstSection:  The pSectionBuffer member of the SectionBuffer structure
	// contains the first section of the Binary Message (section 2.2.5.1.1) packet up to,
	// but not beyond, the MessagePropertiesHeader ([MS-MQMQ] section 2.2.19.3) in the UserMessage.
	SectionTypeBinaryFirstSection SectionType = 1
	// stBinarySecondSection:  The pSectionBuffer member of the SectionBuffer structure
	// contains the second section of the Binary Message packet from beyond the end of the
	// MessagePropertiesHeader in the UserMessage to the end of the packet.
	SectionTypeBinarySecondSection SectionType = 2
	// stSrmpFirstSection:  The pSectionBuffer member of the SectionBuffer structure contains
	// the first section of the SRMP Message packet up to, but not beyond, the CompoundMessageHeader
	// (section 2.2.5.1.2.2) in the UserMessage.
	SectionTypeSrmpFirstSection SectionType = 3
	// stSrmpSecondSection:  The pSectionBuffer member of the SectionBuffer structure contains
	// the second section of the SRMP Message packet from beyond the end of the CompoundMessageHeader
	// in the UserMessage to the end of the packet.
	SectionTypeSrmpSecondSection SectionType = 4
)

func (o SectionType) String() string {
	switch o {
	case SectionTypeFullPacket:
		return "SectionTypeFullPacket"
	case SectionTypeBinaryFirstSection:
		return "SectionTypeBinaryFirstSection"
	case SectionTypeBinarySecondSection:
		return "SectionTypeBinarySecondSection"
	case SectionTypeSrmpFirstSection:
		return "SectionTypeSrmpFirstSection"
	case SectionTypeSrmpSecondSection:
		return "SectionTypeSrmpSecondSection"
	}
	return "Invalid"
}

// SectionBuffer structure represents SectionBuffer RPC structure.
//
// A SectionBuffer represents a fragment or section of a Message Packet (section 2.2.5).
// Operations R_StartReceive (Opnum 7) (section 3.1.4.7) and R_StartTransactionalReceive
// (Opnum 13) (section 3.1.4.13) fragment a Message Packet into an array of one or more
// SectionBuffer structures. The client concatenates these fragments to reconstruct
// a valid Message Packet. There can be up to two sections per message. A Message Packet
// is split into two sections only when a subset of the distinguished message body property
// is returned. The first section always contains the message body property up to the
// size requested.
type SectionBuffer struct {
	// SectionBufferType:  MUST specify a type for the SectionBuffer structure that indicates
	// whether the pSectionBuffer member contains the whole Message Packet or MUST indicate
	// which section is contained. The SectionType (section 2.2.7) enumeration lists possible
	// values. More details are specified in 2.2.7.
	SectionBufferType SectionType `idl:"name:SectionBufferType" json:"section_buffer_type"`
	// SectionSizeAlloc:  MUST specify the original size (in bytes) of the part of the Message
	// Packet that this SectionBuffer represents. When the SectionBuffer represents the
	// first section of the message, this field specifies the size that the SectionBuffer
	// would have been if the entire message body property were included. The difference
	// between the values of the SectionSizeAlloc member and the SectionSize member represents
	// the size of the message body that was not transferred.
	//
	// If the SectionBufferType member value is stFullPacket, stBinarySecondSection, or
	// stSrmpSecondSection, then the SectionSizeAlloc member value MUST be equal to the
	// SectionSize member value.
	SectionSizeAlloc uint32 `idl:"name:SectionSizeAlloc" json:"section_size_alloc"`
	// SectionSize:  MUST be the size (in bytes) of the buffer pointed to by the pSectionBuffer
	// member. The SectionSize member specifies the size of the part of the Message Packet
	// contained in the pSectionBuffer member.
	SectionSize uint32 `idl:"name:SectionSize" json:"section_size"`
	// pSectionBuffer:  MUST be a pointer to an array of bytes containing a section of the
	// Message Packet.
	SectionBuffer []byte `idl:"name:pSectionBuffer;size_is:(SectionSize);pointer:unique" json:"section_buffer"`
}

func (o *SectionBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.SectionBuffer != nil && o.SectionSize == 0 {
		o.SectionSize = uint32(len(o.SectionBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SectionBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.SectionBufferType)); err != nil {
		return err
	}
	if err := w.WriteData(o.SectionSizeAlloc); err != nil {
		return err
	}
	if err := w.WriteData(o.SectionSize); err != nil {
		return err
	}
	if o.SectionBuffer != nil || o.SectionSize > 0 {
		_ptr_pSectionBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SectionSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SectionBuffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.SectionBuffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.SectionBuffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SectionBuffer, _ptr_pSectionBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SectionBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.SectionBufferType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.SectionSizeAlloc); err != nil {
		return err
	}
	if err := w.ReadData(&o.SectionSize); err != nil {
		return err
	}
	_ptr_pSectionBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SectionSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SectionSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SectionBuffer", sizeInfo[0])
		}
		o.SectionBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.SectionBuffer {
			i1 := i1
			if err := w.ReadData(&o.SectionBuffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pSectionBuffer := func(ptr interface{}) { o.SectionBuffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SectionBuffer, _s_pSectionBuffer, _ptr_pSectionBuffer); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultRemoteReadClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultRemoteReadClient) GetServerPort(ctx context.Context, in *GetServerPortRequest, opts ...dcerpc.CallOption) (*GetServerPortResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetServerPortResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) OpenQueue(ctx context.Context, in *OpenQueueRequest, opts ...dcerpc.CallOption) (*OpenQueueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenQueueResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) CloseQueue(ctx context.Context, in *CloseQueueRequest, opts ...dcerpc.CallOption) (*CloseQueueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseQueueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) CreateCursor(ctx context.Context, in *CreateCursorRequest, opts ...dcerpc.CallOption) (*CreateCursorResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateCursorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) CloseCursor(ctx context.Context, in *CloseCursorRequest, opts ...dcerpc.CallOption) (*CloseCursorResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseCursorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) PurgeQueue(ctx context.Context, in *PurgeQueueRequest, opts ...dcerpc.CallOption) (*PurgeQueueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PurgeQueueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) StartReceive(ctx context.Context, in *StartReceiveRequest, opts ...dcerpc.CallOption) (*StartReceiveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StartReceiveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) CancelReceive(ctx context.Context, in *CancelReceiveRequest, opts ...dcerpc.CallOption) (*CancelReceiveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CancelReceiveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) EndReceive(ctx context.Context, in *EndReceiveRequest, opts ...dcerpc.CallOption) (*EndReceiveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EndReceiveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) MoveMessage(ctx context.Context, in *MoveMessageRequest, opts ...dcerpc.CallOption) (*MoveMessageResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MoveMessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) OpenQueueForMove(ctx context.Context, in *OpenQueueForMoveRequest, opts ...dcerpc.CallOption) (*OpenQueueForMoveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenQueueForMoveResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) EnlistRemoteTransaction(ctx context.Context, in *EnlistRemoteTransactionRequest, opts ...dcerpc.CallOption) (*EnlistRemoteTransactionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnlistRemoteTransactionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) StartTransactionalReceive(ctx context.Context, in *StartTransactionalReceiveRequest, opts ...dcerpc.CallOption) (*StartTransactionalReceiveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StartTransactionalReceiveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) SetUserAcknowledgementClass(ctx context.Context, in *SetUserAcknowledgementClassRequest, opts ...dcerpc.CallOption) (*SetUserAcknowledgementClassResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetUserAcknowledgementClassResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) EndTransactionalReceive(ctx context.Context, in *EndTransactionalReceiveRequest, opts ...dcerpc.CallOption) (*EndTransactionalReceiveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EndTransactionalReceiveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteReadClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteReadClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewRemoteReadClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteReadClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteReadSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultRemoteReadClient{cc: cc}, nil
}

// xxx_GetServerPortOperation structure represents the R_GetServerPort operation
type xxx_GetServerPortOperation struct {
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServerPortOperation) OpNum() int { return 0 }

func (o *xxx_GetServerPortOperation) OpName() string { return "/RemoteRead/v1/R_GetServerPort" }

func (o *xxx_GetServerPortOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerPortOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetServerPortOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetServerPortOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerPortOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerPortOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetServerPortRequest structure represents the R_GetServerPort operation request
type GetServerPortRequest struct {
}

func (o *GetServerPortRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServerPortOperation) *xxx_GetServerPortOperation {
	if op == nil {
		op = &xxx_GetServerPortOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *GetServerPortRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServerPortOperation) {
	if o == nil {
		return
	}
}
func (o *GetServerPortRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServerPortRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerPortOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServerPortResponse structure represents the R_GetServerPort operation response
type GetServerPortResponse struct {
	// Return: The R_GetServerPort return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetServerPortResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServerPortOperation) *xxx_GetServerPortOperation {
	if op == nil {
		op = &xxx_GetServerPortOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *GetServerPortResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServerPortOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *GetServerPortResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServerPortResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerPortOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenQueueOperation structure represents the R_OpenQueue operation
type xxx_OpenQueueOperation struct {
	QueueFormat      *mqmq.QueueFormat `idl:"name:pQueueFormat" json:"queue_format"`
	Access           uint32            `idl:"name:dwAccess" json:"access"`
	ShareMode        uint32            `idl:"name:dwShareMode" json:"share_mode"`
	ClientID         *dtyp.GUID        `idl:"name:pClientId" json:"client_id"`
	NonRoutingServer int32             `idl:"name:fNonRoutingServer" json:"non_routing_server"`
	Major            uint8             `idl:"name:Major" json:"major"`
	Minor            uint8             `idl:"name:Minor" json:"minor"`
	BuildNumber      uint16            `idl:"name:BuildNumber" json:"build_number"`
	Workgroup        int32             `idl:"name:fWorkgroup" json:"workgroup"`
	Context          *QueueSerialize   `idl:"name:pphContext" json:"context"`
}

func (o *xxx_OpenQueueOperation) OpNum() int { return 2 }

func (o *xxx_OpenQueueOperation) OpName() string { return "/RemoteRead/v1/R_OpenQueue" }

func (o *xxx_OpenQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pQueueFormat {in} (1:{pointer=ref}*(1))(2:{alias=QUEUE_FORMAT}(struct))
	{
		if o.QueueFormat != nil {
			if err := o.QueueFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqmq.QueueFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwAccess {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Access); err != nil {
			return err
		}
	}
	// dwShareMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ShareMode); err != nil {
			return err
		}
	}
	// pClientId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClientID != nil {
			if err := o.ClientID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// fNonRoutingServer {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.NonRoutingServer); err != nil {
			return err
		}
	}
	// Major {in} (1:(uchar))
	{
		if err := w.WriteData(o.Major); err != nil {
			return err
		}
	}
	// Minor {in} (1:(uchar))
	{
		if err := w.WriteData(o.Minor); err != nil {
			return err
		}
	}
	// BuildNumber {in} (1:{alias=USHORT}(uint16))
	{
		if err := w.WriteData(o.BuildNumber); err != nil {
			return err
		}
	}
	// fWorkgroup {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Workgroup); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pQueueFormat {in} (1:{pointer=ref}*(1))(2:{alias=QUEUE_FORMAT}(struct))
	{
		if o.QueueFormat == nil {
			o.QueueFormat = &mqmq.QueueFormat{}
		}
		if err := o.QueueFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwAccess {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Access); err != nil {
			return err
		}
	}
	// dwShareMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ShareMode); err != nil {
			return err
		}
	}
	// pClientId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClientID == nil {
			o.ClientID = &dtyp.GUID{}
		}
		if err := o.ClientID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// fNonRoutingServer {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.NonRoutingServer); err != nil {
			return err
		}
	}
	// Major {in} (1:(uchar))
	{
		if err := w.ReadData(&o.Major); err != nil {
			return err
		}
	}
	// Minor {in} (1:(uchar))
	{
		if err := w.ReadData(&o.Minor); err != nil {
			return err
		}
	}
	// BuildNumber {in} (1:{alias=USHORT}(uint16))
	{
		if err := w.ReadData(&o.BuildNumber); err != nil {
			return err
		}
	}
	// fWorkgroup {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Workgroup); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=QUEUE_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OpenQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=QUEUE_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OpenQueueRequest structure represents the R_OpenQueue operation request
type OpenQueueRequest struct {
	// pQueueFormat: MUST be a pointer to a QUEUE_FORMAT ([MS-MQMQ] section 2.2.7) structure
	// that identifies the queue to open. NULL is invalid for this parameter. The valid
	// values for the m_qft member are QUEUE_FORMAT_TYPE_PUBLIC, QUEUE_FORMAT_TYPE_PRIVATE,
	// QUEUE_FORMAT_TYPE_DIRECT, QUEUE_FORMAT_TYPE_MACHINE, and QUEUE_FORMAT_TYPE_SUBQUEUE.
	QueueFormat *mqmq.QueueFormat `idl:"name:pQueueFormat" json:"queue_format"`
	// dwAccess: Specifies the requested type of access to the queue. The required dwAccess
	// parameter value for each event is specified in each of the corresponding events.
	// If no requirement is listed, any dwAccess parameter value is accepted.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|                           |                                                                                  |
	//	|           VALUE           |                                     MEANING                                      |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| RECEIVE_ACCESS 0x00000001 | The returned QUEUE_CONTEXT_HANDLE_SERIALIZE handle can be used in the            |
	//	|                           | R_StartReceive or R_StartTransactionalReceive methods with the ulAction          |
	//	|                           | parameter set to either a Peek or Receive action type as defined in the table    |
	//	|                           | under the ulAction parameter in the R_StartReceive method.                       |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| PEEK_ACCESS 0x00000020    | The returned QUEUE_CONTEXT_HANDLE_SERIALIZE handle can be used in the            |
	//	|                           | R_StartReceive method with the ulAction parameter set only to a Peek action      |
	//	|                           | type as defined in the table under the ulAction parameter in the R_StartReceive  |
	//	|                           | method.                                                                          |
	//	+---------------------------+----------------------------------------------------------------------------------+
	Access uint32 `idl:"name:dwAccess" json:"access"`
	// dwShareMode: Specifies whether the client needs exclusive access to the queue. The
	// following values are valid for this parameter:
	//
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	|                          |                                                                                  |
	//	|          VALUE           |                                     MEANING                                      |
	//	|                          |                                                                                  |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| MQ_DENY_NONE 0x00000000  | Permits multiple QUEUE_CONTEXT_HANDLE_SERIALIZE handles to the queue to be       |
	//	|                          | opened concurrently.                                                             |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| MQ_DENY_SHARE 0x00000001 | Permits a single QUEUE_CONTEXT_HANDLE_SERIALIZE handle to the queue at a time,   |
	//	|                          | providing exclusive access to the queue.                                         |
	//	+--------------------------+----------------------------------------------------------------------------------+
	ShareMode uint32 `idl:"name:dwShareMode" json:"share_mode"`
	// pClientId: MUST be set by the client to a pointer to a valid GUID that uniquely identifies
	// the client. When the queue manager acts as the client, the queue manager sets this
	// value to the Identifier ADM attribute of the local QueueManager ADM element instance.
	// The server SHOULD ignore this parameter. The server MAY<14> use this parameter to
	// impose a limit on the number of unique callers. NULL is invalid for this parameter.
	ClientID *dtyp.GUID `idl:"name:pClientId" json:"client_id"`
	// fNonRoutingServer: If the client is configured to operate in the role of an MSMQ
	// routing server, this parameter MUST be set to FALSE (0x00000000); otherwise, it MUST
	// be set to TRUE (0x00000001).<15> If the value of the fNonRoutingServer parameter
	// is FALSE (0x00000000), the server MUST ignore the pClientId parameter.
	//
	//	+-------+------------+
	//	|       |            |
	//	| NAME  |   VALUE    |
	//	|       |            |
	//	+-------+------------+
	//	+-------+------------+
	//	| False | 0x00000000 |
	//	+-------+------------+
	//	| True  | 0x00000001 |
	//	+-------+------------+
	NonRoutingServer int32 `idl:"name:fNonRoutingServer" json:"non_routing_server"`
	// Major: MUST be set by the client to an implementation-specific Major Version number
	// of the client. SHOULD<16> be ignored by the server.
	Major uint8 `idl:"name:Major" json:"major"`
	// Minor: MUST be set by the client to an implementation-specific Minor Version number
	// of the client. SHOULD<17> be ignored by the server.
	Minor uint8 `idl:"name:Minor" json:"minor"`
	// BuildNumber: MUST be set by the client to an implementation-specific Build Number
	// of the client. SHOULD<18> be ignored by the server.
	BuildNumber uint16 `idl:"name:BuildNumber" json:"build_number"`
	// fWorkgroup: MUST be set to TRUE (0x00000001) by the client if the client machine
	// is not a member of a Windows domain; otherwise, it MUST be set to FALSE (0x00000000).
	// The RPC authentication level required by the server MAY<19> be based on this value
	// in subsequent calls on the interface.
	//
	//	+-------+------------+
	//	|       |            |
	//	| NAME  |   VALUE    |
	//	|       |            |
	//	+-------+------------+
	//	+-------+------------+
	//	| False | 0x00000000 |
	//	+-------+------------+
	//	| True  | 0x00000001 |
	//	+-------+------------+
	Workgroup int32 `idl:"name:fWorkgroup" json:"workgroup"`
}

func (o *OpenQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenQueueOperation) *xxx_OpenQueueOperation {
	if op == nil {
		op = &xxx_OpenQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.QueueFormat = o.QueueFormat
	op.Access = o.Access
	op.ShareMode = o.ShareMode
	op.ClientID = o.ClientID
	op.NonRoutingServer = o.NonRoutingServer
	op.Major = o.Major
	op.Minor = o.Minor
	op.BuildNumber = o.BuildNumber
	op.Workgroup = o.Workgroup
	return op
}

func (o *OpenQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenQueueOperation) {
	if o == nil {
		return
	}
	o.QueueFormat = op.QueueFormat
	o.Access = op.Access
	o.ShareMode = op.ShareMode
	o.ClientID = op.ClientID
	o.NonRoutingServer = op.NonRoutingServer
	o.Major = op.Major
	o.Minor = op.Minor
	o.BuildNumber = op.BuildNumber
	o.Workgroup = op.Workgroup
}
func (o *OpenQueueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenQueueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenQueueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenQueueResponse structure represents the R_OpenQueue operation response
type OpenQueueResponse struct {
	// pphContext: MUST be set by the server to a QUEUE_CONTEXT_HANDLE_SERIALIZE handle.
	Context *QueueSerialize `idl:"name:pphContext" json:"context"`
}

func (o *OpenQueueResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenQueueOperation) *xxx_OpenQueueOperation {
	if op == nil {
		op = &xxx_OpenQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *OpenQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenQueueOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *OpenQueueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenQueueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenQueueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseQueueOperation structure represents the R_CloseQueue operation
type xxx_CloseQueueOperation struct {
	Context *QueueSerialize `idl:"name:pphContext" json:"context"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseQueueOperation) OpNum() int { return 3 }

func (o *xxx_CloseQueueOperation) OpName() string { return "/RemoteRead/v1/R_CloseQueue" }

func (o *xxx_CloseQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseQueueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=QUEUE_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=QUEUE_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseQueueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseQueueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=QUEUE_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=QUEUE_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseQueueRequest structure represents the R_CloseQueue operation request
type CloseQueueRequest struct {
	// pphContext: MUST be set by the client to the QUEUE_CONTEXT_HANDLE_SERIALIZE handle
	// to be closed. The handle MUST have been returned by the server in the pphContext
	// parameter of a prior call to the R_OpenQueue method or the R_OpenQueueForMove method
	// and MUST NOT have been closed through a prior call to R_CloseQueue. This value MUST
	// NOT be NULL. If the server returns MQ_OK, it MUST set this value to NULL.
	Context *QueueSerialize `idl:"name:pphContext" json:"context"`
}

func (o *CloseQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseQueueOperation) *xxx_CloseQueueOperation {
	if op == nil {
		op = &xxx_CloseQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *CloseQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseQueueOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *CloseQueueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseQueueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseQueueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseQueueResponse structure represents the R_CloseQueue operation response
type CloseQueueResponse struct {
	// pphContext: MUST be set by the client to the QUEUE_CONTEXT_HANDLE_SERIALIZE handle
	// to be closed. The handle MUST have been returned by the server in the pphContext
	// parameter of a prior call to the R_OpenQueue method or the R_OpenQueueForMove method
	// and MUST NOT have been closed through a prior call to R_CloseQueue. This value MUST
	// NOT be NULL. If the server returns MQ_OK, it MUST set this value to NULL.
	Context *QueueSerialize `idl:"name:pphContext" json:"context"`
	// Return: The R_CloseQueue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseQueueResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseQueueOperation) *xxx_CloseQueueOperation {
	if op == nil {
		op = &xxx_CloseQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Return = o.Return
	return op
}

func (o *CloseQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseQueueOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Return = op.Return
}
func (o *CloseQueueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseQueueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseQueueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateCursorOperation structure represents the R_CreateCursor operation
type xxx_CreateCursorOperation struct {
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
	Cursor  uint32            `idl:"name:phCursor" json:"cursor"`
	Return  int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateCursorOperation) OpNum() int { return 4 }

func (o *xxx_CreateCursorOperation) OpName() string { return "/RemoteRead/v1/R_CreateCursor" }

func (o *xxx_CreateCursorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCursorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateCursorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueNoSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCursorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCursorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phCursor {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cursor); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCursorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phCursor {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cursor); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateCursorRequest structure represents the R_CreateCursor operation request
type CreateCursorRequest struct {
	// phContext: MUST be set by the client to the QUEUE_CONTEXT_HANDLE_NOSERIALIZE (section
	// 2.2.4.1) handle with which to associate the cursor. The handle MUST have been returned
	// by the server in the pphQueue output parameter of a prior call to the R_OpenQueue
	// (Opnum 2) (section 3.1.4.2) method and MUST NOT have been closed through a prior
	// call to the R_CloseQueue (Opnum 3) (section 3.1.4.3) method. This value MUST NOT
	// be NULL.
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
}

func (o *CreateCursorRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateCursorOperation) *xxx_CreateCursorOperation {
	if op == nil {
		op = &xxx_CreateCursorOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *CreateCursorRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateCursorOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *CreateCursorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateCursorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateCursorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateCursorResponse structure represents the R_CreateCursor operation response
type CreateCursorResponse struct {
	// phCursor: MUST be set by the server to a handle for the created cursor.
	Cursor uint32 `idl:"name:phCursor" json:"cursor"`
	// Return: The R_CreateCursor return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateCursorResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateCursorOperation) *xxx_CreateCursorOperation {
	if op == nil {
		op = &xxx_CreateCursorOperation{}
	}
	if o == nil {
		return op
	}
	op.Cursor = o.Cursor
	op.Return = o.Return
	return op
}

func (o *CreateCursorResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateCursorOperation) {
	if o == nil {
		return
	}
	o.Cursor = op.Cursor
	o.Return = op.Return
}
func (o *CreateCursorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateCursorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateCursorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseCursorOperation structure represents the R_CloseCursor operation
type xxx_CloseCursorOperation struct {
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
	Cursor  uint32            `idl:"name:hCursor" json:"cursor"`
	Return  int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseCursorOperation) OpNum() int { return 5 }

func (o *xxx_CloseCursorOperation) OpName() string { return "/RemoteRead/v1/R_CloseCursor" }

func (o *xxx_CloseCursorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseCursorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cursor); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseCursorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueNoSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cursor); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseCursorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseCursorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseCursorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseCursorRequest structure represents the R_CloseCursor operation request
type CloseCursorRequest struct {
	// phContext: MUST be set by the client to the QUEUE_CONTEXT_HANDLE_NOSERIALIZE (section
	// 2.2.4.1) handle with which the cursor was associated in a call to the R_CreateCursor
	// method. The handle MUST have been returned by the server in the pphQueue output parameter
	// of a prior call to the R_OpenQueue (Opnum 2) (section 3.1.4.2) method and MUST NOT
	// have been closed through a prior call to the R_CloseQueue (Opnum 3) (section 3.1.4.3)
	// method. This value MUST NOT be NULL.
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
	// hCursor: MUST be set by the client to the handle of the cursor to be closed. The
	// handle MUST have been obtained by a prior call to the R_CreateCursor method and MUST
	// NOT have been closed through a prior call to the R_CloseCursor method.
	Cursor uint32 `idl:"name:hCursor" json:"cursor"`
}

func (o *CloseCursorRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseCursorOperation) *xxx_CloseCursorOperation {
	if op == nil {
		op = &xxx_CloseCursorOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Cursor = o.Cursor
	return op
}

func (o *CloseCursorRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseCursorOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Cursor = op.Cursor
}
func (o *CloseCursorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseCursorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseCursorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseCursorResponse structure represents the R_CloseCursor operation response
type CloseCursorResponse struct {
	// Return: The R_CloseCursor return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseCursorResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseCursorOperation) *xxx_CloseCursorOperation {
	if op == nil {
		op = &xxx_CloseCursorOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *CloseCursorResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseCursorOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CloseCursorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseCursorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseCursorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PurgeQueueOperation structure represents the R_PurgeQueue operation
type xxx_PurgeQueueOperation struct {
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
	Return  int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_PurgeQueueOperation) OpNum() int { return 6 }

func (o *xxx_PurgeQueueOperation) OpName() string { return "/RemoteRead/v1/R_PurgeQueue" }

func (o *xxx_PurgeQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PurgeQueueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_PurgeQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueNoSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PurgeQueueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PurgeQueueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PurgeQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PurgeQueueRequest structure represents the R_PurgeQueue operation request
type PurgeQueueRequest struct {
	// phContext: MUST be set by the client to a QUEUE_CONTEXT_HANDLE_NOSERIALIZE (section
	// 2.2.4.1) handle of the queue to be purged. The handle MUST have been returned by
	// the server in the pphQueue output parameter of a prior call to the R_OpenQueue (Opnum
	// 2) (section 3.1.4.2) method with the dwAccess parameter set to RECEIVE_ACCESS and
	// MUST NOT have been closed through a prior call to the R_CloseQueue (Opnum 3) (section
	// 3.1.4.3) method. This value MUST NOT be NULL.
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
}

func (o *PurgeQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_PurgeQueueOperation) *xxx_PurgeQueueOperation {
	if op == nil {
		op = &xxx_PurgeQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *PurgeQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_PurgeQueueOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *PurgeQueueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PurgeQueueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PurgeQueueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PurgeQueueResponse structure represents the R_PurgeQueue operation response
type PurgeQueueResponse struct {
	// Return: The R_PurgeQueue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PurgeQueueResponse) xxx_ToOp(ctx context.Context, op *xxx_PurgeQueueOperation) *xxx_PurgeQueueOperation {
	if op == nil {
		op = &xxx_PurgeQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *PurgeQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_PurgeQueueOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PurgeQueueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PurgeQueueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PurgeQueueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StartReceiveOperation structure represents the R_StartReceive operation
type xxx_StartReceiveOperation struct {
	Context                *QueueNoSerialize `idl:"name:phContext" json:"context"`
	LookupID               uint64            `idl:"name:LookupId" json:"lookup_id"`
	Cursor                 uint32            `idl:"name:hCursor" json:"cursor"`
	Action                 uint32            `idl:"name:ulAction" json:"action"`
	Timeout                uint32            `idl:"name:ulTimeout" json:"timeout"`
	RequestID              uint32            `idl:"name:dwRequestId" json:"request_id"`
	MaxBodySize            uint32            `idl:"name:dwMaxBodySize" json:"max_body_size"`
	MaxCompoundMessageSize uint32            `idl:"name:dwMaxCompoundMessageSize" json:"max_compound_message_size"`
	ArriveTime             uint32            `idl:"name:pdwArriveTime" json:"arrive_time"`
	SequenceID             uint64            `idl:"name:pSequenceId" json:"sequence_id"`
	NumberOfSections       uint32            `idl:"name:pdwNumberOfSections" json:"number_of_sections"`
	PacketSections         []*SectionBuffer  `idl:"name:ppPacketSections;size_is:(, pdwNumberOfSections)" json:"packet_sections"`
	Return                 int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_StartReceiveOperation) OpNum() int { return 7 }

func (o *xxx_StartReceiveOperation) OpName() string { return "/RemoteRead/v1/R_StartReceive" }

func (o *xxx_StartReceiveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LookupId {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.LookupID); err != nil {
			return err
		}
	}
	// hCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cursor); err != nil {
			return err
		}
	}
	// ulAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Action); err != nil {
			return err
		}
	}
	// ulTimeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	// dwMaxBodySize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxBodySize); err != nil {
			return err
		}
	}
	// dwMaxCompoundMessageSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxCompoundMessageSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueNoSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LookupId {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.LookupID); err != nil {
			return err
		}
	}
	// hCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cursor); err != nil {
			return err
		}
	}
	// ulAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Action); err != nil {
			return err
		}
	}
	// ulTimeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	// dwMaxBodySize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxBodySize); err != nil {
			return err
		}
	}
	// dwMaxCompoundMessageSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxCompoundMessageSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.PacketSections != nil && o.NumberOfSections == 0 {
		o.NumberOfSections = uint32(len(o.PacketSections))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwArriveTime {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ArriveTime); err != nil {
			return err
		}
	}
	// pSequenceId {out} (1:{pointer=ref}*(1))(2:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.SequenceID); err != nil {
			return err
		}
	}
	// pdwNumberOfSections {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberOfSections); err != nil {
			return err
		}
	}
	// ppPacketSections {out} (1:{pointer=ref}*(2)*(1))(2:{alias=SectionBuffer}[dim:0,size_is=pdwNumberOfSections](struct))
	{
		if o.PacketSections != nil || o.NumberOfSections > 0 {
			_ptr_ppPacketSections := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfSections)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PacketSections {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.PacketSections[i1] != nil {
						if err := o.PacketSections[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&SectionBuffer{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.PacketSections); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&SectionBuffer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PacketSections, _ptr_ppPacketSections); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwArriveTime {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ArriveTime); err != nil {
			return err
		}
	}
	// pSequenceId {out} (1:{pointer=ref}*(1))(2:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.SequenceID); err != nil {
			return err
		}
	}
	// pdwNumberOfSections {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberOfSections); err != nil {
			return err
		}
	}
	// ppPacketSections {out} (1:{pointer=ref}*(2)*(1))(2:{alias=SectionBuffer}[dim:0,size_is=pdwNumberOfSections](struct))
	{
		_ptr_ppPacketSections := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PacketSections", sizeInfo[0])
			}
			o.PacketSections = make([]*SectionBuffer, sizeInfo[0])
			for i1 := range o.PacketSections {
				i1 := i1
				if o.PacketSections[i1] == nil {
					o.PacketSections[i1] = &SectionBuffer{}
				}
				if err := o.PacketSections[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppPacketSections := func(ptr interface{}) { o.PacketSections = *ptr.(*[]*SectionBuffer) }
		if err := w.ReadPointer(&o.PacketSections, _s_ppPacketSections, _ptr_ppPacketSections); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StartReceiveRequest structure represents the R_StartReceive operation request
type StartReceiveRequest struct {
	// phContext: MUST be set by the client to a QUEUE_CONTEXT_HANDLE_NOSERIALIZE (section
	// 2.2.4.1) handle of the queue from which to read a message. The handle MUST have been
	// returned by the server in the pphQueue output parameter of a prior call to the R_OpenQueue
	// (Opnum 2) (section 3.1.4.2) method and MUST NOT have been closed through a call prior
	// to the R_CloseQueue (Opnum 3) (section 3.1.4.3) method. This value MUST NOT be NULL.
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
	// LookupId: If nonzero, specifies the lookup identifier of the message to be acted
	// on.
	//
	// * ulTimeout set to 0x00000000.
	//
	// * hCursor set to 0x00000000.
	//
	// * ulAction set to one of the following:
	//
	// * MQ_LOOKUP_PEEK_PREV
	//
	// * MQ_LOOKUP_PEEK_CURRENT
	//
	// * MQ_LOOKUP_PEEK_NEXT
	//
	// * MQ_LOOKUP_RECEIVE_PREV
	//
	// * MQ_LOOKUP_RECEIVE_CURRENT
	//
	// * MQ_LOOKUP_RECEIVE_NEXT
	//
	// If the client sets the LookupId parameter to 0x0000000000000000, all of the preceding
	// values of the ulAction parameter are invalid.
	LookupID uint64 `idl:"name:LookupId" json:"lookup_id"`
	// hCursor: If nonzero, specifies a handle to a cursor that MUST have been obtained
	// from a prior call to the R_CreateCursor (Opnum 4) (section 3.1.4.4) method. The handle
	// MUST NOT have been closed through a prior call to the R_CloseCursor (Opnum 5) (section
	// 3.1.4.5) method.
	//
	// * LookupId set to 0x0000000000000000
	//
	// * ulAction set to one of the following:
	//
	// * MQ_ACTION_RECEIVE
	//
	// * MQ_ACTION_PEEK_CURRENT
	//
	// * MQ_ACTION_PEEK_NEXT
	Cursor uint32 `idl:"name:hCursor" json:"cursor"`
	// ulAction: Specifies the action to perform on the message. The following table lists
	// possible actions.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                TYPE /                |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_RECEIVE 0x00000000         | If the hCursor parameter is nonzero, read and remove the message for the current |
	//	|                                      | cursor location, and advance the cursor to the next position. If the hCursor     |
	//	|                                      | parameter is 0x00000000, read and remove the message from the front of the       |
	//	|                                      | queue. The valid values for other parameters are as follows: LookupId set to     |
	//	|                                      | 0x0000000000000000.                                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_CURRENT 0x80000000    | If the hCursor parameter is nonzero, read the message at the current cursor      |
	//	|                                      | location, but do not remove it from the queue. If the hCursor parameter is       |
	//	|                                      | 0x00000000, read the message at the front of the queue, but do not remove it     |
	//	|                                      | from the queue. The valid values for other parameters are as follows: LookupId   |
	//	|                                      | set to 0x0000000000000000.                                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_NEXT 0x80000001       | If the hCursor parameter is nonzero, advance the cursor to the next position     |
	//	|                                      | and read the message, but do not remove it from the queue. The valid values for  |
	//	|                                      | other parameters are as follows: LookupId set to 0x0000000000000000. hCursor set |
	//	|                                      | to a nonzero cursor handle obtained from the R_CreateCursor method.              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_CURRENT 0x40000010    | Read the message specified by the LookupId parameter, but do not remove it from  |
	//	|                                      | the queue. The valid values for other parameters are as follows: LookupId set to |
	//	|                                      | a nonzero value. hCursor set to 0x00000000. ulTimeout set to 0x00000000.         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_NEXT 0x40000011       | Read the message following the message specified by the LookupId parameter, but  |
	//	|                                      | do not remove it. The valid values for other parameters are as follows: LookupId |
	//	|                                      | set to a nonzero value. hCursor set to 0x00000000. ulTimeout set to 0x00000000.  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_PREV 0x40000012       | Read the message preceding the message specified by the LookupId parameter, but  |
	//	|                                      | do not remove it from the queue. The valid values for other parameters are as    |
	//	|                                      | follows: LookupId set to a nonzero value. hCursor set to 0x00000000. ulTimeout   |
	//	|                                      | set to 0x00000000.                                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_CURRENT 0x40000020 | Read the message specified by the LookupId parameter, and remove it from the     |
	//	|                                      | queue. The valid values for other parameters are as follows: LookupId set to a   |
	//	|                                      | nonzero value. hCursor set to 0x00000000. ulTimeout set to 0x00000000.           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_NEXT 0x40000021    | Read the message following the message specified by the LookupId parameter, and  |
	//	|                                      | remove it from the queue. The valid values for other parameters are as follows:  |
	//	|                                      | LookupId set to a nonzero value. hCursor set to 0x00000000. ulTimeout set to     |
	//	|                                      | 0x00000000.                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_PREV 0x40000022    | Read the message preceding the message specified by the LookupId parameter, and  |
	//	|                                      | remove it from the queue. The valid values for other parameters are as follows:  |
	//	|                                      | LookupId set to 0x0000000000000000. hCursor set to 0x00000000. ulTimeout set to  |
	//	|                                      | 0x00000000.                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// If the hCursor parameter is 0x00000000 and the LookupId parameter is 0x0000000000000000,
	// the valid values for the ulAction parameter are as follows:
	//
	// * MQ_ACTION_RECEIVE
	//
	// * MQ_ACTION_PEEK_CURRENT
	Action uint32 `idl:"name:ulAction" json:"action"`
	// ulTimeout: Specifies the time-out, in milliseconds, to wait for a message to become
	// available in the queue. The valid value for this parameter is 0x00000000 if the LookupId
	// parameter value is nonzero or if the action is not MQ_ACTION_RECEIVE, MQ_ACTION_PEEK_CURRENT,
	// or MQ_ACTION_PEEK_NEXT.
	Timeout uint32 `idl:"name:ulTimeout" json:"timeout"`
	// dwRequestId: MUST be set by the client to a unique correlation identifier for the
	// receive request. This value MUST be used in a subsequent call to the R_EndReceive
	// method or the R_CancelReceive method to correlate that call with the call to the
	// R_StartReceive method. The value MUST NOT be used in another R_StartReceive method
	// call on the same QUEUE_CONTEXT_HANDLE_NOSERIALIZE handle until a call to either the
	// R_EndReceive method or the R_CancelReceive method with the same dwRequestId parameter
	// value has been completed.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
	// dwMaxBodySize: MUST be set by the client to the maximum size, in bytes, of the message
	// body to be returned. The server SHOULD ignore this parameter when the message is
	// not a Binary Message (section 2.2.5.1.1).
	MaxBodySize uint32 `idl:"name:dwMaxBodySize" json:"max_body_size"`
	// dwMaxCompoundMessageSize: MUST be set by the client to the maximum size, in bytes,
	// of the CompoundMessageHeader (section 2.2.5.1.2.2). The server SHOULD ignore this
	// parameter when the message is not an SRMP Message (section 2.2.5.1.2).
	MaxCompoundMessageSize uint32 `idl:"name:dwMaxCompoundMessageSize" json:"max_compound_message_size"`
}

func (o *StartReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_StartReceiveOperation) *xxx_StartReceiveOperation {
	if op == nil {
		op = &xxx_StartReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.LookupID = o.LookupID
	op.Cursor = o.Cursor
	op.Action = o.Action
	op.Timeout = o.Timeout
	op.RequestID = o.RequestID
	op.MaxBodySize = o.MaxBodySize
	op.MaxCompoundMessageSize = o.MaxCompoundMessageSize
	return op
}

func (o *StartReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_StartReceiveOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.LookupID = op.LookupID
	o.Cursor = op.Cursor
	o.Action = op.Action
	o.Timeout = op.Timeout
	o.RequestID = op.RequestID
	o.MaxBodySize = op.MaxBodySize
	o.MaxCompoundMessageSize = op.MaxCompoundMessageSize
}
func (o *StartReceiveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StartReceiveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartReceiveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StartReceiveResponse structure represents the R_StartReceive operation response
type StartReceiveResponse struct {
	// pdwArriveTime: The server MUST set this value to the time that the message was added
	// to the queue ([MS-MQDMPR] section 3.1.7.3.1), expressed as the number of seconds
	// elapsed since midnight 00:00:00.0, January 1, 1970 UTC.
	ArriveTime uint32 `idl:"name:pdwArriveTime" json:"arrive_time"`
	// pSequenceId: The server MUST set this parameter to the least significant 7 bytes
	// of the Message.LookupIdentifier of the message that is received by this request.
	SequenceID uint64 `idl:"name:pSequenceId" json:"sequence_id"`
	// pdwNumberOfSections: The server MUST set this parameter to the number of entries
	// in the array pointed to by the ppPacketSections parameter.
	NumberOfSections uint32 `idl:"name:pdwNumberOfSections" json:"number_of_sections"`
	// ppPacketSections: The server MUST set this parameter to an array of pointers to SectionBuffer
	// structures. The server MUST fill this array in the following manner:
	//
	// * Create two local variables of type DWORD ( ../ms-dtyp/262627d8-3418-4627-9218-4ffe110850b2
	// ) called maxMessageSize and actualMessageSize. Assign the following values to these
	// variables:
	//
	// If the message is a Binary Message (section 2.2.5.1.1):
	//
	// * maxMessageSize := dwMaxBodySize
	//
	// * actualMessageSize := message packet body size
	//
	// If the message is an SRMP Message (section 2.2.5.1.2):
	//
	// * maxMessageSize := dwMaxCompoundMessageSize
	//
	// * actualMessageSize := size in bytes of *CompoundMessageHeader*
	//
	// * If the value of maxMessageSize is greater than or equal to actualMessageSize ,
	// the ppPacketSections parameter MUST contain a single entry as follows:
	//
	// * SectionType (section 2.2.7) ( f325738c-eca9-4450-b09c-cfe8ade87d81 ) MUST be set
	// to stFullPacket (0x00000000).
	//
	// * The *SectionSize* and *SectionSizeAlloc* elements MUST be set to the message packet
	// size.
	//
	// * The *pSectionBuffer* member MUST contain the entire message packet.
	//
	// * If the value of maxMessageSize is less than actualMessageSize , the array MUST
	// contain a first entry as follows:
	//
	// * *SectionType* MUST be set to one of the following:
	//
	// * stBinaryFirstSection if the message packet is a binary packet.
	//
	// * stSrmpFirstSection if the message packet is an SRMP packet.
	//
	// * The *pSectionBuffer* member MUST contain the message packet headers ( b23ea276-dfad-4083-bc48-d0f8a40255fd#gt_65255b08-c9d4-46d3-adca-3a296d43ae4f
	// ) concatenated with the first maxMessageSize bytes of the message body.
	//
	// * The *SectionSizeAlloc* member MUST be set to the message packet header size plus
	// actualMessageSize.
	//
	// * The *SectionSize* member MUST be set to the size of the *pSectionBuffer* member.
	//
	// * If the value of maxMessageSize is less than actualMessageSize and the message packet
	// trailers ( b23ea276-dfad-4083-bc48-d0f8a40255fd#gt_8ae0387f-5b71-4be9-a326-fd130e7bef81
	// ) are not empty, the array MUST contain a second entry as follows:
	//
	// * *SectionType* MUST be set to one of the following:
	//
	// * stBinarySecondSection if the message packet is a binary packet.
	//
	// * stSrmpSecondSection if the message packet is an SRMP packet.
	//
	// * The *pSectionBuffer* member MUST contain the message packet trailers.
	//
	// * The *SectionSize* member and the *SectionSizeAlloc* member MUST be equal and set
	// to the message packet trailers size.
	//
	// * For the first entry in this array, the *pSectionBuffer* member points to a Message
	// Packet Structure (section 2.2.5) ( f9e71595-339a-4cc4-8341-371e0a4cb232 ). Within
	// this structure, set *UserMessage.BaseHeader.TimeToReachQueue* to *UserHeader.SentTime*
	// + *UserMessage.BaseHeader.TimeToReachQueue*.
	PacketSections []*SectionBuffer `idl:"name:ppPacketSections;size_is:(, pdwNumberOfSections)" json:"packet_sections"`
	// Return: The R_StartReceive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StartReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_StartReceiveOperation) *xxx_StartReceiveOperation {
	if op == nil {
		op = &xxx_StartReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.ArriveTime = o.ArriveTime
	op.SequenceID = o.SequenceID
	op.NumberOfSections = o.NumberOfSections
	op.PacketSections = o.PacketSections
	op.Return = o.Return
	return op
}

func (o *StartReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_StartReceiveOperation) {
	if o == nil {
		return
	}
	o.ArriveTime = op.ArriveTime
	o.SequenceID = op.SequenceID
	o.NumberOfSections = op.NumberOfSections
	o.PacketSections = op.PacketSections
	o.Return = op.Return
}
func (o *StartReceiveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StartReceiveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartReceiveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CancelReceiveOperation structure represents the R_CancelReceive operation
type xxx_CancelReceiveOperation struct {
	Context   *QueueNoSerialize `idl:"name:phContext" json:"context"`
	RequestID uint32            `idl:"name:dwRequestId" json:"request_id"`
	Return    int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelReceiveOperation) OpNum() int { return 8 }

func (o *xxx_CancelReceiveOperation) OpName() string { return "/RemoteRead/v1/R_CancelReceive" }

func (o *xxx_CancelReceiveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelReceiveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueNoSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelReceiveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CancelReceiveRequest structure represents the R_CancelReceive operation request
type CancelReceiveRequest struct {
	// phContext: MUST be set by the client to the QUEUE_CONTEXT_HANDLE_NOSERIALIZE (section
	// 2.2.4.1) handle used in the corresponding call to the R_StartReceive method that
	// is to be canceled. The handle MUST have been returned by the server in the pphQueue
	// output parameter of a prior call to the R_OpenQueue (Opnum 2) (section 3.1.4.2) method
	// and MUST NOT have been previously closed through a call to the R_CloseQueue (Opnum
	// 3) (section 3.1.4.3) method. This value MUST NOT be NULL.
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
	// dwRequestId: MUST be set by the client to the same value as the dwRequestId parameter
	// in the corresponding call to the R_StartReceive method or the R_StartTransactionalReceive
	// method. This parameter acts as an identifier to correlate an R_CancelReceive method
	// call to an R_StartReceive or an R_StartTransactionalReceive method call.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
}

func (o *CancelReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_CancelReceiveOperation) *xxx_CancelReceiveOperation {
	if op == nil {
		op = &xxx_CancelReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.RequestID = o.RequestID
	return op
}

func (o *CancelReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelReceiveOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.RequestID = op.RequestID
}
func (o *CancelReceiveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CancelReceiveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelReceiveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelReceiveResponse structure represents the R_CancelReceive operation response
type CancelReceiveResponse struct {
	// Return: The R_CancelReceive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CancelReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_CancelReceiveOperation) *xxx_CancelReceiveOperation {
	if op == nil {
		op = &xxx_CancelReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *CancelReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelReceiveOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CancelReceiveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CancelReceiveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelReceiveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EndReceiveOperation structure represents the R_EndReceive operation
type xxx_EndReceiveOperation struct {
	Context   *QueueNoSerialize `idl:"name:phContext" json:"context"`
	Ack       uint32            `idl:"name:dwAck" json:"ack"`
	RequestID uint32            `idl:"name:dwRequestId" json:"request_id"`
	Return    int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_EndReceiveOperation) OpNum() int { return 9 }

func (o *xxx_EndReceiveOperation) OpName() string { return "/RemoteRead/v1/R_EndReceive" }

func (o *xxx_EndReceiveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Ack < uint32(1) || o.Ack > uint32(2) {
		return fmt.Errorf("Ack is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndReceiveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwAck {in} (1:{range=(1,2), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Ack); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueNoSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwAck {in} (1:{range=(1,2), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Ack); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndReceiveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EndReceiveRequest structure represents the R_EndReceive operation request
type EndReceiveRequest struct {
	// phContext: MUST be set by the client to the QUEUE_CONTEXT_HANDLE_NOSERIALIZE (section
	// 2.2.4.1) handle used in the corresponding call to the R_StartReceive method. The
	// handle MUST have been returned by the server in the pphQueue output parameter of
	// a prior call to the R_OpenQueue (Opnum 2) (section 3.1.4.2) method and MUST NOT have
	// been closed through a prior call to the R_CloseQueue (Opnum 3) (section 3.1.4.3)
	// method. This value MUST NOT be NULL.
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
	// dwAck: MUST be set to an Acknowledgment (ACK) or a Negative Acknowledgment (NACK)
	// for the message packet received from the server in an R_StartReceive method request.
	// The following table lists possible values.
	//
	//	+--------------------+----------------------------------------------------------------------------------+
	//	|                    |                                                                                  |
	//	|       VALUE        |                                     MEANING                                      |
	//	|                    |                                                                                  |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| RR_ACK 0x00000002  | The client acknowledges that the message packet was received successfully.       |
	//	|                    | The server MUST remove the message from the queue and make it unavailable for    |
	//	|                    | subsequent consumption.                                                          |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| RR_NACK 0x00000001 | The client acknowledges that the message packet was not received successfully.   |
	//	|                    | The server MUST keep the message in the queue and make it available for          |
	//	|                    | subsequent consumption.                                                          |
	//	+--------------------+----------------------------------------------------------------------------------+
	Ack uint32 `idl:"name:dwAck" json:"ack"`
	// dwRequestId: MUST be set by the client to the same value as the dwRequestId parameter
	// in the corresponding call to the R_StartReceive method. This parameter acts as an
	// identifier to correlate an R_EndReceive method call to an R_StartReceive method call.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
}

func (o *EndReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_EndReceiveOperation) *xxx_EndReceiveOperation {
	if op == nil {
		op = &xxx_EndReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Ack = o.Ack
	op.RequestID = o.RequestID
	return op
}

func (o *EndReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_EndReceiveOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Ack = op.Ack
	o.RequestID = op.RequestID
}
func (o *EndReceiveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EndReceiveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EndReceiveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EndReceiveResponse structure represents the R_EndReceive operation response
type EndReceiveResponse struct {
	// Return: The R_EndReceive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EndReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_EndReceiveOperation) *xxx_EndReceiveOperation {
	if op == nil {
		op = &xxx_EndReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *EndReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_EndReceiveOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EndReceiveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EndReceiveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EndReceiveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MoveMessageOperation structure represents the R_MoveMessage operation
type xxx_MoveMessageOperation struct {
	ContextFrom   *QueueNoSerialize    `idl:"name:phContextFrom" json:"context_from"`
	ContextTo     uint64               `idl:"name:ullContextTo" json:"context_to"`
	LookupID      uint64               `idl:"name:LookupId" json:"lookup_id"`
	TransactionID *mqmq.TransactionUOW `idl:"name:pTransactionId" json:"transaction_id"`
	Return        int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_MoveMessageOperation) OpNum() int { return 10 }

func (o *xxx_MoveMessageOperation) OpName() string { return "/RemoteRead/v1/R_MoveMessage" }

func (o *xxx_MoveMessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveMessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContextFrom {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.ContextFrom != nil {
			if err := o.ContextFrom.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ullContextTo {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.ContextTo); err != nil {
			return err
		}
	}
	// LookupId {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.LookupID); err != nil {
			return err
		}
	}
	// pTransactionId {in} (1:{pointer=ref}*(1))(2:{alias=XACTUOW}(struct))
	{
		if o.TransactionID != nil {
			if err := o.TransactionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqmq.TransactionUOW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_MoveMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContextFrom {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.ContextFrom == nil {
			o.ContextFrom = &QueueNoSerialize{}
		}
		if err := o.ContextFrom.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ullContextTo {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.ContextTo); err != nil {
			return err
		}
	}
	// LookupId {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.LookupID); err != nil {
			return err
		}
	}
	// pTransactionId {in} (1:{pointer=ref}*(1))(2:{alias=XACTUOW}(struct))
	{
		if o.TransactionID == nil {
			o.TransactionID = &mqmq.TransactionUOW{}
		}
		if err := o.TransactionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveMessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveMessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MoveMessageRequest structure represents the R_MoveMessage operation request
type MoveMessageRequest struct {
	// phContextFrom: MUST be set by the client to a QUEUE_CONTEXT_HANDLE_NOSERIALIZE (section
	// 2.2.4.1) handle representing the source queue. The handle MUST have been returned
	// by the server in the pphQueue output parameter of a prior call to the R_OpenQueue
	// (Opnum 2) (section 3.1.4.2) method with the dwAccess parameter set to RECEIVE_ACCESS
	// and MUST NOT have been closed through a prior call to the R_CloseQueue (Opnum 3)
	// (section 3.1.4.3) method. This value MUST NOT be NULL.
	ContextFrom *QueueNoSerialize `idl:"name:phContextFrom" json:"context_from"`
	// ullContextTo: MUST be set by the client to a QUEUE_CONTEXT_HANDLE_NOSERIALIZE handle
	// representing the destination queue. The handle MUST have been returned by the server
	// in the pMoveContext output parameter of a prior call to the R_OpenQueueForMove (Opnum
	// 11) (section 3.1.4.11) method and MUST NOT have been closed through a prior call
	// to the R_CloseQueue method. This value MUST NOT be NULL.
	ContextTo uint64 `idl:"name:ullContextTo" json:"context_to"`
	// LookupId: MUST be set by the client to the lookup identifier of the message to be
	// moved.
	LookupID uint64 `idl:"name:LookupId" json:"lookup_id"`
	// pTransactionId: MUST be set by the client as a pointer to a transaction identifier
	// or to a zero value XACTUOW ([MS-MQMQ] section 2.2.18.1.8) structure. If the destination
	// queue is not a transactional queue, this value MUST be a pointer to a zero value
	// XACTUOW structure. If the value of the field is not zero, the transaction identifier
	// MUST have been registered with the server through a prior call to the R_QMEnlistRemoteTransaction
	// (Opnum 12) (section 3.1.4.12) method and MUST NOT be NULL.
	TransactionID *mqmq.TransactionUOW `idl:"name:pTransactionId" json:"transaction_id"`
}

func (o *MoveMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_MoveMessageOperation) *xxx_MoveMessageOperation {
	if op == nil {
		op = &xxx_MoveMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.ContextFrom = o.ContextFrom
	op.ContextTo = o.ContextTo
	op.LookupID = o.LookupID
	op.TransactionID = o.TransactionID
	return op
}

func (o *MoveMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_MoveMessageOperation) {
	if o == nil {
		return
	}
	o.ContextFrom = op.ContextFrom
	o.ContextTo = op.ContextTo
	o.LookupID = op.LookupID
	o.TransactionID = op.TransactionID
}
func (o *MoveMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MoveMessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveMessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MoveMessageResponse structure represents the R_MoveMessage operation response
type MoveMessageResponse struct {
	// Return: The R_MoveMessage return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MoveMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_MoveMessageOperation) *xxx_MoveMessageOperation {
	if op == nil {
		op = &xxx_MoveMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *MoveMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_MoveMessageOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MoveMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MoveMessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveMessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenQueueForMoveOperation structure represents the R_OpenQueueForMove operation
type xxx_OpenQueueForMoveOperation struct {
	QueueFormat      *mqmq.QueueFormat `idl:"name:pQueueFormat" json:"queue_format"`
	Access           uint32            `idl:"name:dwAccess" json:"access"`
	ShareMode        uint32            `idl:"name:dwShareMode" json:"share_mode"`
	ClientID         *dtyp.GUID        `idl:"name:pClientId" json:"client_id"`
	NonRoutingServer int32             `idl:"name:fNonRoutingServer" json:"non_routing_server"`
	Major            uint8             `idl:"name:Major" json:"major"`
	Minor            uint8             `idl:"name:Minor" json:"minor"`
	BuildNumber      uint16            `idl:"name:BuildNumber" json:"build_number"`
	Workgroup        int32             `idl:"name:fWorkgroup" json:"workgroup"`
	MoveContext      uint64            `idl:"name:pMoveContext" json:"move_context"`
	Context          *QueueSerialize   `idl:"name:pphContext" json:"context"`
}

func (o *xxx_OpenQueueForMoveOperation) OpNum() int { return 11 }

func (o *xxx_OpenQueueForMoveOperation) OpName() string { return "/RemoteRead/v1/R_OpenQueueForMove" }

func (o *xxx_OpenQueueForMoveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueForMoveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pQueueFormat {in} (1:{pointer=ref}*(1))(2:{alias=QUEUE_FORMAT}(struct))
	{
		if o.QueueFormat != nil {
			if err := o.QueueFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqmq.QueueFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwAccess {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Access); err != nil {
			return err
		}
	}
	// dwShareMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ShareMode); err != nil {
			return err
		}
	}
	// pClientId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClientID != nil {
			if err := o.ClientID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// fNonRoutingServer {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.NonRoutingServer); err != nil {
			return err
		}
	}
	// Major {in} (1:(uchar))
	{
		if err := w.WriteData(o.Major); err != nil {
			return err
		}
	}
	// Minor {in} (1:(uchar))
	{
		if err := w.WriteData(o.Minor); err != nil {
			return err
		}
	}
	// BuildNumber {in} (1:{alias=USHORT}(uint16))
	{
		if err := w.WriteData(o.BuildNumber); err != nil {
			return err
		}
	}
	// fWorkgroup {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Workgroup); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueForMoveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pQueueFormat {in} (1:{pointer=ref}*(1))(2:{alias=QUEUE_FORMAT}(struct))
	{
		if o.QueueFormat == nil {
			o.QueueFormat = &mqmq.QueueFormat{}
		}
		if err := o.QueueFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwAccess {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Access); err != nil {
			return err
		}
	}
	// dwShareMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ShareMode); err != nil {
			return err
		}
	}
	// pClientId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClientID == nil {
			o.ClientID = &dtyp.GUID{}
		}
		if err := o.ClientID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// fNonRoutingServer {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.NonRoutingServer); err != nil {
			return err
		}
	}
	// Major {in} (1:(uchar))
	{
		if err := w.ReadData(&o.Major); err != nil {
			return err
		}
	}
	// Minor {in} (1:(uchar))
	{
		if err := w.ReadData(&o.Minor); err != nil {
			return err
		}
	}
	// BuildNumber {in} (1:{alias=USHORT}(uint16))
	{
		if err := w.ReadData(&o.BuildNumber); err != nil {
			return err
		}
	}
	// fWorkgroup {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Workgroup); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueForMoveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueForMoveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pMoveContext {out} (1:{pointer=ref}*(1))(2:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.MoveContext); err != nil {
			return err
		}
	}
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=QUEUE_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OpenQueueForMoveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pMoveContext {out} (1:{pointer=ref}*(1))(2:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.MoveContext); err != nil {
			return err
		}
	}
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=QUEUE_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OpenQueueForMoveRequest structure represents the R_OpenQueueForMove operation request
type OpenQueueForMoveRequest struct {
	// pQueueFormat: MUST be a pointer to a QUEUE_FORMAT ([MS-MQMQ] section 2.2.7) structure
	// that identifies the queue to open. This value MUST NOT be NULL. The value of the
	// m_qft member MUST be one of QUEUE_FORMAT_TYPE_PUBLIC, QUEUE_FORMAT_TYPE_PRIVATE,
	// QUEUE_FORMAT_TYPE_DIRECT, QUEUE_FORMAT_TYPE_MACHINE, or QUEUE_FORMAT_TYPE_SUBQUEUE.
	QueueFormat *mqmq.QueueFormat `idl:"name:pQueueFormat" json:"queue_format"`
	// dwAccess: Specifies the required type of access to the queue. MUST be set by the
	// client to MQ_MOVE_ACCESS (0x00000004).
	Access uint32 `idl:"name:dwAccess" json:"access"`
	// dwShareMode: Specifies whether the client needs exclusive access to the queue. MUST
	// be set by the client to MQ_DENY_NONE (0x00000000), which permits multiple QUEUE_CONTEXT_HANDLE_SERIALIZE
	// handles to the queue to be opened concurrently.
	ShareMode uint32 `idl:"name:dwShareMode" json:"share_mode"`
	// pClientId: MUST be set by the client to a pointer to a valid GUID that uniquely identifies
	// the client. When the queue manager acts as the client, the queue manager sets this
	// value to the Identifier ADM attribute of the local QueueManager ADM element instance.
	// The server SHOULD ignore this parameter. This value MUST NOT be NULL.
	ClientID *dtyp.GUID `idl:"name:pClientId" json:"client_id"`
	// fNonRoutingServer: If the client is configured to operate in the role of an MSMQ
	// routing server, this parameter MUST be set to FALSE (0x00000000); otherwise, it MUST
	// be set to TRUE (0x00000001).<22> If the value of the fNonRoutingServer parameter
	// is FALSE (0x00000000), the server MUST ignore the pClientId parameter.
	//
	//	+-------+------------+
	//	|       |            |
	//	| NAME  |   VALUE    |
	//	|       |            |
	//	+-------+------------+
	//	+-------+------------+
	//	| FALSE | 0x00000000 |
	//	+-------+------------+
	//	| TRUE  | 0x00000001 |
	//	+-------+------------+
	NonRoutingServer int32 `idl:"name:fNonRoutingServer" json:"non_routing_server"`
	// Major: MUST be set by the client to an implementation-specific Major Version number
	// of the client. SHOULD<23> be ignored by the server.
	Major uint8 `idl:"name:Major" json:"major"`
	// Minor: MUST be set by the client to an implementation-specific Minor Version number
	// of the client. SHOULD<24> be ignored by the server.
	Minor uint8 `idl:"name:Minor" json:"minor"`
	// BuildNumber: MUST be set by the client to an implementation-specific Build Number
	// of the client. SHOULD<25> be ignored by the server.
	BuildNumber uint16 `idl:"name:BuildNumber" json:"build_number"`
	// fWorkgroup: MUST be set to TRUE (0x00000001) by the client if the client machine
	// is not a member of a Windows domain; otherwise, it MUST be set to FALSE (0x00000000).
	// The RPC authentication level required by the server MAY be based on this value in
	// subsequent calls on the interface.<26>
	//
	//	+-------+------------+
	//	|       |            |
	//	| NAME  |   VALUE    |
	//	|       |            |
	//	+-------+------------+
	//	+-------+------------+
	//	| FALSE | 0x00000000 |
	//	+-------+------------+
	//	| TRUE  | 0x00000001 |
	//	+-------+------------+
	Workgroup int32 `idl:"name:fWorkgroup" json:"workgroup"`
}

func (o *OpenQueueForMoveRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenQueueForMoveOperation) *xxx_OpenQueueForMoveOperation {
	if op == nil {
		op = &xxx_OpenQueueForMoveOperation{}
	}
	if o == nil {
		return op
	}
	op.QueueFormat = o.QueueFormat
	op.Access = o.Access
	op.ShareMode = o.ShareMode
	op.ClientID = o.ClientID
	op.NonRoutingServer = o.NonRoutingServer
	op.Major = o.Major
	op.Minor = o.Minor
	op.BuildNumber = o.BuildNumber
	op.Workgroup = o.Workgroup
	return op
}

func (o *OpenQueueForMoveRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenQueueForMoveOperation) {
	if o == nil {
		return
	}
	o.QueueFormat = op.QueueFormat
	o.Access = op.Access
	o.ShareMode = op.ShareMode
	o.ClientID = op.ClientID
	o.NonRoutingServer = op.NonRoutingServer
	o.Major = op.Major
	o.Minor = op.Minor
	o.BuildNumber = op.BuildNumber
	o.Workgroup = op.Workgroup
}
func (o *OpenQueueForMoveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenQueueForMoveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenQueueForMoveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenQueueForMoveResponse structure represents the R_OpenQueueForMove operation response
type OpenQueueForMoveResponse struct {
	// pMoveContext: The server MUST set this parameter to a pointer to a QUEUE_CONTEXT_HANDLE_SERIALIZE
	// handle and MUST set the value of this parameter to the same value as the contents
	// of the pphContext parameter. The server MUST set this value to a context that can
	// be used as the dwContextTo parameter in a subsequent call to the R_MoveMessage method.
	// Logically, it represents a reference to the QUEUE_CONTEXT_HANDLE_SERIALIZE handle
	// returned in the pphContext parameter.
	MoveContext uint64 `idl:"name:pMoveContext" json:"move_context"`
	// pphContext: MUST be set by the server to a QUEUE_CONTEXT_HANDLE_SERIALIZE handle.
	// A QUEUE_CONTEXT_HANDLE_SERIALIZE handle opened through a call to this method can
	// be closed through a subsequent call to the R_CloseQueue (Opnum 3) (section 3.1.4.3)
	// method.
	Context *QueueSerialize `idl:"name:pphContext" json:"context"`
}

func (o *OpenQueueForMoveResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenQueueForMoveOperation) *xxx_OpenQueueForMoveOperation {
	if op == nil {
		op = &xxx_OpenQueueForMoveOperation{}
	}
	if o == nil {
		return op
	}
	op.MoveContext = o.MoveContext
	op.Context = o.Context
	return op
}

func (o *OpenQueueForMoveResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenQueueForMoveOperation) {
	if o == nil {
		return
	}
	o.MoveContext = op.MoveContext
	o.Context = op.Context
}
func (o *OpenQueueForMoveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenQueueForMoveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenQueueForMoveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnlistRemoteTransactionOperation structure represents the R_QMEnlistRemoteTransaction operation
type xxx_EnlistRemoteTransactionOperation struct {
	TransactionID          *mqmq.TransactionUOW `idl:"name:pTransactionId" json:"transaction_id"`
	PropagationTokenLength uint32               `idl:"name:cbPropagationToken" json:"propagation_token_length"`
	PropagationToken       []byte               `idl:"name:pbPropagationToken;size_is:(cbPropagationToken)" json:"propagation_token"`
	QueueFormat            *mqmq.QueueFormat    `idl:"name:pQueueFormat" json:"queue_format"`
	Return                 int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_EnlistRemoteTransactionOperation) OpNum() int { return 12 }

func (o *xxx_EnlistRemoteTransactionOperation) OpName() string {
	return "/RemoteRead/v1/R_QMEnlistRemoteTransaction"
}

func (o *xxx_EnlistRemoteTransactionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.PropagationToken != nil && o.PropagationTokenLength == 0 {
		o.PropagationTokenLength = uint32(len(o.PropagationToken))
	}
	if o.PropagationTokenLength > uint32(131072) {
		return fmt.Errorf("PropagationTokenLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnlistRemoteTransactionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pTransactionId {in} (1:{pointer=ref}*(1))(2:{alias=XACTUOW}(struct))
	{
		if o.TransactionID != nil {
			if err := o.TransactionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqmq.TransactionUOW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// cbPropagationToken {in} (1:{range=(0,131072), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PropagationTokenLength); err != nil {
			return err
		}
	}
	// pbPropagationToken {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbPropagationToken](uchar))
	{
		dimSize1 := uint64(o.PropagationTokenLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.PropagationToken {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.PropagationToken[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.PropagationToken); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pQueueFormat {in} (1:{pointer=ref}*(1))(2:{alias=QUEUE_FORMAT}(struct))
	{
		if o.QueueFormat != nil {
			if err := o.QueueFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqmq.QueueFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnlistRemoteTransactionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pTransactionId {in} (1:{pointer=ref}*(1))(2:{alias=XACTUOW}(struct))
	{
		if o.TransactionID == nil {
			o.TransactionID = &mqmq.TransactionUOW{}
		}
		if err := o.TransactionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// cbPropagationToken {in} (1:{range=(0,131072), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PropagationTokenLength); err != nil {
			return err
		}
	}
	// pbPropagationToken {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbPropagationToken](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.PropagationToken", sizeInfo[0])
		}
		o.PropagationToken = make([]byte, sizeInfo[0])
		for i1 := range o.PropagationToken {
			i1 := i1
			if err := w.ReadData(&o.PropagationToken[i1]); err != nil {
				return err
			}
		}
	}
	// pQueueFormat {in} (1:{pointer=ref}*(1))(2:{alias=QUEUE_FORMAT}(struct))
	{
		if o.QueueFormat == nil {
			o.QueueFormat = &mqmq.QueueFormat{}
		}
		if err := o.QueueFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnlistRemoteTransactionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnlistRemoteTransactionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnlistRemoteTransactionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnlistRemoteTransactionRequest structure represents the R_QMEnlistRemoteTransaction operation request
type EnlistRemoteTransactionRequest struct {
	// pTransactionId: MUST be a pointer to a transaction identifier obtained as specified
	// in [MS-DTCO] section 3.3.4.1. This value MUST NOT be NULL.
	TransactionID *mqmq.TransactionUOW `idl:"name:pTransactionId" json:"transaction_id"`
	// cbPropagationToken: MUST be the size, in bytes, of the pbPropagationToken parameter.
	PropagationTokenLength uint32 `idl:"name:cbPropagationToken" json:"propagation_token_length"`
	// pbPropagationToken: MUST be a transaction propagation token, as specified in [MS-DTCO]
	// section 2.2.5.4, that represents the transaction identified by the pTransactionId
	// parameter. This parameter MUST NOT be NULL.
	PropagationToken []byte `idl:"name:pbPropagationToken;size_is:(cbPropagationToken)" json:"propagation_token"`
	// pQueueFormat: MUST be a pointer to a QUEUE_FORMAT ([MS-MQMQ] section 2.2.7) structure
	// that identifies the queue to be passed to the R_StartTransactionalReceive method.
	// SHOULD<28> be ignored by the server.
	QueueFormat *mqmq.QueueFormat `idl:"name:pQueueFormat" json:"queue_format"`
}

func (o *EnlistRemoteTransactionRequest) xxx_ToOp(ctx context.Context, op *xxx_EnlistRemoteTransactionOperation) *xxx_EnlistRemoteTransactionOperation {
	if op == nil {
		op = &xxx_EnlistRemoteTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.TransactionID = o.TransactionID
	op.PropagationTokenLength = o.PropagationTokenLength
	op.PropagationToken = o.PropagationToken
	op.QueueFormat = o.QueueFormat
	return op
}

func (o *EnlistRemoteTransactionRequest) xxx_FromOp(ctx context.Context, op *xxx_EnlistRemoteTransactionOperation) {
	if o == nil {
		return
	}
	o.TransactionID = op.TransactionID
	o.PropagationTokenLength = op.PropagationTokenLength
	o.PropagationToken = op.PropagationToken
	o.QueueFormat = op.QueueFormat
}
func (o *EnlistRemoteTransactionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnlistRemoteTransactionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnlistRemoteTransactionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnlistRemoteTransactionResponse structure represents the R_QMEnlistRemoteTransaction operation response
type EnlistRemoteTransactionResponse struct {
	// Return: The R_QMEnlistRemoteTransaction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnlistRemoteTransactionResponse) xxx_ToOp(ctx context.Context, op *xxx_EnlistRemoteTransactionOperation) *xxx_EnlistRemoteTransactionOperation {
	if op == nil {
		op = &xxx_EnlistRemoteTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *EnlistRemoteTransactionResponse) xxx_FromOp(ctx context.Context, op *xxx_EnlistRemoteTransactionOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EnlistRemoteTransactionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnlistRemoteTransactionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnlistRemoteTransactionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StartTransactionalReceiveOperation structure represents the R_StartTransactionalReceive operation
type xxx_StartTransactionalReceiveOperation struct {
	Context                *QueueNoSerialize    `idl:"name:phContext" json:"context"`
	LookupID               uint64               `idl:"name:LookupId" json:"lookup_id"`
	Cursor                 uint32               `idl:"name:hCursor" json:"cursor"`
	Action                 uint32               `idl:"name:ulAction" json:"action"`
	Timeout                uint32               `idl:"name:ulTimeout" json:"timeout"`
	RequestID              uint32               `idl:"name:dwRequestId" json:"request_id"`
	MaxBodySize            uint32               `idl:"name:dwMaxBodySize" json:"max_body_size"`
	MaxCompoundMessageSize uint32               `idl:"name:dwMaxCompoundMessageSize" json:"max_compound_message_size"`
	TransactionID          *mqmq.TransactionUOW `idl:"name:pTransactionId" json:"transaction_id"`
	ArriveTime             uint32               `idl:"name:pdwArriveTime" json:"arrive_time"`
	SequenceID             uint64               `idl:"name:pSequenceId" json:"sequence_id"`
	NumberOfSections       uint32               `idl:"name:pdwNumberOfSections" json:"number_of_sections"`
	PacketSections         []*SectionBuffer     `idl:"name:ppPacketSections;size_is:(, pdwNumberOfSections)" json:"packet_sections"`
	Return                 int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_StartTransactionalReceiveOperation) OpNum() int { return 13 }

func (o *xxx_StartTransactionalReceiveOperation) OpName() string {
	return "/RemoteRead/v1/R_StartTransactionalReceive"
}

func (o *xxx_StartTransactionalReceiveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartTransactionalReceiveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LookupId {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.LookupID); err != nil {
			return err
		}
	}
	// hCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cursor); err != nil {
			return err
		}
	}
	// ulAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Action); err != nil {
			return err
		}
	}
	// ulTimeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	// dwMaxBodySize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxBodySize); err != nil {
			return err
		}
	}
	// dwMaxCompoundMessageSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxCompoundMessageSize); err != nil {
			return err
		}
	}
	// pTransactionId {in} (1:{pointer=ref}*(1))(2:{alias=XACTUOW}(struct))
	{
		if o.TransactionID != nil {
			if err := o.TransactionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqmq.TransactionUOW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_StartTransactionalReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueNoSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LookupId {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.LookupID); err != nil {
			return err
		}
	}
	// hCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cursor); err != nil {
			return err
		}
	}
	// ulAction {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Action); err != nil {
			return err
		}
	}
	// ulTimeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	// dwMaxBodySize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxBodySize); err != nil {
			return err
		}
	}
	// dwMaxCompoundMessageSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxCompoundMessageSize); err != nil {
			return err
		}
	}
	// pTransactionId {in} (1:{pointer=ref}*(1))(2:{alias=XACTUOW}(struct))
	{
		if o.TransactionID == nil {
			o.TransactionID = &mqmq.TransactionUOW{}
		}
		if err := o.TransactionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartTransactionalReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.PacketSections != nil && o.NumberOfSections == 0 {
		o.NumberOfSections = uint32(len(o.PacketSections))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartTransactionalReceiveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwArriveTime {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ArriveTime); err != nil {
			return err
		}
	}
	// pSequenceId {out} (1:{pointer=ref}*(1))(2:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.SequenceID); err != nil {
			return err
		}
	}
	// pdwNumberOfSections {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberOfSections); err != nil {
			return err
		}
	}
	// ppPacketSections {out} (1:{pointer=ref}*(2)*(1))(2:{alias=SectionBuffer}[dim:0,size_is=pdwNumberOfSections](struct))
	{
		if o.PacketSections != nil || o.NumberOfSections > 0 {
			_ptr_ppPacketSections := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfSections)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PacketSections {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.PacketSections[i1] != nil {
						if err := o.PacketSections[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&SectionBuffer{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.PacketSections); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&SectionBuffer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PacketSections, _ptr_ppPacketSections); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartTransactionalReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwArriveTime {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ArriveTime); err != nil {
			return err
		}
	}
	// pSequenceId {out} (1:{pointer=ref}*(1))(2:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.SequenceID); err != nil {
			return err
		}
	}
	// pdwNumberOfSections {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberOfSections); err != nil {
			return err
		}
	}
	// ppPacketSections {out} (1:{pointer=ref}*(2)*(1))(2:{alias=SectionBuffer}[dim:0,size_is=pdwNumberOfSections](struct))
	{
		_ptr_ppPacketSections := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PacketSections", sizeInfo[0])
			}
			o.PacketSections = make([]*SectionBuffer, sizeInfo[0])
			for i1 := range o.PacketSections {
				i1 := i1
				if o.PacketSections[i1] == nil {
					o.PacketSections[i1] = &SectionBuffer{}
				}
				if err := o.PacketSections[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppPacketSections := func(ptr interface{}) { o.PacketSections = *ptr.(*[]*SectionBuffer) }
		if err := w.ReadPointer(&o.PacketSections, _s_ppPacketSections, _ptr_ppPacketSections); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StartTransactionalReceiveRequest structure represents the R_StartTransactionalReceive operation request
type StartTransactionalReceiveRequest struct {
	// phContext: MUST be set by the client to a QUEUE_CONTEXT_HANDLE_NOSERIALIZE (section
	// 2.2.4.1) handle of the queue from which to read a message. The handle MUST have been
	// returned by the server in the pphQueue output parameter of a prior call to the R_OpenQueue
	// (Opnum 2) (section 3.1.4.2) method with the dwAccess parameter set to RECEIVE_ACCESS
	// and MUST NOT have been closed through a prior call to the R_CloseQueue (Opnum 3)
	// (section 3.1.4.3) method. NULL is invalid for this parameter.
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
	// LookupId: If nonzero, specifies the lookup identifier of the message to be acted
	// on.
	//
	// * ulTimeout set to 0x00000000
	//
	// * hCursor set to 0x00000000
	//
	// * ulAction set to one of the following:
	//
	// * MQ_LOOKUP_PEEK_PREV ( pTransactionId set to NULL)
	//
	// * MQ_LOOKUP_PEEK_CURRENT ( pTransactionId set to NULL)
	//
	// * MQ_LOOKUP_PEEK_NEXT ( pTransactionId set to NULL)
	//
	// * MQ_LOOKUP_RECEIVE_PREV
	//
	// * MQ_LOOKUP_RECEIVE_CURRENT
	//
	// * MQ_LOOKUP_RECEIVE_NEXT
	//
	// If the client sets the LookupId parameter to 0x0000000000000000, all of the preceding
	// values of the ulAction parameter are invalid.
	LookupID uint64 `idl:"name:LookupId" json:"lookup_id"`
	// hCursor: If nonzero, specifies a handle to a cursor that MUST have been obtained
	// from a prior call to the R_CreateCursor (Opnum 4) (section 3.1.4.4) method. The handle
	// MUST NOT have been closed through a prior call to the R_CloseCursor (Opnum 5) (section
	// 3.1.4.5) method.
	//
	// * LookupId set to 0x0000000000000000.
	//
	// * ulAction set to one of the following:
	//
	// * MQ_ACTION_RECEIVE
	//
	// * MQ_ACTION_PEEK_CURRENT ( pTransactionId set to NULL)
	//
	// * MQ_ACTION_PEEK_NEXT ( pTransactionId set to NULL)
	Cursor uint32 `idl:"name:hCursor" json:"cursor"`
	// ulAction: Specifies the action to perform on the message. The following table lists
	// possible actions.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                TYPE /                |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_RECEIVE 0x00000000         | If the hCursor parameter is nonzero, read and remove the message at the          |
	//	|                                      | current cursor location from the queue, and advance the cursor. If the hCursor   |
	//	|                                      | parameter is 0x00000000, read and remove the message from the front of the       |
	//	|                                      | queue. The valid values for other parameters are as follows: LookupId set to     |
	//	|                                      | 0x0000000000000000.                                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_CURRENT 0x80000000    | If the hCursor parameter is nonzero, read the message at the current cursor      |
	//	|                                      | location, but do not remove it from the queue. If the hCursor parameter is       |
	//	|                                      | 0x00000000, read the message at the front of the queue, but do not remove it     |
	//	|                                      | from the queue. The valid values for other parameters are as follows: LookupId   |
	//	|                                      | set to 0x0000000000000000. pTransactionId set to NULL.                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_NEXT 0x80000001       | If the hCursor parameter is nonzero, advance the cursor to the next position,    |
	//	|                                      | and read the message, but do not remove it from the queue. The valid values      |
	//	|                                      | for other parameters are as follows: LookupId set to 0x0000000000000000.         |
	//	|                                      | hCursor set to a nonzero cursor handle obtained from the R_CreateCursor method.  |
	//	|                                      | pTransactionId set to NULL.                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_CURRENT 0x40000010    | Read the message specified by the LookupId parameter, but do not remove it from  |
	//	|                                      | the queue. The valid values for other parameters are as follows: LookupId set    |
	//	|                                      | to a nonzero value. hCursor set to 0x00000000. ulTimeout set to 0x00000000.      |
	//	|                                      | pTransactionId set to NULL.                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_NEXT 0x40000011       | Read the message following the message specified by the LookupId parameter, but  |
	//	|                                      | do not remove it. The valid values for other parameters are as follows: LookupId |
	//	|                                      | set to a nonzero value. hCursor set to 0x00000000. ulTimeout set to 0x00000000.  |
	//	|                                      | pTransactionId set to NULL.                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_PREV 0x40000012       | Read the message preceding the message specified by the LookupId parameter, but  |
	//	|                                      | do not remove it from the queue. The valid values for other parameters are as    |
	//	|                                      | follows: LookupId set to a nonzero value. hCursor set to 0x00000000. ulTimeout   |
	//	|                                      | set to 0x00000000. pTransactionId set to NULL.                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_CURRENT 0x40000020 | Read the message specified by the LookupId parameter, and remove it from the     |
	//	|                                      | queue. The valid values for other parameters are as follows: LookupId set to a   |
	//	|                                      | nonzero value. hCursor set to 0x00000000. ulTimeout set to 0x00000000.           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_NEXT 0x40000021    | Read the message following the message specified by the LookupId parameter, and  |
	//	|                                      | remove it from the queue. The valid values for other parameters are as follows:  |
	//	|                                      | LookupId set to a nonzero value. hCursor set to 0x00000000. ulTimeout set to     |
	//	|                                      | 0x00000000.                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_PREV 0x40000022    | Read the message preceding the message specified by the LookupId parameter, and  |
	//	|                                      | remove it from the queue. The valid values for other parameters are as follows:  |
	//	|                                      | LookupId set to a nonzero value. hCursor set to 0x00000000. ulTimeout set to     |
	//	|                                      | 0x00000000.                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// * MQ_ACTION_RECEIVE
	//
	// * MQ_ACTION_PEEK_CURRENT ( pTransactionId set to NULL)
	Action uint32 `idl:"name:ulAction" json:"action"`
	// ulTimeout: Specifies the time-out, in milliseconds, to wait for a message to become
	// available in the queue. The valid value for this parameter is 0x00000000 if the LookupId
	// parameter value is nonzero or if the action is not MQ_ACTION_RECEIVE, MQ_ACTION_PEEK_CURRENT,
	// or MQ_ACTION_PEEK_NEXT.
	Timeout uint32 `idl:"name:ulTimeout" json:"timeout"`
	// dwRequestId: MUST be set by the client to a unique correlation identifier for the
	// receive request. This value MUST be used in a subsequent call to the R_EndTransactionalReceive
	// method or the R_CancelReceive method to correlate that call with the call to the
	// R_StartTransactionalReceive method. The value MUST NOT be used in another R_StartTransactionalReceive
	// method call on the same QUEUE_CONTEXT_HANDLE_NOSERIALIZE handle until a call to either
	// the R_EndTransactionalReceive method or the R_CancelReceive method with the same
	// dwRequestId parameter value has been completed.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
	// dwMaxBodySize: MUST be set by the client to the maximum size, in bytes, of the message
	// body to be returned. The server SHOULD ignore this parameter when the message is
	// not a Binary Message (section 2.2.5.1.1).
	MaxBodySize uint32 `idl:"name:dwMaxBodySize" json:"max_body_size"`
	// dwMaxCompoundMessageSize: MUST be set by the client to the maximum size, in bytes,
	// of the CompoundMessageHeader (section 2.2.5.1.2.2). The server SHOULD ignore this
	// parameter when the message is not an SRMP Message (section 2.2.5.1.2).
	MaxCompoundMessageSize uint32 `idl:"name:dwMaxCompoundMessageSize" json:"max_compound_message_size"`
	// pTransactionId: Set to NULL or set by the client to a transaction identifier that
	// was registered with the server through a prior call to the R_QMEnlistRemoteTransaction
	// (Opnum 12) (section 3.1.4.12) method.
	TransactionID *mqmq.TransactionUOW `idl:"name:pTransactionId" json:"transaction_id"`
}

func (o *StartTransactionalReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_StartTransactionalReceiveOperation) *xxx_StartTransactionalReceiveOperation {
	if op == nil {
		op = &xxx_StartTransactionalReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.LookupID = o.LookupID
	op.Cursor = o.Cursor
	op.Action = o.Action
	op.Timeout = o.Timeout
	op.RequestID = o.RequestID
	op.MaxBodySize = o.MaxBodySize
	op.MaxCompoundMessageSize = o.MaxCompoundMessageSize
	op.TransactionID = o.TransactionID
	return op
}

func (o *StartTransactionalReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_StartTransactionalReceiveOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.LookupID = op.LookupID
	o.Cursor = op.Cursor
	o.Action = op.Action
	o.Timeout = op.Timeout
	o.RequestID = op.RequestID
	o.MaxBodySize = op.MaxBodySize
	o.MaxCompoundMessageSize = op.MaxCompoundMessageSize
	o.TransactionID = op.TransactionID
}
func (o *StartTransactionalReceiveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StartTransactionalReceiveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartTransactionalReceiveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StartTransactionalReceiveResponse structure represents the R_StartTransactionalReceive operation response
type StartTransactionalReceiveResponse struct {
	// pdwArriveTime: The server MUST set this value to the time that the message was added
	// to the queue ([MS-MQDMPR] section 3.1.7.3.1), expressed as the number of seconds
	// elapsed since midnight 00:00:00.0, January 1, 1970 UTC.
	ArriveTime uint32 `idl:"name:pdwArriveTime" json:"arrive_time"`
	// pSequenceId: The server MUST set this parameter to the lower 7 bytes of the Message.LookupIdentifier
	// of the message that is received by this request.
	SequenceID uint64 `idl:"name:pSequenceId" json:"sequence_id"`
	// pdwNumberOfSections: MUST be set by the server to the number of entries in the array
	// that are pointed to by the ppPacketSections parameter.
	NumberOfSections uint32 `idl:"name:pdwNumberOfSections" json:"number_of_sections"`
	// ppPacketSections: MUST be set by the server to an array of pointers to SectionBuffer
	// (section 2.2.6) structures. The server MUST fill this array in the following manner:
	//
	// * Create two local variables of type DWORD ( ../ms-dtyp/262627d8-3418-4627-9218-4ffe110850b2
	// ) called maxMessageSize and actualMessageSize. Assign the following values to these
	// variables:
	//
	// If the message is a *Binary Message:*
	//
	// * maxMessageSize := dwMaxBodySize
	//
	// * actualMessageSize := message packet body size
	//
	// If the message is an *SRMP Message:*
	//
	// * maxMessageSize := dwMaxCompoundMessageSize
	//
	// * actualMessageSize := size in bytes of *CompoundMessageHeader*
	//
	// * If the value of maxMessageSize is greater than or equal to actualMessageSize ,
	// the ppPacketSections parameter MUST contain a single entry as follows:
	//
	// * The *SectionBufferType* member MUST be set to stFullPacket (0x00000000).
	//
	// * The *SectionSize* and *SectionSizeAlloc* members MUST be set to the message packet
	// size.
	//
	// * The *pSectionBuffer* member MUST contain the entire message packet.
	//
	// * If the value of maxMessageSize is less than actualMessageSize , the array MUST
	// contain a first entry as follows:
	//
	// The *SectionBufferType* member MUST be set to one of the following:
	//
	// * stBinaryFirstSection if the message packet is a binary packet.
	//
	// * stSrmpFirstSection if the message packet is an SRMP packet.
	//
	// * The *pSectionBuffer* member MUST contain the message packet headers ( b23ea276-dfad-4083-bc48-d0f8a40255fd#gt_65255b08-c9d4-46d3-adca-3a296d43ae4f
	// ) concatenated with the first maxMessageSize bytes of the message body.
	//
	// * The *SectionSizeAlloc* member MUST be set to the message packet headers plus actualMessageSize.
	//
	// * The *SectionSize* member MUST be set to the size of the *pSectionBuffer* member.
	//
	// * If the value of maxMessageSize is less than actualMessageSize and the message packet
	// trailers ( b23ea276-dfad-4083-bc48-d0f8a40255fd#gt_8ae0387f-5b71-4be9-a326-fd130e7bef81
	// ) are not empty, the array MUST contain a second entry as follows:
	//
	// The *SectionBufferType* member MUST be set to one of the following:
	//
	// * stBinarySecondSection if the message packet is a binary packet.
	//
	// * stSrmpSecondSection if the message packet is an SRMP packet.
	//
	// * The *pSectionBuffer* member MUST contain the message packet trailers.
	//
	// * The *SectionSize* and the *SectionSizeAlloc* members MUST be equal and MUST be
	// set to the message packet trailers size.
	//
	// * For the first entry in this array, the *pSectionBuffer* member points to a Message
	// Packet Structure (section 2.2.5) ( f9e71595-339a-4cc4-8341-371e0a4cb232 ). Within
	// this structure, set *UserMessage.BaseHeader.TimeToReachQueue* to *UserHeader.SentTime*
	// + *UserMessage.BaseHeader.TimeToReachQueue*.
	PacketSections []*SectionBuffer `idl:"name:ppPacketSections;size_is:(, pdwNumberOfSections)" json:"packet_sections"`
	// Return: The R_StartTransactionalReceive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StartTransactionalReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_StartTransactionalReceiveOperation) *xxx_StartTransactionalReceiveOperation {
	if op == nil {
		op = &xxx_StartTransactionalReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.ArriveTime = o.ArriveTime
	op.SequenceID = o.SequenceID
	op.NumberOfSections = o.NumberOfSections
	op.PacketSections = o.PacketSections
	op.Return = o.Return
	return op
}

func (o *StartTransactionalReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_StartTransactionalReceiveOperation) {
	if o == nil {
		return
	}
	o.ArriveTime = op.ArriveTime
	o.SequenceID = op.SequenceID
	o.NumberOfSections = op.NumberOfSections
	o.PacketSections = op.PacketSections
	o.Return = op.Return
}
func (o *StartTransactionalReceiveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StartTransactionalReceiveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartTransactionalReceiveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetUserAcknowledgementClassOperation structure represents the R_SetUserAcknowledgementClass operation
type xxx_SetUserAcknowledgementClassOperation struct {
	Context  *QueueNoSerialize `idl:"name:phContext" json:"context"`
	LookupID uint64            `idl:"name:LookupId" json:"lookup_id"`
	Class    uint16            `idl:"name:usClass" json:"class"`
	Return   int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_SetUserAcknowledgementClassOperation) OpNum() int { return 14 }

func (o *xxx_SetUserAcknowledgementClassOperation) OpName() string {
	return "/RemoteRead/v1/R_SetUserAcknowledgementClass"
}

func (o *xxx_SetUserAcknowledgementClassOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserAcknowledgementClassOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LookupId {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.LookupID); err != nil {
			return err
		}
	}
	// usClass {in} (1:{alias=USHORT}(uint16))
	{
		if err := w.WriteData(o.Class); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserAcknowledgementClassOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueNoSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LookupId {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.LookupID); err != nil {
			return err
		}
	}
	// usClass {in} (1:{alias=USHORT}(uint16))
	{
		if err := w.ReadData(&o.Class); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserAcknowledgementClassOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserAcknowledgementClassOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetUserAcknowledgementClassOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetUserAcknowledgementClassRequest structure represents the R_SetUserAcknowledgementClass operation request
type SetUserAcknowledgementClassRequest struct {
	// phContext: MUST be set by the client to a QUEUE_CONTEXT_HANDLE_NOSERIALIZE handle
	// representing the queue containing the message on which to set the acknowledgment
	// class. The handle MUST have been returned by the server in the pphQueue output parameter
	// of a prior call to the R_OpenQueue (Opnum 2) (section 3.1.4.2) method with the dwAccess
	// parameter set to MQ_RECEIVE_ACCESS and MUST NOT have been closed through a prior
	// call to the R_CloseQueue (Opnum 3) (section 3.1.4.3) method. This value MUST NOT
	// be NULL.
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
	// LookupId: MUST be set by the client to the lookup identifier of the message on which
	// to set the acknowledgment class.
	LookupID uint64 `idl:"name:LookupId" json:"lookup_id"`
	// usClass: The acknowledgment class to set. It MUST be set by the client to one of
	// the following values.
	//
	//	+------------------------------------------+-------------------------------------------------------+
	//	|                                          |                                                       |
	//	|                  VALUE                   |                        MEANING                        |
	//	|                                          |                                                       |
	//	+------------------------------------------+-------------------------------------------------------+
	//	+------------------------------------------+-------------------------------------------------------+
	//	| 0x0000                                   | No-op. No change is made to the acknowledgment class. |
	//	+------------------------------------------+-------------------------------------------------------+
	//	| MQMSG_CLASS_NACK_RECEIVE_REJECTED 0xC004 | Marks the message as rejected.                        |
	//	+------------------------------------------+-------------------------------------------------------+
	Class uint16 `idl:"name:usClass" json:"class"`
}

func (o *SetUserAcknowledgementClassRequest) xxx_ToOp(ctx context.Context, op *xxx_SetUserAcknowledgementClassOperation) *xxx_SetUserAcknowledgementClassOperation {
	if op == nil {
		op = &xxx_SetUserAcknowledgementClassOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.LookupID = o.LookupID
	op.Class = o.Class
	return op
}

func (o *SetUserAcknowledgementClassRequest) xxx_FromOp(ctx context.Context, op *xxx_SetUserAcknowledgementClassOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.LookupID = op.LookupID
	o.Class = op.Class
}
func (o *SetUserAcknowledgementClassRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetUserAcknowledgementClassRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetUserAcknowledgementClassOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetUserAcknowledgementClassResponse structure represents the R_SetUserAcknowledgementClass operation response
type SetUserAcknowledgementClassResponse struct {
	// Return: The R_SetUserAcknowledgementClass return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetUserAcknowledgementClassResponse) xxx_ToOp(ctx context.Context, op *xxx_SetUserAcknowledgementClassOperation) *xxx_SetUserAcknowledgementClassOperation {
	if op == nil {
		op = &xxx_SetUserAcknowledgementClassOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetUserAcknowledgementClassResponse) xxx_FromOp(ctx context.Context, op *xxx_SetUserAcknowledgementClassOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetUserAcknowledgementClassResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetUserAcknowledgementClassResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetUserAcknowledgementClassOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EndTransactionalReceiveOperation structure represents the R_EndTransactionalReceive operation
type xxx_EndTransactionalReceiveOperation struct {
	Context   *QueueNoSerialize `idl:"name:phContext" json:"context"`
	Ack       uint32            `idl:"name:dwAck" json:"ack"`
	RequestID uint32            `idl:"name:dwRequestId" json:"request_id"`
	Return    int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_EndTransactionalReceiveOperation) OpNum() int { return 15 }

func (o *xxx_EndTransactionalReceiveOperation) OpName() string {
	return "/RemoteRead/v1/R_EndTransactionalReceive"
}

func (o *xxx_EndTransactionalReceiveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Ack < uint32(1) || o.Ack > uint32(2) {
		return fmt.Errorf("Ack is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndTransactionalReceiveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwAck {in} (1:{range=(1,2), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Ack); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndTransactionalReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=QUEUE_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &QueueNoSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwAck {in} (1:{range=(1,2), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Ack); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndTransactionalReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndTransactionalReceiveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EndTransactionalReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EndTransactionalReceiveRequest structure represents the R_EndTransactionalReceive operation request
type EndTransactionalReceiveRequest struct {
	// phContext: MUST be set by the client to the QUEUE_CONTEXT_HANDLE_NOSERIALIZE structure
	// used in the corresponding call to the R_StartTransactionalReceive method. The handle
	// MUST have been returned by the server in the pphQueue output parameter of a prior
	// call to the R_OpenQueue (Opnum 2) (section 3.1.4.2) method and MUST NOT have been
	// closed through a prior call to the R_CloseQueue (Opnum 3) (section 3.1.4.3) method.
	// This value MUST NOT be NULL.
	Context *QueueNoSerialize `idl:"name:phContext" json:"context"`
	// dwAck: MUST be set to an Acknowledgment (ACK) or a Negative Acknowledgment (NACK)
	// for the message packet received from the server in an R_StartTransactionalReceive
	// method request. The following table lists possible values.
	//
	//	+--------------------+----------------------------------------------------------------------------------+
	//	|                    |                                                                                  |
	//	|       VALUE        |                                     MEANING                                      |
	//	|                    |                                                                                  |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| RR_ACK 0x00000002  | The client acknowledges that the message packet was received successfully. The   |
	//	|                    | server MUST NOT remove the packet from the queue. The server removes the packet  |
	//	|                    | from the queue when the transaction is committed.                                |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| RR_NACK 0x00000001 | The client acknowledges that the message packet was not received successfully.   |
	//	|                    | The server MUST keep the message packet and make it available for subsequent     |
	//	|                    | consumption.                                                                     |
	//	+--------------------+----------------------------------------------------------------------------------+
	Ack uint32 `idl:"name:dwAck" json:"ack"`
	// dwRequestId: MUST be set by the client to the same value as the dwRequestId  parameter
	// in the corresponding call to the R_StartTransactionalReceive method. This parameter
	// acts as an identifier to correlate an R_EndTransactionalReceive method call to an
	// R_StartTransactionalReceive method call.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
}

func (o *EndTransactionalReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_EndTransactionalReceiveOperation) *xxx_EndTransactionalReceiveOperation {
	if op == nil {
		op = &xxx_EndTransactionalReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Ack = o.Ack
	op.RequestID = o.RequestID
	return op
}

func (o *EndTransactionalReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_EndTransactionalReceiveOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Ack = op.Ack
	o.RequestID = op.RequestID
}
func (o *EndTransactionalReceiveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EndTransactionalReceiveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EndTransactionalReceiveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EndTransactionalReceiveResponse structure represents the R_EndTransactionalReceive operation response
type EndTransactionalReceiveResponse struct {
	// Return: The R_EndTransactionalReceive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EndTransactionalReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_EndTransactionalReceiveOperation) *xxx_EndTransactionalReceiveOperation {
	if op == nil {
		op = &xxx_EndTransactionalReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *EndTransactionalReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_EndTransactionalReceiveOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EndTransactionalReceiveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EndTransactionalReceiveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EndTransactionalReceiveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
