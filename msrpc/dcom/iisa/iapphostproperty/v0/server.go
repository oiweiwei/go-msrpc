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

	// The StringValue method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns a string representation of the value of the specified property.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *pbstrValue is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+--------------------------------------------------------------+
	//	|               RETURN               |                                                              |
	//	|             VALUE/CODE             |                         DESCRIPTION                          |
	//	|                                    |                                                              |
	//	+------------------------------------+--------------------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                        |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0X80070013 ERROR_INVALID_DATA      | The property has a type that is not permitted by the schema. |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.      |
	//	+------------------------------------+--------------------------------------------------------------+
	GetStringValue(context.Context, *GetStringValueRequest) (*GetStringValueResponse, error)

	// The Exception method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns administration system exception information that is related to
	// the processing of the specified property.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppException is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
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
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Value
		op := &xxx_GetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Value
		op := &xxx_SetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Clear
		op := &xxx_ClearOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Clear(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // StringValue
		op := &xxx_GetStringValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStringValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetStringValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Exception
		op := &xxx_GetExceptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExceptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetException(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // GetMetadata
		op := &xxx_GetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // SetMetadata
		op := &xxx_SetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Schema
		op := &xxx_GetSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostProperty
type UnimplementedAppHostPropertyServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostPropertyServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertyServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertyServer) SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertyServer) Clear(context.Context, *ClearRequest) (*ClearResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertyServer) GetStringValue(context.Context, *GetStringValueRequest) (*GetStringValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertyServer) GetException(context.Context, *GetExceptionRequest) (*GetExceptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertyServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertyServer) SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertyServer) GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostPropertyServer = (*UnimplementedAppHostPropertyServer)(nil)
