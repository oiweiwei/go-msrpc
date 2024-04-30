package neteventforwarder

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// NetEventForwarder server interface.
type NetEventForwarderServer interface {

	// RpcNetEventOpenSession operation.
	OpenSession(context.Context, *OpenSessionRequest) (*OpenSessionResponse, error)

	// RpcNetEventReceiveData operation.
	ReceiveData(context.Context, *ReceiveDataRequest) (*ReceiveDataResponse, error)

	// RpcNetEventCloseSession operation.
	CloseSession(context.Context, *CloseSessionRequest) (*CloseSessionResponse, error)
}

func RegisterNetEventForwarderServer(conn dcerpc.Conn, o NetEventForwarderServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNetEventForwarderServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NetEventForwarderSyntaxV1_0))...)
}

func NewNetEventForwarderServerHandle(o NetEventForwarderServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NetEventForwarderServerHandle(ctx, o, opNum, r)
	}
}

func NetEventForwarderServerHandle(ctx context.Context, o NetEventForwarderServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RpcNetEventOpenSession
		in := &OpenSessionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenSession(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // RpcNetEventReceiveData
		in := &ReceiveDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReceiveData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // RpcNetEventCloseSession
		in := &CloseSessionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseSession(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
