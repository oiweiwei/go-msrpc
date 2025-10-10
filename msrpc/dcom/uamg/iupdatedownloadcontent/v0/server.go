package iupdatedownloadcontent

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IUpdateDownloadContent server interface.
type UpdateDownloadContentServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetDownloadURL(context.Context, *GetDownloadURLRequest) (*GetDownloadURLResponse, error)
}

func RegisterUpdateDownloadContentServer(conn dcerpc.Conn, o UpdateDownloadContentServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateDownloadContentServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateDownloadContentSyntaxV0_0))...)
}

func NewUpdateDownloadContentServerHandle(o UpdateDownloadContentServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateDownloadContentServerHandle(ctx, o, opNum, r)
	}
}

func UpdateDownloadContentServerHandle(ctx context.Context, o UpdateDownloadContentServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // DownloadUrl
		op := &xxx_GetDownloadURLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDownloadURLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDownloadURL(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateDownloadContent
type UnimplementedUpdateDownloadContentServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateDownloadContentServer) GetDownloadURL(context.Context, *GetDownloadURLRequest) (*GetDownloadURLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateDownloadContentServer = (*UnimplementedUpdateDownloadContentServer)(nil)
