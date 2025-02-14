package icertadmind

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

// ICertAdminD server interface.
type CertAdminDServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The SetExtension method allows adding, modifying, or disabling of extensions, as
	// specified in [RFC3280]. A CA can include an extension in an issued certificate for
	// a particular pending request.
	SetExtension(context.Context, *SetExtensionRequest) (*SetExtensionResponse, error)

	// The SetAttributes method sets attributes in the specified pending certificate request.
	SetAttributes(context.Context, *SetAttributesRequest) (*SetAttributesResponse, error)

	// The ResubmitRequest method resubmits a specific pending or denied certificate request
	// to the CA.
	ResubmitRequest(context.Context, *ResubmitRequestRequest) (*ResubmitRequestResponse, error)

	// The DenyRequest method denies a specific certificate request that is pending.
	DenyRequest(context.Context, *DenyRequestRequest) (*DenyRequestResponse, error)

	// The IsValidCertificate method verifies the certificate against the CA key and verifies
	// that the certificate has not been revoked.
	IsValidCertificate(context.Context, *IsValidCertificateRequest) (*IsValidCertificateResponse, error)

	// The PublishCRL method sends a request to the CA server to publish a new CRL.
	PublishCRL(context.Context, *PublishCRLRequest) (*PublishCRLResponse, error)

	// The GetCRL method instructs the CA to return the current base CRL for the current
	// CA key.
	GetCRL(context.Context, *GetCRLRequest) (*GetCRLResponse, error)

	// The RevokeCertificate method revokes a certificate either immediately or on a specified
	// date. It instructs the CA to revoke a certificate based on the certificate's serial
	// number and reason code.
	RevokeCertificate(context.Context, *RevokeCertificateRequest) (*RevokeCertificateResponse, error)

	// The EnumViewColumn method returns an array of column information.
	EnumViewColumn(context.Context, *EnumViewColumnRequest) (*EnumViewColumnResponse, error)

	// The GetViewDefaultColumnSet method returns an array of column identifiers that are
	// associated with a specific view.
	GetViewDefaultColumnSet(context.Context, *GetViewDefaultColumnSetRequest) (*GetViewDefaultColumnSetResponse, error)

	// The EnumAttributesOrExtensions method is used to access sets of attributes or extensions
	// for a particular row ID.
	EnumAttributesOrExtensions(context.Context, *EnumAttributesOrExtensionsRequest) (*EnumAttributesOrExtensionsResponse, error)

	// The OpenView method opens a view into the database and returns a set of resultant
	// row data.
	OpenView(context.Context, *OpenViewRequest) (*OpenViewResponse, error)

	// The EnumView method returns a set of resultant row data for the opened view.
	EnumView(context.Context, *EnumViewRequest) (*EnumViewResponse, error)

	// The CloseView method closes a view that was previously opened by using the OpenView
	// method call.
	CloseView(context.Context, *CloseViewRequest) (*CloseViewResponse, error)

	// The ServerControl method is used to force the CA server to unregister the ICertAdminD
	// and ICertAdminD2 interfaces.
	ServerControl(context.Context, *ServerControlRequest) (*ServerControlResponse, error)

	// The Ping method is used to test whether the certificate server is alive.
	Ping(context.Context, *PingRequest) (*PingResponse, error)

	// The GetServerState method is used to validate that the caller has permission to read
	// the CA database.
	GetServerState(context.Context, *GetServerStateRequest) (*GetServerStateResponse, error)

	// The BackupPrepare method is used to prepare the database for performing further backup
	// operations, such as BackupEnd, BackupGetAttachmentInformation, BackupGetBackupLogs,
	// BackupOpenFile, BackupReadFile, BackupCloseFile, and BackupTruncateLogs.
	BackupPrepare(context.Context, *BackupPrepareRequest) (*BackupPrepareResponse, error)

	// The BackupEnd method completes the backup process that is started via a call to ICertAdminD::BackupPrepare.
	//
	// This method has no parameters.
	//
	// If Config_CA_Interface_Flags contains the value IF_NOREMOTEICERTADMINBACKUP, the
	// server SHOULD return an error.<48>
	BackupEnd(context.Context, *BackupEndRequest) (*BackupEndResponse, error)

	// The BackupGetAttachmentInformation method is used to query the CA for the names of
	// database files that should become part of the backup file set.
	BackupGetAttachmentInformation(context.Context, *BackupGetAttachmentInformationRequest) (*BackupGetAttachmentInformationResponse, error)

	// The BackupGetBackupLogs method queries the CA for the names of database transaction
	// log files that should become part of the backup file set.
	BackupGetBackupLogs(context.Context, *BackupGetBackupLogsRequest) (*BackupGetBackupLogsResponse, error)

	// The BackupOpenFile method opens a file for backup.
	BackupOpenFile(context.Context, *BackupOpenFileRequest) (*BackupOpenFileResponse, error)

	// The BackupReadFile method reads the database file and loads the contents into the
	// buffer that is provided. The file MUST be initialized by a prior call to BackupOpenFile.
	BackupReadFile(context.Context, *BackupReadFileRequest) (*BackupReadFileResponse, error)

	// The BackupCloseFile method closes the database file that was initialized by a prior
	// call to the BackupOpenFile.
	//
	// This method has no parameters.
	//
	// If Config_CA_Interface_Flags contains the value IF_NOREMOTEICERTADMINBACKUP, the
	// server SHOULD return an error.<58>
	BackupCloseFile(context.Context, *BackupCloseFileRequest) (*BackupCloseFileResponse, error)

	// The BackupTruncateLogs method function eliminates redundant records from the log
	// files and reduces the disk storage space that is used by log files.
	//
	// This method has no parameters.
	//
	// If Config_CA_Interface_Flags contains the value IF_NOREMOTEICERTADMINBACKUP, the
	// server SHOULD return an error.<59>
	BackupTruncateLogs(context.Context, *BackupTruncateLogsRequest) (*BackupTruncateLogsResponse, error)

	// The ImportCertificate method imports a certificate into the CA database.
	ImportCertificate(context.Context, *ImportCertificateRequest) (*ImportCertificateResponse, error)

	// The BackupGetDynamicFiles method retrieves the list of CA dynamic file names that
	// need to be backed up. The dynamic files are those that are not included in the CA
	// database backup and are created dynamically by the CA, for example: CRL files created
	// by the CA. Note that BackupOpenFile and BackupReadFile cannot be used to open and
	// read the dynamic files whose names are returned by this method. Dynamic files must
	// be backed up by means outside this protocol.
	BackupGetDynamicFiles(context.Context, *BackupGetDynamicFilesRequest) (*BackupGetDynamicFilesResponse, error)

	// The RestoreGetDatabaseLocation method retrieves the list of CA server database location
	// names for all the database files being restored.
	RestoreGetDatabaseLocations(context.Context, *RestoreGetDatabaseLocationsRequest) (*RestoreGetDatabaseLocationsResponse, error)
}

