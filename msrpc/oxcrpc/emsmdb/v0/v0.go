package emsmdb

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	oxcrpc "github.com/oiweiwei/go-msrpc/msrpc/oxcrpc"
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
	_ = oxcrpc.GoPackage
)

var (
	// import guard
	GoPackage = "oxcrpc"
)

var (
	// Syntax UUID
	EmsmdbSyntaxUUID = &uuid.UUID{TimeLow: 0xa4f1db00, TimeMid: 0xca47, TimeHiAndVersion: 0x1067, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x1f, Node: [6]uint8{0x0, 0xdd, 0x1, 0x6, 0x62, 0xda}}
	// Syntax ID
	EmsmdbSyntaxV0_81 = &dcerpc.SyntaxID{IfUUID: EmsmdbSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 81}
)

// emsmdb interface.
type EmsmdbClient interface {

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
	DoDisconnect(context.Context, *DoDisconnectRequest, ...dcerpc.CallOption) (*DoDisconnectResponse, error)

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
	RegisterPushNotification(context.Context, *RegisterPushNotificationRequest, ...dcerpc.CallOption) (*RegisterPushNotificationResponse, error)

	// Opnum5Reserved operation.
	// Opnum5Reserved

	// The EcDummyRpc method returns a SUCCESS. A client can use it to determine whether
	// it can communicate with the server.
	//
	// Return Values: The function MUST succeed and return 0.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	DummyRPC(context.Context, *DummyRPCRequest, ...dcerpc.CallOption) (*DummyRPCResponse, error)

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
	DoConnectEx(context.Context, *DoConnectExRequest, ...dcerpc.CallOption) (*DoConnectExResponse, error)

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
	DoRPCExt2(context.Context, *DoRPCExt2Request, ...dcerpc.CallOption) (*DoRPCExt2Response, error)

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
	DoAsyncConnectEx(context.Context, *DoAsyncConnectExRequest, ...dcerpc.CallOption) (*DoAsyncConnectExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultEmsmdbClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultEmsmdbClient) DoDisconnect(ctx context.Context, in *DoDisconnectRequest, opts ...dcerpc.CallOption) (*DoDisconnectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DoDisconnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEmsmdbClient) RegisterPushNotification(ctx context.Context, in *RegisterPushNotificationRequest, opts ...dcerpc.CallOption) (*RegisterPushNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterPushNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEmsmdbClient) DummyRPC(ctx context.Context, in *DummyRPCRequest, opts ...dcerpc.CallOption) (*DummyRPCResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DummyRPCResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEmsmdbClient) DoConnectEx(ctx context.Context, in *DoConnectExRequest, opts ...dcerpc.CallOption) (*DoConnectExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DoConnectExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEmsmdbClient) DoRPCExt2(ctx context.Context, in *DoRPCExt2Request, opts ...dcerpc.CallOption) (*DoRPCExt2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DoRPCExt2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEmsmdbClient) DoAsyncConnectEx(ctx context.Context, in *DoAsyncConnectExRequest, opts ...dcerpc.CallOption) (*DoAsyncConnectExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DoAsyncConnectExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEmsmdbClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEmsmdbClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewEmsmdbClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EmsmdbClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EmsmdbSyntaxV0_81))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultEmsmdbClient{cc: cc}, nil
}

// xxx_DoDisconnectOperation structure represents the EcDoDisconnect operation
type xxx_DoDisconnectOperation struct {
	Session *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_DoDisconnectOperation) OpNum() int { return 1 }

func (o *xxx_DoDisconnectOperation) OpName() string { return "/emsmdb/v0.81/EcDoDisconnect" }

