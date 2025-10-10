package imsmqqueueinfos

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

// IMSMQQueueInfos server interface.
type ImsmqQueueInfosServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Reset operation.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// Next operation.
	Next(context.Context, *NextRequest) (*NextResponse, error)
}

func RegisterImsmqQueueInfosServer(conn dcerpc.Conn, o ImsmqQueueInfosServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqQueueInfosServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqQueueInfosSyntaxV0_0))...)
}

func NewImsmqQueueInfosServerHandle(o ImsmqQueueInfosServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqQueueInfosServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqQueueInfosServerHandle(ctx context.Context, o ImsmqQueueInfosServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Reset
		op := &xxx_ResetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Reset(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Next
		op := &xxx_NextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Next(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQueueInfos
type UnimplementedImsmqQueueInfosServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqQueueInfosServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueInfosServer) Next(context.Context, *NextRequest) (*NextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqQueueInfosServer = (*UnimplementedImsmqQueueInfosServer)(nil)
