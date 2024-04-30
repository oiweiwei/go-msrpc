package iapphostproperty

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

// IAppHostProperty server interface.
type AppHostPropertyServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Value operation.
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)

	// Value operation.
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)

	// Clear operation.
	Clear(context.Context, *ClearRequest) (*ClearResponse, error)

	// StringValue operation.
	GetStringValue(context.Context, *GetStringValueRequest) (*GetStringValueResponse, error)

	// Exception operation.
	GetException(context.Context, *GetExceptionRequest) (*GetExceptionResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// SetMetadata operation.
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)

	// Schema operation.
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)
}

func RegisterAppHostPropertyServer(conn dcerpc.Conn, o AppHostPropertyServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostPropertyServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostPropertySyntaxV0_0))...)
}

func NewAppHostPropertyServerHandle(o AppHostPropertyServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostPropertyServerHandle(ctx, o, opNum, r)
	}
}

func AppHostPropertyServerHandle(ctx context.Context, o AppHostPropertyServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 4: // Value
		in := &GetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Value
		in := &SetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // Clear
		in := &ClearRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Clear(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // StringValue
		in := &GetStringValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetStringValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Exception
		in := &GetExceptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetException(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // GetMetadata
		in := &GetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // SetMetadata
		in := &SetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Schema
		in := &GetSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
