package ifsrmpropertydefinition2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	ifsrmpropertydefinition "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpropertydefinition/v0"
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
	_ = ifsrmpropertydefinition.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmPropertyDefinition2 interface identifier 47782152-d16c-4229-b4e1-0ddfe308b9f6
	PropertyDefinition2IID = &dcom.IID{Data1: 0x47782152, Data2: 0xd16c, Data3: 0x4229, Data4: []byte{0xb4, 0xe1, 0x0d, 0xdf, 0xe3, 0x08, 0xb9, 0xf6}}
	// Syntax UUID
	PropertyDefinition2SyntaxUUID = &uuid.UUID{TimeLow: 0x47782152, TimeMid: 0xd16c, TimeHiAndVersion: 0x4229, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0xe1, Node: [6]uint8{0xd, 0xdf, 0xe3, 0x8, 0xb9, 0xf6}}
	// Syntax ID
	PropertyDefinition2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: PropertyDefinition2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmPropertyDefinition2 interface.
type PropertyDefinition2Client interface {

	// IFsrmPropertyDefinition retrieval method.
	PropertyDefinition() ifsrmpropertydefinition.PropertyDefinitionClient

	// The PropertyDefinitionFlags (get) method returns the Property Definition.Global,
	// Property Definition.Secure, and Property Definition.Deprecated values for the object.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+------------------------------------------------+
	//	|         RETURN          |                                                |
	//	|       VALUE/CODE        |                  DESCRIPTION                   |
	//	|                         |                                                |
	//	+-------------------------+------------------------------------------------+
	//	+-------------------------+------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The propertyDefinitionFlags parameter is NULL. |
	//	+-------------------------+------------------------------------------------+
	GetPropertyDefinitionFlags(context.Context, *GetPropertyDefinitionFlagsRequest, ...dcerpc.CallOption) (*GetPropertyDefinitionFlagsResponse, error)

	// DisplayName operation.
	GetDisplayName(context.Context, *GetDisplayNameRequest, ...dcerpc.CallOption) (*GetDisplayNameResponse, error)

	// DisplayName operation.
	SetDisplayName(context.Context, *SetDisplayNameRequest, ...dcerpc.CallOption) (*SetDisplayNameResponse, error)

	// AppliesTo operation.
	GetAppliesTo(context.Context, *GetAppliesToRequest, ...dcerpc.CallOption) (*GetAppliesToResponse, error)

	// The ValueDefinitions (get) method returns the property definitions Possible Values.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------------------+
	//	|         RETURN          |                                         |
	//	|       VALUE/CODE        |               DESCRIPTION               |
	//	|                         |                                         |
	//	+-------------------------+-----------------------------------------+
	//	+-------------------------+-----------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The valueDefinitions parameter is NULL. |
	//	+-------------------------+-----------------------------------------+
	GetValueDefinitions(context.Context, *GetValueDefinitionsRequest, ...dcerpc.CallOption) (*GetValueDefinitionsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) PropertyDefinition2Client
}

type xxx_DefaultPropertyDefinition2Client struct {
	ifsrmpropertydefinition.PropertyDefinitionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPropertyDefinition2Client) PropertyDefinition() ifsrmpropertydefinition.PropertyDefinitionClient {
	return o.PropertyDefinitionClient
}

func (o *xxx_DefaultPropertyDefinition2Client) GetPropertyDefinitionFlags(ctx context.Context, in *GetPropertyDefinitionFlagsRequest, opts ...dcerpc.CallOption) (*GetPropertyDefinitionFlagsResponse, error) {
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
	out := &GetPropertyDefinitionFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinition2Client) GetDisplayName(ctx context.Context, in *GetDisplayNameRequest, opts ...dcerpc.CallOption) (*GetDisplayNameResponse, error) {
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
	out := &GetDisplayNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinition2Client) SetDisplayName(ctx context.Context, in *SetDisplayNameRequest, opts ...dcerpc.CallOption) (*SetDisplayNameResponse, error) {
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
	out := &SetDisplayNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinition2Client) GetAppliesTo(ctx context.Context, in *GetAppliesToRequest, opts ...dcerpc.CallOption) (*GetAppliesToResponse, error) {
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
	out := &GetAppliesToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinition2Client) GetValueDefinitions(ctx context.Context, in *GetValueDefinitionsRequest, opts ...dcerpc.CallOption) (*GetValueDefinitionsResponse, error) {
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
	out := &GetValueDefinitionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinition2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPropertyDefinition2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultPropertyDefinition2Client) IPID(ctx context.Context, ipid *dcom.IPID) PropertyDefinition2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPropertyDefinition2Client{
		PropertyDefinitionClient: o.PropertyDefinitionClient.IPID(ctx, ipid),
		cc:                       o.cc,
		ipid:                     ipid,
	}
}

func NewPropertyDefinition2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PropertyDefinition2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PropertyDefinition2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmpropertydefinition.NewPropertyDefinitionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultPropertyDefinition2Client{
		PropertyDefinitionClient: base,
		cc:                       cc,
		ipid:                     ipid,
	}, nil
}

