package imsmqcoordinatedtransactiondispenser

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IMSMQCoordinatedTransactionDispenser server interface.
type CoordinatedTransactionDispenserServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// BeginTransaction operation.
	BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error)
}

func RegisterCoordinatedTransactionDispenserServer(conn dcerpc.Conn, o CoordinatedTransactionDispenserServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCoordinatedTransactionDispenserServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CoordinatedTransactionDispenserSyntaxV0_0))...)
}

func NewCoordinatedTransactionDispenserServerHandle(o CoordinatedTransactionDispenserServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CoordinatedTransactionDispenserServerHandle(ctx, o, opNum, r)
	}
}

func CoordinatedTransactionDispenserServerHandle(ctx context.Context, o CoordinatedTransactionDispenserServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // BeginTransaction
		op := &xxx_BeginTransactionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BeginTransactionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BeginTransaction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQCoordinatedTransactionDispenser
type UnimplementedCoordinatedTransactionDispenserServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedCoordinatedTransactionDispenserServer) BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CoordinatedTransactionDispenserServer = (*UnimplementedCoordinatedTransactionDispenserServer)(nil)
