package icertadmind

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	csra "github.com/oiweiwei/go-msrpc/msrpc/dcom/csra"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = csra.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csra"
)

var (
	// ICertAdminD interface identifier d99e6e71-fc88-11d0-b498-00a0c90312f3
	CertAdminDIID = &dcom.IID{Data1: 0xd99e6e71, Data2: 0xfc88, Data3: 0x11d0, Data4: []byte{0xb4, 0x98, 0x00, 0xa0, 0xc9, 0x03, 0x12, 0xf3}}
	// Syntax UUID
	CertAdminDSyntaxUUID = &uuid.UUID{TimeLow: 0xd99e6e71, TimeMid: 0xfc88, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x98, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x3, 0x12, 0xf3}}
	// Syntax ID
	CertAdminDSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CertAdminDSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICertAdminD interface.
type CertAdminDClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The SetExtension method allows adding, modifying, or disabling of extensions, as
	// specified in [RFC3280]. A CA can include an extension in an issued certificate for
	// a particular pending request.
	SetExtension(context.Context, *SetExtensionRequest, ...dcerpc.CallOption) (*SetExtensionResponse, error)

	// The SetAttributes method sets attributes in the specified pending certificate request.
	SetAttributes(context.Context, *SetAttributesRequest, ...dcerpc.CallOption) (*SetAttributesResponse, error)

	// The ResubmitRequest method resubmits a specific pending or denied certificate request
	// to the CA.
	ResubmitRequest(context.Context, *ResubmitRequestRequest, ...dcerpc.CallOption) (*ResubmitRequestResponse, error)

	// The DenyRequest method denies a specific certificate request that is pending.
	DenyRequest(context.Context, *DenyRequestRequest, ...dcerpc.CallOption) (*DenyRequestResponse, error)

	// The IsValidCertificate method verifies the certificate against the CA key and verifies
	// that the certificate has not been revoked.
	IsValidCertificate(context.Context, *IsValidCertificateRequest, ...dcerpc.CallOption) (*IsValidCertificateResponse, error)

	// The PublishCRL method sends a request to the CA server to publish a new CRL.
	PublishCRL(context.Context, *PublishCRLRequest, ...dcerpc.CallOption) (*PublishCRLResponse, error)

	// The GetCRL method instructs the CA to return the current base CRL for the current
	// CA key.
	GetCRL(context.Context, *GetCRLRequest, ...dcerpc.CallOption) (*GetCRLResponse, error)

	// The RevokeCertificate method revokes a certificate either immediately or on a specified
	// date. It instructs the CA to revoke a certificate based on the certificate's serial
	// number and reason code.
	RevokeCertificate(context.Context, *RevokeCertificateRequest, ...dcerpc.CallOption) (*RevokeCertificateResponse, error)

	// The EnumViewColumn method returns an array of column information.
	EnumViewColumn(context.Context, *EnumViewColumnRequest, ...dcerpc.CallOption) (*EnumViewColumnResponse, error)

	// The GetViewDefaultColumnSet method returns an array of column identifiers that are
	// associated with a specific view.
	GetViewDefaultColumnSet(context.Context, *GetViewDefaultColumnSetRequest, ...dcerpc.CallOption) (*GetViewDefaultColumnSetResponse, error)

	// The EnumAttributesOrExtensions method is used to access sets of attributes or extensions
	// for a particular row ID.
	EnumAttributesOrExtensions(context.Context, *EnumAttributesOrExtensionsRequest, ...dcerpc.CallOption) (*EnumAttributesOrExtensionsResponse, error)

	// The OpenView method opens a view into the database and returns a set of resultant
	// row data.
	OpenView(context.Context, *OpenViewRequest, ...dcerpc.CallOption) (*OpenViewResponse, error)

	// The EnumView method returns a set of resultant row data for the opened view.
	EnumView(context.Context, *EnumViewRequest, ...dcerpc.CallOption) (*EnumViewResponse, error)

	// The CloseView method closes a view that was previously opened by using the OpenView
	// method call.
	CloseView(context.Context, *CloseViewRequest, ...dcerpc.CallOption) (*CloseViewResponse, error)

	// The ServerControl method is used to force the CA server to unregister the ICertAdminD
	// and ICertAdminD2 interfaces.
	ServerControl(context.Context, *ServerControlRequest, ...dcerpc.CallOption) (*ServerControlResponse, error)

	// The Ping method is used to test whether the certificate server is alive.
	Ping(context.Context, *PingRequest, ...dcerpc.CallOption) (*PingResponse, error)

	// The GetServerState method is used to validate that the caller has permission to read
	// the CA database.
	GetServerState(context.Context, *GetServerStateRequest, ...dcerpc.CallOption) (*GetServerStateResponse, error)

	// The BackupPrepare method is used to prepare the database for performing further backup
	// operations, such as BackupEnd, BackupGetAttachmentInformation, BackupGetBackupLogs,
	// BackupOpenFile, BackupReadFile, BackupCloseFile, and BackupTruncateLogs.
	BackupPrepare(context.Context, *BackupPrepareRequest, ...dcerpc.CallOption) (*BackupPrepareResponse, error)

	// The BackupEnd method completes the backup process that is started via a call to ICertAdminD::BackupPrepare.
	//
	// This method has no parameters.
	//
	// If Config_CA_Interface_Flags contains the value IF_NOREMOTEICERTADMINBACKUP, the
	// server SHOULD return an error.<48>
	BackupEnd(context.Context, *BackupEndRequest, ...dcerpc.CallOption) (*BackupEndResponse, error)

	// The BackupGetAttachmentInformation method is used to query the CA for the names of
	// database files that should become part of the backup file set.
	BackupGetAttachmentInformation(context.Context, *BackupGetAttachmentInformationRequest, ...dcerpc.CallOption) (*BackupGetAttachmentInformationResponse, error)

	// The BackupGetBackupLogs method queries the CA for the names of database transaction
	// log files that should become part of the backup file set.
	BackupGetBackupLogs(context.Context, *BackupGetBackupLogsRequest, ...dcerpc.CallOption) (*BackupGetBackupLogsResponse, error)

	// The BackupOpenFile method opens a file for backup.
	BackupOpenFile(context.Context, *BackupOpenFileRequest, ...dcerpc.CallOption) (*BackupOpenFileResponse, error)

	// The BackupReadFile method reads the database file and loads the contents into the
	// buffer that is provided. The file MUST be initialized by a prior call to BackupOpenFile.
	BackupReadFile(context.Context, *BackupReadFileRequest, ...dcerpc.CallOption) (*BackupReadFileResponse, error)

	// The BackupCloseFile method closes the database file that was initialized by a prior
	// call to the BackupOpenFile.
	//
	// This method has no parameters.
	//
	// If Config_CA_Interface_Flags contains the value IF_NOREMOTEICERTADMINBACKUP, the
	// server SHOULD return an error.<58>
	BackupCloseFile(context.Context, *BackupCloseFileRequest, ...dcerpc.CallOption) (*BackupCloseFileResponse, error)

	// The BackupTruncateLogs method function eliminates redundant records from the log
	// files and reduces the disk storage space that is used by log files.
	//
	// This method has no parameters.
	//
	// If Config_CA_Interface_Flags contains the value IF_NOREMOTEICERTADMINBACKUP, the
	// server SHOULD return an error.<59>
	BackupTruncateLogs(context.Context, *BackupTruncateLogsRequest, ...dcerpc.CallOption) (*BackupTruncateLogsResponse, error)

	// The ImportCertificate method imports a certificate into the CA database.
	ImportCertificate(context.Context, *ImportCertificateRequest, ...dcerpc.CallOption) (*ImportCertificateResponse, error)

	// The BackupGetDynamicFiles method retrieves the list of CA dynamic file names that
	// need to be backed up. The dynamic files are those that are not included in the CA
	// database backup and are created dynamically by the CA, for example: CRL files created
	// by the CA. Note that BackupOpenFile and BackupReadFile cannot be used to open and
	// read the dynamic files whose names are returned by this method. Dynamic files must
	// be backed up by means outside this protocol.
	BackupGetDynamicFiles(context.Context, *BackupGetDynamicFilesRequest, ...dcerpc.CallOption) (*BackupGetDynamicFilesResponse, error)

	// The RestoreGetDatabaseLocation method retrieves the list of CA server database location
	// names for all the database files being restored.
	RestoreGetDatabaseLocations(context.Context, *RestoreGetDatabaseLocationsRequest, ...dcerpc.CallOption) (*RestoreGetDatabaseLocationsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CertAdminDClient
}

