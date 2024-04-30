package icertadmind2

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
	icertadmind "github.com/oiweiwei/go-msrpc/msrpc/dcom/csra/icertadmind/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = icertadmind.GoPackage
	_ = dtyp.GoPackage
	_ = csra.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csra"
)

var (
	// ICertAdminD2 interface identifier 7fe0d935-dda6-443f-85d0-1cfb58fe41dd
	CertAdminD2IID = &dcom.IID{Data1: 0x7fe0d935, Data2: 0xdda6, Data3: 0x443f, Data4: []byte{0x85, 0xd0, 0x1c, 0xfb, 0x58, 0xfe, 0x41, 0xdd}}
	// Syntax UUID
	CertAdminD2SyntaxUUID = &uuid.UUID{TimeLow: 0x7fe0d935, TimeMid: 0xdda6, TimeHiAndVersion: 0x443f, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0xd0, Node: [6]uint8{0x1c, 0xfb, 0x58, 0xfe, 0x41, 0xdd}}
	// Syntax ID
	CertAdminD2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CertAdminD2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICertAdminD2 interface.
type CertAdminD2Client interface {

	// ICertAdminD retrieval method.
	CertAdminD() icertadmind.CertAdminDClient

	// The PublishCRLs method instructs a CA to publish CRLs and delta CRLs. This call can
	// either cause the republishing of the current CRLs or cause the CA to create and publish
	// new CRLs.
	PublishCRLs(context.Context, *PublishCRLsRequest, ...dcerpc.CallOption) (*PublishCRLsResponse, error)

	// The GetCAProperty method is used to retrieve the value of a specific property from
	// the CA.
	GetCAProperty(context.Context, *GetCAPropertyRequest, ...dcerpc.CallOption) (*GetCAPropertyResponse, error)

	// The SetCAProperty method is used to set CA properties.
	SetCAProperty(context.Context, *SetCAPropertyRequest, ...dcerpc.CallOption) (*SetCAPropertyResponse, error)

	// The GetCAPropertyInfo method is used to retrieve information about a property on
	// the CA, such as its type and length.
	GetCAPropertyInfo(context.Context, *GetCAPropertyInfoRequest, ...dcerpc.CallOption) (*GetCAPropertyInfoResponse, error)

	// The EnumViewColumnTable method retrieves information about one or more columns from
	// the specified CA database table.
	EnumViewColumnTable(context.Context, *EnumViewColumnTableRequest, ...dcerpc.CallOption) (*EnumViewColumnTableResponse, error)

	// The GetCASecurity method is used to retrieve CA security, as defined in Abstract
	// Data Model (section 3.1.1).
	GetCASecurity(context.Context, *GetCASecurityRequest, ...dcerpc.CallOption) (*GetCASecurityResponse, error)

	// The SetCASecurity method is used to set the CA security, as defined in the Abstract
	// Data Model (section 3.1.1).
	SetCASecurity(context.Context, *SetCASecurityRequest, ...dcerpc.CallOption) (*SetCASecurityResponse, error)

	// The Ping2 method is used to determine if the CA service is started and responding.
	Ping2(context.Context, *Ping2Request, ...dcerpc.CallOption) (*Ping2Response, error)

	// The GetArchivedKey method is used to retrieve an archived private key and the associated
	// certificate.
	GetArchivedKey(context.Context, *GetArchivedKeyRequest, ...dcerpc.CallOption) (*GetArchivedKeyResponse, error)

	// The GetAuditFilter method retrieves the list of events for which the CA server is
	// currently set to create security audit events, as specified in [CIMC-PP].
	GetAuditFilter(context.Context, *GetAuditFilterRequest, ...dcerpc.CallOption) (*GetAuditFilterResponse, error)

	// The SetAuditFilter method sets the list of events for which the CA server MUST create
	// security audit events, as specified in [CIMC-PP].
	SetAuditFilter(context.Context, *SetAuditFilterRequest, ...dcerpc.CallOption) (*SetAuditFilterResponse, error)

	// The GetOfficerRights method is used to retrieve the Officer rights, as specified
	// in [CIMC-PP].
	GetOfficerRights(context.Context, *GetOfficerRightsRequest, ...dcerpc.CallOption) (*GetOfficerRightsResponse, error)

	// The SetOfficerRights method is used to set Officer rights or Enrollment Agent rights.
	// Information on role separation is specified in [CIMC-PP].
	SetOfficerRights(context.Context, *SetOfficerRightsRequest, ...dcerpc.CallOption) (*SetOfficerRightsResponse, error)

	// The GetConfigEntry method retrieves the CAs that persisted the configuration data
	// listed in section 3.1.1.10. Configuration data is represented as a hierarchical data
	// structure with the following format: [\pwszAuthority][\pwszNodePath][\pwszEntry].
	GetConfigEntry(context.Context, *GetConfigEntryRequest, ...dcerpc.CallOption) (*GetConfigEntryResponse, error)

	// The SetConfigEntry method is used to set the CA's persisted configuration data that
	// is listed in section 3.1.1.10.
	SetConfigEntry(context.Context, *SetConfigEntryRequest, ...dcerpc.CallOption) (*SetConfigEntryResponse, error)

	// The ImportKey method adds an encrypted key set to an item in the CA database.
	ImportKey(context.Context, *ImportKeyRequest, ...dcerpc.CallOption) (*ImportKeyResponse, error)

	// The GetMyRoles method retrieves the CA roles, as specified in [CIMC-PP], assigned
	// to the user who calls the method.
	GetMyRoles(context.Context, *GetMyRolesRequest, ...dcerpc.CallOption) (*GetMyRolesResponse, error)

	// The DeleteRow method deletes a row or set of rows from a database table.
	DeleteRow(context.Context, *DeleteRowRequest, ...dcerpc.CallOption) (*DeleteRowResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CertAdminD2Client
}

type xxx_DefaultCertAdminD2Client struct {
	icertadmind.CertAdminDClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCertAdminD2Client) CertAdminD() icertadmind.CertAdminDClient {
	return o.CertAdminDClient
}

func (o *xxx_DefaultCertAdminD2Client) PublishCRLs(ctx context.Context, in *PublishCRLsRequest, opts ...dcerpc.CallOption) (*PublishCRLsResponse, error) {
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
	out := &PublishCRLsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) GetCAProperty(ctx context.Context, in *GetCAPropertyRequest, opts ...dcerpc.CallOption) (*GetCAPropertyResponse, error) {
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
	out := &GetCAPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) SetCAProperty(ctx context.Context, in *SetCAPropertyRequest, opts ...dcerpc.CallOption) (*SetCAPropertyResponse, error) {
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
	out := &SetCAPropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) GetCAPropertyInfo(ctx context.Context, in *GetCAPropertyInfoRequest, opts ...dcerpc.CallOption) (*GetCAPropertyInfoResponse, error) {
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
	out := &GetCAPropertyInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) EnumViewColumnTable(ctx context.Context, in *EnumViewColumnTableRequest, opts ...dcerpc.CallOption) (*EnumViewColumnTableResponse, error) {
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
	out := &EnumViewColumnTableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) GetCASecurity(ctx context.Context, in *GetCASecurityRequest, opts ...dcerpc.CallOption) (*GetCASecurityResponse, error) {
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
	out := &GetCASecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) SetCASecurity(ctx context.Context, in *SetCASecurityRequest, opts ...dcerpc.CallOption) (*SetCASecurityResponse, error) {
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
	out := &SetCASecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) Ping2(ctx context.Context, in *Ping2Request, opts ...dcerpc.CallOption) (*Ping2Response, error) {
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
	out := &Ping2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) GetArchivedKey(ctx context.Context, in *GetArchivedKeyRequest, opts ...dcerpc.CallOption) (*GetArchivedKeyResponse, error) {
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
	out := &GetArchivedKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) GetAuditFilter(ctx context.Context, in *GetAuditFilterRequest, opts ...dcerpc.CallOption) (*GetAuditFilterResponse, error) {
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
	out := &GetAuditFilterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) SetAuditFilter(ctx context.Context, in *SetAuditFilterRequest, opts ...dcerpc.CallOption) (*SetAuditFilterResponse, error) {
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
	out := &SetAuditFilterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) GetOfficerRights(ctx context.Context, in *GetOfficerRightsRequest, opts ...dcerpc.CallOption) (*GetOfficerRightsResponse, error) {
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
	out := &GetOfficerRightsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) SetOfficerRights(ctx context.Context, in *SetOfficerRightsRequest, opts ...dcerpc.CallOption) (*SetOfficerRightsResponse, error) {
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
	out := &SetOfficerRightsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) GetConfigEntry(ctx context.Context, in *GetConfigEntryRequest, opts ...dcerpc.CallOption) (*GetConfigEntryResponse, error) {
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
	out := &GetConfigEntryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) SetConfigEntry(ctx context.Context, in *SetConfigEntryRequest, opts ...dcerpc.CallOption) (*SetConfigEntryResponse, error) {
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
	out := &SetConfigEntryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) ImportKey(ctx context.Context, in *ImportKeyRequest, opts ...dcerpc.CallOption) (*ImportKeyResponse, error) {
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
	out := &ImportKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) GetMyRoles(ctx context.Context, in *GetMyRolesRequest, opts ...dcerpc.CallOption) (*GetMyRolesResponse, error) {
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
	out := &GetMyRolesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) DeleteRow(ctx context.Context, in *DeleteRowRequest, opts ...dcerpc.CallOption) (*DeleteRowResponse, error) {
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
	out := &DeleteRowResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertAdminD2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCertAdminD2Client) IPID(ctx context.Context, ipid *dcom.IPID) CertAdminD2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCertAdminD2Client{
		CertAdminDClient: o.CertAdminDClient.IPID(ctx, ipid),
		cc:               o.cc,
		ipid:             ipid,
	}
}
func NewCertAdminD2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CertAdminD2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CertAdminD2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := icertadmind.NewCertAdminDClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultCertAdminD2Client{
		CertAdminDClient: base,
		cc:               cc,
		ipid:             ipid,
	}, nil
}

