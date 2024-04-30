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
		in := &SetExtensionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetExtension(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // SetAttributes
		in := &SetAttributesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAttributes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // ResubmitRequest
		in := &ResubmitRequestRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResubmitRequest(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // DenyRequest
		in := &DenyRequestRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DenyRequest(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // IsValidCertificate
		in := &IsValidCertificateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsValidCertificate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // PublishCRL
		in := &PublishCRLRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PublishCRL(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // GetCRL
		in := &GetCRLRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCRL(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // RevokeCertificate
		in := &RevokeCertificateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RevokeCertificate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // EnumViewColumn
		in := &EnumViewColumnRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumViewColumn(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // GetViewDefaultColumnSet
		in := &GetViewDefaultColumnSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetViewDefaultColumnSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // EnumAttributesOrExtensions
		in := &EnumAttributesOrExtensionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumAttributesOrExtensions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // OpenView
		in := &OpenViewRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenView(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // EnumView
		in := &EnumViewRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumView(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // CloseView
		in := &CloseViewRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseView(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // ServerControl
		in := &ServerControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // Ping
		in := &PingRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Ping(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // GetServerState
		in := &GetServerStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServerState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // BackupPrepare
		in := &BackupPrepareRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupPrepare(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // BackupEnd
		in := &BackupEndRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupEnd(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // BackupGetAttachmentInformation
		in := &BackupGetAttachmentInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupGetAttachmentInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // BackupGetBackupLogs
		in := &BackupGetBackupLogsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupGetBackupLogs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // BackupOpenFile
		in := &BackupOpenFileRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupOpenFile(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // BackupReadFile
		in := &BackupReadFileRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupReadFile(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // BackupCloseFile
		in := &BackupCloseFileRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupCloseFile(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // BackupTruncateLogs
		in := &BackupTruncateLogsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupTruncateLogs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // ImportCertificate
		in := &ImportCertificateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ImportCertificate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // BackupGetDynamicFiles
		in := &BackupGetDynamicFilesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupGetDynamicFiles(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // RestoreGetDatabaseLocations
		in := &RestoreGetDatabaseLocationsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RestoreGetDatabaseLocations(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
