package intmsnotifysink

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	rsmp "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = rsmp.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// INtmsNotifySink interface identifier bb39332c-bfee-4380-ad8a-badc8aff5bb6
	NotifySinkIID = &dcom.IID{Data1: 0xbb39332c, Data2: 0xbfee, Data3: 0x4380, Data4: []byte{0xad, 0x8a, 0xba, 0xdc, 0x8a, 0xff, 0x5b, 0xb6}}
	// Syntax UUID
	NotifySinkSyntaxUUID = &uuid.UUID{TimeLow: 0xbb39332c, TimeMid: 0xbfee, TimeHiAndVersion: 0x4380, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0x8a, Node: [6]uint8{0xba, 0xdc, 0x8a, 0xff, 0x5b, 0xb6}}
	// Syntax ID
	NotifySinkSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: NotifySinkSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// INtmsNotifySink interface.
type NotifySinkClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	ConnectCallback(context.Context, *ConnectCallbackRequest, ...dcerpc.CallOption) (*ConnectCallbackResponse, error)

	// OnNotify operation.
	OnNotify(context.Context, *OnNotifyRequest, ...dcerpc.CallOption) (*OnNotifyResponse, error)

	ReleaseCallback(context.Context, *ReleaseCallbackRequest, ...dcerpc.CallOption) (*ReleaseCallbackResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) NotifySinkClient
}

type xxx_DefaultNotifySinkClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultNotifySinkClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultNotifySinkClient) ConnectCallback(ctx context.Context, in *ConnectCallbackRequest, opts ...dcerpc.CallOption) (*ConnectCallbackResponse, error) {
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
	out := &ConnectCallbackResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNotifySinkClient) OnNotify(ctx context.Context, in *OnNotifyRequest, opts ...dcerpc.CallOption) (*OnNotifyResponse, error) {
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
	out := &OnNotifyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNotifySinkClient) ReleaseCallback(ctx context.Context, in *ReleaseCallbackRequest, opts ...dcerpc.CallOption) (*ReleaseCallbackResponse, error) {
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
	out := &ReleaseCallbackResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNotifySinkClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultNotifySinkClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultNotifySinkClient) IPID(ctx context.Context, ipid *dcom.IPID) NotifySinkClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultNotifySinkClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewNotifySinkClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (NotifySinkClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(NotifySinkSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultNotifySinkClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ConnectCallbackOperation structure represents the ConnectCallback operation
type xxx_ConnectCallbackOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConnectionPoint *rsmp.Unknown  `idl:"name:pUnkCP" json:"connection_point"`
	Sink            *rsmp.Unknown  `idl:"name:pUnkSink" json:"sink"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectCallbackOperation) OpNum() int { return 3 }

func (o *xxx_ConnectCallbackOperation) OpName() string { return "/INtmsNotifySink/v0/ConnectCallback" }

func (o *xxx_ConnectCallbackOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectCallbackOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pUnkCP {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.ConnectionPoint != nil {
			_ptr_pUnkCP := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConnectionPoint != nil {
					if err := o.ConnectionPoint.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rsmp.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConnectionPoint, _ptr_pUnkCP); err != nil {
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
	// pUnkSink {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Sink != nil {
			_ptr_pUnkSink := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Sink != nil {
					if err := o.Sink.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rsmp.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Sink, _ptr_pUnkSink); err != nil {
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

func (o *xxx_ConnectCallbackOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pUnkCP {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_pUnkCP := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConnectionPoint == nil {
				o.ConnectionPoint = &rsmp.Unknown{}
			}
			if err := o.ConnectionPoint.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pUnkCP := func(ptr interface{}) { o.ConnectionPoint = *ptr.(**rsmp.Unknown) }
		if err := w.ReadPointer(&o.ConnectionPoint, _s_pUnkCP, _ptr_pUnkCP); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pUnkSink {in} (1:{pointer=ref}*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_pUnkSink := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Sink == nil {
				o.Sink = &rsmp.Unknown{}
			}
			if err := o.Sink.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pUnkSink := func(ptr interface{}) { o.Sink = *ptr.(**rsmp.Unknown) }
		if err := w.ReadPointer(&o.Sink, _s_pUnkSink, _ptr_pUnkSink); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectCallbackOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectCallbackOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ConnectCallbackOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ConnectCallbackRequest structure represents the ConnectCallback operation request
type ConnectCallbackRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	ConnectionPoint *rsmp.Unknown  `idl:"name:pUnkCP" json:"connection_point"`
	Sink            *rsmp.Unknown  `idl:"name:pUnkSink" json:"sink"`
}

func (o *ConnectCallbackRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectCallbackOperation) *xxx_ConnectCallbackOperation {
	if op == nil {
		op = &xxx_ConnectCallbackOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ConnectionPoint = o.ConnectionPoint
	op.Sink = o.Sink
	return op
}

func (o *ConnectCallbackRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectCallbackOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConnectionPoint = op.ConnectionPoint
	o.Sink = op.Sink
}
func (o *ConnectCallbackRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectCallbackRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectCallbackOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectCallbackResponse structure represents the ConnectCallback operation response
type ConnectCallbackResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ConnectCallback return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ConnectCallbackResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectCallbackOperation) *xxx_ConnectCallbackOperation {
	if op == nil {
		op = &xxx_ConnectCallbackOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ConnectCallbackResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectCallbackOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ConnectCallbackResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectCallbackResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectCallbackOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OnNotifyOperation structure represents the OnNotify operation
type xxx_OnNotifyOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Type      uint32         `idl:"name:dwType" json:"type"`
	Operation uint32         `idl:"name:dwOperation" json:"operation"`
	ID        *dtyp.GUID     `idl:"name:lpIdentifier" json:"id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_OnNotifyOperation) OpNum() int { return 4 }

func (o *xxx_OnNotifyOperation) OpName() string { return "/INtmsNotifySink/v0/OnNotify" }

func (o *xxx_OnNotifyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwOperation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Operation); err != nil {
			return err
		}
	}
	// lpIdentifier {in} (1:{alias=LPGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ID != nil {
			if err := o.ID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwOperation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Operation); err != nil {
			return err
		}
	}
	// lpIdentifier {in} (1:{alias=LPGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ID == nil {
			o.ID = &dtyp.GUID{}
		}
		if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OnNotifyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// OnNotifyRequest structure represents the OnNotify operation request
type OnNotifyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	Type      uint32         `idl:"name:dwType" json:"type"`
	Operation uint32         `idl:"name:dwOperation" json:"operation"`
	ID        *dtyp.GUID     `idl:"name:lpIdentifier" json:"id"`
}

func (o *OnNotifyRequest) xxx_ToOp(ctx context.Context, op *xxx_OnNotifyOperation) *xxx_OnNotifyOperation {
	if op == nil {
		op = &xxx_OnNotifyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Type = o.Type
	op.Operation = o.Operation
	op.ID = o.ID
	return op
}

func (o *OnNotifyRequest) xxx_FromOp(ctx context.Context, op *xxx_OnNotifyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
	o.Operation = op.Operation
	o.ID = op.ID
}
func (o *OnNotifyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OnNotifyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnNotifyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OnNotifyResponse structure represents the OnNotify operation response
type OnNotifyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OnNotify return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OnNotifyResponse) xxx_ToOp(ctx context.Context, op *xxx_OnNotifyOperation) *xxx_OnNotifyOperation {
	if op == nil {
		op = &xxx_OnNotifyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *OnNotifyResponse) xxx_FromOp(ctx context.Context, op *xxx_OnNotifyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *OnNotifyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OnNotifyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnNotifyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReleaseCallbackOperation structure represents the ReleaseCallback operation
type xxx_ReleaseCallbackOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReleaseCallbackOperation) OpNum() int { return 5 }

func (o *xxx_ReleaseCallbackOperation) OpName() string { return "/INtmsNotifySink/v0/ReleaseCallback" }

func (o *xxx_ReleaseCallbackOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReleaseCallbackOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ReleaseCallbackOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ReleaseCallbackOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReleaseCallbackOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ReleaseCallbackOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ReleaseCallbackRequest structure represents the ReleaseCallback operation request
type ReleaseCallbackRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ReleaseCallbackRequest) xxx_ToOp(ctx context.Context, op *xxx_ReleaseCallbackOperation) *xxx_ReleaseCallbackOperation {
	if op == nil {
		op = &xxx_ReleaseCallbackOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ReleaseCallbackRequest) xxx_FromOp(ctx context.Context, op *xxx_ReleaseCallbackOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ReleaseCallbackRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReleaseCallbackRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReleaseCallbackOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReleaseCallbackResponse structure represents the ReleaseCallback operation response
type ReleaseCallbackResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ReleaseCallback return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReleaseCallbackResponse) xxx_ToOp(ctx context.Context, op *xxx_ReleaseCallbackOperation) *xxx_ReleaseCallbackOperation {
	if op == nil {
		op = &xxx_ReleaseCallbackOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ReleaseCallbackResponse) xxx_FromOp(ctx context.Context, op *xxx_ReleaseCallbackOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ReleaseCallbackResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReleaseCallbackResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReleaseCallbackOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
