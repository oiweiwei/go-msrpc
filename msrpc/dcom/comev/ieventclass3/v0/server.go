package ieventclass3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_GetEventClassPartitionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventClassPartitionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventClassPartitionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // EventClassPartitionID
		op := &xxx_SetEventClassPartitionIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventClassPartitionIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEventClassPartitionID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // EventClassApplicationID
		op := &xxx_GetEventClassApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventClassApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventClassApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // EventClassApplicationID
		op := &xxx_SetEventClassApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventClassApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEventClassApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IEventClass3
type UnimplementedEventClass3Server struct {
	ieventclass2.UnimplementedEventClass2Server
}

func (UnimplementedEventClass3Server) GetEventClassPartitionID(context.Context, *GetEventClassPartitionIDRequest) (*GetEventClassPartitionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClass3Server) SetEventClassPartitionID(context.Context, *SetEventClassPartitionIDRequest) (*SetEventClassPartitionIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClass3Server) GetEventClassApplicationID(context.Context, *GetEventClassApplicationIDRequest) (*GetEventClassApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEventClass3Server) SetEventClassApplicationID(context.Context, *SetEventClassApplicationIDRequest) (*SetEventClassApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EventClass3Server = (*UnimplementedEventClass3Server)(nil)
