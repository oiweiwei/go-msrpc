package frsrpc

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

// frsrpc server interface.
type FrsrpcServer interface {
	SendCommPacket(context.Context, *SendCommPacketRequest) (*SendCommPacketResponse, error)

	VerifyPromotionParent(context.Context, *VerifyPromotionParentRequest) (*VerifyPromotionParentResponse, error)

	StartPromotionParent(context.Context, *StartPromotionParentRequest) (*StartPromotionParentResponse, error)

	Noop(context.Context, *NoopRequest) (*NoopResponse, error)

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire
}

func RegisterFrsrpcServer(conn dcerpc.Conn, o FrsrpcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFrsrpcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FrsrpcSyntaxV1_1))...)
}

func NewFrsrpcServerHandle(o FrsrpcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FrsrpcServerHandle(ctx, o, opNum, r)
	}
}

func FrsrpcServerHandle(ctx context.Context, o FrsrpcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // FrsRpcSendCommPkt
		op := &xxx_SendCommPacketOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendCommPacketRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SendCommPacket(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // FrsRpcVerifyPromotionParent
		op := &xxx_VerifyPromotionParentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &VerifyPromotionParentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.VerifyPromotionParent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // FrsRpcStartPromotionParent
		op := &xxx_StartPromotionParentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartPromotionParentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartPromotionParent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // FrsNOP
		op := &xxx_NoopOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NoopRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Noop(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // Opnum8NotUsedOnWire
		// Opnum8NotUsedOnWire
		return nil, nil
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 10: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented frsrpc
type UnimplementedFrsrpcServer struct {
}

func (UnimplementedFrsrpcServer) SendCommPacket(context.Context, *SendCommPacketRequest) (*SendCommPacketResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsrpcServer) VerifyPromotionParent(context.Context, *VerifyPromotionParentRequest) (*VerifyPromotionParentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsrpcServer) StartPromotionParent(context.Context, *StartPromotionParentRequest) (*StartPromotionParentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsrpcServer) Noop(context.Context, *NoopRequest) (*NoopResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FrsrpcServer = (*UnimplementedFrsrpcServer)(nil)
