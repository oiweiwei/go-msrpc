package qm2qm

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
)

var (
	// import guard
	GoPackage = "mqqp"
)

var (
	// Syntax UUID
	Qm2qmSyntaxUUID = &uuid.UUID{TimeLow: 0x1088a980, TimeMid: 0xeae5, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x9b, Node: [6]uint8{0x0, 0xa0, 0x24, 0x53, 0xc3, 0x37}}
	// Syntax ID
	Qm2qmSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: Qm2qmSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// qm2qm interface.
type Qm2qmClient interface {

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
	StartReceive(context.Context, *StartReceiveRequest, ...dcerpc.CallOption) (*StartReceiveResponse, error)

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
	EndReceive(context.Context, *EndReceiveRequest, ...dcerpc.CallOption) (*EndReceiveResponse, error)

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
	OpenQueue(context.Context, *OpenQueueRequest, ...dcerpc.CallOption) (*OpenQueueResponse, error)

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
	CloseQueue(context.Context, *CloseQueueRequest, ...dcerpc.CallOption) (*CloseQueueResponse, error)

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
	CloseCursor(context.Context, *CloseCursorRequest, ...dcerpc.CallOption) (*CloseCursorResponse, error)

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
	CancelReceive(context.Context, *CancelReceiveRequest, ...dcerpc.CallOption) (*CancelReceiveResponse, error)

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
	PurgeQueue(context.Context, *PurgeQueueRequest, ...dcerpc.CallOption) (*PurgeQueueResponse, error)

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
	GetQMQMServerPort(context.Context, *GetQMQMServerPortRequest, ...dcerpc.CallOption) (*GetQMQMServerPortResponse, error)

	// The RemoteQmGetVersion method retrieves the Message queuing version of the server;
	// this method is called before the RemoteQMOpenQueue method.<21>
	//
	// Return Values: This method has no return values.
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetVersion(context.Context, *GetVersionRequest, ...dcerpc.CallOption) (*GetVersionResponse, error)

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
	StartReceive2(context.Context, *StartReceive2Request, ...dcerpc.CallOption) (*StartReceive2Response, error)

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
	StartReceiveByLookupID(context.Context, *StartReceiveByLookupIDRequest, ...dcerpc.CallOption) (*StartReceiveByLookupIDResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Session structure represents PCTX_RRSESSION_HANDLE_TYPE RPC structure.
type Session dcetypes.ContextHandle

func (o *Session) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Session) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Session) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Session) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Context structure represents PCTX_REMOTEREAD_HANDLE_TYPE RPC structure.
type Context dcetypes.ContextHandle

func (o *Context) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Context) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Context) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Context) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RemoteReadAck type represents REMOTEREADACK RPC enumeration.
//
// The REMOTEREADACK enumeration represents an acknowledgment (ACK) or a negative acknowledgment
// (NACK), indicating a successfully or an unsuccessfully delivered packet, respectively.
type RemoteReadAck uint16

var (
	// RR_UNKNOWN:  No acknowledgment.
	RemoteReadAckUnknown RemoteReadAck = 0
	// RR_NACK:  Negative acknowledgment for a packet.
	RemoteReadAckNack RemoteReadAck = 1
	// RR_ACK:  Acknowledgment for a packet.
	RemoteReadAckAck RemoteReadAck = 2
)

func (o RemoteReadAck) String() string {
	switch o {
	case RemoteReadAckUnknown:
		return "RemoteReadAckUnknown"
	case RemoteReadAckNack:
		return "RemoteReadAckNack"
	case RemoteReadAckAck:
		return "RemoteReadAckAck"
	}
	return "Invalid"
}

// RemoteReadDesc structure represents REMOTEREADDESC RPC structure.
//
// This structure is used to encapsulate the information necessary to perform operations
// RemoteQMStartReceive, RemoteQMStartReceive2, and RemoteQMStartReceiveByLookupId.
type RemoteReadDesc struct {
	RemoteQueue uint32 `idl:"name:hRemoteQueue" json:"remote_queue"`
	Cursor      uint32 `idl:"name:hCursor" json:"cursor"`
	// ulAction:   The following table describes possible actions. The Peek and Receive
	// operations both enable access to the contents of a message. This value is set by
	// the client.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                   TYPE/MEANING                                   |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_RECEIVE 0x00000000         | Type = Receive Reads and removes a message from the current cursor location if   |
	//	|                                      | hCursor is nonzero or from the front of the queue if hCursor is set to zero.     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_CURRENT 0x80000000    | Type = Peek Reads a message from the current cursor location if hCursor is       |
	//	|                                      | nonzero or from the front of the queue if hCursor is set to zero but does not    |
	//	|                                      | remove it from the queue.                                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_NEXT 0x80000001       | Type = Peek Reads a message following the message at the current cursor location |
	//	|                                      | but does not remove it from the queue.                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_CURRENT 0x40000010    | Type = Peek Reads the message specified by a lookup identifier but does not      |
	//	|                                      | remove it from the queue.                                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_NEXT 0x40000011       | Type = Peek Reads the message following the message specified by a lookup        |
	//	|                                      | identifier but does not remove it from the queue.                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_PREV 0x40000012       | Type = Peek Reads the message preceding the message specified by a lookup        |
	//	|                                      | identifier but does not remove it from the queue.                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_CURRENT 0x40000020 | Type = Receive Reads the message specified by a lookup identifier and removes it |
	//	|                                      | from the queue.                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_NEXT 0x40000021    | Type = Receive Reads the message following the message specified by a lookup     |
	//	|                                      | identifier and removes it from the queue.                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_PREV 0x40000022    | Type = Receive Reads the message preceding the message specified by a lookup     |
	//	|                                      | identifier and removes it from the queue.                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Action uint32 `idl:"name:ulAction" json:"action"`
	// ulTimeout:   Specifies a time-out in milliseconds for the server to wait for a message
	// to become available in the queue. This value is set by the client. To specify an
	// infinite time-out, the client MUST set this field to 0xFFFFFFFF.
	Timeout uint32 `idl:"name:ulTimeout" json:"timeout"`
	// dwSize:   Specifies the size, in bytes, of lpBuffer. The valid range is 0 to 0x00420000.
	// This value is set by the server and MUST be set to 0 by the client.
	Size  uint32 `idl:"name:dwSize" json:"size"`
	Queue uint32 `idl:"name:dwQueue" json:"queue"`
	// dwRequestID:  The client MUST set this member to a unique identifier for the receive
	// request, which could later be used to identify and cancel the receive request. This
	// value is set by the client. The client SHOULD NOT<6> reuse this identifier until
	// a call to the RemoteQMEndReceive (Opnum 1) method or to the RemoteQMCancelReceive
	// (Opnum 5) method has been made for that receive request.
	RequestID uint32 `idl:"name:dwRequestID" json:"request_id"`
	// Reserved:  This is a reserved field of type DWORD that MUST be ignored.
	//
	//	+------------+---------------------+
	//	|            |                     |
	//	|   VALUE    |       MEANING       |
	//	|            |                     |
	//	+------------+---------------------+
	//	+------------+---------------------+
	//	| 0x00000000 | Returned by client. |
	//	+------------+---------------------+
	//	| 0x00000001 | Returned by server. |
	//	+------------+---------------------+
	_ uint32 `idl:"name:Reserved"`
	// dwArriveTime:   The server MUST set this value to the time that the message was added
	// to the queue. The time MUST be expressed as the number of seconds elapsed since 00:00:00.0,
	// January 1, 1970 Coordinated Universal Time (UTC).
	ArriveTime uint32 `idl:"name:dwArriveTime" json:"arrive_time"`
	// eAckNack:  This is a reserved field and MUST be ignored by the client and the server.
	AckNack RemoteReadAck `idl:"name:eAckNack" json:"ack_nack"`
	// lpBuffer:  This field represents a pointer to a buffer containing the UserMessage
	// Packet ([MS-MQMQ] section 2.2.20). The size of this field is specified by dwSize.
	// This value is set by the server and MUST be set to NULL by the client.
	Buffer []byte `idl:"name:lpBuffer;size_is:(dwSize);length_is:(dwSize);pointer:unique" json:"buffer"`
}

