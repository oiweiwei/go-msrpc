package ifsrmquotamanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmQuotaManager interface identifier 8bb68c7d-19d8-4ffb-809e-be4fc1734014
	QuotaManagerIID = &dcom.IID{Data1: 0x8bb68c7d, Data2: 0x19d8, Data3: 0x4ffb, Data4: []byte{0x80, 0x9e, 0xbe, 0x4f, 0xc1, 0x73, 0x40, 0x14}}
	// Syntax UUID
	QuotaManagerSyntaxUUID = &uuid.UUID{TimeLow: 0x8bb68c7d, TimeMid: 0x19d8, TimeHiAndVersion: 0x4ffb, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x9e, Node: [6]uint8{0xbe, 0x4f, 0xc1, 0x73, 0x40, 0x14}}
	// Syntax ID
	QuotaManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QuotaManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmQuotaManager interface.
type QuotaManagerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// ActionVariables operation.
	GetActionVariables(context.Context, *GetActionVariablesRequest, ...dcerpc.CallOption) (*GetActionVariablesResponse, error)

	// ActionVariableDescriptions operation.
	GetActionVariableDescriptions(context.Context, *GetActionVariableDescriptionsRequest, ...dcerpc.CallOption) (*GetActionVariableDescriptionsResponse, error)

	// The CreateQuota method creates a blank Non-Persisted Directory Quota Instance (section
	// 3.2.1.2.1.2) for the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+--------------------------------------------------+
	//	|              RETURN              |                                                  |
	//	|            VALUE/CODE            |                   DESCRIPTION                    |
	//	|                                  |                                                  |
	//	+----------------------------------+--------------------------------------------------+
	//	+----------------------------------+--------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | The quota for the specified path already exists. |
	//	+----------------------------------+--------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | One of the quota parameters is NULL.             |
	//	+----------------------------------+--------------------------------------------------+
	CreateQuota(context.Context, *CreateQuotaRequest, ...dcerpc.CallOption) (*CreateQuotaResponse, error)

	// The CreateAutoApplyQuota method creates a Non-Persisted Auto Apply Quota Instance
	// (section 3.2.1.2.2.2) for the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+-------------------------------------------------------------+
	//	|              RETURN              |                                                             |
	//	|            VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                  |                                                             |
	//	+----------------------------------+-------------------------------------------------------------+
	//	+----------------------------------+-------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND      | The specified auto apply quota could not be found.          |
	//	+----------------------------------+-------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | The auto apply quota for the specified path already exists. |
	//	+----------------------------------+-------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | One of the quota parameters is NULL.                        |
	//	+----------------------------------+-------------------------------------------------------------+
	CreateAutoApplyQuota(context.Context, *CreateAutoApplyQuotaRequest, ...dcerpc.CallOption) (*CreateAutoApplyQuotaResponse, error)

	// The GetQuota method returns the directory quota from the List of Persisted Directory
	// Quotas (section 3.2.1.2) for the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND      | The specified quota could not be found.                                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND | The quota for the specified path could not be found.                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH   | The content of the path parameter exceeds the maximum length of 4,000            |
	//	|                                  | characters.                                                                      |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The path parameter is NULL. The |
	//	|                                  | quota parameter is NULL.                                                         |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	GetQuota(context.Context, *GetQuotaRequest, ...dcerpc.CallOption) (*GetQuotaResponse, error)

	// The GetAutoApplyQuota method returns the auto apply quota from the List of Persisted
	// Auto Apply Quotas (section 3.2.1.2) for the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND      | The specified auto apply quota could not be found.                               |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND | The auto apply quota for the specified path could not be found.                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH   | The content of the path parameter exceeds the maximum length of 260 characters.  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The path parameter is NULL. The |
	//	|                                  | quota parameter is NULL.                                                         |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	GetAutoApplyQuota(context.Context, *GetAutoApplyQuotaRequest, ...dcerpc.CallOption) (*GetAutoApplyQuotaResponse, error)

	// The GetRestrictiveQuota method returns the directory quota from the List of Persisted
	// Directory Quotas (section 3.2.1.2) with the lowest quota limit for the specified
	// path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND      | The specified quota could not be found.                                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND | The restrictive quota for the specified path could not be found.                 |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH   | The content of the path parameter exceeds the maximum length of 260 characters.  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The path parameter is NULL. The |
	//	|                                  | quota parameter is NULL.                                                         |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	GetRestrictiveQuota(context.Context, *GetRestrictiveQuotaRequest, ...dcerpc.CallOption) (*GetRestrictiveQuotaResponse, error)

	// The EnumQuotas method returns all the directory quotas from the List of Persisted
	// Directory Quotas (section 3.2.1.2) that fall under the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | The quotas parameter is NULL.                                                    |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | The options parameter contains invalid FsrmEnumOptions (section 2.2.1.2.5)       |
	//	|                                 | values.                                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumQuotas(context.Context, *EnumQuotasRequest, ...dcerpc.CallOption) (*EnumQuotasResponse, error)

	// The EnumAutoApplyQuotas method returns all the auto apply quotas from the List of
	// Persisted Auto Apply Quotas (section 3.2.1.2) that fall under the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | The quotas parameter is NULL.                                                    |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | The options parameter contains invalid FsrmEnumOptions (section 2.2.1.2.5)       |
	//	|                                 | values.                                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumAutoApplyQuotas(context.Context, *EnumAutoApplyQuotasRequest, ...dcerpc.CallOption) (*EnumAutoApplyQuotasResponse, error)

	// The EnumEffectiveQuotas method returns all the directory quotas from the List of
	// Persisted Directory Quotas (section 3.2.1.2) that affect the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | The quotas parameter is NULL.                                                    |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | The options parameter contains invalid FsrmEnumOptions (section 2.2.1.2.5)       |
	//	|                                 | values.                                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumEffectiveQuotas(context.Context, *EnumEffectiveQuotasRequest, ...dcerpc.CallOption) (*EnumEffectiveQuotasResponse, error)

	// The Scan method starts a quota scan on the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND    | A quota for the specified path could not be found.                               |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH | The content of the strPath parameter exceeds the maximum length of 260           |
	//	|                                | characters.                                                                      |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | The strPath parameter is NULL or empty.                                          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	Scan(context.Context, *ScanRequest, ...dcerpc.CallOption) (*ScanResponse, error)

	// The CreateQuotaCollection method creates an empty collection for callers to add quotas
	// to.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------------+
	//	|         RETURN          |                                   |
	//	|       VALUE/CODE        |            DESCRIPTION            |
	//	|                         |                                   |
	//	+-------------------------+-----------------------------------+
	//	+-------------------------+-----------------------------------+
	//	| 0x80070057 E_INVALIDARG | The collection parameter is NULL. |
	//	+-------------------------+-----------------------------------+
	CreateQuotaCollection(context.Context, *CreateQuotaCollectionRequest, ...dcerpc.CallOption) (*CreateQuotaCollectionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QuotaManagerClient
}

