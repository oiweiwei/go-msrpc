package qmcomm

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	mqmq "github.com/oiweiwei/go-msrpc/msrpc/mqmq"
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
	_ = mqmq.GoPackage
	_ = dtyp.GoPackage
	_ = dcetypes.GoPackage
)

var (
	// import guard
	GoPackage = "mqmp"
)

var (
	// Syntax UUID
	QmcommSyntaxUUID = &uuid.UUID{TimeLow: 0xfdb3a030, TimeMid: 0x65f, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0x9b, Node: [6]uint8{0x0, 0xa0, 0x24, 0xea, 0x55, 0x25}}
	// Syntax ID
	QmcommSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: QmcommSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// qmcomm interface.
type QmcommClient interface {

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
	GetRemoteQueueName(context.Context, *GetRemoteQueueNameRequest, ...dcerpc.CallOption) (*GetRemoteQueueNameResponse, error)

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
	OpenRemoteQueue(context.Context, *OpenRemoteQueueRequest, ...dcerpc.CallOption) (*OpenRemoteQueueResponse, error)

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
	CloseRemoteQueueContext(context.Context, *CloseRemoteQueueContextRequest, ...dcerpc.CallOption) (*CloseRemoteQueueContextResponse, error)

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
	CreateRemoteCursor(context.Context, *CreateRemoteCursorRequest, ...dcerpc.CallOption) (*CreateRemoteCursorResponse, error)

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
	CreateObjectInternal(context.Context, *CreateObjectInternalRequest, ...dcerpc.CallOption) (*CreateObjectInternalResponse, error)

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
	SetObjectSecurityInternal(context.Context, *SetObjectSecurityInternalRequest, ...dcerpc.CallOption) (*SetObjectSecurityInternalResponse, error)

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
	GetObjectSecurityInternal(context.Context, *GetObjectSecurityInternalRequest, ...dcerpc.CallOption) (*GetObjectSecurityInternalResponse, error)

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
	DeleteObject(context.Context, *DeleteObjectRequest, ...dcerpc.CallOption) (*DeleteObjectResponse, error)

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
	GetObjectProperties(context.Context, *GetObjectPropertiesRequest, ...dcerpc.CallOption) (*GetObjectPropertiesResponse, error)

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
	SetObjectProperties(context.Context, *SetObjectPropertiesRequest, ...dcerpc.CallOption) (*SetObjectPropertiesResponse, error)

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
	ObjectPathToObjectFormat(context.Context, *ObjectPathToObjectFormatRequest, ...dcerpc.CallOption) (*ObjectPathToObjectFormatResponse, error)

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
	GetTMWhereabouts(context.Context, *GetTMWhereaboutsRequest, ...dcerpc.CallOption) (*GetTMWhereaboutsResponse, error)

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
	EnlistTransaction(context.Context, *EnlistTransactionRequest, ...dcerpc.CallOption) (*EnlistTransactionResponse, error)

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
	EnlistInternalTransaction(context.Context, *EnlistInternalTransactionRequest, ...dcerpc.CallOption) (*EnlistInternalTransactionResponse, error)

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
	CommitTransaction(context.Context, *CommitTransactionRequest, ...dcerpc.CallOption) (*CommitTransactionResponse, error)

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
	AbortTransaction(context.Context, *AbortTransactionRequest, ...dcerpc.CallOption) (*AbortTransactionResponse, error)

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
	OpenQueueInternal(context.Context, *OpenQueueInternalRequest, ...dcerpc.CallOption) (*OpenQueueInternalResponse, error)

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
	CloseHandle(context.Context, *CloseHandleRequest, ...dcerpc.CallOption) (*CloseHandleResponse, error)

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
	CloseCursor(context.Context, *CloseCursorRequest, ...dcerpc.CallOption) (*CloseCursorResponse, error)

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
	SetCursorProperties(context.Context, *SetCursorPropertiesRequest, ...dcerpc.CallOption) (*SetCursorPropertiesResponse, error)

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
	HandleToFormatName(context.Context, *HandleToFormatNameRequest, ...dcerpc.CallOption) (*HandleToFormatNameResponse, error)

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
	PurgeQueue(context.Context, *PurgeQueueRequest, ...dcerpc.CallOption) (*PurgeQueueResponse, error)

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
	QueryQMRegistryInternal(context.Context, *QueryQMRegistryInternalRequest, ...dcerpc.CallOption) (*QueryQMRegistryInternalResponse, error)

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
	GetRTQMServerPort(context.Context, *GetRTQMServerPortRequest, ...dcerpc.CallOption) (*GetRTQMServerPortResponse, error)

	// Opnum32NotUsedOnWire operation.
	// Opnum32NotUsedOnWire

	// Opnum33NotUsedOnWire operation.
	// Opnum33NotUsedOnWire

	// Opnum34NotUsedOnWire operation.
	// Opnum34NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// CreateRemoteCursor structure represents CACCreateRemoteCursor RPC structure.
//
// The CACCreateRemoteCursor structure contains the elements necessary for creating
// a cursor on a queue.
type CreateRemoteCursor struct {
	// hCursor:   The value for this field returned from rpc_ACCreateCursorEx (section 3.1.5.4)
	// contains a DWORD value representing an opened cursor.
	Cursor uint32 `idl:"name:hCursor" json:"cursor"`
	// srv_hACQueue:  The value for this field returned from rpc_ACCreateCursorEx is passed
	// to the hQueue parameter of R_QMCreateRemoteCursor (section 3.1.4.4) when invoked
	// as part of a remote cursor creation call sequence.
	ACQueue uint32 `idl:"name:srv_hACQueue" json:"ac_queue"`
	// cli_pQMQueue:  The value for this field returned from rpc_ACCreateCursorEx is passed
	// to the pQueue parameter of R_QMGetRemoteQueueName (section 3.1.4.1) when invoked
	// as part of a remote cursor creation call sequence.
	QMQueue uint32 `idl:"name:cli_pQMQueue" json:"qm_queue"`
}

func (o *CreateRemoteCursor) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CreateRemoteCursor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Cursor); err != nil {
		return err
	}
	if err := w.WriteData(o.ACQueue); err != nil {
		return err
	}
	if err := w.WriteData(o.QMQueue); err != nil {
		return err
	}
	return nil
}
func (o *CreateRemoteCursor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Cursor); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACQueue); err != nil {
		return err
	}
	if err := w.ReadData(&o.QMQueue); err != nil {
		return err
	}
	return nil
}

// TransferType type represents TRANSFER_TYPE RPC enumeration.
//
// The TRANSFER_TYPE enumeration specifies the valid cases for the unnamed union defined
// in the CACTransferBufferV1 structure (section 2.2.3.2).
type TransferType uint16

var (
	// CACTB_SEND:  A send operation (that is, a message placed into a queue for delivery)
	// is to be performed.
	TransferTypeSend TransferType = 0
	// CACTB_RECEIVE:  A receive operation (that is, a message is to be read from a queue)
	// is to be performed.
	TransferTypeReceive TransferType = 1
	// CACTB_CREATECURSOR:  A cursor creation is to be performed.
	TransferTypeCreateCursor TransferType = 2
)

func (o TransferType) String() string {
	switch o {
	case TransferTypeSend:
		return "TransferTypeSend"
	case TransferTypeReceive:
		return "TransferTypeReceive"
	case TransferTypeCreateCursor:
		return "TransferTypeCreateCursor"
	}
	return "Invalid"
}

// TransferBufferV1 structure represents CACTransferBufferV1 RPC structure.
//
// The CACTransferBufferV1 structure is used to send and receive messages via MSMQ.
//
// Following is the layout of the CACTransferBufferV1 structure with IDL annotations
// followed by descriptions of the structure members.
type TransferBufferV1 struct {
	// uTransferType:  The uTransferType member specifies which of the Send, Receive, or
	// CreateCursor union members is present in the CACTransferBufferV1 structure. The uTransferType
	// member MUST be assigned a value from the TRANSFER_TYPE (section 2.2.2.1) enumeration.
	TransferType     uint32                             `idl:"name:uTransferType" json:"transfer_type"`
	TransferBufferV1 *TransferBufferV1_TransferBufferV1 `idl:"name:TransferBufferV1;switch_is:uTransferType" json:"transfer_buffer_v1"`
	// pClass:  This field indicates the message classification, such as a positive acknowledgment,
	// a system-generated report message, or a normal application-generated message. It
	// contains a 16-bit structure as defined below:
	//
	//	+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+
	//	|   |   |   |   |   |   |   |   |   |   |  1  |   |   |   |   |   |
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |  0  | 1 | 2 | 3 | 4 | 5 |
	//	|   |   |   |   |   |   |   |   |   |   |     |   |   |   |   |   |
	//	+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+
	//	| Class Code                        | Reserved      |
	//	+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+
	//
	//
	//	+------------------------+----------------------------------------------------------------------------------+
	//	|                        |                                                                                  |
	//	|         VALUE          |                                     MEANING                                      |
	//	|                        |                                                                                  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| Class Code 0x00 — 0xFF | Specifies the type of the acknowledgment. This field uniquely classifies the     |
	//	|                        | message type within the groupings defined by the fields described above. If the  |
	//	|                        | H bit is set, this field contains an HTTP status code.                           |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| Reserved 0x0000        | MUST be set to all zeros. Clients and servers MUST ignore the Reserved member.   |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| H 0 — 1                | Specifies whether or not HTTP is being used. A value of 0 MUST be used to        |
	//	|                        | specify that HTTP is not being used. A value of 1 MUST be used to specify that   |
	//	|                        | HTTP is being used. If 1, the Class Code field contains an HTTP status code.     |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| R 0 — 1                | Specifies the stage at which the acknowledgment is to occur. A value of 0 MUST   |
	//	|                        | be used to specify that the acknowledgment is for the delivery (arrival) stage.  |
	//	|                        | A value of 1 MUST be used to specify that the acknowledgment is for the receive  |
	//	|                        | stage.                                                                           |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| S 0 — 1                | Specifies the type of acknowledgment. A value of 0 MUST be used to specify       |
	//	|                        | that normal (positive acknowledgment) message processing has occurred. A value   |
	//	|                        | of 1 MUST be used to specify that abnormal (negative acknowledgment) message     |
	//	|                        | processing has occurred.                                                         |
	//	+------------------------+----------------------------------------------------------------------------------+
	//
	// The following table provides correspondence between the message class values defined
	// in [MS-MQMQ] section 2.2.18.1.6 with the abstract message class types defined in
	// [MS-MQDMPR] section 3.1.1.12.
	//
	//	+-----------------------------------------------------+-------------------------------+
	//	|                    MESSAGE CLASS                    |         MESSAGE CLASS         |
	//	|                        VALUE                        |             TYPE              |
	//	+-----------------------------------------------------+-------------------------------+
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NORMAL 0x0000                           | Normal                        |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_REPORT 0x0001                           | Report                        |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_ACK_REACH_QUEUE 0x0002                  | AckReachQueue                 |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_ACK_RECEIVE 0x4000                      | AckReceive                    |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_BAD_DST_Q 0x8000                   | NackBadDestQueue              |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_DELETED 0x8001                     | NackPurged                    |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_REACH_QUEUE_TIMEOUT 0x8002         | NackReachQueueTimeout         |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_Q_EXCEED_QUOTA 0x8003              | NackQueueExceedQuota          |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_ACCESS_DENIED 0x8004               | NackAccessDenied              |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_HOP_COUNT_EXCEEDED 0x8005          | NackHopCountExceeded          |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_BAD_SIGNATURE 0x8006               | NackBadSignature              |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_BAD_ENCRYPTION 0x8007              | NackBadEncryption             |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_NOT_TRANSACTIONAL_Q 0x8009         | NackNotTransactionalQueue     |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_NOT_TRANSACTIONAL_MSG 0x800a       | NackNotTransactionalMessage   |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_UNSUPPORTED_CRYPTO_PROVIDER 0x800b | NackUnsupportedCryptoProvider |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_Q_DELETED 0xc000                   | NackQueueDeleted              |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_Q_PURGED 0xc001                    | NackQueuePurged               |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_RECEIVE_TIMEOUT 0xc002             | NackReceiveTimeout            |
	//	+-----------------------------------------------------+-------------------------------+
	//	| MQMSG_CLASS_NACK_RECEIVE_REJECTED 0xc004            | NackReceiveRejected           |
	//	+-----------------------------------------------------+-------------------------------+
	Class uint16 `idl:"name:pClass" json:"class"`
	// ppMessageID:  The ppMessageID member, if present, specifies a value that can be used
	// to correlate response messages to sent messages.
	MessageID *mqmq.ObjectID `idl:"name:ppMessageID" json:"message_id"`
	// ppCorrelationID:  If present, the ppCorrelationID member is an array of bytes containing
	// an OBJECTID structure (as specified in [MS-MQMQ] section 2.2.8). The ppCorrelationID
	// member, if present, contains a value copied from the ppMessageID member of a previous
	// request and can be used to correlate responses with previously sent messages. The
	// size (in count of bytes) of ppCorrelationID MUST NOT exceed 20.
	CorrelationID []byte `idl:"name:ppCorrelationID;size_is:(, 20);length_is:(, 20)" json:"correlation_id"`
	// pSentTime:  The pSentTime member is formatted in UTC. The pSentTime member specifies
	// the time that the message was sent.
	SentTime uint32 `idl:"name:pSentTime" json:"sent_time"`
	// pArrivedTime:  The pArrivedTime member is formatted in UTC. The pArrivedTime member
	// specifies the time the message was received.
	ArrivedTime uint32 `idl:"name:pArrivedTime" json:"arrived_time"`
	// pPriority:  The pPriority member is a single byte. The pPriority member specifies
	// the processing priority for the message with larger values indicating a higher priority.
	// The byte value MUST be in the range of 0x00 to 0x07. If no priority is set, the default
	// priority value of 0x03 is used. The pPriority member is ignored for transactional
	// messages. Messages that are not part of a transaction will be processed in arrival
	// sequence within priority. The pPriority member is ignored if the message is a part
	// of a transaction.
	Priority uint8 `idl:"name:pPriority" json:"priority"`
	// pDelivery:  The pDelivery member is a single byte. The pDelivery member MUST specify
	// a value of 0x00 or 0x01.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	| 0x00  | A value of 0x00 specifies that the message is not recoverable. The message       |
	//	|       | can remain in volatile storage and is subject to loss in the event of a system   |
	//	|       | crash. This value corresponds to Message.DeliveryGuarantee.Express as defined in |
	//	|       | [MS-MQDMPR] section 3.1.1.12.                                                    |
	//	+-------+----------------------------------------------------------------------------------+
	//	| 0x01  | A value of 0x01 specifies that the message is recoverable and is to be written   |
	//	|       | to non-volatile storage as it moves through the network to its destination and   |
	//	|       | can survive a system crash. Recoverable messages do not have to be part of a     |
	//	|       | transaction. This value corresponds to Message.DeliveryGuarantee.Recoverable as  |
	//	|       | defined in [MS-MQDMPR] section 3.1.1.12.                                         |
	//	+-------+----------------------------------------------------------------------------------+
	Delivery uint8 `idl:"name:pDelivery" json:"delivery"`
	// pAcknowledge:  The pAcknowledge member is a single byte. The pAcknowledge member
	// value specifies the types of acknowledgment messages that are to be generated for
	// this message. Acknowledgment messages are returned in the administration queue. The
	// pAcknowledge member value MUST be assigned from the following list:
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                            |                                                                                  |
	//	|                   VALUE                    |                                     MEANING                                      |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_ACKNOWLEDGMENT_NONE 0x00             | No acknowledgment needed. This value corresponds to                              |
	//	|                                            | Message.AcknowledgementsRequested.None as defined in [MS-MQDMPR] section         |
	//	|                                            | 3.1.1.12.                                                                        |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_ACKNOWLEDGMENT_POS_ARRIVAL 0x01      | Positive acknowledgment is to be sent when the message is                        |
	//	|                                            | placed in the destination queue. This value corresponds to                       |
	//	|                                            | Message.AcknowledgementsRequested.AckPosArrival as defined in [MS-MQDMPR]        |
	//	|                                            | section 3.1.1.12.                                                                |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_ACKNOWLEDGMENT_POS_RECEIVE 0x02      | Positive acknowledgment is to be sent when the message is                        |
	//	|                                            | received from the destination queue. This value corresponds to                   |
	//	|                                            | Message.AcknowledgementsRequested.AckPosReceive as defined in [MS-MQDMPR]        |
	//	|                                            | section 3.1.1.12.                                                                |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_ACKNOWLEDGMENT_NEG_ARRIVAL 0x04      | Negative acknowledgment is to be sent when the message fails                     |
	//	|                                            | to arrive at the destination queue. This value corresponds to                    |
	//	|                                            | Message.AcknowledgementsRequested.AckNegArrival as defined in [MS-MQDMPR]        |
	//	|                                            | section 3.1.1.12.                                                                |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_ACKNOWLEDGMENT_NACK_REACH_QUEUE 0x04 | Negative acknowledgment is to be sent when the message fails                     |
	//	|                                            | to arrive at the destination queue. This value corresponds to                    |
	//	|                                            | Message.AcknowledgementsRequested.AckNegArrival as defined in [MS-MQDMPR]        |
	//	|                                            | section 3.1.1.12.                                                                |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_ACKNOWLEDGMENT_FULL_REACH_QUEUE 0x05 | Positive acknowledgment is to be sent when the message is placed in the          |
	//	|                                            | destination queue and/or negative acknowledgment is to be sent when the          |
	//	|                                            | message fails to arrive at the destination queue. This value corresponds to a    |
	//	|                                            | combination of Message.AcknowledgementsRequested.AckPosArrival and AckNegArrival |
	//	|                                            | as defined in [MS-MQDMPR] section 3.1.1.12.                                      |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_ACKNOWLEDGMENT_NEG_RECEIVE 0x08      | Negative acknowledgment is to be sent when the message fails to                  |
	//	|                                            | be received from the destination queue. This value corresponds to                |
	//	|                                            | Message.AcknowledgementsRequested.AckNegReceive as defined in [MS-MQDMPR]        |
	//	|                                            | section 3.1.1.12.                                                                |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_ACKNOWLEDGMENT_NACK_RECEIVE 0x0C     | Negative acknowledgment is to be sent when the message fails to arrive           |
	//	|                                            | at the destination queue or when a receive for the message from the              |
	//	|                                            | destination queue fails. This value corresponds to a combination of              |
	//	|                                            | Message.AcknowledgementsRequested.AckNegReceive and AckNegArrival as defined in  |
	//	|                                            | [MS-MQDMPR] section 3.1.1.12.                                                    |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_ACKNOWLEDGMENT_FULL_RECEIVE 0x0E     | Positive acknowledgment is to be sent when the message is received from the      |
	//	|                                            | destination queue and a negative acknowledgment is to be sent when the message   |
	//	|                                            | fails to arrive at the destination queue or a negative acknowledgment is to be   |
	//	|                                            | sent when a receive for the message from the destination queue fails. This value |
	//	|                                            | corresponds to a combination of Message.AcknowledgementsRequested.AckNegReceive, |
	//	|                                            | AckNegArrival, and AckPosReceive as defined in [MS-MQDMPR] section 3.1.1.12.     |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	Acknowledge uint8 `idl:"name:pAcknowledge" json:"acknowledge"`
	// pAuditing:  The pAuditing member is a single byte. The pAuditing member value specifies
	// the conditions under which copies of the message are to be stored as the message
	// is routed to the destination queue. The pAuditing member value MUST be assigned from
	// the following list:
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_JOURNAL_NONE 0x00             | Do not store copies. This value corresponds to a                                 |
	//	|                                     | Message.PositiveJournalingRequested value of False and a                         |
	//	|                                     | Message.NegativeJournalingRequested value of False, as defined in [MS-MQDMPR]    |
	//	|                                     | section 3.1.1.12.                                                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_DEADLETTER 0x01               | Store copy in dead-letter queue on failure. This value corresponds               |
	//	|                                     | to a Message.PositiveJournalingRequested value of False and a                    |
	//	|                                     | Message.NegativeJournalingRequested value of True, as defined in [MS-MQDMPR]     |
	//	|                                     | section 3.1.1.12.                                                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_JOURNAL 0x02                  | Store copy in queue journal upon successful delivery to next computer. This      |
	//	|                                     | value corresponds to a Message.PositiveJournalingRequested value of True and a   |
	//	|                                     | Message.NegativeJournalingRequested value of False, as defined in [MS-MQDMPR]    |
	//	|                                     | section 3.1.1.12.                                                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_DEADLETTER|MQMSG_JOURNAL 0x03 | Store copy in queue journal upon successful delivery to next computer.           |
	//	|                                     | Store copy in dead-letter queue on failure. This value corresponds               |
	//	|                                     | to a Message.PositiveJournalingRequested value of True and a                     |
	//	|                                     | Message.NegativeJournalingRequested value of True, as defined in [MS-MQDMPR]     |
	//	|                                     | section 3.1.1.12.                                                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	Auditing uint8 `idl:"name:pAuditing" json:"auditing"`
	// pApplicationTag:  The pApplicationTag member value is a user-provided item that is
	// passed through unmodified to the message-receiving application. A common use of the
	// pApplicationTag member value is to indicate to the receiving application the type
	// of data contained in the ppMsgExtension member.
	ApplicationTag uint32 `idl:"name:pApplicationTag" json:"application_tag"`
	// ppBody:  The ppBody member is an array of bytes. When the ppBody member is present
	// it contains the user message payload.
	Body []byte `idl:"name:ppBody;size_is:(, ulAllocBodyBufferInBytes);length_is:(, ulBodyBufferSizeInBytes)" json:"body"`
	// ulBodyBufferSizeInBytes:  The ulBodyBufferSizeInBytes member specifies the size (in
	// count of bytes) of the data present in the ppBody member. The value of the ulBodyBufferSizeInBytes
	// member MUST be less than or equal to the value in the ulAllocBodyBufferInBytes member.
	BodyBufferSizeInBytes uint32 `idl:"name:ulBodyBufferSizeInBytes" json:"body_buffer_size_in_bytes"`
	// ulAllocBodyBufferInBytes:  The ulAllocBodyBufferInBytes member specifies the size
	// (in count of bytes) of the buffer that is allocated to contain the ppBody member.
	AllocBodyBufferInBytes uint32 `idl:"name:ulAllocBodyBufferInBytes" json:"alloc_body_buffer_in_bytes"`
	// pBodySize:  The pBodySize member specifies the size (in count of bytes) of the data
	// present in the ppBody member after an encryption or decryption operation has been
	// performed on the ppBody member. The value of the pBodySize member MUST be less than
	// or equal to the value in the ulAllocBodyBufferInBytes member.
	BodySize uint32 `idl:"name:pBodySize" json:"body_size"`
	// ppTitle:  The ppTitle member, when present, is a Unicode string. The ppTitle member
	// specifies a title associated with the message.
	Title string `idl:"name:ppTitle;size_is:(, ulTitleBufferSizeInWCHARs);length_is:(, ulTitleBufferSizeInWCHARs)" json:"title"`
	// ulTitleBufferSizeInWCHARs:  The ulTitleBufferSizeInWCHARs member specifies the size
	// (in count of Unicode characters) of the ppTitle member. The ulTitleBufferSizeInWCHARs
	// member MUST NOT exceed 250.
	TitleBufferSizeInWchars uint32 `idl:"name:ulTitleBufferSizeInWCHARs" json:"title_buffer_size_in_wchars"`
	// pulTitleBufferSizeInWCHARs:  The pulTitleBufferSizeInWCHARs member specifies the
	// actual size (in count of Unicode characters) of the string, if present, in the ppTitle
	// member Unicode string.
	ActualTitleBufferSizeInWchars uint32 `idl:"name:pulTitleBufferSizeInWCHARs" json:"actual_title_buffer_size_in_wchars"`
	// ulAbsoluteTimeToQueue:  The ulAbsoluteTimeToQueue member value provided by the client
	// specifies the number of seconds within which the message MUST reach the destination
	// queue or be discarded. Internally, ulAbsoluteTimeToQueue is converted to a UTC time
	// using the clock of the system on which the queue manager is executing.
	AbsoluteTimeToQueue uint32 `idl:"name:ulAbsoluteTimeToQueue" json:"absolute_time_to_queue"`
	// pulRelativeTimeToQueue:  The pulRelativeTimeToQueue member specifies the number of
	// seconds within which the response message MUST reach the destination queue or be
	// discarded.
	ActualRelativeTimeToQueue uint32 `idl:"name:pulRelativeTimeToQueue" json:"actual_relative_time_to_queue"`
	// ulRelativeTimeToLive:  The ulRelativeTimeToLive member value specifies the number
	// of seconds within which the message MUST be received from the destination queue or
	// be discarded. Internally, ulRelativeTimeToLive is converted to a UTC time using the
	// clock of the system on which the queue manager is executing.
	RelativeTimeToLive uint32 `idl:"name:ulRelativeTimeToLive" json:"relative_time_to_live"`
	// pulRelativeTimeToLive:  The pulRelativeTimeToLive member specifies the number of
	// seconds remaining before the response message will be discarded if it is not received
	// from the destination queue.
	ActualRelativeTimeToLive uint32 `idl:"name:pulRelativeTimeToLive" json:"actual_relative_time_to_live"`
	// pTrace:  The pTrace member MUST be a single byte and indicates whether or not tracing
	// is active.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	| 0x00  | A value of 0x00 MUST be used to specify that tracing is not active. This         |
	//	|       | value corresponds to Message.TracingRequested value of False, as defined in      |
	//	|       | [MS-MQDMPR] section 3.1.1.12.                                                    |
	//	+-------+----------------------------------------------------------------------------------+
	//	| 0x01  | A value of 0x01 MUST be used to specify that tracing is active. This value       |
	//	|       | corresponds to Message.TracingRequested value of True, as defined in [MS-MQDMPR] |
	//	|       | section 3.1.1.12.                                                                |
	//	+-------+----------------------------------------------------------------------------------+
	Trace uint8 `idl:"name:pTrace" json:"trace"`
	// pulSenderIDType:  The pulSenderIDType member specifies the type of the ppSenderID
	// member contents. The pulSenderIDType member value MUST be assigned from the following
	// list:
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_SENDERID_TYPE_NONE 0x00000000 | No sender ID is present. This value corresponds to Message.SenderIdentifierType  |
	//	|                                     | value of None, as defined in [MS-MQDMPR] section 3.1.1.12.                       |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_SENDERID_TYPE_SID 0x00000001  | The sender ID is a SID. This value corresponds to Message.SenderIdentifierType   |
	//	|                                     | value of Sid, as defined in [MS-MQDMPR] section 3.1.1.12.                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_SENDERID_TYPE_QM 0x00000002   | The sender ID is the GUID assigned to a queue manager. This value corresponds    |
	//	|                                     | to Message.SenderIdentifierType value of QueueManagerIdentifier, as defined in   |
	//	|                                     | [MS-MQDMPR] section 3.1.1.12.                                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	SenderIDType uint32 `idl:"name:pulSenderIDType" json:"sender_id_type"`
	// ppSenderID:  The ppSenderID member MUST be an array of bytes. When the value of the
	// pulSenderIDType member is 0x00000000 (MQMSG_SENDERID_TYPE_NONE), the ppSenderID member
	// MUST NOT be present. If the value of the pulSenderIDType member is 0x00000001 (MQMSG_SENDERID_TYPE_SID),
	// the ppSenderID member MUST contain a SID. If the value of the pulSenderIDType member
	// is 0x00000002 (MQMSG_SENDERID_TYPE_QM), the ppSenderID member MUST contain a valid
	// MSMQ Site GUID.
	SenderID []byte `idl:"name:ppSenderID;size_is:(, uSenderIDLen)" json:"sender_id"`
	// pulSenderIDLenProp:  The pulSenderIDLenProp member specifies the size (in count of
	// bytes) of the data present in the ppSenderID member.
	SenderIDLengthProperty uint32 `idl:"name:pulSenderIDLenProp" json:"sender_id_length_property"`
	// pulPrivLevel:  The pulPrivLevel member specifies the privacy level that is used for
	// processing the message. The pulPrivLevel member value MUST be assigned from the following
	// list:
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                           |                                                                                  |
	//	|                   VALUE                   |                                     MEANING                                      |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_PRIV_LEVEL_NONE 0x00000000          | The message is not private. This value corresponds to Message.PrivacyLevel value |
	//	|                                           | of None, as defined in [MS-MQDMPR] section 3.1.1.12.                             |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_PRIV_LEVEL_BODY_BASE 0x00000001     | The message is private and the Cryptographic Service Provider (CSP) will use     |
	//	|                                           | a 40-bit encryption key to encrypt and decrypt the message body. This value      |
	//	|                                           | corresponds to Message.PrivacyLevel value of Base, as defined in [MS-MQDMPR]     |
	//	|                                           | section 3.1.1.12.                                                                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_PRIV_LEVEL_BODY_ENHANCED 0x00000002 | The message is private and the CSP will use a 128-bit encryption key to encrypt  |
	//	|                                           | and decrypt the message body. This value corresponds to Message.PrivacyLevel     |
	//	|                                           | value of Enhanced, as defined in [MS-MQDMPR] section 3.1.1.12.                   |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	PrivLevel uint32 `idl:"name:pulPrivLevel" json:"priv_level"`
	// ulAuthLevel:  The ulAuthLevel member is used only in local interprocess communication
	// and therefore has no meaning when this protocol is used over a network. Servers MUST
	// ignore this field, and clients can specify any value.
	AuthLevel uint32 `idl:"name:ulAuthLevel" json:"auth_level"`
	// pAuthenticated:  The pAuthenticated member is a single byte. The pAuthenticated member
	// value is used to determine the level of authentication that has been performed on
	// the message. The pAuthenticated member value MUST be assigned from the following
	// list:
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_AUTHENTICATION_NOT_REQUESTED 0x00 | Authentication has not been performed. This value corresponds to                 |
	//	|                                         | Message.AuthenticationLevel value of None, as defined in [MS-MQDMPR] section     |
	//	|                                         | 3.1.1.12.                                                                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_AUTHENTICATED_SIG10 0x01          | Authentication has been performed using an MSMQ 1.0 digital signature. This      |
	//	|                                         | value corresponds to Message.AuthenticationLevel value of Sig10, as defined in   |
	//	|                                         | [MS-MQDMPR] section 3.1.1.12.                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_AUTHENTICATED_SIG20 0x03          | Authentication has been performed using an MSMQ 2.0 digital signature. This      |
	//	|                                         | value corresponds to Message.AuthenticationLevel value of Sig20, as defined in   |
	//	|                                         | [MS-MQDMPR] section 3.1.1.12.                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_AUTHENTICATED_SIG30 0x05          | Authentication has been performed using an MSMQ 3.0 digital signature. This      |
	//	|                                         | value corresponds to Message.AuthenticationLevel value of Sig30, as defined in   |
	//	|                                         | [MS-MQDMPR] section 3.1.1.12.                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_AUTHENTICATED_SIGXML 0x09         | Authentication has been performed using an XML digital signature. This value     |
	//	|                                         | corresponds to Message.AuthenticationLevel value of XMLSig, as defined in        |
	//	|                                         | [MS-MQDMPR] section 3.1.1.12.                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	Authenticated uint8 `idl:"name:pAuthenticated" json:"authenticated"`
	// pulHashAlg:  The pulHashAlg member specifies the hashing algorithm that is to be
	// used in the digital signing process and by the authentication process. The pulHashAlg
	// member value MUST be assigned from the following list:
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|                               |                                                                                  |
	//	|             VALUE             |                                     MEANING                                      |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_CALG_MD2 0x00008001     | Use the MD2 algorithm as specified in [RFC1319]. This value corresponds to the   |
	//	|                               | Message.HashAlgorithm value of MD2, as defined in [MS-MQDMPR] section 3.1.1.12.  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_CALG_MD4 0x00008002     | Use the MD4 algorithm as specified in [RFC1320]. This value corresponds to the   |
	//	|                               | Message.HashAlgorithm value of MD4, as defined in [MS-MQDMPR] section 3.1.1.12.  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_CALG_MD5 0x00008003     | Use the MD5 algorithm as specified in [RFC1321]. This value corresponds to the   |
	//	|                               | Message.HashAlgorithm value of MD5, as defined in [MS-MQDMPR] section 3.1.1.12.  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_CALG_SHA1 0x00008004    | Use the SHA-1 algorithm as specified in [RFC3174]. This value corresponds to the |
	//	|                               | Message.HashAlgorithm value of SHA1, as defined in [MS-MQDMPR] section 3.1.1.12. |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_CALG_SHA_256 0x0000800C | Use the SHA-256 algorithm, as specified in [FIPS180-2]. This value corresponds   |
	//	|                               | to the Message.HashAlgorithm value of SHA_256, as defined in [MS-MQDMPR] section |
	//	|                               | 3.1.1.12.                                                                        |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_CALG_SHA_512 0x0000800E | Use the SHA-512 algorithm, as specified in [FIPS180-2]. This value corresponds   |
	//	|                               | to the Message.HashAlgorithm value of SHA_512, as defined in [MS-MQDMPR] section |
	//	|                               | 3.1.1.12.                                                                        |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	HashAlgorithm uint32 `idl:"name:pulHashAlg" json:"hash_algorithm"`
	// pulEncryptAlg:  The pulEncryptAlg member specifies that the encryption algorithm
	// is to be used to encrypt and decrypt the message body. The pulEncryptAlg member value
	// MUST be assigned from the following list:
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|                           |                                                                                  |
	//	|           VALUE           |                                     MEANING                                      |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_CALG_RC2 0x00006602 | Use the RC2 encryption algorithm as specified in [RFC2268]. This value           |
	//	|                           | corresponds to Message.EncryptionAlgorithm value of RC2, as defined in           |
	//	|                           | [MS-MQDMPR] section 3.1.1.12.                                                    |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| MQMSG_CALG_RC4 0x00006801 | Use the RC4 encryption algorithm as specified in [RC4]. This value corresponds   |
	//	|                           | to Message.EncryptionAlgorithm value of RC4, as defined in [MS-MQDMPR] section   |
	//	|                           | 3.1.1.12.                                                                        |
	//	+---------------------------+----------------------------------------------------------------------------------+
	EncryptAlgorithm uint32 `idl:"name:pulEncryptAlg" json:"encrypt_algorithm"`
	// ppSenderCert:  The ppSenderCert member is an array of bytes. If not NULL, the ppSenderCert
	// member MUST contain the message sender's X509 certificate. The byte length of the
	// buffer MUST be indicated by ulSenderCertLen.
	SenderCert []byte `idl:"name:ppSenderCert;size_is:(, ulSenderCertLen)" json:"sender_cert"`
	// ulSenderCertLen:  The ulSenderCertLen member specifies the byte length of the certificate
	// contained in ppSenderCert.
	SenderCertLength uint32 `idl:"name:ulSenderCertLen" json:"sender_cert_length"`
	// pulSenderCertLenProp:  The pulSenderCertLenProp member specifies the length (in count
	// of bytes) of the certificate contained in ppSenderCert.
	SenderCertLengthProperty uint32 `idl:"name:pulSenderCertLenProp" json:"sender_cert_length_property"`
	// ppwcsProvName:  The ppwcsProvName member is a Unicode string. If present, the ppwcsProvName
	// member specifies the name of the Cryptographic Service Provider (CSP) that is used
	// to generate digital signatures for the message.
	ProvName string `idl:"name:ppwcsProvName;size_is:(, ulProvNameLen)" json:"prov_name"`
	// ulProvNameLen:  The ulProvNameLen member specifies the size (in count of Unicode
	// characters) of the buffer that was allocated to contain the ppwcsProvName string.
	ProvNameLength uint32 `idl:"name:ulProvNameLen" json:"prov_name_length"`
	// pulAuthProvNameLenProp:  The pulAuthProvNameLenProp member specifies the size (in
	// count of Unicode characters) of the CSP name contained in ppwcsProvName, plus the
	// size of an enhanced signature appended to the ppSignature buffer. Rules for computing
	// and understanding values for this field are defined in sections 3.1.5.3 and 3.1.5.4.
	AuthProvNameLengthProperty uint32 `idl:"name:pulAuthProvNameLenProp" json:"auth_prov_name_length_property"`
	// pulProvType:  The pulProvType member specifies the type of CSP that is named by ppwcsProvName.
	ProvType uint32 `idl:"name:pulProvType" json:"prov_type"`
	// fDefaultProvider:  The fDefaultProvider member specifies if the CSP named by ppwcsProvName
	// is a default CSP. A value of 0x00000000 MUST be used to specify that the ppwcsProvName
	// is not the default name and all other values MUST be interpreted as specifying that
	// the ppwcsProvName is the default name.
	DefaultProvider int32 `idl:"name:fDefaultProvider" json:"default_provider"`
	// ppSymmKeys:  The ppSymmKeys member is an array of bytes. The ppSymmKeys member, if
	// present, contains an encrypted symmetric key.
	SymmetricKeys []byte `idl:"name:ppSymmKeys;size_is:(, ulSymmKeysSize)" json:"symmetric_keys"`
	// ulSymmKeysSize:  The ulSymmKeysSize member specifies the size (in count of bytes)
	// of the buffer that was allocated to contain the ppSymmKeys member.
	SymmetricKeysSize uint32 `idl:"name:ulSymmKeysSize" json:"symmetric_keys_size"`
	// pulSymmKeysSizeProp:  The pulSymmKeysSizeProp member specifies the size (in count
	// of bytes) of the ppSymmKeys member.
	SymmetricKeysSizeProperty uint32 `idl:"name:pulSymmKeysSizeProp" json:"symmetric_keys_size_property"`
	// bEncrypted:  The bEncrypted member is a single byte. The bEncrypted member specifies
	// if the message body is encrypted or is not encrypted. A bEncrypted member value of
	// 0x00 MUST be interpreted as specifying that the message is not encrypted (FALSE)
	// and all other values MUST be interpreted as specifying that the message is encrypted
	// (TRUE).
	Encrypted uint8 `idl:"name:bEncrypted" json:"encrypted"`
	// bAuthenticated:  The bAuthenticated member is a single byte. The bAuthenticated member
	// specifies if the message has been authenticated or has not been authenticated. A
	// bAuthenticated member value of 0x00 MUST be used to specify that the message has
	// not been authenticated (FALSE) and all other values MUST be interpreted as specifying
	// that the message has been authenticated (TRUE).
	IsAuthenticated uint8 `idl:"name:bAuthenticated" json:"is_authenticated"`
	// uSenderIDLen:  The uSenderIDLen member specifies the maximum size (in count of bytes)
	// that is available to contain data in the ppSenderID member.
	SenderIDLength uint16 `idl:"name:uSenderIDLen" json:"sender_id_length"`
	// ppSignature:  The ppSignature member is an array of bytes. The ppSignature member
	// contains the signature(s) used to authenticate the message.<5>
	Signature []byte `idl:"name:ppSignature;size_is:(, ulSignatureSize)" json:"signature"`
	// ulSignatureSize:  The ulSignatureSize member specifies the size (in count of bytes)
	// allocated to hold the ppSignature member.
	SignatureSize uint32 `idl:"name:ulSignatureSize" json:"signature_size"`
	// pulSignatureSizeProp:  The pulSignatureSizeProp member specifies the size (in count
	// of bytes) of the authentication signature(s) in the ppSignature member.
	SignatureSizeProperty uint32 `idl:"name:pulSignatureSizeProp" json:"signature_size_property"`
	// ppSrcQMID:  The ppSrcQMID member is a GUID. The member contains the GUID assigned
	// to the MSMQ installation that is the source of the message.
	SourceQMID *dtyp.GUID `idl:"name:ppSrcQMID" json:"source_qmid"`
	// pUow:  The pUow member is an XACTUOW structure ([MS-MQMQ] section 2.2.18.1.8). If
	// not NULL, this field identifies a transaction for a Send or Receive operation.
	UOW *mqmq.TransactionUOW `idl:"name:pUow" json:"uow"`
	// ppMsgExtension:  The ppMsgExtension member is an array of bytes. The ppMsgExtension
	// member, when present, contains application-specific data. The ppMsgExtension member
	// is primarily used to pass information to foreign queues.
	MessageExtension []byte `idl:"name:ppMsgExtension;size_is:(, ulMsgExtensionBufferInBytes);length_is:(, ulMsgExtensionBufferInBytes)" json:"message_extension"`
	// ulMsgExtensionBufferInBytes:  The ulMsgExtensionBufferInBytes member specifies the
	// size (in count of bytes) of the buffer allocated for the ppMsgExtension array.
	MessageExtensionBufferInBytes uint32 `idl:"name:ulMsgExtensionBufferInBytes" json:"message_extension_buffer_in_bytes"`
	// pMsgExtensionSize:  The pMsgExtensionSize member specifies the size (in count of
	// bytes) of the data contained in the ppMsgExtension array.
	MessageExtensionSize uint32 `idl:"name:pMsgExtensionSize" json:"message_extension_size"`
	// ppConnectorType:  The ppConnectorType member, if present, is a GUID. The ppConnectorType
	// member specifies the identifier of a foreign queue that is used to communicate with
	// a foreign messaging system.
	ConnectorType *dtyp.GUID `idl:"name:ppConnectorType" json:"connector_type"`
	// pulBodyType:  The pulBodyType member value MUST be one of the valid values allowed
	// for a VARTYPE as specified in [MS-MQMQ] section 2.2.12.
	BodyType uint32 `idl:"name:pulBodyType" json:"body_type"`
	// pulVersion:  The pulVersion member specifies the MSMQ packet version.<6>
	Version uint32 `idl:"name:pulVersion" json:"version"`
}