func (o *RemoteReadDesc) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Buffer != nil && o.Size == 0 {
		o.Size = uint32(len(o.Buffer))
	}
	if o.Buffer != nil && o.Size == 0 {
		o.Size = uint32(len(o.Buffer))
	}
	if o.Size > uint32(4325376) {
		return fmt.Errorf("Size is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RemoteReadDesc) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteQueue); err != nil {
		return err
	}
	if err := w.WriteData(o.Cursor); err != nil {
		return err
	}
	if err := w.WriteData(o.Action); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Queue); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestID); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.ArriveTime); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.AckNack)); err != nil {
		return err
	}
	if o.Buffer != nil || o.Size > 0 {
		_ptr_lpBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64(o.Size)
			if dimLength1 > sizeInfo[0] {
				dimLength1 = sizeInfo[0]
			} else {
				sizeInfo[0] = dimLength1
			}
			if err := w.WriteSize(0); err != nil {
				return err
			}
			if err := w.WriteSize(dimLength1); err != nil {
				return err
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Buffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_lpBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RemoteReadDesc) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteQueue); err != nil {
		return err
	}
	if err := w.ReadData(&o.Cursor); err != nil {
		return err
	}
	if err := w.ReadData(&o.Action); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Queue); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestID); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint32
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.ArriveTime); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.AckNack)); err != nil {
		return err
	}
	_ptr_lpBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpBuffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_lpBuffer, _ptr_lpBuffer); err != nil {
		return err
	}
	return nil
}

// RemoteReadDesc2 structure represents REMOTEREADDESC2 RPC structure.
//
// This structure is used by RemoteQMStartReceive2 and RemoteQMStartReceiveByLookupId
// to encapsulate the parameters necessary for execution of these operations.
type RemoteReadDesc2 struct {
	// pRemoteReadDesc:   A pointer to a REMOTEREADDESC structure, as specified in section
	// 2.2.2.1.
	RemoteReadDesc *RemoteReadDesc `idl:"name:pRemoteReadDesc" json:"remote_read_desc"`
	// SequentialId:   This field is set by the server to the value of a unique message
	// identifier that corresponds to a received message.
	SequentialID uint64 `idl:"name:SequentialId" json:"sequential_id"`
}

