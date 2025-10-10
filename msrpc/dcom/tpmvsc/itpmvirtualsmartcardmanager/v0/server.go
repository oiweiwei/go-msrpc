package itpmvirtualsmartcardmanager

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

// ITpmVirtualSmartCardManager server interface.
type TpmVirtualSmartCardManagerServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	CreateVirtualSmartCard(context.Context, *CreateVirtualSmartCardRequest) (*CreateVirtualSmartCardResponse, error)

	DestroyVirtualSmartCard(context.Context, *DestroyVirtualSmartCardRequest) (*DestroyVirtualSmartCardResponse, error)
}

func RegisterTpmVirtualSmartCardManagerServer(conn dcerpc.Conn, o TpmVirtualSmartCardManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTpmVirtualSmartCardManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TpmVirtualSmartCardManagerSyntaxV0_0))...)
}

func NewTpmVirtualSmartCardManagerServerHandle(o TpmVirtualSmartCardManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TpmVirtualSmartCardManagerServerHandle(ctx, o, opNum, r)
	}
}

func TpmVirtualSmartCardManagerServerHandle(ctx context.Context, o TpmVirtualSmartCardManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CreateVirtualSmartCard
		op := &xxx_CreateVirtualSmartCardOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVirtualSmartCardRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVirtualSmartCard(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // DestroyVirtualSmartCard
		op := &xxx_DestroyVirtualSmartCardOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DestroyVirtualSmartCardRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DestroyVirtualSmartCard(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITpmVirtualSmartCardManager
type UnimplementedTpmVirtualSmartCardManagerServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedTpmVirtualSmartCardManagerServer) CreateVirtualSmartCard(context.Context, *CreateVirtualSmartCardRequest) (*CreateVirtualSmartCardResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTpmVirtualSmartCardManagerServer) DestroyVirtualSmartCard(context.Context, *DestroyVirtualSmartCardRequest) (*DestroyVirtualSmartCardResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TpmVirtualSmartCardManagerServer = (*UnimplementedTpmVirtualSmartCardManagerServer)(nil)