func (o *TransferBufferV1) xxx_PreparePayload(ctx context.Context) error {
	// cannot evaluate expression 20
	// cannot evaluate expression 20
	if o.Body != nil && o.AllocBodyBufferInBytes == 0 {
		o.AllocBodyBufferInBytes = uint32(len(o.Body))
	}
	if o.Body != nil && o.BodyBufferSizeInBytes == 0 {
		o.BodyBufferSizeInBytes = uint32(len(o.Body))
	}
	if o.Title != "" && o.TitleBufferSizeInWchars == 0 {
		o.TitleBufferSizeInWchars = uint32(len(o.Title))
	}
	if o.Title != "" && o.TitleBufferSizeInWchars == 0 {
		o.TitleBufferSizeInWchars = uint32(len(o.Title))
	}
	if o.SenderID != nil && o.SenderIDLength == 0 {
		o.SenderIDLength = uint16(len(o.SenderID))
	}
	if o.SenderCert != nil && o.SenderCertLength == 0 {
		o.SenderCertLength = uint32(len(o.SenderCert))
	}
	if o.ProvName != "" && o.ProvNameLength == 0 {
		o.ProvNameLength = uint32(len(o.ProvName))
	}
	if o.SymmetricKeys != nil && o.SymmetricKeysSize == 0 {
		o.SymmetricKeysSize = uint32(len(o.SymmetricKeys))
	}
	if o.Signature != nil && o.SignatureSize == 0 {
		o.SignatureSize = uint32(len(o.Signature))
	}
	if o.MessageExtension != nil && o.MessageExtensionBufferInBytes == 0 {
		o.MessageExtensionBufferInBytes = uint32(len(o.MessageExtension))
	}
	if o.MessageExtension != nil && o.MessageExtensionBufferInBytes == 0 {
		o.MessageExtensionBufferInBytes = uint32(len(o.MessageExtension))
	}
	if o.TransferType > uint32(2) {
		return fmt.Errorf("TransferType is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransferBufferV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.TransferType); err != nil {
		return err
	}
	_swTransferBufferV1 := uint32(o.TransferType)
	if o.TransferBufferV1 != nil {
		if err := o.TransferBufferV1.MarshalUnionNDR(ctx, w, _swTransferBufferV1); err != nil {
			return err
		}
	} else {
		if err := (&TransferBufferV1_TransferBufferV1{}).MarshalUnionNDR(ctx, w, _swTransferBufferV1); err != nil {
			return err
		}
	}
	_ptr_pClass := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.Class); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.Class, _ptr_pClass); err != nil {
		return err
	}
	if o.MessageID != nil {
		_ptr_ppMessageID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.MessageID != nil {
				_ptr_ppMessageID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.MessageID != nil {
						if err := o.MessageID.MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&mqmq.ObjectID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.MessageID, _ptr_ppMessageID); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MessageID, _ptr_ppMessageID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.CorrelationID != nil {
		_ptr_ppCorrelationID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CorrelationID != nil || 20 > 0 {
				_ptr_ppCorrelationID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(20)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					dimLength1 := uint64(20)
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
					for i1 := range o.CorrelationID {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(o.CorrelationID[i1]); err != nil {
							return err
						}
					}
					for i1 := len(o.CorrelationID); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint8(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.CorrelationID, _ptr_ppCorrelationID); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CorrelationID, _ptr_ppCorrelationID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	_ptr_pSentTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.SentTime); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.SentTime, _ptr_pSentTime); err != nil {
		return err
	}
	_ptr_pArrivedTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.ArrivedTime); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.ArrivedTime, _ptr_pArrivedTime); err != nil {
		return err
	}
	_ptr_pPriority := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.Priority); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.Priority, _ptr_pPriority); err != nil {
		return err
	}
	_ptr_pDelivery := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.Delivery); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.Delivery, _ptr_pDelivery); err != nil {
		return err
	}
	_ptr_pAcknowledge := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.Acknowledge); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.Acknowledge, _ptr_pAcknowledge); err != nil {
		return err
	}
	_ptr_pAuditing := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.Auditing); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.Auditing, _ptr_pAuditing); err != nil {
		return err
	}
	_ptr_pApplicationTag := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.ApplicationTag); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.ApplicationTag, _ptr_pApplicationTag); err != nil {
		return err
	}
	if o.Body != nil {
		_ptr_ppBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Body != nil || o.AllocBodyBufferInBytes > 0 {
				_ptr_ppBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.AllocBodyBufferInBytes)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					dimLength1 := uint64(o.BodyBufferSizeInBytes)
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
					for i1 := range o.Body {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(o.Body[i1]); err != nil {
							return err
						}
					}
					for i1 := len(o.Body); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint8(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Body, _ptr_ppBody); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Body, _ptr_ppBody); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BodyBufferSizeInBytes); err != nil {
		return err
	}
	if err := w.WriteData(o.AllocBodyBufferInBytes); err != nil {
		return err
	}
	_ptr_pBodySize := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.BodySize); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.BodySize, _ptr_pBodySize); err != nil {
		return err
	}
	if o.Title != "" {
		_ptr_ppTitle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Title != "" || o.TitleBufferSizeInWchars > 0 {
				_ptr_ppTitle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.TitleBufferSizeInWchars)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					dimLength1 := uint64(o.TitleBufferSizeInWchars)
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
					_Title_buf := utf16.Encode([]rune(o.Title))
					if uint64(len(_Title_buf)) > sizeInfo[0] {
						_Title_buf = _Title_buf[:sizeInfo[0]]
					}
					for i1 := range _Title_buf {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(_Title_buf[i1]); err != nil {
							return err
						}
					}
					for i1 := len(_Title_buf); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint16(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Title, _ptr_ppTitle); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Title, _ptr_ppTitle); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TitleBufferSizeInWchars); err != nil {
		return err
	}
	_ptr_pulTitleBufferSizeInWCHARs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.ActualTitleBufferSizeInWchars); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.ActualTitleBufferSizeInWchars, _ptr_pulTitleBufferSizeInWCHARs); err != nil {
		return err
	}
	if err := w.WriteData(o.AbsoluteTimeToQueue); err != nil {
		return err
	}
	_ptr_pulRelativeTimeToQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.ActualRelativeTimeToQueue); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.ActualRelativeTimeToQueue, _ptr_pulRelativeTimeToQueue); err != nil {
		return err
	}
	if err := w.WriteData(o.RelativeTimeToLive); err != nil {
		return err
	}
	_ptr_pulRelativeTimeToLive := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.ActualRelativeTimeToLive); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.ActualRelativeTimeToLive, _ptr_pulRelativeTimeToLive); err != nil {
		return err
	}
	_ptr_pTrace := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.Trace); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.Trace, _ptr_pTrace); err != nil {
		return err
	}
	_ptr_pulSenderIDType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.SenderIDType); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.SenderIDType, _ptr_pulSenderIDType); err != nil {
		return err
	}
	if o.SenderID != nil {
		_ptr_ppSenderID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SenderID != nil || o.SenderIDLength > 0 {
				_ptr_ppSenderID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.SenderIDLength)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					for i1 := range o.SenderID {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(o.SenderID[i1]); err != nil {
							return err
						}
					}
					for i1 := len(o.SenderID); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint8(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.SenderID, _ptr_ppSenderID); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SenderID, _ptr_ppSenderID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	_ptr_pulSenderIDLenProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.SenderIDLengthProperty); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.SenderIDLengthProperty, _ptr_pulSenderIDLenProp); err != nil {
		return err
	}
	_ptr_pulPrivLevel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.PrivLevel); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.PrivLevel, _ptr_pulPrivLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthLevel); err != nil {
		return err
	}
	_ptr_pAuthenticated := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.Authenticated); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.Authenticated, _ptr_pAuthenticated); err != nil {
		return err
	}
	_ptr_pulHashAlg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.HashAlgorithm); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.HashAlgorithm, _ptr_pulHashAlg); err != nil {
		return err
	}
	_ptr_pulEncryptAlg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.EncryptAlgorithm); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.EncryptAlgorithm, _ptr_pulEncryptAlg); err != nil {
		return err
	}
	if o.SenderCert != nil {
		_ptr_ppSenderCert := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SenderCert != nil || o.SenderCertLength > 0 {
				_ptr_ppSenderCert := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.SenderCertLength)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					for i1 := range o.SenderCert {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(o.SenderCert[i1]); err != nil {
							return err
						}
					}
					for i1 := len(o.SenderCert); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint8(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.SenderCert, _ptr_ppSenderCert); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SenderCert, _ptr_ppSenderCert); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SenderCertLength); err != nil {
		return err
	}
	_ptr_pulSenderCertLenProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.SenderCertLengthProperty); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.SenderCertLengthProperty, _ptr_pulSenderCertLenProp); err != nil {
		return err
	}
	if o.ProvName != "" {
		_ptr_ppwcsProvName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ProvName != "" || o.ProvNameLength > 0 {
				_ptr_ppwcsProvName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.ProvNameLength)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					_ProvName_buf := utf16.Encode([]rune(o.ProvName))
					if uint64(len(_ProvName_buf)) > sizeInfo[0] {
						_ProvName_buf = _ProvName_buf[:sizeInfo[0]]
					}
					for i1 := range _ProvName_buf {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(_ProvName_buf[i1]); err != nil {
							return err
						}
					}
					for i1 := len(_ProvName_buf); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint16(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.ProvName, _ptr_ppwcsProvName); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ProvName, _ptr_ppwcsProvName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProvNameLength); err != nil {
		return err
	}
	_ptr_pulAuthProvNameLenProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.AuthProvNameLengthProperty); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.AuthProvNameLengthProperty, _ptr_pulAuthProvNameLenProp); err != nil {
		return err
	}
	_ptr_pulProvType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.ProvType); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.ProvType, _ptr_pulProvType); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultProvider); err != nil {
		return err
	}
	if o.SymmetricKeys != nil {
		_ptr_ppSymmKeys := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SymmetricKeys != nil || o.SymmetricKeysSize > 0 {
				_ptr_ppSymmKeys := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.SymmetricKeysSize)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					for i1 := range o.SymmetricKeys {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(o.SymmetricKeys[i1]); err != nil {
							return err
						}
					}
					for i1 := len(o.SymmetricKeys); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint8(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.SymmetricKeys, _ptr_ppSymmKeys); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SymmetricKeys, _ptr_ppSymmKeys); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SymmetricKeysSize); err != nil {
		return err
	}
	_ptr_pulSymmKeysSizeProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.SymmetricKeysSizeProperty); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.SymmetricKeysSizeProperty, _ptr_pulSymmKeysSizeProp); err != nil {
		return err
	}
	if err := w.WriteData(o.Encrypted); err != nil {
		return err
	}
	if err := w.WriteData(o.IsAuthenticated); err != nil {
		return err
	}
	if err := w.WriteData(o.SenderIDLength); err != nil {
		return err
	}
	if o.Signature != nil {
		_ptr_ppSignature := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Signature != nil || o.SignatureSize > 0 {
				_ptr_ppSignature := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.SignatureSize)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					for i1 := range o.Signature {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(o.Signature[i1]); err != nil {
							return err
						}
					}
					for i1 := len(o.Signature); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint8(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Signature, _ptr_ppSignature); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Signature, _ptr_ppSignature); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SignatureSize); err != nil {
		return err
	}
	_ptr_pulSignatureSizeProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.SignatureSizeProperty); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.SignatureSizeProperty, _ptr_pulSignatureSizeProp); err != nil {
		return err
	}
	if o.SourceQMID != nil {
		_ptr_ppSrcQMID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SourceQMID != nil {
				_ptr_ppSrcQMID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.SourceQMID != nil {
						if err := o.SourceQMID.MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.SourceQMID, _ptr_ppSrcQMID); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SourceQMID, _ptr_ppSrcQMID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.UOW != nil {
		_ptr_pUow := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.UOW != nil {
				if err := o.UOW.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.TransactionUOW{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UOW, _ptr_pUow); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.MessageExtension != nil {
		_ptr_ppMsgExtension := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.MessageExtension != nil || o.MessageExtensionBufferInBytes > 0 {
				_ptr_ppMsgExtension := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.MessageExtensionBufferInBytes)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					dimLength1 := uint64(o.MessageExtensionBufferInBytes)
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
					for i1 := range o.MessageExtension {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(o.MessageExtension[i1]); err != nil {
							return err
						}
					}
					for i1 := len(o.MessageExtension); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint8(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.MessageExtension, _ptr_ppMsgExtension); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MessageExtension, _ptr_ppMsgExtension); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MessageExtensionBufferInBytes); err != nil {
		return err
	}
	_ptr_pMsgExtensionSize := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.MessageExtensionSize); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.MessageExtensionSize, _ptr_pMsgExtensionSize); err != nil {
		return err
	}
	if o.ConnectorType != nil {
		_ptr_ppConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ConnectorType != nil {
				_ptr_ppConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.ConnectorType != nil {
						if err := o.ConnectorType.MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.ConnectorType, _ptr_ppConnectorType); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ConnectorType, _ptr_ppConnectorType); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	_ptr_pulBodyType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.BodyType); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.BodyType, _ptr_pulBodyType); err != nil {
		return err
	}
	_ptr_pulVersion := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.Version, _ptr_pulVersion); err != nil {
		return err
	}
	return nil
}
func (o *TransferBufferV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.TransferType); err != nil {
		return err
	}
	if o.TransferBufferV1 == nil {
		o.TransferBufferV1 = &TransferBufferV1_TransferBufferV1{}
	}
	_swTransferBufferV1 := uint32(o.TransferType)
	if err := o.TransferBufferV1.UnmarshalUnionNDR(ctx, w, _swTransferBufferV1); err != nil {
		return err
	}
	_ptr_pClass := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.Class); err != nil {
			return err
		}
		return nil
	})
	_s_pClass := func(ptr interface{}) { o.Class = *ptr.(*uint16) }
	if err := w.ReadPointer(&o.Class, _s_pClass, _ptr_pClass); err != nil {
		return err
	}
	_ptr_ppMessageID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppMessageID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageID == nil {
				o.MessageID = &mqmq.ObjectID{}
			}
			if err := o.MessageID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppMessageID := func(ptr interface{}) { o.MessageID = *ptr.(**mqmq.ObjectID) }
		if err := w.ReadPointer(&o.MessageID, _s_ppMessageID, _ptr_ppMessageID); err != nil {
			return err
		}
		return nil
	})
	_s_ppMessageID := func(ptr interface{}) { o.MessageID = *ptr.(**mqmq.ObjectID) }
	if err := w.ReadPointer(&o.MessageID, _s_ppMessageID, _ptr_ppMessageID); err != nil {
		return err
	}
	_ptr_ppCorrelationID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppCorrelationID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
				return fmt.Errorf("buffer overflow for size %d of array o.CorrelationID", sizeInfo[0])
			}
			o.CorrelationID = make([]byte, sizeInfo[0])
			for i1 := range o.CorrelationID {
				i1 := i1
				if err := w.ReadData(&o.CorrelationID[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppCorrelationID := func(ptr interface{}) { o.CorrelationID = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.CorrelationID, _s_ppCorrelationID, _ptr_ppCorrelationID); err != nil {
			return err
		}
		return nil
	})
	_s_ppCorrelationID := func(ptr interface{}) { o.CorrelationID = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.CorrelationID, _s_ppCorrelationID, _ptr_ppCorrelationID); err != nil {
		return err
	}
	_ptr_pSentTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.SentTime); err != nil {
			return err
		}
		return nil
	})
	_s_pSentTime := func(ptr interface{}) { o.SentTime = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.SentTime, _s_pSentTime, _ptr_pSentTime); err != nil {
		return err
	}
	_ptr_pArrivedTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.ArrivedTime); err != nil {
			return err
		}
		return nil
	})
	_s_pArrivedTime := func(ptr interface{}) { o.ArrivedTime = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.ArrivedTime, _s_pArrivedTime, _ptr_pArrivedTime); err != nil {
		return err
	}
	_ptr_pPriority := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.Priority); err != nil {
			return err
		}
		return nil
	})
	_s_pPriority := func(ptr interface{}) { o.Priority = *ptr.(*uint8) }
	if err := w.ReadPointer(&o.Priority, _s_pPriority, _ptr_pPriority); err != nil {
		return err
	}
	_ptr_pDelivery := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.Delivery); err != nil {
			return err
		}
		return nil
	})
	_s_pDelivery := func(ptr interface{}) { o.Delivery = *ptr.(*uint8) }
	if err := w.ReadPointer(&o.Delivery, _s_pDelivery, _ptr_pDelivery); err != nil {
		return err
	}
	_ptr_pAcknowledge := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.Acknowledge); err != nil {
			return err
		}
		return nil
	})
	_s_pAcknowledge := func(ptr interface{}) { o.Acknowledge = *ptr.(*uint8) }
	if err := w.ReadPointer(&o.Acknowledge, _s_pAcknowledge, _ptr_pAcknowledge); err != nil {
		return err
	}
	_ptr_pAuditing := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.Auditing); err != nil {
			return err
		}
		return nil
	})
	_s_pAuditing := func(ptr interface{}) { o.Auditing = *ptr.(*uint8) }
	if err := w.ReadPointer(&o.Auditing, _s_pAuditing, _ptr_pAuditing); err != nil {
		return err
	}
	_ptr_pApplicationTag := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.ApplicationTag); err != nil {
			return err
		}
		return nil
	})
	_s_pApplicationTag := func(ptr interface{}) { o.ApplicationTag = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.ApplicationTag, _s_pApplicationTag, _ptr_pApplicationTag); err != nil {
		return err
	}
	_ptr_ppBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
				return fmt.Errorf("buffer overflow for size %d of array o.Body", sizeInfo[0])
			}
			o.Body = make([]byte, sizeInfo[0])
			for i1 := range o.Body {
				i1 := i1
				if err := w.ReadData(&o.Body[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppBody := func(ptr interface{}) { o.Body = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Body, _s_ppBody, _ptr_ppBody); err != nil {
			return err
		}
		return nil
	})
	_s_ppBody := func(ptr interface{}) { o.Body = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Body, _s_ppBody, _ptr_ppBody); err != nil {
		return err
	}
	if err := w.ReadData(&o.BodyBufferSizeInBytes); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllocBodyBufferInBytes); err != nil {
		return err
	}
	_ptr_pBodySize := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.BodySize); err != nil {
			return err
		}
		return nil
	})
	_s_pBodySize := func(ptr interface{}) { o.BodySize = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.BodySize, _s_pBodySize, _ptr_pBodySize); err != nil {
		return err
	}
	_ptr_ppTitle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppTitle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			var _Title_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _Title_buf", sizeInfo[0])
			}
			_Title_buf = make([]uint16, sizeInfo[0])
			for i1 := range _Title_buf {
				i1 := i1
				if err := w.ReadData(&_Title_buf[i1]); err != nil {
					return err
				}
			}
			o.Title = strings.TrimRight(string(utf16.Decode(_Title_buf)), ndr.ZeroString)
			return nil
		})
		_s_ppTitle := func(ptr interface{}) { o.Title = *ptr.(*string) }
		if err := w.ReadPointer(&o.Title, _s_ppTitle, _ptr_ppTitle); err != nil {
			return err
		}
		return nil
	})
	_s_ppTitle := func(ptr interface{}) { o.Title = *ptr.(*string) }
	if err := w.ReadPointer(&o.Title, _s_ppTitle, _ptr_ppTitle); err != nil {
		return err
	}
	if err := w.ReadData(&o.TitleBufferSizeInWchars); err != nil {
		return err
	}
	_ptr_pulTitleBufferSizeInWCHARs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.ActualTitleBufferSizeInWchars); err != nil {
			return err
		}
		return nil
	})
	_s_pulTitleBufferSizeInWCHARs := func(ptr interface{}) { o.ActualTitleBufferSizeInWchars = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.ActualTitleBufferSizeInWchars, _s_pulTitleBufferSizeInWCHARs, _ptr_pulTitleBufferSizeInWCHARs); err != nil {
		return err
	}
	if err := w.ReadData(&o.AbsoluteTimeToQueue); err != nil {
		return err
	}
	_ptr_pulRelativeTimeToQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.ActualRelativeTimeToQueue); err != nil {
			return err
		}
		return nil
	})
	_s_pulRelativeTimeToQueue := func(ptr interface{}) { o.ActualRelativeTimeToQueue = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.ActualRelativeTimeToQueue, _s_pulRelativeTimeToQueue, _ptr_pulRelativeTimeToQueue); err != nil {
		return err
	}
	if err := w.ReadData(&o.RelativeTimeToLive); err != nil {
		return err
	}
	_ptr_pulRelativeTimeToLive := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.ActualRelativeTimeToLive); err != nil {
			return err
		}
		return nil
	})
	_s_pulRelativeTimeToLive := func(ptr interface{}) { o.ActualRelativeTimeToLive = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.ActualRelativeTimeToLive, _s_pulRelativeTimeToLive, _ptr_pulRelativeTimeToLive); err != nil {
		return err
	}
	_ptr_pTrace := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.Trace); err != nil {
			return err
		}
		return nil
	})
	_s_pTrace := func(ptr interface{}) { o.Trace = *ptr.(*uint8) }
	if err := w.ReadPointer(&o.Trace, _s_pTrace, _ptr_pTrace); err != nil {
		return err
	}
	_ptr_pulSenderIDType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.SenderIDType); err != nil {
			return err
		}
		return nil
	})
	_s_pulSenderIDType := func(ptr interface{}) { o.SenderIDType = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.SenderIDType, _s_pulSenderIDType, _ptr_pulSenderIDType); err != nil {
		return err
	}
	_ptr_ppSenderID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppSenderID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			// XXX: for opaque unmarshaling
			if o.SenderIDLength > 0 && sizeInfo[0] == 0 {
				sizeInfo[0] = uint64(o.SenderIDLength)
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SenderID", sizeInfo[0])
			}
			o.SenderID = make([]byte, sizeInfo[0])
			for i1 := range o.SenderID {
				i1 := i1
				if err := w.ReadData(&o.SenderID[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppSenderID := func(ptr interface{}) { o.SenderID = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.SenderID, _s_ppSenderID, _ptr_ppSenderID); err != nil {
			return err
		}
		return nil
	})
	_s_ppSenderID := func(ptr interface{}) { o.SenderID = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SenderID, _s_ppSenderID, _ptr_ppSenderID); err != nil {
		return err
	}
	_ptr_pulSenderIDLenProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.SenderIDLengthProperty); err != nil {
			return err
		}
		return nil
	})
	_s_pulSenderIDLenProp := func(ptr interface{}) { o.SenderIDLengthProperty = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.SenderIDLengthProperty, _s_pulSenderIDLenProp, _ptr_pulSenderIDLenProp); err != nil {
		return err
	}
	_ptr_pulPrivLevel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.PrivLevel); err != nil {
			return err
		}
		return nil
	})
	_s_pulPrivLevel := func(ptr interface{}) { o.PrivLevel = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.PrivLevel, _s_pulPrivLevel, _ptr_pulPrivLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthLevel); err != nil {
		return err
	}
	_ptr_pAuthenticated := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.Authenticated); err != nil {
			return err
		}
		return nil
	})
	_s_pAuthenticated := func(ptr interface{}) { o.Authenticated = *ptr.(*uint8) }
	if err := w.ReadPointer(&o.Authenticated, _s_pAuthenticated, _ptr_pAuthenticated); err != nil {
		return err
	}
	_ptr_pulHashAlg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.HashAlgorithm); err != nil {
			return err
		}
		return nil
	})
	_s_pulHashAlg := func(ptr interface{}) { o.HashAlgorithm = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.HashAlgorithm, _s_pulHashAlg, _ptr_pulHashAlg); err != nil {
		return err
	}
	_ptr_pulEncryptAlg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.EncryptAlgorithm); err != nil {
			return err
		}
		return nil
	})
	_s_pulEncryptAlg := func(ptr interface{}) { o.EncryptAlgorithm = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.EncryptAlgorithm, _s_pulEncryptAlg, _ptr_pulEncryptAlg); err != nil {
		return err
	}
	_ptr_ppSenderCert := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppSenderCert := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			// XXX: for opaque unmarshaling
			if o.SenderCertLength > 0 && sizeInfo[0] == 0 {
				sizeInfo[0] = uint64(o.SenderCertLength)
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SenderCert", sizeInfo[0])
			}
			o.SenderCert = make([]byte, sizeInfo[0])
			for i1 := range o.SenderCert {
				i1 := i1
				if err := w.ReadData(&o.SenderCert[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppSenderCert := func(ptr interface{}) { o.SenderCert = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.SenderCert, _s_ppSenderCert, _ptr_ppSenderCert); err != nil {
			return err
		}
		return nil
	})
	_s_ppSenderCert := func(ptr interface{}) { o.SenderCert = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SenderCert, _s_ppSenderCert, _ptr_ppSenderCert); err != nil {
		return err
	}
	if err := w.ReadData(&o.SenderCertLength); err != nil {
		return err
	}
	_ptr_pulSenderCertLenProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.SenderCertLengthProperty); err != nil {
			return err
		}
		return nil
	})
	_s_pulSenderCertLenProp := func(ptr interface{}) { o.SenderCertLengthProperty = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.SenderCertLengthProperty, _s_pulSenderCertLenProp, _ptr_pulSenderCertLenProp); err != nil {
		return err
	}
	_ptr_ppwcsProvName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppwcsProvName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			// XXX: for opaque unmarshaling
			if o.ProvNameLength > 0 && sizeInfo[0] == 0 {
				sizeInfo[0] = uint64(o.ProvNameLength)
			}
			var _ProvName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _ProvName_buf", sizeInfo[0])
			}
			_ProvName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _ProvName_buf {
				i1 := i1
				if err := w.ReadData(&_ProvName_buf[i1]); err != nil {
					return err
				}
			}
			o.ProvName = strings.TrimRight(string(utf16.Decode(_ProvName_buf)), ndr.ZeroString)
			return nil
		})
		_s_ppwcsProvName := func(ptr interface{}) { o.ProvName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ProvName, _s_ppwcsProvName, _ptr_ppwcsProvName); err != nil {
			return err
		}
		return nil
	})
	_s_ppwcsProvName := func(ptr interface{}) { o.ProvName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProvName, _s_ppwcsProvName, _ptr_ppwcsProvName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProvNameLength); err != nil {
		return err
	}
	_ptr_pulAuthProvNameLenProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.AuthProvNameLengthProperty); err != nil {
			return err
		}
		return nil
	})
	_s_pulAuthProvNameLenProp := func(ptr interface{}) { o.AuthProvNameLengthProperty = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.AuthProvNameLengthProperty, _s_pulAuthProvNameLenProp, _ptr_pulAuthProvNameLenProp); err != nil {
		return err
	}
	_ptr_pulProvType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.ProvType); err != nil {
			return err
		}
		return nil
	})
	_s_pulProvType := func(ptr interface{}) { o.ProvType = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.ProvType, _s_pulProvType, _ptr_pulProvType); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultProvider); err != nil {
		return err
	}
	_ptr_ppSymmKeys := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppSymmKeys := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			// XXX: for opaque unmarshaling
			if o.SymmetricKeysSize > 0 && sizeInfo[0] == 0 {
				sizeInfo[0] = uint64(o.SymmetricKeysSize)
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SymmetricKeys", sizeInfo[0])
			}
			o.SymmetricKeys = make([]byte, sizeInfo[0])
			for i1 := range o.SymmetricKeys {
				i1 := i1
				if err := w.ReadData(&o.SymmetricKeys[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppSymmKeys := func(ptr interface{}) { o.SymmetricKeys = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.SymmetricKeys, _s_ppSymmKeys, _ptr_ppSymmKeys); err != nil {
			return err
		}
		return nil
	})
	_s_ppSymmKeys := func(ptr interface{}) { o.SymmetricKeys = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SymmetricKeys, _s_ppSymmKeys, _ptr_ppSymmKeys); err != nil {
		return err
	}
	if err := w.ReadData(&o.SymmetricKeysSize); err != nil {
		return err
	}
	_ptr_pulSymmKeysSizeProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.SymmetricKeysSizeProperty); err != nil {
			return err
		}
		return nil
	})
	_s_pulSymmKeysSizeProp := func(ptr interface{}) { o.SymmetricKeysSizeProperty = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.SymmetricKeysSizeProperty, _s_pulSymmKeysSizeProp, _ptr_pulSymmKeysSizeProp); err != nil {
		return err
	}
	if err := w.ReadData(&o.Encrypted); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsAuthenticated); err != nil {
		return err
	}
	if err := w.ReadData(&o.SenderIDLength); err != nil {
		return err
	}
	_ptr_ppSignature := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppSignature := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			// XXX: for opaque unmarshaling
			if o.SignatureSize > 0 && sizeInfo[0] == 0 {
				sizeInfo[0] = uint64(o.SignatureSize)
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Signature", sizeInfo[0])
			}
			o.Signature = make([]byte, sizeInfo[0])
			for i1 := range o.Signature {
				i1 := i1
				if err := w.ReadData(&o.Signature[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppSignature := func(ptr interface{}) { o.Signature = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Signature, _s_ppSignature, _ptr_ppSignature); err != nil {
			return err
		}
		return nil
	})
	_s_ppSignature := func(ptr interface{}) { o.Signature = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Signature, _s_ppSignature, _ptr_ppSignature); err != nil {
		return err
	}
	if err := w.ReadData(&o.SignatureSize); err != nil {
		return err
	}
	_ptr_pulSignatureSizeProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.SignatureSizeProperty); err != nil {
			return err
		}
		return nil
	})
	_s_pulSignatureSizeProp := func(ptr interface{}) { o.SignatureSizeProperty = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.SignatureSizeProperty, _s_pulSignatureSizeProp, _ptr_pulSignatureSizeProp); err != nil {
		return err
	}
	_ptr_ppSrcQMID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppSrcQMID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SourceQMID == nil {
				o.SourceQMID = &dtyp.GUID{}
			}
			if err := o.SourceQMID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSrcQMID := func(ptr interface{}) { o.SourceQMID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.SourceQMID, _s_ppSrcQMID, _ptr_ppSrcQMID); err != nil {
			return err
		}
		return nil
	})
	_s_ppSrcQMID := func(ptr interface{}) { o.SourceQMID = *ptr.(**dtyp.GUID) }
	if err := w.ReadPointer(&o.SourceQMID, _s_ppSrcQMID, _ptr_ppSrcQMID); err != nil {
		return err
	}
	_ptr_pUow := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UOW == nil {
			o.UOW = &mqmq.TransactionUOW{}
		}
		if err := o.UOW.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pUow := func(ptr interface{}) { o.UOW = *ptr.(**mqmq.TransactionUOW) }
	if err := w.ReadPointer(&o.UOW, _s_pUow, _ptr_pUow); err != nil {
		return err
	}
	_ptr_ppMsgExtension := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppMsgExtension := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
				return fmt.Errorf("buffer overflow for size %d of array o.MessageExtension", sizeInfo[0])
			}
			o.MessageExtension = make([]byte, sizeInfo[0])
			for i1 := range o.MessageExtension {
				i1 := i1
				if err := w.ReadData(&o.MessageExtension[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppMsgExtension := func(ptr interface{}) { o.MessageExtension = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.MessageExtension, _s_ppMsgExtension, _ptr_ppMsgExtension); err != nil {
			return err
		}
		return nil
	})
	_s_ppMsgExtension := func(ptr interface{}) { o.MessageExtension = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.MessageExtension, _s_ppMsgExtension, _ptr_ppMsgExtension); err != nil {
		return err
	}
	if err := w.ReadData(&o.MessageExtensionBufferInBytes); err != nil {
		return err
	}
	_ptr_pMsgExtensionSize := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.MessageExtensionSize); err != nil {
			return err
		}
		return nil
	})
	_s_pMsgExtensionSize := func(ptr interface{}) { o.MessageExtensionSize = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.MessageExtensionSize, _s_pMsgExtensionSize, _ptr_pMsgExtensionSize); err != nil {
		return err
	}
	_ptr_ppConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConnectorType == nil {
				o.ConnectorType = &dtyp.GUID{}
			}
			if err := o.ConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppConnectorType := func(ptr interface{}) { o.ConnectorType = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ConnectorType, _s_ppConnectorType, _ptr_ppConnectorType); err != nil {
			return err
		}
		return nil
	})
	_s_ppConnectorType := func(ptr interface{}) { o.ConnectorType = *ptr.(**dtyp.GUID) }
	if err := w.ReadPointer(&o.ConnectorType, _s_ppConnectorType, _ptr_ppConnectorType); err != nil {
		return err
	}
	_ptr_pulBodyType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.BodyType); err != nil {
			return err
		}
		return nil
	})
	_s_pulBodyType := func(ptr interface{}) { o.BodyType = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.BodyType, _s_pulBodyType, _ptr_pulBodyType); err != nil {
		return err
	}
	_ptr_pulVersion := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
		return nil
	})
	_s_pulVersion := func(ptr interface{}) { o.Version = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.Version, _s_pulVersion, _ptr_pulVersion); err != nil {
		return err
	}
	return nil
}

