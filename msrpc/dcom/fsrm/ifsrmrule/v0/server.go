package ifsrmrule

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

// IFsrmRule server interface.
type RuleServer interface {

	// IFsrmObject base class.
	ifsrmobject.ObjectServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest) (*SetNameResponse, error)

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
	GetRuleType(context.Context, *GetRuleTypeRequest) (*GetRuleTypeResponse, error)

	// ModuleDefinitionName operation.
	GetModuleDefinitionName(context.Context, *GetModuleDefinitionNameRequest) (*GetModuleDefinitionNameResponse, error)

	// ModuleDefinitionName operation.
	SetModuleDefinitionName(context.Context, *SetModuleDefinitionNameRequest) (*SetModuleDefinitionNameResponse, error)

	// NamespaceRoots operation.
	GetNamespaceRoots(context.Context, *GetNamespaceRootsRequest) (*GetNamespaceRootsResponse, error)

	// NamespaceRoots operation.
	SetNamespaceRoots(context.Context, *SetNamespaceRootsRequest) (*SetNamespaceRootsResponse, error)

	// RuleFlags operation.
	GetRuleFlags(context.Context, *GetRuleFlagsRequest) (*GetRuleFlagsResponse, error)

	// RuleFlags operation.
	SetRuleFlags(context.Context, *SetRuleFlagsRequest) (*SetRuleFlagsResponse, error)

	// Parameters operation.
	GetParameters(context.Context, *GetParametersRequest) (*GetParametersResponse, error)

	// Parameters operation.
	SetParameters(context.Context, *SetParametersRequest) (*SetParametersResponse, error)

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
	GetLastModified(context.Context, *GetLastModifiedRequest) (*GetLastModifiedResponse, error)
}

func RegisterRuleServer(conn dcerpc.Conn, o RuleServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRuleServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RuleSyntaxV0_0))...)
}

func NewRuleServerHandle(o RuleServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RuleServerHandle(ctx, o, opNum, r)
	}
}

func RuleServerHandle(ctx context.Context, o RuleServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmObject base method.
		return ifsrmobject.ObjectServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Name
		op := &xxx_SetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // RuleType
		op := &xxx_GetRuleTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRuleTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRuleType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // ModuleDefinitionName
		op := &xxx_GetModuleDefinitionNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetModuleDefinitionNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetModuleDefinitionName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // ModuleDefinitionName
		op := &xxx_SetModuleDefinitionNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetModuleDefinitionNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetModuleDefinitionName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // NamespaceRoots
		op := &xxx_GetNamespaceRootsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNamespaceRootsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNamespaceRoots(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // NamespaceRoots
		op := &xxx_SetNamespaceRootsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNamespaceRootsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNamespaceRoots(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // RuleFlags
		op := &xxx_GetRuleFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRuleFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRuleFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // RuleFlags
		op := &xxx_SetRuleFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRuleFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRuleFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Parameters
		op := &xxx_GetParametersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetParametersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetParameters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // Parameters
		op := &xxx_SetParametersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetParametersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetParameters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // LastModified
		op := &xxx_GetLastModifiedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastModifiedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastModified(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmRule
type UnimplementedRuleServer struct {
	ifsrmobject.UnimplementedObjectServer
}

func (UnimplementedRuleServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) SetName(context.Context, *SetNameRequest) (*SetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) GetRuleType(context.Context, *GetRuleTypeRequest) (*GetRuleTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) GetModuleDefinitionName(context.Context, *GetModuleDefinitionNameRequest) (*GetModuleDefinitionNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) SetModuleDefinitionName(context.Context, *SetModuleDefinitionNameRequest) (*SetModuleDefinitionNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) GetNamespaceRoots(context.Context, *GetNamespaceRootsRequest) (*GetNamespaceRootsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) SetNamespaceRoots(context.Context, *SetNamespaceRootsRequest) (*SetNamespaceRootsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) GetRuleFlags(context.Context, *GetRuleFlagsRequest) (*GetRuleFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) SetRuleFlags(context.Context, *SetRuleFlagsRequest) (*SetRuleFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) GetParameters(context.Context, *GetParametersRequest) (*GetParametersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) SetParameters(context.Context, *SetParametersRequest) (*SetParametersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRuleServer) GetLastModified(context.Context, *GetLastModifiedRequest) (*GetLastModifiedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RuleServer = (*UnimplementedRuleServer)(nil)
