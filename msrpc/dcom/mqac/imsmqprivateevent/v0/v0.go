package imsmqprivateevent

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
	ImsmqPrivateEventIID = &dcom.IID{Data1: 0xd7ab3341, Data2: 0xc9d3, Data3: 0x11d1, Data4: []byte{0xbb, 0x47, 0x00, 0x80, 0xc7, 0xc5, 0xa2, 0xc0}}
	// Syntax UUID
	ImsmqPrivateEventSyntaxUUID = &uuid.UUID{TimeLow: 0xd7ab3341, TimeMid: 0xc9d3, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0x47, Node: [6]uint8{0x0, 0x80, 0xc7, 0xc5, 0xa2, 0xc0}}
	// Syntax ID
	ImsmqPrivateEventSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ImsmqPrivateEventSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQPrivateEvent interface.
type ImsmqPrivateEventClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// Hwnd operation.
	GetHwnd(context.Context, *GetHwndRequest, ...dcerpc.CallOption) (*GetHwndResponse, error)

	// FireArrivedEvent operation.
	FireArrivedEvent(context.Context, *FireArrivedEventRequest, ...dcerpc.CallOption) (*FireArrivedEventResponse, error)

	// FireArrivedErrorEvent operation.
	FireArrivedErrorEvent(context.Context, *FireArrivedErrorEventRequest, ...dcerpc.CallOption) (*FireArrivedErrorEventResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ImsmqPrivateEventClient
}

type xxx_DefaultImsmqPrivateEventClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultImsmqPrivateEventClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultImsmqPrivateEventClient) GetHwnd(ctx context.Context, in *GetHwndRequest, opts ...dcerpc.CallOption) (*GetHwndResponse, error) {
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
	out := &GetHwndResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqPrivateEventClient) FireArrivedEvent(ctx context.Context, in *FireArrivedEventRequest, opts ...dcerpc.CallOption) (*FireArrivedEventResponse, error) {
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
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqPrivateEventClient) FireArrivedErrorEvent(ctx context.Context, in *FireArrivedErrorEventRequest, opts ...dcerpc.CallOption) (*FireArrivedErrorEventResponse, error) {
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
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqPrivateEventClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultImsmqPrivateEventClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultImsmqPrivateEventClient) IPID(ctx context.Context, ipid *dcom.IPID) ImsmqPrivateEventClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultImsmqPrivateEventClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewImsmqPrivateEventClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ImsmqPrivateEventClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ImsmqPrivateEventSyntaxV0_0))...)
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
	return &xxx_DefaultImsmqPrivateEventClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetHwndOperation structure represents the Hwnd operation
