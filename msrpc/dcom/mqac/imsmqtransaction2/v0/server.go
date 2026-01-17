package imsmqtransaction2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = imsmqtransaction.GoPackage
)

// IMSMQTransaction2 server interface.
type Transaction2Server interface {

	// IMSMQTransaction base class.
	imsmqtransaction.TransactionServer

	// The InitNew method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST initialize an MSMQ transaction object to represent an existing underlying
	// transaction object.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the Transaction instance variable is NOT NULL:
	//
	// * Return MQ_ERROR_TRANSACTION_USAGE (0xC00E0050), and take no further action.
	//
	// * Retrieve the transaction object, referred to as rTransObj , from the varTransaction
	// VARIANT based on the type of the VARIANT as follows:
	//
	// * If varTransaction.vt is VT_UNKNOWN then set rTransObj to varTransaction.punkVal.
	//
	//   - Else if varTransaction.vt is VT_UNKNOWN | VT_BYREF then set rTransObj to varTransaction.ppunkVal.
	//
	// * Else if varTransaction.vt is VT_DISPATCH then set rTransObj to varTransaction.pdispVal.
	//
	//   - Else if varTransaction.vt is VT_ DISPATCH | VT_BYREF then set rTransObj to varTransaction.ppdispVal.
	//
	// * Otherwise return E_INVALIDARG (0x80070057).
	//
	// * Retrieve the transaction object that is implementing the ITransaction interface
	// by calling IUnknown::QueryInterface ( ../ms-dcom/2b4db106-fb79-4a67-b45f-63654f19c54c
	// ) (section 3.1 ( 6b1800d7-4e98-48be-997e-57f09ce52d41 ) ) on rTransObj , passing
	// the interface identifier ( 3b7be3f7-651c-4f9c-930b-a9a7c4355ad8#gt_76ad3105-3f05-479d-a40c-c9c8fa2ebd83
	// ) of ITransaction.
	//
	// * Return E_INVALIDARG (0x80070057) if the varTransaction input parameter does not
	// implement the ITransaction interface, and take no further action.
	//
	// * Set the Transaction instance variable to the value of the transaction object previously
	// obtained.
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
