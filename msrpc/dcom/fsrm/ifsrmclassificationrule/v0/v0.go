package ifsrmclassificationrule

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
	ifsrmrule "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmrule/v0"
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
	_ = ifsrmrule.GoPackage
	_ = fsrm.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmClassificationRule interface identifier afc052c2-5315-45ab-841b-c6db0e120148
	ClassificationRuleIID = &dcom.IID{Data1: 0xafc052c2, Data2: 0x5315, Data3: 0x45ab, Data4: []byte{0x84, 0x1b, 0xc6, 0xdb, 0x0e, 0x12, 0x01, 0x48}}
	// Syntax UUID
	ClassificationRuleSyntaxUUID = &uuid.UUID{TimeLow: 0xafc052c2, TimeMid: 0x5315, TimeHiAndVersion: 0x45ab, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0x1b, Node: [6]uint8{0xc6, 0xdb, 0xe, 0x12, 0x1, 0x48}}
	// Syntax ID
	ClassificationRuleSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClassificationRuleSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmClassificationRule interface.
type ClassificationRuleClient interface {

	// IFsrmRule retrieval method.
	Rule() ifsrmrule.RuleClient

	// ExecutionOption operation.
	GetExecutionOption(context.Context, *GetExecutionOptionRequest, ...dcerpc.CallOption) (*GetExecutionOptionResponse, error)

	// ExecutionOption operation.
	SetExecutionOption(context.Context, *SetExecutionOptionRequest, ...dcerpc.CallOption) (*SetExecutionOptionResponse, error)

	// PropertyAffected operation.
	GetPropertyAffected(context.Context, *GetPropertyAffectedRequest, ...dcerpc.CallOption) (*GetPropertyAffectedResponse, error)

	// PropertyAffected operation.
	SetPropertyAffected(context.Context, *SetPropertyAffectedRequest, ...dcerpc.CallOption) (*SetPropertyAffectedResponse, error)

	// Value operation.
	GetValue(context.Context, *GetValueRequest, ...dcerpc.CallOption) (*GetValueResponse, error)

	// Value operation.
	SetValue(context.Context, *SetValueRequest, ...dcerpc.CallOption) (*SetValueResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClassificationRuleClient
}

type xxx_DefaultClassificationRuleClient struct {
	ifsrmrule.RuleClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClassificationRuleClient) Rule() ifsrmrule.RuleClient {
	return o.RuleClient
}

func (o *xxx_DefaultClassificationRuleClient) GetExecutionOption(ctx context.Context, in *GetExecutionOptionRequest, opts ...dcerpc.CallOption) (*GetExecutionOptionResponse, error) {
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
	out := &GetExecutionOptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationRuleClient) SetExecutionOption(ctx context.Context, in *SetExecutionOptionRequest, opts ...dcerpc.CallOption) (*SetExecutionOptionResponse, error) {
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
	out := &SetExecutionOptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationRuleClient) GetPropertyAffected(ctx context.Context, in *GetPropertyAffectedRequest, opts ...dcerpc.CallOption) (*GetPropertyAffectedResponse, error) {
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
	out := &GetPropertyAffectedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationRuleClient) SetPropertyAffected(ctx context.Context, in *SetPropertyAffectedRequest, opts ...dcerpc.CallOption) (*SetPropertyAffectedResponse, error) {
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
	out := &SetPropertyAffectedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationRuleClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...dcerpc.CallOption) (*GetValueResponse, error) {
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
	out := &GetValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationRuleClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...dcerpc.CallOption) (*SetValueResponse, error) {
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
	out := &SetValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationRuleClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClassificationRuleClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClassificationRuleClient) IPID(ctx context.Context, ipid *dcom.IPID) ClassificationRuleClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClassificationRuleClient{
		RuleClient: o.RuleClient.IPID(ctx, ipid),
		cc:         o.cc,
		ipid:       ipid,
	}
}

func NewClassificationRuleClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClassificationRuleClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClassificationRuleSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmrule.NewRuleClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultClassificationRuleClient{
		RuleClient: base,
		cc:         cc,
		ipid:       ipid,
	}, nil
}

// xxx_GetExecutionOptionOperation structure represents the ExecutionOption operation
type xxx_GetExecutionOptionOperation struct {
	This            *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat       `idl:"name:That" json:"that"`
	ExecutionOption fsrm.ExecutionOption `idl:"name:executionOption" json:"execution_option"`
	Return          int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetExecutionOptionOperation) OpNum() int { return 24 }

func (o *xxx_GetExecutionOptionOperation) OpName() string {
	return "/IFsrmClassificationRule/v0/ExecutionOption"
}

func (o *xxx_GetExecutionOptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExecutionOptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetExecutionOptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetExecutionOptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExecutionOptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// executionOption {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmExecutionOption}(enum))
	{
		if err := w.WriteEnum(uint16(o.ExecutionOption)); err != nil {
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

func (o *xxx_GetExecutionOptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// executionOption {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmExecutionOption}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ExecutionOption)); err != nil {
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

// GetExecutionOptionRequest structure represents the ExecutionOption operation request
type GetExecutionOptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetExecutionOptionRequest) xxx_ToOp(ctx context.Context) *xxx_GetExecutionOptionOperation {
	if o == nil {
		return &xxx_GetExecutionOptionOperation{}
	}
	return &xxx_GetExecutionOptionOperation{
		This: o.This,
	}
}

func (o *GetExecutionOptionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetExecutionOptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetExecutionOptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetExecutionOptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExecutionOptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetExecutionOptionResponse structure represents the ExecutionOption operation response
type GetExecutionOptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat       `idl:"name:That" json:"that"`
	ExecutionOption fsrm.ExecutionOption `idl:"name:executionOption" json:"execution_option"`
	// Return: The ExecutionOption return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetExecutionOptionResponse) xxx_ToOp(ctx context.Context) *xxx_GetExecutionOptionOperation {
	if o == nil {
		return &xxx_GetExecutionOptionOperation{}
	}
	return &xxx_GetExecutionOptionOperation{
		That:            o.That,
		ExecutionOption: o.ExecutionOption,
		Return:          o.Return,
	}
}

func (o *GetExecutionOptionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetExecutionOptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ExecutionOption = op.ExecutionOption
	o.Return = op.Return
}
func (o *GetExecutionOptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetExecutionOptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExecutionOptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetExecutionOptionOperation structure represents the ExecutionOption operation
type xxx_SetExecutionOptionOperation struct {
	This            *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat       `idl:"name:That" json:"that"`
	ExecutionOption fsrm.ExecutionOption `idl:"name:executionOption" json:"execution_option"`
	Return          int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_SetExecutionOptionOperation) OpNum() int { return 25 }

func (o *xxx_SetExecutionOptionOperation) OpName() string {
	return "/IFsrmClassificationRule/v0/ExecutionOption"
}

func (o *xxx_SetExecutionOptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExecutionOptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// executionOption {in} (1:{alias=FsrmExecutionOption}(enum))
	{
		if err := w.WriteEnum(uint16(o.ExecutionOption)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExecutionOptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// executionOption {in} (1:{alias=FsrmExecutionOption}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ExecutionOption)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExecutionOptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExecutionOptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetExecutionOptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetExecutionOptionRequest structure represents the ExecutionOption operation request
type SetExecutionOptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This            *dcom.ORPCThis       `idl:"name:This" json:"this"`
	ExecutionOption fsrm.ExecutionOption `idl:"name:executionOption" json:"execution_option"`
}

func (o *SetExecutionOptionRequest) xxx_ToOp(ctx context.Context) *xxx_SetExecutionOptionOperation {
	if o == nil {
		return &xxx_SetExecutionOptionOperation{}
	}
	return &xxx_SetExecutionOptionOperation{
		This:            o.This,
		ExecutionOption: o.ExecutionOption,
	}
}

func (o *SetExecutionOptionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetExecutionOptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ExecutionOption = op.ExecutionOption
}
func (o *SetExecutionOptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetExecutionOptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExecutionOptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetExecutionOptionResponse structure represents the ExecutionOption operation response
type SetExecutionOptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExecutionOption return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetExecutionOptionResponse) xxx_ToOp(ctx context.Context) *xxx_SetExecutionOptionOperation {
	if o == nil {
		return &xxx_SetExecutionOptionOperation{}
	}
	return &xxx_SetExecutionOptionOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetExecutionOptionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetExecutionOptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetExecutionOptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetExecutionOptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExecutionOptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertyAffectedOperation structure represents the PropertyAffected operation
type xxx_GetPropertyAffectedOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Property *oaut.String   `idl:"name:property" json:"property"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertyAffectedOperation) OpNum() int { return 26 }

func (o *xxx_GetPropertyAffectedOperation) OpName() string {
	return "/IFsrmClassificationRule/v0/PropertyAffected"
}

func (o *xxx_GetPropertyAffectedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyAffectedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertyAffectedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertyAffectedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyAffectedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// property {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Property != nil {
			_ptr_property := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Property != nil {
					if err := o.Property.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Property, _ptr_property); err != nil {
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

func (o *xxx_GetPropertyAffectedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// property {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_property := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Property == nil {
				o.Property = &oaut.String{}
			}
			if err := o.Property.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_property := func(ptr interface{}) { o.Property = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Property, _s_property, _ptr_property); err != nil {
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

// GetPropertyAffectedRequest structure represents the PropertyAffected operation request
type GetPropertyAffectedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertyAffectedRequest) xxx_ToOp(ctx context.Context) *xxx_GetPropertyAffectedOperation {
	if o == nil {
		return &xxx_GetPropertyAffectedOperation{}
	}
	return &xxx_GetPropertyAffectedOperation{
		This: o.This,
	}
}

func (o *GetPropertyAffectedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertyAffectedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertyAffectedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertyAffectedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertyAffectedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertyAffectedResponse structure represents the PropertyAffected operation response
type GetPropertyAffectedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Property *oaut.String   `idl:"name:property" json:"property"`
	// Return: The PropertyAffected return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertyAffectedResponse) xxx_ToOp(ctx context.Context) *xxx_GetPropertyAffectedOperation {
	if o == nil {
		return &xxx_GetPropertyAffectedOperation{}
	}
	return &xxx_GetPropertyAffectedOperation{
		That:     o.That,
		Property: o.Property,
		Return:   o.Return,
	}
}

func (o *GetPropertyAffectedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertyAffectedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Property = op.Property
	o.Return = op.Return
}
func (o *GetPropertyAffectedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertyAffectedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertyAffectedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPropertyAffectedOperation structure represents the PropertyAffected operation
type xxx_SetPropertyAffectedOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Property *oaut.String   `idl:"name:property" json:"property"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPropertyAffectedOperation) OpNum() int { return 27 }

func (o *xxx_SetPropertyAffectedOperation) OpName() string {
	return "/IFsrmClassificationRule/v0/PropertyAffected"
}

func (o *xxx_SetPropertyAffectedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertyAffectedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// property {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Property != nil {
			_ptr_property := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Property != nil {
					if err := o.Property.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Property, _ptr_property); err != nil {
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

func (o *xxx_SetPropertyAffectedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// property {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_property := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Property == nil {
				o.Property = &oaut.String{}
			}
			if err := o.Property.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_property := func(ptr interface{}) { o.Property = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Property, _s_property, _ptr_property); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertyAffectedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertyAffectedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPropertyAffectedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPropertyAffectedRequest structure represents the PropertyAffected operation request
type SetPropertyAffectedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	Property *oaut.String   `idl:"name:property" json:"property"`
}

func (o *SetPropertyAffectedRequest) xxx_ToOp(ctx context.Context) *xxx_SetPropertyAffectedOperation {
	if o == nil {
		return &xxx_SetPropertyAffectedOperation{}
	}
	return &xxx_SetPropertyAffectedOperation{
		This:     o.This,
		Property: o.Property,
	}
}

func (o *SetPropertyAffectedRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPropertyAffectedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Property = op.Property
}
func (o *SetPropertyAffectedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetPropertyAffectedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPropertyAffectedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPropertyAffectedResponse structure represents the PropertyAffected operation response
type SetPropertyAffectedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PropertyAffected return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPropertyAffectedResponse) xxx_ToOp(ctx context.Context) *xxx_SetPropertyAffectedOperation {
	if o == nil {
		return &xxx_SetPropertyAffectedOperation{}
	}
	return &xxx_SetPropertyAffectedOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetPropertyAffectedResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPropertyAffectedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPropertyAffectedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetPropertyAffectedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPropertyAffectedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetValueOperation structure represents the Value operation
type xxx_GetValueOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value  *oaut.String   `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetValueOperation) OpNum() int { return 28 }

func (o *xxx_GetValueOperation) OpName() string { return "/IFsrmClassificationRule/v0/Value" }

func (o *xxx_GetValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// value {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Value != nil {
			_ptr_value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_value); err != nil {
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

func (o *xxx_GetValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// value {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.String{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_value := func(ptr interface{}) { o.Value = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Value, _s_value, _ptr_value); err != nil {
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

// GetValueRequest structure represents the Value operation request
type GetValueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetValueRequest) xxx_ToOp(ctx context.Context) *xxx_GetValueOperation {
	if o == nil {
		return &xxx_GetValueOperation{}
	}
	return &xxx_GetValueOperation{
		This: o.This,
	}
}

func (o *GetValueRequest) xxx_FromOp(ctx context.Context, op *xxx_GetValueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetValueResponse structure represents the Value operation response
type GetValueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value *oaut.String   `idl:"name:value" json:"value"`
	// Return: The Value return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetValueResponse) xxx_ToOp(ctx context.Context) *xxx_GetValueOperation {
	if o == nil {
		return &xxx_GetValueOperation{}
	}
	return &xxx_GetValueOperation{
		That:   o.That,
		Value:  o.Value,
		Return: o.Return,
	}
}

func (o *GetValueResponse) xxx_FromOp(ctx context.Context, op *xxx_GetValueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Value = op.Value
	o.Return = op.Return
}
func (o *GetValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetValueOperation structure represents the Value operation
type xxx_SetValueOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value  *oaut.String   `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetValueOperation) OpNum() int { return 29 }

func (o *xxx_SetValueOperation) OpName() string { return "/IFsrmClassificationRule/v0/Value" }

func (o *xxx_SetValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// value {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Value != nil {
			_ptr_value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_value); err != nil {
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

func (o *xxx_SetValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// value {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.String{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_value := func(ptr interface{}) { o.Value = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Value, _s_value, _ptr_value); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetValueRequest structure represents the Value operation request
type SetValueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Value *oaut.String   `idl:"name:value" json:"value"`
}

func (o *SetValueRequest) xxx_ToOp(ctx context.Context) *xxx_SetValueOperation {
	if o == nil {
		return &xxx_SetValueOperation{}
	}
	return &xxx_SetValueOperation{
		This:  o.This,
		Value: o.Value,
	}
}

func (o *SetValueRequest) xxx_FromOp(ctx context.Context, op *xxx_SetValueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Value = op.Value
}
func (o *SetValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetValueResponse structure represents the Value operation response
type SetValueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Value return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetValueResponse) xxx_ToOp(ctx context.Context) *xxx_SetValueOperation {
	if o == nil {
		return &xxx_SetValueOperation{}
	}
	return &xxx_SetValueOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetValueResponse) xxx_FromOp(ctx context.Context, op *xxx_SetValueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
