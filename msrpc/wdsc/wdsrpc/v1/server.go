package wdsrpc

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

// WdsRpc server interface.
type WdsRPCServer interface {

	// The WdsRpcMessage (opnum 0) method sends the request packet to the server and returns
	// the corresponding reply packet.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// non-zero Win32 error code value if an error occurred.
	Message(context.Context, *MessageRequest) (*MessageResponse, error)
}

func RegisterWdsRPCServer(conn dcerpc.Conn, o WdsRPCServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWdsRPCServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WdsRPCSyntaxV1_0))...)
}

func NewWdsRPCServerHandle(o WdsRPCServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WdsRPCServerHandle(ctx, o, opNum, r)
	}
}

func WdsRPCServerHandle(ctx context.Context, o WdsRPCServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // WdsRpcMessage
		op := &xxx_MessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Message(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented WdsRpc
type UnimplementedWdsRPCServer struct {
}

func (UnimplementedWdsRPCServer) Message(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WdsRPCServer = (*UnimplementedWdsRPCServer)(nil)