// TransferBufferV1_TransferBufferV1 structure represents CACTransferBufferV1 union anonymous member.
//
// The CACTransferBufferV1 structure is used to send and receive messages via MSMQ.
//
// Following is the layout of the CACTransferBufferV1 structure with IDL annotations
// followed by descriptions of the structure members.
type TransferBufferV1_TransferBufferV1 struct {
	// Types that are assignable to Value
	//
	// *TransferBufferV1_Send
	// *TransferBufferV1_Receive
	// *TransferBufferV1_CreateCursor
	Value is_TransferBufferV1_TransferBufferV1 `json:"value"`
}

func (o *TransferBufferV1_TransferBufferV1) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *TransferBufferV1_Send:
		if value != nil {
			return value.Send
		}
	case *TransferBufferV1_Receive:
		if value != nil {
			return value.Receive
		}
	case *TransferBufferV1_CreateCursor:
		if value != nil {
			return value.CreateCursor
		}
	}
	return nil
}

type is_TransferBufferV1_TransferBufferV1 interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_TransferBufferV1_TransferBufferV1()
}

func (o *TransferBufferV1_TransferBufferV1) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *TransferBufferV1_Send:
		return uint32(0)
	case *TransferBufferV1_Receive:
		return uint32(1)
	case *TransferBufferV1_CreateCursor:
		return uint32(2)
	}
	return uint32(0)
}

