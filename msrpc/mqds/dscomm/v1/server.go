package dscomm

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

// dscomm server interface.
type DscommServer interface {

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
	CreateObject(context.Context, *CreateObjectRequest) (*CreateObjectResponse, error)

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
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)

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
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

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
	SetProperties(context.Context, *SetPropertiesRequest) (*SetPropertiesResponse, error)

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
	GetObjectSecurity(context.Context, *GetObjectSecurityRequest) (*GetObjectSecurityResponse, error)

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
	SetObjectSecurity(context.Context, *SetObjectSecurityRequest) (*SetObjectSecurityResponse, error)

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
	LookupBegin(context.Context, *LookupBeginRequest) (*LookupBeginResponse, error)

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
	LookupNext(context.Context, *LookupNextRequest) (*LookupNextResponse, error)

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
	LookupEnd(context.Context, *LookupEndRequest) (*LookupEndResponse, error)

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
	DeleteObjectGUID(context.Context, *DeleteObjectGUIDRequest) (*DeleteObjectGUIDResponse, error)

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
	GetPropertiesGUID(context.Context, *GetPropertiesGUIDRequest) (*GetPropertiesGUIDResponse, error)

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
	SetPropertiesGUID(context.Context, *SetPropertiesGUIDRequest) (*SetPropertiesGUIDResponse, error)

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
	GetObjectSecurityGUID(context.Context, *GetObjectSecurityGUIDRequest) (*GetObjectSecurityGUIDResponse, error)

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
	SetObjectSecurityGUID(context.Context, *SetObjectSecurityGUIDRequest) (*SetObjectSecurityGUIDResponse, error)

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
	QMSetMachineProperties(context.Context, *QMSetMachinePropertiesRequest) (*QMSetMachinePropertiesResponse, error)

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
	CreateServersCache(context.Context, *CreateServersCacheRequest) (*CreateServersCacheResponse, error)

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
	QMSetMachinePropertiesSignProc(context.Context, *QMSetMachinePropertiesSignProcRequest) (*QMSetMachinePropertiesSignProcResponse, error)

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
	QMGetObjectSecurity(context.Context, *QMGetObjectSecurityRequest) (*QMGetObjectSecurityResponse, error)

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
	QMGetObjectSecurityChallengeResponseProc(context.Context, *QMGetObjectSecurityChallengeResponseProcRequest) (*QMGetObjectSecurityChallengeResponseProcResponse, error)

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
	InitSecurityContext(context.Context, *InitSecurityContextRequest) (*InitSecurityContextResponse, error)

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
	ValidateServer(context.Context, *ValidateServerRequest) (*ValidateServerResponse, error)

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
	CloseServer(context.Context, *CloseServerRequest) (*CloseServerResponse, error)

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
	GetServerPort(context.Context, *GetServerPortRequest) (*GetServerPortResponse, error)
}

func RegisterDscommServer(conn dcerpc.Conn, o DscommServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDscommServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DscommSyntaxV1_0))...)
}

func NewDscommServerHandle(o DscommServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DscommServerHandle(ctx, o, opNum, r)
	}
}

func DscommServerHandle(ctx context.Context, o DscommServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // S_DSCreateObject
		in := &CreateObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // S_DSDeleteObject
		in := &DeleteObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // S_DSGetProps
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // S_DSSetProps
		in := &SetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // S_DSGetObjectSecurity
		in := &GetObjectSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetObjectSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // S_DSSetObjectSecurity
		in := &SetObjectSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetObjectSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // S_DSLookupBegin
		in := &LookupBeginRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupBegin(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // S_DSLookupNext
		in := &LookupNextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupNext(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // S_DSLookupEnd
		in := &LookupEndRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupEnd(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 10: // S_DSDeleteObjectGuid
		in := &DeleteObjectGUIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteObjectGUID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // S_DSGetPropsGuid
		in := &GetPropertiesGUIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertiesGUID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // S_DSSetPropsGuid
		in := &SetPropertiesGUIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPropertiesGUID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // S_DSGetObjectSecurityGuid
		in := &GetObjectSecurityGUIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetObjectSecurityGUID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // S_DSSetObjectSecurityGuid
		in := &SetObjectSecurityGUIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetObjectSecurityGUID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
	case 16: // Opnum16NotUsedOnWire
		// Opnum16NotUsedOnWire
		return nil, nil
	case 17: // Opnum17NotUsedOnWire
		// Opnum17NotUsedOnWire
		return nil, nil
	case 18: // Opnum18NotUsedOnWire
		// Opnum18NotUsedOnWire
		return nil, nil
	case 19: // S_DSQMSetMachineProperties
		in := &QMSetMachinePropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QMSetMachineProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // S_DSCreateServersCache
		in := &CreateServersCacheRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateServersCache(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // S_DSQMSetMachinePropertiesSignProc
		in := &QMSetMachinePropertiesSignProcRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QMSetMachinePropertiesSignProc(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // S_DSQMGetObjectSecurity
		in := &QMGetObjectSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QMGetObjectSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // S_DSQMGetObjectSecurityChallengeResponceProc
		in := &QMGetObjectSecurityChallengeResponseProcRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QMGetObjectSecurityChallengeResponseProc(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // S_InitSecCtx
		in := &InitSecurityContextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.InitSecurityContext(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // S_DSValidateServer
		in := &ValidateServerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ValidateServer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // S_DSCloseServerHandle
		in := &CloseServerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseServer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // Opnum24NotUsedOnWire
		// Opnum24NotUsedOnWire
		return nil, nil
	case 28: // Opnum25NotUsedOnWire
		// Opnum25NotUsedOnWire
		return nil, nil
	case 29: // Opnum26NotUsedOnWire
		// Opnum26NotUsedOnWire
		return nil, nil
	case 30: // S_DSGetServerPort
		in := &GetServerPortRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServerPort(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