// xxx_PublishCRLsOperation structure represents the PublishCRLs operation
type xxx_PublishCRLsOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	FileTime  *dtyp.Filetime `idl:"name:FileTime" json:"file_time"`
	Flags     uint32         `idl:"name:Flags" json:"flags"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PublishCRLsOperation) OpNum() int { return 31 }

func (o *xxx_PublishCRLsOperation) OpName() string { return "/ICertAdminD2/v0/PublishCRLs" }

func (o *xxx_PublishCRLsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PublishCRLsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PublishCRLsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PublishCRLsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PublishCRLsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PublishCRLsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PublishCRLsRequest structure represents the PublishCRLs operation request
type PublishCRLsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the definition of the pwszAuthority parameter in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// FileTime: Contains a 64-bit value that represents the number of 100-nanosecond intervals
	// since January 1, 1601 (UTC). Specifies the nextUpdate value of the CRL, as specified
	// in [RFC3280] section 5.1.2.5, in Greenwich Mean Time.
	FileTime *dtyp.Filetime `idl:"name:FileTime" json:"file_time"`
	// Flags: An unsigned integer value that specifies the type of CRL to publish and the
	// publishing parameters. This parameter MUST be set to a combination of the following
	// values. Flags uses B as the least-significant bit. It uses B, D and F as shown in
	// the following table.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| B | D | 0 | 0 | F | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	//
	//	+-------+------------------------------------------------+
	//	|       |                                                |
	//	| VALUE |                  DESCRIPTION                   |
	//	|       |                                                |
	//	+-------+------------------------------------------------+
	//	+-------+------------------------------------------------+
	//	| B     | If 1, the CA MUST publish a base CRL.          |
	//	+-------+------------------------------------------------+
	//	| D     | If 1, the CA MUST publish a delta CRL.         |
	//	+-------+------------------------------------------------+
	//	| F     | If 1, the CA MUST republish the existing CRLs. |
	//	+-------+------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *PublishCRLsRequest) xxx_ToOp(ctx context.Context) *xxx_PublishCRLsOperation {
	if o == nil {
		return &xxx_PublishCRLsOperation{}
	}
	return &xxx_PublishCRLsOperation{
		This:      o.This,
		Authority: o.Authority,
		FileTime:  o.FileTime,
		Flags:     o.Flags,
	}
}

func (o *PublishCRLsRequest) xxx_FromOp(ctx context.Context, op *xxx_PublishCRLsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.FileTime = op.FileTime
	o.Flags = op.Flags
}
func (o *PublishCRLsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PublishCRLsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PublishCRLsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PublishCRLsResponse structure represents the PublishCRLs operation response
type PublishCRLsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PublishCRLs return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PublishCRLsResponse) xxx_ToOp(ctx context.Context) *xxx_PublishCRLsOperation {
	if o == nil {
		return &xxx_PublishCRLsOperation{}
	}
	return &xxx_PublishCRLsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PublishCRLsResponse) xxx_FromOp(ctx context.Context, op *xxx_PublishCRLsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PublishCRLsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PublishCRLsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PublishCRLsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCAPropertyOperation structure represents the GetCAProperty operation
type xxx_GetCAPropertyOperation struct {
	This          *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority     string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	PropertyID    int32                   `idl:"name:PropId" json:"property_id"`
	PropertyIndex int32                   `idl:"name:PropIndex" json:"property_index"`
	PropertyType  int32                   `idl:"name:PropType" json:"property_type"`
	PropertyValue *csra.CertTransportBlob `idl:"name:pctbPropertyValue;pointer:ref" json:"property_value"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCAPropertyOperation) OpNum() int { return 32 }

func (o *xxx_GetCAPropertyOperation) OpName() string { return "/ICertAdminD2/v0/GetCAProperty" }

