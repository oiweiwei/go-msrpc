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
	FrsRPCSendCommPacket(context.Context, *FrsRPCSendCommPacketRequest) (*FrsRPCSendCommPacketResponse, error)

	FrsRPCVerifyPromotionParent(context.Context, *FrsRPCVerifyPromotionParentRequest) (*FrsRPCVerifyPromotionParentResponse, error)

	FrsRPCStartPromotionParent(context.Context, *FrsRPCStartPromotionParentRequest) (*FrsRPCStartPromotionParentResponse, error)

	FrsNop(context.Context, *FrsNopRequest) (*FrsNopResponse, error)

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
		op := &xxx_FrsRPCSendCommPacketOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FrsRPCSendCommPacketRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FrsRPCSendCommPacket(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // FrsRpcVerifyPromotionParent
		op := &xxx_FrsRPCVerifyPromotionParentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FrsRPCVerifyPromotionParentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FrsRPCVerifyPromotionParent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // FrsRpcStartPromotionParent
		op := &xxx_FrsRPCStartPromotionParentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FrsRPCStartPromotionParentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FrsRPCStartPromotionParent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // FrsNOP
		op := &xxx_FrsNopOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FrsNopRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FrsNop(ctx, req)
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

func (UnimplementedFrsrpcServer) FrsRPCSendCommPacket(context.Context, *FrsRPCSendCommPacketRequest) (*FrsRPCSendCommPacketResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsrpcServer) FrsRPCVerifyPromotionParent(context.Context, *FrsRPCVerifyPromotionParentRequest) (*FrsRPCVerifyPromotionParentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsrpcServer) FrsRPCStartPromotionParent(context.Context, *FrsRPCStartPromotionParentRequest) (*FrsRPCStartPromotionParentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFrsrpcServer) FrsNop(context.Context, *FrsNopRequest) (*FrsNopResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FrsrpcServer = (*UnimplementedFrsrpcServer)(nil)
