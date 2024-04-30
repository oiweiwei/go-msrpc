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
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Name
		in := &SetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // RuleType
		in := &GetRuleTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRuleType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // ModuleDefinitionName
		in := &GetModuleDefinitionNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetModuleDefinitionName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // ModuleDefinitionName
		in := &SetModuleDefinitionNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetModuleDefinitionName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // NamespaceRoots
		in := &GetNamespaceRootsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNamespaceRoots(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // NamespaceRoots
		in := &SetNamespaceRootsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetNamespaceRoots(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // RuleFlags
		in := &GetRuleFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRuleFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // RuleFlags
		in := &SetRuleFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetRuleFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // Parameters
		in := &GetParametersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetParameters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // Parameters
		in := &SetParametersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetParameters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // LastModified
		in := &GetLastModifiedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLastModified(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
