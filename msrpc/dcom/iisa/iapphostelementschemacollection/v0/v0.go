package iapphostelementschemacollection

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
	// IAppHostElementSchemaCollection interface identifier 0344cdda-151e-4cbf-82da-66ae61e97754
	AppHostElementSchemaCollectionIID = &dcom.IID{Data1: 0x0344cdda, Data2: 0x151e, Data3: 0x4cbf, Data4: []byte{0x82, 0xda, 0x66, 0xae, 0x61, 0xe9, 0x77, 0x54}}
	// Syntax UUID
	AppHostElementSchemaCollectionSyntaxUUID = &uuid.UUID{TimeLow: 0x344cdda, TimeMid: 0x151e, TimeHiAndVersion: 0x4cbf, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0xda, Node: [6]uint8{0x66, 0xae, 0x61, 0xe9, 0x77, 0x54}}
	// Syntax ID
	AppHostElementSchemaCollectionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostElementSchemaCollectionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostElementSchemaCollection interface.
type AppHostElementSchemaCollectionClient interface {

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
	IPID(context.Context, *dcom.IPID) AppHostElementSchemaCollectionClient
}

type xxx_DefaultAppHostElementSchemaCollectionClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostElementSchemaCollectionClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostElementSchemaCollectionClient) GetCount(ctx context.Context, in *GetCountRequest, opts ...dcerpc.CallOption) (*GetCountResponse, error) {
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

func (o *xxx_DefaultAppHostElementSchemaCollectionClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...dcerpc.CallOption) (*GetItemResponse, error) {
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

func (o *xxx_DefaultAppHostElementSchemaCollectionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostElementSchemaCollectionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostElementSchemaCollectionClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostElementSchemaCollectionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostElementSchemaCollectionClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAppHostElementSchemaCollectionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostElementSchemaCollectionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostElementSchemaCollectionSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostElementSchemaCollectionClient{
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

func (o *xxx_GetCountOperation) OpName() string { return "/IAppHostElementSchemaCollection/v0/Count" }

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
	ElementSchema *iisa.AppHostElementSchema `idl:"name:ppElementSchema" json:"element_schema"`
	Return        int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetItemOperation) OpNum() int { return 4 }

func (o *xxx_GetItemOperation) OpName() string { return "/IAppHostElementSchemaCollection/v0/Item" }

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
	// ppElementSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElementSchema}(interface))
	{
		if o.ElementSchema != nil {
			_ptr_ppElementSchema := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ElementSchema != nil {
					if err := o.ElementSchema.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElementSchema{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ElementSchema, _ptr_ppElementSchema); err != nil {
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
	// ppElementSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElementSchema}(interface))
	{
		_ptr_ppElementSchema := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ElementSchema == nil {
				o.ElementSchema = &iisa.AppHostElementSchema{}
			}
			if err := o.ElementSchema.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppElementSchema := func(ptr interface{}) { o.ElementSchema = *ptr.(**iisa.AppHostElementSchema) }
		if err := w.ReadPointer(&o.ElementSchema, _s_ppElementSchema, _ptr_ppElementSchema); err != nil {
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
	ElementSchema *iisa.AppHostElementSchema `idl:"name:ppElementSchema" json:"element_schema"`
	// Return: The Item return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetItemResponse) xxx_ToOp(ctx context.Context) *xxx_GetItemOperation {
	if o == nil {
		return &xxx_GetItemOperation{}
	}
	return &xxx_GetItemOperation{
		That:          o.That,
		ElementSchema: o.ElementSchema,
		Return:        o.Return,
	}
}

func (o *GetItemResponse) xxx_FromOp(ctx context.Context, op *xxx_GetItemOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ElementSchema = op.ElementSchema
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
