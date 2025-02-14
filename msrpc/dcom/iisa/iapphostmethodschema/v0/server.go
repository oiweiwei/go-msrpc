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
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // InputSchema
		op := &xxx_GetInputSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInputSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInputSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // OutputSchema
		op := &xxx_GetOutputSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOutputSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOutputSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // GetMetadata
		op := &xxx_GetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostMethodSchema
type UnimplementedAppHostMethodSchemaServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostMethodSchemaServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMethodSchemaServer) GetInputSchema(context.Context, *GetInputSchemaRequest) (*GetInputSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMethodSchemaServer) GetOutputSchema(context.Context, *GetOutputSchemaRequest) (*GetOutputSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMethodSchemaServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostMethodSchemaServer = (*UnimplementedAppHostMethodSchemaServer)(nil)
