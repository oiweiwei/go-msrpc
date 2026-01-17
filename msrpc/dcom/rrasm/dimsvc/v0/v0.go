package dimsvc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	rrasm "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm"
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
	_ = rrasm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

var (
	// Syntax UUID
	DimsvcSyntaxUUID = &uuid.UUID{TimeLow: 0x8f09f000, TimeMid: 0xb7ed, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0xd2, Node: [6]uint8{0x0, 0x0, 0x1a, 0x18, 0x1c, 0xad}}
	// Syntax ID
	DimsvcSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DimsvcSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// dimsvc interface.
type DimsvcClient interface {

	// The RMprAdminServerGetInfo method retrieves port-related configuration information
	// for the specified RRAS server using the handle hDimServer. The dwLevel defines the
	// type of information requested. The caller MUST pass a pointer to a valid DIM_INFORMATION_CONTAINER
	// (section 2.2.1.2.1). The caller SHOULD free the memory pointed to by pInfoStruct.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the return value contains an error code, as specified in [MS-ERREF] or
	// in section 2.2.4. All values not in the table that follows MUST be treated the same
	// by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges as specified in      |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 0.
	ServerGetInfo(context.Context, *ServerGetInfoRequest, ...dcerpc.CallOption) (*ServerGetInfoResponse, error)

	// The RRasAdminConnectionEnum method retrieves the list of active connections for a
	// specified RRASM server identified by the handle hDimServer. The dwLevel specifies
	// the type of information requested. The caller MUST pass a pointer to a valid DIM_INFORMATION_CONTAINER
	// (section 2.2.1.2.1), where DIM_INFORMATION_CONTAINER.dwBufferSize is initialized
	// to zero (0). After the function returns, the caller SHOULD free the memory pointed
	// to by pInfoStruct.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the return value contains an error code, as specified in [MS-ERREF] or
	// in section 2.2.4. All values that are not in the table that follows MUST be treated
	// the same by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges as specified in      |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_MORE_DATA 0x000000EA     | Not all the data was returned with this call. To obtain additional data, call    |
	//	|                                | the function again using the handle that was returned in the lpdwResumeHandle    |
	//	|                                | parameter.                                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 1.
	ConnectionEnum(context.Context, *ConnectionEnumRequest, ...dcerpc.CallOption) (*ConnectionEnumResponse, error)

	// The RRasAdminConnectionGetInfo method retrieves the connection information for a
	// particular connection identified by the passed handle of the connection given in
	// hDimConnection. The RRAS server is identified by the server handle passed in hDimServer.
	// The dwLevel defines the type of information requested. The caller MUST pass a pointer
	// to a valid DIM_INFORMATION_CONTAINER (section 2.2.1.2.1), where DIM_INFORMATION_CONTAINER.dwBufferSize
	// is initialized to zero (0). The caller SHOULD free the memory pointed to by pInfoStruct.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the value contains an error code, as specified in [MS-ERREF] or section
	// 2.2.4. All values that are not in the table that follows MUST be treated the same
	// by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges as specified in      |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 2.
	ConnectionGetInfo(context.Context, *ConnectionGetInfoRequest, ...dcerpc.CallOption) (*ConnectionGetInfoResponse, error)

	// The RRasAdminConnectionClearStats method resets the statistics counters for the connection
	// identified by the passed handle in hDimConnection. The hDimServer handle specifies
	// the RRASM server on which the call is executed.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges as specified in      |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 3.
	ConnectionClearStats(context.Context, *ConnectionClearStatsRequest, ...dcerpc.CallOption) (*ConnectionClearStatsResponse, error)

	// The RRasAdminPortEnum method retrieves the list of all ports on a RRAS server, or
	// the ports for a specified connection determined by hRasConnection. The hDimServer
	// handle specifies the RRAS server on which the call is executed. The dwLevel defines
	// the type of information requested. The caller MUST pass a pointer to a valid DIM_INFORMATION_CONTAINER
	// (section 2.2.1.2.1), where DIM_INFORMATION_CONTAINER.dwBufferSize is initialized
	// to zero (0). The caller SHOULD free the memory pointed to by pInfoStruct.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or section 2.2.4.
	// All values that are not listed in the table that follows MUST be treated the same
	// by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges.                     |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_MORE_DATA 0x000000EA     | Not all of the data was returned with this call. To obtain additional data, call |
	//	|                                | the function again using the handle that was returned in the lpdwResumeHandle    |
	//	|                                | parameter.                                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 4.
	PortEnum(context.Context, *PortEnumRequest, ...dcerpc.CallOption) (*PortEnumResponse, error)

	// The RRasAdminPortGetInfo method retrieves the port information for a particular port
	// given in hPort for a specified RRAS server using the handle hDimServer. The dwLevel
	// defines the type of information requested. The caller MUST pass a pointer to a valid
	// DIM_INFORMATION_CONTAINER (section 2.2.1.2.1), where DIM_INFORMATION_CONTAINER.dwBufferSize
	// is initialized to zero (0). The caller SHOULD free the memory pointed to by pInfoStruct.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the value contains an error code, as specified in [MS-ERREF] or section
	// 2.2.4. All values that are not listed in the table that follows MUST be treated the
	// same by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 5.
	PortGetInfo(context.Context, *PortGetInfoRequest, ...dcerpc.CallOption) (*PortGetInfoResponse, error)

	// The RRasAdminPortClearStats method resets the counters for the specified port on
	// the given server hPort. The hDimServer handle specifies the RRASM server on which
	// the call is to be executed.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise it contains an error code, as specified in [MS-ERREF] or section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 6.
	PortClearStats(context.Context, *PortClearStatsRequest, ...dcerpc.CallOption) (*PortClearStatsResponse, error)

	// The RRasAdminPortReset function performs no action and always returns ERROR_SUCCESS
	// if the access validation succeeds.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values not in the following table MUST be treated the same by the RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 7.
	PortReset(context.Context, *PortResetRequest, ...dcerpc.CallOption) (*PortResetResponse, error)

	// The RRasAdminPortDisconnect method initiates the disconnect of the connection on
	// a specified port. The port SHOULD have been associated with a connection. For the
	// disconnection to be successful, the connection SHOULD have been in the connected
	// state. The connection is the established dial-up or VPN connection that has the RRAS
	// server as its endpoint.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 8.
	PortDisconnect(context.Context, *PortDisconnectRequest, ...dcerpc.CallOption) (*PortDisconnectResponse, error)

	// The RRouterInterfaceTransportSetGlobalInfo method<254> is used to set global information,
	// for the specified transport (IPX, IPv4, or IPv6), such as disabling IPv6 filtering.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 9.
	RouterInterfaceTransportSetGlobalInfo(context.Context, *RouterInterfaceTransportSetGlobalInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportSetGlobalInfoResponse, error)

	// The RRouterInterfaceTransportGetGlobalInfo method<256> is used to get the entire
	// global information for the specified transport.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 10.
	RouterInterfaceTransportGetGlobalInfo(context.Context, *RouterInterfaceTransportGetGlobalInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportGetGlobalInfoResponse, error)

	// The RRouterInterfaceGetHandle method<260> retrieves the handle of the specified interface
	// lpwsInterfaceName among all the ROUTER_INTERFACE_TYPEs. The hDimServer handle determines
	// the RRAS server on which the call is made.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 11.
	RouterInterfaceGetHandle(context.Context, *RouterInterfaceGetHandleRequest, ...dcerpc.CallOption) (*RouterInterfaceGetHandleResponse, error)

	// The RRouterInterfaceCreate method creates an interface on a specified RRAS server,
	// hDimServer.<261>
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 12.
	RouterInterfaceCreate(context.Context, *RouterInterfaceCreateRequest, ...dcerpc.CallOption) (*RouterInterfaceCreateResponse, error)

	// The RRouterInterfaceGetInfo method<263> retrieves information for a specified interface,
	// hInterface, on a specified RRAS server, hDimServer. This method is used to find information
	// about existing interfaces on the RRAS. The information is returned in pInfoStruct.
	// The caller SHOULD free the memory pointed to by pInfoStruct.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the return value contains an error code, as specified in [MS-ERREF] or
	// in section 2.2.4. All values that are not listed in the table that follows MUST be
	// treated the same by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 13.
	RouterInterfaceGetInfo(context.Context, *RouterInterfaceGetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceGetInfoResponse, error)

	// The RRouterInterfaceSetInfo method<264> sets information for a specified interface,
	// hInterface, on a specified server, hDimServer.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 14.
	RouterInterfaceSetInfo(context.Context, *RouterInterfaceSetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceSetInfoResponse, error)

	// The RRouterInterfaceDelete method<266> deletes an interface on a specified server.
	// The interface MUST have been created with the RRouterInterfaceCreate (section 3.1.4.13)
	// method.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed
	// successfully; otherwise, the value contains an error code, as specified in [MS-ERREF]
	// or in section 2.2.4. All values that are not listed in the table that follows MUST
	// be treated the same by the RRASM client.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000             | The call was successful.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005       | The calling application does not have sufficient privileges, as specified in     |
	//	|                                      | section 2.1.1.1.                                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERFACE_CONNECTED 0x0000038C | The interface is connected and therefore cannot be deleted. This                 |
	//	|                                      | error is returned if the interface is of type ROUTER_IF_TYPE_CLIENT,             |
	//	|                                      | ROUTER_IF_TYPE_HOME_ROUTER, or ROUTER_IF_TYPE_FULL_ROUTER.                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 15.
	RouterInterfaceDelete(context.Context, *RouterInterfaceDeleteRequest, ...dcerpc.CallOption) (*RouterInterfaceDeleteResponse, error)

	// The RRouterInterfaceTransportRemove method<267> is used to remove an existing transport
	// (IPX, IPv4, or IPv6) from the RRAS server on an interface.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the return value contains an error code, as specified in [MS-ERREF] or
	// in section 2.2.4. All values that are not in the table that follows MUST be treated
	// the same by the RRASM client.
	//
	// The return value is one of the following error codes. All other values MUST be treated
	// the same.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 16.
	RouterInterfaceTransportRemove(context.Context, *RouterInterfaceTransportRemoveRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportRemoveResponse, error)

	// The RRouterInterfaceTransportAdd method SHOULD<269> add a transport (IPX, IPv4, or
	// IPv6) to a specified interface. Note that if a transport already exists on an interface
	// it cannot be added.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values not in the following table MUST be treated the same by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | 2.1.1.1.                                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 17.
	RouterInterfaceTransportAdd(context.Context, *RouterInterfaceTransportAddRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportAddResponse, error)

	// The RRouterInterfaceTransportGetInfo method<273> retrieves information about a transport
	// running on a specified interface. The information retrieved is of the type RTR_INFO_BLOCK_HEADER
	// (section 2.2.1.2.3), encapsulated within a DIM_INTERFACE_CONTAINER (section 2.2.1.2.2).
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 18.
	RouterInterfaceTransportGetInfo(context.Context, *RouterInterfaceTransportGetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportGetInfoResponse, error)

	// The RRouterInterfaceTransportSetInfo method SHOULD<276> set information about a transport
	// running on a specified interface. The information is overwritten using the type RTR_INFO_BLOCK_HEADER
	// (section 2.2.1.2.3), encapsulated within a DIM_INTERFACE_CONTAINER (section 2.2.1.2.2).
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The Opnum field value for this method is 19.
	RouterInterfaceTransportSetInfo(context.Context, *RouterInterfaceTransportSetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportSetInfoResponse, error)

	// The RRouterInterfaceEnum method<280> Retrieves the list of all the interfaces from
	// the specified server. The hDimServer handle specifies the RRASM server on which the
	// call is to be executed. The caller MUST pass a pointer to a valid DIM_INFORMATION_CONTAINER
	// (section 2.2.1.2.1), where DIM_INFORMATION_CONTAINER.dwBufferSize is initialized
	// to zero (0). The caller SHOULD free the memory pointed to by pInfoStruct.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client. This error code value can correspond to a RRAS Management Protocol–specific
	// failure, which takes a value between 600 and 975, or any generic failure.
	//
	// The return value is one of the following error codes. All other error values MUST
	// be treated the same.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_MORE_DATA 0x000000EA     | More information is available; the enumeration can be continued.                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 20.
	RouterInterfaceEnum(context.Context, *RouterInterfaceEnumRequest, ...dcerpc.CallOption) (*RouterInterfaceEnumResponse, error)

	// The RRouterInterfaceConnect method<281> establishes a connection for the specified
	// interface hInterface if it is not already connected. The hDimServer handle specifies
	// the RRASM server on which the call is to be executed.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| PENDING 0x00000258             | An operation is pending.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 21. Whether the call returns immediately
	// or is blocked is decided by fBlocking as previously described.
	RouterInterfaceConnect(context.Context, *RouterInterfaceConnectRequest, ...dcerpc.CallOption) (*RouterInterfaceConnectResponse, error)

	// The RRouterInterfaceDisconnect method disconnects the specified interface, hInterface.
	// The hDimServer handle specifies the RRASM server on which the call is to be executed.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the value contains an error code, as specified in [MS-ERREF] or in section
	// 2.2.4. All values that are not in the table that follows MUST be treated the same
	// by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 22.
	RouterInterfaceDisconnect(context.Context, *RouterInterfaceDisconnectRequest, ...dcerpc.CallOption) (*RouterInterfaceDisconnectResponse, error)

	// The RRouterInterfaceUpdateRoutes method<282> updates routing information for a given
	// transport on a specified interface on the RRAS server. If a routing protocol like
	// RIP is running ([RFC1058]), the new routes learned are updated on the interface.
	// This interface MUST be called only when the interface state is ROUTER_IF_STATE_CONNECTED,
	// otherwise an error is returned.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 23.
	RouterInterfaceUpdateRoutes(context.Context, *RouterInterfaceUpdateRoutesRequest, ...dcerpc.CallOption) (*RouterInterfaceUpdateRoutesResponse, error)

	// The RRouterInterfaceQueryUpdateResult method<284> returns the result of the last
	// RRouterInterfaceUpdateRoutes (section 3.1.4.24) request of the RRAS server for a
	// specified transport to update its routes for an interface. This method MUST be called
	// only once after RRouterInterfaceUpdateRoutes.
	//
	// Return Values: A-32 bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the return value contains an error code, as specified in [MS-ERREF] or
	// in section 2.2.4. All values that are not in the table that follows MUST be treated
	// the same by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 24.
	RouterInterfaceQueryUpdateResult(context.Context, *RouterInterfaceQueryUpdateResultRequest, ...dcerpc.CallOption) (*RouterInterfaceQueryUpdateResultResponse, error)

	// The RRouterInterfaceUpdatePhonebookInfo method<286> forces the router to pick up
	// changes made on a specified demand-dial interface, hInterface. The hDimServer handle
	// specifies the RRASM server on which the call is to be executed. Call this method
	// after changes are made to a phone-book entry for a demand-dial interface, such as
	// ROUTER_IF_TYPE_HOME_ROUTER or ROUTER_IF_TYPE_FULL_ROUTER.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 25.
	RouterInterfaceUpdatePhonebookInfo(context.Context, *RouterInterfaceUpdatePhonebookInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceUpdatePhonebookInfoResponse, error)

	// The RMIBEntryCreate method<287> creates an MIB entry that is used by the RRAS to
	// create a route entry in the IPv4 routing table.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 26.
	MIBEntryCreate(context.Context, *MIBEntryCreateRequest, ...dcerpc.CallOption) (*MIBEntryCreateResponse, error)

	// The RMIBEntryDelete method<288> deletes a MIB entry in an IPv4 forwarding table.
	// This causes the IPv4 routing table entry to be deleted.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the following table MUST be treated the same by the RRASM
	// client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 27.
	MIBEntryDelete(context.Context, *MIBEntryDeleteRequest, ...dcerpc.CallOption) (*MIBEntryDeleteResponse, error)

	// The RMIBEntrySet method<290> modifies an MIB entry in the IPv4 forwarding table.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// All error values MUST be treated the same and as long as ERROR_SUCCESS is not returned
	// the call is deemed failed.
	//
	// The opnum field value for this method is 28.
	MIBEntrySet(context.Context, *MIBEntrySetRequest, ...dcerpc.CallOption) (*MIBEntrySetResponse, error)

	// The RMIBEntryGet method<291> retrieves the value of a RRAS MIB entry that corresponds
	// to the transport and that matches the information specified in pInfoStuct.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 29.
	MIBEntryGet(context.Context, *MIBEntryGetRequest, ...dcerpc.CallOption) (*MIBEntryGetResponse, error)

	// The RMIBEntryGetFirst method SHOULD<294> retrieve the first value of an entry corresponding
	// to the transport, protocol, and the MIB entry specified in pInfoStuct.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the return value contains an error code, as specified in [MS-ERREF] or
	// in section 2.2.4. All values that are not in the table that follows MUST be treated
	// the same by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 30.
	MIBEntryGetFirst(context.Context, *MIBEntryGetFirstRequest, ...dcerpc.CallOption) (*MIBEntryGetFirstResponse, error)

	// The RMIBEntryGetNext method<297> retrieves the next (to the previous call) entry
	// corresponding to the transport, protocol, and the MIB entry specified in pInfoStuct.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the value contains an error code, as specified in [MS-ERREF] or in section
	// 2.2.4. All values that are not in the table that follows MUST be treated the same
	// by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NO_MORE_ITEMS 0x00000103 | No more data available.                                                          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 31.
	MIBEntryGetNext(context.Context, *MIBEntryGetNextRequest, ...dcerpc.CallOption) (*MIBEntryGetNextResponse, error)

	// The RMIBGetTrapInfo method<299> queries the module that set a trap event for more
	// information about the trap. This method cannot be called remotely.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the value contains an error code, as specified in [MS-ERREF] or in section
	// 2.2.4. All values that are not in the table that follows MUST be treated the same
	// by the RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 32.
	MIBGetTrapInfo(context.Context, *MIBGetTrapInfoRequest, ...dcerpc.CallOption) (*MIBGetTrapInfoResponse, error)

	// The RMIBSetTrapInfo method<301> specifies a handle to an event that is signaled whenever
	// a trap needs to be issued.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------------+--------------------------------------------------------------+
	//	|                RETURN                |                                                              |
	//	|              VALUE/CODE              |                         DESCRIPTION                          |
	//	|                                      |                                                              |
	//	+--------------------------------------+--------------------------------------------------------------+
	//	+--------------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000             | The call was successful.                                     |
	//	+--------------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005       | The calling application does not have sufficient privileges. |
	//	+--------------------------------------+--------------------------------------------------------------+
	//	| ERROR_INSUFFICIENT_BUFFER 0x0000007A |                                                              |
	//	+--------------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 33.
	MIBSetTrapInfo(context.Context, *MIBSetTrapInfoRequest, ...dcerpc.CallOption) (*MIBSetTrapInfoResponse, error)

	// The RRasAdminConnectionNotification method<302> registers an event object with the
	// RRAS server so that, if an interface connects or disconnects, the event is signaled.
	// The hDimServer handle specifies on which RRAS server the call is to be executed.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 34.
	ConnectionNotification(context.Context, *ConnectionNotificationRequest, ...dcerpc.CallOption) (*ConnectionNotificationResponse, error)

	// The RRasAdminSendUserMessage method<303> sends a message to the user connected on
	// the connection specified by hDimServer.<304>
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the value contains an error code, as specified in [MS-ERREF] or in section
	// 2.2.4. All values that are not in the table that follows MUST be treated the same
	// by the RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 35.
	SendUserMessage(context.Context, *SendUserMessageRequest, ...dcerpc.CallOption) (*SendUserMessageResponse, error)

	// The RRouterDeviceEnum method<305> retrieves the list of all the devices from the
	// specified server. The hDimServer handle specifies the RRASM server on which the call
	// is executed. The caller MUST pass a pointer to a valid DIM_INFORMATION_CONTAINER
	// (section 2.2.1.2.1), where DIM_INFORMATION_CONTAINER.dwBufferSize is initialized
	// to zero (0). The caller SHOULD free the memory pointed to by pInfoStruct->pBuffer.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the return value contains an error code, as specified in [MS-ERREF] or
	// in section 2.2.4. All values that are not in the table that follows MUST be treated
	// the same by the RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 36.
	RouterDeviceEnum(context.Context, *RouterDeviceEnumRequest, ...dcerpc.CallOption) (*RouterDeviceEnumResponse, error)

	// The RRouterInterfaceTransportCreate method<306> is used to create a new transport
	// on the RRAS server.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 37.
	RouterInterfaceTransportCreate(context.Context, *RouterInterfaceTransportCreateRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportCreateResponse, error)

	// The RRouterInterfaceDeviceGetInfo method<309> retrieves information for a specified
	// device, dwIndex, for a specified interface hInterface, on a specified server, hDimServer.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the return value contains an error code, as specified in [MS-ERREF] or
	// in section 2.2.4. All values not in the following table MUST be treated the same
	// by the RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 38.
	RouterInterfaceDeviceGetInfo(context.Context, *RouterInterfaceDeviceGetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceDeviceGetInfoResponse, error)

	// The RRouterInterfaceDeviceSetInfo method<310> sets the information for a specified
	// device, dwIndex, for a specified interface, hInterface, on a specified server, hDimServer.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 39.
	RouterInterfaceDeviceSetInfo(context.Context, *RouterInterfaceDeviceSetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceDeviceSetInfoResponse, error)

	// The RRouterInterfaceSetCredentialsEx method<312> is used to set extended credentials
	// (other than username and password) information for an interface. This function is
	// used to set credentials information used for the EAP methods or used as a key that
	// is shared between two routers (a preshared key).
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 40.
	RouterInterfaceSetCredentialsEx(context.Context, *RouterInterfaceSetCredentialsExRequest, ...dcerpc.CallOption) (*RouterInterfaceSetCredentialsExResponse, error)

	// The RRouterInterfaceGetCredentialsEx method<313> is used to retrieve extended (other
	// than the username or password) credentials information for the specified interface.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 41.
	RouterInterfaceGetCredentialsEx(context.Context, *RouterInterfaceGetCredentialsExRequest, ...dcerpc.CallOption) (*RouterInterfaceGetCredentialsExResponse, error)

	// The RRasAdminConnectionRemoveQuarantine method<314> removes quarantine filters on
	// a dialed-in RRAS client if the filters were applied as a result of Internet Authentication
	// Service (IAS) policies. This function does not remove the IPv6 quarantine filters.
	// The hDimServer handle specifies the RRASM server on which the call is to be executed.<315>
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 42.
	ConnectionRemoveQuarantine(context.Context, *ConnectionRemoveQuarantineRequest, ...dcerpc.CallOption) (*ConnectionRemoveQuarantineResponse, error)

	// The RMprAdminServerSetInfo method<316> sets ports information on a specified server
	// hDimServer.
	//
	// Return Values: A 32-bit, unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the value contains an error code, as specified in [MS-ERREF] or in section
	// 2.2.4. All values that are not listed in the table that follows MUST be treated the
	// same by the RRASM client.
	//
	//	+------------------------------------------+--------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                    |
	//	|                VALUE/CODE                |                            DESCRIPTION                             |
	//	|                                          |                                                                    |
	//	+------------------------------------------+--------------------------------------------------------------------+
	//	+------------------------------------------+--------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000                 | The call was successful.                                           |
	//	+------------------------------------------+--------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005           | The calling application does not have sufficient privileges.       |
	//	+------------------------------------------+--------------------------------------------------------------------+
	//	| ERROR_SUCCESS_REBOOT_REQUIRED 0x00000BC2 | A system reboot is required for such a change to take effect.<317> |
	//	+------------------------------------------+--------------------------------------------------------------------+
	//
	// The opnum field value for this method is 43.
	ServerSetInfo(context.Context, *ServerSetInfoRequest, ...dcerpc.CallOption) (*ServerSetInfoResponse, error)

	// The RMprAdminServerGetInfoEx method<319> gets the device configuration information
	// for PPTP, L2TP, SSTP, and IKEv2 on a server specified by hDimServer.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 44.
	ServerGetInfoEx(context.Context, *ServerGetInfoExRequest, ...dcerpc.CallOption) (*ServerGetInfoExResponse, error)

	// The RRasAdminConnectionEnumEx method<320> retrieves the list of all active connections
	// for a specified RRAS server using handle hDimServer. The caller SHOULD free the memory
	// pointed to by pRasConections.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the value contains an error code, as specified in [MS-ERREF] or in section
	// 2.2.4. All values that are not in the table that follows MUST be treated the same
	// by the RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_MORE_DATA 0x000000EA     | Not all of the data was returned with this call. To obtain additional data, call |
	//	|                                | the function again using the resume handle.                                      |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 45.
	ConnectionEnumEx(context.Context, *ConnectionEnumExRequest, ...dcerpc.CallOption) (*ConnectionEnumExResponse, error)

	// The RRasAdminConnectionGetInfoEx method<321> retrieves the connection information
	// for a specific connection given in hDimConnection for a specified RRAS server using
	// handle hDimServer. The caller SHOULD free the memory pointed to by pRasConnection.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges, as specified in     |
	//	|                                | section 2.1.1.1.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 46.
	ConnectionGetInfoEx(context.Context, *ConnectionGetInfoExRequest, ...dcerpc.CallOption) (*ConnectionGetInfoExResponse, error)

	// The RMprAdminServerSetInfoEx method<322> sets the device configuration information
	// for PPTP, L2TP, SSTP, and IKEv2 on a specified server hDimServer.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values not in the table that follows MUST be treated the same by the RRASM client.
	//
	//	+-------------------------------------------+-------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                               |
	//	|                VALUE/CODE                 |                                  DESCRIPTION                                  |
	//	|                                           |                                                                               |
	//	+-------------------------------------------+-------------------------------------------------------------------------------+
	//	+-------------------------------------------+-------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000                  | The call was successful.                                                      |
	//	+-------------------------------------------+-------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005            | The calling application does not have sufficient privileges.                  |
	//	+-------------------------------------------+-------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS_REBOOT_REQUIRED 0x00000BC2  | A system reboot is required for such a change to take effect.                 |
	//	+-------------------------------------------+-------------------------------------------------------------------------------+
	//	| ERROR_SUCCESS_RESTART_REQUIRED 0x00000BC3 | A remote access service restart is required for such a change to take effect. |
	//	+-------------------------------------------+-------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 47.
	ServerSetInfoEx(context.Context, *ServerSetInfoExRequest, ...dcerpc.CallOption) (*ServerSetInfoExResponse, error)

	// The RRasAdminUpdateConnection method<325> updates the endpoint by sending MOBIKE
	// on a connection specified by hDimConnection on a specified server, hDimServer.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, the return value contains an error code, as specified in [MS-ERREF] or
	// in section 2.2.4. All values that are not included in the table that follows MUST
	// be treated the same by the RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The Opnum field value for this method is 48.
	UpdateConnection(context.Context, *UpdateConnectionRequest, ...dcerpc.CallOption) (*UpdateConnectionResponse, error)

	// The RRouterInterfaceSetCredentialsLocal method<326> is used to set credentials information
	// for an interface. Specifically, the username, domain name, and password that are
	// used in user authentication are set using this method.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates the operation was completed successfully;
	// otherwise it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 49.
	RouterInterfaceSetCredentialsLocal(context.Context, *RouterInterfaceSetCredentialsLocalRequest, ...dcerpc.CallOption) (*RouterInterfaceSetCredentialsLocalResponse, error)

	// The RRouterInterfaceGetCredentialsLocal method<327> is used to get credentials information
	// for an interface. Specifically, the username, domain name, and password used in user
	// authentication are retrieved with this method.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates the operation was completed successfully;
	// otherwise it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 50.
	RouterInterfaceGetCredentialsLocal(context.Context, *RouterInterfaceGetCredentialsLocalRequest, ...dcerpc.CallOption) (*RouterInterfaceGetCredentialsLocalResponse, error)

	// The RRouterInterfaceGetCustomInfoEx method<328> is used to get the tunnel-specific
	// custom configuration for an interface.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	RouterInterfaceGetCustomInfoEx(context.Context, *RouterInterfaceGetCustomInfoExRequest, ...dcerpc.CallOption) (*RouterInterfaceGetCustomInfoExResponse, error)

	// The RRouterInterfaceSetCustomInfoEx method<329> is used to set the tunnel-specific
	// custom configuration for an interface.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// otherwise, it contains an error code, as specified in [MS-ERREF] or in section 2.2.4.
	// All values that are not in the table that follows MUST be treated the same by the
	// RRASM client.
	//
	//	+--------------------------------+--------------------------------------------------------------+
	//	|             RETURN             |                                                              |
	//	|           VALUE/CODE           |                         DESCRIPTION                          |
	//	|                                |                                                              |
	//	+--------------------------------+--------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_SUCCESS 0x00000000       | The call was successful.                                     |
	//	+--------------------------------+--------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The calling application does not have sufficient privileges. |
	//	+--------------------------------+--------------------------------------------------------------+
	RouterInterfaceSetCustomInfoEx(context.Context, *RouterInterfaceSetCustomInfoExRequest, ...dcerpc.CallOption) (*RouterInterfaceSetCustomInfoExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultDimsvcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultDimsvcClient) ServerGetInfo(ctx context.Context, in *ServerGetInfoRequest, opts ...dcerpc.CallOption) (*ServerGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionEnum(ctx context.Context, in *ConnectionEnumRequest, opts ...dcerpc.CallOption) (*ConnectionEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionGetInfo(ctx context.Context, in *ConnectionGetInfoRequest, opts ...dcerpc.CallOption) (*ConnectionGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionClearStats(ctx context.Context, in *ConnectionClearStatsRequest, opts ...dcerpc.CallOption) (*ConnectionClearStatsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionClearStatsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) PortEnum(ctx context.Context, in *PortEnumRequest, opts ...dcerpc.CallOption) (*PortEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PortEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) PortGetInfo(ctx context.Context, in *PortGetInfoRequest, opts ...dcerpc.CallOption) (*PortGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PortGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) PortClearStats(ctx context.Context, in *PortClearStatsRequest, opts ...dcerpc.CallOption) (*PortClearStatsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PortClearStatsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) PortReset(ctx context.Context, in *PortResetRequest, opts ...dcerpc.CallOption) (*PortResetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PortResetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) PortDisconnect(ctx context.Context, in *PortDisconnectRequest, opts ...dcerpc.CallOption) (*PortDisconnectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PortDisconnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportSetGlobalInfo(ctx context.Context, in *RouterInterfaceTransportSetGlobalInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportSetGlobalInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportSetGlobalInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportGetGlobalInfo(ctx context.Context, in *RouterInterfaceTransportGetGlobalInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportGetGlobalInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportGetGlobalInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceGetHandle(ctx context.Context, in *RouterInterfaceGetHandleRequest, opts ...dcerpc.CallOption) (*RouterInterfaceGetHandleResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceGetHandleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceCreate(ctx context.Context, in *RouterInterfaceCreateRequest, opts ...dcerpc.CallOption) (*RouterInterfaceCreateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceCreateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceGetInfo(ctx context.Context, in *RouterInterfaceGetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceSetInfo(ctx context.Context, in *RouterInterfaceSetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceSetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceSetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceDelete(ctx context.Context, in *RouterInterfaceDeleteRequest, opts ...dcerpc.CallOption) (*RouterInterfaceDeleteResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportRemove(ctx context.Context, in *RouterInterfaceTransportRemoveRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportRemoveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportRemoveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportAdd(ctx context.Context, in *RouterInterfaceTransportAddRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportAddResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportAddResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportGetInfo(ctx context.Context, in *RouterInterfaceTransportGetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportSetInfo(ctx context.Context, in *RouterInterfaceTransportSetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportSetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportSetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceEnum(ctx context.Context, in *RouterInterfaceEnumRequest, opts ...dcerpc.CallOption) (*RouterInterfaceEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceConnect(ctx context.Context, in *RouterInterfaceConnectRequest, opts ...dcerpc.CallOption) (*RouterInterfaceConnectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceConnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceDisconnect(ctx context.Context, in *RouterInterfaceDisconnectRequest, opts ...dcerpc.CallOption) (*RouterInterfaceDisconnectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceDisconnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceUpdateRoutes(ctx context.Context, in *RouterInterfaceUpdateRoutesRequest, opts ...dcerpc.CallOption) (*RouterInterfaceUpdateRoutesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceUpdateRoutesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceQueryUpdateResult(ctx context.Context, in *RouterInterfaceQueryUpdateResultRequest, opts ...dcerpc.CallOption) (*RouterInterfaceQueryUpdateResultResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceQueryUpdateResultResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceUpdatePhonebookInfo(ctx context.Context, in *RouterInterfaceUpdatePhonebookInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceUpdatePhonebookInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceUpdatePhonebookInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntryCreate(ctx context.Context, in *MIBEntryCreateRequest, opts ...dcerpc.CallOption) (*MIBEntryCreateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntryCreateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntryDelete(ctx context.Context, in *MIBEntryDeleteRequest, opts ...dcerpc.CallOption) (*MIBEntryDeleteResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntryDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntrySet(ctx context.Context, in *MIBEntrySetRequest, opts ...dcerpc.CallOption) (*MIBEntrySetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntrySetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntryGet(ctx context.Context, in *MIBEntryGetRequest, opts ...dcerpc.CallOption) (*MIBEntryGetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntryGetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntryGetFirst(ctx context.Context, in *MIBEntryGetFirstRequest, opts ...dcerpc.CallOption) (*MIBEntryGetFirstResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntryGetFirstResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntryGetNext(ctx context.Context, in *MIBEntryGetNextRequest, opts ...dcerpc.CallOption) (*MIBEntryGetNextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntryGetNextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBGetTrapInfo(ctx context.Context, in *MIBGetTrapInfoRequest, opts ...dcerpc.CallOption) (*MIBGetTrapInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBGetTrapInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBSetTrapInfo(ctx context.Context, in *MIBSetTrapInfoRequest, opts ...dcerpc.CallOption) (*MIBSetTrapInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBSetTrapInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionNotification(ctx context.Context, in *ConnectionNotificationRequest, opts ...dcerpc.CallOption) (*ConnectionNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) SendUserMessage(ctx context.Context, in *SendUserMessageRequest, opts ...dcerpc.CallOption) (*SendUserMessageResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SendUserMessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterDeviceEnum(ctx context.Context, in *RouterDeviceEnumRequest, opts ...dcerpc.CallOption) (*RouterDeviceEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterDeviceEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportCreate(ctx context.Context, in *RouterInterfaceTransportCreateRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportCreateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportCreateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceDeviceGetInfo(ctx context.Context, in *RouterInterfaceDeviceGetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceDeviceGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceDeviceGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceDeviceSetInfo(ctx context.Context, in *RouterInterfaceDeviceSetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceDeviceSetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceDeviceSetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceSetCredentialsEx(ctx context.Context, in *RouterInterfaceSetCredentialsExRequest, opts ...dcerpc.CallOption) (*RouterInterfaceSetCredentialsExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceSetCredentialsExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceGetCredentialsEx(ctx context.Context, in *RouterInterfaceGetCredentialsExRequest, opts ...dcerpc.CallOption) (*RouterInterfaceGetCredentialsExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceGetCredentialsExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionRemoveQuarantine(ctx context.Context, in *ConnectionRemoveQuarantineRequest, opts ...dcerpc.CallOption) (*ConnectionRemoveQuarantineResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionRemoveQuarantineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ServerSetInfo(ctx context.Context, in *ServerSetInfoRequest, opts ...dcerpc.CallOption) (*ServerSetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerSetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ServerGetInfoEx(ctx context.Context, in *ServerGetInfoExRequest, opts ...dcerpc.CallOption) (*ServerGetInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerGetInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionEnumEx(ctx context.Context, in *ConnectionEnumExRequest, opts ...dcerpc.CallOption) (*ConnectionEnumExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionEnumExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionGetInfoEx(ctx context.Context, in *ConnectionGetInfoExRequest, opts ...dcerpc.CallOption) (*ConnectionGetInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionGetInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ServerSetInfoEx(ctx context.Context, in *ServerSetInfoExRequest, opts ...dcerpc.CallOption) (*ServerSetInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerSetInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...dcerpc.CallOption) (*UpdateConnectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UpdateConnectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceSetCredentialsLocal(ctx context.Context, in *RouterInterfaceSetCredentialsLocalRequest, opts ...dcerpc.CallOption) (*RouterInterfaceSetCredentialsLocalResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceSetCredentialsLocalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceGetCredentialsLocal(ctx context.Context, in *RouterInterfaceGetCredentialsLocalRequest, opts ...dcerpc.CallOption) (*RouterInterfaceGetCredentialsLocalResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceGetCredentialsLocalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceGetCustomInfoEx(ctx context.Context, in *RouterInterfaceGetCustomInfoExRequest, opts ...dcerpc.CallOption) (*RouterInterfaceGetCustomInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceGetCustomInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceSetCustomInfoEx(ctx context.Context, in *RouterInterfaceSetCustomInfoExRequest, opts ...dcerpc.CallOption) (*RouterInterfaceSetCustomInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceSetCustomInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDimsvcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewDimsvcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DimsvcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DimsvcSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultDimsvcClient{cc: cc}, nil
}

// xxx_ServerGetInfoOperation structure represents the RMprAdminServerGetInfo operation
type xxx_ServerGetInfoOperation struct {
	Level  uint32                      `idl:"name:dwLevel" json:"level"`
	Info   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerGetInfoOperation) OpNum() int { return 0 }

func (o *xxx_ServerGetInfoOperation) OpName() string { return "/dimsvc/v0/RMprAdminServerGetInfo" }

func (o *xxx_ServerGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerGetInfoRequest structure represents the RMprAdminServerGetInfo operation request
type ServerGetInfoRequest struct {
	// dwLevel: This is of type DWORD and MUST be set to one of the following values.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | After the method returns, the memory pointed by pInfoStruct->pBuffer is          |
	//	|       | interpreted as MPR_SERVER_0 (section 2.2.1.2.61).                                |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | After the method returns, the memory pointed by pInfoStruct->pBuffer is          |
	//	|       | interpreted as MPR_SERVER_1 (section 2.2.1.2.62).<249>                           |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     2 | After the method returns, the memory pointed by pInfoStruct->pBuffer is          |
	//	|       | interpreted as MPR_SERVER_2 (section 2.2.1.2.63).<250>                           |
	//	+-------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
}

func (o *ServerGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerGetInfoOperation) *xxx_ServerGetInfoOperation {
	if op == nil {
		op = &xxx_ServerGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	return op
}

func (o *ServerGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
}
func (o *ServerGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerGetInfoResponse structure represents the RMprAdminServerGetInfo operation response
type ServerGetInfoResponse struct {
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER. DIM_INFORMATION_CONTAINER.dwBufferSize
	// SHOULD be initialized to zero (0). Upon successful return, the pInfoStruct->pBuffer
	// is typecast to MPR_SERVER_0, MPR_SERVER_1,<251> or MPR_SERVER_2<252> based on the
	// dwLevel value.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RMprAdminServerGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerGetInfoOperation) *xxx_ServerGetInfoOperation {
	if op == nil {
		op = &xxx_ServerGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *ServerGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *ServerGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionEnumOperation structure represents the RRasAdminConnectionEnum operation
type xxx_ConnectionEnumOperation struct {
	Level                  uint32                      `idl:"name:dwLevel" json:"level"`
	Info                   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                      `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	EntriesRead            uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries           uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume                 uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionEnumOperation) OpNum() int { return 1 }

func (o *xxx_ConnectionEnumOperation) OpName() string { return "/dimsvc/v0/RRasAdminConnectionEnum" }

func (o *xxx_ConnectionEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionEnumRequest structure represents the RRasAdminConnectionEnum operation request
type ConnectionEnumRequest struct {
	// dwLevel: This is of type DWORD and SHOULD be set to one of the following values.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as an array of RASI_CONNECTION_0 (section 2.2.1.2.77). The size of   |
	//	|       | the array is determined by lpdwEntriesRead.                                      |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as an array of RASI_CONNECTION_1 (section 2.2.1.2.78). The size of   |
	//	|       | the array is determined by lpdwEntriesRead.                                      |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     2 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as an array of RASI_CONNECTION_2 (section 2.2.1.2.79). The size of   |
	//	|       | the array is determined by lpdwEntriesRead.                                      |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     3 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as an array of RASI_CONNECTION_3 (section 2.2.1.2.80)).The size of   |
	//	|       | the array is determined by lpdwEntriesRead.                                      |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     4 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as an array of RAS_CONNECTION_4_IDL (section 2.2.1.2.236). The size  |
	//	|       | of the array is determined by lpdwEntriesRead.                                   |
	//	+-------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER, and DIM_INFORMATION_CONTAINER.dwBufferSize
	// is initialized to zero (0). Upon successful return, the pInfoStruct->pBuffer is a
	// typecast array of RASI_CONNECTION_0, RASI_CONNECTION_1, RASI_CONNECTION_2, RASI_CONNECTION_3,<253>
	// or RAS_CONNECTION_4_IDL based on the dwLevel value. The array size is determined
	// by the value in memory pointed to by lpdwEntriesRead.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// dwPreferedMaximumLength: This is of type DWORD and SHOULD specify the preferred maximum
	// length of the returned data (pInfoStruct->pBuffer) in bytes.
	PreferredMaximumLength uint32 `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	// lpdwResumeHandle: This is a pointer to type DWORD and specifies a resume handle that
	// is used to continue the enumeration. The lpdwResumeHandle parameter is zero (0) on
	// the first call and left unchanged on subsequent calls. The caller MUST pass the same
	// returned value in the next call to this function; otherwise, an error is returned.
	// If the return code is ERROR_MORE_DATA, another call can be made using this handle
	// to retrieve more data. If the return code is not ERROR_MORE_DATA, the handle returned
	// SHOULD be ignored. A return value of ERROR_SUCCESS indicates a successful completion
	// of the enumeration. Any return value other than ERROR_SUCCESS or ERROR_MORE_DATA
	// indicates the failure of the enumeration.
	Resume uint32 `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *ConnectionEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionEnumOperation) *xxx_ConnectionEnumOperation {
	if op == nil {
		op = &xxx_ConnectionEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.PreferredMaximumLength = o.PreferredMaximumLength
	op.Resume = o.Resume
	return op
}

func (o *ConnectionEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionEnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *ConnectionEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionEnumResponse structure represents the RRasAdminConnectionEnum operation response
type ConnectionEnumResponse struct {
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER, and DIM_INFORMATION_CONTAINER.dwBufferSize
	// is initialized to zero (0). Upon successful return, the pInfoStruct->pBuffer is a
	// typecast array of RASI_CONNECTION_0, RASI_CONNECTION_1, RASI_CONNECTION_2, RASI_CONNECTION_3,<253>
	// or RAS_CONNECTION_4_IDL based on the dwLevel value. The array size is determined
	// by the value in memory pointed to by lpdwEntriesRead.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// lpdwEntriesRead: This is a pointer to type DWORD and upon a successful function-call
	// return specifies the total number of connections enumerated from the current resume
	// position given by lpdwResumeHandle.
	EntriesRead uint32 `idl:"name:lpdwEntriesRead" json:"entries_read"`
	// lpdwTotalEntries: This is a pointer to type DWORD and receives the total number of
	// connections that could have been enumerated from the current resume position given
	// by lpdwResumeHandle.
	TotalEntries uint32 `idl:"name:lpdwTotalEntries" json:"total_entries"`
	// lpdwResumeHandle: This is a pointer to type DWORD and specifies a resume handle that
	// is used to continue the enumeration. The lpdwResumeHandle parameter is zero (0) on
	// the first call and left unchanged on subsequent calls. The caller MUST pass the same
	// returned value in the next call to this function; otherwise, an error is returned.
	// If the return code is ERROR_MORE_DATA, another call can be made using this handle
	// to retrieve more data. If the return code is not ERROR_MORE_DATA, the handle returned
	// SHOULD be ignored. A return value of ERROR_SUCCESS indicates a successful completion
	// of the enumeration. Any return value other than ERROR_SUCCESS or ERROR_MORE_DATA
	// indicates the failure of the enumeration.
	Resume uint32 `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRasAdminConnectionEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionEnumOperation) *xxx_ConnectionEnumOperation {
	if op == nil {
		op = &xxx_ConnectionEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.EntriesRead = o.EntriesRead
	op.TotalEntries = o.TotalEntries
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *ConnectionEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.EntriesRead = op.EntriesRead
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *ConnectionEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionGetInfoOperation structure represents the RRasAdminConnectionGetInfo operation
type xxx_ConnectionGetInfoOperation struct {
	Level      uint32                      `idl:"name:dwLevel" json:"level"`
	Connection uint32                      `idl:"name:hDimConnection" json:"connection"`
	Info       *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return     uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionGetInfoOperation) OpNum() int { return 2 }

func (o *xxx_ConnectionGetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionGetInfo"
}

func (o *xxx_ConnectionGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionGetInfoRequest structure represents the RRasAdminConnectionGetInfo operation request
type ConnectionGetInfoRequest struct {
	// dwLevel: This is of type DWORD and MUST be set to one of the following values.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as RASI_CONNECTION_0 (section 2.2.1.2.77).                           |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as RASI_CONNECTION_1 (section 2.2.1.2.78).                           |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     2 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as an array of RASI_CONNECTION_2 (section 2.2.1.2.79) structures.    |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     3 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as RASI_CONNECTION_3 (section 2.2.1.2.80).                           |
	//	+-------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// hDimConnection: This is of type DWORD and SHOULD be set to a particular connection
	// identifier for which the connection information is required. Obtain this handle by
	// calling RRasAdminConnectionEnum (section 3.1.4.2). Since RRASM server does not maintain
	// connection handle information, the validation of this handle SHOULD be done by the
	// RRAS server implementation.
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
}

func (o *ConnectionGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionGetInfoOperation) *xxx_ConnectionGetInfoOperation {
	if op == nil {
		op = &xxx_ConnectionGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Connection = o.Connection
	return op
}

func (o *ConnectionGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Connection = op.Connection
}
func (o *ConnectionGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionGetInfoResponse structure represents the RRasAdminConnectionGetInfo operation response
type ConnectionGetInfoResponse struct {
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER, and DIM_INFORMATION_CONTAINER.dwBufferSize
	// is initialized to zero (0). Upon successful return, the pInfoStruct->pBuffer is a
	// cast to an array of RASI_CONNECTION_0, RASI_CONNECTION_1, RASI_CONNECTION_2, or RASI_CONNECTION_3
	// structures, based on the dwLevel value.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRasAdminConnectionGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionGetInfoOperation) *xxx_ConnectionGetInfoOperation {
	if op == nil {
		op = &xxx_ConnectionGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *ConnectionGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *ConnectionGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionClearStatsOperation structure represents the RRasAdminConnectionClearStats operation
type xxx_ConnectionClearStatsOperation struct {
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionClearStatsOperation) OpNum() int { return 3 }

func (o *xxx_ConnectionClearStatsOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionClearStats"
}

func (o *xxx_ConnectionClearStatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionClearStatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionClearStatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionClearStatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionClearStatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ConnectionClearStatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionClearStatsRequest structure represents the RRasAdminConnectionClearStats operation request
type ConnectionClearStatsRequest struct {
	// hDimConnection: This is of type DWORD and SHOULD be set to the particular connection
	// identifier for which the connection statistics have to be cleared. Obtain this handle
	// by calling RRasAdminConnectionEnum (section 3.1.4.2). Because RRASM server does not
	// maintain connection handle information, the validation of this handle SHOULD be done
	// by the RRAS server implementation.
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
}

func (o *ConnectionClearStatsRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionClearStatsOperation) *xxx_ConnectionClearStatsOperation {
	if op == nil {
		op = &xxx_ConnectionClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Connection = o.Connection
	return op
}

func (o *ConnectionClearStatsRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionClearStatsOperation) {
	if o == nil {
		return
	}
	o.Connection = op.Connection
}
func (o *ConnectionClearStatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionClearStatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionClearStatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionClearStatsResponse structure represents the RRasAdminConnectionClearStats operation response
type ConnectionClearStatsResponse struct {
	// Return: The RRasAdminConnectionClearStats return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionClearStatsResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionClearStatsOperation) *xxx_ConnectionClearStatsOperation {
	if op == nil {
		op = &xxx_ConnectionClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ConnectionClearStatsResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionClearStatsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ConnectionClearStatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionClearStatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionClearStatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PortEnumOperation structure represents the RRasAdminPortEnum operation
type xxx_PortEnumOperation struct {
	Level                  uint32                      `idl:"name:dwLevel" json:"level"`
	Connection             uint32                      `idl:"name:hRasConnection" json:"connection"`
	Info                   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                      `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	EntriesRead            uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries           uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume                 uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_PortEnumOperation) OpNum() int { return 4 }

func (o *xxx_PortEnumOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortEnum" }

func (o *xxx_PortEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PortEnumRequest structure represents the RRasAdminPortEnum operation request
type PortEnumRequest struct {
	// dwLevel: This is of type DWORD and MUST be set to the following value.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | After the function returns, the memory pointed by pInfoStruct->pBuffer is        |
	//	|       | interpreted as an array of RASI_PORT_0 (section 2.2.1.2.75). The size of the     |
	//	|       | array is determined by lpdwEntriesRead. This includes information related to the |
	//	|       | Port object during runtime. The port objects SHOULD be initialized based on the  |
	//	|       | port configurations defined by WanEndpoints (section 2.2.3.3.1.8). The Device    |
	//	|       | Name is also returned as a part of each port information.                        |
	//	+-------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// hRasConnection: This is of type DWORD and SHOULD be set to a particular connection
	// identifier for which the connection information is required. Obtain this handle by
	// calling RRasAdminConnectionEnum (section 3.1.4.2). If this parameter is passed as
	// INVALID_HANDLE_VALUE (0xFFFFFFFF), the function enumerates all the active ports configured
	// on the RRAS server. Since RRASM server does not maintain connection handle information,
	// the validation of this handle SHOULD be done by the RRAS server implementation.
	Connection uint32 `idl:"name:hRasConnection" json:"connection"`
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1),
	// and DIM_INFORMATION_CONTAINER.dwBufferSize is initialized to zero (0). Upon successful
	// return, the pInfoStruct->pBuffer is typecast to an array of RASI_PORT_0, and the
	// array size is determined by the value to pointer lpdwEntriesRead.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// dwPreferedMaximumLength: This is of type DWORD and SHOULD specify the preferred maximum
	// length of returned data (pInfoStruct->pBuffer) in bytes. If dwPreferedMaximumLength
	// is -1 then all existing port entries for the specified connection are returned in
	// the buffer.
	PreferredMaximumLength uint32 `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	// lpdwResumeHandle: This is a pointer to type DWORD and specifies a resume handle that
	// is used to continue the enumeration. The lpdwResumeHandle parameter is zero (0) on
	// the first call and left unchanged on subsequent calls (the caller MUST pass the same
	// returned value in the next call to this function; otherwise, an error is returned).
	// If the return code is ERROR_MORE_DATA (0x000000EA), another call can be made using
	// this handle to retrieve more data. If the handle is NULL upon return, the enumeration
	// is complete. This handle is invalid (-1) for other types of error returns.
	Resume uint32 `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *PortEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_PortEnumOperation) *xxx_PortEnumOperation {
	if op == nil {
		op = &xxx_PortEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Connection = o.Connection
	op.Info = o.Info
	op.PreferredMaximumLength = o.PreferredMaximumLength
	op.Resume = o.Resume
	return op
}

func (o *PortEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_PortEnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Connection = op.Connection
	o.Info = op.Info
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *PortEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PortEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PortEnumResponse structure represents the RRasAdminPortEnum operation response
type PortEnumResponse struct {
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1),
	// and DIM_INFORMATION_CONTAINER.dwBufferSize is initialized to zero (0). Upon successful
	// return, the pInfoStruct->pBuffer is typecast to an array of RASI_PORT_0, and the
	// array size is determined by the value to pointer lpdwEntriesRead.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// lpdwEntriesRead: This is a pointer to type DWORD. Upon successful return, this determines
	// the total number of ports enumerated from the current resume position given by lpdwResumeHandle.
	EntriesRead uint32 `idl:"name:lpdwEntriesRead" json:"entries_read"`
	// lpdwTotalEntries: This is a pointer to type DWORD and receives the total number of
	// ports that could have been enumerated from the current resume position given by lpdwResumeHandle.
	TotalEntries uint32 `idl:"name:lpdwTotalEntries" json:"total_entries"`
	// lpdwResumeHandle: This is a pointer to type DWORD and specifies a resume handle that
	// is used to continue the enumeration. The lpdwResumeHandle parameter is zero (0) on
	// the first call and left unchanged on subsequent calls (the caller MUST pass the same
	// returned value in the next call to this function; otherwise, an error is returned).
	// If the return code is ERROR_MORE_DATA (0x000000EA), another call can be made using
	// this handle to retrieve more data. If the handle is NULL upon return, the enumeration
	// is complete. This handle is invalid (-1) for other types of error returns.
	Resume uint32 `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRasAdminPortEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PortEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_PortEnumOperation) *xxx_PortEnumOperation {
	if op == nil {
		op = &xxx_PortEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.EntriesRead = o.EntriesRead
	op.TotalEntries = o.TotalEntries
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *PortEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_PortEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.EntriesRead = op.EntriesRead
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *PortEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PortEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PortGetInfoOperation structure represents the RRasAdminPortGetInfo operation
type xxx_PortGetInfoOperation struct {
	Level  uint32                      `idl:"name:dwLevel" json:"level"`
	Port   uint32                      `idl:"name:hPort" json:"port"`
	Info   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_PortGetInfoOperation) OpNum() int { return 5 }

func (o *xxx_PortGetInfoOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortGetInfo" }

func (o *xxx_PortGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PortGetInfoRequest structure represents the RRasAdminPortGetInfo operation request
type PortGetInfoRequest struct {
	// dwLevel: This is of type DWORD and MUST be set to any of the following values.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | After the function returns, the memory pointed to by pInfoStruct->pBuffer        |
	//	|       | is typecast to RASI_PORT_0 (section 2.2.1.2.75). The size of the array is        |
	//	|       | determined by lpdwEntriesRead.                                                   |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | After the function returns, the memory pointed to by pInfoStruct->pBuffer        |
	//	|       | is typecast to RASI_ PORT_1 (section 2.2.1.2.76). The size of the array is       |
	//	|       | determined by lpdwEntriesRead.                                                   |
	//	+-------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// hPort: This is of type DWORD and MUST be set to the particular port identifier for
	// which the port information is required. Obtain this handle by calling RRasAdminPortEnum
	// (section 3.1.4.5). Since RRASM server does not maintain any port handle information,
	// the validation of this handle SHOULD be done by the RRAS server implementation.
	Port uint32 `idl:"name:hPort" json:"port"`
}

func (o *PortGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_PortGetInfoOperation) *xxx_PortGetInfoOperation {
	if op == nil {
		op = &xxx_PortGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Port = o.Port
	return op
}

func (o *PortGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_PortGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Port = op.Port
}
func (o *PortGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PortGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PortGetInfoResponse structure represents the RRasAdminPortGetInfo operation response
type PortGetInfoResponse struct {
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER, and DIM_INFORMATION_CONTAINER.dwBufferSize
	// is initialized to zero (0). Upon successful return, the pInfoStruct->pBuffer is typecast
	// to an array of the RASI_PORT_0 or RASI_ PORT_1 structures, based on the dwLevel value.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRasAdminPortGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PortGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_PortGetInfoOperation) *xxx_PortGetInfoOperation {
	if op == nil {
		op = &xxx_PortGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *PortGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_PortGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *PortGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PortGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PortClearStatsOperation structure represents the RRasAdminPortClearStats operation
type xxx_PortClearStatsOperation struct {
	Port   uint32 `idl:"name:hPort" json:"port"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_PortClearStatsOperation) OpNum() int { return 6 }

func (o *xxx_PortClearStatsOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortClearStats" }

func (o *xxx_PortClearStatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortClearStatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortClearStatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortClearStatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortClearStatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PortClearStatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PortClearStatsRequest structure represents the RRasAdminPortClearStats operation request
type PortClearStatsRequest struct {
	// hPort: This is of type DWORD and MUST be set to the particular port identifier for
	// which the port information is required. Obtain this handle by calling RRasAdminPortEnum
	// (section 3.1.4.5). Since RRASM server does not maintain port handle information,
	// the validation of this handle SHOULD be done by the RRAS server implementation.
	Port uint32 `idl:"name:hPort" json:"port"`
}

func (o *PortClearStatsRequest) xxx_ToOp(ctx context.Context, op *xxx_PortClearStatsOperation) *xxx_PortClearStatsOperation {
	if op == nil {
		op = &xxx_PortClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Port = o.Port
	return op
}

func (o *PortClearStatsRequest) xxx_FromOp(ctx context.Context, op *xxx_PortClearStatsOperation) {
	if o == nil {
		return
	}
	o.Port = op.Port
}
func (o *PortClearStatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PortClearStatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortClearStatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PortClearStatsResponse structure represents the RRasAdminPortClearStats operation response
type PortClearStatsResponse struct {
	// Return: The RRasAdminPortClearStats return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PortClearStatsResponse) xxx_ToOp(ctx context.Context, op *xxx_PortClearStatsOperation) *xxx_PortClearStatsOperation {
	if op == nil {
		op = &xxx_PortClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *PortClearStatsResponse) xxx_FromOp(ctx context.Context, op *xxx_PortClearStatsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PortClearStatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PortClearStatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortClearStatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PortResetOperation structure represents the RRasAdminPortReset operation
type xxx_PortResetOperation struct {
	Port   uint32 `idl:"name:hPort" json:"port"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_PortResetOperation) OpNum() int { return 7 }

func (o *xxx_PortResetOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortReset" }

func (o *xxx_PortResetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortResetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortResetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortResetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortResetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PortResetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PortResetRequest structure represents the RRasAdminPortReset operation request
type PortResetRequest struct {
	// hPort: This is of type DWORD and SHOULD be set to the particular port identifier
	// for which the port information is required. Obtain this handle by calling RRasAdminPortEnum
	// (section 3.1.4.5). This parameter is not used. Because RRASM server does not maintain
	// port handle information, the validation of this handle SHOULD be done by the RRAS
	// server implementation.
	Port uint32 `idl:"name:hPort" json:"port"`
}

func (o *PortResetRequest) xxx_ToOp(ctx context.Context, op *xxx_PortResetOperation) *xxx_PortResetOperation {
	if op == nil {
		op = &xxx_PortResetOperation{}
	}
	if o == nil {
		return op
	}
	op.Port = o.Port
	return op
}

func (o *PortResetRequest) xxx_FromOp(ctx context.Context, op *xxx_PortResetOperation) {
	if o == nil {
		return
	}
	o.Port = op.Port
}
func (o *PortResetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PortResetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortResetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PortResetResponse structure represents the RRasAdminPortReset operation response
type PortResetResponse struct {
	// Return: The RRasAdminPortReset return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PortResetResponse) xxx_ToOp(ctx context.Context, op *xxx_PortResetOperation) *xxx_PortResetOperation {
	if op == nil {
		op = &xxx_PortResetOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *PortResetResponse) xxx_FromOp(ctx context.Context, op *xxx_PortResetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PortResetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PortResetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortResetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PortDisconnectOperation structure represents the RRasAdminPortDisconnect operation
type xxx_PortDisconnectOperation struct {
	Port   uint32 `idl:"name:hPort" json:"port"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_PortDisconnectOperation) OpNum() int { return 8 }

func (o *xxx_PortDisconnectOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortDisconnect" }

func (o *xxx_PortDisconnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortDisconnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortDisconnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortDisconnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortDisconnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PortDisconnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PortDisconnectRequest structure represents the RRasAdminPortDisconnect operation request
type PortDisconnectRequest struct {
	// hPort: This is of type DWORD and MUST be set to the port identifier of the port that
	// is to be disconnected. Obtain this handle by calling RRasAdminPortEnum (section 3.1.4.5).
	// Since RRASM server does not maintain port handle information, the validation of this
	// handle SHOULD be done by the RRAS server implementation.
	Port uint32 `idl:"name:hPort" json:"port"`
}

func (o *PortDisconnectRequest) xxx_ToOp(ctx context.Context, op *xxx_PortDisconnectOperation) *xxx_PortDisconnectOperation {
	if op == nil {
		op = &xxx_PortDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Port = o.Port
	return op
}

func (o *PortDisconnectRequest) xxx_FromOp(ctx context.Context, op *xxx_PortDisconnectOperation) {
	if o == nil {
		return
	}
	o.Port = op.Port
}
func (o *PortDisconnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PortDisconnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortDisconnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PortDisconnectResponse structure represents the RRasAdminPortDisconnect operation response
type PortDisconnectResponse struct {
	// Return: The RRasAdminPortDisconnect return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PortDisconnectResponse) xxx_ToOp(ctx context.Context, op *xxx_PortDisconnectOperation) *xxx_PortDisconnectOperation {
	if op == nil {
		op = &xxx_PortDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *PortDisconnectResponse) xxx_FromOp(ctx context.Context, op *xxx_PortDisconnectOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PortDisconnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PortDisconnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortDisconnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportSetGlobalInfoOperation structure represents the RRouterInterfaceTransportSetGlobalInfo operation
type xxx_RouterInterfaceTransportSetGlobalInfoOperation struct {
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) OpNum() int { return 9 }

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportSetGlobalInfo"
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportSetGlobalInfoRequest structure represents the RRouterInterfaceTransportSetGlobalInfo operation request
type RouterInterfaceTransportSetGlobalInfoRequest struct {
	// dwTransportId: Specifies the transport for which the information is set (IPX, IPv4,
	// or IPv6). It MUST be one of the following values.<255>
	//
	//	+---------------------+---------------+
	//	|                     |               |
	//	|        VALUE        |    MEANING    |
	//	|                     |               |
	//	+---------------------+---------------+
	//	+---------------------+---------------+
	//	| PID_IPX 0x0000002B  | IPX protocol  |
	//	+---------------------+---------------+
	//	| PID_IP 0x00000021   | IPv4 protocol |
	//	+---------------------+---------------+
	//	| PID_IPV6 0x00000057 | IPv6 protocol |
	//	+---------------------+---------------+
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
	// pInfoStruct: Specifies the pointer to a DIM_INTERFACE_CONTAINER (section 2.2.1.2.2).
	//
	// The fGetGlobalInfo member of the DIM_INTERFACE_CONTAINER MUST be set to 0.
	//
	// The pGlobalInfo and dwGlobalInfoSize of DIM_INTERFACE_CONTAINER MUST be set. The
	// rest of the fields SHOULD not be set.
	//
	// The dwGlobalInfoSize field MUST be set to the size of the information passed in pGlobalInfo.
	Info *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportSetGlobalInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetGlobalInfoOperation) *xxx_RouterInterfaceTransportSetGlobalInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportSetGlobalInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportSetGlobalInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetGlobalInfoOperation) {
	if o == nil {
		return
	}
	o.TransportID = op.TransportID
	o.Info = op.Info
}
func (o *RouterInterfaceTransportSetGlobalInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportSetGlobalInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportSetGlobalInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportSetGlobalInfoResponse structure represents the RRouterInterfaceTransportSetGlobalInfo operation response
type RouterInterfaceTransportSetGlobalInfoResponse struct {
	// Return: The RRouterInterfaceTransportSetGlobalInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportSetGlobalInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetGlobalInfoOperation) *xxx_RouterInterfaceTransportSetGlobalInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportSetGlobalInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportSetGlobalInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetGlobalInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceTransportSetGlobalInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportSetGlobalInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportSetGlobalInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportGetGlobalInfoOperation structure represents the RRouterInterfaceTransportGetGlobalInfo operation
type xxx_RouterInterfaceTransportGetGlobalInfoOperation struct {
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) OpNum() int { return 10 }

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportGetGlobalInfo"
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportGetGlobalInfoRequest structure represents the RRouterInterfaceTransportGetGlobalInfo operation request
type RouterInterfaceTransportGetGlobalInfoRequest struct {
	// dwTransportId: Specifies the transport for which the information is set (IPX, IPv4,
	// or IPv6).  It MUST be one of the following values.<257>
	//
	//	+---------------------+---------------+
	//	|                     |               |
	//	|        VALUE        |    MEANING    |
	//	|                     |               |
	//	+---------------------+---------------+
	//	+---------------------+---------------+
	//	| PID_IPX 0x0000002B  | IPX protocol  |
	//	+---------------------+---------------+
	//	| PID_IP 0x00000021   | IPv4 protocol |
	//	+---------------------+---------------+
	//	| PID_IPV6 0x00000057 | IPv6 protocol |
	//	+---------------------+---------------+
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
	// pInfoStruct: Pointer to DIM_INTERFACE_CONTAINER (section 2.2.1.2.1). This pointer
	// MUST NOT be NULL when calling this method and is allocated to the size of the DIM_INTERFACE_CONTAINER
	// structure. On successful completion, the structure is populated with a DIM_INTERFACE_CONTAINER
	// structure having valid values for dwGlobalInfoSize and pGlobalInfo fields which the
	// caller SHOULD free when done using it.
	//
	// The fGetGlobalInfo of DIM_INTERFACE_CONTAINER MUST be set to 1. The rest of the DIM_INTERFACE_CONTAINER
	// fields SHOULD not be set.
	//
	// The pGlobalInfo and dwGlobalInfoSize members of DIM_INTERFACE_CONTAINER will be populated
	// on successful completion of this method.
	//
	// The pGlobalInfo member will point to a valid RTR_INFO_BLOCK_HEADER (section 2.2.1.2.3)
	// and RTR_TOC_ENTRY (section 2.2.1.2.4).
	//
	// On successful completion dwGlobalInfoSize will be set to the size of the information
	// being passed in pGlobalInfo.
	//
	// If dwTransportId is PID_IP or PID_IP6 and if the InfoType field in the RTR_TOC_ENTRY
	// structure is one of the following, pGlobalInfo MUST be interpreted as the structure
	// in the following table.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|            INFOTYPE             |                                                                                  |                             STRUCTURE POINTED TO BY                              |
	//	|              VALUE              |                                     MEANING                                      |                                   PGLOBALINFO                                    |
	//	|                                 |                                                                                  |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_BOOTP 0x0000270F          | IP BOOTP global configuration information.                                       | IPBOOTP_GLOBAL_CONFIG (section 2.2.1.2.149)                                      |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_IGMP 0x4137000A           | IGMP global configuration.                                                       | IGMP_MIB_GLOBAL_CONFIG (section 2.2.1.2.173)                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_RIP 0x00000008            | IP RIP global configuration information.                                         | IPRIP_GLOBAL_CONFIG (section 2.2.1.2.164) (values specified are overwritten)     |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_BGP 0x0137000E            | BGP configuration.                                                               | When dwTransportId is PID_IP, the structure pointed to, at the offset is         |
	//	|                                 |                                                                                  | BGP_CONFIG_HEADER (section 2.2.1.2.252). When dwTransportId is PID_IP6, the      |
	//	|                                 |                                                                                  | structure pointed to, at the offset is BGP_ROUTER_V6 (section 2.2.1.2.265).      |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_DHCP_ALLOCATOR 0x81372714 | DHCP global configuration information.                                           | IP_AUTO_DHCP_GLOBAL_INFO (section 2.2.1.2.191) (values specified are             |
	//	|                                 |                                                                                  | overwritten)                                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_DNS_PROXY 0x81372713      | DNS proxy global configuration information.                                      | IP_DNS_PROXY_GLOBAL_INFO (section 2.2.1.2.193) (values specified are             |
	//	|                                 |                                                                                  | overwritten)                                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_OSPF<258> 0x0000000D      | OSPF global configuration information specified is retrieved. The type field     | OSPF_ROUTE_FILTER_INFO (section 2.2.1.2.209) OSPF_PROTO_FILTER_INFO (section     |
	//	|                                 | is the first field in each of the structures and this defines which of the       | 2.2.1.2.210) OSPF_GLOBAL_PARAM (section 2.2.1.2.211) OSPF_ROUTE_FILTER (section  |
	//	|                                 | structures MUST be used. The type field MUST be OSPF_PARAM_TYPE (section         | 2.2.1.2.208)                                                                     |
	//	|                                 | 2.2.1.1.14) and the value MUST be corresponding to the structures specified.     |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_ALG<259> 0x8137271A       | Application layer gateway global configuration.                                  | IP_ALG_GLOBAL_INFO (section 2.2.1.2.201)                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IPV6_DHCP 0x000003E7         | DHCPv6 Relay global configuration information.                                   | DHCPV6R_GLOBAL_CONFIG (section 2.2.1.2.157)                                      |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_NAT 0x81372715            | IP NAT global configuration information.                                         | IP_NAT_GLOBAL_INFO (section 2.2.1.2.195)                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	Info *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportGetGlobalInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetGlobalInfoOperation) *xxx_RouterInterfaceTransportGetGlobalInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportGetGlobalInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportGetGlobalInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetGlobalInfoOperation) {
	if o == nil {
		return
	}
	o.TransportID = op.TransportID
	o.Info = op.Info
}
func (o *RouterInterfaceTransportGetGlobalInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportGetGlobalInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportGetGlobalInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportGetGlobalInfoResponse structure represents the RRouterInterfaceTransportGetGlobalInfo operation response
type RouterInterfaceTransportGetGlobalInfoResponse struct {
	// pInfoStruct: Pointer to DIM_INTERFACE_CONTAINER (section 2.2.1.2.1). This pointer
	// MUST NOT be NULL when calling this method and is allocated to the size of the DIM_INTERFACE_CONTAINER
	// structure. On successful completion, the structure is populated with a DIM_INTERFACE_CONTAINER
	// structure having valid values for dwGlobalInfoSize and pGlobalInfo fields which the
	// caller SHOULD free when done using it.
	//
	// The fGetGlobalInfo of DIM_INTERFACE_CONTAINER MUST be set to 1. The rest of the DIM_INTERFACE_CONTAINER
	// fields SHOULD not be set.
	//
	// The pGlobalInfo and dwGlobalInfoSize members of DIM_INTERFACE_CONTAINER will be populated
	// on successful completion of this method.
	//
	// The pGlobalInfo member will point to a valid RTR_INFO_BLOCK_HEADER (section 2.2.1.2.3)
	// and RTR_TOC_ENTRY (section 2.2.1.2.4).
	//
	// On successful completion dwGlobalInfoSize will be set to the size of the information
	// being passed in pGlobalInfo.
	//
	// If dwTransportId is PID_IP or PID_IP6 and if the InfoType field in the RTR_TOC_ENTRY
	// structure is one of the following, pGlobalInfo MUST be interpreted as the structure
	// in the following table.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|            INFOTYPE             |                                                                                  |                             STRUCTURE POINTED TO BY                              |
	//	|              VALUE              |                                     MEANING                                      |                                   PGLOBALINFO                                    |
	//	|                                 |                                                                                  |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_BOOTP 0x0000270F          | IP BOOTP global configuration information.                                       | IPBOOTP_GLOBAL_CONFIG (section 2.2.1.2.149)                                      |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_IGMP 0x4137000A           | IGMP global configuration.                                                       | IGMP_MIB_GLOBAL_CONFIG (section 2.2.1.2.173)                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_RIP 0x00000008            | IP RIP global configuration information.                                         | IPRIP_GLOBAL_CONFIG (section 2.2.1.2.164) (values specified are overwritten)     |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_BGP 0x0137000E            | BGP configuration.                                                               | When dwTransportId is PID_IP, the structure pointed to, at the offset is         |
	//	|                                 |                                                                                  | BGP_CONFIG_HEADER (section 2.2.1.2.252). When dwTransportId is PID_IP6, the      |
	//	|                                 |                                                                                  | structure pointed to, at the offset is BGP_ROUTER_V6 (section 2.2.1.2.265).      |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_DHCP_ALLOCATOR 0x81372714 | DHCP global configuration information.                                           | IP_AUTO_DHCP_GLOBAL_INFO (section 2.2.1.2.191) (values specified are             |
	//	|                                 |                                                                                  | overwritten)                                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_DNS_PROXY 0x81372713      | DNS proxy global configuration information.                                      | IP_DNS_PROXY_GLOBAL_INFO (section 2.2.1.2.193) (values specified are             |
	//	|                                 |                                                                                  | overwritten)                                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_OSPF<258> 0x0000000D      | OSPF global configuration information specified is retrieved. The type field     | OSPF_ROUTE_FILTER_INFO (section 2.2.1.2.209) OSPF_PROTO_FILTER_INFO (section     |
	//	|                                 | is the first field in each of the structures and this defines which of the       | 2.2.1.2.210) OSPF_GLOBAL_PARAM (section 2.2.1.2.211) OSPF_ROUTE_FILTER (section  |
	//	|                                 | structures MUST be used. The type field MUST be OSPF_PARAM_TYPE (section         | 2.2.1.2.208)                                                                     |
	//	|                                 | 2.2.1.1.14) and the value MUST be corresponding to the structures specified.     |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_ALG<259> 0x8137271A       | Application layer gateway global configuration.                                  | IP_ALG_GLOBAL_INFO (section 2.2.1.2.201)                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IPV6_DHCP 0x000003E7         | DHCPv6 Relay global configuration information.                                   | DHCPV6R_GLOBAL_CONFIG (section 2.2.1.2.157)                                      |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_NAT 0x81372715            | IP NAT global configuration information.                                         | IP_NAT_GLOBAL_INFO (section 2.2.1.2.195)                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	Info *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRouterInterfaceTransportGetGlobalInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportGetGlobalInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetGlobalInfoOperation) *xxx_RouterInterfaceTransportGetGlobalInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportGetGlobalInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportGetGlobalInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetGlobalInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RouterInterfaceTransportGetGlobalInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportGetGlobalInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportGetGlobalInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceGetHandleOperation structure represents the RRouterInterfaceGetHandle operation
type xxx_RouterInterfaceGetHandleOperation struct {
	InterfaceName           string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
	Interface               uint32 `idl:"name:phInterface" json:"interface"`
	IncludeClientInterfaces uint32 `idl:"name:fIncludeClientInterfaces" json:"include_client_interfaces"`
	Return                  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceGetHandleOperation) OpNum() int { return 11 }

func (o *xxx_RouterInterfaceGetHandleOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceGetHandle"
}

func (o *xxx_RouterInterfaceGetHandleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetHandleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.InterfaceName); err != nil {
			return err
		}
	}
	// phInterface {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// fIncludeClientInterfaces {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.IncludeClientInterfaces); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetHandleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.InterfaceName); err != nil {
			return err
		}
	}
	// phInterface {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// fIncludeClientInterfaces {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.IncludeClientInterfaces); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetHandleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetHandleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phInterface {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetHandleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phInterface {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceGetHandleRequest structure represents the RRouterInterfaceGetHandle operation request
type RouterInterfaceGetHandleRequest struct {
	// lpwsInterfaceName: Pointer to a null-terminated Unicode string that specifies the
	// name of the interface to be retrieved.
	InterfaceName string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
	// phInterface: This is a pointer to a DWORD that receives the unique identifier of
	// the interface specified by lpwsInterfaceName.
	Interface uint32 `idl:"name:phInterface" json:"interface"`
	// fIncludeClientInterfaces: Specifies whether the method includes client interfaces
	// while searching. If this parameter is 0, interfaces of type ROUTER_IF_TYPE_CLIENT
	// are ignored in the search for the interface with the name specified by lpwsInterfaceName.
	// If this parameter is a nonzero value and an interface with the specified name exists,
	// RRouterInterfaceGetHandle returns a handle to an interface of type ROUTER_IF_TYPE_CLIENT.
	// Since it is possible that there are several interfaces of type ROUTER_IF_TYPE_CLIENT,
	// the handle returned references the first interface that is found with the name ROUTER_IF_TYPE_CLIENT
	// specified by lpwsInterfaceName.
	IncludeClientInterfaces uint32 `idl:"name:fIncludeClientInterfaces" json:"include_client_interfaces"`
}

func (o *RouterInterfaceGetHandleRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetHandleOperation) *xxx_RouterInterfaceGetHandleOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceName = o.InterfaceName
	op.Interface = o.Interface
	op.IncludeClientInterfaces = o.IncludeClientInterfaces
	return op
}

func (o *RouterInterfaceGetHandleRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetHandleOperation) {
	if o == nil {
		return
	}
	o.InterfaceName = op.InterfaceName
	o.Interface = op.Interface
	o.IncludeClientInterfaces = op.IncludeClientInterfaces
}
func (o *RouterInterfaceGetHandleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceGetHandleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetHandleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceGetHandleResponse structure represents the RRouterInterfaceGetHandle operation response
type RouterInterfaceGetHandleResponse struct {
	// phInterface: This is a pointer to a DWORD that receives the unique identifier of
	// the interface specified by lpwsInterfaceName.
	Interface uint32 `idl:"name:phInterface" json:"interface"`
	// Return: The RRouterInterfaceGetHandle return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceGetHandleResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetHandleOperation) *xxx_RouterInterfaceGetHandleOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetHandleResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetHandleOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.Return = op.Return
}
func (o *RouterInterfaceGetHandleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceGetHandleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetHandleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceCreateOperation structure represents the RRouterInterfaceCreate operation
type xxx_RouterInterfaceCreateOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:phInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceCreateOperation) OpNum() int { return 12 }

func (o *xxx_RouterInterfaceCreateOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceCreate"
}

func (o *xxx_RouterInterfaceCreateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceCreateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// phInterface {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceCreateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// phInterface {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceCreateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceCreateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phInterface {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceCreateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phInterface {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceCreateRequest structure represents the RRouterInterfaceCreate operation request
type RouterInterfaceCreateRequest struct {
	// dwLevel: This is of type DWORD and MUST be set to any of the following values.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | The pInfoStruct->pBuffer MUST be set to point to a MPRI_INTERFACE_0 (section     |
	//	|       | 2.2.1.2.81).                                                                     |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | The pInfoStruct->pBuffer MUST be set to point to a MPRI_INTERFACE_1 (section     |
	//	|       | 2.2.1.2.82).                                                                     |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     2 | The pInfoStruct->pBuffer MUST be set to point to a MPRI_INTERFACE_2 (section     |
	//	|       | 2.2.1.2.83).                                                                     |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     3 | The pInfoStruct->pBuffer MUST be set to point to a MPRI_INTERFACE_3 (section     |
	//	|       | 2.2.1.2.84).                                                                     |
	//	+-------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1)
	// and MUST be set to following, based on dwLevel.
	//
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |                                                                                  |
	//	| DWLEVEL |                                PINFOSTRUCT->SIZE                                 |                               PINFOSTRUCT->PBUFFER                               |
	//	|         |                                                                                  |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|       0 | This MUST be set to the size of MPRI_INTERFACE_0 (section 2.2.1.2.81).           | This MUST be set to point to MPRI_INTERFACE_0. Only the wszInterfaceName,        |
	//	|         | Otherwise, an error is returned.                                                 | fEnabled, and dwIfType fields of MPRI_INTERFACE_0 can be set. Setting other      |
	//	|         |                                                                                  | values has no effect.                                                            |
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|       1 | This MUST be set to the size of MPRI_INTERFACE_1 (section 2.2.1.2.82).           | This MUST be set to point to MPRI_INTERFACE_1. Only the wszInterfaceName,        |
	//	|         | Otherwise, an error is returned.                                                 | fEnabled, dwIfType fields of MPRI_INTERFACE_1 can be set. Setting other values   |
	//	|         |                                                                                  | has no effect.                                                                   |
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|       2 | This MUST be set to the size of MPRI_INTERFACE_2 (section 2.2.1.2.83).           | This MUST be set to point to MPRI_INTERFACE_2. The dwIfType of MPRI_INTERFACE_2  |
	//	|         | Otherwise, an error is returned.                                                 | SHOULD be set to ROUTER_IF_TYPE_FULL_ROUTER. The szAlternates of                 |
	//	|         |                                                                                  | MPRI_INTERFACE_2 MUST be set to 0.                                               |
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|       3 | This MUST be set to the size of MPRI_INTERFACE_3 (section 2.2.1.2.84).           | This MUST be set to point to MPRI_INTERFACE_3. The dwIfType of MPRI_INTERFACE_3  |
	//	|         | Otherwise, an error is returned.                                                 | SHOULD be set to ROUTER_IF_TYPE_FULL_ROUTER. The szAlternates of                 |
	//	|         |                                                                                  | MPRI_INTERFACE_3 MUST be set to 0. The values of the ipv6addrDns and             |
	//	|         |                                                                                  | ipv6addrDnsAlt members of the MPRI_INTERFACE_3 structure SHOULD be ignored by    |
	//	|         |                                                                                  | the server.                                                                      |
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// If dwIfType is set to ROUTER_IF_TYPE_TUNNEL1 or ROUTER_IF_TYPE_DIALOUT in MPRI_INTERFACE_0,
	// MPRI_INTERFACE_1, MPRI_INTERFACE_2, or MPRI_INTERFACE_3, an error is returned.<262>
	//
	// If dwIfType is set to ROUTER_IF_TYPE_DEDICATED, ROUTER_IF_TYPE_INTERNAL, or ROUTER_IF_TYPE_LOOPBACK,
	// and if fEnabled is set to FALSE, an error other than one of the errors specified
	// in the table that follows MUST be returned.
	//
	// If the dwIfType in MPRI_INTERFACE_0, MPRI_INTERFACE_1, MPRI_INTERFACE_2, or MPRI_INTERFACE_3,
	// is set to ROUTER_IF_TYPE_CLIENT, ROUTER_IF_TYPE_HOME_ROUTER, or ROUTER_IF_TYPE_FULL_ROUTER,
	// and if the RouterType is ROUTER_TYPE_LAN, an error SHOULD be returned.
	//
	// If dwIfType in MPRI_INTERFACE_0 is set to ROUTER_IF_TYPE_FULL_ROUTER, phonebook information
	// for the interface MUST have already been configured in the phonebook file.
	//
	// * MPRIO_RequireEncryptedPw
	//
	// * MPRIO_RequireDataEncryption
	//
	// * MPRIO_RequireCHAP
	//
	// * MPRIO_RequireMsCHAP2
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// phInterface: This is a pointer to a DWORD that specifies the unique identifier of
	// the interface that is created. This is the same as the dwInterface in MPRI_INTERFACE_0,
	// MPRI_INTERFACE_1, MPRI_INTERFACE_2, or MPRI_INTERFACE_3.
	Interface uint32 `idl:"name:phInterface" json:"interface"`
}

func (o *RouterInterfaceCreateRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceCreateOperation) *xxx_RouterInterfaceCreateOperation {
	if op == nil {
		op = &xxx_RouterInterfaceCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceCreateRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceCreateOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Interface = op.Interface
}
func (o *RouterInterfaceCreateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceCreateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceCreateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceCreateResponse structure represents the RRouterInterfaceCreate operation response
type RouterInterfaceCreateResponse struct {
	// phInterface: This is a pointer to a DWORD that specifies the unique identifier of
	// the interface that is created. This is the same as the dwInterface in MPRI_INTERFACE_0,
	// MPRI_INTERFACE_1, MPRI_INTERFACE_2, or MPRI_INTERFACE_3.
	Interface uint32 `idl:"name:phInterface" json:"interface"`
	// Return: The RRouterInterfaceCreate return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceCreateResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceCreateOperation) *xxx_RouterInterfaceCreateOperation {
	if op == nil {
		op = &xxx_RouterInterfaceCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceCreateResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceCreateOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.Return = op.Return
}
func (o *RouterInterfaceCreateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceCreateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceCreateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceGetInfoOperation structure represents the RRouterInterfaceGetInfo operation
type xxx_RouterInterfaceGetInfoOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceGetInfoOperation) OpNum() int { return 13 }

func (o *xxx_RouterInterfaceGetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceGetInfo"
}

func (o *xxx_RouterInterfaceGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceGetInfoRequest structure represents the RRouterInterfaceGetInfo operation request
type RouterInterfaceGetInfoRequest struct {
	// dwLevel: This is of type DWORD and MUST be set to any of the following values.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as MPRI_INTERFACE_0 (section 2.2.1.2.81).                            |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as MPRI_INTERFACE_1 (section 2.2.1.2.82).                            |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     2 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as MPRI_INTERFACE_2 (section 2.2.1.2.83).                            |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     3 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as MPRI_INTERFACE_3 (section 2.2.1.2.84).                            |
	//	+-------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1),
	// and DIM_INFORMATION_CONTAINER.dwBufferSize SHOULD be initialized to zero (0). Upon
	// successful return, the pInfoStruct->pBuffer is cast to an array of MPRI_INTERFACE_0,
	// MPRI_INTERFACE_1, MPRI_INTERFACE_2, or MPRI_INTERFACE_3 structures based on the dwLevel
	// values.
	//
	// * If the *dwfOptions* member has neither MPRIO_RequirePAP nor MPRIO_RequireEAP, then
	// enable MPRIO_RequireEncryptedPw.
	//
	// * If the *dwfOptions* has none of the following flags set, then enable MPRIO_RequireMsEncryptedPw:
	//
	// * MPRIO_RequireCHAP
	//
	// * MPRIO_RequirePAP
	//
	// * MPRIO_RequireEAP
	//
	// * If the *dwEncryptionType* member is not set to MPR_ET_None or MPR_ET_Optional,
	// enable the MPRIO_RequireDataEncryption flag.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). The interface
	// MUST be of type ROUTER_IF_TYPE_FULL_ROUTER if dwLevel is set to 2 or 3. Because the
	// RRASM server does not maintain the interface handles, the RRAS server SHOULD check
	// and ensure that this handle value represents a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetInfoOperation) *xxx_RouterInterfaceGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Interface = op.Interface
}
func (o *RouterInterfaceGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceGetInfoResponse structure represents the RRouterInterfaceGetInfo operation response
type RouterInterfaceGetInfoResponse struct {
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1),
	// and DIM_INFORMATION_CONTAINER.dwBufferSize SHOULD be initialized to zero (0). Upon
	// successful return, the pInfoStruct->pBuffer is cast to an array of MPRI_INTERFACE_0,
	// MPRI_INTERFACE_1, MPRI_INTERFACE_2, or MPRI_INTERFACE_3 structures based on the dwLevel
	// values.
	//
	// * If the *dwfOptions* member has neither MPRIO_RequirePAP nor MPRIO_RequireEAP, then
	// enable MPRIO_RequireEncryptedPw.
	//
	// * If the *dwfOptions* has none of the following flags set, then enable MPRIO_RequireMsEncryptedPw:
	//
	// * MPRIO_RequireCHAP
	//
	// * MPRIO_RequirePAP
	//
	// * MPRIO_RequireEAP
	//
	// * If the *dwEncryptionType* member is not set to MPR_ET_None or MPR_ET_Optional,
	// enable the MPRIO_RequireDataEncryption flag.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRouterInterfaceGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetInfoOperation) *xxx_RouterInterfaceGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RouterInterfaceGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceSetInfoOperation structure represents the RRouterInterfaceSetInfo operation
type xxx_RouterInterfaceSetInfoOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceSetInfoOperation) OpNum() int { return 14 }

func (o *xxx_RouterInterfaceSetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceSetInfo"
}

func (o *xxx_RouterInterfaceSetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceSetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceSetInfoRequest structure represents the RRouterInterfaceSetInfo operation request
type RouterInterfaceSetInfoRequest struct {
	// dwLevel: This is of type DWORD and SHOULD be set to any of the following values.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | pInfoStruct->pBuffer MUST be set to point to a MPRI_INTERFACE_0 (section         |
	//	|       | 2.2.1.2.81).                                                                     |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | pInfoStruct->pBuffer MUST be set to point to a MPRI_INTERFACE_1 (section         |
	//	|       | 2.2.1.2.82).                                                                     |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     2 | pInfoStruct->pBuffer MUST be set to point to a MPRI_INTERFACE_2 (section         |
	//	|       | 2.2.1.2.83).                                                                     |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     3 | pInfoStruct->pBuffer MUST be set to point to a MPRI_INTERFACE_3 (section         |
	//	|       | 2.2.1.2.84).<265>                                                                |
	//	+-------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1)
	// and MUST be set to the following, based on dwLevel.
	//
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |                                                                                  |
	//	| DWLEVEL |                                PINFOSTRUCT->SIZE                                 |                               PINFOSTRUCT->PBUFFER                               |
	//	|         |                                                                                  |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|       0 | This MUST be set to the size of the data being passed in MPRI_INTERFACE_0        | This MUST be set to point to MPRI_INTERFACE_0. Only the fEnabled field of this   |
	//	|         | (section 2.2.1.2.81).                                                            | structure can be modified; the rest of the fields MUST be populated as returned  |
	//	|         |                                                                                  | by RRouterInterfaceGetInfo (section 3.1.4.14).                                   |
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|       1 | This MUST be set to the size of the data being passed in MPRI_INTERFACE_1        | This MUST be set to point to MPRI_INTERFACE_1. In addition to the fields         |
	//	|         | (section 2.2.1.2.82).                                                            | for MPRI_INTERFACE_0, lpwsDialoutHoursRestriction can be set. Setting            |
	//	|         |                                                                                  | other values has no effect. The values MUST be populated as returned by          |
	//	|         |                                                                                  | RRouterInterfaceGetInfo.                                                         |
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|       2 | This MUST be set to the size of MPRI_INTERFACE_2 (section 2.2.1.2.83).           | This MUST be set to point to MPRI_INTERFACE_2. The dwIfType of MPRI_INTERFACE_0  |
	//	|         | Otherwise, an error is returned.                                                 | MUST be set to ROUTER_IF_TYPE_FULL_ROUTER. The szAlternates of MPRI_INTERFACE_2  |
	//	|         |                                                                                  | MUST be set to 0.                                                                |
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|       3 | This MUST be set to the size of MPRI_INTERFACE_3 (section 2.2.1.2.84).           | This MUST be set to point to MPRI_INTERFACE_3. The dwIfType of MPRI_INTERFACE_3  |
	//	|         | Otherwise, an error is returned.                                                 | MUST be set to ROUTER_IF_TYPE_FULL_ROUTER. The szAlternates of MPRI_INTERFACE_0  |
	//	|         |                                                                                  | MUST be set to 0. The values of the ipv6addrDns and ipv6addrDnsAlt members of    |
	//	|         |                                                                                  | MPRI_INTERFACE_3 SHOULD be ignored by the server.                                |
	//	+---------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// If the interface type, specified at the time the interface was created using RRouterInterfaceCreate
	// (section 3.1.4.13), is either ROUTER_IF_TYPE_DEDICATED or ROUTER_IF_TYPE_INTERNAL
	// and if fEnabled is set to FALSE, an error is returned.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// or RRouterInterfaceGetHandle (section 3.1.4.12). Since the RRASM does not manage
	// the interface handles, the RRAS server SHOULD check to ensure that this handle is
	// a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceSetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetInfoOperation) *xxx_RouterInterfaceSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Interface = op.Interface
}
func (o *RouterInterfaceSetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceSetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceSetInfoResponse structure represents the RRouterInterfaceSetInfo operation response
type RouterInterfaceSetInfoResponse struct {
	// Return: The RRouterInterfaceSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceSetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetInfoOperation) *xxx_RouterInterfaceSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceSetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceSetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceSetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceDeleteOperation structure represents the RRouterInterfaceDelete operation
type xxx_RouterInterfaceDeleteOperation struct {
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceDeleteOperation) OpNum() int { return 15 }

func (o *xxx_RouterInterfaceDeleteOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceDelete"
}

func (o *xxx_RouterInterfaceDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceDeleteRequest structure represents the RRouterInterfaceDelete operation request
type RouterInterfaceDeleteRequest struct {
	// hInterface: A unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// or RRouterInterfaceGetHandle (section 3.1.4.12). Because the RRASM server does not
	// maintain the interface handles, the RRAS server SHOULD check to ensure that this
	// handle is a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceDeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeleteOperation) *xxx_RouterInterfaceDeleteOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeleteOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
}
func (o *RouterInterfaceDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceDeleteResponse structure represents the RRouterInterfaceDelete operation response
type RouterInterfaceDeleteResponse struct {
	// Return: The RRouterInterfaceDelete return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceDeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeleteOperation) *xxx_RouterInterfaceDeleteOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportRemoveOperation structure represents the RRouterInterfaceTransportRemove operation
type xxx_RouterInterfaceTransportRemoveOperation struct {
	Interface   uint32 `idl:"name:hInterface" json:"interface"`
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) OpNum() int { return 16 }

func (o *xxx_RouterInterfaceTransportRemoveOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportRemove"
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceTransportRemoveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportRemoveRequest structure represents the RRouterInterfaceTransportRemove operation request
type RouterInterfaceTransportRemoveRequest struct {
	// hInterface: A unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Because the RRASM
	// server does not maintain the interface handles, the RRAS server SHOULD check and
	// ensure that this handle represents a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	// dwTransportId: Specifies the transport (IPX, IPv4, or IPv6). It MUST be one of the
	// following values.<268> Otherwise an error other than those in the returned values
	// table is returned.
	//
	//	+---------------------+---------------+
	//	|                     |               |
	//	|        VALUE        |    MEANING    |
	//	|                     |               |
	//	+---------------------+---------------+
	//	+---------------------+---------------+
	//	| PID_IPX 0x0000002B  | IPX protocol  |
	//	+---------------------+---------------+
	//	| PID_IP 0x00000021   | IPv4 protocol |
	//	+---------------------+---------------+
	//	| PID_IPV6 0x00000057 | IPv6 protocol |
	//	+---------------------+---------------+
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
}

func (o *RouterInterfaceTransportRemoveRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportRemoveOperation) *xxx_RouterInterfaceTransportRemoveOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportRemoveOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	return op
}

func (o *RouterInterfaceTransportRemoveRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportRemoveOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
}
func (o *RouterInterfaceTransportRemoveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportRemoveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportRemoveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportRemoveResponse structure represents the RRouterInterfaceTransportRemove operation response
type RouterInterfaceTransportRemoveResponse struct {
	// Return: The RRouterInterfaceTransportRemove return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportRemoveResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportRemoveOperation) *xxx_RouterInterfaceTransportRemoveOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportRemoveOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportRemoveResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportRemoveOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceTransportRemoveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportRemoveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportRemoveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportAddOperation structure represents the RRouterInterfaceTransportAdd operation
type xxx_RouterInterfaceTransportAddOperation struct {
	Interface   uint32                    `idl:"name:hInterface" json:"interface"`
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportAddOperation) OpNum() int { return 17 }

func (o *xxx_RouterInterfaceTransportAddOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportAdd"
}

func (o *xxx_RouterInterfaceTransportAddOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportAddOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportAddOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportAddOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportAddOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceTransportAddOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportAddRequest structure represents the RRouterInterfaceTransportAdd operation request
type RouterInterfaceTransportAddRequest struct {
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Because the RRASM
	// server does not maintain the interface handles, the RRAS server SHOULD check and
	// ensure that this handle represents a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	// dwTransportId: Specifies the transport (IPX, IPv4, or IPv6) and MUST be one of the
	// following values.<270> Otherwise an error, other than those in the return values
	// table, is returned.
	//
	//	+---------------------+---------------+
	//	|                     |               |
	//	|        VALUE        |    MEANING    |
	//	|                     |               |
	//	+---------------------+---------------+
	//	+---------------------+---------------+
	//	| PID_IPX 0x0000002B  | IPX protocol  |
	//	+---------------------+---------------+
	//	| PID_IP 0x00000021   | IPv4 protocol |
	//	+---------------------+---------------+
	//	| PID_IPV6 0x00000057 | IPv6 protocol |
	//	+---------------------+---------------+
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
	// pInfoStruct: This is a pointer of type DIM_INTERFACE_CONTAINER (section 2.2.1.2.2).
	//
	// pInterfaceInfo and dwInterfaceInfoSize of DIM_INTERFACE_CONTAINER MUST be set to
	// valid values. All other DIM_INTERFACE_CONTAINER fields are ignored.
	//
	// pInterfaceInfo points to a valid RTR_INFO_BLOCK_HEADER (section 2.2.1.2.3) and RTR_TOC_ENTRY
	// (section 2.2.1.2.4). Otherwise, an error is returned. The RRASM server does not store
	// the interface configurations for the various transports that the RRAS server supports.
	// As part of its processing, the RRAS server SHOULD check that InfoType of RTR_TOC_ENTRY
	// is supported.
	//
	// Only a combination of the following entries SHOULD<271> be present in pInterfaceInfo.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	|                                          |          MEANING AND CONDITIONS (INTERFACE MEANS THE ONE IDENTIFIED BY           |              STRUCTURE POINTED TO AT              |
	//	|                  VALUE                   |                                   HINTERFACE)                                    |                      OFFSET                       |
	//	|                                          |                                                                                  |                                                   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_IN_FILTER_INFO 0xFFFF0001             | This is the input filter that MUST be applied to IP packets sent to the RRAS     | FILTER_DESCRIPTOR (section 2.2.1.2.5)             |
	//	|                                          | server. The information is overwritten. The interface MUST NOT be of type        |                                                   |
	//	|                                          | ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or ROUTER_IF_TYPE_DIALOUT.     |                                                   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_OUT_FILTER_INFO 0xFFFF0002            | This is the output filter that MUST be applied to IP packets sent from the       | FILTER_DESCRIPTOR (section 2.2.1.2.5)             |
	//	|                                          | RRAS server. The information is overwritten. The interface MUST NOT be of type   |                                                   |
	//	|                                          | ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or ROUTER_IF_TYPE_DIALOUT.     |                                                   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_INTERFACE_STATUS_INFO 0xFFFF0004      | The interface IP status info is overwritten.                                     | INTERFACE_STATUS_INFO (section 2.2.1.2.18)        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_ROUTER_DISC_INFO 0xFFFF0007           | Router discovery information is overwritten.                                     | RTR_DISC_INFO (section 2.2.1.2.14)                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_MCAST_BOUNDARY_INFO 0xFFFF000B        | Multicast boundary information is added.                                         | MIB_BOUNDARYROW (section 2.2.1.2.24)              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_IFFILTER_INFO 0xFFFF000D              | IP interface filter information is overwritten. Interface MUST NOT be of type    | IFFILTER_INFO (section 2.2.1.2.88)                |
	//	|                                          | ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or ROUTER_IF_TYPE_DIALOUT.     |                                                   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_MCAST_LIMIT_INFO 0xFFFF000E           | Multicast configuration information.                                             | MIB_MCAST_LIMIT_ROW (section 2.2.1.2.16)          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_ROUTE_INFO 0xFFFF0005                 | The dwAdminStatus MUST be IF_ADMIN_STATUS_UP if the route information is added.  | INTERFACE_ROUTE_INFO (section 2.2.1.2.11)         |
	//	|                                          | If bV4 of INTERFACE_ROUTE_INFO is set, it indicates an IPv4 route is added;      |                                                   |
	//	|                                          | otherwise, an IPv6 route is added.                                               |                                                   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_IN_FILTER_INFO_V6 0xFFFF0011          | This is the input filter that MUST be applied to IPv6 packets sent to the        | FILTER_DESCRIPTOR_V6 (section 2.2.1.2.7)          |
	//	|                                          | RRAS server. The information is overwritten. The interface MUST NOT be of type   |                                                   |
	//	|                                          | ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or ROUTER_IF_TYPE_DIALOUT.     |                                                   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_OUT_FILTER_INFO_V6 0xFFFF0012         | This is the output filter that MUST be applied to IPv6 packets sent from the     | FILTER_DESCRIPTOR_V6 (section 2.2.1.2.7)          |
	//	|                                          | RRAS server. The information is overwritten. The interface MUST NOT be of type   |                                                   |
	//	|                                          | ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or ROUTER_IF_TYPE_DIALOUT.     |                                                   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_DEMAND_DIAL_FILTER_INFO_V6 0xFFFF0013 | IPv6 traffic that matches this filter indicates that a site-to-site connection   | FILTER_DESCRIPTOR_V6 (section 2.2.1.2.7)          |
	//	|                                          | MUST be available and all IPv6 packets matching this filter MUST be routed into  |                                                   |
	//	|                                          | the connection. The interface MUST be of type ROUTER_IF_TYPE_FULL_ROUTER or      |                                                   |
	//	|                                          | ROUTER_IF_TYPE_HOME_ROUTER and the filters are overwritten                       |                                                   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| IP_IFFILTER_INFO_V6 0xFFFF0014           | IPv6 interface filter information is overwritten. The interface MUST             | IFFILTER_INFO (section 2.2.1.2.88)                |
	//	|                                          | NOT be of type ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or              |                                                   |
	//	|                                          | ROUTER_IF_TYPE_DIALOUT.                                                          |                                                   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| MS_IP_BOOTP 0x0000270F                   | IP BOOTP interface, information depending.                                       | IPBOOTP_IF_CONFIG (section 2.2.1.2.150)           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| MS_IP_IGMP 0x4137000A                    | IGMP interface information.                                                      | IGMP_MIB_IF_CONFIG (section 2.2.1.2.174)          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| MS_IP_RIP 0x00000008                     | IP RIP interface information.                                                    | IPRIP_IF_CONFIG (section 2.2.1.2.166)             |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| MS_IP_DHCP_ALLOCATOR 0x81372714          | DHCP allocator interface information. Used only with                             | IP_AUTO_DHCP_INTERFACE_INFO (section 2.2.1.2.192) |
	//	|                                          | RRouterInterfaceTransportSetGlobalInfo (section 3.1.4.10).                       |                                                   |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| MS_IP_DNS_PROXY 0x81372713               | DNS proxy interface information.                                                 | IP_DNS_PROXY_INTERFACE_INFO (section 2.2.1.2.194) |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| MS_IP_NAT 0x81372715                     | IP NAT interface information.                                                    | IP_NAT_INTERFACE_INFO (section 2.2.1.2.197)       |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| MS_IP_OSPF<272> 0x0000000D               | OSPF interface information is added.                                             | OSPF_INTERFACE_PARAM (section 2.2.1.2.215)        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| MS_IPV6_DHCP 0x000003E7                  | DHCPv6 Relay interface information.                                              | DHCPV6R_IF_CONFIG (section 2.2.1.2.159)           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+---------------------------------------------------+
	Info *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportAddRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportAddOperation) *xxx_RouterInterfaceTransportAddOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportAddOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportAddRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportAddOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
	o.Info = op.Info
}
func (o *RouterInterfaceTransportAddRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportAddRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportAddOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportAddResponse structure represents the RRouterInterfaceTransportAdd operation response
type RouterInterfaceTransportAddResponse struct {
	// Return: The RRouterInterfaceTransportAdd return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportAddResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportAddOperation) *xxx_RouterInterfaceTransportAddOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportAddOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportAddResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportAddOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceTransportAddResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportAddResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportAddOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportGetInfoOperation structure represents the RRouterInterfaceTransportGetInfo operation
type xxx_RouterInterfaceTransportGetInfoOperation struct {
	Interface   uint32                    `idl:"name:hInterface" json:"interface"`
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) OpNum() int { return 18 }

func (o *xxx_RouterInterfaceTransportGetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportGetInfo"
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportGetInfoRequest structure represents the RRouterInterfaceTransportGetInfo operation request
type RouterInterfaceTransportGetInfoRequest struct {
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Since the RRASM
	// server does not maintain the interface handles, the RRAS server SHOULD check and
	// ensure that this handle is a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	// dwTransportId: Specifies the transport for which the information is retrieved (IPX,
	// IPv4, or IPv6). It MUST be one of the following values.<274> Otherwise an error other
	// than those in the return values table is returned.
	//
	//	+---------------------+--------------------+
	//	|                     |                    |
	//	|        VALUE        |      MEANING       |
	//	|                     |                    |
	//	+---------------------+--------------------+
	//	+---------------------+--------------------+
	//	| PID_IPX 0x0000002B  | IPX protocol       |
	//	+---------------------+--------------------+
	//	| PID_IP 0x00000021   | IPv4 protocol      |
	//	+---------------------+--------------------+
	//	| PID_IPV6 0x00000057 | IPv6 protocol<275> |
	//	+---------------------+--------------------+
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
	// pInfoStruct: This is a pointer of type DIM_INTERFACE_CONTAINER. It MUST NOT be NULL.
	//
	// fGetInterfaceInfo of DIM_INTERFACE_CONTAINER MUST be set to 1.
	Info *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetInfoOperation) *xxx_RouterInterfaceTransportGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetInfoOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
	o.Info = op.Info
}
func (o *RouterInterfaceTransportGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportGetInfoResponse structure represents the RRouterInterfaceTransportGetInfo operation response
type RouterInterfaceTransportGetInfoResponse struct {
	// pInfoStruct: This is a pointer of type DIM_INTERFACE_CONTAINER. It MUST NOT be NULL.
	//
	// fGetInterfaceInfo of DIM_INTERFACE_CONTAINER MUST be set to 1.
	Info *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRouterInterfaceTransportGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetInfoOperation) *xxx_RouterInterfaceTransportGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RouterInterfaceTransportGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportSetInfoOperation structure represents the RRouterInterfaceTransportSetInfo operation
type xxx_RouterInterfaceTransportSetInfoOperation struct {
	Interface   uint32                    `idl:"name:hInterface" json:"interface"`
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) OpNum() int { return 19 }

func (o *xxx_RouterInterfaceTransportSetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportSetInfo"
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceTransportSetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportSetInfoRequest structure represents the RRouterInterfaceTransportSetInfo operation request
type RouterInterfaceTransportSetInfoRequest struct {
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Because the RRASM
	// server does not maintain the interface handles, the RRAS server SHOULD check and
	// ensure that this handle represents a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	// dwTransportId: Specifies the transport for which the information is set (IPX, IPv4,
	// or IPv6). It MUST be one of the following values.<277> Otherwise an error other than
	// those in the return values table is returned.
	//
	//	+---------------------+---------------+
	//	|                     |               |
	//	|        VALUE        |    MEANING    |
	//	|                     |               |
	//	+---------------------+---------------+
	//	+---------------------+---------------+
	//	| PID_IPX 0x0000002B  | IPX protocol  |
	//	+---------------------+---------------+
	//	| PID_IP 0x00000021   | IPv4 protocol |
	//	+---------------------+---------------+
	//	| PID_IPV6 0x00000057 | IPv6 protocol |
	//	+---------------------+---------------+
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
	// pInfoStruct: This is a pointer of type DIM_INTERFACE_CONTAINER.
	//
	// pInterfaceInfo and dwInterfaceInfoSize of PDIM_INTERFACE_CONTAINER MUST be set. The
	// rest of the fields are ignored. The RRASM server does not store the interface configurations
	// for the various transports that the RRAS server supports. The RRAS server SHOULD
	// check that the InfoType field of RTR_TOC_ENTRY (section 2.2.1.2.4) is supported.
	//
	// Only a combination of the following entries of RTR_TOC_ENTRY MUST be present in pInterfaceInfo.<278>
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                          |          MEANING AND CONDITIONS (INTERFACE MEANS THE ONE IDENTIFIED BY           |                             STRUCTURE POINTED TO AT                              |
	//	|                  VALUE                   |                                   HINTERFACE)                                    |                                      OFFSET                                      |
	//	|                                          |                                                                                  |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_INTERFACE_STATUS_INFO 0xFFFF0004      | Interface IP status information is overwritten.                                  | INTERFACE_STATUS_INFO (section 2.2.1.2.18)                                       |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_ROUTER_DISC_INFO 0xFFFF0007           | Router discovery information is overwritten.                                     | RTR_DISC_INFO (section 2.2.1.2.14)                                               |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_MCAST_BOUNDARY_INFO 0xFFFF000B        | Multicast boundary information is added.                                         | MIB_BOUNDARYROW (section 2.2.1.2.24)                                             |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_IFFILTER_INFO 0xFFFF000D              | IP interface filter information is overwritten. The interface MUST               | IFFILTER_INFO (section 2.2.1.2.88)                                               |
	//	|                                          | NOT be of type ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or              |                                                                                  |
	//	|                                          | ROUTER_IF_TYPE_DIALOUT.                                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_MCAST_LIMIT_INFO 0xFFFF000E           | Multicast configuration information.                                             | MIB_MCAST_LIMIT_ROW (section 2.2.1.2.16)                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_ROUTE_INFO 0xFFFF0005                 | The dwAdminStatus MUST be IF_ADMIN_STATUS_UP if the route information is added.  | INTERFACE_ROUTE_INFO (section 2.2.1.2.11)                                        |
	//	|                                          | If bV4 of INTERFACE_ROUTE_INFO is set, it indicates an IPv4 route is added;      |                                                                                  |
	//	|                                          | otherwise, an IPv6 route is added.                                               |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_IN_FILTER_INFO 0xFFFF0001             | This is the input filter that MUST be applied to IP packets sent to the RRAS     | FILTER_DESCRIPTOR (section 2.2.1.2.5)                                            |
	//	|                                          | server. The information is overwritten. The interface MUST NOT be of type        |                                                                                  |
	//	|                                          | ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or ROUTER_IF_TYPE_DIALOUT.     |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_OUT_FILTER_INFO 0xFFFF0002            | This is the output filter that MUST be applied to IP packets sent from the       | FILTER_DESCRIPTOR (section 2.2.1.2.5)                                            |
	//	|                                          | RRAS server. The information is overwritten. The interface MUST NOT be of type   |                                                                                  |
	//	|                                          | ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or ROUTER_IF_TYPE_DIALOUT.     |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_DEMAND_DIAL_FILTER_INFO 0xFFFF0009    | IPv4 traffic that matches this filter indicates that there is a site-to-site     | FILTER_DESCRIPTOR (section 2.2.1.2.5)                                            |
	//	|                                          | connection available into which all the IPv4 packets (matching this filter)      |                                                                                  |
	//	|                                          | are routed. The information is overwritten. The interface MUST be of type        |                                                                                  |
	//	|                                          | ROUTER_IF_TYPE_FULL_ROUTER or ROUTER_IF_TYPE_HOME_ROUTER.                        |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_IN_FILTER_INFO_V6 0xFFFF0011          | This is the input filter that MUST be applied to IPv6 packets sent to the        | FILTER_DESCRIPTOR_V6 (section 2.2.1.2.7)                                         |
	//	|                                          | RRAS server. The information is overwritten. The interface MUST NOT be of type   |                                                                                  |
	//	|                                          | ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or ROUTER_IF_TYPE_DIALOUT.     |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_OUT_FILTER_INFO_V6 0xFFFF0012         | This is the output filter that MUST be applied to IPv6 packets sent from the     | FILTER_DESCRIPTOR_V6 (section 2.2.1.2.7)                                         |
	//	|                                          | RRAS server. The information is overwritten. The interface MUST NOT be of type   |                                                                                  |
	//	|                                          | ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or ROUTER_IF_TYPE_DIALOUT.     |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_DEMAND_DIAL_FILTER_INFO_V6 0xFFFF0013 | IPv6 traffic that matches this filter indicates that a site-to-site connection   | FILTER_DESCRIPTOR_V6 (section 2.2.1.2.7)                                         |
	//	|                                          | MUST be available and all IPv6 packets matching this filter MUST be routed into  |                                                                                  |
	//	|                                          | the connection. The interface MUST be of type ROUTER_IF_TYPE_FULL_ROUTER or      |                                                                                  |
	//	|                                          | ROUTER_IF_TYPE_HOME_ROUTER. The filters are overwritten.                         |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_IFFILTER_INFO_V6 0xFFFF0014           | The IPv6 interface filter information is overwritten. The interface              | IFFILTER_INFO (section 2.2.1.2.88)                                               |
	//	|                                          | MUST NOT be of type ROUTER_IF_TYPE_INTERNAL, ROUTER_IF_TYPE_LOOPBACK, or         |                                                                                  |
	//	|                                          | ROUTER_IF_TYPE_DIALOUT.                                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_BOOTP 0x0000270F                   | IP BOOTP interface information.                                                  | IPBOOTP_IF_CONFIG (section 2.2.1.2.150)                                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_IGMP 0x4137000A                    | IGMP interface information. It can only be set once.                             | IGMP_MIB_IF_CONFIG (section 2.2.1.2.174)                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_RIP 0x00000008                     | IP RIP interface information.                                                    | IPRIP_IF_CONFIG (section 2.2.1.2.166)                                            |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_DHCP_ALLOCATOR 0x81372714          | DHCP allocator interface information.                                            | IP_AUTO_DHCP_INTERFACE_INFO (section 2.2.1.2.192)                                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_DNS_PROXY 0x81372713               | DNS proxy interface information.                                                 | IP_DNS_PROXY_INTERFACE_INFO (section 2.2.1.2.194)                                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_NAT 0x81372715                     | IP NAT interface information.                                                    | IP_NAT_INTERFACE_INFO (section 2.2.1.2.197)                                      |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_OSPF<279> 0x0000000D               | OSPF interface information is set. This MUST end the configuration buffer by     | OSPF_INTERFACE_PARAM (section 2.2.1.2.215, OSPF_NBMA_NEIGHBOR_PARAM (section     |
	//	|                                          | OSPF_END_PARAM_TYPE. If passed with any other structure, it will return an       | 2.2.1.2.216)                                                                     |
	//	|                                          | error.                                                                           |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IPV6_DHCP 0x000003E7                  | DHCPv6 Relay interface information. It can only be set once.                     | DHCPV6R_IF_CONFIG (section 2.2.1.2.159)                                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	Info *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportSetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetInfoOperation) *xxx_RouterInterfaceTransportSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetInfoOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
	o.Info = op.Info
}
func (o *RouterInterfaceTransportSetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportSetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportSetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportSetInfoResponse structure represents the RRouterInterfaceTransportSetInfo operation response
type RouterInterfaceTransportSetInfoResponse struct {
	// Return: The RRouterInterfaceTransportSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportSetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetInfoOperation) *xxx_RouterInterfaceTransportSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportSetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceTransportSetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportSetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportSetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceEnumOperation structure represents the RRouterInterfaceEnum operation
type xxx_RouterInterfaceEnumOperation struct {
	Level                  uint32                      `idl:"name:dwLevel" json:"level"`
	Info                   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                      `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	EntriesRead            uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries           uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume                 uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceEnumOperation) OpNum() int { return 20 }

func (o *xxx_RouterInterfaceEnumOperation) OpName() string { return "/dimsvc/v0/RRouterInterfaceEnum" }

func (o *xxx_RouterInterfaceEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceEnumRequest structure represents the RRouterInterfaceEnum operation request
type RouterInterfaceEnumRequest struct {
	// dwLevel: This is of type DWORD and SHOULD be set to zero (0).
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER. pInfoStruct.dwBufferSize
	// SHOULD be initialized to zero (0). Upon successful return, the pInfoStruct->pBuffer
	// is cast to an array of MPRI_INTERFACE_0 (section 2.2.1.2.81), and the array size
	// is determined by the value to pointer lpdwEntriesRead.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// dwPreferedMaximumLength: This is of type DWORD and SHOULD specify the preferred maximum
	// length of returned data (pInfoStruct->pBuffer) in bytes. If this parameter is -1,
	// the buffer returned is large enough to hold all available information.
	PreferredMaximumLength uint32 `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	// lpdwResumeHandle: This is a pointer to type DWORD and specifies a resume handle used
	// to continue the enumeration. The lpdwResumeHandle parameter is zero (0) on the first
	// call and left unchanged on subsequent calls (the caller MUST pass the same returned
	// value in the next call to this function). If the return code is ERROR_MORE_DATA (0x000000EA),
	// another call MAY be made using this handle to retrieve more data. If the handle is
	// NULL upon return, the enumeration is complete. This handle is invalid for other types
	// of error returns.
	Resume uint32 `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *RouterInterfaceEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceEnumOperation) *xxx_RouterInterfaceEnumOperation {
	if op == nil {
		op = &xxx_RouterInterfaceEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.PreferredMaximumLength = o.PreferredMaximumLength
	op.Resume = o.Resume
	return op
}

func (o *RouterInterfaceEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceEnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *RouterInterfaceEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceEnumResponse structure represents the RRouterInterfaceEnum operation response
type RouterInterfaceEnumResponse struct {
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER. pInfoStruct.dwBufferSize
	// SHOULD be initialized to zero (0). Upon successful return, the pInfoStruct->pBuffer
	// is cast to an array of MPRI_INTERFACE_0 (section 2.2.1.2.81), and the array size
	// is determined by the value to pointer lpdwEntriesRead.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// lpdwEntriesRead: This is a pointer to type DWORD. Upon successful return, this determines
	// the total number of connections enumerated from the current resume position given
	// by lpdwResumeHandle.
	EntriesRead uint32 `idl:"name:lpdwEntriesRead" json:"entries_read"`
	// lpdwTotalEntries: This is a pointer to type DWORD and receives the total number of
	// connections that could have been enumerated from the current resume position given
	// by lpdwResumeHandle.
	TotalEntries uint32 `idl:"name:lpdwTotalEntries" json:"total_entries"`
	// lpdwResumeHandle: This is a pointer to type DWORD and specifies a resume handle used
	// to continue the enumeration. The lpdwResumeHandle parameter is zero (0) on the first
	// call and left unchanged on subsequent calls (the caller MUST pass the same returned
	// value in the next call to this function). If the return code is ERROR_MORE_DATA (0x000000EA),
	// another call MAY be made using this handle to retrieve more data. If the handle is
	// NULL upon return, the enumeration is complete. This handle is invalid for other types
	// of error returns.
	Resume uint32 `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRouterInterfaceEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceEnumOperation) *xxx_RouterInterfaceEnumOperation {
	if op == nil {
		op = &xxx_RouterInterfaceEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.EntriesRead = o.EntriesRead
	op.TotalEntries = o.TotalEntries
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.EntriesRead = op.EntriesRead
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *RouterInterfaceEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceConnectOperation structure represents the RRouterInterfaceConnect operation
type xxx_RouterInterfaceConnectOperation struct {
	Interface        uint32 `idl:"name:hInterface" json:"interface"`
	Event            uint64 `idl:"name:hEvent" json:"event"`
	Blocking         uint32 `idl:"name:fBlocking" json:"blocking"`
	CallersProcessID uint32 `idl:"name:dwCallersProcessId" json:"callers_process_id"`
	Return           uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceConnectOperation) OpNum() int { return 21 }

func (o *xxx_RouterInterfaceConnectOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceConnect"
}

func (o *xxx_RouterInterfaceConnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceConnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.Event)); err != nil {
			return err
		}
	}
	// fBlocking {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Blocking); err != nil {
			return err
		}
	}
	// dwCallersProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CallersProcessID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceConnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.Event)); err != nil {
			return err
		}
	}
	// fBlocking {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Blocking); err != nil {
			return err
		}
	}
	// dwCallersProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CallersProcessID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceConnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceConnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceConnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceConnectRequest structure represents the RRouterInterfaceConnect operation request
type RouterInterfaceConnectRequest struct {
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// or RRouterInterfaceGetHandle (section 3.1.4.12). Since the RRASM server does not
	// maintain the interface handles, the RRAS server SHOULD check and ensure that this
	// handle is a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	// hEvent: The calling application MUST specify NULL for this.
	Event uint64 `idl:"name:hEvent" json:"event"`
	// fBlocking: If this parameter is set to 1, the function does not return until the
	// connection attempt has completed.
	Blocking uint32 `idl:"name:fBlocking" json:"blocking"`
	// dwCallersProcessId: This is for internal use and SHOULD be ignored by the server.
	CallersProcessID uint32 `idl:"name:dwCallersProcessId" json:"callers_process_id"`
}

func (o *RouterInterfaceConnectRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceConnectOperation) *xxx_RouterInterfaceConnectOperation {
	if op == nil {
		op = &xxx_RouterInterfaceConnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.Event = o.Event
	op.Blocking = o.Blocking
	op.CallersProcessID = o.CallersProcessID
	return op
}

func (o *RouterInterfaceConnectRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceConnectOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.Event = op.Event
	o.Blocking = op.Blocking
	o.CallersProcessID = op.CallersProcessID
}
func (o *RouterInterfaceConnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceConnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceConnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceConnectResponse structure represents the RRouterInterfaceConnect operation response
type RouterInterfaceConnectResponse struct {
	// Return: The RRouterInterfaceConnect return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceConnectResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceConnectOperation) *xxx_RouterInterfaceConnectOperation {
	if op == nil {
		op = &xxx_RouterInterfaceConnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceConnectResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceConnectOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceConnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceConnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceConnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceDisconnectOperation structure represents the RRouterInterfaceDisconnect operation
type xxx_RouterInterfaceDisconnectOperation struct {
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceDisconnectOperation) OpNum() int { return 22 }

func (o *xxx_RouterInterfaceDisconnectOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceDisconnect"
}

func (o *xxx_RouterInterfaceDisconnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDisconnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDisconnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDisconnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDisconnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceDisconnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceDisconnectRequest structure represents the RRouterInterfaceDisconnect operation request
type RouterInterfaceDisconnectRequest struct {
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Since the RRASM
	// server does not maintain the interface handles, the RRAS server SHOULD check and
	// ensure that this handle is a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceDisconnectRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDisconnectOperation) *xxx_RouterInterfaceDisconnectOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceDisconnectRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDisconnectOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
}
func (o *RouterInterfaceDisconnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceDisconnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDisconnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceDisconnectResponse structure represents the RRouterInterfaceDisconnect operation response
type RouterInterfaceDisconnectResponse struct {
	// Return: The RRouterInterfaceDisconnect return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceDisconnectResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDisconnectOperation) *xxx_RouterInterfaceDisconnectOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceDisconnectResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDisconnectOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceDisconnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceDisconnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDisconnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceUpdateRoutesOperation structure represents the RRouterInterfaceUpdateRoutes operation
type xxx_RouterInterfaceUpdateRoutesOperation struct {
	Interface       uint32 `idl:"name:hInterface" json:"interface"`
	TransportID     uint32 `idl:"name:dwTransportId" json:"transport_id"`
	Event           uint64 `idl:"name:hEvent" json:"event"`
	ClientProcessID uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
	Return          uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) OpNum() int { return 23 }

func (o *xxx_RouterInterfaceUpdateRoutesOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceUpdateRoutes"
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.Event)); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientProcessID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.Event)); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientProcessID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceUpdateRoutesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceUpdateRoutesRequest structure represents the RRouterInterfaceUpdateRoutes operation request
type RouterInterfaceUpdateRoutesRequest struct {
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Since the RRASM
	// server does not maintain the interface handles, the RRAS server SHOULD check and
	// ensure that this handle represents a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	// dwTransportId: Specifies the transport for which routing information needs to be
	// updated. This MUST be one of the following values.<283>
	//
	//	+--------------------+---------------+
	//	|                    |               |
	//	|       VALUE        |    MEANING    |
	//	|                    |               |
	//	+--------------------+---------------+
	//	+--------------------+---------------+
	//	| PID_IPX 0x0000002B | IPX protocol  |
	//	+--------------------+---------------+
	//	| PID_IP 0x00000021  | IPv4 protocol |
	//	+--------------------+---------------+
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
	// hEvent: The calling application MUST specify NULL for this parameter.
	Event uint64 `idl:"name:hEvent" json:"event"`
	// dwClientProcessId: The current process identifier where the function is called from.
	// If this is a nonexistent process the method will fail with an error code, as specified
	// in [MS-ERREF]. Otherwise the process specified is notified.
	ClientProcessID uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
}

func (o *RouterInterfaceUpdateRoutesRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceUpdateRoutesOperation) *xxx_RouterInterfaceUpdateRoutesOperation {
	if op == nil {
		op = &xxx_RouterInterfaceUpdateRoutesOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	op.Event = o.Event
	op.ClientProcessID = o.ClientProcessID
	return op
}

func (o *RouterInterfaceUpdateRoutesRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceUpdateRoutesOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
	o.Event = op.Event
	o.ClientProcessID = op.ClientProcessID
}
func (o *RouterInterfaceUpdateRoutesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceUpdateRoutesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceUpdateRoutesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceUpdateRoutesResponse structure represents the RRouterInterfaceUpdateRoutes operation response
type RouterInterfaceUpdateRoutesResponse struct {
	// Return: The RRouterInterfaceUpdateRoutes return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceUpdateRoutesResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceUpdateRoutesOperation) *xxx_RouterInterfaceUpdateRoutesOperation {
	if op == nil {
		op = &xxx_RouterInterfaceUpdateRoutesOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceUpdateRoutesResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceUpdateRoutesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceUpdateRoutesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceUpdateRoutesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceUpdateRoutesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceQueryUpdateResultOperation structure represents the RRouterInterfaceQueryUpdateResult operation
type xxx_RouterInterfaceQueryUpdateResultOperation struct {
	Interface    uint32 `idl:"name:hInterface" json:"interface"`
	TransportID  uint32 `idl:"name:dwTransportId" json:"transport_id"`
	UpdateResult uint32 `idl:"name:pUpdateResult" json:"update_result"`
	Return       uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) OpNum() int { return 24 }

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceQueryUpdateResult"
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pUpdateResult {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UpdateResult); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pUpdateResult {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UpdateResult); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceQueryUpdateResultRequest structure represents the RRouterInterfaceQueryUpdateResult operation request
type RouterInterfaceQueryUpdateResultRequest struct {
	// hInterface: The unique identifier of an interface. This identifier can be obtained
	// from RRouterInterfaceCreate (section 3.1.4.13) or RRouterInterfaceGetHandle (section
	// 3.1.4.12). Because the RRASM server does not maintain interface handles, the RRAS
	// server SHOULD check and ensure that this handle represents a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	// dwTransportId: Specifies the transport for which routing information needs to be
	// updated. This MUST be one of the following values.<285> Otherwise an error is returned.
	//
	//	+--------------------+---------------+
	//	|                    |               |
	//	|       VALUE        |    MEANING    |
	//	|                    |               |
	//	+--------------------+---------------+
	//	+--------------------+---------------+
	//	| PID_IPX 0x0000002B | IPX protocol  |
	//	+--------------------+---------------+
	//	| PID_IP 0x00000021  | IPv4 protocol |
	//	+--------------------+---------------+
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
}

func (o *RouterInterfaceQueryUpdateResultRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceQueryUpdateResultOperation) *xxx_RouterInterfaceQueryUpdateResultOperation {
	if op == nil {
		op = &xxx_RouterInterfaceQueryUpdateResultOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	return op
}

func (o *RouterInterfaceQueryUpdateResultRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceQueryUpdateResultOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
}
func (o *RouterInterfaceQueryUpdateResultRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceQueryUpdateResultRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceQueryUpdateResultOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceQueryUpdateResultResponse structure represents the RRouterInterfaceQueryUpdateResult operation response
type RouterInterfaceQueryUpdateResultResponse struct {
	// pUpdateResult: A pointer to a DWORD variable. This variable receives the result of
	// the last call to RRouterInterfaceUpdateRoutes; see the return values of RRouterInterfaceUpdateRoutes.
	UpdateResult uint32 `idl:"name:pUpdateResult" json:"update_result"`
	// Return: The RRouterInterfaceQueryUpdateResult return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceQueryUpdateResultResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceQueryUpdateResultOperation) *xxx_RouterInterfaceQueryUpdateResultOperation {
	if op == nil {
		op = &xxx_RouterInterfaceQueryUpdateResultOperation{}
	}
	if o == nil {
		return op
	}
	op.UpdateResult = o.UpdateResult
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceQueryUpdateResultResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceQueryUpdateResultOperation) {
	if o == nil {
		return
	}
	o.UpdateResult = op.UpdateResult
	o.Return = op.Return
}
func (o *RouterInterfaceQueryUpdateResultResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceQueryUpdateResultResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceQueryUpdateResultOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceUpdatePhonebookInfoOperation structure represents the RRouterInterfaceUpdatePhonebookInfo operation
type xxx_RouterInterfaceUpdatePhonebookInfoOperation struct {
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) OpNum() int { return 25 }

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceUpdatePhonebookInfo"
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceUpdatePhonebookInfoRequest structure represents the RRouterInterfaceUpdatePhonebookInfo operation request
type RouterInterfaceUpdatePhonebookInfoRequest struct {
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Because the RRASM
	// server does not maintain the interface handles, the RRAS server SHOULD check and
	// ensure that this handle is a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceUpdatePhonebookInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceUpdatePhonebookInfoOperation) *xxx_RouterInterfaceUpdatePhonebookInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceUpdatePhonebookInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceUpdatePhonebookInfoOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
}
func (o *RouterInterfaceUpdatePhonebookInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceUpdatePhonebookInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceUpdatePhonebookInfoResponse structure represents the RRouterInterfaceUpdatePhonebookInfo operation response
type RouterInterfaceUpdatePhonebookInfoResponse struct {
	// Return: The RRouterInterfaceUpdatePhonebookInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceUpdatePhonebookInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceUpdatePhonebookInfoOperation) *xxx_RouterInterfaceUpdatePhonebookInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceUpdatePhonebookInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceUpdatePhonebookInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceUpdatePhonebookInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceUpdatePhonebookInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntryCreateOperation structure represents the RMIBEntryCreate operation
type xxx_MIBEntryCreateOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntryCreateOperation) OpNum() int { return 26 }

func (o *xxx_MIBEntryCreateOperation) OpName() string { return "/dimsvc/v0/RMIBEntryCreate" }

func (o *xxx_MIBEntryCreateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryCreateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryCreateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryCreateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryCreateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MIBEntryCreateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBEntryCreateRequest structure represents the RMIBEntryCreate operation request
type MIBEntryCreateRequest struct {
	// dwPid: Specifies the transport protocol.
	PID uint32 `idl:"name:dwPid" json:"pid"`
	// dwRoutingPid: Specifies the routing protocol. The value of this parameter MUST be
	// set to IPRTRMGR_PID (0x00002710). The method MUST return an error other than those
	// specified in the return value table for any other value.
	RoutingPID uint32 `idl:"name:dwRoutingPid" json:"routing_pid"`
	// pInfoStuct: This is a pointer to a DIM_MIB_ENTRY_CONTAINER (section 2.2.1.2.19) structure
	// with valid dwMibInEntrySize and pMibInEntry fields. The dwMibOutEntrySize and pMibOutEntry
	// of the structure are ignored by the server for this method. dwMibInEntrySize MUST
	// be set to the size of data being passed in pMibInEntry. If pInfoStruct is NULL, an
	// error other than those specified in the return value table is returned. Otherwise,
	// the pointer is treated as a pointer to the DIM_MIB_ENTRY_CONTAINER structure.
	//
	// §  dwId set to the value ROUTE_MATCHING (0x0000001F).
	//
	// §  rgbyData MUST be a pointer to a MIB_IPDESTROW (section 2.2.1.2.20) structure.
	// A route corresponding to data in rgbyData is added in the IPv4 route table. The route
	// is created with the following fields specified:
	//
	// § ForwardRow: Contains the MIB_IPFORWARDROW (section 2.2.1.2.35) structure that
	// contains the route information with the following fields set:
	//
	// § dwForwardDest
	//
	// § dwForwardMask
	//
	// § dwForwardPolicy
	//
	// § dwForwardNextHop
	//
	// § ForwardType
	//
	// § ForwardProto
	//
	// § dwForwardAge
	//
	// § dwForwardNextHopAS
	//
	// § dwForwardIfIndex
	//
	// § dwForwardProto
	//
	// § dwForwardMetric1
	//
	// § dwForwardMetric2
	//
	// § dwForwardMetric3
	//
	// § dwForwardMetric4
	//
	// § dwForwardMetric5
	//
	// § dwForwardPreference
	//
	// § dwForwardViewSet: This MUST be the view information of the ForwardRow with the
	// following fields specified, see INTERFACE_ROUTE_INFO (section 2.2.1.2.11):
	//
	// § dwRtInfoDest
	//
	// § dwRtInfoMask
	//
	// § dwRtInfoNextHop
	//
	// § dwRtInfoMetric1
	//
	// § dwRtInfoMetric2
	//
	// § dwRtInfoMetric3
	//
	// § dwRtInfoIfIndex
	//
	// § dwRtInfoProto
	//
	// § pMibInEntry: The following fields are set to these values irrespective of the
	// values specified:
	//
	// § dwForwardPolicy is set to 0.
	//
	// § dwForwardMetric4 and dwForwardMetric5 are set to MIB_IPROUTE_METRIC_UNUSED (-1).
	//
	// § dwForwardPreference is set to IP_PRIORITY_DEFAULT_METRIC (0x0000007F).
	//
	// If pMibInEntry is NULL, an error other than those specified in the return values
	// table is returned. Otherwise the pointer to pMibInEntry is cast to a pointer to MIB_OPAQUE_INFO.
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntryCreateRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryCreateOperation) *xxx_MIBEntryCreateOperation {
	if op == nil {
		op = &xxx_MIBEntryCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntryCreateRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryCreateOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntryCreateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntryCreateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryCreateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntryCreateResponse structure represents the RMIBEntryCreate operation response
type MIBEntryCreateResponse struct {
	// Return: The RMIBEntryCreate return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntryCreateResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryCreateOperation) *xxx_MIBEntryCreateOperation {
	if op == nil {
		op = &xxx_MIBEntryCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *MIBEntryCreateResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryCreateOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MIBEntryCreateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntryCreateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryCreateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntryDeleteOperation structure represents the RMIBEntryDelete operation
type xxx_MIBEntryDeleteOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntryDeleteOperation) OpNum() int { return 27 }

func (o *xxx_MIBEntryDeleteOperation) OpName() string { return "/dimsvc/v0/RMIBEntryDelete" }

func (o *xxx_MIBEntryDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MIBEntryDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBEntryDeleteRequest structure represents the RMIBEntryDelete operation request
type MIBEntryDeleteRequest struct {
	// dwPid: Specifies the transport. The value of this field MUST be as follows.
	//
	//	+-------------------+-------------------------------------+
	//	|                   |                                     |
	//	|       VALUE       |               MEANING               |
	//	|                   |                                     |
	//	+-------------------+-------------------------------------+
	//	+-------------------+-------------------------------------+
	//	| PID_IP 0x00000021 | An IPv4 MIB entry is to be deleted. |
	//	+-------------------+-------------------------------------+
	PID uint32 `idl:"name:dwPid" json:"pid"`
	// dwRoutingPid: Specifies the routing protocol. This MUST be 0x00002710 (10000) and
	// indicates that pInfoStruct MUST point to MIB_OPAQUE_QUERY (section 2.2.1.2.53). Otherwise
	// an error other than those specified in the return values table is returned.
	RoutingPID uint32 `idl:"name:dwRoutingPid" json:"routing_pid"`
	// pInfoStuct: This parameter MUST be a pointer to the structure DIM_MIB_ENTRY_CONTAINER
	// (section 2.2.1.2.19) with valid dwMibInEntrySize and pMibInEntry fields. The parameters
	// dwMibOutEntrySize and pMibOutEntry are ignored by the server. dwMibInEntrySize MUST
	// be set to the size of data being passed in pMibInEntry.
	//
	// pMibInEntry MUST be a pointer to MIB_OPAQUE_QUERY. dwVarId and rgdwVarIndex of MIB_OPAQUE_QUERY
	// MUST be set to one of the following values that need to be deleted.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|          DWVARID          |                       STRUCTURE THAT MUST TO BE POINTED BY                       |                                                                                  |
	//	|           VALUE           |                                   RGDWVARINDEX                                   |                                     REMARKS                                      |
	//	|                           |                                                                                  |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| ROUTE_MATCHING 0x0000001F | MIB_IPDESTROW (section 2.2.1.2.20) rgdwVarIndex[0] MUST be dwForwardDest         | A route with the following matching fields specified in ForwardRow and           |
	//	|                           | rgdwVarIndex[1] MUST be dwForwardMask rgdwVarIndex[2] MUST be dwForwardIfIndex   | MIB_IPDESTROW is deleted from the IPv4 route table: dwForwardDest dwForwardMask  |
	//	|                           | rgdwVarIndex[3] MUST be dwForwardNextHop rgdwVarIndex[4] MUST be dwForwardProto  | dwForwardIfIndex dwForwardNextHop dwForwardProto                                 |
	//	+---------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| IP_NETTABLE 0x00000009    | MIB_IPNETTABLE (section 2.2.1.2.48) rgdwVarIndex[0] MUST be dwIfIndex            | An entry is deleted<289> whose interface index matches the dwIfIndex and whose   |
	//	|                           | rgdwVarIndex[1] MUST be dwAddr                                                   | IPv4 address matches the specified dwAddr.                                       |
	//	+---------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// If pMibInEntry is NULL, an error other than those specified in the return value table
	// is returned. Otherwise, the pointer to pMibInEntry is cast to a pointer to a MIB_OPAQUE_QUERY.
	//
	// If pInfoStruct is NULL, an error other than those specified in the following return
	// value table is returned.
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntryDeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryDeleteOperation) *xxx_MIBEntryDeleteOperation {
	if op == nil {
		op = &xxx_MIBEntryDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntryDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryDeleteOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntryDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntryDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntryDeleteResponse structure represents the RMIBEntryDelete operation response
type MIBEntryDeleteResponse struct {
	// Return: The RMIBEntryDelete return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntryDeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryDeleteOperation) *xxx_MIBEntryDeleteOperation {
	if op == nil {
		op = &xxx_MIBEntryDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *MIBEntryDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryDeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MIBEntryDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntryDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntrySetOperation structure represents the RMIBEntrySet operation
type xxx_MIBEntrySetOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntrySetOperation) OpNum() int { return 28 }

func (o *xxx_MIBEntrySetOperation) OpName() string { return "/dimsvc/v0/RMIBEntrySet" }

func (o *xxx_MIBEntrySetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntrySetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntrySetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntrySetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntrySetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MIBEntrySetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBEntrySetRequest structure represents the RMIBEntrySet operation request
type MIBEntrySetRequest struct {
	// dwPid: Specifies the transport protocol. The value of this parameter MUST be PID_IP.
	// The method MUST return an error other than those specified in the return value table
	// for any other value.
	//
	//	+-------------------+-------------------------------------+
	//	|                   |                                     |
	//	|       VALUE       |               MEANING               |
	//	|                   |                                     |
	//	+-------------------+-------------------------------------+
	//	+-------------------+-------------------------------------+
	//	| PID_IP 0x00000021 | An IPv4 MIB entry is to be deleted. |
	//	+-------------------+-------------------------------------+
	PID uint32 `idl:"name:dwPid" json:"pid"`
	// dwRoutingPid: Specifies the routing protocol. This MUST be 0x00002710 (10000) and
	// indicate that pInfoStruct MUST point to MIB_OPAQUE_QUERY (section 2.2.1.2.53). Otherwise,
	// an error other than those specified in the return table value is returned.
	RoutingPID uint32 `idl:"name:dwRoutingPid" json:"routing_pid"`
	// pInfoStuct: This MUST be a pointer to the structure DIM_MIB_ENTRY_CONTAINER (section
	// 2.2.1.2.19) with valid dwMibInEntrySize and pMibInEntry fields. The dwMibOutEntrySize
	// and pMibOutEntry of the structure are ignored by the server for this method. dwMibInEntrySize
	// MUST be set to the size of data being passed in pMibInEntry. If pInfoStuct is NULL,
	// an error other than those specified in the return value table is returned. Otherwise,
	// the pointer is treated as a pointer to the structure DIM_MIB_ENTRY_CONTAINER.
	//
	// Since the dwRoutingPid parameter of this method MUST take a value IPRTRMGR_PID (0x00002710),
	// the pMibInEntry field of pInfoStuct MUST be a pointer to MIB_OPAQUE_INFO (section
	// 2.2.1.2.52). The dwId and rgbyData fields of MIB_OPAQUE_INFO MUST be set to one of
	// the values in the following table.
	//
	//	+---------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	|           DWID            |     STRUCTURE THAT MUST BE POINTED TO BY      |                                                                                  |
	//	|           VALUE           |                   RGBYDATA                    |                                     REMARKS                                      |
	//	|                           |                                               |                                                                                  |
	//	+---------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ROUTE_MATCHING 0x0000001F | MIB_IPDESTROW (section 2.2.1.2.20)            | A route corresponding to data in rgbyData is added in the IPv4 route table.      |
	//	|                           |                                               | The route is created with the following fields: ForwardRow: Contains the         |
	//	|                           |                                               | MIB_IPFORWARDROW (section 2.2.1.2.35) structure that contains the route          |
	//	|                           |                                               | information with the following fields set: dwForwardDest dwForwardMask           |
	//	|                           |                                               | dwForwardPolicy dwForwardNextHop ForwardType ForwardProto dwForwardAge           |
	//	|                           |                                               | dwForwardNextHopAS dwForwardMetric1 dwForwardMetric2 dwForwardMetric3            |
	//	|                           |                                               | dwForwardMetric4 dwForwardMetric5 dwForwardPreference dwForwardViewSet           |
	//	|                           |                                               | dwRtInfoDest dwRtInfoMask dwRtInfoNextHop dwRtInfoIfIndex dwRtInfoProto          |
	//	|                           |                                               | dwRtInfoMetric1 dwRtInfoMetric2 dwRtInfoMetric3 pMibInEntry: The following       |
	//	|                           |                                               | fields are set to these values irrespective of the values specified              |
	//	|                           |                                               | (INTERFACE_ROUTE_INFO (section 2.2.1.2.11)): dwForwardPolicy is set to 0.        |
	//	|                           |                                               | dwForwardMetric4 and dwForwardMetric5 are set to MIB_IPROUTE_METRIC_UNUSED (-1). |
	//	|                           |                                               | dwForwardPreference is set to IP_PRIORITY_DEFAULT_METRIC (0x0000007F).           |
	//	+---------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| IF_ROW 0x00000002         | MIB_IFROW (section 2.2.1.2.29)                | Only dwAdminStatus can be set to IF_ADMIN_STATUS_DOWN or IF_ADMIN_STATUS_UP (see |
	//	|                           |                                               | dwAdminStatus of INTERFACE_STATUS_INFO (section 2.2.1.2.18).                     |
	//	+---------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_IF_ENTRY 0x00000017 | MIB_IPMCAST_IF_ENTRY (section 2.2.1.2.40)     | dwTtl MUST be set to less than or equal to 255. dwRateLimit MUST NOT be set to   |
	//	|                           |                                               | 0. dwIfIndex MUST be set to the index of the interface for which the entries are |
	//	|                           |                                               | being updated. The following entries can be set only if the operational status   |
	//	|                           |                                               | of the interface is IF_OPER_STATUS_OPERATIONAL (see dwOperStatus of MIB_IFROW:   |
	//	|                           |                                               | ulInMcastOctets ulOutMcastOctets                                                 |
	//	+---------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_BOUNDARY 0x0000001A | MIB_IPMCAST_BOUNDARY (section 2.2.1.2.37)     | If dwStatus is set to ROWSTATUS_CREATEANDGO, the boundary information            |
	//	|                           |                                               | specified in MIB_IPMCAST_BOUNDARY is created. If dwStatus field is set to        |
	//	|                           |                                               | ROWSTATUS_DESTROY, the boundary information specified in MIB_IPMCAST_BOUNDARY is |
	//	|                           |                                               | deleted.                                                                         |
	//	+---------------------------+-----------------------------------------------+----------------------------------------------------------------------------------+
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntrySetRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntrySetOperation) *xxx_MIBEntrySetOperation {
	if op == nil {
		op = &xxx_MIBEntrySetOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntrySetRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntrySetOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntrySetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntrySetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntrySetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntrySetResponse structure represents the RMIBEntrySet operation response
type MIBEntrySetResponse struct {
	// Return: The RMIBEntrySet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntrySetResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntrySetOperation) *xxx_MIBEntrySetOperation {
	if op == nil {
		op = &xxx_MIBEntrySetOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *MIBEntrySetResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntrySetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MIBEntrySetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntrySetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntrySetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntryGetOperation structure represents the RMIBEntryGet operation
type xxx_MIBEntryGetOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntryGetOperation) OpNum() int { return 29 }

func (o *xxx_MIBEntryGetOperation) OpName() string { return "/dimsvc/v0/RMIBEntryGet" }

func (o *xxx_MIBEntryGetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBEntryGetRequest structure represents the RMIBEntryGet operation request
type MIBEntryGetRequest struct {
	// dwPid: Specifies the transport protocol. The value of this parameter MUST be one
	// from the following table.<292> The method MUST return an error for any other values.
	//
	//	+---------------------+------------------------------------------------+
	//	|                     |                                                |
	//	|        VALUE        |                    MEANING                     |
	//	|                     |                                                |
	//	+---------------------+------------------------------------------------+
	//	+---------------------+------------------------------------------------+
	//	| PID_IPX 0x0000002B  | An IPX protocol MIB entry is to be retrieved.  |
	//	+---------------------+------------------------------------------------+
	//	| PID_IP 0x00000021   | An IPv4 protocol MIB entry is to be retrieved. |
	//	+---------------------+------------------------------------------------+
	//	| PID_IPV6 0x00000057 | An IPv6 protocol MIB entry is to be retrieved. |
	//	+---------------------+------------------------------------------------+
	PID uint32 `idl:"name:dwPid" json:"pid"`
	// dwRoutingPid: If dwPid is PID_IP and if dwRoutingPid is 10000, then pMibInEntry MUST
	// point to MIB_OPAQUE_QUERY (section 2.2.1.2.53). The dwVarId field of MIB_OPAQUE_QUERY
	// MUST be one of the following values. pMibOutEntry MUST be interpreted as a pointer
	// to MIB_OPAQUE_QUERY. See section 2.2.1.2.53 for details on the structure retrieved
	// by the rgdwVarIndex field of MIB_OPAQUE_QUERY for different dwVarId values.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|           DWVARID           |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IF_NUMBER 0x00000000        | Number of interfaces on the machine.                                             |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IF_TABLE 0x00000001         | Information about interface table.                                               |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IF_ROW 0x00000002           | Information about a particular interface.                                        |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_STATS 0x00000003         | Information about the IP protocol.                                               |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_ADDRTABLE 0x00000004     | Table of IPv4 address entries.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_ADDRROW 0x00000005       | Information for a particular IPv4 address.                                       |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_FORWARDNUMBER 0x00000006 | Information about number of routes in a particular IP routing table.             |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_FORWARDTABLE 0x00000007  | Table of IPv4 route entries.                                                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_NETTABLE 0x00000009      | Table of ARP entries for IPv4 addresses.                                         |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_NETROW 0x0000000A        | Information about an ARP table entry for an IPv4 address.                        |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ICMP_STATS 0x0000000B       | Statistics for ICMP messages on a particular computer.                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| TCP_STATS 0x0000000C        | Statistics for the TCP protocol running on the local computer.                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| UDP_STATS 0x0000000F        | Statistics for the UDP running on the local computer.                            |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_MFE 0x00000012        | Information for an IP multicast forwarding entry.                                |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_MFE_STATS 0x00000013  | Statistics associated with an MFE.                                               |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_IF_ENTRY 0x00000017   | Information about an IP multicast interface.                                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ROUTE_MATCHING 0x0000001F   | Information about a matching IP route.                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| BEST_IF 0x00000014          | Index of the interface that has the best route to a particular destination IPv4  |
	//	|                             | address.                                                                         |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_GLOBAL 0x00000018     | Global information for IP multicast on a particular computer.                    |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IF_STATUS 0x00000019        | Status information for a particular interface.                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_BOUNDARY 0x0000001A   | Information about a router's scoped IPv4 multicast address boundaries.           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_SCOPE 0x0000001B      | Information about a multicast scope.                                             |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//
	// If dwPid is PID_IPV6 and dwRoutingPid is 10000, then pMibInEntry MUST point to MIB_OPAQUE_QUERY
	// (section 2.2.1.2.53). The dwVarId field of pMibInEntry MUST be one of the following
	// values. pMibOutEntry MUST be interpreted as a pointer to MIB_OPAQUE_QUERY. See section
	// 2.2.1.2.53 for details on the structure retrieved by the rgdwVarIndex field of MIB_OPAQUE_QUERY
	// for different dwVarId values.
	//
	//	+--------------------------+------------------------------------------+
	//	|         DWVARID          |                                          |
	//	|          VALUE           |                 MEANING                  |
	//	|                          |                                          |
	//	+--------------------------+------------------------------------------+
	//	+--------------------------+------------------------------------------+
	//	| 0x00000001 IF_TABLE      | Information about interface table.       |
	//	+--------------------------+------------------------------------------+
	//	| 0x00000008 IP_FORWARDROW | Information about an IPv6 network route. |
	//	+--------------------------+------------------------------------------+
	//
	// If dwPid is PID_IP and if dwRoutingPid is not IPRTRMGR_PID (0x00002710), then dwRoutingPid
	// MUST be one of the entries in the Value column and pMibInEntry MUST be the corresponding
	// entry in the Structure to be pointed to by pMibInEntry column in the following table.
	// The routing protocols specified in the following table are valid only if the protocol
	// is already initialized by RRAS for IPv4. RRAS initializes a protocol when a RTR_TOC_ENTRY
	// with the InfoType corresponding to the protocol is present in the global configuration.
	// The RRouterInterfaceTransportSetGlobalInfo method enables specifying the global configuration
	// through the pGlobalInfo member of pInfoStruct.
	//
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                 |                                                  |                          STRUCTURE TO BE POINTED TO BY                           |                                                                                  |
	//	|              VALUE              |                     MEANING                      |                                   PMIBINENTRY                                    |                                     REMARKS                                      |
	//	|                                 |                                                  |                                                                                  |                                                                                  |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_BOOTP 0x0000270F          | An entry in IBOOTPMIB needs to be retrieved.     | IPBOOTP_MIB_GET_INPUT_DATA                                                       | See section 2.2.1.2.152 for details on how to populate                           |
	//	|                                 |                                                  |                                                                                  | IPBOOTP_MIB_GET_INPUT_DATA to retrieve different types of MIB entries.           |
	//	|                                 |                                                  |                                                                                  | When this method returns pMibOutEntry, pInfoStruct MUST be cast to               |
	//	|                                 |                                                  |                                                                                  | IPBOOTP_MIB_GET_OUTPUT_DATA. See section 2.2.1.2.157 on how to interpret the     |
	//	|                                 |                                                  |                                                                                  | data returned.                                                                   |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_NAT 0x81372715            | An entry in the NAT MIB needs to be retrieved.   | IP_NAT_MIB_QUERY                                                                 | See IP_NAT_MIB_QUERY (section 2.2.1.2.182) for details on how to populate        |
	//	|                                 |                                                  |                                                                                  | this structure to retrieve the entries required. The values are returned in      |
	//	|                                 |                                                  |                                                                                  | IP_NAT_MIB_QUERY.                                                                |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_DNS_PROXY 0x81372713      | An entry in DNS Proxy MIB needs to be retrieved. | IP_DNS_PROXY_MIB_QUERY                                                           | See section 2.2.1.2.186 for details on how to populate IP_DNS_PROXY_MIB_QUERY.   |
	//	|                                 |                                                  |                                                                                  | The values are returned in IP_DNS_PROXY_MIB_QUERY.                               |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_DHCP_ALLOCATOR 0x81372714 | An entry in auto DHCP MIB is retrieved.          | IP_AUTO_DHCP_MIB_QUERY                                                           | See section 2.2.1.2.188 for details on how to populate IP_AUTO_DHCP_MIB_QUERY.   |
	//	|                                 |                                                  |                                                                                  | The values are returned in the same structure.                                   |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                 |                                                  |                                                                                  |                                                                                  |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_OSPF<293> 0x0000000D      | An entry in OSPF MIB needs to be retrieved.      | MIB_DA_MSG MIB_DA_MSG OSPF_GLOBAL_PARAM OSPF_ROUTE_FILTER_INFO                   |                                                                                  |
	//	|                                 |                                                  | OSPF_PROTO_FILTER_INFO OSPF_AREA_PARAM OSPF_AREA_RANGE_PARAM                     |                                                                                  |
	//	|                                 |                                                  | OSPF_VIRT_INTERFACE_PARAM OSPF_NBMA_NEIGHBOR_PARAM                               |                                                                                  |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// If dwPid is PID_IPV6 and if dwRoutingPid is not IPRTRMGR_PID (0x00002710), then dwRoutingPid
	// MUST be one of the entries in the Value column and pMibInEntry MUST be the corresponding
	// entry in the Structure to be pointed to by pMibInEntry in the following table. The
	// routing protocols specified in the following table are valid only if the protocol
	// is already initialized by RRAS for IPv6. RRAS initializes a protocol when a RTR_TOC_ENTRY
	// with the InfoType corresponding to the protocol is present in the global configuration.
	// The RRouterInterfaceTransportSetGlobalInfo method enables specifying the global configuration
	// through the pGlobalInfo member of pInfoStuct.
	//
	//	+-------------------------+-----------------------------------------------------------+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                         |                                                           |       STRUCTURE TO BE POINTED TO BY       |                                                                                  |
	//	|          VALUE          |                          MEANING                          |                PMIBINENTRY                |                                     REMARKS                                      |
	//	|                         |                                                           |                                           |                                                                                  |
	//	+-------------------------+-----------------------------------------------------------+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+-----------------------------------------------------------+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IPV6_DHCP 0x000003E7 | An entry in DHCPv6 Relay agent MIB needs to be retrieved. | DHCPV6R_MIB_GET_INPUT_DATA                | See section 2.2.1.2.160 on how to populate DHCPV6R_MIB_GET_INPUT_DATA. The       |
	//	|                         |                                                           |                                           | values are retuned in DHCPV6R_MIB_GET_OUTPUT_DATA.                               |
	//	+-------------------------+-----------------------------------------------------------+-------------------------------------------+----------------------------------------------------------------------------------+
	RoutingPID uint32 `idl:"name:dwRoutingPid" json:"routing_pid"`
	// pInfoStuct: This MUST be a pointer to the structure DIM_MIB_ENTRY_CONTAINER (section
	// 2.2.1.2.19) with valid dwMibInEntrySize and pMibInEntry fields. dwMibInEntrySize
	// MUST be set to the size of data being passed in pMibInEntry. dwMibOutEntrySize and
	// pMibOutEntry are populated by the server. The caller frees the memory pointed to
	// by pInfoStuct.
	//
	// If dwPid is PID_IPX, dwRoutingPid MUST be one of the entries in the Value column
	// and pMibInEntry MUST be the corresponding entry in the Structure to be pointed to
	// by pMibInEntry in the return values table.
	//
	//	+------------------------------+---------------------------------------------------+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                   |       STRUCTURE TO BE POINTED TO BY        |                                                                                  |
	//	|            VALUE             |                      MEANING                      |                PMIBINENTRY                 |                                     REMARKS                                      |
	//	|                              |                                                   |                                            |                                                                                  |
	//	+------------------------------+---------------------------------------------------+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+---------------------------------------------------+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| IPX_PROTOCOL_BASE 0x00000000 | IPX related information. The call was successful. | IPX_MIB_GET_INPUT_DATA                     | See section 2.2.1.2.105 for details on how to populate IPX_MIB_GET_INPUT_DATA to |
	//	|                              |                                                   |                                            | retrieve different types of MIB entries and how to interpret the data returned.  |
	//	+------------------------------+---------------------------------------------------+--------------------------------------------+----------------------------------------------------------------------------------+
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntryGetRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetOperation) *xxx_MIBEntryGetOperation {
	if op == nil {
		op = &xxx_MIBEntryGetOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntryGetRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntryGetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntryGetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntryGetResponse structure represents the RMIBEntryGet operation response
type MIBEntryGetResponse struct {
	// pInfoStuct: This MUST be a pointer to the structure DIM_MIB_ENTRY_CONTAINER (section
	// 2.2.1.2.19) with valid dwMibInEntrySize and pMibInEntry fields. dwMibInEntrySize
	// MUST be set to the size of data being passed in pMibInEntry. dwMibOutEntrySize and
	// pMibOutEntry are populated by the server. The caller frees the memory pointed to
	// by pInfoStuct.
	//
	// If dwPid is PID_IPX, dwRoutingPid MUST be one of the entries in the Value column
	// and pMibInEntry MUST be the corresponding entry in the Structure to be pointed to
	// by pMibInEntry in the return values table.
	//
	//	+------------------------------+---------------------------------------------------+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                   |       STRUCTURE TO BE POINTED TO BY        |                                                                                  |
	//	|            VALUE             |                      MEANING                      |                PMIBINENTRY                 |                                     REMARKS                                      |
	//	|                              |                                                   |                                            |                                                                                  |
	//	+------------------------------+---------------------------------------------------+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+---------------------------------------------------+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| IPX_PROTOCOL_BASE 0x00000000 | IPX related information. The call was successful. | IPX_MIB_GET_INPUT_DATA                     | See section 2.2.1.2.105 for details on how to populate IPX_MIB_GET_INPUT_DATA to |
	//	|                              |                                                   |                                            | retrieve different types of MIB entries and how to interpret the data returned.  |
	//	+------------------------------+---------------------------------------------------+--------------------------------------------+----------------------------------------------------------------------------------+
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	// Return: The RMIBEntryGet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntryGetResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetOperation) *xxx_MIBEntryGetOperation {
	if op == nil {
		op = &xxx_MIBEntryGetOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MIBEntryGetResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MIBEntryGetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntryGetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntryGetFirstOperation structure represents the RMIBEntryGetFirst operation
type xxx_MIBEntryGetFirstOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntryGetFirstOperation) OpNum() int { return 30 }

func (o *xxx_MIBEntryGetFirstOperation) OpName() string { return "/dimsvc/v0/RMIBEntryGetFirst" }

func (o *xxx_MIBEntryGetFirstOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetFirstOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetFirstOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetFirstOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetFirstOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetFirstOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBEntryGetFirstRequest structure represents the RMIBEntryGetFirst operation request
type MIBEntryGetFirstRequest struct {
	// dwPid: Specifies the transport protocol. The value of this parameter MUST be one
	// from the following table.<295>
	//
	//	+---------------------+---------------------------------------+
	//	|                     |                                       |
	//	|        VALUE        |                MEANING                |
	//	|                     |                                       |
	//	+---------------------+---------------------------------------+
	//	+---------------------+---------------------------------------+
	//	| PID_IPX 0x0000002B  | An IPX MIB entry is to be retrieved.  |
	//	+---------------------+---------------------------------------+
	//	| PID_IP 0x00000021   | An IPv4 MIB entry is to be retrieved. |
	//	+---------------------+---------------------------------------+
	//	| PID_IPV6 0x00000057 | An IPv6 MIB entry is to be retrieved. |
	//	+---------------------+---------------------------------------+
	PID uint32 `idl:"name:dwPid" json:"pid"`
	// dwRoutingPid: If dwPid is PID_IP and if dwRoutingPid is 10000, then pMibInEntry MUST
	// point to MIB_OPAQUE_QUERY (section 2.2.1.2.53). The dwVarId field of MIB_OPAQUE_QUERY
	// MUST be one of the following values. pMibOutEntry MUST be interpreted as a pointer
	// to MIB_OPAQUE_QUERY. See section 2.2.1.2.53 for details on the structure retrieved
	// by the rgdwVarIndex field of MIB_OPAQUE_QUERY for different dwVarId values.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|           DWVARID           |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IF_NUMBER 0x00000000        | Number of interfaces on the machine.                                             |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IF_TABLE 0x00000001         | Information about interface table.                                               |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IF_ROW 0x00000002           | Information about a particular interface.                                        |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_STATS 0x00000003         | Information about the IP protocol.                                               |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_ADDRTABLE 0x00000004     | Table of IPv4 address entries.                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_ADDRROW 0x00000005       | Information for a particular IPv4 address.                                       |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_FORWARDNUMBER 0x00000006 | Information about number of routes in a particular IP routing table.             |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_FORWARDTABLE 0x00000007  | Table of IPv4 route entries.                                                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_NETTABLE 0x00000009      | Table of ARP entries for IPv4 addresses.                                         |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IP_NETROW 0x0000000A        | Information about an ARP table entry for an IPv4 address.                        |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ICMP_STATS 0x0000000B       | Statistics for ICMP messages on a particular computer.                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| TCP_STATS 0x0000000C        | Statistics for the TCP protocol running on the local computer.                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| UDP_STATS 0x0000000F        | Statistics for the UDP running on the local computer.                            |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_MFE 0x00000012        | Information for an IP multicast forwarding entry.                                |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_MFE_STATS 0x00000013  | Statistics associated with an MFE.                                               |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_IF_ENTRY 0x00000017   | Information about an IP multicast interface.                                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ROUTE_MATCHING 0x0000001F   | Information about a matching IP route.                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| BEST_IF 0x00000014          | Index of the interface that has the best route to a particular destination IPv4  |
	//	|                             | address.                                                                         |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_GLOBAL 0x00000018     | Global information for IP multicast on a particular computer.                    |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| IF_STATUS 0x00000019        | Status information for a particular interface.                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_BOUNDARY 0x0000001A   | Information about a router's scoped IPv4 multicast address boundaries.           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| MCAST_SCOPE 0x0000001B      | Information about a multicast scope.                                             |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//
	// If dwPid is PID_IPV6 and dwRoutingPid is 10000, then pMibInEntry MUST point to MIB_OPAQUE_QUERY.
	// The dwVarId field of pMibInEntry MUST be one of the following values. pMibOutEntry
	// MUST be interpreted as a pointer to MIB_OPAQUE_QUERY. See section 2.2.1.2.53 for
	// details on the structure retrieved by the rgdwVarIndex field of MIB_OPAQUE_QUERY
	// for different dwVarId values.
	//
	//	+--------------------------+------------------------------------------+
	//	|         DWVARID          |                                          |
	//	|          VALUE           |                 MEANING                  |
	//	|                          |                                          |
	//	+--------------------------+------------------------------------------+
	//	+--------------------------+------------------------------------------+
	//	|                          |                                          |
	//	+--------------------------+------------------------------------------+
	//	| 0x00000008 IP_FORWARDROW | Information about an IPv6 network route. |
	//	+--------------------------+------------------------------------------+
	//
	// If dwPid is PID_IP and if dwRoutingPid is not IPRTRMGR_PID (0x00002710), then dwRoutingPid
	// MUST be one of the entries in the Value column and pMibInEntry MUST be the corresponding
	// entry in the Structure to be pointed to by pMibInEntry in the following table. The
	// following specified routing protocols are valid only if the protocol is already initialized
	// by RRAS for IPv4. RRAS initializes a protocol when an RTR_TOC_ENTRY with the InfoType
	// corresponding to the protocol is present in the global configuration. The RRouterInterfaceTransportSetGlobalInfo
	// method enables specifying the global configuration through the pGlobalInfo member
	// of pInfoStuct.
	//
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                 |                                                  |                          STRUCTURE TO BE POINTED TO BY                           |                                                                                  |
	//	|              VALUE              |                     MEANING                      |                                   PMIBINENTRY                                    |                                     REMARKS                                      |
	//	|                                 |                                                  |                                                                                  |                                                                                  |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_BOOTP 0x0000270F          | An entry in IBOOTPMIB needs to be retrieved.     | IPBOOTP_MIB_GET_INPUT_DATA                                                       | See section 2.2.1.2.152 for details on how to populate                           |
	//	|                                 |                                                  |                                                                                  | IPBOOTP_MIB_GET_INPUT_DATA to retrieve different types of MIB entries.           |
	//	|                                 |                                                  |                                                                                  | When this method returns pMibOutEntry, pInfoStruct MUST be cast to               |
	//	|                                 |                                                  |                                                                                  | IPBOOTP_MIB_GET_OUTPUT_DATA. See section 2.2.1.2.157 on how to interpret the     |
	//	|                                 |                                                  |                                                                                  | data returned.                                                                   |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_NAT 0x81372715            | An entry in the NAT MIB needs to be retrieved.   | IP_NAT_MIB_QUERY                                                                 | See IP_NAT_MIB_QUERY (section 2.2.1.2.182) for details on how to populate        |
	//	|                                 |                                                  |                                                                                  | this structure to retrieve the entries required. The values are returned in      |
	//	|                                 |                                                  |                                                                                  | IP_NAT_MIB_QUERY.                                                                |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_DNS_PROXY 0x81372713      | An entry in DNS Proxy MIB needs to be retrieved. | IP_DNS_PROXY_MIB_QUERY                                                           | See section 2.2.1.2.186 for details on how to populate IP_DNS_PROXY_MIB_QUERY.   |
	//	|                                 |                                                  |                                                                                  | The values are returned in IP_DNS_PROXY_MIB_QUERY.                               |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_DHCP_ALLOCATOR 0x81372714 | An entry in auto DHCP MIB is retrieved.          | IP_AUTO_DHCP_MIB_QUERY                                                           | See section 2.2.1.2.188 for details on how to populate IP_AUTO_DHCP_MIB_QUERY.   |
	//	|                                 |                                                  |                                                                                  | The values are returned in the same structure.                                   |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                 |                                                  |                                                                                  |                                                                                  |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IP_OSPF<296> 0x0000000D      | An entry in OSPF MIB needs to be retrieved.      | MIB_DA_MSG MIB_DA_MSG OSPF_GLOBAL_PARAM OSPF_ROUTE_FILTER_INFO                   |                                                                                  |
	//	|                                 |                                                  | OSPF_PROTO_FILTER_INFO OSPF_AREA_PARAM OSPF_AREA_RANGE_PARAM                     |                                                                                  |
	//	|                                 |                                                  | OSPF_VIRT_INTERFACE_PARAM OSPF_NBMA_NEIGHBOR_PARAM                               |                                                                                  |
	//	+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// If dwPid is PID_IPV6 and if dwRoutingPid is not IPRTRMGR_PID (0x00002710), then dwRoutingPid
	// MUST be one of the entries in the Value column and pMibInEntry MUST be the corresponding
	// entry in the Structure to be pointed to by pMibInEntry in the following table. The
	// following specified routing protocols are valid only if the protocol is already initialized
	// by RRAS for IPv6. RRAS initializes a protocol when an RTR_TOC_ENTRY with the InfoType
	// corresponding to the protocol is present in the global configuration. The RRouterInterfaceTransportSetGlobalInfo
	// method enables specifying the global configuration through the pGlobalInfo member
	// of pInfoStruct.
	//
	//	+-------------------------+-----------------------------------------------------------+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                         |                                                           |       STRUCTURE TO BE POINTED TO BY       |                                                                                  |
	//	|          VALUE          |                          MEANING                          |                PMIBINENTRY                |                                     REMARKS                                      |
	//	|                         |                                                           |                                           |                                                                                  |
	//	+-------------------------+-----------------------------------------------------------+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+-----------------------------------------------------------+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| MS_IPV6_DHCP 0x000003E7 | An entry in DHCPv6 Relay agent MIB needs to be retrieved. | DHCPV6R_MIB_GET_INPUT_DATA                | See section 2.2.1.2.160 on how to populate DHCPV6R_MIB_GET_INPUT_DATA. The       |
	//	|                         |                                                           |                                           | values are retuned in DHCPV6R_MIB_GET_OUTPUT_DATA.                               |
	//	+-------------------------+-----------------------------------------------------------+-------------------------------------------+----------------------------------------------------------------------------------+
	RoutingPID uint32 `idl:"name:dwRoutingPid" json:"routing_pid"`
	// pInfoStuct: The same as in RMIBEntryGet (section 3.1.4.30).
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntryGetFirstRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetFirstOperation) *xxx_MIBEntryGetFirstOperation {
	if op == nil {
		op = &xxx_MIBEntryGetFirstOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntryGetFirstRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetFirstOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntryGetFirstRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntryGetFirstRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetFirstOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntryGetFirstResponse structure represents the RMIBEntryGetFirst operation response
type MIBEntryGetFirstResponse struct {
	// pInfoStuct: The same as in RMIBEntryGet (section 3.1.4.30).
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	// Return: The RMIBEntryGetFirst return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntryGetFirstResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetFirstOperation) *xxx_MIBEntryGetFirstOperation {
	if op == nil {
		op = &xxx_MIBEntryGetFirstOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MIBEntryGetFirstResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetFirstOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MIBEntryGetFirstResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntryGetFirstResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetFirstOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntryGetNextOperation structure represents the RMIBEntryGetNext operation
type xxx_MIBEntryGetNextOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntryGetNextOperation) OpNum() int { return 31 }

func (o *xxx_MIBEntryGetNextOperation) OpName() string { return "/dimsvc/v0/RMIBEntryGetNext" }

func (o *xxx_MIBEntryGetNextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetNextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetNextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetNextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetNextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetNextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBEntryGetNextRequest structure represents the RMIBEntryGetNext operation request
type MIBEntryGetNextRequest struct {
	// dwPid: The same as in RMIBEntryGet (section 3.1.4.30).
	PID uint32 `idl:"name:dwPid" json:"pid"`
	// dwRoutingPid: The same as in RMIBEntryGetFirst (section 3.1.4.31).
	RoutingPID uint32 `idl:"name:dwRoutingPid" json:"routing_pid"`
	// pInfoStuct: The same as in RMIBEntryGet (section 3.1.4.30).
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntryGetNextRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetNextOperation) *xxx_MIBEntryGetNextOperation {
	if op == nil {
		op = &xxx_MIBEntryGetNextOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntryGetNextRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetNextOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntryGetNextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntryGetNextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetNextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntryGetNextResponse structure represents the RMIBEntryGetNext operation response
type MIBEntryGetNextResponse struct {
	// pInfoStuct: The same as in RMIBEntryGet (section 3.1.4.30).
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	// Return: The RMIBEntryGetNext return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntryGetNextResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetNextOperation) *xxx_MIBEntryGetNextOperation {
	if op == nil {
		op = &xxx_MIBEntryGetNextOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MIBEntryGetNextResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetNextOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MIBEntryGetNextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntryGetNextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetNextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBGetTrapInfoOperation structure represents the RMIBGetTrapInfo operation
type xxx_MIBGetTrapInfoOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBGetTrapInfoOperation) OpNum() int { return 32 }

func (o *xxx_MIBGetTrapInfoOperation) OpName() string { return "/dimsvc/v0/RMIBGetTrapInfo" }

func (o *xxx_MIBGetTrapInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBGetTrapInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBGetTrapInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBGetTrapInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBGetTrapInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBGetTrapInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBGetTrapInfoRequest structure represents the RMIBGetTrapInfo operation request
type MIBGetTrapInfoRequest struct {
	// dwPid: Specifies the transport protocol. The value of this parameter MUST be one
	// from the following table.<300>
	//
	//	+---------------------+---------------------------------------+
	//	|                     |                                       |
	//	|        VALUE        |                MEANING                |
	//	|                     |                                       |
	//	+---------------------+---------------------------------------+
	//	+---------------------+---------------------------------------+
	//	| PID_IPX 0x0000002B  | An IPX MIB entry is to be retrieved.  |
	//	+---------------------+---------------------------------------+
	//	| PID_IP 0x00000021   | An IPv4 MIB entry is to be retrieved. |
	//	+---------------------+---------------------------------------+
	//	| PID_IPV6 0x00000057 | An IPv6 entry is to be retrieved.     |
	//	+---------------------+---------------------------------------+
	PID uint32 `idl:"name:dwPid" json:"pid"`
	// dwRoutingPid: Specifies the routing protocol that exported the variable.
	RoutingPID uint32 `idl:"name:dwRoutingPid" json:"routing_pid"`
	// pInfoStruct: Pointer to an opaque data structure DIM_MIB_ENTRY_CONTAINER (section
	// 2.2.1.2.19). The data structure's format is determined by the router manager or router
	// manager client that is servicing the call. The data structure MUST contain information
	// that specifies the variable being created and the value to assign to the variable.
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *MIBGetTrapInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBGetTrapInfoOperation) *xxx_MIBGetTrapInfoOperation {
	if op == nil {
		op = &xxx_MIBGetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBGetTrapInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBGetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBGetTrapInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBGetTrapInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBGetTrapInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBGetTrapInfoResponse structure represents the RMIBGetTrapInfo operation response
type MIBGetTrapInfoResponse struct {
	// pInfoStruct: Pointer to an opaque data structure DIM_MIB_ENTRY_CONTAINER (section
	// 2.2.1.2.19). The data structure's format is determined by the router manager or router
	// manager client that is servicing the call. The data structure MUST contain information
	// that specifies the variable being created and the value to assign to the variable.
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RMIBGetTrapInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBGetTrapInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBGetTrapInfoOperation) *xxx_MIBGetTrapInfoOperation {
	if op == nil {
		op = &xxx_MIBGetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MIBGetTrapInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBGetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MIBGetTrapInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBGetTrapInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBGetTrapInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBSetTrapInfoOperation structure represents the RMIBSetTrapInfo operation
type xxx_MIBSetTrapInfoOperation struct {
	PID             uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID      uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Event           uint64                   `idl:"name:hEvent" json:"event"`
	ClientProcessID uint32                   `idl:"name:dwClientProcessId" json:"client_process_id"`
	Info            *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	Return          uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBSetTrapInfoOperation) OpNum() int { return 33 }

func (o *xxx_MIBSetTrapInfoOperation) OpName() string { return "/dimsvc/v0/RMIBSetTrapInfo" }

func (o *xxx_MIBSetTrapInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBSetTrapInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.Event)); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientProcessID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBSetTrapInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.Event)); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientProcessID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBSetTrapInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBSetTrapInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBSetTrapInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBSetTrapInfoRequest structure represents the RMIBSetTrapInfo operation request
type MIBSetTrapInfoRequest struct {
	// dwPid: Specifies the protocol ID that exported the variable.
	PID uint32 `idl:"name:dwPid" json:"pid"`
	// dwRoutingPid: Specifies the routing protocol that exported the variable.
	RoutingPID uint32 `idl:"name:dwRoutingPid" json:"routing_pid"`
	// hEvent: A handle to an event that is signaled when a trap needs to be issued. This
	// MUST be a handle to an event on the RRAS server which is created within the process
	// specified by dwClientProcessId which can be signaled whenever a trap needs to be
	// issued. Otherwise the method SHOULD fail with an appropriate error code as specified
	// in [MS-ERREF].
	Event uint64 `idl:"name:hEvent" json:"event"`
	// dwClientProcessId: The current process identifier.
	ClientProcessID uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
	// pInfoStruct: Pointer to an opaque data structure DIM_MIB_ENTRY_CONTAINER (section
	// 2.2.1.2.19). The data structure's format is determined by the router manager or router
	// manager client that is servicing the call. The data structure MUST contain information
	// that specifies the variable being created and the value to assign to the variable.
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *MIBSetTrapInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBSetTrapInfoOperation) *xxx_MIBSetTrapInfoOperation {
	if op == nil {
		op = &xxx_MIBSetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Event = o.Event
	op.ClientProcessID = o.ClientProcessID
	op.Info = o.Info
	return op
}

func (o *MIBSetTrapInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBSetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Event = op.Event
	o.ClientProcessID = op.ClientProcessID
	o.Info = op.Info
}
func (o *MIBSetTrapInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBSetTrapInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBSetTrapInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBSetTrapInfoResponse structure represents the RMIBSetTrapInfo operation response
type MIBSetTrapInfoResponse struct {
	// pInfoStruct: Pointer to an opaque data structure DIM_MIB_ENTRY_CONTAINER (section
	// 2.2.1.2.19). The data structure's format is determined by the router manager or router
	// manager client that is servicing the call. The data structure MUST contain information
	// that specifies the variable being created and the value to assign to the variable.
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RMIBSetTrapInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBSetTrapInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBSetTrapInfoOperation) *xxx_MIBSetTrapInfoOperation {
	if op == nil {
		op = &xxx_MIBSetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MIBSetTrapInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBSetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MIBSetTrapInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBSetTrapInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBSetTrapInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionNotificationOperation structure represents the RRasAdminConnectionNotification operation
type xxx_ConnectionNotificationOperation struct {
	Register          uint32 `idl:"name:fRegister" json:"register"`
	ClientProcessID   uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
	EventNotification uint64 `idl:"name:hEventNotification" json:"event_notification"`
	Return            uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionNotificationOperation) OpNum() int { return 34 }

func (o *xxx_ConnectionNotificationOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionNotification"
}

func (o *xxx_ConnectionNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// fRegister {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Register); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientProcessID); err != nil {
			return err
		}
	}
	// hEventNotification {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.EventNotification)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// fRegister {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Register); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientProcessID); err != nil {
			return err
		}
	}
	// hEventNotification {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.EventNotification)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ConnectionNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionNotificationRequest structure represents the RRasAdminConnectionNotification operation request
type ConnectionNotificationRequest struct {
	// fRegister: This is of type DWORD and SHOULD be set to 1, if notifications are required
	// when an interface is connected or disconnected. Set to 0 to unregister notifications.
	Register uint32 `idl:"name:fRegister" json:"register"`
	// dwClientProcessId: The current process identifier that determines where the function
	// is called from.
	ClientProcessID uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
	// hEventNotification: A handle to an event that is signaled after the connection is
	// connected or disconnected. This MUST be a handle to an event on the RRASM server
	// which is created within the process specified by dwClientProcessId which can be signaled
	// on interface connect and disconnect events. Otherwise the method SHOULD fail with
	// an appropriate error code as specified in [MS-ERREF].
	EventNotification uint64 `idl:"name:hEventNotification" json:"event_notification"`
}

func (o *ConnectionNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionNotificationOperation) *xxx_ConnectionNotificationOperation {
	if op == nil {
		op = &xxx_ConnectionNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Register = o.Register
	op.ClientProcessID = o.ClientProcessID
	op.EventNotification = o.EventNotification
	return op
}

func (o *ConnectionNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionNotificationOperation) {
	if o == nil {
		return
	}
	o.Register = op.Register
	o.ClientProcessID = op.ClientProcessID
	o.EventNotification = op.EventNotification
}
func (o *ConnectionNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionNotificationResponse structure represents the RRasAdminConnectionNotification operation response
type ConnectionNotificationResponse struct {
	// Return: The RRasAdminConnectionNotification return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionNotificationOperation) *xxx_ConnectionNotificationOperation {
	if op == nil {
		op = &xxx_ConnectionNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ConnectionNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionNotificationOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ConnectionNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SendUserMessageOperation structure represents the RRasAdminSendUserMessage operation
type xxx_SendUserMessageOperation struct {
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
	Message    string `idl:"name:lpwszMessage;string" json:"message"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SendUserMessageOperation) OpNum() int { return 35 }

func (o *xxx_SendUserMessageOperation) OpName() string { return "/dimsvc/v0/RRasAdminSendUserMessage" }

func (o *xxx_SendUserMessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendUserMessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	// lpwszMessage {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Message); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendUserMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	// lpwszMessage {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Message); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendUserMessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendUserMessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SendUserMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SendUserMessageRequest structure represents the RRasAdminSendUserMessage operation request
type SendUserMessageRequest struct {
	// hDimConnection: This is of type DWORD and MUST be set to a particular connection
	// identifier for which the connection information is required. Obtain this handle by
	// calling RRasAdminConnectionEnum (section 3.1.4.2). Since RRASM server does not maintain
	// the connection handles, the RRAS server SHOULD check and ensure that this handle
	// represents a valid interface handle.
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
	// lpwszMessage: A pointer to a Unicode string that specifies the message to the user.
	// It MUST NOT be NULL.
	Message string `idl:"name:lpwszMessage;string" json:"message"`
}

func (o *SendUserMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_SendUserMessageOperation) *xxx_SendUserMessageOperation {
	if op == nil {
		op = &xxx_SendUserMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Connection = o.Connection
	op.Message = o.Message
	return op
}

func (o *SendUserMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_SendUserMessageOperation) {
	if o == nil {
		return
	}
	o.Connection = op.Connection
	o.Message = op.Message
}
func (o *SendUserMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SendUserMessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendUserMessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SendUserMessageResponse structure represents the RRasAdminSendUserMessage operation response
type SendUserMessageResponse struct {
	// Return: The RRasAdminSendUserMessage return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SendUserMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_SendUserMessageOperation) *xxx_SendUserMessageOperation {
	if op == nil {
		op = &xxx_SendUserMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SendUserMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_SendUserMessageOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SendUserMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SendUserMessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendUserMessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterDeviceEnumOperation structure represents the RRouterDeviceEnum operation
type xxx_RouterDeviceEnumOperation struct {
	Level        uint32                      `idl:"name:dwLevel" json:"level"`
	Info         *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	TotalEntries uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Return       uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterDeviceEnumOperation) OpNum() int { return 36 }

func (o *xxx_RouterDeviceEnumOperation) OpName() string { return "/dimsvc/v0/RRouterDeviceEnum" }

func (o *xxx_RouterDeviceEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterDeviceEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterDeviceEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterDeviceEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterDeviceEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterDeviceEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterDeviceEnumRequest structure represents the RRouterDeviceEnum operation request
type RouterDeviceEnumRequest struct {
	// dwLevel: This is of type DWORD and MUST be set to the following value.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | After the function returns the memory pointed to by pInfoStruct->pBuffer it is   |
	//	|       | interpreted as array of MPR_DEVICE_0 (section 2.2.1.2.85). The size of the array |
	//	|       | is determined by lpdwEntriesRead.                                                |
	//	+-------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1),
	// and pInfoStruct.dwBufferSize is initialized to zero (0). Upon successful return,
	// the pInfoStruct->pBuffer is a typecast array of MPR_DEVICE_0 (section 2.2.1.2.85)
	// and the array size is determined by the value to pointer lpdwTotalEntries.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// lpdwTotalEntries: This is a pointer to type DWORD and receives the total number of
	// devices that have been enumerated.
	TotalEntries uint32 `idl:"name:lpdwTotalEntries" json:"total_entries"`
}

func (o *RouterDeviceEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterDeviceEnumOperation) *xxx_RouterDeviceEnumOperation {
	if op == nil {
		op = &xxx_RouterDeviceEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.TotalEntries = o.TotalEntries
	return op
}

func (o *RouterDeviceEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterDeviceEnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.TotalEntries = op.TotalEntries
}
func (o *RouterDeviceEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterDeviceEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterDeviceEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterDeviceEnumResponse structure represents the RRouterDeviceEnum operation response
type RouterDeviceEnumResponse struct {
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1),
	// and pInfoStruct.dwBufferSize is initialized to zero (0). Upon successful return,
	// the pInfoStruct->pBuffer is a typecast array of MPR_DEVICE_0 (section 2.2.1.2.85)
	// and the array size is determined by the value to pointer lpdwTotalEntries.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// lpdwTotalEntries: This is a pointer to type DWORD and receives the total number of
	// devices that have been enumerated.
	TotalEntries uint32 `idl:"name:lpdwTotalEntries" json:"total_entries"`
	// Return: The RRouterDeviceEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterDeviceEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterDeviceEnumOperation) *xxx_RouterDeviceEnumOperation {
	if op == nil {
		op = &xxx_RouterDeviceEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.TotalEntries = o.TotalEntries
	op.Return = o.Return
	return op
}

func (o *RouterDeviceEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterDeviceEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.TotalEntries = op.TotalEntries
	o.Return = op.Return
}
func (o *RouterDeviceEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterDeviceEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterDeviceEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportCreateOperation structure represents the RRouterInterfaceTransportCreate operation
type xxx_RouterInterfaceTransportCreateOperation struct {
	TransportID   uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	TransportName string                    `idl:"name:lpwsTransportName;string" json:"transport_name"`
	Info          *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	DLLPath       string                    `idl:"name:lpwsDLLPath;string" json:"dll_path"`
	Return        uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportCreateOperation) OpNum() int { return 37 }

func (o *xxx_RouterInterfaceTransportCreateOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportCreate"
}

func (o *xxx_RouterInterfaceTransportCreateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportCreateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// lpwsTransportName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.TransportName); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpwsDLLPath {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DLLPath); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportCreateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// lpwsTransportName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.TransportName); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpwsDLLPath {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DLLPath); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportCreateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportCreateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceTransportCreateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportCreateRequest structure represents the RRouterInterfaceTransportCreate operation request
type RouterInterfaceTransportCreateRequest struct {
	// dwTransportId: Specifies the transport for which the information is set (IPX, IPv4,
	// or IPv6). It MUST be set to one of the following values.<307>
	//
	//	+---------------------+---------------+
	//	|                     |               |
	//	|        VALUE        |    MEANING    |
	//	|                     |               |
	//	+---------------------+---------------+
	//	+---------------------+---------------+
	//	| PID_IPX 0x0000002B  | IPX protocol  |
	//	+---------------------+---------------+
	//	| PID_IP 0x00000021   | IPv4 protocol |
	//	+---------------------+---------------+
	//	| PID_IPV6 0x00000057 | IPv6 protocol |
	//	+---------------------+---------------+
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
	// lpwsTransportName: Pointer to a null-terminated Unicode string that specifies the
	// name of the transport being added. If this parameter is not specified, the dwTransportId
	// parameter is converted into a string and used as the transport name.
	TransportName string `idl:"name:lpwsTransportName;string" json:"transport_name"`
	// pInfoStruct: Pointer to a DIM_INTERFACE_CONTAINER (section 2.2.1.2.2). This MUST
	// NOT be NULL. The pGlobalInfo member of the DIM_INTERFACE_CONTAINER MUST NOT be NULL
	// and MUST point to a valid RTR_INFO_BLOCK_HEADER (section 2.2.1.2.3) and RTR_TOC_ENTRY
	// (section 2.2.1.2.4).
	//
	// If dwTransportId is PID_IP or PID_IPV6, one or more of the following entries MUST
	// be set in the InfoType field in the RTR_TOC_ENTRY (section 2.2.1.2.4) structure while
	// passing to the method.
	//
	//	+-------------------------------------+--------------------------------------------------------+----------------------------------------+
	//	|                                     |                      MEANING AND                       |        STRUCTURE POINTED TO AT         |
	//	|                VALUE                |                       CONDITIONS                       |                 OFFSET                 |
	//	|                                     |                                                        |                                        |
	//	+-------------------------------------+--------------------------------------------------------+----------------------------------------+
	//	+-------------------------------------+--------------------------------------------------------+----------------------------------------+
	//	| IP_PROT_PRIORITY_INFO 0xFFFF0006    | IPv4 and IPv6 route priority information is specified. | PRIORITY_INFO (section 2.2.1.2.12)     |
	//	+-------------------------------------+--------------------------------------------------------+----------------------------------------+
	//	| IP_PROT_PRIORITY_INFO_EX 0xFFFF0017 | IPv4 and IPv6 route priority information is specified. | PRIORITY_INFO_EX (section 2.2.1.2.266) |
	//	+-------------------------------------+--------------------------------------------------------+----------------------------------------+
	//	| IPV6_GLOBAL_INFO 0xFFFF000F         | Global lPv6 logging and filtering information.         | GLOBAL_INFO (section 2.2.1.2.9)        |
	//	+-------------------------------------+--------------------------------------------------------+----------------------------------------+
	//	| IP_GLOBAL_INFO 0xFFFF0003           | Global lPv4 logging and filtering information.         | GLOBAL_INFO (section 2.2.1.2.9)        |
	//	+-------------------------------------+--------------------------------------------------------+----------------------------------------+
	Info *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	// lpwsDLLPath: Pointer to a null-terminated Unicode string that specifies the name
	// of the router manager DLL for the specified transport. If this name is specified,
	// the function sets the DLL path for this transport to this name.<308>
	DLLPath string `idl:"name:lpwsDLLPath;string" json:"dll_path"`
}

func (o *RouterInterfaceTransportCreateRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportCreateOperation) *xxx_RouterInterfaceTransportCreateOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.TransportID = o.TransportID
	op.TransportName = o.TransportName
	op.Info = o.Info
	op.DLLPath = o.DLLPath
	return op
}

func (o *RouterInterfaceTransportCreateRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportCreateOperation) {
	if o == nil {
		return
	}
	o.TransportID = op.TransportID
	o.TransportName = op.TransportName
	o.Info = op.Info
	o.DLLPath = op.DLLPath
}
func (o *RouterInterfaceTransportCreateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportCreateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportCreateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportCreateResponse structure represents the RRouterInterfaceTransportCreate operation response
type RouterInterfaceTransportCreateResponse struct {
	// Return: The RRouterInterfaceTransportCreate return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportCreateResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportCreateOperation) *xxx_RouterInterfaceTransportCreateOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportCreateResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportCreateOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceTransportCreateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportCreateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportCreateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceDeviceGetInfoOperation structure represents the RRouterInterfaceDeviceGetInfo operation
type xxx_RouterInterfaceDeviceGetInfoOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Index     uint32                      `idl:"name:dwIndex" json:"index"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) OpNum() int { return 38 }

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceDeviceGetInfo"
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceDeviceGetInfoRequest structure represents the RRouterInterfaceDeviceGetInfo operation request
type RouterInterfaceDeviceGetInfoRequest struct {
	// dwLevel: This is of type DWORD and MUST be set to any of the following values.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	|     0 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as MPR_DEVICE_0 (section 2.2.1.2.85).                                |
	//	+-------+----------------------------------------------------------------------------------+
	//	|     1 | After the function returns, the memory pointed to by pInfoStruct->pBuffer is     |
	//	|       | interpreted as MPR_DEVICE_1 (section 2.2.1.2.86).                                |
	//	+-------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1)
	// and pInfoStruct.dwBufferSize SHOULD be initialized to zero (0). Upon successful return,
	// the pInfoStruct->pBuffer is a typecast array of MPR_DEVICE_0 (section 2.2.1.2.85)
	// or MPR_DEVICE_1 (section 2.2.1.2.86), based on the dwLevel value.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// dwIndex: Specifies the one-based index of the device. A multilinked demand-dial interface
	// uses multiple devices.
	Index uint32 `idl:"name:dwIndex" json:"index"`
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Since the RRASM
	// server does not maintain interface handles, the RRAS server SHOULD check and ensure
	// that this handle represents a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceDeviceGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeviceGetInfoOperation) *xxx_RouterInterfaceDeviceGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeviceGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Index = o.Index
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceDeviceGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeviceGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Index = op.Index
	o.Interface = op.Interface
}
func (o *RouterInterfaceDeviceGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceDeviceGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeviceGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceDeviceGetInfoResponse structure represents the RRouterInterfaceDeviceGetInfo operation response
type RouterInterfaceDeviceGetInfoResponse struct {
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1)
	// and pInfoStruct.dwBufferSize SHOULD be initialized to zero (0). Upon successful return,
	// the pInfoStruct->pBuffer is a typecast array of MPR_DEVICE_0 (section 2.2.1.2.85)
	// or MPR_DEVICE_1 (section 2.2.1.2.86), based on the dwLevel value.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRouterInterfaceDeviceGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceDeviceGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeviceGetInfoOperation) *xxx_RouterInterfaceDeviceGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeviceGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceDeviceGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeviceGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RouterInterfaceDeviceGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceDeviceGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeviceGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceDeviceSetInfoOperation structure represents the RRouterInterfaceDeviceSetInfo operation
type xxx_RouterInterfaceDeviceSetInfoOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Index     uint32                      `idl:"name:dwIndex" json:"index"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) OpNum() int { return 39 }

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceDeviceSetInfo"
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceDeviceSetInfoRequest structure represents the RRouterInterfaceDeviceSetInfo operation request
type RouterInterfaceDeviceSetInfoRequest struct {
	// dwLevel: This is of type DWORD and MUST be set to one of the following values.
	//
	//	+-------+---------------------------------------------------------------------------------+
	//	|       |                                                                                 |
	//	| VALUE |                                     MEANING                                     |
	//	|       |                                                                                 |
	//	+-------+---------------------------------------------------------------------------------+
	//	+-------+---------------------------------------------------------------------------------+
	//	|     0 | pInfoStruct->pBuffer MUST be set to point to MPR_DEVICE_0 (section 2.2.1.2.85). |
	//	+-------+---------------------------------------------------------------------------------+
	//	|     1 | pInfoStruct->pBuffer MUST be set to point to MPR_DEVICE_1 (section 2.2.1.2.86). |
	//	+-------+---------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: This is a pointer of type DIM_INFORMATION_CONTAINER (section 2.2.1.2.1)
	// and MUST be set to the following, based on dwLevel. The szDeviceName MUST be one
	// of the devices as specified by RRouterDeviceEnum. The szDeviceType specified in the
	// structure SHOULD be validated against the actual type of the device specified by
	// szDeviceName. If the szDeviceType is incorrect, the actual device type SHOULD be
	// used by the server instead.
	//
	//	+---------+-----------------------------------------------------------------+----------------------------------------------+
	//	|         |                                                                 |                                              |
	//	| DWLEVEL |                        PINFOSTRUCT->SIZE                        |             PINFOSTRUCT->PBUFFER             |
	//	|         |                                                                 |                                              |
	//	+---------+-----------------------------------------------------------------+----------------------------------------------+
	//	+---------+-----------------------------------------------------------------+----------------------------------------------+
	//	|       0 | Should be set to the size of MPR_DEVICE_0 (section 2.2.1.2.85). | This MUST be set to point to MPR_DEVICE_0.   |
	//	+---------+-----------------------------------------------------------------+----------------------------------------------+
	//	|       1 | Should be set to the size of MPR_DEVICE_1 (section 2.2.1.2.86). | This MUST be set to a point to MPR_DEVICE_1. |
	//	+---------+-----------------------------------------------------------------+----------------------------------------------+
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// dwIndex: Specifies the 1-based index of the device. A multilinked demand-dial interface
	// uses multiple devices. If the device information specified for dwIndex 1 is either
	// a PPPoE interface or a VPN interface, any other device information, specified (with
	// dwIndex > 1), SHOULD be ignored. If the device information specified for dwIndex
	// 1 is a serial or ISDN device, the connection SHOULD be treated as a multilink-capable
	// connection. Any other device information, specified (with dwIndex > 1), SHOULD be
	// treated as device information for the individual links. Device information with type
	// szDeviceType, which is neither modem nor ISDN, SHOULD<311> be ignored.
	Index uint32 `idl:"name:dwIndex" json:"index"`
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Since the RRASM
	// server does not maintain the interface handles, the RRAS server implementation SHOULD
	// check and ensure that this handle represents a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceDeviceSetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeviceSetInfoOperation) *xxx_RouterInterfaceDeviceSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeviceSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Index = o.Index
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceDeviceSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeviceSetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Index = op.Index
	o.Interface = op.Interface
}
func (o *RouterInterfaceDeviceSetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceDeviceSetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeviceSetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceDeviceSetInfoResponse structure represents the RRouterInterfaceDeviceSetInfo operation response
type RouterInterfaceDeviceSetInfoResponse struct {
	// Return: The RRouterInterfaceDeviceSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceDeviceSetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeviceSetInfoOperation) *xxx_RouterInterfaceDeviceSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeviceSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceDeviceSetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeviceSetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceDeviceSetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceDeviceSetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeviceSetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceSetCredentialsExOperation structure represents the RRouterInterfaceSetCredentialsEx operation
type xxx_RouterInterfaceSetCredentialsExOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) OpNum() int { return 40 }

func (o *xxx_RouterInterfaceSetCredentialsExOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceSetCredentialsEx"
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceSetCredentialsExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceSetCredentialsExRequest structure represents the RRouterInterfaceSetCredentialsEx operation request
type RouterInterfaceSetCredentialsExRequest struct {
	// dwLevel: Specifies the format of the credentials information. This parameter MUST
	// be one of the following values: 0x00000000, 0x0000001, 0x0000002, or 0x00000003.
	// A value of zero (0) indicates that this is EAP information, and the MPR_CREDENTIALSEX_1
	// structure MUST contain EAPTLS_USER_PROPERTIES (section 2.2.1.2.128). A value of one,
	// two, or three (these can be used interchangeably) indicate that the credentials information
	// is a preshared key and is formatted as an MPR_CREDENTIALSEX_1 structure. The preshared
	// key specifies the preshared key to be used with IPsec for L2TP over IPsec connections.
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: Pointer to the DIM_INFORMATION_CONTAINER (section 2.2.1.2.1), the pBuffer
	// pointer of which points to the MPR_CREDENTIALSEX_1 structure that contains the credential
	// information to be set for the interface.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// hInterface: Handle to the interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). If dwLevel is
	// 0x0000002 and hInterface is NULL, the preshared key is used for L2TP. Since the RRASM
	// server does not maintain the interface handles, the RRAS server SHOULD check and
	// ensure that this handle is a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceSetCredentialsExRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsExOperation) *xxx_RouterInterfaceSetCredentialsExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCredentialsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceSetCredentialsExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsExOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Interface = op.Interface
}
func (o *RouterInterfaceSetCredentialsExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceSetCredentialsExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCredentialsExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceSetCredentialsExResponse structure represents the RRouterInterfaceSetCredentialsEx operation response
type RouterInterfaceSetCredentialsExResponse struct {
	// Return: The RRouterInterfaceSetCredentialsEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceSetCredentialsExResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsExOperation) *xxx_RouterInterfaceSetCredentialsExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCredentialsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceSetCredentialsExResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceSetCredentialsExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceSetCredentialsExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCredentialsExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceGetCredentialsExOperation structure represents the RRouterInterfaceGetCredentialsEx operation
type xxx_RouterInterfaceGetCredentialsExOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) OpNum() int { return 41 }

func (o *xxx_RouterInterfaceGetCredentialsExOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceGetCredentialsEx"
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceGetCredentialsExRequest structure represents the RRouterInterfaceGetCredentialsEx operation request
type RouterInterfaceGetCredentialsExRequest struct {
	// dwLevel: Specifies the format of credentials information. This parameter takes values
	// 0x00000000, 0x00000001, 0x00000002, or 0x00000003. An error other than one of the
	// errors in the return values table is returned for other values. A value of 0x00000000
	// indicates that the credentials information is about the EAP configuration. If the
	// interface does not have EAP information, the dwSize field of the MPR_CREDENTIALSEX_1
	// (section 2.2.1.2.87)structure MUST be set to zero (0). Otherwise, the MPR_CREDENTIALSEX_1
	// structure MUST contain EAPTLS_USER_PROPERTIES (section 2.2.1.2.128). A value of 0x00000001,
	// 0x00000002, or 0x00000003 indicates that the credentials information is a preshared
	// key. If the interface does not have a preshared key configured, an error other than
	// one of the errors in the return values table is returned. Otherwise, the preshared
	// key is formatted as an MPR_CREDENTIALSEX_1 structure.
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: Pointer to a MPR_CREDENTIALSEX_1 structure that contains the preshared
	// key or EAP information for the interface. When the method completes successfully,
	// the client SHOULD free this memory.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// hInterface: A handle to the interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). The interface
	// MUST be of type ROUTER_IF_TYPE_FULL_ROUTER if dwLevel is 0x00000000.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceGetCredentialsExRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsExOperation) *xxx_RouterInterfaceGetCredentialsExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCredentialsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceGetCredentialsExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsExOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Interface = op.Interface
}
func (o *RouterInterfaceGetCredentialsExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceGetCredentialsExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCredentialsExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceGetCredentialsExResponse structure represents the RRouterInterfaceGetCredentialsEx operation response
type RouterInterfaceGetCredentialsExResponse struct {
	// pInfoStruct: Pointer to a MPR_CREDENTIALSEX_1 structure that contains the preshared
	// key or EAP information for the interface. When the method completes successfully,
	// the client SHOULD free this memory.
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRouterInterfaceGetCredentialsEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceGetCredentialsExResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsExOperation) *xxx_RouterInterfaceGetCredentialsExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCredentialsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetCredentialsExResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsExOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RouterInterfaceGetCredentialsExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceGetCredentialsExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCredentialsExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionRemoveQuarantineOperation structure represents the RRasAdminConnectionRemoveQuarantine operation
type xxx_ConnectionRemoveQuarantineOperation struct {
	Connection  uint32 `idl:"name:hRasConnection" json:"connection"`
	IsIPAddress bool   `idl:"name:fIsIpAddress" json:"is_ip_address"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionRemoveQuarantineOperation) OpNum() int { return 42 }

func (o *xxx_ConnectionRemoveQuarantineOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionRemoveQuarantine"
}

func (o *xxx_ConnectionRemoveQuarantineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionRemoveQuarantineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	// fIsIpAddress {in} (1:{alias=BOOL}(int32))
	{
		if !o.IsIPAddress {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ConnectionRemoveQuarantineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	// fIsIpAddress {in} (1:{alias=BOOL}(int32))
	{
		var _bIsIPAddress int32
		if err := w.ReadData(&_bIsIPAddress); err != nil {
			return err
		}
		o.IsIPAddress = _bIsIPAddress != 0
	}
	return nil
}

func (o *xxx_ConnectionRemoveQuarantineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionRemoveQuarantineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ConnectionRemoveQuarantineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionRemoveQuarantineRequest structure represents the RRasAdminConnectionRemoveQuarantine operation request
type ConnectionRemoveQuarantineRequest struct {
	// hRasConnection: This is of type DWORD and SHOULD be set to a particular connection
	// identifier for which the connection information is required. Obtain this handle by
	// calling RRasAdminConnectionEnum (section 3.1.4.2). Since the RRASM server does not
	// maintain the connection handles, the RRAS server SHOULD check and ensure that this
	// handle is a valid connection handle.
	Connection uint32 `idl:"name:hRasConnection" json:"connection"`
	// fIsIpAddress: Specifies a Boolean value that indicates whether the hRasConnection
	// parameter specifies the IP address of the client for which to remove the quarantine
	// filters. If this parameter is a TRUE value, hRasConnection specifies an IP address.
	// Otherwise, hRasConnection specifies a handle to a connection.
	IsIPAddress bool `idl:"name:fIsIpAddress" json:"is_ip_address"`
}

func (o *ConnectionRemoveQuarantineRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionRemoveQuarantineOperation) *xxx_ConnectionRemoveQuarantineOperation {
	if op == nil {
		op = &xxx_ConnectionRemoveQuarantineOperation{}
	}
	if o == nil {
		return op
	}
	op.Connection = o.Connection
	op.IsIPAddress = o.IsIPAddress
	return op
}

func (o *ConnectionRemoveQuarantineRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionRemoveQuarantineOperation) {
	if o == nil {
		return
	}
	o.Connection = op.Connection
	o.IsIPAddress = op.IsIPAddress
}
func (o *ConnectionRemoveQuarantineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionRemoveQuarantineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionRemoveQuarantineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionRemoveQuarantineResponse structure represents the RRasAdminConnectionRemoveQuarantine operation response
type ConnectionRemoveQuarantineResponse struct {
	// Return: The RRasAdminConnectionRemoveQuarantine return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionRemoveQuarantineResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionRemoveQuarantineOperation) *xxx_ConnectionRemoveQuarantineOperation {
	if op == nil {
		op = &xxx_ConnectionRemoveQuarantineOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ConnectionRemoveQuarantineResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionRemoveQuarantineOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ConnectionRemoveQuarantineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionRemoveQuarantineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionRemoveQuarantineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerSetInfoOperation structure represents the RMprAdminServerSetInfo operation
type xxx_ServerSetInfoOperation struct {
	Level  uint32                      `idl:"name:dwLevel" json:"level"`
	Info   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerSetInfoOperation) OpNum() int { return 43 }

func (o *xxx_ServerSetInfoOperation) OpName() string { return "/dimsvc/v0/RMprAdminServerSetInfo" }

func (o *xxx_ServerSetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ServerSetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerSetInfoRequest structure represents the RMprAdminServerSetInfo operation request
type ServerSetInfoRequest struct {
	// dwLevel: This is of type DWORD and MUST be set to any of the following values.
	//
	//	+-------+--------------------------------------------------------------------------------+
	//	|       |                                                                                |
	//	| VALUE |                                    MEANING                                     |
	//	|       |                                                                                |
	//	+-------+--------------------------------------------------------------------------------+
	//	+-------+--------------------------------------------------------------------------------+
	//	|     1 | pInfoStruct->pBuffer MUST be a pointer to a MPR_SERVER_1 (section 2.2.1.2.62). |
	//	+-------+--------------------------------------------------------------------------------+
	//	|     2 | pInfoStruct->pBuffer MUST be a pointer to a MPR_SERVER_2 (section 2.2.1.2.63). |
	//	+-------+--------------------------------------------------------------------------------+
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// pInfoStruct: This is a pointer of type PDIM_INFORMATION_CONTAINER (section 2.2.1.2.1)
	// and MUST be set to the following, based on dwLevel.
	//
	//	+---------+--------------------------------------------------------------------+-----------------------------------------------------------------------+
	//	|         |                                                                    |                                                                       |
	//	| DWLEVEL |                         PINFOSTRUCT->SIZE                          |                         PINFOSTRUCT->PBUFFER                          |
	//	|         |                                                                    |                                                                       |
	//	+---------+--------------------------------------------------------------------+-----------------------------------------------------------------------+
	//	+---------+--------------------------------------------------------------------+-----------------------------------------------------------------------+
	//	|       1 | This MUST be set to the size of MPR_SERVER_1 (section 2.2.1.2.62). | This MUST be set to a pointer to a MPR_SERVER_1 (section 2.2.1.2.62). |
	//	+---------+--------------------------------------------------------------------+-----------------------------------------------------------------------+
	//	|       2 | This MUST be set to the size of MPR_SERVER_2 (section 2.2.1.2.63). | This MUST be set to a pointer to a MPR_SERVER_2 (section 2.2.1.2.63). |
	//	+---------+--------------------------------------------------------------------+-----------------------------------------------------------------------+
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *ServerSetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerSetInfoOperation) *xxx_ServerSetInfoOperation {
	if op == nil {
		op = &xxx_ServerSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	return op
}

func (o *ServerSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerSetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
}
func (o *ServerSetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerSetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerSetInfoResponse structure represents the RMprAdminServerSetInfo operation response
type ServerSetInfoResponse struct {
	// Return: The RMprAdminServerSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerSetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerSetInfoOperation) *xxx_ServerSetInfoOperation {
	if op == nil {
		op = &xxx_ServerSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ServerSetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerSetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ServerSetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerSetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerGetInfoExOperation structure represents the RMprAdminServerGetInfoEx operation
type xxx_ServerGetInfoExOperation struct {
	ServerConfig *rrasm.ServerExIDL `idl:"name:pServerConfig" json:"server_config"`
	Return       uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerGetInfoExOperation) OpNum() int { return 44 }

func (o *xxx_ServerGetInfoExOperation) OpName() string { return "/dimsvc/v0/RMprAdminServerGetInfoEx" }

func (o *xxx_ServerGetInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pServerConfig {in, out} (1:{alias=PMPR_SERVER_EX_IDL}*(1))(2:{alias=MPR_SERVER_EX_IDL}(union))
	{
		if o.ServerConfig != nil {
			if err := o.ServerConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.ServerExIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pServerConfig {in, out} (1:{alias=PMPR_SERVER_EX_IDL,pointer=ref}*(1))(2:{alias=MPR_SERVER_EX_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.ServerExIDL{}
		}
		if err := o.ServerConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pServerConfig {in, out} (1:{alias=PMPR_SERVER_EX_IDL}*(1))(2:{alias=MPR_SERVER_EX_IDL}(union))
	{
		if o.ServerConfig != nil {
			if err := o.ServerConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.ServerExIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pServerConfig {in, out} (1:{alias=PMPR_SERVER_EX_IDL,pointer=ref}*(1))(2:{alias=MPR_SERVER_EX_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.ServerExIDL{}
		}
		if err := o.ServerConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerGetInfoExRequest structure represents the RMprAdminServerGetInfoEx operation request
type ServerGetInfoExRequest struct {
	// pServerConfig: A valid pointer to an MPR_SERVER_EX_IDL structure (section 2.2.1.2.143).
	// This MUST NOT be NULL. On successful return this parameter contains port information
	// for RRAS.
	ServerConfig *rrasm.ServerExIDL `idl:"name:pServerConfig" json:"server_config"`
}

func (o *ServerGetInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerGetInfoExOperation) *xxx_ServerGetInfoExOperation {
	if op == nil {
		op = &xxx_ServerGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerConfig = o.ServerConfig
	return op
}

func (o *ServerGetInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerGetInfoExOperation) {
	if o == nil {
		return
	}
	o.ServerConfig = op.ServerConfig
}
func (o *ServerGetInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerGetInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerGetInfoExResponse structure represents the RMprAdminServerGetInfoEx operation response
type ServerGetInfoExResponse struct {
	// pServerConfig: A valid pointer to an MPR_SERVER_EX_IDL structure (section 2.2.1.2.143).
	// This MUST NOT be NULL. On successful return this parameter contains port information
	// for RRAS.
	ServerConfig *rrasm.ServerExIDL `idl:"name:pServerConfig" json:"server_config"`
	// Return: The RMprAdminServerGetInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerGetInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerGetInfoExOperation) *xxx_ServerGetInfoExOperation {
	if op == nil {
		op = &xxx_ServerGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerConfig = o.ServerConfig
	op.Return = o.Return
	return op
}

func (o *ServerGetInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerGetInfoExOperation) {
	if o == nil {
		return
	}
	o.ServerConfig = op.ServerConfig
	o.Return = op.Return
}
func (o *ServerGetInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerGetInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionEnumExOperation structure represents the RRasAdminConnectionEnumEx operation
type xxx_ConnectionEnumExOperation struct {
	ObjectHeader        *rrasm.ObjectHeaderIDL      `idl:"name:objectHeader" json:"object_header"`
	PreferredMaxLength  uint32                      `idl:"name:dwPreferedMaxLen" json:"preferred_max_length"`
	EntriesRead         uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	LpdNumTotalElements uint32                      `idl:"name:lpdNumTotalElements" json:"lpd_num_total_elements"`
	RASConections       []*rrasm.RASConnectionExIDL `idl:"name:pRasConections;size_is:(, lpdwEntriesRead)" json:"ras_conections"`
	Resume              uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return              uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionEnumExOperation) OpNum() int { return 45 }

func (o *xxx_ConnectionEnumExOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionEnumEx"
}

func (o *xxx_ConnectionEnumExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// objectHeader {in} (1:{alias=PMPRAPI_OBJECT_HEADER_IDL}*(1))(2:{alias=MPRAPI_OBJECT_HEADER_IDL}(struct))
	{
		if o.ObjectHeader != nil {
			if err := o.ObjectHeader.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwPreferedMaxLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaxLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// objectHeader {in} (1:{alias=PMPRAPI_OBJECT_HEADER_IDL,pointer=ref}*(1))(2:{alias=MPRAPI_OBJECT_HEADER_IDL}(struct))
	{
		if o.ObjectHeader == nil {
			o.ObjectHeader = &rrasm.ObjectHeaderIDL{}
		}
		if err := o.ObjectHeader.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwPreferedMaxLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaxLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.RASConections != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.RASConections))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdNumTotalElements {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LpdNumTotalElements); err != nil {
			return err
		}
	}
	// pRasConections {out} (1:{pointer=ref}*(2))(2:{alias=PRAS_CONNECTION_EX_IDL}*(1))(3:{alias=RAS_CONNECTION_EX_IDL}[dim:0,size_is=lpdwEntriesRead](union))
	{
		if o.RASConections != nil || o.EntriesRead > 0 {
			_ptr_pRasConections := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.EntriesRead)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.RASConections {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.RASConections[i1] != nil {
						if err := o.RASConections[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&rrasm.RASConnectionExIDL{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.RASConections); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&rrasm.RASConnectionExIDL{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RASConections, _ptr_pRasConections); err != nil {
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
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpdwEntriesRead {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdNumTotalElements {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LpdNumTotalElements); err != nil {
			return err
		}
	}
	// pRasConections {out} (1:{pointer=ref}*(2))(2:{alias=PRAS_CONNECTION_EX_IDL,pointer=ref}*(1))(3:{alias=RAS_CONNECTION_EX_IDL}[dim:0,size_is=lpdwEntriesRead](union))
	{
		_ptr_pRasConections := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.RASConections", sizeInfo[0])
			}
			o.RASConections = make([]*rrasm.RASConnectionExIDL, sizeInfo[0])
			for i1 := range o.RASConections {
				i1 := i1
				if o.RASConections[i1] == nil {
					o.RASConections[i1] = &rrasm.RASConnectionExIDL{}
				}
				if err := o.RASConections[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pRasConections := func(ptr interface{}) { o.RASConections = *ptr.(*[]*rrasm.RASConnectionExIDL) }
		if err := w.ReadPointer(&o.RASConections, _s_pRasConections, _ptr_pRasConections); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionEnumExRequest structure represents the RRasAdminConnectionEnumEx operation request
type ConnectionEnumExRequest struct {
	// objectHeader: The pointer to an MPRAPI_OBJECT_HEADER_IDL structure (section 2.2.1.2.129).
	// In the structure, the revision field MUST be MPRAPI_RAS_CONNECTION_OBJECT_REVISION_1,
	// type filed MUST be MPRAPI_OBJECT_TYPE_RAS_CONNECTION_OBJECT, and size MUST be size
	// of RAS_CONNECTION_EX_IDL.
	ObjectHeader *rrasm.ObjectHeaderIDL `idl:"name:objectHeader" json:"object_header"`
	// dwPreferedMaxLen: This parameter is of type DWORD and SHOULD specify the preferred
	// maximum length of returned data (pRasConections) in bytes. If dwPreferedMaxLen is
	// -1, the buffer returned is large enough to hold all available information. The number
	// of entries returned is zero (0) in the case that dwPreferedMaxLen is less than the
	// size of one item; otherwise, the number of entries returned is one more than what
	// could be accommodated in dwPreferedMaxLen bytes.
	PreferredMaxLength uint32 `idl:"name:dwPreferedMaxLen" json:"preferred_max_length"`
	// lpdwResumeHandle: This is a pointer to type DWORD and specifies a resume handle used
	// to continue the enumeration. The lpdwResumeHandle parameter is NULL on the first
	// call and left unchanged on subsequent calls (caller MUST pass the same returned value
	// in the next call to this function). If the return code is ERROR_MORE_DATA, another
	// call MAY be made using this handle to retrieve more data. If the handle is NULL upon
	// return, the enumeration is complete. This handle is invalid for other types of error
	// returns.
	Resume uint32 `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *ConnectionEnumExRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionEnumExOperation) *xxx_ConnectionEnumExOperation {
	if op == nil {
		op = &xxx_ConnectionEnumExOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectHeader = o.ObjectHeader
	op.PreferredMaxLength = o.PreferredMaxLength
	op.Resume = o.Resume
	return op
}

func (o *ConnectionEnumExRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionEnumExOperation) {
	if o == nil {
		return
	}
	o.ObjectHeader = op.ObjectHeader
	o.PreferredMaxLength = op.PreferredMaxLength
	o.Resume = op.Resume
}
func (o *ConnectionEnumExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionEnumExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionEnumExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionEnumExResponse structure represents the RRasAdminConnectionEnumEx operation response
type ConnectionEnumExResponse struct {
	// lpdwEntriesRead: This is a pointer to type DWORD. Upon a successful function call
	// return, this parameter determines the total number of connections enumerated from
	// the current resume position given by lpdwResumeHandle.
	EntriesRead uint32 `idl:"name:lpdwEntriesRead" json:"entries_read"`
	// lpdNumTotalElements: This is a pointer to type DWORD and receives the total number
	// of connections that could have been enumerated from the current resume position given
	// by lpdwResumeHandle.
	LpdNumTotalElements uint32 `idl:"name:lpdNumTotalElements" json:"lpd_num_total_elements"`
	// pRasConections: Upon successful return, this is a pointer array of  RAS_CONNECTION_EX_IDL
	// (section 2.2.1.2.134) structures and the array size is determined by value pointed
	// to by lpdwEntriesRead.
	RASConections []*rrasm.RASConnectionExIDL `idl:"name:pRasConections;size_is:(, lpdwEntriesRead)" json:"ras_conections"`
	// lpdwResumeHandle: This is a pointer to type DWORD and specifies a resume handle used
	// to continue the enumeration. The lpdwResumeHandle parameter is NULL on the first
	// call and left unchanged on subsequent calls (caller MUST pass the same returned value
	// in the next call to this function). If the return code is ERROR_MORE_DATA, another
	// call MAY be made using this handle to retrieve more data. If the handle is NULL upon
	// return, the enumeration is complete. This handle is invalid for other types of error
	// returns.
	Resume uint32 `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRasAdminConnectionEnumEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionEnumExResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionEnumExOperation) *xxx_ConnectionEnumExOperation {
	if op == nil {
		op = &xxx_ConnectionEnumExOperation{}
	}
	if o == nil {
		return op
	}
	op.EntriesRead = o.EntriesRead
	op.LpdNumTotalElements = o.LpdNumTotalElements
	op.RASConections = o.RASConections
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *ConnectionEnumExResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionEnumExOperation) {
	if o == nil {
		return
	}
	o.EntriesRead = op.EntriesRead
	o.LpdNumTotalElements = op.LpdNumTotalElements
	o.RASConections = op.RASConections
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *ConnectionEnumExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionEnumExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionEnumExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionGetInfoExOperation structure represents the RRasAdminConnectionGetInfoEx operation
type xxx_ConnectionGetInfoExOperation struct {
	Connection    uint32                    `idl:"name:hDimConnection" json:"connection"`
	ObjectHeader  *rrasm.ObjectHeaderIDL    `idl:"name:objectHeader" json:"object_header"`
	RASConnection *rrasm.RASConnectionExIDL `idl:"name:pRasConnection" json:"ras_connection"`
	Return        uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionGetInfoExOperation) OpNum() int { return 46 }

func (o *xxx_ConnectionGetInfoExOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionGetInfoEx"
}

func (o *xxx_ConnectionGetInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	// objectHeader {in} (1:{alias=PMPRAPI_OBJECT_HEADER_IDL}*(1))(2:{alias=MPRAPI_OBJECT_HEADER_IDL}(struct))
	{
		if o.ObjectHeader != nil {
			if err := o.ObjectHeader.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	// objectHeader {in} (1:{alias=PMPRAPI_OBJECT_HEADER_IDL,pointer=ref}*(1))(2:{alias=MPRAPI_OBJECT_HEADER_IDL}(struct))
	{
		if o.ObjectHeader == nil {
			o.ObjectHeader = &rrasm.ObjectHeaderIDL{}
		}
		if err := o.ObjectHeader.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pRasConnection {out} (1:{alias=PRAS_CONNECTION_EX_IDL}*(1))(2:{alias=RAS_CONNECTION_EX_IDL}(union))
	{
		if o.RASConnection != nil {
			if err := o.RASConnection.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.RASConnectionExIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pRasConnection {out} (1:{alias=PRAS_CONNECTION_EX_IDL,pointer=ref}*(1))(2:{alias=RAS_CONNECTION_EX_IDL}(union))
	{
		if o.RASConnection == nil {
			o.RASConnection = &rrasm.RASConnectionExIDL{}
		}
		if err := o.RASConnection.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionGetInfoExRequest structure represents the RRasAdminConnectionGetInfoEx operation request
type ConnectionGetInfoExRequest struct {
	// hDimConnection: This is of type DWORD and MUST be set to a particular connection
	// identifier for which the connection information is required. Obtain this handle by
	// calling RRasAdminConnectionEnumEx (section 3.1.4.46). Because the RRASM server does
	// not maintain the connection handles, the RRAS server SHOULD check and ensure that
	// this handle represents a valid connection handle.
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
	// objectHeader: A pointer to MPRAPI_OBJECT_HEADER_IDL (section 2.2.1.2.129). In the
	// structure, the revision field MUST be MPRAPI_RAS_CONNECTION_OBJECT_REVISION_1, the
	// type field MUST be MPRAPI_OBJECT_TYPE_RAS_CONNECTION_OBJECT, and the size MUST be
	// the size of RAS_CONNECTION_EX_IDL.
	ObjectHeader *rrasm.ObjectHeaderIDL `idl:"name:objectHeader" json:"object_header"`
}

func (o *ConnectionGetInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionGetInfoExOperation) *xxx_ConnectionGetInfoExOperation {
	if op == nil {
		op = &xxx_ConnectionGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.Connection = o.Connection
	op.ObjectHeader = o.ObjectHeader
	return op
}

func (o *ConnectionGetInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionGetInfoExOperation) {
	if o == nil {
		return
	}
	o.Connection = op.Connection
	o.ObjectHeader = op.ObjectHeader
}
func (o *ConnectionGetInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionGetInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionGetInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionGetInfoExResponse structure represents the RRasAdminConnectionGetInfoEx operation response
type ConnectionGetInfoExResponse struct {
	// pRasConnection: Upon successful return, this is a pointer to a RAS_CONNECTION_EX_IDL
	// (section 2.2.1.2.134) structure.
	RASConnection *rrasm.RASConnectionExIDL `idl:"name:pRasConnection" json:"ras_connection"`
	// Return: The RRasAdminConnectionGetInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionGetInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionGetInfoExOperation) *xxx_ConnectionGetInfoExOperation {
	if op == nil {
		op = &xxx_ConnectionGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.RASConnection = o.RASConnection
	op.Return = o.Return
	return op
}

func (o *ConnectionGetInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionGetInfoExOperation) {
	if o == nil {
		return
	}
	o.RASConnection = op.RASConnection
	o.Return = op.Return
}
func (o *ConnectionGetInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionGetInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionGetInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerSetInfoExOperation structure represents the RMprAdminServerSetInfoEx operation
type xxx_ServerSetInfoExOperation struct {
	ServerConfig *rrasm.ServerSetConfigExIDL `idl:"name:pServerConfig" json:"server_config"`
	Return       uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerSetInfoExOperation) OpNum() int { return 47 }

func (o *xxx_ServerSetInfoExOperation) OpName() string { return "/dimsvc/v0/RMprAdminServerSetInfoEx" }

func (o *xxx_ServerSetInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pServerConfig {in} (1:{alias=PMPR_SERVER_SET_CONFIG_EX_IDL}*(1))(2:{alias=MPR_SERVER_SET_CONFIG_EX_IDL}(union))
	{
		if o.ServerConfig != nil {
			if err := o.ServerConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.ServerSetConfigExIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pServerConfig {in} (1:{alias=PMPR_SERVER_SET_CONFIG_EX_IDL,pointer=ref}*(1))(2:{alias=MPR_SERVER_SET_CONFIG_EX_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.ServerSetConfigExIDL{}
		}
		if err := o.ServerConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ServerSetInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerSetInfoExRequest structure represents the RMprAdminServerSetInfoEx operation request
type ServerSetInfoExRequest struct {
	// pServerConfig: This is a pointer to a structure PMPR_SERVER_SET_CONFIG_EX_IDL (section
	// 2.2.1.2.146) which contains the information required to set values on the RAS server.
	ServerConfig *rrasm.ServerSetConfigExIDL `idl:"name:pServerConfig" json:"server_config"`
}

func (o *ServerSetInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerSetInfoExOperation) *xxx_ServerSetInfoExOperation {
	if op == nil {
		op = &xxx_ServerSetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerConfig = o.ServerConfig
	return op
}

func (o *ServerSetInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerSetInfoExOperation) {
	if o == nil {
		return
	}
	o.ServerConfig = op.ServerConfig
}
func (o *ServerSetInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerSetInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerSetInfoExResponse structure represents the RMprAdminServerSetInfoEx operation response
type ServerSetInfoExResponse struct {
	// Return: The RMprAdminServerSetInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerSetInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerSetInfoExOperation) *xxx_ServerSetInfoExOperation {
	if op == nil {
		op = &xxx_ServerSetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ServerSetInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerSetInfoExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ServerSetInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerSetInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UpdateConnectionOperation structure represents the RRasAdminUpdateConnection operation
type xxx_UpdateConnectionOperation struct {
	Connection   uint32                        `idl:"name:hDimConnection" json:"connection"`
	ServerConfig *rrasm.RASUpdateConnectionIDL `idl:"name:pServerConfig" json:"server_config"`
	Return       uint32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_UpdateConnectionOperation) OpNum() int { return 48 }

func (o *xxx_UpdateConnectionOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminUpdateConnection"
}

func (o *xxx_UpdateConnectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateConnectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	// pServerConfig {in} (1:{alias=PRAS_UPDATE_CONNECTION_IDL}*(1))(2:{alias=RAS_UPDATE_CONNECTION_IDL}(union))
	{
		if o.ServerConfig != nil {
			if err := o.ServerConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.RASUpdateConnectionIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_UpdateConnectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	// pServerConfig {in} (1:{alias=PRAS_UPDATE_CONNECTION_IDL,pointer=ref}*(1))(2:{alias=RAS_UPDATE_CONNECTION_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.RASUpdateConnectionIDL{}
		}
		if err := o.ServerConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateConnectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateConnectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_UpdateConnectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UpdateConnectionRequest structure represents the RRasAdminUpdateConnection operation request
type UpdateConnectionRequest struct {
	// hDimConnection: This is of type DWORD and SHOULD be set to a particular IKEv2 connection
	// identifier for which the endpoint needs to be updated. Obtain this handle by calling
	// RRasAdminConnectionEnum (see section 3.1.4.2). Since the RRASM server does not maintain
	// connection handles, the RRAS server SHOULD validate this handle.
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
	// pServerConfig: This is a pointer to a RAS_UPDATE_CONNECTION_IDL structure (section
	// 2.2.1.2.148) that contains the information required to set values on the RRAS server.
	ServerConfig *rrasm.RASUpdateConnectionIDL `idl:"name:pServerConfig" json:"server_config"`
}

func (o *UpdateConnectionRequest) xxx_ToOp(ctx context.Context, op *xxx_UpdateConnectionOperation) *xxx_UpdateConnectionOperation {
	if op == nil {
		op = &xxx_UpdateConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.Connection = o.Connection
	op.ServerConfig = o.ServerConfig
	return op
}

func (o *UpdateConnectionRequest) xxx_FromOp(ctx context.Context, op *xxx_UpdateConnectionOperation) {
	if o == nil {
		return
	}
	o.Connection = op.Connection
	o.ServerConfig = op.ServerConfig
}
func (o *UpdateConnectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UpdateConnectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateConnectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UpdateConnectionResponse structure represents the RRasAdminUpdateConnection operation response
type UpdateConnectionResponse struct {
	// Return: The RRasAdminUpdateConnection return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UpdateConnectionResponse) xxx_ToOp(ctx context.Context, op *xxx_UpdateConnectionOperation) *xxx_UpdateConnectionOperation {
	if op == nil {
		op = &xxx_UpdateConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *UpdateConnectionResponse) xxx_FromOp(ctx context.Context, op *xxx_UpdateConnectionOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UpdateConnectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UpdateConnectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateConnectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceSetCredentialsLocalOperation structure represents the RRouterInterfaceSetCredentialsLocal operation
type xxx_RouterInterfaceSetCredentialsLocalOperation struct {
	InterfaceName string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
	UserName      string `idl:"name:lpwsUserName;string" json:"user_name"`
	DomainName    string `idl:"name:lpwsDomainName;string" json:"domain_name"`
	Password      string `idl:"name:lpwsPassword;string" json:"password"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) OpNum() int { return 49 }

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceSetCredentialsLocal"
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.InterfaceName); err != nil {
			return err
		}
	}
	// lpwsUserName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.UserName); err != nil {
			return err
		}
	}
	// lpwsDomainName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DomainName); err != nil {
			return err
		}
	}
	// lpwsPassword {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.InterfaceName); err != nil {
			return err
		}
	}
	// lpwsUserName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.UserName); err != nil {
			return err
		}
	}
	// lpwsDomainName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DomainName); err != nil {
			return err
		}
	}
	// lpwsPassword {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceSetCredentialsLocalRequest structure represents the RRouterInterfaceSetCredentialsLocal operation request
type RouterInterfaceSetCredentialsLocalRequest struct {
	// lpwsInterfaceName: A pointer to a null-terminated Unicode string that contains the
	// name of the interface for which credentials need to be set. The length of the string
	// MUST NOT be more than 256 characters, otherwise an error is returned.
	InterfaceName string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
	// lpwsUserName: A pointer to a null-terminated Unicode string that contains the name
	// of the user on this connection. The string length MUST NOT be more than 256 characters,
	// otherwise an error is returned.
	UserName string `idl:"name:lpwsUserName;string" json:"user_name"`
	// lpwsDomainName: A pointer to a null-terminated Unicode string that contains the domain
	// name. The string length MUST NOT be more than 16 characters, otherwise an error is
	// returned.
	DomainName string `idl:"name:lpwsDomainName;string" json:"domain_name"`
	// lpwsPassword: A pointer to a null-terminated Unicode string that contains the password
	// of the user on this connection. The string length MUST NOT be more than 256 characters,
	// otherwise an error is returned.
	Password string `idl:"name:lpwsPassword;string" json:"password"`
}

func (o *RouterInterfaceSetCredentialsLocalRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsLocalOperation) *xxx_RouterInterfaceSetCredentialsLocalOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCredentialsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceName = o.InterfaceName
	op.UserName = o.UserName
	op.DomainName = o.DomainName
	op.Password = o.Password
	return op
}

func (o *RouterInterfaceSetCredentialsLocalRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.InterfaceName = op.InterfaceName
	o.UserName = op.UserName
	o.DomainName = op.DomainName
	o.Password = op.Password
}
func (o *RouterInterfaceSetCredentialsLocalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceSetCredentialsLocalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCredentialsLocalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceSetCredentialsLocalResponse structure represents the RRouterInterfaceSetCredentialsLocal operation response
type RouterInterfaceSetCredentialsLocalResponse struct {
	// Return: The RRouterInterfaceSetCredentialsLocal return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceSetCredentialsLocalResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsLocalOperation) *xxx_RouterInterfaceSetCredentialsLocalOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCredentialsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceSetCredentialsLocalResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceSetCredentialsLocalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceSetCredentialsLocalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCredentialsLocalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceGetCredentialsLocalOperation structure represents the RRouterInterfaceGetCredentialsLocal operation
type xxx_RouterInterfaceGetCredentialsLocalOperation struct {
	InterfaceName string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
	UserName      string `idl:"name:lpwsUserName;string" json:"user_name"`
	DomainName    string `idl:"name:lpwsDomainName;string" json:"domain_name"`
	Password      string `idl:"name:lpwsPassword;string" json:"password"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) OpNum() int { return 50 }

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceGetCredentialsLocal"
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.InterfaceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.InterfaceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpwsUserName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.UserName != "" {
			_ptr_lpwsUserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.UserName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.UserName, _ptr_lpwsUserName); err != nil {
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
	// lpwsDomainName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DomainName != "" {
			_ptr_lpwsDomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DomainName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainName, _ptr_lpwsDomainName); err != nil {
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
	// lpwsPassword {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Password != "" {
			_ptr_lpwsPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_lpwsPassword); err != nil {
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
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpwsUserName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpwsUserName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.UserName); err != nil {
				return err
			}
			return nil
		})
		_s_lpwsUserName := func(ptr interface{}) { o.UserName = *ptr.(*string) }
		if err := w.ReadPointer(&o.UserName, _s_lpwsUserName, _ptr_lpwsUserName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpwsDomainName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpwsDomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DomainName); err != nil {
				return err
			}
			return nil
		})
		_s_lpwsDomainName := func(ptr interface{}) { o.DomainName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DomainName, _s_lpwsDomainName, _ptr_lpwsDomainName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpwsPassword {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpwsPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
				return err
			}
			return nil
		})
		_s_lpwsPassword := func(ptr interface{}) { o.Password = *ptr.(*string) }
		if err := w.ReadPointer(&o.Password, _s_lpwsPassword, _ptr_lpwsPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceGetCredentialsLocalRequest structure represents the RRouterInterfaceGetCredentialsLocal operation request
type RouterInterfaceGetCredentialsLocalRequest struct {
	// lpwsInterfaceName: A pointer to a null-terminated Unicode string that contains the
	// name of the interface for which credentials need to be set. The length of the string
	// MUST NOT be more than 256 characters, otherwise an error is returned. The client
	// SHOULD free this memory.
	InterfaceName string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
}

func (o *RouterInterfaceGetCredentialsLocalRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) *xxx_RouterInterfaceGetCredentialsLocalOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCredentialsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceName = o.InterfaceName
	return op
}

func (o *RouterInterfaceGetCredentialsLocalRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.InterfaceName = op.InterfaceName
}
func (o *RouterInterfaceGetCredentialsLocalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceGetCredentialsLocalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCredentialsLocalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceGetCredentialsLocalResponse structure represents the RRouterInterfaceGetCredentialsLocal operation response
type RouterInterfaceGetCredentialsLocalResponse struct {
	// lpwsUserName: A pointer to a null-terminated Unicode string that contains the name
	// of the user on this connection. The string length MUST NOT be more than 256 characters.
	// The client SHOULD free the memory pointed to by lpwsUserName.
	UserName string `idl:"name:lpwsUserName;string" json:"user_name"`
	// lpwsDomainName: A pointer to a null-terminated Unicode string that contains the domain
	// name. The string length MUST NOT be more than 16 characters. The client SHOULD free
	// the memory pointed to by lpwsDomainName.
	DomainName string `idl:"name:lpwsDomainName;string" json:"domain_name"`
	// lpwsPassword: A pointer to a null-terminated Unicode string that contains the password
	// of the user on this connection. The string length MUST NOT be more than 256 characters.
	// The client SHOULD free the memory pointed to by lpwsPassword.
	Password string `idl:"name:lpwsPassword;string" json:"password"`
	// Return: The RRouterInterfaceGetCredentialsLocal return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceGetCredentialsLocalResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) *xxx_RouterInterfaceGetCredentialsLocalOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCredentialsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.UserName = o.UserName
	op.DomainName = o.DomainName
	op.Password = o.Password
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetCredentialsLocalResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.UserName = op.UserName
	o.DomainName = op.DomainName
	o.Password = op.Password
	o.Return = op.Return
}
func (o *RouterInterfaceGetCredentialsLocalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceGetCredentialsLocalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCredentialsLocalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceGetCustomInfoExOperation structure represents the RRouterInterfaceGetCustomInfoEx operation
type xxx_RouterInterfaceGetCustomInfoExOperation struct {
	Interface             uint32                          `idl:"name:hInterface" json:"interface"`
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
	Return                uint32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) OpNum() int { return 51 }

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceGetCustomInfoEx"
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig != nil {
			if err := o.InterfaceCustomConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL,pointer=ref}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig == nil {
			o.InterfaceCustomConfig = &rrasm.InterfaceCustominfoexIDL{}
		}
		if err := o.InterfaceCustomConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig != nil {
			if err := o.InterfaceCustomConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL,pointer=ref}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig == nil {
			o.InterfaceCustomConfig = &rrasm.InterfaceCustominfoexIDL{}
		}
		if err := o.InterfaceCustomConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceGetCustomInfoExRequest structure represents the RRouterInterfaceGetCustomInfoEx operation request
type RouterInterfaceGetCustomInfoExRequest struct {
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Because the RRASM
	// server does not maintain the interface handles, the RRAS server SHOULD check to ensure
	// that this handle represents a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	// pIfCustomConfig: A valid pointer to an MPR_IF_CUSTOMINFOEX_IDL (section 2.2.1.2.245)
	// structure. This MUST NOT be NULL. On successful return, this parameter contains tunnel-specific
	// custom configuration for the interface whose handle is specified in the hInterface
	// parameter.
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
}

func (o *RouterInterfaceGetCustomInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCustomInfoExOperation) *xxx_RouterInterfaceGetCustomInfoExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCustomInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.InterfaceCustomConfig = o.InterfaceCustomConfig
	return op
}

func (o *RouterInterfaceGetCustomInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCustomInfoExOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.InterfaceCustomConfig = op.InterfaceCustomConfig
}
func (o *RouterInterfaceGetCustomInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceGetCustomInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCustomInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceGetCustomInfoExResponse structure represents the RRouterInterfaceGetCustomInfoEx operation response
type RouterInterfaceGetCustomInfoExResponse struct {
	// pIfCustomConfig: A valid pointer to an MPR_IF_CUSTOMINFOEX_IDL (section 2.2.1.2.245)
	// structure. This MUST NOT be NULL. On successful return, this parameter contains tunnel-specific
	// custom configuration for the interface whose handle is specified in the hInterface
	// parameter.
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
	// Return: The RRouterInterfaceGetCustomInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceGetCustomInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCustomInfoExOperation) *xxx_RouterInterfaceGetCustomInfoExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCustomInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceCustomConfig = o.InterfaceCustomConfig
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetCustomInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCustomInfoExOperation) {
	if o == nil {
		return
	}
	o.InterfaceCustomConfig = op.InterfaceCustomConfig
	o.Return = op.Return
}
func (o *RouterInterfaceGetCustomInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceGetCustomInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCustomInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceSetCustomInfoExOperation structure represents the RRouterInterfaceSetCustomInfoEx operation
type xxx_RouterInterfaceSetCustomInfoExOperation struct {
	Interface             uint32                          `idl:"name:hInterface" json:"interface"`
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
	Return                uint32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) OpNum() int { return 52 }

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceSetCustomInfoEx"
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig != nil {
			if err := o.InterfaceCustomConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL,pointer=ref}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig == nil {
			o.InterfaceCustomConfig = &rrasm.InterfaceCustominfoexIDL{}
		}
		if err := o.InterfaceCustomConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig != nil {
			if err := o.InterfaceCustomConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL,pointer=ref}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig == nil {
			o.InterfaceCustomConfig = &rrasm.InterfaceCustominfoexIDL{}
		}
		if err := o.InterfaceCustomConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceSetCustomInfoExRequest structure represents the RRouterInterfaceSetCustomInfoEx operation request
type RouterInterfaceSetCustomInfoExRequest struct {
	// hInterface: The unique identifier of an interface. This can be obtained from RRouterInterfaceCreate
	// (section 3.1.4.13) or RRouterInterfaceGetHandle (section 3.1.4.12). Because the RRASM
	// server does not maintain the interface handles, the RRAS server SHOULD check to ensure
	// that this handle represents a valid interface handle.
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	// pIfCustomConfig: A valid pointer to an MPR_IF_CUSTOMINFOEX_IDL (section 2.2.1.2.245)
	// structure that contains custom configurations to be set for the interface whose handle
	// is specified in the hInterface parameter. This MUST NOT be NULL.
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
}

func (o *RouterInterfaceSetCustomInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCustomInfoExOperation) *xxx_RouterInterfaceSetCustomInfoExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCustomInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.InterfaceCustomConfig = o.InterfaceCustomConfig
	return op
}

func (o *RouterInterfaceSetCustomInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCustomInfoExOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.InterfaceCustomConfig = op.InterfaceCustomConfig
}
func (o *RouterInterfaceSetCustomInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceSetCustomInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCustomInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceSetCustomInfoExResponse structure represents the RRouterInterfaceSetCustomInfoEx operation response
type RouterInterfaceSetCustomInfoExResponse struct {
	// pIfCustomConfig: A valid pointer to an MPR_IF_CUSTOMINFOEX_IDL (section 2.2.1.2.245)
	// structure that contains custom configurations to be set for the interface whose handle
	// is specified in the hInterface parameter. This MUST NOT be NULL.
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
	// Return: The RRouterInterfaceSetCustomInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceSetCustomInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCustomInfoExOperation) *xxx_RouterInterfaceSetCustomInfoExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCustomInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceCustomConfig = o.InterfaceCustomConfig
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceSetCustomInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCustomInfoExOperation) {
	if o == nil {
		return
	}
	o.InterfaceCustomConfig = op.InterfaceCustomConfig
	o.Return = op.Return
}
func (o *RouterInterfaceSetCustomInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceSetCustomInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCustomInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
