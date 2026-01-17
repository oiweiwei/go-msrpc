package ifsrmpropertydefinitionvalue

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
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // DisplayName
		op := &xxx_GetDisplayNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDisplayNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDisplayName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Description
		op := &xxx_GetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // UniqueID
		op := &xxx_GetUniqueIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUniqueIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUniqueID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmPropertyDefinitionValue
type UnimplementedPropertyDefinitionValueServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedPropertyDefinitionValueServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionValueServer) GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionValueServer) GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPropertyDefinitionValueServer) GetUniqueID(context.Context, *GetUniqueIDRequest) (*GetUniqueIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PropertyDefinitionValueServer = (*UnimplementedPropertyDefinitionValueServer)(nil)
