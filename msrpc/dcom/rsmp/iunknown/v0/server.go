package iunknown

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

// IUnknown server interface.
type UnknownServer interface {

	// QueryInterface operation.
	QueryInterface(context.Context, *QueryInterfaceRequest) (*QueryInterfaceResponse, error)

	// AddRef operation.
	AddReference(context.Context, *AddReferenceRequest) (*AddReferenceResponse, error)

	// Release operation.
	Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error)
}

func RegisterUnknownServer(conn dcerpc.Conn, o UnknownServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUnknownServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UnknownSyntaxV0_0))...)
}

func NewUnknownServerHandle(o UnknownServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UnknownServerHandle(ctx, o, opNum, r)
	}
}

func UnknownServerHandle(ctx context.Context, o UnknownServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // QueryInterface
		op := &xxx_QueryInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // AddRef
		op := &xxx_AddReferenceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddReferenceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddReference(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // Release
		op := &xxx_ReleaseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReleaseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Release(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUnknown
type UnimplementedUnknownServer struct {
}

func (UnimplementedUnknownServer) QueryInterface(context.Context, *QueryInterfaceRequest) (*QueryInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUnknownServer) AddReference(context.Context, *AddReferenceRequest) (*AddReferenceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUnknownServer) Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UnknownServer = (*UnimplementedUnknownServer)(nil)
