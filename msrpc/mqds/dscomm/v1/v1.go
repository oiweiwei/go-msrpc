package dscomm

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
	mqds "github.com/oiweiwei/go-msrpc/msrpc/mqds"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
	_ = mqmq.GoPackage
	_ = mqds.GoPackage
)

var (
	// import guard
	GoPackage = "mqds"
)

var (
	// Syntax UUID
	DscommSyntaxUUID = &uuid.UUID{TimeLow: 0x77df7a80, TimeMid: 0xf298, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0x58, Node: [6]uint8{0x0, 0xa0, 0x24, 0xc4, 0x80, 0xa8}}
	// Syntax ID
	DscommSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: DscommSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// dscomm interface.
type DscommClient interface {

	// This method creates a new directory object, assigns the specified properties and
	// security descriptor to that directory object, and returns a unique GUID identifier
	// for that directory object.
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_ILLEGAL_PROPERTY_VALUE (0xC00E0018)
	//
	// MQ_ERROR_ILLEGAL_ENTERPRISE_OPERATION (0xC00E0071)
	//
	// MQ_ERROR (0xC00E0001)
	//
	// MQ_ERROR_DS_ERROR (0xC00E0043)
	//
	// MQDS_OBJECT_NOT_FOUND (0xC00E050F)
	//
	// E_ADS_PROPERTY_NOT_FOUND (0x8000500D)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	CreateObject(context.Context, *CreateObjectRequest, ...dcerpc.CallOption) (*CreateObjectResponse, error)

	// This method deletes a directory object specified by a directory service pathname.
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// The server SHOULD enforce appropriate security measures to make sure that the caller
	// has the required permissions to execute this routine. If the server enforces security
	// measures, and the caller does not have the required credentials, the server MUST
	// return an error.
	DeleteObject(context.Context, *DeleteObjectRequest, ...dcerpc.CallOption) (*DeleteObjectResponse, error)

	// This method returns the properties associated with a directory object specified by
	// a directory service pathname.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT. Additionally, if a failure HRESULT is returned,
	// the client MUST disregard all out-parameter values.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// MQ_ERROR_ILLEGAL_PROPID (0xC00E0039)
	//
	// MQ_ERROR (0xC00E0001)
	//
	// MQ_ERROR_DS_ERROR (0xC00E0043)
	//
	// MQDS_OBJECT_NOT_FOUND (0xC00E050F)
	//
	// E_ADS_PROPERTY_NOT_FOUND (0x8000500D)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// This method sets the specified properties for a directory object specified by a directory
	// service pathname.
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_ILLEGAL_PROPID (0xC00E0039)
	//
	// MQ_ERROR (0xC00E0001)
	//
	// MQ_ERROR_DS_ERROR (0xC00E0043)
	//
	// MQDS_OBJECT_NOT_FOUND (0xC00E050F)
	//
	// E_ADS_PROPERTY_NOT_FOUND (0x8000500D)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	SetProperties(context.Context, *SetPropertiesRequest, ...dcerpc.CallOption) (*SetPropertiesResponse, error)

	// This method gets security properties for a directory object specified by a directory
	// service pathname.
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028)
	//
	// MQ_ERROR_SECURITY_DESCRIPTOR_TOO_SMALL (0xC00E0023)
	//
	// MQDS_WRONG_OBJ_TYPE (0xC00E0506)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetObjectSecurity(context.Context, *GetObjectSecurityRequest, ...dcerpc.CallOption) (*GetObjectSecurityResponse, error)

	// This method sets security properties for a directory object specified by a directory
	// service pathname.
	//
	// Return Values: If the method succeeds, the return value is 0x00000000. If the method
	// fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	SetObjectSecurity(context.Context, *SetObjectSecurityRequest, ...dcerpc.CallOption) (*SetObjectSecurityResponse, error)

	// This method performs a query over the directory objects and returns an RPC context
	// handle that can be used to retrieve the result set through a subsequent series of
	// calls to S_DSLookupNext (section 3.1.4.18). When the client has no further use of
	// the RPC context handle returned from this method, the client can close the context
	// handle through a call to S_DSLookupEnd (section 3.1.4.19).
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	LookupBegin(context.Context, *LookupBeginRequest, ...dcerpc.CallOption) (*LookupBeginResponse, error)

	// This method returns a portion of the data from the result set computed in a previous
	// call to S_DSLookupBegin (section 3.1.4.17) and updates the cursor index to the first
	// directory object that has not yet been returned to the client.
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	LookupNext(context.Context, *LookupNextRequest, ...dcerpc.CallOption) (*LookupNextResponse, error)

	// This method closes an opened RPC context handle created from a previous call to S_DSLookupBegin
	// (section 3.1.4.17).
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	LookupEnd(context.Context, *LookupEndRequest, ...dcerpc.CallOption) (*LookupEndResponse, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// This method deletes a directory object specified by an object identifier.
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// The server SHOULD enforce appropriate security measures to make sure that the caller
	// has the required permissions to execute this routine. If the server enforces security
	// measures, and the caller does not have the required credentials, the server MUST
	// return an error.
	DeleteObjectGUID(context.Context, *DeleteObjectGUIDRequest, ...dcerpc.CallOption) (*DeleteObjectGUIDResponse, error)

	// This method returns properties associated with a directory object specified by an
	// object identifier.
	//
	// Return Values:  On success, this method MUST return MQ_OK (0x00000000); otherwise,
	// the server MUST return a failure HRESULT. Additionally, if a failure HRESULT is returned,
	// the client MUST disregard all out-parameter values.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028)
	//
	// MQ_ERROR_INVALID_PARAMETER (0xC00E0006)
	//
	// MQ_ERROR_ILLEGAL_PROPID (0xC00E0039)
	//
	// MQ_ERROR (0xC00E0001)
	//
	// MQ_ERROR_DS_ERROR (0xC00E0043)
	//
	// MQDS_OBJECT_NOT_FOUND (0xC00E050F)
	//
	// E_ADS_PROPERTY_NOT_FOUND (0x8000500D)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetPropertiesGUID(context.Context, *GetPropertiesGUIDRequest, ...dcerpc.CallOption) (*GetPropertiesGUIDResponse, error)

	// This method sets properties for a directory object specified by an object identifier.
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_ILLEGAL_PROPID (0xC00E0039)
	//
	// MQ_ERROR (0xC00E0001)
	//
	// MQ_ERROR_DS_ERROR (0xC00E0043)
	//
	// MQDS_OBJECT_NOT_FOUND (0xC00E050F)
	//
	// E_ADS_PROPERTY_NOT_FOUND (0x8000500D)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	SetPropertiesGUID(context.Context, *SetPropertiesGUIDRequest, ...dcerpc.CallOption) (*SetPropertiesGUIDResponse, error)

	// This method retrieves the security descriptor for a directory object specified by
	// an object identifier.
	//
	// Return Values: If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_SECURITY_DESCRIPTOR_TOO_SMALL (0xC00E0023)
	//
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	GetObjectSecurityGUID(context.Context, *GetObjectSecurityGUIDRequest, ...dcerpc.CallOption) (*GetObjectSecurityGUIDResponse, error)

	// This method sets security properties for a directory object specified by an object
	// identifier.
	//
	// Return Values: If the method succeeds, the return value is 0x00000000. If the method
	// fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol extension, as specified in [MS-RPCE].
	SetObjectSecurityGUID(context.Context, *SetObjectSecurityGUIDRequest, ...dcerpc.CallOption) (*SetObjectSecurityGUIDResponse, error)

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	// Opnum18NotUsedOnWire operation.
	// Opnum18NotUsedOnWire

	// This method sets properties associated with a machine object specified by a directory
	// service pathname. This method is intended for use by a queue manager to manipulate
	// its own directory service object of type MQDS_MACHINE. Client applications SHOULD
	// NOT use this method.<76>
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_ILLEGAL_PROPID (0xC00E0039)
	//
	// MQ_ERROR (0xC00E0001)
	//
	// MQ_ERROR_DS_ERROR (0xC00E0043)
	//
	// MQDS_OBJECT_NOT_FOUND (0xC00E050F)
	//
	// E_ADS_PROPERTY_NOT_FOUND (0x8000500D)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	QMSetMachineProperties(context.Context, *QMSetMachinePropertiesRequest, ...dcerpc.CallOption) (*QMSetMachinePropertiesResponse, error)

	// This method returns a list of Backup Site Controllers (BSCs) associated with a specified
	// site. The client calls this method to enumerate the BSCs associated with sites in
	// the configured list of sites in the enterprise.
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code. The server MUST return
	// MQDS_E_NO_MORE_DATA (0xC00E0523), if pdwIndex is an invalid index, into the configured
	// list of sites in the enterprise.
	//
	// MQ_OK (0x00000000)
	//
	// MQDS_E_NO_MORE_DATA (0xC00E0523)
	//
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	//
	// The S_DSCreateServersCache method returns information on Backup Site Controllers
	// (BSCs) associated with sites in the enterprise. The client uses this method to iterate
	// through the list of sites in the enterprise by calling the method repeatedly. Before
	// the first invocation, the client sets the pdwIndex parameter to 0. After each successful
	// invocation, the client increments the pdwIndex parameter by 1 and calls the method
	// again. The client repeats this sequence until the method call returns an error.
	CreateServersCache(context.Context, *CreateServersCacheRequest, ...dcerpc.CallOption) (*CreateServersCacheResponse, error)

	// This method is a callback method called by the server during a client call to S_DSQMSetMachineProperties.
	// Through this method, the server provides a challenge that the client must sign to
	// authenticate itself.
	//
	// Return Values:  This method is obsolete. The server SHOULD NOT call this method,
	// and the client SHOULD return MQ_ERROR_NOT_SUPPORTED (0xC00E03EB).<146> If the method
	// succeeds, the return value is 0. If the method fails, the return value is an implementation-specific
	// error code.
	//
	// ERROR_SUCCESS (0x00000000)
	//
	// MQ_ERROR_NOT_SUPPORTED (0xC00E03EB)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	QMSetMachinePropertiesSignProc(context.Context, *QMSetMachinePropertiesSignProcRequest, ...dcerpc.CallOption) (*QMSetMachinePropertiesSignProcResponse, error)

	// This method retrieves the security descriptor for a directory object specified by
	// an object identifier.
	//
	// Return Values: If the method succeeds, the return value is 0x00000000. If the method
	// fails, the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_SECURITY_DESCRIPTOR_TOO_SMALL (0xC00E0023)
	//
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	QMGetObjectSecurity(context.Context, *QMGetObjectSecurityRequest, ...dcerpc.CallOption) (*QMGetObjectSecurityResponse, error)

	// S_DSQMGetObjectSecurityChallengeResponceProc is a callback method called by the server
	// during a client call to S_DSQMGetObjectSecurity. Through this method, the server
	// provides a challenge that the client must sign to authenticate itself.
	//
	// Return Values:  This method is obsolete. The server SHOULD NOT call this method,
	// and the client SHOULD return MQ_ERROR_NOT_SUPPORTED (0xC00E03EB).<147> If the method
	// succeeds, the return value is 0. If the method fails, the return value is an implementation-specific
	// error code.
	//
	// ERROR_SUCCESS (0x00000000)
	//
	// MQ_ERROR_NOT_SUPPORTED (0xC00E03EB)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol extension, as specified in [MS-RPCE].
	QMGetObjectSecurityChallengeResponseProc(context.Context, *QMGetObjectSecurityChallengeResponseProcRequest, ...dcerpc.CallOption) (*QMGetObjectSecurityChallengeResponseProcResponse, error)

	// This method is a callback method called by the server during a client call to S_DSValidateServer.
	// These two methods are used to tunnel a GSS (as specified in [RFC2743]) security negotiation
	// to provide mutual authentication between the client and server.
	//
	// Return Values:  If the method succeeds, and the negotiation is complete, the return
	// value is 0. If the method succeeds, and the negotiation is not complete, the return
	// value is SEC_I_CONTINUE_NEEDED (0x00090312). If the method fails, the return value
	// is an implementation-specific error code.
	//
	// ERROR_SUCCESS (0x00000000)
	//
	// SEC_I_CONTINUE_NEEDED (0x00090312)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC extension protocol, as specified in [MS-RPCE].
	//
	// The caller MUST supply an input_token in pServerBuff computed through a call to GSS_Accept_sec_context.
	// The receiver MUST process this input_token through a call to GSS_Init_sec_context,
	// generating an output_token that MUST be returned in pServerBuff.
	//
	// If GSS_Init_sec_context returns GSS_S_CONTINUE_NEEDED, this is a signal that the
	// negotiation is not complete. The receiver MUST return SEC_I_CONTINUE_NEEDED (0x00090312).
	//
	// If GSS_Init_sec_context returns GSS_S_COMPLETE, the negotiation is complete. The
	// receiver MUST save the output context handle in the GSS security context state associated
	// with the dwContext parameter. The receiver MUST return SEC_E_OK (0x00000000).
	InitSecurityContext(context.Context, *InitSecurityContextRequest, ...dcerpc.CallOption) (*InitSecurityContextResponse, error)

	// This method performs mutual authentication between the client and server, and establishes
	// a security context, as specified in [RFC2743]. The server uses the security context
	// to construct a digital signature in subsequent method calls of this protocol, and
	// the client uses the security context to validate the digital signature. The digital
	// signature is used in methods that return data to the client. It allows the client
	// to authenticate the source of the data and ensures that it has not been tampered
	// with en route from the server to the client.
	//
	// When the client has no further use of the RPC context handle returned from this method,
	// it releases the handle through a subsequent call to S_DSCloseServerHandle (section
	// 3.1.4.3).
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR_DS_ERROR (0xC00E0043)
	//
	// MQDS_E_CANT_INIT_SERVER_AUTH (0xC00E052B)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	//
	// If the dwClientBuffSize parameter is zero, the server MUST initialize an empty security
	// context in the security context table associated with the pphServerAuth parameter
	// and MUST NOT invoke client callback methods. In addition, all signature parameters
	// in other methods in this protocol that return server signatures MUST be set to all
	// zeros.
	//
	// If the dwClientBuffSize parameter is nonzero, the server MUST construct a valid non-empty
	// security context in the security context table associated with the pphServerAuth
	// parameter.
	//
	// The client MUST supply an input_token in the pClientBuff parameter, acquired through
	// an initial call to GSS_Init_sec_context, as specified in [RFC2743] section 2.2.1.
	// The server MUST perform the following processing.
	//
	// *
	//
	// Let inputToken point to an input_token initialized to the pClientBuff parameter.
	//
	// *
	//
	// Call *GSS_Accept_sec_context* , as specified in [RFC2743] section 2.2.2, with the
	// input_token contained in inputToken and a NULL input_context_handle.
	//
	// *
	//
	// If *GSS_Accept_sec_context* returns GSS_S_COMPLETE, the negotiation is complete.
	// The server MUST allocate a PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle and
	// MUST allocate a security context entry in the security table keyed by the context
	// handle. The server MUST associate the GSS security context (output_context_handle
	// from the GSSAPI call) with the security context entry. The server MUST return the
	// RPC context handle in the pphServerAuth parameter and return MQ_OK.
	//
	// *
	//
	// If *GSS_Accept_sec_context* returns GSS_S_CONTINUE_NEEDED, the server MUST issue
	// a callback to the client through the S_InitSecCtx (section 3.2.4.3) method with the
	// dwContext parameter set to the value supplied by the client in the dwContext parameter
	// and the pServerBuff parameter set to the output_token returned from *GSS_Accept_sec_context*.
	// When the callback to S_InitSecCtx returns, the server MUST set inputToken to the
	// pClientBuff parameter returned by S_InitSecCtx and MUST continue at step 2.
	//
	// *
	//
	// If *GSS_Accept_sec_context* returns any other value, the server MUST take no further
	// action and return a failure HRESULT ( ../ms-dtyp/a9046ed2-bfb2-4d56-a719-2824afce59ac
	// ).
	//
	// The GSS security context is used by the server in subsequent calls to GSSWrap, as
	// specified in [RFC2743] section 2.3.3.
	//
	// On successful return, the client MUST retrieve the GSS security context associated
	// with the dwContext parameter and MUST associate it with the PCONTEXT_HANDLE_SERVER_AUTH_TYPE
	// RPC context handle returned in the pphServerAuth parameter. The GSS security context
	// will be used by the client in subsequent calls to GSSUnwrap, as specified in [RFC2743]
	// section 2.3.4.
	ValidateServer(context.Context, *ValidateServerRequest, ...dcerpc.CallOption) (*ValidateServerResponse, error)

	// This method closes the RPC context handle returned by a previous call to S_DSValidateServer.
	// The server releases resources associated with the RPC context handle.
	//
	// Return Values:  If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	CloseServer(context.Context, *CloseServerRequest, ...dcerpc.CallOption) (*CloseServerResponse, error)

	// Opnum24NotUsedOnWire operation.
	// Opnum24NotUsedOnWire

	// Opnum25NotUsedOnWire operation.
	// Opnum25NotUsedOnWire

	// Opnum26NotUsedOnWire operation.
	// Opnum26NotUsedOnWire

	// This method returns the RPC endpoint port for a transport protocol. The client establishes
	// a new binding to the server by using the returned port number.
	//
	// Return Values:  If the method succeeds, the return value is the RPC endpoint port
	// number. If the method fails, the server MUST return 0.
	//
	// MQ_OK (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	GetServerPort(context.Context, *GetServerPortRequest, ...dcerpc.CallOption) (*GetServerPortResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Type structure represents PCONTEXT_HANDLE_TYPE RPC structure.
type Type dcetypes.ContextHandle

func (o *Type) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Type) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Type) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Type) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ServerAuthType structure represents PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC structure.
type ServerAuthType dcetypes.ContextHandle

