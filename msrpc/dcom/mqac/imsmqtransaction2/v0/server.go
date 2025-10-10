package imsmqtransaction2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	imsmqtransaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqtransaction/v0"
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
	_ = imsmqtransaction.GoPackage
)

// IMSMQTransaction2 server interface.
type Transaction2Server interface {

	// IMSMQTransaction base class.
	imsmqtransaction.TransactionServer

	// InitNew operation.
	InitNew(context.Context, *InitNewRequest) (*InitNewResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterTransaction2Server(conn dcerpc.Conn, o Transaction2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTransaction2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Transaction2SyntaxV0_0))...)
}

func NewTransaction2ServerHandle(o Transaction2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Transaction2ServerHandle(ctx, o, opNum, r)
	}
}

func Transaction2ServerHandle(ctx context.Context, o Transaction2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 10 {
		// IMSMQTransaction base method.
		return imsmqtransaction.TransactionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 10: // InitNew
		op := &xxx_InitNewOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitNewRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitNew(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Properties
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

// Unimplemented IMSMQTransaction2
type UnimplementedTransaction2Server struct {
	imsmqtransaction.UnimplementedTransactionServer
}

func (UnimplementedTransaction2Server) InitNew(context.Context, *InitNewRequest) (*InitNewResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransaction2Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Transaction2Server = (*UnimplementedTransaction2Server)(nil)
