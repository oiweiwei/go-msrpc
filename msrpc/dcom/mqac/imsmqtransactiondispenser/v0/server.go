package imsmqtransactiondispenser

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = idispatch.GoPackage
)

// IMSMQTransactionDispenser server interface.
type ImsmqTransactionDispenserServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// BeginTransaction operation.
	BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error)
}

func RegisterImsmqTransactionDispenserServer(conn dcerpc.Conn, o ImsmqTransactionDispenserServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqTransactionDispenserServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqTransactionDispenserSyntaxV0_0))...)
}

func NewImsmqTransactionDispenserServerHandle(o ImsmqTransactionDispenserServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqTransactionDispenserServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqTransactionDispenserServerHandle(ctx context.Context, o ImsmqTransactionDispenserServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
type UnimplementedImsmqTransactionDispenserServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqTransactionDispenserServer) BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqTransactionDispenserServer = (*UnimplementedImsmqTransactionDispenserServer)(nil)
