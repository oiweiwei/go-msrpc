package ifsrmclassifiermoduledefinition

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmpipelinemoduledefinition "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpipelinemoduledefinition/v0"
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
	_ = ifsrmpipelinemoduledefinition.GoPackage
)

// IFsrmClassifierModuleDefinition server interface.
type ClassifierModuleDefinitionServer interface {

	// IFsrmPipelineModuleDefinition base class.
	ifsrmpipelinemoduledefinition.PipelineModuleDefinitionServer

	// PropertiesAffected operation.
	GetPropertiesAffected(context.Context, *GetPropertiesAffectedRequest) (*GetPropertiesAffectedResponse, error)

	// PropertiesAffected operation.
	SetPropertiesAffected(context.Context, *SetPropertiesAffectedRequest) (*SetPropertiesAffectedResponse, error)

	// PropertiesUsed operation.
	GetPropertiesUsed(context.Context, *GetPropertiesUsedRequest) (*GetPropertiesUsedResponse, error)

	// PropertiesUsed operation.
	SetPropertiesUsed(context.Context, *SetPropertiesUsedRequest) (*SetPropertiesUsedResponse, error)

	// NeedsExplicitValue operation.
	GetNeedsExplicitValue(context.Context, *GetNeedsExplicitValueRequest) (*GetNeedsExplicitValueResponse, error)

	// NeedsExplicitValue operation.
	SetNeedsExplicitValue(context.Context, *SetNeedsExplicitValueRequest) (*SetNeedsExplicitValueResponse, error)
}

func RegisterClassifierModuleDefinitionServer(conn dcerpc.Conn, o ClassifierModuleDefinitionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClassifierModuleDefinitionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClassifierModuleDefinitionSyntaxV0_0))...)
}

func NewClassifierModuleDefinitionServerHandle(o ClassifierModuleDefinitionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClassifierModuleDefinitionServerHandle(ctx, o, opNum, r)
	}
}

func ClassifierModuleDefinitionServerHandle(ctx context.Context, o ClassifierModuleDefinitionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 31 {
		// IFsrmPipelineModuleDefinition base method.
		return ifsrmpipelinemoduledefinition.PipelineModuleDefinitionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 31: // PropertiesAffected
		in := &GetPropertiesAffectedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertiesAffected(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // PropertiesAffected
		in := &SetPropertiesAffectedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPropertiesAffected(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // PropertiesUsed
		in := &GetPropertiesUsedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertiesUsed(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // PropertiesUsed
		in := &SetPropertiesUsedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPropertiesUsed(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // NeedsExplicitValue
		in := &GetNeedsExplicitValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNeedsExplicitValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // NeedsExplicitValue
		in := &SetNeedsExplicitValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetNeedsExplicitValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
