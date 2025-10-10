package imsmqtransactiondispenser3

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

// IMSMQTransactionDispenser3 server interface.
type TransactionDispenser3Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// BeginTransaction operation.
	BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterTransactionDispenser3Server(conn dcerpc.Conn, o TransactionDispenser3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTransactionDispenser3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TransactionDispenser3SyntaxV0_0))...)
}

func NewTransactionDispenser3ServerHandle(o TransactionDispenser3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TransactionDispenser3ServerHandle(ctx, o, opNum, r)
	}
}

func TransactionDispenser3ServerHandle(ctx context.Context, o TransactionDispenser3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 8: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQTransactionDispenser3
type UnimplementedTransactionDispenser3Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedTransactionDispenser3Server) BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransactionDispenser3Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TransactionDispenser3Server = (*UnimplementedTransactionDispenser3Server)(nil)
