package iapphostcollectionschema

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

// IAppHostCollectionSchema server interface.
type AppHostCollectionSchemaServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The AddElementNames method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns a comma-delimited list of names that are supported by
	// the administration system as names to IAppHostElement objects that are collection
	// items of an IAppHostElementCollection object. An administration system typically
	// supports only one name for the IAppHostElement objects that are contained in the
	// collection. However, it could support more names in certain conditions; in which
	// case, all the names are returned by using this method.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *pbstrElementName MUST NOT be NULL.
	// If processing fails, the server MUST return a nonzero HRESULT code as defined in
	// [MS-ERREF]. The following table describes the error conditions that MUST be handled
	// and the corresponding error codes. A server MAY return additional implementation-specific
	// error codes.
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
	GetAddElementNames(context.Context, *GetAddElementNamesRequest) (*GetAddElementNamesResponse, error)

	// The GetAddElementSchema method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns an IAppHostElementSchema that represents the schema
	// and constraints of the IAppHostElement. The IAppHostElement can be a collection item
	// of the specified IAppHostElementCollection from which the specified IAppHostCollectionSchema
	// was retrieved and whose name matches the specified name in the method call.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppSchema MUST NOT be NULL. If processing
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
	//	| 0X80070585 ERROR_INVALID_INDEX     | The element specified by bstrElementName cannot be found. |
	//	+------------------------------------+-----------------------------------------------------------+
	GetAddElementSchema(context.Context, *GetAddElementSchemaRequest) (*GetAddElementSchemaResponse, error)

	// The RemoveElementSchema method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns the IAppHostElementSchema that represents the schema
	// and constraints of an optionally supported "directive IAppHostElement". This directive
	// element can be used by the administration system to control the behavior of the specific
	// IAppHostElementCollection from which the specified IAppHostCollectionSchema was retrieved.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF].
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
	GetRemoveElementSchema(context.Context, *GetRemoveElementSchemaRequest) (*GetRemoveElementSchemaResponse, error)

	// The ClearElementSchema method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns the IAppHostElementSchema that represents the schema
	// and constraints of an optionally supported "directive IAppHostElement". This directive
	// element can be used by the administration system to control behavior of the specific
	// IAppHostElementCollection from which the specified IAppHostCollectionSchema was retrieved.
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
	GetClearElementSchema(context.Context, *GetClearElementSchemaRequest) (*GetClearElementSchemaResponse, error)

	// The IsMergeAppend method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns a Boolean that represents whether the IAppHostElementCollection,
	// from which the specified IAppHostCollectionSchema was retrieved, will prepend collection
	// IAppHostElement objects from lower (deeper) in the hierarchy of the administration
	// system with IAppHostElement objects from higher (shallower) in the hierarchy of the
	// administration system.
	//
	// If the value is false, lower IAppHostElement objects are prepended; otherwise, they
	// are appended.
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
	GetIsMergeAppend(context.Context, *GetIsMergeAppendRequest) (*GetIsMergeAppendResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// The DoesAllowDuplicates method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns whether (through a Boolean) the IAppHostElementCollection
	// from which the specified IAppHostCollectionSchema was retrieved allows duplicate
	// IAppHostElement objects to exist in the collection. The concept of "duplicate" is
	// implementation-specific.
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
	GetDoesAllowDuplicates(context.Context, *GetDoesAllowDuplicatesRequest) (*GetDoesAllowDuplicatesResponse, error)
}

func RegisterAppHostCollectionSchemaServer(conn dcerpc.Conn, o AppHostCollectionSchemaServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostCollectionSchemaServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostCollectionSchemaSyntaxV0_0))...)
}

func NewAppHostCollectionSchemaServerHandle(o AppHostCollectionSchemaServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostCollectionSchemaServerHandle(ctx, o, opNum, r)
	}
}

func AppHostCollectionSchemaServerHandle(ctx context.Context, o AppHostCollectionSchemaServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // AddElementNames
		op := &xxx_GetAddElementNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAddElementNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAddElementNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetAddElementSchema
		op := &xxx_GetAddElementSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAddElementSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAddElementSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // RemoveElementSchema
		op := &xxx_GetRemoveElementSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRemoveElementSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRemoveElementSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // ClearElementSchema
		op := &xxx_GetClearElementSchemaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClearElementSchemaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClearElementSchema(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // IsMergeAppend
		op := &xxx_GetIsMergeAppendOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsMergeAppendRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsMergeAppend(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetMetadata
		op := &xxx_GetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // DoesAllowDuplicates
		op := &xxx_GetDoesAllowDuplicatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDoesAllowDuplicatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDoesAllowDuplicates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostCollectionSchema
type UnimplementedAppHostCollectionSchemaServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostCollectionSchemaServer) GetAddElementNames(context.Context, *GetAddElementNamesRequest) (*GetAddElementNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetAddElementSchema(context.Context, *GetAddElementSchemaRequest) (*GetAddElementSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetRemoveElementSchema(context.Context, *GetRemoveElementSchemaRequest) (*GetRemoveElementSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetClearElementSchema(context.Context, *GetClearElementSchemaRequest) (*GetClearElementSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetIsMergeAppend(context.Context, *GetIsMergeAppendRequest) (*GetIsMergeAppendResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostCollectionSchemaServer) GetDoesAllowDuplicates(context.Context, *GetDoesAllowDuplicatesRequest) (*GetDoesAllowDuplicatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostCollectionSchemaServer = (*UnimplementedAppHostCollectionSchemaServer)(nil)
