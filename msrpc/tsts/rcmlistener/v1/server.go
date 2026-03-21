package rcmlistener

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

// RCMListener server interface.
type RcmListenerServer interface {

	// The RpcOpenListener method returns a handle to the specified Terminal Services listener
	// running on a terminal server. No special permissions are required to call this method.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	OpenListener(context.Context, *OpenListenerRequest) (*OpenListenerResponse, error)

	// The RpcCloseListener method closes the handle for a Terminal Services listener running
	// on a terminal server. This MUST be called after RpcOpenListener. The call to this
	// method MUST be serialized if there are multiple threads running otherwise the behavior
	// of this function is unknown. No special permissions are required to call this method.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.<159>
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	CloseListener(context.Context, *CloseListenerRequest) (*CloseListenerResponse, error)

	// The RpcStopListener method stops the specified Terminal Services listener running
	// on a terminal server. This MUST be called after RpcOpenListener. The caller MUST
	// have WINSTATION_RESET permission to the listener. The method checks whether the caller
	// has WINSTATION_RESET permission (section 3.1.1) by setting it as the Access Request
	// mask, and fails if the caller does not have the permission.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	StopListener(context.Context, *StopListenerRequest) (*StopListenerResponse, error)

	// The RpcStartListener method starts the specified Terminal Services listener on a
	// terminal server. This MUST be called after RpcOpenListener. The caller MUST have
	// WINSTATION_RESET and WINSTATION_QUERY permissions to the listener. The method checks
	// whether the caller has WINSTATION_RESET and WINSTATION_QUERY permission (section
	// 3.1.1) by setting it as the Access Request mask, and fails if the caller does not
	// have the permissions.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	StartListener(context.Context, *StartListenerRequest) (*StartListenerResponse, error)

	// The RpcIsListening method checks whether the specified Terminal Services listener
	// is running on a terminal server. This MUST be called after RpcOpenListener. The caller
	// MUST have WINSTATION_QUERY permission to the listener. The method checks whether
	// the caller has WINSTATION_QUERY permission (section 3.1.1) by setting it as the Access
	// Request mask, and fails if the caller does not have the permission.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	IsListening(context.Context, *IsListeningRequest) (*IsListeningResponse, error)
}

func RegisterRcmListenerServer(conn dcerpc.Conn, o RcmListenerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRcmListenerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RcmListenerSyntaxV1_0))...)
}

func NewRcmListenerServerHandle(o RcmListenerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RcmListenerServerHandle(ctx, o, opNum, r)
	}
}

func RcmListenerServerHandle(ctx context.Context, o RcmListenerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RpcOpenListener
		op := &xxx_OpenListenerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenListenerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenListener(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RpcCloseListener
		op := &xxx_CloseListenerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseListenerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseListener(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // RpcStopListener
		op := &xxx_StopListenerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StopListenerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StopListener(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // RpcStartListener
		op := &xxx_StartListenerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartListenerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartListener(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RpcIsListening
		op := &xxx_IsListeningOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsListeningRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsListening(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented RCMListener
type UnimplementedRcmListenerServer struct {
}

func (UnimplementedRcmListenerServer) OpenListener(context.Context, *OpenListenerRequest) (*OpenListenerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmListenerServer) CloseListener(context.Context, *CloseListenerRequest) (*CloseListenerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmListenerServer) StopListener(context.Context, *StopListenerRequest) (*StopListenerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmListenerServer) StartListener(context.Context, *StartListenerRequest) (*StartListenerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmListenerServer) IsListening(context.Context, *IsListeningRequest) (*IsListeningResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RcmListenerServer = (*UnimplementedRcmListenerServer)(nil)
