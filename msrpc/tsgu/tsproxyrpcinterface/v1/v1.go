package tsproxyrpcinterface

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
)

var (
	// import guard
	GoPackage = "tsgu"
)

var (
	// Syntax UUID
	InterfaceSyntaxUUID = &uuid.UUID{TimeLow: 0x44e265dd, TimeMid: 0x7daf, TimeHiAndVersion: 0x42cd, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x60, Node: [6]uint8{0x3c, 0xdb, 0x6e, 0x7a, 0x27, 0x29}}
	// Syntax ID
	InterfaceSyntaxV1_3 = &dcerpc.SyntaxID{IfUUID: InterfaceSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 3}
)

// TsProxyRpcInterface interface.
type InterfaceClient interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// The TsProxyCreateTunnel method sets up the tunnel (2) in which all further communication
	// between the RDG client and the RDG server occurs. This is also used to exchange versioning
	// and capability information between the RDG client and RDG server. It is used to exchange
	// the RDG server certificate which has already been used to register for an authentication
	// service. After this method call has successfully been completed, a tunnel (2) shutdown
	// can be performed. This is accomplished by using the TsProxyCloseTunnel method call.
	//
	// Prerequisites: The connection state MUST be in Start state.
	//
	// Return Values: The method MUST return ERROR_SUCCESS on success. Other failures MUST
	// be one of the codes listed in the rest of this table. The client MAY interpret failures
	// in any way it deems appropriate. See section 2.2.6 for details on these errors.
	//
	//	+----------------------------------------------------------+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                          RETURN                          |                         STATE                          |                                                                                  |
	//	|                          VALUE                           |                       TRANSITION                       |                                   DESCRIPTION                                    |
	//	|                                                          |                                                        |                                                                                  |
	//	+----------------------------------------------------------+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------------+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS (0x00000000)                               | The connection MUST transition to the connected state. | Returned when a call to the TsProxyCreateTunnel method succeeds.                 |
	//	+----------------------------------------------------------+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| E_PROXY_INTERNALERROR (0x800759D8)                       | The connection MUST transition to end state.           | Returned when the server encounters an unexpected error. The RDG client MUST end |
	//	|                                                          |                                                        | the protocol when this error is received.                                        |
	//	+----------------------------------------------------------+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| E_PROXY_COOKIE_BADPACKET (0x800759F7)                    | The connection MUST transition to end state.           | Returned if the packetAuth field of the TSGPacket parameter is NULL.             |
	//	+----------------------------------------------------------+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| E_PROXY_NOCERTAVAILABLE (0x800759EE)                     | The connection MUST transition to end state.           | Returned when the RDG server cannot find a certificate to register for SCHANNEL  |
	//	|                                                          |                                                        | Authentication Service (AS). The RDG client MUST end the protocol when this      |
	//	|                                                          |                                                        | error is received.                                                               |
	//	+----------------------------------------------------------+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| E_PROXY_UNSUPPORTED_AUTHENTICATION_METHOD(0x800759F9)    | The connection MUST transition to end state.           | Returned to the RDG client when the RDG server is configured for pluggable       |
	//	|                                                          |                                                        | authentication and the value of the packetId member of the TSGPacket parameter   |
	//	|                                                          |                                                        | is not equal to TSG_PACKET_TYPE_AUTH or TSG_PACKET_TYPE_REAUTH. The RDG server   |
	//	|                                                          |                                                        | MUST disconnect the connection.                                                  |
	//	+----------------------------------------------------------+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| E_PROXY_COOKIE_AUTHENTICATION_ACCESS_DENIED (0x800759F8) | The connection MUST transition to end state.           | Returned when the given user does not have access to connect via RDG server.     |
	//	|                                                          |                                                        | The RDG server MUST be in pluggable authentication mode for this error to be     |
	//	|                                                          |                                                        | returned.                                                                        |
	//	+----------------------------------------------------------+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| E_PROXY_CAPABILITYMISMATCH (0x800759E9)                  | The connection MUST transition to end state.           | Returned when the RDG server supports the TSG_MESSAGING_CAP_CONSENT_SIGN         |
	//	|                                                          |                                                        | capability and is configured to allow only a RDG client that supports the        |
	//	|                                                          |                                                        | TSG_MESSAGING_CAP_CONSENT_SIGN capability, but the RDG client doesn't support    |
	//	|                                                          |                                                        | the capability.                                                                  |
	//	+----------------------------------------------------------+--------------------------------------------------------+----------------------------------------------------------------------------------+
	CreateTunnel(context.Context, *CreateTunnelRequest, ...dcerpc.CallOption) (*CreateTunnelResponse, error)

	// The TsProxyAuthorizeTunnel method is used to authorize the tunnel (2) based on rules
	// defined by the RDG server. The RDG server SHOULD perform security authorization for
	// the RDG client. The RDG server SHOULD<39> also use this method to require health
	// checks from the RDG client, which SHOULD result in the RDG client performing health
	// remediation.<40> After this method call has successfully been completed, a tunnel
	// (2) shutdown can be performed. If there are existing channels within the tunnel,
	// the RDG server MUST close all the channels before the tunnel shutdown. The tunnel
	// (2) shutdown is accomplished by using the TsProxyCloseTunnel method call.
	//
	// If this method call completes successfully, the ADM element Number of Connections
	// MUST be incremented by 1.
	//
	// Prerequisites: The connection MUST be in Connected state. If this call is made in
	// any other state, the result is undefined.
	//
	// Return Values: The method MUST return ERROR_SUCCESS on success. Other failures MUST
	// be one of the codes listed. The client MAY interpret failures in any way it deems
	// appropriate. See 2.2.6 for details on these errors.
	//
	//	+------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                        RETURN                        |                             STATE                             |                                                                                  |
	//	|                        VALUE                         |                          TRANSITION                           |                                   DESCRIPTION                                    |
	//	|                                                      |                                                               |                                                                                  |
	//	+------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS (0x00000000)                           | The connection MUST transition to the authorized state.       | Returned when a call to the TsProxyAuthorizeTunnel method succeeds.              |
	//	+------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| E_PROXY_NAP_ACCESSDENIED (0x800759DB)                | The connection MUST transition to Tunnel Close Pending state. | Returned when the RDG server denies the RDG client access due to policy. The RDG |
	//	|                                                      |                                                               | client MUST end the protocol when this error is received.                        |
	//	+------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_NOTSUPPORTED) (0x000059E8)      | The connection MUST transition to Tunnel Close Pending state. | Returned if the packetId field of the TSGPacket parameter is not                 |
	//	|                                                      |                                                               | TSG_PACKET_TYPE_QUARREQUEST. The RDG client MUST end the protocol when this      |
	//	|                                                      |                                                               | error is received.                                                               |
	//	+------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| E_PROXY_QUARANTINE_ACCESSDENIED (0x800759ED)         | The connection MUST transition to Tunnel Close Pending state. | Returned when the RDG server rejects the connection due to quarantine policy.    |
	//	|                                                      |                                                               | The RDG client MUST end the protocol when this error is received.                |
	//	+------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED (0x00000005)                     | The connection MUST transition to Tunnel Close Pending state. | Returned when this call is made either in a state other than the Connected state |
	//	|                                                      |                                                               | or the tunnelContext parameter is NULL. The RDG client MUST end the protocol     |
	//	|                                                      |                                                               | when this error is received.                                                     |
	//	+------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_MAXCONNECTIONSREACHED) (0x59E6) | The connection MUST transition to end state.                  | Returned when the ADM element Number of Connections is equal to the maximum      |
	//	|                                                      |                                                               | number of connections when the call is made.<43> The RDG client MUST end the     |
	//	|                                                      |                                                               | protocol when this error is received.                                            |
	//	+------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER (0x00000057)                 | The connection MUST not transition its state.                 | Returned when the Negotiated Capabilities ADM element contains                   |
	//	|                                                      |                                                               | TSG_NAP_CAPABILITY_QUAR_SOH and TSGPacket->TSGPacket.packetQuarRequest->dataLen  |
	//	|                                                      |                                                               | is not zero and TSGPacket->TSGPacket.packetQuarRequest->data is not NULL and     |
	//	|                                                      |                                                               | TSGPacket->TSGPacket.packetQuarRequest->data is not prefixed with Nonce.         |
	//	+------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	AuthorizeTunnel(context.Context, *AuthorizeTunnelRequest, ...dcerpc.CallOption) (*AuthorizeTunnelResponse, error)

	// The TsProxyMakeTunnelCall method is designed to be used as a general purpose API.
	// If both the client and the server support the administrative message, the client
	// MAY request the same from the RDG server. If the RDG server has any administrative
	// messages, it SHOULD complete the pending call at this point in time. After a call
	// to TsProxyMakeTunnelCall returns, the RDG client SHOULD queue up another request
	// at this point in time. During the shutdown sequence, the client MUST make this call,
	// if a request is pending on the RDG server, to cancel the administrative message request.
	//
	// Prerequisites: The connection MUST be in Authorized state or Channel Created state
	// or Pipe Created state or Channel Close Pending state or Tunnel Close Pending state.
	// If this call is made in any other state, the error ERROR_ACCESS_DENIED is returned.
	//
	// Return Values: The method MUST return ERROR_SUCCESS on success. Other failures MUST
	// be one of the codes listed. The client MAY interpret failures in any way it deems
	// appropriate. See section 2.2.6 for details on these errors. The connection MUST NOT
	// transition its state after completing the TsProxyMakeTunnelCall.
	//
	//	+------------------------------------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                        RETURN                        |                     STATE                     |                                                                                  |
	//	|                        VALUE                         |                  TRANSITION                   |                                   DESCRIPTION                                    |
	//	|                                                      |                                               |                                                                                  |
	//	+------------------------------------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS (0x00000000)                           | The connection MUST NOT transition its state. | Returned when a call to the TsProxyMakeTunnelCall method succeeds.               |
	//	+------------------------------------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED (0x00000005)                     | The connection MUST NOT transition its state. | Returned in the following cases. When the call is made in any                    |
	//	|                                                      |                                               | state other than Authorized, Channel Created, Pipe Created, Channel              |
	//	|                                                      |                                               | Close Pending, or Tunnel Close Pending. If procId is neither                     |
	//	|                                                      |                                               | TSG_TUNNEL_CALL_ASYNC_MSG_REQUEST nor TSG_TUNNEL_CANCEL_ASYNC_MSG_REQUEST.       |
	//	|                                                      |                                               | If procId is TSG_TUNNEL_CALL_ASYNC_MSG_REQUEST and there is already a call to    |
	//	|                                                      |                                               | TsProxyMakeTunnelCall made earlier with procId TSG_TUNNEL_CALL_ASYNC_MSG_REQUEST |
	//	|                                                      |                                               | and it is not yet returned. If procId is TSG_TUNNEL_CANCEL_ASYNC_MSG_REQUEST     |
	//	|                                                      |                                               | and there is no call to TsProxyMakeTunnelCall made earlier with procId           |
	//	|                                                      |                                               | TSG_TUNNEL_CALL_ASYNC_MSG_REQUEST that is not yet returned. If the tunnelContext |
	//	|                                                      |                                               | parameter is NULL. If the tunnel is not authorized. If the Reauthentication      |
	//	|                                                      |                                               | Connection ADM element is TRUE. The RDG client MUST end the protocol when this   |
	//	|                                                      |                                               | error is received.                                                               |
	//	+------------------------------------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_FROM_WIN32(RPC_S_CALL_CANCELLED)(0x8007071A) | The connection MUST not transition its state. | Returned when the call is canceled by the RDG client or the call is canceled     |
	//	|                                                      |                                               | because a shutdown sequence is initiated.                                        |
	//	+------------------------------------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	MakeTunnelCall(context.Context, *MakeTunnelCallRequest, ...dcerpc.CallOption) (*MakeTunnelCallResponse, error)

	// The TsProxyCreateChannel method is used to create a channel between the RDG client
	// and the RDG server.<44> The RDG server SHOULD connect to the target server during
	// this call to start communication between the RDG client and target server. If connection
	// to the target server cannot be done, the RDG server MUST return HRESULT_CODE(E_PROXY_TS_CONNECTFAILED)
	// as noted in the Return Values section.<45> The RDG server MUST return a representation
	// of the channel to the RDG client. After this method call has successfully been completed,
	// a channel shutdown can be performed by using the TsProxyCloseChannel method. Please
	// refer to section 3.1.1 for a state transition diagram.
	//
	// Prerequisites: The tunnel MUST be authorized; otherwise, the error ERROR_ACCESS_DENIED
	// is returned.
	//
	// Return Values: The method MUST return ERROR_SUCCESS on success. Other failures MUST
	// be one of the codes listed. The client MAY interpret failures in any way it deems
	// appropriate. See section 2.2.6 for details on these errors.
	//
	//	+-----------------------------------------------------+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                       RETURN                        |                          STATE                           |                                                                                  |
	//	|                        VALUE                        |                        TRANSITION                        |                                   DESCRIPTION                                    |
	//	|                                                     |                                                          |                                                                                  |
	//	+-----------------------------------------------------+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------------------+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS (0x00000000)                          | The connection MUST transition to Channel Created state. | Returned when a call to the TsProxyCreateChannel method succeeds.                |
	//	+-----------------------------------------------------+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED (0x00000005)                    | The connection MUST NOT transition its state.            | Returned either if tunnelContext parameter is NULL, if this method is called on  |
	//	|                                                     |                                                          | a tunnel which is not authorized, if the tsEndPointInfo parameter is NULL, or if |
	//	|                                                     |                                                          | the numResourceNames member of the tsEndPointInfo parameter is zero.             |
	//	+-----------------------------------------------------+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| E_PROXY_RAP_ACCESSDENIED (0x800759DA)               | The connection MUST NOT transition its state.            | Returned when an attempt to resolve or access a target server is blocked by RDG  |
	//	|                                                     |                                                          | server policies.                                                                 |
	//	+-----------------------------------------------------+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| E_PROXY_INTERNALERROR (0x800759D8)                  | The connection MUST NOT transition its state.            | Returned when the server encounters an unexpected error while creating the       |
	//	|                                                     |                                                          | channel.                                                                         |
	//	+-----------------------------------------------------+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_TS_CONNECTFAILED) (0x000059DD) | The connection MUST NOT transition its state.            | This error is returned in rpc_fault packet when the RDG server fails to          |
	//	|                                                     |                                                          | connect to any of the target server names, as specified in the members of        |
	//	|                                                     |                                                          | tsEndPointInfo.                                                                  |
	//	+-----------------------------------------------------+----------------------------------------------------------+----------------------------------------------------------------------------------+
	CreateChannel(context.Context, *CreateChannelRequest, ...dcerpc.CallOption) (*CreateChannelResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// The TsProxyCloseChannel method is used to terminate the channel from the RDG client
	// to the RDG server. This SHOULD be called only if the RDG client has not received
	// the RPC response PDU with the PFC_LAST_FRAG bit set in the pfc_flags field. All communication
	// between the RDG client and the target server MUST stop after the RDG server executes
	// this method. The RDG client MUST NOT use this context handle in any subsequent operations
	// after calling this method. This will terminate the channel between the RDG client
	// and the target server. If the RDG server has not already sent the RPC response PDU
	// with the PFC_LAST_FRAG bit set in the pfc_flags field, which happens if the RDG server
	// initiated the disconnect, the RDG client will also receive a return code for TsProxySetupReceivePipe
	// in an RPC response PDU with the PFC_LAST_FRAG bit set in the pfc_flags. For a description
	// of RPC response PDU, pfc_flags, and PFC_LAST_FRAG, refer to [C706] sections 12.6.2
	// and 12.6.14.10.
	//
	// The RDG server completes the TsProxyCloseChannel only after sending all of the data
	// it received before this call was made. The RDG client receives the call complete
	// notification only after it receives all of the data that was sent by the RDG server
	// before completing TsProxyCloseChannel. Please refer to section 3.2.6.2.2 for details
	// on how the data is ensured to reach the destination.
	//
	// Prerequisites: The connection MUST be in Channel Created state or Pipe Created state
	// or Channel Close Pending state.
	//
	// Return Values:
	//
	//	+----------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                             STATE                             |                                                                                  |
	//	|              VALUE               |                          TRANSITION                           |                                   DESCRIPTION                                    |
	//	|                                  |                                                               |                                                                                  |
	//	+----------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS (0x00000000)       | The connection MUST transition to Tunnel Close Pending state. | Returned when the call to the TsProxyCloseChannel method succeeds.               |
	//	+----------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED (0x00000005) | The connection MUST NOT transition its state.                 | Returned when the provided context parameter is NULL or not a valid channel      |
	//	|                                  |                                                               | context handle.                                                                  |
	//	+----------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	CloseChannel(context.Context, *CloseChannelRequest, ...dcerpc.CallOption) (*CloseChannelResponse, error)

	// The TsProxyCloseTunnel method is used to terminate the tunnel (1) between the RDG
	// client and the RDG server. All communication between the RDG client and RDG server
	// MUST stop after the RDG server executes this method. The RDG client MUST NOT use
	// this tunnel context handle in any subsequent operations after this method call. This
	// MUST be the final tear down phase of the RDG client to RDG server tunnel. If the
	// ADM element Reauthentication Connection is FALSE, then the ADM element Number of
	// Connections MUST be decremented by 1 in this call. If there is an existing channel
	// within the tunnel, it SHOULD first be closed using TsProxyCloseChannel. If the RDG
	// client calls the TsProxyCloseTunnel method before calling the TsProxyCloseChannel
	// method, the RDG server MUST close the channel and then close the tunnel.
	//
	// Prerequisites: The connection MUST be in any of the following states: Connected state,
	// Authorized state, Channel Created state, Pipe Created state, Channel Close Pending
	// state, or Tunnel Close Pending state.
	//
	// Return Values: The method MUST return 0 on success. This function SHOULD NOT fail
	// from a RDG protocol perspective. If TsProxyCloseTunnel is called while any of the
	// channels are not closed, then the RDG server MUST close all the channels and then
	// close the tunnel.
	//
	//	+----------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                      STATE                       |                                                                                  |
	//	|              VALUE               |                    TRANSITION                    |                                   DESCRIPTION                                    |
	//	|                                  |                                                  |                                                                                  |
	//	+----------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS (0x00000000)       | The connection MUST transition to the end state. | Returned when a call to the TsProxyCloseTunnel method succeeds.                  |
	//	+----------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED (0x00000005) | The connection MUST NOT transition its state.    | Returned when the provided context parameter is NULL or not a valid tunnel       |
	//	|                                  |                                                  | context handle.                                                                  |
	//	+----------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+
	CloseTunnel(context.Context, *CloseTunnelRequest, ...dcerpc.CallOption) (*CloseTunnelResponse, error)

	// The TsProxySetupReceivePipe method is used for data transfer from the RDG server
	// to the RDG client. The RDG server MUST create an RPC out pipe upon receiving this
	// method call from the RDG client. This call bypasses the NDR and hence, the RPC runtime
	// MUST NOT perform a strict NDR data consistency check for this method. Refer to section
	// 3.6.5 for details on NDR-bypassing. Section 3.6.5.4 and section 3.6.5.5 give details
	// on wire representation of data for responses to TsProxySetupReceivePipe. The out
	// pipe MUST be created by the RDG server in the same manner as NDR creates it for a
	// call.<49> The RDG server MUST use this out pipe and Stub Data field in RPC response
	// PDUs to send all data from the target server to the RDG client on the channel. The
	// RDG client MUST use this out pipe to pull data from the target server on the channel.
	// On connection disconnect, the RDG server MUST send the following on the pipe: A DWORD
	// return code in an RPC response PDU and set the PFC_LAST_FRAG bit in the pfc_flags
	// field of the RPC response PDU. The pipe close is indicated when the PFC_LAST_FRAG
	// bit is set in the pfc_flags field of the RPC response PDU. When the RDG client sees
	// that the PFC_LAST_FRAG bit is set in the pfc_flags field of the RPC response PDU,
	// it MUST interpret the 4 bytes Stub Data as the return code of TsProxySetupReceivePipe.
	// For a description of RPC response PDU, pfc_flags, PFC_LAST_FRAG, and Stub Data, refer
	// to sections 12.6.2 and 12.6.4.10 in [C706]. The RDG client and RDG server MUST negotiate
	// a separate out pipe for each channel. Out pipes MUST NOT be used or shared across
	// channels.<50>
	//
	// As long as the channel is not closed, the RPC and Transport layer guarantee that
	// any data that is sent by the RDG server reaches the RDG client. RPC and Transport
	// layer also ensure that the data is delivered to the RDG client in the order it was
	// sent by the RDG server.
	//
	// After the call reaches the RDG server, the connection MUST transition to Pipe Created
	// state after setting up the out pipe.
	//
	// Prerequisites: The connection MUST be in Channel Created state. If this is called
	// in any other state, then the behavior is undefined.
	//
	// Return Values: The method MUST return ERROR_GRACEFUL_DISCONNECT on success, that
	// is, if the RDG client gracefully disconnects the connection by calling TsProxyCloseChannel.
	// Other failures MUST be one of the codes listed. The client MAY interpret failures
	// in any way it deems appropriate. See section 2.2.6 for details on these errors.
	//
	// The error DWORD value is always sent, when the receive pipe closes down. The receive
	// pipe will always close down when a disconnect takes place.
	//
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                           RETURN                           |                             STATE                             |                                                                                  |
	//	|                           VALUE                            |                          TRANSITION                           |                                   DESCRIPTION                                    |
	//	|                                                            |                                                               |                                                                                  |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED (0x00000005)                           | The connection MUST transition to Tunnel Close Pending state. | Returned either if this method is called before TsProxyCreateChannel or if the   |
	//	|                                                            |                                                               | Channel Context Handle ADM element is NULL. The RDG client MUST end the protocol |
	//	|                                                            |                                                               | when this error is received.                                                     |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_INTERNALERROR) (0x000059D8)           | The connection MUST transition to Tunnel Close Pending state. | Returned when an unexpected error occurs in TsProxySetupReceivePipe. The RDG     |
	//	|                                                            |                                                               | client MUST end the protocol when this error is received.                        |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_TS_CONNECTFAILED) (0x000059DD)        | The connection MUST transition to Tunnel Close Pending state. | Returned when the RDG server fails to connect to target server. It is returned   |
	//	|                                                            |                                                               | in an rpc_fault packet.<52> The RDG client MUST end the protocol when this error |
	//	|                                                            |                                                               | is received.                                                                     |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_SESSIONTIMEOUT) (0x000059F6)          | The connection MUST transition to Tunnel Close Pending state. | Returned by RDG server if a session timeout occurs and "disconnect on session    |
	//	|                                                            |                                                               | timeout" is configured at the RDG server and the ADM element Negotiated          |
	//	|                                                            |                                                               | Capabilities contains TSG_NAP_CAPABILITY_IDLE_TIMEOUT. The RDG client MUST end   |
	//	|                                                            |                                                               | the protocol when this error is received.                                        |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_REAUTH_AUTHN_FAILED) (0x000059FA)     | The connection MUST transition to Tunnel Close Pending state. | Returned when a reauthentication attempt by the client has failed because the    |
	//	|                                                            |                                                               | user credentials are no longer valid and the ADM element Negotiated Capabilities |
	//	|                                                            |                                                               | contains TSG_NAP_CAPABILITY_IDLE_TIMEOUT. The RDG client MUST end the protocol   |
	//	|                                                            |                                                               | when this error is received.                                                     |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_REAUTH_CAP_FAILED) (0x000059FB)       | The connection MUST transition to Tunnel Close Pending state. | Returned when a reauthentication attempt by the client has failed because the    |
	//	|                                                            |                                                               | user is not authorized to connect through the RDG server anymore and the ADM     |
	//	|                                                            |                                                               | element Negotiated Capabilities contains TSG_NAP_CAPABILITY_IDLE_TIMEOUT. The    |
	//	|                                                            |                                                               | RDG client MUST end the protocol when this error is received.                    |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_REAUTH_RAP_FAILED) (0x000059FC)       | The connection MUST transition to Tunnel Close Pending state. | Returned when a reauthentication attempt by the client has failed because the    |
	//	|                                                            |                                                               | user is not authorized to connect to the given end resource anymore and the ADM  |
	//	|                                                            |                                                               | element Negotiated Capabilities contains TSG_NAP_CAPABILITY_IDLE_TIMEOUT. The    |
	//	|                                                            |                                                               | RDG client MUST end the protocol when this error is received.                    |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_CONNECTIONABORTED) (0x000004D4)       | The connection MUST transition to Tunnel Close Pending state. | Returned when the following happens: The RDG server administrator forcefully     |
	//	|                                                            |                                                               | disconnects the connection. Or when the ADM element Negotiated Capabilities      |
	//	|                                                            |                                                               | doesn't contain TSG_NAP_CAPABILITY_IDLE_TIMEOUT and any one of the following     |
	//	|                                                            |                                                               | happens: Session timeout occurs and disconnect on session timeout is configured  |
	//	|                                                            |                                                               | at the RDG server. Reauthentication attempt by the client has failed because     |
	//	|                                                            |                                                               | the user credentials are no longer valid. Reauthentication attempt by the client |
	//	|                                                            |                                                               | has failed because the user is not authorized to connect through the RDG server  |
	//	|                                                            |                                                               | anymore. Reauthentication attempt by the client has failed because the user is   |
	//	|                                                            |                                                               | not authorized to connect to the given end resource anymore. Reauthentication    |
	//	|                                                            |                                                               | attempt by the RDG client has failed because the health of the user's computer   |
	//	|                                                            |                                                               | is no longer compliant with the RDG server configuration. The RDG client MUST    |
	//	|                                                            |                                                               | end the protocol when this error is received.                                    |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_SDR_NOT_SUPPORTED_BY_TS) (0x000059FD) | The connection MUST transition to Tunnel Close Pending state. | The RDG server is capable of exchanging policies with some target servers.       |
	//	|                                                            |                                                               | The RDG server MAY be configured to allow connections to only target servers     |
	//	|                                                            |                                                               | that are capable of policy exchange. If such a setting is configured and the     |
	//	|                                                            |                                                               | target server is not capable of exchanging policies with the RDG server, this    |
	//	|                                                            |                                                               | error will be returned. The RDG client MUST end the protocol when this error is  |
	//	|                                                            |                                                               | received.                                                                        |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_GRACEFUL_DISCONNECT (0x000004CA)                     | The connection MUST transition to Tunnel Close Pending state. | Returned when the connection is disconnected gracefully by the RDG client        |
	//	|                                                            |                                                               | calling TsProxyCloseChannel.                                                     |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_REAUTH_NAP_FAILED) (0x00005A00)       | The connection MUST transition to Tunnel Close Pending state. | Returned when a reauthentication attempt by the RDG client has failed            |
	//	|                                                            |                                                               | because the user's computer's health is no longer compliant with the RDG         |
	//	|                                                            |                                                               | server configuration and the ADM element Negotiated Capabilities contains        |
	//	|                                                            |                                                               | TSG_NAP_CAPABILITY_IDLE_TIMEOUT. The RDG client MUST end the protocol when this  |
	//	|                                                            |                                                               | error is received.                                                               |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_OPERATION_ABORTED(0x000003E3)                        | The connection MUST transition to Tunnel Close Pending state. | Returned when the call to TsProxySetupReceivePipe is received after the          |
	//	|                                                            |                                                               | Connection Timer has expired.                                                    |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BAD_ARGUMENTS(0x000000A0)                            | The connection MUST transition to Tunnel Close Pending state. | Returned when the target server unexpectedly closes the connection between the   |
	//	|                                                            |                                                               | RDG server and the target server.                                                |
	//	+------------------------------------------------------------+---------------------------------------------------------------+----------------------------------------------------------------------------------+
	SetupReceivePipe(context.Context, *SetupReceivePipeRequest, ...dcerpc.CallOption) (*SetupReceivePipeResponse, error)

	// The method is used for data transfer from the RDG client to the target server, via
	// the RDG server. The RDG server SHOULD send the buffer data received in this method
	// to the target server. The RPC runtime MUST NOT perform a strict NDR data consistency
	// check for this method. The Remote Desktop Gateway Server Protocol bypasses NDR for
	// this method. The wire data MUST follow the regular RPC specifications as specified
	// in [C706] section 2.1, and [MS-RPCE] minus all NDR headers, trailers, and NDR-specific
	// payload. The RDG server MUST have created the channel to the target server before
	// completing this method call. This method MAY be called multiple times by the RDG
	// client, but only after the previous method call finishes. The RDG server MUST handle
	// multiple sequential invocations of this method call. This method bypasses NDR. For
	// this reason, unlike other RPC methods that return an HRESULT, this method returns
	// a DWORD. This is directly passed to the callee from underlying RPC calls.<48> When
	// this call fails, the RDG server MUST send the final response to TsProxySetupReceivePipe
	// call.
	//
	// Prerequisites: The connection MUST be in Pipe Created state. If this call is made
	// in any other state, ERROR_ONLY_IF_CONNECTED or E_PROXY_TS_CONNECTFAILED is returned.
	//
	// Return Values: The method MUST return ERROR_SUCCESS on success. Other failures MUST
	// be one of the codes listed. The client MAY interpret failures in any way it deems
	// appropriate. See section 2.2.6 for details on these errors.
	//
	//	+--------------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                      |                             STATE                              |                                                                                  |
	//	|                      VALUE                       |                           TRANSITION                           |                                   DESCRIPTION                                    |
	//	|                                                  |                                                                |                                                                                  |
	//	+--------------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS (0x00000000)                       | The connection MUST remain in PipeCreated state.               | Returned when a call to the TsProxySendToServer method succeeds.                 |
	//	+--------------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ONLY_IF_CONNECTED (0x000004E3)             | The connection MUST transition to Channel Close Pending state. | Returned by the RDG server when an attempt is made by the client to send data    |
	//	|                                                  |                                                                | to the target server on connection state other than Pipe Created state. The RDG  |
	//	|                                                  |                                                                | client MUST end the protocol when this error is returned.                        |
	//	+--------------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED (0x00000005)                 | The connection MUST transition to Channel Close Pending state. | Returned if the channel context handle passed in the pRpcMessage parameter is    |
	//	|                                                  |                                                                | NULL. The RDG client MUST end the protocol when this error is received.          |
	//	+--------------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| HRESULT_CODE(E_PROXY_INTERNALERROR) (0x000059D8) | The connection MUST transition to Channel Close Pending state. | Returned when an unexpected error occurs in TsProxySendToServer. The RDG client  |
	//	|                                                  |                                                                | MUST end the protocol when this error is received.                               |
	//	+--------------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------------------------+
	SendToServer(context.Context, *SendToServerRequest, ...dcerpc.CallOption) (*SendToServerResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// MaxResourceNames represents the MAX_RESOURCE_NAMES RPC constant
var MaxResourceNames = 50

// PacketTypeHeader represents the TSG_PACKET_TYPE_HEADER RPC constant
var PacketTypeHeader = 18500

// PacketTypeVersionCaps represents the TSG_PACKET_TYPE_VERSIONCAPS RPC constant
var PacketTypeVersionCaps = 22083

// PacketTypeQuarantineConfigRequest represents the TSG_PACKET_TYPE_QUARCONFIGREQUEST RPC constant
var PacketTypeQuarantineConfigRequest = 20803

// PacketTypeQuarantineRequest represents the TSG_PACKET_TYPE_QUARREQUEST RPC constant
var PacketTypeQuarantineRequest = 20818

// PacketTypeResponse represents the TSG_PACKET_TYPE_RESPONSE RPC constant
var PacketTypeResponse = 20562

// PacketTypeQuarantineEncResponse represents the TSG_PACKET_TYPE_QUARENC_RESPONSE RPC constant
var PacketTypeQuarantineEncResponse = 17746

// CapabilityTypeNap represents the TSG_CAPABILITY_TYPE_NAP RPC constant
var CapabilityTypeNap = 1

// PacketTypeCapsResponse represents the TSG_PACKET_TYPE_CAPS_RESPONSE RPC constant
var PacketTypeCapsResponse = 17232

// PacketTypeMessageRequestPacket represents the TSG_PACKET_TYPE_MSGREQUEST_PACKET RPC constant
var PacketTypeMessageRequestPacket = 18258

// PacketTypeMessagePacket represents the TSG_PACKET_TYPE_MESSAGE_PACKET RPC constant
var PacketTypeMessagePacket = 18256

// PacketTypeAuth represents the TSG_PACKET_TYPE_AUTH RPC constant
var PacketTypeAuth = 16468

// PacketTypeReauth represents the TSG_PACKET_TYPE_REAUTH RPC constant
var PacketTypeReauth = 21072

// AsyncMessageConsentMessage represents the TSG_ASYNC_MESSAGE_CONSENT_MESSAGE RPC constant
var AsyncMessageConsentMessage = 1

// AsyncMessageServiceMessage represents the TSG_ASYNC_MESSAGE_SERVICE_MESSAGE RPC constant
var AsyncMessageServiceMessage = 2

// AsyncMessageReauth represents the TSG_ASYNC_MESSAGE_REAUTH RPC constant
var AsyncMessageReauth = 3

// TunnelCallAsyncMessageRequest represents the TSG_TUNNEL_CALL_ASYNC_MSG_REQUEST RPC constant
var TunnelCallAsyncMessageRequest = 1

// TunnelCancelAsyncMessageRequest represents the TSG_TUNNEL_CANCEL_ASYNC_MSG_REQUEST RPC constant
var TunnelCancelAsyncMessageRequest = 2

// NapCapabilityQuarantineSOH represents the TSG_NAP_CAPABILITY_QUAR_SOH RPC constant
var NapCapabilityQuarantineSOH = 1

// NapCapabilityIdleTimeout represents the TSG_NAP_CAPABILITY_IDLE_TIMEOUT RPC constant
var NapCapabilityIdleTimeout = 2

// MessagingCapConsentSign represents the TSG_MESSAGING_CAP_CONSENT_SIGN RPC constant
var MessagingCapConsentSign = 4

// MessagingCapServiceMessage represents the TSG_MESSAGING_CAP_SERVICE_MSG RPC constant
var MessagingCapServiceMessage = 8

// MessagingCapReauth represents the TSG_MESSAGING_CAP_REAUTH RPC constant
var MessagingCapReauth = 16

// TunnelNoSerialize structure represents PTUNNEL_CONTEXT_HANDLE_NOSERIALIZE RPC structure.
type TunnelNoSerialize dcetypes.ContextHandle

func (o *TunnelNoSerialize) ContextHandle() *dcetypes.ContextHandle {
	return (*dcetypes.ContextHandle)(o)
}

func (o *TunnelNoSerialize) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TunnelNoSerialize) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TunnelNoSerialize) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ChannelNoSerialize structure represents PCHANNEL_CONTEXT_HANDLE_NOSERIALIZE RPC structure.
type ChannelNoSerialize dcetypes.ContextHandle

