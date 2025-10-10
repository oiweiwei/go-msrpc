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
type ImsmqTransaction2Server interface {

	// IMSMQTransaction base class.
	imsmqtransaction.ImsmqTransactionServer

	// InitNew operation.
	InitNew(context.Context, *InitNewRequest) (*InitNewResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterImsmqTransaction2Server(conn dcerpc.Conn, o ImsmqTransaction2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqTransaction2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqTransaction2SyntaxV0_0))...)
}

func NewImsmqTransaction2ServerHandle(o ImsmqTransaction2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqTransaction2ServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqTransaction2ServerHandle(ctx context.Context, o ImsmqTransaction2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 10 {
		// IMSMQTransaction base method.
		return imsmqtransaction.ImsmqTransactionServerHandle(ctx, o, opNum, r)
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
type UnimplementedImsmqTransaction2Server struct {
	imsmqtransaction.UnimplementedImsmqTransactionServer
}

func (UnimplementedImsmqTransaction2Server) InitNew(context.Context, *InitNewRequest) (*InitNewResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqTransaction2Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqTransaction2Server = (*UnimplementedImsmqTransaction2Server)(nil)
