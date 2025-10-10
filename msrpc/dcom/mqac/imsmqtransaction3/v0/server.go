package imsmqtransaction3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	imsmqtransaction2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqtransaction2/v0"
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
	_ = imsmqtransaction2.GoPackage
)

// IMSMQTransaction3 server interface.
type ImsmqTransaction3Server interface {

	// IMSMQTransaction2 base class.
	imsmqtransaction2.ImsmqTransaction2Server

	// ITransaction operation.
	GetITransaction(context.Context, *GetITransactionRequest) (*GetITransactionResponse, error)
}

func RegisterImsmqTransaction3Server(conn dcerpc.Conn, o ImsmqTransaction3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqTransaction3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqTransaction3SyntaxV0_0))...)
}

func NewImsmqTransaction3ServerHandle(o ImsmqTransaction3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqTransaction3ServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqTransaction3ServerHandle(ctx context.Context, o ImsmqTransaction3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IMSMQTransaction2 base method.
		return imsmqtransaction2.ImsmqTransaction2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // ITransaction
		op := &xxx_GetITransactionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetITransactionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetITransaction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQTransaction3
type UnimplementedImsmqTransaction3Server struct {
	imsmqtransaction2.UnimplementedImsmqTransaction2Server
}

func (UnimplementedImsmqTransaction3Server) GetITransaction(context.Context, *GetITransactionRequest) (*GetITransactionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqTransaction3Server = (*UnimplementedImsmqTransaction3Server)(nil)
