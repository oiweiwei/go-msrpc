package remotefw

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

// RemoteFW server interface.
type RemoteFwServer interface {

	// The RRPC_FWOpenPolicyStore method requests the server to open a specified policy
	// store. The store can be opened for reading or for editing the firewall policy. The
	// method also returns a handle to the opened store with which the client can then perform
	// operations on this policy store. The server allocates a PolicyStoreConnection object
	// to track the policy store type and the binary version associated with the handle.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	OpenPolicyStore(context.Context, *OpenPolicyStoreRequest) (*OpenPolicyStoreResponse, error)

	// The RRPC_FWClosePolicyStore method receives an opened store handle, closes it, and
	// deallocates the corresponding PolicyStoreConnection object.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	ClosePolicyStore(context.Context, *ClosePolicyStoreRequest) (*ClosePolicyStoreResponse, error)

	// The RRPC_FWRestoreDefaults method replaces the contents of LocalStore with the contents
	// of DefaultsStore.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	RestoreDefaults(context.Context, *RestoreDefaultsRequest) (*RestoreDefaultsResponse, error)

	// The RRPC_FWGetGlobalConfig method retrieves the value of a global policy configuration
	// option. The client specifies to the server from what store this value MUST be retrieved
	// and in what specific configuration option it is interested.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specific configuration option is not found within the policy. This means     |
	//	|                                    | that it is not configured. If the option is not configured in any other store,   |
	//	|                                    | the firewall uses a default value.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The store type specified does not support this method.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The buffer is not big enough to hold the configuration option value.             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The specific configuration option |
	//	|                                    | is not meant to be available in the specified store. The specified configuration |
	//	|                                    | option is not defined. One of the required values is not specified. The buffer   |
	//	|                                    | size is not enough to hold the specific value.                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	GetGlobalConfig(context.Context, *GetGlobalConfigRequest) (*GetGlobalConfigResponse, error)

	// The RRPC_FWSetGlobalConfig method modifies the value of a global policy configuration
	// option. The client specifies to the server in what store this value MUST be written
	// and what specific configuration option it is interested in modifying.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The store type specified does not support this method.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The specific configuration option |
	//	|                                    | is not meant to be available in the specified store. The specified configuration |
	//	|                                    | option is not defined. One of the required values is not specified. The buffer   |
	//	|                                    | is null but dwBufSize says otherwise. The buffer size is not enough to hold the  |
	//	|                                    | specific value.                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method performs a merge operation of the resultant configuration values, as
	// defined in section 3.1.3. It then determines what modifications are necessary on
	// the rule objects to make sure the policy is enforced.
	SetGlobalConfig(context.Context, *SetGlobalConfigRequest) (*SetGlobalConfigResponse, error)

	// The RRPC_FWAddFirewallRule method requests the server to add the specified firewall
	// rule in the policy contained in the policy store that is referenced by the handle
	// specified in the hPolicyStore parameter.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter of this method is incorrect, or is required and not specified. This  |
	//	|                                    | error can be returned because: The pRule object did not pass the firewall rule   |
	//	|                                    | validations that are specified in the definition of the FW_RULE data type. One   |
	//	|                                    | of the required values is not specified. A policy store does not support rules   |
	//	|                                    | with profile conditions other than ALL profiles. The wszLocalApplication field   |
	//	|                                    | of the rule contains a string that was determined to be an invalid path.<33>     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule to the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule in disk. If
	// called on an online store, the firewall rule is also enforced.
	AddFirewallRule(context.Context, *AddFirewallRuleRequest) (*AddFirewallRuleResponse, error)

	// The RRPC_FWSetFirewallRule method requests the server to modify the specified firewall
	// rule in the policy contained in the policy store that is referenced by the handle
	// specified in the hPolicyStore parameter.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule that is referenced by the wszRuleID member string of the      |
	//	|                                    | FW_RULE data type is not found in the policy store.                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter of this method is incorrect, or is required and not specified. This  |
	//	|                                    | error can be returned because: The pRule object did not pass the firewall rule   |
	//	|                                    | validations that are specified in the definition of the FW_RULE data type. One   |
	//	|                                    | of the required values is not specified. A policy store does not support rules   |
	//	|                                    | that have profile conditions other than ALL profiles. The wszLocalApplication    |
	//	|                                    | field of the rule contains a string that was determined to be an invalid path.   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetFirewallRule(context.Context, *SetFirewallRuleRequest) (*SetFirewallRuleResponse, error)

	// The RRPC_FWDeleteFirewallRule method requests the server to delete the specified
	// firewall rule in the policy contained in the policy store referenced by the handle
	// specified in the hPolicyStore parameter.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following are common.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED  | The specified store does not support this method; the store might be read-only.  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED  | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                 | is also returned if the client does not have the required credentials to call    |
	//	|                                 | the method.                                                                      |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND | The specified rule referenced by the wszRuleID member string of the FW_RULE data |
	//	|                                 | type is not found in the policy store.                                           |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method deletes a firewall rule already stored in the firewall linked list of
	// the memory representation of the store being modified. It uses this list to determine
	// if the rule exists or not. It also writes through and deletes the rule from disk.
	// If called on an online store, the removal of the firewall rule is also enforced.
	DeleteFirewallRule(context.Context, *DeleteFirewallRuleRequest) (*DeleteFirewallRuleResponse, error)

	// The RRPC_FWDeleteAllFirewallRules method deletes all firewall rules in the firewall
	// linked list of the memory representation of the store being modified. It also writes
	// through and deletes all rules from the disk representation. If called on an online
	// store, no firewall rules are enforced after the method returns.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following are common.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED | The specified store does not support this method; the store might be read-only.  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                | is also returned if the client does not have the required credentials to call    |
	//	|                                | the method.                                                                      |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	DeleteAllFirewallRules(context.Context, *DeleteAllFirewallRulesRequest) (*DeleteAllFirewallRulesResponse, error)

	// The RRPC_FWEnumFirewallRules method requests the server to return all the firewall
	// rules contained in the store that is referenced by the hPolicyStore handle. The method
	// returns a linked list of all the firewall rule objects.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The dwProfileFilter parameter contains profiles that are not valid.              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumFirewallRules(context.Context, *EnumFirewallRulesRequest) (*EnumFirewallRulesResponse, error)

	// The RRPC_FWGetConfig method retrieves the value of a profile configuration option.
	// The client specifies to the server from what store and profile this value MUST be
	// retrieved and in what specific configuration option it is interested.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specific configuration option is not found within the policy. This means     |
	//	|                                    | that it is not configured. If the option is not configured in any other store,   |
	//	|                                    | the firewall uses a default value.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The method does not support the specified combination of parameters. This        |
	//	|                                    | can be because: The store type specified does not support this method. The       |
	//	|                                    | configuration option is not supported in this store. The Profile parameter       |
	//	|                                    | contains a combination of profiles (instead of a single profile) or an unknown   |
	//	|                                    | profile.                                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The buffer is not big enough to hold the configuration option value.             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The specified configuration       |
	//	|                                    | option is not defined. One of the required values is not specified.              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)

	// The RRPC_FWSetConfig method modifies the value of a profile configuration option.
	// The client specifies to the server in what store and profile this value MUST be written
	// and what specific configuration option it is interested in modifying.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The method does not support the specified combination of parameters. This can    |
	//	|                                    | be because: The store type specified does not support this method. The Profile   |
	//	|                                    | parameter contains a combination of profiles (instead of a single profile) or an |
	//	|                                    | unknown profile.                                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The specific configuration        |
	//	|                                    | option is not meant to be available in the specified store. The specified        |
	//	|                                    | configuration option is not defined. The size of the buffer does not match the   |
	//	|                                    | size of the type of the configuration value. The buffer is null but dwBufSize    |
	//	|                                    | says otherwise. The caller wants to set a FW_PROFILE_CONFIG_LOG_MAX_FILE_SIZE    |
	//	|                                    | that is not within the valid values [min, max]. The default action               |
	//	|                                    | configuration value specifies a value that maps to neither allow nor block.      |
	//	|                                    | The FW_PROFILE_CONFIG_LOG_FILE_PATH configuration value contains the following   |
	//	|                                    | invalid characters: /,*,?,",<,>,|.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method performs a merge operation of the resultant configuration values, as
	// defined in section 3.1.3. It then determines what modifications are necessary on
	// the rule objects (for example, remove rule enforcement if firewall is off) to make
	// sure the policy is enforced.
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)

	// The RRPC_FWAddConnectionSecurityRule method requests the server to add the connection
	// security rule in the policy contained in the policy store that is referenced by the
	// specified opened policy store handle.
	//
	// Return Values: This method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter of this method is incorrect, or is required and not specified.       |
	//	|                                    | This error can be returned because: The pRule object did not pass the            |
	//	|                                    | connection security rule validations specified in the definition of the          |
	//	|                                    | FW_CS_RULE data type. The rule has a phase 2 crypto set that specified           |
	//	|                                    | FW_CRYPTO_PRPTOCOL_AUTH_NO_ENCAP (see section 2.2.69), and it is a tunnel mode   |
	//	|                                    | rule, or it also has an AuthSet structure (section 2.2.65) that specifies a      |
	//	|                                    | preshared key auth method. A required value is not specified.                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a connection security rule in the connection security link list
	// of the memory representation of the store being modified. It also writes through
	// and saves the rule to disk. If called on an online store, the connection security
	// rule is also enforced.
	AddConnectionSecurityRule(context.Context, *AddConnectionSecurityRuleRequest) (*AddConnectionSecurityRuleResponse, error)

	// The RRPC_FWSetConnectionSecurityRule method requests the server to modify the specified
	// connection security rule in the policy contained in the policy store that is referenced
	// by the handle specified in the hPolicy parameter.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. This error is   |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule that is referenced by the wszRuleID member string of the      |
	//	|                                    | FW_CS_RULE data type is not found in the policy store.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter of this method is incorrect, or is required and not specified.       |
	//	|                                    | This error can be returned because: The pRule object did not pass the            |
	//	|                                    | connection security rule validations that are specified in the definition of     |
	//	|                                    | the FW_CS_RULE data type. The rule has a phase 2 crypto set that specified       |
	//	|                                    | FW_CRYPTO_PRPTOCOL_AUTH_NO_ENCAP (see section 2.2.69), and either it is a tunnel |
	//	|                                    | mode rule or it has an AuthSet that specifies a preshared key auth method. A     |
	//	|                                    | required value is not specified.                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method modifies a connection security rule already stored in the connection
	// security linked list of the memory representation of the store being modified. It
	// uses this list to determine whether the rule exists. It also writes through and saves
	// the rule in disk. If called on an online store, the connection security rule modifications
	// are also enforced.
	SetConnectionSecurityRule(context.Context, *SetConnectionSecurityRuleRequest) (*SetConnectionSecurityRuleResponse, error)

	// The RRPC_FWDeleteConnectionSecurityRule method requests the server to delete the
	// specified connection security rule in the policy contained in the policy store referenced
	// by the handle specified in the hPolicy parameter.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED  | The specified store does not support this method; the store might be read-only.  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED  | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                 | also returned if the client does not have the required credentials to call the   |
	//	|                                 | method.                                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND | The specified rule referenced by the pRuleId member string is not found in the   |
	//	|                                 | policy store.                                                                    |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method deletes a connection security rule already stored in the connection security
	// linked list of the memory representation of the store being modified. It uses this
	// list to determine if the rule exists or not. It also writes through and deletes the
	// rule from disk. If called on an online store, the removal of the connection security
	// rule is also enforced.
	DeleteConnectionSecurityRule(context.Context, *DeleteConnectionSecurityRuleRequest) (*DeleteConnectionSecurityRuleResponse, error)

	// The RRPC_FWDeleteAllConnectionSecurityRules method requests the server to delete
	// all the connection security rules in the policy contained in the policy store referenced
	// by the handle specified in the hPolicy parameter.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED | The specified store does not support this method; the store might be read-only.  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                | also returned if the client does not have the required credentials to call the   |
	//	|                                | method.                                                                          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method deletes all connection security rules in the connection security linked
	// list of the memory representation of the store being modified. It also writes through
	// and deletes all rules from the disk representation. If called on an online store,
	// no connection security rules are enforced after the method returns.
	DeleteAllConnectionSecurityRules(context.Context, *DeleteAllConnectionSecurityRulesRequest) (*DeleteAllConnectionSecurityRulesResponse, error)

	// The RRPC_FWEnumConnectionSecurityRules method requests the server to return all the
	// connection security rules contained in the store that is referenced by the hPolicy
	// handle. The method returns a linked list of all the connection security rule objects.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The dwProfileFilter parameter contains invalid profiles.              |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumConnectionSecurityRules(context.Context, *EnumConnectionSecurityRulesRequest) (*EnumConnectionSecurityRulesResponse, error)

	// The RRPC_FWAddAuthenticationSet method requests the server to add the authentication
	// set in the policy contained in the policy store referenced by the handle specified
	// in the hPolicy parameter.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified set has a set ID that already exists in the specified store.       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The pAuth object did not pass the |
	//	|                                    | authentication set validations specified in the definition of the FW_AUTH_SET    |
	//	|                                    | data type. One of the required values is not specified.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds an authentication set in the authentication linked list of the memory
	// representation of the store being modified. It also writes through and saves the
	// set in disk. If called on an online store and the set is a primary set, the method
	// enumerates the connection security rule list and reapplies each rule referencing
	// this primary set to complete the enforcement of the policy.
	AddAuthenticationSet(context.Context, *AddAuthenticationSetRequest) (*AddAuthenticationSetResponse, error)

	// The RRPC_FWSetAuthenticationSet method requests the server to modify the specified
	// authentication set in the policy contained in the policy store referenced by the
	// handle specified in the hPolicy parameter.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified set referenced by the wszSetId member string of the FW_AUTH_SET    |
	//	|                                    | data type is not found in the policy store.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The pAuth object did not pass the |
	//	|                                    | authentication set validations specified in the definition of the FW_AUTH_SET    |
	//	|                                    | data type. One of the required values is not specified.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method modifies an authentication set in the authentication linked list of the
	// memory representation of the store being modified. It also writes through and saves
	// the set in disk. If called on an online store, the method enumerates the connection
	// security rules list and reapplies each rule referencing this primary set to complete
	// the enforcement of the policy.
	SetAuthenticationSet(context.Context, *SetAuthenticationSetRequest) (*SetAuthenticationSetResponse, error)

	// The RRPC_FWDeleteAuthenticationSet method requests the server to delete the specified
	// authentication set in the policy contained in the policy store referenced by the
	// handle specified in the hPolicy parameter.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000962 ERROR_ACTIVE_CONNECTIONS | The specified set is still referenced by connection security rules. This failure |
	//	|                                     | happens only when the set is not a primary set. There is always a primary set to |
	//	|                                     | use, either from other stores or a hard-coded one.                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED      | The specified store does not support this method; the store might be read-only.  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED      | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                     | also returned if the client does not have the required credentials to call the   |
	//	|                                     | method.                                                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The specified rule referenced by the wszSetId string is not found in the policy  |
	//	|                                     | store.                                                                           |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER  | The specified IPsec phase is not a valid one.                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method deletes an authentication set in the authentication linked list of the
	// memory representation of the store being modified. It also writes through and saves
	// the set in disk. If called on an online store, and the set is not a primary set,
	// the method does not delete the specified set if any connection rule references this
	// set.
	DeleteAuthenticationSet(context.Context, *DeleteAuthenticationSetRequest) (*DeleteAuthenticationSetResponse, error)

	// The RRPC_FWDeleteAllAuthenticationSets method requests the server to delete all the
	// authentication sets of a specific IPsec phase in the policy contained in the policy
	// store referenced by the handle specified in the hPolicy parameter.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000962 ERROR_ACTIVE_CONNECTIONS | The specified set is still referenced by connection security rules. This failure |
	//	|                                     | happens only when the set is not a primary set. There is always a primary set to |
	//	|                                     | use, either from other stores or a hard-coded one.                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED      | The specified store does not support this method; the store might be read-only.  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED      | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                     | also returned if the client does not have the required credentials to call the   |
	//	|                                     | method.                                                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The specified rule referenced by the wszSetId string is not found in the policy  |
	//	|                                     | store.                                                                           |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER  | The specified IPsec phase is not a valid one.                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method deletes all the authentication sets in the authentication linked list
	// of the memory representation of the store being modified. It also writes through
	// and deletes the sets from disk. If called on an online store, the method does not
	// delete the sets if any nonprimary set is referenced by a connection security rule.
	DeleteAllAuthenticationSets(context.Context, *DeleteAllAuthenticationSetsRequest) (*DeleteAllAuthenticationSetsResponse, error)

	// The RRPC_FWEnumAuthenticationSets method requests the server to return all the authentication
	// sets of the specified IPsec phase contained in the store referenced by the hPolicy
	// handle. The method returns a linked list of these objects.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The IpSecPhase parameter          |
	//	|                                    | specifies an invalid IPsec negotiation phase. One of the required values is not  |
	//	|                                    | specified.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumAuthenticationSets(context.Context, *EnumAuthenticationSetsRequest) (*EnumAuthenticationSetsResponse, error)

	// The RRPC_FWAddCryptoSet method adds a cryptographic set in the cryptographic linked
	// list of the memory representation of the store being modified. It also writes through
	// and saves the set to the disk. If called on an online store, and the set is a primary
	// set, the method enumerates the connection security rule list and reapplies each rule
	// referencing this primary set to complete the enforcement of the policy.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The pCrypto object did not        |
	//	|                                    | pass the cryptographic set validations specified in the definition of the        |
	//	|                                    | FW_CRYPTO_SET data type. One of the required values is not specified.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	AddCryptoSet(context.Context, *AddCryptoSetRequest) (*AddCryptoSetResponse, error)

	// The RRPC_FWSetCryptoSet method requests the server to modify the specified cryptographic
	// set in the policy contained in the policy store referenced by the handle specified
	// in the hPolicy parameter.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified set referenced by the wszSetId member string of the FW_CRYPTO_SET  |
	//	|                                    | data type is not found in the policy store.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The pCrypto object did not        |
	//	|                                    | pass the cryptographic set validations specified in the definition of the        |
	//	|                                    | FW_CRYPTO_SET data type. One of the required values is not specified.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method modifies a cryptographic set in the cryptographic linked list of the
	// memory representation of the store being modified. It also writes through and saves
	// the set to the disk. If called on an online store, the method enumerates the connection
	// security rules list and reapplies each rule referencing this primary set to complete
	// the enforcement of the policy.
	SetCryptoSet(context.Context, *SetCryptoSetRequest) (*SetCryptoSetResponse, error)

	// The RRPC_FWDeleteCryptoSet method requests the server to delete the specified cryptographic
	// set in the policy contained in the policy store that is referenced by the handle
	// specified in the hPolicy parameter.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000962 ERROR_ACTIVE_CONNECTIONS | The specified set is still referenced by connection security or main mode rules. |
	//	|                                     | This failure happens only when the set is not a primary set. There is always a   |
	//	|                                     | primary set to use, either from other stores or a hard-coded one.                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED      | The specified store does not support this method; the store might be read-only.  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED      | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                     | also returned if the client does not have the required credentials to call the   |
	//	|                                     | method.                                                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The specified rule that is referenced by the wszSetId string is not found in the |
	//	|                                     | policy store.                                                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER  | The specified IPsec phase is not a valid one.                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method deletes a cryptographic set in the cryptographic linked list of the memory
	// representation of the store being modified. It also writes through and saves the
	// set to disk. If called on an online store and the set is not a primary set, the method
	// does not delete the specified set if any connection rule references this set.
	DeleteCryptoSet(context.Context, *DeleteCryptoSetRequest) (*DeleteCryptoSetResponse, error)

	// The RRPC_FWDeleteAllCryptoSets method requests the server to delete all the cryptographic
	// sets of a specific IPsec phase in the policy contained in the policy store that is
	// referenced by the handle specified in the hPolicy parameter.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000962 ERROR_ACTIVE_CONNECTIONS | There are nonprimary sets still being referenced by connection security or main  |
	//	|                                     | mode rules. There is always a primary set to use, either from other stores or a  |
	//	|                                     | hard-coded one; therefore, this failure never occurs because of primary sets.    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED      | The specified store does not support this method; the store might be read-only.  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED      | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                     | also returned if the client does not have the required credentials to call the   |
	//	|                                     | method.                                                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER  | The specified IPsec phase is not a valid one.                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method deletes all the cryptographic sets in the cryptographic linked list of
	// the memory representation of the store being modified. It also writes through and
	// deletes the sets from disk. If called on an online store, the method does not delete
	// the sets if any nonprimary set is referenced by a connection security rule.
	DeleteAllCryptoSets(context.Context, *DeleteAllCryptoSetsRequest) (*DeleteAllCryptoSetsResponse, error)

	// The RRPC_FWEnumCryptoSets method requests the server to return all the cryptographic
	// sets of the specified IPsec phase contained in the store referenced by the hPolicy
	// handle. The method returns a linked list of all these cryptographic objects.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The IpSecPhase parameter          |
	//	|                                    | specifies an invalid IPsec negotiation phase. One of the required values is not  |
	//	|                                    | specified.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumCryptoSets(context.Context, *EnumCryptoSetsRequest) (*EnumCryptoSetsResponse, error)

	// The RRPC_FWEnumPhase1SAs method requests the server to return all the security associations
	// of the IPsec first negotiation phase contained in the store referenced by the hPolicy
	// handle. The method returns a linked list of all these security associations.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The store handle is not of the dynamic store.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumPhase1SAs(context.Context, *EnumPhase1SAsRequest) (*EnumPhase1SAsResponse, error)

	// The RRPC_FWEnumPhase2SAs method requests the server to return all the security associations
	// of the IPsec second negotiation phase contained in the store referenced by the hPolicy
	// handle. The method returns a linked list of all these security associations.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a non-zero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The store handle is not of the dynamic store.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumPhase2SAs(context.Context, *EnumPhase2SAsRequest) (*EnumPhase2SAsResponse, error)

	// The RRPC_FWDeletePhase1SAs method requests the server to delete all the IPsec first
	// negotiation phase security associations that match the specified endpoints.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The store handle is not of the dynamic store.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required and not         |
	//	|                                    | specified.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	DeletePhase1SAs(context.Context, *DeletePhase1SAsRequest) (*DeletePhase1SAsResponse, error)

	// The RRPC_FWDeletePhase2SAs (Opnum 30) method requests the server to delete all the
	// IPsec second-negotiation-phase security associations that match the specified endpoints.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The store handle is not of the dynamic store.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	DeletePhase2SAs(context.Context, *DeletePhase2SAsRequest) (*DeletePhase2SAsResponse, error)

	// The RRPC_FWEnumProducts (Opnum 31) method requests the server to return all the registered
	// third-party software components registered with the firewall and advanced security
	// component. The only method supported is binary version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The store handle is not of the dynamic store.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumProducts(context.Context, *EnumProductsRequest) (*EnumProductsResponse, error)

	// The RRPC_FWAddMainModeRule (Opnum 32) method requests the server to add the main
	// mode rule in the policy contained in the policy store referenced by the specified
	// opened policy store handle. The only method supported is binary version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	AddMainModeRule(context.Context, *AddMainModeRuleRequest) (*AddMainModeRuleResponse, error)

	// The RRPC_FWSetMainModeRule (Opnum 33) method requests the server to modify the specified
	// main mode rule in the policy contained in the policy store referenced by the handle
	// specified in the hPolicy parameter. The only method supported is binary version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified set referenced by the wszRuleID member STRING of the FW_MM_RULE    |
	//	|                                    | data type is not found in the policy store.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetMainModeRule(context.Context, *SetMainModeRuleRequest) (*SetMainModeRuleResponse, error)

	// The RRPC_FWDeleteMainModeRule (Opnum 34) method requests the server to delete the
	// specified main mode rule in the policy contained in the policy store referenced by
	// the handle specified in the hPolicy parameter. The only method supported is binary
	// version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified set referenced by the wszRuleID member string of the FW_MM_RULE    |
	//	|                                    | data type is not found in the policy store.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	DeleteMainModeRule(context.Context, *DeleteMainModeRuleRequest) (*DeleteMainModeRuleResponse, error)

	// The RRPC_FWDeleteAllMainModeRules (Opnum 35) method requests the server to delete
	// all the main mode rules in the policy contained in the policy store referenced by
	// the handle specified in the hPolicy parameter. The only method supported is binary
	// version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED | The specified store does not support this method; the store might be read-only.  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                | also returned if the client does not have the required credentials to call the   |
	//	|                                | method.                                                                          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	DeleteAllMainModeRules(context.Context, *DeleteAllMainModeRulesRequest) (*DeleteAllMainModeRulesResponse, error)

	// The RRPC_FWEnumMainModeRules (Opnum 36) method requests the server to return all
	// the main mode rules contained in the store referenced by the hPolicy handle. The
	// method returns a linked list of all the main mode rule objects. The only method supported
	// is binary version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters for this method is incorrect or is required but not        |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | parameters did not meet the required constraints. - The dwProfileFilter          |
	//	|                                    | parameter contains invalid profiles.                                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumMainModeRules(context.Context, *EnumMainModeRulesRequest) (*EnumMainModeRulesResponse, error)

	// The RRPC_FWQueryFirewallRules (Opnum 37) method requests the server to return all
	// the firewall rules that match the specified query object that are contained in the
	// store referenced by the hPolicy handle. The method returns a linked list of all the
	// firewall rule objects. The only method supported is binary version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | parameters did not meet the required constraints. - The pQuery parameter         |
	//	|                                    | contains invalid profiles.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryFirewallRules(context.Context, *QueryFirewallRulesRequest) (*QueryFirewallRulesResponse, error)

	// RRPC_FWQueryConnectionSecurityRules2_10 operation.
	QueryConnectionSecurityRules210(context.Context, *QueryConnectionSecurityRules210Request) (*QueryConnectionSecurityRules210Response, error)

	// The RRPC_FWQueryMainModeRules (Opnum 39) method requests the server to return all
	// the main mode rules that match the specified query object that are contained in the
	// store referenced by the hPolicy handle. The method returns a linked list of all the
	// main mode rule objects. The only method supported is binary version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | parameters did not meet the required constraints. - The pQuery parameter         |
	//	|                                    | contains invalid profiles.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryMainModeRules(context.Context, *QueryMainModeRulesRequest) (*QueryMainModeRulesResponse, error)

	// The RRPC_FWQueryAuthenticationSets (Opnum 40) method requests the server to return
	// all the authentication sets that match the specified query object that are contained
	// in the store referenced by the hPolicy handle. The method returns a linked list of
	// all the authentication set objects. The only method supported is binary version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryAuthenticationSets(context.Context, *QueryAuthenticationSetsRequest) (*QueryAuthenticationSetsResponse, error)

	// The RRPC_FWQueryCryptoSets (Opnum 41) method requests the server to return all the
	// crypto sets that match the specified query object that are contained in the store
	// referenced by the hPolicy handle. The method returns a linked list of all the crypto
	// set objects. The only method supported is binary version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | parameters did not meet the required constraints. - The pQuery parameter         |
	//	|                                    | contains invalid profiles.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryCryptoSets(context.Context, *QueryCryptoSetsRequest) (*QueryCryptoSetsResponse, error)

	// The RRPC_FWEnumNetworks (Opnum 42) method requests the server to return all the networks
	// to which the host with the firewall and advanced security component is connected.
	// The only method supported is binary version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter contains invalid profiles.                                |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumNetworks(context.Context, *EnumNetworksRequest) (*EnumNetworksResponse, error)

	// The RRPC_FWEnumAdapters (Opnum 43) method requests the server to return all the networks
	// interfaces that the host with the firewall and advanced security component has. The
	// only method supported is binary version 0x020A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter contains invalid profiles.                                |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumAdapters(context.Context, *EnumAdaptersRequest) (*EnumAdaptersResponse, error)

	// The RRPC_FWGetGlobalConfig2_10 (Opnum 44) method retrieves the value of a global
	// policy configuration option. The client specifies to the server from which store
	// this value MUST be retrieved and in which specific configuration option it is interested.
	// The method is only supported for binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specific configuration option is not found within the policy. This means     |
	//	|                                    | that it is not configured. If the option is not configured in any other store,   |
	//	|                                    | the firewall uses a default value.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store type does not support this method.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The specific configuration option |
	//	|                                    | is not meant to be available in the specified store. The specified configuration |
	//	|                                    | option is not defined. One of the required values is not specified. The buffer   |
	//	|                                    | is not big enough to hold the specific value.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	GetGlobalConfig210(context.Context, *GetGlobalConfig210Request) (*GetGlobalConfig210Response, error)

	// The RRPC_FWGetConfig2_10 (Opnum 45) method retrieves the value of a profile configuration
	// option. The client specifies to the server from which store and profile this value
	// MUST be retrieved and in which specific configuration option it is interested. The
	// method is only supported for binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specific configuration option is not found within the policy. This means     |
	//	|                                    | that it is not configured. If the option is not configured in any other store,   |
	//	|                                    | the firewall uses a default value.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The store type specified does not support this method.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The buffer is not big enough to hold the configuration option value.             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The specific configuration option |
	//	|                                    | is not meant to be available in the specified store. The specified configuration |
	//	|                                    | option is not defined. One of the required values is not specified. The buffer   |
	//	|                                    | is not big enough to hold the specific value.                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	GetConfig210(context.Context, *GetConfig210Request) (*GetConfig210Response, error)

	// The RRPC_FWAddFirewallRule2_10 (Opnum 46) method requests the server to add the specified
	// firewall rule in the policy contained in the policy store referenced by the handle
	// specified in the hPolicyStore parameter. The method is only supported for binary
	// versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pRule object did not pass the |
	//	|                                    | firewall rule validations specified in the definition of the FW_RULE data type.  |
	//	|                                    | One of the required values is not specified. A policy store does not support     |
	//	|                                    | rules with profile conditions other than ALL profiles. The wszLocalApplication   |
	//	|                                    | parameter contains a string that at enforcement time does not represent a valid  |
	//	|                                    | file path.<34>                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// called on an online store, the firewall rule is also enforced.
	AddFirewallRule210(context.Context, *AddFirewallRule210Request) (*AddFirewallRule210Response, error)

	// The RRPC_FWSetFirewallRule2_10 (Opnum 47) method requests the server to modify the
	// specified firewall rule in the policy contained in the policy store referenced by
	// the handle specified in the hPolicyStore parameter. The method is only supported
	// for binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszRuleID member string of the FW_RULE data |
	//	|                                    | type is not found in the policy store.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pRule object did not pass the |
	//	|                                    | firewall rule validations specified in the definition of the FW_RULE data type.  |
	//	|                                    | One of the required values is not specified. A policy store does not support     |
	//	|                                    | rules with profile conditions other than ALL profiles.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetFirewallRule210(context.Context, *SetFirewallRule210Request) (*SetFirewallRule210Response, error)

	// The RRPC_FWEnumFirewallRules2_10 (Opnum 48) method requests the server to return
	// all the firewall rules contained in the store referenced by the hPolicyStore handle.
	// The method returns a linked list of all the firewall rule objects. The method is
	// only supported for binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumFirewallRules210(context.Context, *EnumFirewallRules210Request) (*EnumFirewallRules210Response, error)

	// The RRPC_FWAddConnectionSecurityRule2_10 (Opnum 49) method requests the server to
	// add the specified connection security rule in the policy contained in the policy
	// store referenced by the handle specified in the hPolicyStore parameter. The method
	// is only supported for binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pRule object did not pass     |
	//	|                                    | the firewall rule validations specified in the definition of the FW_CS_RULE      |
	//	|                                    | data type. One of the required values is not specified. A policy store does not  |
	//	|                                    | support rules with profile conditions other than ALL profiles.                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// called on an online store, the firewall rule is also enforced.
	AddConnectionSecurityRule210(context.Context, *AddConnectionSecurityRule210Request) (*AddConnectionSecurityRule210Response, error)

	// The RRPC_FWSetConnectionSecurityRule2_10 (Opnum 50) method requests the server to
	// modify the specified connection security rule in the policy contained in the policy
	// store referenced by the handle specified in the hPolicyStore parameter. The method
	// is only supported for binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszRuleID member string of the FW_CS_RULE   |
	//	|                                    | data type is not found in the policy store.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pRule object did not pass     |
	//	|                                    | the firewall rule validations specified in the definition of the FW_CS_RULE      |
	//	|                                    | data type. One of the required values is not specified. A policy store does not  |
	//	|                                    | support rules with profile conditions other than ALL profiles.                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetConnectionSecurityRule210(context.Context, *SetConnectionSecurityRule210Request) (*SetConnectionSecurityRule210Response, error)

	// The RRPC_FWEnumConnectionSecurityRules2_10 (Opnum 51) method requests the server
	// to return all the connection security rules contained in the store referenced by
	// the hPolicyStore handle. The method returns a linked list of all the connection security
	// rule objects. The method is only supported for binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumConnectionSecurityRules210(context.Context, *EnumConnectionSecurityRules210Request) (*EnumConnectionSecurityRules210Response, error)

	// The RRPC_FWAddAuthenticationSet2_10 (Opnum 52) method requests the server to add
	// the authentication set in the policy contained in the policy store referenced by
	// the handle specified in the hPolicy parameter. The method is only supported for binary
	// versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pAuth object did not pass the |
	//	|                                    | firewall rule validations specified in the definition of the FW_AUTH_SET data    |
	//	|                                    | type. One of the required values is not specified.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// the method is called on an online store, the firewall rule is also enforced.
	AddAuthenticationSet210(context.Context, *AddAuthenticationSet210Request) (*AddAuthenticationSet210Response, error)

	// The RRPC_FWSetAuthenticationSet2_10 (Opnum 53) method requests the server to modify
	// the specified authentication set in the policy contained in the policy store referenced
	// by the handle specified in the hPolicy parameter. The method is only supported for
	// binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszSetId member string of the FW_AUTH_SET   |
	//	|                                    | data type is not found in the policy store.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pAuth object did not pass the |
	//	|                                    | firewall rule validations specified in the definition of the FW_AUTH_SET data    |
	//	|                                    | type. One of the required values is not specified.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetAuthenticationSet210(context.Context, *SetAuthenticationSet210Request) (*SetAuthenticationSet210Response, error)

	// The RRPC_FWEnumAuthenticationSets2_10 (Opnum 54) method requests the server to return
	// all the authentication sets of the specified IPsec phase contained in the store referenced
	// by the hPolicyStore handle. The method returns a linked list of these objects. The
	// method is only supported for binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumAuthenticationSets210(context.Context, *EnumAuthenticationSets210Request) (*EnumAuthenticationSets210Response, error)

	// The RRPC_FWAddCryptoSet2_10 (Opnum 55) method adds a cryptographic set in the cryptographic
	// linked list of the memory representation of the store being modified. The method
	// is only supported for binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified cryptographic set has a cryptographic set ID that already exists   |
	//	|                                    | in the specified store.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pCrypto object did not pass   |
	//	|                                    | the crypto set validations specified in the definition of the FW_CRYPTO_SET data |
	//	|                                    | type. One of the required values is not specified.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// called on an online store, the firewall rule is also enforced.
	AddCryptoSet210(context.Context, *AddCryptoSet210Request) (*AddCryptoSet210Response, error)

	// The RRPC_FWSetCryptoSet2_10 (Opnum 56) method requests the server to modify the specified
	// cryptographic set in the policy contained in the policy store referenced by the handle
	// specified in the hPolicy parameter. The method is only supported for binary versions
	// 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszSetId member string of the FW_CRYPTO_SET |
	//	|                                    | data type is not found in the policy store.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pCrypto object did not pass   |
	//	|                                    | the crypto set validations specified in the definition of the FW_CRYPTO_SET data |
	//	|                                    | type. One of the required values is not specified.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetCryptoSet210(context.Context, *SetCryptoSet210Request) (*SetCryptoSet210Response, error)

	// The RRPC_FWEnumCryptoSets2_10 (Opnum 57) method requests the server to return all
	// the cryptographic sets of the specified IPsec phase contained in the store referenced
	// by the hPolicyStore handle. The method returns a linked list of these objects. The
	// method is only supported for binary versions 0x020A and 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumCryptoSets210(context.Context, *EnumCryptoSets210Request) (*EnumCryptoSets210Response, error)

	// The RRPC_FWAddConnectionSecurityRule2_20 method requests the server to add the specified
	// connection security rule in the policy contained in the policy store referenced by
	// the handle specified in the hPolicyStore parameter. The method is only supported
	// for binary version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pRule object did not pass     |
	//	|                                    | the firewall rule validations specified in the definition of the FW_CS_RULE      |
	//	|                                    | data type. One of the required values is not specified. A policy store does not  |
	//	|                                    | support rules with profile conditions other than ALL profiles.                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// called on an online store, the firewall rule is also enforced.
	AddConnectionSecurityRule220(context.Context, *AddConnectionSecurityRule220Request) (*AddConnectionSecurityRule220Response, error)

	// The RRPC_FWSetConnectionSecurityRule2_20 method requests the server to modify the
	// specified connection security rule in the policy contained in the policy store referenced
	// by the handle specified in the hPolicyStore parameter. The method is only supported
	// for binary version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszRuleID member string of the FW_CS_RULE   |
	//	|                                    | data type is not found in the policy store.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pRule object did not pass     |
	//	|                                    | the firewall rule validations specified in the definition of the FW_CS_RULE      |
	//	|                                    | data type. One of the required values is not specified. A policy store does not  |
	//	|                                    | support rules with profile conditions other than ALL profiles.                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetConnectionSecurityRule220(context.Context, *SetConnectionSecurityRule220Request) (*SetConnectionSecurityRule220Response, error)

	// The RRPC_FWEnumConnectionSecurityRules2_20 (Opnum 60) method requests the server
	// to return all the connection security rules contained in the store referenced by
	// the hPolicyStore handle. The method returns a linked list of all the connection security
	// rule objects. The method is only supported for binary version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The dwProfileFilter parameter contains invalid profiles.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumConnectionSecurityRules220(context.Context, *EnumConnectionSecurityRules220Request) (*EnumConnectionSecurityRules220Response, error)

	// The RRPC_FWQueryConnectionSecurityRules2_20 (Opnum 61) method requests the server
	// to return all the connection security rules that match the specified query object
	// that are contained in the store referenced by the hPolicy handle. The method returns
	// a linked list of all the connection security rule objects. The method is only supported
	// for binary version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryConnectionSecurityRules220(context.Context, *QueryConnectionSecurityRules220Request) (*QueryConnectionSecurityRules220Response, error)

	// The RRPC_FWAddAuthenticationSet2_20 method requests the server to add the authentication
	// set in the policy contained in the policy store referenced by the handle specified
	// in the hPolicy parameter. The method is only supported for binary version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified set has a set ID that already exists in the specified store.       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The pAuth object did not pass the |
	//	|                                    | firewall rule validations specified in the definition of the FW_AUTH_SET data    |
	//	|                                    | type. One of the required values is not specified.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule in disk. If
	// the method is called on an online store, the firewall rule is also enforced.
	AddAuthenticationSet220(context.Context, *AddAuthenticationSet220Request) (*AddAuthenticationSet220Response, error)

	// The RRPC_FWSetAuthenticationSet2_20 method requests the server to modify the specified
	// authentication set in the policy contained in the policy store referenced by the
	// handle specified in the hPolicy parameter. The method is only supported for binary
	// version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszSetId member string of the FW_AUTH_SET   |
	//	|                                    | data type is not found in the policy store.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect, or is required and not        |
	//	|                                    | specified. This error can be returned because: The pAuth object did not pass the |
	//	|                                    | firewall rule validations specified in the definition of the FW_AUTH_SET data    |
	//	|                                    | type. One of the required values is not specified.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	SetAuthenticationSet220(context.Context, *SetAuthenticationSet220Request) (*SetAuthenticationSet220Response, error)

	// The RRPC_FWEnumAuthenticationSets2_20 method requests the server to return all the
	// authentication sets of the specified IPsec phase contained in the store referenced
	// in the hPolicy handle. The method returns a linked list of these objects. The method
	// is only supported for binary version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicy handle was not opened with read/write access rights. The error is    |
	//	|                                    | also returned if the client does not have the required credentials to call the   |
	//	|                                    | method.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The dwProfileFilter parameter contains invalid profiles.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	EnumAuthenticationSets220(context.Context, *EnumAuthenticationSets220Request) (*EnumAuthenticationSets220Response, error)

	// The RRPC_FWQueryAuthenticationSets2_20 method requests the server to return all the
	// authentication sets that match the specified query object that are contained in the
	// store referenced in the hPolicy handle. The method returns a linked list of all the
	// authentication set objects. The method is only supported for binary version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	QueryAuthenticationSets220(context.Context, *QueryAuthenticationSets220Request) (*QueryAuthenticationSets220Response, error)

	// The RRPC_FWAddFirewallRule2_20 method requests the server to add the specified firewall
	// rule in the policy contained in the policy store referenced by the handle specified
	// in the hPolicyStore parameter. The method is only supported for binary version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pRule object did not pass the |
	//	|                                    | firewall rule validations specified in the definition of the FW_RULE data type.  |
	//	|                                    | One of the required values is not specified. A policy store does not support     |
	//	|                                    | rules with profile conditions other than ALL profiles. The wszLocalApplication   |
	//	|                                    | member contains a string that, at enforcement time, does not represent a valid   |
	//	|                                    | file path.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// called on an online store, the firewall rule is also enforced.
	AddFirewallRule220(context.Context, *AddFirewallRule220Request) (*AddFirewallRule220Response, error)

	// The RRPC_FWAddConnectionSecurityRule2_20 method requests the server to modify the
	// specified connection security rule in the policy contained in the policy store referenced
	// by the handle specified in the hPolicyStore parameter. The method is only supported
	// for binary version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszRuleID member string of the FW_RULE data |
	//	|                                    | type is not found in the policy store.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method either is incorrect or is required and not  |
	//	|                                    | specified. This error can be returned because: The pRule object did not pass the |
	//	|                                    | firewall rule validations specified in the definition of the FW_RULE data type.  |
	//	|                                    | One of the required values is not specified. A policy store does not support     |
	//	|                                    | rules with profile conditions other than ALL profiles.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetFirewallRule220(context.Context, *SetFirewallRule220Request) (*SetFirewallRule220Response, error)

	// The RRPC_FWEnumFirewallRules2_20 (Opnum 68) method requests the server to return
	// all the firewall rules contained in the store referenced by the hPolicyStore handle.
	// The method returns a linked list of all the firewall rule objects. The method is
	// only supported for binary version 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. The error  |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumFirewallRules220(context.Context, *EnumFirewallRules220Request) (*EnumFirewallRules220Response, error)

	// The RRPC_FWQueryFirewallRules2_20 (Opnum 69) method requests the server to return
	// all the firewall rules that match the specified query object that are contained in
	// the store referenced by the hPolicy handle. The method returns a linked list of all
	// the connection security rule objects. The method is only supported for binary version
	// 0x0214.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The pQuery parameter contains invalid        |
	//	|                                    | profiles.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryFirewallRules220(context.Context, *QueryFirewallRules220Request) (*QueryFirewallRules220Response, error)

	// The RRPC_FWAddFirewallRule2_24 method requests the server to add the specified firewall
	// rule in the policy contained in the policy store that is referenced by the handle
	// specified in the hPolicyStore parameter. The method is only supported for binary
	// version 0x0218.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - The pRule        |
	//	|                                    | object did not pass the firewall rule validations specified in the definition    |
	//	|                                    | of the FW_RULE data type. - One of the required values is not specified. -       |
	//	|                                    | A policy store does not support rules with profile conditions other than ALL     |
	//	|                                    | profiles. - The wszLocalApplication member of the rule contains a string that,   |
	//	|                                    | at enforcement time, does not represent a valid file path.                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// called on an online store, the firewall rule is also enforced.
	AddFirewallRule224(context.Context, *AddFirewallRule224Request) (*AddFirewallRule224Response, error)

	// The RRPC_FWSetFirewallRule2_24 method requests the server to modify the specified
	// connection security rule in the policy contained in the policy store that is referenced
	// by the handle specified in the hPolicyStore parameter. The method is only supported
	// for binary version 0x0218.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszRuleID member string of the FW_RULE data |
	//	|                                    | type is not found in the policy store.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - The pRule object |
	//	|                                    | did not pass the firewall rule validations specified in the definition of the    |
	//	|                                    | FW_RULE data type. - One of the required values is not specified. - A policy     |
	//	|                                    | store does not support rules with profile conditions other than ALL profiles.    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetFirewallRule224(context.Context, *SetFirewallRule224Request) (*SetFirewallRule224Response, error)

	// The RRPC_FWEnumFirewallRules2_24 method requests the server to return all the firewall
	// rules contained in the store that is referenced by the hPolicyStore handle. The method
	// returns a linked list of all the firewall rule objects. The method is only supported
	// for binary version 0x0218.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumFirewallRules224(context.Context, *EnumFirewallRules224Request) (*EnumFirewallRules224Response, error)

	// The RRPC_FWQueryFirewallRules2_24 method requests the server to return all the firewall
	// rules that match the specified query object that are contained in the store that
	// is referenced by the hPolicyStore handle. The method returns a linked list of all
	// the connection security rule objects. The method is only supported for binary version
	// 0x0218.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The pQuery parameter contains invalid        |
	//	|                                    | conditions.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryFirewallRules224(context.Context, *QueryFirewallRules224Request) (*QueryFirewallRules224Response, error)

	// The RRPC_FWAddFirewallRule2_25 method requests the server to add the specified firewall
	// rule in the policy contained in the policy store that is referenced by the handle
	// specified in the hPolicyStore parameter. The method is only supported for binary
	// version 0x0219.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - The pRule        |
	//	|                                    | object did not pass the firewall rule validations specified in the definition    |
	//	|                                    | of the FW_RULE data type. - One of the required values is not specified. -       |
	//	|                                    | A policy store does not support rules with profile conditions other than ALL     |
	//	|                                    | profiles. - The wszLocalApplication member of the rule contains a string that,   |
	//	|                                    | at enforcement time, does not represent a valid file path.                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// called on an online store, the firewall rule is also enforced.
	AddFirewallRule225(context.Context, *AddFirewallRule225Request) (*AddFirewallRule225Response, error)

	// The RRPC_FWSetFirewallRule2_25 method requests the server to modify the specified
	// connection security rule in the policy contained in the policy store that is referenced
	// by the handle specified in the hPolicyStore parameter. The method is only supported
	// for binary version 0x0219.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszRuleID member string of the FW_RULE data |
	//	|                                    | type is not found in the policy store.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - The pRule object |
	//	|                                    | did not pass the firewall rule validations specified in the definition of the    |
	//	|                                    | FW_RULE data type. - One of the required values is not specified. - A policy     |
	//	|                                    | store does not support rules with profile conditions other than ALL profiles.    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetFirewallRule225(context.Context, *SetFirewallRule225Request) (*SetFirewallRule225Response, error)

	// The RRPC_FWEnumFirewallRules2_25 method requests the server to return all the firewall
	// rules contained in the store that is referenced by the hPolicyStore handle. The method
	// returns a linked list of all the firewall rule objects. The method is only supported
	// for binary version 0x0219.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumFirewallRules225(context.Context, *EnumFirewallRules225Request) (*EnumFirewallRules225Response, error)

	// The RRPC_FWQueryFirewallRules2_25 method requests the server to return all the firewall
	// rules that match the specified query object that are contained in the store that
	// is referenced by the hPolicyStore handle. The method returns a linked list of all
	// the connection security rule objects. The method is only supported for binary version
	// 0x0219.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The pQuery parameter contains invalid        |
	//	|                                    | conditions.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryFirewallRules225(context.Context, *QueryFirewallRules225Request) (*QueryFirewallRules225Response, error)

	// The RRPC_FWAddFirewallRule2_26 method requests the server to add the specified firewall
	// rule in the policy contained in the policy store that is referenced by the handle
	// specified in the hPolicyStore parameter. The method is only supported for binary
	// version 0x021A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - The pRule        |
	//	|                                    | object did not pass the firewall rule validations specified in the definition    |
	//	|                                    | of the FW_RULE data type. - One of the required values is not specified. -       |
	//	|                                    | A policy store does not support rules with profile conditions other than ALL     |
	//	|                                    | profiles. - The wszLocalApplication member of the rule contains a string that,   |
	//	|                                    | at enforcement time, does not represent a valid file path.                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// called on an online store, the firewall rule is also enforced.
	AddFirewallRule226(context.Context, *AddFirewallRule226Request) (*AddFirewallRule226Response, error)

	// The RRPC_FWSetFirewallRule2_26 method requests the server to modify the specified
	// connection security rule in the policy contained in the policy store that is referenced
	// by the handle specified in the hPolicyStore parameter. The method is only supported
	// for binary version 0x021A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszRuleID member string of the FW_RULE data |
	//	|                                    | type is not found in the policy store.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - The pRule object |
	//	|                                    | did not pass the firewall rule validations specified in the definition of the    |
	//	|                                    | FW_RULE data type. - One of the required values is not specified. - A policy     |
	//	|                                    | store does not support rules with profile conditions other than ALL profiles.    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetFirewallRule226(context.Context, *SetFirewallRule226Request) (*SetFirewallRule226Response, error)

	// The RRPC_FWEnumFirewallRules2_26 method requests the server to return all the firewall
	// rules contained in the store that is referenced by the hPolicyStore handle. The method
	// returns a linked list of all the firewall rule objects. The method is only supported
	// for binary version 0x021A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumFirewallRules226(context.Context, *EnumFirewallRules226Request) (*EnumFirewallRules226Response, error)

	// The RRPC_FWQueryFirewallRules2_26 method requests the server to return all the firewall
	// rules that match the specified query object that are contained in the store that
	// is referenced by the hPolicyStore handle. The method returns a linked list of all
	// the connection security rule objects. The method is only supported for binary version
	// 0x021A.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The pQuery parameter contains invalid        |
	//	|                                    | conditions.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryFirewallRules226(context.Context, *QueryFirewallRules226Request) (*QueryFirewallRules226Response, error)

	// The RRPC_FWAddFirewallRule2_27 method requests the server to add the specified firewall
	// rule to the policy contained in the policy store that is referenced by the handle
	// specified in the hPolicyStore parameter. The method is only supported for binary
	// version 0x021B.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - The pRule        |
	//	|                                    | object did not pass the firewall rule validations specified in the definition    |
	//	|                                    | of the FW_RULE data type (section 2.2.37). - One of the required values is       |
	//	|                                    | not specified. - A policy store does not support rules with profile conditions   |
	//	|                                    | other than ALL profiles. - The wszLocalApplication member of the rule contains a |
	//	|                                    | string that, at enforcement time, does not represent a valid file path.          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// called on an online store, the firewall rule is also enforced.
	AddFirewallRule227(context.Context, *AddFirewallRule227Request) (*AddFirewallRule227Response, error)

	// The RRPC_FWSetFirewallRule2_27 method requests the server to modify the specified
	// connection security rule in the policy contained in the policy store that is referenced
	// by the handle specified in the hPolicyStore parameter. The method is only supported
	// for binary version 0x021B.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszRuleID member string of the FW_RULE data |
	//	|                                    | type (section 2.2.37) is not found in the policy store.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - The pRule        |
	//	|                                    | object did not pass the firewall rule validations specified in the definition    |
	//	|                                    | of the FW_RULE data type (section 2.2.37). - One of the required values is not   |
	//	|                                    | specified. - A policy store does not support rules with profile conditions other |
	//	|                                    | than ALL profiles.                                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetFirewallRule227(context.Context, *SetFirewallRule227Request) (*SetFirewallRule227Response, error)

	// The RRPC_FWEnumFirewallRules2_27 method requests the server to return all the firewall
	// rules contained in the store that is referenced by the hPolicyStore handle. The method
	// returns a linked list of all the firewall rule objects. The method is only supported
	// for binary version 0x021B.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumFirewallRules227(context.Context, *EnumFirewallRules227Request) (*EnumFirewallRules227Response, error)

	// The RRPC_FWQueryFirewallRules2_27 method requests the server to return all the firewall
	// rules that match the specified query object that are contained in the store that
	// is referenced by the hPolicyStore handle. The method returns a linked list of all
	// the connection security rule objects. The method is only supported for binary version
	// 0x021B.
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The pQuery parameter contains invalid        |
	//	|                                    | conditions.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryFirewallRules227(context.Context, *QueryFirewallRules227Request) (*QueryFirewallRules227Response, error)

	// The RRPC_FWAddFirewallRule2_31 method requests the server to add the specified firewall
	// rule to the policy contained in the policy store that is referenced by the handle
	// specified in the hPolicyStore parameter. The method is only supported for policy
	// version 0x021F (section 2.2.42).
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | The specified rule has a rule ID that already exists in the specified store.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - The pRule        |
	//	|                                    | object did not pass the firewall rule validations specified in the definition    |
	//	|                                    | of the FW_RULE data type (section 2.2.37). - One of the required values is       |
	//	|                                    | not specified. - A policy store does not support rules with profile conditions   |
	//	|                                    | other than ALL profiles. - The wszLocalApplication member of the rule contains a |
	//	|                                    | string that, at enforcement time, does not represent a valid file path.          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	//
	// This method adds a firewall rule in the firewall linked list of the memory representation
	// of the store being modified. It also writes through and saves the rule on disk. If
	// called on an online store, the firewall rule is also enforced.
	AddFirewallRule231(context.Context, *AddFirewallRule231Request) (*AddFirewallRule231Response, error)

	// The RRPC_FWSetFirewallRule2_31 method requests the server to modify the specified
	// connection security rule in the policy contained in the policy store that is referenced
	// by the handle specified in the hPolicyStore parameter. The method is only supported
	// for policy version 0x021F (section 2.2.42).
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified store does not support this method; the store might be read-only.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified rule referenced by the wszRuleID member string of the FW_RULE data |
	//	|                                    | type (section 2.2.37) is not found in the policy store.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - The pRule        |
	//	|                                    | object did not pass the firewall rule validations specified in the definition    |
	//	|                                    | of the FW_RULE data type (section 2.2.37). - One of the required values is not   |
	//	|                                    | specified. - A policy store does not support rules with profile conditions other |
	//	|                                    | than ALL profiles.                                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	SetFirewallRule231(context.Context, *SetFirewallRule231Request) (*SetFirewallRule231Response, error)

	// The RRPC_FWEnumFirewallRules2_31 method requests the server to return the firewall
	// rules matching the input flags contained in the store that is referenced by the hPolicyStore
	// handle. The method returns a linked list of the corresponding firewall rule objects.
	// The method is only supported for policy version 0x021F (section 2.2.42).
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The hPolicyStore handle was not opened with read/write access rights. This error |
	//	|                                    | is also returned if the client does not have the required credentials to call    |
	//	|                                    | the method.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The dwProfileFilter parameter contains       |
	//	|                                    | invalid profiles.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	EnumFirewallRules231(context.Context, *EnumFirewallRules231Request) (*EnumFirewallRules231Response, error)

	// The RRPC_FWQueryFirewallRules2_31 method requests the server to return all the firewall
	// rules that match the specified query object, as are contained in the store that is
	// referenced by the hPolicyStore handle. The method returns a linked list of all the
	// firewall rules that match the specified query object. The method is only supported
	// for policy version 0x021F (section 2.2.42).
	//
	// Return Values: The method returns 0 if successful; if it fails, it returns a nonzero
	// error code. The field can take any specific error code value, as specified in [MS-ERREF].
	// The following return values are common.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The client does not have the required credentials to call the method.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters of this method is incorrect or is required but not         |
	//	|                                    | specified. This error can be returned in the following cases: - One of the       |
	//	|                                    | required values is not specified. - The pQuery parameter contains invalid        |
	//	|                                    | conditions.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE]. If any lower-layer errors are reported by
	// RPC exception, this exception is converted to an error code and reported to higher-layer
	// protocols via the return value.
	QueryFirewallRules231(context.Context, *QueryFirewallRules231Request) (*QueryFirewallRules231Response, error)
}

