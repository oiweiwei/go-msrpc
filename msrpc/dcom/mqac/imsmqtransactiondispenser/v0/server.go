package imsmqtransactiondispenser

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

// IMSMQTransactionDispenser server interface.
type TransactionDispenserServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// BeginTransaction operation.
	BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error)
}

func RegisterTransactionDispenserServer(conn dcerpc.Conn, o TransactionDispenserServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTransactionDispenserServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TransactionDispenserSyntaxV0_0))...)
}

func NewTransactionDispenserServerHandle(o TransactionDispenserServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TransactionDispenserServerHandle(ctx, o, opNum, r)
	}
}

func TransactionDispenserServerHandle(ctx context.Context, o TransactionDispenserServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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

// Unimplemented IMSMQTransactionDispenser
type UnimplementedTransactionDispenserServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedTransactionDispenserServer) BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TransactionDispenserServer = (*UnimplementedTransactionDispenserServer)(nil)
