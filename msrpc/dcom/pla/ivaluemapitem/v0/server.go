package ivaluemapitem

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

// IValueMapItem server interface.
type ValueMapItemServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	// Description operation.
	SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error)

	// Enabled operation.
	GetEnabled(context.Context, *GetEnabledRequest) (*GetEnabledResponse, error)

	// Enabled operation.
	SetEnabled(context.Context, *SetEnabledRequest) (*SetEnabledResponse, error)

	// Key operation.
	GetKey(context.Context, *GetKeyRequest) (*GetKeyResponse, error)

	// Key operation.
	SetKey(context.Context, *SetKeyRequest) (*SetKeyResponse, error)

	// Value operation.
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)

	// Value operation.
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)

	// The ValueMapType enumeration defines a value map type. A value map defines a named-value
	// pair. A value map can be used in different ways. A value map type defines which way
	// the value map is to be used; each type has a different way of evaluating the "value"
	// of the "value map" based on the "values" of each individual "value map item".
	GetValueMapType(context.Context, *GetValueMapTypeRequest) (*GetValueMapTypeResponse, error)

	// The ValueMapType enumeration defines a value map type. A value map defines a named-value
	// pair. A value map can be used in different ways. A value map type defines which way
	// the value map is to be used; each type has a different way of evaluating the "value"
	// of the "value map" based on the "values" of each individual "value map item".
	SetValueMapType(context.Context, *SetValueMapTypeRequest) (*SetValueMapTypeResponse, error)
}

func RegisterValueMapItemServer(conn dcerpc.Conn, o ValueMapItemServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewValueMapItemServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ValueMapItemSyntaxV0_0))...)
}

func NewValueMapItemServerHandle(o ValueMapItemServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ValueMapItemServerHandle(ctx, o, opNum, r)
	}
}

func ValueMapItemServerHandle(ctx context.Context, o ValueMapItemServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Description
		in := &GetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Description
		in := &SetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Enabled
		in := &GetEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Enabled
		in := &SetEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Key
		in := &GetKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Key
		in := &SetKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Value
		in := &GetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // Value
		in := &SetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // ValueMapType
		in := &GetValueMapTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValueMapType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // ValueMapType
		in := &SetValueMapTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetValueMapType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
