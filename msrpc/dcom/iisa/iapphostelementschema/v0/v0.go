package iapphostelementschema

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
	// IAppHostElementSchema interface identifier ef13d885-642c-4709-99ec-b89561c6bc69
	AppHostElementSchemaIID = &dcom.IID{Data1: 0xef13d885, Data2: 0x642c, Data3: 0x4709, Data4: []byte{0x99, 0xec, 0xb8, 0x95, 0x61, 0xc6, 0xbc, 0x69}}
	// Syntax UUID
	AppHostElementSchemaSyntaxUUID = &uuid.UUID{TimeLow: 0xef13d885, TimeMid: 0x642c, TimeHiAndVersion: 0x4709, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0xec, Node: [6]uint8{0xb8, 0x95, 0x61, 0xc6, 0xbc, 0x69}}
	// Syntax ID
	AppHostElementSchemaSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostElementSchemaSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostElementSchema interface.
type AppHostElementSchemaClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Name operation.
	GetName(context.Context, *GetNameRequest, ...dcerpc.CallOption) (*GetNameResponse, error)

	// DoesAllowUnschematizedProperties operation.
	GetDoesAllowUnschematizedProperties(context.Context, *GetDoesAllowUnschematizedPropertiesRequest, ...dcerpc.CallOption) (*GetDoesAllowUnschematizedPropertiesResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest, ...dcerpc.CallOption) (*GetMetadataResponse, error)

	// CollectionSchema operation.
	GetCollectionSchema(context.Context, *GetCollectionSchemaRequest, ...dcerpc.CallOption) (*GetCollectionSchemaResponse, error)

	// ChildElementSchemas operation.
	GetChildElementSchemas(context.Context, *GetChildElementSchemasRequest, ...dcerpc.CallOption) (*GetChildElementSchemasResponse, error)

	// PropertySchemas operation.
	GetPropertySchemas(context.Context, *GetPropertySchemasRequest, ...dcerpc.CallOption) (*GetPropertySchemasResponse, error)

	// IsCollectionDefault operation.
	GetIsCollectionDefault(context.Context, *GetIsCollectionDefaultRequest, ...dcerpc.CallOption) (*GetIsCollectionDefaultResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostElementSchemaClient
}

type xxx_DefaultAppHostElementSchemaClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostElementSchemaClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostElementSchemaClient) GetName(ctx context.Context, in *GetNameRequest, opts ...dcerpc.CallOption) (*GetNameResponse, error) {
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
	out := &GetNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementSchemaClient) GetDoesAllowUnschematizedProperties(ctx context.Context, in *GetDoesAllowUnschematizedPropertiesRequest, opts ...dcerpc.CallOption) (*GetDoesAllowUnschematizedPropertiesResponse, error) {
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
	out := &GetDoesAllowUnschematizedPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementSchemaClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...dcerpc.CallOption) (*GetMetadataResponse, error) {
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
	out := &GetMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementSchemaClient) GetCollectionSchema(ctx context.Context, in *GetCollectionSchemaRequest, opts ...dcerpc.CallOption) (*GetCollectionSchemaResponse, error) {
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
	out := &GetCollectionSchemaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementSchemaClient) GetChildElementSchemas(ctx context.Context, in *GetChildElementSchemasRequest, opts ...dcerpc.CallOption) (*GetChildElementSchemasResponse, error) {
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
	out := &GetChildElementSchemasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementSchemaClient) GetPropertySchemas(ctx context.Context, in *GetPropertySchemasRequest, opts ...dcerpc.CallOption) (*GetPropertySchemasResponse, error) {
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
	out := &GetPropertySchemasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementSchemaClient) GetIsCollectionDefault(ctx context.Context, in *GetIsCollectionDefaultRequest, opts ...dcerpc.CallOption) (*GetIsCollectionDefaultResponse, error) {
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
	out := &GetIsCollectionDefaultResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostElementSchemaClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostElementSchemaClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostElementSchemaClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostElementSchemaClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewAppHostElementSchemaClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostElementSchemaClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostElementSchemaSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostElementSchemaClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetNameOperation structure represents the Name operation
type xxx_GetNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:pbstrName" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNameOperation) OpNum() int { return 3 }

func (o *xxx_GetNameOperation) OpName() string { return "/IAppHostElementSchema/v0/Name" }

func (o *xxx_GetNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_pbstrName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_pbstrName); err != nil {
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

func (o *xxx_GetNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrName := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_pbstrName, _ptr_pbstrName); err != nil {
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

// GetNameRequest structure represents the Name operation request
type GetNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetNameOperation {
	if o == nil {
		return &xxx_GetNameOperation{}
	}
	return &xxx_GetNameOperation{
		This: o.This,
	}
}

func (o *GetNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNameResponse structure represents the Name operation response
type GetNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name *oaut.String   `idl:"name:pbstrName" json:"name"`
	// Return: The Name return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetNameOperation {
	if o == nil {
		return &xxx_GetNameOperation{}
	}
	return &xxx_GetNameOperation{
		That:   o.That,
		Name:   o.Name,
		Return: o.Return,
	}
}

func (o *GetNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
}
func (o *GetNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDoesAllowUnschematizedPropertiesOperation structure represents the DoesAllowUnschematizedProperties operation
type xxx_GetDoesAllowUnschematizedPropertiesOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	AllowUnschematized int16          `idl:"name:pfAllowUnschematized" json:"allow_unschematized"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDoesAllowUnschematizedPropertiesOperation) OpNum() int { return 4 }

func (o *xxx_GetDoesAllowUnschematizedPropertiesOperation) OpName() string {
	return "/IAppHostElementSchema/v0/DoesAllowUnschematizedProperties"
}

func (o *xxx_GetDoesAllowUnschematizedPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDoesAllowUnschematizedPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDoesAllowUnschematizedPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDoesAllowUnschematizedPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDoesAllowUnschematizedPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfAllowUnschematized {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.AllowUnschematized); err != nil {
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

func (o *xxx_GetDoesAllowUnschematizedPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfAllowUnschematized {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.AllowUnschematized); err != nil {
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

// GetDoesAllowUnschematizedPropertiesRequest structure represents the DoesAllowUnschematizedProperties operation request
type GetDoesAllowUnschematizedPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDoesAllowUnschematizedPropertiesRequest) xxx_ToOp(ctx context.Context) *xxx_GetDoesAllowUnschematizedPropertiesOperation {
	if o == nil {
		return &xxx_GetDoesAllowUnschematizedPropertiesOperation{}
	}
	return &xxx_GetDoesAllowUnschematizedPropertiesOperation{
		This: o.This,
	}
}

func (o *GetDoesAllowUnschematizedPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDoesAllowUnschematizedPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDoesAllowUnschematizedPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDoesAllowUnschematizedPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDoesAllowUnschematizedPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDoesAllowUnschematizedPropertiesResponse structure represents the DoesAllowUnschematizedProperties operation response
type GetDoesAllowUnschematizedPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	AllowUnschematized int16          `idl:"name:pfAllowUnschematized" json:"allow_unschematized"`
	// Return: The DoesAllowUnschematizedProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDoesAllowUnschematizedPropertiesResponse) xxx_ToOp(ctx context.Context) *xxx_GetDoesAllowUnschematizedPropertiesOperation {
	if o == nil {
		return &xxx_GetDoesAllowUnschematizedPropertiesOperation{}
	}
	return &xxx_GetDoesAllowUnschematizedPropertiesOperation{
		That:               o.That,
		AllowUnschematized: o.AllowUnschematized,
		Return:             o.Return,
	}
}

func (o *GetDoesAllowUnschematizedPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDoesAllowUnschematizedPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AllowUnschematized = op.AllowUnschematized
	o.Return = op.Return
}
func (o *GetDoesAllowUnschematizedPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDoesAllowUnschematizedPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDoesAllowUnschematizedPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMetadataOperation structure represents the GetMetadata operation
type xxx_GetMetadataOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
	Value        *oaut.Variant  `idl:"name:pValue" json:"value"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMetadataOperation) OpNum() int { return 5 }

func (o *xxx_GetMetadataOperation) OpName() string { return "/IAppHostElementSchema/v0/GetMetadata" }

func (o *xxx_GetMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MetadataType != nil {
			_ptr_bstrMetadataType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MetadataType != nil {
					if err := o.MetadataType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MetadataType, _ptr_bstrMetadataType); err != nil {
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

func (o *xxx_GetMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMetadataType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MetadataType == nil {
				o.MetadataType = &oaut.String{}
			}
			if err := o.MetadataType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMetadataType := func(ptr interface{}) { o.MetadataType = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MetadataType, _s_bstrMetadataType, _ptr_bstrMetadataType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Value != nil {
			_ptr_pValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_pValue); err != nil {
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

func (o *xxx_GetMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.Variant{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pValue := func(ptr interface{}) { o.Value = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Value, _s_pValue, _ptr_pValue); err != nil {
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

// GetMetadataRequest structure represents the GetMetadata operation request
type GetMetadataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
}

func (o *GetMetadataRequest) xxx_ToOp(ctx context.Context) *xxx_GetMetadataOperation {
	if o == nil {
		return &xxx_GetMetadataOperation{}
	}
	return &xxx_GetMetadataOperation{
		This:         o.This,
		MetadataType: o.MetadataType,
	}
}

func (o *GetMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMetadataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MetadataType = op.MetadataType
}
func (o *GetMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMetadataResponse structure represents the GetMetadata operation response
type GetMetadataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value *oaut.Variant  `idl:"name:pValue" json:"value"`
	// Return: The GetMetadata return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMetadataResponse) xxx_ToOp(ctx context.Context) *xxx_GetMetadataOperation {
	if o == nil {
		return &xxx_GetMetadataOperation{}
	}
	return &xxx_GetMetadataOperation{
		That:   o.That,
		Value:  o.Value,
		Return: o.Return,
	}
}

func (o *GetMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMetadataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Value = op.Value
	o.Return = op.Return
}
func (o *GetMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCollectionSchemaOperation structure represents the CollectionSchema operation
type xxx_GetCollectionSchemaOperation struct {
	This             *dcom.ORPCThis                `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat                `idl:"name:That" json:"that"`
	CollectionSchema *iisa.AppHostCollectionSchema `idl:"name:ppCollectionSchema" json:"collection_schema"`
	Return           int32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCollectionSchemaOperation) OpNum() int { return 6 }

func (o *xxx_GetCollectionSchemaOperation) OpName() string {
	return "/IAppHostElementSchema/v0/CollectionSchema"
}

func (o *xxx_GetCollectionSchemaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCollectionSchemaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCollectionSchemaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCollectionSchemaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCollectionSchemaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppCollectionSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostCollectionSchema}(interface))
	{
		if o.CollectionSchema != nil {
			_ptr_ppCollectionSchema := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CollectionSchema != nil {
					if err := o.CollectionSchema.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostCollectionSchema{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CollectionSchema, _ptr_ppCollectionSchema); err != nil {
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

func (o *xxx_GetCollectionSchemaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppCollectionSchema {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostCollectionSchema}(interface))
	{
		_ptr_ppCollectionSchema := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CollectionSchema == nil {
				o.CollectionSchema = &iisa.AppHostCollectionSchema{}
			}
			if err := o.CollectionSchema.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppCollectionSchema := func(ptr interface{}) { o.CollectionSchema = *ptr.(**iisa.AppHostCollectionSchema) }
		if err := w.ReadPointer(&o.CollectionSchema, _s_ppCollectionSchema, _ptr_ppCollectionSchema); err != nil {
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

// GetCollectionSchemaRequest structure represents the CollectionSchema operation request
type GetCollectionSchemaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCollectionSchemaRequest) xxx_ToOp(ctx context.Context) *xxx_GetCollectionSchemaOperation {
	if o == nil {
		return &xxx_GetCollectionSchemaOperation{}
	}
	return &xxx_GetCollectionSchemaOperation{
		This: o.This,
	}
}

func (o *GetCollectionSchemaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCollectionSchemaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCollectionSchemaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCollectionSchemaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCollectionSchemaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCollectionSchemaResponse structure represents the CollectionSchema operation response
type GetCollectionSchemaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat                `idl:"name:That" json:"that"`
	CollectionSchema *iisa.AppHostCollectionSchema `idl:"name:ppCollectionSchema" json:"collection_schema"`
	// Return: The CollectionSchema return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCollectionSchemaResponse) xxx_ToOp(ctx context.Context) *xxx_GetCollectionSchemaOperation {
	if o == nil {
		return &xxx_GetCollectionSchemaOperation{}
	}
	return &xxx_GetCollectionSchemaOperation{
		That:             o.That,
		CollectionSchema: o.CollectionSchema,
		Return:           o.Return,
	}
}

func (o *GetCollectionSchemaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCollectionSchemaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CollectionSchema = op.CollectionSchema
	o.Return = op.Return
}
func (o *GetCollectionSchemaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCollectionSchemaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCollectionSchemaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetChildElementSchemasOperation structure represents the ChildElementSchemas operation
type xxx_GetChildElementSchemasOperation struct {
	This         *dcom.ORPCThis                       `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat                       `idl:"name:That" json:"that"`
	ChildSchemas *iisa.AppHostElementSchemaCollection `idl:"name:ppChildSchemas" json:"child_schemas"`
	Return       int32                                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetChildElementSchemasOperation) OpNum() int { return 7 }

func (o *xxx_GetChildElementSchemasOperation) OpName() string {
	return "/IAppHostElementSchema/v0/ChildElementSchemas"
}

func (o *xxx_GetChildElementSchemasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChildElementSchemasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetChildElementSchemasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetChildElementSchemasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChildElementSchemasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppChildSchemas {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElementSchemaCollection}(interface))
	{
		if o.ChildSchemas != nil {
			_ptr_ppChildSchemas := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ChildSchemas != nil {
					if err := o.ChildSchemas.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElementSchemaCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ChildSchemas, _ptr_ppChildSchemas); err != nil {
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

func (o *xxx_GetChildElementSchemasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppChildSchemas {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElementSchemaCollection}(interface))
	{
		_ptr_ppChildSchemas := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ChildSchemas == nil {
				o.ChildSchemas = &iisa.AppHostElementSchemaCollection{}
			}
			if err := o.ChildSchemas.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppChildSchemas := func(ptr interface{}) { o.ChildSchemas = *ptr.(**iisa.AppHostElementSchemaCollection) }
		if err := w.ReadPointer(&o.ChildSchemas, _s_ppChildSchemas, _ptr_ppChildSchemas); err != nil {
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

// GetChildElementSchemasRequest structure represents the ChildElementSchemas operation request
type GetChildElementSchemasRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetChildElementSchemasRequest) xxx_ToOp(ctx context.Context) *xxx_GetChildElementSchemasOperation {
	if o == nil {
		return &xxx_GetChildElementSchemasOperation{}
	}
	return &xxx_GetChildElementSchemasOperation{
		This: o.This,
	}
}

func (o *GetChildElementSchemasRequest) xxx_FromOp(ctx context.Context, op *xxx_GetChildElementSchemasOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetChildElementSchemasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetChildElementSchemasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetChildElementSchemasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetChildElementSchemasResponse structure represents the ChildElementSchemas operation response
type GetChildElementSchemasResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat                       `idl:"name:That" json:"that"`
	ChildSchemas *iisa.AppHostElementSchemaCollection `idl:"name:ppChildSchemas" json:"child_schemas"`
	// Return: The ChildElementSchemas return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetChildElementSchemasResponse) xxx_ToOp(ctx context.Context) *xxx_GetChildElementSchemasOperation {
	if o == nil {
		return &xxx_GetChildElementSchemasOperation{}
	}
	return &xxx_GetChildElementSchemasOperation{
		That:         o.That,
		ChildSchemas: o.ChildSchemas,
		Return:       o.Return,
	}
}

func (o *GetChildElementSchemasResponse) xxx_FromOp(ctx context.Context, op *xxx_GetChildElementSchemasOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ChildSchemas = op.ChildSchemas
	o.Return = op.Return
}
func (o *GetChildElementSchemasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetChildElementSchemasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetChildElementSchemasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertySchemasOperation structure represents the PropertySchemas operation
type xxx_GetPropertySchemasOperation struct {
	This            *dcom.ORPCThis                        `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat                        `idl:"name:That" json:"that"`
	PropertySchemas *iisa.AppHostPropertySchemaCollection `idl:"name:ppPropertySchemas" json:"property_schemas"`
	Return          int32                                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertySchemasOperation) OpNum() int { return 8 }

func (o *xxx_GetPropertySchemasOperation) OpName() string {
	return "/IAppHostElementSchema/v0/PropertySchemas"
}

func (o *xxx_GetPropertySchemasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertySchemasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertySchemasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertySchemasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertySchemasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppPropertySchemas {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostPropertySchemaCollection}(interface))
	{
		if o.PropertySchemas != nil {
			_ptr_ppPropertySchemas := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertySchemas != nil {
					if err := o.PropertySchemas.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostPropertySchemaCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertySchemas, _ptr_ppPropertySchemas); err != nil {
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

func (o *xxx_GetPropertySchemasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppPropertySchemas {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostPropertySchemaCollection}(interface))
	{
		_ptr_ppPropertySchemas := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertySchemas == nil {
				o.PropertySchemas = &iisa.AppHostPropertySchemaCollection{}
			}
			if err := o.PropertySchemas.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppPropertySchemas := func(ptr interface{}) { o.PropertySchemas = *ptr.(**iisa.AppHostPropertySchemaCollection) }
		if err := w.ReadPointer(&o.PropertySchemas, _s_ppPropertySchemas, _ptr_ppPropertySchemas); err != nil {
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

// GetPropertySchemasRequest structure represents the PropertySchemas operation request
type GetPropertySchemasRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertySchemasRequest) xxx_ToOp(ctx context.Context) *xxx_GetPropertySchemasOperation {
	if o == nil {
		return &xxx_GetPropertySchemasOperation{}
	}
	return &xxx_GetPropertySchemasOperation{
		This: o.This,
	}
}

func (o *GetPropertySchemasRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertySchemasOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertySchemasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertySchemasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertySchemasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertySchemasResponse structure represents the PropertySchemas operation response
type GetPropertySchemasResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat                        `idl:"name:That" json:"that"`
	PropertySchemas *iisa.AppHostPropertySchemaCollection `idl:"name:ppPropertySchemas" json:"property_schemas"`
	// Return: The PropertySchemas return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertySchemasResponse) xxx_ToOp(ctx context.Context) *xxx_GetPropertySchemasOperation {
	if o == nil {
		return &xxx_GetPropertySchemasOperation{}
	}
	return &xxx_GetPropertySchemasOperation{
		That:            o.That,
		PropertySchemas: o.PropertySchemas,
		Return:          o.Return,
	}
}

func (o *GetPropertySchemasResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertySchemasOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertySchemas = op.PropertySchemas
	o.Return = op.Return
}
func (o *GetPropertySchemasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertySchemasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertySchemasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsCollectionDefaultOperation structure represents the IsCollectionDefault operation
type xxx_GetIsCollectionDefaultOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsCollectionDefault int16          `idl:"name:pfIsCollectionDefault" json:"is_collection_default"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsCollectionDefaultOperation) OpNum() int { return 9 }

func (o *xxx_GetIsCollectionDefaultOperation) OpName() string {
	return "/IAppHostElementSchema/v0/IsCollectionDefault"
}

func (o *xxx_GetIsCollectionDefaultOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsCollectionDefaultOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsCollectionDefaultOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsCollectionDefaultOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsCollectionDefaultOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfIsCollectionDefault {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.IsCollectionDefault); err != nil {
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

func (o *xxx_GetIsCollectionDefaultOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfIsCollectionDefault {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.IsCollectionDefault); err != nil {
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

// GetIsCollectionDefaultRequest structure represents the IsCollectionDefault operation request
type GetIsCollectionDefaultRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsCollectionDefaultRequest) xxx_ToOp(ctx context.Context) *xxx_GetIsCollectionDefaultOperation {
	if o == nil {
		return &xxx_GetIsCollectionDefaultOperation{}
	}
	return &xxx_GetIsCollectionDefaultOperation{
		This: o.This,
	}
}

func (o *GetIsCollectionDefaultRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsCollectionDefaultOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsCollectionDefaultRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetIsCollectionDefaultRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsCollectionDefaultOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsCollectionDefaultResponse structure represents the IsCollectionDefault operation response
type GetIsCollectionDefaultResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsCollectionDefault int16          `idl:"name:pfIsCollectionDefault" json:"is_collection_default"`
	// Return: The IsCollectionDefault return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsCollectionDefaultResponse) xxx_ToOp(ctx context.Context) *xxx_GetIsCollectionDefaultOperation {
	if o == nil {
		return &xxx_GetIsCollectionDefaultOperation{}
	}
	return &xxx_GetIsCollectionDefaultOperation{
		That:                o.That,
		IsCollectionDefault: o.IsCollectionDefault,
		Return:              o.Return,
	}
}

func (o *GetIsCollectionDefaultResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsCollectionDefaultOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsCollectionDefault = op.IsCollectionDefault
	o.Return = op.Return
}
func (o *GetIsCollectionDefaultResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetIsCollectionDefaultResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsCollectionDefaultOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
