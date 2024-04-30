package ifsrmquotamanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = idispatch.GoPackage
)

// IFsrmQuotaManager server interface.
type QuotaManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// ActionVariables operation.
	GetActionVariables(context.Context, *GetActionVariablesRequest) (*GetActionVariablesResponse, error)

	// ActionVariableDescriptions operation.
	GetActionVariableDescriptions(context.Context, *GetActionVariableDescriptionsRequest) (*GetActionVariableDescriptionsResponse, error)

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
	CreateQuota(context.Context, *CreateQuotaRequest) (*CreateQuotaResponse, error)

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
	CreateAutoApplyQuota(context.Context, *CreateAutoApplyQuotaRequest) (*CreateAutoApplyQuotaResponse, error)

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
	GetQuota(context.Context, *GetQuotaRequest) (*GetQuotaResponse, error)

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
	GetAutoApplyQuota(context.Context, *GetAutoApplyQuotaRequest) (*GetAutoApplyQuotaResponse, error)

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
	GetRestrictiveQuota(context.Context, *GetRestrictiveQuotaRequest) (*GetRestrictiveQuotaResponse, error)

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
	EnumQuotas(context.Context, *EnumQuotasRequest) (*EnumQuotasResponse, error)

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
	EnumAutoApplyQuotas(context.Context, *EnumAutoApplyQuotasRequest) (*EnumAutoApplyQuotasResponse, error)

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
	EnumEffectiveQuotas(context.Context, *EnumEffectiveQuotasRequest) (*EnumEffectiveQuotasResponse, error)

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
	Scan(context.Context, *ScanRequest) (*ScanResponse, error)

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
	CreateQuotaCollection(context.Context, *CreateQuotaCollectionRequest) (*CreateQuotaCollectionResponse, error)
}

func RegisterQuotaManagerServer(conn dcerpc.Conn, o QuotaManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuotaManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QuotaManagerSyntaxV0_0))...)
}

func NewQuotaManagerServerHandle(o QuotaManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QuotaManagerServerHandle(ctx, o, opNum, r)
	}
}

func QuotaManagerServerHandle(ctx context.Context, o QuotaManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // ActionVariables
		in := &GetActionVariablesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetActionVariables(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // ActionVariableDescriptions
		in := &GetActionVariableDescriptionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetActionVariableDescriptions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // CreateQuota
		in := &CreateQuotaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateQuota(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // CreateAutoApplyQuota
		in := &CreateAutoApplyQuotaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateAutoApplyQuota(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // GetQuota
		in := &GetQuotaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetQuota(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // GetAutoApplyQuota
		in := &GetAutoApplyQuotaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAutoApplyQuota(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // GetRestrictiveQuota
		in := &GetRestrictiveQuotaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRestrictiveQuota(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // EnumQuotas
		in := &EnumQuotasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumQuotas(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // EnumAutoApplyQuotas
		in := &EnumAutoApplyQuotasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumAutoApplyQuotas(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // EnumEffectiveQuotas
		in := &EnumEffectiveQuotasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumEffectiveQuotas(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // Scan
		in := &ScanRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Scan(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // CreateQuotaCollection
		in := &CreateQuotaCollectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateQuotaCollection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
