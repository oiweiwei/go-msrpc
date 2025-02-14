package ifsrmpipelinemoduledefinition

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
	// IFsrmPipelineModuleDefinition interface identifier 515c1277-2c81-440e-8fcf-367921ed4f59
	PipelineModuleDefinitionIID = &dcom.IID{Data1: 0x515c1277, Data2: 0x2c81, Data3: 0x440e, Data4: []byte{0x8f, 0xcf, 0x36, 0x79, 0x21, 0xed, 0x4f, 0x59}}
	// Syntax UUID
	PipelineModuleDefinitionSyntaxUUID = &uuid.UUID{TimeLow: 0x515c1277, TimeMid: 0x2c81, TimeHiAndVersion: 0x440e, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xcf, Node: [6]uint8{0x36, 0x79, 0x21, 0xed, 0x4f, 0x59}}
	// Syntax ID
	PipelineModuleDefinitionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: PipelineModuleDefinitionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmPipelineModuleDefinition interface.
type PipelineModuleDefinitionClient interface {

	// IFsrmObject retrieval method.
	Object() ifsrmobject.ObjectClient

	// ModuleClsid operation.
	GetModuleClassID(context.Context, *GetModuleClassIDRequest, ...dcerpc.CallOption) (*GetModuleClassIDResponse, error)

	// ModuleClsid operation.
	SetModuleClassID(context.Context, *SetModuleClassIDRequest, ...dcerpc.CallOption) (*SetModuleClassIDResponse, error)

	// Name operation.
	GetName(context.Context, *GetNameRequest, ...dcerpc.CallOption) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest, ...dcerpc.CallOption) (*SetNameResponse, error)

	// Company operation.
	GetCompany(context.Context, *GetCompanyRequest, ...dcerpc.CallOption) (*GetCompanyResponse, error)

	// Company operation.
	SetCompany(context.Context, *SetCompanyRequest, ...dcerpc.CallOption) (*SetCompanyResponse, error)

	// Version operation.
	GetVersion(context.Context, *GetVersionRequest, ...dcerpc.CallOption) (*GetVersionResponse, error)

	// Version operation.
	SetVersion(context.Context, *SetVersionRequest, ...dcerpc.CallOption) (*SetVersionResponse, error)

	// The ModuleType (get) method retrieves the Module type of the module definition as
	// defined in the FsrmPipelineModuleType (section 2.2.1.2.12) enumeration and returns
	// S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------------+
	//	|         RETURN          |                                   |
	//	|       VALUE/CODE        |            DESCRIPTION            |
	//	|                         |                                   |
	//	+-------------------------+-----------------------------------+
	//	+-------------------------+-----------------------------------+
	//	| 0x80070057 E_INVALIDARG | The moduleType parameter is NULL. |
	//	+-------------------------+-----------------------------------+
	GetModuleType(context.Context, *GetModuleTypeRequest, ...dcerpc.CallOption) (*GetModuleTypeResponse, error)

	// Enabled operation.
	GetEnabled(context.Context, *GetEnabledRequest, ...dcerpc.CallOption) (*GetEnabledResponse, error)

	// Enabled operation.
	SetEnabled(context.Context, *SetEnabledRequest, ...dcerpc.CallOption) (*SetEnabledResponse, error)

	// NeedsFileContent operation.
	GetNeedsFileContent(context.Context, *GetNeedsFileContentRequest, ...dcerpc.CallOption) (*GetNeedsFileContentResponse, error)

	// NeedsFileContent operation.
	SetNeedsFileContent(context.Context, *SetNeedsFileContentRequest, ...dcerpc.CallOption) (*SetNeedsFileContentResponse, error)

	// Account operation.
	GetAccount(context.Context, *GetAccountRequest, ...dcerpc.CallOption) (*GetAccountResponse, error)

	// Account operation.
	SetAccount(context.Context, *SetAccountRequest, ...dcerpc.CallOption) (*SetAccountResponse, error)

	// SupportedExtensions operation.
	GetSupportedExtensions(context.Context, *GetSupportedExtensionsRequest, ...dcerpc.CallOption) (*GetSupportedExtensionsResponse, error)

	// SupportedExtensions operation.
	SetSupportedExtensions(context.Context, *SetSupportedExtensionsRequest, ...dcerpc.CallOption) (*SetSupportedExtensionsResponse, error)

	// Parameters operation.
	GetParameters(context.Context, *GetParametersRequest, ...dcerpc.CallOption) (*GetParametersResponse, error)

	// Parameters operation.
	SetParameters(context.Context, *SetParametersRequest, ...dcerpc.CallOption) (*SetParametersResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) PipelineModuleDefinitionClient
}

