package irpcasyncnotify

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

// IRPCAsyncNotify server interface.
type IrpcAsyncNotifyServer interface {

	// The IRPCAsyncNotify_RegisterClient method is called by clients to register to receive
	// notifications and to associate the parameters describing the set of notifications
	// they are registering to receive with a remote object.
	//
	// Return Values: This method MUST return zero to indicate success, or an HRESULT error
	// value ([MS-ERREF] section 2.1.1) to indicate failure. Protocol-specific error values
	// are defined in the following table. The client SHOULD treat all error return values
	// the same, except where noted.
	//
	//	+--------------+----------------------------------------------------------------------------------+
	//	|    RETURN    |                                                                                  |
	//	|    VALUE     |                                   DESCRIPTION                                    |
	//	|              |                                                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x80070005   | The client does not have authorization to register for notifications with the    |
	//	|              | set of parameters specified in this call. If this error value is returned, the   |
	//	|              | client SHOULD NOT retry this call; the client SHOULD consider the error to be    |
	//	|              | fatal and report it to the higher level caller.                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8007000E   | The server does not have enough memory for the new registration.                 |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x80070015   | The server has reached its maximum registration limit.                           |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8007007B   | The pName parameter does not conform to the format specified above. If this      |
	//	|              | error value is returned, the client SHOULD NOT retry this call; the client       |
	//	|              | SHOULD consider the error to be fatal and report it to the higher level caller.  |
	//	+--------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// Unless specified otherwise, if a failure is indicated by an error return or an exception,
	// the client SHOULD retry this method call by performing the following steps:
	//
	// *
	//
	// Call IRPCRemoteObject_Create to generate a new PRPCREMOTEOBJECT (section 2.2.4) (
	// a47aca7c-fcc3-4151-8fb6-1de5225ecfa5 ).
	//
	// *
	//
	// Call IRPCAsyncNotify_RegisterClient with the new PRPCREMOTEOBJECT.
	//
	// Retries SHOULD be separated by time intervals decaying from 1 second to 1 minute
	// to reduce a potential burden on the server.
	//
	// The IRPCAsyncNotify_RegisterClient method MUST be called by clients to register for
	// receiving notifications. Servers MUST associate the given remote object with the
	// registration parameters specified.
	//
	// A client MUST NOT call IRPCAsyncNotify_RegisterClient if a prior call to IRPCAsyncNotify_RegisterClient
	// succeeded using the same PRPCREMOTEOBJECT value, unless a later call to IRPCAsyncNotify_UnregisterClient
	// also succeeded.
	//
	// If registering for unidirectional communication mode, a client SHOULD call IRPCAsyncNotify_GetNotification
	// after a successful call to IRPCAsyncNotify_RegisterClient using the same PRPCREMOTEOBJECT
	// value.
	//
	// If registering for bidirectional communication mode, a client SHOULD call IRPCAsyncNotify_GetNewChannel
	// after a successful call to IRPCAsyncNotify_RegisterClient using the same PRPCREMOTEOBJECT
	// value.
	//
	// Servers MUST support the concurrent registration of multiple remote objects with
	// different registration parameters, including notification type, filter, and communication
	// mode.
	//
	// Servers SHOULD consider the security and privacy context prior to letting clients
	// monitor and receive notifications for all user identities. Relevant access rights
	// are defined in the following table.
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|                               |                                                                                  |
	//	|          NAME/VALUE           |                                   DESCRIPTION                                    |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| SERVER_ALL_ACCESS 0x000F0003  | Combines the WO (Write Owner), WD (Write DACL), RC (Read Control), and DE        |
	//	|                               | (Delete) bits of the ACCESS_MASK data type ([MS-DTYP] section 2.4.3) with        |
	//	|                               | the following protocol-specific bits: 0x00000001 (bit 31): Access rights to      |
	//	|                               | administer print servers. 0x00000002 (bit 30): Access rights to enumerate print  |
	//	|                               | servers. These printing-specific access rights allow a client to administer the  |
	//	|                               | server and to enumerate server components such as print queues.                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| PRINTER_ALL_ACCESS 0x000F000C | Combines the WO (Write Owner), WD (Write DACL), RC (Read Control), and DE        |
	//	|                               | (Delete) bits of the ACCESS_MASK data type with the following protocol-specific  |
	//	|                               | bits: 0x00000004 (bit 29): Access rights for printers to perform administrative  |
	//	|                               | tasks. 0x00000008 (bit 28): Access rights for printers to perform basic printing |
	//	|                               | operations. These printing-specific access rights allow a client basic and       |
	//	|                               | administrative use of print queues.                                              |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//
	// For calls to IRPCAsyncNotify_RegisterClient with NotifyFilter set to kAllUsers, if
	// pName is set to NULL, the server SHOULD fail the call if the calling principal lacks
	// any of the server access rights specified by SERVER_ALL_ACCESS. If pName points to
	// the name of a print queue, the server SHOULD fail the call if the calling principal
	// lacks any of the print queue access rights specified by PRINTER_ALL_ACCESS. For additional
	// information concerning access rights, see [MS-AZOD] section 1.1.1.5.
	RegisterClient(context.Context, *RegisterClientRequest) (*RegisterClientResponse, error)

	// The IRPCAsyncNotify_UnregisterClient method is called by registered clients to unregister
	// previously-registered remote objects. For this call to succeed, the remote object
	// MUST have already successfully called IRPCAsyncNotify_RegisterClient.
	//
	// Return Values: This method MUST return an HRESULT success value ([MS-ERREF] section
	// 2.1.1) to indicate success, or an HRESULT error value to indicate failure. The client
	// MUST consider all error return values fatal and report them to the higher-level caller.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// If a client call to IRPCAsyncNotify_GetNewChannel or IRPCAsyncNotify_GetNotification
	// is blocked on the server waiting for a notification channel or notification to become
	// available, the server MUST process a client call to IRPCAsyncNotify_UnregisterClient
	// without waiting for the notification channel or notification.
	//
	// A server MUST NOT do the following:
	//
	// * Indicate success to a client call of IRPCAsyncNotify_UnregisterClient unless a
	// prior call to IRPCAsyncNotify_RegisterClient succeeded using the same PRPCREMOTEOBJECT
	// ( a47aca7c-fcc3-4151-8fb6-1de5225ecfa5 ) value.
	//
	// * Indicate success to a client call of IRPCAsyncNotify_UnregisterClient following
	// a prior successful call to IRPCAsyncNotify_UnregisterClient by using the same PRPCREMOTEOBJECT
	// value.
	//
	// A client MUST NOT do the following:
	//
	// * Call IRPCAsyncNotify_UnregisterClient, unless a prior call to IRPCAsyncNotify_RegisterClient
	// succeeded by using the same PRPCREMOTEOBJECT value.
	//
	// * Call IRPCAsyncNotify_UnregisterClient following a prior call to IRPCAsyncNotify_UnregisterClient
	// by using the same PRPCREMOTEOBJECT value.
	UnregisterClient(context.Context, *UnregisterClientRequest) (*UnregisterClientResponse, error)

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// The IRPCAsyncNotify_GetNewChannel method returns an array of pointers to print notification
	// channels. This method MUST only be used with bidirectional communication mode.
	//
	// Return Values: This method MUST return zero to indicate success, or an HRESULT error
	// value ([MS-ERREF] section 2.1.1) to indicate failure. Protocol-specific error values
	// are defined in the following table. The client SHOULD treat all error return values
	// the same, except where noted.
	//
	//	+--------------+----------------------------------------------------------------------------------+
	//	|    RETURN    |                                                                                  |
	//	|    VALUE     |                                   DESCRIPTION                                    |
	//	|              |                                                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8004000C   | The server has not yet returned from a previous call to this method with the     |
	//	|              | same remote object. If this error value is returned, the client SHOULD NOT retry |
	//	|              | this call before the previous call to this method with the specified remote      |
	//	|              | object has completed.                                                            |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8007000E   | The server does not have enough memory for the new channel.                      |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8007071A   | Incoming notifications have been terminated. Upon completion of this call with   |
	//	|              | this return value, the server MUST fail subsequent calls to this method with the |
	//	|              | same remote object. If this error value is returned, the client SHOULD NOT retry |
	//	|              | this call.                                                                       |
	//	+--------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: An exception code of 0x8004000C or 0x8007071A SHOULD be thrown
	// by the server under the circumstances described in the preceding table for the corresponding
	// return values. The client MUST treat these exception codes exactly as it would treat
	// the same return values. No additional exceptions are thrown beyond those thrown by
	// the underlying RPC protocol [MS-RPCE].
	//
	// Unless specified otherwise, if a failure is indicated by an error return or an exception,
	// the client SHOULD retry this method call by performing the following steps:
	//
	// *
	//
	// Call IRPCRemoteObject_Create to generate a new PRPCREMOTEOBJECT (section 2.2.4) (
	// a47aca7c-fcc3-4151-8fb6-1de5225ecfa5 ).
	//
	// *
	//
	// Call IRPCAsyncNotify_RegisterClient with the new PRPCREMOTEOBJECT.
	//
	// *
	//
	// Call IRPCAsyncNotify_GetNewChannel with the new PRPCREMOTEOBJECT.
	//
	// Retries SHOULD be separated by time intervals decaying from 1 second to 1 minute
	// to reduce a potential burden on the server. Retries SHOULD terminate when the above
	// sequence succeeds or the client determines that it is no longer interested in notifications
	// for the particular combination of notification type, print queue name, conversation
	// style, and user identity filter that were originally specified in the call to IRPCAsyncNotify_RegisterClient.
	//
	// If successful, the IRPCAsyncNotify_GetNewChannel method MUST return an array of pointers
	// to print notification channels.
	//
	// A server MUST NOT do the following:
	//
	// * Indicate success to a client call of IRPCAsyncNotify_GetNewChannel unless a prior
	// call to IRPCAsyncNotify_RegisterClient succeeded using the same PRPCREMOTEOBJECT
	// value.
	//
	// * Indicate success to a client call of IRPCAsyncNotify_GetNewChannel following a
	// prior successful call to IRPCAsyncNotify_UnregisterClient ( 20fa79b6-4905-4a50-83d5-2bc76525b3c9
	// ) using the same PRPCREMOTEOBJECT value.
	//
	// * Complete a call to IRPCAsyncNotify_GetNewChannel unless an unreturned notification
	// channel is available on the *Bidirectional Notification Channel Queue* associated
	// with the *Client Registration* (Abstract Data Model, section 3.1.1.1), or an abnormal
	// event happened, such as an initiated server shutdown sequence.
	//
	// A client SHOULD do the following:
	//
	// * Call IRPCAsyncNotify_GetNewChannel in response to a prior successful return from
	// IRPCAsyncNotify_RegisterClient or IRPCAsyncNotify_GetNewChannel.
	//
	// * Call IRPCAsyncNotify_GetNotificationSendResponse ( 8c4aab2d-5dfe-469d-a9e3-003869921e45
	// ) in response to a prior successful return from IRPCAsyncNotify_GetNewChannel.
	//
	// A client MUST NOT do the following:
	//
	// * Call IRPCAsyncNotify_GetNewChannel, unless a prior call to IRPCAsyncNotify_RegisterClient
	// succeeded by using the same PRPCREMOTEOBJECT value. <8> ( cdec7a7c-2b2c-4b81-a40a-b12d69f880ee#Appendix_A_8
	// )
	//
	// * Call IRPCAsyncNotify_GetNewChannel following a prior call to IRPCAsyncNotify_UnregisterClient
	// by using the same PRPCREMOTEOBJECT value. <9> ( cdec7a7c-2b2c-4b81-a40a-b12d69f880ee#Appendix_A_9
	// )
	GetNewChannel(context.Context, *GetNewChannelRequest) (*GetNewChannelResponse, error)

	// The IRPCAsyncNotify_GetNotificationSendResponse method sends a client response to
	// the server and returns the next notification sent by way of the same channel when
	// it becomes available. This method MUST be used only with bidirectional communication
	// mode.
	//
	// Return Values: This method MUST return zero to indicate success, or an HRESULT error
	// value ([MS-ERREF] section 2.1.1) to indicate failure. Protocol-specific error values
	// are defined in the following table. The client MUST consider all error return values
	// fatal and report them to the higher-level caller.
	//
	//	+--------------+----------------------------------------------------------------------------------+
	//	|    RETURN    |                                                                                  |
	//	|    VALUE     |                                   DESCRIPTION                                    |
	//	|              |                                                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x80040008   | The notification channel represented by pChannel was previously closed.          |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8004000C   | The server has not yet returned from a previous call to                          |
	//	|              | IRPCAsyncNotify_GetNotificationSendResponse or IRPCAsyncNotify_CloseChannel      |
	//	|              | (section 3.1.1.4.6) with the same notification channel.                          |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x80040012   | The size of the client-to-server response exceeded the maximum size.             |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x80040014   | The notification type identifier is different from the notification type of the  |
	//	|              | notification channel.                                                            |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8007000E   | The server does not have enough memory to complete the request.                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// If a failure is indicated by an error return or an exception, the client SHOULD close
	// the channel.
	//
	// The first call to this method on the newly opened notification channel serves as
	// a mediator among all the clients that registered themselves for the given notification
	// type. This MUST be done by blocking all calls from clients until a matching server-side
	// event occurs, including the following:
	//
	// * The channel issues a notification.
	//
	// * An abnormal condition occurs, such as an initiated server shutdown sequence.
	//
	// * The server receives a client request to close the channel.
	GetNotificationSendResponse(context.Context, *GetNotificationSendResponseRequest) (*GetNotificationSendResponseResponse, error)

	// The IRPCAsyncNotify_GetNotification method returns notification data from the print
	// server. This method MUST NOT be used with bidirectional communication mode.
	//
	// Return Values: This method MUST return zero to indicate success, or an HRESULT error
	// value ([MS-ERREF] section 2.1.1) to indicate failure. Protocol-specific error values
	// are defined in the following table. The client SHOULD treat all error return values
	// the same, except where noted.
	//
	//	+--------------+----------------------------------------------------------------------------------+
	//	|    RETURN    |                                                                                  |
	//	|    VALUE     |                                   DESCRIPTION                                    |
	//	|              |                                                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8004000C   | The server has not yet returned from a previous call to this method with the     |
	//	|              | same remote object. If this error value is returned, the client SHOULD NOT retry |
	//	|              | this call before the previous call to this method with the specified remote      |
	//	|              | object has completed.                                                            |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8007000E   | The server does not have enough memory to complete the request.                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8007071A   | Incoming notifications have been terminated. Upon completion of this call with   |
	//	|              | this return value, the server MUST fail subsequent calls to this method with the |
	//	|              | same remote object. If this error value is returned, the client SHOULD NOT retry |
	//	|              | this call.                                                                       |
	//	+--------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: An exception code of 0x08004000C or 0x8007071A SHOULD be thrown
	// by the server under the circumstances described in the preceding table for the corresponding
	// return values. The client MUST treat these exception codes exactly as it would treat
	// the same return values. No additional exceptions are thrown beyond those thrown by
	// the underlying RPC protocol [MS-RPCE].
	//
	// Unless specified otherwise, if a failure is indicated by an error return or an exception,
	// the client SHOULD retry this method call by performing the following steps:
	//
	// *
	//
	// Call IRPCRemoteObject_Create (section 3.1.2.4.1) ( e3786f60-0b93-4c5e-8cd1-3f0487e4310a
	// ) to generate a new PRPCREMOTEOBJECT (section 2.2.4) ( a47aca7c-fcc3-4151-8fb6-1de5225ecfa5
	// ).
	//
	// *
	//
	// Call IRPCAsyncNotify_RegisterClient with the new PRPCREMOTEOBJECT.
	//
	// *
	//
	// Call IRPCAsyncNotify_GetNotification with the new PRPCREMOTEOBJECT.
	//
	// Retries SHOULD be separated by time intervals decaying from 1 second to 1 minute
	// to reduce a potential burden on the server. Retries SHOULD terminate when the above
	// sequence succeeds or the client determines that it is no longer interested in notifications
	// for the particular combination of notification type, print queue name, conversation
	// style, and user identity filter that were originally specified in the call to IRPCAsyncNotify_RegisterClient.
	//
	// The IRPCAsyncNotify_GetNotification method MUST return data from the server that
	// matches the registration for the given remote object.
	//
	// A server MUST NOT do the following:
	//
	// * Indicate success to a client call of IRPCAsyncNotify_GetNotification unless a prior
	// call to IRPCAsyncNotify_RegisterClient succeeded using the same PRPCREMOTEOBJECT
	// value.
	//
	// * Indicate success to a client call of IRPCAsyncNotify_GetNotification following
	// a prior successful call to IRPCAsyncNotify_UnregisterClient ( 20fa79b6-4905-4a50-83d5-2bc76525b3c9
	// ) using the same PRPCREMOTEOBJECT value.
	//
	// * Complete a call to IRPCAsyncNotify_GetNotification until the *Unidirectional Notification
	// Queue* associated with the *Client Registration* ( Abstract Data Model (section 3.1.1.1)
	// ( bd69159c-f3d8-4f7b-b2c3-f9efec7c9f62 ) ) contains an unreturned notification, or
	// an abnormal condition occurs. An example of an abnormal condition is an initiated
	// server shutdown sequence or remote object unregistration. An abnormal condition will
	// result in a failure error code returned prior to the server having data.
	//
	// A server SHOULD do the following:
	//
	// * Discard unidirectional notifications in the absence of corresponding registered
	// clients.
	//
	// * Buffer unidirectional notifications, up to some implementation-defined limit, <10>
	// ( cdec7a7c-2b2c-4b81-a40a-b12d69f880ee#Appendix_A_10 ) for each registered client
	// that does not have pending IRPCAsyncNotify_GetNotification calls.
	//
	// If a client wants to receive further notifications from the server, the client SHOULD
	// call IRPCAsyncNotify_GetNotification in response to a prior successful return from
	// IRPCAsyncNotify_GetNotification. When the client no longer wants to receive notifications
	// from the server, it SHOULD call IRPCAsyncNotify_UnregisterClient, either before or
	// after the return from IRPCAsyncNotify_GetNotification.
	//
	// A client MUST NOT do the following:
	//
	// * Call IRPCAsyncNotify_GetNotification unless a prior call to IRPCAsyncNotify_RegisterClient
	// succeeded, using the same PRPCREMOTEOBJECT value.
	//
	// * Call IRPCAsyncNotify_GetNotification following a prior call to IRPCAsyncNotify_UnregisterClient
	// by using the same PRPCREMOTEOBJECT value.
	GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationResponse, error)

	// The IRPCAsyncNotify_CloseChannel method sends a final response on the notification
	// channel and closes it. This method MUST NOT be used with unidirectional communication
	// mode.
	//
	// Return Values: This method MUST return zero or an HRESULT success value ([MS-ERREF]
	// section 2.1.1) to indicate success, or an HRESULT error value to indicate failure.
	//
	// Protocol-specific success values are defined in the following table.
	//
	//	+--------------+------------------------------------------+
	//	|    RETURN    |                                          |
	//	|    VALUE     |               DESCRIPTION                |
	//	|              |                                          |
	//	+--------------+------------------------------------------+
	//	+--------------+------------------------------------------+
	//	| 0x00040010   | Another client has acquired the channel. |
	//	+--------------+------------------------------------------+
	//
	// Protocol-specific error values are defined in the following table. The client MUST
	// consider all error return values fatal and report them to the higher-level caller.
	//
	//	+--------------+----------------------------------------------------------------------------------+
	//	|    RETURN    |                                                                                  |
	//	|    VALUE     |                                   DESCRIPTION                                    |
	//	|              |                                                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x80040012   | The response exceeds the maximum size allowed by the server.                     |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x80040014   | The notification type identifier is different from the notification type of the  |
	//	|              | notification channel.                                                            |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x8007000E   | The server does not have enough memory to complete the request.                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// If a client call to IRPCAsyncNotify_GetNotificationSendResponse is blocked on the
	// server, waiting for a notification to become available on a notification channel,
	// then the server MUST process a client call to this method on the same notification
	// channel without waiting for a notification.
	//
	// A client MUST NOT call IRPCAsyncNotify_CloseChannel following a prior successful
	// return from IRPCAsyncNotify_GetNotificationSendResponse with a NULL value of pChannel
	// parameter or following a prior successful return from IRPCAsyncNotify_CloseChannel.<11>
	CloseChannel(context.Context, *CloseChannelRequest) (*CloseChannelResponse, error)
}

func RegisterIrpcAsyncNotifyServer(conn dcerpc.Conn, o IrpcAsyncNotifyServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIrpcAsyncNotifyServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IrpcAsyncNotifySyntaxV1_0))...)
}

func NewIrpcAsyncNotifyServerHandle(o IrpcAsyncNotifyServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IrpcAsyncNotifyServerHandle(ctx, o, opNum, r)
	}
}

func IrpcAsyncNotifyServerHandle(ctx context.Context, o IrpcAsyncNotifyServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // IRPCAsyncNotify_RegisterClient
		in := &RegisterClientRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RegisterClient(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // IRPCAsyncNotify_UnregisterClient
		in := &UnregisterClientRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.UnregisterClient(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // Opnum2NotUsedOnWire
		// Opnum2NotUsedOnWire
		return nil, nil
	case 3: // IRPCAsyncNotify_GetNewChannel
		in := &GetNewChannelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNewChannel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // IRPCAsyncNotify_GetNotificationSendResponse
		in := &GetNotificationSendResponseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNotificationSendResponse(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // IRPCAsyncNotify_GetNotification
		in := &GetNotificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNotification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // IRPCAsyncNotify_CloseChannel
		in := &CloseChannelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseChannel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
