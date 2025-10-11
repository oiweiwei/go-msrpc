package iapphostadminmanager

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

// IAppHostAdminManager server interface.
type AppHostAdminManagerServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetAdminSection operation.
	GetAdminSection(context.Context, *GetAdminSectionRequest) (*GetAdminSectionResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// SetMetadata operation.
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)

	// The ConfigManager method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns an IAppHostConfigManager that provides direct access to the supported
	// path hierarchy of the administration system and access to individual IAppHostElement
	// objects that are contained within.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If successful, *ppConfigManager MUST NOT be NULL.
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
	GetConfigManager(context.Context, *GetConfigManagerRequest) (*GetConfigManagerResponse, error)
}

func RegisterAppHostAdminManagerServer(conn dcerpc.Conn, o AppHostAdminManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostAdminManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostAdminManagerSyntaxV0_0))...)
}

func NewAppHostAdminManagerServerHandle(o AppHostAdminManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostAdminManagerServerHandle(ctx, o, opNum, r)
	}
}

func AppHostAdminManagerServerHandle(ctx context.Context, o AppHostAdminManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetAdminSection
		op := &xxx_GetAdminSectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAdminSectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAdminSection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetMetadata
		op := &xxx_GetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // SetMetadata
		op := &xxx_SetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // ConfigManager
		op := &xxx_GetConfigManagerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigManagerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfigManager(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostAdminManager
type UnimplementedAppHostAdminManagerServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostAdminManagerServer) GetAdminSection(context.Context, *GetAdminSectionRequest) (*GetAdminSectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostAdminManagerServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostAdminManagerServer) SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostAdminManagerServer) GetConfigManager(context.Context, *GetConfigManagerRequest) (*GetConfigManagerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostAdminManagerServer = (*UnimplementedAppHostAdminManagerServer)(nil)
