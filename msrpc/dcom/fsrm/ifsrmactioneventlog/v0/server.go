package ifsrmactioneventlog

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmaction/v0"
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
	_ = ifsrmaction.GoPackage
)

// IFsrmActionEventLog server interface.
type ActionEventLogServer interface {

	// IFsrmAction base class.
	ifsrmaction.ActionServer

	// EventType operation.
	GetEventType(context.Context, *GetEventTypeRequest) (*GetEventTypeResponse, error)

	// EventType operation.
	SetEventType(context.Context, *SetEventTypeRequest) (*SetEventTypeResponse, error)

	// MessageText operation.
	GetMessageText(context.Context, *GetMessageTextRequest) (*GetMessageTextResponse, error)

	// MessageText operation.
	SetMessageText(context.Context, *SetMessageTextRequest) (*SetMessageTextResponse, error)
}

func RegisterActionEventLogServer(conn dcerpc.Conn, o ActionEventLogServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewActionEventLogServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ActionEventLogSyntaxV0_0))...)
}

func NewActionEventLogServerHandle(o ActionEventLogServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ActionEventLogServerHandle(ctx, o, opNum, r)
	}
}

func ActionEventLogServerHandle(ctx context.Context, o ActionEventLogServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmAction base method.
		return ifsrmaction.ActionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // EventType
		op := &xxx_GetEventTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetEventTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetEventType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // EventType
		op := &xxx_SetEventTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEventType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // MessageText
		op := &xxx_GetMessageTextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMessageTextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMessageText(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // MessageText
		op := &xxx_SetMessageTextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMessageTextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMessageText(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmActionEventLog
type UnimplementedActionEventLogServer struct {
	ifsrmaction.UnimplementedActionServer
}

func (UnimplementedActionEventLogServer) GetEventType(context.Context, *GetEventTypeRequest) (*GetEventTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEventLogServer) SetEventType(context.Context, *SetEventTypeRequest) (*SetEventTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEventLogServer) GetMessageText(context.Context, *GetMessageTextRequest) (*GetMessageTextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEventLogServer) SetMessageText(context.Context, *SetMessageTextRequest) (*SetMessageTextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ActionEventLogServer = (*UnimplementedActionEventLogServer)(nil)
