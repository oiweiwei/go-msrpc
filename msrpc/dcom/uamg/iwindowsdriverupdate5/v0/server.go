package iwindowsdriverupdate5

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iwindowsdriverupdate4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate4/v0"
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
	_ = iwindowsdriverupdate4.GoPackage
)

// IWindowsDriverUpdate5 server interface.
type WindowsDriverUpdate5Server interface {

	// IWindowsDriverUpdate4 base class.
	iwindowsdriverupdate4.WindowsDriverUpdate4Server

	GetAutoSelection(context.Context, *GetAutoSelectionRequest) (*GetAutoSelectionResponse, error)

	GetAutoDownload(context.Context, *GetAutoDownloadRequest) (*GetAutoDownloadResponse, error)
}

func RegisterWindowsDriverUpdate5Server(conn dcerpc.Conn, o WindowsDriverUpdate5Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsDriverUpdate5ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdate5SyntaxV0_0))...)
}

func NewWindowsDriverUpdate5ServerHandle(o WindowsDriverUpdate5Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsDriverUpdate5ServerHandle(ctx, o, opNum, r)
	}
}

func WindowsDriverUpdate5ServerHandle(ctx context.Context, o WindowsDriverUpdate5Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 67 {
		// IWindowsDriverUpdate4 base method.
		return iwindowsdriverupdate4.WindowsDriverUpdate4ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 67: // AutoSelection
		op := &xxx_GetAutoSelectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAutoSelectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAutoSelection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 68: // AutoDownload
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

// Unimplemented IWindowsDriverUpdate5
type UnimplementedWindowsDriverUpdate5Server struct {
	iwindowsdriverupdate4.UnimplementedWindowsDriverUpdate4Server
}

func (UnimplementedWindowsDriverUpdate5Server) GetAutoSelection(context.Context, *GetAutoSelectionRequest) (*GetAutoSelectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdate5Server) GetAutoDownload(context.Context, *GetAutoDownloadRequest) (*GetAutoDownloadResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WindowsDriverUpdate5Server = (*UnimplementedWindowsDriverUpdate5Server)(nil)
