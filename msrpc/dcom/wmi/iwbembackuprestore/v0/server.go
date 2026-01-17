package iwbembackuprestore

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

// IWbemBackupRestore server interface.
type BackupRestoreServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// On the IWbemBackupRestore::Backup method invocation, the server MUST back up the
	// contents of the CIM database.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// In case of failure, the server MUST return an HRESULT whose S (severity) bit is set
	// as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	//
	// The IWbemBackupRestore::Backup method MUST be called on the interface that is obtained
	// from the DCOM Remote Protocol activation of a CLSID_WbemBackupRestore interface,
	// as specified in this section.
	Backup(context.Context, *BackupRequest) (*BackupResponse, error)

	// On the IWbemBackupRestore::Restore method invocation, the server MUST restore the
	// contents of the CIM database.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// If the WBEM_FLAG_BACKUP_RESTORE_FORCE_SHUTDOWN flag is not set, the server MUST return
	// WBEM_E_INVALID_PARAMETER.
	//
	// In case of failure, the server MUST return an HRESULT whose S (severity) bit is set
	// as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	Restore(context.Context, *RestoreRequest) (*RestoreResponse, error)
}

func RegisterBackupRestoreServer(conn dcerpc.Conn, o BackupRestoreServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewBackupRestoreServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(BackupRestoreSyntaxV0_0))...)
}

func NewBackupRestoreServerHandle(o BackupRestoreServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return BackupRestoreServerHandle(ctx, o, opNum, r)
	}
}

func BackupRestoreServerHandle(ctx context.Context, o BackupRestoreServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Backup
		op := &xxx_BackupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Backup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Restore
		op := &xxx_RestoreOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestoreRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Restore(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWbemBackupRestore
type UnimplementedBackupRestoreServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedBackupRestoreServer) Backup(context.Context, *BackupRequest) (*BackupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedBackupRestoreServer) Restore(context.Context, *RestoreRequest) (*RestoreResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ BackupRestoreServer = (*UnimplementedBackupRestoreServer)(nil)
