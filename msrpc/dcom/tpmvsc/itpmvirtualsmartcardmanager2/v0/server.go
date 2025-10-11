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
type VirtualSmartCardManager2Server interface {

	// ITpmVirtualSmartCardManager base class.
	itpmvirtualsmartcardmanager.VirtualSmartCardManagerServer

	// This method is invoked by the requestor to create a VSC with the specified PIN policy
	// on the target.
	//
	// Return Values:  The server MUST return 0 if it successfully creates the new VSC,
	// and a nonzero value otherwise.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CreateVirtualSmartCardWithPINPolicy(context.Context, *CreateVirtualSmartCardWithPINPolicyRequest) (*CreateVirtualSmartCardWithPINPolicyResponse, error)
}

func RegisterVirtualSmartCardManager2Server(conn dcerpc.Conn, o VirtualSmartCardManager2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVirtualSmartCardManager2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VirtualSmartCardManager2SyntaxV0_0))...)
}

func NewVirtualSmartCardManager2ServerHandle(o VirtualSmartCardManager2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VirtualSmartCardManager2ServerHandle(ctx, o, opNum, r)
	}
}

func VirtualSmartCardManager2ServerHandle(ctx context.Context, o VirtualSmartCardManager2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 5 {
		// ITpmVirtualSmartCardManager base method.
		return itpmvirtualsmartcardmanager.VirtualSmartCardManagerServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 5: // CreateVirtualSmartCardWithPinPolicy
		op := &xxx_CreateVirtualSmartCardWithPINPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVirtualSmartCardWithPINPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVirtualSmartCardWithPINPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITpmVirtualSmartCardManager2
type UnimplementedVirtualSmartCardManager2Server struct {
	itpmvirtualsmartcardmanager.UnimplementedVirtualSmartCardManagerServer
}

func (UnimplementedVirtualSmartCardManager2Server) CreateVirtualSmartCardWithPINPolicy(context.Context, *CreateVirtualSmartCardWithPINPolicyRequest) (*CreateVirtualSmartCardWithPINPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VirtualSmartCardManager2Server = (*UnimplementedVirtualSmartCardManager2Server)(nil)
