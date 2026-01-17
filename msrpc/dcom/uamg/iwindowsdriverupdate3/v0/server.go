package iwindowsdriverupdate3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iwindowsdriverupdate2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate2/v0"
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
	_ = iwindowsdriverupdate2.GoPackage
)

// IWindowsDriverUpdate3 server interface.
type WindowsDriverUpdate3Server interface {

	// IWindowsDriverUpdate2 base class.
	iwindowsdriverupdate2.WindowsDriverUpdate2Server

	// The IUpdate3::BrowseOnly (opnum 57) method retrieves whether the update is browse-only.
	//
	// The IWindowsDriverUpdate3::BrowseOnly (opnum 65) method retrieves whether the update
	// is browse-only. If an update is browse-only, then it is not recommended for the update
	// to be installed automatically.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the BrowseOnly ADM element.
	GetBrowseOnly(context.Context, *GetBrowseOnlyRequest) (*GetBrowseOnlyResponse, error)
}

func RegisterWindowsDriverUpdate3Server(conn dcerpc.Conn, o WindowsDriverUpdate3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsDriverUpdate3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdate3SyntaxV0_0))...)
}

func NewWindowsDriverUpdate3ServerHandle(o WindowsDriverUpdate3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsDriverUpdate3ServerHandle(ctx, o, opNum, r)
	}
}

func WindowsDriverUpdate3ServerHandle(ctx context.Context, o WindowsDriverUpdate3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 64 {
		// IWindowsDriverUpdate2 base method.
		return iwindowsdriverupdate2.WindowsDriverUpdate2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 64: // BrowseOnly
		op := &xxx_GetBrowseOnlyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBrowseOnlyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBrowseOnly(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWindowsDriverUpdate3
type UnimplementedWindowsDriverUpdate3Server struct {
	iwindowsdriverupdate2.UnimplementedWindowsDriverUpdate2Server
}

func (UnimplementedWindowsDriverUpdate3Server) GetBrowseOnly(context.Context, *GetBrowseOnlyRequest) (*GetBrowseOnlyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WindowsDriverUpdate3Server = (*UnimplementedWindowsDriverUpdate3Server)(nil)
