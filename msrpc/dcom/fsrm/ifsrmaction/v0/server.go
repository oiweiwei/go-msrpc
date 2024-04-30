package ifsrmaction

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

// IFsrmAction server interface.
type ActionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Id operation.
	GetID(context.Context, *GetIDRequest) (*GetIDResponse, error)

	// The ActionType (get) method returns the read-only action type property of the action.
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
	//	| 0x80070057 E_INVALIDARG | The actionType parameter is NULL. |
	//	+-------------------------+-----------------------------------+
	GetActionType(context.Context, *GetActionTypeRequest) (*GetActionTypeResponse, error)

	// RunLimitInterval operation.
	GetRunLimitInterval(context.Context, *GetRunLimitIntervalRequest) (*GetRunLimitIntervalResponse, error)

	// RunLimitInterval operation.
	SetRunLimitInterval(context.Context, *SetRunLimitIntervalRequest) (*SetRunLimitIntervalResponse, error)

	// Delete operation.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterActionServer(conn dcerpc.Conn, o ActionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewActionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ActionSyntaxV0_0))...)
}

func NewActionServerHandle(o ActionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ActionServerHandle(ctx, o, opNum, r)
	}
}

func ActionServerHandle(ctx context.Context, o ActionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Id
		in := &GetIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // ActionType
		in := &GetActionTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetActionType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // RunLimitInterval
		in := &GetRunLimitIntervalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRunLimitInterval(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // RunLimitInterval
		in := &SetRunLimitIntervalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetRunLimitInterval(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Delete
		in := &DeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Delete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
