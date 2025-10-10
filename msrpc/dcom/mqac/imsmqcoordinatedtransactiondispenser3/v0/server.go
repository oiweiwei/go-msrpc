package imsmqcoordinatedtransactiondispenser3

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

// IMSMQCoordinatedTransactionDispenser3 server interface.
type CoordinatedTransactionDispenser3Server interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// BeginTransaction operation.
	BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterCoordinatedTransactionDispenser3Server(conn dcerpc.Conn, o CoordinatedTransactionDispenser3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCoordinatedTransactionDispenser3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CoordinatedTransactionDispenser3SyntaxV0_0))...)
}

func NewCoordinatedTransactionDispenser3ServerHandle(o CoordinatedTransactionDispenser3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CoordinatedTransactionDispenser3ServerHandle(ctx, o, opNum, r)
	}
}

func CoordinatedTransactionDispenser3ServerHandle(ctx context.Context, o CoordinatedTransactionDispenser3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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

// Unimplemented IMSMQCoordinatedTransactionDispenser3
type UnimplementedCoordinatedTransactionDispenser3Server struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedCoordinatedTransactionDispenser3Server) BeginTransaction(context.Context, *BeginTransactionRequest) (*BeginTransactionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCoordinatedTransactionDispenser3Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CoordinatedTransactionDispenser3Server = (*UnimplementedCoordinatedTransactionDispenser3Server)(nil)
