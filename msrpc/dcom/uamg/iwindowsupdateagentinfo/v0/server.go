package iwindowsupdateagentinfo

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IWindowsUpdateAgentInfo server interface.
type WindowsUpdateAgentInfoServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The IWindowsUpdateAgentInfo::GetInfo (opnum 8) method retrieves version information
	// for the server side of the protocol and the update agent.
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
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
}

func RegisterWindowsUpdateAgentInfoServer(conn dcerpc.Conn, o WindowsUpdateAgentInfoServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsUpdateAgentInfoServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsUpdateAgentInfoSyntaxV0_0))...)
}

func NewWindowsUpdateAgentInfoServerHandle(o WindowsUpdateAgentInfoServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsUpdateAgentInfoServerHandle(ctx, o, opNum, r)
	}
}

func WindowsUpdateAgentInfoServerHandle(ctx context.Context, o WindowsUpdateAgentInfoServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetInfo
		op := &xxx_GetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWindowsUpdateAgentInfo
type UnimplementedWindowsUpdateAgentInfoServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedWindowsUpdateAgentInfoServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WindowsUpdateAgentInfoServer = (*UnimplementedWindowsUpdateAgentInfoServer)(nil)
