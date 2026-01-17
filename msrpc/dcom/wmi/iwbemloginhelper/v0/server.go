package iwbemloginhelper

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IWbemLoginHelper server interface.
type LoginHelperServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The IWbemLoginHelper::SetEvent MUST return WBEM_S_NO_ERROR. The SetEvent method SHOULD
	// NOT perform any action.<57>
	//
	// The opnum of the SetEvent method equals 3.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// If the method fails, the server MUST return an HRESULT whose S (severity) bit is
	// set as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	SetEvent(context.Context, *SetEventRequest) (*SetEventResponse, error)
}

func RegisterLoginHelperServer(conn dcerpc.Conn, o LoginHelperServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewLoginHelperServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(LoginHelperSyntaxV0_0))...)
}

func NewLoginHelperServerHandle(o LoginHelperServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return LoginHelperServerHandle(ctx, o, opNum, r)
	}
}

func LoginHelperServerHandle(ctx context.Context, o LoginHelperServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // SetEvent
		op := &xxx_SetEventOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetEventRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetEvent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWbemLoginHelper
type UnimplementedLoginHelperServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedLoginHelperServer) SetEvent(context.Context, *SetEventRequest) (*SetEventResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ LoginHelperServer = (*UnimplementedLoginHelperServer)(nil)
