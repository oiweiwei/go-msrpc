package logon

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

// logon server interface.
type LogonServer interface {
	UASLogon(context.Context, *UASLogonRequest) (*UASLogonResponse, error)

	UASLogoff(context.Context, *UASLogoffRequest) (*UASLogoffResponse, error)

	// The NetrLogonSamLogon method<213> is a predecessor to the NetrLogonSamLogonWithFlags
	// method (section 3.5.4.5.2). All parameters of this method have the same meanings
	// as the identically named parameters of the NetrLogonSamLogonWithFlags method.
	//
	// Message processing<214> is identical to NetrLogonSamLogonEx, as specified in section
	// 3.5.4.5.1, except for the following:
	//
	// * The method uses Netlogon authenticators ( b5e7d25a-40b2-41c8-9611-98f53358af66#gt_e72a2c02-84a2-4ce3-b66f-86f725642dc3
	// ) , so instead of checking for Secure RPC, the server MUST confirm the validity of
	// the Authenticator (section 3.1.4.5 ( da7acaa3-030b-481e-979b-f58f89389806 ) ) that
	// it received using the ComputerName for the secure channel ( b5e7d25a-40b2-41c8-9611-98f53358af66#gt_08ce423c-9f9c-48ed-afa8-8b64c04ecdca
	// ) to find the corresponding record in the ClientSessionInfo table. If the Authenticator
	// parameter is valid, the server MUST compute the ReturnAuthenticator parameter returned
	// (section 3.1.4.5). Otherwise, the server MUST return STATUS_ACCESS_DENIED.
	//
	// * The ExtraFlags parameter is not processed.
	//
	// This method SHOULD only be called by a machine that has established a secure channel
	// with the server.
	SAMLogon(context.Context, *SAMLogonRequest) (*SAMLogonResponse, error)

	// The NetrLogonSamLogoff method SHOULD<215> update the user lastLogoff attribute for
	// the SAM accounts.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	SAMLogoff(context.Context, *SAMLogoffRequest) (*SAMLogoffResponse, error)

	// The NetrServerReqChallenge method SHOULD<172> receive a client challenge and return
	// a server challenge (SC).
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	RequestChallenge(context.Context, *RequestChallengeRequest) (*RequestChallengeResponse, error)

	// The NetrServerAuthenticate method<177> is a predecessor to the NetrServerAuthenticate3
	// method (section 3.5.4.4.2). All parameters of this method have the same meanings
	// as the identically named parameters of the NetrServerAuthenticate3 method.
	//
	// Message processing is identical to NetrServerAuthenticate3, as specified in section
	// 3.5.4.4.2, except for the following:
	//
	// * The NegotiateFlags parameter is not present in NetrServerAuthenticate. Message
	// processing would be identical to an invocation of NetrServerAuthenticate3 with the
	// NegotiateFlags parameter set to 0.
	//
	// * The AccountRid parameter is not present in NetrServerAuthenticate.
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)

	// The NetrServerPasswordSet method SHOULD<182> set a new one-way function (OWF) of
	// a password for an account used by the domain controller for setting up the secure
	// channel from the client.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	PasswordSet(context.Context, *PasswordSetRequest) (*PasswordSetResponse, error)

	// The NetrDatabaseDeltas method SHOULD<217> return a set of changes (or deltas) performed
	// to the SAM database, SAM built-in database, or LSA databases after a particular value
	// of the database serial number. It is used by BDCs to request database changes from
	// the PDC that are missing on the BDC.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// The synchronization that this method performs is not a full synchronization; rather,
	// a subset of database changes is returned. To perform a full synchronization, call
	// NetrDatabaseSync.
	//
	// The server that receives this call MUST do the following:
	//
	// * Verify that the client is a BDC.
	//
	// * Verify the client authenticator. The server MUST return status code STATUS_ACCESS_DENIED
	// if the verification fails.
	//
	// * Validate that DatabaseID is one of the allowed values, 0x00000000 through 0x00000002.
	// If the DatabaseID is not one of these values, the server MUST return the status code
	// STATUS_INVALID_LEVEL.
	//
	// * Given the BDC database serial number, obtain all database records that are missing
	// on the BDC and return the array of deltas, NETLOGON_DELTA_ENUM_ARRAY, for the missing
	// records. The number of elements returned is affected by the value of the PreferredMaximumLength
	// parameter. The server SHOULD <219> ( 0c858a52-732a-43ec-85dd-e44b357c1898#Appendix_A_219
	// ) stop including elements in the returned array after the size of the returned data
	// equals or exceeds the value of the PreferredMaximumLength parameter. The server SHOULD
	// also limit the number of elements per local configuration to avoid large array allocations.
	//
	// * Compute and return the server authenticator.
	//
	// * The server MUST set the value of the DomainModifiedCount parameter to the database
	// serial number of the last delta returned in the array.
	//
	// * If not all missing records are returned, the server MUST return the status code
	// STATUS_MORE_ENTRIES.
	//
	// * The server maintains and updates a state that indicates the client progress in
	// the synchronization protocol, as defined in section 3.6 ( f28f9dc8-eeb2-4112-9eec-a466f639c761
	// ).
	DatabaseDeltas(context.Context, *DatabaseDeltasRequest) (*DatabaseDeltasResponse, error)

	// The NetrDatabaseSync method<223> is a predecessor to the NetrDatabaseSync2 method
	// (section 3.5.4.6.2). All parameters of this method have the same meanings as the
	// identically named parameters of the NetrDatabaseSync2 method.
	//
	// Receiving this method is identical to receiving NetrDatabaseSync2, as specified in
	// section 3.5.4.6.2, except that this call does not use the RestartState parameter.
	// NetrDatabaseSync does not support restarting the synchronization loop.
	DatabaseSync(context.Context, *DatabaseSyncRequest) (*DatabaseSyncResponse, error)

	AccountDeltas(context.Context, *AccountDeltasRequest) (*AccountDeltasResponse, error)

	AccountSync(context.Context, *AccountSyncRequest) (*AccountSyncResponse, error)

	// The NetrGetDCName method MAY<160> be used to retrieve the NetBIOS name of the PDC
	// for the specified domain.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it MUST return
	// a nonzero error code and SHOULD return the following error code.
	//
	//	+--------------------------------+---------------------------------+
	//	|             RETURN             |                                 |
	//	|           VALUE/CODE           |           DESCRIPTION           |
	//	|                                |                                 |
	//	+--------------------------------+---------------------------------+
	//	+--------------------------------+---------------------------------+
	//	| 0x00000035 ERROR_ BAD_ NETPATH | The network path was not found. |
	//	+--------------------------------+---------------------------------+
	GetNetrDCName(context.Context, *GetNetrDCNameRequest) (*GetNetrDCNameResponse, error)

	// The NetrLogonControl method is a predecessor to the NetrLogonControl2Ex method (section
	// 3.5.4.9.1). All parameters of this method SHOULD<260> have the same meanings as the
	// identically named parameters of the NetrLogonControl2Ex method.
	//
	// All restrictions on parameter values in the NetrLogonControl2Ex method (section 3.5.4.9.1)
	// apply. Extra restrictions are applied to the values of the FunctionCode<261> and
	// QueryLevel parameters as follows:
	//
	// * The value of QueryLevel parameter is restricted to 0x00000001. If 0x00000002 is
	// used, the error ERROR_NOT_SUPPORTED is returned; if any value larger than 0x00000002
	// is used, the error ERROR_INVALID_LEVEL is returned.
	//
	// Message processing is identical to NetrLogonControl2Ex (section 3.5.4.9.1), except
	// for the following:
	//
	// * The Data parameter of *NetrLogonControl2Ex* is set to NULL.
	Control(context.Context, *ControlRequest) (*ControlResponse, error)

	// The NetrGetAnyDCName method MAY<162> be used to retrieve the name of a domain controller
	// in the specified primary or directly trusted domain. Only DCs can return the name
	// of a DC in a specified directly trusted domain.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it MUST return
	// a nonzero error code and SHOULD return the following error code.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                  |
	//	|                 VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000712 ERROR_DOMAIN_TRUST_INCONSISTENT | The name or security ID (SID) of the domain specified is inconsistent with the   |
	//	|                                            | trust information for that domain.                                               |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	GetAnyDCName(context.Context, *GetAnyDCNameRequest) (*GetAnyDCNameResponse, error)

	// The NetrLogonControl2 method<259> is a predecessor to the NetrLogonControl2Ex method
	// (section 3.5.4.9.1) and is updated to have the same functionality as NetrLogonControl2Ex.
	// All parameters of this method have the same meanings as the identically named parameters
	// of the NetrLogonControl2Ex method.
	//
	// All restrictions on parameter values in the NetrLoginControl2Ex method (section 3.5.4.9.1)
	// apply. Extra restrictions are applied to the values of the QueryLevel parameter as
	// follows:
	//
	// * If the QueryLevel parameter is set to 0x00000004, the error ERROR_INVALID_LEVEL
	// is returned.
	//
	// Message processing is identical to NetrLogonControl2Ex (section 3.5.4.9.1).
	Control2(context.Context, *Control2Request) (*Control2Response, error)

	// The NetrServerAuthenticate2 method<176> is a predecessor to the NetrServerAuthenticate3
	// method, as specified in section 3.5.4.4.2. All parameters of this method have the
	// same meanings as the identically named parameters of the NetrServerAuthenticate3
	// method.
	//
	// Message processing is identical to NetrServerAuthenticate3, except for the following:
	//
	// The AccountRid parameter is not present in NetrServerAuthenticate2.
	Authenticate2(context.Context, *Authenticate2Request) (*Authenticate2Response, error)

	// The NetrDatabaseSync2 method SHOULD<220> return a set of all changes applied to the
	// specified database since its creation. It provides an interface for a BDC to fully
	// synchronize its databases to those of the PDC. Because returning all changes in one
	// call might be prohibitively expensive due to a large amount of data being returned,
	// this method supports retrieving portions of the database changes in a series of calls
	// using a continuation context until all changes are received. It is possible for the
	// series of calls to be terminated prematurely due to external events, such as system
	// restarts. For that reason, the method also supports restarting the series of calls
	// at a particular point specified by the caller. The caller MUST keep track of synchronization
	// progress during the series of calls as detailed in this section.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// The server that receives this call MUST do the following:
	//
	// * Verify that the client is a backup domain controller (BDC) ( b5e7d25a-40b2-41c8-9611-98f53358af66#gt_ce1138c6-7ab4-4c37-98b4-95599071c3c3
	// ) , the server is a PDC, and is enabled. If any of these conditions are false, the
	// server MUST return the status code STATUS_NOT_SUPPORTED.
	//
	// * Apply Common Error Processing Rule B, specified in section 3 ( 2d776bfc-e81f-4c8f-9da8-4c2920f65413
	// ).
	//
	// * Using the ComputerName for the secure channel ( b5e7d25a-40b2-41c8-9611-98f53358af66#gt_08ce423c-9f9c-48ed-afa8-8b64c04ecdca
	// ) to find the corresponding record in the ClientSessionInfo table, verify the Authenticator
	// parameter (section 3.1.4.5 ( da7acaa3-030b-481e-979b-f58f89389806 ) ). If the Authenticator
	// parameter is valid, compute the ReturnAuthenticator parameter returned (section 3.1.4.5).
	// Otherwise, the server MUST return STATUS_ACCESS_DENIED.
	//
	// * Validate that DatabaseID is one of the allowed values, 0x00000000 through 0x00000002.
	// If the DatabaseID is not one of these values, the server MUST return the status code
	// STATUS_INVALID_LEVEL.
	//
	// * Given the RestartState parameter and the SyncContext parameter, obtain database
	// records that are missing on the BDC and return the array of deltas, NETLOGON_DELTA_ENUM_ARRAY,
	// for the missing records. The number of elements returned is affected by the value
	// of the PreferredMaximumLength parameter. The server SHOULD <222> ( 0c858a52-732a-43ec-85dd-e44b357c1898#Appendix_A_222
	// ) stop including elements in the returned array once the size of the returned data
	// equals or exceeds the value of the PreferredMaximumLength parameter. The server SHOULD
	// also limit the number of elements per local configuration to avoid large array allocations.
	//
	// * The server MUST update and return the SyncContext parameter (section 2.2.1.5.29
	// ( ffce9bbc-6ae5-4478-8f45-e82c3847aaa2 ) ) to continue the synchronization loop on
	// the next client request.
	//
	// * Compute and return the server authenticator.
	//
	// * Initialize *SynchronizationComplete* by setting it to FALSE, and when all the missing
	// records are sent set *SynchronizationComplete* to TRUE.
	//
	// * If *SynchronizationComplete* is FALSE, the server MUST return the status code STATUS_MORE_ENTRIES.
	DatabaseSync2(context.Context, *DatabaseSync2Request) (*DatabaseSync2Response, error)

	// The NetrDatabaseRedo method SHOULD<224> be used by a backup domain controller (BDC)
	// to request information about a single account from the PDC.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// The following CHANGELOG_ENTRY structure pointed to by the ChangeLogEntry parameter
	// carries information about the account object being queried.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| SerialNumber [0..3]                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| SerialNumber [4..7]                                                                                                           |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ObjectRid                                                                                                                     |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| Flags                                                         | DBIndex                       | DeltaType                     |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ObjectSid (optional, variable length) ...                                                                                     |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| ObjectName (optional, variable length) ...                                                                                    |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	DatabaseRedo(context.Context, *DatabaseRedoRequest) (*DatabaseRedoResponse, error)

	// The NetrLogonControl2Ex method SHOULD<243> execute administrative actions that pertain
	// to the Netlogon server operation. It is used to query the status and control the
	// actions of the Netlogon server.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	Control2Ex(context.Context, *Control2ExRequest) (*Control2ExResponse, error)

	// The NetrEnumerateTrustedDomains method SHOULD<227> return a set of NetBIOS names
	// of trusted domains.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it SHOULD return
	// one of the following error codes.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000051F ERROR_NO_LOGON_SERVERS     | There are currently no logon servers available to service the logon request.     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006FA ERROR_NO_TRUST_LSA_SECRET  | The workstation does not have a trust secret.                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006FB ERROR_NO_TRUST_SAM_ACCOUNT | The security database on the server does not have a computer account for this    |
	//	|                                       | workstation trust relationship.                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// The server initializes the DomainNames field of the DOMAIN_NAME_BUFFER to the empty
	// string. The server calls the NetrEnumerateTrustedDomainsEx method and for each PDS_DOMAIN_TRUSTSW
	// element of the NETLOGON_TRUSTED_DOMAIN_ARRAY, appends the NetbiosDomainName field
	// to the DomainNames field of the DOMAIN_NAME_BUFFER (section 2.2.1.6.1). Then the
	// server terminates the DomainNames field with two null bytes.
	//
	// For details, see section 3.5.4.7.2, Receiving NetrEnumerateTrustedDomainsEx.
	EnumerateTrustedDomains(context.Context, *EnumerateTrustedDomainsRequest) (*EnumerateTrustedDomainsResponse, error)

	// The DsrGetDcName method is a predecessor to the DsrGetDcNameEx2 method (section 3.5.4.3.1).
	// The method SHOULD<159> return information about a domain controller in the specified
	// domain. All parameters of this method have the same meanings as the identically named
	// parameters of the DsrGetDcNameEx2 method, except for the SiteGuid parameter, detailed
	// as follows.
	GetDCName(context.Context, *GetDCNameRequest) (*GetDCNameResponse, error)

	// The NetrLogonGetCapabilities method is used by clients to confirm the server capabilities
	// after a secure channel has been established.<194>
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	GetCapabilities(context.Context, *GetCapabilitiesRequest) (*GetCapabilitiesResponse, error)

	// The NetrLogonSetServiceBits method SHOULD<237> be used to notify Netlogon whether
	// a domain controller is running specified services, as detailed in the following section.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// The server SHOULD<240> return ERROR_ACCESS_DENIED if the caller is not local.
	//
	// If the client does not have sufficient privilege, the server MUST return ERROR_ACCESS_DENIED.
	SetServiceBits(context.Context, *SetServiceBitsRequest) (*SetServiceBitsResponse, error)

	// The NetrLogonGetTrustRid method SHOULD<231> be used to obtain the RID of the account
	// whose password is used by domain controllers in the specified domain for establishing
	// the secure channel from the server receiving this call.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// The server SHOULD<232> return ERROR_ACCESS_DENIED if the caller is not local.
	//
	// If ServerName equals NULL and DomainName equals NULL, the server determines if the
	// client has sufficient privilege (as specified in section 3.5.4.2) with the Access
	// Request mask set to NETLOGON_FTINFO_ACCESS.
	//
	// Otherwise, the server determines if the client has sufficient privilege (as specified
	// in section 3.5.4.2) with the Access Request mask set to NETLOGON_SERVICE_ACCESS.
	//
	// If the client does not have sufficient privilege, the server MUST return ERROR_ACCESS_DENIED.
	//
	// If ServerName equals NULL, then the call MUST be made to the local machine. If the
	// DomainName is the same as the domain that the machine is joined to, the call MUST
	// succeed, and the server MUST return the AccountRid of the machine in the domain.
	// If the DomainName is a different domain, the server MUST return ERROR_NO_SUCH_DOMAIN.
	//
	// If both ServerName and DomainName are NULL, the server MUST return the RID for the
	// computer account of the caller. Otherwise, the RID for the account identified by
	// ServerName and DomainName MUST be returned.
	//
	// The server uses the server name passed in the ServerName parameter to look up the
	// domain for the request. If the name is not found, the server MUST return ERROR_INVALID_COMPUTERNAME.
	//
	// If the RID cannot be determined, the server SHOULD return ERROR_TRUSTED_RELATIONSHIP_FAILURE.
	GetTrustRID(context.Context, *GetTrustRIDRequest) (*GetTrustRIDResponse, error)

	// The NetrLogonComputeServerDigest method computes a cryptographic digest of a message
	// by using the MD5 message-digest algorithm, as specified in [RFC1321]. This method
	// SHOULD<233> be called by a client computer against a server and is used to compute
	// a message digest, as specified in this section. The client then calls the NetrLogonComputeClientDigest
	// method (as specified in section 3.4.5.6.3) and compare the digests to ensure that
	// the server that it communicates with knows the shared secret between the client machine
	// and the domain.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// The server uses the server name passed in the ServerName parameter to look up the
	// domain for the request. If the name is not found, the server MUST return ERROR_INVALID_COMPUTERNAME.
	//
	// If the client does not have sufficient privilege, the server MUST return ERROR_ACCESS_DENIED.
	ComputeServerDigest(context.Context, *ComputeServerDigestRequest) (*ComputeServerDigestResponse, error)

	// The NetrLogonComputeClientDigest method is used by a client to compute a cryptographic
	// digest of a message by using the MD5 message-digest algorithm, as specified in [RFC1321].
	// This method is called by a client to compute a message digest, as specified in this
	// section. The client SHOULD<235> use this digest to compare against one that is returned
	// by a call to NetrLogonComputeServerDigest. This comparison allows the client to ensure
	// that the server that it communicates with knows the shared secret between the client
	// machine and the domain.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// If the client does not have sufficient privilege, the server MUST return ERROR_ACCESS_DENIED.
	ComputeClientDigest(context.Context, *ComputeClientDigestRequest) (*ComputeClientDigestResponse, error)

	// The NetrServerAuthenticate3 method SHOULD<173> mutually authenticate the client and
	// the server and establish the session key to be used for the secure channel message
	// protection between the client and the server. It is called after the NetrServerReqChallenge
	// method, as specified in section 3.5.4.4.1.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	Authenticate3(context.Context, *Authenticate3Request) (*Authenticate3Response, error)

	// The DsrGetDcNameEx method is a predecessor to the DsrGetDcNameEx2 (section 3.5.4.3.1)
	// method. The method SHOULD<158> return information about a domain controller in the
	// specified domain and site. All parameters of this method have the same meanings as
	// the identically named parameters of the DsrGetDcNameEx2 method.
	GetDCNameEx(context.Context, *GetDCNameExRequest) (*GetDCNameExResponse, error)

	// The DsrGetSiteName method SHOULD<164> return the site name for the specified computer
	// that receives this call.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it MUST return
	// a nonzero error code and SHOULD return the following error code.
	//
	//	+------------------------------+---------------------------------------------+
	//	|            RETURN            |                                             |
	//	|          VALUE/CODE          |                 DESCRIPTION                 |
	//	|                              |                                             |
	//	+------------------------------+---------------------------------------------+
	//	+------------------------------+---------------------------------------------+
	//	| 0x0000077F ERROR_NO_SITENAME | No site name is available for this machine. |
	//	+------------------------------+---------------------------------------------+
	//
	// If the computer has been configured with a SiteName, it MUST return the SiteName
	// immediately.
	//
	// If the DynamicSiteNameSetTime plus the DynamicSiteNameTimeout is less than the current
	// time (meaning that the DynamicSiteNameSetTime is older than allowed by DynamicSiteNameTimeout),
	// then:
	//
	// * The server MUST locate a domain controller ( b5e7d25a-40b2-41c8-9611-98f53358af66#gt_76a05049-3531-4abd-aec8-30e19954b4bd
	// ) in the domain. The server SHOULD <165> ( 0c858a52-732a-43ec-85dd-e44b357c1898#Appendix_A_165
	// ) implement alternate means to locate DCs: for example, a static list in a file,
	// or the two methods detailed in [MS-ADTS] ( ../ms-adts/d2435927-0999-4c62-8c6d-13ba31a52e1a
	// ) section 6.3.6 ( ../ms-adts/3e9711af-9067-435e-93fb-09182053cef7 ). If the server
	// cannot locate a DC for the domain, then the server MUST return ERROR_NO_SUCH_DOMAIN.
	//
	// * The server then populates the SiteName parameter with the NETLOGON_SAM_LOGON_RESPONSE_EX
	// message ([MS-ADTS] section 6.3.1.9 ( ../ms-adts/8401a33f-34a8-40ca-bf03-c3484b66265f
	// ) ) by setting the SiteName parameter to NETLOGON_SAM_LOGON_RESPONSE_EX.ClientSiteName.
	// The server stores the discovered site name in *DynamicSiteName*.
	//
	// * The server sets the DynamicSiteNameSetTime to the current time.
	//
	// Otherwise, DynamicSiteName MUST be returned immediately as the SiteName parameter.
	//
	// If it is determined that the server that receives this call has no site name, the
	// server MUST return ERROR_NO_SITENAME.
	//
	// This method also returns errors based on Common Error Processing Rules B and C, specified
	// in section 3.
	GetSiteName(context.Context, *GetSiteNameRequest) (*GetSiteNameResponse, error)

	// The NetrLogonGetDomainInfo method SHOULD<188> return information that describes the
	// current domain to which the specified client belongs.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	GetDomainInfo(context.Context, *GetDomainInfoRequest) (*GetDomainInfoResponse, error)

	// The NetrServerPasswordSet2 method SHOULD<178> allow the client to set a new clear
	// text password for an account used by the domain controller for setting up the secure
	// channel from the client. A domain member SHOULD<179> use this function to periodically
	// change its machine account password. A PDC uses this function to periodically change
	// the trust password for all directly trusted domains.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	PasswordSet2(context.Context, *PasswordSet2Request) (*PasswordSet2Response, error)

	// The NetrServerPasswordGet method SHOULD<184> allow a BDC to get a machine account
	// password from the DC with the PDC role in the domain.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	PasswordGet(context.Context, *PasswordGetRequest) (*PasswordGetResponse, error)

	// The NetrLogonSendToSam method allows a BDC or RODC to forward user account password
	// changes to the PDC. It SHOULD<236> be used by the client to deliver an opaque buffer
	// to the SAM database ([MS-SAMR] section 3.1.1) on the server side.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	SendToSAM(context.Context, *SendToSAMRequest) (*SendToSAMResponse, error)

	// The DsrAddressToSiteNamesW method SHOULD<167> translate a list of socket addresses
	// into their corresponding site names. For details about the mapping from socket address
	// to subnet/site name, see [MS-ADTS] sections 6.1.1.2.2.1 and 6.1.1.2.2.2.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters is invalid. This error value is returned if the value of   |
	//	|                                    | EntryCount passed to DsrAddressToSiteNamesW is zero.                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// This method returns errors based on Common Error Processing Rule A, specified in
	// section 3.
	AddressToSiteNamesW(context.Context, *AddressToSiteNamesWRequest) (*AddressToSiteNamesWResponse, error)

	// The DsrGetDcNameEx2 method SHOULD<147> return information about a domain controller
	// (DC) in the specified domain and site. If the AccountName parameter is not NULL,
	// and a DC matching the requested capabilities (as defined in the Flags parameter)
	// responds during this method call, then that DC will have verified that the DC account
	// database contains an account for the AccountName specified. The server that receives
	// this call is not required to be a DC.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	GetDCNameEx2(context.Context, *GetDCNameEx2Request) (*GetDCNameEx2Response, error)

	// The NetrLogonGetTimeServiceParentDomain method SHOULD<241> return the name of the
	// parent domain of the current domain. The domain name returned by this method is suitable
	// for passing into the NetrLogonGetTrustRid method and NetrLogonComputeClientDigest
	// method.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// The server SHOULD<242> return ERROR_ACCESS_DENIED if the caller is not local.
	//
	// If the client does not have sufficient privilege, the server MUST return ERROR_ACCESS_DENIED.
	//
	// The domain name returned MUST be determined according to the following rules:
	//
	// * On a non-DC machine, the returned domain name is the name of the domain of which
	// the ServerName is a member. If ServerName is not valid, the server MUST return ERROR_INVALID_COMPUTERNAME.
	//
	// * On a DC that is at the root of the forest ( b5e7d25a-40b2-41c8-9611-98f53358af66#gt_fd104241-4fb3-457c-b2c4-e0c18bb20b62
	// ) , *rootDomainNamingContext* ( [MS-ADTS] ( ../ms-adts/d2435927-0999-4c62-8c6d-13ba31a52e1a
	// ) section 3.1.1.3.2.16 ( ../ms-adts/7ee04fa8-8027-4c7e-9c4a-4cd4c0198a70 ) ) is equal
	// to *defaultNamingContext* ([MS-ADTS] section 3.1.1.3.2.3 ( ../ms-adts/c1ad47fb-e4a1-4240-a749-b8f07ee0ae91
	// ) ). In this case, ERROR_NO_SUCH_DOMAIN is returned.
	//
	// * On a DC that is at the root of a domain tree ( b5e7d25a-40b2-41c8-9611-98f53358af66#gt_e2c071f1-5977-4749-868f-6c5a04929476
	// ) in the forest, the name of a trusted ( b5e7d25a-40b2-41c8-9611-98f53358af66#gt_5ee032d0-d944-4acb-bbb5-b1cfc7df6db6
	// ) domain that is also at the root of a domain tree in the forest is returned.
	//
	// On any other DC, the name of the domain that is directly the parent domain is returned.
	//
	// The domain's information MUST be retrieved from an implementation-specific directory.
	// Based on this retrieved information, if the domain has a DNS domain name, it MUST
	// be returned; otherwise, the NetBIOS domain name MUST be returned. This behavior is
	// functionally equivalent to locally invoking LsarQueryTrustedDomainInfo ([MS-LSAD]
	// section 3.1.4.7.2) for the domain, where TrustedDomainSid is the domain SID corresponding
	// to the appropriate domain name retrieved from a cached list, and InformationClass
	// is TrustedDomainInformationEx (policy handle is not needed locally), to return the
	// TrustedDomainInformationEx.Name string (DNS name) if it is present or TrustedDomainInformationEx.Flat
	// Name string (NetBIOS name) otherwise.
	//
	// The PdcSameSite returned MUST be determined according to the following rules:
	//
	// * On a non-DC machine, the value of PdcSameSite is set to TRUE.
	//
	// * On a DC machine, the server determines the PDC as specified in [MS-ADTS] section
	// 3.1.1.1.11 ( ../ms-adts/bf8afb5f-1ae0-45de-8445-8a717ea5124a ). Then the server determines
	// the sites of both the server and PDC as specified in [MS-ADTS] section 3.1.1.4.5.29
	// ( ../ms-adts/930d3ab8-7344-49f4-9615-a03d38279dff ). The server MUST compare the
	// PDC site with its own site, and if the two match, the PdcSameSite output parameter
	// MUST be set to TRUE; otherwise, it MUST be set to FALSE.
	GetTimeServiceParentDomain(context.Context, *GetTimeServiceParentDomainRequest) (*GetTimeServiceParentDomainResponse, error)

	// The NetrEnumerateTrustedDomainsEx method SHOULD<226> return a list of trusted domains
	// from a specified server. This method extends NetrEnumerateTrustedDomains by returning
	// an array of domains in a more flexible DS_DOMAIN_TRUSTSW structure, as specified
	// in section 2.2.1.6.2, rather than the array of strings in DOMAIN_NAME_BUFFER structure,
	// as specified in section 2.2.1.6.1. The array is returned as part of the NETLOGON_TRUSTED_DOMAIN_ARRAY
	// structure, as specified in section 2.2.1.6.3.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it SHOULD return
	// one of the following error codes.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000051F ERROR_NO_LOGON_SERVERS     | There are currently no logon servers available to service the logon request.     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006FA ERROR_NO_TRUST_LSA_SECRET  | The workstation does not have a trust secret.                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006FB ERROR_NO_TRUST_SAM_ACCOUNT | The security database on the server does not have a computer account for this    |
	//	|                                       | workstation trust relationship.                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// This method is a wrapper for DsrEnumerateDomainTrusts, which strips off the F flag
	// from the returned data for backward compatibility. For details, see section 3.5.4.7.1.
	EnumerateTrustedDomainsEx(context.Context, *EnumerateTrustedDomainsExRequest) (*EnumerateTrustedDomainsExResponse, error)

	// The DsrAddressToSiteNamesExW method SHOULD<168> translate a list of socket addresses
	// into their corresponding site names and subnet names. For details about the mapping
	// from socket address to subnet/site name, see [MS-ADTS] sections 6.1.1.2.2.1 and 6.1.1.2.2.2.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it SHOULD return
	// one of the following error codes.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One of the parameters is invalid. This error value is returned if the value of   |
	//	|                                    | EntryCount passed to DsrAddressToSiteNamesExW is zero.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// This method returns errors based on Common Error Processing Rule A, specified in
	// section 3.
	AddressToSiteNamesExW(context.Context, *AddressToSiteNamesExWRequest) (*AddressToSiteNamesExWResponse, error)

	// The DsrGetDcSiteCoverageW method SHOULD<166> return a list of sites covered by a
	// domain controller. Site coverage is detailed in [MS-ADTS] section 6.1.1.2.2.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// This method returns errors based on Common Error Processing Rules A and B, specified
	// in section 3.
	GetDCSiteCoverageW(context.Context, *GetDCSiteCoverageWRequest) (*GetDCSiteCoverageWResponse, error)

	// The NetrLogonSamLogonEx method SHOULD<199> provide an extension to NetrLogonSamLogon
	// that accepts an extra flags parameter and uses Secure RPC ([MS-RPCE] section 3.3.1.5.2)
	// instead of Netlogon authenticators. This method handles logon requests for the SAM
	// accounts and allows for generic pass-through authentication, as specified in section
	// 3.2.4.1.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	SAMLogonEx(context.Context, *SAMLogonExRequest) (*SAMLogonExResponse, error)

	// The DsrEnumerateDomainTrusts method SHOULD<225> return an enumerated list of domain
	// trusts, filtered by a set of flags, from the specified server.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it SHOULD return
	// one of the following error codes.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000051F ERROR_NO_LOGON_SERVERS     | There are currently no logon servers available to service the logon request.     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006FA ERROR_NO_TRUST_LSA_SECRET  | The workstation does not have a trust secret.                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006FB ERROR_NO_TRUST_SAM_ACCOUNT | The security database on the server does not have a computer account for this    |
	//	|                                       | workstation trust relationship.                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	EnumerateDomainTrusts(context.Context, *EnumerateDomainTrustsRequest) (*EnumerateDomainTrustsResponse, error)

	// The DsrDeregisterDnsHostRecords method SHOULD<169> delete all of the DNS SRV records
	// registered by a specified domain controller. For the list of SRV records that a domain
	// registers, see [MS-ADTS] section 6.3.2.3, "SRV Records Registered by DC".
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it SHOULD return
	// the following error code.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED | The request is not supported. This error value is returned when                  |
	//	|                                | DsrDeregisterDnsHostRecords is called on a machine that is not a DC.             |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The server determines if the client has sufficient privileges (as specified in section
	// 3.5.4.2) with the Access Request mask set to the NETLOGON_CONTROL_ACCESS mask.
	//
	// If the client does not have sufficient privilege, the server MUST return ERROR_ACCESS_DENIED.
	//
	// If the DnsHostName parameter is not null, the server MUST attempt to delete the DNS
	// SRV records registered for the DC DnsHostName, as specified in [MS-ADTS] section
	// 6.3.2.3.
	//
	// If the DomainGuid parameter is not null, then the server MUST attempt to delete the
	// domain-GUID-based SRV record.
	//
	// If the DsaGuid parameter is not null, then the server MUST attempt to delete the
	// domain CNAME record.
	//
	// The deletion of site-specific records MUST be attempted for every site in the enterprise
	// of the DC on which the method is executed.
	//
	// Unless stated otherwise, if the attempt to delete any records documented previously
	// fails for any reason, then the server MUST ignore the error and continue message
	// processing.
	//
	// It is possible that this method call will create a time-consuming run that generates
	// significant network traffic for enterprises with many sites.
	DeregisterDNSHostRecords(context.Context, *DeregisterDNSHostRecordsRequest) (*DeregisterDNSHostRecordsResponse, error)

	// The NetrServerTrustPasswordsGet method SHOULD<186> return the encrypted current and
	// previous passwords for an account in the domain. This method is called by a client
	// to retrieve the current and previous account passwords from a domain controller.
	// The account name requested MUST be the name used when the secure channel was created,
	// unless the method is called on a PDC by a DC, in which case it can be any valid account
	// name.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// Message processing is identical to NetrServerGetTrustInfo, as specified in section
	// 3.5.4.7.6, except for the following:
	//
	// * The TrustInfo parameter is not present in NetrServerTrustPasswordsGet.
	TrustPasswordsGet(context.Context, *TrustPasswordsGetRequest) (*TrustPasswordsGetResponse, error)

	// The DsrGetForestTrustInformation method SHOULD<229> retrieve the trust information
	// for the forest of the specified domain controller (DC), or for a forest trusted by
	// the forest of the specified DC.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it SHOULD return
	// the following error code.
	//
	//	+-----------------------------------+---------------------+
	//	|              RETURN               |                     |
	//	|            VALUE/CODE             |     DESCRIPTION     |
	//	|                                   |                     |
	//	+-----------------------------------+---------------------+
	//	+-----------------------------------+---------------------+
	//	| 0x00000001 ERROR_INVALID_FUNCTION | Incorrect function. |
	//	+-----------------------------------+---------------------+
	GetForestTrustInformation(context.Context, *GetForestTrustInformationRequest) (*GetForestTrustInformationResponse, error)

	// The NetrGetForestTrustInformation method SHOULD<228> retrieve the trust information
	// for the forest of which the member's domain is itself a member.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	GetNetrForestTrustInformation(context.Context, *GetNetrForestTrustInformationRequest) (*GetNetrForestTrustInformationResponse, error)

	// The NetrLogonSamLogonWithFlags method SHOULD<210> handle logon requests for the SAM
	// accounts.
	//
	// Return Values: The method returns 0x00000000 on success; otherwise, it returns a
	// nonzero error code.
	//
	// Message processing is identical to NetrLogonSamLogon, as specified in section 3.5.4.5.3,
	// except for the following:
	//
	// * NetrLogonSamLogonWithFlags contains an additional parameter named ExtraFlags.
	SAMLogonWithFlags(context.Context, *SAMLogonWithFlagsRequest) (*SAMLogonWithFlagsResponse, error)

	// The NetrServerGetTrustInfo method SHOULD<230> return an information block from a
	// specified server. The information includes encrypted current and previous passwords
	// for a particular account and additional trust data. The account name requested MUST
	// be the name used when the secure channel was created, unless the method is called
	// on a PDC by a domain controller, in which case it can be any valid account name.
	//
	// Return Values: The method returns 0x00000000 to indicate success; otherwise, it returns
	// a nonzero error code.
	GetTrustInfo(context.Context, *GetTrustInfoRequest) (*GetTrustInfoResponse, error)

	// OpnumUnused47 operation.
	// OpnumUnused47

	// The DsrUpdateReadOnlyServerDnsRecords method SHOULD<170> allow an RODC to send a
	// control command to a normal (writable) DC for site-specific and CName types of DNS
	// records update. For registration, site-specific records are for the site in which
	// RODC resides. For the types of DNS records, see [MS-ADTS] section 6.3.2.
	//
	// Return Values: The method returns 0x00000000 (NO_ERROR) on success; otherwise, it
	// returns a nonzero error code.
	UpdateReadOnlyServerDNSRecords(context.Context, *UpdateReadOnlyServerDNSRecordsRequest) (*UpdateReadOnlyServerDNSRecordsResponse, error)

	// The NetrChainSetClientAttributes method SHOULD<198> be invoked by an RODC  on a
	// normal (writable) DC to update to a client's computer account object in Active Directory
	// when it receives either the NetrServerAuthenticate3 method or the NetrLogonGetDomainInfo
	// method with updates requested.
	//
	// Return Values: The method returns 0x00000000 on success.
	ChainSetClientAttributes(context.Context, *ChainSetClientAttributesRequest) (*ChainSetClientAttributesResponse, error)
}

