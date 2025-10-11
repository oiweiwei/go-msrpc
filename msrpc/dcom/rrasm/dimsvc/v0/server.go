package dimsvc

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

// dimsvc server interface.
type DimsvcServer interface {

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
	ServerGetInfo(context.Context, *ServerGetInfoRequest) (*ServerGetInfoResponse, error)

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
	ConnectionEnum(context.Context, *ConnectionEnumRequest) (*ConnectionEnumResponse, error)

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
	ConnectionGetInfo(context.Context, *ConnectionGetInfoRequest) (*ConnectionGetInfoResponse, error)

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
	ConnectionClearStats(context.Context, *ConnectionClearStatsRequest) (*ConnectionClearStatsResponse, error)

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
	PortEnum(context.Context, *PortEnumRequest) (*PortEnumResponse, error)

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
	PortGetInfo(context.Context, *PortGetInfoRequest) (*PortGetInfoResponse, error)

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
	PortClearStats(context.Context, *PortClearStatsRequest) (*PortClearStatsResponse, error)

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
	PortReset(context.Context, *PortResetRequest) (*PortResetResponse, error)

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
	PortDisconnect(context.Context, *PortDisconnectRequest) (*PortDisconnectResponse, error)

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
	RouterInterfaceTransportSetGlobalInfo(context.Context, *RouterInterfaceTransportSetGlobalInfoRequest) (*RouterInterfaceTransportSetGlobalInfoResponse, error)

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
	RouterInterfaceTransportGetGlobalInfo(context.Context, *RouterInterfaceTransportGetGlobalInfoRequest) (*RouterInterfaceTransportGetGlobalInfoResponse, error)

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
	RouterInterfaceGetHandle(context.Context, *RouterInterfaceGetHandleRequest) (*RouterInterfaceGetHandleResponse, error)

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
	RouterInterfaceCreate(context.Context, *RouterInterfaceCreateRequest) (*RouterInterfaceCreateResponse, error)

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
	RouterInterfaceGetInfo(context.Context, *RouterInterfaceGetInfoRequest) (*RouterInterfaceGetInfoResponse, error)

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
	RouterInterfaceSetInfo(context.Context, *RouterInterfaceSetInfoRequest) (*RouterInterfaceSetInfoResponse, error)

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
	RouterInterfaceDelete(context.Context, *RouterInterfaceDeleteRequest) (*RouterInterfaceDeleteResponse, error)

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
	RouterInterfaceTransportRemove(context.Context, *RouterInterfaceTransportRemoveRequest) (*RouterInterfaceTransportRemoveResponse, error)

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
	RouterInterfaceTransportAdd(context.Context, *RouterInterfaceTransportAddRequest) (*RouterInterfaceTransportAddResponse, error)

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
	RouterInterfaceTransportGetInfo(context.Context, *RouterInterfaceTransportGetInfoRequest) (*RouterInterfaceTransportGetInfoResponse, error)

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
	RouterInterfaceTransportSetInfo(context.Context, *RouterInterfaceTransportSetInfoRequest) (*RouterInterfaceTransportSetInfoResponse, error)

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
	RouterInterfaceEnum(context.Context, *RouterInterfaceEnumRequest) (*RouterInterfaceEnumResponse, error)

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
	RouterInterfaceConnect(context.Context, *RouterInterfaceConnectRequest) (*RouterInterfaceConnectResponse, error)

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
	RouterInterfaceDisconnect(context.Context, *RouterInterfaceDisconnectRequest) (*RouterInterfaceDisconnectResponse, error)

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
	RouterInterfaceUpdateRoutes(context.Context, *RouterInterfaceUpdateRoutesRequest) (*RouterInterfaceUpdateRoutesResponse, error)

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
	RouterInterfaceQueryUpdateResult(context.Context, *RouterInterfaceQueryUpdateResultRequest) (*RouterInterfaceQueryUpdateResultResponse, error)

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
	RouterInterfaceUpdatePhonebookInfo(context.Context, *RouterInterfaceUpdatePhonebookInfoRequest) (*RouterInterfaceUpdatePhonebookInfoResponse, error)

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
	MIBEntryCreate(context.Context, *MIBEntryCreateRequest) (*MIBEntryCreateResponse, error)

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
	MIBEntryDelete(context.Context, *MIBEntryDeleteRequest) (*MIBEntryDeleteResponse, error)

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
	MIBEntrySet(context.Context, *MIBEntrySetRequest) (*MIBEntrySetResponse, error)

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
	MIBEntryGet(context.Context, *MIBEntryGetRequest) (*MIBEntryGetResponse, error)

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
	MIBEntryGetFirst(context.Context, *MIBEntryGetFirstRequest) (*MIBEntryGetFirstResponse, error)

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
	MIBEntryGetNext(context.Context, *MIBEntryGetNextRequest) (*MIBEntryGetNextResponse, error)

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
	MIBGetTrapInfo(context.Context, *MIBGetTrapInfoRequest) (*MIBGetTrapInfoResponse, error)

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
	MIBSetTrapInfo(context.Context, *MIBSetTrapInfoRequest) (*MIBSetTrapInfoResponse, error)

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
	ConnectionNotification(context.Context, *ConnectionNotificationRequest) (*ConnectionNotificationResponse, error)

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
	SendUserMessage(context.Context, *SendUserMessageRequest) (*SendUserMessageResponse, error)

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
	RouterDeviceEnum(context.Context, *RouterDeviceEnumRequest) (*RouterDeviceEnumResponse, error)

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
	RouterInterfaceTransportCreate(context.Context, *RouterInterfaceTransportCreateRequest) (*RouterInterfaceTransportCreateResponse, error)

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
	RouterInterfaceDeviceGetInfo(context.Context, *RouterInterfaceDeviceGetInfoRequest) (*RouterInterfaceDeviceGetInfoResponse, error)

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
	RouterInterfaceDeviceSetInfo(context.Context, *RouterInterfaceDeviceSetInfoRequest) (*RouterInterfaceDeviceSetInfoResponse, error)

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
	RouterInterfaceSetCredentialsEx(context.Context, *RouterInterfaceSetCredentialsExRequest) (*RouterInterfaceSetCredentialsExResponse, error)

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
	RouterInterfaceGetCredentialsEx(context.Context, *RouterInterfaceGetCredentialsExRequest) (*RouterInterfaceGetCredentialsExResponse, error)

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
	ConnectionRemoveQuarantine(context.Context, *ConnectionRemoveQuarantineRequest) (*ConnectionRemoveQuarantineResponse, error)

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
	ServerSetInfo(context.Context, *ServerSetInfoRequest) (*ServerSetInfoResponse, error)

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
	ServerGetInfoEx(context.Context, *ServerGetInfoExRequest) (*ServerGetInfoExResponse, error)

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
	ConnectionEnumEx(context.Context, *ConnectionEnumExRequest) (*ConnectionEnumExResponse, error)

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
	ConnectionGetInfoEx(context.Context, *ConnectionGetInfoExRequest) (*ConnectionGetInfoExResponse, error)

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
	ServerSetInfoEx(context.Context, *ServerSetInfoExRequest) (*ServerSetInfoExResponse, error)

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
	UpdateConnection(context.Context, *UpdateConnectionRequest) (*UpdateConnectionResponse, error)

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
	RouterInterfaceSetCredentialsLocal(context.Context, *RouterInterfaceSetCredentialsLocalRequest) (*RouterInterfaceSetCredentialsLocalResponse, error)

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
	RouterInterfaceGetCredentialsLocal(context.Context, *RouterInterfaceGetCredentialsLocalRequest) (*RouterInterfaceGetCredentialsLocalResponse, error)

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
	RouterInterfaceGetCustomInfoEx(context.Context, *RouterInterfaceGetCustomInfoExRequest) (*RouterInterfaceGetCustomInfoExResponse, error)

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
	RouterInterfaceSetCustomInfoEx(context.Context, *RouterInterfaceSetCustomInfoExRequest) (*RouterInterfaceSetCustomInfoExResponse, error)
}

