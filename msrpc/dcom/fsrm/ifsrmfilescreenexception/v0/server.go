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
		in := &GetPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // AllowedFileGroups
		in := &GetAllowedFileGroupsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllowedFileGroups(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // AllowedFileGroups
		in := &SetAllowedFileGroupsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAllowedFileGroups(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
