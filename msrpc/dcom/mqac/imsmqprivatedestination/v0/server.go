package imsmqprivatedestination

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

// IMSMQPrivateDestination server interface.
type PrivateDestinationServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Handle operation.
	GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error)

	// Handle operation.
	SetHandle(context.Context, *SetHandleRequest) (*SetHandleResponse, error)
}

func RegisterPrivateDestinationServer(conn dcerpc.Conn, o PrivateDestinationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPrivateDestinationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PrivateDestinationSyntaxV0_0))...)
}

func NewPrivateDestinationServerHandle(o PrivateDestinationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PrivateDestinationServerHandle(ctx, o, opNum, r)
	}
}

func PrivateDestinationServerHandle(ctx context.Context, o PrivateDestinationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Handle
		op := &xxx_GetHandleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHandleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHandle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Handle
		op := &xxx_SetHandleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetHandleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetHandle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQPrivateDestination
type UnimplementedPrivateDestinationServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedPrivateDestinationServer) GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPrivateDestinationServer) SetHandle(context.Context, *SetHandleRequest) (*SetHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PrivateDestinationServer = (*UnimplementedPrivateDestinationServer)(nil)
