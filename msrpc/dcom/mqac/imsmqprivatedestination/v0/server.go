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
type ImsmqPrivateDestinationServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Handle operation.
	GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error)

	// Handle operation.
	SetHandle(context.Context, *SetHandleRequest) (*SetHandleResponse, error)
}

func RegisterImsmqPrivateDestinationServer(conn dcerpc.Conn, o ImsmqPrivateDestinationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqPrivateDestinationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqPrivateDestinationSyntaxV0_0))...)
}

func NewImsmqPrivateDestinationServerHandle(o ImsmqPrivateDestinationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqPrivateDestinationServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqPrivateDestinationServerHandle(ctx context.Context, o ImsmqPrivateDestinationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
type UnimplementedImsmqPrivateDestinationServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqPrivateDestinationServer) GetHandle(context.Context, *GetHandleRequest) (*GetHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqPrivateDestinationServer) SetHandle(context.Context, *SetHandleRequest) (*SetHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqPrivateDestinationServer = (*UnimplementedImsmqPrivateDestinationServer)(nil)