// xxx_GetPropertyDefinitionFlagsOperation structure represents the PropertyDefinitionFlags operation
type xxx_GetPropertyDefinitionFlagsOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	PropertyDefinitionFlags int32          `idl:"name:propertyDefinitionFlags" json:"property_definition_flags"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertyDefinitionFlagsOperation) OpNum() int { return 22 }

func (o *xxx_GetPropertyDefinitionFlagsOperation) OpName() string {
	return "/IFsrmPropertyDefinition2/v0/PropertyDefinitionFlags"
}

func (o *xxx_GetPropertyDefinitionFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyDefinitionFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertyDefinitionFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertyDefinitionFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyDefinitionFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// propertyDefinitionFlags {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.PropertyDefinitionFlags); err != nil {
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

func (o *xxx_GetPropertyDefinitionFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// propertyDefinitionFlags {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.PropertyDefinitionFlags); err != nil {
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

// GetPropertyDefinitionFlagsRequest structure represents the PropertyDefinitionFlags operation request
type GetPropertyDefinitionFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertyDefinitionFlagsRequest) xxx_ToOp(ctx context.Context) *xxx_GetPropertyDefinitionFlagsOperation {
	if o == nil {
		return &xxx_GetPropertyDefinitionFlagsOperation{}
	}
	return &xxx_GetPropertyDefinitionFlagsOperation{
		This: o.This,
	}
}

func (o *GetPropertyDefinitionFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertyDefinitionFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertyDefinitionFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertyDefinitionFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertyDefinitionFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertyDefinitionFlagsResponse structure represents the PropertyDefinitionFlags operation response
type GetPropertyDefinitionFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// propertyDefinitionFlags: Pointer to a variable that upon completion contains Property
	// Definition.Global, Property Definition.Secure, and Property Definition.Deprecated
	// values for the object.
	PropertyDefinitionFlags int32 `idl:"name:propertyDefinitionFlags" json:"property_definition_flags"`
	// Return: The PropertyDefinitionFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertyDefinitionFlagsResponse) xxx_ToOp(ctx context.Context) *xxx_GetPropertyDefinitionFlagsOperation {
	if o == nil {
		return &xxx_GetPropertyDefinitionFlagsOperation{}
	}
	return &xxx_GetPropertyDefinitionFlagsOperation{
		That:                    o.That,
		PropertyDefinitionFlags: o.PropertyDefinitionFlags,
		Return:                  o.Return,
	}
}

func (o *GetPropertyDefinitionFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertyDefinitionFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyDefinitionFlags = op.PropertyDefinitionFlags
	o.Return = op.Return
}
func (o *GetPropertyDefinitionFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertyDefinitionFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertyDefinitionFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDisplayNameOperation structure represents the DisplayName operation
type xxx_GetDisplayNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDisplayNameOperation) OpNum() int { return 23 }

func (o *xxx_GetDisplayNameOperation) OpName() string {
	return "/IFsrmPropertyDefinition2/v0/DisplayName"
}

func (o *xxx_GetDisplayNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDisplayNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDisplayNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// name {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
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

func (o *xxx_GetDisplayNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// name {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_name := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
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

// GetDisplayNameRequest structure represents the DisplayName operation request
type GetDisplayNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDisplayNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetDisplayNameOperation {
	if o == nil {
		return &xxx_GetDisplayNameOperation{}
	}
	return &xxx_GetDisplayNameOperation{
		This: o.This,
	}
}

func (o *GetDisplayNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDisplayNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDisplayNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDisplayNameResponse structure represents the DisplayName operation response
type GetDisplayNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name *oaut.String   `idl:"name:name" json:"name"`
	// Return: The DisplayName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDisplayNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetDisplayNameOperation {
	if o == nil {
		return &xxx_GetDisplayNameOperation{}
	}
	return &xxx_GetDisplayNameOperation{
		That:   o.That,
		Name:   o.Name,
		Return: o.Return,
	}
}

func (o *GetDisplayNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
}
func (o *GetDisplayNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDisplayNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDisplayNameOperation structure represents the DisplayName operation
type xxx_SetDisplayNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDisplayNameOperation) OpNum() int { return 24 }

func (o *xxx_SetDisplayNameOperation) OpName() string {
	return "/IFsrmPropertyDefinition2/v0/DisplayName"
}

func (o *xxx_SetDisplayNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDisplayNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
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

func (o *xxx_SetDisplayNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_name := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDisplayNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDisplayNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDisplayNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDisplayNameRequest structure represents the DisplayName operation request
type SetDisplayNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Name *oaut.String   `idl:"name:name" json:"name"`
}

func (o *SetDisplayNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetDisplayNameOperation {
	if o == nil {
		return &xxx_SetDisplayNameOperation{}
	}
	return &xxx_SetDisplayNameOperation{
		This: o.This,
		Name: o.Name,
	}
}

func (o *SetDisplayNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDisplayNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *SetDisplayNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDisplayNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDisplayNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDisplayNameResponse structure represents the DisplayName operation response
type SetDisplayNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DisplayName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDisplayNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetDisplayNameOperation {
	if o == nil {
		return &xxx_SetDisplayNameOperation{}
	}
	return &xxx_SetDisplayNameOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetDisplayNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDisplayNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDisplayNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDisplayNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDisplayNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAppliesToOperation structure represents the AppliesTo operation
type xxx_GetAppliesToOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	AppliesTo int32          `idl:"name:appliesTo" json:"applies_to"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAppliesToOperation) OpNum() int { return 25 }