func RegisterRemoteFwServer(conn dcerpc.Conn, o RemoteFwServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteFwServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteFwSyntaxV1_0))...)
}

func NewRemoteFwServerHandle(o RemoteFwServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteFwServerHandle(ctx, o, opNum, r)
	}
}

func RemoteFwServerHandle(ctx context.Context, o RemoteFwServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RRPC_FWOpenPolicyStore
		op := &xxx_OpenPolicyStoreOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenPolicyStoreRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenPolicyStore(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RRPC_FWClosePolicyStore
		op := &xxx_ClosePolicyStoreOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClosePolicyStoreRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClosePolicyStore(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // RRPC_FWRestoreDefaults
		op := &xxx_RestoreDefaultsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestoreDefaultsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RestoreDefaults(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // RRPC_FWGetGlobalConfig
		op := &xxx_GetGlobalConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetGlobalConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetGlobalConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RRPC_FWSetGlobalConfig
		op := &xxx_SetGlobalConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetGlobalConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetGlobalConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // RRPC_FWAddFirewallRule
		op := &xxx_AddFirewallRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddFirewallRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddFirewallRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // RRPC_FWSetFirewallRule
		op := &xxx_SetFirewallRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFirewallRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFirewallRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // RRPC_FWDeleteFirewallRule
		op := &xxx_DeleteFirewallRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteFirewallRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteFirewallRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RRPC_FWDeleteAllFirewallRules
		op := &xxx_DeleteAllFirewallRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAllFirewallRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAllFirewallRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RRPC_FWEnumFirewallRules
		op := &xxx_EnumFirewallRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFirewallRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFirewallRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RRPC_FWGetConfig
		op := &xxx_GetConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RRPC_FWSetConfig
		op := &xxx_SetConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RRPC_FWAddConnectionSecurityRule
		op := &xxx_AddConnectionSecurityRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddConnectionSecurityRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddConnectionSecurityRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // RRPC_FWSetConnectionSecurityRule
		op := &xxx_SetConnectionSecurityRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConnectionSecurityRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConnectionSecurityRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // RRPC_FWDeleteConnectionSecurityRule
		op := &xxx_DeleteConnectionSecurityRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteConnectionSecurityRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteConnectionSecurityRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RRPC_FWDeleteAllConnectionSecurityRules
		op := &xxx_DeleteAllConnectionSecurityRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAllConnectionSecurityRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAllConnectionSecurityRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // RRPC_FWEnumConnectionSecurityRules
		op := &xxx_EnumConnectionSecurityRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumConnectionSecurityRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumConnectionSecurityRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // RRPC_FWAddAuthenticationSet
		op := &xxx_AddAuthenticationSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddAuthenticationSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddAuthenticationSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // RRPC_FWSetAuthenticationSet
		op := &xxx_SetAuthenticationSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAuthenticationSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAuthenticationSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // RRPC_FWDeleteAuthenticationSet
		op := &xxx_DeleteAuthenticationSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAuthenticationSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAuthenticationSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // RRPC_FWDeleteAllAuthenticationSets
		op := &xxx_DeleteAllAuthenticationSetsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAllAuthenticationSetsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAllAuthenticationSets(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // RRPC_FWEnumAuthenticationSets
		op := &xxx_EnumAuthenticationSetsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAuthenticationSetsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAuthenticationSets(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // RRPC_FWAddCryptoSet
		op := &xxx_AddCryptoSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddCryptoSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddCryptoSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // RRPC_FWSetCryptoSet
		op := &xxx_SetCryptoSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCryptoSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCryptoSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // RRPC_FWDeleteCryptoSet
		op := &xxx_DeleteCryptoSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteCryptoSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteCryptoSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // RRPC_FWDeleteAllCryptoSets
		op := &xxx_DeleteAllCryptoSetsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAllCryptoSetsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAllCryptoSets(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // RRPC_FWEnumCryptoSets
		op := &xxx_EnumCryptoSetsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumCryptoSetsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumCryptoSets(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // RRPC_FWEnumPhase1SAs
		op := &xxx_EnumPhase1SAsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPhase1SAsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPhase1SAs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // RRPC_FWEnumPhase2SAs
		op := &xxx_EnumPhase2SAsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPhase2SAsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPhase2SAs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // RRPC_FWDeletePhase1SAs
		op := &xxx_DeletePhase1SAsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePhase1SAsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePhase1SAs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // RRPC_FWDeletePhase2SAs
		op := &xxx_DeletePhase2SAsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePhase2SAsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePhase2SAs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // RRPC_FWEnumProducts
		op := &xxx_EnumProductsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumProductsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumProducts(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // RRPC_FWAddMainModeRule
		op := &xxx_AddMainModeRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddMainModeRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddMainModeRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // RRPC_FWSetMainModeRule
		op := &xxx_SetMainModeRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMainModeRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMainModeRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // RRPC_FWDeleteMainModeRule
		op := &xxx_DeleteMainModeRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteMainModeRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteMainModeRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // RRPC_FWDeleteAllMainModeRules
		op := &xxx_DeleteAllMainModeRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAllMainModeRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAllMainModeRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // RRPC_FWEnumMainModeRules
		op := &xxx_EnumMainModeRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumMainModeRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumMainModeRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // RRPC_FWQueryFirewallRules
		op := &xxx_QueryFirewallRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryFirewallRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryFirewallRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // RRPC_FWQueryConnectionSecurityRules2_10
		op := &xxx_QueryConnectionSecurityRules210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryConnectionSecurityRules210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryConnectionSecurityRules210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // RRPC_FWQueryMainModeRules
		op := &xxx_QueryMainModeRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryMainModeRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryMainModeRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // RRPC_FWQueryAuthenticationSets
		op := &xxx_QueryAuthenticationSetsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryAuthenticationSetsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryAuthenticationSets(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // RRPC_FWQueryCryptoSets
		op := &xxx_QueryCryptoSetsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryCryptoSetsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryCryptoSets(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // RRPC_FWEnumNetworks
		op := &xxx_EnumNetworksOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumNetworksRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumNetworks(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // RRPC_FWEnumAdapters
		op := &xxx_EnumAdaptersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAdaptersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAdapters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // RRPC_FWGetGlobalConfig2_10
		op := &xxx_GetGlobalConfig210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetGlobalConfig210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetGlobalConfig210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // RRPC_FWGetConfig2_10
		op := &xxx_GetConfig210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfig210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfig210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // RRPC_FWAddFirewallRule2_10
		op := &xxx_AddFirewallRule210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddFirewallRule210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddFirewallRule210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // RRPC_FWSetFirewallRule2_10
		op := &xxx_SetFirewallRule210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFirewallRule210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFirewallRule210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // RRPC_FWEnumFirewallRules2_10
		op := &xxx_EnumFirewallRules210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFirewallRules210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFirewallRules210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // RRPC_FWAddConnectionSecurityRule2_10
		op := &xxx_AddConnectionSecurityRule210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddConnectionSecurityRule210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddConnectionSecurityRule210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // RRPC_FWSetConnectionSecurityRule2_10
		op := &xxx_SetConnectionSecurityRule210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConnectionSecurityRule210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConnectionSecurityRule210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // RRPC_FWEnumConnectionSecurityRules2_10
		op := &xxx_EnumConnectionSecurityRules210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumConnectionSecurityRules210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumConnectionSecurityRules210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // RRPC_FWAddAuthenticationSet2_10
		op := &xxx_AddAuthenticationSet210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddAuthenticationSet210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddAuthenticationSet210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // RRPC_FWSetAuthenticationSet2_10
		op := &xxx_SetAuthenticationSet210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAuthenticationSet210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAuthenticationSet210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // RRPC_FWEnumAuthenticationSets2_10
		op := &xxx_EnumAuthenticationSets210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAuthenticationSets210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAuthenticationSets210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // RRPC_FWAddCryptoSet2_10
		op := &xxx_AddCryptoSet210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddCryptoSet210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddCryptoSet210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // RRPC_FWSetCryptoSet2_10
		op := &xxx_SetCryptoSet210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCryptoSet210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCryptoSet210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // RRPC_FWEnumCryptoSets2_10
		op := &xxx_EnumCryptoSets210Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumCryptoSets210Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumCryptoSets210(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // RRPC_FWAddConnectionSecurityRule2_20
		op := &xxx_AddConnectionSecurityRule220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddConnectionSecurityRule220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddConnectionSecurityRule220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // RRPC_FWSetConnectionSecurityRule2_20
		op := &xxx_SetConnectionSecurityRule220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConnectionSecurityRule220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConnectionSecurityRule220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 60: // RRPC_FWEnumConnectionSecurityRules2_20
		op := &xxx_EnumConnectionSecurityRules220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumConnectionSecurityRules220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumConnectionSecurityRules220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 61: // RRPC_FWQueryConnectionSecurityRules2_20
		op := &xxx_QueryConnectionSecurityRules220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryConnectionSecurityRules220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryConnectionSecurityRules220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 62: // RRPC_FWAddAuthenticationSet2_20
		op := &xxx_AddAuthenticationSet220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddAuthenticationSet220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddAuthenticationSet220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 63: // RRPC_FWSetAuthenticationSet2_20
		op := &xxx_SetAuthenticationSet220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAuthenticationSet220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAuthenticationSet220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 64: // RRPC_FWEnumAuthenticationSets2_20
		op := &xxx_EnumAuthenticationSets220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAuthenticationSets220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAuthenticationSets220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 65: // RRPC_FWQueryAuthenticationSets2_20
		op := &xxx_QueryAuthenticationSets220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryAuthenticationSets220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryAuthenticationSets220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 66: // RRPC_FWAddFirewallRule2_20
		op := &xxx_AddFirewallRule220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddFirewallRule220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddFirewallRule220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 67: // RRPC_FWSetFirewallRule2_20
		op := &xxx_SetFirewallRule220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFirewallRule220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFirewallRule220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 68: // RRPC_FWEnumFirewallRules2_20
		op := &xxx_EnumFirewallRules220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFirewallRules220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFirewallRules220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 69: // RRPC_FWQueryFirewallRules2_20
		op := &xxx_QueryFirewallRules220Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryFirewallRules220Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryFirewallRules220(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 70: // RRPC_FWAddFirewallRule2_24
		op := &xxx_AddFirewallRule224Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddFirewallRule224Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddFirewallRule224(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 71: // RRPC_FWSetFirewallRule2_24
		op := &xxx_SetFirewallRule224Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFirewallRule224Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFirewallRule224(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 72: // RRPC_FWEnumFirewallRules2_24
		op := &xxx_EnumFirewallRules224Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFirewallRules224Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFirewallRules224(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 73: // RRPC_FWQueryFirewallRules2_24
		op := &xxx_QueryFirewallRules224Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryFirewallRules224Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryFirewallRules224(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 74: // RRPC_FWAddFirewallRule2_25
		op := &xxx_AddFirewallRule225Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddFirewallRule225Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddFirewallRule225(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 75: // RRPC_FWSetFirewallRule2_25
		op := &xxx_SetFirewallRule225Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFirewallRule225Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFirewallRule225(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 76: // RRPC_FWEnumFirewallRules2_25
		op := &xxx_EnumFirewallRules225Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFirewallRules225Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFirewallRules225(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 77: // RRPC_FWQueryFirewallRules2_25
		op := &xxx_QueryFirewallRules225Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryFirewallRules225Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryFirewallRules225(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 78: // RRPC_FWAddFirewallRule2_26
		op := &xxx_AddFirewallRule226Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddFirewallRule226Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddFirewallRule226(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 79: // RRPC_FWSetFirewallRule2_26
		op := &xxx_SetFirewallRule226Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFirewallRule226Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFirewallRule226(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 80: // RRPC_FWEnumFirewallRules2_26
		op := &xxx_EnumFirewallRules226Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFirewallRules226Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFirewallRules226(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 81: // RRPC_FWQueryFirewallRules2_26
		op := &xxx_QueryFirewallRules226Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryFirewallRules226Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryFirewallRules226(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 82: // RRPC_FWAddFirewallRule2_27
		op := &xxx_AddFirewallRule227Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddFirewallRule227Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddFirewallRule227(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 83: // RRPC_FWSetFirewallRule2_27
		op := &xxx_SetFirewallRule227Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFirewallRule227Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFirewallRule227(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 84: // RRPC_FWEnumFirewallRules2_27
		op := &xxx_EnumFirewallRules227Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFirewallRules227Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFirewallRules227(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 85: // RRPC_FWQueryFirewallRules2_27
		op := &xxx_QueryFirewallRules227Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryFirewallRules227Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryFirewallRules227(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 86: // RRPC_FWAddFirewallRule2_31
		op := &xxx_AddFirewallRule231Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddFirewallRule231Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddFirewallRule231(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 87: // RRPC_FWSetFirewallRule2_31
		op := &xxx_SetFirewallRule231Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFirewallRule231Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFirewallRule231(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 88: // RRPC_FWEnumFirewallRules2_31
		op := &xxx_EnumFirewallRules231Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFirewallRules231Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumFirewallRules231(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 89: // RRPC_FWQueryFirewallRules2_31
		op := &xxx_QueryFirewallRules231Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryFirewallRules231Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryFirewallRules231(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented RemoteFW
type UnimplementedRemoteFwServer struct {
}

func (UnimplementedRemoteFwServer) OpenPolicyStore(context.Context, *OpenPolicyStoreRequest) (*OpenPolicyStoreResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) ClosePolicyStore(context.Context, *ClosePolicyStoreRequest) (*ClosePolicyStoreResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) RestoreDefaults(context.Context, *RestoreDefaultsRequest) (*RestoreDefaultsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) GetGlobalConfig(context.Context, *GetGlobalConfigRequest) (*GetGlobalConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetGlobalConfig(context.Context, *SetGlobalConfigRequest) (*SetGlobalConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddFirewallRule(context.Context, *AddFirewallRuleRequest) (*AddFirewallRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetFirewallRule(context.Context, *SetFirewallRuleRequest) (*SetFirewallRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeleteFirewallRule(context.Context, *DeleteFirewallRuleRequest) (*DeleteFirewallRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeleteAllFirewallRules(context.Context, *DeleteAllFirewallRulesRequest) (*DeleteAllFirewallRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumFirewallRules(context.Context, *EnumFirewallRulesRequest) (*EnumFirewallRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddConnectionSecurityRule(context.Context, *AddConnectionSecurityRuleRequest) (*AddConnectionSecurityRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetConnectionSecurityRule(context.Context, *SetConnectionSecurityRuleRequest) (*SetConnectionSecurityRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeleteConnectionSecurityRule(context.Context, *DeleteConnectionSecurityRuleRequest) (*DeleteConnectionSecurityRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeleteAllConnectionSecurityRules(context.Context, *DeleteAllConnectionSecurityRulesRequest) (*DeleteAllConnectionSecurityRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumConnectionSecurityRules(context.Context, *EnumConnectionSecurityRulesRequest) (*EnumConnectionSecurityRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddAuthenticationSet(context.Context, *AddAuthenticationSetRequest) (*AddAuthenticationSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetAuthenticationSet(context.Context, *SetAuthenticationSetRequest) (*SetAuthenticationSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeleteAuthenticationSet(context.Context, *DeleteAuthenticationSetRequest) (*DeleteAuthenticationSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeleteAllAuthenticationSets(context.Context, *DeleteAllAuthenticationSetsRequest) (*DeleteAllAuthenticationSetsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumAuthenticationSets(context.Context, *EnumAuthenticationSetsRequest) (*EnumAuthenticationSetsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddCryptoSet(context.Context, *AddCryptoSetRequest) (*AddCryptoSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetCryptoSet(context.Context, *SetCryptoSetRequest) (*SetCryptoSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeleteCryptoSet(context.Context, *DeleteCryptoSetRequest) (*DeleteCryptoSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeleteAllCryptoSets(context.Context, *DeleteAllCryptoSetsRequest) (*DeleteAllCryptoSetsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumCryptoSets(context.Context, *EnumCryptoSetsRequest) (*EnumCryptoSetsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumPhase1SAs(context.Context, *EnumPhase1SAsRequest) (*EnumPhase1SAsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumPhase2SAs(context.Context, *EnumPhase2SAsRequest) (*EnumPhase2SAsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeletePhase1SAs(context.Context, *DeletePhase1SAsRequest) (*DeletePhase1SAsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeletePhase2SAs(context.Context, *DeletePhase2SAsRequest) (*DeletePhase2SAsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumProducts(context.Context, *EnumProductsRequest) (*EnumProductsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddMainModeRule(context.Context, *AddMainModeRuleRequest) (*AddMainModeRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetMainModeRule(context.Context, *SetMainModeRuleRequest) (*SetMainModeRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeleteMainModeRule(context.Context, *DeleteMainModeRuleRequest) (*DeleteMainModeRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) DeleteAllMainModeRules(context.Context, *DeleteAllMainModeRulesRequest) (*DeleteAllMainModeRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumMainModeRules(context.Context, *EnumMainModeRulesRequest) (*EnumMainModeRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryFirewallRules(context.Context, *QueryFirewallRulesRequest) (*QueryFirewallRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryConnectionSecurityRules210(context.Context, *QueryConnectionSecurityRules210Request) (*QueryConnectionSecurityRules210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryMainModeRules(context.Context, *QueryMainModeRulesRequest) (*QueryMainModeRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryAuthenticationSets(context.Context, *QueryAuthenticationSetsRequest) (*QueryAuthenticationSetsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryCryptoSets(context.Context, *QueryCryptoSetsRequest) (*QueryCryptoSetsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumNetworks(context.Context, *EnumNetworksRequest) (*EnumNetworksResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumAdapters(context.Context, *EnumAdaptersRequest) (*EnumAdaptersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) GetGlobalConfig210(context.Context, *GetGlobalConfig210Request) (*GetGlobalConfig210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) GetConfig210(context.Context, *GetConfig210Request) (*GetConfig210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddFirewallRule210(context.Context, *AddFirewallRule210Request) (*AddFirewallRule210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetFirewallRule210(context.Context, *SetFirewallRule210Request) (*SetFirewallRule210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumFirewallRules210(context.Context, *EnumFirewallRules210Request) (*EnumFirewallRules210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddConnectionSecurityRule210(context.Context, *AddConnectionSecurityRule210Request) (*AddConnectionSecurityRule210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetConnectionSecurityRule210(context.Context, *SetConnectionSecurityRule210Request) (*SetConnectionSecurityRule210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumConnectionSecurityRules210(context.Context, *EnumConnectionSecurityRules210Request) (*EnumConnectionSecurityRules210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddAuthenticationSet210(context.Context, *AddAuthenticationSet210Request) (*AddAuthenticationSet210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetAuthenticationSet210(context.Context, *SetAuthenticationSet210Request) (*SetAuthenticationSet210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumAuthenticationSets210(context.Context, *EnumAuthenticationSets210Request) (*EnumAuthenticationSets210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddCryptoSet210(context.Context, *AddCryptoSet210Request) (*AddCryptoSet210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetCryptoSet210(context.Context, *SetCryptoSet210Request) (*SetCryptoSet210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumCryptoSets210(context.Context, *EnumCryptoSets210Request) (*EnumCryptoSets210Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddConnectionSecurityRule220(context.Context, *AddConnectionSecurityRule220Request) (*AddConnectionSecurityRule220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetConnectionSecurityRule220(context.Context, *SetConnectionSecurityRule220Request) (*SetConnectionSecurityRule220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumConnectionSecurityRules220(context.Context, *EnumConnectionSecurityRules220Request) (*EnumConnectionSecurityRules220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryConnectionSecurityRules220(context.Context, *QueryConnectionSecurityRules220Request) (*QueryConnectionSecurityRules220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddAuthenticationSet220(context.Context, *AddAuthenticationSet220Request) (*AddAuthenticationSet220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetAuthenticationSet220(context.Context, *SetAuthenticationSet220Request) (*SetAuthenticationSet220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumAuthenticationSets220(context.Context, *EnumAuthenticationSets220Request) (*EnumAuthenticationSets220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryAuthenticationSets220(context.Context, *QueryAuthenticationSets220Request) (*QueryAuthenticationSets220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddFirewallRule220(context.Context, *AddFirewallRule220Request) (*AddFirewallRule220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetFirewallRule220(context.Context, *SetFirewallRule220Request) (*SetFirewallRule220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumFirewallRules220(context.Context, *EnumFirewallRules220Request) (*EnumFirewallRules220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryFirewallRules220(context.Context, *QueryFirewallRules220Request) (*QueryFirewallRules220Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddFirewallRule224(context.Context, *AddFirewallRule224Request) (*AddFirewallRule224Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetFirewallRule224(context.Context, *SetFirewallRule224Request) (*SetFirewallRule224Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumFirewallRules224(context.Context, *EnumFirewallRules224Request) (*EnumFirewallRules224Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryFirewallRules224(context.Context, *QueryFirewallRules224Request) (*QueryFirewallRules224Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddFirewallRule225(context.Context, *AddFirewallRule225Request) (*AddFirewallRule225Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetFirewallRule225(context.Context, *SetFirewallRule225Request) (*SetFirewallRule225Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumFirewallRules225(context.Context, *EnumFirewallRules225Request) (*EnumFirewallRules225Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryFirewallRules225(context.Context, *QueryFirewallRules225Request) (*QueryFirewallRules225Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddFirewallRule226(context.Context, *AddFirewallRule226Request) (*AddFirewallRule226Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetFirewallRule226(context.Context, *SetFirewallRule226Request) (*SetFirewallRule226Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumFirewallRules226(context.Context, *EnumFirewallRules226Request) (*EnumFirewallRules226Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryFirewallRules226(context.Context, *QueryFirewallRules226Request) (*QueryFirewallRules226Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddFirewallRule227(context.Context, *AddFirewallRule227Request) (*AddFirewallRule227Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetFirewallRule227(context.Context, *SetFirewallRule227Request) (*SetFirewallRule227Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumFirewallRules227(context.Context, *EnumFirewallRules227Request) (*EnumFirewallRules227Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryFirewallRules227(context.Context, *QueryFirewallRules227Request) (*QueryFirewallRules227Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) AddFirewallRule231(context.Context, *AddFirewallRule231Request) (*AddFirewallRule231Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) SetFirewallRule231(context.Context, *SetFirewallRule231Request) (*SetFirewallRule231Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) EnumFirewallRules231(context.Context, *EnumFirewallRules231Request) (*EnumFirewallRules231Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteFwServer) QueryFirewallRules231(context.Context, *QueryFirewallRules231Request) (*QueryFirewallRules231Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteFwServer = (*UnimplementedRemoteFwServer)(nil)