func (o *xxx_DoDisconnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoDisconnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session != nil {
			if err := o.Session.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oxcrpc.Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DoDisconnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session == nil {
			o.Session = &oxcrpc.Session{}
		}
		if err := o.Session.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoDisconnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoDisconnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session != nil {
			if err := o.Session.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oxcrpc.Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoDisconnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session == nil {
			o.Session = &oxcrpc.Session{}
		}
		if err := o.Session.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DoDisconnectRequest structure represents the EcDoDisconnect operation request
type DoDisconnectRequest struct {
	// pcxh: A session context handle. On input, this parameter is the session context handle
	// of the Session Context that the client is disconnecting. On output, the server MUST
	// set the pcxh parameter to a zero value. Setting the value to zero instructs the RPC
	// layer of the server to destroy the RPC context handle.
	Session *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
}

func (o *DoDisconnectRequest) xxx_ToOp(ctx context.Context, op *xxx_DoDisconnectOperation) *xxx_DoDisconnectOperation {
	if op == nil {
		op = &xxx_DoDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Session = o.Session
	return op
}

func (o *DoDisconnectRequest) xxx_FromOp(ctx context.Context, op *xxx_DoDisconnectOperation) {
	if o == nil {
		return
	}
	o.Session = op.Session
}
func (o *DoDisconnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DoDisconnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoDisconnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DoDisconnectResponse structure represents the EcDoDisconnect operation response
type DoDisconnectResponse struct {
	// pcxh: A session context handle. On input, this parameter is the session context handle
	// of the Session Context that the client is disconnecting. On output, the server MUST
	// set the pcxh parameter to a zero value. Setting the value to zero instructs the RPC
	// layer of the server to destroy the RPC context handle.
	Session *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
	// Return: The EcDoDisconnect return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DoDisconnectResponse) xxx_ToOp(ctx context.Context, op *xxx_DoDisconnectOperation) *xxx_DoDisconnectOperation {
	if op == nil {
		op = &xxx_DoDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Session = o.Session
	op.Return = o.Return
	return op
}

func (o *DoDisconnectResponse) xxx_FromOp(ctx context.Context, op *xxx_DoDisconnectOperation) {
	if o == nil {
		return
	}
	o.Session = op.Session
	o.Return = op.Return
}
func (o *DoDisconnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DoDisconnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoDisconnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterPushNotificationOperation structure represents the EcRRegisterPushNotification operation
type xxx_RegisterPushNotificationOperation struct {
	Session               *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
	RPC                   uint32          `idl:"name:iRpc" json:"rpc"`
	Context               []byte          `idl:"name:rgbContext;size_is:(cbContext)" json:"context"`
	ContextLength         uint16          `idl:"name:cbContext" json:"context_length"`
	AdviseBits            uint32          `idl:"name:grbitAdviseBits" json:"advise_bits"`
	CallbackAddress       []byte          `idl:"name:rgbCallbackAddress;size_is:(cbCallbackAddress)" json:"callback_address"`
	CallbackAddressLength uint16          `idl:"name:cbCallbackAddress" json:"callback_address_length"`
	Notification          uint32          `idl:"name:hNotification" json:"notification"`
	Return                int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterPushNotificationOperation) OpNum() int { return 4 }

func (o *xxx_RegisterPushNotificationOperation) OpName() string {
	return "/emsmdb/v0.81/EcRRegisterPushNotification"
}

func (o *xxx_RegisterPushNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Context != nil && o.ContextLength == 0 {
		o.ContextLength = uint16(len(o.Context))
	}
	if o.CallbackAddress != nil && o.CallbackAddressLength == 0 {
		o.CallbackAddressLength = uint16(len(o.CallbackAddress))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterPushNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session != nil {
			if err := o.Session.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oxcrpc.Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// iRpc {in} (1:(uint32))
	{
		if err := w.WriteData(o.RPC); err != nil {
			return err
		}
	}
	// rgbContext {in} (1:[dim:0,size_is=cbContext](uchar))
	{
		dimSize1 := uint64(o.ContextLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Context {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Context[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Context); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbContext {in} (1:(uint16))
	{
		if err := w.WriteData(o.ContextLength); err != nil {
			return err
		}
	}
	// grbitAdviseBits {in} (1:(uint32))
	{
		if err := w.WriteData(o.AdviseBits); err != nil {
			return err
		}
	}
	// rgbCallbackAddress {in} (1:[dim:0,size_is=cbCallbackAddress](uchar))
	{
		dimSize1 := uint64(o.CallbackAddressLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.CallbackAddress {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.CallbackAddress[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.CallbackAddress); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbCallbackAddress {in} (1:(uint16))
	{
		if err := w.WriteData(o.CallbackAddressLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterPushNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session == nil {
			o.Session = &oxcrpc.Session{}
		}
		if err := o.Session.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// iRpc {in} (1:(uint32))
	{
		if err := w.ReadData(&o.RPC); err != nil {
			return err
		}
	}
	// rgbContext {in} (1:[dim:0,size_is=cbContext](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Context", sizeInfo[0])
		}
		o.Context = make([]byte, sizeInfo[0])
		for i1 := range o.Context {
			i1 := i1
			if err := w.ReadData(&o.Context[i1]); err != nil {
				return err
			}
		}
	}
	// cbContext {in} (1:(uint16))
	{
		if err := w.ReadData(&o.ContextLength); err != nil {
			return err
		}
	}
	// grbitAdviseBits {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AdviseBits); err != nil {
			return err
		}
	}
	// rgbCallbackAddress {in} (1:[dim:0,size_is=cbCallbackAddress](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.CallbackAddress", sizeInfo[0])
		}
		o.CallbackAddress = make([]byte, sizeInfo[0])
		for i1 := range o.CallbackAddress {
			i1 := i1
			if err := w.ReadData(&o.CallbackAddress[i1]); err != nil {
				return err
			}
		}
	}
	// cbCallbackAddress {in} (1:(uint16))
	{
		if err := w.ReadData(&o.CallbackAddressLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterPushNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterPushNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session != nil {
			if err := o.Session.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oxcrpc.Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNotification {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Notification); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterPushNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session == nil {
			o.Session = &oxcrpc.Session{}
		}
		if err := o.Session.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNotification {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Notification); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RegisterPushNotificationRequest structure represents the EcRRegisterPushNotification operation request
type RegisterPushNotificationRequest struct {
	// pcxh: A session context handle. On input, the client MUST pass a valid session context
	// handle that was created by calling the EcDoConnectEx method. The server uses the
	// session context handle to identify the Session Context to use for this call. On output,
	// the server MUST return the same session context handle on success.
	//
	// The server can destroy the session context handle by returning a zero for the pcxh
	// parameter. Reasons for destroying the session context handle are implementation-dependent;
	// following are examples of why the server might destroy the session context handle:
	//
	// * The session context handle that was passed in is not valid.
	//
	// * An attempt was made to access a mailbox ( c5869b4a-c0d3-4ceb-80f4-c47a250ddf3d#gt_d3ad0e15-adc9-4174-bacf-d929b57278b3
	// ) that is in the process of being moved.
	Session *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
	// iRpc: The server MUST ignore this value. The client MUST pass a value of 0x00000000.
	RPC uint32 `idl:"name:iRpc" json:"rpc"`
	// rgbContext: Opaque client-generated context data that is sent back to the client
	// at the callback address, passed in the rgbCallbackAddress parameter, when the server
	// notifies the client of pending event information. The server MUST save this data
	// within the Session Context and use it when sending a notification to the client.
	Context []byte `idl:"name:rgbContext;size_is:(cbContext)" json:"context"`
	// cbContext: The size of the opaque client context data that is passed in the rgbContext
	// parameter. If the value of this parameter is larger than 0x00000010, the server MUST
	// fail this call with error code ecTooBig.
	ContextLength uint16 `idl:"name:cbContext" json:"context_length"`
	// grbitAdviseBits: This parameter MUST be set to 0xFFFFFFFF.
	AdviseBits uint32 `idl:"name:grbitAdviseBits" json:"advise_bits"`
	// rgbCallbackAddress: The callback address for the server to use to notify the client
	// of a pending event. The size of this data is in the cbCallbackAddress parameter.
	//
	// The data contained in this parameter follows the format of a sockaddr structure.
	// For information about the sockaddr structure, see [MSDN-SOCKADDR].
	//
	// The server supports the address families AF_INET and AF_INET6 for a callback address
	// that corresponds to the protocol sequence types that are specified in section 2.1.
	//
	// If an address family is requested that is not supported, the server MUST return error
	// code ecInvalidParam. If the address family is supported but the communications stack
	// of the server does not support the address type, the server MUST return error code
	// ecNotSupported.
	CallbackAddress []byte `idl:"name:rgbCallbackAddress;size_is:(cbCallbackAddress)" json:"callback_address"`
	// cbCallbackAddress: The length of the callback address in the rgbCallbackAddress parameter.
	// The size of this parameter depends on the address family being used. If this size
	// does not correspond to the size of the sockaddr structure based on address family,
	// the server MUST return error code ecInvalidParam.
	CallbackAddressLength uint16 `idl:"name:cbCallbackAddress" json:"callback_address_length"`
}

func (o *RegisterPushNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterPushNotificationOperation) *xxx_RegisterPushNotificationOperation {
	if op == nil {
		op = &xxx_RegisterPushNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Session = o.Session
	op.RPC = o.RPC
	op.Context = o.Context
	op.ContextLength = o.ContextLength
	op.AdviseBits = o.AdviseBits
	op.CallbackAddress = o.CallbackAddress
	op.CallbackAddressLength = o.CallbackAddressLength
	return op
}

func (o *RegisterPushNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterPushNotificationOperation) {
	if o == nil {
		return
	}
	o.Session = op.Session
	o.RPC = op.RPC
	o.Context = op.Context
	o.ContextLength = op.ContextLength
	o.AdviseBits = op.AdviseBits
	o.CallbackAddress = op.CallbackAddress
	o.CallbackAddressLength = op.CallbackAddressLength
}
func (o *RegisterPushNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterPushNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterPushNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterPushNotificationResponse structure represents the EcRRegisterPushNotification operation response
type RegisterPushNotificationResponse struct {
	// pcxh: A session context handle. On input, the client MUST pass a valid session context
	// handle that was created by calling the EcDoConnectEx method. The server uses the
	// session context handle to identify the Session Context to use for this call. On output,
	// the server MUST return the same session context handle on success.
	//
	// The server can destroy the session context handle by returning a zero for the pcxh
	// parameter. Reasons for destroying the session context handle are implementation-dependent;
	// following are examples of why the server might destroy the session context handle:
	//
	// * The session context handle that was passed in is not valid.
	//
	// * An attempt was made to access a mailbox ( c5869b4a-c0d3-4ceb-80f4-c47a250ddf3d#gt_d3ad0e15-adc9-4174-bacf-d929b57278b3
	// ) that is in the process of being moved.
	Session *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
	// hNotification: A handle to the notification callback. If the call completes successfully,
	// the hNotification parameter contains a handle to the notification callback on the
	// server.
	Notification uint32 `idl:"name:hNotification" json:"notification"`
	// Return: The EcRRegisterPushNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterPushNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterPushNotificationOperation) *xxx_RegisterPushNotificationOperation {
	if op == nil {
		op = &xxx_RegisterPushNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Session = o.Session
	op.Notification = o.Notification
	op.Return = o.Return
	return op
}

func (o *RegisterPushNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterPushNotificationOperation) {
	if o == nil {
		return
	}
	o.Session = op.Session
	o.Notification = op.Notification
	o.Return = op.Return
}
func (o *RegisterPushNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterPushNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterPushNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DummyRPCOperation structure represents the EcDummyRpc operation
type xxx_DummyRPCOperation struct {
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *xxx_DummyRPCOperation) OpNum() int { return 6 }

func (o *xxx_DummyRPCOperation) OpName() string { return "/emsmdb/v0.81/EcDummyRpc" }

func (o *xxx_DummyRPCOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DummyRPCOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_DummyRPCOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_DummyRPCOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DummyRPCOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DummyRPCOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DummyRPCRequest structure represents the EcDummyRpc operation request
type DummyRPCRequest struct {
}

func (o *DummyRPCRequest) xxx_ToOp(ctx context.Context, op *xxx_DummyRPCOperation) *xxx_DummyRPCOperation {
	if op == nil {
		op = &xxx_DummyRPCOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *DummyRPCRequest) xxx_FromOp(ctx context.Context, op *xxx_DummyRPCOperation) {
	if o == nil {
		return
	}
}
func (o *DummyRPCRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DummyRPCRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DummyRPCOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DummyRPCResponse structure represents the EcDummyRpc operation response
type DummyRPCResponse struct {
	// Return: The EcDummyRpc return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DummyRPCResponse) xxx_ToOp(ctx context.Context, op *xxx_DummyRPCOperation) *xxx_DummyRPCOperation {
	if op == nil {
		op = &xxx_DummyRPCOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *DummyRPCResponse) xxx_FromOp(ctx context.Context, op *xxx_DummyRPCOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DummyRPCResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DummyRPCResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DummyRPCOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DoConnectExOperation structure represents the EcDoConnectEx operation
type xxx_DoConnectExOperation struct {
	Session             *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
	UserDN              string          `idl:"name:szUserDN;string" json:"user_dn"`
	Flags               uint32          `idl:"name:ulFlags" json:"flags"`
	UserDNHash          uint32          `idl:"name:ulConMod" json:"user_dn_hash"`
	LimitLength         uint32          `idl:"name:cbLimit" json:"limit_length"`
	CodePageID          uint32          `idl:"name:ulCpid" json:"code_page_id"`
	LocalIDString       uint32          `idl:"name:ulLcidString" json:"local_id_string"`
	LocalIDSort         uint32          `idl:"name:ulLcidSort" json:"local_id_sort"`
	RemoteSessionLink   uint32          `idl:"name:ulIcxrLink" json:"remote_session_link"`
	CanConvertCodePages uint16          `idl:"name:usFCanConvertCodePages" json:"can_convert_code_pages"`
	PollsMax            uint32          `idl:"name:pcmsPollsMax" json:"polls_max"`
	RetryCount          uint32          `idl:"name:pcRetry" json:"retry_count"`
	RetryDelay          uint32          `idl:"name:pcmsRetryDelay" json:"retry_delay"`
	RemoteSession       uint16          `idl:"name:picxr" json:"remote_session"`
	DNPrefix            string          `idl:"name:szDNPrefix;string" json:"dn_prefix"`
	DisplayName         string          `idl:"name:szDisplayName;string" json:"display_name"`
	ClientVersion       []uint16        `idl:"name:rgwClientVersion" json:"client_version"`
	ServerVersion       []uint16        `idl:"name:rgwServerVersion" json:"server_version"`
	BestVersion         []uint16        `idl:"name:rgwBestVersion" json:"best_version"`
	Timestamp           uint32          `idl:"name:pulTimeStamp" json:"timestamp"`
	AuxIn               []byte          `idl:"name:rgbAuxIn;size_is:(cbAuxIn)" json:"aux_in"`
	AuxInLength         uint32          `idl:"name:cbAuxIn" json:"aux_in_length"`
	AuxOut              []byte          `idl:"name:rgbAuxOut;size_is:(pcbAuxOut);length_is:(pcbAuxOut)" json:"aux_out"`
	AuxOutLength        uint32          `idl:"name:pcbAuxOut" json:"aux_out_length"`
	Return              int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_DoConnectExOperation) OpNum() int { return 10 }

func (o *xxx_DoConnectExOperation) OpName() string { return "/emsmdb/v0.81/EcDoConnectEx" }

func (o *xxx_DoConnectExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.AuxIn != nil && o.AuxInLength == 0 {
		o.AuxInLength = uint32(len(o.AuxIn))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoConnectExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// szUserDN {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.UserDN); err != nil {
			return err
		}
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// ulConMod {in} (1:(uint32))
	{
		if err := w.WriteData(o.UserDNHash); err != nil {
			return err
		}
	}
	// cbLimit {in} (1:(uint32))
	{
		if err := w.WriteData(o.LimitLength); err != nil {
			return err
		}
	}
	// ulCpid {in} (1:(uint32))
	{
		if err := w.WriteData(o.CodePageID); err != nil {
			return err
		}
	}
	// ulLcidString {in} (1:(uint32))
	{
		if err := w.WriteData(o.LocalIDString); err != nil {
			return err
		}
	}
	// ulLcidSort {in} (1:(uint32))
	{
		if err := w.WriteData(o.LocalIDSort); err != nil {
			return err
		}
	}
	// ulIcxrLink {in} (1:(uint32))
	{
		if err := w.WriteData(o.RemoteSessionLink); err != nil {
			return err
		}
	}
	// usFCanConvertCodePages {in} (1:(uint16))
	{
		if err := w.WriteData(o.CanConvertCodePages); err != nil {
			return err
		}
	}
	// rgwClientVersion {in} (1:[3](uint16))
	{
		for i1 := range o.ClientVersion {
			i1 := i1
			if uint64(i1) >= 3 {
				break
			}
			if err := w.WriteData(o.ClientVersion[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ClientVersion); uint64(i1) < 3; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// pulTimeStamp {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Timestamp); err != nil {
			return err
		}
	}
	// rgbAuxIn {in} (1:[dim:0,size_is=cbAuxIn](uchar))
	{
		dimSize1 := uint64(o.AuxInLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.AuxIn {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AuxIn[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AuxIn); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbAuxIn {in} (1:(uint32))
	{
		if err := w.WriteData(o.AuxInLength); err != nil {
			return err
		}
	}
	// pcbAuxOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,4104), alias=SMALL_RANGE_ULONG}(uint32))
	{
		if err := w.WriteData(o.AuxOutLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoConnectExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// szUserDN {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.UserDN); err != nil {
			return err
		}
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// ulConMod {in} (1:(uint32))
	{
		if err := w.ReadData(&o.UserDNHash); err != nil {
			return err
		}
	}
	// cbLimit {in} (1:(uint32))
	{
		if err := w.ReadData(&o.LimitLength); err != nil {
			return err
		}
	}
	// ulCpid {in} (1:(uint32))
	{
		if err := w.ReadData(&o.CodePageID); err != nil {
			return err
		}
	}
	// ulLcidString {in} (1:(uint32))
	{
		if err := w.ReadData(&o.LocalIDString); err != nil {
			return err
		}
	}
	// ulLcidSort {in} (1:(uint32))
	{
		if err := w.ReadData(&o.LocalIDSort); err != nil {
			return err
		}
	}
	// ulIcxrLink {in} (1:(uint32))
	{
		if err := w.ReadData(&o.RemoteSessionLink); err != nil {
			return err
		}
	}
	// usFCanConvertCodePages {in} (1:(uint16))
	{
		if err := w.ReadData(&o.CanConvertCodePages); err != nil {
			return err
		}
	}
	// rgwClientVersion {in} (1:[3](uint16))
	{
		o.ClientVersion = make([]uint16, 3)
		for i1 := range o.ClientVersion {
			i1 := i1
			if err := w.ReadData(&o.ClientVersion[i1]); err != nil {
				return err
			}
		}
	}
	// pulTimeStamp {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Timestamp); err != nil {
			return err
		}
	}
	// rgbAuxIn {in} (1:[dim:0,size_is=cbAuxIn](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AuxIn", sizeInfo[0])
		}
		o.AuxIn = make([]byte, sizeInfo[0])
		for i1 := range o.AuxIn {
			i1 := i1
			if err := w.ReadData(&o.AuxIn[i1]); err != nil {
				return err
			}
		}
	}
	// cbAuxIn {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AuxInLength); err != nil {
			return err
		}
	}
	// pcbAuxOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,4104), alias=SMALL_RANGE_ULONG}(uint32))
	{
		if err := w.ReadData(&o.AuxOutLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoConnectExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AuxOut != nil && o.AuxOutLength == 0 {
		o.AuxOutLength = uint32(len(o.AuxOut))
	}
	if o.AuxOut != nil && o.AuxOutLength == 0 {
		o.AuxOutLength = uint32(len(o.AuxOut))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoConnectExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pcxh {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session != nil {
			if err := o.Session.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oxcrpc.Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pcmsPollsMax {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.PollsMax); err != nil {
			return err
		}
	}
	// pcRetry {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RetryCount); err != nil {
			return err
		}
	}
	// pcmsRetryDelay {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RetryDelay); err != nil {
			return err
		}
	}
	// picxr {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.RemoteSession); err != nil {
			return err
		}
	}
	// szDNPrefix {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](uchar))
	{
		if o.DNPrefix != "" {
			_ptr_szDNPrefix := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.DNPrefix); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DNPrefix, _ptr_szDNPrefix); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// szDisplayName {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](uchar))
	{
		if o.DisplayName != "" {
			_ptr_szDisplayName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.DisplayName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DisplayName, _ptr_szDisplayName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// rgwServerVersion {out} (1:[3](uint16))
	{
		for i1 := range o.ServerVersion {
			i1 := i1
			if uint64(i1) >= 3 {
				break
			}
			if err := w.WriteData(o.ServerVersion[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerVersion); uint64(i1) < 3; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// rgwBestVersion {out} (1:[3](uint16))
	{
		for i1 := range o.BestVersion {
			i1 := i1
			if uint64(i1) >= 3 {
				break
			}
			if err := w.WriteData(o.BestVersion[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.BestVersion); uint64(i1) < 3; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// pulTimeStamp {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Timestamp); err != nil {
			return err
		}
	}
	// rgbAuxOut {out} (1:[dim:0,size_is=pcbAuxOut,length_is=pcbAuxOut](uchar))
	{
		dimSize1 := uint64(o.AuxOutLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.AuxOutLength)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.AuxOut {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AuxOut[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AuxOut); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pcbAuxOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,4104), alias=SMALL_RANGE_ULONG}(uint32))
	{
		if err := w.WriteData(o.AuxOutLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoConnectExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pcxh {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session == nil {
			o.Session = &oxcrpc.Session{}
		}
		if err := o.Session.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pcmsPollsMax {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.PollsMax); err != nil {
			return err
		}
	}
	// pcRetry {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RetryCount); err != nil {
			return err
		}
	}
	// pcmsRetryDelay {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RetryDelay); err != nil {
			return err
		}
	}
	// picxr {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.RemoteSession); err != nil {
			return err
		}
	}
	// szDNPrefix {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](uchar))
	{
		_ptr_szDNPrefix := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.DNPrefix); err != nil {
				return err
			}
			return nil
		})
		_s_szDNPrefix := func(ptr interface{}) { o.DNPrefix = *ptr.(*string) }
		if err := w.ReadPointer(&o.DNPrefix, _s_szDNPrefix, _ptr_szDNPrefix); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// szDisplayName {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](uchar))
	{
		_ptr_szDisplayName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.DisplayName); err != nil {
				return err
			}
			return nil
		})
		_s_szDisplayName := func(ptr interface{}) { o.DisplayName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DisplayName, _s_szDisplayName, _ptr_szDisplayName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// rgwServerVersion {out} (1:[3](uint16))
	{
		o.ServerVersion = make([]uint16, 3)
		for i1 := range o.ServerVersion {
			i1 := i1
			if err := w.ReadData(&o.ServerVersion[i1]); err != nil {
				return err
			}
		}
	}
	// rgwBestVersion {out} (1:[3](uint16))
	{
		o.BestVersion = make([]uint16, 3)
		for i1 := range o.BestVersion {
			i1 := i1
			if err := w.ReadData(&o.BestVersion[i1]); err != nil {
				return err
			}
		}
	}
	// pulTimeStamp {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Timestamp); err != nil {
			return err
		}
	}
	// rgbAuxOut {out} (1:[dim:0,size_is=pcbAuxOut,length_is=pcbAuxOut](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AuxOut", sizeInfo[0])
		}
		o.AuxOut = make([]byte, sizeInfo[0])
		for i1 := range o.AuxOut {
			i1 := i1
			if err := w.ReadData(&o.AuxOut[i1]); err != nil {
				return err
			}
		}
	}
	// pcbAuxOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,4104), alias=SMALL_RANGE_ULONG}(uint32))
	{
		if err := w.ReadData(&o.AuxOutLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DoConnectExRequest structure represents the EcDoConnectEx operation request
type DoConnectExRequest struct {
	// szUserDN: The DN of the user who is calling the EcDoConnectEx method in a directory
	// service. The value of the szUserDN parameter is similar to the following: "/o=First
	// Organization/ou=First Administrative Group/cn=recipients/cn=janedow".
	UserDN string `idl:"name:szUserDN;string" json:"user_dn"`
	// ulFlags: A flag value that designates the type of connection being established. On
	// input, this parameter contains connection bits that MUST be set; all flag values
	// not in the following table are reserved connection flags.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 | Requests connection without administrator privilege.                             |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | Requests administrator behavior, which causes the server to check that the user  |
	//	|            | has administrator privilege.                                                     |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00008000 | If this flag is not passed and the client version (as specified by the           |
	//	|            | rgwClientVersion parameter) is less than 12.00.0000.000 and no public folders    |
	//	|            | are configured within the messaging system, the server MUST fail the connection  |
	//	|            | attempt with error code ecClientVerDisallowed. The AUX_EXORGINFO auxiliary       |
	//	|            | block structure, specified in section 2.2.2.2.17, informs the client of the      |
	//	|            | presence of public folders within the organization. The use of the AUX_EXORGINFO |
	//	|            | auxiliary block structure is further defined in section 3.1.4.1.2.1. If this     |
	//	|            | flag is passed and the client version (as specified by the rgwClientVersion      |
	//	|            | parameter) is less than 12.00.0000.000, the server MUST NOT fail the connection  |
	//	|            | attempt due to public folders not being configured within the messaging system.  |
	//	|            | If the client version (as specified by the rgwClientVersion parameter) is        |
	//	|            | greater than or equal to 12.00.0000.000, the server MUST NOT fail the connection |
	//	|            | attempt due to public folders not being configured within the messaging system   |
	//	|            | (regardless of whether or not this flag is passed).                              |
	//	+------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// ulConMod: A client-derived 32-bit hash value of the DN passed in the szUserDN parameter.
	// The server determines which public folder replica to use when accessing public folder
	// information when more than one replica of a folder exists. The hash can be used to
	// distribute client access across replicas in a deterministic way for load balancing.
	UserDNHash uint32 `idl:"name:ulConMod" json:"user_dn_hash"`
	// cbLimit: MUST be set to zero when sent and MUST be ignored when received.
	LimitLength uint32 `idl:"name:cbLimit" json:"limit_length"`
	// ulCpid: The code page in which text data is sent. If the Unicode format is not requested
	// by the client on subsequent calls that use this Session Context, the ulCpid parameter
	// sets the code page to be used in subsequent calls.
	CodePageID uint32 `idl:"name:ulCpid" json:"code_page_id"`
	// ulLcidString: The local ID for everything other than sorting.
	LocalIDString uint32 `idl:"name:ulLcidString" json:"local_id_string"`
	// ulLcidSort: The local ID for sorting.
	LocalIDSort uint32 `idl:"name:ulLcidSort" json:"local_id_sort"`
	// ulIcxrLink: A value used to link the Session Context created by this call with a
	// currently existing Session Context on the server. To request Session Context linking,
	// the client MUST pass the value of 0xFFFFFFFF. To link to an existing Session Context,
	// this value is the session index value returned in the piCxr parameter from a previous
	// call to the EcDoConnectEx method. In addition to passing the session index in the
	// ulIcxrLink parameter, the client sets the pulTimeStamp parameter to the value that
	// was returned in the pulTimeStamp parameter from the previous call to the EcDoConnectEx
	// method. These two values MUST be used by the server to identify an active session
	// with the same session index and session creation time stamp. If a session is found,
	// the server MUST link the Session Context created by this call with the one found.<8>
	//
	// A server allows Session Context linking for the following reasons:
	//
	// *
	//
	// To consume a single Client Access License (CAL) ( c5869b4a-c0d3-4ceb-80f4-c47a250ddf3d#gt_0e48e7fc-c851-4692-9b39-8bfa227f4aba
	// ) for all the connections made from a single client computer. This gives a client
	// the ability to open multiple independent connections using more than one Session
	// Context on the server but be seen to the server as only consuming a single CAL. <9>
	// ( f2038fa2-e1b4-4b77-8e29-08c865c1cb3a#Appendix_A_9 )
	//
	// *
	//
	// To get pending notification information for other sessions on the same client computer.
	// For details, see [MS-OXCNOTIF] ( ../ms-oxcnotif/7c7d1653-5dfb-42f1-9410-fc4e48e10731
	// ).
	//
	// Note that the ulIcxrLink parameter is defined as a 32-bit value. Other than passing
	// 0xFFFFFFFF for no Session Context linking, the server only uses the low-order 16
	// bits as the session index. This value is the value returned in the piCxr parameter
	// from a previous call to the EcDoConnectEx method, which is the session index and
	// defined as a 16-bit value.
	RemoteSessionLink uint32 `idl:"name:ulIcxrLink" json:"remote_session_link"`
	// usFCanConvertCodePages: This parameter is reserved. The client MUST pass a value
	// of 0x0001.
	CanConvertCodePages uint16 `idl:"name:usFCanConvertCodePages" json:"can_convert_code_pages"`
	// rgwClientVersion: The client protocol version that the server uses to determine what
	// protocol functionality the client supports. For more details about how version numbers
	// are interpreted from the wire data, see section 3.2.4.1.3.
	ClientVersion []uint16 `idl:"name:rgwClientVersion" json:"client_version"`
	// pulTimeStamp: The creation time of the newly created Session Context. On input, a
	// value used with the ulIcxrLink parameter to link the Session Context created by this
	// call with an existing Session Context. If the ulIcxrLink parameter is not set to
	// 0xFFFFFFFF, the client MUST pass in the value of the pulTimeStamp parameter returned
	// from the server on a previous call to the EcDoConnectEx method. For more details,
	// see the ulIcxrLink and piCxr parameter descriptions earlier in this section. If the
	// server supports Session Context linking, the server verifies that there is a Session
	// Context state with the unique identifier in the ulIcxrLink parameter, and the Session
	// Context state has a creation time stamp equal to the value passed in this parameter.
	// If so, the server MUST link the Session Context created by this call with the one
	// found. If no such Session Context state is found, the server does not fail the EcDoConnectEx
	// method call but simply does not link the Session Contexts.<11>
	//
	// On output, the server has to return a time stamp in which the new Session Context
	// was created. The server saves the Session Context creation time stamp within the
	// Session Context state for later use if a client attempts to link Session Contexts.
	Timestamp uint32 `idl:"name:pulTimeStamp" json:"timestamp"`
	// rgbAuxIn: An auxiliary payload buffer prefixed by an RPC_HEADER_EXT structure, as
	// specified in section 2.2.2.1. Information stored in this structure determines how
	// to interpret the data that follows the structure. The length of the auxiliary payload
	// buffer that includes the RPC_HEADER_EXT structure is contained in the cbAuxIn parameter.
	//
	// For details about how to access the embedded auxiliary payload buffer, see section
	// 3.1.4.1.1. For details about how to interpret the auxiliary payload data, see section
	// 3.1.4.1.2.
	AuxIn []byte `idl:"name:rgbAuxIn;size_is:(cbAuxIn)" json:"aux_in"`
	// cbAuxIn: The length of the rgbAuxIn parameter. If this value on input is larger than
	// 0x00001008 bytes in size, the server SHOULD<12> fail with the RPC status code RPC_X_BAD_STUB_DATA
	// (0x000006F7). If this value is greater than 0x00000000 and less than 0x00000008,
	// the server SHOULD<13> fail with ecRpcFailed (0x80040115). For more information on
	// returning RPC status codes, see [C706].
	AuxInLength uint32 `idl:"name:cbAuxIn" json:"aux_in_length"`
	// pcbAuxOut: The length of the rgbAuxOut parameter. If this value on input is larger
	// than 0x00001008, the server MUST fail with the RPC status code RPC_X_BAD_STUB_DATA
	// (0x000006F7).
	//
	// On output, this parameter contains the size of the data to be returned in the rgbAuxOut
	// parameter.
	AuxOutLength uint32 `idl:"name:pcbAuxOut" json:"aux_out_length"`
}

func (o *DoConnectExRequest) xxx_ToOp(ctx context.Context, op *xxx_DoConnectExOperation) *xxx_DoConnectExOperation {
	if op == nil {
		op = &xxx_DoConnectExOperation{}
	}
	if o == nil {
		return op
	}
	op.UserDN = o.UserDN
	op.Flags = o.Flags
	op.UserDNHash = o.UserDNHash
	op.LimitLength = o.LimitLength
	op.CodePageID = o.CodePageID
	op.LocalIDString = o.LocalIDString
	op.LocalIDSort = o.LocalIDSort
	op.RemoteSessionLink = o.RemoteSessionLink
	op.CanConvertCodePages = o.CanConvertCodePages
	op.ClientVersion = o.ClientVersion
	op.Timestamp = o.Timestamp
	op.AuxIn = o.AuxIn
	op.AuxInLength = o.AuxInLength
	op.AuxOutLength = o.AuxOutLength
	return op
}

func (o *DoConnectExRequest) xxx_FromOp(ctx context.Context, op *xxx_DoConnectExOperation) {
	if o == nil {
		return
	}
	o.UserDN = op.UserDN
	o.Flags = op.Flags
	o.UserDNHash = op.UserDNHash
	o.LimitLength = op.LimitLength
	o.CodePageID = op.CodePageID
	o.LocalIDString = op.LocalIDString
	o.LocalIDSort = op.LocalIDSort
	o.RemoteSessionLink = op.RemoteSessionLink
	o.CanConvertCodePages = op.CanConvertCodePages
	o.ClientVersion = op.ClientVersion
	o.Timestamp = op.Timestamp
	o.AuxIn = op.AuxIn
	o.AuxInLength = op.AuxInLength
	o.AuxOutLength = op.AuxOutLength
}
func (o *DoConnectExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DoConnectExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoConnectExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DoConnectExResponse structure represents the EcDoConnectEx operation response
type DoConnectExResponse struct {
	// pcxh: A session context handle for the client. On success, the server MUST return
	// a unique value to be used as a session context handle.
	//
	// On failure, the server MUST return a zero value as the session context handle.
	Session *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
	// pcmsPollsMax: An implementation-dependent value that specifies the number of milliseconds
	// that a client waits between polling the server for event information. If the client
	// or server does not support making asynchronous RPCs for notifications as specified
	// in section 3.3.4.1, or the client is unable to receive notifications via UDP datagrams,
	// as specified in [MS-OXCNOTIF] section 3.2.5.4 and [MS-OXCNOTIF] section 3.2.5.5.2,
	// the client can poll the server to determine whether any events are pending for the
	// client.
	PollsMax uint32 `idl:"name:pcmsPollsMax" json:"polls_max"`
	// pcRetry: An implementation-dependent value that specifies the number of times a client
	// retries future RPCs using the session context handle returned in this call. This
	// is for client RPCs that fail with RPC status code RPC_S_SERVER_TOO_BUSY (0x000006BB).
	// This is a suggested retry count for the client and is not to be enforced by the server.
	// For more details about circumstances under which the RPC_S_SERVER_TOO_BUSY status
	// code is returned, see [MS-OXCROPS] section 3.2.4.2. For more details about how the
	// client handles the RPC_S_SERVER_TOO_BUSY status code, see section 3.2.4.4.
	RetryCount uint32 `idl:"name:pcRetry" json:"retry_count"`
	// pcmsRetryDelay: An implementation-dependent value that specifies the number of milliseconds
	// a client waits before retrying a failed RPC. If any future RPC to the server using
	// the session context handle returned in this call fails with RPC status code RPC_S_SERVER_TOO_BUSY
	// (0x000006BB), the client waits the number of milliseconds specified in this output
	// parameter before retrying the call. The number of times a client retries is returned
	// in the pcRetry parameter. This is a suggested delay for the client and is not to
	// be enforced by the server.
	RetryDelay    uint32 `idl:"name:pcmsRetryDelay" json:"retry_delay"`
	RemoteSession uint16 `idl:"name:picxr" json:"remote_session"`
	// szDNPrefix: An implementation-dependent value that specifies a DN prefix that is
	// used to build message recipients. An empty value indicates that there is nothing
	// to prepend to recipient entries on messages.
	DNPrefix string `idl:"name:szDNPrefix;string" json:"dn_prefix"`
	// szDisplayName: The display name of the user associated with the szUserDN parameter.
	DisplayName string `idl:"name:szDisplayName;string" json:"display_name"`
	// rgwServerVersion: The server protocol version that the client uses to determine what
	// protocol functionality the server supports. For details about how version numbers
	// are interpreted from the wire data, see section 3.1.4.1.3.
	ServerVersion []uint16 `idl:"name:rgwServerVersion" json:"server_version"`
	// rgwBestVersion: The minimum client protocol version that the server supports. This
	// information is useful if the call to the EcDoConnectEx method fails with return code
	// ecVersionMismatch. On success, the server returns the value passed in the rgwClientVersion
	// parameter by the client. The server cannot perform any client protocol version negotiation.
	// The server can either return the minimum client protocol version required to access
	// the server and fail the call with ecVersionMismatch (0x80040110) or allow the client
	// and return the value passed by the client in the rgwClientVersion parameter. The
	// server implementation sets the minimum client protocol version that is supported
	// by the server. For details about how version numbers are interpreted from the wire
	// data, see section 3.1.4.1.3.1.
	BestVersion []uint16 `idl:"name:rgwBestVersion" json:"best_version"`
	// pulTimeStamp: The creation time of the newly created Session Context. On input, a
	// value used with the ulIcxrLink parameter to link the Session Context created by this
	// call with an existing Session Context. If the ulIcxrLink parameter is not set to
	// 0xFFFFFFFF, the client MUST pass in the value of the pulTimeStamp parameter returned
	// from the server on a previous call to the EcDoConnectEx method. For more details,
	// see the ulIcxrLink and piCxr parameter descriptions earlier in this section. If the
	// server supports Session Context linking, the server verifies that there is a Session
	// Context state with the unique identifier in the ulIcxrLink parameter, and the Session
	// Context state has a creation time stamp equal to the value passed in this parameter.
	// If so, the server MUST link the Session Context created by this call with the one
	// found. If no such Session Context state is found, the server does not fail the EcDoConnectEx
	// method call but simply does not link the Session Contexts.<11>
	//
	// On output, the server has to return a time stamp in which the new Session Context
	// was created. The server saves the Session Context creation time stamp within the
	// Session Context state for later use if a client attempts to link Session Contexts.
	Timestamp uint32 `idl:"name:pulTimeStamp" json:"timestamp"`
	// rgbAuxOut: An auxiliary payload buffer prefixed by an RPC_HEADER_EXT structure (section
	// 2.2.2.1). On output, the server can return auxiliary payload data to the client in
	// this parameter. The server MUST include an RPC_HEADER_EXT structure before the auxiliary
	// payload data.
	//
	// For details about how to access the embedded auxiliary payload buffer, see section
	// 3.1.4.1.1. For details about how to interpret the auxiliary payload data, see section
	// 3.1.4.1.2.
	AuxOut []byte `idl:"name:rgbAuxOut;size_is:(pcbAuxOut);length_is:(pcbAuxOut)" json:"aux_out"`
	// pcbAuxOut: The length of the rgbAuxOut parameter. If this value on input is larger
	// than 0x00001008, the server MUST fail with the RPC status code RPC_X_BAD_STUB_DATA
	// (0x000006F7).
	//
	// On output, this parameter contains the size of the data to be returned in the rgbAuxOut
	// parameter.
	AuxOutLength uint32 `idl:"name:pcbAuxOut" json:"aux_out_length"`
	// Return: The EcDoConnectEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DoConnectExResponse) xxx_ToOp(ctx context.Context, op *xxx_DoConnectExOperation) *xxx_DoConnectExOperation {
	if op == nil {
		op = &xxx_DoConnectExOperation{}
	}
	if o == nil {
		return op
	}
	op.Session = o.Session
	op.PollsMax = o.PollsMax
	op.RetryCount = o.RetryCount
	op.RetryDelay = o.RetryDelay
	op.RemoteSession = o.RemoteSession
	op.DNPrefix = o.DNPrefix
	op.DisplayName = o.DisplayName
	op.ServerVersion = o.ServerVersion
	op.BestVersion = o.BestVersion
	op.Timestamp = o.Timestamp
	op.AuxOut = o.AuxOut
	op.AuxOutLength = o.AuxOutLength
	op.Return = o.Return
	return op
}

func (o *DoConnectExResponse) xxx_FromOp(ctx context.Context, op *xxx_DoConnectExOperation) {
	if o == nil {
		return
	}
	o.Session = op.Session
	o.PollsMax = op.PollsMax
	o.RetryCount = op.RetryCount
	o.RetryDelay = op.RetryDelay
	o.RemoteSession = op.RemoteSession
	o.DNPrefix = op.DNPrefix
	o.DisplayName = op.DisplayName
	o.ServerVersion = op.ServerVersion
	o.BestVersion = op.BestVersion
	o.Timestamp = op.Timestamp
	o.AuxOut = op.AuxOut
	o.AuxOutLength = op.AuxOutLength
	o.Return = op.Return
}
func (o *DoConnectExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DoConnectExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoConnectExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DoRPCExt2Operation structure represents the EcDoRpcExt2 operation
type xxx_DoRPCExt2Operation struct {
	Session      *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
	Flags        uint32          `idl:"name:pulFlags" json:"flags"`
	In           []byte          `idl:"name:rgbIn;size_is:(cbIn)" json:"in"`
	InLength     uint32          `idl:"name:cbIn" json:"in_length"`
	Out          []byte          `idl:"name:rgbOut;size_is:(pcbOut);length_is:(pcbOut)" json:"out"`
	OutLength    uint32          `idl:"name:pcbOut" json:"out_length"`
	AuxIn        []byte          `idl:"name:rgbAuxIn;size_is:(cbAuxIn)" json:"aux_in"`
	AuxInLength  uint32          `idl:"name:cbAuxIn" json:"aux_in_length"`
	AuxOut       []byte          `idl:"name:rgbAuxOut;size_is:(pcbAuxOut);length_is:(pcbAuxOut)" json:"aux_out"`
	AuxOutLength uint32          `idl:"name:pcbAuxOut" json:"aux_out_length"`
	TransTime    uint32          `idl:"name:pulTransTime" json:"trans_time"`
	Return       int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_DoRPCExt2Operation) OpNum() int { return 11 }

func (o *xxx_DoRPCExt2Operation) OpName() string { return "/emsmdb/v0.81/EcDoRpcExt2" }

func (o *xxx_DoRPCExt2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.In != nil && o.InLength == 0 {
		o.InLength = uint32(len(o.In))
	}
	if o.AuxIn != nil && o.AuxInLength == 0 {
		o.AuxInLength = uint32(len(o.AuxIn))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoRPCExt2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session != nil {
			if err := o.Session.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oxcrpc.Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pulFlags {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// rgbIn {in} (1:[dim:0,size_is=cbIn](uchar))
	{
		dimSize1 := uint64(o.InLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.In {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.In[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.In); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbIn {in} (1:(uint32))
	{
		if err := w.WriteData(o.InLength); err != nil {
			return err
		}
	}
	// pcbOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,262144), alias=BIG_RANGE_ULONG}(uint32))
	{
		if err := w.WriteData(o.OutLength); err != nil {
			return err
		}
	}
	// rgbAuxIn {in} (1:[dim:0,size_is=cbAuxIn](uchar))
	{
		dimSize1 := uint64(o.AuxInLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.AuxIn {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AuxIn[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AuxIn); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbAuxIn {in} (1:(uint32))
	{
		if err := w.WriteData(o.AuxInLength); err != nil {
			return err
		}
	}
	// pcbAuxOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,4104), alias=SMALL_RANGE_ULONG}(uint32))
	{
		if err := w.WriteData(o.AuxOutLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoRPCExt2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session == nil {
			o.Session = &oxcrpc.Session{}
		}
		if err := o.Session.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pulFlags {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// rgbIn {in} (1:[dim:0,size_is=cbIn](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.In", sizeInfo[0])
		}
		o.In = make([]byte, sizeInfo[0])
		for i1 := range o.In {
			i1 := i1
			if err := w.ReadData(&o.In[i1]); err != nil {
				return err
			}
		}
	}
	// cbIn {in} (1:(uint32))
	{
		if err := w.ReadData(&o.InLength); err != nil {
			return err
		}
	}
	// pcbOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,262144), alias=BIG_RANGE_ULONG}(uint32))
	{
		if err := w.ReadData(&o.OutLength); err != nil {
			return err
		}
	}
	// rgbAuxIn {in} (1:[dim:0,size_is=cbAuxIn](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AuxIn", sizeInfo[0])
		}
		o.AuxIn = make([]byte, sizeInfo[0])
		for i1 := range o.AuxIn {
			i1 := i1
			if err := w.ReadData(&o.AuxIn[i1]); err != nil {
				return err
			}
		}
	}
	// cbAuxIn {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AuxInLength); err != nil {
			return err
		}
	}
	// pcbAuxOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,4104), alias=SMALL_RANGE_ULONG}(uint32))
	{
		if err := w.ReadData(&o.AuxOutLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoRPCExt2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Out != nil && o.OutLength == 0 {
		o.OutLength = uint32(len(o.Out))
	}
	if o.Out != nil && o.OutLength == 0 {
		o.OutLength = uint32(len(o.Out))
	}
	if o.AuxOut != nil && o.AuxOutLength == 0 {
		o.AuxOutLength = uint32(len(o.AuxOut))
	}
	if o.AuxOut != nil && o.AuxOutLength == 0 {
		o.AuxOutLength = uint32(len(o.AuxOut))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoRPCExt2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session != nil {
			if err := o.Session.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oxcrpc.Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pulFlags {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// rgbOut {out} (1:[dim:0,size_is=pcbOut,length_is=pcbOut](uchar))
	{
		dimSize1 := uint64(o.OutLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.OutLength)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.Out {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Out[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Out); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pcbOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,262144), alias=BIG_RANGE_ULONG}(uint32))
	{
		if err := w.WriteData(o.OutLength); err != nil {
			return err
		}
	}
	// rgbAuxOut {out} (1:[dim:0,size_is=pcbAuxOut,length_is=pcbAuxOut](uchar))
	{
		dimSize1 := uint64(o.AuxOutLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.AuxOutLength)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.AuxOut {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AuxOut[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AuxOut); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pcbAuxOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,4104), alias=SMALL_RANGE_ULONG}(uint32))
	{
		if err := w.WriteData(o.AuxOutLength); err != nil {
			return err
		}
	}
	// pulTransTime {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TransTime); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoRPCExt2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pcxh {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session == nil {
			o.Session = &oxcrpc.Session{}
		}
		if err := o.Session.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pulFlags {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// rgbOut {out} (1:[dim:0,size_is=pcbOut,length_is=pcbOut](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Out", sizeInfo[0])
		}
		o.Out = make([]byte, sizeInfo[0])
		for i1 := range o.Out {
			i1 := i1
			if err := w.ReadData(&o.Out[i1]); err != nil {
				return err
			}
		}
	}
	// pcbOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,262144), alias=BIG_RANGE_ULONG}(uint32))
	{
		if err := w.ReadData(&o.OutLength); err != nil {
			return err
		}
	}
	// rgbAuxOut {out} (1:[dim:0,size_is=pcbAuxOut,length_is=pcbAuxOut](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AuxOut", sizeInfo[0])
		}
		o.AuxOut = make([]byte, sizeInfo[0])
		for i1 := range o.AuxOut {
			i1 := i1
			if err := w.ReadData(&o.AuxOut[i1]); err != nil {
				return err
			}
		}
	}
	// pcbAuxOut {in, out} (1:{pointer=ref}*(1))(2:{range=(0,4104), alias=SMALL_RANGE_ULONG}(uint32))
	{
		if err := w.ReadData(&o.AuxOutLength); err != nil {
			return err
		}
	}
	// pulTransTime {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TransTime); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DoRPCExt2Request structure represents the EcDoRpcExt2 operation request
type DoRPCExt2Request struct {
	// pcxh: A session context handle. On input, the client MUST pass a valid session context
	// handle that was created by calling the EcDoConnectEx method. The server uses the
	// session context handle to identify the Session Context to use for this call. On output,
	// the server MUST return the same session context handle on success.
	//
	// The server can destroy the session context handle by returning a zero session context
	// handle. Reasons for destroying the session context handle are implementation-dependent;
	// following are examples of why the server might destroy the session context handle:
	//
	// * The server determines that the ROP request ( c5869b4a-c0d3-4ceb-80f4-c47a250ddf3d#gt_edeadb0f-6571-49b7-8cce-5dc77b0793d6
	// ) payload in the rgbIn buffer is malformed or length parameters are not valid.
	//
	// * The session context handle that was passed in is not valid.
	//
	// * An attempt was made to access a mailbox ( c5869b4a-c0d3-4ceb-80f4-c47a250ddf3d#gt_d3ad0e15-adc9-4174-bacf-d929b57278b3
	// ) that is in the process of being moved.
	//
	// * An administrator has blocked a client that has an active connection.
	Session *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
	// pulFlags: The flags that describe the server output characteristics. On input, this
	// parameter contains flags that tell the server how to build the rgbOut parameter.
	//
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	|     FLAG      |            |                                                                                  |
	//	|     NAME      |   VALUE    |                                     MEANING                                      |
	//	|               |            |                                                                                  |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| NoCompression | 0x00000001 | The server MUST NOT compress ROP response payload (rgbOut) or auxiliary payload  |
	//	|               |            | (rgbAuxOut). For details about server behavior when this flag is absent, see     |
	//	|               |            | section 3.1.4.2.1.1.                                                             |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| NoXorMagic    | 0x00000002 | The server MUST NOT obfuscate the ROP response payload (rgbOut) or auxiliary     |
	//	|               |            | payload (rgbAuxOut). For details about server behavior when this flag is absent, |
	//	|               |            | see section 3.1.4.2.1.1.                                                         |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| Chain         | 0x00000004 | The client allows chaining of ROP response payloads.                             |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//
	// For details about how to use these flags, see section 3.1.4.2.1.1.
	//
	// On output, the server SHOULD<17> set this parameter to 0x00000000. The output flags
	// not in the table are reserved for future use.
	Flags uint32 `idl:"name:pulFlags" json:"flags"`
	// rgbIn: The ROP request payload. The ROP request payload is prefixed with an RPC_HEADER_EXT
	// header, as specified in section 2.2.2.1. Information stored in this header determines
	// how to interpret the data following the header. For details about how to access the
	// embedded ROP request payload, see section 3.1.4.2.1. The length of the ROP request
	// payload including the RPC_HEADER_EXT header is contained in the cbIn parameter.
	//
	// For more information about ROP buffers, see [MS-OXCROPS].
	In []byte `idl:"name:rgbIn;size_is:(cbIn)" json:"in"`
	// cbIn: The length of the ROP request payload. On input, this parameter contains the
	// length of the ROP request payload passed in the rgbIn parameter. The ROP request
	// payload includes the size of the ROPs plus the size of the RPC_HEADER_EXT structure.
	// The server SHOULD<18> fail with the RPC status code of RPC_X_BAD_STUB_DATA (0x000006F7)
	// if the request buffer is larger than 0x00040000 bytes in size. For more information
	// on returning RPC status codes, see [C706]. If the request buffer is smaller than
	// the size of the RPC_HEADER_EXT structure (0x00000008 bytes), the server SHOULD<19>
	// fail with error code ecRpcFailed (0x80040115).
	InLength uint32 `idl:"name:cbIn" json:"in_length"`
	// pcbOut: The maximum size of the rgbOut parameter. On input, this parameter contains
	// the maximum size of the rgbOut parameter. If the value in the pcbOut parameter on
	// input is less than 0x00000008, the server SHOULD<20> fail with error code ecRpcFailed
	// (0x80040115). If the value in the pcbOut parameter on input is larger than 0x00040000,
	// the server MUST fail with the RPC status code of RPC_X_BAD_STUB_DATA (0x000006F7).
	// If the server indicates it supports a larger ROP response payload through the ULTRA_LARGE_PACKED_DOWNLOAD_BUFFERS
	// flag being present in the Flags field of the AUX_SERVER_CAPABILITIES (section 2.2.2.2.19)
	// auxiliary block returned in rgbAuxOut field of the EcDoConnectEx method (section
	// 3.1.4.1), then the server MUST fail only if pcbOut is larger than 0x00100000.
	//
	// On output, this parameter contains the size of the ROP response payload, including
	// the size of the RPC_HEADER_EXT header in the rgbOut parameter. The server returns
	// 0x00000000 on failure because there is no ROP response payload. The client ignores
	// any data returned on failure.
	OutLength uint32 `idl:"name:pcbOut" json:"out_length"`
	// rgbAuxIn: The auxiliary payload buffer. The auxiliary payload buffer is prefixed
	// by an RPC_HEADER_EXT structure. Information stored in this header determines how
	// to interpret the data following the header. The length of the auxiliary payload buffer
	// including the RPC_HEADER_EXT header is contained in the cbAuxIn parameter.
	//
	// For details about how to access the embedded auxiliary payload buffer, see section
	// 3.1.4.2.1. For details about how to interpret the auxiliary payload data, see section
	// 3.1.4.2.2.
	AuxIn []byte `idl:"name:rgbAuxIn;size_is:(cbAuxIn)" json:"aux_in"`
	// cbAuxIn: The length of the auxiliary payload buffer. On input, this parameter contains
	// the length of the auxiliary payload buffer passed in the rgbAuxIn parameter. If the
	// request buffer value is larger than 0x00001008 bytes in size, the server SHOULD<21>
	// fail with the RPC status code RPC_X_BAD_STUB_DATA (0x000006F7).<22>
	AuxInLength uint32 `idl:"name:cbAuxIn" json:"aux_in_length"`
	// pcbAuxOut: The maximum length of the auxiliary payload buffer. On input, this parameter
	// contains the maximum length of the rgbAuxOut parameter. If this value on input is
	// larger than 0x00001008, the server MUST fail with the RPC status code RPC_X_BAD_STUB_DATA
	// (0x000006F7).
	//
	// On output, this parameter contains the size of the data to be returned in the rgbAuxOut
	// parameter.
	AuxOutLength uint32 `idl:"name:pcbAuxOut" json:"aux_out_length"`
}

func (o *DoRPCExt2Request) xxx_ToOp(ctx context.Context, op *xxx_DoRPCExt2Operation) *xxx_DoRPCExt2Operation {
	if op == nil {
		op = &xxx_DoRPCExt2Operation{}
	}
	if o == nil {
		return op
	}
	op.Session = o.Session
	op.Flags = o.Flags
	op.In = o.In
	op.InLength = o.InLength
	op.OutLength = o.OutLength
	op.AuxIn = o.AuxIn
	op.AuxInLength = o.AuxInLength
	op.AuxOutLength = o.AuxOutLength
	return op
}

func (o *DoRPCExt2Request) xxx_FromOp(ctx context.Context, op *xxx_DoRPCExt2Operation) {
	if o == nil {
		return
	}
	o.Session = op.Session
	o.Flags = op.Flags
	o.In = op.In
	o.InLength = op.InLength
	o.OutLength = op.OutLength
	o.AuxIn = op.AuxIn
	o.AuxInLength = op.AuxInLength
	o.AuxOutLength = op.AuxOutLength
}
func (o *DoRPCExt2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DoRPCExt2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoRPCExt2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DoRPCExt2Response structure represents the EcDoRpcExt2 operation response
type DoRPCExt2Response struct {
	// pcxh: A session context handle. On input, the client MUST pass a valid session context
	// handle that was created by calling the EcDoConnectEx method. The server uses the
	// session context handle to identify the Session Context to use for this call. On output,
	// the server MUST return the same session context handle on success.
	//
	// The server can destroy the session context handle by returning a zero session context
	// handle. Reasons for destroying the session context handle are implementation-dependent;
	// following are examples of why the server might destroy the session context handle:
	//
	// * The server determines that the ROP request ( c5869b4a-c0d3-4ceb-80f4-c47a250ddf3d#gt_edeadb0f-6571-49b7-8cce-5dc77b0793d6
	// ) payload in the rgbIn buffer is malformed or length parameters are not valid.
	//
	// * The session context handle that was passed in is not valid.
	//
	// * An attempt was made to access a mailbox ( c5869b4a-c0d3-4ceb-80f4-c47a250ddf3d#gt_d3ad0e15-adc9-4174-bacf-d929b57278b3
	// ) that is in the process of being moved.
	//
	// * An administrator has blocked a client that has an active connection.
	Session *oxcrpc.Session `idl:"name:pcxh;pointer:ref" json:"session"`
	// pulFlags: The flags that describe the server output characteristics. On input, this
	// parameter contains flags that tell the server how to build the rgbOut parameter.
	//
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	|     FLAG      |            |                                                                                  |
	//	|     NAME      |   VALUE    |                                     MEANING                                      |
	//	|               |            |                                                                                  |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| NoCompression | 0x00000001 | The server MUST NOT compress ROP response payload (rgbOut) or auxiliary payload  |
	//	|               |            | (rgbAuxOut). For details about server behavior when this flag is absent, see     |
	//	|               |            | section 3.1.4.2.1.1.                                                             |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| NoXorMagic    | 0x00000002 | The server MUST NOT obfuscate the ROP response payload (rgbOut) or auxiliary     |
	//	|               |            | payload (rgbAuxOut). For details about server behavior when this flag is absent, |
	//	|               |            | see section 3.1.4.2.1.1.                                                         |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//	| Chain         | 0x00000004 | The client allows chaining of ROP response payloads.                             |
	//	+---------------+------------+----------------------------------------------------------------------------------+
	//
	// For details about how to use these flags, see section 3.1.4.2.1.1.
	//
	// On output, the server SHOULD<17> set this parameter to 0x00000000. The output flags
	// not in the table are reserved for future use.
	Flags uint32 `idl:"name:pulFlags" json:"flags"`
	// rgbOut: The ROP response payload. The size of the payload is specified in the pcbOut
	// parameter. Like the ROP request payload, the ROP response payload is also prefixed
	// by a RPC_HEADER_EXT header. For details about how to format the ROP response payload,
	// see section 3.1.4.2.1. The size of the ROP response payload plus the RPC_HEADER_EXT
	// header is returned in the pcbOut parameter.
	Out []byte `idl:"name:rgbOut;size_is:(pcbOut);length_is:(pcbOut)" json:"out"`
	// pcbOut: The maximum size of the rgbOut parameter. On input, this parameter contains
	// the maximum size of the rgbOut parameter. If the value in the pcbOut parameter on
	// input is less than 0x00000008, the server SHOULD<20> fail with error code ecRpcFailed
	// (0x80040115). If the value in the pcbOut parameter on input is larger than 0x00040000,
	// the server MUST fail with the RPC status code of RPC_X_BAD_STUB_DATA (0x000006F7).
	// If the server indicates it supports a larger ROP response payload through the ULTRA_LARGE_PACKED_DOWNLOAD_BUFFERS
	// flag being present in the Flags field of the AUX_SERVER_CAPABILITIES (section 2.2.2.2.19)
	// auxiliary block returned in rgbAuxOut field of the EcDoConnectEx method (section
	// 3.1.4.1), then the server MUST fail only if pcbOut is larger than 0x00100000.
	//
	// On output, this parameter contains the size of the ROP response payload, including
	// the size of the RPC_HEADER_EXT header in the rgbOut parameter. The server returns
	// 0x00000000 on failure because there is no ROP response payload. The client ignores
	// any data returned on failure.
	OutLength uint32 `idl:"name:pcbOut" json:"out_length"`
	// rgbAuxOut: The auxiliary payload buffer. On output, the server MAY<23> return auxiliary
	// payload data to the client. The server MUST include a RPC_HEADER_EXT header before
	// the auxiliary payload data.
	AuxOut []byte `idl:"name:rgbAuxOut;size_is:(pcbAuxOut);length_is:(pcbAuxOut)" json:"aux_out"`
	// pcbAuxOut: The maximum length of the auxiliary payload buffer. On input, this parameter
	// contains the maximum length of the rgbAuxOut parameter. If this value on input is
	// larger than 0x00001008, the server MUST fail with the RPC status code RPC_X_BAD_STUB_DATA
	// (0x000006F7).
	//
	// On output, this parameter contains the size of the data to be returned in the rgbAuxOut
	// parameter.
	AuxOutLength uint32 `idl:"name:pcbAuxOut" json:"aux_out_length"`
	// pulTransTime: The time it took to execute this method. On output, the server stores
	// the number of milliseconds the call took to execute. This is the total elapsed time
	// from when the call is dispatched on the server to the point in which the server returns
	// the call. This is diagnostic information the client can use to determine the cause
	// of a slow response time from the server. The client can monitor the total elapsed
	// time across the RPC method call and, using this output parameter, can determine whether
	// time was spent transmitting the request/response on the network or processing it
	// on the server.
	TransTime uint32 `idl:"name:pulTransTime" json:"trans_time"`
	// Return: The EcDoRpcExt2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DoRPCExt2Response) xxx_ToOp(ctx context.Context, op *xxx_DoRPCExt2Operation) *xxx_DoRPCExt2Operation {
	if op == nil {
		op = &xxx_DoRPCExt2Operation{}
	}
	if o == nil {
		return op
	}
	op.Session = o.Session
	op.Flags = o.Flags
	op.Out = o.Out
	op.OutLength = o.OutLength
	op.AuxOut = o.AuxOut
	op.AuxOutLength = o.AuxOutLength
	op.TransTime = o.TransTime
	op.Return = o.Return
	return op
}

func (o *DoRPCExt2Response) xxx_FromOp(ctx context.Context, op *xxx_DoRPCExt2Operation) {
	if o == nil {
		return
	}
	o.Session = op.Session
	o.Flags = op.Flags
	o.Out = op.Out
	o.OutLength = op.OutLength
	o.AuxOut = op.AuxOut
	o.AuxOutLength = op.AuxOutLength
	o.TransTime = op.TransTime
	o.Return = op.Return
}
func (o *DoRPCExt2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DoRPCExt2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoRPCExt2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DoAsyncConnectExOperation structure represents the EcDoAsyncConnectEx operation
type xxx_DoAsyncConnectExOperation struct {
	Session      *oxcrpc.Session      `idl:"name:cxh" json:"session"`
	AsyncSession *oxcrpc.AsyncSession `idl:"name:pacxh;pointer:ref" json:"async_session"`
	Return       int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_DoAsyncConnectExOperation) OpNum() int { return 14 }

func (o *xxx_DoAsyncConnectExOperation) OpName() string { return "/emsmdb/v0.81/EcDoAsyncConnectEx" }

func (o *xxx_DoAsyncConnectExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoAsyncConnectExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// cxh {in} (1:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session != nil {
			if err := o.Session.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oxcrpc.Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DoAsyncConnectExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// cxh {in} (1:{context_handle, alias=CXH, names=ndr_context_handle}(struct))
	{
		if o.Session == nil {
			o.Session = &oxcrpc.Session{}
		}
		if err := o.Session.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoAsyncConnectExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoAsyncConnectExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pacxh {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ACXH, names=ndr_context_handle}(struct))
	{
		if o.AsyncSession != nil {
			if err := o.AsyncSession.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oxcrpc.AsyncSession{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoAsyncConnectExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pacxh {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=ACXH, names=ndr_context_handle}(struct))
	{
		if o.AsyncSession == nil {
			o.AsyncSession = &oxcrpc.AsyncSession{}
		}
		if err := o.AsyncSession.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DoAsyncConnectExRequest structure represents the EcDoAsyncConnectEx operation request
type DoAsyncConnectExRequest struct {
	// cxh: A session context handle. The client MUST pass a valid session context handle
	// that was created by calling the EcDoConnectEx method. The server uses the session
	// context handle to identify the Session Context to use for this call.
	Session *oxcrpc.Session `idl:"name:cxh" json:"session"`
}

func (o *DoAsyncConnectExRequest) xxx_ToOp(ctx context.Context, op *xxx_DoAsyncConnectExOperation) *xxx_DoAsyncConnectExOperation {
	if op == nil {
		op = &xxx_DoAsyncConnectExOperation{}
	}
	if o == nil {
		return op
	}
	op.Session = o.Session
	return op
}

func (o *DoAsyncConnectExRequest) xxx_FromOp(ctx context.Context, op *xxx_DoAsyncConnectExOperation) {
	if o == nil {
		return
	}
	o.Session = op.Session
}
func (o *DoAsyncConnectExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DoAsyncConnectExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoAsyncConnectExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DoAsyncConnectExResponse structure represents the EcDoAsyncConnectEx operation response
type DoAsyncConnectExResponse struct {
	// pacxh: An asynchronous context handle. On success, the server returns an asynchronous
	// context handle that is associated with the Session Context passed in the cxh parameter.
	// On failure, the returned value is NULL. The asynchronous context handle can be used
	// to make a call to the EcDoAsyncWaitEx method on the AsyncEMSMDB interface.
	AsyncSession *oxcrpc.AsyncSession `idl:"name:pacxh;pointer:ref" json:"async_session"`
	// Return: The EcDoAsyncConnectEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DoAsyncConnectExResponse) xxx_ToOp(ctx context.Context, op *xxx_DoAsyncConnectExOperation) *xxx_DoAsyncConnectExOperation {
	if op == nil {
		op = &xxx_DoAsyncConnectExOperation{}
	}
	if o == nil {
		return op
	}
	op.AsyncSession = o.AsyncSession
	op.Return = o.Return
	return op
}

func (o *DoAsyncConnectExResponse) xxx_FromOp(ctx context.Context, op *xxx_DoAsyncConnectExOperation) {
	if o == nil {
		return
	}
	o.AsyncSession = op.AsyncSession
	o.Return = op.Return
}
func (o *DoAsyncConnectExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DoAsyncConnectExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoAsyncConnectExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
