package emsmdb

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
)

// emsmdb server interface.
type EmsmdbServer interface {

	// Opnum0Reserved operation.
	// Opnum0Reserved

	// The EcDoDisconnect method closes a Session Context with the server. The server destroys
	// the Session Context and releases all associated server state, objects, and resources
	// that are associated with the Session Context. This call requires an active session
	// context handle from the EcDoConnectEx method, as specified in section 3.1.4.1.
	//
	// Return Values: If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	DoDisconnect(context.Context, *DoDisconnectRequest) (*DoDisconnectResponse, error)

	// Opnum2Reserved operation.
	// Opnum2Reserved

	// Opnum3Reserved operation.
	// Opnum3Reserved

	// The EcRRegisterPushNotification method registers a callback address with the server
	// for a Session Context. The server MAY<26> support the EcRRegisterPushNotification
	// method.
	//
	// The callback address is used to notify the client of pending events on the server.
	// This call requires an active session context handle from the EcDoConnectEx method,
	// as specified in section 3.1.4.1. The server MUST store the callback address and the
	// opaque context data in the Session Context. To notify the client of pending events,
	// the server sends a packet containing just the opaque context data to the callback
	// address. The callback address specifies which network transport is to be used to
	// send the data packet.
	//
	// For more information about notification handling, see [MS-OXCNOTIF].
	//
	// Return Values: If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code or one of the protocol-defined
	// error codes listed in the following table.
	//
	//	+-----------------+------------+------------------------------------------------------+
	//	|   ERROR CODE    |            |                                                      |
	//	|      NAME       |   VALUE    |                       MEANING                        |
	//	|                 |            |                                                      |
	//	+-----------------+------------+------------------------------------------------------+
	//	+-----------------+------------+------------------------------------------------------+
	//	| ecInvalidParam  | 0x80070057 | A parameter passed was not valid for the call.       |
	//	+-----------------+------------+------------------------------------------------------+
	//	| ecNotSupported  | 0x80040102 | The callback address is not supported on the server. |
	//	+-----------------+------------+------------------------------------------------------+
	//	| ecTooBig        | 0x80040305 | Opaque context data is too large.                    |
	//	+-----------------+------------+------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	RegisterPushNotification(context.Context, *RegisterPushNotificationRequest) (*RegisterPushNotificationResponse, error)

	// Opnum5Reserved operation.
	// Opnum5Reserved

	// The EcDummyRpc method returns a SUCCESS. A client can use it to determine whether
	// it can communicate with the server.
	//
	// Return Values: The function MUST succeed and return 0.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	DummyRPC(context.Context, *DummyRPCRequest) (*DummyRPCResponse, error)

	// Opnum7Reserved operation.
	// Opnum7Reserved

	// Opnum8Reserved operation.
	// Opnum8Reserved

	// Opnum9Reserved operation.
	// Opnum9Reserved

	// The EcDoConnectEx method establishes a new Session Context with the server. The Session
	// Context is persisted on the server until the client disconnects by using the EcDoDisconnect
	// method, as specified in section 3.1.4.3. The EcDoConnectEx method returns a session
	// context handle to be used by a client in subsequent calls.
	//
	// Return Values: If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code or one of the protocol-defined
	// error codes listed in the following table.
	//
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	|      ERROR CODE       |            |                                                                                  |
	//	|         NAME          |   VALUE    |                                     MEANING                                      |
	//	|                       |            |                                                                                  |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	| ecAccessDenied<14>    | 0x80070005 | The authentication context associated with the binding handle does not have      |
	//	|                       |            | enough privilege or the szUserDN parameter is empty.                             |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	| ecNotEncrypted        | 0x00000970 | The server is configured to require encryption and the authentication for        |
	//	|                       |            | the binding handle contained in the hBinding parameter is not set with           |
	//	|                       |            | RPC_C_AUTHN_LEVEL_PKT_PRIVACY. For more information about setting the            |
	//	|                       |            | authentication and authorization, see [MSDN-RpcBindingSetAuthInfoEx]. The client |
	//	|                       |            | attempts the call again with new binding handle that is encrypted.               |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	| ecClientVerDisallowed | 0x000004DF | 1. The server requires encryption, but the client is not encrypted and the       |
	//	|                       |            | client does not support receiving error code ecNotEncrypted being returned by    |
	//	|                       |            | the server. For details about which client versions do not support receiving     |
	//	|                       |            | error code ecNotEncrypted, see section 3.1.4.1.3 and section 3.2.4.1.3. 2. The   |
	//	|                       |            | client version has been blocked by the administrator.                            |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	| ecLoginFailure        | 0x80040111 | Server is unable to log in user to the mailbox or public folder database.        |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	| ecUnknownUser         | 0x000003EB | The server does not recognize the szUserDN parameter as a valid enabled mailbox. |
	//	|                       |            | For more details, see [MS-OXCSTOR] section 3.1.4.1.                              |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	| ecLoginPerm           | 0x000003F2 | The connection is requested for administrative access, but the authentication    |
	//	|                       |            | context associated with the binding handle does not have enough privilege.       |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	| ecVersionMismatch     | 0x80040110 | The client and server versions are not compatible. The client protocol version   |
	//	|                       |            | is earlier than that required by the server.                                     |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	| ecCachedModeRequired  | 0x000004E1 | The server requires the client to be running in cache mode. For details about    |
	//	|                       |            | which client versions understand this error code, see section 3.2.4.1.3.         |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	| ecRpcHttpDisallowed   | 0x000004E0 | The server requires the client to not be connected via RPC over HTTP. For        |
	//	|                       |            | details about which client versions understand this error code, see section      |
	//	|                       |            | 3.1.4.1.3.                                                                       |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//	| ecProtocolDisabled    | 0x000007D8 | The server disallows the user to access the server via this protocol interface.  |
	//	|                       |            | This could be done if the user is only capable of accessing their mailbox        |
	//	|                       |            | information through a different means (for example, Webmail, POP, or IMAP).      |
	//	|                       |            | For details about which client versions understand this error code, see section  |
	//	|                       |            | 3.1.4.1.3.                                                                       |
	//	+-----------------------+------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	DoConnectEx(context.Context, *DoConnectExRequest) (*DoConnectExResponse, error)

	// The EcDoRpcExt2 method passes generic ROP commands to the server for processing within
	// a Session Context. Each call can contain multiple ROP commands. The server returns
	// the results of each ROP command to the client. This call requires an active session
	// context handle returned from the EcDoConnectEx method.
	//
	// Return Values: If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code or the protocol-defined
	// error code listed in the following table.
	//
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	|   ERROR CODE    |            |                                                                                  |
	//	|      NAME       |   VALUE    |                                     MEANING                                      |
	//	|                 |            |                                                                                  |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| ecRpcFormat     | 0x000004B6 | The format of the request was found to be invalid. This is a generic error that  |
	//	|                 |            | means the length was found to be invalid or the content was found to be invalid. |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	DoRPCExt2(context.Context, *DoRPCExt2Request) (*DoRPCExt2Response, error)

	// Opnum12Reserved operation.
	// Opnum12Reserved

	// Opnum13Reserved operation.
	// Opnum13Reserved

	// The EcDoAsyncConnectEx method binds a session context handle returned from the EcDoConnectEx
	// method, as specified in section 3.1.4.1, to a new asynchronous context handle that
	// can be used in calls to the EcDoAsyncWaitEx method in the AsyncEMSMDB interface,
	// as specified in section 3.3.4.1. This call requires that an active session context
	// handle be returned from the EcDoConnectEx method.
	//
	// This method is part of notification handling. For more information about notifications,
	// see [MS-OXCNOTIF].
	//
	// Return Values: If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code or the protocol-defined
	// error code listed in the following table.
	//
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	|   ERROR CODE    |            |                                                                                  |
	//	|      NAME       |   VALUE    |                                     MEANING                                      |
	//	|                 |            |                                                                                  |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| ecRejected<25>  | 0x000007EE | Server has asynchronous RPC notifications disabled. Client either polls          |
	//	|                 |            | for notifications or calls the EcRRegisterPushNotifications method (section      |
	//	|                 |            | 3.1.4.5).                                                                        |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	DoAsyncConnectEx(context.Context, *DoAsyncConnectExRequest) (*DoAsyncConnectExResponse, error)
}

