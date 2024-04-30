package ifsrmpropertycondition

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

// IFsrmPropertyCondition server interface.
type PropertyConditionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest) (*SetNameResponse, error)

	// Type operation.
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

	// Type operation.
	SetType(context.Context, *SetTypeRequest) (*SetTypeResponse, error)

	// Value operation.
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)

	// Value operation.
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)

	// Delete operation.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterPropertyConditionServer(conn dcerpc.Conn, o PropertyConditionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPropertyConditionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PropertyConditionSyntaxV1_0))...)
}

func NewPropertyConditionServerHandle(o PropertyConditionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PropertyConditionServerHandle(ctx, o, opNum, r)
	}
}

func PropertyConditionServerHandle(ctx context.Context, o PropertyConditionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Name
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Name
		in := &SetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Type
		in := &GetTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Type
		in := &SetTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Value
		in := &GetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Value
		in := &SetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Delete
		in := &DeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Delete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
