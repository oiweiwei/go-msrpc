package iapphostmappingextension

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

// IAppHostMappingExtension server interface.
type AppHostMappingExtensionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetSiteNameFromSiteId method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns a string name for the specific integer site ID, which
	// is a concept that is implemented on the administration system (it is an implementation
	// detail).
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *pbstrSiteName is not NULL. If processing
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
	//	| 0X80070002 ERROR_FILE_NOT_FOUND    | The configuration has no site with ID dwSiteId.         |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070490 ERROR_NOT_FOUND         | The configuration contains no sites.                    |
	//	+------------------------------------+---------------------------------------------------------+
	GetSiteNameFromSiteID(context.Context, *GetSiteNameFromSiteIDRequest) (*GetSiteNameFromSiteIDResponse, error)

	// The GetSiteIdFromSiteName method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns a unique integer ID for the specific site name. Site
	// name and ID are implementation details of the administration system.
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
	//	| 0X80070002 ERROR_FILE_NOT_FOUND    | The configuration has no site with name bstrSiteName.   |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070490 ERROR_NOT_FOUND         | The configuration contains no sites.                    |
	//	+------------------------------------+---------------------------------------------------------+
	GetSiteIDFromSiteName(context.Context, *GetSiteIDFromSiteNameRequest) (*GetSiteIDFromSiteNameResponse, error)

	// The GetSiteElementFromSiteId method is received by the server in an RPC_REQUEST packet.
	// In response, the server obtains the site section element from a specific site ID
	// in order to access site configuration and properties.
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
	//	| 0X80070002 ERROR_FILE_NOT_FOUND    | The given site ID could not be found.                   |
	//	+------------------------------------+---------------------------------------------------------+
	GetSiteElementFromSiteID(context.Context, *GetSiteElementFromSiteIDRequest) (*GetSiteElementFromSiteIDResponse, error)

	// MapPath operation.
	MapPath(context.Context, *MapPathRequest) (*MapPathResponse, error)
}

func RegisterAppHostMappingExtensionServer(conn dcerpc.Conn, o AppHostMappingExtensionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostMappingExtensionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostMappingExtensionSyntaxV0_0))...)
}

func NewAppHostMappingExtensionServerHandle(o AppHostMappingExtensionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostMappingExtensionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostMappingExtensionServerHandle(ctx context.Context, o AppHostMappingExtensionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetSiteNameFromSiteId
		op := &xxx_GetSiteNameFromSiteIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSiteNameFromSiteIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSiteNameFromSiteID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetSiteIdFromSiteName
		op := &xxx_GetSiteIDFromSiteNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSiteIDFromSiteNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSiteIDFromSiteName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetSiteElementFromSiteId
		op := &xxx_GetSiteElementFromSiteIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSiteElementFromSiteIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSiteElementFromSiteID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // MapPath
		op := &xxx_MapPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MapPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MapPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostMappingExtension
type UnimplementedAppHostMappingExtensionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostMappingExtensionServer) GetSiteNameFromSiteID(context.Context, *GetSiteNameFromSiteIDRequest) (*GetSiteNameFromSiteIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMappingExtensionServer) GetSiteIDFromSiteName(context.Context, *GetSiteIDFromSiteNameRequest) (*GetSiteIDFromSiteNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMappingExtensionServer) GetSiteElementFromSiteID(context.Context, *GetSiteElementFromSiteIDRequest) (*GetSiteElementFromSiteIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostMappingExtensionServer) MapPath(context.Context, *MapPathRequest) (*MapPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostMappingExtensionServer = (*UnimplementedAppHostMappingExtensionServer)(nil)