func (o *ServerAuthType) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *ServerAuthType) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerAuthType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ServerAuthType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DeleteType structure represents PCONTEXT_HANDLE_DELETE_TYPE RPC structure.
type DeleteType dcetypes.ContextHandle

func (o *DeleteType) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *DeleteType) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DeleteType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DeleteType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type xxx_DefaultDscommClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultDscommClient) CreateObject(ctx context.Context, in *CreateObjectRequest, opts ...dcerpc.CallOption) (*CreateObjectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...dcerpc.CallOption) (*DeleteObjectResponse, error) {
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

func (o *xxx_DefaultDscommClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) SetProperties(ctx context.Context, in *SetPropertiesRequest, opts ...dcerpc.CallOption) (*SetPropertiesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) GetObjectSecurity(ctx context.Context, in *GetObjectSecurityRequest, opts ...dcerpc.CallOption) (*GetObjectSecurityResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetObjectSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) SetObjectSecurity(ctx context.Context, in *SetObjectSecurityRequest, opts ...dcerpc.CallOption) (*SetObjectSecurityResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetObjectSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) LookupBegin(ctx context.Context, in *LookupBeginRequest, opts ...dcerpc.CallOption) (*LookupBeginResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupBeginResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) LookupNext(ctx context.Context, in *LookupNextRequest, opts ...dcerpc.CallOption) (*LookupNextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupNextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) LookupEnd(ctx context.Context, in *LookupEndRequest, opts ...dcerpc.CallOption) (*LookupEndResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupEndResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) DeleteObjectGUID(ctx context.Context, in *DeleteObjectGUIDRequest, opts ...dcerpc.CallOption) (*DeleteObjectGUIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteObjectGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) GetPropertiesGUID(ctx context.Context, in *GetPropertiesGUIDRequest, opts ...dcerpc.CallOption) (*GetPropertiesGUIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPropertiesGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) SetPropertiesGUID(ctx context.Context, in *SetPropertiesGUIDRequest, opts ...dcerpc.CallOption) (*SetPropertiesGUIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetPropertiesGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) GetObjectSecurityGUID(ctx context.Context, in *GetObjectSecurityGUIDRequest, opts ...dcerpc.CallOption) (*GetObjectSecurityGUIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetObjectSecurityGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) SetObjectSecurityGUID(ctx context.Context, in *SetObjectSecurityGUIDRequest, opts ...dcerpc.CallOption) (*SetObjectSecurityGUIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetObjectSecurityGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) QMSetMachineProperties(ctx context.Context, in *QMSetMachinePropertiesRequest, opts ...dcerpc.CallOption) (*QMSetMachinePropertiesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QMSetMachinePropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) CreateServersCache(ctx context.Context, in *CreateServersCacheRequest, opts ...dcerpc.CallOption) (*CreateServersCacheResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateServersCacheResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) QMSetMachinePropertiesSignProc(ctx context.Context, in *QMSetMachinePropertiesSignProcRequest, opts ...dcerpc.CallOption) (*QMSetMachinePropertiesSignProcResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QMSetMachinePropertiesSignProcResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) QMGetObjectSecurity(ctx context.Context, in *QMGetObjectSecurityRequest, opts ...dcerpc.CallOption) (*QMGetObjectSecurityResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QMGetObjectSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) QMGetObjectSecurityChallengeResponseProc(ctx context.Context, in *QMGetObjectSecurityChallengeResponseProcRequest, opts ...dcerpc.CallOption) (*QMGetObjectSecurityChallengeResponseProcResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QMGetObjectSecurityChallengeResponseProcResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) InitSecurityContext(ctx context.Context, in *InitSecurityContextRequest, opts ...dcerpc.CallOption) (*InitSecurityContextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InitSecurityContextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) ValidateServer(ctx context.Context, in *ValidateServerRequest, opts ...dcerpc.CallOption) (*ValidateServerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ValidateServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) CloseServer(ctx context.Context, in *CloseServerRequest, opts ...dcerpc.CallOption) (*CloseServerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) GetServerPort(ctx context.Context, in *GetServerPortRequest, opts ...dcerpc.CallOption) (*GetServerPortResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetServerPortResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDscommClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDscommClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewDscommClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DscommClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DscommSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultDscommClient{cc: cc}, nil
}

// xxx_CreateObjectOperation structure represents the S_DSCreateObject operation
type xxx_CreateObjectOperation struct {
	ObjectType               uint32                  `idl:"name:dwObjectType" json:"object_type"`
	PathName                 string                  `idl:"name:pwcsPathName;string;pointer:unique" json:"path_name"`
	SecurityDescriptorLength uint32                  `idl:"name:dwSDLength" json:"security_descriptor_length"`
	SecurityDescriptor       []byte                  `idl:"name:SecurityDescriptor;size_is:(dwSDLength);pointer:unique" json:"security_descriptor"`
	CreatePartition          uint32                  `idl:"name:cp" json:"create_partition"`
	Property                 []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var                      []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	ObjectGUID               *dtyp.GUID              `idl:"name:pObjGuid;pointer:unique" json:"object_guid"`
	Return                   int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateObjectOperation) OpNum() int { return 0 }

func (o *xxx_CreateObjectOperation) OpName() string { return "/dscomm/v1/S_DSCreateObject" }

func (o *xxx_CreateObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.SecurityDescriptorLength == 0 {
		o.SecurityDescriptorLength = uint32(len(o.SecurityDescriptor))
	}
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
	}
	if o.SecurityDescriptorLength > uint32(524288) {
		return fmt.Errorf("SecurityDescriptorLength is out of range")
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

func (o *xxx_CreateObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.PathName != "" {
			_ptr_pwcsPathName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.PathName, _ptr_pwcsPathName); err != nil {
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
	// dwSDLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.WriteData(o.SecurityDescriptorLength); err != nil {
			return err
		}
	}
	// SecurityDescriptor {in} (1:{pointer=unique}*(1)[dim:0,size_is=dwSDLength](uchar))
	{
		if o.SecurityDescriptor != nil || o.SecurityDescriptorLength > 0 {
			_ptr_SecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.SecurityDescriptorLength)
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
			if err := w.WritePointer(&o.SecurityDescriptor, _ptr_SecurityDescriptor); err != nil {
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
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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
	// pObjGuid {in, out} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectGUID != nil {
			_ptr_pObjGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectGUID != nil {
					if err := o.ObjectGUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectGUID, _ptr_pObjGuid); err != nil {
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

func (o *xxx_CreateObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwcsPathName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
				return err
			}
			return nil
		})
		_s_pwcsPathName := func(ptr interface{}) { o.PathName = *ptr.(*string) }
		if err := w.ReadPointer(&o.PathName, _s_pwcsPathName, _ptr_pwcsPathName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwSDLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.ReadData(&o.SecurityDescriptorLength); err != nil {
			return err
		}
	}
	// SecurityDescriptor {in} (1:{pointer=unique}*(1)[dim:0,size_is=dwSDLength](uchar))
	{
		_ptr_SecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		_s_SecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.SecurityDescriptor, _s_SecurityDescriptor, _ptr_SecurityDescriptor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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
	// pObjGuid {in, out} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_pObjGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectGUID == nil {
				o.ObjectGUID = &dtyp.GUID{}
			}
			if err := o.ObjectGUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pObjGuid := func(ptr interface{}) { o.ObjectGUID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ObjectGUID, _s_pObjGuid, _ptr_pObjGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pObjGuid {in, out} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectGUID != nil {
			_ptr_pObjGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectGUID != nil {
					if err := o.ObjectGUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectGUID, _ptr_pObjGuid); err != nil {
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

func (o *xxx_CreateObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pObjGuid {in, out} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_pObjGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectGUID == nil {
				o.ObjectGUID = &dtyp.GUID{}
			}
			if err := o.ObjectGUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pObjGuid := func(ptr interface{}) { o.ObjectGUID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.ObjectGUID, _s_pObjGuid, _ptr_pObjGuid); err != nil {
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

// CreateObjectRequest structure represents the S_DSCreateObject operation request
type CreateObjectRequest struct {
	// dwObjectType:  Specifies the type of directory object to create. MUST be set to
	// one of the object types, as specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pwcsPathName:  Pointer to a NULL-terminated 16-bit Unicode string that MUST contain
	// the directory service pathname, as specified in section 2.2.9 for the object to be
	// created in the directory service. If the object is of a type that is not referenced
	// by pathname, as specified in section 2.2.9, this pointer MUST be NULL.
	PathName string `idl:"name:pwcsPathName;string;pointer:unique" json:"path_name"`
	// dwSDLength:  MUST contain the size in bytes of the buffer pointed to by SecurityDescriptor.
	SecurityDescriptorLength uint32 `idl:"name:dwSDLength" json:"security_descriptor_length"`
	// SecurityDescriptor:  MUST contain a pointer to a security descriptor, as specified
	// in [MS-DTYP] section 2.4.6, or NULL.
	SecurityDescriptor []byte `idl:"name:SecurityDescriptor;size_is:(dwSDLength);pointer:unique" json:"security_descriptor"`
	// cp:  MUST be set to the size (in elements) of the arrays aProp and apVar. The arrays
	// aProp and apVar MUST have an identical number of elements and MUST contain at least
	// one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  An array of property identifiers of properties to associate with the created
	// object. Each element MUST specify a value from the property identifiers table as
	// specified in section 2.2.10.1, for the directory object type specified in dwObjectType.
	// Each element MUST specify the property identifier for the corresponding property
	// value at the same element index in apVar. The array MUST contain at least one element.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar:  An array of property values to associate with the created object. Each element
	// MUST specify the property value for the corresponding property identifier at the
	// same element index in aProp. The array MUST contain at least one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// pObjGuid: SHOULD be set by the server to the GUID of the created object if the dwObjectType
	// is equal to MQDS_QUEUE, MQDS_ROUTINGLINK, or MQDS_MACHINE.<19>
	ObjectGUID *dtyp.GUID `idl:"name:pObjGuid;pointer:unique" json:"object_guid"`
}

func (o *CreateObjectRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateObjectOperation) *xxx_CreateObjectOperation {
	if op == nil {
		op = &xxx_CreateObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.PathName = o.PathName
	op.SecurityDescriptorLength = o.SecurityDescriptorLength
	op.SecurityDescriptor = o.SecurityDescriptor
	op.CreatePartition = o.CreatePartition
	op.Property = o.Property
	op.Var = o.Var
	op.ObjectGUID = o.ObjectGUID
	return op
}

func (o *CreateObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateObjectOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.PathName = op.PathName
	o.SecurityDescriptorLength = op.SecurityDescriptorLength
	o.SecurityDescriptor = op.SecurityDescriptor
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
	o.ObjectGUID = op.ObjectGUID
}
func (o *CreateObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateObjectResponse structure represents the S_DSCreateObject operation response
type CreateObjectResponse struct {
	// pObjGuid: SHOULD be set by the server to the GUID of the created object if the dwObjectType
	// is equal to MQDS_QUEUE, MQDS_ROUTINGLINK, or MQDS_MACHINE.<19>
	ObjectGUID *dtyp.GUID `idl:"name:pObjGuid;pointer:unique" json:"object_guid"`
	// Return: The S_DSCreateObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateObjectResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateObjectOperation) *xxx_CreateObjectOperation {
	if op == nil {
		op = &xxx_CreateObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectGUID = o.ObjectGUID
	op.Return = o.Return
	return op
}

func (o *CreateObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateObjectOperation) {
	if o == nil {
		return
	}
	o.ObjectGUID = op.ObjectGUID
	o.Return = op.Return
}
func (o *CreateObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteObjectOperation structure represents the S_DSDeleteObject operation
type xxx_DeleteObjectOperation struct {
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	PathName   string `idl:"name:pwcsPathName;string" json:"path_name"`
	Return     int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteObjectOperation) OpNum() int { return 1 }

func (o *xxx_DeleteObjectOperation) OpName() string { return "/dscomm/v1/S_DSDeleteObject" }

func (o *xxx_DeleteObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
	}
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
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
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

// DeleteObjectRequest structure represents the S_DSDeleteObject operation request
type DeleteObjectRequest struct {
	// dwObjectType:  Specifies the type of object to delete. MUST be set to one of the
	// directory object types specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pwcsPathName: Pointer to a NULL-terminated 16-bit Unicode string that MUST contain
	// the directory service pathname, as specified in section 2.2.9, to the object in the
	// directory service.
	PathName string `idl:"name:pwcsPathName;string" json:"path_name"`
}

func (o *DeleteObjectRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteObjectOperation) *xxx_DeleteObjectOperation {
	if op == nil {
		op = &xxx_DeleteObjectOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.PathName = o.PathName
	return op
}

func (o *DeleteObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteObjectOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.PathName = op.PathName
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

// DeleteObjectResponse structure represents the S_DSDeleteObject operation response
type DeleteObjectResponse struct {
	// Return: The S_DSDeleteObject return value.
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

// xxx_GetPropertiesOperation structure represents the S_DSGetProps operation
type xxx_GetPropertiesOperation struct {
	ObjectType          uint32                  `idl:"name:dwObjectType" json:"object_type"`
	PathName            string                  `idl:"name:pwcsPathName;string" json:"path_name"`
	CreatePartition     uint32                  `idl:"name:cp" json:"create_partition"`
	Property            []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var                 []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	ServerAuth          *ServerAuthType         `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte                  `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32                  `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 2 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/dscomm/v1/S_DSGetProps" }

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
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

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// GetPropertiesRequest structure represents the S_DSGetProps operation request
type GetPropertiesRequest struct {
	// dwObjectType:  Specifies the type of object for which properties are to be retrieved.
	// MUST be set to one of the object types specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pwcsPathName:  Pointer to a NULL-terminated 16-bit Unicode string that MUST contain
	// the directory service pathname, as specified in section 2.2.9, to the object in the
	// directory service.
	PathName string `idl:"name:pwcsPathName;string" json:"path_name"`
	// cp:  MUST be set to the size (in elements) of the arrays aProp and apVar. The arrays
	// aProp and apVar MUST have an identical number of elements and MUST contain at least
	// one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  An array of property identifiers specifying the set of object properties
	// to be returned. Each element MUST specify a value from the property identifiers table
	// defined in section 2.2.10.1 for the object type specified in dwObjectType. Each element
	// MUST specify the property identifier for the corresponding property value at the
	// same element index in apVar. The array MUST contain at least one element.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar: On input, each element MUST be initialized to the appropriate VARTYPE ([MS-MQMQ]
	// section 2.2.12.1) for the associated property specified by the same element in aProp,
	// or VT_NULL. On success, the server MUST populate the elements of this array with
	// property values for the properties identified by the corresponding elements of aProp.
	// The array MUST contain at least one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// phServerAuth:  A PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from
	// the pphServerAuth parameter in a previous call to S_DSValidateServer. The server
	// MUST use this parameter as a key to locate the GSS security context used to compute
	// the signature returned in pbServerSignature. See section 3.1.4.2.
	ServerAuth *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize: Contains the maximum length in bytes of the server signature
	// to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.PathName = o.PathName
	op.CreatePartition = o.CreatePartition
	op.Property = o.Property
	op.Var = o.Var
	op.ServerAuth = o.ServerAuth
	op.ServerSignatureSize = o.ServerSignatureSize
	return op
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.PathName = op.PathName
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the S_DSGetProps operation response
type GetPropertiesResponse struct {
	// XXX: cp is an implicit input depedency for output parameters
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// apVar: On input, each element MUST be initialized to the appropriate VARTYPE ([MS-MQMQ]
	// section 2.2.12.1) for the associated property specified by the same element in aProp,
	// or VT_NULL. On success, the server MUST populate the elements of this array with
	// property values for the properties identified by the corresponding elements of aProp.
	// The array MUST contain at least one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// pbServerSignature: Contains a signed hash over the returned property values.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize: Contains the maximum length in bytes of the server signature
	// to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSGetProps return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.CreatePartition == uint32(0) {
		op.CreatePartition = o.CreatePartition
	}

	op.Var = o.Var
	op.ServerSignature = o.ServerSignature
	op.ServerSignatureSize = o.ServerSignatureSize
	op.Return = o.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.CreatePartition = op.CreatePartition

	o.Var = op.Var
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPropertiesOperation structure represents the S_DSSetProps operation
type xxx_SetPropertiesOperation struct {
	ObjectType      uint32                  `idl:"name:dwObjectType" json:"object_type"`
	PathName        string                  `idl:"name:pwcsPathName;string" json:"path_name"`
	CreatePartition uint32                  `idl:"name:cp" json:"create_partition"`
	Property        []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var             []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	Return          int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPropertiesOperation) OpNum() int { return 3 }

func (o *xxx_SetPropertiesOperation) OpName() string { return "/dscomm/v1/S_DSSetProps" }

func (o *xxx_SetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
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

func (o *xxx_SetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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

func (o *xxx_SetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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

func (o *xxx_SetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetPropertiesRequest structure represents the S_DSSetProps operation request
type SetPropertiesRequest struct {
	// dwObjectType:  MUST specify the type of object for which properties are to be set.
	// For supported object types, see the table of object types specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pwcsPathName:  Pointer to a NULL-terminated 16-bit Unicode string that MUST contain
	// the directory service pathname, as specified in section 2.2.9, to the object in the
	// directory service.
	PathName string `idl:"name:pwcsPathName;string" json:"path_name"`
	// cp:  MUST be set to the size (in elements) of the arrays aProp and apVar. The arrays
	// aProp and apVar MUST have an identical number of elements and MUST contain at least
	// one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  An array of identifiers of properties to associate with the object identified
	// by pwcsPathName. Each element MUST specify a value from the property identifiers
	// table defined in section 2.2.10.1 for the object type specified in dwObjectType.
	// Each element MUST specify the property identifier for the corresponding property
	// value at the same element index in apVar. The array MUST contain at least one element.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar:  An array that specifies the property values to associate with the object.
	// Each element MUST specify the property value for the corresponding property identifier
	// at the same element index in aProp. The array MUST contain at least one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
}

func (o *SetPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPropertiesOperation) *xxx_SetPropertiesOperation {
	if op == nil {
		op = &xxx_SetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.PathName = o.PathName
	op.CreatePartition = o.CreatePartition
	op.Property = o.Property
	op.Var = o.Var
	return op
}

func (o *SetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPropertiesOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.PathName = op.PathName
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
}
func (o *SetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPropertiesResponse structure represents the S_DSSetProps operation response
type SetPropertiesResponse struct {
	// Return: The S_DSSetProps return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPropertiesOperation) *xxx_SetPropertiesOperation {
	if op == nil {
		op = &xxx_SetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPropertiesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetObjectSecurityOperation structure represents the S_DSGetObjectSecurity operation
type xxx_GetObjectSecurityOperation struct {
	ObjectType          uint32          `idl:"name:dwObjectType" json:"object_type"`
	PathName            string          `idl:"name:pwcsPathName;string" json:"path_name"`
	SecurityInformation uint32          `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  []byte          `idl:"name:pSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	Length              uint32          `idl:"name:nLength" json:"length"`
	LengthNeeded        uint32          `idl:"name:lpnLengthNeeded" json:"length_needed"`
	ServerAuth          *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte          `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32          `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetObjectSecurityOperation) OpNum() int { return 4 }

func (o *xxx_GetObjectSecurityOperation) OpName() string { return "/dscomm/v1/S_DSGetObjectSecurity" }

func (o *xxx_GetObjectSecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
	}
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

func (o *xxx_GetObjectSecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// nLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// nLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectSecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectSecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpnLengthNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.LengthNeeded); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_GetObjectSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpnLengthNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.LengthNeeded); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// GetObjectSecurityRequest structure represents the S_DSGetObjectSecurity operation request
type GetObjectSecurityRequest struct {
	// dwObjectType:  Specifies the type of object for which security properties are to
	// be retrieved. MUST be set to one of the object types defined in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pwcsPathName: Pointer to a NULL-terminated 16-bit Unicode string that MUST contain
	// the directory service pathname, as specified in section 2.2.9, of the object in the
	// directory service.
	PathName string `idl:"name:pwcsPathName;string" json:"path_name"`
	// SecurityInformation: A bitwise mask of the information to be returned in the pSecurityDescriptor
	// parameter. The bit fields are defined by the following table:
	//
	//	+---------------------------------------+-------------------------------------------------------+
	//	|                                       |                                                       |
	//	|                 VALUE                 |                        MEANING                        |
	//	|                                       |                                                       |
	//	+---------------------------------------+-------------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------------+
	//	| OWNER_SECURITY_INFORMATION 0x00000001 | OWNER field from the security descriptor.             |
	//	+---------------------------------------+-------------------------------------------------------+
	//	| GROUP_SECURITY_INFORMATION 0x00000002 | GROUP field from the security descriptor.             |
	//	+---------------------------------------+-------------------------------------------------------+
	//	| DACL_SECURITY_INFORMATION 0x00000004  | Discretionary ACL field from the security descriptor. |
	//	+---------------------------------------+-------------------------------------------------------+
	//	| SACL_SECURITY_INFORMATION 0x00000008  | System ACL field from the security descriptor.        |
	//	+---------------------------------------+-------------------------------------------------------+
	//	| MQDS_SIGN_PUBLIC_KEY 0x80000000       | Signing public key.                                   |
	//	+---------------------------------------+-------------------------------------------------------+
	//	| MQDS_KEYX_PUBLIC_KEY 0x40000000       | Encrypting public key.                                |
	//	+---------------------------------------+-------------------------------------------------------+
	//
	// The SecurityInformation parameter MUST specify one of:
	//
	// * MQDS_SIGN_PUBLIC_KEY, or
	//
	// * MQDS_KEYX_PUBLIC_KEY, or
	//
	// * A bitwise OR of any combination of:
	//
	// * OWNER_SECURITY_INFORMATION
	//
	// * GROUP_SECURITY_INFORMATION
	//
	// * DACL_SECURITY_INFORMATION
	//
	// * SACL_SECURITY_INFORMATION
	//
	// If the SecurityInformation parameter includes an invalid combination, the server
	// MUST NOT complete the call, and MUST return an error.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// nLength:  MUST be set by the client to the length in bytes of the pSecurityDescriptor
	// buffer.
	Length uint32 `idl:"name:nLength" json:"length"`
	// phServerAuth:  A PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from
	// the pphServerAuth parameter in a previous call to S_DSValidateServer. The server
	// MUST use this parameter as a key to locate the GSS security context used to compute
	// the signature returned in pbServerSignature. See section 3.1.4.2.
	ServerAuth *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize: A DWORD that contains the maximum length in bytes of the
	// server signature to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *GetObjectSecurityRequest) xxx_ToOp(ctx context.Context, op *xxx_GetObjectSecurityOperation) *xxx_GetObjectSecurityOperation {
	if op == nil {
		op = &xxx_GetObjectSecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.PathName = o.PathName
	op.SecurityInformation = o.SecurityInformation
	op.Length = o.Length
	op.ServerAuth = o.ServerAuth
	op.ServerSignatureSize = o.ServerSignatureSize
	return op
}

func (o *GetObjectSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetObjectSecurityOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.PathName = op.PathName
	o.SecurityInformation = op.SecurityInformation
	o.Length = op.Length
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *GetObjectSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetObjectSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetObjectSecurityResponse structure represents the S_DSGetObjectSecurity operation response
type GetObjectSecurityResponse struct {
	// XXX: nLength is an implicit input depedency for output parameters
	Length uint32 `idl:"name:nLength" json:"length"`
	// pSecurityDescriptor: If the SecurityInformation parameter is MQDS_SIGN_PUBLIC_KEY
	// or MQDS_KEYX_PUBLIC_KEY, it SHOULD<55> contain a pointer to a BLOBHEADER (section
	// 2.2.19) structure followed by an RSAPUBKEY (section 2.2.18) structure. Otherwise,
	// this parameter contains a security descriptor, as specified in [MS-DTYP] section
	// 2.4.6.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	// lpnLengthNeeded:  A DWORD representing the length in bytes of the requested security
	// descriptor or public key.
	LengthNeeded uint32 `idl:"name:lpnLengthNeeded" json:"length_needed"`
	// pbServerSignature: Contains a signed hash over the returned property values.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize: A DWORD that contains the maximum length in bytes of the
	// server signature to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSGetObjectSecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetObjectSecurityResponse) xxx_ToOp(ctx context.Context, op *xxx_GetObjectSecurityOperation) *xxx_GetObjectSecurityOperation {
	if op == nil {
		op = &xxx_GetObjectSecurityOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Length == uint32(0) {
		op.Length = o.Length
	}

	op.SecurityDescriptor = o.SecurityDescriptor
	op.LengthNeeded = o.LengthNeeded
	op.ServerSignature = o.ServerSignature
	op.ServerSignatureSize = o.ServerSignatureSize
	op.Return = o.Return
	return op
}

func (o *GetObjectSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetObjectSecurityOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Length = op.Length

	o.SecurityDescriptor = op.SecurityDescriptor
	o.LengthNeeded = op.LengthNeeded
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *GetObjectSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetObjectSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetObjectSecurityOperation structure represents the S_DSSetObjectSecurity operation
type xxx_SetObjectSecurityOperation struct {
	ObjectType          uint32 `idl:"name:dwObjectType" json:"object_type"`
	PathName            string `idl:"name:pwcsPathName;string" json:"path_name"`
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  []byte `idl:"name:pSecurityDescriptor;size_is:(nLength);pointer:unique" json:"security_descriptor"`
	Length              uint32 `idl:"name:nLength" json:"length"`
	Return              int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetObjectSecurityOperation) OpNum() int { return 5 }

func (o *xxx_SetObjectSecurityOperation) OpName() string { return "/dscomm/v1/S_DSSetObjectSecurity" }

func (o *xxx_SetObjectSecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.Length == 0 {
		o.Length = uint32(len(o.SecurityDescriptor))
	}
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
	}
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

func (o *xxx_SetObjectSecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// pSecurityDescriptor {in} (1:{pointer=unique}*(1)[dim:0,size_is=nLength](uchar))
	{
		if o.SecurityDescriptor != nil || o.Length > 0 {
			_ptr_pSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
	// nLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// pSecurityDescriptor {in} (1:{pointer=unique}*(1)[dim:0,size_is=nLength](uchar))
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
	// nLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectSecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectSecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetObjectSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetObjectSecurityRequest structure represents the S_DSSetObjectSecurity operation request
type SetObjectSecurityRequest struct {
	// dwObjectType: Specifies the type of object for which security properties are to be
	// set. MUST be one of the object types specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pwcsPathName: Pointer to a NULL-terminated 16-bit Unicode string that MUST contain
	// the directory service pathname, as specified in section 2.2.9, of the object in the
	// directory service.
	PathName string `idl:"name:pwcsPathName;string" json:"path_name"`
	// SecurityInformation: MUST be set by the client to a bitwise mask specifying the information
	// to set from the pSecurityDescriptor parameter. See the SecurityInformation parameter
	// description in section 3.1.4.11. Information in the pSecurityDescriptor parameter
	// not associated with bits set in this field MUST be ignored.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// pSecurityDescriptor: MUST contain a pointer to a security descriptor, as specified
	// in [MS-DTYP] section 2.4.6, or to an MQDS_PublicKey structure.<63> See the pSecurityDescriptor
	// parameter description in section 3.1.4.11. Note that where 3.1.4.11 indicates that
	// pSecurityDescriptor contains a BLOBHEADER followed by an RSAPUBKEY (section 2.2.18)
	// structure, this method actually contains an MQDS_PublicKey structure, which is the
	// same structure prefixed by a 4-byte length field.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(nLength);pointer:unique" json:"security_descriptor"`
	// nLength:  MUST be set by the client to the length in bytes of the pSecurityDescriptor
	// buffer.
	Length uint32 `idl:"name:nLength" json:"length"`
}

func (o *SetObjectSecurityRequest) xxx_ToOp(ctx context.Context, op *xxx_SetObjectSecurityOperation) *xxx_SetObjectSecurityOperation {
	if op == nil {
		op = &xxx_SetObjectSecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.PathName = o.PathName
	op.SecurityInformation = o.SecurityInformation
	op.SecurityDescriptor = o.SecurityDescriptor
	op.Length = o.Length
	return op
}

func (o *SetObjectSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_SetObjectSecurityOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.PathName = op.PathName
	o.SecurityInformation = op.SecurityInformation
	o.SecurityDescriptor = op.SecurityDescriptor
	o.Length = op.Length
}
func (o *SetObjectSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetObjectSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetObjectSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetObjectSecurityResponse structure represents the S_DSSetObjectSecurity operation response
type SetObjectSecurityResponse struct {
	// Return: The S_DSSetObjectSecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetObjectSecurityResponse) xxx_ToOp(ctx context.Context, op *xxx_SetObjectSecurityOperation) *xxx_SetObjectSecurityOperation {
	if op == nil {
		op = &xxx_SetObjectSecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetObjectSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_SetObjectSecurityOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetObjectSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetObjectSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetObjectSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupBeginOperation structure represents the S_DSLookupBegin operation
type xxx_LookupBeginOperation struct {
	Handle      *Type             `idl:"name:pHandle" json:"handle"`
	Context     string            `idl:"name:pwcsContext;pointer:unique" json:"context"`
	Restriction *mqds.Restriction `idl:"name:pRestriction;pointer:unique" json:"restriction"`
	Columns     *mqds.ColumnSet   `idl:"name:pColumns;pointer:ref" json:"columns"`
	Sort        *mqds.SortSet     `idl:"name:pSort;pointer:unique" json:"sort"`
	ServerAuth  *ServerAuthType   `idl:"name:phServerAuth" json:"server_auth"`
	Return      int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupBeginOperation) OpNum() int { return 6 }

func (o *xxx_LookupBeginOperation) OpName() string { return "/dscomm/v1/S_DSLookupBegin" }

func (o *xxx_LookupBeginOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupBeginOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pwcsContext {in} (1:{pointer=unique}*(1)[dim:0,string](wchar))
	{
		if o.Context != "" {
			_ptr_pwcsContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.Context); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pwcsContext); err != nil {
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
	// pRestriction {in} (1:{pointer=unique}*(1))(2:{alias=MQRESTRICTION}(struct))
	{
		if o.Restriction != nil {
			_ptr_pRestriction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Restriction != nil {
					if err := o.Restriction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqds.Restriction{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Restriction, _ptr_pRestriction); err != nil {
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
	// pColumns {in} (1:{pointer=ref}*(1))(2:{alias=MQCOLUMNSET}(struct))
	{
		if o.Columns != nil {
			if err := o.Columns.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqds.ColumnSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pSort {in} (1:{pointer=unique}*(1))(2:{alias=MQSORTSET}(struct))
	{
		if o.Sort != nil {
			_ptr_pSort := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Sort != nil {
					if err := o.Sort.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqds.SortSet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Sort, _ptr_pSort); err != nil {
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
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_LookupBeginOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pwcsContext {in} (1:{pointer=unique}*(1)[dim:0,string](wchar))
	{
		_ptr_pwcsContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.Context); err != nil {
				return err
			}
			return nil
		})
		_s_pwcsContext := func(ptr interface{}) { o.Context = *ptr.(*string) }
		if err := w.ReadPointer(&o.Context, _s_pwcsContext, _ptr_pwcsContext); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pRestriction {in} (1:{pointer=unique}*(1))(2:{alias=MQRESTRICTION}(struct))
	{
		_ptr_pRestriction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Restriction == nil {
				o.Restriction = &mqds.Restriction{}
			}
			if err := o.Restriction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pRestriction := func(ptr interface{}) { o.Restriction = *ptr.(**mqds.Restriction) }
		if err := w.ReadPointer(&o.Restriction, _s_pRestriction, _ptr_pRestriction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pColumns {in} (1:{pointer=ref}*(1))(2:{alias=MQCOLUMNSET}(struct))
	{
		if o.Columns == nil {
			o.Columns = &mqds.ColumnSet{}
		}
		if err := o.Columns.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pSort {in} (1:{pointer=unique}*(1))(2:{alias=MQSORTSET}(struct))
	{
		_ptr_pSort := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Sort == nil {
				o.Sort = &mqds.SortSet{}
			}
			if err := o.Sort.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pSort := func(ptr interface{}) { o.Sort = *ptr.(**mqds.SortSet) }
		if err := w.ReadPointer(&o.Sort, _s_pSort, _ptr_pSort); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupBeginOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupBeginOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pHandle {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_LookupBeginOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pHandle {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Type{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
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

// LookupBeginRequest structure represents the S_DSLookupBegin operation request
type LookupBeginRequest struct {
	// pwcsContext:  Unicode string that specifies a starting point of the query within
	// the directory service. The client SHOULD set this parameter to NULL, and the server
	// MUST ignore it.
	Context string `idl:"name:pwcsContext;pointer:unique" json:"context"`
	// pRestriction: A pointer to an MQRESTRICTION (section 2.2.12) structure specifying
	// a set of constraints over the objects to be returned. The server MUST restrict the
	// query results to include only objects that have properties that satisfy all of the
	// restrictions specified in this parameter. See section 2.2.12.
	Restriction *mqds.Restriction `idl:"name:pRestriction;pointer:unique" json:"restriction"`
	// pColumns: A pointer to an MQCOLUMNSET (section 2.2.13) structure that specifies the
	// object properties to be returned. The server MUST return (in the result set) only
	// the properties specified by this parameter in the order specified by this parameter.
	// See section 2.2.13.
	Columns *mqds.ColumnSet `idl:"name:pColumns;pointer:ref" json:"columns"`
	// pSort: A pointer to an MQSORTSET (section 2.2.15) structure that defines the sort
	// order of the result set. The server MUST sort the objects in the result set according
	// to this multikey sort order. See section 2.2.15.
	Sort *mqds.SortSet `idl:"name:pSort;pointer:unique" json:"sort"`
	// phServerAuth: A PCONTEXT_HANDLE_SERVER_AUTH_TYPE (section 2.2.5) RPC context handle
	// acquired from the pphServerAuth parameter in a previous call to S_DSValidateServer
	// (section 3.1.4.2).
	ServerAuth *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
}

func (o *LookupBeginRequest) xxx_ToOp(ctx context.Context, op *xxx_LookupBeginOperation) *xxx_LookupBeginOperation {
	if op == nil {
		op = &xxx_LookupBeginOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Restriction = o.Restriction
	op.Columns = o.Columns
	op.Sort = o.Sort
	op.ServerAuth = o.ServerAuth
	return op
}

func (o *LookupBeginRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupBeginOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Restriction = op.Restriction
	o.Columns = op.Columns
	o.Sort = op.Sort
	o.ServerAuth = op.ServerAuth
}
func (o *LookupBeginRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupBeginRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupBeginOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupBeginResponse structure represents the S_DSLookupBegin operation response
type LookupBeginResponse struct {
	// pHandle:  MUST be set by the server to point to an RPC context handle to be used
	// in subsequent calls to S_DSLookupNext and S_DSLookupEnd.
	Handle *Type `idl:"name:pHandle" json:"handle"`
	// Return: The S_DSLookupBegin return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupBeginResponse) xxx_ToOp(ctx context.Context, op *xxx_LookupBeginOperation) *xxx_LookupBeginOperation {
	if op == nil {
		op = &xxx_LookupBeginOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Return = o.Return
	return op
}

func (o *LookupBeginResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupBeginOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Return = op.Return
}
func (o *LookupBeginResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupBeginResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupBeginOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupNextOperation structure represents the S_DSLookupNext operation
type xxx_LookupNextOperation struct {
	Handle              *Type                   `idl:"name:Handle" json:"handle"`
	Size                uint32                  `idl:"name:dwSize" json:"size"`
	OutSize             uint32                  `idl:"name:dwOutSize" json:"out_size"`
	Buffer              []*mqmq.PropertyVariant `idl:"name:pbBuffer;size_is:(dwSize);length_is:(dwOutSize)" json:"buffer"`
	ServerAuth          *ServerAuthType         `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte                  `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32                  `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupNextOperation) OpNum() int { return 7 }

func (o *xxx_LookupNextOperation) OpName() string { return "/dscomm/v1/S_DSLookupNext" }

func (o *xxx_LookupNextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwSize {in} (1:{alias=LPBOUNDED_PROPERTIES}*(1))(2:{range=(0,128), alias=BOUNDED_PROPERTIES, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Type{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwSize {in} (1:{alias=LPBOUNDED_PROPERTIES,pointer=ref}*(1))(2:{range=(0,128), alias=BOUNDED_PROPERTIES, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.OutSize == 0 {
		o.OutSize = uint32(len(o.Buffer))
	}
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// dwOutSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.OutSize); err != nil {
			return err
		}
	}
	// pbBuffer {out} (1:[dim:0,size_is=dwSize,length_is=dwOutSize])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.Size)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.OutSize)
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
		for i1 := range o.Buffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Buffer[i1] != nil {
				if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_LookupNextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// dwOutSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.OutSize); err != nil {
			return err
		}
	}
	// pbBuffer {out} (1:[dim:0,size_is=dwSize,length_is=dwOutSize])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// LookupNextRequest structure represents the S_DSLookupNext operation request
type LookupNextRequest struct {
	// Handle: MUST contain an RPC context handle acquired from a previous call to S_DSLookupBegin.
	// The handle MUST NOT have been used in a previous call to S_DSLookupEnd (section 3.1.4.19).
	Handle *Type `idl:"name:Handle" json:"handle"`
	// dwSize:  MUST point to an unsigned long that contains the maximum number of elements
	// to be returned in the pbBuffer array. If this parameter is less than the number of
	// elements in the pColumns parameter in the corresponding call to S_DSLookupBegin,
	// the server MUST set the dwOutSize parameter to 0x00000000 and return without retrieving
	// any object properties.
	Size uint32 `idl:"name:dwSize" json:"size"`
	// phServerAuth: A PCONTEXT_HANDLE_SERVER_AUTH_TYPE (section 2.2.5) RPC context handle
	// acquired from the pphServerAuth parameter in a previous call to S_DSValidateServer
	// (section 3.1.4.2). The server MUST use this parameter as a key to locate the GSS
	// security context used to compute the signature returned in pbServerSignature. See
	// section 3.1.4.2.
	ServerAuth *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize:  MUST be set by the client to point to an unsigned LONG
	// that contains the maximum length in bytes of the server signature to return. MUST
	// be set by the server to the actual length in bytes of the server signature on output.
	// If the server signature is larger than the supplied buffer, the server MUST return
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028).
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *LookupNextRequest) xxx_ToOp(ctx context.Context, op *xxx_LookupNextOperation) *xxx_LookupNextOperation {
	if op == nil {
		op = &xxx_LookupNextOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Size = o.Size
	op.ServerAuth = o.ServerAuth
	op.ServerSignatureSize = o.ServerSignatureSize
	return op
}

func (o *LookupNextRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupNextOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Size = op.Size
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *LookupNextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupNextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupNextResponse structure represents the S_DSLookupNext operation response
type LookupNextResponse struct {
	// XXX: dwSize is an implicit input depedency for output parameters
	Size uint32 `idl:"name:dwSize" json:"size"`
	// dwOutSize: A pointer to an unsigned LONG that the server MUST set to the number of
	// properties returned in pbBuffer for the set of objects being returned from this invocation
	// of the S_DSLookupNext method. The server MUST return as many completed sets of properties
	// as will fit in the buffer. If no matching objects are found, the server MUST set
	// this parameter to 0 to inform the client that there is no more data.
	OutSize uint32 `idl:"name:dwOutSize" json:"out_size"`
	// pbBuffer: MUST point to an array of PROPVARIANT (section 2.2.3) structures to contain
	// the returned properties.
	Buffer []*mqmq.PropertyVariant `idl:"name:pbBuffer;size_is:(dwSize);length_is:(dwOutSize)" json:"buffer"`
	// pbServerSignature:  MUST point to the signed hash over the property values returned
	// in pbBuffer. See the pbServerSignature parameter description in section 3.1.4.7.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize:  MUST be set by the client to point to an unsigned LONG
	// that contains the maximum length in bytes of the server signature to return. MUST
	// be set by the server to the actual length in bytes of the server signature on output.
	// If the server signature is larger than the supplied buffer, the server MUST return
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028).
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSLookupNext return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupNextResponse) xxx_ToOp(ctx context.Context, op *xxx_LookupNextOperation) *xxx_LookupNextOperation {
	if op == nil {
		op = &xxx_LookupNextOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Size == uint32(0) {
		op.Size = o.Size
	}

	op.OutSize = o.OutSize
	op.Buffer = o.Buffer
	op.ServerSignature = o.ServerSignature
	op.ServerSignatureSize = o.ServerSignatureSize
	op.Return = o.Return
	return op
}

func (o *LookupNextResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupNextOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Size = op.Size

	o.OutSize = op.OutSize
	o.Buffer = op.Buffer
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *LookupNextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupNextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupEndOperation structure represents the S_DSLookupEnd operation
type xxx_LookupEndOperation struct {
	Context *Type `idl:"name:phContext" json:"context"`
	Return  int32 `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupEndOperation) OpNum() int { return 8 }

func (o *xxx_LookupEndOperation) OpName() string { return "/dscomm/v1/S_DSLookupEnd" }

func (o *xxx_LookupEndOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupEndOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_LookupEndOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Type{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupEndOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupEndOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phContext {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_LookupEndOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phContext {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Type{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
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

// LookupEndRequest structure represents the S_DSLookupEnd operation request
type LookupEndRequest struct {
	// phContext:  MUST point to an RPC context handle returned by a previous call to S_DSLookupBegin.
	// MUST NOT have been used in a previous call to S_DSLookupEnd (section 3.1.4.19). The
	// server MUST set this parameter to NULL.
	Context *Type `idl:"name:phContext" json:"context"`
}

func (o *LookupEndRequest) xxx_ToOp(ctx context.Context, op *xxx_LookupEndOperation) *xxx_LookupEndOperation {
	if op == nil {
		op = &xxx_LookupEndOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *LookupEndRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupEndOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *LookupEndRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupEndRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupEndOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupEndResponse structure represents the S_DSLookupEnd operation response
type LookupEndResponse struct {
	// phContext:  MUST point to an RPC context handle returned by a previous call to S_DSLookupBegin.
	// MUST NOT have been used in a previous call to S_DSLookupEnd (section 3.1.4.19). The
	// server MUST set this parameter to NULL.
	Context *Type `idl:"name:phContext" json:"context"`
	// Return: The S_DSLookupEnd return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupEndResponse) xxx_ToOp(ctx context.Context, op *xxx_LookupEndOperation) *xxx_LookupEndOperation {
	if op == nil {
		op = &xxx_LookupEndOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Return = o.Return
	return op
}

func (o *LookupEndResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupEndOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Return = op.Return
}
func (o *LookupEndResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupEndResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupEndOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteObjectGUIDOperation structure represents the S_DSDeleteObjectGuid operation
type xxx_DeleteObjectGUIDOperation struct {
	ObjectType uint32     `idl:"name:dwObjectType" json:"object_type"`
	GUID       *dtyp.GUID `idl:"name:pGuid" json:"guid"`
	Return     int32      `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteObjectGUIDOperation) OpNum() int { return 10 }

func (o *xxx_DeleteObjectGUIDOperation) OpName() string { return "/dscomm/v1/S_DSDeleteObjectGuid" }

func (o *xxx_DeleteObjectGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteObjectGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteObjectGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteObjectGUIDRequest structure represents the S_DSDeleteObjectGuid operation request
type DeleteObjectGUIDRequest struct {
	// dwObjectType:  Specifies the type of object to delete. MUST be set to one of the
	// object types specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pGuid:  A pointer to the GUID of the object to delete. This MUST be the identifier
	// assigned to the object by the server when the object was created. See section 3.1.4.1.
	GUID *dtyp.GUID `idl:"name:pGuid" json:"guid"`
}

func (o *DeleteObjectGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteObjectGUIDOperation) *xxx_DeleteObjectGUIDOperation {
	if op == nil {
		op = &xxx_DeleteObjectGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.GUID = o.GUID
	return op
}

func (o *DeleteObjectGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteObjectGUIDOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.GUID = op.GUID
}
func (o *DeleteObjectGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteObjectGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteObjectGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteObjectGUIDResponse structure represents the S_DSDeleteObjectGuid operation response
type DeleteObjectGUIDResponse struct {
	// Return: The S_DSDeleteObjectGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteObjectGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteObjectGUIDOperation) *xxx_DeleteObjectGUIDOperation {
	if op == nil {
		op = &xxx_DeleteObjectGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *DeleteObjectGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteObjectGUIDOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteObjectGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteObjectGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteObjectGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesGUIDOperation structure represents the S_DSGetPropsGuid operation
type xxx_GetPropertiesGUIDOperation struct {
	ObjectType          uint32                  `idl:"name:dwObjectType" json:"object_type"`
	GUID                *dtyp.GUID              `idl:"name:pGuid;pointer:unique" json:"guid"`
	CreatePartition     uint32                  `idl:"name:cp" json:"create_partition"`
	Property            []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var                 []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	ServerAuth          *ServerAuthType         `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte                  `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32                  `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesGUIDOperation) OpNum() int { return 11 }

func (o *xxx_GetPropertiesGUIDOperation) OpName() string { return "/dscomm/v1/S_DSGetPropsGuid" }

func (o *xxx_GetPropertiesGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
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

func (o *xxx_GetPropertiesGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			_ptr_pGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.GUID != nil {
					if err := o.GUID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.GUID, _ptr_pGuid); err != nil {
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
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_pGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.GUID == nil {
				o.GUID = &dtyp.GUID{}
			}
			if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pGuid := func(ptr interface{}) { o.GUID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.GUID, _s_pGuid, _ptr_pGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_GetPropertiesGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// GetPropertiesGUIDRequest structure represents the S_DSGetPropsGuid operation request
type GetPropertiesGUIDRequest struct {
	// dwObjectType:  Specifies the type of object for which properties are retrieved.
	// MUST be set to one of the object types specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pGuid:  MUST be set by the client to a pointer to the GUID of the object for which
	// properties are retrieved.
	GUID *dtyp.GUID `idl:"name:pGuid;pointer:unique" json:"guid"`
	// cp:  MUST be set to the size (in elements) of the arrays aProp and apVar. The arrays
	// aProp and apVar MUST have an identical number of elements and MUST contain at least
	// one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  An array of identifiers of properties to retrieve from the object designated
	// by pGuid. Each element MUST specify a value from the property identifiers table (defined
	// in section 2.2.10.1) for the object type specified in dwObjectType. Each element
	// MUST specify the property identifier for the corresponding property value at the
	// same element index in apVar. The array MUST contain at least one element.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar: On input, each element MUST be initialized to the appropriate VARTYPE ([MS-MQMQ]
	// section 2.2.12.1) for the associated property specified by the same element in aProp,
	// or VT_NULL. On success, the server MUST populate the elements of this array with
	// property values for the properties identified by the corresponding elements of aProp.
	// The array MUST contain at least one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// phServerAuth:  A PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from
	// the pphServerAuth parameter in a previous call to S_DSValidateServer. The server
	// MUST use this parameter as a key to locate the GSS security context used to compute
	// the signature returned in the pbServerSignature parameter. See section 3.1.4.2.
	ServerAuth *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize:  Contains the maximum length in bytes of the server signature
	// to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *GetPropertiesGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesGUIDOperation) *xxx_GetPropertiesGUIDOperation {
	if op == nil {
		op = &xxx_GetPropertiesGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.GUID = o.GUID
	op.CreatePartition = o.CreatePartition
	op.Property = o.Property
	op.Var = o.Var
	op.ServerAuth = o.ServerAuth
	op.ServerSignatureSize = o.ServerSignatureSize
	return op
}

func (o *GetPropertiesGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesGUIDOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.GUID = op.GUID
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *GetPropertiesGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesGUIDResponse structure represents the S_DSGetPropsGuid operation response
type GetPropertiesGUIDResponse struct {
	// XXX: cp is an implicit input depedency for output parameters
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// apVar: On input, each element MUST be initialized to the appropriate VARTYPE ([MS-MQMQ]
	// section 2.2.12.1) for the associated property specified by the same element in aProp,
	// or VT_NULL. On success, the server MUST populate the elements of this array with
	// property values for the properties identified by the corresponding elements of aProp.
	// The array MUST contain at least one element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// pbServerSignature:  See the pbServerSignature parameter description, as specified
	// in section 3.1.4.7.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize:  Contains the maximum length in bytes of the server signature
	// to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSGetPropsGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesGUIDOperation) *xxx_GetPropertiesGUIDOperation {
	if op == nil {
		op = &xxx_GetPropertiesGUIDOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.CreatePartition == uint32(0) {
		op.CreatePartition = o.CreatePartition
	}

	op.Var = o.Var
	op.ServerSignature = o.ServerSignature
	op.ServerSignatureSize = o.ServerSignatureSize
	op.Return = o.Return
	return op
}

func (o *GetPropertiesGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesGUIDOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.CreatePartition = op.CreatePartition

	o.Var = op.Var
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *GetPropertiesGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPropertiesGUIDOperation structure represents the S_DSSetPropsGuid operation
type xxx_SetPropertiesGUIDOperation struct {
	ObjectType      uint32                  `idl:"name:dwObjectType" json:"object_type"`
	GUID            *dtyp.GUID              `idl:"name:pGuid" json:"guid"`
	CreatePartition uint32                  `idl:"name:cp" json:"create_partition"`
	Property        []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var             []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	Return          int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPropertiesGUIDOperation) OpNum() int { return 12 }

func (o *xxx_SetPropertiesGUIDOperation) OpName() string { return "/dscomm/v1/S_DSSetPropsGuid" }

func (o *xxx_SetPropertiesGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
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

func (o *xxx_SetPropertiesGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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

func (o *xxx_SetPropertiesGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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

func (o *xxx_SetPropertiesGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPropertiesGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPropertiesGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetPropertiesGUIDRequest structure represents the S_DSSetPropsGuid operation request
type SetPropertiesGUIDRequest struct {
	// dwObjectType:  Specifies the type of object for which properties are to be set.
	// MUST be one of the object types specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pGuid:  Pointer to the GUID of the object for which properties are to be set.
	GUID *dtyp.GUID `idl:"name:pGuid" json:"guid"`
	// cp:  MUST be set to the size (in elements) of the arrays aProp and apVar. The arrays
	// aProp and apVar MUST have an identical number of elements and MUST contain at least
	// one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  An array of identifiers of properties to associate with the object designated
	// by pGuid. Each element MUST specify a value from the property identifiers table,
	// as specified in section 2.2.10.1, for the object type specified in dwObjectType.
	// Each element MUST specify the property identifier for the corresponding property
	// value at the same element index in apVar. The array MUST contain at least one element.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar:  An array that specifies the property values to associate with the created
	// object. Each element MUST specify the property value for the corresponding property
	// identifier at the same element index in aProp. The array MUST contain at least one
	// element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
}

func (o *SetPropertiesGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPropertiesGUIDOperation) *xxx_SetPropertiesGUIDOperation {
	if op == nil {
		op = &xxx_SetPropertiesGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.GUID = o.GUID
	op.CreatePartition = o.CreatePartition
	op.Property = o.Property
	op.Var = o.Var
	return op
}

func (o *SetPropertiesGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPropertiesGUIDOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.GUID = op.GUID
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
}
func (o *SetPropertiesGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPropertiesGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPropertiesGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPropertiesGUIDResponse structure represents the S_DSSetPropsGuid operation response
type SetPropertiesGUIDResponse struct {
	// Return: The S_DSSetPropsGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPropertiesGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPropertiesGUIDOperation) *xxx_SetPropertiesGUIDOperation {
	if op == nil {
		op = &xxx_SetPropertiesGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetPropertiesGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPropertiesGUIDOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetPropertiesGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPropertiesGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPropertiesGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetObjectSecurityGUIDOperation structure represents the S_DSGetObjectSecurityGuid operation
type xxx_GetObjectSecurityGUIDOperation struct {
	ObjectType          uint32          `idl:"name:dwObjectType" json:"object_type"`
	GUID                *dtyp.GUID      `idl:"name:pGuid" json:"guid"`
	SecurityInformation uint32          `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  []byte          `idl:"name:pSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	Length              uint32          `idl:"name:nLength" json:"length"`
	LengthNeeded        uint32          `idl:"name:lpnLengthNeeded" json:"length_needed"`
	ServerAuth          *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte          `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32          `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetObjectSecurityGUIDOperation) OpNum() int { return 13 }

func (o *xxx_GetObjectSecurityGUIDOperation) OpName() string {
	return "/dscomm/v1/S_DSGetObjectSecurityGuid"
}

func (o *xxx_GetObjectSecurityGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
	}
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

func (o *xxx_GetObjectSecurityGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecurityInformation {in} (1:(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// nLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectSecurityGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// nLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectSecurityGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetObjectSecurityGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpnLengthNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.LengthNeeded); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_GetObjectSecurityGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpnLengthNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.LengthNeeded); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// GetObjectSecurityGUIDRequest structure represents the S_DSGetObjectSecurityGuid operation request
type GetObjectSecurityGUIDRequest struct {
	// dwObjectType:  Specifies the type of object for which security properties are to
	// be retrieved. MUST be set to one of the object types specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pGuid:  MUST be set by the client to a pointer to the GUID of the object for which
	// to retrieve security information.
	GUID *dtyp.GUID `idl:"name:pGuid" json:"guid"`
	// SecurityInformation:  MUST be set by the client to a bitwise mask specifying the
	// information to return in the pSecurityDescriptor parameter. See the SecurityInformation
	// parameter description in section 3.1.4.11.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// nLength: MUST be set by the client to the length in bytes of the pSecurityDescriptor
	// buffer.
	Length uint32 `idl:"name:nLength" json:"length"`
	// phServerAuth: A PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from
	// the pphServerAuth parameter in a previous call to S_DSValidateServer. The server
	// MUST use this parameter as a key to locate the GSS security context used to compute
	// the signature returned in pbServerSignature. See section 3.1.4.2.
	ServerAuth *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize: A DWORD that contains the maximum length in bytes of the
	// server signature to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *GetObjectSecurityGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetObjectSecurityGUIDOperation) *xxx_GetObjectSecurityGUIDOperation {
	if op == nil {
		op = &xxx_GetObjectSecurityGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.GUID = o.GUID
	op.SecurityInformation = o.SecurityInformation
	op.Length = o.Length
	op.ServerAuth = o.ServerAuth
	op.ServerSignatureSize = o.ServerSignatureSize
	return op
}

func (o *GetObjectSecurityGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetObjectSecurityGUIDOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.GUID = op.GUID
	o.SecurityInformation = op.SecurityInformation
	o.Length = op.Length
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *GetObjectSecurityGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetObjectSecurityGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectSecurityGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetObjectSecurityGUIDResponse structure represents the S_DSGetObjectSecurityGuid operation response
type GetObjectSecurityGUIDResponse struct {
	// XXX: nLength is an implicit input depedency for output parameters
	Length uint32 `idl:"name:nLength" json:"length"`
	// pSecurityDescriptor: If the SecurityInformation parameter is MQDS_SIGN_PUBLIC_KEY
	// or MQDS_KEYX_PUBLIC_KEY, it SHOULD<59> contain a pointer to a BLOBHEADER structure
	// followed by an RSAPUBKEY (section 2.2.18) structure. Otherwise, this parameter contains
	// a pointer to a security descriptor, as specified in [MS-DTYP] section 2.4.6.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	// lpnLengthNeeded: A DWORD that the server MUST set to the length in bytes of the requested
	// security descriptor or public key. If the requested security descriptor or public
	// key is larger than nLength, the server MUST set this parameter to the size in bytes
	// needed for the requested security descriptor or public key, and return MQ_ERROR_SECURITY_DESCRIPTOR_TOO_SMALL
	// (0xC00E0023).
	LengthNeeded uint32 `idl:"name:lpnLengthNeeded" json:"length_needed"`
	// pbServerSignature: MUST point to the signed hash of the security descriptor.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize: A DWORD that contains the maximum length in bytes of the
	// server signature to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSGetObjectSecurityGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetObjectSecurityGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetObjectSecurityGUIDOperation) *xxx_GetObjectSecurityGUIDOperation {
	if op == nil {
		op = &xxx_GetObjectSecurityGUIDOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Length == uint32(0) {
		op.Length = o.Length
	}

	op.SecurityDescriptor = o.SecurityDescriptor
	op.LengthNeeded = o.LengthNeeded
	op.ServerSignature = o.ServerSignature
	op.ServerSignatureSize = o.ServerSignatureSize
	op.Return = o.Return
	return op
}

func (o *GetObjectSecurityGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetObjectSecurityGUIDOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Length = op.Length

	o.SecurityDescriptor = op.SecurityDescriptor
	o.LengthNeeded = op.LengthNeeded
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *GetObjectSecurityGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetObjectSecurityGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetObjectSecurityGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetObjectSecurityGUIDOperation structure represents the S_DSSetObjectSecurityGuid operation
type xxx_SetObjectSecurityGUIDOperation struct {
	ObjectType          uint32     `idl:"name:dwObjectType" json:"object_type"`
	GUID                *dtyp.GUID `idl:"name:pGuid" json:"guid"`
	SecurityInformation uint32     `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  []byte     `idl:"name:pSecurityDescriptor;size_is:(nLength);pointer:unique" json:"security_descriptor"`
	Length              uint32     `idl:"name:nLength" json:"length"`
	Return              int32      `idl:"name:Return" json:"return"`
}

func (o *xxx_SetObjectSecurityGUIDOperation) OpNum() int { return 14 }

func (o *xxx_SetObjectSecurityGUIDOperation) OpName() string {
	return "/dscomm/v1/S_DSSetObjectSecurityGuid"
}

func (o *xxx_SetObjectSecurityGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.Length == 0 {
		o.Length = uint32(len(o.SecurityDescriptor))
	}
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
	}
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

func (o *xxx_SetObjectSecurityGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecurityInformation {in} (1:(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// pSecurityDescriptor {in} (1:{pointer=unique}*(1)[dim:0,size_is=nLength](uchar))
	{
		if o.SecurityDescriptor != nil || o.Length > 0 {
			_ptr_pSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
	// nLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectSecurityGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// pSecurityDescriptor {in} (1:{pointer=unique}*(1)[dim:0,size_is=nLength](uchar))
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
	// nLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectSecurityGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetObjectSecurityGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetObjectSecurityGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetObjectSecurityGUIDRequest structure represents the S_DSSetObjectSecurityGuid operation request
type SetObjectSecurityGUIDRequest struct {
	// dwObjectType:  Specifies the type of object for which security properties are set.
	// MUST be set to one of the object types specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pGuid: MUST be set by the client to a pointer to the GUID of the object for which
	// properties are to be set.
	GUID *dtyp.GUID `idl:"name:pGuid" json:"guid"`
	// SecurityInformation: MUST be set by the client to a bitwise mask specifying the information
	// to set from the pSecurityDescriptor parameter. See the SecurityInformation parameter
	// description in section 3.1.4.11.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// pSecurityDescriptor: MUST contain a pointer to security descriptor, as specified
	// in [MS-DTYP] section 2.4.6, or to an MQDS_PublicKey structure.<68> See the pSecurityDescriptor
	// parameter description in section 3.1.4.11. Note that while section 3.1.4.11 indicates
	// that pSecurityDescriptor contains a BLOBHEADER followed by an RSAPUBKEY (section
	// 2.2.18) structure, this method actually contains an MQDS_PublicKey structure, which
	// is the same structure prefixed by a 4-byte length field.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(nLength);pointer:unique" json:"security_descriptor"`
	// nLength: MUST be set by the client to the length in bytes of the pSecurityDescriptor
	// buffer.
	Length uint32 `idl:"name:nLength" json:"length"`
}

func (o *SetObjectSecurityGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetObjectSecurityGUIDOperation) *xxx_SetObjectSecurityGUIDOperation {
	if op == nil {
		op = &xxx_SetObjectSecurityGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.GUID = o.GUID
	op.SecurityInformation = o.SecurityInformation
	op.SecurityDescriptor = o.SecurityDescriptor
	op.Length = o.Length
	return op
}

func (o *SetObjectSecurityGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetObjectSecurityGUIDOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.GUID = op.GUID
	o.SecurityInformation = op.SecurityInformation
	o.SecurityDescriptor = op.SecurityDescriptor
	o.Length = op.Length
}
func (o *SetObjectSecurityGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetObjectSecurityGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetObjectSecurityGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetObjectSecurityGUIDResponse structure represents the S_DSSetObjectSecurityGuid operation response
type SetObjectSecurityGUIDResponse struct {
	// Return: The S_DSSetObjectSecurityGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetObjectSecurityGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetObjectSecurityGUIDOperation) *xxx_SetObjectSecurityGUIDOperation {
	if op == nil {
		op = &xxx_SetObjectSecurityGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetObjectSecurityGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetObjectSecurityGUIDOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetObjectSecurityGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetObjectSecurityGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetObjectSecurityGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QMSetMachinePropertiesOperation structure represents the S_DSQMSetMachineProperties operation
type xxx_QMSetMachinePropertiesOperation struct {
	PathName        string                  `idl:"name:pwcsPathName;string" json:"path_name"`
	CreatePartition uint32                  `idl:"name:cp" json:"create_partition"`
	Property        []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var             []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	Context         uint32                  `idl:"name:dwContext" json:"context"`
	Return          int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_QMSetMachinePropertiesOperation) OpNum() int { return 19 }

func (o *xxx_QMSetMachinePropertiesOperation) OpName() string {
	return "/dscomm/v1/S_DSQMSetMachineProperties"
}

func (o *xxx_QMSetMachinePropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
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

func (o *xxx_QMSetMachinePropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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
	// dwContext {in} (1:(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMSetMachinePropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pwcsPathName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128)}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp](uint32))
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
	// dwContext {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMSetMachinePropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMSetMachinePropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QMSetMachinePropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QMSetMachinePropertiesRequest structure represents the S_DSQMSetMachineProperties operation request
type QMSetMachinePropertiesRequest struct {
	// pwcsPathName: Pointer to a NULL-terminated 16-bit Unicode string that MUST contain
	// the directory service pathname, as specified in section 2.2.9, of a machine object
	// in the directory service.
	PathName string `idl:"name:pwcsPathName;string" json:"path_name"`
	// cp: MUST be set to the size (in elements) of the arrays aProp and apVar. The arrays
	// aProp and apVar MUST have an identical number of elements, and MUST contain at least
	// one element.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp:  An array of identifiers of properties to associate with the machine object.
	// Each element MUST specify a value from the property identifiers table, as specified
	// in section 2.2.10.1. Each element MUST specify the property identifier for the corresponding
	// property value at the same element index in apVar. The array MUST contain at least
	// one element.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar:  An array that specifies the property values to associate with the machine
	// object. Each element MUST specify the property value for the corresponding property
	// identifier at the same element index in aProp. The array MUST contain at least one
	// element.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// dwContext:  MUST be set by the client to a value that the client can use to correlate
	// callbacks with the initial call to S_DSQMSetMachineProperties. The server MUST supply
	// this value in the dwContext parameter in the related calls of the S_DSQMSetMachinePropertiesSignProc
	// (section 3.2.4.1) callback method.
	Context uint32 `idl:"name:dwContext" json:"context"`
}

func (o *QMSetMachinePropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_QMSetMachinePropertiesOperation) *xxx_QMSetMachinePropertiesOperation {
	if op == nil {
		op = &xxx_QMSetMachinePropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.PathName = o.PathName
	op.CreatePartition = o.CreatePartition
	op.Property = o.Property
	op.Var = o.Var
	op.Context = o.Context
	return op
}

func (o *QMSetMachinePropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_QMSetMachinePropertiesOperation) {
	if o == nil {
		return
	}
	o.PathName = op.PathName
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
	o.Context = op.Context
}
func (o *QMSetMachinePropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QMSetMachinePropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QMSetMachinePropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QMSetMachinePropertiesResponse structure represents the S_DSQMSetMachineProperties operation response
type QMSetMachinePropertiesResponse struct {
	// Return: The S_DSQMSetMachineProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QMSetMachinePropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_QMSetMachinePropertiesOperation) *xxx_QMSetMachinePropertiesOperation {
	if op == nil {
		op = &xxx_QMSetMachinePropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *QMSetMachinePropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_QMSetMachinePropertiesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *QMSetMachinePropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QMSetMachinePropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QMSetMachinePropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateServersCacheOperation structure represents the S_DSCreateServersCache operation
type xxx_CreateServersCacheOperation struct {
	Index               uint32          `idl:"name:pdwIndex" json:"index"`
	SiteServers         string          `idl:"name:lplpSiteServers;string;pointer:ptr" json:"site_servers"`
	ServerAuth          *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte          `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32          `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateServersCacheOperation) OpNum() int { return 20 }

func (o *xxx_CreateServersCacheOperation) OpName() string { return "/dscomm/v1/S_DSCreateServersCache" }

func (o *xxx_CreateServersCacheOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateServersCacheOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pdwIndex {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// lplpSiteServers {in, out} (1:{string, pointer=ptr}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.SiteServers != "" {
			_ptr_lplpSiteServers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SiteServers != "" {
					_ptr_lplpSiteServers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.SiteServers); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.SiteServers, _ptr_lplpSiteServers); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SiteServers, _ptr_lplpSiteServers); err != nil {
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
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateServersCacheOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pdwIndex {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// lplpSiteServers {in, out} (1:{string, pointer=ptr}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_lplpSiteServers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_lplpSiteServers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.SiteServers); err != nil {
					return err
				}
				return nil
			})
			_s_lplpSiteServers := func(ptr interface{}) { o.SiteServers = *ptr.(*string) }
			if err := w.ReadPointer(&o.SiteServers, _s_lplpSiteServers, _ptr_lplpSiteServers); err != nil {
				return err
			}
			return nil
		})
		_s_lplpSiteServers := func(ptr interface{}) { o.SiteServers = *ptr.(*string) }
		if err := w.ReadPointer(&o.SiteServers, _s_lplpSiteServers, _ptr_lplpSiteServers); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateServersCacheOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateServersCacheOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwIndex {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// lplpSiteServers {in, out} (1:{string, pointer=ptr}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.SiteServers != "" {
			_ptr_lplpSiteServers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SiteServers != "" {
					_ptr_lplpSiteServers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.SiteServers); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.SiteServers, _ptr_lplpSiteServers); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SiteServers, _ptr_lplpSiteServers); err != nil {
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
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_CreateServersCacheOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwIndex {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// lplpSiteServers {in, out} (1:{string, pointer=ptr}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_lplpSiteServers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_lplpSiteServers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.SiteServers); err != nil {
					return err
				}
				return nil
			})
			_s_lplpSiteServers := func(ptr interface{}) { o.SiteServers = *ptr.(*string) }
			if err := w.ReadPointer(&o.SiteServers, _s_lplpSiteServers, _ptr_lplpSiteServers); err != nil {
				return err
			}
			return nil
		})
		_s_lplpSiteServers := func(ptr interface{}) { o.SiteServers = *ptr.(*string) }
		if err := w.ReadPointer(&o.SiteServers, _s_lplpSiteServers, _ptr_lplpSiteServers); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// CreateServersCacheRequest structure represents the S_DSCreateServersCache operation request
type CreateServersCacheRequest struct {
	// pdwIndex:  Pointer to an unsigned long that contains an index into the configured
	// list of sites in the enterprise indicating the site to which the list of BSCs is
	// to be returned.
	Index uint32 `idl:"name:pdwIndex" json:"index"`
	// lplpSiteServers:  Pointer to a pointer to a string that contains the list of servers
	// associated with the indexed site. The string MUST be in Server List String (section
	// 2.2.17) format.
	SiteServers string `idl:"name:lplpSiteServers;string;pointer:ptr" json:"site_servers"`
	// phServerAuth: PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from the
	// pphServerAuth parameter in a previous call to S_DSValidateServer. The server MUST
	// use this parameter as a key to locate the GSS security context used to compute the
	// signature returned in pbServerSignature. See section 3.1.4.2.
	ServerAuth *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize:  MUST be set by the client to point to an unsigned LONG
	// that contains the maximum length in bytes of the server signature to return. MUST
	// be set by the server to the actual length in bytes of the server signature on output.
	// If the server signature is larger than the supplied buffer, the server MUST return
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028).
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *CreateServersCacheRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateServersCacheOperation) *xxx_CreateServersCacheOperation {
	if op == nil {
		op = &xxx_CreateServersCacheOperation{}
	}
	if o == nil {
		return op
	}
	op.Index = o.Index
	op.SiteServers = o.SiteServers
	op.ServerAuth = o.ServerAuth
	op.ServerSignatureSize = o.ServerSignatureSize
	return op
}

func (o *CreateServersCacheRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateServersCacheOperation) {
	if o == nil {
		return
	}
	o.Index = op.Index
	o.SiteServers = op.SiteServers
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *CreateServersCacheRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateServersCacheRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateServersCacheOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateServersCacheResponse structure represents the S_DSCreateServersCache operation response
type CreateServersCacheResponse struct {
	// pdwIndex:  Pointer to an unsigned long that contains an index into the configured
	// list of sites in the enterprise indicating the site to which the list of BSCs is
	// to be returned.
	Index uint32 `idl:"name:pdwIndex" json:"index"`
	// lplpSiteServers:  Pointer to a pointer to a string that contains the list of servers
	// associated with the indexed site. The string MUST be in Server List String (section
	// 2.2.17) format.
	SiteServers string `idl:"name:lplpSiteServers;string;pointer:ptr" json:"site_servers"`
	// pbServerSignature:  MUST be set by the server to a buffer that contains a signed
	// hash over the server list returned in lplpSiteServers. The server MUST construct
	// this signature by creating a hash by using the MD5 algorithm (as specified in [RFC1321])
	// and sealing it, as specified by the following pseudocode.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize:  MUST be set by the client to point to an unsigned LONG
	// that contains the maximum length in bytes of the server signature to return. MUST
	// be set by the server to the actual length in bytes of the server signature on output.
	// If the server signature is larger than the supplied buffer, the server MUST return
	// MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028).
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSCreateServersCache return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateServersCacheResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateServersCacheOperation) *xxx_CreateServersCacheOperation {
	if op == nil {
		op = &xxx_CreateServersCacheOperation{}
	}
	if o == nil {
		return op
	}
	op.Index = o.Index
	op.SiteServers = o.SiteServers
	op.ServerSignature = o.ServerSignature
	op.ServerSignatureSize = o.ServerSignatureSize
	op.Return = o.Return
	return op
}

func (o *CreateServersCacheResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateServersCacheOperation) {
	if o == nil {
		return
	}
	o.Index = op.Index
	o.SiteServers = op.SiteServers
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *CreateServersCacheResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateServersCacheResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateServersCacheOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QMSetMachinePropertiesSignProcOperation structure represents the S_DSQMSetMachinePropertiesSignProc operation
type xxx_QMSetMachinePropertiesSignProcOperation struct {
	Challenge        []byte `idl:"name:abChallenge;size_is:(dwCallengeSize)" json:"challenge"`
	ChallengeSize    uint32 `idl:"name:dwCallengeSize" json:"challenge_size"`
	Context          uint32 `idl:"name:dwContext" json:"context"`
	Signature        []byte `idl:"name:abSignature;size_is:(dwSignatureMaxSize);length_is:(pdwSignatureSize)" json:"signature"`
	SignatureSize    uint32 `idl:"name:pdwSignatureSize" json:"signature_size"`
	SignatureMaxSize uint32 `idl:"name:dwSignatureMaxSize" json:"signature_max_size"`
	Return           int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_QMSetMachinePropertiesSignProcOperation) OpNum() int { return 21 }

func (o *xxx_QMSetMachinePropertiesSignProcOperation) OpName() string {
	return "/dscomm/v1/S_DSQMSetMachinePropertiesSignProc"
}

func (o *xxx_QMSetMachinePropertiesSignProcOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Challenge != nil && o.ChallengeSize == 0 {
		o.ChallengeSize = uint32(len(o.Challenge))
	}
	if o.Signature != nil && o.SignatureMaxSize == 0 {
		o.SignatureMaxSize = uint32(len(o.Signature))
	}
	if o.Signature != nil && o.SignatureSize == 0 {
		o.SignatureSize = uint32(len(o.Signature))
	}
	if o.ChallengeSize > uint32(32) {
		return fmt.Errorf("ChallengeSize is out of range")
	}
	if o.SignatureMaxSize > uint32(128) {
		return fmt.Errorf("SignatureMaxSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMSetMachinePropertiesSignProcOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// abChallenge {in} (1:{pointer=ref}*(1)[dim:0,size_is=dwCallengeSize](uint8))
	{
		dimSize1 := uint64(o.ChallengeSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Challenge {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Challenge[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Challenge); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// dwCallengeSize {in} (1:{range=(0,32)}(uint32))
	{
		if err := w.WriteData(o.ChallengeSize); err != nil {
			return err
		}
	}
	// dwContext {in} (1:(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	// abSignature {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwSignatureMaxSize,length_is=pdwSignatureSize](uint8))
	{
		dimSize1 := uint64(o.SignatureMaxSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.SignatureSize)
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
	}
	// pdwSignatureSize {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.SignatureSize); err != nil {
			return err
		}
	}
	// dwSignatureMaxSize {in} (1:{range=(0,128)}(uint32))
	{
		if err := w.WriteData(o.SignatureMaxSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMSetMachinePropertiesSignProcOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// abChallenge {in} (1:{pointer=ref}*(1)[dim:0,size_is=dwCallengeSize](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Challenge", sizeInfo[0])
		}
		o.Challenge = make([]byte, sizeInfo[0])
		for i1 := range o.Challenge {
			i1 := i1
			if err := w.ReadData(&o.Challenge[i1]); err != nil {
				return err
			}
		}
	}
	// dwCallengeSize {in} (1:{range=(0,32)}(uint32))
	{
		if err := w.ReadData(&o.ChallengeSize); err != nil {
			return err
		}
	}
	// dwContext {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	// abSignature {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwSignatureMaxSize,length_is=pdwSignatureSize](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Signature", sizeInfo[0])
		}
		o.Signature = make([]byte, sizeInfo[0])
		for i1 := range o.Signature {
			i1 := i1
			if err := w.ReadData(&o.Signature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwSignatureSize {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.SignatureSize); err != nil {
			return err
		}
	}
	// dwSignatureMaxSize {in} (1:{range=(0,128)}(uint32))
	{
		if err := w.ReadData(&o.SignatureMaxSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMSetMachinePropertiesSignProcOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Signature != nil && o.SignatureSize == 0 {
		o.SignatureSize = uint32(len(o.Signature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMSetMachinePropertiesSignProcOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// abSignature {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwSignatureMaxSize,length_is=pdwSignatureSize](uint8))
	{
		dimSize1 := uint64(o.SignatureMaxSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.SignatureSize)
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
	}
	// pdwSignatureSize {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.SignatureSize); err != nil {
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

func (o *xxx_QMSetMachinePropertiesSignProcOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// abSignature {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwSignatureMaxSize,length_is=pdwSignatureSize](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Signature", sizeInfo[0])
		}
		o.Signature = make([]byte, sizeInfo[0])
		for i1 := range o.Signature {
			i1 := i1
			if err := w.ReadData(&o.Signature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwSignatureSize {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.SignatureSize); err != nil {
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

// QMSetMachinePropertiesSignProcRequest structure represents the S_DSQMSetMachinePropertiesSignProc operation request
type QMSetMachinePropertiesSignProcRequest struct {
	// abChallenge:  MUST be set by the caller to a pointer to a buffer that contains the
	// challenge to be signed. The challenge SHOULD be cryptographically random.
	Challenge []byte `idl:"name:abChallenge;size_is:(dwCallengeSize)" json:"challenge"`
	// dwCallengeSize:  MUST be set by the caller to the size in bytes of the challenge
	// in the abChallenge parameter.
	ChallengeSize uint32 `idl:"name:dwCallengeSize" json:"challenge_size"`
	// dwContext:  MUST be set by the caller to the value supplied by the client in the
	// dwContext parameter of the corresponding call to the S_DSQMSetMachineProperties method.
	// This parameter provides a way for the receiver to correlate the callback with the
	// receiver's in-progress call to S_DSQMSetMachineProperties.
	Context uint32 `idl:"name:dwContext" json:"context"`
	// abSignature: MUST be set by the caller to a pointer to a buffer to contain the returned
	// signature. MUST be set by the receiver to a signature over the challenge in abChallenge.
	// The algorithm for creating this signature is specified by the following pseudocode.
	Signature []byte `idl:"name:abSignature;size_is:(dwSignatureMaxSize);length_is:(pdwSignatureSize)" json:"signature"`
	// pdwSignatureSize:  Size in bytes of the signature in the abSignature parameter.
	// MUST be set by the receiver to the actual length in bytes of the signature returned
	// in abSignature on output.
	SignatureSize uint32 `idl:"name:pdwSignatureSize" json:"signature_size"`
	// dwSignatureMaxSize:  MUST be set by the caller to the maximum length in bytes of
	// the server signature to be returned in abSignature. If the signature is larger than
	// the supplied buffer, the server MUST return MQ_ERROR_USER_BUFFER_TOO_SMALL (0xC00E0028).
	SignatureMaxSize uint32 `idl:"name:dwSignatureMaxSize" json:"signature_max_size"`
}

func (o *QMSetMachinePropertiesSignProcRequest) xxx_ToOp(ctx context.Context, op *xxx_QMSetMachinePropertiesSignProcOperation) *xxx_QMSetMachinePropertiesSignProcOperation {
	if op == nil {
		op = &xxx_QMSetMachinePropertiesSignProcOperation{}
	}
	if o == nil {
		return op
	}
	op.Challenge = o.Challenge
	op.ChallengeSize = o.ChallengeSize
	op.Context = o.Context
	op.Signature = o.Signature
	op.SignatureSize = o.SignatureSize
	op.SignatureMaxSize = o.SignatureMaxSize
	return op
}

func (o *QMSetMachinePropertiesSignProcRequest) xxx_FromOp(ctx context.Context, op *xxx_QMSetMachinePropertiesSignProcOperation) {
	if o == nil {
		return
	}
	o.Challenge = op.Challenge
	o.ChallengeSize = op.ChallengeSize
	o.Context = op.Context
	o.Signature = op.Signature
	o.SignatureSize = op.SignatureSize
	o.SignatureMaxSize = op.SignatureMaxSize
}
func (o *QMSetMachinePropertiesSignProcRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QMSetMachinePropertiesSignProcRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QMSetMachinePropertiesSignProcOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QMSetMachinePropertiesSignProcResponse structure represents the S_DSQMSetMachinePropertiesSignProc operation response
type QMSetMachinePropertiesSignProcResponse struct {
	// XXX: dwSignatureMaxSize is an implicit input depedency for output parameters
	SignatureMaxSize uint32 `idl:"name:dwSignatureMaxSize" json:"signature_max_size"`
	// abSignature: MUST be set by the caller to a pointer to a buffer to contain the returned
	// signature. MUST be set by the receiver to a signature over the challenge in abChallenge.
	// The algorithm for creating this signature is specified by the following pseudocode.
	Signature []byte `idl:"name:abSignature;size_is:(dwSignatureMaxSize);length_is:(pdwSignatureSize)" json:"signature"`
	// pdwSignatureSize:  Size in bytes of the signature in the abSignature parameter.
	// MUST be set by the receiver to the actual length in bytes of the signature returned
	// in abSignature on output.
	SignatureSize uint32 `idl:"name:pdwSignatureSize" json:"signature_size"`
	// Return: The S_DSQMSetMachinePropertiesSignProc return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QMSetMachinePropertiesSignProcResponse) xxx_ToOp(ctx context.Context, op *xxx_QMSetMachinePropertiesSignProcOperation) *xxx_QMSetMachinePropertiesSignProcOperation {
	if op == nil {
		op = &xxx_QMSetMachinePropertiesSignProcOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.SignatureMaxSize == uint32(0) {
		op.SignatureMaxSize = o.SignatureMaxSize
	}

	op.Signature = o.Signature
	op.SignatureSize = o.SignatureSize
	op.Return = o.Return
	return op
}

func (o *QMSetMachinePropertiesSignProcResponse) xxx_FromOp(ctx context.Context, op *xxx_QMSetMachinePropertiesSignProcOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.SignatureMaxSize = op.SignatureMaxSize

	o.Signature = op.Signature
	o.SignatureSize = op.SignatureSize
	o.Return = op.Return
}
func (o *QMSetMachinePropertiesSignProcResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QMSetMachinePropertiesSignProcResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QMSetMachinePropertiesSignProcOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QMGetObjectSecurityOperation structure represents the S_DSQMGetObjectSecurity operation
type xxx_QMGetObjectSecurityOperation struct {
	ObjectType          uint32          `idl:"name:dwObjectType" json:"object_type"`
	GUID                *dtyp.GUID      `idl:"name:pGuid" json:"guid"`
	SecurityInformation uint32          `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  []byte          `idl:"name:pSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	Length              uint32          `idl:"name:nLength" json:"length"`
	LengthNeeded        uint32          `idl:"name:lpnLengthNeeded" json:"length_needed"`
	Context             uint32          `idl:"name:dwContext" json:"context"`
	ServerAuth          *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	ServerSignature     []byte          `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	ServerSignatureSize uint32          `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	Return              int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QMGetObjectSecurityOperation) OpNum() int { return 22 }

func (o *xxx_QMGetObjectSecurityOperation) OpName() string {
	return "/dscomm/v1/S_DSQMGetObjectSecurity"
}

func (o *xxx_QMGetObjectSecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ObjectType < uint32(1) || o.ObjectType > uint32(58) {
		return fmt.Errorf("ObjectType is out of range")
	}
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

func (o *xxx_QMGetObjectSecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.WriteData(o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecurityInformation {in} (1:(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// nLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	// dwContext {in} (1:(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMGetObjectSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwObjectType {in} (1:{range=(1,58)}(uint32))
	{
		if err := w.ReadData(&o.ObjectType); err != nil {
			return err
		}
	}
	// pGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// nLength {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	// dwContext {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	// phServerAuth {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMGetObjectSecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ServerSignature != nil && o.ServerSignatureSize == 0 {
		o.ServerSignatureSize = uint32(len(o.ServerSignature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMGetObjectSecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpnLengthNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.LengthNeeded); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
	{
		dimSize1 := uint64(o.ServerSignatureSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerSignature {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerSignature[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerSignature); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.WriteData(o.ServerSignatureSize); err != nil {
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

func (o *xxx_QMGetObjectSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpnLengthNeeded {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.LengthNeeded); err != nil {
			return err
		}
	}
	// pbServerSignature {out} (1:{pointer=ref}*(1)[dim:0,size_is=pdwServerSignatureSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerSignature", sizeInfo[0])
		}
		o.ServerSignature = make([]byte, sizeInfo[0])
		for i1 := range o.ServerSignature {
			i1 := i1
			if err := w.ReadData(&o.ServerSignature[i1]); err != nil {
				return err
			}
		}
	}
	// pdwServerSignatureSize {in, out} (1:{alias=LPBOUNDED_SIGNATURE_SIZE,pointer=ref}*(1))(2:{range=(0,131072), alias=BOUNDED_SIGNATURE_SIZE}(uint32))
	{
		if err := w.ReadData(&o.ServerSignatureSize); err != nil {
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

// QMGetObjectSecurityRequest structure represents the S_DSQMGetObjectSecurity operation request
type QMGetObjectSecurityRequest struct {
	// dwObjectType: Specifies the type of object for which the security descriptor is retrieved.
	// MUST be set to one of the object types specified in section 2.2.8.
	ObjectType uint32 `idl:"name:dwObjectType" json:"object_type"`
	// pGuid: MUST be set by the client to a pointer to the GUID of the object to retrieve
	// security information.
	GUID *dtyp.GUID `idl:"name:pGuid" json:"guid"`
	// SecurityInformation:  MUST be set by the client to a bitwise mask specifying the
	// information to return in the pSecurityDescriptor parameter. See the SecurityInformation
	// parameter description in section 3.1.4.11.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// nLength: MUST be set by the client to the length in bytes of the pSecurityDescriptor
	// buffer.
	Length uint32 `idl:"name:nLength" json:"length"`
	// dwContext:  MUST be set by the client to a value that the client can use to correlate
	// callbacks with the initial call to S_DSQMGetObjectSecurityChallengeResponceProc.
	// The server MUST supply this value in the dwContext parameter in the related calls
	// of the S_DSQMGetObjectSecurityChallengeResponceProc (section 3.2.4.2) callback method.
	Context uint32 `idl:"name:dwContext" json:"context"`
	// phServerAuth:  A PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle acquired from
	// the pphServerAuth parameter in a previous call to S_DSValidateServer. The server
	// MUST use this parameter as a key to locate the GSS security context used to compute
	// the signature returned in pbServerSignature. See section 3.1.4.2.
	ServerAuth *ServerAuthType `idl:"name:phServerAuth" json:"server_auth"`
	// pdwServerSignatureSize: A DWORD that contains the maximum length in bytes of the
	// server signature to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
}

func (o *QMGetObjectSecurityRequest) xxx_ToOp(ctx context.Context, op *xxx_QMGetObjectSecurityOperation) *xxx_QMGetObjectSecurityOperation {
	if op == nil {
		op = &xxx_QMGetObjectSecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectType = o.ObjectType
	op.GUID = o.GUID
	op.SecurityInformation = o.SecurityInformation
	op.Length = o.Length
	op.Context = o.Context
	op.ServerAuth = o.ServerAuth
	op.ServerSignatureSize = o.ServerSignatureSize
	return op
}

func (o *QMGetObjectSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_QMGetObjectSecurityOperation) {
	if o == nil {
		return
	}
	o.ObjectType = op.ObjectType
	o.GUID = op.GUID
	o.SecurityInformation = op.SecurityInformation
	o.Length = op.Length
	o.Context = op.Context
	o.ServerAuth = op.ServerAuth
	o.ServerSignatureSize = op.ServerSignatureSize
}
func (o *QMGetObjectSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QMGetObjectSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QMGetObjectSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QMGetObjectSecurityResponse structure represents the S_DSQMGetObjectSecurity operation response
type QMGetObjectSecurityResponse struct {
	// XXX: nLength is an implicit input depedency for output parameters
	Length uint32 `idl:"name:nLength" json:"length"`
	// pSecurityDescriptor: If the SecurityInformation parameter is MQDS_SIGN_PUBLIC_KEY
	// or MQDS_KEYX_PUBLIC_KEY, it SHOULD <73> contain a pointer to a BLOBHEADER structure
	// followed by an RSAPUBKEY (section 2.2.18) structure. Otherwise, this parameter contains
	// a security descriptor, as specified in [MS-DTYP] section 2.4.6.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(nLength)" json:"security_descriptor"`
	// lpnLengthNeeded: A DWORD that the server MUST set to the length in bytes of the requested
	// security descriptor or Public Key. If the requested security descriptor or Public
	// Key is larger than nLength, the server MUST set this parameter to the size in bytes
	// needed for the requested security descriptor or Public Key, and return MQ_ERROR_SECURITY_DESCRIPTOR_TOO_SMALL
	// (0xC00E0023).
	LengthNeeded uint32 `idl:"name:lpnLengthNeeded" json:"length_needed"`
	// pbServerSignature: MUST point to the signed hash of the security descriptor.
	ServerSignature []byte `idl:"name:pbServerSignature;size_is:(pdwServerSignatureSize)" json:"server_signature"`
	// pdwServerSignatureSize: A DWORD that contains the maximum length in bytes of the
	// server signature to return.
	ServerSignatureSize uint32 `idl:"name:pdwServerSignatureSize" json:"server_signature_size"`
	// Return: The S_DSQMGetObjectSecurity return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QMGetObjectSecurityResponse) xxx_ToOp(ctx context.Context, op *xxx_QMGetObjectSecurityOperation) *xxx_QMGetObjectSecurityOperation {
	if op == nil {
		op = &xxx_QMGetObjectSecurityOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Length == uint32(0) {
		op.Length = o.Length
	}

	op.SecurityDescriptor = o.SecurityDescriptor
	op.LengthNeeded = o.LengthNeeded
	op.ServerSignature = o.ServerSignature
	op.ServerSignatureSize = o.ServerSignatureSize
	op.Return = o.Return
	return op
}

func (o *QMGetObjectSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_QMGetObjectSecurityOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Length = op.Length

	o.SecurityDescriptor = op.SecurityDescriptor
	o.LengthNeeded = op.LengthNeeded
	o.ServerSignature = op.ServerSignature
	o.ServerSignatureSize = op.ServerSignatureSize
	o.Return = op.Return
}
func (o *QMGetObjectSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QMGetObjectSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QMGetObjectSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QMGetObjectSecurityChallengeResponseProcOperation structure represents the S_DSQMGetObjectSecurityChallengeResponceProc operation
type xxx_QMGetObjectSecurityChallengeResponseProcOperation struct {
	Challenge                []byte `idl:"name:abChallenge;size_is:(dwCallengeSize)" json:"challenge"`
	ChallengeSize            uint32 `idl:"name:dwCallengeSize" json:"challenge_size"`
	Context                  uint32 `idl:"name:dwContext" json:"context"`
	ChallengeResponse        []byte `idl:"name:abCallengeResponce;size_is:(dwCallengeResponceMaxSize);length_is:(pdwCallengeResponceSize)" json:"challenge_response"`
	ChallengeResponseSize    uint32 `idl:"name:pdwCallengeResponceSize" json:"challenge_response_size"`
	ChallengeResponseMaxSize uint32 `idl:"name:dwCallengeResponceMaxSize" json:"challenge_response_max_size"`
	Return                   int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_QMGetObjectSecurityChallengeResponseProcOperation) OpNum() int { return 23 }

func (o *xxx_QMGetObjectSecurityChallengeResponseProcOperation) OpName() string {
	return "/dscomm/v1/S_DSQMGetObjectSecurityChallengeResponceProc"
}

func (o *xxx_QMGetObjectSecurityChallengeResponseProcOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Challenge != nil && o.ChallengeSize == 0 {
		o.ChallengeSize = uint32(len(o.Challenge))
	}
	if o.ChallengeResponse != nil && o.ChallengeResponseMaxSize == 0 {
		o.ChallengeResponseMaxSize = uint32(len(o.ChallengeResponse))
	}
	if o.ChallengeResponse != nil && o.ChallengeResponseSize == 0 {
		o.ChallengeResponseSize = uint32(len(o.ChallengeResponse))
	}
	if o.ChallengeSize > uint32(32) {
		return fmt.Errorf("ChallengeSize is out of range")
	}
	if o.ChallengeResponseMaxSize > uint32(128) {
		return fmt.Errorf("ChallengeResponseMaxSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMGetObjectSecurityChallengeResponseProcOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// abChallenge {in} (1:{pointer=ref}*(1)[dim:0,size_is=dwCallengeSize](uint8))
	{
		dimSize1 := uint64(o.ChallengeSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Challenge {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Challenge[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Challenge); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// dwCallengeSize {in} (1:{range=(0,32)}(uint32))
	{
		if err := w.WriteData(o.ChallengeSize); err != nil {
			return err
		}
	}
	// dwContext {in} (1:(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	// abCallengeResponce {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwCallengeResponceMaxSize,length_is=pdwCallengeResponceSize](uint8))
	{
		dimSize1 := uint64(o.ChallengeResponseMaxSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ChallengeResponseSize)
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
		for i1 := range o.ChallengeResponse {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ChallengeResponse[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ChallengeResponse); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwCallengeResponceSize {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ChallengeResponseSize); err != nil {
			return err
		}
	}
	// dwCallengeResponceMaxSize {in} (1:{range=(0,128)}(uint32))
	{
		if err := w.WriteData(o.ChallengeResponseMaxSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMGetObjectSecurityChallengeResponseProcOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// abChallenge {in} (1:{pointer=ref}*(1)[dim:0,size_is=dwCallengeSize](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Challenge", sizeInfo[0])
		}
		o.Challenge = make([]byte, sizeInfo[0])
		for i1 := range o.Challenge {
			i1 := i1
			if err := w.ReadData(&o.Challenge[i1]); err != nil {
				return err
			}
		}
	}
	// dwCallengeSize {in} (1:{range=(0,32)}(uint32))
	{
		if err := w.ReadData(&o.ChallengeSize); err != nil {
			return err
		}
	}
	// dwContext {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	// abCallengeResponce {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwCallengeResponceMaxSize,length_is=pdwCallengeResponceSize](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ChallengeResponse", sizeInfo[0])
		}
		o.ChallengeResponse = make([]byte, sizeInfo[0])
		for i1 := range o.ChallengeResponse {
			i1 := i1
			if err := w.ReadData(&o.ChallengeResponse[i1]); err != nil {
				return err
			}
		}
	}
	// pdwCallengeResponceSize {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ChallengeResponseSize); err != nil {
			return err
		}
	}
	// dwCallengeResponceMaxSize {in} (1:{range=(0,128)}(uint32))
	{
		if err := w.ReadData(&o.ChallengeResponseMaxSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMGetObjectSecurityChallengeResponseProcOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ChallengeResponse != nil && o.ChallengeResponseSize == 0 {
		o.ChallengeResponseSize = uint32(len(o.ChallengeResponse))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QMGetObjectSecurityChallengeResponseProcOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// abCallengeResponce {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwCallengeResponceMaxSize,length_is=pdwCallengeResponceSize](uint8))
	{
		dimSize1 := uint64(o.ChallengeResponseMaxSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ChallengeResponseSize)
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
		for i1 := range o.ChallengeResponse {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ChallengeResponse[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ChallengeResponse); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwCallengeResponceSize {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ChallengeResponseSize); err != nil {
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

func (o *xxx_QMGetObjectSecurityChallengeResponseProcOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// abCallengeResponce {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwCallengeResponceMaxSize,length_is=pdwCallengeResponceSize](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ChallengeResponse", sizeInfo[0])
		}
		o.ChallengeResponse = make([]byte, sizeInfo[0])
		for i1 := range o.ChallengeResponse {
			i1 := i1
			if err := w.ReadData(&o.ChallengeResponse[i1]); err != nil {
				return err
			}
		}
	}
	// pdwCallengeResponceSize {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ChallengeResponseSize); err != nil {
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

// QMGetObjectSecurityChallengeResponseProcRequest structure represents the S_DSQMGetObjectSecurityChallengeResponceProc operation request
type QMGetObjectSecurityChallengeResponseProcRequest struct {
	// abChallenge:  MUST be set by the caller to a pointer to a buffer that contains the
	// challenge to be signed. The challenge SHOULD be cryptographically random.
	Challenge []byte `idl:"name:abChallenge;size_is:(dwCallengeSize)" json:"challenge"`
	// dwCallengeSize:  MUST be set by the caller to the size, in bytes, of the challenge
	// in the abChallenge parameter.
	ChallengeSize uint32 `idl:"name:dwCallengeSize" json:"challenge_size"`
	// dwContext:  MUST be set by the caller to the value that was supplied in the dwContext
	// parameter of the corresponding call to the S_DSQMGetObjectSecurity method. This parameter
	// provides a way for the receiver to correlate the callback with the receiver's in-progress
	// call to S_DSQMGetObjectSecurity.
	Context uint32 `idl:"name:dwContext" json:"context"`
	// abCallengeResponce:  MUST be set by the caller to a pointer to a buffer to contain
	// the returned signature. MUST be set by the receiver to a signature over the challenge
	// in abChallenge. The algorithm for creating this signature is specified by the following
	// pseudocode.
	ChallengeResponse []byte `idl:"name:abCallengeResponce;size_is:(dwCallengeResponceMaxSize);length_is:(pdwCallengeResponceSize)" json:"challenge_response"`
	// pdwCallengeResponceSize:  Size in bytes of the signature in the abCallengeResponce
	// parameter. MUST be set by the receiver to the actual length, in bytes, of the signature
	// returned in abCallengeResponce on output.
	ChallengeResponseSize uint32 `idl:"name:pdwCallengeResponceSize" json:"challenge_response_size"`
	// dwCallengeResponceMaxSize:  MUST be set by the caller to the maximum length in bytes
	// of the server signature to be returned in abCallengeResponce. If the server signature
	// is larger than the supplied buffer, the server MUST return MQ_ERROR_USER_BUFFER_TOO_SMALL
	// (0xC00E0028).
	ChallengeResponseMaxSize uint32 `idl:"name:dwCallengeResponceMaxSize" json:"challenge_response_max_size"`
}

func (o *QMGetObjectSecurityChallengeResponseProcRequest) xxx_ToOp(ctx context.Context, op *xxx_QMGetObjectSecurityChallengeResponseProcOperation) *xxx_QMGetObjectSecurityChallengeResponseProcOperation {
	if op == nil {
		op = &xxx_QMGetObjectSecurityChallengeResponseProcOperation{}
	}
	if o == nil {
		return op
	}
	op.Challenge = o.Challenge
	op.ChallengeSize = o.ChallengeSize
	op.Context = o.Context
	op.ChallengeResponse = o.ChallengeResponse
	op.ChallengeResponseSize = o.ChallengeResponseSize
	op.ChallengeResponseMaxSize = o.ChallengeResponseMaxSize
	return op
}

func (o *QMGetObjectSecurityChallengeResponseProcRequest) xxx_FromOp(ctx context.Context, op *xxx_QMGetObjectSecurityChallengeResponseProcOperation) {
	if o == nil {
		return
	}
	o.Challenge = op.Challenge
	o.ChallengeSize = op.ChallengeSize
	o.Context = op.Context
	o.ChallengeResponse = op.ChallengeResponse
	o.ChallengeResponseSize = op.ChallengeResponseSize
	o.ChallengeResponseMaxSize = op.ChallengeResponseMaxSize
}
func (o *QMGetObjectSecurityChallengeResponseProcRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QMGetObjectSecurityChallengeResponseProcRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QMGetObjectSecurityChallengeResponseProcOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QMGetObjectSecurityChallengeResponseProcResponse structure represents the S_DSQMGetObjectSecurityChallengeResponceProc operation response
type QMGetObjectSecurityChallengeResponseProcResponse struct {
	// XXX: dwCallengeResponceMaxSize is an implicit input depedency for output parameters
	ChallengeResponseMaxSize uint32 `idl:"name:dwCallengeResponceMaxSize" json:"challenge_response_max_size"`
	// abCallengeResponce:  MUST be set by the caller to a pointer to a buffer to contain
	// the returned signature. MUST be set by the receiver to a signature over the challenge
	// in abChallenge. The algorithm for creating this signature is specified by the following
	// pseudocode.
	ChallengeResponse []byte `idl:"name:abCallengeResponce;size_is:(dwCallengeResponceMaxSize);length_is:(pdwCallengeResponceSize)" json:"challenge_response"`
	// pdwCallengeResponceSize:  Size in bytes of the signature in the abCallengeResponce
	// parameter. MUST be set by the receiver to the actual length, in bytes, of the signature
	// returned in abCallengeResponce on output.
	ChallengeResponseSize uint32 `idl:"name:pdwCallengeResponceSize" json:"challenge_response_size"`
	// Return: The S_DSQMGetObjectSecurityChallengeResponceProc return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QMGetObjectSecurityChallengeResponseProcResponse) xxx_ToOp(ctx context.Context, op *xxx_QMGetObjectSecurityChallengeResponseProcOperation) *xxx_QMGetObjectSecurityChallengeResponseProcOperation {
	if op == nil {
		op = &xxx_QMGetObjectSecurityChallengeResponseProcOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ChallengeResponseMaxSize == uint32(0) {
		op.ChallengeResponseMaxSize = o.ChallengeResponseMaxSize
	}

	op.ChallengeResponse = o.ChallengeResponse
	op.ChallengeResponseSize = o.ChallengeResponseSize
	op.Return = o.Return
	return op
}

func (o *QMGetObjectSecurityChallengeResponseProcResponse) xxx_FromOp(ctx context.Context, op *xxx_QMGetObjectSecurityChallengeResponseProcOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ChallengeResponseMaxSize = op.ChallengeResponseMaxSize

	o.ChallengeResponse = op.ChallengeResponse
	o.ChallengeResponseSize = op.ChallengeResponseSize
	o.Return = op.Return
}
func (o *QMGetObjectSecurityChallengeResponseProcResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QMGetObjectSecurityChallengeResponseProcResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QMGetObjectSecurityChallengeResponseProcOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InitSecurityContextOperation structure represents the S_InitSecCtx operation
type xxx_InitSecurityContextOperation struct {
	Context             uint32 `idl:"name:dwContext" json:"context"`
	ServerBuffer        []byte `idl:"name:pServerbuff;size_is:(dwServerBuffSize)" json:"server_buffer"`
	ServerBufferSize    uint32 `idl:"name:dwServerBuffSize" json:"server_buffer_size"`
	ClientBufferMaxSize uint32 `idl:"name:dwClientBuffMaxSize" json:"client_buffer_max_size"`
	ClientBuffer        []byte `idl:"name:pClientBuff;size_is:(dwClientBuffMaxSize);length_is:(pdwClientBuffSize)" json:"client_buffer"`
	ClientBufferSize    uint32 `idl:"name:pdwClientBuffSize" json:"client_buffer_size"`
	Return              int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_InitSecurityContextOperation) OpNum() int { return 24 }

func (o *xxx_InitSecurityContextOperation) OpName() string { return "/dscomm/v1/S_InitSecCtx" }

func (o *xxx_InitSecurityContextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ServerBuffer != nil && o.ServerBufferSize == 0 {
		o.ServerBufferSize = uint32(len(o.ServerBuffer))
	}
	if o.ServerBufferSize > uint32(524288) {
		return fmt.Errorf("ServerBufferSize is out of range")
	}
	if o.ClientBufferMaxSize > uint32(524288) {
		return fmt.Errorf("ClientBufferMaxSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitSecurityContextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwContext {in} (1:(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	// pServerbuff {in} (1:{pointer=ref}*(1)[dim:0,size_is=dwServerBuffSize](uchar))
	{
		dimSize1 := uint64(o.ServerBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ServerBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ServerBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ServerBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// dwServerBuffSize {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.WriteData(o.ServerBufferSize); err != nil {
			return err
		}
	}
	// dwClientBuffMaxSize {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.WriteData(o.ClientBufferMaxSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitSecurityContextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwContext {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	// pServerbuff {in} (1:{pointer=ref}*(1)[dim:0,size_is=dwServerBuffSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ServerBuffer", sizeInfo[0])
		}
		o.ServerBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.ServerBuffer {
			i1 := i1
			if err := w.ReadData(&o.ServerBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// dwServerBuffSize {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.ReadData(&o.ServerBufferSize); err != nil {
			return err
		}
	}
	// dwClientBuffMaxSize {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.ReadData(&o.ClientBufferMaxSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitSecurityContextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ClientBuffer != nil && o.ClientBufferSize == 0 {
		o.ClientBufferSize = uint32(len(o.ClientBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitSecurityContextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pClientBuff {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwClientBuffMaxSize,length_is=pdwClientBuffSize](uchar))
	{
		dimSize1 := uint64(o.ClientBufferMaxSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ClientBufferSize)
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
		for i1 := range o.ClientBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ClientBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ClientBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwClientBuffSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ClientBufferSize); err != nil {
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

func (o *xxx_InitSecurityContextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pClientBuff {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwClientBuffMaxSize,length_is=pdwClientBuffSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ClientBuffer", sizeInfo[0])
		}
		o.ClientBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.ClientBuffer {
			i1 := i1
			if err := w.ReadData(&o.ClientBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// pdwClientBuffSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ClientBufferSize); err != nil {
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

// InitSecurityContextRequest structure represents the S_InitSecCtx operation request
type InitSecurityContextRequest struct {
	// dwContext:  MUST be set by the caller to the correlation value supplied by the client
	// in the dwContext parameter in the corresponding call to S_DSValidateServer. This
	// parameter provides a way for the receiver to correlate the callback with the receiver's
	// in-progress call to S_DSValidateServer.
	Context uint32 `idl:"name:dwContext" json:"context"`
	// pServerbuff:  MUST be set by the caller to point to a buffer that contains the output_token
	// from the GSS_Accept_sec_context, as specified in [RFC2743].
	ServerBuffer []byte `idl:"name:pServerbuff;size_is:(dwServerBuffSize)" json:"server_buffer"`
	// dwServerBuffSize:  MUST be set by the caller to the length, in bytes, of the output_token
	// within pServerBuff.
	ServerBufferSize uint32 `idl:"name:dwServerBuffSize" json:"server_buffer_size"`
	// dwClientBuffMaxSize:  MUST be set by the caller to the size, in bytes, of the buffer
	// to be returned in pClientBuff.
	ClientBufferMaxSize uint32 `idl:"name:dwClientBuffMaxSize" json:"client_buffer_max_size"`
}

func (o *InitSecurityContextRequest) xxx_ToOp(ctx context.Context, op *xxx_InitSecurityContextOperation) *xxx_InitSecurityContextOperation {
	if op == nil {
		op = &xxx_InitSecurityContextOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.ServerBuffer = o.ServerBuffer
	op.ServerBufferSize = o.ServerBufferSize
	op.ClientBufferMaxSize = o.ClientBufferMaxSize
	return op
}

func (o *InitSecurityContextRequest) xxx_FromOp(ctx context.Context, op *xxx_InitSecurityContextOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.ServerBuffer = op.ServerBuffer
	o.ServerBufferSize = op.ServerBufferSize
	o.ClientBufferMaxSize = op.ClientBufferMaxSize
}
func (o *InitSecurityContextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InitSecurityContextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitSecurityContextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitSecurityContextResponse structure represents the S_InitSecCtx operation response
type InitSecurityContextResponse struct {
	// XXX: dwClientBuffMaxSize is an implicit input depedency for output parameters
	ClientBufferMaxSize uint32 `idl:"name:dwClientBuffMaxSize" json:"client_buffer_max_size"`
	// pClientBuff:  MUST be set by the caller to point to a buffer to hold the returned
	// token. MUST be set by the receiver to the output_token from a call to GSS_Init_sec_context.
	// The buffer length MUST NOT exceed the value specified by dwClientBuffMaxSize. If
	// the negotiated token is larger than the supplied buffer, the server MUST return MQ_ERROR_USER_BUFFER_TOO_SMALL
	// (0xC00E0028).
	ClientBuffer []byte `idl:"name:pClientBuff;size_is:(dwClientBuffMaxSize);length_is:(pdwClientBuffSize)" json:"client_buffer"`
	// pdwClientBuffSize:  MUST be set by the receiver to the actual size, in bytes, of
	// the token in pClientBuff.
	ClientBufferSize uint32 `idl:"name:pdwClientBuffSize" json:"client_buffer_size"`
	// Return: The S_InitSecCtx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InitSecurityContextResponse) xxx_ToOp(ctx context.Context, op *xxx_InitSecurityContextOperation) *xxx_InitSecurityContextOperation {
	if op == nil {
		op = &xxx_InitSecurityContextOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ClientBufferMaxSize == uint32(0) {
		op.ClientBufferMaxSize = o.ClientBufferMaxSize
	}

	op.ClientBuffer = o.ClientBuffer
	op.ClientBufferSize = o.ClientBufferSize
	op.Return = o.Return
	return op
}

func (o *InitSecurityContextResponse) xxx_FromOp(ctx context.Context, op *xxx_InitSecurityContextOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ClientBufferMaxSize = op.ClientBufferMaxSize

	o.ClientBuffer = op.ClientBuffer
	o.ClientBufferSize = op.ClientBufferSize
	o.Return = op.Return
}
func (o *InitSecurityContextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InitSecurityContextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitSecurityContextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ValidateServerOperation structure represents the S_DSValidateServer operation
type xxx_ValidateServerOperation struct {
	EnterpriseID        *dtyp.GUID      `idl:"name:pguidEnterpriseId" json:"enterprise_id"`
	SetupMode           bool            `idl:"name:fSetupMode" json:"setup_mode"`
	Context             uint32          `idl:"name:dwContext" json:"context"`
	ClientBufferMaxSize uint32          `idl:"name:dwClientBuffMaxSize" json:"client_buffer_max_size"`
	ClientBuffer        []byte          `idl:"name:pClientBuff;size_is:(dwClientBuffMaxSize);length_is:(dwClientBuffSize)" json:"client_buffer"`
	ClientBufferSize    uint32          `idl:"name:dwClientBuffSize" json:"client_buffer_size"`
	ServerAuth          *ServerAuthType `idl:"name:pphServerAuth" json:"server_auth"`
	Return              int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_ValidateServerOperation) OpNum() int { return 25 }

func (o *xxx_ValidateServerOperation) OpName() string { return "/dscomm/v1/S_DSValidateServer" }

func (o *xxx_ValidateServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ClientBuffer != nil && o.ClientBufferMaxSize == 0 {
		o.ClientBufferMaxSize = uint32(len(o.ClientBuffer))
	}
	if o.ClientBuffer != nil && o.ClientBufferSize == 0 {
		o.ClientBufferSize = uint32(len(o.ClientBuffer))
	}
	if o.ClientBufferMaxSize > uint32(524288) {
		return fmt.Errorf("ClientBufferMaxSize is out of range")
	}
	if o.ClientBufferSize > uint32(524288) {
		return fmt.Errorf("ClientBufferSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pguidEnterpriseId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.EnterpriseID != nil {
			if err := o.EnterpriseID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// fSetupMode {in} (1:{alias=BOOL}(int32))
	{
		if !o.SetupMode {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// dwContext {in} (1:(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	// dwClientBuffMaxSize {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.WriteData(o.ClientBufferMaxSize); err != nil {
			return err
		}
	}
	// pClientBuff {in} (1:{pointer=ref}*(1)[dim:0,size_is=dwClientBuffMaxSize,length_is=dwClientBuffSize](uchar))
	{
		dimSize1 := uint64(o.ClientBufferMaxSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ClientBufferSize)
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
		for i1 := range o.ClientBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ClientBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ClientBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// dwClientBuffSize {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.WriteData(o.ClientBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pguidEnterpriseId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.EnterpriseID == nil {
			o.EnterpriseID = &dtyp.GUID{}
		}
		if err := o.EnterpriseID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// fSetupMode {in} (1:{alias=BOOL}(int32))
	{
		var _bSetupMode int32
		if err := w.ReadData(&_bSetupMode); err != nil {
			return err
		}
		o.SetupMode = _bSetupMode != 0
	}
	// dwContext {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	// dwClientBuffMaxSize {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.ReadData(&o.ClientBufferMaxSize); err != nil {
			return err
		}
	}
	// pClientBuff {in} (1:{pointer=ref}*(1)[dim:0,size_is=dwClientBuffMaxSize,length_is=dwClientBuffSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ClientBuffer", sizeInfo[0])
		}
		o.ClientBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.ClientBuffer {
			i1 := i1
			if err := w.ReadData(&o.ClientBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// dwClientBuffSize {in} (1:{range=(0,524288)}(uint32))
	{
		if err := w.ReadData(&o.ClientBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphServerAuth {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_SERVER_AUTH_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ValidateServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphServerAuth {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_SERVER_AUTH_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
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

// ValidateServerRequest structure represents the S_DSValidateServer operation request
type ValidateServerRequest struct {
	// pguidEnterpriseId:  MUST be set by the client to a pointer to the GUID of the enterprise
	// that owns the server. The server SHOULD ignore it.<17>
	EnterpriseID *dtyp.GUID `idl:"name:pguidEnterpriseId" json:"enterprise_id"`
	// fSetupMode:  Boolean value that defines the setup mode. Clients SHOULD set this
	// value to FALSE (0x00000000). The server SHOULD ignore it.<18>
	SetupMode bool `idl:"name:fSetupMode" json:"setup_mode"`
	// dwContext:  MUST be set by the client to a value that the client can use to correlate
	// callbacks with the initial call to S_DSValidateServer. The server MUST supply this
	// value in the dwContext parameter in the related calls of the S_InitSecCtx (section
	// 3.2.4.3) callback method.
	Context uint32 `idl:"name:dwContext" json:"context"`
	// dwClientBuffMaxSize:  MUST be set by the client to the size of the buffer pointed
	// to by the pClientBuff parameter.
	ClientBufferMaxSize uint32 `idl:"name:dwClientBuffMaxSize" json:"client_buffer_max_size"`
	// pClientBuff: A pointer that MUST be set by the client to point to a buffer that contains
	// the output_token from an initial call to GSS_Init_sec_context, as specified in [RFC2743]
	// section 2.2.1.
	ClientBuffer []byte `idl:"name:pClientBuff;size_is:(dwClientBuffMaxSize);length_is:(dwClientBuffSize)" json:"client_buffer"`
	// dwClientBuffSize:  MUST be set by the client to the size in bytes of the output_token
	// within the client buffer pointed at by the pClientBuff parameter.
	ClientBufferSize uint32 `idl:"name:dwClientBuffSize" json:"client_buffer_size"`
}

func (o *ValidateServerRequest) xxx_ToOp(ctx context.Context, op *xxx_ValidateServerOperation) *xxx_ValidateServerOperation {
	if op == nil {
		op = &xxx_ValidateServerOperation{}
	}
	if o == nil {
		return op
	}
	op.EnterpriseID = o.EnterpriseID
	op.SetupMode = o.SetupMode
	op.Context = o.Context
	op.ClientBufferMaxSize = o.ClientBufferMaxSize
	op.ClientBuffer = o.ClientBuffer
	op.ClientBufferSize = o.ClientBufferSize
	return op
}

func (o *ValidateServerRequest) xxx_FromOp(ctx context.Context, op *xxx_ValidateServerOperation) {
	if o == nil {
		return
	}
	o.EnterpriseID = op.EnterpriseID
	o.SetupMode = op.SetupMode
	o.Context = op.Context
	o.ClientBufferMaxSize = op.ClientBufferMaxSize
	o.ClientBuffer = op.ClientBuffer
	o.ClientBufferSize = op.ClientBufferSize
}
func (o *ValidateServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ValidateServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ValidateServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ValidateServerResponse structure represents the S_DSValidateServer operation response
type ValidateServerResponse struct {
	// pphServerAuth:  MUST be set by the server to a PCONTEXT_HANDLE_SERVER_AUTH_TYPE
	// (section 2.2.5) RPC context handle.
	ServerAuth *ServerAuthType `idl:"name:pphServerAuth" json:"server_auth"`
	// Return: The S_DSValidateServer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ValidateServerResponse) xxx_ToOp(ctx context.Context, op *xxx_ValidateServerOperation) *xxx_ValidateServerOperation {
	if op == nil {
		op = &xxx_ValidateServerOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerAuth = o.ServerAuth
	op.Return = o.Return
	return op
}

func (o *ValidateServerResponse) xxx_FromOp(ctx context.Context, op *xxx_ValidateServerOperation) {
	if o == nil {
		return
	}
	o.ServerAuth = op.ServerAuth
	o.Return = op.Return
}
func (o *ValidateServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ValidateServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ValidateServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseServerOperation structure represents the S_DSCloseServerHandle operation
type xxx_CloseServerOperation struct {
	ServerAuth *ServerAuthType `idl:"name:pphServerAuth" json:"server_auth"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseServerOperation) OpNum() int { return 26 }

func (o *xxx_CloseServerOperation) OpName() string { return "/dscomm/v1/S_DSCloseServerHandle" }

func (o *xxx_CloseServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pphServerAuth {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_SERVER_AUTH_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pphServerAuth {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_SERVER_AUTH_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphServerAuth {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_SERVER_AUTH_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth != nil {
			if err := o.ServerAuth.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerAuthType{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CloseServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphServerAuth {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE_SERVER_AUTH_TYPE}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_SERVER_AUTH_TYPE, names=ndr_context_handle}(struct))
	{
		if o.ServerAuth == nil {
			o.ServerAuth = &ServerAuthType{}
		}
		if err := o.ServerAuth.UnmarshalNDR(ctx, w); err != nil {
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

// CloseServerRequest structure represents the S_DSCloseServerHandle operation request
type CloseServerRequest struct {
	// pphServerAuth:  The PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle to close.
	// It MUST have been acquired from the pphServerAuth parameter in a previous call to
	// S_DSValidateServer, and MUST NOT have been closed through a previous call to S_DSCloseServerHandle.
	// The server MUST set this parameter to NULL.
	ServerAuth *ServerAuthType `idl:"name:pphServerAuth" json:"server_auth"`
}

func (o *CloseServerRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseServerOperation) *xxx_CloseServerOperation {
	if op == nil {
		op = &xxx_CloseServerOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerAuth = o.ServerAuth
	return op
}

func (o *CloseServerRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseServerOperation) {
	if o == nil {
		return
	}
	o.ServerAuth = op.ServerAuth
}
func (o *CloseServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseServerResponse structure represents the S_DSCloseServerHandle operation response
type CloseServerResponse struct {
	// pphServerAuth:  The PCONTEXT_HANDLE_SERVER_AUTH_TYPE RPC context handle to close.
	// It MUST have been acquired from the pphServerAuth parameter in a previous call to
	// S_DSValidateServer, and MUST NOT have been closed through a previous call to S_DSCloseServerHandle.
	// The server MUST set this parameter to NULL.
	ServerAuth *ServerAuthType `idl:"name:pphServerAuth" json:"server_auth"`
	// Return: The S_DSCloseServerHandle return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseServerResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseServerOperation) *xxx_CloseServerOperation {
	if op == nil {
		op = &xxx_CloseServerOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerAuth = o.ServerAuth
	op.Return = o.Return
	return op
}

func (o *CloseServerResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseServerOperation) {
	if o == nil {
		return
	}
	o.ServerAuth = op.ServerAuth
	o.Return = op.Return
}
func (o *CloseServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServerPortOperation structure represents the S_DSGetServerPort operation
type xxx_GetServerPortOperation struct {
	IP     uint32 `idl:"name:fIP" json:"ip"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServerPortOperation) OpNum() int { return 30 }

func (o *xxx_GetServerPortOperation) OpName() string { return "/dscomm/v1/S_DSGetServerPort" }

func (o *xxx_GetServerPortOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.IP > uint32(1) {
		return fmt.Errorf("IP is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerPortOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// fIP {in} (1:{range=(0,1)}(uint32))
	{
		if err := w.WriteData(o.IP); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerPortOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// fIP {in} (1:{range=(0,1)}(uint32))
	{
		if err := w.ReadData(&o.IP); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerPortOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerPortOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerPortOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetServerPortRequest structure represents the S_DSGetServerPort operation request
type GetServerPortRequest struct {
	// fIP:  Specifies the connected network protocol for which an RPC endpoint port is
	// to be returned.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 | Causes the RPC endpoint port for an RPC over SPX protocol sequence, as specified |
	//	|            | in [MS-RPCE], to be returned.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | Causes the RPC endpoint port for an RPC over TCP/IP protocol sequence, as        |
	//	|            | specified in [MS-RPCE], to be returned.                                          |
	//	+------------+----------------------------------------------------------------------------------+
	IP uint32 `idl:"name:fIP" json:"ip"`
}

func (o *GetServerPortRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServerPortOperation) *xxx_GetServerPortOperation {
	if op == nil {
		op = &xxx_GetServerPortOperation{}
	}
	if o == nil {
		return op
	}
	op.IP = o.IP
	return op
}

func (o *GetServerPortRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServerPortOperation) {
	if o == nil {
		return
	}
	o.IP = op.IP
}
func (o *GetServerPortRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServerPortRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerPortOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServerPortResponse structure represents the S_DSGetServerPort operation response
type GetServerPortResponse struct {
	// Return: The S_DSGetServerPort return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetServerPortResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServerPortOperation) *xxx_GetServerPortOperation {
	if op == nil {
		op = &xxx_GetServerPortOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *GetServerPortResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServerPortOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *GetServerPortResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServerPortResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerPortOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
