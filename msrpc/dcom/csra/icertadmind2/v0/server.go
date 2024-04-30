package icertadmind2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
		in := &PublishCRLsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PublishCRLs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // GetCAProperty
		in := &GetCAPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCAProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // SetCAProperty
		in := &SetCAPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetCAProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // GetCAPropertyInfo
		in := &GetCAPropertyInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCAPropertyInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // EnumViewColumnTable
		in := &EnumViewColumnTableRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumViewColumnTable(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // GetCASecurity
		in := &GetCASecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCASecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // SetCASecurity
		in := &SetCASecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetCASecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // Ping2
		in := &Ping2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Ping2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // GetArchivedKey
		in := &GetArchivedKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetArchivedKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // GetAuditFilter
		in := &GetAuditFilterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAuditFilter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // SetAuditFilter
		in := &SetAuditFilterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAuditFilter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // GetOfficerRights
		in := &GetOfficerRightsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOfficerRights(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // SetOfficerRights
		in := &SetOfficerRightsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOfficerRights(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // GetConfigEntry
		in := &GetConfigEntryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetConfigEntry(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // SetConfigEntry
		in := &SetConfigEntryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetConfigEntry(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // ImportKey
		in := &ImportKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ImportKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // GetMyRoles
		in := &GetMyRolesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMyRoles(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // DeleteRow
		in := &DeleteRowRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteRow(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
