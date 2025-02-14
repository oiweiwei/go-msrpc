package iapphostmethodinstance

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

// IAppHostMethodInstance server interface.
type AppHostMethodInstanceServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Input operation.
	GetInput(context.Context, *GetInputRequest) (*GetInputResponse, error)

	// Output operation.
	GetOutput(context.Context, *GetOutputRequest) (*GetOutputResponse, error)

	// Execute operation.
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// SetMetadata operation.
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)
}

func RegisterAppHostMethodInstanceServer(conn dcerpc.Conn, o AppHostMethodInstanceServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostMethodInstanceServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostMethodInstanceSyntaxV0_0))...)
}

func NewAppHostMethodInstanceServerHandle(o AppHostMethodInstanceServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostMethodInstanceServerHandle(ctx, o, opNum, r)
	}
}

func AppHostMethodInstanceServerHandle(ctx context.Context, o AppHostMethodInstanceServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Input
		op := &xxx_GetInputOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInputRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInput(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Output
		op := &xxx_GetOutputOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOutputRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOutput(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Execute
		op := &xxx_ExecuteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecuteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Execute(ctx, req)
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
	case 7: // SetMetadata
		op := &xxx_SetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostMethodInstance
type UnimplementedAppHostMethodInstanceServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostMethodInstanceServer) GetInput(context.Context, *GetInputRequest) (*GetInputResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMethodInstanceServer) GetOutput(context.Context, *GetOutputRequest) (*GetOutputResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMethodInstanceServer) Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMethodInstanceServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMethodInstanceServer) SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostMethodInstanceServer = (*UnimplementedAppHostMethodInstanceServer)(nil)
