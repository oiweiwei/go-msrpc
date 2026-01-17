package iwbemcontext

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

// IWbemContext server interface.
type ContextServer interface {

	// IUnknown base class.
	iunknown.UnknownServer
}

func RegisterContextServer(conn dcerpc.Conn, o ContextServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewContextServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ContextSyntaxV0_0))...)
}

func NewContextServerHandle(o ContextServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ContextServerHandle(ctx, o, opNum, r)
	}
}

func ContextServerHandle(ctx context.Context, o ContextServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	}
	return nil, nil
}

// Unimplemented IWbemContext
type UnimplementedContextServer struct {
	iunknown.UnimplementedUnknownServer
}

var _ ContextServer = (*UnimplementedContextServer)(nil)