func RegisterEmsmdbServer(conn dcerpc.Conn, o EmsmdbServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEmsmdbServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EmsmdbSyntaxV0_81))...)
}

func NewEmsmdbServerHandle(o EmsmdbServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EmsmdbServerHandle(ctx, o, opNum, r)
	}
}

func EmsmdbServerHandle(ctx context.Context, o EmsmdbServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0Reserved
		// Opnum0Reserved
		return nil, nil
	case 1: // EcDoDisconnect
		op := &xxx_DoDisconnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DoDisconnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DoDisconnect(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // Opnum2Reserved
		// Opnum2Reserved
		return nil, nil
	case 3: // Opnum3Reserved
		// Opnum3Reserved
		return nil, nil
	case 4: // EcRRegisterPushNotification
		op := &xxx_RegisterPushNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterPushNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterPushNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Opnum5Reserved
		// Opnum5Reserved
		return nil, nil
	case 6: // EcDummyRpc
		op := &xxx_DummyRPCOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DummyRPCRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DummyRPC(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // Opnum7Reserved
		// Opnum7Reserved
		return nil, nil
	case 8: // Opnum8Reserved
		// Opnum8Reserved
		return nil, nil
	case 9: // Opnum9Reserved
		// Opnum9Reserved
		return nil, nil
	case 10: // EcDoConnectEx
		op := &xxx_DoConnectExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DoConnectExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DoConnectEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // EcDoRpcExt2
		op := &xxx_DoRPCExt2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DoRPCExt2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DoRPCExt2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Opnum12Reserved
		// Opnum12Reserved
		return nil, nil
	case 13: // Opnum13Reserved
		// Opnum13Reserved
		return nil, nil
	case 14: // EcDoAsyncConnectEx
		op := &xxx_DoAsyncConnectExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DoAsyncConnectExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DoAsyncConnectEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented emsmdb
type UnimplementedEmsmdbServer struct {
}

func (UnimplementedEmsmdbServer) DoDisconnect(context.Context, *DoDisconnectRequest) (*DoDisconnectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEmsmdbServer) RegisterPushNotification(context.Context, *RegisterPushNotificationRequest) (*RegisterPushNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEmsmdbServer) DummyRPC(context.Context, *DummyRPCRequest) (*DummyRPCResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEmsmdbServer) DoConnectEx(context.Context, *DoConnectExRequest) (*DoConnectExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEmsmdbServer) DoRPCExt2(context.Context, *DoRPCExt2Request) (*DoRPCExt2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEmsmdbServer) DoAsyncConnectEx(context.Context, *DoAsyncConnectExRequest) (*DoAsyncConnectExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EmsmdbServer = (*UnimplementedEmsmdbServer)(nil)
