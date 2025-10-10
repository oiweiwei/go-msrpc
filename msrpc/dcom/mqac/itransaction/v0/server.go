package itransaction

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

// ITransaction server interface.
type ITransactionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Commit operation.
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)

	// Abort operation.
	Abort(context.Context, *AbortRequest) (*AbortResponse, error)

	// GetTransactionInfo operation.
	GetTransactionInfo(context.Context, *GetTransactionInfoRequest) (*GetTransactionInfoResponse, error)
}

func RegisterITransactionServer(conn dcerpc.Conn, o ITransactionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewITransactionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ITransactionSyntaxV0_0))...)
}

func NewITransactionServerHandle(o ITransactionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ITransactionServerHandle(ctx, o, opNum, r)
	}
}

func ITransactionServerHandle(ctx context.Context, o ITransactionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Commit
		op := &xxx_CommitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Commit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Abort
		op := &xxx_AbortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AbortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Abort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetTransactionInfo
		op := &xxx_GetTransactionInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTransactionInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTransactionInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITransaction
type UnimplementedITransactionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedITransactionServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedITransactionServer) Abort(context.Context, *AbortRequest) (*AbortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedITransactionServer) GetTransactionInfo(context.Context, *GetTransactionInfoRequest) (*GetTransactionInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ITransactionServer = (*UnimplementedITransactionServer)(nil)
