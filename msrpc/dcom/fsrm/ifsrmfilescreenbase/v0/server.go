package ifsrmfilescreenbase

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

// IFsrmFileScreenBase server interface.
type FileScreenBaseServer interface {

	// IFsrmObject base class.
	ifsrmobject.ObjectServer

	// BlockedFileGroups operation.
	GetBlockedFileGroups(context.Context, *GetBlockedFileGroupsRequest) (*GetBlockedFileGroupsResponse, error)

	// BlockedFileGroups operation.
	SetBlockedFileGroups(context.Context, *SetBlockedFileGroupsRequest) (*SetBlockedFileGroupsResponse, error)

	// FileScreenFlags operation.
	GetFileScreenFlags(context.Context, *GetFileScreenFlagsRequest) (*GetFileScreenFlagsResponse, error)

	// FileScreenFlags operation.
	SetFileScreenFlags(context.Context, *SetFileScreenFlagsRequest) (*SetFileScreenFlagsResponse, error)

	// The CreateAction method creates an action for this file screen object.
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
	//	| 0x80045303 FSRM_E_ALREADY_EXISTS | The object already exists.                                                       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The action parameter            |
	//	|                                  | is NULL. The actionType parameter is not a valid type. If actionType is          |
	//	|                                  | FsrmActionType_Unknown, the parameter MUST be considered an invalid value.       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	CreateAction(context.Context, *CreateActionRequest) (*CreateActionResponse, error)

	// The EnumActions method enumerates all the actions for the file screen object.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+--------------------------------+
	//	|         RETURN          |                                |
	//	|       VALUE/CODE        |          DESCRIPTION           |
	//	|                         |                                |
	//	+-------------------------+--------------------------------+
	//	+-------------------------+--------------------------------+
	//	| 0x80070057 E_INVALIDARG | The actions parameter is NULL. |
	//	+-------------------------+--------------------------------+
	EnumActions(context.Context, *EnumActionsRequest) (*EnumActionsResponse, error)
}

func RegisterFileScreenBaseServer(conn dcerpc.Conn, o FileScreenBaseServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileScreenBaseServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileScreenBaseSyntaxV0_0))...)
}

func NewFileScreenBaseServerHandle(o FileScreenBaseServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileScreenBaseServerHandle(ctx, o, opNum, r)
	}
}

func FileScreenBaseServerHandle(ctx context.Context, o FileScreenBaseServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmObject base method.
		return ifsrmobject.ObjectServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // BlockedFileGroups
		in := &GetBlockedFileGroupsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetBlockedFileGroups(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // BlockedFileGroups
		in := &SetBlockedFileGroupsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetBlockedFileGroups(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // FileScreenFlags
		in := &GetFileScreenFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileScreenFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // FileScreenFlags
		in := &SetFileScreenFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileScreenFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // CreateAction
		in := &CreateActionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateAction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // EnumActions
		in := &EnumActionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumActions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
