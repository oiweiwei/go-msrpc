package iupdatesearcher2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iupdatesearcher "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesearcher/v0"
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
	_ = iupdatesearcher.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdateSearcher2 interface identifier 4cbdcb2d-1589-4beb-bd1c-3e582ff0add0
	UpdateSearcher2IID = &dcom.IID{Data1: 0x4cbdcb2d, Data2: 0x1589, Data3: 0x4beb, Data4: []byte{0xbd, 0x1c, 0x3e, 0x58, 0x2f, 0xf0, 0xad, 0xd0}}
	// Syntax UUID
	UpdateSearcher2SyntaxUUID = &uuid.UUID{TimeLow: 0x4cbdcb2d, TimeMid: 0x1589, TimeHiAndVersion: 0x4beb, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0x1c, Node: [6]uint8{0x3e, 0x58, 0x2f, 0xf0, 0xad, 0xd0}}
	// Syntax ID
	UpdateSearcher2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UpdateSearcher2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdateSearcher2 interface.
type UpdateSearcher2Client interface {

	// IUpdateSearcher retrieval method.
	UpdateSearcher() iupdatesearcher.UpdateSearcherClient

	GetIgnoreDownloadPriority(context.Context, *GetIgnoreDownloadPriorityRequest, ...dcerpc.CallOption) (*GetIgnoreDownloadPriorityResponse, error)

	SetIgnoreDownloadPriority(context.Context, *SetIgnoreDownloadPriorityRequest, ...dcerpc.CallOption) (*SetIgnoreDownloadPriorityResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UpdateSearcher2Client
}

type xxx_DefaultUpdateSearcher2Client struct {
	iupdatesearcher.UpdateSearcherClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdateSearcher2Client) UpdateSearcher() iupdatesearcher.UpdateSearcherClient {
	return o.UpdateSearcherClient
}

func (o *xxx_DefaultUpdateSearcher2Client) GetIgnoreDownloadPriority(ctx context.Context, in *GetIgnoreDownloadPriorityRequest, opts ...dcerpc.CallOption) (*GetIgnoreDownloadPriorityResponse, error) {
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
	out := &GetIgnoreDownloadPriorityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcher2Client) SetIgnoreDownloadPriority(ctx context.Context, in *SetIgnoreDownloadPriorityRequest, opts ...dcerpc.CallOption) (*SetIgnoreDownloadPriorityResponse, error) {
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
	out := &SetIgnoreDownloadPriorityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcher2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdateSearcher2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdateSearcher2Client) IPID(ctx context.Context, ipid *dcom.IPID) UpdateSearcher2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdateSearcher2Client{
		UpdateSearcherClient: o.UpdateSearcherClient.IPID(ctx, ipid),
		cc:                   o.cc,
		ipid:                 ipid,
	}
}

func NewUpdateSearcher2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UpdateSearcher2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UpdateSearcher2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iupdatesearcher.NewUpdateSearcherClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdateSearcher2Client{
		UpdateSearcherClient: base,
		cc:                   cc,
		ipid:                 ipid,
	}, nil
}

// xxx_GetIgnoreDownloadPriorityOperation structure represents the IgnoreDownloadPriority operation
type xxx_GetIgnoreDownloadPriorityOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIgnoreDownloadPriorityOperation) OpNum() int { return 25 }

func (o *xxx_GetIgnoreDownloadPriorityOperation) OpName() string {
	return "/IUpdateSearcher2/v0/IgnoreDownloadPriority"
}

func (o *xxx_GetIgnoreDownloadPriorityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIgnoreDownloadPriorityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIgnoreDownloadPriorityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIgnoreDownloadPriorityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIgnoreDownloadPriorityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
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

func (o *xxx_GetIgnoreDownloadPriorityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
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

// GetIgnoreDownloadPriorityRequest structure represents the IgnoreDownloadPriority operation request
type GetIgnoreDownloadPriorityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIgnoreDownloadPriorityRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIgnoreDownloadPriorityOperation) *xxx_GetIgnoreDownloadPriorityOperation {
	if op == nil {
		op = &xxx_GetIgnoreDownloadPriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIgnoreDownloadPriorityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIgnoreDownloadPriorityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIgnoreDownloadPriorityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIgnoreDownloadPriorityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIgnoreDownloadPriorityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIgnoreDownloadPriorityResponse structure represents the IgnoreDownloadPriority operation response
type GetIgnoreDownloadPriorityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	// Return: The IgnoreDownloadPriority return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIgnoreDownloadPriorityResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIgnoreDownloadPriorityOperation) *xxx_GetIgnoreDownloadPriorityOperation {
	if op == nil {
		op = &xxx_GetIgnoreDownloadPriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetIgnoreDownloadPriorityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIgnoreDownloadPriorityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetIgnoreDownloadPriorityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIgnoreDownloadPriorityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIgnoreDownloadPriorityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetIgnoreDownloadPriorityOperation structure represents the IgnoreDownloadPriority operation
type xxx_SetIgnoreDownloadPriorityOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value  int16          `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetIgnoreDownloadPriorityOperation) OpNum() int { return 26 }

func (o *xxx_SetIgnoreDownloadPriorityOperation) OpName() string {
	return "/IUpdateSearcher2/v0/IgnoreDownloadPriority"
}

func (o *xxx_SetIgnoreDownloadPriorityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIgnoreDownloadPriorityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// value {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Value); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIgnoreDownloadPriorityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// value {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Value); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIgnoreDownloadPriorityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIgnoreDownloadPriorityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetIgnoreDownloadPriorityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetIgnoreDownloadPriorityRequest structure represents the IgnoreDownloadPriority operation request
type SetIgnoreDownloadPriorityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Value int16          `idl:"name:value" json:"value"`
}

func (o *SetIgnoreDownloadPriorityRequest) xxx_ToOp(ctx context.Context, op *xxx_SetIgnoreDownloadPriorityOperation) *xxx_SetIgnoreDownloadPriorityOperation {
	if op == nil {
		op = &xxx_SetIgnoreDownloadPriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Value = o.Value
	return op
}

func (o *SetIgnoreDownloadPriorityRequest) xxx_FromOp(ctx context.Context, op *xxx_SetIgnoreDownloadPriorityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Value = op.Value
}
func (o *SetIgnoreDownloadPriorityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetIgnoreDownloadPriorityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetIgnoreDownloadPriorityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetIgnoreDownloadPriorityResponse structure represents the IgnoreDownloadPriority operation response
type SetIgnoreDownloadPriorityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IgnoreDownloadPriority return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetIgnoreDownloadPriorityResponse) xxx_ToOp(ctx context.Context, op *xxx_SetIgnoreDownloadPriorityOperation) *xxx_SetIgnoreDownloadPriorityOperation {
	if op == nil {
		op = &xxx_SetIgnoreDownloadPriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetIgnoreDownloadPriorityResponse) xxx_FromOp(ctx context.Context, op *xxx_SetIgnoreDownloadPriorityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetIgnoreDownloadPriorityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetIgnoreDownloadPriorityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetIgnoreDownloadPriorityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
