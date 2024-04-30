package ifsrmactionemail2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmactionemail "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmactionemail/v0"
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
	_ = ifsrmactionemail.GoPackage
)

// IFsrmActionEmail2 server interface.
type ActionEmail2Server interface {

	// IFsrmActionEmail base class.
	ifsrmactionemail.ActionEmailServer

	// AttachmentFileListSize operation.
	GetAttachmentFileListSize(context.Context, *GetAttachmentFileListSizeRequest) (*GetAttachmentFileListSizeResponse, error)

	// AttachmentFileListSize operation.
	SetAttachmentFileListSize(context.Context, *SetAttachmentFileListSizeRequest) (*SetAttachmentFileListSizeResponse, error)
}

func RegisterActionEmail2Server(conn dcerpc.Conn, o ActionEmail2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewActionEmail2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ActionEmail2SyntaxV0_0))...)
}

func NewActionEmail2ServerHandle(o ActionEmail2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ActionEmail2ServerHandle(ctx, o, opNum, r)
	}
}

func ActionEmail2ServerHandle(ctx context.Context, o ActionEmail2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 26 {
		// IFsrmActionEmail base method.
		return ifsrmactionemail.ActionEmailServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 26: // AttachmentFileListSize
		in := &GetAttachmentFileListSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAttachmentFileListSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // AttachmentFileListSize
		in := &SetAttachmentFileListSizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAttachmentFileListSize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