type xxx_DefaultPipelineModuleDefinitionClient struct {
	ifsrmobject.ObjectClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) Object() ifsrmobject.ObjectClient {
	return o.ObjectClient
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) GetModuleClassID(ctx context.Context, in *GetModuleClassIDRequest, opts ...dcerpc.CallOption) (*GetModuleClassIDResponse, error) {
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
	out := &GetModuleClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) SetModuleClassID(ctx context.Context, in *SetModuleClassIDRequest, opts ...dcerpc.CallOption) (*SetModuleClassIDResponse, error) {
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
	out := &SetModuleClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) GetName(ctx context.Context, in *GetNameRequest, opts ...dcerpc.CallOption) (*GetNameResponse, error) {
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
	out := &GetNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) SetName(ctx context.Context, in *SetNameRequest, opts ...dcerpc.CallOption) (*SetNameResponse, error) {
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
	out := &SetNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...dcerpc.CallOption) (*GetCompanyResponse, error) {
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
	out := &GetCompanyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) SetCompany(ctx context.Context, in *SetCompanyRequest, opts ...dcerpc.CallOption) (*SetCompanyResponse, error) {
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
	out := &SetCompanyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...dcerpc.CallOption) (*GetVersionResponse, error) {
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
	out := &GetVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) SetVersion(ctx context.Context, in *SetVersionRequest, opts ...dcerpc.CallOption) (*SetVersionResponse, error) {
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
	out := &SetVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) GetModuleType(ctx context.Context, in *GetModuleTypeRequest, opts ...dcerpc.CallOption) (*GetModuleTypeResponse, error) {
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
	out := &GetModuleTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) GetEnabled(ctx context.Context, in *GetEnabledRequest, opts ...dcerpc.CallOption) (*GetEnabledResponse, error) {
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
	out := &GetEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) SetEnabled(ctx context.Context, in *SetEnabledRequest, opts ...dcerpc.CallOption) (*SetEnabledResponse, error) {
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
	out := &SetEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) GetNeedsFileContent(ctx context.Context, in *GetNeedsFileContentRequest, opts ...dcerpc.CallOption) (*GetNeedsFileContentResponse, error) {
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
	out := &GetNeedsFileContentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) SetNeedsFileContent(ctx context.Context, in *SetNeedsFileContentRequest, opts ...dcerpc.CallOption) (*SetNeedsFileContentResponse, error) {
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
	out := &SetNeedsFileContentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...dcerpc.CallOption) (*GetAccountResponse, error) {
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
	out := &GetAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) SetAccount(ctx context.Context, in *SetAccountRequest, opts ...dcerpc.CallOption) (*SetAccountResponse, error) {
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
	out := &SetAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) GetSupportedExtensions(ctx context.Context, in *GetSupportedExtensionsRequest, opts ...dcerpc.CallOption) (*GetSupportedExtensionsResponse, error) {
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
	out := &GetSupportedExtensionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) SetSupportedExtensions(ctx context.Context, in *SetSupportedExtensionsRequest, opts ...dcerpc.CallOption) (*SetSupportedExtensionsResponse, error) {
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
	out := &SetSupportedExtensionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) GetParameters(ctx context.Context, in *GetParametersRequest, opts ...dcerpc.CallOption) (*GetParametersResponse, error) {
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
	out := &GetParametersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) SetParameters(ctx context.Context, in *SetParametersRequest, opts ...dcerpc.CallOption) (*SetParametersResponse, error) {
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
	out := &SetParametersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultPipelineModuleDefinitionClient) IPID(ctx context.Context, ipid *dcom.IPID) PipelineModuleDefinitionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPipelineModuleDefinitionClient{
		ObjectClient: o.ObjectClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}

func NewPipelineModuleDefinitionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PipelineModuleDefinitionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PipelineModuleDefinitionSyntaxV0_0))...)
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
	return &xxx_DefaultPipelineModuleDefinitionClient{
		ObjectClient: base,
		cc:           cc,
		ipid:         ipid,
	}, nil
}

