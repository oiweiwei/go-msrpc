package iapphostsectiondefinitioncollection

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

// IAppHostSectionDefinitionCollection server interface.
type AppHostSectionDefinitionCollectionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// The AddSection method is received by the server in an RPC_REQUEST packet. In response,
	// the server adds a section definition to the administration system.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppConfigSection is not NULL. If
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
	//	| 0X800700B7 ERROR_ALREADY_EXISTS    | A section with name bstrSectionName already exists.     |
	//	+------------------------------------+---------------------------------------------------------+
	AddSection(context.Context, *AddSectionRequest) (*AddSectionResponse, error)

	// The DeleteSection method is received by the server in an RPC_REQUEST packet. In response,
	// the server deletes the specified section definition.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR             | The operation completed successfully.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070585 ERROR_INVALID_INDEX  | The integer index specified by varIndex is invalid, or the section with name     |
	//	|                                 | specified by cIndex could not be found.                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070021 ERROR_LOCK_VIOLATION | The instance is set to read-only.                                                |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000002 ERROR_PATH_NOT_FOUND | The system cannot find the path specified.The section could not be found.        |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	DeleteSection(context.Context, *DeleteSectionRequest) (*DeleteSectionResponse, error)
}

func RegisterAppHostSectionDefinitionCollectionServer(conn dcerpc.Conn, o AppHostSectionDefinitionCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostSectionDefinitionCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostSectionDefinitionCollectionSyntaxV0_0))...)
}

func NewAppHostSectionDefinitionCollectionServerHandle(o AppHostSectionDefinitionCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostSectionDefinitionCollectionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostSectionDefinitionCollectionServerHandle(ctx context.Context, o AppHostSectionDefinitionCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 5: // AddSection
		op := &xxx_AddSectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddSectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddSection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // DeleteSection
		op := &xxx_DeleteSectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteSectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteSection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostSectionDefinitionCollection
type UnimplementedAppHostSectionDefinitionCollectionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostSectionDefinitionCollectionServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionCollectionServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionCollectionServer) AddSection(context.Context, *AddSectionRequest) (*AddSectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionDefinitionCollectionServer) DeleteSection(context.Context, *DeleteSectionRequest) (*DeleteSectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostSectionDefinitionCollectionServer = (*UnimplementedAppHostSectionDefinitionCollectionServer)(nil)
