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
		op := &xxx_ConnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Connect(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // SamrCloseHandle
		op := &xxx_CloseHandleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseHandleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseHandle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // SamrSetSecurityObject
		op := &xxx_SetSecurityObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSecurityObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSecurityObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // SamrQuerySecurityObject
		op := &xxx_QuerySecurityObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QuerySecurityObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QuerySecurityObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // SamrLookupDomainInSamServer
		op := &xxx_LookupDomainInSAMServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupDomainInSAMServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupDomainInSAMServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // SamrEnumerateDomainsInSamServer
		op := &xxx_EnumerateDomainsInSAMServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumerateDomainsInSAMServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumerateDomainsInSAMServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // SamrOpenDomain
		op := &xxx_OpenDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SamrQueryInformationDomain
		op := &xxx_QueryInformationDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryInformationDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryInformationDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // SamrSetInformationDomain
		op := &xxx_SetInformationDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInformationDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInformationDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // SamrCreateGroupInDomain
		op := &xxx_CreateGroupInDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateGroupInDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateGroupInDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // SamrEnumerateGroupsInDomain
		op := &xxx_EnumerateGroupsInDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumerateGroupsInDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumerateGroupsInDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // SamrCreateUserInDomain
		op := &xxx_CreateUserInDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateUserInDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateUserInDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // SamrEnumerateUsersInDomain
		op := &xxx_EnumerateUsersInDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumerateUsersInDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumerateUsersInDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // SamrCreateAliasInDomain
		op := &xxx_CreateAliasInDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateAliasInDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateAliasInDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // SamrEnumerateAliasesInDomain
		op := &xxx_EnumerateAliasesInDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumerateAliasesInDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumerateAliasesInDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // SamrGetAliasMembership
		op := &xxx_GetAliasMembershipOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAliasMembershipRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAliasMembership(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // SamrLookupNamesInDomain
		op := &xxx_LookupNamesInDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupNamesInDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupNamesInDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // SamrLookupIdsInDomain
		op := &xxx_LookupIDsInDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupIDsInDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupIDsInDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // SamrOpenGroup
		op := &xxx_OpenGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // SamrQueryInformationGroup
		op := &xxx_QueryInformationGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryInformationGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryInformationGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // SamrSetInformationGroup
		op := &xxx_SetInformationGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInformationGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInformationGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // SamrAddMemberToGroup
		op := &xxx_AddMemberToGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddMemberToGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddMemberToGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // SamrDeleteGroup
		op := &xxx_DeleteGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // SamrRemoveMemberFromGroup
		op := &xxx_RemoveMemberFromGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveMemberFromGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveMemberFromGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // SamrGetMembersInGroup
		op := &xxx_GetMembersInGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMembersInGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMembersInGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // SamrSetMemberAttributesOfGroup
		op := &xxx_SetMemberAttributesOfGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMemberAttributesOfGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMemberAttributesOfGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // SamrOpenAlias
		op := &xxx_OpenAliasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenAliasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenAlias(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // SamrQueryInformationAlias
		op := &xxx_QueryInformationAliasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryInformationAliasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryInformationAlias(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // SamrSetInformationAlias
		op := &xxx_SetInformationAliasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInformationAliasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInformationAlias(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // SamrDeleteAlias
		op := &xxx_DeleteAliasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAliasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAlias(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // SamrAddMemberToAlias
		op := &xxx_AddMemberToAliasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddMemberToAliasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddMemberToAlias(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // SamrRemoveMemberFromAlias
		op := &xxx_RemoveMemberFromAliasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveMemberFromAliasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveMemberFromAlias(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // SamrGetMembersInAlias
		op := &xxx_GetMembersInAliasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMembersInAliasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMembersInAlias(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // SamrOpenUser
		op := &xxx_OpenUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // SamrDeleteUser
		op := &xxx_DeleteUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // SamrQueryInformationUser
		op := &xxx_QueryInformationUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryInformationUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryInformationUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // SamrSetInformationUser
		op := &xxx_SetInformationUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInformationUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInformationUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // SamrChangePasswordUser
		op := &xxx_ChangePasswordUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangePasswordUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangePasswordUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // SamrGetGroupsForUser
		op := &xxx_GetGroupsForUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetGroupsForUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetGroupsForUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // SamrQueryDisplayInformation
		op := &xxx_QueryDisplayInformationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryDisplayInformationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryDisplayInformation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // SamrGetDisplayEnumerationIndex
		op := &xxx_GetDisplayEnumerationIndexOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDisplayEnumerationIndexRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDisplayEnumerationIndex(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // Opnum42NotUsedOnWire
		// Opnum42NotUsedOnWire
		return nil, nil
	case 43: // Opnum43NotUsedOnWire
		// Opnum43NotUsedOnWire
		return nil, nil
	case 44: // SamrGetUserDomainPasswordInformation
		op := &xxx_GetUserDomainPasswordInformationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserDomainPasswordInformationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserDomainPasswordInformation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // SamrRemoveMemberFromForeignDomain
		op := &xxx_RemoveMemberFromForeignDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveMemberFromForeignDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveMemberFromForeignDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // SamrQueryInformationDomain2
		op := &xxx_QueryInformationDomain2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryInformationDomain2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryInformationDomain2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // SamrQueryInformationUser2
		op := &xxx_QueryInformationUser2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryInformationUser2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryInformationUser2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // SamrQueryDisplayInformation2
		op := &xxx_QueryDisplayInformation2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryDisplayInformation2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryDisplayInformation2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // SamrGetDisplayEnumerationIndex2
		op := &xxx_GetDisplayEnumerationIndex2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDisplayEnumerationIndex2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDisplayEnumerationIndex2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // SamrCreateUser2InDomain
		op := &xxx_CreateUser2InDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateUser2InDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateUser2InDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // SamrQueryDisplayInformation3
		op := &xxx_QueryDisplayInformation3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryDisplayInformation3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryDisplayInformation3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // SamrAddMultipleMembersToAlias
		op := &xxx_AddMultipleMembersToAliasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddMultipleMembersToAliasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddMultipleMembersToAlias(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // SamrRemoveMultipleMembersFromAlias
		op := &xxx_RemoveMultipleMembersFromAliasOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveMultipleMembersFromAliasRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveMultipleMembersFromAlias(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // SamrOemChangePasswordUser2
		op := &xxx_OEMChangePasswordUser2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OEMChangePasswordUser2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OEMChangePasswordUser2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // SamrUnicodeChangePasswordUser2
		op := &xxx_UnicodeChangePasswordUser2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnicodeChangePasswordUser2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UnicodeChangePasswordUser2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // SamrGetDomainPasswordInformation
		op := &xxx_GetDomainPasswordInformationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDomainPasswordInformationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDomainPasswordInformation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // SamrConnect2
		op := &xxx_Connect2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Connect2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Connect2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // SamrSetInformationUser2
		op := &xxx_SetInformationUser2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInformationUser2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInformationUser2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
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
		op := &xxx_Connect4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Connect4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Connect4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 63: // Opnum63NotUsedOnWire
		// Opnum63NotUsedOnWire
		return nil, nil
	case 64: // SamrConnect5
		op := &xxx_Connect5Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Connect5Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Connect5(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 65: // SamrRidToSid
		op := &xxx_RIDToSIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RIDToSIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RIDToSID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 66: // SamrSetDSRMPassword
		op := &xxx_SetDSRMPasswordOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDSRMPasswordRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDSRMPassword(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 67: // SamrValidatePassword
		op := &xxx_ValidatePasswordOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ValidatePasswordRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ValidatePassword(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 68: // Opnum68NotUsedOnWire
		// Opnum68NotUsedOnWire
		return nil, nil
	case 69: // Opnum69NotUsedOnWire
		// Opnum69NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented samr
type UnimplementedSamrServer struct {
}

func (UnimplementedSamrServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) CloseHandle(context.Context, *CloseHandleRequest) (*CloseHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) SetSecurityObject(context.Context, *SetSecurityObjectRequest) (*SetSecurityObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) QuerySecurityObject(context.Context, *QuerySecurityObjectRequest) (*QuerySecurityObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) LookupDomainInSAMServer(context.Context, *LookupDomainInSAMServerRequest) (*LookupDomainInSAMServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) EnumerateDomainsInSAMServer(context.Context, *EnumerateDomainsInSAMServerRequest) (*EnumerateDomainsInSAMServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) OpenDomain(context.Context, *OpenDomainRequest) (*OpenDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) QueryInformationDomain(context.Context, *QueryInformationDomainRequest) (*QueryInformationDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) SetInformationDomain(context.Context, *SetInformationDomainRequest) (*SetInformationDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) CreateGroupInDomain(context.Context, *CreateGroupInDomainRequest) (*CreateGroupInDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) EnumerateGroupsInDomain(context.Context, *EnumerateGroupsInDomainRequest) (*EnumerateGroupsInDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) CreateUserInDomain(context.Context, *CreateUserInDomainRequest) (*CreateUserInDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) EnumerateUsersInDomain(context.Context, *EnumerateUsersInDomainRequest) (*EnumerateUsersInDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) CreateAliasInDomain(context.Context, *CreateAliasInDomainRequest) (*CreateAliasInDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) EnumerateAliasesInDomain(context.Context, *EnumerateAliasesInDomainRequest) (*EnumerateAliasesInDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) GetAliasMembership(context.Context, *GetAliasMembershipRequest) (*GetAliasMembershipResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) LookupNamesInDomain(context.Context, *LookupNamesInDomainRequest) (*LookupNamesInDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) LookupIDsInDomain(context.Context, *LookupIDsInDomainRequest) (*LookupIDsInDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) OpenGroup(context.Context, *OpenGroupRequest) (*OpenGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) QueryInformationGroup(context.Context, *QueryInformationGroupRequest) (*QueryInformationGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) SetInformationGroup(context.Context, *SetInformationGroupRequest) (*SetInformationGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) AddMemberToGroup(context.Context, *AddMemberToGroupRequest) (*AddMemberToGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) RemoveMemberFromGroup(context.Context, *RemoveMemberFromGroupRequest) (*RemoveMemberFromGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) GetMembersInGroup(context.Context, *GetMembersInGroupRequest) (*GetMembersInGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) SetMemberAttributesOfGroup(context.Context, *SetMemberAttributesOfGroupRequest) (*SetMemberAttributesOfGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) OpenAlias(context.Context, *OpenAliasRequest) (*OpenAliasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) QueryInformationAlias(context.Context, *QueryInformationAliasRequest) (*QueryInformationAliasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) SetInformationAlias(context.Context, *SetInformationAliasRequest) (*SetInformationAliasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) DeleteAlias(context.Context, *DeleteAliasRequest) (*DeleteAliasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) AddMemberToAlias(context.Context, *AddMemberToAliasRequest) (*AddMemberToAliasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) RemoveMemberFromAlias(context.Context, *RemoveMemberFromAliasRequest) (*RemoveMemberFromAliasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) GetMembersInAlias(context.Context, *GetMembersInAliasRequest) (*GetMembersInAliasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) OpenUser(context.Context, *OpenUserRequest) (*OpenUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) QueryInformationUser(context.Context, *QueryInformationUserRequest) (*QueryInformationUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) SetInformationUser(context.Context, *SetInformationUserRequest) (*SetInformationUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) ChangePasswordUser(context.Context, *ChangePasswordUserRequest) (*ChangePasswordUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) GetGroupsForUser(context.Context, *GetGroupsForUserRequest) (*GetGroupsForUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) QueryDisplayInformation(context.Context, *QueryDisplayInformationRequest) (*QueryDisplayInformationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) GetDisplayEnumerationIndex(context.Context, *GetDisplayEnumerationIndexRequest) (*GetDisplayEnumerationIndexResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) GetUserDomainPasswordInformation(context.Context, *GetUserDomainPasswordInformationRequest) (*GetUserDomainPasswordInformationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) RemoveMemberFromForeignDomain(context.Context, *RemoveMemberFromForeignDomainRequest) (*RemoveMemberFromForeignDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) QueryInformationDomain2(context.Context, *QueryInformationDomain2Request) (*QueryInformationDomain2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) QueryInformationUser2(context.Context, *QueryInformationUser2Request) (*QueryInformationUser2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) QueryDisplayInformation2(context.Context, *QueryDisplayInformation2Request) (*QueryDisplayInformation2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) GetDisplayEnumerationIndex2(context.Context, *GetDisplayEnumerationIndex2Request) (*GetDisplayEnumerationIndex2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) CreateUser2InDomain(context.Context, *CreateUser2InDomainRequest) (*CreateUser2InDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) QueryDisplayInformation3(context.Context, *QueryDisplayInformation3Request) (*QueryDisplayInformation3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) AddMultipleMembersToAlias(context.Context, *AddMultipleMembersToAliasRequest) (*AddMultipleMembersToAliasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) RemoveMultipleMembersFromAlias(context.Context, *RemoveMultipleMembersFromAliasRequest) (*RemoveMultipleMembersFromAliasResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) OEMChangePasswordUser2(context.Context, *OEMChangePasswordUser2Request) (*OEMChangePasswordUser2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) UnicodeChangePasswordUser2(context.Context, *UnicodeChangePasswordUser2Request) (*UnicodeChangePasswordUser2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) GetDomainPasswordInformation(context.Context, *GetDomainPasswordInformationRequest) (*GetDomainPasswordInformationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) Connect2(context.Context, *Connect2Request) (*Connect2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) SetInformationUser2(context.Context, *SetInformationUser2Request) (*SetInformationUser2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) Connect4(context.Context, *Connect4Request) (*Connect4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) Connect5(context.Context, *Connect5Request) (*Connect5Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) RIDToSID(context.Context, *RIDToSIDRequest) (*RIDToSIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) SetDSRMPassword(context.Context, *SetDSRMPasswordRequest) (*SetDSRMPasswordResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedSamrServer) ValidatePassword(context.Context, *ValidatePasswordRequest) (*ValidatePasswordResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ SamrServer = (*UnimplementedSamrServer)(nil)
