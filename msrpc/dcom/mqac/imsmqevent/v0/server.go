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
type ImsmqEventServer interface {

	// IDispatch base class.
	idispatch.DispatchServer
}

func RegisterImsmqEventServer(conn dcerpc.Conn, o ImsmqEventServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqEventServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqEventSyntaxV0_0))...)
}

func NewImsmqEventServerHandle(o ImsmqEventServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqEventServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqEventServerHandle(ctx context.Context, o ImsmqEventServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	}
	return nil, nil
}

// Unimplemented IMSMQEvent
type UnimplementedImsmqEventServer struct {
	idispatch.UnimplementedDispatchServer
}

var _ ImsmqEventServer = (*UnimplementedImsmqEventServer)(nil)
