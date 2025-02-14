package claims

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

// Claims server interface.
type ClaimsServer interface {
}

func RegisterClaimsServer(conn dcerpc.Conn, o ClaimsServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClaimsServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClaimsSyntaxV1_0))...)
}

func NewClaimsServerHandle(o ClaimsServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClaimsServerHandle(ctx, o, opNum, r)
	}
}

func ClaimsServerHandle(ctx context.Context, o ClaimsServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	}
	return nil, nil
}

// Unimplemented Claims
type UnimplementedClaimsServer struct {
}

var _ ClaimsServer = (*UnimplementedClaimsServer)(nil)
