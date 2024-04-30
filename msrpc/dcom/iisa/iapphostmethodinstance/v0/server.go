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
		in := &GetInputRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetInput(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Output
		in := &GetOutputRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOutput(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Execute
		in := &ExecuteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Execute(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // GetMetadata
		in := &GetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // SetMetadata
		in := &SetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
