package iapphostelementcollection

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

// IAppHostElementCollection server interface.
type AppHostElementCollectionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// The AddElement method is received by the server in an RPC_REQUEST packet. In response,
	// the server attempts to add a newly created IAppHostElement object to the specified
	// collection.
	//
	// This function adds an element to the current collection of elements.
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
	//	| 0X80070021 ERROR_LOCK_VIOLATION    | The instance is not editable.                           |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070585 ERROR_INVALID_INDEX     | The index specified by cPosition is invalid.            |
	//	+------------------------------------+---------------------------------------------------------+
	AddElement(context.Context, *AddElementRequest) (*AddElementResponse, error)

	// The DeleteElement method is received by the server in an RPC_REQUEST packet. In response,
	// the server deletes the IAppHostElement at the specified index.
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
	//	| 0X80070021 ERROR_LOCK_VIOLATION    | The instance is not editable.                           |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070585 ERROR_INVALID_INDEX     | The index specified by cIndex is invalid.               |
	//	+------------------------------------+---------------------------------------------------------+
	DeleteElement(context.Context, *DeleteElementRequest) (*DeleteElementResponse, error)

	// Clear operation.
	Clear(context.Context, *ClearRequest) (*ClearResponse, error)

	// The CreateNewElement method is received by the server in an RPC_REQUEST packet. In
	// response, the server creates a new unattached IAppHostElement object that has the
	// specified name. The name MUST be a supported name as defined by the IAppHostCollectionSchema
	// of the specified IAppHostElementCollection.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppElement is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070021 ERROR_LOCK_VIOLATION    | The instance is not editable.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070585 ERROR_INVALID_INDEX     | The schema does not permit the creation of an element with name bstrElementName. |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	CreateNewElement(context.Context, *CreateNewElementRequest) (*CreateNewElementResponse, error)

	// Schema operation.
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)
}

func RegisterAppHostElementCollectionServer(conn dcerpc.Conn, o AppHostElementCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostElementCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostElementCollectionSyntaxV0_0))...)
}

func NewAppHostElementCollectionServerHandle(o AppHostElementCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostElementCollectionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostElementCollectionServerHandle(ctx context.Context, o AppHostElementCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Count
		op := &xxx_GetCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Item
		op := &xxx_GetItemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetItemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetItem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // AddElement
		op := &xxx_AddElementOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddElementRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddElement(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // DeleteElement
		op := &xxx_DeleteElementOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteElementRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteElement(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // Clear
		op := &xxx_ClearOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Clear(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CreateNewElement
		op := &xxx_CreateNewElementOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNewElementRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNewElement(ctx, req)
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
	}
	return nil, nil
}

// Unimplemented IAppHostElementCollection
type UnimplementedAppHostElementCollectionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostElementCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementCollectionServer) AddElement(context.Context, *AddElementRequest) (*AddElementResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementCollectionServer) DeleteElement(context.Context, *DeleteElementRequest) (*DeleteElementResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementCollectionServer) Clear(context.Context, *ClearRequest) (*ClearResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementCollectionServer) CreateNewElement(context.Context, *CreateNewElementRequest) (*CreateNewElementResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostElementCollectionServer) GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostElementCollectionServer = (*UnimplementedAppHostElementCollectionServer)(nil)
