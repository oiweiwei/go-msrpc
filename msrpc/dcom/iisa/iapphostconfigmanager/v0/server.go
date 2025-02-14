package iapphostconfigmanager

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

// IAppHostConfigManager server interface.
type AppHostConfigManagerServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetConfigFile operation.
	GetConfigFile(context.Context, *GetConfigFileRequest) (*GetConfigFileResponse, error)

	// GetUniqueConfigPath operation.
	GetUniqueConfigPath(context.Context, *GetUniqueConfigPathRequest) (*GetUniqueConfigPathResponse, error)
}

func RegisterAppHostConfigManagerServer(conn dcerpc.Conn, o AppHostConfigManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostConfigManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostConfigManagerSyntaxV0_0))...)
}

func NewAppHostConfigManagerServerHandle(o AppHostConfigManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostConfigManagerServerHandle(ctx, o, opNum, r)
	}
}

func AppHostConfigManagerServerHandle(ctx context.Context, o AppHostConfigManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetConfigFile
		op := &xxx_GetConfigFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfigFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetUniqueConfigPath
		op := &xxx_GetUniqueConfigPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUniqueConfigPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUniqueConfigPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostConfigManager
type UnimplementedAppHostConfigManagerServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostConfigManagerServer) GetConfigFile(context.Context, *GetConfigFileRequest) (*GetConfigFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigManagerServer) GetUniqueConfigPath(context.Context, *GetUniqueConfigPathRequest) (*GetUniqueConfigPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostConfigManagerServer = (*UnimplementedAppHostConfigManagerServer)(nil)
