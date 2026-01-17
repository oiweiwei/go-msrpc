package iwrmprotocol

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

// IWRMProtocol server interface.
type ProtocolServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The GetSupportedClient method retrieves the level of support for clients on the WSRM
	// server.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IWRMProtocol interface methods are specified in section 3.2.4.8.
	GetSupportedClient(context.Context, *GetSupportedClientRequest) (*GetSupportedClientResponse, error)
}

func RegisterProtocolServer(conn dcerpc.Conn, o ProtocolServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewProtocolServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ProtocolSyntaxV0_0))...)
}

func NewProtocolServerHandle(o ProtocolServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ProtocolServerHandle(ctx, o, opNum, r)
	}
}

func ProtocolServerHandle(ctx context.Context, o ProtocolServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetSupportedClient
		op := &xxx_GetSupportedClientOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSupportedClientRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSupportedClient(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMProtocol
type UnimplementedProtocolServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedProtocolServer) GetSupportedClient(context.Context, *GetSupportedClientRequest) (*GetSupportedClientResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ProtocolServer = (*UnimplementedProtocolServer)(nil)
