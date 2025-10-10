package imsmqquery

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

// IMSMQQuery server interface.
type ImsmqQueryServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// LookupQueue operation.
	LookupQueue(context.Context, *LookupQueueRequest) (*LookupQueueResponse, error)
}

func RegisterImsmqQueryServer(conn dcerpc.Conn, o ImsmqQueryServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqQueryServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqQuerySyntaxV0_0))...)
}

func NewImsmqQueryServerHandle(o ImsmqQueryServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqQueryServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqQueryServerHandle(ctx context.Context, o ImsmqQueryServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // LookupQueue
		op := &xxx_LookupQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQuery
type UnimplementedImsmqQueryServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqQueryServer) LookupQueue(context.Context, *LookupQueueRequest) (*LookupQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqQueryServer = (*UnimplementedImsmqQueryServer)(nil)
