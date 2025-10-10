package iwrmpolicy

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
	wsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
	_ = wsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IWRMPolicy interface identifier 59602eb6-57b0-4fd8-aa4b-ebf06971fe15
	PolicyIID = &dcom.IID{Data1: 0x59602eb6, Data2: 0x57b0, Data3: 0x4fd8, Data4: []byte{0xaa, 0x4b, 0xeb, 0xf0, 0x69, 0x71, 0xfe, 0x15}}
	// Syntax UUID
	PolicySyntaxUUID = &uuid.UUID{TimeLow: 0x59602eb6, TimeMid: 0x57b0, TimeHiAndVersion: 0x4fd8, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0xeb, 0xf0, 0x69, 0x71, 0xfe, 0x15}}
	// Syntax ID
	PolicySyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: PolicySyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWRMPolicy interface.
type PolicyClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	GetPolicyInfo(context.Context, *GetPolicyInfoRequest, ...dcerpc.CallOption) (*GetPolicyInfoResponse, error)

	CreatePolicy(context.Context, *CreatePolicyRequest, ...dcerpc.CallOption) (*CreatePolicyResponse, error)

	ModifyPolicy(context.Context, *ModifyPolicyRequest, ...dcerpc.CallOption) (*ModifyPolicyResponse, error)

	DeletePolicy(context.Context, *DeletePolicyRequest, ...dcerpc.CallOption) (*DeletePolicyResponse, error)

	RenameAllocationPolicy(context.Context, *RenameAllocationPolicyRequest, ...dcerpc.CallOption) (*RenameAllocationPolicyResponse, error)

	MoveBefore(context.Context, *MoveBeforeRequest, ...dcerpc.CallOption) (*MoveBeforeResponse, error)

	MoveAfter(context.Context, *MoveAfterRequest, ...dcerpc.CallOption) (*MoveAfterResponse, error)

	SetCALDefaultPolicyName(context.Context, *SetCALDefaultPolicyNameRequest, ...dcerpc.CallOption) (*SetCALDefaultPolicyNameResponse, error)

	GetCALDefaultPolicyName(context.Context, *GetCALDefaultPolicyNameRequest, ...dcerpc.CallOption) (*GetCALDefaultPolicyNameResponse, error)

	GetProcessList(context.Context, *GetProcessListRequest, ...dcerpc.CallOption) (*GetProcessListResponse, error)

	GetCurrentPolicy(context.Context, *GetCurrentPolicyRequest, ...dcerpc.CallOption) (*GetCurrentPolicyResponse, error)

	SetCurrentPolicy(context.Context, *SetCurrentPolicyRequest, ...dcerpc.CallOption) (*SetCurrentPolicyResponse, error)

	GetCurrentStateAndActivePolicyName(context.Context, *GetCurrentStateAndActivePolicyNameRequest, ...dcerpc.CallOption) (*GetCurrentStateAndActivePolicyNameResponse, error)

	GetConditionalPolicy(context.Context, *GetConditionalPolicyRequest, ...dcerpc.CallOption) (*GetConditionalPolicyResponse, error)

	SetConditionalPolicy(context.Context, *SetConditionalPolicyRequest, ...dcerpc.CallOption) (*SetConditionalPolicyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) PolicyClient
}

type xxx_DefaultPolicyClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPolicyClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultPolicyClient) GetPolicyInfo(ctx context.Context, in *GetPolicyInfoRequest, opts ...dcerpc.CallOption) (*GetPolicyInfoResponse, error) {
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
	out := &GetPolicyInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...dcerpc.CallOption) (*CreatePolicyResponse, error) {
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
	out := &CreatePolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) ModifyPolicy(ctx context.Context, in *ModifyPolicyRequest, opts ...dcerpc.CallOption) (*ModifyPolicyResponse, error) {
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
	out := &ModifyPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...dcerpc.CallOption) (*DeletePolicyResponse, error) {
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
	out := &DeletePolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) RenameAllocationPolicy(ctx context.Context, in *RenameAllocationPolicyRequest, opts ...dcerpc.CallOption) (*RenameAllocationPolicyResponse, error) {
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
	out := &RenameAllocationPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) MoveBefore(ctx context.Context, in *MoveBeforeRequest, opts ...dcerpc.CallOption) (*MoveBeforeResponse, error) {
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
	out := &MoveBeforeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) MoveAfter(ctx context.Context, in *MoveAfterRequest, opts ...dcerpc.CallOption) (*MoveAfterResponse, error) {
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
	out := &MoveAfterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) SetCALDefaultPolicyName(ctx context.Context, in *SetCALDefaultPolicyNameRequest, opts ...dcerpc.CallOption) (*SetCALDefaultPolicyNameResponse, error) {
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
	out := &SetCALDefaultPolicyNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) GetCALDefaultPolicyName(ctx context.Context, in *GetCALDefaultPolicyNameRequest, opts ...dcerpc.CallOption) (*GetCALDefaultPolicyNameResponse, error) {
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
	out := &GetCALDefaultPolicyNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) GetProcessList(ctx context.Context, in *GetProcessListRequest, opts ...dcerpc.CallOption) (*GetProcessListResponse, error) {
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
	out := &GetProcessListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) GetCurrentPolicy(ctx context.Context, in *GetCurrentPolicyRequest, opts ...dcerpc.CallOption) (*GetCurrentPolicyResponse, error) {
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
	out := &GetCurrentPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) SetCurrentPolicy(ctx context.Context, in *SetCurrentPolicyRequest, opts ...dcerpc.CallOption) (*SetCurrentPolicyResponse, error) {
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
	out := &SetCurrentPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) GetCurrentStateAndActivePolicyName(ctx context.Context, in *GetCurrentStateAndActivePolicyNameRequest, opts ...dcerpc.CallOption) (*GetCurrentStateAndActivePolicyNameResponse, error) {
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
	out := &GetCurrentStateAndActivePolicyNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) GetConditionalPolicy(ctx context.Context, in *GetConditionalPolicyRequest, opts ...dcerpc.CallOption) (*GetConditionalPolicyResponse, error) {
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
	out := &GetConditionalPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) SetConditionalPolicy(ctx context.Context, in *SetConditionalPolicyRequest, opts ...dcerpc.CallOption) (*SetConditionalPolicyResponse, error) {
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
	out := &SetConditionalPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPolicyClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPolicyClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultPolicyClient) IPID(ctx context.Context, ipid *dcom.IPID) PolicyClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPolicyClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewPolicyClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PolicyClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PolicySyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultPolicyClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetPolicyInfoOperation structure represents the GetPolicyInfo operation
