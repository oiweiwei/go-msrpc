package irpcasyncnotify

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	irpcremoteobject "github.com/oiweiwei/go-msrpc/msrpc/pan/irpcremoteobject/v1"
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
	_ = dtyp.GoPackage
	_ = dcetypes.GoPackage
	_ = irpcremoteobject.GoPackage
)

var (
	// import guard
	GoPackage = "pan"
)

var (
	// Syntax UUID
	IrpcAsyncNotifySyntaxUUID = &uuid.UUID{TimeLow: 0xb6edbfa, TimeMid: 0x4a24, TimeHiAndVersion: 0x4fc6, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x23, Node: [6]uint8{0x94, 0x2b, 0x1e, 0xca, 0x65, 0xd1}}
	// Syntax ID
	IrpcAsyncNotifySyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: IrpcAsyncNotifySyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// IRPCAsyncNotify interface.
type IrpcAsyncNotifyClient interface {

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
	RegisterClient(context.Context, *RegisterClientRequest, ...dcerpc.CallOption) (*RegisterClientResponse, error)

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
	UnregisterClient(context.Context, *UnregisterClientRequest, ...dcerpc.CallOption) (*UnregisterClientResponse, error)

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
	GetNewChannel(context.Context, *GetNewChannelRequest, ...dcerpc.CallOption) (*GetNewChannelResponse, error)

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
	GetNotificationSendResponse(context.Context, *GetNotificationSendResponseRequest, ...dcerpc.CallOption) (*GetNotificationSendResponseResponse, error)

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
	GetNotification(context.Context, *GetNotificationRequest, ...dcerpc.CallOption) (*GetNotificationResponse, error)

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
	CloseChannel(context.Context, *CloseChannelRequest, ...dcerpc.CallOption) (*CloseChannelResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// ConversationStyle type represents PrintAsyncNotifyConversationStyle RPC enumeration.
//
// The PrintAsyncNotifyConversationStyle enumeration MUST specify the communication
// mode expected between the sender and a registered client.
type ConversationStyle uint32

var (
	// kBiDirectional:  Bidirectional communication mode is specified. The sender expects
	// the client to send responses to notifications.
	ConversationStyleBidirectional ConversationStyle = 0
	// kUniDirectional:  Unidirectional communication mode is specified. The sender does
	// not expect the client to respond to notifications.
	ConversationStyleUnidirectional ConversationStyle = 1
)

func (o ConversationStyle) String() string {
	switch o {
	case ConversationStyleBidirectional:
		return "ConversationStyleBidirectional"
	case ConversationStyleUnidirectional:
		return "ConversationStyleUnidirectional"
	}
	return "Invalid"
}

// UserFilter type represents PrintAsyncNotifyUserFilter RPC enumeration.
//
// The PrintAsyncNotifyUserFilter enumeration is used by clients when they register
// to receive notifications from server-resident notification sources. The following
// types of notifications can be requested:
//
// * Notifications intended specifically for a particular client's user identity.
//
// * Notifications intended for all registered client user identities.
//
// typedef [v1_enum] enum
//
// {
//
// kPerUser = 0,
//
// kAllUsers = 1,
//
// } PrintAsyncNotifyUserFilter;
type UserFilter uint32

var (
	// kPerUser:  Indicates that the client is requesting notifications that are intended
	// specifically for its own user identity and notifications that are intended for all
	// registered user identities.
	UserFilterPerUser UserFilter = 0
	// kAllUsers:  Indicates that the client is requesting every notification, whether
	// intended for a specific user identity or for all registered user identities.
	UserFilterAllUsers UserFilter = 1
)

func (o UserFilter) String() string {
	switch o {
	case UserFilterPerUser:
		return "UserFilterPerUser"
	case UserFilterAllUsers:
		return "UserFilterAllUsers"
	}
	return "Invalid"
}

// NotificationType structure represents PrintAsyncNotificationType RPC structure.
type NotificationType dtyp.GUID

func (o *NotificationType) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *NotificationType) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NotificationType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Data1); err != nil {
		return err
	}
	if err := w.WriteData(o.Data2); err != nil {
		return err
	}
	if err := w.WriteData(o.Data3); err != nil {
		return err
	}
	for i1 := range o.Data4 {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Data4[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data4); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *NotificationType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data3); err != nil {
		return err
	}
	o.Data4 = make([]byte, 8)
	for i1 := range o.Data4 {
		i1 := i1
		if err := w.ReadData(&o.Data4[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// NotifyObject structure represents PNOTIFYOBJECT RPC structure.
type NotifyObject dcetypes.ContextHandle

func (o *NotifyObject) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *NotifyObject) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NotifyObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NotifyObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultIrpcAsyncNotifyClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultIrpcAsyncNotifyClient) RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...dcerpc.CallOption) (*RegisterClientResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterClientResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIrpcAsyncNotifyClient) UnregisterClient(ctx context.Context, in *UnregisterClientRequest, opts ...dcerpc.CallOption) (*UnregisterClientResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UnregisterClientResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIrpcAsyncNotifyClient) GetNewChannel(ctx context.Context, in *GetNewChannelRequest, opts ...dcerpc.CallOption) (*GetNewChannelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNewChannelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIrpcAsyncNotifyClient) GetNotificationSendResponse(ctx context.Context, in *GetNotificationSendResponseRequest, opts ...dcerpc.CallOption) (*GetNotificationSendResponseResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNotificationSendResponseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIrpcAsyncNotifyClient) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...dcerpc.CallOption) (*GetNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIrpcAsyncNotifyClient) CloseChannel(ctx context.Context, in *CloseChannelRequest, opts ...dcerpc.CallOption) (*CloseChannelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseChannelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIrpcAsyncNotifyClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIrpcAsyncNotifyClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewIrpcAsyncNotifyClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IrpcAsyncNotifyClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IrpcAsyncNotifySyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultIrpcAsyncNotifyClient{cc: cc}, nil
}

