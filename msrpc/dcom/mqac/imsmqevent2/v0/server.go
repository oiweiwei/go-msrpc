package imsmqevent2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	imsmqevent "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqevent/v0"
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
	_ = imsmqevent.GoPackage
)

// IMSMQEvent2 server interface.
type Event2Server interface {

	// IMSMQEvent base class.
	imsmqevent.EventServer

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterEvent2Server(conn dcerpc.Conn, o Event2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEvent2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Event2SyntaxV0_0))...)
}

func NewEvent2ServerHandle(o Event2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Event2ServerHandle(ctx, o, opNum, r)
	}
}

func Event2ServerHandle(ctx context.Context, o Event2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQEvent2
type UnimplementedEvent2Server struct {
	imsmqevent.UnimplementedEventServer
}

func (UnimplementedEvent2Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Event2Server = (*UnimplementedEvent2Server)(nil)
