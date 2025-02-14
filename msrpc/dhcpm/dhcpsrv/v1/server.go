package dhcpsrv

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

// dhcpsrv server interface.
type DHCPServerServer interface {

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
	CreateSubnet(context.Context, *CreateSubnetRequest) (*CreateSubnetResponse, error)

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
	SetSubnetInfo(context.Context, *SetSubnetInfoRequest) (*SetSubnetInfoResponse, error)

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
	GetSubnetInfo(context.Context, *GetSubnetInfoRequest) (*GetSubnetInfoResponse, error)

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
	EnumSubnets(context.Context, *EnumSubnetsRequest) (*EnumSubnetsResponse, error)

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
	AddSubnetElement(context.Context, *AddSubnetElementRequest) (*AddSubnetElementResponse, error)

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
	EnumSubnetElements(context.Context, *EnumSubnetElementsRequest) (*EnumSubnetElementsResponse, error)

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
	RemoveSubnetElement(context.Context, *RemoveSubnetElementRequest) (*RemoveSubnetElementResponse, error)

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
	DeleteSubnet(context.Context, *DeleteSubnetRequest) (*DeleteSubnetResponse, error)

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
	CreateOption(context.Context, *CreateOptionRequest) (*CreateOptionResponse, error)

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
	SetOptionInfo(context.Context, *SetOptionInfoRequest) (*SetOptionInfoResponse, error)

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
	GetOptionInfo(context.Context, *GetOptionInfoRequest) (*GetOptionInfoResponse, error)

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
	RemoveOption(context.Context, *RemoveOptionRequest) (*RemoveOptionResponse, error)

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
	SetOptionValue(context.Context, *SetOptionValueRequest) (*SetOptionValueResponse, error)

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
	GetOptionValue(context.Context, *GetOptionValueRequest) (*GetOptionValueResponse, error)

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
	EnumOptionValues(context.Context, *EnumOptionValuesRequest) (*EnumOptionValuesResponse, error)

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
	RemoveOptionValue(context.Context, *RemoveOptionValueRequest) (*RemoveOptionValueResponse, error)

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
	CreateClientInfo(context.Context, *CreateClientInfoRequest) (*CreateClientInfoResponse, error)

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
	SetClientInfo(context.Context, *SetClientInfoRequest) (*SetClientInfoResponse, error)

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
	GetClientInfo(context.Context, *GetClientInfoRequest) (*GetClientInfoResponse, error)

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
	DeleteClientInfo(context.Context, *DeleteClientInfoRequest) (*DeleteClientInfoResponse, error)

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
	EnumSubnetClients(context.Context, *EnumSubnetClientsRequest) (*EnumSubnetClientsResponse, error)

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
	GetClientOptions(context.Context, *GetClientOptionsRequest) (*GetClientOptionsResponse, error)

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
	GetMIBInfo(context.Context, *GetMIBInfoRequest) (*GetMIBInfoResponse, error)

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
	EnumOptions(context.Context, *EnumOptionsRequest) (*EnumOptionsResponse, error)

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
	SetOptionValues(context.Context, *SetOptionValuesRequest) (*SetOptionValuesResponse, error)

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
	ServerSetConfig(context.Context, *ServerSetConfigRequest) (*ServerSetConfigResponse, error)

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
	ServerGetConfig(context.Context, *ServerGetConfigRequest) (*ServerGetConfigResponse, error)

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
	ScanDatabase(context.Context, *ScanDatabaseRequest) (*ScanDatabaseResponse, error)

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
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)

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
	AddSubnetElementV4(context.Context, *AddSubnetElementV4Request) (*AddSubnetElementV4Response, error)

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
	EnumSubnetElementsV4(context.Context, *EnumSubnetElementsV4Request) (*EnumSubnetElementsV4Response, error)

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
	RemoveSubnetElementV4(context.Context, *RemoveSubnetElementV4Request) (*RemoveSubnetElementV4Response, error)

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
	CreateClientInfoV4(context.Context, *CreateClientInfoV4Request) (*CreateClientInfoV4Response, error)

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
	SetClientInfoV4(context.Context, *SetClientInfoV4Request) (*SetClientInfoV4Response, error)

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
	GetClientInfoV4(context.Context, *GetClientInfoV4Request) (*GetClientInfoV4Response, error)

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
	EnumSubnetClientsV4(context.Context, *EnumSubnetClientsV4Request) (*EnumSubnetClientsV4Response, error)

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
	SetSuperScopeV4(context.Context, *SetSuperScopeV4Request) (*SetSuperScopeV4Response, error)

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
	GetSuperScopeInfoV4(context.Context, *GetSuperScopeInfoV4Request) (*GetSuperScopeInfoV4Response, error)

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
	DeleteSuperScopeV4(context.Context, *DeleteSuperScopeV4Request) (*DeleteSuperScopeV4Response, error)

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
	ServerSetConfigV4(context.Context, *ServerSetConfigV4Request) (*ServerSetConfigV4Response, error)

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
	ServerGetConfigV4(context.Context, *ServerGetConfigV4Request) (*ServerGetConfigV4Response, error)

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
	ServerSetConfigVQ(context.Context, *ServerSetConfigVQRequest) (*ServerSetConfigVQResponse, error)

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
	ServerGetConfigVQ(context.Context, *ServerGetConfigVQRequest) (*ServerGetConfigVQResponse, error)

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
	GetMIBInfoVQ(context.Context, *GetMIBInfoVQRequest) (*GetMIBInfoVQResponse, error)

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
	CreateClientInfoVQ(context.Context, *CreateClientInfoVQRequest) (*CreateClientInfoVQResponse, error)

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
	SetClientInfoVQ(context.Context, *SetClientInfoVQRequest) (*SetClientInfoVQResponse, error)

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
	GetClientInfoVQ(context.Context, *GetClientInfoVQRequest) (*GetClientInfoVQResponse, error)

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
	EnumSubnetClientsVQ(context.Context, *EnumSubnetClientsVQRequest) (*EnumSubnetClientsVQResponse, error)

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
	CreateSubnetVQ(context.Context, *CreateSubnetVQRequest) (*CreateSubnetVQResponse, error)

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
	GetSubnetInfoVQ(context.Context, *GetSubnetInfoVQRequest) (*GetSubnetInfoVQResponse, error)

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
	SetSubnetInfoVQ(context.Context, *SetSubnetInfoVQRequest) (*SetSubnetInfoVQResponse, error)
}

