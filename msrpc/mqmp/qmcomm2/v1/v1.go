package qmcomm2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	qmcomm "github.com/oiweiwei/go-msrpc/msrpc/mqmp/qmcomm/v1"
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
	_ = mqmq.GoPackage
	_ = qmcomm.GoPackage
)

var (
	// import guard
	GoPackage = "mqmp"
)

var (
	// Syntax UUID
	Qmcomm2SyntaxUUID = &uuid.UUID{TimeLow: 0x76d12b80, TimeMid: 0x3467, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0xff, Node: [6]uint8{0x0, 0x90, 0x27, 0x2f, 0x9e, 0xa3}}
	// Syntax ID
	Qmcomm2SyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: Qmcomm2SyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// qmcomm2 interface.
type Qmcomm2Client interface {

	// A client invokes QMSendMessageInternalEx if the server returns STATUS_RETRY (0xc000022d)
	// from a prior call to rpc_ACSendMessageEx. Implementations of this protocol SHOULD
	// NOT return STATUS_RETRY from rpc_ACSendMessageEx, rendering this method unnecessary.
	// Such implementations MUST take no action when QMSendMessageInternalEx is invoked
	// and return MQ_ERROR_ILLEGAL_OPERATION (0xc00e0064).
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<67> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	QMSendMessageInternalEx(context.Context, *QMSendMessageInternalExRequest, ...dcerpc.CallOption) (*QMSendMessageInternalExResponse, error)

	// A client calls the rpc_ACSendMessageEx method to place a message into a message queue
	// for delivery.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<72> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	//
	// Security Considerations: The caller can request that the server perform security
	// related operations such as signing and encrypting the message. These operations are
	// requested by setting members of the ptb.CACTransferBufferOld structure.
	SendMessageEx(context.Context, *SendMessageExRequest, ...dcerpc.CallOption) (*SendMessageExResponse, error)

	// A client calls rpc_ACReceiveMessageEx to peek or receive a message from a message
	// queue.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<73><74> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	ReceiveMessageEx(context.Context, *ReceiveMessageExRequest, ...dcerpc.CallOption) (*ReceiveMessageExResponse, error)

	// A client calls rpc_ACCreateCursorEx to create a cursor for use when peeking and receiving
	// from a message queue.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<75> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	CreateCursorEx(context.Context, *CreateCursorExRequest, ...dcerpc.CallOption) (*CreateCursorExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultQmcomm2Client struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultQmcomm2Client) QMSendMessageInternalEx(ctx context.Context, in *QMSendMessageInternalExRequest, opts ...dcerpc.CallOption) (*QMSendMessageInternalExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QMSendMessageInternalExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcomm2Client) SendMessageEx(ctx context.Context, in *SendMessageExRequest, opts ...dcerpc.CallOption) (*SendMessageExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SendMessageExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcomm2Client) ReceiveMessageEx(ctx context.Context, in *ReceiveMessageExRequest, opts ...dcerpc.CallOption) (*ReceiveMessageExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReceiveMessageExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcomm2Client) CreateCursorEx(ctx context.Context, in *CreateCursorExRequest, opts ...dcerpc.CallOption) (*CreateCursorExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateCursorExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcomm2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQmcomm2Client) Conn() dcerpc.Conn {
	return o.cc
}

func NewQmcomm2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Qmcomm2Client, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Qmcomm2SyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultQmcomm2Client{cc: cc}, nil
}

// xxx_QMSendMessageInternalExOperation structure represents the QMSendMessageInternalEx operation
type xxx_QMSendMessageInternalExOperation struct {
	QueueFormat *mqmq.QueueFormat        `idl:"name:pQueueFormat" json:"queue_format"`
	Ptb         *qmcomm.TransferBufferV2 `idl:"name:ptb" json:"ptb"`
	MessageID   *mqmq.ObjectID           `idl:"name:pMessageID;pointer:unique" json:"message_id"`
	Return      int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_QMSendMessageInternalExOperation) OpNum() int { return 0 }

func (o *xxx_QMSendMessageInternalExOperation) OpName() string {
	return "/qmcomm2/v1/QMSendMessageInternalEx"
}

