package iunknown

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

// IUnknown server interface.
type UnknownServer interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire
}

func RegisterUnknownServer(conn dcerpc.Conn, o UnknownServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUnknownServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UnknownSyntaxV0_0))...)
}

func NewUnknownServerHandle(o UnknownServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UnknownServerHandle(ctx, o, opNum, r)
	}
}

func UnknownServerHandle(ctx context.Context, o UnknownServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0NotUsedOnWire
		// Opnum0NotUsedOnWire
		return nil, nil
	case 1: // Opnum1NotUsedOnWire
		// Opnum1NotUsedOnWire
		return nil, nil
	case 2: // Opnum2NotUsedOnWire
		// Opnum2NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
