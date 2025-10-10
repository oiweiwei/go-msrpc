package iremotenetworkconfig

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

// IRemoteNetworkConfig server interface.
type RemoteNetworkConfigServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// UpgradeRouterConfig operation.
	UpgradeRouterConfig(context.Context, *UpgradeRouterConfigRequest) (*UpgradeRouterConfigResponse, error)

	// SetUserConfig operation.
	SetUserConfig(context.Context, *SetUserConfigRequest) (*SetUserConfigResponse, error)
}

func RegisterRemoteNetworkConfigServer(conn dcerpc.Conn, o RemoteNetworkConfigServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteNetworkConfigServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteNetworkConfigSyntaxV0_0))...)
}

func NewRemoteNetworkConfigServerHandle(o RemoteNetworkConfigServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteNetworkConfigServerHandle(ctx, o, opNum, r)
	}
}

func RemoteNetworkConfigServerHandle(ctx context.Context, o RemoteNetworkConfigServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // UpgradeRouterConfig
		op := &xxx_UpgradeRouterConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpgradeRouterConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpgradeRouterConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // SetUserConfig
		op := &xxx_SetUserConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetUserConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetUserConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteNetworkConfig
type UnimplementedRemoteNetworkConfigServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRemoteNetworkConfigServer) UpgradeRouterConfig(context.Context, *UpgradeRouterConfigRequest) (*UpgradeRouterConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteNetworkConfigServer) SetUserConfig(context.Context, *SetUserConfigRequest) (*SetUserConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteNetworkConfigServer = (*UnimplementedRemoteNetworkConfigServer)(nil)
