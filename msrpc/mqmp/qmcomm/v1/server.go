package qmcomm

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

// qmcomm server interface.
type QmcommServer interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// During the process of creating a remote cursor, a client calls the R_QMGetRemoteQueueName
	// method to retrieve the name of the remote queue associated with a queue handle. This
	// method is obsolete and the server SHOULD take no action and immediately raise the
	// exception MQ_ERROR_ILLEGAL_OPERATION (0xc00e0064).<11>
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs
	// identically. Additionally, if a failure HRESULT is returned, the client MUST disregard
	// all out-parameter values.
	//
	// Exceptions Thrown: This method SHOULD take no action and SHOULD immediately raise
	// the exception MQ_ERROR_ILLEGAL_OPERATION (0xc00e0064).<12>
	//
	// During the remote cursor creation sequence, the supporting server MAY indicate that
	// the client MUST contact a remote queue manager to proceed.<13> In response, this
	// method is called by the client to determine where to find the remote queue manager.
	// Supporting servers SHOULD contact the remote queue manager on behalf of the client,
	// thus eliminating the purpose of this method.<14>
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	GetRemoteQueueName(context.Context, *GetRemoteQueueNameRequest) (*GetRemoteQueueNameResponse, error)

	// A client calls R_QMOpenRemoteQueue to obtain a valid queue handle on a remote queue
	// as part of the sequence of events involved in opening a remote queue as described
	// in section 4.2.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT <26>, and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: In addition to the exceptions thrown by the underlying RPC protocol
	// [MS-RPCE], the method can throw HRESULT failure codes as RPC exceptions. The client
	// MUST treat all thrown HRESULT codes identically. Additionally, the client MUST disregard
	// all out-parameter values when any failure HRESULT is thrown.
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	OpenRemoteQueue(context.Context, *OpenRemoteQueueRequest) (*OpenRemoteQueueResponse, error)

	// The R_QMCloseRemoteQueueContext method closes a remote queue handle originally obtained
	// from R_QMOpenRemoteQueue (section 3.1.4.2).
	//
	// Return Values: None.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	CloseRemoteQueueContext(context.Context, *CloseRemoteQueueContextRequest) (*CloseRemoteQueueContextResponse, error)

	// The R_QMCreateRemoteCursor method creates a cursor at the server for use during remote
	// read.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs
	// identically. Additionally, if a failure HRESULT is returned, the client MUST disregard
	// all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	CreateRemoteCursor(context.Context, *CreateRemoteCursorRequest) (*CreateRemoteCursorResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// A client calls the R_QMCreateObjectInternal method to create a new private queue
	// located on the supporting server.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT.<30><31> If the returned HRESULT value is
	// MQ_ERROR_QUEUE_EXISTS (0xc00e0005), the client can treat it as a success and continue
	// with other operations. The client MUST treat all other failure HRESULTs identically.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	CreateObjectInternal(context.Context, *CreateObjectInternalRequest) (*CreateObjectInternalResponse, error)

	// A client calls the R_QMSetObjectSecurityInternal method to update the security configuration
	// of a private queue located on the supporting server.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<32> and the client MUST treat all failure
	// HRESULTs identically.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	SetObjectSecurityInternal(context.Context, *SetObjectSecurityInternalRequest) (*SetObjectSecurityInternalResponse, error)

	// A client calls the R_QMGetObjectSecurityInternal method to retrieve the security
	// configuration of a private queue located on the supporting server.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<33> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values with the following exception:
	//
	// If nLength is less than the byte length of the buffer required to contain the SECURITY_DESCRIPTOR
	// for the queue identified by pObjectFormat, the server MUST return the byte length
	// of the buffer required to contain the SECURITY_DESCRIPTOR in the lpnLengthNeeded
	// parameter and MUST return MQ_ERROR_SECURITY_DESCRIPTOR_TOO_SMALL (0xc00e0023).
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	GetObjectSecurityInternal(context.Context, *GetObjectSecurityInternalRequest) (*GetObjectSecurityInternalResponse, error)

	// A client calls R_QMDeleteObject to delete a private queue located on the supporting
	// server.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<34><35> and the client MUST treat all failure
	// HRESULTs identically.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)

	// A client calls R_QMGetObjectProperties to retrieve properties from a private queue
	// located on a supporting server.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<37><38> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	GetObjectProperties(context.Context, *GetObjectPropertiesRequest) (*GetObjectPropertiesResponse, error)

	// The R_QMSetObjectProperties method is called by a client to update properties of
	// a local private queue.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<40> and the client MUST treat all failure
	// HRESULTs identically.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	SetObjectProperties(context.Context, *SetObjectPropertiesRequest) (*SetObjectPropertiesResponse, error)

	// A client calls R_QMObjectPathToObjectFormat to determine a format name for a queue
	// identified by a given path name.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<41><42> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	ObjectPathToObjectFormat(context.Context, *ObjectPathToObjectFormatRequest) (*ObjectPathToObjectFormatResponse, error)

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	// A client calls R_QMGetTmWhereabouts to obtain transaction manager whereabouts, as
	// specified in [MS-DTCO], from the supporting server. The whereabouts enable callers
	// to generate exported transaction cookies, which are required to enlist the supporting
	// server's resource manager (RM) in an external transaction.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT. The client MUST treat all failure HRESULTs
	// identically and disregard all out-parameter values, with the following exception:
	//
	// If cbBufSize is less than the size of the SWhereabouts structure returned by the
	// DTC, the server MUST return MQ_ERROR_USER_BUFFER_TOO_SMALL (0xc00e0028).
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	GetTMWhereabouts(context.Context, *GetTMWhereaboutsRequest) (*GetTMWhereaboutsResponse, error)

	// A client calls the R_QMEnlistTransaction method to enlist the supporting server's
	// resource manager (RM) in an external transaction.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<43> and the client MUST treat all failure
	// HRESULTs identically.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	EnlistTransaction(context.Context, *EnlistTransactionRequest) (*EnlistTransactionResponse, error)

	// A client calls the R_QMEnlistInternalTransaction method to enlist the supporting
	// server's resource manager (RM) in an internal transaction. The server returns a transaction
	// handle associated with the given unit of work identifier (XACTUOW). The returned
	// transaction handle is used when calling R_QMCommitTransaction or R_QMAbortTransaction.
	// The XACTUOW structure ([MS-MQMQ] section 2.2.18.1.8) is provided for calls to rpc_ACSendMessageEx
	// and rpc_ACReceiveMessageEx of the qmcomm2 RPC interface.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs
	// identically. Additionally, if a failure HRESULT is returned, the client MUST disregard
	// all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	EnlistInternalTransaction(context.Context, *EnlistInternalTransactionRequest) (*EnlistInternalTransactionResponse, error)

	// A client calls the R_QMCommitTransaction method to commit an internal transaction.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000) and set phIntXact
	// to NULL; otherwise, the server MUST return a failure HRESULT, and the client MUST
	// treat all failure HRESULTs identically.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	CommitTransaction(context.Context, *CommitTransactionRequest) (*CommitTransactionResponse, error)

	// A client calls the R_QMAbortTransaction method to abort an internal transaction.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000) and MUST
	// set phIntXact to NULL; otherwise, the server MUST return a failure HRESULT, and the
	// client MUST treat all failure HRESULTs identically.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	AbortTransaction(context.Context, *AbortTransactionRequest) (*AbortTransactionResponse, error)

	// A client calls rpc_QMOpenQueueInternal to obtain a local queue context handle, to
	// determine if a queue is located at a remote queue manager (section 4.2), or to obtain
	// a local context handle for an opened remote queue. If the call to RemoteQMOpenQueue
	// ([MS-MQQP] section 3.1.4.3) fails, the result MUST be returned to the client, and
	// the remote open queue sequence is discontinued. In the case of failure, any state
	// changes need to be rolled back.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// if an error occurs, the server MUST return a failure HRESULT,<51> and the client
	// MUST treat all failure HRESULTs identically. Additionally, if a failure HRESULT is
	// returned, the client MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: In addition to the exceptions thrown by the underlying RPC protocol,
	// as specified in [MS-RPCE], the method can throw HRESULT failure codes as RPC exceptions.
	// The client MUST treat all thrown HRESULT codes identically. Additionally, the client
	// MUST disregard all out-parameter values when any failure HRESULT is thrown.
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	OpenQueueInternal(context.Context, *OpenQueueInternalRequest) (*OpenQueueInternalResponse, error)

	// A client calls the rpc_ACCloseHandle method to close context handles acquired from
	// rpc_QMOpenQueueInternal (section 3.1.4.17).
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<52> and the client MUST treat all failure
	// HRESULTs identically. Additionally, if a failure HRESULT is returned, the client
	// MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	CloseHandle(context.Context, *CloseHandleRequest) (*CloseHandleResponse, error)

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire

	// A client calls the rpc_ACCloseCursor method to close a cursor acquired from the rpc_ACCreateCursorEx
	// (section 3.1.5.4) method of the qmcomm2 RPC interface.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<54> and the client MUST treat all failure
	// HRESULTs identically.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	CloseCursor(context.Context, *CloseCursorRequest) (*CloseCursorResponse, error)

	// A client calls the rpc_ACSetCursorProperties method to associate a remote cursor
	// created via R_QMCreateRemoteCursor (section 3.1.4.4) with a local CursorProxy (section
	// 3.1.1.6) created using rpc_ACCreateCursorEx (section 3.1.5.4).
	//
	// Note  This method is obsolete. The server SHOULD take no action and return MQ_ERROR_ILLEGAL_OPERATION
	// (0xc00e0064).<55>
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT, and the client MUST treat all failure HRESULTs
	// identically.
	//
	// This method is obsolete. Servers SHOULD take no action and return MQ_ERROR_ILLEGAL_OPERATION
	// (0xc00e0064). Servers SHOULD contact the remote queue manager on behalf of the client
	// when rpc_ACCreateCursorEx is called to create a remote cursor.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// During the client cursor creation sequence, the supporting server MAY indicate that
	// the client MUST contact a remote queue manager to proceed.<56> In response, the client
	// MUST call R_QMGetRemoteQueueName (section 3.1.4.1) to determine the remote queue
	// manager name and MUST then invoke R_QMCreateRemoteCursor at the remote queue manager.
	// Next, the client MUST call this method to associate the Cursor.Handle obtained from
	// R_QMCreateRemoteCursor with the original CursorProxy.Handle obtained from rpc_ACCreateCursorEx.
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	SetCursorProperties(context.Context, *SetCursorPropertiesRequest) (*SetCursorPropertiesResponse, error)

	// Opnum24NotUsedOnWire operation.
	// Opnum24NotUsedOnWire

	// Opnum25NotUsedOnWire operation.
	// Opnum25NotUsedOnWire

	// A client calls the rpc_ACHandleToFormatName method to retrieve a format name for
	// a queue handle.
	//
	// Return Values:  If the provided buffer is long enough to contain the null-terminated
	// format name for the queue identified by hQueue, the server MUST take the following
	// actions:
	//
	// * Copy the null-terminated format name into the lpwcsFormatName buffer.
	//
	// * Set pdwLength to the length (in Unicode characters) of the format name, including
	// the terminating null character.
	//
	// * Return MQ_OK (0x00000000).
	//
	// If the provided buffer is too small to contain the complete format name for the queue
	// identified by hQueue (including the terminating null character), the server MUST
	// take the following actions:
	//
	// * If the buffer length (indicated by pdwLength ) is greater than 0x00000000, and
	// if lpwcsFormatName is non-NULL, copy the format name to the lpwcsFormatName buffer,
	// truncated to fit the length indicated by the input value for pdwLength. The string
	// MUST be null-terminated post-truncation.
	//
	// * Set pdwLength to the length of the untruncated format name, including the terminating
	// null character.
	//
	// * Take no further action and return MQ_ERROR_FORMATNAME_BUFFER_TOO_SMALL (0xc00e001f).
	//
	// If input parameter values violate constraints specified above, the server MUST take
	// no further action and return a failure HRESULT.
	//
	// If any other error occurs, the server MUST return a failure HRESULT,<57> and the
	// client MUST treat all other failure HRESULTs identically. Additionally, if any other
	// failure HRESULT is returned, the client MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	//
	// The format name to be returned to the client (using the rules defined above) is determined
	// as follows:
	//
	// * Locate a LocalQueueContextHandle (section 3.1.1.3) ( 910315e4-d43e-4c99-b086-555cc271563f
	// ) ADM element instance in the server's iLocalQueueContextHandleTable (section 3.1.1.2
	// ( 6d5edb49-62b1-46de-a235-f2cc97df6a31 ) ) where the value of the *Handle* attribute
	// of the *LocalQueueContextHandle* ADM element instance equals hQueue.
	//
	// * If such a *LocalQueueContextHandle* ADM element instance exists:
	//
	// * Declare iLocatedLocalQueueContextHandle and set it to a reference to the located
	// *LocalQueueContextHandle* ADM element instance.
	//
	// * The format name to be returned to the client is iLocatedLocalQueueContextHandle.
	// *OpenQueueDescriptorReference.FormatName*.
	//
	// * Else:
	//
	// * Locate a RemoteQueueProxyHandle (section 3.1.1.5) ( 01412cba-2803-4644-be30-76cdb8560cec
	// ) ADM element instance in the server's iRemoteQueueProxyHandleTable (section 3.1.1.4
	// ( 50eab75b-1cc0-4023-bb19-a20975e50883 ) ) where the value of the Handle attribute
	// of the *RemoteQueueProxyHandle* ADM element instance equals hQueue.
	//
	// * If no such *RemoteQueueProxyHandle* ADM element instance exists, take no further
	// action and return a failure *HRESULT*.
	//
	// * Declare iLocatedRemoteQueueProxyHandle and set it to a reference to the located
	// *RemoteQueueProxyHandle* ADM element instance.
	//
	// * The format name to be returned to the client is iLocatedRemoteQueueProxyHandle.
	// *FormatName*.
	HandleToFormatName(context.Context, *HandleToFormatNameRequest) (*HandleToFormatNameResponse, error)

	// The rpc_ACPurgeQueue method is called by a client to purge an opened queue.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT,<58> and the client MUST treat all failure
	// HRESULTs identically.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// (section 3.1.4.24) method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002)
	// is the interface specified by the fIP parameter.
	PurgeQueue(context.Context, *PurgeQueueRequest) (*PurgeQueueResponse, error)

	// A client calls the R_QMQueryQMRegistryInternal method to retrieve various string
	// values from the supporting server.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// If input parameter values violate constraints specified above, the server MUST take
	// no further action and return a failure HRESULT.
	//
	// If any other error occurs, the server MUST return a failure HRESULT, and the client
	// MUST treat all other failure HRESULTs identically. Additionally, if any other failure
	// HRESULT is returned, the client MUST disregard all out-parameter values.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method is invoked at the dynamically assigned endpoint returned by the R_QMGetRTQMServerPort
	// method when IP_HANDSHAKE (0x00000000) or IPX_HANDSHAKE (0x00000002) is the interface
	// specified by the fIP parameter.
	QueryQMRegistryInternal(context.Context, *QueryQMRegistryInternalRequest) (*QueryQMRegistryInternalResponse, error)

	// Opnum29NotUsedOnWire operation.
	// Opnum29NotUsedOnWire

	// Opnum30NotUsedOnWire operation.
	// Opnum30NotUsedOnWire

	// The R_QMGetRTQMServerPort method returns an RPC port number, as specified in [MS-RPCE],
	// for the requested combination of interface and protocol. The returned RPC port number
	// can be used for all qmcomm and qmcomm2 methods.
	//
	// Return Values: On success, this method returns a non-zero IP port value for the RPC
	// interface specified by the fIP parameter. If an invalid value is specified for fIP,
	// this method MUST return 0x00000000.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// As specified in section 3.1.3, this protocol configures a fixed listening endpoint
	// at an RPC port number. For the interface and protocol specified by the fIP parameter,
	// this method returns the RPC port number determined at server initialization time.
	// If the default port is already in use, the server SHOULD increment the port number
	// by 11 until an unused port is found.
	//
	// Security consideration: Servers MUST NOT enforce security limitations for this method,
	// since clients can call this method before configuring RPC binding security. See section
	// 5.1 for details.
	GetRTQMServerPort(context.Context, *GetRTQMServerPortRequest) (*GetRTQMServerPortResponse, error)

	// Opnum32NotUsedOnWire operation.
	// Opnum32NotUsedOnWire

	// Opnum33NotUsedOnWire operation.
	// Opnum33NotUsedOnWire

	// Opnum34NotUsedOnWire operation.
	// Opnum34NotUsedOnWire
}

