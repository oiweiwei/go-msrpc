package imsmqqueueinfos

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
	// IMSMQQueueInfos interface identifier d7d6e07d-dccd-11d0-aa4b-0060970debae
	ImsmqQueueInfosIID = &dcom.IID{Data1: 0xd7d6e07d, Data2: 0xdccd, Data3: 0x11d0, Data4: []byte{0xaa, 0x4b, 0x00, 0x60, 0x97, 0x0d, 0xeb, 0xae}}
	// Syntax UUID
	ImsmqQueueInfosSyntaxUUID = &uuid.UUID{TimeLow: 0xd7d6e07d, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	// Syntax ID
	ImsmqQueueInfosSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ImsmqQueueInfosSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQQueueInfos interface.
type ImsmqQueueInfosClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// Reset operation.
	Reset(context.Context, *ResetRequest, ...dcerpc.CallOption) (*ResetResponse, error)

	// Next operation.
	Next(context.Context, *NextRequest, ...dcerpc.CallOption) (*NextResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ImsmqQueueInfosClient
}

type xxx_DefaultImsmqQueueInfosClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultImsmqQueueInfosClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultImsmqQueueInfosClient) Reset(ctx context.Context, in *ResetRequest, opts ...dcerpc.CallOption) (*ResetResponse, error) {
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
	out := &ResetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqQueueInfosClient) Next(ctx context.Context, in *NextRequest, opts ...dcerpc.CallOption) (*NextResponse, error) {
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
	out := &NextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqQueueInfosClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultImsmqQueueInfosClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultImsmqQueueInfosClient) IPID(ctx context.Context, ipid *dcom.IPID) ImsmqQueueInfosClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultImsmqQueueInfosClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewImsmqQueueInfosClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ImsmqQueueInfosClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ImsmqQueueInfosSyntaxV0_0))...)
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
	return &xxx_DefaultImsmqQueueInfosClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_ResetOperation structure represents the Reset operation
type xxx_ResetOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResetOperation) OpNum() int { return 7 }

func (o *xxx_ResetOperation) OpName() string { return "/IMSMQQueueInfos/v0/Reset" }

func (o *xxx_ResetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ResetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ResetRequest structure represents the Reset operation request
type ResetRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ResetRequest) xxx_ToOp(ctx context.Context, op *xxx_ResetOperation) *xxx_ResetOperation {
	if op == nil {
		op = &xxx_ResetOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ResetRequest) xxx_FromOp(ctx context.Context, op *xxx_ResetOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ResetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResetResponse structure represents the Reset operation response
type ResetResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Reset return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResetResponse) xxx_ToOp(ctx context.Context, op *xxx_ResetOperation) *xxx_ResetOperation {
	if op == nil {
		op = &xxx_ResetOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ResetResponse) xxx_FromOp(ctx context.Context, op *xxx_ResetOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ResetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NextOperation structure represents the Next operation
type xxx_NextOperation struct {
	This        *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat       `idl:"name:That" json:"that"`
	PpqinfoNext *mqac.ImsmqQueueInfo `idl:"name:ppqinfoNext" json:"ppqinfo_next"`
	Return      int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_NextOperation) OpNum() int { return 8 }

func (o *xxx_NextOperation) OpName() string { return "/IMSMQQueueInfos/v0/Next" }

func (o *xxx_NextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_NextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_NextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfoNext {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		if o.PpqinfoNext != nil {
			_ptr_ppqinfoNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PpqinfoNext != nil {
					if err := o.PpqinfoNext.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.ImsmqQueueInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PpqinfoNext, _ptr_ppqinfoNext); err != nil {
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

func (o *xxx_NextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfoNext {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		_ptr_ppqinfoNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PpqinfoNext == nil {
				o.PpqinfoNext = &mqac.ImsmqQueueInfo{}
			}
			if err := o.PpqinfoNext.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfoNext := func(ptr interface{}) { o.PpqinfoNext = *ptr.(**mqac.ImsmqQueueInfo) }
		if err := w.ReadPointer(&o.PpqinfoNext, _s_ppqinfoNext, _ptr_ppqinfoNext); err != nil {
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

// NextRequest structure represents the Next operation request
type NextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *NextRequest) xxx_ToOp(ctx context.Context, op *xxx_NextOperation) *xxx_NextOperation {
	if op == nil {
		op = &xxx_NextOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *NextRequest) xxx_FromOp(ctx context.Context, op *xxx_NextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *NextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *NextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NextResponse structure represents the Next operation response
type NextResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat       `idl:"name:That" json:"that"`
	PpqinfoNext *mqac.ImsmqQueueInfo `idl:"name:ppqinfoNext" json:"ppqinfo_next"`
	// Return: The Next return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *NextResponse) xxx_ToOp(ctx context.Context, op *xxx_NextOperation) *xxx_NextOperation {
	if op == nil {
		op = &xxx_NextOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PpqinfoNext = o.PpqinfoNext
	op.Return = o.Return
	return op
}

func (o *NextResponse) xxx_FromOp(ctx context.Context, op *xxx_NextOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PpqinfoNext = op.PpqinfoNext
	o.Return = op.Return
}
func (o *NextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *NextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
