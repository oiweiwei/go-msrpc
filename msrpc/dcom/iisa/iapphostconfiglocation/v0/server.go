package iapphostconfiglocation

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

// IAppHostConfigLocation server interface.
type AppHostConfigLocationServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The Path method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the subpath for the specific IAppHostConfigLocation.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *pbstrLocationPath is not NULL. If
	// processing fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF].
	// The following table describes the error conditions that MUST be handled and the corresponding
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
	GetPath(context.Context, *GetPathRequest) (*GetPathResponse, error)

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// The AddConfigSection method is received by the server in an RPC_REQUEST packet. In
	// response, the server attempts to create a new empty IAppHostElement and add it to
	// the specified IAppHostConfigLocation. The server MAY choose to create the IAppHostElement
	// object in memory only and not persist it to permanent storage, such as a disk file,
	// until later.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppAdminElement is not NULL. If processing
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
	//	| 0X800700B7 ERROR_ALREADY_EXISTS    | A configuration element with the name specified by bstrSectionName already       |
	//	|                                    | exists.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted.           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070021 ERROR_LOCK_VIOLATION    | The instance is not editable.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	AddConfigSection(context.Context, *AddConfigSectionRequest) (*AddConfigSectionResponse, error)

	// The DeleteConfigSection method is received by the server in an RPC_REQUEST packet.
	// In response, the server attempts to delete the IAppHostElement of the specified index.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
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
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted.           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070021 ERROR_LOCK_VIOLATION    | The instance is not editable.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070585 ERROR_INVALID_INDEX     | The integer index specified by cIndex is invalid, or the element with name       |
	//	|                                    | specified by cIndex could not be found.                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	DeleteConfigSection(context.Context, *DeleteConfigSectionRequest) (*DeleteConfigSectionResponse, error)
}

func RegisterAppHostConfigLocationServer(conn dcerpc.Conn, o AppHostConfigLocationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostConfigLocationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostConfigLocationSyntaxV0_0))...)
}

func NewAppHostConfigLocationServerHandle(o AppHostConfigLocationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostConfigLocationServerHandle(ctx, o, opNum, r)
	}
}

func AppHostConfigLocationServerHandle(ctx context.Context, o AppHostConfigLocationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Path
		op := &xxx_GetPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Count
		op := &xxx_GetCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Item
		op := &xxx_GetItemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetItemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetItem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // AddConfigSection
		op := &xxx_AddConfigSectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddConfigSectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddConfigSection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // DeleteConfigSection
		op := &xxx_DeleteConfigSectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteConfigSectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteConfigSection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostConfigLocation
type UnimplementedAppHostConfigLocationServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostConfigLocationServer) GetPath(context.Context, *GetPathRequest) (*GetPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationServer) AddConfigSection(context.Context, *AddConfigSectionRequest) (*AddConfigSectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigLocationServer) DeleteConfigSection(context.Context, *DeleteConfigSectionRequest) (*DeleteConfigSectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostConfigLocationServer = (*UnimplementedAppHostConfigLocationServer)(nil)