type xxx_DefaultCertAdminDClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCertAdminDClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCertAdminDClient) SetExtension(ctx context.Context, in *SetExtensionRequest, opts ...dcerpc.CallOption) (*SetExtensionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetExtensionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) SetAttributes(ctx context.Context, in *SetAttributesRequest, opts ...dcerpc.CallOption) (*SetAttributesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetAttributesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) ResubmitRequest(ctx context.Context, in *ResubmitRequestRequest, opts ...dcerpc.CallOption) (*ResubmitRequestResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResubmitRequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) DenyRequest(ctx context.Context, in *DenyRequestRequest, opts ...dcerpc.CallOption) (*DenyRequestResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DenyRequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) IsValidCertificate(ctx context.Context, in *IsValidCertificateRequest, opts ...dcerpc.CallOption) (*IsValidCertificateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsValidCertificateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) PublishCRL(ctx context.Context, in *PublishCRLRequest, opts ...dcerpc.CallOption) (*PublishCRLResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PublishCRLResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) GetCRL(ctx context.Context, in *GetCRLRequest, opts ...dcerpc.CallOption) (*GetCRLResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetCRLResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) RevokeCertificate(ctx context.Context, in *RevokeCertificateRequest, opts ...dcerpc.CallOption) (*RevokeCertificateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RevokeCertificateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) EnumViewColumn(ctx context.Context, in *EnumViewColumnRequest, opts ...dcerpc.CallOption) (*EnumViewColumnResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumViewColumnResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) GetViewDefaultColumnSet(ctx context.Context, in *GetViewDefaultColumnSetRequest, opts ...dcerpc.CallOption) (*GetViewDefaultColumnSetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetViewDefaultColumnSetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) EnumAttributesOrExtensions(ctx context.Context, in *EnumAttributesOrExtensionsRequest, opts ...dcerpc.CallOption) (*EnumAttributesOrExtensionsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumAttributesOrExtensionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) OpenView(ctx context.Context, in *OpenViewRequest, opts ...dcerpc.CallOption) (*OpenViewResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenViewResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) EnumView(ctx context.Context, in *EnumViewRequest, opts ...dcerpc.CallOption) (*EnumViewResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumViewResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) CloseView(ctx context.Context, in *CloseViewRequest, opts ...dcerpc.CallOption) (*CloseViewResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseViewResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) ServerControl(ctx context.Context, in *ServerControlRequest, opts ...dcerpc.CallOption) (*ServerControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) Ping(ctx context.Context, in *PingRequest, opts ...dcerpc.CallOption) (*PingResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) GetServerState(ctx context.Context, in *GetServerStateRequest, opts ...dcerpc.CallOption) (*GetServerStateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetServerStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) BackupPrepare(ctx context.Context, in *BackupPrepareRequest, opts ...dcerpc.CallOption) (*BackupPrepareResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupPrepareResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) BackupEnd(ctx context.Context, in *BackupEndRequest, opts ...dcerpc.CallOption) (*BackupEndResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupEndResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) BackupGetAttachmentInformation(ctx context.Context, in *BackupGetAttachmentInformationRequest, opts ...dcerpc.CallOption) (*BackupGetAttachmentInformationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupGetAttachmentInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) BackupGetBackupLogs(ctx context.Context, in *BackupGetBackupLogsRequest, opts ...dcerpc.CallOption) (*BackupGetBackupLogsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupGetBackupLogsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) BackupOpenFile(ctx context.Context, in *BackupOpenFileRequest, opts ...dcerpc.CallOption) (*BackupOpenFileResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupOpenFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) BackupReadFile(ctx context.Context, in *BackupReadFileRequest, opts ...dcerpc.CallOption) (*BackupReadFileResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupReadFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) BackupCloseFile(ctx context.Context, in *BackupCloseFileRequest, opts ...dcerpc.CallOption) (*BackupCloseFileResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupCloseFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) BackupTruncateLogs(ctx context.Context, in *BackupTruncateLogsRequest, opts ...dcerpc.CallOption) (*BackupTruncateLogsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupTruncateLogsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) ImportCertificate(ctx context.Context, in *ImportCertificateRequest, opts ...dcerpc.CallOption) (*ImportCertificateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ImportCertificateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) BackupGetDynamicFiles(ctx context.Context, in *BackupGetDynamicFilesRequest, opts ...dcerpc.CallOption) (*BackupGetDynamicFilesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupGetDynamicFilesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) RestoreGetDatabaseLocations(ctx context.Context, in *RestoreGetDatabaseLocationsRequest, opts ...dcerpc.CallOption) (*RestoreGetDatabaseLocationsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RestoreGetDatabaseLocationsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminDClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCertAdminDClient) IPID(ctx context.Context, ipid *dcom.IPID) CertAdminDClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCertAdminDClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewCertAdminDClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CertAdminDClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CertAdminDSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultCertAdminDClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_SetExtensionOperation structure represents the SetExtension operation
type xxx_SetExtensionOperation struct {
	This          *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority     string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	RequestID     uint32                  `idl:"name:dwRequestId" json:"request_id"`
	ExtensionName string                  `idl:"name:pwszExtensionName;string;pointer:unique" json:"extension_name"`
	Type          uint32                  `idl:"name:dwType" json:"type"`
	Flags         uint32                  `idl:"name:dwFlags" json:"flags"`
	Value         *csra.CertTransportBlob `idl:"name:pctbValue;pointer:ref" json:"value"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetExtensionOperation) OpNum() int { return 3 }

func (o *xxx_SetExtensionOperation) OpName() string { return "/ICertAdminD/v0/SetExtension" }

func (o *xxx_SetExtensionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtensionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	// pwszExtensionName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ExtensionName != "" {
			_ptr_pwszExtensionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ExtensionName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ExtensionName, _ptr_pwszExtensionName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pctbValue {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Value != nil {
			if err := o.Value.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csra.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtensionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	// pwszExtensionName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszExtensionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ExtensionName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszExtensionName := func(ptr interface{}) { o.ExtensionName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ExtensionName, _s_pwszExtensionName, _ptr_pwszExtensionName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pctbValue {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Value == nil {
			o.Value = &csra.CertTransportBlob{}
		}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtensionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtensionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExtensionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetExtensionRequest structure represents the SetExtension operation request
type SetExtensionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: A null-terminated Unicode string that contains the name of the certificate
	// server. The pwszAuthority is a Unicode string in the form of a distinguished name
	// (DN) value, such as "CAName", where CAName MUST be the full common name (CN) or sanitized
	// name of the CA, as specified in [MS-WCCE] section 3.1.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// dwRequestId: A 32-bit nonzero unsigned integer value that specifies the ID of the
	// certificate request.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
	// pwszExtensionName: A null-terminated Unicode string that specifies the OID (1) for
	// the extension to set, as specified in [X680]. The string MUST be 31 or fewer characters
	// in length and the characters MUST NOT be NULL.
	ExtensionName string `idl:"name:pwszExtensionName;string;pointer:unique" json:"extension_name"`
	// dwType: An unsigned integer value that specifies the type of extension being set.
	// The dwType parameter MUST agree with the data type of the pb member of the pctbValue
	// parameter. This parameter can be one of the following values.
	//
	//	+------------+--------------------+
	//	|            |                    |
	//	|   VALUE    |      MEANING       |
	//	|            |                    |
	//	+------------+--------------------+
	//	+------------+--------------------+
	//	| 0x00000001 | Unsigned long data |
	//	+------------+--------------------+
	//	| 0x00000002 | Date/time          |
	//	+------------+--------------------+
	//	| 0x00000003 | Binary data        |
	//	+------------+--------------------+
	//	| 0x00000004 | Unicode            |
	//	+------------+--------------------+
	Type uint32 `idl:"name:dwType" json:"type"`
	// dwFlags: An unsigned integer value that specifies the flags for the extension being
	// set. This parameter can be one of the following values.
	//
	//	+-------+--------------------------------------------------------+
	//	|       |                                                        |
	//	| VALUE |                        MEANING                         |
	//	|       |                                                        |
	//	+-------+--------------------------------------------------------+
	//	+-------+--------------------------------------------------------+
	//	|     1 | This is a critical extension.                          |
	//	+-------+--------------------------------------------------------+
	//	|     2 | The extension MUST NOT be used on issued certificates. |
	//	+-------+--------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pctbValue: A pointer to a CERTTRANSBLOB structure. The pb member MUST point to the
	// binary data for the extension and the cb member MUST contain the length, in bytes,
	// of the value. Depending on the value of the dwType parameter, the format of the binary
	// data that is pointed to by the pb member is shown in the following table.
	//
	//	+-----------------+----------------------------------------------------------------------------------+
	//	|    VALUE OF     |                                                                                  |
	//	|     DWTYPE      |                                     MEANING                                      |
	//	|                 |                                                                                  |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000001      | The CERTTRANSBLOB structure pb member MUST point to an unsigned long data value  |
	//	|                 | in little-endian format.                                                         |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000002      | The CERTTRANSBLOB structure pb member MUST point to data using little-endian     |
	//	|                 | encoding format.                                                                 |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000003      | The CERTTRANSBLOB structure pb member MUST point to an array of bytes that are   |
	//	|                 | not in need of endian forcing.                                                   |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000004      | The CERTTRANSBLOB structure pb member MUST point to a null-terminated Unicode    |
	//	|                 | string in little-endian format.                                                  |
	//	+-----------------+----------------------------------------------------------------------------------+
	//
	// This method instructs the CA to add, modify, or disable an extension that is associated
	// with a previously submitted certificate request that is in a pending state, as specified
	// in [MS-WCCE] section 3.2.1.4.2.1.3. If the certificate request does not contain an
	// extension with the name specified in pwszExtensionName, then the extension is added
	// to the certificate request. If the request already contains an extension of that
	// name, then the extension specified in the SetExtension call will replace the old
	// one, hence modifying the contents. To disable an extension, a value of 2 can be specified
	// in dwFlags parameter, described above, when calling SetExtension.
	Value *csra.CertTransportBlob `idl:"name:pctbValue;pointer:ref" json:"value"`
}

func (o *SetExtensionRequest) xxx_ToOp(ctx context.Context) *xxx_SetExtensionOperation {
	if o == nil {
		return &xxx_SetExtensionOperation{}
	}
	return &xxx_SetExtensionOperation{
		This:          o.This,
		Authority:     o.Authority,
		RequestID:     o.RequestID,
		ExtensionName: o.ExtensionName,
		Type:          o.Type,
		Flags:         o.Flags,
		Value:         o.Value,
	}
}

func (o *SetExtensionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetExtensionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.RequestID = op.RequestID
	o.ExtensionName = op.ExtensionName
	o.Type = op.Type
	o.Flags = op.Flags
	o.Value = op.Value
}
func (o *SetExtensionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetExtensionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExtensionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetExtensionResponse structure represents the SetExtension operation response
type SetExtensionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetExtension return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetExtensionResponse) xxx_ToOp(ctx context.Context) *xxx_SetExtensionOperation {
	if o == nil {
		return &xxx_SetExtensionOperation{}
	}
	return &xxx_SetExtensionOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetExtensionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetExtensionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetExtensionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetExtensionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExtensionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAttributesOperation structure represents the SetAttributes operation
type xxx_SetAttributesOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority  string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	RequestID  uint32         `idl:"name:dwRequestId" json:"request_id"`
	Attributes string         `idl:"name:pwszAttributes;string;pointer:unique" json:"attributes"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAttributesOperation) OpNum() int { return 4 }

func (o *xxx_SetAttributesOperation) OpName() string { return "/ICertAdminD/v0/SetAttributes" }

