package ifsrmpropertydefinition

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
	ifsrmobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmobject/v0"
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
	_ = ifsrmobject.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmPropertyDefinition interface identifier ede0150f-e9a3-419c-877c-01fe5d24c5d3
	PropertyDefinitionIID = &dcom.IID{Data1: 0xede0150f, Data2: 0xe9a3, Data3: 0x419c, Data4: []byte{0x87, 0x7c, 0x01, 0xfe, 0x5d, 0x24, 0xc5, 0xd3}}
	// Syntax UUID
	PropertyDefinitionSyntaxUUID = &uuid.UUID{TimeLow: 0xede0150f, TimeMid: 0xe9a3, TimeHiAndVersion: 0x419c, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0x7c, Node: [6]uint8{0x1, 0xfe, 0x5d, 0x24, 0xc5, 0xd3}}
	// Syntax ID
	PropertyDefinitionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: PropertyDefinitionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmPropertyDefinition interface.
type PropertyDefinitionClient interface {

	// IFsrmObject retrieval method.
	Object() ifsrmobject.ObjectClient

	// Name operation.
	GetName(context.Context, *GetNameRequest, ...dcerpc.CallOption) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest, ...dcerpc.CallOption) (*SetNameResponse, error)

	// Type operation.
	GetType(context.Context, *GetTypeRequest, ...dcerpc.CallOption) (*GetTypeResponse, error)

	// Type operation.
	SetType(context.Context, *SetTypeRequest, ...dcerpc.CallOption) (*SetTypeResponse, error)

	// PossibleValues operation.
	GetPossibleValues(context.Context, *GetPossibleValuesRequest, ...dcerpc.CallOption) (*GetPossibleValuesResponse, error)

	// PossibleValues operation.
	SetPossibleValues(context.Context, *SetPossibleValuesRequest, ...dcerpc.CallOption) (*SetPossibleValuesResponse, error)

	// ValueDescriptions operation.
	GetValueDescriptions(context.Context, *GetValueDescriptionsRequest, ...dcerpc.CallOption) (*GetValueDescriptionsResponse, error)

	// ValueDescriptions operation.
	SetValueDescriptions(context.Context, *SetValueDescriptionsRequest, ...dcerpc.CallOption) (*SetValueDescriptionsResponse, error)

	// Parameters operation.
	GetParameters(context.Context, *GetParametersRequest, ...dcerpc.CallOption) (*GetParametersResponse, error)

	// Parameters operation.
	SetParameters(context.Context, *SetParametersRequest, ...dcerpc.CallOption) (*SetParametersResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) PropertyDefinitionClient
}

type xxx_DefaultPropertyDefinitionClient struct {
	ifsrmobject.ObjectClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPropertyDefinitionClient) Object() ifsrmobject.ObjectClient {
	return o.ObjectClient
}

func (o *xxx_DefaultPropertyDefinitionClient) GetName(ctx context.Context, in *GetNameRequest, opts ...dcerpc.CallOption) (*GetNameResponse, error) {
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

func (o *xxx_DefaultPropertyDefinitionClient) SetName(ctx context.Context, in *SetNameRequest, opts ...dcerpc.CallOption) (*SetNameResponse, error) {
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
	out := &SetNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinitionClient) GetType(ctx context.Context, in *GetTypeRequest, opts ...dcerpc.CallOption) (*GetTypeResponse, error) {
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
	out := &GetTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinitionClient) SetType(ctx context.Context, in *SetTypeRequest, opts ...dcerpc.CallOption) (*SetTypeResponse, error) {
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
	out := &SetTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinitionClient) GetPossibleValues(ctx context.Context, in *GetPossibleValuesRequest, opts ...dcerpc.CallOption) (*GetPossibleValuesResponse, error) {
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
	out := &GetPossibleValuesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinitionClient) SetPossibleValues(ctx context.Context, in *SetPossibleValuesRequest, opts ...dcerpc.CallOption) (*SetPossibleValuesResponse, error) {
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
	out := &SetPossibleValuesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinitionClient) GetValueDescriptions(ctx context.Context, in *GetValueDescriptionsRequest, opts ...dcerpc.CallOption) (*GetValueDescriptionsResponse, error) {
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
	out := &GetValueDescriptionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinitionClient) SetValueDescriptions(ctx context.Context, in *SetValueDescriptionsRequest, opts ...dcerpc.CallOption) (*SetValueDescriptionsResponse, error) {
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
	out := &SetValueDescriptionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinitionClient) GetParameters(ctx context.Context, in *GetParametersRequest, opts ...dcerpc.CallOption) (*GetParametersResponse, error) {
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
	out := &GetParametersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinitionClient) SetParameters(ctx context.Context, in *SetParametersRequest, opts ...dcerpc.CallOption) (*SetParametersResponse, error) {
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
	out := &SetParametersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPropertyDefinitionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPropertyDefinitionClient) IPID(ctx context.Context, ipid *dcom.IPID) PropertyDefinitionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPropertyDefinitionClient{
		ObjectClient: o.ObjectClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}
func NewPropertyDefinitionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PropertyDefinitionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PropertyDefinitionSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmobject.NewObjectClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultPropertyDefinitionClient{
		ObjectClient: base,
		cc:           cc,
		ipid:         ipid,
	}, nil
}

