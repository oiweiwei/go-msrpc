package rcmpublic

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

// RCMPublic server interface.
type RcmPublicServer interface {

	// The RpcGetClientData method returns information about the client that is connected
	// to a particular session running on a terminal server. The caller must have WINSTATION_QUERY
	// permission. The method checks whether the caller has WINSTATION_QUERY permission
	// (section 3.1.1) by setting it as the Access Request mask, and fails if the caller
	// does not have the permission.
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
	GetClientData(context.Context, *GetClientDataRequest) (*GetClientDataResponse, error)

	// The RpcGetConfigData method returns the configuration data associated with the user
	// connected to a particular session running on a terminal server. The caller MUST have
	// WINSTATION_QUERY permission. The method checks whether the caller has WINSTATION_QUERY
	// permission (section 3.1.1) by setting it as the Access Request mask, and fails if
	// the caller does not have the permission.<158>
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
	GetConfigData(context.Context, *GetConfigDataRequest) (*GetConfigDataResponse, error)

	// The RpcGetProtocolStatus method retrieves information about the status of the protocol
	// used to connect to a particular session running on a terminal server. The caller
	// MUST have WINSTATION_QUERY permission for the session. The method checks whether
	// the caller has WINSTATION_QUERY permission (section 3.1.1) by setting it as the Access
	// Request mask, and fails if the caller does not have the permission.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Successful completion |
	//	+-------------------+-----------------------+
	GetProtocolStatus(context.Context, *GetProtocolStatusRequest) (*GetProtocolStatusResponse, error)

	// The RpcGetLastInputTime method returns the time the last user input was received
	// by the associated protocol for the specified sessions running on a terminal server.
	// The caller MUST have WINSTATION_QUERY permission for the session. The method checks
	// whether the caller has WINSTATION_QUERY permission (section 3.1.1) by setting it
	// as the Access Request mask, and fails if the caller does not have the permission.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Successful completion |
	//	+-------------------+-----------------------+
	GetLastInputTime(context.Context, *GetLastInputTimeRequest) (*GetLastInputTimeResponse, error)

	// The RpcGetRemoteAddress method retrieves the IP address of the client computer connected
	// to the session on the terminal server. The caller MUST have WINSTATION_QUERY permission
	// for the session. The method checks whether the caller has WINSTATION_QUERY permission
	// (section 3.1.1) by setting it as the Access Request mask, and fails if the caller
	// does not have the permission.
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
	GetRemoteAddress(context.Context, *GetRemoteAddressRequest) (*GetRemoteAddressResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// The RpcGetAllListeners method returns a list of all Terminal Services listeners running
	// on a terminal server. No special permissions are required to call this method. However,
	// only listeners for which the caller has WINSTATION_QUERY permission are enumerated.
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
	GetAllListeners(context.Context, *GetAllListenersRequest) (*GetAllListenersResponse, error)

	// The RpcGetSessionProtocolLastInputTime method returns the protocol status and the
	// time the last input was received by the protocol associated with a specific session
	// running on a terminal server. The caller MUST have WINSTATION_QUERY permission for
	// the session. The method checks whether the caller has WINSTATION_QUERY permission
	// (section 3.1.1) by setting it as the Access Request mask, and fails if the caller
	// does not have the permission.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Successful completion |
	//	+-------------------+-----------------------+
	GetSessionProtocolLastInputTime(context.Context, *GetSessionProtocolLastInputTimeRequest) (*GetSessionProtocolLastInputTimeResponse, error)

	// The RpcGetUserCertificates method returns a client X509 certificate if the client
	// used the certificate to connect to a session running on a terminal server. For more
	// information, see [X509]. The caller MUST have WINSTATION_QUERY permission for the
	// session. The method checks whether the caller has WINSTATION_QUERY permission (section
	// 3.1.1) by setting it as the Access Request mask, and fails if the caller does not
	// have the permission.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	GetUserCertificates(context.Context, *GetUserCertificatesRequest) (*GetUserCertificatesResponse, error)

	// The RpcQuerySessionData method returns information about a particular session running
	// on a terminal server. The caller MUST have WINSTATION_QUERY permission to the session
	// being queried. The method checks whether the caller has WINSTATION_QUERY permission
	// (section 3.1.1) by setting it as the Access Request mask, and fails if the caller
	// does not have the permission.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Successful completion |
	//	+-------------------+-----------------------+
	QuerySessionData(context.Context, *QuerySessionDataRequest) (*QuerySessionDataResponse, error)
}

func RegisterRcmPublicServer(conn dcerpc.Conn, o RcmPublicServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRcmPublicServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RcmPublicSyntaxV1_0))...)
}

func NewRcmPublicServerHandle(o RcmPublicServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RcmPublicServerHandle(ctx, o, opNum, r)
	}
}

func RcmPublicServerHandle(ctx context.Context, o RcmPublicServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RpcGetClientData
		op := &xxx_GetClientDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RpcGetConfigData
		op := &xxx_GetConfigDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfigData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // RpcGetProtocolStatus
		op := &xxx_GetProtocolStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetProtocolStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProtocolStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // RpcGetLastInputTime
		op := &xxx_GetLastInputTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastInputTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastInputTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RpcGetRemoteAddress
		op := &xxx_GetRemoteAddressOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRemoteAddressRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRemoteAddress(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // RpcGetAllListeners
		op := &xxx_GetAllListenersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllListenersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllListeners(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RpcGetSessionProtocolLastInputTime
		op := &xxx_GetSessionProtocolLastInputTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSessionProtocolLastInputTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSessionProtocolLastInputTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RpcGetUserCertificates
		op := &xxx_GetUserCertificatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserCertificatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserCertificates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RpcQuerySessionData
		op := &xxx_QuerySessionDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QuerySessionDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QuerySessionData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented RCMPublic
type UnimplementedRcmPublicServer struct {
}

func (UnimplementedRcmPublicServer) GetClientData(context.Context, *GetClientDataRequest) (*GetClientDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmPublicServer) GetConfigData(context.Context, *GetConfigDataRequest) (*GetConfigDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmPublicServer) GetProtocolStatus(context.Context, *GetProtocolStatusRequest) (*GetProtocolStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmPublicServer) GetLastInputTime(context.Context, *GetLastInputTimeRequest) (*GetLastInputTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmPublicServer) GetRemoteAddress(context.Context, *GetRemoteAddressRequest) (*GetRemoteAddressResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmPublicServer) GetAllListeners(context.Context, *GetAllListenersRequest) (*GetAllListenersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmPublicServer) GetSessionProtocolLastInputTime(context.Context, *GetSessionProtocolLastInputTimeRequest) (*GetSessionProtocolLastInputTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmPublicServer) GetUserCertificates(context.Context, *GetUserCertificatesRequest) (*GetUserCertificatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRcmPublicServer) QuerySessionData(context.Context, *QuerySessionDataRequest) (*QuerySessionDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RcmPublicServer = (*UnimplementedRcmPublicServer)(nil)