func RegisterDimsvcServer(conn dcerpc.Conn, o DimsvcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDimsvcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DimsvcSyntaxV0_0))...)
}

func NewDimsvcServerHandle(o DimsvcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DimsvcServerHandle(ctx, o, opNum, r)
	}
}

func DimsvcServerHandle(ctx context.Context, o DimsvcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RMprAdminServerGetInfo
		op := &xxx_ServerGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RRasAdminConnectionEnum
		op := &xxx_ConnectionEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // RRasAdminConnectionGetInfo
		op := &xxx_ConnectionGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // RRasAdminConnectionClearStats
		op := &xxx_ConnectionClearStatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionClearStatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionClearStats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RRasAdminPortEnum
		op := &xxx_PortEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PortEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PortEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // RRasAdminPortGetInfo
		op := &xxx_PortGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PortGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PortGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // RRasAdminPortClearStats
		op := &xxx_PortClearStatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PortClearStatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PortClearStats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // RRasAdminPortReset
		op := &xxx_PortResetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PortResetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PortReset(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RRasAdminPortDisconnect
		op := &xxx_PortDisconnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PortDisconnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PortDisconnect(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RRouterInterfaceTransportSetGlobalInfo
		op := &xxx_RouterInterfaceTransportSetGlobalInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportSetGlobalInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportSetGlobalInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RRouterInterfaceTransportGetGlobalInfo
		op := &xxx_RouterInterfaceTransportGetGlobalInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportGetGlobalInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportGetGlobalInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RRouterInterfaceGetHandle
		op := &xxx_RouterInterfaceGetHandleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceGetHandleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceGetHandle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RRouterInterfaceCreate
		op := &xxx_RouterInterfaceCreateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceCreateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceCreate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // RRouterInterfaceGetInfo
		op := &xxx_RouterInterfaceGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // RRouterInterfaceSetInfo
		op := &xxx_RouterInterfaceSetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceSetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceSetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RRouterInterfaceDelete
		op := &xxx_RouterInterfaceDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // RRouterInterfaceTransportRemove
		op := &xxx_RouterInterfaceTransportRemoveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportRemoveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportRemove(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // RRouterInterfaceTransportAdd
		op := &xxx_RouterInterfaceTransportAddOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportAddRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportAdd(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // RRouterInterfaceTransportGetInfo
		op := &xxx_RouterInterfaceTransportGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // RRouterInterfaceTransportSetInfo
		op := &xxx_RouterInterfaceTransportSetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportSetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportSetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // RRouterInterfaceEnum
		op := &xxx_RouterInterfaceEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // RRouterInterfaceConnect
		op := &xxx_RouterInterfaceConnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceConnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceConnect(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // RRouterInterfaceDisconnect
		op := &xxx_RouterInterfaceDisconnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceDisconnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceDisconnect(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // RRouterInterfaceUpdateRoutes
		op := &xxx_RouterInterfaceUpdateRoutesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceUpdateRoutesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceUpdateRoutes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // RRouterInterfaceQueryUpdateResult
		op := &xxx_RouterInterfaceQueryUpdateResultOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceQueryUpdateResultRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceQueryUpdateResult(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // RRouterInterfaceUpdatePhonebookInfo
		op := &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceUpdatePhonebookInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceUpdatePhonebookInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // RMIBEntryCreate
		op := &xxx_MIBEntryCreateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntryCreateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntryCreate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // RMIBEntryDelete
		op := &xxx_MIBEntryDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntryDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntryDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // RMIBEntrySet
		op := &xxx_MIBEntrySetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntrySetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntrySet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // RMIBEntryGet
		op := &xxx_MIBEntryGetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntryGetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntryGet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // RMIBEntryGetFirst
		op := &xxx_MIBEntryGetFirstOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntryGetFirstRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntryGetFirst(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // RMIBEntryGetNext
		op := &xxx_MIBEntryGetNextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntryGetNextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntryGetNext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // RMIBGetTrapInfo
		op := &xxx_MIBGetTrapInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBGetTrapInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBGetTrapInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // RMIBSetTrapInfo
		op := &xxx_MIBSetTrapInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBSetTrapInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBSetTrapInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // RRasAdminConnectionNotification
		op := &xxx_ConnectionNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // RRasAdminSendUserMessage
		op := &xxx_SendUserMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendUserMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SendUserMessage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // RRouterDeviceEnum
		op := &xxx_RouterDeviceEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterDeviceEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterDeviceEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // RRouterInterfaceTransportCreate
		op := &xxx_RouterInterfaceTransportCreateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportCreateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportCreate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // RRouterInterfaceDeviceGetInfo
		op := &xxx_RouterInterfaceDeviceGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceDeviceGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceDeviceGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // RRouterInterfaceDeviceSetInfo
		op := &xxx_RouterInterfaceDeviceSetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceDeviceSetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceDeviceSetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // RRouterInterfaceSetCredentialsEx
		op := &xxx_RouterInterfaceSetCredentialsExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceSetCredentialsExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceSetCredentialsEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // RRouterInterfaceGetCredentialsEx
		op := &xxx_RouterInterfaceGetCredentialsExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceGetCredentialsExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceGetCredentialsEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // RRasAdminConnectionRemoveQuarantine
		op := &xxx_ConnectionRemoveQuarantineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionRemoveQuarantineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionRemoveQuarantine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // RMprAdminServerSetInfo
		op := &xxx_ServerSetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerSetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerSetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // RMprAdminServerGetInfoEx
		op := &xxx_ServerGetInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerGetInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerGetInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // RRasAdminConnectionEnumEx
		op := &xxx_ConnectionEnumExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionEnumExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionEnumEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // RRasAdminConnectionGetInfoEx
		op := &xxx_ConnectionGetInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionGetInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionGetInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // RMprAdminServerSetInfoEx
		op := &xxx_ServerSetInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerSetInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerSetInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // RRasAdminUpdateConnection
		op := &xxx_UpdateConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpdateConnection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // RRouterInterfaceSetCredentialsLocal
		op := &xxx_RouterInterfaceSetCredentialsLocalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceSetCredentialsLocalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceSetCredentialsLocal(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // RRouterInterfaceGetCredentialsLocal
		op := &xxx_RouterInterfaceGetCredentialsLocalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceGetCredentialsLocalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceGetCredentialsLocal(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // RRouterInterfaceGetCustomInfoEx
		op := &xxx_RouterInterfaceGetCustomInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceGetCustomInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceGetCustomInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // RRouterInterfaceSetCustomInfoEx
		op := &xxx_RouterInterfaceSetCustomInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceSetCustomInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceSetCustomInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented dimsvc
type UnimplementedDimsvcServer struct {
}

func (UnimplementedDimsvcServer) ServerGetInfo(context.Context, *ServerGetInfoRequest) (*ServerGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionEnum(context.Context, *ConnectionEnumRequest) (*ConnectionEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionGetInfo(context.Context, *ConnectionGetInfoRequest) (*ConnectionGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionClearStats(context.Context, *ConnectionClearStatsRequest) (*ConnectionClearStatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) PortEnum(context.Context, *PortEnumRequest) (*PortEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) PortGetInfo(context.Context, *PortGetInfoRequest) (*PortGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) PortClearStats(context.Context, *PortClearStatsRequest) (*PortClearStatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) PortReset(context.Context, *PortResetRequest) (*PortResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) PortDisconnect(context.Context, *PortDisconnectRequest) (*PortDisconnectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportSetGlobalInfo(context.Context, *RouterInterfaceTransportSetGlobalInfoRequest) (*RouterInterfaceTransportSetGlobalInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportGetGlobalInfo(context.Context, *RouterInterfaceTransportGetGlobalInfoRequest) (*RouterInterfaceTransportGetGlobalInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceGetHandle(context.Context, *RouterInterfaceGetHandleRequest) (*RouterInterfaceGetHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceCreate(context.Context, *RouterInterfaceCreateRequest) (*RouterInterfaceCreateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceGetInfo(context.Context, *RouterInterfaceGetInfoRequest) (*RouterInterfaceGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceSetInfo(context.Context, *RouterInterfaceSetInfoRequest) (*RouterInterfaceSetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceDelete(context.Context, *RouterInterfaceDeleteRequest) (*RouterInterfaceDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportRemove(context.Context, *RouterInterfaceTransportRemoveRequest) (*RouterInterfaceTransportRemoveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportAdd(context.Context, *RouterInterfaceTransportAddRequest) (*RouterInterfaceTransportAddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportGetInfo(context.Context, *RouterInterfaceTransportGetInfoRequest) (*RouterInterfaceTransportGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportSetInfo(context.Context, *RouterInterfaceTransportSetInfoRequest) (*RouterInterfaceTransportSetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceEnum(context.Context, *RouterInterfaceEnumRequest) (*RouterInterfaceEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceConnect(context.Context, *RouterInterfaceConnectRequest) (*RouterInterfaceConnectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceDisconnect(context.Context, *RouterInterfaceDisconnectRequest) (*RouterInterfaceDisconnectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceUpdateRoutes(context.Context, *RouterInterfaceUpdateRoutesRequest) (*RouterInterfaceUpdateRoutesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceQueryUpdateResult(context.Context, *RouterInterfaceQueryUpdateResultRequest) (*RouterInterfaceQueryUpdateResultResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceUpdatePhonebookInfo(context.Context, *RouterInterfaceUpdatePhonebookInfoRequest) (*RouterInterfaceUpdatePhonebookInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntryCreate(context.Context, *MIBEntryCreateRequest) (*MIBEntryCreateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntryDelete(context.Context, *MIBEntryDeleteRequest) (*MIBEntryDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntrySet(context.Context, *MIBEntrySetRequest) (*MIBEntrySetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntryGet(context.Context, *MIBEntryGetRequest) (*MIBEntryGetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntryGetFirst(context.Context, *MIBEntryGetFirstRequest) (*MIBEntryGetFirstResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntryGetNext(context.Context, *MIBEntryGetNextRequest) (*MIBEntryGetNextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBGetTrapInfo(context.Context, *MIBGetTrapInfoRequest) (*MIBGetTrapInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBSetTrapInfo(context.Context, *MIBSetTrapInfoRequest) (*MIBSetTrapInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionNotification(context.Context, *ConnectionNotificationRequest) (*ConnectionNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) SendUserMessage(context.Context, *SendUserMessageRequest) (*SendUserMessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterDeviceEnum(context.Context, *RouterDeviceEnumRequest) (*RouterDeviceEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportCreate(context.Context, *RouterInterfaceTransportCreateRequest) (*RouterInterfaceTransportCreateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceDeviceGetInfo(context.Context, *RouterInterfaceDeviceGetInfoRequest) (*RouterInterfaceDeviceGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceDeviceSetInfo(context.Context, *RouterInterfaceDeviceSetInfoRequest) (*RouterInterfaceDeviceSetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceSetCredentialsEx(context.Context, *RouterInterfaceSetCredentialsExRequest) (*RouterInterfaceSetCredentialsExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceGetCredentialsEx(context.Context, *RouterInterfaceGetCredentialsExRequest) (*RouterInterfaceGetCredentialsExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionRemoveQuarantine(context.Context, *ConnectionRemoveQuarantineRequest) (*ConnectionRemoveQuarantineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ServerSetInfo(context.Context, *ServerSetInfoRequest) (*ServerSetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ServerGetInfoEx(context.Context, *ServerGetInfoExRequest) (*ServerGetInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionEnumEx(context.Context, *ConnectionEnumExRequest) (*ConnectionEnumExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionGetInfoEx(context.Context, *ConnectionGetInfoExRequest) (*ConnectionGetInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ServerSetInfoEx(context.Context, *ServerSetInfoExRequest) (*ServerSetInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) UpdateConnection(context.Context, *UpdateConnectionRequest) (*UpdateConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceSetCredentialsLocal(context.Context, *RouterInterfaceSetCredentialsLocalRequest) (*RouterInterfaceSetCredentialsLocalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceGetCredentialsLocal(context.Context, *RouterInterfaceGetCredentialsLocalRequest) (*RouterInterfaceGetCredentialsLocalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceGetCustomInfoEx(context.Context, *RouterInterfaceGetCustomInfoExRequest) (*RouterInterfaceGetCustomInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceSetCustomInfoEx(context.Context, *RouterInterfaceSetCustomInfoExRequest) (*RouterInterfaceSetCustomInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DimsvcServer = (*UnimplementedDimsvcServer)(nil)