// xxx_RegisterClientOperation structure represents the IRPCAsyncNotify_RegisterClient operation
type xxx_RegisterClientOperation struct {
	RegistrationObject   *irpcremoteobject.RemoteObject `idl:"name:pRegistrationObj" json:"registration_object"`
	Name                 string                         `idl:"name:pName;string;pointer:unique" json:"name"`
	InNotificationType   *NotificationType              `idl:"name:pInNotificationType" json:"in_notification_type"`
	NotifyFilter         UserFilter                     `idl:"name:NotifyFilter" json:"notify_filter"`
	ConversationStyle    ConversationStyle              `idl:"name:conversationStyle" json:"conversation_style"`
	RemoteServerReferral string                         `idl:"name:ppRmtServerReferral;string" json:"remote_server_referral"`
	Return               int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterClientOperation) OpNum() int { return 0 }

func (o *xxx_RegisterClientOperation) OpName() string {
	return "/IRPCAsyncNotify/v1/IRPCAsyncNotify_RegisterClient"
}

func (o *xxx_RegisterClientOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterClientOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pRegistrationObj {in} (1:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RegistrationObject != nil {
			if err := o.RegistrationObject.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&irpcremoteobject.RemoteObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.Name != "" {
			_ptr_pName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_pName); err != nil {
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
	// pInNotificationType {in} (1:{pointer=ref}*(1))(2:{alias=PrintAsyncNotificationType, names=GUID}(struct))
	{
		if o.InNotificationType != nil {
			if err := o.InNotificationType.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NotificationType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// NotifyFilter {in} (1:{v1_enum, alias=PrintAsyncNotifyUserFilter}(enum))
	{
		if err := w.WriteEnum(uint32(o.NotifyFilter)); err != nil {
			return err
		}
	}
	// conversationStyle {in} (1:{v1_enum, alias=PrintAsyncNotifyConversationStyle}(enum))
	{
		if err := w.WriteEnum(uint32(o.ConversationStyle)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterClientOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pRegistrationObj {in} (1:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RegistrationObject == nil {
			o.RegistrationObject = &irpcremoteobject.RemoteObject{}
		}
		if err := o.RegistrationObject.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
				return err
			}
			return nil
		})
		_s_pName := func(ptr interface{}) { o.Name = *ptr.(*string) }
		if err := w.ReadPointer(&o.Name, _s_pName, _ptr_pName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pInNotificationType {in} (1:{pointer=ref}*(1))(2:{alias=PrintAsyncNotificationType, names=GUID}(struct))
	{
		if o.InNotificationType == nil {
			o.InNotificationType = &NotificationType{}
		}
		if err := o.InNotificationType.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// NotifyFilter {in} (1:{v1_enum, alias=PrintAsyncNotifyUserFilter}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.NotifyFilter)); err != nil {
			return err
		}
	}
	// conversationStyle {in} (1:{v1_enum, alias=PrintAsyncNotifyConversationStyle}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.ConversationStyle)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterClientOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterClientOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppRmtServerReferral {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.RemoteServerReferral != "" {
			_ptr_ppRmtServerReferral := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.RemoteServerReferral); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.RemoteServerReferral, _ptr_ppRmtServerReferral); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterClientOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppRmtServerReferral {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppRmtServerReferral := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteServerReferral); err != nil {
				return err
			}
			return nil
		})
		_s_ppRmtServerReferral := func(ptr interface{}) { o.RemoteServerReferral = *ptr.(*string) }
		if err := w.ReadPointer(&o.RemoteServerReferral, _s_ppRmtServerReferral, _ptr_ppRmtServerReferral); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RegisterClientRequest structure represents the IRPCAsyncNotify_RegisterClient operation request
type RegisterClientRequest struct {
	// pRegistrationObj: MUST be the remote object context handle that was returned by the
	// server in the ppRemoteObject output parameter of a prior call to IRPCRemoteObject_Create
	// (section 3.1.2.4.1). This value MUST NOT be NULL.
	RegistrationObject *irpcremoteobject.RemoteObject `idl:"name:pRegistrationObj" json:"registration_object"`
	// pName: MUST be NULL or a pointer to a NULL-terminated string, encoded in Unicode
	// UTF-16LE ([RFC2781] section 4.2), which specifies the full UNC name of the print
	// queue from which the print client is registering to receive notifications.
	//
	// This UNC name MUST be in the following format:
	//
	// SERVER_NAME is a DNS, NetBIOS, IPv4, or IPv6 host name.
	//
	// LOCAL_PRINTER_NAME is a string that MUST NOT contain the characters "\" or ",".
	//
	// DNS names are specified in [RFC819] section 2, and NetBIOS names are specified in
	// [RFC1001] section 14. Basic notational conventions are specified in [RFC2616] section
	// 2, and "host" is defined in [RFC3986] section 3.2.2.
	//
	// If pName is NULL, the registration MUST be made for the remote print server itself.
	Name string `idl:"name:pName;string;pointer:unique" json:"name"`
	// pInNotificationType: MUST be a pointer to a PrintAsyncNotificationType structure
	// that specifies the notification type identifier for the notifications that the client
	// is registering to receive.
	InNotificationType *NotificationType `idl:"name:pInNotificationType" json:"in_notification_type"`
	// NotifyFilter: MUST be a value of type PrintAsyncNotifyUserFilter that specifies whether
	// the client is registering to receive notifications that are issued to all registered
	// clients, irrespective of their authenticated user identity, or to receive notifications
	// that are issued only to the specific authenticated user identity of the registering
	// RPC client.
	NotifyFilter UserFilter `idl:"name:NotifyFilter" json:"notify_filter"`
	// conversationStyle: MUST be a value of type PrintAsyncNotifyConversationStyle that
	// specifies whether the client is registering for bidirectional communication mode
	// or unidirectional communication mode.
	ConversationStyle ConversationStyle `idl:"name:conversationStyle" json:"conversation_style"`
}

func (o *RegisterClientRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterClientOperation) *xxx_RegisterClientOperation {
	if op == nil {
		op = &xxx_RegisterClientOperation{}
	}
	if o == nil {
		return op
	}
	op.RegistrationObject = o.RegistrationObject
	op.Name = o.Name
	op.InNotificationType = o.InNotificationType
	op.NotifyFilter = o.NotifyFilter
	op.ConversationStyle = o.ConversationStyle
	return op
}

func (o *RegisterClientRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterClientOperation) {
	if o == nil {
		return
	}
	o.RegistrationObject = op.RegistrationObject
	o.Name = op.Name
	o.InNotificationType = op.InNotificationType
	o.NotifyFilter = op.NotifyFilter
	o.ConversationStyle = op.ConversationStyle
}
func (o *RegisterClientRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterClientRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterClientOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterClientResponse structure represents the IRPCAsyncNotify_RegisterClient operation response
type RegisterClientResponse struct {
	// ppRmtServerReferral: Servers SHOULD return NULL for this parameter, and clients MUST
	// ignore it on receipt.
	RemoteServerReferral string `idl:"name:ppRmtServerReferral;string" json:"remote_server_referral"`
	// Return: The IRPCAsyncNotify_RegisterClient return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterClientResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterClientOperation) *xxx_RegisterClientOperation {
	if op == nil {
		op = &xxx_RegisterClientOperation{}
	}
	if o == nil {
		return op
	}
	op.RemoteServerReferral = o.RemoteServerReferral
	op.Return = o.Return
	return op
}

func (o *RegisterClientResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterClientOperation) {
	if o == nil {
		return
	}
	o.RemoteServerReferral = op.RemoteServerReferral
	o.Return = op.Return
}
func (o *RegisterClientResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterClientResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterClientOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnregisterClientOperation structure represents the IRPCAsyncNotify_UnregisterClient operation
type xxx_UnregisterClientOperation struct {
	RegistrationObject *irpcremoteobject.RemoteObject `idl:"name:pRegistrationObj" json:"registration_object"`
	Return             int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_UnregisterClientOperation) OpNum() int { return 1 }

func (o *xxx_UnregisterClientOperation) OpName() string {
	return "/IRPCAsyncNotify/v1/IRPCAsyncNotify_UnregisterClient"
}

func (o *xxx_UnregisterClientOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterClientOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pRegistrationObj {in} (1:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RegistrationObject != nil {
			if err := o.RegistrationObject.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&irpcremoteobject.RemoteObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_UnregisterClientOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pRegistrationObj {in} (1:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RegistrationObject == nil {
			o.RegistrationObject = &irpcremoteobject.RemoteObject{}
		}
		if err := o.RegistrationObject.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterClientOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterClientOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterClientOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UnregisterClientRequest structure represents the IRPCAsyncNotify_UnregisterClient operation request
type UnregisterClientRequest struct {
	// pRegistrationObj: MUST be the remote object context handle that MUST have been successfully
	// registered by a prior call to IRPCAsyncNotify_RegisterClient. This value MUST NOT
	// be NULL.
	RegistrationObject *irpcremoteobject.RemoteObject `idl:"name:pRegistrationObj" json:"registration_object"`
}

func (o *UnregisterClientRequest) xxx_ToOp(ctx context.Context, op *xxx_UnregisterClientOperation) *xxx_UnregisterClientOperation {
	if op == nil {
		op = &xxx_UnregisterClientOperation{}
	}
	if o == nil {
		return op
	}
	op.RegistrationObject = o.RegistrationObject
	return op
}

func (o *UnregisterClientRequest) xxx_FromOp(ctx context.Context, op *xxx_UnregisterClientOperation) {
	if o == nil {
		return
	}
	o.RegistrationObject = op.RegistrationObject
}
func (o *UnregisterClientRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UnregisterClientRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnregisterClientOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UnregisterClientResponse structure represents the IRPCAsyncNotify_UnregisterClient operation response
type UnregisterClientResponse struct {
	// Return: The IRPCAsyncNotify_UnregisterClient return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UnregisterClientResponse) xxx_ToOp(ctx context.Context, op *xxx_UnregisterClientOperation) *xxx_UnregisterClientOperation {
	if op == nil {
		op = &xxx_UnregisterClientOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *UnregisterClientResponse) xxx_FromOp(ctx context.Context, op *xxx_UnregisterClientOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UnregisterClientResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UnregisterClientResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnregisterClientOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNewChannelOperation structure represents the IRPCAsyncNotify_GetNewChannel operation
type xxx_GetNewChannelOperation struct {
	RemoteObject     *irpcremoteobject.RemoteObject `idl:"name:pRemoteObj" json:"remote_object"`
	NumberOfChannels uint32                         `idl:"name:pNoOfChannels" json:"number_of_channels"`
	ChannelContext   []*NotifyObject                `idl:"name:ppChannelCtxt;size_is:(, pNoOfChannels)" json:"channel_context"`
	Return           int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNewChannelOperation) OpNum() int { return 3 }

func (o *xxx_GetNewChannelOperation) OpName() string {
	return "/IRPCAsyncNotify/v1/IRPCAsyncNotify_GetNewChannel"
}

func (o *xxx_GetNewChannelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNewChannelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pRemoteObj {in} (1:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RemoteObject != nil {
			if err := o.RemoteObject.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&irpcremoteobject.RemoteObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetNewChannelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pRemoteObj {in} (1:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RemoteObject == nil {
			o.RemoteObject = &irpcremoteobject.RemoteObject{}
		}
		if err := o.RemoteObject.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNewChannelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ChannelContext != nil && o.NumberOfChannels == 0 {
		o.NumberOfChannels = uint32(len(o.ChannelContext))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNewChannelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pNoOfChannels {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.NumberOfChannels); err != nil {
			return err
		}
	}
	// ppChannelCtxt {out} (1:{pointer=ref}*(2)*(1))(2:{context_handle, alias=PNOTIFYOBJECT, names=ndr_context_handle}[dim:0,size_is=pNoOfChannels](struct))
	{
		if o.ChannelContext != nil || o.NumberOfChannels > 0 {
			_ptr_ppChannelCtxt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfChannels)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ChannelContext {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ChannelContext[i1] != nil {
						if err := o.ChannelContext[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&NotifyObject{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ChannelContext); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&NotifyObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ChannelContext, _ptr_ppChannelCtxt); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNewChannelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pNoOfChannels {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.NumberOfChannels); err != nil {
			return err
		}
	}
	// ppChannelCtxt {out} (1:{pointer=ref}*(2)*(1))(2:{context_handle, alias=PNOTIFYOBJECT, names=ndr_context_handle}[dim:0,size_is=pNoOfChannels](struct))
	{
		_ptr_ppChannelCtxt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ChannelContext", sizeInfo[0])
			}
			o.ChannelContext = make([]*NotifyObject, sizeInfo[0])
			for i1 := range o.ChannelContext {
				i1 := i1
				if o.ChannelContext[i1] == nil {
					o.ChannelContext[i1] = &NotifyObject{}
				}
				if err := o.ChannelContext[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppChannelCtxt := func(ptr interface{}) { o.ChannelContext = *ptr.(*[]*NotifyObject) }
		if err := w.ReadPointer(&o.ChannelContext, _s_ppChannelCtxt, _ptr_ppChannelCtxt); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNewChannelRequest structure represents the IRPCAsyncNotify_GetNewChannel operation request
type GetNewChannelRequest struct {
	// pRemoteObj: MUST be the remote object context handle. This handle is obtained from
	// IRPCRemoteObject_Create (section 3.1.2.4.1). This remote object MUST have been registered
	// for bidirectional communication mode by a prior successful call to IRPCAsyncNotify_RegisterClient
	// (section 3.1.1.4.1).
	RemoteObject *irpcremoteobject.RemoteObject `idl:"name:pRemoteObj" json:"remote_object"`
}

func (o *GetNewChannelRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNewChannelOperation) *xxx_GetNewChannelOperation {
	if op == nil {
		op = &xxx_GetNewChannelOperation{}
	}
	if o == nil {
		return op
	}
	op.RemoteObject = o.RemoteObject
	return op
}

func (o *GetNewChannelRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNewChannelOperation) {
	if o == nil {
		return
	}
	o.RemoteObject = op.RemoteObject
}
func (o *GetNewChannelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNewChannelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNewChannelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNewChannelResponse structure represents the IRPCAsyncNotify_GetNewChannel operation response
type GetNewChannelResponse struct {
	// pNoOfChannels: MUST specify the number of notification channels returned. The array
	// of notification channels is specified by the ppChannelCtxt parameter.
	NumberOfChannels uint32 `idl:"name:pNoOfChannels" json:"number_of_channels"`
	// ppChannelCtxt: MUST specify a pointer to the array of returned notification channels.
	// This data is represented by a Bidirectional Notification Channel structure in the
	// Abstract Data Model (section 3.1.1.1).
	ChannelContext []*NotifyObject `idl:"name:ppChannelCtxt;size_is:(, pNoOfChannels)" json:"channel_context"`
	// Return: The IRPCAsyncNotify_GetNewChannel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNewChannelResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNewChannelOperation) *xxx_GetNewChannelOperation {
	if op == nil {
		op = &xxx_GetNewChannelOperation{}
	}
	if o == nil {
		return op
	}
	op.NumberOfChannels = o.NumberOfChannels
	op.ChannelContext = o.ChannelContext
	op.Return = o.Return
	return op
}

func (o *GetNewChannelResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNewChannelOperation) {
	if o == nil {
		return
	}
	o.NumberOfChannels = op.NumberOfChannels
	o.ChannelContext = op.ChannelContext
	o.Return = op.Return
}
func (o *GetNewChannelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNewChannelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNewChannelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNotificationSendResponseOperation structure represents the IRPCAsyncNotify_GetNotificationSendResponse operation
type xxx_GetNotificationSendResponseOperation struct {
	Channel             *NotifyObject     `idl:"name:pChannel" json:"channel"`
	InNotificationType  *NotificationType `idl:"name:pInNotificationType;pointer:unique" json:"in_notification_type"`
	InSize              uint32            `idl:"name:InSize" json:"in_size"`
	InNotificationData  []byte            `idl:"name:pInNotificationData;size_is:(InSize);pointer:unique" json:"in_notification_data"`
	OutNotificationType *NotificationType `idl:"name:ppOutNotificationType" json:"out_notification_type"`
	OutSize             uint32            `idl:"name:pOutSize" json:"out_size"`
	OutNotificationData []byte            `idl:"name:ppOutNotificationData;size_is:(, pOutSize)" json:"out_notification_data"`
	Return              int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNotificationSendResponseOperation) OpNum() int { return 4 }

func (o *xxx_GetNotificationSendResponseOperation) OpName() string {
	return "/IRPCAsyncNotify/v1/IRPCAsyncNotify_GetNotificationSendResponse"
}

func (o *xxx_GetNotificationSendResponseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InNotificationData != nil && o.InSize == 0 {
		o.InSize = uint32(len(o.InNotificationData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotificationSendResponseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pChannel {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PNOTIFYOBJECT, names=ndr_context_handle}(struct))
	{
		if o.Channel != nil {
			if err := o.Channel.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NotifyObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pInNotificationType {in} (1:{pointer=unique}*(1))(2:{alias=PrintAsyncNotificationType, names=GUID}(struct))
	{
		if o.InNotificationType != nil {
			_ptr_pInNotificationType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InNotificationType != nil {
					if err := o.InNotificationType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NotificationType{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InNotificationType, _ptr_pInNotificationType); err != nil {
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
	// InSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.InSize); err != nil {
			return err
		}
	}
	// pInNotificationData {in} (1:{pointer=unique}*(1)[dim:0,size_is=InSize](uint8))
	{
		if o.InNotificationData != nil || o.InSize > 0 {
			_ptr_pInNotificationData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InNotificationData {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InNotificationData[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InNotificationData); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InNotificationData, _ptr_pInNotificationData); err != nil {
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
	return nil
}

func (o *xxx_GetNotificationSendResponseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pChannel {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PNOTIFYOBJECT, names=ndr_context_handle}(struct))
	{
		if o.Channel == nil {
			o.Channel = &NotifyObject{}
		}
		if err := o.Channel.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pInNotificationType {in} (1:{pointer=unique}*(1))(2:{alias=PrintAsyncNotificationType, names=GUID}(struct))
	{
		_ptr_pInNotificationType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InNotificationType == nil {
				o.InNotificationType = &NotificationType{}
			}
			if err := o.InNotificationType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInNotificationType := func(ptr interface{}) { o.InNotificationType = *ptr.(**NotificationType) }
		if err := w.ReadPointer(&o.InNotificationType, _s_pInNotificationType, _ptr_pInNotificationType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// InSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.InSize); err != nil {
			return err
		}
	}
	// pInNotificationData {in} (1:{pointer=unique}*(1)[dim:0,size_is=InSize](uint8))
	{
		_ptr_pInNotificationData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InNotificationData", sizeInfo[0])
			}
			o.InNotificationData = make([]byte, sizeInfo[0])
			for i1 := range o.InNotificationData {
				i1 := i1
				if err := w.ReadData(&o.InNotificationData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pInNotificationData := func(ptr interface{}) { o.InNotificationData = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InNotificationData, _s_pInNotificationData, _ptr_pInNotificationData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotificationSendResponseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutNotificationData != nil && o.OutSize == 0 {
		o.OutSize = uint32(len(o.OutNotificationData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotificationSendResponseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pChannel {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PNOTIFYOBJECT, names=ndr_context_handle}(struct))
	{
		if o.Channel != nil {
			if err := o.Channel.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NotifyObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppOutNotificationType {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PrintAsyncNotificationType, names=GUID}(struct))
	{
		if o.OutNotificationType != nil {
			_ptr_ppOutNotificationType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OutNotificationType != nil {
					if err := o.OutNotificationType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NotificationType{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutNotificationType, _ptr_ppOutNotificationType); err != nil {
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
	// pOutSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.OutSize); err != nil {
			return err
		}
	}
	// ppOutNotificationData {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pOutSize](uint8))
	{
		if o.OutNotificationData != nil || o.OutSize > 0 {
			_ptr_ppOutNotificationData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.OutSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.OutNotificationData {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.OutNotificationData[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.OutNotificationData); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutNotificationData, _ptr_ppOutNotificationData); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotificationSendResponseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pChannel {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PNOTIFYOBJECT, names=ndr_context_handle}(struct))
	{
		if o.Channel == nil {
			o.Channel = &NotifyObject{}
		}
		if err := o.Channel.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppOutNotificationType {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PrintAsyncNotificationType, names=GUID}(struct))
	{
		_ptr_ppOutNotificationType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OutNotificationType == nil {
				o.OutNotificationType = &NotificationType{}
			}
			if err := o.OutNotificationType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppOutNotificationType := func(ptr interface{}) { o.OutNotificationType = *ptr.(**NotificationType) }
		if err := w.ReadPointer(&o.OutNotificationType, _s_ppOutNotificationType, _ptr_ppOutNotificationType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pOutSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.OutSize); err != nil {
			return err
		}
	}
	// ppOutNotificationData {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pOutSize](uint8))
	{
		_ptr_ppOutNotificationData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.OutNotificationData", sizeInfo[0])
			}
			o.OutNotificationData = make([]byte, sizeInfo[0])
			for i1 := range o.OutNotificationData {
				i1 := i1
				if err := w.ReadData(&o.OutNotificationData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppOutNotificationData := func(ptr interface{}) { o.OutNotificationData = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.OutNotificationData, _s_ppOutNotificationData, _ptr_ppOutNotificationData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNotificationSendResponseRequest structure represents the IRPCAsyncNotify_GetNotificationSendResponse operation request
type GetNotificationSendResponseRequest struct {
	// pChannel: A pointer to a notification channel that MUST NOT be closed or zero, and
	// which MUST have been returned by the server in the ppChannelCtxt output parameter
	// of a prior call to IRPCAsyncNotify_GetNewChannel. If the server closes the notification
	// channel, it MUST set the pChannel value to NULL upon return from this method. If
	// the notification channel was acquired by a different client, the server MUST set
	// the pChannel value to NULL upon return from this method.
	Channel *NotifyObject `idl:"name:pChannel" json:"channel"`
	// pInNotificationType: A NULL value or a pointer to a PrintAsyncNotificationType structure
	// that specifies the notification type identifier of the notification type in which
	// the registered client is interested.
	InNotificationType *NotificationType `idl:"name:pInNotificationType;pointer:unique" json:"in_notification_type"`
	// InSize: The size, in bytes, of the input data that the pInNotificationData parameter
	// points to. The server SHOULD impose an upper limit of 0x00A00000 on this value. If
	// the client exceeds the server-imposed limit, the server MUST return an error result.
	InSize uint32 `idl:"name:InSize" json:"in_size"`
	// pInNotificationData: A pointer to input data holding the client's response to the
	// previous notification that was received on the same bidirectional notification channel.
	//
	// On the first call to this method for a given channel, the client SHOULD provide zero
	// bytes of response data and the server MUST ignore any response data sent. On subsequent
	// calls to this method, the response format MUST conform to the requirements of the
	// notification channel's notification type, and those notification type requirements
	// determine whether or a not a zero-byte response is acceptable.
	InNotificationData []byte `idl:"name:pInNotificationData;size_is:(InSize);pointer:unique" json:"in_notification_data"`
}

func (o *GetNotificationSendResponseRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNotificationSendResponseOperation) *xxx_GetNotificationSendResponseOperation {
	if op == nil {
		op = &xxx_GetNotificationSendResponseOperation{}
	}
	if o == nil {
		return op
	}
	op.Channel = o.Channel
	op.InNotificationType = o.InNotificationType
	op.InSize = o.InSize
	op.InNotificationData = o.InNotificationData
	return op
}

func (o *GetNotificationSendResponseRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNotificationSendResponseOperation) {
	if o == nil {
		return
	}
	o.Channel = op.Channel
	o.InNotificationType = op.InNotificationType
	o.InSize = op.InSize
	o.InNotificationData = op.InNotificationData
}
func (o *GetNotificationSendResponseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNotificationSendResponseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNotificationSendResponseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNotificationSendResponseResponse structure represents the IRPCAsyncNotify_GetNotificationSendResponse operation response
type GetNotificationSendResponseResponse struct {
	// pChannel: A pointer to a notification channel that MUST NOT be closed or zero, and
	// which MUST have been returned by the server in the ppChannelCtxt output parameter
	// of a prior call to IRPCAsyncNotify_GetNewChannel. If the server closes the notification
	// channel, it MUST set the pChannel value to NULL upon return from this method. If
	// the notification channel was acquired by a different client, the server MUST set
	// the pChannel value to NULL upon return from this method.
	Channel *NotifyObject `idl:"name:pChannel" json:"channel"`
	// ppOutNotificationType: A pointer to the returned pointer to the notification type
	// identifier of the server-to-client notification. If the notification channel was
	// acquired by a different client, the value MUST be NOTIFICATION_RELEASE (section 2.2.1).
	// If the server needs to close the notification channel without sending a final response,
	// the value SHOULD be NOTIFICATION_RELEASE. In all other cases, the value MUST be the
	// same as the notification type identifier of the notification type for which the client
	// has registered.
	OutNotificationType *NotificationType `idl:"name:ppOutNotificationType" json:"out_notification_type"`
	// pOutSize: A pointer to the returned length of server-to-client notification data,
	// in number of bytes. The client MAY impose an upper limit on this value that is smaller
	// than the maximum unsigned 32-bit integer. If the notification channel was acquired
	// by a different client, the server SHOULD set the value of pOutSize to 0x00000000.
	// If the value of ppOutNotificationType points to NOTIFICATION_RELEASE, the server
	// SHOULD set the value of pOutSize to 0x00000000.
	OutSize uint32 `idl:"name:pOutSize" json:"out_size"`
	// ppOutNotificationData: A pointer to the returned pointer to server-to-client notification
	// data in a format that MUST conform to the notification channel's notification type.
	// If the notification channel was acquired by a different client, the server SHOULD
	// set the value of ppOutNotificationData to NULL. If the value of ppOutNotificationType
	// points to NOTIFICATION_RELEASE, the client MUST ignore the content of ppOutNotificationData.
	OutNotificationData []byte `idl:"name:ppOutNotificationData;size_is:(, pOutSize)" json:"out_notification_data"`
	// Return: The IRPCAsyncNotify_GetNotificationSendResponse return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNotificationSendResponseResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNotificationSendResponseOperation) *xxx_GetNotificationSendResponseOperation {
	if op == nil {
		op = &xxx_GetNotificationSendResponseOperation{}
	}
	if o == nil {
		return op
	}
	op.Channel = o.Channel
	op.OutNotificationType = o.OutNotificationType
	op.OutSize = o.OutSize
	op.OutNotificationData = o.OutNotificationData
	op.Return = o.Return
	return op
}

func (o *GetNotificationSendResponseResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNotificationSendResponseOperation) {
	if o == nil {
		return
	}
	o.Channel = op.Channel
	o.OutNotificationType = op.OutNotificationType
	o.OutSize = op.OutSize
	o.OutNotificationData = op.OutNotificationData
	o.Return = op.Return
}
func (o *GetNotificationSendResponseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNotificationSendResponseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNotificationSendResponseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNotificationOperation structure represents the IRPCAsyncNotify_GetNotification operation
type xxx_GetNotificationOperation struct {
	RemoteObject        *irpcremoteobject.RemoteObject `idl:"name:pRemoteObj" json:"remote_object"`
	OutNotificationType *NotificationType              `idl:"name:ppOutNotificationType" json:"out_notification_type"`
	OutSize             uint32                         `idl:"name:pOutSize" json:"out_size"`
	OutNotificationData []byte                         `idl:"name:ppOutNotificationData;size_is:(, pOutSize)" json:"out_notification_data"`
	Return              int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNotificationOperation) OpNum() int { return 5 }

func (o *xxx_GetNotificationOperation) OpName() string {
	return "/IRPCAsyncNotify/v1/IRPCAsyncNotify_GetNotification"
}

func (o *xxx_GetNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pRemoteObj {in} (1:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RemoteObject != nil {
			if err := o.RemoteObject.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&irpcremoteobject.RemoteObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pRemoteObj {in} (1:{context_handle, alias=PRPCREMOTEOBJECT, names=ndr_context_handle}(struct))
	{
		if o.RemoteObject == nil {
			o.RemoteObject = &irpcremoteobject.RemoteObject{}
		}
		if err := o.RemoteObject.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutNotificationData != nil && o.OutSize == 0 {
		o.OutSize = uint32(len(o.OutNotificationData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppOutNotificationType {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PrintAsyncNotificationType, names=GUID}(struct))
	{
		if o.OutNotificationType != nil {
			_ptr_ppOutNotificationType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OutNotificationType != nil {
					if err := o.OutNotificationType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NotificationType{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutNotificationType, _ptr_ppOutNotificationType); err != nil {
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
	// pOutSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.OutSize); err != nil {
			return err
		}
	}
	// ppOutNotificationData {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pOutSize](uint8))
	{
		if o.OutNotificationData != nil || o.OutSize > 0 {
			_ptr_ppOutNotificationData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.OutSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.OutNotificationData {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.OutNotificationData[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.OutNotificationData); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutNotificationData, _ptr_ppOutNotificationData); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppOutNotificationType {out} (1:{pointer=ref}*(2)*(1))(2:{alias=PrintAsyncNotificationType, names=GUID}(struct))
	{
		_ptr_ppOutNotificationType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OutNotificationType == nil {
				o.OutNotificationType = &NotificationType{}
			}
			if err := o.OutNotificationType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppOutNotificationType := func(ptr interface{}) { o.OutNotificationType = *ptr.(**NotificationType) }
		if err := w.ReadPointer(&o.OutNotificationType, _s_ppOutNotificationType, _ptr_ppOutNotificationType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pOutSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.OutSize); err != nil {
			return err
		}
	}
	// ppOutNotificationData {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pOutSize](uint8))
	{
		_ptr_ppOutNotificationData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.OutNotificationData", sizeInfo[0])
			}
			o.OutNotificationData = make([]byte, sizeInfo[0])
			for i1 := range o.OutNotificationData {
				i1 := i1
				if err := w.ReadData(&o.OutNotificationData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppOutNotificationData := func(ptr interface{}) { o.OutNotificationData = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.OutNotificationData, _s_ppOutNotificationData, _ptr_ppOutNotificationData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNotificationRequest structure represents the IRPCAsyncNotify_GetNotification operation request
type GetNotificationRequest struct {
	// pRemoteObj: MUST be the remote object context handle. This remote object MUST have
	// been registered for unidirectional communication mode by a prior successful call
	// to IRPCAsyncNotify_RegisterClient (section 3.1.1.4.1).
	RemoteObject *irpcremoteobject.RemoteObject `idl:"name:pRemoteObj" json:"remote_object"`
}

func (o *GetNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNotificationOperation) *xxx_GetNotificationOperation {
	if op == nil {
		op = &xxx_GetNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.RemoteObject = o.RemoteObject
	return op
}

func (o *GetNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNotificationOperation) {
	if o == nil {
		return
	}
	o.RemoteObject = op.RemoteObject
}
func (o *GetNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNotificationResponse structure represents the IRPCAsyncNotify_GetNotification operation response
type GetNotificationResponse struct {
	// ppOutNotificationType: MUST return a pointer to the notification type identifier
	// of the server-to-client notification. If the registered remote object has been deleted,
	// the value MUST be NOTIFICATION_RELEASE (section 2.2.1). In all other cases the value
	// MUST be the same as the notification type identifier of the notification type for
	// which the print client has registered.
	OutNotificationType *NotificationType `idl:"name:ppOutNotificationType" json:"out_notification_type"`
	// pOutSize: MUST be the length of server-to-client notification data, in number of
	// bytes. The client MAY impose an upper limit on this value that is smaller than the
	// maximum unsigned 32-bit integer.
	OutSize uint32 `idl:"name:pOutSize" json:"out_size"`
	// ppOutNotificationData: MUST be a pointer to server-to-client notification data in
	// a format that MUST conform to the channel's notification type.
	OutNotificationData []byte `idl:"name:ppOutNotificationData;size_is:(, pOutSize)" json:"out_notification_data"`
	// Return: The IRPCAsyncNotify_GetNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNotificationOperation) *xxx_GetNotificationOperation {
	if op == nil {
		op = &xxx_GetNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.OutNotificationType = o.OutNotificationType
	op.OutSize = o.OutSize
	op.OutNotificationData = o.OutNotificationData
	op.Return = o.Return
	return op
}

func (o *GetNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNotificationOperation) {
	if o == nil {
		return
	}
	o.OutNotificationType = op.OutNotificationType
	o.OutSize = op.OutSize
	o.OutNotificationData = op.OutNotificationData
	o.Return = op.Return
}
func (o *GetNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseChannelOperation structure represents the IRPCAsyncNotify_CloseChannel operation
type xxx_CloseChannelOperation struct {
	Channel            *NotifyObject     `idl:"name:pChannel" json:"channel"`
	InNotificationType *NotificationType `idl:"name:pInNotificationType" json:"in_notification_type"`
	InSize             uint32            `idl:"name:InSize" json:"in_size"`
	Reason             []byte            `idl:"name:pReason;size_is:(InSize);pointer:unique" json:"reason"`
	Return             int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseChannelOperation) OpNum() int { return 6 }

func (o *xxx_CloseChannelOperation) OpName() string {
	return "/IRPCAsyncNotify/v1/IRPCAsyncNotify_CloseChannel"
}

func (o *xxx_CloseChannelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Reason != nil && o.InSize == 0 {
		o.InSize = uint32(len(o.Reason))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseChannelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pChannel {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PNOTIFYOBJECT, names=ndr_context_handle}(struct))
	{
		if o.Channel != nil {
			if err := o.Channel.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NotifyObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pInNotificationType {in} (1:{pointer=ref}*(1))(2:{alias=PrintAsyncNotificationType, names=GUID}(struct))
	{
		if o.InNotificationType != nil {
			if err := o.InNotificationType.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NotificationType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.InSize); err != nil {
			return err
		}
	}
	// pReason {in} (1:{pointer=unique}*(1)[dim:0,size_is=InSize](uint8))
	{
		if o.Reason != nil || o.InSize > 0 {
			_ptr_pReason := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Reason {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Reason[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Reason); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Reason, _ptr_pReason); err != nil {
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
	return nil
}

func (o *xxx_CloseChannelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pChannel {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PNOTIFYOBJECT, names=ndr_context_handle}(struct))
	{
		if o.Channel == nil {
			o.Channel = &NotifyObject{}
		}
		if err := o.Channel.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pInNotificationType {in} (1:{pointer=ref}*(1))(2:{alias=PrintAsyncNotificationType, names=GUID}(struct))
	{
		if o.InNotificationType == nil {
			o.InNotificationType = &NotificationType{}
		}
		if err := o.InNotificationType.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.InSize); err != nil {
			return err
		}
	}
	// pReason {in} (1:{pointer=unique}*(1)[dim:0,size_is=InSize](uint8))
	{
		_ptr_pReason := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Reason", sizeInfo[0])
			}
			o.Reason = make([]byte, sizeInfo[0])
			for i1 := range o.Reason {
				i1 := i1
				if err := w.ReadData(&o.Reason[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pReason := func(ptr interface{}) { o.Reason = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Reason, _s_pReason, _ptr_pReason); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseChannelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseChannelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pChannel {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PNOTIFYOBJECT, names=ndr_context_handle}(struct))
	{
		if o.Channel != nil {
			if err := o.Channel.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NotifyObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseChannelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pChannel {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PNOTIFYOBJECT, names=ndr_context_handle}(struct))
	{
		if o.Channel == nil {
			o.Channel = &NotifyObject{}
		}
		if err := o.Channel.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseChannelRequest structure represents the IRPCAsyncNotify_CloseChannel operation request
type CloseChannelRequest struct {
	// pChannel: MUST be a pointer to a notification channel that MUST NOT be closed or
	// zero and that MUST have been returned by the server in the ppChannelCtxt output parameter
	// of a prior call to IRPCAsyncNotify_GetNewChannel. Upon receipt, the server MUST set
	// the pChannel value to NULL.
	Channel *NotifyObject `idl:"name:pChannel" json:"channel"`
	// pInNotificationType: MUST be a pointer to a PrintAsyncNotificationType value. If
	// the client needs to close the notification channels without sending a final response,
	// then this value SHOULD point to NOTIFICATION_RELEASE. In all other cases, this value
	// MUST point to the notification type identifier of the notification type for which
	// the client has registered.
	InNotificationType *NotificationType `idl:"name:pInNotificationType" json:"in_notification_type"`
	// InSize: The server SHOULD impose an upper limit on this value that is smaller than
	// the maximum unsigned 32-bit integer. That limit SHOULD be 0x00A00000. If the client
	// exceeds the server-imposed limit, the server MUST return an error result.
	//
	// If pInNotificationType is NOTIFICATION_RELEASE, then InSize SHOULD be 0x00000000.
	InSize uint32 `idl:"name:InSize" json:"in_size"`
	// pReason: MUST be a pointer to a sequence of bytes conveying final client-to-server
	// response data. The number of bytes MUST be provided in the InSize parameter. If InSize
	// is not 0x00000000, then pReason MUST NOT be NULL.
	//
	// If pInNotificationType is NOTIFICATION_RELEASE, then the client SHOULD provide zero
	// bytes of response data and the server MUST ignore any response data pointed to by
	// pReason. If pInNotificationType is not NOTIFICATION_RELEASE, then the response format
	// MUST conform to the requirements of the notification channel's notification type
	// and those notification type requirements determine whether or not a zero-byte response
	// is acceptable.
	Reason []byte `idl:"name:pReason;size_is:(InSize);pointer:unique" json:"reason"`
}

func (o *CloseChannelRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseChannelOperation) *xxx_CloseChannelOperation {
	if op == nil {
		op = &xxx_CloseChannelOperation{}
	}
	if o == nil {
		return op
	}
	op.Channel = o.Channel
	op.InNotificationType = o.InNotificationType
	op.InSize = o.InSize
	op.Reason = o.Reason
	return op
}

func (o *CloseChannelRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseChannelOperation) {
	if o == nil {
		return
	}
	o.Channel = op.Channel
	o.InNotificationType = op.InNotificationType
	o.InSize = op.InSize
	o.Reason = op.Reason
}
func (o *CloseChannelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseChannelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseChannelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseChannelResponse structure represents the IRPCAsyncNotify_CloseChannel operation response
type CloseChannelResponse struct {
	// pChannel: MUST be a pointer to a notification channel that MUST NOT be closed or
	// zero and that MUST have been returned by the server in the ppChannelCtxt output parameter
	// of a prior call to IRPCAsyncNotify_GetNewChannel. Upon receipt, the server MUST set
	// the pChannel value to NULL.
	Channel *NotifyObject `idl:"name:pChannel" json:"channel"`
	// Return: The IRPCAsyncNotify_CloseChannel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseChannelResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseChannelOperation) *xxx_CloseChannelOperation {
	if op == nil {
		op = &xxx_CloseChannelOperation{}
	}
	if o == nil {
		return op
	}
	op.Channel = o.Channel
	op.Return = o.Return
	return op
}

func (o *CloseChannelResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseChannelOperation) {
	if o == nil {
		return
	}
	o.Channel = op.Channel
	o.Return = op.Return
}
func (o *CloseChannelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseChannelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseChannelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
