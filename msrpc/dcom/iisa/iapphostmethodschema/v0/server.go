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

	// The InputSchema method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the schema and constraints of the input parameters to the method
	// call. This can be NULL if no input parameters are defined for the method.
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
	GetInputSchema(context.Context, *GetInputSchemaRequest) (*GetInputSchemaResponse, error)

	// The OutputSchema method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the schema and constraints of the output parameters to the method
	// call. This can be NULL if no output parameters are defined for the method.
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
