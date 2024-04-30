package lsarpc

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

// lsarpc server interface.
type LsarpcServer interface {

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
	Close(context.Context, *CloseRequest) (*CloseResponse, error)

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
	EnumeratePrivileges(context.Context, *EnumeratePrivilegesRequest) (*EnumeratePrivilegesResponse, error)

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
	QuerySecurityObject(context.Context, *QuerySecurityObjectRequest) (*QuerySecurityObjectResponse, error)

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
	SetSecurityObject(context.Context, *SetSecurityObjectRequest) (*SetSecurityObjectResponse, error)

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// The LsarOpenPolicy method is exactly the same as LsarOpenPolicy2, except that the
	// SystemName parameter in this function, because of its syntactic definition, contains
	// only one character instead of a full string. This SystemName parameter does not have
	// any effect on message processing in any environment. It MUST be ignored.
	OpenPolicy(context.Context, *OpenPolicyRequest) (*OpenPolicyResponse, error)

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
	QueryInformationPolicy(context.Context, *QueryInformationPolicyRequest) (*QueryInformationPolicyResponse, error)

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
	SetInformationPolicy(context.Context, *SetInformationPolicyRequest) (*SetInformationPolicyResponse, error)

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
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)

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
	EnumerateAccounts(context.Context, *EnumerateAccountsRequest) (*EnumerateAccountsResponse, error)

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
	CreateTrustedDomain(context.Context, *CreateTrustedDomainRequest) (*CreateTrustedDomainResponse, error)

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
	EnumerateTrustedDomains(context.Context, *EnumerateTrustedDomainsRequest) (*EnumerateTrustedDomainsResponse, error)

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
	CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error)

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
	OpenAccount(context.Context, *OpenAccountRequest) (*OpenAccountResponse, error)

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
	EnumeratePrivilegesAccount(context.Context, *EnumeratePrivilegesAccountRequest) (*EnumeratePrivilegesAccountResponse, error)

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
	AddPrivilegesToAccount(context.Context, *AddPrivilegesToAccountRequest) (*AddPrivilegesToAccountResponse, error)

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
	RemovePrivilegesFromAccount(context.Context, *RemovePrivilegesFromAccountRequest) (*RemovePrivilegesFromAccountResponse, error)

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
	GetSystemAccessAccount(context.Context, *GetSystemAccessAccountRequest) (*GetSystemAccessAccountResponse, error)

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
	SetSystemAccessAccount(context.Context, *SetSystemAccessAccountRequest) (*SetSystemAccessAccountResponse, error)

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
	OpenTrustedDomain(context.Context, *OpenTrustedDomainRequest) (*OpenTrustedDomainResponse, error)

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
	QueryInfoTrustedDomain(context.Context, *QueryInfoTrustedDomainRequest) (*QueryInfoTrustedDomainResponse, error)

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
	SetInformationTrustedDomain(context.Context, *SetInformationTrustedDomainRequest) (*SetInformationTrustedDomainResponse, error)

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
	OpenSecret(context.Context, *OpenSecretRequest) (*OpenSecretResponse, error)

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
	SetSecret(context.Context, *SetSecretRequest) (*SetSecretResponse, error)

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
	QuerySecret(context.Context, *QuerySecretRequest) (*QuerySecretResponse, error)

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
	LookupPrivilegeValue(context.Context, *LookupPrivilegeValueRequest) (*LookupPrivilegeValueResponse, error)

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
	LookupPrivilegeName(context.Context, *LookupPrivilegeNameRequest) (*LookupPrivilegeNameResponse, error)

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
	LookupPrivilegeDisplayName(context.Context, *LookupPrivilegeDisplayNameRequest) (*LookupPrivilegeDisplayNameResponse, error)

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
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)

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
	EnumerateAccountsWithUserRight(context.Context, *EnumerateAccountsWithUserRightRequest) (*EnumerateAccountsWithUserRightResponse, error)

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
	EnumerateAccountRights(context.Context, *EnumerateAccountRightsRequest) (*EnumerateAccountRightsResponse, error)

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
	AddAccountRights(context.Context, *AddAccountRightsRequest) (*AddAccountRightsResponse, error)

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
	RemoveAccountRights(context.Context, *RemoveAccountRightsRequest) (*RemoveAccountRightsResponse, error)

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
	QueryTrustedDomainInfo(context.Context, *QueryTrustedDomainInfoRequest) (*QueryTrustedDomainInfoResponse, error)

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
	SetTrustedDomainInfo(context.Context, *SetTrustedDomainInfoRequest) (*SetTrustedDomainInfoResponse, error)

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
	DeleteTrustedDomain(context.Context, *DeleteTrustedDomainRequest) (*DeleteTrustedDomainResponse, error)

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
	StorePrivateData(context.Context, *StorePrivateDataRequest) (*StorePrivateDataResponse, error)

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
	RetrievePrivateData(context.Context, *RetrievePrivateDataRequest) (*RetrievePrivateDataResponse, error)

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
	OpenPolicy2(context.Context, *OpenPolicy2Request) (*OpenPolicy2Response, error)

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
	QueryInformationPolicy2(context.Context, *QueryInformationPolicy2Request) (*QueryInformationPolicy2Response, error)

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
	SetInformationPolicy2(context.Context, *SetInformationPolicy2Request) (*SetInformationPolicy2Response, error)

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
	QueryTrustedDomainInfoByName(context.Context, *QueryTrustedDomainInfoByNameRequest) (*QueryTrustedDomainInfoByNameResponse, error)

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
	SetTrustedDomainInfoByName(context.Context, *SetTrustedDomainInfoByNameRequest) (*SetTrustedDomainInfoByNameResponse, error)

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
	EnumerateTrustedDomainsEx(context.Context, *EnumerateTrustedDomainsExRequest) (*EnumerateTrustedDomainsExResponse, error)

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
	CreateTrustedDomainEx(context.Context, *CreateTrustedDomainExRequest) (*CreateTrustedDomainExResponse, error)

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
	QueryDomainInformationPolicy(context.Context, *QueryDomainInformationPolicyRequest) (*QueryDomainInformationPolicyResponse, error)

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
	SetDomainInformationPolicy(context.Context, *SetDomainInformationPolicyRequest) (*SetDomainInformationPolicyResponse, error)

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
	OpenTrustedDomainByName(context.Context, *OpenTrustedDomainByNameRequest) (*OpenTrustedDomainByNameResponse, error)

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
	CreateTrustedDomainEx2(context.Context, *CreateTrustedDomainEx2Request) (*CreateTrustedDomainEx2Response, error)

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
	QueryForestTrustInformation(context.Context, *QueryForestTrustInformationRequest) (*QueryForestTrustInformationResponse, error)

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
	SetForestTrustInformation(context.Context, *SetForestTrustInformationRequest) (*SetForestTrustInformationResponse, error)
}

