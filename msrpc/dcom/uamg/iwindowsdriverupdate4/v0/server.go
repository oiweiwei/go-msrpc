package iwindowsdriverupdate4

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iwindowsdriverupdate3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate3/v0"
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
	_ = iwindowsdriverupdate3.GoPackage
)

// IWindowsDriverUpdate4 server interface.
type WindowsDriverUpdate4Server interface {

	// IWindowsDriverUpdate3 base class.
	iwindowsdriverupdate3.WindowsDriverUpdate3Server

	// The IWindowsDriverUpdate4::WindowsDriverUpdateEntries (opnum 66) method retrieves
	// a collection of driver update entries associated with this driver.
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
	// This method SHOULD return the value of the WindowsDriverUpdateEntries ADM element.
	GetWindowsDriverUpdateEntries(context.Context, *GetWindowsDriverUpdateEntriesRequest) (*GetWindowsDriverUpdateEntriesResponse, error)

	// The IUpdate4::PerUser (opnum 58) method retrieves whether the update is per user.
	//
	// The IWindowsDriverUpdate4::PerUser (opnum 67) method retrieves whether the update
	// is per user.
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
	// This method SHOULD return the value of the PerUser ADM element.
	GetPerUser(context.Context, *GetPerUserRequest) (*GetPerUserResponse, error)
}

func RegisterWindowsDriverUpdate4Server(conn dcerpc.Conn, o WindowsDriverUpdate4Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsDriverUpdate4ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdate4SyntaxV0_0))...)
}

func NewWindowsDriverUpdate4ServerHandle(o WindowsDriverUpdate4Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsDriverUpdate4ServerHandle(ctx, o, opNum, r)
	}
}

func WindowsDriverUpdate4ServerHandle(ctx context.Context, o WindowsDriverUpdate4Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 65 {
		// IWindowsDriverUpdate3 base method.
		return iwindowsdriverupdate3.WindowsDriverUpdate3ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 65: // WindowsDriverUpdateEntries
		op := &xxx_GetWindowsDriverUpdateEntriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetWindowsDriverUpdateEntriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetWindowsDriverUpdateEntries(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 66: // PerUser
		op := &xxx_GetPerUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPerUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPerUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWindowsDriverUpdate4
type UnimplementedWindowsDriverUpdate4Server struct {
	iwindowsdriverupdate3.UnimplementedWindowsDriverUpdate3Server
}

func (UnimplementedWindowsDriverUpdate4Server) GetWindowsDriverUpdateEntries(context.Context, *GetWindowsDriverUpdateEntriesRequest) (*GetWindowsDriverUpdateEntriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdate4Server) GetPerUser(context.Context, *GetPerUserRequest) (*GetPerUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WindowsDriverUpdate4Server = (*UnimplementedWindowsDriverUpdate4Server)(nil)
