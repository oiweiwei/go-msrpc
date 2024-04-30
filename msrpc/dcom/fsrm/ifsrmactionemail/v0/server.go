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
		in := &GetMailFromRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMailFrom(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // MailFrom
		in := &SetMailFromRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMailFrom(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // MailReplyTo
		in := &GetMailReplyToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMailReplyTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // MailReplyTo
		in := &SetMailReplyToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMailReplyTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // MailTo
		in := &GetMailToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMailTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // MailTo
		in := &SetMailToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMailTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // MailCc
		in := &GetMailCCRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMailCC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // MailCc
		in := &SetMailCCRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMailCC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // MailBcc
		in := &GetMailBCCRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMailBCC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // MailBcc
		in := &SetMailBCCRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMailBCC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // MailSubject
		in := &GetMailSubjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMailSubject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // MailSubject
		in := &SetMailSubjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMailSubject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // MessageText
		in := &GetMessageTextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMessageText(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // MessageText
		in := &SetMessageTextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMessageText(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
