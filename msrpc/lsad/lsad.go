// The lsad package implements the LSAD client protocol.
//
// # Introduction
//
// The Local Security Authority (Domain Policy) Remote Protocol is used to manage various
// machine and domain security policies. All versions of Windows NT operating system–based
// products, in all configurations, implement and listen on the server side of this
// protocol. However, not all operations are meaningful in all configurations.
//
// This protocol, with minor exceptions, enables remote policy-management scenarios.
// Therefore, the majority of this interface does not need to be implemented to achieve
// Windows client-to-server (domain controller configuration and otherwise) interoperability,
// as defined by the ability for Windows clients to retrieve policy settings from servers.
//
// Policy settings controlled by this protocol relate to the following:
//
// * *Account objects:* The rights and privileges ( 31ca2a31-0be4-4773-bcef-05ad6cd3ccfb#gt_d8092e10-b227-4b44-b015-511bb8178940
// ) that security principals ( 31ca2a31-0be4-4773-bcef-05ad6cd3ccfb#gt_f3ef2572-95cf-4c5c-b3c9-551fd648f409
// ) have on the server.
//
// * *Secret objects:* Mechanisms that securely store data on the server.
//
// * *Trusted domain objects:* Mechanisms that the Windows operating system uses for
// describing trust ( 31ca2a31-0be4-4773-bcef-05ad6cd3ccfb#gt_5ee032d0-d944-4acb-bbb5-b1cfc7df6db6
// ) relationships between domains and forests ( 31ca2a31-0be4-4773-bcef-05ad6cd3ccfb#gt_fd104241-4fb3-457c-b2c4-e0c18bb20b62
// ).
//
// * Other miscellaneous settings, such as lifetimes of Kerberos tickets, states of
// domain controller (backup or primary), and other unrelated pieces of policy.
//
// All of these types of policy are addressed in sections of this document that specify
// the server data model.
//
// # Overview
//
// The Local Security Authority (Domain Policy) Remote Protocol provides a remote procedure
// call (RPC) interface used for providing remote management for policy settings related
// to account objects, secret objects, trusted domain objects (TDOs), and other miscellaneous
// security-related policy settings. The client end of the Local Security Authority
// (Domain Policy) Remote Protocol is an application that issues method calls on the
// RPC interface. The server end of the Local Security Authority (Domain Policy) Remote
// Protocol is a service that implements support for this RPC interface.
//
// The following represent primary use cases for remote management:
//
// * Creating, deleting, enumerating, and modifying trusts ( 31ca2a31-0be4-4773-bcef-05ad6cd3ccfb#gt_5ee032d0-d944-4acb-bbb5-b1cfc7df6db6
// ) , account objects, and secret objects.
//
// * Querying and modifying policy settings unrelated to TDOs, account objects or secret
// objects, such as lifetimes of Kerberos tickets.
//
// This protocol is used by Windows clients for the "domain join" operation (as specified
// in [MS-ADTS] section 6.4) as an implementation choice to achieve the end state, as
// specified in [MS-ADTS]. The specific profile of the Local Security Authority (Domain
// Policy) Remote Protocol for the "domain join" scenario is specified in section 1.6
// as "Retrieval of policy settings by clients".
//
// The server end of the Local Security Authority (Domain Policy) Remote Protocol can
// be implemented on a domain controller (DC), including primary domain controllers
// (PDCs), backup domain controllers (BDCs), global catalog servers (GC servers), and
// read-only domain controllers (RODCs), or on a non–domain controller. In the case
// of a DC, including PDCs, BDCs, GC servers, and RODCs, the server end of this protocol
// can be in one of the forest functional levels. The behavior of the server end of
// the Local Security Authority (Domain Policy) Remote Protocol is the same in these
// cases, except when noted in the message processing descriptions for the methods of
// this protocol. See sections 3.1.4.4.1, 3.1.4.4.3, 3.1.4.4.5, 3.1.4.7, 3.1.4.7.3,
// 3.1.4.7.4, 3.1.4.7.10, 3.1.4.7.14, and 3.1.4.7.16 for details.
//
// This protocol is a simple request/response-based RPC protocol. Typically, there are
// no long-lived sessions, although clients can cache the RPC connection and reuse it
// over time. A sample sequence of requests and responses is specified in section 4.
//
// It is helpful to consider two perspectives when understanding and implementing this
// protocol: an object-based perspective and a method-based perspective.
//
// The object-based perspective shows that the protocol exposes four main object abstractions:
// a policy object, an account object, a secret object, and a trusted domain object.
// A requester obtains a "handle" (an RPC context handle) to one of these objects and
// then performs one or more actions on the object. The following is a brief listing
// of methods that operate on each of the respective object types.
//
// Policy object:
//
// * LsarOpenPolicy3
//
// * LsarOpenPolicy2
//
// * LsarQueryInformationPolicy2
//
// * LsarSetInformationPolicy2
//
// * LsarClose
//
// * LsarQueryDomainInformationPolicy
//
// * LsarEnumeratePrivileges
//
// * LsarLookupPrivilegeName
//
// * LsarLookupPrivilegeValue
//
// * LsarLookupPrivilegeDisplayName
//
// * LsarSetDomainInformationPolicy
//
// * LsarQuerySecurityObject
//
// * LsarSetSecurityObject
//
// Account object:
//
// * LsarCreateAccount
//
// * LsarOpenAccount
//
// * LsarEnumerateAccounts
//
// * LsarClose
//
// * LsarDeleteObject
//
// * LsarSetSystemAccessAccount
//
// * LsarQuerySecurityObject
//
// * LsarAddAccountRights
//
// * LsarRemoveAccountRights
//
// * LsarAddPrivilegesToAccount
//
// * LsarRemovePrivilegesFromAccount
//
// * LsarEnumerateAccountsWithUserRight
//
// * LsarGetSystemAccessAccount
//
// * LsarSetSecurityObject
//
// * LsarEnumeratePrivilegesAccount
//
// * LsarEnumerateAccountRights
//
// Secret object:
//
// * LsarCreateSecret
//
// * LsarOpenSecret
//
// * LsarClose
//
// * LsarDeleteObject
//
// * LsarRetrievePrivateData
//
// * LsarStorePrivateData
//
// * LsarSetSecret
//
// * LsarQuerySecret
//
// * LsarQuerySecurityObject
//
// * LsarSetSecurityObject
//
// Trusted domain object:
//
// * LsarCreateTrustedDomainEx3
//
// * LsarCreateTrustedDomainEx2
//
// * LsarOpenTrustedDomain
//
// * LsarClose
//
// * LsarDeleteObject
//
// * LsarOpenTrustedDomainByName
//
// * LsarDeleteTrustedDomain
//
// * LsarEnumerateTrustedDomainsEx
//
// * LsarQueryInfoTrustedDomain
//
// * LsarSetInformationTrustedDomain
//
// * LsarQueryForestTrustInformation
//
// * LsarSetForestTrustInformation
//
// * LsarQueryTrustedDomainInfo
//
// * LsarSetTrustedDomainInfo
//
// * LsarQueryTrustedDomainInfoByName
//
// * LsarSetTrustedDomainInfoByName
//
// For example, to set a policy that controls the lifetime of Kerberos tickets, a requester
// opens a handle to the Policy object and updates the maximum service ticket age policy
// setting via a parameter called MaxServiceTicketAge. The call sequence from the requester
// appears as follows (with the parameter information removed for brevity):
//
// *
//
// Send LsarOpenPolicy3 request; receive LsarOpenPolicy3 reply.
//
// *
//
// Send LsarQueryDomainInformationPolicy request; receive LsarQueryDomainInformationPolicy
// reply.
//
// *
//
// Send LsarSetDomainInformationPolicy request; receive LsarSetDomainInformationPolicy
// reply.
//
// *
//
// Send LsarClose request; receive LsarClose reply.
//
// The following is a brief explanation of the call sequence:
//
// *
//
// Using the network address of a responder that implements this protocol, a requester
// makes an LsarOpenPolicy3 request to obtain a handle to the policy object. This handle
// is necessary to examine and manipulate domain ( 31ca2a31-0be4-4773-bcef-05ad6cd3ccfb#gt_b0276eb2-4e65-4cf1-a718-e0920a614aca
// ) policy information.
//
// *
//
// Using the handle returned from LsarOpenPolicy3, the requester makes an LsarQueryDomainInformationPolicy
// request to retrieve the current policy settings that affect Kerberos tickets.
//
// *
//
// After modifying the portions of the Kerberos ticket policy information to suit the
// requester, the requester makes the LsarSetDomainInformationPolicy request to set
// the policy to the new values.
//
// *
//
// The requester closes the policy handle returned from LsarOpenPolicy3. This releases
// responder resources associated with the handle.
//
// In the method-based perspective, there is a common set of operations for each object
// type. The operations fall into patterns. The following is a list of the patterns
// and associated methods, along with a description of the pattern.
//
// * *Open pattern:* This pattern returns an RPC context handle that references a specific
// object type. A requester uses this pattern by specifying a specific access for the
// handle in the request and using the returned handle to call other methods that require
// the returned handle and the associated access. For example, calling the LsarSetSecret
// method requires a secret object handle that has been opened with SECRET_WRITE access.
//
// LsarOpenPolicy3 is distinguished from the other methods in this pattern in two ways.
// First, the requestor calls this method before calling any other handle-based methods.
// Second, a network address, rather than a context handle, is required to indicate
// the responder.
//
// The following are the methods that follow the open pattern:
//
// * LsarOpenPolicy3
//
// * LsarOpenPolicy2
//
// * LsarOpenPolicy
//
// * LsarOpenAccount
//
// * LsarOpenSecret
//
// * LsarOpenTrustedDomain
//
// * LsarOpenTrustedDomainByName
//
// * *Enumerate pattern:* This pattern enables a requester to obtain a complete listing
// of all objects of a certain type (account or trusted domain ( 31ca2a31-0be4-4773-bcef-05ad6cd3ccfb#gt_f2f00d47-6cf2-4b32-b8f7-b63e38e2e9c4
// ) ) or to get all values of a certain type out of an object (for example, privileges
// ( 31ca2a31-0be4-4773-bcef-05ad6cd3ccfb#gt_d8092e10-b227-4b44-b015-511bb8178940 )
// known to the server).
//
// The following are the methods that follow the enumerate pattern:
//
// * LsarEnumerateTrustedDomainsEx
//
// * LsarEnumerateAccounts
//
// * LsarEnumeratePrivileges
//
// * LsarEnumeratePrivilegesAccount
//
// * LsarEnumerateAccountRights
//
// * LsarEnumerateAccountsWithUserRight
//
// * *Create pattern:* Methods in this pattern enable specified objects to be created.
// A handle to the newly created object is also returned.
//
// The following are the methods that follow the create pattern:
//
// * LsarCreateAccount
//
// * LsarCreateSecret
//
// * LsarCreateTrustedDomainEx2
//
// * LsarCreateTrustedDomainEx3
//
// * *Query pattern:* This pattern enables specified attributes of an object to be returned.
// The requester indicates which attributes to return by specifying an "information
// class". This is an enumeration that the responder understands and translates to a
// specific structure to return (the structure contains the attributes indicated by
// the information class).
//
// For example, to retrieve the name of a trusted domain, a requester would specify
// the information level "TrustedDomainNameInformation" to the LsarQueryTrustedDomainInfo
// method.
//
// The following are the methods that follow the query pattern:
//
// * LsarQueryDomainInformationPolicy
//
// * LsarQueryForestTrustInformation
//
// * LsarQueryInformationPolicy2
//
// * LsarQuerySecret
//
// * LsarQueryTrustedDomainInfo
//
// * LsarQueryTrustedDomainInfoByName
//
// * LsarQueryInfoTrustedDomain
//
// * *Set pattern:* This pattern enables specified object attributes to be set. The
// requester makes a request for which attributes to update by specifying an "information
// class". Similar to the Query pattern, this information level allows the caller to
// specify to the responder which attributes are being sent in the request.
//
// The following are the methods that follow the set pattern:
//
// * LsarSetDomainInformationPolicy
//
// * LsarSetForestTrustInformation
//
// * LsarSetInformationPolicy2
//
// * LsarSetSecret
//
// * LsarAddPrivilegesToAccount
//
// * LsarRemovePrivilegesFromAccount
//
// * LsarAddAccountRights
//
// * LsarRemoveAccountRights
//
// * *Delete pattern:* This pattern enables a requester to delete a specified object.
//
// The following are the methods that follow the delete pattern:
//
// * LsarDeleteObject
//
// * LsarDeleteTrustedDomain
//
// * *Lookup pattern:* This pattern enables a caller to translate between different
// representations of an entity (in the case of this protocol, names and identifiers
// of privileges).
//
// The following are the methods that follow the lookup pattern:
//
// * LsarLookupPrivilegeName
//
// * LsarLookupPrivilegeValue
//
// * LsarLookupPrivilegeDisplayName
//
// * *Security pattern:* This pattern enables a caller to specify or query the access
// control at the level of individual objects.
//
// The following are the methods that follow the security pattern:
//
// * LsarSetSecurityObject
//
// * LsarQuerySecurityObject
//
// * *Miscellaneous:* The following method does not fall into a general pattern. A brief
// description is given here. See the message processing section for details.
//
// LsarClose: This method releases responder resources associated with the RPC context
// handle that is passed as a parameter.
package lsad

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

var (
	// import guard
	GoPackage = "lsad"
)
