package iupdate5

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdate4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate4/v0"
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
	_ = iupdate4.GoPackage
)

// IUpdate5 server interface.
type Update5Server interface {

	// IUpdate4 base class.
	iupdate4.Update4Server

	GetAutoSelection(context.Context, *GetAutoSelectionRequest) (*GetAutoSelectionResponse, error)

	GetAutoDownload(context.Context, *GetAutoDownloadRequest) (*GetAutoDownloadResponse, error)
}

func RegisterUpdate5Server(conn dcerpc.Conn, o Update5Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdate5ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Update5SyntaxV0_0))...)
}

func NewUpdate5ServerHandle(o Update5Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Update5ServerHandle(ctx, o, opNum, r)
	}
}

func Update5ServerHandle(ctx context.Context, o Update5Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 58 {
		// IUpdate4 base method.
		return iupdate4.Update4ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 58: // AutoSelection
		op := &xxx_GetAutoSelectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAutoSelectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAutoSelection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // AutoDownload
		op := &xxx_GetAutoDownloadOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAutoDownloadRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAutoDownload(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdate5
type UnimplementedUpdate5Server struct {
	iupdate4.UnimplementedUpdate4Server
}

func (UnimplementedUpdate5Server) GetAutoSelection(context.Context, *GetAutoSelectionRequest) (*GetAutoSelectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdate5Server) GetAutoDownload(context.Context, *GetAutoDownloadRequest) (*GetAutoDownloadResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Update5Server = (*UnimplementedUpdate5Server)(nil)
