package imsmqtransaction

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

// IMSMQTransaction server interface.
type ImsmqTransactionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Transaction operation.
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)

	// Commit operation.
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)

	// Abort operation.
	Abort(context.Context, *AbortRequest) (*AbortResponse, error)
}

func RegisterImsmqTransactionServer(conn dcerpc.Conn, o ImsmqTransactionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqTransactionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqTransactionSyntaxV0_0))...)
}

func NewImsmqTransactionServerHandle(o ImsmqTransactionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqTransactionServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqTransactionServerHandle(ctx context.Context, o ImsmqTransactionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Transaction
		op := &xxx_GetTransactionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTransactionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTransaction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Commit
		op := &xxx_CommitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Commit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Abort
		op := &xxx_AbortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AbortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Abort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQTransaction
type UnimplementedImsmqTransactionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqTransactionServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqTransactionServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqTransactionServer) Abort(context.Context, *AbortRequest) (*AbortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqTransactionServer = (*UnimplementedImsmqTransactionServer)(nil)
