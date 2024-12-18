package iapphostconstantvaluecollection

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iisa "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = oaut.GoPackage
	_ = iisa.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/iisa"
)

var (
	// IAppHostConstantValueCollection interface identifier 5b5a68e6-8b9f-45e1-8199-a95ffccdffff
	AppHostConstantValueCollectionIID = &dcom.IID{Data1: 0x5b5a68e6, Data2: 0x8b9f, Data3: 0x45e1, Data4: []byte{0x81, 0x99, 0xa9, 0x5f, 0xfc, 0xcd, 0xff, 0xff}}
	// Syntax UUID
	AppHostConstantValueCollectionSyntaxUUID = &uuid.UUID{TimeLow: 0x5b5a68e6, TimeMid: 0x8b9f, TimeHiAndVersion: 0x45e1, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0x99, Node: [6]uint8{0xa9, 0x5f, 0xfc, 0xcd, 0xff, 0xff}}
	// Syntax ID
	AppHostConstantValueCollectionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostConstantValueCollectionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostConstantValueCollection interface.
type AppHostConstantValueCollectionClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Count operation.
	GetCount(context.Context, *GetCountRequest, ...dcerpc.CallOption) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest, ...dcerpc.CallOption) (*GetItemResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostConstantValueCollectionClient
}

type xxx_DefaultAppHostConstantValueCollectionClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostConstantValueCollectionClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostConstantValueCollectionClient) GetCount(ctx context.Context, in *GetCountRequest, opts ...dcerpc.CallOption) (*GetCountResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConstantValueCollectionClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...dcerpc.CallOption) (*GetItemResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetItemResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConstantValueCollectionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostConstantValueCollectionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostConstantValueCollectionClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostConstantValueCollectionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostConstantValueCollectionClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAppHostConstantValueCollectionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostConstantValueCollectionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostConstantValueCollectionSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostConstantValueCollectionClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetCountOperation structure represents the Count operation
type xxx_GetCountOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Count  uint32         `idl:"name:pcCount" json:"count"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCountOperation) OpNum() int { return 3 }

func (o *xxx_GetCountOperation) OpName() string { return "/IAppHostConstantValueCollection/v0/Count" }

func (o *xxx_GetCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcCount {out, retval} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
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

func (o *xxx_GetCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcCount {out, retval} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
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

// GetCountRequest structure represents the Count operation request
type GetCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCountRequest) xxx_ToOp(ctx context.Context) *xxx_GetCountOperation {
	if o == nil {
		return &xxx_GetCountOperation{}
	}
	return &xxx_GetCountOperation{
		This: o.This,
	}
}

func (o *GetCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCountResponse structure represents the Count operation response
type GetCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Count uint32         `idl:"name:pcCount" json:"count"`
	// Return: The Count return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCountResponse) xxx_ToOp(ctx context.Context) *xxx_GetCountOperation {
	if o == nil {
		return &xxx_GetCountOperation{}
	}
	return &xxx_GetCountOperation{
		That:   o.That,
		Count:  o.Count,
		Return: o.Return,
	}
}

func (o *GetCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Count = op.Count
	o.Return = op.Return
}
func (o *GetCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetItemOperation structure represents the Item operation
type xxx_GetItemOperation struct {
	This          *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Index         *oaut.Variant              `idl:"name:cIndex" json:"index"`
	ConstantValue *iisa.AppHostConstantValue `idl:"name:ppConstantValue" json:"constant_value"`
	Return        int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetItemOperation) OpNum() int { return 4 }

func (o *xxx_GetItemOperation) OpName() string { return "/IAppHostConstantValueCollection/v0/Item" }

func (o *xxx_GetItemOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// cIndex {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Index != nil {
			if err := o.Index.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// cIndex {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Index == nil {
			o.Index = &oaut.Variant{}
		}
		if err := o.Index.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppConstantValue {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostConstantValue}(interface))
	{
		if o.ConstantValue != nil {
			_ptr_ppConstantValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConstantValue != nil {
					if err := o.ConstantValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostConstantValue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConstantValue, _ptr_ppConstantValue); err != nil {
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

func (o *xxx_GetItemOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppConstantValue {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostConstantValue}(interface))
	{
		_ptr_ppConstantValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConstantValue == nil {
				o.ConstantValue = &iisa.AppHostConstantValue{}
			}
			if err := o.ConstantValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppConstantValue := func(ptr interface{}) { o.ConstantValue = *ptr.(**iisa.AppHostConstantValue) }
		if err := w.ReadPointer(&o.ConstantValue, _s_ppConstantValue, _ptr_ppConstantValue); err != nil {
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

// GetItemRequest structure represents the Item operation request
type GetItemRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Index *oaut.Variant  `idl:"name:cIndex" json:"index"`
}

func (o *GetItemRequest) xxx_ToOp(ctx context.Context) *xxx_GetItemOperation {
	if o == nil {
		return &xxx_GetItemOperation{}
	}
	return &xxx_GetItemOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *GetItemRequest) xxx_FromOp(ctx context.Context, op *xxx_GetItemOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetItemRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetItemRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetItemOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetItemResponse structure represents the Item operation response
type GetItemResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat             `idl:"name:That" json:"that"`
	ConstantValue *iisa.AppHostConstantValue `idl:"name:ppConstantValue" json:"constant_value"`
	// Return: The Item return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetItemResponse) xxx_ToOp(ctx context.Context) *xxx_GetItemOperation {
	if o == nil {
		return &xxx_GetItemOperation{}
	}
	return &xxx_GetItemOperation{
		That:          o.That,
		ConstantValue: o.ConstantValue,
		Return:        o.Return,
	}
}

func (o *GetItemResponse) xxx_FromOp(ctx context.Context, op *xxx_GetItemOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ConstantValue = op.ConstantValue
	o.Return = op.Return
}
func (o *GetItemResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetItemResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetItemOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