type xxx_GetHwndOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Phwnd  int32          `idl:"name:phwnd" json:"phwnd"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetHwndOperation) OpNum() int { return 7 }

func (o *xxx_GetHwndOperation) OpName() string { return "/IMSMQPrivateEvent/v0/Hwnd" }

func (o *xxx_GetHwndOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHwndOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetHwndOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetHwndOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHwndOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WriteData(o.Phwnd); err != nil {
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

func (o *xxx_GetHwndOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
		if err := w.ReadData(&o.Phwnd); err != nil {
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

// GetHwndRequest structure represents the Hwnd operation request
type GetHwndRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetHwndRequest) xxx_ToOp(ctx context.Context, op *xxx_GetHwndOperation) *xxx_GetHwndOperation {
	if op == nil {
		op = &xxx_GetHwndOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetHwndRequest) xxx_FromOp(ctx context.Context, op *xxx_GetHwndOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetHwndRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetHwndRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHwndOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetHwndResponse structure represents the Hwnd operation response
type GetHwndResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Phwnd int32          `idl:"name:phwnd" json:"phwnd"`
	// Return: The Hwnd return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetHwndResponse) xxx_ToOp(ctx context.Context, op *xxx_GetHwndOperation) *xxx_GetHwndOperation {
	if op == nil {
		op = &xxx_GetHwndOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Phwnd = o.Phwnd
	op.Return = o.Return
	return op
}

func (o *GetHwndResponse) xxx_FromOp(ctx context.Context, op *xxx_GetHwndOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Phwnd = op.Phwnd
	o.Return = op.Return
}
func (o *GetHwndResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetHwndResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHwndOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FireArrivedEventOperation structure represents the FireArrivedEvent operation
type xxx_FireArrivedEventOperation struct {
	This      *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Pq        *mqac.ImsmqQueue `idl:"name:pq" json:"pq"`
	Msgcursor int32            `idl:"name:msgcursor" json:"msgcursor"`
	Return    int32            `idl:"name:Return" json:"return"`
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
		if o.Pq != nil {
			_ptr_pq := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Pq != nil {
					if err := o.Pq.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.ImsmqQueue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Pq, _ptr_pq); err != nil {
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
		if err := w.WriteData(o.Msgcursor); err != nil {
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
			if o.Pq == nil {
				o.Pq = &mqac.ImsmqQueue{}
			}
			if err := o.Pq.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pq := func(ptr interface{}) { o.Pq = *ptr.(**mqac.ImsmqQueue) }
		if err := w.ReadPointer(&o.Pq, _s_pq, _ptr_pq); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// msgcursor {in} (1:(int32))
	{
		if err := w.ReadData(&o.Msgcursor); err != nil {
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
	This      *dcom.ORPCThis   `idl:"name:This" json:"this"`
	Pq        *mqac.ImsmqQueue `idl:"name:pq" json:"pq"`
	Msgcursor int32            `idl:"name:msgcursor" json:"msgcursor"`
}

func (o *FireArrivedEventRequest) xxx_ToOp(ctx context.Context, op *xxx_FireArrivedEventOperation) *xxx_FireArrivedEventOperation {
	if op == nil {
		op = &xxx_FireArrivedEventOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Pq = o.Pq
	op.Msgcursor = o.Msgcursor
	return op
}

func (o *FireArrivedEventRequest) xxx_FromOp(ctx context.Context, op *xxx_FireArrivedEventOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Pq = op.Pq
	o.Msgcursor = op.Msgcursor
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
	This      *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Pq        *mqac.ImsmqQueue `idl:"name:pq" json:"pq"`
	Status    int32            `idl:"name:hrStatus" json:"status"`
	Msgcursor int32            `idl:"name:msgcursor" json:"msgcursor"`
	Return    int32            `idl:"name:Return" json:"return"`
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
		if o.Pq != nil {
			_ptr_pq := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Pq != nil {
					if err := o.Pq.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.ImsmqQueue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Pq, _ptr_pq); err != nil {
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
		if err := w.WriteData(o.Msgcursor); err != nil {
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
			if o.Pq == nil {
				o.Pq = &mqac.ImsmqQueue{}
			}
			if err := o.Pq.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pq := func(ptr interface{}) { o.Pq = *ptr.(**mqac.ImsmqQueue) }
		if err := w.ReadPointer(&o.Pq, _s_pq, _ptr_pq); err != nil {
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
		if err := w.ReadData(&o.Msgcursor); err != nil {
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
	This      *dcom.ORPCThis   `idl:"name:This" json:"this"`
	Pq        *mqac.ImsmqQueue `idl:"name:pq" json:"pq"`
	Status    int32            `idl:"name:hrStatus" json:"status"`
	Msgcursor int32            `idl:"name:msgcursor" json:"msgcursor"`
}

func (o *FireArrivedErrorEventRequest) xxx_ToOp(ctx context.Context, op *xxx_FireArrivedErrorEventOperation) *xxx_FireArrivedErrorEventOperation {
	if op == nil {
		op = &xxx_FireArrivedErrorEventOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Pq = o.Pq
	op.Status = o.Status
	op.Msgcursor = o.Msgcursor
	return op
}

func (o *FireArrivedErrorEventRequest) xxx_FromOp(ctx context.Context, op *xxx_FireArrivedErrorEventOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Pq = op.Pq
	o.Status = op.Status
	o.Msgcursor = op.Msgcursor
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
