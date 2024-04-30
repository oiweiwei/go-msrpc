package ifsrmpropertydefinitionvalue

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

// IFsrmPropertyDefinitionValue server interface.
type PropertyDefinitionValueServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// DisplayName operation.
	GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error)

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	// The UniqueID (get) method returns the Property Value Definition.UniqueId of the property
	// definition value.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+---------------------------------+
	//	|         RETURN          |                                 |
	//	|       VALUE/CODE        |           DESCRIPTION           |
	//	|                         |                                 |
	//	+-------------------------+---------------------------------+
	//	+-------------------------+---------------------------------+
	//	| 0x80070057 E_INVALIDARG | The uniqueID parameter is NULL. |
	//	+-------------------------+---------------------------------+
	GetUniqueID(context.Context, *GetUniqueIDRequest) (*GetUniqueIDResponse, error)
}

func RegisterPropertyDefinitionValueServer(conn dcerpc.Conn, o PropertyDefinitionValueServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPropertyDefinitionValueServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PropertyDefinitionValueSyntaxV0_0))...)
}

func NewPropertyDefinitionValueServerHandle(o PropertyDefinitionValueServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PropertyDefinitionValueServerHandle(ctx, o, opNum, r)
	}
}

func PropertyDefinitionValueServerHandle(ctx context.Context, o PropertyDefinitionValueServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 8: // DisplayName
		in := &GetDisplayNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDisplayName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Description
		in := &GetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // UniqueID
		in := &GetUniqueIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUniqueID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
