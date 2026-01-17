package iapphostelement

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

// IAppHostElement server interface.
type AppHostElementServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// The Collection method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns an IAppHostElementCollection that represents a collection of "collection
	// IAppHostElement" objects. If the specific IAppHostElement does not support this type
	// of child object, it indicates this in the return.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
	//
	//	+------------------------------------+------------------------------------------------------------------------+
	//	|               RETURN               |                                                                        |
	//	|             VALUE/CODE             |                              DESCRIPTION                               |
	//	|                                    |                                                                        |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                                  |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                          |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted. |
	//	+------------------------------------+------------------------------------------------------------------------+
	GetCollection(context.Context, *GetCollectionRequest) (*GetCollectionResponse, error)

	// The Properties method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns an IAppHostPropertyCollection that contains the IAppHostProperty
	// objects that are available for this IAppHostElement.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppProperties is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+------------------------------------------------------------------------+
	//	|               RETURN               |                                                                        |
	//	|             VALUE/CODE             |                              DESCRIPTION                               |
	//	|                                    |                                                                        |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                                  |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                          |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted. |
	//	+------------------------------------+------------------------------------------------------------------------+
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The ChildElements method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns an IAppHostChildElementCollection that contains child IAppHostElement
	// objects if any child IAppHostElement objects exist.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
	//
	//	+------------------------------------+------------------------------------------------------------------------+
	//	|               RETURN               |                                                                        |
	//	|             VALUE/CODE             |                              DESCRIPTION                               |
	//	|                                    |                                                                        |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                                  |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                          |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted. |
	//	+------------------------------------+------------------------------------------------------------------------+
	GetChildElements(context.Context, *GetChildElementsRequest) (*GetChildElementsResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// SetMetadata operation.
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)

	// Schema operation.
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)

	// The GetElementByName method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns the child IAppHostElement object with the specified
	// name.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppElement is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+----------------------------------------------------------------+
	//	|               RETURN               |                                                                |
	//	|             VALUE/CODE             |                          DESCRIPTION                           |
	//	|                                    |                                                                |
	//	+------------------------------------+----------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                          |
	//	+------------------------------------+----------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                  |
	//	+------------------------------------+----------------------------------------------------------------+
	//	| 0X80070585 ERROR_INVALID_INDEX     | The child element specified by bstrSubName could not be found. |
	//	+------------------------------------+----------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.        |
	//	+------------------------------------+----------------------------------------------------------------+
	GetElementByName(context.Context, *GetElementByNameRequest) (*GetElementByNameResponse, error)

	// The GetPropertyByName method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns the IAppHostProperty of the specified name.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppProperty is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+-----------------------------------------------------------+
	//	|               RETURN               |                                                           |
	//	|             VALUE/CODE             |                        DESCRIPTION                        |
	//	|                                    |                                                           |
	//	+------------------------------------+-----------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                     |
	//	+------------------------------------+-----------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.             |
	//	+------------------------------------+-----------------------------------------------------------+
	//	| 0X80070585 ERROR_INVALID_INDEX     | The property specified by bstrSubName could not be found. |
	//	+------------------------------------+-----------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.   |
	//	+------------------------------------+-----------------------------------------------------------+
	GetPropertyByName(context.Context, *GetPropertyByNameRequest) (*GetPropertyByNameResponse, error)

	// Clear operation.
	Clear(context.Context, *ClearRequest) (*ClearResponse, error)

	// The Methods method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns an IAppHostMethodCollection, which is the collection of methods
	// that are supported for the specific IAppHostElement object.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppMethods is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+---------------------------------------------------------+
	//	|               RETURN               |                                                         |
	//	|             VALUE/CODE             |                       DESCRIPTION                       |
	//	|                                    |                                                         |
	//	+------------------------------------+---------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                   |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.           |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command. |
	//	+------------------------------------+---------------------------------------------------------+
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
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Collection
		op := &xxx_GetCollectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCollectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCollection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // ChildElements
		op := &xxx_GetChildElementsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetChildElementsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetChildElements(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // GetMetadata
		op := &xxx_GetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SetMetadata
		op := &xxx_SetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Schema
		op := &xxx_GetSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetElementByName
		op := &xxx_GetElementByNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetElementByNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetElementByName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // GetPropertyByName
		op := &xxx_GetPropertyByNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertyByNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPropertyByName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Clear
		op := &xxx_ClearOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Clear(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Methods
		op := &xxx_GetMethodsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMethodsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMethods(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostElement
type UnimplementedAppHostElementServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostElementServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementServer) GetCollection(context.Context, *GetCollectionRequest) (*GetCollectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementServer) GetChildElements(context.Context, *GetChildElementsRequest) (*GetChildElementsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementServer) SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementServer) GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementServer) GetElementByName(context.Context, *GetElementByNameRequest) (*GetElementByNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementServer) GetPropertyByName(context.Context, *GetPropertyByNameRequest) (*GetPropertyByNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementServer) Clear(context.Context, *ClearRequest) (*ClearResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementServer) GetMethods(context.Context, *GetMethodsRequest) (*GetMethodsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostElementServer = (*UnimplementedAppHostElementServer)(nil)
