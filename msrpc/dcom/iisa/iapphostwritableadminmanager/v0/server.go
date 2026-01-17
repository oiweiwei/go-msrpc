package iapphostwritableadminmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iapphostadminmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostadminmanager/v0"
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
	_ = iapphostadminmanager.GoPackage
)

// IAppHostWritableAdminManager server interface.
type AppHostWritableAdminManagerServer interface {

	// IAppHostAdminManager base class.
	iapphostadminmanager.AppHostAdminManagerServer

	// The CommitChanges method is received by the server in an RPC_REQUEST packet. In response,
	// the server commits any in-memory changes that it accumulates to a persisted store.
	// This behavior essentially writes out the changes that are made by the client.
	//
	// This method has no parameters.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
	//
	//	+------------------------------------+------------------------------------------------------------------------+
	//	|               RETURN               |                                                                        |
	//	|             VALUE/CODE             |                              DESCRIPTION                               |
	//	|                                    |                                                                        |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                                  |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted. |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X00000002 ERROR_PATH_NOT_FOUND    | The system cannot find the path specified.                             |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070002 ERROR_FILE_NOT_FOUND    | The system cannot find the file specified.                             |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070005 ERROR_ACCESS_DENIED     | Access is denied.                                                      |
	//	+------------------------------------+------------------------------------------------------------------------+
	CommitChanges(context.Context, *CommitChangesRequest) (*CommitChangesResponse, error)

	// CommitPath operation.
	GetCommitPath(context.Context, *GetCommitPathRequest) (*GetCommitPathResponse, error)

	// CommitPath operation.
	SetCommitPath(context.Context, *SetCommitPathRequest) (*SetCommitPathResponse, error)
}

func RegisterAppHostWritableAdminManagerServer(conn dcerpc.Conn, o AppHostWritableAdminManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostWritableAdminManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostWritableAdminManagerSyntaxV0_0))...)
}

func NewAppHostWritableAdminManagerServerHandle(o AppHostWritableAdminManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostWritableAdminManagerServerHandle(ctx, o, opNum, r)
	}
}

func AppHostWritableAdminManagerServerHandle(ctx context.Context, o AppHostWritableAdminManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IAppHostAdminManager base method.
		return iapphostadminmanager.AppHostAdminManagerServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CommitChanges
		op := &xxx_CommitChangesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommitChangesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CommitChanges(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CommitPath
		op := &xxx_GetCommitPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCommitPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCommitPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // CommitPath
		op := &xxx_SetCommitPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCommitPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCommitPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostWritableAdminManager
type UnimplementedAppHostWritableAdminManagerServer struct {
	iapphostadminmanager.UnimplementedAppHostAdminManagerServer
}

func (UnimplementedAppHostWritableAdminManagerServer) CommitChanges(context.Context, *CommitChangesRequest) (*CommitChangesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostWritableAdminManagerServer) GetCommitPath(context.Context, *GetCommitPathRequest) (*GetCommitPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostWritableAdminManagerServer) SetCommitPath(context.Context, *SetCommitPathRequest) (*SetCommitPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostWritableAdminManagerServer = (*UnimplementedAppHostWritableAdminManagerServer)(nil)
