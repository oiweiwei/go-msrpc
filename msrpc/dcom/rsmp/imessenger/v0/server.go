package imessenger

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = iunknown.GoPackage
)

// IMessenger server interface.
type MessengerServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The SendMessage method adds a message to the send queue.
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)

	// The RecallMessage method retrieves a message from the send queue.
	RecallMessage(context.Context, *RecallMessageRequest) (*RecallMessageResponse, error)
}

func RegisterMessengerServer(conn dcerpc.Conn, o MessengerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewMessengerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(MessengerSyntaxV0_0))...)
}

func NewMessengerServerHandle(o MessengerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return MessengerServerHandle(ctx, o, opNum, r)
	}
}

func MessengerServerHandle(ctx context.Context, o MessengerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // SendMessage
		op := &xxx_SendMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SendMessage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RecallMessage
		op := &xxx_RecallMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RecallMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RecallMessage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMessenger
type UnimplementedMessengerServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedMessengerServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMessengerServer) RecallMessage(context.Context, *RecallMessageRequest) (*RecallMessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ MessengerServer = (*UnimplementedMessengerServer)(nil)
