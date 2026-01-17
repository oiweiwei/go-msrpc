package ifsrmactionemail2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_GetAttachmentFileListSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAttachmentFileListSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAttachmentFileListSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // AttachmentFileListSize
		op := &xxx_SetAttachmentFileListSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAttachmentFileListSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAttachmentFileListSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmActionEmail2
type UnimplementedActionEmail2Server struct {
	ifsrmactionemail.UnimplementedActionEmailServer
}

func (UnimplementedActionEmail2Server) GetAttachmentFileListSize(context.Context, *GetAttachmentFileListSizeRequest) (*GetAttachmentFileListSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionEmail2Server) SetAttachmentFileListSize(context.Context, *SetAttachmentFileListSizeRequest) (*SetAttachmentFileListSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ActionEmail2Server = (*UnimplementedActionEmail2Server)(nil)