func (o *xxx_SetAttributesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAttributesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	// pwszAttributes {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Attributes != "" {
			_ptr_pwszAttributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Attributes); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Attributes, _ptr_pwszAttributes); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAttributesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	// pwszAttributes {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAttributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Attributes); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAttributes := func(ptr interface{}) { o.Attributes = *ptr.(*string) }
		if err := w.ReadPointer(&o.Attributes, _s_pwszAttributes, _ptr_pwszAttributes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAttributesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAttributesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAttributesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetAttributesRequest structure represents the SetAttributes operation request
type SetAttributesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in ICertAdminD::SetExtension (section
	// 3.1.4.1.1).
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// dwRequestId: A 32-bit nonzero unsigned integer value that specifies the ID of the
	// certificate request.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
	// pwszAttributes: A null-terminated Unicode string. The value of the string MUST have
	// the same format as specified in [MS-WCCE] section 3.2.1.4.2.1.2.
	//
	// This method instructs the CA to add or modify a name-value pair that is associated
	// with a previously submitted certificate request that is in a pending state. Information
	// about a pending certificate request is specified in section 3.1.1.1.1.
	Attributes string `idl:"name:pwszAttributes;string;pointer:unique" json:"attributes"`
}

func (o *SetAttributesRequest) xxx_ToOp(ctx context.Context) *xxx_SetAttributesOperation {
	if o == nil {
		return &xxx_SetAttributesOperation{}
	}
	return &xxx_SetAttributesOperation{
		This:       o.This,
		Authority:  o.Authority,
		RequestID:  o.RequestID,
		Attributes: o.Attributes,
	}
}

func (o *SetAttributesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAttributesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.RequestID = op.RequestID
	o.Attributes = op.Attributes
}
func (o *SetAttributesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetAttributesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAttributesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAttributesResponse structure represents the SetAttributes operation response
type SetAttributesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetAttributes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAttributesResponse) xxx_ToOp(ctx context.Context) *xxx_SetAttributesOperation {
	if o == nil {
		return &xxx_SetAttributesOperation{}
	}
	return &xxx_SetAttributesOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetAttributesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAttributesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAttributesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetAttributesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAttributesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResubmitRequestOperation structure represents the ResubmitRequest operation
type xxx_ResubmitRequestOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority   string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	RequestID   uint32         `idl:"name:dwRequestId" json:"request_id"`
	Disposition uint32         `idl:"name:pdwDisposition" json:"disposition"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResubmitRequestOperation) OpNum() int { return 5 }

func (o *xxx_ResubmitRequestOperation) OpName() string { return "/ICertAdminD/v0/ResubmitRequest" }

func (o *xxx_ResubmitRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResubmitRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResubmitRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResubmitRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResubmitRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwDisposition {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Disposition); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResubmitRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwDisposition {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Disposition); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResubmitRequestRequest structure represents the ResubmitRequest operation request
type ResubmitRequestRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in ICertAdminD::SetExtension (section
	// 3.1.4.1.1).
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// dwRequestId: A 32-bit nonzero unsigned integer value that specifies the ID of the
	// certificate request.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
}

func (o *ResubmitRequestRequest) xxx_ToOp(ctx context.Context) *xxx_ResubmitRequestOperation {
	if o == nil {
		return &xxx_ResubmitRequestOperation{}
	}
	return &xxx_ResubmitRequestOperation{
		This:      o.This,
		Authority: o.Authority,
		RequestID: o.RequestID,
	}
}

func (o *ResubmitRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_ResubmitRequestOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.RequestID = op.RequestID
}
func (o *ResubmitRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ResubmitRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResubmitRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResubmitRequestResponse structure represents the ResubmitRequest operation response
type ResubmitRequestResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwDisposition: A pointer to an unsigned integer value that receives the disposition
	// status of the certificate (upon resubmission).
	//
	// This method instructs the CA to try again to process a previously submitted certificate
	// request, which is in a pending or denied state.
	Disposition uint32 `idl:"name:pdwDisposition" json:"disposition"`
	// Return: The ResubmitRequest return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResubmitRequestResponse) xxx_ToOp(ctx context.Context) *xxx_ResubmitRequestOperation {
	if o == nil {
		return &xxx_ResubmitRequestOperation{}
	}
	return &xxx_ResubmitRequestOperation{
		That:        o.That,
		Disposition: o.Disposition,
		Return:      o.Return,
	}
}

func (o *ResubmitRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_ResubmitRequestOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Disposition = op.Disposition
	o.Return = op.Return
}
func (o *ResubmitRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ResubmitRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResubmitRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DenyRequestOperation structure represents the DenyRequest operation
type xxx_DenyRequestOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	RequestID uint32         `idl:"name:dwRequestId" json:"request_id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DenyRequestOperation) OpNum() int { return 6 }

func (o *xxx_DenyRequestOperation) OpName() string { return "/ICertAdminD/v0/DenyRequest" }

func (o *xxx_DenyRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DenyRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DenyRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwRequestId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DenyRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DenyRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DenyRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DenyRequestRequest structure represents the DenyRequest operation request
type DenyRequestRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in ICertAdminD::SetExtension (section
	// 3.1.4.1.1).
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// dwRequestId: A 32-bit nonzero unsigned integer value that specifies the ID of the
	// certificate request.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
}

func (o *DenyRequestRequest) xxx_ToOp(ctx context.Context) *xxx_DenyRequestOperation {
	if o == nil {
		return &xxx_DenyRequestOperation{}
	}
	return &xxx_DenyRequestOperation{
		This:      o.This,
		Authority: o.Authority,
		RequestID: o.RequestID,
	}
}

func (o *DenyRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_DenyRequestOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.RequestID = op.RequestID
}
func (o *DenyRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DenyRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DenyRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DenyRequestResponse structure represents the DenyRequest operation response
type DenyRequestResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DenyRequest return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DenyRequestResponse) xxx_ToOp(ctx context.Context) *xxx_DenyRequestOperation {
	if o == nil {
		return &xxx_DenyRequestOperation{}
	}
	return &xxx_DenyRequestOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DenyRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_DenyRequestOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DenyRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DenyRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DenyRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsValidCertificateOperation structure represents the IsValidCertificate operation
type xxx_IsValidCertificateOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority        string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	SerialNumber     string         `idl:"name:pSerialNumber;string;pointer:unique" json:"serial_number"`
	RevocationReason int32          `idl:"name:pRevocationReason" json:"revocation_reason"`
	Disposition      int32          `idl:"name:pDisposition" json:"disposition"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsValidCertificateOperation) OpNum() int { return 7 }

func (o *xxx_IsValidCertificateOperation) OpName() string {
	return "/ICertAdminD/v0/IsValidCertificate"
}

func (o *xxx_IsValidCertificateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsValidCertificateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pSerialNumber {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.SerialNumber != "" {
			_ptr_pSerialNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SerialNumber); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SerialNumber, _ptr_pSerialNumber); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsValidCertificateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pSerialNumber {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pSerialNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SerialNumber); err != nil {
				return err
			}
			return nil
		})
		_s_pSerialNumber := func(ptr interface{}) { o.SerialNumber = *ptr.(*string) }
		if err := w.ReadPointer(&o.SerialNumber, _s_pSerialNumber, _ptr_pSerialNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsValidCertificateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsValidCertificateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pRevocationReason {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.RevocationReason); err != nil {
			return err
		}
	}
	// pDisposition {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Disposition); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsValidCertificateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pRevocationReason {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.RevocationReason); err != nil {
			return err
		}
	}
	// pDisposition {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Disposition); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsValidCertificateRequest structure represents the IsValidCertificate operation request
type IsValidCertificateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in ICertAdminD::SetExtension (section
	// 3.1.4.1.1).
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// pSerialNumber: A null-terminated Unicode string specifying a serial number that identifies
	// the certificate to be reviewed. The string MUST specify the serial number as an even
	// number of hexadecimal digits. If necessary, a zero can be prefixed to the number
	// to produce an even number of digits. The string MUST NOT contain more than one leading
	// zero. Information about the serial number is as specified in [RFC3280] section 4.1.2.2.
	SerialNumber string `idl:"name:pSerialNumber;string;pointer:unique" json:"serial_number"`
}

func (o *IsValidCertificateRequest) xxx_ToOp(ctx context.Context) *xxx_IsValidCertificateOperation {
	if o == nil {
		return &xxx_IsValidCertificateOperation{}
	}
	return &xxx_IsValidCertificateOperation{
		This:         o.This,
		Authority:    o.Authority,
		SerialNumber: o.SerialNumber,
	}
}

func (o *IsValidCertificateRequest) xxx_FromOp(ctx context.Context, op *xxx_IsValidCertificateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.SerialNumber = op.SerialNumber
}
func (o *IsValidCertificateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsValidCertificateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsValidCertificateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsValidCertificateResponse structure represents the IsValidCertificate operation response
type IsValidCertificateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pRevocationReason: A pointer to a LONG value that receives the revocation reason
	// code. The revocation reason code MUST be one of the following values that are defined
	// for CRLReason, as specified in [RFC3280] section 5.3.1.
	//
	//	+-------+----------------------+
	//	|       |                      |
	//	| VALUE |       MEANING        |
	//	|       |                      |
	//	+-------+----------------------+
	//	+-------+----------------------+
	//	|     0 | unspecified          |
	//	+-------+----------------------+
	//	|     1 | keyCompromise        |
	//	+-------+----------------------+
	//	|     2 | cACompromise         |
	//	+-------+----------------------+
	//	|     3 | affiliationChanged   |
	//	+-------+----------------------+
	//	|     4 | superseded           |
	//	+-------+----------------------+
	//	|     5 | cessationOfOperation |
	//	+-------+----------------------+
	//	|     6 | certificateHold      |
	//	+-------+----------------------+
	RevocationReason int32 `idl:"name:pRevocationReason" json:"revocation_reason"`
	// pDisposition: A pointer to a LONG that receives the disposition status of the request.
	// This parameter MUST be one of the following values.
	//
	//	+------------+-----------------------------------+
	//	|            |                                   |
	//	|   VALUE    |              MEANING              |
	//	|            |                                   |
	//	+------------+-----------------------------------+
	//	+------------+-----------------------------------+
	//	| 0x00000002 | The certificate has been revoked. |
	//	+------------+-----------------------------------+
	//	| 0x00000003 | The certificate is still valid.   |
	//	+------------+-----------------------------------+
	//	| 0x00000004 | The certificate was never issued. |
	//	+------------+-----------------------------------+
	Disposition int32 `idl:"name:pDisposition" json:"disposition"`
	// Return: The IsValidCertificate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsValidCertificateResponse) xxx_ToOp(ctx context.Context) *xxx_IsValidCertificateOperation {
	if o == nil {
		return &xxx_IsValidCertificateOperation{}
	}
	return &xxx_IsValidCertificateOperation{
		That:             o.That,
		RevocationReason: o.RevocationReason,
		Disposition:      o.Disposition,
		Return:           o.Return,
	}
}

func (o *IsValidCertificateResponse) xxx_FromOp(ctx context.Context, op *xxx_IsValidCertificateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RevocationReason = op.RevocationReason
	o.Disposition = op.Disposition
	o.Return = op.Return
}
func (o *IsValidCertificateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsValidCertificateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsValidCertificateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PublishCRLOperation structure represents the PublishCRL operation
type xxx_PublishCRLOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	FileTime  *dtyp.Filetime `idl:"name:FileTime" json:"file_time"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PublishCRLOperation) OpNum() int { return 8 }

func (o *xxx_PublishCRLOperation) OpName() string { return "/ICertAdminD/v0/PublishCRL" }

func (o *xxx_PublishCRLOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PublishCRLOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// FileTime {in} (1:{alias=FILETIME}(struct))
	{
		if o.FileTime != nil {
			if err := o.FileTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_PublishCRLOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// FileTime {in} (1:{alias=FILETIME}(struct))
	{
		if o.FileTime == nil {
			o.FileTime = &dtyp.Filetime{}
		}
		if err := o.FileTime.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PublishCRLOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PublishCRLOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PublishCRLOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PublishCRLRequest structure represents the PublishCRL operation request
type PublishCRLRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in ICertAdminD::SetExtension (section
	// 3.1.4.1.1).
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// FileTime: Contains a 64-bit value that represents the number of 100-nanosecond intervals
	// since January 1, 1601, according to Coordinated Universal Time (UTC). This is used
	// to calculate the nextUpdate value of the CRL as specified in [RFC3280] section 5
	// in UTC-Greenwich Mean Time.
	FileTime *dtyp.Filetime `idl:"name:FileTime" json:"file_time"`
}

func (o *PublishCRLRequest) xxx_ToOp(ctx context.Context) *xxx_PublishCRLOperation {
	if o == nil {
		return &xxx_PublishCRLOperation{}
	}
	return &xxx_PublishCRLOperation{
		This:      o.This,
		Authority: o.Authority,
		FileTime:  o.FileTime,
	}
}

func (o *PublishCRLRequest) xxx_FromOp(ctx context.Context, op *xxx_PublishCRLOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.FileTime = op.FileTime
}
func (o *PublishCRLRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PublishCRLRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PublishCRLOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PublishCRLResponse structure represents the PublishCRL operation response
type PublishCRLResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PublishCRL return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PublishCRLResponse) xxx_ToOp(ctx context.Context) *xxx_PublishCRLOperation {
	if o == nil {
		return &xxx_PublishCRLOperation{}
	}
	return &xxx_PublishCRLOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PublishCRLResponse) xxx_FromOp(ctx context.Context, op *xxx_PublishCRLOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PublishCRLResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PublishCRLResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PublishCRLOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCRLOperation structure represents the GetCRL operation
type xxx_GetCRLOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	CRL       *csra.CertTransportBlob `idl:"name:pctbCRL;pointer:ref" json:"crl"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCRLOperation) OpNum() int { return 9 }

func (o *xxx_GetCRLOperation) OpName() string { return "/ICertAdminD/v0/GetCRL" }

func (o *xxx_GetCRLOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCRLOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCRLOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCRLOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCRLOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pctbCRL {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.CRL != nil {
			if err := o.CRL.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csra.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCRLOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pctbCRL {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.CRL == nil {
			o.CRL = &csra.CertTransportBlob{}
		}
		if err := o.CRL.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetCRLRequest structure represents the GetCRL operation request
type GetCRLRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *GetCRLRequest) xxx_ToOp(ctx context.Context) *xxx_GetCRLOperation {
	if o == nil {
		return &xxx_GetCRLOperation{}
	}
	return &xxx_GetCRLOperation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *GetCRLRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCRLOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *GetCRLRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCRLRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCRLOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCRLResponse structure represents the GetCRL operation response
type GetCRLResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pctbCRL: If the function succeeds, this method MUST return a CERTTRANSBLOB structure
	// that contains the ASN.1 DER (as specified in [X660] and [X690]) encoded CRL (CRLRawCRL)
	// for the CA server's current signing certificate.
	//
	// The GetCRL method instructs the CA to return the recent base CRL, which is signed
	// with the current CA key to the caller. If a CRL cannot be found, the CA MUST return
	// ERROR_FILE_NOT_FOUND, as specified in [MS-ERREF].<35>
	CRL *csra.CertTransportBlob `idl:"name:pctbCRL;pointer:ref" json:"crl"`
	// Return: The GetCRL return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCRLResponse) xxx_ToOp(ctx context.Context) *xxx_GetCRLOperation {
	if o == nil {
		return &xxx_GetCRLOperation{}
	}
	return &xxx_GetCRLOperation{
		That:   o.That,
		CRL:    o.CRL,
		Return: o.Return,
	}
}

func (o *GetCRLResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCRLOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CRL = op.CRL
	o.Return = op.Return
}
func (o *GetCRLResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCRLResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCRLOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RevokeCertificateOperation structure represents the RevokeCertificate operation
type xxx_RevokeCertificateOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority    string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	SerialNumber string         `idl:"name:pwszSerialNumber;string;pointer:unique" json:"serial_number"`
	Reason       uint32         `idl:"name:Reason" json:"reason"`
	FileTime     *dtyp.Filetime `idl:"name:FileTime" json:"file_time"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RevokeCertificateOperation) OpNum() int { return 10 }

func (o *xxx_RevokeCertificateOperation) OpName() string { return "/ICertAdminD/v0/RevokeCertificate" }

func (o *xxx_RevokeCertificateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeCertificateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszSerialNumber {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.SerialNumber != "" {
			_ptr_pwszSerialNumber := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SerialNumber); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SerialNumber, _ptr_pwszSerialNumber); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Reason {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Reason); err != nil {
			return err
		}
	}
	// FileTime {in} (1:{alias=FILETIME}(struct))
	{
		if o.FileTime != nil {
			if err := o.FileTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RevokeCertificateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszSerialNumber {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszSerialNumber := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SerialNumber); err != nil {
				return err
			}
			return nil
		})
		_s_pwszSerialNumber := func(ptr interface{}) { o.SerialNumber = *ptr.(*string) }
		if err := w.ReadPointer(&o.SerialNumber, _s_pwszSerialNumber, _ptr_pwszSerialNumber); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Reason {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Reason); err != nil {
			return err
		}
	}
	// FileTime {in} (1:{alias=FILETIME}(struct))
	{
		if o.FileTime == nil {
			o.FileTime = &dtyp.Filetime{}
		}
		if err := o.FileTime.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeCertificateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeCertificateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RevokeCertificateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RevokeCertificateRequest structure represents the RevokeCertificate operation request
type RevokeCertificateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in ICertAdminD::SetExtension (section
	// 3.1.4.1.1).
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// pwszSerialNumber: A null-terminated Unicode string that specifies a serial number
	// that identifies the certificate to be revoked. The string MUST specify the serial
	// number as plain hexadecimal digits (no leading 0x) as specified in [RFC3280] section
	// 4.1.2.2.<36>
	SerialNumber string `idl:"name:pwszSerialNumber;string;pointer:unique" json:"serial_number"`
	// Reason: An unsigned integer value that specifies the revocation reason code. The
	// revocation reason code MUST be either one of the values listed in the following table
	// (and specified in [RFC3280] section 5.3.1), or one of the following values: 0xfffffffd,
	// 0xfffffffe, or 0xffffffff.
	//
	//	+------------+--------------------------------------------------------------------+
	//	|            |                                                                    |
	//	|   VALUE    |                              MEANING                               |
	//	|            |                                                                    |
	//	+------------+--------------------------------------------------------------------+
	//	+------------+--------------------------------------------------------------------+
	//	|          0 | unspecified                                                        |
	//	+------------+--------------------------------------------------------------------+
	//	|          1 | keyCompromise                                                      |
	//	+------------+--------------------------------------------------------------------+
	//	|          2 | cACompromise                                                       |
	//	+------------+--------------------------------------------------------------------+
	//	|          3 | affiliationChanged                                                 |
	//	+------------+--------------------------------------------------------------------+
	//	|          4 | superseded                                                         |
	//	+------------+--------------------------------------------------------------------+
	//	|          5 | cessationOfOperation                                               |
	//	+------------+--------------------------------------------------------------------+
	//	|          6 | certificateHold                                                    |
	//	+------------+--------------------------------------------------------------------+
	//	|          8 | removeFromCRL                                                      |
	//	+------------+--------------------------------------------------------------------+
	//	| 0xfffffffd | See processing rules, beginning with rule 2.                       |
	//	+------------+--------------------------------------------------------------------+
	//	| 0xfffffffe | See processing rules, beginning with rule 3.                       |
	//	+------------+--------------------------------------------------------------------+
	//	| 0xffffffff | Released from hold. (See processing rules, beginning with rule 4.) |
	//	+------------+--------------------------------------------------------------------+
	Reason uint32 `idl:"name:Reason" json:"reason"`
	// FileTime: Contains a 64-bit value that represents the number of 100-nanosecond intervals
	// since January 1, 1601 (UTC). This value specifies the date, according to Greenwich
	// Mean Time, when the certificate became invalid. The FileTime corresponds to the Request_Revocation_Date
	// that is defined in section 3.1.1.1.
	FileTime *dtyp.Filetime `idl:"name:FileTime" json:"file_time"`
}

func (o *RevokeCertificateRequest) xxx_ToOp(ctx context.Context) *xxx_RevokeCertificateOperation {
	if o == nil {
		return &xxx_RevokeCertificateOperation{}
	}
	return &xxx_RevokeCertificateOperation{
		This:         o.This,
		Authority:    o.Authority,
		SerialNumber: o.SerialNumber,
		Reason:       o.Reason,
		FileTime:     o.FileTime,
	}
}

func (o *RevokeCertificateRequest) xxx_FromOp(ctx context.Context, op *xxx_RevokeCertificateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.SerialNumber = op.SerialNumber
	o.Reason = op.Reason
	o.FileTime = op.FileTime
}
func (o *RevokeCertificateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RevokeCertificateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RevokeCertificateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RevokeCertificateResponse structure represents the RevokeCertificate operation response
type RevokeCertificateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RevokeCertificate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RevokeCertificateResponse) xxx_ToOp(ctx context.Context) *xxx_RevokeCertificateOperation {
	if o == nil {
		return &xxx_RevokeCertificateOperation{}
	}
	return &xxx_RevokeCertificateOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RevokeCertificateResponse) xxx_FromOp(ctx context.Context, op *xxx_RevokeCertificateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RevokeCertificateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RevokeCertificateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RevokeCertificateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumViewColumnOperation structure represents the EnumViewColumn operation
type xxx_EnumViewColumnOperation struct {
	This           *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority      string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Column         uint32                  `idl:"name:iColumn" json:"column"`
	ColumnCount    uint32                  `idl:"name:cColumn" json:"column_count"`
	ColumnOutCount uint32                  `idl:"name:pcColumnOut" json:"column_out_count"`
	ColumnInfo     *csra.CertTransportBlob `idl:"name:pctbColumnInfo;pointer:ref" json:"column_info"`
	Return         int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumViewColumnOperation) OpNum() int { return 11 }

func (o *xxx_EnumViewColumnOperation) OpName() string { return "/ICertAdminD/v0/EnumViewColumn" }

func (o *xxx_EnumViewColumnOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewColumnOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// iColumn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Column); err != nil {
			return err
		}
	}
	// cColumn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ColumnCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewColumnOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// iColumn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Column); err != nil {
			return err
		}
	}
	// cColumn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ColumnCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewColumnOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewColumnOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pcColumnOut {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ColumnOutCount); err != nil {
			return err
		}
	}
	// pctbColumnInfo {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.ColumnInfo != nil {
			if err := o.ColumnInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csra.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewColumnOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcColumnOut {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ColumnOutCount); err != nil {
			return err
		}
	}
	// pctbColumnInfo {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.ColumnInfo == nil {
			o.ColumnInfo = &csra.CertTransportBlob{}
		}
		if err := o.ColumnInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumViewColumnRequest structure represents the EnumViewColumn operation request
type EnumViewColumnRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the definition of the pwszAuthority parameter in ICertAdminD::SetExtension
	// (section 3.1.4.1.1).
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// iColumn: An unsigned integer that specifies the identifier of the column with which
	// to begin the enumeration. Valid values are from 0 to one less than the maximum number
	// of columns for the Request table.
	Column uint32 `idl:"name:iColumn" json:"column"`
	// cColumn: An unsigned integer that specifies the requested number of columns to return.
	ColumnCount uint32 `idl:"name:cColumn" json:"column_count"`
}

