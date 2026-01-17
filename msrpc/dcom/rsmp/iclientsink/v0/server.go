package iclientsink

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

// IClientSink server interface.
type ClientSinkServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// OnNotify operation.
	OnNotify(context.Context, *OnNotifyRequest) (*OnNotifyResponse, error)
}

func RegisterClientSinkServer(conn dcerpc.Conn, o ClientSinkServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClientSinkServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClientSinkSyntaxV0_0))...)
}

func NewClientSinkServerHandle(o ClientSinkServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClientSinkServerHandle(ctx, o, opNum, r)
	}
}

func ClientSinkServerHandle(ctx context.Context, o ClientSinkServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // OnNotify
		op := &xxx_OnNotifyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OnNotifyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OnNotify(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IClientSink
type UnimplementedClientSinkServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedClientSinkServer) OnNotify(context.Context, *OnNotifyRequest) (*OnNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ClientSinkServer = (*UnimplementedClientSinkServer)(nil)
