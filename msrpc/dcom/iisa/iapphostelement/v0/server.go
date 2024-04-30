package iapphostelement

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IAppHostElement server interface.
type AppHostElementServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Collection operation.
	GetCollection(context.Context, *GetCollectionRequest) (*GetCollectionResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// ChildElements operation.
	GetChildElements(context.Context, *GetChildElementsRequest) (*GetChildElementsResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// SetMetadata operation.
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)

	// Schema operation.
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)

	// GetElementByName operation.
	GetElementByName(context.Context, *GetElementByNameRequest) (*GetElementByNameResponse, error)

	// GetPropertyByName operation.
	GetPropertyByName(context.Context, *GetPropertyByNameRequest) (*GetPropertyByNameResponse, error)

	// Clear operation.
	Clear(context.Context, *ClearRequest) (*ClearResponse, error)

	// Methods operation.
	GetMethods(context.Context, *GetMethodsRequest) (*GetMethodsResponse, error)
}

func RegisterAppHostElementServer(conn dcerpc.Conn, o AppHostElementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostElementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostElementSyntaxV0_0))...)
}

func NewAppHostElementServerHandle(o AppHostElementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostElementServerHandle(ctx, o, opNum, r)
	}
}

func AppHostElementServerHandle(ctx context.Context, o AppHostElementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Name
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Collection
		in := &GetCollectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCollection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Properties
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // ChildElements
		in := &GetChildElementsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetChildElements(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // GetMetadata
		in := &GetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // SetMetadata
		in := &SetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Schema
		in := &GetSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // GetElementByName
		in := &GetElementByNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetElementByName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // GetPropertyByName
		in := &GetPropertyByNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertyByName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Clear
		in := &ClearRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Clear(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Methods
		in := &GetMethodsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMethods(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
