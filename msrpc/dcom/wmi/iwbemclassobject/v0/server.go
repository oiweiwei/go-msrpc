package iwbemclassobject

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

// IWbemClassObject server interface.
type ClassObjectServer interface {

	// IUnknown base class.
	iunknown.UnknownServer
}

func RegisterClassObjectServer(conn dcerpc.Conn, o ClassObjectServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClassObjectServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClassObjectSyntaxV0_0))...)
}

func NewClassObjectServerHandle(o ClassObjectServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClassObjectServerHandle(ctx, o, opNum, r)
	}
}

func ClassObjectServerHandle(ctx context.Context, o ClassObjectServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	}
	return nil, nil
}

// Unimplemented IWbemClassObject
type UnimplementedClassObjectServer struct {
	iunknown.UnimplementedUnknownServer
}

var _ ClassObjectServer = (*UnimplementedClassObjectServer)(nil)
