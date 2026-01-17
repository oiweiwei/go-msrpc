package iapphostmethodinstance

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = iunknown.GoPackage
)

// IAppHostMethodInstance server interface.
type AppHostMethodInstanceServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The Input method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns an IAppHostElement in which input parameters can be specified
	// for the specified method instance call. If the method does not support input parameters,
	// no IAppHostElement is returned.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
	//
	//	+------------------------------------+-----------------------------------------------+
	//	|               RETURN               |                                               |
	//	|             VALUE/CODE             |                  DESCRIPTION                  |
	//	|                                    |                                               |
	//	+------------------------------------+-----------------------------------------------+
	//	+------------------------------------+-----------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.         |
	//	+------------------------------------+-----------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null. |
	//	+------------------------------------+-----------------------------------------------+
	GetInput(context.Context, *GetInputRequest) (*GetInputResponse, error)

	// The Output method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns an IAppHostElement from which output parameters are retrieved
	// after the specified method instance has been executed. If the method does not support
	// output parameters, no IAppHostElement is returned.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
	//
	//	+------------------------------------+-----------------------------------------------+
	//	|               RETURN               |                                               |
	//	|             VALUE/CODE             |                  DESCRIPTION                  |
	//	|                                    |                                               |
	//	+------------------------------------+-----------------------------------------------+
	//	+------------------------------------+-----------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.         |
	//	+------------------------------------+-----------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null. |
	//	+------------------------------------+-----------------------------------------------+
	GetOutput(context.Context, *GetOutputRequest) (*GetOutputResponse, error)

	// The Execute method is received by the server in an RPC_REQUEST packet. In response,
	// the server actually executes the specified custom method.
	//
	// This method has no parameters.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
	//
	//	+---------------------+---------------------------------------+
	//	|       RETURN        |                                       |
	//	|     VALUE/CODE      |              DESCRIPTION              |
	//	|                     |                                       |
	//	+---------------------+---------------------------------------+
	//	+---------------------+---------------------------------------+
	//	| 0X00000000 NO_ERROR | The operation completed successfully. |
	//	+---------------------+---------------------------------------+
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
