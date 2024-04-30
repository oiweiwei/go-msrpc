package ifsrmpropertydefinition2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmpropertydefinition "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpropertydefinition/v0"
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
	_ = ifsrmpropertydefinition.GoPackage
)

// IFsrmPropertyDefinition2 server interface.
type PropertyDefinition2Server interface {

	// IFsrmPropertyDefinition base class.
	ifsrmpropertydefinition.PropertyDefinitionServer

	// The PropertyDefinitionFlags (get) method returns the Property Definition.Global,
	// Property Definition.Secure, and Property Definition.Deprecated values for the object.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+------------------------------------------------+
	//	|         RETURN          |                                                |
	//	|       VALUE/CODE        |                  DESCRIPTION                   |
	//	|                         |                                                |
	//	+-------------------------+------------------------------------------------+
	//	+-------------------------+------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The propertyDefinitionFlags parameter is NULL. |
	//	+-------------------------+------------------------------------------------+
	GetPropertyDefinitionFlags(context.Context, *GetPropertyDefinitionFlagsRequest) (*GetPropertyDefinitionFlagsResponse, error)

	// DisplayName operation.
	GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error)

	// DisplayName operation.
	SetDisplayName(context.Context, *SetDisplayNameRequest) (*SetDisplayNameResponse, error)

	// AppliesTo operation.
	GetAppliesTo(context.Context, *GetAppliesToRequest) (*GetAppliesToResponse, error)

	// The ValueDefinitions (get) method returns the property definitions Possible Values.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------------------+
	//	|         RETURN          |                                         |
	//	|       VALUE/CODE        |               DESCRIPTION               |
	//	|                         |                                         |
	//	+-------------------------+-----------------------------------------+
	//	+-------------------------+-----------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The valueDefinitions parameter is NULL. |
	//	+-------------------------+-----------------------------------------+
	GetValueDefinitions(context.Context, *GetValueDefinitionsRequest) (*GetValueDefinitionsResponse, error)
}

func RegisterPropertyDefinition2Server(conn dcerpc.Conn, o PropertyDefinition2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPropertyDefinition2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PropertyDefinition2SyntaxV0_0))...)
}

func NewPropertyDefinition2ServerHandle(o PropertyDefinition2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PropertyDefinition2ServerHandle(ctx, o, opNum, r)
	}
}

func PropertyDefinition2ServerHandle(ctx context.Context, o PropertyDefinition2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 22 {
		// IFsrmPropertyDefinition base method.
		return ifsrmpropertydefinition.PropertyDefinitionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 22: // PropertyDefinitionFlags
		in := &GetPropertyDefinitionFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertyDefinitionFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // DisplayName
		in := &GetDisplayNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDisplayName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // DisplayName
		in := &SetDisplayNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDisplayName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // AppliesTo
		in := &GetAppliesToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAppliesTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // ValueDefinitions
		in := &GetValueDefinitionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValueDefinitions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
