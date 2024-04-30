package samr

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

// samr server interface.
type SamrServer interface {

	// The SamrConnect method returns a handle to a server object.
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)

	// The SamrCloseHandle method closes (that is, releases server-side resources used by)
	// any context handle obtained from this RPC interface.
	CloseHandle(context.Context, *CloseHandleRequest) (*CloseHandleResponse, error)

	// The SamrSetSecurityObject method sets the access control on a server, domain, user,
	// group, or alias object.
	SetSecurityObject(context.Context, *SetSecurityObjectRequest) (*SetSecurityObjectResponse, error)

	// The SamrQuerySecurityObject method queries the access control on a server, domain,
	// user, group, or alias object.
	QuerySecurityObject(context.Context, *QuerySecurityObjectRequest) (*QuerySecurityObjectResponse, error)

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// The SamrLookupDomainInSamServer method obtains the SID of a domain object, given
	// the object's name.
	LookupDomainInSAMServer(context.Context, *LookupDomainInSAMServerRequest) (*LookupDomainInSAMServerResponse, error)

	// The SamrEnumerateDomainsInSamServer method obtains a listing of all domains hosted
	// by the server side of this protocol.
	EnumerateDomainsInSAMServer(context.Context, *EnumerateDomainsInSAMServerRequest) (*EnumerateDomainsInSAMServerResponse, error)

	// The SamrOpenDomain method obtains a handle to a domain object, given a SID.
	OpenDomain(context.Context, *OpenDomainRequest) (*OpenDomainResponse, error)

	// The SamrQueryInformationDomain method obtains attributes from a domain object.
	//
	// See the description of SamrQueryInformationDomain2 (section 3.1.5.5.1) for details,
	// because the method interface arguments and message processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	QueryInformationDomain(context.Context, *QueryInformationDomainRequest) (*QueryInformationDomainResponse, error)

	// The SamrSetInformationDomain method updates attributes on a domain object.
	SetInformationDomain(context.Context, *SetInformationDomainRequest) (*SetInformationDomainResponse, error)

	// The SamrCreateGroupInDomain method creates a group object within a domain.
	CreateGroupInDomain(context.Context, *CreateGroupInDomainRequest) (*CreateGroupInDomainResponse, error)

	// The SamrEnumerateGroupsInDomain method enumerates all groups.
	EnumerateGroupsInDomain(context.Context, *EnumerateGroupsInDomainRequest) (*EnumerateGroupsInDomainResponse, error)

	// The SamrCreateUserInDomain method creates a user.
	CreateUserInDomain(context.Context, *CreateUserInDomainRequest) (*CreateUserInDomainResponse, error)

	// The SamrEnumerateUsersInDomain method enumerates all users.
	EnumerateUsersInDomain(context.Context, *EnumerateUsersInDomainRequest) (*EnumerateUsersInDomainResponse, error)

	// The SamrCreateAliasInDomain method creates an alias.
	CreateAliasInDomain(context.Context, *CreateAliasInDomainRequest) (*CreateAliasInDomainResponse, error)

	// The SamrEnumerateAliasesInDomain method enumerates all aliases.
	EnumerateAliasesInDomain(context.Context, *EnumerateAliasesInDomainRequest) (*EnumerateAliasesInDomainResponse, error)

	// The SamrGetAliasMembership method obtains the union of all aliases that a given set
	// of SIDs is a member of.
	GetAliasMembership(context.Context, *GetAliasMembershipRequest) (*GetAliasMembershipResponse, error)

	// The SamrLookupNamesInDomain method translates a set of account names into a set of
	// RIDs.
	LookupNamesInDomain(context.Context, *LookupNamesInDomainRequest) (*LookupNamesInDomainResponse, error)

	// The SamrLookupIdsInDomain method translates a set of RIDs into account names.
	LookupIDsInDomain(context.Context, *LookupIDsInDomainRequest) (*LookupIDsInDomainResponse, error)

	// The SamrOpenGroup method obtains a handle to a group, given a RID.
	OpenGroup(context.Context, *OpenGroupRequest) (*OpenGroupResponse, error)

	// The SamrQueryInformationGroup method obtains attributes from a group object.
	QueryInformationGroup(context.Context, *QueryInformationGroupRequest) (*QueryInformationGroupResponse, error)

	// The SamrSetInformationGroup method updates attributes on a group object.
	SetInformationGroup(context.Context, *SetInformationGroupRequest) (*SetInformationGroupResponse, error)

	// The SamrAddMemberToGroup method adds a member to a group.
	AddMemberToGroup(context.Context, *AddMemberToGroupRequest) (*AddMemberToGroupResponse, error)

	// The SamrDeleteGroup method removes a group object.
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error)

	// The SamrRemoveMemberFromGroup method removes a member from a group.
	RemoveMemberFromGroup(context.Context, *RemoveMemberFromGroupRequest) (*RemoveMemberFromGroupResponse, error)

	// The SamrGetMembersInGroup method reads the members of a group.
	GetMembersInGroup(context.Context, *GetMembersInGroupRequest) (*GetMembersInGroupResponse, error)

	// The SamrSetMemberAttributesOfGroup method sets the attributes of a member relationship.
	SetMemberAttributesOfGroup(context.Context, *SetMemberAttributesOfGroupRequest) (*SetMemberAttributesOfGroupResponse, error)

	// The SamrOpenAlias method obtains a handle to an alias, given a RID.
	OpenAlias(context.Context, *OpenAliasRequest) (*OpenAliasResponse, error)

	// The SamrQueryInformationAlias method obtains attributes from an alias object.
	QueryInformationAlias(context.Context, *QueryInformationAliasRequest) (*QueryInformationAliasResponse, error)

	// The SamrSetInformationAlias method updates attributes on an alias object.
	SetInformationAlias(context.Context, *SetInformationAliasRequest) (*SetInformationAliasResponse, error)

	// The SamrDeleteAlias method removes an alias object.
	DeleteAlias(context.Context, *DeleteAliasRequest) (*DeleteAliasResponse, error)

	// The SamrAddMemberToAlias method adds a member to an alias.
	AddMemberToAlias(context.Context, *AddMemberToAliasRequest) (*AddMemberToAliasResponse, error)

	// The SamrRemoveMemberFromAlias method removes a member from an alias.
	RemoveMemberFromAlias(context.Context, *RemoveMemberFromAliasRequest) (*RemoveMemberFromAliasResponse, error)

	// The SamrGetMembersInAlias method obtains the membership list of an alias.
	GetMembersInAlias(context.Context, *GetMembersInAliasRequest) (*GetMembersInAliasResponse, error)

	// The SamrOpenUser method obtains a handle to a user, given a RID.
	OpenUser(context.Context, *OpenUserRequest) (*OpenUserResponse, error)

	// The SamrDeleteUser method removes a user object.
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)

	// The SamrQueryInformationUser method obtains attributes from a user object.
	//
	// See the description of SamrQueryInformationUser2 (section 3.1.5.5.5) for details,
	// because the method interface arguments and message processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	QueryInformationUser(context.Context, *QueryInformationUserRequest) (*QueryInformationUserResponse, error)

	// The SamrSetInformationUser method updates attributes on a user object.
	//
	// See the description of SamrSetInformationUser2 (section 3.1.5.6.4) for details, because
	// the method interface arguments and message processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	SetInformationUser(context.Context, *SetInformationUserRequest) (*SetInformationUserResponse, error)

	// The SamrChangePasswordUser method changes the password of a user object.
	ChangePasswordUser(context.Context, *ChangePasswordUserRequest) (*ChangePasswordUserResponse, error)

	// The SamrGetGroupsForUser method obtains a listing of groups that a user is a member
	// of.
	GetGroupsForUser(context.Context, *GetGroupsForUserRequest) (*GetGroupsForUserResponse, error)

	// The SamrQueryDisplayInformation method obtains a list of accounts in ascending name-sorted
	// order, starting at a specified index.
	//
	// See the description of SamrQueryDisplayInformation3 (section 3.1.5.3.1) for details,
	// because the method interface arguments and message processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	QueryDisplayInformation(context.Context, *QueryDisplayInformationRequest) (*QueryDisplayInformationResponse, error)

	// The SamrGetDisplayEnumerationIndex method obtains an index into an ascending account-name–sorted
	// list of accounts.
	//
	// See the description of SamrGetDisplayEnumerationIndex2 (section 3.1.5.3.4) for details,
	// because the method-interface arguments and processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	GetDisplayEnumerationIndex(context.Context, *GetDisplayEnumerationIndexRequest) (*GetDisplayEnumerationIndexResponse, error)

	// Opnum42NotUsedOnWire operation.
	// Opnum42NotUsedOnWire

	// Opnum43NotUsedOnWire operation.
	// Opnum43NotUsedOnWire

	// The SamrGetUserDomainPasswordInformation method obtains select password policy information
	// (without requiring a domain handle).
	GetUserDomainPasswordInformation(context.Context, *GetUserDomainPasswordInformationRequest) (*GetUserDomainPasswordInformationResponse, error)

	// The SamrRemoveMemberFromForeignDomain method removes a member from all aliases.
	RemoveMemberFromForeignDomain(context.Context, *RemoveMemberFromForeignDomainRequest) (*RemoveMemberFromForeignDomainResponse, error)

	// The SamrQueryInformationDomain2 method obtains attributes from a domain object.
	QueryInformationDomain2(context.Context, *QueryInformationDomain2Request) (*QueryInformationDomain2Response, error)

	// The SamrQueryInformationUser2 method obtains attributes from a user object.
	QueryInformationUser2(context.Context, *QueryInformationUser2Request) (*QueryInformationUser2Response, error)

	// The SamrQueryDisplayInformation2 method obtains a list of accounts in ascending name-sorted
	// order, starting at a specified index.
	//
	// See the description of SamrQueryDisplayInformation3 (section 3.1.5.3.1) for details,
	// because the method-interface arguments and message processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	QueryDisplayInformation2(context.Context, *QueryDisplayInformation2Request) (*QueryDisplayInformation2Response, error)

	// The SamrGetDisplayEnumerationIndex2 method obtains an index into an ascending account-name–sorted
	// list of accounts, such that the index is the position in the list of the accounts
	// whose account name best matches a client-provided string.
	GetDisplayEnumerationIndex2(context.Context, *GetDisplayEnumerationIndex2Request) (*GetDisplayEnumerationIndex2Response, error)

	// The SamrCreateUser2InDomain method creates a user.
	CreateUser2InDomain(context.Context, *CreateUser2InDomainRequest) (*CreateUser2InDomainResponse, error)

	// The SamrQueryDisplayInformation3 method obtains a listing of accounts in ascending
	// name-sorted order, starting at a specified index.
	QueryDisplayInformation3(context.Context, *QueryDisplayInformation3Request) (*QueryDisplayInformation3Response, error)

	// The SamrAddMultipleMembersToAlias method adds multiple members to an alias.
	AddMultipleMembersToAlias(context.Context, *AddMultipleMembersToAliasRequest) (*AddMultipleMembersToAliasResponse, error)

	// The SamrRemoveMultipleMembersFromAlias method removes multiple members from an alias.
	RemoveMultipleMembersFromAlias(context.Context, *RemoveMultipleMembersFromAliasRequest) (*RemoveMultipleMembersFromAliasResponse, error)

	// The SamrOemChangePasswordUser2 method changes a user's password.
	//
	// );
	OEMChangePasswordUser2(context.Context, *OEMChangePasswordUser2Request) (*OEMChangePasswordUser2Response, error)

	// The SamrUnicodeChangePasswordUser2 method changes a user account's password.
	UnicodeChangePasswordUser2(context.Context, *UnicodeChangePasswordUser2Request) (*UnicodeChangePasswordUser2Response, error)

	// The SamrGetDomainPasswordInformation method obtains select password policy information
	// (without authenticating to the server).
	GetDomainPasswordInformation(context.Context, *GetDomainPasswordInformationRequest) (*GetDomainPasswordInformationResponse, error)

	// The SamrConnect2 method returns a handle to a server object.
	Connect2(context.Context, *Connect2Request) (*Connect2Response, error)

	// The SamrSetInformationUser2 method updates attributes on a user object.
	SetInformationUser2(context.Context, *SetInformationUser2Request) (*SetInformationUser2Response, error)

	// Opnum59NotUsedOnWire operation.
	// Opnum59NotUsedOnWire

	// Opnum60NotUsedOnWire operation.
	// Opnum60NotUsedOnWire

	// Opnum61NotUsedOnWire operation.
	// Opnum61NotUsedOnWire

	// The SamrConnect4 method obtains a handle to a server object.
	Connect4(context.Context, *Connect4Request) (*Connect4Response, error)

	// Opnum63NotUsedOnWire operation.
	// Opnum63NotUsedOnWire

	// The SamrConnect5 method obtains a handle to a server object.
	Connect5(context.Context, *Connect5Request) (*Connect5Response, error)

	// The SamrRidToSid method obtains the SID of an account, given a RID.
	RIDToSID(context.Context, *RIDToSIDRequest) (*RIDToSIDResponse, error)

	// The SamrSetDSRMPassword method sets a local recovery password.
	SetDSRMPassword(context.Context, *SetDSRMPasswordRequest) (*SetDSRMPasswordResponse, error)

	// The SamrValidatePassword method validates an application password against the locally
	// stored policy.
	ValidatePassword(context.Context, *ValidatePasswordRequest) (*ValidatePasswordResponse, error)

	// Opnum68NotUsedOnWire operation.
	// Opnum68NotUsedOnWire

	// Opnum69NotUsedOnWire operation.
	// Opnum69NotUsedOnWire
}