func (o *TransferBufferV1_TransferBufferV1) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*TransferBufferV1_Send)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TransferBufferV1_Send{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1):
		_o, _ := o.Value.(*TransferBufferV1_Receive)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TransferBufferV1_Receive{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*TransferBufferV1_CreateCursor)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TransferBufferV1_CreateCursor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *TransferBufferV1_TransferBufferV1) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &TransferBufferV1_Send{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1):
		o.Value = &TransferBufferV1_Receive{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &TransferBufferV1_CreateCursor{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// TransferBufferV1_Send structure represents TransferBufferV1_TransferBufferV1 RPC union arm.
//
// It has following labels: 0
type TransferBufferV1_Send struct {
	// Send:  The Send structure is present in the CACTransferBufferV1 structure when the
	// value of the uTransferType member is 0x00000000 (CACTB_SEND). The Send structure
	// is defined inline to the CACTransferBufferV1 structure. The Send structure members
	// are defined as follows:
	Send *TransferBufferV1_TransferBufferV1_Send `idl:"name:Send" json:"send"`
}

func (*TransferBufferV1_Send) is_TransferBufferV1_TransferBufferV1() {}

func (o *TransferBufferV1_Send) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Send != nil {
		if err := o.Send.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TransferBufferV1_TransferBufferV1_Send{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransferBufferV1_Send) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Send == nil {
		o.Send = &TransferBufferV1_TransferBufferV1_Send{}
	}
	if err := o.Send.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TransferBufferV1_TransferBufferV1_Send structure represents CACTransferBufferV1 structure anonymous member.
//
// The CACTransferBufferV1 structure is used to send and receive messages via MSMQ.
//
// Following is the layout of the CACTransferBufferV1 structure with IDL annotations
// followed by descriptions of the structure members.
type TransferBufferV1_TransferBufferV1_Send struct {
	// pAdminQueueFormat:  The pAdminQueueFormat member is a QUEUE_FORMAT ([MS-MQMQ] section
	// 2.2.7) structure. If present, the pAdminQueueFormat member describes the administration
	// queue that is to be used for send operation acknowledgments.
	AdminQueueFormat *mqmq.QueueFormat `idl:"name:pAdminQueueFormat" json:"admin_queue_format"`
	// pResponseQueueFormat:  The pResponseQueueFormat member is a QUEUE_FORMAT structure.
	// If present, the pResponseQueueFormat member describes the queue that is to be used
	// for application-specific responses. As an application-specific value, this field
	// SHOULD be ignored by the server.
	ResponseQueueFormat *mqmq.QueueFormat `idl:"name:pResponseQueueFormat" json:"response_queue_format"`
}

func (o *TransferBufferV1_TransferBufferV1_Send) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransferBufferV1_TransferBufferV1_Send) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.AdminQueueFormat != nil {
		_ptr_pAdminQueueFormat := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AdminQueueFormat != nil {
				if err := o.AdminQueueFormat.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.QueueFormat{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AdminQueueFormat, _ptr_pAdminQueueFormat); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ResponseQueueFormat != nil {
		_ptr_pResponseQueueFormat := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ResponseQueueFormat != nil {
				if err := o.ResponseQueueFormat.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.QueueFormat{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResponseQueueFormat, _ptr_pResponseQueueFormat); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransferBufferV1_TransferBufferV1_Send) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_pAdminQueueFormat := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AdminQueueFormat == nil {
			o.AdminQueueFormat = &mqmq.QueueFormat{}
		}
		if err := o.AdminQueueFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pAdminQueueFormat := func(ptr interface{}) { o.AdminQueueFormat = *ptr.(**mqmq.QueueFormat) }
	if err := w.ReadPointer(&o.AdminQueueFormat, _s_pAdminQueueFormat, _ptr_pAdminQueueFormat); err != nil {
		return err
	}
	_ptr_pResponseQueueFormat := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ResponseQueueFormat == nil {
			o.ResponseQueueFormat = &mqmq.QueueFormat{}
		}
		if err := o.ResponseQueueFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pResponseQueueFormat := func(ptr interface{}) { o.ResponseQueueFormat = *ptr.(**mqmq.QueueFormat) }
	if err := w.ReadPointer(&o.ResponseQueueFormat, _s_pResponseQueueFormat, _ptr_pResponseQueueFormat); err != nil {
		return err
	}
	return nil
}

// TransferBufferV1_Receive structure represents TransferBufferV1_TransferBufferV1 RPC union arm.
//
// It has following labels: 1
type TransferBufferV1_Receive struct {
	// Receive:  The Receive structure is present in the CACTransferBufferV1 structure when
	// the value of the uTransferType member is 0x00000001 (CACTB_RECEIVE). The Receive
	// structure is defined inline to the CACTransferBufferV1 structure. The Receive structure
	// members are defined as follows:
	Receive *TransferBufferV1_TransferBufferV1_Receive `idl:"name:Receive" json:"receive"`
}

func (*TransferBufferV1_Receive) is_TransferBufferV1_TransferBufferV1() {}

func (o *TransferBufferV1_Receive) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Receive != nil {
		if err := o.Receive.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TransferBufferV1_TransferBufferV1_Receive{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransferBufferV1_Receive) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Receive == nil {
		o.Receive = &TransferBufferV1_TransferBufferV1_Receive{}
	}
	if err := o.Receive.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TransferBufferV1_TransferBufferV1_Receive structure represents CACTransferBufferV1 structure anonymous member.
//
// The CACTransferBufferV1 structure is used to send and receive messages via MSMQ.
//
// Following is the layout of the CACTransferBufferV1 structure with IDL annotations
// followed by descriptions of the structure members.
type TransferBufferV1_TransferBufferV1_Receive struct {
	// RequestTimeout:  The RequestTimeout member specifies the amount of time (in milliseconds)
	// to wait for a message to be returned before returning a failure.
	RequestTimeout uint32 `idl:"name:RequestTimeout" json:"request_timeout"`
	// Action:  The Action member specifies the type of receive operation that is to be
	// performed. The Action member MUST specify one of the values: 0x00000000 (MQ_ACTION_RECEIVE),
	// 0x80000000 (MQ_ACTION_PEEK_CURRENT), or 0x80000001 (MQ_ACTION_PEEK_NEXT).
	//
	//	+------------------------+------------+
	//	|                        |            |
	//	|          NAME          |   VALUE    |
	//	|                        |            |
	//	+------------------------+------------+
	//	+------------------------+------------+
	//	| MQ_ACTION_RECEIVE      | 0x00000000 |
	//	+------------------------+------------+
	//	| MQ_ACTION_PEEK_CURRENT | 0x80000000 |
	//	+------------------------+------------+
	//	| MQ_ACTION_PEEK_NEXT    | 0x80000001 |
	//	+------------------------+------------+
	Action uint32 `idl:"name:Action" json:"action"`
	// Asynchronous:  The Asynchronous member is used as a Boolean variable to indicate
	// if the receive is to be performed asynchronously. An Asynchronous member value of
	// 0x00000000 SHOULD be interpreted as specifying FALSE (receive operation is not to
	// be performed asynchronously) and all other values SHOULD be interpreted as TRUE (receive
	// operation is to be performed asynchronously).<4>
	Asynchronous uint32 `idl:"name:Asynchronous" json:"asynchronous"`
	// Cursor:  A cursor handle obtained from rpc_ACCreateCursorEx (section 3.1.5.4). A
	// cursor can be used to reference a specific position within the message queue, rather
	// than the first message in the queue, from which the message will be retrieved.
	Cursor uint32 `idl:"name:Cursor" json:"cursor"`
	// ulResponseFormatNameLen:  The ulResponseFormatNameLen member specifies the size (in
	// count of Unicode characters) of the string allocated for the ppResponseFormatName
	// member. The ulResponseFormatNameLen member MUST have a value in the range of 0 to
	// 1024, inclusive.
	ResponseFormatNameLength uint32 `idl:"name:ulResponseFormatNameLen" json:"response_format_name_length"`
	// ppResponseFormatName:  A null-terminated Unicode string containing a format name
	// (as specified in [MS-MQMQ]) which indicates an application-defined queue which can
	// be used for response messages. This value is used only by MSMQ applications, and
	// it MUST be ignored by MSMQ queue managers.
	ResponseFormatName string `idl:"name:ppResponseFormatName;size_is:(, ulResponseFormatNameLen)" json:"response_format_name"`
	// pulResponseFormatNameLenProp:  The pulResponseFormatNameLenProp member specifies
	// the size (in count of Unicode characters) of the string contained in the ppResponseFormatName
	// member.
	ResponseFormatNameLengthProperty uint32 `idl:"name:pulResponseFormatNameLenProp" json:"response_format_name_length_property"`
	// ulAdminFormatNameLen:  The ulAdminFormatNameLen member specifies the size (in count
	// of Unicode characters) of the string allocated for the ppAdminFormatName member.
	// The ulAdminFormatNameLen member MUST have a value in the range of 0 to 1024, inclusive.
	AdminFormatNameLength uint32 `idl:"name:ulAdminFormatNameLen" json:"admin_format_name_length"`
	// ppAdminFormatName:  A null-terminated Unicode string containing a format name (as
	// specified in [MS-MQMQ]) which indicates an application-defined administration queue
	// to which acknowledgment messages will be directed.
	AdminFormatName string `idl:"name:ppAdminFormatName;size_is:(, ulAdminFormatNameLen)" json:"admin_format_name"`
	// pulAdminFormatNameLenProp:  The pulAdminFormatNameLenProp member specifies the size
	// (in count of Unicode characters) of the string contained in the ppAdminFormatName
	// member.
	AdminFormatNameLengthProperty uint32 `idl:"name:pulAdminFormatNameLenProp" json:"admin_format_name_length_property"`
	// ulDestFormatNameLen:  The ulDestFormatNameLen member specifies the size (in count
	// of Unicode characters) of the string allocated for the ppDestFormatName member. The
	// ulDestFormatNameLen member MUST have a value in the range of 0 to 1024, inclusive.
	DestinationFormatNameLength uint32 `idl:"name:ulDestFormatNameLen" json:"destination_format_name_length"`
	// ppDestFormatName:  A null-terminated Unicode string containing a format name (as
	// specified in [MS-MQMQ]) that indicates the name of a message's destination queue.
	DestinationFormatName string `idl:"name:ppDestFormatName;size_is:(, ulDestFormatNameLen)" json:"destination_format_name"`
	// pulDestFormatNameLenProp:  The pulDestFormatNameLenProp member specifies the size
	// (in count of Unicode characters) of the string contained in the ppDestFormatName
	// member.
	DestinationFormatNameLengthProperty uint32 `idl:"name:pulDestFormatNameLenProp" json:"destination_format_name_length_property"`
	// ulOrderingFormatNameLen:  The ulOrderingFormatNameLen member specifies the size (in
	// count of Unicode characters) of the string allocated for the ppOrderingFormatName
	// member. The ulOrderingFormatNameLen member MUST have a value in the range of 0 to
	// 1024, inclusive.
	OrderingFormatNameLength uint32 `idl:"name:ulOrderingFormatNameLen" json:"ordering_format_name_length"`
	// ppOrderingFormatName:  A null-terminated Unicode string containing a format name
	// (as specified in [MS-MQMQ]) that indicates the name of the MSMQ order queue that
	// tracks the ordering of transactional messages.
	OrderingFormatName string `idl:"name:ppOrderingFormatName;size_is:(, ulOrderingFormatNameLen)" json:"ordering_format_name"`
	// pulOrderingFormatNameLenProp:  The pulOrderingFormatNameLenProp member specifies
	// the size (in count of Unicode characters) of the string contained in the ppOrderingFormatName
	// member.
	OrderingFormatNameLengthProperty uint32 `idl:"name:pulOrderingFormatNameLenProp" json:"ordering_format_name_length_property"`
}

func (o *TransferBufferV1_TransferBufferV1_Receive) xxx_PreparePayload(ctx context.Context) error {
	if o.ResponseFormatName != "" && o.ResponseFormatNameLength == 0 {
		o.ResponseFormatNameLength = uint32(len(o.ResponseFormatName))
	}
	if o.AdminFormatName != "" && o.AdminFormatNameLength == 0 {
		o.AdminFormatNameLength = uint32(len(o.AdminFormatName))
	}
	if o.DestinationFormatName != "" && o.DestinationFormatNameLength == 0 {
		o.DestinationFormatNameLength = uint32(len(o.DestinationFormatName))
	}
	if o.OrderingFormatName != "" && o.OrderingFormatNameLength == 0 {
		o.OrderingFormatNameLength = uint32(len(o.OrderingFormatName))
	}
	if o.ResponseFormatNameLength > uint32(1024) {
		return fmt.Errorf("ResponseFormatNameLength is out of range")
	}
	if o.AdminFormatNameLength > uint32(1024) {
		return fmt.Errorf("AdminFormatNameLength is out of range")
	}
	if o.DestinationFormatNameLength > uint32(1024) {
		return fmt.Errorf("DestinationFormatNameLength is out of range")
	}
	if o.OrderingFormatNameLength > uint32(1024) {
		return fmt.Errorf("OrderingFormatNameLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransferBufferV1_TransferBufferV1_Receive) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.Action); err != nil {
		return err
	}
	if err := w.WriteData(o.Asynchronous); err != nil {
		return err
	}
	if err := w.WriteData(o.Cursor); err != nil {
		return err
	}
	if err := w.WriteData(o.ResponseFormatNameLength); err != nil {
		return err
	}
	if o.ResponseFormatName != "" {
		_ptr_ppResponseFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ResponseFormatName != "" || o.ResponseFormatNameLength > 0 {
				_ptr_ppResponseFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.ResponseFormatNameLength)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					_ResponseFormatName_buf := utf16.Encode([]rune(o.ResponseFormatName))
					if uint64(len(_ResponseFormatName_buf)) > sizeInfo[0] {
						_ResponseFormatName_buf = _ResponseFormatName_buf[:sizeInfo[0]]
					}
					for i1 := range _ResponseFormatName_buf {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(_ResponseFormatName_buf[i1]); err != nil {
							return err
						}
					}
					for i1 := len(_ResponseFormatName_buf); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint16(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.ResponseFormatName, _ptr_ppResponseFormatName); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResponseFormatName, _ptr_ppResponseFormatName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	_ptr_pulResponseFormatNameLenProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.ResponseFormatNameLengthProperty); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.ResponseFormatNameLengthProperty, _ptr_pulResponseFormatNameLenProp); err != nil {
		return err
	}
	if err := w.WriteData(o.AdminFormatNameLength); err != nil {
		return err
	}
	if o.AdminFormatName != "" {
		_ptr_ppAdminFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AdminFormatName != "" || o.AdminFormatNameLength > 0 {
				_ptr_ppAdminFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.AdminFormatNameLength)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					_AdminFormatName_buf := utf16.Encode([]rune(o.AdminFormatName))
					if uint64(len(_AdminFormatName_buf)) > sizeInfo[0] {
						_AdminFormatName_buf = _AdminFormatName_buf[:sizeInfo[0]]
					}
					for i1 := range _AdminFormatName_buf {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(_AdminFormatName_buf[i1]); err != nil {
							return err
						}
					}
					for i1 := len(_AdminFormatName_buf); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint16(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.AdminFormatName, _ptr_ppAdminFormatName); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AdminFormatName, _ptr_ppAdminFormatName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	_ptr_pulAdminFormatNameLenProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.AdminFormatNameLengthProperty); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.AdminFormatNameLengthProperty, _ptr_pulAdminFormatNameLenProp); err != nil {
		return err
	}
	if err := w.WriteData(o.DestinationFormatNameLength); err != nil {
		return err
	}
	if o.DestinationFormatName != "" {
		_ptr_ppDestFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DestinationFormatName != "" || o.DestinationFormatNameLength > 0 {
				_ptr_ppDestFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.DestinationFormatNameLength)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					_DestinationFormatName_buf := utf16.Encode([]rune(o.DestinationFormatName))
					if uint64(len(_DestinationFormatName_buf)) > sizeInfo[0] {
						_DestinationFormatName_buf = _DestinationFormatName_buf[:sizeInfo[0]]
					}
					for i1 := range _DestinationFormatName_buf {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(_DestinationFormatName_buf[i1]); err != nil {
							return err
						}
					}
					for i1 := len(_DestinationFormatName_buf); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint16(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.DestinationFormatName, _ptr_ppDestFormatName); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DestinationFormatName, _ptr_ppDestFormatName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	_ptr_pulDestFormatNameLenProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.DestinationFormatNameLengthProperty); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.DestinationFormatNameLengthProperty, _ptr_pulDestFormatNameLenProp); err != nil {
		return err
	}
	if err := w.WriteData(o.OrderingFormatNameLength); err != nil {
		return err
	}
	if o.OrderingFormatName != "" {
		_ptr_ppOrderingFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.OrderingFormatName != "" || o.OrderingFormatNameLength > 0 {
				_ptr_ppOrderingFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					dimSize1 := uint64(o.OrderingFormatNameLength)
					if err := w.WriteSize(dimSize1); err != nil {
						return err
					}
					sizeInfo := []uint64{
						dimSize1,
					}
					_OrderingFormatName_buf := utf16.Encode([]rune(o.OrderingFormatName))
					if uint64(len(_OrderingFormatName_buf)) > sizeInfo[0] {
						_OrderingFormatName_buf = _OrderingFormatName_buf[:sizeInfo[0]]
					}
					for i1 := range _OrderingFormatName_buf {
						i1 := i1
						if uint64(i1) >= sizeInfo[0] {
							break
						}
						if err := w.WriteData(_OrderingFormatName_buf[i1]); err != nil {
							return err
						}
					}
					for i1 := len(_OrderingFormatName_buf); uint64(i1) < sizeInfo[0]; i1++ {
						if err := w.WriteData(uint16(0)); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.OrderingFormatName, _ptr_ppOrderingFormatName); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OrderingFormatName, _ptr_ppOrderingFormatName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	_ptr_pulOrderingFormatNameLenProp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.OrderingFormatNameLengthProperty); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.OrderingFormatNameLengthProperty, _ptr_pulOrderingFormatNameLenProp); err != nil {
		return err
	}
	return nil
}
func (o *TransferBufferV1_TransferBufferV1_Receive) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.Action); err != nil {
		return err
	}
	if err := w.ReadData(&o.Asynchronous); err != nil {
		return err
	}
	if err := w.ReadData(&o.Cursor); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResponseFormatNameLength); err != nil {
		return err
	}
	_ptr_ppResponseFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppResponseFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			// XXX: for opaque unmarshaling
			if o.ResponseFormatNameLength > 0 && sizeInfo[0] == 0 {
				sizeInfo[0] = uint64(o.ResponseFormatNameLength)
			}
			var _ResponseFormatName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _ResponseFormatName_buf", sizeInfo[0])
			}
			_ResponseFormatName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _ResponseFormatName_buf {
				i1 := i1
				if err := w.ReadData(&_ResponseFormatName_buf[i1]); err != nil {
					return err
				}
			}
			o.ResponseFormatName = strings.TrimRight(string(utf16.Decode(_ResponseFormatName_buf)), ndr.ZeroString)
			return nil
		})
		_s_ppResponseFormatName := func(ptr interface{}) { o.ResponseFormatName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ResponseFormatName, _s_ppResponseFormatName, _ptr_ppResponseFormatName); err != nil {
			return err
		}
		return nil
	})
	_s_ppResponseFormatName := func(ptr interface{}) { o.ResponseFormatName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ResponseFormatName, _s_ppResponseFormatName, _ptr_ppResponseFormatName); err != nil {
		return err
	}
	_ptr_pulResponseFormatNameLenProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.ResponseFormatNameLengthProperty); err != nil {
			return err
		}
		return nil
	})
	_s_pulResponseFormatNameLenProp := func(ptr interface{}) { o.ResponseFormatNameLengthProperty = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.ResponseFormatNameLengthProperty, _s_pulResponseFormatNameLenProp, _ptr_pulResponseFormatNameLenProp); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdminFormatNameLength); err != nil {
		return err
	}
	_ptr_ppAdminFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppAdminFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			// XXX: for opaque unmarshaling
			if o.AdminFormatNameLength > 0 && sizeInfo[0] == 0 {
				sizeInfo[0] = uint64(o.AdminFormatNameLength)
			}
			var _AdminFormatName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _AdminFormatName_buf", sizeInfo[0])
			}
			_AdminFormatName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _AdminFormatName_buf {
				i1 := i1
				if err := w.ReadData(&_AdminFormatName_buf[i1]); err != nil {
					return err
				}
			}
			o.AdminFormatName = strings.TrimRight(string(utf16.Decode(_AdminFormatName_buf)), ndr.ZeroString)
			return nil
		})
		_s_ppAdminFormatName := func(ptr interface{}) { o.AdminFormatName = *ptr.(*string) }
		if err := w.ReadPointer(&o.AdminFormatName, _s_ppAdminFormatName, _ptr_ppAdminFormatName); err != nil {
			return err
		}
		return nil
	})
	_s_ppAdminFormatName := func(ptr interface{}) { o.AdminFormatName = *ptr.(*string) }
	if err := w.ReadPointer(&o.AdminFormatName, _s_ppAdminFormatName, _ptr_ppAdminFormatName); err != nil {
		return err
	}
	_ptr_pulAdminFormatNameLenProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.AdminFormatNameLengthProperty); err != nil {
			return err
		}
		return nil
	})
	_s_pulAdminFormatNameLenProp := func(ptr interface{}) { o.AdminFormatNameLengthProperty = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.AdminFormatNameLengthProperty, _s_pulAdminFormatNameLenProp, _ptr_pulAdminFormatNameLenProp); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationFormatNameLength); err != nil {
		return err
	}
	_ptr_ppDestFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppDestFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			// XXX: for opaque unmarshaling
			if o.DestinationFormatNameLength > 0 && sizeInfo[0] == 0 {
				sizeInfo[0] = uint64(o.DestinationFormatNameLength)
			}
			var _DestinationFormatName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _DestinationFormatName_buf", sizeInfo[0])
			}
			_DestinationFormatName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _DestinationFormatName_buf {
				i1 := i1
				if err := w.ReadData(&_DestinationFormatName_buf[i1]); err != nil {
					return err
				}
			}
			o.DestinationFormatName = strings.TrimRight(string(utf16.Decode(_DestinationFormatName_buf)), ndr.ZeroString)
			return nil
		})
		_s_ppDestFormatName := func(ptr interface{}) { o.DestinationFormatName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DestinationFormatName, _s_ppDestFormatName, _ptr_ppDestFormatName); err != nil {
			return err
		}
		return nil
	})
	_s_ppDestFormatName := func(ptr interface{}) { o.DestinationFormatName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DestinationFormatName, _s_ppDestFormatName, _ptr_ppDestFormatName); err != nil {
		return err
	}
	_ptr_pulDestFormatNameLenProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.DestinationFormatNameLengthProperty); err != nil {
			return err
		}
		return nil
	})
	_s_pulDestFormatNameLenProp := func(ptr interface{}) { o.DestinationFormatNameLengthProperty = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.DestinationFormatNameLengthProperty, _s_pulDestFormatNameLenProp, _ptr_pulDestFormatNameLenProp); err != nil {
		return err
	}
	if err := w.ReadData(&o.OrderingFormatNameLength); err != nil {
		return err
	}
	_ptr_ppOrderingFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppOrderingFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			// XXX: for opaque unmarshaling
			if o.OrderingFormatNameLength > 0 && sizeInfo[0] == 0 {
				sizeInfo[0] = uint64(o.OrderingFormatNameLength)
			}
			var _OrderingFormatName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _OrderingFormatName_buf", sizeInfo[0])
			}
			_OrderingFormatName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _OrderingFormatName_buf {
				i1 := i1
				if err := w.ReadData(&_OrderingFormatName_buf[i1]); err != nil {
					return err
				}
			}
			o.OrderingFormatName = strings.TrimRight(string(utf16.Decode(_OrderingFormatName_buf)), ndr.ZeroString)
			return nil
		})
		_s_ppOrderingFormatName := func(ptr interface{}) { o.OrderingFormatName = *ptr.(*string) }
		if err := w.ReadPointer(&o.OrderingFormatName, _s_ppOrderingFormatName, _ptr_ppOrderingFormatName); err != nil {
			return err
		}
		return nil
	})
	_s_ppOrderingFormatName := func(ptr interface{}) { o.OrderingFormatName = *ptr.(*string) }
	if err := w.ReadPointer(&o.OrderingFormatName, _s_ppOrderingFormatName, _ptr_ppOrderingFormatName); err != nil {
		return err
	}
	_ptr_pulOrderingFormatNameLenProp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.OrderingFormatNameLengthProperty); err != nil {
			return err
		}
		return nil
	})
	_s_pulOrderingFormatNameLenProp := func(ptr interface{}) { o.OrderingFormatNameLengthProperty = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.OrderingFormatNameLengthProperty, _s_pulOrderingFormatNameLenProp, _ptr_pulOrderingFormatNameLenProp); err != nil {
		return err
	}
	return nil
}

// TransferBufferV1_CreateCursor structure represents TransferBufferV1_TransferBufferV1 RPC union arm.
//
// It has following labels: 2
type TransferBufferV1_CreateCursor struct {
	// CreateCursor:  The CreateCursor member contains information for creating a cursor
	// which is used when receiving messages from a queue. The CreateCursor member is present
	// in the CACTransferBufferV1 structure when the value of the uTransferType member is
	// 0x00000002 (CACTB_CREATECURSOR). The CreateCursor member is not used by any of the
	// methods defined by the qmcomm and qmcomm2 interfaces.
	CreateCursor *CreateRemoteCursor `idl:"name:CreateCursor" json:"create_cursor"`
}

func (*TransferBufferV1_CreateCursor) is_TransferBufferV1_TransferBufferV1() {}

func (o *TransferBufferV1_CreateCursor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CreateCursor != nil {
		if err := o.CreateCursor.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CreateRemoteCursor{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransferBufferV1_CreateCursor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.CreateCursor == nil {
		o.CreateCursor = &CreateRemoteCursor{}
	}
	if err := o.CreateCursor.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TransferBufferV2 structure represents CACTransferBufferV2 RPC structure.
//
// The CACTransferBufferV2 structure is used to send and receive messages via MSMQ.
//
// Following is the layout of the CACTransferBufferV2 structure followed by descriptions
// of the structure members.
type TransferBufferV2 struct {
	// old:   The CACTransferBufferOld MUST be a CACTransferBufferV1, as defined in section
	// 2.2.3.2.
	Old *TransferBufferV1 `idl:"name:old" json:"old"`
	// pbFirstInXact:   The pbFirstInXact member MUST be a single byte. The pbFirstInXact
	// member MUST be set to a value of 0x00 (FALSE) when the associated message is not
	// the first message in a transaction. A value other than 0x00 MUST be interpreted as
	// indicating (TRUE) that the associated message is the first message in a transaction.
	FirstInTransaction uint8 `idl:"name:pbFirstInXact" json:"first_in_transaction"`
	// pbLastInXact:   The pbLastInXact member MUST be a single byte. The pbLastInXact member
	// MUST be set to a value of 0x00 (FALSE) when the associated message is not the last
	// message in a transaction. A value other than 0x00 MUST be interpreted as indicating
	// (TRUE) that the associated message is the last message in a transaction.
	LastInTransaction uint8 `idl:"name:pbLastInXact" json:"last_in_transaction"`
	// ppXactID:   The ppXactID member, if present, MUST be an OBJECTID structure, as specified
	// in [MS-MQMQ] section 2.2.8.
	TransactionID *mqmq.ObjectID `idl:"name:ppXactID" json:"transaction_id"`
}

func (o *TransferBufferV2) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransferBufferV2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Old != nil {
		if err := o.Old.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TransferBufferV1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	_ptr_pbFirstInXact := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.FirstInTransaction); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.FirstInTransaction, _ptr_pbFirstInXact); err != nil {
		return err
	}
	_ptr_pbLastInXact := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.LastInTransaction); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.LastInTransaction, _ptr_pbLastInXact); err != nil {
		return err
	}
	if o.TransactionID != nil {
		_ptr_ppXactID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.TransactionID != nil {
				_ptr_ppXactID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.TransactionID != nil {
						if err := o.TransactionID.MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&mqmq.ObjectID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.TransactionID, _ptr_ppXactID); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.TransactionID, _ptr_ppXactID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransferBufferV2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Old == nil {
		o.Old = &TransferBufferV1{}
	}
	if err := o.Old.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pbFirstInXact := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.FirstInTransaction); err != nil {
			return err
		}
		return nil
	})
	_s_pbFirstInXact := func(ptr interface{}) { o.FirstInTransaction = *ptr.(*uint8) }
	if err := w.ReadPointer(&o.FirstInTransaction, _s_pbFirstInXact, _ptr_pbFirstInXact); err != nil {
		return err
	}
	_ptr_pbLastInXact := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.LastInTransaction); err != nil {
			return err
		}
		return nil
	})
	_s_pbLastInXact := func(ptr interface{}) { o.LastInTransaction = *ptr.(*uint8) }
	if err := w.ReadPointer(&o.LastInTransaction, _s_pbLastInXact, _ptr_pbLastInXact); err != nil {
		return err
	}
	_ptr_ppXactID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppXactID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TransactionID == nil {
				o.TransactionID = &mqmq.ObjectID{}
			}
			if err := o.TransactionID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppXactID := func(ptr interface{}) { o.TransactionID = *ptr.(**mqmq.ObjectID) }
		if err := w.ReadPointer(&o.TransactionID, _s_ppXactID, _ptr_ppXactID); err != nil {
			return err
		}
		return nil
	})
	_s_ppXactID := func(ptr interface{}) { o.TransactionID = *ptr.(**mqmq.ObjectID) }
	if err := w.ReadPointer(&o.TransactionID, _s_ppXactID, _ptr_ppXactID); err != nil {
		return err
	}
	return nil
}

// ObjectFormat structure represents OBJECT_FORMAT RPC structure.
type ObjectFormat struct {
	ObjectType   uint32                     `idl:"name:ObjType" json:"object_type"`
	ObjectFormat *ObjectFormat_ObjectFormat `idl:"name:ObjectFormat;switch_is:ObjType" json:"object_format"`
}

func (o *ObjectFormat) xxx_PreparePayload(ctx context.Context) error {
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(2) {
		return fmt.Errorf("ObjectType is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectFormat) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ObjectType); err != nil {
		return err
	}
	_swObjectFormat := uint32(o.ObjectType)
	if o.ObjectFormat != nil {
		if err := o.ObjectFormat.MarshalUnionNDR(ctx, w, _swObjectFormat); err != nil {
			return err
		}
	} else {
		if err := (&ObjectFormat_ObjectFormat{}).MarshalUnionNDR(ctx, w, _swObjectFormat); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectFormat) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ObjectType); err != nil {
		return err
	}
	if o.ObjectFormat == nil {
		o.ObjectFormat = &ObjectFormat_ObjectFormat{}
	}
	_swObjectFormat := uint32(o.ObjectType)
	if err := o.ObjectFormat.UnmarshalUnionNDR(ctx, w, _swObjectFormat); err != nil {
		return err
	}
	return nil
}

// ObjectFormat_ObjectFormat structure represents OBJECT_FORMAT union anonymous member.
type ObjectFormat_ObjectFormat struct {
	// Types that are assignable to Value
	//
	// *ObjectFormat_QueueFormat
	Value is_ObjectFormat_ObjectFormat `json:"value"`
}

func (o *ObjectFormat_ObjectFormat) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ObjectFormat_QueueFormat:
		if value != nil {
			return value.QueueFormat
		}
	}
	return nil
}

type is_ObjectFormat_ObjectFormat interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ObjectFormat_ObjectFormat()
}

func (o *ObjectFormat_ObjectFormat) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ObjectFormat_QueueFormat:
		return uint32(1)
	}
	return uint32(0)
}

func (o *ObjectFormat_ObjectFormat) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*ObjectFormat_QueueFormat)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectFormat_QueueFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ObjectFormat_ObjectFormat) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &ObjectFormat_QueueFormat{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ObjectFormat_QueueFormat structure represents ObjectFormat_ObjectFormat RPC union arm.
//
// It has following labels: 1
type ObjectFormat_QueueFormat struct {
	QueueFormat *mqmq.QueueFormat `idl:"name:pQueueFormat" json:"queue_format"`
}

func (*ObjectFormat_QueueFormat) is_ObjectFormat_ObjectFormat() {}

func (o *ObjectFormat_QueueFormat) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.QueueFormat != nil {
		_ptr_pQueueFormat := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.QueueFormat != nil {
				if err := o.QueueFormat.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.QueueFormat{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.QueueFormat, _ptr_pQueueFormat); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectFormat_QueueFormat) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pQueueFormat := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.QueueFormat == nil {
			o.QueueFormat = &mqmq.QueueFormat{}
		}
		if err := o.QueueFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pQueueFormat := func(ptr interface{}) { o.QueueFormat = *ptr.(**mqmq.QueueFormat) }
	if err := w.ReadPointer(&o.QueueFormat, _s_pQueueFormat, _ptr_pQueueFormat); err != nil {
		return err
	}
	return nil
}

// Context structure represents PCTX_OPENREMOTE_HANDLE_TYPE RPC structure.
type Context dcetypes.ContextHandle

func (o *Context) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Context) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Context) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Context) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Queue structure represents RPC_QUEUE_HANDLE RPC structure.
type Queue dcetypes.ContextHandle

func (o *Queue) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Queue) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Queue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Queue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// InternalTransaction structure represents RPC_INT_XACT_HANDLE RPC structure.
type InternalTransaction dcetypes.ContextHandle

func (o *InternalTransaction) ContextHandle() *dcetypes.ContextHandle {
	return (*dcetypes.ContextHandle)(o)
}

func (o *InternalTransaction) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *InternalTransaction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *InternalTransaction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type xxx_DefaultQmcommClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultQmcommClient) GetRemoteQueueName(ctx context.Context, in *GetRemoteQueueNameRequest, opts ...dcerpc.CallOption) (*GetRemoteQueueNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetRemoteQueueNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) OpenRemoteQueue(ctx context.Context, in *OpenRemoteQueueRequest, opts ...dcerpc.CallOption) (*OpenRemoteQueueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenRemoteQueueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) CloseRemoteQueueContext(ctx context.Context, in *CloseRemoteQueueContextRequest, opts ...dcerpc.CallOption) (*CloseRemoteQueueContextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseRemoteQueueContextResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultQmcommClient) CreateRemoteCursor(ctx context.Context, in *CreateRemoteCursorRequest, opts ...dcerpc.CallOption) (*CreateRemoteCursorResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateRemoteCursorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) CreateObjectInternal(ctx context.Context, in *CreateObjectInternalRequest, opts ...dcerpc.CallOption) (*CreateObjectInternalResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateObjectInternalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) SetObjectSecurityInternal(ctx context.Context, in *SetObjectSecurityInternalRequest, opts ...dcerpc.CallOption) (*SetObjectSecurityInternalResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetObjectSecurityInternalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) GetObjectSecurityInternal(ctx context.Context, in *GetObjectSecurityInternalRequest, opts ...dcerpc.CallOption) (*GetObjectSecurityInternalResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetObjectSecurityInternalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...dcerpc.CallOption) (*DeleteObjectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) GetObjectProperties(ctx context.Context, in *GetObjectPropertiesRequest, opts ...dcerpc.CallOption) (*GetObjectPropertiesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetObjectPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) SetObjectProperties(ctx context.Context, in *SetObjectPropertiesRequest, opts ...dcerpc.CallOption) (*SetObjectPropertiesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetObjectPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) ObjectPathToObjectFormat(ctx context.Context, in *ObjectPathToObjectFormatRequest, opts ...dcerpc.CallOption) (*ObjectPathToObjectFormatResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ObjectPathToObjectFormatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) GetTMWhereabouts(ctx context.Context, in *GetTMWhereaboutsRequest, opts ...dcerpc.CallOption) (*GetTMWhereaboutsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetTMWhereaboutsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) EnlistTransaction(ctx context.Context, in *EnlistTransactionRequest, opts ...dcerpc.CallOption) (*EnlistTransactionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnlistTransactionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) EnlistInternalTransaction(ctx context.Context, in *EnlistInternalTransactionRequest, opts ...dcerpc.CallOption) (*EnlistInternalTransactionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnlistInternalTransactionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) CommitTransaction(ctx context.Context, in *CommitTransactionRequest, opts ...dcerpc.CallOption) (*CommitTransactionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CommitTransactionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) AbortTransaction(ctx context.Context, in *AbortTransactionRequest, opts ...dcerpc.CallOption) (*AbortTransactionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AbortTransactionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) OpenQueueInternal(ctx context.Context, in *OpenQueueInternalRequest, opts ...dcerpc.CallOption) (*OpenQueueInternalResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenQueueInternalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) CloseHandle(ctx context.Context, in *CloseHandleRequest, opts ...dcerpc.CallOption) (*CloseHandleResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseHandleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) CloseCursor(ctx context.Context, in *CloseCursorRequest, opts ...dcerpc.CallOption) (*CloseCursorResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseCursorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) SetCursorProperties(ctx context.Context, in *SetCursorPropertiesRequest, opts ...dcerpc.CallOption) (*SetCursorPropertiesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetCursorPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) HandleToFormatName(ctx context.Context, in *HandleToFormatNameRequest, opts ...dcerpc.CallOption) (*HandleToFormatNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &HandleToFormatNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) PurgeQueue(ctx context.Context, in *PurgeQueueRequest, opts ...dcerpc.CallOption) (*PurgeQueueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PurgeQueueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) QueryQMRegistryInternal(ctx context.Context, in *QueryQMRegistryInternalRequest, opts ...dcerpc.CallOption) (*QueryQMRegistryInternalResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryQMRegistryInternalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) GetRTQMServerPort(ctx context.Context, in *GetRTQMServerPortRequest, opts ...dcerpc.CallOption) (*GetRTQMServerPortResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetRTQMServerPortResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmcommClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQmcommClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewQmcommClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QmcommClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QmcommSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultQmcommClient{cc: cc}, nil
}

