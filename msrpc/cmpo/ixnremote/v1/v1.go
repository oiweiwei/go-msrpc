package ixnremote

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
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
)

var (
	// import guard
	GoPackage = "cmpo"
)

var (
	// Syntax UUID
	IxnRemoteSyntaxUUID = &uuid.UUID{TimeLow: 0x906b0ce0, TimeMid: 0xc70b, TimeHiAndVersion: 0x1067, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x17, Node: [6]uint8{0x0, 0xdd, 0x1, 0x6, 0x62, 0xda}}
	// Syntax ID
	IxnRemoteSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: IxnRemoteSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// IXnRemote interface.
type IxnRemoteClient interface {

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
	Poke(context.Context, *PokeRequest, ...dcerpc.CallOption) (*PokeResponse, error)

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
	BuildContext(context.Context, *BuildContextRequest, ...dcerpc.CallOption) (*BuildContextResponse, error)

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
	NegotiateResources(context.Context, *NegotiateResourcesRequest, ...dcerpc.CallOption) (*NegotiateResourcesResponse, error)

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
	SendReceive(context.Context, *SendReceiveRequest, ...dcerpc.CallOption) (*SendReceiveResponse, error)

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
	TearDownContext(context.Context, *TearDownContextRequest, ...dcerpc.CallOption) (*TearDownContextResponse, error)

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
	BeginTearDown(context.Context, *BeginTearDownRequest, ...dcerpc.CallOption) (*BeginTearDownResponse, error)

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
	PokeW(context.Context, *PokeWRequest, ...dcerpc.CallOption) (*PokeWResponse, error)

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
	BuildContextW(context.Context, *BuildContextWRequest, ...dcerpc.CallOption) (*BuildContextWResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// MaxComputerNameLength represents the MAX_COMPUTERNAME_LENGTH RPC constant
var MaxComputerNameLength = 15

// GUIDLength represents the GUID_LENGTH RPC constant
var GUIDLength = 37

// TeardownType type represents TEARDOWN_TYPE RPC enumeration.
type TeardownType uint16

var (
	TeardownTypeForce   TeardownType = 0
	TeardownTypeProblem TeardownType = 2
)

func (o TeardownType) String() string {
	switch o {
	case TeardownTypeForce:
		return "TeardownTypeForce"
	case TeardownTypeProblem:
		return "TeardownTypeProblem"
	}
	return "Invalid"
}

// SessionRank type represents SESSION_RANK RPC enumeration.
type SessionRank uint16

var (
	SessionRankSrankPrimary   SessionRank = 1
	SessionRankSrankSecondary SessionRank = 2
)

func (o SessionRank) String() string {
	switch o {
	case SessionRankSrankPrimary:
		return "SessionRankSrankPrimary"
	case SessionRankSrankSecondary:
		return "SessionRankSrankSecondary"
	}
	return "Invalid"
}

// ResourceType type represents RESOURCE_TYPE RPC enumeration.
type ResourceType uint16

var (
	ResourceTypeConnections ResourceType = 0
)

func (o ResourceType) String() string {
	switch o {
	case ResourceTypeConnections:
		return "ResourceTypeConnections"
	}
	return "Invalid"
}

// BindVersionSet structure represents BIND_VERSION_SET RPC structure.
//
// The BIND_VERSION_SET structure holds three sets of version range values that specify
// the version ranges supported by a partner for three protocols: this protocol, MSDTC
// Connection Manager: OleTx Transports Protocol, and two other protocols that are layered
// on top of this protocol. This is because MSDTC Connection Manager: OleTx Transports
// Protocol is designed to be a transport protocol over which two other protocols are
// layered. For the rest of this specification, the protocol that is layered immediately
// on top of the MSDTC Connection Manager: OleTx Transports Protocol is referred to
// as the level-two protocol, and the protocol layered on top of the level-two protocol
// is the level-three protocol. The ranges of level-two version number values and level-three
// version number values are specific to the level-two protocol and level-three protocol,
// respectively.
type BindVersionSet struct {
	// dwMinLevelOne:  A 4-byte unsigned integer value containing the minimum supported
	// MSDTC Connection Manager: OleTx Transports Protocol version. dwMinLevelOne MUST be
	// less than or equal to dwMaxLevelOne.
	//
	// This field indicates whether the unsigned_char_t [C706] version of the Session creation
	// API calls (Poke/BuildContext) or the wchar_t [C706] version of the Session creation
	// API calls (PokeW/BuildContextW) are used. This field MUST be one of the following
	// values:
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The unsigned_char_t version of the Session creation API (Poke and BuildContext)  |
	//	|            | is used.                                                                         |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The wchar_t version of the Session creation API (PokeW and BuildContextW) is     |
	//	|            | used.                                                                            |
	//	+------------+----------------------------------------------------------------------------------+
	MinLevelOne uint32 `idl:"name:dwMinLevelOne" json:"min_level_one"`
	// dwMaxLevelOne:  A 4-byte unsigned integer value containing the maximum version supported
	// for a level-one session. dwMaxLevelOne MUST be greater than or equal to dwMinLevelOne.
	//
	// This field indicates whether the unsigned_char_t version of the Session creation
	// API calls (Poke/BuildContext) or the wchar_t version of the Session creation API
	// calls (PokeW/BuildContextW) are used. This field MUST be one of the following values:
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The unsigned_char_t version of the Session creation API (Poke and BuildContext)  |
	//	|            | is used.                                                                         |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The wchar_t version of the Session creation API (PokeW and BuildContextW) is     |
	//	|            | used.                                                                            |
	//	+------------+----------------------------------------------------------------------------------+
	MaxLevelOne uint32 `idl:"name:dwMaxLevelOne" json:"max_level_one"`
	// dwMinLevelTwo:  A 4-byte unsigned integer value containing the minimum version supported
	// for the level-two protocol session. The value for dwMinLevelTwo MUST be less than
	// or equal to dwMaxLevelTwo.
	MinLevelTwo uint32 `idl:"name:dwMinLevelTwo" json:"min_level_two"`
	// dwMaxLevelTwo:  A 4-byte unsigned integer value containing the maximum version supported
	// for the level-two protocol session. The value for dwMaxLevelTwo MUST be greater than
	// or equal to dwMinLevelTwo.
	MaxLevelTwo uint32 `idl:"name:dwMaxLevelTwo" json:"max_level_two"`
	// dwMinLevelThree:  A 4-byte unsigned integer value containing the minimum version
	// supported for the level-three protocol session. The value for dwMinLevelThree MUST
	// be less than or equal to dwMaxLevelThree.
	MinLevelThree uint32 `idl:"name:dwMinLevelThree" json:"min_level_three"`
	// dwMaxLevelThree:  A 4-byte unsigned integer value containing the maximum version
	// supported for the level-three protocol session. dwMaxLevelThree MUST be greater than
	// or equal to dwMinLevelThree.
	MaxLevelThree uint32 `idl:"name:dwMaxLevelThree" json:"max_level_three"`
}

func (o *BindVersionSet) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindVersionSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.MinLevelOne); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxLevelOne); err != nil {
		return err
	}
	if err := w.WriteData(o.MinLevelTwo); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxLevelTwo); err != nil {
		return err
	}
	if err := w.WriteData(o.MinLevelThree); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxLevelThree); err != nil {
		return err
	}
	return nil
}
func (o *BindVersionSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinLevelOne); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxLevelOne); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinLevelTwo); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxLevelTwo); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinLevelThree); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxLevelThree); err != nil {
		return err
	}
	return nil
}

// BoundVersionSet structure represents BOUND_VERSION_SET RPC structure.
//
// The BOUND_VERSION_SET is a structure containing the MSDTC Connection Manager: OleTx
// Transports Protocol version numbers that were successfully negotiated during a BuildContext
// call or a BuildContextW call.
type BoundVersionSet struct {
	// dwLevelOneAccepted:  A session level-one bind was successfully created.
	//
	// A 4-byte unsigned integer value containing the MSDTC Connection Manager: OleTx Transports
	// Protocol version that was negotiated with the partner and MUST be used in MSDTC Connection
	// Manager: OleTx Transports Protocol exchanges with the partner.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The unsigned_char_t version of the Session creation API (Poke and BuildContext)  |
	//	|            | is used.                                                                         |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The wchar_t version of the Session creation API (PokeW and BuildContextW) is     |
	//	|            | used.                                                                            |
	//	+------------+----------------------------------------------------------------------------------+
	LevelOneAccepted uint32 `idl:"name:dwLevelOneAccepted" json:"level_one_accepted"`
	// dwLevelTwoAccepted:  A 4-byte unsigned integer value containing the level-two protocol
	// version that was negotiated with the partner and MUST be used in level-two protocol
	// exchanges with the partner.
	LevelTwoAccepted uint32 `idl:"name:dwLevelTwoAccepted" json:"level_two_accepted"`
	// dwLevelThreeAccepted:  A 4-byte unsigned integer value containing the level-three
	// protocol version that was negotiated with the partner and MUST be used in level-three
	// protocol exchanges with the partner.
	LevelThreeAccepted uint32 `idl:"name:dwLevelThreeAccepted" json:"level_three_accepted"`
}

