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

	// GetConnectionInterface operation.
	GetConnectionInterface(context.Context, *GetConnectionInterfaceRequest) (*GetConnectionInterfaceResponse, error)

	// GetConnectionPointContainer operation.
	GetConnectionPointContainer(context.Context, *GetConnectionPointContainerRequest) (*GetConnectionPointContainerResponse, error)

	// Advise operation.
	Advise(context.Context, *AdviseRequest) (*AdviseResponse, error)

	// Unadvise operation.
	Unadvise(context.Context, *UnadviseRequest) (*UnadviseResponse, error)

	// EnumConnections operation.
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
