package imsmqevent3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	imsmqevent2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqevent2/v0"
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
	_ = imsmqevent2.GoPackage
)

// IMSMQEvent3 server interface.
type ImsmqEvent3Server interface {

	// IMSMQEvent2 base class.
	imsmqevent2.ImsmqEvent2Server
}

func RegisterImsmqEvent3Server(conn dcerpc.Conn, o ImsmqEvent3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqEvent3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqEvent3SyntaxV0_0))...)
}

func NewImsmqEvent3ServerHandle(o ImsmqEvent3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqEvent3ServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqEvent3ServerHandle(ctx context.Context, o ImsmqEvent3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	}
	return nil, nil
}

// Unimplemented IMSMQEvent3
type UnimplementedImsmqEvent3Server struct {
	imsmqevent2.UnimplementedImsmqEvent2Server
}

var _ ImsmqEvent3Server = (*UnimplementedImsmqEvent3Server)(nil)
