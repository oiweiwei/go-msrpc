package ifsrmactionemail

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

// IFsrmActionEmail server interface.
type ActionEmailServer interface {

	// IFsrmAction base class.
	ifsrmaction.ActionServer

	// MailFrom operation.
	GetMailFrom(context.Context, *GetMailFromRequest) (*GetMailFromResponse, error)

	// MailFrom operation.
	SetMailFrom(context.Context, *SetMailFromRequest) (*SetMailFromResponse, error)

	// MailReplyTo operation.
	GetMailReplyTo(context.Context, *GetMailReplyToRequest) (*GetMailReplyToResponse, error)

	// MailReplyTo operation.
	SetMailReplyTo(context.Context, *SetMailReplyToRequest) (*SetMailReplyToResponse, error)

	// MailTo operation.
	GetMailTo(context.Context, *GetMailToRequest) (*GetMailToResponse, error)

	// MailTo operation.
	SetMailTo(context.Context, *SetMailToRequest) (*SetMailToResponse, error)

	// MailCc operation.
	GetMailCC(context.Context, *GetMailCCRequest) (*GetMailCCResponse, error)

	// MailCc operation.
	SetMailCC(context.Context, *SetMailCCRequest) (*SetMailCCResponse, error)

	// MailBcc operation.
	GetMailBCC(context.Context, *GetMailBCCRequest) (*GetMailBCCResponse, error)

	// MailBcc operation.
	SetMailBCC(context.Context, *SetMailBCCRequest) (*SetMailBCCResponse, error)

	// MailSubject operation.
	GetMailSubject(context.Context, *GetMailSubjectRequest) (*GetMailSubjectResponse, error)

	// MailSubject operation.
	SetMailSubject(context.Context, *SetMailSubjectRequest) (*SetMailSubjectResponse, error)

	// MessageText operation.
	GetMessageText(context.Context, *GetMessageTextRequest) (*GetMessageTextResponse, error)

	// MessageText operation.
	SetMessageText(context.Context, *SetMessageTextRequest) (*SetMessageTextResponse, error)
}

func RegisterActionEmailServer(conn dcerpc.Conn, o ActionEmailServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewActionEmailServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ActionEmailSyntaxV0_0))...)
}

func NewActionEmailServerHandle(o ActionEmailServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ActionEmailServerHandle(ctx, o, opNum, r)
	}
}

func ActionEmailServerHandle(ctx context.Context, o ActionEmailServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmAction base method.
		return ifsrmaction.ActionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // MailFrom
		op := &xxx_GetMailFromOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMailFromRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMailFrom(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // MailFrom
		op := &xxx_SetMailFromOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMailFromRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMailFrom(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // MailReplyTo
		op := &xxx_GetMailReplyToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMailReplyToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMailReplyTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // MailReplyTo
		op := &xxx_SetMailReplyToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMailReplyToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMailReplyTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // MailTo
		op := &xxx_GetMailToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMailToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMailTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // MailTo
		op := &xxx_SetMailToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMailToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMailTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // MailCc
		op := &xxx_GetMailCCOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMailCCRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMailCC(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // MailCc
		op := &xxx_SetMailCCOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMailCCRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMailCC(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // MailBcc
		op := &xxx_GetMailBCCOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMailBCCRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMailBCC(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // MailBcc
		op := &xxx_SetMailBCCOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMailBCCRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMailBCC(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // MailSubject
		op := &xxx_GetMailSubjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMailSubjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMailSubject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // MailSubject
		op := &xxx_SetMailSubjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMailSubjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMailSubject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // MessageText
		op := &xxx_GetMessageTextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMessageTextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMessageText(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // MessageText
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

// Unimplemented IFsrmActionEmail
type UnimplementedActionEmailServer struct {
	ifsrmaction.UnimplementedActionServer
}

func (UnimplementedActionEmailServer) GetMailFrom(context.Context, *GetMailFromRequest) (*GetMailFromResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) SetMailFrom(context.Context, *SetMailFromRequest) (*SetMailFromResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) GetMailReplyTo(context.Context, *GetMailReplyToRequest) (*GetMailReplyToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) SetMailReplyTo(context.Context, *SetMailReplyToRequest) (*SetMailReplyToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) GetMailTo(context.Context, *GetMailToRequest) (*GetMailToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) SetMailTo(context.Context, *SetMailToRequest) (*SetMailToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) GetMailCC(context.Context, *GetMailCCRequest) (*GetMailCCResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) SetMailCC(context.Context, *SetMailCCRequest) (*SetMailCCResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) GetMailBCC(context.Context, *GetMailBCCRequest) (*GetMailBCCResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) SetMailBCC(context.Context, *SetMailBCCRequest) (*SetMailBCCResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) GetMailSubject(context.Context, *GetMailSubjectRequest) (*GetMailSubjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) SetMailSubject(context.Context, *SetMailSubjectRequest) (*SetMailSubjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) GetMessageText(context.Context, *GetMessageTextRequest) (*GetMessageTextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmailServer) SetMessageText(context.Context, *SetMessageTextRequest) (*SetMessageTextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ActionEmailServer = (*UnimplementedActionEmailServer)(nil)
