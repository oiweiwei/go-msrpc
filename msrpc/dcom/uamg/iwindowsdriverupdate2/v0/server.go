package iwindowsdriverupdate2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iwindowsdriverupdate "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate/v0"
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
	_ = iwindowsdriverupdate.GoPackage
)

// IWindowsDriverUpdate2 server interface.
type WindowsDriverUpdate2Server interface {

	// IWindowsDriverUpdate base class.
	iwindowsdriverupdate.WindowsDriverUpdateServer

	GetRebootRequired(context.Context, *GetRebootRequiredRequest) (*GetRebootRequiredResponse, error)

	GetIsPresent(context.Context, *GetIsPresentRequest) (*GetIsPresentResponse, error)

	GetCveIDs(context.Context, *GetCveIDsRequest) (*GetCveIDsResponse, error)

	// Opnum64NotUsedOnWire operation.
	// Opnum64NotUsedOnWire
}

func RegisterWindowsDriverUpdate2Server(conn dcerpc.Conn, o WindowsDriverUpdate2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsDriverUpdate2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdate2SyntaxV0_0))...)
}

func NewWindowsDriverUpdate2ServerHandle(o WindowsDriverUpdate2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsDriverUpdate2ServerHandle(ctx, o, opNum, r)
	}
}

func WindowsDriverUpdate2ServerHandle(ctx context.Context, o WindowsDriverUpdate2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 60 {
		// IWindowsDriverUpdate base method.
		return iwindowsdriverupdate.WindowsDriverUpdateServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 60: // RebootRequired
		op := &xxx_GetRebootRequiredOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRebootRequiredRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRebootRequired(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 61: // IsPresent
		op := &xxx_GetIsPresentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsPresentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsPresent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 62: // CveIDs
		op := &xxx_GetCveIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCveIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCveIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 63: // Opnum64NotUsedOnWire
		// Opnum64NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IWindowsDriverUpdate2
type UnimplementedWindowsDriverUpdate2Server struct {
	iwindowsdriverupdate.UnimplementedWindowsDriverUpdateServer
}

func (UnimplementedWindowsDriverUpdate2Server) GetRebootRequired(context.Context, *GetRebootRequiredRequest) (*GetRebootRequiredResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdate2Server) GetIsPresent(context.Context, *GetIsPresentRequest) (*GetIsPresentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdate2Server) GetCveIDs(context.Context, *GetCveIDsRequest) (*GetCveIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WindowsDriverUpdate2Server = (*UnimplementedWindowsDriverUpdate2Server)(nil)
