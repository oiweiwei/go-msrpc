package ifsrmclassifiermoduledefinition

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ifsrmpipelinemoduledefinition "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpipelinemoduledefinition/v0"
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
	_ = ifsrmpipelinemoduledefinition.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmClassifierModuleDefinition interface identifier bb36ea26-6318-4b8c-8592-f72dd602e7a5
	ClassifierModuleDefinitionIID = &dcom.IID{Data1: 0xbb36ea26, Data2: 0x6318, Data3: 0x4b8c, Data4: []byte{0x85, 0x92, 0xf7, 0x2d, 0xd6, 0x02, 0xe7, 0xa5}}
	// Syntax UUID
	ClassifierModuleDefinitionSyntaxUUID = &uuid.UUID{TimeLow: 0xbb36ea26, TimeMid: 0x6318, TimeHiAndVersion: 0x4b8c, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x92, Node: [6]uint8{0xf7, 0x2d, 0xd6, 0x2, 0xe7, 0xa5}}
	// Syntax ID
	ClassifierModuleDefinitionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClassifierModuleDefinitionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmClassifierModuleDefinition interface.
type ClassifierModuleDefinitionClient interface {

	// IFsrmPipelineModuleDefinition retrieval method.
	PipelineModuleDefinition() ifsrmpipelinemoduledefinition.PipelineModuleDefinitionClient

	// PropertiesAffected operation.
	GetPropertiesAffected(context.Context, *GetPropertiesAffectedRequest, ...dcerpc.CallOption) (*GetPropertiesAffectedResponse, error)

	// PropertiesAffected operation.
	SetPropertiesAffected(context.Context, *SetPropertiesAffectedRequest, ...dcerpc.CallOption) (*SetPropertiesAffectedResponse, error)

	// PropertiesUsed operation.
	GetPropertiesUsed(context.Context, *GetPropertiesUsedRequest, ...dcerpc.CallOption) (*GetPropertiesUsedResponse, error)

	// PropertiesUsed operation.
	SetPropertiesUsed(context.Context, *SetPropertiesUsedRequest, ...dcerpc.CallOption) (*SetPropertiesUsedResponse, error)

	// NeedsExplicitValue operation.
	GetNeedsExplicitValue(context.Context, *GetNeedsExplicitValueRequest, ...dcerpc.CallOption) (*GetNeedsExplicitValueResponse, error)

	// NeedsExplicitValue operation.
	SetNeedsExplicitValue(context.Context, *SetNeedsExplicitValueRequest, ...dcerpc.CallOption) (*SetNeedsExplicitValueResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClassifierModuleDefinitionClient
}

type xxx_DefaultClassifierModuleDefinitionClient struct {
	ifsrmpipelinemoduledefinition.PipelineModuleDefinitionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClassifierModuleDefinitionClient) PipelineModuleDefinition() ifsrmpipelinemoduledefinition.PipelineModuleDefinitionClient {
	return o.PipelineModuleDefinitionClient
}

func (o *xxx_DefaultClassifierModuleDefinitionClient) GetPropertiesAffected(ctx context.Context, in *GetPropertiesAffectedRequest, opts ...dcerpc.CallOption) (*GetPropertiesAffectedResponse, error) {
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
	out := &GetPropertiesAffectedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassifierModuleDefinitionClient) SetPropertiesAffected(ctx context.Context, in *SetPropertiesAffectedRequest, opts ...dcerpc.CallOption) (*SetPropertiesAffectedResponse, error) {
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
	out := &SetPropertiesAffectedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassifierModuleDefinitionClient) GetPropertiesUsed(ctx context.Context, in *GetPropertiesUsedRequest, opts ...dcerpc.CallOption) (*GetPropertiesUsedResponse, error) {
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
	out := &GetPropertiesUsedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassifierModuleDefinitionClient) SetPropertiesUsed(ctx context.Context, in *SetPropertiesUsedRequest, opts ...dcerpc.CallOption) (*SetPropertiesUsedResponse, error) {
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
	out := &SetPropertiesUsedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassifierModuleDefinitionClient) GetNeedsExplicitValue(ctx context.Context, in *GetNeedsExplicitValueRequest, opts ...dcerpc.CallOption) (*GetNeedsExplicitValueResponse, error) {
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
	out := &GetNeedsExplicitValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassifierModuleDefinitionClient) SetNeedsExplicitValue(ctx context.Context, in *SetNeedsExplicitValueRequest, opts ...dcerpc.CallOption) (*SetNeedsExplicitValueResponse, error) {
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
	out := &SetNeedsExplicitValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassifierModuleDefinitionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClassifierModuleDefinitionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClassifierModuleDefinitionClient) IPID(ctx context.Context, ipid *dcom.IPID) ClassifierModuleDefinitionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClassifierModuleDefinitionClient{
		PipelineModuleDefinitionClient: o.PipelineModuleDefinitionClient.IPID(ctx, ipid),
		cc:                             o.cc,
		ipid:                           ipid,
	}
}

func NewClassifierModuleDefinitionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClassifierModuleDefinitionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClassifierModuleDefinitionSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmpipelinemoduledefinition.NewPipelineModuleDefinitionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultClassifierModuleDefinitionClient{
		PipelineModuleDefinitionClient: base,
		cc:                             cc,
		ipid:                           ipid,
	}, nil
}

// xxx_GetPropertiesAffectedOperation structure represents the PropertiesAffected operation
type xxx_GetPropertiesAffectedOperation struct {
	This               *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat  `idl:"name:That" json:"that"`
	PropertiesAffected *oaut.SafeArray `idl:"name:propertiesAffected" json:"properties_affected"`
	Return             int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesAffectedOperation) OpNum() int { return 31 }

func (o *xxx_GetPropertiesAffectedOperation) OpName() string {
	return "/IFsrmClassifierModuleDefinition/v0/PropertiesAffected"
}

func (o *xxx_GetPropertiesAffectedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesAffectedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertiesAffectedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertiesAffectedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesAffectedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// propertiesAffected {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.PropertiesAffected != nil {
			_ptr_propertiesAffected := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertiesAffected != nil {
					if err := o.PropertiesAffected.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertiesAffected, _ptr_propertiesAffected); err != nil {
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

func (o *xxx_GetPropertiesAffectedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// propertiesAffected {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_propertiesAffected := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertiesAffected == nil {
				o.PropertiesAffected = &oaut.SafeArray{}
			}
			if err := o.PropertiesAffected.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertiesAffected := func(ptr interface{}) { o.PropertiesAffected = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.PropertiesAffected, _s_propertiesAffected, _ptr_propertiesAffected); err != nil {
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

// GetPropertiesAffectedRequest structure represents the PropertiesAffected operation request
type GetPropertiesAffectedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertiesAffectedRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesAffectedOperation) *xxx_GetPropertiesAffectedOperation {
	if op == nil {
		op = &xxx_GetPropertiesAffectedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPropertiesAffectedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesAffectedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesAffectedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesAffectedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesAffectedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesAffectedResponse structure represents the PropertiesAffected operation response
type GetPropertiesAffectedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That               *dcom.ORPCThat  `idl:"name:That" json:"that"`
	PropertiesAffected *oaut.SafeArray `idl:"name:propertiesAffected" json:"properties_affected"`
	// Return: The PropertiesAffected return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesAffectedResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesAffectedOperation) *xxx_GetPropertiesAffectedOperation {
	if op == nil {
		op = &xxx_GetPropertiesAffectedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PropertiesAffected = o.PropertiesAffected
	op.Return = o.Return
	return op
}

func (o *GetPropertiesAffectedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesAffectedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertiesAffected = op.PropertiesAffected
	o.Return = op.Return
}
func (o *GetPropertiesAffectedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesAffectedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesAffectedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPropertiesAffectedOperation structure represents the PropertiesAffected operation
type xxx_SetPropertiesAffectedOperation struct {
	This               *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat  `idl:"name:That" json:"that"`
	PropertiesAffected *oaut.SafeArray `idl:"name:propertiesAffected" json:"properties_affected"`
	Return             int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPropertiesAffectedOperation) OpNum() int { return 32 }

func (o *xxx_SetPropertiesAffectedOperation) OpName() string {
	return "/IFsrmClassifierModuleDefinition/v0/PropertiesAffected"
}

func (o *xxx_SetPropertiesAffectedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertiesAffectedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// propertiesAffected {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.PropertiesAffected != nil {
			_ptr_propertiesAffected := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertiesAffected != nil {
					if err := o.PropertiesAffected.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertiesAffected, _ptr_propertiesAffected); err != nil {
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

func (o *xxx_SetPropertiesAffectedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// propertiesAffected {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_propertiesAffected := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertiesAffected == nil {
				o.PropertiesAffected = &oaut.SafeArray{}
			}
			if err := o.PropertiesAffected.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertiesAffected := func(ptr interface{}) { o.PropertiesAffected = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.PropertiesAffected, _s_propertiesAffected, _ptr_propertiesAffected); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertiesAffectedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertiesAffectedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPropertiesAffectedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPropertiesAffectedRequest structure represents the PropertiesAffected operation request
type SetPropertiesAffectedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis  `idl:"name:This" json:"this"`
	PropertiesAffected *oaut.SafeArray `idl:"name:propertiesAffected" json:"properties_affected"`
}

func (o *SetPropertiesAffectedRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPropertiesAffectedOperation) *xxx_SetPropertiesAffectedOperation {
	if op == nil {
		op = &xxx_SetPropertiesAffectedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PropertiesAffected = o.PropertiesAffected
	return op
}

func (o *SetPropertiesAffectedRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPropertiesAffectedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PropertiesAffected = op.PropertiesAffected
}
func (o *SetPropertiesAffectedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPropertiesAffectedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPropertiesAffectedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPropertiesAffectedResponse structure represents the PropertiesAffected operation response
type SetPropertiesAffectedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PropertiesAffected return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPropertiesAffectedResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPropertiesAffectedOperation) *xxx_SetPropertiesAffectedOperation {
	if op == nil {
		op = &xxx_SetPropertiesAffectedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPropertiesAffectedResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPropertiesAffectedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPropertiesAffectedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPropertiesAffectedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPropertiesAffectedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesUsedOperation structure represents the PropertiesUsed operation
type xxx_GetPropertiesUsedOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	PropertiesUsed *oaut.SafeArray `idl:"name:propertiesUsed" json:"properties_used"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesUsedOperation) OpNum() int { return 33 }

func (o *xxx_GetPropertiesUsedOperation) OpName() string {
	return "/IFsrmClassifierModuleDefinition/v0/PropertiesUsed"
}

func (o *xxx_GetPropertiesUsedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesUsedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertiesUsedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertiesUsedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesUsedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// propertiesUsed {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.PropertiesUsed != nil {
			_ptr_propertiesUsed := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertiesUsed != nil {
					if err := o.PropertiesUsed.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertiesUsed, _ptr_propertiesUsed); err != nil {
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

func (o *xxx_GetPropertiesUsedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// propertiesUsed {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_propertiesUsed := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertiesUsed == nil {
				o.PropertiesUsed = &oaut.SafeArray{}
			}
			if err := o.PropertiesUsed.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertiesUsed := func(ptr interface{}) { o.PropertiesUsed = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.PropertiesUsed, _s_propertiesUsed, _ptr_propertiesUsed); err != nil {
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

// GetPropertiesUsedRequest structure represents the PropertiesUsed operation request
type GetPropertiesUsedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertiesUsedRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesUsedOperation) *xxx_GetPropertiesUsedOperation {
	if op == nil {
		op = &xxx_GetPropertiesUsedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPropertiesUsedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesUsedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesUsedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesUsedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesUsedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesUsedResponse structure represents the PropertiesUsed operation response
type GetPropertiesUsedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	PropertiesUsed *oaut.SafeArray `idl:"name:propertiesUsed" json:"properties_used"`
	// Return: The PropertiesUsed return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesUsedResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesUsedOperation) *xxx_GetPropertiesUsedOperation {
	if op == nil {
		op = &xxx_GetPropertiesUsedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PropertiesUsed = o.PropertiesUsed
	op.Return = o.Return
	return op
}

func (o *GetPropertiesUsedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesUsedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertiesUsed = op.PropertiesUsed
	o.Return = op.Return
}
func (o *GetPropertiesUsedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesUsedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesUsedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPropertiesUsedOperation structure represents the PropertiesUsed operation
type xxx_SetPropertiesUsedOperation struct {
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat  `idl:"name:That" json:"that"`
	PropertiesUsed *oaut.SafeArray `idl:"name:propertiesUsed" json:"properties_used"`
	Return         int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPropertiesUsedOperation) OpNum() int { return 34 }

func (o *xxx_SetPropertiesUsedOperation) OpName() string {
	return "/IFsrmClassifierModuleDefinition/v0/PropertiesUsed"
}

func (o *xxx_SetPropertiesUsedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertiesUsedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// propertiesUsed {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.PropertiesUsed != nil {
			_ptr_propertiesUsed := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertiesUsed != nil {
					if err := o.PropertiesUsed.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertiesUsed, _ptr_propertiesUsed); err != nil {
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

func (o *xxx_SetPropertiesUsedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// propertiesUsed {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_propertiesUsed := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertiesUsed == nil {
				o.PropertiesUsed = &oaut.SafeArray{}
			}
			if err := o.PropertiesUsed.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertiesUsed := func(ptr interface{}) { o.PropertiesUsed = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.PropertiesUsed, _s_propertiesUsed, _ptr_propertiesUsed); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertiesUsedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertiesUsedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPropertiesUsedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPropertiesUsedRequest structure represents the PropertiesUsed operation request
type SetPropertiesUsedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis  `idl:"name:This" json:"this"`
	PropertiesUsed *oaut.SafeArray `idl:"name:propertiesUsed" json:"properties_used"`
}

func (o *SetPropertiesUsedRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPropertiesUsedOperation) *xxx_SetPropertiesUsedOperation {
	if op == nil {
		op = &xxx_SetPropertiesUsedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PropertiesUsed = o.PropertiesUsed
	return op
}

func (o *SetPropertiesUsedRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPropertiesUsedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PropertiesUsed = op.PropertiesUsed
}
func (o *SetPropertiesUsedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPropertiesUsedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPropertiesUsedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPropertiesUsedResponse structure represents the PropertiesUsed operation response
type SetPropertiesUsedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PropertiesUsed return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPropertiesUsedResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPropertiesUsedOperation) *xxx_SetPropertiesUsedOperation {
	if op == nil {
		op = &xxx_SetPropertiesUsedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPropertiesUsedResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPropertiesUsedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPropertiesUsedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPropertiesUsedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPropertiesUsedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNeedsExplicitValueOperation structure represents the NeedsExplicitValue operation
type xxx_GetNeedsExplicitValueOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	NeedsExplicitValue int16          `idl:"name:needsExplicitValue" json:"needs_explicit_value"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNeedsExplicitValueOperation) OpNum() int { return 35 }

func (o *xxx_GetNeedsExplicitValueOperation) OpName() string {
	return "/IFsrmClassifierModuleDefinition/v0/NeedsExplicitValue"
}

func (o *xxx_GetNeedsExplicitValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNeedsExplicitValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNeedsExplicitValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNeedsExplicitValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNeedsExplicitValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// needsExplicitValue {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.NeedsExplicitValue); err != nil {
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

func (o *xxx_GetNeedsExplicitValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// needsExplicitValue {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.NeedsExplicitValue); err != nil {
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

// GetNeedsExplicitValueRequest structure represents the NeedsExplicitValue operation request
type GetNeedsExplicitValueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNeedsExplicitValueRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNeedsExplicitValueOperation) *xxx_GetNeedsExplicitValueOperation {
	if op == nil {
		op = &xxx_GetNeedsExplicitValueOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNeedsExplicitValueRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNeedsExplicitValueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNeedsExplicitValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNeedsExplicitValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNeedsExplicitValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNeedsExplicitValueResponse structure represents the NeedsExplicitValue operation response
type GetNeedsExplicitValueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	NeedsExplicitValue int16          `idl:"name:needsExplicitValue" json:"needs_explicit_value"`
	// Return: The NeedsExplicitValue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNeedsExplicitValueResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNeedsExplicitValueOperation) *xxx_GetNeedsExplicitValueOperation {
	if op == nil {
		op = &xxx_GetNeedsExplicitValueOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.NeedsExplicitValue = o.NeedsExplicitValue
	op.Return = o.Return
	return op
}

func (o *GetNeedsExplicitValueResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNeedsExplicitValueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.NeedsExplicitValue = op.NeedsExplicitValue
	o.Return = op.Return
}
func (o *GetNeedsExplicitValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNeedsExplicitValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNeedsExplicitValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNeedsExplicitValueOperation structure represents the NeedsExplicitValue operation
type xxx_SetNeedsExplicitValueOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	NeedsExplicitValue int16          `idl:"name:needsExplicitValue" json:"needs_explicit_value"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNeedsExplicitValueOperation) OpNum() int { return 36 }

func (o *xxx_SetNeedsExplicitValueOperation) OpName() string {
	return "/IFsrmClassifierModuleDefinition/v0/NeedsExplicitValue"
}

func (o *xxx_SetNeedsExplicitValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNeedsExplicitValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// needsExplicitValue {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.NeedsExplicitValue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNeedsExplicitValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// needsExplicitValue {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.NeedsExplicitValue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNeedsExplicitValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNeedsExplicitValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNeedsExplicitValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNeedsExplicitValueRequest structure represents the NeedsExplicitValue operation request
type SetNeedsExplicitValueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	NeedsExplicitValue int16          `idl:"name:needsExplicitValue" json:"needs_explicit_value"`
}

func (o *SetNeedsExplicitValueRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNeedsExplicitValueOperation) *xxx_SetNeedsExplicitValueOperation {
	if op == nil {
		op = &xxx_SetNeedsExplicitValueOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.NeedsExplicitValue = o.NeedsExplicitValue
	return op
}

func (o *SetNeedsExplicitValueRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNeedsExplicitValueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NeedsExplicitValue = op.NeedsExplicitValue
}
func (o *SetNeedsExplicitValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNeedsExplicitValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNeedsExplicitValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNeedsExplicitValueResponse structure represents the NeedsExplicitValue operation response
type SetNeedsExplicitValueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The NeedsExplicitValue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNeedsExplicitValueResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNeedsExplicitValueOperation) *xxx_SetNeedsExplicitValueOperation {
	if op == nil {
		op = &xxx_SetNeedsExplicitValueOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNeedsExplicitValueResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNeedsExplicitValueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNeedsExplicitValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNeedsExplicitValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNeedsExplicitValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
