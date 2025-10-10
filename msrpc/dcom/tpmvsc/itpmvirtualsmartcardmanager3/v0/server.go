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
type TpmVirtualSmartCardManager3Server interface {

	// ITpmVirtualSmartCardManager2 base class.
	itpmvirtualsmartcardmanager2.TpmVirtualSmartCardManager2Server

	CreateVirtualSmartCardWithAttestation(context.Context, *CreateVirtualSmartCardWithAttestationRequest) (*CreateVirtualSmartCardWithAttestationResponse, error)
}

func RegisterTpmVirtualSmartCardManager3Server(conn dcerpc.Conn, o TpmVirtualSmartCardManager3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTpmVirtualSmartCardManager3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TpmVirtualSmartCardManager3SyntaxV0_0))...)
}

func NewTpmVirtualSmartCardManager3ServerHandle(o TpmVirtualSmartCardManager3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TpmVirtualSmartCardManager3ServerHandle(ctx, o, opNum, r)
	}
}

func TpmVirtualSmartCardManager3ServerHandle(ctx context.Context, o TpmVirtualSmartCardManager3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 6 {
		// ITpmVirtualSmartCardManager2 base method.
		return itpmvirtualsmartcardmanager2.TpmVirtualSmartCardManager2ServerHandle(ctx, o, opNum, r)
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
type UnimplementedTpmVirtualSmartCardManager3Server struct {
	itpmvirtualsmartcardmanager2.UnimplementedTpmVirtualSmartCardManager2Server
}

func (UnimplementedTpmVirtualSmartCardManager3Server) CreateVirtualSmartCardWithAttestation(context.Context, *CreateVirtualSmartCardWithAttestationRequest) (*CreateVirtualSmartCardWithAttestationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TpmVirtualSmartCardManager3Server = (*UnimplementedTpmVirtualSmartCardManager3Server)(nil)