func (o *xxx_GetAppliesToOperation) OpName() string { return "/IFsrmPropertyDefinition2/v0/AppliesTo" }

func (o *xxx_GetAppliesToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAppliesToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAppliesToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAppliesToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAppliesToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// appliesTo {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.AppliesTo); err != nil {
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

func (o *xxx_GetAppliesToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// appliesTo {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.AppliesTo); err != nil {
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

// GetAppliesToRequest structure represents the AppliesTo operation request
type GetAppliesToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAppliesToRequest) xxx_ToOp(ctx context.Context) *xxx_GetAppliesToOperation {
	if o == nil {
		return &xxx_GetAppliesToOperation{}
	}
	return &xxx_GetAppliesToOperation{
		This: o.This,
	}
}

func (o *GetAppliesToRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAppliesToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAppliesToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAppliesToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAppliesToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAppliesToResponse structure represents the AppliesTo operation response
type GetAppliesToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	AppliesTo int32          `idl:"name:appliesTo" json:"applies_to"`
	// Return: The AppliesTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAppliesToResponse) xxx_ToOp(ctx context.Context) *xxx_GetAppliesToOperation {
	if o == nil {
		return &xxx_GetAppliesToOperation{}
	}
	return &xxx_GetAppliesToOperation{
		That:      o.That,
		AppliesTo: o.AppliesTo,
		Return:    o.Return,
	}
}

func (o *GetAppliesToResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAppliesToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AppliesTo = op.AppliesTo
	o.Return = op.Return
}
func (o *GetAppliesToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAppliesToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAppliesToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetValueDefinitionsOperation structure represents the ValueDefinitions operation
type xxx_GetValueDefinitionsOperation struct {
	This             *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat   `idl:"name:That" json:"that"`
	ValueDefinitions *fsrm.Collection `idl:"name:valueDefinitions" json:"value_definitions"`
	Return           int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetValueDefinitionsOperation) OpNum() int { return 26 }

func (o *xxx_GetValueDefinitionsOperation) OpName() string {
	return "/IFsrmPropertyDefinition2/v0/ValueDefinitions"
}

func (o *xxx_GetValueDefinitionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValueDefinitionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetValueDefinitionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetValueDefinitionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValueDefinitionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// valueDefinitions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.ValueDefinitions != nil {
			_ptr_valueDefinitions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ValueDefinitions != nil {
					if err := o.ValueDefinitions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ValueDefinitions, _ptr_valueDefinitions); err != nil {
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

func (o *xxx_GetValueDefinitionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// valueDefinitions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_valueDefinitions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ValueDefinitions == nil {
				o.ValueDefinitions = &fsrm.Collection{}
			}
			if err := o.ValueDefinitions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_valueDefinitions := func(ptr interface{}) { o.ValueDefinitions = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.ValueDefinitions, _s_valueDefinitions, _ptr_valueDefinitions); err != nil {
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

// GetValueDefinitionsRequest structure represents the ValueDefinitions operation request
type GetValueDefinitionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetValueDefinitionsRequest) xxx_ToOp(ctx context.Context) *xxx_GetValueDefinitionsOperation {
	if o == nil {
		return &xxx_GetValueDefinitionsOperation{}
	}
	return &xxx_GetValueDefinitionsOperation{
		This: o.This,
	}
}

func (o *GetValueDefinitionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetValueDefinitionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetValueDefinitionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetValueDefinitionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValueDefinitionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetValueDefinitionsResponse structure represents the ValueDefinitions operation response
type GetValueDefinitionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// valueDefinitions: Pointer to a variable that, upon completion, contains the array
	// of IFsrmPropertyDefinitionValue elements as defined in the property definition.
	ValueDefinitions *fsrm.Collection `idl:"name:valueDefinitions" json:"value_definitions"`
	// Return: The ValueDefinitions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetValueDefinitionsResponse) xxx_ToOp(ctx context.Context) *xxx_GetValueDefinitionsOperation {
	if o == nil {
		return &xxx_GetValueDefinitionsOperation{}
	}
	return &xxx_GetValueDefinitionsOperation{
		That:             o.That,
		ValueDefinitions: o.ValueDefinitions,
		Return:           o.Return,
	}
}

func (o *GetValueDefinitionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetValueDefinitionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ValueDefinitions = op.ValueDefinitions
	o.Return = op.Return
}
func (o *GetValueDefinitionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetValueDefinitionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValueDefinitionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
