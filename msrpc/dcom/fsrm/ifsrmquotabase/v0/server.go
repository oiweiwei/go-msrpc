package ifsrmquotabase

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
		in := &GetQuotaLimitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetQuotaLimit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // QuotaLimit
		in := &SetQuotaLimitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetQuotaLimit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // QuotaFlags
		in := &GetQuotaFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetQuotaFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // QuotaFlags
		in := &SetQuotaFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetQuotaFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // Thresholds
		in := &GetThresholdsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetThresholds(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // AddThreshold
		in := &AddThresholdRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddThreshold(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // DeleteThreshold
		in := &DeleteThresholdRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteThreshold(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // ModifyThreshold
		in := &ModifyThresholdRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ModifyThreshold(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // CreateThresholdAction
		in := &CreateThresholdActionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateThresholdAction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // EnumThresholdActions
		in := &EnumThresholdActionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumThresholdActions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