func RegisterQmcommServer(conn dcerpc.Conn, o QmcommServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQmcommServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QmcommSyntaxV1_0))...)
}

func NewQmcommServerHandle(o QmcommServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QmcommServerHandle(ctx, o, opNum, r)
	}
}

func QmcommServerHandle(ctx context.Context, o QmcommServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0NotUsedOnWire
		// Opnum0NotUsedOnWire
		return nil, nil
	case 1: // R_QMGetRemoteQueueName
		in := &GetRemoteQueueNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRemoteQueueName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // R_QMOpenRemoteQueue
		in := &OpenRemoteQueueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenRemoteQueue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // R_QMCloseRemoteQueueContext
		in := &CloseRemoteQueueContextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseRemoteQueueContext(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // R_QMCreateRemoteCursor
		in := &CreateRemoteCursorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateRemoteCursor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // R_QMCreateObjectInternal
		in := &CreateObjectInternalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateObjectInternal(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // R_QMSetObjectSecurityInternal
		in := &SetObjectSecurityInternalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetObjectSecurityInternal(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // R_QMGetObjectSecurityInternal
		in := &GetObjectSecurityInternalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetObjectSecurityInternal(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // R_QMDeleteObject
		in := &DeleteObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // R_QMGetObjectProperties
		in := &GetObjectPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetObjectProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // R_QMSetObjectProperties
		in := &SetObjectPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetObjectProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // R_QMObjectPathToObjectFormat
		in := &ObjectPathToObjectFormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ObjectPathToObjectFormat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Opnum13NotUsedOnWire
		// Opnum13NotUsedOnWire
		return nil, nil
	case 14: // R_QMGetTmWhereabouts
		in := &GetTMWhereaboutsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTMWhereabouts(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // R_QMEnlistTransaction
		in := &EnlistTransactionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnlistTransaction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // R_QMEnlistInternalTransaction
		in := &EnlistInternalTransactionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnlistInternalTransaction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // R_QMCommitTransaction
		in := &CommitTransactionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CommitTransaction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // R_QMAbortTransaction
		in := &AbortTransactionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AbortTransaction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // rpc_QMOpenQueueInternal
		in := &OpenQueueInternalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenQueueInternal(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // rpc_ACCloseHandle
		in := &CloseHandleRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseHandle(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // Opnum21NotUsedOnWire
		// Opnum21NotUsedOnWire
		return nil, nil
	case 22: // rpc_ACCloseCursor
		in := &CloseCursorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseCursor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // rpc_ACSetCursorProperties
		in := &SetCursorPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetCursorProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // Opnum24NotUsedOnWire
		// Opnum24NotUsedOnWire
		return nil, nil
	case 25: // Opnum25NotUsedOnWire
		// Opnum25NotUsedOnWire
		return nil, nil
	case 26: // rpc_ACHandleToFormatName
		in := &HandleToFormatNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.HandleToFormatName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // rpc_ACPurgeQueue
		in := &PurgeQueueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PurgeQueue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // R_QMQueryQMRegistryInternal
		in := &QueryQMRegistryInternalRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryQMRegistryInternal(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // Opnum29NotUsedOnWire
		// Opnum29NotUsedOnWire
		return nil, nil
	case 30: // Opnum30NotUsedOnWire
		// Opnum30NotUsedOnWire
		return nil, nil
	case 31: // R_QMGetRTQMServerPort
		in := &GetRTQMServerPortRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRTQMServerPort(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // Opnum32NotUsedOnWire
		// Opnum32NotUsedOnWire
		return nil, nil
	case 33: // Opnum33NotUsedOnWire
		// Opnum33NotUsedOnWire
		return nil, nil
	case 34: // Opnum34NotUsedOnWire
		// Opnum34NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
