package iapphostconfigfile

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = iunknown.GoPackage
)

// IAppHostConfigFile server interface.
type AppHostConfigFileServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// ConfigPath operation.
	GetConfigPath(context.Context, *GetConfigPathRequest) (*GetConfigPathResponse, error)

	// The FilePath method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the server operating system file path that corresponds to the
	// specific IAppHostConfigFile, if that path applies to the administration system implementation.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *pbstrFilePath is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+---------------------------------------------------------+
	//	|               RETURN               |                                                         |
	//	|             VALUE/CODE             |                       DESCRIPTION                       |
	//	|                                    |                                                         |
	//	+------------------------------------+---------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                   |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.           |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command. |
	//	+------------------------------------+---------------------------------------------------------+
	GetFilePath(context.Context, *GetFilePathRequest) (*GetFilePathResponse, error)

	// The Locations method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns any hierarchy subpaths that exist in the specified IAppHostConfigFile.
	// These subpaths are returned in the form of a collection object that implements the
	// IAppHostConfigLocationCollection.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppLocations is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	|               RETURN               |                                                                         |
	//	|             VALUE/CODE             |                               DESCRIPTION                               |
	//	|                                    |                                                                         |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                                   |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                           |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                 |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X80070002 ERROR_FILE_NOT_FOUND    | A server resource (for example, a file on a disk) could not be found.   |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X80070005 ERROR_ACCESS_DENIED     | Access to a server resource (for example, a file on a disk) was denied. |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted.  |
	//	+------------------------------------+-------------------------------------------------------------------------+
	GetLocations(context.Context, *GetLocationsRequest) (*GetLocationsResponse, error)

	// GetAdminSection operation.
	GetAdminSection(context.Context, *GetAdminSectionRequest) (*GetAdminSectionResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// SetMetadata operation.
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)

	// The ClearInvalidSections method is received by the server in an RPC_REQUEST packet.
	// In response, the server MAY clear any invalid IAppHostElement objects that exist
	// in the specific IAppHostConfigFile.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result (as specified in [MS-ERREF] section 2) on failure. All failure results MUST
	// be treated identically. The following table specifies failure results specific to
	// this method.
	//
	//	+--------------------------------+-------------------------------+
	//	|             RETURN             |                               |
	//	|           VALUE/CODE           |          DESCRIPTION          |
	//	|                                |                               |
	//	+--------------------------------+-------------------------------+
	//	+--------------------------------+-------------------------------+
	//	| 0x80070032 ERROR_NOT_SUPPORTED | The request is not supported. |
	//	+--------------------------------+-------------------------------+
	ClearInvalidSections(context.Context, *ClearInvalidSectionsRequest) (*ClearInvalidSectionsResponse, error)

	// The RootSectionGroup method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns an IAppHostSectionGroup object, which represents a declaration
	// of IAppHostElement objects that apply to the specified IAppHostConfigFile and that
	// potentially apply to other IAppHostConfigFile at deeper hierarchy paths than the
	// current file.
	//
	// A declaration means that an IAppHostElement of the name in the declaration MAY exist
	// in the specified IAppHostConfigFile or potentially in deeper paths. A declaration
	// is NOT a definition/instance. A declaration only controls whether an actual IAppHostElement
	// instance is supported.
	//
	// This function returns the section group object in the configSections section group
	// that defines the root section group for the current section.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppSectionGroups is not NULL. If
	// processing fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF].
	// The following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+---------------------------------------------------------+
	//	|               RETURN               |                                                         |
	//	|             VALUE/CODE             |                       DESCRIPTION                       |
	//	|                                    |                                                         |
	//	+------------------------------------+---------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                   |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.           |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command. |
	//	+------------------------------------+---------------------------------------------------------+
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