type xxx_DefaultQuotaManagerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQuotaManagerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultQuotaManagerClient) GetActionVariables(ctx context.Context, in *GetActionVariablesRequest, opts ...dcerpc.CallOption) (*GetActionVariablesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetActionVariablesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) GetActionVariableDescriptions(ctx context.Context, in *GetActionVariableDescriptionsRequest, opts ...dcerpc.CallOption) (*GetActionVariableDescriptionsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetActionVariableDescriptionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) CreateQuota(ctx context.Context, in *CreateQuotaRequest, opts ...dcerpc.CallOption) (*CreateQuotaResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CreateQuotaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) CreateAutoApplyQuota(ctx context.Context, in *CreateAutoApplyQuotaRequest, opts ...dcerpc.CallOption) (*CreateAutoApplyQuotaResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CreateAutoApplyQuotaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) GetQuota(ctx context.Context, in *GetQuotaRequest, opts ...dcerpc.CallOption) (*GetQuotaResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetQuotaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) GetAutoApplyQuota(ctx context.Context, in *GetAutoApplyQuotaRequest, opts ...dcerpc.CallOption) (*GetAutoApplyQuotaResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetAutoApplyQuotaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) GetRestrictiveQuota(ctx context.Context, in *GetRestrictiveQuotaRequest, opts ...dcerpc.CallOption) (*GetRestrictiveQuotaResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &GetRestrictiveQuotaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) EnumQuotas(ctx context.Context, in *EnumQuotasRequest, opts ...dcerpc.CallOption) (*EnumQuotasResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &EnumQuotasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) EnumAutoApplyQuotas(ctx context.Context, in *EnumAutoApplyQuotasRequest, opts ...dcerpc.CallOption) (*EnumAutoApplyQuotasResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &EnumAutoApplyQuotasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) EnumEffectiveQuotas(ctx context.Context, in *EnumEffectiveQuotasRequest, opts ...dcerpc.CallOption) (*EnumEffectiveQuotasResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &EnumEffectiveQuotasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) Scan(ctx context.Context, in *ScanRequest, opts ...dcerpc.CallOption) (*ScanResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &ScanResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) CreateQuotaCollection(ctx context.Context, in *CreateQuotaCollectionRequest, opts ...dcerpc.CallOption) (*CreateQuotaCollectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
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
	out := &CreateQuotaCollectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQuotaManagerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQuotaManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) QuotaManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQuotaManagerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewQuotaManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QuotaManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QuotaManagerSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultQuotaManagerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetActionVariablesOperation structure represents the ActionVariables operation