func (o *ChannelNoSerialize) ContextHandle() *dcetypes.ContextHandle {
	return (*dcetypes.ContextHandle)(o)
}

func (o *ChannelNoSerialize) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChannelNoSerialize) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ChannelNoSerialize) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TunnelSerialize structure represents PTUNNEL_CONTEXT_HANDLE_SERIALIZE RPC structure.
type TunnelSerialize dcetypes.ContextHandle

func (o *TunnelSerialize) ContextHandle() *dcetypes.ContextHandle {
	return (*dcetypes.ContextHandle)(o)
}

func (o *TunnelSerialize) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TunnelSerialize) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TunnelSerialize) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ChannelSerialize structure represents PCHANNEL_CONTEXT_HANDLE_SERIALIZE RPC structure.
type ChannelSerialize dcetypes.ContextHandle

func (o *ChannelSerialize) ContextHandle() *dcetypes.ContextHandle {
	return (*dcetypes.ContextHandle)(o)
}

func (o *ChannelSerialize) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChannelSerialize) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ChannelSerialize) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EndpointInfo structure represents TSENDPOINTINFO RPC structure.
//
// The TSENDPOINTINFO structure contains information about the target server to which
// the RDG server attempts to connect.
type EndpointInfo struct {
	// resourceName:  An array of RESOURCENAME strings, as specified in section 2.2.1.1.
	// The range is from 0 to numResourceNames. This array, in conjunction with alternateResourceNames
	// parameter array, comprises the alias names of the target server to which the RDG
	// server can connect. As specified in the Protocol Overview (section 1.3), the RDG
	// server acts as a proxy to target server. The RDP client and target server MUST use
	// [MS-RDPBCGR] to communicate.
	ResourceName []string `idl:"name:resourceName;size_is:(numResourceNames)" json:"resource_name"`
	// numResourceNames:  The number of RESOURCENAME datatypes in the resourceName array.
	// The value MUST be in the range of 1 to 50, inclusive.
	ResourceNamesLength uint32 `idl:"name:numResourceNames" json:"resource_names_length"`
	// alternateResourceNames:  An array of RESOURCENAME strings to be used as alternative
	// names for the target server. The range is from 0 to numAlternateResourceNames.<10>
	AlternateResourceNames []string `idl:"name:alternateResourceNames;size_is:(numAlternateResourceNames);pointer:unique" json:"alternate_resource_names"`
	// numAlternateResourceNames:  The number of allowed alternateResourceNames. The value
	// MUST be in the range of 0 to 3, inclusive.
	AlternateResourceNamesLength uint16 `idl:"name:numAlternateResourceNames" json:"alternate_resource_names_length"`
	// Port:  Specifies the protocol ID and TCP port number for the target server endpoint
	// to which the RDG server connects. The protocol ID is in the low order 16 bits of
	// this field and port number is in the high order 16 bits. The value of the protocol
	// ID must be set to 3.
	Port uint32 `idl:"name:Port" json:"port"`
}