// xxx_GetModuleClassIDOperation structure represents the ModuleClsid operation
type xxx_GetModuleClassIDOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ModuleClassID *oaut.String   `idl:"name:moduleClsid" json:"module_class_id"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetModuleClassIDOperation) OpNum() int { return 12 }

func (o *xxx_GetModuleClassIDOperation) OpName() string {
	return "/IFsrmPipelineModuleDefinition/v0/ModuleClsid"
}

func (o *xxx_GetModuleClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetModuleClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetModuleClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// moduleClsid {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ModuleClassID != nil {
			_ptr_moduleClsid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModuleClassID != nil {
					if err := o.ModuleClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleClassID, _ptr_moduleClsid); err != nil {
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

func (o *xxx_GetModuleClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// moduleClsid {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_moduleClsid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModuleClassID == nil {
				o.ModuleClassID = &oaut.String{}
			}
			if err := o.ModuleClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_moduleClsid := func(ptr interface{}) { o.ModuleClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ModuleClassID, _s_moduleClsid, _ptr_moduleClsid); err != nil {
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

// GetModuleClassIDRequest structure represents the ModuleClsid operation request
type GetModuleClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetModuleClassIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetModuleClassIDOperation) *xxx_GetModuleClassIDOperation {
	if op == nil {
		op = &xxx_GetModuleClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetModuleClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetModuleClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetModuleClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetModuleClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetModuleClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetModuleClassIDResponse structure represents the ModuleClsid operation response
type GetModuleClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ModuleClassID *oaut.String   `idl:"name:moduleClsid" json:"module_class_id"`
	// Return: The ModuleClsid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetModuleClassIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetModuleClassIDOperation) *xxx_GetModuleClassIDOperation {
	if op == nil {
		op = &xxx_GetModuleClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ModuleClassID = o.ModuleClassID
	op.Return = o.Return
	return op
}

func (o *GetModuleClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetModuleClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ModuleClassID = op.ModuleClassID
	o.Return = op.Return
}
func (o *GetModuleClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetModuleClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetModuleClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetModuleClassIDOperation structure represents the ModuleClsid operation
type xxx_SetModuleClassIDOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ModuleClassID *oaut.String   `idl:"name:moduleClsid" json:"module_class_id"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetModuleClassIDOperation) OpNum() int { return 13 }

func (o *xxx_SetModuleClassIDOperation) OpName() string {
	return "/IFsrmPipelineModuleDefinition/v0/ModuleClsid"
}

