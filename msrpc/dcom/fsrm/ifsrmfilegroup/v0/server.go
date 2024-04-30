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
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Name
		in := &SetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // Members
		in := &GetMembersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMembers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Members
		in := &SetMembersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMembers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // NonMembers
		in := &GetNonMembersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNonMembers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // NonMembers
		in := &SetNonMembersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetNonMembers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
