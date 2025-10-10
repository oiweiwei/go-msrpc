package itpmvirtualsmartcardmanager2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	itpmvirtualsmartcardmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc/itpmvirtualsmartcardmanager/v0"
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
	_ = itpmvirtualsmartcardmanager.GoPackage
)

// ITpmVirtualSmartCardManager2 server interface.
type TpmVirtualSmartCardManager2Server interface {

	// ITpmVirtualSmartCardManager base class.
	itpmvirtualsmartcardmanager.TpmVirtualSmartCardManagerServer

	CreateVirtualSmartCardWithPinPolicy(context.Context, *CreateVirtualSmartCardWithPinPolicyRequest) (*CreateVirtualSmartCardWithPinPolicyResponse, error)
}

func RegisterTpmVirtualSmartCardManager2Server(conn dcerpc.Conn, o TpmVirtualSmartCardManager2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTpmVirtualSmartCardManager2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TpmVirtualSmartCardManager2SyntaxV0_0))...)
}

func NewTpmVirtualSmartCardManager2ServerHandle(o TpmVirtualSmartCardManager2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TpmVirtualSmartCardManager2ServerHandle(ctx, o, opNum, r)
	}
}

func TpmVirtualSmartCardManager2ServerHandle(ctx context.Context, o TpmVirtualSmartCardManager2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 5 {
		// ITpmVirtualSmartCardManager base method.
		return itpmvirtualsmartcardmanager.TpmVirtualSmartCardManagerServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 5: // CreateVirtualSmartCardWithPinPolicy
		op := &xxx_CreateVirtualSmartCardWithPinPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVirtualSmartCardWithPinPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVirtualSmartCardWithPinPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITpmVirtualSmartCardManager2
type UnimplementedTpmVirtualSmartCardManager2Server struct {
	itpmvirtualsmartcardmanager.UnimplementedTpmVirtualSmartCardManagerServer
}

func (UnimplementedTpmVirtualSmartCardManager2Server) CreateVirtualSmartCardWithPinPolicy(context.Context, *CreateVirtualSmartCardWithPinPolicyRequest) (*CreateVirtualSmartCardWithPinPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TpmVirtualSmartCardManager2Server = (*UnimplementedTpmVirtualSmartCardManager2Server)(nil)
