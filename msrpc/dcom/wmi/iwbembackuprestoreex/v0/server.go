package iwbembackuprestoreex

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iwbembackuprestore "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbembackuprestore/v0"
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
	_ = iwbembackuprestore.GoPackage
)

// IWbemBackupRestoreEx server interface.
type BackupRestoreExServer interface {

	// IWbemBackupRestore base class.
	iwbembackuprestore.BackupRestoreServer

	// On the IWbemBackupRestoreEx::Pause method invocation, the server MUST set the IsServerPaused
	// flag to True and MUST persist the CIM database in a consistent state.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// If Pause is called and the IsServerPaused flag is set to True, the server MUST return
	// WBEM_E_INVALID_OPERATION. In case of any other failure, the server MUST return an
	// HRESULT whose S (severity) bit is set as specified in [MS-ERREF] section 2.1. The
	// actual HRESULT value is implementation dependent.
	//
	// The IWbemBackupRestoreEx::Pause method MUST be called on the interface that is obtained
	// from the DCOM Remote Protocol activation of the CLSID_WbemBackupRestore interface,
	// as specified in this section.
	Pause(context.Context, *PauseRequest) (*PauseResponse, error)

	// On the IWbemBackupRestoreEx::Resume method invocation, the server MUST set the IsServerPaused
	// flag to False.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return a WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// If Resume is called and the IsServerPaused flag is set to False, the server MUST
	// return WBEM_E_INVALID_OPERATION.
	//
	// In case of any other failure, the server MUST return an HRESULT whose S (severity)
	// bit is set as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	Resume(context.Context, *ResumeRequest) (*ResumeResponse, error)
}

func RegisterBackupRestoreExServer(conn dcerpc.Conn, o BackupRestoreExServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewBackupRestoreExServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(BackupRestoreExSyntaxV0_0))...)
}

func NewBackupRestoreExServerHandle(o BackupRestoreExServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return BackupRestoreExServerHandle(ctx, o, opNum, r)
	}
}

func BackupRestoreExServerHandle(ctx context.Context, o BackupRestoreExServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 5 {
		// IWbemBackupRestore base method.
		return iwbembackuprestore.BackupRestoreServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 5: // Pause
		in := &PauseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Pause(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // Resume
		in := &ResumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Resume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
