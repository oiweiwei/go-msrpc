package ixnremote

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

// IXnRemote server interface.
type IxnRemoteServer interface {

	// The Poke method is used by a secondary partner to request the primary partner session
	// initiation. The parameter values specified in the call identify both participants.
	//
	// Return Values: This method MUST return zero (0x00000000) on success. On failure,
	// it MUST return either one of the values described in the following table or an implementation-specific
	// HRESULT. A client MUST NOT depend on implementation-specific failure HRESULT values.
	// For more information about how the client SHOULD behave based on the possible return
	// values, see section 3.4.6.1.2. Standard errors are defined in [MS-ERREF] section
	// 2.2.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_STATUS                  | The return value indicates success.                                              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000123 E_CM_SERVER_NOT_READY         | The session object is not in the Connecting state.<16>                           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                  | The return value indicates that one of the specified arguments is invalid.<17>   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006BB RPC_S_SERVER_TOO_BUSY         | The return value indicates that the partner is too busy to complete this         |
	//	|                                          | operation. For more information, see [MS-RPCE] section 3.1.1.5.5                 |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000173 E_CM_S_PROTOCOL_NOT_SUPPORTED | The return value indicates that none of the protocols described in the rguchBlob |
	//	|                                          | parameter is supported by the partner.                                           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is zero.
	//
	// Poke SHOULD NOT be invoked on a secondary partner. If it is, the secondary partner
	// SHOULD respond by making a Poke callback on the primary partner.<18> In this case,
	// the parameters to the Poke call MUST be calculated from the incoming parameters and
	// the secondary partner's local name object; specifically, the pszCalleeUuid parameter
	// MUST be set to the value of the pszUuidString parameter; the pszHostName parameter
	// MUST be the Hostname field of the secondary partner's local name object; and the
	// pszUuidString parameter MUST be the string form of the CID field of the secondary
	// partner's local name object. The secondary partner MAY return from the Poke method
	// before this call has completed.
	//
	// When Poke is invoked on a primary partner, the primary partner MUST construct a name
	// object using the host name specified in the pszHostName parameter, the contact identifier
	// (CID) specified in the pszUuidString parameter, and the RPC protocols specified in
	// the grbitComProtocols field of the BIND_INFO_BLOB structure.
	//
	// The primary partner MUST use this name object to check whether or not an existing
	// session with a matching name object already exists in the session table.
	//
	// If an existing session is found, the primary partner MUST check the State field of
	// the session object.
	//
	// * If the value is set to Connecting, the existing session will be used during the
	// rest of the call.
	//
	// * Otherwise, the primary partner MUST return an implementation-specific error code.
	// <19> ( 2f0b4979-92b7-46f5-9e94-81531e68f3fe#Appendix_A_19 )
	//
	// If an existing session is not found, a new session object MUST be created and added
	// to the session table. The new session object MUST be initialized with the created
	// name object. An RPC binding handle to the secondary partner MUST be created and stored
	// in the session object. For binding handles, see [C706]. The State field MUST be set
	// to Connecting.
	//
	// At this point, the primary partner does not have to wait until the entire process
	// is completed. It SHOULD return success from the method, while it continues to perform
	// the following actions.<20>
	//
	// After identifying a valid existing session or initializing a new session object and
	// adding it to the session table, the primary partner MUST attempt to call either the
	// BuildContextW method or the BuildContext method on the secondary partner with the
	// RPC binding handle stored in the session object. For details on making BuildContext
	// calls to a partner, see section 3.3.4.2 and section 3.4.6.1.1.
	//
	// To determine whether the secondary partner supports BuildContextW, the primary partner
	// calls BuildContextW on the secondary partner and waits for a return value.
	//
	// If the secondary partner does not support the BuildContextW method, the primary partner
	// MUST call the BuildContext method.
	//
	// If the secondary partner does support the BuildContextW method, the primary partner
	// MUST NOT call the BuildContext method. During this call, the secondary partner will
	// make a nested synchronous callback to the primary partner to complete the session
	// establishment. See section 3.4.6.1.1.
	//
	// If the call completes successfully, the primary partner MUST examine the State field
	// of the session object; if the value is "Confirming Connection", it MUST set the state
	// of the session object to Active and cancel the Session Setup timer associated with
	// that session object.
	//
	// If the call completes unsuccessfully, the primary partner SHOULD behave according
	// to the error code that was returned:
	//
	// * If the error code is 0x80000712 (E_CM_VERSION_SET_NOTSUPPORTED), or 0x800000173
	// (E_CM_S_PROTOCOL_NOT_SUPPORTED), or it retried the nested call for more than the
	// number of times specified in the *Session Setup Retry Count* ADM element, or if the
	// State field of the session object is not "Confirming Connection", the primary partner
	// MUST remove the session object from the session table and clean it up. For instructions
	// on cleaning up a session object, see section 3.2.1.3 ( ff2a0bdc-9952-4357-971e-9e659c8824c8
	// ).
	//
	// * If the error code is ox800000123 (E_CM_SERVER_NOT_READY) or 0x000006BB (RPC_S_SERVER_TOO_BUSY),
	// or any other implementation-specific error code, the primary partner SHOULD retry
	// the call for the number of times specified in the *Session Setup Retry Count* ADM
	// element.
	Poke(context.Context, *PokeRequest) (*PokeResponse, error)

	// The BuildContext method is invoked by either a primary partner or a secondary partner.
	// When invoked by a primary partner, the BuildContext method requests that the secondary
	// partner begin the next step of establishing a session. When invoked by a secondary
	// partner, the BuildContext method requests that the primary partner complete the establishment
	// of the session.
	//
	// Return Values: This method MUST return zero (0x00000000) on success. On failure,
	// it MUST return either one of the values described in the following table of return
	// values or an implementation-specific HRESULT. A client MUST NOT depend on implementation-specific
	// failure HRESULT values. For more information about how the client SHOULD behave based
	// on the possible return values, see section 3.4.6.1.1. Standard errors are defined
	// in [MS-ERREF] section 2.2.
	//
	// Standard errors are defined in [MS-ERREF] section 4.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_STATUS                  | The return value indicates success.                                              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000172 E_CM_VERSION_SET_NOTSUPPORTED | The return value indicates that the callee partner does not support the caller’s |
	//	|                                          | BindVersionSet parameter and will not execute the requested operation.           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000124 E_CM_S_TIMEDOUT               | The return value indicates that the callee timed out while waiting for the       |
	//	|                                          | caller to complete the bind. This is returned by a secondary partner to a        |
	//	|                                          | primary partner if the primary partner does not return from the secondary        |
	//	|                                          | partner's call to BuildContext within half of the Session Setup Timer (section   |
	//	|                                          | 3.2.2.1) interval.                                                               |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006BB RPC_S_SERVER_TOO_BUSY         | The return value indicates that the partner is too busy to complete this         |
	//	|                                          | operation. For more information, see [MS-RPCE] section 3.1.1.5.5.                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000173 E_CM_S_PROTOCOL_NOT_SUPPORTED | The return value indicates that none of the protocols described in the rguchBlob |
	//	|                                          | parameter are supported by the partner.                                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                  | The return value indicates that one of the specified arguments is invalid.       |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The following table of return values describes the possible errors that SHOULD be
	// returned by this method.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000120 E_CM_SESSION_DOWN     | In a scenario where the value of the sRank parameter is SRANK_SECONDARY, if      |
	//	|                                  | BuildContext is called and an existing session object is not found, the call     |
	//	|                                  | SHOULD return this value.<21>                                                    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000123 E_CM_SERVER_NOT_READY | The session object is not in the Connecting state.<22>                           |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 1. For more information, see [C706].
	//
	// This method has different effects depending on the value of the sRank parameter.
	//
	// For the structure and sequence of data on the wire, see [C706] Transfer Syntax Network
	// Data Representation (NDR) topics.
	BuildContext(context.Context, *BuildContextRequest) (*BuildContextResponse, error)

	// The NegotiateResources method is invoked by one partner to request that the other
	// partner allocate resources for future use.
	//
	// Return Values: This method MUST return zero (0x00000000) on success. On failure,
	// it MUST return either one of the values described in the following table of return
	// values or an implementation-specific HRESULT. A client MUST NOT depend on implementation-specific
	// failure HRESULT values. For more information about how the client SHOULD behave based
	// on the possible return values, see section 3.4.6.4. Standard errors are defined in
	// [MS-ERREF] section 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_STATUS        | The return value indicates success.                                              |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000127 E_CM_OUTOFRESOURCES | The server was unable to allocate the resources requested and will continue to   |
	//	|                                | operate with the current set of resources.                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The following table of return values describes the possible errors that SHOULD be
	// returned by this method.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This value is returned in the following scenarios: If the resource type that was |
	//	|                                  | passed in the resourceType parameter is not a valid resource. If the value of    |
	//	|                                  | the dwcRequested parameter is not between 1 and 1000.                            |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000123 E_CM_SERVER_NOT_READY | The session object is not in the Active state.                                   |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 2. See [C706].
	//
	// For the structure and sequence of data on the wire, see [C706] Transfer Syntax Network
	// Data Representation (NDR) topics.
	NegotiateResources(context.Context, *NegotiateResourcesRequest) (*NegotiateResourcesResponse, error)

	// The SendReceive method is invoked by one partner to transmit messages to the other
	// partner. Both the primary and the secondary participants have the option to call
	// this method multiple times after a session has been established between them.
	//
	// Return Values: This method MUST return zero (0x00000000) on success. On failure,
	// it MUST return either one of the values described in the following table of return
	// values or an implementation-specific HRESULT. A client MUST NOT depend on implementation-specific
	// failure HRESULT values. For more information about how the client SHOULD behave based
	// on the possible return values, see section 3.4.6.4. Standard errors are defined in
	// [MS-ERREF] section 2.2.
	//
	//	+-------------------------+-------------------------------------+
	//	|         RETURN          |                                     |
	//	|       VALUE/CODE        |             DESCRIPTION             |
	//	|                         |                                     |
	//	+-------------------------+-------------------------------------+
	//	+-------------------------+-------------------------------------+
	//	| 0x00000000 ERROR_STATUS | The return value indicates success. |
	//	+-------------------------+-------------------------------------+
	//
	// The table below describes the possible errors that SHOULD be returned by this method.
	//
	//	+----------------------------------+---------------------------------------------------------------------+
	//	|              RETURN              |                                                                     |
	//	|            VALUE/CODE            |                             DESCRIPTION                             |
	//	|                                  |                                                                     |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0x80000119 E_CM_TEARING_DOWN     | The session object is in the Requesting Teardown or Teardown state. |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0x80000123 E_CM_SERVER_NOT_READY | The session object is not in the Active state.                      |
	//	+----------------------------------+---------------------------------------------------------------------+
	//
	// The opnum field value for this method is 3, as specified in [C706].
	//
	// For the structure and sequence of data on the wire, see [C706] section 14.
	SendReceive(context.Context, *SendReceiveRequest) (*SendReceiveResponse, error)

	// The TearDownContext method is invoked by either a primary partner or a secondary
	// partner. When invoked by a primary partner, the TearDownContext method requests that
	// the secondary partner begin the next step of tearing down a session. When invoked
	// by a secondary partner, the TearDownContext method requests that the primary partner
	// complete the teardown of the session. The Microsoft Interface Definition Language
	// (MIDL) syntax of the method is as follows.
	//
	// Return Values: This method MUST return zero (0x00000000) on success. On failure,
	// it MUST return an implementation-specific HRESULT. A client MUST NOT depend on implementation-specific
	// failure HRESULT values. From an over-the-wire communication point of view, the client
	// MUST implement only a behavior for the case when the call succeeds and another behavior
	// for the case when the call does not succeed, (see section 3.4.6.2). Standard errors
	// are defined in [MS-ERREF] section 2.2. A client MUST NOT exhibit behavior observable
	// on the wire that is dependent on implementation-specific failure HRESULT values.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_STATUS | The return value indicates success.                                              |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This value MAY be returned when an invalid sRank value is passed as a            |
	//	|                         | parameter.<27>                                                                   |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80004005 E_FAIL       | This return value indicates that the session failed to tear down within the      |
	//	|                         | interval specified by the Session Teardown Timer (section 3.2.2.2).              |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//
	// Thereafter, the method has a different effect depending on the value of the sRank
	// parameter and the value of the teardownType parameter.
	TearDownContext(context.Context, *TearDownContextRequest) (*TearDownContextResponse, error)

	// The BeginTearDown method is invoked by a secondary partner to request that a primary
	// partner begin session teardown.
	//
	// Return Values: This method MUST return zero (0x00000000) on success. On failure,
	// it MUST return an implementation-specific HRESULT. A client MUST NOT depend on implementation-specific
	// failure HRESULT values. From an over-the-wire communication point of view, the client
	// MUST implement only a behavior for the case when the call succeeds and another behavior
	// for the case when the call does not succeed, (see section 3.4.6.2). Standard errors
	// are defined in [MS-ERREF] section 2.2.
	//
	//	+-------------------------+-------------------------------------+
	//	|         RETURN          |                                     |
	//	|       VALUE/CODE        |             DESCRIPTION             |
	//	|                         |                                     |
	//	+-------------------------+-------------------------------------+
	//	+-------------------------+-------------------------------------+
	//	| 0x00000000 ERROR_STATUS | The return value indicates success. |
	//	+-------------------------+-------------------------------------+
	//
	// BeginTearDown MUST NOT be invoked on a secondary partner.
	//
	// If the session object is in the Teardown state, the primary partner MUST immediately
	// return from the method with S_OK. Otherwise, the primary partner MUST set the state
	// of the session object associated with the context handle to Teardown and return S_OK
	// from the method. Also, it MUST start the Session Teardown timer associated with that
	// session object and attempt to call the TearDownContext method on the secondary partner.
	// The secondary partner SHOULD choose to perform these actions asynchronously.
	BeginTearDown(context.Context, *BeginTearDownRequest) (*BeginTearDownResponse, error)

	// The PokeW method is equivalent in all ways to the Poke method except that its string
	// parameters are encoded in UTF-16.
	//
	// Return Values: This method MUST return zero (0x00000000) on success. On failure,
	// it MUST return an implementation-specific HRESULT. A client MUST NOT depend on implementation-specific
	// failure HRESULT values. From an over-the-wire communication point of view, the client
	// MUST implement only a behavior for the case when the call succeeds and another behavior
	// for the case when the call does not succeed, (see section 3.4.6.1.2). Standard errors
	// are defined in [MS-ERREF] section 2.2.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_STATUS                  | The return value indicates success.                                              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006D1 RPC_S_PROCNUM_OUT_OF_RANGE    | The return value indicates that the caller does not support this call.           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000123 E_CM_SERVER_NOT_READY         | The session object is not in the Connecting state.<28>                           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                  | The return value indicates that one of the specified arguments is invalid.<29>   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006BB RPC_S_SERVER_TOO_BUSY         | The return value indicates that the partner is too busy to complete this         |
	//	|                                          | operation. For more information, see [MS-RPCE] section 3.1.1.5.5.                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000173 E_CM_S_PROTOCOL_NOT_SUPPORTED | The return value indicates that none of the protocols described in the rguchBlob |
	//	|                                          | parameter is supported by the partner.                                           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//
	// When a partner calls PokeW on another partner, an error code of RPC_S_PROCNUM_OUT_OF_RANGE
	// means that the callee does not support PokeW.
	PokeW(context.Context, *PokeWRequest) (*PokeWResponse, error)

	// The BuildContextW method is equivalent in all ways to the BuildContext method, except
	// that its string parameters are encoded in UTF-16. The MIDL syntax of the method is
	// as follows.
	//
	// Return Values: This method MUST return zero (0x00000000) on success. On failure,
	// it MUST return either 0x80000172 (E_CM_VERSION_SET_NOTSUPPORTED) or an implementation-specific
	// HRESULT. A client SHOULD distinguish between 0x80000172 and other error codes, as
	// specified in sections 3.3.4.2.1 and 3.3.4.2.2, but MUST NOT depend on implementation-specific
	// failure HRESULT values. From an over-the-wire communication point of view, the client
	// MUST implement only behaviors for the errors described in the following table.
	//
	// Standard errors are defined in [MS-ERREF] section 4.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_STATUS                  | The return value indicates success.                                              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000172 E_CM_VERSION_SET_NOTSUPPORTED | The return value indicates that the callee partner does not support the caller's |
	//	|                                          | BindVersionSet parameter and will not execute the requested operation.           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006D1 RPC_S_PROCNUM_OUT_OF_RANGE    | The return value indicates that the caller does not support this call.           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000124 E_CM_S_TIMEDOUT               | The return value indicates that the callee timed out while waiting for the       |
	//	|                                          | caller to complete the bind. This value is returned by a secondary partner to    |
	//	|                                          | a primary partner if the primary partner does not return from the secondary      |
	//	|                                          | partner's call to BuildContext within half the amount of time specified in the   |
	//	|                                          | Session Setup Timer (section 3.2.2.1).                                           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006BB RPC_S_SERVER_TOO_BUSY         | The return value indicates that the partner is too busy to complete this         |
	//	|                                          | operation. For more information, see [MS-RPCE] section 3.1.1.5.5.                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000173 E_CM_S_PROTOCOL_NOT_SUPPORTED | The return value indicates that none of the protocols described in the rguchBlob |
	//	|                                          | parameter is supported by the partner.                                           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                  | The return value indicates that one of the specified arguments is invalid.       |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The following table describes the possible implementation-specific errors that SHOULD
	// be returned by this method.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000120 E_CM_SESSION_DOWN     | In a scenario where the value of the sRank parameter is SRANK_SECONDARY, if      |
	//	|                                  | BuildContextW is called and an existing session object is not found, the call    |
	//	|                                  | SHOULD return this value.<30>                                                    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000123 E_CM_SERVER_NOT_READY | The session object is not in the Connecting state.<31>                           |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//
	// When a partner calls BuildContextW on another partner, an error code of RPC_S_PROCNUM_OUT_OF_RANGE
	// means that the callee does not support BuildContextW.
	BuildContextW(context.Context, *BuildContextWRequest) (*BuildContextWResponse, error)
}