func (o *xxx_QMSendMessageInternalExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMSendMessageInternalExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ptb {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.Ptb != nil {
			if err := o.Ptb.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&qmcomm.TransferBufferV2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pMessageID {in, out} (1:{pointer=unique}*(1))(2:{alias=OBJECTID}(struct))
	{
		if o.MessageID != nil {
			_ptr_pMessageID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MessageID != nil {
					if err := o.MessageID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqmq.ObjectID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MessageID, _ptr_pMessageID); err != nil {
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
	return nil
}

func (o *xxx_QMSendMessageInternalExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ptb {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.Ptb == nil {
			o.Ptb = &qmcomm.TransferBufferV2{}
		}
		if err := o.Ptb.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pMessageID {in, out} (1:{pointer=unique}*(1))(2:{alias=OBJECTID}(struct))
	{
		_ptr_pMessageID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageID == nil {
				o.MessageID = &mqmq.ObjectID{}
			}
			if err := o.MessageID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pMessageID := func(ptr interface{}) { o.MessageID = *ptr.(**mqmq.ObjectID) }
		if err := w.ReadPointer(&o.MessageID, _s_pMessageID, _ptr_pMessageID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMSendMessageInternalExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMSendMessageInternalExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pMessageID {in, out} (1:{pointer=unique}*(1))(2:{alias=OBJECTID}(struct))
	{
		if o.MessageID != nil {
			_ptr_pMessageID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MessageID != nil {
					if err := o.MessageID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqmq.ObjectID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MessageID, _ptr_pMessageID); err != nil {
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

func (o *xxx_QMSendMessageInternalExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pMessageID {in, out} (1:{pointer=unique}*(1))(2:{alias=OBJECTID}(struct))
	{
		_ptr_pMessageID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageID == nil {
				o.MessageID = &mqmq.ObjectID{}
			}
			if err := o.MessageID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pMessageID := func(ptr interface{}) { o.MessageID = *ptr.(**mqmq.ObjectID) }
		if err := w.ReadPointer(&o.MessageID, _s_pMessageID, _ptr_pMessageID); err != nil {
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

// QMSendMessageInternalExRequest structure represents the QMSendMessageInternalEx operation request
type QMSendMessageInternalExRequest struct {
	// pQueueFormat: MUST be a pointer to a QUEUE_FORMAT ([MS-MQMQ] section 2.2.7) structure,
	// which identifies an existing queue to be opened. MUST NOT be NULL, and MUST conform
	// to the format name syntax rules defined in [MS-MQMQ]. The queue identified by pQueueFormat
	// MUST be local to the supporting server, and MUST be successfully openable via a call
	// to rpc_QMOpenQueueInternal with a dwDesiredAccess level of MQ_SEND_ACCESS (0x00000002).
	QueueFormat *mqmq.QueueFormat `idl:"name:pQueueFormat" json:"queue_format"`
	// ptb: A CACTransferBufferV2 structure pointer as described in section 2.2.3.3. See
	// the identical parameter in section 3.1.5.2 for details on this parameter.
	Ptb *qmcomm.TransferBufferV2 `idl:"name:ptb" json:"ptb"`
	// pMessageID: An OBJECTID as defined in [MS-MQMQ] section 2.2.8. See the identical
	// parameter in section 3.1.5.2 for details on this parameter.
	MessageID *mqmq.ObjectID `idl:"name:pMessageID;pointer:unique" json:"message_id"`
}

func (o *QMSendMessageInternalExRequest) xxx_ToOp(ctx context.Context) *xxx_QMSendMessageInternalExOperation {
	if o == nil {
		return &xxx_QMSendMessageInternalExOperation{}
	}
	return &xxx_QMSendMessageInternalExOperation{
		QueueFormat: o.QueueFormat,
		Ptb:         o.Ptb,
		MessageID:   o.MessageID,
	}
}

func (o *QMSendMessageInternalExRequest) xxx_FromOp(ctx context.Context, op *xxx_QMSendMessageInternalExOperation) {
	if o == nil {
		return
	}
	o.QueueFormat = op.QueueFormat
	o.Ptb = op.Ptb
	o.MessageID = op.MessageID
}
func (o *QMSendMessageInternalExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QMSendMessageInternalExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QMSendMessageInternalExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QMSendMessageInternalExResponse structure represents the QMSendMessageInternalEx operation response
type QMSendMessageInternalExResponse struct {
	// pMessageID: An OBJECTID as defined in [MS-MQMQ] section 2.2.8. See the identical
	// parameter in section 3.1.5.2 for details on this parameter.
	MessageID *mqmq.ObjectID `idl:"name:pMessageID;pointer:unique" json:"message_id"`
	// Return: The QMSendMessageInternalEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QMSendMessageInternalExResponse) xxx_ToOp(ctx context.Context) *xxx_QMSendMessageInternalExOperation {
	if o == nil {
		return &xxx_QMSendMessageInternalExOperation{}
	}
	return &xxx_QMSendMessageInternalExOperation{
		MessageID: o.MessageID,
		Return:    o.Return,
	}
}

func (o *QMSendMessageInternalExResponse) xxx_FromOp(ctx context.Context, op *xxx_QMSendMessageInternalExOperation) {
	if o == nil {
		return
	}
	o.MessageID = op.MessageID
	o.Return = op.Return
}
func (o *QMSendMessageInternalExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QMSendMessageInternalExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QMSendMessageInternalExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SendMessageExOperation structure represents the rpc_ACSendMessageEx operation
type xxx_SendMessageExOperation struct {
	Queue     *qmcomm.Queue            `idl:"name:hQueue" json:"queue"`
	Ptb       *qmcomm.TransferBufferV2 `idl:"name:ptb" json:"ptb"`
	MessageID *mqmq.ObjectID           `idl:"name:pMessageID;pointer:unique" json:"message_id"`
	Return    int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_SendMessageExOperation) OpNum() int { return 1 }

func (o *xxx_SendMessageExOperation) OpName() string { return "/qmcomm2/v1/rpc_ACSendMessageEx" }

func (o *xxx_SendMessageExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hQueue {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue != nil {
			if err := o.Queue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&qmcomm.Queue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ptb {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.Ptb != nil {
			if err := o.Ptb.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&qmcomm.TransferBufferV2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pMessageID {in, out} (1:{pointer=unique}*(1))(2:{alias=OBJECTID}(struct))
	{
		if o.MessageID != nil {
			_ptr_pMessageID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MessageID != nil {
					if err := o.MessageID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqmq.ObjectID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MessageID, _ptr_pMessageID); err != nil {
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
	return nil
}

func (o *xxx_SendMessageExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQueue {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue == nil {
			o.Queue = &qmcomm.Queue{}
		}
		if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ptb {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.Ptb == nil {
			o.Ptb = &qmcomm.TransferBufferV2{}
		}
		if err := o.Ptb.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pMessageID {in, out} (1:{pointer=unique}*(1))(2:{alias=OBJECTID}(struct))
	{
		_ptr_pMessageID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageID == nil {
				o.MessageID = &mqmq.ObjectID{}
			}
			if err := o.MessageID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pMessageID := func(ptr interface{}) { o.MessageID = *ptr.(**mqmq.ObjectID) }
		if err := w.ReadPointer(&o.MessageID, _s_pMessageID, _ptr_pMessageID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pMessageID {in, out} (1:{pointer=unique}*(1))(2:{alias=OBJECTID}(struct))
	{
		if o.MessageID != nil {
			_ptr_pMessageID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MessageID != nil {
					if err := o.MessageID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqmq.ObjectID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MessageID, _ptr_pMessageID); err != nil {
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

func (o *xxx_SendMessageExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pMessageID {in, out} (1:{pointer=unique}*(1))(2:{alias=OBJECTID}(struct))
	{
		_ptr_pMessageID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageID == nil {
				o.MessageID = &mqmq.ObjectID{}
			}
			if err := o.MessageID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pMessageID := func(ptr interface{}) { o.MessageID = *ptr.(**mqmq.ObjectID) }
		if err := w.ReadPointer(&o.MessageID, _s_pMessageID, _ptr_pMessageID); err != nil {
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

// SendMessageExRequest structure represents the rpc_ACSendMessageEx operation request
type SendMessageExRequest struct {
	// hQueue:  MUST be an RPC_QUEUE_HANDLE (section 2.2.1.1.2) obtained from the phQueue
	// parameter of the rpc_QMOpenQueueInternal (section 3.1.4.17) method called with the
	// dwDesiredAccess parameter set to MQ_SEND_ACCESS. Prior to this method being invoked,
	// the queue MUST NOT have been deleted, and the queue handle MUST NOT have been closed.
	Queue *qmcomm.Queue `idl:"name:hQueue" json:"queue"`
	// ptb: MUST NOT be NULL. ptb points to a CACTransferBufferV2 (section 2.2.3.3) structure.
	// Refer to section 2.2.3.3 for definitions of the CACTransferBufferV2 members. Constraints
	// for the members are defined following. In the section following, "ptb.old" is used
	// as shorthand to refer to the CACTransferBufferOld member of the CACTransferBufferV2
	// structure.
	//
	// ptb.old.uTransferType MUST be CACTB_SEND (0x00000000).
	//
	// ptb.old.Send.pAdminQueueFormat can be NULL, in which case no administration queue
	// format name is associated with the message. If not NULL, ptb.old.Send.pAdminQueueFormat
	// MUST point to a QUEUE_FORMAT ([MS-MQMQ] section 2.2.7) structure.
	//
	// ptb.old.Send.pResponseQueueFormat can be NULL, in which case no response queue format
	// name is associated with the message. If not NULL, ptb.old.Send.pResponseQueueFormat
	// MUST point to a QUEUE_FORMAT structure.
	//
	// If the queue identified hQueue was opened using a direct format name, as specified
	// in [MS-MQMQ] section 2.1.2, ptb.old.pulPrivLevel MUST be NULL or, if not NULL, MUST
	// point to the value MQMSG_PRIV_LEVEL_NONE (0x00000000). Encryption MUST NOT be requested
	// for queues opened with direct format name.
	//
	// If the queue identified by hQueue is not an outgoing queue (rather, it is a queue
	// which is local to the supporting server), and ptb.bEncrypted is not 0x00, the server
	// MAY return STATUS_RETRY (0xc000022d) and take no action.<68>
	//
	// ptb.old.pPriority can be NULL; otherwise, the value MUST be from 0x00 to 0x07 inclusive.
	// If the value is NULL, the server MUST substitute the default value of 0x03.
	//
	// ptb.old.pTrace can be NULL, in which case the server MUST substitute the default
	// value of 0x00.
	//
	// If ptb.old.ulAbsoluteTimeToQueue is 0x00000000, the server MUST substitute the default
	// value of 0xffffffff.
	//
	// ptb.old.ppMessageID can be NULL. If not NULL, the server MUST ignore the in-value.
	//
	// ptb.old.ppConnectorType can be NULL. If NULL, then no connector type value is associated
	// with the message.
	//
	// ptb.old.pDelivery can be NULL, in which case the server MUST substitute the default
	// value of 0x00. However, if ptb.old.pUow contains a nonzero value, the server MUST
	// substitute the value 0x01 for ptb.old.pDelivery, since transactional messages are
	// by definition stored as recoverable.
	//
	// ptb.old.pAuditing can be NULL, in which case the server MUST substitute the default
	// value of 0x00.
	//
	// ptb.old.pClass can be NULL, in which case the server MUST substitute the default
	// value of 0x0000. This field can be used by connector applications to produce acknowledgment
	// messages. Typical applications will always specify MQMSG_CLASS_NORMAL (0x0000).
	//
	// ptb.old.ppCorrelationID can be NULL, in which case the server MUST substitute the
	// default value by filling the array of bytes with hexadecimal zeros (0x00).
	//
	// ptb.old.pAcknowledge can be NULL, in which case the server MUST substitute the default
	// value of 0x00.
	//
	// ptb.old.pApplicationTag can be NULL, in which case the server MUST substitute the
	// default value of 0x00000000.
	//
	// ptb.old.ppTitle can be NULL, in which case the server MUST treat the value as an
	// empty string and MUST ignore the value of ptb.old.ulTitleBufferSizeInWCHARs. If ptb.old.ppTitle
	// is NOT NULL, the server MUST take the number of Unicode characters indicated by ptb.old.ulTitleBufferSizeInWCHARs.
	// If ptb.old.ulTitleBufferSizeInWCHARs is greater than 250, the value MUST be truncated
	// to 250. The server MUST null-terminate the resulting character array.
	//
	// ptb.old.ppMsgExtension can be NULL, in which case no extension array is associated
	// with the message and the server MUST ignore the value of ptb.old.ulMsgExtensionBufferInBytes.
	// If ptb.old.ppMsgExtension is NOT NULL, the server MUST take the number of bytes indicated
	// by ptb.old.ulMsgExtensionBufferInBytes. The buffer is an opaque array of bytes and
	// a terminating null character is not required.
	//
	// ptb.old.ppBody can be NULL, in which case no body array is associated with the message
	// and the server MUST ignore the values of ptb.old.ulBodyBufferSizeInBytes and ptb.old.ulAllocBodyBufferInBytes.
	// If ptb.old.ppBody is NOT NULL, the server MUST take the number of bytes indicated
	// by ptb.old.ulBodyBufferSizeInBytes, and allocate body storage for the number of bytes
	// indicated by ptb.old.ulAllocBodyBufferInBytes. The message body is an opaque array
	// of bytes and a terminating null character is not required.
	//
	// ptb.old.pulPrivLevel can be NULL, in which case the server MUST substitute the default
	// value of 0x00000000.
	//
	// ptb.old.pulHashAlg can be NULL if ptb.old.ulSignatureSize is 0x00000000; otherwise,
	// it MUST be set to the hash algorithm used to produce the signature of the message,
	// as specified in section 2.2.3.2.<69> If it is set to NULL, the server MUST substitute
	// the value of 0x00000000.
	//
	// ptb.old.pulEncryptAlg can be NULL if ptb.old.pulPrivLevel is set to NULL; otherwise,
	// it MUST be set to the encryption algorithm associated with ptb.old.pulPrivLevel,
	// as specified in section 2.2.3.2.<70> If it is set to NULL, the server MUST substitute
	// the value of 0x00000000.
	//
	// ptb.old.pulBodyType can be NULL, in which case the server MUST substitute the default
	// value of 0x00000000.
	//
	// ptb.old.ppSenderCert can be NULL if ptb.old.ulSenderCertLen is 0x00000000, in which
	// case an X509 certificate for the sender is not associated with the message.
	//
	// ptb.old.pulSenderIDType MUST NOT be NULL if ptb.old.uSenderIDLen is nonzero.
	//
	// ptb.old.pSenderID can be NULL if ptb.old.uSenderIDLen is zero and ptb.old.pulSenderIDType
	// is MQMSG_SENDERID_TYPE_NONE (0x00000000), in which case a SID is not associated with
	// the message.
	//
	// ptb.old.ppSymmKeys can be NULL if ptb.old.ulSymmKeysSize is zero (0x00000000), in
	// which case an encrypted symmetric key is not associated with the message. Otherwise,
	// ptb.old.ppSymKeys MUST contain the symmetric key used to encrypt the message body.
	// The symmetric key MUST be encrypted with the public key of the recipient QM. The
	// manner by which the public key for the recipient QM is obtained is beyond the scope
	// of this network protocol.
	//
	// If ptb.old.ulSignatureSize is 0x00000000: no digital signature is associated with
	// the message.
	//
	// * If ptb.old.fDefaultProvider is 0x00000000, ptb.old.ppwcsProvName MUST NOT be NULL.
	// If ptb.old.pulProvType is NOT NULL, it MUST specify the provider type of the CSP
	// named by ptb.old.ppwcsProvName; otherwise, the server MUST substitute the value of
	// 0x00000000. Note that ptb.old.ulProvNameLen is used only to affect RPC ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_8a7f6700-8311-45bc-af10-82e10accd331
	// ) marshaling of the ptb.old.ppwcsProvName buffer. The server MUST otherwise ignore
	// ptb.old.ulProvNameLen and treat ptb.old.ppwcsProvName as a null-terminated string.
	//
	// * If ptb.old.fDefaultProvider is not 0x00000000, ptb.old.pulProvType MUST NOT be
	// NULL, and MUST contain PROV_RSA_FULL (0x00000001).
	//
	// * If ptb.old.pulAuthProvNameLenProp is NULL:
	//
	// * If not NULL, the ptb.old.ppSignature buffer contains a simple array of bytes containing
	// the MSMQ 1.0 digital signature ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_cf9a8e0f-8060-464b-a673-fe4f815d3d8a
	// ). The byte length of the buffer is indicated by ptb.old.ulSignatureSize.
	//
	// * Else, if ptb.old.pulAuthProvNameLenProp is NOT NULL:
	//
	// * If not NULL, the ptb.old.ppSignature buffer contains two distinct byte array parts.
	// The first part MUST be ignored by the server. The second part contains an MSMQ 2.0
	// digital signature ( 102555d1-0dbf-4b2e-b78c-e388823d252c#gt_46782277-690b-4a09-ad6d-8c36b11e51cd
	// ).
	//
	// * The byte length of the first part is indicated by subtracting the length of the
	// second part from ptb.old.ulSignatureSize. (Thus, length( [first part] ) + length(
	// [second part] ) = ptb.old.ulSignatureSize.)
	//
	// * The byte length of the second part is indicated by subtracting ptb.old.ulProvNameLen
	// from ptb.old.pulAuthProvNameLenProp.
	//
	// * The second part begins immediately after the first.
	//
	// * ptb.old.Receive
	//
	// * ptb.old.CreateCursor
	//
	// * ptb.old.pSentTime
	//
	// * ptb.old.pArrivedTime
	//
	// * ptb.old.pBodySize
	//
	// * ptb.old.pulTitleBufferSizeInWCHARs
	//
	// * ptb.old.pulRelativeTimeToQueue
	//
	// * ptb.old.pulRelativeTimeToLive
	//
	// * ptb.old.pulSenderIDLenProp
	//
	// * ptb.old.ulAuthLevel
	//
	// * ptb.old.pAuthenticated
	//
	// * ptb.old.bAuthenticated
	//
	// * ptb.old.pulSenderCertLenProp
	//
	// * ptb.old.pulSymmKeysSizeProp
	//
	// * ptb.old.pulSignatureSizeProp
	//
	// * ptb.old.ppSrcQMID
	//
	// * ptb.old.pMsgExtensionSize
	//
	// * ptb.old.pulVersion
	//
	// * ptb.pbFirstInXact
	//
	// * ptb.pbLastInXact
	//
	// * ptb.ppXactID
	Ptb *qmcomm.TransferBufferV2 `idl:"name:ptb" json:"ptb"`
	// pMessageID: An OBJECTID as defined in [MS-MQMQ] section 2.2.8. This value can be
	// NULL. If not NULL, the server MUST return a message identifier for the new message
	// if this method succeeds.
	MessageID *mqmq.ObjectID `idl:"name:pMessageID;pointer:unique" json:"message_id"`
}

func (o *SendMessageExRequest) xxx_ToOp(ctx context.Context) *xxx_SendMessageExOperation {
	if o == nil {
		return &xxx_SendMessageExOperation{}
	}
	return &xxx_SendMessageExOperation{
		Queue:     o.Queue,
		Ptb:       o.Ptb,
		MessageID: o.MessageID,
	}
}

func (o *SendMessageExRequest) xxx_FromOp(ctx context.Context, op *xxx_SendMessageExOperation) {
	if o == nil {
		return
	}
	o.Queue = op.Queue
	o.Ptb = op.Ptb
	o.MessageID = op.MessageID
}
func (o *SendMessageExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SendMessageExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendMessageExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SendMessageExResponse structure represents the rpc_ACSendMessageEx operation response
type SendMessageExResponse struct {
	// pMessageID: An OBJECTID as defined in [MS-MQMQ] section 2.2.8. This value can be
	// NULL. If not NULL, the server MUST return a message identifier for the new message
	// if this method succeeds.
	MessageID *mqmq.ObjectID `idl:"name:pMessageID;pointer:unique" json:"message_id"`
	// Return: The rpc_ACSendMessageEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SendMessageExResponse) xxx_ToOp(ctx context.Context) *xxx_SendMessageExOperation {
	if o == nil {
		return &xxx_SendMessageExOperation{}
	}
	return &xxx_SendMessageExOperation{
		MessageID: o.MessageID,
		Return:    o.Return,
	}
}

func (o *SendMessageExResponse) xxx_FromOp(ctx context.Context, op *xxx_SendMessageExOperation) {
	if o == nil {
		return
	}
	o.MessageID = op.MessageID
	o.Return = op.Return
}
func (o *SendMessageExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SendMessageExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendMessageExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveMessageExOperation structure represents the rpc_ACReceiveMessageEx operation
type xxx_ReceiveMessageExOperation struct {
	HQMContext uint32                   `idl:"name:hQMContext" json:"h_qm_context"`
	Ptb        *qmcomm.TransferBufferV2 `idl:"name:ptb" json:"ptb"`
	Return     int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveMessageExOperation) OpNum() int { return 2 }

func (o *xxx_ReceiveMessageExOperation) OpName() string { return "/qmcomm2/v1/rpc_ACReceiveMessageEx" }

func (o *xxx_ReceiveMessageExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveMessageExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hQMContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HQMContext); err != nil {
			return err
		}
	}
	// ptb {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.Ptb != nil {
			if err := o.Ptb.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&qmcomm.TransferBufferV2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveMessageExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQMContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HQMContext); err != nil {
			return err
		}
	}
	// ptb {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.Ptb == nil {
			o.Ptb = &qmcomm.TransferBufferV2{}
		}
		if err := o.Ptb.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveMessageExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveMessageExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ptb {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.Ptb != nil {
			if err := o.Ptb.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&qmcomm.TransferBufferV2{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ReceiveMessageExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ptb {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.Ptb == nil {
			o.Ptb = &qmcomm.TransferBufferV2{}
		}
		if err := o.Ptb.UnmarshalNDR(ctx, w); err != nil {
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

// ReceiveMessageExRequest structure represents the rpc_ACReceiveMessageEx operation request
type ReceiveMessageExRequest struct {
	// hQMContext: A queue context value obtained from the pdwQMContext parameter of rpc_QMOpenQueueInternal.
	// The queue MUST have been opened with MQ_PEEK_ACCESS or MQ_RECEIVE_ACCESS as the dwDesiredAccess
	// parameter when rpc_QMOpenQueueInternal was called. Prior to this method being invoked,
	// the queue MUST NOT have been deleted, and the queue handle for the open queue MUST
	// NOT have been closed.
	HQMContext uint32 `idl:"name:hQMContext" json:"h_qm_context"`
	// ptb: MUST NOT be NULL. The ptb parameter points to a CACTransferBufferV2 (section
	// 2.2.3.3) structure. Constraints for the member fields are defined following. In the
	// sections following, "ptb.old" is used as shorthand to refer to the CACTransferBufferOld
	// member of the CACTransferBufferV2 structure.
	//
	// ptb.old.uTransferType MUST be CACTB_RECEIVE (0x00000001).
	//
	// ptb.old.Receive.Action MUST contain one of the following values: 0x00000000 (MQ_ACTION_RECEIVE),
	// 0x80000000 (MQ_ACTION_PEEK_CURRENT) or 0x80000001 (MQ_ACTION_PEEK_NEXT).
	//
	// On input, ptb.old.Receive.Cursor can be 0x00000000, in which case no cursor is associated
	// with the receive operation. Otherwise, ptb.old.Receive.Cursor MUST contain a Cursor
	// Handle obtained from the pcc.hCursor parameter of rpc_ACCreateCursorEx. The cursor
	// MUST have been created using the queue handle associated with the queue context value
	// provided for the hQMContext parameter, and the cursor MUST NOT have been closed prior
	// to this call. On output, the value of ptb.old.Receive.Cursor MUST be the same as
	// it was on input.
	//
	// ptb.old.Receive.ulResponseFormatNameLen is used for RPC marshaling of the ppResponseFormatName
	// buffer. On input, the client MUST set this value to the minimum of pulResponseFormatNameLenProp
	// and 1024 bytes. If ptb.old.Receive.ppResponseFormatName is NULL, this value MUST
	// be 0x00000000. On output, the server MUST set this value to the minimum of ulResponseFormatNameLen
	// and pulResponseFormatNameLenProp.
	//
	// On input, ptb.old.Receive.pulResponseFormatNameLenProp indicates the Unicode character
	// length of the buffer contained in ppResponseFormatName. On output, the server MUST
	// set this value to indicate the full length of the response queue format name associated
	// with the message being retrieved.
	//
	// On input, ptb.old.Receive.ppResponseFormatName can be NULL, in which case it MUST
	// be NULL on output. Otherwise, on successful retrieval of a message and prior to filling
	// the buffer, the server MUST verify that the pulResponseFormatNameLenProp field indicates
	// that the buffer is large enough to contain the response queue format name for the
	// retrieved message.
	//
	// ptb.old.Receive.ulAdminFormatNameLen is used for RPC marshaling of the ppAdminFormatName
	// buffer. On input, the client MUST set this value to the minimum of pulAdminFormatNameLenProp
	// and 1024 bytes. If ptb.old.Receive.ppAdminFormatName is NULL, this value MUST be
	// 0x00000000. On output, the server MUST set this value to the minimum of ulAdminFormatNameLen
	// and pulAdminFormatNameLenProp.
	//
	// On input, ptb.old.Receive.pulAdminFormatNameLenProp indicates the Unicode character
	// length of the buffer contained in ppAdminFormatName. On output, the server MUST set
	// this value to indicate the full length of the administration queue format name associated
	// with the message being retrieved.
	//
	// On input, ptb.old.Receive.ppAdminFormatName can be NULL, in which case it MUST be
	// NULL on output. Otherwise, on successful retrieval of a message and prior to filling
	// the buffer, the server MUST verify that the pulAdminFormatNameLenProp field indicates
	// that the buffer is large enough to contain the administration queue format name for
	// the retrieved message.
	//
	// ptb.old.Receive.ulDestFormatNameLen is used for RPC marshaling of the ppDestFormatName
	// buffer. On input, the client MUST set this value to the minimum of pulDestFormatNameLenProp
	// and 1024 bytes. If ptb.old.Receive.ppDestFormatName is NULL, this value MUST be 0x00000000.
	// On output, the server MUST set this value to the minimum of ulDestFormatNameLen and
	// pulDestFormatNameLenProp.
	//
	// On input, ptb.old.Receive.pulDestFormatNameLenProp indicates the Unicode character
	// length of the buffer contained in ppDestFormatName. On output, the server MUST set
	// this value to indicate the full length of the destination queue format name associated
	// with the message being retrieved.
	//
	// On input, ptb.old.Receive.ppDestFormatName can be NULL, in which case it MUST be
	// NULL on output. Otherwise, on successful retrieval of a message and prior to filling
	// the buffer, the server MUST verify that the pulDestFormatNameLenProp field indicates
	// that the buffer is large enough to contain the destination queue format name for
	// the retrieved message.
	//
	// ptb.old.Receive.ulOrderingFormatNameLen is used for RPC marshaling of the ppOrderingFormatName
	// buffer. On input, the client MUST set this value to the minimum of pulOrderingFormatNameLenProp
	// and 1024 bytes. If ptb.old.Receive.ppOrderingFormatName is NULL, this value MUST
	// be 0x00000000. On output, the server MUST set this value to the minimum of ulOrderingFormatNameLen
	// and pulOrderingFormatNameLenProp.
	//
	// On input, ptb.old.Receive.pulOrderingFormatNameLenProp indicates the Unicode character
	// length of the buffer contained in ppOrderingFormatName. On output, the server MUST
	// set this value to indicate the full length of the order queue format name associated
	// with the message being retrieved.
	//
	// On input, ptb.old.Receive.ppOrderingFormatName can be NULL, in which case it MUST
	// be NULL on output. Otherwise, on successful retrieval of a message and prior to filling
	// the buffer, the server MUST verify that the pulOrderingFormatNameLenProp field indicates
	// that the buffer is large enough to contain the order queue format name for the retrieved
	// message.
	//
	// On input, ptb.old.ppBody can be NULL, in which case it MUST be NULL on output. Otherwise,
	// on successful retrieval of a message, prior to filling the buffer, the server MUST
	// verify that the ulBodyBufferSizeInBytes field indicates that the buffer is large
	// enough to contain the message body for the retrieved message. On output, the byte
	// length of the complete body for the retrieved message MUST be returned in the pBodySize
	// field, if it is not NULL.
	//
	// On input, ptb.old.ulBodyBufferSizeInBytes MUST be 0x00000000 if ptb.old.ppBody is
	// NULL. On output, the value of ptb.old.ulBodyBufferSizeInBytes MUST be the same as
	// it was on input.
	//
	// ptb.old.ulAllocBodyBufferInBytes is used for RPC marshaling of the ppBody buffer.
	// If ppBody is NULL, this value MUST be 0x00000000.
	//
	// On input, ptb.old.pBodySize can be NULL, in which case it MUST be NULL on output.
	// Otherwise, on successful retrieval of a message, the server MUST set this value to
	// the byte length of the message body.
	//
	// ptb.old.ulTitleBufferSizeInWCHARs is used for RPC marshaling of the ptb.old.ppTitle
	// buffer. On input, the client MUST set this value to the minimum of pulTitleBufferSizeInWCHARs
	// and 250. If ptb.old.ppTitle is NULL, this value MUST be 0x00000000. On output, the
	// server MUST set this value to the minimum of ulTitleBufferSizeInWCHARs and pulTitleBufferSizeInWCHARs.
	//
	// On input, ptb.old.pulTitleBufferSizeInWCHARs indicates the Unicode character length
	// of the buffer contained in ppTitle. On output, the server MUST set this value to
	// indicate the full length of the message label associated with the message being retrieved.
	//
	// On input, ptb.old.ppTitle can be NULL, in which case it MUST be NULL on output. Otherwise,
	// on successful retrieval of a message, prior to filling the buffer, the server MUST
	// verify that the pulTitleBufferSizeInWCHARs field indicates that the buffer is large
	// enough to contain the message label for the retrieved message.
	//
	// On input, ptb.old.ppMsgExtension can be NULL, in which case it MUST be NULL on output.
	// Otherwise, on successful retrieval of a message, prior to filling the buffer, the
	// server MUST verify that the ptb.old.ulMsgExtensionBufferInBytes field indicates that
	// the buffer is large enough to contain the message extension array for the retrieved
	// message.
	//
	// On input, ptb.old.ulMsgExtensionBufferInBytes MUST be 0x00000000 if ptb.old.ppMsgExtension
	// is NULL. On output, the value of ptb.old.ulMsgExtensionBufferInBytes MUST be the
	// same as it was on input.
	//
	// On input, ptb.old.pMsgExtensionSize can be NULL, in which case it MUST be NULL on
	// output. Otherwise, the server MUST return the full length of the retrieved message
	// extension array in ptb.old.pMsgExtensionSize.
	//
	// On input, ptb.old.pUow can be NULL, in which case the Receive operation is not associated
	// with a transaction. Otherwise, ptb.old.pUow MUST contain a 16-byte transaction identifier
	// which has been enlisted by a prior call to R_QMEnlistTransaction or R_QMEnlistInternalTransaction.
	// On output, the value of ptb.old.pUow MUST be the same as it was on input.
	//
	// On input, ptb.old.ppSenderID can be NULL, in which case it MUST be NULL on output.
	// Otherwise, on successful retrieval of a message and prior to filling the buffer,
	// the server MUST verify that the ptb.old.uSenderIDLen field indicates that the buffer
	// is large enough to contain the sender ID buffer for the retrieved message.
	//
	// On input, ptb.old.pulSenderIDLenProp can be NULL; otherwise, on output, the server
	// MUST return the full length of the sender ID buffer for the retrieved message in
	// ptb.old.pulSenderIDLenProp, or 0x00000000 if the value was not included in the retrieved
	// message.
	//
	// On input, ptb.old.ppwcsProvName can be NULL, in which case it MUST be NULL on output.
	// Otherwise, prior to filling the buffer, the server MUST verify that the ptb.old.ulProvNameLen
	// field indicates that the buffer is large enough to contain the null-terminated CSP
	// name string. If the retrieved message does not include a CSP name buffer, the server
	// MUST return 0x00000000 for ptb.old.pulAuthProvNameLenProp if the pulAuthProvNameLenProp
	// pointer is not NULL.
	//
	// On input, ptb.old.pulAuthProvNameLenProp can be NULL, in which case it MUST be NULL
	// on output. Otherwise, the server MUST return the length of the CSP name buffer for
	// the retrieved message in ptb.old.pulAuthProvNameLenProp, or 0x00000000 if the value
	// was not included in the retrieved message.
	//
	// On input, ptb.old.ppSenderCert can be NULL, in which case it MUST be NULL on output.
	// Otherwise, prior to filling the buffer, the server MUST verify that the ptb.old.ulSenderCertLen
	// field indicates that the buffer is large enough to contain the sender certificate
	// buffer. If the retrieved message does not include a sender certificate, the server
	// MUST return 0x00000000 for ptb.old.pulSenderCertLenProp if the pulSenderCertLenProp
	// pointer is not NULL.
	//
	// On input, ptb.old.pulSenderCertLenProp can be NULL, in which case it MUST be NULL
	// on output. Otherwise, the server MUST return the length of the sender certificate
	// buffer for the retrieved message in ptb.old.pulSenderCertLenProp, or 0x00000000 if
	// the value is not included in the retrieved message.
	//
	// On input, ptb.old.ppSymmKeys can be NULL, in which case it MUST be NULL on output.
	// Otherwise, prior to filling the buffer, the server MUST verify that the ptb.old.ulSymmKeysSize
	// field indicates that the buffer is large enough to contain the symmetric key buffer.
	// If the retrieved message does not include a symmetric key, the server MUST return
	// 0x00000000 for ptb.old.pulSymmKeysSizeProp if the pulSymmKeysSizeProp pointer is
	// not NULL.
	//
	// On input, ptb.old.pulSymmKeysSizeProp can be NULL, in which case it MUST be NULL
	// on output. Otherwise, the server MUST return the length of the symmetric key buffer
	// for the retrieved message in ptb.old.pulSymmKeysSizeProp or 0x00000000 if the value
	// is not included in the retrieved message.
	//
	// On input, ptb.old.ppSignature can be NULL, in which case it MUST be NULL on output.
	// Otherwise, prior to filling the buffer, the server MUST verify that the ptb.old.ulSignatureSize
	// field indicates that the buffer is large enough to contain the signed hash buffer.
	// If the retrieved message does not include a signed hash, the server MUST return 0x00000000
	// for ptb.old.pulSignatureSizeProp if the pulSignatureSizeProp pointer is not NULL.
	//
	// On input, ptb.old.pulSignatureSizeProp can be NULL, in which case it MUST be NULL
	// on output. Otherwise, the server MUST return the length of the signed hash buffer
	// for the retrieved message in ptb.old.pulSignatureSizeProp, or 0x00000000 if the value
	// is not included in the retrieved message.
	//
	// * ptb.old.pClass
	//
	// * ptb.old.ppMessageID
	//
	// * ptb.old.ppCorrelationID
	//
	// * ptb.old.pSentTime
	//
	// * ptb.old.pArrivedTime
	//
	// * ptb.old.pPriority
	//
	// * ptb.old.pDelivery
	//
	// * ptb.old.pAcknowledge
	//
	// * ptb.old.pAuditing
	//
	// * ptb.old.pApplicationTag
	//
	// * ptb.old.pulRelativeTimeToQueue
	//
	// * ptb.old.pulRelativeTimeToLive
	//
	// * ptb.old.pTrace
	//
	// * ptb.old.pulPrivLevel
	//
	// * ptb.old.pAuthenticated
	//
	// * ptb.old.pulHashAlg
	//
	// * ptb.old.pulEncryptAlg
	//
	// * ptb.old.pulProvType
	//
	// * ptb.old.pulSenderIDType
	//
	// * ptb.old.ppSrcQMID
	//
	// * ptb.old.ppConnectorType
	//
	// * ptb.old.pulBodyType
	//
	// * ptb.old.pulVersion
	//
	// * ptb.pbFirstInXact
	//
	// * ptb.pbLastInXact
	//
	// * ptb.ppXactID
	//
	// * ptb.old.Send
	//
	// * ptb.old.CreateCursor
	//
	// * ptb.old.Receive.Asynchronous
	//
	// * ptb.old.ulAbsoluteTimeToQueue
	//
	// * ptb.old.ulRelativeTimeToLive
	//
	// * ptb.old.ulAuthLevel
	//
	// * ptb.old.bEncrypted
	//
	// * ptb.old.bAuthenticated
	//
	// * ptb.old.fDefaultProvider
	Ptb *qmcomm.TransferBufferV2 `idl:"name:ptb" json:"ptb"`
}

func (o *ReceiveMessageExRequest) xxx_ToOp(ctx context.Context) *xxx_ReceiveMessageExOperation {
	if o == nil {
		return &xxx_ReceiveMessageExOperation{}
	}
	return &xxx_ReceiveMessageExOperation{
		HQMContext: o.HQMContext,
		Ptb:        o.Ptb,
	}
}

func (o *ReceiveMessageExRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceiveMessageExOperation) {
	if o == nil {
		return
	}
	o.HQMContext = op.HQMContext
	o.Ptb = op.Ptb
}
func (o *ReceiveMessageExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveMessageExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveMessageExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveMessageExResponse structure represents the rpc_ACReceiveMessageEx operation response
type ReceiveMessageExResponse struct {
	// ptb: MUST NOT be NULL. The ptb parameter points to a CACTransferBufferV2 (section
	// 2.2.3.3) structure. Constraints for the member fields are defined following. In the
	// sections following, "ptb.old" is used as shorthand to refer to the CACTransferBufferOld
	// member of the CACTransferBufferV2 structure.
	//
	// ptb.old.uTransferType MUST be CACTB_RECEIVE (0x00000001).
	//
	// ptb.old.Receive.Action MUST contain one of the following values: 0x00000000 (MQ_ACTION_RECEIVE),
	// 0x80000000 (MQ_ACTION_PEEK_CURRENT) or 0x80000001 (MQ_ACTION_PEEK_NEXT).
	//
	// On input, ptb.old.Receive.Cursor can be 0x00000000, in which case no cursor is associated
	// with the receive operation. Otherwise, ptb.old.Receive.Cursor MUST contain a Cursor
	// Handle obtained from the pcc.hCursor parameter of rpc_ACCreateCursorEx. The cursor
	// MUST have been created using the queue handle associated with the queue context value
	// provided for the hQMContext parameter, and the cursor MUST NOT have been closed prior
	// to this call. On output, the value of ptb.old.Receive.Cursor MUST be the same as
	// it was on input.
	//
	// ptb.old.Receive.ulResponseFormatNameLen is used for RPC marshaling of the ppResponseFormatName
	// buffer. On input, the client MUST set this value to the minimum of pulResponseFormatNameLenProp
	// and 1024 bytes. If ptb.old.Receive.ppResponseFormatName is NULL, this value MUST
	// be 0x00000000. On output, the server MUST set this value to the minimum of ulResponseFormatNameLen
	// and pulResponseFormatNameLenProp.
	//
	// On input, ptb.old.Receive.pulResponseFormatNameLenProp indicates the Unicode character
	// length of the buffer contained in ppResponseFormatName. On output, the server MUST
	// set this value to indicate the full length of the response queue format name associated
	// with the message being retrieved.
	//
	// On input, ptb.old.Receive.ppResponseFormatName can be NULL, in which case it MUST
	// be NULL on output. Otherwise, on successful retrieval of a message and prior to filling
	// the buffer, the server MUST verify that the pulResponseFormatNameLenProp field indicates
	// that the buffer is large enough to contain the response queue format name for the
	// retrieved message.
	//
	// ptb.old.Receive.ulAdminFormatNameLen is used for RPC marshaling of the ppAdminFormatName
	// buffer. On input, the client MUST set this value to the minimum of pulAdminFormatNameLenProp
	// and 1024 bytes. If ptb.old.Receive.ppAdminFormatName is NULL, this value MUST be
	// 0x00000000. On output, the server MUST set this value to the minimum of ulAdminFormatNameLen
	// and pulAdminFormatNameLenProp.
	//
	// On input, ptb.old.Receive.pulAdminFormatNameLenProp indicates the Unicode character
	// length of the buffer contained in ppAdminFormatName. On output, the server MUST set
	// this value to indicate the full length of the administration queue format name associated
	// with the message being retrieved.
	//
	// On input, ptb.old.Receive.ppAdminFormatName can be NULL, in which case it MUST be
	// NULL on output. Otherwise, on successful retrieval of a message and prior to filling
	// the buffer, the server MUST verify that the pulAdminFormatNameLenProp field indicates
	// that the buffer is large enough to contain the administration queue format name for
	// the retrieved message.
	//
	// ptb.old.Receive.ulDestFormatNameLen is used for RPC marshaling of the ppDestFormatName
	// buffer. On input, the client MUST set this value to the minimum of pulDestFormatNameLenProp
	// and 1024 bytes. If ptb.old.Receive.ppDestFormatName is NULL, this value MUST be 0x00000000.
	// On output, the server MUST set this value to the minimum of ulDestFormatNameLen and
	// pulDestFormatNameLenProp.
	//
	// On input, ptb.old.Receive.pulDestFormatNameLenProp indicates the Unicode character
	// length of the buffer contained in ppDestFormatName. On output, the server MUST set
	// this value to indicate the full length of the destination queue format name associated
	// with the message being retrieved.
	//
	// On input, ptb.old.Receive.ppDestFormatName can be NULL, in which case it MUST be
	// NULL on output. Otherwise, on successful retrieval of a message and prior to filling
	// the buffer, the server MUST verify that the pulDestFormatNameLenProp field indicates
	// that the buffer is large enough to contain the destination queue format name for
	// the retrieved message.
	//
	// ptb.old.Receive.ulOrderingFormatNameLen is used for RPC marshaling of the ppOrderingFormatName
	// buffer. On input, the client MUST set this value to the minimum of pulOrderingFormatNameLenProp
	// and 1024 bytes. If ptb.old.Receive.ppOrderingFormatName is NULL, this value MUST
	// be 0x00000000. On output, the server MUST set this value to the minimum of ulOrderingFormatNameLen
	// and pulOrderingFormatNameLenProp.
	//
	// On input, ptb.old.Receive.pulOrderingFormatNameLenProp indicates the Unicode character
	// length of the buffer contained in ppOrderingFormatName. On output, the server MUST
	// set this value to indicate the full length of the order queue format name associated
	// with the message being retrieved.
	//
	// On input, ptb.old.Receive.ppOrderingFormatName can be NULL, in which case it MUST
	// be NULL on output. Otherwise, on successful retrieval of a message and prior to filling
	// the buffer, the server MUST verify that the pulOrderingFormatNameLenProp field indicates
	// that the buffer is large enough to contain the order queue format name for the retrieved
	// message.
	//
	// On input, ptb.old.ppBody can be NULL, in which case it MUST be NULL on output. Otherwise,
	// on successful retrieval of a message, prior to filling the buffer, the server MUST
	// verify that the ulBodyBufferSizeInBytes field indicates that the buffer is large
	// enough to contain the message body for the retrieved message. On output, the byte
	// length of the complete body for the retrieved message MUST be returned in the pBodySize
	// field, if it is not NULL.
	//
	// On input, ptb.old.ulBodyBufferSizeInBytes MUST be 0x00000000 if ptb.old.ppBody is
	// NULL. On output, the value of ptb.old.ulBodyBufferSizeInBytes MUST be the same as
	// it was on input.
	//
	// ptb.old.ulAllocBodyBufferInBytes is used for RPC marshaling of the ppBody buffer.
	// If ppBody is NULL, this value MUST be 0x00000000.
	//
	// On input, ptb.old.pBodySize can be NULL, in which case it MUST be NULL on output.
	// Otherwise, on successful retrieval of a message, the server MUST set this value to
	// the byte length of the message body.
	//
	// ptb.old.ulTitleBufferSizeInWCHARs is used for RPC marshaling of the ptb.old.ppTitle
	// buffer. On input, the client MUST set this value to the minimum of pulTitleBufferSizeInWCHARs
	// and 250. If ptb.old.ppTitle is NULL, this value MUST be 0x00000000. On output, the
	// server MUST set this value to the minimum of ulTitleBufferSizeInWCHARs and pulTitleBufferSizeInWCHARs.
	//
	// On input, ptb.old.pulTitleBufferSizeInWCHARs indicates the Unicode character length
	// of the buffer contained in ppTitle. On output, the server MUST set this value to
	// indicate the full length of the message label associated with the message being retrieved.
	//
	// On input, ptb.old.ppTitle can be NULL, in which case it MUST be NULL on output. Otherwise,
	// on successful retrieval of a message, prior to filling the buffer, the server MUST
	// verify that the pulTitleBufferSizeInWCHARs field indicates that the buffer is large
	// enough to contain the message label for the retrieved message.
	//
	// On input, ptb.old.ppMsgExtension can be NULL, in which case it MUST be NULL on output.
	// Otherwise, on successful retrieval of a message, prior to filling the buffer, the
	// server MUST verify that the ptb.old.ulMsgExtensionBufferInBytes field indicates that
	// the buffer is large enough to contain the message extension array for the retrieved
	// message.
	//
	// On input, ptb.old.ulMsgExtensionBufferInBytes MUST be 0x00000000 if ptb.old.ppMsgExtension
	// is NULL. On output, the value of ptb.old.ulMsgExtensionBufferInBytes MUST be the
	// same as it was on input.
	//
	// On input, ptb.old.pMsgExtensionSize can be NULL, in which case it MUST be NULL on
	// output. Otherwise, the server MUST return the full length of the retrieved message
	// extension array in ptb.old.pMsgExtensionSize.
	//
	// On input, ptb.old.pUow can be NULL, in which case the Receive operation is not associated
	// with a transaction. Otherwise, ptb.old.pUow MUST contain a 16-byte transaction identifier
	// which has been enlisted by a prior call to R_QMEnlistTransaction or R_QMEnlistInternalTransaction.
	// On output, the value of ptb.old.pUow MUST be the same as it was on input.
	//
	// On input, ptb.old.ppSenderID can be NULL, in which case it MUST be NULL on output.
	// Otherwise, on successful retrieval of a message and prior to filling the buffer,
	// the server MUST verify that the ptb.old.uSenderIDLen field indicates that the buffer
	// is large enough to contain the sender ID buffer for the retrieved message.
	//
	// On input, ptb.old.pulSenderIDLenProp can be NULL; otherwise, on output, the server
	// MUST return the full length of the sender ID buffer for the retrieved message in
	// ptb.old.pulSenderIDLenProp, or 0x00000000 if the value was not included in the retrieved
	// message.
	//
	// On input, ptb.old.ppwcsProvName can be NULL, in which case it MUST be NULL on output.
	// Otherwise, prior to filling the buffer, the server MUST verify that the ptb.old.ulProvNameLen
	// field indicates that the buffer is large enough to contain the null-terminated CSP
	// name string. If the retrieved message does not include a CSP name buffer, the server
	// MUST return 0x00000000 for ptb.old.pulAuthProvNameLenProp if the pulAuthProvNameLenProp
	// pointer is not NULL.
	//
	// On input, ptb.old.pulAuthProvNameLenProp can be NULL, in which case it MUST be NULL
	// on output. Otherwise, the server MUST return the length of the CSP name buffer for
	// the retrieved message in ptb.old.pulAuthProvNameLenProp, or 0x00000000 if the value
	// was not included in the retrieved message.
	//
	// On input, ptb.old.ppSenderCert can be NULL, in which case it MUST be NULL on output.
	// Otherwise, prior to filling the buffer, the server MUST verify that the ptb.old.ulSenderCertLen
	// field indicates that the buffer is large enough to contain the sender certificate
	// buffer. If the retrieved message does not include a sender certificate, the server
	// MUST return 0x00000000 for ptb.old.pulSenderCertLenProp if the pulSenderCertLenProp
	// pointer is not NULL.
	//
	// On input, ptb.old.pulSenderCertLenProp can be NULL, in which case it MUST be NULL
	// on output. Otherwise, the server MUST return the length of the sender certificate
	// buffer for the retrieved message in ptb.old.pulSenderCertLenProp, or 0x00000000 if
	// the value is not included in the retrieved message.
	//
	// On input, ptb.old.ppSymmKeys can be NULL, in which case it MUST be NULL on output.
	// Otherwise, prior to filling the buffer, the server MUST verify that the ptb.old.ulSymmKeysSize
	// field indicates that the buffer is large enough to contain the symmetric key buffer.
	// If the retrieved message does not include a symmetric key, the server MUST return
	// 0x00000000 for ptb.old.pulSymmKeysSizeProp if the pulSymmKeysSizeProp pointer is
	// not NULL.
	//
	// On input, ptb.old.pulSymmKeysSizeProp can be NULL, in which case it MUST be NULL
	// on output. Otherwise, the server MUST return the length of the symmetric key buffer
	// for the retrieved message in ptb.old.pulSymmKeysSizeProp or 0x00000000 if the value
	// is not included in the retrieved message.
	//
	// On input, ptb.old.ppSignature can be NULL, in which case it MUST be NULL on output.
	// Otherwise, prior to filling the buffer, the server MUST verify that the ptb.old.ulSignatureSize
	// field indicates that the buffer is large enough to contain the signed hash buffer.
	// If the retrieved message does not include a signed hash, the server MUST return 0x00000000
	// for ptb.old.pulSignatureSizeProp if the pulSignatureSizeProp pointer is not NULL.
	//
	// On input, ptb.old.pulSignatureSizeProp can be NULL, in which case it MUST be NULL
	// on output. Otherwise, the server MUST return the length of the signed hash buffer
	// for the retrieved message in ptb.old.pulSignatureSizeProp, or 0x00000000 if the value
	// is not included in the retrieved message.
	//
	// * ptb.old.pClass
	//
	// * ptb.old.ppMessageID
	//
	// * ptb.old.ppCorrelationID
	//
	// * ptb.old.pSentTime
	//
	// * ptb.old.pArrivedTime
	//
	// * ptb.old.pPriority
	//
	// * ptb.old.pDelivery
	//
	// * ptb.old.pAcknowledge
	//
	// * ptb.old.pAuditing
	//
	// * ptb.old.pApplicationTag
	//
	// * ptb.old.pulRelativeTimeToQueue
	//
	// * ptb.old.pulRelativeTimeToLive
	//
	// * ptb.old.pTrace
	//
	// * ptb.old.pulPrivLevel
	//
	// * ptb.old.pAuthenticated
	//
	// * ptb.old.pulHashAlg
	//
	// * ptb.old.pulEncryptAlg
	//
	// * ptb.old.pulProvType
	//
	// * ptb.old.pulSenderIDType
	//
	// * ptb.old.ppSrcQMID
	//
	// * ptb.old.ppConnectorType
	//
	// * ptb.old.pulBodyType
	//
	// * ptb.old.pulVersion
	//
	// * ptb.pbFirstInXact
	//
	// * ptb.pbLastInXact
	//
	// * ptb.ppXactID
	//
	// * ptb.old.Send
	//
	// * ptb.old.CreateCursor
	//
	// * ptb.old.Receive.Asynchronous
	//
	// * ptb.old.ulAbsoluteTimeToQueue
	//
	// * ptb.old.ulRelativeTimeToLive
	//
	// * ptb.old.ulAuthLevel
	//
	// * ptb.old.bEncrypted
	//
	// * ptb.old.bAuthenticated
	//
	// * ptb.old.fDefaultProvider
	Ptb *qmcomm.TransferBufferV2 `idl:"name:ptb" json:"ptb"`
	// Return: The rpc_ACReceiveMessageEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveMessageExResponse) xxx_ToOp(ctx context.Context) *xxx_ReceiveMessageExOperation {
	if o == nil {
		return &xxx_ReceiveMessageExOperation{}
	}
	return &xxx_ReceiveMessageExOperation{
		Ptb:    o.Ptb,
		Return: o.Return,
	}
}

func (o *ReceiveMessageExResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceiveMessageExOperation) {
	if o == nil {
		return
	}
	o.Ptb = op.Ptb
	o.Return = op.Return
}
func (o *ReceiveMessageExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveMessageExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveMessageExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateCursorExOperation structure represents the rpc_ACCreateCursorEx operation
type xxx_CreateCursorExOperation struct {
	Queue  *qmcomm.Queue              `idl:"name:hQueue" json:"queue"`
	Pcc    *qmcomm.CreateRemoteCursor `idl:"name:pcc" json:"pcc"`
	Return int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateCursorExOperation) OpNum() int { return 3 }

func (o *xxx_CreateCursorExOperation) OpName() string { return "/qmcomm2/v1/rpc_ACCreateCursorEx" }

func (o *xxx_CreateCursorExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCursorExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hQueue {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue != nil {
			if err := o.Queue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&qmcomm.Queue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pcc {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.Pcc != nil {
			if err := o.Pcc.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&qmcomm.CreateRemoteCursor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateCursorExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQueue {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue == nil {
			o.Queue = &qmcomm.Queue{}
		}
		if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pcc {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.Pcc == nil {
			o.Pcc = &qmcomm.CreateRemoteCursor{}
		}
		if err := o.Pcc.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCursorExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateCursorExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pcc {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.Pcc != nil {
			if err := o.Pcc.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&qmcomm.CreateRemoteCursor{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CreateCursorExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pcc {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.Pcc == nil {
			o.Pcc = &qmcomm.CreateRemoteCursor{}
		}
		if err := o.Pcc.UnmarshalNDR(ctx, w); err != nil {
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

// CreateCursorExRequest structure represents the rpc_ACCreateCursorEx operation request
type CreateCursorExRequest struct {
	// hQueue:  MUST be an RPC_QUEUE_HANDLE (section 2.2.1.1.2) acquired from the phQueue
	// parameter of rpc_QMOpenQueueInternal (section 3.1.4.17). Prior to this method being
	// invoked, the queue MUST NOT have been deleted, and the queue handle MUST NOT have
	// been closed.
	Queue *qmcomm.Queue `idl:"name:hQueue" json:"queue"`
	// pcc:  A pointer to a CACCreateRemoteCursor (section 2.2.3.4) structure. MUST NOT
	// be NULL.
	Pcc *qmcomm.CreateRemoteCursor `idl:"name:pcc" json:"pcc"`
}

func (o *CreateCursorExRequest) xxx_ToOp(ctx context.Context) *xxx_CreateCursorExOperation {
	if o == nil {
		return &xxx_CreateCursorExOperation{}
	}
	return &xxx_CreateCursorExOperation{
		Queue: o.Queue,
		Pcc:   o.Pcc,
	}
}

func (o *CreateCursorExRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateCursorExOperation) {
	if o == nil {
		return
	}
	o.Queue = op.Queue
	o.Pcc = op.Pcc
}
func (o *CreateCursorExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateCursorExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateCursorExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateCursorExResponse structure represents the rpc_ACCreateCursorEx operation response
type CreateCursorExResponse struct {
	// pcc:  A pointer to a CACCreateRemoteCursor (section 2.2.3.4) structure. MUST NOT
	// be NULL.
	Pcc *qmcomm.CreateRemoteCursor `idl:"name:pcc" json:"pcc"`
	// Return: The rpc_ACCreateCursorEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateCursorExResponse) xxx_ToOp(ctx context.Context) *xxx_CreateCursorExOperation {
	if o == nil {
		return &xxx_CreateCursorExOperation{}
	}
	return &xxx_CreateCursorExOperation{
		Pcc:    o.Pcc,
		Return: o.Return,
	}
}

func (o *CreateCursorExResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateCursorExOperation) {
	if o == nil {
		return
	}
	o.Pcc = op.Pcc
	o.Return = op.Return
}
func (o *CreateCursorExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateCursorExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateCursorExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
