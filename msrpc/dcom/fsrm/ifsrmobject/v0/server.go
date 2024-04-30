package ifsrmobject

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

// IFsrmObject server interface.
type ObjectServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Id operation.
	GetID(context.Context, *GetIDRequest) (*GetIDResponse, error)

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	// Description operation.
	SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error)

	// Delete operation.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)

	// Commit operation.
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)
}

func RegisterObjectServer(conn dcerpc.Conn, o ObjectServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewObjectServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ObjectSyntaxV0_0))...)
}

func NewObjectServerHandle(o ObjectServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ObjectServerHandle(ctx, o, opNum, r)
	}
}

func ObjectServerHandle(ctx context.Context, o ObjectServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 8: // Description
		in := &GetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Description
		in := &SetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Delete
		in := &DeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Delete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Commit
		in := &CommitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Commit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
