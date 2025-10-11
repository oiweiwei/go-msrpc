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
type Transaction3Server interface {

	// IMSMQTransaction2 base class.
	imsmqtransaction2.Transaction2Server

	// The ITransaction method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the ITransaction interface on the underlying transaction object.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// The pvarITransaction output parameter MUST be set to the ITransaction interface pointer
	// of the Transaction instance variable.
	GetITransaction(context.Context, *GetITransactionRequest) (*GetITransactionResponse, error)
}

func RegisterTransaction3Server(conn dcerpc.Conn, o Transaction3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTransaction3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Transaction3SyntaxV0_0))...)
}

func NewTransaction3ServerHandle(o Transaction3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Transaction3ServerHandle(ctx, o, opNum, r)
	}
}

func Transaction3ServerHandle(ctx context.Context, o Transaction3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IMSMQTransaction2 base method.
		return imsmqtransaction2.Transaction2ServerHandle(ctx, o, opNum, r)
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
type UnimplementedTransaction3Server struct {
	imsmqtransaction2.UnimplementedTransaction2Server
}

func (UnimplementedTransaction3Server) GetITransaction(context.Context, *GetITransactionRequest) (*GetITransactionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Transaction3Server = (*UnimplementedTransaction3Server)(nil)
