package dhcpsrv

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dhcpm "github.com/oiweiwei/go-msrpc/msrpc/dhcpm"
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
	_ = dhcpm.GoPackage
)

var (
	// import guard
	GoPackage = "dhcpm"
)

var (
	// Syntax UUID
	DHCPServerSyntaxUUID = &uuid.UUID{TimeLow: 0x6bffd098, TimeMid: 0xa112, TimeHiAndVersion: 0x3610, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x33, Node: [6]uint8{0x46, 0xc3, 0xf8, 0x74, 0x53, 0x2d}}
	// Syntax ID
	DHCPServerSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: DHCPServerSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// dhcpsrv interface.
type DHCPServerClient interface {

	// The R_DhcpCreateSubnet method is used to create a new IPv4 subnet on the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS            | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR     | An error occurred while accessing the DHCP server database.                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E54 ERROR_DHCP_SUBNET_EXISTS | The IPv4 scope parameters are incorrect. Either the IPv4 scope already exists,   |
	//	|                                     | corresponding to the SubnetAddress and SubnetMask members of the structure       |
	//	|                                     | DHCP_SUBNET_INFO (section 2.2.1.2.8), or there is a range overlap of IPv4        |
	//	|                                     | addresses between those associated with the SubnetAddress and SubnetMask fields  |
	//	|                                     | of the new IPv4 scope and the subnet address and mask of an already existing     |
	//	|                                     | IPv4 scope.                                                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 0.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	CreateSubnet(context.Context, *CreateSubnetRequest, ...dcerpc.CallOption) (*CreateSubnetResponse, error)

	// The R_DhcpSetSubnetInfo method sets/modifies the information about an IPv4 subnet
	// defined on the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed
	// successfully, else it contains a Win32 error code, as specified in [MS-ERREF]. This
	// error code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------------------------+
	//	|                  RETURN                  |                                                             |
	//	|                VALUE/CODE                |                         DESCRIPTION                         |
	//	|                                          |                                                             |
	//	+------------------------------------------+-------------------------------------------------------------+
	//	+------------------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                    |
	//	+------------------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR          | An error occurred while accessing the DHCP server database. |
	//	+------------------------------------------+-------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist.                   |
	//	+------------------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 1.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	SetSubnetInfo(context.Context, *SetSubnetInfoRequest, ...dcerpc.CallOption) (*SetSubnetInfoResponse, error)

	// The R_DhcpGetSubnetInfo method retrieves the information about a specific IPv4 subnet
	// defined on the DHCPv4 server. The caller of this function is responsible for freeing
	// the memory pointed to by SubnetInfo by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------+
	//	|                  RETURN                  |                                           |
	//	|                VALUE/CODE                |                DESCRIPTION                |
	//	|                                          |                                           |
	//	+------------------------------------------+-------------------------------------------+
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                  |
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist. |
	//	+------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 2.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetSubnetInfo(context.Context, *GetSubnetInfoRequest, ...dcerpc.CallOption) (*GetSubnetInfoResponse, error)

	// The R_DhcpEnumSubnets method enumerates IPv4 subnets configured on the DHCPv4 server.
	// The caller of this function can free the memory pointed to by the EnumInfo parameter
	// and its member the Elements array by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------+-----------------------------------------------+
	//	|             RETURN             |                                               |
	//	|           VALUE/CODE           |                  DESCRIPTION                  |
	//	|                                |                                               |
	//	+--------------------------------+-----------------------------------------------+
	//	+--------------------------------+-----------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                      |
	//	+--------------------------------+-----------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate. |
	//	+--------------------------------+-----------------------------------------------+
	//
	// The opnum field value for this method is 3.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnets(context.Context, *EnumSubnetsRequest, ...dcerpc.CallOption) (*EnumSubnetsResponse, error)

	// The R_DhcpAddSubnetElement method adds an IPv4 subnet element (IPv4 reservation,
	// IPv4 exclusion range, or IPv4 range) to the IPv4 subnet in the DHCPv4 server. There
	// is an extension of this method in R_DhcpAddSubnetElementV4 (section 3.1.4.30).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                       RETURN                       |                                                                                  |
	//	|                     VALUE/CODE                     |                                   DESCRIPTION                                    |
	//	|                                                    |                                                                                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                           | The call was successful.                                                         |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                    | An error occurred while accessing the DHCP server database.                      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT           | The specified IPv4 subnet does not exist.                                        |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E37 ERROR_DHCP_INVALID_RANGE                | The specified IPv4 range either overlaps an existing range or is not valid.      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E35 ERROR_DHCP_IPRANGE_EXITS                | The specified IPv4 address range already exists.                                 |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT          | The specified DHCPv4 client is not an IPv4 reserved client.                      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E36 ERROR_DHCP_RESERVEDIP_EXITS             | The specified IPv4 address or hardware address is being used by another DHCPv4   |
	//	|                                                    | client.                                                                          |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E51 ERROR_DHCP_IPRANGE_CONV_ILLEGAL         | Conversion of a BOOTP scope to a DHCP-only scope is illegal, since BOOTP clients |
	//	|                                                    | exist in the scope. Manually delete BOOTP clients from the scope when converting |
	//	|                                                    | the range to a DHCP-only range.                                                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E90 ERROR_SCOPE_RANGE_POLICY_RANGE_CONFLICT | There is an IP address range configured for a policy in this scope. This         |
	//	|                                                    | operation on the scope IP address range cannot be performed until the policy IP  |
	//	|                                                    | address range is suitably modified.                                              |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004EA1 ERROR_DHCP_FO_IPRANGE_TYPE_CONV_ILLEGAL | Conversion of a failover scope to a scope of type BOOTP or BOTH could not be     |
	//	|                                                    | performed. Failover is supported only for DHCP scopes.                           |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 4.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	AddSubnetElement(context.Context, *AddSubnetElementRequest, ...dcerpc.CallOption) (*AddSubnetElementResponse, error)

	// The R_DhcpEnumSubnetElements method enumerates the list of a specific type of IPv4
	// subnet elements from a specific DHCPv4 IPv4 subnet. The caller of this function can
	// free the memory pointed to by the EnumElementInfo parameter and its member the Elements
	// array by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------------+
	//	|                  RETURN                  |                                                 |
	//	|                VALUE/CODE                |                   DESCRIPTION                   |
	//	|                                          |                                                 |
	//	+------------------------------------------+-------------------------------------------------+
	//	+------------------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                        |
	//	+------------------------------------------+-------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA               | There are more elements available to enumerate. |
	//	+------------------------------------------+-------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS           | There are no more elements left to enumerate.   |
	//	+------------------------------------------+-------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist.       |
	//	+------------------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 5.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetElements(context.Context, *EnumSubnetElementsRequest, ...dcerpc.CallOption) (*EnumSubnetElementsResponse, error)

	// The R_DhcpRemoveSubnetElement method removes an IPv4 subnet element from an IPv4
	// subnet defined on the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                       RETURN                       |                                                                                  |
	//	|                     VALUE/CODE                     |                                   DESCRIPTION                                    |
	//	|                                                    |                                                                                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                           | The call was successful.                                                         |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT           | The specified IPv4 subnet does not exist.                                        |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E27 ERROR_DHCP_ELEMENT_CANT_REMOVE          | The specified IPv4 subnet element cannot be removed because at least one IPv4    |
	//	|                                                    | address has been leased out to a client from the subnet. The starting address of |
	//	|                                                    | the specified IPv4 exclusion range is not part of any exclusion range configured |
	//	|                                                    | on the server. There is an error in deleting the exclusion range from the        |
	//	|                                                    | database.                                                                        |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                    | An error occurred while accessing the DHCP server database.                      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E37 ERROR_DHCP_INVALID_RANGE                | The specified IPv4 range does not match an existing IPv4 range.                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E90 ERROR_SCOPE_RANGE_POLICY_RANGE_CONFLICT | There is an IP address range configured for a policy in this scope. This         |
	//	|                                                    | operation cannot be performed until the policy IP range is suitably modified.    |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 6.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveSubnetElement(context.Context, *RemoveSubnetElementRequest, ...dcerpc.CallOption) (*RemoveSubnetElementResponse, error)

	// The R_DhcpDeleteSubnet method deletes an IPv4 subnet from the DHCPv4 server. The
	// ForceFlag defines the behavior of the operation when an IP address from the subnet
	// has been allocated to some DHCPv4 client.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                         RETURN                         |                                                                                  |
	//	|                       VALUE/CODE                       |                                   DESCRIPTION                                    |
	//	|                                                        |                                                                                  |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                               | The call was successful.                                                         |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT               | The specified IPv4 subnet does not exist.                                        |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E27 ERROR_DHCP_ELEMENT_CANT_REMOVE              | The specified IPv4 subnet cannot be removed because at least one IPv4 address    |
	//	|                                                        | has been leased out to some client from the subnet.                              |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                        | An error occurred while accessing the DHCP server database.                      |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E90 ERROR_DHCP_FO_SCOPE_ALREADY_IN_RELATIONSHIP | The specified IPv4 subnet is part of a failover relationship and cannot be       |
	//	|                                                        | deleted without first removing the failover relationship.                        |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 7.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DeleteSubnet(context.Context, *DeleteSubnetRequest, ...dcerpc.CallOption) (*DeleteSubnetResponse, error)

	// The R_DhcpCreateOption method creates an option definition of the specified option
	// for the default user class and vendor class pair at the default option level. The
	// OptionID parameter specifies the identifier of the option.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------+------------------------------------------------------------------+
	//	|               RETURN               |                                                                  |
	//	|             VALUE/CODE             |                           DESCRIPTION                            |
	//	|                                    |                                                                  |
	//	+------------------------------------+------------------------------------------------------------------+
	//	+------------------------------------+------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The call was successful.                                         |
	//	+------------------------------------+------------------------------------------------------------------+
	//	| 0x00004E29 ERROR_DHCP_OPTION_EXITS | The specified option already exists on the DHCP server database. |
	//	+------------------------------------+------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR    | An error occurred while accessing the DHCP server database.      |
	//	+------------------------------------+------------------------------------------------------------------+
	//
	// The opnum field value for this method is 8.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CreateOption(context.Context, *CreateOptionRequest, ...dcerpc.CallOption) (*CreateOptionResponse, error)

	// The R_DhcpSetOptionInfo method modifies the option definition of the specified option
	// for the default user class and vendor class pair at the default option level. There
	// is an extension method R_DhcpSetOptionInfoV5 that sets the option definition for
	// a specific user class and vendor class pair.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                             |
	//	|                VALUE/CODE                |                                 DESCRIPTION                                 |
	//	|                                          |                                                                             |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                    |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The specified option definition does not exist on the DHCP server database. |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR          | An error occurred while accessing the DHCP server database.                 |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 9.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOptionInfo(context.Context, *SetOptionInfoRequest, ...dcerpc.CallOption) (*SetOptionInfoResponse, error)

	// The R_DhcpGetOptionInfo method retrieves the option definition of the specified option
	// for the default user class and vendor class pair at the default option level. There
	// is an extension method R_DhcpGetOptionInfoV5 (section 3.2.4.17) that retrieves the
	// option definition for a specific user class and vendor class pair. The caller of
	// this function can free the memory pointed to by the OptionInfo parameter, by calling
	// the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                             |
	//	|                VALUE/CODE                |                                 DESCRIPTION                                 |
	//	|                                          |                                                                             |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                    |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The specified option definition does not exist on the DHCP server database. |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 10.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetOptionInfo(context.Context, *GetOptionInfoRequest, ...dcerpc.CallOption) (*GetOptionInfoResponse, error)

	// The R_DhcpRemoveOption method removes the option definition of a specific option
	// for the default user class and vendor class pair at the default option level. The
	// OptionID parameter specifies the identifier of the option definition.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS indicates that the operation was completed successfully, else
	// it contains a Win32 error code, as specified in [MS-ERREF]. This error code value
	// can correspond to a DHCP-specific failure, which takes a value between 20000 and
	// 20099, or any generic failure.
	//
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                             |
	//	|                VALUE/CODE                |                                 DESCRIPTION                                 |
	//	|                                          |                                                                             |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                    |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The specified option definition does not exist on the DHCP server database. |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 11.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveOption(context.Context, *RemoveOptionRequest, ...dcerpc.CallOption) (*RemoveOptionResponse, error)

	// The R_DhcpSetOptionValue method creates the option value when called for the first
	// time. Otherwise, it modifies the option value for a specific option associated with
	// the default user class and vendor class pair. The values can be set at the default,
	// server, scope, multicast scope, or reservation level on the DHCPv4 server. The ScopeInfo
	// parameter defines the level at which this option value is set.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                             |
	//	|                VALUE/CODE                 |                                 DESCRIPTION                                 |
	//	|                                           |                                                                             |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                    |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist on the DHCP server.                |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT  | The specified option definition does not exist on the DHCP server database. |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND     | The class name being used is unknown or incorrect.                          |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.                         |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 12.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	SetOptionValue(context.Context, *SetOptionValueRequest, ...dcerpc.CallOption) (*SetOptionValueResponse, error)

	// The R_DhcpGetOptionValue method retrieves the option value for a specific option
	// associated with the default user class and vendor class pair. The values can be retrieved
	// from the default, server, scope, multicast scope, or reservation level on the DHCPv4
	// server. The ScopeInfo parameter defines the level from which the option value needs
	// to be retrieved. The caller of this function can free the memory pointed to by the
	// OptionValue parameter by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                             |
	//	|                VALUE/CODE                 |                                 DESCRIPTION                                 |
	//	|                                           |                                                                             |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                    |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist on the DHCP server.                |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT  | The specified option definition does not exist on the DHCP server database. |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.                         |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 13.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetOptionValue(context.Context, *GetOptionValueRequest, ...dcerpc.CallOption) (*GetOptionValueResponse, error)

	// The R_DhcpEnumOptionValues method enumerates all the option values for the default
	// user class and vendor class pair. The values can be enumerated at a specified level
	// (that is, default, server, scope, multicast scope, or reservation level) defined
	// by the ScopeInfo parameter. The extension of this API is R_DhcpEnumOptionValuesV5
	// (section 3.2.4.23), which retrieves the option values for a specific user class and
	// vendor class at a specific scope defined by the ScopeInfo parameter. The caller of
	// this function can free the memory pointed to by the OptionValues parameter and its
	// member the Values array by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	|                  RETURN                   |                                                              |
	//	|                VALUE/CODE                 |                         DESCRIPTION                          |
	//	|                                           |                                                              |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                     |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA                | There are more elements available to enumerate.              |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS            | There are no more elements left to enumerate.                |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist on the DHCP server. |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.          |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 14.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumOptionValues(context.Context, *EnumOptionValuesRequest, ...dcerpc.CallOption) (*EnumOptionValuesResponse, error)

	// The R_DhcpRemoveOptionValue method removes the option value for a specific option
	// on the DHCPv4 server for the default user class and vendor class. ScopeInfo defines
	// the level (that is, server, scope, multicast scope, or IPv4 reservation level) on
	// which this option value is removed.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                             |
	//	|                VALUE/CODE                 |                                 DESCRIPTION                                 |
	//	|                                           |                                                                             |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                    |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist on the DHCP server.                |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT  | The specified option definition does not exist on the DHCP server database. |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.                         |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 15.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	RemoveOptionValue(context.Context, *RemoveOptionValueRequest, ...dcerpc.CallOption) (*RemoveOptionValueResponse, error)

	// The R_DhcpCreateClientInfo method creates DHCPv4 client lease records on the DHCPv4
	// server database.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+---------------------------------------------------------------+
	//	|             RETURN              |                                                               |
	//	|           VALUE/CODE            |                          DESCRIPTION                          |
	//	|                                 |                                                               |
	//	+---------------------------------+---------------------------------------------------------------+
	//	+---------------------------------+---------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                      |
	//	+---------------------------------+---------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCPv4 server database. |
	//	+---------------------------------+---------------------------------------------------------------+
	//
	// The opnum field value for this method is 16.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CreateClientInfo(context.Context, *CreateClientInfoRequest, ...dcerpc.CallOption) (*CreateClientInfoResponse, error)

	// The R_DhcpSetClientInfo method modifies existing DHCPv4 client lease records on the
	// DHCPv4 server database.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCPv4 server database or the client entry |
	//	|                                 | is not present in the database.                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 17.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetClientInfo(context.Context, *SetClientInfoRequest, ...dcerpc.CallOption) (*SetClientInfoResponse, error)

	// The R_DhcpGetClientInfo method retrieves DHCPv4 client lease record information from
	// the DHCPv4 server database. The caller of this function can free the memory pointed
	// to by the ClientInfo parameter by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database, or the client entry  |
	//	|                                 | is not present in the database.                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 18.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetClientInfo(context.Context, *GetClientInfoRequest, ...dcerpc.CallOption) (*GetClientInfoResponse, error)

	// The R_DhcpDeleteClientInfo method deletes the specified DHCPv4 client lease record
	// from the DHCPv4 server database. It also frees up the DHCPv4 client IPv4 address
	// for redistribution.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | The call was successful.                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR       | An error occurred while accessing the DHCP server database, or the client entry  |
	//	|                                       | is not present in the database.                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E33 ERROR_DHCP_RESERVED_CLIENT | The specified DHCP client is a reserved DHCP client.                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 19.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DeleteClientInfo(context.Context, *DeleteClientInfoRequest, ...dcerpc.CallOption) (*DeleteClientInfoResponse, error)

	// The R_DhcpEnumSubnetClients method is used to retrieve all DHCPv4 clients serviced
	// on the specified IPv4 subnet. This method returns DHCPv4 clients from all IPv4 subnets
	// if the subnet address provided is zero. The caller of this function can free the
	// memory pointed to by the ClientInfo parameter and its member the Clients array by
	// calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA      | There are more elements available to enumerate.             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS  | There are no more elements left to enumerate.               |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 20.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetClients(context.Context, *EnumSubnetClientsRequest, ...dcerpc.CallOption) (*EnumSubnetClientsResponse, error)

	// The R_DhcpGetClientOptions method is never used.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value contains a Win32 error code, as specified in [MS-ERREF]. This error code value
	// can correspond to a DHCP-specific failure, which takes a value between 20000 and
	// 20099, or any generic failure.
	//
	// The opnum field value for this method is 21.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	GetClientOptions(context.Context, *GetClientOptionsRequest, ...dcerpc.CallOption) (*GetClientOptionsResponse, error)

	// The R_DhcpGetMibInfo method is used to retrieve the statistics of the DHCPv4 server.
	// The caller of this function can free the memory pointed to by the MibInfo parameter
	// and its member the ScopeInfo array by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 22.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	GetMIBInfo(context.Context, *GetMIBInfoRequest, ...dcerpc.CallOption) (*GetMIBInfoResponse, error)

	// The R_DhcpEnumOptions method enumerates the option definition for a default user
	// class and vendor class pair specified at the default level. The extension of this
	// method is R_DhcpEnumOptionsV5 (section 3.2.4.18), which enumerates the option definition
	// for the specific user class and vendor class. The caller of this function can free
	// the memory pointed to by the Options parameter and its member the Options array by
	// calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA     | There are more elements available to enumerate. |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate.   |
	//	+--------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 23.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumOptions(context.Context, *EnumOptionsRequest, ...dcerpc.CallOption) (*EnumOptionsResponse, error)

	// The R_DhcpSetOptionValues method creates the option values when called for the first
	// time. Otherwise, it modifies the option values of one or more options at a specified
	// level for a default user class and vendor class pair (that is, at the default, server,
	// scope, multicast scope, or IPv4 reservation level). The ScopeInfo parameter defines
	// the scope on which these option values are modified. The extension of this method
	// is R_DhcpSetOptionValuesV5 (section 3.2.4.21), which sets/modifies the option values
	// of one or more options for a specific user class and vendor class pair.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                             |
	//	|                VALUE/CODE                 |                                 DESCRIPTION                                 |
	//	|                                           |                                                                             |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                    |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist on the DHCP server.                |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT  | The specified option definition does not exist in the DHCP server database. |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND     | The class name being used is unknown or incorrect.                          |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.                         |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 24.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOptionValues(context.Context, *SetOptionValuesRequest, ...dcerpc.CallOption) (*SetOptionValuesResponse, error)

	// The R_DhcpServerSetConfig method sets/modifies the DHCPv4 server settings. There
	// is an extension method R_DhcpServerSetConfigV4 (section 3.1.4.40) that sets some
	// additional settings on the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, the return value contains a Win32 error code, as specified in [MS-ERREF].
	// This error code value can correspond to a DHCP-specific failure, which takes a value
	// between 20000 and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 25.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ServerSetConfig(context.Context, *ServerSetConfigRequest, ...dcerpc.CallOption) (*ServerSetConfigResponse, error)

	// The R_DhcpServerGetConfig method retrieves the current DHCPv4 server setting. There
	// is an extension method R_DhcpServerGetConfigV4 (section 3.1.4.41) that retrieves
	// some additional settings on the DHCPv4 server. The caller of this function can free
	// the memory pointed to by the ConfigInfo parameter by calling the function midl_user_free
	// as specified in 3.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 26.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	ServerGetConfig(context.Context, *ServerGetConfigRequest, ...dcerpc.CallOption) (*ServerGetConfigResponse, error)

	// The R_DhcpScanDatabase method is used by DHCP servers that enumerate and/or fix inconsistencies
	// between the ADM elements DHCPv4 client lease records specified in DHCPv4Scope.DHCPv4ClientsList
	// and the bitmask representation in memory specified in DHCPv4IpRange.BitMask. The
	// caller of this function can free the memory pointed to by the ScanList parameter
	// and its member the ScanItems array by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+--------------------------------------------------------------+
	//	|                  RETURN                  |                                                              |
	//	|                VALUE/CODE                |                         DESCRIPTION                          |
	//	|                                          |                                                              |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                     |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist on the DHCP server. |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR          | An error occurred while accessing the DHCP server database.  |
	//	+------------------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 27.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	ScanDatabase(context.Context, *ScanDatabaseRequest, ...dcerpc.CallOption) (*ScanDatabaseResponse, error)

	// The R_DhcpGetVersion method retrieves the major and minor version numbers of the
	// DHCP server. The version numbers can be used to determine the functionality supported
	// by the DHCP server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 28.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetVersion(context.Context, *GetVersionRequest, ...dcerpc.CallOption) (*GetVersionResponse, error)

	// The R_DhcpAddSubnetElementV4 method adds an IPv4 subnet element (IPv4 reservation
	// for DHCPv4 or BOOTP clients, IPv4 exclusion range, or IPv4 range) to the IPv4 subnet
	// in the DHCPv4 server. There is an extension of this method in R_DhcpAddSubnetElementV5
	// (section 3.2.4.38).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                       RETURN                       |                                                                                  |
	//	|                     VALUE/CODE                     |                                   DESCRIPTION                                    |
	//	|                                                    |                                                                                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                           | The call was successful.                                                         |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT           | The specified IPv4 subnet does not exist on the DHCP server.                     |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                    | An error occurred while accessing the DHCP server database.                      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT          | The specified DHCP client is not a reserved client.                              |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E35 ERROR_DHCP_IPRANGE_EXITS                | The specified IPv4 address range already exists.                                 |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E36 ERROR_DHCP_RESERVEDIP_EXITS             | The specified IPv4 address or hardware address is being used by another DHCP     |
	//	|                                                    | client.                                                                          |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E37 ERROR_DHCP_INVALID_RANGE                | The specified IPv4 range either overlaps an existing range or is not valid.      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E51 ERROR_DHCP_IPRANGE_CONV_ILLEGAL         | Conversion of a BOOTP scope to a DHCP-only scope is illegal, since BOOTP clients |
	//	|                                                    | exist in the scope. Manually delete BOOTP clients from the scope when converting |
	//	|                                                    | the range to a DHCP-only range.                                                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E90 ERROR_SCOPE_RANGE_POLICY_RANGE_CONFLICT | An IP address range is configured for a policy in this scope. This operation     |
	//	|                                                    | cannot be performed on the scope IP address range until the policy IP address    |
	//	|                                                    | range is suitably modified.                                                      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004EA1 ERROR_DHCP_FO_IPRANGE_TYPE_CONV_ILLEGAL | Conversion of a failover scope to a BOOTP-only or BOTH scope cannot be           |
	//	|                                                    | performed. Failover is supported only for DHCP scopes.                           |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 29.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	AddSubnetElementV4(context.Context, *AddSubnetElementV4Request, ...dcerpc.CallOption) (*AddSubnetElementV4Response, error)

	// The R_DhcpEnumSubnetElementsV4 method enumerates the list of a specific type of IPv4
	// subnet element (IPv4 reservation for DHCPv4 or BOOTP clients, IPv4 exclusion range,
	// or IPv4 range) from a specific DHCPv4 IPv4 subnet. The caller of this function can
	// free the memory pointed to by the EnumElementInfo parameter and its member the Elements
	// array by calling the function midl_user_free specified in section 3.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+--------------------------------------------------------------+
	//	|                  RETURN                  |                                                              |
	//	|                VALUE/CODE                |                         DESCRIPTION                          |
	//	|                                          |                                                              |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                     |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA               | There are more elements available to enumerate.              |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS           | There are no more elements left to enumerate.                |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist on the DHCP server. |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR          | An error occurred while accessing the DHCP server database.  |
	//	+------------------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 30.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetElementsV4(context.Context, *EnumSubnetElementsV4Request, ...dcerpc.CallOption) (*EnumSubnetElementsV4Response, error)

	// The R_DhcpRemoveSubnetElementV4 method removes an IPv4 subnet element (IPv4 reservation
	// for DHCPv4 or BOOTP clients, IPv4 exclusion range, or IPv4 range) from an IPv4 subnet
	// defined on the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                       RETURN                       |                                                                                  |
	//	|                     VALUE/CODE                     |                                   DESCRIPTION                                    |
	//	|                                                    |                                                                                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                           | The call was successful.                                                         |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT           | The specified IPv4 subnet does not exist on the DHCP server.                     |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E27 ERROR_DHCP_ELEMENT_CANT_REMOVE          | Failures can occur for any of the following reasons: The specified IPv4 subnet   |
	//	|                                                    | element cannot be removed because at least one IPv4 address has been leased out  |
	//	|                                                    | to a client in the subnet. The starting address of the specified IPv4 exclusion  |
	//	|                                                    | range is not part of any exclusion range configured on the server. There is an   |
	//	|                                                    | error in deleting the exclusion range from the database.                         |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                    | An error occurred while accessing the DHCP server database.                      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E37 ERROR_DHCP_INVALID_RANGE                | The specified IPv4 range does not match an existing IPv4 range.                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E90 ERROR_SCOPE_RANGE_POLICY_RANGE_CONFLICT | An IP address range is configured for a policy in this scope. This operation     |
	//	|                                                    | cannot be performed on the scope IP address range until the policy IP address    |
	//	|                                                    | range is suitably modified.                                                      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 31.
	RemoveSubnetElementV4(context.Context, *RemoveSubnetElementV4Request, ...dcerpc.CallOption) (*RemoveSubnetElementV4Response, error)

	// The R_DhcpCreateClientInfoV4 method sets/modifies the DHCPv4 client lease record
	// on the DHCPv4 server database.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 32.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CreateClientInfoV4(context.Context, *CreateClientInfoV4Request, ...dcerpc.CallOption) (*CreateClientInfoV4Response, error)

	// The R_DhcpSetClientInfoV4 method sets/modifies the existing DHCPv4 client lease record
	// on the DHCPv4 server database.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCPv4 server database or the client entry |
	//	|                                 | is not present in the database.                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 33.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetClientInfoV4(context.Context, *SetClientInfoV4Request, ...dcerpc.CallOption) (*SetClientInfoV4Response, error)

	// The R_DhcpGetClientInfoV4 method retrieves the DHCPv4 client lease record information
	// from the DHCPv4 server database. The caller of this function can free the memory
	// pointed to by the ClientInfo parameter, by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database or the client entry   |
	//	|                                 | is not present in the database.                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 34.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	GetClientInfoV4(context.Context, *GetClientInfoV4Request, ...dcerpc.CallOption) (*GetClientInfoV4Response, error)

	// The R_DhcpEnumSubnetClientsV4 method is used to retrieve all DHCPv4 clients serviced
	// from the specified IPv4 subnet. This method returns the DHCPv4 clients from all IPv4
	// subnets if the subnet address specified zero. The caller of this function can free
	// the memory pointed to by the ClientInfo parameter and its member the Clients array
	// by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA      | There are more elements available to enumerate.             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS  | There are no more elements left to enumerate.               |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 35.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	EnumSubnetClientsV4(context.Context, *EnumSubnetClientsV4Request, ...dcerpc.CallOption) (*EnumSubnetClientsV4Response, error)

	// The R_DhcpSetSuperScopeV4 method adds, modifies, or deletes the IPv4 subnet to/from
	// the superscope information from the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+--------------------------------------------------------------+
	//	|                  RETURN                  |                                                              |
	//	|                VALUE/CODE                |                         DESCRIPTION                          |
	//	|                                          |                                                              |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                     |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist on the DHCP server. |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR          | An error occurred while accessing the DHCP server database.  |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E24 ERROR_DHCP_SUBNET_EXITS       | The specified IPv4 subnet already exists.                    |
	//	+------------------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 36.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	SetSuperScopeV4(context.Context, *SetSuperScopeV4Request, ...dcerpc.CallOption) (*SetSuperScopeV4Response, error)

	// The R_DhcpGetSuperScopeInfoV4 method retrieves all the superscope information from
	// the DHCPv4 server. The caller of this function can free the memory pointed to by
	// the SuperScopeTable parameter and its member the pEntries array by calling the function
	// midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 37.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetSuperScopeInfoV4(context.Context, *GetSuperScopeInfoV4Request, ...dcerpc.CallOption) (*GetSuperScopeInfoV4Response, error)

	// The R_DhcpDeleteSuperScopeV4 method deletes the specified superscope from the DHCPv4
	// server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+--------------------------------------------------------------+
	//	|                  RETURN                  |                                                              |
	//	|                VALUE/CODE                |                         DESCRIPTION                          |
	//	|                                          |                                                              |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                     |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist on the DHCP server. |
	//	+------------------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 38.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DeleteSuperScopeV4(context.Context, *DeleteSuperScopeV4Request, ...dcerpc.CallOption) (*DeleteSuperScopeV4Response, error)

	// The R_DhcpServerSetConfigV4 method sets/modifies the DHCPv4 server settings.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, the return value contains a Win32 error code, as specified in [MS-ERREF].
	// This error code value can correspond to a DHCP-specific failure, which takes a value
	// between 20000 and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 39.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ServerSetConfigV4(context.Context, *ServerSetConfigV4Request, ...dcerpc.CallOption) (*ServerSetConfigV4Response, error)

	// The R_DhcpServerGetConfigV4 method retrieves the current DHCPv4 server setting. The
	// caller of this function can free the memory pointed to by the ConfigInfo parameter
	// by calling the function midl_user_free, specified in section 3.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 40.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ServerGetConfigV4(context.Context, *ServerGetConfigV4Request, ...dcerpc.CallOption) (*ServerGetConfigV4Response, error)

	// The R_DhcpServerSetConfigVQ method sets/modifies the DHCPv4 server settings.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, the return value contains a Win32 error code, as specified in [MS-ERREF].
	// This error code value can correspond to a DHCP-specific failure, which takes a value
	// between 20000 and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 41.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ServerSetConfigVQ(context.Context, *ServerSetConfigVQRequest, ...dcerpc.CallOption) (*ServerSetConfigVQResponse, error)

	// The R_DhcpServerGetConfigVQ method retrieves the current DHCPv4 server setting. The
	// caller of this function can free the memory pointed to by the ConfigInfo parameter
	// by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 42.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	ServerGetConfigVQ(context.Context, *ServerGetConfigVQRequest, ...dcerpc.CallOption) (*ServerGetConfigVQResponse, error)

	// The R_DhcpGetMibInfoVQ method just returns ERROR_SUCCESS. It is never used.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The R_DhcpGetMibInfoVQ method returns only ERROR_SUCCESS. It is never used.
	//
	// The opnum field value for this method is 43.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetMIBInfoVQ(context.Context, *GetMIBInfoVQRequest, ...dcerpc.CallOption) (*GetMIBInfoVQResponse, error)

	// The R_DhcpCreateClientInfoVQ method sets\modifies the DHCPv4 client lease record
	// on the DHCP server database.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 44.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CreateClientInfoVQ(context.Context, *CreateClientInfoVQRequest, ...dcerpc.CallOption) (*CreateClientInfoVQResponse, error)

	// The R_DhcpSetClientInfoVQ method sets/modifies an existing DHCPv4 client lease record
	// on the DHCPv4 server database.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database or the client entry   |
	//	|                                 | is not present in the database.                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 45.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetClientInfoVQ(context.Context, *SetClientInfoVQRequest, ...dcerpc.CallOption) (*SetClientInfoVQResponse, error)

	// The R_DhcpGetClientInfoVQ method retrieves DHCPv4 client lease record information
	// from the DHCPv4 server database. The caller of this function can free the memory
	// pointed to by the ClientInfo parameter, by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database, or the client entry  |
	//	|                                 | is not present in the database.                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 46.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetClientInfoVQ(context.Context, *GetClientInfoVQRequest, ...dcerpc.CallOption) (*GetClientInfoVQResponse, error)

	// The R_DhcpEnumSubnetClientsVQ method is used to retrieve all DHCPv4 clients serviced
	// from the specified IPv4 subnet. This method returns DHCPv4 clients from all IPv4
	// subnets if the subnet address specified is zero. The caller of this function can
	// free the memory pointed to by the ClientInfo parameter and its member the Clients
	// array by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA      | There are more elements available to enumerate.             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS  | There are no more elements left to enumerate.               |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 47.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetClientsVQ(context.Context, *EnumSubnetClientsVQRequest, ...dcerpc.CallOption) (*EnumSubnetClientsVQResponse, error)

	// The R_DhcpCreateSubnetVQ method is used to create the new IPv4 subnet along with
	// its NAP state on the DHCPv4 server. This method is an extension of R_DhcpCreateSubnet
	// (section 3.1.4.1).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS            | The call was successful.                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR     | An error occurred while accessing the DHCP server database.                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E54 ERROR_DHCP_SUBNET_EXISTS | The IPv4 scope parameters are incorrect. Either the IPv4 scope already           |
	//	|                                     | exists, corresponding to the SubnetAddress and SubnetMask members of the         |
	//	|                                     | DHCP_SUBNET_INFO_VQ (section 2.2.1.2.45) structure, or there is a range overlap  |
	//	|                                     | of IPv4 addresses between those associated with the SubnetAddress and SubnetMask |
	//	|                                     | members of the new IPv4 scope and the subnet address and mask of an already      |
	//	|                                     | existing IPv4 scope.                                                             |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 48.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CreateSubnetVQ(context.Context, *CreateSubnetVQRequest, ...dcerpc.CallOption) (*CreateSubnetVQResponse, error)

	// The R_DhcpGetSubnetInfoVQ method retrieves the information about a specific IPv4
	// subnet defined on the DHCPv4 server. This method is an extension of R_DhcpGetSubnetInfo
	// method in which the NAP state is not returned. The caller of this function can free
	// the memory pointed to by the SubnetInfoVQ parameter, by calling the function midl_user_free
	// (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------+
	//	|                  RETURN                  |                                           |
	//	|                VALUE/CODE                |                DESCRIPTION                |
	//	|                                          |                                           |
	//	+------------------------------------------+-------------------------------------------+
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                  |
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist. |
	//	+------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 49.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetSubnetInfoVQ(context.Context, *GetSubnetInfoVQRequest, ...dcerpc.CallOption) (*GetSubnetInfoVQResponse, error)

	// The R_DhcpSetSubnetInfoVQ method sets/modifies the information about an IPv4 subnet
	// defined on the DHCPv4 server. This method is an extension of the R_DhcpSetSubnetInfo
	// method in which NAP state is not set.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------------------------+
	//	|                  RETURN                  |                                                             |
	//	|                VALUE/CODE                |                         DESCRIPTION                         |
	//	|                                          |                                                             |
	//	+------------------------------------------+-------------------------------------------------------------+
	//	+------------------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                    |
	//	+------------------------------------------+-------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist.                   |
	//	+------------------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR          | An error occurred while accessing the DHCP server database. |
	//	+------------------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 50.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetSubnetInfoVQ(context.Context, *SetSubnetInfoVQRequest, ...dcerpc.CallOption) (*SetSubnetInfoVQResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultDHCPServerClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultDHCPServerClient) CreateSubnet(ctx context.Context, in *CreateSubnetRequest, opts ...dcerpc.CallOption) (*CreateSubnetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateSubnetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) SetSubnetInfo(ctx context.Context, in *SetSubnetInfoRequest, opts ...dcerpc.CallOption) (*SetSubnetInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSubnetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetSubnetInfo(ctx context.Context, in *GetSubnetInfoRequest, opts ...dcerpc.CallOption) (*GetSubnetInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSubnetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) EnumSubnets(ctx context.Context, in *EnumSubnetsRequest, opts ...dcerpc.CallOption) (*EnumSubnetsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumSubnetsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) AddSubnetElement(ctx context.Context, in *AddSubnetElementRequest, opts ...dcerpc.CallOption) (*AddSubnetElementResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddSubnetElementResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) EnumSubnetElements(ctx context.Context, in *EnumSubnetElementsRequest, opts ...dcerpc.CallOption) (*EnumSubnetElementsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumSubnetElementsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) RemoveSubnetElement(ctx context.Context, in *RemoveSubnetElementRequest, opts ...dcerpc.CallOption) (*RemoveSubnetElementResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveSubnetElementResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) DeleteSubnet(ctx context.Context, in *DeleteSubnetRequest, opts ...dcerpc.CallOption) (*DeleteSubnetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteSubnetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) CreateOption(ctx context.Context, in *CreateOptionRequest, opts ...dcerpc.CallOption) (*CreateOptionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateOptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) SetOptionInfo(ctx context.Context, in *SetOptionInfoRequest, opts ...dcerpc.CallOption) (*SetOptionInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetOptionInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetOptionInfo(ctx context.Context, in *GetOptionInfoRequest, opts ...dcerpc.CallOption) (*GetOptionInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetOptionInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) RemoveOption(ctx context.Context, in *RemoveOptionRequest, opts ...dcerpc.CallOption) (*RemoveOptionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveOptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) SetOptionValue(ctx context.Context, in *SetOptionValueRequest, opts ...dcerpc.CallOption) (*SetOptionValueResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetOptionValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetOptionValue(ctx context.Context, in *GetOptionValueRequest, opts ...dcerpc.CallOption) (*GetOptionValueResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetOptionValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) EnumOptionValues(ctx context.Context, in *EnumOptionValuesRequest, opts ...dcerpc.CallOption) (*EnumOptionValuesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumOptionValuesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) RemoveOptionValue(ctx context.Context, in *RemoveOptionValueRequest, opts ...dcerpc.CallOption) (*RemoveOptionValueResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveOptionValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) CreateClientInfo(ctx context.Context, in *CreateClientInfoRequest, opts ...dcerpc.CallOption) (*CreateClientInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateClientInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) SetClientInfo(ctx context.Context, in *SetClientInfoRequest, opts ...dcerpc.CallOption) (*SetClientInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetClientInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetClientInfo(ctx context.Context, in *GetClientInfoRequest, opts ...dcerpc.CallOption) (*GetClientInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClientInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) DeleteClientInfo(ctx context.Context, in *DeleteClientInfoRequest, opts ...dcerpc.CallOption) (*DeleteClientInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteClientInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) EnumSubnetClients(ctx context.Context, in *EnumSubnetClientsRequest, opts ...dcerpc.CallOption) (*EnumSubnetClientsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumSubnetClientsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetClientOptions(ctx context.Context, in *GetClientOptionsRequest, opts ...dcerpc.CallOption) (*GetClientOptionsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClientOptionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetMIBInfo(ctx context.Context, in *GetMIBInfoRequest, opts ...dcerpc.CallOption) (*GetMIBInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetMIBInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) EnumOptions(ctx context.Context, in *EnumOptionsRequest, opts ...dcerpc.CallOption) (*EnumOptionsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumOptionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) SetOptionValues(ctx context.Context, in *SetOptionValuesRequest, opts ...dcerpc.CallOption) (*SetOptionValuesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetOptionValuesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) ServerSetConfig(ctx context.Context, in *ServerSetConfigRequest, opts ...dcerpc.CallOption) (*ServerSetConfigResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerSetConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) ServerGetConfig(ctx context.Context, in *ServerGetConfigRequest, opts ...dcerpc.CallOption) (*ServerGetConfigResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerGetConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) ScanDatabase(ctx context.Context, in *ScanDatabaseRequest, opts ...dcerpc.CallOption) (*ScanDatabaseResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ScanDatabaseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...dcerpc.CallOption) (*GetVersionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) AddSubnetElementV4(ctx context.Context, in *AddSubnetElementV4Request, opts ...dcerpc.CallOption) (*AddSubnetElementV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddSubnetElementV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) EnumSubnetElementsV4(ctx context.Context, in *EnumSubnetElementsV4Request, opts ...dcerpc.CallOption) (*EnumSubnetElementsV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumSubnetElementsV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) RemoveSubnetElementV4(ctx context.Context, in *RemoveSubnetElementV4Request, opts ...dcerpc.CallOption) (*RemoveSubnetElementV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveSubnetElementV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) CreateClientInfoV4(ctx context.Context, in *CreateClientInfoV4Request, opts ...dcerpc.CallOption) (*CreateClientInfoV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateClientInfoV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) SetClientInfoV4(ctx context.Context, in *SetClientInfoV4Request, opts ...dcerpc.CallOption) (*SetClientInfoV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetClientInfoV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetClientInfoV4(ctx context.Context, in *GetClientInfoV4Request, opts ...dcerpc.CallOption) (*GetClientInfoV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClientInfoV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) EnumSubnetClientsV4(ctx context.Context, in *EnumSubnetClientsV4Request, opts ...dcerpc.CallOption) (*EnumSubnetClientsV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumSubnetClientsV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) SetSuperScopeV4(ctx context.Context, in *SetSuperScopeV4Request, opts ...dcerpc.CallOption) (*SetSuperScopeV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSuperScopeV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetSuperScopeInfoV4(ctx context.Context, in *GetSuperScopeInfoV4Request, opts ...dcerpc.CallOption) (*GetSuperScopeInfoV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSuperScopeInfoV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) DeleteSuperScopeV4(ctx context.Context, in *DeleteSuperScopeV4Request, opts ...dcerpc.CallOption) (*DeleteSuperScopeV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteSuperScopeV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) ServerSetConfigV4(ctx context.Context, in *ServerSetConfigV4Request, opts ...dcerpc.CallOption) (*ServerSetConfigV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerSetConfigV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) ServerGetConfigV4(ctx context.Context, in *ServerGetConfigV4Request, opts ...dcerpc.CallOption) (*ServerGetConfigV4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerGetConfigV4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) ServerSetConfigVQ(ctx context.Context, in *ServerSetConfigVQRequest, opts ...dcerpc.CallOption) (*ServerSetConfigVQResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerSetConfigVQResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) ServerGetConfigVQ(ctx context.Context, in *ServerGetConfigVQRequest, opts ...dcerpc.CallOption) (*ServerGetConfigVQResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerGetConfigVQResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetMIBInfoVQ(ctx context.Context, in *GetMIBInfoVQRequest, opts ...dcerpc.CallOption) (*GetMIBInfoVQResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetMIBInfoVQResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) CreateClientInfoVQ(ctx context.Context, in *CreateClientInfoVQRequest, opts ...dcerpc.CallOption) (*CreateClientInfoVQResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateClientInfoVQResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) SetClientInfoVQ(ctx context.Context, in *SetClientInfoVQRequest, opts ...dcerpc.CallOption) (*SetClientInfoVQResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetClientInfoVQResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetClientInfoVQ(ctx context.Context, in *GetClientInfoVQRequest, opts ...dcerpc.CallOption) (*GetClientInfoVQResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClientInfoVQResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) EnumSubnetClientsVQ(ctx context.Context, in *EnumSubnetClientsVQRequest, opts ...dcerpc.CallOption) (*EnumSubnetClientsVQResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumSubnetClientsVQResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) CreateSubnetVQ(ctx context.Context, in *CreateSubnetVQRequest, opts ...dcerpc.CallOption) (*CreateSubnetVQResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateSubnetVQResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) GetSubnetInfoVQ(ctx context.Context, in *GetSubnetInfoVQRequest, opts ...dcerpc.CallOption) (*GetSubnetInfoVQResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSubnetInfoVQResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) SetSubnetInfoVQ(ctx context.Context, in *SetSubnetInfoVQRequest, opts ...dcerpc.CallOption) (*SetSubnetInfoVQResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSubnetInfoVQResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDHCPServerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDHCPServerClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewDHCPServerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DHCPServerClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DHCPServerSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultDHCPServerClient{cc: cc}, nil
}

// xxx_CreateSubnetOperation structure represents the R_DhcpCreateSubnet operation
type xxx_CreateSubnetOperation struct {
	ServerIPAddress string            `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32            `idl:"name:SubnetAddress" json:"subnet_address"`
	SubnetInfo      *dhcpm.SubnetInfo `idl:"name:SubnetInfo;pointer:ref" json:"subnet_info"`
	Return          uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateSubnetOperation) OpNum() int { return 0 }

func (o *xxx_CreateSubnetOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpCreateSubnet" }

func (o *xxx_CreateSubnetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSubnetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// SubnetInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_INFO}*(1))(2:{alias=DHCP_SUBNET_INFO}(struct))
	{
		if o.SubnetInfo != nil {
			if err := o.SubnetInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SubnetInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSubnetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// SubnetInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_INFO}*(1))(2:{alias=DHCP_SUBNET_INFO}(struct))
	{
		if o.SubnetInfo == nil {
			o.SubnetInfo = &dhcpm.SubnetInfo{}
		}
		if err := o.SubnetInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSubnetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSubnetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateSubnetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateSubnetRequest structure represents the R_DhcpCreateSubnet operation request
type CreateSubnetRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: A DHCP_IP_ADDRESS that contains the IPv4 subnet address.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// SubnetInfo: This is a pointer to a structure of type LPDHCP_SUBNET_INFO (section
	// 2.2.1.2.8) that contains information about the IPv4 subnet, including the IPv4 subnet
	// mask and the IPv4 address of the subnet. The structure DHCP_HOST_INFO (section 2.2.1.2.7)
	// (referred by PrimaryHost) stored in SubnetInfo MUST be ignored by both the caller
	// and the server.
	SubnetInfo *dhcpm.SubnetInfo `idl:"name:SubnetInfo;pointer:ref" json:"subnet_info"`
}

func (o *CreateSubnetRequest) xxx_ToOp(ctx context.Context) *xxx_CreateSubnetOperation {
	if o == nil {
		return &xxx_CreateSubnetOperation{}
	}
	return &xxx_CreateSubnetOperation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
		SubnetInfo:      o.SubnetInfo,
	}
}

func (o *CreateSubnetRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateSubnetOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.SubnetInfo = op.SubnetInfo
}
func (o *CreateSubnetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateSubnetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateSubnetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateSubnetResponse structure represents the R_DhcpCreateSubnet operation response
type CreateSubnetResponse struct {
	// Return: The R_DhcpCreateSubnet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateSubnetResponse) xxx_ToOp(ctx context.Context) *xxx_CreateSubnetOperation {
	if o == nil {
		return &xxx_CreateSubnetOperation{}
	}
	return &xxx_CreateSubnetOperation{
		Return: o.Return,
	}
}

func (o *CreateSubnetResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateSubnetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CreateSubnetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateSubnetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateSubnetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubnetInfoOperation structure represents the R_DhcpSetSubnetInfo operation
type xxx_SetSubnetInfoOperation struct {
	ServerIPAddress string            `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32            `idl:"name:SubnetAddress" json:"subnet_address"`
	SubnetInfo      *dhcpm.SubnetInfo `idl:"name:SubnetInfo;pointer:ref" json:"subnet_info"`
	Return          uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubnetInfoOperation) OpNum() int { return 1 }

func (o *xxx_SetSubnetInfoOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpSetSubnetInfo" }

func (o *xxx_SetSubnetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubnetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// SubnetInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_INFO}*(1))(2:{alias=DHCP_SUBNET_INFO}(struct))
	{
		if o.SubnetInfo != nil {
			if err := o.SubnetInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SubnetInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubnetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// SubnetInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_INFO}*(1))(2:{alias=DHCP_SUBNET_INFO}(struct))
	{
		if o.SubnetInfo == nil {
			o.SubnetInfo = &dhcpm.SubnetInfo{}
		}
		if err := o.SubnetInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubnetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubnetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSubnetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSubnetInfoRequest structure represents the R_DhcpSetSubnetInfo operation request
type SetSubnetInfoRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS, a DWORD structure containing the
	// IPv4 subnet ID for which the subnet information is modified.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// SubnetInfo: This is a pointer to a DHCP_SUBNET_INFO (section 2.2.1.2.8) structure
	// that contains the information of the IPv4 subnet that is modified in the existing
	// IPv4 subnet identified by SubnetAddress. The DHCP_HOST_INFO (section 2.2.1.2.7) structure
	// (referred by PrimaryHost) stored in SubnetInfo MUST be ignored by both the caller
	// and the server.
	SubnetInfo *dhcpm.SubnetInfo `idl:"name:SubnetInfo;pointer:ref" json:"subnet_info"`
}

func (o *SetSubnetInfoRequest) xxx_ToOp(ctx context.Context) *xxx_SetSubnetInfoOperation {
	if o == nil {
		return &xxx_SetSubnetInfoOperation{}
	}
	return &xxx_SetSubnetInfoOperation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
		SubnetInfo:      o.SubnetInfo,
	}
}

func (o *SetSubnetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubnetInfoOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.SubnetInfo = op.SubnetInfo
}
func (o *SetSubnetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSubnetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubnetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubnetInfoResponse structure represents the R_DhcpSetSubnetInfo operation response
type SetSubnetInfoResponse struct {
	// Return: The R_DhcpSetSubnetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetSubnetInfoResponse) xxx_ToOp(ctx context.Context) *xxx_SetSubnetInfoOperation {
	if o == nil {
		return &xxx_SetSubnetInfoOperation{}
	}
	return &xxx_SetSubnetInfoOperation{
		Return: o.Return,
	}
}

func (o *SetSubnetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubnetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetSubnetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSubnetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubnetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubnetInfoOperation structure represents the R_DhcpGetSubnetInfo operation
type xxx_GetSubnetInfoOperation struct {
	ServerIPAddress string            `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32            `idl:"name:SubnetAddress" json:"subnet_address"`
	SubnetInfo      *dhcpm.SubnetInfo `idl:"name:SubnetInfo" json:"subnet_info"`
	Return          uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubnetInfoOperation) OpNum() int { return 2 }

func (o *xxx_GetSubnetInfoOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetSubnetInfo" }

func (o *xxx_GetSubnetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubnetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubnetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubnetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubnetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SubnetInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SUBNET_INFO}*(1))(3:{alias=DHCP_SUBNET_INFO}(struct))
	{
		if o.SubnetInfo != nil {
			_ptr_SubnetInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubnetInfo != nil {
					if err := o.SubnetInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.SubnetInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubnetInfo, _ptr_SubnetInfo); err != nil {
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

func (o *xxx_GetSubnetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SubnetInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SUBNET_INFO,pointer=ref}*(1))(3:{alias=DHCP_SUBNET_INFO}(struct))
	{
		_ptr_SubnetInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubnetInfo == nil {
				o.SubnetInfo = &dhcpm.SubnetInfo{}
			}
			if err := o.SubnetInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_SubnetInfo := func(ptr interface{}) { o.SubnetInfo = *ptr.(**dhcpm.SubnetInfo) }
		if err := w.ReadPointer(&o.SubnetInfo, _s_SubnetInfo, _ptr_SubnetInfo); err != nil {
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

// GetSubnetInfoRequest structure represents the R_DhcpGetSubnetInfo operation request
type GetSubnetInfoRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS, a DWORD structure containing the
	// IPv4 subnet ID for which the information is retrieved.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
}

func (o *GetSubnetInfoRequest) xxx_ToOp(ctx context.Context) *xxx_GetSubnetInfoOperation {
	if o == nil {
		return &xxx_GetSubnetInfoOperation{}
	}
	return &xxx_GetSubnetInfoOperation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
	}
}

func (o *GetSubnetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubnetInfoOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
}
func (o *GetSubnetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSubnetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubnetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubnetInfoResponse structure represents the R_DhcpGetSubnetInfo operation response
type GetSubnetInfoResponse struct {
	// SubnetInfo: This is a pointer of type LPDHCP_SUBNET_INFO in which the information
	// for the subnet matching the ID specified by SubnetAddress is retrieved.
	SubnetInfo *dhcpm.SubnetInfo `idl:"name:SubnetInfo" json:"subnet_info"`
	// Return: The R_DhcpGetSubnetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetSubnetInfoResponse) xxx_ToOp(ctx context.Context) *xxx_GetSubnetInfoOperation {
	if o == nil {
		return &xxx_GetSubnetInfoOperation{}
	}
	return &xxx_GetSubnetInfoOperation{
		SubnetInfo: o.SubnetInfo,
		Return:     o.Return,
	}
}

func (o *GetSubnetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubnetInfoOperation) {
	if o == nil {
		return
	}
	o.SubnetInfo = op.SubnetInfo
	o.Return = op.Return
}
func (o *GetSubnetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSubnetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubnetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumSubnetsOperation structure represents the R_DhcpEnumSubnets operation
type xxx_EnumSubnetsOperation struct {
	ServerIPAddress  string         `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	Resume           uint32         `idl:"name:ResumeHandle" json:"resume"`
	PreferredMaximum uint32         `idl:"name:PreferredMaximum" json:"preferred_maximum"`
	EnumInfo         *dhcpm.IPArray `idl:"name:EnumInfo" json:"enum_info"`
	ElementsRead     uint32         `idl:"name:ElementsRead" json:"elements_read"`
	ElementsTotal    uint32         `idl:"name:ElementsTotal" json:"elements_total"`
	Return           uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumSubnetsOperation) OpNum() int { return 3 }

func (o *xxx_EnumSubnetsOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpEnumSubnets" }

func (o *xxx_EnumSubnetsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// EnumInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_IP_ARRAY}*(1))(3:{alias=DHCP_IP_ARRAY}(struct))
	{
		if o.EnumInfo != nil {
			_ptr_EnumInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EnumInfo != nil {
					if err := o.EnumInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.IPArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EnumInfo, _ptr_EnumInfo); err != nil {
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
	// ElementsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ElementsRead); err != nil {
			return err
		}
	}
	// ElementsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ElementsTotal); err != nil {
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

func (o *xxx_EnumSubnetsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// EnumInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_IP_ARRAY,pointer=ref}*(1))(3:{alias=DHCP_IP_ARRAY}(struct))
	{
		_ptr_EnumInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EnumInfo == nil {
				o.EnumInfo = &dhcpm.IPArray{}
			}
			if err := o.EnumInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EnumInfo := func(ptr interface{}) { o.EnumInfo = *ptr.(**dhcpm.IPArray) }
		if err := w.ReadPointer(&o.EnumInfo, _s_EnumInfo, _ptr_EnumInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ElementsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ElementsRead); err != nil {
			return err
		}
	}
	// ElementsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ElementsTotal); err != nil {
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

// EnumSubnetsRequest structure represents the R_DhcpEnumSubnets operation request
type EnumSubnetsRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6), a
	// DWORD that identifies the enumeration operation. Initially, this value MUST be set
	// to zero, with a successful call returning the handle value used for subsequent enumeration
	// requests. For example, if PreferredMaximum is set to 100, and 200 IPv4 subnet addresses
	// are stored on the DHCPv4 server, the resume handle can be used after the first 100
	// IPv4 subnets are retrieved to obtain the next 100 on a subsequent call.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// PreferredMaximum: This is of type DWORD, specifying the preferred maximum number
	// of IPv4 subnet addresses to return. If the number of remaining unenumerated IPv4
	// subnets is less than this value, all the IPv4 subnets configured on DHCPv4 server
	// are returned. To retrieve all the IPv4 subnets defined on the DHCPv4 server, 0xFFFFFFFF
	// is specified.
	PreferredMaximum uint32 `idl:"name:PreferredMaximum" json:"preferred_maximum"`
}

func (o *EnumSubnetsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetsOperation {
	if o == nil {
		return &xxx_EnumSubnetsOperation{}
	}
	return &xxx_EnumSubnetsOperation{
		ServerIPAddress:  o.ServerIPAddress,
		Resume:           o.Resume,
		PreferredMaximum: o.PreferredMaximum,
	}
}

func (o *EnumSubnetsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetsOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.Resume = op.Resume
	o.PreferredMaximum = op.PreferredMaximum
}
func (o *EnumSubnetsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumSubnetsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumSubnetsResponse structure represents the R_DhcpEnumSubnets operation response
type EnumSubnetsResponse struct {
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6), a
	// DWORD that identifies the enumeration operation. Initially, this value MUST be set
	// to zero, with a successful call returning the handle value used for subsequent enumeration
	// requests. For example, if PreferredMaximum is set to 100, and 200 IPv4 subnet addresses
	// are stored on the DHCPv4 server, the resume handle can be used after the first 100
	// IPv4 subnets are retrieved to obtain the next 100 on a subsequent call.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// EnumInfo: This is a pointer of type LPDHCP_IP_ARRAY that points to the location in
	// which the IPv4 subnet configured on the DHCPv4 server is returned.
	EnumInfo *dhcpm.IPArray `idl:"name:EnumInfo" json:"enum_info"`
	// ElementsRead: This is a pointer to a DWORD value that specifies the number of IPv4
	// subnet addresses returned in EnumInfo. The caller MUST allocate memory for this parameter
	// equal to the size of data type DWORD.
	ElementsRead uint32 `idl:"name:ElementsRead" json:"elements_read"`
	// ElementsTotal: This is a pointer to a DWORD value that specifies the number of IPv4
	// subnets defined on the DHCPv4 server that have not yet been enumerated with respect
	// to the resume handle that is returned. The caller MUST allocate memory for this parameter
	// equal to the size of data type DWORD.
	ElementsTotal uint32 `idl:"name:ElementsTotal" json:"elements_total"`
	// Return: The R_DhcpEnumSubnets return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumSubnetsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetsOperation {
	if o == nil {
		return &xxx_EnumSubnetsOperation{}
	}
	return &xxx_EnumSubnetsOperation{
		Resume:        o.Resume,
		EnumInfo:      o.EnumInfo,
		ElementsRead:  o.ElementsRead,
		ElementsTotal: o.ElementsTotal,
		Return:        o.Return,
	}
}

func (o *EnumSubnetsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetsOperation) {
	if o == nil {
		return
	}
	o.Resume = op.Resume
	o.EnumInfo = op.EnumInfo
	o.ElementsRead = op.ElementsRead
	o.ElementsTotal = op.ElementsTotal
	o.Return = op.Return
}
func (o *EnumSubnetsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumSubnetsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddSubnetElementOperation structure represents the R_DhcpAddSubnetElement operation
type xxx_AddSubnetElementOperation struct {
	ServerIPAddress string                   `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32                   `idl:"name:SubnetAddress" json:"subnet_address"`
	AddElementInfo  *dhcpm.SubnetElementData `idl:"name:AddElementInfo;pointer:ref" json:"add_element_info"`
	Return          uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_AddSubnetElementOperation) OpNum() int { return 4 }

func (o *xxx_AddSubnetElementOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpAddSubnetElement" }

func (o *xxx_AddSubnetElementOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddSubnetElementOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// AddElementInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_ELEMENT_DATA}*(1))(2:{alias=DHCP_SUBNET_ELEMENT_DATA}(struct))
	{
		if o.AddElementInfo != nil {
			if err := o.AddElementInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SubnetElementData{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddSubnetElementOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// AddElementInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_ELEMENT_DATA}*(1))(2:{alias=DHCP_SUBNET_ELEMENT_DATA}(struct))
	{
		if o.AddElementInfo == nil {
			o.AddElementInfo = &dhcpm.SubnetElementData{}
		}
		if err := o.AddElementInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddSubnetElementOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddSubnetElementOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddSubnetElementOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddSubnetElementRequest structure represents the R_DhcpAddSubnetElement operation request
type AddSubnetElementRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS, containing the IPv4 subnet ID in
	// which the IPv4 subnet element is added.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// AddElementInfo: This is a pointer to a structure DHCP_SUBNET_ELEMENT_DATA (section
	// 2.2.1.2.33) that contains the IPv4 subnet element that needs to be added to the IPv4
	// subnet.
	AddElementInfo *dhcpm.SubnetElementData `idl:"name:AddElementInfo;pointer:ref" json:"add_element_info"`
}

func (o *AddSubnetElementRequest) xxx_ToOp(ctx context.Context) *xxx_AddSubnetElementOperation {
	if o == nil {
		return &xxx_AddSubnetElementOperation{}
	}
	return &xxx_AddSubnetElementOperation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
		AddElementInfo:  o.AddElementInfo,
	}
}

func (o *AddSubnetElementRequest) xxx_FromOp(ctx context.Context, op *xxx_AddSubnetElementOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.AddElementInfo = op.AddElementInfo
}
func (o *AddSubnetElementRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddSubnetElementRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddSubnetElementOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddSubnetElementResponse structure represents the R_DhcpAddSubnetElement operation response
type AddSubnetElementResponse struct {
	// Return: The R_DhcpAddSubnetElement return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddSubnetElementResponse) xxx_ToOp(ctx context.Context) *xxx_AddSubnetElementOperation {
	if o == nil {
		return &xxx_AddSubnetElementOperation{}
	}
	return &xxx_AddSubnetElementOperation{
		Return: o.Return,
	}
}

func (o *AddSubnetElementResponse) xxx_FromOp(ctx context.Context, op *xxx_AddSubnetElementOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddSubnetElementResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddSubnetElementResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddSubnetElementOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumSubnetElementsOperation structure represents the R_DhcpEnumSubnetElements operation
type xxx_EnumSubnetElementsOperation struct {
	ServerIPAddress  string                        `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress    uint32                        `idl:"name:SubnetAddress" json:"subnet_address"`
	EnumElementType  dhcpm.SubnetElementType       `idl:"name:EnumElementType" json:"enum_element_type"`
	Resume           uint32                        `idl:"name:ResumeHandle" json:"resume"`
	PreferredMaximum uint32                        `idl:"name:PreferredMaximum" json:"preferred_maximum"`
	EnumElementInfo  *dhcpm.SubnetElementInfoArray `idl:"name:EnumElementInfo" json:"enum_element_info"`
	ElementsRead     uint32                        `idl:"name:ElementsRead" json:"elements_read"`
	ElementsTotal    uint32                        `idl:"name:ElementsTotal" json:"elements_total"`
	Return           uint32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumSubnetElementsOperation) OpNum() int { return 5 }

func (o *xxx_EnumSubnetElementsOperation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpEnumSubnetElements"
}

func (o *xxx_EnumSubnetElementsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetElementsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// EnumElementType {in} (1:{alias=DHCP_SUBNET_ELEMENT_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.EnumElementType)); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetElementsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// EnumElementType {in} (1:{alias=DHCP_SUBNET_ELEMENT_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.EnumElementType)); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetElementsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetElementsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// EnumElementInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SUBNET_ELEMENT_INFO_ARRAY}*(1))(3:{alias=DHCP_SUBNET_ELEMENT_INFO_ARRAY}(struct))
	{
		if o.EnumElementInfo != nil {
			_ptr_EnumElementInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EnumElementInfo != nil {
					if err := o.EnumElementInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.SubnetElementInfoArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EnumElementInfo, _ptr_EnumElementInfo); err != nil {
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
	// ElementsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ElementsRead); err != nil {
			return err
		}
	}
	// ElementsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ElementsTotal); err != nil {
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

func (o *xxx_EnumSubnetElementsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// EnumElementInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SUBNET_ELEMENT_INFO_ARRAY,pointer=ref}*(1))(3:{alias=DHCP_SUBNET_ELEMENT_INFO_ARRAY}(struct))
	{
		_ptr_EnumElementInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EnumElementInfo == nil {
				o.EnumElementInfo = &dhcpm.SubnetElementInfoArray{}
			}
			if err := o.EnumElementInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EnumElementInfo := func(ptr interface{}) { o.EnumElementInfo = *ptr.(**dhcpm.SubnetElementInfoArray) }
		if err := w.ReadPointer(&o.EnumElementInfo, _s_EnumElementInfo, _ptr_EnumElementInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ElementsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ElementsRead); err != nil {
			return err
		}
	}
	// ElementsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ElementsTotal); err != nil {
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

// EnumSubnetElementsRequest structure represents the R_DhcpEnumSubnetElements operation request
type EnumSubnetElementsRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing the
	// IPv4 subnet ID from which subnet elements are enumerated.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// EnumElementType: This is of type DHCP_SUBNET_ELEMENT_TYPE (section 2.2.1.1.7) enumeration,
	// indicating the type of IPv4 subnet element to enumerate.
	EnumElementType dhcpm.SubnetElementType `idl:"name:EnumElementType" json:"enum_element_type"`
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) a
	// DWORD that identifies the enumeration operation. Initially, this value MUST be set
	// to zero, with a successful call returning the handle value used for subsequent enumeration
	// requests. For example, if the PreferredMaximum parameter is set to 1,000 bytes, and
	// 2,000 bytes worth of IPv4 subnet elements are stored on the DHCPv4 server, the resume
	// handle can be used after the first 1,000 bytes are retrieved to obtain the next 1,000
	// on a subsequent call, and so forth.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// PreferredMaximum: This is of type DWORD that specifies the preferred maximum number
	// of bytes to return. If the number of remaining unenumerated subnet elements (in bytes)
	// is less than this value, all IPv4 subnet elements for a specific type are returned.
	// To retrieve all the IPv4 subnet elements of a specific type, 0xFFFFFFFF is specified.
	PreferredMaximum uint32 `idl:"name:PreferredMaximum" json:"preferred_maximum"`
}

func (o *EnumSubnetElementsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetElementsOperation {
	if o == nil {
		return &xxx_EnumSubnetElementsOperation{}
	}
	return &xxx_EnumSubnetElementsOperation{
		ServerIPAddress:  o.ServerIPAddress,
		SubnetAddress:    o.SubnetAddress,
		EnumElementType:  o.EnumElementType,
		Resume:           o.Resume,
		PreferredMaximum: o.PreferredMaximum,
	}
}

func (o *EnumSubnetElementsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetElementsOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.EnumElementType = op.EnumElementType
	o.Resume = op.Resume
	o.PreferredMaximum = op.PreferredMaximum
}
func (o *EnumSubnetElementsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumSubnetElementsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetElementsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumSubnetElementsResponse structure represents the R_DhcpEnumSubnetElements operation response
type EnumSubnetElementsResponse struct {
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) a
	// DWORD that identifies the enumeration operation. Initially, this value MUST be set
	// to zero, with a successful call returning the handle value used for subsequent enumeration
	// requests. For example, if the PreferredMaximum parameter is set to 1,000 bytes, and
	// 2,000 bytes worth of IPv4 subnet elements are stored on the DHCPv4 server, the resume
	// handle can be used after the first 1,000 bytes are retrieved to obtain the next 1,000
	// on a subsequent call, and so forth.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// EnumElementInfo: This is a pointer of type LPDHCP_SUBNET_ELEMENT_INFO_ARRAY in which
	// an IPv4 subnet element of type EnumElementType is returned for a specific IPv4 subnet
	// SubnetAddress. If no IPv4 subnet element of a specific type is available for enumeration,
	// this value is null.
	EnumElementInfo *dhcpm.SubnetElementInfoArray `idl:"name:EnumElementInfo" json:"enum_element_info"`
	// ElementsRead: This is a pointer to a DWORD value that specifies the number of IPv4
	// subnet elements read in the EnumElementInfo parameter for a specific type of IPv4
	// subnet element. The caller MUST allocate memory for this parameter equal to the size
	// of data type DWORD.
	ElementsRead uint32 `idl:"name:ElementsRead" json:"elements_read"`
	// ElementsTotal: This is a pointer to a DWORD value that specifies the number of IPv4
	// subnet elements of a specific type from a specific IPv4 subnet that are not yet enumerated
	// with respect to the resume handle that is returned. The caller MUST allocate memory
	// for this parameter equal to the size of data type DWORD.
	ElementsTotal uint32 `idl:"name:ElementsTotal" json:"elements_total"`
	// Return: The R_DhcpEnumSubnetElements return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumSubnetElementsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetElementsOperation {
	if o == nil {
		return &xxx_EnumSubnetElementsOperation{}
	}
	return &xxx_EnumSubnetElementsOperation{
		Resume:          o.Resume,
		EnumElementInfo: o.EnumElementInfo,
		ElementsRead:    o.ElementsRead,
		ElementsTotal:   o.ElementsTotal,
		Return:          o.Return,
	}
}

func (o *EnumSubnetElementsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetElementsOperation) {
	if o == nil {
		return
	}
	o.Resume = op.Resume
	o.EnumElementInfo = op.EnumElementInfo
	o.ElementsRead = op.ElementsRead
	o.ElementsTotal = op.ElementsTotal
	o.Return = op.Return
}
func (o *EnumSubnetElementsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumSubnetElementsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetElementsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveSubnetElementOperation structure represents the R_DhcpRemoveSubnetElement operation
type xxx_RemoveSubnetElementOperation struct {
	ServerIPAddress   string                   `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress     uint32                   `idl:"name:SubnetAddress" json:"subnet_address"`
	RemoveElementInfo *dhcpm.SubnetElementData `idl:"name:RemoveElementInfo;pointer:ref" json:"remove_element_info"`
	ForceFlag         dhcpm.ForceFlag          `idl:"name:ForceFlag" json:"force_flag"`
	Return            uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveSubnetElementOperation) OpNum() int { return 6 }

func (o *xxx_RemoveSubnetElementOperation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpRemoveSubnetElement"
}

func (o *xxx_RemoveSubnetElementOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubnetElementOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// RemoveElementInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_ELEMENT_DATA}*(1))(2:{alias=DHCP_SUBNET_ELEMENT_DATA}(struct))
	{
		if o.RemoveElementInfo != nil {
			if err := o.RemoveElementInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SubnetElementData{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ForceFlag {in} (1:{alias=DHCP_FORCE_FLAG}(enum))
	{
		if err := w.WriteEnum(uint16(o.ForceFlag)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubnetElementOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// RemoveElementInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_ELEMENT_DATA}*(1))(2:{alias=DHCP_SUBNET_ELEMENT_DATA}(struct))
	{
		if o.RemoveElementInfo == nil {
			o.RemoveElementInfo = &dhcpm.SubnetElementData{}
		}
		if err := o.RemoveElementInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ForceFlag {in} (1:{alias=DHCP_FORCE_FLAG}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ForceFlag)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubnetElementOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubnetElementOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveSubnetElementOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveSubnetElementRequest structure represents the R_DhcpRemoveSubnetElement operation request
type RemoveSubnetElementRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1) that contains
	// the IPv4 subnet ID from which subnet elements are enumerated.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// RemoveElementInfo: This is a pointer of type DHCP_SUBNET_ELEMENT_DATA (section 2.2.1.2.33),
	// containing the IPv4 subnet element that needs to be removed from the IPv4 subnet.
	RemoveElementInfo *dhcpm.SubnetElementData `idl:"name:RemoveElementInfo;pointer:ref" json:"remove_element_info"`
	// ForceFlag: This is of type DHCP_FORCE_FLAG (section 2.2.1.1.9), defining the behavior
	// of this method. If the flag is set to DhcpNoForce and this subnet has served the
	// IPv4 address to some DHCPv4\BOOTP clients, the IPv4 range is not deleted; if the
	// flag is set to DhcpFullForce, the IPv4 range is deleted along with DHCPv4 client
	// lease record on the DHCPv4 server.
	ForceFlag dhcpm.ForceFlag `idl:"name:ForceFlag" json:"force_flag"`
}

func (o *RemoveSubnetElementRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveSubnetElementOperation {
	if o == nil {
		return &xxx_RemoveSubnetElementOperation{}
	}
	return &xxx_RemoveSubnetElementOperation{
		ServerIPAddress:   o.ServerIPAddress,
		SubnetAddress:     o.SubnetAddress,
		RemoveElementInfo: o.RemoveElementInfo,
		ForceFlag:         o.ForceFlag,
	}
}

func (o *RemoveSubnetElementRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveSubnetElementOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.RemoveElementInfo = op.RemoveElementInfo
	o.ForceFlag = op.ForceFlag
}
func (o *RemoveSubnetElementRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveSubnetElementRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveSubnetElementOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveSubnetElementResponse structure represents the R_DhcpRemoveSubnetElement operation response
type RemoveSubnetElementResponse struct {
	// Return: The R_DhcpRemoveSubnetElement return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveSubnetElementResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveSubnetElementOperation {
	if o == nil {
		return &xxx_RemoveSubnetElementOperation{}
	}
	return &xxx_RemoveSubnetElementOperation{
		Return: o.Return,
	}
}

func (o *RemoveSubnetElementResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveSubnetElementOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveSubnetElementResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveSubnetElementResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveSubnetElementOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteSubnetOperation structure represents the R_DhcpDeleteSubnet operation
type xxx_DeleteSubnetOperation struct {
	ServerIPAddress string          `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32          `idl:"name:SubnetAddress" json:"subnet_address"`
	ForceFlag       dhcpm.ForceFlag `idl:"name:ForceFlag" json:"force_flag"`
	Return          uint32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteSubnetOperation) OpNum() int { return 7 }

func (o *xxx_DeleteSubnetOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpDeleteSubnet" }

func (o *xxx_DeleteSubnetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteSubnetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// ForceFlag {in} (1:{alias=DHCP_FORCE_FLAG}(enum))
	{
		if err := w.WriteEnum(uint16(o.ForceFlag)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteSubnetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// ForceFlag {in} (1:{alias=DHCP_FORCE_FLAG}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ForceFlag)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteSubnetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteSubnetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteSubnetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteSubnetRequest structure represents the R_DhcpDeleteSubnet operation request
type DeleteSubnetRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), uniquely identifying
	// the IPv4 subnet that needs to be removed from the DHCPv4 server.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// ForceFlag: This is of type DHCP_FORCE_FLAG (section 2.2.1.1.9) enumeration. If the
	// flag is set to DhcpNoForce and an IPv4 address from this subnet has been served to
	// some DHCPv4/BOOTP client, the IPv4 subnet is not deleted; if the flag is set to DhcpFullForce,
	// the IPv4 subnet is deleted along with the DHCPv4 client lease record on the DHCPv4
	// server; if the flag is set to DhcpFailoverForce, the IPv4 subnet is deleted along
	// with the DHCPv4 client lease record on the DHCPv4 server but the DNS resource records
	// corresponding to the DHCPv4 client lease records are not deleted from the DNS server.
	ForceFlag dhcpm.ForceFlag `idl:"name:ForceFlag" json:"force_flag"`
}

func (o *DeleteSubnetRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteSubnetOperation {
	if o == nil {
		return &xxx_DeleteSubnetOperation{}
	}
	return &xxx_DeleteSubnetOperation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
		ForceFlag:       o.ForceFlag,
	}
}

func (o *DeleteSubnetRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteSubnetOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.ForceFlag = op.ForceFlag
}
func (o *DeleteSubnetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteSubnetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteSubnetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteSubnetResponse structure represents the R_DhcpDeleteSubnet operation response
type DeleteSubnetResponse struct {
	// Return: The R_DhcpDeleteSubnet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteSubnetResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteSubnetOperation {
	if o == nil {
		return &xxx_DeleteSubnetOperation{}
	}
	return &xxx_DeleteSubnetOperation{
		Return: o.Return,
	}
}

func (o *DeleteSubnetResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteSubnetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteSubnetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteSubnetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteSubnetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateOptionOperation structure represents the R_DhcpCreateOption operation
type xxx_CreateOptionOperation struct {
	ServerIPAddress string        `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	OptionID        uint32        `idl:"name:OptionID" json:"option_id"`
	OptionInfo      *dhcpm.Option `idl:"name:OptionInfo;pointer:ref" json:"option_info"`
	Return          uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateOptionOperation) OpNum() int { return 8 }

func (o *xxx_CreateOptionOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpCreateOption" }

func (o *xxx_CreateOptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateOptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionID); err != nil {
			return err
		}
	}
	// OptionInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION}*(1))(2:{alias=DHCP_OPTION}(struct))
	{
		if o.OptionInfo != nil {
			if err := o.OptionInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.Option{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateOptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionID); err != nil {
			return err
		}
	}
	// OptionInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION}*(1))(2:{alias=DHCP_OPTION}(struct))
	{
		if o.OptionInfo == nil {
			o.OptionInfo = &dhcpm.Option{}
		}
		if err := o.OptionInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateOptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateOptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateOptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateOptionRequest structure represents the R_DhcpCreateOption operation request
type CreateOptionRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// OptionID: This is of type DHCP_OPTION_ID (section 2.2.1.2.3), containing the option
	// identifier for the option being created.
	OptionID uint32 `idl:"name:OptionID" json:"option_id"`
	// OptionInfo: This is a pointer to a DHCP_OPTION (section 2.2.1.2.25) structure that
	// contains the information about the option definition.
	OptionInfo *dhcpm.Option `idl:"name:OptionInfo;pointer:ref" json:"option_info"`
}

func (o *CreateOptionRequest) xxx_ToOp(ctx context.Context) *xxx_CreateOptionOperation {
	if o == nil {
		return &xxx_CreateOptionOperation{}
	}
	return &xxx_CreateOptionOperation{
		ServerIPAddress: o.ServerIPAddress,
		OptionID:        o.OptionID,
		OptionInfo:      o.OptionInfo,
	}
}

func (o *CreateOptionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateOptionOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.OptionID = op.OptionID
	o.OptionInfo = op.OptionInfo
}
func (o *CreateOptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateOptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateOptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateOptionResponse structure represents the R_DhcpCreateOption operation response
type CreateOptionResponse struct {
	// Return: The R_DhcpCreateOption return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateOptionResponse) xxx_ToOp(ctx context.Context) *xxx_CreateOptionOperation {
	if o == nil {
		return &xxx_CreateOptionOperation{}
	}
	return &xxx_CreateOptionOperation{
		Return: o.Return,
	}
}

func (o *CreateOptionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateOptionOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CreateOptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateOptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateOptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOptionInfoOperation structure represents the R_DhcpSetOptionInfo operation
type xxx_SetOptionInfoOperation struct {
	ServerIPAddress string        `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	OptionID        uint32        `idl:"name:OptionID" json:"option_id"`
	OptionInfo      *dhcpm.Option `idl:"name:OptionInfo;pointer:ref" json:"option_info"`
	Return          uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOptionInfoOperation) OpNum() int { return 9 }

func (o *xxx_SetOptionInfoOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpSetOptionInfo" }

func (o *xxx_SetOptionInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionID); err != nil {
			return err
		}
	}
	// OptionInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION}*(1))(2:{alias=DHCP_OPTION}(struct))
	{
		if o.OptionInfo != nil {
			if err := o.OptionInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.Option{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionID); err != nil {
			return err
		}
	}
	// OptionInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION}*(1))(2:{alias=DHCP_OPTION}(struct))
	{
		if o.OptionInfo == nil {
			o.OptionInfo = &dhcpm.Option{}
		}
		if err := o.OptionInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOptionInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetOptionInfoRequest structure represents the R_DhcpSetOptionInfo operation request
type SetOptionInfoRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// OptionID: This is of type DHCP_OPTION_ID (section 2.2.1.2.3), containing the option
	// identifier for the option being modified.
	OptionID uint32 `idl:"name:OptionID" json:"option_id"`
	// OptionInfo: This is a pointer of type DHCP_OPTION (section 2.2.1.2.25) structure,
	// containing the option definition for the option being modified.
	OptionInfo *dhcpm.Option `idl:"name:OptionInfo;pointer:ref" json:"option_info"`
}

func (o *SetOptionInfoRequest) xxx_ToOp(ctx context.Context) *xxx_SetOptionInfoOperation {
	if o == nil {
		return &xxx_SetOptionInfoOperation{}
	}
	return &xxx_SetOptionInfoOperation{
		ServerIPAddress: o.ServerIPAddress,
		OptionID:        o.OptionID,
		OptionInfo:      o.OptionInfo,
	}
}

func (o *SetOptionInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOptionInfoOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.OptionID = op.OptionID
	o.OptionInfo = op.OptionInfo
}
func (o *SetOptionInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetOptionInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOptionInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOptionInfoResponse structure represents the R_DhcpSetOptionInfo operation response
type SetOptionInfoResponse struct {
	// Return: The R_DhcpSetOptionInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetOptionInfoResponse) xxx_ToOp(ctx context.Context) *xxx_SetOptionInfoOperation {
	if o == nil {
		return &xxx_SetOptionInfoOperation{}
	}
	return &xxx_SetOptionInfoOperation{
		Return: o.Return,
	}
}

func (o *SetOptionInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOptionInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetOptionInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetOptionInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOptionInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetOptionInfoOperation structure represents the R_DhcpGetOptionInfo operation
type xxx_GetOptionInfoOperation struct {
	ServerIPAddress string        `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	OptionID        uint32        `idl:"name:OptionID" json:"option_id"`
	OptionInfo      *dhcpm.Option `idl:"name:OptionInfo" json:"option_info"`
	Return          uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOptionInfoOperation) OpNum() int { return 10 }

func (o *xxx_GetOptionInfoOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetOptionInfo" }

func (o *xxx_GetOptionInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOptionInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOptionInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOptionInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOptionInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// OptionInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_OPTION}*(1))(3:{alias=DHCP_OPTION}(struct))
	{
		if o.OptionInfo != nil {
			_ptr_OptionInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OptionInfo != nil {
					if err := o.OptionInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.Option{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OptionInfo, _ptr_OptionInfo); err != nil {
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

func (o *xxx_GetOptionInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// OptionInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_OPTION,pointer=ref}*(1))(3:{alias=DHCP_OPTION}(struct))
	{
		_ptr_OptionInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OptionInfo == nil {
				o.OptionInfo = &dhcpm.Option{}
			}
			if err := o.OptionInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OptionInfo := func(ptr interface{}) { o.OptionInfo = *ptr.(**dhcpm.Option) }
		if err := w.ReadPointer(&o.OptionInfo, _s_OptionInfo, _ptr_OptionInfo); err != nil {
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

// GetOptionInfoRequest structure represents the R_DhcpGetOptionInfo operation request
type GetOptionInfoRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// OptionID: This is of type DHCP_OPTION_ID (section 2.2.1.2.3), containing the option
	// identifier for the option being retrieved.
	OptionID uint32 `idl:"name:OptionID" json:"option_id"`
}

func (o *GetOptionInfoRequest) xxx_ToOp(ctx context.Context) *xxx_GetOptionInfoOperation {
	if o == nil {
		return &xxx_GetOptionInfoOperation{}
	}
	return &xxx_GetOptionInfoOperation{
		ServerIPAddress: o.ServerIPAddress,
		OptionID:        o.OptionID,
	}
}

func (o *GetOptionInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOptionInfoOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.OptionID = op.OptionID
}
func (o *GetOptionInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetOptionInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOptionInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOptionInfoResponse structure represents the R_DhcpGetOptionInfo operation response
type GetOptionInfoResponse struct {
	// OptionInfo: This is a pointer of type LPDHCP_OPTION.
	OptionInfo *dhcpm.Option `idl:"name:OptionInfo" json:"option_info"`
	// Return: The R_DhcpGetOptionInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetOptionInfoResponse) xxx_ToOp(ctx context.Context) *xxx_GetOptionInfoOperation {
	if o == nil {
		return &xxx_GetOptionInfoOperation{}
	}
	return &xxx_GetOptionInfoOperation{
		OptionInfo: o.OptionInfo,
		Return:     o.Return,
	}
}

func (o *GetOptionInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOptionInfoOperation) {
	if o == nil {
		return
	}
	o.OptionInfo = op.OptionInfo
	o.Return = op.Return
}
func (o *GetOptionInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetOptionInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOptionInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveOptionOperation structure represents the R_DhcpRemoveOption operation
type xxx_RemoveOptionOperation struct {
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	OptionID        uint32 `idl:"name:OptionID" json:"option_id"`
	Return          uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveOptionOperation) OpNum() int { return 11 }

func (o *xxx_RemoveOptionOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpRemoveOption" }

func (o *xxx_RemoveOptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveOptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveOptionRequest structure represents the R_DhcpRemoveOption operation request
type RemoveOptionRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// OptionID: This is of type DHCP_OPTION_ID (section 2.2.1.2.3), containing the option
	// identifier for the option definition being removed.
	OptionID uint32 `idl:"name:OptionID" json:"option_id"`
}

func (o *RemoveOptionRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveOptionOperation {
	if o == nil {
		return &xxx_RemoveOptionOperation{}
	}
	return &xxx_RemoveOptionOperation{
		ServerIPAddress: o.ServerIPAddress,
		OptionID:        o.OptionID,
	}
}

func (o *RemoveOptionRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveOptionOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.OptionID = op.OptionID
}
func (o *RemoveOptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveOptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveOptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveOptionResponse structure represents the R_DhcpRemoveOption operation response
type RemoveOptionResponse struct {
	// Return: The R_DhcpRemoveOption return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveOptionResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveOptionOperation {
	if o == nil {
		return &xxx_RemoveOptionOperation{}
	}
	return &xxx_RemoveOptionOperation{
		Return: o.Return,
	}
}

func (o *RemoveOptionResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveOptionOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveOptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveOptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveOptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOptionValueOperation structure represents the R_DhcpSetOptionValue operation
type xxx_SetOptionValueOperation struct {
	ServerIPAddress string                 `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	OptionID        uint32                 `idl:"name:OptionID" json:"option_id"`
	ScopeInfo       *dhcpm.OptionScopeInfo `idl:"name:ScopeInfo;pointer:ref" json:"scope_info"`
	OptionValue     *dhcpm.OptionData      `idl:"name:OptionValue;pointer:ref" json:"option_value"`
	Return          uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOptionValueOperation) OpNum() int { return 12 }

func (o *xxx_SetOptionValueOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpSetOptionValue" }

func (o *xxx_SetOptionValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionID); err != nil {
			return err
		}
	}
	// ScopeInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION_SCOPE_INFO}*(1))(2:{alias=DHCP_OPTION_SCOPE_INFO}(struct))
	{
		if o.ScopeInfo != nil {
			if err := o.ScopeInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.OptionScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// OptionValue {in} (1:{pointer=ref, alias=LPDHCP_OPTION_DATA}*(1))(2:{alias=DHCP_OPTION_DATA}(struct))
	{
		if o.OptionValue != nil {
			if err := o.OptionValue.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.OptionData{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionID); err != nil {
			return err
		}
	}
	// ScopeInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION_SCOPE_INFO}*(1))(2:{alias=DHCP_OPTION_SCOPE_INFO}(struct))
	{
		if o.ScopeInfo == nil {
			o.ScopeInfo = &dhcpm.OptionScopeInfo{}
		}
		if err := o.ScopeInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionValue {in} (1:{pointer=ref, alias=LPDHCP_OPTION_DATA}*(1))(2:{alias=DHCP_OPTION_DATA}(struct))
	{
		if o.OptionValue == nil {
			o.OptionValue = &dhcpm.OptionData{}
		}
		if err := o.OptionValue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOptionValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetOptionValueRequest structure represents the R_DhcpSetOptionValue operation request
type SetOptionValueRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// OptionID: This is of type DHCP_OPTION_ID (section 2.2.1.2.3), containing the option
	// identifier for the option being set or modified.
	OptionID uint32 `idl:"name:OptionID" json:"option_id"`
	// ScopeInfo: This is a pointer to a DHCP_OPTION_SCOPE_INFO (section 2.2.1.2.41) structure
	// that contains information describing the level (that is, default, server, scope,
	// multicast scope, or reservation level) at which this option value is set on.
	ScopeInfo *dhcpm.OptionScopeInfo `idl:"name:ScopeInfo;pointer:ref" json:"scope_info"`
	// OptionValue: A pointer to a DHCP_OPTION_DATA (section 2.2.1.2.24) structure that
	// contains the option value to be set. For Dynamic DNS update settings, see section
	// 3.3.1.
	OptionValue *dhcpm.OptionData `idl:"name:OptionValue;pointer:ref" json:"option_value"`
}

func (o *SetOptionValueRequest) xxx_ToOp(ctx context.Context) *xxx_SetOptionValueOperation {
	if o == nil {
		return &xxx_SetOptionValueOperation{}
	}
	return &xxx_SetOptionValueOperation{
		ServerIPAddress: o.ServerIPAddress,
		OptionID:        o.OptionID,
		ScopeInfo:       o.ScopeInfo,
		OptionValue:     o.OptionValue,
	}
}

func (o *SetOptionValueRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOptionValueOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.OptionID = op.OptionID
	o.ScopeInfo = op.ScopeInfo
	o.OptionValue = op.OptionValue
}
func (o *SetOptionValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetOptionValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOptionValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOptionValueResponse structure represents the R_DhcpSetOptionValue operation response
type SetOptionValueResponse struct {
	// Return: The R_DhcpSetOptionValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetOptionValueResponse) xxx_ToOp(ctx context.Context) *xxx_SetOptionValueOperation {
	if o == nil {
		return &xxx_SetOptionValueOperation{}
	}
	return &xxx_SetOptionValueOperation{
		Return: o.Return,
	}
}

func (o *SetOptionValueResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOptionValueOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetOptionValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetOptionValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOptionValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetOptionValueOperation structure represents the R_DhcpGetOptionValue operation
type xxx_GetOptionValueOperation struct {
	ServerIPAddress string                 `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	OptionID        uint32                 `idl:"name:OptionID" json:"option_id"`
	ScopeInfo       *dhcpm.OptionScopeInfo `idl:"name:ScopeInfo;pointer:ref" json:"scope_info"`
	OptionValue     *dhcpm.OptionValue     `idl:"name:OptionValue" json:"option_value"`
	Return          uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOptionValueOperation) OpNum() int { return 13 }

func (o *xxx_GetOptionValueOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetOptionValue" }

func (o *xxx_GetOptionValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOptionValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionID); err != nil {
			return err
		}
	}
	// ScopeInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION_SCOPE_INFO}*(1))(2:{alias=DHCP_OPTION_SCOPE_INFO}(struct))
	{
		if o.ScopeInfo != nil {
			if err := o.ScopeInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.OptionScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOptionValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionID); err != nil {
			return err
		}
	}
	// ScopeInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION_SCOPE_INFO}*(1))(2:{alias=DHCP_OPTION_SCOPE_INFO}(struct))
	{
		if o.ScopeInfo == nil {
			o.ScopeInfo = &dhcpm.OptionScopeInfo{}
		}
		if err := o.ScopeInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOptionValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOptionValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// OptionValue {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_OPTION_VALUE}*(1))(3:{alias=DHCP_OPTION_VALUE}(struct))
	{
		if o.OptionValue != nil {
			_ptr_OptionValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OptionValue != nil {
					if err := o.OptionValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.OptionValue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OptionValue, _ptr_OptionValue); err != nil {
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

func (o *xxx_GetOptionValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// OptionValue {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_OPTION_VALUE,pointer=ref}*(1))(3:{alias=DHCP_OPTION_VALUE}(struct))
	{
		_ptr_OptionValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OptionValue == nil {
				o.OptionValue = &dhcpm.OptionValue{}
			}
			if err := o.OptionValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OptionValue := func(ptr interface{}) { o.OptionValue = *ptr.(**dhcpm.OptionValue) }
		if err := w.ReadPointer(&o.OptionValue, _s_OptionValue, _ptr_OptionValue); err != nil {
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

// GetOptionValueRequest structure represents the R_DhcpGetOptionValue operation request
type GetOptionValueRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// OptionID: This is of type DHCP_OPTION_ID (section 2.2.1.2.3), containing the option
	// identifier for the option being retrieved.
	OptionID uint32 `idl:"name:OptionID" json:"option_id"`
	// ScopeInfo: This is a pointer to a DHCP_OPTION_SCOPE_INFO (section 2.2.1.2.41) structure
	// that contains information describing the level (that is, default, server, scope,
	// multicast scope, or reservation level) from which the option value is retrieved.
	ScopeInfo *dhcpm.OptionScopeInfo `idl:"name:ScopeInfo;pointer:ref" json:"scope_info"`
}

func (o *GetOptionValueRequest) xxx_ToOp(ctx context.Context) *xxx_GetOptionValueOperation {
	if o == nil {
		return &xxx_GetOptionValueOperation{}
	}
	return &xxx_GetOptionValueOperation{
		ServerIPAddress: o.ServerIPAddress,
		OptionID:        o.OptionID,
		ScopeInfo:       o.ScopeInfo,
	}
}

func (o *GetOptionValueRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOptionValueOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.OptionID = op.OptionID
	o.ScopeInfo = op.ScopeInfo
}
func (o *GetOptionValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetOptionValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOptionValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOptionValueResponse structure represents the R_DhcpGetOptionValue operation response
type GetOptionValueResponse struct {
	// OptionValue: This is a pointer of type LPDHCP_OPTION_VALUE in which the option value
	// is retrieved corresponding to the OptionID parameter. For Dynamic DNS update settings,
	// see section 3.3.1.
	OptionValue *dhcpm.OptionValue `idl:"name:OptionValue" json:"option_value"`
	// Return: The R_DhcpGetOptionValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetOptionValueResponse) xxx_ToOp(ctx context.Context) *xxx_GetOptionValueOperation {
	if o == nil {
		return &xxx_GetOptionValueOperation{}
	}
	return &xxx_GetOptionValueOperation{
		OptionValue: o.OptionValue,
		Return:      o.Return,
	}
}

func (o *GetOptionValueResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOptionValueOperation) {
	if o == nil {
		return
	}
	o.OptionValue = op.OptionValue
	o.Return = op.Return
}
func (o *GetOptionValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetOptionValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOptionValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumOptionValuesOperation structure represents the R_DhcpEnumOptionValues operation
type xxx_EnumOptionValuesOperation struct {
	ServerIPAddress  string                  `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ScopeInfo        *dhcpm.OptionScopeInfo  `idl:"name:ScopeInfo;pointer:ref" json:"scope_info"`
	Resume           uint32                  `idl:"name:ResumeHandle" json:"resume"`
	PreferredMaximum uint32                  `idl:"name:PreferredMaximum" json:"preferred_maximum"`
	OptionValues     *dhcpm.OptionValueArray `idl:"name:OptionValues" json:"option_values"`
	OptionsRead      uint32                  `idl:"name:OptionsRead" json:"options_read"`
	OptionsTotal     uint32                  `idl:"name:OptionsTotal" json:"options_total"`
	Return           uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumOptionValuesOperation) OpNum() int { return 14 }

func (o *xxx_EnumOptionValuesOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpEnumOptionValues" }

func (o *xxx_EnumOptionValuesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOptionValuesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ScopeInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION_SCOPE_INFO}*(1))(2:{alias=DHCP_OPTION_SCOPE_INFO}(struct))
	{
		if o.ScopeInfo != nil {
			if err := o.ScopeInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.OptionScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOptionValuesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ScopeInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION_SCOPE_INFO}*(1))(2:{alias=DHCP_OPTION_SCOPE_INFO}(struct))
	{
		if o.ScopeInfo == nil {
			o.ScopeInfo = &dhcpm.OptionScopeInfo{}
		}
		if err := o.ScopeInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOptionValuesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOptionValuesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// OptionValues {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_OPTION_VALUE_ARRAY}*(1))(3:{alias=DHCP_OPTION_VALUE_ARRAY}(struct))
	{
		if o.OptionValues != nil {
			_ptr_OptionValues := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OptionValues != nil {
					if err := o.OptionValues.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.OptionValueArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OptionValues, _ptr_OptionValues); err != nil {
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
	// OptionsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionsRead); err != nil {
			return err
		}
	}
	// OptionsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionsTotal); err != nil {
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

func (o *xxx_EnumOptionValuesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// OptionValues {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_OPTION_VALUE_ARRAY,pointer=ref}*(1))(3:{alias=DHCP_OPTION_VALUE_ARRAY}(struct))
	{
		_ptr_OptionValues := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OptionValues == nil {
				o.OptionValues = &dhcpm.OptionValueArray{}
			}
			if err := o.OptionValues.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OptionValues := func(ptr interface{}) { o.OptionValues = *ptr.(**dhcpm.OptionValueArray) }
		if err := w.ReadPointer(&o.OptionValues, _s_OptionValues, _ptr_OptionValues); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionsRead); err != nil {
			return err
		}
	}
	// OptionsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionsTotal); err != nil {
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

// EnumOptionValuesRequest structure represents the R_DhcpEnumOptionValues operation request
type EnumOptionValuesRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ScopeInfo: This is a pointer to a DHCP_OPTION_SCOPE_INFO (section 2.2.1.2.41) structure
	// that contains information describing the level (that is, default, server, scope,
	// multicast scope, or IPv4 reservation level) at which the option values are enumerated
	// on.
	ScopeInfo *dhcpm.OptionScopeInfo `idl:"name:ScopeInfo;pointer:ref" json:"scope_info"`
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies whether the enumeration operation is a continuation of a previous operation.
	// Initially, this value MUST be set to zero, with a successful call returning the handle
	// value used for subsequent enumeration requests. For example, if the PreferredMaximum
	// parameter is set to 1,000 bytes, and 2,000 bytes worth of option values are stored
	// on the DHCPv4 server, the resume handle can be used after the first 1,000 bytes are
	// retrieved to obtain the next 1,000 on a subsequent call, and so forth.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// PreferredMaximum: This is of type DWORD, specifying the preferred maximum number
	// of bytes to return. If the number of remaining unenumerated option values (in bytes)
	// is less than this value, all option values are returned. To retrieve all the option
	// values for the default user class and vendor class at the desired level, 0xFFFFFFFF
	// is specified.
	PreferredMaximum uint32 `idl:"name:PreferredMaximum" json:"preferred_maximum"`
}

func (o *EnumOptionValuesRequest) xxx_ToOp(ctx context.Context) *xxx_EnumOptionValuesOperation {
	if o == nil {
		return &xxx_EnumOptionValuesOperation{}
	}
	return &xxx_EnumOptionValuesOperation{
		ServerIPAddress:  o.ServerIPAddress,
		ScopeInfo:        o.ScopeInfo,
		Resume:           o.Resume,
		PreferredMaximum: o.PreferredMaximum,
	}
}

func (o *EnumOptionValuesRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumOptionValuesOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.ScopeInfo = op.ScopeInfo
	o.Resume = op.Resume
	o.PreferredMaximum = op.PreferredMaximum
}
func (o *EnumOptionValuesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumOptionValuesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumOptionValuesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumOptionValuesResponse structure represents the R_DhcpEnumOptionValues operation response
type EnumOptionValuesResponse struct {
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies whether the enumeration operation is a continuation of a previous operation.
	// Initially, this value MUST be set to zero, with a successful call returning the handle
	// value used for subsequent enumeration requests. For example, if the PreferredMaximum
	// parameter is set to 1,000 bytes, and 2,000 bytes worth of option values are stored
	// on the DHCPv4 server, the resume handle can be used after the first 1,000 bytes are
	// retrieved to obtain the next 1,000 on a subsequent call, and so forth.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// OptionValues: This is a pointer of type LPDHCP_OPTION_VALUE_ARRAY in which all the
	// option values for the default user class and vendor class are retrieved at a specific
	// level (that is, default, server, scope, multicast scope, or IPv4 reservation level)
	// corresponding to the ScopeInfo parameter.
	OptionValues *dhcpm.OptionValueArray `idl:"name:OptionValues" json:"option_values"`
	// OptionsRead: This is a pointer to a DWORD value that specifies the number of option
	// values read in the OptionValues parameter. The caller MUST allocate memory for this
	// parameter equal to the size of data type DWORD.
	OptionsRead uint32 `idl:"name:OptionsRead" json:"options_read"`
	// OptionsTotal: This is a pointer to a DWORD value that specifies the number of option
	// values that have not yet been read. The caller MUST allocate memory for this parameter
	// equal to the size of data type DWORD.
	OptionsTotal uint32 `idl:"name:OptionsTotal" json:"options_total"`
	// Return: The R_DhcpEnumOptionValues return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumOptionValuesResponse) xxx_ToOp(ctx context.Context) *xxx_EnumOptionValuesOperation {
	if o == nil {
		return &xxx_EnumOptionValuesOperation{}
	}
	return &xxx_EnumOptionValuesOperation{
		Resume:       o.Resume,
		OptionValues: o.OptionValues,
		OptionsRead:  o.OptionsRead,
		OptionsTotal: o.OptionsTotal,
		Return:       o.Return,
	}
}

func (o *EnumOptionValuesResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumOptionValuesOperation) {
	if o == nil {
		return
	}
	o.Resume = op.Resume
	o.OptionValues = op.OptionValues
	o.OptionsRead = op.OptionsRead
	o.OptionsTotal = op.OptionsTotal
	o.Return = op.Return
}
func (o *EnumOptionValuesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumOptionValuesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumOptionValuesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveOptionValueOperation structure represents the R_DhcpRemoveOptionValue operation
type xxx_RemoveOptionValueOperation struct {
	ServerIPAddress string                 `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	OptionID        uint32                 `idl:"name:OptionID" json:"option_id"`
	ScopeInfo       *dhcpm.OptionScopeInfo `idl:"name:ScopeInfo;pointer:ref" json:"scope_info"`
	Return          uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveOptionValueOperation) OpNum() int { return 15 }

func (o *xxx_RemoveOptionValueOperation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpRemoveOptionValue"
}

func (o *xxx_RemoveOptionValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOptionValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionID); err != nil {
			return err
		}
	}
	// ScopeInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION_SCOPE_INFO}*(1))(2:{alias=DHCP_OPTION_SCOPE_INFO}(struct))
	{
		if o.ScopeInfo != nil {
			if err := o.ScopeInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.OptionScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOptionValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionID {in} (1:{alias=DHCP_OPTION_ID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionID); err != nil {
			return err
		}
	}
	// ScopeInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION_SCOPE_INFO}*(1))(2:{alias=DHCP_OPTION_SCOPE_INFO}(struct))
	{
		if o.ScopeInfo == nil {
			o.ScopeInfo = &dhcpm.OptionScopeInfo{}
		}
		if err := o.ScopeInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOptionValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOptionValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveOptionValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveOptionValueRequest structure represents the R_DhcpRemoveOptionValue operation request
type RemoveOptionValueRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// OptionID: This is of type DHCP_OPTION_ID (section 2.2.1.2.3), containing the option
	// identifier for the option being removed.
	OptionID uint32 `idl:"name:OptionID" json:"option_id"`
	// ScopeInfo: This is a pointer to a DHCP_OPTION_SCOPE_INFO (section 2.2.1.2.41) structure
	// that contains information describing the level (that is, server, scope, multicast
	// scope, or IPv4 reservation level) from which this option value is removed.
	ScopeInfo *dhcpm.OptionScopeInfo `idl:"name:ScopeInfo;pointer:ref" json:"scope_info"`
}

func (o *RemoveOptionValueRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveOptionValueOperation {
	if o == nil {
		return &xxx_RemoveOptionValueOperation{}
	}
	return &xxx_RemoveOptionValueOperation{
		ServerIPAddress: o.ServerIPAddress,
		OptionID:        o.OptionID,
		ScopeInfo:       o.ScopeInfo,
	}
}

func (o *RemoveOptionValueRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveOptionValueOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.OptionID = op.OptionID
	o.ScopeInfo = op.ScopeInfo
}
func (o *RemoveOptionValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveOptionValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveOptionValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveOptionValueResponse structure represents the R_DhcpRemoveOptionValue operation response
type RemoveOptionValueResponse struct {
	// Return: The R_DhcpRemoveOptionValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveOptionValueResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveOptionValueOperation {
	if o == nil {
		return &xxx_RemoveOptionValueOperation{}
	}
	return &xxx_RemoveOptionValueOperation{
		Return: o.Return,
	}
}

func (o *RemoveOptionValueResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveOptionValueOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveOptionValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveOptionValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveOptionValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateClientInfoOperation structure represents the R_DhcpCreateClientInfo operation
type xxx_CreateClientInfoOperation struct {
	ServerIPAddress string            `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ClientInfo      *dhcpm.ClientInfo `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
	Return          uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateClientInfoOperation) OpNum() int { return 16 }

func (o *xxx_CreateClientInfoOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpCreateClientInfo" }

func (o *xxx_CreateClientInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO}*(1))(2:{alias=DHCP_CLIENT_INFO}(struct))
	{
		if o.ClientInfo != nil {
			if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.ClientInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO}*(1))(2:{alias=DHCP_CLIENT_INFO}(struct))
	{
		if o.ClientInfo == nil {
			o.ClientInfo = &dhcpm.ClientInfo{}
		}
		if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateClientInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateClientInfoRequest structure represents the R_DhcpCreateClientInfo operation request
type CreateClientInfoRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ClientInfo: A pointer of type DHCP_CLIENT_INFO (section 2.2.1.2.12) structure that
	// contains the DHCPv4 client lease record information that needs to be set on the DHCPv4
	// server. The caller MUST pass the ClientIPAddress and ClientHardwareAddress member
	// fields to add the DHCPv4 client lease record to the DHCPv4 server database. The ClientHardwareAddress
	// member represents a DHCPv4 client-identifier (section 2.2.1.2.5.1). Members ClientName,
	// ClientComment, ClientLeaseExpires, and OwnerHost are modified on the DHCPv4 client
	// lease record identified by the ClientIPaddress member.
	ClientInfo *dhcpm.ClientInfo `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
}

func (o *CreateClientInfoRequest) xxx_ToOp(ctx context.Context) *xxx_CreateClientInfoOperation {
	if o == nil {
		return &xxx_CreateClientInfoOperation{}
	}
	return &xxx_CreateClientInfoOperation{
		ServerIPAddress: o.ServerIPAddress,
		ClientInfo:      o.ClientInfo,
	}
}

func (o *CreateClientInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateClientInfoOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.ClientInfo = op.ClientInfo
}
func (o *CreateClientInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateClientInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateClientInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateClientInfoResponse structure represents the R_DhcpCreateClientInfo operation response
type CreateClientInfoResponse struct {
	// Return: The R_DhcpCreateClientInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateClientInfoResponse) xxx_ToOp(ctx context.Context) *xxx_CreateClientInfoOperation {
	if o == nil {
		return &xxx_CreateClientInfoOperation{}
	}
	return &xxx_CreateClientInfoOperation{
		Return: o.Return,
	}
}

func (o *CreateClientInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateClientInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CreateClientInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateClientInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateClientInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClientInfoOperation structure represents the R_DhcpSetClientInfo operation
type xxx_SetClientInfoOperation struct {
	ServerIPAddress string            `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ClientInfo      *dhcpm.ClientInfo `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
	Return          uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClientInfoOperation) OpNum() int { return 17 }

func (o *xxx_SetClientInfoOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpSetClientInfo" }

func (o *xxx_SetClientInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO}*(1))(2:{alias=DHCP_CLIENT_INFO}(struct))
	{
		if o.ClientInfo != nil {
			if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.ClientInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO}*(1))(2:{alias=DHCP_CLIENT_INFO}(struct))
	{
		if o.ClientInfo == nil {
			o.ClientInfo = &dhcpm.ClientInfo{}
		}
		if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetClientInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetClientInfoRequest structure represents the R_DhcpSetClientInfo operation request
type SetClientInfoRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ClientInfo: A pointer of type DHCP_CLIENT_INFO (section 2.2.1.2.12) structure that
	// contains the DHCPv4 client lease record information that needs to be modified within
	// the DHCPv4 server database. The caller MUST pass the ClientIPAddress and ClientHardwareAddress
	// member fields for modification of a DHCPv4 client lease record stored in the DHCPv4
	// server database. The ClientHardwareAddress member field represents a DHCPv4 client-identifier
	// (section 2.2.1.2.5.1). The members ClientName, ClientComment, ClientLeaseExpires,
	// and OwnerHost are modified in the DHCPv4 client lease record identified by the ClientIPAddress
	// member.
	ClientInfo *dhcpm.ClientInfo `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
}

func (o *SetClientInfoRequest) xxx_ToOp(ctx context.Context) *xxx_SetClientInfoOperation {
	if o == nil {
		return &xxx_SetClientInfoOperation{}
	}
	return &xxx_SetClientInfoOperation{
		ServerIPAddress: o.ServerIPAddress,
		ClientInfo:      o.ClientInfo,
	}
}

func (o *SetClientInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClientInfoOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.ClientInfo = op.ClientInfo
}
func (o *SetClientInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetClientInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClientInfoResponse structure represents the R_DhcpSetClientInfo operation response
type SetClientInfoResponse struct {
	// Return: The R_DhcpSetClientInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetClientInfoResponse) xxx_ToOp(ctx context.Context) *xxx_SetClientInfoOperation {
	if o == nil {
		return &xxx_SetClientInfoOperation{}
	}
	return &xxx_SetClientInfoOperation{
		Return: o.Return,
	}
}

func (o *SetClientInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClientInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetClientInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetClientInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClientInfoOperation structure represents the R_DhcpGetClientInfo operation
type xxx_GetClientInfoOperation struct {
	ServerIPAddress string            `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SearchInfo      *dhcpm.SearchInfo `idl:"name:SearchInfo;pointer:ref" json:"search_info"`
	ClientInfo      *dhcpm.ClientInfo `idl:"name:ClientInfo" json:"client_info"`
	Return          uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClientInfoOperation) OpNum() int { return 18 }

func (o *xxx_GetClientInfoOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetClientInfo" }

func (o *xxx_GetClientInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SearchInfo {in} (1:{pointer=ref, alias=LPDHCP_SEARCH_INFO}*(1))(2:{alias=DHCP_SEARCH_INFO}(struct))
	{
		if o.SearchInfo != nil {
			if err := o.SearchInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SearchInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SearchInfo {in} (1:{pointer=ref, alias=LPDHCP_SEARCH_INFO}*(1))(2:{alias=DHCP_SEARCH_INFO}(struct))
	{
		if o.SearchInfo == nil {
			o.SearchInfo = &dhcpm.SearchInfo{}
		}
		if err := o.SearchInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO}*(1))(3:{alias=DHCP_CLIENT_INFO}(struct))
	{
		if o.ClientInfo != nil {
			_ptr_ClientInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientInfo != nil {
					if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.ClientInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientInfo, _ptr_ClientInfo); err != nil {
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

func (o *xxx_GetClientInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO,pointer=ref}*(1))(3:{alias=DHCP_CLIENT_INFO}(struct))
	{
		_ptr_ClientInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientInfo == nil {
				o.ClientInfo = &dhcpm.ClientInfo{}
			}
			if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ClientInfo := func(ptr interface{}) { o.ClientInfo = *ptr.(**dhcpm.ClientInfo) }
		if err := w.ReadPointer(&o.ClientInfo, _s_ClientInfo, _ptr_ClientInfo); err != nil {
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

// GetClientInfoRequest structure represents the R_DhcpGetClientInfo operation request
type GetClientInfoRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SearchInfo: This is a pointer of type DHCP_SEARCH_INFO (section 2.2.1.2.18) structure
	// that defines the key to be used to search the DHCPv4 client lease record on the DHCPv4
	// server. In case the SearchType member is DhcpClientName and there are multiple lease
	// records with the same ClientName, the server will return client information for the
	// client having the lowest numerical IP address.
	SearchInfo *dhcpm.SearchInfo `idl:"name:SearchInfo;pointer:ref" json:"search_info"`
}

func (o *GetClientInfoRequest) xxx_ToOp(ctx context.Context) *xxx_GetClientInfoOperation {
	if o == nil {
		return &xxx_GetClientInfoOperation{}
	}
	return &xxx_GetClientInfoOperation{
		ServerIPAddress: o.ServerIPAddress,
		SearchInfo:      o.SearchInfo,
	}
}

func (o *GetClientInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClientInfoOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SearchInfo = op.SearchInfo
}
func (o *GetClientInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClientInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClientInfoResponse structure represents the R_DhcpGetClientInfo operation response
type GetClientInfoResponse struct {
	// ClientInfo: This is a pointer of type LPDHCP_CLIENT_INFO that points to the location
	// in which specific DHCPv4 client lease record information is retrieved. The ClientHardwareAddress
	// member represents a DHCPv4 client unique ID (section 2.2.1.2.5.2)
	ClientInfo *dhcpm.ClientInfo `idl:"name:ClientInfo" json:"client_info"`
	// Return: The R_DhcpGetClientInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetClientInfoResponse) xxx_ToOp(ctx context.Context) *xxx_GetClientInfoOperation {
	if o == nil {
		return &xxx_GetClientInfoOperation{}
	}
	return &xxx_GetClientInfoOperation{
		ClientInfo: o.ClientInfo,
		Return:     o.Return,
	}
}

func (o *GetClientInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClientInfoOperation) {
	if o == nil {
		return
	}
	o.ClientInfo = op.ClientInfo
	o.Return = op.Return
}
func (o *GetClientInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClientInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteClientInfoOperation structure represents the R_DhcpDeleteClientInfo operation
type xxx_DeleteClientInfoOperation struct {
	ServerIPAddress string            `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ClientInfo      *dhcpm.SearchInfo `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
	Return          uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteClientInfoOperation) OpNum() int { return 19 }

func (o *xxx_DeleteClientInfoOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpDeleteClientInfo" }

func (o *xxx_DeleteClientInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClientInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_SEARCH_INFO}*(1))(2:{alias=DHCP_SEARCH_INFO}(struct))
	{
		if o.ClientInfo != nil {
			if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SearchInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClientInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_SEARCH_INFO}*(1))(2:{alias=DHCP_SEARCH_INFO}(struct))
	{
		if o.ClientInfo == nil {
			o.ClientInfo = &dhcpm.SearchInfo{}
		}
		if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClientInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteClientInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteClientInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteClientInfoRequest structure represents the R_DhcpDeleteClientInfo operation request
type DeleteClientInfoRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is not
	// used.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ClientInfo: This is of type DHCP_SEARCH_INFO (section 2.2.1.2.18) structure, defining
	// the key to be used to search the DHCPv4 client lease record that needs to be deleted
	// on the DHCPv4 server. In case the SearchType member is DhcpClientName and there are
	// multiple lease records with the same ClientName member, the server will delete the
	// lease record for any of the clients with that client name.
	ClientInfo *dhcpm.SearchInfo `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
}

func (o *DeleteClientInfoRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteClientInfoOperation {
	if o == nil {
		return &xxx_DeleteClientInfoOperation{}
	}
	return &xxx_DeleteClientInfoOperation{
		ServerIPAddress: o.ServerIPAddress,
		ClientInfo:      o.ClientInfo,
	}
}

func (o *DeleteClientInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteClientInfoOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.ClientInfo = op.ClientInfo
}
func (o *DeleteClientInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteClientInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteClientInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteClientInfoResponse structure represents the R_DhcpDeleteClientInfo operation response
type DeleteClientInfoResponse struct {
	// Return: The R_DhcpDeleteClientInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteClientInfoResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteClientInfoOperation {
	if o == nil {
		return &xxx_DeleteClientInfoOperation{}
	}
	return &xxx_DeleteClientInfoOperation{
		Return: o.Return,
	}
}

func (o *DeleteClientInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteClientInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteClientInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteClientInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteClientInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumSubnetClientsOperation structure represents the R_DhcpEnumSubnetClients operation
type xxx_EnumSubnetClientsOperation struct {
	ServerIPAddress  string                 `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress    uint32                 `idl:"name:SubnetAddress" json:"subnet_address"`
	Resume           uint32                 `idl:"name:ResumeHandle" json:"resume"`
	PreferredMaximum uint32                 `idl:"name:PreferredMaximum" json:"preferred_maximum"`
	ClientInfo       *dhcpm.ClientInfoArray `idl:"name:ClientInfo" json:"client_info"`
	ClientsRead      uint32                 `idl:"name:ClientsRead" json:"clients_read"`
	ClientsTotal     uint32                 `idl:"name:ClientsTotal" json:"clients_total"`
	Return           uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumSubnetClientsOperation) OpNum() int { return 20 }

func (o *xxx_EnumSubnetClientsOperation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpEnumSubnetClients"
}

func (o *xxx_EnumSubnetClientsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO_ARRAY}*(1))(3:{alias=DHCP_CLIENT_INFO_ARRAY}(struct))
	{
		if o.ClientInfo != nil {
			_ptr_ClientInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientInfo != nil {
					if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.ClientInfoArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientInfo, _ptr_ClientInfo); err != nil {
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
	// ClientsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientsRead); err != nil {
			return err
		}
	}
	// ClientsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientsTotal); err != nil {
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

func (o *xxx_EnumSubnetClientsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO_ARRAY,pointer=ref}*(1))(3:{alias=DHCP_CLIENT_INFO_ARRAY}(struct))
	{
		_ptr_ClientInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientInfo == nil {
				o.ClientInfo = &dhcpm.ClientInfoArray{}
			}
			if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ClientInfo := func(ptr interface{}) { o.ClientInfo = *ptr.(**dhcpm.ClientInfoArray) }
		if err := w.ReadPointer(&o.ClientInfo, _s_ClientInfo, _ptr_ClientInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientsRead); err != nil {
			return err
		}
	}
	// ClientsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientsTotal); err != nil {
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

// EnumSubnetClientsRequest structure represents the R_DhcpEnumSubnetClients operation request
type EnumSubnetClientsRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing the
	// IPv4 subnet ID from which DHCPv4 clients are enumerated. If this parameter is set
	// to 0, the DHCPv4 clients from all the IPv4 subnets are returned.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies the enumeration operation. Initially, this value MUST be set to zero,
	// with a successful call returning the handle value used for subsequent enumeration
	// requests. This field contains the last IPv4 address retrieved.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// PreferredMaximum: This is of type DWORD, specifying the preferred maximum number
	// of bytes to return. The minimum value is 1,024 bytes (1 kilobyte), and the maximum
	// value is 65,536 bytes (64 kilobytes). If the input value is greater or less than
	// this range, it MUST be set to the maximum or minimum value, respectively. To retrieve
	// all the DHCPv4 clients serviced on the specific IPv4 subnet, 0xFFFFFFFF is specified.
	PreferredMaximum uint32 `idl:"name:PreferredMaximum" json:"preferred_maximum"`
}

func (o *EnumSubnetClientsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetClientsOperation {
	if o == nil {
		return &xxx_EnumSubnetClientsOperation{}
	}
	return &xxx_EnumSubnetClientsOperation{
		ServerIPAddress:  o.ServerIPAddress,
		SubnetAddress:    o.SubnetAddress,
		Resume:           o.Resume,
		PreferredMaximum: o.PreferredMaximum,
	}
}

func (o *EnumSubnetClientsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetClientsOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.Resume = op.Resume
	o.PreferredMaximum = op.PreferredMaximum
}
func (o *EnumSubnetClientsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumSubnetClientsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetClientsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumSubnetClientsResponse structure represents the R_DhcpEnumSubnetClients operation response
type EnumSubnetClientsResponse struct {
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies the enumeration operation. Initially, this value MUST be set to zero,
	// with a successful call returning the handle value used for subsequent enumeration
	// requests. This field contains the last IPv4 address retrieved.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// ClientInfo: This is a pointer of type LPDHCP_CLIENT_INFO_ARRAY that points to the
	// location which contains the DHCPv4 client lease record array.
	ClientInfo *dhcpm.ClientInfoArray `idl:"name:ClientInfo" json:"client_info"`
	// ClientsRead: This is a pointer to a DWORD that specifies the number of DHCPv4 client
	// lease records read in the ClientInfo parameter. The caller MUST allocate memory for
	// this parameter equal to the size of data type DWORD.
	ClientsRead uint32 `idl:"name:ClientsRead" json:"clients_read"`
	// ClientsTotal: This is a pointer to a DWORD that specifies the number of DHCPv4 client
	// lease records remaining from the current position. For example, if there are 100
	// DHCPv4 lease record clients for an IPv4 subnet, and if 10 DHCPv4 lease records are
	// enumerated per call, then for the first time this would have a value of 90.<31> The
	// caller MUST allocate memory for this parameter equal to the size of data type DWORD.
	ClientsTotal uint32 `idl:"name:ClientsTotal" json:"clients_total"`
	// Return: The R_DhcpEnumSubnetClients return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumSubnetClientsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetClientsOperation {
	if o == nil {
		return &xxx_EnumSubnetClientsOperation{}
	}
	return &xxx_EnumSubnetClientsOperation{
		Resume:       o.Resume,
		ClientInfo:   o.ClientInfo,
		ClientsRead:  o.ClientsRead,
		ClientsTotal: o.ClientsTotal,
		Return:       o.Return,
	}
}

func (o *EnumSubnetClientsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetClientsOperation) {
	if o == nil {
		return
	}
	o.Resume = op.Resume
	o.ClientInfo = op.ClientInfo
	o.ClientsRead = op.ClientsRead
	o.ClientsTotal = op.ClientsTotal
	o.Return = op.Return
}
func (o *EnumSubnetClientsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumSubnetClientsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetClientsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClientOptionsOperation structure represents the R_DhcpGetClientOptions operation
type xxx_GetClientOptionsOperation struct {
	ServerIPAddress  string            `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ClientIPAddress  uint32            `idl:"name:ClientIpAddress" json:"client_ip_address"`
	ClientSubnetMask uint32            `idl:"name:ClientSubnetMask" json:"client_subnet_mask"`
	ClientOptions    *dhcpm.OptionList `idl:"name:ClientOptions" json:"client_options"`
	Return           uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClientOptionsOperation) OpNum() int { return 21 }

func (o *xxx_GetClientOptionsOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetClientOptions" }

func (o *xxx_GetClientOptionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientOptionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ClientIpAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientIPAddress); err != nil {
			return err
		}
	}
	// ClientSubnetMask {in} (1:{alias=DHCP_IP_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientSubnetMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientOptionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientIpAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientIPAddress); err != nil {
			return err
		}
	}
	// ClientSubnetMask {in} (1:{alias=DHCP_IP_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientSubnetMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientOptionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientOptionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ClientOptions {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_OPTION_LIST}*(1))(3:{alias=DHCP_OPTION_LIST}(struct))
	{
		if o.ClientOptions != nil {
			_ptr_ClientOptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientOptions != nil {
					if err := o.ClientOptions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.OptionList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientOptions, _ptr_ClientOptions); err != nil {
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

func (o *xxx_GetClientOptionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ClientOptions {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_OPTION_LIST,pointer=ref}*(1))(3:{alias=DHCP_OPTION_LIST}(struct))
	{
		_ptr_ClientOptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientOptions == nil {
				o.ClientOptions = &dhcpm.OptionList{}
			}
			if err := o.ClientOptions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ClientOptions := func(ptr interface{}) { o.ClientOptions = *ptr.(**dhcpm.OptionList) }
		if err := w.ReadPointer(&o.ClientOptions, _s_ClientOptions, _ptr_ClientOptions); err != nil {
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

// GetClientOptionsRequest structure represents the R_DhcpGetClientOptions operation request
type GetClientOptionsRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ClientIpAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1). Currently it
	// is not used, and any value set to this parameter will not affect the behavior of
	// this method.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// ClientSubnetMask: This is of type DHCP_IP_MASK (section 2.2.1.2.2). Currently it
	// is not used, and any value set to this parameter will not affect the behavior of
	// this method.
	ClientSubnetMask uint32 `idl:"name:ClientSubnetMask" json:"client_subnet_mask"`
}

func (o *GetClientOptionsRequest) xxx_ToOp(ctx context.Context) *xxx_GetClientOptionsOperation {
	if o == nil {
		return &xxx_GetClientOptionsOperation{}
	}
	return &xxx_GetClientOptionsOperation{
		ServerIPAddress:  o.ServerIPAddress,
		ClientIPAddress:  o.ClientIPAddress,
		ClientSubnetMask: o.ClientSubnetMask,
	}
}

func (o *GetClientOptionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClientOptionsOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.ClientIPAddress = op.ClientIPAddress
	o.ClientSubnetMask = op.ClientSubnetMask
}
func (o *GetClientOptionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClientOptionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientOptionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClientOptionsResponse structure represents the R_DhcpGetClientOptions operation response
type GetClientOptionsResponse struct {
	// ClientOptions: This is a pointer of type LPDHCP_OPTION_LIST. Currently it is not
	// used, and any value set to this parameter will not affect the behavior of this method.
	ClientOptions *dhcpm.OptionList `idl:"name:ClientOptions" json:"client_options"`
	// Return: The R_DhcpGetClientOptions return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetClientOptionsResponse) xxx_ToOp(ctx context.Context) *xxx_GetClientOptionsOperation {
	if o == nil {
		return &xxx_GetClientOptionsOperation{}
	}
	return &xxx_GetClientOptionsOperation{
		ClientOptions: o.ClientOptions,
		Return:        o.Return,
	}
}

func (o *GetClientOptionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClientOptionsOperation) {
	if o == nil {
		return
	}
	o.ClientOptions = op.ClientOptions
	o.Return = op.Return
}
func (o *GetClientOptionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClientOptionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientOptionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMIBInfoOperation structure represents the R_DhcpGetMibInfo operation
type xxx_GetMIBInfoOperation struct {
	ServerIPAddress string         `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	MIBInfo         *dhcpm.MIBInfo `idl:"name:MibInfo" json:"mib_info"`
	Return          uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMIBInfoOperation) OpNum() int { return 22 }

func (o *xxx_GetMIBInfoOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetMibInfo" }

func (o *xxx_GetMIBInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMIBInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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

func (o *xxx_GetMIBInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMIBInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMIBInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// MibInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_MIB_INFO}*(1))(3:{alias=DHCP_MIB_INFO}(struct))
	{
		if o.MIBInfo != nil {
			_ptr_MibInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MIBInfo != nil {
					if err := o.MIBInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.MIBInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MIBInfo, _ptr_MibInfo); err != nil {
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

func (o *xxx_GetMIBInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// MibInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_MIB_INFO,pointer=ref}*(1))(3:{alias=DHCP_MIB_INFO}(struct))
	{
		_ptr_MibInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MIBInfo == nil {
				o.MIBInfo = &dhcpm.MIBInfo{}
			}
			if err := o.MIBInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_MibInfo := func(ptr interface{}) { o.MIBInfo = *ptr.(**dhcpm.MIBInfo) }
		if err := w.ReadPointer(&o.MIBInfo, _s_MibInfo, _ptr_MibInfo); err != nil {
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

// GetMIBInfoRequest structure represents the R_DhcpGetMibInfo operation request
type GetMIBInfoRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
}

func (o *GetMIBInfoRequest) xxx_ToOp(ctx context.Context) *xxx_GetMIBInfoOperation {
	if o == nil {
		return &xxx_GetMIBInfoOperation{}
	}
	return &xxx_GetMIBInfoOperation{
		ServerIPAddress: o.ServerIPAddress,
	}
}

func (o *GetMIBInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMIBInfoOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
}
func (o *GetMIBInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMIBInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMIBInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMIBInfoResponse structure represents the R_DhcpGetMibInfo operation response
type GetMIBInfoResponse struct {
	// MibInfo: This is a pointer of type LPDHCP_MIB_INFO that points to the location that
	// contains DHCPv4 server statistics.
	MIBInfo *dhcpm.MIBInfo `idl:"name:MibInfo" json:"mib_info"`
	// Return: The R_DhcpGetMibInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetMIBInfoResponse) xxx_ToOp(ctx context.Context) *xxx_GetMIBInfoOperation {
	if o == nil {
		return &xxx_GetMIBInfoOperation{}
	}
	return &xxx_GetMIBInfoOperation{
		MIBInfo: o.MIBInfo,
		Return:  o.Return,
	}
}

func (o *GetMIBInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMIBInfoOperation) {
	if o == nil {
		return
	}
	o.MIBInfo = op.MIBInfo
	o.Return = op.Return
}
func (o *GetMIBInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMIBInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMIBInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumOptionsOperation structure represents the R_DhcpEnumOptions operation
type xxx_EnumOptionsOperation struct {
	ServerIPAddress  string             `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	Resume           uint32             `idl:"name:ResumeHandle" json:"resume"`
	PreferredMaximum uint32             `idl:"name:PreferredMaximum" json:"preferred_maximum"`
	Options          *dhcpm.OptionArray `idl:"name:Options" json:"options"`
	OptionsRead      uint32             `idl:"name:OptionsRead" json:"options_read"`
	OptionsTotal     uint32             `idl:"name:OptionsTotal" json:"options_total"`
	Return           uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumOptionsOperation) OpNum() int { return 23 }

func (o *xxx_EnumOptionsOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpEnumOptions" }

func (o *xxx_EnumOptionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOptionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOptionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOptionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOptionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// Options {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_OPTION_ARRAY}*(1))(3:{alias=DHCP_OPTION_ARRAY}(struct))
	{
		if o.Options != nil {
			_ptr_Options := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Options != nil {
					if err := o.Options.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.OptionArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Options, _ptr_Options); err != nil {
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
	// OptionsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionsRead); err != nil {
			return err
		}
	}
	// OptionsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OptionsTotal); err != nil {
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

func (o *xxx_EnumOptionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// Options {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_OPTION_ARRAY,pointer=ref}*(1))(3:{alias=DHCP_OPTION_ARRAY}(struct))
	{
		_ptr_Options := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Options == nil {
				o.Options = &dhcpm.OptionArray{}
			}
			if err := o.Options.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Options := func(ptr interface{}) { o.Options = *ptr.(**dhcpm.OptionArray) }
		if err := w.ReadPointer(&o.Options, _s_Options, _ptr_Options); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionsRead); err != nil {
			return err
		}
	}
	// OptionsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OptionsTotal); err != nil {
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

// EnumOptionsRequest structure represents the R_DhcpEnumOptions operation request
type EnumOptionsRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies the enumeration operation. Initially, this value MUST be set to zero,
	// with a successful call returning the handle value used for subsequent enumeration
	// requests. For example, if the PreferredMaximum parameter is set to 1,000 bytes, and
	// 2,000 bytes worth of option definitions are stored on the DHCPv4 server, the resume
	// handle can be used after the first 1,000 bytes are retrieved to obtain the next 1,000
	// on a subsequent call, and so forth.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// PreferredMaximum: This is of type DWORD, specifying the preferred maximum number
	// of bytes to return. If the number of remaining unenumerated option definitions (in
	// bytes) is less than this value, all option definitions are returned. To retrieve
	// all the option definitions defined for a default user class and vendor class on the
	// DHCPv4 server, 0xFFFFFFFF is specified.
	PreferredMaximum uint32 `idl:"name:PreferredMaximum" json:"preferred_maximum"`
}

func (o *EnumOptionsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumOptionsOperation {
	if o == nil {
		return &xxx_EnumOptionsOperation{}
	}
	return &xxx_EnumOptionsOperation{
		ServerIPAddress:  o.ServerIPAddress,
		Resume:           o.Resume,
		PreferredMaximum: o.PreferredMaximum,
	}
}

func (o *EnumOptionsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumOptionsOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.Resume = op.Resume
	o.PreferredMaximum = op.PreferredMaximum
}
func (o *EnumOptionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumOptionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumOptionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumOptionsResponse structure represents the R_DhcpEnumOptions operation response
type EnumOptionsResponse struct {
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies the enumeration operation. Initially, this value MUST be set to zero,
	// with a successful call returning the handle value used for subsequent enumeration
	// requests. For example, if the PreferredMaximum parameter is set to 1,000 bytes, and
	// 2,000 bytes worth of option definitions are stored on the DHCPv4 server, the resume
	// handle can be used after the first 1,000 bytes are retrieved to obtain the next 1,000
	// on a subsequent call, and so forth.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// Options: This is a pointer of type LPDHCP_OPTION_ARRAY that points to the location
	// where all the option definitions for the default user class and vendor class are
	// retrieved from the DHCPv4 server.
	Options *dhcpm.OptionArray `idl:"name:Options" json:"options"`
	// OptionsRead: This is a pointer to a DWORD value that specifies the number of option
	// definitions read in the Options parameter. The caller MUST allocate memory for this
	// parameter equal to the size of data type DWORD.
	OptionsRead uint32 `idl:"name:OptionsRead" json:"options_read"`
	// OptionsTotal: This is a pointer to a DWORD value that specifies the number of option
	// definitions that have not yet been enumerated. The caller MUST allocate memory for
	// this parameter that is equal to the size of data type DWORD.
	OptionsTotal uint32 `idl:"name:OptionsTotal" json:"options_total"`
	// Return: The R_DhcpEnumOptions return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumOptionsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumOptionsOperation {
	if o == nil {
		return &xxx_EnumOptionsOperation{}
	}
	return &xxx_EnumOptionsOperation{
		Resume:       o.Resume,
		Options:      o.Options,
		OptionsRead:  o.OptionsRead,
		OptionsTotal: o.OptionsTotal,
		Return:       o.Return,
	}
}

func (o *EnumOptionsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumOptionsOperation) {
	if o == nil {
		return
	}
	o.Resume = op.Resume
	o.Options = op.Options
	o.OptionsRead = op.OptionsRead
	o.OptionsTotal = op.OptionsTotal
	o.Return = op.Return
}
func (o *EnumOptionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumOptionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumOptionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOptionValuesOperation structure represents the R_DhcpSetOptionValues operation
type xxx_SetOptionValuesOperation struct {
	ServerIPAddress string                  `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ScopeInfo       *dhcpm.OptionScopeInfo  `idl:"name:ScopeInfo;pointer:ref" json:"scope_info"`
	OptionValues    *dhcpm.OptionValueArray `idl:"name:OptionValues;pointer:ref" json:"option_values"`
	Return          uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOptionValuesOperation) OpNum() int { return 24 }

func (o *xxx_SetOptionValuesOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpSetOptionValues" }

func (o *xxx_SetOptionValuesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionValuesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ScopeInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION_SCOPE_INFO}*(1))(2:{alias=DHCP_OPTION_SCOPE_INFO}(struct))
	{
		if o.ScopeInfo != nil {
			if err := o.ScopeInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.OptionScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// OptionValues {in} (1:{pointer=ref, alias=LPDHCP_OPTION_VALUE_ARRAY}*(1))(2:{alias=DHCP_OPTION_VALUE_ARRAY}(struct))
	{
		if o.OptionValues != nil {
			if err := o.OptionValues.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.OptionValueArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionValuesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ScopeInfo {in} (1:{pointer=ref, alias=LPDHCP_OPTION_SCOPE_INFO}*(1))(2:{alias=DHCP_OPTION_SCOPE_INFO}(struct))
	{
		if o.ScopeInfo == nil {
			o.ScopeInfo = &dhcpm.OptionScopeInfo{}
		}
		if err := o.ScopeInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OptionValues {in} (1:{pointer=ref, alias=LPDHCP_OPTION_VALUE_ARRAY}*(1))(2:{alias=DHCP_OPTION_VALUE_ARRAY}(struct))
	{
		if o.OptionValues == nil {
			o.OptionValues = &dhcpm.OptionValueArray{}
		}
		if err := o.OptionValues.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionValuesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOptionValuesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOptionValuesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetOptionValuesRequest structure represents the R_DhcpSetOptionValues operation request
type SetOptionValuesRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ScopeInfo: This is a pointer to a DHCP_OPTION_SCOPE_INFO (section 2.2.1.2.41) structure
	// that contains the level (that is, default, server, scope, multicast scope, or IPv4
	// reservation level) at which the option values are set.
	ScopeInfo *dhcpm.OptionScopeInfo `idl:"name:ScopeInfo;pointer:ref" json:"scope_info"`
	// OptionValues: This is a pointer of type DHCP_OPTION_VALUE_ARRAY (section 2.2.1.2.43)
	// structure that points to the location that contains one or more option identifiers
	// along with the values.
	OptionValues *dhcpm.OptionValueArray `idl:"name:OptionValues;pointer:ref" json:"option_values"`
}

func (o *SetOptionValuesRequest) xxx_ToOp(ctx context.Context) *xxx_SetOptionValuesOperation {
	if o == nil {
		return &xxx_SetOptionValuesOperation{}
	}
	return &xxx_SetOptionValuesOperation{
		ServerIPAddress: o.ServerIPAddress,
		ScopeInfo:       o.ScopeInfo,
		OptionValues:    o.OptionValues,
	}
}

func (o *SetOptionValuesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOptionValuesOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.ScopeInfo = op.ScopeInfo
	o.OptionValues = op.OptionValues
}
func (o *SetOptionValuesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetOptionValuesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOptionValuesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOptionValuesResponse structure represents the R_DhcpSetOptionValues operation response
type SetOptionValuesResponse struct {
	// Return: The R_DhcpSetOptionValues return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetOptionValuesResponse) xxx_ToOp(ctx context.Context) *xxx_SetOptionValuesOperation {
	if o == nil {
		return &xxx_SetOptionValuesOperation{}
	}
	return &xxx_SetOptionValuesOperation{
		Return: o.Return,
	}
}

func (o *SetOptionValuesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOptionValuesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetOptionValuesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetOptionValuesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOptionValuesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerSetConfigOperation structure represents the R_DhcpServerSetConfig operation
type xxx_ServerSetConfigOperation struct {
	ServerIPAddress string                  `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	FieldsToSet     uint32                  `idl:"name:FieldsToSet" json:"fields_to_set"`
	ConfigInfo      *dhcpm.ServerConfigInfo `idl:"name:ConfigInfo;pointer:ref" json:"config_info"`
	Return          uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerSetConfigOperation) OpNum() int { return 25 }

func (o *xxx_ServerSetConfigOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpServerSetConfig" }

func (o *xxx_ServerSetConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// FieldsToSet {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.FieldsToSet); err != nil {
			return err
		}
	}
	// ConfigInfo {in} (1:{pointer=ref, alias=LPDHCP_SERVER_CONFIG_INFO}*(1))(2:{alias=DHCP_SERVER_CONFIG_INFO}(struct))
	{
		if o.ConfigInfo != nil {
			if err := o.ConfigInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.ServerConfigInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// FieldsToSet {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.FieldsToSet); err != nil {
			return err
		}
	}
	// ConfigInfo {in} (1:{pointer=ref, alias=LPDHCP_SERVER_CONFIG_INFO}*(1))(2:{alias=DHCP_SERVER_CONFIG_INFO}(struct))
	{
		if o.ConfigInfo == nil {
			o.ConfigInfo = &dhcpm.ServerConfigInfo{}
		}
		if err := o.ConfigInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ServerSetConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerSetConfigRequest structure represents the R_DhcpServerSetConfig operation request
type ServerSetConfigRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// FieldsToSet: A DWORD that contains the bitmask of the fields in the ConfigInfo structure
	// to set. This method can be called with a value for FieldsToSet.
	//
	// The bit mapping for the various values for FieldsToSet is listed in the following
	// table:
	//
	//	+-----------------------------+------------+
	//	|                             |            |
	//	|         FIELDSTOSET         |    BIT     |
	//	|                             |            |
	//	+-----------------------------+------------+
	//	+-----------------------------+------------+
	//	| Set_APIProtocolSupport      | 0x00000001 |
	//	+-----------------------------+------------+
	//	| Set_DatabaseName            | 0x00000002 |
	//	+-----------------------------+------------+
	//	| Set_DatabasePath            | 0x00000004 |
	//	+-----------------------------+------------+
	//	| Set_BackupPath              | 0x00000008 |
	//	+-----------------------------+------------+
	//	| Set_BackupInterval          | 0x00000010 |
	//	+-----------------------------+------------+
	//	| Set_DatabaseLoggingFlag     | 0x00000020 |
	//	+-----------------------------+------------+
	//	| Set_RestoreFlag             | 0x00000040 |
	//	+-----------------------------+------------+
	//	| Set_DatabaseCleanupInterval | 0x00000080 |
	//	+-----------------------------+------------+
	//	| Set_DebugFlag               | 0x00000100 |
	//	+-----------------------------+------------+
	//
	// The DHCP Server ignores the bits not specified in the table.
	//
	// * Set_APIProtocolSupport
	//
	// * Set_DatabaseName
	//
	// * Set_DatabasePath
	//
	// * Set_DatabaseLoggingFlag
	//
	// * Set_RestoreFlag
	FieldsToSet uint32 `idl:"name:FieldsToSet" json:"fields_to_set"`
	// ConfigInfo: A pointer of type DHCP_SERVER_CONFIG_INFO (section 2.2.1.2.53) structure
	// that contains the settings for the DHCPv4 server. The value that is passed here depends
	// on the FieldsToSet parameter. Details of the dependencies follow the return value
	// description.
	ConfigInfo *dhcpm.ServerConfigInfo `idl:"name:ConfigInfo;pointer:ref" json:"config_info"`
}

func (o *ServerSetConfigRequest) xxx_ToOp(ctx context.Context) *xxx_ServerSetConfigOperation {
	if o == nil {
		return &xxx_ServerSetConfigOperation{}
	}
	return &xxx_ServerSetConfigOperation{
		ServerIPAddress: o.ServerIPAddress,
		FieldsToSet:     o.FieldsToSet,
		ConfigInfo:      o.ConfigInfo,
	}
}

func (o *ServerSetConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerSetConfigOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.FieldsToSet = op.FieldsToSet
	o.ConfigInfo = op.ConfigInfo
}
func (o *ServerSetConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ServerSetConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerSetConfigResponse structure represents the R_DhcpServerSetConfig operation response
type ServerSetConfigResponse struct {
	// Return: The R_DhcpServerSetConfig return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerSetConfigResponse) xxx_ToOp(ctx context.Context) *xxx_ServerSetConfigOperation {
	if o == nil {
		return &xxx_ServerSetConfigOperation{}
	}
	return &xxx_ServerSetConfigOperation{
		Return: o.Return,
	}
}

func (o *ServerSetConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerSetConfigOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ServerSetConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ServerSetConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerGetConfigOperation structure represents the R_DhcpServerGetConfig operation
type xxx_ServerGetConfigOperation struct {
	ServerIPAddress string                  `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ConfigInfo      *dhcpm.ServerConfigInfo `idl:"name:ConfigInfo" json:"config_info"`
	Return          uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerGetConfigOperation) OpNum() int { return 26 }

func (o *xxx_ServerGetConfigOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpServerGetConfig" }

func (o *xxx_ServerGetConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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

func (o *xxx_ServerGetConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ConfigInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SERVER_CONFIG_INFO}*(1))(3:{alias=DHCP_SERVER_CONFIG_INFO}(struct))
	{
		if o.ConfigInfo != nil {
			_ptr_ConfigInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigInfo != nil {
					if err := o.ConfigInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.ServerConfigInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigInfo, _ptr_ConfigInfo); err != nil {
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

func (o *xxx_ServerGetConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ConfigInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SERVER_CONFIG_INFO,pointer=ref}*(1))(3:{alias=DHCP_SERVER_CONFIG_INFO}(struct))
	{
		_ptr_ConfigInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigInfo == nil {
				o.ConfigInfo = &dhcpm.ServerConfigInfo{}
			}
			if err := o.ConfigInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ConfigInfo := func(ptr interface{}) { o.ConfigInfo = *ptr.(**dhcpm.ServerConfigInfo) }
		if err := w.ReadPointer(&o.ConfigInfo, _s_ConfigInfo, _ptr_ConfigInfo); err != nil {
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

// ServerGetConfigRequest structure represents the R_DhcpServerGetConfig operation request
type ServerGetConfigRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
}

func (o *ServerGetConfigRequest) xxx_ToOp(ctx context.Context) *xxx_ServerGetConfigOperation {
	if o == nil {
		return &xxx_ServerGetConfigOperation{}
	}
	return &xxx_ServerGetConfigOperation{
		ServerIPAddress: o.ServerIPAddress,
	}
}

func (o *ServerGetConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerGetConfigOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
}
func (o *ServerGetConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ServerGetConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerGetConfigResponse structure represents the R_DhcpServerGetConfig operation response
type ServerGetConfigResponse struct {
	// ConfigInfo: This is a pointer of type LPDHCP_SERVER_CONFIG_INFO that points to the
	// location where the DHCPv4 server settings are retrieved. The caller of this method
	// can free up this structure after use.
	ConfigInfo *dhcpm.ServerConfigInfo `idl:"name:ConfigInfo" json:"config_info"`
	// Return: The R_DhcpServerGetConfig return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerGetConfigResponse) xxx_ToOp(ctx context.Context) *xxx_ServerGetConfigOperation {
	if o == nil {
		return &xxx_ServerGetConfigOperation{}
	}
	return &xxx_ServerGetConfigOperation{
		ConfigInfo: o.ConfigInfo,
		Return:     o.Return,
	}
}

func (o *ServerGetConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerGetConfigOperation) {
	if o == nil {
		return
	}
	o.ConfigInfo = op.ConfigInfo
	o.Return = op.Return
}
func (o *ServerGetConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ServerGetConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ScanDatabaseOperation structure represents the R_DhcpScanDatabase operation
type xxx_ScanDatabaseOperation struct {
	ServerIPAddress string          `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32          `idl:"name:SubnetAddress" json:"subnet_address"`
	FixFlag         uint32          `idl:"name:FixFlag" json:"fix_flag"`
	ScanList        *dhcpm.ScanList `idl:"name:ScanList" json:"scan_list"`
	Return          uint32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ScanDatabaseOperation) OpNum() int { return 27 }

func (o *xxx_ScanDatabaseOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpScanDatabase" }

func (o *xxx_ScanDatabaseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScanDatabaseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// FixFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.FixFlag); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScanDatabaseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// FixFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.FixFlag); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScanDatabaseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ScanDatabaseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ScanList {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SCAN_LIST}*(1))(3:{alias=DHCP_SCAN_LIST}(struct))
	{
		if o.ScanList != nil {
			_ptr_ScanList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ScanList != nil {
					if err := o.ScanList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.ScanList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ScanList, _ptr_ScanList); err != nil {
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

func (o *xxx_ScanDatabaseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ScanList {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SCAN_LIST,pointer=ref}*(1))(3:{alias=DHCP_SCAN_LIST}(struct))
	{
		_ptr_ScanList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ScanList == nil {
				o.ScanList = &dhcpm.ScanList{}
			}
			if err := o.ScanList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ScanList := func(ptr interface{}) { o.ScanList = *ptr.(**dhcpm.ScanList) }
		if err := w.ReadPointer(&o.ScanList, _s_ScanList, _ptr_ScanList); err != nil {
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

// ScanDatabaseRequest structure represents the R_DhcpScanDatabase operation request
type ScanDatabaseRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing the
	// IPv4 subnet ID of the subnet in which the scan is done for the IPv4 addresses that
	// are not in sync.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// FixFlag: This is of type DWORD, defining the behavior of this method. This method
	// enumerates the DHCPv4 client IPv4 addresses that are not in sync in both the stores,
	// and if the FixFlag parameter is set to TRUE, it fixes those unmatched IPv4 addresses
	// also.
	FixFlag uint32 `idl:"name:FixFlag" json:"fix_flag"`
}

func (o *ScanDatabaseRequest) xxx_ToOp(ctx context.Context) *xxx_ScanDatabaseOperation {
	if o == nil {
		return &xxx_ScanDatabaseOperation{}
	}
	return &xxx_ScanDatabaseOperation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
		FixFlag:         o.FixFlag,
	}
}

func (o *ScanDatabaseRequest) xxx_FromOp(ctx context.Context, op *xxx_ScanDatabaseOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.FixFlag = op.FixFlag
}
func (o *ScanDatabaseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ScanDatabaseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ScanDatabaseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ScanDatabaseResponse structure represents the R_DhcpScanDatabase operation response
type ScanDatabaseResponse struct {
	// ScanList: This is a pointer of type LPDHCP_SCAN_LIST that points to the location
	// containing the DHCPv4 client IPv4 addresses that are not in sync in both the stores.
	ScanList *dhcpm.ScanList `idl:"name:ScanList" json:"scan_list"`
	// Return: The R_DhcpScanDatabase return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ScanDatabaseResponse) xxx_ToOp(ctx context.Context) *xxx_ScanDatabaseOperation {
	if o == nil {
		return &xxx_ScanDatabaseOperation{}
	}
	return &xxx_ScanDatabaseOperation{
		ScanList: o.ScanList,
		Return:   o.Return,
	}
}

func (o *ScanDatabaseResponse) xxx_FromOp(ctx context.Context, op *xxx_ScanDatabaseOperation) {
	if o == nil {
		return
	}
	o.ScanList = op.ScanList
	o.Return = op.Return
}
func (o *ScanDatabaseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ScanDatabaseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ScanDatabaseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVersionOperation structure represents the R_DhcpGetVersion operation
type xxx_GetVersionOperation struct {
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	MajorVersion    uint32 `idl:"name:MajorVersion" json:"major_version"`
	MinorVersion    uint32 `idl:"name:MinorVersion" json:"minor_version"`
	Return          uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVersionOperation) OpNum() int { return 28 }

func (o *xxx_GetVersionOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetVersion" }

func (o *xxx_GetVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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

func (o *xxx_GetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// MajorVersion {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
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

func (o *xxx_GetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// MajorVersion {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// MinorVersion {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
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

// GetVersionRequest structure represents the R_DhcpGetVersion operation request
type GetVersionRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
}

func (o *GetVersionRequest) xxx_ToOp(ctx context.Context) *xxx_GetVersionOperation {
	if o == nil {
		return &xxx_GetVersionOperation{}
	}
	return &xxx_GetVersionOperation{
		ServerIPAddress: o.ServerIPAddress,
	}
}

func (o *GetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
}
func (o *GetVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVersionResponse structure represents the R_DhcpGetVersion operation response
type GetVersionResponse struct {
	// MajorVersion: This is a pointer to a DWORD in which the major version of the DHCP
	// server is returned. The MajorVersion parameter MUST be allocated by the client before
	// the call.
	MajorVersion uint32 `idl:"name:MajorVersion" json:"major_version"`
	// MinorVersion: This is a pointer to a DWORD in which the minor version of the DHCP
	// server is returned. The MinorVersion parameter MUST be allocated by the client before
	// the call.
	MinorVersion uint32 `idl:"name:MinorVersion" json:"minor_version"`
	// Return: The R_DhcpGetVersion return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetVersionResponse) xxx_ToOp(ctx context.Context) *xxx_GetVersionOperation {
	if o == nil {
		return &xxx_GetVersionOperation{}
	}
	return &xxx_GetVersionOperation{
		MajorVersion: o.MajorVersion,
		MinorVersion: o.MinorVersion,
		Return:       o.Return,
	}
}

func (o *GetVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
	o.Return = op.Return
}
func (o *GetVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddSubnetElementV4Operation structure represents the R_DhcpAddSubnetElementV4 operation
type xxx_AddSubnetElementV4Operation struct {
	ServerIPAddress string                     `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32                     `idl:"name:SubnetAddress" json:"subnet_address"`
	AddElementInfo  *dhcpm.SubnetElementDataV4 `idl:"name:AddElementInfo;pointer:ref" json:"add_element_info"`
	Return          uint32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_AddSubnetElementV4Operation) OpNum() int { return 29 }

func (o *xxx_AddSubnetElementV4Operation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpAddSubnetElementV4"
}

func (o *xxx_AddSubnetElementV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddSubnetElementV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// AddElementInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_ELEMENT_DATA_V4}*(1))(2:{alias=DHCP_SUBNET_ELEMENT_DATA_V4}(struct))
	{
		if o.AddElementInfo != nil {
			if err := o.AddElementInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SubnetElementDataV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddSubnetElementV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// AddElementInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_ELEMENT_DATA_V4}*(1))(2:{alias=DHCP_SUBNET_ELEMENT_DATA_V4}(struct))
	{
		if o.AddElementInfo == nil {
			o.AddElementInfo = &dhcpm.SubnetElementDataV4{}
		}
		if err := o.AddElementInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddSubnetElementV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddSubnetElementV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddSubnetElementV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddSubnetElementV4Request structure represents the R_DhcpAddSubnetElementV4 operation request
type AddSubnetElementV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS, containing the IPv4 subnet ID in
	// which the IPv4 subnet element is added.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// AddElementInfo: This is a pointer to a DHCP_SUBNET_ELEMENT_DATA_V4 (section 2.2.1.2.35)
	// structure that contains the IPv4 subnet element that needs to be added to the IPv4
	// subnet.
	AddElementInfo *dhcpm.SubnetElementDataV4 `idl:"name:AddElementInfo;pointer:ref" json:"add_element_info"`
}

func (o *AddSubnetElementV4Request) xxx_ToOp(ctx context.Context) *xxx_AddSubnetElementV4Operation {
	if o == nil {
		return &xxx_AddSubnetElementV4Operation{}
	}
	return &xxx_AddSubnetElementV4Operation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
		AddElementInfo:  o.AddElementInfo,
	}
}

func (o *AddSubnetElementV4Request) xxx_FromOp(ctx context.Context, op *xxx_AddSubnetElementV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.AddElementInfo = op.AddElementInfo
}
func (o *AddSubnetElementV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddSubnetElementV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddSubnetElementV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddSubnetElementV4Response structure represents the R_DhcpAddSubnetElementV4 operation response
type AddSubnetElementV4Response struct {
	// Return: The R_DhcpAddSubnetElementV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddSubnetElementV4Response) xxx_ToOp(ctx context.Context) *xxx_AddSubnetElementV4Operation {
	if o == nil {
		return &xxx_AddSubnetElementV4Operation{}
	}
	return &xxx_AddSubnetElementV4Operation{
		Return: o.Return,
	}
}

func (o *AddSubnetElementV4Response) xxx_FromOp(ctx context.Context, op *xxx_AddSubnetElementV4Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddSubnetElementV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddSubnetElementV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddSubnetElementV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumSubnetElementsV4Operation structure represents the R_DhcpEnumSubnetElementsV4 operation
type xxx_EnumSubnetElementsV4Operation struct {
	ServerIPAddress  string                          `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress    uint32                          `idl:"name:SubnetAddress" json:"subnet_address"`
	EnumElementType  dhcpm.SubnetElementType         `idl:"name:EnumElementType" json:"enum_element_type"`
	Resume           uint32                          `idl:"name:ResumeHandle" json:"resume"`
	PreferredMaximum uint32                          `idl:"name:PreferredMaximum" json:"preferred_maximum"`
	EnumElementInfo  *dhcpm.SubnetElementInfoArrayV4 `idl:"name:EnumElementInfo" json:"enum_element_info"`
	ElementsRead     uint32                          `idl:"name:ElementsRead" json:"elements_read"`
	ElementsTotal    uint32                          `idl:"name:ElementsTotal" json:"elements_total"`
	Return           uint32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumSubnetElementsV4Operation) OpNum() int { return 30 }

func (o *xxx_EnumSubnetElementsV4Operation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpEnumSubnetElementsV4"
}

func (o *xxx_EnumSubnetElementsV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetElementsV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// EnumElementType {in} (1:{alias=DHCP_SUBNET_ELEMENT_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.EnumElementType)); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetElementsV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// EnumElementType {in} (1:{alias=DHCP_SUBNET_ELEMENT_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.EnumElementType)); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetElementsV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetElementsV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// EnumElementInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SUBNET_ELEMENT_INFO_ARRAY_V4}*(1))(3:{alias=DHCP_SUBNET_ELEMENT_INFO_ARRAY_V4}(struct))
	{
		if o.EnumElementInfo != nil {
			_ptr_EnumElementInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EnumElementInfo != nil {
					if err := o.EnumElementInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.SubnetElementInfoArrayV4{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EnumElementInfo, _ptr_EnumElementInfo); err != nil {
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
	// ElementsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ElementsRead); err != nil {
			return err
		}
	}
	// ElementsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ElementsTotal); err != nil {
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

func (o *xxx_EnumSubnetElementsV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// EnumElementInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SUBNET_ELEMENT_INFO_ARRAY_V4,pointer=ref}*(1))(3:{alias=DHCP_SUBNET_ELEMENT_INFO_ARRAY_V4}(struct))
	{
		_ptr_EnumElementInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EnumElementInfo == nil {
				o.EnumElementInfo = &dhcpm.SubnetElementInfoArrayV4{}
			}
			if err := o.EnumElementInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EnumElementInfo := func(ptr interface{}) { o.EnumElementInfo = *ptr.(**dhcpm.SubnetElementInfoArrayV4) }
		if err := w.ReadPointer(&o.EnumElementInfo, _s_EnumElementInfo, _ptr_EnumElementInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ElementsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ElementsRead); err != nil {
			return err
		}
	}
	// ElementsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ElementsTotal); err != nil {
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

// EnumSubnetElementsV4Request structure represents the R_DhcpEnumSubnetElementsV4 operation request
type EnumSubnetElementsV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS, containing the IPv4 subnet ID from
	// which subnet elements are enumerated.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// EnumElementType: This is of type DHCP_SUBNET_ELEMENT_TYPE (section 2.2.1.1.7) enumeration
	// value, indicating the type of IPv4 subnet element to enumerate.
	EnumElementType dhcpm.SubnetElementType `idl:"name:EnumElementType" json:"enum_element_type"`
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies the enumeration operation. Initially, this value MUST be set to zero,
	// with a successful call returning the handle value used for subsequent enumeration
	// requests. For example, if the PreferredMaximum parameter is set to 1,000 bytes, and
	// 2,000 bytes worth of IPv4 subnet elements are stored on the DHCPv4 server, the resume
	// handle can be used after the first 1,000 bytes are retrieved to obtain the next 1,000
	// on a subsequent call, and so forth.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// PreferredMaximum: This is of type DWORD, specifying the preferred maximum number
	// of bytes to return. If the number of remaining un-enumerated subnet elements (in
	// bytes) is less than this value, all IPv4 subnet elements for the specific type are
	// returned. To retrieve all the IPv4 subnet elements of a specific type, 0xFFFFFFFF
	// is specified.
	PreferredMaximum uint32 `idl:"name:PreferredMaximum" json:"preferred_maximum"`
}

func (o *EnumSubnetElementsV4Request) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetElementsV4Operation {
	if o == nil {
		return &xxx_EnumSubnetElementsV4Operation{}
	}
	return &xxx_EnumSubnetElementsV4Operation{
		ServerIPAddress:  o.ServerIPAddress,
		SubnetAddress:    o.SubnetAddress,
		EnumElementType:  o.EnumElementType,
		Resume:           o.Resume,
		PreferredMaximum: o.PreferredMaximum,
	}
}

func (o *EnumSubnetElementsV4Request) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetElementsV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.EnumElementType = op.EnumElementType
	o.Resume = op.Resume
	o.PreferredMaximum = op.PreferredMaximum
}
func (o *EnumSubnetElementsV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumSubnetElementsV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetElementsV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumSubnetElementsV4Response structure represents the R_DhcpEnumSubnetElementsV4 operation response
type EnumSubnetElementsV4Response struct {
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies the enumeration operation. Initially, this value MUST be set to zero,
	// with a successful call returning the handle value used for subsequent enumeration
	// requests. For example, if the PreferredMaximum parameter is set to 1,000 bytes, and
	// 2,000 bytes worth of IPv4 subnet elements are stored on the DHCPv4 server, the resume
	// handle can be used after the first 1,000 bytes are retrieved to obtain the next 1,000
	// on a subsequent call, and so forth.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// EnumElementInfo: This is a pointer of type LPDHCP_SUBNET_ELEMENT_INFO_ARRAY_V4 in
	// which an IPv4 subnet element of type EnumElementType is returned for a specific IPv4
	// subnet SubnetAddress. If no IPv4 subnet element of the specific type is available
	// for enumeration, this value is null. The caller is responsible for freeing this memory.
	EnumElementInfo *dhcpm.SubnetElementInfoArrayV4 `idl:"name:EnumElementInfo" json:"enum_element_info"`
	// ElementsRead: This is a pointer to a DWORD value that specifies the number of IPv4
	// subnet elements read in the EnumElementInfo parameter for a specific type of IPv4
	// subnet element. The caller MUST allocate memory for this parameter equal to the size
	// of DWORD data type.
	ElementsRead uint32 `idl:"name:ElementsRead" json:"elements_read"`
	// ElementsTotal: This is a pointer to a DWORD value that specifies the number of IPv4
	// subnet elements of a specific type from a specific IPv4 subnet that are not yet enumerated
	// with respect to the resume handle that is returned. The caller MUST allocate memory
	// for this parameter equal to the size of DWORD data type.
	ElementsTotal uint32 `idl:"name:ElementsTotal" json:"elements_total"`
	// Return: The R_DhcpEnumSubnetElementsV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumSubnetElementsV4Response) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetElementsV4Operation {
	if o == nil {
		return &xxx_EnumSubnetElementsV4Operation{}
	}
	return &xxx_EnumSubnetElementsV4Operation{
		Resume:          o.Resume,
		EnumElementInfo: o.EnumElementInfo,
		ElementsRead:    o.ElementsRead,
		ElementsTotal:   o.ElementsTotal,
		Return:          o.Return,
	}
}

func (o *EnumSubnetElementsV4Response) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetElementsV4Operation) {
	if o == nil {
		return
	}
	o.Resume = op.Resume
	o.EnumElementInfo = op.EnumElementInfo
	o.ElementsRead = op.ElementsRead
	o.ElementsTotal = op.ElementsTotal
	o.Return = op.Return
}
func (o *EnumSubnetElementsV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumSubnetElementsV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetElementsV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveSubnetElementV4Operation structure represents the R_DhcpRemoveSubnetElementV4 operation
type xxx_RemoveSubnetElementV4Operation struct {
	ServerIPAddress   string                     `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress     uint32                     `idl:"name:SubnetAddress" json:"subnet_address"`
	RemoveElementInfo *dhcpm.SubnetElementDataV4 `idl:"name:RemoveElementInfo;pointer:ref" json:"remove_element_info"`
	ForceFlag         dhcpm.ForceFlag            `idl:"name:ForceFlag" json:"force_flag"`
	Return            uint32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveSubnetElementV4Operation) OpNum() int { return 31 }

func (o *xxx_RemoveSubnetElementV4Operation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpRemoveSubnetElementV4"
}

func (o *xxx_RemoveSubnetElementV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubnetElementV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// RemoveElementInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_ELEMENT_DATA_V4}*(1))(2:{alias=DHCP_SUBNET_ELEMENT_DATA_V4}(struct))
	{
		if o.RemoveElementInfo != nil {
			if err := o.RemoveElementInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SubnetElementDataV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ForceFlag {in} (1:{alias=DHCP_FORCE_FLAG}(enum))
	{
		if err := w.WriteEnum(uint16(o.ForceFlag)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubnetElementV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// RemoveElementInfo {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_ELEMENT_DATA_V4}*(1))(2:{alias=DHCP_SUBNET_ELEMENT_DATA_V4}(struct))
	{
		if o.RemoveElementInfo == nil {
			o.RemoveElementInfo = &dhcpm.SubnetElementDataV4{}
		}
		if err := o.RemoveElementInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ForceFlag {in} (1:{alias=DHCP_FORCE_FLAG}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ForceFlag)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubnetElementV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveSubnetElementV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveSubnetElementV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveSubnetElementV4Request structure represents the R_DhcpRemoveSubnetElementV4 operation request
type RemoveSubnetElementV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1) that contains
	// the IPv4 subnet ID from which the IPv4 subnet element is removed.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// RemoveElementInfo: This is a pointer of type DHCP_SUBNET_ELEMENT_DATA_V4 (section
	// 2.2.1.2.35) structure that contains the IPv4 subnet element that needs to be removed
	// from the IPv4 subnet.
	RemoveElementInfo *dhcpm.SubnetElementDataV4 `idl:"name:RemoveElementInfo;pointer:ref" json:"remove_element_info"`
	// ForceFlag: This is of type DHCP_FORCE_FLAG (section 2.2.1.1.9) enumeration, defining
	// the behavior of this method. If the flag is set to DhcpNoForce value and this subnet
	// has served the IPv4 address to some DHCPv4\BOOTP clients, the IPv4 range is not deleted.
	// If the flag is set to DhcpFullForce value, the IPv4 range is deleted along with the
	// DHCPv4 client lease record on the DHCPv4 server.
	ForceFlag dhcpm.ForceFlag `idl:"name:ForceFlag" json:"force_flag"`
}

func (o *RemoveSubnetElementV4Request) xxx_ToOp(ctx context.Context) *xxx_RemoveSubnetElementV4Operation {
	if o == nil {
		return &xxx_RemoveSubnetElementV4Operation{}
	}
	return &xxx_RemoveSubnetElementV4Operation{
		ServerIPAddress:   o.ServerIPAddress,
		SubnetAddress:     o.SubnetAddress,
		RemoveElementInfo: o.RemoveElementInfo,
		ForceFlag:         o.ForceFlag,
	}
}

func (o *RemoveSubnetElementV4Request) xxx_FromOp(ctx context.Context, op *xxx_RemoveSubnetElementV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.RemoveElementInfo = op.RemoveElementInfo
	o.ForceFlag = op.ForceFlag
}
func (o *RemoveSubnetElementV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveSubnetElementV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveSubnetElementV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveSubnetElementV4Response structure represents the R_DhcpRemoveSubnetElementV4 operation response
type RemoveSubnetElementV4Response struct {
	// Return: The R_DhcpRemoveSubnetElementV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveSubnetElementV4Response) xxx_ToOp(ctx context.Context) *xxx_RemoveSubnetElementV4Operation {
	if o == nil {
		return &xxx_RemoveSubnetElementV4Operation{}
	}
	return &xxx_RemoveSubnetElementV4Operation{
		Return: o.Return,
	}
}

func (o *RemoveSubnetElementV4Response) xxx_FromOp(ctx context.Context, op *xxx_RemoveSubnetElementV4Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveSubnetElementV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveSubnetElementV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveSubnetElementV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateClientInfoV4Operation structure represents the R_DhcpCreateClientInfoV4 operation
type xxx_CreateClientInfoV4Operation struct {
	ServerIPAddress string              `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ClientInfo      *dhcpm.ClientInfoV4 `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
	Return          uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateClientInfoV4Operation) OpNum() int { return 32 }

func (o *xxx_CreateClientInfoV4Operation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpCreateClientInfoV4"
}

func (o *xxx_CreateClientInfoV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO_V4}*(1))(2:{alias=DHCP_CLIENT_INFO_V4}(struct))
	{
		if o.ClientInfo != nil {
			if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.ClientInfoV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO_V4}*(1))(2:{alias=DHCP_CLIENT_INFO_V4}(struct))
	{
		if o.ClientInfo == nil {
			o.ClientInfo = &dhcpm.ClientInfoV4{}
		}
		if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateClientInfoV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateClientInfoV4Request structure represents the R_DhcpCreateClientInfoV4 operation request
type CreateClientInfoV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ClientInfo: A pointer of type DHCP_CLIENT_INFO_V4 (section 2.2.1.2.14) structure
	// that contains the DHCPv4 client lease record information that needs to be set on
	// the DHCPv4 server. The caller MUST pass the ClientIPAddress and ClientHardwareAddress
	// member fields when adding a DHCPv4 client lease record to the DHCPv4 server database.
	// The ClientHardwareAddress member represents a DHCPv4 client-identifier (section 2.2.1.2.5.1).
	// Members ClientName, ClientComment, ClientLeaseExpires, and OwnerHost are modified
	// on the DHCPv4 client lease record identified by the ClientIPAddress member.
	ClientInfo *dhcpm.ClientInfoV4 `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
}

func (o *CreateClientInfoV4Request) xxx_ToOp(ctx context.Context) *xxx_CreateClientInfoV4Operation {
	if o == nil {
		return &xxx_CreateClientInfoV4Operation{}
	}
	return &xxx_CreateClientInfoV4Operation{
		ServerIPAddress: o.ServerIPAddress,
		ClientInfo:      o.ClientInfo,
	}
}

func (o *CreateClientInfoV4Request) xxx_FromOp(ctx context.Context, op *xxx_CreateClientInfoV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.ClientInfo = op.ClientInfo
}
func (o *CreateClientInfoV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateClientInfoV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateClientInfoV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateClientInfoV4Response structure represents the R_DhcpCreateClientInfoV4 operation response
type CreateClientInfoV4Response struct {
	// Return: The R_DhcpCreateClientInfoV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateClientInfoV4Response) xxx_ToOp(ctx context.Context) *xxx_CreateClientInfoV4Operation {
	if o == nil {
		return &xxx_CreateClientInfoV4Operation{}
	}
	return &xxx_CreateClientInfoV4Operation{
		Return: o.Return,
	}
}

func (o *CreateClientInfoV4Response) xxx_FromOp(ctx context.Context, op *xxx_CreateClientInfoV4Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CreateClientInfoV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateClientInfoV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateClientInfoV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClientInfoV4Operation structure represents the R_DhcpSetClientInfoV4 operation
type xxx_SetClientInfoV4Operation struct {
	ServerIPAddress string              `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ClientInfo      *dhcpm.ClientInfoV4 `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
	Return          uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClientInfoV4Operation) OpNum() int { return 33 }

func (o *xxx_SetClientInfoV4Operation) OpName() string { return "/dhcpsrv/v1/R_DhcpSetClientInfoV4" }

func (o *xxx_SetClientInfoV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO_V4}*(1))(2:{alias=DHCP_CLIENT_INFO_V4}(struct))
	{
		if o.ClientInfo != nil {
			if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.ClientInfoV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO_V4}*(1))(2:{alias=DHCP_CLIENT_INFO_V4}(struct))
	{
		if o.ClientInfo == nil {
			o.ClientInfo = &dhcpm.ClientInfoV4{}
		}
		if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetClientInfoV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetClientInfoV4Request structure represents the R_DhcpSetClientInfoV4 operation request
type SetClientInfoV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ClientInfo: A pointer of type DHCP_CLIENT_INFO_V4 (section 2.2.1.2.14) structure
	// that contains the DHCPv4 client lease record information that needs to be modified
	// on the DHCPv4 server database. The caller MUST pass the ClientIPAddress and ClientHardwareAddress
	// member fields when modifying a DHCPv4 client lease record in the DHCPv4 server database.
	// The ClientHardwareAddress member represents a DHCPv4 client-identifier (section 2.2.1.2.5.1).
	// The members ClientName, ClientComment, ClientLeaseExpires, and OwnerHost are modified
	// in the DHCPv4 client lease record identified by the ClientIPaddress member.
	ClientInfo *dhcpm.ClientInfoV4 `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
}

func (o *SetClientInfoV4Request) xxx_ToOp(ctx context.Context) *xxx_SetClientInfoV4Operation {
	if o == nil {
		return &xxx_SetClientInfoV4Operation{}
	}
	return &xxx_SetClientInfoV4Operation{
		ServerIPAddress: o.ServerIPAddress,
		ClientInfo:      o.ClientInfo,
	}
}

func (o *SetClientInfoV4Request) xxx_FromOp(ctx context.Context, op *xxx_SetClientInfoV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.ClientInfo = op.ClientInfo
}
func (o *SetClientInfoV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetClientInfoV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientInfoV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClientInfoV4Response structure represents the R_DhcpSetClientInfoV4 operation response
type SetClientInfoV4Response struct {
	// Return: The R_DhcpSetClientInfoV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetClientInfoV4Response) xxx_ToOp(ctx context.Context) *xxx_SetClientInfoV4Operation {
	if o == nil {
		return &xxx_SetClientInfoV4Operation{}
	}
	return &xxx_SetClientInfoV4Operation{
		Return: o.Return,
	}
}

func (o *SetClientInfoV4Response) xxx_FromOp(ctx context.Context, op *xxx_SetClientInfoV4Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetClientInfoV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetClientInfoV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientInfoV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClientInfoV4Operation structure represents the R_DhcpGetClientInfoV4 operation
type xxx_GetClientInfoV4Operation struct {
	ServerIPAddress string              `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SearchInfo      *dhcpm.SearchInfo   `idl:"name:SearchInfo;pointer:ref" json:"search_info"`
	ClientInfo      *dhcpm.ClientInfoV4 `idl:"name:ClientInfo" json:"client_info"`
	Return          uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClientInfoV4Operation) OpNum() int { return 34 }

func (o *xxx_GetClientInfoV4Operation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetClientInfoV4" }

func (o *xxx_GetClientInfoV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SearchInfo {in} (1:{pointer=ref, alias=LPDHCP_SEARCH_INFO}*(1))(2:{alias=DHCP_SEARCH_INFO}(struct))
	{
		if o.SearchInfo != nil {
			if err := o.SearchInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SearchInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SearchInfo {in} (1:{pointer=ref, alias=LPDHCP_SEARCH_INFO}*(1))(2:{alias=DHCP_SEARCH_INFO}(struct))
	{
		if o.SearchInfo == nil {
			o.SearchInfo = &dhcpm.SearchInfo{}
		}
		if err := o.SearchInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO_V4}*(1))(3:{alias=DHCP_CLIENT_INFO_V4}(struct))
	{
		if o.ClientInfo != nil {
			_ptr_ClientInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientInfo != nil {
					if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.ClientInfoV4{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientInfo, _ptr_ClientInfo); err != nil {
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

func (o *xxx_GetClientInfoV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO_V4,pointer=ref}*(1))(3:{alias=DHCP_CLIENT_INFO_V4}(struct))
	{
		_ptr_ClientInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientInfo == nil {
				o.ClientInfo = &dhcpm.ClientInfoV4{}
			}
			if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ClientInfo := func(ptr interface{}) { o.ClientInfo = *ptr.(**dhcpm.ClientInfoV4) }
		if err := w.ReadPointer(&o.ClientInfo, _s_ClientInfo, _ptr_ClientInfo); err != nil {
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

// GetClientInfoV4Request structure represents the R_DhcpGetClientInfoV4 operation request
type GetClientInfoV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SearchInfo: This is a pointer of type DHCP_SEARCH_INFO (section 2.2.1.2.18) structure
	// that defines the key to be used to search the DHCPv4 client lease record on the DHCPv4
	// server. In case the SearchType member is DhcpClientName and there are multiple lease
	// records with the same ClientName, the server will return client information for the
	// client having the lowest numerical IP address.
	SearchInfo *dhcpm.SearchInfo `idl:"name:SearchInfo;pointer:ref" json:"search_info"`
}

func (o *GetClientInfoV4Request) xxx_ToOp(ctx context.Context) *xxx_GetClientInfoV4Operation {
	if o == nil {
		return &xxx_GetClientInfoV4Operation{}
	}
	return &xxx_GetClientInfoV4Operation{
		ServerIPAddress: o.ServerIPAddress,
		SearchInfo:      o.SearchInfo,
	}
}

func (o *GetClientInfoV4Request) xxx_FromOp(ctx context.Context, op *xxx_GetClientInfoV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SearchInfo = op.SearchInfo
}
func (o *GetClientInfoV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClientInfoV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientInfoV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClientInfoV4Response structure represents the R_DhcpGetClientInfoV4 operation response
type GetClientInfoV4Response struct {
	// ClientInfo: This is a pointer of type LPDHCP_CLIENT_INFO_V4 that points to the location
	// to which a specific DHCPv4 client lease record is retrieved. The caller is responsible
	// for freeing this memory. The ClientHardwareAddress member represents a DHCPv4 client
	// unique ID (section 2.2.1.2.5.2).
	ClientInfo *dhcpm.ClientInfoV4 `idl:"name:ClientInfo" json:"client_info"`
	// Return: The R_DhcpGetClientInfoV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetClientInfoV4Response) xxx_ToOp(ctx context.Context) *xxx_GetClientInfoV4Operation {
	if o == nil {
		return &xxx_GetClientInfoV4Operation{}
	}
	return &xxx_GetClientInfoV4Operation{
		ClientInfo: o.ClientInfo,
		Return:     o.Return,
	}
}

func (o *GetClientInfoV4Response) xxx_FromOp(ctx context.Context, op *xxx_GetClientInfoV4Operation) {
	if o == nil {
		return
	}
	o.ClientInfo = op.ClientInfo
	o.Return = op.Return
}
func (o *GetClientInfoV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClientInfoV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientInfoV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumSubnetClientsV4Operation structure represents the R_DhcpEnumSubnetClientsV4 operation
type xxx_EnumSubnetClientsV4Operation struct {
	ServerIPAddress  string                   `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress    uint32                   `idl:"name:SubnetAddress" json:"subnet_address"`
	Resume           uint32                   `idl:"name:ResumeHandle" json:"resume"`
	PreferredMaximum uint32                   `idl:"name:PreferredMaximum" json:"preferred_maximum"`
	ClientInfo       *dhcpm.ClientInfoArrayV4 `idl:"name:ClientInfo" json:"client_info"`
	ClientsRead      uint32                   `idl:"name:ClientsRead" json:"clients_read"`
	ClientsTotal     uint32                   `idl:"name:ClientsTotal" json:"clients_total"`
	Return           uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumSubnetClientsV4Operation) OpNum() int { return 35 }

func (o *xxx_EnumSubnetClientsV4Operation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpEnumSubnetClientsV4"
}

func (o *xxx_EnumSubnetClientsV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO_ARRAY_V4}*(1))(3:{alias=DHCP_CLIENT_INFO_ARRAY_V4}(struct))
	{
		if o.ClientInfo != nil {
			_ptr_ClientInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientInfo != nil {
					if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.ClientInfoArrayV4{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientInfo, _ptr_ClientInfo); err != nil {
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
	// ClientsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientsRead); err != nil {
			return err
		}
	}
	// ClientsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientsTotal); err != nil {
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

func (o *xxx_EnumSubnetClientsV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO_ARRAY_V4,pointer=ref}*(1))(3:{alias=DHCP_CLIENT_INFO_ARRAY_V4}(struct))
	{
		_ptr_ClientInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientInfo == nil {
				o.ClientInfo = &dhcpm.ClientInfoArrayV4{}
			}
			if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ClientInfo := func(ptr interface{}) { o.ClientInfo = *ptr.(**dhcpm.ClientInfoArrayV4) }
		if err := w.ReadPointer(&o.ClientInfo, _s_ClientInfo, _ptr_ClientInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientsRead); err != nil {
			return err
		}
	}
	// ClientsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientsTotal); err != nil {
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

// EnumSubnetClientsV4Request structure represents the R_DhcpEnumSubnetClientsV4 operation request
type EnumSubnetClientsV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing the
	// IPv4 subnet ID from which DHCPv4 clients are enumerated. If this parameter is set
	// to 0, the DHCPv4 clients from all the IPv4 subnets are returned.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies the enumeration operation. Initially, this value MUST be set to zero,
	// with a successful call returning the handle value used for subsequent enumeration
	// requests. This field contains the last IPv4 address retrieved.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// PreferredMaximum: This is of type DWORD, specifying the preferred maximum number
	// of bytes to return. The minimum value is 1,024 bytes (1 kilobyte), and the maximum
	// value is 65,536 bytes (64 kilobytes); if the input value is greater or less than
	// this range, it MUST be set to the maximum or minimum value, respectively. To retrieve
	// all the DHCPv4 clients serviced from a specific IPv4 subnet, 0xFFFFFFFF is specified.
	PreferredMaximum uint32 `idl:"name:PreferredMaximum" json:"preferred_maximum"`
}

func (o *EnumSubnetClientsV4Request) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetClientsV4Operation {
	if o == nil {
		return &xxx_EnumSubnetClientsV4Operation{}
	}
	return &xxx_EnumSubnetClientsV4Operation{
		ServerIPAddress:  o.ServerIPAddress,
		SubnetAddress:    o.SubnetAddress,
		Resume:           o.Resume,
		PreferredMaximum: o.PreferredMaximum,
	}
}

func (o *EnumSubnetClientsV4Request) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetClientsV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.Resume = op.Resume
	o.PreferredMaximum = op.PreferredMaximum
}
func (o *EnumSubnetClientsV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumSubnetClientsV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetClientsV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumSubnetClientsV4Response structure represents the R_DhcpEnumSubnetClientsV4 operation response
type EnumSubnetClientsV4Response struct {
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies the enumeration operation. Initially, this value MUST be set to zero,
	// with a successful call returning the handle value used for subsequent enumeration
	// requests. This field contains the last IPv4 address retrieved.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// ClientInfo: This is a pointer of type LPDHCP_CLIENT_INFO_ARRAY_V4 that points to
	// the location that contains the DHCPv4 client lease record array.
	ClientInfo *dhcpm.ClientInfoArrayV4 `idl:"name:ClientInfo" json:"client_info"`
	// ClientsRead: This is a pointer to a DWORD that specifies the number of DHCPv4 client
	// lease records read in the ClientInfo parameter. The caller MUST allocate memory for
	// this parameter equal to the size of data type DWORD.
	ClientsRead uint32 `idl:"name:ClientsRead" json:"clients_read"`
	// ClientsTotal: This is a pointer to a DWORD that specifies the number of DHCPv4 client
	// lease records remaining from the current position. For example, if there are 100
	// DHCPv4 lease record clients for an IPv4 subnet, and if 10 DHCPv4 lease records are
	// enumerated per call, then for the first time this would have a value of 90.<37> The
	// caller MUST allocate memory for this parameter equal to the size of data type DWORD.
	ClientsTotal uint32 `idl:"name:ClientsTotal" json:"clients_total"`
	// Return: The R_DhcpEnumSubnetClientsV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumSubnetClientsV4Response) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetClientsV4Operation {
	if o == nil {
		return &xxx_EnumSubnetClientsV4Operation{}
	}
	return &xxx_EnumSubnetClientsV4Operation{
		Resume:       o.Resume,
		ClientInfo:   o.ClientInfo,
		ClientsRead:  o.ClientsRead,
		ClientsTotal: o.ClientsTotal,
		Return:       o.Return,
	}
}

func (o *EnumSubnetClientsV4Response) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetClientsV4Operation) {
	if o == nil {
		return
	}
	o.Resume = op.Resume
	o.ClientInfo = op.ClientInfo
	o.ClientsRead = op.ClientsRead
	o.ClientsTotal = op.ClientsTotal
	o.Return = op.Return
}
func (o *EnumSubnetClientsV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumSubnetClientsV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetClientsV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSuperScopeV4Operation structure represents the R_DhcpSetSuperScopeV4 operation
type xxx_SetSuperScopeV4Operation struct {
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	SuperScopeName  string `idl:"name:SuperScopeName;string;pointer:unique" json:"super_scope_name"`
	ChangeExisting  bool   `idl:"name:ChangeExisting" json:"change_existing"`
	Return          uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSuperScopeV4Operation) OpNum() int { return 36 }

func (o *xxx_SetSuperScopeV4Operation) OpName() string { return "/dhcpsrv/v1/R_DhcpSetSuperScopeV4" }

func (o *xxx_SetSuperScopeV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSuperScopeV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// SuperScopeName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.SuperScopeName != "" {
			_ptr_SuperScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SuperScopeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SuperScopeName, _ptr_SuperScopeName); err != nil {
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
	// ChangeExisting {in} (1:{alias=BOOL}(int32))
	{
		if !o.ChangeExisting {
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

func (o *xxx_SetSuperScopeV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// SuperScopeName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_SuperScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SuperScopeName); err != nil {
				return err
			}
			return nil
		})
		_s_SuperScopeName := func(ptr interface{}) { o.SuperScopeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.SuperScopeName, _s_SuperScopeName, _ptr_SuperScopeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ChangeExisting {in} (1:{alias=BOOL}(int32))
	{
		var _bChangeExisting int32
		if err := w.ReadData(&_bChangeExisting); err != nil {
			return err
		}
		o.ChangeExisting = _bChangeExisting != 0
	}
	return nil
}

func (o *xxx_SetSuperScopeV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSuperScopeV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSuperScopeV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSuperScopeV4Request structure represents the R_DhcpSetSuperScopeV4 operation request
type SetSuperScopeV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing the
	// IPv4 address of the subnet that is being set in a superscope or is removed from a
	// superscope.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// SuperScopeName: This is of a pointer of WCHAR that points to the location that contains
	// the superscope name. If NULL is specified in this field, the subnet corresponding
	// to the SubnetAddress parameter is removed from any superscope that it was part of.
	SuperScopeName string `idl:"name:SuperScopeName;string;pointer:unique" json:"super_scope_name"`
	// ChangeExisting: This is a BOOL that MUST be set to TRUE if the IPv4 subnet is already
	// part of any superscope.
	ChangeExisting bool `idl:"name:ChangeExisting" json:"change_existing"`
}

func (o *SetSuperScopeV4Request) xxx_ToOp(ctx context.Context) *xxx_SetSuperScopeV4Operation {
	if o == nil {
		return &xxx_SetSuperScopeV4Operation{}
	}
	return &xxx_SetSuperScopeV4Operation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
		SuperScopeName:  o.SuperScopeName,
		ChangeExisting:  o.ChangeExisting,
	}
}

func (o *SetSuperScopeV4Request) xxx_FromOp(ctx context.Context, op *xxx_SetSuperScopeV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.SuperScopeName = op.SuperScopeName
	o.ChangeExisting = op.ChangeExisting
}
func (o *SetSuperScopeV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSuperScopeV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSuperScopeV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSuperScopeV4Response structure represents the R_DhcpSetSuperScopeV4 operation response
type SetSuperScopeV4Response struct {
	// Return: The R_DhcpSetSuperScopeV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetSuperScopeV4Response) xxx_ToOp(ctx context.Context) *xxx_SetSuperScopeV4Operation {
	if o == nil {
		return &xxx_SetSuperScopeV4Operation{}
	}
	return &xxx_SetSuperScopeV4Operation{
		Return: o.Return,
	}
}

func (o *SetSuperScopeV4Response) xxx_FromOp(ctx context.Context, op *xxx_SetSuperScopeV4Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetSuperScopeV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSuperScopeV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSuperScopeV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSuperScopeInfoV4Operation structure represents the R_DhcpGetSuperScopeInfoV4 operation
type xxx_GetSuperScopeInfoV4Operation struct {
	ServerIPAddress string                 `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SuperScopeTable *dhcpm.SuperScopeTable `idl:"name:SuperScopeTable" json:"super_scope_table"`
	Return          uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSuperScopeInfoV4Operation) OpNum() int { return 37 }

func (o *xxx_GetSuperScopeInfoV4Operation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpGetSuperScopeInfoV4"
}

func (o *xxx_GetSuperScopeInfoV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSuperScopeInfoV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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

func (o *xxx_GetSuperScopeInfoV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSuperScopeInfoV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSuperScopeInfoV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SuperScopeTable {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SUPER_SCOPE_TABLE}*(1))(3:{alias=DHCP_SUPER_SCOPE_TABLE}(struct))
	{
		if o.SuperScopeTable != nil {
			_ptr_SuperScopeTable := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SuperScopeTable != nil {
					if err := o.SuperScopeTable.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.SuperScopeTable{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SuperScopeTable, _ptr_SuperScopeTable); err != nil {
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

func (o *xxx_GetSuperScopeInfoV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SuperScopeTable {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SUPER_SCOPE_TABLE,pointer=ref}*(1))(3:{alias=DHCP_SUPER_SCOPE_TABLE}(struct))
	{
		_ptr_SuperScopeTable := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SuperScopeTable == nil {
				o.SuperScopeTable = &dhcpm.SuperScopeTable{}
			}
			if err := o.SuperScopeTable.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_SuperScopeTable := func(ptr interface{}) { o.SuperScopeTable = *ptr.(**dhcpm.SuperScopeTable) }
		if err := w.ReadPointer(&o.SuperScopeTable, _s_SuperScopeTable, _ptr_SuperScopeTable); err != nil {
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

// GetSuperScopeInfoV4Request structure represents the R_DhcpGetSuperScopeInfoV4 operation request
type GetSuperScopeInfoV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
}

func (o *GetSuperScopeInfoV4Request) xxx_ToOp(ctx context.Context) *xxx_GetSuperScopeInfoV4Operation {
	if o == nil {
		return &xxx_GetSuperScopeInfoV4Operation{}
	}
	return &xxx_GetSuperScopeInfoV4Operation{
		ServerIPAddress: o.ServerIPAddress,
	}
}

func (o *GetSuperScopeInfoV4Request) xxx_FromOp(ctx context.Context, op *xxx_GetSuperScopeInfoV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
}
func (o *GetSuperScopeInfoV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSuperScopeInfoV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSuperScopeInfoV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSuperScopeInfoV4Response structure represents the R_DhcpGetSuperScopeInfoV4 operation response
type GetSuperScopeInfoV4Response struct {
	// SuperScopeTable: It is a pointer to type LPDHCP_SUPER_SCOPE_TABLE that points to
	// a location that contains the information for all the superscopes.
	SuperScopeTable *dhcpm.SuperScopeTable `idl:"name:SuperScopeTable" json:"super_scope_table"`
	// Return: The R_DhcpGetSuperScopeInfoV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetSuperScopeInfoV4Response) xxx_ToOp(ctx context.Context) *xxx_GetSuperScopeInfoV4Operation {
	if o == nil {
		return &xxx_GetSuperScopeInfoV4Operation{}
	}
	return &xxx_GetSuperScopeInfoV4Operation{
		SuperScopeTable: o.SuperScopeTable,
		Return:          o.Return,
	}
}

func (o *GetSuperScopeInfoV4Response) xxx_FromOp(ctx context.Context, op *xxx_GetSuperScopeInfoV4Operation) {
	if o == nil {
		return
	}
	o.SuperScopeTable = op.SuperScopeTable
	o.Return = op.Return
}
func (o *GetSuperScopeInfoV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSuperScopeInfoV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSuperScopeInfoV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteSuperScopeV4Operation structure represents the R_DhcpDeleteSuperScopeV4 operation
type xxx_DeleteSuperScopeV4Operation struct {
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SuperScopeName  string `idl:"name:SuperScopeName;string;pointer:ref" json:"super_scope_name"`
	Return          uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteSuperScopeV4Operation) OpNum() int { return 38 }

func (o *xxx_DeleteSuperScopeV4Operation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpDeleteSuperScopeV4"
}

func (o *xxx_DeleteSuperScopeV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteSuperScopeV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SuperScopeName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.SuperScopeName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteSuperScopeV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SuperScopeName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.SuperScopeName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteSuperScopeV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteSuperScopeV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteSuperScopeV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteSuperScopeV4Request structure represents the R_DhcpDeleteSuperScopeV4 operation request
type DeleteSuperScopeV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SuperScopeName: This is a pointer to WCHAR that points to the name of the superscope
	// that needs to be deleted.
	SuperScopeName string `idl:"name:SuperScopeName;string;pointer:ref" json:"super_scope_name"`
}

func (o *DeleteSuperScopeV4Request) xxx_ToOp(ctx context.Context) *xxx_DeleteSuperScopeV4Operation {
	if o == nil {
		return &xxx_DeleteSuperScopeV4Operation{}
	}
	return &xxx_DeleteSuperScopeV4Operation{
		ServerIPAddress: o.ServerIPAddress,
		SuperScopeName:  o.SuperScopeName,
	}
}

func (o *DeleteSuperScopeV4Request) xxx_FromOp(ctx context.Context, op *xxx_DeleteSuperScopeV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SuperScopeName = op.SuperScopeName
}
func (o *DeleteSuperScopeV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteSuperScopeV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteSuperScopeV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteSuperScopeV4Response structure represents the R_DhcpDeleteSuperScopeV4 operation response
type DeleteSuperScopeV4Response struct {
	// Return: The R_DhcpDeleteSuperScopeV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteSuperScopeV4Response) xxx_ToOp(ctx context.Context) *xxx_DeleteSuperScopeV4Operation {
	if o == nil {
		return &xxx_DeleteSuperScopeV4Operation{}
	}
	return &xxx_DeleteSuperScopeV4Operation{
		Return: o.Return,
	}
}

func (o *DeleteSuperScopeV4Response) xxx_FromOp(ctx context.Context, op *xxx_DeleteSuperScopeV4Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteSuperScopeV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteSuperScopeV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteSuperScopeV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerSetConfigV4Operation structure represents the R_DhcpServerSetConfigV4 operation
type xxx_ServerSetConfigV4Operation struct {
	ServerIPAddress string                    `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	FieldsToSet     uint32                    `idl:"name:FieldsToSet" json:"fields_to_set"`
	ConfigInfo      *dhcpm.ServerConfigInfoV4 `idl:"name:ConfigInfo;pointer:ref" json:"config_info"`
	Return          uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerSetConfigV4Operation) OpNum() int { return 39 }

func (o *xxx_ServerSetConfigV4Operation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpServerSetConfigV4"
}

func (o *xxx_ServerSetConfigV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// FieldsToSet {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.FieldsToSet); err != nil {
			return err
		}
	}
	// ConfigInfo {in} (1:{pointer=ref, alias=LPDHCP_SERVER_CONFIG_INFO_V4}*(1))(2:{alias=DHCP_SERVER_CONFIG_INFO_V4}(struct))
	{
		if o.ConfigInfo != nil {
			if err := o.ConfigInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.ServerConfigInfoV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// FieldsToSet {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.FieldsToSet); err != nil {
			return err
		}
	}
	// ConfigInfo {in} (1:{pointer=ref, alias=LPDHCP_SERVER_CONFIG_INFO_V4}*(1))(2:{alias=DHCP_SERVER_CONFIG_INFO_V4}(struct))
	{
		if o.ConfigInfo == nil {
			o.ConfigInfo = &dhcpm.ServerConfigInfoV4{}
		}
		if err := o.ConfigInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ServerSetConfigV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerSetConfigV4Request structure represents the R_DhcpServerSetConfigV4 operation request
type ServerSetConfigV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// FieldsToSet: A DWORD that contains the bitmask of the fields in the ConfigInfo parameter
	// to set. This method can be called with a value for the FieldsToSet parameter.
	//
	// The bit mapping for the various values for the FieldsToSet parameter is listed in
	// the following table.
	//
	//	+-----------------------------+------------+
	//	|                             |            |
	//	|         FIELDSTOSET         |    BIT     |
	//	|                             |            |
	//	+-----------------------------+------------+
	//	+-----------------------------+------------+
	//	| Set_APIProtocolSupport      | 0x00000001 |
	//	+-----------------------------+------------+
	//	| Set_DatabaseName            | 0x00000002 |
	//	+-----------------------------+------------+
	//	| Set_DatabasePath            | 0x00000004 |
	//	+-----------------------------+------------+
	//	| Set_BackupPath              | 0x00000008 |
	//	+-----------------------------+------------+
	//	| Set_BackupInterval          | 0x00000010 |
	//	+-----------------------------+------------+
	//	| Set_DatabaseLoggingFlag     | 0x00000020 |
	//	+-----------------------------+------------+
	//	| Set_RestoreFlag             | 0x00000040 |
	//	+-----------------------------+------------+
	//	| Set_DatabaseCleanupInterval | 0x00000080 |
	//	+-----------------------------+------------+
	//	| Set_DebugFlag               | 0x00000100 |
	//	+-----------------------------+------------+
	//	| Set_PingRetries             | 0x00000200 |
	//	+-----------------------------+------------+
	//	| Set_BootFileTable           | 0x00000400 |
	//	+-----------------------------+------------+
	//	| Set_AuditLogState           | 0x00000800 |
	//	+-----------------------------+------------+
	//
	// The DHCP server behavior ignores the bits not specified in the table.
	//
	// * Set_APIProtocolSupport
	//
	// * Set_DatabaseName
	//
	// * Set_DatabasePath
	//
	// * Set_DatabaseLoggingFlag
	//
	// * Set_RestoreFlag
	FieldsToSet uint32 `idl:"name:FieldsToSet" json:"fields_to_set"`
	// ConfigInfo: A pointer of type DHCP_SERVER_CONFIG_INFO_V4 (section 2.2.1.2.54) structure,
	// containing the settings for the DHCPv4 server. The value that is passed here depends
	// on the FieldsToSet parameter. Details of the dependencies follow the return value
	// description.
	ConfigInfo *dhcpm.ServerConfigInfoV4 `idl:"name:ConfigInfo;pointer:ref" json:"config_info"`
}

func (o *ServerSetConfigV4Request) xxx_ToOp(ctx context.Context) *xxx_ServerSetConfigV4Operation {
	if o == nil {
		return &xxx_ServerSetConfigV4Operation{}
	}
	return &xxx_ServerSetConfigV4Operation{
		ServerIPAddress: o.ServerIPAddress,
		FieldsToSet:     o.FieldsToSet,
		ConfigInfo:      o.ConfigInfo,
	}
}

func (o *ServerSetConfigV4Request) xxx_FromOp(ctx context.Context, op *xxx_ServerSetConfigV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.FieldsToSet = op.FieldsToSet
	o.ConfigInfo = op.ConfigInfo
}
func (o *ServerSetConfigV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ServerSetConfigV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetConfigV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerSetConfigV4Response structure represents the R_DhcpServerSetConfigV4 operation response
type ServerSetConfigV4Response struct {
	// Return: The R_DhcpServerSetConfigV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerSetConfigV4Response) xxx_ToOp(ctx context.Context) *xxx_ServerSetConfigV4Operation {
	if o == nil {
		return &xxx_ServerSetConfigV4Operation{}
	}
	return &xxx_ServerSetConfigV4Operation{
		Return: o.Return,
	}
}

func (o *ServerSetConfigV4Response) xxx_FromOp(ctx context.Context, op *xxx_ServerSetConfigV4Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ServerSetConfigV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ServerSetConfigV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetConfigV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerGetConfigV4Operation structure represents the R_DhcpServerGetConfigV4 operation
type xxx_ServerGetConfigV4Operation struct {
	ServerIPAddress string                    `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ConfigInfo      *dhcpm.ServerConfigInfoV4 `idl:"name:ConfigInfo" json:"config_info"`
	Return          uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerGetConfigV4Operation) OpNum() int { return 40 }

func (o *xxx_ServerGetConfigV4Operation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpServerGetConfigV4"
}

func (o *xxx_ServerGetConfigV4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetConfigV4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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

func (o *xxx_ServerGetConfigV4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetConfigV4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetConfigV4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ConfigInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SERVER_CONFIG_INFO_V4}*(1))(3:{alias=DHCP_SERVER_CONFIG_INFO_V4}(struct))
	{
		if o.ConfigInfo != nil {
			_ptr_ConfigInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigInfo != nil {
					if err := o.ConfigInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.ServerConfigInfoV4{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigInfo, _ptr_ConfigInfo); err != nil {
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

func (o *xxx_ServerGetConfigV4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ConfigInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SERVER_CONFIG_INFO_V4,pointer=ref}*(1))(3:{alias=DHCP_SERVER_CONFIG_INFO_V4}(struct))
	{
		_ptr_ConfigInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigInfo == nil {
				o.ConfigInfo = &dhcpm.ServerConfigInfoV4{}
			}
			if err := o.ConfigInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ConfigInfo := func(ptr interface{}) { o.ConfigInfo = *ptr.(**dhcpm.ServerConfigInfoV4) }
		if err := w.ReadPointer(&o.ConfigInfo, _s_ConfigInfo, _ptr_ConfigInfo); err != nil {
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

// ServerGetConfigV4Request structure represents the R_DhcpServerGetConfigV4 operation request
type ServerGetConfigV4Request struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
}

func (o *ServerGetConfigV4Request) xxx_ToOp(ctx context.Context) *xxx_ServerGetConfigV4Operation {
	if o == nil {
		return &xxx_ServerGetConfigV4Operation{}
	}
	return &xxx_ServerGetConfigV4Operation{
		ServerIPAddress: o.ServerIPAddress,
	}
}

func (o *ServerGetConfigV4Request) xxx_FromOp(ctx context.Context, op *xxx_ServerGetConfigV4Operation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
}
func (o *ServerGetConfigV4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ServerGetConfigV4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetConfigV4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerGetConfigV4Response structure represents the R_DhcpServerGetConfigV4 operation response
type ServerGetConfigV4Response struct {
	// ConfigInfo: This is a pointer of type LPDHCP_SERVER_CONFIG_INFO_V4 that points to
	// the location where the DHCPv4 server settings are retrieved. The caller of this method
	// can free up this structure after use.
	ConfigInfo *dhcpm.ServerConfigInfoV4 `idl:"name:ConfigInfo" json:"config_info"`
	// Return: The R_DhcpServerGetConfigV4 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerGetConfigV4Response) xxx_ToOp(ctx context.Context) *xxx_ServerGetConfigV4Operation {
	if o == nil {
		return &xxx_ServerGetConfigV4Operation{}
	}
	return &xxx_ServerGetConfigV4Operation{
		ConfigInfo: o.ConfigInfo,
		Return:     o.Return,
	}
}

func (o *ServerGetConfigV4Response) xxx_FromOp(ctx context.Context, op *xxx_ServerGetConfigV4Operation) {
	if o == nil {
		return
	}
	o.ConfigInfo = op.ConfigInfo
	o.Return = op.Return
}
func (o *ServerGetConfigV4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ServerGetConfigV4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetConfigV4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerSetConfigVQOperation structure represents the R_DhcpServerSetConfigVQ operation
type xxx_ServerSetConfigVQOperation struct {
	ServerIPAddress string                    `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	FieldsToSet     uint32                    `idl:"name:FieldsToSet" json:"fields_to_set"`
	ConfigInfo      *dhcpm.ServerConfigInfoVQ `idl:"name:ConfigInfo;pointer:ref" json:"config_info"`
	Return          uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerSetConfigVQOperation) OpNum() int { return 41 }

func (o *xxx_ServerSetConfigVQOperation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpServerSetConfigVQ"
}

func (o *xxx_ServerSetConfigVQOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigVQOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// FieldsToSet {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.FieldsToSet); err != nil {
			return err
		}
	}
	// ConfigInfo {in} (1:{pointer=ref, alias=LPDHCP_SERVER_CONFIG_INFO_VQ}*(1))(2:{alias=DHCP_SERVER_CONFIG_INFO_VQ}(struct))
	{
		if o.ConfigInfo != nil {
			if err := o.ConfigInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.ServerConfigInfoVQ{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigVQOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// FieldsToSet {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.FieldsToSet); err != nil {
			return err
		}
	}
	// ConfigInfo {in} (1:{pointer=ref, alias=LPDHCP_SERVER_CONFIG_INFO_VQ}*(1))(2:{alias=DHCP_SERVER_CONFIG_INFO_VQ}(struct))
	{
		if o.ConfigInfo == nil {
			o.ConfigInfo = &dhcpm.ServerConfigInfoVQ{}
		}
		if err := o.ConfigInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigVQOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetConfigVQOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ServerSetConfigVQOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerSetConfigVQRequest structure represents the R_DhcpServerSetConfigVQ operation request
type ServerSetConfigVQRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// FieldsToSet: A DWORD that contains the bitmask of the fields in the ConfigInfo parameter
	// to set. This method can be called with a value for the FieldsToSet parameter.
	//
	// The bit mapping for the various values for the FieldsToSet parameter is listed in
	// the following table.
	//
	//	+-----------------------------+------------+
	//	|                             |            |
	//	|         FIELDSTOSET         |    BIT     |
	//	|                             |            |
	//	+-----------------------------+------------+
	//	+-----------------------------+------------+
	//	| Set_APIProtocolSupport      | 0x00000001 |
	//	+-----------------------------+------------+
	//	| Set_DatabaseName            | 0x00000002 |
	//	+-----------------------------+------------+
	//	| Set_DatabasePath            | 0x00000004 |
	//	+-----------------------------+------------+
	//	| Set_BackupPath              | 0x00000008 |
	//	+-----------------------------+------------+
	//	| Set_BackupInterval          | 0x00000010 |
	//	+-----------------------------+------------+
	//	| Set_DatabaseLoggingFlag     | 0x00000020 |
	//	+-----------------------------+------------+
	//	| Set_RestoreFlag             | 0x00000040 |
	//	+-----------------------------+------------+
	//	| Set_DatabaseCleanupInterval | 0x00000080 |
	//	+-----------------------------+------------+
	//	| Set_DebugFlag               | 0x00000100 |
	//	+-----------------------------+------------+
	//	| Set_PingRetries             | 0x00000200 |
	//	+-----------------------------+------------+
	//	| Set_BootFileTable           | 0x00000400 |
	//	+-----------------------------+------------+
	//	| Set_AuditLogState           | 0x00000800 |
	//	+-----------------------------+------------+
	//	| Set_QuarantineON            | 0x00001000 |
	//	+-----------------------------+------------+
	//	| Set_QuarantineDefFail       | 0x00002000 |
	//	+-----------------------------+------------+
	//
	// * Set_APIProtocolSupport
	//
	// * Set_DatabaseName
	//
	// * Set_DatabasePath
	//
	// * Set_DatabaseLoggingFlag
	//
	// * Set_RestoreFlag
	FieldsToSet uint32 `idl:"name:FieldsToSet" json:"fields_to_set"`
	// ConfigInfo: A pointer of type DHCP_SERVER_CONFIG_INFO_VQ (section 2.2.1.2.55) structure
	// that contains the settings for the DHCPv4 server. The value that is passed here depends
	// on the FieldsToSet parameter. Details of the dependencies follow the return value
	// description.
	ConfigInfo *dhcpm.ServerConfigInfoVQ `idl:"name:ConfigInfo;pointer:ref" json:"config_info"`
}

func (o *ServerSetConfigVQRequest) xxx_ToOp(ctx context.Context) *xxx_ServerSetConfigVQOperation {
	if o == nil {
		return &xxx_ServerSetConfigVQOperation{}
	}
	return &xxx_ServerSetConfigVQOperation{
		ServerIPAddress: o.ServerIPAddress,
		FieldsToSet:     o.FieldsToSet,
		ConfigInfo:      o.ConfigInfo,
	}
}

func (o *ServerSetConfigVQRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerSetConfigVQOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.FieldsToSet = op.FieldsToSet
	o.ConfigInfo = op.ConfigInfo
}
func (o *ServerSetConfigVQRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ServerSetConfigVQRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetConfigVQOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerSetConfigVQResponse structure represents the R_DhcpServerSetConfigVQ operation response
type ServerSetConfigVQResponse struct {
	// Return: The R_DhcpServerSetConfigVQ return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerSetConfigVQResponse) xxx_ToOp(ctx context.Context) *xxx_ServerSetConfigVQOperation {
	if o == nil {
		return &xxx_ServerSetConfigVQOperation{}
	}
	return &xxx_ServerSetConfigVQOperation{
		Return: o.Return,
	}
}

func (o *ServerSetConfigVQResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerSetConfigVQOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ServerSetConfigVQResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ServerSetConfigVQResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetConfigVQOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerGetConfigVQOperation structure represents the R_DhcpServerGetConfigVQ operation
type xxx_ServerGetConfigVQOperation struct {
	ServerIPAddress string                    `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ConfigInfo      *dhcpm.ServerConfigInfoVQ `idl:"name:ConfigInfo" json:"config_info"`
	Return          uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerGetConfigVQOperation) OpNum() int { return 42 }

func (o *xxx_ServerGetConfigVQOperation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpServerGetConfigVQ"
}

func (o *xxx_ServerGetConfigVQOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetConfigVQOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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

func (o *xxx_ServerGetConfigVQOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetConfigVQOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetConfigVQOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ConfigInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SERVER_CONFIG_INFO_VQ}*(1))(3:{alias=DHCP_SERVER_CONFIG_INFO_VQ}(struct))
	{
		if o.ConfigInfo != nil {
			_ptr_ConfigInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigInfo != nil {
					if err := o.ConfigInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.ServerConfigInfoVQ{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigInfo, _ptr_ConfigInfo); err != nil {
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

func (o *xxx_ServerGetConfigVQOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ConfigInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SERVER_CONFIG_INFO_VQ,pointer=ref}*(1))(3:{alias=DHCP_SERVER_CONFIG_INFO_VQ}(struct))
	{
		_ptr_ConfigInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigInfo == nil {
				o.ConfigInfo = &dhcpm.ServerConfigInfoVQ{}
			}
			if err := o.ConfigInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ConfigInfo := func(ptr interface{}) { o.ConfigInfo = *ptr.(**dhcpm.ServerConfigInfoVQ) }
		if err := w.ReadPointer(&o.ConfigInfo, _s_ConfigInfo, _ptr_ConfigInfo); err != nil {
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

// ServerGetConfigVQRequest structure represents the R_DhcpServerGetConfigVQ operation request
type ServerGetConfigVQRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
}

func (o *ServerGetConfigVQRequest) xxx_ToOp(ctx context.Context) *xxx_ServerGetConfigVQOperation {
	if o == nil {
		return &xxx_ServerGetConfigVQOperation{}
	}
	return &xxx_ServerGetConfigVQOperation{
		ServerIPAddress: o.ServerIPAddress,
	}
}

func (o *ServerGetConfigVQRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerGetConfigVQOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
}
func (o *ServerGetConfigVQRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ServerGetConfigVQRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetConfigVQOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerGetConfigVQResponse structure represents the R_DhcpServerGetConfigVQ operation response
type ServerGetConfigVQResponse struct {
	// ConfigInfo: This is a pointer of type LPDHCP_SERVER_CONFIG_INFO_VQ that points to
	// the location where the DHCPv4 server settings are retrieved. The caller of this method
	// can free up this structure after use.
	ConfigInfo *dhcpm.ServerConfigInfoVQ `idl:"name:ConfigInfo" json:"config_info"`
	// Return: The R_DhcpServerGetConfigVQ return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerGetConfigVQResponse) xxx_ToOp(ctx context.Context) *xxx_ServerGetConfigVQOperation {
	if o == nil {
		return &xxx_ServerGetConfigVQOperation{}
	}
	return &xxx_ServerGetConfigVQOperation{
		ConfigInfo: o.ConfigInfo,
		Return:     o.Return,
	}
}

func (o *ServerGetConfigVQResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerGetConfigVQOperation) {
	if o == nil {
		return
	}
	o.ConfigInfo = op.ConfigInfo
	o.Return = op.Return
}
func (o *ServerGetConfigVQResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ServerGetConfigVQResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetConfigVQOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMIBInfoVQOperation structure represents the R_DhcpGetMibInfoVQ operation
type xxx_GetMIBInfoVQOperation struct {
	ServerIPAddress string           `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	MIBInfo         *dhcpm.MIBInfoVQ `idl:"name:MibInfo" json:"mib_info"`
	Return          uint32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMIBInfoVQOperation) OpNum() int { return 43 }

func (o *xxx_GetMIBInfoVQOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetMibInfoVQ" }

func (o *xxx_GetMIBInfoVQOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMIBInfoVQOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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

func (o *xxx_GetMIBInfoVQOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMIBInfoVQOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMIBInfoVQOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// MibInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_MIB_INFO_VQ}*(1))(3:{alias=DHCP_MIB_INFO_VQ}(struct))
	{
		if o.MIBInfo != nil {
			_ptr_MibInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MIBInfo != nil {
					if err := o.MIBInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.MIBInfoVQ{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MIBInfo, _ptr_MibInfo); err != nil {
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

func (o *xxx_GetMIBInfoVQOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// MibInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_MIB_INFO_VQ,pointer=ref}*(1))(3:{alias=DHCP_MIB_INFO_VQ}(struct))
	{
		_ptr_MibInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MIBInfo == nil {
				o.MIBInfo = &dhcpm.MIBInfoVQ{}
			}
			if err := o.MIBInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_MibInfo := func(ptr interface{}) { o.MIBInfo = *ptr.(**dhcpm.MIBInfoVQ) }
		if err := w.ReadPointer(&o.MIBInfo, _s_MibInfo, _ptr_MibInfo); err != nil {
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

// GetMIBInfoVQRequest structure represents the R_DhcpGetMibInfoVQ operation request
type GetMIBInfoVQRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
}

func (o *GetMIBInfoVQRequest) xxx_ToOp(ctx context.Context) *xxx_GetMIBInfoVQOperation {
	if o == nil {
		return &xxx_GetMIBInfoVQOperation{}
	}
	return &xxx_GetMIBInfoVQOperation{
		ServerIPAddress: o.ServerIPAddress,
	}
}

func (o *GetMIBInfoVQRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMIBInfoVQOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
}
func (o *GetMIBInfoVQRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMIBInfoVQRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMIBInfoVQOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMIBInfoVQResponse structure represents the R_DhcpGetMibInfoVQ operation response
type GetMIBInfoVQResponse struct {
	// MibInfo: This is a pointer of type LPDHCP_MIB_INFO_VQ that points to the location
	// that contains the MIB information of the DHCPv4 server.
	MIBInfo *dhcpm.MIBInfoVQ `idl:"name:MibInfo" json:"mib_info"`
	// Return: The R_DhcpGetMibInfoVQ return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetMIBInfoVQResponse) xxx_ToOp(ctx context.Context) *xxx_GetMIBInfoVQOperation {
	if o == nil {
		return &xxx_GetMIBInfoVQOperation{}
	}
	return &xxx_GetMIBInfoVQOperation{
		MIBInfo: o.MIBInfo,
		Return:  o.Return,
	}
}

func (o *GetMIBInfoVQResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMIBInfoVQOperation) {
	if o == nil {
		return
	}
	o.MIBInfo = op.MIBInfo
	o.Return = op.Return
}
func (o *GetMIBInfoVQResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMIBInfoVQResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMIBInfoVQOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateClientInfoVQOperation structure represents the R_DhcpCreateClientInfoVQ operation
type xxx_CreateClientInfoVQOperation struct {
	ServerIPAddress string              `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ClientInfo      *dhcpm.ClientInfoVQ `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
	Return          uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateClientInfoVQOperation) OpNum() int { return 44 }

func (o *xxx_CreateClientInfoVQOperation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpCreateClientInfoVQ"
}

func (o *xxx_CreateClientInfoVQOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoVQOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO_VQ}*(1))(2:{alias=DHCP_CLIENT_INFO_VQ}(struct))
	{
		if o.ClientInfo != nil {
			if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.ClientInfoVQ{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoVQOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO_VQ}*(1))(2:{alias=DHCP_CLIENT_INFO_VQ}(struct))
	{
		if o.ClientInfo == nil {
			o.ClientInfo = &dhcpm.ClientInfoVQ{}
		}
		if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoVQOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateClientInfoVQOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateClientInfoVQOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateClientInfoVQRequest structure represents the R_DhcpCreateClientInfoVQ operation request
type CreateClientInfoVQRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ClientInfo: A pointer of type DHCP_CLIENT_INFO_VQ (section 2.2.1.2.19) structure
	// that contains the DHCPv4 client lease record information that needs to be set on
	// the DHCPv4 server. The caller MUST pass the ClientIPAddress and ClientHardwareAddress
	// members when adding a DHCPv4 client lease record to the DHCPv4 server database. The
	// ClientHardwareAddress member represents a DHCPv4 client-identifier (section 2.2.1.2.5.1).
	// Members ClientName, ClientComment, ClientLeaseExpires, and OwnerHost are modified
	// on the DHCPv4 client lease record identified by the ClientIPaddress member.
	ClientInfo *dhcpm.ClientInfoVQ `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
}

func (o *CreateClientInfoVQRequest) xxx_ToOp(ctx context.Context) *xxx_CreateClientInfoVQOperation {
	if o == nil {
		return &xxx_CreateClientInfoVQOperation{}
	}
	return &xxx_CreateClientInfoVQOperation{
		ServerIPAddress: o.ServerIPAddress,
		ClientInfo:      o.ClientInfo,
	}
}

func (o *CreateClientInfoVQRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateClientInfoVQOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.ClientInfo = op.ClientInfo
}
func (o *CreateClientInfoVQRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateClientInfoVQRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateClientInfoVQOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateClientInfoVQResponse structure represents the R_DhcpCreateClientInfoVQ operation response
type CreateClientInfoVQResponse struct {
	// Return: The R_DhcpCreateClientInfoVQ return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateClientInfoVQResponse) xxx_ToOp(ctx context.Context) *xxx_CreateClientInfoVQOperation {
	if o == nil {
		return &xxx_CreateClientInfoVQOperation{}
	}
	return &xxx_CreateClientInfoVQOperation{
		Return: o.Return,
	}
}

func (o *CreateClientInfoVQResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateClientInfoVQOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CreateClientInfoVQResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateClientInfoVQResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateClientInfoVQOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClientInfoVQOperation structure represents the R_DhcpSetClientInfoVQ operation
type xxx_SetClientInfoVQOperation struct {
	ServerIPAddress string              `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	ClientInfo      *dhcpm.ClientInfoVQ `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
	Return          uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClientInfoVQOperation) OpNum() int { return 45 }

func (o *xxx_SetClientInfoVQOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpSetClientInfoVQ" }

func (o *xxx_SetClientInfoVQOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoVQOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO_VQ}*(1))(2:{alias=DHCP_CLIENT_INFO_VQ}(struct))
	{
		if o.ClientInfo != nil {
			if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.ClientInfoVQ{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoVQOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientInfo {in} (1:{pointer=ref, alias=LPDHCP_CLIENT_INFO_VQ}*(1))(2:{alias=DHCP_CLIENT_INFO_VQ}(struct))
	{
		if o.ClientInfo == nil {
			o.ClientInfo = &dhcpm.ClientInfoVQ{}
		}
		if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoVQOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientInfoVQOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetClientInfoVQOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetClientInfoVQRequest structure represents the R_DhcpSetClientInfoVQ operation request
type SetClientInfoVQRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// ClientInfo: A pointer of type DHCP_CLIENT_INFO_VQ (section 2.2.1.2.19) structure
	// that contains the DHCPv4 client lease record information that needs to be modified
	// on the DHCPv4 server database. The caller MUST pass the ClientIPAddress and ClientHardwareAddress
	// members when modifying a DHCPv4 client lease record stored in the DHCPv4 server database.
	// The ClientHardwareAddress member represents a DHCPv4 client-identifier (section 2.2.1.2.5.1).
	// The members ClientName, ClientComment, ClientLeaseExpires, and OwnerHost are modified
	// in the DHCPv4 client lease record identified by the ClientIPAddress member.
	ClientInfo *dhcpm.ClientInfoVQ `idl:"name:ClientInfo;pointer:ref" json:"client_info"`
}

func (o *SetClientInfoVQRequest) xxx_ToOp(ctx context.Context) *xxx_SetClientInfoVQOperation {
	if o == nil {
		return &xxx_SetClientInfoVQOperation{}
	}
	return &xxx_SetClientInfoVQOperation{
		ServerIPAddress: o.ServerIPAddress,
		ClientInfo:      o.ClientInfo,
	}
}

func (o *SetClientInfoVQRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClientInfoVQOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.ClientInfo = op.ClientInfo
}
func (o *SetClientInfoVQRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetClientInfoVQRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientInfoVQOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClientInfoVQResponse structure represents the R_DhcpSetClientInfoVQ operation response
type SetClientInfoVQResponse struct {
	// Return: The R_DhcpSetClientInfoVQ return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetClientInfoVQResponse) xxx_ToOp(ctx context.Context) *xxx_SetClientInfoVQOperation {
	if o == nil {
		return &xxx_SetClientInfoVQOperation{}
	}
	return &xxx_SetClientInfoVQOperation{
		Return: o.Return,
	}
}

func (o *SetClientInfoVQResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClientInfoVQOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetClientInfoVQResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetClientInfoVQResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientInfoVQOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClientInfoVQOperation structure represents the R_DhcpGetClientInfoVQ operation
type xxx_GetClientInfoVQOperation struct {
	ServerIPAddress string              `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SearchInfo      *dhcpm.SearchInfo   `idl:"name:SearchInfo;pointer:ref" json:"search_info"`
	ClientInfo      *dhcpm.ClientInfoVQ `idl:"name:ClientInfo" json:"client_info"`
	Return          uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClientInfoVQOperation) OpNum() int { return 46 }

func (o *xxx_GetClientInfoVQOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetClientInfoVQ" }

func (o *xxx_GetClientInfoVQOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoVQOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SearchInfo {in} (1:{pointer=ref, alias=LPDHCP_SEARCH_INFO}*(1))(2:{alias=DHCP_SEARCH_INFO}(struct))
	{
		if o.SearchInfo != nil {
			if err := o.SearchInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SearchInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoVQOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SearchInfo {in} (1:{pointer=ref, alias=LPDHCP_SEARCH_INFO}*(1))(2:{alias=DHCP_SEARCH_INFO}(struct))
	{
		if o.SearchInfo == nil {
			o.SearchInfo = &dhcpm.SearchInfo{}
		}
		if err := o.SearchInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoVQOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientInfoVQOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO_VQ}*(1))(3:{alias=DHCP_CLIENT_INFO_VQ}(struct))
	{
		if o.ClientInfo != nil {
			_ptr_ClientInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientInfo != nil {
					if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.ClientInfoVQ{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientInfo, _ptr_ClientInfo); err != nil {
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

func (o *xxx_GetClientInfoVQOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO_VQ,pointer=ref}*(1))(3:{alias=DHCP_CLIENT_INFO_VQ}(struct))
	{
		_ptr_ClientInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientInfo == nil {
				o.ClientInfo = &dhcpm.ClientInfoVQ{}
			}
			if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ClientInfo := func(ptr interface{}) { o.ClientInfo = *ptr.(**dhcpm.ClientInfoVQ) }
		if err := w.ReadPointer(&o.ClientInfo, _s_ClientInfo, _ptr_ClientInfo); err != nil {
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

// GetClientInfoVQRequest structure represents the R_DhcpGetClientInfoVQ operation request
type GetClientInfoVQRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SearchInfo: This is a pointer of type DHCP_SEARCH_INFO (section 2.2.1.2.18) structure
	// that defines the key to be used to search the DHCPv4 client lease record on the DHCPv4
	// server. In case the SearchType member is DhcpClientName and there are multiple lease
	// records with the same ClientName member, the server will return client information
	// for the client having the lowest numerical IP address.
	SearchInfo *dhcpm.SearchInfo `idl:"name:SearchInfo;pointer:ref" json:"search_info"`
}

func (o *GetClientInfoVQRequest) xxx_ToOp(ctx context.Context) *xxx_GetClientInfoVQOperation {
	if o == nil {
		return &xxx_GetClientInfoVQOperation{}
	}
	return &xxx_GetClientInfoVQOperation{
		ServerIPAddress: o.ServerIPAddress,
		SearchInfo:      o.SearchInfo,
	}
}

func (o *GetClientInfoVQRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClientInfoVQOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SearchInfo = op.SearchInfo
}
func (o *GetClientInfoVQRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClientInfoVQRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientInfoVQOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClientInfoVQResponse structure represents the R_DhcpGetClientInfoVQ operation response
type GetClientInfoVQResponse struct {
	// ClientInfo: This is a pointer of type LPDHCP_CLIENT_INFO_VQ that points to the location
	// in which specific DHCPv4 client lease record information is retrieved. The caller
	// can free up this buffer after using this. The ClientHardwareAddress member represents
	// a DHCPv4 client unique ID (section 2.2.1.2.5.2).
	ClientInfo *dhcpm.ClientInfoVQ `idl:"name:ClientInfo" json:"client_info"`
	// Return: The R_DhcpGetClientInfoVQ return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetClientInfoVQResponse) xxx_ToOp(ctx context.Context) *xxx_GetClientInfoVQOperation {
	if o == nil {
		return &xxx_GetClientInfoVQOperation{}
	}
	return &xxx_GetClientInfoVQOperation{
		ClientInfo: o.ClientInfo,
		Return:     o.Return,
	}
}

func (o *GetClientInfoVQResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClientInfoVQOperation) {
	if o == nil {
		return
	}
	o.ClientInfo = op.ClientInfo
	o.Return = op.Return
}
func (o *GetClientInfoVQResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClientInfoVQResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientInfoVQOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumSubnetClientsVQOperation structure represents the R_DhcpEnumSubnetClientsVQ operation
type xxx_EnumSubnetClientsVQOperation struct {
	ServerIPAddress  string                   `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress    uint32                   `idl:"name:SubnetAddress" json:"subnet_address"`
	Resume           uint32                   `idl:"name:ResumeHandle" json:"resume"`
	PreferredMaximum uint32                   `idl:"name:PreferredMaximum" json:"preferred_maximum"`
	ClientInfo       *dhcpm.ClientInfoArrayVQ `idl:"name:ClientInfo" json:"client_info"`
	ClientsRead      uint32                   `idl:"name:ClientsRead" json:"clients_read"`
	ClientsTotal     uint32                   `idl:"name:ClientsTotal" json:"clients_total"`
	Return           uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumSubnetClientsVQOperation) OpNum() int { return 47 }

func (o *xxx_EnumSubnetClientsVQOperation) OpName() string {
	return "/dhcpsrv/v1/R_DhcpEnumSubnetClientsVQ"
}

func (o *xxx_EnumSubnetClientsVQOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsVQOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsVQOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// PreferredMaximum {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximum); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsVQOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumSubnetClientsVQOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Resume); err != nil {
			return err
		}
	}
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO_ARRAY_VQ}*(1))(3:{alias=DHCP_CLIENT_INFO_ARRAY_VQ}(struct))
	{
		if o.ClientInfo != nil {
			_ptr_ClientInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientInfo != nil {
					if err := o.ClientInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.ClientInfoArrayVQ{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientInfo, _ptr_ClientInfo); err != nil {
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
	// ClientsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientsRead); err != nil {
			return err
		}
	}
	// ClientsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientsTotal); err != nil {
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

func (o *xxx_EnumSubnetClientsVQOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ResumeHandle {in, out} (1:{pointer=ref}*(1))(2:{alias=DHCP_RESUME_HANDLE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Resume); err != nil {
			return err
		}
	}
	// ClientInfo {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_CLIENT_INFO_ARRAY_VQ,pointer=ref}*(1))(3:{alias=DHCP_CLIENT_INFO_ARRAY_VQ}(struct))
	{
		_ptr_ClientInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientInfo == nil {
				o.ClientInfo = &dhcpm.ClientInfoArrayVQ{}
			}
			if err := o.ClientInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ClientInfo := func(ptr interface{}) { o.ClientInfo = *ptr.(**dhcpm.ClientInfoArrayVQ) }
		if err := w.ReadPointer(&o.ClientInfo, _s_ClientInfo, _ptr_ClientInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientsRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientsRead); err != nil {
			return err
		}
	}
	// ClientsTotal {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientsTotal); err != nil {
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

// EnumSubnetClientsVQRequest structure represents the R_DhcpEnumSubnetClientsVQ operation request
type EnumSubnetClientsVQRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1) that contains
	// the IPv4 subnet ID from which DHCPv4 clients are enumerated. If this parameter is
	// set to 0, the DHCPv4 clients from all the IPv4 subnets are returned.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies the enumeration operation. Initially, this value MUST be set to zero,
	// with a successful call returning the handle value used for subsequent enumeration
	// requests. This field contains the last IPv4 address retrieved.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// PreferredMaximum: This is of type DWORD, specifying the preferred maximum number
	// of bytes to return. The minimum value is 1,024 bytes (1 kilobyte), and the maximum
	// value is 65,536 bytes (64 kilobytes); if the input value is greater or less than
	// this range, it MUST be set to the maximum or minimum value, respectively. To retrieve
	// all the DHCPv4 clients serviced by a specific IPv4 subnet, 0xFFFFFFFF is specified.
	PreferredMaximum uint32 `idl:"name:PreferredMaximum" json:"preferred_maximum"`
}

func (o *EnumSubnetClientsVQRequest) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetClientsVQOperation {
	if o == nil {
		return &xxx_EnumSubnetClientsVQOperation{}
	}
	return &xxx_EnumSubnetClientsVQOperation{
		ServerIPAddress:  o.ServerIPAddress,
		SubnetAddress:    o.SubnetAddress,
		Resume:           o.Resume,
		PreferredMaximum: o.PreferredMaximum,
	}
}

func (o *EnumSubnetClientsVQRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetClientsVQOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.Resume = op.Resume
	o.PreferredMaximum = op.PreferredMaximum
}
func (o *EnumSubnetClientsVQRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumSubnetClientsVQRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetClientsVQOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumSubnetClientsVQResponse structure represents the R_DhcpEnumSubnetClientsVQ operation response
type EnumSubnetClientsVQResponse struct {
	// ResumeHandle: This is a pointer of type DHCP_RESUME_HANDLE (section 2.2.1.2.6) that
	// identifies the enumeration operation. Initially, this value MUST be set to zero,
	// with a successful call returning the handle value used for subsequent enumeration
	// requests. This field contains the last IPv4 address retrieved.
	Resume uint32 `idl:"name:ResumeHandle" json:"resume"`
	// ClientInfo: This is a pointer of type LPDHCP_CLIENT_INFO_ARRAY_VQ that points to
	// the location that contains the DHCPv4 client lease record array.
	ClientInfo *dhcpm.ClientInfoArrayVQ `idl:"name:ClientInfo" json:"client_info"`
	// ClientsRead: This is a pointer to a DWORD that specifies the number of DHCPv4 client
	// lease records read in the ClientInfo parameter. The caller MUST allocate memory for
	// this parameter equal to the size of data type DWORD.
	ClientsRead uint32 `idl:"name:ClientsRead" json:"clients_read"`
	// ClientsTotal: This is a pointer to a DWORD that specifies the number of DHCPv4 client
	// lease records remaining from the current position. For example, if there are 100
	// DHCPv4 lease record clients for an IPv4 subnet, and if 10 DHCPv4 lease records are
	// enumerated per call, then for the first time this would have a value of 90.<42> The
	// caller MUST allocate memory for this parameter equal to the size of data type DWORD.
	ClientsTotal uint32 `idl:"name:ClientsTotal" json:"clients_total"`
	// Return: The R_DhcpEnumSubnetClientsVQ return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumSubnetClientsVQResponse) xxx_ToOp(ctx context.Context) *xxx_EnumSubnetClientsVQOperation {
	if o == nil {
		return &xxx_EnumSubnetClientsVQOperation{}
	}
	return &xxx_EnumSubnetClientsVQOperation{
		Resume:       o.Resume,
		ClientInfo:   o.ClientInfo,
		ClientsRead:  o.ClientsRead,
		ClientsTotal: o.ClientsTotal,
		Return:       o.Return,
	}
}

func (o *EnumSubnetClientsVQResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumSubnetClientsVQOperation) {
	if o == nil {
		return
	}
	o.Resume = op.Resume
	o.ClientInfo = op.ClientInfo
	o.ClientsRead = op.ClientsRead
	o.ClientsTotal = op.ClientsTotal
	o.Return = op.Return
}
func (o *EnumSubnetClientsVQResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumSubnetClientsVQResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumSubnetClientsVQOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateSubnetVQOperation structure represents the R_DhcpCreateSubnetVQ operation
type xxx_CreateSubnetVQOperation struct {
	ServerIPAddress string              `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32              `idl:"name:SubnetAddress" json:"subnet_address"`
	SubnetInfoVQ    *dhcpm.SubnetInfoVQ `idl:"name:SubnetInfoVQ;pointer:ref" json:"subnet_info_vq"`
	Return          uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateSubnetVQOperation) OpNum() int { return 48 }

func (o *xxx_CreateSubnetVQOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpCreateSubnetVQ" }

func (o *xxx_CreateSubnetVQOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSubnetVQOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// SubnetInfoVQ {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_INFO_VQ}*(1))(2:{alias=DHCP_SUBNET_INFO_VQ}(struct))
	{
		if o.SubnetInfoVQ != nil {
			if err := o.SubnetInfoVQ.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SubnetInfoVQ{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSubnetVQOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// SubnetInfoVQ {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_INFO_VQ}*(1))(2:{alias=DHCP_SUBNET_INFO_VQ}(struct))
	{
		if o.SubnetInfoVQ == nil {
			o.SubnetInfoVQ = &dhcpm.SubnetInfoVQ{}
		}
		if err := o.SubnetInfoVQ.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSubnetVQOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSubnetVQOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateSubnetVQOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateSubnetVQRequest structure represents the R_DhcpCreateSubnetVQ operation request
type CreateSubnetVQRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing the
	// IPv4 address of the subnet.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// SubnetInfoVQ: This is a pointer to a structure of type DHCP_SUBNET_INFO_VQ (section
	// 2.2.1.2.45) that contains information about the IPv4 subnet, including the IPv4 subnet
	// mask and the IPv4 address of the subnet. The structure DHCP_HOST_INFO (section 2.2.1.2.7)
	// (referred by the PrimaryHost member) stored in the SubnetInfoVQ parameter MUST be
	// ignored by both the caller and the server.
	SubnetInfoVQ *dhcpm.SubnetInfoVQ `idl:"name:SubnetInfoVQ;pointer:ref" json:"subnet_info_vq"`
}

func (o *CreateSubnetVQRequest) xxx_ToOp(ctx context.Context) *xxx_CreateSubnetVQOperation {
	if o == nil {
		return &xxx_CreateSubnetVQOperation{}
	}
	return &xxx_CreateSubnetVQOperation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
		SubnetInfoVQ:    o.SubnetInfoVQ,
	}
}

func (o *CreateSubnetVQRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateSubnetVQOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.SubnetInfoVQ = op.SubnetInfoVQ
}
func (o *CreateSubnetVQRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateSubnetVQRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateSubnetVQOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateSubnetVQResponse structure represents the R_DhcpCreateSubnetVQ operation response
type CreateSubnetVQResponse struct {
	// Return: The R_DhcpCreateSubnetVQ return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateSubnetVQResponse) xxx_ToOp(ctx context.Context) *xxx_CreateSubnetVQOperation {
	if o == nil {
		return &xxx_CreateSubnetVQOperation{}
	}
	return &xxx_CreateSubnetVQOperation{
		Return: o.Return,
	}
}

func (o *CreateSubnetVQResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateSubnetVQOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CreateSubnetVQResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateSubnetVQResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateSubnetVQOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubnetInfoVQOperation structure represents the R_DhcpGetSubnetInfoVQ operation
type xxx_GetSubnetInfoVQOperation struct {
	ServerIPAddress string              `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32              `idl:"name:SubnetAddress" json:"subnet_address"`
	SubnetInfoVQ    *dhcpm.SubnetInfoVQ `idl:"name:SubnetInfoVQ" json:"subnet_info_vq"`
	Return          uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubnetInfoVQOperation) OpNum() int { return 49 }

func (o *xxx_GetSubnetInfoVQOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpGetSubnetInfoVQ" }

func (o *xxx_GetSubnetInfoVQOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubnetInfoVQOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubnetInfoVQOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubnetInfoVQOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubnetInfoVQOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SubnetInfoVQ {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SUBNET_INFO_VQ}*(1))(3:{alias=DHCP_SUBNET_INFO_VQ}(struct))
	{
		if o.SubnetInfoVQ != nil {
			_ptr_SubnetInfoVQ := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubnetInfoVQ != nil {
					if err := o.SubnetInfoVQ.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dhcpm.SubnetInfoVQ{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubnetInfoVQ, _ptr_SubnetInfoVQ); err != nil {
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

func (o *xxx_GetSubnetInfoVQOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SubnetInfoVQ {out} (1:{pointer=ref}*(2))(2:{alias=LPDHCP_SUBNET_INFO_VQ,pointer=ref}*(1))(3:{alias=DHCP_SUBNET_INFO_VQ}(struct))
	{
		_ptr_SubnetInfoVQ := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubnetInfoVQ == nil {
				o.SubnetInfoVQ = &dhcpm.SubnetInfoVQ{}
			}
			if err := o.SubnetInfoVQ.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_SubnetInfoVQ := func(ptr interface{}) { o.SubnetInfoVQ = *ptr.(**dhcpm.SubnetInfoVQ) }
		if err := w.ReadPointer(&o.SubnetInfoVQ, _s_SubnetInfoVQ, _ptr_SubnetInfoVQ); err != nil {
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

// GetSubnetInfoVQRequest structure represents the R_DhcpGetSubnetInfoVQ operation request
type GetSubnetInfoVQRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS, containing the IPv4 subnet ID for
	// which the information is retrieved.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
}

func (o *GetSubnetInfoVQRequest) xxx_ToOp(ctx context.Context) *xxx_GetSubnetInfoVQOperation {
	if o == nil {
		return &xxx_GetSubnetInfoVQOperation{}
	}
	return &xxx_GetSubnetInfoVQOperation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
	}
}

func (o *GetSubnetInfoVQRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubnetInfoVQOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
}
func (o *GetSubnetInfoVQRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSubnetInfoVQRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubnetInfoVQOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubnetInfoVQResponse structure represents the R_DhcpGetSubnetInfoVQ operation response
type GetSubnetInfoVQResponse struct {
	// SubnetInfoVQ: This is a pointer of type LPDHCP_SUBNET_INFO_VQ, in which the information
	// for the subnet matching the ID specified by the SubnetAddress parameter is retrieved.
	SubnetInfoVQ *dhcpm.SubnetInfoVQ `idl:"name:SubnetInfoVQ" json:"subnet_info_vq"`
	// Return: The R_DhcpGetSubnetInfoVQ return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetSubnetInfoVQResponse) xxx_ToOp(ctx context.Context) *xxx_GetSubnetInfoVQOperation {
	if o == nil {
		return &xxx_GetSubnetInfoVQOperation{}
	}
	return &xxx_GetSubnetInfoVQOperation{
		SubnetInfoVQ: o.SubnetInfoVQ,
		Return:       o.Return,
	}
}

func (o *GetSubnetInfoVQResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubnetInfoVQOperation) {
	if o == nil {
		return
	}
	o.SubnetInfoVQ = op.SubnetInfoVQ
	o.Return = op.Return
}
func (o *GetSubnetInfoVQResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSubnetInfoVQResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubnetInfoVQOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubnetInfoVQOperation structure represents the R_DhcpSetSubnetInfoVQ operation
type xxx_SetSubnetInfoVQOperation struct {
	ServerIPAddress string              `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	SubnetAddress   uint32              `idl:"name:SubnetAddress" json:"subnet_address"`
	SubnetInfoVQ    *dhcpm.SubnetInfoVQ `idl:"name:SubnetInfoVQ;pointer:ref" json:"subnet_info_vq"`
	Return          uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubnetInfoVQOperation) OpNum() int { return 50 }

func (o *xxx_SetSubnetInfoVQOperation) OpName() string { return "/dhcpsrv/v1/R_DhcpSetSubnetInfoVQ" }

func (o *xxx_SetSubnetInfoVQOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubnetInfoVQOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerIPAddress != "" {
			_ptr_ServerIpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerIPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerIPAddress, _ptr_ServerIpAddress); err != nil {
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
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubnetAddress); err != nil {
			return err
		}
	}
	// SubnetInfoVQ {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_INFO_VQ}*(1))(2:{alias=DHCP_SUBNET_INFO_VQ}(struct))
	{
		if o.SubnetInfoVQ != nil {
			if err := o.SubnetInfoVQ.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dhcpm.SubnetInfoVQ{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubnetInfoVQOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerIpAddress {in} (1:{handle, string, pointer=unique, alias=DHCP_SRV_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerIpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerIPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_ServerIpAddress := func(ptr interface{}) { o.ServerIPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerIPAddress, _s_ServerIpAddress, _ptr_ServerIpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SubnetAddress {in} (1:{alias=DHCP_IP_ADDRESS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubnetAddress); err != nil {
			return err
		}
	}
	// SubnetInfoVQ {in} (1:{pointer=ref, alias=LPDHCP_SUBNET_INFO_VQ}*(1))(2:{alias=DHCP_SUBNET_INFO_VQ}(struct))
	{
		if o.SubnetInfoVQ == nil {
			o.SubnetInfoVQ = &dhcpm.SubnetInfoVQ{}
		}
		if err := o.SubnetInfoVQ.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubnetInfoVQOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubnetInfoVQOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSubnetInfoVQOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSubnetInfoVQRequest structure represents the R_DhcpSetSubnetInfoVQ operation request
type SetSubnetInfoVQRequest struct {
	// ServerIpAddress: The IP address/host name of the DHCP server. This parameter is unused.
	ServerIPAddress string `idl:"name:ServerIpAddress;string;pointer:unique" json:"server_ip_address"`
	// SubnetAddress: This is of type DHCP_IP_ADDRESS, containing the IPv4 subnet ID for
	// which the subnet information is modified.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// SubnetInfoVQ: This is a pointer to a DHCP_SUBNET_INFO_VQ (section 2.2.1.2.45) structure
	// that contains the information about the IPv4 subnet that is modified in the existing
	// IPv4 subnet identified by the SubnetAddress parameter. The structure DHCP_HOST_INFO
	// (section 2.2.1.2.7) (referred by the PrimaryHost member) stored in the SubnetInfoVQ
	// parameter MUST be ignored by both the caller and the server.
	SubnetInfoVQ *dhcpm.SubnetInfoVQ `idl:"name:SubnetInfoVQ;pointer:ref" json:"subnet_info_vq"`
}

func (o *SetSubnetInfoVQRequest) xxx_ToOp(ctx context.Context) *xxx_SetSubnetInfoVQOperation {
	if o == nil {
		return &xxx_SetSubnetInfoVQOperation{}
	}
	return &xxx_SetSubnetInfoVQOperation{
		ServerIPAddress: o.ServerIPAddress,
		SubnetAddress:   o.SubnetAddress,
		SubnetInfoVQ:    o.SubnetInfoVQ,
	}
}

func (o *SetSubnetInfoVQRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubnetInfoVQOperation) {
	if o == nil {
		return
	}
	o.ServerIPAddress = op.ServerIPAddress
	o.SubnetAddress = op.SubnetAddress
	o.SubnetInfoVQ = op.SubnetInfoVQ
}
func (o *SetSubnetInfoVQRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSubnetInfoVQRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubnetInfoVQOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubnetInfoVQResponse structure represents the R_DhcpSetSubnetInfoVQ operation response
type SetSubnetInfoVQResponse struct {
	// Return: The R_DhcpSetSubnetInfoVQ return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetSubnetInfoVQResponse) xxx_ToOp(ctx context.Context) *xxx_SetSubnetInfoVQOperation {
	if o == nil {
		return &xxx_SetSubnetInfoVQOperation{}
	}
	return &xxx_SetSubnetInfoVQOperation{
		Return: o.Return,
	}
}

func (o *SetSubnetInfoVQResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubnetInfoVQOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetSubnetInfoVQResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSubnetInfoVQResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubnetInfoVQOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
