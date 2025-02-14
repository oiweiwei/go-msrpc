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
		op := &xxx_GetPropertiesAffectedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesAffectedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPropertiesAffected(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // PropertiesAffected
		op := &xxx_SetPropertiesAffectedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPropertiesAffectedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPropertiesAffected(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // PropertiesUsed
		op := &xxx_GetPropertiesUsedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesUsedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPropertiesUsed(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // PropertiesUsed
		op := &xxx_SetPropertiesUsedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPropertiesUsedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPropertiesUsed(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // NeedsExplicitValue
		op := &xxx_GetNeedsExplicitValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNeedsExplicitValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNeedsExplicitValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // NeedsExplicitValue
		op := &xxx_SetNeedsExplicitValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNeedsExplicitValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNeedsExplicitValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmClassifierModuleDefinition
type UnimplementedClassifierModuleDefinitionServer struct {
	ifsrmpipelinemoduledefinition.UnimplementedPipelineModuleDefinitionServer
}

func (UnimplementedClassifierModuleDefinitionServer) GetPropertiesAffected(context.Context, *GetPropertiesAffectedRequest) (*GetPropertiesAffectedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClassifierModuleDefinitionServer) SetPropertiesAffected(context.Context, *SetPropertiesAffectedRequest) (*SetPropertiesAffectedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClassifierModuleDefinitionServer) GetPropertiesUsed(context.Context, *GetPropertiesUsedRequest) (*GetPropertiesUsedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClassifierModuleDefinitionServer) SetPropertiesUsed(context.Context, *SetPropertiesUsedRequest) (*SetPropertiesUsedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClassifierModuleDefinitionServer) GetNeedsExplicitValue(context.Context, *GetNeedsExplicitValueRequest) (*GetNeedsExplicitValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClassifierModuleDefinitionServer) SetNeedsExplicitValue(context.Context, *SetNeedsExplicitValueRequest) (*SetNeedsExplicitValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ClassifierModuleDefinitionServer = (*UnimplementedClassifierModuleDefinitionServer)(nil)
