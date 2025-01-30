package lsarpc

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
	GoPackage = "lsad"
)

var (
	// Syntax UUID
	LsarpcSyntaxUUID = &uuid.UUID{TimeLow: 0x12345778, TimeMid: 0x1234, TimeHiAndVersion: 0xabcd, ClockSeqHiAndReserved: 0xef, ClockSeqLow: 0x0, Node: [6]uint8{0x1, 0x23, 0x45, 0x67, 0x89, 0xab}}
	// Syntax ID
	LsarpcSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: LsarpcSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// lsarpc interface.
type LsarpcClient interface {

	// The LsarClose method frees the resources held by a context handle that was opened
	// earlier. After response, the context handle will no longer be usable, and any subsequent
	// uses of this handle will fail.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------+-----------------------------------------+
	//	|              RETURN              |                                         |
	//	|            VALUE/CODE            |               DESCRIPTION               |
	//	|                                  |                                         |
	//	+----------------------------------+-----------------------------------------+
	//	+----------------------------------+-----------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS        | The request was successfully completed. |
	//	+----------------------------------+-----------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE | ObjectHandle is not a valid handle.     |
	//	+----------------------------------+-----------------------------------------+
	Close(context.Context, *CloseRequest, ...dcerpc.CallOption) (*CloseResponse, error)

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// The LsarEnumeratePrivileges method is invoked to enumerate all privileges known to
	// the system. This method can be called multiple times to return its output in fragments.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation.              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000105 STATUS_MORE_ENTRIES      | More information is available to successive calls.                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8000001A STATUS_NO_MORE_ENTRIES   | No more entries are available from the enumeration.                              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the parameters supplied was invalid. This can happen if EnumerationBuffer |
	//	|                                     | is NULL or EnumerationContext is NULL.                                           |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | PolicyHandle is not a valid handle.                                              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	EnumeratePrivileges(context.Context, *EnumeratePrivilegesRequest, ...dcerpc.CallOption) (*EnumeratePrivilegesResponse, error)

	// The LsarQuerySecurityObject method is invoked to query security information that
	// is assigned to a database object. It returns the security descriptor of the object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------+---------------------------------------------------------------------+
	//	|              RETURN              |                                                                     |
	//	|            VALUE/CODE            |                             DESCRIPTION                             |
	//	|                                  |                                                                     |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS        | The request was successfully completed.                             |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED  | The caller does not have the permissions to perform this operation. |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0xC00000BB STATUS_NOT_SUPPORTED  | The request is not supported.                                       |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE | ObjectHandle is not a valid handle.                                 |
	//	+----------------------------------+---------------------------------------------------------------------+
	QuerySecurityObject(context.Context, *QuerySecurityObjectRequest, ...dcerpc.CallOption) (*QuerySecurityObjectResponse, error)

	// The LsarSetSecurityObject method is invoked to set a security descriptor on an object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                | The request was successfully completed.                                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000009A STATUS_INSUFFICIENT_RESOURCES | There are insufficient resources to complete the request.                        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED          | The caller does not have the permissions to perform this operation.              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000079 STATUS_INVALID_SECURITY_DESCR | The supplied security descriptor is invalid.                                     |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER      | One of the parameters supplied was invalid. For instance, SecurityDescriptor is  |
	//	|                                          | NULL.                                                                            |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00000BB STATUS_NOT_SUPPORTED          | The operation is not supported for this object.                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE         | ObjectHandle is not a valid handle.                                              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	SetSecurityObject(context.Context, *SetSecurityObjectRequest, ...dcerpc.CallOption) (*SetSecurityObjectResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// The LsarOpenPolicy method is exactly the same as LsarOpenPolicy2, except that the
	// SystemName parameter in this function, because of its syntactic definition, contains
	// only one character instead of a full string. This SystemName parameter does not have
	// any effect on message processing in any environment. It MUST be ignored.
	OpenPolicy(context.Context, *OpenPolicyRequest, ...dcerpc.CallOption) (*OpenPolicyResponse, error)

	// The LsarQueryInformationPolicy method is invoked to query values that represent the
	// server's information policy.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing below.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                | The request was successfully completed.                                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000009A STATUS_INSUFFICIENT_RESOURCES | There are insufficient resources to complete the request.                        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED          | The caller does not have the permissions to perform the operation.               |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER      | One of the parameters is incorrect. For instance, this can happen if             |
	//	|                                          | InformationClass is out of range or if PolicyInformation is NULL.                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE         | PolicyHandle is not a valid handle.                                              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	QueryInformationPolicy(context.Context, *QueryInformationPolicyRequest, ...dcerpc.CallOption) (*QueryInformationPolicyResponse, error)

	// The LsarSetInformationPolicy method is invoked to set a policy on the server.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation.              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the parameters is incorrect. For instance, this can happen if             |
	//	|                                     | InformationClass is not supported or some of the supplied policy data is         |
	//	|                                     | invalid.                                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000002 STATUS_NOT_IMPLEMENTED   | This information class cannot be set.                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | PolicyHandle is not a valid handle.                                              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	SetInformationPolicy(context.Context, *SetInformationPolicyRequest, ...dcerpc.CallOption) (*SetInformationPolicyResponse, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// The LsarCreateAccount method is invoked to create a new account object in the server's
	// database.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                     |
	//	|               VALUE/CODE                |                             DESCRIPTION                             |
	//	|                                         |                                                                     |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                             |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation. |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000035 STATUS_OBJECT_NAME_COLLISION | An account with this SID already exists.                            |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                 |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | Some of the parameters supplied were invalid.                       |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	CreateAccount(context.Context, *CreateAccountRequest, ...dcerpc.CallOption) (*CreateAccountResponse, error)

	// The LsarEnumerateAccounts method is invoked to request a list of account objects
	// in the server's database. The method can be called multiple times to return its output
	// in fragments.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	|              RETURN               |                                                                     |
	//	|            VALUE/CODE             |                             DESCRIPTION                             |
	//	|                                   |                                                                     |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS         | The request was successfully completed.                             |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED   | The caller does not have the permissions to perform this operation. |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0x00000105 STATUS_MORE_ENTRIES    | More information is available to successive calls.                  |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0x8000001A STATUS_NO_MORE_ENTRIES | No more entries are available from the enumeration.                 |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE  | PolicyHandle is not a valid handle.                                 |
	//	+-----------------------------------+---------------------------------------------------------------------+
	EnumerateAccounts(context.Context, *EnumerateAccountsRequest, ...dcerpc.CallOption) (*EnumerateAccountsResponse, error)

	// The LsarCreateTrustedDomain method is invoked to create an object of type trusted
	// domain in the server's database.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                     RETURN                      |                                                                                  |
	//	|                   VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                 |                                                                                  |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                       | The request was successfully completed.                                          |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED                 | The caller does not have the permissions to perform this operation.              |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER             | One of the supplied arguments is invalid.                                        |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000300 STATUS_NOT_SUPPORTED_ON_SBS          | The operation is not supported on a particular product.<101>                     |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00002B1 STATUS_DIRECTORY_SERVICE_REQUIRED    | The Active Directory service was not available on the server.                    |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000078 STATUS_INVALID_SID                   | The security identifier of the trusted domain is not valid.                      |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00002E9 STATUS_CURRENT_DOMAIN_NOT_ALLOWED    | Trust cannot be established with the current domain.                             |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000035 STATUS_OBJECT_NAME_COLLISION         | Another TDO already exists that matches some of the identifying information of   |
	//	|                                                 | the supplied information.                                                        |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE                | PolicyHandle is not a valid handle.                                              |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000401 STATUS_PER_USER_TRUST_QUOTA_EXCEEDED | The caller's quota for the maximum number of TDOs that can be created by control |
	//	|                                                 | access right Create-Inbound-Trust is exceeded.                                   |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000402 STATUS_ALL_USER_TRUST_QUOTA_EXCEEDED | The combined users' quota for the maximum number of TDOs that can be created by  |
	//	|                                                 | control access right Create-Inbound-Trust is exceeded.                           |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	CreateTrustedDomain(context.Context, *CreateTrustedDomainRequest, ...dcerpc.CallOption) (*CreateTrustedDomainResponse, error)

	// The LsarEnumerateTrustedDomains method is invoked to request a list of trusted domain
	// objects in the server's database. The method can be called multiple times to return
	// its output in fragments.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	|              RETURN               |                                                                     |
	//	|            VALUE/CODE             |                             DESCRIPTION                             |
	//	|                                   |                                                                     |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS         | The request was successfully completed.                             |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED   | The caller does not have the permissions to perform this operation. |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000105 STATUS_MORE_ENTRIES    | More information is available to successive calls.                  |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0xC000001A STATUS_NO_MORE_ENTRIES | No more entries are available from the enumeration.                 |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE  | PolicyHandle is not a valid handle.                                 |
	//	+-----------------------------------+---------------------------------------------------------------------+
	EnumerateTrustedDomains(context.Context, *EnumerateTrustedDomainsRequest, ...dcerpc.CallOption) (*EnumerateTrustedDomainsResponse, error)

	// Lsar_LSA_TM_14 operation.
	// LSATM14

	// Lsar_LSA_TM_15 operation.
	// LSATM15

	// The LsarCreateSecret method is invoked to create a new secret object in the server's
	// database.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation.              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | One of the supplied parameters is invalid. This can happen, for example, if      |
	//	|                                         | SecretHandle is NULL or if SecretName is not a valid name for a secret object.   |
	//	|                                         | Secret naming rules are specified in the processing rules shown below for the    |
	//	|                                         | SecretName parameter.                                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000035 STATUS_OBJECT_NAME_COLLISION | The secret object by the specified name already exists.                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000106 STATUS_NAME_TOO_LONG         | The length of specified secret name exceeds the maximum set by the server.       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	CreateSecret(context.Context, *CreateSecretRequest, ...dcerpc.CallOption) (*CreateSecretResponse, error)

	// The LsarOpenAccount method is invoked to obtain a handle to an account object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation.              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | Some of the parameters supplied are incorrect.  For instance, this can happen    |
	//	|                                         | when AccountSid is NULL.                                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000034 STATUS_OBJECT_NAME_NOT_FOUND | An account with this SID does not exist in the server's database.                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	OpenAccount(context.Context, *OpenAccountRequest, ...dcerpc.CallOption) (*OpenAccountResponse, error)

	// The LsarEnumeratePrivilegesAccount method is invoked to retrieve a list of privileges
	// granted to an account on the server.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+------------------------------------------+---------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                     |
	//	|                VALUE/CODE                |                             DESCRIPTION                             |
	//	|                                          |                                                                     |
	//	+------------------------------------------+---------------------------------------------------------------------+
	//	+------------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                | The request was successfully completed.                             |
	//	+------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000009A STATUS_INSUFFICIENT_RESOURCES | There are insufficient resources to complete the request.           |
	//	+------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED          | The caller does not have the permissions to perform this operation. |
	//	+------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE         | AccountHandle is not a valid handle.                                |
	//	+------------------------------------------+---------------------------------------------------------------------+
	EnumeratePrivilegesAccount(context.Context, *EnumeratePrivilegesAccountRequest, ...dcerpc.CallOption) (*EnumeratePrivilegesAccountResponse, error)

	// The LsarAddPrivilegesToAccount method is invoked to add new privileges to an existing
	// account object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+-----------------------------------------------------------------+
	//	|               RETURN                |                                                                 |
	//	|             VALUE/CODE              |                           DESCRIPTION                           |
	//	|                                     |                                                                 |
	//	+-------------------------------------+-----------------------------------------------------------------+
	//	+-------------------------------------+-----------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                         |
	//	+-------------------------------------+-----------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have permissions to perform this operation. |
	//	+-------------------------------------+-----------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | Some of the parameters supplied were invalid.                   |
	//	+-------------------------------------+-----------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | AccountHandle is not a valid handle.                            |
	//	+-------------------------------------+-----------------------------------------------------------------+
	AddPrivilegesToAccount(context.Context, *AddPrivilegesToAccountRequest, ...dcerpc.CallOption) (*AddPrivilegesToAccountResponse, error)

	// The LsarRemovePrivilegesFromAccount method is invoked to remove privileges from an
	// account object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | Some of the parameters supplied were invalid.                       |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | AccountHandle is not a valid handle.                                |
	//	+-------------------------------------+---------------------------------------------------------------------+
	RemovePrivilegesFromAccount(context.Context, *RemovePrivilegesFromAccountRequest, ...dcerpc.CallOption) (*RemovePrivilegesFromAccountResponse, error)

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire

	// Opnum22NotUsedOnWire operation.
	// Opnum22NotUsedOnWire

	// The LsarGetSystemAccessAccount method is invoked to retrieve system access account
	// flags for an account object. System access account flags are described as part of
	// the account object data model, as specified in section 3.1.1.3.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------+---------------------------------------------------------------------+
	//	|              RETURN              |                                                                     |
	//	|            VALUE/CODE            |                             DESCRIPTION                             |
	//	|                                  |                                                                     |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS        | The request was successfully completed.                             |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED  | The caller does not have the permissions to perform this operation. |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE | AccountHandle is not a valid handle.                                |
	//	+----------------------------------+---------------------------------------------------------------------+
	GetSystemAccessAccount(context.Context, *GetSystemAccessAccountRequest, ...dcerpc.CallOption) (*GetSystemAccessAccountResponse, error)

	// The LsarSetSystemAccessAccount method is invoked to set system access account flags
	// for an account object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the supplied parameters was invalid.                         |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | AccountHandle is not a valid handle.                                |
	//	+-------------------------------------+---------------------------------------------------------------------+
	SetSystemAccessAccount(context.Context, *SetSystemAccessAccountRequest, ...dcerpc.CallOption) (*SetSystemAccessAccountResponse, error)

	// The LsarOpenTrustedDomain method is invoked to obtain a handle to a trusted domain
	// object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                    | The request was successfully completed.                                          |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED              | The caller does not have the permissions to perform this operation.              |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER          | One of the supplied parameters is invalid. For instance, this can happen if the  |
	//	|                                              | security identifier TrustedDomainSid is not a valid domain security identifier.  |
	//	|                                              | Section 3.1.4.10 specifies data validation rules, including what constitutes a   |
	//	|                                              | valid domain security identifier.                                                |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE             | PolicyHandle is not a valid handle.                                              |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00000DF STATUS_NO_SUCH_DOMAIN             | The specified trusted domain object does not exist.                              |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00002B1 STATUS_DIRECTORY_SERVICE_REQUIRED | The Active Directory service was not available on the server.                    |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	OpenTrustedDomain(context.Context, *OpenTrustedDomainRequest, ...dcerpc.CallOption) (*OpenTrustedDomainResponse, error)

	// The LsarQueryInfoTrustedDomain method is invoked to retrieve information about the
	// trusted domain object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	|                RETURN                |                                                                     |
	//	|              VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                      |                                                                     |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS            | The request was successfully completed.                             |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED      | The caller does not have the permissions to perform this operation. |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER  | One of the arguments supplied to the function was invalid.          |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000003 STATUS_INVALID_INFO_CLASS | The InformationClass argument is outside the allowed range.         |
	//	+--------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE     | TrustedDomainHandle is not a valid handle.                          |
	//	+--------------------------------------+---------------------------------------------------------------------+
	QueryInfoTrustedDomain(context.Context, *QueryInfoTrustedDomainRequest, ...dcerpc.CallOption) (*QueryInfoTrustedDomainResponse, error)

	// The LsarSetInformationTrustedDomain method is invoked to set information on a trusted
	// domain object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	|                 RETURN                 |                                                                     |
	//	|               VALUE/CODE               |                             DESCRIPTION                             |
	//	|                                        |                                                                     |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS              | The request was successfully completed.                             |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED        | The caller does not have the permissions to perform this operation. |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER    | One of the arguments supplied to the function was invalid.          |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC00000DD STATUS_INVALID_DOMAIN_STATE | The domain is in the wrong state to perform the stated operation.   |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE       | TrustedDomainHandle is not a valid handle.                          |
	//	+----------------------------------------+---------------------------------------------------------------------+
	SetInformationTrustedDomain(context.Context, *SetInformationTrustedDomainRequest, ...dcerpc.CallOption) (*SetInformationTrustedDomainResponse, error)

	// The LsarOpenSecret method is invoked to obtain a handle to an existing secret object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                     |
	//	|               VALUE/CODE                |                             DESCRIPTION                             |
	//	|                                         |                                                                     |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                             |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation. |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000034 STATUS_OBJECT_NAME_NOT_FOUND | The secret with the specified name was not found.                   |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                 |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | Some of the parameters supplied were invalid.                       |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	OpenSecret(context.Context, *OpenSecretRequest, ...dcerpc.CallOption) (*OpenSecretResponse, error)

	// The LsarSetSecret method is invoked to set the current and old values of the secret
	// object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One or more of the supplied parameters was invalid.                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | SecretHandle is not a valid handle.                                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	SetSecret(context.Context, *SetSecretRequest, ...dcerpc.CallOption) (*SetSecretResponse, error)

	// The LsarQuerySecret method is invoked to retrieve the current and old (or previous)
	// value of the secret object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------+---------------------------------------------------------------------+
	//	|              RETURN              |                                                                     |
	//	|            VALUE/CODE            |                             DESCRIPTION                             |
	//	|                                  |                                                                     |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS        | The request was successfully completed.                             |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED  | The caller does not have the permissions to perform this operation. |
	//	+----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE | SecretHandle is not a valid handle.                                 |
	//	+----------------------------------+---------------------------------------------------------------------+
	QuerySecret(context.Context, *QuerySecretRequest, ...dcerpc.CallOption) (*QuerySecretResponse, error)

	// The LsarLookupPrivilegeValue method is invoked to map the name of a privilege into
	// a locally unique identifier (LUID) by which the privilege is known on the server.
	// The locally unique value of the privilege can then be used in subsequent calls to
	// other methods, such as LsarAddPrivilegesToAccount.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One or more of the supplied parameters was invalid.                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000060 STATUS_NO_SUCH_PRIVILEGE | The privilege name is not recognized by the server.                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | PolicyHandle is not a valid handle.                                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	LookupPrivilegeValue(context.Context, *LookupPrivilegeValueRequest, ...dcerpc.CallOption) (*LookupPrivilegeValueResponse, error)

	// The LsarLookupPrivilegeName method is invoked to map the LUID of a privilege into
	// a string name by which the privilege is known on the server.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000060 STATUS_NO_SUCH_PRIVILEGE | The supplied LUID is not recognized by the server.                  |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | PolicyHandle is not a valid handle.                                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	LookupPrivilegeName(context.Context, *LookupPrivilegeNameRequest, ...dcerpc.CallOption) (*LookupPrivilegeNameResponse, error)

	// The LsarLookupPrivilegeDisplayName method is invoked to map the name of a privilege
	// into a display text string in the caller's language.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One or more of the supplied parameters was invalid.                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000060 STATUS_NO_SUCH_PRIVILEGE | The supplied LUID is not recognized by the server.                  |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | PolicyHandle is not a valid handle.                                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	LookupPrivilegeDisplayName(context.Context, *LookupPrivilegeDisplayNameRequest, ...dcerpc.CallOption) (*LookupPrivilegeDisplayNameResponse, error)

	// The LsarDeleteObject method is invoked to delete an open account object, secret object,
	// or trusted domain object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One or more of the supplied parameters was invalid.                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | ObjectHandle is not a valid handle.                                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	DeleteObject(context.Context, *DeleteObjectRequest, ...dcerpc.CallOption) (*DeleteObjectResponse, error)

	// The LsarEnumerateAccountsWithUserRight method is invoked to return a list of account
	// objects that have the user right equal to the passed-in value.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | PolicyHandle is not a valid handle.                                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000060 STATUS_NO_SUCH_PRIVILEGE | The supplied name is not recognized by the server.                  |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the supplied arguments is invalid.                           |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x8000001A STATUS_NO_MORE_ENTRIES   | No account was found with the specified privilege.                  |
	//	+-------------------------------------+---------------------------------------------------------------------+
	EnumerateAccountsWithUserRight(context.Context, *EnumerateAccountsWithUserRightRequest, ...dcerpc.CallOption) (*EnumerateAccountsWithUserRightResponse, error)

	// The LsarEnumerateAccountRights method is invoked to retrieve a list of rights associated
	// with an existing account.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                     |
	//	|               VALUE/CODE                |                             DESCRIPTION                             |
	//	|                                         |                                                                     |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                             |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation. |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | One or more of the supplied parameters was invalid.                 |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000034 STATUS_OBJECT_NAME_NOT_FOUND | The specified account object does not exist.                        |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                 |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	EnumerateAccountRights(context.Context, *EnumerateAccountRightsRequest, ...dcerpc.CallOption) (*EnumerateAccountRightsResponse, error)

	// The LsarAddAccountRights method is invoked to add new rights to an account object.
	// If the account object does not exist, the system will attempt to create one.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One or more of the supplied parameters was invalid.                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000060 STATUS_NO_SUCH_PRIVILEGE | The rights supplied were not recognized.                            |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | PolicyHandle is not a valid handle.                                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	AddAccountRights(context.Context, *AddAccountRightsRequest, ...dcerpc.CallOption) (*AddAccountRightsResponse, error)

	// The LsarRemoveAccountRights method is invoked to remove rights from an account object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                     |
	//	|               VALUE/CODE                |                             DESCRIPTION                             |
	//	|                                         |                                                                     |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                             |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation. |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | One or more of the supplied parameters was invalid.                 |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000060 STATUS_NO_SUCH_PRIVILEGE     | The rights supplied were not recognized.                            |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000034 STATUS_OBJECT_NAME_NOT_FOUND | An account with this SID does not exist.                            |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                 |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC00000BB STATUS_NOT_SUPPORTED         | The operation is not supported by the server.                       |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	RemoveAccountRights(context.Context, *RemoveAccountRightsRequest, ...dcerpc.CallOption) (*RemoveAccountRightsResponse, error)

	// The LsarQueryTrustedDomainInfo method is invoked to retrieve information on a trusted
	// domain object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                     |
	//	|                  VALUE/CODE                  |                             DESCRIPTION                             |
	//	|                                              |                                                                     |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                    | The request was successfully completed.                             |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED              | The caller does not have the permissions to perform this operation. |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER          | One or more of the supplied parameters was invalid.                 |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000002 STATUS_NOT_IMPLEMENTED            | The specified information class is not supported.                   |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000003 STATUS_INVALID_INFO_CLASS         | The InformationClass argument is outside the allowed range.         |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE             | PolicyHandle is not a valid handle.                                 |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC00000DF STATUS_NO_SUCH_DOMAIN             | The specified trusted domain object does not exist.                 |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC00002B1 STATUS_DIRECTORY_SERVICE_REQUIRED | The Active Directory service was not available on the server.       |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	QueryTrustedDomainInfo(context.Context, *QueryTrustedDomainInfoRequest, ...dcerpc.CallOption) (*QueryTrustedDomainInfoResponse, error)

	// The LsarSetTrustedDomainInfo method is invoked to set information on a trusted domain
	// object. In some cases, if the trusted domain object does not exist, it will be created.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                     |
	//	|                  VALUE/CODE                  |                             DESCRIPTION                             |
	//	|                                              |                                                                     |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                    | The request was successfully completed.                             |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED              | The caller does not have the permissions to perform this operation. |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER          | One or more of the supplied parameters was invalid.                 |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE             | PolicyHandle is not a valid handle.                                 |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC00000DF STATUS_NO_SUCH_DOMAIN             | The specified trusted domain object does not exist.                 |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	//	| 0xC00002B1 STATUS_DIRECTORY_SERVICE_REQUIRED | The Active Directory service was not available on the server.       |
	//	+----------------------------------------------+---------------------------------------------------------------------+
	SetTrustedDomainInfo(context.Context, *SetTrustedDomainInfoRequest, ...dcerpc.CallOption) (*SetTrustedDomainInfoResponse, error)

	// The LsarDeleteTrustedDomain method is invoked to delete a trusted domain object (TDO).
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	|                       RETURN                       |                                                                                |
	//	|                     VALUE/CODE                     |                                  DESCRIPTION                                   |
	//	|                                                    |                                                                                |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                          | The request was successfully completed.                                        |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED                    | The caller does not have the permissions to perform this operation.            |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0xC00000DF STATUS_NO_SUCH_DOMAIN                   | The specified TDO does not exist.                                              |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER                | One or more of the supplied parameters was invalid.                            |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE                   | PolicyHandle is not a valid handle.                                            |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0xC00002B1 STATUS_DIRECTORY_SERVICE_REQUIRED       | The Active Directory service was not available on the server.                  |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0xC0000403 STATUS_USER_DELETE_TRUST_QUOTA_EXCEEDED | The caller's quota for the maximum allowed number of deleted TDOs is exceeded. |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	DeleteTrustedDomain(context.Context, *DeleteTrustedDomainRequest, ...dcerpc.CallOption) (*DeleteTrustedDomainResponse, error)

	// The LsarStorePrivateData method is invoked to store a secret value.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | PolicyHandle is not a valid handle.                                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One or more of the supplied parameters was invalid.                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	StorePrivateData(context.Context, *StorePrivateDataRequest, ...dcerpc.CallOption) (*StorePrivateDataResponse, error)

	// The LsarRetrievePrivateData method is invoked to retrieve a secret value.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                     |
	//	|               VALUE/CODE                |                             DESCRIPTION                             |
	//	|                                         |                                                                     |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                             |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation. |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | One of the supplied parameters was invalid.                         |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                 |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000034 STATUS_OBJECT_NAME_NOT_FOUND | The key with the specified name was not found.                      |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	RetrievePrivateData(context.Context, *RetrievePrivateDataRequest, ...dcerpc.CallOption) (*RetrievePrivateDataResponse, error)

	// The LsarOpenPolicy2 method opens a context handle to the RPC server. This is the
	// first function that MUST be called to contact the Local Security Authority (Domain
	// Policy) Remote Protocol database.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing below.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation.              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the supplied parameters is incorrect. For example, this can happen when   |
	//	|                                     | ObjectAttributes is NULL or DesiredAccess is zero.                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	OpenPolicy2(context.Context, *OpenPolicy2Request, ...dcerpc.CallOption) (*OpenPolicy2Response, error)

	// Lsar_LSA_TM_45 operation.
	// LSATM45

	// The LsarQueryInformationPolicy2 method is invoked to query values that represent
	// the server's security policy.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing below.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                | The request was successfully completed.                                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000009A STATUS_INSUFFICIENT_RESOURCES | There are insufficient resources to complete the request.                        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED          | The caller does not have the permissions to perform the operation.               |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER      | One of the parameters is incorrect. For instance, this can happen if             |
	//	|                                          | InformationClass is out of range or if PolicyInformation is NULL.                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE         | PolicyHandle is not a valid handle.                                              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	QueryInformationPolicy2(context.Context, *QueryInformationPolicy2Request, ...dcerpc.CallOption) (*QueryInformationPolicy2Response, error)

	// The LsarSetInformationPolicy2 method is invoked to set a policy on the server.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation.              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the parameters is incorrect. For instance, this can happen if             |
	//	|                                     | InformationClass is not supported or some of the supplied policy data is         |
	//	|                                     | invalid.                                                                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000002 STATUS_NOT_IMPLEMENTED   | This information class cannot be set.                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | PolicyHandle is not a valid handle.                                              |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	SetInformationPolicy2(context.Context, *SetInformationPolicy2Request, ...dcerpc.CallOption) (*SetInformationPolicy2Response, error)

	// The LsarQueryTrustedDomainInfoByName method is invoked to retrieve information about
	// a trusted domain object by its string name.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                       |
	//	|               VALUE/CODE                |                              DESCRIPTION                              |
	//	|                                         |                                                                       |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                               |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation.   |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | One of the supplied parameters was invalid.                           |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| 0xC0000034 STATUS_OBJECT_NAME_NOT_FOUND | The trusted domain object with the specified name could not be found. |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                   |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	QueryTrustedDomainInfoByName(context.Context, *QueryTrustedDomainInfoByNameRequest, ...dcerpc.CallOption) (*QueryTrustedDomainInfoByNameResponse, error)

	// The LsarSetTrustedDomainInfoByName method is invoked to set information about a trusted
	// domain object by its string name.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                       |
	//	|               VALUE/CODE                |                              DESCRIPTION                              |
	//	|                                         |                                                                       |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                               |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation.   |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | One of the supplied arguments is invalid.                             |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| 0xC0000034 STATUS_OBJECT_NAME_NOT_FOUND | The trusted domain object with the specified name could not be found. |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                   |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	SetTrustedDomainInfoByName(context.Context, *SetTrustedDomainInfoByNameRequest, ...dcerpc.CallOption) (*SetTrustedDomainInfoByNameResponse, error)

	// The LsarEnumerateTrustedDomainsEx method is invoked to enumerate trusted domain objects
	// in the server's database. The method is designed to be invoked multiple times to
	// retrieve the data in fragments.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	|              RETURN               |                                                                     |
	//	|            VALUE/CODE             |                             DESCRIPTION                             |
	//	|                                   |                                                                     |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS         | The request was successfully completed.                             |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED   | The caller does not have the permissions to perform this operation. |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0x8000001A STATUS_NO_MORE_ENTRIES | No more information is available.                                   |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0x00000105 STATUS_MORE_ENTRIES    | More information is available by calling this method again.         |
	//	+-----------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE  | PolicyHandle is not a valid handle.                                 |
	//	+-----------------------------------+---------------------------------------------------------------------+
	EnumerateTrustedDomainsEx(context.Context, *EnumerateTrustedDomainsExRequest, ...dcerpc.CallOption) (*EnumerateTrustedDomainsExResponse, error)

	// The LsarCreateTrustedDomainEx method is invoked to create a new trusted domain object
	// (TDO).
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                     RETURN                      |                                                                                  |
	//	|                   VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                 |                                                                                  |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                       | The request was successfully completed.                                          |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED                 | The caller does not have the permissions to perform this operation.              |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER             | One of the supplied arguments is invalid.                                        |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000300 STATUS_NOT_SUPPORTED_ON_SBS          | The operation is not supported on a particular product.<100>                     |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00000DD STATUS_INVALID_DOMAIN_STATE          | The operation cannot complete in the current state of the domain.                |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00002B1 STATUS_DIRECTORY_SERVICE_REQUIRED    | The Active Directory service was not available on the server.                    |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000078 STATUS_INVALID_SID                   | The security identifier of the trusted domain is not valid.                      |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00002E9 STATUS_CURRENT_DOMAIN_NOT_ALLOWED    | Trust cannot be established with the current domain.                             |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000035 STATUS_OBJECT_NAME_COLLISION         | Another TDO already exists that matches some of the identifying information of   |
	//	|                                                 | the supplied information.                                                        |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE                | PolicyHandle is not a valid handle.                                              |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000401 STATUS_PER_USER_TRUST_QUOTA_EXCEEDED | The caller's quota for the maximum number of TDOs that can be created by control |
	//	|                                                 | access right Create-Inbound-Trust is exceeded.                                   |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000402 STATUS_ALL_USER_TRUST_QUOTA_EXCEEDED | The combined users' quota for the maximum number of TDOs that can be created by  |
	//	|                                                 | control access right Create-Inbound-Trust is exceeded.                           |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	CreateTrustedDomainEx(context.Context, *CreateTrustedDomainExRequest, ...dcerpc.CallOption) (*CreateTrustedDomainExResponse, error)

	// Opnum52NotUsedOnWire operation.
	// Opnum52NotUsedOnWire

	// The LsarQueryDomainInformationPolicy method is invoked to retrieve policy settings
	// in addition to those exposed through LsarQueryInformationPolicy and LsarSetInformationPolicy2.
	// Despite the term "Domain" in the name of the method, processing of this message occurs
	// with local data, and furthermore, there is no requirement that this data have any
	// relationship with the LSA information in the domain to which the machine is joined.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                     |
	//	|               VALUE/CODE                |                             DESCRIPTION                             |
	//	|                                         |                                                                     |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                             |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation. |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | One of the supplied arguments was invalid.                          |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000034 STATUS_OBJECT_NAME_NOT_FOUND | No value has been set for this policy.                              |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                 |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	QueryDomainInformationPolicy(context.Context, *QueryDomainInformationPolicyRequest, ...dcerpc.CallOption) (*QueryDomainInformationPolicyResponse, error)

	// The LsarSetDomainInformationPolicy method is invoked to change policy settings in
	// addition to those exposed through LsarQueryInformationPolicy and LsarSetInformationPolicy2.
	// Despite the term "Domain" in the name of the method, processing of this message occurs
	// with local data. Also, there is no requirement that this data have any relationship
	// with the LSA information in the domain in which the machine is joined.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the following message processing.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the supplied arguments was invalid.                          |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE    | PolicyHandle is not a valid handle.                                 |
	//	+-------------------------------------+---------------------------------------------------------------------+
	SetDomainInformationPolicy(context.Context, *SetDomainInformationPolicyRequest, ...dcerpc.CallOption) (*SetDomainInformationPolicyResponse, error)

	// The LsarOpenTrustedDomainByName method is invoked to open a trusted domain object
	// handle by supplying the name of the trusted domain.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                     |
	//	|               VALUE/CODE                |                             DESCRIPTION                             |
	//	|                                         |                                                                     |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS               | The request was successfully completed.                             |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED         | The caller does not have the permissions to perform this operation. |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER     | One of the supplied arguments was invalid.                          |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000034 STATUS_OBJECT_NAME_NOT_FOUND | A trusted domain object by this name was not found.                 |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE        | PolicyHandle is not a valid handle.                                 |
	//	+-----------------------------------------+---------------------------------------------------------------------+
	OpenTrustedDomainByName(context.Context, *OpenTrustedDomainByNameRequest, ...dcerpc.CallOption) (*OpenTrustedDomainByNameResponse, error)

	// Opnum56NotUsedOnWire operation.
	// Opnum56NotUsedOnWire

	// Lsar_LSA_TM_57 operation.
	// LSATM57

	// Lsar_LSA_TM_58 operation.
	// LSATM58

	// The LsarCreateTrustedDomainEx2 method is invoked to create a new trusted domain object
	// (TDO).<96>
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                     RETURN                      |                                                                                  |
	//	|                   VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                 |                                                                                  |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS                       | The request was successfully completed.                                          |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED                 | The caller does not have the permissions to perform this operation.              |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER             | One of the supplied arguments is invalid.                                        |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000300 STATUS_NOT_SUPPORTED_ON_SBS          | The operation is not supported on a particular product.<97>                      |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00000DD STATUS_INVALID_DOMAIN_STATE          | The operation cannot complete in the current state of the domain.                |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00002B1 STATUS_DIRECTORY_SERVICE_REQUIRED    | The Active Directory service was not available on the server.                    |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000078 STATUS_INVALID_SID                   | The security identifier of the trusted domain is not valid.                      |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00002E9 STATUS_CURRENT_DOMAIN_NOT_ALLOWED    | Trust cannot be established with the current domain.                             |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000035 STATUS_OBJECT_NAME_COLLISION         | Another TDO already exists that matches some of the identifying information of   |
	//	|                                                 | the supplied information.                                                        |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE                | PolicyHandle is not a valid handle.                                              |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000401 STATUS_PER_USER_TRUST_QUOTA_EXCEEDED | The caller's quota for maximum number of TDOs that can be created by control     |
	//	|                                                 | access right Create-Inbound-Trust is exceeded.                                   |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000402 STATUS_ALL_USER_TRUST_QUOTA_EXCEEDED | The combined users' quota for maximum number of TDOs that can be created by      |
	//	|                                                 | control access right Create-Inbound-Trust is exceeded.                           |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	CreateTrustedDomainEx2(context.Context, *CreateTrustedDomainEx2Request, ...dcerpc.CallOption) (*CreateTrustedDomainEx2Response, error)

	// Opnum60NotUsedOnWire operation.
	// Opnum60NotUsedOnWire

	// Opnum61NotUsedOnWire operation.
	// Opnum61NotUsedOnWire

	// Opnum62NotUsedOnWire operation.
	// Opnum62NotUsedOnWire

	// Opnum63NotUsedOnWire operation.
	// Opnum63NotUsedOnWire

	// Opnum64NotUsedOnWire operation.
	// Opnum64NotUsedOnWire

	// Opnum65NotUsedOnWire operation.
	// Opnum65NotUsedOnWire

	// Opnum66NotUsedOnWire operation.
	// Opnum66NotUsedOnWire

	// Opnum67NotUsedOnWire operation.
	// Opnum67NotUsedOnWire

	// Lsar_LSA_TM_68 operation.
	// LSATM68

	// Opnum69NotUsedOnWire operation.
	// Opnum69NotUsedOnWire

	// Opnum70NotUsedOnWire operation.
	// Opnum70NotUsedOnWire

	// Opnum71NotUsedOnWire operation.
	// Opnum71NotUsedOnWire

	// Opnum72NotUsedOnWire operation.
	// Opnum72NotUsedOnWire

	// The LsarQueryForestTrustInformation method is invoked to retrieve information about
	// a trust relationship with another forest.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------------+-------------------------------------------------------------------------+
	//	|                 RETURN                 |                                                                         |
	//	|               VALUE/CODE               |                               DESCRIPTION                               |
	//	|                                        |                                                                         |
	//	+----------------------------------------+-------------------------------------------------------------------------+
	//	+----------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS              | The request was successfully completed.                                 |
	//	+----------------------------------------+-------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED        | The caller does not have the permissions to perform this operation.     |
	//	+----------------------------------------+-------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER    | One of the supplied arguments was invalid.                              |
	//	+----------------------------------------+-------------------------------------------------------------------------+
	//	| 0xC00000DD STATUS_INVALID_DOMAIN_STATE | The domain is in the wrong state of this operation.                     |
	//	+----------------------------------------+-------------------------------------------------------------------------+
	//	| 0xC00000DF STATUS_NO_SUCH_DOMAIN       | The TrustedDomainName is not a recognized domain name.                  |
	//	+----------------------------------------+-------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE       | PolicyHandle is not a valid handle.                                     |
	//	+----------------------------------------+-------------------------------------------------------------------------+
	//	| 0xC0000225 STATUS_NOT_FOUND            | Forest trust information does not exist for this trusted domain object. |
	//	+----------------------------------------+-------------------------------------------------------------------------+
	QueryForestTrustInformation(context.Context, *QueryForestTrustInformationRequest, ...dcerpc.CallOption) (*QueryForestTrustInformationResponse, error)

	// The LsarSetForestTrustInformation method is invoked to establish a trust relationship
	// with another forest by attaching a set of records called the forest trust information
	// to the trusted domain object.
	//
	// Return Values: The following is a summary of the return values that an implementation
	// MUST return, as specified by the message processing that follows.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                 |                                                                                  |
	//	|               VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS              | The request was successfully completed.                                          |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED        | The caller does not have the permissions to perform this operation.              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00000DD STATUS_INVALID_DOMAIN_STATE | The domain is not the root domain of the forest, or the forest is not at         |
	//	|                                        | DS_BEHAVIOR_WIN2003 forest functional level.                                     |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00000DE STATUS_INVALID_DOMAIN_ROLE  | The server is not the primary domain controller.                                 |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC00000DF STATUS_NO_SUCH_DOMAIN       | The trusted domain object with the name in the TrustedDomainName parameter does  |
	//	|                                        | not exist.                                                                       |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC0000008 STATUS_INVALID_HANDLE       | PolicyHandle is not a valid handle.                                              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER    | Some of the parameters supplied were invalid.                                    |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	SetForestTrustInformation(context.Context, *SetForestTrustInformationRequest, ...dcerpc.CallOption) (*SetForestTrustInformationResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Handle structure represents LSAPR_HANDLE RPC structure.
type Handle dcetypes.ContextHandle

func (o *Handle) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Handle) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Handle) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Handle) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// String structure represents STRING RPC structure.
//
// The STRING structure holds a counted string encoded in the OEM code page.
//
// This structure has no effect on message processing in any environment.
type String struct {
	// Length:  The length, in bytes, of the string pointed to by the Buffer member, not
	// including the terminating null character (if any).
	Length uint16 `idl:"name:Length" json:"length"`
	// MaximumLength:  This field contains the total number of bytes in the Buffer field.
	MaximumLength uint16 `idl:"name:MaximumLength" json:"maximum_length"`
	// Buffer:  A pointer to the actual string. If Length is greater than 0, this field
	// MUST contain a non-NULL value. If Length is 0, this field MUST be ignored.
	Buffer []byte `idl:"name:Buffer;size_is:(MaximumLength);length_is:(Length)" json:"buffer"`
}

func (o *String) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.MaximumLength == 0 {
		o.MaximumLength = uint16(len(o.Buffer))
	}
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint16(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *String) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumLength); err != nil {
		return err
	}
	if o.Buffer != nil || o.MaximumLength > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MaximumLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64(o.Length)
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
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Buffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *String) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumLength); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// ACL structure represents LSAPR_ACL RPC structure.
//
// The LSAPR_ACL structure defines the header of an access control list (ACL) that specifies
// a list of security protections applied to an object.
//
// This structure has no effect on message processing in any environment.
type ACL struct {
	// AclRevision:  The revision level of the LSAPR_ACL structure. This field MUST be ignored.
	// The content is unspecified, and no requirements are placed on its value because it
	// is never used.
	ACLRevision uint8 `idl:"name:AclRevision" json:"acl_revision"`
	// Sbz1:  This field is used for alignment. This field MUST be ignored. The content
	// is unspecified, and no requirements are placed on its value because it is never used.
	SBZ1 uint8 `idl:"name:Sbz1" json:"sbz1"`
	// AclSize:  The size of this structure in bytes, including the size of the variable
	// sized Dummy1 field.
	ACLSize uint16 `idl:"name:AclSize" json:"acl_size"`
	// Dummy1:  This field MUST be ignored. The content is unspecified, and no requirements
	// are placed on its value because it is never used.
	//
	// The ACL structure is specified in [MS-DTYP] section 2.4.5.
	_ []byte `idl:"name:Dummy1;size_is:((AclSize-4))"`
}

func (o *ACL) xxx_PreparePayload(ctx context.Context) error {
	if o.ACLSize < 4 {
		o.ACLSize = 4
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *ACL) NDRSizeInfo() []uint64 {
	dimSize1 := uint64((o.ACLSize - 4))
	if o.ACLSize < 4 {
		dimSize1 = uint64(0)
	}
	return []uint64{
		dimSize1,
	}
}
func (o *ACL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.ACLRevision); err != nil {
		return err
	}
	if err := w.WriteData(o.SBZ1); err != nil {
		return err
	}
	if err := w.WriteData(o.ACLSize); err != nil {
		return err
	}
	// reserved Dummy1
	for i1 := 0; uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACLRevision); err != nil {
		return err
	}
	if err := w.ReadData(&o.SBZ1); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACLSize); err != nil {
		return err
	}
	// reserved Dummy1
	var _Dummy1 []byte
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array _Dummy1", sizeInfo[0])
	}
	_Dummy1 = make([]byte, sizeInfo[0])
	for i1 := range _Dummy1 {
		i1 := i1
		if err := w.ReadData(&_Dummy1[i1]); err != nil {
			return err
		}
	}
	return nil
}

// SecurityDescriptor structure represents LSAPR_SECURITY_DESCRIPTOR RPC structure.
//
// The LSAPR_SECURITY_DESCRIPTOR structure defines an object's security descriptor.
//
// This structure has no effect on message processing in any environment.
type SecurityDescriptor struct {
	// Revision:  The security descriptor revision level. This field MUST be ignored. The
	// content is unspecified, and no requirements are placed on its value because it is
	// never used.
	Revision uint8 `idl:"name:Revision" json:"revision"`
	// Sbz1:  This field is used for alignment. This field MUST be ignored. The content
	// is unspecified, and no requirements are placed on its value because it is never used.
	SBZ1 uint8 `idl:"name:Sbz1" json:"sbz1"`
	// Control:  A set of flags (as specified in section 2.2.3.3) that qualify the meaning
	// of the security descriptor or its individual fields.
	Control uint16 `idl:"name:Control" json:"control"`
	// Owner:  A pointer to the RPC_SID structure that represents an object's owner as a
	// SID.
	Owner *dtyp.SID `idl:"name:Owner" json:"owner"`
	// Group:  A pointer to the RPC_SID structure that represents an object's primary group
	// as a SID.
	Group *dtyp.SID `idl:"name:Group" json:"group"`
	// Sacl:  A pointer to an ACL structure (as specified in 2.2.3.2) that contains a system
	// access control list (SACL).
	SACL *ACL `idl:"name:Sacl" json:"sacl"`
	// Dacl:  A pointer to an ACL structure that contains a discretionary access control
	// list (DACL).
	//
	// The SECURITY_DESCRIPTOR structure is specified in [MS-DTYP] section 2.4.6.
	DACL *ACL `idl:"name:Dacl" json:"dacl"`
}

func (o *SecurityDescriptor) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityDescriptor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Revision); err != nil {
		return err
	}
	if err := w.WriteData(o.SBZ1); err != nil {
		return err
	}
	if err := w.WriteData(o.Control); err != nil {
		return err
	}
	if o.Owner != nil {
		_ptr_Owner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Owner != nil {
				if err := o.Owner.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Owner, _ptr_Owner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Group != nil {
		_ptr_Group := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Group != nil {
				if err := o.Group.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Group, _ptr_Group); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SACL != nil {
		_ptr_Sacl := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SACL != nil {
				if err := o.SACL.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ACL{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SACL, _ptr_Sacl); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DACL != nil {
		_ptr_Dacl := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DACL != nil {
				if err := o.DACL.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ACL{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DACL, _ptr_Dacl); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityDescriptor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Revision); err != nil {
		return err
	}
	if err := w.ReadData(&o.SBZ1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Control); err != nil {
		return err
	}
	_ptr_Owner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Owner == nil {
			o.Owner = &dtyp.SID{}
		}
		if err := o.Owner.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Owner := func(ptr interface{}) { o.Owner = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.Owner, _s_Owner, _ptr_Owner); err != nil {
		return err
	}
	_ptr_Group := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Group == nil {
			o.Group = &dtyp.SID{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Group := func(ptr interface{}) { o.Group = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.Group, _s_Group, _ptr_Group); err != nil {
		return err
	}
	_ptr_Sacl := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SACL == nil {
			o.SACL = &ACL{}
		}
		if err := o.SACL.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sacl := func(ptr interface{}) { o.SACL = *ptr.(**ACL) }
	if err := w.ReadPointer(&o.SACL, _s_Sacl, _ptr_Sacl); err != nil {
		return err
	}
	_ptr_Dacl := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DACL == nil {
			o.DACL = &ACL{}
		}
		if err := o.DACL.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Dacl := func(ptr interface{}) { o.DACL = *ptr.(**ACL) }
	if err := w.ReadPointer(&o.DACL, _s_Dacl, _ptr_Dacl); err != nil {
		return err
	}
	return nil
}

// SecurityImpersonationLevel type represents SECURITY_IMPERSONATION_LEVEL RPC enumeration.
//
// The SECURITY_IMPERSONATION_LEVEL enumeration defines a set of values that specifies
// security impersonation levels. These levels govern the degree to which a server process
// can act on behalf of a client process.
//
// This enumeration has no effect on message processing in any environment.
type SecurityImpersonationLevel uint16

var (
	// SecurityAnonymous:  The server cannot obtain information about the client and cannot
	// impersonate the client.
	SecurityImpersonationLevelAnonymous SecurityImpersonationLevel = 0
	// SecurityIdentification:  The server can obtain information such as security identifiers
	// and privileges, but the server cannot impersonate the client.
	SecurityImpersonationLevelIdentification SecurityImpersonationLevel = 1
	// SecurityImpersonation:  The server can impersonate the client's security context
	// on its local system, but cannot impersonate the client when communicating with services
	// on remote systems.
	SecurityImpersonationLevelImpersonation SecurityImpersonationLevel = 2
	// SecurityDelegation:  The server can impersonate the client's security context when
	// communicating with services on remote systems.
	SecurityImpersonationLevelDelegation SecurityImpersonationLevel = 3
)

func (o SecurityImpersonationLevel) String() string {
	switch o {
	case SecurityImpersonationLevelAnonymous:
		return "SecurityImpersonationLevelAnonymous"
	case SecurityImpersonationLevelIdentification:
		return "SecurityImpersonationLevelIdentification"
	case SecurityImpersonationLevelImpersonation:
		return "SecurityImpersonationLevelImpersonation"
	case SecurityImpersonationLevelDelegation:
		return "SecurityImpersonationLevelDelegation"
	}
	return "Invalid"
}

// SecurityQualityOfService structure represents SECURITY_QUALITY_OF_SERVICE RPC structure.
//
// The SECURITY_QUALITY_OF_SERVICE structure defines information used to support client
// impersonation.
//
// This structure has no effect on message processing in any environment.
type SecurityQualityOfService struct {
	// Length:  This value MUST be ignored. No requirements are placed on its value because
	// it is never used.
	Length uint32 `idl:"name:Length" json:"length"`
	// ImpersonationLevel:  This field contains information (as specified in section 2.2.3.5)
	// given to the server about the client that describes how the server can represent,
	// or impersonate, the client.
	ImpersonationLevel SecurityImpersonationLevel `idl:"name:ImpersonationLevel" json:"impersonation_level"`
	// ContextTrackingMode:  This field specifies how the server tracks changes to the client's
	// security context (as specified in section 2.2.3.6).
	ContextTrackingMode uint8 `idl:"name:ContextTrackingMode" json:"context_tracking_mode"`
	// EffectiveOnly:  This field specifies whether the server can enable or disable privileges
	// and groups that the client's security context might include. This value MUST be TRUE
	// (nonzero) if the server has this right; otherwise, it MUST be FALSE (0).
	EffectiveOnly uint8 `idl:"name:EffectiveOnly" json:"effective_only"`
}

func (o *SecurityQualityOfService) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityQualityOfService) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ImpersonationLevel)); err != nil {
		return err
	}
	if err := w.WriteData(o.ContextTrackingMode); err != nil {
		return err
	}
	if err := w.WriteData(o.EffectiveOnly); err != nil {
		return err
	}
	return nil
}
func (o *SecurityQualityOfService) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ImpersonationLevel)); err != nil {
		return err
	}
	if err := w.ReadData(&o.ContextTrackingMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.EffectiveOnly); err != nil {
		return err
	}
	return nil
}

// ObjectAttributes structure represents LSAPR_OBJECT_ATTRIBUTES RPC structure.
//
// The LSAPR_OBJECT_ATTRIBUTES structure specifies an object and its properties. This
// structure MUST be ignored except for the RootDirectory field, which MUST be NULL.<17>
type ObjectAttributes struct {
	// Length:  The length of the structure, in bytes. This field is not used and MUST be
	// ignored.
	Length uint32 `idl:"name:Length" json:"length"`
	// RootDirectory:  This field is not used and MUST be NULL.
	RootDirectory []byte `idl:"name:RootDirectory" json:"root_directory"`
	// ObjectName:  A pointer to a STRING structure that contains the object name. This
	// field MUST be ignored. The content is unspecified and no requirements are placed
	// on its value because it is never used.
	ObjectName *String `idl:"name:ObjectName" json:"object_name"`
	// Attributes:  This field MUST be ignored. The content is unspecified and no requirements
	// are placed on its value because it is never used.
	Attributes uint32 `idl:"name:Attributes" json:"attributes"`
	// SecurityDescriptor:  This field contains the security attributes of the object. This
	// field MUST be ignored. The content is unspecified and no requirements are placed
	// on its value because it is never used.
	SecurityDescriptor *SecurityDescriptor `idl:"name:SecurityDescriptor" json:"security_descriptor"`
	// SecurityQualityOfService:  This field MUST be ignored. The content is unspecified
	// and no requirements are placed on its value because it is never used.
	SecurityQualityOfService *SecurityQualityOfService `idl:"name:SecurityQualityOfService" json:"security_quality_of_service"`
}

func (o *ObjectAttributes) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectAttributes) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.RootDirectory != nil {
		_ptr_RootDirectory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(len(o.RootDirectory))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.RootDirectory {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.RootDirectory[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.RootDirectory); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RootDirectory, _ptr_RootDirectory); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ObjectName != nil {
		_ptr_ObjectName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ObjectName != nil {
				if err := o.ObjectName.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectName, _ptr_ObjectName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil {
		_ptr_SecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SecurityDescriptor != nil {
				if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityDescriptor, _ptr_SecurityDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SecurityQualityOfService != nil {
		_ptr_SecurityQualityOfService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SecurityQualityOfService != nil {
				if err := o.SecurityQualityOfService.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SecurityQualityOfService{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityQualityOfService, _ptr_SecurityQualityOfService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectAttributes) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_RootDirectory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.RootDirectory", sizeInfo[0])
		}
		o.RootDirectory = make([]byte, sizeInfo[0])
		for i1 := range o.RootDirectory {
			i1 := i1
			if err := w.ReadData(&o.RootDirectory[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_RootDirectory := func(ptr interface{}) { o.RootDirectory = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.RootDirectory, _s_RootDirectory, _ptr_RootDirectory); err != nil {
		return err
	}
	_ptr_ObjectName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ObjectName == nil {
			o.ObjectName = &String{}
		}
		if err := o.ObjectName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ObjectName := func(ptr interface{}) { o.ObjectName = *ptr.(**String) }
	if err := w.ReadPointer(&o.ObjectName, _s_ObjectName, _ptr_ObjectName); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	_ptr_SecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SecurityDescriptor == nil {
			o.SecurityDescriptor = &SecurityDescriptor{}
		}
		if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(**SecurityDescriptor) }
	if err := w.ReadPointer(&o.SecurityDescriptor, _s_SecurityDescriptor, _ptr_SecurityDescriptor); err != nil {
		return err
	}
	_ptr_SecurityQualityOfService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SecurityQualityOfService == nil {
			o.SecurityQualityOfService = &SecurityQualityOfService{}
		}
		if err := o.SecurityQualityOfService.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SecurityQualityOfService := func(ptr interface{}) { o.SecurityQualityOfService = *ptr.(**SecurityQualityOfService) }
	if err := w.ReadPointer(&o.SecurityQualityOfService, _s_SecurityQualityOfService, _ptr_SecurityQualityOfService); err != nil {
		return err
	}
	return nil
}

// TrustInformation structure represents LSAPR_TRUST_INFORMATION RPC structure.
//
// The LSAPR_TRUST_INFORMATION structure identifies a domain.
type TrustInformation struct {
	// Name:  This field contains a name for the domain that is subject to the restrictions
	// of a NetBIOS name, as specified in [RFC1088]. This value SHOULD be used (by implementations
	// external to this protocol) to identify the domain via the NetBIOS, as specified in
	// [RFC1088].
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// Sid:  The SID of the domain. This field MUST NOT be NULL.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
}

func (o *TrustInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SID, _ptr_Sid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// PolicyInformationClass type represents POLICY_INFORMATION_CLASS RPC enumeration.
//
// The POLICY_INFORMATION_CLASS enumeration type contains values that specify the type
// of policy being queried or set by the client.
type PolicyInformationClass uint16

var (
	// PolicyAuditLogInformation: Information about audit log.
	PolicyInformationClassAuditLogInformation PolicyInformationClass = 1
	// PolicyAuditEventsInformation: Auditing options.
	PolicyInformationClassAuditEventsInformation PolicyInformationClass = 2
	// PolicyPrimaryDomainInformation: Primary domain information.
	PolicyInformationClassPrimaryDomainInformation PolicyInformationClass = 3
	// PolicyPdAccountInformation: Obsolete information class.
	PolicyInformationClassPDAccountInformation PolicyInformationClass = 4
	// PolicyAccountDomainInformation: Account domain information.
	PolicyInformationClassAccountDomainInformation PolicyInformationClass = 5
	// PolicyLsaServerRoleInformation: Server role information.
	PolicyInformationClassLSAServerRoleInformation PolicyInformationClass = 6
	// PolicyReplicaSourceInformation: Replica source information.
	PolicyInformationClassReplicaSourceInformation PolicyInformationClass = 7
	// PolicyInformationNotUsedOnWire: This enumeration value does not appear on the wire.
	PolicyInformationClassNotUsedOnWire PolicyInformationClass = 8
	// PolicyModificationInformation: Obsolete information class.
	PolicyInformationClassModificationInformation PolicyInformationClass = 9
	// PolicyAuditFullSetInformation: Obsolete information class.
	PolicyInformationClassAuditFullSetInformation PolicyInformationClass = 10
	// PolicyAuditFullQueryInformation: Audit log state.
	PolicyInformationClassAuditFullQueryInformation PolicyInformationClass = 11
	// PolicyDnsDomainInformation: DNS domain information.
	PolicyInformationClassDNSDomainInformation PolicyInformationClass = 12
	// PolicyDnsDomainInformationInt: DNS domain information.
	PolicyInformationClassDNSDomainInformationInt PolicyInformationClass = 13
	// PolicyLocalAccountDomainInformation: Local account domain information.
	PolicyInformationClassLocalAccountDomainInformation PolicyInformationClass = 14
	// PolicyMachineAccountInformation: Machine account information.
	PolicyInformationClassMachineAccountInformation PolicyInformationClass = 15
	// PolicyLastEntry: Not used in this protocol. Present to mark the end of the enumeration.
	//
	// The following citation contains a timeline of when each enumeration value was introduced.<20>
	//
	// The values in this enumeration are used to define the contents of the LSAPR_POLICY_INFORMATION
	// (section 2.2.4.2) union, where the structure associated with each enumeration value
	// is specified. The structure associated with each enumeration value defines the meaning
	// of that value to this protocol.
	PolicyInformationClassLastEntry PolicyInformationClass = 16
)

func (o PolicyInformationClass) String() string {
	switch o {
	case PolicyInformationClassAuditLogInformation:
		return "PolicyInformationClassAuditLogInformation"
	case PolicyInformationClassAuditEventsInformation:
		return "PolicyInformationClassAuditEventsInformation"
	case PolicyInformationClassPrimaryDomainInformation:
		return "PolicyInformationClassPrimaryDomainInformation"
	case PolicyInformationClassPDAccountInformation:
		return "PolicyInformationClassPDAccountInformation"
	case PolicyInformationClassAccountDomainInformation:
		return "PolicyInformationClassAccountDomainInformation"
	case PolicyInformationClassLSAServerRoleInformation:
		return "PolicyInformationClassLSAServerRoleInformation"
	case PolicyInformationClassReplicaSourceInformation:
		return "PolicyInformationClassReplicaSourceInformation"
	case PolicyInformationClassNotUsedOnWire:
		return "PolicyInformationClassNotUsedOnWire"
	case PolicyInformationClassModificationInformation:
		return "PolicyInformationClassModificationInformation"
	case PolicyInformationClassAuditFullSetInformation:
		return "PolicyInformationClassAuditFullSetInformation"
	case PolicyInformationClassAuditFullQueryInformation:
		return "PolicyInformationClassAuditFullQueryInformation"
	case PolicyInformationClassDNSDomainInformation:
		return "PolicyInformationClassDNSDomainInformation"
	case PolicyInformationClassDNSDomainInformationInt:
		return "PolicyInformationClassDNSDomainInformationInt"
	case PolicyInformationClassLocalAccountDomainInformation:
		return "PolicyInformationClassLocalAccountDomainInformation"
	case PolicyInformationClassMachineAccountInformation:
		return "PolicyInformationClassMachineAccountInformation"
	case PolicyInformationClassLastEntry:
		return "PolicyInformationClassLastEntry"
	}
	return "Invalid"
}

// PolicyAuditEventType type represents POLICY_AUDIT_EVENT_TYPE RPC enumeration.
type PolicyAuditEventType uint16

var (
	// AuditCategorySystem:  Manages auditing of system-related events
	PolicyAuditEventTypeAuditCategorySystem PolicyAuditEventType = 0
	// AuditCategoryLogon:  Manages auditing of account logon events
	PolicyAuditEventTypeAuditCategoryLogon PolicyAuditEventType = 1
	// AuditCategoryObjectAccess:  Manages auditing of object access events
	PolicyAuditEventTypeAuditCategoryObjectAccess PolicyAuditEventType = 2
	// AuditCategoryPrivilegeUse:  Manages auditing of privilege use events
	PolicyAuditEventTypeAuditCategoryPrivilegeUse PolicyAuditEventType = 3
	// AuditCategoryDetailedTracking:  Manages detailed auditing
	PolicyAuditEventTypeAuditCategoryDetailedTracking PolicyAuditEventType = 4
	// AuditCategoryPolicyChange:  Manages auditing of policy change events
	PolicyAuditEventTypeAuditCategoryPolicyChange PolicyAuditEventType = 5
	// AuditCategoryAccountManagement:  Manages auditing of account management events
	PolicyAuditEventTypeAuditCategoryAccountManagement PolicyAuditEventType = 6
	// AuditCategoryDirectoryServiceAccess:  Manages auditing of Active Directory access
	// events
	PolicyAuditEventTypeAuditCategoryDirectoryServiceAccess PolicyAuditEventType = 7
	// AuditCategoryAccountLogon:  Manages auditing of account logon events
	//
	// The values in this enumeration are used as indices into the EventAuditingOptions
	// field of the LSAPR_POLICY_AUDIT_EVENTS_INFO structure (see section 2.2.4.4).
	PolicyAuditEventTypeAuditCategoryAccountLogon PolicyAuditEventType = 8
)

func (o PolicyAuditEventType) String() string {
	switch o {
	case PolicyAuditEventTypeAuditCategorySystem:
		return "PolicyAuditEventTypeAuditCategorySystem"
	case PolicyAuditEventTypeAuditCategoryLogon:
		return "PolicyAuditEventTypeAuditCategoryLogon"
	case PolicyAuditEventTypeAuditCategoryObjectAccess:
		return "PolicyAuditEventTypeAuditCategoryObjectAccess"
	case PolicyAuditEventTypeAuditCategoryPrivilegeUse:
		return "PolicyAuditEventTypeAuditCategoryPrivilegeUse"
	case PolicyAuditEventTypeAuditCategoryDetailedTracking:
		return "PolicyAuditEventTypeAuditCategoryDetailedTracking"
	case PolicyAuditEventTypeAuditCategoryPolicyChange:
		return "PolicyAuditEventTypeAuditCategoryPolicyChange"
	case PolicyAuditEventTypeAuditCategoryAccountManagement:
		return "PolicyAuditEventTypeAuditCategoryAccountManagement"
	case PolicyAuditEventTypeAuditCategoryDirectoryServiceAccess:
		return "PolicyAuditEventTypeAuditCategoryDirectoryServiceAccess"
	case PolicyAuditEventTypeAuditCategoryAccountLogon:
		return "PolicyAuditEventTypeAuditCategoryAccountLogon"
	}
	return "Invalid"
}

// UnicodeString structure represents LSA_UNICODE_STRING RPC structure.
type UnicodeString dtyp.UnicodeString

func (o *UnicodeString) UnicodeString() *dtyp.UnicodeString { return (*dtyp.UnicodeString)(o) }

func (o *UnicodeString) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != "" && o.MaximumLength == 0 {
		o.MaximumLength = uint16((len(o.Buffer) * 2))
	}
	if o.Buffer != "" && o.Length == 0 {
		o.Length = uint16((len(o.Buffer) * 2))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumLength); err != nil {
		return err
	}
	if o.Buffer != "" || (o.MaximumLength/2) > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64((o.MaximumLength / 2))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64((o.Length / 2))
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
			_Buffer_buf := utf16.Encode([]rune(o.Buffer))
			if uint64(len(_Buffer_buf)) > sizeInfo[0] {
				_Buffer_buf = _Buffer_buf[:sizeInfo[0]]
			}
			for i1 := range _Buffer_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Buffer_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumLength); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*string) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// PolicyAuditLogInfo structure represents POLICY_AUDIT_LOG_INFO RPC structure.
//
// The POLICY_AUDIT_LOG_INFO structure contains information about the state of the audit
// log. The following structure corresponds to the PolicyAuditLogInformation information
// class.
type PolicyAuditLogInfo struct {
	// AuditLogPercentFull:  A measure of how full the audit log is, as a percentage.
	AuditLogPercentFull uint32 `idl:"name:AuditLogPercentFull" json:"audit_log_percent_full"`
	// MaximumLogSize:  The maximum size of the auditing log, in kilobytes (KB).
	MaximumLogSize uint32 `idl:"name:MaximumLogSize" json:"maximum_log_size"`
	// AuditRetentionPeriod:  The auditing log retention period (64-bit signed integer),
	// a 64-bit value that represents the number of 100-nanosecond intervals since January
	// 1, 1601, UTC. An audit record can be discarded if its time stamp predates the current
	// time minus the retention period.
	AuditRetentionPeriod *dtyp.LargeInteger `idl:"name:AuditRetentionPeriod" json:"audit_retention_period"`
	// AuditLogFullShutdownInProgress:  A Boolean flag; indicates whether or not a system
	// shutdown is being initiated due to the security audit log becoming full. This condition
	// occurs only if the system is configured to shut down when the log becomes full.
	//
	// After a shutdown has been initiated, this flag MUST be set to TRUE (nonzero). If
	// an administrator can correct the situation before the shutdown becomes irreversible,
	// this flag MUST be reset to FALSE (0).
	AuditLogFullShutdownInProgress uint8 `idl:"name:AuditLogFullShutdownInProgress" json:"audit_log_full_shutdown_in_progress"`
	// TimeToShutdown:  A 64-bit value that represents the number of 100-nanosecond intervals
	// since January 1, 1601, UTC. If the AuditLogFullShutdownInProgress flag is set, this
	// field MUST contain the time left before the shutdown becomes irreversible.
	TimeToShutdown *dtyp.LargeInteger `idl:"name:TimeToShutdown" json:"time_to_shutdown"`
	// NextAuditRecordId:  Not in use. This field SHOULD be set to zero when sent, and MUST
	// be ignored on receipt.
	NextAuditRecordID uint32 `idl:"name:NextAuditRecordId" json:"next_audit_record_id"`
}

func (o *PolicyAuditLogInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyAuditLogInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.AuditLogPercentFull); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumLogSize); err != nil {
		return err
	}
	if o.AuditRetentionPeriod != nil {
		if err := o.AuditRetentionPeriod.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AuditLogFullShutdownInProgress); err != nil {
		return err
	}
	if o.TimeToShutdown != nil {
		if err := o.TimeToShutdown.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NextAuditRecordID); err != nil {
		return err
	}
	return nil
}
func (o *PolicyAuditLogInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuditLogPercentFull); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumLogSize); err != nil {
		return err
	}
	if o.AuditRetentionPeriod == nil {
		o.AuditRetentionPeriod = &dtyp.LargeInteger{}
	}
	if err := o.AuditRetentionPeriod.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuditLogFullShutdownInProgress); err != nil {
		return err
	}
	if o.TimeToShutdown == nil {
		o.TimeToShutdown = &dtyp.LargeInteger{}
	}
	if err := o.TimeToShutdown.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.NextAuditRecordID); err != nil {
		return err
	}
	return nil
}

// PolicyLSAServerRole type represents POLICY_LSA_SERVER_ROLE RPC enumeration.
//
// The POLICY_LSA_SERVER_ROLE enumeration takes one of two possible values, depending
// on which capacity the account domain database is in: primary or backup. Certain operations
// of the protocol are allowed only against a primary account database. On non–domain
// controller machines, the account domain database is in primary state. On domain controllers,
// if the machine is the primary domain controller (PDC) role owner, then the account
// domain database is in primary state; otherwise, it is in backup state.
type PolicyLSAServerRole uint16

var (
	// PolicyServerRoleBackup:  A backup account database.
	PolicyLSAServerRoleBackup PolicyLSAServerRole = 2
	// PolicyServerRolePrimary:  A primary account database.
	PolicyLSAServerRolePrimary PolicyLSAServerRole = 3
)

func (o PolicyLSAServerRole) String() string {
	switch o {
	case PolicyLSAServerRoleBackup:
		return "PolicyLSAServerRoleBackup"
	case PolicyLSAServerRolePrimary:
		return "PolicyLSAServerRolePrimary"
	}
	return "Invalid"
}

// PolicyLSAServerRoleInfo structure represents POLICY_LSA_SERVER_ROLE_INFO RPC structure.
//
// The POLICY_LSA_SERVER_ROLE_INFO structure is used to allow callers to query and set
// whether the account domain database acts as the primary copy or backup copy. The
// following structure corresponds to the PolicyLsaServerRoleInformation information
// class.
type PolicyLSAServerRoleInfo struct {
	// LsaServerRole:  One of the values of the POLICY_LSA_SERVER_ROLE enumeration on return.
	LSAServerRole PolicyLSAServerRole `idl:"name:LsaServerRole" json:"lsa_server_role"`
}

func (o *PolicyLSAServerRoleInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyLSAServerRoleInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.LSAServerRole)); err != nil {
		return err
	}
	return nil
}
func (o *PolicyLSAServerRoleInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.LSAServerRole)); err != nil {
		return err
	}
	return nil
}

// PolicyModificationInfo structure represents POLICY_MODIFICATION_INFO RPC structure.
//
// The POLICY_MODIFICATION_INFO structure is obsolete and exists for backward compatibility
// purposes only. Callers of this protocol MUST NOT be able to set or retrieve this
// structure.
type PolicyModificationInfo struct {
	// ModifiedId:  A 64-bit unsigned integer that is incremented each time anything in
	// the Local Security Authority (LSA) database is modified.
	ModifiedID *dtyp.LargeInteger `idl:"name:ModifiedId" json:"modified_id"`
	// DatabaseCreationTime:  The date and time when the LSA database was created. It is
	// a 64-bit value that represents the number of 100-nanosecond intervals since January
	// 1, 1601, UTC.
	DatabaseCreationTime *dtyp.LargeInteger `idl:"name:DatabaseCreationTime" json:"database_creation_time"`
}

func (o *PolicyModificationInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyModificationInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ModifiedID != nil {
		if err := o.ModifiedID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DatabaseCreationTime != nil {
		if err := o.DatabaseCreationTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyModificationInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ModifiedID == nil {
		o.ModifiedID = &dtyp.LargeInteger{}
	}
	if err := o.ModifiedID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DatabaseCreationTime == nil {
		o.DatabaseCreationTime = &dtyp.LargeInteger{}
	}
	if err := o.DatabaseCreationTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyAuditFullSetInfo structure represents POLICY_AUDIT_FULL_SET_INFO RPC structure.
//
// The POLICY_AUDIT_FULL_SET_INFO structure contains information to set on the server
// that is controlling audit log behavior. The following structure corresponds to the
// PolicyAuditFullSetInformation information class. This information class is not supported.
type PolicyAuditFullSetInfo struct {
	// ShutDownOnFull:  A nonzero value means that the system MUST shut down when the event
	// log is full, while zero means that the system MUST NOT shut down when the event log
	// is full.
	ShutDownOnFull uint8 `idl:"name:ShutDownOnFull" json:"shut_down_on_full"`
}

func (o *PolicyAuditFullSetInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyAuditFullSetInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteData(o.ShutDownOnFull); err != nil {
		return err
	}
	return nil
}
func (o *PolicyAuditFullSetInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.ShutDownOnFull); err != nil {
		return err
	}
	return nil
}

// PolicyAuditFullQueryInfo structure represents POLICY_AUDIT_FULL_QUERY_INFO RPC structure.
//
// The POLICY_AUDIT_FULL_QUERY_INFO structure is used to query information about the
// state of the audit log on the server. The following structure corresponds to the
// PolicyAuditFullQueryInformation information class.
//
// This information class is obsolete and exists for backward compatibility purposes
// only.
type PolicyAuditFullQueryInfo struct {
	// ShutDownOnFull:  This field indicates whether the system MUST shut down when the
	// event log is full.
	ShutDownOnFull uint8 `idl:"name:ShutDownOnFull" json:"shut_down_on_full"`
	// LogIsFull:  This field indicates whether the event log is full or not.
	LogIsFull uint8 `idl:"name:LogIsFull" json:"log_is_full"`
}

func (o *PolicyAuditFullQueryInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyAuditFullQueryInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteData(o.ShutDownOnFull); err != nil {
		return err
	}
	if err := w.WriteData(o.LogIsFull); err != nil {
		return err
	}
	return nil
}
func (o *PolicyAuditFullQueryInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.ShutDownOnFull); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogIsFull); err != nil {
		return err
	}
	return nil
}

// PolicyDomainInformationClass type represents POLICY_DOMAIN_INFORMATION_CLASS RPC enumeration.
//
// The POLICY_DOMAIN_INFORMATION_CLASS enumeration type contains values that specify
// the type of policy being queried or set by the client.
//
// The values in this enumeration are used in defining the contents of the LSAPR_POLICY_DOMAIN_INFORMATION
// union.
type PolicyDomainInformationClass uint16

var (
	PolicyDomainInformationClassQualityOfServiceInformation PolicyDomainInformationClass = 1
	PolicyDomainInformationClassEFSInformation              PolicyDomainInformationClass = 2
	PolicyDomainInformationClassKerberosTicketInformation   PolicyDomainInformationClass = 3
)

func (o PolicyDomainInformationClass) String() string {
	switch o {
	case PolicyDomainInformationClassQualityOfServiceInformation:
		return "PolicyDomainInformationClassQualityOfServiceInformation"
	case PolicyDomainInformationClassEFSInformation:
		return "PolicyDomainInformationClassEFSInformation"
	case PolicyDomainInformationClassKerberosTicketInformation:
		return "PolicyDomainInformationClassKerberosTicketInformation"
	}
	return "Invalid"
}

// PolicyDomainKerberosTicketInfo structure represents POLICY_DOMAIN_KERBEROS_TICKET_INFO RPC structure.
//
// The POLICY_DOMAIN_KERBEROS_TICKET_INFO structure communicates policy information
// about the Kerberos security provider.
type PolicyDomainKerberosTicketInfo struct {
	// AuthenticationOptions:  Optional flags that affect validations performed during authentication.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 |  4  | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |     |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | V C | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+-----+---+---+---+---+---+---+---+
	//
	// Where the bits are defined as:
	//
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                 |                                                                                  |
	//	|                      VALUE                      |                                   DESCRIPTION                                    |
	//	|                                                 |                                                                                  |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	//	| VC POLICY_KERBEROS_VALIDATE_CLIENT (0x00000080) | This is the only flag that is currently defined. When this bit is set, the       |
	//	|                                                 | AuthenticationOptions flag of the Key Distribution Center (KDC) configuration    |
	//	|                                                 | setting will be set to POLICY_KERBEROS_VALIDATE_CLIENT (as described in          |
	//	|                                                 | [MS-KILE] section 3.3.1). All other bits SHOULD be set to 0 and ignored upon     |
	//	|                                                 | receipt.                                                                         |
	//	+-------------------------------------------------+----------------------------------------------------------------------------------+
	AuthenticationOptions uint32 `idl:"name:AuthenticationOptions" json:"authentication_options"`
	// MaxServiceTicketAge:  This is in units of 10^(-7) seconds. It corresponds to Maximum
	// ticket lifetime (as specified in [RFC4120] section 8.2) for service tickets only.
	// The default value of this setting is 10 hours.
	MaxServiceTicketAge *dtyp.LargeInteger `idl:"name:MaxServiceTicketAge" json:"max_service_ticket_age"`
	// MaxTicketAge:  This is in units of 10^(-7) seconds. It corresponds to the Maximum
	// ticket lifetime (as specified in [RFC4120] section 8.2) for ticket-granting ticket
	// (TGT) only. The default value of this setting is 10 hours.
	MaxTicketAge *dtyp.LargeInteger `idl:"name:MaxTicketAge" json:"max_ticket_age"`
	// MaxRenewAge:  This is in units of 10^(-7) seconds. It corresponds to the Maximum
	// renewable lifetime, as specified in [RFC4120] section 8.2. The default value of this
	// setting is one week.
	MaxRenewAge *dtyp.LargeInteger `idl:"name:MaxRenewAge" json:"max_renew_age"`
	// MaxClockSkew:  This is in units of 10^(-7) seconds. It corresponds to the Acceptable
	// clock skew, as specified in [RFC4120] section 8.2. The default value of this setting
	// is five minutes.
	MaxClockSkew *dtyp.LargeInteger `idl:"name:MaxClockSkew" json:"max_clock_skew"`
	// Reserved:  The value of this field SHOULD be set to zero when sent or on receipt.
	_ *dtyp.LargeInteger `idl:"name:Reserved"`
}

func (o *PolicyDomainKerberosTicketInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyDomainKerberosTicketInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthenticationOptions); err != nil {
		return err
	}
	if o.MaxServiceTicketAge != nil {
		if err := o.MaxServiceTicketAge.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MaxTicketAge != nil {
		if err := o.MaxTicketAge.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MaxRenewAge != nil {
		if err := o.MaxRenewAge.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MaxClockSkew != nil {
		if err := o.MaxClockSkew.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved Reserved
	if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}
func (o *PolicyDomainKerberosTicketInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthenticationOptions); err != nil {
		return err
	}
	if o.MaxServiceTicketAge == nil {
		o.MaxServiceTicketAge = &dtyp.LargeInteger{}
	}
	if err := o.MaxServiceTicketAge.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.MaxTicketAge == nil {
		o.MaxTicketAge = &dtyp.LargeInteger{}
	}
	if err := o.MaxTicketAge.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.MaxRenewAge == nil {
		o.MaxRenewAge = &dtyp.LargeInteger{}
	}
	if err := o.MaxRenewAge.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.MaxClockSkew == nil {
		o.MaxClockSkew = &dtyp.LargeInteger{}
	}
	if err := o.MaxClockSkew.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved *dtyp.LargeInteger
	if _Reserved == nil {
		_Reserved = &dtyp.LargeInteger{}
	}
	if err := _Reserved.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedPOSIXOffsetInfo structure represents TRUSTED_POSIX_OFFSET_INFO RPC structure.
//
// The TRUSTED_POSIX_OFFSET_INFO structure communicates any offset necessary for POSIX
// compliance. The following structure corresponds to the TrustedPosixOffsetInformation
// information class.
type TrustedPOSIXOffsetInfo struct {
	// Offset:  The offset to use for the generation of POSIX IDs for users and groups,
	// as specified in "trustPosixOffset" in [MS-ADTS] section 6.1.6.7.14.
	Offset uint32 `idl:"name:Offset" json:"offset"`
}

func (o *TrustedPOSIXOffsetInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedPOSIXOffsetInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Offset); err != nil {
		return err
	}
	return nil
}
func (o *TrustedPOSIXOffsetInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offset); err != nil {
		return err
	}
	return nil
}

// TrustedInformationClass type represents TRUSTED_INFORMATION_CLASS RPC enumeration.
//
// The TRUSTED_INFORMATION_CLASS enumeration type contains values that specify the type
// of trusted domain information queried or set by the client.
type TrustedInformationClass uint16

var (
	// TrustedDomainNameInformation: The trusted domain information contains the LSAPR_TRUSTED_DOMAIN_NAME_INFO
	// structure specified in section 2.2.7.4.
	TrustedInformationClassDomainNameInformation TrustedInformationClass = 1
	// TrustedControllersInformation: The trusted domain information contains the LSAPR_TRUSTED_CONTROLLERS_INFO
	// structure specified in section 2.2.7.5.
	TrustedInformationClassControllersInformation TrustedInformationClass = 2
	// TrustedPosixOffsetInformation: The trusted domain information contains the TRUSTED_POSIX_OFFSET_INFO
	// structure specified in section 2.2.7.6.
	TrustedInformationClassPOSIXOffsetInformation TrustedInformationClass = 3
	// TrustedPasswordInformation: The trusted domain information contains the LSAPR_TRUSTED_PASSWORD_INFO
	// structure specified in section 2.2.7.7.
	TrustedInformationClassPasswordInformation TrustedInformationClass = 4
	// TrustedDomainInformationBasic: The trusted domain information contains the LSAPR_TRUSTED_DOMAIN_INFORMATION_BASIC
	// structure specified in section 2.2.7.8.
	TrustedInformationClassDomainInformationBasic TrustedInformationClass = 5
	// TrustedDomainInformationEx: The trusted domain information contains the LSAPR_TRUSTED_
	// DOMAIN_INFORMATION_EX structure specified in section 2.2.7.9.
	TrustedInformationClassDomainInformationEx TrustedInformationClass = 6
	// TrustedDomainAuthInformation: The trusted domain information contains the LSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION
	// structure specified in section 2.2.7.11.
	TrustedInformationClassDomainAuthInformation TrustedInformationClass = 7
	// TrustedDomainFullInformation: The trusted domain information contains the LSAPR_TRUSTED_DOMAIN_FULL_INFORMATION
	// structure specified in section 2.2.7.13.
	TrustedInformationClassDomainFullInformation TrustedInformationClass = 8
	// TrustedDomainAuthInformationInternal: The trusted domain information contains the
	// LSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION_INTERNAL structure specified in section 2.2.7.12.
	TrustedInformationClassDomainAuthInformationInternal TrustedInformationClass = 9
	// TrustedDomainFullInformationInternal: The trusted domain information contains the
	// LSAPR_TRUSTED_DOMAIN_FULL_INFORMATION_INTERNAL structure specified in section 2.2.7.14.
	TrustedInformationClassDomainFullInformationInternal TrustedInformationClass = 10
	// TrustedDomainInformationEx2Internal: The trusted domain information contains the
	// LSAPR_TRUSTED_DOMAIN_INFORMATION_EX2 structure specified in section 2.2.7.10.
	TrustedInformationClassDomainInformationEx2Internal TrustedInformationClass = 11
	// TrustedDomainFullInformation2Internal: The trusted domain information contains the
	// LSAPR_TRUSTED_DOMAIN_FULL_INFORMATION2 structure specified in section 2.2.7.15.
	TrustedInformationClassDomainFullInformation2Internal TrustedInformationClass = 12
	// TrustedDomainSupportedEncryptionTypes: The trusted domain information contains the
	// TRUSTED_DOMAIN_SUPPORTED_ENCRYPTION_TYPES structure specified in section 2.2.7.18.
	TrustedInformationClassDomainSupportedEncryptionTypes TrustedInformationClass = 13
)

func (o TrustedInformationClass) String() string {
	switch o {
	case TrustedInformationClassDomainNameInformation:
		return "TrustedInformationClassDomainNameInformation"
	case TrustedInformationClassControllersInformation:
		return "TrustedInformationClassControllersInformation"
	case TrustedInformationClassPOSIXOffsetInformation:
		return "TrustedInformationClassPOSIXOffsetInformation"
	case TrustedInformationClassPasswordInformation:
		return "TrustedInformationClassPasswordInformation"
	case TrustedInformationClassDomainInformationBasic:
		return "TrustedInformationClassDomainInformationBasic"
	case TrustedInformationClassDomainInformationEx:
		return "TrustedInformationClassDomainInformationEx"
	case TrustedInformationClassDomainAuthInformation:
		return "TrustedInformationClassDomainAuthInformation"
	case TrustedInformationClassDomainFullInformation:
		return "TrustedInformationClassDomainFullInformation"
	case TrustedInformationClassDomainAuthInformationInternal:
		return "TrustedInformationClassDomainAuthInformationInternal"
	case TrustedInformationClassDomainFullInformationInternal:
		return "TrustedInformationClassDomainFullInformationInternal"
	case TrustedInformationClassDomainInformationEx2Internal:
		return "TrustedInformationClassDomainInformationEx2Internal"
	case TrustedInformationClassDomainFullInformation2Internal:
		return "TrustedInformationClassDomainFullInformation2Internal"
	case TrustedInformationClassDomainSupportedEncryptionTypes:
		return "TrustedInformationClassDomainSupportedEncryptionTypes"
	}
	return "Invalid"
}

// ForestTrustRecordType type represents LSA_FOREST_TRUST_RECORD_TYPE RPC enumeration.
//
// The LSA_FOREST_TRUST_RECORD_TYPE enumeration specifies a type of forest trust record.
type ForestTrustRecordType uint16

var (
	// ForestTrustTopLevelName:  The DNS name of the trusted forest. The structure used
	// for this record type is equivalent to LSA_UNICODE_STRING (section 2.2.2.3).
	ForestTrustRecordTypeTopLevelName ForestTrustRecordType = 0
	// ForestTrustTopLevelNameEx:  The DNS name of the trusted forest. This is the same
	// as ForestTrustTopLevelName. The structure used for this record type is equivalent
	// to LSA_UNICODE_STRING.
	ForestTrustRecordTypeTopLevelNameEx ForestTrustRecordType = 1
	// ForestTrustDomainInfo:  This field specifies a record containing identification
	// and name information.
	ForestTrustRecordTypeDomainInfo ForestTrustRecordType = 2
	// ForestTrustRecordTypeLast: The highest record value for this type is equal to the
	// ForestTrustScannerInfo enum value (4).
	ForestTrustRecordTypeLast ForestTrustRecordType = 2
)

func (o ForestTrustRecordType) String() string {
	switch o {
	case ForestTrustRecordTypeTopLevelName:
		return "ForestTrustRecordTypeTopLevelName"
	case ForestTrustRecordTypeTopLevelNameEx:
		return "ForestTrustRecordTypeTopLevelNameEx"
	case ForestTrustRecordTypeDomainInfo:
		return "ForestTrustRecordTypeDomainInfo"
	case ForestTrustRecordTypeLast:
		return "ForestTrustRecordTypeLast"
	}
	return "Invalid"
}

// ForestTrustBinaryData structure represents LSA_FOREST_TRUST_BINARY_DATA RPC structure.
//
// The LSA_FOREST_TRUST_BINARY_DATA structure is used to communicate a forest trust
// record. This structure is not used in the current version of the protocol.
type ForestTrustBinaryData struct {
	// Length:  The count of bytes in Buffer.<37>
	Length uint32 `idl:"name:Length" json:"length"`
	// Buffer:  The trust record. If the Length field has a value other than 0, this field
	// MUST NOT be NULL.
	Buffer []byte `idl:"name:Buffer;size_is:(Length)" json:"buffer"`
}

func (o *ForestTrustBinaryData) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint32(len(o.Buffer))
	}
	if o.Length > uint32(131072) {
		return fmt.Errorf("Length is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustBinaryData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.Buffer != nil || o.Length > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Buffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustBinaryData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Length > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Length)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// ForestTrustDomainInfo structure represents LSA_FOREST_TRUST_DOMAIN_INFO RPC structure.
//
// The LSA_FOREST_TRUST_DOMAIN_INFO structure is used to communicate a forest trust
// record corresponding to the LSA_FOREST_TRUST_DOMAIN_INFO value of ForestTrustDomainInfo.
type ForestTrustDomainInfo struct {
	// Sid:  Domain SID for the trusted domain.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
	// DnsName:  The DNS name of the trusted domain.
	DNSName *UnicodeString `idl:"name:DnsName" json:"dns_name"`
	// NetbiosName:  The NetBIOS name of the trusted domain, as specified in [RFC1088].
	NetBIOSName *UnicodeString `idl:"name:NetbiosName" json:"netbios_name"`
}

func (o *ForestTrustDomainInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustDomainInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SID, _ptr_Sid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DNSName != nil {
		if err := o.DNSName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.NetBIOSName != nil {
		if err := o.NetBIOSName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustDomainInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	if o.DNSName == nil {
		o.DNSName = &UnicodeString{}
	}
	if err := o.DNSName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.NetBIOSName == nil {
		o.NetBIOSName = &UnicodeString{}
	}
	if err := o.NetBIOSName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ForestTrustRecord structure represents LSA_FOREST_TRUST_RECORD RPC structure.
//
// The LSA_FOREST_TRUST_RECORD structure is used to communicate the type, creation time,
// and data for a forest trust record. The data is determined by the trust type as follows
// in the definition of the contained union.
type ForestTrustRecord struct {
	// Flags:   Contains zero or more flags from LSA Trust Record Flags (section 2.2.1.5).
	// See the Meaning column in the table of that section for related information.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ForestTrustType:  This value is one of LSA_FOREST_TRUST_RECORD_TYPE.
	ForestTrustType ForestTrustRecordType `idl:"name:ForestTrustType" json:"forest_trust_type"`
	// Time:  The date and time when this entry was created. It is a 64-bit value that represents
	// the number of 100-nanosecond intervals since January 1, 1601, UTC.
	Time *dtyp.LargeInteger `idl:"name:Time" json:"time"`
	// ForestTrustData:  An LSA_UNICODE_STRING, LSA_FOREST_TRUST_DOMAIN_INFO, or LSA_FOREST_TRUST_SCANNER_INFO
	// structure, depending on the value of ForestTrustType as specified in the structure
	// definition for LSA_FOREST_TRUST_RECORD.
	ForestTrustData *ForestTrustRecord_ForestTrustData `idl:"name:ForestTrustData;switch_is:ForestTrustType" json:"forest_trust_data"`
}

func (o *ForestTrustRecord) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustRecord) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ForestTrustType)); err != nil {
		return err
	}
	if o.Time != nil {
		if err := o.Time.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	_swForestTrustData := uint16(o.ForestTrustType)
	if o.ForestTrustData != nil {
		if err := o.ForestTrustData.MarshalUnionNDR(ctx, w, _swForestTrustData); err != nil {
			return err
		}
	} else {
		if err := (&ForestTrustRecord_ForestTrustData{}).MarshalUnionNDR(ctx, w, _swForestTrustData); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustRecord) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ForestTrustType)); err != nil {
		return err
	}
	if o.Time == nil {
		o.Time = &dtyp.LargeInteger{}
	}
	if err := o.Time.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ForestTrustData == nil {
		o.ForestTrustData = &ForestTrustRecord_ForestTrustData{}
	}
	_swForestTrustData := uint16(o.ForestTrustType)
	if err := o.ForestTrustData.UnmarshalUnionNDR(ctx, w, _swForestTrustData); err != nil {
		return err
	}
	return nil
}

// ForestTrustRecord_ForestTrustData structure represents LSA_FOREST_TRUST_RECORD union anonymous member.
//
// The LSA_FOREST_TRUST_RECORD structure is used to communicate the type, creation time,
// and data for a forest trust record. The data is determined by the trust type as follows
// in the definition of the contained union.
type ForestTrustRecord_ForestTrustData struct {
	// Types that are assignable to Value
	//
	// *ForestTrustData_TopLevelName
	// *ForestTrustData_DomainInfo
	// *ForestTrustData_Data
	Value is_ForestTrustRecord_ForestTrustData `json:"value"`
}

func (o *ForestTrustRecord_ForestTrustData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ForestTrustData_TopLevelName:
		if value != nil {
			return value.TopLevelName
		}
	case *ForestTrustData_DomainInfo:
		if value != nil {
			return value.DomainInfo
		}
	case *ForestTrustData_Data:
		if value != nil {
			return value.Data
		}
	}
	return nil
}

type is_ForestTrustRecord_ForestTrustData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ForestTrustRecord_ForestTrustData()
}

func (o *ForestTrustRecord_ForestTrustData) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ForestTrustData_TopLevelName:
		switch sw {
		case uint16(0),
			uint16(1):
			return sw
		}
		return uint16(0)
	case *ForestTrustData_DomainInfo:
		return uint16(2)
	}
	return uint16(0)
}

func (o *ForestTrustRecord_ForestTrustData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(0),
		uint16(1):
		_o, _ := o.Value.(*ForestTrustData_TopLevelName)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ForestTrustData_TopLevelName{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*ForestTrustData_DomainInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ForestTrustData_DomainInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		_o, _ := o.Value.(*ForestTrustData_Data)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ForestTrustData_Data{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *ForestTrustRecord_ForestTrustData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(0),
		uint16(1):
		o.Value = &ForestTrustData_TopLevelName{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &ForestTrustData_DomainInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &ForestTrustData_Data{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// ForestTrustData_TopLevelName structure represents ForestTrustRecord_ForestTrustData RPC union arm.
//
// It has following labels: 0, 1
type ForestTrustData_TopLevelName struct {
	TopLevelName *UnicodeString `idl:"name:TopLevelName" json:"top_level_name"`
}

func (*ForestTrustData_TopLevelName) is_ForestTrustRecord_ForestTrustData() {}

func (o *ForestTrustData_TopLevelName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TopLevelName != nil {
		if err := o.TopLevelName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustData_TopLevelName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TopLevelName == nil {
		o.TopLevelName = &UnicodeString{}
	}
	if err := o.TopLevelName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ForestTrustData_DomainInfo structure represents ForestTrustRecord_ForestTrustData RPC union arm.
//
// It has following labels: 2
type ForestTrustData_DomainInfo struct {
	DomainInfo *ForestTrustDomainInfo `idl:"name:DomainInfo" json:"domain_info"`
}

func (*ForestTrustData_DomainInfo) is_ForestTrustRecord_ForestTrustData() {}

func (o *ForestTrustData_DomainInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DomainInfo != nil {
		if err := o.DomainInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ForestTrustDomainInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustData_DomainInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DomainInfo == nil {
		o.DomainInfo = &ForestTrustDomainInfo{}
	}
	if err := o.DomainInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ForestTrustData_Data structure represents ForestTrustRecord_ForestTrustData RPC default union arm.
type ForestTrustData_Data struct {
	Data *ForestTrustBinaryData `idl:"name:Data" json:"data"`
}

func (*ForestTrustData_Data) is_ForestTrustRecord_ForestTrustData() {}

func (o *ForestTrustData_Data) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Data != nil {
		if err := o.Data.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ForestTrustBinaryData{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustData_Data) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Data == nil {
		o.Data = &ForestTrustBinaryData{}
	}
	if err := o.Data.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ForestTrustInformation structure represents LSA_FOREST_TRUST_INFORMATION RPC structure.
//
// The LSA_FOREST_TRUST_INFORMATION structure is a collection of LSA_FOREST_TRUST_RECORD
// (section 2.2.7.21) structures.
type ForestTrustInformation struct {
	// RecordCount:  A count of elements in the Entries array.<38>
	RecordCount uint32 `idl:"name:RecordCount" json:"record_count"`
	// Entries:  An array of LSA_FOREST_TRUST_RECORD structures. If the RecordCount field
	// has a value other than 0, this field MUST NOT be NULL.
	Entries []*ForestTrustRecord `idl:"name:Entries;size_is:(RecordCount)" json:"entries"`
}

func (o *ForestTrustInformation) xxx_PreparePayload(ctx context.Context) error {
	if o.Entries != nil && o.RecordCount == 0 {
		o.RecordCount = uint32(len(o.Entries))
	}
	if o.RecordCount > uint32(4000) {
		return fmt.Errorf("RecordCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RecordCount); err != nil {
		return err
	}
	if o.Entries != nil || o.RecordCount > 0 {
		_ptr_Entries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RecordCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Entries {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Entries[i1] != nil {
					_ptr_Entries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Entries[i1] != nil {
							if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ForestTrustRecord{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Entries[i1], _ptr_Entries); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Entries, _ptr_Entries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecordCount); err != nil {
		return err
	}
	_ptr_Entries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RecordCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RecordCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*ForestTrustRecord, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			_ptr_Entries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Entries[i1] == nil {
					o.Entries[i1] = &ForestTrustRecord{}
				}
				if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Entries := func(ptr interface{}) { o.Entries[i1] = *ptr.(**ForestTrustRecord) }
			if err := w.ReadPointer(&o.Entries[i1], _s_Entries, _ptr_Entries); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Entries := func(ptr interface{}) { o.Entries = *ptr.(*[]*ForestTrustRecord) }
	if err := w.ReadPointer(&o.Entries, _s_Entries, _ptr_Entries); err != nil {
		return err
	}
	return nil
}

// ForestTrustCollisionRecordType type represents LSA_FOREST_TRUST_COLLISION_RECORD_TYPE RPC enumeration.
//
// The LSA_FOREST_TRUST_COLLISION_RECORD_TYPE type specifies the type of a collision
// record in the message.
type ForestTrustCollisionRecordType uint16

var (
	// CollisionTdo: A forest trust record that a caller attempted to set on a trusted domain
	// object has suffered a collision with another trusted domain object in Active Directory,
	// as specified in [MS-ADTS], section 6.1.6.
	ForestTrustCollisionRecordTypeCollisionTDO ForestTrustCollisionRecordType = 0
	// CollisionXref: A forest trust record that a caller attempted to set on a trusted
	// domain object has suffered a collision with a cross-reference object belonging to
	// the forest to which the server belongs, as specified in [MS-ADTS], section 6.1.6.
	ForestTrustCollisionRecordTypeCollisionXref ForestTrustCollisionRecordType = 1
	// CollisionOther: A forest trust record that a caller attempted to set on a trusted
	// domain object has suffered a collision for an unknown reason.
	ForestTrustCollisionRecordTypeCollisionOther ForestTrustCollisionRecordType = 2
)

func (o ForestTrustCollisionRecordType) String() string {
	switch o {
	case ForestTrustCollisionRecordTypeCollisionTDO:
		return "ForestTrustCollisionRecordTypeCollisionTDO"
	case ForestTrustCollisionRecordTypeCollisionXref:
		return "ForestTrustCollisionRecordTypeCollisionXref"
	case ForestTrustCollisionRecordTypeCollisionOther:
		return "ForestTrustCollisionRecordTypeCollisionOther"
	}
	return "Invalid"
}

// ForestTrustCollisionRecord structure represents LSA_FOREST_TRUST_COLLISION_RECORD RPC structure.
//
// The LSA_FOREST_TRUST_COLLISION_RECORD structure is used to communicate forest trust
// collision information. For more information about trusted domain objects, see [MS-ADTS]
// section 6.1.6.
type ForestTrustCollisionRecord struct {
	// Index:  An ordinal number of a forest trust record in the forest trust information
	// supplied by the caller that suffered a collision. For rules about collisions, see
	// sections 3.1.4.7.16 and 3.1.4.7.16.1.
	Index uint32 `idl:"name:Index" json:"index"`
	// Type:  The type of collision record, as specified in section 2.2.7.26.
	Type ForestTrustCollisionRecordType `idl:"name:Type" json:"type"`
	// Flags:  A set of bits specifying the nature of the collision. These flags and the
	// rules for generating them are specified in sections 3.1.4.7.16 and 3.1.4.7.16.1.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// Name:  The name of the existing entity (a top-level name entry, a domain information
	// entry, or a top-level name exclusion entry) that caused the collision.
	Name *UnicodeString `idl:"name:Name" json:"name"`
}

func (o *ForestTrustCollisionRecord) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustCollisionRecord) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustCollisionRecord) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ForestTrustCollisionInformation structure represents LSA_FOREST_TRUST_COLLISION_INFORMATION RPC structure.
//
// The LSA_FOREST_TRUST_COLLISION_INFORMATION structure is used to communicate a set
// of LSA_FOREST_TRUST_COLLISION_RECORD structures.
type ForestTrustCollisionInformation struct {
	// RecordCount:  The count of elements in the Entries array.
	RecordCount uint32 `idl:"name:RecordCount" json:"record_count"`
	// Entries:  An array of LSA_FOREST_TRUST_COLLISION_RECORD (section 2.2.7.27) structures.
	// If the RecordCount field has a value other than zero, this field MUST NOT be NULL.
	Entries []*ForestTrustCollisionRecord `idl:"name:Entries;size_is:(RecordCount)" json:"entries"`
}

func (o *ForestTrustCollisionInformation) xxx_PreparePayload(ctx context.Context) error {
	if o.Entries != nil && o.RecordCount == 0 {
		o.RecordCount = uint32(len(o.Entries))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustCollisionInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RecordCount); err != nil {
		return err
	}
	if o.Entries != nil || o.RecordCount > 0 {
		_ptr_Entries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RecordCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Entries {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Entries[i1] != nil {
					_ptr_Entries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Entries[i1] != nil {
							if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ForestTrustCollisionRecord{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Entries[i1], _ptr_Entries); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Entries, _ptr_Entries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ForestTrustCollisionInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecordCount); err != nil {
		return err
	}
	_ptr_Entries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RecordCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RecordCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*ForestTrustCollisionRecord, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			_ptr_Entries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Entries[i1] == nil {
					o.Entries[i1] = &ForestTrustCollisionRecord{}
				}
				if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Entries := func(ptr interface{}) { o.Entries[i1] = *ptr.(**ForestTrustCollisionRecord) }
			if err := w.ReadPointer(&o.Entries[i1], _s_Entries, _ptr_Entries); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Entries := func(ptr interface{}) { o.Entries = *ptr.(*[]*ForestTrustCollisionRecord) }
	if err := w.ReadPointer(&o.Entries, _s_Entries, _ptr_Entries); err != nil {
		return err
	}
	return nil
}

// AccountInformation structure represents LSAPR_ACCOUNT_INFORMATION RPC structure.
//
// The LSAPR_ACCOUNT_INFORMATION structure specifies a security principal security identifier
// (SID).
type AccountInformation struct {
	// Sid:  This field contains the SID of the security principal. This field MUST NOT
	// be NULL.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
}

func (o *AccountInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccountInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SID, _ptr_Sid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccountInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// AccountEnumBuffer structure represents LSAPR_ACCOUNT_ENUM_BUFFER RPC structure.
//
// The LSAPR_ACCOUNT_ENUM_BUFFER structure specifies a collection of security principal
// SIDs represented in an array of structures of type LSAPR_ACCOUNT_INFORMATION.
type AccountEnumBuffer struct {
	// EntriesRead:  This field contains the number of security principals.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Information:  This field contains a set of structures that define the security principal
	// SID, as specified in section 2.2.5.1. If the EntriesRead field has a value other
	// than 0, this field MUST NOT be NULL.
	Information []*AccountInformation `idl:"name:Information;size_is:(EntriesRead)" json:"information"`
}

func (o *AccountEnumBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Information != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Information))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccountEnumBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Information != nil || o.EntriesRead > 0 {
		_ptr_Information := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Information {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Information[i1] != nil {
					if err := o.Information[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&AccountInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Information); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&AccountInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Information, _ptr_Information); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccountEnumBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_Information := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Information", sizeInfo[0])
		}
		o.Information = make([]*AccountInformation, sizeInfo[0])
		for i1 := range o.Information {
			i1 := i1
			if o.Information[i1] == nil {
				o.Information[i1] = &AccountInformation{}
			}
			if err := o.Information[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Information := func(ptr interface{}) { o.Information = *ptr.(*[]*AccountInformation) }
	if err := w.ReadPointer(&o.Information, _s_Information, _ptr_Information); err != nil {
		return err
	}
	return nil
}

// SrSecurityDescriptor structure represents LSAPR_SR_SECURITY_DESCRIPTOR RPC structure.
//
// The LSAPR_SR_SECURITY_DESCRIPTOR structure is used to communicate a self-relative
// security descriptor, as specified in [MS-DTYP] section 2.4.6.
type SrSecurityDescriptor struct {
	// Length:  The count of bytes in SecurityDescriptor.<18>
	Length uint32 `idl:"name:Length" json:"length"`
	// SecurityDescriptor:  The contiguous buffer containing the self-relative security
	// descriptor. This field MUST contain the Length number of bytes. If the Length field
	// has a value other than 0, this field MUST NOT be NULL.
	SecurityDescriptor []byte `idl:"name:SecurityDescriptor;size_is:(Length)" json:"security_descriptor"`
}

func (o *SrSecurityDescriptor) xxx_PreparePayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.Length == 0 {
		o.Length = uint32(len(o.SecurityDescriptor))
	}
	if o.Length > uint32(262144) {
		return fmt.Errorf("Length is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SrSecurityDescriptor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil || o.Length > 0 {
		_ptr_SecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SecurityDescriptor {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.SecurityDescriptor[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.SecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityDescriptor, _ptr_SecurityDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SrSecurityDescriptor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_SecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Length > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Length)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SecurityDescriptor", sizeInfo[0])
		}
		o.SecurityDescriptor = make([]byte, sizeInfo[0])
		for i1 := range o.SecurityDescriptor {
			i1 := i1
			if err := w.ReadData(&o.SecurityDescriptor[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_SecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SecurityDescriptor, _s_SecurityDescriptor, _ptr_SecurityDescriptor); err != nil {
		return err
	}
	return nil
}

// LUIDAndAttributes structure represents LSAPR_LUID_AND_ATTRIBUTES RPC structure.
//
// The LSAPR_LUID_AND_ATTRIBUTES structure is a tuple defining a locally unique identifier
// (LUID) and a field defining the attributes of the LUID.
type LUIDAndAttributes struct {
	// Luid:  The locally unique identifier.
	LUID *dtyp.LUID `idl:"name:Luid" json:"luid"`
	// Attributes:  This field contains bitmapped values that define the properties of the
	// privilege set. One or more of the following flags can be set.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | E | D |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// D: The privilege is enabled by default.
	//
	// E: The privilege is enabled.
	Attributes uint32 `idl:"name:Attributes" json:"attributes"`
}

func (o *LUIDAndAttributes) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LUIDAndAttributes) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.LUID != nil {
		if err := o.LUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	return nil
}
func (o *LUIDAndAttributes) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.LUID == nil {
		o.LUID = &dtyp.LUID{}
	}
	if err := o.LUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	return nil
}

// PrivilegeSet structure represents LSAPR_PRIVILEGE_SET RPC structure.
//
// The LSAPR_PRIVILEGE_SET structure defines a set of privileges that belong to an account.
type PrivilegeSet struct {
	// PrivilegeCount:  This field contains the number of privileges.<26>
	PrivilegeCount uint32 `idl:"name:PrivilegeCount" json:"privilege_count"`
	// Control:  This field contains bitmapped values that define the properties of the
	// privilege set.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | O |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// O: Valid for a set operation indicating that all specified privileges that are not
	// already assigned are to be assigned.
	Control uint32 `idl:"name:Control" json:"control"`
	// Privilege:  An array of LSAPR_LUID_AND_ATTRIBUTES structures. If the PrivilegeCount
	// field has a value different than 0, this field MUST NOT be NULL.
	Privilege []*LUIDAndAttributes `idl:"name:Privilege;size_is:(PrivilegeCount)" json:"privilege"`
}

func (o *PrivilegeSet) xxx_PreparePayload(ctx context.Context) error {
	if o.Privilege != nil && o.PrivilegeCount == 0 {
		o.PrivilegeCount = uint32(len(o.Privilege))
	}
	if o.PrivilegeCount > uint32(1000) {
		return fmt.Errorf("PrivilegeCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *PrivilegeSet) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.PrivilegeCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PrivilegeSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PrivilegeCount); err != nil {
		return err
	}
	if err := w.WriteData(o.Control); err != nil {
		return err
	}
	for i1 := range o.Privilege {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Privilege[i1] != nil {
			if err := o.Privilege[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&LUIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Privilege); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&LUIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PrivilegeSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrivilegeCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.Control); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.PrivilegeCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.PrivilegeCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Privilege", sizeInfo[0])
	}
	o.Privilege = make([]*LUIDAndAttributes, sizeInfo[0])
	for i1 := range o.Privilege {
		i1 := i1
		if o.Privilege[i1] == nil {
			o.Privilege[i1] = &LUIDAndAttributes{}
		}
		if err := o.Privilege[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// PolicyPrivilegeDefinition structure represents LSAPR_POLICY_PRIVILEGE_DEF RPC structure.
//
// The LSAPR_POLICY_PRIVILEGE_DEF structure specifies a privilege definition, which
// consists of a pairing of a human-readable name with a locally unique identifier (LUID).
type PolicyPrivilegeDefinition struct {
	// Name:  An RPC_UNICODE_STRING that contains the privilege name.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// LocalValue:  This field contains the LUID value assigned locally for efficient representation
	// of the privilege. This value is meaningful only on the system where it was assigned.
	LocalValue *dtyp.LUID `idl:"name:LocalValue" json:"local_value"`
}

func (o *PolicyPrivilegeDefinition) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyPrivilegeDefinition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalValue != nil {
		if err := o.LocalValue.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyPrivilegeDefinition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalValue == nil {
		o.LocalValue = &dtyp.LUID{}
	}
	if err := o.LocalValue.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PrivilegeEnumBuffer structure represents LSAPR_PRIVILEGE_ENUM_BUFFER RPC structure.
//
// The LSAPR_PRIVILEGE_ENUM_BUFFER structure specifies a collection of privilege definitions
// of type LSAPR_POLICY_PRIVILEGE_DEF.
type PrivilegeEnumBuffer struct {
	// Entries:  This field contains the number of privileges in the structure.
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// Privileges:  This field contains a set of structures that define the privileges,
	// as specified in section 2.2.8.1. If the Entries field has a value other than 0, this
	// field MUST NOT be NULL.
	Privileges []*PolicyPrivilegeDefinition `idl:"name:Privileges;size_is:(Entries)" json:"privileges"`
}

func (o *PrivilegeEnumBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Privileges != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.Privileges))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PrivilegeEnumBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.Privileges != nil || o.Entries > 0 {
		_ptr_Privileges := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Privileges {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Privileges[i1] != nil {
					if err := o.Privileges[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PolicyPrivilegeDefinition{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Privileges); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&PolicyPrivilegeDefinition{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Privileges, _ptr_Privileges); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PrivilegeEnumBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_Privileges := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Privileges", sizeInfo[0])
		}
		o.Privileges = make([]*PolicyPrivilegeDefinition, sizeInfo[0])
		for i1 := range o.Privileges {
			i1 := i1
			if o.Privileges[i1] == nil {
				o.Privileges[i1] = &PolicyPrivilegeDefinition{}
			}
			if err := o.Privileges[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Privileges := func(ptr interface{}) { o.Privileges = *ptr.(*[]*PolicyPrivilegeDefinition) }
	if err := w.ReadPointer(&o.Privileges, _s_Privileges, _ptr_Privileges); err != nil {
		return err
	}
	return nil
}

// CRCipherValue structure represents LSAPR_CR_CIPHER_VALUE RPC structure.
//
// The LSAPR_CR_CIPHER_VALUE structure is a counted buffer of bytes containing a secret
// object.
type CRCipherValue struct {
	// Length:  This field contains the number of valid bytes in the Buffer field.<27>
	Length uint32 `idl:"name:Length" json:"length"`
	// MaximumLength:  This field contains the number of allocated bytes in the Buffer field.<28>
	MaximumLength uint32 `idl:"name:MaximumLength" json:"maximum_length"`
	// Buffer:  This field contains the actual secret data. If the value of the MaximumLength
	// field is greater than 0, this field MUST contain a non-NULL value. This field is
	// always encrypted using algorithms as specified in section 5.1.2.
	Buffer []byte `idl:"name:Buffer;size_is:(MaximumLength);length_is:(Length)" json:"buffer"`
}

func (o *CRCipherValue) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.MaximumLength == 0 {
		o.MaximumLength = uint32(len(o.Buffer))
	}
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint32(len(o.Buffer))
	}
	if o.Length > uint32(131088) {
		return fmt.Errorf("Length is out of range")
	}
	if o.MaximumLength > uint32(131088) {
		return fmt.Errorf("MaximumLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CRCipherValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumLength); err != nil {
		return err
	}
	if o.Buffer != nil || o.MaximumLength > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MaximumLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64(o.Length)
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
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Buffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CRCipherValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumLength); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// TrustedEnumBuffer structure represents LSAPR_TRUSTED_ENUM_BUFFER RPC structure.
//
// The LSAPR_TRUSTED_ENUM_BUFFER structure specifies a collection of trust information
// structures of type LSAPR_TRUST_INFORMATION.
type TrustedEnumBuffer struct {
	// EntriesRead:  This field contains the number of trust information structures.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Information:  This field contains a set of structures that define the trust information,
	// as specified in section 2.2.7.1. If the EntriesRead field has a value other than
	// 0, this field MUST NOT be NULL.
	Information []*TrustInformation `idl:"name:Information;size_is:(EntriesRead)" json:"information"`
}

func (o *TrustedEnumBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Information != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Information))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedEnumBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Information != nil || o.EntriesRead > 0 {
		_ptr_Information := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Information {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Information[i1] != nil {
					if err := o.Information[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TrustInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Information); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&TrustInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Information, _ptr_Information); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedEnumBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_Information := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Information", sizeInfo[0])
		}
		o.Information = make([]*TrustInformation, sizeInfo[0])
		for i1 := range o.Information {
			i1 := i1
			if o.Information[i1] == nil {
				o.Information[i1] = &TrustInformation{}
			}
			if err := o.Information[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Information := func(ptr interface{}) { o.Information = *ptr.(*[]*TrustInformation) }
	if err := w.ReadPointer(&o.Information, _s_Information, _ptr_Information); err != nil {
		return err
	}
	return nil
}

// PolicyAccountDomInfo structure represents LSAPR_POLICY_ACCOUNT_DOM_INFO RPC structure.
//
// The LSAPR_POLICY_ACCOUNT_DOM_INFO structure contains information about the server's
// account domain. The following structure corresponds to the PolicyAccountDomainInformation
// and PolicyLocalAccountDomainInformation information classes.
type PolicyAccountDomInfo struct {
	// DomainName:  This field contains a name for the account domain that is subjected
	// to the restrictions of a NetBIOS name, as specified in [RFC1088]. This value SHOULD
	// be used (by implementations external to this protocol) to identify the domain  via
	// the NetBIOS API, as specified in [RFC1088].
	DomainName *dtyp.UnicodeString `idl:"name:DomainName" json:"domain_name"`
	// DomainSid:  The SID of the account domain. This field MUST NOT be NULL.
	DomainSID *dtyp.SID `idl:"name:DomainSid" json:"domain_sid"`
}

func (o *PolicyAccountDomInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyAccountDomInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.DomainName != nil {
		if err := o.DomainName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DomainSID != nil {
		_ptr_DomainSid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DomainSID != nil {
				if err := o.DomainSID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainSID, _ptr_DomainSid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyAccountDomInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.DomainName == nil {
		o.DomainName = &dtyp.UnicodeString{}
	}
	if err := o.DomainName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_DomainSid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DomainSID == nil {
			o.DomainSID = &dtyp.SID{}
		}
		if err := o.DomainSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DomainSid := func(ptr interface{}) { o.DomainSID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.DomainSID, _s_DomainSid, _ptr_DomainSid); err != nil {
		return err
	}
	return nil
}

// PolicyPrimaryDomInfo structure represents LSAPR_POLICY_PRIMARY_DOM_INFO RPC structure.
//
// The LSAPR_POLICY_PRIMARY_DOM_INFO structure defines the server's primary domain.
//
// The following structure corresponds to the PolicyPrimaryDomainInformation information
// class.
type PolicyPrimaryDomInfo struct {
	// Name:  This field contains a name for the primary domain that is subject to the restrictions
	// of a NetBIOS name, as specified in [RFC1088]. The value SHOULD be used (by implementations
	// external to this protocol) to identify the domain via the NetBIOS API, as specified
	// in [RFC1088].
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// Sid:  The SID of the primary domain.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
}

func (o *PolicyPrimaryDomInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyPrimaryDomInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SID, _ptr_Sid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyPrimaryDomInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// PolicyDNSDomainInfo structure represents LSAPR_POLICY_DNS_DOMAIN_INFO RPC structure.
//
// The LSAPR_POLICY_DNS_DOMAIN_INFO structure is used to allow callers to query and
// set the server's primary domain.<22>
//
// The following structure corresponds to the PolicyDnsDomainInformation and PolicyDnsDomainInformationInt
// information classes.
type PolicyDNSDomainInfo struct {
	// Name:  This field contains a name for the domain that is subject to the restrictions
	// of a NetBIOS name, as specified in [RFC1088]. This value SHOULD be used (by implementations
	// external to this protocol) to identify the domain via the NetBIOS API, as specified
	// in [RFC1088].
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// DnsDomainName:  The fully qualified DNS name of the domain.
	DNSDomainName *dtyp.UnicodeString `idl:"name:DnsDomainName" json:"dns_domain_name"`
	// DnsForestName:  The fully qualified DNS name of the forest containing this domain.
	DNSForestName *dtyp.UnicodeString `idl:"name:DnsForestName" json:"dns_forest_name"`
	// DomainGuid:  The globally unique identifier (GUID), as specified in [MS-DTYP] section
	// 2.3.4.1, of the domain.
	DomainGUID *dtyp.GUID `idl:"name:DomainGuid" json:"domain_guid"`
	// Sid:  The SID of the domain.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
}

func (o *PolicyDNSDomainInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyDNSDomainInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DNSDomainName != nil {
		if err := o.DNSDomainName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DNSForestName != nil {
		if err := o.DNSForestName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DomainGUID != nil {
		if err := o.DomainGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SID, _ptr_Sid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyDNSDomainInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DNSDomainName == nil {
		o.DNSDomainName = &dtyp.UnicodeString{}
	}
	if err := o.DNSDomainName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DNSForestName == nil {
		o.DNSForestName = &dtyp.UnicodeString{}
	}
	if err := o.DNSForestName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DomainGUID == nil {
		o.DomainGUID = &dtyp.GUID{}
	}
	if err := o.DomainGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// PolicyPDAccountInfo structure represents LSAPR_POLICY_PD_ACCOUNT_INFO RPC structure.
//
// The LSAPR_POLICY_PD_ACCOUNT_INFO structure is obsolete and exists for backward compatibility
// purposes only.
type PolicyPDAccountInfo struct {
	// Name: Represents the name of an account in the domain that is to be used for authentication
	// and name/ID lookup requests.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
}

func (o *PolicyPDAccountInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyPDAccountInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyPDAccountInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyReplicaSourceInfo structure represents LSAPR_POLICY_REPLICA_SRCE_INFO RPC structure.
//
// The LSAPR_POLICY_REPLICA_SRCE_INFO structure corresponds to the PolicyReplicaSourceInformation
// information class.
type PolicyReplicaSourceInfo struct {
	// ReplicaSource:  A string.
	SourceReplica *dtyp.UnicodeString `idl:"name:ReplicaSource" json:"source_replica"`
	// ReplicaAccountName:  A string.
	AccountNameReplica *dtyp.UnicodeString `idl:"name:ReplicaAccountName" json:"account_name_replica"`
}

func (o *PolicyReplicaSourceInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyReplicaSourceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.SourceReplica != nil {
		if err := o.SourceReplica.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AccountNameReplica != nil {
		if err := o.AccountNameReplica.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyReplicaSourceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.SourceReplica == nil {
		o.SourceReplica = &dtyp.UnicodeString{}
	}
	if err := o.SourceReplica.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AccountNameReplica == nil {
		o.AccountNameReplica = &dtyp.UnicodeString{}
	}
	if err := o.AccountNameReplica.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyAuditEventsInfo structure represents LSAPR_POLICY_AUDIT_EVENTS_INFO RPC structure.
//
// The LSAPR_POLICY_AUDIT_EVENTS_INFO structure contains auditing options on the server.
type PolicyAuditEventsInfo struct {
	// AuditingMode:  0 indicates that auditing is disabled. All other values indicate that
	// auditing is enabled.
	AuditingMode uint8 `idl:"name:AuditingMode" json:"auditing_mode"`
	// EventAuditingOptions:  An array of values specifying the auditing options for a particular
	// audit type. The auditing type of an element is represented by its index in the array,
	// which is identified by the POLICY_AUDIT_EVENT_TYPE enumeration (see section 2.2.4.20).
	// Each element MUST contain one or more of the values in the table below.
	//
	// If the MaximumAuditEventCount field has a value other than 0, this field MUST NOT
	// be NULL.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| POLICY_AUDIT_EVENT_UNCHANGED 0x00000000 | Leave existing auditing options unchanged for events of this type; used only for |
	//	|                                         | set operations. This value cannot be combined with values in this table.         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| POLICY_AUDIT_EVENT_NONE 0x00000004      | Upon updates, this value causes the existing auditing options for events of      |
	//	|                                         | this type to be deleted and replaced with any other new values specified. If     |
	//	|                                         | specified by itself, this value cancels all auditing options for events of this  |
	//	|                                         | type. This value is used only for set operations.                                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| POLICY_AUDIT_EVENT_SUCCESS 0x00000001   | When auditing is enabled, audit all successful occurrences of events of the      |
	//	|                                         | given type.                                                                      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| POLICY_AUDIT_EVENT_FAILURE 0x00000002   | When auditing is enabled, audit all unsuccessful occurrences of events of the    |
	//	|                                         | given type.                                                                      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	EventAuditingOptions []uint32 `idl:"name:EventAuditingOptions;size_is:(MaximumAuditEventCount)" json:"event_auditing_options"`
	// MaximumAuditEventCount:  The number of entries in the EventAuditingOptions array.<21>
	MaximumAuditEventCount uint32 `idl:"name:MaximumAuditEventCount" json:"maximum_audit_event_count"`
}

func (o *PolicyAuditEventsInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.EventAuditingOptions != nil && o.MaximumAuditEventCount == 0 {
		o.MaximumAuditEventCount = uint32(len(o.EventAuditingOptions))
	}
	if o.MaximumAuditEventCount > uint32(1000) {
		return fmt.Errorf("MaximumAuditEventCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyAuditEventsInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.AuditingMode); err != nil {
		return err
	}
	if o.EventAuditingOptions != nil || o.MaximumAuditEventCount > 0 {
		_ptr_EventAuditingOptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MaximumAuditEventCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.EventAuditingOptions {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.EventAuditingOptions[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.EventAuditingOptions); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EventAuditingOptions, _ptr_EventAuditingOptions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MaximumAuditEventCount); err != nil {
		return err
	}
	return nil
}
func (o *PolicyAuditEventsInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuditingMode); err != nil {
		return err
	}
	_ptr_EventAuditingOptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.MaximumAuditEventCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.MaximumAuditEventCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.EventAuditingOptions", sizeInfo[0])
		}
		o.EventAuditingOptions = make([]uint32, sizeInfo[0])
		for i1 := range o.EventAuditingOptions {
			i1 := i1
			if err := w.ReadData(&o.EventAuditingOptions[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_EventAuditingOptions := func(ptr interface{}) { o.EventAuditingOptions = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.EventAuditingOptions, _s_EventAuditingOptions, _ptr_EventAuditingOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumAuditEventCount); err != nil {
		return err
	}
	return nil
}

// PolicyMachineAccountInfo structure represents LSAPR_POLICY_MACHINE_ACCT_INFO RPC structure.
type PolicyMachineAccountInfo struct {
	RID uint32    `idl:"name:Rid" json:"rid"`
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
}

func (o *PolicyMachineAccountInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyMachineAccountInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RID); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SID, _ptr_Sid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyMachineAccountInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RID); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// PolicyInformation structure represents LSAPR_POLICY_INFORMATION RPC union.
type PolicyInformation struct {
	// Types that are assignable to Value
	//
	// *PolicyInformation_PolicyAuditLogInfo
	// *PolicyInformation_PolicyAuditEventsInfo
	// *PolicyInformation_PolicyPrimaryDomainInfo
	// *PolicyInformation_PolicyAccountDomainInfo
	// *PolicyInformation_PolicyPDAccountInfo
	// *PolicyInformation_PolicyServerRoleInfo
	// *PolicyInformation_PolicyReplicaSourceInfo
	// *PolicyInformation_PolicyModificationInfo
	// *PolicyInformation_PolicyAuditFullSetInfo
	// *PolicyInformation_PolicyAuditFullQueryInfo
	// *PolicyInformation_PolicyDNSDomainInfo
	// *PolicyInformation_PolicyDNSDomainInfoInt
	// *PolicyInformation_PolicyLocalAccountDomainInfo
	// *PolicyInformation_PolicyMachineAccountInfo
	Value is_PolicyInformation `json:"value"`
}

func (o *PolicyInformation) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *PolicyInformation_PolicyAuditLogInfo:
		if value != nil {
			return value.PolicyAuditLogInfo
		}
	case *PolicyInformation_PolicyAuditEventsInfo:
		if value != nil {
			return value.PolicyAuditEventsInfo
		}
	case *PolicyInformation_PolicyPrimaryDomainInfo:
		if value != nil {
			return value.PolicyPrimaryDomainInfo
		}
	case *PolicyInformation_PolicyAccountDomainInfo:
		if value != nil {
			return value.PolicyAccountDomainInfo
		}
	case *PolicyInformation_PolicyPDAccountInfo:
		if value != nil {
			return value.PolicyPDAccountInfo
		}
	case *PolicyInformation_PolicyServerRoleInfo:
		if value != nil {
			return value.PolicyServerRoleInfo
		}
	case *PolicyInformation_PolicyReplicaSourceInfo:
		if value != nil {
			return value.PolicyReplicaSourceInfo
		}
	case *PolicyInformation_PolicyModificationInfo:
		if value != nil {
			return value.PolicyModificationInfo
		}
	case *PolicyInformation_PolicyAuditFullSetInfo:
		if value != nil {
			return value.PolicyAuditFullSetInfo
		}
	case *PolicyInformation_PolicyAuditFullQueryInfo:
		if value != nil {
			return value.PolicyAuditFullQueryInfo
		}
	case *PolicyInformation_PolicyDNSDomainInfo:
		if value != nil {
			return value.PolicyDNSDomainInfo
		}
	case *PolicyInformation_PolicyDNSDomainInfoInt:
		if value != nil {
			return value.PolicyDNSDomainInfoInt
		}
	case *PolicyInformation_PolicyLocalAccountDomainInfo:
		if value != nil {
			return value.PolicyLocalAccountDomainInfo
		}
	case *PolicyInformation_PolicyMachineAccountInfo:
		if value != nil {
			return value.PolicyMachineAccountInfo
		}
	}
	return nil
}

type is_PolicyInformation interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_PolicyInformation()
}

func (o *PolicyInformation) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *PolicyInformation_PolicyAuditLogInfo:
		return uint16(1)
	case *PolicyInformation_PolicyAuditEventsInfo:
		return uint16(2)
	case *PolicyInformation_PolicyPrimaryDomainInfo:
		return uint16(3)
	case *PolicyInformation_PolicyAccountDomainInfo:
		return uint16(5)
	case *PolicyInformation_PolicyPDAccountInfo:
		return uint16(4)
	case *PolicyInformation_PolicyServerRoleInfo:
		return uint16(6)
	case *PolicyInformation_PolicyReplicaSourceInfo:
		return uint16(7)
	case *PolicyInformation_PolicyModificationInfo:
		return uint16(9)
	case *PolicyInformation_PolicyAuditFullSetInfo:
		return uint16(10)
	case *PolicyInformation_PolicyAuditFullQueryInfo:
		return uint16(11)
	case *PolicyInformation_PolicyDNSDomainInfo:
		return uint16(12)
	case *PolicyInformation_PolicyDNSDomainInfoInt:
		return uint16(13)
	case *PolicyInformation_PolicyLocalAccountDomainInfo:
		return uint16(14)
	case *PolicyInformation_PolicyMachineAccountInfo:
		return uint16(15)
	}
	return uint16(0)
}

func (o *PolicyInformation) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*PolicyInformation_PolicyAuditLogInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyAuditLogInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*PolicyInformation_PolicyAuditEventsInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyAuditEventsInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*PolicyInformation_PolicyPrimaryDomainInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyPrimaryDomainInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5):
		_o, _ := o.Value.(*PolicyInformation_PolicyAccountDomainInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyAccountDomainInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*PolicyInformation_PolicyPDAccountInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyPDAccountInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(6):
		_o, _ := o.Value.(*PolicyInformation_PolicyServerRoleInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyServerRoleInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(7):
		_o, _ := o.Value.(*PolicyInformation_PolicyReplicaSourceInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyReplicaSourceInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(9):
		_o, _ := o.Value.(*PolicyInformation_PolicyModificationInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyModificationInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(10):
		_o, _ := o.Value.(*PolicyInformation_PolicyAuditFullSetInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyAuditFullSetInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(11):
		_o, _ := o.Value.(*PolicyInformation_PolicyAuditFullQueryInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyAuditFullQueryInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(12):
		_o, _ := o.Value.(*PolicyInformation_PolicyDNSDomainInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyDNSDomainInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(13):
		_o, _ := o.Value.(*PolicyInformation_PolicyDNSDomainInfoInt)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyDNSDomainInfoInt{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(14):
		_o, _ := o.Value.(*PolicyInformation_PolicyLocalAccountDomainInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyLocalAccountDomainInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(15):
		_o, _ := o.Value.(*PolicyInformation_PolicyMachineAccountInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation_PolicyMachineAccountInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *PolicyInformation) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &PolicyInformation_PolicyAuditLogInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &PolicyInformation_PolicyAuditEventsInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &PolicyInformation_PolicyPrimaryDomainInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5):
		o.Value = &PolicyInformation_PolicyAccountDomainInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &PolicyInformation_PolicyPDAccountInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(6):
		o.Value = &PolicyInformation_PolicyServerRoleInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(7):
		o.Value = &PolicyInformation_PolicyReplicaSourceInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(9):
		o.Value = &PolicyInformation_PolicyModificationInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(10):
		o.Value = &PolicyInformation_PolicyAuditFullSetInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(11):
		o.Value = &PolicyInformation_PolicyAuditFullQueryInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(12):
		o.Value = &PolicyInformation_PolicyDNSDomainInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(13):
		o.Value = &PolicyInformation_PolicyDNSDomainInfoInt{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(14):
		o.Value = &PolicyInformation_PolicyLocalAccountDomainInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(15):
		o.Value = &PolicyInformation_PolicyMachineAccountInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// PolicyInformation_PolicyAuditLogInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 1
type PolicyInformation_PolicyAuditLogInfo struct {
	PolicyAuditLogInfo *PolicyAuditLogInfo `idl:"name:PolicyAuditLogInfo" json:"policy_audit_log_info"`
}

func (*PolicyInformation_PolicyAuditLogInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyAuditLogInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyAuditLogInfo != nil {
		if err := o.PolicyAuditLogInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyAuditLogInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyAuditLogInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyAuditLogInfo == nil {
		o.PolicyAuditLogInfo = &PolicyAuditLogInfo{}
	}
	if err := o.PolicyAuditLogInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyAuditEventsInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 2
type PolicyInformation_PolicyAuditEventsInfo struct {
	PolicyAuditEventsInfo *PolicyAuditEventsInfo `idl:"name:PolicyAuditEventsInfo" json:"policy_audit_events_info"`
}

func (*PolicyInformation_PolicyAuditEventsInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyAuditEventsInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyAuditEventsInfo != nil {
		if err := o.PolicyAuditEventsInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyAuditEventsInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyAuditEventsInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyAuditEventsInfo == nil {
		o.PolicyAuditEventsInfo = &PolicyAuditEventsInfo{}
	}
	if err := o.PolicyAuditEventsInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyPrimaryDomainInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 3
type PolicyInformation_PolicyPrimaryDomainInfo struct {
	PolicyPrimaryDomainInfo *PolicyPrimaryDomInfo `idl:"name:PolicyPrimaryDomainInfo" json:"policy_primary_domain_info"`
}

func (*PolicyInformation_PolicyPrimaryDomainInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyPrimaryDomainInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyPrimaryDomainInfo != nil {
		if err := o.PolicyPrimaryDomainInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyPrimaryDomInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyPrimaryDomainInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyPrimaryDomainInfo == nil {
		o.PolicyPrimaryDomainInfo = &PolicyPrimaryDomInfo{}
	}
	if err := o.PolicyPrimaryDomainInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyAccountDomainInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 5
type PolicyInformation_PolicyAccountDomainInfo struct {
	PolicyAccountDomainInfo *PolicyAccountDomInfo `idl:"name:PolicyAccountDomainInfo" json:"policy_account_domain_info"`
}

func (*PolicyInformation_PolicyAccountDomainInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyAccountDomainInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyAccountDomainInfo != nil {
		if err := o.PolicyAccountDomainInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyAccountDomInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyAccountDomainInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyAccountDomainInfo == nil {
		o.PolicyAccountDomainInfo = &PolicyAccountDomInfo{}
	}
	if err := o.PolicyAccountDomainInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyPDAccountInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 4
type PolicyInformation_PolicyPDAccountInfo struct {
	PolicyPDAccountInfo *PolicyPDAccountInfo `idl:"name:PolicyPdAccountInfo" json:"policy_pd_account_info"`
}

func (*PolicyInformation_PolicyPDAccountInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyPDAccountInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyPDAccountInfo != nil {
		if err := o.PolicyPDAccountInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyPDAccountInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyPDAccountInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyPDAccountInfo == nil {
		o.PolicyPDAccountInfo = &PolicyPDAccountInfo{}
	}
	if err := o.PolicyPDAccountInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyServerRoleInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 6
type PolicyInformation_PolicyServerRoleInfo struct {
	PolicyServerRoleInfo *PolicyLSAServerRoleInfo `idl:"name:PolicyServerRoleInfo" json:"policy_server_role_info"`
}

func (*PolicyInformation_PolicyServerRoleInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyServerRoleInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyServerRoleInfo != nil {
		if err := o.PolicyServerRoleInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyLSAServerRoleInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyServerRoleInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyServerRoleInfo == nil {
		o.PolicyServerRoleInfo = &PolicyLSAServerRoleInfo{}
	}
	if err := o.PolicyServerRoleInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyReplicaSourceInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 7
type PolicyInformation_PolicyReplicaSourceInfo struct {
	PolicyReplicaSourceInfo *PolicyReplicaSourceInfo `idl:"name:PolicyReplicaSourceInfo" json:"policy_replica_source_info"`
}

func (*PolicyInformation_PolicyReplicaSourceInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyReplicaSourceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyReplicaSourceInfo != nil {
		if err := o.PolicyReplicaSourceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyReplicaSourceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyReplicaSourceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyReplicaSourceInfo == nil {
		o.PolicyReplicaSourceInfo = &PolicyReplicaSourceInfo{}
	}
	if err := o.PolicyReplicaSourceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyModificationInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 9
type PolicyInformation_PolicyModificationInfo struct {
	PolicyModificationInfo *PolicyModificationInfo `idl:"name:PolicyModificationInfo" json:"policy_modification_info"`
}

func (*PolicyInformation_PolicyModificationInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyModificationInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyModificationInfo != nil {
		if err := o.PolicyModificationInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyModificationInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyModificationInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyModificationInfo == nil {
		o.PolicyModificationInfo = &PolicyModificationInfo{}
	}
	if err := o.PolicyModificationInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyAuditFullSetInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 10
type PolicyInformation_PolicyAuditFullSetInfo struct {
	PolicyAuditFullSetInfo *PolicyAuditFullSetInfo `idl:"name:PolicyAuditFullSetInfo" json:"policy_audit_full_set_info"`
}

func (*PolicyInformation_PolicyAuditFullSetInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyAuditFullSetInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyAuditFullSetInfo != nil {
		if err := o.PolicyAuditFullSetInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyAuditFullSetInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyAuditFullSetInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyAuditFullSetInfo == nil {
		o.PolicyAuditFullSetInfo = &PolicyAuditFullSetInfo{}
	}
	if err := o.PolicyAuditFullSetInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyAuditFullQueryInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 11
type PolicyInformation_PolicyAuditFullQueryInfo struct {
	PolicyAuditFullQueryInfo *PolicyAuditFullQueryInfo `idl:"name:PolicyAuditFullQueryInfo" json:"policy_audit_full_query_info"`
}

func (*PolicyInformation_PolicyAuditFullQueryInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyAuditFullQueryInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyAuditFullQueryInfo != nil {
		if err := o.PolicyAuditFullQueryInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyAuditFullQueryInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyAuditFullQueryInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyAuditFullQueryInfo == nil {
		o.PolicyAuditFullQueryInfo = &PolicyAuditFullQueryInfo{}
	}
	if err := o.PolicyAuditFullQueryInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyDNSDomainInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 12
type PolicyInformation_PolicyDNSDomainInfo struct {
	PolicyDNSDomainInfo *PolicyDNSDomainInfo `idl:"name:PolicyDnsDomainInfo" json:"policy_dns_domain_info"`
}

func (*PolicyInformation_PolicyDNSDomainInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyDNSDomainInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyDNSDomainInfo != nil {
		if err := o.PolicyDNSDomainInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyDNSDomainInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyDNSDomainInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyDNSDomainInfo == nil {
		o.PolicyDNSDomainInfo = &PolicyDNSDomainInfo{}
	}
	if err := o.PolicyDNSDomainInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyDNSDomainInfoInt structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 13
type PolicyInformation_PolicyDNSDomainInfoInt struct {
	PolicyDNSDomainInfoInt *PolicyDNSDomainInfo `idl:"name:PolicyDnsDomainInfoInt" json:"policy_dns_domain_info_int"`
}

func (*PolicyInformation_PolicyDNSDomainInfoInt) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyDNSDomainInfoInt) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyDNSDomainInfoInt != nil {
		if err := o.PolicyDNSDomainInfoInt.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyDNSDomainInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyDNSDomainInfoInt) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyDNSDomainInfoInt == nil {
		o.PolicyDNSDomainInfoInt = &PolicyDNSDomainInfo{}
	}
	if err := o.PolicyDNSDomainInfoInt.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyLocalAccountDomainInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 14
type PolicyInformation_PolicyLocalAccountDomainInfo struct {
	PolicyLocalAccountDomainInfo *PolicyAccountDomInfo `idl:"name:PolicyLocalAccountDomainInfo" json:"policy_local_account_domain_info"`
}

func (*PolicyInformation_PolicyLocalAccountDomainInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyLocalAccountDomainInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyLocalAccountDomainInfo != nil {
		if err := o.PolicyLocalAccountDomainInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyAccountDomInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyLocalAccountDomainInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyLocalAccountDomainInfo == nil {
		o.PolicyLocalAccountDomainInfo = &PolicyAccountDomInfo{}
	}
	if err := o.PolicyLocalAccountDomainInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyInformation_PolicyMachineAccountInfo structure represents LSAPR_POLICY_INFORMATION RPC union arm.
//
// It has following labels: 15
type PolicyInformation_PolicyMachineAccountInfo struct {
	PolicyMachineAccountInfo *PolicyMachineAccountInfo `idl:"name:PolicyMachineAccountInfo" json:"policy_machine_account_info"`
}

func (*PolicyInformation_PolicyMachineAccountInfo) is_PolicyInformation() {}

func (o *PolicyInformation_PolicyMachineAccountInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyMachineAccountInfo != nil {
		if err := o.PolicyMachineAccountInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyMachineAccountInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyInformation_PolicyMachineAccountInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyMachineAccountInfo == nil {
		o.PolicyMachineAccountInfo = &PolicyMachineAccountInfo{}
	}
	if err := o.PolicyMachineAccountInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyDomainQualityOfServiceInfo structure represents POLICY_DOMAIN_QUALITY_OF_SERVICE_INFO RPC structure.
//
// The POLICY_DOMAIN_QUALITY_OF_SERVICE_INFO structure is obsolete and exists for backward
// compatibility purposes only.
type PolicyDomainQualityOfServiceInfo struct {
	// QualityOfService:  Quality of service of the responder. MUST be set to zero when
	// sent and MUST be ignored on receipt.
	QualityOfService uint32 `idl:"name:QualityOfService" json:"quality_of_service"`
}

func (o *PolicyDomainQualityOfServiceInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyDomainQualityOfServiceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.QualityOfService); err != nil {
		return err
	}
	return nil
}
func (o *PolicyDomainQualityOfServiceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.QualityOfService); err != nil {
		return err
	}
	return nil
}

// PolicyDomainEFSInfo structure represents LSAPR_POLICY_DOMAIN_EFS_INFO RPC structure.
//
// The LSAPR_POLICY_DOMAIN_EFS_INFO structure communicates a counted binary byte array.
type PolicyDomainEFSInfo struct {
	// InfoLength:  The count of bytes in the EfsBlob.
	InfoLength uint32 `idl:"name:InfoLength" json:"info_length"`
	// EfsBlob:  An array of bytes, of size InfoLength bytes. If the value of InfoLength
	// is other than 0, this field MUST NOT be NULL. The syntax of this blob SHOULD<24>
	// conform to the layout specified in [MS-GPEF] section 2.2.1.2.1.
	EFSBlob []byte `idl:"name:EfsBlob;size_is:(InfoLength)" json:"efs_blob"`
}

func (o *PolicyDomainEFSInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.EFSBlob != nil && o.InfoLength == 0 {
		o.InfoLength = uint32(len(o.EFSBlob))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyDomainEFSInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.InfoLength); err != nil {
		return err
	}
	if o.EFSBlob != nil || o.InfoLength > 0 {
		_ptr_EfsBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InfoLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.EFSBlob {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.EFSBlob[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.EFSBlob); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EFSBlob, _ptr_EfsBlob); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyDomainEFSInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.InfoLength); err != nil {
		return err
	}
	_ptr_EfsBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InfoLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InfoLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.EFSBlob", sizeInfo[0])
		}
		o.EFSBlob = make([]byte, sizeInfo[0])
		for i1 := range o.EFSBlob {
			i1 := i1
			if err := w.ReadData(&o.EFSBlob[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_EfsBlob := func(ptr interface{}) { o.EFSBlob = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.EFSBlob, _s_EfsBlob, _ptr_EfsBlob); err != nil {
		return err
	}
	return nil
}

// PolicyDomainInformation structure represents LSAPR_POLICY_DOMAIN_INFORMATION RPC union.
type PolicyDomainInformation struct {
	// Types that are assignable to Value
	//
	// *PolicyDomainInformation_PolicyDomainQualityOfServiceInfo
	// *PolicyDomainInformation_PolicyDomainEFSInfo
	// *PolicyDomainInformation_PolicyDomainKerberosTicketInfo
	Value is_PolicyDomainInformation `json:"value"`
}

func (o *PolicyDomainInformation) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *PolicyDomainInformation_PolicyDomainQualityOfServiceInfo:
		if value != nil {
			return value.PolicyDomainQualityOfServiceInfo
		}
	case *PolicyDomainInformation_PolicyDomainEFSInfo:
		if value != nil {
			return value.PolicyDomainEFSInfo
		}
	case *PolicyDomainInformation_PolicyDomainKerberosTicketInfo:
		if value != nil {
			return value.PolicyDomainKerberosTicketInfo
		}
	}
	return nil
}

type is_PolicyDomainInformation interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_PolicyDomainInformation()
}

func (o *PolicyDomainInformation) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *PolicyDomainInformation_PolicyDomainQualityOfServiceInfo:
		return uint16(1)
	case *PolicyDomainInformation_PolicyDomainEFSInfo:
		return uint16(2)
	case *PolicyDomainInformation_PolicyDomainKerberosTicketInfo:
		return uint16(3)
	}
	return uint16(0)
}

func (o *PolicyDomainInformation) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*PolicyDomainInformation_PolicyDomainQualityOfServiceInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyDomainInformation_PolicyDomainQualityOfServiceInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*PolicyDomainInformation_PolicyDomainEFSInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyDomainInformation_PolicyDomainEFSInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*PolicyDomainInformation_PolicyDomainKerberosTicketInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PolicyDomainInformation_PolicyDomainKerberosTicketInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *PolicyDomainInformation) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &PolicyDomainInformation_PolicyDomainQualityOfServiceInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &PolicyDomainInformation_PolicyDomainEFSInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &PolicyDomainInformation_PolicyDomainKerberosTicketInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// PolicyDomainInformation_PolicyDomainQualityOfServiceInfo structure represents LSAPR_POLICY_DOMAIN_INFORMATION RPC union arm.
//
// It has following labels: 1
type PolicyDomainInformation_PolicyDomainQualityOfServiceInfo struct {
	PolicyDomainQualityOfServiceInfo *PolicyDomainQualityOfServiceInfo `idl:"name:PolicyDomainQualityOfServiceInfo" json:"policy_domain_quality_of_service_info"`
}

func (*PolicyDomainInformation_PolicyDomainQualityOfServiceInfo) is_PolicyDomainInformation() {}

func (o *PolicyDomainInformation_PolicyDomainQualityOfServiceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyDomainQualityOfServiceInfo != nil {
		if err := o.PolicyDomainQualityOfServiceInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyDomainQualityOfServiceInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyDomainInformation_PolicyDomainQualityOfServiceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyDomainQualityOfServiceInfo == nil {
		o.PolicyDomainQualityOfServiceInfo = &PolicyDomainQualityOfServiceInfo{}
	}
	if err := o.PolicyDomainQualityOfServiceInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyDomainInformation_PolicyDomainEFSInfo structure represents LSAPR_POLICY_DOMAIN_INFORMATION RPC union arm.
//
// It has following labels: 2
type PolicyDomainInformation_PolicyDomainEFSInfo struct {
	PolicyDomainEFSInfo *PolicyDomainEFSInfo `idl:"name:PolicyDomainEfsInfo" json:"policy_domain_efs_info"`
}

func (*PolicyDomainInformation_PolicyDomainEFSInfo) is_PolicyDomainInformation() {}

func (o *PolicyDomainInformation_PolicyDomainEFSInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyDomainEFSInfo != nil {
		if err := o.PolicyDomainEFSInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyDomainEFSInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyDomainInformation_PolicyDomainEFSInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyDomainEFSInfo == nil {
		o.PolicyDomainEFSInfo = &PolicyDomainEFSInfo{}
	}
	if err := o.PolicyDomainEFSInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PolicyDomainInformation_PolicyDomainKerberosTicketInfo structure represents LSAPR_POLICY_DOMAIN_INFORMATION RPC union arm.
//
// It has following labels: 3
type PolicyDomainInformation_PolicyDomainKerberosTicketInfo struct {
	PolicyDomainKerberosTicketInfo *PolicyDomainKerberosTicketInfo `idl:"name:PolicyDomainKerbTicketInfo" json:"policy_domain_kerberos_ticket_info"`
}

func (*PolicyDomainInformation_PolicyDomainKerberosTicketInfo) is_PolicyDomainInformation() {}

func (o *PolicyDomainInformation_PolicyDomainKerberosTicketInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PolicyDomainKerberosTicketInfo != nil {
		if err := o.PolicyDomainKerberosTicketInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PolicyDomainKerberosTicketInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyDomainInformation_PolicyDomainKerberosTicketInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PolicyDomainKerberosTicketInfo == nil {
		o.PolicyDomainKerberosTicketInfo = &PolicyDomainKerberosTicketInfo{}
	}
	if err := o.PolicyDomainKerberosTicketInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainNameInfo structure represents LSAPR_TRUSTED_DOMAIN_NAME_INFO RPC structure.
//
// The LSAPR_TRUSTED_DOMAIN_NAME_INFO structure is used to communicate the name of a
// trusted domain. The following structure corresponds to the TrustedDomainNameInformation
// information class.
type TrustedDomainNameInfo struct {
	// Name:  This field contains a name for the domain that is subject to the restrictions
	// of a NetBIOS name, as specified in [RFC1088]. This field SHOULD be used (by implementations
	// external to this protocol) to identify the domain via the NetBIOS API, as specified
	// in [RFC1088].
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
}

func (o *TrustedDomainNameInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainNameInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainNameInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedControllersInfo structure represents LSAPR_TRUSTED_CONTROLLERS_INFO RPC structure.
//
// The LSAPR_TRUSTED_CONTROLLERS_INFO structure is used to communicate a set of names
// of domain controllers (DCs) in a trusted domain. The following structure corresponds
// to the TrustedControllersInformation information class.
type TrustedControllersInfo struct {
	// Entries:  The count of names.<31>
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// Names:  This field contains an array of DC names that are subject to the restrictions
	// of a NetBIOS name, as specified in [RFC1088]. This field SHOULD be used (by implementations
	// external to this protocol) to identify the DCs via the NetBIOS API, as specified
	// in [RFC1088]. If the Entries field has a value other than 0, this field MUST NOT
	// be NULL.
	Names []*dtyp.UnicodeString `idl:"name:Names;size_is:(Entries)" json:"names"`
}

func (o *TrustedControllersInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.Names != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.Names))
	}
	if o.Entries > uint32(5) {
		return fmt.Errorf("Entries is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedControllersInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.Names != nil || o.Entries > 0 {
		_ptr_Names := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Names {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Names[i1] != nil {
					if err := o.Names[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Names, _ptr_Names); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedControllersInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_Names := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			if o.Names[i1] == nil {
				o.Names[i1] = &dtyp.UnicodeString{}
			}
			if err := o.Names[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Names := func(ptr interface{}) { o.Names = *ptr.(*[]*dtyp.UnicodeString) }
	if err := w.ReadPointer(&o.Names, _s_Names, _ptr_Names); err != nil {
		return err
	}
	return nil
}

// TrustedPasswordInfo structure represents LSAPR_TRUSTED_PASSWORD_INFO RPC structure.
//
// The LSAPR_TRUSTED_PASSWORD_INFO structure is used to communicate trust-authentication
// material. The following structure corresponds to the TrustedPasswordInformation information
// class.
type TrustedPasswordInfo struct {
	// Password:  The current authentication material. See section 2.2.6.1.
	Password *CRCipherValue `idl:"name:Password" json:"password"`
	// OldPassword:  The version prior to the current version of the authentication material.
	// See section 2.2.6.1.
	OldPassword *CRCipherValue `idl:"name:OldPassword" json:"old_password"`
}

func (o *TrustedPasswordInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedPasswordInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Password != nil {
		_ptr_Password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Password != nil {
				if err := o.Password.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Password, _ptr_Password); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.OldPassword != nil {
		_ptr_OldPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.OldPassword != nil {
				if err := o.OldPassword.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OldPassword, _ptr_OldPassword); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedPasswordInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Password == nil {
			o.Password = &CRCipherValue{}
		}
		if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Password := func(ptr interface{}) { o.Password = *ptr.(**CRCipherValue) }
	if err := w.ReadPointer(&o.Password, _s_Password, _ptr_Password); err != nil {
		return err
	}
	_ptr_OldPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.OldPassword == nil {
			o.OldPassword = &CRCipherValue{}
		}
		if err := o.OldPassword.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_OldPassword := func(ptr interface{}) { o.OldPassword = *ptr.(**CRCipherValue) }
	if err := w.ReadPointer(&o.OldPassword, _s_OldPassword, _ptr_OldPassword); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInformationEx structure represents LSAPR_TRUSTED_DOMAIN_INFORMATION_EX RPC structure.
//
// The LSAPR_TRUSTED_DOMAIN_INFORMATION_EX structure communicates properties of a trusted
// domain. The following structure corresponds to the TrustedDomainInformationEx information
// class. Domain trusts are specified in [MS-ADTS] section 6.1.6.
type TrustedDomainInformationEx struct {
	// Name:  The DNS name of the domain. Maps to the Name field, as specified in section
	// 3.1.1.5.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// FlatName:  The NetBIOS name of the trusted domain, as specified in [RFC1088]. Maps
	// to the Flat Name field, as specified in section 3.1.1.5.
	FlatName *dtyp.UnicodeString `idl:"name:FlatName" json:"flat_name"`
	// Sid:  The domain SID. Maps to the Security Identifier field, as specified in section
	// 3.1.1.5.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
	// TrustDirection:  This field contains bitmapped values that define the properties
	// of the direction of trust between the local domain and the named domain. One or more
	// of the valid flags can be set. If all bits are 0, the trust is said to be disabled.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | I | O |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// I: The trust is inbound.
	//
	// O: The trust is outbound.
	//
	// All other bits SHOULD be 0 and ignored upon receipt.
	TrustDirection uint32 `idl:"name:TrustDirection" json:"trust_direction"`
	// TrustType:  This field specifies the type of trust between the local domain and the
	// named domain.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | Trust with a Windows domain that is not running Active Directory.                |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | Trust with a Windows domain that is running Active Directory.                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000003 | Trust with a non–Windows-compliant Kerberos distribution, as specified in        |
	//	|            | [RFC4120].                                                                       |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 | Trust with a distributed computing environment (DCE) realm. This is a historical |
	//	|            | reference and is not used.                                                       |
	//	+------------+----------------------------------------------------------------------------------+
	//
	// Note  Other values SHOULD NOT be set.
	TrustType uint32 `idl:"name:TrustType" json:"trust_type"`
	// TrustAttributes:  This field contains bitmapped values that define the attributes
	// of the trust.<32>
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---------+---------+---+---------+---------+---------+---------+---------+---------+---------+---------+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 |    1    |    2    | 3 |    4    |    5    |    6    |    7    |    8    |    9    |    3    |    1    |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |         |         |   |         |         |         |         |         |         |    0    |         |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---------+---------+---+---------+---------+---------+---------+---------+---------+---------+---------+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---------+---------+---+---------+---------+---------+---------+---------+---------+---------+---------+
	//	| R | R | R | R | R | R | R | R | O | O | R | R | R | R | R | R | R | R | R | R | R | T A P T | T A N C | R | T A R C | T A T E | T A W F | T A C O | T A F T | T A Q D | T A U O | T A N T |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---------+---------+---+---------+---------+---------+---------+---------+---------+---------+---------+
	//
	// TrustAttribute values are described in section 3.1.1.5. The following table shows
	// how these values map to the Trust Attributes field in section 3.1.1.5.
	//
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                             |                                                                                  |
	//	|                            VALUE                            |                                     MAPPING                                      |
	//	|                                                             |                                                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TANT (TRUST_ATTRIBUTE_NON_TRANSITIVE)                       | Trust Attributes: Non-transitive                                                 |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TAUO (TRUST_ATTRIBUTE_UPLEVEL_ONLY)                         | Trust Attributes: Uplevel only                                                   |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TAQD (TRUST_ATTRIBUTE_QUARANTINED_DOMAIN)                   | Trust Attributes: Quarantined                                                    |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TAFT (TRUST_ATTRIBUTE_FOREST_TRANSITIVE)                    | Trust Attributes: Forest trust                                                   |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TACO (TRUST_ATTRIBUTE_CROSS_ORGANIZATION)                   | Trust Attributes: Cross organization                                             |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TAWF (TRUST_ATTRIBUTE_WITHIN_FOREST)                        | Trust Attributes: Within forest                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TATE (TRUST_ATTRIBUTE_TREAT_AS_EXTERNAL)                    | Trust Attributes: Treat as external                                              |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TARC (TRUST_ATTRIBUTE_USES_RC4_ENCRYPTION)                  | Trust Attributes: Use RC4 Encryption (for more information about RC4, see        |
	//	|                                                             | [SCHNEIER] section 17.1).                                                        |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TANC (TRUST_ATTRIBUTE_CROSS_ORGANIZATION_NO_TGT_DELEGATION) | Trust Attributes: Tokens must not be trusted for delegation.                     |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| TAPT (TRUST_ATTRIBUTE_PIM_TRUST)                            | Trust Attributes: PrivilegedIdentityManagement (PIM) trust.                      |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| O                                                           | Obsolete. SHOULD be set to 0.                                                    |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| R                                                           | Reserved for future use. SHOULD be set to zero.                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	TrustAttributes uint32 `idl:"name:TrustAttributes" json:"trust_attributes"`
}

func (o *TrustedDomainInformationEx) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInformationEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FlatName != nil {
		if err := o.FlatName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SID, _ptr_Sid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TrustDirection); err != nil {
		return err
	}
	if err := w.WriteData(o.TrustType); err != nil {
		return err
	}
	if err := w.WriteData(o.TrustAttributes); err != nil {
		return err
	}
	return nil
}
func (o *TrustedDomainInformationEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.FlatName == nil {
		o.FlatName = &dtyp.UnicodeString{}
	}
	if err := o.FlatName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustDirection); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustType); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustAttributes); err != nil {
		return err
	}
	return nil
}

// AuthInformation structure represents LSAPR_AUTH_INFORMATION RPC structure.
//
// The LSAPR_AUTH_INFORMATION structure communicates information about authentication
// between trusted domains. Domain trust authentication is specified in [MS-ADTS] section
// 6.1.6.9.1.
type AuthInformation struct {
	// LastUpdateTime:  The date and time when this authentication information was last
	// updated. It is a 64-bit value that represents the number of 100-nanosecond intervals
	// since January 1, 1601, UTC.
	LastUpdateTime *dtyp.LargeInteger `idl:"name:LastUpdateTime" json:"last_update_time"`
	// AuthType:  A type for the AuthInfo, as specified in the following table.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 | This type MUST be ignored.                                                       |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | Derived RC4HMAC key. For more information, see [RFC4757].                        |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | A plaintext password. Indicates that the information stored in the attribute     |
	//	|            | is a Unicode plaintext password. If this AuthType is present, Kerberos can then  |
	//	|            | use this password to derive additional key types that are needed to encrypt and  |
	//	|            | decrypt cross-realm TGTs.                                                        |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000003 | A plaintext password version number that is a single, unsigned long integer      |
	//	|            | consisting of 32 bits.                                                           |
	//	+------------+----------------------------------------------------------------------------------+
	AuthType uint32 `idl:"name:AuthType" json:"auth_type"`
	// AuthInfoLength:  The count of bytes in AuthInfo buffer.<36>
	AuthInfoLength uint32 `idl:"name:AuthInfoLength" json:"auth_info_length"`
	// AuthInfo:  Authentication data that depends on the AuthType.
	//
	// The self-relative form of the LSAPR_AUTH_INFORMATION structure is used in LSAPR_TRUSTED_DOMAIN_AUTH_BLOB;
	// in that case, the structure memory layout looks like the following.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| LastUpdateTime                                                                                                                |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| AuthType                                                                                                                      |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| AuthInfoLength                                                                                                                |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| AuthInfo [1 ... AuthInfoLength]                                                                                               |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	AuthInfo []byte `idl:"name:AuthInfo;size_is:(AuthInfoLength)" json:"auth_info"`
}

func (o *AuthInformation) xxx_PreparePayload(ctx context.Context) error {
	if o.AuthInfo != nil && o.AuthInfoLength == 0 {
		o.AuthInfoLength = uint32(len(o.AuthInfo))
	}
	if o.AuthInfoLength > uint32(65536) {
		return fmt.Errorf("AuthInfoLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.LastUpdateTime != nil {
		if err := o.LastUpdateTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AuthType); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthInfoLength); err != nil {
		return err
	}
	if o.AuthInfo != nil || o.AuthInfoLength > 0 {
		_ptr_AuthInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AuthInfoLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.AuthInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.AuthInfo[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.AuthInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AuthInfo, _ptr_AuthInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.LastUpdateTime == nil {
		o.LastUpdateTime = &dtyp.LargeInteger{}
	}
	if err := o.LastUpdateTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthInfoLength); err != nil {
		return err
	}
	_ptr_AuthInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AuthInfoLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AuthInfoLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AuthInfo", sizeInfo[0])
		}
		o.AuthInfo = make([]byte, sizeInfo[0])
		for i1 := range o.AuthInfo {
			i1 := i1
			if err := w.ReadData(&o.AuthInfo[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_AuthInfo := func(ptr interface{}) { o.AuthInfo = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.AuthInfo, _s_AuthInfo, _ptr_AuthInfo); err != nil {
		return err
	}
	return nil
}

// TrustedDomainAuthInformation structure represents LSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION RPC structure.
//
// The LSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION structure communicates authentication material.
// The following structure corresponds to the TrustedDomainAuthInformation information
// class. Domain trust authentication is specified in [MS-ADTS] section 6.1.6.9.1. This
// structure maps to the Incoming and Outgoing Trust Password fields, as specified in
// section 3.1.1.5.
type TrustedDomainAuthInformation struct {
	// IncomingAuthInfos:  The count of LSAPR_AUTH_INFORMATION entries (section 2.2.7.17)
	// in the IncomingAuthenticationInformation field.<33>
	IncomingAuthInfos uint32 `idl:"name:IncomingAuthInfos" json:"incoming_auth_infos"`
	// IncomingAuthenticationInformation:  An array of LSAPR_AUTH_INFORMATION structures.
	// The values are used to compute keys used in inbound trust validation, as specified
	// in [MS-ADTS] section 6.1.6.9.1.
	IncomingAuthenticationInformation *AuthInformation `idl:"name:IncomingAuthenticationInformation" json:"incoming_authentication_information"`
	// IncomingPreviousAuthenticationInformation:  Same as IncomingAuthenticationInformation,
	// but the data is the previous version of the authentication information.
	IncomingPreviousAuthenticationInformation *AuthInformation `idl:"name:IncomingPreviousAuthenticationInformation" json:"incoming_previous_authentication_information"`
	// OutgoingAuthInfos:  The count of LSAPR_AUTH_INFORMATION entries in the OutgoingAuthenticationInformation
	// field.<34>
	OutgoingAuthInfos uint32 `idl:"name:OutgoingAuthInfos" json:"outgoing_auth_infos"`
	// OutgoingAuthenticationInformation:  An array of LSAPR_AUTH_INFORMATION structures.
	// The values are used to compute keys used in outbound trust validation, as specified
	// in [MS-ADTS] section 6.1.6.9.1.
	OutgoingAuthenticationInformation *AuthInformation `idl:"name:OutgoingAuthenticationInformation" json:"outgoing_authentication_information"`
	// OutgoingPreviousAuthenticationInformation:  Same as OutgoingAuthenticationInformation,
	// but the data is the previous version of the authentication information.
	OutgoingPreviousAuthenticationInformation *AuthInformation `idl:"name:OutgoingPreviousAuthenticationInformation" json:"outgoing_previous_authentication_information"`
}

func (o *TrustedDomainAuthInformation) xxx_PreparePayload(ctx context.Context) error {
	if o.IncomingAuthInfos > uint32(1) {
		return fmt.Errorf("IncomingAuthInfos is out of range")
	}
	if o.OutgoingAuthInfos > uint32(1) {
		return fmt.Errorf("OutgoingAuthInfos is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainAuthInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.IncomingAuthInfos); err != nil {
		return err
	}
	if o.IncomingAuthenticationInformation != nil {
		_ptr_IncomingAuthenticationInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IncomingAuthenticationInformation != nil {
				if err := o.IncomingAuthenticationInformation.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AuthInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IncomingAuthenticationInformation, _ptr_IncomingAuthenticationInformation); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.IncomingPreviousAuthenticationInformation != nil {
		_ptr_IncomingPreviousAuthenticationInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IncomingPreviousAuthenticationInformation != nil {
				if err := o.IncomingPreviousAuthenticationInformation.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AuthInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IncomingPreviousAuthenticationInformation, _ptr_IncomingPreviousAuthenticationInformation); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.OutgoingAuthInfos); err != nil {
		return err
	}
	if o.OutgoingAuthenticationInformation != nil {
		_ptr_OutgoingAuthenticationInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.OutgoingAuthenticationInformation != nil {
				if err := o.OutgoingAuthenticationInformation.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AuthInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OutgoingAuthenticationInformation, _ptr_OutgoingAuthenticationInformation); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.OutgoingPreviousAuthenticationInformation != nil {
		_ptr_OutgoingPreviousAuthenticationInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.OutgoingPreviousAuthenticationInformation != nil {
				if err := o.OutgoingPreviousAuthenticationInformation.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AuthInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OutgoingPreviousAuthenticationInformation, _ptr_OutgoingPreviousAuthenticationInformation); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainAuthInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.IncomingAuthInfos); err != nil {
		return err
	}
	_ptr_IncomingAuthenticationInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IncomingAuthenticationInformation == nil {
			o.IncomingAuthenticationInformation = &AuthInformation{}
		}
		if err := o.IncomingAuthenticationInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IncomingAuthenticationInformation := func(ptr interface{}) { o.IncomingAuthenticationInformation = *ptr.(**AuthInformation) }
	if err := w.ReadPointer(&o.IncomingAuthenticationInformation, _s_IncomingAuthenticationInformation, _ptr_IncomingAuthenticationInformation); err != nil {
		return err
	}
	_ptr_IncomingPreviousAuthenticationInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IncomingPreviousAuthenticationInformation == nil {
			o.IncomingPreviousAuthenticationInformation = &AuthInformation{}
		}
		if err := o.IncomingPreviousAuthenticationInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IncomingPreviousAuthenticationInformation := func(ptr interface{}) { o.IncomingPreviousAuthenticationInformation = *ptr.(**AuthInformation) }
	if err := w.ReadPointer(&o.IncomingPreviousAuthenticationInformation, _s_IncomingPreviousAuthenticationInformation, _ptr_IncomingPreviousAuthenticationInformation); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutgoingAuthInfos); err != nil {
		return err
	}
	_ptr_OutgoingAuthenticationInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.OutgoingAuthenticationInformation == nil {
			o.OutgoingAuthenticationInformation = &AuthInformation{}
		}
		if err := o.OutgoingAuthenticationInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_OutgoingAuthenticationInformation := func(ptr interface{}) { o.OutgoingAuthenticationInformation = *ptr.(**AuthInformation) }
	if err := w.ReadPointer(&o.OutgoingAuthenticationInformation, _s_OutgoingAuthenticationInformation, _ptr_OutgoingAuthenticationInformation); err != nil {
		return err
	}
	_ptr_OutgoingPreviousAuthenticationInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.OutgoingPreviousAuthenticationInformation == nil {
			o.OutgoingPreviousAuthenticationInformation = &AuthInformation{}
		}
		if err := o.OutgoingPreviousAuthenticationInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_OutgoingPreviousAuthenticationInformation := func(ptr interface{}) { o.OutgoingPreviousAuthenticationInformation = *ptr.(**AuthInformation) }
	if err := w.ReadPointer(&o.OutgoingPreviousAuthenticationInformation, _s_OutgoingPreviousAuthenticationInformation, _ptr_OutgoingPreviousAuthenticationInformation); err != nil {
		return err
	}
	return nil
}

// TrustedDomainFullInformation structure represents LSAPR_TRUSTED_DOMAIN_FULL_INFORMATION RPC structure.
//
// The LSAPR_TRUSTED_DOMAIN_FULL_INFORMATION structure communicates identification,
// POSIX compatibility, and authentication information for a trusted domain. The following
// structure corresponds to the TrustedDomainFullInformation information class.
type TrustedDomainFullInformation struct {
	// Information:  A structure containing name, SID, and trust attributes, as specified
	// in section 2.2.7.9.
	Information *TrustedDomainInformationEx `idl:"name:Information" json:"information"`
	// PosixOffset:  Any offset required for POSIX compliance, as specified in section 2.2.7.6.
	POSIXOffset *TrustedPOSIXOffsetInfo `idl:"name:PosixOffset" json:"posix_offset"`
	// AuthInformation:  Authentication material, as specified in section 2.2.7.11.
	AuthInformation *TrustedDomainAuthInformation `idl:"name:AuthInformation" json:"auth_information"`
}

func (o *TrustedDomainFullInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainFullInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Information != nil {
		if err := o.Information.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainInformationEx{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.POSIXOffset != nil {
		if err := o.POSIXOffset.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedPOSIXOffsetInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AuthInformation != nil {
		if err := o.AuthInformation.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainAuthInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainFullInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Information == nil {
		o.Information = &TrustedDomainInformationEx{}
	}
	if err := o.Information.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.POSIXOffset == nil {
		o.POSIXOffset = &TrustedPOSIXOffsetInfo{}
	}
	if err := o.POSIXOffset.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AuthInformation == nil {
		o.AuthInformation = &TrustedDomainAuthInformation{}
	}
	if err := o.AuthInformation.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInformationBasic structure represents LSAPR_TRUSTED_DOMAIN_INFORMATION_BASIC RPC structure.
type TrustedDomainInformationBasic TrustInformation

func (o *TrustedDomainInformationBasic) TrustInformation() *TrustInformation {
	return (*TrustInformation)(o)
}

func (o *TrustedDomainInformationBasic) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInformationBasic) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SID, _ptr_Sid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInformationBasic) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	return nil
}

// TrustedDomainAuthBlob structure represents LSAPR_TRUSTED_DOMAIN_AUTH_BLOB RPC structure.
//
// The LSAPR_TRUSTED_DOMAIN_AUTH_BLOB structure contains a counted buffer of authentication
// material. Domain trust authentication is specified in [MS-ADTS] section 6.1.6.9.1.
type TrustedDomainAuthBlob struct {
	// AuthSize:  The count of bytes in AuthBlob.<35>
	AuthSize uint32 `idl:"name:AuthSize" json:"auth_size"`
	// AuthBlob:  An array of bytes containing the authentication material. If the AuthSize
	// field has a value other than 0, this field MUST NOT be NULL. Always encrypted using
	// algorithms, as specified in section 5.1.1. The plaintext layout is in the following
	// format.
	//
	// The incoming and outgoing authentication information buffer size included at the
	// end of the LSAPR_TRUSTED_DOMAIN_AUTH_BLOB can be used to extract the incoming and
	// outgoing authentication information buffers from the LSAPR_TRUSTED_DOMAIN_AUTH_BLOB.
	// Each of these buffers contains the byte offset to both the current and the previous
	// authentication information. This information can be used to extract current and (if
	// any) previous authentication information.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 512 bytes of random data (512 bytes)                                                                                          |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| CountOutgoingAuthInfos                                                                                                        |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ByteOffsetCurrentOutgoingAuthInfo                                                                                             |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ByteOffsetPreviousOutgoingAuthInfo                                                                                            |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| CurrentOutgoingAuthInfos (variable)                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| PreviousOutgoingAuthInfos (variable)                                                                                          |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| CountIncomingAuthInfos                                                                                                        |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ByteOffsetCurrentIncomingAuthInfo                                                                                             |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ByteOffsetPreviousIncomingAuthInfo                                                                                            |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| CurrentIncomingAuthInfos (variable)                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| PreviousIncomingAuthInfos (variable)                                                                                          |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ...                                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| OutgoingAuthInfoSize                                                                                                          |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| IncomingAuthInfoSize                                                                                                          |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// CountOutgoingAuthInfos (4 bytes): Specifies the count of entries present in the CurrentOutgoingAuthInfos
	// field. If optional field PreviousOutgoingAuthInfos is present, the number of entries
	// in PreviousOutgoingAuthInfos is also equal to CountOutgoingAuthInfos.
	//
	// ByteOffsetCurrentOutgoingAuthInfo (4 bytes): Specifies the byte offset from the beginning
	// of CountOutgoingAuthInfos to the start of the CurrentOutgoingAuthInfos field. If
	// CountOutgoingAuthInfos is 0, this field MUST be ignored.
	//
	// ByteOffsetPreviousOutgoingAuthInfo (4 bytes): Specifies the byte offset from the
	// beginning of CountOutgoingAuthInfos to the start of the PreviousOutgoingAuthInfos
	// field. If the difference between ByteOffsetPreviousOutgoingAuthInfo and OutgoingAuthInfoSize
	// is 0, the PreviousOutgoingAuthInfos field has zero entries.
	//
	// CurrentOutgoingAuthInfos: Contains an array of CountOutgoingAuthInfos of LSAPR_AUTH_INFORMATION
	// (section 2.2.7.17) entries in self-relative format. Each LSAPR_AUTH_INFORMATION entry
	// in the array MUST be 4-byte aligned. When it is necessary to insert unused padding
	// bytes into a buffer for data alignment, such bytes MUST be set to 0.
	//
	// PreviousOutgoingAuthInfos: Contains an array of CountOutgoingAuthInfos LSAPR_AUTH_INFORMATION
	// entries in self-relative format. See the comments for the ByteOffsetPreviousOutgoingAuthInfo
	// field to determine when this field is present. Each LSAPR_AUTH_INFORMATION entry
	// in the array MUST be 4-byte aligned. When it is necessary to insert unused padding
	// bytes into a buffer for data alignment, such bytes MUST be set to 0.
	//
	// CountIncomingAuthInfos (4 bytes): Specifies the count of entries present in the CurrentIncomingAuthInfos
	// field. If optional field PreviousIncomingAuthInfos is present, the number of entries
	// in PreviousIncomingAuthInfos is also equal to CountIncomingAuthInfos.
	//
	// ByteOffsetCurrentIncomingAuthInfo (4 bytes): Specifies the byte offset from the beginning
	// of CountIncomingAuthInfos to the start of the CurrentIncomingAuthInfos field. If
	// CountIncomingAuthInfos is 0, this field MUST be ignored.
	//
	// ByteOffsetPreviousIncomingAuthInfo (4 bytes): Specifies the byte offset from the
	// beginning of CountIncomingAuthInfos to the start of the PreviousIncomingAuthInfos
	// field. If the difference between ByteOffsetPreviousIncomingAuthInfo and IncomingAuthInfoSize
	// is 0, the PreviousIncomingAuthInfos field has zero entries.
	//
	// CurrentIncomingAuthInfos: Contains an array of CountIncomingAuthInfos LSAPR_AUTH_INFORMATION
	// entries in self-relative format. Each LSAPR_AUTH_INFORMATION entry in the array MUST
	// be 4-byte aligned. When it is necessary to insert unused padding bytes into a buffer
	// for data alignment, such bytes MUST be set to 0.
	//
	// PreviousIncomingAuthInfos: Contains an array of CountIncomingAuthInfos LSAPR_AUTH_INFORMATION
	// entries in self-relative format. See the comments for the ByteOffsetPreviousIncomingAuthInfo
	// field to determine when this field is present. Each LSAPR_AUTH_INFORMATION entry
	// in the array MUST be 4-byte aligned. When it is necessary to insert unused padding
	// bytes into a buffer for data alignment, such bytes MUST be set to 0.
	//
	// OutgoingAuthInfoSize (4 bytes): Specifies the size, in bytes, of the subportion of
	// the structure from the beginning of the CountOutgoingAuthInfos field through the
	// end of the of the PreviousOutgoingAuthInfos field.
	AuthBlob []byte `idl:"name:AuthBlob;size_is:(AuthSize)" json:"auth_blob"`
}

func (o *TrustedDomainAuthBlob) xxx_PreparePayload(ctx context.Context) error {
	if o.AuthBlob != nil && o.AuthSize == 0 {
		o.AuthSize = uint32(len(o.AuthBlob))
	}
	if o.AuthSize > uint32(65536) {
		return fmt.Errorf("AuthSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainAuthBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthSize); err != nil {
		return err
	}
	if o.AuthBlob != nil || o.AuthSize > 0 {
		_ptr_AuthBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AuthSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.AuthBlob {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.AuthBlob[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.AuthBlob); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AuthBlob, _ptr_AuthBlob); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainAuthBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthSize); err != nil {
		return err
	}
	_ptr_AuthBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AuthSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AuthSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AuthBlob", sizeInfo[0])
		}
		o.AuthBlob = make([]byte, sizeInfo[0])
		for i1 := range o.AuthBlob {
			i1 := i1
			if err := w.ReadData(&o.AuthBlob[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_AuthBlob := func(ptr interface{}) { o.AuthBlob = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.AuthBlob, _s_AuthBlob, _ptr_AuthBlob); err != nil {
		return err
	}
	return nil
}

// TrustedDomainAuthInformationInternal structure represents LSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION_INTERNAL RPC structure.
//
// The LSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION_INTERNAL structure communicates authentication
// material. The following structure corresponds to the TrustedDomainAuthInformationInternal
// information class. For more information about domain trust authentication material,
// see [MS-ADTS] section 6.1.6.9.1.
type TrustedDomainAuthInformationInternal struct {
	// AuthBlob:  An LSAPR_TRUSTED_DOMAIN_AUTH_BLOB.
	AuthBlob *TrustedDomainAuthBlob `idl:"name:AuthBlob" json:"auth_blob"`
}

func (o *TrustedDomainAuthInformationInternal) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainAuthInformationInternal) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.AuthBlob != nil {
		if err := o.AuthBlob.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainAuthBlob{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainAuthInformationInternal) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.AuthBlob == nil {
		o.AuthBlob = &TrustedDomainAuthBlob{}
	}
	if err := o.AuthBlob.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainFullInformationInternal structure represents LSAPR_TRUSTED_DOMAIN_FULL_INFORMATION_INTERNAL RPC structure.
//
// The LSAPR_TRUSTED_DOMAIN_FULL_INFORMATION_INTERNAL structure communicates identification
// and authentication information for a trusted domain. The following structure corresponds
// to the TrustedDomainFullInformationInternal information class.
type TrustedDomainFullInformationInternal struct {
	// Information:  A structure containing name, SID, and trust attributes, as specified
	// in section 2.2.7.9.
	Information *TrustedDomainInformationEx `idl:"name:Information" json:"information"`
	// PosixOffset:  Any offset required for POSIX compliance, as specified in section 2.2.7.6.
	POSIXOffset *TrustedPOSIXOffsetInfo `idl:"name:PosixOffset" json:"posix_offset"`
	// AuthInformation:  Authentication material, as specified in section 2.2.7.12.
	AuthInformation *TrustedDomainAuthInformationInternal `idl:"name:AuthInformation" json:"auth_information"`
}

func (o *TrustedDomainFullInformationInternal) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainFullInformationInternal) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Information != nil {
		if err := o.Information.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainInformationEx{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.POSIXOffset != nil {
		if err := o.POSIXOffset.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedPOSIXOffsetInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AuthInformation != nil {
		if err := o.AuthInformation.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainAuthInformationInternal{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainFullInformationInternal) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Information == nil {
		o.Information = &TrustedDomainInformationEx{}
	}
	if err := o.Information.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.POSIXOffset == nil {
		o.POSIXOffset = &TrustedPOSIXOffsetInfo{}
	}
	if err := o.POSIXOffset.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AuthInformation == nil {
		o.AuthInformation = &TrustedDomainAuthInformationInternal{}
	}
	if err := o.AuthInformation.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInformationEx2 structure represents LSAPR_TRUSTED_DOMAIN_INFORMATION_EX2 RPC structure.
//
// The LSAPR_TRUSTED_DOMAIN_INFORMATION_EX2 structure communicates properties of a trusted
// domain. The following structure corresponds to the TrustedDomainInformationEx2Internal
// information class. Domain trusts are specified in [MS-ADTS] section 6.1.6.
type TrustedDomainInformationEx2 struct {
	// Name:  The DNS name of the domain. Maps to the Name field, as specified in section
	// 3.1.1.5.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// FlatName:  The NetBIOS name of the trusted domain, as specified in [RFC1088]. Maps
	// to the Flat Name field, as specified in section 3.1.1.5.
	FlatName *dtyp.UnicodeString `idl:"name:FlatName" json:"flat_name"`
	// Sid:  The domain SID. Maps to the Security Identifier field, as specified in section
	// 3.1.1.5.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
	// TrustDirection:  This field contains bitmapped values that define the properties
	// of the direction of trust between the local domain and the named domain. See section
	// 2.2.7.9 for valid values and a description of each flag. Maps to the Trusted Direction
	// field, as specified in section 3.1.1.5.
	TrustDirection uint32 `idl:"name:TrustDirection" json:"trust_direction"`
	// TrustType:  This field specifies the type of trust between the local domain and the
	// named domain. See section 2.2.7.9 for valid values and a description of each value.
	// Maps to the Trusted Type field, as specified in section 3.1.1.5.
	TrustType uint32 `idl:"name:TrustType" json:"trust_type"`
	// TrustAttributes:  This field contains bitmapped values that define the attributes
	// of the trust. See section 2.2.7.9 for valid values and a description of each flag.
	// Maps to the Trusted Attributes field, as specified in section 3.1.1.5.
	TrustAttributes uint32 `idl:"name:TrustAttributes" json:"trust_attributes"`
	// ForestTrustLength:  The count of bytes in ForestTrustInfo.
	ForestTrustLength uint32 `idl:"name:ForestTrustLength" json:"forest_trust_length"`
	// ForestTrustInfo:  Binary data for the forest trust. For more information, see "Trust
	// Objects" in [MS-ADTS] section 6.1.6. Maps to the Forest Trust Information field,
	// as specified in section 3.1.1.5. Conversion from this binary format to the LSA_FOREST_TRUST_INFORMATION
	// format is specified in [MS-ADTS] section 6.1.6.9.3. If the ForestTrustLength field
	// has a value other than 0, this field MUST NOT be NULL.
	ForestTrustInfo []byte `idl:"name:ForestTrustInfo;size_is:(ForestTrustLength)" json:"forest_trust_info"`
}

func (o *TrustedDomainInformationEx2) xxx_PreparePayload(ctx context.Context) error {
	if o.ForestTrustInfo != nil && o.ForestTrustLength == 0 {
		o.ForestTrustLength = uint32(len(o.ForestTrustInfo))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInformationEx2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FlatName != nil {
		if err := o.FlatName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SID, _ptr_Sid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TrustDirection); err != nil {
		return err
	}
	if err := w.WriteData(o.TrustType); err != nil {
		return err
	}
	if err := w.WriteData(o.TrustAttributes); err != nil {
		return err
	}
	if err := w.WriteData(o.ForestTrustLength); err != nil {
		return err
	}
	if o.ForestTrustInfo != nil || o.ForestTrustLength > 0 {
		_ptr_ForestTrustInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ForestTrustLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ForestTrustInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.ForestTrustInfo[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.ForestTrustInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ForestTrustInfo, _ptr_ForestTrustInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInformationEx2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.FlatName == nil {
		o.FlatName = &dtyp.UnicodeString{}
	}
	if err := o.FlatName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustDirection); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustType); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustAttributes); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForestTrustLength); err != nil {
		return err
	}
	_ptr_ForestTrustInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ForestTrustLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ForestTrustLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ForestTrustInfo", sizeInfo[0])
		}
		o.ForestTrustInfo = make([]byte, sizeInfo[0])
		for i1 := range o.ForestTrustInfo {
			i1 := i1
			if err := w.ReadData(&o.ForestTrustInfo[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ForestTrustInfo := func(ptr interface{}) { o.ForestTrustInfo = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ForestTrustInfo, _s_ForestTrustInfo, _ptr_ForestTrustInfo); err != nil {
		return err
	}
	return nil
}

// TrustedDomainFullInformation2 structure represents LSAPR_TRUSTED_DOMAIN_FULL_INFORMATION2 RPC structure.
//
// The LSAPR_TRUSTED_DOMAIN_FULL_INFORMATION2 structure is used to communicate identification,
// POSIX compatibility, and authentication information for a trusted domain. The following
// structure corresponds to the TrustedDomainFullInformation2Internal information class.
type TrustedDomainFullInformation2 struct {
	// Information:  A structure containing name, SID, and trust attributes, as specified
	// in section 2.2.7.10.
	Information *TrustedDomainInformationEx2 `idl:"name:Information" json:"information"`
	// PosixOffset:  Any offset required for POSIX compliance, as specified in section 2.2.7.6.
	POSIXOffset *TrustedPOSIXOffsetInfo `idl:"name:PosixOffset" json:"posix_offset"`
	// AuthInformation:  Authentication material, as specified in section 2.2.7.11.
	AuthInformation *TrustedDomainAuthInformation `idl:"name:AuthInformation" json:"auth_information"`
}

func (o *TrustedDomainFullInformation2) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainFullInformation2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Information != nil {
		if err := o.Information.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainInformationEx2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.POSIXOffset != nil {
		if err := o.POSIXOffset.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedPOSIXOffsetInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AuthInformation != nil {
		if err := o.AuthInformation.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainAuthInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainFullInformation2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Information == nil {
		o.Information = &TrustedDomainInformationEx2{}
	}
	if err := o.Information.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.POSIXOffset == nil {
		o.POSIXOffset = &TrustedPOSIXOffsetInfo{}
	}
	if err := o.POSIXOffset.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AuthInformation == nil {
		o.AuthInformation = &TrustedDomainAuthInformation{}
	}
	if err := o.AuthInformation.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainSupportedEncryptionTypes structure represents TRUSTED_DOMAIN_SUPPORTED_ENCRYPTION_TYPES RPC structure.
//
// The TRUSTED_DOMAIN_SUPPORTED_ENCRYPTION_TYPES structure is used to present the encryption
// types that are allowed through a trust.
type TrustedDomainSupportedEncryptionTypes struct {
	// SupportedEncryptionTypes:  This field contains bitmapped values that define the encryption
	// types supported by this trust relationship. The flags can be set in any combination.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | S | A | R | M | C |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// C: Supports CRC32, as specified in [RFC3961] page 31.
	//
	// M: Supports RSA-MD5, as specified in [RFC3961] page 31.
	//
	// R: Supports RC4-HMAC-MD5, as specified in [RFC4757].
	//
	// A: Supports HMAC-SHA1-96-AES128, as specified in [RFC3961] page 31.
	//
	// S: Supports HMAC-SHA1-96-AES256, as specified in [RFC3961] page 31.
	SupportedEncryptionTypes uint32 `idl:"name:SupportedEncryptionTypes" json:"supported_encryption_types"`
}

func (o *TrustedDomainSupportedEncryptionTypes) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainSupportedEncryptionTypes) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SupportedEncryptionTypes); err != nil {
		return err
	}
	return nil
}
func (o *TrustedDomainSupportedEncryptionTypes) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SupportedEncryptionTypes); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union.
type TrustedDomainInfo struct {
	// Types that are assignable to Value
	//
	// *TrustedDomainInfo_TrustedDomainNameInfo
	// *TrustedDomainInfo_TrustedControllersInfo
	// *TrustedDomainInfo_TrustedPOSIXOffsetInfo
	// *TrustedDomainInfo_TrustedPasswordInfo
	// *TrustedDomainInfo_Basic
	// *TrustedDomainInfo_Ex
	// *TrustedDomainInfo_TrustedAuthInfo
	// *TrustedDomainInfo_TrustedFullInfo
	// *TrustedDomainInfo_TrustedAuthInfoInternal
	// *TrustedDomainInfo_TrustedFullInfoInternal
	// *TrustedDomainInfo_Ex2
	// *TrustedDomainInfo_TrustedFullInfo2
	// *TrustedDomainInfo_TrustedDomainSETs
	Value is_TrustedDomainInfo `json:"value"`
}

func (o *TrustedDomainInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *TrustedDomainInfo_TrustedDomainNameInfo:
		if value != nil {
			return value.TrustedDomainNameInfo
		}
	case *TrustedDomainInfo_TrustedControllersInfo:
		if value != nil {
			return value.TrustedControllersInfo
		}
	case *TrustedDomainInfo_TrustedPOSIXOffsetInfo:
		if value != nil {
			return value.TrustedPOSIXOffsetInfo
		}
	case *TrustedDomainInfo_TrustedPasswordInfo:
		if value != nil {
			return value.TrustedPasswordInfo
		}
	case *TrustedDomainInfo_Basic:
		if value != nil {
			return value.TrustedDomainInfoBasic
		}
	case *TrustedDomainInfo_Ex:
		if value != nil {
			return value.TrustedDomainInfoEx
		}
	case *TrustedDomainInfo_TrustedAuthInfo:
		if value != nil {
			return value.TrustedAuthInfo
		}
	case *TrustedDomainInfo_TrustedFullInfo:
		if value != nil {
			return value.TrustedFullInfo
		}
	case *TrustedDomainInfo_TrustedAuthInfoInternal:
		if value != nil {
			return value.TrustedAuthInfoInternal
		}
	case *TrustedDomainInfo_TrustedFullInfoInternal:
		if value != nil {
			return value.TrustedFullInfoInternal
		}
	case *TrustedDomainInfo_Ex2:
		if value != nil {
			return value.TrustedDomainInfoEx2
		}
	case *TrustedDomainInfo_TrustedFullInfo2:
		if value != nil {
			return value.TrustedFullInfo2
		}
	case *TrustedDomainInfo_TrustedDomainSETs:
		if value != nil {
			return value.TrustedDomainSETs
		}
	}
	return nil
}

type is_TrustedDomainInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_TrustedDomainInfo()
}

func (o *TrustedDomainInfo) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *TrustedDomainInfo_TrustedDomainNameInfo:
		return uint16(1)
	case *TrustedDomainInfo_TrustedControllersInfo:
		return uint16(2)
	case *TrustedDomainInfo_TrustedPOSIXOffsetInfo:
		return uint16(3)
	case *TrustedDomainInfo_TrustedPasswordInfo:
		return uint16(4)
	case *TrustedDomainInfo_Basic:
		return uint16(5)
	case *TrustedDomainInfo_Ex:
		return uint16(6)
	case *TrustedDomainInfo_TrustedAuthInfo:
		return uint16(7)
	case *TrustedDomainInfo_TrustedFullInfo:
		return uint16(8)
	case *TrustedDomainInfo_TrustedAuthInfoInternal:
		return uint16(9)
	case *TrustedDomainInfo_TrustedFullInfoInternal:
		return uint16(10)
	case *TrustedDomainInfo_Ex2:
		return uint16(11)
	case *TrustedDomainInfo_TrustedFullInfo2:
		return uint16(12)
	case *TrustedDomainInfo_TrustedDomainSETs:
		return uint16(13)
	}
	return uint16(0)
}

func (o *TrustedDomainInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*TrustedDomainInfo_TrustedDomainNameInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_TrustedDomainNameInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*TrustedDomainInfo_TrustedControllersInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_TrustedControllersInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*TrustedDomainInfo_TrustedPOSIXOffsetInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_TrustedPOSIXOffsetInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*TrustedDomainInfo_TrustedPasswordInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_TrustedPasswordInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5):
		_o, _ := o.Value.(*TrustedDomainInfo_Basic)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_Basic{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(6):
		_o, _ := o.Value.(*TrustedDomainInfo_Ex)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_Ex{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(7):
		_o, _ := o.Value.(*TrustedDomainInfo_TrustedAuthInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_TrustedAuthInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(8):
		_o, _ := o.Value.(*TrustedDomainInfo_TrustedFullInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_TrustedFullInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(9):
		_o, _ := o.Value.(*TrustedDomainInfo_TrustedAuthInfoInternal)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_TrustedAuthInfoInternal{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(10):
		_o, _ := o.Value.(*TrustedDomainInfo_TrustedFullInfoInternal)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_TrustedFullInfoInternal{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(11):
		_o, _ := o.Value.(*TrustedDomainInfo_Ex2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_Ex2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(12):
		_o, _ := o.Value.(*TrustedDomainInfo_TrustedFullInfo2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_TrustedFullInfo2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(13):
		_o, _ := o.Value.(*TrustedDomainInfo_TrustedDomainSETs)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo_TrustedDomainSETs{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *TrustedDomainInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &TrustedDomainInfo_TrustedDomainNameInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &TrustedDomainInfo_TrustedControllersInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &TrustedDomainInfo_TrustedPOSIXOffsetInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &TrustedDomainInfo_TrustedPasswordInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5):
		o.Value = &TrustedDomainInfo_Basic{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(6):
		o.Value = &TrustedDomainInfo_Ex{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(7):
		o.Value = &TrustedDomainInfo_TrustedAuthInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(8):
		o.Value = &TrustedDomainInfo_TrustedFullInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(9):
		o.Value = &TrustedDomainInfo_TrustedAuthInfoInternal{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(10):
		o.Value = &TrustedDomainInfo_TrustedFullInfoInternal{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(11):
		o.Value = &TrustedDomainInfo_Ex2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(12):
		o.Value = &TrustedDomainInfo_TrustedFullInfo2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(13):
		o.Value = &TrustedDomainInfo_TrustedDomainSETs{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// TrustedDomainInfo_TrustedDomainNameInfo structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 1
type TrustedDomainInfo_TrustedDomainNameInfo struct {
	TrustedDomainNameInfo *TrustedDomainNameInfo `idl:"name:TrustedDomainNameInfo" json:"trusted_domain_name_info"`
}

func (*TrustedDomainInfo_TrustedDomainNameInfo) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_TrustedDomainNameInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedDomainNameInfo != nil {
		if err := o.TrustedDomainNameInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainNameInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_TrustedDomainNameInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedDomainNameInfo == nil {
		o.TrustedDomainNameInfo = &TrustedDomainNameInfo{}
	}
	if err := o.TrustedDomainNameInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_TrustedControllersInfo structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 2
type TrustedDomainInfo_TrustedControllersInfo struct {
	TrustedControllersInfo *TrustedControllersInfo `idl:"name:TrustedControllersInfo" json:"trusted_controllers_info"`
}

func (*TrustedDomainInfo_TrustedControllersInfo) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_TrustedControllersInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedControllersInfo != nil {
		if err := o.TrustedControllersInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedControllersInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_TrustedControllersInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedControllersInfo == nil {
		o.TrustedControllersInfo = &TrustedControllersInfo{}
	}
	if err := o.TrustedControllersInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_TrustedPOSIXOffsetInfo structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 3
type TrustedDomainInfo_TrustedPOSIXOffsetInfo struct {
	TrustedPOSIXOffsetInfo *TrustedPOSIXOffsetInfo `idl:"name:TrustedPosixOffsetInfo" json:"trusted_posix_offset_info"`
}

func (*TrustedDomainInfo_TrustedPOSIXOffsetInfo) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_TrustedPOSIXOffsetInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedPOSIXOffsetInfo != nil {
		if err := o.TrustedPOSIXOffsetInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedPOSIXOffsetInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_TrustedPOSIXOffsetInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedPOSIXOffsetInfo == nil {
		o.TrustedPOSIXOffsetInfo = &TrustedPOSIXOffsetInfo{}
	}
	if err := o.TrustedPOSIXOffsetInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_TrustedPasswordInfo structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 4
type TrustedDomainInfo_TrustedPasswordInfo struct {
	TrustedPasswordInfo *TrustedPasswordInfo `idl:"name:TrustedPasswordInfo" json:"trusted_password_info"`
}

func (*TrustedDomainInfo_TrustedPasswordInfo) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_TrustedPasswordInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedPasswordInfo != nil {
		if err := o.TrustedPasswordInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedPasswordInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_TrustedPasswordInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedPasswordInfo == nil {
		o.TrustedPasswordInfo = &TrustedPasswordInfo{}
	}
	if err := o.TrustedPasswordInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_Basic structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 5
type TrustedDomainInfo_Basic struct {
	TrustedDomainInfoBasic *TrustedDomainInformationBasic `idl:"name:TrustedDomainInfoBasic" json:"trusted_domain_info_basic"`
}

func (*TrustedDomainInfo_Basic) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_Basic) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedDomainInfoBasic != nil {
		if err := o.TrustedDomainInfoBasic.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainInformationBasic{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_Basic) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedDomainInfoBasic == nil {
		o.TrustedDomainInfoBasic = &TrustedDomainInformationBasic{}
	}
	if err := o.TrustedDomainInfoBasic.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_Ex structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 6
type TrustedDomainInfo_Ex struct {
	TrustedDomainInfoEx *TrustedDomainInformationEx `idl:"name:TrustedDomainInfoEx" json:"trusted_domain_info_ex"`
}

func (*TrustedDomainInfo_Ex) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_Ex) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedDomainInfoEx != nil {
		if err := o.TrustedDomainInfoEx.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainInformationEx{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_Ex) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedDomainInfoEx == nil {
		o.TrustedDomainInfoEx = &TrustedDomainInformationEx{}
	}
	if err := o.TrustedDomainInfoEx.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_TrustedAuthInfo structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 7
type TrustedDomainInfo_TrustedAuthInfo struct {
	TrustedAuthInfo *TrustedDomainAuthInformation `idl:"name:TrustedAuthInfo" json:"trusted_auth_info"`
}

func (*TrustedDomainInfo_TrustedAuthInfo) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_TrustedAuthInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedAuthInfo != nil {
		if err := o.TrustedAuthInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainAuthInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_TrustedAuthInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedAuthInfo == nil {
		o.TrustedAuthInfo = &TrustedDomainAuthInformation{}
	}
	if err := o.TrustedAuthInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_TrustedFullInfo structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 8
type TrustedDomainInfo_TrustedFullInfo struct {
	TrustedFullInfo *TrustedDomainFullInformation `idl:"name:TrustedFullInfo" json:"trusted_full_info"`
}

func (*TrustedDomainInfo_TrustedFullInfo) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_TrustedFullInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedFullInfo != nil {
		if err := o.TrustedFullInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainFullInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_TrustedFullInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedFullInfo == nil {
		o.TrustedFullInfo = &TrustedDomainFullInformation{}
	}
	if err := o.TrustedFullInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_TrustedAuthInfoInternal structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 9
type TrustedDomainInfo_TrustedAuthInfoInternal struct {
	TrustedAuthInfoInternal *TrustedDomainAuthInformationInternal `idl:"name:TrustedAuthInfoInternal" json:"trusted_auth_info_internal"`
}

func (*TrustedDomainInfo_TrustedAuthInfoInternal) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_TrustedAuthInfoInternal) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedAuthInfoInternal != nil {
		if err := o.TrustedAuthInfoInternal.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainAuthInformationInternal{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_TrustedAuthInfoInternal) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedAuthInfoInternal == nil {
		o.TrustedAuthInfoInternal = &TrustedDomainAuthInformationInternal{}
	}
	if err := o.TrustedAuthInfoInternal.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_TrustedFullInfoInternal structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 10
type TrustedDomainInfo_TrustedFullInfoInternal struct {
	TrustedFullInfoInternal *TrustedDomainFullInformationInternal `idl:"name:TrustedFullInfoInternal" json:"trusted_full_info_internal"`
}

func (*TrustedDomainInfo_TrustedFullInfoInternal) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_TrustedFullInfoInternal) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedFullInfoInternal != nil {
		if err := o.TrustedFullInfoInternal.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainFullInformationInternal{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_TrustedFullInfoInternal) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedFullInfoInternal == nil {
		o.TrustedFullInfoInternal = &TrustedDomainFullInformationInternal{}
	}
	if err := o.TrustedFullInfoInternal.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_Ex2 structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 11
type TrustedDomainInfo_Ex2 struct {
	TrustedDomainInfoEx2 *TrustedDomainInformationEx2 `idl:"name:TrustedDomainInfoEx2" json:"trusted_domain_info_ex2"`
}

func (*TrustedDomainInfo_Ex2) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_Ex2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedDomainInfoEx2 != nil {
		if err := o.TrustedDomainInfoEx2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainInformationEx2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_Ex2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedDomainInfoEx2 == nil {
		o.TrustedDomainInfoEx2 = &TrustedDomainInformationEx2{}
	}
	if err := o.TrustedDomainInfoEx2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_TrustedFullInfo2 structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 12
type TrustedDomainInfo_TrustedFullInfo2 struct {
	TrustedFullInfo2 *TrustedDomainFullInformation2 `idl:"name:TrustedFullInfo2" json:"trusted_full_info2"`
}

func (*TrustedDomainInfo_TrustedFullInfo2) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_TrustedFullInfo2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedFullInfo2 != nil {
		if err := o.TrustedFullInfo2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainFullInformation2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_TrustedFullInfo2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedFullInfo2 == nil {
		o.TrustedFullInfo2 = &TrustedDomainFullInformation2{}
	}
	if err := o.TrustedFullInfo2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustedDomainInfo_TrustedDomainSETs structure represents LSAPR_TRUSTED_DOMAIN_INFO RPC union arm.
//
// It has following labels: 13
type TrustedDomainInfo_TrustedDomainSETs struct {
	TrustedDomainSETs *TrustedDomainSupportedEncryptionTypes `idl:"name:TrustedDomainSETs" json:"trusted_domain_sets"`
}

func (*TrustedDomainInfo_TrustedDomainSETs) is_TrustedDomainInfo() {}

func (o *TrustedDomainInfo_TrustedDomainSETs) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TrustedDomainSETs != nil {
		if err := o.TrustedDomainSETs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TrustedDomainSupportedEncryptionTypes{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedDomainInfo_TrustedDomainSETs) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TrustedDomainSETs == nil {
		o.TrustedDomainSETs = &TrustedDomainSupportedEncryptionTypes{}
	}
	if err := o.TrustedDomainSETs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserRightSet structure represents LSAPR_USER_RIGHT_SET RPC structure.
//
// The LSAPR_USER_RIGHT_SET structure specifies a collection of user rights.
type UserRightSet struct {
	// Entries:  This field contains the number of rights.<25>
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// UserRights:  An array of strings specifying the rights. These can be string names
	// corresponding to either privilege names or system access names, as specified in section
	// 3.1.1.2. If the Entries field has a value other than 0, this field MUST NOT be NULL.
	UserRights []*dtyp.UnicodeString `idl:"name:UserRights;size_is:(Entries)" json:"user_rights"`
}

func (o *UserRightSet) xxx_PreparePayload(ctx context.Context) error {
	if o.UserRights != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.UserRights))
	}
	if o.Entries > uint32(256) {
		return fmt.Errorf("Entries is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserRightSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.UserRights != nil || o.Entries > 0 {
		_ptr_UserRights := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.UserRights {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.UserRights[i1] != nil {
					if err := o.UserRights[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.UserRights); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UserRights, _ptr_UserRights); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserRightSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_UserRights := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.UserRights", sizeInfo[0])
		}
		o.UserRights = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.UserRights {
			i1 := i1
			if o.UserRights[i1] == nil {
				o.UserRights[i1] = &dtyp.UnicodeString{}
			}
			if err := o.UserRights[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_UserRights := func(ptr interface{}) { o.UserRights = *ptr.(*[]*dtyp.UnicodeString) }
	if err := w.ReadPointer(&o.UserRights, _s_UserRights, _ptr_UserRights); err != nil {
		return err
	}
	return nil
}

// TrustedEnumBufferEx structure represents LSAPR_TRUSTED_ENUM_BUFFER_EX RPC structure.
//
// The LSAPR_TRUSTED_ENUM_BUFFER_EX structure specifies a collection of trust information
// structures of type LSAPR_TRUSTED_DOMAIN_INFORMATION_EX.
type TrustedEnumBufferEx struct {
	// EntriesRead:  This field contains the number of trust information structures.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// EnumerationBuffer:  This field contains a set of structures that define the trust
	// information, as specified in section 2.2.7.9. If the EntriesRead field has a value
	// other than 0, this field MUST NOT be NULL.
	EnumerationBuffer []*TrustedDomainInformationEx `idl:"name:EnumerationBuffer;size_is:(EntriesRead)" json:"enumeration_buffer"`
}

func (o *TrustedEnumBufferEx) xxx_PreparePayload(ctx context.Context) error {
	if o.EnumerationBuffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.EnumerationBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedEnumBufferEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.EnumerationBuffer != nil || o.EntriesRead > 0 {
		_ptr_EnumerationBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.EnumerationBuffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.EnumerationBuffer[i1] != nil {
					if err := o.EnumerationBuffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TrustedDomainInformationEx{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.EnumerationBuffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&TrustedDomainInformationEx{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EnumerationBuffer, _ptr_EnumerationBuffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TrustedEnumBufferEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_EnumerationBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.EnumerationBuffer", sizeInfo[0])
		}
		o.EnumerationBuffer = make([]*TrustedDomainInformationEx, sizeInfo[0])
		for i1 := range o.EnumerationBuffer {
			i1 := i1
			if o.EnumerationBuffer[i1] == nil {
				o.EnumerationBuffer[i1] = &TrustedDomainInformationEx{}
			}
			if err := o.EnumerationBuffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_EnumerationBuffer := func(ptr interface{}) { o.EnumerationBuffer = *ptr.(*[]*TrustedDomainInformationEx) }
	if err := w.ReadPointer(&o.EnumerationBuffer, _s_EnumerationBuffer, _ptr_EnumerationBuffer); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultLsarpcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultLsarpcClient) Close(ctx context.Context, in *CloseRequest, opts ...dcerpc.CallOption) (*CloseResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) EnumeratePrivileges(ctx context.Context, in *EnumeratePrivilegesRequest, opts ...dcerpc.CallOption) (*EnumeratePrivilegesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumeratePrivilegesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) QuerySecurityObject(ctx context.Context, in *QuerySecurityObjectRequest, opts ...dcerpc.CallOption) (*QuerySecurityObjectResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QuerySecurityObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) SetSecurityObject(ctx context.Context, in *SetSecurityObjectRequest, opts ...dcerpc.CallOption) (*SetSecurityObjectResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSecurityObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) OpenPolicy(ctx context.Context, in *OpenPolicyRequest, opts ...dcerpc.CallOption) (*OpenPolicyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) QueryInformationPolicy(ctx context.Context, in *QueryInformationPolicyRequest, opts ...dcerpc.CallOption) (*QueryInformationPolicyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryInformationPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) SetInformationPolicy(ctx context.Context, in *SetInformationPolicyRequest, opts ...dcerpc.CallOption) (*SetInformationPolicyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInformationPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...dcerpc.CallOption) (*CreateAccountResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) EnumerateAccounts(ctx context.Context, in *EnumerateAccountsRequest, opts ...dcerpc.CallOption) (*EnumerateAccountsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateAccountsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) CreateTrustedDomain(ctx context.Context, in *CreateTrustedDomainRequest, opts ...dcerpc.CallOption) (*CreateTrustedDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateTrustedDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) EnumerateTrustedDomains(ctx context.Context, in *EnumerateTrustedDomainsRequest, opts ...dcerpc.CallOption) (*EnumerateTrustedDomainsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateTrustedDomainsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...dcerpc.CallOption) (*CreateSecretResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateSecretResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) OpenAccount(ctx context.Context, in *OpenAccountRequest, opts ...dcerpc.CallOption) (*OpenAccountResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) EnumeratePrivilegesAccount(ctx context.Context, in *EnumeratePrivilegesAccountRequest, opts ...dcerpc.CallOption) (*EnumeratePrivilegesAccountResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumeratePrivilegesAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) AddPrivilegesToAccount(ctx context.Context, in *AddPrivilegesToAccountRequest, opts ...dcerpc.CallOption) (*AddPrivilegesToAccountResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddPrivilegesToAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) RemovePrivilegesFromAccount(ctx context.Context, in *RemovePrivilegesFromAccountRequest, opts ...dcerpc.CallOption) (*RemovePrivilegesFromAccountResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemovePrivilegesFromAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) GetSystemAccessAccount(ctx context.Context, in *GetSystemAccessAccountRequest, opts ...dcerpc.CallOption) (*GetSystemAccessAccountResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSystemAccessAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) SetSystemAccessAccount(ctx context.Context, in *SetSystemAccessAccountRequest, opts ...dcerpc.CallOption) (*SetSystemAccessAccountResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSystemAccessAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) OpenTrustedDomain(ctx context.Context, in *OpenTrustedDomainRequest, opts ...dcerpc.CallOption) (*OpenTrustedDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenTrustedDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) QueryInfoTrustedDomain(ctx context.Context, in *QueryInfoTrustedDomainRequest, opts ...dcerpc.CallOption) (*QueryInfoTrustedDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryInfoTrustedDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) SetInformationTrustedDomain(ctx context.Context, in *SetInformationTrustedDomainRequest, opts ...dcerpc.CallOption) (*SetInformationTrustedDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInformationTrustedDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) OpenSecret(ctx context.Context, in *OpenSecretRequest, opts ...dcerpc.CallOption) (*OpenSecretResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenSecretResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) SetSecret(ctx context.Context, in *SetSecretRequest, opts ...dcerpc.CallOption) (*SetSecretResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSecretResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) QuerySecret(ctx context.Context, in *QuerySecretRequest, opts ...dcerpc.CallOption) (*QuerySecretResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QuerySecretResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) LookupPrivilegeValue(ctx context.Context, in *LookupPrivilegeValueRequest, opts ...dcerpc.CallOption) (*LookupPrivilegeValueResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupPrivilegeValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) LookupPrivilegeName(ctx context.Context, in *LookupPrivilegeNameRequest, opts ...dcerpc.CallOption) (*LookupPrivilegeNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupPrivilegeNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) LookupPrivilegeDisplayName(ctx context.Context, in *LookupPrivilegeDisplayNameRequest, opts ...dcerpc.CallOption) (*LookupPrivilegeDisplayNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupPrivilegeDisplayNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...dcerpc.CallOption) (*DeleteObjectResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) EnumerateAccountsWithUserRight(ctx context.Context, in *EnumerateAccountsWithUserRightRequest, opts ...dcerpc.CallOption) (*EnumerateAccountsWithUserRightResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateAccountsWithUserRightResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) EnumerateAccountRights(ctx context.Context, in *EnumerateAccountRightsRequest, opts ...dcerpc.CallOption) (*EnumerateAccountRightsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateAccountRightsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) AddAccountRights(ctx context.Context, in *AddAccountRightsRequest, opts ...dcerpc.CallOption) (*AddAccountRightsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddAccountRightsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) RemoveAccountRights(ctx context.Context, in *RemoveAccountRightsRequest, opts ...dcerpc.CallOption) (*RemoveAccountRightsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveAccountRightsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) QueryTrustedDomainInfo(ctx context.Context, in *QueryTrustedDomainInfoRequest, opts ...dcerpc.CallOption) (*QueryTrustedDomainInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryTrustedDomainInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) SetTrustedDomainInfo(ctx context.Context, in *SetTrustedDomainInfoRequest, opts ...dcerpc.CallOption) (*SetTrustedDomainInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetTrustedDomainInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) DeleteTrustedDomain(ctx context.Context, in *DeleteTrustedDomainRequest, opts ...dcerpc.CallOption) (*DeleteTrustedDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteTrustedDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) StorePrivateData(ctx context.Context, in *StorePrivateDataRequest, opts ...dcerpc.CallOption) (*StorePrivateDataResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StorePrivateDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) RetrievePrivateData(ctx context.Context, in *RetrievePrivateDataRequest, opts ...dcerpc.CallOption) (*RetrievePrivateDataResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RetrievePrivateDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) OpenPolicy2(ctx context.Context, in *OpenPolicy2Request, opts ...dcerpc.CallOption) (*OpenPolicy2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenPolicy2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) QueryInformationPolicy2(ctx context.Context, in *QueryInformationPolicy2Request, opts ...dcerpc.CallOption) (*QueryInformationPolicy2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryInformationPolicy2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) SetInformationPolicy2(ctx context.Context, in *SetInformationPolicy2Request, opts ...dcerpc.CallOption) (*SetInformationPolicy2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInformationPolicy2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) QueryTrustedDomainInfoByName(ctx context.Context, in *QueryTrustedDomainInfoByNameRequest, opts ...dcerpc.CallOption) (*QueryTrustedDomainInfoByNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryTrustedDomainInfoByNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) SetTrustedDomainInfoByName(ctx context.Context, in *SetTrustedDomainInfoByNameRequest, opts ...dcerpc.CallOption) (*SetTrustedDomainInfoByNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetTrustedDomainInfoByNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) EnumerateTrustedDomainsEx(ctx context.Context, in *EnumerateTrustedDomainsExRequest, opts ...dcerpc.CallOption) (*EnumerateTrustedDomainsExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateTrustedDomainsExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) CreateTrustedDomainEx(ctx context.Context, in *CreateTrustedDomainExRequest, opts ...dcerpc.CallOption) (*CreateTrustedDomainExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateTrustedDomainExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) QueryDomainInformationPolicy(ctx context.Context, in *QueryDomainInformationPolicyRequest, opts ...dcerpc.CallOption) (*QueryDomainInformationPolicyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryDomainInformationPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) SetDomainInformationPolicy(ctx context.Context, in *SetDomainInformationPolicyRequest, opts ...dcerpc.CallOption) (*SetDomainInformationPolicyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetDomainInformationPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) OpenTrustedDomainByName(ctx context.Context, in *OpenTrustedDomainByNameRequest, opts ...dcerpc.CallOption) (*OpenTrustedDomainByNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenTrustedDomainByNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) CreateTrustedDomainEx2(ctx context.Context, in *CreateTrustedDomainEx2Request, opts ...dcerpc.CallOption) (*CreateTrustedDomainEx2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateTrustedDomainEx2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) QueryForestTrustInformation(ctx context.Context, in *QueryForestTrustInformationRequest, opts ...dcerpc.CallOption) (*QueryForestTrustInformationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryForestTrustInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) SetForestTrustInformation(ctx context.Context, in *SetForestTrustInformationRequest, opts ...dcerpc.CallOption) (*SetForestTrustInformationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetForestTrustInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultLsarpcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewLsarpcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (LsarpcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(LsarpcSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultLsarpcClient{cc: cc}, nil
}

// xxx_CloseOperation structure represents the LsarClose operation
type xxx_CloseOperation struct {
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	Return int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseOperation) OpNum() int { return 0 }

func (o *xxx_CloseOperation) OpName() string { return "/lsarpc/v0/LsarClose" }

func (o *xxx_CloseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object != nil {
			if err := o.Object.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object == nil {
			o.Object = &Handle{}
		}
		if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object != nil {
			if err := o.Object.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object == nil {
			o.Object = &Handle{}
		}
		if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseRequest structure represents the LsarClose operation request
type CloseRequest struct {
	// ObjectHandle: The context handle to be freed. On response, it MUST be set to 0.
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
}

func (o *CloseRequest) xxx_ToOp(ctx context.Context) *xxx_CloseOperation {
	if o == nil {
		return &xxx_CloseOperation{}
	}
	return &xxx_CloseOperation{
		Object: o.Object,
	}
}

func (o *CloseRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
}
func (o *CloseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseResponse structure represents the LsarClose operation response
type CloseResponse struct {
	// ObjectHandle: The context handle to be freed. On response, it MUST be set to 0.
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	// Return: The LsarClose return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseResponse) xxx_ToOp(ctx context.Context) *xxx_CloseOperation {
	if o == nil {
		return &xxx_CloseOperation{}
	}
	return &xxx_CloseOperation{
		Object: o.Object,
		Return: o.Return,
	}
}

func (o *CloseResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
	o.Return = op.Return
}
func (o *CloseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumeratePrivilegesOperation structure represents the LsarEnumeratePrivileges operation
type xxx_EnumeratePrivilegesOperation struct {
	Policy                 *Handle              `idl:"name:PolicyHandle" json:"policy"`
	EnumerationContext     uint32               `idl:"name:EnumerationContext" json:"enumeration_context"`
	EnumerationBuffer      *PrivilegeEnumBuffer `idl:"name:EnumerationBuffer" json:"enumeration_buffer"`
	PreferredMaximumLength uint32               `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
	Return                 int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumeratePrivilegesOperation) OpNum() int { return 2 }

func (o *xxx_EnumeratePrivilegesOperation) OpName() string {
	return "/lsarpc/v0/LsarEnumeratePrivileges"
}

func (o *xxx_EnumeratePrivilegesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumeratePrivilegesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// PreferedMaximumLength {in} (1:(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumeratePrivilegesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// PreferedMaximumLength {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumeratePrivilegesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumeratePrivilegesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// EnumerationBuffer {out} (1:{alias=PLSAPR_PRIVILEGE_ENUM_BUFFER}*(1))(2:{alias=LSAPR_PRIVILEGE_ENUM_BUFFER}(struct))
	{
		if o.EnumerationBuffer != nil {
			if err := o.EnumerationBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PrivilegeEnumBuffer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumeratePrivilegesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// EnumerationBuffer {out} (1:{alias=PLSAPR_PRIVILEGE_ENUM_BUFFER,pointer=ref}*(1))(2:{alias=LSAPR_PRIVILEGE_ENUM_BUFFER}(struct))
	{
		if o.EnumerationBuffer == nil {
			o.EnumerationBuffer = &PrivilegeEnumBuffer{}
		}
		if err := o.EnumerationBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumeratePrivilegesRequest structure represents the LsarEnumeratePrivileges operation request
type EnumeratePrivilegesRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// EnumerationContext: A pointer to a context value that is used to resume enumeration,
	// if necessary.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// PreferedMaximumLength: A value that indicates the approximate size of the data to
	// be returned.
	PreferredMaximumLength uint32 `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
}

func (o *EnumeratePrivilegesRequest) xxx_ToOp(ctx context.Context) *xxx_EnumeratePrivilegesOperation {
	if o == nil {
		return &xxx_EnumeratePrivilegesOperation{}
	}
	return &xxx_EnumeratePrivilegesOperation{
		Policy:                 o.Policy,
		EnumerationContext:     o.EnumerationContext,
		PreferredMaximumLength: o.PreferredMaximumLength,
	}
}

func (o *EnumeratePrivilegesRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumeratePrivilegesOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.EnumerationContext = op.EnumerationContext
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *EnumeratePrivilegesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumeratePrivilegesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumeratePrivilegesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumeratePrivilegesResponse structure represents the LsarEnumeratePrivileges operation response
type EnumeratePrivilegesResponse struct {
	// EnumerationContext: A pointer to a context value that is used to resume enumeration,
	// if necessary.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// EnumerationBuffer: A pointer to a structure that will contain the results of the
	// enumeration.
	EnumerationBuffer *PrivilegeEnumBuffer `idl:"name:EnumerationBuffer" json:"enumeration_buffer"`
	// Return: The LsarEnumeratePrivileges return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumeratePrivilegesResponse) xxx_ToOp(ctx context.Context) *xxx_EnumeratePrivilegesOperation {
	if o == nil {
		return &xxx_EnumeratePrivilegesOperation{}
	}
	return &xxx_EnumeratePrivilegesOperation{
		EnumerationContext: o.EnumerationContext,
		EnumerationBuffer:  o.EnumerationBuffer,
		Return:             o.Return,
	}
}

func (o *EnumeratePrivilegesResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumeratePrivilegesOperation) {
	if o == nil {
		return
	}
	o.EnumerationContext = op.EnumerationContext
	o.EnumerationBuffer = op.EnumerationBuffer
	o.Return = op.Return
}
func (o *EnumeratePrivilegesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumeratePrivilegesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumeratePrivilegesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QuerySecurityObjectOperation structure represents the LsarQuerySecurityObject operation
type xxx_QuerySecurityObjectOperation struct {
	Object              *Handle               `idl:"name:ObjectHandle" json:"object"`
	SecurityInformation uint32                `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  *SrSecurityDescriptor `idl:"name:SecurityDescriptor" json:"security_descriptor"`
	Return              int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_QuerySecurityObjectOperation) OpNum() int { return 3 }

func (o *xxx_QuerySecurityObjectOperation) OpName() string {
	return "/lsarpc/v0/LsarQuerySecurityObject"
}

func (o *xxx_QuerySecurityObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySecurityObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ObjectHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object != nil {
			if err := o.Object.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecurityInformation {in} (1:{alias=SECURITY_INFORMATION, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySecurityObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ObjectHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object == nil {
			o.Object = &Handle{}
		}
		if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=SECURITY_INFORMATION, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySecurityObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySecurityObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SecurityDescriptor {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_SR_SECURITY_DESCRIPTOR}*(1))(3:{alias=LSAPR_SR_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor != nil {
			_ptr_SecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SecurityDescriptor != nil {
					if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SrSecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecurityDescriptor, _ptr_SecurityDescriptor); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySecurityObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SecurityDescriptor {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_SR_SECURITY_DESCRIPTOR,pointer=ref}*(1))(3:{alias=LSAPR_SR_SECURITY_DESCRIPTOR}(struct))
	{
		_ptr_SecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SecurityDescriptor == nil {
				o.SecurityDescriptor = &SrSecurityDescriptor{}
			}
			if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_SecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(**SrSecurityDescriptor) }
		if err := w.ReadPointer(&o.SecurityDescriptor, _s_SecurityDescriptor, _ptr_SecurityDescriptor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QuerySecurityObjectRequest structure represents the LsarQuerySecurityObject operation request
type QuerySecurityObjectRequest struct {
	// ObjectHandle: An open object handle of any type.
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	// SecurityInformation: A bitmask specifying which portions of the security descriptor
	// the caller is interested in.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
}

func (o *QuerySecurityObjectRequest) xxx_ToOp(ctx context.Context) *xxx_QuerySecurityObjectOperation {
	if o == nil {
		return &xxx_QuerySecurityObjectOperation{}
	}
	return &xxx_QuerySecurityObjectOperation{
		Object:              o.Object,
		SecurityInformation: o.SecurityInformation,
	}
}

func (o *QuerySecurityObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_QuerySecurityObjectOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
	o.SecurityInformation = op.SecurityInformation
}
func (o *QuerySecurityObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QuerySecurityObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySecurityObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QuerySecurityObjectResponse structure represents the LsarQuerySecurityObject operation response
type QuerySecurityObjectResponse struct {
	// SecurityDescriptor: Used to return the security descriptor containing the elements
	// requested by the caller.
	SecurityDescriptor *SrSecurityDescriptor `idl:"name:SecurityDescriptor" json:"security_descriptor"`
	// Return: The LsarQuerySecurityObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QuerySecurityObjectResponse) xxx_ToOp(ctx context.Context) *xxx_QuerySecurityObjectOperation {
	if o == nil {
		return &xxx_QuerySecurityObjectOperation{}
	}
	return &xxx_QuerySecurityObjectOperation{
		SecurityDescriptor: o.SecurityDescriptor,
		Return:             o.Return,
	}
}

func (o *QuerySecurityObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_QuerySecurityObjectOperation) {
	if o == nil {
		return
	}
	o.SecurityDescriptor = op.SecurityDescriptor
	o.Return = op.Return
}
func (o *QuerySecurityObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QuerySecurityObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySecurityObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSecurityObjectOperation structure represents the LsarSetSecurityObject operation
type xxx_SetSecurityObjectOperation struct {
	Object              *Handle               `idl:"name:ObjectHandle" json:"object"`
	SecurityInformation uint32                `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  *SrSecurityDescriptor `idl:"name:SecurityDescriptor" json:"security_descriptor"`
	Return              int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSecurityObjectOperation) OpNum() int { return 4 }

func (o *xxx_SetSecurityObjectOperation) OpName() string { return "/lsarpc/v0/LsarSetSecurityObject" }

func (o *xxx_SetSecurityObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ObjectHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object != nil {
			if err := o.Object.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecurityInformation {in} (1:{alias=SECURITY_INFORMATION, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// SecurityDescriptor {in} (1:{alias=PLSAPR_SR_SECURITY_DESCRIPTOR}*(1))(2:{alias=LSAPR_SR_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor != nil {
			if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SrSecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ObjectHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object == nil {
			o.Object = &Handle{}
		}
		if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=SECURITY_INFORMATION, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// SecurityDescriptor {in} (1:{alias=PLSAPR_SR_SECURITY_DESCRIPTOR,pointer=ref}*(1))(2:{alias=LSAPR_SR_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor == nil {
			o.SecurityDescriptor = &SrSecurityDescriptor{}
		}
		if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSecurityObjectRequest structure represents the LsarSetSecurityObject operation request
type SetSecurityObjectRequest struct {
	// ObjectHandle: An open handle to an existing object.
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	// SecurityInformation: A bitmask specifying which portions of the security descriptor
	// are to be set.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// SecurityDescriptor: The security descriptor to be set.
	SecurityDescriptor *SrSecurityDescriptor `idl:"name:SecurityDescriptor" json:"security_descriptor"`
}

func (o *SetSecurityObjectRequest) xxx_ToOp(ctx context.Context) *xxx_SetSecurityObjectOperation {
	if o == nil {
		return &xxx_SetSecurityObjectOperation{}
	}
	return &xxx_SetSecurityObjectOperation{
		Object:              o.Object,
		SecurityInformation: o.SecurityInformation,
		SecurityDescriptor:  o.SecurityDescriptor,
	}
}

func (o *SetSecurityObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSecurityObjectOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
	o.SecurityInformation = op.SecurityInformation
	o.SecurityDescriptor = op.SecurityDescriptor
}
func (o *SetSecurityObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSecurityObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecurityObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSecurityObjectResponse structure represents the LsarSetSecurityObject operation response
type SetSecurityObjectResponse struct {
	// Return: The LsarSetSecurityObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSecurityObjectResponse) xxx_ToOp(ctx context.Context) *xxx_SetSecurityObjectOperation {
	if o == nil {
		return &xxx_SetSecurityObjectOperation{}
	}
	return &xxx_SetSecurityObjectOperation{
		Return: o.Return,
	}
}

func (o *SetSecurityObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSecurityObjectOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetSecurityObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSecurityObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecurityObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenPolicyOperation structure represents the LsarOpenPolicy operation
type xxx_OpenPolicyOperation struct {
	SystemName       string            `idl:"name:SystemName;pointer:unique" json:"system_name"`
	ObjectAttributes *ObjectAttributes `idl:"name:ObjectAttributes" json:"object_attributes"`
	DesiredAccess    uint32            `idl:"name:DesiredAccess" json:"desired_access"`
	Policy           *Handle           `idl:"name:PolicyHandle" json:"policy"`
	Return           int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenPolicyOperation) OpNum() int { return 6 }

func (o *xxx_OpenPolicyOperation) OpName() string { return "/lsarpc/v0/LsarOpenPolicy" }

func (o *xxx_OpenPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SystemName {in} (1:{pointer=unique}*(1)[dim:0,string](wchar))
	{
		if o.SystemName != "" {
			_ptr_SystemName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.SystemName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SystemName, _ptr_SystemName); err != nil {
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
	// ObjectAttributes {in} (1:{alias=PLSAPR_OBJECT_ATTRIBUTES}*(1))(2:{alias=LSAPR_OBJECT_ATTRIBUTES}(struct))
	{
		if o.ObjectAttributes != nil {
			if err := o.ObjectAttributes.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectAttributes{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SystemName {in} (1:{pointer=unique}*(1)[dim:0,string](wchar))
	{
		_ptr_SystemName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.SystemName); err != nil {
				return err
			}
			return nil
		})
		_s_SystemName := func(ptr interface{}) { o.SystemName = *ptr.(*string) }
		if err := w.ReadPointer(&o.SystemName, _s_SystemName, _ptr_SystemName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ObjectAttributes {in} (1:{alias=PLSAPR_OBJECT_ATTRIBUTES,pointer=ref}*(1))(2:{alias=LSAPR_OBJECT_ATTRIBUTES}(struct))
	{
		if o.ObjectAttributes == nil {
			o.ObjectAttributes = &ObjectAttributes{}
		}
		if err := o.ObjectAttributes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenPolicyRequest structure represents the LsarOpenPolicy operation request
type OpenPolicyRequest struct {
	// SystemName: This parameter does not have any effect on message processing in any
	// environment. It MUST be ignored on receipt.
	SystemName string `idl:"name:SystemName;pointer:unique" json:"system_name"`
	// ObjectAttributes: This parameter does not have any effect on message processing in
	// any environment. All fields MUST<60> be ignored except RootDirectory, which MUST
	// be NULL.
	ObjectAttributes *ObjectAttributes `idl:"name:ObjectAttributes" json:"object_attributes"`
	// DesiredAccess: An ACCESS_MASK value that specifies the requested access rights that
	// MUST be granted on the returned PolicyHandle, if the request is successful.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *OpenPolicyRequest) xxx_ToOp(ctx context.Context) *xxx_OpenPolicyOperation {
	if o == nil {
		return &xxx_OpenPolicyOperation{}
	}
	return &xxx_OpenPolicyOperation{
		SystemName:       o.SystemName,
		ObjectAttributes: o.ObjectAttributes,
		DesiredAccess:    o.DesiredAccess,
	}
}

func (o *OpenPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenPolicyOperation) {
	if o == nil {
		return
	}
	o.SystemName = op.SystemName
	o.ObjectAttributes = op.ObjectAttributes
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenPolicyResponse structure represents the LsarOpenPolicy operation response
type OpenPolicyResponse struct {
	// PolicyHandle: An RPC context handle (as specified in section 2.2.2.1) that represents
	// a reference to the abstract data model of a policy object, as specified in section
	// 3.1.1.1.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// Return: The LsarOpenPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenPolicyResponse) xxx_ToOp(ctx context.Context) *xxx_OpenPolicyOperation {
	if o == nil {
		return &xxx_OpenPolicyOperation{}
	}
	return &xxx_OpenPolicyOperation{
		Policy: o.Policy,
		Return: o.Return,
	}
}

func (o *OpenPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenPolicyOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.Return = op.Return
}
func (o *OpenPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInformationPolicyOperation structure represents the LsarQueryInformationPolicy operation
type xxx_QueryInformationPolicyOperation struct {
	Policy            *Handle                `idl:"name:PolicyHandle" json:"policy"`
	InformationClass  PolicyInformationClass `idl:"name:InformationClass" json:"information_class"`
	PolicyInformation *PolicyInformation     `idl:"name:PolicyInformation;switch_is:InformationClass" json:"policy_information"`
	Return            int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInformationPolicyOperation) OpNum() int { return 7 }

func (o *xxx_QueryInformationPolicyOperation) OpName() string {
	return "/lsarpc/v0/LsarQueryInformationPolicy"
}

func (o *xxx_QueryInformationPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InformationClass {in} (1:{alias=POLICY_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=POLICY_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PolicyInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=PLSAPR_POLICY_INFORMATION}*(1))(3:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_INFORMATION}(union))
	{
		if o.PolicyInformation != nil {
			_ptr_PolicyInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swPolicyInformation := uint16(o.InformationClass)
				if o.PolicyInformation != nil {
					if err := o.PolicyInformation.MarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
						return err
					}
				} else {
					if err := (&PolicyInformation{}).MarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyInformation, _ptr_PolicyInformation); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PolicyInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=PLSAPR_POLICY_INFORMATION,pointer=ref}*(1))(3:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_INFORMATION}(union))
	{
		_ptr_PolicyInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyInformation == nil {
				o.PolicyInformation = &PolicyInformation{}
			}
			_swPolicyInformation := uint16(o.InformationClass)
			if err := o.PolicyInformation.UnmarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
				return err
			}
			return nil
		})
		_s_PolicyInformation := func(ptr interface{}) { o.PolicyInformation = *ptr.(**PolicyInformation) }
		if err := w.ReadPointer(&o.PolicyInformation, _s_PolicyInformation, _ptr_PolicyInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryInformationPolicyRequest structure represents the LsarQueryInformationPolicy operation request
type QueryInformationPolicyRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// InformationClass: A parameter that specifies what type of information the caller
	// is requesting.
	InformationClass PolicyInformationClass `idl:"name:InformationClass" json:"information_class"`
}

func (o *QueryInformationPolicyRequest) xxx_ToOp(ctx context.Context) *xxx_QueryInformationPolicyOperation {
	if o == nil {
		return &xxx_QueryInformationPolicyOperation{}
	}
	return &xxx_QueryInformationPolicyOperation{
		Policy:           o.Policy,
		InformationClass: o.InformationClass,
	}
}

func (o *QueryInformationPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationPolicyOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.InformationClass = op.InformationClass
}
func (o *QueryInformationPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInformationPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInformationPolicyResponse structure represents the LsarQueryInformationPolicy operation response
type QueryInformationPolicyResponse struct {
	// PolicyInformation: A parameter that references policy information structure on return.
	PolicyInformation *PolicyInformation `idl:"name:PolicyInformation;switch_is:InformationClass" json:"policy_information"`
	// Return: The LsarQueryInformationPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryInformationPolicyResponse) xxx_ToOp(ctx context.Context) *xxx_QueryInformationPolicyOperation {
	if o == nil {
		return &xxx_QueryInformationPolicyOperation{}
	}
	return &xxx_QueryInformationPolicyOperation{
		PolicyInformation: o.PolicyInformation,
		Return:            o.Return,
	}
}

func (o *QueryInformationPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationPolicyOperation) {
	if o == nil {
		return
	}
	o.PolicyInformation = op.PolicyInformation
	o.Return = op.Return
}
func (o *QueryInformationPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInformationPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInformationPolicyOperation structure represents the LsarSetInformationPolicy operation
type xxx_SetInformationPolicyOperation struct {
	Policy            *Handle                `idl:"name:PolicyHandle" json:"policy"`
	InformationClass  PolicyInformationClass `idl:"name:InformationClass" json:"information_class"`
	PolicyInformation *PolicyInformation     `idl:"name:PolicyInformation;switch_is:InformationClass" json:"policy_information"`
	Return            int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInformationPolicyOperation) OpNum() int { return 8 }

func (o *xxx_SetInformationPolicyOperation) OpName() string {
	return "/lsarpc/v0/LsarSetInformationPolicy"
}

func (o *xxx_SetInformationPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InformationClass {in} (1:{alias=POLICY_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	// PolicyInformation {in} (1:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=PLSAPR_POLICY_INFORMATION}*(1))(2:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_INFORMATION}(union))
	{
		_swPolicyInformation := uint16(o.InformationClass)
		if o.PolicyInformation != nil {
			if err := o.PolicyInformation.MarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation{}).MarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=POLICY_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	// PolicyInformation {in} (1:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=PLSAPR_POLICY_INFORMATION,pointer=ref}*(1))(2:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_INFORMATION}(union))
	{
		if o.PolicyInformation == nil {
			o.PolicyInformation = &PolicyInformation{}
		}
		_swPolicyInformation := uint16(o.InformationClass)
		if err := o.PolicyInformation.UnmarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInformationPolicyRequest structure represents the LsarSetInformationPolicy operation request
type SetInformationPolicyRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// InformationClass: A parameter that specifies what type of information the caller
	// is setting.
	InformationClass PolicyInformationClass `idl:"name:InformationClass" json:"information_class"`
	// PolicyInformation: Data that represents the policy being set.
	PolicyInformation *PolicyInformation `idl:"name:PolicyInformation;switch_is:InformationClass" json:"policy_information"`
}

func (o *SetInformationPolicyRequest) xxx_ToOp(ctx context.Context) *xxx_SetInformationPolicyOperation {
	if o == nil {
		return &xxx_SetInformationPolicyOperation{}
	}
	return &xxx_SetInformationPolicyOperation{
		Policy:            o.Policy,
		InformationClass:  o.InformationClass,
		PolicyInformation: o.PolicyInformation,
	}
}

func (o *SetInformationPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInformationPolicyOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.InformationClass = op.InformationClass
	o.PolicyInformation = op.PolicyInformation
}
func (o *SetInformationPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetInformationPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInformationPolicyResponse structure represents the LsarSetInformationPolicy operation response
type SetInformationPolicyResponse struct {
	// Return: The LsarSetInformationPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInformationPolicyResponse) xxx_ToOp(ctx context.Context) *xxx_SetInformationPolicyOperation {
	if o == nil {
		return &xxx_SetInformationPolicyOperation{}
	}
	return &xxx_SetInformationPolicyOperation{
		Return: o.Return,
	}
}

func (o *SetInformationPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInformationPolicyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetInformationPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetInformationPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateAccountOperation structure represents the LsarCreateAccount operation
type xxx_CreateAccountOperation struct {
	Policy        *Handle   `idl:"name:PolicyHandle" json:"policy"`
	AccountSID    *dtyp.SID `idl:"name:AccountSid" json:"account_sid"`
	DesiredAccess uint32    `idl:"name:DesiredAccess" json:"desired_access"`
	Account       *Handle   `idl:"name:AccountHandle" json:"account"`
	Return        int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateAccountOperation) OpNum() int { return 10 }

func (o *xxx_CreateAccountOperation) OpName() string { return "/lsarpc/v0/LsarCreateAccount" }

func (o *xxx_CreateAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// AccountSid {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.AccountSID != nil {
			if err := o.AccountSID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// AccountSid {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.AccountSID == nil {
			o.AccountSID = &dtyp.SID{}
		}
		if err := o.AccountSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// AccountHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account != nil {
			if err := o.Account.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// AccountHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account == nil {
			o.Account = &Handle{}
		}
		if err := o.Account.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateAccountRequest structure represents the LsarCreateAccount operation request
type CreateAccountRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// AccountSid: The security identifier (SID) of the account to be created.
	AccountSID *dtyp.SID `idl:"name:AccountSid" json:"account_sid"`
	// DesiredAccess: A bitmask specifying accesses to be granted to the newly created and
	// opened account at this time.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *CreateAccountRequest) xxx_ToOp(ctx context.Context) *xxx_CreateAccountOperation {
	if o == nil {
		return &xxx_CreateAccountOperation{}
	}
	return &xxx_CreateAccountOperation{
		Policy:        o.Policy,
		AccountSID:    o.AccountSID,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *CreateAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateAccountOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.AccountSID = op.AccountSID
	o.DesiredAccess = op.DesiredAccess
}
func (o *CreateAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateAccountResponse structure represents the LsarCreateAccount operation response
type CreateAccountResponse struct {
	// AccountHandle: Used to return a handle to the newly created account object.
	Account *Handle `idl:"name:AccountHandle" json:"account"`
	// Return: The LsarCreateAccount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateAccountResponse) xxx_ToOp(ctx context.Context) *xxx_CreateAccountOperation {
	if o == nil {
		return &xxx_CreateAccountOperation{}
	}
	return &xxx_CreateAccountOperation{
		Account: o.Account,
		Return:  o.Return,
	}
}

func (o *CreateAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateAccountOperation) {
	if o == nil {
		return
	}
	o.Account = op.Account
	o.Return = op.Return
}
func (o *CreateAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateAccountsOperation structure represents the LsarEnumerateAccounts operation
type xxx_EnumerateAccountsOperation struct {
	Policy                 *Handle            `idl:"name:PolicyHandle" json:"policy"`
	EnumerationContext     uint32             `idl:"name:EnumerationContext" json:"enumeration_context"`
	EnumerationBuffer      *AccountEnumBuffer `idl:"name:EnumerationBuffer" json:"enumeration_buffer"`
	PreferredMaximumLength uint32             `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
	Return                 int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateAccountsOperation) OpNum() int { return 11 }

func (o *xxx_EnumerateAccountsOperation) OpName() string { return "/lsarpc/v0/LsarEnumerateAccounts" }

func (o *xxx_EnumerateAccountsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// PreferedMaximumLength {in} (1:(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// PreferedMaximumLength {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// EnumerationBuffer {out} (1:{alias=PLSAPR_ACCOUNT_ENUM_BUFFER}*(1))(2:{alias=LSAPR_ACCOUNT_ENUM_BUFFER}(struct))
	{
		if o.EnumerationBuffer != nil {
			if err := o.EnumerationBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AccountEnumBuffer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// EnumerationBuffer {out} (1:{alias=PLSAPR_ACCOUNT_ENUM_BUFFER,pointer=ref}*(1))(2:{alias=LSAPR_ACCOUNT_ENUM_BUFFER}(struct))
	{
		if o.EnumerationBuffer == nil {
			o.EnumerationBuffer = &AccountEnumBuffer{}
		}
		if err := o.EnumerationBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateAccountsRequest structure represents the LsarEnumerateAccounts operation request
type EnumerateAccountsRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// EnumerationContext: A pointer to a context value that is used to resume enumeration,
	// if necessary.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// PreferedMaximumLength: A value that indicates the approximate size of the data to
	// return.
	PreferredMaximumLength uint32 `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
}

func (o *EnumerateAccountsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateAccountsOperation {
	if o == nil {
		return &xxx_EnumerateAccountsOperation{}
	}
	return &xxx_EnumerateAccountsOperation{
		Policy:                 o.Policy,
		EnumerationContext:     o.EnumerationContext,
		PreferredMaximumLength: o.PreferredMaximumLength,
	}
}

func (o *EnumerateAccountsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateAccountsOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.EnumerationContext = op.EnumerationContext
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *EnumerateAccountsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateAccountsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateAccountsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateAccountsResponse structure represents the LsarEnumerateAccounts operation response
type EnumerateAccountsResponse struct {
	// EnumerationContext: A pointer to a context value that is used to resume enumeration,
	// if necessary.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// EnumerationBuffer: A pointer to a structure that will contain the results of the
	// enumeration.
	EnumerationBuffer *AccountEnumBuffer `idl:"name:EnumerationBuffer" json:"enumeration_buffer"`
	// Return: The LsarEnumerateAccounts return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateAccountsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateAccountsOperation {
	if o == nil {
		return &xxx_EnumerateAccountsOperation{}
	}
	return &xxx_EnumerateAccountsOperation{
		EnumerationContext: o.EnumerationContext,
		EnumerationBuffer:  o.EnumerationBuffer,
		Return:             o.Return,
	}
}

func (o *EnumerateAccountsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateAccountsOperation) {
	if o == nil {
		return
	}
	o.EnumerationContext = op.EnumerationContext
	o.EnumerationBuffer = op.EnumerationBuffer
	o.Return = op.Return
}
func (o *EnumerateAccountsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateAccountsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateAccountsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateTrustedDomainOperation structure represents the LsarCreateTrustedDomain operation
type xxx_CreateTrustedDomainOperation struct {
	Policy                   *Handle           `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainInformation *TrustInformation `idl:"name:TrustedDomainInformation" json:"trusted_domain_information"`
	DesiredAccess            uint32            `idl:"name:DesiredAccess" json:"desired_access"`
	TrustedDomain            *Handle           `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	Return                   int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateTrustedDomainOperation) OpNum() int { return 12 }

func (o *xxx_CreateTrustedDomainOperation) OpName() string {
	return "/lsarpc/v0/LsarCreateTrustedDomain"
}

func (o *xxx_CreateTrustedDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainInformation {in} (1:{alias=PLSAPR_TRUST_INFORMATION}*(1))(2:{alias=LSAPR_TRUST_INFORMATION}(struct))
	{
		if o.TrustedDomainInformation != nil {
			if err := o.TrustedDomainInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainInformation {in} (1:{alias=PLSAPR_TRUST_INFORMATION,pointer=ref}*(1))(2:{alias=LSAPR_TRUST_INFORMATION}(struct))
	{
		if o.TrustedDomainInformation == nil {
			o.TrustedDomainInformation = &TrustInformation{}
		}
		if err := o.TrustedDomainInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TrustedDomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain != nil {
			if err := o.TrustedDomain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TrustedDomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain == nil {
			o.TrustedDomain = &Handle{}
		}
		if err := o.TrustedDomain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateTrustedDomainRequest structure represents the LsarCreateTrustedDomain operation request
type CreateTrustedDomainRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainInformation: Information about the new trusted domain object (TDO) to
	// be created.
	TrustedDomainInformation *TrustInformation `idl:"name:TrustedDomainInformation" json:"trusted_domain_information"`
	// DesiredAccess: An access mask that specifies the desired access to the TDO handle.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *CreateTrustedDomainRequest) xxx_ToOp(ctx context.Context) *xxx_CreateTrustedDomainOperation {
	if o == nil {
		return &xxx_CreateTrustedDomainOperation{}
	}
	return &xxx_CreateTrustedDomainOperation{
		Policy:                   o.Policy,
		TrustedDomainInformation: o.TrustedDomainInformation,
		DesiredAccess:            o.DesiredAccess,
	}
}

func (o *CreateTrustedDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateTrustedDomainOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainInformation = op.TrustedDomainInformation
	o.DesiredAccess = op.DesiredAccess
}
func (o *CreateTrustedDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateTrustedDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateTrustedDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateTrustedDomainResponse structure represents the LsarCreateTrustedDomain operation response
type CreateTrustedDomainResponse struct {
	// TrustedDomainHandle: Used to return the handle for the newly created TDO.
	TrustedDomain *Handle `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	// Return: The LsarCreateTrustedDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateTrustedDomainResponse) xxx_ToOp(ctx context.Context) *xxx_CreateTrustedDomainOperation {
	if o == nil {
		return &xxx_CreateTrustedDomainOperation{}
	}
	return &xxx_CreateTrustedDomainOperation{
		TrustedDomain: o.TrustedDomain,
		Return:        o.Return,
	}
}

func (o *CreateTrustedDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateTrustedDomainOperation) {
	if o == nil {
		return
	}
	o.TrustedDomain = op.TrustedDomain
	o.Return = op.Return
}
func (o *CreateTrustedDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateTrustedDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateTrustedDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateTrustedDomainsOperation structure represents the LsarEnumerateTrustedDomains operation
type xxx_EnumerateTrustedDomainsOperation struct {
	Policy                 *Handle            `idl:"name:PolicyHandle" json:"policy"`
	EnumerationContext     uint32             `idl:"name:EnumerationContext" json:"enumeration_context"`
	EnumerationBuffer      *TrustedEnumBuffer `idl:"name:EnumerationBuffer" json:"enumeration_buffer"`
	PreferredMaximumLength uint32             `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
	Return                 int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateTrustedDomainsOperation) OpNum() int { return 13 }

func (o *xxx_EnumerateTrustedDomainsOperation) OpName() string {
	return "/lsarpc/v0/LsarEnumerateTrustedDomains"
}

func (o *xxx_EnumerateTrustedDomainsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateTrustedDomainsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// PreferedMaximumLength {in} (1:(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateTrustedDomainsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// PreferedMaximumLength {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateTrustedDomainsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateTrustedDomainsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// EnumerationBuffer {out} (1:{alias=PLSAPR_TRUSTED_ENUM_BUFFER}*(1))(2:{alias=LSAPR_TRUSTED_ENUM_BUFFER}(struct))
	{
		if o.EnumerationBuffer != nil {
			if err := o.EnumerationBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedEnumBuffer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateTrustedDomainsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// EnumerationBuffer {out} (1:{alias=PLSAPR_TRUSTED_ENUM_BUFFER,pointer=ref}*(1))(2:{alias=LSAPR_TRUSTED_ENUM_BUFFER}(struct))
	{
		if o.EnumerationBuffer == nil {
			o.EnumerationBuffer = &TrustedEnumBuffer{}
		}
		if err := o.EnumerationBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateTrustedDomainsRequest structure represents the LsarEnumerateTrustedDomains operation request
type EnumerateTrustedDomainsRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// EnumerationContext: A pointer to a context value that is used to resume enumeration,
	// if necessary.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// PreferedMaximumLength: A value that indicates the approximate size of the data to
	// be returned.
	PreferredMaximumLength uint32 `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
}

func (o *EnumerateTrustedDomainsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateTrustedDomainsOperation {
	if o == nil {
		return &xxx_EnumerateTrustedDomainsOperation{}
	}
	return &xxx_EnumerateTrustedDomainsOperation{
		Policy:                 o.Policy,
		EnumerationContext:     o.EnumerationContext,
		PreferredMaximumLength: o.PreferredMaximumLength,
	}
}

func (o *EnumerateTrustedDomainsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateTrustedDomainsOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.EnumerationContext = op.EnumerationContext
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *EnumerateTrustedDomainsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateTrustedDomainsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateTrustedDomainsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateTrustedDomainsResponse structure represents the LsarEnumerateTrustedDomains operation response
type EnumerateTrustedDomainsResponse struct {
	// EnumerationContext: A pointer to a context value that is used to resume enumeration,
	// if necessary.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// EnumerationBuffer: A pointer to a structure that will contain the results of the
	// enumeration.
	EnumerationBuffer *TrustedEnumBuffer `idl:"name:EnumerationBuffer" json:"enumeration_buffer"`
	// Return: The LsarEnumerateTrustedDomains return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateTrustedDomainsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateTrustedDomainsOperation {
	if o == nil {
		return &xxx_EnumerateTrustedDomainsOperation{}
	}
	return &xxx_EnumerateTrustedDomainsOperation{
		EnumerationContext: o.EnumerationContext,
		EnumerationBuffer:  o.EnumerationBuffer,
		Return:             o.Return,
	}
}

func (o *EnumerateTrustedDomainsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateTrustedDomainsOperation) {
	if o == nil {
		return
	}
	o.EnumerationContext = op.EnumerationContext
	o.EnumerationBuffer = op.EnumerationBuffer
	o.Return = op.Return
}
func (o *EnumerateTrustedDomainsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateTrustedDomainsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateTrustedDomainsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateSecretOperation structure represents the LsarCreateSecret operation
type xxx_CreateSecretOperation struct {
	Policy        *Handle             `idl:"name:PolicyHandle" json:"policy"`
	SecretName    *dtyp.UnicodeString `idl:"name:SecretName" json:"secret_name"`
	DesiredAccess uint32              `idl:"name:DesiredAccess" json:"desired_access"`
	Secret        *Handle             `idl:"name:SecretHandle" json:"secret"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateSecretOperation) OpNum() int { return 16 }

func (o *xxx_CreateSecretOperation) OpName() string { return "/lsarpc/v0/LsarCreateSecret" }

func (o *xxx_CreateSecretOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSecretOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecretName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.SecretName != nil {
			if err := o.SecretName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSecretOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecretName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.SecretName == nil {
			o.SecretName = &dtyp.UnicodeString{}
		}
		if err := o.SecretName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSecretOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSecretOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SecretHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Secret != nil {
			if err := o.Secret.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateSecretOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SecretHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Secret == nil {
			o.Secret = &Handle{}
		}
		if err := o.Secret.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateSecretRequest structure represents the LsarCreateSecret operation request
type CreateSecretRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// SecretName: The name of the secret object to be created.
	SecretName *dtyp.UnicodeString `idl:"name:SecretName" json:"secret_name"`
	// DesiredAccess: A bitmask that specifies accesses to be granted to the newly created
	// and opened secret object at this time.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *CreateSecretRequest) xxx_ToOp(ctx context.Context) *xxx_CreateSecretOperation {
	if o == nil {
		return &xxx_CreateSecretOperation{}
	}
	return &xxx_CreateSecretOperation{
		Policy:        o.Policy,
		SecretName:    o.SecretName,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *CreateSecretRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateSecretOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.SecretName = op.SecretName
	o.DesiredAccess = op.DesiredAccess
}
func (o *CreateSecretRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateSecretRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateSecretOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateSecretResponse structure represents the LsarCreateSecret operation response
type CreateSecretResponse struct {
	// SecretHandle: Used to return a handle to the newly created secret object.
	Secret *Handle `idl:"name:SecretHandle" json:"secret"`
	// Return: The LsarCreateSecret return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateSecretResponse) xxx_ToOp(ctx context.Context) *xxx_CreateSecretOperation {
	if o == nil {
		return &xxx_CreateSecretOperation{}
	}
	return &xxx_CreateSecretOperation{
		Secret: o.Secret,
		Return: o.Return,
	}
}

func (o *CreateSecretResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateSecretOperation) {
	if o == nil {
		return
	}
	o.Secret = op.Secret
	o.Return = op.Return
}
func (o *CreateSecretResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateSecretResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateSecretOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenAccountOperation structure represents the LsarOpenAccount operation
type xxx_OpenAccountOperation struct {
	Policy        *Handle   `idl:"name:PolicyHandle" json:"policy"`
	AccountSID    *dtyp.SID `idl:"name:AccountSid" json:"account_sid"`
	DesiredAccess uint32    `idl:"name:DesiredAccess" json:"desired_access"`
	Account       *Handle   `idl:"name:AccountHandle" json:"account"`
	Return        int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenAccountOperation) OpNum() int { return 17 }

func (o *xxx_OpenAccountOperation) OpName() string { return "/lsarpc/v0/LsarOpenAccount" }

func (o *xxx_OpenAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// AccountSid {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.AccountSID != nil {
			if err := o.AccountSID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// AccountSid {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.AccountSID == nil {
			o.AccountSID = &dtyp.SID{}
		}
		if err := o.AccountSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// AccountHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account != nil {
			if err := o.Account.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// AccountHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account == nil {
			o.Account = &Handle{}
		}
		if err := o.Account.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenAccountRequest structure represents the LsarOpenAccount operation request
type OpenAccountRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// AccountSid: A SID of the account to be opened.
	AccountSID *dtyp.SID `idl:"name:AccountSid" json:"account_sid"`
	// DesiredAccess: A bitmask specifying accesses to be granted to the opened account
	// at this time.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *OpenAccountRequest) xxx_ToOp(ctx context.Context) *xxx_OpenAccountOperation {
	if o == nil {
		return &xxx_OpenAccountOperation{}
	}
	return &xxx_OpenAccountOperation{
		Policy:        o.Policy,
		AccountSID:    o.AccountSID,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *OpenAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenAccountOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.AccountSID = op.AccountSID
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenAccountResponse structure represents the LsarOpenAccount operation response
type OpenAccountResponse struct {
	// AccountHandle: Used to return a handle to the opened account object.
	Account *Handle `idl:"name:AccountHandle" json:"account"`
	// Return: The LsarOpenAccount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenAccountResponse) xxx_ToOp(ctx context.Context) *xxx_OpenAccountOperation {
	if o == nil {
		return &xxx_OpenAccountOperation{}
	}
	return &xxx_OpenAccountOperation{
		Account: o.Account,
		Return:  o.Return,
	}
}

func (o *OpenAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenAccountOperation) {
	if o == nil {
		return
	}
	o.Account = op.Account
	o.Return = op.Return
}
func (o *OpenAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumeratePrivilegesAccountOperation structure represents the LsarEnumeratePrivilegesAccount operation
type xxx_EnumeratePrivilegesAccountOperation struct {
	Account    *Handle       `idl:"name:AccountHandle" json:"account"`
	Privileges *PrivilegeSet `idl:"name:Privileges" json:"privileges"`
	Return     int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumeratePrivilegesAccountOperation) OpNum() int { return 18 }

func (o *xxx_EnumeratePrivilegesAccountOperation) OpName() string {
	return "/lsarpc/v0/LsarEnumeratePrivilegesAccount"
}

func (o *xxx_EnumeratePrivilegesAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumeratePrivilegesAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AccountHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account != nil {
			if err := o.Account.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EnumeratePrivilegesAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AccountHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account == nil {
			o.Account = &Handle{}
		}
		if err := o.Account.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumeratePrivilegesAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumeratePrivilegesAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Privileges {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_PRIVILEGE_SET}*(1))(3:{alias=LSAPR_PRIVILEGE_SET}(struct))
	{
		if o.Privileges != nil {
			_ptr_Privileges := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Privileges != nil {
					if err := o.Privileges.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PrivilegeSet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Privileges, _ptr_Privileges); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumeratePrivilegesAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Privileges {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_PRIVILEGE_SET,pointer=ref}*(1))(3:{alias=LSAPR_PRIVILEGE_SET}(struct))
	{
		_ptr_Privileges := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Privileges == nil {
				o.Privileges = &PrivilegeSet{}
			}
			if err := o.Privileges.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Privileges := func(ptr interface{}) { o.Privileges = *ptr.(**PrivilegeSet) }
		if err := w.ReadPointer(&o.Privileges, _s_Privileges, _ptr_Privileges); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumeratePrivilegesAccountRequest structure represents the LsarEnumeratePrivilegesAccount operation request
type EnumeratePrivilegesAccountRequest struct {
	// AccountHandle: An open account object handle obtained from either LsarCreateAccount
	// (section 3.1.4.5.1) or LsarOpenAccount (section 3.1.4.5.3).
	Account *Handle `idl:"name:AccountHandle" json:"account"`
}

func (o *EnumeratePrivilegesAccountRequest) xxx_ToOp(ctx context.Context) *xxx_EnumeratePrivilegesAccountOperation {
	if o == nil {
		return &xxx_EnumeratePrivilegesAccountOperation{}
	}
	return &xxx_EnumeratePrivilegesAccountOperation{
		Account: o.Account,
	}
}

func (o *EnumeratePrivilegesAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumeratePrivilegesAccountOperation) {
	if o == nil {
		return
	}
	o.Account = op.Account
}
func (o *EnumeratePrivilegesAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumeratePrivilegesAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumeratePrivilegesAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumeratePrivilegesAccountResponse structure represents the LsarEnumeratePrivilegesAccount operation response
type EnumeratePrivilegesAccountResponse struct {
	// Privileges: Used to return a list of privileges granted to the account.
	Privileges *PrivilegeSet `idl:"name:Privileges" json:"privileges"`
	// Return: The LsarEnumeratePrivilegesAccount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumeratePrivilegesAccountResponse) xxx_ToOp(ctx context.Context) *xxx_EnumeratePrivilegesAccountOperation {
	if o == nil {
		return &xxx_EnumeratePrivilegesAccountOperation{}
	}
	return &xxx_EnumeratePrivilegesAccountOperation{
		Privileges: o.Privileges,
		Return:     o.Return,
	}
}

func (o *EnumeratePrivilegesAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumeratePrivilegesAccountOperation) {
	if o == nil {
		return
	}
	o.Privileges = op.Privileges
	o.Return = op.Return
}
func (o *EnumeratePrivilegesAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumeratePrivilegesAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumeratePrivilegesAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddPrivilegesToAccountOperation structure represents the LsarAddPrivilegesToAccount operation
type xxx_AddPrivilegesToAccountOperation struct {
	Account    *Handle       `idl:"name:AccountHandle" json:"account"`
	Privileges *PrivilegeSet `idl:"name:Privileges" json:"privileges"`
	Return     int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_AddPrivilegesToAccountOperation) OpNum() int { return 19 }

func (o *xxx_AddPrivilegesToAccountOperation) OpName() string {
	return "/lsarpc/v0/LsarAddPrivilegesToAccount"
}

func (o *xxx_AddPrivilegesToAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddPrivilegesToAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AccountHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account != nil {
			if err := o.Account.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Privileges {in} (1:{alias=PLSAPR_PRIVILEGE_SET}*(1))(2:{alias=LSAPR_PRIVILEGE_SET}(struct))
	{
		if o.Privileges != nil {
			if err := o.Privileges.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PrivilegeSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_AddPrivilegesToAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AccountHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account == nil {
			o.Account = &Handle{}
		}
		if err := o.Account.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Privileges {in} (1:{alias=PLSAPR_PRIVILEGE_SET,pointer=ref}*(1))(2:{alias=LSAPR_PRIVILEGE_SET}(struct))
	{
		if o.Privileges == nil {
			o.Privileges = &PrivilegeSet{}
		}
		if err := o.Privileges.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddPrivilegesToAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddPrivilegesToAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddPrivilegesToAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddPrivilegesToAccountRequest structure represents the LsarAddPrivilegesToAccount operation request
type AddPrivilegesToAccountRequest struct {
	// AccountHandle: An open account object handle obtained from either LsarCreateAccount
	// (section 3.1.4.5.1) or LsarOpenAccount (section 3.1.4.5.3).
	Account *Handle `idl:"name:AccountHandle" json:"account"`
	// Privileges: Contains a list of privileges to add to the account.
	Privileges *PrivilegeSet `idl:"name:Privileges" json:"privileges"`
}

func (o *AddPrivilegesToAccountRequest) xxx_ToOp(ctx context.Context) *xxx_AddPrivilegesToAccountOperation {
	if o == nil {
		return &xxx_AddPrivilegesToAccountOperation{}
	}
	return &xxx_AddPrivilegesToAccountOperation{
		Account:    o.Account,
		Privileges: o.Privileges,
	}
}

func (o *AddPrivilegesToAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_AddPrivilegesToAccountOperation) {
	if o == nil {
		return
	}
	o.Account = op.Account
	o.Privileges = op.Privileges
}
func (o *AddPrivilegesToAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddPrivilegesToAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddPrivilegesToAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddPrivilegesToAccountResponse structure represents the LsarAddPrivilegesToAccount operation response
type AddPrivilegesToAccountResponse struct {
	// Return: The LsarAddPrivilegesToAccount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddPrivilegesToAccountResponse) xxx_ToOp(ctx context.Context) *xxx_AddPrivilegesToAccountOperation {
	if o == nil {
		return &xxx_AddPrivilegesToAccountOperation{}
	}
	return &xxx_AddPrivilegesToAccountOperation{
		Return: o.Return,
	}
}

func (o *AddPrivilegesToAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_AddPrivilegesToAccountOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddPrivilegesToAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddPrivilegesToAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddPrivilegesToAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemovePrivilegesFromAccountOperation structure represents the LsarRemovePrivilegesFromAccount operation
type xxx_RemovePrivilegesFromAccountOperation struct {
	Account       *Handle       `idl:"name:AccountHandle" json:"account"`
	AllPrivileges uint8         `idl:"name:AllPrivileges" json:"all_privileges"`
	Privileges    *PrivilegeSet `idl:"name:Privileges;pointer:unique" json:"privileges"`
	Return        int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_RemovePrivilegesFromAccountOperation) OpNum() int { return 20 }

func (o *xxx_RemovePrivilegesFromAccountOperation) OpName() string {
	return "/lsarpc/v0/LsarRemovePrivilegesFromAccount"
}

func (o *xxx_RemovePrivilegesFromAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePrivilegesFromAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AccountHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account != nil {
			if err := o.Account.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// AllPrivileges {in} (1:(uchar))
	{
		if err := w.WriteData(o.AllPrivileges); err != nil {
			return err
		}
	}
	// Privileges {in} (1:{pointer=unique, alias=PLSAPR_PRIVILEGE_SET}*(1))(2:{alias=LSAPR_PRIVILEGE_SET}(struct))
	{
		if o.Privileges != nil {
			_ptr_Privileges := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Privileges != nil {
					if err := o.Privileges.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PrivilegeSet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Privileges, _ptr_Privileges); err != nil {
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

func (o *xxx_RemovePrivilegesFromAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AccountHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account == nil {
			o.Account = &Handle{}
		}
		if err := o.Account.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// AllPrivileges {in} (1:(uchar))
	{
		if err := w.ReadData(&o.AllPrivileges); err != nil {
			return err
		}
	}
	// Privileges {in} (1:{pointer=unique, alias=PLSAPR_PRIVILEGE_SET}*(1))(2:{alias=LSAPR_PRIVILEGE_SET}(struct))
	{
		_ptr_Privileges := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Privileges == nil {
				o.Privileges = &PrivilegeSet{}
			}
			if err := o.Privileges.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Privileges := func(ptr interface{}) { o.Privileges = *ptr.(**PrivilegeSet) }
		if err := w.ReadPointer(&o.Privileges, _s_Privileges, _ptr_Privileges); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePrivilegesFromAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePrivilegesFromAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePrivilegesFromAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemovePrivilegesFromAccountRequest structure represents the LsarRemovePrivilegesFromAccount operation request
type RemovePrivilegesFromAccountRequest struct {
	// AccountHandle: An open account object handle obtained from either LsarCreateAccount
	// (section 3.1.4.5.1) or LsarOpenAccount (section 3.1.4.5.3).
	Account *Handle `idl:"name:AccountHandle" json:"account"`
	// AllPrivileges: If this parameter is not FALSE (0), all privileges will be stripped
	// from the account object.
	AllPrivileges uint8 `idl:"name:AllPrivileges" json:"all_privileges"`
	// Privileges: Contains a (possibly empty) list of privileges to remove from the account
	// object.
	Privileges *PrivilegeSet `idl:"name:Privileges;pointer:unique" json:"privileges"`
}

func (o *RemovePrivilegesFromAccountRequest) xxx_ToOp(ctx context.Context) *xxx_RemovePrivilegesFromAccountOperation {
	if o == nil {
		return &xxx_RemovePrivilegesFromAccountOperation{}
	}
	return &xxx_RemovePrivilegesFromAccountOperation{
		Account:       o.Account,
		AllPrivileges: o.AllPrivileges,
		Privileges:    o.Privileges,
	}
}

func (o *RemovePrivilegesFromAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_RemovePrivilegesFromAccountOperation) {
	if o == nil {
		return
	}
	o.Account = op.Account
	o.AllPrivileges = op.AllPrivileges
	o.Privileges = op.Privileges
}
func (o *RemovePrivilegesFromAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemovePrivilegesFromAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemovePrivilegesFromAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemovePrivilegesFromAccountResponse structure represents the LsarRemovePrivilegesFromAccount operation response
type RemovePrivilegesFromAccountResponse struct {
	// Return: The LsarRemovePrivilegesFromAccount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemovePrivilegesFromAccountResponse) xxx_ToOp(ctx context.Context) *xxx_RemovePrivilegesFromAccountOperation {
	if o == nil {
		return &xxx_RemovePrivilegesFromAccountOperation{}
	}
	return &xxx_RemovePrivilegesFromAccountOperation{
		Return: o.Return,
	}
}

func (o *RemovePrivilegesFromAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_RemovePrivilegesFromAccountOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemovePrivilegesFromAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemovePrivilegesFromAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemovePrivilegesFromAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSystemAccessAccountOperation structure represents the LsarGetSystemAccessAccount operation
type xxx_GetSystemAccessAccountOperation struct {
	Account      *Handle `idl:"name:AccountHandle" json:"account"`
	SystemAccess uint32  `idl:"name:SystemAccess" json:"system_access"`
	Return       int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSystemAccessAccountOperation) OpNum() int { return 23 }

func (o *xxx_GetSystemAccessAccountOperation) OpName() string {
	return "/lsarpc/v0/LsarGetSystemAccessAccount"
}

func (o *xxx_GetSystemAccessAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemAccessAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AccountHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account != nil {
			if err := o.Account.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetSystemAccessAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AccountHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account == nil {
			o.Account = &Handle{}
		}
		if err := o.Account.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemAccessAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemAccessAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SystemAccess {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.SystemAccess); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSystemAccessAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SystemAccess {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.SystemAccess); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSystemAccessAccountRequest structure represents the LsarGetSystemAccessAccount operation request
type GetSystemAccessAccountRequest struct {
	// AccountHandle: An open account object handle obtained from either LsarCreateAccount
	// (section 3.1.4.5.1) or LsarOpenAccount (section 3.1.4.5.3).
	Account *Handle `idl:"name:AccountHandle" json:"account"`
}

func (o *GetSystemAccessAccountRequest) xxx_ToOp(ctx context.Context) *xxx_GetSystemAccessAccountOperation {
	if o == nil {
		return &xxx_GetSystemAccessAccountOperation{}
	}
	return &xxx_GetSystemAccessAccountOperation{
		Account: o.Account,
	}
}

func (o *GetSystemAccessAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSystemAccessAccountOperation) {
	if o == nil {
		return
	}
	o.Account = op.Account
}
func (o *GetSystemAccessAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSystemAccessAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSystemAccessAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSystemAccessAccountResponse structure represents the LsarGetSystemAccessAccount operation response
type GetSystemAccessAccountResponse struct {
	// SystemAccess: Used to return a bitmask of access flags associated with the account.
	SystemAccess uint32 `idl:"name:SystemAccess" json:"system_access"`
	// Return: The LsarGetSystemAccessAccount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSystemAccessAccountResponse) xxx_ToOp(ctx context.Context) *xxx_GetSystemAccessAccountOperation {
	if o == nil {
		return &xxx_GetSystemAccessAccountOperation{}
	}
	return &xxx_GetSystemAccessAccountOperation{
		SystemAccess: o.SystemAccess,
		Return:       o.Return,
	}
}

func (o *GetSystemAccessAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSystemAccessAccountOperation) {
	if o == nil {
		return
	}
	o.SystemAccess = op.SystemAccess
	o.Return = op.Return
}
func (o *GetSystemAccessAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSystemAccessAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSystemAccessAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSystemAccessAccountOperation structure represents the LsarSetSystemAccessAccount operation
type xxx_SetSystemAccessAccountOperation struct {
	Account      *Handle `idl:"name:AccountHandle" json:"account"`
	SystemAccess uint32  `idl:"name:SystemAccess" json:"system_access"`
	Return       int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSystemAccessAccountOperation) OpNum() int { return 24 }

func (o *xxx_SetSystemAccessAccountOperation) OpName() string {
	return "/lsarpc/v0/LsarSetSystemAccessAccount"
}

func (o *xxx_SetSystemAccessAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSystemAccessAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AccountHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account != nil {
			if err := o.Account.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SystemAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.SystemAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSystemAccessAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AccountHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Account == nil {
			o.Account = &Handle{}
		}
		if err := o.Account.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SystemAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.SystemAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSystemAccessAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSystemAccessAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSystemAccessAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSystemAccessAccountRequest structure represents the LsarSetSystemAccessAccount operation request
type SetSystemAccessAccountRequest struct {
	// AccountHandle: An open account object handle obtained from either LsarCreateAccount
	// (section 3.1.4.5.1) or LsarOpenAccount (section 3.1.4.5.3).
	Account *Handle `idl:"name:AccountHandle" json:"account"`
	// SystemAccess: A bitmask containing the account flags to be set on the account.
	SystemAccess uint32 `idl:"name:SystemAccess" json:"system_access"`
}

func (o *SetSystemAccessAccountRequest) xxx_ToOp(ctx context.Context) *xxx_SetSystemAccessAccountOperation {
	if o == nil {
		return &xxx_SetSystemAccessAccountOperation{}
	}
	return &xxx_SetSystemAccessAccountOperation{
		Account:      o.Account,
		SystemAccess: o.SystemAccess,
	}
}

func (o *SetSystemAccessAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSystemAccessAccountOperation) {
	if o == nil {
		return
	}
	o.Account = op.Account
	o.SystemAccess = op.SystemAccess
}
func (o *SetSystemAccessAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSystemAccessAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSystemAccessAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSystemAccessAccountResponse structure represents the LsarSetSystemAccessAccount operation response
type SetSystemAccessAccountResponse struct {
	// Return: The LsarSetSystemAccessAccount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSystemAccessAccountResponse) xxx_ToOp(ctx context.Context) *xxx_SetSystemAccessAccountOperation {
	if o == nil {
		return &xxx_SetSystemAccessAccountOperation{}
	}
	return &xxx_SetSystemAccessAccountOperation{
		Return: o.Return,
	}
}

func (o *SetSystemAccessAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSystemAccessAccountOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetSystemAccessAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSystemAccessAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSystemAccessAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenTrustedDomainOperation structure represents the LsarOpenTrustedDomain operation
type xxx_OpenTrustedDomainOperation struct {
	Policy           *Handle   `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainSID *dtyp.SID `idl:"name:TrustedDomainSid" json:"trusted_domain_sid"`
	DesiredAccess    uint32    `idl:"name:DesiredAccess" json:"desired_access"`
	TrustedDomain    *Handle   `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	Return           int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenTrustedDomainOperation) OpNum() int { return 25 }

func (o *xxx_OpenTrustedDomainOperation) OpName() string { return "/lsarpc/v0/LsarOpenTrustedDomain" }

func (o *xxx_OpenTrustedDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenTrustedDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainSid {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.TrustedDomainSID != nil {
			if err := o.TrustedDomainSID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenTrustedDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainSid {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.TrustedDomainSID == nil {
			o.TrustedDomainSID = &dtyp.SID{}
		}
		if err := o.TrustedDomainSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenTrustedDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenTrustedDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TrustedDomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain != nil {
			if err := o.TrustedDomain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenTrustedDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TrustedDomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain == nil {
			o.TrustedDomain = &Handle{}
		}
		if err := o.TrustedDomain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenTrustedDomainRequest structure represents the LsarOpenTrustedDomain operation request
type OpenTrustedDomainRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainSid: A security identifier of the trusted domain that is being opened.
	TrustedDomainSID *dtyp.SID `idl:"name:TrustedDomainSid" json:"trusted_domain_sid"`
	// DesiredAccess: A bitmask of access rights to open the object with.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *OpenTrustedDomainRequest) xxx_ToOp(ctx context.Context) *xxx_OpenTrustedDomainOperation {
	if o == nil {
		return &xxx_OpenTrustedDomainOperation{}
	}
	return &xxx_OpenTrustedDomainOperation{
		Policy:           o.Policy,
		TrustedDomainSID: o.TrustedDomainSID,
		DesiredAccess:    o.DesiredAccess,
	}
}

func (o *OpenTrustedDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenTrustedDomainOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainSID = op.TrustedDomainSID
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenTrustedDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenTrustedDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenTrustedDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenTrustedDomainResponse structure represents the LsarOpenTrustedDomain operation response
type OpenTrustedDomainResponse struct {
	// TrustedDomainHandle: Used to return the trusted domain object handle.
	TrustedDomain *Handle `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	// Return: The LsarOpenTrustedDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenTrustedDomainResponse) xxx_ToOp(ctx context.Context) *xxx_OpenTrustedDomainOperation {
	if o == nil {
		return &xxx_OpenTrustedDomainOperation{}
	}
	return &xxx_OpenTrustedDomainOperation{
		TrustedDomain: o.TrustedDomain,
		Return:        o.Return,
	}
}

func (o *OpenTrustedDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenTrustedDomainOperation) {
	if o == nil {
		return
	}
	o.TrustedDomain = op.TrustedDomain
	o.Return = op.Return
}
func (o *OpenTrustedDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenTrustedDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenTrustedDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInfoTrustedDomainOperation structure represents the LsarQueryInfoTrustedDomain operation
type xxx_QueryInfoTrustedDomainOperation struct {
	TrustedDomain            *Handle                 `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	InformationClass         TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
	TrustedDomainInformation *TrustedDomainInfo      `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
	Return                   int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInfoTrustedDomainOperation) OpNum() int { return 26 }

func (o *xxx_QueryInfoTrustedDomainOperation) OpName() string {
	return "/lsarpc/v0/LsarQueryInfoTrustedDomain"
}

func (o *xxx_QueryInfoTrustedDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInfoTrustedDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// TrustedDomainHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain != nil {
			if err := o.TrustedDomain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInfoTrustedDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// TrustedDomainHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain == nil {
			o.TrustedDomain = &Handle{}
		}
		if err := o.TrustedDomain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInfoTrustedDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInfoTrustedDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TrustedDomainInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO}*(1))(3:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		if o.TrustedDomainInformation != nil {
			_ptr_TrustedDomainInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swTrustedDomainInformation := uint16(o.InformationClass)
				if o.TrustedDomainInformation != nil {
					if err := o.TrustedDomainInformation.MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
						return err
					}
				} else {
					if err := (&TrustedDomainInfo{}).MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TrustedDomainInformation, _ptr_TrustedDomainInformation); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInfoTrustedDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TrustedDomainInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO,pointer=ref}*(1))(3:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		_ptr_TrustedDomainInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TrustedDomainInformation == nil {
				o.TrustedDomainInformation = &TrustedDomainInfo{}
			}
			_swTrustedDomainInformation := uint16(o.InformationClass)
			if err := o.TrustedDomainInformation.UnmarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
				return err
			}
			return nil
		})
		_s_TrustedDomainInformation := func(ptr interface{}) { o.TrustedDomainInformation = *ptr.(**TrustedDomainInfo) }
		if err := w.ReadPointer(&o.TrustedDomainInformation, _s_TrustedDomainInformation, _ptr_TrustedDomainInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryInfoTrustedDomainRequest structure represents the LsarQueryInfoTrustedDomain operation request
type QueryInfoTrustedDomainRequest struct {
	// TrustedDomainHandle: An open trusted domain object handle.
	TrustedDomain *Handle `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	// InformationClass: One of the TRUSTED_INFORMATION_CLASS values indicating the type
	// of information the caller is interested in.
	InformationClass TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
}

func (o *QueryInfoTrustedDomainRequest) xxx_ToOp(ctx context.Context) *xxx_QueryInfoTrustedDomainOperation {
	if o == nil {
		return &xxx_QueryInfoTrustedDomainOperation{}
	}
	return &xxx_QueryInfoTrustedDomainOperation{
		TrustedDomain:    o.TrustedDomain,
		InformationClass: o.InformationClass,
	}
}

func (o *QueryInfoTrustedDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryInfoTrustedDomainOperation) {
	if o == nil {
		return
	}
	o.TrustedDomain = op.TrustedDomain
	o.InformationClass = op.InformationClass
}
func (o *QueryInfoTrustedDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInfoTrustedDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInfoTrustedDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInfoTrustedDomainResponse structure represents the LsarQueryInfoTrustedDomain operation response
type QueryInfoTrustedDomainResponse struct {
	// TrustedDomainInformation: Used to return requested information about the trusted
	// domain object.
	TrustedDomainInformation *TrustedDomainInfo `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
	// Return: The LsarQueryInfoTrustedDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryInfoTrustedDomainResponse) xxx_ToOp(ctx context.Context) *xxx_QueryInfoTrustedDomainOperation {
	if o == nil {
		return &xxx_QueryInfoTrustedDomainOperation{}
	}
	return &xxx_QueryInfoTrustedDomainOperation{
		TrustedDomainInformation: o.TrustedDomainInformation,
		Return:                   o.Return,
	}
}

func (o *QueryInfoTrustedDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryInfoTrustedDomainOperation) {
	if o == nil {
		return
	}
	o.TrustedDomainInformation = op.TrustedDomainInformation
	o.Return = op.Return
}
func (o *QueryInfoTrustedDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInfoTrustedDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInfoTrustedDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInformationTrustedDomainOperation structure represents the LsarSetInformationTrustedDomain operation
type xxx_SetInformationTrustedDomainOperation struct {
	TrustedDomain            *Handle                 `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	InformationClass         TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
	TrustedDomainInformation *TrustedDomainInfo      `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
	Return                   int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInformationTrustedDomainOperation) OpNum() int { return 27 }

func (o *xxx_SetInformationTrustedDomainOperation) OpName() string {
	return "/lsarpc/v0/LsarSetInformationTrustedDomain"
}

func (o *xxx_SetInformationTrustedDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationTrustedDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// TrustedDomainHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain != nil {
			if err := o.TrustedDomain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	// TrustedDomainInformation {in} (1:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO}*(1))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		_swTrustedDomainInformation := uint16(o.InformationClass)
		if o.TrustedDomainInformation != nil {
			if err := o.TrustedDomainInformation.MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo{}).MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationTrustedDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// TrustedDomainHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain == nil {
			o.TrustedDomain = &Handle{}
		}
		if err := o.TrustedDomain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	// TrustedDomainInformation {in} (1:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO,pointer=ref}*(1))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		if o.TrustedDomainInformation == nil {
			o.TrustedDomainInformation = &TrustedDomainInfo{}
		}
		_swTrustedDomainInformation := uint16(o.InformationClass)
		if err := o.TrustedDomainInformation.UnmarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationTrustedDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationTrustedDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationTrustedDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInformationTrustedDomainRequest structure represents the LsarSetInformationTrustedDomain operation request
type SetInformationTrustedDomainRequest struct {
	// TrustedDomainHandle: A handle to a trusted domain object.
	TrustedDomain *Handle `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	// InformationClass: A value indicating the type of information requested by the caller.
	InformationClass TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
	// TrustedDomainInformation: Used to supply the information to be set.
	TrustedDomainInformation *TrustedDomainInfo `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
}

func (o *SetInformationTrustedDomainRequest) xxx_ToOp(ctx context.Context) *xxx_SetInformationTrustedDomainOperation {
	if o == nil {
		return &xxx_SetInformationTrustedDomainOperation{}
	}
	return &xxx_SetInformationTrustedDomainOperation{
		TrustedDomain:            o.TrustedDomain,
		InformationClass:         o.InformationClass,
		TrustedDomainInformation: o.TrustedDomainInformation,
	}
}

func (o *SetInformationTrustedDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInformationTrustedDomainOperation) {
	if o == nil {
		return
	}
	o.TrustedDomain = op.TrustedDomain
	o.InformationClass = op.InformationClass
	o.TrustedDomainInformation = op.TrustedDomainInformation
}
func (o *SetInformationTrustedDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetInformationTrustedDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationTrustedDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInformationTrustedDomainResponse structure represents the LsarSetInformationTrustedDomain operation response
type SetInformationTrustedDomainResponse struct {
	// Return: The LsarSetInformationTrustedDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInformationTrustedDomainResponse) xxx_ToOp(ctx context.Context) *xxx_SetInformationTrustedDomainOperation {
	if o == nil {
		return &xxx_SetInformationTrustedDomainOperation{}
	}
	return &xxx_SetInformationTrustedDomainOperation{
		Return: o.Return,
	}
}

func (o *SetInformationTrustedDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInformationTrustedDomainOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetInformationTrustedDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetInformationTrustedDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationTrustedDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenSecretOperation structure represents the LsarOpenSecret operation
type xxx_OpenSecretOperation struct {
	Policy        *Handle             `idl:"name:PolicyHandle" json:"policy"`
	SecretName    *dtyp.UnicodeString `idl:"name:SecretName" json:"secret_name"`
	DesiredAccess uint32              `idl:"name:DesiredAccess" json:"desired_access"`
	Secret        *Handle             `idl:"name:SecretHandle" json:"secret"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenSecretOperation) OpNum() int { return 28 }

func (o *xxx_OpenSecretOperation) OpName() string { return "/lsarpc/v0/LsarOpenSecret" }

func (o *xxx_OpenSecretOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenSecretOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecretName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.SecretName != nil {
			if err := o.SecretName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenSecretOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecretName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.SecretName == nil {
			o.SecretName = &dtyp.UnicodeString{}
		}
		if err := o.SecretName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenSecretOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenSecretOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SecretHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Secret != nil {
			if err := o.Secret.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenSecretOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SecretHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Secret == nil {
			o.Secret = &Handle{}
		}
		if err := o.Secret.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenSecretRequest structure represents the LsarOpenSecret operation request
type OpenSecretRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// SecretName: The name of the secret object to open.
	SecretName *dtyp.UnicodeString `idl:"name:SecretName" json:"secret_name"`
	// DesiredAccess: The requested type of access.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *OpenSecretRequest) xxx_ToOp(ctx context.Context) *xxx_OpenSecretOperation {
	if o == nil {
		return &xxx_OpenSecretOperation{}
	}
	return &xxx_OpenSecretOperation{
		Policy:        o.Policy,
		SecretName:    o.SecretName,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *OpenSecretRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenSecretOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.SecretName = op.SecretName
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenSecretRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenSecretRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenSecretOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenSecretResponse structure represents the LsarOpenSecret operation response
type OpenSecretResponse struct {
	// SecretHandle: Used to return the handle to the opened secret object.
	Secret *Handle `idl:"name:SecretHandle" json:"secret"`
	// Return: The LsarOpenSecret return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenSecretResponse) xxx_ToOp(ctx context.Context) *xxx_OpenSecretOperation {
	if o == nil {
		return &xxx_OpenSecretOperation{}
	}
	return &xxx_OpenSecretOperation{
		Secret: o.Secret,
		Return: o.Return,
	}
}

func (o *OpenSecretResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenSecretOperation) {
	if o == nil {
		return
	}
	o.Secret = op.Secret
	o.Return = op.Return
}
func (o *OpenSecretResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenSecretResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenSecretOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSecretOperation structure represents the LsarSetSecret operation
type xxx_SetSecretOperation struct {
	Secret                *Handle        `idl:"name:SecretHandle" json:"secret"`
	EncryptedCurrentValue *CRCipherValue `idl:"name:EncryptedCurrentValue;pointer:unique" json:"encrypted_current_value"`
	EncryptedOldValue     *CRCipherValue `idl:"name:EncryptedOldValue;pointer:unique" json:"encrypted_old_value"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSecretOperation) OpNum() int { return 29 }

func (o *xxx_SetSecretOperation) OpName() string { return "/lsarpc/v0/LsarSetSecret" }

func (o *xxx_SetSecretOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecretOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SecretHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Secret != nil {
			if err := o.Secret.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// EncryptedCurrentValue {in} (1:{pointer=unique, alias=PLSAPR_CR_CIPHER_VALUE}*(1))(2:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		if o.EncryptedCurrentValue != nil {
			_ptr_EncryptedCurrentValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedCurrentValue != nil {
					if err := o.EncryptedCurrentValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedCurrentValue, _ptr_EncryptedCurrentValue); err != nil {
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
	// EncryptedOldValue {in} (1:{pointer=unique, alias=PLSAPR_CR_CIPHER_VALUE}*(1))(2:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		if o.EncryptedOldValue != nil {
			_ptr_EncryptedOldValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedOldValue != nil {
					if err := o.EncryptedOldValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedOldValue, _ptr_EncryptedOldValue); err != nil {
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

func (o *xxx_SetSecretOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SecretHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Secret == nil {
			o.Secret = &Handle{}
		}
		if err := o.Secret.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EncryptedCurrentValue {in} (1:{pointer=unique, alias=PLSAPR_CR_CIPHER_VALUE}*(1))(2:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		_ptr_EncryptedCurrentValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EncryptedCurrentValue == nil {
				o.EncryptedCurrentValue = &CRCipherValue{}
			}
			if err := o.EncryptedCurrentValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedCurrentValue := func(ptr interface{}) { o.EncryptedCurrentValue = *ptr.(**CRCipherValue) }
		if err := w.ReadPointer(&o.EncryptedCurrentValue, _s_EncryptedCurrentValue, _ptr_EncryptedCurrentValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// EncryptedOldValue {in} (1:{pointer=unique, alias=PLSAPR_CR_CIPHER_VALUE}*(1))(2:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		_ptr_EncryptedOldValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EncryptedOldValue == nil {
				o.EncryptedOldValue = &CRCipherValue{}
			}
			if err := o.EncryptedOldValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedOldValue := func(ptr interface{}) { o.EncryptedOldValue = *ptr.(**CRCipherValue) }
		if err := w.ReadPointer(&o.EncryptedOldValue, _s_EncryptedOldValue, _ptr_EncryptedOldValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecretOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecretOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecretOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSecretRequest structure represents the LsarSetSecret operation request
type SetSecretRequest struct {
	// SecretHandle: An open secret object handle.
	Secret *Handle `idl:"name:SecretHandle" json:"secret"`
	// EncryptedCurrentValue: A binary large object (BLOB) representing a new encrypted
	// cipher value. It is valid for this parameter to be NULL, in which case the value
	// is deleted from the server's policy database.
	EncryptedCurrentValue *CRCipherValue `idl:"name:EncryptedCurrentValue;pointer:unique" json:"encrypted_current_value"`
	// EncryptedOldValue: A BLOB representing the encrypted old value. It is valid for this
	// parameter to be NULL, in which case the current value in the policy database is copied.
	EncryptedOldValue *CRCipherValue `idl:"name:EncryptedOldValue;pointer:unique" json:"encrypted_old_value"`
}

func (o *SetSecretRequest) xxx_ToOp(ctx context.Context) *xxx_SetSecretOperation {
	if o == nil {
		return &xxx_SetSecretOperation{}
	}
	return &xxx_SetSecretOperation{
		Secret:                o.Secret,
		EncryptedCurrentValue: o.EncryptedCurrentValue,
		EncryptedOldValue:     o.EncryptedOldValue,
	}
}

func (o *SetSecretRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSecretOperation) {
	if o == nil {
		return
	}
	o.Secret = op.Secret
	o.EncryptedCurrentValue = op.EncryptedCurrentValue
	o.EncryptedOldValue = op.EncryptedOldValue
}
func (o *SetSecretRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSecretRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecretOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSecretResponse structure represents the LsarSetSecret operation response
type SetSecretResponse struct {
	// Return: The LsarSetSecret return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSecretResponse) xxx_ToOp(ctx context.Context) *xxx_SetSecretOperation {
	if o == nil {
		return &xxx_SetSecretOperation{}
	}
	return &xxx_SetSecretOperation{
		Return: o.Return,
	}
}

func (o *SetSecretResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSecretOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetSecretResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSecretResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecretOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QuerySecretOperation structure represents the LsarQuerySecret operation
type xxx_QuerySecretOperation struct {
	Secret                *Handle            `idl:"name:SecretHandle" json:"secret"`
	EncryptedCurrentValue *CRCipherValue     `idl:"name:EncryptedCurrentValue;pointer:unique" json:"encrypted_current_value"`
	CurrentValueSetTime   *dtyp.LargeInteger `idl:"name:CurrentValueSetTime;pointer:unique" json:"current_value_set_time"`
	EncryptedOldValue     *CRCipherValue     `idl:"name:EncryptedOldValue;pointer:unique" json:"encrypted_old_value"`
	OldValueSetTime       *dtyp.LargeInteger `idl:"name:OldValueSetTime;pointer:unique" json:"old_value_set_time"`
	Return                int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_QuerySecretOperation) OpNum() int { return 30 }

func (o *xxx_QuerySecretOperation) OpName() string { return "/lsarpc/v0/LsarQuerySecret" }

func (o *xxx_QuerySecretOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySecretOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SecretHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Secret != nil {
			if err := o.Secret.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// EncryptedCurrentValue {in, out} (1:{pointer=unique}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		if o.EncryptedCurrentValue != nil {
			_ptr_EncryptedCurrentValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedCurrentValue != nil {
					_ptr_EncryptedCurrentValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.EncryptedCurrentValue != nil {
							if err := o.EncryptedCurrentValue.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.EncryptedCurrentValue, _ptr_EncryptedCurrentValue); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedCurrentValue, _ptr_EncryptedCurrentValue); err != nil {
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
	// CurrentValueSetTime {in, out} (1:{pointer=unique, alias=PLARGE_INTEGER}*(1))(2:{alias=LARGE_INTEGER}(struct))
	{
		if o.CurrentValueSetTime != nil {
			_ptr_CurrentValueSetTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CurrentValueSetTime != nil {
					if err := o.CurrentValueSetTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CurrentValueSetTime, _ptr_CurrentValueSetTime); err != nil {
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
	// EncryptedOldValue {in, out} (1:{pointer=unique}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		if o.EncryptedOldValue != nil {
			_ptr_EncryptedOldValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedOldValue != nil {
					_ptr_EncryptedOldValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.EncryptedOldValue != nil {
							if err := o.EncryptedOldValue.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.EncryptedOldValue, _ptr_EncryptedOldValue); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedOldValue, _ptr_EncryptedOldValue); err != nil {
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
	// OldValueSetTime {in, out} (1:{pointer=unique, alias=PLARGE_INTEGER}*(1))(2:{alias=LARGE_INTEGER}(struct))
	{
		if o.OldValueSetTime != nil {
			_ptr_OldValueSetTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldValueSetTime != nil {
					if err := o.OldValueSetTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldValueSetTime, _ptr_OldValueSetTime); err != nil {
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

func (o *xxx_QuerySecretOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SecretHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Secret == nil {
			o.Secret = &Handle{}
		}
		if err := o.Secret.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EncryptedCurrentValue {in, out} (1:{pointer=unique}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE,pointer=ref}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		_ptr_EncryptedCurrentValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_EncryptedCurrentValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.EncryptedCurrentValue == nil {
					o.EncryptedCurrentValue = &CRCipherValue{}
				}
				if err := o.EncryptedCurrentValue.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_EncryptedCurrentValue := func(ptr interface{}) { o.EncryptedCurrentValue = *ptr.(**CRCipherValue) }
			if err := w.ReadPointer(&o.EncryptedCurrentValue, _s_EncryptedCurrentValue, _ptr_EncryptedCurrentValue); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedCurrentValue := func(ptr interface{}) { o.EncryptedCurrentValue = *ptr.(**CRCipherValue) }
		if err := w.ReadPointer(&o.EncryptedCurrentValue, _s_EncryptedCurrentValue, _ptr_EncryptedCurrentValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CurrentValueSetTime {in, out} (1:{pointer=unique, alias=PLARGE_INTEGER}*(1))(2:{alias=LARGE_INTEGER}(struct))
	{
		_ptr_CurrentValueSetTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CurrentValueSetTime == nil {
				o.CurrentValueSetTime = &dtyp.LargeInteger{}
			}
			if err := o.CurrentValueSetTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_CurrentValueSetTime := func(ptr interface{}) { o.CurrentValueSetTime = *ptr.(**dtyp.LargeInteger) }
		if err := w.ReadPointer(&o.CurrentValueSetTime, _s_CurrentValueSetTime, _ptr_CurrentValueSetTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// EncryptedOldValue {in, out} (1:{pointer=unique}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE,pointer=ref}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		_ptr_EncryptedOldValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_EncryptedOldValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.EncryptedOldValue == nil {
					o.EncryptedOldValue = &CRCipherValue{}
				}
				if err := o.EncryptedOldValue.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_EncryptedOldValue := func(ptr interface{}) { o.EncryptedOldValue = *ptr.(**CRCipherValue) }
			if err := w.ReadPointer(&o.EncryptedOldValue, _s_EncryptedOldValue, _ptr_EncryptedOldValue); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedOldValue := func(ptr interface{}) { o.EncryptedOldValue = *ptr.(**CRCipherValue) }
		if err := w.ReadPointer(&o.EncryptedOldValue, _s_EncryptedOldValue, _ptr_EncryptedOldValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OldValueSetTime {in, out} (1:{pointer=unique, alias=PLARGE_INTEGER}*(1))(2:{alias=LARGE_INTEGER}(struct))
	{
		_ptr_OldValueSetTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldValueSetTime == nil {
				o.OldValueSetTime = &dtyp.LargeInteger{}
			}
			if err := o.OldValueSetTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OldValueSetTime := func(ptr interface{}) { o.OldValueSetTime = *ptr.(**dtyp.LargeInteger) }
		if err := w.ReadPointer(&o.OldValueSetTime, _s_OldValueSetTime, _ptr_OldValueSetTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySecretOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySecretOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EncryptedCurrentValue {in, out} (1:{pointer=unique}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		if o.EncryptedCurrentValue != nil {
			_ptr_EncryptedCurrentValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedCurrentValue != nil {
					_ptr_EncryptedCurrentValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.EncryptedCurrentValue != nil {
							if err := o.EncryptedCurrentValue.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.EncryptedCurrentValue, _ptr_EncryptedCurrentValue); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedCurrentValue, _ptr_EncryptedCurrentValue); err != nil {
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
	// CurrentValueSetTime {in, out} (1:{pointer=unique, alias=PLARGE_INTEGER}*(1))(2:{alias=LARGE_INTEGER}(struct))
	{
		if o.CurrentValueSetTime != nil {
			_ptr_CurrentValueSetTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CurrentValueSetTime != nil {
					if err := o.CurrentValueSetTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CurrentValueSetTime, _ptr_CurrentValueSetTime); err != nil {
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
	// EncryptedOldValue {in, out} (1:{pointer=unique}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		if o.EncryptedOldValue != nil {
			_ptr_EncryptedOldValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedOldValue != nil {
					_ptr_EncryptedOldValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.EncryptedOldValue != nil {
							if err := o.EncryptedOldValue.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.EncryptedOldValue, _ptr_EncryptedOldValue); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedOldValue, _ptr_EncryptedOldValue); err != nil {
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
	// OldValueSetTime {in, out} (1:{pointer=unique, alias=PLARGE_INTEGER}*(1))(2:{alias=LARGE_INTEGER}(struct))
	{
		if o.OldValueSetTime != nil {
			_ptr_OldValueSetTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldValueSetTime != nil {
					if err := o.OldValueSetTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldValueSetTime, _ptr_OldValueSetTime); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySecretOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EncryptedCurrentValue {in, out} (1:{pointer=unique}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE,pointer=ref}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		_ptr_EncryptedCurrentValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_EncryptedCurrentValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.EncryptedCurrentValue == nil {
					o.EncryptedCurrentValue = &CRCipherValue{}
				}
				if err := o.EncryptedCurrentValue.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_EncryptedCurrentValue := func(ptr interface{}) { o.EncryptedCurrentValue = *ptr.(**CRCipherValue) }
			if err := w.ReadPointer(&o.EncryptedCurrentValue, _s_EncryptedCurrentValue, _ptr_EncryptedCurrentValue); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedCurrentValue := func(ptr interface{}) { o.EncryptedCurrentValue = *ptr.(**CRCipherValue) }
		if err := w.ReadPointer(&o.EncryptedCurrentValue, _s_EncryptedCurrentValue, _ptr_EncryptedCurrentValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CurrentValueSetTime {in, out} (1:{pointer=unique, alias=PLARGE_INTEGER}*(1))(2:{alias=LARGE_INTEGER}(struct))
	{
		_ptr_CurrentValueSetTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CurrentValueSetTime == nil {
				o.CurrentValueSetTime = &dtyp.LargeInteger{}
			}
			if err := o.CurrentValueSetTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_CurrentValueSetTime := func(ptr interface{}) { o.CurrentValueSetTime = *ptr.(**dtyp.LargeInteger) }
		if err := w.ReadPointer(&o.CurrentValueSetTime, _s_CurrentValueSetTime, _ptr_CurrentValueSetTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// EncryptedOldValue {in, out} (1:{pointer=unique}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE,pointer=ref}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		_ptr_EncryptedOldValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_EncryptedOldValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.EncryptedOldValue == nil {
					o.EncryptedOldValue = &CRCipherValue{}
				}
				if err := o.EncryptedOldValue.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_EncryptedOldValue := func(ptr interface{}) { o.EncryptedOldValue = *ptr.(**CRCipherValue) }
			if err := w.ReadPointer(&o.EncryptedOldValue, _s_EncryptedOldValue, _ptr_EncryptedOldValue); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedOldValue := func(ptr interface{}) { o.EncryptedOldValue = *ptr.(**CRCipherValue) }
		if err := w.ReadPointer(&o.EncryptedOldValue, _s_EncryptedOldValue, _ptr_EncryptedOldValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OldValueSetTime {in, out} (1:{pointer=unique, alias=PLARGE_INTEGER}*(1))(2:{alias=LARGE_INTEGER}(struct))
	{
		_ptr_OldValueSetTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldValueSetTime == nil {
				o.OldValueSetTime = &dtyp.LargeInteger{}
			}
			if err := o.OldValueSetTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OldValueSetTime := func(ptr interface{}) { o.OldValueSetTime = *ptr.(**dtyp.LargeInteger) }
		if err := w.ReadPointer(&o.OldValueSetTime, _s_OldValueSetTime, _ptr_OldValueSetTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QuerySecretRequest structure represents the LsarQuerySecret operation request
type QuerySecretRequest struct {
	// SecretHandle: An open secret object handle.
	Secret *Handle `idl:"name:SecretHandle" json:"secret"`
	// EncryptedCurrentValue: Used to return the encrypted current value of the secret object.
	EncryptedCurrentValue *CRCipherValue `idl:"name:EncryptedCurrentValue;pointer:unique" json:"encrypted_current_value"`
	// CurrentValueSetTime: Used to return the time when the current value was set.
	CurrentValueSetTime *dtyp.LargeInteger `idl:"name:CurrentValueSetTime;pointer:unique" json:"current_value_set_time"`
	// EncryptedOldValue: A BLOB representing the encrypted old value. It is valid for this
	// parameter to be NULL, in which case the current value in the policy database is copied.
	EncryptedOldValue *CRCipherValue `idl:"name:EncryptedOldValue;pointer:unique" json:"encrypted_old_value"`
	// OldValueSetTime: The time corresponding to the instant that the old value was last
	// changed.
	OldValueSetTime *dtyp.LargeInteger `idl:"name:OldValueSetTime;pointer:unique" json:"old_value_set_time"`
}

func (o *QuerySecretRequest) xxx_ToOp(ctx context.Context) *xxx_QuerySecretOperation {
	if o == nil {
		return &xxx_QuerySecretOperation{}
	}
	return &xxx_QuerySecretOperation{
		Secret:                o.Secret,
		EncryptedCurrentValue: o.EncryptedCurrentValue,
		CurrentValueSetTime:   o.CurrentValueSetTime,
		EncryptedOldValue:     o.EncryptedOldValue,
		OldValueSetTime:       o.OldValueSetTime,
	}
}

func (o *QuerySecretRequest) xxx_FromOp(ctx context.Context, op *xxx_QuerySecretOperation) {
	if o == nil {
		return
	}
	o.Secret = op.Secret
	o.EncryptedCurrentValue = op.EncryptedCurrentValue
	o.CurrentValueSetTime = op.CurrentValueSetTime
	o.EncryptedOldValue = op.EncryptedOldValue
	o.OldValueSetTime = op.OldValueSetTime
}
func (o *QuerySecretRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QuerySecretRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySecretOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QuerySecretResponse structure represents the LsarQuerySecret operation response
type QuerySecretResponse struct {
	// EncryptedCurrentValue: Used to return the encrypted current value of the secret object.
	EncryptedCurrentValue *CRCipherValue `idl:"name:EncryptedCurrentValue;pointer:unique" json:"encrypted_current_value"`
	// CurrentValueSetTime: Used to return the time when the current value was set.
	CurrentValueSetTime *dtyp.LargeInteger `idl:"name:CurrentValueSetTime;pointer:unique" json:"current_value_set_time"`
	// EncryptedOldValue: A BLOB representing the encrypted old value. It is valid for this
	// parameter to be NULL, in which case the current value in the policy database is copied.
	EncryptedOldValue *CRCipherValue `idl:"name:EncryptedOldValue;pointer:unique" json:"encrypted_old_value"`
	// OldValueSetTime: The time corresponding to the instant that the old value was last
	// changed.
	OldValueSetTime *dtyp.LargeInteger `idl:"name:OldValueSetTime;pointer:unique" json:"old_value_set_time"`
	// Return: The LsarQuerySecret return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QuerySecretResponse) xxx_ToOp(ctx context.Context) *xxx_QuerySecretOperation {
	if o == nil {
		return &xxx_QuerySecretOperation{}
	}
	return &xxx_QuerySecretOperation{
		EncryptedCurrentValue: o.EncryptedCurrentValue,
		CurrentValueSetTime:   o.CurrentValueSetTime,
		EncryptedOldValue:     o.EncryptedOldValue,
		OldValueSetTime:       o.OldValueSetTime,
		Return:                o.Return,
	}
}

func (o *QuerySecretResponse) xxx_FromOp(ctx context.Context, op *xxx_QuerySecretOperation) {
	if o == nil {
		return
	}
	o.EncryptedCurrentValue = op.EncryptedCurrentValue
	o.CurrentValueSetTime = op.CurrentValueSetTime
	o.EncryptedOldValue = op.EncryptedOldValue
	o.OldValueSetTime = op.OldValueSetTime
	o.Return = op.Return
}
func (o *QuerySecretResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QuerySecretResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySecretOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupPrivilegeValueOperation structure represents the LsarLookupPrivilegeValue operation
type xxx_LookupPrivilegeValueOperation struct {
	Policy *Handle             `idl:"name:PolicyHandle" json:"policy"`
	Name   *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	Value  *dtyp.LUID          `idl:"name:Value" json:"value"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupPrivilegeValueOperation) OpNum() int { return 31 }

func (o *xxx_LookupPrivilegeValueOperation) OpName() string {
	return "/lsarpc/v0/LsarLookupPrivilegeValue"
}

func (o *xxx_LookupPrivilegeValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Name {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.Name != nil {
			if err := o.Name.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Name {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.Name == nil {
			o.Name = &dtyp.UnicodeString{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Value {out} (1:{alias=PLUID}*(1))(2:{alias=LUID}(struct))
	{
		if o.Value != nil {
			if err := o.Value.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.LUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Value {out} (1:{alias=PLUID,pointer=ref}*(1))(2:{alias=LUID}(struct))
	{
		if o.Value == nil {
			o.Value = &dtyp.LUID{}
		}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupPrivilegeValueRequest structure represents the LsarLookupPrivilegeValue operation request
type LookupPrivilegeValueRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// Name: A string containing the name of a privilege.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
}

func (o *LookupPrivilegeValueRequest) xxx_ToOp(ctx context.Context) *xxx_LookupPrivilegeValueOperation {
	if o == nil {
		return &xxx_LookupPrivilegeValueOperation{}
	}
	return &xxx_LookupPrivilegeValueOperation{
		Policy: o.Policy,
		Name:   o.Name,
	}
}

func (o *LookupPrivilegeValueRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupPrivilegeValueOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.Name = op.Name
}
func (o *LookupPrivilegeValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *LookupPrivilegeValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupPrivilegeValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupPrivilegeValueResponse structure represents the LsarLookupPrivilegeValue operation response
type LookupPrivilegeValueResponse struct {
	// Value: Used to return a LUID assigned by the server to the privilege by this name.
	Value *dtyp.LUID `idl:"name:Value" json:"value"`
	// Return: The LsarLookupPrivilegeValue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupPrivilegeValueResponse) xxx_ToOp(ctx context.Context) *xxx_LookupPrivilegeValueOperation {
	if o == nil {
		return &xxx_LookupPrivilegeValueOperation{}
	}
	return &xxx_LookupPrivilegeValueOperation{
		Value:  o.Value,
		Return: o.Return,
	}
}

func (o *LookupPrivilegeValueResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupPrivilegeValueOperation) {
	if o == nil {
		return
	}
	o.Value = op.Value
	o.Return = op.Return
}
func (o *LookupPrivilegeValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *LookupPrivilegeValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupPrivilegeValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupPrivilegeNameOperation structure represents the LsarLookupPrivilegeName operation
type xxx_LookupPrivilegeNameOperation struct {
	Policy *Handle             `idl:"name:PolicyHandle" json:"policy"`
	Value  *dtyp.LUID          `idl:"name:Value" json:"value"`
	Name   *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupPrivilegeNameOperation) OpNum() int { return 32 }

func (o *xxx_LookupPrivilegeNameOperation) OpName() string {
	return "/lsarpc/v0/LsarLookupPrivilegeName"
}

func (o *xxx_LookupPrivilegeNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Value {in} (1:{alias=PLUID}*(1))(2:{alias=LUID}(struct))
	{
		if o.Value != nil {
			if err := o.Value.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.LUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Value {in} (1:{alias=PLUID,pointer=ref}*(1))(2:{alias=LUID}(struct))
	{
		if o.Value == nil {
			o.Value = &dtyp.LUID{}
		}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Name {out} (1:{pointer=ref}*(2))(2:{alias=PRPC_UNICODE_STRING}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.Name != nil {
			_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_Name); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Name {out} (1:{pointer=ref}*(2))(2:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &dtyp.UnicodeString{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Name := func(ptr interface{}) { o.Name = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupPrivilegeNameRequest structure represents the LsarLookupPrivilegeName operation request
type LookupPrivilegeNameRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// Value: A LUID that the caller wishes to map to a string name.
	Value *dtyp.LUID `idl:"name:Value" json:"value"`
}

func (o *LookupPrivilegeNameRequest) xxx_ToOp(ctx context.Context) *xxx_LookupPrivilegeNameOperation {
	if o == nil {
		return &xxx_LookupPrivilegeNameOperation{}
	}
	return &xxx_LookupPrivilegeNameOperation{
		Policy: o.Policy,
		Value:  o.Value,
	}
}

func (o *LookupPrivilegeNameRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupPrivilegeNameOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.Value = op.Value
}
func (o *LookupPrivilegeNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *LookupPrivilegeNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupPrivilegeNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupPrivilegeNameResponse structure represents the LsarLookupPrivilegeName operation response
type LookupPrivilegeNameResponse struct {
	// Name: Used to return the string name corresponding to the supplied LUID.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// Return: The LsarLookupPrivilegeName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupPrivilegeNameResponse) xxx_ToOp(ctx context.Context) *xxx_LookupPrivilegeNameOperation {
	if o == nil {
		return &xxx_LookupPrivilegeNameOperation{}
	}
	return &xxx_LookupPrivilegeNameOperation{
		Name:   o.Name,
		Return: o.Return,
	}
}

func (o *LookupPrivilegeNameResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupPrivilegeNameOperation) {
	if o == nil {
		return
	}
	o.Name = op.Name
	o.Return = op.Return
}
func (o *LookupPrivilegeNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *LookupPrivilegeNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupPrivilegeNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupPrivilegeDisplayNameOperation structure represents the LsarLookupPrivilegeDisplayName operation
type xxx_LookupPrivilegeDisplayNameOperation struct {
	Policy                      *Handle             `idl:"name:PolicyHandle" json:"policy"`
	Name                        *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	ClientLanguage              int16               `idl:"name:ClientLanguage" json:"client_language"`
	ClientSystemDefaultLanguage int16               `idl:"name:ClientSystemDefaultLanguage" json:"client_system_default_language"`
	DisplayName                 *dtyp.UnicodeString `idl:"name:DisplayName" json:"display_name"`
	LanguageReturned            uint16              `idl:"name:LanguageReturned" json:"language_returned"`
	Return                      int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupPrivilegeDisplayNameOperation) OpNum() int { return 33 }

func (o *xxx_LookupPrivilegeDisplayNameOperation) OpName() string {
	return "/lsarpc/v0/LsarLookupPrivilegeDisplayName"
}

func (o *xxx_LookupPrivilegeDisplayNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeDisplayNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Name {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.Name != nil {
			if err := o.Name.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClientLanguage {in} (1:(int16))
	{
		if err := w.WriteData(o.ClientLanguage); err != nil {
			return err
		}
	}
	// ClientSystemDefaultLanguage {in} (1:(int16))
	{
		if err := w.WriteData(o.ClientSystemDefaultLanguage); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeDisplayNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Name {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.Name == nil {
			o.Name = &dtyp.UnicodeString{}
		}
		if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientLanguage {in} (1:(int16))
	{
		if err := w.ReadData(&o.ClientLanguage); err != nil {
			return err
		}
	}
	// ClientSystemDefaultLanguage {in} (1:(int16))
	{
		if err := w.ReadData(&o.ClientSystemDefaultLanguage); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeDisplayNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeDisplayNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// DisplayName {out} (1:{pointer=ref}*(2))(2:{alias=PRPC_UNICODE_STRING}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.DisplayName != nil {
			_ptr_DisplayName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DisplayName != nil {
					if err := o.DisplayName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DisplayName, _ptr_DisplayName); err != nil {
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
	// LanguageReturned {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.LanguageReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupPrivilegeDisplayNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// DisplayName {out} (1:{pointer=ref}*(2))(2:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		_ptr_DisplayName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DisplayName == nil {
				o.DisplayName = &dtyp.UnicodeString{}
			}
			if err := o.DisplayName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DisplayName := func(ptr interface{}) { o.DisplayName = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&o.DisplayName, _s_DisplayName, _ptr_DisplayName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// LanguageReturned {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.LanguageReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupPrivilegeDisplayNameRequest structure represents the LsarLookupPrivilegeDisplayName operation request
type LookupPrivilegeDisplayNameRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// Name: A string containing the name of a privilege.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// ClientLanguage: An identifier of the client's language.
	ClientLanguage int16 `idl:"name:ClientLanguage" json:"client_language"`
	// ClientSystemDefaultLanguage: An identifier of the default language of the caller's
	// machine.
	ClientSystemDefaultLanguage int16 `idl:"name:ClientSystemDefaultLanguage" json:"client_system_default_language"`
}

func (o *LookupPrivilegeDisplayNameRequest) xxx_ToOp(ctx context.Context) *xxx_LookupPrivilegeDisplayNameOperation {
	if o == nil {
		return &xxx_LookupPrivilegeDisplayNameOperation{}
	}
	return &xxx_LookupPrivilegeDisplayNameOperation{
		Policy:                      o.Policy,
		Name:                        o.Name,
		ClientLanguage:              o.ClientLanguage,
		ClientSystemDefaultLanguage: o.ClientSystemDefaultLanguage,
	}
}

func (o *LookupPrivilegeDisplayNameRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupPrivilegeDisplayNameOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.Name = op.Name
	o.ClientLanguage = op.ClientLanguage
	o.ClientSystemDefaultLanguage = op.ClientSystemDefaultLanguage
}
func (o *LookupPrivilegeDisplayNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *LookupPrivilegeDisplayNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupPrivilegeDisplayNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupPrivilegeDisplayNameResponse structure represents the LsarLookupPrivilegeDisplayName operation response
type LookupPrivilegeDisplayNameResponse struct {
	// DisplayName: Used to return the display name of the privilege in the language pointed
	// to by the LanguageReturned value.
	DisplayName *dtyp.UnicodeString `idl:"name:DisplayName" json:"display_name"`
	// LanguageReturned: An identifier of the language in which DisplayName was returned.
	LanguageReturned uint16 `idl:"name:LanguageReturned" json:"language_returned"`
	// Return: The LsarLookupPrivilegeDisplayName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupPrivilegeDisplayNameResponse) xxx_ToOp(ctx context.Context) *xxx_LookupPrivilegeDisplayNameOperation {
	if o == nil {
		return &xxx_LookupPrivilegeDisplayNameOperation{}
	}
	return &xxx_LookupPrivilegeDisplayNameOperation{
		DisplayName:      o.DisplayName,
		LanguageReturned: o.LanguageReturned,
		Return:           o.Return,
	}
}

func (o *LookupPrivilegeDisplayNameResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupPrivilegeDisplayNameOperation) {
	if o == nil {
		return
	}
	o.DisplayName = op.DisplayName
	o.LanguageReturned = op.LanguageReturned
	o.Return = op.Return
}
func (o *LookupPrivilegeDisplayNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *LookupPrivilegeDisplayNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupPrivilegeDisplayNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteObjectOperation structure represents the LsarDeleteObject operation
type xxx_DeleteObjectOperation struct {
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	Return int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteObjectOperation) OpNum() int { return 34 }

func (o *xxx_DeleteObjectOperation) OpName() string { return "/lsarpc/v0/LsarDeleteObject" }

func (o *xxx_DeleteObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object != nil {
			if err := o.Object.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object == nil {
			o.Object = &Handle{}
		}
		if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object != nil {
			if err := o.Object.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object == nil {
			o.Object = &Handle{}
		}
		if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteObjectRequest structure represents the LsarDeleteObject operation request
type DeleteObjectRequest struct {
	// ObjectHandle: A handle to an open object of the correct type to be deleted. After
	// successful completion of the call, the handle value cannot be reused.
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
}

func (o *DeleteObjectRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteObjectOperation {
	if o == nil {
		return &xxx_DeleteObjectOperation{}
	}
	return &xxx_DeleteObjectOperation{
		Object: o.Object,
	}
}

func (o *DeleteObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteObjectOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
}
func (o *DeleteObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteObjectResponse structure represents the LsarDeleteObject operation response
type DeleteObjectResponse struct {
	// ObjectHandle: A handle to an open object of the correct type to be deleted. After
	// successful completion of the call, the handle value cannot be reused.
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	// Return: The LsarDeleteObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteObjectResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteObjectOperation {
	if o == nil {
		return &xxx_DeleteObjectOperation{}
	}
	return &xxx_DeleteObjectOperation{
		Object: o.Object,
		Return: o.Return,
	}
}

func (o *DeleteObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteObjectOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
	o.Return = op.Return
}
func (o *DeleteObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateAccountsWithUserRightOperation structure represents the LsarEnumerateAccountsWithUserRight operation
type xxx_EnumerateAccountsWithUserRightOperation struct {
	Policy            *Handle             `idl:"name:PolicyHandle" json:"policy"`
	UserRight         *dtyp.UnicodeString `idl:"name:UserRight;pointer:unique" json:"user_right"`
	EnumerationBuffer *AccountEnumBuffer  `idl:"name:EnumerationBuffer" json:"enumeration_buffer"`
	Return            int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateAccountsWithUserRightOperation) OpNum() int { return 35 }

func (o *xxx_EnumerateAccountsWithUserRightOperation) OpName() string {
	return "/lsarpc/v0/LsarEnumerateAccountsWithUserRight"
}

func (o *xxx_EnumerateAccountsWithUserRightOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountsWithUserRightOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// UserRight {in} (1:{pointer=unique, alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.UserRight != nil {
			_ptr_UserRight := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserRight != nil {
					if err := o.UserRight.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserRight, _ptr_UserRight); err != nil {
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

func (o *xxx_EnumerateAccountsWithUserRightOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// UserRight {in} (1:{pointer=unique, alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		_ptr_UserRight := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserRight == nil {
				o.UserRight = &dtyp.UnicodeString{}
			}
			if err := o.UserRight.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserRight := func(ptr interface{}) { o.UserRight = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&o.UserRight, _s_UserRight, _ptr_UserRight); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountsWithUserRightOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountsWithUserRightOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EnumerationBuffer {out} (1:{alias=PLSAPR_ACCOUNT_ENUM_BUFFER}*(1))(2:{alias=LSAPR_ACCOUNT_ENUM_BUFFER}(struct))
	{
		if o.EnumerationBuffer != nil {
			if err := o.EnumerationBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AccountEnumBuffer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountsWithUserRightOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EnumerationBuffer {out} (1:{alias=PLSAPR_ACCOUNT_ENUM_BUFFER,pointer=ref}*(1))(2:{alias=LSAPR_ACCOUNT_ENUM_BUFFER}(struct))
	{
		if o.EnumerationBuffer == nil {
			o.EnumerationBuffer = &AccountEnumBuffer{}
		}
		if err := o.EnumerationBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateAccountsWithUserRightRequest structure represents the LsarEnumerateAccountsWithUserRight operation request
type EnumerateAccountsWithUserRightRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// UserRight: The name of the right to use in enumeration.
	UserRight *dtyp.UnicodeString `idl:"name:UserRight;pointer:unique" json:"user_right"`
}

func (o *EnumerateAccountsWithUserRightRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateAccountsWithUserRightOperation {
	if o == nil {
		return &xxx_EnumerateAccountsWithUserRightOperation{}
	}
	return &xxx_EnumerateAccountsWithUserRightOperation{
		Policy:    o.Policy,
		UserRight: o.UserRight,
	}
}

func (o *EnumerateAccountsWithUserRightRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateAccountsWithUserRightOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.UserRight = op.UserRight
}
func (o *EnumerateAccountsWithUserRightRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateAccountsWithUserRightRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateAccountsWithUserRightOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateAccountsWithUserRightResponse structure represents the LsarEnumerateAccountsWithUserRight operation response
type EnumerateAccountsWithUserRightResponse struct {
	// EnumerationBuffer: Used to return the list of account objects that have the specified
	// right.
	EnumerationBuffer *AccountEnumBuffer `idl:"name:EnumerationBuffer" json:"enumeration_buffer"`
	// Return: The LsarEnumerateAccountsWithUserRight return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateAccountsWithUserRightResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateAccountsWithUserRightOperation {
	if o == nil {
		return &xxx_EnumerateAccountsWithUserRightOperation{}
	}
	return &xxx_EnumerateAccountsWithUserRightOperation{
		EnumerationBuffer: o.EnumerationBuffer,
		Return:            o.Return,
	}
}

func (o *EnumerateAccountsWithUserRightResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateAccountsWithUserRightOperation) {
	if o == nil {
		return
	}
	o.EnumerationBuffer = op.EnumerationBuffer
	o.Return = op.Return
}
func (o *EnumerateAccountsWithUserRightResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateAccountsWithUserRightResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateAccountsWithUserRightOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateAccountRightsOperation structure represents the LsarEnumerateAccountRights operation
type xxx_EnumerateAccountRightsOperation struct {
	Policy     *Handle       `idl:"name:PolicyHandle" json:"policy"`
	AccountSID *dtyp.SID     `idl:"name:AccountSid" json:"account_sid"`
	UserRights *UserRightSet `idl:"name:UserRights" json:"user_rights"`
	Return     int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateAccountRightsOperation) OpNum() int { return 36 }

func (o *xxx_EnumerateAccountRightsOperation) OpName() string {
	return "/lsarpc/v0/LsarEnumerateAccountRights"
}

func (o *xxx_EnumerateAccountRightsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountRightsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// AccountSid {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.AccountSID != nil {
			if err := o.AccountSID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountRightsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// AccountSid {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.AccountSID == nil {
			o.AccountSID = &dtyp.SID{}
		}
		if err := o.AccountSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountRightsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountRightsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// UserRights {out} (1:{alias=PLSAPR_USER_RIGHT_SET}*(1))(2:{alias=LSAPR_USER_RIGHT_SET}(struct))
	{
		if o.UserRights != nil {
			if err := o.UserRights.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserRightSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAccountRightsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// UserRights {out} (1:{alias=PLSAPR_USER_RIGHT_SET,pointer=ref}*(1))(2:{alias=LSAPR_USER_RIGHT_SET}(struct))
	{
		if o.UserRights == nil {
			o.UserRights = &UserRightSet{}
		}
		if err := o.UserRights.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateAccountRightsRequest structure represents the LsarEnumerateAccountRights operation request
type EnumerateAccountRightsRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// AccountSid: A SID of the account object that the caller is inquiring about.
	AccountSID *dtyp.SID `idl:"name:AccountSid" json:"account_sid"`
}

func (o *EnumerateAccountRightsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateAccountRightsOperation {
	if o == nil {
		return &xxx_EnumerateAccountRightsOperation{}
	}
	return &xxx_EnumerateAccountRightsOperation{
		Policy:     o.Policy,
		AccountSID: o.AccountSID,
	}
}

func (o *EnumerateAccountRightsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateAccountRightsOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.AccountSID = op.AccountSID
}
func (o *EnumerateAccountRightsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateAccountRightsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateAccountRightsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateAccountRightsResponse structure represents the LsarEnumerateAccountRights operation response
type EnumerateAccountRightsResponse struct {
	// UserRights: Used to return a list of right names associated with the account.
	UserRights *UserRightSet `idl:"name:UserRights" json:"user_rights"`
	// Return: The LsarEnumerateAccountRights return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateAccountRightsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateAccountRightsOperation {
	if o == nil {
		return &xxx_EnumerateAccountRightsOperation{}
	}
	return &xxx_EnumerateAccountRightsOperation{
		UserRights: o.UserRights,
		Return:     o.Return,
	}
}

func (o *EnumerateAccountRightsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateAccountRightsOperation) {
	if o == nil {
		return
	}
	o.UserRights = op.UserRights
	o.Return = op.Return
}
func (o *EnumerateAccountRightsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateAccountRightsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateAccountRightsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddAccountRightsOperation structure represents the LsarAddAccountRights operation
type xxx_AddAccountRightsOperation struct {
	Policy     *Handle       `idl:"name:PolicyHandle" json:"policy"`
	AccountSID *dtyp.SID     `idl:"name:AccountSid" json:"account_sid"`
	UserRights *UserRightSet `idl:"name:UserRights" json:"user_rights"`
	Return     int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_AddAccountRightsOperation) OpNum() int { return 37 }

func (o *xxx_AddAccountRightsOperation) OpName() string { return "/lsarpc/v0/LsarAddAccountRights" }

func (o *xxx_AddAccountRightsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccountRightsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// AccountSid {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.AccountSID != nil {
			if err := o.AccountSID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// UserRights {in} (1:{alias=PLSAPR_USER_RIGHT_SET}*(1))(2:{alias=LSAPR_USER_RIGHT_SET}(struct))
	{
		if o.UserRights != nil {
			if err := o.UserRights.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserRightSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccountRightsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// AccountSid {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.AccountSID == nil {
			o.AccountSID = &dtyp.SID{}
		}
		if err := o.AccountSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// UserRights {in} (1:{alias=PLSAPR_USER_RIGHT_SET,pointer=ref}*(1))(2:{alias=LSAPR_USER_RIGHT_SET}(struct))
	{
		if o.UserRights == nil {
			o.UserRights = &UserRightSet{}
		}
		if err := o.UserRights.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccountRightsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccountRightsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccountRightsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddAccountRightsRequest structure represents the LsarAddAccountRights operation request
type AddAccountRightsRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// AccountSid: A security identifier of an account to add the rights to.
	AccountSID *dtyp.SID `idl:"name:AccountSid" json:"account_sid"`
	// UserRights: A set of right names to add to the account.
	UserRights *UserRightSet `idl:"name:UserRights" json:"user_rights"`
}

func (o *AddAccountRightsRequest) xxx_ToOp(ctx context.Context) *xxx_AddAccountRightsOperation {
	if o == nil {
		return &xxx_AddAccountRightsOperation{}
	}
	return &xxx_AddAccountRightsOperation{
		Policy:     o.Policy,
		AccountSID: o.AccountSID,
		UserRights: o.UserRights,
	}
}

func (o *AddAccountRightsRequest) xxx_FromOp(ctx context.Context, op *xxx_AddAccountRightsOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.AccountSID = op.AccountSID
	o.UserRights = op.UserRights
}
func (o *AddAccountRightsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddAccountRightsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddAccountRightsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddAccountRightsResponse structure represents the LsarAddAccountRights operation response
type AddAccountRightsResponse struct {
	// Return: The LsarAddAccountRights return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddAccountRightsResponse) xxx_ToOp(ctx context.Context) *xxx_AddAccountRightsOperation {
	if o == nil {
		return &xxx_AddAccountRightsOperation{}
	}
	return &xxx_AddAccountRightsOperation{
		Return: o.Return,
	}
}

func (o *AddAccountRightsResponse) xxx_FromOp(ctx context.Context, op *xxx_AddAccountRightsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddAccountRightsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddAccountRightsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddAccountRightsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveAccountRightsOperation structure represents the LsarRemoveAccountRights operation
type xxx_RemoveAccountRightsOperation struct {
	Policy     *Handle       `idl:"name:PolicyHandle" json:"policy"`
	AccountSID *dtyp.SID     `idl:"name:AccountSid" json:"account_sid"`
	AllRights  uint8         `idl:"name:AllRights" json:"all_rights"`
	UserRights *UserRightSet `idl:"name:UserRights" json:"user_rights"`
	Return     int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveAccountRightsOperation) OpNum() int { return 38 }

func (o *xxx_RemoveAccountRightsOperation) OpName() string {
	return "/lsarpc/v0/LsarRemoveAccountRights"
}

func (o *xxx_RemoveAccountRightsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveAccountRightsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// AccountSid {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.AccountSID != nil {
			if err := o.AccountSID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// AllRights {in} (1:(uchar))
	{
		if err := w.WriteData(o.AllRights); err != nil {
			return err
		}
	}
	// UserRights {in} (1:{alias=PLSAPR_USER_RIGHT_SET}*(1))(2:{alias=LSAPR_USER_RIGHT_SET}(struct))
	{
		if o.UserRights != nil {
			if err := o.UserRights.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserRightSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveAccountRightsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// AccountSid {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.AccountSID == nil {
			o.AccountSID = &dtyp.SID{}
		}
		if err := o.AccountSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// AllRights {in} (1:(uchar))
	{
		if err := w.ReadData(&o.AllRights); err != nil {
			return err
		}
	}
	// UserRights {in} (1:{alias=PLSAPR_USER_RIGHT_SET,pointer=ref}*(1))(2:{alias=LSAPR_USER_RIGHT_SET}(struct))
	{
		if o.UserRights == nil {
			o.UserRights = &UserRightSet{}
		}
		if err := o.UserRights.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveAccountRightsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveAccountRightsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveAccountRightsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveAccountRightsRequest structure represents the LsarRemoveAccountRights operation request
type RemoveAccountRightsRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// AccountSid: A security descriptor of an account object.
	AccountSID *dtyp.SID `idl:"name:AccountSid" json:"account_sid"`
	// AllRights: If this field is not set to 0, all rights will be removed.
	AllRights uint8 `idl:"name:AllRights" json:"all_rights"`
	// UserRights: A set of rights to remove from the account.
	UserRights *UserRightSet `idl:"name:UserRights" json:"user_rights"`
}

func (o *RemoveAccountRightsRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveAccountRightsOperation {
	if o == nil {
		return &xxx_RemoveAccountRightsOperation{}
	}
	return &xxx_RemoveAccountRightsOperation{
		Policy:     o.Policy,
		AccountSID: o.AccountSID,
		AllRights:  o.AllRights,
		UserRights: o.UserRights,
	}
}

func (o *RemoveAccountRightsRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveAccountRightsOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.AccountSID = op.AccountSID
	o.AllRights = op.AllRights
	o.UserRights = op.UserRights
}
func (o *RemoveAccountRightsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveAccountRightsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveAccountRightsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveAccountRightsResponse structure represents the LsarRemoveAccountRights operation response
type RemoveAccountRightsResponse struct {
	// Return: The LsarRemoveAccountRights return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveAccountRightsResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveAccountRightsOperation {
	if o == nil {
		return &xxx_RemoveAccountRightsOperation{}
	}
	return &xxx_RemoveAccountRightsOperation{
		Return: o.Return,
	}
}

func (o *RemoveAccountRightsResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveAccountRightsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveAccountRightsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveAccountRightsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveAccountRightsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryTrustedDomainInfoOperation structure represents the LsarQueryTrustedDomainInfo operation
type xxx_QueryTrustedDomainInfoOperation struct {
	Policy                   *Handle                 `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainSID         *dtyp.SID               `idl:"name:TrustedDomainSid" json:"trusted_domain_sid"`
	InformationClass         TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
	TrustedDomainInformation *TrustedDomainInfo      `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
	Return                   int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryTrustedDomainInfoOperation) OpNum() int { return 39 }

func (o *xxx_QueryTrustedDomainInfoOperation) OpName() string {
	return "/lsarpc/v0/LsarQueryTrustedDomainInfo"
}

func (o *xxx_QueryTrustedDomainInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryTrustedDomainInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainSid {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.TrustedDomainSID != nil {
			if err := o.TrustedDomainSID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryTrustedDomainInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainSid {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.TrustedDomainSID == nil {
			o.TrustedDomainSID = &dtyp.SID{}
		}
		if err := o.TrustedDomainSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryTrustedDomainInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryTrustedDomainInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TrustedDomainInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO}*(1))(3:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		if o.TrustedDomainInformation != nil {
			_ptr_TrustedDomainInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swTrustedDomainInformation := uint16(o.InformationClass)
				if o.TrustedDomainInformation != nil {
					if err := o.TrustedDomainInformation.MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
						return err
					}
				} else {
					if err := (&TrustedDomainInfo{}).MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TrustedDomainInformation, _ptr_TrustedDomainInformation); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryTrustedDomainInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TrustedDomainInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO,pointer=ref}*(1))(3:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		_ptr_TrustedDomainInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TrustedDomainInformation == nil {
				o.TrustedDomainInformation = &TrustedDomainInfo{}
			}
			_swTrustedDomainInformation := uint16(o.InformationClass)
			if err := o.TrustedDomainInformation.UnmarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
				return err
			}
			return nil
		})
		_s_TrustedDomainInformation := func(ptr interface{}) { o.TrustedDomainInformation = *ptr.(**TrustedDomainInfo) }
		if err := w.ReadPointer(&o.TrustedDomainInformation, _s_TrustedDomainInformation, _ptr_TrustedDomainInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryTrustedDomainInfoRequest structure represents the LsarQueryTrustedDomainInfo operation request
type QueryTrustedDomainInfoRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainSid: A security descriptor of the trusted domain object.
	TrustedDomainSID *dtyp.SID `idl:"name:TrustedDomainSid" json:"trusted_domain_sid"`
	// InformationClass: Identifies the type of information the caller is interested in.
	InformationClass TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
}

func (o *QueryTrustedDomainInfoRequest) xxx_ToOp(ctx context.Context) *xxx_QueryTrustedDomainInfoOperation {
	if o == nil {
		return &xxx_QueryTrustedDomainInfoOperation{}
	}
	return &xxx_QueryTrustedDomainInfoOperation{
		Policy:           o.Policy,
		TrustedDomainSID: o.TrustedDomainSID,
		InformationClass: o.InformationClass,
	}
}

func (o *QueryTrustedDomainInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryTrustedDomainInfoOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainSID = op.TrustedDomainSID
	o.InformationClass = op.InformationClass
}
func (o *QueryTrustedDomainInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryTrustedDomainInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryTrustedDomainInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryTrustedDomainInfoResponse structure represents the LsarQueryTrustedDomainInfo operation response
type QueryTrustedDomainInfoResponse struct {
	// TrustedDomainInformation: Used to return the information on the trusted domain object
	// to the caller.
	TrustedDomainInformation *TrustedDomainInfo `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
	// Return: The LsarQueryTrustedDomainInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryTrustedDomainInfoResponse) xxx_ToOp(ctx context.Context) *xxx_QueryTrustedDomainInfoOperation {
	if o == nil {
		return &xxx_QueryTrustedDomainInfoOperation{}
	}
	return &xxx_QueryTrustedDomainInfoOperation{
		TrustedDomainInformation: o.TrustedDomainInformation,
		Return:                   o.Return,
	}
}

func (o *QueryTrustedDomainInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryTrustedDomainInfoOperation) {
	if o == nil {
		return
	}
	o.TrustedDomainInformation = op.TrustedDomainInformation
	o.Return = op.Return
}
func (o *QueryTrustedDomainInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryTrustedDomainInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryTrustedDomainInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetTrustedDomainInfoOperation structure represents the LsarSetTrustedDomainInfo operation
type xxx_SetTrustedDomainInfoOperation struct {
	Policy                   *Handle                 `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainSID         *dtyp.SID               `idl:"name:TrustedDomainSid" json:"trusted_domain_sid"`
	InformationClass         TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
	TrustedDomainInformation *TrustedDomainInfo      `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
	Return                   int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTrustedDomainInfoOperation) OpNum() int { return 40 }

func (o *xxx_SetTrustedDomainInfoOperation) OpName() string {
	return "/lsarpc/v0/LsarSetTrustedDomainInfo"
}

func (o *xxx_SetTrustedDomainInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTrustedDomainInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainSid {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.TrustedDomainSID != nil {
			if err := o.TrustedDomainSID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	// TrustedDomainInformation {in} (1:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO}*(1))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		_swTrustedDomainInformation := uint16(o.InformationClass)
		if o.TrustedDomainInformation != nil {
			if err := o.TrustedDomainInformation.MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo{}).MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTrustedDomainInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainSid {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.TrustedDomainSID == nil {
			o.TrustedDomainSID = &dtyp.SID{}
		}
		if err := o.TrustedDomainSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	// TrustedDomainInformation {in} (1:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO,pointer=ref}*(1))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		if o.TrustedDomainInformation == nil {
			o.TrustedDomainInformation = &TrustedDomainInfo{}
		}
		_swTrustedDomainInformation := uint16(o.InformationClass)
		if err := o.TrustedDomainInformation.UnmarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTrustedDomainInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTrustedDomainInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTrustedDomainInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetTrustedDomainInfoRequest structure represents the LsarSetTrustedDomainInfo operation request
type SetTrustedDomainInfoRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainSid: A SID of the trusted domain object to be modified.
	TrustedDomainSID *dtyp.SID `idl:"name:TrustedDomainSid" json:"trusted_domain_sid"`
	// InformationClass: Identifies the type of information to be set on the trusted domain
	// object.
	InformationClass TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
	// TrustedDomainInformation: Information to be set on the trusted domain object.
	TrustedDomainInformation *TrustedDomainInfo `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
}

func (o *SetTrustedDomainInfoRequest) xxx_ToOp(ctx context.Context) *xxx_SetTrustedDomainInfoOperation {
	if o == nil {
		return &xxx_SetTrustedDomainInfoOperation{}
	}
	return &xxx_SetTrustedDomainInfoOperation{
		Policy:                   o.Policy,
		TrustedDomainSID:         o.TrustedDomainSID,
		InformationClass:         o.InformationClass,
		TrustedDomainInformation: o.TrustedDomainInformation,
	}
}

func (o *SetTrustedDomainInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTrustedDomainInfoOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainSID = op.TrustedDomainSID
	o.InformationClass = op.InformationClass
	o.TrustedDomainInformation = op.TrustedDomainInformation
}
func (o *SetTrustedDomainInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetTrustedDomainInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTrustedDomainInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTrustedDomainInfoResponse structure represents the LsarSetTrustedDomainInfo operation response
type SetTrustedDomainInfoResponse struct {
	// Return: The LsarSetTrustedDomainInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetTrustedDomainInfoResponse) xxx_ToOp(ctx context.Context) *xxx_SetTrustedDomainInfoOperation {
	if o == nil {
		return &xxx_SetTrustedDomainInfoOperation{}
	}
	return &xxx_SetTrustedDomainInfoOperation{
		Return: o.Return,
	}
}

func (o *SetTrustedDomainInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTrustedDomainInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetTrustedDomainInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetTrustedDomainInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTrustedDomainInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteTrustedDomainOperation structure represents the LsarDeleteTrustedDomain operation
type xxx_DeleteTrustedDomainOperation struct {
	Policy           *Handle   `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainSID *dtyp.SID `idl:"name:TrustedDomainSid" json:"trusted_domain_sid"`
	Return           int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteTrustedDomainOperation) OpNum() int { return 41 }

func (o *xxx_DeleteTrustedDomainOperation) OpName() string {
	return "/lsarpc/v0/LsarDeleteTrustedDomain"
}

func (o *xxx_DeleteTrustedDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteTrustedDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainSid {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.TrustedDomainSID != nil {
			if err := o.TrustedDomainSID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteTrustedDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainSid {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.TrustedDomainSID == nil {
			o.TrustedDomainSID = &dtyp.SID{}
		}
		if err := o.TrustedDomainSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteTrustedDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteTrustedDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteTrustedDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteTrustedDomainRequest structure represents the LsarDeleteTrustedDomain operation request
type DeleteTrustedDomainRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainSid: A security descriptor of the TDO to be deleted.
	TrustedDomainSID *dtyp.SID `idl:"name:TrustedDomainSid" json:"trusted_domain_sid"`
}

func (o *DeleteTrustedDomainRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteTrustedDomainOperation {
	if o == nil {
		return &xxx_DeleteTrustedDomainOperation{}
	}
	return &xxx_DeleteTrustedDomainOperation{
		Policy:           o.Policy,
		TrustedDomainSID: o.TrustedDomainSID,
	}
}

func (o *DeleteTrustedDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteTrustedDomainOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainSID = op.TrustedDomainSID
}
func (o *DeleteTrustedDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteTrustedDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteTrustedDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteTrustedDomainResponse structure represents the LsarDeleteTrustedDomain operation response
type DeleteTrustedDomainResponse struct {
	// Return: The LsarDeleteTrustedDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteTrustedDomainResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteTrustedDomainOperation {
	if o == nil {
		return &xxx_DeleteTrustedDomainOperation{}
	}
	return &xxx_DeleteTrustedDomainOperation{
		Return: o.Return,
	}
}

func (o *DeleteTrustedDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteTrustedDomainOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteTrustedDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteTrustedDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteTrustedDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StorePrivateDataOperation structure represents the LsarStorePrivateData operation
type xxx_StorePrivateDataOperation struct {
	Policy        *Handle             `idl:"name:PolicyHandle" json:"policy"`
	KeyName       *dtyp.UnicodeString `idl:"name:KeyName" json:"key_name"`
	EncryptedData *CRCipherValue      `idl:"name:EncryptedData;pointer:unique" json:"encrypted_data"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_StorePrivateDataOperation) OpNum() int { return 42 }

func (o *xxx_StorePrivateDataOperation) OpName() string { return "/lsarpc/v0/LsarStorePrivateData" }

func (o *xxx_StorePrivateDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StorePrivateDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// KeyName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.KeyName != nil {
			if err := o.KeyName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// EncryptedData {in} (1:{pointer=unique, alias=PLSAPR_CR_CIPHER_VALUE}*(1))(2:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		if o.EncryptedData != nil {
			_ptr_EncryptedData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedData != nil {
					if err := o.EncryptedData.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedData, _ptr_EncryptedData); err != nil {
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

func (o *xxx_StorePrivateDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// KeyName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.KeyName == nil {
			o.KeyName = &dtyp.UnicodeString{}
		}
		if err := o.KeyName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// EncryptedData {in} (1:{pointer=unique, alias=PLSAPR_CR_CIPHER_VALUE}*(1))(2:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		_ptr_EncryptedData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EncryptedData == nil {
				o.EncryptedData = &CRCipherValue{}
			}
			if err := o.EncryptedData.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedData := func(ptr interface{}) { o.EncryptedData = *ptr.(**CRCipherValue) }
		if err := w.ReadPointer(&o.EncryptedData, _s_EncryptedData, _ptr_EncryptedData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StorePrivateDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StorePrivateDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StorePrivateDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StorePrivateDataRequest structure represents the LsarStorePrivateData operation request
type StorePrivateDataRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// KeyName: The name under which private data will be stored.
	KeyName *dtyp.UnicodeString `idl:"name:KeyName" json:"key_name"`
	// EncryptedData: The secret value to be stored.
	EncryptedData *CRCipherValue `idl:"name:EncryptedData;pointer:unique" json:"encrypted_data"`
}

func (o *StorePrivateDataRequest) xxx_ToOp(ctx context.Context) *xxx_StorePrivateDataOperation {
	if o == nil {
		return &xxx_StorePrivateDataOperation{}
	}
	return &xxx_StorePrivateDataOperation{
		Policy:        o.Policy,
		KeyName:       o.KeyName,
		EncryptedData: o.EncryptedData,
	}
}

func (o *StorePrivateDataRequest) xxx_FromOp(ctx context.Context, op *xxx_StorePrivateDataOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.KeyName = op.KeyName
	o.EncryptedData = op.EncryptedData
}
func (o *StorePrivateDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StorePrivateDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StorePrivateDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StorePrivateDataResponse structure represents the LsarStorePrivateData operation response
type StorePrivateDataResponse struct {
	// Return: The LsarStorePrivateData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StorePrivateDataResponse) xxx_ToOp(ctx context.Context) *xxx_StorePrivateDataOperation {
	if o == nil {
		return &xxx_StorePrivateDataOperation{}
	}
	return &xxx_StorePrivateDataOperation{
		Return: o.Return,
	}
}

func (o *StorePrivateDataResponse) xxx_FromOp(ctx context.Context, op *xxx_StorePrivateDataOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *StorePrivateDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StorePrivateDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StorePrivateDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RetrievePrivateDataOperation structure represents the LsarRetrievePrivateData operation
type xxx_RetrievePrivateDataOperation struct {
	Policy        *Handle             `idl:"name:PolicyHandle" json:"policy"`
	KeyName       *dtyp.UnicodeString `idl:"name:KeyName" json:"key_name"`
	EncryptedData *CRCipherValue      `idl:"name:EncryptedData" json:"encrypted_data"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_RetrievePrivateDataOperation) OpNum() int { return 43 }

func (o *xxx_RetrievePrivateDataOperation) OpName() string {
	return "/lsarpc/v0/LsarRetrievePrivateData"
}

func (o *xxx_RetrievePrivateDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrievePrivateDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// KeyName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.KeyName != nil {
			if err := o.KeyName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// EncryptedData {in, out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		if o.EncryptedData != nil {
			_ptr_EncryptedData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedData != nil {
					if err := o.EncryptedData.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedData, _ptr_EncryptedData); err != nil {
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

func (o *xxx_RetrievePrivateDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// KeyName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.KeyName == nil {
			o.KeyName = &dtyp.UnicodeString{}
		}
		if err := o.KeyName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// EncryptedData {in, out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE,pointer=ref}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		_ptr_EncryptedData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EncryptedData == nil {
				o.EncryptedData = &CRCipherValue{}
			}
			if err := o.EncryptedData.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedData := func(ptr interface{}) { o.EncryptedData = *ptr.(**CRCipherValue) }
		if err := w.ReadPointer(&o.EncryptedData, _s_EncryptedData, _ptr_EncryptedData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrievePrivateDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrievePrivateDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EncryptedData {in, out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		if o.EncryptedData != nil {
			_ptr_EncryptedData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedData != nil {
					if err := o.EncryptedData.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CRCipherValue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedData, _ptr_EncryptedData); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrievePrivateDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EncryptedData {in, out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_CR_CIPHER_VALUE,pointer=ref}*(1))(3:{alias=LSAPR_CR_CIPHER_VALUE}(struct))
	{
		_ptr_EncryptedData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EncryptedData == nil {
				o.EncryptedData = &CRCipherValue{}
			}
			if err := o.EncryptedData.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedData := func(ptr interface{}) { o.EncryptedData = *ptr.(**CRCipherValue) }
		if err := w.ReadPointer(&o.EncryptedData, _s_EncryptedData, _ptr_EncryptedData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RetrievePrivateDataRequest structure represents the LsarRetrievePrivateData operation request
type RetrievePrivateDataRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// KeyName: The name identifying the secret value to be retrieved.
	KeyName *dtyp.UnicodeString `idl:"name:KeyName" json:"key_name"`
	// EncryptedData: Receives the encrypted value of the secret object.
	EncryptedData *CRCipherValue `idl:"name:EncryptedData" json:"encrypted_data"`
}

func (o *RetrievePrivateDataRequest) xxx_ToOp(ctx context.Context) *xxx_RetrievePrivateDataOperation {
	if o == nil {
		return &xxx_RetrievePrivateDataOperation{}
	}
	return &xxx_RetrievePrivateDataOperation{
		Policy:        o.Policy,
		KeyName:       o.KeyName,
		EncryptedData: o.EncryptedData,
	}
}

func (o *RetrievePrivateDataRequest) xxx_FromOp(ctx context.Context, op *xxx_RetrievePrivateDataOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.KeyName = op.KeyName
	o.EncryptedData = op.EncryptedData
}
func (o *RetrievePrivateDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RetrievePrivateDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrievePrivateDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RetrievePrivateDataResponse structure represents the LsarRetrievePrivateData operation response
type RetrievePrivateDataResponse struct {
	// EncryptedData: Receives the encrypted value of the secret object.
	EncryptedData *CRCipherValue `idl:"name:EncryptedData" json:"encrypted_data"`
	// Return: The LsarRetrievePrivateData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RetrievePrivateDataResponse) xxx_ToOp(ctx context.Context) *xxx_RetrievePrivateDataOperation {
	if o == nil {
		return &xxx_RetrievePrivateDataOperation{}
	}
	return &xxx_RetrievePrivateDataOperation{
		EncryptedData: o.EncryptedData,
		Return:        o.Return,
	}
}

func (o *RetrievePrivateDataResponse) xxx_FromOp(ctx context.Context, op *xxx_RetrievePrivateDataOperation) {
	if o == nil {
		return
	}
	o.EncryptedData = op.EncryptedData
	o.Return = op.Return
}
func (o *RetrievePrivateDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RetrievePrivateDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrievePrivateDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenPolicy2Operation structure represents the LsarOpenPolicy2 operation
type xxx_OpenPolicy2Operation struct {
	SystemName       string            `idl:"name:SystemName;string;pointer:unique" json:"system_name"`
	ObjectAttributes *ObjectAttributes `idl:"name:ObjectAttributes" json:"object_attributes"`
	DesiredAccess    uint32            `idl:"name:DesiredAccess" json:"desired_access"`
	Policy           *Handle           `idl:"name:PolicyHandle" json:"policy"`
	Return           int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenPolicy2Operation) OpNum() int { return 44 }

func (o *xxx_OpenPolicy2Operation) OpName() string { return "/lsarpc/v0/LsarOpenPolicy2" }

func (o *xxx_OpenPolicy2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicy2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SystemName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.SystemName != "" {
			_ptr_SystemName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SystemName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SystemName, _ptr_SystemName); err != nil {
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
	// ObjectAttributes {in} (1:{alias=PLSAPR_OBJECT_ATTRIBUTES}*(1))(2:{alias=LSAPR_OBJECT_ATTRIBUTES}(struct))
	{
		if o.ObjectAttributes != nil {
			if err := o.ObjectAttributes.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectAttributes{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicy2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SystemName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_SystemName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SystemName); err != nil {
				return err
			}
			return nil
		})
		_s_SystemName := func(ptr interface{}) { o.SystemName = *ptr.(*string) }
		if err := w.ReadPointer(&o.SystemName, _s_SystemName, _ptr_SystemName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ObjectAttributes {in} (1:{alias=PLSAPR_OBJECT_ATTRIBUTES,pointer=ref}*(1))(2:{alias=LSAPR_OBJECT_ATTRIBUTES}(struct))
	{
		if o.ObjectAttributes == nil {
			o.ObjectAttributes = &ObjectAttributes{}
		}
		if err := o.ObjectAttributes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicy2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicy2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicy2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenPolicy2Request structure represents the LsarOpenPolicy2 operation request
type OpenPolicy2Request struct {
	// SystemName: This parameter does not have any effect on message processing in any
	// environment. It MUST be ignored on receipt.
	SystemName string `idl:"name:SystemName;string;pointer:unique" json:"system_name"`
	// ObjectAttributes: This parameter does not have any effect on message processing in
	// any environment. All fields MUST<59> be ignored except RootDirectory which MUST be
	// NULL.
	ObjectAttributes *ObjectAttributes `idl:"name:ObjectAttributes" json:"object_attributes"`
	// DesiredAccess: An ACCESS_MASK value that specifies the requested access rights that
	// MUST be granted on the returned PolicyHandle if the request is successful.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *OpenPolicy2Request) xxx_ToOp(ctx context.Context) *xxx_OpenPolicy2Operation {
	if o == nil {
		return &xxx_OpenPolicy2Operation{}
	}
	return &xxx_OpenPolicy2Operation{
		SystemName:       o.SystemName,
		ObjectAttributes: o.ObjectAttributes,
		DesiredAccess:    o.DesiredAccess,
	}
}

func (o *OpenPolicy2Request) xxx_FromOp(ctx context.Context, op *xxx_OpenPolicy2Operation) {
	if o == nil {
		return
	}
	o.SystemName = op.SystemName
	o.ObjectAttributes = op.ObjectAttributes
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenPolicy2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenPolicy2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPolicy2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenPolicy2Response structure represents the LsarOpenPolicy2 operation response
type OpenPolicy2Response struct {
	// PolicyHandle: An RPC context handle (as specified in section 2.2.2.1) that represents
	// a reference to the abstract data model of a policy object, as specified in section
	// 3.1.1.1.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// Return: The LsarOpenPolicy2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenPolicy2Response) xxx_ToOp(ctx context.Context) *xxx_OpenPolicy2Operation {
	if o == nil {
		return &xxx_OpenPolicy2Operation{}
	}
	return &xxx_OpenPolicy2Operation{
		Policy: o.Policy,
		Return: o.Return,
	}
}

func (o *OpenPolicy2Response) xxx_FromOp(ctx context.Context, op *xxx_OpenPolicy2Operation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.Return = op.Return
}
func (o *OpenPolicy2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenPolicy2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPolicy2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInformationPolicy2Operation structure represents the LsarQueryInformationPolicy2 operation
type xxx_QueryInformationPolicy2Operation struct {
	Policy            *Handle                `idl:"name:PolicyHandle" json:"policy"`
	InformationClass  PolicyInformationClass `idl:"name:InformationClass" json:"information_class"`
	PolicyInformation *PolicyInformation     `idl:"name:PolicyInformation;switch_is:InformationClass" json:"policy_information"`
	Return            int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInformationPolicy2Operation) OpNum() int { return 46 }

func (o *xxx_QueryInformationPolicy2Operation) OpName() string {
	return "/lsarpc/v0/LsarQueryInformationPolicy2"
}

func (o *xxx_QueryInformationPolicy2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationPolicy2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InformationClass {in} (1:{alias=POLICY_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationPolicy2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=POLICY_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationPolicy2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationPolicy2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PolicyInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=PLSAPR_POLICY_INFORMATION}*(1))(3:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_INFORMATION}(union))
	{
		if o.PolicyInformation != nil {
			_ptr_PolicyInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swPolicyInformation := uint16(o.InformationClass)
				if o.PolicyInformation != nil {
					if err := o.PolicyInformation.MarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
						return err
					}
				} else {
					if err := (&PolicyInformation{}).MarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyInformation, _ptr_PolicyInformation); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationPolicy2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PolicyInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=PLSAPR_POLICY_INFORMATION,pointer=ref}*(1))(3:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_INFORMATION}(union))
	{
		_ptr_PolicyInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyInformation == nil {
				o.PolicyInformation = &PolicyInformation{}
			}
			_swPolicyInformation := uint16(o.InformationClass)
			if err := o.PolicyInformation.UnmarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
				return err
			}
			return nil
		})
		_s_PolicyInformation := func(ptr interface{}) { o.PolicyInformation = *ptr.(**PolicyInformation) }
		if err := w.ReadPointer(&o.PolicyInformation, _s_PolicyInformation, _ptr_PolicyInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryInformationPolicy2Request structure represents the LsarQueryInformationPolicy2 operation request
type QueryInformationPolicy2Request struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// InformationClass: A parameter that specifies what type of information the caller
	// is requesting.
	InformationClass PolicyInformationClass `idl:"name:InformationClass" json:"information_class"`
}

func (o *QueryInformationPolicy2Request) xxx_ToOp(ctx context.Context) *xxx_QueryInformationPolicy2Operation {
	if o == nil {
		return &xxx_QueryInformationPolicy2Operation{}
	}
	return &xxx_QueryInformationPolicy2Operation{
		Policy:           o.Policy,
		InformationClass: o.InformationClass,
	}
}

func (o *QueryInformationPolicy2Request) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationPolicy2Operation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.InformationClass = op.InformationClass
}
func (o *QueryInformationPolicy2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInformationPolicy2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationPolicy2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInformationPolicy2Response structure represents the LsarQueryInformationPolicy2 operation response
type QueryInformationPolicy2Response struct {
	// PolicyInformation: A parameter that references policy information structure on return.
	PolicyInformation *PolicyInformation `idl:"name:PolicyInformation;switch_is:InformationClass" json:"policy_information"`
	// Return: The LsarQueryInformationPolicy2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryInformationPolicy2Response) xxx_ToOp(ctx context.Context) *xxx_QueryInformationPolicy2Operation {
	if o == nil {
		return &xxx_QueryInformationPolicy2Operation{}
	}
	return &xxx_QueryInformationPolicy2Operation{
		PolicyInformation: o.PolicyInformation,
		Return:            o.Return,
	}
}

func (o *QueryInformationPolicy2Response) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationPolicy2Operation) {
	if o == nil {
		return
	}
	o.PolicyInformation = op.PolicyInformation
	o.Return = op.Return
}
func (o *QueryInformationPolicy2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInformationPolicy2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationPolicy2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInformationPolicy2Operation structure represents the LsarSetInformationPolicy2 operation
type xxx_SetInformationPolicy2Operation struct {
	Policy            *Handle                `idl:"name:PolicyHandle" json:"policy"`
	InformationClass  PolicyInformationClass `idl:"name:InformationClass" json:"information_class"`
	PolicyInformation *PolicyInformation     `idl:"name:PolicyInformation;switch_is:InformationClass" json:"policy_information"`
	Return            int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInformationPolicy2Operation) OpNum() int { return 47 }

func (o *xxx_SetInformationPolicy2Operation) OpName() string {
	return "/lsarpc/v0/LsarSetInformationPolicy2"
}

func (o *xxx_SetInformationPolicy2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationPolicy2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InformationClass {in} (1:{alias=POLICY_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	// PolicyInformation {in} (1:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=PLSAPR_POLICY_INFORMATION}*(1))(2:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_INFORMATION}(union))
	{
		_swPolicyInformation := uint16(o.InformationClass)
		if o.PolicyInformation != nil {
			if err := o.PolicyInformation.MarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
				return err
			}
		} else {
			if err := (&PolicyInformation{}).MarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationPolicy2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=POLICY_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	// PolicyInformation {in} (1:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=PLSAPR_POLICY_INFORMATION,pointer=ref}*(1))(2:{switch_type={alias=POLICY_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_INFORMATION}(union))
	{
		if o.PolicyInformation == nil {
			o.PolicyInformation = &PolicyInformation{}
		}
		_swPolicyInformation := uint16(o.InformationClass)
		if err := o.PolicyInformation.UnmarshalUnionNDR(ctx, w, _swPolicyInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationPolicy2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationPolicy2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationPolicy2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInformationPolicy2Request structure represents the LsarSetInformationPolicy2 operation request
type SetInformationPolicy2Request struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// InformationClass: A parameter that specifies what type of information the caller
	// is setting.
	InformationClass PolicyInformationClass `idl:"name:InformationClass" json:"information_class"`
	// PolicyInformation: Data that represents policy being set.
	PolicyInformation *PolicyInformation `idl:"name:PolicyInformation;switch_is:InformationClass" json:"policy_information"`
}

func (o *SetInformationPolicy2Request) xxx_ToOp(ctx context.Context) *xxx_SetInformationPolicy2Operation {
	if o == nil {
		return &xxx_SetInformationPolicy2Operation{}
	}
	return &xxx_SetInformationPolicy2Operation{
		Policy:            o.Policy,
		InformationClass:  o.InformationClass,
		PolicyInformation: o.PolicyInformation,
	}
}

func (o *SetInformationPolicy2Request) xxx_FromOp(ctx context.Context, op *xxx_SetInformationPolicy2Operation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.InformationClass = op.InformationClass
	o.PolicyInformation = op.PolicyInformation
}
func (o *SetInformationPolicy2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetInformationPolicy2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationPolicy2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInformationPolicy2Response structure represents the LsarSetInformationPolicy2 operation response
type SetInformationPolicy2Response struct {
	// Return: The LsarSetInformationPolicy2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInformationPolicy2Response) xxx_ToOp(ctx context.Context) *xxx_SetInformationPolicy2Operation {
	if o == nil {
		return &xxx_SetInformationPolicy2Operation{}
	}
	return &xxx_SetInformationPolicy2Operation{
		Return: o.Return,
	}
}

func (o *SetInformationPolicy2Response) xxx_FromOp(ctx context.Context, op *xxx_SetInformationPolicy2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetInformationPolicy2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetInformationPolicy2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationPolicy2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryTrustedDomainInfoByNameOperation structure represents the LsarQueryTrustedDomainInfoByName operation
type xxx_QueryTrustedDomainInfoByNameOperation struct {
	Policy                   *Handle                 `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainName        *dtyp.UnicodeString     `idl:"name:TrustedDomainName" json:"trusted_domain_name"`
	InformationClass         TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
	TrustedDomainInformation *TrustedDomainInfo      `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
	Return                   int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryTrustedDomainInfoByNameOperation) OpNum() int { return 48 }

func (o *xxx_QueryTrustedDomainInfoByNameOperation) OpName() string {
	return "/lsarpc/v0/LsarQueryTrustedDomainInfoByName"
}

func (o *xxx_QueryTrustedDomainInfoByNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryTrustedDomainInfoByNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.TrustedDomainName != nil {
			if err := o.TrustedDomainName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryTrustedDomainInfoByNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.TrustedDomainName == nil {
			o.TrustedDomainName = &dtyp.UnicodeString{}
		}
		if err := o.TrustedDomainName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryTrustedDomainInfoByNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryTrustedDomainInfoByNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TrustedDomainInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO}*(1))(3:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		if o.TrustedDomainInformation != nil {
			_ptr_TrustedDomainInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swTrustedDomainInformation := uint16(o.InformationClass)
				if o.TrustedDomainInformation != nil {
					if err := o.TrustedDomainInformation.MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
						return err
					}
				} else {
					if err := (&TrustedDomainInfo{}).MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TrustedDomainInformation, _ptr_TrustedDomainInformation); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryTrustedDomainInfoByNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TrustedDomainInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO,pointer=ref}*(1))(3:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		_ptr_TrustedDomainInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TrustedDomainInformation == nil {
				o.TrustedDomainInformation = &TrustedDomainInfo{}
			}
			_swTrustedDomainInformation := uint16(o.InformationClass)
			if err := o.TrustedDomainInformation.UnmarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
				return err
			}
			return nil
		})
		_s_TrustedDomainInformation := func(ptr interface{}) { o.TrustedDomainInformation = *ptr.(**TrustedDomainInfo) }
		if err := w.ReadPointer(&o.TrustedDomainInformation, _s_TrustedDomainInformation, _ptr_TrustedDomainInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryTrustedDomainInfoByNameRequest structure represents the LsarQueryTrustedDomainInfoByName operation request
type QueryTrustedDomainInfoByNameRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainName: The name of the trusted domain object to query.
	TrustedDomainName *dtyp.UnicodeString `idl:"name:TrustedDomainName" json:"trusted_domain_name"`
	// InformationClass: One of the TRUSTED_INFORMATION_CLASS values identifying the type
	// of information the caller is interested in.
	InformationClass TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
}

func (o *QueryTrustedDomainInfoByNameRequest) xxx_ToOp(ctx context.Context) *xxx_QueryTrustedDomainInfoByNameOperation {
	if o == nil {
		return &xxx_QueryTrustedDomainInfoByNameOperation{}
	}
	return &xxx_QueryTrustedDomainInfoByNameOperation{
		Policy:            o.Policy,
		TrustedDomainName: o.TrustedDomainName,
		InformationClass:  o.InformationClass,
	}
}

func (o *QueryTrustedDomainInfoByNameRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryTrustedDomainInfoByNameOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainName = op.TrustedDomainName
	o.InformationClass = op.InformationClass
}
func (o *QueryTrustedDomainInfoByNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryTrustedDomainInfoByNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryTrustedDomainInfoByNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryTrustedDomainInfoByNameResponse structure represents the LsarQueryTrustedDomainInfoByName operation response
type QueryTrustedDomainInfoByNameResponse struct {
	// TrustedDomainInformation: Used to return the information requested by the caller.
	TrustedDomainInformation *TrustedDomainInfo `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
	// Return: The LsarQueryTrustedDomainInfoByName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryTrustedDomainInfoByNameResponse) xxx_ToOp(ctx context.Context) *xxx_QueryTrustedDomainInfoByNameOperation {
	if o == nil {
		return &xxx_QueryTrustedDomainInfoByNameOperation{}
	}
	return &xxx_QueryTrustedDomainInfoByNameOperation{
		TrustedDomainInformation: o.TrustedDomainInformation,
		Return:                   o.Return,
	}
}

func (o *QueryTrustedDomainInfoByNameResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryTrustedDomainInfoByNameOperation) {
	if o == nil {
		return
	}
	o.TrustedDomainInformation = op.TrustedDomainInformation
	o.Return = op.Return
}
func (o *QueryTrustedDomainInfoByNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryTrustedDomainInfoByNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryTrustedDomainInfoByNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetTrustedDomainInfoByNameOperation structure represents the LsarSetTrustedDomainInfoByName operation
type xxx_SetTrustedDomainInfoByNameOperation struct {
	Policy                   *Handle                 `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainName        *dtyp.UnicodeString     `idl:"name:TrustedDomainName" json:"trusted_domain_name"`
	InformationClass         TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
	TrustedDomainInformation *TrustedDomainInfo      `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
	Return                   int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTrustedDomainInfoByNameOperation) OpNum() int { return 49 }

func (o *xxx_SetTrustedDomainInfoByNameOperation) OpName() string {
	return "/lsarpc/v0/LsarSetTrustedDomainInfoByName"
}

func (o *xxx_SetTrustedDomainInfoByNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTrustedDomainInfoByNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.TrustedDomainName != nil {
			if err := o.TrustedDomainName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	// TrustedDomainInformation {in} (1:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO}*(1))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		_swTrustedDomainInformation := uint16(o.InformationClass)
		if o.TrustedDomainInformation != nil {
			if err := o.TrustedDomainInformation.MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInfo{}).MarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTrustedDomainInfoByNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.TrustedDomainName == nil {
			o.TrustedDomainName = &dtyp.UnicodeString{}
		}
		if err := o.TrustedDomainName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=TRUSTED_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	// TrustedDomainInformation {in} (1:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=PLSAPR_TRUSTED_DOMAIN_INFO,pointer=ref}*(1))(2:{switch_type={alias=TRUSTED_INFORMATION_CLASS}(enum), alias=LSAPR_TRUSTED_DOMAIN_INFO}(union))
	{
		if o.TrustedDomainInformation == nil {
			o.TrustedDomainInformation = &TrustedDomainInfo{}
		}
		_swTrustedDomainInformation := uint16(o.InformationClass)
		if err := o.TrustedDomainInformation.UnmarshalUnionNDR(ctx, w, _swTrustedDomainInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTrustedDomainInfoByNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTrustedDomainInfoByNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTrustedDomainInfoByNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetTrustedDomainInfoByNameRequest structure represents the LsarSetTrustedDomainInfoByName operation request
type SetTrustedDomainInfoByNameRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainName: The name of the trusted domain object to set information on.
	TrustedDomainName *dtyp.UnicodeString `idl:"name:TrustedDomainName" json:"trusted_domain_name"`
	// InformationClass: One of the TRUSTED_INFORMATION_CLASS values indicating the type
	// of information the caller is trying to set.
	InformationClass TrustedInformationClass `idl:"name:InformationClass" json:"information_class"`
	// TrustedDomainInformation: The data being set.
	TrustedDomainInformation *TrustedDomainInfo `idl:"name:TrustedDomainInformation;switch_is:InformationClass" json:"trusted_domain_information"`
}

func (o *SetTrustedDomainInfoByNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetTrustedDomainInfoByNameOperation {
	if o == nil {
		return &xxx_SetTrustedDomainInfoByNameOperation{}
	}
	return &xxx_SetTrustedDomainInfoByNameOperation{
		Policy:                   o.Policy,
		TrustedDomainName:        o.TrustedDomainName,
		InformationClass:         o.InformationClass,
		TrustedDomainInformation: o.TrustedDomainInformation,
	}
}

func (o *SetTrustedDomainInfoByNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTrustedDomainInfoByNameOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainName = op.TrustedDomainName
	o.InformationClass = op.InformationClass
	o.TrustedDomainInformation = op.TrustedDomainInformation
}
func (o *SetTrustedDomainInfoByNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetTrustedDomainInfoByNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTrustedDomainInfoByNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTrustedDomainInfoByNameResponse structure represents the LsarSetTrustedDomainInfoByName operation response
type SetTrustedDomainInfoByNameResponse struct {
	// Return: The LsarSetTrustedDomainInfoByName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetTrustedDomainInfoByNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetTrustedDomainInfoByNameOperation {
	if o == nil {
		return &xxx_SetTrustedDomainInfoByNameOperation{}
	}
	return &xxx_SetTrustedDomainInfoByNameOperation{
		Return: o.Return,
	}
}

func (o *SetTrustedDomainInfoByNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTrustedDomainInfoByNameOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetTrustedDomainInfoByNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetTrustedDomainInfoByNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTrustedDomainInfoByNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateTrustedDomainsExOperation structure represents the LsarEnumerateTrustedDomainsEx operation
type xxx_EnumerateTrustedDomainsExOperation struct {
	Policy                 *Handle              `idl:"name:PolicyHandle" json:"policy"`
	EnumerationContext     uint32               `idl:"name:EnumerationContext" json:"enumeration_context"`
	EnumerationBuffer      *TrustedEnumBufferEx `idl:"name:EnumerationBuffer" json:"enumeration_buffer"`
	PreferredMaximumLength uint32               `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
	Return                 int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateTrustedDomainsExOperation) OpNum() int { return 50 }

func (o *xxx_EnumerateTrustedDomainsExOperation) OpName() string {
	return "/lsarpc/v0/LsarEnumerateTrustedDomainsEx"
}

func (o *xxx_EnumerateTrustedDomainsExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateTrustedDomainsExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// PreferedMaximumLength {in} (1:(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateTrustedDomainsExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// PreferedMaximumLength {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateTrustedDomainsExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateTrustedDomainsExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// EnumerationBuffer {out} (1:{alias=PLSAPR_TRUSTED_ENUM_BUFFER_EX}*(1))(2:{alias=LSAPR_TRUSTED_ENUM_BUFFER_EX}(struct))
	{
		if o.EnumerationBuffer != nil {
			if err := o.EnumerationBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedEnumBufferEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateTrustedDomainsExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// EnumerationBuffer {out} (1:{alias=PLSAPR_TRUSTED_ENUM_BUFFER_EX,pointer=ref}*(1))(2:{alias=LSAPR_TRUSTED_ENUM_BUFFER_EX}(struct))
	{
		if o.EnumerationBuffer == nil {
			o.EnumerationBuffer = &TrustedEnumBufferEx{}
		}
		if err := o.EnumerationBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateTrustedDomainsExRequest structure represents the LsarEnumerateTrustedDomainsEx operation request
type EnumerateTrustedDomainsExRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// EnumerationContext: Used to keep track of the state of the enumeration in cases where
	// the caller obtains its information in several fragments.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// PreferedMaximumLength: A value that indicates the approximate size of the data to
	// be returned.
	PreferredMaximumLength uint32 `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
}

func (o *EnumerateTrustedDomainsExRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateTrustedDomainsExOperation {
	if o == nil {
		return &xxx_EnumerateTrustedDomainsExOperation{}
	}
	return &xxx_EnumerateTrustedDomainsExOperation{
		Policy:                 o.Policy,
		EnumerationContext:     o.EnumerationContext,
		PreferredMaximumLength: o.PreferredMaximumLength,
	}
}

func (o *EnumerateTrustedDomainsExRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateTrustedDomainsExOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.EnumerationContext = op.EnumerationContext
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *EnumerateTrustedDomainsExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateTrustedDomainsExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateTrustedDomainsExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateTrustedDomainsExResponse structure represents the LsarEnumerateTrustedDomainsEx operation response
type EnumerateTrustedDomainsExResponse struct {
	// EnumerationContext: Used to keep track of the state of the enumeration in cases where
	// the caller obtains its information in several fragments.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// EnumerationBuffer: Contains a fragment of requested information.
	EnumerationBuffer *TrustedEnumBufferEx `idl:"name:EnumerationBuffer" json:"enumeration_buffer"`
	// Return: The LsarEnumerateTrustedDomainsEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateTrustedDomainsExResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateTrustedDomainsExOperation {
	if o == nil {
		return &xxx_EnumerateTrustedDomainsExOperation{}
	}
	return &xxx_EnumerateTrustedDomainsExOperation{
		EnumerationContext: o.EnumerationContext,
		EnumerationBuffer:  o.EnumerationBuffer,
		Return:             o.Return,
	}
}

func (o *EnumerateTrustedDomainsExResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateTrustedDomainsExOperation) {
	if o == nil {
		return
	}
	o.EnumerationContext = op.EnumerationContext
	o.EnumerationBuffer = op.EnumerationBuffer
	o.Return = op.Return
}
func (o *EnumerateTrustedDomainsExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateTrustedDomainsExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateTrustedDomainsExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateTrustedDomainExOperation structure represents the LsarCreateTrustedDomainEx operation
type xxx_CreateTrustedDomainExOperation struct {
	Policy                    *Handle                       `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainInformation  *TrustedDomainInformationEx   `idl:"name:TrustedDomainInformation" json:"trusted_domain_information"`
	AuthenticationInformation *TrustedDomainAuthInformation `idl:"name:AuthenticationInformation" json:"authentication_information"`
	DesiredAccess             uint32                        `idl:"name:DesiredAccess" json:"desired_access"`
	TrustedDomain             *Handle                       `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	Return                    int32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateTrustedDomainExOperation) OpNum() int { return 51 }

func (o *xxx_CreateTrustedDomainExOperation) OpName() string {
	return "/lsarpc/v0/LsarCreateTrustedDomainEx"
}

func (o *xxx_CreateTrustedDomainExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainInformation {in} (1:{alias=PLSAPR_TRUSTED_DOMAIN_INFORMATION_EX}*(1))(2:{alias=LSAPR_TRUSTED_DOMAIN_INFORMATION_EX}(struct))
	{
		if o.TrustedDomainInformation != nil {
			if err := o.TrustedDomainInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInformationEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// AuthenticationInformation {in} (1:{alias=PLSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION}*(1))(2:{alias=LSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION}(struct))
	{
		if o.AuthenticationInformation != nil {
			if err := o.AuthenticationInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainAuthInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainInformation {in} (1:{alias=PLSAPR_TRUSTED_DOMAIN_INFORMATION_EX,pointer=ref}*(1))(2:{alias=LSAPR_TRUSTED_DOMAIN_INFORMATION_EX}(struct))
	{
		if o.TrustedDomainInformation == nil {
			o.TrustedDomainInformation = &TrustedDomainInformationEx{}
		}
		if err := o.TrustedDomainInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// AuthenticationInformation {in} (1:{alias=PLSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION,pointer=ref}*(1))(2:{alias=LSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION}(struct))
	{
		if o.AuthenticationInformation == nil {
			o.AuthenticationInformation = &TrustedDomainAuthInformation{}
		}
		if err := o.AuthenticationInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TrustedDomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain != nil {
			if err := o.TrustedDomain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TrustedDomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain == nil {
			o.TrustedDomain = &Handle{}
		}
		if err := o.TrustedDomain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateTrustedDomainExRequest structure represents the LsarCreateTrustedDomainEx operation request
type CreateTrustedDomainExRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainInformation: Information about the new TDO to be created.
	TrustedDomainInformation *TrustedDomainInformationEx `idl:"name:TrustedDomainInformation" json:"trusted_domain_information"`
	// AuthenticationInformation: Encrypted authentication information for the new TDO.
	AuthenticationInformation *TrustedDomainAuthInformation `idl:"name:AuthenticationInformation" json:"authentication_information"`
	// DesiredAccess: An access mask that specifies desired access to the TDO handle.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *CreateTrustedDomainExRequest) xxx_ToOp(ctx context.Context) *xxx_CreateTrustedDomainExOperation {
	if o == nil {
		return &xxx_CreateTrustedDomainExOperation{}
	}
	return &xxx_CreateTrustedDomainExOperation{
		Policy:                    o.Policy,
		TrustedDomainInformation:  o.TrustedDomainInformation,
		AuthenticationInformation: o.AuthenticationInformation,
		DesiredAccess:             o.DesiredAccess,
	}
}

func (o *CreateTrustedDomainExRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateTrustedDomainExOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainInformation = op.TrustedDomainInformation
	o.AuthenticationInformation = op.AuthenticationInformation
	o.DesiredAccess = op.DesiredAccess
}
func (o *CreateTrustedDomainExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateTrustedDomainExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateTrustedDomainExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateTrustedDomainExResponse structure represents the LsarCreateTrustedDomainEx operation response
type CreateTrustedDomainExResponse struct {
	// TrustedDomainHandle: Used to return the handle for the newly created TDO.
	TrustedDomain *Handle `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	// Return: The LsarCreateTrustedDomainEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateTrustedDomainExResponse) xxx_ToOp(ctx context.Context) *xxx_CreateTrustedDomainExOperation {
	if o == nil {
		return &xxx_CreateTrustedDomainExOperation{}
	}
	return &xxx_CreateTrustedDomainExOperation{
		TrustedDomain: o.TrustedDomain,
		Return:        o.Return,
	}
}

func (o *CreateTrustedDomainExResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateTrustedDomainExOperation) {
	if o == nil {
		return
	}
	o.TrustedDomain = op.TrustedDomain
	o.Return = op.Return
}
func (o *CreateTrustedDomainExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateTrustedDomainExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateTrustedDomainExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryDomainInformationPolicyOperation structure represents the LsarQueryDomainInformationPolicy operation
type xxx_QueryDomainInformationPolicyOperation struct {
	Policy                  *Handle                      `idl:"name:PolicyHandle" json:"policy"`
	InformationClass        PolicyDomainInformationClass `idl:"name:InformationClass" json:"information_class"`
	PolicyDomainInformation *PolicyDomainInformation     `idl:"name:PolicyDomainInformation;switch_is:InformationClass" json:"policy_domain_information"`
	Return                  int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryDomainInformationPolicyOperation) OpNum() int { return 53 }

func (o *xxx_QueryDomainInformationPolicyOperation) OpName() string {
	return "/lsarpc/v0/LsarQueryDomainInformationPolicy"
}

func (o *xxx_QueryDomainInformationPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDomainInformationPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InformationClass {in} (1:{alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDomainInformationPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDomainInformationPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDomainInformationPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PolicyDomainInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum), alias=PLSAPR_POLICY_DOMAIN_INFORMATION}*(1))(3:{switch_type={alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_DOMAIN_INFORMATION}(union))
	{
		if o.PolicyDomainInformation != nil {
			_ptr_PolicyDomainInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swPolicyDomainInformation := uint16(o.InformationClass)
				if o.PolicyDomainInformation != nil {
					if err := o.PolicyDomainInformation.MarshalUnionNDR(ctx, w, _swPolicyDomainInformation); err != nil {
						return err
					}
				} else {
					if err := (&PolicyDomainInformation{}).MarshalUnionNDR(ctx, w, _swPolicyDomainInformation); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyDomainInformation, _ptr_PolicyDomainInformation); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDomainInformationPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PolicyDomainInformation {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum), alias=PLSAPR_POLICY_DOMAIN_INFORMATION,pointer=ref}*(1))(3:{switch_type={alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_DOMAIN_INFORMATION}(union))
	{
		_ptr_PolicyDomainInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyDomainInformation == nil {
				o.PolicyDomainInformation = &PolicyDomainInformation{}
			}
			_swPolicyDomainInformation := uint16(o.InformationClass)
			if err := o.PolicyDomainInformation.UnmarshalUnionNDR(ctx, w, _swPolicyDomainInformation); err != nil {
				return err
			}
			return nil
		})
		_s_PolicyDomainInformation := func(ptr interface{}) { o.PolicyDomainInformation = *ptr.(**PolicyDomainInformation) }
		if err := w.ReadPointer(&o.PolicyDomainInformation, _s_PolicyDomainInformation, _ptr_PolicyDomainInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryDomainInformationPolicyRequest structure represents the LsarQueryDomainInformationPolicy operation request
type QueryDomainInformationPolicyRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// InformationClass: A parameter that specifies what type of information the caller
	// is requesting.
	InformationClass PolicyDomainInformationClass `idl:"name:InformationClass" json:"information_class"`
}

func (o *QueryDomainInformationPolicyRequest) xxx_ToOp(ctx context.Context) *xxx_QueryDomainInformationPolicyOperation {
	if o == nil {
		return &xxx_QueryDomainInformationPolicyOperation{}
	}
	return &xxx_QueryDomainInformationPolicyOperation{
		Policy:           o.Policy,
		InformationClass: o.InformationClass,
	}
}

func (o *QueryDomainInformationPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryDomainInformationPolicyOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.InformationClass = op.InformationClass
}
func (o *QueryDomainInformationPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryDomainInformationPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDomainInformationPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryDomainInformationPolicyResponse structure represents the LsarQueryDomainInformationPolicy operation response
type QueryDomainInformationPolicyResponse struct {
	// PolicyDomainInformation: A parameter that references policy information structure
	// on return.
	PolicyDomainInformation *PolicyDomainInformation `idl:"name:PolicyDomainInformation;switch_is:InformationClass" json:"policy_domain_information"`
	// Return: The LsarQueryDomainInformationPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryDomainInformationPolicyResponse) xxx_ToOp(ctx context.Context) *xxx_QueryDomainInformationPolicyOperation {
	if o == nil {
		return &xxx_QueryDomainInformationPolicyOperation{}
	}
	return &xxx_QueryDomainInformationPolicyOperation{
		PolicyDomainInformation: o.PolicyDomainInformation,
		Return:                  o.Return,
	}
}

func (o *QueryDomainInformationPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryDomainInformationPolicyOperation) {
	if o == nil {
		return
	}
	o.PolicyDomainInformation = op.PolicyDomainInformation
	o.Return = op.Return
}
func (o *QueryDomainInformationPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryDomainInformationPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDomainInformationPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDomainInformationPolicyOperation structure represents the LsarSetDomainInformationPolicy operation
type xxx_SetDomainInformationPolicyOperation struct {
	Policy                  *Handle                      `idl:"name:PolicyHandle" json:"policy"`
	InformationClass        PolicyDomainInformationClass `idl:"name:InformationClass" json:"information_class"`
	PolicyDomainInformation *PolicyDomainInformation     `idl:"name:PolicyDomainInformation;switch_is:InformationClass;pointer:unique" json:"policy_domain_information"`
	Return                  int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDomainInformationPolicyOperation) OpNum() int { return 54 }

func (o *xxx_SetDomainInformationPolicyOperation) OpName() string {
	return "/lsarpc/v0/LsarSetDomainInformationPolicy"
}

func (o *xxx_SetDomainInformationPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDomainInformationPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InformationClass {in} (1:{alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InformationClass)); err != nil {
			return err
		}
	}
	// PolicyDomainInformation {in} (1:{switch_type={alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum), pointer=unique, alias=PLSAPR_POLICY_DOMAIN_INFORMATION}*(1))(2:{switch_type={alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_DOMAIN_INFORMATION}(union))
	{
		if o.PolicyDomainInformation != nil {
			_ptr_PolicyDomainInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swPolicyDomainInformation := uint16(o.InformationClass)
				if o.PolicyDomainInformation != nil {
					if err := o.PolicyDomainInformation.MarshalUnionNDR(ctx, w, _swPolicyDomainInformation); err != nil {
						return err
					}
				} else {
					if err := (&PolicyDomainInformation{}).MarshalUnionNDR(ctx, w, _swPolicyDomainInformation); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PolicyDomainInformation, _ptr_PolicyDomainInformation); err != nil {
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

func (o *xxx_SetDomainInformationPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InformationClass {in} (1:{alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InformationClass)); err != nil {
			return err
		}
	}
	// PolicyDomainInformation {in} (1:{switch_type={alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum), pointer=unique, alias=PLSAPR_POLICY_DOMAIN_INFORMATION}*(1))(2:{switch_type={alias=POLICY_DOMAIN_INFORMATION_CLASS}(enum), alias=LSAPR_POLICY_DOMAIN_INFORMATION}(union))
	{
		_ptr_PolicyDomainInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PolicyDomainInformation == nil {
				o.PolicyDomainInformation = &PolicyDomainInformation{}
			}
			_swPolicyDomainInformation := uint16(o.InformationClass)
			if err := o.PolicyDomainInformation.UnmarshalUnionNDR(ctx, w, _swPolicyDomainInformation); err != nil {
				return err
			}
			return nil
		})
		_s_PolicyDomainInformation := func(ptr interface{}) { o.PolicyDomainInformation = *ptr.(**PolicyDomainInformation) }
		if err := w.ReadPointer(&o.PolicyDomainInformation, _s_PolicyDomainInformation, _ptr_PolicyDomainInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDomainInformationPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDomainInformationPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDomainInformationPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetDomainInformationPolicyRequest structure represents the LsarSetDomainInformationPolicy operation request
type SetDomainInformationPolicyRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// InformationClass: A parameter that specifies what type of information the caller
	// is setting.
	InformationClass PolicyDomainInformationClass `idl:"name:InformationClass" json:"information_class"`
	// PolicyDomainInformation: Data representing policy being set.
	PolicyDomainInformation *PolicyDomainInformation `idl:"name:PolicyDomainInformation;switch_is:InformationClass;pointer:unique" json:"policy_domain_information"`
}

func (o *SetDomainInformationPolicyRequest) xxx_ToOp(ctx context.Context) *xxx_SetDomainInformationPolicyOperation {
	if o == nil {
		return &xxx_SetDomainInformationPolicyOperation{}
	}
	return &xxx_SetDomainInformationPolicyOperation{
		Policy:                  o.Policy,
		InformationClass:        o.InformationClass,
		PolicyDomainInformation: o.PolicyDomainInformation,
	}
}

func (o *SetDomainInformationPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDomainInformationPolicyOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.InformationClass = op.InformationClass
	o.PolicyDomainInformation = op.PolicyDomainInformation
}
func (o *SetDomainInformationPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDomainInformationPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDomainInformationPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDomainInformationPolicyResponse structure represents the LsarSetDomainInformationPolicy operation response
type SetDomainInformationPolicyResponse struct {
	// Return: The LsarSetDomainInformationPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDomainInformationPolicyResponse) xxx_ToOp(ctx context.Context) *xxx_SetDomainInformationPolicyOperation {
	if o == nil {
		return &xxx_SetDomainInformationPolicyOperation{}
	}
	return &xxx_SetDomainInformationPolicyOperation{
		Return: o.Return,
	}
}

func (o *SetDomainInformationPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDomainInformationPolicyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetDomainInformationPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDomainInformationPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDomainInformationPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenTrustedDomainByNameOperation structure represents the LsarOpenTrustedDomainByName operation
type xxx_OpenTrustedDomainByNameOperation struct {
	Policy            *Handle             `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainName *dtyp.UnicodeString `idl:"name:TrustedDomainName" json:"trusted_domain_name"`
	DesiredAccess     uint32              `idl:"name:DesiredAccess" json:"desired_access"`
	TrustedDomain     *Handle             `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	Return            int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenTrustedDomainByNameOperation) OpNum() int { return 55 }

func (o *xxx_OpenTrustedDomainByNameOperation) OpName() string {
	return "/lsarpc/v0/LsarOpenTrustedDomainByName"
}

func (o *xxx_OpenTrustedDomainByNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenTrustedDomainByNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.TrustedDomainName != nil {
			if err := o.TrustedDomainName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenTrustedDomainByNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.TrustedDomainName == nil {
			o.TrustedDomainName = &dtyp.UnicodeString{}
		}
		if err := o.TrustedDomainName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenTrustedDomainByNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenTrustedDomainByNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TrustedDomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain != nil {
			if err := o.TrustedDomain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenTrustedDomainByNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TrustedDomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain == nil {
			o.TrustedDomain = &Handle{}
		}
		if err := o.TrustedDomain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenTrustedDomainByNameRequest structure represents the LsarOpenTrustedDomainByName operation request
type OpenTrustedDomainByNameRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainName: The name of the trusted domain object.
	TrustedDomainName *dtyp.UnicodeString `idl:"name:TrustedDomainName" json:"trusted_domain_name"`
	// DesiredAccess: The type of access requested by the caller.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *OpenTrustedDomainByNameRequest) xxx_ToOp(ctx context.Context) *xxx_OpenTrustedDomainByNameOperation {
	if o == nil {
		return &xxx_OpenTrustedDomainByNameOperation{}
	}
	return &xxx_OpenTrustedDomainByNameOperation{
		Policy:            o.Policy,
		TrustedDomainName: o.TrustedDomainName,
		DesiredAccess:     o.DesiredAccess,
	}
}

func (o *OpenTrustedDomainByNameRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenTrustedDomainByNameOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainName = op.TrustedDomainName
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenTrustedDomainByNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenTrustedDomainByNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenTrustedDomainByNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenTrustedDomainByNameResponse structure represents the LsarOpenTrustedDomainByName operation response
type OpenTrustedDomainByNameResponse struct {
	// TrustedDomainHandle: Used to return the opened trusted domain handle.
	TrustedDomain *Handle `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	// Return: The LsarOpenTrustedDomainByName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenTrustedDomainByNameResponse) xxx_ToOp(ctx context.Context) *xxx_OpenTrustedDomainByNameOperation {
	if o == nil {
		return &xxx_OpenTrustedDomainByNameOperation{}
	}
	return &xxx_OpenTrustedDomainByNameOperation{
		TrustedDomain: o.TrustedDomain,
		Return:        o.Return,
	}
}

func (o *OpenTrustedDomainByNameResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenTrustedDomainByNameOperation) {
	if o == nil {
		return
	}
	o.TrustedDomain = op.TrustedDomain
	o.Return = op.Return
}
func (o *OpenTrustedDomainByNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenTrustedDomainByNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenTrustedDomainByNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateTrustedDomainEx2Operation structure represents the LsarCreateTrustedDomainEx2 operation
type xxx_CreateTrustedDomainEx2Operation struct {
	Policy                    *Handle                               `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainInformation  *TrustedDomainInformationEx           `idl:"name:TrustedDomainInformation" json:"trusted_domain_information"`
	AuthenticationInformation *TrustedDomainAuthInformationInternal `idl:"name:AuthenticationInformation" json:"authentication_information"`
	DesiredAccess             uint32                                `idl:"name:DesiredAccess" json:"desired_access"`
	TrustedDomain             *Handle                               `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	Return                    int32                                 `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateTrustedDomainEx2Operation) OpNum() int { return 59 }

func (o *xxx_CreateTrustedDomainEx2Operation) OpName() string {
	return "/lsarpc/v0/LsarCreateTrustedDomainEx2"
}

func (o *xxx_CreateTrustedDomainEx2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainEx2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainInformation {in} (1:{alias=PLSAPR_TRUSTED_DOMAIN_INFORMATION_EX}*(1))(2:{alias=LSAPR_TRUSTED_DOMAIN_INFORMATION_EX}(struct))
	{
		if o.TrustedDomainInformation != nil {
			if err := o.TrustedDomainInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainInformationEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// AuthenticationInformation {in} (1:{alias=PLSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION_INTERNAL}*(1))(2:{alias=LSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION_INTERNAL}(struct))
	{
		if o.AuthenticationInformation != nil {
			if err := o.AuthenticationInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TrustedDomainAuthInformationInternal{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainEx2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainInformation {in} (1:{alias=PLSAPR_TRUSTED_DOMAIN_INFORMATION_EX,pointer=ref}*(1))(2:{alias=LSAPR_TRUSTED_DOMAIN_INFORMATION_EX}(struct))
	{
		if o.TrustedDomainInformation == nil {
			o.TrustedDomainInformation = &TrustedDomainInformationEx{}
		}
		if err := o.TrustedDomainInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// AuthenticationInformation {in} (1:{alias=PLSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION_INTERNAL,pointer=ref}*(1))(2:{alias=LSAPR_TRUSTED_DOMAIN_AUTH_INFORMATION_INTERNAL}(struct))
	{
		if o.AuthenticationInformation == nil {
			o.AuthenticationInformation = &TrustedDomainAuthInformationInternal{}
		}
		if err := o.AuthenticationInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainEx2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainEx2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TrustedDomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain != nil {
			if err := o.TrustedDomain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTrustedDomainEx2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TrustedDomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.TrustedDomain == nil {
			o.TrustedDomain = &Handle{}
		}
		if err := o.TrustedDomain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateTrustedDomainEx2Request structure represents the LsarCreateTrustedDomainEx2 operation request
type CreateTrustedDomainEx2Request struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainInformation: Information about the new TDO to be created.
	TrustedDomainInformation *TrustedDomainInformationEx `idl:"name:TrustedDomainInformation" json:"trusted_domain_information"`
	// AuthenticationInformation: Encrypted authentication information for the new TDO.
	AuthenticationInformation *TrustedDomainAuthInformationInternal `idl:"name:AuthenticationInformation" json:"authentication_information"`
	// DesiredAccess: An access mask specifying desired access to the TDO handle.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *CreateTrustedDomainEx2Request) xxx_ToOp(ctx context.Context) *xxx_CreateTrustedDomainEx2Operation {
	if o == nil {
		return &xxx_CreateTrustedDomainEx2Operation{}
	}
	return &xxx_CreateTrustedDomainEx2Operation{
		Policy:                    o.Policy,
		TrustedDomainInformation:  o.TrustedDomainInformation,
		AuthenticationInformation: o.AuthenticationInformation,
		DesiredAccess:             o.DesiredAccess,
	}
}

func (o *CreateTrustedDomainEx2Request) xxx_FromOp(ctx context.Context, op *xxx_CreateTrustedDomainEx2Operation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainInformation = op.TrustedDomainInformation
	o.AuthenticationInformation = op.AuthenticationInformation
	o.DesiredAccess = op.DesiredAccess
}
func (o *CreateTrustedDomainEx2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateTrustedDomainEx2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateTrustedDomainEx2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateTrustedDomainEx2Response structure represents the LsarCreateTrustedDomainEx2 operation response
type CreateTrustedDomainEx2Response struct {
	// TrustedDomainHandle: Used to return the handle for the newly created TDO.
	TrustedDomain *Handle `idl:"name:TrustedDomainHandle" json:"trusted_domain"`
	// Return: The LsarCreateTrustedDomainEx2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateTrustedDomainEx2Response) xxx_ToOp(ctx context.Context) *xxx_CreateTrustedDomainEx2Operation {
	if o == nil {
		return &xxx_CreateTrustedDomainEx2Operation{}
	}
	return &xxx_CreateTrustedDomainEx2Operation{
		TrustedDomain: o.TrustedDomain,
		Return:        o.Return,
	}
}

func (o *CreateTrustedDomainEx2Response) xxx_FromOp(ctx context.Context, op *xxx_CreateTrustedDomainEx2Operation) {
	if o == nil {
		return
	}
	o.TrustedDomain = op.TrustedDomain
	o.Return = op.Return
}
func (o *CreateTrustedDomainEx2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateTrustedDomainEx2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateTrustedDomainEx2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryForestTrustInformationOperation structure represents the LsarQueryForestTrustInformation operation
type xxx_QueryForestTrustInformationOperation struct {
	Policy            *Handle                 `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainName *UnicodeString          `idl:"name:TrustedDomainName" json:"trusted_domain_name"`
	HighestRecordType ForestTrustRecordType   `idl:"name:HighestRecordType" json:"highest_record_type"`
	ForestTrustInfo   *ForestTrustInformation `idl:"name:ForestTrustInfo" json:"forest_trust_info"`
	Return            int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryForestTrustInformationOperation) OpNum() int { return 73 }

func (o *xxx_QueryForestTrustInformationOperation) OpName() string {
	return "/lsarpc/v0/LsarQueryForestTrustInformation"
}

func (o *xxx_QueryForestTrustInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryForestTrustInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainName {in} (1:{alias=PLSA_UNICODE_STRING}*(1))(2:{alias=LSA_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.TrustedDomainName != nil {
			if err := o.TrustedDomainName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// HighestRecordType {in} (1:{alias=LSA_FOREST_TRUST_RECORD_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.HighestRecordType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryForestTrustInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainName {in} (1:{alias=PLSA_UNICODE_STRING,pointer=ref}*(1))(2:{alias=LSA_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.TrustedDomainName == nil {
			o.TrustedDomainName = &UnicodeString{}
		}
		if err := o.TrustedDomainName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// HighestRecordType {in} (1:{alias=LSA_FOREST_TRUST_RECORD_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.HighestRecordType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryForestTrustInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryForestTrustInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ForestTrustInfo {out} (1:{pointer=ref}*(2))(2:{alias=PLSA_FOREST_TRUST_INFORMATION}*(1))(3:{alias=LSA_FOREST_TRUST_INFORMATION}(struct))
	{
		if o.ForestTrustInfo != nil {
			_ptr_ForestTrustInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ForestTrustInfo != nil {
					if err := o.ForestTrustInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ForestTrustInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ForestTrustInfo, _ptr_ForestTrustInfo); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryForestTrustInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ForestTrustInfo {out} (1:{pointer=ref}*(2))(2:{alias=PLSA_FOREST_TRUST_INFORMATION,pointer=ref}*(1))(3:{alias=LSA_FOREST_TRUST_INFORMATION}(struct))
	{
		_ptr_ForestTrustInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ForestTrustInfo == nil {
				o.ForestTrustInfo = &ForestTrustInformation{}
			}
			if err := o.ForestTrustInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ForestTrustInfo := func(ptr interface{}) { o.ForestTrustInfo = *ptr.(**ForestTrustInformation) }
		if err := w.ReadPointer(&o.ForestTrustInfo, _s_ForestTrustInfo, _ptr_ForestTrustInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryForestTrustInformationRequest structure represents the LsarQueryForestTrustInformation operation request
type QueryForestTrustInformationRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainName: The name of the trusted domain to query.
	TrustedDomainName *UnicodeString `idl:"name:TrustedDomainName" json:"trusted_domain_name"`
	// HighestRecordType: The highest ordinal number of forest trust record type that the
	// caller understands.
	HighestRecordType ForestTrustRecordType `idl:"name:HighestRecordType" json:"highest_record_type"`
}

func (o *QueryForestTrustInformationRequest) xxx_ToOp(ctx context.Context) *xxx_QueryForestTrustInformationOperation {
	if o == nil {
		return &xxx_QueryForestTrustInformationOperation{}
	}
	return &xxx_QueryForestTrustInformationOperation{
		Policy:            o.Policy,
		TrustedDomainName: o.TrustedDomainName,
		HighestRecordType: o.HighestRecordType,
	}
}

func (o *QueryForestTrustInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryForestTrustInformationOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainName = op.TrustedDomainName
	o.HighestRecordType = op.HighestRecordType
}
func (o *QueryForestTrustInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryForestTrustInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryForestTrustInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryForestTrustInformationResponse structure represents the LsarQueryForestTrustInformation operation response
type QueryForestTrustInformationResponse struct {
	// ForestTrustInfo: Used to return the forest trust information.
	ForestTrustInfo *ForestTrustInformation `idl:"name:ForestTrustInfo" json:"forest_trust_info"`
	// Return: The LsarQueryForestTrustInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryForestTrustInformationResponse) xxx_ToOp(ctx context.Context) *xxx_QueryForestTrustInformationOperation {
	if o == nil {
		return &xxx_QueryForestTrustInformationOperation{}
	}
	return &xxx_QueryForestTrustInformationOperation{
		ForestTrustInfo: o.ForestTrustInfo,
		Return:          o.Return,
	}
}

func (o *QueryForestTrustInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryForestTrustInformationOperation) {
	if o == nil {
		return
	}
	o.ForestTrustInfo = op.ForestTrustInfo
	o.Return = op.Return
}
func (o *QueryForestTrustInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryForestTrustInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryForestTrustInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetForestTrustInformationOperation structure represents the LsarSetForestTrustInformation operation
type xxx_SetForestTrustInformationOperation struct {
	Policy            *Handle                          `idl:"name:PolicyHandle" json:"policy"`
	TrustedDomainName *UnicodeString                   `idl:"name:TrustedDomainName" json:"trusted_domain_name"`
	HighestRecordType ForestTrustRecordType            `idl:"name:HighestRecordType" json:"highest_record_type"`
	ForestTrustInfo   *ForestTrustInformation          `idl:"name:ForestTrustInfo" json:"forest_trust_info"`
	CheckOnly         uint8                            `idl:"name:CheckOnly" json:"check_only"`
	CollisionInfo     *ForestTrustCollisionInformation `idl:"name:CollisionInfo" json:"collision_info"`
	Return            int32                            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetForestTrustInformationOperation) OpNum() int { return 74 }

func (o *xxx_SetForestTrustInformationOperation) OpName() string {
	return "/lsarpc/v0/LsarSetForestTrustInformation"
}

func (o *xxx_SetForestTrustInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetForestTrustInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrustedDomainName {in} (1:{alias=PLSA_UNICODE_STRING}*(1))(2:{alias=LSA_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.TrustedDomainName != nil {
			if err := o.TrustedDomainName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// HighestRecordType {in} (1:{alias=LSA_FOREST_TRUST_RECORD_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.HighestRecordType)); err != nil {
			return err
		}
	}
	// ForestTrustInfo {in} (1:{alias=PLSA_FOREST_TRUST_INFORMATION}*(1))(2:{alias=LSA_FOREST_TRUST_INFORMATION}(struct))
	{
		if o.ForestTrustInfo != nil {
			if err := o.ForestTrustInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ForestTrustInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// CheckOnly {in} (1:(uchar))
	{
		if err := w.WriteData(o.CheckOnly); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetForestTrustInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrustedDomainName {in} (1:{alias=PLSA_UNICODE_STRING,pointer=ref}*(1))(2:{alias=LSA_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.TrustedDomainName == nil {
			o.TrustedDomainName = &UnicodeString{}
		}
		if err := o.TrustedDomainName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// HighestRecordType {in} (1:{alias=LSA_FOREST_TRUST_RECORD_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.HighestRecordType)); err != nil {
			return err
		}
	}
	// ForestTrustInfo {in} (1:{alias=PLSA_FOREST_TRUST_INFORMATION,pointer=ref}*(1))(2:{alias=LSA_FOREST_TRUST_INFORMATION}(struct))
	{
		if o.ForestTrustInfo == nil {
			o.ForestTrustInfo = &ForestTrustInformation{}
		}
		if err := o.ForestTrustInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CheckOnly {in} (1:(uchar))
	{
		if err := w.ReadData(&o.CheckOnly); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetForestTrustInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetForestTrustInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// CollisionInfo {out} (1:{pointer=ref}*(2))(2:{alias=PLSA_FOREST_TRUST_COLLISION_INFORMATION}*(1))(3:{alias=LSA_FOREST_TRUST_COLLISION_INFORMATION}(struct))
	{
		if o.CollisionInfo != nil {
			_ptr_CollisionInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CollisionInfo != nil {
					if err := o.CollisionInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ForestTrustCollisionInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CollisionInfo, _ptr_CollisionInfo); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetForestTrustInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// CollisionInfo {out} (1:{pointer=ref}*(2))(2:{alias=PLSA_FOREST_TRUST_COLLISION_INFORMATION,pointer=ref}*(1))(3:{alias=LSA_FOREST_TRUST_COLLISION_INFORMATION}(struct))
	{
		_ptr_CollisionInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CollisionInfo == nil {
				o.CollisionInfo = &ForestTrustCollisionInformation{}
			}
			if err := o.CollisionInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_CollisionInfo := func(ptr interface{}) { o.CollisionInfo = *ptr.(**ForestTrustCollisionInformation) }
		if err := w.ReadPointer(&o.CollisionInfo, _s_CollisionInfo, _ptr_CollisionInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetForestTrustInformationRequest structure represents the LsarSetForestTrustInformation operation request
type SetForestTrustInformationRequest struct {
	// PolicyHandle: An RPC context handle obtained from either LsarOpenPolicy or LsarOpenPolicy2.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// TrustedDomainName: The name of the trusted domain object on which to set the forest
	// trust information.
	TrustedDomainName *UnicodeString `idl:"name:TrustedDomainName" json:"trusted_domain_name"`
	// HighestRecordType: The highest ordinal forest trust record type that the caller understands.
	HighestRecordType ForestTrustRecordType `idl:"name:HighestRecordType" json:"highest_record_type"`
	// ForestTrustInfo: The forest trust information that the caller is trying to set on
	// the trusted domain object.
	ForestTrustInfo *ForestTrustInformation `idl:"name:ForestTrustInfo" json:"forest_trust_info"`
	// CheckOnly: If not 0, the operation is read-only and does not alter the state of the
	// server's database.
	CheckOnly uint8 `idl:"name:CheckOnly" json:"check_only"`
}

func (o *SetForestTrustInformationRequest) xxx_ToOp(ctx context.Context) *xxx_SetForestTrustInformationOperation {
	if o == nil {
		return &xxx_SetForestTrustInformationOperation{}
	}
	return &xxx_SetForestTrustInformationOperation{
		Policy:            o.Policy,
		TrustedDomainName: o.TrustedDomainName,
		HighestRecordType: o.HighestRecordType,
		ForestTrustInfo:   o.ForestTrustInfo,
		CheckOnly:         o.CheckOnly,
	}
}

func (o *SetForestTrustInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetForestTrustInformationOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.TrustedDomainName = op.TrustedDomainName
	o.HighestRecordType = op.HighestRecordType
	o.ForestTrustInfo = op.ForestTrustInfo
	o.CheckOnly = op.CheckOnly
}
func (o *SetForestTrustInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetForestTrustInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetForestTrustInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetForestTrustInformationResponse structure represents the LsarSetForestTrustInformation operation response
type SetForestTrustInformationResponse struct {
	// CollisionInfo: Used to return information about collisions between different sets
	// of forest trust information in the server's database.
	CollisionInfo *ForestTrustCollisionInformation `idl:"name:CollisionInfo" json:"collision_info"`
	// Return: The LsarSetForestTrustInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetForestTrustInformationResponse) xxx_ToOp(ctx context.Context) *xxx_SetForestTrustInformationOperation {
	if o == nil {
		return &xxx_SetForestTrustInformationOperation{}
	}
	return &xxx_SetForestTrustInformationOperation{
		CollisionInfo: o.CollisionInfo,
		Return:        o.Return,
	}
}

func (o *SetForestTrustInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetForestTrustInformationOperation) {
	if o == nil {
		return
	}
	o.CollisionInfo = op.CollisionInfo
	o.Return = op.Return
}
func (o *SetForestTrustInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetForestTrustInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetForestTrustInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
