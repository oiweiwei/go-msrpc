package extendederror

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

// ExtendedError server interface.
type ExtendedErrorServer interface {
}

func RegisterExtendedErrorServer(conn dcerpc.Conn, o ExtendedErrorServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewExtendedErrorServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ExtendedErrorSyntaxV1_0))...)
}

func NewExtendedErrorServerHandle(o ExtendedErrorServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ExtendedErrorServerHandle(ctx, o, opNum, r)
	}
}

func ExtendedErrorServerHandle(ctx context.Context, o ExtendedErrorServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	}
	return nil, nil
}

// Unimplemented ExtendedError
type UnimplementedExtendedErrorServer struct {
}

var _ ExtendedErrorServer = (*UnimplementedExtendedErrorServer)(nil)
