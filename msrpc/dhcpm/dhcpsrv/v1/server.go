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
		in := &CreateSubnetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateSubnet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // R_DhcpSetSubnetInfo
		in := &SetSubnetInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubnetInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // R_DhcpGetSubnetInfo
		in := &GetSubnetInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubnetInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // R_DhcpEnumSubnets
		in := &EnumSubnetsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnets(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // R_DhcpAddSubnetElement
		in := &AddSubnetElementRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddSubnetElement(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // R_DhcpEnumSubnetElements
		in := &EnumSubnetElementsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetElements(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // R_DhcpRemoveSubnetElement
		in := &RemoveSubnetElementRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveSubnetElement(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // R_DhcpDeleteSubnet
		in := &DeleteSubnetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteSubnet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // R_DhcpCreateOption
		in := &CreateOptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateOption(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // R_DhcpSetOptionInfo
		in := &SetOptionInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOptionInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // R_DhcpGetOptionInfo
		in := &GetOptionInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOptionInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // R_DhcpRemoveOption
		in := &RemoveOptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveOption(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // R_DhcpSetOptionValue
		in := &SetOptionValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOptionValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // R_DhcpGetOptionValue
		in := &GetOptionValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOptionValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // R_DhcpEnumOptionValues
		in := &EnumOptionValuesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumOptionValues(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // R_DhcpRemoveOptionValue
		in := &RemoveOptionValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveOptionValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // R_DhcpCreateClientInfo
		in := &CreateClientInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateClientInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // R_DhcpSetClientInfo
		in := &SetClientInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetClientInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // R_DhcpGetClientInfo
		in := &GetClientInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClientInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // R_DhcpDeleteClientInfo
		in := &DeleteClientInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteClientInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // R_DhcpEnumSubnetClients
		in := &EnumSubnetClientsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetClients(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // R_DhcpGetClientOptions
		in := &GetClientOptionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClientOptions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // R_DhcpGetMibInfo
		in := &GetMIBInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMIBInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // R_DhcpEnumOptions
		in := &EnumOptionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumOptions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // R_DhcpSetOptionValues
		in := &SetOptionValuesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOptionValues(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // R_DhcpServerSetConfig
		in := &ServerSetConfigRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerSetConfig(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // R_DhcpServerGetConfig
		in := &ServerGetConfigRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerGetConfig(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // R_DhcpScanDatabase
		in := &ScanDatabaseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ScanDatabase(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // R_DhcpGetVersion
		in := &GetVersionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetVersion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // R_DhcpAddSubnetElementV4
		in := &AddSubnetElementV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddSubnetElementV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // R_DhcpEnumSubnetElementsV4
		in := &EnumSubnetElementsV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetElementsV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // R_DhcpRemoveSubnetElementV4
		in := &RemoveSubnetElementV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveSubnetElementV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // R_DhcpCreateClientInfoV4
		in := &CreateClientInfoV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateClientInfoV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // R_DhcpSetClientInfoV4
		in := &SetClientInfoV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetClientInfoV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // R_DhcpGetClientInfoV4
		in := &GetClientInfoV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClientInfoV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // R_DhcpEnumSubnetClientsV4
		in := &EnumSubnetClientsV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetClientsV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // R_DhcpSetSuperScopeV4
		in := &SetSuperScopeV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSuperScopeV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // R_DhcpGetSuperScopeInfoV4
		in := &GetSuperScopeInfoV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSuperScopeInfoV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // R_DhcpDeleteSuperScopeV4
		in := &DeleteSuperScopeV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteSuperScopeV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // R_DhcpServerSetConfigV4
		in := &ServerSetConfigV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerSetConfigV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // R_DhcpServerGetConfigV4
		in := &ServerGetConfigV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerGetConfigV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // R_DhcpServerSetConfigVQ
		in := &ServerSetConfigVQRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerSetConfigVQ(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // R_DhcpServerGetConfigVQ
		in := &ServerGetConfigVQRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerGetConfigVQ(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // R_DhcpGetMibInfoVQ
		in := &GetMIBInfoVQRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMIBInfoVQ(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // R_DhcpCreateClientInfoVQ
		in := &CreateClientInfoVQRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateClientInfoVQ(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // R_DhcpSetClientInfoVQ
		in := &SetClientInfoVQRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetClientInfoVQ(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // R_DhcpGetClientInfoVQ
		in := &GetClientInfoVQRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClientInfoVQ(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // R_DhcpEnumSubnetClientsVQ
		in := &EnumSubnetClientsVQRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetClientsVQ(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // R_DhcpCreateSubnetVQ
		in := &CreateSubnetVQRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateSubnetVQ(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // R_DhcpGetSubnetInfoVQ
		in := &GetSubnetInfoVQRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubnetInfoVQ(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // R_DhcpSetSubnetInfoVQ
		in := &SetSubnetInfoVQRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubnetInfoVQ(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
