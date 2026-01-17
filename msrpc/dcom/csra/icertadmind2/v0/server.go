package icertadmind2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	icertadmind "github.com/oiweiwei/go-msrpc/msrpc/dcom/csra/icertadmind/v0"
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
	_ = icertadmind.GoPackage
)

// ICertAdminD2 server interface.
type CertAdminD2Server interface {

	// ICertAdminD base class.
	icertadmind.CertAdminDServer

	// The PublishCRLs method instructs a CA to publish CRLs and delta CRLs. This call can
	// either cause the republishing of the current CRLs or cause the CA to create and publish
	// new CRLs.
	PublishCRLs(context.Context, *PublishCRLsRequest) (*PublishCRLsResponse, error)

	// The GetCAProperty method is used to retrieve the value of a specific property from
	// the CA.
	GetCAProperty(context.Context, *GetCAPropertyRequest) (*GetCAPropertyResponse, error)

	// The SetCAProperty method is used to set CA properties.
	SetCAProperty(context.Context, *SetCAPropertyRequest) (*SetCAPropertyResponse, error)

	// The GetCAPropertyInfo method is used to retrieve information about a property on
	// the CA, such as its type and length.
	GetCAPropertyInfo(context.Context, *GetCAPropertyInfoRequest) (*GetCAPropertyInfoResponse, error)

	// The EnumViewColumnTable method retrieves information about one or more columns from
	// the specified CA database table.
	EnumViewColumnTable(context.Context, *EnumViewColumnTableRequest) (*EnumViewColumnTableResponse, error)

	// The GetCASecurity method is used to retrieve CA security, as defined in Abstract
	// Data Model (section 3.1.1).
	GetCASecurity(context.Context, *GetCASecurityRequest) (*GetCASecurityResponse, error)

	// The SetCASecurity method is used to set the CA security, as defined in the Abstract
	// Data Model (section 3.1.1).
	SetCASecurity(context.Context, *SetCASecurityRequest) (*SetCASecurityResponse, error)

	// The Ping2 method is used to determine if the CA service is started and responding.
	Ping2(context.Context, *Ping2Request) (*Ping2Response, error)

	// The GetArchivedKey method is used to retrieve an archived private key and the associated
	// certificate.
	GetArchivedKey(context.Context, *GetArchivedKeyRequest) (*GetArchivedKeyResponse, error)

	// The GetAuditFilter method retrieves the list of events for which the CA server is
	// currently set to create security audit events, as specified in [CIMC-PP].
	GetAuditFilter(context.Context, *GetAuditFilterRequest) (*GetAuditFilterResponse, error)

	// The SetAuditFilter method sets the list of events for which the CA server MUST create
	// security audit events, as specified in [CIMC-PP].
	SetAuditFilter(context.Context, *SetAuditFilterRequest) (*SetAuditFilterResponse, error)

	// The GetOfficerRights method is used to retrieve the Officer rights, as specified
	// in [CIMC-PP].
	GetOfficerRights(context.Context, *GetOfficerRightsRequest) (*GetOfficerRightsResponse, error)

	// The SetOfficerRights method is used to set Officer rights or Enrollment Agent rights.
	// Information on role separation is specified in [CIMC-PP].
	SetOfficerRights(context.Context, *SetOfficerRightsRequest) (*SetOfficerRightsResponse, error)

	// The GetConfigEntry method retrieves the CAs that persisted the configuration data
	// listed in section 3.1.1.10. Configuration data is represented as a hierarchical data
	// structure with the following format: [\pwszAuthority][\pwszNodePath][\pwszEntry].
	GetConfigEntry(context.Context, *GetConfigEntryRequest) (*GetConfigEntryResponse, error)

	// The SetConfigEntry method is used to set the CA's persisted configuration data that
	// is listed in section 3.1.1.10.
	SetConfigEntry(context.Context, *SetConfigEntryRequest) (*SetConfigEntryResponse, error)

	// The ImportKey method adds an encrypted key set to an item in the CA database.
	ImportKey(context.Context, *ImportKeyRequest) (*ImportKeyResponse, error)

	// The GetMyRoles method retrieves the CA roles, as specified in [CIMC-PP], assigned
	// to the user who calls the method.
	GetMyRoles(context.Context, *GetMyRolesRequest) (*GetMyRolesResponse, error)

	// The DeleteRow method deletes a row or set of rows from a database table.
	DeleteRow(context.Context, *DeleteRowRequest) (*DeleteRowResponse, error)
}