func (o *EndpointInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.ResourceName != nil && o.ResourceNamesLength == 0 {
		o.ResourceNamesLength = uint32(len(o.ResourceName))
	}
	if o.AlternateResourceNames != nil && o.AlternateResourceNamesLength == 0 {
		o.AlternateResourceNamesLength = uint16(len(o.AlternateResourceNames))
	}
	if o.ResourceNamesLength > uint32(50) {
		return fmt.Errorf("ResourceNamesLength is out of range")
	}
	if o.AlternateResourceNamesLength > uint16(3) {
		return fmt.Errorf("AlternateResourceNamesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EndpointInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ResourceName != nil || o.ResourceNamesLength > 0 {
		_ptr_resourceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ResourceNamesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ResourceName {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ResourceName[i1] != "" {
					_ptr_resourceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16String(ctx, w, o.ResourceName[i1]); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.ResourceName[i1], _ptr_resourceName); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ResourceName); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResourceName, _ptr_resourceName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ResourceNamesLength); err != nil {
		return err
	}
	if o.AlternateResourceNames != nil || o.AlternateResourceNamesLength > 0 {
		_ptr_alternateResourceNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AlternateResourceNamesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.AlternateResourceNames {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.AlternateResourceNames[i1] != "" {
					_ptr_alternateResourceNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16String(ctx, w, o.AlternateResourceNames[i1]); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.AlternateResourceNames[i1], _ptr_alternateResourceNames); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.AlternateResourceNames); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AlternateResourceNames, _ptr_alternateResourceNames); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AlternateResourceNamesLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *EndpointInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_resourceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ResourceNamesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ResourceNamesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ResourceName", sizeInfo[0])
		}
		o.ResourceName = make([]string, sizeInfo[0])
		for i1 := range o.ResourceName {
			i1 := i1
			_ptr_resourceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16String(ctx, w, &o.ResourceName[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_resourceName := func(ptr interface{}) { o.ResourceName[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.ResourceName[i1], _s_resourceName, _ptr_resourceName); err != nil {
				return err
			}
		}
		return nil
	})
	_s_resourceName := func(ptr interface{}) { o.ResourceName = *ptr.(*[]string) }
	if err := w.ReadPointer(&o.ResourceName, _s_resourceName, _ptr_resourceName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResourceNamesLength); err != nil {
		return err
	}
	_ptr_alternateResourceNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AlternateResourceNamesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AlternateResourceNamesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AlternateResourceNames", sizeInfo[0])
		}
		o.AlternateResourceNames = make([]string, sizeInfo[0])
		for i1 := range o.AlternateResourceNames {
			i1 := i1
			_ptr_alternateResourceNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16String(ctx, w, &o.AlternateResourceNames[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_alternateResourceNames := func(ptr interface{}) { o.AlternateResourceNames[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.AlternateResourceNames[i1], _s_alternateResourceNames, _ptr_alternateResourceNames); err != nil {
				return err
			}
		}
		return nil
	})
	_s_alternateResourceNames := func(ptr interface{}) { o.AlternateResourceNames = *ptr.(*[]string) }
	if err := w.ReadPointer(&o.AlternateResourceNames, _s_alternateResourceNames, _ptr_alternateResourceNames); err != nil {
		return err
	}
	if err := w.ReadData(&o.AlternateResourceNamesLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// PacketHeader structure represents TSG_PACKET_HEADER RPC structure.
//
// The TSG_PACKET_HEADER structure contains information about the ComponentId and PacketId
// fields of the TSG_PACKET structure. The value of PacketId in TSG_PACKET MUST be set
// to TSG_PACKET_TYPE_HEADER.
type PacketHeader struct {
	// ComponentId:  Represents the component sending the packet. This MUST be the following
	// value:
	//
	//	+--------+----------------------+
	//	|        |                      |
	//	| VALUE  |       MEANING        |
	//	|        |                      |
	//	+--------+----------------------+
	//	+--------+----------------------+
	//	| 0x5452 | TS Gateway Transport |
	//	+--------+----------------------+
	ComponentID uint16 `idl:"name:ComponentId" json:"component_id"`
	// PacketId:  Unused.
	//
	// This structure cannot be used by itself as part of any method call. It can be used
	// only in the context of other structures.
	PacketID uint16 `idl:"name:PacketId" json:"packet_id"`
}

func (o *PacketHeader) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketHeader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.ComponentID); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketID); err != nil {
		return err
	}
	return nil
}
func (o *PacketHeader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.ComponentID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketID); err != nil {
		return err
	}
	return nil
}

// CapabilityNap structure represents TSG_CAPABILITY_NAP RPC structure.
//
// The TSG_CAPABILITY_NAP structure contains information about the NAP capabilities
// of the RDG client and RDG server.
//
// This structure MUST be embedded in the TSG_PACKET_CAPABILITIES structure.
type CapabilityNap struct {
	// capabilities:  Indicates the NAP capabilities supported by the RDG client and RDG
	// server. This bit field MUST be 0 or one or more of the following values.
	//
	//	+---------------------------------+
	//	|                                 |
	//	|              VALUE              |
	//	|                                 |
	//	+---------------------------------+
	//	+---------------------------------+
	//	| TSG_NAP_CAPABILITY_QUAR_SOH     |
	//	+---------------------------------+
	//	| TSG_NAP_CAPABILITY_IDLE_TIMEOUT |
	//	+---------------------------------+
	//	| TSG_MESSAGING_CAP_CONSENT_SIGN  |
	//	+---------------------------------+
	//	| TSG_MESSAGING_CAP_SERVICE_MSG   |
	//	+---------------------------------+
	//	| TSG_MESSAGING_CAP_REAUTH        |
	//	+---------------------------------+
	Capabilities uint32 `idl:"name:capabilities" json:"capabilities"`
}

func (o *CapabilityNap) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CapabilityNap) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Capabilities); err != nil {
		return err
	}
	return nil
}
func (o *CapabilityNap) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Capabilities); err != nil {
		return err
	}
	return nil
}

// CapabilitiesUnion structure represents TSG_CAPABILITIES_UNION RPC union.
//
// The TSG_CAPABILITIES_UNION union specifies an RPC switch_type union of structures
// as follows.
type CapabilitiesUnion struct {
	// Types that are assignable to Value
	//
	// *CapabilitiesUnion_CapNap
	Value is_CapabilitiesUnion `json:"value"`
}

func (o *CapabilitiesUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *CapabilitiesUnion_CapNap:
		if value != nil {
			return value.CapNap
		}
	}
	return nil
}

type is_CapabilitiesUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_CapabilitiesUnion()
}

func (o *CapabilitiesUnion) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *CapabilitiesUnion_CapNap:
		return uint32(1)
	}
	return uint32(0)
}

func (o *CapabilitiesUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*CapabilitiesUnion_CapNap)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&CapabilitiesUnion_CapNap{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *CapabilitiesUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &CapabilitiesUnion_CapNap{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// CapabilitiesUnion_CapNap structure represents TSG_CAPABILITIES_UNION RPC union arm.
//
// It has following labels: 1
type CapabilitiesUnion_CapNap struct {
	// TSGCapNap:  A TSG_CAPABILITY_NAP structure.
	CapNap *CapabilityNap `idl:"name:TSGCapNap" json:"cap_nap"`
}

func (*CapabilitiesUnion_CapNap) is_CapabilitiesUnion() {}

func (o *CapabilitiesUnion_CapNap) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CapNap != nil {
		if err := o.CapNap.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CapabilityNap{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *CapabilitiesUnion_CapNap) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.CapNap == nil {
		o.CapNap = &CapabilityNap{}
	}
	if err := o.CapNap.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PacketCapabilities structure represents TSG_PACKET_CAPABILITIES RPC structure.
//
// The TSG_PACKET_CAPABILITIES structure contains information about the capabilities
// of the RDG client and RDG server.
//
// This structure MUST be embedded in the TSG_PACKET_VERSIONCAPS structure.
type PacketCapabilities struct {
	// capabilityType:  Indicates the type of NAP capability supported by the RDG client
	// or the RDG server. This member MUST be the following value:
	//
	//	+------------+----------------------------------------------------------------------------+
	//	|            |                                                                            |
	//	|   VALUE    |                                  MEANING                                   |
	//	|            |                                                                            |
	//	+------------+----------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------+
	//	| 0x00000001 | The RDG server supports NAP capability type (TSG_CAPABILITY_TYPE_NAP).<12> |
	//	+------------+----------------------------------------------------------------------------+
	CapabilityType uint32 `idl:"name:capabilityType" json:"capability_type"`
	// TSGPacket:  Specifies the union containing the actual structure corresponding to
	// the value defined in the capabilityType field. Valid structures are specified in
	// sections 2.2.9.2.1.2.1.1 and 2.2.9.2.1.2.1.2.
	Packet *CapabilitiesUnion `idl:"name:TSGPacket;switch_is:capabilityType" json:"packet"`
}

func (o *PacketCapabilities) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketCapabilities) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.CapabilityType); err != nil {
		return err
	}
	_swPacket := uint32(o.CapabilityType)
	if o.Packet != nil {
		if err := o.Packet.MarshalUnionNDR(ctx, w, _swPacket); err != nil {
			return err
		}
	} else {
		if err := (&CapabilitiesUnion{}).MarshalUnionNDR(ctx, w, _swPacket); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketCapabilities) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.CapabilityType); err != nil {
		return err
	}
	if o.Packet == nil {
		o.Packet = &CapabilitiesUnion{}
	}
	_swPacket := uint32(o.CapabilityType)
	if err := o.Packet.UnmarshalUnionNDR(ctx, w, _swPacket); err != nil {
		return err
	}
	return nil
}

// PacketVersionCaps structure represents TSG_PACKET_VERSIONCAPS RPC structure.
//
// The TSG_PACKET_VERSIONCAPS structure is used for version and capabilities negotiation.
// The value of the packetId field in TSG_PACKET MUST be set to TSG_PACKET_TYPE_VERSIONCAPS.
//
// This structure MUST be embedded in the TSG_PACKET_QUARENC_RESPONSE.
type PacketVersionCaps struct {
	// tsgHeader:  Specified in 2.2.9.2.1.1.
	Header *PacketHeader `idl:"name:tsgHeader" json:"header"`
	// TSGCaps:  An array of TSG_PACKET_CAPABILITIES structures. The number of elements
	// in the array is indicated by the numCapabilities field.
	Caps []*PacketCapabilities `idl:"name:TSGCaps;size_is:(numCapabilities)" json:"caps"`
	// numCapabilities:  The number of array elements for the TSGCaps field. This value
	// MUST be in the range of 0 and 32. If the TSGCaps field is ignored, then this field
	// MUST also be ignored.
	CapabilitiesLength uint32 `idl:"name:numCapabilities" json:"capabilities_length"`
	// majorVersion:  Indicates the major version of the RDG client or RDG server, depending
	// on the sender. This MUST be the following value:
	//
	//	+--------+-------------------------------------------------------------------------+
	//	|        |                                                                         |
	//	| VALUE  |                                 MEANING                                 |
	//	|        |                                                                         |
	//	+--------+-------------------------------------------------------------------------+
	//	+--------+-------------------------------------------------------------------------+
	//	| 0x0001 | Current major version of the Terminal Services Gateway Server Protocol. |
	//	+--------+-------------------------------------------------------------------------+
	MajorVersion uint16 `idl:"name:majorVersion" json:"major_version"`
	// minorVersion:  Indicates the minor version of the RDG client or RDG server, depending
	// on the sender. This MUST be the following value.
	//
	//	+--------+-------------------------------------------------------------------------+
	//	|        |                                                                         |
	//	| VALUE  |                                 MEANING                                 |
	//	|        |                                                                         |
	//	+--------+-------------------------------------------------------------------------+
	//	+--------+-------------------------------------------------------------------------+
	//	| 0x0001 | Current minor version of the Terminal Services Gateway Server Protocol. |
	//	+--------+-------------------------------------------------------------------------+
	MinorVersion uint16 `idl:"name:minorVersion" json:"minor_version"`
	// quarantineCapabilities:  Indicates quarantine capabilities of the RDG client and
	// RDG server, depending on the sender. This MAY be the following value:<11>
	//
	//	+--------+---------------------------------------------------------+
	//	|        |                                                         |
	//	| VALUE  |                         MEANING                         |
	//	|        |                                                         |
	//	+--------+---------------------------------------------------------+
	//	+--------+---------------------------------------------------------+
	//	| 0x0001 | Quarantine is supported and required by the RDG server. |
	//	+--------+---------------------------------------------------------+
	QuarantineCapabilities uint16 `idl:"name:quarantineCapabilities" json:"quarantine_capabilities"`
}