func (o *xxx_SetModuleClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetModuleClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// moduleClsid {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ModuleClassID != nil {
			_ptr_moduleClsid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModuleClassID != nil {
					if err := o.ModuleClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleClassID, _ptr_moduleClsid); err != nil {
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

func (o *xxx_SetModuleClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// moduleClsid {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_moduleClsid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModuleClassID == nil {
				o.ModuleClassID = &oaut.String{}
			}
			if err := o.ModuleClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_moduleClsid := func(ptr interface{}) { o.ModuleClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ModuleClassID, _s_moduleClsid, _ptr_moduleClsid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetModuleClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetModuleClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetModuleClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetModuleClassIDRequest structure represents the ModuleClsid operation request
type SetModuleClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	ModuleClassID *oaut.String   `idl:"name:moduleClsid" json:"module_class_id"`
}

func (o *SetModuleClassIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetModuleClassIDOperation) *xxx_SetModuleClassIDOperation {
	if op == nil {
		op = &xxx_SetModuleClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ModuleClassID = o.ModuleClassID
	return op
}

func (o *SetModuleClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetModuleClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ModuleClassID = op.ModuleClassID
}
func (o *SetModuleClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetModuleClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetModuleClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetModuleClassIDResponse structure represents the ModuleClsid operation response
type SetModuleClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModuleClsid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetModuleClassIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetModuleClassIDOperation) *xxx_SetModuleClassIDOperation {
	if op == nil {
		op = &xxx_SetModuleClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetModuleClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetModuleClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetModuleClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetModuleClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetModuleClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNameOperation structure represents the Name operation
type xxx_GetNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNameOperation) OpNum() int { return 14 }

func (o *xxx_GetNameOperation) OpName() string { return "/IFsrmPipelineModuleDefinition/v0/Name" }

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

func (o *GetNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNameOperation) *xxx_GetNameOperation {
	if op == nil {
		op = &xxx_GetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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

func (o *GetNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNameOperation) *xxx_GetNameOperation {
	if op == nil {
		op = &xxx_GetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Name = o.Name
	op.Return = o.Return
	return op
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
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
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

func (o *xxx_SetNameOperation) OpNum() int { return 15 }

func (o *xxx_SetNameOperation) OpName() string { return "/IFsrmPipelineModuleDefinition/v0/Name" }

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

func (o *SetNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNameOperation) *xxx_SetNameOperation {
	if op == nil {
		op = &xxx_SetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Name = o.Name
	return op
}

func (o *SetNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *SetNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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

func (o *SetNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNameOperation) *xxx_SetNameOperation {
	if op == nil {
		op = &xxx_SetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCompanyOperation structure represents the Company operation
type xxx_GetCompanyOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Company *oaut.String   `idl:"name:company" json:"company"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCompanyOperation) OpNum() int { return 16 }

func (o *xxx_GetCompanyOperation) OpName() string { return "/IFsrmPipelineModuleDefinition/v0/Company" }

func (o *xxx_GetCompanyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompanyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCompanyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCompanyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCompanyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// company {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Company != nil {
			_ptr_company := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Company != nil {
					if err := o.Company.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Company, _ptr_company); err != nil {
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

func (o *xxx_GetCompanyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// company {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_company := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Company == nil {
				o.Company = &oaut.String{}
			}
			if err := o.Company.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_company := func(ptr interface{}) { o.Company = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Company, _s_company, _ptr_company); err != nil {
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

// GetCompanyRequest structure represents the Company operation request
type GetCompanyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCompanyRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCompanyOperation) *xxx_GetCompanyOperation {
	if op == nil {
		op = &xxx_GetCompanyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCompanyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCompanyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCompanyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCompanyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCompanyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCompanyResponse structure represents the Company operation response
type GetCompanyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Company *oaut.String   `idl:"name:company" json:"company"`
	// Return: The Company return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCompanyResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCompanyOperation) *xxx_GetCompanyOperation {
	if op == nil {
		op = &xxx_GetCompanyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Company = o.Company
	op.Return = o.Return
	return op
}

func (o *GetCompanyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCompanyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Company = op.Company
	o.Return = op.Return
}
func (o *GetCompanyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCompanyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCompanyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCompanyOperation structure represents the Company operation
type xxx_SetCompanyOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Company *oaut.String   `idl:"name:company" json:"company"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCompanyOperation) OpNum() int { return 17 }

func (o *xxx_SetCompanyOperation) OpName() string { return "/IFsrmPipelineModuleDefinition/v0/Company" }

func (o *xxx_SetCompanyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCompanyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// company {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Company != nil {
			_ptr_company := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Company != nil {
					if err := o.Company.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Company, _ptr_company); err != nil {
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

func (o *xxx_SetCompanyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// company {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_company := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Company == nil {
				o.Company = &oaut.String{}
			}
			if err := o.Company.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_company := func(ptr interface{}) { o.Company = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Company, _s_company, _ptr_company); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCompanyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCompanyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCompanyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCompanyRequest structure represents the Company operation request
type SetCompanyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Company *oaut.String   `idl:"name:company" json:"company"`
}

func (o *SetCompanyRequest) xxx_ToOp(ctx context.Context, op *xxx_SetCompanyOperation) *xxx_SetCompanyOperation {
	if op == nil {
		op = &xxx_SetCompanyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Company = o.Company
	return op
}

func (o *SetCompanyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCompanyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Company = op.Company
}
func (o *SetCompanyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetCompanyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCompanyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCompanyResponse structure represents the Company operation response
type SetCompanyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Company return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCompanyResponse) xxx_ToOp(ctx context.Context, op *xxx_SetCompanyOperation) *xxx_SetCompanyOperation {
	if op == nil {
		op = &xxx_SetCompanyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetCompanyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCompanyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCompanyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetCompanyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCompanyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVersionOperation structure represents the Version operation
type xxx_GetVersionOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Version *oaut.String   `idl:"name:version" json:"version"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVersionOperation) OpNum() int { return 18 }

func (o *xxx_GetVersionOperation) OpName() string { return "/IFsrmPipelineModuleDefinition/v0/Version" }

func (o *xxx_GetVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// version {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Version != nil {
			_ptr_version := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Version != nil {
					if err := o.Version.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Version, _ptr_version); err != nil {
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

func (o *xxx_GetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// version {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_version := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Version == nil {
				o.Version = &oaut.String{}
			}
			if err := o.Version.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_version := func(ptr interface{}) { o.Version = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Version, _s_version, _ptr_version); err != nil {
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

// GetVersionRequest structure represents the Version operation request
type GetVersionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVersionOperation) *xxx_GetVersionOperation {
	if op == nil {
		op = &xxx_GetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVersionResponse structure represents the Version operation response
type GetVersionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Version *oaut.String   `idl:"name:version" json:"version"`
	// Return: The Version return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVersionOperation) *xxx_GetVersionOperation {
	if op == nil {
		op = &xxx_GetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Version = o.Version
	op.Return = o.Return
	return op
}

func (o *GetVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Version = op.Version
	o.Return = op.Return
}
func (o *GetVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetVersionOperation structure represents the Version operation
type xxx_SetVersionOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Version *oaut.String   `idl:"name:version" json:"version"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetVersionOperation) OpNum() int { return 19 }

func (o *xxx_SetVersionOperation) OpName() string { return "/IFsrmPipelineModuleDefinition/v0/Version" }

func (o *xxx_SetVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// version {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Version != nil {
			_ptr_version := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Version != nil {
					if err := o.Version.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Version, _ptr_version); err != nil {
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

func (o *xxx_SetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// version {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_version := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Version == nil {
				o.Version = &oaut.String{}
			}
			if err := o.Version.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_version := func(ptr interface{}) { o.Version = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Version, _s_version, _ptr_version); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetVersionRequest structure represents the Version operation request
type SetVersionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Version *oaut.String   `idl:"name:version" json:"version"`
}

func (o *SetVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_SetVersionOperation) *xxx_SetVersionOperation {
	if op == nil {
		op = &xxx_SetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Version = o.Version
	return op
}

func (o *SetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetVersionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Version = op.Version
}
func (o *SetVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetVersionResponse structure represents the Version operation response
type SetVersionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Version return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_SetVersionOperation) *xxx_SetVersionOperation {
	if op == nil {
		op = &xxx_SetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetVersionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetModuleTypeOperation structure represents the ModuleType operation
type xxx_GetModuleTypeOperation struct {
	This       *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat          `idl:"name:That" json:"that"`
	ModuleType fsrm.PipelineModuleType `idl:"name:moduleType" json:"module_type"`
	Return     int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetModuleTypeOperation) OpNum() int { return 20 }

func (o *xxx_GetModuleTypeOperation) OpName() string {
	return "/IFsrmPipelineModuleDefinition/v0/ModuleType"
}

func (o *xxx_GetModuleTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetModuleTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetModuleTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// moduleType {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmPipelineModuleType}(enum))
	{
		if err := w.WriteEnum(uint16(o.ModuleType)); err != nil {
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

func (o *xxx_GetModuleTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// moduleType {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmPipelineModuleType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ModuleType)); err != nil {
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

// GetModuleTypeRequest structure represents the ModuleType operation request
type GetModuleTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetModuleTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetModuleTypeOperation) *xxx_GetModuleTypeOperation {
	if op == nil {
		op = &xxx_GetModuleTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetModuleTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetModuleTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetModuleTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetModuleTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetModuleTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetModuleTypeResponse structure represents the ModuleType operation response
type GetModuleTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// moduleType: Pointer to a variable that upon completion contains the Module type of
	// the module definition.
	ModuleType fsrm.PipelineModuleType `idl:"name:moduleType" json:"module_type"`
	// Return: The ModuleType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetModuleTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetModuleTypeOperation) *xxx_GetModuleTypeOperation {
	if op == nil {
		op = &xxx_GetModuleTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ModuleType = o.ModuleType
	op.Return = o.Return
	return op
}

func (o *GetModuleTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetModuleTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ModuleType = op.ModuleType
	o.Return = op.Return
}
func (o *GetModuleTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetModuleTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetModuleTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEnabledOperation structure represents the Enabled operation
type xxx_GetEnabledOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEnabledOperation) OpNum() int { return 21 }

func (o *xxx_GetEnabledOperation) OpName() string { return "/IFsrmPipelineModuleDefinition/v0/Enabled" }

func (o *xxx_GetEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// enabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Enabled); err != nil {
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

func (o *xxx_GetEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// enabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Enabled); err != nil {
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

// GetEnabledRequest structure represents the Enabled operation request
type GetEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEnabledOperation) *xxx_GetEnabledOperation {
	if op == nil {
		op = &xxx_GetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEnabledResponse structure represents the Enabled operation response
type GetEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
	// Return: The Enabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEnabledOperation) *xxx_GetEnabledOperation {
	if op == nil {
		op = &xxx_GetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enabled = o.Enabled
	op.Return = o.Return
	return op
}

func (o *GetEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enabled = op.Enabled
	o.Return = op.Return
}
func (o *GetEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEnabledOperation structure represents the Enabled operation
type xxx_SetEnabledOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEnabledOperation) OpNum() int { return 22 }

func (o *xxx_SetEnabledOperation) OpName() string { return "/IFsrmPipelineModuleDefinition/v0/Enabled" }

func (o *xxx_SetEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// enabled {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Enabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// enabled {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Enabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEnabledRequest structure represents the Enabled operation request
type SetEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Enabled int16          `idl:"name:enabled" json:"enabled"`
}

func (o *SetEnabledRequest) xxx_ToOp(ctx context.Context, op *xxx_SetEnabledOperation) *xxx_SetEnabledOperation {
	if op == nil {
		op = &xxx_SetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Enabled = o.Enabled
	return op
}

func (o *SetEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Enabled = op.Enabled
}
func (o *SetEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEnabledResponse structure represents the Enabled operation response
type SetEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Enabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEnabledResponse) xxx_ToOp(ctx context.Context, op *xxx_SetEnabledOperation) *xxx_SetEnabledOperation {
	if op == nil {
		op = &xxx_SetEnabledOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNeedsFileContentOperation structure represents the NeedsFileContent operation
type xxx_GetNeedsFileContentOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	NeedsFileContent int16          `idl:"name:needsFileContent" json:"needs_file_content"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNeedsFileContentOperation) OpNum() int { return 23 }

func (o *xxx_GetNeedsFileContentOperation) OpName() string {
	return "/IFsrmPipelineModuleDefinition/v0/NeedsFileContent"
}

func (o *xxx_GetNeedsFileContentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNeedsFileContentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNeedsFileContentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNeedsFileContentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNeedsFileContentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// needsFileContent {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.NeedsFileContent); err != nil {
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

func (o *xxx_GetNeedsFileContentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// needsFileContent {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.NeedsFileContent); err != nil {
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

// GetNeedsFileContentRequest structure represents the NeedsFileContent operation request
type GetNeedsFileContentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNeedsFileContentRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNeedsFileContentOperation) *xxx_GetNeedsFileContentOperation {
	if op == nil {
		op = &xxx_GetNeedsFileContentOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNeedsFileContentRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNeedsFileContentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNeedsFileContentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNeedsFileContentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNeedsFileContentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNeedsFileContentResponse structure represents the NeedsFileContent operation response
type GetNeedsFileContentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	NeedsFileContent int16          `idl:"name:needsFileContent" json:"needs_file_content"`
	// Return: The NeedsFileContent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNeedsFileContentResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNeedsFileContentOperation) *xxx_GetNeedsFileContentOperation {
	if op == nil {
		op = &xxx_GetNeedsFileContentOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.NeedsFileContent = o.NeedsFileContent
	op.Return = o.Return
	return op
}

func (o *GetNeedsFileContentResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNeedsFileContentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.NeedsFileContent = op.NeedsFileContent
	o.Return = op.Return
}
func (o *GetNeedsFileContentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNeedsFileContentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNeedsFileContentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNeedsFileContentOperation structure represents the NeedsFileContent operation
type xxx_SetNeedsFileContentOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	NeedsFileContent int16          `idl:"name:needsFileContent" json:"needs_file_content"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNeedsFileContentOperation) OpNum() int { return 24 }

func (o *xxx_SetNeedsFileContentOperation) OpName() string {
	return "/IFsrmPipelineModuleDefinition/v0/NeedsFileContent"
}

func (o *xxx_SetNeedsFileContentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNeedsFileContentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// needsFileContent {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.NeedsFileContent); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNeedsFileContentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// needsFileContent {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.NeedsFileContent); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNeedsFileContentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNeedsFileContentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNeedsFileContentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNeedsFileContentRequest structure represents the NeedsFileContent operation request
type SetNeedsFileContentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	NeedsFileContent int16          `idl:"name:needsFileContent" json:"needs_file_content"`
}

func (o *SetNeedsFileContentRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNeedsFileContentOperation) *xxx_SetNeedsFileContentOperation {
	if op == nil {
		op = &xxx_SetNeedsFileContentOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.NeedsFileContent = o.NeedsFileContent
	return op
}

func (o *SetNeedsFileContentRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNeedsFileContentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NeedsFileContent = op.NeedsFileContent
}
func (o *SetNeedsFileContentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNeedsFileContentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNeedsFileContentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNeedsFileContentResponse structure represents the NeedsFileContent operation response
type SetNeedsFileContentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The NeedsFileContent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNeedsFileContentResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNeedsFileContentOperation) *xxx_SetNeedsFileContentOperation {
	if op == nil {
		op = &xxx_SetNeedsFileContentOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNeedsFileContentResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNeedsFileContentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNeedsFileContentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNeedsFileContentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNeedsFileContentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAccountOperation structure represents the Account operation
type xxx_GetAccountOperation struct {
	This             *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat   `idl:"name:That" json:"that"`
	RetrievalAccount fsrm.AccountType `idl:"name:retrievalAccount" json:"retrieval_account"`
	Return           int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAccountOperation) OpNum() int { return 25 }

func (o *xxx_GetAccountOperation) OpName() string { return "/IFsrmPipelineModuleDefinition/v0/Account" }

func (o *xxx_GetAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retrievalAccount {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmAccountType}(enum))
	{
		if err := w.WriteEnum(uint16(o.RetrievalAccount)); err != nil {
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

func (o *xxx_GetAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retrievalAccount {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmAccountType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.RetrievalAccount)); err != nil {
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

// GetAccountRequest structure represents the Account operation request
type GetAccountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAccountRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAccountOperation) *xxx_GetAccountOperation {
	if op == nil {
		op = &xxx_GetAccountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAccountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAccountResponse structure represents the Account operation response
type GetAccountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat   `idl:"name:That" json:"that"`
	RetrievalAccount fsrm.AccountType `idl:"name:retrievalAccount" json:"retrieval_account"`
	// Return: The Account return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAccountResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAccountOperation) *xxx_GetAccountOperation {
	if op == nil {
		op = &xxx_GetAccountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.RetrievalAccount = o.RetrievalAccount
	op.Return = o.Return
	return op
}

func (o *GetAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAccountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RetrievalAccount = op.RetrievalAccount
	o.Return = op.Return
}
func (o *GetAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAccountOperation structure represents the Account operation
type xxx_SetAccountOperation struct {
	This             *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat   `idl:"name:That" json:"that"`
	RetrievalAccount fsrm.AccountType `idl:"name:retrievalAccount" json:"retrieval_account"`
	Return           int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAccountOperation) OpNum() int { return 26 }

func (o *xxx_SetAccountOperation) OpName() string { return "/IFsrmPipelineModuleDefinition/v0/Account" }

func (o *xxx_SetAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// retrievalAccount {in} (1:{alias=FsrmAccountType}(enum))
	{
		if err := w.WriteEnum(uint16(o.RetrievalAccount)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// retrievalAccount {in} (1:{alias=FsrmAccountType}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.RetrievalAccount)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAccountRequest structure represents the Account operation request
type SetAccountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This             *dcom.ORPCThis   `idl:"name:This" json:"this"`
	RetrievalAccount fsrm.AccountType `idl:"name:retrievalAccount" json:"retrieval_account"`
}

func (o *SetAccountRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAccountOperation) *xxx_SetAccountOperation {
	if op == nil {
		op = &xxx_SetAccountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RetrievalAccount = o.RetrievalAccount
	return op
}

func (o *SetAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAccountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RetrievalAccount = op.RetrievalAccount
}
func (o *SetAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAccountResponse structure represents the Account operation response
type SetAccountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Account return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAccountResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAccountOperation) *xxx_SetAccountOperation {
	if op == nil {
		op = &xxx_SetAccountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAccountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSupportedExtensionsOperation structure represents the SupportedExtensions operation
type xxx_GetSupportedExtensionsOperation struct {
	This                *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat  `idl:"name:That" json:"that"`
	SupportedExtensions *oaut.SafeArray `idl:"name:supportedExtensions" json:"supported_extensions"`
	Return              int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSupportedExtensionsOperation) OpNum() int { return 27 }

func (o *xxx_GetSupportedExtensionsOperation) OpName() string {
	return "/IFsrmPipelineModuleDefinition/v0/SupportedExtensions"
}

func (o *xxx_GetSupportedExtensionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedExtensionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSupportedExtensionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSupportedExtensionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedExtensionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// supportedExtensions {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.SupportedExtensions != nil {
			_ptr_supportedExtensions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SupportedExtensions != nil {
					if err := o.SupportedExtensions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SupportedExtensions, _ptr_supportedExtensions); err != nil {
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

func (o *xxx_GetSupportedExtensionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// supportedExtensions {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_supportedExtensions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SupportedExtensions == nil {
				o.SupportedExtensions = &oaut.SafeArray{}
			}
			if err := o.SupportedExtensions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_supportedExtensions := func(ptr interface{}) { o.SupportedExtensions = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.SupportedExtensions, _s_supportedExtensions, _ptr_supportedExtensions); err != nil {
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

// GetSupportedExtensionsRequest structure represents the SupportedExtensions operation request
type GetSupportedExtensionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSupportedExtensionsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSupportedExtensionsOperation) *xxx_GetSupportedExtensionsOperation {
	if op == nil {
		op = &xxx_GetSupportedExtensionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetSupportedExtensionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSupportedExtensionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSupportedExtensionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSupportedExtensionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSupportedExtensionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSupportedExtensionsResponse structure represents the SupportedExtensions operation response
type GetSupportedExtensionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                *dcom.ORPCThat  `idl:"name:That" json:"that"`
	SupportedExtensions *oaut.SafeArray `idl:"name:supportedExtensions" json:"supported_extensions"`
	// Return: The SupportedExtensions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSupportedExtensionsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSupportedExtensionsOperation) *xxx_GetSupportedExtensionsOperation {
	if op == nil {
		op = &xxx_GetSupportedExtensionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SupportedExtensions = o.SupportedExtensions
	op.Return = o.Return
	return op
}

func (o *GetSupportedExtensionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSupportedExtensionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SupportedExtensions = op.SupportedExtensions
	o.Return = op.Return
}
func (o *GetSupportedExtensionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSupportedExtensionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSupportedExtensionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSupportedExtensionsOperation structure represents the SupportedExtensions operation
type xxx_SetSupportedExtensionsOperation struct {
	This                *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat  `idl:"name:That" json:"that"`
	SupportedExtensions *oaut.SafeArray `idl:"name:supportedExtensions" json:"supported_extensions"`
	Return              int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSupportedExtensionsOperation) OpNum() int { return 28 }

func (o *xxx_SetSupportedExtensionsOperation) OpName() string {
	return "/IFsrmPipelineModuleDefinition/v0/SupportedExtensions"
}

func (o *xxx_SetSupportedExtensionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSupportedExtensionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// supportedExtensions {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.SupportedExtensions != nil {
			_ptr_supportedExtensions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SupportedExtensions != nil {
					if err := o.SupportedExtensions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SupportedExtensions, _ptr_supportedExtensions); err != nil {
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

func (o *xxx_SetSupportedExtensionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// supportedExtensions {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_supportedExtensions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SupportedExtensions == nil {
				o.SupportedExtensions = &oaut.SafeArray{}
			}
			if err := o.SupportedExtensions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_supportedExtensions := func(ptr interface{}) { o.SupportedExtensions = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.SupportedExtensions, _s_supportedExtensions, _ptr_supportedExtensions); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSupportedExtensionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSupportedExtensionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSupportedExtensionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSupportedExtensionsRequest structure represents the SupportedExtensions operation request
type SetSupportedExtensionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                *dcom.ORPCThis  `idl:"name:This" json:"this"`
	SupportedExtensions *oaut.SafeArray `idl:"name:supportedExtensions" json:"supported_extensions"`
}

func (o *SetSupportedExtensionsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSupportedExtensionsOperation) *xxx_SetSupportedExtensionsOperation {
	if op == nil {
		op = &xxx_SetSupportedExtensionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SupportedExtensions = o.SupportedExtensions
	return op
}

func (o *SetSupportedExtensionsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSupportedExtensionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SupportedExtensions = op.SupportedExtensions
}
func (o *SetSupportedExtensionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSupportedExtensionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSupportedExtensionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSupportedExtensionsResponse structure represents the SupportedExtensions operation response
type SetSupportedExtensionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SupportedExtensions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSupportedExtensionsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSupportedExtensionsOperation) *xxx_SetSupportedExtensionsOperation {
	if op == nil {
		op = &xxx_SetSupportedExtensionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetSupportedExtensionsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSupportedExtensionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSupportedExtensionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSupportedExtensionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSupportedExtensionsOperation{}
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

func (o *xxx_GetParametersOperation) OpNum() int { return 29 }

func (o *xxx_GetParametersOperation) OpName() string {
	return "/IFsrmPipelineModuleDefinition/v0/Parameters"
}

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

func (o *GetParametersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetParametersOperation) *xxx_GetParametersOperation {
	if op == nil {
		op = &xxx_GetParametersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetParametersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetParametersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetParametersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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

func (o *GetParametersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetParametersOperation) *xxx_GetParametersOperation {
	if op == nil {
		op = &xxx_GetParametersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Parameters = o.Parameters
	op.Return = o.Return
	return op
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
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
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

func (o *xxx_SetParametersOperation) OpNum() int { return 30 }

func (o *xxx_SetParametersOperation) OpName() string {
	return "/IFsrmPipelineModuleDefinition/v0/Parameters"
}

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

func (o *SetParametersRequest) xxx_ToOp(ctx context.Context, op *xxx_SetParametersOperation) *xxx_SetParametersOperation {
	if op == nil {
		op = &xxx_SetParametersOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Parameters = o.Parameters
	return op
}

func (o *SetParametersRequest) xxx_FromOp(ctx context.Context, op *xxx_SetParametersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Parameters = op.Parameters
}
func (o *SetParametersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
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

func (o *SetParametersResponse) xxx_ToOp(ctx context.Context, op *xxx_SetParametersOperation) *xxx_SetParametersOperation {
	if op == nil {
		op = &xxx_SetParametersOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetParametersResponse) xxx_FromOp(ctx context.Context, op *xxx_SetParametersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetParametersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetParametersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetParametersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
