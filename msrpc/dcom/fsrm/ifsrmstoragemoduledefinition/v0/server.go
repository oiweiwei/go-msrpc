package ifsrmstoragemoduledefinition

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

// IFsrmStorageModuleDefinition server interface.
type StorageModuleDefinitionServer interface {

	// IFsrmPipelineModuleDefinition base class.
	ifsrmpipelinemoduledefinition.PipelineModuleDefinitionServer

	// Capabilities operation.
	GetCapabilities(context.Context, *GetCapabilitiesRequest) (*GetCapabilitiesResponse, error)

	// Capabilities operation.
	SetCapabilities(context.Context, *SetCapabilitiesRequest) (*SetCapabilitiesResponse, error)

	// StorageType operation.
	GetStorageType(context.Context, *GetStorageTypeRequest) (*GetStorageTypeResponse, error)

	// StorageType operation.
	SetStorageType(context.Context, *SetStorageTypeRequest) (*SetStorageTypeResponse, error)

	// UpdatesFileContent operation.
	GetUpdatesFileContent(context.Context, *GetUpdatesFileContentRequest) (*GetUpdatesFileContentResponse, error)

	// UpdatesFileContent operation.
	SetUpdatesFileContent(context.Context, *SetUpdatesFileContentRequest) (*SetUpdatesFileContentResponse, error)
}

func RegisterStorageModuleDefinitionServer(conn dcerpc.Conn, o StorageModuleDefinitionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewStorageModuleDefinitionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(StorageModuleDefinitionSyntaxV0_0))...)
}

func NewStorageModuleDefinitionServerHandle(o StorageModuleDefinitionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return StorageModuleDefinitionServerHandle(ctx, o, opNum, r)
	}
}

func StorageModuleDefinitionServerHandle(ctx context.Context, o StorageModuleDefinitionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 31 {
		// IFsrmPipelineModuleDefinition base method.
		return ifsrmpipelinemoduledefinition.PipelineModuleDefinitionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 31: // Capabilities
		op := &xxx_GetCapabilitiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCapabilitiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCapabilities(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // Capabilities
		op := &xxx_SetCapabilitiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCapabilitiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCapabilities(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // StorageType
		op := &xxx_GetStorageTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStorageTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetStorageType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // StorageType
		op := &xxx_SetStorageTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetStorageTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetStorageType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // UpdatesFileContent
		op := &xxx_GetUpdatesFileContentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUpdatesFileContentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUpdatesFileContent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // UpdatesFileContent
		op := &xxx_SetUpdatesFileContentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetUpdatesFileContentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetUpdatesFileContent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmStorageModuleDefinition
type UnimplementedStorageModuleDefinitionServer struct {
	ifsrmpipelinemoduledefinition.UnimplementedPipelineModuleDefinitionServer
}

func (UnimplementedStorageModuleDefinitionServer) GetCapabilities(context.Context, *GetCapabilitiesRequest) (*GetCapabilitiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStorageModuleDefinitionServer) SetCapabilities(context.Context, *SetCapabilitiesRequest) (*SetCapabilitiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStorageModuleDefinitionServer) GetStorageType(context.Context, *GetStorageTypeRequest) (*GetStorageTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStorageModuleDefinitionServer) SetStorageType(context.Context, *SetStorageTypeRequest) (*SetStorageTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStorageModuleDefinitionServer) GetUpdatesFileContent(context.Context, *GetUpdatesFileContentRequest) (*GetUpdatesFileContentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedStorageModuleDefinitionServer) SetUpdatesFileContent(context.Context, *SetUpdatesFileContentRequest) (*SetUpdatesFileContentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ StorageModuleDefinitionServer = (*UnimplementedStorageModuleDefinitionServer)(nil)