func RegisterCertAdminDServer(conn dcerpc.Conn, o CertAdminDServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCertAdminDServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CertAdminDSyntaxV0_0))...)
}

func NewCertAdminDServerHandle(o CertAdminDServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CertAdminDServerHandle(ctx, o, opNum, r)
	}
}

func CertAdminDServerHandle(ctx context.Context, o CertAdminDServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // SetExtension
		op := &xxx_SetExtensionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExtensionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExtension(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // SetAttributes
		op := &xxx_SetAttributesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAttributesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAttributes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ResubmitRequest
		op := &xxx_ResubmitRequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResubmitRequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResubmitRequest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // DenyRequest
		op := &xxx_DenyRequestOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DenyRequestRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DenyRequest(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // IsValidCertificate
		op := &xxx_IsValidCertificateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsValidCertificateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsValidCertificate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // PublishCRL
		op := &xxx_PublishCRLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PublishCRLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PublishCRL(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // GetCRL
		op := &xxx_GetCRLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCRLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCRL(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RevokeCertificate
		op := &xxx_RevokeCertificateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RevokeCertificateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RevokeCertificate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // EnumViewColumn
		op := &xxx_EnumViewColumnOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumViewColumnRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumViewColumn(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // GetViewDefaultColumnSet
		op := &xxx_GetViewDefaultColumnSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetViewDefaultColumnSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetViewDefaultColumnSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // EnumAttributesOrExtensions
		op := &xxx_EnumAttributesOrExtensionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAttributesOrExtensionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAttributesOrExtensions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // OpenView
		op := &xxx_OpenViewOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenViewRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenView(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // EnumView
		op := &xxx_EnumViewOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumViewRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumView(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // CloseView
		op := &xxx_CloseViewOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseViewRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseView(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // ServerControl
		op := &xxx_ServerControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // Ping
		op := &xxx_PingOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PingRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Ping(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // GetServerState
		op := &xxx_GetServerStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // BackupPrepare
		op := &xxx_BackupPrepareOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupPrepareRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupPrepare(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // BackupEnd
		op := &xxx_BackupEndOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupEndRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupEnd(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // BackupGetAttachmentInformation
		op := &xxx_BackupGetAttachmentInformationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupGetAttachmentInformationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupGetAttachmentInformation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // BackupGetBackupLogs
		op := &xxx_BackupGetBackupLogsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupGetBackupLogsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupGetBackupLogs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // BackupOpenFile
		op := &xxx_BackupOpenFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupOpenFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupOpenFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // BackupReadFile
		op := &xxx_BackupReadFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupReadFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupReadFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // BackupCloseFile
		op := &xxx_BackupCloseFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupCloseFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupCloseFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // BackupTruncateLogs
		op := &xxx_BackupTruncateLogsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupTruncateLogsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupTruncateLogs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // ImportCertificate
		op := &xxx_ImportCertificateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportCertificateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportCertificate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // BackupGetDynamicFiles
		op := &xxx_BackupGetDynamicFilesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupGetDynamicFilesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupGetDynamicFiles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // RestoreGetDatabaseLocations
		op := &xxx_RestoreGetDatabaseLocationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestoreGetDatabaseLocationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RestoreGetDatabaseLocations(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICertAdminD
type UnimplementedCertAdminDServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedCertAdminDServer) SetExtension(context.Context, *SetExtensionRequest) (*SetExtensionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) SetAttributes(context.Context, *SetAttributesRequest) (*SetAttributesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) ResubmitRequest(context.Context, *ResubmitRequestRequest) (*ResubmitRequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) DenyRequest(context.Context, *DenyRequestRequest) (*DenyRequestResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) IsValidCertificate(context.Context, *IsValidCertificateRequest) (*IsValidCertificateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) PublishCRL(context.Context, *PublishCRLRequest) (*PublishCRLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) GetCRL(context.Context, *GetCRLRequest) (*GetCRLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) RevokeCertificate(context.Context, *RevokeCertificateRequest) (*RevokeCertificateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) EnumViewColumn(context.Context, *EnumViewColumnRequest) (*EnumViewColumnResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) GetViewDefaultColumnSet(context.Context, *GetViewDefaultColumnSetRequest) (*GetViewDefaultColumnSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) EnumAttributesOrExtensions(context.Context, *EnumAttributesOrExtensionsRequest) (*EnumAttributesOrExtensionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) OpenView(context.Context, *OpenViewRequest) (*OpenViewResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) EnumView(context.Context, *EnumViewRequest) (*EnumViewResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) CloseView(context.Context, *CloseViewRequest) (*CloseViewResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) ServerControl(context.Context, *ServerControlRequest) (*ServerControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) GetServerState(context.Context, *GetServerStateRequest) (*GetServerStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) BackupPrepare(context.Context, *BackupPrepareRequest) (*BackupPrepareResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) BackupEnd(context.Context, *BackupEndRequest) (*BackupEndResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) BackupGetAttachmentInformation(context.Context, *BackupGetAttachmentInformationRequest) (*BackupGetAttachmentInformationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) BackupGetBackupLogs(context.Context, *BackupGetBackupLogsRequest) (*BackupGetBackupLogsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) BackupOpenFile(context.Context, *BackupOpenFileRequest) (*BackupOpenFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) BackupReadFile(context.Context, *BackupReadFileRequest) (*BackupReadFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) BackupCloseFile(context.Context, *BackupCloseFileRequest) (*BackupCloseFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) BackupTruncateLogs(context.Context, *BackupTruncateLogsRequest) (*BackupTruncateLogsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) ImportCertificate(context.Context, *ImportCertificateRequest) (*ImportCertificateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) BackupGetDynamicFiles(context.Context, *BackupGetDynamicFilesRequest) (*BackupGetDynamicFilesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminDServer) RestoreGetDatabaseLocations(context.Context, *RestoreGetDatabaseLocationsRequest) (*RestoreGetDatabaseLocationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CertAdminDServer = (*UnimplementedCertAdminDServer)(nil)
