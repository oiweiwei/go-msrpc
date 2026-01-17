package iwbemloginclientid

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

// IWbemLoginClientID server interface.
type LoginClientIDServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The IWbemLoginClientID::SetClientInfo method MUST pass the client NETBIOS name and
	// a unique client-generated number to the server.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// In case of failure, the server MUST return an HRESULT whose S (severity) bit is set
	// as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	SetClientInfo(context.Context, *SetClientInfoRequest) (*SetClientInfoResponse, error)
}

func RegisterLoginClientIDServer(conn dcerpc.Conn, o LoginClientIDServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewLoginClientIDServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(LoginClientIDSyntaxV0_0))...)
}

func NewLoginClientIDServerHandle(o LoginClientIDServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return LoginClientIDServerHandle(ctx, o, opNum, r)
	}
}

func LoginClientIDServerHandle(ctx context.Context, o LoginClientIDServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // SetClientInfo
		op := &xxx_SetClientInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWbemLoginClientID
type UnimplementedLoginClientIDServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedLoginClientIDServer) SetClientInfo(context.Context, *SetClientInfoRequest) (*SetClientInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ LoginClientIDServer = (*UnimplementedLoginClientIDServer)(nil)
