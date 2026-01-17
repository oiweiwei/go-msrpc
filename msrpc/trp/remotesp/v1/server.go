package remotesp

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

// remotesp server interface.
type RemotespServer interface {

	// The RemoteSPAttach method is called by the server to establish a binding instance
	// in response to a client call to the server's ClientAttach method.
	//
	// Return Values: The method returns 0 on success; otherwise, it returns a nonzero error
	// code, as specified in [MS-ERREF]. On success, the Server List is updated with the
	// binding instance.
	//
	// Exceptions Thrown:
	//
	// The client raises an RPC_S_ACCESS_DENIED exception if it fails to obtain the RPC
	// call attributes. The client also raises an RPC_S_ACCESS_DENIED exception if it determines
	// from the call attributes that the server did not specify RPC_C_AUTHN_LEVEL_PKT_PRIVACY,
	// and the client configuration requires this authentication level.
	//
	// Except as noted above, no exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// The opnum field value for this method is 0.<10>
	Attach(context.Context, *AttachRequest) (*AttachResponse, error)

	// The RemoteSPEventProc method is called by the server to "push" completion notifications
	// and unsolicited events to the client. The server MUST call this method of the remotesp
	// interface with the endpoint and protocol sequence as specified by the connection-oriented
	// client in the ClientAttach RPC packet.
	//
	// Return Values: This method has no return values.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	//
	// The opnum field value for this method is 1.
	EventProc(context.Context, *EventProcRequest) (*EventProcResponse, error)

	// The RemoteSPDetach method is called by the server in response to a Client call to
	// the server's ClientDetach method to free the binding instance and to release the
	// associated resources.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	//
	// The opnum field value for this method is 2.
	Detach(context.Context, *DetachRequest) (*DetachResponse, error)
}

func RegisterRemotespServer(conn dcerpc.Conn, o RemotespServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemotespServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemotespSyntaxV1_0))...)
}

func NewRemotespServerHandle(o RemotespServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemotespServerHandle(ctx, o, opNum, r)
	}
}

func RemotespServerHandle(ctx context.Context, o RemotespServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RemoteSPAttach
		op := &xxx_AttachOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AttachRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Attach(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RemoteSPEventProc
		op := &xxx_EventProcOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EventProcRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EventProc(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // RemoteSPDetach
		op := &xxx_DetachOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DetachRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Detach(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented remotesp
type UnimplementedRemotespServer struct {
}

func (UnimplementedRemotespServer) Attach(context.Context, *AttachRequest) (*AttachResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemotespServer) EventProc(context.Context, *EventProcRequest) (*EventProcResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemotespServer) Detach(context.Context, *DetachRequest) (*DetachResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemotespServer = (*UnimplementedRemotespServer)(nil)