func (o *xxx_GetCAPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// PropId {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.PropertyID); err != nil {
			return err
		}
	}
	// PropIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.PropertyIndex); err != nil {
			return err
		}
	}
	// PropType {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.PropertyType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// PropId {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.PropertyID); err != nil {
			return err
		}
	}
	// PropIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.PropertyIndex); err != nil {
			return err
		}
	}
	// PropType {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.PropertyType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pctbPropertyValue {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.PropertyValue != nil {
			if err := o.PropertyValue.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetCAPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pctbPropertyValue {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.PropertyValue == nil {
			o.PropertyValue = &csra.CertTransportBlob{}
		}
		if err := o.PropertyValue.UnmarshalNDR(ctx, w); err != nil {
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

// GetCAPropertyRequest structure represents the GetCAProperty operation request
type GetCAPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// PropId: An integer value specifying the property to be returned. The PropID value
	// MUST be one of the values in the table labeled PropId in [MS-WCCE] section 3.2.1.4.3.2.
	// If a value other than one of the listed values is used, the error E_INVALIDARG is
	// returned.
	PropertyID int32 `idl:"name:PropId" json:"property_id"`
	// PropIndex: Some of these properties (the ones labeled "indexed" in the table in [MS-WCCE]
	// section 3.2.1.4.3.2) have arrays of values. This parameter MUST be used as the index
	// into such an array. For properties that are not arrays, this parameter MUST be ignored.
	PropertyIndex int32 `idl:"name:PropIndex" json:"property_index"`
	// PropType: An integer value that specifies the property data type.
	//
	//	+----------------------------+-------------------------------------------------------------+
	//	|                            |                                                             |
	//	|           VALUE            |                           MEANING                           |
	//	|                            |                                                             |
	//	+----------------------------+-------------------------------------------------------------+
	//	+----------------------------+-------------------------------------------------------------+
	//	| PROPTYPE_LONG 0x00000001   | The property type is a signed long integer or a byte array. |
	//	+----------------------------+-------------------------------------------------------------+
	//	| PROPTYPE_BINARY 0x00000003 | The property type is binary data.                           |
	//	+----------------------------+-------------------------------------------------------------+
	//	| PROPTYPE_STRING 0x00000004 | The property type is a Unicode string.                      |
	//	+----------------------------+-------------------------------------------------------------+
	PropertyType int32 `idl:"name:PropType" json:"property_type"`
}

func (o *GetCAPropertyRequest) xxx_ToOp(ctx context.Context) *xxx_GetCAPropertyOperation {
	if o == nil {
		return &xxx_GetCAPropertyOperation{}
	}
	return &xxx_GetCAPropertyOperation{
		This:          o.This,
		Authority:     o.Authority,
		PropertyID:    o.PropertyID,
		PropertyIndex: o.PropertyIndex,
		PropertyType:  o.PropertyType,
	}
}

func (o *GetCAPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCAPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.PropertyID = op.PropertyID
	o.PropertyIndex = op.PropertyIndex
	o.PropertyType = op.PropertyType
}
func (o *GetCAPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCAPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCAPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCAPropertyResponse structure represents the GetCAProperty operation response
type GetCAPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pctbPropertyValue:  If the function succeeds, this method MUST return a CERTTRANSBLOB
	// structure that contains the property value. If the function fails, the contents are
	// undefined.
	//
	// Note  The numeric values for the constants listed in this topic are defined in the
	// table for the PropID parameter.
	//
	// * If PROPTYPE_STRING is specified in the PropType parameter, pctbPropertyValue MUST
	// be a pointer to a CERTTRANSBLOB structure. The *pb* member of the structure points
	// to the little-endian ( c6451297-197d-4b4b-b786-3f3187b67b8f#gt_079478cb-f4c5-4ce5-b72b-2144da5d2ce7
	// ) encoded Unicode string. The length, in bytes, of the string MUST be contained in
	// the *cb* member.
	//
	// * If PROPTYPE_LONG is specified in the PropType parameter, there are two possible
	// return types depending on the PropID. The first type is the return of a CAINFO structure
	// (as specified in [MS-WCCE] section 2.2.2.4 ( ../ms-wcce/4fa5241c-d10e-4011-87e0-c74753d725a3
	// ) ) and the second type is for the return of a BYTE array:
	//
	// * If the value passed in PropId maps to one of the following properties, pctbPropertyValue
	// is a pointer to a CERTTRANSBLOB structure, and the *pb* member of that structure
	// MUST contain a pointer to a CAINFO structure that contains the values of the properties
	// listed as follows. The marshaling rules for a CAINFO structure in a CERTTRANSBLOB
	// are specified in [MS-WCCE] section 2.2.2.2.5 ( ../ms-wcce/cd9656c0-6be3-4887-84b0-aacedc017b0b
	// ) :
	//
	// * CR_PROP_CATYPE
	//
	// * CR_PROP_CASIGCERTCOUNT
	//
	// * CR_PROP_CAXCHGCERTCOUNT
	//
	// * CR_PROP_EXITCOUNT
	//
	// * CR_PROP_CAPROPIDMAX
	//
	// * CR_PROP_KRACERTUSEDCOUNT
	//
	// * CR_PROP_ROLESEPARATIONENABLED
	//
	// * CR_PROP_KRACERTCOUNT
	//
	// * CR_PROP_ADVANCEDSERVER
	//
	// * If the value passed in PropId maps to one of the following properties, pctbPropertyValue
	// is a pointer to a CERTTRANSBLOB structure, and the *pb* member of the structure points
	// to a byte array containing the value for the requested property. The marshaling rules
	// for each property are specified in the subsection of [MS-WCCE] section 3.2.1.4.3.2
	// that corresponds to the property name. The *cb* member contains the length of the
	// byte array:
	//
	// * CR_PROP_CACERTSTATE
	//
	// * CR_PROP_CRLSTATE
	//
	// * CR_PROP_KRACERTSTATE
	//
	// * CR_PROP_BASECRLPUBLISHSTATE
	//
	// * CR_PROP_DELTACRLPUBLISHSTATE
	//
	// * CR_PROP_CACERTSTATUSCODE
	//
	// * CR_PROP_CAFORWARDCROSSCERTSTATE
	//
	// * CR_PROP_CABACKWARDCROSSCERTSTATE
	//
	// * If PROPTYPE_BINARY is specified in the PropType parameter, pctbPropertyValue MUST
	// be a pointer to a CERTTRANSBLOB structure. The *pb* member of the structure points
	// to the requested binary large object (BLOB).
	//
	// Based on the property identifier passed in PropId , the binary data pointed to by
	// the *pb* member MUST be populated as follows:
	//
	// * CR_PROP_CASIGCERT: MUST be an X.509 certificate ( c6451297-197d-4b4b-b786-3f3187b67b8f#gt_7a0f4b71-23ba-434f-b781-28053ed64879
	// ) encoded using DER, as specified in [X660] ( https://go.microsoft.com/fwlink/?LinkId=90592
	// ).
	//
	// * CR_PROP_BASECRL: MUST be a X.509 CRL ( c6451297-197d-4b4b-b786-3f3187b67b8f#gt_4f22841f-249b-42fb-a31a-5049c00be939
	// ) encoded using DER, as specified in [X660].
	//
	// * CR_PROP_CAFORWARDCROSSCERT: MUST be a X.509 certificate encoded using DER, as specified
	// in [X660].
	//
	// * CR_PROP_CABACKWARDCROSSCERT: MUST be a X.509 certificate encoded using DER, as
	// specified in [X660].
	//
	// * CR_PROP_CAXCHGCERT: MUST be a X.509 certificate encoded using DER, as specified
	// in [X660].
	//
	// The CA MUST execute the processing rules specified in [MS-WCCE] section 3.2.1.4.3.2.15
	// ( ../ms-wcce/585d9359-4bc5-471f-bba7-2d9a336debdc ) , "PropID = 0x0000000F (CR_PROP_CAXCHGCERT)
	// "CA Exchange Certificate"".
	//
	// * CR_PROP_CAXCHGCERTCHAIN: MUST be a CMS message, as specified in [RFC2797] ( https://go.microsoft.com/fwlink/?LinkId=90382
	// ) encoded using DER, as specified in [X660].
	//
	// The CA MUST execute the processing rules specified in [MS-WCCE] section 3.2.1.4.3.2.16
	// ( ../ms-wcce/d38e7259-d0d6-4adb-b111-2bac47c64bed ) , "PropID = 0x00000010 (CR_PROP_CAXCHGCERTCHAIN)
	// "CA Exchange Certificate Chain"".
	//
	// * CR_PROP_CASIGCERTCHAIN: MUST be a CMS message [RFC2797] encoded using DER. [X660].
	//
	// * CR_PROP_CASIGCERTCRLCHAIN: MUST be a CMS message, as specified in [RFC2797], encoded
	// using DER, as specified in [X660].
	//
	// * CR_PROP_CASIGCERTCRLCHAIN: MUST be a CMS message, as specified in [RFC2797], encoded
	// using DER, as specified in [X660].
	//
	// * CR_PROP_CAXCHGCERTCRLCHAIN: CR_PROP_CASIGCERTCRLCHAIN: MUST be a CMS message, as
	// specified in [RFC2797], encoded using DER, as specified in [X660].
	//
	// The CA MUST execute the processing rules specified in [MS-WCCE] section 3.2.1.4.3.2.33
	// ( ../ms-wcce/5dbf4c4f-3ac1-426a-9425-ff96968a1b6a ) , "PropID = 0x00000021 (CR_PROP_CAXCHGCERTCRLCHAIN)
	// "CA Exchange Certificate Chain and CRL"".
	//
	// * CR_PROP_DELTACRL: MUST be a X.509 CRL encoded using DER [X660].
	//
	// * CR_PROP_KRACERT: MUST be a X.509 CRL encoded using DER, as specified in [X660].
	//
	// The marshaling rules for each of the preceding properties into a CERTTRANSBLOB are
	// specified in [MS-WCCE] sections 2.2.2.2.2 (for X.509 certificate), 2.2.2.2.3 (X.509
	// CRL), and 2.2.2.2.4 (CMS message).
	PropertyValue *csra.CertTransportBlob `idl:"name:pctbPropertyValue;pointer:ref" json:"property_value"`
	// Return: The GetCAProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCAPropertyResponse) xxx_ToOp(ctx context.Context) *xxx_GetCAPropertyOperation {
	if o == nil {
		return &xxx_GetCAPropertyOperation{}
	}
	return &xxx_GetCAPropertyOperation{
		That:          o.That,
		PropertyValue: o.PropertyValue,
		Return:        o.Return,
	}
}

func (o *GetCAPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCAPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyValue = op.PropertyValue
	o.Return = op.Return
}
func (o *GetCAPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCAPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCAPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCAPropertyOperation structure represents the SetCAProperty operation
type xxx_SetCAPropertyOperation struct {
	This          *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority     string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	PropertyID    int32                   `idl:"name:PropId" json:"property_id"`
	PropertyIndex int32                   `idl:"name:PropIndex" json:"property_index"`
	PropertyType  int32                   `idl:"name:PropType" json:"property_type"`
	PropertyValue *csra.CertTransportBlob `idl:"name:pctbPropertyValue" json:"property_value"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCAPropertyOperation) OpNum() int { return 33 }

func (o *xxx_SetCAPropertyOperation) OpName() string { return "/ICertAdminD2/v0/SetCAProperty" }

func (o *xxx_SetCAPropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCAPropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// PropId {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.PropertyID); err != nil {
			return err
		}
	}
	// PropIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.PropertyIndex); err != nil {
			return err
		}
	}
	// PropType {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.PropertyType); err != nil {
			return err
		}
	}
	// pctbPropertyValue {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.PropertyValue != nil {
			if err := o.PropertyValue.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetCAPropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// PropId {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.PropertyID); err != nil {
			return err
		}
	}
	// PropIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.PropertyIndex); err != nil {
			return err
		}
	}
	// PropType {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.PropertyType); err != nil {
			return err
		}
	}
	// pctbPropertyValue {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.PropertyValue == nil {
			o.PropertyValue = &csra.CertTransportBlob{}
		}
		if err := o.PropertyValue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCAPropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCAPropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCAPropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCAPropertyRequest structure represents the SetCAProperty operation request
type SetCAPropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// PropId: A LONG value that specifies one and exactly one of the following property
	// identifiers. The use of PropIds, is as specified in [MS-WCCE] section 3.2.1.4.3.2.
	// If a value other than one of the listed values is used, the error E_INVALIDARG is
	// returned.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x0000001a | A binary object that contains the CA's key recovery agent (KRA) certificate to   |
	//	|            | be added at the index specified by PropIndex parameter.                          |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000019 | The maximum number of KRA certificates available on the CA.                      |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000018 | The minimum number of KRAs to use when archiving a private key. For more         |
	//	|            | information on KRA usage, see [MSFT-ARCHIVE].                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x0000001d | A collection of name and OID (1) pairs that identify the templates supported by  |
	//	|            | a CA.                                                                            |
	//	+------------+----------------------------------------------------------------------------------+
	PropertyID int32 `idl:"name:PropId" json:"property_id"`
	// PropIndex: A LONG value for the index of the KRA certificate to set when the provided
	// PropId is 0x1a. For other PropId values, it MUST be 0.
	PropertyIndex int32 `idl:"name:PropIndex" json:"property_index"`
	// PropType: A LONG value that specifies the type of the property. This parameter MUST
	// be one of the following values.
	//
	//	+----------------------------+---------------------+
	//	|                            |                     |
	//	|           VALUE            |       MEANING       |
	//	|                            |                     |
	//	+----------------------------+---------------------+
	//	+----------------------------+---------------------+
	//	| PROPTYPE_LONG 0x00000001   | Signed LONG data    |
	//	+----------------------------+---------------------+
	//	| PROPTYPE_BINARY 0x00000003 | Binary data         |
	//	+----------------------------+---------------------+
	//	| PROPTYPE_STRING 0x00000004 | Unicode String data |
	//	+----------------------------+---------------------+
	PropertyType int32 `idl:"name:PropType" json:"property_type"`
	// pctbPropertyValue: A pointer to CERTTRANSBLOB that specifies the new property value.
	// The format for the value contained in CERTTRANSBLOB is specific to the PropId defined
	// as follows.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|              VALUE OF               |                               FORMAT FOR VALUES IN                               |
	//	|               PROPID                |                                  CERTTRANSBLOB                                   |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_KRACERTUSEDCOUNT 0x00000018 | The pb member of CERTTRANSBLOB MUST point to an unsigned integer value           |
	//	|                                     | (little-endian format) and the cb member of CERTTRANSBLOB MUST contain the       |
	//	|                                     | length of the bytes containing the value.                                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_KRACERTCOUNT 0x00000019     | The pb member of CERTTRANSBLOB MUST point to an unsigned integer value           |
	//	|                                     | (little-endian format) and the cb member of CERTTRANSBLOB MUST contain the       |
	//	|                                     | length of the bytes containing the value.                                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_KRACERT 0x0000001a          | The pb member of CERTTRANSBLOB MUST point to an ASN.1 DER (as specified in       |
	//	|                                     | [ITUX690]) encoded byte array of Certificate. The cb member of CERTTRANSBLOB     |
	//	|                                     | MUST contain the length of the array.                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| CR_PROP_TEMPLATES 0x0000001d        | As specified in [MS-WCCE] section 3.2.1.4.3.2.29.                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// The following  table defines the values that MUST be set for PropIndex and PropType
	// for each one of the property values passed via PropID.
	//
	//	+--------------+-------------------------+------------------+
	//	|    PROPID    |     PROPINDEX MUST      |  PROPTYPE MUST   |
	//	|    VALUE     |           BE            |        BE        |
	//	+--------------+-------------------------+------------------+
	//	+--------------+-------------------------+------------------+
	//	| 0x0000001a   | The minimum index is 0. | 0x00000003       |
	//	+--------------+-------------------------+------------------+
	//	| 0x00000019   | 0x00000000              | 0x00000001       |
	//	+--------------+-------------------------+------------------+
	//	| 0x00000018   | 0x00000000              | 0x00000001       |
	//	+--------------+-------------------------+------------------+
	//	| 0x0000001d   | 0x00000000              | 0x00000004       |
	//	+--------------+-------------------------+------------------+
	PropertyValue *csra.CertTransportBlob `idl:"name:pctbPropertyValue" json:"property_value"`
}

func (o *SetCAPropertyRequest) xxx_ToOp(ctx context.Context) *xxx_SetCAPropertyOperation {
	if o == nil {
		return &xxx_SetCAPropertyOperation{}
	}
	return &xxx_SetCAPropertyOperation{
		This:          o.This,
		Authority:     o.Authority,
		PropertyID:    o.PropertyID,
		PropertyIndex: o.PropertyIndex,
		PropertyType:  o.PropertyType,
		PropertyValue: o.PropertyValue,
	}
}

func (o *SetCAPropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCAPropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.PropertyID = op.PropertyID
	o.PropertyIndex = op.PropertyIndex
	o.PropertyType = op.PropertyType
	o.PropertyValue = op.PropertyValue
}
func (o *SetCAPropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetCAPropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCAPropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCAPropertyResponse structure represents the SetCAProperty operation response
type SetCAPropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetCAProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCAPropertyResponse) xxx_ToOp(ctx context.Context) *xxx_SetCAPropertyOperation {
	if o == nil {
		return &xxx_SetCAPropertyOperation{}
	}
	return &xxx_SetCAPropertyOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetCAPropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCAPropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCAPropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetCAPropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCAPropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCAPropertyInfoOperation structure represents the GetCAPropertyInfo operation
type xxx_GetCAPropertyInfoOperation struct {
	This          *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority     string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	PropertyCount int32                   `idl:"name:pcProperty" json:"property_count"`
	PropertyInfo  *csra.CertTransportBlob `idl:"name:pctbPropInfo;pointer:ref" json:"property_info"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCAPropertyInfoOperation) OpNum() int { return 34 }

func (o *xxx_GetCAPropertyInfoOperation) OpName() string { return "/ICertAdminD2/v0/GetCAPropertyInfo" }

func (o *xxx_GetCAPropertyInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAPropertyInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCAPropertyInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCAPropertyInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCAPropertyInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcProperty {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.PropertyCount); err != nil {
			return err
		}
	}
	// pctbPropInfo {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.PropertyInfo != nil {
			if err := o.PropertyInfo.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetCAPropertyInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcProperty {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.PropertyCount); err != nil {
			return err
		}
	}
	// pctbPropInfo {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.PropertyInfo == nil {
			o.PropertyInfo = &csra.CertTransportBlob{}
		}
		if err := o.PropertyInfo.UnmarshalNDR(ctx, w); err != nil {
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

// GetCAPropertyInfoRequest structure represents the GetCAPropertyInfo operation request
type GetCAPropertyInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the pwszAuthority definition in ICertAdminD::SetExtension.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *GetCAPropertyInfoRequest) xxx_ToOp(ctx context.Context) *xxx_GetCAPropertyInfoOperation {
	if o == nil {
		return &xxx_GetCAPropertyInfoOperation{}
	}
	return &xxx_GetCAPropertyInfoOperation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *GetCAPropertyInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCAPropertyInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *GetCAPropertyInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCAPropertyInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCAPropertyInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCAPropertyInfoResponse structure represents the GetCAPropertyInfo operation response
type GetCAPropertyInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcProperty: An integer value containing the number of property structures returned.
	PropertyCount int32 `idl:"name:pcProperty" json:"property_count"`
	// pctbPropInfo: A CERTTRANSBLOB structure containing zero or more CATRANSPROP structures.
	// For more information on CERTTRANSBLOB and CATRANSPROP structures, see section 2.2.1.
	//
	// The processing of the ICertAdminD2::GetCAPropertyInfo method is the same as that
	// specified in [MS-WCCE] section 3.2.1.4.3.3.
	PropertyInfo *csra.CertTransportBlob `idl:"name:pctbPropInfo;pointer:ref" json:"property_info"`
	// Return: The GetCAPropertyInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCAPropertyInfoResponse) xxx_ToOp(ctx context.Context) *xxx_GetCAPropertyInfoOperation {
	if o == nil {
		return &xxx_GetCAPropertyInfoOperation{}
	}
	return &xxx_GetCAPropertyInfoOperation{
		That:          o.That,
		PropertyCount: o.PropertyCount,
		PropertyInfo:  o.PropertyInfo,
		Return:        o.Return,
	}
}