func (o *BoundVersionSet) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BoundVersionSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LevelOneAccepted); err != nil {
		return err
	}
	if err := w.WriteData(o.LevelTwoAccepted); err != nil {
		return err
	}
	if err := w.WriteData(o.LevelThreeAccepted); err != nil {
		return err
	}
	return nil
}
func (o *BoundVersionSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LevelOneAccepted); err != nil {
		return err
	}
	if err := w.ReadData(&o.LevelTwoAccepted); err != nil {
		return err
	}
	if err := w.ReadData(&o.LevelThreeAccepted); err != nil {
		return err
	}
	return nil
}

// BindInfoBlob structure represents BIND_INFO_BLOB RPC structure.
//
// The BIND_INFO_BLOB packet is a structure containing details on how to bind to a partner.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwcbThisStruct                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| grbitComProtocols                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type BindInfoBlob struct {
	// dwcbThisStruct (4 bytes): An unsigned 4-byte integer. The size of this structure
	// in bytes. This value MUST be set to 8.
	ThisStructureLength uint32 `idl:"name:dwcbThisStruct" json:"this_structure_length"`
	// grbitComProtocols (4 bytes): A COM_PROTOCOL bit field specifying the RPC protocol
	// sequences that the partner supports.
	COMProtocols uint32 `idl:"name:grbitComProtocols" json:"com_protocols"`
}

func (o *BindInfoBlob) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindInfoBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ThisStructureLength); err != nil {
		return err
	}
	if err := w.WriteData(o.COMProtocols); err != nil {
		return err
	}
	return nil
}
func (o *BindInfoBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ThisStructureLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.COMProtocols); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultIxnRemoteClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultIxnRemoteClient) Poke(ctx context.Context, in *PokeRequest, opts ...dcerpc.CallOption) (*PokeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PokeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIxnRemoteClient) BuildContext(ctx context.Context, in *BuildContextRequest, opts ...dcerpc.CallOption) (*BuildContextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BuildContextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIxnRemoteClient) NegotiateResources(ctx context.Context, in *NegotiateResourcesRequest, opts ...dcerpc.CallOption) (*NegotiateResourcesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NegotiateResourcesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIxnRemoteClient) SendReceive(ctx context.Context, in *SendReceiveRequest, opts ...dcerpc.CallOption) (*SendReceiveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SendReceiveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIxnRemoteClient) TearDownContext(ctx context.Context, in *TearDownContextRequest, opts ...dcerpc.CallOption) (*TearDownContextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &TearDownContextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIxnRemoteClient) BeginTearDown(ctx context.Context, in *BeginTearDownRequest, opts ...dcerpc.CallOption) (*BeginTearDownResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BeginTearDownResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIxnRemoteClient) PokeW(ctx context.Context, in *PokeWRequest, opts ...dcerpc.CallOption) (*PokeWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PokeWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIxnRemoteClient) BuildContextW(ctx context.Context, in *BuildContextWRequest, opts ...dcerpc.CallOption) (*BuildContextWResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BuildContextWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIxnRemoteClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIxnRemoteClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewIxnRemoteClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IxnRemoteClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IxnRemoteSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultIxnRemoteClient{cc: cc}, nil
}

// xxx_PokeOperation structure represents the Poke operation
type xxx_PokeOperation struct {
	Rank       SessionRank `idl:"name:sRank" json:"rank"`
	CalleeUUID string      `idl:"name:pszCalleeUuid;string" json:"callee_uuid"`
	HostName   string      `idl:"name:pszHostName;string" json:"host_name"`
	UUIDString string      `idl:"name:pszUuidString;string" json:"uuid_string"`
	SizeOfBlob uint32      `idl:"name:dwcbSizeOfBlob" json:"size_of_blob"`
	Blob       []byte      `idl:"name:rguchBlob;size_is:(dwcbSizeOfBlob)" json:"blob"`
	Return     int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_PokeOperation) OpNum() int { return 0 }

func (o *xxx_PokeOperation) OpName() string { return "/IXnRemote/v1/Poke" }

