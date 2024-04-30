package ifsrmactioneventlog

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
		in := &GetEventTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // EventType
		in := &SetEventTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetEventType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // MessageText
		in := &GetMessageTextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMessageText(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // MessageText
		in := &SetMessageTextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMessageText(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