func RegisterCertAdminD2Server(conn dcerpc.Conn, o CertAdminD2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCertAdminD2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CertAdminD2SyntaxV0_0))...)
}

func NewCertAdminD2ServerHandle(o CertAdminD2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CertAdminD2ServerHandle(ctx, o, opNum, r)
	}
}

func CertAdminD2ServerHandle(ctx context.Context, o CertAdminD2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 31 {
		// ICertAdminD base method.
		return icertadmind.CertAdminDServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 31: // PublishCRLs
		op := &xxx_PublishCRLsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PublishCRLsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PublishCRLs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // GetCAProperty
		op := &xxx_GetCAPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCAPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCAProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // SetCAProperty
		op := &xxx_SetCAPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCAPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCAProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // GetCAPropertyInfo
		op := &xxx_GetCAPropertyInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCAPropertyInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCAPropertyInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // EnumViewColumnTable
		op := &xxx_EnumViewColumnTableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumViewColumnTableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumViewColumnTable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // GetCASecurity
		op := &xxx_GetCASecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCASecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCASecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // SetCASecurity
		op := &xxx_SetCASecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCASecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCASecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // Ping2
		op := &xxx_Ping2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Ping2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Ping2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // GetArchivedKey
		op := &xxx_GetArchivedKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetArchivedKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetArchivedKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // GetAuditFilter
		op := &xxx_GetAuditFilterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAuditFilterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAuditFilter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // SetAuditFilter
		op := &xxx_SetAuditFilterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAuditFilterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAuditFilter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // GetOfficerRights
		op := &xxx_GetOfficerRightsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOfficerRightsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOfficerRights(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // SetOfficerRights
		op := &xxx_SetOfficerRightsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOfficerRightsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOfficerRights(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // GetConfigEntry
		op := &xxx_GetConfigEntryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigEntryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfigEntry(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // SetConfigEntry
		op := &xxx_SetConfigEntryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConfigEntryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConfigEntry(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // ImportKey
		op := &xxx_ImportKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ImportKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ImportKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // GetMyRoles
		op := &xxx_GetMyRolesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMyRolesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMyRoles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // DeleteRow
		op := &xxx_DeleteRowOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteRowRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteRow(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICertAdminD2
type UnimplementedCertAdminD2Server struct {
	icertadmind.UnimplementedCertAdminDServer
}

func (UnimplementedCertAdminD2Server) PublishCRLs(context.Context, *PublishCRLsRequest) (*PublishCRLsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) GetCAProperty(context.Context, *GetCAPropertyRequest) (*GetCAPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) SetCAProperty(context.Context, *SetCAPropertyRequest) (*SetCAPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) GetCAPropertyInfo(context.Context, *GetCAPropertyInfoRequest) (*GetCAPropertyInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) EnumViewColumnTable(context.Context, *EnumViewColumnTableRequest) (*EnumViewColumnTableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) GetCASecurity(context.Context, *GetCASecurityRequest) (*GetCASecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) SetCASecurity(context.Context, *SetCASecurityRequest) (*SetCASecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) Ping2(context.Context, *Ping2Request) (*Ping2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) GetArchivedKey(context.Context, *GetArchivedKeyRequest) (*GetArchivedKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) GetAuditFilter(context.Context, *GetAuditFilterRequest) (*GetAuditFilterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) SetAuditFilter(context.Context, *SetAuditFilterRequest) (*SetAuditFilterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) GetOfficerRights(context.Context, *GetOfficerRightsRequest) (*GetOfficerRightsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) SetOfficerRights(context.Context, *SetOfficerRightsRequest) (*SetOfficerRightsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) GetConfigEntry(context.Context, *GetConfigEntryRequest) (*GetConfigEntryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) SetConfigEntry(context.Context, *SetConfigEntryRequest) (*SetConfigEntryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) ImportKey(context.Context, *ImportKeyRequest) (*ImportKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) GetMyRoles(context.Context, *GetMyRolesRequest) (*GetMyRolesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertAdminD2Server) DeleteRow(context.Context, *DeleteRowRequest) (*DeleteRowResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CertAdminD2Server = (*UnimplementedCertAdminD2Server)(nil)