func RegisterSamrServer(conn dcerpc.Conn, o SamrServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSamrServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SamrSyntaxV1_0))...)
}

func NewSamrServerHandle(o SamrServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SamrServerHandle(ctx, o, opNum, r)
	}
}

func SamrServerHandle(ctx context.Context, o SamrServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // SamrConnect
		in := &ConnectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Connect(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // SamrCloseHandle
		in := &CloseHandleRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseHandle(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // SamrSetSecurityObject
		in := &SetSecurityObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSecurityObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // SamrQuerySecurityObject
		in := &QuerySecurityObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QuerySecurityObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // SamrLookupDomainInSamServer
		in := &LookupDomainInSAMServerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupDomainInSAMServer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // SamrEnumerateDomainsInSamServer
		in := &EnumerateDomainsInSAMServerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateDomainsInSAMServer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // SamrOpenDomain
		in := &OpenDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // SamrQueryInformationDomain
		in := &QueryInformationDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInformationDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // SamrSetInformationDomain
		in := &SetInformationDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetInformationDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // SamrCreateGroupInDomain
		in := &CreateGroupInDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateGroupInDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // SamrEnumerateGroupsInDomain
		in := &EnumerateGroupsInDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateGroupsInDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // SamrCreateUserInDomain
		in := &CreateUserInDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateUserInDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // SamrEnumerateUsersInDomain
		in := &EnumerateUsersInDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateUsersInDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // SamrCreateAliasInDomain
		in := &CreateAliasInDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateAliasInDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // SamrEnumerateAliasesInDomain
		in := &EnumerateAliasesInDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateAliasesInDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // SamrGetAliasMembership
		in := &GetAliasMembershipRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAliasMembership(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // SamrLookupNamesInDomain
		in := &LookupNamesInDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupNamesInDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // SamrLookupIdsInDomain
		in := &LookupIDsInDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupIDsInDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // SamrOpenGroup
		in := &OpenGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // SamrQueryInformationGroup
		in := &QueryInformationGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInformationGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // SamrSetInformationGroup
		in := &SetInformationGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetInformationGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // SamrAddMemberToGroup
		in := &AddMemberToGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddMemberToGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // SamrDeleteGroup
		in := &DeleteGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // SamrRemoveMemberFromGroup
		in := &RemoveMemberFromGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveMemberFromGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // SamrGetMembersInGroup
		in := &GetMembersInGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMembersInGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // SamrSetMemberAttributesOfGroup
		in := &SetMemberAttributesOfGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMemberAttributesOfGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // SamrOpenAlias
		in := &OpenAliasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenAlias(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // SamrQueryInformationAlias
		in := &QueryInformationAliasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInformationAlias(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // SamrSetInformationAlias
		in := &SetInformationAliasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetInformationAlias(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // SamrDeleteAlias
		in := &DeleteAliasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteAlias(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // SamrAddMemberToAlias
		in := &AddMemberToAliasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddMemberToAlias(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // SamrRemoveMemberFromAlias
		in := &RemoveMemberFromAliasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveMemberFromAlias(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // SamrGetMembersInAlias
		in := &GetMembersInAliasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMembersInAlias(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // SamrOpenUser
		in := &OpenUserRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenUser(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // SamrDeleteUser
		in := &DeleteUserRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteUser(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // SamrQueryInformationUser
		in := &QueryInformationUserRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInformationUser(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // SamrSetInformationUser
		in := &SetInformationUserRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetInformationUser(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // SamrChangePasswordUser
		in := &ChangePasswordUserRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ChangePasswordUser(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // SamrGetGroupsForUser
		in := &GetGroupsForUserRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetGroupsForUser(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // SamrQueryDisplayInformation
		in := &QueryDisplayInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryDisplayInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // SamrGetDisplayEnumerationIndex
		in := &GetDisplayEnumerationIndexRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDisplayEnumerationIndex(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // Opnum42NotUsedOnWire
		// Opnum42NotUsedOnWire
		return nil, nil
	case 43: // Opnum43NotUsedOnWire
		// Opnum43NotUsedOnWire
		return nil, nil
	case 44: // SamrGetUserDomainPasswordInformation
		in := &GetUserDomainPasswordInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUserDomainPasswordInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // SamrRemoveMemberFromForeignDomain
		in := &RemoveMemberFromForeignDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveMemberFromForeignDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // SamrQueryInformationDomain2
		in := &QueryInformationDomain2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInformationDomain2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // SamrQueryInformationUser2
		in := &QueryInformationUser2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInformationUser2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // SamrQueryDisplayInformation2
		in := &QueryDisplayInformation2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryDisplayInformation2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // SamrGetDisplayEnumerationIndex2
		in := &GetDisplayEnumerationIndex2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDisplayEnumerationIndex2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // SamrCreateUser2InDomain
		in := &CreateUser2InDomainRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateUser2InDomain(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // SamrQueryDisplayInformation3
		in := &QueryDisplayInformation3Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryDisplayInformation3(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // SamrAddMultipleMembersToAlias
		in := &AddMultipleMembersToAliasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddMultipleMembersToAlias(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 53: // SamrRemoveMultipleMembersFromAlias
		in := &RemoveMultipleMembersFromAliasRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveMultipleMembersFromAlias(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // SamrOemChangePasswordUser2
		in := &OEMChangePasswordUser2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OEMChangePasswordUser2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 55: // SamrUnicodeChangePasswordUser2
		in := &UnicodeChangePasswordUser2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.UnicodeChangePasswordUser2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 56: // SamrGetDomainPasswordInformation
		in := &GetDomainPasswordInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDomainPasswordInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // SamrConnect2
		in := &Connect2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Connect2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 58: // SamrSetInformationUser2
		in := &SetInformationUser2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetInformationUser2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 59: // Opnum59NotUsedOnWire
		// Opnum59NotUsedOnWire
		return nil, nil
	case 60: // Opnum60NotUsedOnWire
		// Opnum60NotUsedOnWire
		return nil, nil
	case 61: // Opnum61NotUsedOnWire
		// Opnum61NotUsedOnWire
		return nil, nil
	case 62: // SamrConnect4
		in := &Connect4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Connect4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 63: // Opnum63NotUsedOnWire
		// Opnum63NotUsedOnWire
		return nil, nil
	case 64: // SamrConnect5
		in := &Connect5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Connect5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 65: // SamrRidToSid
		in := &RIDToSIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RIDToSID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 66: // SamrSetDSRMPassword
		in := &SetDSRMPasswordRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDSRMPassword(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 67: // SamrValidatePassword
		in := &ValidatePasswordRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ValidatePassword(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 68: // Opnum68NotUsedOnWire
		// Opnum68NotUsedOnWire
		return nil, nil
	case 69: // Opnum69NotUsedOnWire
		// Opnum69NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
