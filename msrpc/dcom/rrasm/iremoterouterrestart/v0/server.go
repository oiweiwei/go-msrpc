package iremoterouterrestart

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

// IRemoteRouterRestart server interface.
type RemoteRouterRestartServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// RestartRouter operation.
	RestartRouter(context.Context, *RestartRouterRequest) (*RestartRouterResponse, error)
}

func RegisterRemoteRouterRestartServer(conn dcerpc.Conn, o RemoteRouterRestartServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteRouterRestartServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteRouterRestartSyntaxV0_0))...)
}

func NewRemoteRouterRestartServerHandle(o RemoteRouterRestartServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteRouterRestartServerHandle(ctx, o, opNum, r)
	}
}

func RemoteRouterRestartServerHandle(ctx context.Context, o RemoteRouterRestartServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // RestartRouter
		op := &xxx_RestartRouterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestartRouterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RestartRouter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteRouterRestart
type UnimplementedRemoteRouterRestartServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRemoteRouterRestartServer) RestartRouter(context.Context, *RestartRouterRequest) (*RestartRouterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteRouterRestartServer = (*UnimplementedRemoteRouterRestartServer)(nil)