func RegisterLsarpcServer(conn dcerpc.Conn, o LsarpcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewLsarpcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(LsarpcSyntaxV0_0))...)
}

func NewLsarpcServerHandle(o LsarpcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return LsarpcServerHandle(ctx, o, opNum, r)
	}
}

func LsarpcServerHandle(ctx context.Context, o LsarpcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // LsarClose
		in := &CloseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Close(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // Opnum1NotUsedOnWire
		// Opnum1NotUsedOnWire
		return nil, nil
	case 2: // LsarEnumeratePrivileges
		in := &EnumeratePrivilegesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumeratePrivileges(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // LsarQuerySecurityObject
		in := &QuerySecurityObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QuerySecurityObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // LsarSetSecurityObject
		in := &SetSecurityObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSecurityObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // LsarOpenPolicy
		in := &OpenPolicyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenPolicy(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // LsarQueryInformationPolicy
		in := &QueryInformationPolicyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInformationPolicy(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // LsarSetInformationPolicy
		in := &SetInformationPolicyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetInformationPolicy(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 10: // LsarCreateAccount
		in := &CreateAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // LsarEnumerateAccounts
		in := &EnumerateAccountsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateAccounts(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // LsarCreateTrustedDomain
		in := &CreateTrustedDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateTrustedDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // LsarEnumerateTrustedDomains
		in := &EnumerateTrustedDomainsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateTrustedDomains(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // Lsar_LSA_TM_14
		// Lsar_LSA_TM_14
		return nil, nil
	case 15: // Lsar_LSA_TM_15
		// Lsar_LSA_TM_15
		return nil, nil
	case 16: // LsarCreateSecret
		in := &CreateSecretRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateSecret(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // LsarOpenAccount
		in := &OpenAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // LsarEnumeratePrivilegesAccount
		in := &EnumeratePrivilegesAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumeratePrivilegesAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // LsarAddPrivilegesToAccount
		in := &AddPrivilegesToAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPrivilegesToAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // LsarRemovePrivilegesFromAccount
		in := &RemovePrivilegesFromAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemovePrivilegesFromAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // Opnum21NotUsedOnWire
		// Opnum21NotUsedOnWire
		return nil, nil
	case 22: // Opnum22NotUsedOnWire
		// Opnum22NotUsedOnWire
		return nil, nil
	case 23: // LsarGetSystemAccessAccount
		in := &GetSystemAccessAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSystemAccessAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // LsarSetSystemAccessAccount
		in := &SetSystemAccessAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSystemAccessAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // LsarOpenTrustedDomain
		in := &OpenTrustedDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenTrustedDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // LsarQueryInfoTrustedDomain
		in := &QueryInfoTrustedDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInfoTrustedDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // LsarSetInformationTrustedDomain
		in := &SetInformationTrustedDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetInformationTrustedDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // LsarOpenSecret
		in := &OpenSecretRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenSecret(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // LsarSetSecret
		in := &SetSecretRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSecret(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // LsarQuerySecret
		in := &QuerySecretRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QuerySecret(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // LsarLookupPrivilegeValue
		in := &LookupPrivilegeValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupPrivilegeValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // LsarLookupPrivilegeName
		in := &LookupPrivilegeNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupPrivilegeName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // LsarLookupPrivilegeDisplayName
		in := &LookupPrivilegeDisplayNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupPrivilegeDisplayName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // LsarDeleteObject
		in := &DeleteObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // LsarEnumerateAccountsWithUserRight
		in := &EnumerateAccountsWithUserRightRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateAccountsWithUserRight(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // LsarEnumerateAccountRights
		in := &EnumerateAccountRightsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateAccountRights(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // LsarAddAccountRights
		in := &AddAccountRightsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddAccountRights(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // LsarRemoveAccountRights
		in := &RemoveAccountRightsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveAccountRights(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // LsarQueryTrustedDomainInfo
		in := &QueryTrustedDomainInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryTrustedDomainInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // LsarSetTrustedDomainInfo
		in := &SetTrustedDomainInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTrustedDomainInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // LsarDeleteTrustedDomain
		in := &DeleteTrustedDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteTrustedDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // LsarStorePrivateData
		in := &StorePrivateDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StorePrivateData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // LsarRetrievePrivateData
		in := &RetrievePrivateDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RetrievePrivateData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // LsarOpenPolicy2
		in := &OpenPolicy2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenPolicy2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // Lsar_LSA_TM_45
		// Lsar_LSA_TM_45
		return nil, nil
	case 46: // LsarQueryInformationPolicy2
		in := &QueryInformationPolicy2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInformationPolicy2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // LsarSetInformationPolicy2
		in := &SetInformationPolicy2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetInformationPolicy2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // LsarQueryTrustedDomainInfoByName
		in := &QueryTrustedDomainInfoByNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryTrustedDomainInfoByName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // LsarSetTrustedDomainInfoByName
		in := &SetTrustedDomainInfoByNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetTrustedDomainInfoByName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // LsarEnumerateTrustedDomainsEx
		in := &EnumerateTrustedDomainsExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateTrustedDomainsEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // LsarCreateTrustedDomainEx
		in := &CreateTrustedDomainExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateTrustedDomainEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // Opnum52NotUsedOnWire
		// Opnum52NotUsedOnWire
		return nil, nil
	case 53: // LsarQueryDomainInformationPolicy
		in := &QueryDomainInformationPolicyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryDomainInformationPolicy(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // LsarSetDomainInformationPolicy
		in := &SetDomainInformationPolicyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDomainInformationPolicy(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 55: // LsarOpenTrustedDomainByName
		in := &OpenTrustedDomainByNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenTrustedDomainByName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 56: // Opnum56NotUsedOnWire
		// Opnum56NotUsedOnWire
		return nil, nil
	case 57: // Lsar_LSA_TM_57
		// Lsar_LSA_TM_57
		return nil, nil
	case 58: // Lsar_LSA_TM_58
		// Lsar_LSA_TM_58
		return nil, nil
	case 59: // LsarCreateTrustedDomainEx2
		in := &CreateTrustedDomainEx2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateTrustedDomainEx2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 60: // Opnum60NotUsedOnWire
		// Opnum60NotUsedOnWire
		return nil, nil
	case 61: // Opnum61NotUsedOnWire
		// Opnum61NotUsedOnWire
		return nil, nil
	case 62: // Opnum62NotUsedOnWire
		// Opnum62NotUsedOnWire
		return nil, nil
	case 63: // Opnum63NotUsedOnWire
		// Opnum63NotUsedOnWire
		return nil, nil
	case 64: // Opnum64NotUsedOnWire
		// Opnum64NotUsedOnWire
		return nil, nil
	case 65: // Opnum65NotUsedOnWire
		// Opnum65NotUsedOnWire
		return nil, nil
	case 66: // Opnum66NotUsedOnWire
		// Opnum66NotUsedOnWire
		return nil, nil
	case 67: // Opnum67NotUsedOnWire
		// Opnum67NotUsedOnWire
		return nil, nil
	case 68: // Lsar_LSA_TM_68
		// Lsar_LSA_TM_68
		return nil, nil
	case 69: // Opnum69NotUsedOnWire
		// Opnum69NotUsedOnWire
		return nil, nil
	case 70: // Opnum70NotUsedOnWire
		// Opnum70NotUsedOnWire
		return nil, nil
	case 71: // Opnum71NotUsedOnWire
		// Opnum71NotUsedOnWire
		return nil, nil
	case 72: // Opnum72NotUsedOnWire
		// Opnum72NotUsedOnWire
		return nil, nil
	case 73: // LsarQueryForestTrustInformation
		in := &QueryForestTrustInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryForestTrustInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 74: // LsarSetForestTrustInformation
		in := &SetForestTrustInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetForestTrustInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
