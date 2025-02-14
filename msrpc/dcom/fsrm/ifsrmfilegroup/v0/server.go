package ifsrmfilegroup

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

// IFsrmFileGroup server interface.
type FileGroupServer interface {

	// IFsrmObject base class.
	ifsrmobject.ObjectServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest) (*SetNameResponse, error)

	// Members operation.
	GetMembers(context.Context, *GetMembersRequest) (*GetMembersResponse, error)

	// Members operation.
	SetMembers(context.Context, *SetMembersRequest) (*SetMembersResponse, error)

	// NonMembers operation.
	GetNonMembers(context.Context, *GetNonMembersRequest) (*GetNonMembersResponse, error)

	// NonMembers operation.
	SetNonMembers(context.Context, *SetNonMembersRequest) (*SetNonMembersResponse, error)
}

func RegisterFileGroupServer(conn dcerpc.Conn, o FileGroupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileGroupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileGroupSyntaxV0_0))...)
}

func NewFileGroupServerHandle(o FileGroupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileGroupServerHandle(ctx, o, opNum, r)
	}
}

func FileGroupServerHandle(ctx context.Context, o FileGroupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmObject base method.
		return ifsrmobject.ObjectServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Name
		op := &xxx_SetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Members
		op := &xxx_GetMembersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMembersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMembers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Members
		op := &xxx_SetMembersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMembersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMembers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // NonMembers
		op := &xxx_GetNonMembersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNonMembersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNonMembers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // NonMembers
		op := &xxx_SetNonMembersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNonMembersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNonMembers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmFileGroup
type UnimplementedFileGroupServer struct {
	ifsrmobject.UnimplementedObjectServer
}

func (UnimplementedFileGroupServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileGroupServer) SetName(context.Context, *SetNameRequest) (*SetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileGroupServer) GetMembers(context.Context, *GetMembersRequest) (*GetMembersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileGroupServer) SetMembers(context.Context, *SetMembersRequest) (*SetMembersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileGroupServer) GetNonMembers(context.Context, *GetNonMembersRequest) (*GetNonMembersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileGroupServer) SetNonMembers(context.Context, *SetNonMembersRequest) (*SetNonMembersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FileGroupServer = (*UnimplementedFileGroupServer)(nil)
