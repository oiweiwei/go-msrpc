package ifsrmrule

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
	// IFsrmRule interface identifier cb0df960-16f5-4495-9079-3f9360d831df
	RuleIID = &dcom.IID{Data1: 0xcb0df960, Data2: 0x16f5, Data3: 0x4495, Data4: []byte{0x90, 0x79, 0x3f, 0x93, 0x60, 0xd8, 0x31, 0xdf}}
	// Syntax UUID
	RuleSyntaxUUID = &uuid.UUID{TimeLow: 0xcb0df960, TimeMid: 0x16f5, TimeHiAndVersion: 0x4495, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0x79, Node: [6]uint8{0x3f, 0x93, 0x60, 0xd8, 0x31, 0xdf}}
	// Syntax ID
	RuleSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RuleSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmRule interface.
type RuleClient interface {

	// IFsrmObject retrieval method.
	Object() ifsrmobject.ObjectClient

	// Name operation.
	GetName(context.Context, *GetNameRequest, ...dcerpc.CallOption) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest, ...dcerpc.CallOption) (*SetNameResponse, error)

	// The RuleType (get) method retrieves the rule type of the rule as defined in the FsrmRuleType
	// (section 2.2.1.2.11) enumeration and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------+
	//	|         RETURN          |                             |
	//	|       VALUE/CODE        |         DESCRIPTION         |
	//	|                         |                             |
	//	+-------------------------+-----------------------------+
	//	+-------------------------+-----------------------------+
	//	| 0x80070057 E_INVALIDARG | The type parameter is NULL. |
	//	+-------------------------+-----------------------------+
	GetRuleType(context.Context, *GetRuleTypeRequest, ...dcerpc.CallOption) (*GetRuleTypeResponse, error)

	// ModuleDefinitionName operation.
	GetModuleDefinitionName(context.Context, *GetModuleDefinitionNameRequest, ...dcerpc.CallOption) (*GetModuleDefinitionNameResponse, error)

	// ModuleDefinitionName operation.
	SetModuleDefinitionName(context.Context, *SetModuleDefinitionNameRequest, ...dcerpc.CallOption) (*SetModuleDefinitionNameResponse, error)

	// NamespaceRoots operation.
	GetNamespaceRoots(context.Context, *GetNamespaceRootsRequest, ...dcerpc.CallOption) (*GetNamespaceRootsResponse, error)

	// NamespaceRoots operation.
	SetNamespaceRoots(context.Context, *SetNamespaceRootsRequest, ...dcerpc.CallOption) (*SetNamespaceRootsResponse, error)

	// RuleFlags operation.
	GetRuleFlags(context.Context, *GetRuleFlagsRequest, ...dcerpc.CallOption) (*GetRuleFlagsResponse, error)

	// RuleFlags operation.
	SetRuleFlags(context.Context, *SetRuleFlagsRequest, ...dcerpc.CallOption) (*SetRuleFlagsResponse, error)

	// Parameters operation.
	GetParameters(context.Context, *GetParametersRequest, ...dcerpc.CallOption) (*GetParametersResponse, error)

	// Parameters operation.
	SetParameters(context.Context, *SetParametersRequest, ...dcerpc.CallOption) (*SetParametersResponse, error)

	// The LastModified (get) method retrieves the last modified time corresponding to the
	// time the rule was last modified and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-------------------------------------+
	//	|         RETURN          |                                     |
	//	|       VALUE/CODE        |             DESCRIPTION             |
	//	|                         |                                     |
	//	+-------------------------+-------------------------------------+
	//	+-------------------------+-------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The lastModified parameter is NULL. |
	//	+-------------------------+-------------------------------------+
	GetLastModified(context.Context, *GetLastModifiedRequest, ...dcerpc.CallOption) (*GetLastModifiedResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RuleClient
}

type xxx_DefaultRuleClient struct {
	ifsrmobject.ObjectClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRuleClient) Object() ifsrmobject.ObjectClient {
	return o.ObjectClient
}

