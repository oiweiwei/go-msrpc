package ifsrmpropertydefinition

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

// IFsrmPropertyDefinition server interface.
type PropertyDefinitionServer interface {

	// IFsrmObject base class.
	ifsrmobject.ObjectServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest) (*SetNameResponse, error)

	// Type operation.
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

	// Type operation.
	SetType(context.Context, *SetTypeRequest) (*SetTypeResponse, error)

	// PossibleValues operation.
	GetPossibleValues(context.Context, *GetPossibleValuesRequest) (*GetPossibleValuesResponse, error)

	// PossibleValues operation.
	SetPossibleValues(context.Context, *SetPossibleValuesRequest) (*SetPossibleValuesResponse, error)

	// ValueDescriptions operation.
	GetValueDescriptions(context.Context, *GetValueDescriptionsRequest) (*GetValueDescriptionsResponse, error)

	// ValueDescriptions operation.
	SetValueDescriptions(context.Context, *SetValueDescriptionsRequest) (*SetValueDescriptionsResponse, error)

	// Parameters operation.
	GetParameters(context.Context, *GetParametersRequest) (*GetParametersResponse, error)

	// Parameters operation.
	SetParameters(context.Context, *SetParametersRequest) (*SetParametersResponse, error)
}

func RegisterPropertyDefinitionServer(conn dcerpc.Conn, o PropertyDefinitionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPropertyDefinitionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PropertyDefinitionSyntaxV0_0))...)
}

func NewPropertyDefinitionServerHandle(o PropertyDefinitionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PropertyDefinitionServerHandle(ctx, o, opNum, r)
	}
}

func PropertyDefinitionServerHandle(ctx context.Context, o PropertyDefinitionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 14: // Type
		op := &xxx_GetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Type
		op := &xxx_SetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // PossibleValues
		op := &xxx_GetPossibleValuesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPossibleValuesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPossibleValues(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // PossibleValues
		op := &xxx_SetPossibleValuesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPossibleValuesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPossibleValues(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // ValueDescriptions
		op := &xxx_GetValueDescriptionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValueDescriptionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValueDescriptions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // ValueDescriptions
		op := &xxx_SetValueDescriptionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetValueDescriptionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetValueDescriptions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Parameters
		op := &xxx_GetParametersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetParametersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetParameters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Parameters
		op := &xxx_SetParametersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetParametersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetParameters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmPropertyDefinition
type UnimplementedPropertyDefinitionServer struct {
	ifsrmobject.UnimplementedObjectServer
}

func (UnimplementedPropertyDefinitionServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionServer) SetName(context.Context, *SetNameRequest) (*SetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionServer) GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionServer) SetType(context.Context, *SetTypeRequest) (*SetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionServer) GetPossibleValues(context.Context, *GetPossibleValuesRequest) (*GetPossibleValuesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionServer) SetPossibleValues(context.Context, *SetPossibleValuesRequest) (*SetPossibleValuesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionServer) GetValueDescriptions(context.Context, *GetValueDescriptionsRequest) (*GetValueDescriptionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionServer) SetValueDescriptions(context.Context, *SetValueDescriptionsRequest) (*SetValueDescriptionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionServer) GetParameters(context.Context, *GetParametersRequest) (*GetParametersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionServer) SetParameters(context.Context, *SetParametersRequest) (*SetParametersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PropertyDefinitionServer = (*UnimplementedPropertyDefinitionServer)(nil)