type xxx_GetPolicyInfoOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyName *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
	PolicyInfo *oaut.String   `idl:"name:pbstrPolicyInfo" json:"policy_info"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPolicyInfoOperation) OpNum() int { return 7 }

func (o *xxx_GetPolicyInfoOperation) OpName() string { return "/IWRMPolicy/v0/GetPolicyInfo" }

func (o *xxx_GetPolicyInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPolicyInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyName != nil {
			_ptr_bstrPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyName != nil {
					if err := o.PolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyName, _ptr_bstrPolicyName); err != nil {
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

func (o *xxx_GetPolicyInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyName == nil {
				o.PolicyName = &oaut.String{}
			}
			if err := o.PolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyName, _s_bstrPolicyName, _ptr_bstrPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPolicyInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPolicyInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrPolicyInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyInfo != nil {
			_ptr_pbstrPolicyInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyInfo != nil {
					if err := o.PolicyInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyInfo, _ptr_pbstrPolicyInfo); err != nil {
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

func (o *xxx_GetPolicyInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrPolicyInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPolicyInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyInfo == nil {
				o.PolicyInfo = &oaut.String{}
			}
			if err := o.PolicyInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPolicyInfo := func(ptr interface{}) { o.PolicyInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyInfo, _s_pbstrPolicyInfo, _ptr_pbstrPolicyInfo); err != nil {
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

// GetPolicyInfoRequest structure represents the GetPolicyInfo operation request
type GetPolicyInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	PolicyName *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
}

func (o *GetPolicyInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPolicyInfoOperation) *xxx_GetPolicyInfoOperation {
	if op == nil {
		op = &xxx_GetPolicyInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PolicyName = o.PolicyName
	return op
}

func (o *GetPolicyInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPolicyInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PolicyName = op.PolicyName
}
func (o *GetPolicyInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPolicyInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPolicyInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPolicyInfoResponse structure represents the GetPolicyInfo operation response
type GetPolicyInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyInfo *oaut.String   `idl:"name:pbstrPolicyInfo" json:"policy_info"`
	// Return: The GetPolicyInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPolicyInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPolicyInfoOperation) *xxx_GetPolicyInfoOperation {
	if op == nil {
		op = &xxx_GetPolicyInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PolicyInfo = o.PolicyInfo
	op.Return = o.Return
	return op
}

func (o *GetPolicyInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPolicyInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PolicyInfo = op.PolicyInfo
	o.Return = op.Return
}
func (o *GetPolicyInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPolicyInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPolicyInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreatePolicyOperation structure represents the CreatePolicy operation
type xxx_CreatePolicyOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyInfo *oaut.String   `idl:"name:bstrPolicyInfo" json:"policy_info"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreatePolicyOperation) OpNum() int { return 8 }

func (o *xxx_CreatePolicyOperation) OpName() string { return "/IWRMPolicy/v0/CreatePolicy" }