type xxx_GetActionVariablesOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Variables *oaut.SafeArray `idl:"name:variables" json:"variables"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetActionVariablesOperation) OpNum() int { return 7 }

func (o *xxx_GetActionVariablesOperation) OpName() string {
	return "/IFsrmQuotaManager/v0/ActionVariables"
}

func (o *xxx_GetActionVariablesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariablesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetActionVariablesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetActionVariablesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariablesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// variables {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Variables != nil {
			_ptr_variables := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Variables != nil {
					if err := o.Variables.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Variables, _ptr_variables); err != nil {
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

func (o *xxx_GetActionVariablesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// variables {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_variables := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Variables == nil {
				o.Variables = &oaut.SafeArray{}
			}
			if err := o.Variables.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_variables := func(ptr interface{}) { o.Variables = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Variables, _s_variables, _ptr_variables); err != nil {
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

// GetActionVariablesRequest structure represents the ActionVariables operation request
type GetActionVariablesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetActionVariablesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetActionVariablesOperation) *xxx_GetActionVariablesOperation {
	if op == nil {
		op = &xxx_GetActionVariablesOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetActionVariablesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariablesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetActionVariablesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetActionVariablesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariablesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetActionVariablesResponse structure represents the ActionVariables operation response
type GetActionVariablesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Variables *oaut.SafeArray `idl:"name:variables" json:"variables"`
	// Return: The ActionVariables return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetActionVariablesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetActionVariablesOperation) *xxx_GetActionVariablesOperation {
	if op == nil {
		op = &xxx_GetActionVariablesOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Variables = op.Variables
	o.Return = op.Return
	return op
}

func (o *GetActionVariablesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariablesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Variables = op.Variables
	o.Return = op.Return
}
func (o *GetActionVariablesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetActionVariablesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariablesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetActionVariableDescriptionsOperation structure represents the ActionVariableDescriptions operation
type xxx_GetActionVariableDescriptionsOperation struct {
	This         *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Descriptions *oaut.SafeArray `idl:"name:descriptions" json:"descriptions"`
	Return       int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetActionVariableDescriptionsOperation) OpNum() int { return 8 }

func (o *xxx_GetActionVariableDescriptionsOperation) OpName() string {
	return "/IFsrmQuotaManager/v0/ActionVariableDescriptions"
}

func (o *xxx_GetActionVariableDescriptionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariableDescriptionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetActionVariableDescriptionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetActionVariableDescriptionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariableDescriptionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// descriptions {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Descriptions != nil {
			_ptr_descriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Descriptions != nil {
					if err := o.Descriptions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Descriptions, _ptr_descriptions); err != nil {
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

func (o *xxx_GetActionVariableDescriptionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// descriptions {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_descriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Descriptions == nil {
				o.Descriptions = &oaut.SafeArray{}
			}
			if err := o.Descriptions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_descriptions := func(ptr interface{}) { o.Descriptions = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Descriptions, _s_descriptions, _ptr_descriptions); err != nil {
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

// GetActionVariableDescriptionsRequest structure represents the ActionVariableDescriptions operation request
type GetActionVariableDescriptionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetActionVariableDescriptionsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetActionVariableDescriptionsOperation) *xxx_GetActionVariableDescriptionsOperation {
	if op == nil {
		op = &xxx_GetActionVariableDescriptionsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetActionVariableDescriptionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariableDescriptionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetActionVariableDescriptionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetActionVariableDescriptionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariableDescriptionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetActionVariableDescriptionsResponse structure represents the ActionVariableDescriptions operation response
type GetActionVariableDescriptionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Descriptions *oaut.SafeArray `idl:"name:descriptions" json:"descriptions"`
	// Return: The ActionVariableDescriptions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetActionVariableDescriptionsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetActionVariableDescriptionsOperation) *xxx_GetActionVariableDescriptionsOperation {
	if op == nil {
		op = &xxx_GetActionVariableDescriptionsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Descriptions = op.Descriptions
	o.Return = op.Return
	return op
}

func (o *GetActionVariableDescriptionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariableDescriptionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Descriptions = op.Descriptions
	o.Return = op.Return
}
func (o *GetActionVariableDescriptionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetActionVariableDescriptionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariableDescriptionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateQuotaOperation structure represents the CreateQuota operation
type xxx_CreateQuotaOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   *oaut.String   `idl:"name:path" json:"path"`
	Quota  *fsrm.Quota    `idl:"name:quota" json:"quota"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateQuotaOperation) OpNum() int { return 9 }

func (o *xxx_CreateQuotaOperation) OpName() string { return "/IFsrmQuotaManager/v0/CreateQuota" }

