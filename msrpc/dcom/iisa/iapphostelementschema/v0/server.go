package iapphostelementschema

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

// IAppHostElementSchema server interface.
type AppHostElementSchemaServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// DoesAllowUnschematizedProperties operation.
	GetDoesAllowUnschematizedProperties(context.Context, *GetDoesAllowUnschematizedPropertiesRequest) (*GetDoesAllowUnschematizedPropertiesResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// CollectionSchema operation.
	GetCollectionSchema(context.Context, *GetCollectionSchemaRequest) (*GetCollectionSchemaResponse, error)

	// ChildElementSchemas operation.
	GetChildElementSchemas(context.Context, *GetChildElementSchemasRequest) (*GetChildElementSchemasResponse, error)

	// PropertySchemas operation.
	GetPropertySchemas(context.Context, *GetPropertySchemasRequest) (*GetPropertySchemasResponse, error)

	// IsCollectionDefault operation.
	GetIsCollectionDefault(context.Context, *GetIsCollectionDefaultRequest) (*GetIsCollectionDefaultResponse, error)
}

func RegisterAppHostElementSchemaServer(conn dcerpc.Conn, o AppHostElementSchemaServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostElementSchemaServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostElementSchemaSyntaxV0_0))...)
}

func NewAppHostElementSchemaServerHandle(o AppHostElementSchemaServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostElementSchemaServerHandle(ctx, o, opNum, r)
	}
}

func AppHostElementSchemaServerHandle(ctx context.Context, o AppHostElementSchemaServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 4: // DoesAllowUnschematizedProperties
		in := &GetDoesAllowUnschematizedPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDoesAllowUnschematizedProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GetMetadata
		in := &GetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // CollectionSchema
		in := &GetCollectionSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCollectionSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // ChildElementSchemas
		in := &GetChildElementSchemasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetChildElementSchemas(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // PropertySchemas
		in := &GetPropertySchemasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertySchemas(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // IsCollectionDefault
		in := &GetIsCollectionDefaultRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIsCollectionDefault(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
