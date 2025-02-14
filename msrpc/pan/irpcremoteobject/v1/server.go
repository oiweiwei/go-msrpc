package irpcremoteobject

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// IRPCRemoteObject server interface.
type IrpcRemoteObjectServer interface {

	// The IRPCRemoteObject_Create method creates a remote object on a server and returns
	// it to the client.
	//
	// Return Values: This method MUST return zero to indicate success, or an HRESULT error
	// value ([MS-ERREF] section 2.1.1) to indicate failure. The client MUST consider all
	// error return values fatal and report them to the higher-level caller.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	Create(context.Context, *CreateRequest) (*CreateResponse, error)

	// The IRPCRemoteObject_Delete method destroys the specified remote object.
	//
	// Return Values: This method has no return values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterIrpcRemoteObjectServer(conn dcerpc.Conn, o IrpcRemoteObjectServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIrpcRemoteObjectServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IrpcRemoteObjectSyntaxV1_0))...)
}

func NewIrpcRemoteObjectServerHandle(o IrpcRemoteObjectServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IrpcRemoteObjectServerHandle(ctx, o, opNum, r)
	}
}

func IrpcRemoteObjectServerHandle(ctx context.Context, o IrpcRemoteObjectServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // IRPCRemoteObject_Create
		op := &xxx_CreateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Create(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // IRPCRemoteObject_Delete
		op := &xxx_DeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Delete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRPCRemoteObject
type UnimplementedIrpcRemoteObjectServer struct {
}

func (UnimplementedIrpcRemoteObjectServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIrpcRemoteObjectServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IrpcRemoteObjectServer = (*UnimplementedIrpcRemoteObjectServer)(nil)