func (o *xxx_CreatePolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPolicyInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyInfo != nil {
			_ptr_bstrPolicyInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyInfo != nil {
					if err := o.PolicyInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyInfo, _ptr_bstrPolicyInfo); err != nil {
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

func (o *xxx_CreatePolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPolicyInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyInfo == nil {
				o.PolicyInfo = &oaut.String{}
			}
			if err := o.PolicyInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyInfo := func(ptr interface{}) { o.PolicyInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyInfo, _s_bstrPolicyInfo, _ptr_bstrPolicyInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreatePolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreatePolicyRequest structure represents the CreatePolicy operation request
type CreatePolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	PolicyInfo *oaut.String   `idl:"name:bstrPolicyInfo" json:"policy_info"`
}

func (o *CreatePolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_CreatePolicyOperation) *xxx_CreatePolicyOperation {
	if op == nil {
		op = &xxx_CreatePolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PolicyInfo = o.PolicyInfo
	return op
}

func (o *CreatePolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_CreatePolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PolicyInfo = op.PolicyInfo
}
func (o *CreatePolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreatePolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreatePolicyResponse structure represents the CreatePolicy operation response
type CreatePolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreatePolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreatePolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_CreatePolicyOperation) *xxx_CreatePolicyOperation {
	if op == nil {
		op = &xxx_CreatePolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreatePolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_CreatePolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreatePolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreatePolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyPolicyOperation structure represents the ModifyPolicy operation
type xxx_ModifyPolicyOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyInfo *oaut.String   `idl:"name:bstrPolicyInfo" json:"policy_info"`
	Overwrite  bool           `idl:"name:bOverwrite" json:"overwrite"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyPolicyOperation) OpNum() int { return 9 }

func (o *xxx_ModifyPolicyOperation) OpName() string { return "/IWRMPolicy/v0/ModifyPolicy" }

func (o *xxx_ModifyPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPolicyInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyInfo != nil {
			_ptr_bstrPolicyInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyInfo != nil {
					if err := o.PolicyInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyInfo, _ptr_bstrPolicyInfo); err != nil {
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
	// bOverwrite {in} (1:{alias=BOOL}(int32))
	{
		if !o.Overwrite {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ModifyPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPolicyInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyInfo == nil {
				o.PolicyInfo = &oaut.String{}
			}
			if err := o.PolicyInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyInfo := func(ptr interface{}) { o.PolicyInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyInfo, _s_bstrPolicyInfo, _ptr_bstrPolicyInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bOverwrite {in} (1:{alias=BOOL}(int32))
	{
		var _bOverwrite int32
		if err := w.ReadData(&_bOverwrite); err != nil {
			return err
		}
		o.Overwrite = _bOverwrite != 0
	}
	return nil
}

func (o *xxx_ModifyPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyPolicyRequest structure represents the ModifyPolicy operation request
type ModifyPolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	PolicyInfo *oaut.String   `idl:"name:bstrPolicyInfo" json:"policy_info"`
	Overwrite  bool           `idl:"name:bOverwrite" json:"overwrite"`
}

func (o *ModifyPolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyPolicyOperation) *xxx_ModifyPolicyOperation {
	if op == nil {
		op = &xxx_ModifyPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PolicyInfo = o.PolicyInfo
	op.Overwrite = o.Overwrite
	return op
}

func (o *ModifyPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyPolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PolicyInfo = op.PolicyInfo
	o.Overwrite = op.Overwrite
}
func (o *ModifyPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyPolicyResponse structure represents the ModifyPolicy operation response
type ModifyPolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyPolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyPolicyOperation) *xxx_ModifyPolicyOperation {
	if op == nil {
		op = &xxx_ModifyPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ModifyPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyPolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeletePolicyOperation structure represents the DeletePolicy operation
type xxx_DeletePolicyOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyName *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeletePolicyOperation) OpNum() int { return 10 }

func (o *xxx_DeletePolicyOperation) OpName() string { return "/IWRMPolicy/v0/DeletePolicy" }

func (o *xxx_DeletePolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyName != nil {
			_ptr_bstrPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyName != nil {
					if err := o.PolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyName, _ptr_bstrPolicyName); err != nil {
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

func (o *xxx_DeletePolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyName == nil {
				o.PolicyName = &oaut.String{}
			}
			if err := o.PolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyName, _s_bstrPolicyName, _ptr_bstrPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeletePolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeletePolicyRequest structure represents the DeletePolicy operation request
type DeletePolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	PolicyName *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
}

func (o *DeletePolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_DeletePolicyOperation) *xxx_DeletePolicyOperation {
	if op == nil {
		op = &xxx_DeletePolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PolicyName = o.PolicyName
	return op
}

func (o *DeletePolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_DeletePolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PolicyName = op.PolicyName
}
func (o *DeletePolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeletePolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeletePolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeletePolicyResponse structure represents the DeletePolicy operation response
type DeletePolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeletePolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeletePolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_DeletePolicyOperation) *xxx_DeletePolicyOperation {
	if op == nil {
		op = &xxx_DeletePolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeletePolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_DeletePolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeletePolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeletePolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeletePolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RenameAllocationPolicyOperation structure represents the RenameAllocationPolicy operation
type xxx_RenameAllocationPolicyOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	NewPolicyName *oaut.String   `idl:"name:bstrNewPolicyName" json:"new_policy_name"`
	OldPolicyName *oaut.String   `idl:"name:bstrOldPolicyName" json:"old_policy_name"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RenameAllocationPolicyOperation) OpNum() int { return 11 }

func (o *xxx_RenameAllocationPolicyOperation) OpName() string {
	return "/IWRMPolicy/v0/RenameAllocationPolicy"
}

func (o *xxx_RenameAllocationPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameAllocationPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrNewPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NewPolicyName != nil {
			_ptr_bstrNewPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewPolicyName != nil {
					if err := o.NewPolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewPolicyName, _ptr_bstrNewPolicyName); err != nil {
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
	// bstrOldPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OldPolicyName != nil {
			_ptr_bstrOldPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldPolicyName != nil {
					if err := o.OldPolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldPolicyName, _ptr_bstrOldPolicyName); err != nil {
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

func (o *xxx_RenameAllocationPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrNewPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrNewPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewPolicyName == nil {
				o.NewPolicyName = &oaut.String{}
			}
			if err := o.NewPolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrNewPolicyName := func(ptr interface{}) { o.NewPolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NewPolicyName, _s_bstrNewPolicyName, _ptr_bstrNewPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrOldPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrOldPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldPolicyName == nil {
				o.OldPolicyName = &oaut.String{}
			}
			if err := o.OldPolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrOldPolicyName := func(ptr interface{}) { o.OldPolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OldPolicyName, _s_bstrOldPolicyName, _ptr_bstrOldPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameAllocationPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameAllocationPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RenameAllocationPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RenameAllocationPolicyRequest structure represents the RenameAllocationPolicy operation request
type RenameAllocationPolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	NewPolicyName *oaut.String   `idl:"name:bstrNewPolicyName" json:"new_policy_name"`
	OldPolicyName *oaut.String   `idl:"name:bstrOldPolicyName" json:"old_policy_name"`
}

func (o *RenameAllocationPolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_RenameAllocationPolicyOperation) *xxx_RenameAllocationPolicyOperation {
	if op == nil {
		op = &xxx_RenameAllocationPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.NewPolicyName = o.NewPolicyName
	op.OldPolicyName = o.OldPolicyName
	return op
}

func (o *RenameAllocationPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_RenameAllocationPolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NewPolicyName = op.NewPolicyName
	o.OldPolicyName = op.OldPolicyName
}
func (o *RenameAllocationPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RenameAllocationPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameAllocationPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RenameAllocationPolicyResponse structure represents the RenameAllocationPolicy operation response
type RenameAllocationPolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RenameAllocationPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RenameAllocationPolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_RenameAllocationPolicyOperation) *xxx_RenameAllocationPolicyOperation {
	if op == nil {
		op = &xxx_RenameAllocationPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RenameAllocationPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_RenameAllocationPolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RenameAllocationPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RenameAllocationPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameAllocationPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MoveBeforeOperation structure represents the MoveBefore operation
type xxx_MoveBeforeOperation struct {
	This                       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyName                 *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
	ResourceGroupName          *oaut.String   `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
	ReferenceResourceGroupName *oaut.String   `idl:"name:bstrRefResourceGroupName" json:"reference_resource_group_name"`
	Return                     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MoveBeforeOperation) OpNum() int { return 12 }

func (o *xxx_MoveBeforeOperation) OpName() string { return "/IWRMPolicy/v0/MoveBefore" }

func (o *xxx_MoveBeforeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveBeforeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyName != nil {
			_ptr_bstrPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyName != nil {
					if err := o.PolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyName, _ptr_bstrPolicyName); err != nil {
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
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupName != nil {
			_ptr_bstrResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupName != nil {
					if err := o.ResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
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
	// bstrRefResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReferenceResourceGroupName != nil {
			_ptr_bstrRefResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferenceResourceGroupName != nil {
					if err := o.ReferenceResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferenceResourceGroupName, _ptr_bstrRefResourceGroupName); err != nil {
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

func (o *xxx_MoveBeforeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyName == nil {
				o.PolicyName = &oaut.String{}
			}
			if err := o.PolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyName, _s_bstrPolicyName, _ptr_bstrPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupName == nil {
				o.ResourceGroupName = &oaut.String{}
			}
			if err := o.ResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrResourceGroupName := func(ptr interface{}) { o.ResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupName, _s_bstrResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrRefResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrRefResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferenceResourceGroupName == nil {
				o.ReferenceResourceGroupName = &oaut.String{}
			}
			if err := o.ReferenceResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrRefResourceGroupName := func(ptr interface{}) { o.ReferenceResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReferenceResourceGroupName, _s_bstrRefResourceGroupName, _ptr_bstrRefResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveBeforeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveBeforeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MoveBeforeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// MoveBeforeRequest structure represents the MoveBefore operation request
type MoveBeforeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                       *dcom.ORPCThis `idl:"name:This" json:"this"`
	PolicyName                 *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
	ResourceGroupName          *oaut.String   `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
	ReferenceResourceGroupName *oaut.String   `idl:"name:bstrRefResourceGroupName" json:"reference_resource_group_name"`
}

func (o *MoveBeforeRequest) xxx_ToOp(ctx context.Context, op *xxx_MoveBeforeOperation) *xxx_MoveBeforeOperation {
	if op == nil {
		op = &xxx_MoveBeforeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PolicyName = o.PolicyName
	op.ResourceGroupName = o.ResourceGroupName
	op.ReferenceResourceGroupName = o.ReferenceResourceGroupName
	return op
}

func (o *MoveBeforeRequest) xxx_FromOp(ctx context.Context, op *xxx_MoveBeforeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PolicyName = op.PolicyName
	o.ResourceGroupName = op.ResourceGroupName
	o.ReferenceResourceGroupName = op.ReferenceResourceGroupName
}
func (o *MoveBeforeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MoveBeforeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveBeforeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MoveBeforeResponse structure represents the MoveBefore operation response
type MoveBeforeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MoveBefore return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MoveBeforeResponse) xxx_ToOp(ctx context.Context, op *xxx_MoveBeforeOperation) *xxx_MoveBeforeOperation {
	if op == nil {
		op = &xxx_MoveBeforeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *MoveBeforeResponse) xxx_FromOp(ctx context.Context, op *xxx_MoveBeforeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *MoveBeforeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MoveBeforeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveBeforeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MoveAfterOperation structure represents the MoveAfter operation
type xxx_MoveAfterOperation struct {
	This                       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyName                 *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
	ResourceGroupName          *oaut.String   `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
	ReferenceResourceGroupName *oaut.String   `idl:"name:bstrRefResourceGroupName" json:"reference_resource_group_name"`
	Return                     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MoveAfterOperation) OpNum() int { return 13 }

func (o *xxx_MoveAfterOperation) OpName() string { return "/IWRMPolicy/v0/MoveAfter" }

func (o *xxx_MoveAfterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveAfterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyName != nil {
			_ptr_bstrPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyName != nil {
					if err := o.PolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyName, _ptr_bstrPolicyName); err != nil {
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
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupName != nil {
			_ptr_bstrResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupName != nil {
					if err := o.ResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
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
	// bstrRefResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReferenceResourceGroupName != nil {
			_ptr_bstrRefResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferenceResourceGroupName != nil {
					if err := o.ReferenceResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferenceResourceGroupName, _ptr_bstrRefResourceGroupName); err != nil {
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

func (o *xxx_MoveAfterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyName == nil {
				o.PolicyName = &oaut.String{}
			}
			if err := o.PolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyName, _s_bstrPolicyName, _ptr_bstrPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupName == nil {
				o.ResourceGroupName = &oaut.String{}
			}
			if err := o.ResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrResourceGroupName := func(ptr interface{}) { o.ResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupName, _s_bstrResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrRefResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrRefResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferenceResourceGroupName == nil {
				o.ReferenceResourceGroupName = &oaut.String{}
			}
			if err := o.ReferenceResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrRefResourceGroupName := func(ptr interface{}) { o.ReferenceResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReferenceResourceGroupName, _s_bstrRefResourceGroupName, _ptr_bstrRefResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveAfterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveAfterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MoveAfterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// MoveAfterRequest structure represents the MoveAfter operation request
type MoveAfterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                       *dcom.ORPCThis `idl:"name:This" json:"this"`
	PolicyName                 *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
	ResourceGroupName          *oaut.String   `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
	ReferenceResourceGroupName *oaut.String   `idl:"name:bstrRefResourceGroupName" json:"reference_resource_group_name"`
}

func (o *MoveAfterRequest) xxx_ToOp(ctx context.Context, op *xxx_MoveAfterOperation) *xxx_MoveAfterOperation {
	if op == nil {
		op = &xxx_MoveAfterOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PolicyName = o.PolicyName
	op.ResourceGroupName = o.ResourceGroupName
	op.ReferenceResourceGroupName = o.ReferenceResourceGroupName
	return op
}

func (o *MoveAfterRequest) xxx_FromOp(ctx context.Context, op *xxx_MoveAfterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PolicyName = op.PolicyName
	o.ResourceGroupName = op.ResourceGroupName
	o.ReferenceResourceGroupName = op.ReferenceResourceGroupName
}
func (o *MoveAfterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MoveAfterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveAfterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MoveAfterResponse structure represents the MoveAfter operation response
type MoveAfterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MoveAfter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MoveAfterResponse) xxx_ToOp(ctx context.Context, op *xxx_MoveAfterOperation) *xxx_MoveAfterOperation {
	if op == nil {
		op = &xxx_MoveAfterOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *MoveAfterResponse) xxx_FromOp(ctx context.Context, op *xxx_MoveAfterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *MoveAfterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MoveAfterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveAfterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCALDefaultPolicyNameOperation structure represents the SetCalDefaultPolicyName operation
type xxx_SetCALDefaultPolicyNameOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	DefaultPolicyName *oaut.String   `idl:"name:bstrDefaultPolicyName" json:"default_policy_name"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCALDefaultPolicyNameOperation) OpNum() int { return 14 }

func (o *xxx_SetCALDefaultPolicyNameOperation) OpName() string {
	return "/IWRMPolicy/v0/SetCalDefaultPolicyName"
}

func (o *xxx_SetCALDefaultPolicyNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCALDefaultPolicyNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrDefaultPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DefaultPolicyName != nil {
			_ptr_bstrDefaultPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DefaultPolicyName != nil {
					if err := o.DefaultPolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DefaultPolicyName, _ptr_bstrDefaultPolicyName); err != nil {
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

func (o *xxx_SetCALDefaultPolicyNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrDefaultPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrDefaultPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DefaultPolicyName == nil {
				o.DefaultPolicyName = &oaut.String{}
			}
			if err := o.DefaultPolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrDefaultPolicyName := func(ptr interface{}) { o.DefaultPolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DefaultPolicyName, _s_bstrDefaultPolicyName, _ptr_bstrDefaultPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCALDefaultPolicyNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCALDefaultPolicyNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCALDefaultPolicyNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCALDefaultPolicyNameRequest structure represents the SetCalDefaultPolicyName operation request
type SetCALDefaultPolicyNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	DefaultPolicyName *oaut.String   `idl:"name:bstrDefaultPolicyName" json:"default_policy_name"`
}

func (o *SetCALDefaultPolicyNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetCALDefaultPolicyNameOperation) *xxx_SetCALDefaultPolicyNameOperation {
	if op == nil {
		op = &xxx_SetCALDefaultPolicyNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DefaultPolicyName = o.DefaultPolicyName
	return op
}

func (o *SetCALDefaultPolicyNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCALDefaultPolicyNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DefaultPolicyName = op.DefaultPolicyName
}
func (o *SetCALDefaultPolicyNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetCALDefaultPolicyNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCALDefaultPolicyNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCALDefaultPolicyNameResponse structure represents the SetCalDefaultPolicyName operation response
type SetCALDefaultPolicyNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetCalDefaultPolicyName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCALDefaultPolicyNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetCALDefaultPolicyNameOperation) *xxx_SetCALDefaultPolicyNameOperation {
	if op == nil {
		op = &xxx_SetCALDefaultPolicyNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetCALDefaultPolicyNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCALDefaultPolicyNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCALDefaultPolicyNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetCALDefaultPolicyNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCALDefaultPolicyNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCALDefaultPolicyNameOperation structure represents the GetCalDefaultPolicyName operation
type xxx_GetCALDefaultPolicyNameOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	DefaultPolicyName *oaut.String   `idl:"name:pbstrDefaultPolicyName" json:"default_policy_name"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCALDefaultPolicyNameOperation) OpNum() int { return 15 }

func (o *xxx_GetCALDefaultPolicyNameOperation) OpName() string {
	return "/IWRMPolicy/v0/GetCalDefaultPolicyName"
}

func (o *xxx_GetCALDefaultPolicyNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCALDefaultPolicyNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCALDefaultPolicyNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCALDefaultPolicyNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCALDefaultPolicyNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrDefaultPolicyName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DefaultPolicyName != nil {
			_ptr_pbstrDefaultPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DefaultPolicyName != nil {
					if err := o.DefaultPolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DefaultPolicyName, _ptr_pbstrDefaultPolicyName); err != nil {
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

func (o *xxx_GetCALDefaultPolicyNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrDefaultPolicyName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrDefaultPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DefaultPolicyName == nil {
				o.DefaultPolicyName = &oaut.String{}
			}
			if err := o.DefaultPolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrDefaultPolicyName := func(ptr interface{}) { o.DefaultPolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DefaultPolicyName, _s_pbstrDefaultPolicyName, _ptr_pbstrDefaultPolicyName); err != nil {
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

// GetCALDefaultPolicyNameRequest structure represents the GetCalDefaultPolicyName operation request
type GetCALDefaultPolicyNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCALDefaultPolicyNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCALDefaultPolicyNameOperation) *xxx_GetCALDefaultPolicyNameOperation {
	if op == nil {
		op = &xxx_GetCALDefaultPolicyNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCALDefaultPolicyNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCALDefaultPolicyNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCALDefaultPolicyNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCALDefaultPolicyNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCALDefaultPolicyNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCALDefaultPolicyNameResponse structure represents the GetCalDefaultPolicyName operation response
type GetCALDefaultPolicyNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	DefaultPolicyName *oaut.String   `idl:"name:pbstrDefaultPolicyName" json:"default_policy_name"`
	// Return: The GetCalDefaultPolicyName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCALDefaultPolicyNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCALDefaultPolicyNameOperation) *xxx_GetCALDefaultPolicyNameOperation {
	if op == nil {
		op = &xxx_GetCALDefaultPolicyNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DefaultPolicyName = o.DefaultPolicyName
	op.Return = o.Return
	return op
}

func (o *GetCALDefaultPolicyNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCALDefaultPolicyNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DefaultPolicyName = op.DefaultPolicyName
	o.Return = op.Return
}
func (o *GetCALDefaultPolicyNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCALDefaultPolicyNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCALDefaultPolicyNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetProcessListOperation structure represents the GetProcessList operation
type xxx_GetProcessListOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyName  *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
	ProcessList *oaut.String   `idl:"name:pbstrProcessList" json:"process_list"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetProcessListOperation) OpNum() int { return 16 }

func (o *xxx_GetProcessListOperation) OpName() string { return "/IWRMPolicy/v0/GetProcessList" }

func (o *xxx_GetProcessListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProcessListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyName != nil {
			_ptr_bstrPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyName != nil {
					if err := o.PolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyName, _ptr_bstrPolicyName); err != nil {
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

func (o *xxx_GetProcessListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyName == nil {
				o.PolicyName = &oaut.String{}
			}
			if err := o.PolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyName, _s_bstrPolicyName, _ptr_bstrPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProcessListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProcessListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrProcessList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ProcessList != nil {
			_ptr_pbstrProcessList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ProcessList != nil {
					if err := o.ProcessList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProcessList, _ptr_pbstrProcessList); err != nil {
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

func (o *xxx_GetProcessListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrProcessList {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrProcessList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ProcessList == nil {
				o.ProcessList = &oaut.String{}
			}
			if err := o.ProcessList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrProcessList := func(ptr interface{}) { o.ProcessList = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ProcessList, _s_pbstrProcessList, _ptr_pbstrProcessList); err != nil {
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

// GetProcessListRequest structure represents the GetProcessList operation request
type GetProcessListRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	PolicyName *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
}

func (o *GetProcessListRequest) xxx_ToOp(ctx context.Context, op *xxx_GetProcessListOperation) *xxx_GetProcessListOperation {
	if op == nil {
		op = &xxx_GetProcessListOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PolicyName = o.PolicyName
	return op
}

func (o *GetProcessListRequest) xxx_FromOp(ctx context.Context, op *xxx_GetProcessListOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PolicyName = op.PolicyName
}
func (o *GetProcessListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetProcessListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProcessListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetProcessListResponse structure represents the GetProcessList operation response
type GetProcessListResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ProcessList *oaut.String   `idl:"name:pbstrProcessList" json:"process_list"`
	// Return: The GetProcessList return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetProcessListResponse) xxx_ToOp(ctx context.Context, op *xxx_GetProcessListOperation) *xxx_GetProcessListOperation {
	if op == nil {
		op = &xxx_GetProcessListOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ProcessList = o.ProcessList
	op.Return = o.Return
	return op
}

func (o *GetProcessListResponse) xxx_FromOp(ctx context.Context, op *xxx_GetProcessListOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ProcessList = op.ProcessList
	o.Return = op.Return
}
func (o *GetProcessListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetProcessListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProcessListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCurrentPolicyOperation structure represents the GetCurrentPolicy operation
type xxx_GetCurrentPolicyOperation struct {
	This              *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat      `idl:"name:That" json:"that"`
	CurrentPolicyInfo *oaut.String        `idl:"name:pbstrCurrentPolicyInfo" json:"current_policy_info"`
	EnumManage        wsrm.ManagementType `idl:"name:enumManage" json:"enum_manage"`
	Return            int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCurrentPolicyOperation) OpNum() int { return 17 }

func (o *xxx_GetCurrentPolicyOperation) OpName() string { return "/IWRMPolicy/v0/GetCurrentPolicy" }

func (o *xxx_GetCurrentPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCurrentPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCurrentPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCurrentPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCurrentPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrCurrentPolicyInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CurrentPolicyInfo != nil {
			_ptr_pbstrCurrentPolicyInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CurrentPolicyInfo != nil {
					if err := o.CurrentPolicyInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CurrentPolicyInfo, _ptr_pbstrCurrentPolicyInfo); err != nil {
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
	// enumManage {out} (1:{pointer=ref}*(1))(2:{v1_enum, alias=MANAGEMENT_TYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumManage)); err != nil {
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

func (o *xxx_GetCurrentPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrCurrentPolicyInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrCurrentPolicyInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CurrentPolicyInfo == nil {
				o.CurrentPolicyInfo = &oaut.String{}
			}
			if err := o.CurrentPolicyInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrCurrentPolicyInfo := func(ptr interface{}) { o.CurrentPolicyInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CurrentPolicyInfo, _s_pbstrCurrentPolicyInfo, _ptr_pbstrCurrentPolicyInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// enumManage {out} (1:{pointer=ref}*(1))(2:{v1_enum, alias=MANAGEMENT_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumManage)); err != nil {
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

// GetCurrentPolicyRequest structure represents the GetCurrentPolicy operation request
type GetCurrentPolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCurrentPolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCurrentPolicyOperation) *xxx_GetCurrentPolicyOperation {
	if op == nil {
		op = &xxx_GetCurrentPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCurrentPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCurrentPolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCurrentPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCurrentPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCurrentPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCurrentPolicyResponse structure represents the GetCurrentPolicy operation response
type GetCurrentPolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That              *dcom.ORPCThat      `idl:"name:That" json:"that"`
	CurrentPolicyInfo *oaut.String        `idl:"name:pbstrCurrentPolicyInfo" json:"current_policy_info"`
	EnumManage        wsrm.ManagementType `idl:"name:enumManage" json:"enum_manage"`
	// Return: The GetCurrentPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCurrentPolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCurrentPolicyOperation) *xxx_GetCurrentPolicyOperation {
	if op == nil {
		op = &xxx_GetCurrentPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CurrentPolicyInfo = o.CurrentPolicyInfo
	op.EnumManage = o.EnumManage
	op.Return = o.Return
	return op
}

func (o *GetCurrentPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCurrentPolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CurrentPolicyInfo = op.CurrentPolicyInfo
	o.EnumManage = op.EnumManage
	o.Return = op.Return
}
func (o *GetCurrentPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCurrentPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCurrentPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCurrentPolicyOperation structure represents the SetCurrentPolicy operation
type xxx_SetCurrentPolicyOperation struct {
	This       *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat      `idl:"name:That" json:"that"`
	PolicyName *oaut.String        `idl:"name:bstrPolicyName" json:"policy_name"`
	EnumManage wsrm.ManagementType `idl:"name:enumManage" json:"enum_manage"`
	Return     int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCurrentPolicyOperation) OpNum() int { return 18 }

func (o *xxx_SetCurrentPolicyOperation) OpName() string { return "/IWRMPolicy/v0/SetCurrentPolicy" }

func (o *xxx_SetCurrentPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCurrentPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyName != nil {
			_ptr_bstrPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyName != nil {
					if err := o.PolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyName, _ptr_bstrPolicyName); err != nil {
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
	// enumManage {in} (1:{v1_enum, alias=MANAGEMENT_TYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumManage)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCurrentPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyName == nil {
				o.PolicyName = &oaut.String{}
			}
			if err := o.PolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyName, _s_bstrPolicyName, _ptr_bstrPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// enumManage {in} (1:{v1_enum, alias=MANAGEMENT_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumManage)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCurrentPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCurrentPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCurrentPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCurrentPolicyRequest structure represents the SetCurrentPolicy operation request
type SetCurrentPolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis      `idl:"name:This" json:"this"`
	PolicyName *oaut.String        `idl:"name:bstrPolicyName" json:"policy_name"`
	EnumManage wsrm.ManagementType `idl:"name:enumManage" json:"enum_manage"`
}

func (o *SetCurrentPolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_SetCurrentPolicyOperation) *xxx_SetCurrentPolicyOperation {
	if op == nil {
		op = &xxx_SetCurrentPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PolicyName = o.PolicyName
	op.EnumManage = o.EnumManage
	return op
}

func (o *SetCurrentPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCurrentPolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PolicyName = op.PolicyName
	o.EnumManage = op.EnumManage
}
func (o *SetCurrentPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetCurrentPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCurrentPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCurrentPolicyResponse structure represents the SetCurrentPolicy operation response
type SetCurrentPolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetCurrentPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCurrentPolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_SetCurrentPolicyOperation) *xxx_SetCurrentPolicyOperation {
	if op == nil {
		op = &xxx_SetCurrentPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetCurrentPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCurrentPolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCurrentPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetCurrentPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCurrentPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCurrentStateAndActivePolicyNameOperation structure represents the GetCurrentStateAndActivePolicyName operation
type xxx_GetCurrentStateAndActivePolicyNameOperation struct {
	This              *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat      `idl:"name:That" json:"that"`
	CurrentPolicyName *oaut.String        `idl:"name:pbstrCurrentPolicyName" json:"current_policy_name"`
	EnumManage        wsrm.ManagementType `idl:"name:enumManage" json:"enum_manage"`
	Return            int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCurrentStateAndActivePolicyNameOperation) OpNum() int { return 19 }

func (o *xxx_GetCurrentStateAndActivePolicyNameOperation) OpName() string {
	return "/IWRMPolicy/v0/GetCurrentStateAndActivePolicyName"
}

func (o *xxx_GetCurrentStateAndActivePolicyNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCurrentStateAndActivePolicyNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCurrentStateAndActivePolicyNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCurrentStateAndActivePolicyNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCurrentStateAndActivePolicyNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrCurrentPolicyName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CurrentPolicyName != nil {
			_ptr_pbstrCurrentPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CurrentPolicyName != nil {
					if err := o.CurrentPolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CurrentPolicyName, _ptr_pbstrCurrentPolicyName); err != nil {
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
	// enumManage {out} (1:{pointer=ref}*(1))(2:{v1_enum, alias=MANAGEMENT_TYPE}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumManage)); err != nil {
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

func (o *xxx_GetCurrentStateAndActivePolicyNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrCurrentPolicyName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrCurrentPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CurrentPolicyName == nil {
				o.CurrentPolicyName = &oaut.String{}
			}
			if err := o.CurrentPolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrCurrentPolicyName := func(ptr interface{}) { o.CurrentPolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CurrentPolicyName, _s_pbstrCurrentPolicyName, _ptr_pbstrCurrentPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// enumManage {out} (1:{pointer=ref}*(1))(2:{v1_enum, alias=MANAGEMENT_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumManage)); err != nil {
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

// GetCurrentStateAndActivePolicyNameRequest structure represents the GetCurrentStateAndActivePolicyName operation request
type GetCurrentStateAndActivePolicyNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCurrentStateAndActivePolicyNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCurrentStateAndActivePolicyNameOperation) *xxx_GetCurrentStateAndActivePolicyNameOperation {
	if op == nil {
		op = &xxx_GetCurrentStateAndActivePolicyNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCurrentStateAndActivePolicyNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCurrentStateAndActivePolicyNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCurrentStateAndActivePolicyNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCurrentStateAndActivePolicyNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCurrentStateAndActivePolicyNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCurrentStateAndActivePolicyNameResponse structure represents the GetCurrentStateAndActivePolicyName operation response
type GetCurrentStateAndActivePolicyNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That              *dcom.ORPCThat      `idl:"name:That" json:"that"`
	CurrentPolicyName *oaut.String        `idl:"name:pbstrCurrentPolicyName" json:"current_policy_name"`
	EnumManage        wsrm.ManagementType `idl:"name:enumManage" json:"enum_manage"`
	// Return: The GetCurrentStateAndActivePolicyName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCurrentStateAndActivePolicyNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCurrentStateAndActivePolicyNameOperation) *xxx_GetCurrentStateAndActivePolicyNameOperation {
	if op == nil {
		op = &xxx_GetCurrentStateAndActivePolicyNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CurrentPolicyName = o.CurrentPolicyName
	op.EnumManage = o.EnumManage
	op.Return = o.Return
	return op
}

func (o *GetCurrentStateAndActivePolicyNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCurrentStateAndActivePolicyNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CurrentPolicyName = op.CurrentPolicyName
	o.EnumManage = op.EnumManage
	o.Return = op.Return
}
func (o *GetCurrentStateAndActivePolicyNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCurrentStateAndActivePolicyNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCurrentStateAndActivePolicyNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetConditionalPolicyOperation structure represents the GetConditionalPolicy operation
type xxx_GetConditionalPolicyOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyName *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
	PolicyInfo *oaut.String   `idl:"name:pbstrPolicyInfo" json:"policy_info"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetConditionalPolicyOperation) OpNum() int { return 20 }

func (o *xxx_GetConditionalPolicyOperation) OpName() string {
	return "/IWRMPolicy/v0/GetConditionalPolicy"
}

func (o *xxx_GetConditionalPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConditionalPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyName != nil {
			_ptr_bstrPolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyName != nil {
					if err := o.PolicyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyName, _ptr_bstrPolicyName); err != nil {
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

func (o *xxx_GetConditionalPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPolicyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyName == nil {
				o.PolicyName = &oaut.String{}
			}
			if err := o.PolicyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyName, _s_bstrPolicyName, _ptr_bstrPolicyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConditionalPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConditionalPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrPolicyInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyInfo != nil {
			_ptr_pbstrPolicyInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyInfo != nil {
					if err := o.PolicyInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyInfo, _ptr_pbstrPolicyInfo); err != nil {
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

func (o *xxx_GetConditionalPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrPolicyInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPolicyInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyInfo == nil {
				o.PolicyInfo = &oaut.String{}
			}
			if err := o.PolicyInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPolicyInfo := func(ptr interface{}) { o.PolicyInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyInfo, _s_pbstrPolicyInfo, _ptr_pbstrPolicyInfo); err != nil {
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

// GetConditionalPolicyRequest structure represents the GetConditionalPolicy operation request
type GetConditionalPolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	PolicyName *oaut.String   `idl:"name:bstrPolicyName" json:"policy_name"`
}

func (o *GetConditionalPolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_GetConditionalPolicyOperation) *xxx_GetConditionalPolicyOperation {
	if op == nil {
		op = &xxx_GetConditionalPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PolicyName = o.PolicyName
	return op
}

func (o *GetConditionalPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConditionalPolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PolicyName = op.PolicyName
}
func (o *GetConditionalPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetConditionalPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConditionalPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetConditionalPolicyResponse structure represents the GetConditionalPolicy operation response
type GetConditionalPolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyInfo *oaut.String   `idl:"name:pbstrPolicyInfo" json:"policy_info"`
	// Return: The GetConditionalPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetConditionalPolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_GetConditionalPolicyOperation) *xxx_GetConditionalPolicyOperation {
	if op == nil {
		op = &xxx_GetConditionalPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PolicyInfo = o.PolicyInfo
	op.Return = o.Return
	return op
}

func (o *GetConditionalPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConditionalPolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PolicyInfo = op.PolicyInfo
	o.Return = op.Return
}
func (o *GetConditionalPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetConditionalPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConditionalPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetConditionalPolicyOperation structure represents the SetConditionalPolicy operation
type xxx_SetConditionalPolicyOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	PolicyInfo *oaut.String   `idl:"name:bstrPolicyInfo" json:"policy_info"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetConditionalPolicyOperation) OpNum() int { return 21 }

func (o *xxx_SetConditionalPolicyOperation) OpName() string {
	return "/IWRMPolicy/v0/SetConditionalPolicy"
}

func (o *xxx_SetConditionalPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConditionalPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPolicyInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PolicyInfo != nil {
			_ptr_bstrPolicyInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PolicyInfo != nil {
					if err := o.PolicyInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyInfo, _ptr_bstrPolicyInfo); err != nil {
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

func (o *xxx_SetConditionalPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPolicyInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPolicyInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyInfo == nil {
				o.PolicyInfo = &oaut.String{}
			}
			if err := o.PolicyInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPolicyInfo := func(ptr interface{}) { o.PolicyInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PolicyInfo, _s_bstrPolicyInfo, _ptr_bstrPolicyInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConditionalPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConditionalPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetConditionalPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetConditionalPolicyRequest structure represents the SetConditionalPolicy operation request
type SetConditionalPolicyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	PolicyInfo *oaut.String   `idl:"name:bstrPolicyInfo" json:"policy_info"`
}

func (o *SetConditionalPolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_SetConditionalPolicyOperation) *xxx_SetConditionalPolicyOperation {
	if op == nil {
		op = &xxx_SetConditionalPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PolicyInfo = o.PolicyInfo
	return op
}

func (o *SetConditionalPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetConditionalPolicyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PolicyInfo = op.PolicyInfo
}
func (o *SetConditionalPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetConditionalPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetConditionalPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetConditionalPolicyResponse structure represents the SetConditionalPolicy operation response
type SetConditionalPolicyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetConditionalPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetConditionalPolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_SetConditionalPolicyOperation) *xxx_SetConditionalPolicyOperation {
	if op == nil {
		op = &xxx_SetConditionalPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetConditionalPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetConditionalPolicyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetConditionalPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetConditionalPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetConditionalPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
