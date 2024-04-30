package ifsrmpipelinemoduledefinition

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmobject/v0"
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
	_ = ifsrmobject.GoPackage
)

// IFsrmPipelineModuleDefinition server interface.
type PipelineModuleDefinitionServer interface {

	// IFsrmObject base class.
	ifsrmobject.ObjectServer

	// ModuleClsid operation.
	GetModuleClassID(context.Context, *GetModuleClassIDRequest) (*GetModuleClassIDResponse, error)

	// ModuleClsid operation.
	SetModuleClassID(context.Context, *SetModuleClassIDRequest) (*SetModuleClassIDResponse, error)

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest) (*SetNameResponse, error)

	// Company operation.
	GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error)

	// Company operation.
	SetCompany(context.Context, *SetCompanyRequest) (*SetCompanyResponse, error)

	// Version operation.
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)

	// Version operation.
	SetVersion(context.Context, *SetVersionRequest) (*SetVersionResponse, error)

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
	GetModuleType(context.Context, *GetModuleTypeRequest) (*GetModuleTypeResponse, error)

	// Enabled operation.
	GetEnabled(context.Context, *GetEnabledRequest) (*GetEnabledResponse, error)

	// Enabled operation.
	SetEnabled(context.Context, *SetEnabledRequest) (*SetEnabledResponse, error)

	// NeedsFileContent operation.
	GetNeedsFileContent(context.Context, *GetNeedsFileContentRequest) (*GetNeedsFileContentResponse, error)

	// NeedsFileContent operation.
	SetNeedsFileContent(context.Context, *SetNeedsFileContentRequest) (*SetNeedsFileContentResponse, error)

	// Account operation.
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)

	// Account operation.
	SetAccount(context.Context, *SetAccountRequest) (*SetAccountResponse, error)

	// SupportedExtensions operation.
	GetSupportedExtensions(context.Context, *GetSupportedExtensionsRequest) (*GetSupportedExtensionsResponse, error)

	// SupportedExtensions operation.
	SetSupportedExtensions(context.Context, *SetSupportedExtensionsRequest) (*SetSupportedExtensionsResponse, error)

	// Parameters operation.
	GetParameters(context.Context, *GetParametersRequest) (*GetParametersResponse, error)

	// Parameters operation.
	SetParameters(context.Context, *SetParametersRequest) (*SetParametersResponse, error)
}

func RegisterPipelineModuleDefinitionServer(conn dcerpc.Conn, o PipelineModuleDefinitionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPipelineModuleDefinitionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PipelineModuleDefinitionSyntaxV0_0))...)
}

func NewPipelineModuleDefinitionServerHandle(o PipelineModuleDefinitionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PipelineModuleDefinitionServerHandle(ctx, o, opNum, r)
	}
}

func PipelineModuleDefinitionServerHandle(ctx context.Context, o PipelineModuleDefinitionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmObject base method.
		return ifsrmobject.ObjectServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // ModuleClsid
		in := &GetModuleClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetModuleClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // ModuleClsid
		in := &SetModuleClassIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetModuleClassID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // Name
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Name
		in := &SetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // Company
		in := &GetCompanyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCompany(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // Company
		in := &SetCompanyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetCompany(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // Version
		in := &GetVersionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetVersion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // Version
		in := &SetVersionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetVersion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // ModuleType
		in := &GetModuleTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetModuleType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // Enabled
		in := &GetEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // Enabled
		in := &SetEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // NeedsFileContent
		in := &GetNeedsFileContentRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNeedsFileContent(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // NeedsFileContent
		in := &SetNeedsFileContentRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetNeedsFileContent(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // Account
		in := &GetAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // Account
		in := &SetAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // SupportedExtensions
		in := &GetSupportedExtensionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSupportedExtensions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // SupportedExtensions
		in := &SetSupportedExtensionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSupportedExtensions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // Parameters
		in := &GetParametersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetParameters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // Parameters
		in := &SetParametersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetParameters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
