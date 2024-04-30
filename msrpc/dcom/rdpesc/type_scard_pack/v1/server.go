package type_scard_pack

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

// type_scard_pack server interface.
type TypeSmartCardPackServer interface {
}

func RegisterTypeSmartCardPackServer(conn dcerpc.Conn, o TypeSmartCardPackServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTypeSmartCardPackServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TypeSmartCardPackSyntaxV1_0))...)
}

func NewTypeSmartCardPackServerHandle(o TypeSmartCardPackServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TypeSmartCardPackServerHandle(ctx, o, opNum, r)
	}
}

func TypeSmartCardPackServerHandle(ctx context.Context, o TypeSmartCardPackServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	}
	return nil, nil
}
