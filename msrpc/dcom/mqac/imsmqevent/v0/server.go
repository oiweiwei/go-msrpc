package imsmqevent

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

// IMSMQEvent server interface.
type EventServer interface {

	// IDispatch base class.
	idispatch.DispatchServer
}

func RegisterEventServer(conn dcerpc.Conn, o EventServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventSyntaxV0_0))...)
}

func NewEventServerHandle(o EventServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventServerHandle(ctx, o, opNum, r)
	}
}

func EventServerHandle(ctx context.Context, o EventServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	}
	return nil, nil
}

// Unimplemented IMSMQEvent
type UnimplementedEventServer struct {
	idispatch.UnimplementedDispatchServer
}

var _ EventServer = (*UnimplementedEventServer)(nil)
