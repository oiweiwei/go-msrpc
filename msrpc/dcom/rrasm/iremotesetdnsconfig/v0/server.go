package iremotesetdnsconfig

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

// IRemoteSetDnsConfig server interface.
type RemoteSetDNSConfigServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// SetDnsConfig operation.
	SetDNSConfig(context.Context, *SetDNSConfigRequest) (*SetDNSConfigResponse, error)
}

func RegisterRemoteSetDNSConfigServer(conn dcerpc.Conn, o RemoteSetDNSConfigServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteSetDNSConfigServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteSetDNSConfigSyntaxV0_0))...)
}

func NewRemoteSetDNSConfigServerHandle(o RemoteSetDNSConfigServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteSetDNSConfigServerHandle(ctx, o, opNum, r)
	}
}

func RemoteSetDNSConfigServerHandle(ctx context.Context, o RemoteSetDNSConfigServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // SetDnsConfig
		op := &xxx_SetDNSConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDNSConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDNSConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteSetDnsConfig
type UnimplementedRemoteSetDNSConfigServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRemoteSetDNSConfigServer) SetDNSConfig(context.Context, *SetDNSConfigRequest) (*SetDNSConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteSetDNSConfigServer = (*UnimplementedRemoteSetDNSConfigServer)(nil)
