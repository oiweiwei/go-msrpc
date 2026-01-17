package tsproxyrpcinterface

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

// TsProxyRpcInterface server interface.
type InterfaceServer interface {

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
	CreateTunnel(context.Context, *CreateTunnelRequest) (*CreateTunnelResponse, error)

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
	AuthorizeTunnel(context.Context, *AuthorizeTunnelRequest) (*AuthorizeTunnelResponse, error)

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
	MakeTunnelCall(context.Context, *MakeTunnelCallRequest) (*MakeTunnelCallResponse, error)

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
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)

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
	CloseChannel(context.Context, *CloseChannelRequest) (*CloseChannelResponse, error)

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
	CloseTunnel(context.Context, *CloseTunnelRequest) (*CloseTunnelResponse, error)

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
	SetupReceivePipe(context.Context, *SetupReceivePipeRequest) (*SetupReceivePipeResponse, error)

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
	SendToServer(context.Context, *SendToServerRequest) (*SendToServerResponse, error)
}

func RegisterInterfaceServer(conn dcerpc.Conn, o InterfaceServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewInterfaceServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(InterfaceSyntaxV1_3))...)
}

func NewInterfaceServerHandle(o InterfaceServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return InterfaceServerHandle(ctx, o, opNum, r)
	}
}

func InterfaceServerHandle(ctx context.Context, o InterfaceServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0NotUsedOnWire
		// Opnum0NotUsedOnWire
		return nil, nil
	case 1: // TsProxyCreateTunnel
		op := &xxx_CreateTunnelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateTunnelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateTunnel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // TsProxyAuthorizeTunnel
		op := &xxx_AuthorizeTunnelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AuthorizeTunnelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AuthorizeTunnel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // TsProxyMakeTunnelCall
		op := &xxx_MakeTunnelCallOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MakeTunnelCallRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MakeTunnelCall(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // TsProxyCreateChannel
		op := &xxx_CreateChannelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateChannelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateChannel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // TsProxyCloseChannel
		op := &xxx_CloseChannelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseChannelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseChannel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // TsProxyCloseTunnel
		op := &xxx_CloseTunnelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseTunnelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseTunnel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // TsProxySetupReceivePipe
		op := &xxx_SetupReceivePipeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetupReceivePipeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetupReceivePipe(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // TsProxySendToServer
		op := &xxx_SendToServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendToServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SendToServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented TsProxyRpcInterface
type UnimplementedInterfaceServer struct {
}

func (UnimplementedInterfaceServer) CreateTunnel(context.Context, *CreateTunnelRequest) (*CreateTunnelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInterfaceServer) AuthorizeTunnel(context.Context, *AuthorizeTunnelRequest) (*AuthorizeTunnelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInterfaceServer) MakeTunnelCall(context.Context, *MakeTunnelCallRequest) (*MakeTunnelCallResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInterfaceServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInterfaceServer) CloseChannel(context.Context, *CloseChannelRequest) (*CloseChannelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInterfaceServer) CloseTunnel(context.Context, *CloseTunnelRequest) (*CloseTunnelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInterfaceServer) SetupReceivePipe(context.Context, *SetupReceivePipeRequest) (*SetupReceivePipeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInterfaceServer) SendToServer(context.Context, *SendToServerRequest) (*SendToServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ InterfaceServer = (*UnimplementedInterfaceServer)(nil)
