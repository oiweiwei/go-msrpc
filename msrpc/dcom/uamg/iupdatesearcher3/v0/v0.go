package iupdatesearcher3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	uamg "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg"
	iupdatesearcher2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesearcher2/v0"
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
	_ = iupdatesearcher2.GoPackage
	_ = uamg.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdateSearcher3 interface identifier 04c6895d-eaf2-4034-97f3-311de9be413a
	UpdateSearcher3IID = &dcom.IID{Data1: 0x04c6895d, Data2: 0xeaf2, Data3: 0x4034, Data4: []byte{0x97, 0xf3, 0x31, 0x1d, 0xe9, 0xbe, 0x41, 0x3a}}
	// Syntax UUID
	UpdateSearcher3SyntaxUUID = &uuid.UUID{TimeLow: 0x4c6895d, TimeMid: 0xeaf2, TimeHiAndVersion: 0x4034, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0xf3, Node: [6]uint8{0x31, 0x1d, 0xe9, 0xbe, 0x41, 0x3a}}
	// Syntax ID
	UpdateSearcher3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UpdateSearcher3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdateSearcher3 interface.
type UpdateSearcher3Client interface {

	// IUpdateSearcher2 retrieval method.
	UpdateSearcher2() iupdatesearcher2.UpdateSearcher2Client

	GetSearchScope(context.Context, *GetSearchScopeRequest, ...dcerpc.CallOption) (*GetSearchScopeResponse, error)

	SetSearchScope(context.Context, *SetSearchScopeRequest, ...dcerpc.CallOption) (*SetSearchScopeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UpdateSearcher3Client
}

type xxx_DefaultUpdateSearcher3Client struct {
	iupdatesearcher2.UpdateSearcher2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdateSearcher3Client) UpdateSearcher2() iupdatesearcher2.UpdateSearcher2Client {
	return o.UpdateSearcher2Client
}

func (o *xxx_DefaultUpdateSearcher3Client) GetSearchScope(ctx context.Context, in *GetSearchScopeRequest, opts ...dcerpc.CallOption) (*GetSearchScopeResponse, error) {
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
	out := &GetSearchScopeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcher3Client) SetSearchScope(ctx context.Context, in *SetSearchScopeRequest, opts ...dcerpc.CallOption) (*SetSearchScopeResponse, error) {
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
	out := &SetSearchScopeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcher3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdateSearcher3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdateSearcher3Client) IPID(ctx context.Context, ipid *dcom.IPID) UpdateSearcher3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdateSearcher3Client{
		UpdateSearcher2Client: o.UpdateSearcher2Client.IPID(ctx, ipid),
		cc:                    o.cc,
		ipid:                  ipid,
	}
}

func NewUpdateSearcher3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UpdateSearcher3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UpdateSearcher3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iupdatesearcher2.NewUpdateSearcher2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdateSearcher3Client{
		UpdateSearcher2Client: base,
		cc:                    cc,
		ipid:                  ipid,
	}, nil
}

// xxx_GetSearchScopeOperation structure represents the SearchScope operation
type xxx_GetSearchScopeOperation struct {
	This        *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat   `idl:"name:That" json:"that"`
	ReturnValue uamg.SearchScope `idl:"name:retval" json:"return_value"`
	Return      int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSearchScopeOperation) OpNum() int { return 27 }

func (o *xxx_GetSearchScopeOperation) OpName() string { return "/IUpdateSearcher3/v0/SearchScope" }

func (o *xxx_GetSearchScopeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSearchScopeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSearchScopeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSearchScopeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSearchScopeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{v1_enum, alias=SearchScope}(enum))
	{
		if err := w.WriteEnum(uint32(o.ReturnValue)); err != nil {
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

func (o *xxx_GetSearchScopeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{v1_enum, alias=SearchScope}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.ReturnValue)); err != nil {
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

// GetSearchScopeRequest structure represents the SearchScope operation request
type GetSearchScopeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSearchScopeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSearchScopeOperation) *xxx_GetSearchScopeOperation {
	if op == nil {
		op = &xxx_GetSearchScopeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSearchScopeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSearchScopeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSearchScopeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSearchScopeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSearchScopeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSearchScopeResponse structure represents the SearchScope operation response
type GetSearchScopeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat   `idl:"name:That" json:"that"`
	ReturnValue uamg.SearchScope `idl:"name:retval" json:"return_value"`
	// Return: The SearchScope return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSearchScopeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSearchScopeOperation) *xxx_GetSearchScopeOperation {
	if op == nil {
		op = &xxx_GetSearchScopeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetSearchScopeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSearchScopeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetSearchScopeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSearchScopeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSearchScopeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSearchScopeOperation structure represents the SearchScope operation
type xxx_SetSearchScopeOperation struct {
	This   *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Value  uamg.SearchScope `idl:"name:value" json:"value"`
	Return int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSearchScopeOperation) OpNum() int { return 28 }

func (o *xxx_SetSearchScopeOperation) OpName() string { return "/IUpdateSearcher3/v0/SearchScope" }

func (o *xxx_SetSearchScopeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSearchScopeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// value {in} (1:{v1_enum, alias=SearchScope}(enum))
	{
		if err := w.WriteEnum(uint32(o.Value)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSearchScopeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// value {in} (1:{v1_enum, alias=SearchScope}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.Value)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSearchScopeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSearchScopeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSearchScopeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSearchScopeRequest structure represents the SearchScope operation request
type SetSearchScopeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis   `idl:"name:This" json:"this"`
	Value uamg.SearchScope `idl:"name:value" json:"value"`
}

func (o *SetSearchScopeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSearchScopeOperation) *xxx_SetSearchScopeOperation {
	if op == nil {
		op = &xxx_SetSearchScopeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Value = o.Value
	return op
}

func (o *SetSearchScopeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSearchScopeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Value = op.Value
}
func (o *SetSearchScopeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSearchScopeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSearchScopeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSearchScopeResponse structure represents the SearchScope operation response
type SetSearchScopeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SearchScope return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSearchScopeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSearchScopeOperation) *xxx_SetSearchScopeOperation {
	if op == nil {
		op = &xxx_SetSearchScopeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSearchScopeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSearchScopeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSearchScopeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSearchScopeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSearchScopeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
