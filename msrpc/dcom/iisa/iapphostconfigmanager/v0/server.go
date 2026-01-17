package iapphostconfigmanager

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

// IAppHostConfigManager server interface.
type AppHostConfigManagerServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetConfigFile method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns an IAppHostConfigFile for the specific hierarchy path.
	//
	// The administration system implementation can choose to fail if the specified hierarchy
	// path does not have an IAppHostConfigFile container for it. Or it can choose to succeed
	// and create an empty IAppHostConfigFile container instead.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppConfigFile is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted.           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070002 ERROR_FILE_NOT_FOUND    | The server resource (for example, a file or database) corresponding to           |
	//	|                                    | bstrConfigPath could not be found.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070005 ERROR_ACCESS_DENIED     | Access to a server resource (for example, a file on a disk) was denied.          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	GetConfigFile(context.Context, *GetConfigFileRequest) (*GetConfigFileResponse, error)

	// The GetUniqueConfigPath method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns the deepest hierarchy path (up to the specified hierarchy
	// path) that contains a unique set of IAppHostElement objects. For example:
	//
	// Assume: At hierarchy path A, a set of IAppHostElement objects exist.
	//
	// Assume: At hierarchy path B (deeper than A), the identical set of objects exists.
	//
	// Given these assumptions, GetUniqueConfigPath( B ) returns path A. In other words,
	// the method returns the shallowest path that contains the identical set of IAppHostElement
	// objects as the specified path.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *pbstrUniquePath is not NULL. If
	// processing fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF].
	// The following table describes the error conditions that MUST be handled and the corresponding
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
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted.  |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X00000002 ERROR_PATH_NOT_FOUND    | A server resource (for example, a file on a disk) could not be found.   |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X80070005 ERROR_ACCESS_DENIED     | Access to a server resource (for example, a file on a disk) was denied. |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                 |
	//	+------------------------------------+-------------------------------------------------------------------------+
	GetUniqueConfigPath(context.Context, *GetUniqueConfigPathRequest) (*GetUniqueConfigPathResponse, error)
}

func RegisterAppHostConfigManagerServer(conn dcerpc.Conn, o AppHostConfigManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostConfigManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostConfigManagerSyntaxV0_0))...)
}

func NewAppHostConfigManagerServerHandle(o AppHostConfigManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostConfigManagerServerHandle(ctx, o, opNum, r)
	}
}

func AppHostConfigManagerServerHandle(ctx context.Context, o AppHostConfigManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetConfigFile
		op := &xxx_GetConfigFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfigFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetUniqueConfigPath
		op := &xxx_GetUniqueConfigPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUniqueConfigPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUniqueConfigPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostConfigManager
type UnimplementedAppHostConfigManagerServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostConfigManagerServer) GetConfigFile(context.Context, *GetConfigFileRequest) (*GetConfigFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigManagerServer) GetUniqueConfigPath(context.Context, *GetUniqueConfigPathRequest) (*GetUniqueConfigPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostConfigManagerServer = (*UnimplementedAppHostConfigManagerServer)(nil)