func (o *EnumViewColumnRequest) xxx_ToOp(ctx context.Context) *xxx_EnumViewColumnOperation {
	if o == nil {
		return &xxx_EnumViewColumnOperation{}
	}
	return &xxx_EnumViewColumnOperation{
		This:        o.This,
		Authority:   o.Authority,
		Column:      o.Column,
		ColumnCount: o.ColumnCount,
	}
}

func (o *EnumViewColumnRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumViewColumnOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.Column = op.Column
	o.ColumnCount = op.ColumnCount
}
func (o *EnumViewColumnRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumViewColumnRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumViewColumnOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumViewColumnResponse structure represents the EnumViewColumn operation response
type EnumViewColumnResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	ColumnOutCount uint32         `idl:"name:pcColumnOut" json:"column_out_count"`
	// pctbColumnInfo: A pointer to a CERTTRANSBLOB structure. Upon return, the pb member
	// of this structure points to an array of the marshaled CERTTRANSDBCOLUMN structures
	// as described in section 2.2.1.7.
	//
	// The EnumViewColumn method returns information about the columns that are associated
	// with the Request table to the client. The processing rules for this method are the
	// same as for the EnumViewColumnTable method with the iTable parameter set to 0x00000000.
	ColumnInfo *csra.CertTransportBlob `idl:"name:pctbColumnInfo;pointer:ref" json:"column_info"`
	// Return: The EnumViewColumn return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumViewColumnResponse) xxx_ToOp(ctx context.Context) *xxx_EnumViewColumnOperation {
	if o == nil {
		return &xxx_EnumViewColumnOperation{}
	}
	return &xxx_EnumViewColumnOperation{
		That:           o.That,
		ColumnOutCount: o.ColumnOutCount,
		ColumnInfo:     o.ColumnInfo,
		Return:         o.Return,
	}
}

func (o *EnumViewColumnResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumViewColumnOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ColumnOutCount = op.ColumnOutCount
	o.ColumnInfo = op.ColumnInfo
	o.Return = op.Return
}
func (o *EnumViewColumnResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumViewColumnResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumViewColumnOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetViewDefaultColumnSetOperation structure represents the GetViewDefaultColumnSet operation
type xxx_GetViewDefaultColumnSetOperation struct {
	This             *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority        string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	ColumnSetDefault uint32                  `idl:"name:iColumnSetDefault" json:"column_set_default"`
	ColumnOutCount   uint32                  `idl:"name:pcColumnOut" json:"column_out_count"`
	ColumnInfo       *csra.CertTransportBlob `idl:"name:pctbColumnInfo;pointer:ref" json:"column_info"`
	Return           int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetViewDefaultColumnSetOperation) OpNum() int { return 12 }

func (o *xxx_GetViewDefaultColumnSetOperation) OpName() string {
	return "/ICertAdminD/v0/GetViewDefaultColumnSet"
}

func (o *xxx_GetViewDefaultColumnSetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetViewDefaultColumnSetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// iColumnSetDefault {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ColumnSetDefault); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetViewDefaultColumnSetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// iColumnSetDefault {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ColumnSetDefault); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetViewDefaultColumnSetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetViewDefaultColumnSetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pcColumnOut {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ColumnOutCount); err != nil {
			return err
		}
	}
	// pctbColumnInfo {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.ColumnInfo != nil {
			if err := o.ColumnInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csra.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetViewDefaultColumnSetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcColumnOut {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ColumnOutCount); err != nil {
			return err
		}
	}
	// pctbColumnInfo {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.ColumnInfo == nil {
			o.ColumnInfo = &csra.CertTransportBlob{}
		}
		if err := o.ColumnInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetViewDefaultColumnSetRequest structure represents the GetViewDefaultColumnSet operation request
type GetViewDefaultColumnSetRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the definition of the pwszAuthority parameter in ICertAdminD::SetExtension
	// (section 3.1.4.1.1).
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// iColumnSetDefault: An unsigned integer value that specifies the requested default
	// column set to get. The value MUST be one of the values in the following table. If
	// a value other than one of the listed values is used, the error E_INVALIDARG is returned.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0xFFFFFFFF | The caller attempts to retrieve the list of column identifiers that are useful   |
	//	|            | for viewing pending requests.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0xFFFFFFFE | The caller attempts to retrieve the list of column identifiers that are useful   |
	//	|            | for viewing issued certificates.                                                 |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0xFFFFFFFD | The caller attempts to retrieve the list of column identifiers that are useful   |
	//	|            | for viewing failed requests.                                                     |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0xFFFFFFFC | The caller attempts to retrieve the list of column identifiers that are useful   |
	//	|            | for viewing extensions.                                                          |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0xFFFFFFFB | The caller attempts to retrieve the list of column identifiers that are useful   |
	//	|            | for viewing attributes.                                                          |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0xFFFFFFFA | The caller attempts to retrieve the list of column identifiers that are useful   |
	//	|            | for viewing CRLs.                                                                |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0xFFFFFFF9 | The caller attempts to retrieve the list of column identifiers that are useful   |
	//	|            | for viewing revoked certificates.                                                |
	//	+------------+----------------------------------------------------------------------------------+
	ColumnSetDefault uint32 `idl:"name:iColumnSetDefault" json:"column_set_default"`
}

func (o *GetViewDefaultColumnSetRequest) xxx_ToOp(ctx context.Context) *xxx_GetViewDefaultColumnSetOperation {
	if o == nil {
		return &xxx_GetViewDefaultColumnSetOperation{}
	}
	return &xxx_GetViewDefaultColumnSetOperation{
		This:             o.This,
		Authority:        o.Authority,
		ColumnSetDefault: o.ColumnSetDefault,
	}
}

func (o *GetViewDefaultColumnSetRequest) xxx_FromOp(ctx context.Context, op *xxx_GetViewDefaultColumnSetOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.ColumnSetDefault = op.ColumnSetDefault
}
func (o *GetViewDefaultColumnSetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetViewDefaultColumnSetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetViewDefaultColumnSetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetViewDefaultColumnSetResponse structure represents the GetViewDefaultColumnSet operation response
type GetViewDefaultColumnSetResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	ColumnOutCount uint32         `idl:"name:pcColumnOut" json:"column_out_count"`
	// pctbColumnInfo: A pointer to a CERTTRANSBLOB structure. Its cb member MUST contain
	// the length, in bytes, of the array that is referenced by the pb member. The pb member
	// MUST point to an array of DWORDs, where each DWORD value represents the identifier
	// for a column. Each DWORD in the array is marshaled by using little-endian format.
	ColumnInfo *csra.CertTransportBlob `idl:"name:pctbColumnInfo;pointer:ref" json:"column_info"`
	// Return: The GetViewDefaultColumnSet return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetViewDefaultColumnSetResponse) xxx_ToOp(ctx context.Context) *xxx_GetViewDefaultColumnSetOperation {
	if o == nil {
		return &xxx_GetViewDefaultColumnSetOperation{}
	}
	return &xxx_GetViewDefaultColumnSetOperation{
		That:           o.That,
		ColumnOutCount: o.ColumnOutCount,
		ColumnInfo:     o.ColumnInfo,
		Return:         o.Return,
	}
}

func (o *GetViewDefaultColumnSetResponse) xxx_FromOp(ctx context.Context, op *xxx_GetViewDefaultColumnSetOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ColumnOutCount = op.ColumnOutCount
	o.ColumnInfo = op.ColumnInfo
	o.Return = op.Return
}
func (o *GetViewDefaultColumnSetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetViewDefaultColumnSetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetViewDefaultColumnSetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumAttributesOrExtensionsOperation structure represents the EnumAttributesOrExtensions operation
type xxx_EnumAttributesOrExtensionsOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	RowID     uint32                  `idl:"name:RowId" json:"row_id"`
	Flags     uint32                  `idl:"name:Flags" json:"flags"`
	Last      string                  `idl:"name:pwszLast;string;pointer:unique" json:"last"`
	Count     uint32                  `idl:"name:celt" json:"count"`
	Fetched   uint32                  `idl:"name:pceltFetched" json:"fetched"`
	Out       *csra.CertTransportBlob `idl:"name:pctbOut;pointer:ref" json:"out"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumAttributesOrExtensionsOperation) OpNum() int { return 13 }

func (o *xxx_EnumAttributesOrExtensionsOperation) OpName() string {
	return "/ICertAdminD/v0/EnumAttributesOrExtensions"
}

func (o *xxx_EnumAttributesOrExtensionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAttributesOrExtensionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// RowId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RowID); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pwszLast {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Last != "" {
			_ptr_pwszLast := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Last); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Last, _ptr_pwszLast); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// celt {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAttributesOrExtensionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RowId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RowID); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pwszLast {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszLast := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Last); err != nil {
				return err
			}
			return nil
		})
		_s_pwszLast := func(ptr interface{}) { o.Last = *ptr.(*string) }
		if err := w.ReadPointer(&o.Last, _s_pwszLast, _ptr_pwszLast); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// celt {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAttributesOrExtensionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAttributesOrExtensionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pceltFetched {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Fetched); err != nil {
			return err
		}
	}
	// pctbOut {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Out != nil {
			if err := o.Out.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csra.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAttributesOrExtensionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pceltFetched {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Fetched); err != nil {
			return err
		}
	}
	// pctbOut {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Out == nil {
			o.Out = &csra.CertTransportBlob{}
		}
		if err := o.Out.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumAttributesOrExtensionsRequest structure represents the EnumAttributesOrExtensions operation request
type EnumAttributesOrExtensionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the pwszAuthority definition in section ICertAdminD::SetExtension
	// (section 3.1.4.1.1).
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// RowId: An unsigned integer that specifies the RequestID value of the row to retrieve
	// attributes or extensions.
	RowID uint32 `idl:"name:RowId" json:"row_id"`
	// Flags: An unsigned integer value that MUST take either of the following values.
	//
	//	+------------+----------------------+
	//	|            |                      |
	//	|   VALUE    |       MEANING        |
	//	|            |                      |
	//	+------------+----------------------+
	//	+------------+----------------------+
	//	| 0x00000000 | Enumerate attributes |
	//	+------------+----------------------+
	//	| 0x00000001 | Enumerate extensions |
	//	+------------+----------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// pwszLast: A pointer to a null-terminated UNICODE string that specifies the name of
	// the attribute or extension beyond which the data is requested. If the value of Flags
	// is 1, the name MUST be an OID (1) string as specified in [X680].
	Last string `idl:"name:pwszLast;string;pointer:unique" json:"last"`
	// celt: An unsigned integer value that specifies the requested count of attributes
	// (CERTTRANSDBATTRIBUTE) or extensions (CERTTRANSDBEXTENSION) structures to be returned
	// to the client.
	Count uint32 `idl:"name:celt" json:"count"`
}

func (o *EnumAttributesOrExtensionsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumAttributesOrExtensionsOperation {
	if o == nil {
		return &xxx_EnumAttributesOrExtensionsOperation{}
	}
	return &xxx_EnumAttributesOrExtensionsOperation{
		This:      o.This,
		Authority: o.Authority,
		RowID:     o.RowID,
		Flags:     o.Flags,
		Last:      o.Last,
		Count:     o.Count,
	}
}

func (o *EnumAttributesOrExtensionsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumAttributesOrExtensionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.RowID = op.RowID
	o.Flags = op.Flags
	o.Last = op.Last
	o.Count = op.Count
}
func (o *EnumAttributesOrExtensionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumAttributesOrExtensionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumAttributesOrExtensionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumAttributesOrExtensionsResponse structure represents the EnumAttributesOrExtensions operation response
type EnumAttributesOrExtensionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pceltFetched: A pointer to the unsigned integer that receives the actual count of
	// the attributes (CERTTRANSDBATTRIBUTE) or extensions (CERTTRANSDBEXTENSION) structure
	// data returned by the server in the pctbOut parameter.
	Fetched uint32 `idl:"name:pceltFetched" json:"fetched"`
	// pctbOut: A pointer to the CERTTRANSBLOB structure. The data returned is marshaled
	// CERTTRANSDBATTRIBUTE or CERTTRANSDBEXTENSION structure array as described in CERTTRANSDBATTRIBUTE
	// and CERTTRANSDBEXTENSION.
	//
	// The EnumAttributesOrExtensions method obtains information about the attributes or
	// extensions (as specified in [MS-WCCE] section 2.2.2.7) that are associated with a
	// specific request in the Request table.
	Out *csra.CertTransportBlob `idl:"name:pctbOut;pointer:ref" json:"out"`
	// Return: The EnumAttributesOrExtensions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumAttributesOrExtensionsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumAttributesOrExtensionsOperation {
	if o == nil {
		return &xxx_EnumAttributesOrExtensionsOperation{}
	}
	return &xxx_EnumAttributesOrExtensionsOperation{
		That:    o.That,
		Fetched: o.Fetched,
		Out:     o.Out,
		Return:  o.Return,
	}
}

func (o *EnumAttributesOrExtensionsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumAttributesOrExtensionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Fetched = op.Fetched
	o.Out = op.Out
	o.Return = op.Return
}
func (o *EnumAttributesOrExtensionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumAttributesOrExtensionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumAttributesOrExtensionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenViewOperation structure represents the OpenView operation
type xxx_OpenViewOperation struct {
	This                     *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Authority                string                      `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	CertViewRestrictionCount uint32                      `idl:"name:ccvr" json:"cert_view_restriction_count"`
	CertViewRestrictions     []*csra.CertViewRestriction `idl:"name:acvr;size_is:(ccvr)" json:"cert_view_restrictions"`
	ColumnsCountOut          uint32                      `idl:"name:ccolOut" json:"columns_count_out"`
	ColumnsOut               []uint32                    `idl:"name:acolOut;size_is:(ccolOut)" json:"columns_out"`
	ID                       uint32                      `idl:"name:ielt" json:"id"`
	Count                    uint32                      `idl:"name:celt" json:"count"`
	Fetched                  uint32                      `idl:"name:pceltFetched" json:"fetched"`
	ResultRows               *csra.CertTransportBlob     `idl:"name:pctbResultRows;pointer:ref" json:"result_rows"`
	Return                   int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenViewOperation) OpNum() int { return 14 }

