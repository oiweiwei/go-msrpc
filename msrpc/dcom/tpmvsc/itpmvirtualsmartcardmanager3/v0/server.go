package itpmvirtualsmartcardmanager3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	itpmvirtualsmartcardmanager2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/tpmvsc/itpmvirtualsmartcardmanager2/v0"
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
	_ = itpmvirtualsmartcardmanager2.GoPackage
)

// ITpmVirtualSmartCardManager3 server interface.
type VirtualSmartCardManager3Server interface {

	// ITpmVirtualSmartCardManager2 base class.
	itpmvirtualsmartcardmanager2.VirtualSmartCardManager2Server

	CreateVirtualSmartCardWithAttestation(context.Context, *CreateVirtualSmartCardWithAttestationRequest) (*CreateVirtualSmartCardWithAttestationResponse, error)
}

func RegisterVirtualSmartCardManager3Server(conn dcerpc.Conn, o VirtualSmartCardManager3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVirtualSmartCardManager3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VirtualSmartCardManager3SyntaxV0_0))...)
}

func NewVirtualSmartCardManager3ServerHandle(o VirtualSmartCardManager3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VirtualSmartCardManager3ServerHandle(ctx, o, opNum, r)
	}
}

func VirtualSmartCardManager3ServerHandle(ctx context.Context, o VirtualSmartCardManager3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 6 {
		// ITpmVirtualSmartCardManager2 base method.
		return itpmvirtualsmartcardmanager2.VirtualSmartCardManager2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 6: // CreateVirtualSmartCardWithAttestation
		op := &xxx_CreateVirtualSmartCardWithAttestationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateVirtualSmartCardWithAttestationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateVirtualSmartCardWithAttestation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITpmVirtualSmartCardManager3
type UnimplementedVirtualSmartCardManager3Server struct {
	itpmvirtualsmartcardmanager2.UnimplementedVirtualSmartCardManager2Server
}

func (UnimplementedVirtualSmartCardManager3Server) CreateVirtualSmartCardWithAttestation(context.Context, *CreateVirtualSmartCardWithAttestationRequest) (*CreateVirtualSmartCardWithAttestationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ VirtualSmartCardManager3Server = (*UnimplementedVirtualSmartCardManager3Server)(nil)
