package iconnectionpointcontainer

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

// IConnectionPointContainer server interface.
type ConnectionPointContainerServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The EnumConnectionPoints method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return a pointer to an IEnumConnectionPoints interface
	// pointer, as defined in [MSDN-ECP], that can be used by the client to enumerate all
	// the IConnectionPoint implementations for the MSMQEvent object.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<97>
	EnumConnectionPoints(context.Context, *EnumConnectionPointsRequest) (*EnumConnectionPointsResponse, error)

	// The FindConnectionPoint method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return a pointer to an IConnectionPoint interface pointer
	// for a specified IID.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.<98>
	FindConnectionPoint(context.Context, *FindConnectionPointRequest) (*FindConnectionPointResponse, error)
}

func RegisterConnectionPointContainerServer(conn dcerpc.Conn, o ConnectionPointContainerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewConnectionPointContainerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ConnectionPointContainerSyntaxV0_0))...)
}

func NewConnectionPointContainerServerHandle(o ConnectionPointContainerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ConnectionPointContainerServerHandle(ctx, o, opNum, r)
	}
}

func ConnectionPointContainerServerHandle(ctx context.Context, o ConnectionPointContainerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // EnumConnectionPoints
		op := &xxx_EnumConnectionPointsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumConnectionPointsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumConnectionPoints(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // FindConnectionPoint
		op := &xxx_FindConnectionPointOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FindConnectionPointRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FindConnectionPoint(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IConnectionPointContainer
type UnimplementedConnectionPointContainerServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedConnectionPointContainerServer) EnumConnectionPoints(context.Context, *EnumConnectionPointsRequest) (*EnumConnectionPointsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedConnectionPointContainerServer) FindConnectionPoint(context.Context, *FindConnectionPointRequest) (*FindConnectionPointResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ConnectionPointContainerServer = (*UnimplementedConnectionPointContainerServer)(nil)
