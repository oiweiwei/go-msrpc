package iapphostconfiglocationcollection

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

// IAppHostConfigLocationCollection server interface.
type AppHostConfigLocationCollectionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// The AddLocation method is received by the server in an RPC_REQUEST packet. In response,
	// the server attempts to create a new subpath container in the IAppHostConfigFile that
	// provides the specified IAppHostConfigLocationCollection.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppNewLocation is not NULL. If processing
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
	//	| 0X80070005 ERROR_ACCESS_DENIED     | The instance is not editable.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X800700B7 ERROR_ALREADY_EXISTS    | The location path specified by bstrLocationPath cannot be added since it already |
	//	|                                    | exists.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X800700A1 ERROR_BAD_PATHNAME      | The server resource (for example, a file or database) corresponding to the path  |
	//	|                                    | bstrLocationPath could not be found.                                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	AddLocation(context.Context, *AddLocationRequest) (*AddLocationResponse, error)

	// The DeleteLocation method is received by the server in an RPC_REQUEST packet. In
	// response, the server attempts to delete the specific subpath container (the IAppHostConfigLocation
	// object).
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR            | The operation completed successfully.                                            |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070005 ERROR_ACCESS_DENIED | The instance is not editable.                                                    |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070585 ERROR_INVALID_INDEX | The integer index specified by cIndex is invalid, or the location with name      |
	//	|                                | specified by cIndex could not be found.                                          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error)
}

func RegisterAppHostConfigLocationCollectionServer(conn dcerpc.Conn, o AppHostConfigLocationCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostConfigLocationCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostConfigLocationCollectionSyntaxV0_0))...)
}

func NewAppHostConfigLocationCollectionServerHandle(o AppHostConfigLocationCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostConfigLocationCollectionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostConfigLocationCollectionServerHandle(ctx context.Context, o AppHostConfigLocationCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 5: // AddLocation
		op := &xxx_AddLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // DeleteLocation
		op := &xxx_DeleteLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostConfigLocationCollection
type UnimplementedAppHostConfigLocationCollectionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostConfigLocationCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationCollectionServer) AddLocation(context.Context, *AddLocationRequest) (*AddLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationCollectionServer) DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostConfigLocationCollectionServer = (*UnimplementedAppHostConfigLocationCollectionServer)(nil)