func (o *xxx_DefaultRuleClient) GetName(ctx context.Context, in *GetNameRequest, opts ...dcerpc.CallOption) (*GetNameResponse, error) {
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

func (o *xxx_DefaultRuleClient) SetName(ctx context.Context, in *SetNameRequest, opts ...dcerpc.CallOption) (*SetNameResponse, error) {
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

func (o *xxx_DefaultRuleClient) GetRuleType(ctx context.Context, in *GetRuleTypeRequest, opts ...dcerpc.CallOption) (*GetRuleTypeResponse, error) {
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
	out := &GetRuleTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRuleClient) GetModuleDefinitionName(ctx context.Context, in *GetModuleDefinitionNameRequest, opts ...dcerpc.CallOption) (*GetModuleDefinitionNameResponse, error) {
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
	out := &GetModuleDefinitionNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRuleClient) SetModuleDefinitionName(ctx context.Context, in *SetModuleDefinitionNameRequest, opts ...dcerpc.CallOption) (*SetModuleDefinitionNameResponse, error) {
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
	out := &SetModuleDefinitionNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRuleClient) GetNamespaceRoots(ctx context.Context, in *GetNamespaceRootsRequest, opts ...dcerpc.CallOption) (*GetNamespaceRootsResponse, error) {
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
	out := &GetNamespaceRootsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRuleClient) SetNamespaceRoots(ctx context.Context, in *SetNamespaceRootsRequest, opts ...dcerpc.CallOption) (*SetNamespaceRootsResponse, error) {
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
	out := &SetNamespaceRootsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRuleClient) GetRuleFlags(ctx context.Context, in *GetRuleFlagsRequest, opts ...dcerpc.CallOption) (*GetRuleFlagsResponse, error) {
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
	out := &GetRuleFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRuleClient) SetRuleFlags(ctx context.Context, in *SetRuleFlagsRequest, opts ...dcerpc.CallOption) (*SetRuleFlagsResponse, error) {
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
	out := &SetRuleFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRuleClient) GetParameters(ctx context.Context, in *GetParametersRequest, opts ...dcerpc.CallOption) (*GetParametersResponse, error) {
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

func (o *xxx_DefaultRuleClient) SetParameters(ctx context.Context, in *SetParametersRequest, opts ...dcerpc.CallOption) (*SetParametersResponse, error) {
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

func (o *xxx_DefaultRuleClient) GetLastModified(ctx context.Context, in *GetLastModifiedRequest, opts ...dcerpc.CallOption) (*GetLastModifiedResponse, error) {
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
	out := &GetLastModifiedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRuleClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRuleClient) IPID(ctx context.Context, ipid *dcom.IPID) RuleClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRuleClient{
		ObjectClient: o.ObjectClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}
func NewRuleClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RuleClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RuleSyntaxV0_0))...)
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
	return &xxx_DefaultRuleClient{
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

func (o *xxx_GetNameOperation) OpName() string { return "/IFsrmRule/v0/Name" }

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

func (o *xxx_SetNameOperation) OpName() string { return "/IFsrmRule/v0/Name" }

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

// xxx_GetRuleTypeOperation structure represents the RuleType operation
type xxx_GetRuleTypeOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	RuleType fsrm.RuleType  `idl:"name:ruleType" json:"rule_type"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRuleTypeOperation) OpNum() int { return 14 }

func (o *xxx_GetRuleTypeOperation) OpName() string { return "/IFsrmRule/v0/RuleType" }

func (o *xxx_GetRuleTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRuleTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRuleTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ruleType {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmRuleType}(enum))
	{
		if err := w.WriteData(uint16(o.RuleType)); err != nil {
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

func (o *xxx_GetRuleTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ruleType {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmRuleType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.RuleType)); err != nil {
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

// GetRuleTypeRequest structure represents the RuleType operation request
type GetRuleTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRuleTypeRequest) xxx_ToOp(ctx context.Context) *xxx_GetRuleTypeOperation {
	if o == nil {
		return &xxx_GetRuleTypeOperation{}
	}
	return &xxx_GetRuleTypeOperation{
		This: o.This,
	}
}

func (o *GetRuleTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRuleTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRuleTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRuleTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRuleTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRuleTypeResponse structure represents the RuleType operation response
type GetRuleTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ruleType: Pointer to a variable that upon completion contains the rule type of the
	// rule.
	RuleType fsrm.RuleType `idl:"name:ruleType" json:"rule_type"`
	// Return: The RuleType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRuleTypeResponse) xxx_ToOp(ctx context.Context) *xxx_GetRuleTypeOperation {
	if o == nil {
		return &xxx_GetRuleTypeOperation{}
	}
	return &xxx_GetRuleTypeOperation{
		That:     o.That,
		RuleType: o.RuleType,
		Return:   o.Return,
	}
}

func (o *GetRuleTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRuleTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RuleType = op.RuleType
	o.Return = op.Return
}
func (o *GetRuleTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRuleTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRuleTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetModuleDefinitionNameOperation structure represents the ModuleDefinitionName operation
type xxx_GetModuleDefinitionNameOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	ModuleDefinitionName *oaut.String   `idl:"name:moduleDefinitionName" json:"module_definition_name"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetModuleDefinitionNameOperation) OpNum() int { return 15 }

func (o *xxx_GetModuleDefinitionNameOperation) OpName() string {
	return "/IFsrmRule/v0/ModuleDefinitionName"
}

func (o *xxx_GetModuleDefinitionNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleDefinitionNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetModuleDefinitionNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetModuleDefinitionNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleDefinitionNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// moduleDefinitionName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ModuleDefinitionName != nil {
			_ptr_moduleDefinitionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModuleDefinitionName != nil {
					if err := o.ModuleDefinitionName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleDefinitionName, _ptr_moduleDefinitionName); err != nil {
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

func (o *xxx_GetModuleDefinitionNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// moduleDefinitionName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_moduleDefinitionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModuleDefinitionName == nil {
				o.ModuleDefinitionName = &oaut.String{}
			}
			if err := o.ModuleDefinitionName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_moduleDefinitionName := func(ptr interface{}) { o.ModuleDefinitionName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ModuleDefinitionName, _s_moduleDefinitionName, _ptr_moduleDefinitionName); err != nil {
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

// GetModuleDefinitionNameRequest structure represents the ModuleDefinitionName operation request
type GetModuleDefinitionNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetModuleDefinitionNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetModuleDefinitionNameOperation {
	if o == nil {
		return &xxx_GetModuleDefinitionNameOperation{}
	}
	return &xxx_GetModuleDefinitionNameOperation{
		This: o.This,
	}
}

func (o *GetModuleDefinitionNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetModuleDefinitionNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetModuleDefinitionNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetModuleDefinitionNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetModuleDefinitionNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetModuleDefinitionNameResponse structure represents the ModuleDefinitionName operation response
type GetModuleDefinitionNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	ModuleDefinitionName *oaut.String   `idl:"name:moduleDefinitionName" json:"module_definition_name"`
	// Return: The ModuleDefinitionName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetModuleDefinitionNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetModuleDefinitionNameOperation {
	if o == nil {
		return &xxx_GetModuleDefinitionNameOperation{}
	}
	return &xxx_GetModuleDefinitionNameOperation{
		That:                 o.That,
		ModuleDefinitionName: o.ModuleDefinitionName,
		Return:               o.Return,
	}
}

func (o *GetModuleDefinitionNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetModuleDefinitionNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ModuleDefinitionName = op.ModuleDefinitionName
	o.Return = op.Return
}
func (o *GetModuleDefinitionNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetModuleDefinitionNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetModuleDefinitionNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetModuleDefinitionNameOperation structure represents the ModuleDefinitionName operation
type xxx_SetModuleDefinitionNameOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	ModuleDefinitionName *oaut.String   `idl:"name:moduleDefinitionName" json:"module_definition_name"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetModuleDefinitionNameOperation) OpNum() int { return 16 }

func (o *xxx_SetModuleDefinitionNameOperation) OpName() string {
	return "/IFsrmRule/v0/ModuleDefinitionName"
}

func (o *xxx_SetModuleDefinitionNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetModuleDefinitionNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// moduleDefinitionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ModuleDefinitionName != nil {
			_ptr_moduleDefinitionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModuleDefinitionName != nil {
					if err := o.ModuleDefinitionName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleDefinitionName, _ptr_moduleDefinitionName); err != nil {
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

func (o *xxx_SetModuleDefinitionNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// moduleDefinitionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_moduleDefinitionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModuleDefinitionName == nil {
				o.ModuleDefinitionName = &oaut.String{}
			}
			if err := o.ModuleDefinitionName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_moduleDefinitionName := func(ptr interface{}) { o.ModuleDefinitionName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ModuleDefinitionName, _s_moduleDefinitionName, _ptr_moduleDefinitionName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetModuleDefinitionNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetModuleDefinitionNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetModuleDefinitionNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetModuleDefinitionNameRequest structure represents the ModuleDefinitionName operation request
type SetModuleDefinitionNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	ModuleDefinitionName *oaut.String   `idl:"name:moduleDefinitionName" json:"module_definition_name"`
}

func (o *SetModuleDefinitionNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetModuleDefinitionNameOperation {
	if o == nil {
		return &xxx_SetModuleDefinitionNameOperation{}
	}
	return &xxx_SetModuleDefinitionNameOperation{
		This:                 o.This,
		ModuleDefinitionName: o.ModuleDefinitionName,
	}
}

func (o *SetModuleDefinitionNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetModuleDefinitionNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ModuleDefinitionName = op.ModuleDefinitionName
}
func (o *SetModuleDefinitionNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetModuleDefinitionNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetModuleDefinitionNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetModuleDefinitionNameResponse structure represents the ModuleDefinitionName operation response
type SetModuleDefinitionNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModuleDefinitionName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetModuleDefinitionNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetModuleDefinitionNameOperation {
	if o == nil {
		return &xxx_SetModuleDefinitionNameOperation{}
	}
	return &xxx_SetModuleDefinitionNameOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetModuleDefinitionNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetModuleDefinitionNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetModuleDefinitionNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetModuleDefinitionNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetModuleDefinitionNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNamespaceRootsOperation structure represents the NamespaceRoots operation
type xxx_GetNamespaceRootsOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	NamespaceRoots *oaut.SafeArray `idl:"name:namespaceRoots" json:"namespace_roots"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNamespaceRootsOperation) OpNum() int { return 17 }

func (o *xxx_GetNamespaceRootsOperation) OpName() string { return "/IFsrmRule/v0/NamespaceRoots" }

func (o *xxx_GetNamespaceRootsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamespaceRootsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNamespaceRootsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNamespaceRootsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamespaceRootsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// namespaceRoots {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.NamespaceRoots != nil {
			_ptr_namespaceRoots := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NamespaceRoots != nil {
					if err := o.NamespaceRoots.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NamespaceRoots, _ptr_namespaceRoots); err != nil {
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

func (o *xxx_GetNamespaceRootsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// namespaceRoots {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_namespaceRoots := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NamespaceRoots == nil {
				o.NamespaceRoots = &oaut.SafeArray{}
			}
			if err := o.NamespaceRoots.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_namespaceRoots := func(ptr interface{}) { o.NamespaceRoots = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.NamespaceRoots, _s_namespaceRoots, _ptr_namespaceRoots); err != nil {
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

// GetNamespaceRootsRequest structure represents the NamespaceRoots operation request
type GetNamespaceRootsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNamespaceRootsRequest) xxx_ToOp(ctx context.Context) *xxx_GetNamespaceRootsOperation {
	if o == nil {
		return &xxx_GetNamespaceRootsOperation{}
	}
	return &xxx_GetNamespaceRootsOperation{
		This: o.This,
	}
}

func (o *GetNamespaceRootsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNamespaceRootsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNamespaceRootsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNamespaceRootsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNamespaceRootsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNamespaceRootsResponse structure represents the NamespaceRoots operation response
type GetNamespaceRootsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	NamespaceRoots *oaut.SafeArray `idl:"name:namespaceRoots" json:"namespace_roots"`
	// Return: The NamespaceRoots return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNamespaceRootsResponse) xxx_ToOp(ctx context.Context) *xxx_GetNamespaceRootsOperation {
	if o == nil {
		return &xxx_GetNamespaceRootsOperation{}
	}
	return &xxx_GetNamespaceRootsOperation{
		That:           o.That,
		NamespaceRoots: o.NamespaceRoots,
		Return:         o.Return,
	}
}

func (o *GetNamespaceRootsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNamespaceRootsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.NamespaceRoots = op.NamespaceRoots
	o.Return = op.Return
}
func (o *GetNamespaceRootsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNamespaceRootsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNamespaceRootsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNamespaceRootsOperation structure represents the NamespaceRoots operation
type xxx_SetNamespaceRootsOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	NamespaceRoots *oaut.SafeArray `idl:"name:namespaceRoots" json:"namespace_roots"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNamespaceRootsOperation) OpNum() int { return 18 }

func (o *xxx_SetNamespaceRootsOperation) OpName() string { return "/IFsrmRule/v0/NamespaceRoots" }

func (o *xxx_SetNamespaceRootsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNamespaceRootsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// namespaceRoots {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.NamespaceRoots != nil {
			_ptr_namespaceRoots := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NamespaceRoots != nil {
					if err := o.NamespaceRoots.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NamespaceRoots, _ptr_namespaceRoots); err != nil {
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

func (o *xxx_SetNamespaceRootsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// namespaceRoots {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_namespaceRoots := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NamespaceRoots == nil {
				o.NamespaceRoots = &oaut.SafeArray{}
			}
			if err := o.NamespaceRoots.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_namespaceRoots := func(ptr interface{}) { o.NamespaceRoots = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.NamespaceRoots, _s_namespaceRoots, _ptr_namespaceRoots); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNamespaceRootsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNamespaceRootsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNamespaceRootsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNamespaceRootsRequest structure represents the NamespaceRoots operation request
type SetNamespaceRootsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	NamespaceRoots *oaut.SafeArray `idl:"name:namespaceRoots" json:"namespace_roots"`
}

func (o *SetNamespaceRootsRequest) xxx_ToOp(ctx context.Context) *xxx_SetNamespaceRootsOperation {
	if o == nil {
		return &xxx_SetNamespaceRootsOperation{}
	}
	return &xxx_SetNamespaceRootsOperation{
		This:           o.This,
		NamespaceRoots: o.NamespaceRoots,
	}
}

func (o *SetNamespaceRootsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNamespaceRootsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NamespaceRoots = op.NamespaceRoots
}
func (o *SetNamespaceRootsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetNamespaceRootsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNamespaceRootsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNamespaceRootsResponse structure represents the NamespaceRoots operation response
type SetNamespaceRootsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The NamespaceRoots return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNamespaceRootsResponse) xxx_ToOp(ctx context.Context) *xxx_SetNamespaceRootsOperation {
	if o == nil {
		return &xxx_SetNamespaceRootsOperation{}
	}
	return &xxx_SetNamespaceRootsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetNamespaceRootsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNamespaceRootsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNamespaceRootsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetNamespaceRootsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNamespaceRootsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRuleFlagsOperation structure represents the RuleFlags operation
type xxx_GetRuleFlagsOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RuleFlags int32          `idl:"name:ruleFlags" json:"rule_flags"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRuleFlagsOperation) OpNum() int { return 19 }

func (o *xxx_GetRuleFlagsOperation) OpName() string { return "/IFsrmRule/v0/RuleFlags" }

func (o *xxx_GetRuleFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRuleFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRuleFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ruleFlags {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.RuleFlags); err != nil {
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

func (o *xxx_GetRuleFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ruleFlags {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.RuleFlags); err != nil {
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

// GetRuleFlagsRequest structure represents the RuleFlags operation request
type GetRuleFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRuleFlagsRequest) xxx_ToOp(ctx context.Context) *xxx_GetRuleFlagsOperation {
	if o == nil {
		return &xxx_GetRuleFlagsOperation{}
	}
	return &xxx_GetRuleFlagsOperation{
		This: o.This,
	}
}

func (o *GetRuleFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRuleFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRuleFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRuleFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRuleFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRuleFlagsResponse structure represents the RuleFlags operation response
type GetRuleFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RuleFlags int32          `idl:"name:ruleFlags" json:"rule_flags"`
	// Return: The RuleFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRuleFlagsResponse) xxx_ToOp(ctx context.Context) *xxx_GetRuleFlagsOperation {
	if o == nil {
		return &xxx_GetRuleFlagsOperation{}
	}
	return &xxx_GetRuleFlagsOperation{
		That:      o.That,
		RuleFlags: o.RuleFlags,
		Return:    o.Return,
	}
}

func (o *GetRuleFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRuleFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RuleFlags = op.RuleFlags
	o.Return = op.Return
}
func (o *GetRuleFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRuleFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRuleFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetRuleFlagsOperation structure represents the RuleFlags operation
type xxx_SetRuleFlagsOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	RuleFlags int32          `idl:"name:ruleFlags" json:"rule_flags"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetRuleFlagsOperation) OpNum() int { return 20 }

func (o *xxx_SetRuleFlagsOperation) OpName() string { return "/IFsrmRule/v0/RuleFlags" }

func (o *xxx_SetRuleFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRuleFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ruleFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.RuleFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRuleFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ruleFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.RuleFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRuleFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRuleFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetRuleFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetRuleFlagsRequest structure represents the RuleFlags operation request
type SetRuleFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	RuleFlags int32          `idl:"name:ruleFlags" json:"rule_flags"`
}

func (o *SetRuleFlagsRequest) xxx_ToOp(ctx context.Context) *xxx_SetRuleFlagsOperation {
	if o == nil {
		return &xxx_SetRuleFlagsOperation{}
	}
	return &xxx_SetRuleFlagsOperation{
		This:      o.This,
		RuleFlags: o.RuleFlags,
	}
}

func (o *SetRuleFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetRuleFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RuleFlags = op.RuleFlags
}
func (o *SetRuleFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetRuleFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRuleFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetRuleFlagsResponse structure represents the RuleFlags operation response
type SetRuleFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RuleFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetRuleFlagsResponse) xxx_ToOp(ctx context.Context) *xxx_SetRuleFlagsOperation {
	if o == nil {
		return &xxx_SetRuleFlagsOperation{}
	}
	return &xxx_SetRuleFlagsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetRuleFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetRuleFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetRuleFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetRuleFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRuleFlagsOperation{}
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

func (o *xxx_GetParametersOperation) OpNum() int { return 21 }

func (o *xxx_GetParametersOperation) OpName() string { return "/IFsrmRule/v0/Parameters" }

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

func (o *xxx_SetParametersOperation) OpNum() int { return 22 }

func (o *xxx_SetParametersOperation) OpName() string { return "/IFsrmRule/v0/Parameters" }

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

// xxx_GetLastModifiedOperation structure represents the LastModified operation
type xxx_GetLastModifiedOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	LastModified float64        `idl:"name:lastModified" json:"last_modified"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLastModifiedOperation) OpNum() int { return 23 }

func (o *xxx_GetLastModifiedOperation) OpName() string { return "/IFsrmRule/v0/LastModified" }

func (o *xxx_GetLastModifiedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastModifiedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLastModifiedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLastModifiedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLastModifiedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lastModified {out, retval} (1:{pointer=ref}*(1))(2:{alias=DATE}(float64))
	{
		if err := w.WriteData(o.LastModified); err != nil {
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

func (o *xxx_GetLastModifiedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lastModified {out, retval} (1:{pointer=ref}*(1))(2:{alias=DATE}(float64))
	{
		if err := w.ReadData(&o.LastModified); err != nil {
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

// GetLastModifiedRequest structure represents the LastModified operation request
type GetLastModifiedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLastModifiedRequest) xxx_ToOp(ctx context.Context) *xxx_GetLastModifiedOperation {
	if o == nil {
		return &xxx_GetLastModifiedOperation{}
	}
	return &xxx_GetLastModifiedOperation{
		This: o.This,
	}
}

func (o *GetLastModifiedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLastModifiedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLastModifiedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetLastModifiedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastModifiedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLastModifiedResponse structure represents the LastModified operation response
type GetLastModifiedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lastModified: Pointer to a variable that upon completion contains the last modified
	// time when the rule was last modified.
	LastModified float64 `idl:"name:lastModified" json:"last_modified"`
	// Return: The LastModified return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLastModifiedResponse) xxx_ToOp(ctx context.Context) *xxx_GetLastModifiedOperation {
	if o == nil {
		return &xxx_GetLastModifiedOperation{}
	}
	return &xxx_GetLastModifiedOperation{
		That:         o.That,
		LastModified: o.LastModified,
		Return:       o.Return,
	}
}

func (o *GetLastModifiedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLastModifiedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LastModified = op.LastModified
	o.Return = op.Return
}
func (o *GetLastModifiedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetLastModifiedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLastModifiedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