// xxx_GetRemoteQueueNameOperation structure represents the R_QMGetRemoteQueueName operation
type xxx_GetRemoteQueueNameOperation struct {
	QueueID         uint32 `idl:"name:pQueue" json:"queue_id"`
	RemoteQueueName string `idl:"name:lplpRemoteQueueName;string;pointer:ptr" json:"remote_queue_name"`
	Return          int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRemoteQueueNameOperation) OpNum() int { return 1 }

func (o *xxx_GetRemoteQueueNameOperation) OpName() string { return "/qmcomm/v1/R_QMGetRemoteQueueName" }

func (o *xxx_GetRemoteQueueNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteQueueNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QueueID); err != nil {
			return err
		}
	}
	// lplpRemoteQueueName {in, out} (1:{string, pointer=ptr}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.RemoteQueueName != "" {
			_ptr_lplpRemoteQueueName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RemoteQueueName != "" {
					_ptr_lplpRemoteQueueName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.RemoteQueueName); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.RemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
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

func (o *xxx_GetRemoteQueueNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QueueID); err != nil {
			return err
		}
	}
	// lplpRemoteQueueName {in, out} (1:{string, pointer=ptr}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_lplpRemoteQueueName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_lplpRemoteQueueName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteQueueName); err != nil {
					return err
				}
				return nil
			})
			_s_lplpRemoteQueueName := func(ptr interface{}) { o.RemoteQueueName = *ptr.(*string) }
			if err := w.ReadPointer(&o.RemoteQueueName, _s_lplpRemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
				return err
			}
			return nil
		})
		_s_lplpRemoteQueueName := func(ptr interface{}) { o.RemoteQueueName = *ptr.(*string) }
		if err := w.ReadPointer(&o.RemoteQueueName, _s_lplpRemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteQueueNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteQueueNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lplpRemoteQueueName {in, out} (1:{string, pointer=ptr}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.RemoteQueueName != "" {
			_ptr_lplpRemoteQueueName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RemoteQueueName != "" {
					_ptr_lplpRemoteQueueName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.RemoteQueueName); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.RemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
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

func (o *xxx_GetRemoteQueueNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lplpRemoteQueueName {in, out} (1:{string, pointer=ptr}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_lplpRemoteQueueName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_lplpRemoteQueueName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteQueueName); err != nil {
					return err
				}
				return nil
			})
			_s_lplpRemoteQueueName := func(ptr interface{}) { o.RemoteQueueName = *ptr.(*string) }
			if err := w.ReadPointer(&o.RemoteQueueName, _s_lplpRemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
				return err
			}
			return nil
		})
		_s_lplpRemoteQueueName := func(ptr interface{}) { o.RemoteQueueName = *ptr.(*string) }
		if err := w.ReadPointer(&o.RemoteQueueName, _s_lplpRemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
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

// GetRemoteQueueNameRequest structure represents the R_QMGetRemoteQueueName operation request
type GetRemoteQueueNameRequest struct {
	// pQueue:  MUST be a DWORD that contains a queue context value obtained from the cli_pQMQueue
	// member of the structure returned by the rpc_ACCreateCursorEx method of the qmcomm2
	// interface. See section 4.4 for an example illustrating this value being obtained.
	QueueID uint32 `idl:"name:pQueue" json:"queue_id"`
	// lplpRemoteQueueName:  A pointer to a buffer to receive the null-terminated name
	// of the remote queue associated with pQueue. On input, this value MUST be NULL.
	RemoteQueueName string `idl:"name:lplpRemoteQueueName;string;pointer:ptr" json:"remote_queue_name"`
}

func (o *GetRemoteQueueNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRemoteQueueNameOperation) *xxx_GetRemoteQueueNameOperation {
	if op == nil {
		op = &xxx_GetRemoteQueueNameOperation{}
	}
	if o == nil {
		return op
	}
	op.QueueID = o.QueueID
	op.RemoteQueueName = o.RemoteQueueName
	return op
}

func (o *GetRemoteQueueNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRemoteQueueNameOperation) {
	if o == nil {
		return
	}
	o.QueueID = op.QueueID
	o.RemoteQueueName = op.RemoteQueueName
}
func (o *GetRemoteQueueNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRemoteQueueNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRemoteQueueNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRemoteQueueNameResponse structure represents the R_QMGetRemoteQueueName operation response
type GetRemoteQueueNameResponse struct {
	// lplpRemoteQueueName:  A pointer to a buffer to receive the null-terminated name
	// of the remote queue associated with pQueue. On input, this value MUST be NULL.
	RemoteQueueName string `idl:"name:lplpRemoteQueueName;string;pointer:ptr" json:"remote_queue_name"`
	// Return: The R_QMGetRemoteQueueName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRemoteQueueNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRemoteQueueNameOperation) *xxx_GetRemoteQueueNameOperation {
	if op == nil {
		op = &xxx_GetRemoteQueueNameOperation{}
	}
	if o == nil {
		return op
	}
	op.RemoteQueueName = o.RemoteQueueName
	op.Return = o.Return
	return op
}

func (o *GetRemoteQueueNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRemoteQueueNameOperation) {
	if o == nil {
		return
	}
	o.RemoteQueueName = op.RemoteQueueName
	o.Return = op.Return
}
func (o *GetRemoteQueueNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRemoteQueueNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRemoteQueueNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenRemoteQueueOperation structure represents the R_QMOpenRemoteQueue operation
type xxx_OpenRemoteQueueOperation struct {
	Context          *Context          `idl:"name:pphContext" json:"context"`
	ContextID        uint32            `idl:"name:pdwContext" json:"context_id"`
	QueueFormat      *mqmq.QueueFormat `idl:"name:pQueueFormat;pointer:unique" json:"queue_format"`
	CallingProcessID uint32            `idl:"name:dwCallingProcessID" json:"calling_process_id"`
	DesiredAccess    uint32            `idl:"name:dwDesiredAccess" json:"desired_access"`
	ShareMode        uint32            `idl:"name:dwShareMode" json:"share_mode"`
	ClientGUID       *dtyp.GUID        `idl:"name:pLicGuid" json:"client_guid"`
	MQS              uint32            `idl:"name:dwMQS" json:"mqs"`
	QueueID          uint32            `idl:"name:dwpQueue" json:"queue_id"`
	Queue            uint32            `idl:"name:phQueue" json:"queue"`
	Return           int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenRemoteQueueOperation) OpNum() int { return 2 }

func (o *xxx_OpenRemoteQueueOperation) OpName() string { return "/qmcomm/v1/R_QMOpenRemoteQueue" }

func (o *xxx_OpenRemoteQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenRemoteQueueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pQueueFormat {in} (1:{pointer=unique}*(1))(2:{alias=QUEUE_FORMAT}(struct))
	{
		if o.QueueFormat != nil {
			_ptr_pQueueFormat := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueFormat != nil {
					if err := o.QueueFormat.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqmq.QueueFormat{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueueFormat, _ptr_pQueueFormat); err != nil {
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
	// dwCallingProcessID {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CallingProcessID); err != nil {
			return err
		}
	}
	// dwDesiredAccess {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	// dwShareMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ShareMode); err != nil {
			return err
		}
	}
	// pLicGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClientGUID != nil {
			if err := o.ClientGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwMQS {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MQS); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenRemoteQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pQueueFormat {in} (1:{pointer=unique}*(1))(2:{alias=QUEUE_FORMAT}(struct))
	{
		_ptr_pQueueFormat := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueFormat == nil {
				o.QueueFormat = &mqmq.QueueFormat{}
			}
			if err := o.QueueFormat.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pQueueFormat := func(ptr interface{}) { o.QueueFormat = *ptr.(**mqmq.QueueFormat) }
		if err := w.ReadPointer(&o.QueueFormat, _s_pQueueFormat, _ptr_pQueueFormat); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwCallingProcessID {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CallingProcessID); err != nil {
			return err
		}
	}
	// dwDesiredAccess {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	// dwShareMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ShareMode); err != nil {
			return err
		}
	}
	// pLicGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClientGUID == nil {
			o.ClientGUID = &dtyp.GUID{}
		}
		if err := o.ClientGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwMQS {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MQS); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenRemoteQueueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenRemoteQueueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_OPENREMOTE_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ContextID); err != nil {
			return err
		}
	}
	// dwpQueue {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QueueID); err != nil {
			return err
		}
	}
	// phQueue {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Queue); err != nil {
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

func (o *xxx_OpenRemoteQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_OPENREMOTE_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ContextID); err != nil {
			return err
		}
	}
	// dwpQueue {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QueueID); err != nil {
			return err
		}
	}
	// phQueue {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Queue); err != nil {
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

