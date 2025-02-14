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
		op := &xxx_OpenSessionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenSessionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenSession(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RpcNetEventReceiveData
		op := &xxx_ReceiveDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReceiveDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReceiveData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // RpcNetEventCloseSession
		op := &xxx_CloseSessionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseSessionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseSession(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented NetEventForwarder
type UnimplementedNetEventForwarderServer struct {
}

func (UnimplementedNetEventForwarderServer) OpenSession(context.Context, *OpenSessionRequest) (*OpenSessionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetEventForwarderServer) ReceiveData(context.Context, *ReceiveDataRequest) (*ReceiveDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetEventForwarderServer) CloseSession(context.Context, *CloseSessionRequest) (*CloseSessionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NetEventForwarderServer = (*UnimplementedNetEventForwarderServer)(nil)
