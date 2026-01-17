package imsmqprivateevent

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = mqac.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQPrivateEvent interface identifier d7ab3341-c9d3-11d1-bb47-0080c7c5a2c0
	PrivateEventIID = &dcom.IID{Data1: 0xd7ab3341, Data2: 0xc9d3, Data3: 0x11d1, Data4: []byte{0xbb, 0x47, 0x00, 0x80, 0xc7, 0xc5, 0xa2, 0xc0}}
	// Syntax UUID
	PrivateEventSyntaxUUID = &uuid.UUID{TimeLow: 0xd7ab3341, TimeMid: 0xc9d3, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0x47, Node: [6]uint8{0x0, 0x80, 0xc7, 0xc5, 0xa2, 0xc0}}
	// Syntax ID
	PrivateEventSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: PrivateEventSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQPrivateEvent interface.
type PrivateEventClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The Hwnd method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a LONG value, which the client MUST ignore. Because the returned
	// LONG value serves no purpose, the server MAY<90> return 0x00000000.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	//
	// This interface is optional for communication with the client. If implemented, this
	// interface MUST reside on the server.
	GetHandle(context.Context, *GetHandleRequest, ...dcerpc.CallOption) (*GetHandleResponse, error)

	// The FireArrivedEvent method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST provide notification of the availability of a Message.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	FireArrivedEvent(context.Context, *FireArrivedEventRequest, ...dcerpc.CallOption) (*FireArrivedEventResponse, error)

	// The FireArrivedErrorEvent method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST provide notification of an error relating to the arrival
	// of a message.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	FireArrivedErrorEvent(context.Context, *FireArrivedErrorEventRequest, ...dcerpc.CallOption) (*FireArrivedErrorEventResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) PrivateEventClient
}

type xxx_DefaultPrivateEventClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPrivateEventClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultPrivateEventClient) GetHandle(ctx context.Context, in *GetHandleRequest, opts ...dcerpc.CallOption) (*GetHandleResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetHandleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPrivateEventClient) FireArrivedEvent(ctx context.Context, in *FireArrivedEventRequest, opts ...dcerpc.CallOption) (*FireArrivedEventResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FireArrivedEventResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPrivateEventClient) FireArrivedErrorEvent(ctx context.Context, in *FireArrivedErrorEventRequest, opts ...dcerpc.CallOption) (*FireArrivedErrorEventResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FireArrivedErrorEventResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPrivateEventClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPrivateEventClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultPrivateEventClient) IPID(ctx context.Context, ipid *dcom.IPID) PrivateEventClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPrivateEventClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewPrivateEventClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PrivateEventClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PrivateEventSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultPrivateEventClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetHandleOperation structure represents the Hwnd operation
type xxx_GetHandleOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle int32          `idl:"name:phwnd" json:"handle"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetHandleOperation) OpNum() int { return 7 }

func (o *xxx_GetHandleOperation) OpName() string { return "/IMSMQPrivateEvent/v0/Hwnd" }

func (o *xxx_GetHandleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// phwnd {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Handle); err != nil {
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

func (o *xxx_GetHandleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// phwnd {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
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

// GetHandleRequest structure represents the Hwnd operation request
type GetHandleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetHandleRequest) xxx_ToOp(ctx context.Context, op *xxx_GetHandleOperation) *xxx_GetHandleOperation {
	if op == nil {
		op = &xxx_GetHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetHandleRequest) xxx_FromOp(ctx context.Context, op *xxx_GetHandleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetHandleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetHandleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHandleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetHandleResponse structure represents the Hwnd operation response
type GetHandleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// phwnd: A pointer to a long that MAY contain 0x00000000. The value returned via this
	// parameter MUST be ignored by the client, and it serves no purpose.
	Handle int32 `idl:"name:phwnd" json:"handle"`
	// Return: The Hwnd return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetHandleResponse) xxx_ToOp(ctx context.Context, op *xxx_GetHandleOperation) *xxx_GetHandleOperation {
	if op == nil {
		op = &xxx_GetHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Handle = o.Handle
	op.Return = o.Return
	return op
}

func (o *GetHandleResponse) xxx_FromOp(ctx context.Context, op *xxx_GetHandleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Handle = op.Handle
	o.Return = op.Return
}
func (o *GetHandleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetHandleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHandleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FireArrivedEventOperation structure represents the FireArrivedEvent operation
type xxx_FireArrivedEventOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	Queue         *mqac.Queue    `idl:"name:pq" json:"queue"`
	MessageCursor int32          `idl:"name:msgcursor" json:"message_cursor"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FireArrivedEventOperation) OpNum() int { return 8 }

func (o *xxx_FireArrivedEventOperation) OpName() string {
	return "/IMSMQPrivateEvent/v0/FireArrivedEvent"
}