func (o *PacketVersionCaps) xxx_PreparePayload(ctx context.Context) error {
	if o.Caps != nil && o.CapabilitiesLength == 0 {
		o.CapabilitiesLength = uint32(len(o.Caps))
	}
	if o.CapabilitiesLength > uint32(32) {
		return fmt.Errorf("CapabilitiesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketVersionCaps) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Header != nil {
		if err := o.Header.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PacketHeader{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Caps != nil || o.CapabilitiesLength > 0 {
		_ptr_TSGCaps := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.CapabilitiesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Caps {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Caps[i1] != nil {
					if err := o.Caps[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PacketCapabilities{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Caps); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&PacketCapabilities{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Caps, _ptr_TSGCaps); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CapabilitiesLength); err != nil {
		return err
	}
	if err := w.WriteData(o.MajorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.MinorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.QuarantineCapabilities); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *PacketVersionCaps) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Header == nil {
		o.Header = &PacketHeader{}
	}
	if err := o.Header.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_TSGCaps := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.CapabilitiesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.CapabilitiesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Caps", sizeInfo[0])
		}
		o.Caps = make([]*PacketCapabilities, sizeInfo[0])
		for i1 := range o.Caps {
			i1 := i1
			if o.Caps[i1] == nil {
				o.Caps[i1] = &PacketCapabilities{}
			}
			if err := o.Caps[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_TSGCaps := func(ptr interface{}) { o.Caps = *ptr.(*[]*PacketCapabilities) }
	if err := w.ReadPointer(&o.Caps, _s_TSGCaps, _ptr_TSGCaps); err != nil {
		return err
	}
	if err := w.ReadData(&o.CapabilitiesLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.MajorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.QuarantineCapabilities); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// PacketQuarantineConfigRequest structure represents TSG_PACKET_QUARCONFIGREQUEST RPC structure.
//
// The TSG_PACKET_QUARCONFIGREQUEST structure contains information about quarantine
// configuration. RDG server and RDG client MAY support this structure.<13> If the RDG
// server or RDG client do not support the TSG_PACKET_QUARCONFIGREQUEST structure, then
// the error code HRESULT_CODE(E_PROXY_NOTSUPPORTED) is returned.
type PacketQuarantineConfigRequest struct {
	// flags:  Contains information about quarantine configuration.
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *PacketQuarantineConfigRequest) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketQuarantineConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *PacketQuarantineConfigRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// PacketQuarantineRequest structure represents TSG_PACKET_QUARREQUEST RPC structure.
//
// The TSG_PACKET_QUARREQUEST structure<14> contains information about the RDG client's
// statement of health (SoH) and the name of the RDG client machine. The value of the
// packetId field in TSG_PACKET MUST be set to TSG_PACKET_TYPE_QUARREQUEST.
type PacketQuarantineRequest struct {
	// flags:  This field can be any value when sending and ignored on receipt.
	Flags uint32 `idl:"name:flags" json:"flags"`
	// machineName:  A string representing the name of the RDG Client Machine name (section
	// 3.5.1).<15> This field can be ignored. The length of the name, including the terminating
	// null character, MUST be equal to the size specified by the nameLength field.
	MachineName string `idl:"name:machineName;size_is:(nameLength);string" json:"machine_name"`
	// nameLength:  An unsigned long specifying the number of characters in machineName,
	// including the terminating null character. The specified value MUST be in the range
	// from 0 to 513 characters.
	NameLength uint32 `idl:"name:nameLength" json:"name_length"`
	// data:  An array of bytes that specifies the statement of health prepended with nonce,
	// which is obtained in TSG_PACKET_QUARENC_RESPONSE (section 2.2.9.2.1.6) from the RDG
	// server in response to TsProxyCreateTunnel.<16> This field can be ignored. The length
	// of this data is specified by the dataLen field.
	Data []byte `idl:"name:data;size_is:(dataLen);pointer:unique" json:"data"`
	// dataLen:  The length, in bytes, of the data field. This value MUST be in the range
	// between 0 and 8000, both inclusive.
	DataLength uint32 `idl:"name:dataLen" json:"data_length"`
}

func (o *PacketQuarantineRequest) xxx_PreparePayload(ctx context.Context) error {
	if o.MachineName != "" && o.NameLength == 0 {
		o.NameLength = uint32(len(o.MachineName))
	}
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if o.NameLength > uint32(513) {
		return fmt.Errorf("NameLength is out of range")
	}
	if o.DataLength > uint32(8000) {
		return fmt.Errorf("DataLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketQuarantineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.MachineName != "" || o.NameLength > 0 {
		_ptr_machineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NameLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.MachineName)
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
			_MachineName_buf := utf16.Encode([]rune(o.MachineName))
			if uint64(len(_MachineName_buf)) > sizeInfo[0]-1 {
				_MachineName_buf = _MachineName_buf[:sizeInfo[0]-1]
			}
			if o.MachineName != ndr.ZeroString {
				_MachineName_buf = append(_MachineName_buf, uint16(0))
			}
			for i1 := range _MachineName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_MachineName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_MachineName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MachineName, _ptr_machineName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NameLength); err != nil {
		return err
	}
	if o.Data != nil || o.DataLength > 0 {
		_ptr_data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Data {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Data[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_data); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *PacketQuarantineRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_machineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _MachineName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _MachineName_buf", sizeInfo[0])
		}
		_MachineName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _MachineName_buf {
			i1 := i1
			if err := w.ReadData(&_MachineName_buf[i1]); err != nil {
				return err
			}
		}
		o.MachineName = strings.TrimRight(string(utf16.Decode(_MachineName_buf)), ndr.ZeroString)
		return nil
	})
	_s_machineName := func(ptr interface{}) { o.MachineName = *ptr.(*string) }
	if err := w.ReadPointer(&o.MachineName, _s_machineName, _ptr_machineName); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameLength); err != nil {
		return err
	}
	_ptr_data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_data, _ptr_data); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// RedirectionFlags structure represents TSG_REDIRECTION_FLAGS RPC structure.
//
// The TSG_REDIRECTION_FLAGS structure specifies the device redirection settings that
// MUST be enforced by the RDG client. For details about device redirection, see  [MS-RDSOD]
// section 2.1.1.2.
//
// This structure MUST be embedded in the TSG_PACKET_RESPONSE structure.
//
// Note  Both enableAllRedirections and disableAllRedirections MUST NOT be TRUE.
type RedirectionFlags struct {
	// enableAllRedirections:  A Boolean value indicating whether the RDG server specifies
	// any control over the device redirection on the RDG client.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | Device redirection is not enabled for all devices. Other fields of this          |
	//	|                  | structure specify which device redirection is enabled or disabled.               |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | Device redirection is enabled for all devices. All other fields of this          |
	//	|                  | structure MUST be ignored.                                                       |
	//	+------------------+----------------------------------------------------------------------------------+
	EnableAllRedirections bool `idl:"name:enableAllRedirections" json:"enable_all_redirections"`
	// disableAllRedirections:  A Boolean value indicating whether the RDG server specifies
	// any control over disabling all device redirection on the RDG client.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | Device redirection is not disabled for all devices. Other fields of this         |
	//	|                  | structure specify which device redirection is enabled or disabled.               |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | Device redirection is disabled for all devices. All other fields of this         |
	//	|                  | structure MUST be ignored.                                                       |
	//	+------------------+----------------------------------------------------------------------------------+
	DisableAllRedirections bool `idl:"name:disableAllRedirections" json:"disable_all_redirections"`
	// driveRedirectionDisabled:  A Boolean value indicating whether the RDG server specifies
	// any control over disabling drive redirection on the RDG client.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The RDG client is allowed to choose its own redirection settings for enabling or |
	//	|                  | disabling drive redirection.                                                     |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | Drive redirection is disabled.                                                   |
	//	+------------------+----------------------------------------------------------------------------------+
	DriveRedirectionDisabled bool `idl:"name:driveRedirectionDisabled" json:"drive_redirection_disabled"`
	// printerRedirectionDisabled:  A Boolean value indicating whether the RDG server specifies
	// any control over disabling printer redirection on the RDG client.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The RDG client is allowed to choose its own redirection settings for enabling or |
	//	|                  | disabling printer redirection.                                                   |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | Printer redirection is disabled.                                                 |
	//	+------------------+----------------------------------------------------------------------------------+
	PrinterRedirectionDisabled bool `idl:"name:printerRedirectionDisabled" json:"printer_redirection_disabled"`
	// portRedirectionDisabled:  A Boolean value indicating whether the RDG server specifies
	// any control over disabling port redirection on the RDG client.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The RDG client is allowed to choose its own redirection settings for enabling    |
	//	|                  | or disabling port redirection. Port redirection applies to both serial (COM) and |
	//	|                  | parallel ports (LPT).                                                            |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | Port redirection is disabled.                                                    |
	//	+------------------+----------------------------------------------------------------------------------+
	PortRedirectionDisabled bool `idl:"name:portRedirectionDisabled" json:"port_redirection_disabled"`
	// reserved:  Unused. MUST be 0.
	_ bool `idl:"name:reserved"`
	// clipboardRedirectionDisabled:  A Boolean value indicating whether the RDG server
	// specifies any control over disabling clipboard redirection on the RDG client.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The RDG client is allowed to choose its own redirection settings for enabling or |
	//	|                  | disabling clipboard redirection.                                                 |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | Clipboard redirection is disabled.                                               |
	//	+------------------+----------------------------------------------------------------------------------+
	ClipboardRedirectionDisabled bool `idl:"name:clipboardRedirectionDisabled" json:"clipboard_redirection_disabled"`
	// pnpRedirectionDisabled:  A Boolean value indicating whether the RDG server specifies
	// any control over disabling Plug and Play redirection on the RDG client.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The RDG client is allowed to choose its own redirection settings for enabling or |
	//	|                  | disabling PnP redirection.                                                       |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | PnP redirection is disabled.                                                     |
	//	+------------------+----------------------------------------------------------------------------------+
	PnPRedirectionDisabled bool `idl:"name:pnpRedirectionDisabled" json:"pnp_redirection_disabled"`
}

func (o *RedirectionFlags) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RedirectionFlags) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.EnableAllRedirections {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.DisableAllRedirections {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.DriveRedirectionDisabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.PrinterRedirectionDisabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.PortRedirectionDisabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	// reserved reserved
	if err := w.WriteData(false); err != nil {
		return err
	}
	if !o.ClipboardRedirectionDisabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.PnPRedirectionDisabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *RedirectionFlags) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bEnableAllRedirections int32
	if err := w.ReadData(&_bEnableAllRedirections); err != nil {
		return err
	}
	o.EnableAllRedirections = _bEnableAllRedirections != 0
	var _bDisableAllRedirections int32
	if err := w.ReadData(&_bDisableAllRedirections); err != nil {
		return err
	}
	o.DisableAllRedirections = _bDisableAllRedirections != 0
	var _bDriveRedirectionDisabled int32
	if err := w.ReadData(&_bDriveRedirectionDisabled); err != nil {
		return err
	}
	o.DriveRedirectionDisabled = _bDriveRedirectionDisabled != 0
	var _bPrinterRedirectionDisabled int32
	if err := w.ReadData(&_bPrinterRedirectionDisabled); err != nil {
		return err
	}
	o.PrinterRedirectionDisabled = _bPrinterRedirectionDisabled != 0
	var _bPortRedirectionDisabled int32
	if err := w.ReadData(&_bPortRedirectionDisabled); err != nil {
		return err
	}
	o.PortRedirectionDisabled = _bPortRedirectionDisabled != 0
	// reserved reserved
	var _reserved bool
	var _b_reserved int32
	if err := w.ReadData(&_b_reserved); err != nil {
		return err
	}
	_reserved = _b_reserved != 0
	_ = _reserved
	var _bClipboardRedirectionDisabled int32
	if err := w.ReadData(&_bClipboardRedirectionDisabled); err != nil {
		return err
	}
	o.ClipboardRedirectionDisabled = _bClipboardRedirectionDisabled != 0
	var _bPnPRedirectionDisabled int32
	if err := w.ReadData(&_bPnPRedirectionDisabled); err != nil {
		return err
	}
	o.PnPRedirectionDisabled = _bPnPRedirectionDisabled != 0
	return nil
}

// PacketResponse structure represents TSG_PACKET_RESPONSE RPC structure.
//
// The TSG_PACKET_RESPONSE structure contains the response of the RDG server to the
// RDG client for the TsProxyAuthorizeTunnel method call. The value of the packetId
// field in TSG_PACKET MUST be set to TSG_PACKET_TYPE_RESPONSE.
type PacketResponse struct {
	// flags:  The RDG server MUST set this value to TSG_PACKET_TYPE_QUARREQUEST to indicate
	// that this structure is in response to the TsProxyAuthorizeTunnel method call. The
	// RDG client MAY ignore this field.
	Flags uint32 `idl:"name:flags" json:"flags"`
	// reserved:  This field is unused and can be any value when sending and ignored on
	// receipt.
	_ uint32 `idl:"name:reserved"`
	// responseData:  Byte data representing the response from the RDG server for the TsProxyAuthorizeTunnel
	// method call. If the Negotiated Capabilities ADM element contains TSG_NAP_CAPABILITY_QUAR_SOH
	// and TSG_NAP_CAPABILITY_IDLE_TIMEOUT and the value of the dataLen member specified
	// in the TSG_PACKET_QUARREQUEST structure (section 2.2.9.2.1.4) is greater than zero,
	// then responseData MUST contain both the statement of health response (SoHR) and the
	// idle timeout value. If Negotiated Capabilities contains only TSG_NAP_CAPABILITY_QUAR_SOH
	// and the value of the dataLen member specified in the TSG_PACKET_QUARREQUEST structure
	// (section 2.2.9.2.1.4) is greater than zero, then responseData MUST contain only the
	// statement of health response. If Negotiated Capabilities contains only TSG_NAP_CAPABILITY_IDLE_TIMEOUT,
	// then responseData MUST contain only the idle timeout value. The length of the data
	// MUST be equal to that specified by responseDataLen. If Negotiated Capabilities does
	// not contain both TSG_NAP_CAPABILITY_QUAR_SOH and TSG_NAP_CAPABILITY_IDLE_TIMEOUT,
	// then responseData is ignored and responseDataLen is set to zero.<17>
	ResponseData []byte `idl:"name:responseData;size_is:(responseDataLen)" json:"response_data"`
	// responseDataLen:  Length, in bytes, of the data specified by the responseData field.
	ResponseDataLength uint32 `idl:"name:responseDataLen" json:"response_data_length"`
	// redirectionFlags:  A TSG_REDIRECTION_FLAGS structure.<18>
	RedirectionFlags *RedirectionFlags `idl:"name:redirectionFlags" json:"redirection_flags"`
}

func (o *PacketResponse) xxx_PreparePayload(ctx context.Context) error {
	if o.ResponseData != nil && o.ResponseDataLength == 0 {
		o.ResponseDataLength = uint32(len(o.ResponseData))
	}
	if o.ResponseDataLength > uint32(24000) {
		return fmt.Errorf("ResponseDataLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	// reserved reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.ResponseData != nil || o.ResponseDataLength > 0 {
		_ptr_responseData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ResponseDataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ResponseData {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.ResponseData[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.ResponseData); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResponseData, _ptr_responseData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ResponseDataLength); err != nil {
		return err
	}
	if o.RedirectionFlags != nil {
		if err := o.RedirectionFlags.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RedirectionFlags{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *PacketResponse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	// reserved reserved
	var _reserved uint32
	if err := w.ReadData(&_reserved); err != nil {
		return err
	}
	_ptr_responseData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ResponseDataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ResponseDataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ResponseData", sizeInfo[0])
		}
		o.ResponseData = make([]byte, sizeInfo[0])
		for i1 := range o.ResponseData {
			i1 := i1
			if err := w.ReadData(&o.ResponseData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_responseData := func(ptr interface{}) { o.ResponseData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ResponseData, _s_responseData, _ptr_responseData); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResponseDataLength); err != nil {
		return err
	}
	if o.RedirectionFlags == nil {
		o.RedirectionFlags = &RedirectionFlags{}
	}
	if err := o.RedirectionFlags.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// PacketQuarantineEncResponse structure represents TSG_PACKET_QUARENC_RESPONSE RPC structure.
//
// The TSG_PACKET_QUARENC_RESPONSE structure contains the response of the RDG server
// for the TsProxyCreateTunnel method call. The value of the packetId field in TSG_PACKET
// MUST be set to TSG_PACKET_TYPE_QUARENC_RESPONSE.
type PacketQuarantineEncResponse struct {
	// flags:  Unused. MUST be 0.
	Flags uint32 `idl:"name:flags" json:"flags"`
	// certChainLen:  An unsigned long specifying the number of characters in certChainData,
	// including the terminating null character. If the quarantineCapabilities field of
	// the TSG_PACKET_VERSIONCAPS structure is set to 1, this MUST be a nonzero value. This
	// field MUST be ignored if certChainData is ignored. The value MUST be in the range
	// of 0 and 24000; both inclusive.
	CertChainLength uint32 `idl:"name:certChainLen" json:"cert_chain_length"`
	// certChainData:  The certificate, along with the chain, that the RDG server used for
	// the SCHANNEL authentication service as part of registering the RPC interfaces and
	// initialization. It MUST be a string representation of the certificate chain if certChainLen
	// is nonzero.<19> This field can be ignored.
	CertChainData string `idl:"name:certChainData;size_is:(certChainLen);string" json:"cert_chain_data"`
	// nonce:  A GUID to uniquely identify this connection to prevent replay attacks by
	// the RDG client. This can be used for auditing purposes. A GUID is a unique ID using
	// opaque sequence of bytes as specified in [MS-DTYP] section 2.3.4.2.
	Nonce *dtyp.GUID `idl:"name:nonce" json:"nonce"`
	// versionCaps:  A pointer to a TSG_PACKET_VERSIONCAPS structure, as specified in section
	// 2.2.9.2.1.2.
	VersionCaps *PacketVersionCaps `idl:"name:versionCaps" json:"version_caps"`
}

func (o *PacketQuarantineEncResponse) xxx_PreparePayload(ctx context.Context) error {
	if o.CertChainData != "" && o.CertChainLength == 0 {
		o.CertChainLength = uint32(len(o.CertChainData))
	}
	if o.CertChainLength > uint32(24000) {
		return fmt.Errorf("CertChainLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketQuarantineEncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.CertChainLength); err != nil {
		return err
	}
	if o.CertChainData != "" || o.CertChainLength > 0 {
		_ptr_certChainData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.CertChainLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.CertChainData)
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
			_CertChainData_buf := utf16.Encode([]rune(o.CertChainData))
			if uint64(len(_CertChainData_buf)) > sizeInfo[0]-1 {
				_CertChainData_buf = _CertChainData_buf[:sizeInfo[0]-1]
			}
			if o.CertChainData != ndr.ZeroString {
				_CertChainData_buf = append(_CertChainData_buf, uint16(0))
			}
			for i1 := range _CertChainData_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_CertChainData_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_CertChainData_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CertChainData, _ptr_certChainData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Nonce != nil {
		if err := o.Nonce.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.VersionCaps != nil {
		_ptr_versionCaps := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.VersionCaps != nil {
				if err := o.VersionCaps.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketVersionCaps{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VersionCaps, _ptr_versionCaps); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketQuarantineEncResponse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.CertChainLength); err != nil {
		return err
	}
	_ptr_certChainData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _CertChainData_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _CertChainData_buf", sizeInfo[0])
		}
		_CertChainData_buf = make([]uint16, sizeInfo[0])
		for i1 := range _CertChainData_buf {
			i1 := i1
			if err := w.ReadData(&_CertChainData_buf[i1]); err != nil {
				return err
			}
		}
		o.CertChainData = strings.TrimRight(string(utf16.Decode(_CertChainData_buf)), ndr.ZeroString)
		return nil
	})
	_s_certChainData := func(ptr interface{}) { o.CertChainData = *ptr.(*string) }
	if err := w.ReadPointer(&o.CertChainData, _s_certChainData, _ptr_certChainData); err != nil {
		return err
	}
	if o.Nonce == nil {
		o.Nonce = &dtyp.GUID{}
	}
	if err := o.Nonce.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_versionCaps := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.VersionCaps == nil {
			o.VersionCaps = &PacketVersionCaps{}
		}
		if err := o.VersionCaps.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_versionCaps := func(ptr interface{}) { o.VersionCaps = *ptr.(**PacketVersionCaps) }
	if err := w.ReadPointer(&o.VersionCaps, _s_versionCaps, _ptr_versionCaps); err != nil {
		return err
	}
	return nil
}

// PacketMessageRequest structure represents TSG_PACKET_MSG_REQUEST RPC structure.
//
// The TSG_PACKET_MSG_REQUEST structure contains the request from the client to the
// RDG server to send across an administrative message whenever there is any. The value
// of the packetId field in TSG_PACKET MUST be set to TSG_PACKET_TYPE_MSGREQUEST_PACKET.
type PacketMessageRequest struct {
	// maxMessagesPerBatch:  An unsigned long that specifies how many messages can be sent
	// by the server at one time.
	MaxMessagesPerBatch uint32 `idl:"name:maxMessagesPerBatch" json:"max_messages_per_batch"`
}

func (o *PacketMessageRequest) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxMessagesPerBatch); err != nil {
		return err
	}
	return nil
}
func (o *PacketMessageRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxMessagesPerBatch); err != nil {
		return err
	}
	return nil
}

// PacketStringMessage structure represents TSG_PACKET_STRING_MESSAGE RPC structure.
//
// The TSG_PACKET_STRING_MESSAGE structure contains either the Consent Signing Message
// or the Administrative Message that is being sent from the RDG server to the client.
type PacketStringMessage struct {
	// isDisplayMandatory:  A Boolean that specifies whether the client needs to display
	// this message.
	IsDisplayMandatory int32 `idl:"name:isDisplayMandatory" json:"is_display_mandatory"`
	// isConsentMandatory:  A Boolean that specifies whether the user needs to give its
	// consent before the connection can proceed.
	IsConsentMandatory int32 `idl:"name:isConsentMandatory" json:"is_consent_mandatory"`
	// msgBytes:  An unsigned long specifying the number of characters in msgBuffer, including
	// the terminating null character. The size of the message SHOULD<21> be determined
	// by the serverCert field in the HTTP_TUNNEL_RESPONSE_OPTIONAL structure (section 2.2.10.21).
	// The consent message is embedded in the HTTP_TUNNEL_RESPONSE as part of the HTTP_TUNNEL_RESPONSE_OPTIONAL
	// structure. When the HTTP_TUNNEL_RESPONSE_FIELD_CONSENT_MSG flag is set in the HTTP_TUNNEL_RESPONSE_FIELDS_PRESENT_FLAGS
	// (section 2.2.5.3.8), the HTTP_TUNNEL_RESPONSE_OPTIONAL data structure contains a
	// consent message in the HTTP_UNICODE_STRING format (section 2.2.10.22).
	MessageBytes uint32 `idl:"name:msgBytes" json:"message_bytes"`
	// msgBuffer:  An array of wchar_t specifying the string. The size of the buffer is
	// as indicated by msgBytes.
	MessageBuffer string `idl:"name:msgBuffer;size_is:(msgBytes)" json:"message_buffer"`
}

func (o *PacketStringMessage) xxx_PreparePayload(ctx context.Context) error {
	if o.MessageBuffer != "" && o.MessageBytes == 0 {
		o.MessageBytes = uint32(len(o.MessageBuffer))
	}
	if o.MessageBytes > uint32(65536) {
		return fmt.Errorf("MessageBytes is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketStringMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.IsDisplayMandatory); err != nil {
		return err
	}
	if err := w.WriteData(o.IsConsentMandatory); err != nil {
		return err
	}
	if err := w.WriteData(o.MessageBytes); err != nil {
		return err
	}
	if o.MessageBuffer != "" || o.MessageBytes > 0 {
		_ptr_msgBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MessageBytes)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_MessageBuffer_buf := utf16.Encode([]rune(o.MessageBuffer))
			if uint64(len(_MessageBuffer_buf)) > sizeInfo[0] {
				_MessageBuffer_buf = _MessageBuffer_buf[:sizeInfo[0]]
			}
			for i1 := range _MessageBuffer_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_MessageBuffer_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_MessageBuffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MessageBuffer, _ptr_msgBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketStringMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsDisplayMandatory); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsConsentMandatory); err != nil {
		return err
	}
	if err := w.ReadData(&o.MessageBytes); err != nil {
		return err
	}
	_ptr_msgBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.MessageBytes > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.MessageBytes)
		}
		var _MessageBuffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _MessageBuffer_buf", sizeInfo[0])
		}
		_MessageBuffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _MessageBuffer_buf {
			i1 := i1
			if err := w.ReadData(&_MessageBuffer_buf[i1]); err != nil {
				return err
			}
		}
		o.MessageBuffer = strings.TrimRight(string(utf16.Decode(_MessageBuffer_buf)), ndr.ZeroString)
		return nil
	})
	_s_msgBuffer := func(ptr interface{}) { o.MessageBuffer = *ptr.(*string) }
	if err := w.ReadPointer(&o.MessageBuffer, _s_msgBuffer, _ptr_msgBuffer); err != nil {
		return err
	}
	return nil
}