// OpenRemoteQueueRequest structure represents the R_QMOpenRemoteQueue operation request
type OpenRemoteQueueRequest struct {
	// pQueueFormat: A QUEUE_FORMAT ([MS-MQMQ] section 2.2.7) structure that identifies
	// the queue to be opened. It MUST NOT be NULL and MUST conform to the format name syntax
	// rules defined in [MS-MQMQ]. It MUST NOT be a distribution list or multicast format
	// name. For direct format names, the protocol MUST NOT be HTTP.
	QueueFormat *mqmq.QueueFormat `idl:"name:pQueueFormat;pointer:unique" json:"queue_format"`
	// dwCallingProcessID:  MUST be ignored. Clients MAY pass 0x00000000.<15>
	CallingProcessID uint32 `idl:"name:dwCallingProcessID" json:"calling_process_id"`
	// dwDesiredAccess:  A DWORD that specifies the access mode requested for the queue.
	// The access mode defines the set of operations that can be invoked using the returned
	// queue handle. The value MUST be one of the following:
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_RECEIVE_ACCESS 0x00000001 | The returned queue handle MUST only permit message peek, message receive (peek   |
	//	|                              | and delete), and queue purge operations.                                         |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_PEEK_ACCESS 0x00000020    | The returned queue handle MUST only permit message peek operations.              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	DesiredAccess uint32 `idl:"name:dwDesiredAccess" json:"desired_access"`
	// dwShareMode:  Specifies the exclusivity level for the opened queue. The value MUST
	// be one of the following:
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                     MEANING                                      |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_DENY_NONE 0x00000000          | The queue is not opened exclusively.                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_DENY_RECEIVE_SHARE 0x00000001 | The queue is to be opened for exclusive read access. If the queue has already    |
	//	|                                  | been opened for read access, the server MUST return STATUS_SHARING_VIOLATION     |
	//	|                                  | (0xc0000043). If the queue is opened successfully for exclusive read access,     |
	//	|                                  | subsequent attempts to open the same queue for read access MUST return           |
	//	|                                  | STATUS_SHARING_VIOLATION (0xc0000043) until the queue has been closed.           |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	ShareMode uint32 `idl:"name:dwShareMode" json:"share_mode"`
	// pLicGuid:  MUST be a pointer to a valid GUID that uniquely identifies the client.<16><17>
	// The server MAY ignore this parameter.<18>
	ClientGUID *dtyp.GUID `idl:"name:pLicGuid" json:"client_guid"`
	// dwMQS:  MUST be set by clients to indicate the client operating system category.
	// Servers MAY ignore this value.<19> The following values are defined:
	//
	//	+----------------+---------------------------------------------------------------+
	//	|                |          ARE SERVER CONNECTION LICENSING LIMITATIONS          |
	//	|     VALUE      |                       ENFORCED?/MEANING                       |
	//	|                |                                                               |
	//	+----------------+---------------------------------------------------------------+
	//	+----------------+---------------------------------------------------------------+
	//	| 0x00000000<20> | None. The operating system (OS) version is not declared.      |
	//	+----------------+---------------------------------------------------------------+
	//	| 0x00000100     | Yes. For supported operating systems.<21>                     |
	//	+----------------+---------------------------------------------------------------+
	//	| 0x00000200     | Yes. For supported operating systems.<22>                     |
	//	+----------------+---------------------------------------------------------------+
	//	| 0x00000300     | Yes. For supported operating systems.<23>                     |
	//	+----------------+---------------------------------------------------------------+
	//	| 0x00000400     | No. For supported operating systems.<24>                      |
	//	+----------------+---------------------------------------------------------------+
	//	| 0x00000500     | No. For supported operating systems.<25>                      |
	//	+----------------+---------------------------------------------------------------+
	MQS uint32 `idl:"name:dwMQS" json:"mqs"`
}

func (o *OpenRemoteQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenRemoteQueueOperation) *xxx_OpenRemoteQueueOperation {
	if op == nil {
		op = &xxx_OpenRemoteQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.QueueFormat = o.QueueFormat
	op.CallingProcessID = o.CallingProcessID
	op.DesiredAccess = o.DesiredAccess
	op.ShareMode = o.ShareMode
	op.ClientGUID = o.ClientGUID
	op.MQS = o.MQS
	return op
}

func (o *OpenRemoteQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenRemoteQueueOperation) {
	if o == nil {
		return
	}
	o.QueueFormat = op.QueueFormat
	o.CallingProcessID = op.CallingProcessID
	o.DesiredAccess = op.DesiredAccess
	o.ShareMode = op.ShareMode
	o.ClientGUID = op.ClientGUID
	o.MQS = op.MQS
}
func (o *OpenRemoteQueueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenRemoteQueueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenRemoteQueueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenRemoteQueueResponse structure represents the R_QMOpenRemoteQueue operation response
type OpenRemoteQueueResponse struct {
	// pphContext:  A pointer to a variable to receive the PCTX_OPENREMOTE_HANDLE_TYPE
	// (section 2.2.1.1.3) context handle.
	Context *Context `idl:"name:pphContext" json:"context"`
	// pdwContext: A pointer to a variable to receive the value of the Handle attribute
	// for the new OpenQueueDescriptor ([MS-MQDMPR] section 3.1.1.16) ADM element instance
	// created by this method. It MUST NOT be NULL.
	ContextID uint32 `idl:"name:pdwContext" json:"context_id"`
	// dwpQueue: A pointer to a variable to receive a value that identifies the new OpenQueueDescriptor
	// ADM element instance created by this method, as specified in the processing rules
	// section for this method. It MUST NOT be NULL.
	QueueID uint32 `idl:"name:dwpQueue" json:"queue_id"`
	// phQueue: A pointer to a variable to receive the value of the Handle attribute for
	// the new OpenQueueDescriptor ADM element instance created by this method. It MUST
	// NOT be NULL.
	Queue uint32 `idl:"name:phQueue" json:"queue"`
	// Return: The R_QMOpenRemoteQueue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenRemoteQueueResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenRemoteQueueOperation) *xxx_OpenRemoteQueueOperation {
	if op == nil {
		op = &xxx_OpenRemoteQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.ContextID = o.ContextID
	op.QueueID = o.QueueID
	op.Queue = o.Queue
	op.Return = o.Return
	return op
}

func (o *OpenRemoteQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenRemoteQueueOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.ContextID = op.ContextID
	o.QueueID = op.QueueID
	o.Queue = op.Queue
	o.Return = op.Return
}
func (o *OpenRemoteQueueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenRemoteQueueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenRemoteQueueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseRemoteQueueContextOperation structure represents the R_QMCloseRemoteQueueContext operation
type xxx_CloseRemoteQueueContextOperation struct {
	Context *Context `idl:"name:pphContext" json:"context"`
}

func (o *xxx_CloseRemoteQueueContextOperation) OpNum() int { return 3 }

func (o *xxx_CloseRemoteQueueContextOperation) OpName() string {
	return "/qmcomm/v1/R_QMCloseRemoteQueueContext"
}

func (o *xxx_CloseRemoteQueueContextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseRemoteQueueContextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_OPENREMOTE_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseRemoteQueueContextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_OPENREMOTE_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseRemoteQueueContextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseRemoteQueueContextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_OPENREMOTE_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Context{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseRemoteQueueContextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCTX_OPENREMOTE_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Context{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// CloseRemoteQueueContextRequest structure represents the R_QMCloseRemoteQueueContext operation request
type CloseRemoteQueueContextRequest struct {
	// pphContext:  An RPC context handle as defined in [MS-RPCE] section 2. This handle
	// MUST have been acquired from the pphContext parameter of the R_QMOpenRemoteQueue
	// method.
	Context *Context `idl:"name:pphContext" json:"context"`
}

func (o *CloseRemoteQueueContextRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseRemoteQueueContextOperation) *xxx_CloseRemoteQueueContextOperation {
	if op == nil {
		op = &xxx_CloseRemoteQueueContextOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *CloseRemoteQueueContextRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseRemoteQueueContextOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *CloseRemoteQueueContextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseRemoteQueueContextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseRemoteQueueContextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseRemoteQueueContextResponse structure represents the R_QMCloseRemoteQueueContext operation response
type CloseRemoteQueueContextResponse struct {
	// pphContext:  An RPC context handle as defined in [MS-RPCE] section 2. This handle
	// MUST have been acquired from the pphContext parameter of the R_QMOpenRemoteQueue
	// method.
	Context *Context `idl:"name:pphContext" json:"context"`
}

func (o *CloseRemoteQueueContextResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseRemoteQueueContextOperation) *xxx_CloseRemoteQueueContextOperation {
	if op == nil {
		op = &xxx_CloseRemoteQueueContextOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *CloseRemoteQueueContextResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseRemoteQueueContextOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *CloseRemoteQueueContextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseRemoteQueueContextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseRemoteQueueContextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateRemoteCursorOperation structure represents the R_QMCreateRemoteCursor operation
type xxx_CreateRemoteCursorOperation struct {
	Buffer1 *TransferBufferV1 `idl:"name:ptb1" json:"buffer1"`
	Queue   uint32            `idl:"name:hQueue" json:"queue"`
	Cursor  uint32            `idl:"name:phCursor" json:"cursor"`
	Return  int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateRemoteCursorOperation) OpNum() int { return 4 }

func (o *xxx_CreateRemoteCursorOperation) OpName() string { return "/qmcomm/v1/R_QMCreateRemoteCursor" }

func (o *xxx_CreateRemoteCursorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRemoteCursorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ptb1 {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.Buffer1 != nil {
			if err := o.Buffer1.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TransferBufferV1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Queue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRemoteCursorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ptb1 {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.Buffer1 == nil {
			o.Buffer1 = &TransferBufferV1{}
		}
		if err := o.Buffer1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Queue); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRemoteCursorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRemoteCursorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phCursor {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cursor); err != nil {
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

func (o *xxx_CreateRemoteCursorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phCursor {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cursor); err != nil {
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

// CreateRemoteCursorRequest structure represents the R_QMCreateRemoteCursor operation request
type CreateRemoteCursorRequest struct {
	// ptb1:  MUST be ignored. Clients SHOULD pass NULL.<29>
	Buffer1 *TransferBufferV1 `idl:"name:ptb1" json:"buffer1"`
	// hQueue: A DWORD that contains the value of the Handle attribute of an OpenQueueDescriptor
	// ([MS-MQDMPR] section 3.1.1.16) ADM element instance. The client obtains this value
	// from either the pcc.srv_hACQueue out-parameter of rpc_ACCreateCursorEx or the phQueue
	// out-parameter of R_QMOpenRemoteQueue.
	Queue uint32 `idl:"name:hQueue" json:"queue"`
}

func (o *CreateRemoteCursorRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateRemoteCursorOperation) *xxx_CreateRemoteCursorOperation {
	if op == nil {
		op = &xxx_CreateRemoteCursorOperation{}
	}
	if o == nil {
		return op
	}
	op.Buffer1 = o.Buffer1
	op.Queue = o.Queue
	return op
}

func (o *CreateRemoteCursorRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateRemoteCursorOperation) {
	if o == nil {
		return
	}
	o.Buffer1 = op.Buffer1
	o.Queue = op.Queue
}
func (o *CreateRemoteCursorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateRemoteCursorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRemoteCursorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateRemoteCursorResponse structure represents the R_QMCreateRemoteCursor operation response
type CreateRemoteCursorResponse struct {
	// phCursor: A pointer to a DWORD to receive the value of the Handle attribute of the
	// Cursor ([MS-MQDMPR] section 3.2) ADM element instance that references the created
	// cursor. It MUST NOT be NULL on input.
	Cursor uint32 `idl:"name:phCursor" json:"cursor"`
	// Return: The R_QMCreateRemoteCursor return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateRemoteCursorResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateRemoteCursorOperation) *xxx_CreateRemoteCursorOperation {
	if op == nil {
		op = &xxx_CreateRemoteCursorOperation{}
	}
	if o == nil {
		return op
	}
	op.Cursor = o.Cursor
	op.Return = o.Return
	return op
}

func (o *CreateRemoteCursorResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateRemoteCursorOperation) {
	if o == nil {
		return
	}
	o.Cursor = op.Cursor
	o.Return = op.Return
}
func (o *CreateRemoteCursorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateRemoteCursorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRemoteCursorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateObjectInternalOperation structure represents the R_QMCreateObjectInternal operation
type xxx_CreateObjectInternalOperation struct {
	ObjectType             uint32                  `idl:"name:dwObjectType" json:"object_type"`
	PathName               string                  `idl:"name:lpwcsPathName;string" json:"path_name"`
	SecurityDescriptorSize uint32                  `idl:"name:SDSize" json:"security_descriptor_size"`
	SecurityDescriptor     []byte                  `idl:"name:pSecurityDescriptor;size_is:(SDSize);pointer:unique" json:"security_descriptor"`
	CreatePartition        uint32                  `idl:"name:cp" json:"create_partition"`
	Property               []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var                    []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	Return                 int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateObjectInternalOperation) OpNum() int { return 6 }

func (o *xxx_CreateObjectInternalOperation) OpName() string {
	return "/qmcomm/v1/R_QMCreateObjectInternal"
}

func (o *xxx_CreateObjectInternalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.SecurityDescriptorSize == 0 {
		o.SecurityDescriptorSize = uint32(len(o.SecurityDescriptor))
	}
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.SecurityDescriptorSize > uint32(524288) {
		return fmt.Errorf("SecurityDescriptorSize is out of range")
	}
	if o.CreatePartition < uint32(1) || o.CreatePartition > uint32(128) {
		return fmt.Errorf("CreatePartition is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObjectInternalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// lpwcsPathName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	// SDSize {in} (1:{range=(0,524288), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityDescriptorSize); err != nil {
			return err
		}
	}
	// pSecurityDescriptor {in} (1:{pointer=unique}*(1)[dim:0,size_is=SDSize](uchar))
	{
		if o.SecurityDescriptor != nil || o.SecurityDescriptorSize > 0 {
			_ptr_pSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.SecurityDescriptorSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.SecurityDescriptor {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.SecurityDescriptor[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.SecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
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
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp])(2:{alias=DWORD}(uint32))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Property {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Property[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Property); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		}
	}
	// apVar {in} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Var {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Var[i1] != nil {
				if err := o.Var[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Var); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObjectInternalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// lpwcsPathName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
	}
	// SDSize {in} (1:{range=(0,524288), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityDescriptorSize); err != nil {
			return err
		}
	}
	// pSecurityDescriptor {in} (1:{pointer=unique}*(1)[dim:0,size_is=SDSize](uchar))
	{
		_ptr_pSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SecurityDescriptor", sizeInfo[0])
			}
			o.SecurityDescriptor = make([]byte, sizeInfo[0])
			for i1 := range o.SecurityDescriptor {
				i1 := i1
				if err := w.ReadData(&o.SecurityDescriptor[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pSecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.SecurityDescriptor, _s_pSecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp])(2:{alias=DWORD}(uint32))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Property", sizeInfo[0])
		}
		o.Property = make([]uint32, sizeInfo[0])
		for i1 := range o.Property {
			i1 := i1
			if err := w.ReadData(&o.Property[i1]); err != nil {
				return err
			}
		}
	}
	// apVar {in} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Var", sizeInfo[0])
		}
		o.Var = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Var {
			i1 := i1
			if o.Var[i1] == nil {
				o.Var[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Var[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObjectInternalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObjectInternalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateObjectInternalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateObjectInternalRequest structure represents the R_QMCreateObjectInternal operation request
type CreateObjectInternalRequest struct {
	// dwObjectType:  MUST be 0x00000001 in order to specify a queue.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// lpwcsPathName:  MUST be a pointer to a null-terminated string containing a path
	// name for the queue to be created. The path name MUST identify a private queue local
	// to the supporting server by including "." as the computer name or by using the supporting
	// server computer name.
	PathName string `idl:"name:lpwcsPathName;string" json:"path_name"`
	// SDSize:  MUST be set to the byte length of the SECURITY_DESCRIPTOR buffer pointed
	// to by pSecurityDescriptor. If pSecurityDescriptor is NULL, this parameter MUST be
	// 0x00000000.
	SecurityDescriptorSize uint32 `idl:"name:SDSize" json:"security_descriptor_size"`
	// pSecurityDescriptor:  Must be a pointer to an array of bytes containing a SECURITY_DESCRIPTOR
	// structure. The SECURITY_DESCRIPTOR specifies the initial security configuration for
	// the queue to be created. This value can be NULL, in which case the server MUST provide
	// a default security configuration for the new queue. The SECURITY_DESCRIPTOR structure
	// is defined in [MS-DTYP] section 2.4.6.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(SDSize);pointer:unique" json:"security_descriptor"`
	// cp:  MUST be set to the size (in elements) of the arrays aProp and apVar. The arrays
	// aProp and apVar MUST have an identical number of elements and MUST contain at least
	// one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  MUST be an array of queue property identifiers that, together with the apVar
	// array, specify the initial queue property values for the new queue. Each element
	// MUST specify a value from the queue property identifiers table defined in [MS-MQMQ]
	// section 2.3.1. Each element MUST specify the property identifier for the corresponding
	// property value at the same element index in apVar and MUST contain at least one element.
	// Each element MUST contain a queue property identifier; identifiers for other properties
	// are not permitted.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar:  MUST be an array that specifies the property values to associate with the
	// new queue. Each element MUST specify the property value for the corresponding property
	// identifier at the same element index in aProp and MUST contain at least one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
}

func (o *CreateObjectInternalRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateObjectInternalOperation) *xxx_CreateObjectInternalOperation {
	if op == nil {
		op = &xxx_CreateObjectInternalOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.PathName = o.PathName
	op.SecurityDescriptorSize = o.SecurityDescriptorSize
	op.SecurityDescriptor = o.SecurityDescriptor
	op.CreatePartition = o.CreatePartition
	op.Property = o.Property
	op.Var = o.Var
	return op
}

func (o *CreateObjectInternalRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateObjectInternalOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.PathName = op.PathName
	o.SecurityDescriptorSize = op.SecurityDescriptorSize
	o.SecurityDescriptor = op.SecurityDescriptor
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
}
func (o *CreateObjectInternalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateObjectInternalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateObjectInternalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateObjectInternalResponse structure represents the R_QMCreateObjectInternal operation response
type CreateObjectInternalResponse struct {
	// Return: The R_QMCreateObjectInternal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateObjectInternalResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateObjectInternalOperation) *xxx_CreateObjectInternalOperation {
	if op == nil {
		op = &xxx_CreateObjectInternalOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *CreateObjectInternalResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateObjectInternalOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CreateObjectInternalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateObjectInternalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateObjectInternalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetObjectSecurityInternalOperation structure represents the R_QMSetObjectSecurityInternal operation
type xxx_SetObjectSecurityInternalOperation struct {
	ObjectFormat           *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
	SecurityInformation    uint32        `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptorSize uint32        `idl:"name:SDSize" json:"security_descriptor_size"`
	SecurityDescriptor     []byte        `idl:"name:pSecurityDescriptor;size_is:(SDSize);pointer:unique" json:"security_descriptor"`
	Return                 int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_SetObjectSecurityInternalOperation) OpNum() int { return 7 }

func (o *xxx_SetObjectSecurityInternalOperation) OpName() string {
	return "/qmcomm/v1/R_QMSetObjectSecurityInternal"
}

func (o *xxx_SetObjectSecurityInternalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.SecurityDescriptorSize == 0 {
		o.SecurityDescriptorSize = uint32(len(o.SecurityDescriptor))
	}
	if o.SecurityDescriptorSize > uint32(524288) {
		return fmt.Errorf("SecurityDescriptorSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectSecurityInternalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pObjectFormat {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat != nil {
			if err := o.ObjectFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// SDSize {in} (1:{range=(0,524288), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityDescriptorSize); err != nil {
			return err
		}
	}
	// pSecurityDescriptor {in} (1:{pointer=unique}*(1)[dim:0,size_is=SDSize](uchar))
	{
		if o.SecurityDescriptor != nil || o.SecurityDescriptorSize > 0 {
			_ptr_pSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.SecurityDescriptorSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.SecurityDescriptor {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.SecurityDescriptor[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.SecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
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

func (o *xxx_SetObjectSecurityInternalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pObjectFormat {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat == nil {
			o.ObjectFormat = &ObjectFormat{}
		}
		if err := o.ObjectFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// SDSize {in} (1:{range=(0,524288), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityDescriptorSize); err != nil {
			return err
		}
	}
	// pSecurityDescriptor {in} (1:{pointer=unique}*(1)[dim:0,size_is=SDSize](uchar))
	{
		_ptr_pSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SecurityDescriptor", sizeInfo[0])
			}
			o.SecurityDescriptor = make([]byte, sizeInfo[0])
			for i1 := range o.SecurityDescriptor {
				i1 := i1
				if err := w.ReadData(&o.SecurityDescriptor[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pSecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.SecurityDescriptor, _s_pSecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectSecurityInternalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectSecurityInternalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetObjectSecurityInternalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetObjectSecurityInternalRequest structure represents the R_QMSetObjectSecurityInternal operation request
type SetObjectSecurityInternalRequest struct {
	// pObjectFormat:  MUST point to an OBJECT_FORMAT structure that identifies an existing
	// local private queue on the supporting server for which the security configuration
	// will be updated. This MUST NOT be NULL. The ObjType member of the structure MUST
	// be 0x00000001. The pQueueFormat member MUST NOT be NULL.
	ObjectFormat *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
	// SecurityInformation:  MUST contain a value from the SECURITY_INFORMATION enumeration
	// which indicates the portions of the provided SECURITY_DESCRIPTOR to be applied to
	// the queue identified by pObjectFormat. The SECURITY_INFORMATION enumeration is defined
	// in [MS-MQMQ] section 2.2.3.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// SDSize:  MUST be set to the byte length of the buffer pointed to by pSecurityDescriptor.
	SecurityDescriptorSize uint32 `idl:"name:SDSize" json:"security_descriptor_size"`
	// pSecurityDescriptor:  MUST be a pointer to an array of bytes containing a SECURITY_DESCRIPTOR
	// structure (see [MS-DTYP] section 2.4.6).
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(SDSize);pointer:unique" json:"security_descriptor"`
}

func (o *SetObjectSecurityInternalRequest) xxx_ToOp(ctx context.Context, op *xxx_SetObjectSecurityInternalOperation) *xxx_SetObjectSecurityInternalOperation {
	if op == nil {
		op = &xxx_SetObjectSecurityInternalOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectFormat = o.ObjectFormat
	op.SecurityInformation = o.SecurityInformation
	op.SecurityDescriptorSize = o.SecurityDescriptorSize
	op.SecurityDescriptor = o.SecurityDescriptor
	return op
}

func (o *SetObjectSecurityInternalRequest) xxx_FromOp(ctx context.Context, op *xxx_SetObjectSecurityInternalOperation) {
	if o == nil {
		return
	}
	o.ObjectFormat = op.ObjectFormat
	o.SecurityInformation = op.SecurityInformation
	o.SecurityDescriptorSize = op.SecurityDescriptorSize
	o.SecurityDescriptor = op.SecurityDescriptor
}
func (o *SetObjectSecurityInternalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetObjectSecurityInternalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetObjectSecurityInternalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetObjectSecurityInternalResponse structure represents the R_QMSetObjectSecurityInternal operation response
type SetObjectSecurityInternalResponse struct {
	// Return: The R_QMSetObjectSecurityInternal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetObjectSecurityInternalResponse) xxx_ToOp(ctx context.Context, op *xxx_SetObjectSecurityInternalOperation) *xxx_SetObjectSecurityInternalOperation {
	if op == nil {
		op = &xxx_SetObjectSecurityInternalOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetObjectSecurityInternalResponse) xxx_FromOp(ctx context.Context, op *xxx_SetObjectSecurityInternalOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetObjectSecurityInternalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetObjectSecurityInternalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetObjectSecurityInternalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetObjectSecurityInternalOperation structure represents the R_QMGetObjectSecurityInternal operation
type xxx_GetObjectSecurityInternalOperation struct {
	ObjectFormat         *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
	RequestedInformation uint32        `idl:"name:RequestedInformation" json:"requested_information"`
	SecurityDescriptor   []byte        `idl:"name:pSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	Length               uint32        `idl:"name:nLength" json:"length"`
	LengthNeeded         uint32        `idl:"name:lpnLengthNeeded" json:"length_needed"`
	Return               int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetObjectSecurityInternalOperation) OpNum() int { return 8 }

func (o *xxx_GetObjectSecurityInternalOperation) OpName() string {
	return "/qmcomm/v1/R_QMGetObjectSecurityInternal"
}

func (o *xxx_GetObjectSecurityInternalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Length > uint32(524288) {
		return fmt.Errorf("Length is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectSecurityInternalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pObjectFormat {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat != nil {
			if err := o.ObjectFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// RequestedInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedInformation); err != nil {
			return err
		}
	}
	// nLength {in} (1:{range=(0,524288), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectSecurityInternalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pObjectFormat {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat == nil {
			o.ObjectFormat = &ObjectFormat{}
		}
		if err := o.ObjectFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RequestedInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedInformation); err != nil {
			return err
		}
	}
	// nLength {in} (1:{range=(0,524288), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectSecurityInternalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectSecurityInternalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pSecurityDescriptor {out} (1:{pointer=ref}*(1)[dim:0,size_is=nLength](uchar))
	{
		dimSize1 := uint64(o.Length)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.SecurityDescriptor {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.SecurityDescriptor[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.SecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpnLengthNeeded {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LengthNeeded); err != nil {
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

func (o *xxx_GetObjectSecurityInternalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pSecurityDescriptor {out} (1:{pointer=ref}*(1)[dim:0,size_is=nLength](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.SecurityDescriptor", sizeInfo[0])
		}
		o.SecurityDescriptor = make([]byte, sizeInfo[0])
		for i1 := range o.SecurityDescriptor {
			i1 := i1
			if err := w.ReadData(&o.SecurityDescriptor[i1]); err != nil {
				return err
			}
		}
	}
	// lpnLengthNeeded {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LengthNeeded); err != nil {
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

// GetObjectSecurityInternalRequest structure represents the R_QMGetObjectSecurityInternal operation request
type GetObjectSecurityInternalRequest struct {
	// pObjectFormat:  MUST point to an OBJECT_FORMAT structure which identifies an existing
	// local private queue on the supporting server for which the security configuration
	// is to be retrieved. It MUST NOT be NULL. The ObjType member of the structure MUST
	// be 0x00000001, and the pQueueFormat member MUST NOT be NULL.
	ObjectFormat *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
	// RequestedInformation:  MUST contain a value from the SECURITY_INFORMATION enumeration
	// which indicates the portions of the SECURITY_DESCRIPTOR ([MS-DTYP] section 2.4.6)
	// to be retrieved from the queue identified by pObjectFormat. The SECURITY_INFORMATION
	// enumeration is defined in [MS-MQMQ] section 2.2.3.
	RequestedInformation uint32 `idl:"name:RequestedInformation" json:"requested_information"`
	// nLength:  MUST indicate the byte length of the buffer pointed to by pSecurityDescriptor.
	// This value can be 0x00000000.
	Length uint32 `idl:"name:nLength" json:"length"`
}

func (o *GetObjectSecurityInternalRequest) xxx_ToOp(ctx context.Context, op *xxx_GetObjectSecurityInternalOperation) *xxx_GetObjectSecurityInternalOperation {
	if op == nil {
		op = &xxx_GetObjectSecurityInternalOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectFormat = o.ObjectFormat
	op.RequestedInformation = o.RequestedInformation
	op.Length = o.Length
	return op
}

func (o *GetObjectSecurityInternalRequest) xxx_FromOp(ctx context.Context, op *xxx_GetObjectSecurityInternalOperation) {
	if o == nil {
		return
	}
	o.ObjectFormat = op.ObjectFormat
	o.RequestedInformation = op.RequestedInformation
	o.Length = op.Length
}
func (o *GetObjectSecurityInternalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetObjectSecurityInternalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectSecurityInternalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetObjectSecurityInternalResponse structure represents the R_QMGetObjectSecurityInternal operation response
type GetObjectSecurityInternalResponse struct {
	// pSecurityDescriptor:  MUST be a pointer to an array of bytes into which the server
	// MUST write a self-relative SECURITY_DESCRIPTOR structure. The server MUST NOT write
	// more than nLength bytes to the buffer. If the buffer provided by the client is too
	// small (as indicated by the nLength parameter) to contain the SECURITY_DESCRIPTOR
	// for the queue identified by pObjectFormat, the server MUST return MQ_ERROR_SECURITY_DESCRIPTOR_TOO_SMALL
	// (0xc00e0023). This parameter can be NULL if nLength is 0x00000000.
	//
	// The SECURITY_DESCRIPTOR structure is defined in [MS-DTYP] section 2.4.6.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	// lpnLengthNeeded: MUST NOT be NULL. The server MUST set this DWORD to the byte length
	// of the SECURITY_DESCRIPTOR structure for the queue identified by pObjectFormat.
	LengthNeeded uint32 `idl:"name:lpnLengthNeeded" json:"length_needed"`
	// Return: The R_QMGetObjectSecurityInternal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetObjectSecurityInternalResponse) xxx_ToOp(ctx context.Context, op *xxx_GetObjectSecurityInternalOperation) *xxx_GetObjectSecurityInternalOperation {
	if op == nil {
		op = &xxx_GetObjectSecurityInternalOperation{}
	}
	if o == nil {
		return op
	}
	op.SecurityDescriptor = o.SecurityDescriptor
	op.LengthNeeded = o.LengthNeeded
	op.Return = o.Return
	return op
}

func (o *GetObjectSecurityInternalResponse) xxx_FromOp(ctx context.Context, op *xxx_GetObjectSecurityInternalOperation) {
	if o == nil {
		return
	}
	o.SecurityDescriptor = op.SecurityDescriptor
	o.LengthNeeded = op.LengthNeeded
	o.Return = op.Return
}
func (o *GetObjectSecurityInternalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetObjectSecurityInternalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectSecurityInternalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteObjectOperation structure represents the R_QMDeleteObject operation
type xxx_DeleteObjectOperation struct {
	ObjectFormat *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
	Return       int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteObjectOperation) OpNum() int { return 9 }

func (o *xxx_DeleteObjectOperation) OpName() string { return "/qmcomm/v1/R_QMDeleteObject" }

func (o *xxx_DeleteObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pObjectFormat {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat != nil {
			if err := o.ObjectFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pObjectFormat {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat == nil {
			o.ObjectFormat = &ObjectFormat{}
		}
		if err := o.ObjectFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteObjectRequest structure represents the R_QMDeleteObject operation request
type DeleteObjectRequest struct {
	// pObjectFormat:  MUST point to an OBJECT_FORMAT structure that identifies an existing
	// local private queue on the supporting server. MUST NOT be NULL. The ObjType member
	// of the structure MUST be 0x00000001. The pQueueFormat member MUST NOT be NULL.
	ObjectFormat *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
}

func (o *DeleteObjectRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteObjectOperation) *xxx_DeleteObjectOperation {
	if op == nil {
		op = &xxx_DeleteObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectFormat = o.ObjectFormat
	return op
}

func (o *DeleteObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteObjectOperation) {
	if o == nil {
		return
	}
	o.ObjectFormat = op.ObjectFormat
}
func (o *DeleteObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteObjectResponse structure represents the R_QMDeleteObject operation response
type DeleteObjectResponse struct {
	// Return: The R_QMDeleteObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteObjectResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteObjectOperation) *xxx_DeleteObjectOperation {
	if op == nil {
		op = &xxx_DeleteObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *DeleteObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteObjectOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetObjectPropertiesOperation structure represents the R_QMGetObjectProperties operation
type xxx_GetObjectPropertiesOperation struct {
	ObjectFormat    *ObjectFormat           `idl:"name:pObjectFormat" json:"object_format"`
	CreatePartition uint32                  `idl:"name:cp" json:"create_partition"`
	Property        []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var             []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	Return          int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetObjectPropertiesOperation) OpNum() int { return 10 }

func (o *xxx_GetObjectPropertiesOperation) OpName() string {
	return "/qmcomm/v1/R_QMGetObjectProperties"
}

func (o *xxx_GetObjectPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.CreatePartition < uint32(1) || o.CreatePartition > uint32(128) {
		return fmt.Errorf("CreatePartition is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pObjectFormat {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat != nil {
			if err := o.ObjectFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp])(2:{alias=DWORD}(uint32))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Property {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Property[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Property); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		}
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Var {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Var[i1] != nil {
				if err := o.Var[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Var); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pObjectFormat {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat == nil {
			o.ObjectFormat = &ObjectFormat{}
		}
		if err := o.ObjectFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp])(2:{alias=DWORD}(uint32))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Property", sizeInfo[0])
		}
		o.Property = make([]uint32, sizeInfo[0])
		for i1 := range o.Property {
			i1 := i1
			if err := w.ReadData(&o.Property[i1]); err != nil {
				return err
			}
		}
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Var", sizeInfo[0])
		}
		o.Var = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Var {
			i1 := i1
			if o.Var[i1] == nil {
				o.Var[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Var[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Var {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Var[i1] != nil {
				if err := o.Var[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Var); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetObjectPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Var", sizeInfo[0])
		}
		o.Var = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Var {
			i1 := i1
			if o.Var[i1] == nil {
				o.Var[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Var[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
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

// GetObjectPropertiesRequest structure represents the R_QMGetObjectProperties operation request
type GetObjectPropertiesRequest struct {
	// pObjectFormat:  MUST point to an OBJECT_FORMAT structure which identifies an existing
	// local private queue on the supporting server. MUST NOT be NULL. The ObjType member
	// of the structure MUST be 0x00000001. The pQueueFormat member MUST NOT be NULL.
	ObjectFormat *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
	// cp:  MUST be set to the size (in elements) of the arrays aProp and apVar. The arrays
	// aProp and apVar MUST have an identical number of elements and MUST contain at least
	// one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  MUST be an array of queue property identifiers of properties to retrieve.
	// Each element MUST specify a value from the queue property identifiers table defined
	// in [MS-MQMQ] section 2.3.1. Each element MUST specify the queue property identifier
	// for the corresponding queue property value at the same element index in apVar. MUST
	// contain at least one element.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar: MUST contain at least one element. On input, each element MUST be initialized
	// to the appropriate VARTYPE for the associated property specified by the same element
	// in aProp, or VT_NULL. Otherwise, the server SHOULD return the failure HRESULT MQ_ERROR_PROPERTY
	// (0xc00e0002).<36> On success, the server MUST populate the elements of this array
	// with property values for the properties identified by the corresponding elements
	// of aProp.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
}

func (o *GetObjectPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetObjectPropertiesOperation) *xxx_GetObjectPropertiesOperation {
	if op == nil {
		op = &xxx_GetObjectPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectFormat = o.ObjectFormat
	op.CreatePartition = o.CreatePartition
	op.Property = o.Property
	op.Var = o.Var
	return op
}

func (o *GetObjectPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetObjectPropertiesOperation) {
	if o == nil {
		return
	}
	o.ObjectFormat = op.ObjectFormat
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
}
func (o *GetObjectPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetObjectPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetObjectPropertiesResponse structure represents the R_QMGetObjectProperties operation response
type GetObjectPropertiesResponse struct {
	// apVar: MUST contain at least one element. On input, each element MUST be initialized
	// to the appropriate VARTYPE for the associated property specified by the same element
	// in aProp, or VT_NULL. Otherwise, the server SHOULD return the failure HRESULT MQ_ERROR_PROPERTY
	// (0xc00e0002).<36> On success, the server MUST populate the elements of this array
	// with property values for the properties identified by the corresponding elements
	// of aProp.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// Return: The R_QMGetObjectProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetObjectPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetObjectPropertiesOperation) *xxx_GetObjectPropertiesOperation {
	if op == nil {
		op = &xxx_GetObjectPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Var = o.Var
	op.Return = o.Return
	return op
}

func (o *GetObjectPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetObjectPropertiesOperation) {
	if o == nil {
		return
	}
	o.Var = op.Var
	o.Return = op.Return
}
func (o *GetObjectPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetObjectPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetObjectPropertiesOperation structure represents the R_QMSetObjectProperties operation
type xxx_SetObjectPropertiesOperation struct {
	ObjectFormat    *ObjectFormat           `idl:"name:pObjectFormat" json:"object_format"`
	CreatePartition uint32                  `idl:"name:cp" json:"create_partition"`
	Property        []uint32                `idl:"name:aProp;size_is:(cp);pointer:unique" json:"property"`
	Var             []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp);pointer:unique" json:"var"`
	Return          int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetObjectPropertiesOperation) OpNum() int { return 11 }

func (o *xxx_SetObjectPropertiesOperation) OpName() string {
	return "/qmcomm/v1/R_QMSetObjectProperties"
}

func (o *xxx_SetObjectPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.CreatePartition < uint32(1) || o.CreatePartition > uint32(128) {
		return fmt.Errorf("CreatePartition is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pObjectFormat {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat != nil {
			if err := o.ObjectFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:{pointer=unique}[dim:0,size_is=cp])(2:{alias=DWORD}(uint32))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Property {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Property[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Property); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		}
	}
	// apVar {in} (1:{pointer=unique}[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Var {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Var[i1] != nil {
				if err := o.Var[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Var); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pObjectFormat {in} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat == nil {
			o.ObjectFormat = &ObjectFormat{}
		}
		if err := o.ObjectFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:{pointer=unique}[dim:0,size_is=cp])(2:{alias=DWORD}(uint32))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Property", sizeInfo[0])
		}
		o.Property = make([]uint32, sizeInfo[0])
		for i1 := range o.Property {
			i1 := i1
			if err := w.ReadData(&o.Property[i1]); err != nil {
				return err
			}
		}
	}
	// apVar {in} (1:{pointer=unique}[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Var", sizeInfo[0])
		}
		o.Var = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Var {
			i1 := i1
			if o.Var[i1] == nil {
				o.Var[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Var[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetObjectPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetObjectPropertiesRequest structure represents the R_QMSetObjectProperties operation request
type SetObjectPropertiesRequest struct {
	// pObjectFormat:  MUST point to an OBJECT_FORMAT structure which identifies an existing
	// local private queue on the supporting server. MUST NOT be NULL. The ObjType member
	// of the structure MUST be 0x00000001. The pQueueFormat member MUST NOT be NULL.
	ObjectFormat *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
	// cp:  MUST be set to the size (in elements) of the arrays aProp and apVar. The arrays
	// aProp and apVar MUST have an identical number of elements, and MUST contain at least
	// one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  MUST be an array of queue property identifiers for properties to be updated.
	// Each element MUST specify a value from the queue property identifiers table defined
	// in [MS-MQMQ] section 2.3.1. Each element MUST specify the queue property identifier
	// for the corresponding queue property value at the same element index in apVar. MUST
	// contain at least one element.
	Property []uint32 `idl:"name:aProp;size_is:(cp);pointer:unique" json:"property"`
	// apVar:  MUST be an array that specifies the property values to update. Each element
	// MUST specify the property value for the corresponding property identifier at the
	// same element index in aProp. MUST contain at least one element. The vt (VARTYPE)
	// member of each PROPVARIANT element MUST be set to the appropriate type for the property
	// being updated; otherwise, the server SHOULD return the failure HRESULT MQ_ERROR_PROPERTY
	// (0xc00e0002).<39> Queue properties and their appropriate VARTYPEs are specified by
	// [MS-MQMQ] section 2.3.1.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp);pointer:unique" json:"var"`
}

func (o *SetObjectPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetObjectPropertiesOperation) *xxx_SetObjectPropertiesOperation {
	if op == nil {
		op = &xxx_SetObjectPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectFormat = o.ObjectFormat
	op.CreatePartition = o.CreatePartition
	op.Property = o.Property
	op.Var = o.Var
	return op
}

func (o *SetObjectPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetObjectPropertiesOperation) {
	if o == nil {
		return
	}
	o.ObjectFormat = op.ObjectFormat
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
}
func (o *SetObjectPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetObjectPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetObjectPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetObjectPropertiesResponse structure represents the R_QMSetObjectProperties operation response
type SetObjectPropertiesResponse struct {
	// Return: The R_QMSetObjectProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetObjectPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetObjectPropertiesOperation) *xxx_SetObjectPropertiesOperation {
	if op == nil {
		op = &xxx_SetObjectPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetObjectPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetObjectPropertiesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetObjectPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetObjectPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetObjectPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ObjectPathToObjectFormatOperation structure represents the R_QMObjectPathToObjectFormat operation
type xxx_ObjectPathToObjectFormatOperation struct {
	PathName     string        `idl:"name:lpwcsPathName;string" json:"path_name"`
	ObjectFormat *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
	Return       int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_ObjectPathToObjectFormatOperation) OpNum() int { return 12 }

func (o *xxx_ObjectPathToObjectFormatOperation) OpName() string {
	return "/qmcomm/v1/R_QMObjectPathToObjectFormat"
}

func (o *xxx_ObjectPathToObjectFormatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ObjectPathToObjectFormatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpwcsPathName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	// pObjectFormat {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat != nil {
			if err := o.ObjectFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ObjectPathToObjectFormatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpwcsPathName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
	}
	// pObjectFormat {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat == nil {
			o.ObjectFormat = &ObjectFormat{}
		}
		if err := o.ObjectFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ObjectPathToObjectFormatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ObjectPathToObjectFormatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pObjectFormat {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat != nil {
			if err := o.ObjectFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectFormat{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ObjectPathToObjectFormatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pObjectFormat {in, out} (1:{pointer=ref}*(1)(struct))
	{
		if o.ObjectFormat == nil {
			o.ObjectFormat = &ObjectFormat{}
		}
		if err := o.ObjectFormat.UnmarshalNDR(ctx, w); err != nil {
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

// ObjectPathToObjectFormatRequest structure represents the R_QMObjectPathToObjectFormat operation request
type ObjectPathToObjectFormatRequest struct {
	// lpwcsPathName:  MUST be a pointer to a null-terminated path name string, as defined
	// by [MS-MQMQ] section 2.1.1. The path name MUST identify an existing private queue
	// located on a supporting server.
	PathName string `idl:"name:lpwcsPathName;string" json:"path_name"`
	// pObjectFormat:  MUST be a pointer to an OBJECT_FORMAT structure, as specified in
	// section 2.2.3.5. On success, this structure MUST be populated with a direct format
	// name or private format name for the queue identified by lpwcsPathName. This specification
	// does not mandate the process through which a server produces a format name for a
	// given path name.
	ObjectFormat *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
}

func (o *ObjectPathToObjectFormatRequest) xxx_ToOp(ctx context.Context, op *xxx_ObjectPathToObjectFormatOperation) *xxx_ObjectPathToObjectFormatOperation {
	if op == nil {
		op = &xxx_ObjectPathToObjectFormatOperation{}
	}
	if o == nil {
		return op
	}
	op.PathName = o.PathName
	op.ObjectFormat = o.ObjectFormat
	return op
}

func (o *ObjectPathToObjectFormatRequest) xxx_FromOp(ctx context.Context, op *xxx_ObjectPathToObjectFormatOperation) {
	if o == nil {
		return
	}
	o.PathName = op.PathName
	o.ObjectFormat = op.ObjectFormat
}
func (o *ObjectPathToObjectFormatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ObjectPathToObjectFormatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ObjectPathToObjectFormatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ObjectPathToObjectFormatResponse structure represents the R_QMObjectPathToObjectFormat operation response
type ObjectPathToObjectFormatResponse struct {
	// pObjectFormat:  MUST be a pointer to an OBJECT_FORMAT structure, as specified in
	// section 2.2.3.5. On success, this structure MUST be populated with a direct format
	// name or private format name for the queue identified by lpwcsPathName. This specification
	// does not mandate the process through which a server produces a format name for a
	// given path name.
	ObjectFormat *ObjectFormat `idl:"name:pObjectFormat" json:"object_format"`
	// Return: The R_QMObjectPathToObjectFormat return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ObjectPathToObjectFormatResponse) xxx_ToOp(ctx context.Context, op *xxx_ObjectPathToObjectFormatOperation) *xxx_ObjectPathToObjectFormatOperation {
	if op == nil {
		op = &xxx_ObjectPathToObjectFormatOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectFormat = o.ObjectFormat
	op.Return = o.Return
	return op
}

func (o *ObjectPathToObjectFormatResponse) xxx_FromOp(ctx context.Context, op *xxx_ObjectPathToObjectFormatOperation) {
	if o == nil {
		return
	}
	o.ObjectFormat = op.ObjectFormat
	o.Return = op.Return
}
func (o *ObjectPathToObjectFormatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ObjectPathToObjectFormatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ObjectPathToObjectFormatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTMWhereaboutsOperation structure represents the R_QMGetTmWhereabouts operation
type xxx_GetTMWhereaboutsOperation struct {
	BufferLength      uint32 `idl:"name:cbBufSize" json:"buffer_length"`
	Whereabouts       []byte `idl:"name:pbWhereabouts;size_is:(cbBufSize)" json:"whereabouts"`
	WhereaboutsLength uint32 `idl:"name:pcbWhereabouts" json:"whereabouts_length"`
	Return            int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTMWhereaboutsOperation) OpNum() int { return 14 }

func (o *xxx_GetTMWhereaboutsOperation) OpName() string { return "/qmcomm/v1/R_QMGetTmWhereabouts" }

func (o *xxx_GetTMWhereaboutsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.BufferLength > uint32(131072) {
		return fmt.Errorf("BufferLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTMWhereaboutsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// cbBufSize {in} (1:{range=(0,131072), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTMWhereaboutsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// cbBufSize {in} (1:{range=(0,131072), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTMWhereaboutsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTMWhereaboutsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pbWhereabouts {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbBufSize](uchar))
	{
		dimSize1 := uint64(o.BufferLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Whereabouts {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Whereabouts[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Whereabouts); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pcbWhereabouts {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.WhereaboutsLength); err != nil {
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

func (o *xxx_GetTMWhereaboutsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pbWhereabouts {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbBufSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Whereabouts", sizeInfo[0])
		}
		o.Whereabouts = make([]byte, sizeInfo[0])
		for i1 := range o.Whereabouts {
			i1 := i1
			if err := w.ReadData(&o.Whereabouts[i1]); err != nil {
				return err
			}
		}
	}
	// pcbWhereabouts {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.WhereaboutsLength); err != nil {
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

// GetTMWhereaboutsRequest structure represents the R_QMGetTmWhereabouts operation request
type GetTMWhereaboutsRequest struct {
	// cbBufSize:  MUST be set to the byte length of the buffer pointed to by pbWhereabouts.
	// If this value is 0x00000000, the server MUST ignore the pbWhereabouts parameter.
	BufferLength uint32 `idl:"name:cbBufSize" json:"buffer_length"`
}

func (o *GetTMWhereaboutsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTMWhereaboutsOperation) *xxx_GetTMWhereaboutsOperation {
	if op == nil {
		op = &xxx_GetTMWhereaboutsOperation{}
	}
	if o == nil {
		return op
	}
	op.BufferLength = o.BufferLength
	return op
}

func (o *GetTMWhereaboutsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTMWhereaboutsOperation) {
	if o == nil {
		return
	}
	o.BufferLength = op.BufferLength
}
func (o *GetTMWhereaboutsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTMWhereaboutsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTMWhereaboutsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTMWhereaboutsResponse structure represents the R_QMGetTmWhereabouts operation response
type GetTMWhereaboutsResponse struct {
	// pbWhereabouts:  On success, points to an array of bytes containing a Distributed
	// Transaction Coordinator (DTC) SWhereabouts structure, as specified in [MS-DTCO] section
	// 2.2.5.11.
	Whereabouts []byte `idl:"name:pbWhereabouts;size_is:(cbBufSize)" json:"whereabouts"`
	// pcbWhereabouts:  On success, or ifMQ_ERROR_USER_BUFFER_TOO_SMALL (0xc00e0028) is
	// returned, pcbWhereabouts points to a DWORD containing the byte length of the SWhereabouts
	// structure retrieved from the DTC; otherwise, this parameter MUST be ignored.
	WhereaboutsLength uint32 `idl:"name:pcbWhereabouts" json:"whereabouts_length"`
	// Return: The R_QMGetTmWhereabouts return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTMWhereaboutsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTMWhereaboutsOperation) *xxx_GetTMWhereaboutsOperation {
	if op == nil {
		op = &xxx_GetTMWhereaboutsOperation{}
	}
	if o == nil {
		return op
	}
	op.Whereabouts = o.Whereabouts
	op.WhereaboutsLength = o.WhereaboutsLength
	op.Return = o.Return
	return op
}

func (o *GetTMWhereaboutsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTMWhereaboutsOperation) {
	if o == nil {
		return
	}
	o.Whereabouts = op.Whereabouts
	o.WhereaboutsLength = op.WhereaboutsLength
	o.Return = op.Return
}
func (o *GetTMWhereaboutsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTMWhereaboutsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTMWhereaboutsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnlistTransactionOperation structure represents the R_QMEnlistTransaction operation
type xxx_EnlistTransactionOperation struct {
	UOW          *mqmq.TransactionUOW `idl:"name:pUow" json:"uow"`
	CookieLength uint32               `idl:"name:cbCookie" json:"cookie_length"`
	Cookie       []byte               `idl:"name:pbCookie;size_is:(cbCookie)" json:"cookie"`
	Return       int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_EnlistTransactionOperation) OpNum() int { return 15 }

func (o *xxx_EnlistTransactionOperation) OpName() string { return "/qmcomm/v1/R_QMEnlistTransaction" }

func (o *xxx_EnlistTransactionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Cookie != nil && o.CookieLength == 0 {
		o.CookieLength = uint32(len(o.Cookie))
	}
	if o.CookieLength > uint32(131072) {
		return fmt.Errorf("CookieLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnlistTransactionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pUow {in} (1:{pointer=ref}*(1))(2:{alias=XACTUOW}(struct))
	{
		if o.UOW != nil {
			if err := o.UOW.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqmq.TransactionUOW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// cbCookie {in} (1:{range=(0,131072), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CookieLength); err != nil {
			return err
		}
	}
	// pbCookie {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbCookie](uchar))
	{
		dimSize1 := uint64(o.CookieLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Cookie {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Cookie[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Cookie); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EnlistTransactionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pUow {in} (1:{pointer=ref}*(1))(2:{alias=XACTUOW}(struct))
	{
		if o.UOW == nil {
			o.UOW = &mqmq.TransactionUOW{}
		}
		if err := o.UOW.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// cbCookie {in} (1:{range=(0,131072), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CookieLength); err != nil {
			return err
		}
	}
	// pbCookie {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbCookie](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Cookie", sizeInfo[0])
		}
		o.Cookie = make([]byte, sizeInfo[0])
		for i1 := range o.Cookie {
			i1 := i1
			if err := w.ReadData(&o.Cookie[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EnlistTransactionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnlistTransactionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnlistTransactionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnlistTransactionRequest structure represents the R_QMEnlistTransaction operation request
type EnlistTransactionRequest struct {
	// pUow:  MUST point to an XACTUOW structure ([MS-MQMQ] section 2.2.18.1.8) that identifies
	// the external transaction in which the server is to enlist, as specified in section
	// 2.2.3.1.
	UOW *mqmq.TransactionUOW `idl:"name:pUow" json:"uow"`
	// cbCookie:  MUST be set to the byte length of the buffer pointed to by pbCookie.
	CookieLength uint32 `idl:"name:cbCookie" json:"cookie_length"`
	// pbCookie: MUST be a pointer to an array of bytes containing an exported transaction
	// cookie, which can be obtained as specified in [MS-DTCO] section 3.3.4.14.
	Cookie []byte `idl:"name:pbCookie;size_is:(cbCookie)" json:"cookie"`
}

func (o *EnlistTransactionRequest) xxx_ToOp(ctx context.Context, op *xxx_EnlistTransactionOperation) *xxx_EnlistTransactionOperation {
	if op == nil {
		op = &xxx_EnlistTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.UOW = o.UOW
	op.CookieLength = o.CookieLength
	op.Cookie = o.Cookie
	return op
}

func (o *EnlistTransactionRequest) xxx_FromOp(ctx context.Context, op *xxx_EnlistTransactionOperation) {
	if o == nil {
		return
	}
	o.UOW = op.UOW
	o.CookieLength = op.CookieLength
	o.Cookie = op.Cookie
}
func (o *EnlistTransactionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnlistTransactionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnlistTransactionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnlistTransactionResponse structure represents the R_QMEnlistTransaction operation response
type EnlistTransactionResponse struct {
	// Return: The R_QMEnlistTransaction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnlistTransactionResponse) xxx_ToOp(ctx context.Context, op *xxx_EnlistTransactionOperation) *xxx_EnlistTransactionOperation {
	if op == nil {
		op = &xxx_EnlistTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *EnlistTransactionResponse) xxx_FromOp(ctx context.Context, op *xxx_EnlistTransactionOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EnlistTransactionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnlistTransactionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnlistTransactionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnlistInternalTransactionOperation structure represents the R_QMEnlistInternalTransaction operation
type xxx_EnlistInternalTransactionOperation struct {
	UOW                 *mqmq.TransactionUOW `idl:"name:pUow" json:"uow"`
	InternalTransaction *InternalTransaction `idl:"name:phIntXact" json:"internal_transaction"`
	Return              int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_EnlistInternalTransactionOperation) OpNum() int { return 16 }

func (o *xxx_EnlistInternalTransactionOperation) OpName() string {
	return "/qmcomm/v1/R_QMEnlistInternalTransaction"
}

func (o *xxx_EnlistInternalTransactionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnlistInternalTransactionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pUow {in} (1:{pointer=ref}*(1))(2:{alias=XACTUOW}(struct))
	{
		if o.UOW != nil {
			if err := o.UOW.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqmq.TransactionUOW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EnlistInternalTransactionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pUow {in} (1:{pointer=ref}*(1))(2:{alias=XACTUOW}(struct))
	{
		if o.UOW == nil {
			o.UOW = &mqmq.TransactionUOW{}
		}
		if err := o.UOW.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnlistInternalTransactionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnlistInternalTransactionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phIntXact {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_INT_XACT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.InternalTransaction != nil {
			if err := o.InternalTransaction.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InternalTransaction{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnlistInternalTransactionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phIntXact {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_INT_XACT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.InternalTransaction == nil {
			o.InternalTransaction = &InternalTransaction{}
		}
		if err := o.InternalTransaction.UnmarshalNDR(ctx, w); err != nil {
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

// EnlistInternalTransactionRequest structure represents the R_QMEnlistInternalTransaction operation request
type EnlistInternalTransactionRequest struct {
	// pUow:  MUST point to an XACTUOW structure that uniquely identifies the internal
	// transaction in which the server is to enlist.<44>
	UOW *mqmq.TransactionUOW `idl:"name:pUow" json:"uow"`
}

func (o *EnlistInternalTransactionRequest) xxx_ToOp(ctx context.Context, op *xxx_EnlistInternalTransactionOperation) *xxx_EnlistInternalTransactionOperation {
	if op == nil {
		op = &xxx_EnlistInternalTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.UOW = o.UOW
	return op
}

func (o *EnlistInternalTransactionRequest) xxx_FromOp(ctx context.Context, op *xxx_EnlistInternalTransactionOperation) {
	if o == nil {
		return
	}
	o.UOW = op.UOW
}
func (o *EnlistInternalTransactionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnlistInternalTransactionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnlistInternalTransactionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnlistInternalTransactionResponse structure represents the R_QMEnlistInternalTransaction operation response
type EnlistInternalTransactionResponse struct {
	// phIntXact:  A pointer to receive the new RPC_INT_XACT_HANDLE which represents the
	// new internal transaction context.
	InternalTransaction *InternalTransaction `idl:"name:phIntXact" json:"internal_transaction"`
	// Return: The R_QMEnlistInternalTransaction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnlistInternalTransactionResponse) xxx_ToOp(ctx context.Context, op *xxx_EnlistInternalTransactionOperation) *xxx_EnlistInternalTransactionOperation {
	if op == nil {
		op = &xxx_EnlistInternalTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.InternalTransaction = o.InternalTransaction
	op.Return = o.Return
	return op
}

func (o *EnlistInternalTransactionResponse) xxx_FromOp(ctx context.Context, op *xxx_EnlistInternalTransactionOperation) {
	if o == nil {
		return
	}
	o.InternalTransaction = op.InternalTransaction
	o.Return = op.Return
}
func (o *EnlistInternalTransactionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnlistInternalTransactionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnlistInternalTransactionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CommitTransactionOperation structure represents the R_QMCommitTransaction operation
type xxx_CommitTransactionOperation struct {
	InternalTransaction *InternalTransaction `idl:"name:phIntXact" json:"internal_transaction"`
	Return              int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_CommitTransactionOperation) OpNum() int { return 17 }

func (o *xxx_CommitTransactionOperation) OpName() string { return "/qmcomm/v1/R_QMCommitTransaction" }

func (o *xxx_CommitTransactionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitTransactionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phIntXact {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_INT_XACT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.InternalTransaction != nil {
			if err := o.InternalTransaction.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InternalTransaction{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CommitTransactionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phIntXact {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_INT_XACT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.InternalTransaction == nil {
			o.InternalTransaction = &InternalTransaction{}
		}
		if err := o.InternalTransaction.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitTransactionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitTransactionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phIntXact {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_INT_XACT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.InternalTransaction != nil {
			if err := o.InternalTransaction.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InternalTransaction{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CommitTransactionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phIntXact {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_INT_XACT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.InternalTransaction == nil {
			o.InternalTransaction = &InternalTransaction{}
		}
		if err := o.InternalTransaction.UnmarshalNDR(ctx, w); err != nil {
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

// CommitTransactionRequest structure represents the R_QMCommitTransaction operation request
type CommitTransactionRequest struct {
	// phIntXact:  MUST be an RPC_INT_XACT_HANDLE (section 2.2.1.1.1) identifying the internal
	// transaction to commit. MUST NOT be NULL. The value of this handle MUST have been
	// acquired from R_QMEnlistInternalTransaction (section 3.1.4.14). On return, the server
	// MUST set this parameter to NULL.
	InternalTransaction *InternalTransaction `idl:"name:phIntXact" json:"internal_transaction"`
}

func (o *CommitTransactionRequest) xxx_ToOp(ctx context.Context, op *xxx_CommitTransactionOperation) *xxx_CommitTransactionOperation {
	if op == nil {
		op = &xxx_CommitTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.InternalTransaction = o.InternalTransaction
	return op
}

func (o *CommitTransactionRequest) xxx_FromOp(ctx context.Context, op *xxx_CommitTransactionOperation) {
	if o == nil {
		return
	}
	o.InternalTransaction = op.InternalTransaction
}
func (o *CommitTransactionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CommitTransactionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitTransactionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CommitTransactionResponse structure represents the R_QMCommitTransaction operation response
type CommitTransactionResponse struct {
	// phIntXact:  MUST be an RPC_INT_XACT_HANDLE (section 2.2.1.1.1) identifying the internal
	// transaction to commit. MUST NOT be NULL. The value of this handle MUST have been
	// acquired from R_QMEnlistInternalTransaction (section 3.1.4.14). On return, the server
	// MUST set this parameter to NULL.
	InternalTransaction *InternalTransaction `idl:"name:phIntXact" json:"internal_transaction"`
	// Return: The R_QMCommitTransaction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CommitTransactionResponse) xxx_ToOp(ctx context.Context, op *xxx_CommitTransactionOperation) *xxx_CommitTransactionOperation {
	if op == nil {
		op = &xxx_CommitTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.InternalTransaction = o.InternalTransaction
	op.Return = o.Return
	return op
}

func (o *CommitTransactionResponse) xxx_FromOp(ctx context.Context, op *xxx_CommitTransactionOperation) {
	if o == nil {
		return
	}
	o.InternalTransaction = op.InternalTransaction
	o.Return = op.Return
}
func (o *CommitTransactionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CommitTransactionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitTransactionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AbortTransactionOperation structure represents the R_QMAbortTransaction operation
type xxx_AbortTransactionOperation struct {
	InternalTransaction *InternalTransaction `idl:"name:phIntXact" json:"internal_transaction"`
	Return              int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_AbortTransactionOperation) OpNum() int { return 18 }

func (o *xxx_AbortTransactionOperation) OpName() string { return "/qmcomm/v1/R_QMAbortTransaction" }

func (o *xxx_AbortTransactionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortTransactionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phIntXact {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_INT_XACT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.InternalTransaction != nil {
			if err := o.InternalTransaction.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InternalTransaction{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_AbortTransactionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phIntXact {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_INT_XACT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.InternalTransaction == nil {
			o.InternalTransaction = &InternalTransaction{}
		}
		if err := o.InternalTransaction.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortTransactionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortTransactionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phIntXact {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_INT_XACT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.InternalTransaction != nil {
			if err := o.InternalTransaction.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InternalTransaction{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_AbortTransactionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phIntXact {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_INT_XACT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.InternalTransaction == nil {
			o.InternalTransaction = &InternalTransaction{}
		}
		if err := o.InternalTransaction.UnmarshalNDR(ctx, w); err != nil {
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

// AbortTransactionRequest structure represents the R_QMAbortTransaction operation request
type AbortTransactionRequest struct {
	// phIntXact:  MUST be an RPC_INT_XACT_HANDLE identifying the internal transaction
	// to abort. MUST NOT be NULL. The value of this handle MUST have been acquired from
	// R_QMEnlistInternalTransaction. On return, the server MUST set this parameter to NULL.
	InternalTransaction *InternalTransaction `idl:"name:phIntXact" json:"internal_transaction"`
}

func (o *AbortTransactionRequest) xxx_ToOp(ctx context.Context, op *xxx_AbortTransactionOperation) *xxx_AbortTransactionOperation {
	if op == nil {
		op = &xxx_AbortTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.InternalTransaction = o.InternalTransaction
	return op
}

func (o *AbortTransactionRequest) xxx_FromOp(ctx context.Context, op *xxx_AbortTransactionOperation) {
	if o == nil {
		return
	}
	o.InternalTransaction = op.InternalTransaction
}
func (o *AbortTransactionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AbortTransactionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortTransactionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AbortTransactionResponse structure represents the R_QMAbortTransaction operation response
type AbortTransactionResponse struct {
	// phIntXact:  MUST be an RPC_INT_XACT_HANDLE identifying the internal transaction
	// to abort. MUST NOT be NULL. The value of this handle MUST have been acquired from
	// R_QMEnlistInternalTransaction. On return, the server MUST set this parameter to NULL.
	InternalTransaction *InternalTransaction `idl:"name:phIntXact" json:"internal_transaction"`
	// Return: The R_QMAbortTransaction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AbortTransactionResponse) xxx_ToOp(ctx context.Context, op *xxx_AbortTransactionOperation) *xxx_AbortTransactionOperation {
	if op == nil {
		op = &xxx_AbortTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.InternalTransaction = o.InternalTransaction
	op.Return = o.Return
	return op
}

func (o *AbortTransactionResponse) xxx_FromOp(ctx context.Context, op *xxx_AbortTransactionOperation) {
	if o == nil {
		return
	}
	o.InternalTransaction = op.InternalTransaction
	o.Return = op.Return
}
func (o *AbortTransactionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AbortTransactionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortTransactionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenQueueInternalOperation structure represents the rpc_QMOpenQueueInternal operation
type xxx_OpenQueueInternalOperation struct {
	QueueFormat     *mqmq.QueueFormat `idl:"name:pQueueFormat" json:"queue_format"`
	DesiredAccess   uint32            `idl:"name:dwDesiredAccess" json:"desired_access"`
	ShareMode       uint32            `idl:"name:dwShareMode" json:"share_mode"`
	RemoteQueue     uint32            `idl:"name:hRemoteQueue" json:"remote_queue"`
	RemoteQueueName string            `idl:"name:lplpRemoteQueueName;string;pointer:ptr" json:"remote_queue_name"`
	QueueID         uint32            `idl:"name:dwpQueue" json:"queue_id"`
	ClientGUID      *dtyp.GUID        `idl:"name:pLicGuid" json:"client_guid"`
	ClientName      string            `idl:"name:lpClientName;string" json:"client_name"`
	QMContext       uint32            `idl:"name:pdwQMContext" json:"qm_context"`
	Queue           *Queue            `idl:"name:phQueue" json:"queue"`
	RemoteProtocol  uint32            `idl:"name:dwRemoteProtocol" json:"remote_protocol"`
	RemoteContext   uint32            `idl:"name:dwpRemoteContext" json:"remote_context"`
	Return          int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenQueueInternalOperation) OpNum() int { return 19 }

func (o *xxx_OpenQueueInternalOperation) OpName() string { return "/qmcomm/v1/rpc_QMOpenQueueInternal" }

func (o *xxx_OpenQueueInternalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueInternalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pQueueFormat {in} (1:{pointer=ref}*(1))(2:{alias=QUEUE_FORMAT}(struct))
	{
		if o.QueueFormat != nil {
			if err := o.QueueFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqmq.QueueFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwDesiredAccess {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	// dwShareMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ShareMode); err != nil {
			return err
		}
	}
	// hRemoteQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RemoteQueue); err != nil {
			return err
		}
	}
	// lplpRemoteQueueName {in, out} (1:{string, pointer=ptr}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.RemoteQueueName != "" {
			_ptr_lplpRemoteQueueName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RemoteQueueName != "" {
					_ptr_lplpRemoteQueueName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.RemoteQueueName); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.RemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
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
	// dwpQueue {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QueueID); err != nil {
			return err
		}
	}
	// pLicGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClientGUID != nil {
			if err := o.ClientGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpClientName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
			return err
		}
	}
	// dwRemoteProtocol {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RemoteProtocol); err != nil {
			return err
		}
	}
	// dwpRemoteContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RemoteContext); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueInternalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pQueueFormat {in} (1:{pointer=ref}*(1))(2:{alias=QUEUE_FORMAT}(struct))
	{
		if o.QueueFormat == nil {
			o.QueueFormat = &mqmq.QueueFormat{}
		}
		if err := o.QueueFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwDesiredAccess {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	// dwShareMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ShareMode); err != nil {
			return err
		}
	}
	// hRemoteQueue {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RemoteQueue); err != nil {
			return err
		}
	}
	// lplpRemoteQueueName {in, out} (1:{string, pointer=ptr}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_lplpRemoteQueueName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_lplpRemoteQueueName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteQueueName); err != nil {
					return err
				}
				return nil
			})
			_s_lplpRemoteQueueName := func(ptr interface{}) { o.RemoteQueueName = *ptr.(*string) }
			if err := w.ReadPointer(&o.RemoteQueueName, _s_lplpRemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
				return err
			}
			return nil
		})
		_s_lplpRemoteQueueName := func(ptr interface{}) { o.RemoteQueueName = *ptr.(*string) }
		if err := w.ReadPointer(&o.RemoteQueueName, _s_lplpRemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwpQueue {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QueueID); err != nil {
			return err
		}
	}
	// pLicGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClientGUID == nil {
			o.ClientGUID = &dtyp.GUID{}
		}
		if err := o.ClientGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpClientName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
	}
	// dwRemoteProtocol {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RemoteProtocol); err != nil {
			return err
		}
	}
	// dwpRemoteContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RemoteContext); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueInternalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueueInternalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lplpRemoteQueueName {in, out} (1:{string, pointer=ptr}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.RemoteQueueName != "" {
			_ptr_lplpRemoteQueueName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RemoteQueueName != "" {
					_ptr_lplpRemoteQueueName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.RemoteQueueName); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.RemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
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
	// pdwQMContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QMContext); err != nil {
			return err
		}
	}
	// phQueue {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue != nil {
			if err := o.Queue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Queue{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenQueueInternalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lplpRemoteQueueName {in, out} (1:{string, pointer=ptr}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_lplpRemoteQueueName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_lplpRemoteQueueName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteQueueName); err != nil {
					return err
				}
				return nil
			})
			_s_lplpRemoteQueueName := func(ptr interface{}) { o.RemoteQueueName = *ptr.(*string) }
			if err := w.ReadPointer(&o.RemoteQueueName, _s_lplpRemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
				return err
			}
			return nil
		})
		_s_lplpRemoteQueueName := func(ptr interface{}) { o.RemoteQueueName = *ptr.(*string) }
		if err := w.ReadPointer(&o.RemoteQueueName, _s_lplpRemoteQueueName, _ptr_lplpRemoteQueueName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwQMContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QMContext); err != nil {
			return err
		}
	}
	// phQueue {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue == nil {
			o.Queue = &Queue{}
		}
		if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
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

// OpenQueueInternalRequest structure represents the rpc_QMOpenQueueInternal operation request
type OpenQueueInternalRequest struct {
	// pQueueFormat: MUST be a pointer to a QUEUE_FORMAT ([MS-MQMQ] section 2.2.7) structure,
	// which identifies an existing queue to be opened. MUST NOT be NULL and MUST conform
	// to the format name syntax rules defined in [MS-MQMQ].
	QueueFormat *mqmq.QueueFormat `idl:"name:pQueueFormat" json:"queue_format"`
	// dwDesiredAccess: A DWORD that specifies the access mode requested for the queue.
	// The access mode defines the set of operations which can be invoked using the returned
	// queue handle. The value MUST be one of the following:
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                              |                                                                                  |
	//	|                    VALUE                     |                                     MEANING                                      |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_RECEIVE_ACCESS 0x00000001                 | The server MUST permit only the following operations using the returned queue    |
	//	|                                              | handle: Message peek Message receive (peek and delete) Queue purge               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_SEND_ACCESS 0x00000002                    | The server MUST permit only message send operations using the returned queue     |
	//	|                                              | handle.                                                                          |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_PEEK_ACCESS 0x00000020                    | The server MUST permit only message peek operations using the returned queue     |
	//	|                                              | handle.                                                                          |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_RECEIVE_ACCESS|MQ_ADMIN_ACCESS 0x00000081 | The returned queue handle MUST perform operations on the outgoing queue          |
	//	|                                              | associated with the queue identified by pQueueFormat. Additionally, the server   |
	//	|                                              | MUST permit only the following operations using the returned queue handle:       |
	//	|                                              | Message peek Message receive (peek and delete) Queue purge                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_PEEK_ACCESS|MQ_ADMIN_ACCESS 0x000000a0    | The returned queue handle MUST perform operations on the outgoing queue          |
	//	|                                              | associated with the queue identified by pQueueFormat. Additionally, the server   |
	//	|                                              | MUST permit only message peek operations using the returned queue handle.        |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// If pQueueFormat contains an HTTP or multicast format name, R_QMOpenRemoteQueue (section
	// 3.1.4.2) MUST be MQ_SEND_ACCESS (0x00000002).
	//
	// If pQueueFormat identifies a sub-queue, dwDesiredAccess MUST NOT be MQ_SEND_ACCESS
	// (0x00000002).
	//
	// If pQueueFormat identifies a system, journal, machine, or connector queue, dwDesiredAccess
	// MUST be MQ_RECEIVE_ACCESS (0x00000001) or MQ_PEEK_ACCESS (0x00000020).
	DesiredAccess uint32 `idl:"name:dwDesiredAccess" json:"desired_access"`
	// dwShareMode:  Specifies the exclusivity level for the opened queue. The value MUST
	// be one of the following:
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                     MEANING                                      |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_DENY_NONE 0x00000000          | The queue is not opened exclusively.                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_DENY_RECEIVE_SHARE 0x00000001 | The queue is opened for exclusive read access. If the queue has already been     |
	//	|                                  | opened for read access, the server MUST return a failure HRESULT. If the queue   |
	//	|                                  | is opened successfully for exclusive read access, subsequent attempts to open    |
	//	|                                  | the same queue for read access MUST return a failure HRESULT until the queue has |
	//	|                                  | been closed.                                                                     |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	ShareMode uint32 `idl:"name:dwShareMode" json:"share_mode"`
	// hRemoteQueue: MUST be 0x00000000, or MUST contain a DWORD value obtained from the
	// phQueue out-parameter of the R_QMOpenRemoteQueue method invoked at a remote queue
	// manager.
	RemoteQueue uint32 `idl:"name:hRemoteQueue" json:"remote_queue"`
	// lplpRemoteQueueName:  On input, the server MUST ignore lplpRemoteQueueName. If hRemoteQueue
	// is 0x00000000 and the queue identified by pQueueFormat is located at a remote queue
	// manager, the server MUST set this string to a null-terminated path name, from which
	// the client can determine the computer name of the remote queue manager, as specified
	// in [MS-MQMQ] section 2.1.1.
	RemoteQueueName string `idl:"name:lplpRemoteQueueName;string;pointer:ptr" json:"remote_queue_name"`
	// dwpQueue: If hRemoteQueue is 0x00000000, dwpQueue MUST be NULL; otherwise, dwpQueue
	// MUST contain a DWORD value obtained from the dwpQueue out-parameter of the R_QMOpenRemoteQueue
	// method invoked at a remote queue manager.
	QueueID uint32 `idl:"name:dwpQueue" json:"queue_id"`
	// pLicGuid:  MUST be a pointer to a valid GUID which uniquely identifies the client.<45><46>
	// The server MAY ignore this parameter.<47>
	ClientGUID *dtyp.GUID `idl:"name:pLicGuid" json:"client_guid"`
	// lpClientName:  MUST be a null-terminated string containing the client's computer
	// name.<48> Servers MAY use this parameter in concert with the pLicGuid parameter to
	// implement limits on the number of unique clients which can open queue handles.<49>
	// Implementing connection limits is optional and not recommended.
	ClientName string `idl:"name:lpClientName;string" json:"client_name"`
	// dwRemoteProtocol: Clients MUST set this parameter to 0x00000000. Servers SHOULD ignore
	// this parameter.<50>
	//
	//	+------------+----------------------------------------------+
	//	|            |                                              |
	//	|   VALUE    |                   MEANING                    |
	//	|            |                                              |
	//	+------------+----------------------------------------------+
	//	+------------+----------------------------------------------+
	//	| 0x00000000 | The TCP/IP protocol sequence is to be used.  |
	//	+------------+----------------------------------------------+
	//	| 0x00000003 | The IPX/SPX protocol sequence is to be used. |
	//	+------------+----------------------------------------------+
	RemoteProtocol uint32 `idl:"name:dwRemoteProtocol" json:"remote_protocol"`
	// dwpRemoteContext: If hRemoteQueue is 0x00000000, dwpRemoteContext MUST contain 0x000000000;
	// otherwise, dwpRemoteContext MUST contain a DWORD value obtained from the pdwContext
	// out-parameter of the R_QMOpenRemoteQueue (section 3.1.4.2) method invoked at a remote
	// queue manager.
	RemoteContext uint32 `idl:"name:dwpRemoteContext" json:"remote_context"`
}

func (o *OpenQueueInternalRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenQueueInternalOperation) *xxx_OpenQueueInternalOperation {
	if op == nil {
		op = &xxx_OpenQueueInternalOperation{}
	}
	if o == nil {
		return op
	}
	op.QueueFormat = o.QueueFormat
	op.DesiredAccess = o.DesiredAccess
	op.ShareMode = o.ShareMode
	op.RemoteQueue = o.RemoteQueue
	op.RemoteQueueName = o.RemoteQueueName
	op.QueueID = o.QueueID
	op.ClientGUID = o.ClientGUID
	op.ClientName = o.ClientName
	op.RemoteProtocol = o.RemoteProtocol
	op.RemoteContext = o.RemoteContext
	return op
}

func (o *OpenQueueInternalRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenQueueInternalOperation) {
	if o == nil {
		return
	}
	o.QueueFormat = op.QueueFormat
	o.DesiredAccess = op.DesiredAccess
	o.ShareMode = op.ShareMode
	o.RemoteQueue = op.RemoteQueue
	o.RemoteQueueName = op.RemoteQueueName
	o.QueueID = op.QueueID
	o.ClientGUID = op.ClientGUID
	o.ClientName = op.ClientName
	o.RemoteProtocol = op.RemoteProtocol
	o.RemoteContext = op.RemoteContext
}
func (o *OpenQueueInternalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenQueueInternalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenQueueInternalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenQueueInternalResponse structure represents the rpc_QMOpenQueueInternal operation response
type OpenQueueInternalResponse struct {
	// lplpRemoteQueueName:  On input, the server MUST ignore lplpRemoteQueueName. If hRemoteQueue
	// is 0x00000000 and the queue identified by pQueueFormat is located at a remote queue
	// manager, the server MUST set this string to a null-terminated path name, from which
	// the client can determine the computer name of the remote queue manager, as specified
	// in [MS-MQMQ] section 2.1.1.
	RemoteQueueName string `idl:"name:lplpRemoteQueueName;string;pointer:ptr" json:"remote_queue_name"`
	// pdwQMContext: A pointer to a variable to receive a DWORD value that identifies either
	// an OpenQueueDescriptor ([MS-MQDMPR] section 3.1.1.16) ADM element instance at the
	// server or a RemoteQueueProxyHandle (section 3.1.1.5) ADM element instance that contains
	// information pertaining to an OpenQueueDescriptor ADM element instance at a remote
	// server. When the client calls rpc_ACReceiveMessageEx (section 3.1.5.3), it specifies
	// a queue by providing the value that is returned by this parameter. On return, the
	// client MUST ignore pdwQMContext if the value returned via lplpRemoteQueueName is
	// non-NULL.
	QMContext uint32 `idl:"name:pdwQMContext" json:"qm_context"`
	// phQueue:  A pointer to a variable to receive a new RPC_QUEUE_HANDLE (section 2.2.1.1.2)
	// context handle. On return, the client MUST ignore phQueue if the value returned via
	// lplpRemoteQueueName is non-NULL.
	Queue *Queue `idl:"name:phQueue" json:"queue"`
	// Return: The rpc_QMOpenQueueInternal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenQueueInternalResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenQueueInternalOperation) *xxx_OpenQueueInternalOperation {
	if op == nil {
		op = &xxx_OpenQueueInternalOperation{}
	}
	if o == nil {
		return op
	}
	op.RemoteQueueName = o.RemoteQueueName
	op.QMContext = o.QMContext
	op.Queue = o.Queue
	op.Return = o.Return
	return op
}

func (o *OpenQueueInternalResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenQueueInternalOperation) {
	if o == nil {
		return
	}
	o.RemoteQueueName = op.RemoteQueueName
	o.QMContext = op.QMContext
	o.Queue = op.Queue
	o.Return = op.Return
}
func (o *OpenQueueInternalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenQueueInternalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenQueueInternalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseHandleOperation structure represents the rpc_ACCloseHandle operation
type xxx_CloseHandleOperation struct {
	Queue  *Queue `idl:"name:phQueue" json:"queue"`
	Return int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseHandleOperation) OpNum() int { return 20 }

func (o *xxx_CloseHandleOperation) OpName() string { return "/qmcomm/v1/rpc_ACCloseHandle" }

func (o *xxx_CloseHandleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseHandleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phQueue {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue != nil {
			if err := o.Queue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Queue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseHandleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phQueue {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue == nil {
			o.Queue = &Queue{}
		}
		if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseHandleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseHandleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phQueue {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue != nil {
			if err := o.Queue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Queue{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CloseHandleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phQueue {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue == nil {
			o.Queue = &Queue{}
		}
		if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
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

// CloseHandleRequest structure represents the rpc_ACCloseHandle operation request
type CloseHandleRequest struct {
	// phQueue:  MUST be a context handle acquired from the phQueue out-parameter of the
	// rpc_QMOpenQueueInternal method. On success, the server MUST set this parameter to
	// NULL.
	Queue *Queue `idl:"name:phQueue" json:"queue"`
}

func (o *CloseHandleRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseHandleOperation) *xxx_CloseHandleOperation {
	if op == nil {
		op = &xxx_CloseHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.Queue = o.Queue
	return op
}

func (o *CloseHandleRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseHandleOperation) {
	if o == nil {
		return
	}
	o.Queue = op.Queue
}
func (o *CloseHandleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseHandleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseHandleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseHandleResponse structure represents the rpc_ACCloseHandle operation response
type CloseHandleResponse struct {
	// phQueue:  MUST be a context handle acquired from the phQueue out-parameter of the
	// rpc_QMOpenQueueInternal method. On success, the server MUST set this parameter to
	// NULL.
	Queue *Queue `idl:"name:phQueue" json:"queue"`
	// Return: The rpc_ACCloseHandle return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseHandleResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseHandleOperation) *xxx_CloseHandleOperation {
	if op == nil {
		op = &xxx_CloseHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.Queue = o.Queue
	op.Return = o.Return
	return op
}

func (o *CloseHandleResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseHandleOperation) {
	if o == nil {
		return
	}
	o.Queue = op.Queue
	o.Return = op.Return
}
func (o *CloseHandleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseHandleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseHandleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseCursorOperation structure represents the rpc_ACCloseCursor operation
type xxx_CloseCursorOperation struct {
	Queue  *Queue `idl:"name:hQueue" json:"queue"`
	Cursor uint32 `idl:"name:hCursor" json:"cursor"`
	Return int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseCursorOperation) OpNum() int { return 22 }

func (o *xxx_CloseCursorOperation) OpName() string { return "/qmcomm/v1/rpc_ACCloseCursor" }

func (o *xxx_CloseCursorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseCursorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hQueue {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue != nil {
			if err := o.Queue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Queue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cursor); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseCursorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQueue {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue == nil {
			o.Queue = &Queue{}
		}
		if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cursor); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseCursorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseCursorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloseCursorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseCursorRequest structure represents the rpc_ACCloseCursor operation request
type CloseCursorRequest struct {
	// hQueue:  MUST contain the RPC_QUEUE_HANDLE (section 2.2.1.1.2) context handle passed
	// to rpc_ACCreateCursorEx when the cursor specified by hCursor was created.
	Queue *Queue `idl:"name:hQueue" json:"queue"`
	// hCursor:  MUST contain a DWORD value obtained from the pcc.hCursor out-parameter
	// of rpc_ACCreateCursorEx, or the reserved value 0x0000000b.
	Cursor uint32 `idl:"name:hCursor" json:"cursor"`
}

func (o *CloseCursorRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseCursorOperation) *xxx_CloseCursorOperation {
	if op == nil {
		op = &xxx_CloseCursorOperation{}
	}
	if o == nil {
		return op
	}
	op.Queue = o.Queue
	op.Cursor = o.Cursor
	return op
}

func (o *CloseCursorRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseCursorOperation) {
	if o == nil {
		return
	}
	o.Queue = op.Queue
	o.Cursor = op.Cursor
}
func (o *CloseCursorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseCursorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseCursorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseCursorResponse structure represents the rpc_ACCloseCursor operation response
type CloseCursorResponse struct {
	// Return: The rpc_ACCloseCursor return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseCursorResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseCursorOperation) *xxx_CloseCursorOperation {
	if op == nil {
		op = &xxx_CloseCursorOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *CloseCursorResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseCursorOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CloseCursorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseCursorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseCursorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCursorPropertiesOperation structure represents the rpc_ACSetCursorProperties operation
type xxx_SetCursorPropertiesOperation struct {
	Proxy        *Queue `idl:"name:hProxy" json:"proxy"`
	Cursor       uint32 `idl:"name:hCursor" json:"cursor"`
	RemoteCursor uint32 `idl:"name:hRemoteCursor" json:"remote_cursor"`
	Return       int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCursorPropertiesOperation) OpNum() int { return 23 }

func (o *xxx_SetCursorPropertiesOperation) OpName() string {
	return "/qmcomm/v1/rpc_ACSetCursorProperties"
}

func (o *xxx_SetCursorPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCursorPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hProxy {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Proxy != nil {
			if err := o.Proxy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Queue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Cursor); err != nil {
			return err
		}
	}
	// hRemoteCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RemoteCursor); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCursorPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hProxy {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Proxy == nil {
			o.Proxy = &Queue{}
		}
		if err := o.Proxy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Cursor); err != nil {
			return err
		}
	}
	// hRemoteCursor {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RemoteCursor); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCursorPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCursorPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCursorPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetCursorPropertiesRequest structure represents the rpc_ACSetCursorProperties operation request
type SetCursorPropertiesRequest struct {
	// hProxy: MUST contain the RPC_QUEUE_HANDLE (section 2.2.1.1.2) context handle passed
	// to rpc_ACCreateCursorEx when the cursor specified by hCursor was created.
	Proxy *Queue `idl:"name:hProxy" json:"proxy"`
	// hCursor: MUST contain a CursorProxy.Handle obtained from the pcc.hCursor out-parameter
	// of rpc_ACCreateCursorEx.
	Cursor uint32 `idl:"name:hCursor" json:"cursor"`
	// hRemoteCursor: MUST contain a Cursor.Handle for a remote cursor acquired from the
	// phCursor out-parameter of R_QMCreateRemoteCursor invoked at a remote queue manager.
	RemoteCursor uint32 `idl:"name:hRemoteCursor" json:"remote_cursor"`
}

func (o *SetCursorPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetCursorPropertiesOperation) *xxx_SetCursorPropertiesOperation {
	if op == nil {
		op = &xxx_SetCursorPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Proxy = o.Proxy
	op.Cursor = o.Cursor
	op.RemoteCursor = o.RemoteCursor
	return op
}

func (o *SetCursorPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCursorPropertiesOperation) {
	if o == nil {
		return
	}
	o.Proxy = op.Proxy
	o.Cursor = op.Cursor
	o.RemoteCursor = op.RemoteCursor
}
func (o *SetCursorPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetCursorPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCursorPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCursorPropertiesResponse structure represents the rpc_ACSetCursorProperties operation response
type SetCursorPropertiesResponse struct {
	// Return: The rpc_ACSetCursorProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCursorPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetCursorPropertiesOperation) *xxx_SetCursorPropertiesOperation {
	if op == nil {
		op = &xxx_SetCursorPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetCursorPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCursorPropertiesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetCursorPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetCursorPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCursorPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_HandleToFormatNameOperation structure represents the rpc_ACHandleToFormatName operation
type xxx_HandleToFormatNameOperation struct {
	Queue                     *Queue `idl:"name:hQueue" json:"queue"`
	FormatNameRPCBufferLength uint32 `idl:"name:dwFormatNameRPCBufferLen" json:"format_name_rpc_buffer_length"`
	FormatName                string `idl:"name:lpwcsFormatName;size_is:(dwFormatNameRPCBufferLen);length_is:(dwFormatNameRPCBufferLen);pointer:unique" json:"format_name"`
	Length                    uint32 `idl:"name:pdwLength" json:"length"`
	Return                    int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_HandleToFormatNameOperation) OpNum() int { return 26 }

func (o *xxx_HandleToFormatNameOperation) OpName() string {
	return "/qmcomm/v1/rpc_ACHandleToFormatName"
}

func (o *xxx_HandleToFormatNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.FormatName != "" && o.FormatNameRPCBufferLength == 0 {
		o.FormatNameRPCBufferLength = uint32(len(o.FormatName))
	}
	if o.FormatName != "" && o.FormatNameRPCBufferLength == 0 {
		o.FormatNameRPCBufferLength = uint32(len(o.FormatName))
	}
	if o.FormatNameRPCBufferLength > uint32(524288) {
		return fmt.Errorf("FormatNameRPCBufferLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HandleToFormatNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hQueue {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue != nil {
			if err := o.Queue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Queue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFormatNameRPCBufferLen {in} (1:{range=(0,524288), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.FormatNameRPCBufferLength); err != nil {
			return err
		}
	}
	// lpwcsFormatName {in, out} (1:{pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,size_is=dwFormatNameRPCBufferLen,length_is=dwFormatNameRPCBufferLen,string](wchar))
	{
		if o.FormatName != "" || o.FormatNameRPCBufferLength > 0 {
			_ptr_lpwcsFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.FormatNameRPCBufferLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.FormatNameRPCBufferLength)
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
				_FormatName_buf := utf16.Encode([]rune(o.FormatName))
				if uint64(len(_FormatName_buf)) > sizeInfo[0] {
					_FormatName_buf = _FormatName_buf[:sizeInfo[0]]
				}
				for i1 := range _FormatName_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_FormatName_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_FormatName_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FormatName, _ptr_lpwcsFormatName); err != nil {
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
	// pdwLength {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HandleToFormatNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQueue {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue == nil {
			o.Queue = &Queue{}
		}
		if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFormatNameRPCBufferLen {in} (1:{range=(0,524288), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.FormatNameRPCBufferLength); err != nil {
			return err
		}
	}
	// lpwcsFormatName {in, out} (1:{pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,size_is=dwFormatNameRPCBufferLen,length_is=dwFormatNameRPCBufferLen,string](wchar))
	{
		_ptr_lpwcsFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			var _FormatName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _FormatName_buf", sizeInfo[0])
			}
			_FormatName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _FormatName_buf {
				i1 := i1
				if err := w.ReadData(&_FormatName_buf[i1]); err != nil {
					return err
				}
			}
			o.FormatName = strings.TrimRight(string(utf16.Decode(_FormatName_buf)), ndr.ZeroString)
			return nil
		})
		_s_lpwcsFormatName := func(ptr interface{}) { o.FormatName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FormatName, _s_lpwcsFormatName, _ptr_lpwcsFormatName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwLength {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HandleToFormatNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_HandleToFormatNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpwcsFormatName {in, out} (1:{pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,size_is=dwFormatNameRPCBufferLen,length_is=dwFormatNameRPCBufferLen,string](wchar))
	{
		if o.FormatName != "" || o.FormatNameRPCBufferLength > 0 {
			_ptr_lpwcsFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.FormatNameRPCBufferLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.FormatNameRPCBufferLength)
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
				_FormatName_buf := utf16.Encode([]rune(o.FormatName))
				if uint64(len(_FormatName_buf)) > sizeInfo[0] {
					_FormatName_buf = _FormatName_buf[:sizeInfo[0]]
				}
				for i1 := range _FormatName_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_FormatName_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_FormatName_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FormatName, _ptr_lpwcsFormatName); err != nil {
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
	// pdwLength {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
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

func (o *xxx_HandleToFormatNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpwcsFormatName {in, out} (1:{pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,size_is=dwFormatNameRPCBufferLen,length_is=dwFormatNameRPCBufferLen,string](wchar))
	{
		_ptr_lpwcsFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			var _FormatName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _FormatName_buf", sizeInfo[0])
			}
			_FormatName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _FormatName_buf {
				i1 := i1
				if err := w.ReadData(&_FormatName_buf[i1]); err != nil {
					return err
				}
			}
			o.FormatName = strings.TrimRight(string(utf16.Decode(_FormatName_buf)), ndr.ZeroString)
			return nil
		})
		_s_lpwcsFormatName := func(ptr interface{}) { o.FormatName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FormatName, _s_lpwcsFormatName, _ptr_lpwcsFormatName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwLength {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
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

// HandleToFormatNameRequest structure represents the rpc_ACHandleToFormatName operation request
type HandleToFormatNameRequest struct {
	// hQueue:  MUST be an RPC_QUEUE_HANDLE (section 2.2.1.1.2) acquired from the phQueue
	// parameter of rpc_QMOpenQueueInternal (section 3.1.4.17). Prior to this method being
	// invoked, the queue MUST NOT have been deleted, and the queue handle MUST NOT have
	// been closed.
	Queue *Queue `idl:"name:hQueue" json:"queue"`
	// dwFormatNameRPCBufferLen:  Length of the buffer (in Unicode characters) provided
	// for the lpwcsFormatName parameter.
	FormatNameRPCBufferLength uint32 `idl:"name:dwFormatNameRPCBufferLen" json:"format_name_rpc_buffer_length"`
	// lpwcsFormatName:  Pointer to a Unicode character buffer into which the server writes
	// the format name (as specified in [MS-MQMQ]) for the queue identified by the hQueue
	// parameter. The character buffer MUST be null-terminated by the server prior to returning,
	// even if the provided buffer is not large enough to contain the entire format name
	// string. Can be NULL if dwFormatNameRPCBufferLen is 0x00000000. MUST NOT be NULL if
	// dwFormatNameRPCBufferLen is nonzero.
	FormatName string `idl:"name:lpwcsFormatName;size_is:(dwFormatNameRPCBufferLen);length_is:(dwFormatNameRPCBufferLen);pointer:unique" json:"format_name"`
	// pdwLength:  On input, the maximum number of Unicode characters to write to the lpwcsFormatName
	// buffer. This value MUST be equal to the dwFormatNameRPCBufferLen parameter. On return,
	// the server MUST update the value of this parameter to indicate the complete length
	// of the format name string for the queue identified by hQueue, without regard for
	// the size of the provided buffer.
	Length uint32 `idl:"name:pdwLength" json:"length"`
}

func (o *HandleToFormatNameRequest) xxx_ToOp(ctx context.Context, op *xxx_HandleToFormatNameOperation) *xxx_HandleToFormatNameOperation {
	if op == nil {
		op = &xxx_HandleToFormatNameOperation{}
	}
	if o == nil {
		return op
	}
	op.Queue = o.Queue
	op.FormatNameRPCBufferLength = o.FormatNameRPCBufferLength
	op.FormatName = o.FormatName
	op.Length = o.Length
	return op
}

func (o *HandleToFormatNameRequest) xxx_FromOp(ctx context.Context, op *xxx_HandleToFormatNameOperation) {
	if o == nil {
		return
	}
	o.Queue = op.Queue
	o.FormatNameRPCBufferLength = op.FormatNameRPCBufferLength
	o.FormatName = op.FormatName
	o.Length = op.Length
}
func (o *HandleToFormatNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *HandleToFormatNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_HandleToFormatNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// HandleToFormatNameResponse structure represents the rpc_ACHandleToFormatName operation response
type HandleToFormatNameResponse struct {
	// lpwcsFormatName:  Pointer to a Unicode character buffer into which the server writes
	// the format name (as specified in [MS-MQMQ]) for the queue identified by the hQueue
	// parameter. The character buffer MUST be null-terminated by the server prior to returning,
	// even if the provided buffer is not large enough to contain the entire format name
	// string. Can be NULL if dwFormatNameRPCBufferLen is 0x00000000. MUST NOT be NULL if
	// dwFormatNameRPCBufferLen is nonzero.
	FormatName string `idl:"name:lpwcsFormatName;size_is:(dwFormatNameRPCBufferLen);length_is:(dwFormatNameRPCBufferLen);pointer:unique" json:"format_name"`
	// pdwLength:  On input, the maximum number of Unicode characters to write to the lpwcsFormatName
	// buffer. This value MUST be equal to the dwFormatNameRPCBufferLen parameter. On return,
	// the server MUST update the value of this parameter to indicate the complete length
	// of the format name string for the queue identified by hQueue, without regard for
	// the size of the provided buffer.
	Length uint32 `idl:"name:pdwLength" json:"length"`
	// Return: The rpc_ACHandleToFormatName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *HandleToFormatNameResponse) xxx_ToOp(ctx context.Context, op *xxx_HandleToFormatNameOperation) *xxx_HandleToFormatNameOperation {
	if op == nil {
		op = &xxx_HandleToFormatNameOperation{}
	}
	if o == nil {
		return op
	}
	op.FormatName = o.FormatName
	op.Length = o.Length
	op.Return = o.Return
	return op
}

func (o *HandleToFormatNameResponse) xxx_FromOp(ctx context.Context, op *xxx_HandleToFormatNameOperation) {
	if o == nil {
		return
	}
	o.FormatName = op.FormatName
	o.Length = op.Length
	o.Return = op.Return
}
func (o *HandleToFormatNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *HandleToFormatNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_HandleToFormatNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PurgeQueueOperation structure represents the rpc_ACPurgeQueue operation
type xxx_PurgeQueueOperation struct {
	Queue  *Queue `idl:"name:hQueue" json:"queue"`
	Return int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_PurgeQueueOperation) OpNum() int { return 27 }

func (o *xxx_PurgeQueueOperation) OpName() string { return "/qmcomm/v1/rpc_ACPurgeQueue" }

func (o *xxx_PurgeQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PurgeQueueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hQueue {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue != nil {
			if err := o.Queue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Queue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_PurgeQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQueue {in} (1:{context_handle, alias=RPC_QUEUE_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Queue == nil {
			o.Queue = &Queue{}
		}
		if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PurgeQueueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PurgeQueueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PurgeQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PurgeQueueRequest structure represents the rpc_ACPurgeQueue operation request
type PurgeQueueRequest struct {
	// hQueue:  MUST be an RPC_QUEUE_HANDLE (section 2.2.1.1.2) obtained from the phQueue
	// parameter of the rpc_QMOpenQueueInternal (section 3.1.4.17) method. Prior to this
	// method being invoked, the queue MUST NOT have been deleted, and the queue handle
	// MUST NOT have been closed.
	Queue *Queue `idl:"name:hQueue" json:"queue"`
}

func (o *PurgeQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_PurgeQueueOperation) *xxx_PurgeQueueOperation {
	if op == nil {
		op = &xxx_PurgeQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.Queue = o.Queue
	return op
}

func (o *PurgeQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_PurgeQueueOperation) {
	if o == nil {
		return
	}
	o.Queue = op.Queue
}
func (o *PurgeQueueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PurgeQueueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PurgeQueueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PurgeQueueResponse structure represents the rpc_ACPurgeQueue operation response
type PurgeQueueResponse struct {
	// Return: The rpc_ACPurgeQueue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PurgeQueueResponse) xxx_ToOp(ctx context.Context, op *xxx_PurgeQueueOperation) *xxx_PurgeQueueOperation {
	if op == nil {
		op = &xxx_PurgeQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *PurgeQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_PurgeQueueOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PurgeQueueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PurgeQueueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PurgeQueueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryQMRegistryInternalOperation structure represents the R_QMQueryQMRegistryInternal operation
type xxx_QueryQMRegistryInternalOperation struct {
	QueryType  uint32 `idl:"name:dwQueryType" json:"query_type"`
	MQISServer string `idl:"name:lplpMQISServer;string" json:"mqis_server"`
	Return     int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryQMRegistryInternalOperation) OpNum() int { return 28 }

func (o *xxx_QueryQMRegistryInternalOperation) OpName() string {
	return "/qmcomm/v1/R_QMQueryQMRegistryInternal"
}

func (o *xxx_QueryQMRegistryInternalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryQMRegistryInternalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwQueryType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QueryType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryQMRegistryInternalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwQueryType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QueryType); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryQMRegistryInternalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryQMRegistryInternalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lplpMQISServer {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.MQISServer != "" {
			_ptr_lplpMQISServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.MQISServer); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.MQISServer, _ptr_lplpMQISServer); err != nil {
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

func (o *xxx_QueryQMRegistryInternalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lplpMQISServer {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_lplpMQISServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.MQISServer); err != nil {
				return err
			}
			return nil
		})
		_s_lplpMQISServer := func(ptr interface{}) { o.MQISServer = *ptr.(*string) }
		if err := w.ReadPointer(&o.MQISServer, _s_lplpMQISServer, _ptr_lplpMQISServer); err != nil {
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

// QueryQMRegistryInternalRequest structure represents the R_QMQueryQMRegistryInternal operation request
type QueryQMRegistryInternalRequest struct {
	// dwQueryType: Specifies the type and format of the data to return to the caller via
	// the lplpMQISServer parameter. MUST be one of the values in the following table:
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 | A comma-delimited list of MQIS server names configured on the supporting server. |
	//	|            | This value is retrieved from the DirectoryServerList attribute of the server's   |
	//	|            | LocalQueueManager ADM element instance.                                          |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The server's default time-to-reach-queue message property value, expressed in    |
	//	|            | seconds, converted to a string.<59><60>                                          |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The GUID that represents the entire MSMQ forest.<61> See following for the       |
	//	|            | curly braced GUID string representation to use. The string uses the "braceless"  |
	//	|            | format.                                                                          |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000003 | A string representation of the supporting server version.<62>                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 | The content of the Identifier attribute of the server's LocalQueueManager ADM    |
	//	|            | element instance. The curly braced GUID string representation uses a "braceless" |
	//	|            | format given following.<63><64>                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//
	// The format for the comma-delimited list of MQIS server names (0x00000000) is given
	// by the following augmented BNF:
	//
	// The GUID string for the MSMQ forest (0x00000002) uses the "braceless" format depicted
	// in the following augmented BNF:
	//
	// The string format used for the supporting server version (0x00000003), depicted in
	// augmented BNF, is as follows:
	//
	// The GUID for the server queue manager (0x00000004) uses the following "braceless"
	// format, depicted in augmented BNF:
	QueryType uint32 `idl:"name:dwQueryType" json:"query_type"`
}

func (o *QueryQMRegistryInternalRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryQMRegistryInternalOperation) *xxx_QueryQMRegistryInternalOperation {
	if op == nil {
		op = &xxx_QueryQMRegistryInternalOperation{}
	}
	if o == nil {
		return op
	}
	op.QueryType = o.QueryType
	return op
}

func (o *QueryQMRegistryInternalRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryQMRegistryInternalOperation) {
	if o == nil {
		return
	}
	o.QueryType = op.QueryType
}
func (o *QueryQMRegistryInternalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryQMRegistryInternalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryQMRegistryInternalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryQMRegistryInternalResponse structure represents the R_QMQueryQMRegistryInternal operation response
type QueryQMRegistryInternalResponse struct {
	// lplpMQISServer:  On success, the server returns the string indicated by dwQueryType
	// through this parameter. The server can set this parameter to NULL in the event of
	// an error.
	MQISServer string `idl:"name:lplpMQISServer;string" json:"mqis_server"`
	// Return: The R_QMQueryQMRegistryInternal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryQMRegistryInternalResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryQMRegistryInternalOperation) *xxx_QueryQMRegistryInternalOperation {
	if op == nil {
		op = &xxx_QueryQMRegistryInternalOperation{}
	}
	if o == nil {
		return op
	}
	op.MQISServer = o.MQISServer
	op.Return = o.Return
	return op
}

func (o *QueryQMRegistryInternalResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryQMRegistryInternalOperation) {
	if o == nil {
		return
	}
	o.MQISServer = op.MQISServer
	o.Return = op.Return
}
func (o *QueryQMRegistryInternalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryQMRegistryInternalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryQMRegistryInternalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRTQMServerPortOperation structure represents the R_QMGetRTQMServerPort operation
type xxx_GetRTQMServerPortOperation struct {
	IP     uint32 `idl:"name:fIP" json:"ip"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRTQMServerPortOperation) OpNum() int { return 31 }

func (o *xxx_GetRTQMServerPortOperation) OpName() string { return "/qmcomm/v1/R_QMGetRTQMServerPort" }

func (o *xxx_GetRTQMServerPortOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRTQMServerPortOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// fIP {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.IP); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRTQMServerPortOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// fIP {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.IP); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRTQMServerPortOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRTQMServerPortOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRTQMServerPortOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetRTQMServerPortRequest structure represents the R_QMGetRTQMServerPort operation request
type GetRTQMServerPortRequest struct {
	// fIP:  Specifies the interface for which a port value is to be returned. One of the
	// following values MUST be specified; otherwise, this method MUST return 0x00000000
	// to indicate failure.
	//
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	|                          |                                                                                  |
	//	|          VALUE           |                                     MEANING                                      |
	//	|                          |                                                                                  |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| IP_HANDSHAKE 0x00000000  | Requests that the server return the RPC port number for the qmcomm and qmcomm2   |
	//	|                          | interfaces bound to TCP/IP. The default port number is 2103.                     |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| IP_READ 0x00000001       | Requests that the server return the RPC port number for the qm2qm interface, as  |
	//	|                          | specified in [MS-MQQP], bound to TCP/IP. The default port number is 2105.        |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| IPX_HANDSHAKE 0x00000002 | Requests that the server return the RPC port number for the qmcomm and qmcomm2   |
	//	|                          | interfaces bound to SPX.<65> The default port number is 2103.                    |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| IPX_READ 0x00000003      | Requests that the server return the RPC port number for the qm2qm interface, as  |
	//	|                          | specified in [MS-MQQP], bound to SPX.<66> The default port number is 2105.       |
	//	+--------------------------+----------------------------------------------------------------------------------+
	IP uint32 `idl:"name:fIP" json:"ip"`
}

func (o *GetRTQMServerPortRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRTQMServerPortOperation) *xxx_GetRTQMServerPortOperation {
	if op == nil {
		op = &xxx_GetRTQMServerPortOperation{}
	}
	if o == nil {
		return op
	}
	op.IP = o.IP
	return op
}

func (o *GetRTQMServerPortRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRTQMServerPortOperation) {
	if o == nil {
		return
	}
	o.IP = op.IP
}
func (o *GetRTQMServerPortRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRTQMServerPortRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRTQMServerPortOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRTQMServerPortResponse structure represents the R_QMGetRTQMServerPort operation response
type GetRTQMServerPortResponse struct {
	// Return: The R_QMGetRTQMServerPort return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetRTQMServerPortResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRTQMServerPortOperation) *xxx_GetRTQMServerPortOperation {
	if op == nil {
		op = &xxx_GetRTQMServerPortOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *GetRTQMServerPortResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRTQMServerPortOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *GetRTQMServerPortResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRTQMServerPortResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRTQMServerPortOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
