package ifsrmfilescreenbase

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
		op := &xxx_GetBlockedFileGroupsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBlockedFileGroupsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBlockedFileGroups(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // BlockedFileGroups
		op := &xxx_SetBlockedFileGroupsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetBlockedFileGroupsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetBlockedFileGroups(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // FileScreenFlags
		op := &xxx_GetFileScreenFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileScreenFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileScreenFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // FileScreenFlags
		op := &xxx_SetFileScreenFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFileScreenFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFileScreenFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // CreateAction
		op := &xxx_CreateActionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateActionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateAction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // EnumActions
		op := &xxx_EnumActionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumActionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumActions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmFileScreenBase
type UnimplementedFileScreenBaseServer struct {
	ifsrmobject.UnimplementedObjectServer
}

func (UnimplementedFileScreenBaseServer) GetBlockedFileGroups(context.Context, *GetBlockedFileGroupsRequest) (*GetBlockedFileGroupsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenBaseServer) SetBlockedFileGroups(context.Context, *SetBlockedFileGroupsRequest) (*SetBlockedFileGroupsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenBaseServer) GetFileScreenFlags(context.Context, *GetFileScreenFlagsRequest) (*GetFileScreenFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenBaseServer) SetFileScreenFlags(context.Context, *SetFileScreenFlagsRequest) (*SetFileScreenFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenBaseServer) CreateAction(context.Context, *CreateActionRequest) (*CreateActionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenBaseServer) EnumActions(context.Context, *EnumActionsRequest) (*EnumActionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FileScreenBaseServer = (*UnimplementedFileScreenBaseServer)(nil)
