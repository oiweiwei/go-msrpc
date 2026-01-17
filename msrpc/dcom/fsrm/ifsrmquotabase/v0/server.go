package ifsrmquotabase

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmobject/v0"
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
	_ = ifsrmobject.GoPackage
)

// IFsrmQuotaBase server interface.
type QuotaBaseServer interface {

	// IFsrmObject base class.
	ifsrmobject.ObjectServer

	// QuotaLimit operation.
	GetQuotaLimit(context.Context, *GetQuotaLimitRequest) (*GetQuotaLimitResponse, error)

	// QuotaLimit operation.
	SetQuotaLimit(context.Context, *SetQuotaLimitRequest) (*SetQuotaLimitResponse, error)

	// QuotaFlags operation.
	GetQuotaFlags(context.Context, *GetQuotaFlagsRequest) (*GetQuotaFlagsResponse, error)

	// QuotaFlags operation.
	SetQuotaFlags(context.Context, *SetQuotaFlagsRequest) (*SetQuotaFlagsResponse, error)

	// The Thresholds (get) method returns the thresholds for the quota object.
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
	//	| 0x80070057 E_INVALIDARG | The thresholds parameter is NULL. |
	//	+-------------------------+-----------------------------------+
	GetThresholds(context.Context, *GetThresholdsRequest) (*GetThresholdsResponse, error)

	// The AddThreshold method adds a threshold to the quota objects list of thresholds.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------+
	//	|              RETURN              |                                                                            |
	//	|            VALUE/CODE            |                                DESCRIPTION                                 |
	//	|                                  |                                                                            |
	//	+----------------------------------+----------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | The object being created already exists.                                   |
	//	+----------------------------------+----------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE   | The content of the threshold parameter is less than 1 or greater than 250. |
	//	+----------------------------------+----------------------------------------------------------------------------+
	AddThreshold(context.Context, *AddThresholdRequest) (*AddThresholdResponse, error)

	// The DeleteThreshold method deletes a threshold from the quota objects list of thresholds.
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
	//	| 0x80045301 FSRM_E_NOT_FOUND    | No threshold was found on the quota object equal to the value specified in the   |
	//	|                                | threshold parameter.                                                             |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE | The content of the threshold parameter is less than 1 or greater than 250.       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	DeleteThreshold(context.Context, *DeleteThresholdRequest) (*DeleteThresholdResponse, error)

	// The ModifyThreshold method modifies the disk usage percentage of an existing threshold
	// in the quota object's list of thresholds.
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
	//	| 0x80045301 FSRM_E_NOT_FOUND      | No threshold was found on the quota object equal to the value specified in the   |
	//	|                                  | threshold parameter.                                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | A threshold equal to the newThreshold parameter already exists for the quota     |
	//	|                                  | object.                                                                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE   | The content of the threshold parameter is less than 1 or greater than 250.       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	ModifyThreshold(context.Context, *ModifyThresholdRequest) (*ModifyThresholdResponse, error)

	// The CreateThresholdAction method creates an action and associates it with the specified
	// threshold.
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
	//	| 0x80045301 FSRM_E_NOT_FOUND      | An object with the specified threshold could not be found.                       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | An action of the same type as the actionType parameter already exists for the    |
	//	|                                  | threshold.                                                                       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE   | The content of the threshold parameter is less than 1 or greater than 250.       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The action parameter is NULL.   |
	//	|                                  | The actionsType parameter is not a valid type.                                   |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	CreateThresholdAction(context.Context, *CreateThresholdActionRequest) (*CreateThresholdActionResponse, error)

	// The EnumThresholdActions method enumerates all the actions for the specified threshold.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+----------------------------------------------------------------------------+
	//	|             RETURN             |                                                                            |
	//	|           VALUE/CODE           |                                DESCRIPTION                                 |
	//	|                                |                                                                            |
	//	+--------------------------------+----------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND    | An object with the specified threshold could not be found.                 |
	//	+--------------------------------+----------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE | The content of the threshold parameter is less than 1 or greater than 250. |
	//	+--------------------------------+----------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | The actions parameter is NULL.                                             |
	//	+--------------------------------+----------------------------------------------------------------------------+
	EnumThresholdActions(context.Context, *EnumThresholdActionsRequest) (*EnumThresholdActionsResponse, error)
}

func RegisterQuotaBaseServer(conn dcerpc.Conn, o QuotaBaseServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuotaBaseServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QuotaBaseSyntaxV0_0))...)
}

func NewQuotaBaseServerHandle(o QuotaBaseServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QuotaBaseServerHandle(ctx, o, opNum, r)
	}
}

func QuotaBaseServerHandle(ctx context.Context, o QuotaBaseServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmObject base method.
		return ifsrmobject.ObjectServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // QuotaLimit
		op := &xxx_GetQuotaLimitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQuotaLimitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQuotaLimit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // QuotaLimit
		op := &xxx_SetQuotaLimitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetQuotaLimitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetQuotaLimit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // QuotaFlags
		op := &xxx_GetQuotaFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQuotaFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQuotaFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // QuotaFlags
		op := &xxx_SetQuotaFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetQuotaFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetQuotaFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Thresholds
		op := &xxx_GetThresholdsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetThresholdsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetThresholds(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // AddThreshold
		op := &xxx_AddThresholdOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddThresholdRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddThreshold(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // DeleteThreshold
		op := &xxx_DeleteThresholdOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteThresholdRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteThreshold(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // ModifyThreshold
		op := &xxx_ModifyThresholdOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyThresholdRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyThreshold(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // CreateThresholdAction
		op := &xxx_CreateThresholdActionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateThresholdActionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateThresholdAction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // EnumThresholdActions
		op := &xxx_EnumThresholdActionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumThresholdActionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumThresholdActions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmQuotaBase
type UnimplementedQuotaBaseServer struct {
	ifsrmobject.UnimplementedObjectServer
}

func (UnimplementedQuotaBaseServer) GetQuotaLimit(context.Context, *GetQuotaLimitRequest) (*GetQuotaLimitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaBaseServer) SetQuotaLimit(context.Context, *SetQuotaLimitRequest) (*SetQuotaLimitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaBaseServer) GetQuotaFlags(context.Context, *GetQuotaFlagsRequest) (*GetQuotaFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaBaseServer) SetQuotaFlags(context.Context, *SetQuotaFlagsRequest) (*SetQuotaFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaBaseServer) GetThresholds(context.Context, *GetThresholdsRequest) (*GetThresholdsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaBaseServer) AddThreshold(context.Context, *AddThresholdRequest) (*AddThresholdResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaBaseServer) DeleteThreshold(context.Context, *DeleteThresholdRequest) (*DeleteThresholdResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaBaseServer) ModifyThreshold(context.Context, *ModifyThresholdRequest) (*ModifyThresholdResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaBaseServer) CreateThresholdAction(context.Context, *CreateThresholdActionRequest) (*CreateThresholdActionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaBaseServer) EnumThresholdActions(context.Context, *EnumThresholdActionsRequest) (*EnumThresholdActionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QuotaBaseServer = (*UnimplementedQuotaBaseServer)(nil)