// xxx_GetNameOperation structure represents the Name operation
type xxx_GetNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNameOperation) OpNum() int { return 12 }

func (o *xxx_GetNameOperation) OpName() string { return "/IFsrmPropertyDefinition/v0/Name" }

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
	Name *oaut.String   `idl:"name:name" json:"name"`
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

// xxx_SetNameOperation structure represents the Name operation
type xxx_SetNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNameOperation) OpNum() int { return 13 }

func (o *xxx_SetNameOperation) OpName() string { return "/IFsrmPropertyDefinition/v0/Name" }

func (o *xxx_SetNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_SetNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNameRequest structure represents the Name operation request
type SetNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Name *oaut.String   `idl:"name:name" json:"name"`
}

func (o *SetNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetNameOperation {
	if o == nil {
		return &xxx_SetNameOperation{}
	}
	return &xxx_SetNameOperation{
		This: o.This,
		Name: o.Name,
	}
}

func (o *SetNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *SetNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNameResponse structure represents the Name operation response
type SetNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Name return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetNameOperation {
	if o == nil {
		return &xxx_SetNameOperation{}
	}
	return &xxx_SetNameOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTypeOperation structure represents the Type operation
type xxx_GetTypeOperation struct {
	This   *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Type   fsrm.PropertyDefinitionType `idl:"name:type" json:"type"`
	Return int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeOperation) OpNum() int { return 14 }

func (o *xxx_GetTypeOperation) OpName() string { return "/IFsrmPropertyDefinition/v0/Type" }

func (o *xxx_GetTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// type {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmPropertyDefinitionType}(enum))
	{
		if err := w.WriteData(uint16(o.Type)); err != nil {
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

func (o *xxx_GetTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// type {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmPropertyDefinitionType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Type)); err != nil {
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

// GetTypeRequest structure represents the Type operation request
type GetTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTypeRequest) xxx_ToOp(ctx context.Context) *xxx_GetTypeOperation {
	if o == nil {
		return &xxx_GetTypeOperation{}
	}
	return &xxx_GetTypeOperation{
		This: o.This,
	}
}

func (o *GetTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeResponse structure represents the Type operation response
type GetTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Type fsrm.PropertyDefinitionType `idl:"name:type" json:"type"`
	// Return: The Type return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeResponse) xxx_ToOp(ctx context.Context) *xxx_GetTypeOperation {
	if o == nil {
		return &xxx_GetTypeOperation{}
	}
	return &xxx_GetTypeOperation{
		That:   o.That,
		Type:   o.Type,
		Return: o.Return,
	}
}

func (o *GetTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Type = op.Type
	o.Return = op.Return
}
func (o *GetTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetTypeOperation structure represents the Type operation
type xxx_SetTypeOperation struct {
	This   *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Type   fsrm.PropertyDefinitionType `idl:"name:type" json:"type"`
	Return int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTypeOperation) OpNum() int { return 15 }

func (o *xxx_SetTypeOperation) OpName() string { return "/IFsrmPropertyDefinition/v0/Type" }

func (o *xxx_SetTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// type {in} (1:{alias=FsrmPropertyDefinitionType}(enum))
	{
		if err := w.WriteData(uint16(o.Type)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// type {in} (1:{alias=FsrmPropertyDefinitionType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Type)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetTypeRequest structure represents the Type operation request
type SetTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis              `idl:"name:This" json:"this"`
	Type fsrm.PropertyDefinitionType `idl:"name:type" json:"type"`
}

func (o *SetTypeRequest) xxx_ToOp(ctx context.Context) *xxx_SetTypeOperation {
	if o == nil {
		return &xxx_SetTypeOperation{}
	}
	return &xxx_SetTypeOperation{
		This: o.This,
		Type: o.Type,
	}
}

func (o *SetTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
}
func (o *SetTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTypeResponse structure represents the Type operation response
type SetTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Type return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetTypeResponse) xxx_ToOp(ctx context.Context) *xxx_SetTypeOperation {
	if o == nil {
		return &xxx_SetTypeOperation{}
	}
	return &xxx_SetTypeOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPossibleValuesOperation structure represents the PossibleValues operation
type xxx_GetPossibleValuesOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	PossibleValues *oaut.SafeArray `idl:"name:possibleValues" json:"possible_values"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPossibleValuesOperation) OpNum() int { return 16 }

func (o *xxx_GetPossibleValuesOperation) OpName() string {
	return "/IFsrmPropertyDefinition/v0/PossibleValues"
}

func (o *xxx_GetPossibleValuesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPossibleValuesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPossibleValuesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPossibleValuesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPossibleValuesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// possibleValues {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.PossibleValues != nil {
			_ptr_possibleValues := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PossibleValues != nil {
					if err := o.PossibleValues.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PossibleValues, _ptr_possibleValues); err != nil {
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

func (o *xxx_GetPossibleValuesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// possibleValues {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_possibleValues := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PossibleValues == nil {
				o.PossibleValues = &oaut.SafeArray{}
			}
			if err := o.PossibleValues.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_possibleValues := func(ptr interface{}) { o.PossibleValues = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.PossibleValues, _s_possibleValues, _ptr_possibleValues); err != nil {
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

// GetPossibleValuesRequest structure represents the PossibleValues operation request
type GetPossibleValuesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPossibleValuesRequest) xxx_ToOp(ctx context.Context) *xxx_GetPossibleValuesOperation {
	if o == nil {
		return &xxx_GetPossibleValuesOperation{}
	}
	return &xxx_GetPossibleValuesOperation{
		This: o.This,
	}
}

func (o *GetPossibleValuesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPossibleValuesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPossibleValuesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPossibleValuesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPossibleValuesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPossibleValuesResponse structure represents the PossibleValues operation response
type GetPossibleValuesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	PossibleValues *oaut.SafeArray `idl:"name:possibleValues" json:"possible_values"`
	// Return: The PossibleValues return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPossibleValuesResponse) xxx_ToOp(ctx context.Context) *xxx_GetPossibleValuesOperation {
	if o == nil {
		return &xxx_GetPossibleValuesOperation{}
	}
	return &xxx_GetPossibleValuesOperation{
		That:           o.That,
		PossibleValues: o.PossibleValues,
		Return:         o.Return,
	}
}

func (o *GetPossibleValuesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPossibleValuesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PossibleValues = op.PossibleValues
	o.Return = op.Return
}
func (o *GetPossibleValuesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPossibleValuesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPossibleValuesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPossibleValuesOperation structure represents the PossibleValues operation
type xxx_SetPossibleValuesOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	PossibleValues *oaut.SafeArray `idl:"name:possibleValues" json:"possible_values"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPossibleValuesOperation) OpNum() int { return 17 }

func (o *xxx_SetPossibleValuesOperation) OpName() string {
	return "/IFsrmPropertyDefinition/v0/PossibleValues"
}

func (o *xxx_SetPossibleValuesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPossibleValuesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// possibleValues {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.PossibleValues != nil {
			_ptr_possibleValues := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PossibleValues != nil {
					if err := o.PossibleValues.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PossibleValues, _ptr_possibleValues); err != nil {
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

func (o *xxx_SetPossibleValuesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// possibleValues {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_possibleValues := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PossibleValues == nil {
				o.PossibleValues = &oaut.SafeArray{}
			}
			if err := o.PossibleValues.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_possibleValues := func(ptr interface{}) { o.PossibleValues = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.PossibleValues, _s_possibleValues, _ptr_possibleValues); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPossibleValuesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPossibleValuesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPossibleValuesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPossibleValuesRequest structure represents the PossibleValues operation request
type SetPossibleValuesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	PossibleValues *oaut.SafeArray `idl:"name:possibleValues" json:"possible_values"`
}

func (o *SetPossibleValuesRequest) xxx_ToOp(ctx context.Context) *xxx_SetPossibleValuesOperation {
	if o == nil {
		return &xxx_SetPossibleValuesOperation{}
	}
	return &xxx_SetPossibleValuesOperation{
		This:           o.This,
		PossibleValues: o.PossibleValues,
	}
}

func (o *SetPossibleValuesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPossibleValuesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PossibleValues = op.PossibleValues
}
func (o *SetPossibleValuesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetPossibleValuesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPossibleValuesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPossibleValuesResponse structure represents the PossibleValues operation response
type SetPossibleValuesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PossibleValues return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPossibleValuesResponse) xxx_ToOp(ctx context.Context) *xxx_SetPossibleValuesOperation {
	if o == nil {
		return &xxx_SetPossibleValuesOperation{}
	}
	return &xxx_SetPossibleValuesOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetPossibleValuesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPossibleValuesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPossibleValuesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetPossibleValuesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPossibleValuesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetValueDescriptionsOperation structure represents the ValueDescriptions operation
type xxx_GetValueDescriptionsOperation struct {
	This              *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ValueDescriptions *oaut.SafeArray `idl:"name:valueDescriptions" json:"value_descriptions"`
	Return            int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetValueDescriptionsOperation) OpNum() int { return 18 }

func (o *xxx_GetValueDescriptionsOperation) OpName() string {
	return "/IFsrmPropertyDefinition/v0/ValueDescriptions"
}

func (o *xxx_GetValueDescriptionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValueDescriptionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetValueDescriptionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetValueDescriptionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValueDescriptionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// valueDescriptions {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ValueDescriptions != nil {
			_ptr_valueDescriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ValueDescriptions != nil {
					if err := o.ValueDescriptions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ValueDescriptions, _ptr_valueDescriptions); err != nil {
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

func (o *xxx_GetValueDescriptionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// valueDescriptions {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_valueDescriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ValueDescriptions == nil {
				o.ValueDescriptions = &oaut.SafeArray{}
			}
			if err := o.ValueDescriptions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_valueDescriptions := func(ptr interface{}) { o.ValueDescriptions = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ValueDescriptions, _s_valueDescriptions, _ptr_valueDescriptions); err != nil {
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

// GetValueDescriptionsRequest structure represents the ValueDescriptions operation request
type GetValueDescriptionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetValueDescriptionsRequest) xxx_ToOp(ctx context.Context) *xxx_GetValueDescriptionsOperation {
	if o == nil {
		return &xxx_GetValueDescriptionsOperation{}
	}
	return &xxx_GetValueDescriptionsOperation{
		This: o.This,
	}
}

func (o *GetValueDescriptionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetValueDescriptionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetValueDescriptionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetValueDescriptionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValueDescriptionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetValueDescriptionsResponse structure represents the ValueDescriptions operation response
type GetValueDescriptionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That              *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ValueDescriptions *oaut.SafeArray `idl:"name:valueDescriptions" json:"value_descriptions"`
	// Return: The ValueDescriptions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetValueDescriptionsResponse) xxx_ToOp(ctx context.Context) *xxx_GetValueDescriptionsOperation {
	if o == nil {
		return &xxx_GetValueDescriptionsOperation{}
	}
	return &xxx_GetValueDescriptionsOperation{
		That:              o.That,
		ValueDescriptions: o.ValueDescriptions,
		Return:            o.Return,
	}
}

func (o *GetValueDescriptionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetValueDescriptionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ValueDescriptions = op.ValueDescriptions
	o.Return = op.Return
}
func (o *GetValueDescriptionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetValueDescriptionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValueDescriptionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetValueDescriptionsOperation structure represents the ValueDescriptions operation
type xxx_SetValueDescriptionsOperation struct {
	This              *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ValueDescriptions *oaut.SafeArray `idl:"name:valueDescriptions" json:"value_descriptions"`
	Return            int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetValueDescriptionsOperation) OpNum() int { return 19 }

func (o *xxx_SetValueDescriptionsOperation) OpName() string {
	return "/IFsrmPropertyDefinition/v0/ValueDescriptions"
}

func (o *xxx_SetValueDescriptionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueDescriptionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// valueDescriptions {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ValueDescriptions != nil {
			_ptr_valueDescriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ValueDescriptions != nil {
					if err := o.ValueDescriptions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ValueDescriptions, _ptr_valueDescriptions); err != nil {
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

func (o *xxx_SetValueDescriptionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// valueDescriptions {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_valueDescriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ValueDescriptions == nil {
				o.ValueDescriptions = &oaut.SafeArray{}
			}
			if err := o.ValueDescriptions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_valueDescriptions := func(ptr interface{}) { o.ValueDescriptions = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ValueDescriptions, _s_valueDescriptions, _ptr_valueDescriptions); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueDescriptionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueDescriptionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetValueDescriptionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetValueDescriptionsRequest structure represents the ValueDescriptions operation request
type SetValueDescriptionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis  `idl:"name:This" json:"this"`
	ValueDescriptions *oaut.SafeArray `idl:"name:valueDescriptions" json:"value_descriptions"`
}

func (o *SetValueDescriptionsRequest) xxx_ToOp(ctx context.Context) *xxx_SetValueDescriptionsOperation {
	if o == nil {
		return &xxx_SetValueDescriptionsOperation{}
	}
	return &xxx_SetValueDescriptionsOperation{
		This:              o.This,
		ValueDescriptions: o.ValueDescriptions,
	}
}

func (o *SetValueDescriptionsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetValueDescriptionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ValueDescriptions = op.ValueDescriptions
}
func (o *SetValueDescriptionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetValueDescriptionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetValueDescriptionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetValueDescriptionsResponse structure represents the ValueDescriptions operation response
type SetValueDescriptionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ValueDescriptions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetValueDescriptionsResponse) xxx_ToOp(ctx context.Context) *xxx_SetValueDescriptionsOperation {
	if o == nil {
		return &xxx_SetValueDescriptionsOperation{}
	}
	return &xxx_SetValueDescriptionsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetValueDescriptionsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetValueDescriptionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetValueDescriptionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetValueDescriptionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetValueDescriptionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetParametersOperation structure represents the Parameters operation
type xxx_GetParametersOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Parameters *oaut.SafeArray `idl:"name:parameters" json:"parameters"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetParametersOperation) OpNum() int { return 20 }

func (o *xxx_GetParametersOperation) OpName() string { return "/IFsrmPropertyDefinition/v0/Parameters" }

func (o *xxx_GetParametersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetParametersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetParametersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetParametersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetParametersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// parameters {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Parameters != nil {
			_ptr_parameters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Parameters != nil {
					if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Parameters, _ptr_parameters); err != nil {
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

func (o *xxx_GetParametersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// parameters {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_parameters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Parameters == nil {
				o.Parameters = &oaut.SafeArray{}
			}
			if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_parameters := func(ptr interface{}) { o.Parameters = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Parameters, _s_parameters, _ptr_parameters); err != nil {
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

// GetParametersRequest structure represents the Parameters operation request
type GetParametersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetParametersRequest) xxx_ToOp(ctx context.Context) *xxx_GetParametersOperation {
	if o == nil {
		return &xxx_GetParametersOperation{}
	}
	return &xxx_GetParametersOperation{
		This: o.This,
	}
}

func (o *GetParametersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetParametersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetParametersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetParametersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetParametersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetParametersResponse structure represents the Parameters operation response
type GetParametersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Parameters *oaut.SafeArray `idl:"name:parameters" json:"parameters"`
	// Return: The Parameters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetParametersResponse) xxx_ToOp(ctx context.Context) *xxx_GetParametersOperation {
	if o == nil {
		return &xxx_GetParametersOperation{}
	}
	return &xxx_GetParametersOperation{
		That:       o.That,
		Parameters: o.Parameters,
		Return:     o.Return,
	}
}

func (o *GetParametersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetParametersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Parameters = op.Parameters
	o.Return = op.Return
}
func (o *GetParametersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetParametersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetParametersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetParametersOperation structure represents the Parameters operation
type xxx_SetParametersOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Parameters *oaut.SafeArray `idl:"name:parameters" json:"parameters"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetParametersOperation) OpNum() int { return 21 }

func (o *xxx_SetParametersOperation) OpName() string { return "/IFsrmPropertyDefinition/v0/Parameters" }

func (o *xxx_SetParametersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetParametersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// parameters {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Parameters != nil {
			_ptr_parameters := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Parameters != nil {
					if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Parameters, _ptr_parameters); err != nil {
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

func (o *xxx_SetParametersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// parameters {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_parameters := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Parameters == nil {
				o.Parameters = &oaut.SafeArray{}
			}
			if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_parameters := func(ptr interface{}) { o.Parameters = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Parameters, _s_parameters, _ptr_parameters); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetParametersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetParametersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetParametersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetParametersRequest structure represents the Parameters operation request
type SetParametersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Parameters *oaut.SafeArray `idl:"name:parameters" json:"parameters"`
}

func (o *SetParametersRequest) xxx_ToOp(ctx context.Context) *xxx_SetParametersOperation {
	if o == nil {
		return &xxx_SetParametersOperation{}
	}
	return &xxx_SetParametersOperation{
		This:       o.This,
		Parameters: o.Parameters,
	}
}

func (o *SetParametersRequest) xxx_FromOp(ctx context.Context, op *xxx_SetParametersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Parameters = op.Parameters
}
func (o *SetParametersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetParametersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetParametersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetParametersResponse structure represents the Parameters operation response
type SetParametersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Parameters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetParametersResponse) xxx_ToOp(ctx context.Context) *xxx_SetParametersOperation {
	if o == nil {
		return &xxx_SetParametersOperation{}
	}
	return &xxx_SetParametersOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetParametersResponse) xxx_FromOp(ctx context.Context, op *xxx_SetParametersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetParametersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetParametersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetParametersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
