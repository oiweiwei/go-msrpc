package windowsshutdown

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// WindowsShutdown server interface.
type WindowsShutdownServer interface {

	// The WsdrInitiateShutdown method is used to initiate the shutdown of the remote computer.<14>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	InitiateShutdown(context.Context, *InitiateShutdownRequest) (*InitiateShutdownResponse, error)

	// The WsdrAbortShutdown method is used to terminate the shutdown of the remote computer
	// within the waiting period.<15>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	AbortShutdown(context.Context, *AbortShutdownRequest) (*AbortShutdownResponse, error)
}

func RegisterWindowsShutdownServer(conn dcerpc.Conn, o WindowsShutdownServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsShutdownServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsShutdownSyntaxV1_0))...)
}

func NewWindowsShutdownServerHandle(o WindowsShutdownServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsShutdownServerHandle(ctx, o, opNum, r)
	}
}

func WindowsShutdownServerHandle(ctx context.Context, o WindowsShutdownServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // WsdrInitiateShutdown
		op := &xxx_InitiateShutdownOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitiateShutdownRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitiateShutdown(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // WsdrAbortShutdown
		op := &xxx_AbortShutdownOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AbortShutdownRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AbortShutdown(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented WindowsShutdown
type UnimplementedWindowsShutdownServer struct {
}

func (UnimplementedWindowsShutdownServer) InitiateShutdown(context.Context, *InitiateShutdownRequest) (*InitiateShutdownResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsShutdownServer) AbortShutdown(context.Context, *AbortShutdownRequest) (*AbortShutdownResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WindowsShutdownServer = (*UnimplementedWindowsShutdownServer)(nil)
