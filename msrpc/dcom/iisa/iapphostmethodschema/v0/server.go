package iapphostmethodschema

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

// IAppHostMethodSchema server interface.
type AppHostMethodSchemaServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// InputSchema operation.
	GetInputSchema(context.Context, *GetInputSchemaRequest) (*GetInputSchemaResponse, error)

	// OutputSchema operation.
	GetOutputSchema(context.Context, *GetOutputSchemaRequest) (*GetOutputSchemaResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)
}

func RegisterAppHostMethodSchemaServer(conn dcerpc.Conn, o AppHostMethodSchemaServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostMethodSchemaServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostMethodSchemaSyntaxV0_0))...)
}

func NewAppHostMethodSchemaServerHandle(o AppHostMethodSchemaServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostMethodSchemaServerHandle(ctx, o, opNum, r)
	}
}

func AppHostMethodSchemaServerHandle(ctx context.Context, o AppHostMethodSchemaServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 4: // InputSchema
		in := &GetInputSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetInputSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // OutputSchema
		in := &GetOutputSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOutputSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // GetMetadata
		in := &GetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