func (o *xxx_FireArrivedEventOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FireArrivedEventOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pq {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueue}(interface))
	{
		if o.Queue != nil {
			_ptr_pq := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Queue != nil {
					if err := o.Queue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Queue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Queue, _ptr_pq); err != nil {
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
	// msgcursor {in} (1:(int32))
	{
		if err := w.WriteData(o.MessageCursor); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FireArrivedEventOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pq {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueue}(interface))
	{
		_ptr_pq := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Queue == nil {
				o.Queue = &mqac.Queue{}
			}
			if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pq := func(ptr interface{}) { o.Queue = *ptr.(**mqac.Queue) }
		if err := w.ReadPointer(&o.Queue, _s_pq, _ptr_pq); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// msgcursor {in} (1:(int32))
	{
		if err := w.ReadData(&o.MessageCursor); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FireArrivedEventOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FireArrivedEventOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_FireArrivedEventOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
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

// FireArrivedEventRequest structure represents the FireArrivedEvent operation request
type FireArrivedEventRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pq: A pointer to an IMSMQQueue interface that upon success will be cast to an IDispatch
	// pointer.
	Queue *mqac.Queue `idl:"name:pq" json:"queue"`
	// msgcursor: A long value that specifies the value of the cursor option that was specified
	// through the Cursor input parameter that was passed to the IMSMQQueue4::EnableNotification
	// operation to associate this MSMQEvent with MSMQQueue. This parameter corresponds
	// to the MQMSGCURSOR (section 2.2.2.8) enum.
	MessageCursor int32 `idl:"name:msgcursor" json:"message_cursor"`
}

func (o *FireArrivedEventRequest) xxx_ToOp(ctx context.Context, op *xxx_FireArrivedEventOperation) *xxx_FireArrivedEventOperation {
	if op == nil {
		op = &xxx_FireArrivedEventOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Queue = o.Queue
	op.MessageCursor = o.MessageCursor
	return op
}

func (o *FireArrivedEventRequest) xxx_FromOp(ctx context.Context, op *xxx_FireArrivedEventOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Queue = op.Queue
	o.MessageCursor = op.MessageCursor
}
func (o *FireArrivedEventRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FireArrivedEventRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FireArrivedEventOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FireArrivedEventResponse structure represents the FireArrivedEvent operation response
type FireArrivedEventResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FireArrivedEvent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FireArrivedEventResponse) xxx_ToOp(ctx context.Context, op *xxx_FireArrivedEventOperation) *xxx_FireArrivedEventOperation {
	if op == nil {
		op = &xxx_FireArrivedEventOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *FireArrivedEventResponse) xxx_FromOp(ctx context.Context, op *xxx_FireArrivedEventOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *FireArrivedEventResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FireArrivedEventResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FireArrivedEventOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FireArrivedErrorEventOperation structure represents the FireArrivedErrorEvent operation
type xxx_FireArrivedErrorEventOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	Queue         *mqac.Queue    `idl:"name:pq" json:"queue"`
	Status        int32          `idl:"name:hrStatus" json:"status"`
	MessageCursor int32          `idl:"name:msgcursor" json:"message_cursor"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FireArrivedErrorEventOperation) OpNum() int { return 9 }

func (o *xxx_FireArrivedErrorEventOperation) OpName() string {
	return "/IMSMQPrivateEvent/v0/FireArrivedErrorEvent"
}

func (o *xxx_FireArrivedErrorEventOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FireArrivedErrorEventOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pq {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueue}(interface))
	{
		if o.Queue != nil {
			_ptr_pq := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Queue != nil {
					if err := o.Queue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Queue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Queue, _ptr_pq); err != nil {
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
	// hrStatus {in} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// msgcursor {in} (1:(int32))
	{
		if err := w.WriteData(o.MessageCursor); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FireArrivedErrorEventOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pq {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQQueue}(interface))
	{
		_ptr_pq := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Queue == nil {
				o.Queue = &mqac.Queue{}
			}
			if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pq := func(ptr interface{}) { o.Queue = *ptr.(**mqac.Queue) }
		if err := w.ReadPointer(&o.Queue, _s_pq, _ptr_pq); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hrStatus {in} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// msgcursor {in} (1:(int32))
	{
		if err := w.ReadData(&o.MessageCursor); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FireArrivedErrorEventOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FireArrivedErrorEventOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_FireArrivedErrorEventOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
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

// FireArrivedErrorEventRequest structure represents the FireArrivedErrorEvent operation request
type FireArrivedErrorEventRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pq: A pointer to an IMSMQQueue interface that upon success will be cast to an IDispatch
	// pointer and forwarded on to the client.
	Queue *mqac.Queue `idl:"name:pq" json:"queue"`
	// hrStatus: An HRESULT value that specifies the error code that was received from the
	// Queue where the Message was delivered.
	Status int32 `idl:"name:hrStatus" json:"status"`
	// msgcursor: A long value that specifies the value of the cursor option that was specified
	// through the Cursor input parameter that was passed to the IMSMQQueue4::EnableNotification
	// operation to associate this MSMQEvent with the queue. This parameter corresponds
	// to the MQMSGCURSOR (section 2.2.2.8) enum.
	MessageCursor int32 `idl:"name:msgcursor" json:"message_cursor"`
}

func (o *FireArrivedErrorEventRequest) xxx_ToOp(ctx context.Context, op *xxx_FireArrivedErrorEventOperation) *xxx_FireArrivedErrorEventOperation {
	if op == nil {
		op = &xxx_FireArrivedErrorEventOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Queue = o.Queue
	op.Status = o.Status
	op.MessageCursor = o.MessageCursor
	return op
}

func (o *FireArrivedErrorEventRequest) xxx_FromOp(ctx context.Context, op *xxx_FireArrivedErrorEventOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Queue = op.Queue
	o.Status = op.Status
	o.MessageCursor = op.MessageCursor
}
func (o *FireArrivedErrorEventRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FireArrivedErrorEventRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FireArrivedErrorEventOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FireArrivedErrorEventResponse structure represents the FireArrivedErrorEvent operation response
type FireArrivedErrorEventResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FireArrivedErrorEvent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FireArrivedErrorEventResponse) xxx_ToOp(ctx context.Context, op *xxx_FireArrivedErrorEventOperation) *xxx_FireArrivedErrorEventOperation {
	if op == nil {
		op = &xxx_FireArrivedErrorEventOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *FireArrivedErrorEventResponse) xxx_FromOp(ctx context.Context, op *xxx_FireArrivedErrorEventOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *FireArrivedErrorEventResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FireArrivedErrorEventResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FireArrivedErrorEventOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
