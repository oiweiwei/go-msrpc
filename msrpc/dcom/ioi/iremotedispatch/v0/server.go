package iremotedispatch

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IRemoteDispatch server interface.
type RemoteDispatchServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The RemoteDispatchAutoDone method is called by the client to invoke a method on the
	// server.
	//
	// Return Values: An HRESULT that specifies success or failure. All success HRESULT
	// values MUST be treated as success and all failure HRESULT values MUST be treated
	// as failure.
	RemoteDispatchAutoDone(context.Context, *RemoteDispatchAutoDoneRequest) (*RemoteDispatchAutoDoneResponse, error)

	// The RemoteDispatchNotAutoDone method is called by the client to invoke a method on
	// the server.
	//
	// Return Values: An HRESULT that specifies success or failure. All success HRESULT
	// values MUST be treated as success and all failure HRESULT values MUST be treated
	// as failure.
	RemoteDispatchNotAutoDone(context.Context, *RemoteDispatchNotAutoDoneRequest) (*RemoteDispatchNotAutoDoneResponse, error)
}

func RegisterRemoteDispatchServer(conn dcerpc.Conn, o RemoteDispatchServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteDispatchServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteDispatchSyntaxV0_0))...)
}

func NewRemoteDispatchServerHandle(o RemoteDispatchServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteDispatchServerHandle(ctx, o, opNum, r)
	}
}

func RemoteDispatchServerHandle(ctx context.Context, o RemoteDispatchServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // RemoteDispatchAutoDone
		op := &xxx_RemoteDispatchAutoDoneOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoteDispatchAutoDoneRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoteDispatchAutoDone(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RemoteDispatchNotAutoDone
		op := &xxx_RemoteDispatchNotAutoDoneOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoteDispatchNotAutoDoneRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoteDispatchNotAutoDone(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteDispatch
type UnimplementedRemoteDispatchServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedRemoteDispatchServer) RemoteDispatchAutoDone(context.Context, *RemoteDispatchAutoDoneRequest) (*RemoteDispatchAutoDoneResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteDispatchServer) RemoteDispatchNotAutoDone(context.Context, *RemoteDispatchNotAutoDoneRequest) (*RemoteDispatchNotAutoDoneResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteDispatchServer = (*UnimplementedRemoteDispatchServer)(nil)