func (o *RemoteReadDesc2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RemoteReadDesc2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.RemoteReadDesc != nil {
		_ptr_pRemoteReadDesc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RemoteReadDesc != nil {
				if err := o.RemoteReadDesc.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RemoteReadDesc{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteReadDesc, _ptr_pRemoteReadDesc); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SequentialID); err != nil {
		return err
	}
	return nil
}
func (o *RemoteReadDesc2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	_ptr_pRemoteReadDesc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RemoteReadDesc == nil {
			o.RemoteReadDesc = &RemoteReadDesc{}
		}
		if err := o.RemoteReadDesc.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pRemoteReadDesc := func(ptr interface{}) { o.RemoteReadDesc = *ptr.(**RemoteReadDesc) }
	if err := w.ReadPointer(&o.RemoteReadDesc, _s_pRemoteReadDesc, _ptr_pRemoteReadDesc); err != nil {
		return err
	}
	if err := w.ReadData(&o.SequentialID); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultQm2qmClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultQm2qmClient) StartReceive(ctx context.Context, in *StartReceiveRequest, opts ...dcerpc.CallOption) (*StartReceiveResponse, error) {
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

func (o *xxx_DefaultQm2qmClient) EndReceive(ctx context.Context, in *EndReceiveRequest, opts ...dcerpc.CallOption) (*EndReceiveResponse, error) {
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

func (o *xxx_DefaultQm2qmClient) OpenQueue(ctx context.Context, in *OpenQueueRequest, opts ...dcerpc.CallOption) (*OpenQueueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenQueueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQm2qmClient) CloseQueue(ctx context.Context, in *CloseQueueRequest, opts ...dcerpc.CallOption) (*CloseQueueResponse, error) {
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

func (o *xxx_DefaultQm2qmClient) CloseCursor(ctx context.Context, in *CloseCursorRequest, opts ...dcerpc.CallOption) (*CloseCursorResponse, error) {
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

func (o *xxx_DefaultQm2qmClient) CancelReceive(ctx context.Context, in *CancelReceiveRequest, opts ...dcerpc.CallOption) (*CancelReceiveResponse, error) {
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

func (o *xxx_DefaultQm2qmClient) PurgeQueue(ctx context.Context, in *PurgeQueueRequest, opts ...dcerpc.CallOption) (*PurgeQueueResponse, error) {
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

func (o *xxx_DefaultQm2qmClient) GetQMQMServerPort(ctx context.Context, in *GetQMQMServerPortRequest, opts ...dcerpc.CallOption) (*GetQMQMServerPortResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetQMQMServerPortResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQm2qmClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...dcerpc.CallOption) (*GetVersionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetVersionResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultQm2qmClient) StartReceive2(ctx context.Context, in *StartReceive2Request, opts ...dcerpc.CallOption) (*StartReceive2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StartReceive2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQm2qmClient) StartReceiveByLookupID(ctx context.Context, in *StartReceiveByLookupIDRequest, opts ...dcerpc.CallOption) (*StartReceiveByLookupIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StartReceiveByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQm2qmClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQm2qmClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewQm2qmClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Qm2qmClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Qm2qmSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultQm2qmClient{cc: cc}, nil
}

// xxx_StartReceiveOperation structure represents the RemoteQMStartReceive operation
type xxx_StartReceiveOperation struct {
	Context        *Context        `idl:"name:pphContext" json:"context"`
	RemoteReadDesc *RemoteReadDesc `idl:"name:lpRemoteReadDesc" json:"remote_read_desc"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_StartReceiveOperation) OpNum() int { return 0 }

func (o *xxx_StartReceiveOperation) OpName() string { return "/qm2qm/v1/RemoteQMStartReceive" }

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
	// lpRemoteReadDesc {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC}(struct))
	{
		if o.RemoteReadDesc != nil {
			if err := o.RemoteReadDesc.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteReadDesc{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpRemoteReadDesc {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC}(struct))
	{
		if o.RemoteReadDesc == nil {
			o.RemoteReadDesc = &RemoteReadDesc{}
		}
		if err := o.RemoteReadDesc.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
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
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_REMOTEREAD_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpRemoteReadDesc {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC}(struct))
	{
		if o.RemoteReadDesc != nil {
			if err := o.RemoteReadDesc.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteReadDesc{}).MarshalNDR(ctx, w); err != nil {
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
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_REMOTEREAD_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpRemoteReadDesc {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC}(struct))
	{
		if o.RemoteReadDesc == nil {
			o.RemoteReadDesc = &RemoteReadDesc{}
		}
		if err := o.RemoteReadDesc.UnmarshalNDR(ctx, w); err != nil {
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

// StartReceiveRequest structure represents the RemoteQMStartReceive operation request
type StartReceiveRequest struct {
	// lpRemoteReadDesc: A pointer to an instance of a REMOTEREADDESC (section 2.2.2.1)
	// structure.
	//
	// In addition, the ulAction member of the lpRemoteReadDesc parameter MUST be one of
	// the following values.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|             VALUE OF              |                                                                                  |
	//	|             ULACTION              |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_RECEIVE 0x00000000      | If hCursor is nonzero, read and remove the first message available at the        |
	//	|                                   | current cursor's location walking towards the end of the queue. If hCursor is    |
	//	|                                   | zero, read and remove the message from the front of the queue.                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_CURRENT 0x80000000 | If hCursor is nonzero, read the message at the current cursor location, but      |
	//	|                                   | do not remove it from the queue. The cursor location does not change after the   |
	//	|                                   | operation. If hCursor is zero, read the message at the front of the queue, but   |
	//	|                                   | do not remove it from the queue.                                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_NEXT 0x80000001    | Read the message following the message at the current cursor location, but       |
	//	|                                   | do not remove it. The cursor location will then change to the next available     |
	//	|                                   | message, walking towards the end of the queue. The hCursor parameter MUST be set |
	//	|                                   | to a nonzero cursor handle.                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//
	// The hCursor member of lpRemoteReadDesc specifies a handle to an opened cursor. A
	// value of zero indicates that a cursor is not used for this operation.
	//
	// The dwRequestID member of the lpRemoteReadDesc parameter is used in a subsequent
	// call to RemoteQMCancelReceive to correlate that call with the call to RemoteQMStartReceive.
	RemoteReadDesc *RemoteReadDesc `idl:"name:lpRemoteReadDesc" json:"remote_read_desc"`
}

func (o *StartReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_StartReceiveOperation) *xxx_StartReceiveOperation {
	if op == nil {
		op = &xxx_StartReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.RemoteReadDesc = o.RemoteReadDesc
	return op
}

func (o *StartReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_StartReceiveOperation) {
	if o == nil {
		return
	}
	o.RemoteReadDesc = op.RemoteReadDesc
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

// StartReceiveResponse structure represents the RemoteQMStartReceive operation response
type StartReceiveResponse struct {
	// pphContext: The server MUST return a non-NULL value for this handle upon success
	// for receive calls. This handle will be used by the client in subsequent calls to
	// RemoteQMEndReceive. This handle MUST NOT be set upon failure, or for peek calls.
	// If this method returns an error, pphContext is undefined and MUST NOT be used as
	// an argument for a call to RemoteQMEndReceive.
	Context *Context `idl:"name:pphContext" json:"context"`
	// lpRemoteReadDesc: A pointer to an instance of a REMOTEREADDESC (section 2.2.2.1)
	// structure.
	//
	// In addition, the ulAction member of the lpRemoteReadDesc parameter MUST be one of
	// the following values.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|             VALUE OF              |                                                                                  |
	//	|             ULACTION              |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_RECEIVE 0x00000000      | If hCursor is nonzero, read and remove the first message available at the        |
	//	|                                   | current cursor's location walking towards the end of the queue. If hCursor is    |
	//	|                                   | zero, read and remove the message from the front of the queue.                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_CURRENT 0x80000000 | If hCursor is nonzero, read the message at the current cursor location, but      |
	//	|                                   | do not remove it from the queue. The cursor location does not change after the   |
	//	|                                   | operation. If hCursor is zero, read the message at the front of the queue, but   |
	//	|                                   | do not remove it from the queue.                                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_NEXT 0x80000001    | Read the message following the message at the current cursor location, but       |
	//	|                                   | do not remove it. The cursor location will then change to the next available     |
	//	|                                   | message, walking towards the end of the queue. The hCursor parameter MUST be set |
	//	|                                   | to a nonzero cursor handle.                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//
	// The hCursor member of lpRemoteReadDesc specifies a handle to an opened cursor. A
	// value of zero indicates that a cursor is not used for this operation.
	//
	// The dwRequestID member of the lpRemoteReadDesc parameter is used in a subsequent
	// call to RemoteQMCancelReceive to correlate that call with the call to RemoteQMStartReceive.
	RemoteReadDesc *RemoteReadDesc `idl:"name:lpRemoteReadDesc" json:"remote_read_desc"`
	// Return: The RemoteQMStartReceive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StartReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_StartReceiveOperation) *xxx_StartReceiveOperation {
	if op == nil {
		op = &xxx_StartReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.RemoteReadDesc = o.RemoteReadDesc
	op.Return = o.Return
	return op
}

func (o *StartReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_StartReceiveOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.RemoteReadDesc = op.RemoteReadDesc
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

// xxx_EndReceiveOperation structure represents the RemoteQMEndReceive operation
type xxx_EndReceiveOperation struct {
	Context *Context `idl:"name:pphContext" json:"context"`
	Ack     uint32   `idl:"name:dwAck" json:"ack"`
	Return  int32    `idl:"name:Return" json:"return"`
}

func (o *xxx_EndReceiveOperation) OpNum() int { return 1 }

func (o *xxx_EndReceiveOperation) OpName() string { return "/qm2qm/v1/RemoteQMEndReceive" }

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
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_REMOTEREAD_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
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
	return nil
}

func (o *xxx_EndReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_REMOTEREAD_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
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
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_REMOTEREAD_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EndReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_REMOTEREAD_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
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

// EndReceiveRequest structure represents the RemoteQMEndReceive operation request
type EndReceiveRequest struct {
	// pphContext: A pointer to a context handle of a pending remote read operation.
	Context *Context `idl:"name:pphContext" json:"context"`
	// dwAck: An ACK or NACK about the status of the message packet of the pending remote
	// read operation.
	//
	//	+--------------------+----------------------------------------------------------------------------------+
	//	|                    |                                                                                  |
	//	|       VALUE        |                                     MEANING                                      |
	//	|                    |                                                                                  |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| RR_NACK 0x00000001 | The client acknowledges that the message packet was not delivered successfully.  |
	//	|                    | The server MUST keep the message in the queue and make it available for          |
	//	|                    | subsequent consumption.                                                          |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| RR_ACK 0x00000002  | The client acknowledges that the message packet was delivered successfully.      |
	//	|                    | The server MUST remove the message from the queue and make it unavailable for    |
	//	|                    | subsequent consumption.                                                          |
	//	+--------------------+----------------------------------------------------------------------------------+
	Ack uint32 `idl:"name:dwAck" json:"ack"`
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
	return op
}

func (o *EndReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_EndReceiveOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Ack = op.Ack
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

// EndReceiveResponse structure represents the RemoteQMEndReceive operation response
type EndReceiveResponse struct {
	// pphContext: A pointer to a context handle of a pending remote read operation.
	Context *Context `idl:"name:pphContext" json:"context"`
	// Return: The RemoteQMEndReceive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EndReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_EndReceiveOperation) *xxx_EndReceiveOperation {
	if op == nil {
		op = &xxx_EndReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Return = o.Return
	return op
}

func (o *EndReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_EndReceiveOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
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

// xxx_OpenQueueOperation structure represents the RemoteQMOpenQueue operation
type xxx_OpenQueueOperation struct {
	Context    *Session   `idl:"name:phContext" json:"context"`
	ClientGUID *dtyp.GUID `idl:"name:pLicGuid" json:"client_guid"`
	MQS        uint32     `idl:"name:dwMQS" json:"mqs"`
	Queue      uint32     `idl:"name:hQueue" json:"queue"`
	QueueID    uint32     `idl:"name:pQueue" json:"queue_id"`
	ContextID  uint32     `idl:"name:dwpContext" json:"context_id"`
	Return     int32      `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenQueueOperation) OpNum() int { return 2 }

func (o *xxx_OpenQueueOperation) OpName() string { return "/qm2qm/v1/RemoteQMOpenQueue" }

func (o *xxx_OpenQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.MQS > uint32(16) {
		return fmt.Errorf("MQS is out of range")
	}
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
	// pLicGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClientGUID != nil {
			if err := o.ClientGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwMQS {in} (1:{range=(0,16), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MQS); err != nil {
			return err
		}
	}
	// hQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Queue); err != nil {
			return err
		}
	}
	// pQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QueueID); err != nil {
			return err
		}
	}
	// dwpContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ContextID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pLicGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClientGUID == nil {
			o.ClientGUID = &dtyp.GUID{}
		}
		if err := o.ClientGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwMQS {in} (1:{range=(0,16), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MQS); err != nil {
			return err
		}
	}
	// hQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Queue); err != nil {
			return err
		}
	}
	// pQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QueueID); err != nil {
			return err
		}
	}
	// dwpContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ContextID); err != nil {
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
	// phContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_RRSESSION_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Session{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_RRSESSION_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Session{}
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

