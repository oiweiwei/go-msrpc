package iapphostelementcollection

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
	// IAppHostElementCollection interface identifier c8550bff-5281-4b1e-ac34-99b6fa38464d
	AppHostElementCollectionIID = &dcom.IID{Data1: 0xc8550bff, Data2: 0x5281, Data3: 0x4b1e, Data4: []byte{0xac, 0x34, 0x99, 0xb6, 0xfa, 0x38, 0x46, 0x4d}}
	// Syntax UUID
	AppHostElementCollectionSyntaxUUID = &uuid.UUID{TimeLow: 0xc8550bff, TimeMid: 0x5281, TimeHiAndVersion: 0x4b1e, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0x34, Node: [6]uint8{0x99, 0xb6, 0xfa, 0x38, 0x46, 0x4d}}
	// Syntax ID
	AppHostElementCollectionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostElementCollectionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostElementCollection interface.
type AppHostElementCollectionClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Count operation.
	GetCount(context.Context, *GetCountRequest, ...dcerpc.CallOption) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest, ...dcerpc.CallOption) (*GetItemResponse, error)

	// AddElement operation.
	AddElement(context.Context, *AddElementRequest, ...dcerpc.CallOption) (*AddElementResponse, error)

	// DeleteElement operation.
	DeleteElement(context.Context, *DeleteElementRequest, ...dcerpc.CallOption) (*DeleteElementResponse, error)

	// Clear operation.
	Clear(context.Context, *ClearRequest, ...dcerpc.CallOption) (*ClearResponse, error)

	// CreateNewElement operation.
	CreateNewElement(context.Context, *CreateNewElementRequest, ...dcerpc.CallOption) (*CreateNewElementResponse, error)

	// Schema operation.
	GetSchema(context.Context, *GetSchemaRequest, ...dcerpc.CallOption) (*GetSchemaResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostElementCollectionClient
}

type xxx_DefaultAppHostElementCollectionClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostElementCollectionClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostElementCollectionClient) GetCount(ctx context.Context, in *GetCountRequest, opts ...dcerpc.CallOption) (*GetCountResponse, error) {
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

func (o *xxx_DefaultAppHostElementCollectionClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...dcerpc.CallOption) (*GetItemResponse, error) {
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

func (o *xxx_DefaultAppHostElementCollectionClient) AddElement(ctx context.Context, in *AddElementRequest, opts ...dcerpc.CallOption) (*AddElementResponse, error) {
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
	out := &AddElementResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementCollectionClient) DeleteElement(ctx context.Context, in *DeleteElementRequest, opts ...dcerpc.CallOption) (*DeleteElementResponse, error) {
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
	out := &DeleteElementResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementCollectionClient) Clear(ctx context.Context, in *ClearRequest, opts ...dcerpc.CallOption) (*ClearResponse, error) {
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
	out := &ClearResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementCollectionClient) CreateNewElement(ctx context.Context, in *CreateNewElementRequest, opts ...dcerpc.CallOption) (*CreateNewElementResponse, error) {
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
	out := &CreateNewElementResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementCollectionClient) GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...dcerpc.CallOption) (*GetSchemaResponse, error) {
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
	out := &GetSchemaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementCollectionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostElementCollectionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostElementCollectionClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostElementCollectionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostElementCollectionClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAppHostElementCollectionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostElementCollectionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostElementCollectionSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostElementCollectionClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetCountOperation structure represents the Count operation
type xxx_GetCountOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ElementCount uint32         `idl:"name:pcElementCount" json:"element_count"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCountOperation) OpNum() int { return 3 }

func (o *xxx_GetCountOperation) OpName() string { return "/IAppHostElementCollection/v0/Count" }

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
	// pcElementCount {out, retval} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ElementCount); err != nil {
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
	// pcElementCount {out, retval} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ElementCount); err != nil {
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
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ElementCount uint32         `idl:"name:pcElementCount" json:"element_count"`
	// Return: The Count return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCountResponse) xxx_ToOp(ctx context.Context) *xxx_GetCountOperation {
	if o == nil {
		return &xxx_GetCountOperation{}
	}
	return &xxx_GetCountOperation{
		That:         o.That,
		ElementCount: o.ElementCount,
		Return:       o.Return,
	}
}

func (o *GetCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ElementCount = op.ElementCount
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
	This    *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Index   *oaut.Variant        `idl:"name:cIndex" json:"index"`
	Element *iisa.AppHostElement `idl:"name:ppElement" json:"element"`
	Return  int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetItemOperation) OpNum() int { return 4 }

func (o *xxx_GetItemOperation) OpName() string { return "/IAppHostElementCollection/v0/Item" }

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
	// ppElement {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		if o.Element != nil {
			_ptr_ppElement := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Element != nil {
					if err := o.Element.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Element, _ptr_ppElement); err != nil {
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
	// ppElement {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		_ptr_ppElement := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Element == nil {
				o.Element = &iisa.AppHostElement{}
			}
			if err := o.Element.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppElement := func(ptr interface{}) { o.Element = *ptr.(**iisa.AppHostElement) }
		if err := w.ReadPointer(&o.Element, _s_ppElement, _ptr_ppElement); err != nil {
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
	That    *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Element *iisa.AppHostElement `idl:"name:ppElement" json:"element"`
	// Return: The Item return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetItemResponse) xxx_ToOp(ctx context.Context) *xxx_GetItemOperation {
	if o == nil {
		return &xxx_GetItemOperation{}
	}
	return &xxx_GetItemOperation{
		That:    o.That,
		Element: o.Element,
		Return:  o.Return,
	}
}

func (o *GetItemResponse) xxx_FromOp(ctx context.Context, op *xxx_GetItemOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Element = op.Element
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

// xxx_AddElementOperation structure represents the AddElement operation
type xxx_AddElementOperation struct {
	This          *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Element       *iisa.AppHostElement `idl:"name:pElement" json:"element"`
	PositionCount int32                `idl:"name:cPosition" json:"position_count"`
	Return        int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_AddElementOperation) OpNum() int { return 5 }

func (o *xxx_AddElementOperation) OpName() string { return "/IAppHostElementCollection/v0/AddElement" }

func (o *xxx_AddElementOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddElementOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pElement {in} (1:{pointer=ref}*(1))(2:{alias=IAppHostElement}(interface))
	{
		if o.Element != nil {
			_ptr_pElement := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Element != nil {
					if err := o.Element.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Element, _ptr_pElement); err != nil {
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
	// cPosition {in, default_value={-1}} (1:(int32))
	{
		if err := w.WriteData(o.PositionCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddElementOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pElement {in} (1:{pointer=ref}*(1))(2:{alias=IAppHostElement}(interface))
	{
		_ptr_pElement := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Element == nil {
				o.Element = &iisa.AppHostElement{}
			}
			if err := o.Element.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pElement := func(ptr interface{}) { o.Element = *ptr.(**iisa.AppHostElement) }
		if err := w.ReadPointer(&o.Element, _s_pElement, _ptr_pElement); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cPosition {in, default_value={-1}} (1:(int32))
	{
		if err := w.ReadData(&o.PositionCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddElementOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddElementOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddElementOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddElementRequest structure represents the AddElement operation request
type AddElementRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis       `idl:"name:This" json:"this"`
	Element       *iisa.AppHostElement `idl:"name:pElement" json:"element"`
	PositionCount int32                `idl:"name:cPosition" json:"position_count"`
}

func (o *AddElementRequest) xxx_ToOp(ctx context.Context) *xxx_AddElementOperation {
	if o == nil {
		return &xxx_AddElementOperation{}
	}
	return &xxx_AddElementOperation{
		This:          o.This,
		Element:       o.Element,
		PositionCount: o.PositionCount,
	}
}

func (o *AddElementRequest) xxx_FromOp(ctx context.Context, op *xxx_AddElementOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Element = op.Element
	o.PositionCount = op.PositionCount
}
func (o *AddElementRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddElementRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddElementOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddElementResponse structure represents the AddElement operation response
type AddElementResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddElement return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddElementResponse) xxx_ToOp(ctx context.Context) *xxx_AddElementOperation {
	if o == nil {
		return &xxx_AddElementOperation{}
	}
	return &xxx_AddElementOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AddElementResponse) xxx_FromOp(ctx context.Context, op *xxx_AddElementOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddElementResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddElementResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddElementOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteElementOperation structure represents the DeleteElement operation
type xxx_DeleteElementOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index  *oaut.Variant  `idl:"name:cIndex" json:"index"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteElementOperation) OpNum() int { return 6 }

func (o *xxx_DeleteElementOperation) OpName() string {
	return "/IAppHostElementCollection/v0/DeleteElement"
}

func (o *xxx_DeleteElementOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteElementOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteElementOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_DeleteElementOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteElementOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteElementOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteElementRequest structure represents the DeleteElement operation request
type DeleteElementRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Index *oaut.Variant  `idl:"name:cIndex" json:"index"`
}

func (o *DeleteElementRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteElementOperation {
	if o == nil {
		return &xxx_DeleteElementOperation{}
	}
	return &xxx_DeleteElementOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *DeleteElementRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteElementOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *DeleteElementRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteElementRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteElementOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteElementResponse structure represents the DeleteElement operation response
type DeleteElementResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteElement return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteElementResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteElementOperation {
	if o == nil {
		return &xxx_DeleteElementOperation{}
	}
	return &xxx_DeleteElementOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteElementResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteElementOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteElementResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteElementResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteElementOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearOperation structure represents the Clear operation
type xxx_ClearOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearOperation) OpNum() int { return 7 }

func (o *xxx_ClearOperation) OpName() string { return "/IAppHostElementCollection/v0/Clear" }

func (o *xxx_ClearOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ClearOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ClearRequest structure represents the Clear operation request
type ClearRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ClearRequest) xxx_ToOp(ctx context.Context) *xxx_ClearOperation {
	if o == nil {
		return &xxx_ClearOperation{}
	}
	return &xxx_ClearOperation{
		This: o.This,
	}
}

func (o *ClearRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ClearRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClearRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearResponse structure represents the Clear operation response
type ClearResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Clear return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearResponse) xxx_ToOp(ctx context.Context) *xxx_ClearOperation {
	if o == nil {
		return &xxx_ClearOperation{}
	}
	return &xxx_ClearOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ClearResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ClearResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClearResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateNewElementOperation structure represents the CreateNewElement operation
type xxx_CreateNewElementOperation struct {
	This        *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat       `idl:"name:That" json:"that"`
	ElementName *oaut.String         `idl:"name:bstrElementName" json:"element_name"`
	Element     *iisa.AppHostElement `idl:"name:ppElement" json:"element"`
	Return      int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateNewElementOperation) OpNum() int { return 8 }

func (o *xxx_CreateNewElementOperation) OpName() string {
	return "/IAppHostElementCollection/v0/CreateNewElement"
}

func (o *xxx_CreateNewElementOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNewElementOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrElementName {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ElementName != nil {
			_ptr_bstrElementName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ElementName != nil {
					if err := o.ElementName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ElementName, _ptr_bstrElementName); err != nil {
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

func (o *xxx_CreateNewElementOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrElementName {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrElementName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ElementName == nil {
				o.ElementName = &oaut.String{}
			}
			if err := o.ElementName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrElementName := func(ptr interface{}) { o.ElementName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ElementName, _s_bstrElementName, _ptr_bstrElementName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNewElementOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNewElementOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppElement {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		if o.Element != nil {
			_ptr_ppElement := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Element != nil {
					if err := o.Element.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Element, _ptr_ppElement); err != nil {
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

func (o *xxx_CreateNewElementOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppElement {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		_ptr_ppElement := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Element == nil {
				o.Element = &iisa.AppHostElement{}
			}
			if err := o.Element.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppElement := func(ptr interface{}) { o.Element = *ptr.(**iisa.AppHostElement) }
		if err := w.ReadPointer(&o.Element, _s_ppElement, _ptr_ppElement); err != nil {
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

// CreateNewElementRequest structure represents the CreateNewElement operation request
type CreateNewElementRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	ElementName *oaut.String   `idl:"name:bstrElementName" json:"element_name"`
}

func (o *CreateNewElementRequest) xxx_ToOp(ctx context.Context) *xxx_CreateNewElementOperation {
	if o == nil {
		return &xxx_CreateNewElementOperation{}
	}
	return &xxx_CreateNewElementOperation{
		This:        o.This,
		ElementName: o.ElementName,
	}
}

func (o *CreateNewElementRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateNewElementOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ElementName = op.ElementName
}
func (o *CreateNewElementRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateNewElementRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNewElementOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateNewElementResponse structure represents the CreateNewElement operation response
type CreateNewElementResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Element *iisa.AppHostElement `idl:"name:ppElement" json:"element"`
	// Return: The CreateNewElement return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateNewElementResponse) xxx_ToOp(ctx context.Context) *xxx_CreateNewElementOperation {
	if o == nil {
		return &xxx_CreateNewElementOperation{}
	}
	return &xxx_CreateNewElementOperation{
		That:    o.That,
		Element: o.Element,
		Return:  o.Return,
	}
}

func (o *CreateNewElementResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateNewElementOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Element = op.Element
	o.Return = op.Return
}
func (o *CreateNewElementResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateNewElementResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNewElementOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSchemaOperation structure represents the Schema operation
type xxx_GetSchemaOperation struct {
	This   *dcom.ORPCThis                `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat                `idl:"name:That" json:"that"`
	Schema *iisa.AppHostCollectionSchema `idl:"name:ppSchema" json:"schema"`
	Return int32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSchemaOperation) OpNum() int { return 9 }

func (o *xxx_GetSchemaOperation) OpName() string { return "/IAppHostElementCollection/v0/Schema" }

func (o *xxx_GetSchemaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSchemaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSchemaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSchemaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSchemaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostCollectionSchema}(interface))
	{
		if o.Schema != nil {
			_ptr_ppSchema := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Schema != nil {
					if err := o.Schema.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostCollectionSchema{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Schema, _ptr_ppSchema); err != nil {
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

func (o *xxx_GetSchemaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostCollectionSchema}(interface))
	{
		_ptr_ppSchema := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Schema == nil {
				o.Schema = &iisa.AppHostCollectionSchema{}
			}
			if err := o.Schema.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSchema := func(ptr interface{}) { o.Schema = *ptr.(**iisa.AppHostCollectionSchema) }
		if err := w.ReadPointer(&o.Schema, _s_ppSchema, _ptr_ppSchema); err != nil {
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

// GetSchemaRequest structure represents the Schema operation request
type GetSchemaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSchemaRequest) xxx_ToOp(ctx context.Context) *xxx_GetSchemaOperation {
	if o == nil {
		return &xxx_GetSchemaOperation{}
	}
	return &xxx_GetSchemaOperation{
		This: o.This,
	}
}

func (o *GetSchemaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSchemaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSchemaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSchemaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSchemaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSchemaResponse structure represents the Schema operation response
type GetSchemaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat                `idl:"name:That" json:"that"`
	Schema *iisa.AppHostCollectionSchema `idl:"name:ppSchema" json:"schema"`
	// Return: The Schema return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSchemaResponse) xxx_ToOp(ctx context.Context) *xxx_GetSchemaOperation {
	if o == nil {
		return &xxx_GetSchemaOperation{}
	}
	return &xxx_GetSchemaOperation{
		That:   o.That,
		Schema: o.Schema,
		Return: o.Return,
	}
}

func (o *GetSchemaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSchemaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Schema = op.Schema
	o.Return = op.Return
}
func (o *GetSchemaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSchemaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSchemaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
