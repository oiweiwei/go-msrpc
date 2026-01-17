package ifsrmquotamanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_GetActionVariablesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetActionVariablesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetActionVariables(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ActionVariableDescriptions
		op := &xxx_GetActionVariableDescriptionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetActionVariableDescriptionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetActionVariableDescriptions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // CreateQuota
		op := &xxx_CreateQuotaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateQuotaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateQuota(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // CreateAutoApplyQuota
		op := &xxx_CreateAutoApplyQuotaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateAutoApplyQuotaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateAutoApplyQuota(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // GetQuota
		op := &xxx_GetQuotaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQuotaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQuota(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // GetAutoApplyQuota
		op := &xxx_GetAutoApplyQuotaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAutoApplyQuotaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAutoApplyQuota(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetRestrictiveQuota
		op := &xxx_GetRestrictiveQuotaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRestrictiveQuotaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRestrictiveQuota(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // EnumQuotas
		op := &xxx_EnumQuotasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumQuotasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumQuotas(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // EnumAutoApplyQuotas
		op := &xxx_EnumAutoApplyQuotasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAutoApplyQuotasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAutoApplyQuotas(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // EnumEffectiveQuotas
		op := &xxx_EnumEffectiveQuotasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumEffectiveQuotasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumEffectiveQuotas(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Scan
		op := &xxx_ScanOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ScanRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Scan(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // CreateQuotaCollection
		op := &xxx_CreateQuotaCollectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateQuotaCollectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateQuotaCollection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmQuotaManager
type UnimplementedQuotaManagerServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedQuotaManagerServer) GetActionVariables(context.Context, *GetActionVariablesRequest) (*GetActionVariablesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) GetActionVariableDescriptions(context.Context, *GetActionVariableDescriptionsRequest) (*GetActionVariableDescriptionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) CreateQuota(context.Context, *CreateQuotaRequest) (*CreateQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) CreateAutoApplyQuota(context.Context, *CreateAutoApplyQuotaRequest) (*CreateAutoApplyQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) GetQuota(context.Context, *GetQuotaRequest) (*GetQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) GetAutoApplyQuota(context.Context, *GetAutoApplyQuotaRequest) (*GetAutoApplyQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) GetRestrictiveQuota(context.Context, *GetRestrictiveQuotaRequest) (*GetRestrictiveQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) EnumQuotas(context.Context, *EnumQuotasRequest) (*EnumQuotasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) EnumAutoApplyQuotas(context.Context, *EnumAutoApplyQuotasRequest) (*EnumAutoApplyQuotasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) EnumEffectiveQuotas(context.Context, *EnumEffectiveQuotasRequest) (*EnumEffectiveQuotasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) Scan(context.Context, *ScanRequest) (*ScanResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaManagerServer) CreateQuotaCollection(context.Context, *CreateQuotaCollectionRequest) (*CreateQuotaCollectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QuotaManagerServer = (*UnimplementedQuotaManagerServer)(nil)
