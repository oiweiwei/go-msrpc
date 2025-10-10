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
type Event3Server interface {

	// IMSMQEvent2 base class.
	imsmqevent2.Event2Server
}

func RegisterEvent3Server(conn dcerpc.Conn, o Event3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEvent3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Event3SyntaxV0_0))...)
}

func NewEvent3ServerHandle(o Event3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Event3ServerHandle(ctx, o, opNum, r)
	}
}

func Event3ServerHandle(ctx context.Context, o Event3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	}
	return nil, nil
}

// Unimplemented IMSMQEvent3
type UnimplementedEvent3Server struct {
	imsmqevent2.UnimplementedEvent2Server
}

var _ Event3Server = (*UnimplementedEvent3Server)(nil)