// PacketReauthMessage structure represents TSG_PACKET_REAUTH_MESSAGE RPC structure.
//
// The TSG_PACKET_REAUTH_MESSAGE structure is sent by the RDG server to the client when
// the server requires the user credential to be reauthenticated.
type PacketReauthMessage struct {
	// tunnelContext:  A unsigned __int64 that is sent by the server to client. When the
	// client initiates the reauthentication sequence, it MUST include this context. This
	// is used by the server to validate successful reauthentication by the client.
	TunnelContext uint64 `idl:"name:tunnelContext" json:"tunnel_context"`
}

func (o *PacketReauthMessage) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketReauthMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.TunnelContext); err != nil {
		return err
	}
	return nil
}
func (o *PacketReauthMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.TunnelContext); err != nil {
		return err
	}
	return nil
}

// PacketTypeMessageUnion structure represents TSG_PACKET_TYPE_MESSAGE_UNION RPC union.
type PacketTypeMessageUnion struct {
	// Types that are assignable to Value
	//
	// *PacketTypeMessageUnion_ConsentMessage
	// *PacketTypeMessageUnion_ServiceMessage
	// *PacketTypeMessageUnion_ReauthMessage
	Value is_PacketTypeMessageUnion `json:"value"`
}

func (o *PacketTypeMessageUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *PacketTypeMessageUnion_ConsentMessage:
		if value != nil {
			return value.ConsentMessage
		}
	case *PacketTypeMessageUnion_ServiceMessage:
		if value != nil {
			return value.ServiceMessage
		}
	case *PacketTypeMessageUnion_ReauthMessage:
		if value != nil {
			return value.ReauthMessage
		}
	}
	return nil
}

type is_PacketTypeMessageUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_PacketTypeMessageUnion()
}

func (o *PacketTypeMessageUnion) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *PacketTypeMessageUnion_ConsentMessage:
		return uint32(1)
	case *PacketTypeMessageUnion_ServiceMessage:
		return uint32(2)
	case *PacketTypeMessageUnion_ReauthMessage:
		return uint32(3)
	}
	return uint32(0)
}

func (o *PacketTypeMessageUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
		_o, _ := o.Value.(*PacketTypeMessageUnion_ConsentMessage)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeMessageUnion_ConsentMessage{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*PacketTypeMessageUnion_ServiceMessage)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeMessageUnion_ServiceMessage{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*PacketTypeMessageUnion_ReauthMessage)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeMessageUnion_ReauthMessage{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *PacketTypeMessageUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
		o.Value = &PacketTypeMessageUnion_ConsentMessage{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &PacketTypeMessageUnion_ServiceMessage{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &PacketTypeMessageUnion_ReauthMessage{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// PacketTypeMessageUnion_ConsentMessage structure represents TSG_PACKET_TYPE_MESSAGE_UNION RPC union arm.
//
// It has following labels: 1
type PacketTypeMessageUnion_ConsentMessage struct {
	ConsentMessage *PacketStringMessage `idl:"name:consentMessage" json:"consent_message"`
}

func (*PacketTypeMessageUnion_ConsentMessage) is_PacketTypeMessageUnion() {}

func (o *PacketTypeMessageUnion_ConsentMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ConsentMessage != nil {
		_ptr_consentMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ConsentMessage != nil {
				if err := o.ConsentMessage.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketStringMessage{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ConsentMessage, _ptr_consentMessage); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeMessageUnion_ConsentMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_consentMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ConsentMessage == nil {
			o.ConsentMessage = &PacketStringMessage{}
		}
		if err := o.ConsentMessage.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_consentMessage := func(ptr interface{}) { o.ConsentMessage = *ptr.(**PacketStringMessage) }
	if err := w.ReadPointer(&o.ConsentMessage, _s_consentMessage, _ptr_consentMessage); err != nil {
		return err
	}
	return nil
}

// PacketTypeMessageUnion_ServiceMessage structure represents TSG_PACKET_TYPE_MESSAGE_UNION RPC union arm.
//
// It has following labels: 2
type PacketTypeMessageUnion_ServiceMessage struct {
	ServiceMessage *PacketStringMessage `idl:"name:serviceMessage" json:"service_message"`
}

func (*PacketTypeMessageUnion_ServiceMessage) is_PacketTypeMessageUnion() {}

func (o *PacketTypeMessageUnion_ServiceMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ServiceMessage != nil {
		_ptr_serviceMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ServiceMessage != nil {
				if err := o.ServiceMessage.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketStringMessage{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ServiceMessage, _ptr_serviceMessage); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeMessageUnion_ServiceMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_serviceMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ServiceMessage == nil {
			o.ServiceMessage = &PacketStringMessage{}
		}
		if err := o.ServiceMessage.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_serviceMessage := func(ptr interface{}) { o.ServiceMessage = *ptr.(**PacketStringMessage) }
	if err := w.ReadPointer(&o.ServiceMessage, _s_serviceMessage, _ptr_serviceMessage); err != nil {
		return err
	}
	return nil
}

// PacketTypeMessageUnion_ReauthMessage structure represents TSG_PACKET_TYPE_MESSAGE_UNION RPC union arm.
//
// It has following labels: 3
type PacketTypeMessageUnion_ReauthMessage struct {
	ReauthMessage *PacketReauthMessage `idl:"name:reauthMessage" json:"reauth_message"`
}

func (*PacketTypeMessageUnion_ReauthMessage) is_PacketTypeMessageUnion() {}

func (o *PacketTypeMessageUnion_ReauthMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ReauthMessage != nil {
		_ptr_reauthMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ReauthMessage != nil {
				if err := o.ReauthMessage.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketReauthMessage{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReauthMessage, _ptr_reauthMessage); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeMessageUnion_ReauthMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_reauthMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ReauthMessage == nil {
			o.ReauthMessage = &PacketReauthMessage{}
		}
		if err := o.ReauthMessage.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_reauthMessage := func(ptr interface{}) { o.ReauthMessage = *ptr.(**PacketReauthMessage) }
	if err := w.ReadPointer(&o.ReauthMessage, _s_reauthMessage, _ptr_reauthMessage); err != nil {
		return err
	}
	return nil
}

// PacketMessageResponse structure represents TSG_PACKET_MSG_RESPONSE RPC structure.
//
// The TSG_PACKET_MSG_RESPONSE structure contains the response of the RDG server to
// the client when a message needs to be sent to the client. The value of the packetId
// field in TSG_PACKET MUST be set to TSG_PACKET_TYPE_MESSAGE_PACKET.
type PacketMessageResponse struct {
	// msgID:  This field is unused.<20> This field can be ignored.
	MessageID uint32 `idl:"name:msgID" json:"message_id"`
	// msgType:  An unsigned long specifying what type of message is being sent by the server.
	// This MUST be one of the following values.
	//
	//	+----------------------------------------------+--------------------------------------------------+
	//	|                                              |                                                  |
	//	|                    VALUE                     |                     MEANING                      |
	//	|                                              |                                                  |
	//	+----------------------------------------------+--------------------------------------------------+
	//	+----------------------------------------------+--------------------------------------------------+
	//	| TSG_ASYNC_MESSAGE_CONSENT_MESSAGE 0x00000001 | The server is sending a Consent Signing Message. |
	//	+----------------------------------------------+--------------------------------------------------+
	//	| TSG_ASYNC_MESSAGE_SERVICE_MESSAGE 0x00000002 | The server is sending an Administrative Message. |
	//	+----------------------------------------------+--------------------------------------------------+
	//	| TSG_ASYNC_MESSAGE_REAUTH 0x00000003          | The server expects the client to Reauthenticate. |
	//	+----------------------------------------------+--------------------------------------------------+
	MessageType uint32 `idl:"name:msgType" json:"message_type"`
	// isMsgPresent:  A Boolean that indicates whether the messagePacket parameter is present
	// or not. If the value is TRUE, then messagePacket contains valid data and can be processed.
	// If the value is FALSE, messagePacket parameter MUST be ignored.
	IsMessagePresent int32 `idl:"name:isMsgPresent" json:"is_message_present"`
	// messagePacket:  A TSG_PACKET_TYPE_MESSAGE_UNION union, as specified in section 2.2.9.2.1.9.1.
	MessagePacket *PacketTypeMessageUnion `idl:"name:messagePacket;switch_is:msgType" json:"message_packet"`
}

func (o *PacketMessageResponse) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MessageID); err != nil {
		return err
	}
	if err := w.WriteData(o.MessageType); err != nil {
		return err
	}
	if err := w.WriteData(o.IsMessagePresent); err != nil {
		return err
	}
	_swMessagePacket := uint32(o.MessageType)
	if o.MessagePacket != nil {
		if err := o.MessagePacket.MarshalUnionNDR(ctx, w, _swMessagePacket); err != nil {
			return err
		}
	} else {
		if err := (&PacketTypeMessageUnion{}).MarshalUnionNDR(ctx, w, _swMessagePacket); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketMessageResponse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MessageID); err != nil {
		return err
	}
	if err := w.ReadData(&o.MessageType); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsMessagePresent); err != nil {
		return err
	}
	if o.MessagePacket == nil {
		o.MessagePacket = &PacketTypeMessageUnion{}
	}
	_swMessagePacket := uint32(o.MessageType)
	if err := o.MessagePacket.UnmarshalUnionNDR(ctx, w, _swMessagePacket); err != nil {
		return err
	}
	return nil
}

// PacketCapsResponse structure represents TSG_PACKET_CAPS_RESPONSE RPC structure.
//
// The TSG_PACKET_CAPS_RESPONSE structure contains the response of the RDG server, which
// supports Consent Signing capability, to the RDG client for the TsProxyCreateTunnel
// method call. This structure contains TSG_PACKET_QUARENC_RESPONSE followed by the
// consent signing string. The value of the packetId field in TSG_PACKET MUST be set
// to TSG_PACKET_TYPE_CAPS_RESPONSE.
type PacketCapsResponse struct {
	// pktQuarEncResponse:  A TSG_PACKET_QUARENC_RESPONSE structure as specified in section
	// 2.2.9.2.1.6.
	PacketQuarantineEncResponse *PacketQuarantineEncResponse `idl:"name:pktQuarEncResponse" json:"packet_quarantine_enc_response"`
	// pktConsentMessage:  A TSG_PACKET_MSG_RESPONSE structure as specified in section 2.2.9.2.1.9.
	PacketConsentMessage *PacketMessageResponse `idl:"name:pktConsentMessage" json:"packet_consent_message"`
}

func (o *PacketCapsResponse) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketCapsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.PacketQuarantineEncResponse != nil {
		if err := o.PacketQuarantineEncResponse.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PacketQuarantineEncResponse{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PacketConsentMessage != nil {
		if err := o.PacketConsentMessage.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PacketMessageResponse{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketCapsResponse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.PacketQuarantineEncResponse == nil {
		o.PacketQuarantineEncResponse = &PacketQuarantineEncResponse{}
	}
	if err := o.PacketQuarantineEncResponse.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PacketConsentMessage == nil {
		o.PacketConsentMessage = &PacketMessageResponse{}
	}
	if err := o.PacketConsentMessage.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PacketAuth structure represents TSG_PACKET_AUTH RPC structure.
//
// The TSG_PACKET_AUTH structure is sent by the client to the TS Gateway server when
// Pluggable Authentication is used. This packet includes TSG_PACKET_VERSIONCAPS, which
// is used for capability negotiation, and cookie, which is used for user authentication.
// This MUST be the first packet from the client to the server if the server has Pluggable
// Authentication turned on. The value of the packetId field in TSG_PACKET MUST be set
// to TSG_PACKET_TYPE_AUTH.
type PacketAuth struct {
	// TSGVersionCaps:  A TSG_PACKET_VERSIONCAPS structure as specified in section 2.2.9.2.1.2.
	VersionCaps *PacketVersionCaps `idl:"name:TSGVersionCaps" json:"version_caps"`
	// cookieLen:  An unsigned long that specifies the size in bytes for the field cookie.
	CookieLength uint32 `idl:"name:cookieLen" json:"cookie_length"`
	// cookie:  A byte pointer that points to the cookie data. The cookie is used for authentication.
	Cookie []byte `idl:"name:cookie;size_is:(cookieLen)" json:"cookie"`
}

func (o *PacketAuth) xxx_PreparePayload(ctx context.Context) error {
	if o.Cookie != nil && o.CookieLength == 0 {
		o.CookieLength = uint32(len(o.Cookie))
	}
	if o.CookieLength > uint32(65536) {
		return fmt.Errorf("CookieLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketAuth) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.VersionCaps != nil {
		if err := o.VersionCaps.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PacketVersionCaps{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CookieLength); err != nil {
		return err
	}
	if o.Cookie != nil || o.CookieLength > 0 {
		_ptr_cookie := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			return nil
		})
		if err := w.WritePointer(&o.Cookie, _ptr_cookie); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketAuth) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.VersionCaps == nil {
		o.VersionCaps = &PacketVersionCaps{}
	}
	if err := o.VersionCaps.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.CookieLength); err != nil {
		return err
	}
	_ptr_cookie := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.CookieLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.CookieLength)
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
		return nil
	})
	_s_cookie := func(ptr interface{}) { o.Cookie = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Cookie, _s_cookie, _ptr_cookie); err != nil {
		return err
	}
	return nil
}

// InitialPacketTypeUnion structure represents TSG_INITIAL_PACKET_TYPE_UNION RPC union.
type InitialPacketTypeUnion struct {
	// Types that are assignable to Value
	//
	// *InitialPacketTypeUnion_PacketVersionCaps
	// *InitialPacketTypeUnion_PacketAuth
	Value is_InitialPacketTypeUnion `json:"value"`
}

func (o *InitialPacketTypeUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *InitialPacketTypeUnion_PacketVersionCaps:
		if value != nil {
			return value.PacketVersionCaps
		}
	case *InitialPacketTypeUnion_PacketAuth:
		if value != nil {
			return value.PacketAuth
		}
	}
	return nil
}

type is_InitialPacketTypeUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_InitialPacketTypeUnion()
}

func (o *InitialPacketTypeUnion) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *InitialPacketTypeUnion_PacketVersionCaps:
		return uint32(22083)
	case *InitialPacketTypeUnion_PacketAuth:
		return uint32(16468)
	}
	return uint32(0)
}

func (o *InitialPacketTypeUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
	case uint32(22083):
		_o, _ := o.Value.(*InitialPacketTypeUnion_PacketVersionCaps)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InitialPacketTypeUnion_PacketVersionCaps{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16468):
		_o, _ := o.Value.(*InitialPacketTypeUnion_PacketAuth)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InitialPacketTypeUnion_PacketAuth{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *InitialPacketTypeUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
	case uint32(22083):
		o.Value = &InitialPacketTypeUnion_PacketVersionCaps{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16468):
		o.Value = &InitialPacketTypeUnion_PacketAuth{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// InitialPacketTypeUnion_PacketVersionCaps structure represents TSG_INITIAL_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 22083
type InitialPacketTypeUnion_PacketVersionCaps struct {
	PacketVersionCaps *PacketVersionCaps `idl:"name:packetVersionCaps" json:"packet_version_caps"`
}

func (*InitialPacketTypeUnion_PacketVersionCaps) is_InitialPacketTypeUnion() {}

func (o *InitialPacketTypeUnion_PacketVersionCaps) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketVersionCaps != nil {
		_ptr_packetVersionCaps := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketVersionCaps != nil {
				if err := o.PacketVersionCaps.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketVersionCaps{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketVersionCaps, _ptr_packetVersionCaps); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InitialPacketTypeUnion_PacketVersionCaps) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetVersionCaps := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketVersionCaps == nil {
			o.PacketVersionCaps = &PacketVersionCaps{}
		}
		if err := o.PacketVersionCaps.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetVersionCaps := func(ptr interface{}) { o.PacketVersionCaps = *ptr.(**PacketVersionCaps) }
	if err := w.ReadPointer(&o.PacketVersionCaps, _s_packetVersionCaps, _ptr_packetVersionCaps); err != nil {
		return err
	}
	return nil
}

// InitialPacketTypeUnion_PacketAuth structure represents TSG_INITIAL_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 16468
type InitialPacketTypeUnion_PacketAuth struct {
	PacketAuth *PacketAuth `idl:"name:packetAuth" json:"packet_auth"`
}

func (*InitialPacketTypeUnion_PacketAuth) is_InitialPacketTypeUnion() {}

func (o *InitialPacketTypeUnion_PacketAuth) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketAuth != nil {
		_ptr_packetAuth := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketAuth != nil {
				if err := o.PacketAuth.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketAuth{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketAuth, _ptr_packetAuth); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InitialPacketTypeUnion_PacketAuth) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetAuth := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketAuth == nil {
			o.PacketAuth = &PacketAuth{}
		}
		if err := o.PacketAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetAuth := func(ptr interface{}) { o.PacketAuth = *ptr.(**PacketAuth) }
	if err := w.ReadPointer(&o.PacketAuth, _s_packetAuth, _ptr_packetAuth); err != nil {
		return err
	}
	return nil
}

// PacketReauth structure represents TSG_PACKET_REAUTH RPC structure.
//
// The TSG_PACKET_REAUTH structure is sent by the client to the TS Gateway server when
// the client is reauthenticating the connection. The value of the packetId field in
// TSG_PACKET MUST be set to TSG_PACKET_TYPE_REAUTH.
type PacketReauth struct {
	// tunnelContext:  An unsigned __int64 that identifies which tunnel is being reauthenticated.
	TunnelContext uint64 `idl:"name:tunnelContext" json:"tunnel_context"`
	// packetId:  An unsigned long that specifies what type of packet is present inside
	// TSGInitialPacket.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TSG_PACKET_TYPE_VERSIONCAPS 0x00005643 | This packet is sent when Pluggable Authentication is off.                        |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| TSG_PACKET_TYPE_AUTH 0x00004054        | This packet is sent when Pluggable Authentication is on. This packet             |
	//	|                                        | includes TSG_PACKET_VERSIONCAPS as well as the cookie that is required for       |
	//	|                                        | authentication.                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	PacketID uint32 `idl:"name:packetId" json:"packet_id"`
	// TSGInitialPacket:  A TSG_INITIAL_PACKET_TYPE_UNION union as specified in section
	// 2.2.9.2.1.11.1.
	InitialPacket *InitialPacketTypeUnion `idl:"name:TSGInitialPacket;switch_is:packetId" json:"initial_packet"`
}

func (o *PacketReauth) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketReauth) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.TunnelContext); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketID); err != nil {
		return err
	}
	_swInitialPacket := uint32(o.PacketID)
	if o.InitialPacket != nil {
		if err := o.InitialPacket.MarshalUnionNDR(ctx, w, _swInitialPacket); err != nil {
			return err
		}
	} else {
		if err := (&InitialPacketTypeUnion{}).MarshalUnionNDR(ctx, w, _swInitialPacket); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *PacketReauth) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.TunnelContext); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketID); err != nil {
		return err
	}
	if o.InitialPacket == nil {
		o.InitialPacket = &InitialPacketTypeUnion{}
	}
	_swInitialPacket := uint32(o.PacketID)
	if err := o.InitialPacket.UnmarshalUnionNDR(ctx, w, _swInitialPacket); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion structure represents TSG_PACKET_TYPE_UNION RPC union.
//
// The TSG_PACKET_TYPE_UNION union specifies an RPC switch_type union of structures
// as follows.
type PacketTypeUnion struct {
	// Types that are assignable to Value
	//
	// *PacketTypeUnion_PacketHeader
	// *PacketTypeUnion_PacketVersionCaps
	// *PacketTypeUnion_PacketQuarantineConfigRequest
	// *PacketTypeUnion_PacketQuarantineRequest
	// *PacketTypeUnion_PacketResponse
	// *PacketTypeUnion_PacketQuarantineEncResponse
	// *PacketTypeUnion_PacketCapsResponse
	// *PacketTypeUnion_PacketMessageRequest
	// *PacketTypeUnion_PacketMessageResponse
	// *PacketTypeUnion_PacketAuth
	// *PacketTypeUnion_PacketReauth
	Value is_PacketTypeUnion `json:"value"`
}

func (o *PacketTypeUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *PacketTypeUnion_PacketHeader:
		if value != nil {
			return value.PacketHeader
		}
	case *PacketTypeUnion_PacketVersionCaps:
		if value != nil {
			return value.PacketVersionCaps
		}
	case *PacketTypeUnion_PacketQuarantineConfigRequest:
		if value != nil {
			return value.PacketQuarantineConfigRequest
		}
	case *PacketTypeUnion_PacketQuarantineRequest:
		if value != nil {
			return value.PacketQuarantineRequest
		}
	case *PacketTypeUnion_PacketResponse:
		if value != nil {
			return value.PacketResponse
		}
	case *PacketTypeUnion_PacketQuarantineEncResponse:
		if value != nil {
			return value.PacketQuarantineEncResponse
		}
	case *PacketTypeUnion_PacketCapsResponse:
		if value != nil {
			return value.PacketCapsResponse
		}
	case *PacketTypeUnion_PacketMessageRequest:
		if value != nil {
			return value.PacketMessageRequest
		}
	case *PacketTypeUnion_PacketMessageResponse:
		if value != nil {
			return value.PacketMessageResponse
		}
	case *PacketTypeUnion_PacketAuth:
		if value != nil {
			return value.PacketAuth
		}
	case *PacketTypeUnion_PacketReauth:
		if value != nil {
			return value.PacketReauth
		}
	}
	return nil
}

type is_PacketTypeUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_PacketTypeUnion()
}