func RegisterLogonServer(conn dcerpc.Conn, o LogonServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewLogonServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(LogonSyntaxV1_0))...)
}

func NewLogonServerHandle(o LogonServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return LogonServerHandle(ctx, o, opNum, r)
	}
}

func LogonServerHandle(ctx context.Context, o LogonServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // NetrLogonUasLogon
		in := &UASLogonRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.UASLogon(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // NetrLogonUasLogoff
		in := &UASLogoffRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.UASLogoff(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // NetrLogonSamLogon
		in := &SAMLogonRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SAMLogon(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // NetrLogonSamLogoff
		in := &SAMLogoffRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SAMLogoff(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // NetrServerReqChallenge
		in := &RequestChallengeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RequestChallenge(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // NetrServerAuthenticate
		in := &AuthenticateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Authenticate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // NetrServerPasswordSet
		in := &PasswordSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PasswordSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // NetrDatabaseDeltas
		in := &DatabaseDeltasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DatabaseDeltas(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // NetrDatabaseSync
		in := &DatabaseSyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DatabaseSync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // NetrAccountDeltas
		in := &AccountDeltasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AccountDeltas(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // NetrAccountSync
		in := &AccountSyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AccountSync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // NetrGetDCName
		in := &GetNetrDCNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNetrDCName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // NetrLogonControl
		in := &ControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Control(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // NetrGetAnyDCName
		in := &GetAnyDCNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAnyDCName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // NetrLogonControl2
		in := &Control2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Control2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // NetrServerAuthenticate2
		in := &Authenticate2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Authenticate2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // NetrDatabaseSync2
		in := &DatabaseSync2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DatabaseSync2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // NetrDatabaseRedo
		in := &DatabaseRedoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DatabaseRedo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // NetrLogonControl2Ex
		in := &Control2ExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Control2Ex(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // NetrEnumerateTrustedDomains
		in := &EnumerateTrustedDomainsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateTrustedDomains(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // DsrGetDcName
		in := &GetDCNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDCName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // NetrLogonGetCapabilities
		in := &GetCapabilitiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCapabilities(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // NetrLogonSetServiceBits
		in := &SetServiceBitsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetServiceBits(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // NetrLogonGetTrustRid
		in := &GetTrustRIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTrustRID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // NetrLogonComputeServerDigest
		in := &ComputeServerDigestRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ComputeServerDigest(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // NetrLogonComputeClientDigest
		in := &ComputeClientDigestRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ComputeClientDigest(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // NetrServerAuthenticate3
		in := &Authenticate3Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Authenticate3(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // DsrGetDcNameEx
		in := &GetDCNameExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDCNameEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // DsrGetSiteName
		in := &GetSiteNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSiteName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // NetrLogonGetDomainInfo
		in := &GetDomainInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDomainInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // NetrServerPasswordSet2
		in := &PasswordSet2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PasswordSet2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // NetrServerPasswordGet
		in := &PasswordGetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PasswordGet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // NetrLogonSendToSam
		in := &SendToSAMRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SendToSAM(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // DsrAddressToSiteNamesW
		in := &AddressToSiteNamesWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddressToSiteNamesW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // DsrGetDcNameEx2
		in := &GetDCNameEx2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDCNameEx2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // NetrLogonGetTimeServiceParentDomain
		in := &GetTimeServiceParentDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTimeServiceParentDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // NetrEnumerateTrustedDomainsEx
		in := &EnumerateTrustedDomainsExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateTrustedDomainsEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // DsrAddressToSiteNamesExW
		in := &AddressToSiteNamesExWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddressToSiteNamesExW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // DsrGetDcSiteCoverageW
		in := &GetDCSiteCoverageWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDCSiteCoverageW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // NetrLogonSamLogonEx
		in := &SAMLogonExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SAMLogonEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // DsrEnumerateDomainTrusts
		in := &EnumerateDomainTrustsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateDomainTrusts(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // DsrDeregisterDnsHostRecords
		in := &DeregisterDNSHostRecordsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeregisterDNSHostRecords(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // NetrServerTrustPasswordsGet
		in := &TrustPasswordsGetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.TrustPasswordsGet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // DsrGetForestTrustInformation
		in := &GetForestTrustInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetForestTrustInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // NetrGetForestTrustInformation
		in := &GetNetrForestTrustInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNetrForestTrustInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // NetrLogonSamLogonWithFlags
		in := &SAMLogonWithFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SAMLogonWithFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // NetrServerGetTrustInfo
		in := &GetTrustInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTrustInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // OpnumUnused47
		// OpnumUnused47
		return nil, nil
	case 48: // DsrUpdateReadOnlyServerDnsRecords
		in := &UpdateReadOnlyServerDNSRecordsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.UpdateReadOnlyServerDNSRecords(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // NetrChainSetClientAttributes
		in := &ChainSetClientAttributesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ChainSetClientAttributes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
