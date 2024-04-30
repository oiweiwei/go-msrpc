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
		in := &GetCapabilitiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCapabilities(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // Capabilities
		in := &SetCapabilitiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetCapabilities(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // StorageType
		in := &GetStorageTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetStorageType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // StorageType
		in := &SetStorageTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetStorageType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // UpdatesFileContent
		in := &GetUpdatesFileContentRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUpdatesFileContent(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // UpdatesFileContent
		in := &SetUpdatesFileContentRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetUpdatesFileContent(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