func (o *xxx_OpenViewOperation) OpName() string { return "/ICertAdminD/v0/OpenView" }

func (o *xxx_OpenViewOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.CertViewRestrictions != nil && o.CertViewRestrictionCount == 0 {
		o.CertViewRestrictionCount = uint32(len(o.CertViewRestrictions))
	}
	if o.ColumnsOut != nil && o.ColumnsCountOut == 0 {
		o.ColumnsCountOut = uint32(len(o.ColumnsOut))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenViewOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ccvr {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CertViewRestrictionCount); err != nil {
			return err
		}
	}
	// acvr {in} (1:{pointer=ref}*(1))(2:{alias=CERTVIEWRESTRICTION}[dim:0,size_is=ccvr](struct))
	{
		dimSize1 := uint64(o.CertViewRestrictionCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.CertViewRestrictions {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.CertViewRestrictions[i1] != nil {
				if err := o.CertViewRestrictions[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&csra.CertViewRestriction{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.CertViewRestrictions); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&csra.CertViewRestriction{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ccolOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ColumnsCountOut); err != nil {
			return err
		}
	}
	// acolOut {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}[dim:0,size_is=ccolOut](uint32))
	{
		dimSize1 := uint64(o.ColumnsCountOut)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ColumnsOut {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ColumnsOut[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ColumnsOut); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		}
	}
	// ielt {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ID); err != nil {
			return err
		}
	}
	// celt {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenViewOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ccvr {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CertViewRestrictionCount); err != nil {
			return err
		}
	}
	// acvr {in} (1:{pointer=ref}*(1))(2:{alias=CERTVIEWRESTRICTION}[dim:0,size_is=ccvr](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.CertViewRestrictions", sizeInfo[0])
		}
		o.CertViewRestrictions = make([]*csra.CertViewRestriction, sizeInfo[0])
		for i1 := range o.CertViewRestrictions {
			i1 := i1
			if o.CertViewRestrictions[i1] == nil {
				o.CertViewRestrictions[i1] = &csra.CertViewRestriction{}
			}
			if err := o.CertViewRestrictions[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ccolOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ColumnsCountOut); err != nil {
			return err
		}
	}
	// acolOut {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}[dim:0,size_is=ccolOut](uint32))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ColumnsOut", sizeInfo[0])
		}
		o.ColumnsOut = make([]uint32, sizeInfo[0])
		for i1 := range o.ColumnsOut {
			i1 := i1
			if err := w.ReadData(&o.ColumnsOut[i1]); err != nil {
				return err
			}
		}
	}
	// ielt {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ID); err != nil {
			return err
		}
	}
	// celt {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenViewOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenViewOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pceltFetched {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Fetched); err != nil {
			return err
		}
	}
	// pctbResultRows {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.ResultRows != nil {
			if err := o.ResultRows.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csra.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenViewOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pceltFetched {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Fetched); err != nil {
			return err
		}
	}
	// pctbResultRows {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.ResultRows == nil {
			o.ResultRows = &csra.CertTransportBlob{}
		}
		if err := o.ResultRows.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenViewRequest structure represents the OpenView operation request
type OpenViewRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the definition of the pwszAuthority parameter in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// ccvr: An unsigned integer value that specifies the count of a CERTVIEWRESTRICTION
	// structure array that is pointed to by the acvr parameter.
	CertViewRestrictionCount uint32 `idl:"name:ccvr" json:"cert_view_restriction_count"`
	// acvr: A pointer to an array of CERTVIEWRESTRICTION structures. For more information,
	// see section 2.2.1.3.
	CertViewRestrictions []*csra.CertViewRestriction `idl:"name:acvr;size_is:(ccvr)" json:"cert_view_restrictions"`
	// ccolOut: An unsigned integer value that specifies the count of a DWORD array that
	// is pointed to by the acolOut parameter.
	ColumnsCountOut uint32 `idl:"name:ccolOut" json:"columns_count_out"`
	// acolOut: A pointer to an array of DWORDs. Each DWORD value specifies the column identifier
	// for the resultant set of rows.
	ColumnsOut []uint32 `idl:"name:acolOut;size_is:(ccolOut)" json:"columns_out"`
	// ielt: An unsigned integer value that specifies the identifier of the first row to
	// return from the resultant set of rows.
	ID uint32 `idl:"name:ielt" json:"id"`
	// celt: An unsigned integer value that specifies the requested count of the row data
	// to be returned from the resultant set of rows.
	Count uint32 `idl:"name:celt" json:"count"`
}

func (o *OpenViewRequest) xxx_ToOp(ctx context.Context) *xxx_OpenViewOperation {
	if o == nil {
		return &xxx_OpenViewOperation{}
	}
	return &xxx_OpenViewOperation{
		This:                     o.This,
		Authority:                o.Authority,
		CertViewRestrictionCount: o.CertViewRestrictionCount,
		CertViewRestrictions:     o.CertViewRestrictions,
		ColumnsCountOut:          o.ColumnsCountOut,
		ColumnsOut:               o.ColumnsOut,
		ID:                       o.ID,
		Count:                    o.Count,
	}
}

func (o *OpenViewRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenViewOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.CertViewRestrictionCount = op.CertViewRestrictionCount
	o.CertViewRestrictions = op.CertViewRestrictions
	o.ColumnsCountOut = op.ColumnsCountOut
	o.ColumnsOut = op.ColumnsOut
	o.ID = op.ID
	o.Count = op.Count
}
func (o *OpenViewRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenViewRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenViewOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenViewResponse structure represents the OpenView operation response
type OpenViewResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pceltFetched: A pointer to an unsigned integer value that receives the actual count
	// of row data that is returned by the server in the pctbResultRows parameter.
	Fetched uint32 `idl:"name:pceltFetched" json:"fetched"`
	// pctbResultRows: A pointer to a CERTTRANSBLOB structure. The pb byte array of the
	// CERTTRANSBLOB structure MUST contain (on successful return) an array of n marshaled
	// CERTTRANSDBRESULTROW structures (section 2.2.3), where n is the value returned in
	// pceltFetched. Each CERTTRANSDBRESULTROW contains one or more CERTTRANSDBRESULTCOLUMN
	// structures (section 2.2.1.10).
	//
	// The OpenView method opens a view into the database and returns a set of resultant
	// row data.
	ResultRows *csra.CertTransportBlob `idl:"name:pctbResultRows;pointer:ref" json:"result_rows"`
	// Return: The OpenView return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenViewResponse) xxx_ToOp(ctx context.Context) *xxx_OpenViewOperation {
	if o == nil {
		return &xxx_OpenViewOperation{}
	}
	return &xxx_OpenViewOperation{
		That:       o.That,
		Fetched:    o.Fetched,
		ResultRows: o.ResultRows,
		Return:     o.Return,
	}
}

func (o *OpenViewResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenViewOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Fetched = op.Fetched
	o.ResultRows = op.ResultRows
	o.Return = op.Return
}
func (o *OpenViewResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenViewResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenViewOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumViewOperation structure represents the EnumView operation
type xxx_EnumViewOperation struct {
	This       *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority  string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	ID         uint32                  `idl:"name:ielt" json:"id"`
	Count      uint32                  `idl:"name:celt" json:"count"`
	Fetched    uint32                  `idl:"name:pceltFetched" json:"fetched"`
	ResultRows *csra.CertTransportBlob `idl:"name:pctbResultRows;pointer:ref" json:"result_rows"`
	Return     int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumViewOperation) OpNum() int { return 15 }

func (o *xxx_EnumViewOperation) OpName() string { return "/ICertAdminD/v0/EnumView" }

func (o *xxx_EnumViewOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ielt {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ID); err != nil {
			return err
		}
	}
	// celt {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ielt {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ID); err != nil {
			return err
		}
	}
	// celt {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pceltFetched {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Fetched); err != nil {
			return err
		}
	}
	// pctbResultRows {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.ResultRows != nil {
			if err := o.ResultRows.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csra.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pceltFetched {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Fetched); err != nil {
			return err
		}
	}
	// pctbResultRows {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.ResultRows == nil {
			o.ResultRows = &csra.CertTransportBlob{}
		}
		if err := o.ResultRows.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumViewRequest structure represents the EnumView operation request
type EnumViewRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the definition of the pwszAuthority parameter in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// ielt: An unsigned integer value that specifies the identifier of the first row to
	// return from the resultant set of rows.
	ID uint32 `idl:"name:ielt" json:"id"`
	// celt: An unsigned integer value that specifies the requested count of the row data
	// to be returned from the resultant set of rows.
	Count uint32 `idl:"name:celt" json:"count"`
}

func (o *EnumViewRequest) xxx_ToOp(ctx context.Context) *xxx_EnumViewOperation {
	if o == nil {
		return &xxx_EnumViewOperation{}
	}
	return &xxx_EnumViewOperation{
		This:      o.This,
		Authority: o.Authority,
		ID:        o.ID,
		Count:     o.Count,
	}
}

func (o *EnumViewRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumViewOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.ID = op.ID
	o.Count = op.Count
}
func (o *EnumViewRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumViewRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumViewOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumViewResponse structure represents the EnumView operation response
type EnumViewResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pceltFetched: A pointer to an unsigned integer value that receives the actual count
	// of row data that is returned by the server in the pctbResultRows parameter.
	Fetched uint32 `idl:"name:pceltFetched" json:"fetched"`
	// pctbResultRows: A pointer to a CERTTRANSBLOB structure. The pb byte array of the
	// CERTTRANSBLOB structure MUST contain (on successful return) an array of n marshaled
	// CERTTRANSDBRESULTROW structures (section 2.2.3), where n is the value returned in
	// pceltFetched. Each CERTTRANSDBRESULTROW contains one or more CERTTRANSDBRESULTCOLUMN
	// structures (section 2.2.1.10). In addition, an extra CERTTRANSDBRESULTROW structure
	// is included in the array when the server encounters the end of the enumeration, as
	// described in the following rules.
	ResultRows *csra.CertTransportBlob `idl:"name:pctbResultRows;pointer:ref" json:"result_rows"`
	// Return: The EnumView return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumViewResponse) xxx_ToOp(ctx context.Context) *xxx_EnumViewOperation {
	if o == nil {
		return &xxx_EnumViewOperation{}
	}
	return &xxx_EnumViewOperation{
		That:       o.That,
		Fetched:    o.Fetched,
		ResultRows: o.ResultRows,
		Return:     o.Return,
	}
}

func (o *EnumViewResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumViewOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Fetched = op.Fetched
	o.ResultRows = op.ResultRows
	o.Return = op.Return
}
func (o *EnumViewResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumViewResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumViewOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseViewOperation structure represents the CloseView operation
type xxx_CloseViewOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseViewOperation) OpNum() int { return 16 }

func (o *xxx_CloseViewOperation) OpName() string { return "/ICertAdminD/v0/CloseView" }