func (o *xxx_PokeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Blob != nil && o.SizeOfBlob == 0 {
		o.SizeOfBlob = uint32(len(o.Blob))
	}
	if len(o.CalleeUUID) > int(37) {
		return fmt.Errorf("CalleeUUID is out of range")
	}
	if len(o.HostName) > int(16) {
		return fmt.Errorf("HostName is out of range")
	}
	if len(o.UUIDString) > int(37) {
		return fmt.Errorf("UUIDString is out of range")
	}
	if o.SizeOfBlob < uint32(8) || o.SizeOfBlob > uint32(8) {
		return fmt.Errorf("SizeOfBlob is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PokeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// sRank {in} (1:{alias=SESSION_RANK}(enum))
	{
		if err := w.WriteEnum(uint16(o.Rank)); err != nil {
			return err
		}
	}
	// pszCalleeUuid {in} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.CalleeUUID); err != nil {
			return err
		}
	}
	// pszHostName {in} (1:{string, range=(1,16)}[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.HostName); err != nil {
			return err
		}
	}
	// pszUuidString {in} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.UUIDString); err != nil {
			return err
		}
	}
	// dwcbSizeOfBlob {in} (1:{range=(8,8), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeOfBlob); err != nil {
			return err
		}
	}
	// rguchBlob {in} (1:[dim:0,size_is=dwcbSizeOfBlob](uchar))
	{
		dimSize1 := uint64(o.SizeOfBlob)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Blob {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Blob[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Blob); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_PokeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// sRank {in} (1:{alias=SESSION_RANK}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Rank)); err != nil {
			return err
		}
	}
	// pszCalleeUuid {in} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.CalleeUUID); err != nil {
			return err
		}
	}
	// pszHostName {in} (1:{string, range=(1,16)}[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.HostName); err != nil {
			return err
		}
	}
	// pszUuidString {in} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.UUIDString); err != nil {
			return err
		}
	}
	// dwcbSizeOfBlob {in} (1:{range=(8,8), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeOfBlob); err != nil {
			return err
		}
	}
	// rguchBlob {in} (1:[dim:0,size_is=dwcbSizeOfBlob](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Blob", sizeInfo[0])
		}
		o.Blob = make([]byte, sizeInfo[0])
		for i1 := range o.Blob {
			i1 := i1
			if err := w.ReadData(&o.Blob[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_PokeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PokeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PokeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PokeRequest structure represents the Poke operation request
type PokeRequest struct {
	// sRank: The session rank of the partner making the call. This parameter MUST be set
	// to 0x02 (SRANK_SECONDARY).
	//
	//	+----------------------+------------------------------------------+
	//	|                      |                                          |
	//	|        VALUE         |                 MEANING                  |
	//	|                      |                                          |
	//	+----------------------+------------------------------------------+
	//	+----------------------+------------------------------------------+
	//	| SRANK_SECONDARY 0x02 | The caller is the secondary participant. |
	//	+----------------------+------------------------------------------+
	Rank SessionRank `idl:"name:sRank" json:"rank"`
	// pszCalleeUuid: A string containing the primary partner's contact identifier (CID)
	// in the form of a GUID. The contact identifier (CID) MUST match the CID in the primary
	// partner's local name object and MUST be formatted into a string.
	CalleeUUID string `idl:"name:pszCalleeUuid;string" json:"callee_uuid"`
	// pszHostName: The string form of the caller's host name. This host name identifies
	// the machine on which the caller's instance of the MSDTC Connection Manager: OleTx
	// Transports Protocol is running. This value is used by the primary participant to
	// establish the RPC binding handle for its subsequent call to BuildContext. This MUST
	// be a NetBIOS name. For NetBIOS, see [NETBEUI], [RFC1001], and [RFC1002].
	HostName string `idl:"name:pszHostName;string" json:"host_name"`
	// pszUuidString: The string form of the caller's contact identifier (CID) in the form
	// of a GUID. This contact identifier (CID) identifies the caller's instance of the
	// MSDTC Connection Manager: OleTx Transports Protocol. It MUST match the CID in the
	// caller's local name object, and MUST be formatted into a string. This value is used
	// by the primary participant to establish the RPC binding handle for its subsequent
	// call to BuildContext.
	UUIDString string `idl:"name:pszUuidString;string" json:"uuid_string"`
	// dwcbSizeOfBlob: The count, in bytes, of the size of the binding info structure. This
	// parameter MUST be set to 0x00000008.
	SizeOfBlob uint32 `idl:"name:dwcbSizeOfBlob" json:"size_of_blob"`
	// rguchBlob: A byte array containing a BIND_INFO_BLOB structure specifying the transport
	// protocols supported. This information is used to build the RPC binding for the reverse
	// connection.
	Blob []byte `idl:"name:rguchBlob;size_is:(dwcbSizeOfBlob)" json:"blob"`
}

func (o *PokeRequest) xxx_ToOp(ctx context.Context, op *xxx_PokeOperation) *xxx_PokeOperation {
	if op == nil {
		op = &xxx_PokeOperation{}
	}
	if o == nil {
		return op
	}
	o.Rank = op.Rank
	o.CalleeUUID = op.CalleeUUID
	o.HostName = op.HostName
	o.UUIDString = op.UUIDString
	o.SizeOfBlob = op.SizeOfBlob
	o.Blob = op.Blob
	return op
}

func (o *PokeRequest) xxx_FromOp(ctx context.Context, op *xxx_PokeOperation) {
	if o == nil {
		return
	}
	o.Rank = op.Rank
	o.CalleeUUID = op.CalleeUUID
	o.HostName = op.HostName
	o.UUIDString = op.UUIDString
	o.SizeOfBlob = op.SizeOfBlob
	o.Blob = op.Blob
}
func (o *PokeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PokeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PokeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PokeResponse structure represents the Poke operation response
type PokeResponse struct {
	// Return: The Poke return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PokeResponse) xxx_ToOp(ctx context.Context, op *xxx_PokeOperation) *xxx_PokeOperation {
	if op == nil {
		op = &xxx_PokeOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *PokeResponse) xxx_FromOp(ctx context.Context, op *xxx_PokeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PokeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PokeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PokeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BuildContextOperation structure represents the BuildContext operation
type xxx_BuildContextOperation struct {
	Rank            SessionRank             `idl:"name:sRank" json:"rank"`
	BindVersionSet  *BindVersionSet         `idl:"name:BindVersionSet" json:"bind_version_set"`
	CalleeUUID      string                  `idl:"name:pszCalleeUuid;string" json:"callee_uuid"`
	HostName        string                  `idl:"name:pszHostName;string" json:"host_name"`
	UUIDString      string                  `idl:"name:pszUuidString;string" json:"uuid_string"`
	GUIDIn          string                  `idl:"name:pszGuidIn;string" json:"guid_in"`
	GUIDOut         string                  `idl:"name:pszGuidOut;string" json:"guid_out"`
	BoundVersionSet *BoundVersionSet        `idl:"name:pBoundVersionSet" json:"bound_version_set"`
	SizeOfBlob      uint32                  `idl:"name:dwcbSizeOfBlob" json:"size_of_blob"`
	Blob            []byte                  `idl:"name:rguchBlob;size_is:(dwcbSizeOfBlob)" json:"blob"`
	Handle          *dcetypes.ContextHandle `idl:"name:ppHandle" json:"handle"`
	Return          int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_BuildContextOperation) OpNum() int { return 1 }

func (o *xxx_BuildContextOperation) OpName() string { return "/IXnRemote/v1/BuildContext" }

func (o *xxx_BuildContextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Blob != nil && o.SizeOfBlob == 0 {
		o.SizeOfBlob = uint32(len(o.Blob))
	}
	if len(o.CalleeUUID) > int(37) {
		return fmt.Errorf("CalleeUUID is out of range")
	}
	if len(o.HostName) > int(16) {
		return fmt.Errorf("HostName is out of range")
	}
	if len(o.UUIDString) > int(37) {
		return fmt.Errorf("UUIDString is out of range")
	}
	if len(o.GUIDIn) > int(37) {
		return fmt.Errorf("GUIDIn is out of range")
	}
	if len(o.GUIDOut) > int(37) {
		return fmt.Errorf("GUIDOut is out of range")
	}
	if o.SizeOfBlob < uint32(8) || o.SizeOfBlob > uint32(8) {
		return fmt.Errorf("SizeOfBlob is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BuildContextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// sRank {in} (1:{alias=SESSION_RANK}(enum))
	{
		if err := w.WriteEnum(uint16(o.Rank)); err != nil {
			return err
		}
	}
	// BindVersionSet {in} (1:{alias=BIND_VERSION_SET}(struct))
	{
		if o.BindVersionSet != nil {
			if err := o.BindVersionSet.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&BindVersionSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pszCalleeUuid {in} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.CalleeUUID); err != nil {
			return err
		}
	}
	// pszHostName {in} (1:{string, range=(1,16)}[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.HostName); err != nil {
			return err
		}
	}
	// pszUuidString {in} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.UUIDString); err != nil {
			return err
		}
	}
	// pszGuidIn {in} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.GUIDIn); err != nil {
			return err
		}
	}
	// pszGuidOut {in, out} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.GUIDOut); err != nil {
			return err
		}
	}
	// pBoundVersionSet {in, out} (1:{pointer=ref}*(1))(2:{alias=BOUND_VERSION_SET}(struct))
	{
		if o.BoundVersionSet != nil {
			if err := o.BoundVersionSet.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&BoundVersionSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwcbSizeOfBlob {in} (1:{range=(8,8), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeOfBlob); err != nil {
			return err
		}
	}
	// rguchBlob {in} (1:[dim:0,size_is=dwcbSizeOfBlob](uchar))
	{
		dimSize1 := uint64(o.SizeOfBlob)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Blob {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Blob[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Blob); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_BuildContextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// sRank {in} (1:{alias=SESSION_RANK}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Rank)); err != nil {
			return err
		}
	}
	// BindVersionSet {in} (1:{alias=BIND_VERSION_SET}(struct))
	{
		if o.BindVersionSet == nil {
			o.BindVersionSet = &BindVersionSet{}
		}
		if err := o.BindVersionSet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pszCalleeUuid {in} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.CalleeUUID); err != nil {
			return err
		}
	}
	// pszHostName {in} (1:{string, range=(1,16)}[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.HostName); err != nil {
			return err
		}
	}
	// pszUuidString {in} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.UUIDString); err != nil {
			return err
		}
	}
	// pszGuidIn {in} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.GUIDIn); err != nil {
			return err
		}
	}
	// pszGuidOut {in, out} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.GUIDOut); err != nil {
			return err
		}
	}
	// pBoundVersionSet {in, out} (1:{pointer=ref}*(1))(2:{alias=BOUND_VERSION_SET}(struct))
	{
		if o.BoundVersionSet == nil {
			o.BoundVersionSet = &BoundVersionSet{}
		}
		if err := o.BoundVersionSet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwcbSizeOfBlob {in} (1:{range=(8,8), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeOfBlob); err != nil {
			return err
		}
	}
	// rguchBlob {in} (1:[dim:0,size_is=dwcbSizeOfBlob](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Blob", sizeInfo[0])
		}
		o.Blob = make([]byte, sizeInfo[0])
		for i1 := range o.Blob {
			i1 := i1
			if err := w.ReadData(&o.Blob[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_BuildContextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if len(o.GUIDOut) > int(37) {
		return fmt.Errorf("GUIDOut is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BuildContextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pszGuidOut {in, out} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.GUIDOut); err != nil {
			return err
		}
	}
	// pBoundVersionSet {in, out} (1:{pointer=ref}*(1))(2:{alias=BOUND_VERSION_SET}(struct))
	{
		if o.BoundVersionSet != nil {
			if err := o.BoundVersionSet.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&BoundVersionSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppHandle {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_BuildContextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pszGuidOut {in, out} (1:{string, range=(37,37)}[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.GUIDOut); err != nil {
			return err
		}
	}
	// pBoundVersionSet {in, out} (1:{pointer=ref}*(1))(2:{alias=BOUND_VERSION_SET}(struct))
	{
		if o.BoundVersionSet == nil {
			o.BoundVersionSet = &BoundVersionSet{}
		}
		if err := o.BoundVersionSet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppHandle {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &dcetypes.ContextHandle{}
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

// BuildContextRequest structure represents the BuildContext operation request
type BuildContextRequest struct {
	// sRank: The session rank of the partner making the call. It MUST be one of the following
	// values.
	//
	//	+-------------------+------------------------------------------------------+
	//	|                   |                                                      |
	//	|       VALUE       |                       MEANING                        |
	//	|                   |                                                      |
	//	+-------------------+------------------------------------------------------+
	//	+-------------------+------------------------------------------------------+
	//	| SRANK_PRIMARY 1   | The caller is the primary partner in this session.   |
	//	+-------------------+------------------------------------------------------+
	//	| SRANK_SECONDARY 2 | The caller is the secondary partner in this session. |
	//	+-------------------+------------------------------------------------------+
	Rank SessionRank `idl:"name:sRank" json:"rank"`
	// BindVersionSet: A BIND_VERSION_SET structure that contains the minimum and maximum
	// versions supported by the partner, as specified by the Minimum Level 1 Version Number,
	// Maximum Level 1 Version Number, Minimum Level 2 Version Number, Maximum Level 2 Version
	// Number, Minimum Level 3 Version Number, and Maximum Level 3 Version Number ADM elements
	// (see section 3.2.1.1).
	BindVersionSet *BindVersionSet `idl:"name:BindVersionSet" json:"bind_version_set"`
	// pszCalleeUuid: A string containing the callee's contact identifier (CID) in the form
	// of a GUID. The contact identifier (CID) MUST match the contact identifier (CID) in
	// the callee's local name object and MUST be formatted into a string.
	CalleeUUID string `idl:"name:pszCalleeUuid;string" json:"callee_uuid"`
	// pszHostName: The string form of the caller's host name. This host name identifies
	// the machine in which the caller's instance of the MSDTC Connection Manager: OleTx
	// Transports Protocol is running. This MUST be a NetBIOS name. For NetBIOS, see [NETBEUI],
	// [RFC1001], and [RFC1002].
	HostName string `idl:"name:pszHostName;string" json:"host_name"`
	// pszUuidString: The string form of the caller's contact identifier (CID) in the form
	// of a GUID. This CID identifies the caller's instance of the MSDTC Connection Manager:
	// OleTx Transports Protocol. It MUST match the contact identifier (CID) in the caller's
	// local name object and MUST be formatted into a string.
	UUIDString string `idl:"name:pszUuidString;string" json:"uuid_string"`
	// pszGuidIn: A string form of a GUID that represents a unique identifier for this bind
	// attempt. The GUID MUST be formatted as a string.
	GUIDIn string `idl:"name:pszGuidIn;string" json:"guid_in"`
	// pszGuidOut: A string form of a GUID that represents a global identifier for this
	// bind attempt. On input, the pszGuidOut parameter MUST be set to 00000000-0000-0000-0000-000000000000.
	// On return, if the bind attempt is ultimately successful, the pszGuidOut parameter
	// MUST be equal to the value of the pszGuidIn parameter. Otherwise, if the bind attempt
	// is ultimately unsuccessful, the pszGuidOut parameter MUST be set to 00000000-0000-0000-0000-000000000000
	// on return.
	GUIDOut string `idl:"name:pszGuidOut;string" json:"guid_out"`
	// pBoundVersionSet: A pointer to a BOUND_VERSION_SET structure. This structure is filled
	// in by the callee. If any error is returned, this structure MUST be filled with zeros
	// before returning. On successful completion, the caller receives a BOUND_VERSION_SET
	// on return.
	BoundVersionSet *BoundVersionSet `idl:"name:pBoundVersionSet" json:"bound_version_set"`
	// dwcbSizeOfBlob: The count in bytes of the size of the binding info structure. This
	// parameter MUST be set to the size of the BIND_INFO_BLOB, 8.
	SizeOfBlob uint32 `idl:"name:dwcbSizeOfBlob" json:"size_of_blob"`
	// rguchBlob: A byte array containing the BIND_INFO_BLOB structure specifying the supported
	// transport protocols. This information is used to build the RPC binding for the reverse
	// connection.
	Blob []byte `idl:"name:rguchBlob;size_is:(dwcbSizeOfBlob)" json:"blob"`
}

func (o *BuildContextRequest) xxx_ToOp(ctx context.Context, op *xxx_BuildContextOperation) *xxx_BuildContextOperation {
	if op == nil {
		op = &xxx_BuildContextOperation{}
	}
	if o == nil {
		return op
	}
	o.Rank = op.Rank
	o.BindVersionSet = op.BindVersionSet
	o.CalleeUUID = op.CalleeUUID
	o.HostName = op.HostName
	o.UUIDString = op.UUIDString
	o.GUIDIn = op.GUIDIn
	o.GUIDOut = op.GUIDOut
	o.BoundVersionSet = op.BoundVersionSet
	o.SizeOfBlob = op.SizeOfBlob
	o.Blob = op.Blob
	return op
}

func (o *BuildContextRequest) xxx_FromOp(ctx context.Context, op *xxx_BuildContextOperation) {
	if o == nil {
		return
	}
	o.Rank = op.Rank
	o.BindVersionSet = op.BindVersionSet
	o.CalleeUUID = op.CalleeUUID
	o.HostName = op.HostName
	o.UUIDString = op.UUIDString
	o.GUIDIn = op.GUIDIn
	o.GUIDOut = op.GUIDOut
	o.BoundVersionSet = op.BoundVersionSet
	o.SizeOfBlob = op.SizeOfBlob
	o.Blob = op.Blob
}
func (o *BuildContextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BuildContextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BuildContextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BuildContextResponse structure represents the BuildContext operation response
type BuildContextResponse struct {
	// pszGuidOut: A string form of a GUID that represents a global identifier for this
	// bind attempt. On input, the pszGuidOut parameter MUST be set to 00000000-0000-0000-0000-000000000000.
	// On return, if the bind attempt is ultimately successful, the pszGuidOut parameter
	// MUST be equal to the value of the pszGuidIn parameter. Otherwise, if the bind attempt
	// is ultimately unsuccessful, the pszGuidOut parameter MUST be set to 00000000-0000-0000-0000-000000000000
	// on return.
	GUIDOut string `idl:"name:pszGuidOut;string" json:"guid_out"`
	// pBoundVersionSet: A pointer to a BOUND_VERSION_SET structure. This structure is filled
	// in by the callee. If any error is returned, this structure MUST be filled with zeros
	// before returning. On successful completion, the caller receives a BOUND_VERSION_SET
	// on return.
	BoundVersionSet *BoundVersionSet `idl:"name:pBoundVersionSet" json:"bound_version_set"`
	// ppHandle: On successful return, an RPC context handle that correlates with the session
	// object created by (or referenced by) this method. For RPC context handles, see [C706].
	Handle *dcetypes.ContextHandle `idl:"name:ppHandle" json:"handle"`
	// Return: The BuildContext return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BuildContextResponse) xxx_ToOp(ctx context.Context, op *xxx_BuildContextOperation) *xxx_BuildContextOperation {
	if op == nil {
		op = &xxx_BuildContextOperation{}
	}
	if o == nil {
		return op
	}
	o.GUIDOut = op.GUIDOut
	o.BoundVersionSet = op.BoundVersionSet
	o.Handle = op.Handle
	o.Return = op.Return
	return op
}

func (o *BuildContextResponse) xxx_FromOp(ctx context.Context, op *xxx_BuildContextOperation) {
	if o == nil {
		return
	}
	o.GUIDOut = op.GUIDOut
	o.BoundVersionSet = op.BoundVersionSet
	o.Handle = op.Handle
	o.Return = op.Return
}
func (o *BuildContextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BuildContextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BuildContextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NegotiateResourcesOperation structure represents the NegotiateResources operation
type xxx_NegotiateResourcesOperation struct {
	Context        *dcetypes.ContextHandle `idl:"name:phContext" json:"context"`
	ResourceType   ResourceType            `idl:"name:resourceType" json:"resource_type"`
	RequestedCount uint32                  `idl:"name:dwcRequested" json:"requested_count"`
	AcceptedCount  uint32                  `idl:"name:pdwcAccepted" json:"accepted_count"`
	Return         int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_NegotiateResourcesOperation) OpNum() int { return 2 }

func (o *xxx_NegotiateResourcesOperation) OpName() string { return "/IXnRemote/v1/NegotiateResources" }

func (o *xxx_NegotiateResourcesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NegotiateResourcesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{alias=PCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// resourceType {in} (1:{alias=RESOURCE_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.ResourceType)); err != nil {
			return err
		}
	}
	// dwcRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedCount); err != nil {
			return err
		}
	}
	// pdwcAccepted {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AcceptedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NegotiateResourcesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{alias=PCONTEXT_HANDLE,pointer=ref}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &dcetypes.ContextHandle{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// resourceType {in} (1:{alias=RESOURCE_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ResourceType)); err != nil {
			return err
		}
	}
	// dwcRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedCount); err != nil {
			return err
		}
	}
	// pdwcAccepted {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AcceptedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NegotiateResourcesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NegotiateResourcesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwcAccepted {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AcceptedCount); err != nil {
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

func (o *xxx_NegotiateResourcesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwcAccepted {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AcceptedCount); err != nil {
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

// NegotiateResourcesRequest structure represents the NegotiateResources operation request
type NegotiateResourcesRequest struct {
	// phContext: An RPC context, returned by a call to BuildContext or BuildContextW, correlated
	// with a session object that is in the Active state. For context handles, see [C706].
	Context *dcetypes.ContextHandle `idl:"name:phContext" json:"context"`
	// resourceType: A RESOURCE_TYPE enumerated value indicating the resource type to be
	// negotiated.
	//
	//	+---------------------+------------------------------------------------+
	//	|                     |                                                |
	//	|        VALUE        |                    MEANING                     |
	//	|                     |                                                |
	//	+---------------------+------------------------------------------------+
	//	+---------------------+------------------------------------------------+
	//	| RT_CONNECTIONS 0x00 | The resource to be negotiated is a connection. |
	//	+---------------------+------------------------------------------------+
	ResourceType ResourceType `idl:"name:resourceType" json:"resource_type"`
	// dwcRequested: An unsigned 32-bit integer that specifies the number of resources to
	// allocate. This value MUST be greater than 0x00 and less than 1,000.
	RequestedCount uint32 `idl:"name:dwcRequested" json:"requested_count"`
	// pdwcAccepted: A pointer to an unsigned 32-bit integer that receives the number of
	// resources that were allocated on behalf of the caller. This value SHOULD be smaller
	// than the value of dwcRequested if the partner was incapable of allocating all of
	// the requested resources. On input, this value MUST be set to 0x00000000.
	AcceptedCount uint32 `idl:"name:pdwcAccepted" json:"accepted_count"`
}

func (o *NegotiateResourcesRequest) xxx_ToOp(ctx context.Context, op *xxx_NegotiateResourcesOperation) *xxx_NegotiateResourcesOperation {
	if op == nil {
		op = &xxx_NegotiateResourcesOperation{}
	}
	if o == nil {
		return op
	}
	o.Context = op.Context
	o.ResourceType = op.ResourceType
	o.RequestedCount = op.RequestedCount
	o.AcceptedCount = op.AcceptedCount
	return op
}

func (o *NegotiateResourcesRequest) xxx_FromOp(ctx context.Context, op *xxx_NegotiateResourcesOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.ResourceType = op.ResourceType
	o.RequestedCount = op.RequestedCount
	o.AcceptedCount = op.AcceptedCount
}
func (o *NegotiateResourcesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *NegotiateResourcesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NegotiateResourcesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NegotiateResourcesResponse structure represents the NegotiateResources operation response
type NegotiateResourcesResponse struct {
	// pdwcAccepted: A pointer to an unsigned 32-bit integer that receives the number of
	// resources that were allocated on behalf of the caller. This value SHOULD be smaller
	// than the value of dwcRequested if the partner was incapable of allocating all of
	// the requested resources. On input, this value MUST be set to 0x00000000.
	AcceptedCount uint32 `idl:"name:pdwcAccepted" json:"accepted_count"`
	// Return: The NegotiateResources return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *NegotiateResourcesResponse) xxx_ToOp(ctx context.Context, op *xxx_NegotiateResourcesOperation) *xxx_NegotiateResourcesOperation {
	if op == nil {
		op = &xxx_NegotiateResourcesOperation{}
	}
	if o == nil {
		return op
	}
	o.AcceptedCount = op.AcceptedCount
	o.Return = op.Return
	return op
}

func (o *NegotiateResourcesResponse) xxx_FromOp(ctx context.Context, op *xxx_NegotiateResourcesOperation) {
	if o == nil {
		return
	}
	o.AcceptedCount = op.AcceptedCount
	o.Return = op.Return
}
func (o *NegotiateResourcesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *NegotiateResourcesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NegotiateResourcesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SendReceiveOperation structure represents the SendReceive operation
type xxx_SendReceiveOperation struct {
	Context       *dcetypes.ContextHandle `idl:"name:phContext" json:"context"`
	MessagesCount uint32                  `idl:"name:dwcMessages" json:"messages_count"`
	SizeOfBoxCar  uint32                  `idl:"name:dwcbSizeOfBoxCar" json:"size_of_box_car"`
	BoxCar        []byte                  `idl:"name:rguchBoxCar;size_is:(dwcbSizeOfBoxCar)" json:"box_car"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SendReceiveOperation) OpNum() int { return 3 }

func (o *xxx_SendReceiveOperation) OpName() string { return "/IXnRemote/v1/SendReceive" }

func (o *xxx_SendReceiveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.BoxCar != nil && o.SizeOfBoxCar == 0 {
		o.SizeOfBoxCar = uint32(len(o.BoxCar))
	}
	if o.MessagesCount < uint32(1) || o.MessagesCount > uint32(4095) {
		return fmt.Errorf("MessagesCount is out of range")
	}
	if o.SizeOfBoxCar < uint32(40) || o.SizeOfBoxCar > uint32(81920) {
		return fmt.Errorf("SizeOfBoxCar is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendReceiveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{alias=PCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwcMessages {in} (1:{range=(1,4095), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MessagesCount); err != nil {
			return err
		}
	}
	// dwcbSizeOfBoxCar {in} (1:{range=(40,81920), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeOfBoxCar); err != nil {
			return err
		}
	}
	// rguchBoxCar {in} (1:[dim:0,size_is=dwcbSizeOfBoxCar](uchar))
	{
		dimSize1 := uint64(o.SizeOfBoxCar)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.BoxCar {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.BoxCar[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.BoxCar); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SendReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{alias=PCONTEXT_HANDLE,pointer=ref}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &dcetypes.ContextHandle{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwcMessages {in} (1:{range=(1,4095), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MessagesCount); err != nil {
			return err
		}
	}
	// dwcbSizeOfBoxCar {in} (1:{range=(40,81920), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeOfBoxCar); err != nil {
			return err
		}
	}
	// rguchBoxCar {in} (1:[dim:0,size_is=dwcbSizeOfBoxCar](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.BoxCar", sizeInfo[0])
		}
		o.BoxCar = make([]byte, sizeInfo[0])
		for i1 := range o.BoxCar {
			i1 := i1
			if err := w.ReadData(&o.BoxCar[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SendReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendReceiveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SendReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SendReceiveRequest structure represents the SendReceive operation request
type SendReceiveRequest struct {
	// phContext: An RPC context handle, returned by a call to BuildContext or BuildContextW,
	// correlated with a session object in the Active state. For context handles, see [C706].
	Context *dcetypes.ContextHandle `idl:"name:phContext" json:"context"`
	// dwcMessages: An unsigned 32-bit integer specifying the number of messages being sent.
	MessagesCount uint32 `idl:"name:dwcMessages" json:"messages_count"`
	// dwcbSizeOfBoxCar: Size in bytes of the box car specified by rguchBoxCar.
	SizeOfBoxCar uint32 `idl:"name:dwcbSizeOfBoxCar" json:"size_of_box_car"`
	// rguchBoxCar: An array of bytes that contains the messages being sent.
	BoxCar []byte `idl:"name:rguchBoxCar;size_is:(dwcbSizeOfBoxCar)" json:"box_car"`
}

func (o *SendReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_SendReceiveOperation) *xxx_SendReceiveOperation {
	if op == nil {
		op = &xxx_SendReceiveOperation{}
	}
	if o == nil {
		return op
	}
	o.Context = op.Context
	o.MessagesCount = op.MessagesCount
	o.SizeOfBoxCar = op.SizeOfBoxCar
	o.BoxCar = op.BoxCar
	return op
}

func (o *SendReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_SendReceiveOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.MessagesCount = op.MessagesCount
	o.SizeOfBoxCar = op.SizeOfBoxCar
	o.BoxCar = op.BoxCar
}
func (o *SendReceiveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SendReceiveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendReceiveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SendReceiveResponse structure represents the SendReceive operation response
type SendReceiveResponse struct {
	// Return: The SendReceive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SendReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_SendReceiveOperation) *xxx_SendReceiveOperation {
	if op == nil {
		op = &xxx_SendReceiveOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SendReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_SendReceiveOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SendReceiveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SendReceiveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendReceiveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_TearDownContextOperation structure represents the TearDownContext operation
type xxx_TearDownContextOperation struct {
	ContextHandle *dcetypes.ContextHandle `idl:"name:contextHandle" json:"context_handle"`
	Rank          SessionRank             `idl:"name:sRank" json:"rank"`
	TearDownType  TeardownType            `idl:"name:tearDownType" json:"tear_down_type"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_TearDownContextOperation) OpNum() int { return 4 }

func (o *xxx_TearDownContextOperation) OpName() string { return "/IXnRemote/v1/TearDownContext" }

func (o *xxx_TearDownContextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TearDownContextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// contextHandle {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// sRank {in} (1:{alias=SESSION_RANK}(enum))
	{
		if err := w.WriteEnum(uint16(o.Rank)); err != nil {
			return err
		}
	}
	// tearDownType {in} (1:{alias=TEARDOWN_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.TearDownType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TearDownContextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// contextHandle {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &dcetypes.ContextHandle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// sRank {in} (1:{alias=SESSION_RANK}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Rank)); err != nil {
			return err
		}
	}
	// tearDownType {in} (1:{alias=TEARDOWN_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.TearDownType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TearDownContextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TearDownContextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// contextHandle {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_TearDownContextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// contextHandle {in, out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &dcetypes.ContextHandle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
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

// TearDownContextRequest structure represents the TearDownContext operation request
type TearDownContextRequest struct {
	// contextHandle: An RPC context handle, returned by a call to BuildContext or BuildContextW,
	// is correlated with a session object that is in the Active state. After TearDownContext
	// is executed, on either success or failure requests, contextHandle will be set to
	// null. For context handles, see [C706].
	ContextHandle *dcetypes.ContextHandle `idl:"name:contextHandle" json:"context_handle"`
	// sRank: A SESSION_RANK enumerated value indicating whether the teardown request is
	// being made by a primary partner or secondary partner. The teardown request MUST be
	// sent from a primary partner only.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|                      |                                                                                  |
	//	|        VALUE         |                                     MEANING                                      |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| SRANK_PRIMARY 0x01   | The caller is the primary partner in this session. The callee MUST be a          |
	//	|                      | secondary partner in this session, and the caller MUST be a primary partner in   |
	//	|                      | this session.                                                                    |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| SRANK_SECONDARY 0x02 | The caller is the secondary partner in this session. The callee MUST be a        |
	//	|                      | primary partner in this session, and the caller MUST be a secondary partner in   |
	//	|                      | this session.                                                                    |
	//	+----------------------+----------------------------------------------------------------------------------+
	Rank SessionRank `idl:"name:sRank" json:"rank"`
	// tearDownType: The reason for tearing down the session. It MUST be one of the following
	// values.
	//
	//	+-----------------+---------------------------------------------------------------+
	//	|                 |                                                               |
	//	|      VALUE      |                            MEANING                            |
	//	|                 |                                                               |
	//	+-----------------+---------------------------------------------------------------+
	//	+-----------------+---------------------------------------------------------------+
	//	| TT_FORCE 0x00   | The session is being forcefully torn down.                    |
	//	+-----------------+---------------------------------------------------------------+
	//	| TT_PROBLEM 0x02 | The session is being torn down because an error has occurred. |
	//	+-----------------+---------------------------------------------------------------+
	TearDownType TeardownType `idl:"name:tearDownType" json:"tear_down_type"`
}

func (o *TearDownContextRequest) xxx_ToOp(ctx context.Context, op *xxx_TearDownContextOperation) *xxx_TearDownContextOperation {
	if op == nil {
		op = &xxx_TearDownContextOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextHandle = op.ContextHandle
	o.Rank = op.Rank
	o.TearDownType = op.TearDownType
	return op
}

func (o *TearDownContextRequest) xxx_FromOp(ctx context.Context, op *xxx_TearDownContextOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
	o.Rank = op.Rank
	o.TearDownType = op.TearDownType
}
func (o *TearDownContextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *TearDownContextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TearDownContextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// TearDownContextResponse structure represents the TearDownContext operation response
type TearDownContextResponse struct {
	// contextHandle: An RPC context handle, returned by a call to BuildContext or BuildContextW,
	// is correlated with a session object that is in the Active state. After TearDownContext
	// is executed, on either success or failure requests, contextHandle will be set to
	// null. For context handles, see [C706].
	ContextHandle *dcetypes.ContextHandle `idl:"name:contextHandle" json:"context_handle"`
	// Return: The TearDownContext return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *TearDownContextResponse) xxx_ToOp(ctx context.Context, op *xxx_TearDownContextOperation) *xxx_TearDownContextOperation {
	if op == nil {
		op = &xxx_TearDownContextOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextHandle = op.ContextHandle
	o.Return = op.Return
	return op
}

func (o *TearDownContextResponse) xxx_FromOp(ctx context.Context, op *xxx_TearDownContextOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
	o.Return = op.Return
}
func (o *TearDownContextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *TearDownContextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TearDownContextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BeginTearDownOperation structure represents the BeginTearDown operation
type xxx_BeginTearDownOperation struct {
	ContextHandle *dcetypes.ContextHandle `idl:"name:contextHandle" json:"context_handle"`
	TearDownType  TeardownType            `idl:"name:tearDownType" json:"tear_down_type"`
	Return        int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_BeginTearDownOperation) OpNum() int { return 5 }

func (o *xxx_BeginTearDownOperation) OpName() string { return "/IXnRemote/v1/BeginTearDown" }

func (o *xxx_BeginTearDownOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginTearDownOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// contextHandle {in} (1:{alias=PCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// tearDownType {in} (1:{alias=TEARDOWN_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.TearDownType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginTearDownOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// contextHandle {in} (1:{alias=PCONTEXT_HANDLE,pointer=ref}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &dcetypes.ContextHandle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// tearDownType {in} (1:{alias=TEARDOWN_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.TearDownType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginTearDownOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginTearDownOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BeginTearDownOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BeginTearDownRequest structure represents the BeginTearDown operation request
type BeginTearDownRequest struct {
	// contextHandle: An RPC context handle that is correlated with a session object that
	// is in the Active state. For context handles, see [C706].
	ContextHandle *dcetypes.ContextHandle `idl:"name:contextHandle" json:"context_handle"`
	// tearDownType: The reason for tearing down the session. It MUST be set to 0x00 (TT_FORCE).
	//
	//	+-------+----------+
	//	|       |          |
	//	| VALUE | MEANING  |
	//	|       |          |
	//	+-------+----------+
	//	+-------+----------+
	//	| 0x00  | TT_FORCE |
	//	+-------+----------+
	TearDownType TeardownType `idl:"name:tearDownType" json:"tear_down_type"`
}

func (o *BeginTearDownRequest) xxx_ToOp(ctx context.Context, op *xxx_BeginTearDownOperation) *xxx_BeginTearDownOperation {
	if op == nil {
		op = &xxx_BeginTearDownOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextHandle = op.ContextHandle
	o.TearDownType = op.TearDownType
	return op
}

func (o *BeginTearDownRequest) xxx_FromOp(ctx context.Context, op *xxx_BeginTearDownOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
	o.TearDownType = op.TearDownType
}
func (o *BeginTearDownRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BeginTearDownRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BeginTearDownOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BeginTearDownResponse structure represents the BeginTearDown operation response
type BeginTearDownResponse struct {
	// Return: The BeginTearDown return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BeginTearDownResponse) xxx_ToOp(ctx context.Context, op *xxx_BeginTearDownOperation) *xxx_BeginTearDownOperation {
	if op == nil {
		op = &xxx_BeginTearDownOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *BeginTearDownResponse) xxx_FromOp(ctx context.Context, op *xxx_BeginTearDownOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BeginTearDownResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BeginTearDownResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BeginTearDownOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PokeWOperation structure represents the PokeW operation
type xxx_PokeWOperation struct {
	Rank       SessionRank `idl:"name:sRank" json:"rank"`
	CalleeUUID string      `idl:"name:pwszCalleeUuid;string" json:"callee_uuid"`
	HostName   string      `idl:"name:pwszHostName;string" json:"host_name"`
	UUIDString string      `idl:"name:pwszUuidString;string" json:"uuid_string"`
	SizeOfBlob uint32      `idl:"name:dwcbSizeOfBlob" json:"size_of_blob"`
	Blob       []byte      `idl:"name:rguchBlob;size_is:(dwcbSizeOfBlob)" json:"blob"`
	Return     int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_PokeWOperation) OpNum() int { return 6 }

func (o *xxx_PokeWOperation) OpName() string { return "/IXnRemote/v1/PokeW" }

func (o *xxx_PokeWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Blob != nil && o.SizeOfBlob == 0 {
		o.SizeOfBlob = uint32(len(o.Blob))
	}
	if len(o.CalleeUUID) > int(37) {
		return fmt.Errorf("CalleeUUID is out of range")
	}
	if len(o.HostName) > int(16) {
		return fmt.Errorf("HostName is out of range")
	}
	if len(o.UUIDString) > int(37) {
		return fmt.Errorf("UUIDString is out of range")
	}
	if o.SizeOfBlob < uint32(8) || o.SizeOfBlob > uint32(8) {
		return fmt.Errorf("SizeOfBlob is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PokeWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// sRank {in} (1:{alias=SESSION_RANK}(enum))
	{
		if err := w.WriteEnum(uint16(o.Rank)); err != nil {
			return err
		}
	}
	// pwszCalleeUuid {in} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.CalleeUUID); err != nil {
			return err
		}
	}
	// pwszHostName {in} (1:{string, range=(1,16)}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.HostName); err != nil {
			return err
		}
	}
	// pwszUuidString {in} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.UUIDString); err != nil {
			return err
		}
	}
	// dwcbSizeOfBlob {in} (1:{range=(8,8), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeOfBlob); err != nil {
			return err
		}
	}
	// rguchBlob {in} (1:[dim:0,size_is=dwcbSizeOfBlob](uchar))
	{
		dimSize1 := uint64(o.SizeOfBlob)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Blob {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Blob[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Blob); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_PokeWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// sRank {in} (1:{alias=SESSION_RANK}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Rank)); err != nil {
			return err
		}
	}
	// pwszCalleeUuid {in} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.CalleeUUID); err != nil {
			return err
		}
	}
	// pwszHostName {in} (1:{string, range=(1,16)}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.HostName); err != nil {
			return err
		}
	}
	// pwszUuidString {in} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.UUIDString); err != nil {
			return err
		}
	}
	// dwcbSizeOfBlob {in} (1:{range=(8,8), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeOfBlob); err != nil {
			return err
		}
	}
	// rguchBlob {in} (1:[dim:0,size_is=dwcbSizeOfBlob](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Blob", sizeInfo[0])
		}
		o.Blob = make([]byte, sizeInfo[0])
		for i1 := range o.Blob {
			i1 := i1
			if err := w.ReadData(&o.Blob[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_PokeWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PokeWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PokeWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PokeWRequest structure represents the PokeW operation request
type PokeWRequest struct {
	// sRank: The SESSION_RANK of the partner making the call. This parameter MUST be set
	// to 0x02 (SRANK_SECONDARY).
	//
	//	+----------------------+------------------------------------------+
	//	|                      |                                          |
	//	|        VALUE         |                 MEANING                  |
	//	|                      |                                          |
	//	+----------------------+------------------------------------------+
	//	+----------------------+------------------------------------------+
	//	| SRANK_SECONDARY 0x02 | The caller is the secondary participant. |
	//	+----------------------+------------------------------------------+
	Rank SessionRank `idl:"name:sRank" json:"rank"`
	// pwszCalleeUuid: The string form of the primary partner contact identifier (CID).
	// The contact identifier (CID) MUST match the contact identifier (CID) in the primary
	// partner local name object, and MUST be formatted into a string.
	CalleeUUID string `idl:"name:pwszCalleeUuid;string" json:"callee_uuid"`
	// pwszHostName: The string form of the caller's host name. This host name identifies
	// the machine in which the caller's instance of the MSDTC Connection Manager: OleTx
	// Transports Protocol is running. This MUST be a NetBIOS name. For NetBIOS, see [NETBEUI],
	// [RFC1001], and [RFC1002].
	HostName string `idl:"name:pwszHostName;string" json:"host_name"`
	// pwszUuidString: The string form of the caller's contact identifier (CID). This contact
	// identifier (CID) identifies the caller's instance of the MSDTC Connection Manager:
	// OleTx Transports Protocol; it MUST match the contact identifier (CID) in the caller's
	// local name object and MUST be formatted into a string.
	UUIDString string `idl:"name:pwszUuidString;string" json:"uuid_string"`
	// dwcbSizeOfBlob: The count, in bytes, of the size of the binding info structure. This
	// parameter MUST be set to the size of the BIND_INFO_BLOB, 8.
	SizeOfBlob uint32 `idl:"name:dwcbSizeOfBlob" json:"size_of_blob"`
	// rguchBlob: A byte array that contains a BIND_INFO_BLOB structure.
	Blob []byte `idl:"name:rguchBlob;size_is:(dwcbSizeOfBlob)" json:"blob"`
}

func (o *PokeWRequest) xxx_ToOp(ctx context.Context, op *xxx_PokeWOperation) *xxx_PokeWOperation {
	if op == nil {
		op = &xxx_PokeWOperation{}
	}
	if o == nil {
		return op
	}
	o.Rank = op.Rank
	o.CalleeUUID = op.CalleeUUID
	o.HostName = op.HostName
	o.UUIDString = op.UUIDString
	o.SizeOfBlob = op.SizeOfBlob
	o.Blob = op.Blob
	return op
}

func (o *PokeWRequest) xxx_FromOp(ctx context.Context, op *xxx_PokeWOperation) {
	if o == nil {
		return
	}
	o.Rank = op.Rank
	o.CalleeUUID = op.CalleeUUID
	o.HostName = op.HostName
	o.UUIDString = op.UUIDString
	o.SizeOfBlob = op.SizeOfBlob
	o.Blob = op.Blob
}
func (o *PokeWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PokeWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PokeWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PokeWResponse structure represents the PokeW operation response
type PokeWResponse struct {
	// Return: The PokeW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PokeWResponse) xxx_ToOp(ctx context.Context, op *xxx_PokeWOperation) *xxx_PokeWOperation {
	if op == nil {
		op = &xxx_PokeWOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *PokeWResponse) xxx_FromOp(ctx context.Context, op *xxx_PokeWOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PokeWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PokeWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PokeWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BuildContextWOperation structure represents the BuildContextW operation
type xxx_BuildContextWOperation struct {
	Rank            SessionRank             `idl:"name:sRank" json:"rank"`
	BindVersionSet  *BindVersionSet         `idl:"name:BindVersionSet" json:"bind_version_set"`
	CalleeUUID      string                  `idl:"name:pwszCalleeUuid;string" json:"callee_uuid"`
	HostName        string                  `idl:"name:pwszHostName;string" json:"host_name"`
	UUIDString      string                  `idl:"name:pwszUuidString;string" json:"uuid_string"`
	GUIDIn          string                  `idl:"name:pwszGuidIn;string" json:"guid_in"`
	GUIDOut         string                  `idl:"name:pwszGuidOut;string" json:"guid_out"`
	BoundVersionSet *BoundVersionSet        `idl:"name:pBoundVersionSet" json:"bound_version_set"`
	SizeOfBlob      uint32                  `idl:"name:dwcbSizeOfBlob" json:"size_of_blob"`
	Blob            []byte                  `idl:"name:rguchBlob;size_is:(dwcbSizeOfBlob)" json:"blob"`
	Handle          *dcetypes.ContextHandle `idl:"name:ppHandle" json:"handle"`
	Return          int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_BuildContextWOperation) OpNum() int { return 7 }

func (o *xxx_BuildContextWOperation) OpName() string { return "/IXnRemote/v1/BuildContextW" }

func (o *xxx_BuildContextWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Blob != nil && o.SizeOfBlob == 0 {
		o.SizeOfBlob = uint32(len(o.Blob))
	}
	if len(o.CalleeUUID) > int(37) {
		return fmt.Errorf("CalleeUUID is out of range")
	}
	if len(o.HostName) > int(16) {
		return fmt.Errorf("HostName is out of range")
	}
	if len(o.UUIDString) > int(37) {
		return fmt.Errorf("UUIDString is out of range")
	}
	if len(o.GUIDIn) > int(37) {
		return fmt.Errorf("GUIDIn is out of range")
	}
	if len(o.GUIDOut) > int(37) {
		return fmt.Errorf("GUIDOut is out of range")
	}
	if o.SizeOfBlob < uint32(8) || o.SizeOfBlob > uint32(8) {
		return fmt.Errorf("SizeOfBlob is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BuildContextWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// sRank {in} (1:{alias=SESSION_RANK}(enum))
	{
		if err := w.WriteEnum(uint16(o.Rank)); err != nil {
			return err
		}
	}
	// BindVersionSet {in} (1:{alias=BIND_VERSION_SET}(struct))
	{
		if o.BindVersionSet != nil {
			if err := o.BindVersionSet.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&BindVersionSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pwszCalleeUuid {in} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.CalleeUUID); err != nil {
			return err
		}
	}
	// pwszHostName {in} (1:{string, range=(1,16)}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.HostName); err != nil {
			return err
		}
	}
	// pwszUuidString {in} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.UUIDString); err != nil {
			return err
		}
	}
	// pwszGuidIn {in} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.GUIDIn); err != nil {
			return err
		}
	}
	// pwszGuidOut {in, out} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.GUIDOut); err != nil {
			return err
		}
	}
	// pBoundVersionSet {in, out} (1:{pointer=ref}*(1))(2:{alias=BOUND_VERSION_SET}(struct))
	{
		if o.BoundVersionSet != nil {
			if err := o.BoundVersionSet.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&BoundVersionSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwcbSizeOfBlob {in} (1:{range=(8,8), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeOfBlob); err != nil {
			return err
		}
	}
	// rguchBlob {in} (1:[dim:0,size_is=dwcbSizeOfBlob](uchar))
	{
		dimSize1 := uint64(o.SizeOfBlob)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Blob {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Blob[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Blob); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_BuildContextWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// sRank {in} (1:{alias=SESSION_RANK}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Rank)); err != nil {
			return err
		}
	}
	// BindVersionSet {in} (1:{alias=BIND_VERSION_SET}(struct))
	{
		if o.BindVersionSet == nil {
			o.BindVersionSet = &BindVersionSet{}
		}
		if err := o.BindVersionSet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pwszCalleeUuid {in} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.CalleeUUID); err != nil {
			return err
		}
	}
	// pwszHostName {in} (1:{string, range=(1,16)}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.HostName); err != nil {
			return err
		}
	}
	// pwszUuidString {in} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.UUIDString); err != nil {
			return err
		}
	}
	// pwszGuidIn {in} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.GUIDIn); err != nil {
			return err
		}
	}
	// pwszGuidOut {in, out} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.GUIDOut); err != nil {
			return err
		}
	}
	// pBoundVersionSet {in, out} (1:{pointer=ref}*(1))(2:{alias=BOUND_VERSION_SET}(struct))
	{
		if o.BoundVersionSet == nil {
			o.BoundVersionSet = &BoundVersionSet{}
		}
		if err := o.BoundVersionSet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwcbSizeOfBlob {in} (1:{range=(8,8), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeOfBlob); err != nil {
			return err
		}
	}
	// rguchBlob {in} (1:[dim:0,size_is=dwcbSizeOfBlob](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Blob", sizeInfo[0])
		}
		o.Blob = make([]byte, sizeInfo[0])
		for i1 := range o.Blob {
			i1 := i1
			if err := w.ReadData(&o.Blob[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_BuildContextWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if len(o.GUIDOut) > int(37) {
		return fmt.Errorf("GUIDOut is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BuildContextWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pwszGuidOut {in, out} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.GUIDOut); err != nil {
			return err
		}
	}
	// pBoundVersionSet {in, out} (1:{pointer=ref}*(1))(2:{alias=BOUND_VERSION_SET}(struct))
	{
		if o.BoundVersionSet != nil {
			if err := o.BoundVersionSet.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&BoundVersionSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppHandle {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_BuildContextWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pwszGuidOut {in, out} (1:{string, range=(37,37)}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.GUIDOut); err != nil {
			return err
		}
	}
	// pBoundVersionSet {in, out} (1:{pointer=ref}*(1))(2:{alias=BOUND_VERSION_SET}(struct))
	{
		if o.BoundVersionSet == nil {
			o.BoundVersionSet = &BoundVersionSet{}
		}
		if err := o.BoundVersionSet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppHandle {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &dcetypes.ContextHandle{}
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

// BuildContextWRequest structure represents the BuildContextW operation request
type BuildContextWRequest struct {
	// sRank: The rank of the caller.
	//
	//	+----------------------+------------------------------------------------------+
	//	|                      |                                                      |
	//	|        VALUE         |                       MEANING                        |
	//	|                      |                                                      |
	//	+----------------------+------------------------------------------------------+
	//	+----------------------+------------------------------------------------------+
	//	| SRANK_PRIMARY 0x01   | The caller is the primary partner in this session.   |
	//	+----------------------+------------------------------------------------------+
	//	| SRANK_SECONDARY 0x02 | The caller is the secondary partner in this session. |
	//	+----------------------+------------------------------------------------------+
	Rank SessionRank `idl:"name:sRank" json:"rank"`
	// BindVersionSet: A BIND_VERSION_SET structure that contains the minimum and maximum
	// versions supported by the partner.
	BindVersionSet *BindVersionSet `idl:"name:BindVersionSet" json:"bind_version_set"`
	// pwszCalleeUuid: The string form of the callee's contact identifier (CID). The contact
	// identifier (CID) MUST match the contact identifier (CID) in the callee's local name
	// object and MUST be formatted into a string.
	CalleeUUID string `idl:"name:pwszCalleeUuid;string" json:"callee_uuid"`
	// pwszHostName: The string form of the caller's host name. This host name identifies
	// the machine in which the caller's instance of the MSDTC Connection Manager: OleTx
	// Transports Protocol is running. This MUST be a NetBIOS name. For NetBIOS, see [NETBEUI],
	// [RFC1001], and [RFC1002].
	HostName string `idl:"name:pwszHostName;string" json:"host_name"`
	// pwszUuidString: The string form of the caller's contact identifier (CID). This contact
	// identifier (CID) identifies the caller's instance of the MSDTC Connection Manager:
	// OleTx Transports Protocol. This MUST match the contact identifier (CID) in the caller's
	// local name object and MUST be formatted into a string.
	UUIDString string `idl:"name:pwszUuidString;string" json:"uuid_string"`
	// pwszGuidIn: A string form of a UUID that represents a unique identifier for this
	// bind attempt. The UUID MUST be formatted into a string.
	GUIDIn string `idl:"name:pwszGuidIn;string" json:"guid_in"`
	// pwszGuidOut: A string form of a UUID that represents a unique identifier for this
	// bind attempt. On input, the pwszGuidOut parameter MUST be set to 00000000-0000-0000-0000-000000000000.
	// On return, if the bind attempt is ultimately successful, the pwszGuidOut parameter
	// MUST be equal to the value of the pszGuidIn parameter. Otherwise, if the bind attempt
	// is ultimately unsuccessful, the pwszGuidOut parameter MUST be set to 00000000-0000-0000-0000-000000000000
	// on return.
	GUIDOut string `idl:"name:pwszGuidOut;string" json:"guid_out"`
	// pBoundVersionSet: A pointer to a BOUND_VERSION_SET structure. When the method is
	// called, every field of the BOUND_VERSION_SET structure MUST be initialized to zero.
	// This parameter receives a BOUND_VERSION_SET on successful completion and also on
	// return.
	BoundVersionSet *BoundVersionSet `idl:"name:pBoundVersionSet" json:"bound_version_set"`
	// dwcbSizeOfBlob: The count in bytes of the size of the binding info structure. This
	// parameter MUST be set to the size of BIND_INFO_BLOB, 8.
	SizeOfBlob uint32 `idl:"name:dwcbSizeOfBlob" json:"size_of_blob"`
	// rguchBlob: A byte array that contains a BIND_INFO_BLOB structure.
	Blob []byte `idl:"name:rguchBlob;size_is:(dwcbSizeOfBlob)" json:"blob"`
}

func (o *BuildContextWRequest) xxx_ToOp(ctx context.Context, op *xxx_BuildContextWOperation) *xxx_BuildContextWOperation {
	if op == nil {
		op = &xxx_BuildContextWOperation{}
	}
	if o == nil {
		return op
	}
	o.Rank = op.Rank
	o.BindVersionSet = op.BindVersionSet
	o.CalleeUUID = op.CalleeUUID
	o.HostName = op.HostName
	o.UUIDString = op.UUIDString
	o.GUIDIn = op.GUIDIn
	o.GUIDOut = op.GUIDOut
	o.BoundVersionSet = op.BoundVersionSet
	o.SizeOfBlob = op.SizeOfBlob
	o.Blob = op.Blob
	return op
}

func (o *BuildContextWRequest) xxx_FromOp(ctx context.Context, op *xxx_BuildContextWOperation) {
	if o == nil {
		return
	}
	o.Rank = op.Rank
	o.BindVersionSet = op.BindVersionSet
	o.CalleeUUID = op.CalleeUUID
	o.HostName = op.HostName
	o.UUIDString = op.UUIDString
	o.GUIDIn = op.GUIDIn
	o.GUIDOut = op.GUIDOut
	o.BoundVersionSet = op.BoundVersionSet
	o.SizeOfBlob = op.SizeOfBlob
	o.Blob = op.Blob
}
func (o *BuildContextWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BuildContextWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BuildContextWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BuildContextWResponse structure represents the BuildContextW operation response
type BuildContextWResponse struct {
	// pwszGuidOut: A string form of a UUID that represents a unique identifier for this
	// bind attempt. On input, the pwszGuidOut parameter MUST be set to 00000000-0000-0000-0000-000000000000.
	// On return, if the bind attempt is ultimately successful, the pwszGuidOut parameter
	// MUST be equal to the value of the pszGuidIn parameter. Otherwise, if the bind attempt
	// is ultimately unsuccessful, the pwszGuidOut parameter MUST be set to 00000000-0000-0000-0000-000000000000
	// on return.
	GUIDOut string `idl:"name:pwszGuidOut;string" json:"guid_out"`
	// pBoundVersionSet: A pointer to a BOUND_VERSION_SET structure. When the method is
	// called, every field of the BOUND_VERSION_SET structure MUST be initialized to zero.
	// This parameter receives a BOUND_VERSION_SET on successful completion and also on
	// return.
	BoundVersionSet *BoundVersionSet `idl:"name:pBoundVersionSet" json:"bound_version_set"`
	// ppHandle: On successful return, an RPC context handle (see [C706]) that correlates
	// with the session object created by, or referenced by, this method.
	Handle *dcetypes.ContextHandle `idl:"name:ppHandle" json:"handle"`
	// Return: The BuildContextW return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BuildContextWResponse) xxx_ToOp(ctx context.Context, op *xxx_BuildContextWOperation) *xxx_BuildContextWOperation {
	if op == nil {
		op = &xxx_BuildContextWOperation{}
	}
	if o == nil {
		return op
	}
	o.GUIDOut = op.GUIDOut
	o.BoundVersionSet = op.BoundVersionSet
	o.Handle = op.Handle
	o.Return = op.Return
	return op
}

func (o *BuildContextWResponse) xxx_FromOp(ctx context.Context, op *xxx_BuildContextWOperation) {
	if o == nil {
		return
	}
	o.GUIDOut = op.GUIDOut
	o.BoundVersionSet = op.BoundVersionSet
	o.Handle = op.Handle
	o.Return = op.Return
}
func (o *BuildContextWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BuildContextWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BuildContextWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
