package icaapi

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	tsts "github.com/oiweiwei/go-msrpc/msrpc/tsts"
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
	_ = tsts.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "tsts"
)

var (
	// Syntax UUID
	IcaAPISyntaxUUID = &uuid.UUID{TimeLow: 0x5ca4a760, TimeMid: 0xebb1, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x11, Node: [6]uint8{0x0, 0xa0, 0x24, 0x54, 0x20, 0xed}}
	// Syntax ID
	IcaAPISyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: IcaAPISyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// IcaApi interface.
type IcaAPIClient interface {

	// The RpcWinStationOpenServer method returns a server handle that can be used in other
	// WinStation API methods for querying information about the WinStation (sessions) on
	// the server. No special permissions are required to call this method.
	//
	// Return Values: Returns TRUE if the call succeeded, or FALSE if the call failed. On
	// failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationOpenServer(context.Context, *WinStationOpenServerRequest, ...dcerpc.CallOption) (*WinStationOpenServerResponse, error)

	// The RpcWinStationCloseServer method closes the server handle for WinStation APIs.
	// No special permissions are required to call this method.
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the call failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationCloseServer(context.Context, *WinStationCloseServerRequest, ...dcerpc.CallOption) (*WinStationCloseServerResponse, error)

	// The RpcIcaServerPing method is called to verify that the server is alive. No special
	// permissions are required to call this method.<161>
	//
	// Return Values:  Returns TRUE if the call succeeded and the server is alive, or FALSE
	// if the method failed. On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	IcaServerPing(context.Context, *IcaServerPingRequest, ...dcerpc.CallOption) (*IcaServerPingResponse, error)

	// The RpcWinStationEnumerate method retrieves a list of LOGONID structures for sessions
	// on a terminal server. No special permissions are required to call this method. However,
	// only sessions to which the caller has WINSTATION_QUERY permission are enumerated.
	// The method checks whether the caller has WINSTATION_QUERY permission (section 3.1.1)
	// by setting it as the Access Request mask, and fails if the caller does not have the
	// permission.
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the lookup failed.
	// On failure, pResult indicates the failure status code. If all of the logon IDs have
	// already been retrieved from the server, TRUE will be returned, and pResult will be
	// STATUS_NO_MORE_ENTRIES (as specified in [MS-ERREF]), indicating to the call that
	// all logon IDs have been retrieved.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationEnumerate(context.Context, *WinStationEnumerateRequest, ...dcerpc.CallOption) (*WinStationEnumerateResponse, error)

	// The RpcWinStationRename method enables the caller to change the name of the session.
	// The caller MUST have DELETE permission, as specified in [MS-DTYP] section 2.4.3,
	// on the session that is identified by the old name.<162>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationRename(context.Context, *WinStationRenameRequest, ...dcerpc.CallOption) (*WinStationRenameResponse, error)

	// The RpcWinStationQueryInformation method retrieves various types of configuration
	// information on a session. The caller MUST have the WINSTATION_QUERY permission right
	// as well as specific permission rights for some operations as indicated in the following
	// sections. The method checks whether the caller has WINSTATION_QUERY permission and
	// the specific permission required for some operations (section 3.1.1) by setting it
	// as the Access Request mask, and fails if the caller does not have the permission.<163>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationQueryInformation(context.Context, *WinStationQueryInformationRequest, ...dcerpc.CallOption) (*WinStationQueryInformationResponse, error)

	// The RpcWinStationSetInformation method sets various types of configuration information
	// for a session. The caller MUST have the WINSTATION_SET permission. Some operations
	// MUST have more specific permissions as indicated in more detail in the sections that
	// follow. The method checks whether the caller has WINSTATION_SET permission and the
	// specific permission for the configuration information (section 3.1.1) by setting
	// it as the Access Request mask, and fails if the caller does not have the permissions.<173>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationSetInformation(context.Context, *WinStationSetInformationRequest, ...dcerpc.CallOption) (*WinStationSetInformationResponse, error)

	// The RpcWinStationSendMessage method displays a message box on a given terminal server
	// session and, optionally, waits for a reply. The caller MUST have WINSTATION_MSG permission
	// for this method to succeed. The method checks whether the caller has WINSTATION_MSG
	// permission (section 3.1.1) by setting it as the Access Request mask, and fails if
	// the caller does not have the permission.<177>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationSendMessage(context.Context, *WinStationSendMessageRequest, ...dcerpc.CallOption) (*WinStationSendMessageResponse, error)

	// The RpcLogonIdFromWinStationName method returns a session's session ID given its
	// session name. The caller MUST have WINSTATION_QUERY permission. The method checks
	// whether the caller has WINSTATION_QUERY permission (section 3.1.1) by setting it
	// as the Access Request mask, and fails if the caller does not have the permission.<178>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	LogonIDFromWinStationName(context.Context, *LogonIDFromWinStationNameRequest, ...dcerpc.CallOption) (*LogonIDFromWinStationNameResponse, error)

	// The RpcWinStationNameFromLogonId method retrieves the Windows Station (WinStation)
	// name for a specific session. The caller MUST have WINSTATION_QUERY permission. The
	// method checks whether the caller has WINSTATION_QUERY permission (section 3.1.1)
	// by setting it as the Access Request mask, and fails if the caller does not have the
	// permission.<179>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationNameFromLogonID(context.Context, *WinStationNameFromLogonIDRequest, ...dcerpc.CallOption) (*WinStationNameFromLogonIDResponse, error)

	// The RpcWinStationConnect method connects a user's terminal server client from a given
	// terminal server session to a different terminal server session. If there is a user
	// connected to the client session, it will be disconnected at the end of this call.
	// If the method succeeds, the state of the session is State_Active as defined in the
	// WINSTATIONSTATECLASS enumeration (section 2.2.1.9).
	//
	// The client indicated by ConnectLogonId MUST have WINSTATION_DISCONNECT permission.
	// Similarly, TargetLogonId MUST have WINSTATION_CONNECT and WINSTATION_DISCONNECT permissions.
	// For each of the aforementioned permissions, the method checks whether the caller
	// has the permission (section 3.1.1) by setting the Access Request mask to the specific
	// permission, and fails if the caller does not have the permission.<180>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationConnect(context.Context, *WinStationConnectRequest, ...dcerpc.CallOption) (*WinStationConnectResponse, error)

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire

	// The RpcWinStationDisconnect method disconnects, on the server, the terminal server
	// client from a session. If the method succeeds, the state of the session is State_Disconnected
	// as defined in the WINSTATIONSTATECLASS enumeration (section 2.2.1.9).
	//
	// The caller of this method MUST have WINSTATION_DISCONNECT permission on the session
	// to disconnect. The method checks whether the caller has WINSTATION_DISCONNECT permission
	// (section 3.1.1) by setting it as the Access Request mask, and fails if the caller
	// does not have the permission.<182>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationDisconnect(context.Context, *WinStationDisconnectRequest, ...dcerpc.CallOption) (*WinStationDisconnectResponse, error)

	// The RpcWinStationReset method resets a session. Resetting a session will lead to
	// the user being logged off and his or her terminal server client being disconnected.
	// The caller MUST have WINSTATION_RESET permissions. The method checks whether the
	// caller has WINSTATION_RESET permission (section 3.1.1) by setting it as the Access
	// Request mask, and fails if the caller does not have the permission.<184>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationReset(context.Context, *WinStationResetRequest, ...dcerpc.CallOption) (*WinStationResetResponse, error)

	// The RpcWinStationShutdownSystem method shuts down the system and optionally logs
	// off all sessions and/or reboots the system. The caller requires SeShutdownPrivilege
	// (see [MSDN-PRVLGECNSTS]) when performing the shutdown locally and SeRemoteShutdownPrivilege
	// (see [MSDN-PRVLGECNSTS]) when performing the shutdown remotely. The caller calls
	// ExitWindowsEx (see [MSDN-ExitWindowsEx]) to perform the actual shutdown once all
	// checks have been completed.
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationShutdownSystem(context.Context, *WinStationShutdownSystemRequest, ...dcerpc.CallOption) (*WinStationShutdownSystemResponse, error)

	// The RpcWinStationWaitSystemEvent method waits synchronously for a system event from
	// an RPC API request on behalf of the caller. There is no time-out on the wait. Only
	// one event wait at a time can be posted per server handle. If an event wait is already
	// outstanding and the new request is not a cancel, the new request will fail. The caller
	// is not required to have any specific permission to call RpcWinStationWaitSystemEvent.
	// The first time this is called, the server will create an event block for the handle
	// specified by hServer. This event block will be cleared if RpcWinStationWaitSystemEvent
	// is called with EventMask equal to WEVENT_NONE or if RpcWinStationCloseServer or RpcWinStationCloseServerEx
	// are closed for the handle hServer.
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationWaitSystemEvent(context.Context, *WinStationWaitSystemEventRequest, ...dcerpc.CallOption) (*WinStationWaitSystemEventResponse, error)

	// The RpcWinStationShadow method starts a shadow (remote control) operation of another
	// terminal server session. If the method succeeds, the state of the session that started
	// the shadow operation is State_Shadow and the state of the session being shadowed
	// is State_Active as defined in the WINSTATIONSTATECLASS enumeration (section 2.2.1.9).
	//
	// The caller MUST have WINSTATION_SHADOW permission. The other session can be local
	// or on a terminal server. The method MUST be called from inside a remote terminal
	// server session. The session to shadow MUST be in the active state with a user logged
	// on. The method checks whether the caller has WINSTATION_SHADOW permission (section
	// 3.1.1) by setting it as the Access Request mask, and fails if the caller does not
	// have the permission.<186>
	//
	// Return Values: Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationShadow(context.Context, *WinStationShadowRequest, ...dcerpc.CallOption) (*WinStationShadowResponse, error)

	// Opnum18NotUsedOnWire operation.
	// Opnum18NotUsedOnWire

	// Opnum19NotUsedOnWire operation.
	// Opnum19NotUsedOnWire

	// Opnum20NotUsedOnWire operation.
	// Opnum20NotUsedOnWire

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire

	// Opnum22NotUsedOnWire operation.
	// Opnum22NotUsedOnWire

	// Opnum23NotUsedOnWire operation.
	// Opnum23NotUsedOnWire

	// Opnum24NotUsedOnWire operation.
	// Opnum24NotUsedOnWire

	// Opnum25NotUsedOnWire operation.
	// Opnum25NotUsedOnWire

	// Opnum26NotUsedOnWire operation.
	// Opnum26NotUsedOnWire

	// Opnum27NotUsedOnWire operation.
	// Opnum27NotUsedOnWire

	// Opnum28NotUsedOnWire operation.
	// Opnum28NotUsedOnWire

	// The RpcWinStationBreakPoint method breaks into the debugger in either the session
	// process of a specific session or in the terminal server service process. When this
	// method is called, the server impersonates the caller and then tries to enable SeShutdownPrivilege
	// (see [MSDN-PRVLGECNSTS]). If the attempt to enable this privilege fails, the RpcWinStationBreakPoint
	// call fails.<187>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationBreakPoint(context.Context, *WinStationBreakPointRequest, ...dcerpc.CallOption) (*WinStationBreakPointResponse, error)

	// The RpcWinStationReadRegistry method tells the terminal server to reread, from the
	// registry, the configuration data for all the WinStations, and to update the memory
	// locations where this data was stored with the values read from the registry.<189>
	// The caller of this RPC method MUST be running either as SYSTEM or as an Administrator.<190>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationReadRegistry(context.Context, *WinStationReadRegistryRequest, ...dcerpc.CallOption) (*WinStationReadRegistryResponse, error)

	// Opnum31NotUsedOnWire operation.
	// Opnum31NotUsedOnWire

	// Opnum32NotUsedOnWire operation.
	// Opnum32NotUsedOnWire

	// Opnum33NotUsedOnWire operation.
	// Opnum33NotUsedOnWire

	// The OldRpcWinStationEnumerateProcesses method calls the RpcWinStationEnumerateProcesses
	// method and returns whatever is returned by that method. It has the same parameters
	// as the RpcWinStationEnumerateProcesses method. No special permissions are required
	// to call this method.<191>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	OldRPCWinStationEnumerateProcesses(context.Context, *OldRPCWinStationEnumerateProcessesRequest, ...dcerpc.CallOption) (*OldRPCWinStationEnumerateProcessesResponse, error)

	// Opnum35NotUsedOnWire operation.
	// Opnum35NotUsedOnWire

	// The RpcWinStationEnumerateProcesses method retrieves the processes running on a remote
	// server on which the caller has WINSTATION_QUERY permission to retrieve information.
	// The method checks whether the caller has WINSTATION_QUERY permission (section 3.1.1)
	// by setting it as the Access Request mask, and fails if the caller does not have the
	// permission.<192>
	//
	// Return Values:  Returns TRUE if the call succeeded, or FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationEnumerateProcesses(context.Context, *WinStationEnumerateProcessesRequest, ...dcerpc.CallOption) (*WinStationEnumerateProcessesResponse, error)

	// The RpcWinStationTerminateProcess method terminates the specified process. An attempt
	// is made to enable the SE_DEBUG_PRIVILEGE privilege to kill processes not owned by
	// the current user, including processes running in other terminal server sessions.
	// Caller MUST have terminate permission to terminate the process.
	//
	// Return Values:  Returns TRUE if the call succeeded, and FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationTerminateProcess(context.Context, *WinStationTerminateProcessRequest, ...dcerpc.CallOption) (*WinStationTerminateProcessResponse, error)

	// Opnum38NotUsedOnWire operation.
	// Opnum38NotUsedOnWire

	// Opnum39NotUsedOnWire operation.
	// Opnum39NotUsedOnWire

	// Opnum40NotUsedOnWire operation.
	// Opnum40NotUsedOnWire

	// Opnum41NotUsedOnWire operation.
	// Opnum41NotUsedOnWire

	// Opnum42NotUsedOnWire operation.
	// Opnum42NotUsedOnWire

	// The RpcWinStationGetAllProcesses method retrieves the list of processes running on
	// the server machine. Only the processes from the sessions on which the user has WINSTATION_QUERY
	// permission will be retrieved. The method checks whether the caller has WINSTATION_QUERY
	// permission (section 3.1.1) by setting it as the Access Request mask, and fails if
	// the caller does not have the permission.
	//
	// Return Values:  Returns TRUE if the call succeeded, and FALSE if the lookup failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationGetAllProcesses(context.Context, *WinStationGetAllProcessesRequest, ...dcerpc.CallOption) (*WinStationGetAllProcessesResponse, error)

	// The RpcWinStationGetProcessSid method retrieves the process security identifier (SID),
	// as specified in [MS-DTYP] section 2.4.2, for a given process ID and process start
	// time combination.<193>The caller MUST have the PROCESS_QUERY_INFORMATION access right
	// to the process being queried and the TOKEN_QUERY access right to the access token
	// associated with the process. For more information on the process access rights, see
	// [MSDN-PROCRIGHTS]. For more information on access rights for access tokens, see [MSDN-TOKENRIGHTS].
	//
	// Return Values:  Returns TRUE if the call succeeded, and FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationGetProcessSID(context.Context, *WinStationGetProcessSIDRequest, ...dcerpc.CallOption) (*WinStationGetProcessSIDResponse, error)

	// The RpcWinStationGetTermSrvCountersValue method retrieves the current value of requested
	// terminal server performance counters. The caller is not required to have any specific
	// permission to call this method.
	//
	// Return Values:  Returns TRUE if the call succeeded, and FALSE if the method failed.
	// On failure, pResult indicates the failure status code. Individual entries in the
	// array pCounter will indicate whether or not the counter data for that counter could
	// be retrieved.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationGetTerminateServerCountersValue(context.Context, *WinStationGetTerminateServerCountersValueRequest, ...dcerpc.CallOption) (*WinStationGetTerminateServerCountersValueResponse, error)

	// The RpcWinStationReInitializeSecurity method reinitializes security for all non-console
	// WinStations (remote connection protocols). <202> Existing sessions will not be affected,
	// but future sessions will have the new security descriptor read from the registry
	// applied to them. This method MUST be called by processes running as SYSTEM or as
	// an Administrator.<203>
	//
	// Return Values:  Returns TRUE if the call succeeded, and FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationReInitializeSecurity(context.Context, *WinStationReInitializeSecurityRequest, ...dcerpc.CallOption) (*WinStationReInitializeSecurityResponse, error)

	// Opnum47NotUsedOnWire operation.
	// Opnum47NotUsedOnWire

	// Opnum48NotUsedOnWire operation.
	// Opnum48NotUsedOnWire

	// Opnum49NotUsedOnWire operation.
	// Opnum49NotUsedOnWire

	// Opnum50NotUsedOnWire operation.
	// Opnum50NotUsedOnWire

	// Opnum51NotUsedOnWire operation.
	// Opnum51NotUsedOnWire

	// Opnum52NotUsedOnWire operation.
	// Opnum52NotUsedOnWire

	// The RpcWinStationGetLanAdapterName method returns the name of the LAN adapter with
	// a specific LAN adapter number (lana) and transport type, if it is configured to be
	// used for a Terminal Services protocol connection. No special permissions are required
	// to call this method.<204><205>
	//
	// Return Values:  Returns TRUE if the call is successful, and FALSE if the method
	// fails. On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationGetLANAdapterName(context.Context, *WinStationGetLANAdapterNameRequest, ...dcerpc.CallOption) (*WinStationGetLANAdapterNameResponse, error)

	// Opnum54NotUsedOnWire operation.
	// Opnum54NotUsedOnWire

	// Opnum55NotUsedOnWire operation.
	// Opnum55NotUsedOnWire

	// Opnum56NotUsedOnWire operation.
	// Opnum56NotUsedOnWire

	// Opnum57NotUsedOnWire operation.
	// Opnum57NotUsedOnWire

	// The RpcWinStationUpdateSettings method rereads settings for all WinStations. The
	// caller MUST have WINSTATION_QUERY permission. The method checks whether the caller
	// has WINSTATION_QUERY permission (section 3.1.1) by setting it as the Access Request
	// mask, and fails if the caller does not have the permission.<206>
	//
	// Return Values:  Returns TRUE if the call succeeded, and FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationUpdateSettings(context.Context, *WinStationUpdateSettingsRequest, ...dcerpc.CallOption) (*WinStationUpdateSettingsResponse, error)

	// The RpcWinStationShadowStop method stops all shadow operations on the specified session,
	// including whether the session is acting as a shadow client (a session that is shadowing
	// another session) or as a shadow target (a session being shadowed by another session).
	// If the method succeeds, both the state of the session that started the shadow operation
	// and the state of the session being shadowed are State_Active as defined in the WINSTATIONSTATECLASS
	// enumeration (section 2.2.1.9).
	//
	// Caller MUST have WINSTATION_DISCONNECT and WINSTATION_RESET permissions. For each
	// aforementioned required permission, the method checks whether the caller has the
	// permission (section 3.1.1) by setting the Access Request mask to the specific permission,
	// and fails if the caller does not have the permission.<208>
	//
	// Return Values: Returns TRUE if the call succeeded, and FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationShadowStop(context.Context, *WinStationShadowStopRequest, ...dcerpc.CallOption) (*WinStationShadowStopResponse, error)

	// The RpcWinStationCloseServerEx method closes the server handle for WinStation APIs.
	// The call to this method MUST be serialized if there are multiple threads running;
	// otherwise, the behavior of this function is unknown. No special permissions are required
	// to call this method.
	//
	// Return Values: Returns TRUE if the call succeeded, and FALSE if the call failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationCloseServerEx(context.Context, *WinStationCloseServerExRequest, ...dcerpc.CallOption) (*WinStationCloseServerExResponse, error)

	// The RpcWinStationIsHelpAssistantSession method determines if a session is created
	// by the built-in HelpAssistant user account.<209> The caller is not required to have
	// any specific permission to call this method.
	//
	// Return Values:  Returns TRUE if the session is running as HelpAssistant, and FALSE
	// if this is not a HelpAssistant session or if an error was encountered during the
	// test. On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationIsHelpAssistantSession(context.Context, *WinStationIsHelpAssistantSessionRequest, ...dcerpc.CallOption) (*WinStationIsHelpAssistantSessionResponse, error)

	// The RpcWinStationGetMachinePolicy method returns a copy of the terminal server machine
	// policy to the caller.<210> The caller is not required to have any specific permission
	// to call this method.
	//
	// Return Values:  Returns TRUE if the call succeeded, and FALSE if the method failed.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationGetMachinePolicy(context.Context, *WinStationGetMachinePolicyRequest, ...dcerpc.CallOption) (*WinStationGetMachinePolicyResponse, error)

	// Opnum63NotUsedOnWire operation.
	// Opnum63NotUsedOnWire

	// Opnum64NotUsedOnWire operation.
	// Opnum64NotUsedOnWire

	// The RpcWinStationCheckLoopBack method checks if there is a loopback when a client
	// tries to connect. Loopback refers to opening a terminal server session on the local
	// machine.<211> The caller is not required to have any specific permission to call
	// this method.
	//
	// Return Values:  Returns FALSE if there is no loopback, and TRUE if a loopback was
	// detected. This method returns TRUE also in the case when an error was encountered
	// during the loopback test. In this case, the pResult value contains the relevant error
	// code.
	//
	//	+-------------------+----------------------------------------------------+
	//	|      RETURN       |                                                    |
	//	|    VALUE/CODE     |                    DESCRIPTION                     |
	//	|                   |                                                    |
	//	+-------------------+----------------------------------------------------+
	//	+-------------------+----------------------------------------------------+
	//	| 0x01 TRUE         | A loopback was detected or the method call failed. |
	//	+-------------------+----------------------------------------------------+
	//	| 0x00 FALSE        | A loopback was not detected.                       |
	//	+-------------------+----------------------------------------------------+
	WinStationCheckLoopBack(context.Context, *WinStationCheckLoopBackRequest, ...dcerpc.CallOption) (*WinStationCheckLoopBackResponse, error)

	// The RpcConnectCallback method initiates a TCP connection to the specified IP address
	// and waits for the party on the other end of the connection to start the Remote Desktop
	// Protocol (RDP) connection sequence. More information on Remote Desktop Protocol can
	// be found in [MS-RDPBCGR]. This method MUST be called by processes running as SYSTEM.
	// Note that this function assumes that the address being passed in is an IPv4 address.
	// IPv6 addresses are not supported.<212>
	//
	// Return Values:  Returns TRUE if the call succeeded, and FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	ConnectCallback(context.Context, *ConnectCallbackRequest, ...dcerpc.CallOption) (*ConnectCallbackResponse, error)

	// Opnum67NotUsedOnWire operation.
	// Opnum67NotUsedOnWire

	// Opnum68NotUsedOnWire operation.
	// Opnum68NotUsedOnWire

	// Opnum69NotUsedOnWire operation.
	// Opnum69NotUsedOnWire

	// The RpcWinStationGetAllProcesses_NT6 method retrieves the processes running a remote
	// server machine. Only the processes from the sessions on which the caller has WINSTATION_QUERY
	// permission will be retrieved. The method checks whether the caller has WINSTATION_QUERY
	// permission (section 3.1.1) by setting it as the Access Request mask, and fails if
	// the caller does not have the permission.
	//
	// Return Values:  Returns TRUE if the call succeeded, and FALSE if the lookup failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationGetAllProcessesNT6(context.Context, *WinStationGetAllProcessesNT6Request, ...dcerpc.CallOption) (*WinStationGetAllProcessesNT6Response, error)

	// Opnum71NotUsedOnWire operation.
	// Opnum71NotUsedOnWire

	// Opnum72NotUsedOnWire operation.
	// Opnum72NotUsedOnWire

	// Opnum73NotUsedOnWire operation.
	// Opnum73NotUsedOnWire

	// Opnum74NotUsedOnWire operation.
	// Opnum74NotUsedOnWire

	// The RpcWinStationOpenSessionDirectory method pings the Session Directory to see if
	// it can accept RPC calls. The caller MUST be either SYSTEM or an administrator. The
	// method performs access checks as defined in sections 3.1.2 and 3.1.3. The method
	// fails if both checks fail. For more information about the Session Directory, see
	// [MSFT-SDLBTS].<213>
	//
	// Return Values:  Returns TRUE if the call succeeded, and FALSE if the method failed.
	// On failure, pResult indicates the failure status code.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x01 TRUE         | Successful completion. |
	//	+-------------------+------------------------+
	//	| 0x00 FALSE        | Method call failed.    |
	//	+-------------------+------------------------+
	WinStationOpenSessionDirectory(context.Context, *WinStationOpenSessionDirectoryRequest, ...dcerpc.CallOption) (*WinStationOpenSessionDirectoryResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultIcaAPIClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultIcaAPIClient) WinStationOpenServer(ctx context.Context, in *WinStationOpenServerRequest, opts ...dcerpc.CallOption) (*WinStationOpenServerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationOpenServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationCloseServer(ctx context.Context, in *WinStationCloseServerRequest, opts ...dcerpc.CallOption) (*WinStationCloseServerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationCloseServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) IcaServerPing(ctx context.Context, in *IcaServerPingRequest, opts ...dcerpc.CallOption) (*IcaServerPingResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IcaServerPingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationEnumerate(ctx context.Context, in *WinStationEnumerateRequest, opts ...dcerpc.CallOption) (*WinStationEnumerateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationEnumerateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationRename(ctx context.Context, in *WinStationRenameRequest, opts ...dcerpc.CallOption) (*WinStationRenameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationRenameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationQueryInformation(ctx context.Context, in *WinStationQueryInformationRequest, opts ...dcerpc.CallOption) (*WinStationQueryInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationQueryInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationSetInformation(ctx context.Context, in *WinStationSetInformationRequest, opts ...dcerpc.CallOption) (*WinStationSetInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationSetInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationSendMessage(ctx context.Context, in *WinStationSendMessageRequest, opts ...dcerpc.CallOption) (*WinStationSendMessageResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationSendMessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) LogonIDFromWinStationName(ctx context.Context, in *LogonIDFromWinStationNameRequest, opts ...dcerpc.CallOption) (*LogonIDFromWinStationNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LogonIDFromWinStationNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationNameFromLogonID(ctx context.Context, in *WinStationNameFromLogonIDRequest, opts ...dcerpc.CallOption) (*WinStationNameFromLogonIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationNameFromLogonIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationConnect(ctx context.Context, in *WinStationConnectRequest, opts ...dcerpc.CallOption) (*WinStationConnectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationConnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationDisconnect(ctx context.Context, in *WinStationDisconnectRequest, opts ...dcerpc.CallOption) (*WinStationDisconnectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationDisconnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationReset(ctx context.Context, in *WinStationResetRequest, opts ...dcerpc.CallOption) (*WinStationResetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationResetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationShutdownSystem(ctx context.Context, in *WinStationShutdownSystemRequest, opts ...dcerpc.CallOption) (*WinStationShutdownSystemResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationShutdownSystemResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationWaitSystemEvent(ctx context.Context, in *WinStationWaitSystemEventRequest, opts ...dcerpc.CallOption) (*WinStationWaitSystemEventResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationWaitSystemEventResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationShadow(ctx context.Context, in *WinStationShadowRequest, opts ...dcerpc.CallOption) (*WinStationShadowResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationShadowResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationBreakPoint(ctx context.Context, in *WinStationBreakPointRequest, opts ...dcerpc.CallOption) (*WinStationBreakPointResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationBreakPointResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationReadRegistry(ctx context.Context, in *WinStationReadRegistryRequest, opts ...dcerpc.CallOption) (*WinStationReadRegistryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationReadRegistryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) OldRPCWinStationEnumerateProcesses(ctx context.Context, in *OldRPCWinStationEnumerateProcessesRequest, opts ...dcerpc.CallOption) (*OldRPCWinStationEnumerateProcessesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OldRPCWinStationEnumerateProcessesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationEnumerateProcesses(ctx context.Context, in *WinStationEnumerateProcessesRequest, opts ...dcerpc.CallOption) (*WinStationEnumerateProcessesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationEnumerateProcessesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationTerminateProcess(ctx context.Context, in *WinStationTerminateProcessRequest, opts ...dcerpc.CallOption) (*WinStationTerminateProcessResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationTerminateProcessResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationGetAllProcesses(ctx context.Context, in *WinStationGetAllProcessesRequest, opts ...dcerpc.CallOption) (*WinStationGetAllProcessesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationGetAllProcessesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationGetProcessSID(ctx context.Context, in *WinStationGetProcessSIDRequest, opts ...dcerpc.CallOption) (*WinStationGetProcessSIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationGetProcessSIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationGetTerminateServerCountersValue(ctx context.Context, in *WinStationGetTerminateServerCountersValueRequest, opts ...dcerpc.CallOption) (*WinStationGetTerminateServerCountersValueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationGetTerminateServerCountersValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationReInitializeSecurity(ctx context.Context, in *WinStationReInitializeSecurityRequest, opts ...dcerpc.CallOption) (*WinStationReInitializeSecurityResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationReInitializeSecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationGetLANAdapterName(ctx context.Context, in *WinStationGetLANAdapterNameRequest, opts ...dcerpc.CallOption) (*WinStationGetLANAdapterNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationGetLANAdapterNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationUpdateSettings(ctx context.Context, in *WinStationUpdateSettingsRequest, opts ...dcerpc.CallOption) (*WinStationUpdateSettingsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationUpdateSettingsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationShadowStop(ctx context.Context, in *WinStationShadowStopRequest, opts ...dcerpc.CallOption) (*WinStationShadowStopResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationShadowStopResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationCloseServerEx(ctx context.Context, in *WinStationCloseServerExRequest, opts ...dcerpc.CallOption) (*WinStationCloseServerExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationCloseServerExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationIsHelpAssistantSession(ctx context.Context, in *WinStationIsHelpAssistantSessionRequest, opts ...dcerpc.CallOption) (*WinStationIsHelpAssistantSessionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationIsHelpAssistantSessionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationGetMachinePolicy(ctx context.Context, in *WinStationGetMachinePolicyRequest, opts ...dcerpc.CallOption) (*WinStationGetMachinePolicyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationGetMachinePolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationCheckLoopBack(ctx context.Context, in *WinStationCheckLoopBackRequest, opts ...dcerpc.CallOption) (*WinStationCheckLoopBackResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationCheckLoopBackResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) ConnectCallback(ctx context.Context, in *ConnectCallbackRequest, opts ...dcerpc.CallOption) (*ConnectCallbackResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectCallbackResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationGetAllProcessesNT6(ctx context.Context, in *WinStationGetAllProcessesNT6Request, opts ...dcerpc.CallOption) (*WinStationGetAllProcessesNT6Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationGetAllProcessesNT6Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) WinStationOpenSessionDirectory(ctx context.Context, in *WinStationOpenSessionDirectoryRequest, opts ...dcerpc.CallOption) (*WinStationOpenSessionDirectoryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WinStationOpenSessionDirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != false {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIcaAPIClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIcaAPIClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewIcaAPIClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IcaAPIClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IcaAPISyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultIcaAPIClient{cc: cc}, nil
}

// xxx_WinStationOpenServerOperation structure represents the RpcWinStationOpenServer operation
type xxx_WinStationOpenServerOperation struct {
	Result uint32       `idl:"name:pResult" json:"result"`
	Server *tsts.Server `idl:"name:phServer" json:"server"`
	Return bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationOpenServerOperation) OpNum() int { return 0 }

func (o *xxx_WinStationOpenServerOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationOpenServer"
}

func (o *xxx_WinStationOpenServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationOpenServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_WinStationOpenServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_WinStationOpenServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationOpenServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// phServer {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationOpenServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// phServer {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationOpenServerRequest structure represents the RpcWinStationOpenServer operation request
type WinStationOpenServerRequest struct {
}

func (o *WinStationOpenServerRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationOpenServerOperation) *xxx_WinStationOpenServerOperation {
	if op == nil {
		op = &xxx_WinStationOpenServerOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *WinStationOpenServerRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationOpenServerOperation) {
	if o == nil {
		return
	}
}
func (o *WinStationOpenServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationOpenServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationOpenServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationOpenServerResponse structure represents the RpcWinStationOpenServer operation response
type WinStationOpenServerResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationOpenServer failed. If the
	// call was successful, this parameter is STATUS_SUCCESS (0x00000000) (as specified
	// in [MS-ERREF]).
	//
	//	+-----------------------------+----------------------------------------------+
	//	|                             |                                              |
	//	|            VALUE            |                   MEANING                    |
	//	|                             |                                              |
	//	+-----------------------------+----------------------------------------------+
	//	+-----------------------------+----------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000   | Successful call.                             |
	//	+-----------------------------+----------------------------------------------+
	//	| STATUS_CANCELLED 0xC0000120 | The server is shutting down.                 |
	//	+-----------------------------+----------------------------------------------+
	//	| STATUS_NO_MEMORY 0xC0000017 | Not enough memory to complete the operation. |
	//	+-----------------------------+----------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// phServer: Handle to the server object. This is of type SERVER_HANDLE. This handle
	// is used by other RpcWinStation methods.
	Server *tsts.Server `idl:"name:phServer" json:"server"`
	// Return: The RpcWinStationOpenServer return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationOpenServerResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationOpenServerOperation) *xxx_WinStationOpenServerOperation {
	if op == nil {
		op = &xxx_WinStationOpenServerOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Server = o.Server
	op.Return = o.Return
	return op
}

func (o *WinStationOpenServerResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationOpenServerOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Server = op.Server
	o.Return = op.Return
}
func (o *WinStationOpenServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationOpenServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationOpenServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationCloseServerOperation structure represents the RpcWinStationCloseServer operation
type xxx_WinStationCloseServerOperation struct {
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	Result uint32       `idl:"name:pResult" json:"result"`
	Return bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationCloseServerOperation) OpNum() int { return 1 }

func (o *xxx_WinStationCloseServerOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationCloseServer"
}

func (o *xxx_WinStationCloseServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCloseServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_WinStationCloseServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCloseServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCloseServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCloseServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationCloseServerRequest structure represents the RpcWinStationCloseServer operation request
type WinStationCloseServerRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// MUST be returned from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
}

func (o *WinStationCloseServerRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationCloseServerOperation) *xxx_WinStationCloseServerOperation {
	if op == nil {
		op = &xxx_WinStationCloseServerOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	return op
}

func (o *WinStationCloseServerRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationCloseServerOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
}
func (o *WinStationCloseServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationCloseServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationCloseServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationCloseServerResponse structure represents the RpcWinStationCloseServer operation response
type WinStationCloseServerResponse struct {
	// pResult: If the call was successful, this parameter is STATUS_SUCCESS (0x00000000),
	// as specified in [MS-ERREF]; otherwise, it MUST be an implementation-specific negative
	// value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationCloseServer return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationCloseServerResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationCloseServerOperation) *xxx_WinStationCloseServerOperation {
	if op == nil {
		op = &xxx_WinStationCloseServerOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationCloseServerResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationCloseServerOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationCloseServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationCloseServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationCloseServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IcaServerPingOperation structure represents the RpcIcaServerPing operation
type xxx_IcaServerPingOperation struct {
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	Result uint32       `idl:"name:pResult" json:"result"`
	Return bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_IcaServerPingOperation) OpNum() int { return 2 }

func (o *xxx_IcaServerPingOperation) OpName() string { return "/IcaApi/v1/RpcIcaServerPing" }

func (o *xxx_IcaServerPingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IcaServerPingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_IcaServerPingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IcaServerPingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IcaServerPingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IcaServerPingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IcaServerPingRequest structure represents the RpcIcaServerPing operation request
type IcaServerPingRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
}

func (o *IcaServerPingRequest) xxx_ToOp(ctx context.Context, op *xxx_IcaServerPingOperation) *xxx_IcaServerPingOperation {
	if op == nil {
		op = &xxx_IcaServerPingOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	return op
}

func (o *IcaServerPingRequest) xxx_FromOp(ctx context.Context, op *xxx_IcaServerPingOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
}
func (o *IcaServerPingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IcaServerPingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IcaServerPingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IcaServerPingResponse structure represents the RpcIcaServerPing operation response
type IcaServerPingResponse struct {
	// pResult: If the call was successful, this parameter is STATUS_SUCCESS (0x00000000),
	// as specified in [MS-ERREF]; otherwise it MUST be an implementation-specific negative
	// value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcIcaServerPing return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *IcaServerPingResponse) xxx_ToOp(ctx context.Context, op *xxx_IcaServerPingOperation) *xxx_IcaServerPingOperation {
	if op == nil {
		op = &xxx_IcaServerPingOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *IcaServerPingResponse) xxx_FromOp(ctx context.Context, op *xxx_IcaServerPingOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *IcaServerPingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IcaServerPingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IcaServerPingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationEnumerateOperation structure represents the RpcWinStationEnumerate operation
type xxx_WinStationEnumerateOperation struct {
	Server    *tsts.Server `idl:"name:hServer" json:"server"`
	Result    uint32       `idl:"name:pResult" json:"result"`
	Entries   uint32       `idl:"name:pEntries" json:"entries"`
	LogonID   []byte       `idl:"name:pLogonId;size_is:(pByteCount);pointer:unique" json:"logon_id"`
	ByteCount uint32       `idl:"name:pByteCount" json:"byte_count"`
	Index     uint32       `idl:"name:pIndex" json:"index"`
	Return    bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationEnumerateOperation) OpNum() int { return 3 }

func (o *xxx_WinStationEnumerateOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationEnumerate"
}

func (o *xxx_WinStationEnumerateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.LogonID != nil && o.ByteCount == 0 {
		o.ByteCount = uint32(len(o.LogonID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationEnumerateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pEntries {in, out} (1:{alias=PULONG}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Entries); err != nil {
			return err
		}
	}
	// pLogonId {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=pByteCount](char))
	{
		if o.LogonID != nil || o.ByteCount > 0 {
			_ptr_pLogonId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ByteCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.LogonID {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.LogonID[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.LogonID); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LogonID, _ptr_pLogonId); err != nil {
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
	// pByteCount {in, out} (1:{alias=PULONG}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ByteCount); err != nil {
			return err
		}
	}
	// pIndex {in, out} (1:{alias=PULONG}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationEnumerateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pEntries {in, out} (1:{alias=PULONG,pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Entries); err != nil {
			return err
		}
	}
	// pLogonId {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=pByteCount](char))
	{
		_ptr_pLogonId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.LogonID", sizeInfo[0])
			}
			o.LogonID = make([]byte, sizeInfo[0])
			for i1 := range o.LogonID {
				i1 := i1
				if err := w.ReadData(&o.LogonID[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pLogonId := func(ptr interface{}) { o.LogonID = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.LogonID, _s_pLogonId, _ptr_pLogonId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pByteCount {in, out} (1:{alias=PULONG,pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ByteCount); err != nil {
			return err
		}
	}
	// pIndex {in, out} (1:{alias=PULONG,pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationEnumerateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.LogonID != nil && o.ByteCount == 0 {
		o.ByteCount = uint32(len(o.LogonID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationEnumerateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pEntries {in, out} (1:{alias=PULONG}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Entries); err != nil {
			return err
		}
	}
	// pLogonId {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=pByteCount](char))
	{
		if o.LogonID != nil || o.ByteCount > 0 {
			_ptr_pLogonId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ByteCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.LogonID {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.LogonID[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.LogonID); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LogonID, _ptr_pLogonId); err != nil {
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
	// pByteCount {in, out} (1:{alias=PULONG}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ByteCount); err != nil {
			return err
		}
	}
	// pIndex {in, out} (1:{alias=PULONG}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationEnumerateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pEntries {in, out} (1:{alias=PULONG,pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Entries); err != nil {
			return err
		}
	}
	// pLogonId {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=pByteCount](char))
	{
		_ptr_pLogonId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.LogonID", sizeInfo[0])
			}
			o.LogonID = make([]byte, sizeInfo[0])
			for i1 := range o.LogonID {
				i1 := i1
				if err := w.ReadData(&o.LogonID[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pLogonId := func(ptr interface{}) { o.LogonID = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.LogonID, _s_pLogonId, _ptr_pLogonId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pByteCount {in, out} (1:{alias=PULONG,pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ByteCount); err != nil {
			return err
		}
	}
	// pIndex {in, out} (1:{alias=PULONG,pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationEnumerateRequest structure represents the RpcWinStationEnumerate operation request
type WinStationEnumerateRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// pEntries:  Pointer to the number of entries to return to the caller. On return from
	// this method, this is the number of logon IDs actually returned in this call to RpcWinStationEnumerate.
	Entries uint32 `idl:"name:pEntries" json:"entries"`
	// pLogonId:  Buffer where the logon IDs are stored when the method returns. This will
	// be an array of LOGONID structures. Caller MUST cast this to PCHAR before calling
	// this method.
	LogonID []byte `idl:"name:pLogonId;size_is:(pByteCount);pointer:unique" json:"logon_id"`
	// pByteCount:  Size of the buffer, in bytes, to which pLogonId points.
	ByteCount uint32 `idl:"name:pByteCount" json:"byte_count"`
	// pIndex:  Last index of the logon ID lookup from this call, passed to the server
	// the next time this method is called. Initial value of this passed by the caller MUST
	// be 0.
	Index uint32 `idl:"name:pIndex" json:"index"`
}

func (o *WinStationEnumerateRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationEnumerateOperation) *xxx_WinStationEnumerateOperation {
	if op == nil {
		op = &xxx_WinStationEnumerateOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Entries = o.Entries
	op.LogonID = o.LogonID
	op.ByteCount = o.ByteCount
	op.Index = o.Index
	return op
}

func (o *WinStationEnumerateRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationEnumerateOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Entries = op.Entries
	o.LogonID = op.LogonID
	o.ByteCount = op.ByteCount
	o.Index = op.Index
}
func (o *WinStationEnumerateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationEnumerateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationEnumerateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationEnumerateResponse structure represents the RpcWinStationEnumerate operation response
type WinStationEnumerateResponse struct {
	// pResult: If the call was successful, this parameter will be STATUS_SUCCESS (0x00000000),
	// as specified in [MS-ERREF]; otherwise, it MUST be an implementation-specific negative
	// value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// pEntries:  Pointer to the number of entries to return to the caller. On return from
	// this method, this is the number of logon IDs actually returned in this call to RpcWinStationEnumerate.
	Entries uint32 `idl:"name:pEntries" json:"entries"`
	// pLogonId:  Buffer where the logon IDs are stored when the method returns. This will
	// be an array of LOGONID structures. Caller MUST cast this to PCHAR before calling
	// this method.
	LogonID []byte `idl:"name:pLogonId;size_is:(pByteCount);pointer:unique" json:"logon_id"`
	// pByteCount:  Size of the buffer, in bytes, to which pLogonId points.
	ByteCount uint32 `idl:"name:pByteCount" json:"byte_count"`
	// pIndex:  Last index of the logon ID lookup from this call, passed to the server
	// the next time this method is called. Initial value of this passed by the caller MUST
	// be 0.
	Index uint32 `idl:"name:pIndex" json:"index"`
	// Return: The RpcWinStationEnumerate return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationEnumerateResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationEnumerateOperation) *xxx_WinStationEnumerateOperation {
	if op == nil {
		op = &xxx_WinStationEnumerateOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Entries = o.Entries
	op.LogonID = o.LogonID
	op.ByteCount = o.ByteCount
	op.Index = o.Index
	op.Return = o.Return
	return op
}

func (o *WinStationEnumerateResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationEnumerateOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Entries = op.Entries
	o.LogonID = op.LogonID
	o.ByteCount = op.ByteCount
	o.Index = op.Index
	o.Return = op.Return
}
func (o *WinStationEnumerateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationEnumerateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationEnumerateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationRenameOperation structure represents the RpcWinStationRename operation
type xxx_WinStationRenameOperation struct {
	Server            *tsts.Server `idl:"name:hServer" json:"server"`
	Result            uint32       `idl:"name:pResult" json:"result"`
	WinStationNameOld string       `idl:"name:pWinStationNameOld;size_is:(NameOldSize)" json:"win_station_name_old"`
	NameOldSize       uint32       `idl:"name:NameOldSize" json:"name_old_size"`
	WinStationNameNew string       `idl:"name:pWinStationNameNew;size_is:(NameNewSize)" json:"win_station_name_new"`
	NameNewSize       uint32       `idl:"name:NameNewSize" json:"name_new_size"`
	Return            bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationRenameOperation) OpNum() int { return 4 }

func (o *xxx_WinStationRenameOperation) OpName() string { return "/IcaApi/v1/RpcWinStationRename" }

func (o *xxx_WinStationRenameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.WinStationNameOld != "" && o.NameOldSize == 0 {
		o.NameOldSize = uint32(ndr.UTF16Len(o.WinStationNameOld))
	}
	if o.WinStationNameNew != "" && o.NameNewSize == 0 {
		o.NameNewSize = uint32(ndr.UTF16Len(o.WinStationNameNew))
	}
	if o.NameOldSize > uint32(256) {
		return fmt.Errorf("NameOldSize is out of range")
	}
	if o.NameNewSize > uint32(256) {
		return fmt.Errorf("NameNewSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationRenameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pWinStationNameOld {in} (1:{alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameOldSize,string](wchar))
	{
		dimSize1 := uint64(o.NameOldSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_WinStationNameOld_buf := utf16.Encode([]rune(o.WinStationNameOld))
		if uint64(len(_WinStationNameOld_buf)) > sizeInfo[0] {
			_WinStationNameOld_buf = _WinStationNameOld_buf[:sizeInfo[0]]
		}
		for i1 := range _WinStationNameOld_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_WinStationNameOld_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_WinStationNameOld_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// NameOldSize {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameOldSize); err != nil {
			return err
		}
	}
	// pWinStationNameNew {in} (1:{alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameNewSize,string](wchar))
	{
		dimSize1 := uint64(o.NameNewSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_WinStationNameNew_buf := utf16.Encode([]rune(o.WinStationNameNew))
		if uint64(len(_WinStationNameNew_buf)) > sizeInfo[0] {
			_WinStationNameNew_buf = _WinStationNameNew_buf[:sizeInfo[0]]
		}
		for i1 := range _WinStationNameNew_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_WinStationNameNew_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_WinStationNameNew_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// NameNewSize {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameNewSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationRenameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pWinStationNameOld {in} (1:{alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameOldSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _WinStationNameOld_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _WinStationNameOld_buf", sizeInfo[0])
		}
		_WinStationNameOld_buf = make([]uint16, sizeInfo[0])
		for i1 := range _WinStationNameOld_buf {
			i1 := i1
			if err := w.ReadData(&_WinStationNameOld_buf[i1]); err != nil {
				return err
			}
		}
		o.WinStationNameOld = strings.TrimRight(string(utf16.Decode(_WinStationNameOld_buf)), ndr.ZeroString)
	}
	// NameOldSize {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameOldSize); err != nil {
			return err
		}
	}
	// pWinStationNameNew {in} (1:{alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameNewSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _WinStationNameNew_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _WinStationNameNew_buf", sizeInfo[0])
		}
		_WinStationNameNew_buf = make([]uint16, sizeInfo[0])
		for i1 := range _WinStationNameNew_buf {
			i1 := i1
			if err := w.ReadData(&_WinStationNameNew_buf[i1]); err != nil {
				return err
			}
		}
		o.WinStationNameNew = strings.TrimRight(string(utf16.Decode(_WinStationNameNew_buf)), ndr.ZeroString)
	}
	// NameNewSize {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameNewSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationRenameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationRenameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationRenameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationRenameRequest structure represents the RpcWinStationRename operation request
type WinStationRenameRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// pWinStationNameOld:  The pointer to a string that is the old name of the session
	// being renamed.
	WinStationNameOld string `idl:"name:pWinStationNameOld;size_is:(NameOldSize)" json:"win_station_name_old"`
	// NameOldSize:  The length of the string in characters pointed to by pWinStationNameOld
	// including the terminating NULL character.
	NameOldSize uint32 `idl:"name:NameOldSize" json:"name_old_size"`
	// pWinStationNameNew:  The pointer to a string that is the new name of the session
	// being renamed.
	WinStationNameNew string `idl:"name:pWinStationNameNew;size_is:(NameNewSize)" json:"win_station_name_new"`
	// NameNewSize: The length of the string in characters pointed to by pWinStationNameNew
	// including the terminating NULL character. Name MUST be shorter than or equal to WINSTATIONNAME_LENGTH.
	NameNewSize uint32 `idl:"name:NameNewSize" json:"name_new_size"`
}

func (o *WinStationRenameRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationRenameOperation) *xxx_WinStationRenameOperation {
	if op == nil {
		op = &xxx_WinStationRenameOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.WinStationNameOld = o.WinStationNameOld
	op.NameOldSize = o.NameOldSize
	op.WinStationNameNew = o.WinStationNameNew
	op.NameNewSize = o.NameNewSize
	return op
}

func (o *WinStationRenameRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationRenameOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.WinStationNameOld = op.WinStationNameOld
	o.NameOldSize = op.NameOldSize
	o.WinStationNameNew = op.WinStationNameNew
	o.NameNewSize = op.NameNewSize
}
func (o *WinStationRenameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationRenameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationRenameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationRenameResponse structure represents the RpcWinStationRename operation response
type WinStationRenameResponse struct {
	// pResult:  The failure error code if the call to RpcWinStationRename fails. If the
	// call is successful, this parameter MUST be STATUS_SUCCESS (0x00000000), as specified
	// in [MS-ERREF].
	//
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                 |                                                                                  |
	//	|                      VALUE                      |                                     MEANING                                      |
	//	|                                                 |                                                                                  |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000                       | The call is successful.                                                          |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022                 | The caller does not have DELETE permission.                                      |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_CTX_WINSTATION_NAME_INVALID 0xC00A0001   | The sizes are 0, one or the other of the pointers is NULL, or a pointer is       |
	//	|                                                 | invalid.                                                                         |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_CTX_WINSTATION_NOT_FOUND 0xC00A0015      | No session exists with the name given in pWinStationNameOld.                     |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_CTX_WINSTATION_NAME_COLLISION 0xC00A0016 | A session already exists with the name given in pWinStationNameNew.              |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationRename return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationRenameResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationRenameOperation) *xxx_WinStationRenameOperation {
	if op == nil {
		op = &xxx_WinStationRenameOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationRenameResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationRenameOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationRenameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationRenameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationRenameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationQueryInformationOperation structure represents the RpcWinStationQueryInformation operation
type xxx_WinStationQueryInformationOperation struct {
	Server                      *tsts.Server `idl:"name:hServer" json:"server"`
	Result                      uint32       `idl:"name:pResult" json:"result"`
	LogonID                     uint32       `idl:"name:LogonId" json:"logon_id"`
	WinStationInformationClass  uint32       `idl:"name:WinStationInformationClass" json:"win_station_information_class"`
	WinStationInformation       []byte       `idl:"name:pWinStationInformation;size_is:(WinStationInformationLength);pointer:unique" json:"win_station_information"`
	WinStationInformationLength uint32       `idl:"name:WinStationInformationLength" json:"win_station_information_length"`
	ReturnLength                uint32       `idl:"name:pReturnLength" json:"return_length"`
	Return                      bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationQueryInformationOperation) OpNum() int { return 5 }

func (o *xxx_WinStationQueryInformationOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationQueryInformation"
}

func (o *xxx_WinStationQueryInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.WinStationInformation != nil && o.WinStationInformationLength == 0 {
		o.WinStationInformationLength = uint32(len(o.WinStationInformation))
	}
	if o.WinStationInformationLength > uint32(32768) {
		return fmt.Errorf("WinStationInformationLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationQueryInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LogonID); err != nil {
			return err
		}
	}
	// WinStationInformationClass {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.WinStationInformationClass); err != nil {
			return err
		}
	}
	// pWinStationInformation {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=WinStationInformationLength](char))
	{
		if o.WinStationInformation != nil || o.WinStationInformationLength > 0 {
			_ptr_pWinStationInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.WinStationInformationLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.WinStationInformation {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.WinStationInformation[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.WinStationInformation); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WinStationInformation, _ptr_pWinStationInformation); err != nil {
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
	// WinStationInformationLength {in} (1:{range=(0,32768), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.WinStationInformationLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationQueryInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LogonID); err != nil {
			return err
		}
	}
	// WinStationInformationClass {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.WinStationInformationClass); err != nil {
			return err
		}
	}
	// pWinStationInformation {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=WinStationInformationLength](char))
	{
		_ptr_pWinStationInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.WinStationInformation", sizeInfo[0])
			}
			o.WinStationInformation = make([]byte, sizeInfo[0])
			for i1 := range o.WinStationInformation {
				i1 := i1
				if err := w.ReadData(&o.WinStationInformation[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pWinStationInformation := func(ptr interface{}) { o.WinStationInformation = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.WinStationInformation, _s_pWinStationInformation, _ptr_pWinStationInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WinStationInformationLength {in} (1:{range=(0,32768), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.WinStationInformationLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationQueryInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationQueryInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pWinStationInformation {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=WinStationInformationLength](char))
	{
		if o.WinStationInformation != nil || o.WinStationInformationLength > 0 {
			_ptr_pWinStationInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.WinStationInformationLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.WinStationInformation {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.WinStationInformation[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.WinStationInformation); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WinStationInformation, _ptr_pWinStationInformation); err != nil {
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
	// pReturnLength {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ReturnLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationQueryInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pWinStationInformation {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=WinStationInformationLength](char))
	{
		_ptr_pWinStationInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.WinStationInformation", sizeInfo[0])
			}
			o.WinStationInformation = make([]byte, sizeInfo[0])
			for i1 := range o.WinStationInformation {
				i1 := i1
				if err := w.ReadData(&o.WinStationInformation[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pWinStationInformation := func(ptr interface{}) { o.WinStationInformation = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.WinStationInformation, _s_pWinStationInformation, _ptr_pWinStationInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pReturnLength {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ReturnLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationQueryInformationRequest structure represents the RpcWinStationQueryInformation operation request
type WinStationQueryInformationRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. hServer MUST
	// be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// LogonId: The session ID of the session for which to retrieve information.
	LogonID uint32 `idl:"name:LogonId" json:"logon_id"`
	// WinStationInformationClass: The class of data to retrieve. These values come from
	// the enum type WINSTATIONINFOCLASS.
	//
	// The following classes are supported.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationCreateData 0                | Retrieves general information on the type of terminal server session (protocol)  |
	//	|                                       | to which the session belongs. The pWinStationInformation argument points         |
	//	|                                       | to a WINSTATIONCREATE structure, and WinStationInformationLength MUST be         |
	//	|                                       | sizeof(WINSTATIONCREATE).<164>                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationConfiguration 1             | Retrieves general configuration data on the terminal server session. The         |
	//	|                                       | pWinStationInformation argument points to a WINSTACONFIGWIRE structure           |
	//	|                                       | followed by a USERCONFIG structure. The WinStationInformationLength MUST be      |
	//	|                                       | sizeof(WINSTACONFIGWIRE) + sizeof(USERCONFIG). The Size field in the USERCONFIG  |
	//	|                                       | structure inside WINSTACONFIGWIRE MUST be set to sizeof(USERCONFIG) and          |
	//	|                                       | the Offset set to sizeof(WINSTACONFIGWIRE). The Size field in the NewFields      |
	//	|                                       | structure inside WINSTACONFIGWIRE MUST be set to 0, and the offset MUST be set   |
	//	|                                       | to sizeof(WINSTACONFIGWIRE) + sizeof(USERCONFIG).                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationPdParams 2                  | Retrieves transport protocol driver parameters.<165> The structure coming into   |
	//	|                                       | the function indicates via SDClass the specific protocol driver on which to      |
	//	|                                       | receive parameter information. The result will be returned in the union in       |
	//	|                                       | the structure. The pWinStationInformation argument points to a PDPARAMSWIRE      |
	//	|                                       | structure followed by a PDPARAMS structure. The WinStationInformationLength MUST |
	//	|                                       | be sizeof(PDPARAMSWIRE) + sizeof(PDPARAMS). The Size field in SdClassSpecific    |
	//	|                                       | field inside PDPARAMSWIRE MUST be set to sizeof(PDPARAMS) and the offset MUST be |
	//	|                                       | set to sizeof(PDPARAMSWIRE).                                                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationWd 3                        | Retrieves WinStation protocol driver configuration data for the session.<166>    |
	//	|                                       | The pWinStationInformation argument points to a VARDATA_WIRE structure           |
	//	|                                       | followed by a WDCONFIG structure. The WinStationInformationLength                |
	//	|                                       | MUST be sizeof(VARDATA_WIRE) + sizeof(WDCONFIG). The Size field in the           |
	//	|                                       | VARDATA_WIRE structure MUST be set to sizeof(WDCONFIG) and the Offset set to     |
	//	|                                       | sizeof(VARDATA_WIRE).                                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationPd 4                        | Retrieves transport protocol driver configuration data for the session.<167>     |
	//	|                                       | The pWinStationInformation argument points to a VARDATA_WIRE, followed by        |
	//	|                                       | a PDPARAMSWIRE structure, followed by a PDCONFIG2 structure and finally          |
	//	|                                       | ending with a PDPARAMS structure. The WinStationInformationLength MUST           |
	//	|                                       | be sizeof(VARDATA_WIRE) + sizeof(PDPARAMSWIRE) + size of(PDCONFIG2) +            |
	//	|                                       | sizeof(PDPARAMS). The Size field in the VARDATA_WIRE structure MUST be           |
	//	|                                       | set to sizeof(PDCONFIG2) and the Offset set to sizeof(VARDATA_WIRE) +            |
	//	|                                       | sizeof(PDPARAMSWIRE). The Size field in SdClassSpecific field inside             |
	//	|                                       | PDPARAMSWIRE MUST be set to sizeof(PDPARAMS) - sizeof(SDCLASS), and the offset   |
	//	|                                       | MUST be set to Offset + Size of the VARDATA_WIRE structure.                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationPrinter 5                   | Not supported.                                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationClient 6                    | Retrieves data on the terminal server client of the session. The                 |
	//	|                                       | pWinStationInformation argument points to a VARDATA_WIRE structure followed      |
	//	|                                       | by a WINSTATIONCLIENT structure. The WinStationInformationLength MUST be         |
	//	|                                       | sizeof(VARDATA_WIRE) + sizeof(WINSTATIONCLIENT). The Size field in the           |
	//	|                                       | VARDATA_WIRE structure MUST be set to sizeof(WINSTATIONCLIENT) and the Offset    |
	//	|                                       | set to sizeof(VARDATA_WIRE).                                                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationModules 7                   | Internal get function to retrieve data on protocol-specific binaries             |
	//	|                                       | loaded for the given terminal server session. The structure pointed to           |
	//	|                                       | by pWinStationInformation and the size of the buffer is Terminal Service         |
	//	|                                       | protocol-specific.                                                               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationInformation 8               | Retrieves information on the session, including connect state, session's         |
	//	|                                       | name, connect time, disconnect time, time last input was received from           |
	//	|                                       | the client, logon time, user's username and domain, and the current time.        |
	//	|                                       | pWinStationInformation points to a VARDATA_WIRE structure followed by a          |
	//	|                                       | WINSTATIONINFORMATION structure. The WinStationInformationLength MUST be         |
	//	|                                       | sizeof(VARDATA_WIRE) + sizeof(WINSTATIONINFORMATION). The Size field in the      |
	//	|                                       | VARDATA_WIRE structure MUST be set to sizeof(WINSTATIONINFORMATION) and the      |
	//	|                                       | Offset set to sizeof(VARDATA_WIRE).                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationUserToken 14                | Retrieves the user's token in the session. Caller requires WINSTATION_ALL_ACCESS |
	//	|                                       | permission. The pWinStationInformation argument points to a WINSTATIONUSERTOKEN  |
	//	|                                       | structure, and WinStationInformationLength MUST be sizeof(WINSTATIONUSERTOKEN).  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationVideoData 16                | Retrieves resolution and color depth of the session. The                         |
	//	|                                       | pWinStationInformation argument points to a WINSTATIONVIDEODATA structure, and   |
	//	|                                       | WinStationInformationLength MUST be sizeof(WINSTATIONVIDEODATA).                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationCd 18                       | Retrieves connection driver configuration data. The pWinStationInformation       |
	//	|                                       | points to a CDCONFIG structure, and WinStationInformationLength MUST be          |
	//	|                                       | sizeof(CDCONFIG).                                                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationVirtualData 20              | Query client virtual data. The pWinStationInformation argument MUST point to 8   |
	//	|                                       | bytes and WinStationInformationLength MUST be 8.                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationLoadBalanceSessionTarget 24 | Retrieves the target session ID for a client redirected from another server      |
	//	|                                       | in a load balancing cluster. The pWinStationInformation points to a ULONG, and   |
	//	|                                       | WinStationInformationLength MUST be sizeof(ULONG). If there is no redirection,   |
	//	|                                       | -1 is returned in pWinStationInformation.<168>                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationLoadIndicator 25            | Retrieves an indicator of the load on the server. The pWinStationInformation     |
	//	|                                       | argument points to a WINSTATIONLOADINDICATORDATA structure.                      |
	//	|                                       | WinStationInformationLength MUST be sizeof(WINSTATIONLOADINDICATORDATA).         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationShadowInfo 26               | Retrieves the current shadow state of a session. The pWinStationInformation      |
	//	|                                       | argument points to a WINSTATIONSHADOW structure. WinStationInformationLength     |
	//	|                                       | MUST be sizeof(WINSTATIONSHADOW).<169>                                           |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationDigProductId 27             | Retrieves the client product ID and current product ID of the session. The       |
	//	|                                       | pWinStationInformation argument points to a WINSTATIONPRODID structure.          |
	//	|                                       | WinStationInformationLength MUST be sizeof(WINSTATIONPRODID).<170>               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationLockedState 28              | Retrieves the current locked state of the session, TRUE or FALSE.                |
	//	|                                       | The pWinStationInformation argument points to a BOOL variable.                   |
	//	|                                       | WinStationInformationLength MUST be sizeof(BOOL).                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationRemoteAddress 29            | Retrieves the remote IP address of the terminal server client in the session.    |
	//	|                                       | The pWinStationInformation argument points to a WINSTATIONREMOTEADDRESS          |
	//	|                                       | structure. WinStationInformationLength MUST be sizeof(WINSTATIONREMOTEADDRESS).  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationIdleTime 30                 | Retrieves the idle time for the session, in seconds. The pWinStationInformation  |
	//	|                                       | argument points to a ULONG variable. WinStationInformationLength MUST be         |
	//	|                                       | sizeof(ULONG).<171>                                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationLastReconnectType 31        | Retrieves the last reconnect type for the session. The value placed              |
	//	|                                       | in pWinStationInformation will come from the enum RECONNECT_TYPE.                |
	//	|                                       | The pWinStationInformation argument points to a ULONG variable.                  |
	//	|                                       | WinStationInformationLength MUST be sizeof(ULONG).<172>                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationDisallowAutoReconnect 32    | Retrieves the allow (1) or disallow (0) state for auto-reconnect, 1 or           |
	//	|                                       | 0. The pWinStationInformation argument points to a BOOLEAN variable.             |
	//	|                                       | WinStationInformationLength MUST be sizeof(BOOLEAN).                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationReconnectedFromId 37        | In case of reconnected sessions, this will return the session ID of the          |
	//	|                                       | temporary session from which it was reconnected, or -1 if no temporary session   |
	//	|                                       | was created. The WinStationInformationLength argument points to a ULONG          |
	//	|                                       | variable. WinStationInformationLength MUST be sizeof(ULONG).                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationEffectsPolicy 38            | Return policies that differentiate among implementations. The                    |
	//	|                                       | pWinStationInformation argument points to a ULONG variable.                      |
	//	|                                       | WinStationInformationLength MUST be sizeof(ULONG).                               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationType 39                     | Returns the type associated with this WinStation. The pWinStationInformation     |
	//	|                                       | argument points to a ULONG variable. WinStationInformationLength MUST be         |
	//	|                                       | sizeof(ULONG).                                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationInformationEx 40            | Retrieves extended information on the session, including connect state, flags,   |
	//	|                                       | session's name, connect time, disconnect time, time last input was received      |
	//	|                                       | from the client, logon time, user's username and domain, and the current         |
	//	|                                       | time. pWinStationInformation points to a VARDATA_WIRE structure followed by      |
	//	|                                       | a WINSTATIONINFORMATIONEX structure. The WinStationInformationLength MUST be     |
	//	|                                       | sizeof(VARDATA_WIRE) + sizeof(WINSTATIONINFORMATIONEX). The Size field in the    |
	//	|                                       | VARDATA_WIRE structure MUST be set to sizeof(WINSTATIONINFORMATIONEX) and the    |
	//	|                                       | Offset set to sizeof(VARDATA_WIRE).                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	WinStationInformationClass uint32 `idl:"name:WinStationInformationClass" json:"win_station_information_class"`
	// pWinStationInformation: Pointer to buffer allocated by the caller in which to retrieve
	// the data. The data type or structure that pWinStationInformation points to is determined
	// by the value of WinStationInformationClass. See previous sections for what pWinStationInformation
	// SHOULD point to for each class.
	WinStationInformation []byte `idl:"name:pWinStationInformation;size_is:(WinStationInformationLength);pointer:unique" json:"win_station_information"`
	// WinStationInformationLength: Size of the data pointed to by pWinStationInformation,
	// in bytes.
	WinStationInformationLength uint32 `idl:"name:WinStationInformationLength" json:"win_station_information_length"`
}

func (o *WinStationQueryInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationQueryInformationOperation) *xxx_WinStationQueryInformationOperation {
	if op == nil {
		op = &xxx_WinStationQueryInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.LogonID = o.LogonID
	op.WinStationInformationClass = o.WinStationInformationClass
	op.WinStationInformation = o.WinStationInformation
	op.WinStationInformationLength = o.WinStationInformationLength
	return op
}

func (o *WinStationQueryInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationQueryInformationOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.LogonID = op.LogonID
	o.WinStationInformationClass = op.WinStationInformationClass
	o.WinStationInformation = op.WinStationInformation
	o.WinStationInformationLength = op.WinStationInformationLength
}
func (o *WinStationQueryInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationQueryInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationQueryInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationQueryInformationResponse structure represents the RpcWinStationQueryInformation operation response
type WinStationQueryInformationResponse struct {
	// XXX: WinStationInformationLength is an implicit input depedency for output parameters
	WinStationInformationLength uint32 `idl:"name:WinStationInformationLength" json:"win_station_information_length"`

	// pResult: Failure error code if the call to RpcWinStationQueryInformation failed.
	// If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+--------------------------------------+--------------------------------------------------------+
	//	|                                      |                                                        |
	//	|                VALUE                 |                        MEANING                         |
	//	|                                      |                                                        |
	//	+--------------------------------------+--------------------------------------------------------+
	//	+--------------------------------------+--------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000            | Successful completion.                                 |
	//	+--------------------------------------+--------------------------------------------------------+
	//	| STATUS_INVALID_INFO_CLASS 0xC0000003 | The class is not recognized.                           |
	//	+--------------------------------------+--------------------------------------------------------+
	//	| STATUS_BUFFER_TOO_SMALL 0xC0000023   | WinStationInformationLength is too small.              |
	//	+--------------------------------------+--------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022      | The caller does not have permission for the operation. |
	//	+--------------------------------------+--------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// pWinStationInformation: Pointer to buffer allocated by the caller in which to retrieve
	// the data. The data type or structure that pWinStationInformation points to is determined
	// by the value of WinStationInformationClass. See previous sections for what pWinStationInformation
	// SHOULD point to for each class.
	WinStationInformation []byte `idl:"name:pWinStationInformation;size_is:(WinStationInformationLength);pointer:unique" json:"win_station_information"`
	// pReturnLength: Pointer to a variable to receive the size, in bytes, of the data retrieved.
	// If WinStationInformationLength is too small, pReturnLength indicates the correct
	// number of bytes for the caller to allocate.
	ReturnLength uint32 `idl:"name:pReturnLength" json:"return_length"`
	// Return: The RpcWinStationQueryInformation return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationQueryInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationQueryInformationOperation) *xxx_WinStationQueryInformationOperation {
	if op == nil {
		op = &xxx_WinStationQueryInformationOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.WinStationInformationLength == uint32(0) {
		op.WinStationInformationLength = o.WinStationInformationLength
	}

	op.Result = o.Result
	op.WinStationInformation = o.WinStationInformation
	op.ReturnLength = o.ReturnLength
	op.Return = o.Return
	return op
}

func (o *WinStationQueryInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationQueryInformationOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.WinStationInformationLength = op.WinStationInformationLength

	o.Result = op.Result
	o.WinStationInformation = op.WinStationInformation
	o.ReturnLength = op.ReturnLength
	o.Return = op.Return
}
func (o *WinStationQueryInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationQueryInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationQueryInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationSetInformationOperation structure represents the RpcWinStationSetInformation operation
type xxx_WinStationSetInformationOperation struct {
	Server                      *tsts.Server `idl:"name:hServer" json:"server"`
	Result                      uint32       `idl:"name:pResult" json:"result"`
	LogonID                     uint32       `idl:"name:LogonId" json:"logon_id"`
	WinStationInformationClass  uint32       `idl:"name:WinStationInformationClass" json:"win_station_information_class"`
	WinStationInformation       []byte       `idl:"name:pWinStationInformation;size_is:(WinStationInformationLength);pointer:unique" json:"win_station_information"`
	WinStationInformationLength uint32       `idl:"name:WinStationInformationLength" json:"win_station_information_length"`
	Return                      bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationSetInformationOperation) OpNum() int { return 6 }

func (o *xxx_WinStationSetInformationOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationSetInformation"
}

func (o *xxx_WinStationSetInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.WinStationInformation != nil && o.WinStationInformationLength == 0 {
		o.WinStationInformationLength = uint32(len(o.WinStationInformation))
	}
	if o.WinStationInformationLength > uint32(32768) {
		return fmt.Errorf("WinStationInformationLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationSetInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LogonID); err != nil {
			return err
		}
	}
	// WinStationInformationClass {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.WinStationInformationClass); err != nil {
			return err
		}
	}
	// pWinStationInformation {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=WinStationInformationLength](char))
	{
		if o.WinStationInformation != nil || o.WinStationInformationLength > 0 {
			_ptr_pWinStationInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.WinStationInformationLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.WinStationInformation {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.WinStationInformation[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.WinStationInformation); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WinStationInformation, _ptr_pWinStationInformation); err != nil {
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
	// WinStationInformationLength {in} (1:{range=(0,32768), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.WinStationInformationLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationSetInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LogonID); err != nil {
			return err
		}
	}
	// WinStationInformationClass {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.WinStationInformationClass); err != nil {
			return err
		}
	}
	// pWinStationInformation {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=WinStationInformationLength](char))
	{
		_ptr_pWinStationInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.WinStationInformation", sizeInfo[0])
			}
			o.WinStationInformation = make([]byte, sizeInfo[0])
			for i1 := range o.WinStationInformation {
				i1 := i1
				if err := w.ReadData(&o.WinStationInformation[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pWinStationInformation := func(ptr interface{}) { o.WinStationInformation = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.WinStationInformation, _s_pWinStationInformation, _ptr_pWinStationInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WinStationInformationLength {in} (1:{range=(0,32768), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.WinStationInformationLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationSetInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationSetInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pWinStationInformation {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=WinStationInformationLength](char))
	{
		if o.WinStationInformation != nil || o.WinStationInformationLength > 0 {
			_ptr_pWinStationInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.WinStationInformationLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.WinStationInformation {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.WinStationInformation[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.WinStationInformation); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WinStationInformation, _ptr_pWinStationInformation); err != nil {
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
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationSetInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pWinStationInformation {in, out} (1:{pointer=unique, alias=PCHAR}*(1))(2:{alias=CHAR}[dim:0,size_is=WinStationInformationLength](char))
	{
		_ptr_pWinStationInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.WinStationInformation", sizeInfo[0])
			}
			o.WinStationInformation = make([]byte, sizeInfo[0])
			for i1 := range o.WinStationInformation {
				i1 := i1
				if err := w.ReadData(&o.WinStationInformation[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pWinStationInformation := func(ptr interface{}) { o.WinStationInformation = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.WinStationInformation, _s_pWinStationInformation, _ptr_pWinStationInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationSetInformationRequest structure represents the RpcWinStationSetInformation operation request
type WinStationSetInformationRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. hServer MUST
	// be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// LogonId:  The ID of the session for which to set information.
	LogonID uint32 `idl:"name:LogonId" json:"logon_id"`
	// WinStationInformationClass:  The class of data to set. These values come from the
	// enum type WINSTATIONINFOCLASS. See the following sections for the supported classes.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationConfiguration 1          | Merges configuration data into the terminal server session's data. The           |
	//	|                                    | pWinStationInformation argument points to a WINSTACONFIGWIRE structure           |
	//	|                                    | followed by a USERCONFIG structure. The WinStationInformationLength MUST be      |
	//	|                                    | sizeof(WINSTACONFIGWIRE) + sizeof(USERCONFIG). The Size field in the USERCONFIG  |
	//	|                                    | structure inside WINSTACONFIGWIRE MUST be set to sizeof(USERCONFIG) and          |
	//	|                                    | the Offset set to sizeof(WINSTACONFIGWIRE). The Size field in the NewFields      |
	//	|                                    | structure inside WINSTACONFIGWIRE MUST be set to 0, and the offset MUST be set   |
	//	|                                    | to sizeof(WINSTACONFIGWIRE) + sizeof(USERCONFIG).                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationPdParams 2               | Sets transport protocol driver parameters. The structure coming into the         |
	//	|                                    | function indicates via SDClass the specific protocol driver for which            |
	//	|                                    | to set the parameter information. The pWinStationInformation argument            |
	//	|                                    | points to a PDPARAMSWIRE structure followed by a PDPARAMS structure. The         |
	//	|                                    | WinStationInformationLength MUST be sizeof(PDPARAMSWIRE) + sizeof(PDPARAMS).     |
	//	|                                    | The Size field in SdClassSpecific field inside PDPARAMSWIRE MUST be set to       |
	//	|                                    | sizeof(PDPARAMS) and the offset MUST be set to sizeof(PDPARAMSWIRE).             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationTrace 9                  | Enables tracing on the lower-level terminal server drivers for this session.     |
	//	|                                    | This MUST be called by a process running as SYSTEM or as an administrator. The   |
	//	|                                    | method performs access checks as defined in sections 3.1.2 and 3.1.3. The method |
	//	|                                    | fails if both checks fail. The pWinStationInformation argument points to a       |
	//	|                                    | TS_TRACE structure, and WinStationInformationLength MUST be sizeof(TS_TRACE).    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationBeep 10                  | Sends a beep to the session. The pWinStationInformation argument points to a     |
	//	|                                    | BEEPINPUT structure, and WinStationInformationLength MUST be sizeof(BEEPINPUT).  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationEncryptionOff 11         | Turns encryption off.<174>                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationEncryptionPerm 12        | Turns encryption permanently on.<175>                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationNtSecurity 13            | Sends logon service in the session a CTRL+ALT+DEL message.<176> The              |
	//	|                                    | pWinStationInformation argument and WinStationInformationLength are not used     |
	//	|                                    | for this class. Set them to dummy valid data, however, as there always has to be |
	//	|                                    | something in these parameters.                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationInitialProgram 17        | Not used.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationSystemTrace 19           | Enables global tracing on the lower-level terminal server drivers for            |
	//	|                                    | this session. This MUST be called by a process running as SYSTEM or as an        |
	//	|                                    | administrator. The method performs access checks as defined in sections 3.1.2    |
	//	|                                    | and 3.1.3. The method fails if both checks fail. The pWinStationInformation      |
	//	|                                    | argument points to a TS_TRACE structure, and WinStationInformationLength MUST be |
	//	|                                    | sizeof(TS_TRACE).                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationClientData 21            | Sends data to the terminal server client. WinStationInformationLength            |
	//	|                                    | represents the complete length of all items to send and MUST be at               |
	//	|                                    | least sizeof(VARDATA_WIRE) + sizeof(WINSTATIONCLIENTDATA ). Otherwise,           |
	//	|                                    | STATUS_INFO_LENGTH_MISMATCH (as specified in [MS-ERREF]) is returned in          |
	//	|                                    | pResult. If the data is larger than this but still less than what is expected,   |
	//	|                                    | STATUS_INVALID_USER_BUFFER (as specified in [MS-ERREF]) is returned in pResult.  |
	//	|                                    | The pWinStationInformation points to a VARDATA_WIRE structure followed by        |
	//	|                                    | the WINSTATIONCLIENTDATA structure itself. The Size field in the VARDATA_WIRE    |
	//	|                                    | structure MUST be set to sizeof(WINSTATIONCLIENTDATA) and the Offset set to      |
	//	|                                    | sizeof(VARDATA_WIRE).                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationSecureDesktopEnter 22    | Not used.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationSecureDesktopExit 23     | Not used.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationShadowInfo 26            | Not used.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationLockedState 28           | Notifies processes of the new locked state of the session. TRUE or               |
	//	|                                    | FALSE. The pWinStationInformation argument points to a BOOL variable.            |
	//	|                                    | WinStationInformationLength MUST be sizeof (BOOL).                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| WinStationDisallowAutoReconnect 32 | Allows or disallows auto-reconnect behavior for this session, TRUE or FALSE.     |
	//	|                                    | This MUST be called by a process running as SYSTEM. The pWinStationInformation   |
	//	|                                    | argument points to a BOOL variable. WinStationInformationLength MUST be sizeof   |
	//	|                                    | (BOOL).                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	WinStationInformationClass uint32 `idl:"name:WinStationInformationClass" json:"win_station_information_class"`
	// pWinStationInformation: Pointer to buffer allocated by the caller in which the data
	// for the operation is located. The data type or structure to which pWinStationInformation
	// points is determined by the value of WinStationInformationClass.
	WinStationInformation []byte `idl:"name:pWinStationInformation;size_is:(WinStationInformationLength);pointer:unique" json:"win_station_information"`
	// WinStationInformationLength: Size of the data pointed to by pWinStationInformation,
	// in bytes.
	WinStationInformationLength uint32 `idl:"name:WinStationInformationLength" json:"win_station_information_length"`
}

func (o *WinStationSetInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationSetInformationOperation) *xxx_WinStationSetInformationOperation {
	if op == nil {
		op = &xxx_WinStationSetInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.LogonID = o.LogonID
	op.WinStationInformationClass = o.WinStationInformationClass
	op.WinStationInformation = o.WinStationInformation
	op.WinStationInformationLength = o.WinStationInformationLength
	return op
}

func (o *WinStationSetInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationSetInformationOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.LogonID = op.LogonID
	o.WinStationInformationClass = op.WinStationInformationClass
	o.WinStationInformation = op.WinStationInformation
	o.WinStationInformationLength = op.WinStationInformationLength
}
func (o *WinStationSetInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationSetInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationSetInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationSetInformationResponse structure represents the RpcWinStationSetInformation operation response
type WinStationSetInformationResponse struct {
	// XXX: WinStationInformationLength is an implicit input depedency for output parameters
	WinStationInformationLength uint32 `idl:"name:WinStationInformationLength" json:"win_station_information_length"`

	// pResult: Failure error code if the call to RpcWinStationSetInformation failed. If
	// the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+--------------------------------------+--------------------------------------------------------+
	//	|                                      |                                                        |
	//	|                VALUE                 |                        MEANING                         |
	//	|                                      |                                                        |
	//	+--------------------------------------+--------------------------------------------------------+
	//	+--------------------------------------+--------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000            | Successful completion.                                 |
	//	+--------------------------------------+--------------------------------------------------------+
	//	| STATUS_INVALID_INFO_CLASS 0xC0000003 | The class is not recognized.                           |
	//	+--------------------------------------+--------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022      | The caller does not have permission for the operation. |
	//	+--------------------------------------+--------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// pWinStationInformation: Pointer to buffer allocated by the caller in which the data
	// for the operation is located. The data type or structure to which pWinStationInformation
	// points is determined by the value of WinStationInformationClass.
	WinStationInformation []byte `idl:"name:pWinStationInformation;size_is:(WinStationInformationLength);pointer:unique" json:"win_station_information"`
	// Return: The RpcWinStationSetInformation return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationSetInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationSetInformationOperation) *xxx_WinStationSetInformationOperation {
	if op == nil {
		op = &xxx_WinStationSetInformationOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.WinStationInformationLength == uint32(0) {
		op.WinStationInformationLength = o.WinStationInformationLength
	}

	op.Result = o.Result
	op.WinStationInformation = o.WinStationInformation
	op.Return = o.Return
	return op
}

func (o *WinStationSetInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationSetInformationOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.WinStationInformationLength = op.WinStationInformationLength

	o.Result = op.Result
	o.WinStationInformation = op.WinStationInformation
	o.Return = op.Return
}
func (o *WinStationSetInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationSetInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationSetInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationSendMessageOperation structure represents the RpcWinStationSendMessage operation
type xxx_WinStationSendMessageOperation struct {
	Server        *tsts.Server `idl:"name:hServer" json:"server"`
	Result        uint32       `idl:"name:pResult" json:"result"`
	LogonID       uint32       `idl:"name:LogonId" json:"logon_id"`
	Title         string       `idl:"name:pTitle;size_is:(TitleLength)" json:"title"`
	TitleLength   uint32       `idl:"name:TitleLength" json:"title_length"`
	Message       string       `idl:"name:pMessage;size_is:(MessageLength)" json:"message"`
	MessageLength uint32       `idl:"name:MessageLength" json:"message_length"`
	Style         uint32       `idl:"name:Style" json:"style"`
	Timeout       uint32       `idl:"name:Timeout" json:"timeout"`
	Response      uint32       `idl:"name:pResponse" json:"response"`
	DoNotWait     bool         `idl:"name:DoNotWait" json:"do_not_wait"`
	Return        bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationSendMessageOperation) OpNum() int { return 7 }

func (o *xxx_WinStationSendMessageOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationSendMessage"
}

func (o *xxx_WinStationSendMessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Title != "" && o.TitleLength == 0 {
		o.TitleLength = uint32(ndr.UTF16Len(o.Title))
	}
	if o.Message != "" && o.MessageLength == 0 {
		o.MessageLength = uint32(ndr.UTF16Len(o.Message))
	}
	if o.TitleLength > uint32(1024) {
		return fmt.Errorf("TitleLength is out of range")
	}
	if o.MessageLength > uint32(1024) {
		return fmt.Errorf("MessageLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationSendMessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LogonID); err != nil {
			return err
		}
	}
	// pTitle {in} (1:{alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=TitleLength,string](wchar))
	{
		dimSize1 := uint64(o.TitleLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
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
	}
	// TitleLength {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TitleLength); err != nil {
			return err
		}
	}
	// pMessage {in} (1:{alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=MessageLength,string](wchar))
	{
		dimSize1 := uint64(o.MessageLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_Message_buf := utf16.Encode([]rune(o.Message))
		if uint64(len(_Message_buf)) > sizeInfo[0] {
			_Message_buf = _Message_buf[:sizeInfo[0]]
		}
		for i1 := range _Message_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Message_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Message_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// MessageLength {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MessageLength); err != nil {
			return err
		}
	}
	// Style {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Style); err != nil {
			return err
		}
	}
	// Timeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// DoNotWait {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.DoNotWait); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationSendMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LogonID); err != nil {
			return err
		}
	}
	// pTitle {in} (1:{alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=TitleLength,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
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
	}
	// TitleLength {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TitleLength); err != nil {
			return err
		}
	}
	// pMessage {in} (1:{alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=MessageLength,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Message_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Message_buf", sizeInfo[0])
		}
		_Message_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Message_buf {
			i1 := i1
			if err := w.ReadData(&_Message_buf[i1]); err != nil {
				return err
			}
		}
		o.Message = strings.TrimRight(string(utf16.Decode(_Message_buf)), ndr.ZeroString)
	}
	// MessageLength {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MessageLength); err != nil {
			return err
		}
	}
	// Style {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Style); err != nil {
			return err
		}
	}
	// Timeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// DoNotWait {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.DoNotWait); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationSendMessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationSendMessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pResponse {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Response); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationSendMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pResponse {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Response); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationSendMessageRequest structure represents the RpcWinStationSendMessage operation request
type WinStationSendMessageRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// LogonId:  The session ID of the session on which to display the message box.
	LogonID uint32 `idl:"name:LogonId" json:"logon_id"`
	// pTitle:  Pointer to the title for the message box to display.
	Title string `idl:"name:pTitle;size_is:(TitleLength)" json:"title"`
	// TitleLength:  The length, in bytes, of the title to display.
	TitleLength uint32 `idl:"name:TitleLength" json:"title_length"`
	// pMessage:  Pointer to the message to display.
	Message string `idl:"name:pMessage;size_is:(MessageLength)" json:"message"`
	// MessageLength:  The length, in bytes, of the message to display in the specified
	// session.
	MessageLength uint32 `idl:"name:MessageLength" json:"message_length"`
	// Style:  Can be any value that the standard MessageBox() method's Style parameter
	// takes. For more information, see [MSDN-MSGBOX].
	Style uint32 `idl:"name:Style" json:"style"`
	// Timeout:  The response time-out, in seconds. If the message box is not responded
	// to in Timeout seconds, a response code of IDTIMEOUT MUST be returned in pResponse
	// to indicate that the message box timed out. This time-out value is managed by another
	// system component which dismisses the message box if no user input is entered during
	// this interval.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
	// DoNotWait:  If set to TRUE, do not wait for the response to the message. On return,
	// if no errors occur in queuing the message, the pResponse parameter will be set to
	// IDASYNC.
	DoNotWait bool `idl:"name:DoNotWait" json:"do_not_wait"`
}

func (o *WinStationSendMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationSendMessageOperation) *xxx_WinStationSendMessageOperation {
	if op == nil {
		op = &xxx_WinStationSendMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.LogonID = o.LogonID
	op.Title = o.Title
	op.TitleLength = o.TitleLength
	op.Message = o.Message
	op.MessageLength = o.MessageLength
	op.Style = o.Style
	op.Timeout = o.Timeout
	op.DoNotWait = o.DoNotWait
	return op
}

func (o *WinStationSendMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationSendMessageOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.LogonID = op.LogonID
	o.Title = op.Title
	o.TitleLength = op.TitleLength
	o.Message = op.Message
	o.MessageLength = op.MessageLength
	o.Style = op.Style
	o.Timeout = op.Timeout
	o.DoNotWait = op.DoNotWait
}
func (o *WinStationSendMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationSendMessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationSendMessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationSendMessageResponse structure represents the RpcWinStationSendMessage operation response
type WinStationSendMessageResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationSendMessage failed. If
	// the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+---------------------------------+-----------------------------------------------------+
	//	|                                 |                                                     |
	//	|              VALUE              |                       MEANING                       |
	//	|                                 |                                                     |
	//	+---------------------------------+-----------------------------------------------------+
	//	+---------------------------------+-----------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000       | The call succeeded.                                 |
	//	+---------------------------------+-----------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022 | The caller does not have WINSTATION_MSG permission. |
	//	+---------------------------------+-----------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// pResponse:  The return code from the MessageBox method. This value will be a standard
	// MessageBox return value. For more information, see [MSDN-MSGBOX].
	Response uint32 `idl:"name:pResponse" json:"response"`
	// Return: The RpcWinStationSendMessage return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationSendMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationSendMessageOperation) *xxx_WinStationSendMessageOperation {
	if op == nil {
		op = &xxx_WinStationSendMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Response = o.Response
	op.Return = o.Return
	return op
}

func (o *WinStationSendMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationSendMessageOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Response = op.Response
	o.Return = op.Return
}
func (o *WinStationSendMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationSendMessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationSendMessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LogonIDFromWinStationNameOperation structure represents the RpcLogonIdFromWinStationName operation
type xxx_LogonIDFromWinStationNameOperation struct {
	Server         *tsts.Server `idl:"name:hServer" json:"server"`
	Result         uint32       `idl:"name:pResult" json:"result"`
	WinStationName string       `idl:"name:pWinStationName;size_is:(NameSize)" json:"win_station_name"`
	NameSize       uint32       `idl:"name:NameSize" json:"name_size"`
	LogonID        uint32       `idl:"name:pLogonId" json:"logon_id"`
	Return         bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_LogonIDFromWinStationNameOperation) OpNum() int { return 8 }

func (o *xxx_LogonIDFromWinStationNameOperation) OpName() string {
	return "/IcaApi/v1/RpcLogonIdFromWinStationName"
}

func (o *xxx_LogonIDFromWinStationNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.WinStationName != "" && o.NameSize == 0 {
		o.NameSize = uint32(ndr.UTF16Len(o.WinStationName))
	}
	if o.NameSize > uint32(256) {
		return fmt.Errorf("NameSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LogonIDFromWinStationNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pWinStationName {in} (1:{alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameSize,string](wchar))
	{
		dimSize1 := uint64(o.NameSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_WinStationName_buf := utf16.Encode([]rune(o.WinStationName))
		if uint64(len(_WinStationName_buf)) > sizeInfo[0] {
			_WinStationName_buf = _WinStationName_buf[:sizeInfo[0]]
		}
		for i1 := range _WinStationName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_WinStationName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_WinStationName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// NameSize {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LogonIDFromWinStationNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pWinStationName {in} (1:{alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _WinStationName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _WinStationName_buf", sizeInfo[0])
		}
		_WinStationName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _WinStationName_buf {
			i1 := i1
			if err := w.ReadData(&_WinStationName_buf[i1]); err != nil {
				return err
			}
		}
		o.WinStationName = strings.TrimRight(string(utf16.Decode(_WinStationName_buf)), ndr.ZeroString)
	}
	// NameSize {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LogonIDFromWinStationNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LogonIDFromWinStationNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pLogonId {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LogonID); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LogonIDFromWinStationNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pLogonId {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LogonID); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LogonIDFromWinStationNameRequest structure represents the RpcLogonIdFromWinStationName operation request
type LogonIDFromWinStationNameRequest struct {
	// hServer: The Handle to the server object. This is of type SERVER_HANDLE. hServer
	// MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// pWinStationName:  The pointer to a buffer holding the session name.
	WinStationName string `idl:"name:pWinStationName;size_is:(NameSize)" json:"win_station_name"`
	// NameSize:  The length of the string in characters pointed to by pWinStationName
	// including the terminating NULL character. MUST be less than or equal to WINSTATIONNAME_LENGTH.
	NameSize uint32 `idl:"name:NameSize" json:"name_size"`
}

func (o *LogonIDFromWinStationNameRequest) xxx_ToOp(ctx context.Context, op *xxx_LogonIDFromWinStationNameOperation) *xxx_LogonIDFromWinStationNameOperation {
	if op == nil {
		op = &xxx_LogonIDFromWinStationNameOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.WinStationName = o.WinStationName
	op.NameSize = o.NameSize
	return op
}

func (o *LogonIDFromWinStationNameRequest) xxx_FromOp(ctx context.Context, op *xxx_LogonIDFromWinStationNameOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.WinStationName = op.WinStationName
	o.NameSize = op.NameSize
}
func (o *LogonIDFromWinStationNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LogonIDFromWinStationNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LogonIDFromWinStationNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LogonIDFromWinStationNameResponse structure represents the RpcLogonIdFromWinStationName operation response
type LogonIDFromWinStationNameResponse struct {
	// pResult:  The failure error code if the call to RpcLogonIdFromWinStationName fails.
	// If the call is successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+---------------------------------+--------------------------------------------------------+
	//	|                                 |                                                        |
	//	|              VALUE              |                        MEANING                         |
	//	|                                 |                                                        |
	//	+---------------------------------+--------------------------------------------------------+
	//	+---------------------------------+--------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000       | The call is successful.                                |
	//	+---------------------------------+--------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022 | The caller does not have permission for the operation. |
	//	+---------------------------------+--------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// pLogonId: The matching session ID for the session specified by pWinStationName.
	LogonID uint32 `idl:"name:pLogonId" json:"logon_id"`
	// Return: The RpcLogonIdFromWinStationName return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *LogonIDFromWinStationNameResponse) xxx_ToOp(ctx context.Context, op *xxx_LogonIDFromWinStationNameOperation) *xxx_LogonIDFromWinStationNameOperation {
	if op == nil {
		op = &xxx_LogonIDFromWinStationNameOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.LogonID = o.LogonID
	op.Return = o.Return
	return op
}

func (o *LogonIDFromWinStationNameResponse) xxx_FromOp(ctx context.Context, op *xxx_LogonIDFromWinStationNameOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.LogonID = op.LogonID
	o.Return = op.Return
}
func (o *LogonIDFromWinStationNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LogonIDFromWinStationNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LogonIDFromWinStationNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationNameFromLogonIDOperation structure represents the RpcWinStationNameFromLogonId operation
type xxx_WinStationNameFromLogonIDOperation struct {
	Server         *tsts.Server `idl:"name:hServer" json:"server"`
	Result         uint32       `idl:"name:pResult" json:"result"`
	LoginID        uint32       `idl:"name:LoginId" json:"login_id"`
	WinStationName string       `idl:"name:pWinStationName;size_is:(NameSize)" json:"win_station_name"`
	NameSize       uint32       `idl:"name:NameSize" json:"name_size"`
	Return         bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationNameFromLogonIDOperation) OpNum() int { return 9 }

func (o *xxx_WinStationNameFromLogonIDOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationNameFromLogonId"
}

func (o *xxx_WinStationNameFromLogonIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.WinStationName != "" && o.NameSize == 0 {
		o.NameSize = uint32(ndr.UTF16Len(o.WinStationName))
	}
	if o.NameSize > uint32(256) {
		return fmt.Errorf("NameSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationNameFromLogonIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LoginId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LoginID); err != nil {
			return err
		}
	}
	// pWinStationName {in, out} (1:{alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameSize,string](wchar))
	{
		dimSize1 := uint64(o.NameSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_WinStationName_buf := utf16.Encode([]rune(o.WinStationName))
		if uint64(len(_WinStationName_buf)) > sizeInfo[0] {
			_WinStationName_buf = _WinStationName_buf[:sizeInfo[0]]
		}
		for i1 := range _WinStationName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_WinStationName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_WinStationName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// NameSize {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationNameFromLogonIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LoginId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LoginID); err != nil {
			return err
		}
	}
	// pWinStationName {in, out} (1:{alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _WinStationName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _WinStationName_buf", sizeInfo[0])
		}
		_WinStationName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _WinStationName_buf {
			i1 := i1
			if err := w.ReadData(&_WinStationName_buf[i1]); err != nil {
				return err
			}
		}
		o.WinStationName = strings.TrimRight(string(utf16.Decode(_WinStationName_buf)), ndr.ZeroString)
	}
	// NameSize {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationNameFromLogonIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationNameFromLogonIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pWinStationName {in, out} (1:{alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameSize,string](wchar))
	{
		dimSize1 := uint64(o.NameSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_WinStationName_buf := utf16.Encode([]rune(o.WinStationName))
		if uint64(len(_WinStationName_buf)) > sizeInfo[0] {
			_WinStationName_buf = _WinStationName_buf[:sizeInfo[0]]
		}
		for i1 := range _WinStationName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_WinStationName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_WinStationName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationNameFromLogonIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pWinStationName {in, out} (1:{alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _WinStationName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _WinStationName_buf", sizeInfo[0])
		}
		_WinStationName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _WinStationName_buf {
			i1 := i1
			if err := w.ReadData(&_WinStationName_buf[i1]); err != nil {
				return err
			}
		}
		o.WinStationName = strings.TrimRight(string(utf16.Decode(_WinStationName_buf)), ndr.ZeroString)
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationNameFromLogonIDRequest structure represents the RpcWinStationNameFromLogonId operation request
type WinStationNameFromLogonIDRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// LoginId:  The ID of the session for which to retrieve the WinStation name.
	LoginID uint32 `idl:"name:LoginId" json:"login_id"`
	// pWinStationName:  Pointer to a buffer holding the session name. The length of the
	// buffer MUST be equal to or greater than (WINSTATIONNAME_LENGTH + 1).
	WinStationName string `idl:"name:pWinStationName;size_is:(NameSize)" json:"win_station_name"`
	// NameSize: The size, in bytes, of the buffer where the WinStation name will be stored.
	NameSize uint32 `idl:"name:NameSize" json:"name_size"`
}

func (o *WinStationNameFromLogonIDRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationNameFromLogonIDOperation) *xxx_WinStationNameFromLogonIDOperation {
	if op == nil {
		op = &xxx_WinStationNameFromLogonIDOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.LoginID = o.LoginID
	op.WinStationName = o.WinStationName
	op.NameSize = o.NameSize
	return op
}

func (o *WinStationNameFromLogonIDRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationNameFromLogonIDOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.LoginID = op.LoginID
	o.WinStationName = op.WinStationName
	o.NameSize = op.NameSize
}
func (o *WinStationNameFromLogonIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationNameFromLogonIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationNameFromLogonIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationNameFromLogonIDResponse structure represents the RpcWinStationNameFromLogonId operation response
type WinStationNameFromLogonIDResponse struct {
	// XXX: NameSize is an implicit input depedency for output parameters
	NameSize uint32 `idl:"name:NameSize" json:"name_size"`

	// pResult:  Failure error code if the call to RpcWinStationNameFromLogonId failed.
	// If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000           | The call succeeded.                                                              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_INVALID_PARAMETER 0xC000000D | NameSize value is less than WINSTATIONNAME_LENGTH + 1.                           |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80071B6E                          | The session does not exist or the caller does not have WINSTATION_QUERY          |
	//	|                                     | permission.                                                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// pWinStationName:  Pointer to a buffer holding the session name. The length of the
	// buffer MUST be equal to or greater than (WINSTATIONNAME_LENGTH + 1).
	WinStationName string `idl:"name:pWinStationName;size_is:(NameSize)" json:"win_station_name"`
	// Return: The RpcWinStationNameFromLogonId return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationNameFromLogonIDResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationNameFromLogonIDOperation) *xxx_WinStationNameFromLogonIDOperation {
	if op == nil {
		op = &xxx_WinStationNameFromLogonIDOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.NameSize == uint32(0) {
		op.NameSize = o.NameSize
	}

	op.Result = o.Result
	op.WinStationName = o.WinStationName
	op.Return = o.Return
	return op
}

func (o *WinStationNameFromLogonIDResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationNameFromLogonIDOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.NameSize = op.NameSize

	o.Result = op.Result
	o.WinStationName = op.WinStationName
	o.Return = op.Return
}
func (o *WinStationNameFromLogonIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationNameFromLogonIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationNameFromLogonIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationConnectOperation structure represents the RpcWinStationConnect operation
type xxx_WinStationConnectOperation struct {
	Server         *tsts.Server `idl:"name:hServer" json:"server"`
	Result         uint32       `idl:"name:pResult" json:"result"`
	ClientLogonID  uint32       `idl:"name:ClientLogonId" json:"client_logon_id"`
	ConnectLogonID uint32       `idl:"name:ConnectLogonId" json:"connect_logon_id"`
	TargetLogonID  uint32       `idl:"name:TargetLogonId" json:"target_logon_id"`
	Password       string       `idl:"name:pPassword;size_is:(PasswordSize)" json:"password"`
	PasswordSize   uint32       `idl:"name:PasswordSize" json:"password_size"`
	Wait           bool         `idl:"name:Wait" json:"wait"`
	Return         bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationConnectOperation) OpNum() int { return 10 }

func (o *xxx_WinStationConnectOperation) OpName() string { return "/IcaApi/v1/RpcWinStationConnect" }

func (o *xxx_WinStationConnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Password != "" && o.PasswordSize == 0 {
		o.PasswordSize = uint32(ndr.UTF16Len(o.Password))
	}
	if o.PasswordSize > uint32(1024) {
		return fmt.Errorf("PasswordSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationConnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ClientLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientLogonID); err != nil {
			return err
		}
	}
	// ConnectLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ConnectLogonID); err != nil {
			return err
		}
	}
	// TargetLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TargetLogonID); err != nil {
			return err
		}
	}
	// pPassword {in} (1:{alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=PasswordSize,string](wchar))
	{
		dimSize1 := uint64(o.PasswordSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_Password_buf := utf16.Encode([]rune(o.Password))
		if uint64(len(_Password_buf)) > sizeInfo[0] {
			_Password_buf = _Password_buf[:sizeInfo[0]]
		}
		for i1 := range _Password_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Password_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Password_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// PasswordSize {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PasswordSize); err != nil {
			return err
		}
	}
	// Wait {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Wait); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationConnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ClientLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientLogonID); err != nil {
			return err
		}
	}
	// ConnectLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ConnectLogonID); err != nil {
			return err
		}
	}
	// TargetLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TargetLogonID); err != nil {
			return err
		}
	}
	// pPassword {in} (1:{alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=PasswordSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Password_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Password_buf", sizeInfo[0])
		}
		_Password_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Password_buf {
			i1 := i1
			if err := w.ReadData(&_Password_buf[i1]); err != nil {
				return err
			}
		}
		o.Password = strings.TrimRight(string(utf16.Decode(_Password_buf)), ndr.ZeroString)
	}
	// PasswordSize {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PasswordSize); err != nil {
			return err
		}
	}
	// Wait {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Wait); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationConnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationConnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationConnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationConnectRequest structure represents the RpcWinStationConnect operation request
type WinStationConnectRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// ClientLogonId:  The session ID of the caller of this method.
	ClientLogonID uint32 `idl:"name:ClientLogonId" json:"client_logon_id"`
	// ConnectLogonId: The ID of the session from which the connection is being made. This
	// MUST be the same as ClientLogonId and MUST be an existing session ID. The user MUST
	// be logged on. To indicate the current session, this MUST be LOGONID_CURRENT.
	ConnectLogonID uint32 `idl:"name:ConnectLogonId" json:"connect_logon_id"`
	// TargetLogonId:  The session ID of the session to which the connection is being made.
	// Cannot be the same as ConnectLogonId and MUST be an existing session ID.
	TargetLogonID uint32 `idl:"name:TargetLogonId" json:"target_logon_id"`
	// pPassword:  The password of TargetLogonId's session. The password MUST be valid.
	// The password MAY be NULL if the same user is making the call as the user logged on
	// to TargetLogonId's session.
	Password string `idl:"name:pPassword;size_is:(PasswordSize)" json:"password"`
	// PasswordSize: The length of the string pPassword in characters including the terminating
	// NULL character.
	PasswordSize uint32 `idl:"name:PasswordSize" json:"password_size"`
	// Wait:  TRUE indicates to wait for the connection to complete, FALSE otherwise.<181>
	Wait bool `idl:"name:Wait" json:"wait"`
}

func (o *WinStationConnectRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationConnectOperation) *xxx_WinStationConnectOperation {
	if op == nil {
		op = &xxx_WinStationConnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ClientLogonID = o.ClientLogonID
	op.ConnectLogonID = o.ConnectLogonID
	op.TargetLogonID = o.TargetLogonID
	op.Password = o.Password
	op.PasswordSize = o.PasswordSize
	op.Wait = o.Wait
	return op
}

func (o *WinStationConnectRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationConnectOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ClientLogonID = op.ClientLogonID
	o.ConnectLogonID = op.ConnectLogonID
	o.TargetLogonID = op.TargetLogonID
	o.Password = op.Password
	o.PasswordSize = op.PasswordSize
	o.Wait = op.Wait
}
func (o *WinStationConnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationConnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationConnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationConnectResponse structure represents the RpcWinStationConnect operation response
type WinStationConnectResponse struct {
	// pResult: If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000);
	// otherwise, it MUST be an implementation-specific negative value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationConnect return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationConnectResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationConnectOperation) *xxx_WinStationConnectOperation {
	if op == nil {
		op = &xxx_WinStationConnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationConnectResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationConnectOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationConnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationConnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationConnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationDisconnectOperation structure represents the RpcWinStationDisconnect operation
type xxx_WinStationDisconnectOperation struct {
	Server  *tsts.Server `idl:"name:hServer" json:"server"`
	Result  uint32       `idl:"name:pResult" json:"result"`
	LogonID uint32       `idl:"name:LogonId" json:"logon_id"`
	Wait    bool         `idl:"name:bWait" json:"wait"`
	Return  bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationDisconnectOperation) OpNum() int { return 13 }

func (o *xxx_WinStationDisconnectOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationDisconnect"
}

func (o *xxx_WinStationDisconnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationDisconnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LogonID); err != nil {
			return err
		}
	}
	// bWait {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Wait); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationDisconnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LogonID); err != nil {
			return err
		}
	}
	// bWait {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Wait); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationDisconnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationDisconnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationDisconnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationDisconnectRequest structure represents the RpcWinStationDisconnect operation request
type WinStationDisconnectRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. hServer MUST
	// be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// LogonId:  The ID of the session to disconnect. Can be LOGONID_CURRENT to indicate
	// the current session.
	LogonID uint32 `idl:"name:LogonId" json:"logon_id"`
	// bWait:  TRUE to wait for the disconnect to complete before returning, FALSE otherwise.<183>
	Wait bool `idl:"name:bWait" json:"wait"`
}

func (o *WinStationDisconnectRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationDisconnectOperation) *xxx_WinStationDisconnectOperation {
	if op == nil {
		op = &xxx_WinStationDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.LogonID = o.LogonID
	op.Wait = o.Wait
	return op
}

func (o *WinStationDisconnectRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationDisconnectOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.LogonID = op.LogonID
	o.Wait = op.Wait
}
func (o *WinStationDisconnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationDisconnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationDisconnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationDisconnectResponse structure represents the RpcWinStationDisconnect operation response
type WinStationDisconnectResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationDisconnect failed. If the
	// call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+---------------------------------+----------------------------------------------------------------+
	//	|                                 |                                                                |
	//	|              VALUE              |                            MEANING                             |
	//	|                                 |                                                                |
	//	+---------------------------------+----------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000       | The call succeeded.                                            |
	//	+---------------------------------+----------------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022 | The caller does not have permission to disconnect the session. |
	//	+---------------------------------+----------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationDisconnect return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationDisconnectResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationDisconnectOperation) *xxx_WinStationDisconnectOperation {
	if op == nil {
		op = &xxx_WinStationDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationDisconnectResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationDisconnectOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationDisconnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationDisconnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationDisconnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationResetOperation structure represents the RpcWinStationReset operation
type xxx_WinStationResetOperation struct {
	Server  *tsts.Server `idl:"name:hServer" json:"server"`
	Result  uint32       `idl:"name:pResult" json:"result"`
	LogonID uint32       `idl:"name:LogonId" json:"logon_id"`
	Wait    bool         `idl:"name:bWait" json:"wait"`
	Return  bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationResetOperation) OpNum() int { return 14 }

func (o *xxx_WinStationResetOperation) OpName() string { return "/IcaApi/v1/RpcWinStationReset" }

func (o *xxx_WinStationResetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationResetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LogonID); err != nil {
			return err
		}
	}
	// bWait {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Wait); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationResetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LogonID); err != nil {
			return err
		}
	}
	// bWait {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Wait); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationResetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationResetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationResetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationResetRequest structure represents the RpcWinStationReset operation request
type WinStationResetRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. hServer MUST
	// be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// LogonId:  The ID of the session to reset.
	LogonID uint32 `idl:"name:LogonId" json:"logon_id"`
	// bWait: TRUE to wait for the disconnect to complete before returning, FALSE otherwise.
	Wait bool `idl:"name:bWait" json:"wait"`
}

func (o *WinStationResetRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationResetOperation) *xxx_WinStationResetOperation {
	if op == nil {
		op = &xxx_WinStationResetOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.LogonID = o.LogonID
	op.Wait = o.Wait
	return op
}

func (o *WinStationResetRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationResetOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.LogonID = op.LogonID
	o.Wait = op.Wait
}
func (o *WinStationResetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationResetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationResetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationResetResponse structure represents the RpcWinStationReset operation response
type WinStationResetResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationReset failed. If the call
	// was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+---------------------------------+-------------------------------------------------------+
	//	|                                 |                                                       |
	//	|              VALUE              |                        MEANING                        |
	//	|                                 |                                                       |
	//	+---------------------------------+-------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000       | The call succeeded.                                   |
	//	+---------------------------------+-------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022 | The caller does not have WINSTATION_RESET permission. |
	//	+---------------------------------+-------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationReset return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationResetResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationResetOperation) *xxx_WinStationResetOperation {
	if op == nil {
		op = &xxx_WinStationResetOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationResetResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationResetOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationResetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationResetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationResetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationShutdownSystemOperation structure represents the RpcWinStationShutdownSystem operation
type xxx_WinStationShutdownSystemOperation struct {
	Server        *tsts.Server `idl:"name:hServer" json:"server"`
	Result        uint32       `idl:"name:pResult" json:"result"`
	ClientLogonID uint32       `idl:"name:ClientLogonId" json:"client_logon_id"`
	ShutdownFlags uint32       `idl:"name:ShutdownFlags" json:"shutdown_flags"`
	Return        bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationShutdownSystemOperation) OpNum() int { return 15 }

func (o *xxx_WinStationShutdownSystemOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationShutdownSystem"
}

func (o *xxx_WinStationShutdownSystemOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShutdownSystemOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ClientLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientLogonID); err != nil {
			return err
		}
	}
	// ShutdownFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ShutdownFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShutdownSystemOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ClientLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientLogonID); err != nil {
			return err
		}
	}
	// ShutdownFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ShutdownFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShutdownSystemOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShutdownSystemOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShutdownSystemOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationShutdownSystemRequest structure represents the RpcWinStationShutdownSystem operation request
type WinStationShutdownSystemRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// ClientLogonId: The session requesting to shut down the system. Ignored when the RPC
	// call is remote.
	ClientLogonID uint32 `idl:"name:ClientLogonId" json:"client_logon_id"`
	// ShutdownFlags:  Shutdown flags. It MUST be any bitwise OR combination of the following
	// flags.
	//
	//	+-------------------------+----------------------------+
	//	|                         |                            |
	//	|          VALUE          |          MEANING           |
	//	|                         |                            |
	//	+-------------------------+----------------------------+
	//	+-------------------------+----------------------------+
	//	| WSD_LOGOFF 0x00000001   | Forces sessions to logoff. |
	//	+-------------------------+----------------------------+
	//	| WSD_SHUTDOWN 0x00000002 | Shuts down the system.     |
	//	+-------------------------+----------------------------+
	//	| WSD_REBOOT 0x00000004   | Reboots after shutdown.    |
	//	+-------------------------+----------------------------+
	//	| WSD_POWEROFF 0x00000008 | Powers off after shutdown. |
	//	+-------------------------+----------------------------+
	ShutdownFlags uint32 `idl:"name:ShutdownFlags" json:"shutdown_flags"`
}

func (o *WinStationShutdownSystemRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationShutdownSystemOperation) *xxx_WinStationShutdownSystemOperation {
	if op == nil {
		op = &xxx_WinStationShutdownSystemOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ClientLogonID = o.ClientLogonID
	op.ShutdownFlags = o.ShutdownFlags
	return op
}

func (o *WinStationShutdownSystemRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationShutdownSystemOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ClientLogonID = op.ClientLogonID
	o.ShutdownFlags = op.ShutdownFlags
}
func (o *WinStationShutdownSystemRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationShutdownSystemRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationShutdownSystemOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationShutdownSystemResponse structure represents the RpcWinStationShutdownSystem operation response
type WinStationShutdownSystemResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationShutdownSystem failed.
	// If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+---------------------------------+--------------------------------------------------------------+
	//	|                                 |                                                              |
	//	|              VALUE              |                           MEANING                            |
	//	|                                 |                                                              |
	//	+---------------------------------+--------------------------------------------------------------+
	//	+---------------------------------+--------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000       | The call succeeded.                                          |
	//	+---------------------------------+--------------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022 | The caller does not have permission to shut down the system. |
	//	+---------------------------------+--------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationShutdownSystem return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationShutdownSystemResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationShutdownSystemOperation) *xxx_WinStationShutdownSystemOperation {
	if op == nil {
		op = &xxx_WinStationShutdownSystemOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationShutdownSystemResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationShutdownSystemOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationShutdownSystemResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationShutdownSystemResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationShutdownSystemOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationWaitSystemEventOperation structure represents the RpcWinStationWaitSystemEvent operation
type xxx_WinStationWaitSystemEventOperation struct {
	Server     *tsts.Server `idl:"name:hServer" json:"server"`
	Result     uint32       `idl:"name:pResult" json:"result"`
	EventMask  uint32       `idl:"name:EventMask" json:"event_mask"`
	EventFlags uint32       `idl:"name:pEventFlags" json:"event_flags"`
	Return     bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationWaitSystemEventOperation) OpNum() int { return 16 }

func (o *xxx_WinStationWaitSystemEventOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationWaitSystemEvent"
}

func (o *xxx_WinStationWaitSystemEventOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationWaitSystemEventOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// EventMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EventMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationWaitSystemEventOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EventMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EventMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationWaitSystemEventOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationWaitSystemEventOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pEventFlags {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EventFlags); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationWaitSystemEventOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pEventFlags {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EventFlags); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationWaitSystemEventRequest structure represents the RpcWinStationWaitSystemEvent operation request
type WinStationWaitSystemEventRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. hServer MUST
	// be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// EventMask:  The mask of events for which to wait. It MUST be any bitwise OR combination
	// of the following except for WEVENT_NONE.
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|                               |                                                                                  |
	//	|             VALUE             |                                     MEANING                                      |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_NONE 0x00000000        | The client requests to clear its event wait block. This MUST be called when      |
	//	|                               | completing waiting for the event. When RpcWinStationCloseServer is called for    |
	//	|                               | hServer, this method and mask value is called on the client's behalf.            |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_CREATE 0x00000001      | Wait for a new session to be created.                                            |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_DELETE 0x00000002      | Wait for an existing session to be deleted.                                      |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_RENAME 0x00000004      | Wait for a session to be renamed.                                                |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_CONNECT 0x00000008     | The session connected to a client.                                               |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_DISCONNECT 0x00000010  | A client disconnected from the session.                                          |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_LOGON 0x00000020       | A user logged on to the session.                                                 |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_LOGOFF 0x00000040      | A user logged off from the session.                                              |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_STATECHANGE 0x00000080 | The session state changed.                                                       |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_LICENSE 0x00000100     | The license state changed.<185>                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_ALL 0x7fffffff         | Wait for all event types.                                                        |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| WEVENT_FLUSH 0x80000000       | Release all waiting clients.                                                     |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	EventMask uint32 `idl:"name:EventMask" json:"event_mask"`
}

func (o *WinStationWaitSystemEventRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationWaitSystemEventOperation) *xxx_WinStationWaitSystemEventOperation {
	if op == nil {
		op = &xxx_WinStationWaitSystemEventOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.EventMask = o.EventMask
	return op
}

func (o *WinStationWaitSystemEventRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationWaitSystemEventOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.EventMask = op.EventMask
}
func (o *WinStationWaitSystemEventRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationWaitSystemEventRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationWaitSystemEventOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationWaitSystemEventResponse structure represents the RpcWinStationWaitSystemEvent operation response
type WinStationWaitSystemEventResponse struct {
	// pResult: If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000);
	// otherwise, it MUST be an implementation-specific negative value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// pEventFlags:  Pointer to a variable to receive a bitmask that is a subset of EventMask
	// indicating which events actually occurred during this wait operation.
	EventFlags uint32 `idl:"name:pEventFlags" json:"event_flags"`
	// Return: The RpcWinStationWaitSystemEvent return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationWaitSystemEventResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationWaitSystemEventOperation) *xxx_WinStationWaitSystemEventOperation {
	if op == nil {
		op = &xxx_WinStationWaitSystemEventOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.EventFlags = o.EventFlags
	op.Return = o.Return
	return op
}

func (o *WinStationWaitSystemEventResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationWaitSystemEventOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.EventFlags = op.EventFlags
	o.Return = op.Return
}
func (o *WinStationWaitSystemEventResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationWaitSystemEventResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationWaitSystemEventOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationShadowOperation structure represents the RpcWinStationShadow operation
type xxx_WinStationShadowOperation struct {
	Server           *tsts.Server `idl:"name:hServer" json:"server"`
	Result           uint32       `idl:"name:pResult" json:"result"`
	LogonID          uint32       `idl:"name:LogonId" json:"logon_id"`
	TargetServerName string       `idl:"name:pTargetServerName;size_is:(NameSize);pointer:unique" json:"target_server_name"`
	NameSize         uint32       `idl:"name:NameSize" json:"name_size"`
	TargetLogonID    uint32       `idl:"name:TargetLogonId" json:"target_logon_id"`
	HotKeyVirtualKey uint8        `idl:"name:HotKeyVk" json:"hot_key_virtual_key"`
	HotKeyModifiers  uint16       `idl:"name:HotkeyModifiers" json:"hot_key_modifiers"`
	Return           bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationShadowOperation) OpNum() int { return 17 }

func (o *xxx_WinStationShadowOperation) OpName() string { return "/IcaApi/v1/RpcWinStationShadow" }

func (o *xxx_WinStationShadowOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.TargetServerName != "" && o.NameSize == 0 {
		o.NameSize = uint32(ndr.UTF16Len(o.TargetServerName))
	}
	if o.NameSize > uint32(1024) {
		return fmt.Errorf("NameSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShadowOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LogonID); err != nil {
			return err
		}
	}
	// pTargetServerName {in} (1:{pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameSize,string](wchar))
	{
		if o.TargetServerName != "" || o.NameSize > 0 {
			_ptr_pTargetServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NameSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				_TargetServerName_buf := utf16.Encode([]rune(o.TargetServerName))
				if uint64(len(_TargetServerName_buf)) > sizeInfo[0] {
					_TargetServerName_buf = _TargetServerName_buf[:sizeInfo[0]]
				}
				for i1 := range _TargetServerName_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_TargetServerName_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_TargetServerName_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TargetServerName, _ptr_pTargetServerName); err != nil {
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
	// NameSize {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSize); err != nil {
			return err
		}
	}
	// TargetLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TargetLogonID); err != nil {
			return err
		}
	}
	// HotKeyVk {in} (1:{alias=BYTE}(uchar))
	{
		if err := w.WriteData(o.HotKeyVirtualKey); err != nil {
			return err
		}
	}
	// HotkeyModifiers {in} (1:{alias=USHORT}(uint16))
	{
		if err := w.WriteData(o.HotKeyModifiers); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShadowOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LogonID); err != nil {
			return err
		}
	}
	// pTargetServerName {in} (1:{pointer=unique, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameSize,string](wchar))
	{
		_ptr_pTargetServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _TargetServerName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _TargetServerName_buf", sizeInfo[0])
			}
			_TargetServerName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _TargetServerName_buf {
				i1 := i1
				if err := w.ReadData(&_TargetServerName_buf[i1]); err != nil {
					return err
				}
			}
			o.TargetServerName = strings.TrimRight(string(utf16.Decode(_TargetServerName_buf)), ndr.ZeroString)
			return nil
		})
		_s_pTargetServerName := func(ptr interface{}) { o.TargetServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.TargetServerName, _s_pTargetServerName, _ptr_pTargetServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NameSize {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSize); err != nil {
			return err
		}
	}
	// TargetLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TargetLogonID); err != nil {
			return err
		}
	}
	// HotKeyVk {in} (1:{alias=BYTE}(uchar))
	{
		if err := w.ReadData(&o.HotKeyVirtualKey); err != nil {
			return err
		}
	}
	// HotkeyModifiers {in} (1:{alias=USHORT}(uint16))
	{
		if err := w.ReadData(&o.HotKeyModifiers); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShadowOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShadowOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShadowOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationShadowRequest structure represents the RpcWinStationShadow operation request
type WinStationShadowRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument value MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// LogonId: The ID of the session to shadow from.
	LogonID uint32 `idl:"name:LogonId" json:"logon_id"`
	// pTargetServerName: The shadow target server name. Set pTargetServerName to NULL to
	// indicate the current server.
	TargetServerName string `idl:"name:pTargetServerName;size_is:(NameSize);pointer:unique" json:"target_server_name"`
	// NameSize: The size of the string pTargetServerName, in bytes. MAY be 0 if pTargetServerName
	// is NULL.
	NameSize uint32 `idl:"name:NameSize" json:"name_size"`
	// TargetLogonId: The shadow target session ID.
	TargetLogonID uint32 `idl:"name:TargetLogonId" json:"target_logon_id"`
	// HotKeyVk: The virtual key code of the key to press to stop shadowing. This key is
	// used in combination with the HotkeyModifiers parameter.
	HotKeyVirtualKey uint8 `idl:"name:HotKeyVk" json:"hot_key_virtual_key"`
	// HotkeyModifiers: The virtual modifier that signifies the modifier key, such as shift
	// or control, to press to stop shadowing. The modifier key is used in combination with
	// the key signified by the HotKeyVk parameter. This parameter MAY be any combination
	// of KBDSHIFT, KBDCTRL, and KBDALT to indicate the SHIFT key, the CTRL key, and the
	// ALT key, respectively.
	HotKeyModifiers uint16 `idl:"name:HotkeyModifiers" json:"hot_key_modifiers"`
}

func (o *WinStationShadowRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationShadowOperation) *xxx_WinStationShadowOperation {
	if op == nil {
		op = &xxx_WinStationShadowOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.LogonID = o.LogonID
	op.TargetServerName = o.TargetServerName
	op.NameSize = o.NameSize
	op.TargetLogonID = o.TargetLogonID
	op.HotKeyVirtualKey = o.HotKeyVirtualKey
	op.HotKeyModifiers = o.HotKeyModifiers
	return op
}

func (o *WinStationShadowRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationShadowOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.LogonID = op.LogonID
	o.TargetServerName = op.TargetServerName
	o.NameSize = op.NameSize
	o.TargetLogonID = op.TargetLogonID
	o.HotKeyVirtualKey = op.HotKeyVirtualKey
	o.HotKeyModifiers = op.HotKeyModifiers
}
func (o *WinStationShadowRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationShadowRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationShadowOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationShadowResponse structure represents the RpcWinStationShadow operation response
type WinStationShadowResponse struct {
	// pResult: If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000);
	// otherwise, it MUST be an implementation-specific negative value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationShadow return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationShadowResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationShadowOperation) *xxx_WinStationShadowOperation {
	if op == nil {
		op = &xxx_WinStationShadowOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationShadowResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationShadowOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationShadowResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationShadowResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationShadowOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationBreakPointOperation structure represents the RpcWinStationBreakPoint operation
type xxx_WinStationBreakPointOperation struct {
	Server     *tsts.Server `idl:"name:hServer" json:"server"`
	Result     uint32       `idl:"name:pResult" json:"result"`
	LogonID    uint32       `idl:"name:LogonId" json:"logon_id"`
	KernelFlag bool         `idl:"name:KernelFlag" json:"kernel_flag"`
	Return     bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationBreakPointOperation) OpNum() int { return 29 }

func (o *xxx_WinStationBreakPointOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationBreakPoint"
}

func (o *xxx_WinStationBreakPointOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationBreakPointOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LogonID); err != nil {
			return err
		}
	}
	// KernelFlag {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.KernelFlag); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationBreakPointOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LogonID); err != nil {
			return err
		}
	}
	// KernelFlag {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.KernelFlag); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationBreakPointOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationBreakPointOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationBreakPointOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationBreakPointRequest structure represents the RpcWinStationBreakPoint operation request
type WinStationBreakPointRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// LogonId:  The ID of the session to break into the debugger. If this parameter is
	// -2, the terminal server service MUST break into the debugger instead.<188>
	LogonID uint32 `idl:"name:LogonId" json:"logon_id"`
	// KernelFlag:  Set to TRUE to indicate that the server will break into the debugger
	// in a particular session in kernel mode. If LogonId is -2, the server MUST break into
	// the debugger in user mode.
	KernelFlag bool `idl:"name:KernelFlag" json:"kernel_flag"`
}

func (o *WinStationBreakPointRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationBreakPointOperation) *xxx_WinStationBreakPointOperation {
	if op == nil {
		op = &xxx_WinStationBreakPointOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.LogonID = o.LogonID
	op.KernelFlag = o.KernelFlag
	return op
}

func (o *WinStationBreakPointRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationBreakPointOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.LogonID = op.LogonID
	o.KernelFlag = op.KernelFlag
}
func (o *WinStationBreakPointRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationBreakPointRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationBreakPointOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationBreakPointResponse structure represents the RpcWinStationBreakPoint operation response
type WinStationBreakPointResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationBreakPoint failed. If the
	// call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+---------------------------------+-------------------------------------------------------------------------------+
	//	|                                 |                                                                               |
	//	|              VALUE              |                                    MEANING                                    |
	//	|                                 |                                                                               |
	//	+---------------------------------+-------------------------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000       | The call succeeded.                                                           |
	//	+---------------------------------+-------------------------------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022 | The SeShutdownPrivilege (see [MSDN-PRVLGECNSTS]) privilege cannot be enabled. |
	//	+---------------------------------+-------------------------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationBreakPoint return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationBreakPointResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationBreakPointOperation) *xxx_WinStationBreakPointOperation {
	if op == nil {
		op = &xxx_WinStationBreakPointOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationBreakPointResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationBreakPointOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationBreakPointResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationBreakPointResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationBreakPointOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationReadRegistryOperation structure represents the RpcWinStationReadRegistry operation
type xxx_WinStationReadRegistryOperation struct {
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	Result uint32       `idl:"name:pResult" json:"result"`
	Return bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationReadRegistryOperation) OpNum() int { return 30 }

func (o *xxx_WinStationReadRegistryOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationReadRegistry"
}

func (o *xxx_WinStationReadRegistryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationReadRegistryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_WinStationReadRegistryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationReadRegistryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationReadRegistryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationReadRegistryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationReadRegistryRequest structure represents the RpcWinStationReadRegistry operation request
type WinStationReadRegistryRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. hServer MUST
	// be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
}

func (o *WinStationReadRegistryRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationReadRegistryOperation) *xxx_WinStationReadRegistryOperation {
	if op == nil {
		op = &xxx_WinStationReadRegistryOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	return op
}

func (o *WinStationReadRegistryRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationReadRegistryOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
}
func (o *WinStationReadRegistryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationReadRegistryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationReadRegistryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationReadRegistryResponse structure represents the RpcWinStationReadRegistry operation response
type WinStationReadRegistryResponse struct {
	// pResult:  If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000),
	// as specified in [MS-ERREF]; otherwise, it MUST be an implementation-specific negative
	// value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationReadRegistry return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationReadRegistryResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationReadRegistryOperation) *xxx_WinStationReadRegistryOperation {
	if op == nil {
		op = &xxx_WinStationReadRegistryOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationReadRegistryResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationReadRegistryOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationReadRegistryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationReadRegistryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationReadRegistryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OldRPCWinStationEnumerateProcessesOperation structure represents the OldRpcWinStationEnumerateProcesses operation
type xxx_OldRPCWinStationEnumerateProcessesOperation struct {
	Server        *tsts.Server `idl:"name:hServer" json:"server"`
	Result        uint32       `idl:"name:pResult" json:"result"`
	ProcessBuffer []byte       `idl:"name:pProcessBuffer;size_is:(ByteCount)" json:"process_buffer"`
	ByteCount     uint32       `idl:"name:ByteCount" json:"byte_count"`
	Return        bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_OldRPCWinStationEnumerateProcessesOperation) OpNum() int { return 34 }

func (o *xxx_OldRPCWinStationEnumerateProcessesOperation) OpName() string {
	return "/IcaApi/v1/OldRpcWinStationEnumerateProcesses"
}

func (o *xxx_OldRPCWinStationEnumerateProcessesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ByteCount > uint32(32768) {
		return fmt.Errorf("ByteCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OldRPCWinStationEnumerateProcessesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ByteCount {in} (1:{range=(0,32768), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ByteCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OldRPCWinStationEnumerateProcessesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ByteCount {in} (1:{range=(0,32768), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ByteCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OldRPCWinStationEnumerateProcessesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OldRPCWinStationEnumerateProcessesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pProcessBuffer {out} (1:{alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=ByteCount](uchar))
	{
		dimSize1 := uint64(o.ByteCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ProcessBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ProcessBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ProcessBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OldRPCWinStationEnumerateProcessesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pProcessBuffer {out} (1:{alias=PBYTE,pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=ByteCount](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ProcessBuffer", sizeInfo[0])
		}
		o.ProcessBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.ProcessBuffer {
			i1 := i1
			if err := w.ReadData(&o.ProcessBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OldRPCWinStationEnumerateProcessesRequest structure represents the OldRpcWinStationEnumerateProcesses operation request
type OldRPCWinStationEnumerateProcessesRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. hServer MUST
	// be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// ByteCount:  The size, in bytes, of the pProcessBuffer parameter. If ByteCount is
	// too small to receive the data, the method returns an error code (STATUS_INFO_LENGTH_MISMATCH,
	// as specified in [MS-ERREF]) in the pResult parameter. Note that during failure no
	// indication is given to the caller specifying the correct size if pProcessBuffer is
	// too small.
	ByteCount uint32 `idl:"name:ByteCount" json:"byte_count"`
}

func (o *OldRPCWinStationEnumerateProcessesRequest) xxx_ToOp(ctx context.Context, op *xxx_OldRPCWinStationEnumerateProcessesOperation) *xxx_OldRPCWinStationEnumerateProcessesOperation {
	if op == nil {
		op = &xxx_OldRPCWinStationEnumerateProcessesOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ByteCount = o.ByteCount
	return op
}

func (o *OldRPCWinStationEnumerateProcessesRequest) xxx_FromOp(ctx context.Context, op *xxx_OldRPCWinStationEnumerateProcessesOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ByteCount = op.ByteCount
}
func (o *OldRPCWinStationEnumerateProcessesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OldRPCWinStationEnumerateProcessesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OldRPCWinStationEnumerateProcessesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OldRPCWinStationEnumerateProcessesResponse structure represents the OldRpcWinStationEnumerateProcesses operation response
type OldRPCWinStationEnumerateProcessesResponse struct {
	// XXX: ByteCount is an implicit input depedency for output parameters
	ByteCount uint32 `idl:"name:ByteCount" json:"byte_count"`

	// pResult:  Failure error code if the call to RpcWinStationEnumerateProcesses failed.
	// If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000), as
	// specified in [MS-ERREF].
	//
	//	+----------------------------------------+---------------------------------------------+
	//	|                                        |                                             |
	//	|                 VALUE                  |                   MEANING                   |
	//	|                                        |                                             |
	//	+----------------------------------------+---------------------------------------------+
	//	+----------------------------------------+---------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000              | The call succeeded.                         |
	//	+----------------------------------------+---------------------------------------------+
	//	| STATUS_INFO_LENGTH_MISMATCH 0xC0000004 | ByteCount is too small to receive the data. |
	//	+----------------------------------------+---------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// pProcessBuffer:  Pointer to a buffer receiving the list of processes.
	//
	// *
	//
	// A TS_PROCESS_INFORMATION_NT4 ( 0a8a73d9-3c9e-4d2d-91e4-7fef48b6c29d ) structure.
	//
	// *
	//
	// A number of SYSTEM_THREAD_INFORMATION (see [WININTERNALS] or [MSFT-WINSYSINTERNALS]
	// ( https://go.microsoft.com/fwlink/?LinkId=208768 ) ) structures equal to the maximum
	// number of threads in the process.
	//
	// *
	//
	// A TS_SYS_PROCESS_INFORMATION ( 96702c7d-6f5f-4965-8f29-5fdbea73278e ) structure for
	// the process.
	ProcessBuffer []byte `idl:"name:pProcessBuffer;size_is:(ByteCount)" json:"process_buffer"`
	// Return: The OldRpcWinStationEnumerateProcesses return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *OldRPCWinStationEnumerateProcessesResponse) xxx_ToOp(ctx context.Context, op *xxx_OldRPCWinStationEnumerateProcessesOperation) *xxx_OldRPCWinStationEnumerateProcessesOperation {
	if op == nil {
		op = &xxx_OldRPCWinStationEnumerateProcessesOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ByteCount == uint32(0) {
		op.ByteCount = o.ByteCount
	}

	op.Result = o.Result
	op.ProcessBuffer = o.ProcessBuffer
	op.Return = o.Return
	return op
}

func (o *OldRPCWinStationEnumerateProcessesResponse) xxx_FromOp(ctx context.Context, op *xxx_OldRPCWinStationEnumerateProcessesOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ByteCount = op.ByteCount

	o.Result = op.Result
	o.ProcessBuffer = op.ProcessBuffer
	o.Return = op.Return
}
func (o *OldRPCWinStationEnumerateProcessesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OldRPCWinStationEnumerateProcessesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OldRPCWinStationEnumerateProcessesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationEnumerateProcessesOperation structure represents the RpcWinStationEnumerateProcesses operation
type xxx_WinStationEnumerateProcessesOperation struct {
	Server        *tsts.Server `idl:"name:hServer" json:"server"`
	Result        uint32       `idl:"name:pResult" json:"result"`
	ProcessBuffer []byte       `idl:"name:pProcessBuffer;size_is:(ByteCount)" json:"process_buffer"`
	ByteCount     uint32       `idl:"name:ByteCount" json:"byte_count"`
	Return        bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationEnumerateProcessesOperation) OpNum() int { return 36 }

func (o *xxx_WinStationEnumerateProcessesOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationEnumerateProcesses"
}

func (o *xxx_WinStationEnumerateProcessesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ByteCount > uint32(32768) {
		return fmt.Errorf("ByteCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationEnumerateProcessesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ByteCount {in} (1:{range=(0,32768), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ByteCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationEnumerateProcessesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ByteCount {in} (1:{range=(0,32768), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ByteCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationEnumerateProcessesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationEnumerateProcessesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pProcessBuffer {out} (1:{alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=ByteCount](uchar))
	{
		dimSize1 := uint64(o.ByteCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ProcessBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ProcessBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ProcessBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationEnumerateProcessesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pProcessBuffer {out} (1:{alias=PBYTE,pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=ByteCount](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ProcessBuffer", sizeInfo[0])
		}
		o.ProcessBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.ProcessBuffer {
			i1 := i1
			if err := w.ReadData(&o.ProcessBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationEnumerateProcessesRequest structure represents the RpcWinStationEnumerateProcesses operation request
type WinStationEnumerateProcessesRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. hServer MUST
	// be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// ByteCount:  The size, in bytes, of the pProcessBuffer parameter. If ByteCount is
	// too small to receive the data, the method returns an error code (STATUS_INFO_LENGTH_MISMATCH,
	// as specified in [MS-ERREF]) in the pResult parameter. Note that during failure no
	// indication is given to the caller specifying the correct size if pProcessBuffer is
	// too small.
	ByteCount uint32 `idl:"name:ByteCount" json:"byte_count"`
}

func (o *WinStationEnumerateProcessesRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationEnumerateProcessesOperation) *xxx_WinStationEnumerateProcessesOperation {
	if op == nil {
		op = &xxx_WinStationEnumerateProcessesOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ByteCount = o.ByteCount
	return op
}

func (o *WinStationEnumerateProcessesRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationEnumerateProcessesOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ByteCount = op.ByteCount
}
func (o *WinStationEnumerateProcessesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationEnumerateProcessesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationEnumerateProcessesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationEnumerateProcessesResponse structure represents the RpcWinStationEnumerateProcesses operation response
type WinStationEnumerateProcessesResponse struct {
	// XXX: ByteCount is an implicit input depedency for output parameters
	ByteCount uint32 `idl:"name:ByteCount" json:"byte_count"`

	// pResult:  Failure error code if the call to RpcWinStationEnumerateProcesses failed.
	// If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000), as
	// specified in [MS-ERREF].
	//
	//	+----------------------------------------+---------------------------------------------+
	//	|                                        |                                             |
	//	|                 VALUE                  |                   MEANING                   |
	//	|                                        |                                             |
	//	+----------------------------------------+---------------------------------------------+
	//	+----------------------------------------+---------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000              | The call succeeded.                         |
	//	+----------------------------------------+---------------------------------------------+
	//	| STATUS_INFO_LENGTH_MISMATCH 0xC0000004 | ByteCount is too small to receive the data. |
	//	+----------------------------------------+---------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// pProcessBuffer:  Pointer to a buffer receiving the list of processes.
	//
	// *
	//
	// A TS_PROCESS_INFORMATION_NT4 ( 0a8a73d9-3c9e-4d2d-91e4-7fef48b6c29d ) structure.
	//
	// *
	//
	// A number of SYSTEM_THREAD_INFORMATION (see [WININTERNALS] or [MSFT-WINSYSINTERNALS]
	// ( https://go.microsoft.com/fwlink/?LinkId=208768 ) ) structures equal to the maximum
	// number of threads in the process.
	//
	// *
	//
	// A TS_SYS_PROCESS_INFORMATION ( 96702c7d-6f5f-4965-8f29-5fdbea73278e ) structure for
	// the process.
	ProcessBuffer []byte `idl:"name:pProcessBuffer;size_is:(ByteCount)" json:"process_buffer"`
	// Return: The RpcWinStationEnumerateProcesses return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationEnumerateProcessesResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationEnumerateProcessesOperation) *xxx_WinStationEnumerateProcessesOperation {
	if op == nil {
		op = &xxx_WinStationEnumerateProcessesOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ByteCount == uint32(0) {
		op.ByteCount = o.ByteCount
	}

	op.Result = o.Result
	op.ProcessBuffer = o.ProcessBuffer
	op.Return = o.Return
	return op
}

func (o *WinStationEnumerateProcessesResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationEnumerateProcessesOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ByteCount = op.ByteCount

	o.Result = op.Result
	o.ProcessBuffer = op.ProcessBuffer
	o.Return = op.Return
}
func (o *WinStationEnumerateProcessesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationEnumerateProcessesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationEnumerateProcessesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationTerminateProcessOperation structure represents the RpcWinStationTerminateProcess operation
type xxx_WinStationTerminateProcessOperation struct {
	Server    *tsts.Server `idl:"name:hServer" json:"server"`
	Result    uint32       `idl:"name:pResult" json:"result"`
	ProcessID uint32       `idl:"name:ProcessId" json:"process_id"`
	ExitCode  uint32       `idl:"name:ExitCode" json:"exit_code"`
	Return    bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationTerminateProcessOperation) OpNum() int { return 37 }

func (o *xxx_WinStationTerminateProcessOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationTerminateProcess"
}

func (o *xxx_WinStationTerminateProcessOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationTerminateProcessOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ProcessID); err != nil {
			return err
		}
	}
	// ExitCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ExitCode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationTerminateProcessOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ProcessID); err != nil {
			return err
		}
	}
	// ExitCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ExitCode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationTerminateProcessOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationTerminateProcessOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationTerminateProcessOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationTerminateProcessRequest structure represents the RpcWinStationTerminateProcess operation request
type WinStationTerminateProcessRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. hServer MUST
	// be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// ProcessId: The ID of the process to terminate.
	ProcessID uint32 `idl:"name:ProcessId" json:"process_id"`
	// ExitCode: The exit code to be used by the process and threads that are terminated
	// as a result of this call.
	ExitCode uint32 `idl:"name:ExitCode" json:"exit_code"`
}

func (o *WinStationTerminateProcessRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationTerminateProcessOperation) *xxx_WinStationTerminateProcessOperation {
	if op == nil {
		op = &xxx_WinStationTerminateProcessOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ProcessID = o.ProcessID
	op.ExitCode = o.ExitCode
	return op
}

func (o *WinStationTerminateProcessRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationTerminateProcessOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ProcessID = op.ProcessID
	o.ExitCode = op.ExitCode
}
func (o *WinStationTerminateProcessRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationTerminateProcessRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationTerminateProcessOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationTerminateProcessResponse structure represents the RpcWinStationTerminateProcess operation response
type WinStationTerminateProcessResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationTerminateProcess failed.
	// If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000), as
	// specified in [MS-ERREF].
	//
	//	+---------------------------------+---------------------------------------------------------------+
	//	|                                 |                                                               |
	//	|              VALUE              |                            MEANING                            |
	//	|                                 |                                                               |
	//	+---------------------------------+---------------------------------------------------------------+
	//	+---------------------------------+---------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000       | The call succeeded.                                           |
	//	+---------------------------------+---------------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022 | The caller does not have permission to terminate the process. |
	//	+---------------------------------+---------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationTerminateProcess return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationTerminateProcessResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationTerminateProcessOperation) *xxx_WinStationTerminateProcessOperation {
	if op == nil {
		op = &xxx_WinStationTerminateProcessOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationTerminateProcessResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationTerminateProcessOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationTerminateProcessResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationTerminateProcessResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationTerminateProcessOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationGetAllProcessesOperation structure represents the RpcWinStationGetAllProcesses operation
type xxx_WinStationGetAllProcessesOperation struct {
	Server            *tsts.Server             `idl:"name:hServer" json:"server"`
	Result            uint32                   `idl:"name:pResult" json:"result"`
	Level             uint32                   `idl:"name:Level" json:"level"`
	NumberOfProcesses uint32                   `idl:"name:pNumberOfProcesses" json:"number_of_processes"`
	AllProcessesInfo  []*tsts.AllProcessesInfo `idl:"name:ppTsAllProcessesInfo;size_is:(, pNumberOfProcesses)" json:"all_processes_info"`
	Return            bool                     `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationGetAllProcessesOperation) OpNum() int { return 43 }

func (o *xxx_WinStationGetAllProcessesOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationGetAllProcesses"
}

func (o *xxx_WinStationGetAllProcessesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetAllProcessesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Level {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pNumberOfProcesses {in, out} (1:{pointer=ref}*(1))(2:{range=(0,32768), alias=BOUNDED_ULONG, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.NumberOfProcesses); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetAllProcessesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Level {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pNumberOfProcesses {in, out} (1:{pointer=ref}*(1))(2:{range=(0,32768), alias=BOUNDED_ULONG, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.NumberOfProcesses); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetAllProcessesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AllProcessesInfo != nil && o.NumberOfProcesses == 0 {
		o.NumberOfProcesses = uint32(len(o.AllProcessesInfo))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetAllProcessesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pNumberOfProcesses {in, out} (1:{pointer=ref}*(1))(2:{range=(0,32768), alias=BOUNDED_ULONG, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.NumberOfProcesses); err != nil {
			return err
		}
	}
	// ppTsAllProcessesInfo {out} (1:{pointer=ref}*(2))(2:{alias=PTS_ALL_PROCESSES_INFO}*(1))(3:{alias=TS_ALL_PROCESSES_INFO}[dim:0,size_is=pNumberOfProcesses](struct))
	{
		if o.AllProcessesInfo != nil || o.NumberOfProcesses > 0 {
			_ptr_ppTsAllProcessesInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfProcesses)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.AllProcessesInfo {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.AllProcessesInfo[i1] != nil {
						if err := o.AllProcessesInfo[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&tsts.AllProcessesInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.AllProcessesInfo); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&tsts.AllProcessesInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AllProcessesInfo, _ptr_ppTsAllProcessesInfo); err != nil {
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
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetAllProcessesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pNumberOfProcesses {in, out} (1:{pointer=ref}*(1))(2:{range=(0,32768), alias=BOUNDED_ULONG, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.NumberOfProcesses); err != nil {
			return err
		}
	}
	// ppTsAllProcessesInfo {out} (1:{pointer=ref}*(2))(2:{alias=PTS_ALL_PROCESSES_INFO,pointer=ref}*(1))(3:{alias=TS_ALL_PROCESSES_INFO}[dim:0,size_is=pNumberOfProcesses](struct))
	{
		_ptr_ppTsAllProcessesInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.AllProcessesInfo", sizeInfo[0])
			}
			o.AllProcessesInfo = make([]*tsts.AllProcessesInfo, sizeInfo[0])
			for i1 := range o.AllProcessesInfo {
				i1 := i1
				if o.AllProcessesInfo[i1] == nil {
					o.AllProcessesInfo[i1] = &tsts.AllProcessesInfo{}
				}
				if err := o.AllProcessesInfo[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppTsAllProcessesInfo := func(ptr interface{}) { o.AllProcessesInfo = *ptr.(*[]*tsts.AllProcessesInfo) }
		if err := w.ReadPointer(&o.AllProcessesInfo, _s_ppTsAllProcessesInfo, _ptr_ppTsAllProcessesInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationGetAllProcessesRequest structure represents the RpcWinStationGetAllProcesses operation request
type WinStationGetAllProcessesRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// Level:  MUST be 0. Any other value will lead to FALSE being returned by the method.
	Level uint32 `idl:"name:Level" json:"level"`
	// pNumberOfProcesses:  The number of processes requested by the caller. On return,
	// this indicates the number of processes actually stored in the ppTsAllProcessesInfo
	// parameter.
	NumberOfProcesses uint32 `idl:"name:pNumberOfProcesses" json:"number_of_processes"`
}

func (o *WinStationGetAllProcessesRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetAllProcessesOperation) *xxx_WinStationGetAllProcessesOperation {
	if op == nil {
		op = &xxx_WinStationGetAllProcessesOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Level = o.Level
	op.NumberOfProcesses = o.NumberOfProcesses
	return op
}

func (o *WinStationGetAllProcessesRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetAllProcessesOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Level = op.Level
	o.NumberOfProcesses = op.NumberOfProcesses
}
func (o *WinStationGetAllProcessesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationGetAllProcessesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetAllProcessesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationGetAllProcessesResponse structure represents the RpcWinStationGetAllProcesses operation response
type WinStationGetAllProcessesResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationGetAllProcesses failed.
	// If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+---------------------------------+--------------------------------------+
	//	|                                 |                                      |
	//	|              VALUE              |               MEANING                |
	//	|                                 |                                      |
	//	+---------------------------------+--------------------------------------+
	//	+---------------------------------+--------------------------------------+
	//	| STATUS_SUCCESS 0x00000000       | The call was successful.             |
	//	+---------------------------------+--------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022 | The caller does not have permission. |
	//	+---------------------------------+--------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// pNumberOfProcesses:  The number of processes requested by the caller. On return,
	// this indicates the number of processes actually stored in the ppTsAllProcessesInfo
	// parameter.
	NumberOfProcesses uint32 `idl:"name:pNumberOfProcesses" json:"number_of_processes"`
	// ppTsAllProcessesInfo:  Pointer to an array of processes allocated and returned by
	// the method. *ppTsAllProcessesInfo is allocated by the method to be an array of TS_ALL_PROCESSES_INFO
	// structures. The array returned by the method MUST be freed by the caller.
	AllProcessesInfo []*tsts.AllProcessesInfo `idl:"name:ppTsAllProcessesInfo;size_is:(, pNumberOfProcesses)" json:"all_processes_info"`
	// Return: The RpcWinStationGetAllProcesses return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationGetAllProcessesResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetAllProcessesOperation) *xxx_WinStationGetAllProcessesOperation {
	if op == nil {
		op = &xxx_WinStationGetAllProcessesOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.NumberOfProcesses = o.NumberOfProcesses
	op.AllProcessesInfo = o.AllProcessesInfo
	op.Return = o.Return
	return op
}

func (o *WinStationGetAllProcessesResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetAllProcessesOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.NumberOfProcesses = op.NumberOfProcesses
	o.AllProcessesInfo = op.AllProcessesInfo
	o.Return = op.Return
}
func (o *WinStationGetAllProcessesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationGetAllProcessesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetAllProcessesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationGetProcessSIDOperation structure represents the RpcWinStationGetProcessSid operation
type xxx_WinStationGetProcessSIDOperation struct {
	Server           *tsts.Server       `idl:"name:hServer" json:"server"`
	UniqueProcessID  uint32             `idl:"name:dwUniqueProcessId" json:"unique_process_id"`
	ProcessStartTime *dtyp.LargeInteger `idl:"name:ProcessStartTime" json:"process_start_time"`
	Result           int32              `idl:"name:pResult" json:"result"`
	ProcessUserSID   []byte             `idl:"name:pProcessUserSid;size_is:(dwSidSize);pointer:unique" json:"process_user_sid"`
	SIDSize          uint32             `idl:"name:dwSidSize" json:"sid_size"`
	SizeNeeded       uint32             `idl:"name:pdwSizeNeeded" json:"size_needed"`
	Return           bool               `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationGetProcessSIDOperation) OpNum() int { return 44 }

func (o *xxx_WinStationGetProcessSIDOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationGetProcessSid"
}

func (o *xxx_WinStationGetProcessSIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ProcessUserSID != nil && o.SIDSize == 0 {
		o.SIDSize = uint32(len(o.ProcessUserSID))
	}
	if o.SIDSize > uint32(1024) {
		return fmt.Errorf("SIDSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetProcessSIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwUniqueProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UniqueProcessID); err != nil {
			return err
		}
	}
	// ProcessStartTime {in} (1:{alias=LARGE_INTEGER}(struct))
	{
		if o.ProcessStartTime != nil {
			if err := o.ProcessStartTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pProcessUserSid {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwSidSize](uchar))
	{
		if o.ProcessUserSID != nil || o.SIDSize > 0 {
			_ptr_pProcessUserSid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.SIDSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ProcessUserSID {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ProcessUserSID[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ProcessUserSID); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProcessUserSID, _ptr_pProcessUserSid); err != nil {
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
	// dwSidSize {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SIDSize); err != nil {
			return err
		}
	}
	// pdwSizeNeeded {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeNeeded); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetProcessSIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwUniqueProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UniqueProcessID); err != nil {
			return err
		}
	}
	// ProcessStartTime {in} (1:{alias=LARGE_INTEGER}(struct))
	{
		if o.ProcessStartTime == nil {
			o.ProcessStartTime = &dtyp.LargeInteger{}
		}
		if err := o.ProcessStartTime.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pProcessUserSid {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwSidSize](uchar))
	{
		_ptr_pProcessUserSid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ProcessUserSID", sizeInfo[0])
			}
			o.ProcessUserSID = make([]byte, sizeInfo[0])
			for i1 := range o.ProcessUserSID {
				i1 := i1
				if err := w.ReadData(&o.ProcessUserSID[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pProcessUserSid := func(ptr interface{}) { o.ProcessUserSID = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ProcessUserSID, _s_pProcessUserSid, _ptr_pProcessUserSid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwSidSize {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SIDSize); err != nil {
			return err
		}
	}
	// pdwSizeNeeded {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeNeeded); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetProcessSIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetProcessSIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pProcessUserSid {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwSidSize](uchar))
	{
		if o.ProcessUserSID != nil || o.SIDSize > 0 {
			_ptr_pProcessUserSid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.SIDSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ProcessUserSID {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ProcessUserSID[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ProcessUserSID); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProcessUserSID, _ptr_pProcessUserSid); err != nil {
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
	// pdwSizeNeeded {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeNeeded); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetProcessSIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pProcessUserSid {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwSidSize](uchar))
	{
		_ptr_pProcessUserSid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ProcessUserSID", sizeInfo[0])
			}
			o.ProcessUserSID = make([]byte, sizeInfo[0])
			for i1 := range o.ProcessUserSID {
				i1 := i1
				if err := w.ReadData(&o.ProcessUserSID[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pProcessUserSid := func(ptr interface{}) { o.ProcessUserSID = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ProcessUserSID, _s_pProcessUserSid, _ptr_pProcessUserSid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwSizeNeeded {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeNeeded); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationGetProcessSIDRequest structure represents the RpcWinStationGetProcessSid operation request
type WinStationGetProcessSIDRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// dwUniqueProcessId:  The process ID to retrieve the SID.
	UniqueProcessID uint32 `idl:"name:dwUniqueProcessId" json:"unique_process_id"`
	// ProcessStartTime:  The start time of the process indicated by dwUniqueProcessId.
	// This is a 64-bit value representing the number of 100-nanosecond intervals since
	// January 1, 1601 (UTC). ProcessStartTime combined with dwUniqueProcessId is used to
	// identify a process.
	ProcessStartTime *dtyp.LargeInteger `idl:"name:ProcessStartTime" json:"process_start_time"`
	// pProcessUserSid:  The buffer into which the method MUST copy the SID of the process.
	// MUST be NULL if dwSidSize is zero, in which case the correct size will be returned
	// in pdwSizeNeeded.
	ProcessUserSID []byte `idl:"name:pProcessUserSid;size_is:(dwSidSize);pointer:unique" json:"process_user_sid"`
	// dwSidSize:  The size of the buffer, in bytes, pointed to by pProcessUserSid. If
	// the buffer is too small, STATUS_BUFFER_TOO_SMALL<198> or 0x8007007A<199> is returned
	// in pResult.
	SIDSize uint32 `idl:"name:dwSidSize" json:"sid_size"`
	// pdwSizeNeeded:  Indicates the length of the SID. If STATUS_BUFFER_TOO_SMALL<200>
	// or 0x8007007A<201> is returned in pResult, this indicates to the caller the correct
	// size to allocate to a buffer prior to calling the method again.
	SizeNeeded uint32 `idl:"name:pdwSizeNeeded" json:"size_needed"`
}

func (o *WinStationGetProcessSIDRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetProcessSIDOperation) *xxx_WinStationGetProcessSIDOperation {
	if op == nil {
		op = &xxx_WinStationGetProcessSIDOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.UniqueProcessID = o.UniqueProcessID
	op.ProcessStartTime = o.ProcessStartTime
	op.ProcessUserSID = o.ProcessUserSID
	op.SIDSize = o.SIDSize
	op.SizeNeeded = o.SizeNeeded
	return op
}

func (o *WinStationGetProcessSIDRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetProcessSIDOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.UniqueProcessID = op.UniqueProcessID
	o.ProcessStartTime = op.ProcessStartTime
	o.ProcessUserSID = op.ProcessUserSID
	o.SIDSize = op.SIDSize
	o.SizeNeeded = op.SizeNeeded
}
func (o *WinStationGetProcessSIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationGetProcessSIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetProcessSIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationGetProcessSIDResponse structure represents the RpcWinStationGetProcessSid operation response
type WinStationGetProcessSIDResponse struct {
	// XXX: dwSidSize is an implicit input depedency for output parameters
	SIDSize uint32 `idl:"name:dwSidSize" json:"sid_size"`

	// pResult:  Failure error code if the call to RpcWinStationGetProcessSid failed. If
	// the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+------------------------------------+-------------------------------------------------------+
	//	|                                    |                                                       |
	//	|               VALUE                |                        MEANING                        |
	//	|                                    |                                                       |
	//	+------------------------------------+-------------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000          | The call succeeded.                                   |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x8007007A                         | The size of pProcessUserSid buffer is too small.<194> |
	//	+------------------------------------+-------------------------------------------------------+
	//	| STATUS_BUFFER_TOO_SMALL 0xC0000023 | The size of pProcessUserSid buffer is too small.<195> |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x80070005                         | The caller does not have necessary permissions.<196>  |
	//	+------------------------------------+-------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022    | The caller does not have necessary permissions.<197>  |
	//	+------------------------------------+-------------------------------------------------------+
	Result int32 `idl:"name:pResult" json:"result"`
	// pProcessUserSid:  The buffer into which the method MUST copy the SID of the process.
	// MUST be NULL if dwSidSize is zero, in which case the correct size will be returned
	// in pdwSizeNeeded.
	ProcessUserSID []byte `idl:"name:pProcessUserSid;size_is:(dwSidSize);pointer:unique" json:"process_user_sid"`
	// pdwSizeNeeded:  Indicates the length of the SID. If STATUS_BUFFER_TOO_SMALL<200>
	// or 0x8007007A<201> is returned in pResult, this indicates to the caller the correct
	// size to allocate to a buffer prior to calling the method again.
	SizeNeeded uint32 `idl:"name:pdwSizeNeeded" json:"size_needed"`
	// Return: The RpcWinStationGetProcessSid return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationGetProcessSIDResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetProcessSIDOperation) *xxx_WinStationGetProcessSIDOperation {
	if op == nil {
		op = &xxx_WinStationGetProcessSIDOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.SIDSize == uint32(0) {
		op.SIDSize = o.SIDSize
	}

	op.Result = o.Result
	op.ProcessUserSID = o.ProcessUserSID
	op.SizeNeeded = o.SizeNeeded
	op.Return = o.Return
	return op
}

func (o *WinStationGetProcessSIDResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetProcessSIDOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.SIDSize = op.SIDSize

	o.Result = op.Result
	o.ProcessUserSID = op.ProcessUserSID
	o.SizeNeeded = op.SizeNeeded
	o.Return = op.Return
}
func (o *WinStationGetProcessSIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationGetProcessSIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetProcessSIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationGetTerminateServerCountersValueOperation structure represents the RpcWinStationGetTermSrvCountersValue operation
type xxx_WinStationGetTerminateServerCountersValueOperation struct {
	Server  *tsts.Server    `idl:"name:hServer" json:"server"`
	Result  uint32          `idl:"name:pResult" json:"result"`
	Entries uint32          `idl:"name:dwEntries" json:"entries"`
	Counter []*tsts.Counter `idl:"name:pCounter;size_is:(dwEntries)" json:"counter"`
	Return  bool            `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationGetTerminateServerCountersValueOperation) OpNum() int { return 45 }

func (o *xxx_WinStationGetTerminateServerCountersValueOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationGetTermSrvCountersValue"
}

func (o *xxx_WinStationGetTerminateServerCountersValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Counter != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.Counter))
	}
	if o.Entries > uint32(4096) {
		return fmt.Errorf("Entries is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetTerminateServerCountersValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwEntries {in} (1:{range=(0,4096), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Entries); err != nil {
			return err
		}
	}
	// pCounter {in, out} (1:{alias=PTS_COUNTER}*(1))(2:{alias=TS_COUNTER}[dim:0,size_is=dwEntries](struct))
	{
		dimSize1 := uint64(o.Entries)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Counter {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Counter[i1] != nil {
				if err := o.Counter[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&tsts.Counter{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Counter); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&tsts.Counter{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_WinStationGetTerminateServerCountersValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwEntries {in} (1:{range=(0,4096), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Entries); err != nil {
			return err
		}
	}
	// pCounter {in, out} (1:{alias=PTS_COUNTER,pointer=ref}*(1))(2:{alias=TS_COUNTER}[dim:0,size_is=dwEntries](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Counter", sizeInfo[0])
		}
		o.Counter = make([]*tsts.Counter, sizeInfo[0])
		for i1 := range o.Counter {
			i1 := i1
			if o.Counter[i1] == nil {
				o.Counter[i1] = &tsts.Counter{}
			}
			if err := o.Counter[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_WinStationGetTerminateServerCountersValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetTerminateServerCountersValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pCounter {in, out} (1:{alias=PTS_COUNTER}*(1))(2:{alias=TS_COUNTER}[dim:0,size_is=dwEntries](struct))
	{
		dimSize1 := uint64(o.Entries)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Counter {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Counter[i1] != nil {
				if err := o.Counter[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&tsts.Counter{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Counter); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&tsts.Counter{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetTerminateServerCountersValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pCounter {in, out} (1:{alias=PTS_COUNTER,pointer=ref}*(1))(2:{alias=TS_COUNTER}[dim:0,size_is=dwEntries](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Counter", sizeInfo[0])
		}
		o.Counter = make([]*tsts.Counter, sizeInfo[0])
		for i1 := range o.Counter {
			i1 := i1
			if o.Counter[i1] == nil {
				o.Counter[i1] = &tsts.Counter{}
			}
			if err := o.Counter[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationGetTerminateServerCountersValueRequest structure represents the RpcWinStationGetTermSrvCountersValue operation request
type WinStationGetTerminateServerCountersValueRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// dwEntries:  The number of performance counters to query. Indicates the size of the
	// array pointed to by pCounter.
	Entries uint32 `idl:"name:dwEntries" json:"entries"`
	// pCounter:  An array of TS_COUNTER structures. The caller MUST set the dwCounterId
	// in the TS_COUNTER structures for each entry in the array to indicate the counter
	// whose current value to retrieve. On return, the method MUST set the value for that
	// performance counter. If the performance counter ID is not recognized or is not supported,
	// it will set the bResult to 0.
	Counter []*tsts.Counter `idl:"name:pCounter;size_is:(dwEntries)" json:"counter"`
}

func (o *WinStationGetTerminateServerCountersValueRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetTerminateServerCountersValueOperation) *xxx_WinStationGetTerminateServerCountersValueOperation {
	if op == nil {
		op = &xxx_WinStationGetTerminateServerCountersValueOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Entries = o.Entries
	op.Counter = o.Counter
	return op
}

func (o *WinStationGetTerminateServerCountersValueRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetTerminateServerCountersValueOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Entries = op.Entries
	o.Counter = op.Counter
}
func (o *WinStationGetTerminateServerCountersValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationGetTerminateServerCountersValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetTerminateServerCountersValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationGetTerminateServerCountersValueResponse structure represents the RpcWinStationGetTermSrvCountersValue operation response
type WinStationGetTerminateServerCountersValueResponse struct {
	// XXX: dwEntries is an implicit input depedency for output parameters
	Entries uint32 `idl:"name:dwEntries" json:"entries"`

	// pResult:  If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000);
	// otherwise it MUST be an implementation-specific negative value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// pCounter:  An array of TS_COUNTER structures. The caller MUST set the dwCounterId
	// in the TS_COUNTER structures for each entry in the array to indicate the counter
	// whose current value to retrieve. On return, the method MUST set the value for that
	// performance counter. If the performance counter ID is not recognized or is not supported,
	// it will set the bResult to 0.
	Counter []*tsts.Counter `idl:"name:pCounter;size_is:(dwEntries)" json:"counter"`
	// Return: The RpcWinStationGetTermSrvCountersValue return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationGetTerminateServerCountersValueResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetTerminateServerCountersValueOperation) *xxx_WinStationGetTerminateServerCountersValueOperation {
	if op == nil {
		op = &xxx_WinStationGetTerminateServerCountersValueOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Entries == uint32(0) {
		op.Entries = o.Entries
	}

	op.Result = o.Result
	op.Counter = o.Counter
	op.Return = o.Return
	return op
}

func (o *WinStationGetTerminateServerCountersValueResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetTerminateServerCountersValueOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Entries = op.Entries

	o.Result = op.Result
	o.Counter = op.Counter
	o.Return = op.Return
}
func (o *WinStationGetTerminateServerCountersValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationGetTerminateServerCountersValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetTerminateServerCountersValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationReInitializeSecurityOperation structure represents the RpcWinStationReInitializeSecurity operation
type xxx_WinStationReInitializeSecurityOperation struct {
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	Result uint32       `idl:"name:pResult" json:"result"`
	Return bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationReInitializeSecurityOperation) OpNum() int { return 46 }

func (o *xxx_WinStationReInitializeSecurityOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationReInitializeSecurity"
}

func (o *xxx_WinStationReInitializeSecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationReInitializeSecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_WinStationReInitializeSecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationReInitializeSecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationReInitializeSecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationReInitializeSecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationReInitializeSecurityRequest structure represents the RpcWinStationReInitializeSecurity operation request
type WinStationReInitializeSecurityRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
}

func (o *WinStationReInitializeSecurityRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationReInitializeSecurityOperation) *xxx_WinStationReInitializeSecurityOperation {
	if op == nil {
		op = &xxx_WinStationReInitializeSecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	return op
}

func (o *WinStationReInitializeSecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationReInitializeSecurityOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
}
func (o *WinStationReInitializeSecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationReInitializeSecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationReInitializeSecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationReInitializeSecurityResponse structure represents the RpcWinStationReInitializeSecurity operation response
type WinStationReInitializeSecurityResponse struct {
	// pResult:  If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000);
	// otherwise, it MUST be an implementation-specific negative value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationReInitializeSecurity return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationReInitializeSecurityResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationReInitializeSecurityOperation) *xxx_WinStationReInitializeSecurityOperation {
	if op == nil {
		op = &xxx_WinStationReInitializeSecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationReInitializeSecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationReInitializeSecurityOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationReInitializeSecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationReInitializeSecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationReInitializeSecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationGetLANAdapterNameOperation structure represents the RpcWinStationGetLanAdapterName operation
type xxx_WinStationGetLANAdapterNameOperation struct {
	Server         *tsts.Server `idl:"name:hServer" json:"server"`
	Result         uint32       `idl:"name:pResult" json:"result"`
	PDNameSize     uint32       `idl:"name:PdNameSize" json:"pd_name_size"`
	PDName         string       `idl:"name:pPdName;size_is:(PdNameSize)" json:"pd_name"`
	LANAdapter     uint32       `idl:"name:LanAdapter" json:"lan_adapter"`
	Length         uint32       `idl:"name:pLength" json:"length"`
	LanAdapterName string       `idl:"name:ppLanAdapter;size_is:(, pLength)" json:"lan_adapter_name"`
	Return         bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationGetLANAdapterNameOperation) OpNum() int { return 53 }

func (o *xxx_WinStationGetLANAdapterNameOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationGetLanAdapterName"
}

func (o *xxx_WinStationGetLANAdapterNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.PDName != "" && o.PDNameSize == 0 {
		o.PDNameSize = uint32(ndr.UTF16Len(o.PDName))
	}
	if o.PDNameSize > uint32(4096) {
		return fmt.Errorf("PDNameSize is out of range")
	}
	if o.LANAdapter > uint32(1024) {
		return fmt.Errorf("LANAdapter is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetLANAdapterNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// PdNameSize {in} (1:{range=(0,4096), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PDNameSize); err != nil {
			return err
		}
	}
	// pPdName {in} (1:{alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=PdNameSize,string](wchar))
	{
		dimSize1 := uint64(o.PDNameSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_PDName_buf := utf16.Encode([]rune(o.PDName))
		if uint64(len(_PDName_buf)) > sizeInfo[0] {
			_PDName_buf = _PDName_buf[:sizeInfo[0]]
		}
		for i1 := range _PDName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_PDName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_PDName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// LanAdapter {in} (1:{range=(0,1024), alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.LANAdapter); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetLANAdapterNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// PdNameSize {in} (1:{range=(0,4096), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PDNameSize); err != nil {
			return err
		}
	}
	// pPdName {in} (1:{alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=PdNameSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _PDName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _PDName_buf", sizeInfo[0])
		}
		_PDName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _PDName_buf {
			i1 := i1
			if err := w.ReadData(&_PDName_buf[i1]); err != nil {
				return err
			}
		}
		o.PDName = strings.TrimRight(string(utf16.Decode(_PDName_buf)), ndr.ZeroString)
	}
	// LanAdapter {in} (1:{range=(0,1024), alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.LANAdapter); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetLANAdapterNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.LanAdapterName != "" && o.Length == 0 {
		o.Length = uint32(ndr.UTF16Len(o.LanAdapterName))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetLANAdapterNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pLength {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	// ppLanAdapter {out} (1:{pointer=ref}*(2))(2:{alias=PWCHAR}*(1))(3:{alias=WCHAR}[dim:0,size_is=pLength,string](wchar))
	{
		if o.LanAdapterName != "" || o.Length > 0 {
			_ptr_ppLanAdapter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Length)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				_LanAdapterName_buf := utf16.Encode([]rune(o.LanAdapterName))
				if uint64(len(_LanAdapterName_buf)) > sizeInfo[0] {
					_LanAdapterName_buf = _LanAdapterName_buf[:sizeInfo[0]]
				}
				for i1 := range _LanAdapterName_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_LanAdapterName_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_LanAdapterName_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LanAdapterName, _ptr_ppLanAdapter); err != nil {
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
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetLANAdapterNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pLength {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	// ppLanAdapter {out} (1:{pointer=ref}*(2))(2:{alias=PWCHAR,pointer=ref}*(1))(3:{alias=WCHAR}[dim:0,size_is=pLength,string](wchar))
	{
		_ptr_ppLanAdapter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _LanAdapterName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _LanAdapterName_buf", sizeInfo[0])
			}
			_LanAdapterName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _LanAdapterName_buf {
				i1 := i1
				if err := w.ReadData(&_LanAdapterName_buf[i1]); err != nil {
					return err
				}
			}
			o.LanAdapterName = strings.TrimRight(string(utf16.Decode(_LanAdapterName_buf)), ndr.ZeroString)
			return nil
		})
		_s_ppLanAdapter := func(ptr interface{}) { o.LanAdapterName = *ptr.(*string) }
		if err := w.ReadPointer(&o.LanAdapterName, _s_ppLanAdapter, _ptr_ppLanAdapter); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationGetLANAdapterNameRequest structure represents the RpcWinStationGetLanAdapterName operation request
type WinStationGetLANAdapterNameRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// PdNameSize: The size, in bytes, of pPdName including the terminating NULL character.
	PDNameSize uint32 `idl:"name:PdNameSize" json:"pd_name_size"`
	// pPdName: The transport protocol type on which to retrieve information. This MUST
	// be any of the following strings: tcp, netbios, ipx, spx.
	PDName string `idl:"name:pPdName;size_is:(PdNameSize)" json:"pd_name"`
	// LanAdapter: The number of the LAN adapter to retrieve information (also known as
	// lana). If this is set to "0", it will always return a LAN adapter name to indicate
	// all LAN adapters configured with the protocol, irrespective of the transport protocol
	// type specified in pPdName.
	LANAdapter uint32 `idl:"name:LanAdapter" json:"lan_adapter"`
}

func (o *WinStationGetLANAdapterNameRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetLANAdapterNameOperation) *xxx_WinStationGetLANAdapterNameOperation {
	if op == nil {
		op = &xxx_WinStationGetLANAdapterNameOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.PDNameSize = o.PDNameSize
	op.PDName = o.PDName
	op.LANAdapter = o.LANAdapter
	return op
}

func (o *WinStationGetLANAdapterNameRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetLANAdapterNameOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.PDNameSize = op.PDNameSize
	o.PDName = op.PDName
	o.LANAdapter = op.LANAdapter
}
func (o *WinStationGetLANAdapterNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationGetLANAdapterNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetLANAdapterNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationGetLANAdapterNameResponse structure represents the RpcWinStationGetLanAdapterName operation response
type WinStationGetLANAdapterNameResponse struct {
	// pResult: If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000);
	// otherwise, it MUST be an implementation-specific negative value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// pLength: The pointer to a ULONG containing the length of the string ppLanAdapter,
	// in characters including the terminating NULL character. If LanAdapter is 0, this
	// value MUST be DEVICENAME_LENGTH+1.
	Length uint32 `idl:"name:pLength" json:"length"`
	// ppLanAdapter:  The pointer to a string allocated by this method for retrieving the
	// LAN adapter's name. This memory MUST be freed by the caller.
	LanAdapterName string `idl:"name:ppLanAdapter;size_is:(, pLength)" json:"lan_adapter_name"`
	// Return: The RpcWinStationGetLanAdapterName return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationGetLANAdapterNameResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetLANAdapterNameOperation) *xxx_WinStationGetLANAdapterNameOperation {
	if op == nil {
		op = &xxx_WinStationGetLANAdapterNameOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Length = o.Length
	op.LanAdapterName = o.LanAdapterName
	op.Return = o.Return
	return op
}

func (o *WinStationGetLANAdapterNameResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetLANAdapterNameOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Length = op.Length
	o.LanAdapterName = op.LanAdapterName
	o.Return = op.Return
}
func (o *WinStationGetLANAdapterNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationGetLANAdapterNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetLANAdapterNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationUpdateSettingsOperation structure represents the RpcWinStationUpdateSettings operation
type xxx_WinStationUpdateSettingsOperation struct {
	Server             *tsts.Server `idl:"name:hServer" json:"server"`
	Result             uint32       `idl:"name:pResult" json:"result"`
	SettingsClass      uint32       `idl:"name:SettingsClass" json:"settings_class"`
	SettingsParameters uint32       `idl:"name:SettingsParameters" json:"settings_parameters"`
	Return             bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationUpdateSettingsOperation) OpNum() int { return 58 }

func (o *xxx_WinStationUpdateSettingsOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationUpdateSettings"
}

func (o *xxx_WinStationUpdateSettingsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationUpdateSettingsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SettingsClass {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingsClass); err != nil {
			return err
		}
	}
	// SettingsParameters {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingsParameters); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationUpdateSettingsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SettingsClass {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingsClass); err != nil {
			return err
		}
	}
	// SettingsParameters {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingsParameters); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationUpdateSettingsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationUpdateSettingsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationUpdateSettingsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationUpdateSettingsRequest structure represents the RpcWinStationUpdateSettings operation request
type WinStationUpdateSettingsRequest struct {
	// hServer:  A handle to the server object of type SERVER_HANDLE. The hServer argument
	// MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// SettingsClass:  The class for which to update settings.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| WINSTACFG_SESSDIR 0x00000001 | Contacts Session Directory to reread the WinStation settings.<207>               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| WINSTACFG_LEGACY 0x00000000  | Rereads settings from the local registry for the configured winstations. This    |
	//	|                              | does the same thing as RpcWinStationReadRegistry.                                |
	//	+------------------------------+----------------------------------------------------------------------------------+
	SettingsClass uint32 `idl:"name:SettingsClass" json:"settings_class"`
	// SettingsParameters:  MUST be 0.
	SettingsParameters uint32 `idl:"name:SettingsParameters" json:"settings_parameters"`
}

func (o *WinStationUpdateSettingsRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationUpdateSettingsOperation) *xxx_WinStationUpdateSettingsOperation {
	if op == nil {
		op = &xxx_WinStationUpdateSettingsOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.SettingsClass = o.SettingsClass
	op.SettingsParameters = o.SettingsParameters
	return op
}

func (o *WinStationUpdateSettingsRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationUpdateSettingsOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.SettingsClass = op.SettingsClass
	o.SettingsParameters = op.SettingsParameters
}
func (o *WinStationUpdateSettingsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationUpdateSettingsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationUpdateSettingsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationUpdateSettingsResponse structure represents the RpcWinStationUpdateSettings operation response
type WinStationUpdateSettingsResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationUpdateSettings failed.
	// If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+-------------------------------------+-----------------------------------------------------------+
	//	|                                     |                                                           |
	//	|                VALUE                |                          MEANING                          |
	//	|                                     |                                                           |
	//	+-------------------------------------+-----------------------------------------------------------+
	//	+-------------------------------------+-----------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000           | The call was successful.                                  |
	//	+-------------------------------------+-----------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022     | The caller does not have permission to read the settings. |
	//	+-------------------------------------+-----------------------------------------------------------+
	//	| STATUS_INVALID_PARAMETER 0xC000000D | Unrecognized SettingsClass.                               |
	//	+-------------------------------------+-----------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationUpdateSettings return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationUpdateSettingsResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationUpdateSettingsOperation) *xxx_WinStationUpdateSettingsOperation {
	if op == nil {
		op = &xxx_WinStationUpdateSettingsOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationUpdateSettingsResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationUpdateSettingsOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationUpdateSettingsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationUpdateSettingsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationUpdateSettingsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationShadowStopOperation structure represents the RpcWinStationShadowStop operation
type xxx_WinStationShadowStopOperation struct {
	Server  *tsts.Server `idl:"name:hServer" json:"server"`
	Result  uint32       `idl:"name:pResult" json:"result"`
	LogonID uint32       `idl:"name:LogonId" json:"logon_id"`
	Wait    bool         `idl:"name:bWait" json:"wait"`
	Return  bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationShadowStopOperation) OpNum() int { return 59 }

func (o *xxx_WinStationShadowStopOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationShadowStop"
}

func (o *xxx_WinStationShadowStopOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShadowStopOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LogonID); err != nil {
			return err
		}
	}
	// bWait {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Wait); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShadowStopOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LogonID); err != nil {
			return err
		}
	}
	// bWait {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Wait); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShadowStopOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShadowStopOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationShadowStopOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationShadowStopRequest structure represents the RpcWinStationShadowStop operation request
type WinStationShadowStopRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// LogonId: The ID of the session on which to stop shadowing operations.
	LogonID uint32 `idl:"name:LogonId" json:"logon_id"`
	// bWait: TRUE indicates wait for reset to complete, FALSE otherwise.
	Wait bool `idl:"name:bWait" json:"wait"`
}

func (o *WinStationShadowStopRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationShadowStopOperation) *xxx_WinStationShadowStopOperation {
	if op == nil {
		op = &xxx_WinStationShadowStopOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.LogonID = o.LogonID
	op.Wait = o.Wait
	return op
}

func (o *WinStationShadowStopRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationShadowStopOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.LogonID = op.LogonID
	o.Wait = op.Wait
}
func (o *WinStationShadowStopRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationShadowStopRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationShadowStopOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationShadowStopResponse structure represents the RpcWinStationShadowStop operation response
type WinStationShadowStopResponse struct {
	// pResult: Failure error code if the call to RpcWinStationShadowStop failed. If the
	// call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	|                                            |                                                                                |
	//	|                   VALUE                    |                                    MEANING                                     |
	//	|                                            |                                                                                |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000                  | The call was successful.                                                       |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| STATUS_CTX_WINSTATION_NOT_FOUND 0xC00A0015 | Indicates the session does not exist.                                          |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| STATUS_CTX_SHADOW_NOT_RUNNING 0xC00A0036   | Indicates the session is either not active or not being shadowed.              |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022            | Indicates the caller does not have permission to end shadowing on the session. |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationShadowStop return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationShadowStopResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationShadowStopOperation) *xxx_WinStationShadowStopOperation {
	if op == nil {
		op = &xxx_WinStationShadowStopOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationShadowStopResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationShadowStopOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationShadowStopResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationShadowStopResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationShadowStopOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationCloseServerExOperation structure represents the RpcWinStationCloseServerEx operation
type xxx_WinStationCloseServerExOperation struct {
	Server *tsts.Server `idl:"name:phServer" json:"server"`
	Result uint32       `idl:"name:pResult" json:"result"`
	Return bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationCloseServerExOperation) OpNum() int { return 60 }

func (o *xxx_WinStationCloseServerExOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationCloseServerEx"
}

func (o *xxx_WinStationCloseServerExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCloseServerExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phServer {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_WinStationCloseServerExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phServer {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCloseServerExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCloseServerExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phServer {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCloseServerExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phServer {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationCloseServerExRequest structure represents the RpcWinStationCloseServerEx operation request
type WinStationCloseServerExRequest struct {
	// phServer: Pointer to a variable that is a handle to the server. The variable is of
	// type SERVER_HANDLE. The handle MUST be returned from a previous call to RpcWinStationOpenServer.
	// On return from this method, *phServer is set to NULL.
	Server *tsts.Server `idl:"name:phServer" json:"server"`
}

func (o *WinStationCloseServerExRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationCloseServerExOperation) *xxx_WinStationCloseServerExOperation {
	if op == nil {
		op = &xxx_WinStationCloseServerExOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	return op
}

func (o *WinStationCloseServerExRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationCloseServerExOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
}
func (o *WinStationCloseServerExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationCloseServerExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationCloseServerExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationCloseServerExResponse structure represents the RpcWinStationCloseServerEx operation response
type WinStationCloseServerExResponse struct {
	// phServer: Pointer to a variable that is a handle to the server. The variable is of
	// type SERVER_HANDLE. The handle MUST be returned from a previous call to RpcWinStationOpenServer.
	// On return from this method, *phServer is set to NULL.
	Server *tsts.Server `idl:"name:phServer" json:"server"`
	// pResult: If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000);
	// otherwise, it MUST be an implementation-specific negative value.
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationCloseServerEx return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationCloseServerExResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationCloseServerExOperation) *xxx_WinStationCloseServerExOperation {
	if op == nil {
		op = &xxx_WinStationCloseServerExOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationCloseServerExResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationCloseServerExOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationCloseServerExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationCloseServerExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationCloseServerExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationIsHelpAssistantSessionOperation structure represents the RpcWinStationIsHelpAssistantSession operation
type xxx_WinStationIsHelpAssistantSessionOperation struct {
	Server    *tsts.Server `idl:"name:hServer" json:"server"`
	Result    uint32       `idl:"name:pResult" json:"result"`
	SessionID uint32       `idl:"name:SessionId" json:"session_id"`
	Return    bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationIsHelpAssistantSessionOperation) OpNum() int { return 61 }

func (o *xxx_WinStationIsHelpAssistantSessionOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationIsHelpAssistantSession"
}

func (o *xxx_WinStationIsHelpAssistantSessionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationIsHelpAssistantSessionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationIsHelpAssistantSessionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SessionId {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationIsHelpAssistantSessionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationIsHelpAssistantSessionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationIsHelpAssistantSessionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationIsHelpAssistantSessionRequest structure represents the RpcWinStationIsHelpAssistantSession operation request
type WinStationIsHelpAssistantSessionRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// SessionId: The ID of the session to check.
	SessionID uint32 `idl:"name:SessionId" json:"session_id"`
}

func (o *WinStationIsHelpAssistantSessionRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationIsHelpAssistantSessionOperation) *xxx_WinStationIsHelpAssistantSessionOperation {
	if op == nil {
		op = &xxx_WinStationIsHelpAssistantSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.SessionID = o.SessionID
	return op
}

func (o *WinStationIsHelpAssistantSessionRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationIsHelpAssistantSessionOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.SessionID = op.SessionID
}
func (o *WinStationIsHelpAssistantSessionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationIsHelpAssistantSessionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationIsHelpAssistantSessionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationIsHelpAssistantSessionResponse structure represents the RpcWinStationIsHelpAssistantSession operation response
type WinStationIsHelpAssistantSessionResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationIsHelpAssistantSession
	// failed. If no error was encountered while executing the call, this parameter MUST
	// be STATUS_SUCCESS (0x00000000).
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                            |                                                                                  |
	//	|                   VALUE                    |                                     MEANING                                      |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000                  | The call was executed successfully.                                              |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_UNSUCCESSFUL 0xC0000001             | Execution of the call failed.                                                    |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_CTX_WINSTATION_NOT_FOUND 0xC00A0015 | The SessionId does not exist.                                                    |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_WRONG_PASSWORD 0xC000006A           | This is a Help Assistant session but the help assistance ticket associated with  |
	//	|                                            | the session is no longer valid.                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationIsHelpAssistantSession return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationIsHelpAssistantSessionResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationIsHelpAssistantSessionOperation) *xxx_WinStationIsHelpAssistantSessionOperation {
	if op == nil {
		op = &xxx_WinStationIsHelpAssistantSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationIsHelpAssistantSessionResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationIsHelpAssistantSessionOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationIsHelpAssistantSessionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationIsHelpAssistantSessionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationIsHelpAssistantSessionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationGetMachinePolicyOperation structure represents the RpcWinStationGetMachinePolicy operation
type xxx_WinStationGetMachinePolicyOperation struct {
	Server     *tsts.Server `idl:"name:hServer" json:"server"`
	Policy     []byte       `idl:"name:pPolicy;size_is:(bufferSize)" json:"policy"`
	BufferSize uint32       `idl:"name:bufferSize" json:"buffer_size"`
	Return     bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationGetMachinePolicyOperation) OpNum() int { return 62 }

func (o *xxx_WinStationGetMachinePolicyOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationGetMachinePolicy"
}

func (o *xxx_WinStationGetMachinePolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Policy != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Policy))
	}
	if o.BufferSize > uint32(32768) {
		return fmt.Errorf("BufferSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetMachinePolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pPolicy {in, out} (1:{alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=bufferSize](uchar))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Policy {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Policy[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Policy); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// bufferSize {in} (1:{range=(0,32768), alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetMachinePolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pPolicy {in, out} (1:{alias=PBYTE,pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=bufferSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Policy", sizeInfo[0])
		}
		o.Policy = make([]byte, sizeInfo[0])
		for i1 := range o.Policy {
			i1 := i1
			if err := w.ReadData(&o.Policy[i1]); err != nil {
				return err
			}
		}
	}
	// bufferSize {in} (1:{range=(0,32768), alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetMachinePolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetMachinePolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pPolicy {in, out} (1:{alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=bufferSize](uchar))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Policy {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Policy[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Policy); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetMachinePolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pPolicy {in, out} (1:{alias=PBYTE,pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=bufferSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Policy", sizeInfo[0])
		}
		o.Policy = make([]byte, sizeInfo[0])
		for i1 := range o.Policy {
			i1 := i1
			if err := w.ReadData(&o.Policy[i1]); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationGetMachinePolicyRequest structure represents the RpcWinStationGetMachinePolicy operation request
type WinStationGetMachinePolicyRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// pPolicy:  Pointer to a buffer to receive the machine policy. This buffer MUST be
	// of type POLICY_TS_MACHINE.
	Policy []byte `idl:"name:pPolicy;size_is:(bufferSize)" json:"policy"`
	// bufferSize:  Size of the buffer, in bytes, pointed to by pPolicy. This size MUST
	// NOT be less than sizeof(POLICY_TS_MACHINE).
	BufferSize uint32 `idl:"name:bufferSize" json:"buffer_size"`
}

func (o *WinStationGetMachinePolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetMachinePolicyOperation) *xxx_WinStationGetMachinePolicyOperation {
	if op == nil {
		op = &xxx_WinStationGetMachinePolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Policy = o.Policy
	op.BufferSize = o.BufferSize
	return op
}

func (o *WinStationGetMachinePolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetMachinePolicyOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Policy = op.Policy
	o.BufferSize = op.BufferSize
}
func (o *WinStationGetMachinePolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationGetMachinePolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetMachinePolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationGetMachinePolicyResponse structure represents the RpcWinStationGetMachinePolicy operation response
type WinStationGetMachinePolicyResponse struct {
	// XXX: bufferSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:bufferSize" json:"buffer_size"`

	// pPolicy:  Pointer to a buffer to receive the machine policy. This buffer MUST be
	// of type POLICY_TS_MACHINE.
	Policy []byte `idl:"name:pPolicy;size_is:(bufferSize)" json:"policy"`
	// Return: The RpcWinStationGetMachinePolicy return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationGetMachinePolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetMachinePolicyOperation) *xxx_WinStationGetMachinePolicyOperation {
	if op == nil {
		op = &xxx_WinStationGetMachinePolicyOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.Policy = o.Policy
	op.Return = o.Return
	return op
}

func (o *WinStationGetMachinePolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetMachinePolicyOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.Policy = op.Policy
	o.Return = op.Return
}
func (o *WinStationGetMachinePolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationGetMachinePolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetMachinePolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationCheckLoopBackOperation structure represents the RpcWinStationCheckLoopBack operation
type xxx_WinStationCheckLoopBackOperation struct {
	Server           *tsts.Server `idl:"name:hServer" json:"server"`
	Result           uint32       `idl:"name:pResult" json:"result"`
	ClientLogonID    uint32       `idl:"name:ClientLogonId" json:"client_logon_id"`
	TargetLogonID    uint32       `idl:"name:TargetLogonId" json:"target_logon_id"`
	TargetServerName string       `idl:"name:pTargetServerName;size_is:(NameSize)" json:"target_server_name"`
	NameSize         uint32       `idl:"name:NameSize" json:"name_size"`
	Return           bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationCheckLoopBackOperation) OpNum() int { return 65 }

func (o *xxx_WinStationCheckLoopBackOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationCheckLoopBack"
}

func (o *xxx_WinStationCheckLoopBackOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.TargetServerName != "" && o.NameSize == 0 {
		o.NameSize = uint32(ndr.UTF16Len(o.TargetServerName))
	}
	if o.NameSize > uint32(1024) {
		return fmt.Errorf("NameSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCheckLoopBackOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ClientLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientLogonID); err != nil {
			return err
		}
	}
	// TargetLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TargetLogonID); err != nil {
			return err
		}
	}
	// pTargetServerName {in} (1:{alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameSize,string](wchar))
	{
		dimSize1 := uint64(o.NameSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		_TargetServerName_buf := utf16.Encode([]rune(o.TargetServerName))
		if uint64(len(_TargetServerName_buf)) > sizeInfo[0] {
			_TargetServerName_buf = _TargetServerName_buf[:sizeInfo[0]]
		}
		for i1 := range _TargetServerName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_TargetServerName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_TargetServerName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// NameSize {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCheckLoopBackOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ClientLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientLogonID); err != nil {
			return err
		}
	}
	// TargetLogonId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TargetLogonID); err != nil {
			return err
		}
	}
	// pTargetServerName {in} (1:{alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,size_is=NameSize,string](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _TargetServerName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _TargetServerName_buf", sizeInfo[0])
		}
		_TargetServerName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _TargetServerName_buf {
			i1 := i1
			if err := w.ReadData(&_TargetServerName_buf[i1]); err != nil {
				return err
			}
		}
		o.TargetServerName = strings.TrimRight(string(utf16.Decode(_TargetServerName_buf)), ndr.ZeroString)
	}
	// NameSize {in} (1:{range=(0,1024), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCheckLoopBackOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCheckLoopBackOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationCheckLoopBackOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationCheckLoopBackRequest structure represents the RpcWinStationCheckLoopBack operation request
type WinStationCheckLoopBackRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// ClientLogonId: The ID of the session from which the terminal server client was started.
	ClientLogonID uint32 `idl:"name:ClientLogonId" json:"client_logon_id"`
	// TargetLogonId: The session ID to which the client is trying to connect.
	TargetLogonID uint32 `idl:"name:TargetLogonId" json:"target_logon_id"`
	// pTargetServerName: The name of the target server to which the client is connecting.
	// The string MUST contain the terminating NULL character.
	TargetServerName string `idl:"name:pTargetServerName;size_is:(NameSize)" json:"target_server_name"`
	// NameSize: The length of the pTargetServerName string in characters including the
	// terminating NULL character.
	NameSize uint32 `idl:"name:NameSize" json:"name_size"`
}

func (o *WinStationCheckLoopBackRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationCheckLoopBackOperation) *xxx_WinStationCheckLoopBackOperation {
	if op == nil {
		op = &xxx_WinStationCheckLoopBackOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ClientLogonID = o.ClientLogonID
	op.TargetLogonID = o.TargetLogonID
	op.TargetServerName = o.TargetServerName
	op.NameSize = o.NameSize
	return op
}

func (o *WinStationCheckLoopBackRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationCheckLoopBackOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ClientLogonID = op.ClientLogonID
	o.TargetLogonID = op.TargetLogonID
	o.TargetServerName = op.TargetServerName
	o.NameSize = op.NameSize
}
func (o *WinStationCheckLoopBackRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationCheckLoopBackRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationCheckLoopBackOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationCheckLoopBackResponse structure represents the RpcWinStationCheckLoopBack operation response
type WinStationCheckLoopBackResponse struct {
	// pResult: Failure error code if the call to RpcWinStationCheckLoopBack failed. If
	// the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+------------------------------------------------+--------------------------------------------------------------------------------+
	//	|                                                |                                                                                |
	//	|                     VALUE                      |                                    MEANING                                     |
	//	|                                                |                                                                                |
	//	+------------------------------------------------+--------------------------------------------------------------------------------+
	//	+------------------------------------------------+--------------------------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000                      | The call was successful.                                                       |
	//	+------------------------------------------------+--------------------------------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022                | A loopback was detected.                                                       |
	//	+------------------------------------------------+--------------------------------------------------------------------------------+
	//	| STATUS_CTX_WINSTATION_ACCESS_DENIED 0xC00A002B | The server is in the process of shutting down and cannot complete the request. |
	//	+------------------------------------------------+--------------------------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationCheckLoopBack return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationCheckLoopBackResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationCheckLoopBackOperation) *xxx_WinStationCheckLoopBackOperation {
	if op == nil {
		op = &xxx_WinStationCheckLoopBackOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationCheckLoopBackResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationCheckLoopBackOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationCheckLoopBackResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationCheckLoopBackResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationCheckLoopBackOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectCallbackOperation structure represents the RpcConnectCallback operation
type xxx_ConnectCallbackOperation struct {
	Server      *tsts.Server `idl:"name:hServer" json:"server"`
	Result      uint32       `idl:"name:pResult" json:"result"`
	Timeout     uint32       `idl:"name:TimeOut" json:"timeout"`
	AddressType uint32       `idl:"name:AddressType" json:"address_type"`
	Address     []byte       `idl:"name:pAddress;size_is:(AddressSize)" json:"address"`
	AddressSize uint32       `idl:"name:AddressSize" json:"address_size"`
	Return      bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectCallbackOperation) OpNum() int { return 66 }

func (o *xxx_ConnectCallbackOperation) OpName() string { return "/IcaApi/v1/RpcConnectCallback" }

func (o *xxx_ConnectCallbackOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Address != nil && o.AddressSize == 0 {
		o.AddressSize = uint32(len(o.Address))
	}
	if o.AddressSize > uint32(4096) {
		return fmt.Errorf("AddressSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectCallbackOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TimeOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// AddressType {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.AddressType); err != nil {
			return err
		}
	}
	// pAddress {in} (1:{alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=AddressSize](uchar))
	{
		dimSize1 := uint64(o.AddressSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Address {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Address[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Address); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// AddressSize {in} (1:{range=(0,4096), alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.AddressSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectCallbackOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TimeOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// AddressType {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.AddressType); err != nil {
			return err
		}
	}
	// pAddress {in} (1:{alias=PBYTE,pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=AddressSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Address", sizeInfo[0])
		}
		o.Address = make([]byte, sizeInfo[0])
		for i1 := range o.Address {
			i1 := i1
			if err := w.ReadData(&o.Address[i1]); err != nil {
				return err
			}
		}
	}
	// AddressSize {in} (1:{range=(0,4096), alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.AddressSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectCallbackOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectCallbackOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectCallbackOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectCallbackRequest structure represents the RpcConnectCallback operation request
type ConnectCallbackRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// TimeOut: Not used.
	Timeout uint32 `idl:"name:TimeOut" json:"timeout"`
	// AddressType:  MUST be TDI_ADDRESS_TYPE_IP. For more information, see [MSDN-TDIADDRESS].
	AddressType uint32 `idl:"name:AddressType" json:"address_type"`
	// pAddress:  Pointer to the address itself. MUST be TDI_ADDRESS_IP. This is a standard
	// representation for an IP address. For more information, see [MSDN-TDIADDRESS].
	Address []byte `idl:"name:pAddress;size_is:(AddressSize)" json:"address"`
	// AddressSize:  MUST be TDI_ADDRESS_LENGTH_IP. This is a standard representation for
	// the length of an IP address. For more information, see [MSDN-TDIADDRESS].
	AddressSize uint32 `idl:"name:AddressSize" json:"address_size"`
}

func (o *ConnectCallbackRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectCallbackOperation) *xxx_ConnectCallbackOperation {
	if op == nil {
		op = &xxx_ConnectCallbackOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Timeout = o.Timeout
	op.AddressType = o.AddressType
	op.Address = o.Address
	op.AddressSize = o.AddressSize
	return op
}

func (o *ConnectCallbackRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectCallbackOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Timeout = op.Timeout
	o.AddressType = op.AddressType
	o.Address = op.Address
	o.AddressSize = op.AddressSize
}
func (o *ConnectCallbackRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectCallbackRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectCallbackOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectCallbackResponse structure represents the RpcConnectCallback operation response
type ConnectCallbackResponse struct {
	// pResult:  Failure error code if the call to RpcConnectCallback failed. If the call
	// was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000           | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_NOT_SUPPORTED 0xC00000BB     | AddressType is not TDI_ADDRESS_TYPE_IP. This is a standard representation of a   |
	//	|                                     | type for an IP address. For more information, see [MSDN-TDIADDRESS].             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_INVALID_PARAMETER 0xC000000D | AddressSize is not TDI_ADDRESS_LENGTH_IP. For more information, see              |
	//	|                                     | [MSDN-TDIADDRESS].                                                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022     | The caller is not SYSTEM.                                                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcConnectCallback return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *ConnectCallbackResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectCallbackOperation) *xxx_ConnectCallbackOperation {
	if op == nil {
		op = &xxx_ConnectCallbackOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *ConnectCallbackResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectCallbackOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *ConnectCallbackResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectCallbackResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectCallbackOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationGetAllProcessesNT6Operation structure represents the RpcWinStationGetAllProcesses_NT6 operation
type xxx_WinStationGetAllProcessesNT6Operation struct {
	Server            *tsts.Server                `idl:"name:hServer" json:"server"`
	Result            uint32                      `idl:"name:pResult" json:"result"`
	Level             uint32                      `idl:"name:Level" json:"level"`
	NumberOfProcesses uint32                      `idl:"name:pNumberOfProcesses" json:"number_of_processes"`
	AllProcessesInfo  []*tsts.AllProcessesInfoNT6 `idl:"name:ppTsAllProcessesInfo;size_is:(, pNumberOfProcesses)" json:"all_processes_info"`
	Return            bool                        `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationGetAllProcessesNT6Operation) OpNum() int { return 70 }

func (o *xxx_WinStationGetAllProcessesNT6Operation) OpName() string {
	return "/IcaApi/v1/RpcWinStationGetAllProcesses_NT6"
}

func (o *xxx_WinStationGetAllProcessesNT6Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetAllProcessesNT6Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Level {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pNumberOfProcesses {in, out} (1:{pointer=ref}*(1))(2:{range=(0,32768), alias=BOUNDED_ULONG, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.NumberOfProcesses); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetAllProcessesNT6Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Level {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pNumberOfProcesses {in, out} (1:{pointer=ref}*(1))(2:{range=(0,32768), alias=BOUNDED_ULONG, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.NumberOfProcesses); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetAllProcessesNT6Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AllProcessesInfo != nil && o.NumberOfProcesses == 0 {
		o.NumberOfProcesses = uint32(len(o.AllProcessesInfo))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetAllProcessesNT6Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// pNumberOfProcesses {in, out} (1:{pointer=ref}*(1))(2:{range=(0,32768), alias=BOUNDED_ULONG, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.NumberOfProcesses); err != nil {
			return err
		}
	}
	// ppTsAllProcessesInfo {out} (1:{pointer=ref}*(2))(2:{alias=PTS_ALL_PROCESSES_INFO_NT6}*(1))(3:{alias=TS_ALL_PROCESSES_INFO_NT6}[dim:0,size_is=pNumberOfProcesses](struct))
	{
		if o.AllProcessesInfo != nil || o.NumberOfProcesses > 0 {
			_ptr_ppTsAllProcessesInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfProcesses)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.AllProcessesInfo {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.AllProcessesInfo[i1] != nil {
						if err := o.AllProcessesInfo[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&tsts.AllProcessesInfoNT6{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.AllProcessesInfo); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&tsts.AllProcessesInfoNT6{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AllProcessesInfo, _ptr_ppTsAllProcessesInfo); err != nil {
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
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationGetAllProcessesNT6Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// pNumberOfProcesses {in, out} (1:{pointer=ref}*(1))(2:{range=(0,32768), alias=BOUNDED_ULONG, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.NumberOfProcesses); err != nil {
			return err
		}
	}
	// ppTsAllProcessesInfo {out} (1:{pointer=ref}*(2))(2:{alias=PTS_ALL_PROCESSES_INFO_NT6,pointer=ref}*(1))(3:{alias=TS_ALL_PROCESSES_INFO_NT6}[dim:0,size_is=pNumberOfProcesses](struct))
	{
		_ptr_ppTsAllProcessesInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.AllProcessesInfo", sizeInfo[0])
			}
			o.AllProcessesInfo = make([]*tsts.AllProcessesInfoNT6, sizeInfo[0])
			for i1 := range o.AllProcessesInfo {
				i1 := i1
				if o.AllProcessesInfo[i1] == nil {
					o.AllProcessesInfo[i1] = &tsts.AllProcessesInfoNT6{}
				}
				if err := o.AllProcessesInfo[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppTsAllProcessesInfo := func(ptr interface{}) { o.AllProcessesInfo = *ptr.(*[]*tsts.AllProcessesInfoNT6) }
		if err := w.ReadPointer(&o.AllProcessesInfo, _s_ppTsAllProcessesInfo, _ptr_ppTsAllProcessesInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationGetAllProcessesNT6Request structure represents the RpcWinStationGetAllProcesses_NT6 operation request
type WinStationGetAllProcessesNT6Request struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// Level:  MUST be GAP_LEVEL_BASIC (0). Any other value will lead to the method returning
	// FALSE.
	Level uint32 `idl:"name:Level" json:"level"`
	// pNumberOfProcesses:  The number of processes requested by the caller. On return,
	// this indicates the number of processes actually stored in the ppTsAllProcessesInfo
	// parameter.
	NumberOfProcesses uint32 `idl:"name:pNumberOfProcesses" json:"number_of_processes"`
}

func (o *WinStationGetAllProcessesNT6Request) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetAllProcessesNT6Operation) *xxx_WinStationGetAllProcessesNT6Operation {
	if op == nil {
		op = &xxx_WinStationGetAllProcessesNT6Operation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Level = o.Level
	op.NumberOfProcesses = o.NumberOfProcesses
	return op
}

func (o *WinStationGetAllProcessesNT6Request) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetAllProcessesNT6Operation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Level = op.Level
	o.NumberOfProcesses = op.NumberOfProcesses
}
func (o *WinStationGetAllProcessesNT6Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationGetAllProcessesNT6Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetAllProcessesNT6Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationGetAllProcessesNT6Response structure represents the RpcWinStationGetAllProcesses_NT6 operation response
type WinStationGetAllProcessesNT6Response struct {
	// pResult:  Failure error code if the call to RpcWinStationGetAllProcesses_NT6 failed.
	// If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+---------------------------------+--------------------------------------------------------+
	//	|                                 |                                                        |
	//	|              VALUE              |                        MEANING                         |
	//	|                                 |                                                        |
	//	+---------------------------------+--------------------------------------------------------+
	//	+---------------------------------+--------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000       | The call was successful.                               |
	//	+---------------------------------+--------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022 | The caller does not have permission for the operation. |
	//	+---------------------------------+--------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// pNumberOfProcesses:  The number of processes requested by the caller. On return,
	// this indicates the number of processes actually stored in the ppTsAllProcessesInfo
	// parameter.
	NumberOfProcesses uint32 `idl:"name:pNumberOfProcesses" json:"number_of_processes"`
	// ppTsAllProcessesInfo:  Pointer to an array of processes allocated and returned by
	// the method. *ppTsAllProcessesInfo is allocated by the method to be an array of TS_ALL_PROCESSES_INFO_NT6
	// structures. The array returned by the method MUST be freed by the caller.
	AllProcessesInfo []*tsts.AllProcessesInfoNT6 `idl:"name:ppTsAllProcessesInfo;size_is:(, pNumberOfProcesses)" json:"all_processes_info"`
	// Return: The RpcWinStationGetAllProcesses_NT6 return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationGetAllProcessesNT6Response) xxx_ToOp(ctx context.Context, op *xxx_WinStationGetAllProcessesNT6Operation) *xxx_WinStationGetAllProcessesNT6Operation {
	if op == nil {
		op = &xxx_WinStationGetAllProcessesNT6Operation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.NumberOfProcesses = o.NumberOfProcesses
	op.AllProcessesInfo = o.AllProcessesInfo
	op.Return = o.Return
	return op
}

func (o *WinStationGetAllProcessesNT6Response) xxx_FromOp(ctx context.Context, op *xxx_WinStationGetAllProcessesNT6Operation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.NumberOfProcesses = op.NumberOfProcesses
	o.AllProcessesInfo = op.AllProcessesInfo
	o.Return = op.Return
}
func (o *WinStationGetAllProcessesNT6Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationGetAllProcessesNT6Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationGetAllProcessesNT6Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WinStationOpenSessionDirectoryOperation structure represents the RpcWinStationOpenSessionDirectory operation
type xxx_WinStationOpenSessionDirectoryOperation struct {
	Server     *tsts.Server `idl:"name:hServer" json:"server"`
	Result     uint32       `idl:"name:pResult" json:"result"`
	ServerName string       `idl:"name:pszServerName;max_is:(64);string" json:"server_name"`
	Return     bool         `idl:"name:Return" json:"return"`
}

func (o *xxx_WinStationOpenSessionDirectoryOperation) OpNum() int { return 75 }

func (o *xxx_WinStationOpenSessionDirectoryOperation) OpName() string {
	return "/IcaApi/v1/RpcWinStationOpenSessionDirectory"
}

func (o *xxx_WinStationOpenSessionDirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	// cannot evaluate expression 65
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationOpenSessionDirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&tsts.Server{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pszServerName {in} (1:{string, alias=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,max_is=64,string,null](wchar))
	{
		dimSize1 := uint64(65)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.UTF16NLen(o.ServerName)
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
		_ServerName_buf := utf16.Encode([]rune(o.ServerName))
		if uint64(len(_ServerName_buf)) > sizeInfo[0]-1 {
			_ServerName_buf = _ServerName_buf[:sizeInfo[0]-1]
		}
		if o.ServerName != ndr.ZeroString {
			_ServerName_buf = append(_ServerName_buf, uint16(0))
		}
		for i1 := range _ServerName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_ServerName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_ServerName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_WinStationOpenSessionDirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hServer {in} (1:{context_handle, alias=SERVER_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &tsts.Server{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pszServerName {in} (1:{string, alias=PWCHAR,pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,max_is=64,string,null](wchar))
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
		var _ServerName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _ServerName_buf", sizeInfo[0])
		}
		_ServerName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _ServerName_buf {
			i1 := i1
			if err := w.ReadData(&_ServerName_buf[i1]); err != nil {
				return err
			}
		}
		o.ServerName = strings.TrimRight(string(utf16.Decode(_ServerName_buf)), ndr.ZeroString)
	}
	return nil
}

func (o *xxx_WinStationOpenSessionDirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationOpenSessionDirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WinStationOpenSessionDirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResult {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Result); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WinStationOpenSessionDirectoryRequest structure represents the RpcWinStationOpenSessionDirectory operation request
type WinStationOpenSessionDirectoryRequest struct {
	// hServer: Handle to the server object. This is of type SERVER_HANDLE. The hServer
	// argument MUST be obtained from a previous call to RpcWinStationOpenServer.
	Server *tsts.Server `idl:"name:hServer" json:"server"`
	// pszServerName: The name of the server hosting session directory to which to attempt
	// connection.
	ServerName string `idl:"name:pszServerName;max_is:(64);string" json:"server_name"`
}

func (o *WinStationOpenSessionDirectoryRequest) xxx_ToOp(ctx context.Context, op *xxx_WinStationOpenSessionDirectoryOperation) *xxx_WinStationOpenSessionDirectoryOperation {
	if op == nil {
		op = &xxx_WinStationOpenSessionDirectoryOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ServerName = o.ServerName
	return op
}

func (o *WinStationOpenSessionDirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_WinStationOpenSessionDirectoryOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ServerName = op.ServerName
}
func (o *WinStationOpenSessionDirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WinStationOpenSessionDirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationOpenSessionDirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WinStationOpenSessionDirectoryResponse structure represents the RpcWinStationOpenSessionDirectory operation response
type WinStationOpenSessionDirectoryResponse struct {
	// pResult:  Failure error code if the call to RpcWinStationOpenSessionDirectory failed.
	// If the call was successful, this parameter MUST be STATUS_SUCCESS (0x00000000).
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|                                 |                                                                                  |
	//	|              VALUE              |                                     MEANING                                      |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_UNSUCCESSFUL 0xC0000001  | The server is not in application server mode on an advanced servers Stock        |
	//	|                                 | Keeping Unit (SKU).                                                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_ACCESS_DENIED 0xC0000022 | The caller is not SYSTEM nor an administrator.                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	Result uint32 `idl:"name:pResult" json:"result"`
	// Return: The RpcWinStationOpenSessionDirectory return value.
	Return bool `idl:"name:Return" json:"return"`
}

func (o *WinStationOpenSessionDirectoryResponse) xxx_ToOp(ctx context.Context, op *xxx_WinStationOpenSessionDirectoryOperation) *xxx_WinStationOpenSessionDirectoryOperation {
	if op == nil {
		op = &xxx_WinStationOpenSessionDirectoryOperation{}
	}
	if o == nil {
		return op
	}
	op.Result = o.Result
	op.Return = o.Return
	return op
}

func (o *WinStationOpenSessionDirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_WinStationOpenSessionDirectoryOperation) {
	if o == nil {
		return
	}
	o.Result = op.Result
	o.Return = op.Return
}
func (o *WinStationOpenSessionDirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WinStationOpenSessionDirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WinStationOpenSessionDirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