func (o *GetCAPropertyInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCAPropertyInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyCount = op.PropertyCount
	o.PropertyInfo = op.PropertyInfo
	o.Return = op.Return
}
func (o *GetCAPropertyInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCAPropertyInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCAPropertyInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumViewColumnTableOperation structure represents the EnumViewColumnTable operation
type xxx_EnumViewColumnTableOperation struct {
	This           *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority      string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Table          uint32                  `idl:"name:iTable" json:"table"`
	Column         uint32                  `idl:"name:iColumn" json:"column"`
	ColumnCount    uint32                  `idl:"name:cColumn" json:"column_count"`
	ColumnOutCount uint32                  `idl:"name:pcColumnOut" json:"column_out_count"`
	ColumnInfo     *csra.CertTransportBlob `idl:"name:pctbColumnInfo;pointer:ref" json:"column_info"`
	Return         int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumViewColumnTableOperation) OpNum() int { return 35 }

func (o *xxx_EnumViewColumnTableOperation) OpName() string {
	return "/ICertAdminD2/v0/EnumViewColumnTable"
}

func (o *xxx_EnumViewColumnTableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewColumnTableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// iTable {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Table); err != nil {
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

func (o *xxx_EnumViewColumnTableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// iTable {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Table); err != nil {
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

func (o *xxx_EnumViewColumnTableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumViewColumnTableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnumViewColumnTableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// EnumViewColumnTableRequest structure represents the EnumViewColumnTable operation request
type EnumViewColumnTableRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the definition of the pwszAuthority parameter in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// iTable: An unsigned integer that specifies the database table to be used for the
	// enumeration. This MUST be set from the following values.
	//
	//	+------------+-----------------+
	//	|            |                 |
	//	|   VALUE    |     MEANING     |
	//	|            |                 |
	//	+------------+-----------------+
	//	+------------+-----------------+
	//	| 0x00000000 | Request table   |
	//	+------------+-----------------+
	//	| 0x00003000 | Extension table |
	//	+------------+-----------------+
	//	| 0x00004000 | Attribute table |
	//	+------------+-----------------+
	//	| 0x00005000 | CRL table       |
	//	+------------+-----------------+
	Table uint32 `idl:"name:iTable" json:"table"`
	// iColumn: An unsigned integer that specifies the column number with which to begin
	// the enumeration. Valid values are from 0 to one less than the maximum number of columns
	// for the table.
	Column uint32 `idl:"name:iColumn" json:"column"`
	// cColumn: An unsigned integer that specifies the requested number of columns to return.
	ColumnCount uint32 `idl:"name:cColumn" json:"column_count"`
}

func (o *EnumViewColumnTableRequest) xxx_ToOp(ctx context.Context) *xxx_EnumViewColumnTableOperation {
	if o == nil {
		return &xxx_EnumViewColumnTableOperation{}
	}
	return &xxx_EnumViewColumnTableOperation{
		This:        o.This,
		Authority:   o.Authority,
		Table:       o.Table,
		Column:      o.Column,
		ColumnCount: o.ColumnCount,
	}
}

func (o *EnumViewColumnTableRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumViewColumnTableOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.Table = op.Table
	o.Column = op.Column
	o.ColumnCount = op.ColumnCount
}
func (o *EnumViewColumnTableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumViewColumnTableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumViewColumnTableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumViewColumnTableResponse structure represents the EnumViewColumnTable operation response
type EnumViewColumnTableResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	ColumnOutCount uint32         `idl:"name:pcColumnOut" json:"column_out_count"`
	// pctbColumnInfo: A pointer to a CERTTRANSBLOB structure. Upon return, the pb member
	// of this structure points to an array of the marshaled CERTTRANSDBCOLUMN structures.
	// The format and marshaling for the value of pctbColumnInfo MUST be as specified in
	// section 2.2.1.7.
	//
	// The EnumViewColumnTable method returns information to the client about columns that
	// are associated with a specific table. The CA server MUST enforce the following processing
	// rules:
	//
	// * The CA server MUST enforce that the iTable parameter has a value as specified in
	// the previous table; otherwise, it MUST fail with the error ERROR_INVALID_PARAMETER.
	//
	// * The CA server MUST enforce that iColumn is less than the number of columns associated
	// with the table; otherwise, it MUST fail with the error ERROR_ARITHMETIC_OVERFLOW.
	//
	// * The CA server MUST enforce that cColumn is greater than 0; otherwise, it MUST fail
	// with the error ERROR_INVALID_PARAMETER. <71> ( 5f06c74c-1a29-4fdf-b8dd-ae3300d1b90d#Appendix_A_71
	// )
	//
	// * The CA server MUST use the value of *iColumn* to identify the column identifier
	// that is associated with the table (identified by the value of the iTable parameter).
	//
	// * The number of column information returned MUST be a minimum of the *cColumn* value
	// and the remaining number of columns in the table (starting from *iColumn* ). The
	// value of *pcColumn MUST be set to the number of the column information returned.
	ColumnInfo *csra.CertTransportBlob `idl:"name:pctbColumnInfo;pointer:ref" json:"column_info"`
	// Return: The EnumViewColumnTable return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumViewColumnTableResponse) xxx_ToOp(ctx context.Context) *xxx_EnumViewColumnTableOperation {
	if o == nil {
		return &xxx_EnumViewColumnTableOperation{}
	}
	return &xxx_EnumViewColumnTableOperation{
		That:           o.That,
		ColumnOutCount: o.ColumnOutCount,
		ColumnInfo:     o.ColumnInfo,
		Return:         o.Return,
	}
}

func (o *EnumViewColumnTableResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumViewColumnTableOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ColumnOutCount = op.ColumnOutCount
	o.ColumnInfo = op.ColumnInfo
	o.Return = op.Return
}
func (o *EnumViewColumnTableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumViewColumnTableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumViewColumnTableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCASecurityOperation structure represents the GetCASecurity operation
type xxx_GetCASecurityOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	SD        *csra.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCASecurityOperation) OpNum() int { return 36 }

func (o *xxx_GetCASecurityOperation) OpName() string { return "/ICertAdminD2/v0/GetCASecurity" }

func (o *xxx_GetCASecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCASecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCASecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCASecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCASecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pctbSD {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD != nil {
			if err := o.SD.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetCASecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pctbSD {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD == nil {
			o.SD = &csra.CertTransportBlob{}
		}
		if err := o.SD.UnmarshalNDR(ctx, w); err != nil {
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

// GetCASecurityRequest structure represents the GetCASecurity operation request
type GetCASecurityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *GetCASecurityRequest) xxx_ToOp(ctx context.Context) *xxx_GetCASecurityOperation {
	if o == nil {
		return &xxx_GetCASecurityOperation{}
	}
	return &xxx_GetCASecurityOperation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *GetCASecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCASecurityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *GetCASecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCASecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCASecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCASecurityResponse structure represents the GetCASecurity operation response
type GetCASecurityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pctbSD: A pointer to a CERTTRANSBLOB data structure that contains the CA's security
	// descriptor. Security descriptors are specified in [MS-DTYP] section 2.4.6.
	SD *csra.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
	// Return: The GetCASecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCASecurityResponse) xxx_ToOp(ctx context.Context) *xxx_GetCASecurityOperation {
	if o == nil {
		return &xxx_GetCASecurityOperation{}
	}
	return &xxx_GetCASecurityOperation{
		That:   o.That,
		SD:     o.SD,
		Return: o.Return,
	}
}

func (o *GetCASecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCASecurityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SD = op.SD
	o.Return = op.Return
}
func (o *GetCASecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCASecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCASecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCASecurityOperation structure represents the SetCASecurity operation
type xxx_SetCASecurityOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	SD        *csra.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCASecurityOperation) OpNum() int { return 37 }

func (o *xxx_SetCASecurityOperation) OpName() string { return "/ICertAdminD2/v0/SetCASecurity" }

func (o *xxx_SetCASecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCASecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pctbSD {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD != nil {
			if err := o.SD.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetCASecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pctbSD {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD == nil {
			o.SD = &csra.CertTransportBlob{}
		}
		if err := o.SD.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCASecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCASecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCASecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCASecurityRequest structure represents the SetCASecurity operation request
type SetCASecurityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// pctbSD: A pointer to a CERTTRANSBLOB data structure that holds the security descriptor.
	// Security descriptors are specified in [MS-DTYP] section 2.4.6.
	//
	// The CA SHOULD use the permissions set in pctbSD to deny and allow permissions to
	// CA functionality. Microsoft CA permissions are defined in section 3.1.1.7.
	SD *csra.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
}

func (o *SetCASecurityRequest) xxx_ToOp(ctx context.Context) *xxx_SetCASecurityOperation {
	if o == nil {
		return &xxx_SetCASecurityOperation{}
	}
	return &xxx_SetCASecurityOperation{
		This:      o.This,
		Authority: o.Authority,
		SD:        o.SD,
	}
}

func (o *SetCASecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCASecurityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.SD = op.SD
}
func (o *SetCASecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetCASecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCASecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCASecurityResponse structure represents the SetCASecurity operation response
type SetCASecurityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetCASecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCASecurityResponse) xxx_ToOp(ctx context.Context) *xxx_SetCASecurityOperation {
	if o == nil {
		return &xxx_SetCASecurityOperation{}
	}
	return &xxx_SetCASecurityOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetCASecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCASecurityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCASecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetCASecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCASecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Ping2Operation structure represents the Ping2 operation
type xxx_Ping2Operation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_Ping2Operation) OpNum() int { return 38 }

func (o *xxx_Ping2Operation) OpName() string { return "/ICertAdminD2/v0/Ping2" }

func (o *xxx_Ping2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Ping2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_Ping2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_Ping2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Ping2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_Ping2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// Ping2Request structure represents the Ping2 operation request
type Ping2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the pwszAuthority definition in section 3.1.4.1.1.
	//
	// ICertAdminD2::Ping2 is as specified in [MS-WCCE] section 3.2.1.4.3.4.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *Ping2Request) xxx_ToOp(ctx context.Context) *xxx_Ping2Operation {
	if o == nil {
		return &xxx_Ping2Operation{}
	}
	return &xxx_Ping2Operation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *Ping2Request) xxx_FromOp(ctx context.Context, op *xxx_Ping2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *Ping2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Ping2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Ping2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Ping2Response structure represents the Ping2 operation response
type Ping2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Ping2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Ping2Response) xxx_ToOp(ctx context.Context) *xxx_Ping2Operation {
	if o == nil {
		return &xxx_Ping2Operation{}
	}
	return &xxx_Ping2Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *Ping2Response) xxx_FromOp(ctx context.Context, op *xxx_Ping2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *Ping2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Ping2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Ping2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetArchivedKeyOperation structure represents the GetArchivedKey operation
type xxx_GetArchivedKeyOperation struct {
	This        *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority   string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	RequestID   uint32                  `idl:"name:dwRequestId" json:"request_id"`
	ArchivedKey *csra.CertTransportBlob `idl:"name:pctbArchivedKey;pointer:ref" json:"archived_key"`
	Return      int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetArchivedKeyOperation) OpNum() int { return 39 }

func (o *xxx_GetArchivedKeyOperation) OpName() string { return "/ICertAdminD2/v0/GetArchivedKey" }

func (o *xxx_GetArchivedKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetArchivedKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetArchivedKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetArchivedKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetArchivedKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pctbArchivedKey {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.ArchivedKey != nil {
			if err := o.ArchivedKey.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetArchivedKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pctbArchivedKey {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.ArchivedKey == nil {
			o.ArchivedKey = &csra.CertTransportBlob{}
		}
		if err := o.ArchivedKey.UnmarshalNDR(ctx, w); err != nil {
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

// GetArchivedKeyRequest structure represents the GetArchivedKey operation request
type GetArchivedKeyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// dwRequestId: An unsigned integer value that specifies the RequestId of the certificate
	// request for which the archived private key and associated certificate are being requested.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
}

func (o *GetArchivedKeyRequest) xxx_ToOp(ctx context.Context) *xxx_GetArchivedKeyOperation {
	if o == nil {
		return &xxx_GetArchivedKeyOperation{}
	}
	return &xxx_GetArchivedKeyOperation{
		This:      o.This,
		Authority: o.Authority,
		RequestID: o.RequestID,
	}
}

func (o *GetArchivedKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetArchivedKeyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.RequestID = op.RequestID
}
func (o *GetArchivedKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetArchivedKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetArchivedKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetArchivedKeyResponse structure represents the GetArchivedKey operation response
type GetArchivedKeyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pctbArchivedKey: A pointer to a CERTTRANSBLOB structure that MUST contain, on successful
	// response, the archived private key and associated certificate.
	//
	// ArchivedKey Property Value Processing and Format
	ArchivedKey *csra.CertTransportBlob `idl:"name:pctbArchivedKey;pointer:ref" json:"archived_key"`
	// Return: The GetArchivedKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetArchivedKeyResponse) xxx_ToOp(ctx context.Context) *xxx_GetArchivedKeyOperation {
	if o == nil {
		return &xxx_GetArchivedKeyOperation{}
	}
	return &xxx_GetArchivedKeyOperation{
		That:        o.That,
		ArchivedKey: o.ArchivedKey,
		Return:      o.Return,
	}
}

func (o *GetArchivedKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetArchivedKeyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ArchivedKey = op.ArchivedKey
	o.Return = op.Return
}
func (o *GetArchivedKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetArchivedKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetArchivedKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAuditFilterOperation structure represents the GetAuditFilter operation
type xxx_GetAuditFilterOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Filter    uint32         `idl:"name:pdwFilter" json:"filter"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAuditFilterOperation) OpNum() int { return 40 }

func (o *xxx_GetAuditFilterOperation) OpName() string { return "/ICertAdminD2/v0/GetAuditFilter" }

func (o *xxx_GetAuditFilterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAuditFilterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAuditFilterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAuditFilterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAuditFilterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwFilter {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
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

func (o *xxx_GetAuditFilterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwFilter {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
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

// GetAuditFilterRequest structure represents the GetAuditFilter operation request
type GetAuditFilterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *GetAuditFilterRequest) xxx_ToOp(ctx context.Context) *xxx_GetAuditFilterOperation {
	if o == nil {
		return &xxx_GetAuditFilterOperation{}
	}
	return &xxx_GetAuditFilterOperation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *GetAuditFilterRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAuditFilterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *GetAuditFilterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAuditFilterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAuditFilterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAuditFilterResponse structure represents the GetAuditFilter operation response
type GetAuditFilterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwFilter: An unsigned integer that specifies the current audit settings. This is
	// a bitwise-OR combination of zero or more of the following values.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | Audit CA server for the following events: ServerControl Registration of the      |
	//	|            | ICertAdminD interface. Unregistration of the ICertAdminD interface.              |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | Audit CA server for the following method calls: BackupPrepare BackupEnd          |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 | Audit CA server for the following method calls: ICertRequestD::Request           |
	//	|            | ResubmitRequest DenyRequest SetAttributes SetExtensions ImportCertificate        |
	//	|            | DeleteRow                                                                        |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 | Audit CA server for the following method calls: RevokeCertificate PublishCRL     |
	//	|            | PublishCRLs                                                                      |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000010 | Audit CA server for the following method calls: SetCASecurity SetOfficerRights   |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000020 | Audit CA server for the following method calls: GetArchivedKey ImportKey         |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000040 | Audit CA server for the following method calls: SetCAProperty SetConfigEntry     |
	//	+------------+----------------------------------------------------------------------------------+
	//
	// The GetAuditFilter method is used to retrieve the audit filter currently in use (initialize
	// to 0 during the registration of the interfaces and can be modified by a call to the
	// SetAuditFilter method).
	Filter uint32 `idl:"name:pdwFilter" json:"filter"`
	// Return: The GetAuditFilter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAuditFilterResponse) xxx_ToOp(ctx context.Context) *xxx_GetAuditFilterOperation {
	if o == nil {
		return &xxx_GetAuditFilterOperation{}
	}
	return &xxx_GetAuditFilterOperation{
		That:   o.That,
		Filter: o.Filter,
		Return: o.Return,
	}
}

func (o *GetAuditFilterResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAuditFilterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Filter = op.Filter
	o.Return = op.Return
}
func (o *GetAuditFilterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAuditFilterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAuditFilterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAuditFilterOperation structure represents the SetAuditFilter operation
type xxx_SetAuditFilterOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Filter    uint32         `idl:"name:dwFilter" json:"filter"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAuditFilterOperation) OpNum() int { return 41 }

func (o *xxx_SetAuditFilterOperation) OpName() string { return "/ICertAdminD2/v0/SetAuditFilter" }

func (o *xxx_SetAuditFilterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuditFilterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuditFilterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuditFilterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuditFilterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAuditFilterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAuditFilterRequest structure represents the SetAuditFilter operation request
type SetAuditFilterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// dwFilter: An unsigned integer that specifies the events to be audited by the CA.
	// For possible values, see section 3.1.4.2.10.
	//
	// The SetAuditFilter method is used to set the audit filter value that is passed in
	// by the client. The audit filter value is used to determine which actions are audited.
	Filter uint32 `idl:"name:dwFilter" json:"filter"`
}

func (o *SetAuditFilterRequest) xxx_ToOp(ctx context.Context) *xxx_SetAuditFilterOperation {
	if o == nil {
		return &xxx_SetAuditFilterOperation{}
	}
	return &xxx_SetAuditFilterOperation{
		This:      o.This,
		Authority: o.Authority,
		Filter:    o.Filter,
	}
}

func (o *SetAuditFilterRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAuditFilterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.Filter = op.Filter
}
func (o *SetAuditFilterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetAuditFilterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAuditFilterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAuditFilterResponse structure represents the SetAuditFilter operation response
type SetAuditFilterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetAuditFilter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAuditFilterResponse) xxx_ToOp(ctx context.Context) *xxx_SetAuditFilterOperation {
	if o == nil {
		return &xxx_SetAuditFilterOperation{}
	}
	return &xxx_SetAuditFilterOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetAuditFilterResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAuditFilterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAuditFilterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetAuditFilterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAuditFilterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetOfficerRightsOperation structure represents the GetOfficerRights operation
type xxx_GetOfficerRightsOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Enabled   bool                    `idl:"name:pfEnabled" json:"enabled"`
	SD        *csra.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOfficerRightsOperation) OpNum() int { return 42 }

func (o *xxx_GetOfficerRightsOperation) OpName() string { return "/ICertAdminD2/v0/GetOfficerRights" }

func (o *xxx_GetOfficerRightsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOfficerRightsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetOfficerRightsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetOfficerRightsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOfficerRightsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.Enabled {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// pctbSD {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD != nil {
			if err := o.SD.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetOfficerRightsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfEnabled {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bEnabled int32
		if err := w.ReadData(&_bEnabled); err != nil {
			return err
		}
		o.Enabled = _bEnabled != 0
	}
	// pctbSD {out} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD == nil {
			o.SD = &csra.CertTransportBlob{}
		}
		if err := o.SD.UnmarshalNDR(ctx, w); err != nil {
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

// GetOfficerRightsRequest structure represents the GetOfficerRights operation request
type GetOfficerRightsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *GetOfficerRightsRequest) xxx_ToOp(ctx context.Context) *xxx_GetOfficerRightsOperation {
	if o == nil {
		return &xxx_GetOfficerRightsOperation{}
	}
	return &xxx_GetOfficerRightsOperation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *GetOfficerRightsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOfficerRightsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *GetOfficerRightsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetOfficerRightsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOfficerRightsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOfficerRightsResponse structure represents the GetOfficerRights operation response
type GetOfficerRightsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pfEnabled: A pointer to a Boolean value.
	Enabled bool `idl:"name:pfEnabled" json:"enabled"`
	// pctbSD: A pointer to the CERTTRANSBLOB structure that contains the marshaled information
	// specified in section 2.2.1.11.1.
	SD *csra.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
	// Return: The GetOfficerRights return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetOfficerRightsResponse) xxx_ToOp(ctx context.Context) *xxx_GetOfficerRightsOperation {
	if o == nil {
		return &xxx_GetOfficerRightsOperation{}
	}
	return &xxx_GetOfficerRightsOperation{
		That:    o.That,
		Enabled: o.Enabled,
		SD:      o.SD,
		Return:  o.Return,
	}
}

func (o *GetOfficerRightsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOfficerRightsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enabled = op.Enabled
	o.SD = op.SD
	o.Return = op.Return
}
func (o *GetOfficerRightsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetOfficerRightsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOfficerRightsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOfficerRightsOperation structure represents the SetOfficerRights operation
type xxx_SetOfficerRightsOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Enable    bool                    `idl:"name:fEnable" json:"enable"`
	SD        *csra.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOfficerRightsOperation) OpNum() int { return 43 }

func (o *xxx_SetOfficerRightsOperation) OpName() string { return "/ICertAdminD2/v0/SetOfficerRights" }

func (o *xxx_SetOfficerRightsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOfficerRightsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fEnable {in} (1:{alias=BOOL}(int32))
	{
		if !o.Enable {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// pctbSD {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD != nil {
			if err := o.SD.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetOfficerRightsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fEnable {in} (1:{alias=BOOL}(int32))
	{
		var _bEnable int32
		if err := w.ReadData(&_bEnable); err != nil {
			return err
		}
		o.Enable = _bEnable != 0
	}
	// pctbSD {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.SD == nil {
			o.SD = &csra.CertTransportBlob{}
		}
		if err := o.SD.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOfficerRightsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOfficerRightsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOfficerRightsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetOfficerRightsRequest structure represents the SetOfficerRights operation request
type SetOfficerRightsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// fEnable: A 32-bit BOOL parameter composed of two 16-bit fields. Each of these fields
	// can be set to zero or to a nonzero value as follows.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| F | F | F | F | F | F | F | F | F | F | F | F | F | F | F | F | R | R | R | R | R | R | R | R | R | R | R | R | R | R | R | R |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// F - fRightsEnable: If bits 0 through 15 are 0, then disable access rights (officer
	// or enrollment agent) and ignore the value of pctbSD.
	//
	// R - RightsType: If bits 16 through 31 are 0, then the security descriptor in the
	// pctbSD parameter is for officer rights.
	Enable bool `idl:"name:fEnable" json:"enable"`
	// pctbSD: A pointer to the CERTTRANSBLOB structure that holds the marshaled security
	// descriptor, as specified in [MS-DTYP] section 2.4.6.
	SD *csra.CertTransportBlob `idl:"name:pctbSD;pointer:ref" json:"sd"`
}

func (o *SetOfficerRightsRequest) xxx_ToOp(ctx context.Context) *xxx_SetOfficerRightsOperation {
	if o == nil {
		return &xxx_SetOfficerRightsOperation{}
	}
	return &xxx_SetOfficerRightsOperation{
		This:      o.This,
		Authority: o.Authority,
		Enable:    o.Enable,
		SD:        o.SD,
	}
}

func (o *SetOfficerRightsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOfficerRightsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.Enable = op.Enable
	o.SD = op.SD
}
func (o *SetOfficerRightsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetOfficerRightsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOfficerRightsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOfficerRightsResponse structure represents the SetOfficerRights operation response
type SetOfficerRightsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetOfficerRights return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetOfficerRightsResponse) xxx_ToOp(ctx context.Context) *xxx_SetOfficerRightsOperation {
	if o == nil {
		return &xxx_SetOfficerRightsOperation{}
	}
	return &xxx_SetOfficerRightsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetOfficerRightsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOfficerRightsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetOfficerRightsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetOfficerRightsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOfficerRightsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetConfigEntryOperation structure represents the GetConfigEntry operation
type xxx_GetConfigEntryOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	NodePath  string         `idl:"name:pwszNodePath;string;pointer:unique" json:"node_path"`
	Entry     string         `idl:"name:pwszEntry;string;pointer:ref" json:"entry"`
	Variant   *oaut.Variant  `idl:"name:pVariant;pointer:ref" json:"variant"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetConfigEntryOperation) OpNum() int { return 44 }

func (o *xxx_GetConfigEntryOperation) OpName() string { return "/ICertAdminD2/v0/GetConfigEntry" }

func (o *xxx_GetConfigEntryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigEntryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszNodePath {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.NodePath != "" {
			_ptr_pwszNodePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NodePath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NodePath, _ptr_pwszNodePath); err != nil {
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
	// pwszEntry {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Entry); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigEntryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszNodePath {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszNodePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NodePath); err != nil {
				return err
			}
			return nil
		})
		_s_pwszNodePath := func(ptr interface{}) { o.NodePath = *ptr.(*string) }
		if err := w.ReadPointer(&o.NodePath, _s_pwszNodePath, _ptr_pwszNodePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszEntry {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Entry); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigEntryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigEntryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVariant {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Variant != nil {
			_ptr_pVariant := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Variant != nil {
					if err := o.Variant.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Variant, _ptr_pVariant); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigEntryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVariant {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pVariant := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Variant == nil {
				o.Variant = &oaut.Variant{}
			}
			if err := o.Variant.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVariant := func(ptr interface{}) { o.Variant = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Variant, _s_pVariant, _ptr_pVariant); err != nil {
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

// GetConfigEntryRequest structure represents the GetConfigEntry operation request
type GetConfigEntryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// pwszNodePath: A string value that represents the node path for the configuration
	// information. This parameter can be an empty string and MUST NOT be NULL.<72>
	NodePath string `idl:"name:pwszNodePath;string;pointer:unique" json:"node_path"`
	// pwszEntry: A string value that represents the name of the leaf entry whose information
	// is being retrieved. This value can be an EMPTY string and MUST NOT be NULL.<73>
	Entry string `idl:"name:pwszEntry;string;pointer:ref" json:"entry"`
}

func (o *GetConfigEntryRequest) xxx_ToOp(ctx context.Context) *xxx_GetConfigEntryOperation {
	if o == nil {
		return &xxx_GetConfigEntryOperation{}
	}
	return &xxx_GetConfigEntryOperation{
		This:      o.This,
		Authority: o.Authority,
		NodePath:  o.NodePath,
		Entry:     o.Entry,
	}
}

func (o *GetConfigEntryRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConfigEntryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.NodePath = op.NodePath
	o.Entry = op.Entry
}
func (o *GetConfigEntryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetConfigEntryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigEntryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetConfigEntryResponse structure represents the GetConfigEntry operation response
type GetConfigEntryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pVariant: A pointer to a VARIANT that receives the requested information.
	//
	// On Windows, the CA uses these datatypes to set the data that it stores in the registry:
	//
	//	REG_BINARY – The vt member of VARIANT is set to VT_ARRAY|VT_UI1 and the pArray member references a single dimension SAFEARRAY the binary data. The number of elements of the SAFEARRAY reference by pArray is equal to the length of binary data.
	//
	// REG_DWORD – The vt member of VARIANT is set to VT_I4 and the lVal member is the
	// registry value.
	//
	// REG_SZ – The vt member of VARIANT is set to VT_BSTR and the bstrVal member is set
	// to BSTR for Unicode string in the registry value.
	//
	// The GetConfigEntry method retrieves the CA configuration data or configuration data
	// hierarchy information.
	Variant *oaut.Variant `idl:"name:pVariant;pointer:ref" json:"variant"`
	// Return: The GetConfigEntry return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetConfigEntryResponse) xxx_ToOp(ctx context.Context) *xxx_GetConfigEntryOperation {
	if o == nil {
		return &xxx_GetConfigEntryOperation{}
	}
	return &xxx_GetConfigEntryOperation{
		That:    o.That,
		Variant: o.Variant,
		Return:  o.Return,
	}
}

func (o *GetConfigEntryResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConfigEntryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Variant = op.Variant
	o.Return = op.Return
}
func (o *GetConfigEntryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetConfigEntryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigEntryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetConfigEntryOperation structure represents the SetConfigEntry operation
type xxx_SetConfigEntryOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	NodePath  string         `idl:"name:pwszNodePath;string;pointer:unique" json:"node_path"`
	Entry     string         `idl:"name:pwszEntry;string;pointer:ref" json:"entry"`
	Variant   *oaut.Variant  `idl:"name:pVariant;pointer:ref" json:"variant"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetConfigEntryOperation) OpNum() int { return 45 }

func (o *xxx_SetConfigEntryOperation) OpName() string { return "/ICertAdminD2/v0/SetConfigEntry" }

func (o *xxx_SetConfigEntryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigEntryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszNodePath {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.NodePath != "" {
			_ptr_pwszNodePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NodePath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NodePath, _ptr_pwszNodePath); err != nil {
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
	// pwszEntry {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Entry); err != nil {
			return err
		}
	}
	// pVariant {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Variant != nil {
			_ptr_pVariant := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Variant != nil {
					if err := o.Variant.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Variant, _ptr_pVariant); err != nil {
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

func (o *xxx_SetConfigEntryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszNodePath {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszNodePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NodePath); err != nil {
				return err
			}
			return nil
		})
		_s_pwszNodePath := func(ptr interface{}) { o.NodePath = *ptr.(*string) }
		if err := w.ReadPointer(&o.NodePath, _s_pwszNodePath, _ptr_pwszNodePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszEntry {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Entry); err != nil {
			return err
		}
	}
	// pVariant {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pVariant := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Variant == nil {
				o.Variant = &oaut.Variant{}
			}
			if err := o.Variant.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVariant := func(ptr interface{}) { o.Variant = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Variant, _s_pVariant, _ptr_pVariant); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigEntryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigEntryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetConfigEntryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetConfigEntryRequest structure represents the SetConfigEntry operation request
type SetConfigEntryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// pwszNodePath: A string value that represents the node path for the configuration
	// information. This parameter can be an EMPTY string and MUST NOT be NULL.
	NodePath string `idl:"name:pwszNodePath;string;pointer:unique" json:"node_path"`
	// pwszEntry: A string value that represents the name of the leaf entry whose information
	// is being set. This value can be an EMPTY string and MUST NOT be NULL.
	Entry string `idl:"name:pwszEntry;string;pointer:ref" json:"entry"`
	// pVariant: A pointer to VARIANT that specifies the information to set. If this value
	// is EMPTY, the indicated entry MUST be deleted.
	Variant *oaut.Variant `idl:"name:pVariant;pointer:ref" json:"variant"`
}

func (o *SetConfigEntryRequest) xxx_ToOp(ctx context.Context) *xxx_SetConfigEntryOperation {
	if o == nil {
		return &xxx_SetConfigEntryOperation{}
	}
	return &xxx_SetConfigEntryOperation{
		This:      o.This,
		Authority: o.Authority,
		NodePath:  o.NodePath,
		Entry:     o.Entry,
		Variant:   o.Variant,
	}
}

func (o *SetConfigEntryRequest) xxx_FromOp(ctx context.Context, op *xxx_SetConfigEntryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.NodePath = op.NodePath
	o.Entry = op.Entry
	o.Variant = op.Variant
}
func (o *SetConfigEntryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetConfigEntryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetConfigEntryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetConfigEntryResponse structure represents the SetConfigEntry operation response
type SetConfigEntryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetConfigEntry return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetConfigEntryResponse) xxx_ToOp(ctx context.Context) *xxx_SetConfigEntryOperation {
	if o == nil {
		return &xxx_SetConfigEntryOperation{}
	}
	return &xxx_SetConfigEntryOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetConfigEntryResponse) xxx_FromOp(ctx context.Context, op *xxx_SetConfigEntryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetConfigEntryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetConfigEntryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetConfigEntryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportKeyOperation structure represents the ImportKey operation
type xxx_ImportKeyOperation struct {
	This      *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Authority string                  `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	RequestID uint32                  `idl:"name:dwRequestId" json:"request_id"`
	CertHash  string                  `idl:"name:pwszCertHash;string;pointer:unique" json:"cert_hash"`
	Flags     uint32                  `idl:"name:dwFlags" json:"flags"`
	Key       *csra.CertTransportBlob `idl:"name:pctbKey;pointer:ref" json:"key"`
	Return    int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportKeyOperation) OpNum() int { return 46 }

func (o *xxx_ImportKeyOperation) OpName() string { return "/ICertAdminD2/v0/ImportKey" }

func (o *xxx_ImportKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszCertHash {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.CertHash != "" {
			_ptr_pwszCertHash := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.CertHash); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.CertHash, _ptr_pwszCertHash); err != nil {
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
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pctbKey {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ImportKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszCertHash {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszCertHash := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.CertHash); err != nil {
				return err
			}
			return nil
		})
		_s_pwszCertHash := func(ptr interface{}) { o.CertHash = *ptr.(*string) }
		if err := w.ReadPointer(&o.CertHash, _s_pwszCertHash, _ptr_pwszCertHash); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pctbKey {in} (1:{pointer=ref}*(1))(2:{alias=CERTTRANSBLOB}(struct))
	{
		if o.Key == nil {
			o.Key = &csra.CertTransportBlob{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ImportKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ImportKeyRequest structure represents the ImportKey operation request
type ImportKeyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// dwRequestId: An unsigned integer value that represents the certificate request ID
	// in the CA database.
	RequestID uint32 `idl:"name:dwRequestId" json:"request_id"`
	// pwszCertHash: A null-terminated Unicode string value that represents the SHA-1 hash
	// of the ASN.1 DER–encoded certificate data (as specified in [X660]) and that is
	// formatted as a hexadecimal string.
	CertHash string `idl:"name:pwszCertHash;string;pointer:unique" json:"cert_hash"`
	// dwFlags: An unsigned integer that specifies the optional flags for this method.
	//
	//	+------------+--------------------------------------------------+
	//	|            |                                                  |
	//	|   VALUE    |                     MEANING                      |
	//	|            |                                                  |
	//	+------------+--------------------------------------------------+
	//	+------------+--------------------------------------------------+
	//	| 0x00010000 | Overwrite the existing archived key, if present. |
	//	+------------+--------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// pctbKey: A CERTTRANSBLOB structure that contains the ASN.1 DER–encoded (as specified
	// in [X660] and [X690]) PKCS#7 message (as specified in [RFC2315]) that contains the
	// private key to be archived. The content of the enveloped PKCS#7 is as specified in
	// [MS-WCCE] section 3.2.1.4.2.1.4.
	Key *csra.CertTransportBlob `idl:"name:pctbKey;pointer:ref" json:"key"`
}

func (o *ImportKeyRequest) xxx_ToOp(ctx context.Context) *xxx_ImportKeyOperation {
	if o == nil {
		return &xxx_ImportKeyOperation{}
	}
	return &xxx_ImportKeyOperation{
		This:      o.This,
		Authority: o.Authority,
		RequestID: o.RequestID,
		CertHash:  o.CertHash,
		Flags:     o.Flags,
		Key:       o.Key,
	}
}

func (o *ImportKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportKeyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.RequestID = op.RequestID
	o.CertHash = op.CertHash
	o.Flags = op.Flags
	o.Key = op.Key
}
func (o *ImportKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ImportKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportKeyResponse structure represents the ImportKey operation response
type ImportKeyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ImportKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportKeyResponse) xxx_ToOp(ctx context.Context) *xxx_ImportKeyOperation {
	if o == nil {
		return &xxx_ImportKeyOperation{}
	}
	return &xxx_ImportKeyOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ImportKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportKeyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ImportKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ImportKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMyRolesOperation structure represents the GetMyRoles operation
type xxx_GetMyRolesOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Roles     int32          `idl:"name:pdwRoles" json:"roles"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMyRolesOperation) OpNum() int { return 47 }

func (o *xxx_GetMyRolesOperation) OpName() string { return "/ICertAdminD2/v0/GetMyRoles" }

func (o *xxx_GetMyRolesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMyRolesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMyRolesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMyRolesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMyRolesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwRoles {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Roles); err != nil {
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

func (o *xxx_GetMyRolesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwRoles {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Roles); err != nil {
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

// GetMyRolesRequest structure represents the GetMyRoles operation request
type GetMyRolesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority: See the pwszAuthority definition in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
}

func (o *GetMyRolesRequest) xxx_ToOp(ctx context.Context) *xxx_GetMyRolesOperation {
	if o == nil {
		return &xxx_GetMyRolesOperation{}
	}
	return &xxx_GetMyRolesOperation{
		This:      o.This,
		Authority: o.Authority,
	}
}

func (o *GetMyRolesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMyRolesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
}
func (o *GetMyRolesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMyRolesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMyRolesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMyRolesResponse structure represents the GetMyRoles operation response
type GetMyRolesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwRoles: A bitwise-OR combination of zero or more CA security values based on the
	// CA implementation. Microsoft CA permissions are defined in section 3.1.1.7.
	//
	// For pdwRoles, the server MUST return a signed integer that represents the CA roles
	// assigned to the caller.
	Roles int32 `idl:"name:pdwRoles" json:"roles"`
	// Return: The GetMyRoles return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMyRolesResponse) xxx_ToOp(ctx context.Context) *xxx_GetMyRolesOperation {
	if o == nil {
		return &xxx_GetMyRolesOperation{}
	}
	return &xxx_GetMyRolesOperation{
		That:   o.That,
		Roles:  o.Roles,
		Return: o.Return,
	}
}

func (o *GetMyRolesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMyRolesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Roles = op.Roles
	o.Return = op.Return
}
func (o *GetMyRolesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMyRolesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMyRolesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteRowOperation structure represents the DeleteRow operation
type xxx_DeleteRowOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authority    string         `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	Flags        uint32         `idl:"name:dwFlags" json:"flags"`
	FileTime     *dtyp.Filetime `idl:"name:FileTime" json:"file_time"`
	Table        uint32         `idl:"name:dwTable" json:"table"`
	RowID        uint32         `idl:"name:dwRowId" json:"row_id"`
	DeletedCount int32          `idl:"name:pcDeleted" json:"deleted_count"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteRowOperation) OpNum() int { return 48 }

func (o *xxx_DeleteRowOperation) OpName() string { return "/ICertAdminD2/v0/DeleteRow" }

func (o *xxx_DeleteRowOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteRowOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
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
	// dwTable {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Table); err != nil {
			return err
		}
	}
	// dwRowId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RowID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteRowOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
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
	// dwTable {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Table); err != nil {
			return err
		}
	}
	// dwRowId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RowID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteRowOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteRowOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcDeleted {out, retval} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.DeletedCount); err != nil {
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

func (o *xxx_DeleteRowOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcDeleted {out, retval} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.DeletedCount); err != nil {
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

// DeleteRowRequest structure represents the DeleteRow operation request
type DeleteRowRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszAuthority:  See the definition of the pwszAuthority parameter in section 3.1.4.1.1.
	Authority string `idl:"name:pwszAuthority;string;pointer:unique" json:"authority"`
	// dwFlags: An unsigned integer value that specifies the type of rows to be deleted.
	// This parameter can be one of the following values.
	//
	//	+------------+----------------------------------------------------------+
	//	|            |                                                          |
	//	|   VALUE    |                         MEANING                          |
	//	|            |                                                          |
	//	+------------+----------------------------------------------------------+
	//	+------------+----------------------------------------------------------+
	//	| 0x00000000 | Delete the individual row.                               |
	//	+------------+----------------------------------------------------------+
	//	| 0x00000001 | Delete the rows that contain expired certificates.       |
	//	+------------+----------------------------------------------------------+
	//	| 0x00000002 | Delete the rows that contain pending or failed requests. |
	//	+------------+----------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// FileTime:  Contains a 64-bit value that represents the number of 100-nanosecond
	// intervals since January 1, 1601 (UTC). The value is used to query for multiple rows
	// to be deleted. It MUST contain all zeros if the dwRowId parameter is nonzero.
	FileTime *dtyp.Filetime `idl:"name:FileTime" json:"file_time"`
	// dwTable: An unsigned integer value that specifies the table in which to delete rows.
	// This parameter can be one of the following values.
	//
	//	+------------+----------------------------------+
	//	|            |                                  |
	//	|   VALUE    |             MEANING              |
	//	|            |                                  |
	//	+------------+----------------------------------+
	//	+------------+----------------------------------+
	//	| 0x00000000 | Delete the Request table rows.   |
	//	+------------+----------------------------------+
	//	| 0x00003000 | Delete the Extension table rows. |
	//	+------------+----------------------------------+
	//	| 0x00004000 | Delete the Attribute table rows. |
	//	+------------+----------------------------------+
	//	| 0x00005000 | Delete the CRL table rows.       |
	//	+------------+----------------------------------+
	Table uint32 `idl:"name:dwTable" json:"table"`
	// dwRowId: An unsigned integer value that represents the row identifier in the CA data
	// table. MUST be set to 0 if FileTime is nonzero.
	RowID uint32 `idl:"name:dwRowId" json:"row_id"`
}

func (o *DeleteRowRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteRowOperation {
	if o == nil {
		return &xxx_DeleteRowOperation{}
	}
	return &xxx_DeleteRowOperation{
		This:      o.This,
		Authority: o.Authority,
		Flags:     o.Flags,
		FileTime:  o.FileTime,
		Table:     o.Table,
		RowID:     o.RowID,
	}
}

func (o *DeleteRowRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteRowOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authority = op.Authority
	o.Flags = op.Flags
	o.FileTime = op.FileTime
	o.Table = op.Table
	o.RowID = op.RowID
}
func (o *DeleteRowRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteRowRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteRowOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteRowResponse structure represents the DeleteRow operation response
type DeleteRowResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcDeleted:  Returns the count of successfully deleted table rows.
	//
	// The DeleteRow method is used to instruct the CA to delete rows from the specified
	// table.
	DeletedCount int32 `idl:"name:pcDeleted" json:"deleted_count"`
	// Return: The DeleteRow return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteRowResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteRowOperation {
	if o == nil {
		return &xxx_DeleteRowOperation{}
	}
	return &xxx_DeleteRowOperation{
		That:         o.That,
		DeletedCount: o.DeletedCount,
		Return:       o.Return,
	}
}

func (o *DeleteRowResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteRowOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DeletedCount = op.DeletedCount
	o.Return = op.Return
}
func (o *DeleteRowResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteRowResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteRowOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
