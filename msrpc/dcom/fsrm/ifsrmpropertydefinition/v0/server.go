package ifsrmpropertydefinition

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
	case 14: // Type
		in := &GetTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Type
		in := &SetTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // PossibleValues
		in := &GetPossibleValuesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPossibleValues(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // PossibleValues
		in := &SetPossibleValuesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPossibleValues(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // ValueDescriptions
		in := &GetValueDescriptionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValueDescriptions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // ValueDescriptions
		in := &SetValueDescriptionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetValueDescriptions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // Parameters
		in := &GetParametersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetParameters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // Parameters
		in := &SetParametersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetParameters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
