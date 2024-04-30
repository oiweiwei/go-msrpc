package ieventclass3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ieventclass2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventclass2/v0"
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
	_ = ieventclass2.GoPackage
)

// IEventClass3 server interface.
type EventClass3Server interface {

	// IEventClass2 base class.
	ieventclass2.EventClass2Server

	// EventClassPartitionID operation.
	GetEventClassPartitionID(context.Context, *GetEventClassPartitionIDRequest) (*GetEventClassPartitionIDResponse, error)

	// EventClassPartitionID operation.
	SetEventClassPartitionID(context.Context, *SetEventClassPartitionIDRequest) (*SetEventClassPartitionIDResponse, error)

	// EventClassApplicationID operation.
	GetEventClassApplicationID(context.Context, *GetEventClassApplicationIDRequest) (*GetEventClassApplicationIDResponse, error)

	// EventClassApplicationID operation.
	SetEventClassApplicationID(context.Context, *SetEventClassApplicationIDRequest) (*SetEventClassApplicationIDResponse, error)
}

func RegisterEventClass3Server(conn dcerpc.Conn, o EventClass3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventClass3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventClass3SyntaxV0_0))...)
}

func NewEventClass3ServerHandle(o EventClass3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventClass3ServerHandle(ctx, o, opNum, r)
	}
}

func EventClass3ServerHandle(ctx context.Context, o EventClass3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 29 {
		// IEventClass2 base method.
		return ieventclass2.EventClass2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 29: // EventClassPartitionID
		in := &GetEventClassPartitionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventClassPartitionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // EventClassPartitionID
		in := &SetEventClassPartitionIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEventClassPartitionID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // EventClassApplicationID
		in := &GetEventClassApplicationIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventClassApplicationID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // EventClassApplicationID
		in := &SetEventClassApplicationIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEventClassApplicationID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
