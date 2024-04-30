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
		in := &GetConfigPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetConfigPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // FilePath
		in := &GetFilePathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFilePath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Locations
		in := &GetLocationsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLocations(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // GetAdminSection
		in := &GetAdminSectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAdminSection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // GetMetadata
		in := &GetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // SetMetadata
		in := &SetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // ClearInvalidSections
		in := &ClearInvalidSectionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClearInvalidSections(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // RootSectionGroup
		in := &GetRootSectionGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRootSectionGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