func RegisterDHCPServerServer(conn dcerpc.Conn, o DHCPServerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDHCPServerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DHCPServerSyntaxV1_0))...)
}

func NewDHCPServerServerHandle(o DHCPServerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DHCPServerServerHandle(ctx, o, opNum, r)
	}
}

func DHCPServerServerHandle(ctx context.Context, o DHCPServerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // R_DhcpCreateSubnet
		op := &xxx_CreateSubnetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateSubnetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateSubnet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // R_DhcpSetSubnetInfo
		op := &xxx_SetSubnetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubnetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubnetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // R_DhcpGetSubnetInfo
		op := &xxx_GetSubnetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubnetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubnetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // R_DhcpEnumSubnets
		op := &xxx_EnumSubnetsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumSubnetsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumSubnets(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // R_DhcpAddSubnetElement
		op := &xxx_AddSubnetElementOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddSubnetElementRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddSubnetElement(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // R_DhcpEnumSubnetElements
		op := &xxx_EnumSubnetElementsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumSubnetElementsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumSubnetElements(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // R_DhcpRemoveSubnetElement
		op := &xxx_RemoveSubnetElementOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveSubnetElementRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveSubnetElement(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // R_DhcpDeleteSubnet
		op := &xxx_DeleteSubnetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteSubnetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteSubnet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // R_DhcpCreateOption
		op := &xxx_CreateOptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateOptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateOption(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // R_DhcpSetOptionInfo
		op := &xxx_SetOptionInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOptionInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOptionInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // R_DhcpGetOptionInfo
		op := &xxx_GetOptionInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOptionInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOptionInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // R_DhcpRemoveOption
		op := &xxx_RemoveOptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveOptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveOption(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // R_DhcpSetOptionValue
		op := &xxx_SetOptionValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOptionValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOptionValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // R_DhcpGetOptionValue
		op := &xxx_GetOptionValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOptionValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOptionValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // R_DhcpEnumOptionValues
		op := &xxx_EnumOptionValuesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumOptionValuesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumOptionValues(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // R_DhcpRemoveOptionValue
		op := &xxx_RemoveOptionValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveOptionValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveOptionValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // R_DhcpCreateClientInfo
		op := &xxx_CreateClientInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateClientInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateClientInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // R_DhcpSetClientInfo
		op := &xxx_SetClientInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // R_DhcpGetClientInfo
		op := &xxx_GetClientInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // R_DhcpDeleteClientInfo
		op := &xxx_DeleteClientInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteClientInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteClientInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // R_DhcpEnumSubnetClients
		op := &xxx_EnumSubnetClientsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumSubnetClientsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumSubnetClients(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // R_DhcpGetClientOptions
		op := &xxx_GetClientOptionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientOptionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientOptions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // R_DhcpGetMibInfo
		op := &xxx_GetMIBInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMIBInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMIBInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // R_DhcpEnumOptions
		op := &xxx_EnumOptionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumOptionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumOptions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // R_DhcpSetOptionValues
		op := &xxx_SetOptionValuesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOptionValuesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOptionValues(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // R_DhcpServerSetConfig
		op := &xxx_ServerSetConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerSetConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerSetConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // R_DhcpServerGetConfig
		op := &xxx_ServerGetConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerGetConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerGetConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // R_DhcpScanDatabase
		op := &xxx_ScanDatabaseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ScanDatabaseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ScanDatabase(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // R_DhcpGetVersion
		op := &xxx_GetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // R_DhcpAddSubnetElementV4
		op := &xxx_AddSubnetElementV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddSubnetElementV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddSubnetElementV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // R_DhcpEnumSubnetElementsV4
		op := &xxx_EnumSubnetElementsV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumSubnetElementsV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumSubnetElementsV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // R_DhcpRemoveSubnetElementV4
		op := &xxx_RemoveSubnetElementV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveSubnetElementV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveSubnetElementV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // R_DhcpCreateClientInfoV4
		op := &xxx_CreateClientInfoV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateClientInfoV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateClientInfoV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // R_DhcpSetClientInfoV4
		op := &xxx_SetClientInfoV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientInfoV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientInfoV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // R_DhcpGetClientInfoV4
		op := &xxx_GetClientInfoV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientInfoV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientInfoV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // R_DhcpEnumSubnetClientsV4
		op := &xxx_EnumSubnetClientsV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumSubnetClientsV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumSubnetClientsV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // R_DhcpSetSuperScopeV4
		op := &xxx_SetSuperScopeV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSuperScopeV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSuperScopeV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // R_DhcpGetSuperScopeInfoV4
		op := &xxx_GetSuperScopeInfoV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSuperScopeInfoV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSuperScopeInfoV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // R_DhcpDeleteSuperScopeV4
		op := &xxx_DeleteSuperScopeV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteSuperScopeV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteSuperScopeV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // R_DhcpServerSetConfigV4
		op := &xxx_ServerSetConfigV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerSetConfigV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerSetConfigV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // R_DhcpServerGetConfigV4
		op := &xxx_ServerGetConfigV4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerGetConfigV4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerGetConfigV4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // R_DhcpServerSetConfigVQ
		op := &xxx_ServerSetConfigVQOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerSetConfigVQRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerSetConfigVQ(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // R_DhcpServerGetConfigVQ
		op := &xxx_ServerGetConfigVQOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerGetConfigVQRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerGetConfigVQ(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // R_DhcpGetMibInfoVQ
		op := &xxx_GetMIBInfoVQOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMIBInfoVQRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMIBInfoVQ(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // R_DhcpCreateClientInfoVQ
		op := &xxx_CreateClientInfoVQOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateClientInfoVQRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateClientInfoVQ(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // R_DhcpSetClientInfoVQ
		op := &xxx_SetClientInfoVQOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientInfoVQRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientInfoVQ(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // R_DhcpGetClientInfoVQ
		op := &xxx_GetClientInfoVQOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientInfoVQRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientInfoVQ(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // R_DhcpEnumSubnetClientsVQ
		op := &xxx_EnumSubnetClientsVQOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumSubnetClientsVQRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumSubnetClientsVQ(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // R_DhcpCreateSubnetVQ
		op := &xxx_CreateSubnetVQOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateSubnetVQRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateSubnetVQ(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // R_DhcpGetSubnetInfoVQ
		op := &xxx_GetSubnetInfoVQOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSubnetInfoVQRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSubnetInfoVQ(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // R_DhcpSetSubnetInfoVQ
		op := &xxx_SetSubnetInfoVQOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSubnetInfoVQRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSubnetInfoVQ(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented dhcpsrv
type UnimplementedDHCPServerServer struct {
}

func (UnimplementedDHCPServerServer) CreateSubnet(context.Context, *CreateSubnetRequest) (*CreateSubnetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) SetSubnetInfo(context.Context, *SetSubnetInfoRequest) (*SetSubnetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetSubnetInfo(context.Context, *GetSubnetInfoRequest) (*GetSubnetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) EnumSubnets(context.Context, *EnumSubnetsRequest) (*EnumSubnetsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) AddSubnetElement(context.Context, *AddSubnetElementRequest) (*AddSubnetElementResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) EnumSubnetElements(context.Context, *EnumSubnetElementsRequest) (*EnumSubnetElementsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) RemoveSubnetElement(context.Context, *RemoveSubnetElementRequest) (*RemoveSubnetElementResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) DeleteSubnet(context.Context, *DeleteSubnetRequest) (*DeleteSubnetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) CreateOption(context.Context, *CreateOptionRequest) (*CreateOptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) SetOptionInfo(context.Context, *SetOptionInfoRequest) (*SetOptionInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetOptionInfo(context.Context, *GetOptionInfoRequest) (*GetOptionInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) RemoveOption(context.Context, *RemoveOptionRequest) (*RemoveOptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) SetOptionValue(context.Context, *SetOptionValueRequest) (*SetOptionValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetOptionValue(context.Context, *GetOptionValueRequest) (*GetOptionValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) EnumOptionValues(context.Context, *EnumOptionValuesRequest) (*EnumOptionValuesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) RemoveOptionValue(context.Context, *RemoveOptionValueRequest) (*RemoveOptionValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) CreateClientInfo(context.Context, *CreateClientInfoRequest) (*CreateClientInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) SetClientInfo(context.Context, *SetClientInfoRequest) (*SetClientInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetClientInfo(context.Context, *GetClientInfoRequest) (*GetClientInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) DeleteClientInfo(context.Context, *DeleteClientInfoRequest) (*DeleteClientInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) EnumSubnetClients(context.Context, *EnumSubnetClientsRequest) (*EnumSubnetClientsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetClientOptions(context.Context, *GetClientOptionsRequest) (*GetClientOptionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetMIBInfo(context.Context, *GetMIBInfoRequest) (*GetMIBInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) EnumOptions(context.Context, *EnumOptionsRequest) (*EnumOptionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) SetOptionValues(context.Context, *SetOptionValuesRequest) (*SetOptionValuesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) ServerSetConfig(context.Context, *ServerSetConfigRequest) (*ServerSetConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) ServerGetConfig(context.Context, *ServerGetConfigRequest) (*ServerGetConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) ScanDatabase(context.Context, *ScanDatabaseRequest) (*ScanDatabaseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) AddSubnetElementV4(context.Context, *AddSubnetElementV4Request) (*AddSubnetElementV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) EnumSubnetElementsV4(context.Context, *EnumSubnetElementsV4Request) (*EnumSubnetElementsV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) RemoveSubnetElementV4(context.Context, *RemoveSubnetElementV4Request) (*RemoveSubnetElementV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) CreateClientInfoV4(context.Context, *CreateClientInfoV4Request) (*CreateClientInfoV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) SetClientInfoV4(context.Context, *SetClientInfoV4Request) (*SetClientInfoV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetClientInfoV4(context.Context, *GetClientInfoV4Request) (*GetClientInfoV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) EnumSubnetClientsV4(context.Context, *EnumSubnetClientsV4Request) (*EnumSubnetClientsV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) SetSuperScopeV4(context.Context, *SetSuperScopeV4Request) (*SetSuperScopeV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetSuperScopeInfoV4(context.Context, *GetSuperScopeInfoV4Request) (*GetSuperScopeInfoV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) DeleteSuperScopeV4(context.Context, *DeleteSuperScopeV4Request) (*DeleteSuperScopeV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) ServerSetConfigV4(context.Context, *ServerSetConfigV4Request) (*ServerSetConfigV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) ServerGetConfigV4(context.Context, *ServerGetConfigV4Request) (*ServerGetConfigV4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) ServerSetConfigVQ(context.Context, *ServerSetConfigVQRequest) (*ServerSetConfigVQResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) ServerGetConfigVQ(context.Context, *ServerGetConfigVQRequest) (*ServerGetConfigVQResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetMIBInfoVQ(context.Context, *GetMIBInfoVQRequest) (*GetMIBInfoVQResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) CreateClientInfoVQ(context.Context, *CreateClientInfoVQRequest) (*CreateClientInfoVQResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) SetClientInfoVQ(context.Context, *SetClientInfoVQRequest) (*SetClientInfoVQResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetClientInfoVQ(context.Context, *GetClientInfoVQRequest) (*GetClientInfoVQResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) EnumSubnetClientsVQ(context.Context, *EnumSubnetClientsVQRequest) (*EnumSubnetClientsVQResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) CreateSubnetVQ(context.Context, *CreateSubnetVQRequest) (*CreateSubnetVQResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) GetSubnetInfoVQ(context.Context, *GetSubnetInfoVQRequest) (*GetSubnetInfoVQResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDHCPServerServer) SetSubnetInfoVQ(context.Context, *SetSubnetInfoVQRequest) (*SetSubnetInfoVQResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DHCPServerServer = (*UnimplementedDHCPServerServer)(nil)