// OpenQueueRequest structure represents the RemoteQMOpenQueue operation request
type OpenQueueRequest struct {
	// pLicGuid: A pointer to a valid GUID ([MS-DTYP] section 2.3.4) that uniquely identifies
	// the client. This value is set to the QueueManager.Identifier ADM element of the queue
	// manager at the client end.
	ClientGUID *dtyp.GUID `idl:"name:pLicGuid" json:"client_guid"`
	// dwMQS: This value MAY be used by the server to impose an implementation-specific
	// limit on the number of concurrent callers.<12>
	MQS uint32 `idl:"name:dwMQS" json:"mqs"`
	// hQueue: A queue identifier. This value SHOULD be ignored by the server.<13>
	Queue uint32 `idl:"name:hQueue" json:"queue"`
	// pQueue: A DWORD that references an OpenQueueDescriptor of a remote opened queue.
	QueueID uint32 `idl:"name:pQueue" json:"queue_id"`
	// dwpContext: A DWORD that references an OpenQueueDescriptor of a remote opened queue.
	ContextID uint32 `idl:"name:dwpContext" json:"context_id"`
}

func (o *OpenQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenQueueOperation) *xxx_OpenQueueOperation {
	if op == nil {
		op = &xxx_OpenQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.ClientGUID = o.ClientGUID
	op.MQS = o.MQS
	op.Queue = o.Queue
	op.QueueID = o.QueueID
	op.ContextID = o.ContextID
	return op
}

func (o *OpenQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenQueueOperation) {
	if o == nil {
		return
	}
	o.ClientGUID = op.ClientGUID
	o.MQS = op.MQS
	o.Queue = op.Queue
	o.QueueID = op.QueueID
	o.ContextID = op.ContextID
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

// OpenQueueResponse structure represents the RemoteQMOpenQueue operation response
type OpenQueueResponse struct {
	// phContext:  A pointer to a context handle that contains the information about the
	// opened queue, which corresponds to the abstract data model's OpenQueueEntry. The
	// server MUST set this value; it gets deleted on a call to RemoteQMCloseQueue.
	Context *Session `idl:"name:phContext" json:"context"`
	// Return: The RemoteQMOpenQueue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenQueueResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenQueueOperation) *xxx_OpenQueueOperation {
	if op == nil {
		op = &xxx_OpenQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Return = o.Return
	return op
}

func (o *OpenQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenQueueOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Return = op.Return
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

// xxx_CloseQueueOperation structure represents the RemoteQMCloseQueue operation
type xxx_CloseQueueOperation struct {
	Context *Session `idl:"name:pphContext" json:"context"`
	Return  int32    `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseQueueOperation) OpNum() int { return 3 }

func (o *xxx_CloseQueueOperation) OpName() string { return "/qm2qm/v1/RemoteQMCloseQueue" }

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
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_RRSESSION_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_RRSESSION_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Session{}
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
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_RRSESSION_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Session{}).MarshalNDR(ctx, w); err != nil {
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
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_RRSESSION_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Session{}
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

// CloseQueueRequest structure represents the RemoteQMCloseQueue operation request
type CloseQueueRequest struct {
	// pphContext: A PCTX_RRSESSION_HANDLE_TYPE to a remote opened queue.
	Context *Session `idl:"name:pphContext" json:"context"`
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

// CloseQueueResponse structure represents the RemoteQMCloseQueue operation response
type CloseQueueResponse struct {
	// pphContext: A PCTX_RRSESSION_HANDLE_TYPE to a remote opened queue.
	Context *Session `idl:"name:pphContext" json:"context"`
	// Return: The RemoteQMCloseQueue return value.
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

// xxx_CloseCursorOperation structure represents the RemoteQMCloseCursor operation
type xxx_CloseCursorOperation struct {
	Queue  uint32 `idl:"name:hQueue" json:"queue"`
	Cursor uint32 `idl:"name:hCursor" json:"cursor"`
	Return int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseCursorOperation) OpNum() int { return 4 }

func (o *xxx_CloseCursorOperation) OpName() string { return "/qm2qm/v1/RemoteQMCloseCursor" }

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
	// hQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Queue); err != nil {
			return err
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
	// hQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Queue); err != nil {
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

// CloseCursorRequest structure represents the RemoteQMCloseCursor operation request
type CloseCursorRequest struct {
	// hQueue:  A queue handle value upon which the cursor operates.
	Queue uint32 `idl:"name:hQueue" json:"queue"`
	// hCursor: Specifies the cursor handle to be closed.
	Cursor uint32 `idl:"name:hCursor" json:"cursor"`
}

func (o *CloseCursorRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseCursorOperation) *xxx_CloseCursorOperation {
	if op == nil {
		op = &xxx_CloseCursorOperation{}
	}
	if o == nil {
		return op
	}
	op.Queue = o.Queue
	op.Cursor = o.Cursor
	return op
}

func (o *CloseCursorRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseCursorOperation) {
	if o == nil {
		return
	}
	o.Queue = op.Queue
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

// CloseCursorResponse structure represents the RemoteQMCloseCursor operation response
type CloseCursorResponse struct {
	// Return: The RemoteQMCloseCursor return value.
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

// xxx_CancelReceiveOperation structure represents the RemoteQMCancelReceive operation
type xxx_CancelReceiveOperation struct {
	Queue     uint32 `idl:"name:hQueue" json:"queue"`
	QueueID   uint32 `idl:"name:pQueue" json:"queue_id"`
	RequestID uint32 `idl:"name:dwRequestID" json:"request_id"`
	Return    int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelReceiveOperation) OpNum() int { return 5 }

func (o *xxx_CancelReceiveOperation) OpName() string { return "/qm2qm/v1/RemoteQMCancelReceive" }

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
	// hQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Queue); err != nil {
			return err
		}
	}
	// pQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QueueID); err != nil {
			return err
		}
	}
	// dwRequestID {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Queue); err != nil {
			return err
		}
	}
	// pQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QueueID); err != nil {
			return err
		}
	}
	// dwRequestID {in} (1:{alias=DWORD}(uint32))
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

// CancelReceiveRequest structure represents the RemoteQMCancelReceive operation request
type CancelReceiveRequest struct {
	// hQueue:  Queue identifier to cancel receive. Its value is validated in the method's
	// processing rules.
	Queue uint32 `idl:"name:hQueue" json:"queue"`
	// pQueue:  Queue descriptor to cancel receive. Its value is validated in the method's
	// processing rules.
	QueueID uint32 `idl:"name:pQueue" json:"queue_id"`
	// dwRequestID: A unique value that identifies a pending remote read operation.
	RequestID uint32 `idl:"name:dwRequestID" json:"request_id"`
}

func (o *CancelReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_CancelReceiveOperation) *xxx_CancelReceiveOperation {
	if op == nil {
		op = &xxx_CancelReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.Queue = o.Queue
	op.QueueID = o.QueueID
	op.RequestID = o.RequestID
	return op
}

func (o *CancelReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelReceiveOperation) {
	if o == nil {
		return
	}
	o.Queue = op.Queue
	o.QueueID = op.QueueID
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

// CancelReceiveResponse structure represents the RemoteQMCancelReceive operation response
type CancelReceiveResponse struct {
	// Return: The RemoteQMCancelReceive return value.
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

// xxx_PurgeQueueOperation structure represents the RemoteQMPurgeQueue operation
type xxx_PurgeQueueOperation struct {
	Queue  uint32 `idl:"name:hQueue" json:"queue"`
	Return int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_PurgeQueueOperation) OpNum() int { return 6 }

func (o *xxx_PurgeQueueOperation) OpName() string { return "/qm2qm/v1/RemoteQMPurgeQueue" }

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
	// hQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Queue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PurgeQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Queue); err != nil {
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

// PurgeQueueRequest structure represents the RemoteQMPurgeQueue operation request
type PurgeQueueRequest struct {
	// hQueue: A queue handle value acquired from the phQueue parameter of the qmcomm:R_QMOpenRemoteQueue
	// method as specified in [MS-MQMP] section 3.1.4.2.
	Queue uint32 `idl:"name:hQueue" json:"queue"`
}

func (o *PurgeQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_PurgeQueueOperation) *xxx_PurgeQueueOperation {
	if op == nil {
		op = &xxx_PurgeQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.Queue = o.Queue
	return op
}

func (o *PurgeQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_PurgeQueueOperation) {
	if o == nil {
		return
	}
	o.Queue = op.Queue
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

// PurgeQueueResponse structure represents the RemoteQMPurgeQueue operation response
type PurgeQueueResponse struct {
	// Return: The RemoteQMPurgeQueue return value.
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

// xxx_GetQMQMServerPortOperation structure represents the RemoteQMGetQMQMServerPort operation
type xxx_GetQMQMServerPortOperation struct {
	PortType uint32 `idl:"name:dwPortType" json:"port_type"`
	Return   uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQMQMServerPortOperation) OpNum() int { return 7 }

func (o *xxx_GetQMQMServerPortOperation) OpName() string {
	return "/qm2qm/v1/RemoteQMGetQMQMServerPort"
}

func (o *xxx_GetQMQMServerPortOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.PortType > uint32(3) {
		return fmt.Errorf("PortType is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQMQMServerPortOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPortType {in} (1:{range=(0,3), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PortType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQMQMServerPortOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPortType {in} (1:{range=(0,3), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PortType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQMQMServerPortOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQMQMServerPortOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQMQMServerPortOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetQMQMServerPortRequest structure represents the RemoteQMGetQMQMServerPort operation request
type GetQMQMServerPortRequest struct {
	// dwPortType:  Specifies the interface for which a port value is to be returned. One
	// of the following values MUST be specified; otherwise, this method MUST return 0x00000000
	// to indicate failure.
	//
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	|                          |                                                                                  |
	//	|          VALUE           |                                     MEANING                                      |
	//	|                          |                                                                                  |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| IP_HANDSHAKE 0x00000000  | Requests that the server return the RPC port number for the qmcomm and qmcomm2   |
	//	|                          | interfaces bound to TCP/IP. For more information on the qmcomm and qmcomm2       |
	//	|                          | interfaces, see [MS-MQMP]. The default port number is 2103.                      |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| IP_READ 0x00000001       | Requests that the server return the RPC port number for the qm2qm interface      |
	//	|                          | bound to TCP/IP. For more information on the qm2qm interface, see section 3.1.4. |
	//	|                          | The default port number is 2105.                                                 |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| IPX_HANDSHAKE 0x00000002 | Requests that the server return the RPC port number for the qmcomm and qmcomm2   |
	//	|                          | interfaces bound to SPX.<19> For more information on the qmcomm and qmcomm2      |
	//	|                          | interfaces, see [MS-MQMP]. The default port number is 2103.                      |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| IPX_READ 0x00000003      | Requests that the server return the RPC port number for the qm2qm interface      |
	//	|                          | bound to SPX.<20> For more information on the qm2qm interface, see section       |
	//	|                          | 3.1.4. The default port number is 2105.                                          |
	//	+--------------------------+----------------------------------------------------------------------------------+
	PortType uint32 `idl:"name:dwPortType" json:"port_type"`
}

func (o *GetQMQMServerPortRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQMQMServerPortOperation) *xxx_GetQMQMServerPortOperation {
	if op == nil {
		op = &xxx_GetQMQMServerPortOperation{}
	}
	if o == nil {
		return op
	}
	op.PortType = o.PortType
	return op
}

func (o *GetQMQMServerPortRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQMQMServerPortOperation) {
	if o == nil {
		return
	}
	o.PortType = op.PortType
}
func (o *GetQMQMServerPortRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQMQMServerPortRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQMQMServerPortOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQMQMServerPortResponse structure represents the RemoteQMGetQMQMServerPort operation response
type GetQMQMServerPortResponse struct {
	// Return: The RemoteQMGetQMQMServerPort return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetQMQMServerPortResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQMQMServerPortOperation) *xxx_GetQMQMServerPortOperation {
	if op == nil {
		op = &xxx_GetQMQMServerPortOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *GetQMQMServerPortResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQMQMServerPortOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *GetQMQMServerPortResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQMQMServerPortResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQMQMServerPortOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVersionOperation structure represents the RemoteQmGetVersion operation
type xxx_GetVersionOperation struct {
	Major       uint8  `idl:"name:pMajor" json:"major"`
	Minor       uint8  `idl:"name:pMinor" json:"minor"`
	BuildNumber uint16 `idl:"name:pBuildNumber" json:"build_number"`
}

func (o *xxx_GetVersionOperation) OpNum() int { return 8 }

func (o *xxx_GetVersionOperation) OpName() string { return "/qm2qm/v1/RemoteQmGetVersion" }

func (o *xxx_GetVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pMajor {out} (1:{pointer=ref}*(1)(uchar))
	{
		if err := w.WriteData(o.Major); err != nil {
			return err
		}
	}
	// pMinor {out} (1:{pointer=ref}*(1)(uchar))
	{
		if err := w.WriteData(o.Minor); err != nil {
			return err
		}
	}
	// pBuildNumber {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.BuildNumber); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pMajor {out} (1:{pointer=ref}*(1)(uchar))
	{
		if err := w.ReadData(&o.Major); err != nil {
			return err
		}
	}
	// pMinor {out} (1:{pointer=ref}*(1)(uchar))
	{
		if err := w.ReadData(&o.Minor); err != nil {
			return err
		}
	}
	// pBuildNumber {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.BuildNumber); err != nil {
			return err
		}
	}
	return nil
}

// GetVersionRequest structure represents the RemoteQmGetVersion operation request
type GetVersionRequest struct {
}

func (o *GetVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVersionOperation) *xxx_GetVersionOperation {
	if op == nil {
		op = &xxx_GetVersionOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *GetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
}
func (o *GetVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVersionResponse structure represents the RemoteQmGetVersion operation response
type GetVersionResponse struct {
	// pMajor: A pointer to an unsigned character. The server SHOULD<22> set this parameter
	// to 0x06.
	Major uint8 `idl:"name:pMajor" json:"major"`
	// pMinor: A pointer to an unsigned character. The server SHOULD<23> set this parameter
	// to 0x01.
	Minor uint8 `idl:"name:pMinor" json:"minor"`
	// pBuildNumber: A pointer to an unsigned short. The server SHOULD<24> set this parameter
	// to an implementation-specific build number.
	BuildNumber uint16 `idl:"name:pBuildNumber" json:"build_number"`
}

func (o *GetVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVersionOperation) *xxx_GetVersionOperation {
	if op == nil {
		op = &xxx_GetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Major = o.Major
	op.Minor = o.Minor
	op.BuildNumber = o.BuildNumber
	return op
}

func (o *GetVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.Major = op.Major
	o.Minor = op.Minor
	o.BuildNumber = op.BuildNumber
}
func (o *GetVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StartReceive2Operation structure represents the RemoteQMStartReceive2 operation
type xxx_StartReceive2Operation struct {
	Context         *Context         `idl:"name:pphContext" json:"context"`
	RemoteReadDesc2 *RemoteReadDesc2 `idl:"name:lpRemoteReadDesc2" json:"remote_read_desc2"`
	Return          int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_StartReceive2Operation) OpNum() int { return 9 }

func (o *xxx_StartReceive2Operation) OpName() string { return "/qm2qm/v1/RemoteQMStartReceive2" }

func (o *xxx_StartReceive2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceive2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpRemoteReadDesc2 {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC2}(struct))
	{
		if o.RemoteReadDesc2 != nil {
			if err := o.RemoteReadDesc2.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteReadDesc2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceive2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpRemoteReadDesc2 {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC2}(struct))
	{
		if o.RemoteReadDesc2 == nil {
			o.RemoteReadDesc2 = &RemoteReadDesc2{}
		}
		if err := o.RemoteReadDesc2.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceive2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceive2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_REMOTEREAD_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpRemoteReadDesc2 {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC2}(struct))
	{
		if o.RemoteReadDesc2 != nil {
			if err := o.RemoteReadDesc2.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteReadDesc2{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_StartReceive2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_REMOTEREAD_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpRemoteReadDesc2 {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC2}(struct))
	{
		if o.RemoteReadDesc2 == nil {
			o.RemoteReadDesc2 = &RemoteReadDesc2{}
		}
		if err := o.RemoteReadDesc2.UnmarshalNDR(ctx, w); err != nil {
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

// StartReceive2Request structure represents the RemoteQMStartReceive2 operation request
type StartReceive2Request struct {
	// lpRemoteReadDesc2: A pointer to an instance of a REMOTEREADDESC2 (section 2.2.2.2)
	// structure. The SequentialId member MUST be set to the least significant 7 bytes of
	// the Message.LookupIdentifier ([MS-MQDMPR] section 3.1.1.12) of the message that is
	// returned by this method.
	//
	// The client MUST provide all parameters of lpRemoteReadDesc2.pRemoteReadDesc that
	// are marked as to be set by the client in section 2.2.2.1.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|             VALUE OF              |                                                                                  |
	//	|             ULACTION              |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_RECEIVE 0x00000000      | If hCursor is nonzero, read and remove the first message available at the        |
	//	|                                   | current cursor location walking toward the end of the queue. If hCursor is zero, |
	//	|                                   | read and remove the message from the front of the queue.                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_CURRENT 0x80000000 | If hCursor is nonzero, read the message at the current cursor location, but      |
	//	|                                   | do not remove it from the queue. The cursor location does not change after the   |
	//	|                                   | operation. If hCursor is zero, read the message at the front of the queue, but   |
	//	|                                   | do not remove it from the queue.                                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_NEXT 0x80000001    | Read the message following the message at the current cursor location, but       |
	//	|                                   | do not remove it. The cursor location will then change to the next available     |
	//	|                                   | message, walking toward the end of the queue. The hCursor parameter MUST be set  |
	//	|                                   | to a nonzero cursor handle.                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//
	// The dwRequestID member of the lpRemoteReadDesc parameter is used in a subsequent
	// call to RemoteQMEndReceive or RemoteQMCancelReceive to correlate that call with the
	// call to RemoteQMStartReceive2.
	RemoteReadDesc2 *RemoteReadDesc2 `idl:"name:lpRemoteReadDesc2" json:"remote_read_desc2"`
}

func (o *StartReceive2Request) xxx_ToOp(ctx context.Context, op *xxx_StartReceive2Operation) *xxx_StartReceive2Operation {
	if op == nil {
		op = &xxx_StartReceive2Operation{}
	}
	if o == nil {
		return op
	}
	op.RemoteReadDesc2 = o.RemoteReadDesc2
	return op
}

func (o *StartReceive2Request) xxx_FromOp(ctx context.Context, op *xxx_StartReceive2Operation) {
	if o == nil {
		return
	}
	o.RemoteReadDesc2 = op.RemoteReadDesc2
}
func (o *StartReceive2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StartReceive2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartReceive2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StartReceive2Response structure represents the RemoteQMStartReceive2 operation response
type StartReceive2Response struct {
	// pphContext: The server MUST return a non-NULL value for this handle upon success
	// for receive calls. This handle will be used by the client in subsequent calls to
	// RemoteQMEndReceive (section 3.1.4.2). This handle MUST NOT be set upon failure or
	// for peek calls. If this method returns an error, pphContext is undefined and MUST
	// NOT be used as an argument for a call to RemoteQMEndReceive.
	Context *Context `idl:"name:pphContext" json:"context"`
	// lpRemoteReadDesc2: A pointer to an instance of a REMOTEREADDESC2 (section 2.2.2.2)
	// structure. The SequentialId member MUST be set to the least significant 7 bytes of
	// the Message.LookupIdentifier ([MS-MQDMPR] section 3.1.1.12) of the message that is
	// returned by this method.
	//
	// The client MUST provide all parameters of lpRemoteReadDesc2.pRemoteReadDesc that
	// are marked as to be set by the client in section 2.2.2.1.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|             VALUE OF              |                                                                                  |
	//	|             ULACTION              |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_RECEIVE 0x00000000      | If hCursor is nonzero, read and remove the first message available at the        |
	//	|                                   | current cursor location walking toward the end of the queue. If hCursor is zero, |
	//	|                                   | read and remove the message from the front of the queue.                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_CURRENT 0x80000000 | If hCursor is nonzero, read the message at the current cursor location, but      |
	//	|                                   | do not remove it from the queue. The cursor location does not change after the   |
	//	|                                   | operation. If hCursor is zero, read the message at the front of the queue, but   |
	//	|                                   | do not remove it from the queue.                                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_ACTION_PEEK_NEXT 0x80000001    | Read the message following the message at the current cursor location, but       |
	//	|                                   | do not remove it. The cursor location will then change to the next available     |
	//	|                                   | message, walking toward the end of the queue. The hCursor parameter MUST be set  |
	//	|                                   | to a nonzero cursor handle.                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//
	// The dwRequestID member of the lpRemoteReadDesc parameter is used in a subsequent
	// call to RemoteQMEndReceive or RemoteQMCancelReceive to correlate that call with the
	// call to RemoteQMStartReceive2.
	RemoteReadDesc2 *RemoteReadDesc2 `idl:"name:lpRemoteReadDesc2" json:"remote_read_desc2"`
	// Return: The RemoteQMStartReceive2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StartReceive2Response) xxx_ToOp(ctx context.Context, op *xxx_StartReceive2Operation) *xxx_StartReceive2Operation {
	if op == nil {
		op = &xxx_StartReceive2Operation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.RemoteReadDesc2 = o.RemoteReadDesc2
	op.Return = o.Return
	return op
}

func (o *StartReceive2Response) xxx_FromOp(ctx context.Context, op *xxx_StartReceive2Operation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.RemoteReadDesc2 = op.RemoteReadDesc2
	o.Return = op.Return
}
func (o *StartReceive2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StartReceive2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartReceive2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StartReceiveByLookupIDOperation structure represents the RemoteQMStartReceiveByLookupId operation
type xxx_StartReceiveByLookupIDOperation struct {
	LookupID        uint64           `idl:"name:LookupId" json:"lookup_id"`
	Context         *Context         `idl:"name:pphContext" json:"context"`
	RemoteReadDesc2 *RemoteReadDesc2 `idl:"name:lpRemoteReadDesc2" json:"remote_read_desc2"`
	Return          int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_StartReceiveByLookupIDOperation) OpNum() int { return 10 }

func (o *xxx_StartReceiveByLookupIDOperation) OpName() string {
	return "/qm2qm/v1/RemoteQMStartReceiveByLookupId"
}

func (o *xxx_StartReceiveByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LookupId {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.LookupID); err != nil {
			return err
		}
	}
	// lpRemoteReadDesc2 {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC2}(struct))
	{
		if o.RemoteReadDesc2 != nil {
			if err := o.RemoteReadDesc2.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteReadDesc2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LookupId {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.LookupID); err != nil {
			return err
		}
	}
	// lpRemoteReadDesc2 {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC2}(struct))
	{
		if o.RemoteReadDesc2 == nil {
			o.RemoteReadDesc2 = &RemoteReadDesc2{}
		}
		if err := o.RemoteReadDesc2.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartReceiveByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_REMOTEREAD_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpRemoteReadDesc2 {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC2}(struct))
	{
		if o.RemoteReadDesc2 != nil {
			if err := o.RemoteReadDesc2.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteReadDesc2{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_StartReceiveByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_REMOTEREAD_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpRemoteReadDesc2 {in, out} (1:{pointer=ref}*(1))(2:{alias=REMOTEREADDESC2}(struct))
	{
		if o.RemoteReadDesc2 == nil {
			o.RemoteReadDesc2 = &RemoteReadDesc2{}
		}
		if err := o.RemoteReadDesc2.UnmarshalNDR(ctx, w); err != nil {
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

// StartReceiveByLookupIDRequest structure represents the RemoteQMStartReceiveByLookupId operation request
type StartReceiveByLookupIDRequest struct {
	// LookupId: Lookup identifier of the message to be returned.
	LookupID uint64 `idl:"name:LookupId" json:"lookup_id"`
	// lpRemoteReadDesc2: A REMOTEREADDESC2 (section 2.2.2.2) instance that contains the
	// remote description accompanied by a sequential ID. The members of the pRemoteReadDesc
	// member of the lpRemoteReadDesc2 parameter MUST be assigned in the same manner as
	// that specified in RemoteQMStartReceive and section 2.2.2.1. In addition, the SequentialId
	// member MUST be set to the least significant 7 bytes of the Message.LookupIdentifier
	// ([MS-MQDMPR] section 3.1.1.12) of the message that is returned by this method.
	//
	// The client must provide all parameters of lpRemoteReadDesc2.pRemoteReadDesc that
	// are marked as to be set by the client in section 2.2.2.1.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|               VALUE OF               |                                                                                  |
	//	|               ULACTION               |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_CURRENT 0x40000010    | Read the message that is specified by the LookupId parameter, but do not         |
	//	|                                      | remove it from the queue. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor          |
	//	|                                      | parameter MUST be set to zero. The LookupId parameter MUST NOT be set to 0. The  |
	//	|                                      | lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set to 0x00000000. |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_NEXT 0x40000011       | If LookupId is 0, read the first message. Otherwise, read the message following  |
	//	|                                      | the message that is specified by LookupId. In either case, do not remove the     |
	//	|                                      | message. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor parameter MUST be set to  |
	//	|                                      | zero. The lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set to   |
	//	|                                      | 0x00000000.                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_PREV 0x40000012       | If LookupId is 0xFFFFFFFFFFFFFFFF, read the last message. Otherwise,             |
	//	|                                      | read the message preceding the message that is specified by the                  |
	//	|                                      | LookupId parameter. In either case, do not remove the message from               |
	//	|                                      | the queue. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor parameter               |
	//	|                                      | MUST be set to zero. The LookupId parameter MUST NOT be set to 0. The            |
	//	|                                      | lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set to 0x00000000. |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_CURRENT 0x40000020 | Read the message that is specified by the LookupId parameter, and remove         |
	//	|                                      | it from the queue. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor parameter       |
	//	|                                      | MUST be set to zero. The LookupId parameter MUST NOT be set to 0. The            |
	//	|                                      | lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set to 0x00000000. |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_NEXT 0x40000021    | If LookupId is 0, read the first message. Otherwise, read the message following  |
	//	|                                      | the message that is specified by the LookupId parameter. Remove the message from |
	//	|                                      | the queue. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor parameter MUST be set   |
	//	|                                      | to zero. The lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set   |
	//	|                                      | to 0x00000000.                                                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_PREV 0x40000022    | If LookupId is 0xFFFFFFFFFFFFFFFF, read the last message. Otherwise, read the    |
	//	|                                      | message preceding the message that is specified by the LookupId parameter.       |
	//	|                                      | Remove the message from the queue. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor |
	//	|                                      | parameter MUST be set to zero. The LookupId parameter MUST NOT be set to 0. The  |
	//	|                                      | lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set to 0x00000000. |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	RemoteReadDesc2 *RemoteReadDesc2 `idl:"name:lpRemoteReadDesc2" json:"remote_read_desc2"`
}

func (o *StartReceiveByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_StartReceiveByLookupIDOperation) *xxx_StartReceiveByLookupIDOperation {
	if op == nil {
		op = &xxx_StartReceiveByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.LookupID = o.LookupID
	op.RemoteReadDesc2 = o.RemoteReadDesc2
	return op
}

func (o *StartReceiveByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_StartReceiveByLookupIDOperation) {
	if o == nil {
		return
	}
	o.LookupID = op.LookupID
	o.RemoteReadDesc2 = op.RemoteReadDesc2
}
func (o *StartReceiveByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StartReceiveByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartReceiveByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StartReceiveByLookupIDResponse structure represents the RemoteQMStartReceiveByLookupId operation response
type StartReceiveByLookupIDResponse struct {
	// pphContext: The server MUST return a non-NULL value for this handle, on success for
	// receive calls. This handle is used by the client in subsequent calls to RemoteQMEndReceive.
	// This handle MUST NOT be set on failure, or for peek calls. If this method returns
	// an error, pphContext is undefined and MUST NOT be used as an argument for a call
	// to RemoteQMEndReceive.
	Context *Context `idl:"name:pphContext" json:"context"`
	// lpRemoteReadDesc2: A REMOTEREADDESC2 (section 2.2.2.2) instance that contains the
	// remote description accompanied by a sequential ID. The members of the pRemoteReadDesc
	// member of the lpRemoteReadDesc2 parameter MUST be assigned in the same manner as
	// that specified in RemoteQMStartReceive and section 2.2.2.1. In addition, the SequentialId
	// member MUST be set to the least significant 7 bytes of the Message.LookupIdentifier
	// ([MS-MQDMPR] section 3.1.1.12) of the message that is returned by this method.
	//
	// The client must provide all parameters of lpRemoteReadDesc2.pRemoteReadDesc that
	// are marked as to be set by the client in section 2.2.2.1.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|               VALUE OF               |                                                                                  |
	//	|               ULACTION               |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_CURRENT 0x40000010    | Read the message that is specified by the LookupId parameter, but do not         |
	//	|                                      | remove it from the queue. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor          |
	//	|                                      | parameter MUST be set to zero. The LookupId parameter MUST NOT be set to 0. The  |
	//	|                                      | lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set to 0x00000000. |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_NEXT 0x40000011       | If LookupId is 0, read the first message. Otherwise, read the message following  |
	//	|                                      | the message that is specified by LookupId. In either case, do not remove the     |
	//	|                                      | message. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor parameter MUST be set to  |
	//	|                                      | zero. The lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set to   |
	//	|                                      | 0x00000000.                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_PEEK_PREV 0x40000012       | If LookupId is 0xFFFFFFFFFFFFFFFF, read the last message. Otherwise,             |
	//	|                                      | read the message preceding the message that is specified by the                  |
	//	|                                      | LookupId parameter. In either case, do not remove the message from               |
	//	|                                      | the queue. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor parameter               |
	//	|                                      | MUST be set to zero. The LookupId parameter MUST NOT be set to 0. The            |
	//	|                                      | lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set to 0x00000000. |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_CURRENT 0x40000020 | Read the message that is specified by the LookupId parameter, and remove         |
	//	|                                      | it from the queue. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor parameter       |
	//	|                                      | MUST be set to zero. The LookupId parameter MUST NOT be set to 0. The            |
	//	|                                      | lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set to 0x00000000. |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_NEXT 0x40000021    | If LookupId is 0, read the first message. Otherwise, read the message following  |
	//	|                                      | the message that is specified by the LookupId parameter. Remove the message from |
	//	|                                      | the queue. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor parameter MUST be set   |
	//	|                                      | to zero. The lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set   |
	//	|                                      | to 0x00000000.                                                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_LOOKUP_RECEIVE_PREV 0x40000022    | If LookupId is 0xFFFFFFFFFFFFFFFF, read the last message. Otherwise, read the    |
	//	|                                      | message preceding the message that is specified by the LookupId parameter.       |
	//	|                                      | Remove the message from the queue. The lpRemoteReadDesc2.pRemoteReadDesc.hCursor |
	//	|                                      | parameter MUST be set to zero. The LookupId parameter MUST NOT be set to 0. The  |
	//	|                                      | lpRemoteReadDesc2.pRemoteReadDesc.ulTimeout parameter MUST be set to 0x00000000. |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	RemoteReadDesc2 *RemoteReadDesc2 `idl:"name:lpRemoteReadDesc2" json:"remote_read_desc2"`
	// Return: The RemoteQMStartReceiveByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StartReceiveByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_StartReceiveByLookupIDOperation) *xxx_StartReceiveByLookupIDOperation {
	if op == nil {
		op = &xxx_StartReceiveByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.RemoteReadDesc2 = o.RemoteReadDesc2
	op.Return = o.Return
	return op
}

func (o *StartReceiveByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_StartReceiveByLookupIDOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.RemoteReadDesc2 = op.RemoteReadDesc2
	o.Return = op.Return
}
func (o *StartReceiveByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StartReceiveByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartReceiveByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