func (o *xxx_CreateQuotaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateQuotaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_CreateQuotaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateQuotaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateQuotaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quota {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmQuota}(interface))
	{
		if o.Quota != nil {
			_ptr_quota := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Quota != nil {
					if err := o.Quota.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Quota{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Quota, _ptr_quota); err != nil {
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

func (o *xxx_CreateQuotaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quota {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmQuota}(interface))
	{
		_ptr_quota := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Quota == nil {
				o.Quota = &fsrm.Quota{}
			}
			if err := o.Quota.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quota := func(ptr interface{}) { o.Quota = *ptr.(**fsrm.Quota) }
		if err := w.ReadPointer(&o.Quota, _s_quota, _ptr_quota); err != nil {
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

// CreateQuotaRequest structure represents the CreateQuota operation request
type CreateQuotaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path of the local directory for the directory quota. The maximum
	// length of this string MUST be 260 characters.
	Path *oaut.String `idl:"name:path" json:"path"`
}

func (o *CreateQuotaRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateQuotaOperation) *xxx_CreateQuotaOperation {
	if op == nil {
		op = &xxx_CreateQuotaOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Path = op.Path
	return op
}

func (o *CreateQuotaRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateQuotaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *CreateQuotaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateQuotaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateQuotaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateQuotaResponse structure represents the CreateQuota operation response
type CreateQuotaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// quota: Pointer to an IFsrmQuota interface pointer (section 3.2.4.2.16) that upon
	// completion points to the newly created Non-Persisted Directory Quota Instance. To
	// have the Non-Persisted Directory Quota Instance added to the server's List of Persisted
	// Directory Quota Instances (section 3.2.4.2.16.1), the caller MUST call Commit (section
	// 3.2.4.2.10.5).
	Quota *fsrm.Quota `idl:"name:quota" json:"quota"`
	// Return: The CreateQuota return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateQuotaResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateQuotaOperation) *xxx_CreateQuotaOperation {
	if op == nil {
		op = &xxx_CreateQuotaOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
	return op
}

func (o *CreateQuotaResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateQuotaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
}
func (o *CreateQuotaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateQuotaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateQuotaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateAutoApplyQuotaOperation structure represents the CreateAutoApplyQuota operation
type xxx_CreateAutoApplyQuotaOperation struct {
	This              *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat       `idl:"name:That" json:"that"`
	QuotaTemplateName *oaut.String         `idl:"name:quotaTemplateName" json:"quota_template_name"`
	Path              *oaut.String         `idl:"name:path" json:"path"`
	Quota             *fsrm.AutoApplyQuota `idl:"name:quota" json:"quota"`
	Return            int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateAutoApplyQuotaOperation) OpNum() int { return 10 }

func (o *xxx_CreateAutoApplyQuotaOperation) OpName() string {
	return "/IFsrmQuotaManager/v0/CreateAutoApplyQuota"
}

func (o *xxx_CreateAutoApplyQuotaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAutoApplyQuotaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// quotaTemplateName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QuotaTemplateName != nil {
			_ptr_quotaTemplateName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QuotaTemplateName != nil {
					if err := o.QuotaTemplateName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QuotaTemplateName, _ptr_quotaTemplateName); err != nil {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_CreateAutoApplyQuotaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// quotaTemplateName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_quotaTemplateName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QuotaTemplateName == nil {
				o.QuotaTemplateName = &oaut.String{}
			}
			if err := o.QuotaTemplateName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotaTemplateName := func(ptr interface{}) { o.QuotaTemplateName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QuotaTemplateName, _s_quotaTemplateName, _ptr_quotaTemplateName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAutoApplyQuotaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAutoApplyQuotaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quota {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmAutoApplyQuota}(interface))
	{
		if o.Quota != nil {
			_ptr_quota := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Quota != nil {
					if err := o.Quota.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.AutoApplyQuota{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Quota, _ptr_quota); err != nil {
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

func (o *xxx_CreateAutoApplyQuotaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quota {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmAutoApplyQuota}(interface))
	{
		_ptr_quota := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Quota == nil {
				o.Quota = &fsrm.AutoApplyQuota{}
			}
			if err := o.Quota.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quota := func(ptr interface{}) { o.Quota = *ptr.(**fsrm.AutoApplyQuota) }
		if err := w.ReadPointer(&o.Quota, _s_quota, _ptr_quota); err != nil {
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

// CreateAutoApplyQuotaRequest structure represents the CreateAutoApplyQuota operation request
type CreateAutoApplyQuotaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// quotaTemplateName: Contains the Name property of the directory quota template from
	// which to derive the auto apply quota. The maximum length of this string MUST be 4,000
	// characters.
	QuotaTemplateName *oaut.String `idl:"name:quotaTemplateName" json:"quota_template_name"`
	// path: Contains the path of the local directory for the auto apply quota. The maximum
	// length of this string MUST be 260 characters.
	Path *oaut.String `idl:"name:path" json:"path"`
}

func (o *CreateAutoApplyQuotaRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateAutoApplyQuotaOperation) *xxx_CreateAutoApplyQuotaOperation {
	if op == nil {
		op = &xxx_CreateAutoApplyQuotaOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.QuotaTemplateName = op.QuotaTemplateName
	o.Path = op.Path
	return op
}

func (o *CreateAutoApplyQuotaRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateAutoApplyQuotaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QuotaTemplateName = op.QuotaTemplateName
	o.Path = op.Path
}
func (o *CreateAutoApplyQuotaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateAutoApplyQuotaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateAutoApplyQuotaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateAutoApplyQuotaResponse structure represents the CreateAutoApplyQuota operation response
type CreateAutoApplyQuotaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// quota: Pointer to an IFsrmAutoApplyQuota interface pointer (section 3.2.4.2.17) that
	// upon completion points to the newly created Non-Persisted Auto Apply Quota Instance.
	// To have the Non-Persisted Auto Apply Quota Instance added to the server's List of
	// Persisted Auto Apply Quota Instances (section 3.2.1.2), the caller MUST call Commit
	// (section 3.2.4.2.17.1).
	Quota *fsrm.AutoApplyQuota `idl:"name:quota" json:"quota"`
	// Return: The CreateAutoApplyQuota return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateAutoApplyQuotaResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateAutoApplyQuotaOperation) *xxx_CreateAutoApplyQuotaOperation {
	if op == nil {
		op = &xxx_CreateAutoApplyQuotaOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
	return op
}

func (o *CreateAutoApplyQuotaResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateAutoApplyQuotaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
}
func (o *CreateAutoApplyQuotaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateAutoApplyQuotaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateAutoApplyQuotaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQuotaOperation structure represents the GetQuota operation
type xxx_GetQuotaOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   *oaut.String   `idl:"name:path" json:"path"`
	Quota  *fsrm.Quota    `idl:"name:quota" json:"quota"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQuotaOperation) OpNum() int { return 11 }

func (o *xxx_GetQuotaOperation) OpName() string { return "/IFsrmQuotaManager/v0/GetQuota" }

func (o *xxx_GetQuotaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_GetQuotaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quota {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmQuota}(interface))
	{
		if o.Quota != nil {
			_ptr_quota := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Quota != nil {
					if err := o.Quota.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Quota{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Quota, _ptr_quota); err != nil {
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

func (o *xxx_GetQuotaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quota {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmQuota}(interface))
	{
		_ptr_quota := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Quota == nil {
				o.Quota = &fsrm.Quota{}
			}
			if err := o.Quota.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quota := func(ptr interface{}) { o.Quota = *ptr.(**fsrm.Quota) }
		if err := w.ReadPointer(&o.Quota, _s_quota, _ptr_quota); err != nil {
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

// GetQuotaRequest structure represents the GetQuota operation request
type GetQuotaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path to the requested directory quota.
	Path *oaut.String `idl:"name:path" json:"path"`
}

func (o *GetQuotaRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaOperation) *xxx_GetQuotaOperation {
	if op == nil {
		op = &xxx_GetQuotaOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Path = op.Path
	return op
}

func (o *GetQuotaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *GetQuotaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQuotaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQuotaResponse structure represents the GetQuota operation response
type GetQuotaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// quota: Pointer to an IFsrmQuota interface pointer (section 3.2.4.2.16) that upon
	// completion points to the directory quota for the specified path. The caller MUST
	// release the quota when it is done with it.
	Quota *fsrm.Quota `idl:"name:quota" json:"quota"`
	// Return: The GetQuota return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQuotaResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaOperation) *xxx_GetQuotaOperation {
	if op == nil {
		op = &xxx_GetQuotaOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
	return op
}

func (o *GetQuotaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
}
func (o *GetQuotaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQuotaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAutoApplyQuotaOperation structure represents the GetAutoApplyQuota operation
type xxx_GetAutoApplyQuotaOperation struct {
	This   *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Path   *oaut.String         `idl:"name:path" json:"path"`
	Quota  *fsrm.AutoApplyQuota `idl:"name:quota" json:"quota"`
	Return int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAutoApplyQuotaOperation) OpNum() int { return 12 }

func (o *xxx_GetAutoApplyQuotaOperation) OpName() string {
	return "/IFsrmQuotaManager/v0/GetAutoApplyQuota"
}

func (o *xxx_GetAutoApplyQuotaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAutoApplyQuotaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_GetAutoApplyQuotaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAutoApplyQuotaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAutoApplyQuotaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quota {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmAutoApplyQuota}(interface))
	{
		if o.Quota != nil {
			_ptr_quota := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Quota != nil {
					if err := o.Quota.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.AutoApplyQuota{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Quota, _ptr_quota); err != nil {
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

func (o *xxx_GetAutoApplyQuotaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quota {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmAutoApplyQuota}(interface))
	{
		_ptr_quota := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Quota == nil {
				o.Quota = &fsrm.AutoApplyQuota{}
			}
			if err := o.Quota.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quota := func(ptr interface{}) { o.Quota = *ptr.(**fsrm.AutoApplyQuota) }
		if err := w.ReadPointer(&o.Quota, _s_quota, _ptr_quota); err != nil {
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

// GetAutoApplyQuotaRequest structure represents the GetAutoApplyQuota operation request
type GetAutoApplyQuotaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path to the requested auto apply quota.
	Path *oaut.String `idl:"name:path" json:"path"`
}

func (o *GetAutoApplyQuotaRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAutoApplyQuotaOperation) *xxx_GetAutoApplyQuotaOperation {
	if op == nil {
		op = &xxx_GetAutoApplyQuotaOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Path = op.Path
	return op
}

func (o *GetAutoApplyQuotaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAutoApplyQuotaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *GetAutoApplyQuotaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAutoApplyQuotaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAutoApplyQuotaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAutoApplyQuotaResponse structure represents the GetAutoApplyQuota operation response
type GetAutoApplyQuotaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// quota: Pointer to an IFsrmAutoApplyQuota interface pointer (section 3.2.4.2.17) that
	// upon completion points to the auto apply quota for the specified path. The caller
	// MUST release the quota when it is done with it.
	Quota *fsrm.AutoApplyQuota `idl:"name:quota" json:"quota"`
	// Return: The GetAutoApplyQuota return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAutoApplyQuotaResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAutoApplyQuotaOperation) *xxx_GetAutoApplyQuotaOperation {
	if op == nil {
		op = &xxx_GetAutoApplyQuotaOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
	return op
}

func (o *GetAutoApplyQuotaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAutoApplyQuotaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
}
func (o *GetAutoApplyQuotaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAutoApplyQuotaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAutoApplyQuotaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRestrictiveQuotaOperation structure represents the GetRestrictiveQuota operation
type xxx_GetRestrictiveQuotaOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   *oaut.String   `idl:"name:path" json:"path"`
	Quota  *fsrm.Quota    `idl:"name:quota" json:"quota"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRestrictiveQuotaOperation) OpNum() int { return 13 }

func (o *xxx_GetRestrictiveQuotaOperation) OpName() string {
	return "/IFsrmQuotaManager/v0/GetRestrictiveQuota"
}

func (o *xxx_GetRestrictiveQuotaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRestrictiveQuotaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_GetRestrictiveQuotaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRestrictiveQuotaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRestrictiveQuotaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quota {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmQuota}(interface))
	{
		if o.Quota != nil {
			_ptr_quota := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Quota != nil {
					if err := o.Quota.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Quota{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Quota, _ptr_quota); err != nil {
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

func (o *xxx_GetRestrictiveQuotaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quota {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmQuota}(interface))
	{
		_ptr_quota := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Quota == nil {
				o.Quota = &fsrm.Quota{}
			}
			if err := o.Quota.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quota := func(ptr interface{}) { o.Quota = *ptr.(**fsrm.Quota) }
		if err := w.ReadPointer(&o.Quota, _s_quota, _ptr_quota); err != nil {
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

// GetRestrictiveQuotaRequest structure represents the GetRestrictiveQuota operation request
type GetRestrictiveQuotaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path to return for the restrictive quota. The maximum length of
	// this string MUST be 260 characters.
	Path *oaut.String `idl:"name:path" json:"path"`
}

func (o *GetRestrictiveQuotaRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRestrictiveQuotaOperation) *xxx_GetRestrictiveQuotaOperation {
	if op == nil {
		op = &xxx_GetRestrictiveQuotaOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Path = op.Path
	return op
}

func (o *GetRestrictiveQuotaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRestrictiveQuotaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *GetRestrictiveQuotaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRestrictiveQuotaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRestrictiveQuotaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRestrictiveQuotaResponse structure represents the GetRestrictiveQuota operation response
type GetRestrictiveQuotaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// quota: Pointer to an IFsrmQuota interface pointer (section 3.2.4.2.16) that upon
	// completion points to the most restrictive quota for the specified path. The caller
	// MUST release the quota when it is done with it.
	Quota *fsrm.Quota `idl:"name:quota" json:"quota"`
	// Return: The GetRestrictiveQuota return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRestrictiveQuotaResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRestrictiveQuotaOperation) *xxx_GetRestrictiveQuotaOperation {
	if op == nil {
		op = &xxx_GetRestrictiveQuotaOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
	return op
}

func (o *GetRestrictiveQuotaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRestrictiveQuotaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
}
func (o *GetRestrictiveQuotaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRestrictiveQuotaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRestrictiveQuotaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumQuotasOperation structure represents the EnumQuotas operation
type xxx_EnumQuotasOperation struct {
	This    *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Path    *oaut.String                `idl:"name:path" json:"path"`
	Options fsrm.EnumOptions            `idl:"name:options" json:"options"`
	Quotas  *fsrm.CommittableCollection `idl:"name:quotas" json:"quotas"`
	Return  int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumQuotasOperation) OpNum() int { return 14 }

func (o *xxx_EnumQuotasOperation) OpName() string { return "/IFsrmQuotaManager/v0/EnumQuotas" }

func (o *xxx_EnumQuotasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumQuotasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumQuotasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumQuotasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumQuotasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quotas {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.Quotas != nil {
			_ptr_quotas := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Quotas != nil {
					if err := o.Quotas.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Quotas, _ptr_quotas); err != nil {
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

func (o *xxx_EnumQuotasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quotas {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_quotas := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Quotas == nil {
				o.Quotas = &fsrm.CommittableCollection{}
			}
			if err := o.Quotas.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotas := func(ptr interface{}) { o.Quotas = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.Quotas, _s_quotas, _ptr_quotas); err != nil {
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

// EnumQuotasRequest structure represents the EnumQuotas operation request
type EnumQuotasRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path indicating the search location for quota numeration.
	Path *oaut.String `idl:"name:path" json:"path"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the quotas.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumQuotasRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumQuotasOperation) *xxx_EnumQuotasOperation {
	if op == nil {
		op = &xxx_EnumQuotasOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Path = op.Path
	o.Options = op.Options
	return op
}

func (o *EnumQuotasRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumQuotasOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Options = op.Options
}
func (o *EnumQuotasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumQuotasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumQuotasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumQuotasResponse structure represents the EnumQuotas operation response
type EnumQuotasResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// quotas: Pointer to an IFsrmCommittableCollection interface pointer (section 3.2.4.2.3)
	// that upon completion contains pointers to directory quotas belonging to the path
	// specified based on the wildcard characters specified in the path. The caller MUST
	// release the collection when it is done with it.
	Quotas *fsrm.CommittableCollection `idl:"name:quotas" json:"quotas"`
	// Return: The EnumQuotas return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumQuotasResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumQuotasOperation) *xxx_EnumQuotasOperation {
	if op == nil {
		op = &xxx_EnumQuotasOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Quotas = op.Quotas
	o.Return = op.Return
	return op
}

func (o *EnumQuotasResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumQuotasOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Quotas = op.Quotas
	o.Return = op.Return
}
func (o *EnumQuotasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumQuotasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumQuotasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumAutoApplyQuotasOperation structure represents the EnumAutoApplyQuotas operation
type xxx_EnumAutoApplyQuotasOperation struct {
	This    *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Path    *oaut.String                `idl:"name:path" json:"path"`
	Options fsrm.EnumOptions            `idl:"name:options" json:"options"`
	Quotas  *fsrm.CommittableCollection `idl:"name:quotas" json:"quotas"`
	Return  int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumAutoApplyQuotasOperation) OpNum() int { return 15 }

func (o *xxx_EnumAutoApplyQuotasOperation) OpName() string {
	return "/IFsrmQuotaManager/v0/EnumAutoApplyQuotas"
}

func (o *xxx_EnumAutoApplyQuotasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAutoApplyQuotasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAutoApplyQuotasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAutoApplyQuotasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumAutoApplyQuotasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quotas {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.Quotas != nil {
			_ptr_quotas := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Quotas != nil {
					if err := o.Quotas.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Quotas, _ptr_quotas); err != nil {
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

func (o *xxx_EnumAutoApplyQuotasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quotas {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_quotas := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Quotas == nil {
				o.Quotas = &fsrm.CommittableCollection{}
			}
			if err := o.Quotas.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotas := func(ptr interface{}) { o.Quotas = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.Quotas, _s_quotas, _ptr_quotas); err != nil {
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

// EnumAutoApplyQuotasRequest structure represents the EnumAutoApplyQuotas operation request
type EnumAutoApplyQuotasRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path indicating the search location for quota enumeration.
	Path *oaut.String `idl:"name:path" json:"path"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the quotas.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumAutoApplyQuotasRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumAutoApplyQuotasOperation) *xxx_EnumAutoApplyQuotasOperation {
	if op == nil {
		op = &xxx_EnumAutoApplyQuotasOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Path = op.Path
	o.Options = op.Options
	return op
}

func (o *EnumAutoApplyQuotasRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumAutoApplyQuotasOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Options = op.Options
}
func (o *EnumAutoApplyQuotasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumAutoApplyQuotasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumAutoApplyQuotasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumAutoApplyQuotasResponse structure represents the EnumAutoApplyQuotas operation response
type EnumAutoApplyQuotasResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// quotas: Pointer to an IFsrmCommittableCollection interface pointer (section 3.2.4.2.3)
	// that upon completion contains pointers to auto apply quotas belonging to the specified
	// path based on the wildcard specified in path. The caller MUST release the collection
	// when it is done with it.
	Quotas *fsrm.CommittableCollection `idl:"name:quotas" json:"quotas"`
	// Return: The EnumAutoApplyQuotas return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumAutoApplyQuotasResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumAutoApplyQuotasOperation) *xxx_EnumAutoApplyQuotasOperation {
	if op == nil {
		op = &xxx_EnumAutoApplyQuotasOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Quotas = op.Quotas
	o.Return = op.Return
	return op
}

func (o *EnumAutoApplyQuotasResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumAutoApplyQuotasOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Quotas = op.Quotas
	o.Return = op.Return
}
func (o *EnumAutoApplyQuotasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumAutoApplyQuotasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumAutoApplyQuotasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumEffectiveQuotasOperation structure represents the EnumEffectiveQuotas operation
type xxx_EnumEffectiveQuotasOperation struct {
	This    *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Path    *oaut.String                `idl:"name:path" json:"path"`
	Options fsrm.EnumOptions            `idl:"name:options" json:"options"`
	Quotas  *fsrm.CommittableCollection `idl:"name:quotas" json:"quotas"`
	Return  int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumEffectiveQuotasOperation) OpNum() int { return 16 }

func (o *xxx_EnumEffectiveQuotasOperation) OpName() string {
	return "/IFsrmQuotaManager/v0/EnumEffectiveQuotas"
}

func (o *xxx_EnumEffectiveQuotasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumEffectiveQuotasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumEffectiveQuotasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumEffectiveQuotasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumEffectiveQuotasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quotas {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.Quotas != nil {
			_ptr_quotas := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Quotas != nil {
					if err := o.Quotas.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Quotas, _ptr_quotas); err != nil {
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

func (o *xxx_EnumEffectiveQuotasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quotas {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_quotas := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Quotas == nil {
				o.Quotas = &fsrm.CommittableCollection{}
			}
			if err := o.Quotas.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotas := func(ptr interface{}) { o.Quotas = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.Quotas, _s_quotas, _ptr_quotas); err != nil {
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

// EnumEffectiveQuotasRequest structure represents the EnumEffectiveQuotas operation request
type EnumEffectiveQuotasRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path to return for the quotas. The maximum length of this string
	// MUST be 260 characters.
	Path *oaut.String `idl:"name:path" json:"path"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the quotas.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumEffectiveQuotasRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumEffectiveQuotasOperation) *xxx_EnumEffectiveQuotasOperation {
	if op == nil {
		op = &xxx_EnumEffectiveQuotasOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Path = op.Path
	o.Options = op.Options
	return op
}

func (o *EnumEffectiveQuotasRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumEffectiveQuotasOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Options = op.Options
}
func (o *EnumEffectiveQuotasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumEffectiveQuotasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumEffectiveQuotasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumEffectiveQuotasResponse structure represents the EnumEffectiveQuotas operation response
type EnumEffectiveQuotasResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// quotas: Pointer to an IFsrmCommittableCollection interface pointer (section 3.2.4.2.3)
	// that upon completion contains pointers to every directory quota that affects the
	// specified path. The caller MUST release the collection when it is done with it.
	Quotas *fsrm.CommittableCollection `idl:"name:quotas" json:"quotas"`
	// Return: The EnumEffectiveQuotas return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumEffectiveQuotasResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumEffectiveQuotasOperation) *xxx_EnumEffectiveQuotasOperation {
	if op == nil {
		op = &xxx_EnumEffectiveQuotasOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Quotas = op.Quotas
	o.Return = op.Return
	return op
}

func (o *EnumEffectiveQuotasResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumEffectiveQuotasOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Quotas = op.Quotas
	o.Return = op.Return
}
func (o *EnumEffectiveQuotasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumEffectiveQuotasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumEffectiveQuotasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ScanOperation structure represents the Scan operation
type xxx_ScanOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   *oaut.String   `idl:"name:strPath" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ScanOperation) OpNum() int { return 17 }

func (o *xxx_ScanOperation) OpName() string { return "/IFsrmQuotaManager/v0/Scan" }

func (o *xxx_ScanOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScanOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_strPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_strPath); err != nil {
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

func (o *xxx_ScanOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strPath := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_strPath, _ptr_strPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScanOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScanOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ScanOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ScanRequest structure represents the Scan operation request
type ScanRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// strPath: Contains the path at which to start the quota scan. The maximum length of
	// this string MUST be 260 characters.
	Path *oaut.String `idl:"name:strPath" json:"path"`
}

func (o *ScanRequest) xxx_ToOp(ctx context.Context, op *xxx_ScanOperation) *xxx_ScanOperation {
	if op == nil {
		op = &xxx_ScanOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Path = op.Path
	return op
}

func (o *ScanRequest) xxx_FromOp(ctx context.Context, op *xxx_ScanOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *ScanRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ScanRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ScanOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ScanResponse structure represents the Scan operation response
type ScanResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Scan return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ScanResponse) xxx_ToOp(ctx context.Context, op *xxx_ScanOperation) *xxx_ScanOperation {
	if op == nil {
		op = &xxx_ScanOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *ScanResponse) xxx_FromOp(ctx context.Context, op *xxx_ScanOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ScanResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ScanResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ScanOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateQuotaCollectionOperation structure represents the CreateQuotaCollection operation
type xxx_CreateQuotaCollectionOperation struct {
	This       *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Collection *fsrm.CommittableCollection `idl:"name:collection" json:"collection"`
	Return     int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateQuotaCollectionOperation) OpNum() int { return 18 }

func (o *xxx_CreateQuotaCollectionOperation) OpName() string {
	return "/IFsrmQuotaManager/v0/CreateQuotaCollection"
}

func (o *xxx_CreateQuotaCollectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateQuotaCollectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateQuotaCollectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CreateQuotaCollectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateQuotaCollectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// collection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.Collection != nil {
			_ptr_collection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collection != nil {
					if err := o.Collection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collection, _ptr_collection); err != nil {
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

func (o *xxx_CreateQuotaCollectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// collection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_collection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collection == nil {
				o.Collection = &fsrm.CommittableCollection{}
			}
			if err := o.Collection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_collection := func(ptr interface{}) { o.Collection = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.Collection, _s_collection, _ptr_collection); err != nil {
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

// CreateQuotaCollectionRequest structure represents the CreateQuotaCollection operation request
type CreateQuotaCollectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreateQuotaCollectionRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateQuotaCollectionOperation) *xxx_CreateQuotaCollectionOperation {
	if op == nil {
		op = &xxx_CreateQuotaCollectionOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *CreateQuotaCollectionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateQuotaCollectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreateQuotaCollectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateQuotaCollectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateQuotaCollectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateQuotaCollectionResponse structure represents the CreateQuotaCollection operation response
type CreateQuotaCollectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// collection: Pointer to an IFsrmCommittableCollection interface pointer (section 3.2.4.2.3)
	// that upon completion points to an empty IFsrmCommittableCollection specific to quota
	// objects. A caller MUST release the collection received when it is done with it.
	Collection *fsrm.CommittableCollection `idl:"name:collection" json:"collection"`
	// Return: The CreateQuotaCollection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateQuotaCollectionResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateQuotaCollectionOperation) *xxx_CreateQuotaCollectionOperation {
	if op == nil {
		op = &xxx_CreateQuotaCollectionOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Collection = op.Collection
	o.Return = op.Return
	return op
}

func (o *CreateQuotaCollectionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateQuotaCollectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collection = op.Collection
	o.Return = op.Return
}
func (o *CreateQuotaCollectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateQuotaCollectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateQuotaCollectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