func (o *xxx_CloseViewOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseViewOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseViewOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseViewOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseViewOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseViewOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseViewRequest structure represents the CloseView operation request
type CloseViewRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *CloseViewRequest) xxx_ToOp(ctx context.Context) *xxx_CloseViewOperation {
	if o == nil {
		return &xxx_CloseViewOperation{}
	}
	return &xxx_CloseViewOperation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *CloseViewRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseViewOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *CloseViewRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseViewRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseViewOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseViewResponse structure represents the CloseView operation response
type CloseViewResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CloseView return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseViewResponse) xxx_ToOp(ctx context.Context) *xxx_CloseViewOperation {
	if o == nil {
		return &xxx_CloseViewOperation{}
	}
	return &xxx_CloseViewOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CloseViewResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseViewOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CloseViewResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseViewResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseViewOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerControlOperation structure represents the ServerControl operation
type xxx_ServerControlOperation struct {
	This         *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority    string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	ControlFlags uint32                  `idl:"name:dwControlFlags" json:"control_flags"`
	Out          *csra.CertTransportBlob `idl:"name:pctbOut;pointer:ref" json:"out"`
	Return       int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerControlOperation) OpNum() int { return 17 }

func (o *xxx_ServerControlOperation) OpName() string { return "/ICertAdminD/v0/ServerControl" }

func (o *xxx_ServerControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwControlFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwControlFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pctbOut {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Out != nil {
			if err := o.Out.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csra.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pctbOut {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Out == nil {
			o.Out = &csra.CertTransportBlob{}
		}
		if err := o.Out.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerControlRequest structure represents the ServerControl operation request
type ServerControlRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// dwControlFlags: An unsigned integer value that specifies the control to be sent to
	// the certificate server. It MUST take one of the following values.
	//
	//	+-------------+--------------------------------------------------------------------+
	//	|             |                                                                    |
	//	|    VALUE    |                              MEANING                               |
	//	|             |                                                                    |
	//	+-------------+--------------------------------------------------------------------+
	//	+-------------+--------------------------------------------------------------------+
	//	| 0x000000001 | Request unregister for DCOM interfaces for the certificate server. |
	//	+-------------+--------------------------------------------------------------------+
	//	| 0x000000002 | Not currently used.                                                |
	//	+-------------+--------------------------------------------------------------------+
	//	| 0x000000003 | Not currently used.                                                |
	//	+-------------+--------------------------------------------------------------------+
	ControlFlags uint32 `idl:"name:dwControlFlags" json:"control_flags"`
}

func (o *ServerControlRequest) xxx_ToOp(ctx context.Context) *xxx_ServerControlOperation {
	if o == nil {
		return &xxx_ServerControlOperation{}
	}
	return &xxx_ServerControlOperation{
		This:         o.This,
		Authority:    o.Authority,
		ControlFlags: o.ControlFlags,
	}
}

func (o *ServerControlRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerControlOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.ControlFlags = op.ControlFlags
}
func (o *ServerControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ServerControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerControlResponse structure represents the ServerControl operation response
type ServerControlResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pctbOut: All fields of this parameter MUST be set to 0 on return.
	Out *csra.CertTransportBlob `idl:"name:pctbOut;pointer:ref" json:"out"`
	// Return: The ServerControl return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ServerControlResponse) xxx_ToOp(ctx context.Context) *xxx_ServerControlOperation {
	if o == nil {
		return &xxx_ServerControlOperation{}
	}
	return &xxx_ServerControlOperation{
		That:   o.That,
		Out:    o.Out,
		Return: o.Return,
	}
}

func (o *ServerControlResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerControlOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Out = op.Out
	o.Return = op.Return
}
func (o *ServerControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ServerControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PingOperation structure represents the Ping operation
type xxx_PingOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PingOperation) OpNum() int { return 18 }

func (o *xxx_PingOperation) OpName() string { return "/ICertAdminD/v0/Ping" }

func (o *xxx_PingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PingRequest structure represents the Ping operation request
type PingRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the pwszAuthority definition in section 3.1.4.1.1.
	//
	// Windows formats return values per the definition of HRESULT as specified in [MS-ERREF].
	// Negative values indicate errors, positive values indicate success. Specific values
	// are as specified in [MS-ERREF].
	//
	// The ICertAdminD::Ping method is as specified in [MS-WCCE] section 3.2.1.4.2.3.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *PingRequest) xxx_ToOp(ctx context.Context) *xxx_PingOperation {
	if o == nil {
		return &xxx_PingOperation{}
	}
	return &xxx_PingOperation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *PingRequest) xxx_FromOp(ctx context.Context, op *xxx_PingOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *PingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PingResponse structure represents the Ping operation response
type PingResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Ping return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PingResponse) xxx_ToOp(ctx context.Context) *xxx_PingOperation {
	if o == nil {
		return &xxx_PingOperation{}
	}
	return &xxx_PingOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PingResponse) xxx_FromOp(ctx context.Context, op *xxx_PingOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServerStateOperation structure represents the GetServerState operation
type xxx_GetServerStateOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	State     uint32         `idl:"name:pdwState" json:"state"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServerStateOperation) OpNum() int { return 19 }

func (o *xxx_GetServerStateOperation) OpName() string { return "/ICertAdminD/v0/GetServerState" }

func (o *xxx_GetServerStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwState {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwState {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetServerStateRequest structure represents the GetServerState operation request
type GetServerStateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *GetServerStateRequest) xxx_ToOp(ctx context.Context) *xxx_GetServerStateOperation {
	if o == nil {
		return &xxx_GetServerStateOperation{}
	}
	return &xxx_GetServerStateOperation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *GetServerStateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServerStateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *GetServerStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetServerStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServerStateResponse structure represents the GetServerState operation response
type GetServerStateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwState: A pointer to an unsigned integer value that specifies whether the caller
	// has permission to read from the CA database.
	//
	// The CA MUST return 1 for pdwState if the caller has permission to read from the CA
	// database. Otherwise, the CA MUST return 0.
	State uint32 `idl:"name:pdwState" json:"state"`
	// Return: The GetServerState return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServerStateResponse) xxx_ToOp(ctx context.Context) *xxx_GetServerStateOperation {
	if o == nil {
		return &xxx_GetServerStateOperation{}
	}
	return &xxx_GetServerStateOperation{
		That:   o.That,
		State:  o.State,
		Return: o.Return,
	}
}

func (o *GetServerStateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServerStateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.State = op.State
	o.Return = op.Return
}
func (o *GetServerStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetServerStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupPrepareOperation structure represents the BackupPrepare operation
type xxx_BackupPrepareOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority        string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Jet              uint32         `idl:"name:grbitJet" json:"jet"`
	BackupFlags      uint32         `idl:"name:dwBackupFlags" json:"backup_flags"`
	BackupAnnotation uint16         `idl:"name:pwszBackupAnnotation" json:"backup_annotation"`
	ClientID         uint32         `idl:"name:dwClientIdentifier" json:"client_id"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupPrepareOperation) OpNum() int { return 20 }

func (o *xxx_BackupPrepareOperation) OpName() string { return "/ICertAdminD/v0/BackupPrepare" }

func (o *xxx_BackupPrepareOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupPrepareOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// grbitJet {in} (1:(uint32))
	{
		if err := w.WriteData(o.Jet); err != nil {
			return err
		}
	}
	// dwBackupFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.BackupFlags); err != nil {
			return err
		}
	}
	// pwszBackupAnnotation {in} (1:{pointer=ref}*(1))(2:{alias=WCHAR}(wchar))
	{
		if err := w.WriteData(o.BackupAnnotation); err != nil {
			return err
		}
	}
	// dwClientIdentifier {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupPrepareOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// grbitJet {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Jet); err != nil {
			return err
		}
	}
	// dwBackupFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.BackupFlags); err != nil {
			return err
		}
	}
	// pwszBackupAnnotation {in} (1:{pointer=ref}*(1))(2:{alias=WCHAR}(wchar))
	{
		if err := w.ReadData(&o.BackupAnnotation); err != nil {
			return err
		}
	}
	// dwClientIdentifier {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupPrepareOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupPrepareOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupPrepareOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupPrepareRequest structure represents the BackupPrepare operation request
type BackupPrepareRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// grbitJet: An unsigned long value. This MUST be one of the following values.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 | Prepare for full backup of the CA database.                                      |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | Prepare for an incremental backup as opposed to a full backup. This means that   |
	//	|            | only the log files since the last full or incremental backup will be backed up.  |
	//	+------------+----------------------------------------------------------------------------------+
	Jet uint32 `idl:"name:grbitJet" json:"jet"`
	// dwBackupFlags: An unsigned long value. MUST be 0. MUST be ignored on receipt.
	BackupFlags uint32 `idl:"name:dwBackupFlags" json:"backup_flags"`
	// pwszBackupAnnotation: Not Used. Can be set to any arbitrary value, and MUST be ignored
	// on receipt.
	BackupAnnotation uint16 `idl:"name:pwszBackupAnnotation" json:"backup_annotation"`
	// dwClientIdentifier: An unsigned long value. Not used. MUST be 0. MUST be ignored
	// on receipt.
	//
	// If Config_CA_Interface_Flags contains the value IF_NOREMOTEICERTADMINBACKUP, the
	// server SHOULD return an error.<46>
	ClientID uint32 `idl:"name:dwClientIdentifier" json:"client_id"`
}

func (o *BackupPrepareRequest) xxx_ToOp(ctx context.Context) *xxx_BackupPrepareOperation {
	if o == nil {
		return &xxx_BackupPrepareOperation{}
	}
	return &xxx_BackupPrepareOperation{
		This:             o.This,
		Authority:        o.Authority,
		Jet:              o.Jet,
		BackupFlags:      o.BackupFlags,
		BackupAnnotation: o.BackupAnnotation,
		ClientID:         o.ClientID,
	}
}

func (o *BackupPrepareRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupPrepareOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.Jet = op.Jet
	o.BackupFlags = op.BackupFlags
	o.BackupAnnotation = op.BackupAnnotation
	o.ClientID = op.ClientID
}
func (o *BackupPrepareRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupPrepareRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupPrepareOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupPrepareResponse structure represents the BackupPrepare operation response
type BackupPrepareResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The BackupPrepare return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupPrepareResponse) xxx_ToOp(ctx context.Context) *xxx_BackupPrepareOperation {
	if o == nil {
		return &xxx_BackupPrepareOperation{}
	}
	return &xxx_BackupPrepareOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *BackupPrepareResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupPrepareOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *BackupPrepareResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupPrepareResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupPrepareOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupEndOperation structure represents the BackupEnd operation
type xxx_BackupEndOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupEndOperation) OpNum() int { return 21 }

func (o *xxx_BackupEndOperation) OpName() string { return "/ICertAdminD/v0/BackupEnd" }

func (o *xxx_BackupEndOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEndOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEndOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEndOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEndOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupEndOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupEndRequest structure represents the BackupEnd operation request
type BackupEndRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *BackupEndRequest) xxx_ToOp(ctx context.Context) *xxx_BackupEndOperation {
	if o == nil {
		return &xxx_BackupEndOperation{}
	}
	return &xxx_BackupEndOperation{
		This: o.This,
	}
}

func (o *BackupEndRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupEndOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *BackupEndRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupEndRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupEndOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupEndResponse structure represents the BackupEnd operation response
type BackupEndResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The BackupEnd return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupEndResponse) xxx_ToOp(ctx context.Context) *xxx_BackupEndOperation {
	if o == nil {
		return &xxx_BackupEndOperation{}
	}
	return &xxx_BackupEndOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *BackupEndResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupEndOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *BackupEndResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupEndResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupEndOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupGetAttachmentInformationOperation structure represents the BackupGetAttachmentInformation operation
type xxx_BackupGetAttachmentInformationOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	DBFiles       []string       `idl:"name:ppwszzDBFiles;size_is:(, pcwcDBFiles)" json:"db_files"`
	DBFilesLength int32          `idl:"name:pcwcDBFiles" json:"db_files_length"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupGetAttachmentInformationOperation) OpNum() int { return 22 }

func (o *xxx_BackupGetAttachmentInformationOperation) OpName() string {
	return "/ICertAdminD/v0/BackupGetAttachmentInformation"
}

func (o *xxx_BackupGetAttachmentInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetAttachmentInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetAttachmentInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetAttachmentInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.DBFiles != nil && o.DBFilesLength == 0 {
		o.DBFilesLength = int32(ndr.MultiSzLen(o.DBFiles))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetAttachmentInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppwszzDBFiles {out} (1:{multi_size, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=pcwcDBFiles,string](wchar))
	{
		if o.DBFiles != nil || o.DBFilesLength > 0 {
			_ptr_ppwszzDBFiles := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DBFilesLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				var _DBFiles_buf []uint16
				for i1 := range o.DBFiles {
					_DBFiles_buf = append(_DBFiles_buf, utf16.Encode([]rune(o.DBFiles[i1]))...)
					_DBFiles_buf = append(_DBFiles_buf, uint16(0))
				}
				_DBFiles_buf = append(_DBFiles_buf, uint16(0))
				for i1 := range _DBFiles_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_DBFiles_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_DBFiles_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DBFiles, _ptr_ppwszzDBFiles); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pcwcDBFiles {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.DBFilesLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetAttachmentInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppwszzDBFiles {out} (1:{multi_size, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=pcwcDBFiles,string](wchar))
	{
		_ptr_ppwszzDBFiles := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _DBFiles_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _DBFiles_buf", sizeInfo[0])
			}
			_DBFiles_buf = make([]uint16, sizeInfo[0])
			for i1 := range _DBFiles_buf {
				i1 := i1
				if err := w.ReadData(&_DBFiles_buf[i1]); err != nil {
					return err
				}
			}
			_tmp_DBFiles_buf := string(utf16.Decode(_DBFiles_buf))
			if _tmp_DBFiles_buf = strings.TrimRight(_tmp_DBFiles_buf, ndr.ZeroString); _tmp_DBFiles_buf != "" {
				o.DBFiles = strings.Split(_tmp_DBFiles_buf, ndr.ZeroString)
			}
			return nil
		})
		_s_ppwszzDBFiles := func(ptr interface{}) { o.DBFiles = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.DBFiles, _s_ppwszzDBFiles, _ptr_ppwszzDBFiles); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcwcDBFiles {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.DBFilesLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupGetAttachmentInformationRequest structure represents the BackupGetAttachmentInformation operation request
type BackupGetAttachmentInformationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *BackupGetAttachmentInformationRequest) xxx_ToOp(ctx context.Context) *xxx_BackupGetAttachmentInformationOperation {
	if o == nil {
		return &xxx_BackupGetAttachmentInformationOperation{}
	}
	return &xxx_BackupGetAttachmentInformationOperation{
		This: o.This,
	}
}

func (o *BackupGetAttachmentInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupGetAttachmentInformationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *BackupGetAttachmentInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupGetAttachmentInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupGetAttachmentInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupGetAttachmentInformationResponse structure represents the BackupGetAttachmentInformation operation response
type BackupGetAttachmentInformationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppwszzDBFiles: A pointer to a WCHAR pointer that will receive the list of null-terminated
	// database file names. Detailed information on database file name structure formatting
	// is specified in section 2.2.4.
	DBFiles []string `idl:"name:ppwszzDBFiles;size_is:(, pcwcDBFiles)" json:"db_files"`
	// pcwcDBFiles: A pointer to an integer value that contains the total length, in characters,
	// of all strings (including a NULL-terminator character) returned in ppwszzDBFiles.
	DBFilesLength int32 `idl:"name:pcwcDBFiles" json:"db_files_length"`
	// Return: The BackupGetAttachmentInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupGetAttachmentInformationResponse) xxx_ToOp(ctx context.Context) *xxx_BackupGetAttachmentInformationOperation {
	if o == nil {
		return &xxx_BackupGetAttachmentInformationOperation{}
	}
	return &xxx_BackupGetAttachmentInformationOperation{
		That:          o.That,
		DBFiles:       o.DBFiles,
		DBFilesLength: o.DBFilesLength,
		Return:        o.Return,
	}
}

func (o *BackupGetAttachmentInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupGetAttachmentInformationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DBFiles = op.DBFiles
	o.DBFilesLength = op.DBFilesLength
	o.Return = op.Return
}
func (o *BackupGetAttachmentInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupGetAttachmentInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupGetAttachmentInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupGetBackupLogsOperation structure represents the BackupGetBackupLogs operation
type xxx_BackupGetBackupLogsOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogFiles       []string       `idl:"name:ppwszzLogFiles;size_is:(, pcwcLogFiles)" json:"log_files"`
	LogFilesLength int32          `idl:"name:pcwcLogFiles" json:"log_files_length"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupGetBackupLogsOperation) OpNum() int { return 23 }

