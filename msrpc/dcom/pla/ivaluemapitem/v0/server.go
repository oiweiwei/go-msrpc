package ivaluemapitem

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_GetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Description
		op := &xxx_SetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Enabled
		op := &xxx_GetEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Enabled
		op := &xxx_SetEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Key
		op := &xxx_GetKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Key
		op := &xxx_SetKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Value
		op := &xxx_GetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Value
		op := &xxx_SetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // ValueMapType
		op := &xxx_GetValueMapTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValueMapTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValueMapType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // ValueMapType
		op := &xxx_SetValueMapTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetValueMapTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetValueMapType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IValueMapItem
type UnimplementedValueMapItemServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedValueMapItemServer) GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedValueMapItemServer) SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedValueMapItemServer) GetEnabled(context.Context, *GetEnabledRequest) (*GetEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedValueMapItemServer) SetEnabled(context.Context, *SetEnabledRequest) (*SetEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedValueMapItemServer) GetKey(context.Context, *GetKeyRequest) (*GetKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedValueMapItemServer) SetKey(context.Context, *SetKeyRequest) (*SetKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedValueMapItemServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedValueMapItemServer) SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedValueMapItemServer) GetValueMapType(context.Context, *GetValueMapTypeRequest) (*GetValueMapTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedValueMapItemServer) SetValueMapType(context.Context, *SetValueMapTypeRequest) (*SetValueMapTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ValueMapItemServer = (*UnimplementedValueMapItemServer)(nil)
