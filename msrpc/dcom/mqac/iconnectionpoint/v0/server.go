package iconnectionpoint

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IConnectionPoint server interface.
type ConnectionPointServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetConnectionInterface method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return a pointer to the IID of the interface that client
	// applications MUST implement for the MSMQEvent object to be able to notify them when
	// an event is raised.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<92>
	GetConnectionInterface(context.Context, *GetConnectionInterfaceRequest) (*GetConnectionInterfaceResponse, error)

	// The GetConnectionPointContainer method is received by the server in an RPC_REQUEST
	// packet. In response, the server MUST return a pointer to a pointer to an IConnectionPointContainer
	// interface for the MSMQEvent object.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<93>
	GetConnectionPointContainer(context.Context, *GetConnectionPointContainerRequest) (*GetConnectionPointContainerResponse, error)

	// The Advise method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST register a client's callback object to receive event notifications.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<94>
	Advise(context.Context, *AdviseRequest) (*AdviseResponse, error)

	// The Unadvise method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST deregister a client's callback object to cease receiving event notifications.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<95>
	Unadvise(context.Context, *UnadviseRequest) (*UnadviseResponse, error)

	// The EnumConnections method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return a pointer to an IEnumConnections interface pointer,
	// as defined in [MSDN-EC], for the client to enumerate all the currently registered
	// callback objects.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<96>
	EnumConnections(context.Context, *EnumConnectionsRequest) (*EnumConnectionsResponse, error)
}

func RegisterConnectionPointServer(conn dcerpc.Conn, o ConnectionPointServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewConnectionPointServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ConnectionPointSyntaxV0_0))...)
}

func NewConnectionPointServerHandle(o ConnectionPointServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ConnectionPointServerHandle(ctx, o, opNum, r)
	}
}

func ConnectionPointServerHandle(ctx context.Context, o ConnectionPointServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetConnectionInterface
		op := &xxx_GetConnectionInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConnectionInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConnectionInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetConnectionPointContainer
		op := &xxx_GetConnectionPointContainerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConnectionPointContainerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConnectionPointContainer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Advise
		op := &xxx_AdviseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AdviseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Advise(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Unadvise
		op := &xxx_UnadviseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnadviseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Unadvise(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // EnumConnections
		op := &xxx_EnumConnectionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumConnectionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumConnections(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IConnectionPoint
type UnimplementedConnectionPointServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedConnectionPointServer) GetConnectionInterface(context.Context, *GetConnectionInterfaceRequest) (*GetConnectionInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConnectionPointServer) GetConnectionPointContainer(context.Context, *GetConnectionPointContainerRequest) (*GetConnectionPointContainerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConnectionPointServer) Advise(context.Context, *AdviseRequest) (*AdviseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConnectionPointServer) Unadvise(context.Context, *UnadviseRequest) (*UnadviseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConnectionPointServer) EnumConnections(context.Context, *EnumConnectionsRequest) (*EnumConnectionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ConnectionPointServer = (*UnimplementedConnectionPointServer)(nil)