func (o *xxx_BackupGetBackupLogsOperation) OpName() string {
	return "/ICertAdminD/v0/BackupGetBackupLogs"
}

func (o *xxx_BackupGetBackupLogsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetBackupLogsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetBackupLogsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetBackupLogsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.LogFiles != nil && o.LogFilesLength == 0 {
		o.LogFilesLength = int32(ndr.MultiSzLen(o.LogFiles))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetBackupLogsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppwszzLogFiles {out} (1:{multi_size, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=pcwcLogFiles,string](wchar))
	{
		if o.LogFiles != nil || o.LogFilesLength > 0 {
			_ptr_ppwszzLogFiles := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.LogFilesLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				var _LogFiles_buf []uint16
				for i1 := range o.LogFiles {
					_LogFiles_buf = append(_LogFiles_buf, utf16.Encode([]rune(o.LogFiles[i1]))...)
					_LogFiles_buf = append(_LogFiles_buf, uint16(0))
				}
				_LogFiles_buf = append(_LogFiles_buf, uint16(0))
				for i1 := range _LogFiles_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_LogFiles_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_LogFiles_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LogFiles, _ptr_ppwszzLogFiles); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pcwcLogFiles {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.LogFilesLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetBackupLogsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppwszzLogFiles {out} (1:{multi_size, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=pcwcLogFiles,string](wchar))
	{
		_ptr_ppwszzLogFiles := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _LogFiles_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _LogFiles_buf", sizeInfo[0])
			}
			_LogFiles_buf = make([]uint16, sizeInfo[0])
			for i1 := range _LogFiles_buf {
				i1 := i1
				if err := w.ReadData(&_LogFiles_buf[i1]); err != nil {
					return err
				}
			}
			_tmp_LogFiles_buf := string(utf16.Decode(_LogFiles_buf))
			if _tmp_LogFiles_buf = strings.TrimRight(_tmp_LogFiles_buf, ndr.ZeroString); _tmp_LogFiles_buf != "" {
				o.LogFiles = strings.Split(_tmp_LogFiles_buf, ndr.ZeroString)
			}
			return nil
		})
		_s_ppwszzLogFiles := func(ptr interface{}) { o.LogFiles = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.LogFiles, _s_ppwszzLogFiles, _ptr_ppwszzLogFiles); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcwcLogFiles {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.LogFilesLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupGetBackupLogsRequest structure represents the BackupGetBackupLogs operation request
type BackupGetBackupLogsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *BackupGetBackupLogsRequest) xxx_ToOp(ctx context.Context) *xxx_BackupGetBackupLogsOperation {
	if o == nil {
		return &xxx_BackupGetBackupLogsOperation{}
	}
	return &xxx_BackupGetBackupLogsOperation{
		This: o.This,
	}
}

func (o *BackupGetBackupLogsRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupGetBackupLogsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *BackupGetBackupLogsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupGetBackupLogsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupGetBackupLogsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupGetBackupLogsResponse structure represents the BackupGetBackupLogs operation response
type BackupGetBackupLogsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppwszzLogFiles: A pointer to the WCHAR pointer that receives the list of null-terminated
	// log file names. Detailed information about database file name structure formatting
	// is specified in section 2.2.4.
	LogFiles []string `idl:"name:ppwszzLogFiles;size_is:(, pcwcLogFiles)" json:"log_files"`
	// pcwcLogFiles: A pointer to an integer value that contains the total length, in characters,
	// of all strings (including the NULL terminator character) returned in ppwszzLogFiles.
	LogFilesLength int32 `idl:"name:pcwcLogFiles" json:"log_files_length"`
	// Return: The BackupGetBackupLogs return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupGetBackupLogsResponse) xxx_ToOp(ctx context.Context) *xxx_BackupGetBackupLogsOperation {
	if o == nil {
		return &xxx_BackupGetBackupLogsOperation{}
	}
	return &xxx_BackupGetBackupLogsOperation{
		That:           o.That,
		LogFiles:       o.LogFiles,
		LogFilesLength: o.LogFilesLength,
		Return:         o.Return,
	}
}

func (o *BackupGetBackupLogsResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupGetBackupLogsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LogFiles = op.LogFiles
	o.LogFilesLength = op.LogFilesLength
	o.Return = op.Return
}
func (o *BackupGetBackupLogsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupGetBackupLogsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupGetBackupLogsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupOpenFileOperation structure represents the BackupOpenFile operation
type xxx_BackupOpenFileOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   string         `idl:"name:pwszPath;string;pointer:unique" json:"path"`
	Length uint64         `idl:"name:pliLength" json:"length"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupOpenFileOperation) OpNum() int { return 24 }

func (o *xxx_BackupOpenFileOperation) OpName() string { return "/ICertAdminD/v0/BackupOpenFile" }

func (o *xxx_BackupOpenFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOpenFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszPath {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_pwszPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_pwszPath); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOpenFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszPath {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_pwszPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_pwszPath, _ptr_pwszPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOpenFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOpenFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pliLength {out} (1:{pointer=ref}*(1)(uint64))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOpenFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pliLength {out} (1:{pointer=ref}*(1)(uint64))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupOpenFileRequest structure represents the BackupOpenFile operation request
type BackupOpenFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszPath: A null-terminated UNICODE string that specifies the path to the targeted
	// file. The file name MUST be UNC form, for example: "\\server\sharepoint\...path...\filename.ext".
	Path string `idl:"name:pwszPath;string;pointer:unique" json:"path"`
}

func (o *BackupOpenFileRequest) xxx_ToOp(ctx context.Context) *xxx_BackupOpenFileOperation {
	if o == nil {
		return &xxx_BackupOpenFileOperation{}
	}
	return &xxx_BackupOpenFileOperation{
		This: o.This,
		Path: o.Path,
	}
}

func (o *BackupOpenFileRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupOpenFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *BackupOpenFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupOpenFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupOpenFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupOpenFileResponse structure represents the BackupOpenFile operation response
type BackupOpenFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pliLength: A pointer to a signed 64-bit integer that receives the size, in bytes,
	// of the targeted file.
	Length uint64 `idl:"name:pliLength" json:"length"`
	// Return: The BackupOpenFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupOpenFileResponse) xxx_ToOp(ctx context.Context) *xxx_BackupOpenFileOperation {
	if o == nil {
		return &xxx_BackupOpenFileOperation{}
	}
	return &xxx_BackupOpenFileOperation{
		That:   o.That,
		Length: o.Length,
		Return: o.Return,
	}
}

func (o *BackupOpenFileResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupOpenFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Length = op.Length
	o.Return = op.Return
}
func (o *BackupOpenFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupOpenFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupOpenFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupReadFileOperation structure represents the BackupReadFile operation
type xxx_BackupReadFileOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	Buffer       []byte         `idl:"name:pbBuffer;size_is:(cbBuffer);pointer:ref" json:"buffer"`
	BufferLength int32          `idl:"name:cbBuffer" json:"buffer_length"`
	ReadLength   int32          `idl:"name:pcbRead" json:"read_length"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupReadFileOperation) OpNum() int { return 25 }

func (o *xxx_BackupReadFileOperation) OpName() string { return "/ICertAdminD/v0/BackupReadFile" }

func (o *xxx_BackupReadFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupReadFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cbBuffer {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.BufferLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupReadFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbBuffer {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.BufferLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupReadFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupReadFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbBuffer {out} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=cbBuffer](uint8))
	{
		dimSize1 := uint64(o.BufferLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Buffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Buffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pcbRead {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.ReadLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupReadFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbBuffer {out} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=cbBuffer](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
	}
	// pcbRead {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.ReadLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupReadFileRequest structure represents the BackupReadFile operation request
type BackupReadFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// cbBuffer: The size, in bytes, of the preceding buffer. This parameter MUST be a multiple
	// of the page size of the operating system.
	BufferLength int32 `idl:"name:cbBuffer" json:"buffer_length"`
}

func (o *BackupReadFileRequest) xxx_ToOp(ctx context.Context) *xxx_BackupReadFileOperation {
	if o == nil {
		return &xxx_BackupReadFileOperation{}
	}
	return &xxx_BackupReadFileOperation{
		This:         o.This,
		BufferLength: o.BufferLength,
	}
}

func (o *BackupReadFileRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupReadFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BufferLength = op.BufferLength
}
func (o *BackupReadFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupReadFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupReadFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupReadFileResponse structure represents the BackupReadFile operation response
type BackupReadFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbBuffer: A pointer to the buffer that receives the read data.
	Buffer []byte `idl:"name:pbBuffer;size_is:(cbBuffer);pointer:ref" json:"buffer"`
	// pcbRead: A pointer to an integer that receives the actual number of bytes read.
	ReadLength int32 `idl:"name:pcbRead" json:"read_length"`
	// Return: The BackupReadFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupReadFileResponse) xxx_ToOp(ctx context.Context) *xxx_BackupReadFileOperation {
	if o == nil {
		return &xxx_BackupReadFileOperation{}
	}
	return &xxx_BackupReadFileOperation{
		That:       o.That,
		Buffer:     o.Buffer,
		ReadLength: o.ReadLength,
		Return:     o.Return,
	}
}

func (o *BackupReadFileResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupReadFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Buffer = op.Buffer
	o.ReadLength = op.ReadLength
	o.Return = op.Return
}
func (o *BackupReadFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupReadFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupReadFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupCloseFileOperation structure represents the BackupCloseFile operation
type xxx_BackupCloseFileOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupCloseFileOperation) OpNum() int { return 26 }

func (o *xxx_BackupCloseFileOperation) OpName() string { return "/ICertAdminD/v0/BackupCloseFile" }

func (o *xxx_BackupCloseFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupCloseFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupCloseFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupCloseFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupCloseFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupCloseFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupCloseFileRequest structure represents the BackupCloseFile operation request
type BackupCloseFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *BackupCloseFileRequest) xxx_ToOp(ctx context.Context) *xxx_BackupCloseFileOperation {
	if o == nil {
		return &xxx_BackupCloseFileOperation{}
	}
	return &xxx_BackupCloseFileOperation{
		This: o.This,
	}
}

func (o *BackupCloseFileRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupCloseFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *BackupCloseFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupCloseFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupCloseFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupCloseFileResponse structure represents the BackupCloseFile operation response
type BackupCloseFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The BackupCloseFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupCloseFileResponse) xxx_ToOp(ctx context.Context) *xxx_BackupCloseFileOperation {
	if o == nil {
		return &xxx_BackupCloseFileOperation{}
	}
	return &xxx_BackupCloseFileOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *BackupCloseFileResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupCloseFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *BackupCloseFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupCloseFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupCloseFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupTruncateLogsOperation structure represents the BackupTruncateLogs operation
type xxx_BackupTruncateLogsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupTruncateLogsOperation) OpNum() int { return 27 }

func (o *xxx_BackupTruncateLogsOperation) OpName() string {
	return "/ICertAdminD/v0/BackupTruncateLogs"
}

func (o *xxx_BackupTruncateLogsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupTruncateLogsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupTruncateLogsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupTruncateLogsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupTruncateLogsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupTruncateLogsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupTruncateLogsRequest structure represents the BackupTruncateLogs operation request
type BackupTruncateLogsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *BackupTruncateLogsRequest) xxx_ToOp(ctx context.Context) *xxx_BackupTruncateLogsOperation {
	if o == nil {
		return &xxx_BackupTruncateLogsOperation{}
	}
	return &xxx_BackupTruncateLogsOperation{
		This: o.This,
	}
}

func (o *BackupTruncateLogsRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupTruncateLogsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *BackupTruncateLogsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupTruncateLogsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupTruncateLogsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupTruncateLogsResponse structure represents the BackupTruncateLogs operation response
type BackupTruncateLogsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The BackupTruncateLogs return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupTruncateLogsResponse) xxx_ToOp(ctx context.Context) *xxx_BackupTruncateLogsOperation {
	if o == nil {
		return &xxx_BackupTruncateLogsOperation{}
	}
	return &xxx_BackupTruncateLogsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *BackupTruncateLogsResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupTruncateLogsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *BackupTruncateLogsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupTruncateLogsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupTruncateLogsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportCertificateOperation structure represents the ImportCertificate operation
type xxx_ImportCertificateOperation struct {
	This        *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority   string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Certificate *csra.CertTransportBlob `idl:"name:pctbCertificate;pointer:ref" json:"certificate"`
	Flags       int32                   `idl:"name:dwFlags" json:"flags"`
	RequestID   int32                   `idl:"name:pdwRequestId" json:"request_id"`
	Return      int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportCertificateOperation) OpNum() int { return 28 }

func (o *xxx_ImportCertificateOperation) OpName() string { return "/ICertAdminD/v0/ImportCertificate" }

func (o *xxx_ImportCertificateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportCertificateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Authority != "" {
			_ptr_pwszAuthority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Authority); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Authority, _ptr_pwszAuthority); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pctbCertificate {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Certificate != nil {
			if err := o.Certificate.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csra.CertTransportBlob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportCertificateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszAuthority {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszAuthority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Authority); err != nil {
				return err
			}
			return nil
		})
		_s_pwszAuthority := func(ptr interface{}) { o.Authority = *ptr.(*string) }
		if err := w.ReadPointer(&o.Authority, _s_pwszAuthority, _ptr_pwszAuthority); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pctbCertificate {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Certificate == nil {
			o.Certificate = &csra.CertTransportBlob{}
		}
		if err := o.Certificate.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportCertificateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportCertificateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwRequestId {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.RequestID); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportCertificateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwRequestId {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.RequestID); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ImportCertificateRequest structure represents the ImportCertificate operation request
type ImportCertificateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// pctbCertificate: A CERTTRANSBLOB that contains an ASN.1 DER–encoded (as specified
	// in [X660] and [X690]) certificate that is inserted into the CA database.
	Certificate *csra.CertTransportBlob `idl:"name:pctbCertificate;pointer:ref" json:"certificate"`
	// dwFlags: A LONG value that MUST take one of the following values.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    0 | If this value is set, the CA server does not allow certificates that are not     |
	//	|                                      | issued by it to be imported into its database.                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| FLAG_ALLOW_IMPORT_FOREIGN 0x00010000 | A request to the CA server to allow certificates that are not issued by it to be |
	//	|                                      | imported into its database.                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ICF_EXISTINGROW 0x00020000           | A request to the CA to associate the imported certificates with an existing      |
	//	|                                      | request row.                                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Flags int32 `idl:"name:dwFlags" json:"flags"`
}

func (o *ImportCertificateRequest) xxx_ToOp(ctx context.Context) *xxx_ImportCertificateOperation {
	if o == nil {
		return &xxx_ImportCertificateOperation{}
	}
	return &xxx_ImportCertificateOperation{
		This:        o.This,
		Authority:   o.Authority,
		Certificate: o.Certificate,
		Flags:       o.Flags,
	}
}

func (o *ImportCertificateRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportCertificateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.Certificate = op.Certificate
	o.Flags = op.Flags
}
func (o *ImportCertificateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ImportCertificateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportCertificateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportCertificateResponse structure represents the ImportCertificate operation response
type ImportCertificateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwRequestId: Returns the request ID for the imported certificate. This is used to
	// refer to the certificate after it is imported into the database.
	//
	// ImportCertificate imports a certificate into the CA database Request table.
	RequestID int32 `idl:"name:pdwRequestId" json:"request_id"`
	// Return: The ImportCertificate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportCertificateResponse) xxx_ToOp(ctx context.Context) *xxx_ImportCertificateOperation {
	if o == nil {
		return &xxx_ImportCertificateOperation{}
	}
	return &xxx_ImportCertificateOperation{
		That:      o.That,
		RequestID: o.RequestID,
		Return:    o.Return,
	}
}

func (o *ImportCertificateResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportCertificateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RequestID = op.RequestID
	o.Return = op.Return
}
func (o *ImportCertificateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ImportCertificateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportCertificateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupGetDynamicFilesOperation structure represents the BackupGetDynamicFiles operation
type xxx_BackupGetDynamicFilesOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Files       []string       `idl:"name:ppwszzFiles;size_is:(, pcwcFiles)" json:"files"`
	FilesLength int32          `idl:"name:pcwcFiles" json:"files_length"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupGetDynamicFilesOperation) OpNum() int { return 29 }

func (o *xxx_BackupGetDynamicFilesOperation) OpName() string {
	return "/ICertAdminD/v0/BackupGetDynamicFiles"
}

func (o *xxx_BackupGetDynamicFilesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetDynamicFilesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetDynamicFilesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetDynamicFilesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Files != nil && o.FilesLength == 0 {
		o.FilesLength = int32(ndr.MultiSzLen(o.Files))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetDynamicFilesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppwszzFiles {out} (1:{multi_size, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=pcwcFiles,string](wchar))
	{
		if o.Files != nil || o.FilesLength > 0 {
			_ptr_ppwszzFiles := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.FilesLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				var _Files_buf []uint16
				for i1 := range o.Files {
					_Files_buf = append(_Files_buf, utf16.Encode([]rune(o.Files[i1]))...)
					_Files_buf = append(_Files_buf, uint16(0))
				}
				_Files_buf = append(_Files_buf, uint16(0))
				for i1 := range _Files_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_Files_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_Files_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Files, _ptr_ppwszzFiles); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pcwcFiles {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.FilesLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupGetDynamicFilesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppwszzFiles {out} (1:{multi_size, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=pcwcFiles,string](wchar))
	{
		_ptr_ppwszzFiles := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _Files_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _Files_buf", sizeInfo[0])
			}
			_Files_buf = make([]uint16, sizeInfo[0])
			for i1 := range _Files_buf {
				i1 := i1
				if err := w.ReadData(&_Files_buf[i1]); err != nil {
					return err
				}
			}
			_tmp_Files_buf := string(utf16.Decode(_Files_buf))
			if _tmp_Files_buf = strings.TrimRight(_tmp_Files_buf, ndr.ZeroString); _tmp_Files_buf != "" {
				o.Files = strings.Split(_tmp_Files_buf, ndr.ZeroString)
			}
			return nil
		})
		_s_ppwszzFiles := func(ptr interface{}) { o.Files = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Files, _s_ppwszzFiles, _ptr_ppwszzFiles); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcwcFiles {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.FilesLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupGetDynamicFilesRequest structure represents the BackupGetDynamicFiles operation request
type BackupGetDynamicFilesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *BackupGetDynamicFilesRequest) xxx_ToOp(ctx context.Context) *xxx_BackupGetDynamicFilesOperation {
	if o == nil {
		return &xxx_BackupGetDynamicFilesOperation{}
	}
	return &xxx_BackupGetDynamicFilesOperation{
		This: o.This,
	}
}

func (o *BackupGetDynamicFilesRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupGetDynamicFilesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *BackupGetDynamicFilesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupGetDynamicFilesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupGetDynamicFilesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupGetDynamicFilesResponse structure represents the BackupGetDynamicFiles operation response
type BackupGetDynamicFilesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppwszzFiles: A pointer to a WCHAR pointer that receives the list of null-terminated
	// dynamic file names that are used by a CA.
	Files []string `idl:"name:ppwszzFiles;size_is:(, pcwcFiles)" json:"files"`
	// pcwcFiles: A pointer to the LONG value that specifies the number of characters in
	// ppwszzFiles.
	FilesLength int32 `idl:"name:pcwcFiles" json:"files_length"`
	// Return: The BackupGetDynamicFiles return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BackupGetDynamicFilesResponse) xxx_ToOp(ctx context.Context) *xxx_BackupGetDynamicFilesOperation {
	if o == nil {
		return &xxx_BackupGetDynamicFilesOperation{}
	}
	return &xxx_BackupGetDynamicFilesOperation{
		That:        o.That,
		Files:       o.Files,
		FilesLength: o.FilesLength,
		Return:      o.Return,
	}
}

func (o *BackupGetDynamicFilesResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupGetDynamicFilesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Files = op.Files
	o.FilesLength = op.FilesLength
	o.Return = op.Return
}
func (o *BackupGetDynamicFilesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupGetDynamicFilesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupGetDynamicFilesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RestoreGetDatabaseLocationsOperation structure represents the RestoreGetDatabaseLocations operation
type xxx_RestoreGetDatabaseLocationsOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	DatabaseLocations []string       `idl:"name:ppwszzDatabaseLocations;size_is:(, pcwcPaths)" json:"database_locations"`
	PathsLength       int32          `idl:"name:pcwcPaths" json:"paths_length"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RestoreGetDatabaseLocationsOperation) OpNum() int { return 30 }

func (o *xxx_RestoreGetDatabaseLocationsOperation) OpName() string {
	return "/ICertAdminD/v0/RestoreGetDatabaseLocations"
}

func (o *xxx_RestoreGetDatabaseLocationsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreGetDatabaseLocationsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreGetDatabaseLocationsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreGetDatabaseLocationsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.DatabaseLocations != nil && o.PathsLength == 0 {
		o.PathsLength = int32(ndr.MultiSzLen(o.DatabaseLocations))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreGetDatabaseLocationsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppwszzDatabaseLocations {out} (1:{multi_size, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=pcwcPaths,string](wchar))
	{
		if o.DatabaseLocations != nil || o.PathsLength > 0 {
			_ptr_ppwszzDatabaseLocations := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.PathsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				var _DatabaseLocations_buf []uint16
				for i1 := range o.DatabaseLocations {
					_DatabaseLocations_buf = append(_DatabaseLocations_buf, utf16.Encode([]rune(o.DatabaseLocations[i1]))...)
					_DatabaseLocations_buf = append(_DatabaseLocations_buf, uint16(0))
				}
				_DatabaseLocations_buf = append(_DatabaseLocations_buf, uint16(0))
				for i1 := range _DatabaseLocations_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_DatabaseLocations_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_DatabaseLocations_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DatabaseLocations, _ptr_ppwszzDatabaseLocations); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pcwcPaths {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.PathsLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RestoreGetDatabaseLocationsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppwszzDatabaseLocations {out} (1:{multi_size, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=pcwcPaths,string](wchar))
	{
		_ptr_ppwszzDatabaseLocations := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _DatabaseLocations_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _DatabaseLocations_buf", sizeInfo[0])
			}
			_DatabaseLocations_buf = make([]uint16, sizeInfo[0])
			for i1 := range _DatabaseLocations_buf {
				i1 := i1
				if err := w.ReadData(&_DatabaseLocations_buf[i1]); err != nil {
					return err
				}
			}
			_tmp_DatabaseLocations_buf := string(utf16.Decode(_DatabaseLocations_buf))
			if _tmp_DatabaseLocations_buf = strings.TrimRight(_tmp_DatabaseLocations_buf, ndr.ZeroString); _tmp_DatabaseLocations_buf != "" {
				o.DatabaseLocations = strings.Split(_tmp_DatabaseLocations_buf, ndr.ZeroString)
			}
			return nil
		})
		_s_ppwszzDatabaseLocations := func(ptr interface{}) { o.DatabaseLocations = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.DatabaseLocations, _s_ppwszzDatabaseLocations, _ptr_ppwszzDatabaseLocations); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcwcPaths {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.PathsLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RestoreGetDatabaseLocationsRequest structure represents the RestoreGetDatabaseLocations operation request
type RestoreGetDatabaseLocationsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RestoreGetDatabaseLocationsRequest) xxx_ToOp(ctx context.Context) *xxx_RestoreGetDatabaseLocationsOperation {
	if o == nil {
		return &xxx_RestoreGetDatabaseLocationsOperation{}
	}
	return &xxx_RestoreGetDatabaseLocationsOperation{
		This: o.This,
	}
}

func (o *RestoreGetDatabaseLocationsRequest) xxx_FromOp(ctx context.Context, op *xxx_RestoreGetDatabaseLocationsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RestoreGetDatabaseLocationsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RestoreGetDatabaseLocationsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreGetDatabaseLocationsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RestoreGetDatabaseLocationsResponse structure represents the RestoreGetDatabaseLocations operation response
type RestoreGetDatabaseLocationsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppwszzDatabaseLocations: A pointer to a WCHAR pointer that will receive the list
	// of null-terminated database location names and the log directory name. Detailed information
	// about database file name structure formatting is specified in section 2.2.4.
	DatabaseLocations []string `idl:"name:ppwszzDatabaseLocations;size_is:(, pcwcPaths)" json:"database_locations"`
	// pcwcPaths: A pointer to the LONG value that specifies the number of characters in
	// ppwszzDatabaseLocations.
	PathsLength int32 `idl:"name:pcwcPaths" json:"paths_length"`
	// Return: The RestoreGetDatabaseLocations return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RestoreGetDatabaseLocationsResponse) xxx_ToOp(ctx context.Context) *xxx_RestoreGetDatabaseLocationsOperation {
	if o == nil {
		return &xxx_RestoreGetDatabaseLocationsOperation{}
	}
	return &xxx_RestoreGetDatabaseLocationsOperation{
		That:              o.That,
		DatabaseLocations: o.DatabaseLocations,
		PathsLength:       o.PathsLength,
		Return:            o.Return,
	}
}

func (o *RestoreGetDatabaseLocationsResponse) xxx_FromOp(ctx context.Context, op *xxx_RestoreGetDatabaseLocationsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DatabaseLocations = op.DatabaseLocations
	o.PathsLength = op.PathsLength
	o.Return = op.Return
}
func (o *RestoreGetDatabaseLocationsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RestoreGetDatabaseLocationsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RestoreGetDatabaseLocationsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
