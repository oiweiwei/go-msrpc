package ifsrmpipelinemoduledefinition

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_GetModuleClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetModuleClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetModuleClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // ModuleClsid
		op := &xxx_SetModuleClassIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetModuleClassIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetModuleClassID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Name
		op := &xxx_SetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Company
		op := &xxx_GetCompanyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCompanyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCompany(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Company
		op := &xxx_SetCompanyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCompanyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCompany(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // Version
		op := &xxx_GetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Version
		op := &xxx_SetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // ModuleType
		op := &xxx_GetModuleTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetModuleTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetModuleType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Enabled
		op := &xxx_GetEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // Enabled
		op := &xxx_SetEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // NeedsFileContent
		op := &xxx_GetNeedsFileContentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNeedsFileContentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNeedsFileContent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // NeedsFileContent
		op := &xxx_SetNeedsFileContentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNeedsFileContentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNeedsFileContent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // Account
		op := &xxx_GetAccountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAccountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAccount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // Account
		op := &xxx_SetAccountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAccountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAccount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // SupportedExtensions
		op := &xxx_GetSupportedExtensionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSupportedExtensionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSupportedExtensions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // SupportedExtensions
		op := &xxx_SetSupportedExtensionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSupportedExtensionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSupportedExtensions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // Parameters
		op := &xxx_GetParametersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetParametersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetParameters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // Parameters
		op := &xxx_SetParametersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetParametersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetParameters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmPipelineModuleDefinition
type UnimplementedPipelineModuleDefinitionServer struct {
	ifsrmobject.UnimplementedObjectServer
}

func (UnimplementedPipelineModuleDefinitionServer) GetModuleClassID(context.Context, *GetModuleClassIDRequest) (*GetModuleClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) SetModuleClassID(context.Context, *SetModuleClassIDRequest) (*SetModuleClassIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) SetName(context.Context, *SetNameRequest) (*SetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) SetCompany(context.Context, *SetCompanyRequest) (*SetCompanyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) SetVersion(context.Context, *SetVersionRequest) (*SetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) GetModuleType(context.Context, *GetModuleTypeRequest) (*GetModuleTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) GetEnabled(context.Context, *GetEnabledRequest) (*GetEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) SetEnabled(context.Context, *SetEnabledRequest) (*SetEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) GetNeedsFileContent(context.Context, *GetNeedsFileContentRequest) (*GetNeedsFileContentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) SetNeedsFileContent(context.Context, *SetNeedsFileContentRequest) (*SetNeedsFileContentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) SetAccount(context.Context, *SetAccountRequest) (*SetAccountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) GetSupportedExtensions(context.Context, *GetSupportedExtensionsRequest) (*GetSupportedExtensionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) SetSupportedExtensions(context.Context, *SetSupportedExtensionsRequest) (*SetSupportedExtensionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) GetParameters(context.Context, *GetParametersRequest) (*GetParametersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPipelineModuleDefinitionServer) SetParameters(context.Context, *SetParametersRequest) (*SetParametersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PipelineModuleDefinitionServer = (*UnimplementedPipelineModuleDefinitionServer)(nil)
