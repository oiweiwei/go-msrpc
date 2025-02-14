package iapphostconfigfile

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IAppHostConfigFile server interface.
type AppHostConfigFileServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// ConfigPath operation.
	GetConfigPath(context.Context, *GetConfigPathRequest) (*GetConfigPathResponse, error)

	// FilePath operation.
	GetFilePath(context.Context, *GetFilePathRequest) (*GetFilePathResponse, error)

	// Locations operation.
	GetLocations(context.Context, *GetLocationsRequest) (*GetLocationsResponse, error)

	// GetAdminSection operation.
	GetAdminSection(context.Context, *GetAdminSectionRequest) (*GetAdminSectionResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// SetMetadata operation.
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)

	// ClearInvalidSections operation.
	ClearInvalidSections(context.Context, *ClearInvalidSectionsRequest) (*ClearInvalidSectionsResponse, error)

	// RootSectionGroup operation.
	GetRootSectionGroup(context.Context, *GetRootSectionGroupRequest) (*GetRootSectionGroupResponse, error)
}

func RegisterAppHostConfigFileServer(conn dcerpc.Conn, o AppHostConfigFileServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostConfigFileServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostConfigFileSyntaxV0_0))...)
}

func NewAppHostConfigFileServerHandle(o AppHostConfigFileServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostConfigFileServerHandle(ctx, o, opNum, r)
	}
}

func AppHostConfigFileServerHandle(ctx context.Context, o AppHostConfigFileServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ConfigPath
		op := &xxx_GetConfigPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfigPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // FilePath
		op := &xxx_GetFilePathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFilePathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFilePath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Locations
		op := &xxx_GetLocationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLocationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLocations(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // GetAdminSection
		op := &xxx_GetAdminSectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAdminSectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAdminSection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // GetMetadata
		op := &xxx_GetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SetMetadata
		op := &xxx_SetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ClearInvalidSections
		op := &xxx_ClearInvalidSectionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearInvalidSectionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClearInvalidSections(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RootSectionGroup
		op := &xxx_GetRootSectionGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRootSectionGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRootSectionGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostConfigFile
type UnimplementedAppHostConfigFileServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostConfigFileServer) GetConfigPath(context.Context, *GetConfigPathRequest) (*GetConfigPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigFileServer) GetFilePath(context.Context, *GetFilePathRequest) (*GetFilePathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigFileServer) GetLocations(context.Context, *GetLocationsRequest) (*GetLocationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigFileServer) GetAdminSection(context.Context, *GetAdminSectionRequest) (*GetAdminSectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigFileServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigFileServer) SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigFileServer) ClearInvalidSections(context.Context, *ClearInvalidSectionsRequest) (*ClearInvalidSectionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigFileServer) GetRootSectionGroup(context.Context, *GetRootSectionGroupRequest) (*GetRootSectionGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostConfigFileServer = (*UnimplementedAppHostConfigFileServer)(nil)
