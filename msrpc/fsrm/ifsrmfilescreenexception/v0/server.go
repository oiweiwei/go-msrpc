package ifsrmfilescreenexception

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

// IFsrmFileScreenException server interface.
type FileScreenExceptionServer interface {

	// IFsrmObject base class.
	ifsrmobject.ObjectServer

	// Path operation.
	GetPath(context.Context, *GetPathRequest) (*GetPathResponse, error)

	// AllowedFileGroups operation.
	GetAllowedFileGroups(context.Context, *GetAllowedFileGroupsRequest) (*GetAllowedFileGroupsResponse, error)

	// AllowedFileGroups operation.
	SetAllowedFileGroups(context.Context, *SetAllowedFileGroupsRequest) (*SetAllowedFileGroupsResponse, error)
}

func RegisterFileScreenExceptionServer(conn dcerpc.Conn, o FileScreenExceptionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileScreenExceptionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileScreenExceptionSyntaxV0_0))...)
}

func NewFileScreenExceptionServerHandle(o FileScreenExceptionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileScreenExceptionServerHandle(ctx, o, opNum, r)
	}
}

func FileScreenExceptionServerHandle(ctx context.Context, o FileScreenExceptionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmObject base method.
		return ifsrmobject.ObjectServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // Path
		op := &xxx_GetPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // AllowedFileGroups
		op := &xxx_GetAllowedFileGroupsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllowedFileGroupsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllowedFileGroups(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // AllowedFileGroups
		op := &xxx_SetAllowedFileGroupsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAllowedFileGroupsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAllowedFileGroups(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmFileScreenException
type UnimplementedFileScreenExceptionServer struct {
	ifsrmobject.UnimplementedObjectServer
}

func (UnimplementedFileScreenExceptionServer) GetPath(context.Context, *GetPathRequest) (*GetPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenExceptionServer) GetAllowedFileGroups(context.Context, *GetAllowedFileGroupsRequest) (*GetAllowedFileGroupsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileScreenExceptionServer) SetAllowedFileGroups(context.Context, *SetAllowedFileGroupsRequest) (*SetAllowedFileGroupsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FileScreenExceptionServer = (*UnimplementedFileScreenExceptionServer)(nil)
