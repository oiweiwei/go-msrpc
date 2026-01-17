package iremoteipv6config

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

// IRemoteIPV6Config server interface.
type RemoteIPv6ConfigServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetAddressList operation.
	GetAddressList(context.Context, *GetAddressListRequest) (*GetAddressListResponse, error)
}

func RegisterRemoteIPv6ConfigServer(conn dcerpc.Conn, o RemoteIPv6ConfigServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteIPv6ConfigServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteIPv6ConfigSyntaxV0_0))...)
}

func NewRemoteIPv6ConfigServerHandle(o RemoteIPv6ConfigServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteIPv6ConfigServerHandle(ctx, o, opNum, r)
	}
}

func RemoteIPv6ConfigServerHandle(ctx context.Context, o RemoteIPv6ConfigServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetAddressList
		op := &xxx_GetAddressListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAddressListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAddressList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteIPV6Config
type UnimplementedRemoteIPv6ConfigServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRemoteIPv6ConfigServer) GetAddressList(context.Context, *GetAddressListRequest) (*GetAddressListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteIPv6ConfigServer = (*UnimplementedRemoteIPv6ConfigServer)(nil)
