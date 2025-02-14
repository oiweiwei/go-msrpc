package icluscfgasyncevictcleanup

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

// IClusCfgAsyncEvictCleanup server interface.
type AsyncEvictCleanupServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// CleanupNode operation.
	CleanupNode(context.Context, *CleanupNodeRequest) (*CleanupNodeResponse, error)
}

func RegisterAsyncEvictCleanupServer(conn dcerpc.Conn, o AsyncEvictCleanupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAsyncEvictCleanupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AsyncEvictCleanupSyntaxV0_0))...)
}

func NewAsyncEvictCleanupServerHandle(o AsyncEvictCleanupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AsyncEvictCleanupServerHandle(ctx, o, opNum, r)
	}
}

func AsyncEvictCleanupServerHandle(ctx context.Context, o AsyncEvictCleanupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CleanupNode
		op := &xxx_CleanupNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CleanupNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CleanupNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IClusCfgAsyncEvictCleanup
type UnimplementedAsyncEvictCleanupServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedAsyncEvictCleanupServer) CleanupNode(context.Context, *CleanupNodeRequest) (*CleanupNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AsyncEvictCleanupServer = (*UnimplementedAsyncEvictCleanupServer)(nil)
