package iapphostelementschema

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

// IAppHostElementSchema server interface.
type AppHostElementSchemaServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// The DoesAllowUnschematizedProperties method is received by the server in an RPC_REQUEST
	// packet. In response, the server returns whether the corresponding IAppHostElement
	// supports IAppHostProperty objects.
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
	GetDoesAllowUnschematizedProperties(context.Context, *GetDoesAllowUnschematizedPropertiesRequest) (*GetDoesAllowUnschematizedPropertiesResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// The CollectionSchema method is received by the server in an RPC_REQUEST packet. If
	// the specified IAppHostElement object supports child collection elements, the server
	// returns the schema and constraints of the collection that is contained in the corresponding
	// IAppHostElement object.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppCollectionSchema is not NULL.
	// If processing fails, the server MUST return a nonzero HRESULT code as defined in
	// [MS-ERREF]. The following table describes the error conditions that MUST be handled
	// and the corresponding error codes. A server MAY return additional implementation-specific
	// error codes.
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
	GetCollectionSchema(context.Context, *GetCollectionSchemaRequest) (*GetCollectionSchemaResponse, error)

	// The ChildElementSchemas method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns the schema and constraints of any child elements
	// that are contained in the corresponding IAppHostElement object.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppChildSchemas is not NULL. If processing
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
	GetChildElementSchemas(context.Context, *GetChildElementSchemasRequest) (*GetChildElementSchemasResponse, error)

	// The PropertySchemas method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns the schema and constraints for the IAppHostProperty
	// objects that are contained in the corresponding IAppHostElement.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppPropertySchemas is not NULL. If
	// processing fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF].
	// The following table describes the error conditions that MUST be handled and the corresponding
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
	GetPropertySchemas(context.Context, *GetPropertySchemasRequest) (*GetPropertySchemasResponse, error)

	// The IsCollectionDefault method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns whether the corresponding IAppHostElement object
	// is also considered to be a supported default for other IAppHostElement objects in
	// the administration system.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
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
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // DoesAllowUnschematizedProperties
		op := &xxx_GetDoesAllowUnschematizedPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDoesAllowUnschematizedPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDoesAllowUnschematizedProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetMetadata
		op := &xxx_GetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // CollectionSchema
		op := &xxx_GetCollectionSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCollectionSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCollectionSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // ChildElementSchemas
		op := &xxx_GetChildElementSchemasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetChildElementSchemasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetChildElementSchemas(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // PropertySchemas
		op := &xxx_GetPropertySchemasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertySchemasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPropertySchemas(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // IsCollectionDefault
		op := &xxx_GetIsCollectionDefaultOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsCollectionDefaultRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsCollectionDefault(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostElementSchema
type UnimplementedAppHostElementSchemaServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostElementSchemaServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementSchemaServer) GetDoesAllowUnschematizedProperties(context.Context, *GetDoesAllowUnschematizedPropertiesRequest) (*GetDoesAllowUnschematizedPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementSchemaServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementSchemaServer) GetCollectionSchema(context.Context, *GetCollectionSchemaRequest) (*GetCollectionSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementSchemaServer) GetChildElementSchemas(context.Context, *GetChildElementSchemasRequest) (*GetChildElementSchemasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementSchemaServer) GetPropertySchemas(context.Context, *GetPropertySchemasRequest) (*GetPropertySchemasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementSchemaServer) GetIsCollectionDefault(context.Context, *GetIsCollectionDefaultRequest) (*GetIsCollectionDefaultResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostElementSchemaServer = (*UnimplementedAppHostElementSchemaServer)(nil)