func (o *PacketTypeUnion) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *PacketTypeUnion_PacketHeader:
		return uint32(18500)
	case *PacketTypeUnion_PacketVersionCaps:
		return uint32(22083)
	case *PacketTypeUnion_PacketQuarantineConfigRequest:
		return uint32(20803)
	case *PacketTypeUnion_PacketQuarantineRequest:
		return uint32(20818)
	case *PacketTypeUnion_PacketResponse:
		return uint32(20562)
	case *PacketTypeUnion_PacketQuarantineEncResponse:
		return uint32(17746)
	case *PacketTypeUnion_PacketCapsResponse:
		return uint32(17232)
	case *PacketTypeUnion_PacketMessageRequest:
		return uint32(18258)
	case *PacketTypeUnion_PacketMessageResponse:
		return uint32(18256)
	case *PacketTypeUnion_PacketAuth:
		return uint32(16468)
	case *PacketTypeUnion_PacketReauth:
		return uint32(21072)
	}
	return uint32(0)
}

func (o *PacketTypeUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
	case uint32(18500):
		_o, _ := o.Value.(*PacketTypeUnion_PacketHeader)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketHeader{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(22083):
		_o, _ := o.Value.(*PacketTypeUnion_PacketVersionCaps)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketVersionCaps{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(20803):
		_o, _ := o.Value.(*PacketTypeUnion_PacketQuarantineConfigRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketQuarantineConfigRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(20818):
		_o, _ := o.Value.(*PacketTypeUnion_PacketQuarantineRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketQuarantineRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(20562):
		_o, _ := o.Value.(*PacketTypeUnion_PacketResponse)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketResponse{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(17746):
		_o, _ := o.Value.(*PacketTypeUnion_PacketQuarantineEncResponse)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketQuarantineEncResponse{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(17232):
		_o, _ := o.Value.(*PacketTypeUnion_PacketCapsResponse)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketCapsResponse{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(18258):
		_o, _ := o.Value.(*PacketTypeUnion_PacketMessageRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketMessageRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(18256):
		_o, _ := o.Value.(*PacketTypeUnion_PacketMessageResponse)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketMessageResponse{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16468):
		_o, _ := o.Value.(*PacketTypeUnion_PacketAuth)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketAuth{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(21072):
		_o, _ := o.Value.(*PacketTypeUnion_PacketReauth)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PacketTypeUnion_PacketReauth{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *PacketTypeUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
	case uint32(18500):
		o.Value = &PacketTypeUnion_PacketHeader{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(22083):
		o.Value = &PacketTypeUnion_PacketVersionCaps{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(20803):
		o.Value = &PacketTypeUnion_PacketQuarantineConfigRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(20818):
		o.Value = &PacketTypeUnion_PacketQuarantineRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(20562):
		o.Value = &PacketTypeUnion_PacketResponse{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(17746):
		o.Value = &PacketTypeUnion_PacketQuarantineEncResponse{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(17232):
		o.Value = &PacketTypeUnion_PacketCapsResponse{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(18258):
		o.Value = &PacketTypeUnion_PacketMessageRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(18256):
		o.Value = &PacketTypeUnion_PacketMessageResponse{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16468):
		o.Value = &PacketTypeUnion_PacketAuth{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(21072):
		o.Value = &PacketTypeUnion_PacketReauth{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// PacketTypeUnion_PacketHeader structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 18500
type PacketTypeUnion_PacketHeader struct {
	// packetHeader:  A PTSG_PACKET_HEADER as specified in section 2.2.9.2.1.1.
	PacketHeader *PacketHeader `idl:"name:packetHeader" json:"packet_header"`
}

func (*PacketTypeUnion_PacketHeader) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketHeader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketHeader != nil {
		_ptr_packetHeader := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketHeader != nil {
				if err := o.PacketHeader.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketHeader{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketHeader, _ptr_packetHeader); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketHeader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetHeader := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketHeader == nil {
			o.PacketHeader = &PacketHeader{}
		}
		if err := o.PacketHeader.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetHeader := func(ptr interface{}) { o.PacketHeader = *ptr.(**PacketHeader) }
	if err := w.ReadPointer(&o.PacketHeader, _s_packetHeader, _ptr_packetHeader); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion_PacketVersionCaps structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 22083
type PacketTypeUnion_PacketVersionCaps struct {
	// packetVersionCaps:  A PTSG_PACKET_VERSIONCAPS as specified in section 2.2.9.2.1.2.
	PacketVersionCaps *PacketVersionCaps `idl:"name:packetVersionCaps" json:"packet_version_caps"`
}

func (*PacketTypeUnion_PacketVersionCaps) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketVersionCaps) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketVersionCaps != nil {
		_ptr_packetVersionCaps := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketVersionCaps != nil {
				if err := o.PacketVersionCaps.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketVersionCaps{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketVersionCaps, _ptr_packetVersionCaps); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketVersionCaps) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetVersionCaps := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketVersionCaps == nil {
			o.PacketVersionCaps = &PacketVersionCaps{}
		}
		if err := o.PacketVersionCaps.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetVersionCaps := func(ptr interface{}) { o.PacketVersionCaps = *ptr.(**PacketVersionCaps) }
	if err := w.ReadPointer(&o.PacketVersionCaps, _s_packetVersionCaps, _ptr_packetVersionCaps); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion_PacketQuarantineConfigRequest structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 20803
type PacketTypeUnion_PacketQuarantineConfigRequest struct {
	// packetQuarConfigRequest:  A PTSG_PACKET_QUARCONFIGREQUEST as specified in section
	// 2.2.9.2.1.3.
	PacketQuarantineConfigRequest *PacketQuarantineConfigRequest `idl:"name:packetQuarConfigRequest" json:"packet_quarantine_config_request"`
}

func (*PacketTypeUnion_PacketQuarantineConfigRequest) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketQuarantineConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketQuarantineConfigRequest != nil {
		_ptr_packetQuarConfigRequest := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketQuarantineConfigRequest != nil {
				if err := o.PacketQuarantineConfigRequest.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketQuarantineConfigRequest{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketQuarantineConfigRequest, _ptr_packetQuarConfigRequest); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketQuarantineConfigRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetQuarConfigRequest := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketQuarantineConfigRequest == nil {
			o.PacketQuarantineConfigRequest = &PacketQuarantineConfigRequest{}
		}
		if err := o.PacketQuarantineConfigRequest.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetQuarConfigRequest := func(ptr interface{}) { o.PacketQuarantineConfigRequest = *ptr.(**PacketQuarantineConfigRequest) }
	if err := w.ReadPointer(&o.PacketQuarantineConfigRequest, _s_packetQuarConfigRequest, _ptr_packetQuarConfigRequest); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion_PacketQuarantineRequest structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 20818
type PacketTypeUnion_PacketQuarantineRequest struct {
	// packetQuarRequest:  A PTSG_PACKET_QUARREQUEST as specified in section 2.2.9.2.1.4.
	PacketQuarantineRequest *PacketQuarantineRequest `idl:"name:packetQuarRequest" json:"packet_quarantine_request"`
}

func (*PacketTypeUnion_PacketQuarantineRequest) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketQuarantineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketQuarantineRequest != nil {
		_ptr_packetQuarRequest := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketQuarantineRequest != nil {
				if err := o.PacketQuarantineRequest.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketQuarantineRequest{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketQuarantineRequest, _ptr_packetQuarRequest); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketQuarantineRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetQuarRequest := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketQuarantineRequest == nil {
			o.PacketQuarantineRequest = &PacketQuarantineRequest{}
		}
		if err := o.PacketQuarantineRequest.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetQuarRequest := func(ptr interface{}) { o.PacketQuarantineRequest = *ptr.(**PacketQuarantineRequest) }
	if err := w.ReadPointer(&o.PacketQuarantineRequest, _s_packetQuarRequest, _ptr_packetQuarRequest); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion_PacketResponse structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 20562
type PacketTypeUnion_PacketResponse struct {
	// packetResponse:  A PTSG_PACKET_RESPONSE as specified in section 2.2.9.2.1.5.
	PacketResponse *PacketResponse `idl:"name:packetResponse" json:"packet_response"`
}

func (*PacketTypeUnion_PacketResponse) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketResponse != nil {
		_ptr_packetResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketResponse != nil {
				if err := o.PacketResponse.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketResponse{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketResponse, _ptr_packetResponse); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketResponse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketResponse == nil {
			o.PacketResponse = &PacketResponse{}
		}
		if err := o.PacketResponse.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetResponse := func(ptr interface{}) { o.PacketResponse = *ptr.(**PacketResponse) }
	if err := w.ReadPointer(&o.PacketResponse, _s_packetResponse, _ptr_packetResponse); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion_PacketQuarantineEncResponse structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 17746
type PacketTypeUnion_PacketQuarantineEncResponse struct {
	// packetQuarEncResponse:  A PTSG_PACKET_QUARENC_RESPONSE as specified in section 2.2.9.2.1.6.
	PacketQuarantineEncResponse *PacketQuarantineEncResponse `idl:"name:packetQuarEncResponse" json:"packet_quarantine_enc_response"`
}

func (*PacketTypeUnion_PacketQuarantineEncResponse) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketQuarantineEncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketQuarantineEncResponse != nil {
		_ptr_packetQuarEncResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketQuarantineEncResponse != nil {
				if err := o.PacketQuarantineEncResponse.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketQuarantineEncResponse{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketQuarantineEncResponse, _ptr_packetQuarEncResponse); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketQuarantineEncResponse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetQuarEncResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketQuarantineEncResponse == nil {
			o.PacketQuarantineEncResponse = &PacketQuarantineEncResponse{}
		}
		if err := o.PacketQuarantineEncResponse.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetQuarEncResponse := func(ptr interface{}) { o.PacketQuarantineEncResponse = *ptr.(**PacketQuarantineEncResponse) }
	if err := w.ReadPointer(&o.PacketQuarantineEncResponse, _s_packetQuarEncResponse, _ptr_packetQuarEncResponse); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion_PacketCapsResponse structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 17232
type PacketTypeUnion_PacketCapsResponse struct {
	// packetCapsResponse:  A PTSG_PACKET_CAPS_RESPONSE as specified in section 2.2.9.2.1.7.
	PacketCapsResponse *PacketCapsResponse `idl:"name:packetCapsResponse" json:"packet_caps_response"`
}

func (*PacketTypeUnion_PacketCapsResponse) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketCapsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketCapsResponse != nil {
		_ptr_packetCapsResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketCapsResponse != nil {
				if err := o.PacketCapsResponse.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketCapsResponse{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketCapsResponse, _ptr_packetCapsResponse); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketCapsResponse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetCapsResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketCapsResponse == nil {
			o.PacketCapsResponse = &PacketCapsResponse{}
		}
		if err := o.PacketCapsResponse.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetCapsResponse := func(ptr interface{}) { o.PacketCapsResponse = *ptr.(**PacketCapsResponse) }
	if err := w.ReadPointer(&o.PacketCapsResponse, _s_packetCapsResponse, _ptr_packetCapsResponse); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion_PacketMessageRequest structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 18258
type PacketTypeUnion_PacketMessageRequest struct {
	// packetMsgRequest:  A PTSG_PACKET_MSG_REQUEST as specified in section 2.2.9.2.1.8.
	PacketMessageRequest *PacketMessageRequest `idl:"name:packetMsgRequest" json:"packet_message_request"`
}

func (*PacketTypeUnion_PacketMessageRequest) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketMessageRequest != nil {
		_ptr_packetMsgRequest := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketMessageRequest != nil {
				if err := o.PacketMessageRequest.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketMessageRequest{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketMessageRequest, _ptr_packetMsgRequest); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketMessageRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetMsgRequest := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketMessageRequest == nil {
			o.PacketMessageRequest = &PacketMessageRequest{}
		}
		if err := o.PacketMessageRequest.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetMsgRequest := func(ptr interface{}) { o.PacketMessageRequest = *ptr.(**PacketMessageRequest) }
	if err := w.ReadPointer(&o.PacketMessageRequest, _s_packetMsgRequest, _ptr_packetMsgRequest); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion_PacketMessageResponse structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 18256
type PacketTypeUnion_PacketMessageResponse struct {
	// packetMsgResponse:  A PTSG_PACKET_MSG_RESPONSE as specified in section 2.2.9.2.1.9.
	PacketMessageResponse *PacketMessageResponse `idl:"name:packetMsgResponse" json:"packet_message_response"`
}

func (*PacketTypeUnion_PacketMessageResponse) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketMessageResponse != nil {
		_ptr_packetMsgResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketMessageResponse != nil {
				if err := o.PacketMessageResponse.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketMessageResponse{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketMessageResponse, _ptr_packetMsgResponse); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketMessageResponse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetMsgResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketMessageResponse == nil {
			o.PacketMessageResponse = &PacketMessageResponse{}
		}
		if err := o.PacketMessageResponse.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetMsgResponse := func(ptr interface{}) { o.PacketMessageResponse = *ptr.(**PacketMessageResponse) }
	if err := w.ReadPointer(&o.PacketMessageResponse, _s_packetMsgResponse, _ptr_packetMsgResponse); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion_PacketAuth structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 16468
type PacketTypeUnion_PacketAuth struct {
	// packetAuth:  A PTSG_PACKET_AUTH as specified in section 2.2.9.2.1.10.
	PacketAuth *PacketAuth `idl:"name:packetAuth" json:"packet_auth"`
}

func (*PacketTypeUnion_PacketAuth) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketAuth) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketAuth != nil {
		_ptr_packetAuth := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketAuth != nil {
				if err := o.PacketAuth.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketAuth{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketAuth, _ptr_packetAuth); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketAuth) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetAuth := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketAuth == nil {
			o.PacketAuth = &PacketAuth{}
		}
		if err := o.PacketAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetAuth := func(ptr interface{}) { o.PacketAuth = *ptr.(**PacketAuth) }
	if err := w.ReadPointer(&o.PacketAuth, _s_packetAuth, _ptr_packetAuth); err != nil {
		return err
	}
	return nil
}

// PacketTypeUnion_PacketReauth structure represents TSG_PACKET_TYPE_UNION RPC union arm.
//
// It has following labels: 21072
type PacketTypeUnion_PacketReauth struct {
	// packetReauth:  A PTSG_PACKET_REAUTH as specified in section 2.2.9.2.1.11.
	PacketReauth *PacketReauth `idl:"name:packetReauth" json:"packet_reauth"`
}

func (*PacketTypeUnion_PacketReauth) is_PacketTypeUnion() {}

func (o *PacketTypeUnion_PacketReauth) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PacketReauth != nil {
		_ptr_packetReauth := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PacketReauth != nil {
				if err := o.PacketReauth.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PacketReauth{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PacketReauth, _ptr_packetReauth); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PacketTypeUnion_PacketReauth) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_packetReauth := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PacketReauth == nil {
			o.PacketReauth = &PacketReauth{}
		}
		if err := o.PacketReauth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_packetReauth := func(ptr interface{}) { o.PacketReauth = *ptr.(**PacketReauth) }
	if err := w.ReadPointer(&o.PacketReauth, _s_packetReauth, _ptr_packetReauth); err != nil {
		return err
	}
	return nil
}

// Packet structure represents TSG_PACKET RPC structure.
//
// The TSG_PACKET structure specifies the type of structure to be used by the RDG client
// and RDG server.
type Packet struct {
	// packetId:  This value specifies the type of structure pointer contained in the TSGPacket
	// field. Valid values are specified in sections 2.2.5.2.2, 2.2.5.2.3, 2.2.5.2.4, 2.2.5.2.5,
	// 2.2.5.2.6, 2.2.5.2.7,  2.2.5.2.9, 2.2.5.2.10, 2.2.5.2.11, 2.2.5.2.12, and 2.2.5.2.13.
	PacketID uint32 `idl:"name:packetId" json:"packet_id"`
	// TSGPacket:  A union field containing the actual structure pointer corresponding to
	// the value contained in the packetId field. Valid structures for this field are specified
	// in sections 2.2.9.2.1.1, 2.2.9.2.1.2, 2.2.9.2.1.3, 2.2.9.2.1.4, 2.2.9.2.1.5, 2.2.9.2.1.6,
	// 2.2.9.2.1.7, 2.2.9.2.1.8, 2.2.9.2.1.9, 2.2.9.2.1.10, and 2.2.9.2.1.11.
	Packet *PacketTypeUnion `idl:"name:TSGPacket;switch_is:packetId" json:"packet"`
}

func (o *Packet) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Packet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PacketID); err != nil {
		return err
	}
	_swPacket := uint32(o.PacketID)
	if o.Packet != nil {
		if err := o.Packet.MarshalUnionNDR(ctx, w, _swPacket); err != nil {
			return err
		}
	} else {
		if err := (&PacketTypeUnion{}).MarshalUnionNDR(ctx, w, _swPacket); err != nil {
			return err
		}
	}
	return nil
}
func (o *Packet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PacketID); err != nil {
		return err
	}
	if o.Packet == nil {
		o.Packet = &PacketTypeUnion{}
	}
	_swPacket := uint32(o.PacketID)
	if err := o.Packet.UnmarshalUnionNDR(ctx, w, _swPacket); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultInterfaceClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultInterfaceClient) CreateTunnel(ctx context.Context, in *CreateTunnelRequest, opts ...dcerpc.CallOption) (*CreateTunnelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateTunnelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInterfaceClient) AuthorizeTunnel(ctx context.Context, in *AuthorizeTunnelRequest, opts ...dcerpc.CallOption) (*AuthorizeTunnelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AuthorizeTunnelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInterfaceClient) MakeTunnelCall(ctx context.Context, in *MakeTunnelCallRequest, opts ...dcerpc.CallOption) (*MakeTunnelCallResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MakeTunnelCallResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInterfaceClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...dcerpc.CallOption) (*CreateChannelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateChannelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInterfaceClient) CloseChannel(ctx context.Context, in *CloseChannelRequest, opts ...dcerpc.CallOption) (*CloseChannelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseChannelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInterfaceClient) CloseTunnel(ctx context.Context, in *CloseTunnelRequest, opts ...dcerpc.CallOption) (*CloseTunnelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseTunnelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInterfaceClient) SetupReceivePipe(ctx context.Context, in *SetupReceivePipeRequest, opts ...dcerpc.CallOption) (*SetupReceivePipeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetupReceivePipeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInterfaceClient) SendToServer(ctx context.Context, in *SendToServerRequest, opts ...dcerpc.CallOption) (*SendToServerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SendToServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInterfaceClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultInterfaceClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewInterfaceClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (InterfaceClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(InterfaceSyntaxV1_3))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultInterfaceClient{cc: cc}, nil
}

// xxx_CreateTunnelOperation structure represents the TsProxyCreateTunnel operation
type xxx_CreateTunnelOperation struct {
	Packet         *Packet          `idl:"name:TSGPacket;pointer:ref" json:"packet"`
	PacketResponse *Packet          `idl:"name:TSGPacketResponse;pointer:ref" json:"packet_response"`
	TunnelContext  *TunnelSerialize `idl:"name:tunnelContext" json:"tunnel_context"`
	TunnelID       uint32           `idl:"name:tunnelId" json:"tunnel_id"`
	Return         int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateTunnelOperation) OpNum() int { return 1 }

func (o *xxx_CreateTunnelOperation) OpName() string {
	return "/TsProxyRpcInterface/v1.3/TsProxyCreateTunnel"
}

func (o *xxx_CreateTunnelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTunnelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// TSGPacket {in} (1:{pointer=ref, alias=PTSG_PACKET}*(1))(2:{alias=TSG_PACKET}(struct))
	{
		if o.Packet != nil {
			if err := o.Packet.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Packet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTunnelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// TSGPacket {in} (1:{pointer=ref, alias=PTSG_PACKET}*(1))(2:{alias=TSG_PACKET}(struct))
	{
		if o.Packet == nil {
			o.Packet = &Packet{}
		}
		if err := o.Packet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTunnelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTunnelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TSGPacketResponse {out} (1:{pointer=ref}*(2))(2:{alias=PTSG_PACKET}*(1))(3:{alias=TSG_PACKET}(struct))
	{
		if o.PacketResponse != nil {
			_ptr_TSGPacketResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PacketResponse != nil {
					if err := o.PacketResponse.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Packet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PacketResponse, _ptr_TSGPacketResponse); err != nil {
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
	// tunnelContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.TunnelContext != nil {
			if err := o.TunnelContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TunnelSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// tunnelId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TunnelID); err != nil {
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

func (o *xxx_CreateTunnelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TSGPacketResponse {out} (1:{pointer=ref}*(2))(2:{alias=PTSG_PACKET,pointer=ref}*(1))(3:{alias=TSG_PACKET}(struct))
	{
		_ptr_TSGPacketResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PacketResponse == nil {
				o.PacketResponse = &Packet{}
			}
			if err := o.PacketResponse.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_TSGPacketResponse := func(ptr interface{}) { o.PacketResponse = *ptr.(**Packet) }
		if err := w.ReadPointer(&o.PacketResponse, _s_TSGPacketResponse, _ptr_TSGPacketResponse); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// tunnelContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.TunnelContext == nil {
			o.TunnelContext = &TunnelSerialize{}
		}
		if err := o.TunnelContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// tunnelId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TunnelID); err != nil {
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

// CreateTunnelRequest structure represents the TsProxyCreateTunnel operation request
type CreateTunnelRequest struct {
	// TSGPacket: Pointer to the TSG_PACKET structure. If this call is made for a reauthentication,
	// then the packetId field MUST be set to TSG_PACKET_TYPE_REAUTH and the packetReauth
	// field of the TSGPacket union field MUST be a pointer to the TSG_PACKET_REAUTH structure.
	// Otherwise, if this call is made for a new connection and the RDG server is configured
	// for RPC authentication, then the value of the packetId field MUST be set to TSG_PACKET_TYPE_VERSIONCAPS
	// and the packetVersionCaps field of the TSGPacket union field MUST be a pointer to
	// the TSG_PACKET_VERSIONCAPS structure. Otherwise, if this call is made for a new connection
	// and the RDG server is configured for pluggable authentication <37>, then the value
	// of the packetId field MUST be set to TSG_PACKET_TYPE_AUTH and the packetAuth field
	// of the TSGPacket union field MUST be a pointer to the TSG_PACKET_AUTH structure.
	// If TSG_PACKET_AUTH is not populated correctly, the error E_PROXY_COOKIE_BADPACKET
	// is returned.<38>
	Packet *Packet `idl:"name:TSGPacket;pointer:ref" json:"packet"`
}

func (o *CreateTunnelRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateTunnelOperation) *xxx_CreateTunnelOperation {
	if op == nil {
		op = &xxx_CreateTunnelOperation{}
	}
	if o == nil {
		return op
	}
	op.Packet = o.Packet
	return op
}

func (o *CreateTunnelRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateTunnelOperation) {
	if o == nil {
		return
	}
	o.Packet = op.Packet
}
func (o *CreateTunnelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateTunnelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateTunnelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateTunnelResponse structure represents the TsProxyCreateTunnel operation response
type CreateTunnelResponse struct {
	// TSGPacketResponse: Pointer to the TSG_PACKET structure. If TSG_MESSAGING_CAP_CONSENT_SIGN
	// capability is negotiated, the packetId member of the TSGPacketResponse out parameter
	// MUST be TSG_PACKET_TYPE_CAPS_RESPONSE and the packetCapsResponse field of the TSGPacket
	// union field MUST be a pointer to the TSG_PACKET_CAPS_RESPONSE (section 2.2.9.2.1.7).
	// Otherwise, the packetId member of TSGPacketResponse MUST be TSG_PACKET_TYPE_QUARENC_RESPONSE,
	// and the packetQuarEncResponse field of the TSGPacket union field MUST be a pointer
	// to the TSG_PACKET_QUARENC_RESPONSE structure. The ADM element Nonce MUST be initialized
	// to a unique GUID and assigned to the nonce field of the TSG_PACKET_QUARENC_RESPONSE
	// structure either in TSGPacketResponse->TSGPacket.packetQuarEncResponse or TSGPacketResponse->TSGPacket.packetCapsResponse->pktQuarEncResponse.
	PacketResponse *Packet `idl:"name:TSGPacketResponse;pointer:ref" json:"packet_response"`
	// tunnelContext: An RPC context handle that represents context-specific information
	// for the tunnel (2). The RDG server MUST provide a non-NULL value. The RDG client
	// MUST save and use this context handle on all subsequent methods calls on the tunnel
	// (2). The methods are TsProxyAuthorizeTunnel, TsProxyCreateChannel, and TsProxyCloseTunnel.
	TunnelContext *TunnelSerialize `idl:"name:tunnelContext" json:"tunnel_context"`
	// tunnelId: An unsigned long identifier representing the tunnel (2). The RDG server
	// MUST save this value in the ADM element Tunnel id and SHOULD provide this value to
	// the RDG client. The RDG client SHOULD save the tunnel id for future use on the RDG
	// client itself. This tunnel id is not required on any future method calls to the RDG
	// server; the tunnelContext is used instead.
	TunnelID uint32 `idl:"name:tunnelId" json:"tunnel_id"`
	// Return: The TsProxyCreateTunnel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateTunnelResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateTunnelOperation) *xxx_CreateTunnelOperation {
	if op == nil {
		op = &xxx_CreateTunnelOperation{}
	}
	if o == nil {
		return op
	}
	op.PacketResponse = o.PacketResponse
	op.TunnelContext = o.TunnelContext
	op.TunnelID = o.TunnelID
	op.Return = o.Return
	return op
}

func (o *CreateTunnelResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateTunnelOperation) {
	if o == nil {
		return
	}
	o.PacketResponse = op.PacketResponse
	o.TunnelContext = op.TunnelContext
	o.TunnelID = op.TunnelID
	o.Return = op.Return
}
func (o *CreateTunnelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateTunnelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateTunnelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AuthorizeTunnelOperation structure represents the TsProxyAuthorizeTunnel operation
type xxx_AuthorizeTunnelOperation struct {
	TunnelContext  *TunnelNoSerialize `idl:"name:tunnelContext" json:"tunnel_context"`
	Packet         *Packet            `idl:"name:TSGPacket;pointer:ref" json:"packet"`
	PacketResponse *Packet            `idl:"name:TSGPacketResponse;pointer:ref" json:"packet_response"`
	Return         int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_AuthorizeTunnelOperation) OpNum() int { return 2 }

func (o *xxx_AuthorizeTunnelOperation) OpName() string {
	return "/TsProxyRpcInterface/v1.3/TsProxyAuthorizeTunnel"
}

func (o *xxx_AuthorizeTunnelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AuthorizeTunnelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// tunnelContext {in} (1:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.TunnelContext != nil {
			if err := o.TunnelContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TunnelNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TSGPacket {in} (1:{pointer=ref, alias=PTSG_PACKET}*(1))(2:{alias=TSG_PACKET}(struct))
	{
		if o.Packet != nil {
			if err := o.Packet.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Packet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AuthorizeTunnelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// tunnelContext {in} (1:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.TunnelContext == nil {
			o.TunnelContext = &TunnelNoSerialize{}
		}
		if err := o.TunnelContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TSGPacket {in} (1:{pointer=ref, alias=PTSG_PACKET}*(1))(2:{alias=TSG_PACKET}(struct))
	{
		if o.Packet == nil {
			o.Packet = &Packet{}
		}
		if err := o.Packet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AuthorizeTunnelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AuthorizeTunnelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TSGPacketResponse {out} (1:{pointer=ref}*(2))(2:{alias=PTSG_PACKET}*(1))(3:{alias=TSG_PACKET}(struct))
	{
		if o.PacketResponse != nil {
			_ptr_TSGPacketResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PacketResponse != nil {
					if err := o.PacketResponse.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Packet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PacketResponse, _ptr_TSGPacketResponse); err != nil {
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

func (o *xxx_AuthorizeTunnelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TSGPacketResponse {out} (1:{pointer=ref}*(2))(2:{alias=PTSG_PACKET,pointer=ref}*(1))(3:{alias=TSG_PACKET}(struct))
	{
		_ptr_TSGPacketResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PacketResponse == nil {
				o.PacketResponse = &Packet{}
			}
			if err := o.PacketResponse.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_TSGPacketResponse := func(ptr interface{}) { o.PacketResponse = *ptr.(**Packet) }
		if err := w.ReadPointer(&o.PacketResponse, _s_TSGPacketResponse, _ptr_TSGPacketResponse); err != nil {
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

// AuthorizeTunnelRequest structure represents the TsProxyAuthorizeTunnel operation request
type AuthorizeTunnelRequest struct {
	// tunnelContext: The RDG client MUST provide the RDG server with the same context handle
	// it received from the TsProxyCreateTunnel method call. The RDG server SHOULD throw
	// an exception if the RPC validation and verification fails.
	TunnelContext *TunnelNoSerialize `idl:"name:tunnelContext" json:"tunnel_context"`
	// TSGPacket: Pointer to the TSG_PACKET structure. The value of the packetId field MUST
	// be set to TSG_PACKET_TYPE_QUARREQUEST. If this is set to any other value, the error
	// E_PROXY_NOT_SUPPORTED is returned. The packetQuarRequest field of the TSGPacket union
	// field MUST be a pointer to the TSG_PACKET_QUARREQUEST structure.
	Packet *Packet `idl:"name:TSGPacket;pointer:ref" json:"packet"`
}

func (o *AuthorizeTunnelRequest) xxx_ToOp(ctx context.Context, op *xxx_AuthorizeTunnelOperation) *xxx_AuthorizeTunnelOperation {
	if op == nil {
		op = &xxx_AuthorizeTunnelOperation{}
	}
	if o == nil {
		return op
	}
	op.TunnelContext = o.TunnelContext
	op.Packet = o.Packet
	return op
}

func (o *AuthorizeTunnelRequest) xxx_FromOp(ctx context.Context, op *xxx_AuthorizeTunnelOperation) {
	if o == nil {
		return
	}
	o.TunnelContext = op.TunnelContext
	o.Packet = op.Packet
}
func (o *AuthorizeTunnelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AuthorizeTunnelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AuthorizeTunnelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AuthorizeTunnelResponse structure represents the TsProxyAuthorizeTunnel operation response
type AuthorizeTunnelResponse struct {
	// TSGPacketResponse: Pointer to the TSG_PACKET structure. The value of the packetId
	// field MUST be TSG_PACKET_TYPE_RESPONSE. The packetResponse field of the TSGPacket
	// union field MUST be a pointer to the TSG_PACKET_RESPONSE structure.
	PacketResponse *Packet `idl:"name:TSGPacketResponse;pointer:ref" json:"packet_response"`
	// Return: The TsProxyAuthorizeTunnel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AuthorizeTunnelResponse) xxx_ToOp(ctx context.Context, op *xxx_AuthorizeTunnelOperation) *xxx_AuthorizeTunnelOperation {
	if op == nil {
		op = &xxx_AuthorizeTunnelOperation{}
	}
	if o == nil {
		return op
	}
	op.PacketResponse = o.PacketResponse
	op.Return = o.Return
	return op
}

func (o *AuthorizeTunnelResponse) xxx_FromOp(ctx context.Context, op *xxx_AuthorizeTunnelOperation) {
	if o == nil {
		return
	}
	o.PacketResponse = op.PacketResponse
	o.Return = op.Return
}
func (o *AuthorizeTunnelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AuthorizeTunnelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AuthorizeTunnelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MakeTunnelCallOperation structure represents the TsProxyMakeTunnelCall operation
type xxx_MakeTunnelCallOperation struct {
	TunnelContext  *TunnelNoSerialize `idl:"name:tunnelContext" json:"tunnel_context"`
	ProcID         uint32             `idl:"name:procId" json:"proc_id"`
	Packet         *Packet            `idl:"name:TSGPacket;pointer:ref" json:"packet"`
	PacketResponse *Packet            `idl:"name:TSGPacketResponse;pointer:ref" json:"packet_response"`
	Return         int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_MakeTunnelCallOperation) OpNum() int { return 3 }

func (o *xxx_MakeTunnelCallOperation) OpName() string {
	return "/TsProxyRpcInterface/v1.3/TsProxyMakeTunnelCall"
}

func (o *xxx_MakeTunnelCallOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MakeTunnelCallOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// tunnelContext {in} (1:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.TunnelContext != nil {
			if err := o.TunnelContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TunnelNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// procId {in} (1:(uint32))
	{
		if err := w.WriteData(o.ProcID); err != nil {
			return err
		}
	}
	// TSGPacket {in} (1:{pointer=ref, alias=PTSG_PACKET}*(1))(2:{alias=TSG_PACKET}(struct))
	{
		if o.Packet != nil {
			if err := o.Packet.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Packet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MakeTunnelCallOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// tunnelContext {in} (1:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.TunnelContext == nil {
			o.TunnelContext = &TunnelNoSerialize{}
		}
		if err := o.TunnelContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// procId {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ProcID); err != nil {
			return err
		}
	}
	// TSGPacket {in} (1:{pointer=ref, alias=PTSG_PACKET}*(1))(2:{alias=TSG_PACKET}(struct))
	{
		if o.Packet == nil {
			o.Packet = &Packet{}
		}
		if err := o.Packet.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MakeTunnelCallOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MakeTunnelCallOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TSGPacketResponse {out} (1:{pointer=ref}*(2))(2:{alias=PTSG_PACKET}*(1))(3:{alias=TSG_PACKET}(struct))
	{
		if o.PacketResponse != nil {
			_ptr_TSGPacketResponse := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PacketResponse != nil {
					if err := o.PacketResponse.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Packet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PacketResponse, _ptr_TSGPacketResponse); err != nil {
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

func (o *xxx_MakeTunnelCallOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TSGPacketResponse {out} (1:{pointer=ref}*(2))(2:{alias=PTSG_PACKET,pointer=ref}*(1))(3:{alias=TSG_PACKET}(struct))
	{
		_ptr_TSGPacketResponse := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PacketResponse == nil {
				o.PacketResponse = &Packet{}
			}
			if err := o.PacketResponse.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_TSGPacketResponse := func(ptr interface{}) { o.PacketResponse = *ptr.(**Packet) }
		if err := w.ReadPointer(&o.PacketResponse, _s_TSGPacketResponse, _ptr_TSGPacketResponse); err != nil {
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

// MakeTunnelCallRequest structure represents the TsProxyMakeTunnelCall operation request
type MakeTunnelCallRequest struct {
	// tunnelContext: The RDG client MUST provide the RDG server with the same context handle
	// it received from the TsProxyCreateTunnel method call. The RDG server SHOULD throw
	// an exception if the RPC validation and verification fail.
	TunnelContext *TunnelNoSerialize `idl:"name:tunnelContext" json:"tunnel_context"`
	// procId: This field identifies the work that is performed by the API. This field can
	// have the following values.
	//
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                |                                                                                  |
	//	|                     VALUE                      |                                     MEANING                                      |
	//	|                                                |                                                                                  |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TSG_TUNNEL_CALL_ASYNC_MSG_REQUEST 0x00000001   | Used to request an administrative message when the same is available on the      |
	//	|                                                | server.                                                                          |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TSG_TUNNEL_CANCEL_ASYNC_MSG_REQUEST 0x00000002 | Used to cancel a pending administrative message request.                         |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	ProcID uint32 `idl:"name:procId" json:"proc_id"`
	// TSGPacket: Pointer to the TSG_PACKET structure. The value of the packetId field MUST
	// be set to TSG_PACKET_TYPE_MSGREQUEST_PACKET. The packetMsgRequest field of the TSGPacket
	// union field MUST be a pointer to the TSG_PACKET_MSG_REQUEST structure.
	Packet *Packet `idl:"name:TSGPacket;pointer:ref" json:"packet"`
}

func (o *MakeTunnelCallRequest) xxx_ToOp(ctx context.Context, op *xxx_MakeTunnelCallOperation) *xxx_MakeTunnelCallOperation {
	if op == nil {
		op = &xxx_MakeTunnelCallOperation{}
	}
	if o == nil {
		return op
	}
	op.TunnelContext = o.TunnelContext
	op.ProcID = o.ProcID
	op.Packet = o.Packet
	return op
}

func (o *MakeTunnelCallRequest) xxx_FromOp(ctx context.Context, op *xxx_MakeTunnelCallOperation) {
	if o == nil {
		return
	}
	o.TunnelContext = op.TunnelContext
	o.ProcID = op.ProcID
	o.Packet = op.Packet
}
func (o *MakeTunnelCallRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MakeTunnelCallRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MakeTunnelCallOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeTunnelCallResponse structure represents the TsProxyMakeTunnelCall operation response
type MakeTunnelCallResponse struct {
	// TSGPacketResponse: Pointer to the TSG_PACKET structure. If procId is TSG_TUNNEL_CANCEL_ASYNC_MSG_REQUEST
	// or if the return value is HRESULT_FROM_WIN32(RPC_S_CALL_CANCELLED), *TSGPacketResponse
	// MUST be set to NULL. Otherwise, the value of the packetId field MUST be TSG_PACKET_TYPE_MESSAGE_PACKET.
	// The packetMsgResponse field of the TSGPacket union field MUST be a pointer to the
	// TSG_PACKET_MSG_RESPONSE structure.
	PacketResponse *Packet `idl:"name:TSGPacketResponse;pointer:ref" json:"packet_response"`
	// Return: The TsProxyMakeTunnelCall return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MakeTunnelCallResponse) xxx_ToOp(ctx context.Context, op *xxx_MakeTunnelCallOperation) *xxx_MakeTunnelCallOperation {
	if op == nil {
		op = &xxx_MakeTunnelCallOperation{}
	}
	if o == nil {
		return op
	}
	op.PacketResponse = o.PacketResponse
	op.Return = o.Return
	return op
}

func (o *MakeTunnelCallResponse) xxx_FromOp(ctx context.Context, op *xxx_MakeTunnelCallOperation) {
	if o == nil {
		return
	}
	o.PacketResponse = op.PacketResponse
	o.Return = op.Return
}
func (o *MakeTunnelCallResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MakeTunnelCallResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MakeTunnelCallOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateChannelOperation structure represents the TsProxyCreateChannel operation
type xxx_CreateChannelOperation struct {
	TunnelContext  *TunnelNoSerialize `idl:"name:tunnelContext" json:"tunnel_context"`
	EndpointInfo   *EndpointInfo      `idl:"name:tsEndPointInfo;pointer:ref" json:"endpoint_info"`
	ChannelContext *ChannelSerialize  `idl:"name:channelContext" json:"channel_context"`
	ChannelID      uint32             `idl:"name:channelId" json:"channel_id"`
	Return         int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateChannelOperation) OpNum() int { return 4 }

func (o *xxx_CreateChannelOperation) OpName() string {
	return "/TsProxyRpcInterface/v1.3/TsProxyCreateChannel"
}

func (o *xxx_CreateChannelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateChannelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// tunnelContext {in} (1:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.TunnelContext != nil {
			if err := o.TunnelContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TunnelNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// tsEndPointInfo {in} (1:{pointer=ref, alias=PTSENDPOINTINFO}*(1))(2:{alias=TSENDPOINTINFO}(struct))
	{
		if o.EndpointInfo != nil {
			if err := o.EndpointInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&EndpointInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateChannelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// tunnelContext {in} (1:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.TunnelContext == nil {
			o.TunnelContext = &TunnelNoSerialize{}
		}
		if err := o.TunnelContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// tsEndPointInfo {in} (1:{pointer=ref, alias=PTSENDPOINTINFO}*(1))(2:{alias=TSENDPOINTINFO}(struct))
	{
		if o.EndpointInfo == nil {
			o.EndpointInfo = &EndpointInfo{}
		}
		if err := o.EndpointInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateChannelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateChannelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// channelContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCHANNEL_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.ChannelContext != nil {
			if err := o.ChannelContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ChannelSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// channelId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ChannelID); err != nil {
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

func (o *xxx_CreateChannelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// channelContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCHANNEL_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.ChannelContext == nil {
			o.ChannelContext = &ChannelSerialize{}
		}
		if err := o.ChannelContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// channelId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ChannelID); err != nil {
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

// CreateChannelRequest structure represents the TsProxyCreateChannel operation request
type CreateChannelRequest struct {
	// tunnelContext: The RDG client MUST provide the RDG server with the same context handle
	// it received from the TsProxyCreateTunnel method call. The RDG server SHOULD throw
	// an exception if the RPC validation and verification fails.
	TunnelContext *TunnelNoSerialize `idl:"name:tunnelContext" json:"tunnel_context"`
	// tsEndPointInfo: Pointer to the TSENDPOINTINFO structure. The RDG client MUST provide
	// a non-NULL pointer to the RDG server for this structure. The RDG server initializes
	// the ADM element Target server names with an array of resourceName and alternateResourceNames
	// members of TSENDPOINTINFO structure. The RDG server SHOULD try to connect to the
	// target server by each name in the array until it succeeds or until the array is traversed
	// completely. If connection fails for all target server names, HRESULT_CODE(E_PROXY_TS_CONNECTFAILED)
	// (0x000059DD) is returned.<47> The rules for determining a valid server name are specified
	// in section 2.2.1.1.
	EndpointInfo *EndpointInfo `idl:"name:tsEndPointInfo;pointer:ref" json:"endpoint_info"`
}

func (o *CreateChannelRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateChannelOperation) *xxx_CreateChannelOperation {
	if op == nil {
		op = &xxx_CreateChannelOperation{}
	}
	if o == nil {
		return op
	}
	op.TunnelContext = o.TunnelContext
	op.EndpointInfo = o.EndpointInfo
	return op
}

func (o *CreateChannelRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateChannelOperation) {
	if o == nil {
		return
	}
	o.TunnelContext = op.TunnelContext
	o.EndpointInfo = op.EndpointInfo
}
func (o *CreateChannelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateChannelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateChannelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateChannelResponse structure represents the TsProxyCreateChannel operation response
type CreateChannelResponse struct {
	// channelContext: A RPC context handle that represents context-specific information
	// for the channel. The RDG server MUST provide a non-NULL value. The RDG client MUST
	// save and use this context handle on all subsequent method calls on the channel. Specifically,
	// these methods are TsProxySendToServer, TsProxySetupReceivePipe, and TsProxyCloseChannel.
	ChannelContext *ChannelSerialize `idl:"name:channelContext" json:"channel_context"`
	// channelId: An unsigned long identifying the channel. The RDG server MUST provide
	// this value to the RDG client. The RDG client MUST save the returned channel ID for
	// future use in the ADM element Channel id (section 3.5.1). This channel ID is not
	// required on any future method calls.
	ChannelID uint32 `idl:"name:channelId" json:"channel_id"`
	// Return: The TsProxyCreateChannel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateChannelResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateChannelOperation) *xxx_CreateChannelOperation {
	if op == nil {
		op = &xxx_CreateChannelOperation{}
	}
	if o == nil {
		return op
	}
	op.ChannelContext = o.ChannelContext
	op.ChannelID = o.ChannelID
	op.Return = o.Return
	return op
}

func (o *CreateChannelResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateChannelOperation) {
	if o == nil {
		return
	}
	o.ChannelContext = op.ChannelContext
	o.ChannelID = op.ChannelID
	o.Return = op.Return
}
func (o *CreateChannelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateChannelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateChannelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseChannelOperation structure represents the TsProxyCloseChannel operation
type xxx_CloseChannelOperation struct {
	Context *ChannelNoSerialize `idl:"name:context" json:"context"`
	Return  int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseChannelOperation) OpNum() int { return 6 }

func (o *xxx_CloseChannelOperation) OpName() string {
	return "/TsProxyRpcInterface/v1.3/TsProxyCloseChannel"
}

func (o *xxx_CloseChannelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
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
	// context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCHANNEL_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ChannelNoSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseChannelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCHANNEL_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &ChannelNoSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
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
	// context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCHANNEL_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ChannelNoSerialize{}).MarshalNDR(ctx, w); err != nil {
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
	// context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCHANNEL_CONTEXT_HANDLE_NOSERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &ChannelNoSerialize{}
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

// CloseChannelRequest structure represents the TsProxyCloseChannel operation request
type CloseChannelRequest struct {
	// context: The RDG client MUST provide the RDG server with the same context handle
	// it received from the TsProxyCreateChannel method call.
	Context *ChannelNoSerialize `idl:"name:context" json:"context"`
}

func (o *CloseChannelRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseChannelOperation) *xxx_CloseChannelOperation {
	if op == nil {
		op = &xxx_CloseChannelOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *CloseChannelRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseChannelOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
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

// CloseChannelResponse structure represents the TsProxyCloseChannel operation response
type CloseChannelResponse struct {
	// context: The RDG client MUST provide the RDG server with the same context handle
	// it received from the TsProxyCreateChannel method call.
	Context *ChannelNoSerialize `idl:"name:context" json:"context"`
	// Return: The TsProxyCloseChannel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseChannelResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseChannelOperation) *xxx_CloseChannelOperation {
	if op == nil {
		op = &xxx_CloseChannelOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Return = o.Return
	return op
}

func (o *CloseChannelResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseChannelOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
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

// xxx_CloseTunnelOperation structure represents the TsProxyCloseTunnel operation
type xxx_CloseTunnelOperation struct {
	Context *TunnelSerialize `idl:"name:context" json:"context"`
	Return  int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseTunnelOperation) OpNum() int { return 7 }

func (o *xxx_CloseTunnelOperation) OpName() string {
	return "/TsProxyRpcInterface/v1.3/TsProxyCloseTunnel"
}

func (o *xxx_CloseTunnelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseTunnelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TunnelSerialize{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseTunnelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &TunnelSerialize{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseTunnelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseTunnelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TunnelSerialize{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CloseTunnelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// context {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PTUNNEL_CONTEXT_HANDLE_SERIALIZE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &TunnelSerialize{}
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

// CloseTunnelRequest structure represents the TsProxyCloseTunnel operation request
type CloseTunnelRequest struct {
	// context: The RDG client MUST provide the RDG server with the same context handle
	// it received from the TsProxyCreateTunnel method call.
	Context *TunnelSerialize `idl:"name:context" json:"context"`
}

func (o *CloseTunnelRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseTunnelOperation) *xxx_CloseTunnelOperation {
	if op == nil {
		op = &xxx_CloseTunnelOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *CloseTunnelRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseTunnelOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *CloseTunnelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseTunnelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseTunnelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseTunnelResponse structure represents the TsProxyCloseTunnel operation response
type CloseTunnelResponse struct {
	// context: The RDG client MUST provide the RDG server with the same context handle
	// it received from the TsProxyCreateTunnel method call.
	Context *TunnelSerialize `idl:"name:context" json:"context"`
	// Return: The TsProxyCloseTunnel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseTunnelResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseTunnelOperation) *xxx_CloseTunnelOperation {
	if op == nil {
		op = &xxx_CloseTunnelOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Return = o.Return
	return op
}

func (o *CloseTunnelResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseTunnelOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Return = op.Return
}
func (o *CloseTunnelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseTunnelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseTunnelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetupReceivePipeOperation structure represents the TsProxySetupReceivePipe operation
type xxx_SetupReceivePipeOperation struct {
	Message []byte `idl:"name:pRpcMessage;max_is:(32767)" json:"message"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetupReceivePipeOperation) OpNum() int { return 8 }

func (o *xxx_SetupReceivePipeOperation) OpName() string {
	return "/TsProxyRpcInterface/v1.3/TsProxySetupReceivePipe"
}

func (o *xxx_SetupReceivePipeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	// cannot evaluate expression 32768
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetupReceivePipeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pRpcMessage {in} (1:[dim:0,max_is=32767](uint8))
	{
		dimSize1 := uint64(32768)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Message {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Message[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Message); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetupReceivePipeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pRpcMessage {in} (1:[dim:0,max_is=32767](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Message", sizeInfo[0])
		}
		o.Message = make([]byte, sizeInfo[0])
		for i1 := range o.Message {
			i1 := i1
			if err := w.ReadData(&o.Message[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetupReceivePipeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetupReceivePipeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetupReceivePipeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetupReceivePipeRequest structure represents the TsProxySetupReceivePipe operation request
type SetupReceivePipeRequest struct {
	// pRpcMessage: The protocol data between RDG client and RDG server MUST be decoded
	// as specified in section 2.2.9.4. RPC stub information is specified in [MS-RPCE] sections
	// 1.1 and 1.5.
	Message []byte `idl:"name:pRpcMessage;max_is:(32767)" json:"message"`
}

func (o *SetupReceivePipeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetupReceivePipeOperation) *xxx_SetupReceivePipeOperation {
	if op == nil {
		op = &xxx_SetupReceivePipeOperation{}
	}
	if o == nil {
		return op
	}
	op.Message = o.Message
	return op
}

func (o *SetupReceivePipeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetupReceivePipeOperation) {
	if o == nil {
		return
	}
	o.Message = op.Message
}
func (o *SetupReceivePipeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetupReceivePipeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetupReceivePipeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetupReceivePipeResponse structure represents the TsProxySetupReceivePipe operation response
type SetupReceivePipeResponse struct {
	// Return: The TsProxySetupReceivePipe return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetupReceivePipeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetupReceivePipeOperation) *xxx_SetupReceivePipeOperation {
	if op == nil {
		op = &xxx_SetupReceivePipeOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetupReceivePipeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetupReceivePipeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetupReceivePipeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetupReceivePipeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetupReceivePipeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SendToServerOperation structure represents the TsProxySendToServer operation
type xxx_SendToServerOperation struct {
	Message []byte `idl:"name:pRpcMessage;max_is:(32767)" json:"message"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SendToServerOperation) OpNum() int { return 9 }

func (o *xxx_SendToServerOperation) OpName() string {
	return "/TsProxyRpcInterface/v1.3/TsProxySendToServer"
}

func (o *xxx_SendToServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	// cannot evaluate expression 32768
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendToServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pRpcMessage {in} (1:[dim:0,max_is=32767](uint8))
	{
		dimSize1 := uint64(32768)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Message {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Message[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Message); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SendToServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pRpcMessage {in} (1:[dim:0,max_is=32767](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Message", sizeInfo[0])
		}
		o.Message = make([]byte, sizeInfo[0])
		for i1 := range o.Message {
			i1 := i1
			if err := w.ReadData(&o.Message[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SendToServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendToServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SendToServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SendToServerRequest structure represents the TsProxySendToServer operation request
type SendToServerRequest struct {
	// pRpcMessage: The protocol data between RDG client and RDG server MUST be decoded
	// as specified in section 2.2.9.3. RPC stub information is specified in [MS-RPCE] sections
	// 1.1 and 1.5.
	Message []byte `idl:"name:pRpcMessage;max_is:(32767)" json:"message"`
}

func (o *SendToServerRequest) xxx_ToOp(ctx context.Context, op *xxx_SendToServerOperation) *xxx_SendToServerOperation {
	if op == nil {
		op = &xxx_SendToServerOperation{}
	}
	if o == nil {
		return op
	}
	op.Message = o.Message
	return op
}

func (o *SendToServerRequest) xxx_FromOp(ctx context.Context, op *xxx_SendToServerOperation) {
	if o == nil {
		return
	}
	o.Message = op.Message
}
func (o *SendToServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SendToServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendToServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SendToServerResponse structure represents the TsProxySendToServer operation response
type SendToServerResponse struct {
	// Return: The TsProxySendToServer return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SendToServerResponse) xxx_ToOp(ctx context.Context, op *xxx_SendToServerOperation) *xxx_SendToServerOperation {
	if op == nil {
		op = &xxx_SendToServerOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SendToServerResponse) xxx_FromOp(ctx context.Context, op *xxx_SendToServerOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SendToServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SendToServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendToServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