func RegisterIxnRemoteServer(conn dcerpc.Conn, o IxnRemoteServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIxnRemoteServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IxnRemoteSyntaxV1_0))...)
}

func NewIxnRemoteServerHandle(o IxnRemoteServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IxnRemoteServerHandle(ctx, o, opNum, r)
	}
}

func IxnRemoteServerHandle(ctx context.Context, o IxnRemoteServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Poke
		op := &xxx_PokeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PokeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Poke(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // BuildContext
		op := &xxx_BuildContextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BuildContextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BuildContext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // NegotiateResources
		op := &xxx_NegotiateResourcesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NegotiateResourcesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NegotiateResources(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // SendReceive
		op := &xxx_SendReceiveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendReceiveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SendReceive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // TearDownContext
		op := &xxx_TearDownContextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &TearDownContextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.TearDownContext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // BeginTearDown
		op := &xxx_BeginTearDownOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BeginTearDownRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BeginTearDown(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // PokeW
		op := &xxx_PokeWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PokeWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PokeW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // BuildContextW
		op := &xxx_BuildContextWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BuildContextWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BuildContextW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IXnRemote
type UnimplementedIxnRemoteServer struct {
}

func (UnimplementedIxnRemoteServer) Poke(context.Context, *PokeRequest) (*PokeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIxnRemoteServer) BuildContext(context.Context, *BuildContextRequest) (*BuildContextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIxnRemoteServer) NegotiateResources(context.Context, *NegotiateResourcesRequest) (*NegotiateResourcesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIxnRemoteServer) SendReceive(context.Context, *SendReceiveRequest) (*SendReceiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIxnRemoteServer) TearDownContext(context.Context, *TearDownContextRequest) (*TearDownContextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIxnRemoteServer) BeginTearDown(context.Context, *BeginTearDownRequest) (*BeginTearDownResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIxnRemoteServer) PokeW(context.Context, *PokeWRequest) (*PokeWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIxnRemoteServer) BuildContextW(context.Context, *BuildContextWRequest) (*BuildContextWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IxnRemoteServer = (*UnimplementedIxnRemoteServer)(nil)
