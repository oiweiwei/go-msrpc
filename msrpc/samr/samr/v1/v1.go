package samr

import (
	"context"
	"encoding/hex"
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
	_ = hex.DecodeString
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "samr"
)

var (
	// Syntax UUID
	SamrSyntaxUUID = &uuid.UUID{TimeLow: 0x12345778, TimeMid: 0x1234, TimeHiAndVersion: 0xabcd, ClockSeqHiAndReserved: 0xef, ClockSeqLow: 0x0, Node: [6]uint8{0x1, 0x23, 0x45, 0x67, 0x89, 0xac}}
	// Syntax ID
	SamrSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: SamrSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// samr interface.
type SamrClient interface {

	// The SamrConnect method returns a handle to a server object.
	Connect(context.Context, *ConnectRequest, ...dcerpc.CallOption) (*ConnectResponse, error)

	// The SamrCloseHandle method closes (that is, releases server-side resources used by)
	// any context handle obtained from this RPC interface.
	CloseHandle(context.Context, *CloseHandleRequest, ...dcerpc.CallOption) (*CloseHandleResponse, error)

	// The SamrSetSecurityObject method sets the access control on a server, domain, user,
	// group, or alias object.
	SetSecurityObject(context.Context, *SetSecurityObjectRequest, ...dcerpc.CallOption) (*SetSecurityObjectResponse, error)

	// The SamrQuerySecurityObject method queries the access control on a server, domain,
	// user, group, or alias object.
	QuerySecurityObject(context.Context, *QuerySecurityObjectRequest, ...dcerpc.CallOption) (*QuerySecurityObjectResponse, error)

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// The SamrLookupDomainInSamServer method obtains the SID of a domain object, given
	// the object's name.
	LookupDomainInSAMServer(context.Context, *LookupDomainInSAMServerRequest, ...dcerpc.CallOption) (*LookupDomainInSAMServerResponse, error)

	// The SamrEnumerateDomainsInSamServer method obtains a listing of all domains hosted
	// by the server side of this protocol.
	EnumerateDomainsInSAMServer(context.Context, *EnumerateDomainsInSAMServerRequest, ...dcerpc.CallOption) (*EnumerateDomainsInSAMServerResponse, error)

	// The SamrOpenDomain method obtains a handle to a domain object, given a SID.
	OpenDomain(context.Context, *OpenDomainRequest, ...dcerpc.CallOption) (*OpenDomainResponse, error)

	// The SamrQueryInformationDomain method obtains attributes from a domain object.
	//
	// See the description of SamrQueryInformationDomain2 (section 3.1.5.5.1) for details,
	// because the method interface arguments and message processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	QueryInformationDomain(context.Context, *QueryInformationDomainRequest, ...dcerpc.CallOption) (*QueryInformationDomainResponse, error)

	// The SamrSetInformationDomain method updates attributes on a domain object.
	SetInformationDomain(context.Context, *SetInformationDomainRequest, ...dcerpc.CallOption) (*SetInformationDomainResponse, error)

	// The SamrCreateGroupInDomain method creates a group object within a domain.
	CreateGroupInDomain(context.Context, *CreateGroupInDomainRequest, ...dcerpc.CallOption) (*CreateGroupInDomainResponse, error)

	// The SamrEnumerateGroupsInDomain method enumerates all groups.
	EnumerateGroupsInDomain(context.Context, *EnumerateGroupsInDomainRequest, ...dcerpc.CallOption) (*EnumerateGroupsInDomainResponse, error)

	// The SamrCreateUserInDomain method creates a user.
	CreateUserInDomain(context.Context, *CreateUserInDomainRequest, ...dcerpc.CallOption) (*CreateUserInDomainResponse, error)

	// The SamrEnumerateUsersInDomain method enumerates all users.
	EnumerateUsersInDomain(context.Context, *EnumerateUsersInDomainRequest, ...dcerpc.CallOption) (*EnumerateUsersInDomainResponse, error)

	// The SamrCreateAliasInDomain method creates an alias.
	CreateAliasInDomain(context.Context, *CreateAliasInDomainRequest, ...dcerpc.CallOption) (*CreateAliasInDomainResponse, error)

	// The SamrEnumerateAliasesInDomain method enumerates all aliases.
	EnumerateAliasesInDomain(context.Context, *EnumerateAliasesInDomainRequest, ...dcerpc.CallOption) (*EnumerateAliasesInDomainResponse, error)

	// The SamrGetAliasMembership method obtains the union of all aliases that a given set
	// of SIDs is a member of.
	GetAliasMembership(context.Context, *GetAliasMembershipRequest, ...dcerpc.CallOption) (*GetAliasMembershipResponse, error)

	// The SamrLookupNamesInDomain method translates a set of account names into a set of
	// RIDs.
	LookupNamesInDomain(context.Context, *LookupNamesInDomainRequest, ...dcerpc.CallOption) (*LookupNamesInDomainResponse, error)

	// The SamrLookupIdsInDomain method translates a set of RIDs into account names.
	LookupIDsInDomain(context.Context, *LookupIDsInDomainRequest, ...dcerpc.CallOption) (*LookupIDsInDomainResponse, error)

	// The SamrOpenGroup method obtains a handle to a group, given a RID.
	OpenGroup(context.Context, *OpenGroupRequest, ...dcerpc.CallOption) (*OpenGroupResponse, error)

	// The SamrQueryInformationGroup method obtains attributes from a group object.
	QueryInformationGroup(context.Context, *QueryInformationGroupRequest, ...dcerpc.CallOption) (*QueryInformationGroupResponse, error)

	// The SamrSetInformationGroup method updates attributes on a group object.
	SetInformationGroup(context.Context, *SetInformationGroupRequest, ...dcerpc.CallOption) (*SetInformationGroupResponse, error)

	// The SamrAddMemberToGroup method adds a member to a group.
	AddMemberToGroup(context.Context, *AddMemberToGroupRequest, ...dcerpc.CallOption) (*AddMemberToGroupResponse, error)

	// The SamrDeleteGroup method removes a group object.
	DeleteGroup(context.Context, *DeleteGroupRequest, ...dcerpc.CallOption) (*DeleteGroupResponse, error)

	// The SamrRemoveMemberFromGroup method removes a member from a group.
	RemoveMemberFromGroup(context.Context, *RemoveMemberFromGroupRequest, ...dcerpc.CallOption) (*RemoveMemberFromGroupResponse, error)

	// The SamrGetMembersInGroup method reads the members of a group.
	GetMembersInGroup(context.Context, *GetMembersInGroupRequest, ...dcerpc.CallOption) (*GetMembersInGroupResponse, error)

	// The SamrSetMemberAttributesOfGroup method sets the attributes of a member relationship.
	SetMemberAttributesOfGroup(context.Context, *SetMemberAttributesOfGroupRequest, ...dcerpc.CallOption) (*SetMemberAttributesOfGroupResponse, error)

	// The SamrOpenAlias method obtains a handle to an alias, given a RID.
	OpenAlias(context.Context, *OpenAliasRequest, ...dcerpc.CallOption) (*OpenAliasResponse, error)

	// The SamrQueryInformationAlias method obtains attributes from an alias object.
	QueryInformationAlias(context.Context, *QueryInformationAliasRequest, ...dcerpc.CallOption) (*QueryInformationAliasResponse, error)

	// The SamrSetInformationAlias method updates attributes on an alias object.
	SetInformationAlias(context.Context, *SetInformationAliasRequest, ...dcerpc.CallOption) (*SetInformationAliasResponse, error)

	// The SamrDeleteAlias method removes an alias object.
	DeleteAlias(context.Context, *DeleteAliasRequest, ...dcerpc.CallOption) (*DeleteAliasResponse, error)

	// The SamrAddMemberToAlias method adds a member to an alias.
	AddMemberToAlias(context.Context, *AddMemberToAliasRequest, ...dcerpc.CallOption) (*AddMemberToAliasResponse, error)

	// The SamrRemoveMemberFromAlias method removes a member from an alias.
	RemoveMemberFromAlias(context.Context, *RemoveMemberFromAliasRequest, ...dcerpc.CallOption) (*RemoveMemberFromAliasResponse, error)

	// The SamrGetMembersInAlias method obtains the membership list of an alias.
	GetMembersInAlias(context.Context, *GetMembersInAliasRequest, ...dcerpc.CallOption) (*GetMembersInAliasResponse, error)

	// The SamrOpenUser method obtains a handle to a user, given a RID.
	OpenUser(context.Context, *OpenUserRequest, ...dcerpc.CallOption) (*OpenUserResponse, error)

	// The SamrDeleteUser method removes a user object.
	DeleteUser(context.Context, *DeleteUserRequest, ...dcerpc.CallOption) (*DeleteUserResponse, error)

	// The SamrQueryInformationUser method obtains attributes from a user object.
	//
	// See the description of SamrQueryInformationUser2 (section 3.1.5.5.5) for details,
	// because the method interface arguments and message processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	QueryInformationUser(context.Context, *QueryInformationUserRequest, ...dcerpc.CallOption) (*QueryInformationUserResponse, error)

	// The SamrSetInformationUser method updates attributes on a user object.
	//
	// See the description of SamrSetInformationUser2 (section 3.1.5.6.4) for details, because
	// the method interface arguments and message processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	SetInformationUser(context.Context, *SetInformationUserRequest, ...dcerpc.CallOption) (*SetInformationUserResponse, error)

	// The SamrChangePasswordUser method changes the password of a user object.
	ChangePasswordUser(context.Context, *ChangePasswordUserRequest, ...dcerpc.CallOption) (*ChangePasswordUserResponse, error)

	// The SamrGetGroupsForUser method obtains a listing of groups that a user is a member
	// of.
	GetGroupsForUser(context.Context, *GetGroupsForUserRequest, ...dcerpc.CallOption) (*GetGroupsForUserResponse, error)

	// The SamrQueryDisplayInformation method obtains a list of accounts in ascending name-sorted
	// order, starting at a specified index.
	//
	// See the description of SamrQueryDisplayInformation3 (section 3.1.5.3.1) for details,
	// because the method interface arguments and message processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	QueryDisplayInformation(context.Context, *QueryDisplayInformationRequest, ...dcerpc.CallOption) (*QueryDisplayInformationResponse, error)

	// The SamrGetDisplayEnumerationIndex method obtains an index into an ascending account-name–sorted
	// list of accounts.
	//
	// See the description of SamrGetDisplayEnumerationIndex2 (section 3.1.5.3.4) for details,
	// because the method-interface arguments and processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	GetDisplayEnumerationIndex(context.Context, *GetDisplayEnumerationIndexRequest, ...dcerpc.CallOption) (*GetDisplayEnumerationIndexResponse, error)

	// Opnum42NotUsedOnWire operation.
	// Opnum42NotUsedOnWire

	// Opnum43NotUsedOnWire operation.
	// Opnum43NotUsedOnWire

	// The SamrGetUserDomainPasswordInformation method obtains select password policy information
	// (without requiring a domain handle).
	GetUserDomainPasswordInformation(context.Context, *GetUserDomainPasswordInformationRequest, ...dcerpc.CallOption) (*GetUserDomainPasswordInformationResponse, error)

	// The SamrRemoveMemberFromForeignDomain method removes a member from all aliases.
	RemoveMemberFromForeignDomain(context.Context, *RemoveMemberFromForeignDomainRequest, ...dcerpc.CallOption) (*RemoveMemberFromForeignDomainResponse, error)

	// The SamrQueryInformationDomain2 method obtains attributes from a domain object.
	QueryInformationDomain2(context.Context, *QueryInformationDomain2Request, ...dcerpc.CallOption) (*QueryInformationDomain2Response, error)

	// The SamrQueryInformationUser2 method obtains attributes from a user object.
	QueryInformationUser2(context.Context, *QueryInformationUser2Request, ...dcerpc.CallOption) (*QueryInformationUser2Response, error)

	// The SamrQueryDisplayInformation2 method obtains a list of accounts in ascending name-sorted
	// order, starting at a specified index.
	//
	// See the description of SamrQueryDisplayInformation3 (section 3.1.5.3.1) for details,
	// because the method-interface arguments and message processing are identical.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	QueryDisplayInformation2(context.Context, *QueryDisplayInformation2Request, ...dcerpc.CallOption) (*QueryDisplayInformation2Response, error)

	// The SamrGetDisplayEnumerationIndex2 method obtains an index into an ascending account-name–sorted
	// list of accounts, such that the index is the position in the list of the accounts
	// whose account name best matches a client-provided string.
	GetDisplayEnumerationIndex2(context.Context, *GetDisplayEnumerationIndex2Request, ...dcerpc.CallOption) (*GetDisplayEnumerationIndex2Response, error)

	// The SamrCreateUser2InDomain method creates a user.
	CreateUser2InDomain(context.Context, *CreateUser2InDomainRequest, ...dcerpc.CallOption) (*CreateUser2InDomainResponse, error)

	// The SamrQueryDisplayInformation3 method obtains a listing of accounts in ascending
	// name-sorted order, starting at a specified index.
	QueryDisplayInformation3(context.Context, *QueryDisplayInformation3Request, ...dcerpc.CallOption) (*QueryDisplayInformation3Response, error)

	// The SamrAddMultipleMembersToAlias method adds multiple members to an alias.
	AddMultipleMembersToAlias(context.Context, *AddMultipleMembersToAliasRequest, ...dcerpc.CallOption) (*AddMultipleMembersToAliasResponse, error)

	// The SamrRemoveMultipleMembersFromAlias method removes multiple members from an alias.
	RemoveMultipleMembersFromAlias(context.Context, *RemoveMultipleMembersFromAliasRequest, ...dcerpc.CallOption) (*RemoveMultipleMembersFromAliasResponse, error)

	// The SamrOemChangePasswordUser2 method changes a user's password.
	//
	// );
	OEMChangePasswordUser2(context.Context, *OEMChangePasswordUser2Request, ...dcerpc.CallOption) (*OEMChangePasswordUser2Response, error)

	// The SamrUnicodeChangePasswordUser2 method changes a user account's password.
	UnicodeChangePasswordUser2(context.Context, *UnicodeChangePasswordUser2Request, ...dcerpc.CallOption) (*UnicodeChangePasswordUser2Response, error)

	// The SamrGetDomainPasswordInformation method obtains select password policy information
	// (without authenticating to the server).
	GetDomainPasswordInformation(context.Context, *GetDomainPasswordInformationRequest, ...dcerpc.CallOption) (*GetDomainPasswordInformationResponse, error)

	// The SamrConnect2 method returns a handle to a server object.
	Connect2(context.Context, *Connect2Request, ...dcerpc.CallOption) (*Connect2Response, error)

	// The SamrSetInformationUser2 method updates attributes on a user object.
	SetInformationUser2(context.Context, *SetInformationUser2Request, ...dcerpc.CallOption) (*SetInformationUser2Response, error)

	// Opnum59NotUsedOnWire operation.
	// Opnum59NotUsedOnWire

	// Opnum60NotUsedOnWire operation.
	// Opnum60NotUsedOnWire

	// Opnum61NotUsedOnWire operation.
	// Opnum61NotUsedOnWire

	// The SamrConnect4 method obtains a handle to a server object.
	Connect4(context.Context, *Connect4Request, ...dcerpc.CallOption) (*Connect4Response, error)

	// Opnum63NotUsedOnWire operation.
	// Opnum63NotUsedOnWire

	// The SamrConnect5 method obtains a handle to a server object.
	Connect5(context.Context, *Connect5Request, ...dcerpc.CallOption) (*Connect5Response, error)

	// The SamrRidToSid method obtains the SID of an account, given a RID.
	RIDToSID(context.Context, *RIDToSIDRequest, ...dcerpc.CallOption) (*RIDToSIDResponse, error)

	// The SamrSetDSRMPassword method sets a local recovery password.
	SetDSRMPassword(context.Context, *SetDSRMPasswordRequest, ...dcerpc.CallOption) (*SetDSRMPasswordResponse, error)

	// The SamrValidatePassword method validates an application password against the locally
	// stored policy.
	ValidatePassword(context.Context, *ValidatePasswordRequest, ...dcerpc.CallOption) (*ValidatePasswordResponse, error)

	// Opnum68NotUsedOnWire operation.
	// Opnum68NotUsedOnWire

	// Opnum69NotUsedOnWire operation.
	// Opnum69NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// PackagesCredentialsName represents the PACKAGES_CREDENTIALS_NAME RPC constant
var PackagesCredentialsName = "Packages"

// WdigestCredentialsName represents the WDIGEST_CREDENTIALS_NAME RPC constant
var WdigestCredentialsName = "Primary:WDigest"

// KerberosStoredCredentialName represents the KERB_STORED_CREDENTIAL_NAME RPC constant
var KerberosStoredCredentialName = "Primary:Kerberos"

// KerberosStoredCredentialNewName represents the KERB_STORED_CREDENTIAL_NEW_NAME RPC constant
var KerberosStoredCredentialNewName = "Primary:Kerberos-Newer-Keys"

// NTLMStrongNTOWFName represents the NTLM_STRONG_NTOWF_NAME RPC constant
var NTLMStrongNTOWFName = "Primary:NTLM-Strong-NTOWF"

// CleartextCredentialsName represents the CLEARTEXT_CREDENTIALS_NAME RPC constant
var CleartextCredentialsName = "Primary:CLEARTEXT"

// RawCredentials structure represents RAW_CREDENTIALS RPC structure.
type RawCredentials struct {
	RawCredentials []byte `idl:"name:RawCredentials" json:"raw_credentials"`
}

func (o *RawCredentials) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RawCredentials) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.RawCredentials != nil {
		_ptr_RawCredentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.RawCredentials {
				i1 := i1
				if err := w.WriteData(o.RawCredentials[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RawCredentials, _ptr_RawCredentials); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RawCredentials) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_RawCredentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.RawCredentials = append(o.RawCredentials, uint8(0))
			if err := w.ReadData(&o.RawCredentials[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_RawCredentials := func(ptr interface{}) { o.RawCredentials = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.RawCredentials, _s_RawCredentials, _ptr_RawCredentials); err != nil {
		return err
	}
	return nil
}

// CleartextCredentials structure represents CLEARTEXT_CREDENTIALS RPC structure.
type CleartextCredentials struct {
	CleartextCredentials string `idl:"name:CleartextCredentials" json:"cleartext_credentials"`
}

func (o *CleartextCredentials) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CleartextCredentials) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.CleartextCredentials != "" {
		_ptr_CleartextCredentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_CleartextCredentials_buf := utf16.Encode([]rune(o.CleartextCredentials))
			for i1 := range _CleartextCredentials_buf {
				i1 := i1
				if err := w.WriteData(_CleartextCredentials_buf[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CleartextCredentials, _ptr_CleartextCredentials); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CleartextCredentials) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_CleartextCredentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		var _CleartextCredentials_buf []uint16
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			_CleartextCredentials_buf = append(_CleartextCredentials_buf, uint16(0))
			if err := w.ReadData(&_CleartextCredentials_buf[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_CleartextCredentials := func(ptr interface{}) { o.CleartextCredentials = *ptr.(*string) }
	if err := w.ReadPointer(&o.CleartextCredentials, _s_CleartextCredentials, _ptr_CleartextCredentials); err != nil {
		return err
	}
	return nil
}

// NTLMStrongNTOWF structure represents NTLM_STRONG_NTOWF RPC structure.
type NTLMStrongNTOWF struct {
	NTLMStrongNTOWF []byte `idl:"name:NTLMStrongNTOWF" json:"ntlm_strong_ntowf"`
}

func (o *NTLMStrongNTOWF) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTLMStrongNTOWF) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.NTLMStrongNTOWF != nil {
		_ptr_NTLMStrongNTOWF := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.NTLMStrongNTOWF {
				i1 := i1
				if uint64(i1) >= 16 {
					break
				}
				if err := w.WriteData(o.NTLMStrongNTOWF[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.NTLMStrongNTOWF); uint64(i1) < 16; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NTLMStrongNTOWF, _ptr_NTLMStrongNTOWF); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTLMStrongNTOWF) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_NTLMStrongNTOWF := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		o.NTLMStrongNTOWF = make([]byte, 16)
		for i1 := range o.NTLMStrongNTOWF {
			i1 := i1
			if err := w.ReadData(&o.NTLMStrongNTOWF[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_NTLMStrongNTOWF := func(ptr interface{}) { o.NTLMStrongNTOWF = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.NTLMStrongNTOWF, _s_NTLMStrongNTOWF, _ptr_NTLMStrongNTOWF); err != nil {
		return err
	}
	return nil
}

// PackagesCredentials structure represents PACKAGES_CREDENTIALS RPC structure.
type PackagesCredentials struct {
	Packages []string `idl:"name:Packages" json:"packages"`
}

func (o *PackagesCredentials) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PackagesCredentials) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Packages != nil {
		_ptr_Packages := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			var _Packages_buf []uint16
			for i1 := range o.Packages {
				_Packages_buf = append(_Packages_buf, utf16.Encode([]rune(o.Packages[i1]))...)
				_Packages_buf = append(_Packages_buf, uint16(0))
			}
			_Packages_buf = append(_Packages_buf, uint16(0))
			for i1 := range _Packages_buf {
				i1 := i1
				if err := w.WriteData(_Packages_buf[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Packages, _ptr_Packages); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PackagesCredentials) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Packages := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		var _Packages_buf []uint16
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			_Packages_buf = append(_Packages_buf, uint16(0))
			if err := w.ReadData(&_Packages_buf[i1]); err != nil {
				return err
			}
		}
		_tmp_Packages_buf := string(utf16.Decode(_Packages_buf))
		if _tmp_Packages_buf = strings.TrimRight(_tmp_Packages_buf, ndr.ZeroString); _tmp_Packages_buf != "" {
			o.Packages = strings.Split(_tmp_Packages_buf, ndr.ZeroString)
		}
		return nil
	})
	_s_Packages := func(ptr interface{}) { o.Packages = *ptr.(*[]string) }
	if err := w.ReadPointer(&o.Packages, _s_Packages, _ptr_Packages); err != nil {
		return err
	}
	return nil
}

// WdigestCredentials structure represents WDIGEST_CREDENTIALS RPC structure.
//
// The WDIGEST_CREDENTIALS structure defines the format of the Primary:WDigest property
// within the supplementalCredentials attribute. This structure is stored as a property
// value in a USER_PROPERTY structure.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved1                     | Reserved2                     | Version                       | NumberOfHashes                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved3                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash1 (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash2 (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash3 (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash4 (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash5 (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash6 (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash7 (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash8 (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash9 (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash10 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash11 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash12 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash13 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash14 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash15 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash16 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash17 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash18 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash19 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash20 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash21 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash22 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash23 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash24 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash25 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash26 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash27 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash28 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Hash29 (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type WdigestCredentials struct {
	// Reserved1 (1 byte): This value MUST be ignored by the recipient and MAY<25> be set
	// to arbitrary values upon an update to the supplementalCredentials attribute.
	_ uint8 `idl:"name:Reserved1"`
	// Reserved2 (1 byte): This value MUST be ignored by the recipient and MUST be set to
	// zero.
	_ uint8 `idl:"name:Reserved2"`
	// Version (1 byte): This value MUST be set to 1.
	Version uint8 `idl:"name:Version" json:"version"`
	// NumberOfHashes (1 byte): This value MUST be set to 29 because there are 29 hashes
	// in the array.
	NumberOfHashes uint8 `idl:"name:NumberOfHashes" json:"number_of_hashes"`
	// Reserved3 (12 bytes): This value MUST be ignored by the recipient and MUST be set
	// to zero.
	//
	// For information on the Hash fields, see section 3.1.1.8.11.
	_      []byte   `idl:"name:Reserved3"`
	Hashes [][]byte `idl:"name:Hashes" json:"hashes"`
}

func (o *WdigestCredentials) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WdigestCredentials) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	// reserved Reserved1
	if err := w.WriteData(uint8(0)); err != nil {
		return err
	}
	// reserved Reserved2
	if err := w.WriteData(uint8(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfHashes); err != nil {
		return err
	}
	// reserved Reserved3
	for i1 := 0; uint64(i1) < 12; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Hashes {
		i1 := i1
		if uint64(i1) >= 29 {
			break
		}
		for i2 := range o.Hashes[i1] {
			i2 := i2
			if uint64(i2) >= 16 {
				break
			}
			if err := w.WriteData(o.Hashes[i1][i2]); err != nil {
				return err
			}
		}
		for i2 := len(o.Hashes[i1]); uint64(i2) < 16; i2++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Hashes); uint64(i1) < 29; i1++ {
		for i2 := 0; uint64(i2) < 16; i2++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}
func (o *WdigestCredentials) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	// reserved Reserved1
	var _Reserved1 uint8
	if err := w.ReadData(&_Reserved1); err != nil {
		return err
	}
	// reserved Reserved2
	var _Reserved2 uint8
	if err := w.ReadData(&_Reserved2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfHashes); err != nil {
		return err
	}
	// reserved Reserved3
	var _Reserved3 []byte
	_Reserved3 = make([]byte, 12)
	for i1 := range _Reserved3 {
		i1 := i1
		if err := w.ReadData(&_Reserved3[i1]); err != nil {
			return err
		}
	}
	o.Hashes = make([][]byte, 29)
	for i1 := range o.Hashes {
		i1 := i1
		o.Hashes[i1] = make([]byte, 16)
		for i2 := range o.Hashes[i1] {
			i2 := i2
			if err := w.ReadData(&o.Hashes[i1][i2]); err != nil {
				return err
			}
		}
	}
	return nil
}

// KerberosKeyData structure represents KERB_KEY_DATA RPC structure.
//
// The KERB_KEY_DATA structure holds a cryptographic key. This structure is used in
// conjunction with KERB_STORED_CREDENTIAL. For more information, see section 3.1.1.8.11.4.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved1                                                     | Reserved2                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved3                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| KeyType                                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| KeyLength                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| KeyOffset                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type KerberosKeyData struct {
	// Reserved1 (2 bytes): This value MUST be ignored by the recipient and MUST be set
	// to zero.
	_ uint16 `idl:"name:Reserved1"`
	// Reserved2 (2 bytes): This value MUST be ignored by the recipient and MUST be set
	// to zero.
	_ uint16 `idl:"name:Reserved2"`
	// Reserved3 (4 bytes): This value MUST be ignored by the recipient and MUST be set
	// to zero.
	_ uint32 `idl:"name:Reserved3"`
	// KeyType (4 bytes): Indicates the type of key, stored as a 32-bit unsigned integer
	// in little-endian byte order. This MUST be set to one of the following values, which
	// are defined in section 2.2.10.8.
	//
	//	+-------+-------------+
	//	|       |             |
	//	| VALUE |   MEANING   |
	//	|       |             |
	//	+-------+-------------+
	//	+-------+-------------+
	//	|     1 | dec-cbc-crc |
	//	+-------+-------------+
	//	|     3 | des-cbc-md5 |
	//	+-------+-------------+
	KeyType uint32 `idl:"name:KeyType" json:"key_type"`
	// KeyLength (4 bytes): The length, in bytes, of the value beginning at KeyOffset. The
	// value of this field is stored in little-endian byte order.
	KeyLength uint32 `idl:"name:KeyLength" json:"key_length"`
	// KeyOffset (4 bytes): An offset, in little-endian byte order, from the beginning of
	// the property value (that is, from the beginning of the Revision field of KERB_STORED_CREDENTIAL)
	// to where the key value starts. The key value is the hash value specified according
	// to the KeyType.
	KeyOffset uint32 `idl:"name:KeyOffset" json:"key_offset"`
	KeyData   []byte `idl:"name:KeyData" json:"key_data"`
}

func (o *KerberosKeyData) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosKeyData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	// reserved Reserved1
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	// reserved Reserved2
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	// reserved Reserved3
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyType); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyLength); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyOffset); err != nil {
		return err
	}
	return nil
}
func (o *KerberosKeyData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	// reserved Reserved1
	var _Reserved1 uint16
	if err := w.ReadData(&_Reserved1); err != nil {
		return err
	}
	// reserved Reserved2
	var _Reserved2 uint16
	if err := w.ReadData(&_Reserved2); err != nil {
		return err
	}
	// reserved Reserved3
	var _Reserved3 uint32
	if err := w.ReadData(&_Reserved3); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyType); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyOffset); err != nil {
		return err
	}
	return nil
}

// KerberosStoredCredential structure represents KERB_STORED_CREDENTIAL RPC structure.
//
// The KERB_STORED_CREDENTIAL structure is a variable-length structure that defines
// the format of the Primary:Kerberos property within the supplementalCredentials attribute.
// For information on how this structure is created, see section 3.1.1.8.11.4.
//
// This structure is stored as a property value in a USER_PROPERTY structure.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Revision                                                      | Flags                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| CredentialCount                                               | OldCredentialCount                                            |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| DefaultSaltLength                                             | DefaultSaltMaximumLength                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| DefaultSaltOffset                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Credentials (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| OldCredentials (variable)                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| DefaultSalt (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| KeyValues (variable)                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type KerberosStoredCredential struct {
	// Revision (2 bytes): This value MUST be set to 3.
	Revision uint16 `idl:"name:Revision" json:"revision"`
	// Flags (2 bytes): This value MUST be zero and ignored on read.
	Flags uint16 `idl:"name:Flags" json:"flags"`
	// CredentialCount (2 bytes): This is the count of elements in the Credentials array.
	// This value MUST be set to 2.
	CredentialCount uint16 `idl:"name:CredentialCount" json:"credential_count"`
	// OldCredentialCount (2 bytes): This is the count of elements in the OldCredentials
	// array that contain the keys for the previous password. This value MUST be set to
	// 0 or 2.
	OldCredentialCount uint16 `idl:"name:OldCredentialCount" json:"old_credential_count"`
	// DefaultSaltLength (2 bytes): The length, in bytes, of a salt value.
	DefaultSaltLength uint16 `idl:"name:DefaultSaltLength" json:"default_salt_length"`
	// DefaultSaltMaximumLength (2 bytes): The length, in bytes, of the buffer containing
	// the salt value.
	DefaultSaltMaximumLength uint16 `idl:"name:DefaultSaltMaximumLength" json:"default_salt_maximum_length"`
	// DefaultSaltOffset (4 bytes): An offset, in little-endian byte order, from the beginning
	// of the attribute value (that is, from the beginning of the Revision field of KERB_STORED_CREDENTIAL)
	// to where the salt value starts. This value SHOULD be ignored on read.
	DefaultSaltOffset uint32 `idl:"name:DefaultSaltOffset" json:"default_salt_offset"`
	// Credentials (variable): An array of CredentialCount KERB_KEY_DATA (section 2.2.10.5)
	// elements.
	Credentials []*KerberosKeyData `idl:"name:Credentials;size_is:(CredentialCount)" json:"credentials"`
	// OldCredentials (variable): An array of OldCredentialCount KERB_KEY_DATA elements.
	OldCredentials []*KerberosKeyData `idl:"name:OldCredentials;size_is:(OldCredentialCount)" json:"old_credentials"`
	// DefaultSalt (variable): The default salt value.
	DefaultSalt []byte `idl:"name:DefaultSalt;size_is:(DefaultSaltMaximumLength)" json:"default_salt"`
	// KeyValues (variable): An array of CredentialCount + OldCredentialCount key values.
	// Each key value MUST be located at the offset specified by the corresponding KeyOffset
	// values specified in Credentials and OldCredentials.
	KeyValues []byte `idl:"name:KeyValues" json:"key_values"`
}

func (o *KerberosStoredCredential) xxx_PreparePayload(ctx context.Context) error {
	if o.Credentials != nil && o.CredentialCount == 0 {
		o.CredentialCount = uint16(len(o.Credentials))
	}
	if o.OldCredentials != nil && o.OldCredentialCount == 0 {
		o.OldCredentialCount = uint16(len(o.OldCredentials))
	}
	if o.DefaultSalt != nil && o.DefaultSaltMaximumLength == 0 {
		o.DefaultSaltMaximumLength = uint16(len(o.DefaultSalt))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosStoredCredential) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Revision); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.CredentialCount); err != nil {
		return err
	}
	if err := w.WriteData(o.OldCredentialCount); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultSaltLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultSaltMaximumLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultSaltOffset); err != nil {
		return err
	}
	if o.Credentials != nil || o.CredentialCount > 0 {
		_ptr_Credentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.CredentialCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Credentials {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Credentials[i1] != nil {
					_ptr_Credentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Credentials[i1] != nil {
							if err := o.Credentials[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&KerberosKeyData{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Credentials[i1], _ptr_Credentials); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Credentials); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Credentials, _ptr_Credentials); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.OldCredentials != nil || o.OldCredentialCount > 0 {
		_ptr_OldCredentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.OldCredentialCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.OldCredentials {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.OldCredentials[i1] != nil {
					_ptr_OldCredentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.OldCredentials[i1] != nil {
							if err := o.OldCredentials[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&KerberosKeyData{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.OldCredentials[i1], _ptr_OldCredentials); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.OldCredentials); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OldCredentials, _ptr_OldCredentials); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DefaultSalt != nil || o.DefaultSaltMaximumLength > 0 {
		_ptr_DefaultSalt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DefaultSaltMaximumLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.DefaultSalt {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.DefaultSalt[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.DefaultSalt); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DefaultSalt, _ptr_DefaultSalt); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.KeyValues != nil {
		_ptr_KeyValues := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.KeyValues {
				i1 := i1
				if err := w.WriteData(o.KeyValues[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.KeyValues, _ptr_KeyValues); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosStoredCredential) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Revision); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.CredentialCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.OldCredentialCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultSaltLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultSaltMaximumLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultSaltOffset); err != nil {
		return err
	}
	_ptr_Credentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.CredentialCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.CredentialCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Credentials", sizeInfo[0])
		}
		o.Credentials = make([]*KerberosKeyData, sizeInfo[0])
		for i1 := range o.Credentials {
			i1 := i1
			_ptr_Credentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Credentials[i1] == nil {
					o.Credentials[i1] = &KerberosKeyData{}
				}
				if err := o.Credentials[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Credentials := func(ptr interface{}) { o.Credentials[i1] = *ptr.(**KerberosKeyData) }
			if err := w.ReadPointer(&o.Credentials[i1], _s_Credentials, _ptr_Credentials); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Credentials := func(ptr interface{}) { o.Credentials = *ptr.(*[]*KerberosKeyData) }
	if err := w.ReadPointer(&o.Credentials, _s_Credentials, _ptr_Credentials); err != nil {
		return err
	}
	_ptr_OldCredentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.OldCredentialCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.OldCredentialCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OldCredentials", sizeInfo[0])
		}
		o.OldCredentials = make([]*KerberosKeyData, sizeInfo[0])
		for i1 := range o.OldCredentials {
			i1 := i1
			_ptr_OldCredentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.OldCredentials[i1] == nil {
					o.OldCredentials[i1] = &KerberosKeyData{}
				}
				if err := o.OldCredentials[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_OldCredentials := func(ptr interface{}) { o.OldCredentials[i1] = *ptr.(**KerberosKeyData) }
			if err := w.ReadPointer(&o.OldCredentials[i1], _s_OldCredentials, _ptr_OldCredentials); err != nil {
				return err
			}
		}
		return nil
	})
	_s_OldCredentials := func(ptr interface{}) { o.OldCredentials = *ptr.(*[]*KerberosKeyData) }
	if err := w.ReadPointer(&o.OldCredentials, _s_OldCredentials, _ptr_OldCredentials); err != nil {
		return err
	}
	_ptr_DefaultSalt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DefaultSaltMaximumLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DefaultSaltMaximumLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DefaultSalt", sizeInfo[0])
		}
		o.DefaultSalt = make([]byte, sizeInfo[0])
		for i1 := range o.DefaultSalt {
			i1 := i1
			if err := w.ReadData(&o.DefaultSalt[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_DefaultSalt := func(ptr interface{}) { o.DefaultSalt = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.DefaultSalt, _s_DefaultSalt, _ptr_DefaultSalt); err != nil {
		return err
	}
	_ptr_KeyValues := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.KeyValues = append(o.KeyValues, uint8(0))
			if err := w.ReadData(&o.KeyValues[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_KeyValues := func(ptr interface{}) { o.KeyValues = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.KeyValues, _s_KeyValues, _ptr_KeyValues); err != nil {
		return err
	}
	return nil
}

// KerberosKeyDataNew structure represents KERB_KEY_DATA_NEW RPC structure.
//
// The KERB_KEY_DATA_NEW structure holds a cryptographic key. This structure is used
// in conjunction with KERB_STORED_CREDENTIAL_NEW. For more information, see section
// 3.1.1.8.11.6.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved1                                                     | Reserved2                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved3                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| IterationCount                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| KeyType                                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| KeyLength                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| KeyOffset                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type KerberosKeyDataNew struct {
	// Reserved1 (2 bytes): This value MUST be ignored by the recipient and MUST be set
	// to zero.
	_ uint16 `idl:"name:Reserved1"`
	// Reserved2 (2 bytes): This value MUST be ignored by the recipient and MUST be set
	// to zero.
	_ uint16 `idl:"name:Reserved2"`
	// Reserved3 (4 bytes): This value MUST be ignored by the recipient and MUST be set
	// to zero.
	_ uint32 `idl:"name:Reserved3"`
	// IterationCount (4 bytes):  Indicates the iteration count used to calculate the password
	// hashes.
	IterationCount uint32 `idl:"name:IterationCount" json:"iteration_count"`
	// KeyType (4 bytes): Indicates the type of key, stored as a 32-bit unsigned integer
	// in little-endian byte order. This MUST be one of the values listed in section 2.2.10.8.
	KeyType uint32 `idl:"name:KeyType" json:"key_type"`
	// KeyLength (4 bytes): The length, in bytes, of the value beginning at KeyOffset. The
	// value of this field is stored in little-endian byte order.
	KeyLength uint32 `idl:"name:KeyLength" json:"key_length"`
	// KeyOffset (4 bytes): An offset, in little-endian byte order, from the beginning of
	// the property value (that is, from the beginning of the Revision field of KERB_STORED_CREDENTIAL_NEW)
	// to where the key value starts.
	KeyOffset uint32 `idl:"name:KeyOffset" json:"key_offset"`
	KeyData   []byte `idl:"name:KeyData" json:"key_data"`
}

func (o *KerberosKeyDataNew) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosKeyDataNew) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	// reserved Reserved1
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	// reserved Reserved2
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	// reserved Reserved3
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.IterationCount); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyType); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyLength); err != nil {
		return err
	}
	if err := w.WriteData(o.KeyOffset); err != nil {
		return err
	}
	return nil
}
func (o *KerberosKeyDataNew) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	// reserved Reserved1
	var _Reserved1 uint16
	if err := w.ReadData(&_Reserved1); err != nil {
		return err
	}
	// reserved Reserved2
	var _Reserved2 uint16
	if err := w.ReadData(&_Reserved2); err != nil {
		return err
	}
	// reserved Reserved3
	var _Reserved3 uint32
	if err := w.ReadData(&_Reserved3); err != nil {
		return err
	}
	if err := w.ReadData(&o.IterationCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyType); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyOffset); err != nil {
		return err
	}
	return nil
}

// KerberosStoredCredentialNew structure represents KERB_STORED_CREDENTIAL_NEW RPC structure.
//
// The KERB_STORED_CREDENTIAL_NEW structure is a variable-length structure that defines
// the format of the Primary:Kerberos-Newer-Keys property within the supplementalCredentials
// attribute. For information on how this structure is created, see section 3.1.1.8.11.6.
//
// This structure is stored as a property value in a USER_PROPERTY structure.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Revision                                                      | Flags                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| CredentialCount                                               | ServiceCredentialCount                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| OldCredentialCount                                            | OlderCredentialCount                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| DefaultSaltLength                                             | DefaultSaltMaximumLength                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| DefaultSaltOffset                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| DefaultIterationCount                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Credentials (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ServiceCredentials (variable)                                                                                                 |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| OldCredentials (variable)                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| OlderCredentials (variable)                                                                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| DefaultSalt (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| KeyValues (variable)                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type KerberosStoredCredentialNew struct {
	// Revision (2 bytes): This value MUST be set to 4.
	Revision uint16 `idl:"name:Revision" json:"revision"`
	// Flags (2 bytes): This value MUST be zero and ignored on read.
	Flags uint16 `idl:"name:Flags" json:"flags"`
	// CredentialCount (2 bytes): This is the count of elements in the Credentials field.
	CredentialCount uint16 `idl:"name:CredentialCount" json:"credential_count"`
	// ServiceCredentialCount (2 bytes):  This is the count of elements in the ServiceCredentials
	// field. It MUST be zero.
	ServiceCredentialCount uint16 `idl:"name:ServiceCredentialCount" json:"service_credential_count"`
	// OldCredentialCount (2 bytes): This is the count of elements in the OldCredentials
	// field that contain the keys for the previous password.
	OldCredentialCount uint16 `idl:"name:OldCredentialCount" json:"old_credential_count"`
	// OlderCredentialCount (2 bytes):  This is the count of elements in the OlderCredentials
	// field that contain the keys for the previous password.
	OlderCredentialCount uint16 `idl:"name:OlderCredentialCount" json:"older_credential_count"`
	// DefaultSaltLength (2 bytes): The length, in bytes, of a salt value.
	DefaultSaltLength uint16 `idl:"name:DefaultSaltLength" json:"default_salt_length"`
	// DefaultSaltMaximumLength (2 bytes): The length, in bytes, of the buffer containing
	// the salt value.
	DefaultSaltMaximumLength uint16 `idl:"name:DefaultSaltMaximumLength" json:"default_salt_maximum_length"`
	// DefaultSaltOffset (4 bytes): An offset, in little-endian byte order, from the beginning
	// of the attribute value (that is, from the beginning of the Revision field of KERB_STORED_CREDENTIAL)
	// to where DefaultSalt starts. This value SHOULD be ignored on read.
	DefaultSaltOffset uint32 `idl:"name:DefaultSaltOffset" json:"default_salt_offset"`
	// DefaultIterationCount (4 bytes): The default iteration count used to calculate the
	// password hashes.
	DefaultIterationCount uint32 `idl:"name:DefaultIterationCount" json:"default_iteration_count"`
	// Credentials (variable): An array of CredentialCount KERB_KEY_DATA_NEW (section 2.2.10.7)
	// elements.
	Credentials []*KerberosKeyDataNew `idl:"name:Credentials;size_is:(CredentialCount)" json:"credentials"`
	// ServiceCredentials (variable): (This field is optional.) An array of ServiceCredentialCount
	// KERB_KEY_DATA_NEW elements.
	ServiceCredentials []*KerberosKeyDataNew `idl:"name:ServiceCredentials;size_is:(ServiceCredentialCount)" json:"service_credentials"`
	// OldCredentials (variable): (This field is optional.) An array of OldCredentialCount
	// KERB_KEY_DATA_NEW elements.
	OldCredentials []*KerberosKeyDataNew `idl:"name:OldCredentials;size_is:(OldCredentialCount)" json:"old_credentials"`
	// OlderCredentials (variable): (This field is optional.) An array of OlderCredentialCount
	// KERB_KEY_DATA_NEW elements.
	OlderCredentials []*KerberosKeyDataNew `idl:"name:OlderCredentials;size_is:(OlderCredentialCount)" json:"older_credentials"`
	// DefaultSalt (variable): The default salt value.
	DefaultSalt []byte `idl:"name:DefaultSalt;size_is:(DefaultSaltMaximumLength)" json:"default_salt"`
	// KeyValues (variable): An array of CredentialCount + ServiceCredentialCount + OldCredentialCount
	// + OlderCredentialCount key values. Each key value MUST be located at the offset specified
	// by the corresponding KeyOffset values specified in Credentials, ServiceCredentials,
	// OldCredentials, and OlderCredentials.
	KeyValues []byte `idl:"name:KeyValues" json:"key_values"`
}

func (o *KerberosStoredCredentialNew) xxx_PreparePayload(ctx context.Context) error {
	if o.Credentials != nil && o.CredentialCount == 0 {
		o.CredentialCount = uint16(len(o.Credentials))
	}
	if o.ServiceCredentials != nil && o.ServiceCredentialCount == 0 {
		o.ServiceCredentialCount = uint16(len(o.ServiceCredentials))
	}
	if o.OldCredentials != nil && o.OldCredentialCount == 0 {
		o.OldCredentialCount = uint16(len(o.OldCredentials))
	}
	if o.OlderCredentials != nil && o.OlderCredentialCount == 0 {
		o.OlderCredentialCount = uint16(len(o.OlderCredentials))
	}
	if o.DefaultSalt != nil && o.DefaultSaltMaximumLength == 0 {
		o.DefaultSaltMaximumLength = uint16(len(o.DefaultSalt))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosStoredCredentialNew) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Revision); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.CredentialCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ServiceCredentialCount); err != nil {
		return err
	}
	if err := w.WriteData(o.OldCredentialCount); err != nil {
		return err
	}
	if err := w.WriteData(o.OlderCredentialCount); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultSaltLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultSaltMaximumLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultSaltOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultIterationCount); err != nil {
		return err
	}
	if o.Credentials != nil || o.CredentialCount > 0 {
		_ptr_Credentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.CredentialCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Credentials {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Credentials[i1] != nil {
					_ptr_Credentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Credentials[i1] != nil {
							if err := o.Credentials[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&KerberosKeyDataNew{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Credentials[i1], _ptr_Credentials); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Credentials); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Credentials, _ptr_Credentials); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ServiceCredentials != nil || o.ServiceCredentialCount > 0 {
		_ptr_ServiceCredentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ServiceCredentialCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ServiceCredentials {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ServiceCredentials[i1] != nil {
					_ptr_ServiceCredentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.ServiceCredentials[i1] != nil {
							if err := o.ServiceCredentials[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&KerberosKeyDataNew{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.ServiceCredentials[i1], _ptr_ServiceCredentials); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ServiceCredentials); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ServiceCredentials, _ptr_ServiceCredentials); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.OldCredentials != nil || o.OldCredentialCount > 0 {
		_ptr_OldCredentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.OldCredentialCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.OldCredentials {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.OldCredentials[i1] != nil {
					_ptr_OldCredentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.OldCredentials[i1] != nil {
							if err := o.OldCredentials[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&KerberosKeyDataNew{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.OldCredentials[i1], _ptr_OldCredentials); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.OldCredentials); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OldCredentials, _ptr_OldCredentials); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.OlderCredentials != nil || o.OlderCredentialCount > 0 {
		_ptr_OlderCredentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.OlderCredentialCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.OlderCredentials {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.OlderCredentials[i1] != nil {
					_ptr_OlderCredentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.OlderCredentials[i1] != nil {
							if err := o.OlderCredentials[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&KerberosKeyDataNew{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.OlderCredentials[i1], _ptr_OlderCredentials); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.OlderCredentials); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OlderCredentials, _ptr_OlderCredentials); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DefaultSalt != nil || o.DefaultSaltMaximumLength > 0 {
		_ptr_DefaultSalt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DefaultSaltMaximumLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.DefaultSalt {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.DefaultSalt[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.DefaultSalt); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DefaultSalt, _ptr_DefaultSalt); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.KeyValues != nil {
		_ptr_KeyValues := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.KeyValues {
				i1 := i1
				if err := w.WriteData(o.KeyValues[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.KeyValues, _ptr_KeyValues); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosStoredCredentialNew) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Revision); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.CredentialCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ServiceCredentialCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.OldCredentialCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.OlderCredentialCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultSaltLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultSaltMaximumLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultSaltOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultIterationCount); err != nil {
		return err
	}
	_ptr_Credentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.CredentialCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.CredentialCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Credentials", sizeInfo[0])
		}
		o.Credentials = make([]*KerberosKeyDataNew, sizeInfo[0])
		for i1 := range o.Credentials {
			i1 := i1
			_ptr_Credentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Credentials[i1] == nil {
					o.Credentials[i1] = &KerberosKeyDataNew{}
				}
				if err := o.Credentials[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Credentials := func(ptr interface{}) { o.Credentials[i1] = *ptr.(**KerberosKeyDataNew) }
			if err := w.ReadPointer(&o.Credentials[i1], _s_Credentials, _ptr_Credentials); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Credentials := func(ptr interface{}) { o.Credentials = *ptr.(*[]*KerberosKeyDataNew) }
	if err := w.ReadPointer(&o.Credentials, _s_Credentials, _ptr_Credentials); err != nil {
		return err
	}
	_ptr_ServiceCredentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ServiceCredentialCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ServiceCredentialCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ServiceCredentials", sizeInfo[0])
		}
		o.ServiceCredentials = make([]*KerberosKeyDataNew, sizeInfo[0])
		for i1 := range o.ServiceCredentials {
			i1 := i1
			_ptr_ServiceCredentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.ServiceCredentials[i1] == nil {
					o.ServiceCredentials[i1] = &KerberosKeyDataNew{}
				}
				if err := o.ServiceCredentials[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ServiceCredentials := func(ptr interface{}) { o.ServiceCredentials[i1] = *ptr.(**KerberosKeyDataNew) }
			if err := w.ReadPointer(&o.ServiceCredentials[i1], _s_ServiceCredentials, _ptr_ServiceCredentials); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ServiceCredentials := func(ptr interface{}) { o.ServiceCredentials = *ptr.(*[]*KerberosKeyDataNew) }
	if err := w.ReadPointer(&o.ServiceCredentials, _s_ServiceCredentials, _ptr_ServiceCredentials); err != nil {
		return err
	}
	_ptr_OldCredentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.OldCredentialCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.OldCredentialCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OldCredentials", sizeInfo[0])
		}
		o.OldCredentials = make([]*KerberosKeyDataNew, sizeInfo[0])
		for i1 := range o.OldCredentials {
			i1 := i1
			_ptr_OldCredentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.OldCredentials[i1] == nil {
					o.OldCredentials[i1] = &KerberosKeyDataNew{}
				}
				if err := o.OldCredentials[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_OldCredentials := func(ptr interface{}) { o.OldCredentials[i1] = *ptr.(**KerberosKeyDataNew) }
			if err := w.ReadPointer(&o.OldCredentials[i1], _s_OldCredentials, _ptr_OldCredentials); err != nil {
				return err
			}
		}
		return nil
	})
	_s_OldCredentials := func(ptr interface{}) { o.OldCredentials = *ptr.(*[]*KerberosKeyDataNew) }
	if err := w.ReadPointer(&o.OldCredentials, _s_OldCredentials, _ptr_OldCredentials); err != nil {
		return err
	}
	_ptr_OlderCredentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.OlderCredentialCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.OlderCredentialCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OlderCredentials", sizeInfo[0])
		}
		o.OlderCredentials = make([]*KerberosKeyDataNew, sizeInfo[0])
		for i1 := range o.OlderCredentials {
			i1 := i1
			_ptr_OlderCredentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.OlderCredentials[i1] == nil {
					o.OlderCredentials[i1] = &KerberosKeyDataNew{}
				}
				if err := o.OlderCredentials[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_OlderCredentials := func(ptr interface{}) { o.OlderCredentials[i1] = *ptr.(**KerberosKeyDataNew) }
			if err := w.ReadPointer(&o.OlderCredentials[i1], _s_OlderCredentials, _ptr_OlderCredentials); err != nil {
				return err
			}
		}
		return nil
	})
	_s_OlderCredentials := func(ptr interface{}) { o.OlderCredentials = *ptr.(*[]*KerberosKeyDataNew) }
	if err := w.ReadPointer(&o.OlderCredentials, _s_OlderCredentials, _ptr_OlderCredentials); err != nil {
		return err
	}
	_ptr_DefaultSalt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DefaultSaltMaximumLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DefaultSaltMaximumLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DefaultSalt", sizeInfo[0])
		}
		o.DefaultSalt = make([]byte, sizeInfo[0])
		for i1 := range o.DefaultSalt {
			i1 := i1
			if err := w.ReadData(&o.DefaultSalt[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_DefaultSalt := func(ptr interface{}) { o.DefaultSalt = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.DefaultSalt, _s_DefaultSalt, _ptr_DefaultSalt); err != nil {
		return err
	}
	_ptr_KeyValues := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.KeyValues = append(o.KeyValues, uint8(0))
			if err := w.ReadData(&o.KeyValues[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_KeyValues := func(ptr interface{}) { o.KeyValues = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.KeyValues, _s_KeyValues, _ptr_KeyValues); err != nil {
		return err
	}
	return nil
}

// UserProperty structure represents USER_PROPERTY RPC structure.
//
// The USER_PROPERTY structure defines an array element that contains a single property
// name and value for the supplementalCredentials attribute.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| NameLength                                                    | ValueLength                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved                                                      | PropertyName (variable)                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| PropertyValue (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type UserProperty struct {
	// NameLength (2 bytes): The number of bytes, in little-endian byte order, of PropertyName.
	// The property name is located at an offset of zero bytes just following the Reserved
	// field. For more information, see the message processing section for supplementalCredentials
	// (section 3.1.1.8.11).
	NameLength uint16 `idl:"name:NameLength" json:"name_length"`
	// ValueLength (2 bytes): The number of bytes contained in PropertyValue.
	ValueLength uint16 `idl:"name:ValueLength" json:"value_length"`
	// Reserved (2 bytes): This value MUST be ignored by the recipient and MAY<24> be set
	// to arbitrary values on update.
	_ uint16 `idl:"name:Reserved"`
	// PropertyName (variable): The name of this property as a UTF-16 encoded string.
	PropertyName     string `idl:"name:PropertyName;size_is:((NameLength/2))" json:"property_name"`
	PropertyValueRaw []byte `idl:"name:PropertyValueRaw;size_is:(ValueLength)" json:"property_value_raw"`
	// PropertyValue (variable): The value of this property. The value MUST be hexadecimal-encoded
	// using an 8-bit character size, and the values '0' through '9' inclusive and 'a' through
	// 'f' inclusive (the specification of 'a' through 'f' is case-sensitive).
	PropertyValue *UserProperty_PropertyValue `idl:"name:PropertyValue;switch_is:PropertyName" json:"property_value"`
}

func (o *UserProperty) xxx_PreparePayload(ctx context.Context) error {
	if o.PropertyName != "" && o.NameLength == 0 {
		o.NameLength = uint16((len(o.PropertyName) * 2))
	}
	if o.PropertyValueRaw != nil && o.ValueLength == 0 {
		o.ValueLength = uint16(len(o.PropertyValueRaw))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.NameLength); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueLength); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	if o.PropertyName != "" || (o.NameLength/2) > 0 {
		_ptr_PropertyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64((o.NameLength / 2))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_PropertyName_buf := utf16.Encode([]rune(o.PropertyName))
			if uint64(len(_PropertyName_buf)) > sizeInfo[0] {
				_PropertyName_buf = _PropertyName_buf[:sizeInfo[0]]
			}
			for i1 := range _PropertyName_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_PropertyName_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_PropertyName_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PropertyName, _ptr_PropertyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PropertyValueRaw != nil || o.ValueLength > 0 {
		_ptr_PropertyValueRaw := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValueLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.PropertyValueRaw {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.PropertyValueRaw[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.PropertyValueRaw); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PropertyValueRaw, _ptr_PropertyValueRaw); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueLength); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint16
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	_ptr_PropertyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if (o.NameLength/2) > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64((o.NameLength / 2))
		}
		var _PropertyName_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _PropertyName_buf", sizeInfo[0])
		}
		_PropertyName_buf = make([]uint16, sizeInfo[0])
		for i1 := range _PropertyName_buf {
			i1 := i1
			if err := w.ReadData(&_PropertyName_buf[i1]); err != nil {
				return err
			}
		}
		o.PropertyName = strings.TrimRight(string(utf16.Decode(_PropertyName_buf)), ndr.ZeroString)
		return nil
	})
	_s_PropertyName := func(ptr interface{}) { o.PropertyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PropertyName, _s_PropertyName, _ptr_PropertyName); err != nil {
		return err
	}
	_ptr_PropertyValueRaw := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValueLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValueLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.PropertyValueRaw", sizeInfo[0])
		}
		o.PropertyValueRaw = make([]byte, sizeInfo[0])
		for i1 := range o.PropertyValueRaw {
			i1 := i1
			if err := w.ReadData(&o.PropertyValueRaw[i1]); err != nil {
				return err
			}
		}
		_hex_PropertyValueRaw, err := hex.DecodeString(string(o.PropertyValueRaw))
		if err != nil {
			return err
		}
		o.PropertyValueRaw = _hex_PropertyValueRaw
		_layout_PropertyValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyValue == nil {
				o.PropertyValue = &UserProperty_PropertyValue{}
			}
			_swPropertyValue := string(o.PropertyName)
			if err := o.PropertyValue.UnmarshalUnionNDR(ctx, w, _swPropertyValue); err != nil {
				return err
			}
			return nil
		})
		if len(o.PropertyValueRaw) > 0 {
			if err := w.WithBytes(o.PropertyValueRaw).Unmarshal(ctx, _layout_PropertyValue); err != nil {
				return err
			}
		}
		return nil
	})
	_s_PropertyValueRaw := func(ptr interface{}) { o.PropertyValueRaw = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.PropertyValueRaw, _s_PropertyValueRaw, _ptr_PropertyValueRaw); err != nil {
		return err
	}
	return nil
}

// UserProperty_PropertyValue structure represents USER_PROPERTY union anonymous member.
//
// The USER_PROPERTY structure defines an array element that contains a single property
// name and value for the supplementalCredentials attribute.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| NameLength                                                    | ValueLength                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved                                                      | PropertyName (variable)                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| PropertyValue (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type UserProperty_PropertyValue struct {
	// Types that are assignable to Value
	//
	// *UserProperty_PropertyValue_PackagesCredential
	// *UserProperty_PropertyValue_WDigestCredential
	// *UserProperty_PropertyValue_KerberosStoredCredential
	// *UserProperty_PropertyValue_KerberosStoredCredentialNew
	// *UserProperty_PropertyValue_NTLMStrongNTOWF
	// *UserProperty_PropertyValue_CleartextCredential
	// *UserProperty_PropertyValue_RawCredential
	Value is_UserProperty_PropertyValue `json:"value"`
}

func (o *UserProperty_PropertyValue) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *UserProperty_PropertyValue_PackagesCredential:
		if value != nil {
			return value.PackagesCredential
		}
	case *UserProperty_PropertyValue_WDigestCredential:
		if value != nil {
			return value.WDigestCredential
		}
	case *UserProperty_PropertyValue_KerberosStoredCredential:
		if value != nil {
			return value.KerberosStoredCredential
		}
	case *UserProperty_PropertyValue_KerberosStoredCredentialNew:
		if value != nil {
			return value.KerberosStoredCredentialNew
		}
	case *UserProperty_PropertyValue_NTLMStrongNTOWF:
		if value != nil {
			return value.NTLMStrongNTOWF
		}
	case *UserProperty_PropertyValue_CleartextCredential:
		if value != nil {
			return value.CleartextCredential
		}
	case *UserProperty_PropertyValue_RawCredential:
		if value != nil {
			return value.RawCredential
		}
	}
	return nil
}

type is_UserProperty_PropertyValue interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_UserProperty_PropertyValue()
}

func (o *UserProperty_PropertyValue) NDRSwitchValue(sw string) string {
	if o == nil {
		return string("")
	}
	switch (interface{})(o.Value).(type) {
	case *UserProperty_PropertyValue_PackagesCredential:
		return string("Packages")
	case *UserProperty_PropertyValue_WDigestCredential:
		return string("Primary:WDigest")
	case *UserProperty_PropertyValue_KerberosStoredCredential:
		return string("Primary:Kerberos")
	case *UserProperty_PropertyValue_KerberosStoredCredentialNew:
		return string("Primary:Kerberos-Newer-Keys")
	case *UserProperty_PropertyValue_NTLMStrongNTOWF:
		return string("Primary:NTLM-Strong-NTOWF")
	case *UserProperty_PropertyValue_CleartextCredential:
		return string("Primary:CLEARTEXT")
	}
	return string("")
}
func (o *UserProperty_PropertyValue) NDRLayout() {}

func (o *UserProperty_PropertyValue) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw string) error {
	if err := w.WriteSwitch(string(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	switch sw {
	case string("Packages"):
		_o, _ := o.Value.(*UserProperty_PropertyValue_PackagesCredential)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserProperty_PropertyValue_PackagesCredential{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case string("Primary:WDigest"):
		_o, _ := o.Value.(*UserProperty_PropertyValue_WDigestCredential)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserProperty_PropertyValue_WDigestCredential{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case string("Primary:Kerberos"):
		_o, _ := o.Value.(*UserProperty_PropertyValue_KerberosStoredCredential)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserProperty_PropertyValue_KerberosStoredCredential{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case string("Primary:Kerberos-Newer-Keys"):
		_o, _ := o.Value.(*UserProperty_PropertyValue_KerberosStoredCredentialNew)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserProperty_PropertyValue_KerberosStoredCredentialNew{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case string("Primary:NTLM-Strong-NTOWF"):
		_o, _ := o.Value.(*UserProperty_PropertyValue_NTLMStrongNTOWF)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserProperty_PropertyValue_NTLMStrongNTOWF{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case string("Primary:CLEARTEXT"):
		_o, _ := o.Value.(*UserProperty_PropertyValue_CleartextCredential)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserProperty_PropertyValue_CleartextCredential{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		_o, _ := o.Value.(*UserProperty_PropertyValue_RawCredential)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserProperty_PropertyValue_RawCredential{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *UserProperty_PropertyValue) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw string) error {
	if err := w.ReadSwitch((*string)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	switch sw {
	case string("Packages"):
		o.Value = &UserProperty_PropertyValue_PackagesCredential{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case string("Primary:WDigest"):
		o.Value = &UserProperty_PropertyValue_WDigestCredential{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case string("Primary:Kerberos"):
		o.Value = &UserProperty_PropertyValue_KerberosStoredCredential{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case string("Primary:Kerberos-Newer-Keys"):
		o.Value = &UserProperty_PropertyValue_KerberosStoredCredentialNew{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case string("Primary:NTLM-Strong-NTOWF"):
		o.Value = &UserProperty_PropertyValue_NTLMStrongNTOWF{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case string("Primary:CLEARTEXT"):
		o.Value = &UserProperty_PropertyValue_CleartextCredential{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		o.Value = &UserProperty_PropertyValue_RawCredential{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// UserProperty_PropertyValue_PackagesCredential structure represents UserProperty_PropertyValue RPC union arm.
//
// It has following labels: "Packages"
type UserProperty_PropertyValue_PackagesCredential struct {
	PackagesCredential *PackagesCredentials `idl:"name:PackagesCredential" json:"packages_credential"`
}

func (*UserProperty_PropertyValue_PackagesCredential) is_UserProperty_PropertyValue() {}

func (o *UserProperty_PropertyValue_PackagesCredential) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PackagesCredential != nil {
		_ptr_PackagesCredential := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PackagesCredential != nil {
				if err := o.PackagesCredential.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PackagesCredentials{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PackagesCredential, _ptr_PackagesCredential); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProperty_PropertyValue_PackagesCredential) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_PackagesCredential := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PackagesCredential == nil {
			o.PackagesCredential = &PackagesCredentials{}
		}
		if err := o.PackagesCredential.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_PackagesCredential := func(ptr interface{}) { o.PackagesCredential = *ptr.(**PackagesCredentials) }
	if err := w.ReadPointer(&o.PackagesCredential, _s_PackagesCredential, _ptr_PackagesCredential); err != nil {
		return err
	}
	return nil
}

// UserProperty_PropertyValue_WDigestCredential structure represents UserProperty_PropertyValue RPC union arm.
//
// It has following labels: "Primary:WDigest"
type UserProperty_PropertyValue_WDigestCredential struct {
	WDigestCredential *WdigestCredentials `idl:"name:WDigestCredential" json:"w_digest_credential"`
}

func (*UserProperty_PropertyValue_WDigestCredential) is_UserProperty_PropertyValue() {}

func (o *UserProperty_PropertyValue_WDigestCredential) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.WDigestCredential != nil {
		_ptr_WDigestCredential := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.WDigestCredential != nil {
				if err := o.WDigestCredential.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WdigestCredentials{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.WDigestCredential, _ptr_WDigestCredential); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProperty_PropertyValue_WDigestCredential) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_WDigestCredential := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.WDigestCredential == nil {
			o.WDigestCredential = &WdigestCredentials{}
		}
		if err := o.WDigestCredential.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_WDigestCredential := func(ptr interface{}) { o.WDigestCredential = *ptr.(**WdigestCredentials) }
	if err := w.ReadPointer(&o.WDigestCredential, _s_WDigestCredential, _ptr_WDigestCredential); err != nil {
		return err
	}
	return nil
}

// UserProperty_PropertyValue_KerberosStoredCredential structure represents UserProperty_PropertyValue RPC union arm.
//
// It has following labels: "Primary:Kerberos"
type UserProperty_PropertyValue_KerberosStoredCredential struct {
	KerberosStoredCredential *KerberosStoredCredential `idl:"name:KerbStoredCredential" json:"kerberos_stored_credential"`
}

func (*UserProperty_PropertyValue_KerberosStoredCredential) is_UserProperty_PropertyValue() {}

func (o *UserProperty_PropertyValue_KerberosStoredCredential) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.KerberosStoredCredential != nil {
		_ptr_KerbStoredCredential := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.KerberosStoredCredential != nil {
				if err := o.KerberosStoredCredential.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&KerberosStoredCredential{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.KerberosStoredCredential, _ptr_KerbStoredCredential); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProperty_PropertyValue_KerberosStoredCredential) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_KerbStoredCredential := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.KerberosStoredCredential == nil {
			o.KerberosStoredCredential = &KerberosStoredCredential{}
		}
		if err := o.KerberosStoredCredential.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_KerbStoredCredential := func(ptr interface{}) { o.KerberosStoredCredential = *ptr.(**KerberosStoredCredential) }
	if err := w.ReadPointer(&o.KerberosStoredCredential, _s_KerbStoredCredential, _ptr_KerbStoredCredential); err != nil {
		return err
	}
	return nil
}

// UserProperty_PropertyValue_KerberosStoredCredentialNew structure represents UserProperty_PropertyValue RPC union arm.
//
// It has following labels: "Primary:Kerberos-Newer-Keys"
type UserProperty_PropertyValue_KerberosStoredCredentialNew struct {
	KerberosStoredCredentialNew *KerberosStoredCredentialNew `idl:"name:KerbStoredCredentialNew" json:"kerberos_stored_credential_new"`
}

func (*UserProperty_PropertyValue_KerberosStoredCredentialNew) is_UserProperty_PropertyValue() {}

func (o *UserProperty_PropertyValue_KerberosStoredCredentialNew) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.KerberosStoredCredentialNew != nil {
		_ptr_KerbStoredCredentialNew := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.KerberosStoredCredentialNew != nil {
				if err := o.KerberosStoredCredentialNew.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&KerberosStoredCredentialNew{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.KerberosStoredCredentialNew, _ptr_KerbStoredCredentialNew); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProperty_PropertyValue_KerberosStoredCredentialNew) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_KerbStoredCredentialNew := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.KerberosStoredCredentialNew == nil {
			o.KerberosStoredCredentialNew = &KerberosStoredCredentialNew{}
		}
		if err := o.KerberosStoredCredentialNew.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_KerbStoredCredentialNew := func(ptr interface{}) { o.KerberosStoredCredentialNew = *ptr.(**KerberosStoredCredentialNew) }
	if err := w.ReadPointer(&o.KerberosStoredCredentialNew, _s_KerbStoredCredentialNew, _ptr_KerbStoredCredentialNew); err != nil {
		return err
	}
	return nil
}

// UserProperty_PropertyValue_NTLMStrongNTOWF structure represents UserProperty_PropertyValue RPC union arm.
//
// It has following labels: "Primary:NTLM-Strong-NTOWF"
type UserProperty_PropertyValue_NTLMStrongNTOWF struct {
	NTLMStrongNTOWF *NTLMStrongNTOWF `idl:"name:NTLMStrongNTOWF" json:"ntlm_strong_ntowf"`
}

func (*UserProperty_PropertyValue_NTLMStrongNTOWF) is_UserProperty_PropertyValue() {}

func (o *UserProperty_PropertyValue_NTLMStrongNTOWF) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.NTLMStrongNTOWF != nil {
		_ptr_NTLMStrongNTOWF := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.NTLMStrongNTOWF != nil {
				if err := o.NTLMStrongNTOWF.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&NTLMStrongNTOWF{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NTLMStrongNTOWF, _ptr_NTLMStrongNTOWF); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProperty_PropertyValue_NTLMStrongNTOWF) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_NTLMStrongNTOWF := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.NTLMStrongNTOWF == nil {
			o.NTLMStrongNTOWF = &NTLMStrongNTOWF{}
		}
		if err := o.NTLMStrongNTOWF.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_NTLMStrongNTOWF := func(ptr interface{}) { o.NTLMStrongNTOWF = *ptr.(**NTLMStrongNTOWF) }
	if err := w.ReadPointer(&o.NTLMStrongNTOWF, _s_NTLMStrongNTOWF, _ptr_NTLMStrongNTOWF); err != nil {
		return err
	}
	return nil
}

// UserProperty_PropertyValue_CleartextCredential structure represents UserProperty_PropertyValue RPC union arm.
//
// It has following labels: "Primary:CLEARTEXT"
type UserProperty_PropertyValue_CleartextCredential struct {
	CleartextCredential *CleartextCredentials `idl:"name:CleartextCredential" json:"cleartext_credential"`
}

func (*UserProperty_PropertyValue_CleartextCredential) is_UserProperty_PropertyValue() {}

func (o *UserProperty_PropertyValue_CleartextCredential) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CleartextCredential != nil {
		_ptr_CleartextCredential := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CleartextCredential != nil {
				if err := o.CleartextCredential.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CleartextCredentials{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CleartextCredential, _ptr_CleartextCredential); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProperty_PropertyValue_CleartextCredential) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_CleartextCredential := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CleartextCredential == nil {
			o.CleartextCredential = &CleartextCredentials{}
		}
		if err := o.CleartextCredential.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_CleartextCredential := func(ptr interface{}) { o.CleartextCredential = *ptr.(**CleartextCredentials) }
	if err := w.ReadPointer(&o.CleartextCredential, _s_CleartextCredential, _ptr_CleartextCredential); err != nil {
		return err
	}
	return nil
}

// UserProperty_PropertyValue_RawCredential structure represents UserProperty_PropertyValue RPC default union arm.
type UserProperty_PropertyValue_RawCredential struct {
	RawCredential *RawCredentials `idl:"name:RawCredential" json:"raw_credential"`
}

func (*UserProperty_PropertyValue_RawCredential) is_UserProperty_PropertyValue() {}

func (o *UserProperty_PropertyValue_RawCredential) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.RawCredential != nil {
		_ptr_RawCredential := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RawCredential != nil {
				if err := o.RawCredential.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RawCredentials{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RawCredential, _ptr_RawCredential); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProperty_PropertyValue_RawCredential) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_RawCredential := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RawCredential == nil {
			o.RawCredential = &RawCredentials{}
		}
		if err := o.RawCredential.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_RawCredential := func(ptr interface{}) { o.RawCredential = *ptr.(**RawCredentials) }
	if err := w.ReadPointer(&o.RawCredential, _s_RawCredential, _ptr_RawCredential); err != nil {
		return err
	}
	return nil
}

// UserProperties structure represents USER_PROPERTIES RPC structure.
//
// The USER_PROPERTIES structure defines the format of the supplementalCredentials attribute.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved1                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Length                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved2                                                     | Reserved3                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved4 (96 bytes)                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| PropertySignature                                             | PropertyCount (optional)                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| UserProperties (variable)                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved5                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type UserProperties struct {
	// Reserved1 (4 bytes): This value MUST be set to zero and MUST be ignored by the recipient.
	_ uint32 `idl:"name:Reserved1"`
	// Length (4 bytes): This value MUST be set to the length, in bytes, of the entire structure,
	// starting from the Reserved4 field.
	Length uint32 `idl:"name:Length" json:"length"`
	// Reserved2 (2 bytes): This value MUST be set to zero and MUST be ignored by the recipient.
	_ uint16 `idl:"name:Reserved2"`
	// Reserved3 (2 bytes): This value MUST be set to zero and MUST be ignored by the recipient.
	_ uint16 `idl:"name:Reserved3"`
	// Reserved4 (96 bytes): This value MUST be ignored by the recipient and MAY<22> contain
	// arbitrary values.
	_ []byte `idl:"name:Reserved4"`
	// PropertySignature (2 bytes): This field MUST be the value 0x50, in little-endian
	// byte order. This is an arbitrary value used to indicate whether the structure is
	// corrupt. That is, if this value is not 0x50 on read, the structure is considered
	// corrupt, processing MUST be aborted, and an error code MUST be returned.
	PropertySignature []byte `idl:"name:PropertySignature" json:"property_signature"`
	// PropertyCount (2 bytes): The number of USER_PROPERTY elements in the UserProperties
	// field. When there are zero USER_PROPERTY elements in the UserProperties field, this
	// field MUST be omitted; the resultant USER_PROPERTIES structure has a constant size
	// of 0x6F bytes.
	PropertyCount     uint16 `idl:"name:PropertyCount" json:"property_count"`
	UserPropertiesRaw []byte `idl:"name:UserPropertiesRaw;size_is:((Length-100))" json:"user_properties_raw"`
	// UserProperties (variable): An array of PropertyCount USER_PROPERTY elements.
	UserProperties []*UserProperty `idl:"name:UserProperties;size_is:(PropertyCount)" json:"user_properties"`
	// Reserved5 (1 byte): This value SHOULD<23> be set to zero and MUST be ignored by the
	// recipient.
	_ uint8 `idl:"name:Reserved5"`
}

func (o *UserProperties) xxx_PreparePayload(ctx context.Context) error {
	if o.UserPropertiesRaw != nil && o.Length == 0 {
		o.Length = uint32((len(o.UserPropertiesRaw) + 100))
	}
	if o.Length < 100 {
		o.Length = 100
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProperties) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	// reserved Reserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	// reserved Reserved2
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	// reserved Reserved3
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	// reserved Reserved4
	for i1 := 0; uint64(i1) < 96; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.PropertySignature {
		i1 := i1
		if uint64(i1) >= 2 {
			break
		}
		if err := w.WriteData(o.PropertySignature[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.PropertySignature); uint64(i1) < 2; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PropertyCount); err != nil {
		return err
	}
	if o.UserPropertiesRaw != nil || (o.Length-100) > 0 {
		_ptr_UserPropertiesRaw := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64((o.Length - 100))
			if o.Length < 100 {
				dimSize1 = uint64(0)
			}
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.UserPropertiesRaw {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.UserPropertiesRaw[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.UserPropertiesRaw); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UserPropertiesRaw, _ptr_UserPropertiesRaw); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved Reserved5
	if err := w.WriteData(uint8(0)); err != nil {
		return err
	}
	return nil
}
func (o *UserProperties) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	// reserved Reserved1
	var _Reserved1 uint32
	if err := w.ReadData(&_Reserved1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	// reserved Reserved2
	var _Reserved2 uint16
	if err := w.ReadData(&_Reserved2); err != nil {
		return err
	}
	// reserved Reserved3
	var _Reserved3 uint16
	if err := w.ReadData(&_Reserved3); err != nil {
		return err
	}
	// reserved Reserved4
	var _Reserved4 []byte
	_Reserved4 = make([]byte, 96)
	for i1 := range _Reserved4 {
		i1 := i1
		if err := w.ReadData(&_Reserved4[i1]); err != nil {
			return err
		}
	}
	o.PropertySignature = make([]byte, 2)
	for i1 := range o.PropertySignature {
		i1 := i1
		if err := w.ReadData(&o.PropertySignature[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PropertyCount); err != nil {
		return err
	}
	_ptr_UserPropertiesRaw := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Length > 100 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64((o.Length - 100))
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.UserPropertiesRaw", sizeInfo[0])
		}
		o.UserPropertiesRaw = make([]byte, sizeInfo[0])
		for i1 := range o.UserPropertiesRaw {
			i1 := i1
			if err := w.ReadData(&o.UserPropertiesRaw[i1]); err != nil {
				return err
			}
		}
		_layout_UserProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			// XXX: for opaque unmarshaling
			if o.PropertyCount > 0 && sizeInfo[0] == 0 {
				sizeInfo[0] = uint64(o.PropertyCount)
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.UserProperties", sizeInfo[0])
			}
			o.UserProperties = make([]*UserProperty, sizeInfo[0])
			for i1 := range o.UserProperties {
				i1 := i1
				_ptr_UserProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if o.UserProperties[i1] == nil {
						o.UserProperties[i1] = &UserProperty{}
					}
					if err := o.UserProperties[i1].UnmarshalNDR(ctx, w); err != nil {
						return err
					}
					return nil
				})
				_s_UserProperties := func(ptr interface{}) { o.UserProperties[i1] = *ptr.(**UserProperty) }
				if err := w.ReadPointer(&o.UserProperties[i1], _s_UserProperties, _ptr_UserProperties); err != nil {
					return err
				}
			}
			return nil
		})
		if len(o.UserPropertiesRaw) > 0 {
			if err := w.WithBytes(o.UserPropertiesRaw).Unmarshal(ctx, _layout_UserProperties); err != nil {
				return err
			}
		}
		return nil
	})
	_s_UserPropertiesRaw := func(ptr interface{}) { o.UserPropertiesRaw = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.UserPropertiesRaw, _s_UserPropertiesRaw, _ptr_UserPropertiesRaw); err != nil {
		return err
	}
	// reserved Reserved5
	var _Reserved5 uint8
	if err := w.ReadData(&_Reserved5); err != nil {
		return err
	}
	return nil
}

// String structure represents RPC_STRING RPC structure.
//
// The RPC_STRING structure holds a counted string encoded in the OEM code page.
type String struct {
	// Length:  The size, in bytes, not including a terminating null character, of the string
	// contained in Buffer.
	Length uint16 `idl:"name:Length" json:"length"`
	// MaximumLength:  The size, in bytes, of the Buffer member.
	MaximumLength uint16 `idl:"name:MaximumLength" json:"maximum_length"`
	// Buffer:  A buffer containing a string encoded in the OEM code page. The string is
	// counted (by the Length member), and therefore is not null-terminated.
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

// OldLargeInteger structure represents OLD_LARGE_INTEGER RPC structure.
//
// The OLD_LARGE_INTEGER structure defines a 64-bit value that is accessible in two
// 4-byte chunks.
type OldLargeInteger struct {
	// LowPart:  The least-significant portion of a 64-bit value.
	LowPart uint32 `idl:"name:LowPart" json:"low_part"`
	// HighPart:  The most-significant portion of a 64-bit value.
	HighPart int32 `idl:"name:HighPart" json:"high_part"`
}

func (o *OldLargeInteger) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OldLargeInteger) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LowPart); err != nil {
		return err
	}
	if err := w.WriteData(o.HighPart); err != nil {
		return err
	}
	return nil
}
func (o *OldLargeInteger) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LowPart); err != nil {
		return err
	}
	if err := w.ReadData(&o.HighPart); err != nil {
		return err
	}
	return nil
}

// Handle structure represents SAMPR_HANDLE RPC structure.
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

// EncryptedLMOWFPassword structure represents ENCRYPTED_LM_OWF_PASSWORD RPC structure.
//
// The ENCRYPTED_LM_OWF_PASSWORD structure defines a block of encrypted data used in
// various methods to communicate sensitive information.
type EncryptedLMOWFPassword struct {
	// data:  16 bytes of unstructured data used to hold an encrypted 16-byte hash (either
	// an LM hash or an NT hash). The encryption algorithm is specified in section 2.2.11.1.
	// The methods specified in sections 3.1.5.10 and 3.1.5.13.6 use this structure and
	// specify the type of hash and the encryption key.
	Data []byte `idl:"name:data" json:"data"`
}

func (o *EncryptedLMOWFPassword) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedLMOWFPassword) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedLMOWFPassword) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Data = make([]byte, 16)
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// EncryptedNTOWFPassword structure represents ENCRYPTED_NT_OWF_PASSWORD RPC structure.
type EncryptedNTOWFPassword struct {
	Data []byte `idl:"name:data" json:"data"`
}

func (o *EncryptedNTOWFPassword) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedNTOWFPassword) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedNTOWFPassword) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Data = make([]byte, 16)
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Uint32Array structure represents SAMPR_ULONG_ARRAY RPC structure.
//
// The SAMPR_ULONG_ARRAY structure holds a counted array of unsigned long values.
type Uint32Array struct {
	// Count:  The number of elements in Element. If zero, Element MUST be ignored. If nonzero,
	// Element MUST point to at least Count * sizeof(unsigned long) bytes of memory.
	Count uint32 `idl:"name:Count" json:"count"`
	// Element:  A pointer to an array of unsigned integers with Count elements. The semantic
	// meaning is dependent on the method in which the structure is being used.
	Element []uint32 `idl:"name:Element;size_is:(Count)" json:"element"`
}

func (o *Uint32Array) xxx_PreparePayload(ctx context.Context) error {
	if o.Element != nil && o.Count == 0 {
		o.Count = uint32(len(o.Element))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Uint32Array) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Element != nil || o.Count > 0 {
		_ptr_Element := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Element {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Element[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Element); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Element, _ptr_Element); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Uint32Array) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_Element := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Element", sizeInfo[0])
		}
		o.Element = make([]uint32, sizeInfo[0])
		for i1 := range o.Element {
			i1 := i1
			if err := w.ReadData(&o.Element[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Element := func(ptr interface{}) { o.Element = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Element, _s_Element, _ptr_Element); err != nil {
		return err
	}
	return nil
}

// SIDInformation structure represents SAMPR_SID_INFORMATION RPC structure.
//
// The SAMPR_PSID_ARRAY structure holds an array of SID values.
type SIDInformation struct {
	SIDPointer *dtyp.SID `idl:"name:SidPointer" json:"sid_pointer"`
}

func (o *SIDInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SIDInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.SIDPointer != nil {
		_ptr_SidPointer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SIDPointer != nil {
				if err := o.SIDPointer.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SIDPointer, _ptr_SidPointer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SIDInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_SidPointer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SIDPointer == nil {
			o.SIDPointer = &dtyp.SID{}
		}
		if err := o.SIDPointer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SidPointer := func(ptr interface{}) { o.SIDPointer = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SIDPointer, _s_SidPointer, _ptr_SidPointer); err != nil {
		return err
	}
	return nil
}

// SIDArray structure represents SAMPR_PSID_ARRAY RPC structure.
//
// The SAMPR_SID_INFORMATION structure holds a SID pointer.
type SIDArray struct {
	Count uint32            `idl:"name:Count" json:"count"`
	SIDs  []*SIDInformation `idl:"name:Sids;size_is:(Count)" json:"sids"`
}

func (o *SIDArray) xxx_PreparePayload(ctx context.Context) error {
	if o.SIDs != nil && o.Count == 0 {
		o.Count = uint32(len(o.SIDs))
	}
	if o.Count > uint32(1024) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SIDArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.SIDs != nil || o.Count > 0 {
		_ptr_Sids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.SIDs[i1] != nil {
					if err := o.SIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SIDInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.SIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SIDInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SIDs, _ptr_Sids); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SIDArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_Sids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SIDs", sizeInfo[0])
		}
		o.SIDs = make([]*SIDInformation, sizeInfo[0])
		for i1 := range o.SIDs {
			i1 := i1
			if o.SIDs[i1] == nil {
				o.SIDs[i1] = &SIDInformation{}
			}
			if err := o.SIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Sids := func(ptr interface{}) { o.SIDs = *ptr.(*[]*SIDInformation) }
	if err := w.ReadPointer(&o.SIDs, _s_Sids, _ptr_Sids); err != nil {
		return err
	}
	return nil
}

// SIDArrayOut structure represents SAMPR_PSID_ARRAY_OUT RPC structure.
//
// The SAMPR_PSID_ARRAY_OUT structure holds an array of SID values.
type SIDArrayOut struct {
	// Count:  The number of elements in Sids. If zero, Sids MUST be ignored. If nonzero,
	// Sids MUST point to at least Count * sizeof(SAMPR_SID_INFORMATION) bytes of memory.
	Count uint32 `idl:"name:Count" json:"count"`
	// Sids:  An array of pointers to SID values. For more information, see section 2.2.7.5.
	SIDs []*SIDInformation `idl:"name:Sids;size_is:(Count)" json:"sids"`
}

func (o *SIDArrayOut) xxx_PreparePayload(ctx context.Context) error {
	if o.SIDs != nil && o.Count == 0 {
		o.Count = uint32(len(o.SIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SIDArrayOut) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.SIDs != nil || o.Count > 0 {
		_ptr_Sids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.SIDs[i1] != nil {
					if err := o.SIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SIDInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.SIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SIDInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SIDs, _ptr_Sids); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SIDArrayOut) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_Sids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SIDs", sizeInfo[0])
		}
		o.SIDs = make([]*SIDInformation, sizeInfo[0])
		for i1 := range o.SIDs {
			i1 := i1
			if o.SIDs[i1] == nil {
				o.SIDs[i1] = &SIDInformation{}
			}
			if err := o.SIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Sids := func(ptr interface{}) { o.SIDs = *ptr.(*[]*SIDInformation) }
	if err := w.ReadPointer(&o.SIDs, _s_Sids, _ptr_Sids); err != nil {
		return err
	}
	return nil
}

// ReturnedUstringArray structure represents SAMPR_RETURNED_USTRING_ARRAY RPC structure.
//
// The SAMPR_RETURNED_USTRING_ARRAY structure holds an array of counted UTF-16 encoded
// strings.
type ReturnedUstringArray struct {
	// Count:  The number of elements in Element. If zero, Element MUST be ignored. If nonzero,
	// Element MUST point to at least Count * sizeof(RPC_UNICODE_STRING) bytes of memory.
	Count uint32 `idl:"name:Count" json:"count"`
	// Element:  Array of counted strings (see RPC_UNICODE_STRING in [MS-DTYP] section 2.3.10).
	// The semantic meaning is method-dependent.
	Element []*dtyp.UnicodeString `idl:"name:Element;size_is:(Count)" json:"element"`
}

func (o *ReturnedUstringArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Element != nil && o.Count == 0 {
		o.Count = uint32(len(o.Element))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReturnedUstringArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Element != nil || o.Count > 0 {
		_ptr_Element := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Element {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Element[i1] != nil {
					if err := o.Element[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Element); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Element, _ptr_Element); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReturnedUstringArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_Element := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Element", sizeInfo[0])
		}
		o.Element = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.Element {
			i1 := i1
			if o.Element[i1] == nil {
				o.Element[i1] = &dtyp.UnicodeString{}
			}
			if err := o.Element[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Element := func(ptr interface{}) { o.Element = *ptr.(*[]*dtyp.UnicodeString) }
	if err := w.ReadPointer(&o.Element, _s_Element, _ptr_Element); err != nil {
		return err
	}
	return nil
}

// SIDNameUse type represents SID_NAME_USE RPC enumeration.
//
// The SID_NAME_USE enumeration specifies the type of account that a SID references.
type SIDNameUse uint16

var (
	// SidTypeUser: Indicates a user object.
	SIDNameUseTypeUser SIDNameUse = 1
	// SidTypeGroup: Indicates a group object.
	SIDNameUseTypeGroup SIDNameUse = 2
	// SidTypeDomain: Indicates a domain object.
	SIDNameUseTypeDomain SIDNameUse = 3
	// SidTypeAlias: Indicates an alias object.
	SIDNameUseTypeAlias SIDNameUse = 4
	// SidTypeWellKnownGroup: Indicates an object whose SID is invariant.
	SIDNameUseTypeWellKnownGroup SIDNameUse = 5
	// SidTypeDeletedAccount: Indicates an object that has been deleted.
	SIDNameUseTypeDeletedAccount SIDNameUse = 6
	// SidTypeInvalid: This member is not used.
	SIDNameUseTypeInvalid SIDNameUse = 7
	// SidTypeUnknown: Indicates that the type of object could not be determined. For example,
	// no object with that SID exists.
	SIDNameUseTypeUnknown SIDNameUse = 8
	// SidTypeComputer: This member is not used.
	SIDNameUseTypeComputer SIDNameUse = 9
	// SidTypeLabel: This member is not used.
	SIDNameUseTypeLabel SIDNameUse = 10
)

func (o SIDNameUse) String() string {
	switch o {
	case SIDNameUseTypeUser:
		return "SIDNameUseTypeUser"
	case SIDNameUseTypeGroup:
		return "SIDNameUseTypeGroup"
	case SIDNameUseTypeDomain:
		return "SIDNameUseTypeDomain"
	case SIDNameUseTypeAlias:
		return "SIDNameUseTypeAlias"
	case SIDNameUseTypeWellKnownGroup:
		return "SIDNameUseTypeWellKnownGroup"
	case SIDNameUseTypeDeletedAccount:
		return "SIDNameUseTypeDeletedAccount"
	case SIDNameUseTypeInvalid:
		return "SIDNameUseTypeInvalid"
	case SIDNameUseTypeUnknown:
		return "SIDNameUseTypeUnknown"
	case SIDNameUseTypeComputer:
		return "SIDNameUseTypeComputer"
	case SIDNameUseTypeLabel:
		return "SIDNameUseTypeLabel"
	}
	return "Invalid"
}

// ShortBlob structure represents RPC_SHORT_BLOB RPC structure.
//
// The RPC_SHORT_BLOB structure holds a counted array of unsigned short values.
type ShortBlob struct {
	// Length:  The number of bytes of data contained in the Buffer member.
	Length uint16 `idl:"name:Length" json:"length"`
	// MaximumLength:  The length, in bytes, of the Buffer member.
	MaximumLength uint16 `idl:"name:MaximumLength" json:"maximum_length"`
	// Buffer:  A buffer containing Length/2 unsigned short values.
	Buffer []uint16 `idl:"name:Buffer;size_is:((MaximumLength/2));length_is:((Length/2))" json:"buffer"`
}

func (o *ShortBlob) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.MaximumLength == 0 {
		o.MaximumLength = uint16((len(o.Buffer) * 2))
	}
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint16((len(o.Buffer) * 2))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ShortBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.Buffer != nil || (o.MaximumLength/2) > 0 {
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
func (o *ShortBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]uint16, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]uint16) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// RIDEnumeration structure represents SAMPR_RID_ENUMERATION RPC structure.
//
// The SAMPR_RID_ENUMERATION structure holds the name and RID information about an account.
type RIDEnumeration struct {
	// RelativeId:  A RID.
	RelativeID uint32 `idl:"name:RelativeId" json:"relative_id"`
	// Name:  The UTF-16 encoded name of the account that is associated with RelativeId.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
}

func (o *RIDEnumeration) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RIDEnumeration) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RelativeID); err != nil {
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
func (o *RIDEnumeration) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RelativeID); err != nil {
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

// EnumerationBuffer structure represents SAMPR_ENUMERATION_BUFFER RPC structure.
//
// The SAMPR_ENUMERATION_BUFFER structure holds an array of SAMPR_RID_ENUMERATION elements.
type EnumerationBuffer struct {
	// EntriesRead:  The number of elements in Buffer. If zero, Buffer MUST be ignored.
	// If nonzero, Buffer MUST point to at least EntriesRead * sizeof(SAMPR_RID_ENUMERATION)
	// bytes of memory.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  An array of SAMPR_RID_ENUMERATION elements.
	Buffer []*RIDEnumeration `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *EnumerationBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EnumerationBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
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
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&RIDEnumeration{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&RIDEnumeration{}).MarshalNDR(ctx, w); err != nil {
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
func (o *EnumerationBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
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
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*RIDEnumeration, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &RIDEnumeration{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*RIDEnumeration) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// SrSecurityDescriptor structure represents SAMPR_SR_SECURITY_DESCRIPTOR RPC structure.
//
// The SAMPR_SR_SECURITY_DESCRIPTOR structure holds a formatted security descriptor.
type SrSecurityDescriptor struct {
	// Length:  The size, in bytes, of SecurityDescriptor. If zero, SecurityDescriptor MUST
	// be ignored. The maximum size of 256 * 1024 is an arbitrary value chosen to limit
	// the amount of memory a client can force the server to allocate.
	Length uint32 `idl:"name:Length" json:"length"`
	// SecurityDescriptor:  A binary format per the SECURITY_DESCRIPTOR format in [MS-DTYP]
	// section 2.4.6.
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

// GroupMembership structure represents GROUP_MEMBERSHIP RPC structure.
//
// The GROUP_MEMBERSHIP structure holds information on a group membership.
type GroupMembership struct {
	// RelativeId:  A RID that represents one membership value.
	RelativeID uint32 `idl:"name:RelativeId" json:"relative_id"`
	// Attributes:  Characteristics about the membership represented as a bitmask. Values
	// are defined in section 2.2.1.10.
	Attributes uint32 `idl:"name:Attributes" json:"attributes"`
}

func (o *GroupMembership) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupMembership) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RelativeID); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	return nil
}
func (o *GroupMembership) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RelativeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	return nil
}

// GetGroupsBuffer structure represents SAMPR_GET_GROUPS_BUFFER RPC structure.
//
// The SAMPR_GET_GROUPS_BUFFER structure represents the members of a group.
type GetGroupsBuffer struct {
	// MembershipCount:  The number of elements in Groups. If zero, Groups MUST be ignored.
	// If nonzero, Groups MUST point to at least MembershipCount * sizeof(GROUP_MEMBERSHIP)
	// bytes of memory.
	MembershipCount uint32 `idl:"name:MembershipCount" json:"membership_count"`
	// Groups:  An array to hold information about the members of the group.
	Groups []*GroupMembership `idl:"name:Groups;size_is:(MembershipCount)" json:"groups"`
}

func (o *GetGroupsBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Groups != nil && o.MembershipCount == 0 {
		o.MembershipCount = uint32(len(o.Groups))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetGroupsBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MembershipCount); err != nil {
		return err
	}
	if o.Groups != nil || o.MembershipCount > 0 {
		_ptr_Groups := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MembershipCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Groups {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Groups[i1] != nil {
					if err := o.Groups[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Groups); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Groups, _ptr_Groups); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetGroupsBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MembershipCount); err != nil {
		return err
	}
	_ptr_Groups := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.MembershipCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.MembershipCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Groups", sizeInfo[0])
		}
		o.Groups = make([]*GroupMembership, sizeInfo[0])
		for i1 := range o.Groups {
			i1 := i1
			if o.Groups[i1] == nil {
				o.Groups[i1] = &GroupMembership{}
			}
			if err := o.Groups[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Groups := func(ptr interface{}) { o.Groups = *ptr.(*[]*GroupMembership) }
	if err := w.ReadPointer(&o.Groups, _s_Groups, _ptr_Groups); err != nil {
		return err
	}
	return nil
}

// GetMembersBuffer structure represents SAMPR_GET_MEMBERS_BUFFER RPC structure.
//
// The SAMPR_GET_MEMBERS_BUFFER structure represents the membership of a group.
type GetMembersBuffer struct {
	// MemberCount:  The number of elements in Members and Attributes. If zero, Members
	// and Attributes MUST be ignored. If nonzero, Members and Attributes MUST point to
	// at least MemberCount * sizeof(unsigned long) bytes of memory.
	MemberCount uint32 `idl:"name:MemberCount" json:"member_count"`
	// Members:  An array of RIDs.
	Members []uint32 `idl:"name:Members;size_is:(MemberCount)" json:"members"`
	// Attributes:  Characteristics about the membership, represented as a bitmask. Values
	// are defined in section 2.2.1.10.
	Attributes []uint32 `idl:"name:Attributes;size_is:(MemberCount)" json:"attributes"`
}

func (o *GetMembersBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Members != nil && o.MemberCount == 0 {
		o.MemberCount = uint32(len(o.Members))
	}
	if o.Attributes != nil && o.MemberCount == 0 {
		o.MemberCount = uint32(len(o.Attributes))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetMembersBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MemberCount); err != nil {
		return err
	}
	if o.Members != nil || o.MemberCount > 0 {
		_ptr_Members := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MemberCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Members {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Members[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Members); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Members, _ptr_Members); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Attributes != nil || o.MemberCount > 0 {
		_ptr_Attributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MemberCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Attributes {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Attributes[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Attributes); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Attributes, _ptr_Attributes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *GetMembersBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MemberCount); err != nil {
		return err
	}
	_ptr_Members := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.MemberCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.MemberCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Members", sizeInfo[0])
		}
		o.Members = make([]uint32, sizeInfo[0])
		for i1 := range o.Members {
			i1 := i1
			if err := w.ReadData(&o.Members[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Members := func(ptr interface{}) { o.Members = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Members, _s_Members, _ptr_Members); err != nil {
		return err
	}
	_ptr_Attributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.MemberCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.MemberCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Attributes", sizeInfo[0])
		}
		o.Attributes = make([]uint32, sizeInfo[0])
		for i1 := range o.Attributes {
			i1 := i1
			if err := w.ReadData(&o.Attributes[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Attributes := func(ptr interface{}) { o.Attributes = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Attributes, _s_Attributes, _ptr_Attributes); err != nil {
		return err
	}
	return nil
}

// RevisionInfoV1 structure represents SAMPR_REVISION_INFO_V1 RPC structure.
//
// The SAMPR_REVISION_INFO_V1 structure is used to communicate the revision and capabilities
// of client and server. For more information, see SamrConnect5.
type RevisionInfoV1 struct {
	// Revision:  The revision of the client or server side of this protocol (depending
	// on which side sends the structure). The value MUST be set to 3 and MUST be ignored.
	Revision uint32 `idl:"name:Revision" json:"revision"`
	// SupportedFeatures:  A bit field. When sent from the client, this field MUST be zero
	// and ignored on receipt by the server. When returned from the server, the following
	// fields are handled by the client; all other bits are ignored by the client and MUST
	// be zero when returned from the server.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | On receipt by the client, this value, when set, indicates that RID values        |
	//	|            | returned from the server MUST NOT be concatenated with the domain SID to create  |
	//	|            | the SID for the account referenced by the RID. Instead, the client MUST call     |
	//	|            | SamrRidToSid to obtain the SID. This field can be combined with other bits using |
	//	|            | a logical OR. See the product behavior citation at the end of this section for   |
	//	|            | more information (about Windows implementations).                                |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | Reserved. See the product behavior citation at the end of this section for       |
	//	|            | additional details.                                                              |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 | Reserved. See the product behavior citation at the end of this section for       |
	//	|            | additional details.                                                              |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000010 | On receipt by the client, this value, when set, indicates that the client should |
	//	|            | use AES Encryption with the SAMPR_ENCRYPTED_PASSWORD_AES structure to encrypt    |
	//	|            | password buffers when sent over the wire. See AES Cipher Usage (section 3.2.2.4) |
	//	|            | and SAMPR_ENCRYPTED_PASSWORD_AES (section 2.2.6.32).                             |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000020 | On receipt of this value by the client, when set, indicates that the server      |
	//	|            | supports the validation of computer account re-use through client calls to the   |
	//	|            | SamrValidateComputerAccountReuseAttempt method.<19>                              |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000040 | On receipt of this value by the client, when set, indicates that the server      |
	//	|            | supports the validation of computer account re-use through client calls to the   |
	//	|            | SamrValidateComputerAccountReuseAttempt method.<20>                              |
	//	+------------+----------------------------------------------------------------------------------+
	SupportedFeatures uint32 `idl:"name:SupportedFeatures" json:"supported_features"`
}

func (o *RevisionInfoV1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RevisionInfoV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Revision); err != nil {
		return err
	}
	if err := w.WriteData(o.SupportedFeatures); err != nil {
		return err
	}
	return nil
}
func (o *RevisionInfoV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Revision); err != nil {
		return err
	}
	if err := w.ReadData(&o.SupportedFeatures); err != nil {
		return err
	}
	return nil
}

// RevisionInfo structure represents SAMPR_REVISION_INFO RPC union.
//
// The SAMPR_REVISION_INFO union holds revision information structures that are used
// in the SamrConnect5 method.
type RevisionInfo struct {
	// Types that are assignable to Value
	//
	// *RevisionInfo_V1
	Value is_RevisionInfo `json:"value"`
}

func (o *RevisionInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *RevisionInfo_V1:
		if value != nil {
			return value.V1
		}
	}
	return nil
}

type is_RevisionInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_RevisionInfo()
}

func (o *RevisionInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *RevisionInfo_V1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *RevisionInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*RevisionInfo_V1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RevisionInfo_V1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *RevisionInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &RevisionInfo_V1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// RevisionInfo_V1 structure represents SAMPR_REVISION_INFO RPC union arm.
//
// It has following labels: 1
type RevisionInfo_V1 struct {
	// V1:  Version 1 revision information, as described in SAMPR_REVISION_INFO_V1 (section
	// 2.2.7.15).
	V1 *RevisionInfoV1 `idl:"name:V1" json:"v1"`
}

func (*RevisionInfo_V1) is_RevisionInfo() {}

func (o *RevisionInfo_V1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.V1 != nil {
		if err := o.V1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RevisionInfoV1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RevisionInfo_V1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.V1 == nil {
		o.V1 = &RevisionInfoV1{}
	}
	if err := o.V1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserDomainPasswordInformation structure represents USER_DOMAIN_PASSWORD_INFORMATION RPC structure.
//
// The USER_DOMAIN_PASSWORD_INFORMATION structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
type UserDomainPasswordInformation struct {
	MinPasswordLength  uint16 `idl:"name:MinPasswordLength" json:"min_password_length"`
	PasswordProperties uint32 `idl:"name:PasswordProperties" json:"password_properties"`
}

func (o *UserDomainPasswordInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserDomainPasswordInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.MinPasswordLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PasswordProperties); err != nil {
		return err
	}
	return nil
}
func (o *UserDomainPasswordInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinPasswordLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordProperties); err != nil {
		return err
	}
	return nil
}

// DomainServerEnableState type represents DOMAIN_SERVER_ENABLE_STATE RPC enumeration.
//
// The DOMAIN_SERVER_ENABLE_STATE enumeration describes the enabled or disabled state
// of a server.
type DomainServerEnableState uint16

var (
	// DomainServerEnabled:  The server is considered "enabled" to the client.
	DomainServerEnableStateEnabled DomainServerEnableState = 1
	// DomainServerDisabled:  This field is not used.
	DomainServerEnableStateDisabled DomainServerEnableState = 2
)

func (o DomainServerEnableState) String() string {
	switch o {
	case DomainServerEnableStateEnabled:
		return "DomainServerEnableStateEnabled"
	case DomainServerEnableStateDisabled:
		return "DomainServerEnableStateDisabled"
	}
	return "Invalid"
}

// DomainStateInformation structure represents DOMAIN_STATE_INFORMATION RPC structure.
//
// The DOMAIN_STATE_INFORMATION structure holds the enabled/disabled state of the server.
//
// For information on each field, see section 2.2.3.1.
type DomainStateInformation struct {
	DomainServerState DomainServerEnableState `idl:"name:DomainServerState" json:"domain_server_state"`
}

func (o *DomainStateInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainStateInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.DomainServerState)); err != nil {
		return err
	}
	return nil
}
func (o *DomainStateInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.DomainServerState)); err != nil {
		return err
	}
	return nil
}

// DomainServerRole type represents DOMAIN_SERVER_ROLE RPC enumeration.
//
// The DOMAIN_SERVER_ROLE enumeration indicates whether a server is a PDC.
type DomainServerRole uint16

var (
	// DomainServerRoleBackup:  The DC is not the PDC.
	DomainServerRoleBackup DomainServerRole = 2
	// DomainServerRolePrimary:  The DC is the PDC.
	DomainServerRolePrimary DomainServerRole = 3
)

func (o DomainServerRole) String() string {
	switch o {
	case DomainServerRoleBackup:
		return "DomainServerRoleBackup"
	case DomainServerRolePrimary:
		return "DomainServerRolePrimary"
	}
	return "Invalid"
}

// DomainPasswordInformation structure represents DOMAIN_PASSWORD_INFORMATION RPC structure.
//
// The DOMAIN_PASSWORD_INFORMATION structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
type DomainPasswordInformation struct {
	MinPasswordLength     uint16           `idl:"name:MinPasswordLength" json:"min_password_length"`
	PasswordHistoryLength uint16           `idl:"name:PasswordHistoryLength" json:"password_history_length"`
	PasswordProperties    uint32           `idl:"name:PasswordProperties" json:"password_properties"`
	MaxPasswordAge        *OldLargeInteger `idl:"name:MaxPasswordAge" json:"max_password_age"`
	MinPasswordAge        *OldLargeInteger `idl:"name:MinPasswordAge" json:"min_password_age"`
}

func (o *DomainPasswordInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainPasswordInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.MinPasswordLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PasswordHistoryLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PasswordProperties); err != nil {
		return err
	}
	if o.MaxPasswordAge != nil {
		if err := o.MaxPasswordAge.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MinPasswordAge != nil {
		if err := o.MinPasswordAge.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainPasswordInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinPasswordLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordHistoryLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordProperties); err != nil {
		return err
	}
	if o.MaxPasswordAge == nil {
		o.MaxPasswordAge = &OldLargeInteger{}
	}
	if err := o.MaxPasswordAge.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.MinPasswordAge == nil {
		o.MinPasswordAge = &OldLargeInteger{}
	}
	if err := o.MinPasswordAge.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainLogoffInformation structure represents DOMAIN_LOGOFF_INFORMATION RPC structure.
//
// The DOMAIN_LOGOFF_INFORMATION structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
type DomainLogoffInformation struct {
	ForceLogoff *OldLargeInteger `idl:"name:ForceLogoff" json:"force_logoff"`
}

func (o *DomainLogoffInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainLogoffInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.ForceLogoff != nil {
		if err := o.ForceLogoff.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainLogoffInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.ForceLogoff == nil {
		o.ForceLogoff = &OldLargeInteger{}
	}
	if err := o.ForceLogoff.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainServerRoleInformation structure represents DOMAIN_SERVER_ROLE_INFORMATION RPC structure.
//
// The DOMAIN_SERVER_ROLE_INFORMATION structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
type DomainServerRoleInformation struct {
	DomainServerRole DomainServerRole `idl:"name:DomainServerRole" json:"domain_server_role"`
}

func (o *DomainServerRoleInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainServerRoleInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.DomainServerRole)); err != nil {
		return err
	}
	return nil
}
func (o *DomainServerRoleInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.DomainServerRole)); err != nil {
		return err
	}
	return nil
}

// DomainModifiedInformation structure represents DOMAIN_MODIFIED_INFORMATION RPC structure.
//
// The DOMAIN_MODIFIED_INFORMATION structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
type DomainModifiedInformation struct {
	DomainModifiedCount *OldLargeInteger `idl:"name:DomainModifiedCount" json:"domain_modified_count"`
	CreationTime        *OldLargeInteger `idl:"name:CreationTime" json:"creation_time"`
}

func (o *DomainModifiedInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainModifiedInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.DomainModifiedCount != nil {
		if err := o.DomainModifiedCount.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CreationTime != nil {
		if err := o.CreationTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainModifiedInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.DomainModifiedCount == nil {
		o.DomainModifiedCount = &OldLargeInteger{}
	}
	if err := o.DomainModifiedCount.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.CreationTime == nil {
		o.CreationTime = &OldLargeInteger{}
	}
	if err := o.CreationTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainModifiedInformation2 structure represents DOMAIN_MODIFIED_INFORMATION2 RPC structure.
//
// The DOMAIN_MODIFIED_INFORMATION2 structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
type DomainModifiedInformation2 struct {
	DomainModifiedCount          *OldLargeInteger `idl:"name:DomainModifiedCount" json:"domain_modified_count"`
	CreationTime                 *OldLargeInteger `idl:"name:CreationTime" json:"creation_time"`
	ModifiedCountATLastPromotion *OldLargeInteger `idl:"name:ModifiedCountAtLastPromotion" json:"modified_count_at_last_promotion"`
}

func (o *DomainModifiedInformation2) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainModifiedInformation2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.DomainModifiedCount != nil {
		if err := o.DomainModifiedCount.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CreationTime != nil {
		if err := o.CreationTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ModifiedCountATLastPromotion != nil {
		if err := o.ModifiedCountATLastPromotion.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainModifiedInformation2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.DomainModifiedCount == nil {
		o.DomainModifiedCount = &OldLargeInteger{}
	}
	if err := o.DomainModifiedCount.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.CreationTime == nil {
		o.CreationTime = &OldLargeInteger{}
	}
	if err := o.CreationTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ModifiedCountATLastPromotion == nil {
		o.ModifiedCountATLastPromotion = &OldLargeInteger{}
	}
	if err := o.ModifiedCountATLastPromotion.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainGeneralInformation structure represents SAMPR_DOMAIN_GENERAL_INFORMATION RPC structure.
//
// The SAMPR_DOMAIN_GENERAL_INFORMATION structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
//
// Note  In section 2.2.3.1, the types for the DomainServerState and DomainServerRole
// members are the DOMAIN_SERVER_ENABLE_STATE and DOMAIN_SERVER_ROLE enumerations, respectively.
// These fields have the same purpose as the enumeration values, but the data types
// are different. The following tables show the corresponding mappings.
type DomainGeneralInformation struct {
	ForceLogoff           *OldLargeInteger    `idl:"name:ForceLogoff" json:"force_logoff"`
	OEMInformation        *dtyp.UnicodeString `idl:"name:OemInformation" json:"oem_information"`
	DomainName            *dtyp.UnicodeString `idl:"name:DomainName" json:"domain_name"`
	SourceNodeNameReplica *dtyp.UnicodeString `idl:"name:ReplicaSourceNodeName" json:"source_node_name_replica"`
	DomainModifiedCount   *OldLargeInteger    `idl:"name:DomainModifiedCount" json:"domain_modified_count"`
	// For DomainServerState:
	//
	//	+----------------------------------------------+---------------------+
	//	|    ENUMERATION DOMAIN SERVER ENABLE STATE    |    UNSIGNED LONG    |
	//	|                    VALUE                     |        VALUE        |
	//	+----------------------------------------------+---------------------+
	//	+----------------------------------------------+---------------------+
	//	| DomainServerEnabled                          |                   1 |
	//	+----------------------------------------------+---------------------+
	//	| DomainServerDisabled                         |                   2 |
	//	+----------------------------------------------+---------------------+
	DomainServerState uint32 `idl:"name:DomainServerState" json:"domain_server_state"`
	// For DomainServerRole:
	//
	//	+--------------------------------------+---------------------+
	//	|    ENUMERATION DOMAIN SERVER ROLE    |    UNSIGNED LONG    |
	//	|                VALUE                 |        VALUE        |
	//	+--------------------------------------+---------------------+
	//	+--------------------------------------+---------------------+
	//	| DomainServerRoleBackup               |                   2 |
	//	+--------------------------------------+---------------------+
	//	| DomainServerRolePrimary              |                   3 |
	//	+--------------------------------------+---------------------+
	DomainServerRole         uint32 `idl:"name:DomainServerRole" json:"domain_server_role"`
	UASCompatibilityRequired uint8  `idl:"name:UasCompatibilityRequired" json:"uas_compatibility_required"`
	UserCount                uint32 `idl:"name:UserCount" json:"user_count"`
	GroupCount               uint32 `idl:"name:GroupCount" json:"group_count"`
	AliasCount               uint32 `idl:"name:AliasCount" json:"alias_count"`
}

func (o *DomainGeneralInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainGeneralInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ForceLogoff != nil {
		if err := o.ForceLogoff.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OEMInformation != nil {
		if err := o.OEMInformation.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
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
	if o.SourceNodeNameReplica != nil {
		if err := o.SourceNodeNameReplica.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DomainModifiedCount != nil {
		if err := o.DomainModifiedCount.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DomainServerState); err != nil {
		return err
	}
	if err := w.WriteData(o.DomainServerRole); err != nil {
		return err
	}
	if err := w.WriteData(o.UASCompatibilityRequired); err != nil {
		return err
	}
	if err := w.WriteData(o.UserCount); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupCount); err != nil {
		return err
	}
	if err := w.WriteData(o.AliasCount); err != nil {
		return err
	}
	return nil
}
func (o *DomainGeneralInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.ForceLogoff == nil {
		o.ForceLogoff = &OldLargeInteger{}
	}
	if err := o.ForceLogoff.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OEMInformation == nil {
		o.OEMInformation = &dtyp.UnicodeString{}
	}
	if err := o.OEMInformation.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DomainName == nil {
		o.DomainName = &dtyp.UnicodeString{}
	}
	if err := o.DomainName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SourceNodeNameReplica == nil {
		o.SourceNodeNameReplica = &dtyp.UnicodeString{}
	}
	if err := o.SourceNodeNameReplica.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DomainModifiedCount == nil {
		o.DomainModifiedCount = &OldLargeInteger{}
	}
	if err := o.DomainModifiedCount.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainServerState); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainServerRole); err != nil {
		return err
	}
	if err := w.ReadData(&o.UASCompatibilityRequired); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.AliasCount); err != nil {
		return err
	}
	return nil
}

// DomainGeneralInformation2 structure represents SAMPR_DOMAIN_GENERAL_INFORMATION2 RPC structure.
//
// The SAMPR_DOMAIN_GENERAL_INFORMATION2 structure contains domain fields.
//
// For information on each field, see section 2.2.3.1, except for I1, which is specified
// in section 2.2.3.10.
type DomainGeneralInformation2 struct {
	I1                       *DomainGeneralInformation `idl:"name:I1" json:"i1"`
	LockoutDuration          *dtyp.LargeInteger        `idl:"name:LockoutDuration" json:"lockout_duration"`
	LockoutObservationWindow *dtyp.LargeInteger        `idl:"name:LockoutObservationWindow" json:"lockout_observation_window"`
	LockoutThreshold         uint16                    `idl:"name:LockoutThreshold" json:"lockout_threshold"`
}

func (o *DomainGeneralInformation2) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainGeneralInformation2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.I1 != nil {
		if err := o.I1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainGeneralInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LockoutDuration != nil {
		if err := o.LockoutDuration.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LockoutObservationWindow != nil {
		if err := o.LockoutObservationWindow.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LockoutThreshold); err != nil {
		return err
	}
	return nil
}
func (o *DomainGeneralInformation2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.I1 == nil {
		o.I1 = &DomainGeneralInformation{}
	}
	if err := o.I1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LockoutDuration == nil {
		o.LockoutDuration = &dtyp.LargeInteger{}
	}
	if err := o.LockoutDuration.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LockoutObservationWindow == nil {
		o.LockoutObservationWindow = &dtyp.LargeInteger{}
	}
	if err := o.LockoutObservationWindow.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LockoutThreshold); err != nil {
		return err
	}
	return nil
}

// DomainOEMInformation structure represents SAMPR_DOMAIN_OEM_INFORMATION RPC structure.
//
// The SAMPR_DOMAIN_OEM_INFORMATION structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
type DomainOEMInformation struct {
	OEMInformation *dtyp.UnicodeString `idl:"name:OemInformation" json:"oem_information"`
}

func (o *DomainOEMInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainOEMInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.OEMInformation != nil {
		if err := o.OEMInformation.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainOEMInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.OEMInformation == nil {
		o.OEMInformation = &dtyp.UnicodeString{}
	}
	if err := o.OEMInformation.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainNameInformation structure represents SAMPR_DOMAIN_NAME_INFORMATION RPC structure.
//
// The SAMPR_DOMAIN_NAME_INFORMATION structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
type DomainNameInformation struct {
	DomainName *dtyp.UnicodeString `idl:"name:DomainName" json:"domain_name"`
}

func (o *DomainNameInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainNameInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	return nil
}
func (o *DomainNameInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.DomainName == nil {
		o.DomainName = &dtyp.UnicodeString{}
	}
	if err := o.DomainName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainReplicationInformation structure represents SAMPR_DOMAIN_REPLICATION_INFORMATION RPC structure.
//
// The SAMPR_DOMAIN_REPLICATION_INFORMATION structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
type DomainReplicationInformation struct {
	SourceNodeNameReplica *dtyp.UnicodeString `idl:"name:ReplicaSourceNodeName" json:"source_node_name_replica"`
}

func (o *DomainReplicationInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainReplicationInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.SourceNodeNameReplica != nil {
		if err := o.SourceNodeNameReplica.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainReplicationInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.SourceNodeNameReplica == nil {
		o.SourceNodeNameReplica = &dtyp.UnicodeString{}
	}
	if err := o.SourceNodeNameReplica.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainLockoutInformation structure represents SAMPR_DOMAIN_LOCKOUT_INFORMATION RPC structure.
//
// The SAMPR_DOMAIN_LOCKOUT_INFORMATION structure contains domain fields.
//
// For information on each field, see section 2.2.3.1.
type DomainLockoutInformation struct {
	LockoutDuration          *dtyp.LargeInteger `idl:"name:LockoutDuration" json:"lockout_duration"`
	LockoutObservationWindow *dtyp.LargeInteger `idl:"name:LockoutObservationWindow" json:"lockout_observation_window"`
	LockoutThreshold         uint16             `idl:"name:LockoutThreshold" json:"lockout_threshold"`
}

func (o *DomainLockoutInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainLockoutInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.LockoutDuration != nil {
		if err := o.LockoutDuration.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LockoutObservationWindow != nil {
		if err := o.LockoutObservationWindow.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LockoutThreshold); err != nil {
		return err
	}
	return nil
}
func (o *DomainLockoutInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.LockoutDuration == nil {
		o.LockoutDuration = &dtyp.LargeInteger{}
	}
	if err := o.LockoutDuration.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LockoutObservationWindow == nil {
		o.LockoutObservationWindow = &dtyp.LargeInteger{}
	}
	if err := o.LockoutObservationWindow.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LockoutThreshold); err != nil {
		return err
	}
	return nil
}

// DomainInformationClass type represents DOMAIN_INFORMATION_CLASS RPC enumeration.
//
// The DOMAIN_INFORMATION_CLASS enumeration indicates how to interpret the Buffer parameter
// for SamrSetInformationDomain and SamrQueryInformationDomain. For a list of associated
// structures, see section 2.2.3.17.
type DomainInformationClass uint16

var (
	// DomainPasswordInformation:  Indicates the Buffer parameter is to be interpreted
	// as a DOMAIN_PASSWORD_INFORMATION structure (see section 2.2.3.5).
	DomainInformationClassPasswordInformation DomainInformationClass = 1
	// DomainGeneralInformation:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_DOMAIN_GENERAL_INFORMATION structure (see section 2.2.3.10).
	DomainInformationClassGeneralInformation DomainInformationClass = 2
	// DomainLogoffInformation:  Indicates the Buffer parameter is to be interpreted as
	// a DOMAIN_LOGOFF_INFORMATION structure (see section 2.2.3.6).
	DomainInformationClassLogoffInformation DomainInformationClass = 3
	// DomainOemInformation:  Indicates the Buffer parameter is to be interpreted as a
	// SAMPR_DOMAIN_OEM_INFORMATION structure (see section 2.2.3.12).
	DomainInformationClassOEMInformation DomainInformationClass = 4
	// DomainNameInformation:  Indicates the Buffer parameter is to be interpreted as a
	// SAMPR_DOMAIN_NAME_INFORMATION structure (see section 2.2.3.13).
	DomainInformationClassNameInformation DomainInformationClass = 5
	// DomainReplicationInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_DOMAIN_REPLICATION_INFORMATION structure (see section 2.2.3.14).
	DomainInformationClassReplicationInformation DomainInformationClass = 6
	// DomainServerRoleInformation:  Indicates the Buffer parameter is to be interpreted
	// as a DOMAIN_SERVER_ROLE_INFORMATION structure (see section 2.2.3.7).
	DomainInformationClassServerRoleInformation DomainInformationClass = 7
	// DomainModifiedInformation:  Indicates the Buffer parameter is to be interpreted
	// as a DOMAIN_MODIFIED_INFORMATION structure (see section 2.2.3.8).
	DomainInformationClassModifiedInformation DomainInformationClass = 8
	// DomainStateInformation:  Indicates the Buffer parameter is to be interpreted as
	// a DOMAIN_STATE_INFORMATION structure (see section 2.2.3.3).
	DomainInformationClassStateInformation DomainInformationClass = 9
	// DomainGeneralInformation2:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_DOMAIN_GENERAL_INFORMATION2 structure (see section 2.2.3.11).
	DomainInformationClassGeneralInformation2 DomainInformationClass = 11
	// DomainLockoutInformation:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_DOMAIN_LOCKOUT_INFORMATION structure (see section 2.2.3.15).
	DomainInformationClassLockoutInformation DomainInformationClass = 12
	// DomainModifiedInformation2:  Indicates the Buffer parameter is to be interpreted
	// as a DOMAIN_MODIFIED_INFORMATION2 structure (see section 2.2.3.9).
	DomainInformationClassModifiedInformation2 DomainInformationClass = 13
)

func (o DomainInformationClass) String() string {
	switch o {
	case DomainInformationClassPasswordInformation:
		return "DomainInformationClassPasswordInformation"
	case DomainInformationClassGeneralInformation:
		return "DomainInformationClassGeneralInformation"
	case DomainInformationClassLogoffInformation:
		return "DomainInformationClassLogoffInformation"
	case DomainInformationClassOEMInformation:
		return "DomainInformationClassOEMInformation"
	case DomainInformationClassNameInformation:
		return "DomainInformationClassNameInformation"
	case DomainInformationClassReplicationInformation:
		return "DomainInformationClassReplicationInformation"
	case DomainInformationClassServerRoleInformation:
		return "DomainInformationClassServerRoleInformation"
	case DomainInformationClassModifiedInformation:
		return "DomainInformationClassModifiedInformation"
	case DomainInformationClassStateInformation:
		return "DomainInformationClassStateInformation"
	case DomainInformationClassGeneralInformation2:
		return "DomainInformationClassGeneralInformation2"
	case DomainInformationClassLockoutInformation:
		return "DomainInformationClassLockoutInformation"
	case DomainInformationClassModifiedInformation2:
		return "DomainInformationClassModifiedInformation2"
	}
	return "Invalid"
}

// DomainInfoBuffer structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union.
//
// The SAMPR_DOMAIN_INFO_BUFFER union combines all possible structures used in the SamrSetInformationDomain
// and SamrQueryInformationDomain methods. For details on each field, see the associated
// section for each field structure.
type DomainInfoBuffer struct {
	// Types that are assignable to Value
	//
	// *DomainInfoBuffer_Password
	// *DomainInfoBuffer_General
	// *DomainInfoBuffer_Logoff
	// *DomainInfoBuffer_OEM
	// *DomainInfoBuffer_Name
	// *DomainInfoBuffer_Role
	// *DomainInfoBuffer_Replication
	// *DomainInfoBuffer_Modified
	// *DomainInfoBuffer_State
	// *DomainInfoBuffer_General2
	// *DomainInfoBuffer_Lockout
	// *DomainInfoBuffer_Modified2
	Value is_DomainInfoBuffer `json:"value"`
}

func (o *DomainInfoBuffer) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *DomainInfoBuffer_Password:
		if value != nil {
			return value.Password
		}
	case *DomainInfoBuffer_General:
		if value != nil {
			return value.General
		}
	case *DomainInfoBuffer_Logoff:
		if value != nil {
			return value.Logoff
		}
	case *DomainInfoBuffer_OEM:
		if value != nil {
			return value.OEM
		}
	case *DomainInfoBuffer_Name:
		if value != nil {
			return value.Name
		}
	case *DomainInfoBuffer_Role:
		if value != nil {
			return value.Role
		}
	case *DomainInfoBuffer_Replication:
		if value != nil {
			return value.Replication
		}
	case *DomainInfoBuffer_Modified:
		if value != nil {
			return value.Modified
		}
	case *DomainInfoBuffer_State:
		if value != nil {
			return value.State
		}
	case *DomainInfoBuffer_General2:
		if value != nil {
			return value.General2
		}
	case *DomainInfoBuffer_Lockout:
		if value != nil {
			return value.Lockout
		}
	case *DomainInfoBuffer_Modified2:
		if value != nil {
			return value.Modified2
		}
	}
	return nil
}

type is_DomainInfoBuffer interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_DomainInfoBuffer()
}

func (o *DomainInfoBuffer) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *DomainInfoBuffer_Password:
		return uint16(1)
	case *DomainInfoBuffer_General:
		return uint16(2)
	case *DomainInfoBuffer_Logoff:
		return uint16(3)
	case *DomainInfoBuffer_OEM:
		return uint16(4)
	case *DomainInfoBuffer_Name:
		return uint16(5)
	case *DomainInfoBuffer_Role:
		return uint16(7)
	case *DomainInfoBuffer_Replication:
		return uint16(6)
	case *DomainInfoBuffer_Modified:
		return uint16(8)
	case *DomainInfoBuffer_State:
		return uint16(9)
	case *DomainInfoBuffer_General2:
		return uint16(11)
	case *DomainInfoBuffer_Lockout:
		return uint16(12)
	case *DomainInfoBuffer_Modified2:
		return uint16(13)
	}
	return uint16(0)
}

func (o *DomainInfoBuffer) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*DomainInfoBuffer_Password)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_Password{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*DomainInfoBuffer_General)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_General{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*DomainInfoBuffer_Logoff)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_Logoff{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*DomainInfoBuffer_OEM)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_OEM{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5):
		_o, _ := o.Value.(*DomainInfoBuffer_Name)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_Name{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(7):
		_o, _ := o.Value.(*DomainInfoBuffer_Role)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_Role{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(6):
		_o, _ := o.Value.(*DomainInfoBuffer_Replication)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_Replication{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(8):
		_o, _ := o.Value.(*DomainInfoBuffer_Modified)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_Modified{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(9):
		_o, _ := o.Value.(*DomainInfoBuffer_State)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_State{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(11):
		_o, _ := o.Value.(*DomainInfoBuffer_General2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_General2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(12):
		_o, _ := o.Value.(*DomainInfoBuffer_Lockout)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_Lockout{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(13):
		_o, _ := o.Value.(*DomainInfoBuffer_Modified2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer_Modified2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *DomainInfoBuffer) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &DomainInfoBuffer_Password{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &DomainInfoBuffer_General{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &DomainInfoBuffer_Logoff{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &DomainInfoBuffer_OEM{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5):
		o.Value = &DomainInfoBuffer_Name{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(7):
		o.Value = &DomainInfoBuffer_Role{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(6):
		o.Value = &DomainInfoBuffer_Replication{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(8):
		o.Value = &DomainInfoBuffer_Modified{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(9):
		o.Value = &DomainInfoBuffer_State{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(11):
		o.Value = &DomainInfoBuffer_General2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(12):
		o.Value = &DomainInfoBuffer_Lockout{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(13):
		o.Value = &DomainInfoBuffer_Modified2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// DomainInfoBuffer_Password structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 1
type DomainInfoBuffer_Password struct {
	Password *DomainPasswordInformation `idl:"name:Password" json:"password"`
}

func (*DomainInfoBuffer_Password) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_Password) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Password != nil {
		if err := o.Password.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainPasswordInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_Password) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Password == nil {
		o.Password = &DomainPasswordInformation{}
	}
	if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_General structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 2
type DomainInfoBuffer_General struct {
	General *DomainGeneralInformation `idl:"name:General" json:"general"`
}

func (*DomainInfoBuffer_General) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_General) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.General != nil {
		if err := o.General.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainGeneralInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_General) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.General == nil {
		o.General = &DomainGeneralInformation{}
	}
	if err := o.General.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_Logoff structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 3
type DomainInfoBuffer_Logoff struct {
	Logoff *DomainLogoffInformation `idl:"name:Logoff" json:"logoff"`
}

func (*DomainInfoBuffer_Logoff) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_Logoff) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Logoff != nil {
		if err := o.Logoff.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainLogoffInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_Logoff) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Logoff == nil {
		o.Logoff = &DomainLogoffInformation{}
	}
	if err := o.Logoff.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_OEM structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 4
type DomainInfoBuffer_OEM struct {
	OEM *DomainOEMInformation `idl:"name:Oem" json:"oem"`
}

func (*DomainInfoBuffer_OEM) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_OEM) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.OEM != nil {
		if err := o.OEM.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainOEMInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_OEM) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.OEM == nil {
		o.OEM = &DomainOEMInformation{}
	}
	if err := o.OEM.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_Name structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 5
type DomainInfoBuffer_Name struct {
	Name *DomainNameInformation `idl:"name:Name" json:"name"`
}

func (*DomainInfoBuffer_Name) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_Name) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainNameInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_Name) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Name == nil {
		o.Name = &DomainNameInformation{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_Role structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 7
type DomainInfoBuffer_Role struct {
	Role *DomainServerRoleInformation `idl:"name:Role" json:"role"`
}

func (*DomainInfoBuffer_Role) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_Role) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Role != nil {
		if err := o.Role.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainServerRoleInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_Role) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Role == nil {
		o.Role = &DomainServerRoleInformation{}
	}
	if err := o.Role.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_Replication structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 6
type DomainInfoBuffer_Replication struct {
	Replication *DomainReplicationInformation `idl:"name:Replication" json:"replication"`
}

func (*DomainInfoBuffer_Replication) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_Replication) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Replication != nil {
		if err := o.Replication.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainReplicationInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_Replication) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Replication == nil {
		o.Replication = &DomainReplicationInformation{}
	}
	if err := o.Replication.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_Modified structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 8
type DomainInfoBuffer_Modified struct {
	Modified *DomainModifiedInformation `idl:"name:Modified" json:"modified"`
}

func (*DomainInfoBuffer_Modified) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_Modified) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Modified != nil {
		if err := o.Modified.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainModifiedInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_Modified) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Modified == nil {
		o.Modified = &DomainModifiedInformation{}
	}
	if err := o.Modified.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_State structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 9
type DomainInfoBuffer_State struct {
	State *DomainStateInformation `idl:"name:State" json:"state"`
}

func (*DomainInfoBuffer_State) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_State) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.State != nil {
		if err := o.State.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainStateInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_State) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.State == nil {
		o.State = &DomainStateInformation{}
	}
	if err := o.State.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_General2 structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 11
type DomainInfoBuffer_General2 struct {
	General2 *DomainGeneralInformation2 `idl:"name:General2" json:"general2"`
}

func (*DomainInfoBuffer_General2) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_General2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.General2 != nil {
		if err := o.General2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainGeneralInformation2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_General2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.General2 == nil {
		o.General2 = &DomainGeneralInformation2{}
	}
	if err := o.General2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_Lockout structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 12
type DomainInfoBuffer_Lockout struct {
	Lockout *DomainLockoutInformation `idl:"name:Lockout" json:"lockout"`
}

func (*DomainInfoBuffer_Lockout) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_Lockout) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Lockout != nil {
		if err := o.Lockout.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainLockoutInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_Lockout) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Lockout == nil {
		o.Lockout = &DomainLockoutInformation{}
	}
	if err := o.Lockout.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainInfoBuffer_Modified2 structure represents SAMPR_DOMAIN_INFO_BUFFER RPC union arm.
//
// It has following labels: 13
type DomainInfoBuffer_Modified2 struct {
	Modified2 *DomainModifiedInformation2 `idl:"name:Modified2" json:"modified2"`
}

func (*DomainInfoBuffer_Modified2) is_DomainInfoBuffer() {}

func (o *DomainInfoBuffer_Modified2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Modified2 != nil {
		if err := o.Modified2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainModifiedInformation2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainInfoBuffer_Modified2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Modified2 == nil {
		o.Modified2 = &DomainModifiedInformation2{}
	}
	if err := o.Modified2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainDisplayInformation type represents DOMAIN_DISPLAY_INFORMATION RPC enumeration.
//
// The DOMAIN_DISPLAY_INFORMATION enumeration indicates how to interpret the Buffer
// parameter for SamrQueryDisplayInformation, SamrQueryDisplayInformation2, SamrQueryDisplayInformation3,
// SamrGetDisplayEnumerationIndex, and SamrGetDisplayEnumerationIndex2. See section
// 2.2.8.13 for the list of the structures that are associated with each enumeration.
//
// *PDOMAIN_DISPLAY_INFORMATION;
type DomainDisplayInformation uint16

var (
	// DomainDisplayUser: Indicates the Buffer parameter is to be interpreted as a SAMPR_DOMAIN_DISPLAY_USER_BUFFER
	// structure (see section 2.2.8.7).
	DomainDisplayInformationUser DomainDisplayInformation = 1
	// DomainDisplayMachine: Indicates the Buffer parameter is to be interpreted as a SAMPR_DOMAIN_DISPLAY_MACHINE_BUFFER
	// structure (see section 2.2.8.8).
	DomainDisplayInformationMachine DomainDisplayInformation = 2
	// DomainDisplayGroup: Indicates the Buffer parameter is to be interpreted as a SAMPR_DOMAIN_DISPLAY_GROUP_BUFFER
	// structure (see section 2.2.8.9).
	DomainDisplayInformationGroup DomainDisplayInformation = 3
	// DomainDisplayOemUser: Indicates the Buffer parameter is to be interpreted as a SAMPR_DOMAIN_DISPLAY_OEM_USER_BUFFER
	// structure (see section 2.2.8.10).
	DomainDisplayInformationOEMUser DomainDisplayInformation = 4
	// DomainDisplayOemGroup: Indicates the Buffer parameter is to be interpreted as a SAMPR_DOMAIN_DISPLAY_OEM_GROUP_BUFFER
	// structure (see section 2.2.8.11).
	DomainDisplayInformationOEMGroup DomainDisplayInformation = 5
)

func (o DomainDisplayInformation) String() string {
	switch o {
	case DomainDisplayInformationUser:
		return "DomainDisplayInformationUser"
	case DomainDisplayInformationMachine:
		return "DomainDisplayInformationMachine"
	case DomainDisplayInformationGroup:
		return "DomainDisplayInformationGroup"
	case DomainDisplayInformationOEMUser:
		return "DomainDisplayInformationOEMUser"
	case DomainDisplayInformationOEMGroup:
		return "DomainDisplayInformationOEMGroup"
	}
	return "Invalid"
}

// DomainDisplayUser structure represents SAMPR_DOMAIN_DISPLAY_USER RPC structure.
//
// The SAMPR_DOMAIN_DISPLAY_USER structure contains a subset of user account information
// sufficient to show a summary of the account for an account management application.
//
// For information on each field, see section 2.2.8.1.
type DomainDisplayUser struct {
	Index          uint32              `idl:"name:Index" json:"index"`
	RID            uint32              `idl:"name:Rid" json:"rid"`
	AccountControl uint32              `idl:"name:AccountControl" json:"account_control"`
	AccountName    *dtyp.UnicodeString `idl:"name:AccountName" json:"account_name"`
	AdminComment   *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
	FullName       *dtyp.UnicodeString `idl:"name:FullName" json:"full_name"`
}

func (o *DomainDisplayUser) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayUser) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteData(o.RID); err != nil {
		return err
	}
	if err := w.WriteData(o.AccountControl); err != nil {
		return err
	}
	if o.AccountName != nil {
		if err := o.AccountName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FullName != nil {
		if err := o.FullName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayUser) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadData(&o.RID); err != nil {
		return err
	}
	if err := w.ReadData(&o.AccountControl); err != nil {
		return err
	}
	if o.AccountName == nil {
		o.AccountName = &dtyp.UnicodeString{}
	}
	if err := o.AccountName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.FullName == nil {
		o.FullName = &dtyp.UnicodeString{}
	}
	if err := o.FullName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainDisplayMachine structure represents SAMPR_DOMAIN_DISPLAY_MACHINE RPC structure.
//
// The SAMPR_DOMAIN_DISPLAY_MACHINE structure contains a subset of machine account information
// sufficient to show a summary of the account for an account management application.
//
// For information on each field, see section 2.2.8.1.
type DomainDisplayMachine struct {
	Index          uint32              `idl:"name:Index" json:"index"`
	RID            uint32              `idl:"name:Rid" json:"rid"`
	AccountControl uint32              `idl:"name:AccountControl" json:"account_control"`
	AccountName    *dtyp.UnicodeString `idl:"name:AccountName" json:"account_name"`
	AdminComment   *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
}

func (o *DomainDisplayMachine) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayMachine) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteData(o.RID); err != nil {
		return err
	}
	if err := w.WriteData(o.AccountControl); err != nil {
		return err
	}
	if o.AccountName != nil {
		if err := o.AccountName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayMachine) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadData(&o.RID); err != nil {
		return err
	}
	if err := w.ReadData(&o.AccountControl); err != nil {
		return err
	}
	if o.AccountName == nil {
		o.AccountName = &dtyp.UnicodeString{}
	}
	if err := o.AccountName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainDisplayGroup structure represents SAMPR_DOMAIN_DISPLAY_GROUP RPC structure.
//
// The SAMPR_DOMAIN_DISPLAY_GROUP structure contains a subset of group information sufficient
// to show a summary of the account for an account management application.
//
// For information on each field, see section 2.2.8.1.
type DomainDisplayGroup struct {
	Index        uint32              `idl:"name:Index" json:"index"`
	RID          uint32              `idl:"name:Rid" json:"rid"`
	Attributes   uint32              `idl:"name:Attributes" json:"attributes"`
	AccountName  *dtyp.UnicodeString `idl:"name:AccountName" json:"account_name"`
	AdminComment *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
}

func (o *DomainDisplayGroup) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayGroup) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteData(o.RID); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.AccountName != nil {
		if err := o.AccountName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayGroup) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadData(&o.RID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.AccountName == nil {
		o.AccountName = &dtyp.UnicodeString{}
	}
	if err := o.AccountName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainDisplayOEMUser structure represents SAMPR_DOMAIN_DISPLAY_OEM_USER RPC structure.
//
// The SAMPR_DOMAIN_DISPLAY_OEM_USER structure contains a subset of user account information
// sufficient to show a summary of the account for an account management application.
// This structure exists to support non–Unicode-based systems.
//
// For information on each field, see section 2.2.8.1.
type DomainDisplayOEMUser struct {
	Index          uint32  `idl:"name:Index" json:"index"`
	OEMAccountName *String `idl:"name:OemAccountName" json:"oem_account_name"`
}

func (o *DomainDisplayOEMUser) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayOEMUser) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if o.OEMAccountName != nil {
		if err := o.OEMAccountName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&String{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayOEMUser) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if o.OEMAccountName == nil {
		o.OEMAccountName = &String{}
	}
	if err := o.OEMAccountName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainDisplayOEMGroup structure represents SAMPR_DOMAIN_DISPLAY_OEM_GROUP RPC structure.
//
// The SAMPR_DOMAIN_DISPLAY_OEM_GROUP structure contains a subset of group information
// sufficient to show a summary of the account for an account management application.
// This structure exists to support non–Unicode-based systems.
//
// For information on each field, see section 2.2.8.1.
type DomainDisplayOEMGroup struct {
	Index          uint32  `idl:"name:Index" json:"index"`
	OEMAccountName *String `idl:"name:OemAccountName" json:"oem_account_name"`
}

func (o *DomainDisplayOEMGroup) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayOEMGroup) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if o.OEMAccountName != nil {
		if err := o.OEMAccountName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&String{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayOEMGroup) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if o.OEMAccountName == nil {
		o.OEMAccountName = &String{}
	}
	if err := o.OEMAccountName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DomainDisplayUserBuffer structure represents SAMPR_DOMAIN_DISPLAY_USER_BUFFER RPC structure.
//
// The SAMPR_DOMAIN_DISPLAY_USER_BUFFER structure holds an array of SAMPR_DOMAIN_DISPLAY_USER
// elements used to return a list of users through the SamrQueryDisplayInformation family
// of methods (see section 3.1.5.3).
type DomainDisplayUserBuffer struct {
	// EntriesRead:  The number of elements in Buffer. If zero, Buffer MUST be ignored.
	// If nonzero, Buffer MUST point to at least EntriesRead number of elements.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  An array of SAMPR_DOMAIN_DISPLAY_USER elements.
	Buffer []*DomainDisplayUser `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *DomainDisplayUserBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayUserBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
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
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&DomainDisplayUser{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&DomainDisplayUser{}).MarshalNDR(ctx, w); err != nil {
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
func (o *DomainDisplayUserBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
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
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*DomainDisplayUser, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &DomainDisplayUser{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*DomainDisplayUser) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// DomainDisplayMachineBuffer structure represents SAMPR_DOMAIN_DISPLAY_MACHINE_BUFFER RPC structure.
//
// The SAMPR_DOMAIN_DISPLAY_MACHINE_BUFFER structure holds an array of SAMPR_DOMAIN_DISPLAY_MACHINE
// elements used to return a list of machine accounts through the SamrQueryDisplayInformation
// family of methods (see section 3.1.5.3).
type DomainDisplayMachineBuffer struct {
	// EntriesRead:  The number of elements in Buffer. If zero, Buffer MUST be ignored.
	// If nonzero, Buffer MUST point to at least EntriesRead number of elements.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  An array of SAMPR_DOMAIN_DISPLAY_MACHINE elements.
	Buffer []*DomainDisplayMachine `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *DomainDisplayMachineBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayMachineBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
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
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&DomainDisplayMachine{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&DomainDisplayMachine{}).MarshalNDR(ctx, w); err != nil {
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
func (o *DomainDisplayMachineBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
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
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*DomainDisplayMachine, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &DomainDisplayMachine{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*DomainDisplayMachine) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// DomainDisplayGroupBuffer structure represents SAMPR_DOMAIN_DISPLAY_GROUP_BUFFER RPC structure.
//
// The SAMPR_DOMAIN_DISPLAY_GROUP_BUFFER structure holds an array of SAMPR_DOMAIN_DISPLAY_GROUP
// elements used to return a list of groups through the SamrQueryDisplayInformation
// family of methods (see section 3.1.5.3).
type DomainDisplayGroupBuffer struct {
	// EntriesRead:  The number of elements in Buffer. If zero, Buffer MUST be ignored.
	// If nonzero, Buffer MUST point to at least EntriesRead number of elements.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  An array of SAMPR_DOMAIN_DISPLAY_GROUP elements.
	Buffer []*DomainDisplayGroup `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *DomainDisplayGroupBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayGroupBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
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
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&DomainDisplayGroup{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&DomainDisplayGroup{}).MarshalNDR(ctx, w); err != nil {
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
func (o *DomainDisplayGroupBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
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
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*DomainDisplayGroup, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &DomainDisplayGroup{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*DomainDisplayGroup) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// DomainDisplayOEMUserBuffer structure represents SAMPR_DOMAIN_DISPLAY_OEM_USER_BUFFER RPC structure.
//
// The SAMPR_DOMAIN_DISPLAY_OEM_USER_BUFFER structure holds an array of SAMPR_DOMAIN_DISPLAY_OEM_USER
// elements used to return a list of users through the SamrQueryDisplayInformation family
// of methods (see section 3.1.5.3).
type DomainDisplayOEMUserBuffer struct {
	// EntriesRead:  The number of elements in Buffer. If zero, Buffer MUST be ignored.
	// If nonzero, Buffer MUST point to at least EntriesRead number of elements.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  An array of SAMPR_DOMAIN_DISPLAY_OEM_USER elements.
	Buffer []*DomainDisplayOEMUser `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *DomainDisplayOEMUserBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayOEMUserBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
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
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&DomainDisplayOEMUser{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&DomainDisplayOEMUser{}).MarshalNDR(ctx, w); err != nil {
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
func (o *DomainDisplayOEMUserBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
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
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*DomainDisplayOEMUser, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &DomainDisplayOEMUser{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*DomainDisplayOEMUser) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// DomainDisplayOEMGroupBuffer structure represents SAMPR_DOMAIN_DISPLAY_OEM_GROUP_BUFFER RPC structure.
//
// The SAMPR_DOMAIN_DISPLAY_OEM_GROUP_BUFFER structure holds an array of SAMPR_DOMAIN_DISPLAY_OEM_GROUP
// elements used to return a list of user accounts through the SamrQueryDisplayInformation
// family of methods (see section 3.1.5.3).
type DomainDisplayOEMGroupBuffer struct {
	// EntriesRead:  The number of elements in Buffer. If zero, Buffer MUST be ignored.
	// If nonzero, Buffer MUST point to at least EntriesRead number of elements.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  An array of SAMPR_DOMAIN_DISPLAY_OEM_GROUP elements.
	Buffer []*DomainDisplayOEMGroup `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *DomainDisplayOEMGroupBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainDisplayOEMGroupBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
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
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&DomainDisplayOEMGroup{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&DomainDisplayOEMGroup{}).MarshalNDR(ctx, w); err != nil {
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
func (o *DomainDisplayOEMGroupBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
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
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*DomainDisplayOEMGroup, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &DomainDisplayOEMGroup{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*DomainDisplayOEMGroup) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// DisplayInfoBuffer structure represents SAMPR_DISPLAY_INFO_BUFFER RPC union.
//
// The SAMPR_DISPLAY_INFO_BUFFER union is a union of display structures returned by
// the SamrQueryDisplayInformation family of methods (see section 3.1.5.3). For details
// on each field, see the associated section for the field structure.
type DisplayInfoBuffer struct {
	// Types that are assignable to Value
	//
	// *DisplayInfoBuffer_UserInformation
	// *DisplayInfoBuffer_MachineInformation
	// *DisplayInfoBuffer_GroupInformation
	// *DisplayInfoBuffer_OEMUserInformation
	// *DisplayInfoBuffer_OEMGroupInformation
	Value is_DisplayInfoBuffer `json:"value"`
}

func (o *DisplayInfoBuffer) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *DisplayInfoBuffer_UserInformation:
		if value != nil {
			return value.UserInformation
		}
	case *DisplayInfoBuffer_MachineInformation:
		if value != nil {
			return value.MachineInformation
		}
	case *DisplayInfoBuffer_GroupInformation:
		if value != nil {
			return value.GroupInformation
		}
	case *DisplayInfoBuffer_OEMUserInformation:
		if value != nil {
			return value.OEMUserInformation
		}
	case *DisplayInfoBuffer_OEMGroupInformation:
		if value != nil {
			return value.OEMGroupInformation
		}
	}
	return nil
}

type is_DisplayInfoBuffer interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_DisplayInfoBuffer()
}

func (o *DisplayInfoBuffer) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *DisplayInfoBuffer_UserInformation:
		return uint16(1)
	case *DisplayInfoBuffer_MachineInformation:
		return uint16(2)
	case *DisplayInfoBuffer_GroupInformation:
		return uint16(3)
	case *DisplayInfoBuffer_OEMUserInformation:
		return uint16(4)
	case *DisplayInfoBuffer_OEMGroupInformation:
		return uint16(5)
	}
	return uint16(0)
}

func (o *DisplayInfoBuffer) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*DisplayInfoBuffer_UserInformation)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DisplayInfoBuffer_UserInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*DisplayInfoBuffer_MachineInformation)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DisplayInfoBuffer_MachineInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*DisplayInfoBuffer_GroupInformation)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DisplayInfoBuffer_GroupInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*DisplayInfoBuffer_OEMUserInformation)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DisplayInfoBuffer_OEMUserInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5):
		_o, _ := o.Value.(*DisplayInfoBuffer_OEMGroupInformation)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&DisplayInfoBuffer_OEMGroupInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *DisplayInfoBuffer) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &DisplayInfoBuffer_UserInformation{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &DisplayInfoBuffer_MachineInformation{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &DisplayInfoBuffer_GroupInformation{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &DisplayInfoBuffer_OEMUserInformation{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5):
		o.Value = &DisplayInfoBuffer_OEMGroupInformation{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// DisplayInfoBuffer_UserInformation structure represents SAMPR_DISPLAY_INFO_BUFFER RPC union arm.
//
// It has following labels: 1
type DisplayInfoBuffer_UserInformation struct {
	UserInformation *DomainDisplayUserBuffer `idl:"name:UserInformation" json:"user_information"`
}

func (*DisplayInfoBuffer_UserInformation) is_DisplayInfoBuffer() {}

func (o *DisplayInfoBuffer_UserInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UserInformation != nil {
		if err := o.UserInformation.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainDisplayUserBuffer{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DisplayInfoBuffer_UserInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UserInformation == nil {
		o.UserInformation = &DomainDisplayUserBuffer{}
	}
	if err := o.UserInformation.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DisplayInfoBuffer_MachineInformation structure represents SAMPR_DISPLAY_INFO_BUFFER RPC union arm.
//
// It has following labels: 2
type DisplayInfoBuffer_MachineInformation struct {
	MachineInformation *DomainDisplayMachineBuffer `idl:"name:MachineInformation" json:"machine_information"`
}

func (*DisplayInfoBuffer_MachineInformation) is_DisplayInfoBuffer() {}

func (o *DisplayInfoBuffer_MachineInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MachineInformation != nil {
		if err := o.MachineInformation.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainDisplayMachineBuffer{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DisplayInfoBuffer_MachineInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MachineInformation == nil {
		o.MachineInformation = &DomainDisplayMachineBuffer{}
	}
	if err := o.MachineInformation.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DisplayInfoBuffer_GroupInformation structure represents SAMPR_DISPLAY_INFO_BUFFER RPC union arm.
//
// It has following labels: 3
type DisplayInfoBuffer_GroupInformation struct {
	GroupInformation *DomainDisplayGroupBuffer `idl:"name:GroupInformation" json:"group_information"`
}

func (*DisplayInfoBuffer_GroupInformation) is_DisplayInfoBuffer() {}

func (o *DisplayInfoBuffer_GroupInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GroupInformation != nil {
		if err := o.GroupInformation.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainDisplayGroupBuffer{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DisplayInfoBuffer_GroupInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.GroupInformation == nil {
		o.GroupInformation = &DomainDisplayGroupBuffer{}
	}
	if err := o.GroupInformation.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DisplayInfoBuffer_OEMUserInformation structure represents SAMPR_DISPLAY_INFO_BUFFER RPC union arm.
//
// It has following labels: 4
type DisplayInfoBuffer_OEMUserInformation struct {
	OEMUserInformation *DomainDisplayOEMUserBuffer `idl:"name:OemUserInformation" json:"oem_user_information"`
}

func (*DisplayInfoBuffer_OEMUserInformation) is_DisplayInfoBuffer() {}

func (o *DisplayInfoBuffer_OEMUserInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.OEMUserInformation != nil {
		if err := o.OEMUserInformation.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainDisplayOEMUserBuffer{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DisplayInfoBuffer_OEMUserInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.OEMUserInformation == nil {
		o.OEMUserInformation = &DomainDisplayOEMUserBuffer{}
	}
	if err := o.OEMUserInformation.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DisplayInfoBuffer_OEMGroupInformation structure represents SAMPR_DISPLAY_INFO_BUFFER RPC union arm.
//
// It has following labels: 5
type DisplayInfoBuffer_OEMGroupInformation struct {
	OEMGroupInformation *DomainDisplayOEMGroupBuffer `idl:"name:OemGroupInformation" json:"oem_group_information"`
}

func (*DisplayInfoBuffer_OEMGroupInformation) is_DisplayInfoBuffer() {}

func (o *DisplayInfoBuffer_OEMGroupInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.OEMGroupInformation != nil {
		if err := o.OEMGroupInformation.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DomainDisplayOEMGroupBuffer{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DisplayInfoBuffer_OEMGroupInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.OEMGroupInformation == nil {
		o.OEMGroupInformation = &DomainDisplayOEMGroupBuffer{}
	}
	if err := o.OEMGroupInformation.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// GroupAttributeInformation structure represents GROUP_ATTRIBUTE_INFORMATION RPC structure.
//
// The GROUP_ATTRIBUTE_INFORMATION structure contains group fields.
//
// For information on each field, see section 2.2.4.1.
type GroupAttributeInformation struct {
	Attributes uint32 `idl:"name:Attributes" json:"attributes"`
}

func (o *GroupAttributeInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupAttributeInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	return nil
}
func (o *GroupAttributeInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	return nil
}

// GroupGeneralInformation structure represents SAMPR_GROUP_GENERAL_INFORMATION RPC structure.
//
// The SAMPR_GROUP_GENERAL_INFORMATION structure contains group fields.
//
// For information on each field, see section 2.2.4.1.
type GroupGeneralInformation struct {
	Name         *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	Attributes   uint32              `idl:"name:Attributes" json:"attributes"`
	MemberCount  uint32              `idl:"name:MemberCount" json:"member_count"`
	AdminComment *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
}

func (o *GroupGeneralInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupGeneralInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if err := w.WriteData(o.MemberCount); err != nil {
		return err
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupGeneralInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if err := w.ReadData(&o.MemberCount); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// GroupNameInformation structure represents SAMPR_GROUP_NAME_INFORMATION RPC structure.
//
// The SAMPR_GROUP_NAME_INFORMATION structure contains group fields.
//
// For information on each field, see section 2.2.4.1.
type GroupNameInformation struct {
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
}

func (o *GroupNameInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupNameInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *GroupNameInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// GroupAdmCommentInformation structure represents SAMPR_GROUP_ADM_COMMENT_INFORMATION RPC structure.
//
// The SAMPR_GROUP_ADM_COMMENT_INFORMATION structure contains group fields.
//
// For information on each field, see section 2.2.4.1.
type GroupAdmCommentInformation struct {
	AdminComment *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
}

func (o *GroupAdmCommentInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupAdmCommentInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupAdmCommentInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// GroupInformationClass type represents GROUP_INFORMATION_CLASS RPC enumeration.
//
// The GROUP_INFORMATION_CLASS enumeration indicates how to interpret the Buffer parameter
// for SamrSetInformationGroup and SamrQueryInformationGroup. For a list of associated
// structures, see section 2.2.4.7.
type GroupInformationClass uint16

var (
	// GroupGeneralInformation:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_GROUP_GENERAL_INFORMATION structure (see section 2.2.4.3).
	GroupInformationClassGeneralInformation GroupInformationClass = 1
	// GroupNameInformation:  Indicates the Buffer parameter is to be interpreted as a
	// SAMPR_GROUP_NAME_INFORMATION structure (see section 2.2.4.4).
	GroupInformationClassNameInformation GroupInformationClass = 2
	// GroupAttributeInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_GROUP_ATTRIBUTE_INFORMATION structure (see section 2.2.4.2).
	GroupInformationClassAttributeInformation GroupInformationClass = 3
	// GroupAdminCommentInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_GROUP_ADM_COMMENT_INFORMATION structure (see section 2.2.4.5).
	GroupInformationClassAdminCommentInformation GroupInformationClass = 4
	// GroupReplicationInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_GROUP_GENERAL_INFORMATION structure (see section 2.2.4.3).
	GroupInformationClassReplicationInformation GroupInformationClass = 5
)

func (o GroupInformationClass) String() string {
	switch o {
	case GroupInformationClassGeneralInformation:
		return "GroupInformationClassGeneralInformation"
	case GroupInformationClassNameInformation:
		return "GroupInformationClassNameInformation"
	case GroupInformationClassAttributeInformation:
		return "GroupInformationClassAttributeInformation"
	case GroupInformationClassAdminCommentInformation:
		return "GroupInformationClassAdminCommentInformation"
	case GroupInformationClassReplicationInformation:
		return "GroupInformationClassReplicationInformation"
	}
	return "Invalid"
}

// GroupInfoBuffer structure represents SAMPR_GROUP_INFO_BUFFER RPC union.
//
// The SAMPR_GROUP_INFO_BUFFER union combines all possible structures used in the SamrSetInformationGroup
// and SamrQueryInformationGroup methods. For information on each field, with the exception
// of the DoNotUse field, see the associated section for the field structure.
type GroupInfoBuffer struct {
	// Types that are assignable to Value
	//
	// *GroupInfoBuffer_General
	// *GroupInfoBuffer_Name
	// *GroupInfoBuffer_Attribute
	// *GroupInfoBuffer_AdminComment
	// *GroupInfoBuffer_DoNotUse
	Value is_GroupInfoBuffer `json:"value"`
}

func (o *GroupInfoBuffer) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *GroupInfoBuffer_General:
		if value != nil {
			return value.General
		}
	case *GroupInfoBuffer_Name:
		if value != nil {
			return value.Name
		}
	case *GroupInfoBuffer_Attribute:
		if value != nil {
			return value.Attribute
		}
	case *GroupInfoBuffer_AdminComment:
		if value != nil {
			return value.AdminComment
		}
	case *GroupInfoBuffer_DoNotUse:
		if value != nil {
			return value.DoNotUse
		}
	}
	return nil
}

type is_GroupInfoBuffer interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_GroupInfoBuffer()
}

func (o *GroupInfoBuffer) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *GroupInfoBuffer_General:
		return uint16(1)
	case *GroupInfoBuffer_Name:
		return uint16(2)
	case *GroupInfoBuffer_Attribute:
		return uint16(3)
	case *GroupInfoBuffer_AdminComment:
		return uint16(4)
	case *GroupInfoBuffer_DoNotUse:
		return uint16(5)
	}
	return uint16(0)
}

func (o *GroupInfoBuffer) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*GroupInfoBuffer_General)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&GroupInfoBuffer_General{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*GroupInfoBuffer_Name)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&GroupInfoBuffer_Name{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*GroupInfoBuffer_Attribute)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&GroupInfoBuffer_Attribute{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*GroupInfoBuffer_AdminComment)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&GroupInfoBuffer_AdminComment{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5):
		_o, _ := o.Value.(*GroupInfoBuffer_DoNotUse)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&GroupInfoBuffer_DoNotUse{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *GroupInfoBuffer) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &GroupInfoBuffer_General{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &GroupInfoBuffer_Name{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &GroupInfoBuffer_Attribute{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &GroupInfoBuffer_AdminComment{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5):
		o.Value = &GroupInfoBuffer_DoNotUse{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// GroupInfoBuffer_General structure represents SAMPR_GROUP_INFO_BUFFER RPC union arm.
//
// It has following labels: 1
type GroupInfoBuffer_General struct {
	General *GroupGeneralInformation `idl:"name:General" json:"general"`
}

func (*GroupInfoBuffer_General) is_GroupInfoBuffer() {}

func (o *GroupInfoBuffer_General) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.General != nil {
		if err := o.General.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GroupGeneralInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupInfoBuffer_General) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.General == nil {
		o.General = &GroupGeneralInformation{}
	}
	if err := o.General.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// GroupInfoBuffer_Name structure represents SAMPR_GROUP_INFO_BUFFER RPC union arm.
//
// It has following labels: 2
type GroupInfoBuffer_Name struct {
	Name *GroupNameInformation `idl:"name:Name" json:"name"`
}

func (*GroupInfoBuffer_Name) is_GroupInfoBuffer() {}

func (o *GroupInfoBuffer_Name) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GroupNameInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupInfoBuffer_Name) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Name == nil {
		o.Name = &GroupNameInformation{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// GroupInfoBuffer_Attribute structure represents SAMPR_GROUP_INFO_BUFFER RPC union arm.
//
// It has following labels: 3
type GroupInfoBuffer_Attribute struct {
	Attribute *GroupAttributeInformation `idl:"name:Attribute" json:"attribute"`
}

func (*GroupInfoBuffer_Attribute) is_GroupInfoBuffer() {}

func (o *GroupInfoBuffer_Attribute) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Attribute != nil {
		if err := o.Attribute.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GroupAttributeInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupInfoBuffer_Attribute) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Attribute == nil {
		o.Attribute = &GroupAttributeInformation{}
	}
	if err := o.Attribute.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// GroupInfoBuffer_AdminComment structure represents SAMPR_GROUP_INFO_BUFFER RPC union arm.
//
// It has following labels: 4
type GroupInfoBuffer_AdminComment struct {
	AdminComment *GroupAdmCommentInformation `idl:"name:AdminComment" json:"admin_comment"`
}

func (*GroupInfoBuffer_AdminComment) is_GroupInfoBuffer() {}

func (o *GroupInfoBuffer_AdminComment) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GroupAdmCommentInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupInfoBuffer_AdminComment) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.AdminComment == nil {
		o.AdminComment = &GroupAdmCommentInformation{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// GroupInfoBuffer_DoNotUse structure represents SAMPR_GROUP_INFO_BUFFER RPC union arm.
//
// It has following labels: 5
type GroupInfoBuffer_DoNotUse struct {
	// DoNotUse:  This field exists to allow the GroupReplicationInformation enumeration
	// to be specified by the client.
	DoNotUse *GroupGeneralInformation `idl:"name:DoNotUse" json:"do_not_use"`
}

func (*GroupInfoBuffer_DoNotUse) is_GroupInfoBuffer() {}

func (o *GroupInfoBuffer_DoNotUse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DoNotUse != nil {
		if err := o.DoNotUse.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GroupGeneralInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupInfoBuffer_DoNotUse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DoNotUse == nil {
		o.DoNotUse = &GroupGeneralInformation{}
	}
	if err := o.DoNotUse.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AliasGeneralInformation structure represents SAMPR_ALIAS_GENERAL_INFORMATION RPC structure.
//
// The SAMPR_ALIAS_GENERAL_INFORMATION structure contains alias fields.
//
// For information on each field, see section 2.2.5.1.
type AliasGeneralInformation struct {
	Name         *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	MemberCount  uint32              `idl:"name:MemberCount" json:"member_count"`
	AdminComment *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
}

func (o *AliasGeneralInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AliasGeneralInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.MemberCount); err != nil {
		return err
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AliasGeneralInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.MemberCount); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AliasNameInformation structure represents SAMPR_ALIAS_NAME_INFORMATION RPC structure.
//
// The SAMPR_ALIAS_NAME_INFORMATION structure contains alias fields.
//
// For information on each field, see section 2.2.5.1.
type AliasNameInformation struct {
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
}

func (o *AliasNameInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AliasNameInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *AliasNameInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// AliasAdmCommentInformation structure represents SAMPR_ALIAS_ADM_COMMENT_INFORMATION RPC structure.
//
// The SAMPR_ALIAS_ADM_COMMENT_INFORMATION structure contains alias fields.
//
// For information on each field, see section 2.2.5.1.
type AliasAdmCommentInformation struct {
	AdminComment *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
}

func (o *AliasAdmCommentInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AliasAdmCommentInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AliasAdmCommentInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AliasInformationClass type represents ALIAS_INFORMATION_CLASS RPC enumeration.
//
// The ALIAS_INFORMATION_CLASS enumeration indicates how to interpret the Buffer parameter
// for SamrQueryInformationAlias and SamrSetInformationAlias. For a list of the structures
// associated with each enumeration, see section 2.2.5.6.
type AliasInformationClass uint16

var (
	// AliasGeneralInformation:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_ALIAS_GENERAL_INFORMATION structure (see section 2.2.5.2).
	AliasInformationClassGeneralInformation AliasInformationClass = 1
	// AliasNameInformation:  Indicates the Buffer parameter is to be interpreted as a
	// SAMPR_ALIAS_NAME_INFORMATION structure (see section 2.2.5.3).
	AliasInformationClassNameInformation AliasInformationClass = 2
	// AliasAdminCommentInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_ALIAS_ADM_COMMENT_INFORMATION structure (see section 2.2.5.4).
	AliasInformationClassAdminCommentInformation AliasInformationClass = 3
)

func (o AliasInformationClass) String() string {
	switch o {
	case AliasInformationClassGeneralInformation:
		return "AliasInformationClassGeneralInformation"
	case AliasInformationClassNameInformation:
		return "AliasInformationClassNameInformation"
	case AliasInformationClassAdminCommentInformation:
		return "AliasInformationClassAdminCommentInformation"
	}
	return "Invalid"
}

// AliasInfoBuffer structure represents SAMPR_ALIAS_INFO_BUFFER RPC union.
//
// The SAMPR_ALIAS_INFO_BUFFER union combines all possible structures used in the SamrSetInformationAlias
// and SamrQueryInformationAlias methods. For information on each field, see the associated
// section for the field structure.
type AliasInfoBuffer struct {
	// Types that are assignable to Value
	//
	// *AliasInfoBuffer_General
	// *AliasInfoBuffer_Name
	// *AliasInfoBuffer_AdminComment
	Value is_AliasInfoBuffer `json:"value"`
}

func (o *AliasInfoBuffer) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *AliasInfoBuffer_General:
		if value != nil {
			return value.General
		}
	case *AliasInfoBuffer_Name:
		if value != nil {
			return value.Name
		}
	case *AliasInfoBuffer_AdminComment:
		if value != nil {
			return value.AdminComment
		}
	}
	return nil
}

type is_AliasInfoBuffer interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_AliasInfoBuffer()
}

func (o *AliasInfoBuffer) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *AliasInfoBuffer_General:
		return uint16(1)
	case *AliasInfoBuffer_Name:
		return uint16(2)
	case *AliasInfoBuffer_AdminComment:
		return uint16(3)
	}
	return uint16(0)
}

func (o *AliasInfoBuffer) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*AliasInfoBuffer_General)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AliasInfoBuffer_General{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*AliasInfoBuffer_Name)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AliasInfoBuffer_Name{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*AliasInfoBuffer_AdminComment)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AliasInfoBuffer_AdminComment{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *AliasInfoBuffer) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &AliasInfoBuffer_General{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &AliasInfoBuffer_Name{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &AliasInfoBuffer_AdminComment{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// AliasInfoBuffer_General structure represents SAMPR_ALIAS_INFO_BUFFER RPC union arm.
//
// It has following labels: 1
type AliasInfoBuffer_General struct {
	General *AliasGeneralInformation `idl:"name:General" json:"general"`
}

func (*AliasInfoBuffer_General) is_AliasInfoBuffer() {}

func (o *AliasInfoBuffer_General) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.General != nil {
		if err := o.General.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AliasGeneralInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AliasInfoBuffer_General) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.General == nil {
		o.General = &AliasGeneralInformation{}
	}
	if err := o.General.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AliasInfoBuffer_Name structure represents SAMPR_ALIAS_INFO_BUFFER RPC union arm.
//
// It has following labels: 2
type AliasInfoBuffer_Name struct {
	Name *AliasNameInformation `idl:"name:Name" json:"name"`
}

func (*AliasInfoBuffer_Name) is_AliasInfoBuffer() {}

func (o *AliasInfoBuffer_Name) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AliasNameInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AliasInfoBuffer_Name) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Name == nil {
		o.Name = &AliasNameInformation{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AliasInfoBuffer_AdminComment structure represents SAMPR_ALIAS_INFO_BUFFER RPC union arm.
//
// It has following labels: 3
type AliasInfoBuffer_AdminComment struct {
	AdminComment *AliasAdmCommentInformation `idl:"name:AdminComment" json:"admin_comment"`
}

func (*AliasInfoBuffer_AdminComment) is_AliasInfoBuffer() {}

func (o *AliasInfoBuffer_AdminComment) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AliasAdmCommentInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AliasInfoBuffer_AdminComment) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.AdminComment == nil {
		o.AdminComment = &AliasAdmCommentInformation{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// EncryptedUserPassword structure represents SAMPR_ENCRYPTED_USER_PASSWORD RPC structure.
//
// The SAMPR_ENCRYPTED_USER_PASSWORD structure carries an encrypted string.
type EncryptedUserPassword struct {
	// Buffer:  An array to carry encrypted cleartext password data. The encryption key
	// is method-specific, while the algorithm specified in section 3.2.2.1 is common for
	// all methods that use this structure. See the message syntax for SamrOemChangePasswordUser2
	// (section 3.1.5.10.2) and SamrUnicodeChangePasswordUser2 (section 3.1.5.10.3), and
	// the message processing for SamrSetInformationUser2 (section 3.1.5.6.4), for details
	// on the encryption key selection. The size of (256 * 2) + 4 for Buffer is determined
	// by the size of the structure that is encrypted, SAMPR_USER_PASSWORD; see below for
	// more details.
	//
	// For all protocol uses, the decrypted format of Buffer is the following structure.
	//
	// typedef struct _SAMPR_USER_PASSWORD { wchar_t       Buffer[256]; unsigned long
	// Length; } SAMPR_USER_PASSWORD, *PSAMPR_USER_PASSWORD;
	//
	// Buffer: This array contains the cleartext value at the end of the buffer. The start
	// of the string is Length number of bytes from the end of the buffer. The cleartext
	// value can be no more than 512 bytes. The unused portions of SAMPR_USER_PASSWORD.Buffer
	// SHOULD be filled with random bytes by the client. The value 512 is chosen because
	// that is the longest password allowed by this protocol (and enforced by the server).
	//
	// Implementations of this protocol MUST protect the SAMPR_ENCRYPTED_USER_PASSWORD structure
	// by encrypting the 516 bytes of data referenced in its Buffer field on request (and
	// reply), and decrypting on receipt. See section 3.2.2.1 for the specification of the
	// algorithm performing encryption and decryption.
	Buffer []byte `idl:"name:Buffer" json:"buffer"`
}

func (o *EncryptedUserPassword) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedUserPassword) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 516 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 516; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedUserPassword) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Buffer = make([]byte, 516)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	return nil
}

// EncryptedUserPasswordNew structure represents SAMPR_ENCRYPTED_USER_PASSWORD_NEW RPC structure.
//
// The SAMPR_ENCRYPTED_USER_PASSWORD_NEW structure carries an encrypted string.
type EncryptedUserPasswordNew struct {
	// Buffer:  An array to carry encrypted cleartext password data.
	//
	// For all protocol uses, the decrypted format of Buffer is the following structure.
	//
	// typedef struct _SAMPR_USER_PASSWORD_NEW { WCHAR Buffer[256]; ULONG Length; UCHAR
	// ClearSalt[16]; } SAMPR_USER_PASSWORD_NEW, *PSAMPR_USER_PASSWORD_NEW;
	//
	// Buffer: This array contains the cleartext value at the end of the buffer. The cleartext
	// value can be no more than 512 bytes. The start of the string is Length number of
	// bytes from the end of the buffer. The unused portions of SAMPR_USER_PASSWORD_NEW.Buffer
	// SHOULD be filled with random bytes by the client.
	//
	// Length: An unsigned integer, in little-endian byte order, that indicates the number
	// of bytes of the cleartext value (located in SAMPR_USER_PASSWORD_NEW.Buffer).
	//
	// Implementations of this protocol MUST protect the SAMPR_ENCRYPTED_USER_PASSWORD_NEW
	// structure by encrypting the first 516 bytes of data referenced in its Buffer field
	// on request (and reply) and by decrypting on receipt. See section 3.2.2.1 for the
	// specification of the algorithm performing encryption and decryption.
	//
	// The first 516 bytes are defined as the first 516 bytes of the SAMPR_USER_PASSWORD_NEW
	// structure defined previously. The last 16 bytes of the SAMPR_ENCRYPTED_USER_PASSWORD_NEW
	// structure are defined as the last 16 bytes of the SAMPR_USER_PASSWORD_NEW structure
	// and MUST NOT be encrypted or decrypted.
	Buffer []byte `idl:"name:Buffer" json:"buffer"`
}

func (o *EncryptedUserPasswordNew) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedUserPasswordNew) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 532 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 532; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedUserPasswordNew) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Buffer = make([]byte, 532)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	return nil
}

// UserPrimaryGroupInformation structure represents USER_PRIMARY_GROUP_INFORMATION RPC structure.
//
// The USER_PRIMARY_GROUP_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserPrimaryGroupInformation struct {
	PrimaryGroupID uint32 `idl:"name:PrimaryGroupId" json:"primary_group_id"`
}

func (o *UserPrimaryGroupInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserPrimaryGroupInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PrimaryGroupID); err != nil {
		return err
	}
	return nil
}
func (o *UserPrimaryGroupInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrimaryGroupID); err != nil {
		return err
	}
	return nil
}

// UserControlInformation structure represents USER_CONTROL_INFORMATION RPC structure.
//
// The USER_CONTROL_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserControlInformation struct {
	UserAccountControl uint32 `idl:"name:UserAccountControl" json:"user_account_control"`
}

func (o *UserControlInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserControlInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.UserAccountControl); err != nil {
		return err
	}
	return nil
}
func (o *UserControlInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserAccountControl); err != nil {
		return err
	}
	return nil
}

// UserExpiresInformation structure represents USER_EXPIRES_INFORMATION RPC structure.
//
// The USER_EXPIRES_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserExpiresInformation struct {
	AccountExpires *OldLargeInteger `idl:"name:AccountExpires" json:"account_expires"`
}

func (o *UserExpiresInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserExpiresInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.AccountExpires != nil {
		if err := o.AccountExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserExpiresInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.AccountExpires == nil {
		o.AccountExpires = &OldLargeInteger{}
	}
	if err := o.AccountExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// LogonHours structure represents SAMPR_LOGON_HOURS RPC structure.
//
// The SAMPR_LOGON_HOURS structure contains logon policy information that describes
// when a user account is permitted to authenticate.
type LogonHours struct {
	// UnitsPerWeek:  A division of the week (7 days). For example, the value 7 means that
	// each unit is a day; a value of (7*24) means that the units are hours. The minimum
	// granularity of time is one minute, where the UnitsPerWeek would be 10080; therefore,
	// the maximum size of LogonHours is 10080/8, or 1,260 bytes.
	UnitsPerWeek uint16 `idl:"name:UnitsPerWeek" json:"units_per_week"`
	// LogonHours:  A pointer to a bit field containing at least UnitsPerWeek number of
	// bits. The leftmost bit represents the first unit, starting at Sunday, 12 A.M. If
	// a bit is set, authentication is allowed to occur; otherwise, authentication is not
	// allowed to occur.
	LogonHours []byte `idl:"name:LogonHours;size_is:(1260);length_is:(((UnitsPerWeek+7)/8))" json:"logon_hours"`
}

func (o *LogonHours) xxx_PreparePayload(ctx context.Context) error {
	// cannot evaluate expression 1260
	if o.LogonHours != nil && o.UnitsPerWeek == 0 {
		o.UnitsPerWeek = uint16(((len(o.LogonHours) * 8) - 7))
		if len(o.LogonHours) < 7 {
			o.UnitsPerWeek = uint16(0)
		}
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LogonHours) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.UnitsPerWeek); err != nil {
		return err
	}
	if o.LogonHours != nil || 1260 > 0 {
		_ptr_LogonHours := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(1260)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64(((o.UnitsPerWeek + 7) / 8))
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
			for i1 := range o.LogonHours {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.LogonHours[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.LogonHours); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LogonHours, _ptr_LogonHours); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *LogonHours) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.UnitsPerWeek); err != nil {
		return err
	}
	_ptr_LogonHours := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			return fmt.Errorf("buffer overflow for size %d of array o.LogonHours", sizeInfo[0])
		}
		o.LogonHours = make([]byte, sizeInfo[0])
		for i1 := range o.LogonHours {
			i1 := i1
			if err := w.ReadData(&o.LogonHours[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_LogonHours := func(ptr interface{}) { o.LogonHours = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.LogonHours, _s_LogonHours, _ptr_LogonHours); err != nil {
		return err
	}
	return nil
}

// UserAllInformation structure represents SAMPR_USER_ALL_INFORMATION RPC structure.
//
// The SAMPR_USER_ALL_INFORMATION structure contains user attribute information. Most
// fields are described in section 2.2.6.1. The exceptions are described below.
type UserAllInformation struct {
	LastLogon          *OldLargeInteger    `idl:"name:LastLogon" json:"last_logon"`
	LastLogoff         *OldLargeInteger    `idl:"name:LastLogoff" json:"last_logoff"`
	PasswordLastSet    *OldLargeInteger    `idl:"name:PasswordLastSet" json:"password_last_set"`
	AccountExpires     *OldLargeInteger    `idl:"name:AccountExpires" json:"account_expires"`
	PasswordCanChange  *OldLargeInteger    `idl:"name:PasswordCanChange" json:"password_can_change"`
	PasswordMustChange *OldLargeInteger    `idl:"name:PasswordMustChange" json:"password_must_change"`
	UserName           *dtyp.UnicodeString `idl:"name:UserName" json:"user_name"`
	FullName           *dtyp.UnicodeString `idl:"name:FullName" json:"full_name"`
	HomeDirectory      *dtyp.UnicodeString `idl:"name:HomeDirectory" json:"home_directory"`
	HomeDirectoryDrive *dtyp.UnicodeString `idl:"name:HomeDirectoryDrive" json:"home_directory_drive"`
	ScriptPath         *dtyp.UnicodeString `idl:"name:ScriptPath" json:"script_path"`
	ProfilePath        *dtyp.UnicodeString `idl:"name:ProfilePath" json:"profile_path"`
	AdminComment       *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
	WorkStations       *dtyp.UnicodeString `idl:"name:WorkStations" json:"work_stations"`
	UserComment        *dtyp.UnicodeString `idl:"name:UserComment" json:"user_comment"`
	Parameters         *dtyp.UnicodeString `idl:"name:Parameters" json:"parameters"`
	// LmOwfPassword:  An RPC_SHORT_BLOB structure where Length and MaximumLength MUST be
	// 16, and the Buffer MUST be formatted with an ENCRYPTED_LM_OWF_PASSWORD structure
	// with the cleartext value being an LM hash, and the encryption key being the 16-byte
	// SMB session key obtained as specified in either section 3.1.2.4 or section 3.2.2.3.
	LMOWFPassword *ShortBlob `idl:"name:LmOwfPassword" json:"lm_owf_password"`
	// NtOwfPassword:  An RPC_SHORT_BLOB structure where Length and MaximumLength MUST be
	// 16, and the Buffer MUST be formatted with an ENCRYPTED_NT_OWF_PASSWORD structure
	// with the cleartext value being an NT hash, and the encryption key being the 16-byte
	// SMB session key obtained as specified in either section 3.1.2.4 or section 3.2.2.3.
	NTOWFPassword *ShortBlob `idl:"name:NtOwfPassword" json:"nt_owf_password"`
	// PrivateData:  Not used. Ignored on receipt at the server and client. Clients MUST
	// set to zero when sent, and servers MUST set to zero on return.
	PrivateData *dtyp.UnicodeString `idl:"name:PrivateData" json:"private_data"`
	// SecurityDescriptor:  Not used. Ignored on receipt at the server and client. Clients
	// MUST set to zero when sent, and servers MUST set to zero on return.
	SecurityDescriptor *SrSecurityDescriptor `idl:"name:SecurityDescriptor" json:"security_descriptor"`
	UserID             uint32                `idl:"name:UserId" json:"user_id"`
	PrimaryGroupID     uint32                `idl:"name:PrimaryGroupId" json:"primary_group_id"`
	UserAccountControl uint32                `idl:"name:UserAccountControl" json:"user_account_control"`
	// WhichFields:  A 32-bit bit field indicating which fields within the SAMPR_USER_ALL_INFORMATION
	// structure will be processed by the server. Section 2.2.1.8 specifies the valid bits
	// and also specifies the structure field to which each bit corresponds.
	WhichFields      uint32      `idl:"name:WhichFields" json:"which_fields"`
	LogonHours       *LogonHours `idl:"name:LogonHours" json:"logon_hours"`
	BadPasswordCount uint16      `idl:"name:BadPasswordCount" json:"bad_password_count"`
	LogonCount       uint16      `idl:"name:LogonCount" json:"logon_count"`
	CountryCode      uint16      `idl:"name:CountryCode" json:"country_code"`
	CodePage         uint16      `idl:"name:CodePage" json:"code_page"`
	// LmPasswordPresent:  If zero, LmOwfPassword MUST be ignored; otherwise, LmOwfPassword
	// MUST be processed.
	LMPasswordPresent uint8 `idl:"name:LmPasswordPresent" json:"lm_password_present"`
	// NtPasswordPresent:  If zero, NtOwfPassword MUST be ignored; otherwise, NtOwfPassword
	// MUST be processed.
	NTPasswordPresent uint8 `idl:"name:NtPasswordPresent" json:"nt_password_present"`
	PasswordExpired   uint8 `idl:"name:PasswordExpired" json:"password_expired"`
	// PrivateDataSensitive:  Not used. Ignored on receipt at the server and client.
	PrivateDataSensitive uint8 `idl:"name:PrivateDataSensitive" json:"private_data_sensitive"`
}

func (o *UserAllInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserAllInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LastLogon != nil {
		if err := o.LastLogon.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastLogoff != nil {
		if err := o.LastLogoff.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordLastSet != nil {
		if err := o.PasswordLastSet.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AccountExpires != nil {
		if err := o.AccountExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordCanChange != nil {
		if err := o.PasswordCanChange.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordMustChange != nil {
		if err := o.PasswordMustChange.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.UserName != nil {
		if err := o.UserName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FullName != nil {
		if err := o.FullName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeDirectory != nil {
		if err := o.HomeDirectory.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeDirectoryDrive != nil {
		if err := o.HomeDirectoryDrive.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ScriptPath != nil {
		if err := o.ScriptPath.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ProfilePath != nil {
		if err := o.ProfilePath.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.WorkStations != nil {
		if err := o.WorkStations.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.UserComment != nil {
		if err := o.UserComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Parameters != nil {
		if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LMOWFPassword != nil {
		if err := o.LMOWFPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ShortBlob{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.NTOWFPassword != nil {
		if err := o.NTOWFPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ShortBlob{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PrivateData != nil {
		if err := o.PrivateData.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SecurityDescriptor != nil {
		if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SrSecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UserID); err != nil {
		return err
	}
	if err := w.WriteData(o.PrimaryGroupID); err != nil {
		return err
	}
	if err := w.WriteData(o.UserAccountControl); err != nil {
		return err
	}
	if err := w.WriteData(o.WhichFields); err != nil {
		return err
	}
	if o.LogonHours != nil {
		if err := o.LogonHours.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LogonHours{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.WriteData(o.LogonCount); err != nil {
		return err
	}
	if err := w.WriteData(o.CountryCode); err != nil {
		return err
	}
	if err := w.WriteData(o.CodePage); err != nil {
		return err
	}
	if err := w.WriteData(o.LMPasswordPresent); err != nil {
		return err
	}
	if err := w.WriteData(o.NTPasswordPresent); err != nil {
		return err
	}
	if err := w.WriteData(o.PasswordExpired); err != nil {
		return err
	}
	if err := w.WriteData(o.PrivateDataSensitive); err != nil {
		return err
	}
	return nil
}
func (o *UserAllInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LastLogon == nil {
		o.LastLogon = &OldLargeInteger{}
	}
	if err := o.LastLogon.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastLogoff == nil {
		o.LastLogoff = &OldLargeInteger{}
	}
	if err := o.LastLogoff.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordLastSet == nil {
		o.PasswordLastSet = &OldLargeInteger{}
	}
	if err := o.PasswordLastSet.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AccountExpires == nil {
		o.AccountExpires = &OldLargeInteger{}
	}
	if err := o.AccountExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordCanChange == nil {
		o.PasswordCanChange = &OldLargeInteger{}
	}
	if err := o.PasswordCanChange.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordMustChange == nil {
		o.PasswordMustChange = &OldLargeInteger{}
	}
	if err := o.PasswordMustChange.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.UserName == nil {
		o.UserName = &dtyp.UnicodeString{}
	}
	if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.FullName == nil {
		o.FullName = &dtyp.UnicodeString{}
	}
	if err := o.FullName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeDirectory == nil {
		o.HomeDirectory = &dtyp.UnicodeString{}
	}
	if err := o.HomeDirectory.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeDirectoryDrive == nil {
		o.HomeDirectoryDrive = &dtyp.UnicodeString{}
	}
	if err := o.HomeDirectoryDrive.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ScriptPath == nil {
		o.ScriptPath = &dtyp.UnicodeString{}
	}
	if err := o.ScriptPath.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ProfilePath == nil {
		o.ProfilePath = &dtyp.UnicodeString{}
	}
	if err := o.ProfilePath.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.WorkStations == nil {
		o.WorkStations = &dtyp.UnicodeString{}
	}
	if err := o.WorkStations.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.UserComment == nil {
		o.UserComment = &dtyp.UnicodeString{}
	}
	if err := o.UserComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Parameters == nil {
		o.Parameters = &dtyp.UnicodeString{}
	}
	if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LMOWFPassword == nil {
		o.LMOWFPassword = &ShortBlob{}
	}
	if err := o.LMOWFPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.NTOWFPassword == nil {
		o.NTOWFPassword = &ShortBlob{}
	}
	if err := o.NTOWFPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PrivateData == nil {
		o.PrivateData = &dtyp.UnicodeString{}
	}
	if err := o.PrivateData.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SecurityDescriptor == nil {
		o.SecurityDescriptor = &SrSecurityDescriptor{}
	}
	if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrimaryGroupID); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserAccountControl); err != nil {
		return err
	}
	if err := w.ReadData(&o.WhichFields); err != nil {
		return err
	}
	if o.LogonHours == nil {
		o.LogonHours = &LogonHours{}
	}
	if err := o.LogonHours.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogonCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.CountryCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.CodePage); err != nil {
		return err
	}
	if err := w.ReadData(&o.LMPasswordPresent); err != nil {
		return err
	}
	if err := w.ReadData(&o.NTPasswordPresent); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordExpired); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrivateDataSensitive); err != nil {
		return err
	}
	return nil
}

// UserGeneralInformation structure represents SAMPR_USER_GENERAL_INFORMATION RPC structure.
//
// The SAMPR_USER_GENERAL_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserGeneralInformation struct {
	UserName       *dtyp.UnicodeString `idl:"name:UserName" json:"user_name"`
	FullName       *dtyp.UnicodeString `idl:"name:FullName" json:"full_name"`
	PrimaryGroupID uint32              `idl:"name:PrimaryGroupId" json:"primary_group_id"`
	AdminComment   *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
	UserComment    *dtyp.UnicodeString `idl:"name:UserComment" json:"user_comment"`
}

func (o *UserGeneralInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserGeneralInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.UserName != nil {
		if err := o.UserName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FullName != nil {
		if err := o.FullName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PrimaryGroupID); err != nil {
		return err
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.UserComment != nil {
		if err := o.UserComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserGeneralInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.UserName == nil {
		o.UserName = &dtyp.UnicodeString{}
	}
	if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.FullName == nil {
		o.FullName = &dtyp.UnicodeString{}
	}
	if err := o.FullName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrimaryGroupID); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.UserComment == nil {
		o.UserComment = &dtyp.UnicodeString{}
	}
	if err := o.UserComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserPreferencesInformation structure represents SAMPR_USER_PREFERENCES_INFORMATION RPC structure.
//
// The SAMPR_USER_PREFERENCES_INFORMATION structure contains user fields.
type UserPreferencesInformation struct {
	UserComment *dtyp.UnicodeString `idl:"name:UserComment" json:"user_comment"`
	// Reserved1:  Ignored by the client and server and MUST be a zero-length string when
	// sent and returned.
	//
	// For information on all other fields, see section 2.2.6.1.
	_           *dtyp.UnicodeString `idl:"name:Reserved1"`
	CountryCode uint16              `idl:"name:CountryCode" json:"country_code"`
	CodePage    uint16              `idl:"name:CodePage" json:"code_page"`
}

func (o *UserPreferencesInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserPreferencesInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.UserComment != nil {
		if err := o.UserComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved Reserved1
	if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.WriteData(o.CountryCode); err != nil {
		return err
	}
	if err := w.WriteData(o.CodePage); err != nil {
		return err
	}
	return nil
}
func (o *UserPreferencesInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.UserComment == nil {
		o.UserComment = &dtyp.UnicodeString{}
	}
	if err := o.UserComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	// reserved Reserved1
	var _Reserved1 *dtyp.UnicodeString
	if _Reserved1 == nil {
		_Reserved1 = &dtyp.UnicodeString{}
	}
	if err := _Reserved1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.CountryCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.CodePage); err != nil {
		return err
	}
	return nil
}

// UserParametersInformation structure represents SAMPR_USER_PARAMETERS_INFORMATION RPC structure.
//
// The SAMPR_USER_PARAMETERS_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserParametersInformation struct {
	Parameters *dtyp.UnicodeString `idl:"name:Parameters" json:"parameters"`
}

func (o *UserParametersInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserParametersInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.Parameters != nil {
		if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserParametersInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.Parameters == nil {
		o.Parameters = &dtyp.UnicodeString{}
	}
	if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserLogonInformation structure represents SAMPR_USER_LOGON_INFORMATION RPC structure.
//
// The SAMPR_USER_LOGON_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserLogonInformation struct {
	UserName           *dtyp.UnicodeString `idl:"name:UserName" json:"user_name"`
	FullName           *dtyp.UnicodeString `idl:"name:FullName" json:"full_name"`
	UserID             uint32              `idl:"name:UserId" json:"user_id"`
	PrimaryGroupID     uint32              `idl:"name:PrimaryGroupId" json:"primary_group_id"`
	HomeDirectory      *dtyp.UnicodeString `idl:"name:HomeDirectory" json:"home_directory"`
	HomeDirectoryDrive *dtyp.UnicodeString `idl:"name:HomeDirectoryDrive" json:"home_directory_drive"`
	ScriptPath         *dtyp.UnicodeString `idl:"name:ScriptPath" json:"script_path"`
	ProfilePath        *dtyp.UnicodeString `idl:"name:ProfilePath" json:"profile_path"`
	WorkStations       *dtyp.UnicodeString `idl:"name:WorkStations" json:"work_stations"`
	LastLogon          *OldLargeInteger    `idl:"name:LastLogon" json:"last_logon"`
	LastLogoff         *OldLargeInteger    `idl:"name:LastLogoff" json:"last_logoff"`
	PasswordLastSet    *OldLargeInteger    `idl:"name:PasswordLastSet" json:"password_last_set"`
	PasswordCanChange  *OldLargeInteger    `idl:"name:PasswordCanChange" json:"password_can_change"`
	PasswordMustChange *OldLargeInteger    `idl:"name:PasswordMustChange" json:"password_must_change"`
	LogonHours         *LogonHours         `idl:"name:LogonHours" json:"logon_hours"`
	BadPasswordCount   uint16              `idl:"name:BadPasswordCount" json:"bad_password_count"`
	LogonCount         uint16              `idl:"name:LogonCount" json:"logon_count"`
	UserAccountControl uint32              `idl:"name:UserAccountControl" json:"user_account_control"`
}

func (o *UserLogonInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserLogonInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.UserName != nil {
		if err := o.UserName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FullName != nil {
		if err := o.FullName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UserID); err != nil {
		return err
	}
	if err := w.WriteData(o.PrimaryGroupID); err != nil {
		return err
	}
	if o.HomeDirectory != nil {
		if err := o.HomeDirectory.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeDirectoryDrive != nil {
		if err := o.HomeDirectoryDrive.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ScriptPath != nil {
		if err := o.ScriptPath.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ProfilePath != nil {
		if err := o.ProfilePath.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.WorkStations != nil {
		if err := o.WorkStations.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastLogon != nil {
		if err := o.LastLogon.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastLogoff != nil {
		if err := o.LastLogoff.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordLastSet != nil {
		if err := o.PasswordLastSet.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordCanChange != nil {
		if err := o.PasswordCanChange.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordMustChange != nil {
		if err := o.PasswordMustChange.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogonHours != nil {
		if err := o.LogonHours.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LogonHours{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.WriteData(o.LogonCount); err != nil {
		return err
	}
	if err := w.WriteData(o.UserAccountControl); err != nil {
		return err
	}
	return nil
}
func (o *UserLogonInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.UserName == nil {
		o.UserName = &dtyp.UnicodeString{}
	}
	if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.FullName == nil {
		o.FullName = &dtyp.UnicodeString{}
	}
	if err := o.FullName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrimaryGroupID); err != nil {
		return err
	}
	if o.HomeDirectory == nil {
		o.HomeDirectory = &dtyp.UnicodeString{}
	}
	if err := o.HomeDirectory.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeDirectoryDrive == nil {
		o.HomeDirectoryDrive = &dtyp.UnicodeString{}
	}
	if err := o.HomeDirectoryDrive.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ScriptPath == nil {
		o.ScriptPath = &dtyp.UnicodeString{}
	}
	if err := o.ScriptPath.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ProfilePath == nil {
		o.ProfilePath = &dtyp.UnicodeString{}
	}
	if err := o.ProfilePath.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.WorkStations == nil {
		o.WorkStations = &dtyp.UnicodeString{}
	}
	if err := o.WorkStations.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastLogon == nil {
		o.LastLogon = &OldLargeInteger{}
	}
	if err := o.LastLogon.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastLogoff == nil {
		o.LastLogoff = &OldLargeInteger{}
	}
	if err := o.LastLogoff.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordLastSet == nil {
		o.PasswordLastSet = &OldLargeInteger{}
	}
	if err := o.PasswordLastSet.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordCanChange == nil {
		o.PasswordCanChange = &OldLargeInteger{}
	}
	if err := o.PasswordCanChange.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordMustChange == nil {
		o.PasswordMustChange = &OldLargeInteger{}
	}
	if err := o.PasswordMustChange.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogonHours == nil {
		o.LogonHours = &LogonHours{}
	}
	if err := o.LogonHours.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogonCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserAccountControl); err != nil {
		return err
	}
	return nil
}

// UserAccountInformation structure represents SAMPR_USER_ACCOUNT_INFORMATION RPC structure.
//
// The SAMPR_USER_ACCOUNT_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserAccountInformation struct {
	UserName           *dtyp.UnicodeString `idl:"name:UserName" json:"user_name"`
	FullName           *dtyp.UnicodeString `idl:"name:FullName" json:"full_name"`
	UserID             uint32              `idl:"name:UserId" json:"user_id"`
	PrimaryGroupID     uint32              `idl:"name:PrimaryGroupId" json:"primary_group_id"`
	HomeDirectory      *dtyp.UnicodeString `idl:"name:HomeDirectory" json:"home_directory"`
	HomeDirectoryDrive *dtyp.UnicodeString `idl:"name:HomeDirectoryDrive" json:"home_directory_drive"`
	ScriptPath         *dtyp.UnicodeString `idl:"name:ScriptPath" json:"script_path"`
	ProfilePath        *dtyp.UnicodeString `idl:"name:ProfilePath" json:"profile_path"`
	AdminComment       *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
	WorkStations       *dtyp.UnicodeString `idl:"name:WorkStations" json:"work_stations"`
	LastLogon          *OldLargeInteger    `idl:"name:LastLogon" json:"last_logon"`
	LastLogoff         *OldLargeInteger    `idl:"name:LastLogoff" json:"last_logoff"`
	LogonHours         *LogonHours         `idl:"name:LogonHours" json:"logon_hours"`
	BadPasswordCount   uint16              `idl:"name:BadPasswordCount" json:"bad_password_count"`
	LogonCount         uint16              `idl:"name:LogonCount" json:"logon_count"`
	PasswordLastSet    *OldLargeInteger    `idl:"name:PasswordLastSet" json:"password_last_set"`
	AccountExpires     *OldLargeInteger    `idl:"name:AccountExpires" json:"account_expires"`
	UserAccountControl uint32              `idl:"name:UserAccountControl" json:"user_account_control"`
}

func (o *UserAccountInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserAccountInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.UserName != nil {
		if err := o.UserName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FullName != nil {
		if err := o.FullName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UserID); err != nil {
		return err
	}
	if err := w.WriteData(o.PrimaryGroupID); err != nil {
		return err
	}
	if o.HomeDirectory != nil {
		if err := o.HomeDirectory.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeDirectoryDrive != nil {
		if err := o.HomeDirectoryDrive.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ScriptPath != nil {
		if err := o.ScriptPath.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ProfilePath != nil {
		if err := o.ProfilePath.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.WorkStations != nil {
		if err := o.WorkStations.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastLogon != nil {
		if err := o.LastLogon.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LastLogoff != nil {
		if err := o.LastLogoff.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogonHours != nil {
		if err := o.LogonHours.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LogonHours{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.WriteData(o.LogonCount); err != nil {
		return err
	}
	if o.PasswordLastSet != nil {
		if err := o.PasswordLastSet.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AccountExpires != nil {
		if err := o.AccountExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OldLargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UserAccountControl); err != nil {
		return err
	}
	return nil
}
func (o *UserAccountInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.UserName == nil {
		o.UserName = &dtyp.UnicodeString{}
	}
	if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.FullName == nil {
		o.FullName = &dtyp.UnicodeString{}
	}
	if err := o.FullName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrimaryGroupID); err != nil {
		return err
	}
	if o.HomeDirectory == nil {
		o.HomeDirectory = &dtyp.UnicodeString{}
	}
	if err := o.HomeDirectory.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeDirectoryDrive == nil {
		o.HomeDirectoryDrive = &dtyp.UnicodeString{}
	}
	if err := o.HomeDirectoryDrive.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ScriptPath == nil {
		o.ScriptPath = &dtyp.UnicodeString{}
	}
	if err := o.ScriptPath.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ProfilePath == nil {
		o.ProfilePath = &dtyp.UnicodeString{}
	}
	if err := o.ProfilePath.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.WorkStations == nil {
		o.WorkStations = &dtyp.UnicodeString{}
	}
	if err := o.WorkStations.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastLogon == nil {
		o.LastLogon = &OldLargeInteger{}
	}
	if err := o.LastLogon.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LastLogoff == nil {
		o.LastLogoff = &OldLargeInteger{}
	}
	if err := o.LastLogoff.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogonHours == nil {
		o.LogonHours = &LogonHours{}
	}
	if err := o.LogonHours.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogonCount); err != nil {
		return err
	}
	if o.PasswordLastSet == nil {
		o.PasswordLastSet = &OldLargeInteger{}
	}
	if err := o.PasswordLastSet.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AccountExpires == nil {
		o.AccountExpires = &OldLargeInteger{}
	}
	if err := o.AccountExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserAccountControl); err != nil {
		return err
	}
	return nil
}

// UserANameInformation structure represents SAMPR_USER_A_NAME_INFORMATION RPC structure.
//
// The SAMPR_USER_A_NAME_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserANameInformation struct {
	UserName *dtyp.UnicodeString `idl:"name:UserName" json:"user_name"`
}

func (o *UserANameInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserANameInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.UserName != nil {
		if err := o.UserName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserANameInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.UserName == nil {
		o.UserName = &dtyp.UnicodeString{}
	}
	if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserFNameInformation structure represents SAMPR_USER_F_NAME_INFORMATION RPC structure.
//
// The SAMPR_USER_F_NAME_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserFNameInformation struct {
	FullName *dtyp.UnicodeString `idl:"name:FullName" json:"full_name"`
}

func (o *UserFNameInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserFNameInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.FullName != nil {
		if err := o.FullName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserFNameInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.FullName == nil {
		o.FullName = &dtyp.UnicodeString{}
	}
	if err := o.FullName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserNameInformation structure represents SAMPR_USER_NAME_INFORMATION RPC structure.
//
// The SAMPR_USER_NAME_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserNameInformation struct {
	UserName *dtyp.UnicodeString `idl:"name:UserName" json:"user_name"`
	FullName *dtyp.UnicodeString `idl:"name:FullName" json:"full_name"`
}

func (o *UserNameInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserNameInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.UserName != nil {
		if err := o.UserName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FullName != nil {
		if err := o.FullName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserNameInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.UserName == nil {
		o.UserName = &dtyp.UnicodeString{}
	}
	if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.FullName == nil {
		o.FullName = &dtyp.UnicodeString{}
	}
	if err := o.FullName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserHomeInformation structure represents SAMPR_USER_HOME_INFORMATION RPC structure.
//
// The SAMPR_USER_HOME_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserHomeInformation struct {
	HomeDirectory      *dtyp.UnicodeString `idl:"name:HomeDirectory" json:"home_directory"`
	HomeDirectoryDrive *dtyp.UnicodeString `idl:"name:HomeDirectoryDrive" json:"home_directory_drive"`
}

func (o *UserHomeInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserHomeInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.HomeDirectory != nil {
		if err := o.HomeDirectory.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeDirectoryDrive != nil {
		if err := o.HomeDirectoryDrive.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserHomeInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.HomeDirectory == nil {
		o.HomeDirectory = &dtyp.UnicodeString{}
	}
	if err := o.HomeDirectory.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeDirectoryDrive == nil {
		o.HomeDirectoryDrive = &dtyp.UnicodeString{}
	}
	if err := o.HomeDirectoryDrive.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserScriptInformation structure represents SAMPR_USER_SCRIPT_INFORMATION RPC structure.
//
// The SAMPR_USER_SCRIPT_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserScriptInformation struct {
	ScriptPath *dtyp.UnicodeString `idl:"name:ScriptPath" json:"script_path"`
}

func (o *UserScriptInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserScriptInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.ScriptPath != nil {
		if err := o.ScriptPath.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserScriptInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.ScriptPath == nil {
		o.ScriptPath = &dtyp.UnicodeString{}
	}
	if err := o.ScriptPath.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserProfileInformation structure represents SAMPR_USER_PROFILE_INFORMATION RPC structure.
//
// The SAMPR_USER_PROFILE_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserProfileInformation struct {
	ProfilePath *dtyp.UnicodeString `idl:"name:ProfilePath" json:"profile_path"`
}

func (o *UserProfileInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProfileInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.ProfilePath != nil {
		if err := o.ProfilePath.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserProfileInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.ProfilePath == nil {
		o.ProfilePath = &dtyp.UnicodeString{}
	}
	if err := o.ProfilePath.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserAdminCommentInformation structure represents SAMPR_USER_ADMIN_COMMENT_INFORMATION RPC structure.
//
// The SAMPR_USER_ADMIN_COMMENT_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserAdminCommentInformation struct {
	AdminComment *dtyp.UnicodeString `idl:"name:AdminComment" json:"admin_comment"`
}

func (o *UserAdminCommentInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserAdminCommentInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserAdminCommentInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.AdminComment == nil {
		o.AdminComment = &dtyp.UnicodeString{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserWorkstationsInformation structure represents SAMPR_USER_WORKSTATIONS_INFORMATION RPC structure.
//
// The SAMPR_USER_WORKSTATIONS_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserWorkstationsInformation struct {
	WorkStations *dtyp.UnicodeString `idl:"name:WorkStations" json:"work_stations"`
}

func (o *UserWorkstationsInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserWorkstationsInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.WorkStations != nil {
		if err := o.WorkStations.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserWorkstationsInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.WorkStations == nil {
		o.WorkStations = &dtyp.UnicodeString{}
	}
	if err := o.WorkStations.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserLogonHoursInformation structure represents SAMPR_USER_LOGON_HOURS_INFORMATION RPC structure.
//
// The SAMPR_USER_LOGON_HOURS_INFORMATION structure contains user fields.
//
// For information on each field, see section 2.2.6.1.
type UserLogonHoursInformation struct {
	LogonHours *LogonHours `idl:"name:LogonHours" json:"logon_hours"`
}

func (o *UserLogonHoursInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserLogonHoursInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.LogonHours != nil {
		if err := o.LogonHours.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LogonHours{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserLogonHoursInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.LogonHours == nil {
		o.LogonHours = &LogonHours{}
	}
	if err := o.LogonHours.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInternal1Information structure represents SAMPR_USER_INTERNAL1_INFORMATION RPC structure.
//
// The SAMPR_USER_INTERNAL1_INFORMATION structure holds the hashed form of a cleartext
// password.
type UserInternal1Information struct {
	// EncryptedNtOwfPassword:  An NT hash encrypted with the 16-byte SMB session key obtained
	// as specified in either section 3.1.2.4 or section 3.2.2.3.
	EncryptedNTOWFPassword *EncryptedNTOWFPassword `idl:"name:EncryptedNtOwfPassword" json:"encrypted_nt_owf_password"`
	// EncryptedLmOwfPassword:  An LM hash encrypted with the 16-byte SMB session key obtained
	// as specified in either section 3.1.2.4 or section 3.2.2.3.
	EncryptedLMOWFPassword *EncryptedLMOWFPassword `idl:"name:EncryptedLmOwfPassword" json:"encrypted_lm_owf_password"`
	// NtPasswordPresent:  If nonzero, indicates that the EncryptedNtOwfPassword value is
	// valid; otherwise, EncryptedNtOwfPassword MUST be ignored.
	NTPasswordPresent uint8 `idl:"name:NtPasswordPresent" json:"nt_password_present"`
	// LmPasswordPresent:  If nonzero, indicates that the EncryptedLmOwfPassword value is
	// valid; otherwise, EncryptedLmOwfPassword MUST be ignored.
	LMPasswordPresent uint8 `idl:"name:LmPasswordPresent" json:"lm_password_present"`
	// PasswordExpired:  See section 2.2.6.1.
	PasswordExpired uint8 `idl:"name:PasswordExpired" json:"password_expired"`
}

func (o *UserInternal1Information) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInternal1Information) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if o.EncryptedNTOWFPassword != nil {
		if err := o.EncryptedNTOWFPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&EncryptedNTOWFPassword{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.EncryptedLMOWFPassword != nil {
		if err := o.EncryptedLMOWFPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&EncryptedLMOWFPassword{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NTPasswordPresent); err != nil {
		return err
	}
	if err := w.WriteData(o.LMPasswordPresent); err != nil {
		return err
	}
	if err := w.WriteData(o.PasswordExpired); err != nil {
		return err
	}
	return nil
}
func (o *UserInternal1Information) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.EncryptedNTOWFPassword == nil {
		o.EncryptedNTOWFPassword = &EncryptedNTOWFPassword{}
	}
	if err := o.EncryptedNTOWFPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.EncryptedLMOWFPassword == nil {
		o.EncryptedLMOWFPassword = &EncryptedLMOWFPassword{}
	}
	if err := o.EncryptedLMOWFPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.NTPasswordPresent); err != nil {
		return err
	}
	if err := w.ReadData(&o.LMPasswordPresent); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordExpired); err != nil {
		return err
	}
	return nil
}

// UserInternal4Information structure represents SAMPR_USER_INTERNAL4_INFORMATION RPC structure.
//
// The SAMPR_USER_INTERNAL4_INFORMATION structure holds all attributes of a user, along
// with an encrypted password.
type UserInternal4Information struct {
	// I1:  See section 2.2.6.6.
	I1 *UserAllInformation `idl:"name:I1" json:"i1"`
	// UserPassword:  See section 2.2.6.21.
	UserPassword *EncryptedUserPassword `idl:"name:UserPassword" json:"user_password"`
}

func (o *UserInternal4Information) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInternal4Information) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.I1 != nil {
		if err := o.I1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserAllInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.UserPassword != nil {
		if err := o.UserPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInternal4Information) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.I1 == nil {
		o.I1 = &UserAllInformation{}
	}
	if err := o.I1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.UserPassword == nil {
		o.UserPassword = &EncryptedUserPassword{}
	}
	if err := o.UserPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInternal4InformationNew structure represents SAMPR_USER_INTERNAL4_INFORMATION_NEW RPC structure.
//
// The SAMPR_USER_INTERNAL4_INFORMATION_NEW structure holds all attributes of a user,
// along with an encrypted password. The encrypted password uses a salt to improve the
// encryption algorithm. See the specification for SAMPR_ENCRYPTED_USER_PASSWORD_NEW
// (section 2.2.6.22) for details on salt value selection.
type UserInternal4InformationNew struct {
	// I1:  See section 2.2.6.6.
	I1 *UserAllInformation `idl:"name:I1" json:"i1"`
	// UserPassword:  See section 2.2.6.22.
	UserPassword *EncryptedUserPasswordNew `idl:"name:UserPassword" json:"user_password"`
}

func (o *UserInternal4InformationNew) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInternal4InformationNew) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.I1 != nil {
		if err := o.I1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserAllInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.UserPassword != nil {
		if err := o.UserPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&EncryptedUserPasswordNew{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInternal4InformationNew) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.I1 == nil {
		o.I1 = &UserAllInformation{}
	}
	if err := o.I1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.UserPassword == nil {
		o.UserPassword = &EncryptedUserPasswordNew{}
	}
	if err := o.UserPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInternal5Information structure represents SAMPR_USER_INTERNAL5_INFORMATION RPC structure.
//
// The SAMPR_USER_INTERNAL5_INFORMATION structure holds an encrypted password.
//
// This structure is used to carry a new password for a particular account from the
// client to the server, encrypted in a way that protects it from disclosure or tampering
// while in transit.
type UserInternal5Information struct {
	// UserPassword:  A cleartext password, encrypted according to the specification for
	// SAMPR_ENCRYPTED_USER_PASSWORD, with the encryption key being the 16-byte SMB session
	// key obtained as specified in either section 3.1.2.4 or section 3.2.2.3.
	UserPassword *EncryptedUserPassword `idl:"name:UserPassword" json:"user_password"`
	// PasswordExpired:  See section 2.2.6.1.
	PasswordExpired uint8 `idl:"name:PasswordExpired" json:"password_expired"`
}

func (o *UserInternal5Information) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInternal5Information) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if o.UserPassword != nil {
		if err := o.UserPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PasswordExpired); err != nil {
		return err
	}
	return nil
}
func (o *UserInternal5Information) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UserPassword == nil {
		o.UserPassword = &EncryptedUserPassword{}
	}
	if err := o.UserPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordExpired); err != nil {
		return err
	}
	return nil
}

// UserInternal5InformationNew structure represents SAMPR_USER_INTERNAL5_INFORMATION_NEW RPC structure.
//
// The SAMPR_USER_INTERNAL5_INFORMATION_NEW structure communicates an encrypted password.
// The encrypted password uses a salt to improve the encryption algorithm. See the specification
// for SAMPR_ENCRYPTED_USER_PASSWORD_NEW (section 2.2.6.22) for details on salt value
// selection.
//
// This structure is used to carry a new password for a particular account from the
// client to the server, encrypted in a way that protects it from disclosure or tampering
// while in transit. A random value, a salt, is used by the client to seed the encryption
// routine; see section 2.2.6.22 for details.
type UserInternal5InformationNew struct {
	// UserPassword:  A password, encrypted according to the specification for SAMPR_ENCRYPTED_USER_PASSWORD_NEW,
	// with the encryption key being the 16-byte SMB session key obtained as specified in
	// either section 3.1.2.4 or section 3.2.2.3.
	UserPassword *EncryptedUserPasswordNew `idl:"name:UserPassword" json:"user_password"`
	// PasswordExpired:  See section 2.2.6.1.
	PasswordExpired uint8 `idl:"name:PasswordExpired" json:"password_expired"`
}

func (o *UserInternal5InformationNew) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInternal5InformationNew) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if o.UserPassword != nil {
		if err := o.UserPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&EncryptedUserPasswordNew{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PasswordExpired); err != nil {
		return err
	}
	return nil
}
func (o *UserInternal5InformationNew) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UserPassword == nil {
		o.UserPassword = &EncryptedUserPasswordNew{}
	}
	if err := o.UserPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordExpired); err != nil {
		return err
	}
	return nil
}

// UserInformationClass type represents USER_INFORMATION_CLASS RPC enumeration.
//
// The USER_INFORMATION_CLASS enumeration indicates how to interpret the Buffer parameter
// for SamrSetInformationUser, SamrQueryInformationUser, SamrSetInformationUser2, and
// SamrQueryInformationUser2. For a list of associated structures, see section 2.2.6.29.
type UserInformationClass uint16

var (
	// UserGeneralInformation:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_USER_GENERAL_INFORMATION structure (see section 2.2.6.7).
	UserInformationClassGeneralInformation UserInformationClass = 1
	// UserPreferencesInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_USER_PREFERENCES_INFORMATION structure (see section 2.2.6.8).
	UserInformationClassPreferencesInformation UserInformationClass = 2
	// UserLogonInformation:  Indicates the Buffer parameter is to be interpreted as a
	// SAMPR_USER_LOGON_INFORMATION structure (see section 2.2.6.10).
	UserInformationClassLogonInformation UserInformationClass = 3
	// UserLogonHoursInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_USER_LOGON_HOURS_INFORMATION structure (see section 2.2.6.20).
	UserInformationClassLogonHoursInformation UserInformationClass = 4
	// UserAccountInformation:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_USER_ACCOUNT_INFORMATION structure (see section 2.2.6.11).
	UserInformationClassAccountInformation UserInformationClass = 5
	// UserNameInformation:  Indicates the Buffer parameter is to be interpreted as a SAMPR_USER_NAME_INFORMATION
	// structure (see section 2.2.6.14).
	UserInformationClassNameInformation UserInformationClass = 6
	// UserAccountNameInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_USER_A_NAME_INFORMATION structure (see section 2.2.6.12).
	UserInformationClassAccountNameInformation UserInformationClass = 7
	// UserFullNameInformation:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_USER_F_NAME_INFORMATION structure (see section 2.2.6.13).
	UserInformationClassFullNameInformation UserInformationClass = 8
	// UserPrimaryGroupInformation:  Indicates the Buffer parameter is to be interpreted
	// as a USER_PRIMARY_GROUP_INFORMATION structure (see section 2.2.6.2).
	UserInformationClassPrimaryGroupInformation UserInformationClass = 9
	// UserHomeInformation:  Indicates the Buffer parameter is to be interpreted as a SAMPR_USER_HOME_INFORMATION
	// structure (see section 2.2.6.15).
	UserInformationClassHomeInformation UserInformationClass = 10
	// UserScriptInformation:  Indicates the Buffer parameter is to be interpreted as a
	// SAMPR_USER_SCRIPT_INFORMATION structure (see section 2.2.6.16).
	UserInformationClassScriptInformation UserInformationClass = 11
	// UserProfileInformation:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_USER_PROFILE_INFORMATION structure (see section 2.2.6.17).
	UserInformationClassProfileInformation UserInformationClass = 12
	// UserAdminCommentInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_USER_ADMIN_COMMENT_INFORMATION structure (see section 2.2.6.18).
	UserInformationClassAdminCommentInformation UserInformationClass = 13
	// UserWorkStationsInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_USER_WORKSTATIONS_INFORMATION structure (see section 2.2.6.19).
	UserInformationClassWorkStationsInformation UserInformationClass = 14
	// UserControlInformation:  Indicates the Buffer parameter is to be interpreted as
	// a USER_CONTROL_INFORMATION structure (see section 2.2.6.3).
	UserInformationClassControlInformation UserInformationClass = 16
	// UserExpiresInformation:  Indicates the Buffer parameter is to be interpreted as
	// a USER_EXPIRES_INFORMATION structure (see section 2.2.6.4).
	UserInformationClassExpiresInformation UserInformationClass = 17
	// UserInternal1Information:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_USER_INTERNAL1_INFORMATION structure (see section 2.2.6.23).
	UserInformationClassInternal1Information UserInformationClass = 18
	// UserParametersInformation:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_USER_PARAMETERS_INFORMATION structure (see section 2.2.6.9).
	UserInformationClassParametersInformation UserInformationClass = 20
	// UserAllInformation:  Indicates the Buffer parameter is to be interpreted as a SAMPR_USER_ALL_INFORMATION
	// structure (see section 2.2.6.6).
	UserInformationClassAllInformation UserInformationClass = 21
	// UserInternal4Information:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_USER_INTERNAL4_INFORMATION structure (see section 2.2.6.24).
	UserInformationClassInternal4Information UserInformationClass = 23
	// UserInternal5Information:  Indicates the Buffer parameter is to be interpreted as
	// a SAMPR_USER_INTERNAL5_INFORMATION structure (see section 2.2.6.26).
	UserInformationClassInternal5Information UserInformationClass = 24
	// UserInternal4InformationNew:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_USER_INTERNAL4_INFORMATION_NEW structure (see section 2.2.6.25).
	UserInformationClassInternal4InformationNew UserInformationClass = 25
	// UserInternal5InformationNew:  Indicates the Buffer parameter is to be interpreted
	// as a SAMPR_USER_INTERNAL5_INFORMATION_NEW structure (see section 2.2.6.27).
	UserInformationClassInternal5InformationNew UserInformationClass = 26
)

func (o UserInformationClass) String() string {
	switch o {
	case UserInformationClassGeneralInformation:
		return "UserInformationClassGeneralInformation"
	case UserInformationClassPreferencesInformation:
		return "UserInformationClassPreferencesInformation"
	case UserInformationClassLogonInformation:
		return "UserInformationClassLogonInformation"
	case UserInformationClassLogonHoursInformation:
		return "UserInformationClassLogonHoursInformation"
	case UserInformationClassAccountInformation:
		return "UserInformationClassAccountInformation"
	case UserInformationClassNameInformation:
		return "UserInformationClassNameInformation"
	case UserInformationClassAccountNameInformation:
		return "UserInformationClassAccountNameInformation"
	case UserInformationClassFullNameInformation:
		return "UserInformationClassFullNameInformation"
	case UserInformationClassPrimaryGroupInformation:
		return "UserInformationClassPrimaryGroupInformation"
	case UserInformationClassHomeInformation:
		return "UserInformationClassHomeInformation"
	case UserInformationClassScriptInformation:
		return "UserInformationClassScriptInformation"
	case UserInformationClassProfileInformation:
		return "UserInformationClassProfileInformation"
	case UserInformationClassAdminCommentInformation:
		return "UserInformationClassAdminCommentInformation"
	case UserInformationClassWorkStationsInformation:
		return "UserInformationClassWorkStationsInformation"
	case UserInformationClassControlInformation:
		return "UserInformationClassControlInformation"
	case UserInformationClassExpiresInformation:
		return "UserInformationClassExpiresInformation"
	case UserInformationClassInternal1Information:
		return "UserInformationClassInternal1Information"
	case UserInformationClassParametersInformation:
		return "UserInformationClassParametersInformation"
	case UserInformationClassAllInformation:
		return "UserInformationClassAllInformation"
	case UserInformationClassInternal4Information:
		return "UserInformationClassInternal4Information"
	case UserInformationClassInternal5Information:
		return "UserInformationClassInternal5Information"
	case UserInformationClassInternal4InformationNew:
		return "UserInformationClassInternal4InformationNew"
	case UserInformationClassInternal5InformationNew:
		return "UserInformationClassInternal5InformationNew"
	}
	return "Invalid"
}

// UserInfoBuffer structure represents SAMPR_USER_INFO_BUFFER RPC union.
//
// The SAMPR_USER_INFO_BUFFER union combines all possible structures used in the SamrSetInformationUser,
// SamrSetInformationUser2, SamrQueryInformationUser, and SamrQueryInformationUser2
// methods (see sections 3.1.5.6.5, 3.1.5.6.4, 3.1.5.5.6, and 3.1.5.5.5). For details
// on each field, see the associated section for the field structure.
type UserInfoBuffer struct {
	// Types that are assignable to Value
	//
	// *UserInfoBuffer_General
	// *UserInfoBuffer_Preferences
	// *UserInfoBuffer_Logon
	// *UserInfoBuffer_LogonHours
	// *UserInfoBuffer_Account
	// *UserInfoBuffer_Name
	// *UserInfoBuffer_AccountName
	// *UserInfoBuffer_FullName
	// *UserInfoBuffer_PrimaryGroup
	// *UserInfoBuffer_Home
	// *UserInfoBuffer_Script
	// *UserInfoBuffer_Profile
	// *UserInfoBuffer_AdminComment
	// *UserInfoBuffer_WorkStations
	// *UserInfoBuffer_Control
	// *UserInfoBuffer_Expires
	// *UserInfoBuffer_Internal1
	// *UserInfoBuffer_Parameters
	// *UserInfoBuffer_All
	// *UserInfoBuffer_Internal4
	// *UserInfoBuffer_Internal5
	// *UserInfoBuffer_Internal4New
	// *UserInfoBuffer_Internal5New
	Value is_UserInfoBuffer `json:"value"`
}

func (o *UserInfoBuffer) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *UserInfoBuffer_General:
		if value != nil {
			return value.General
		}
	case *UserInfoBuffer_Preferences:
		if value != nil {
			return value.Preferences
		}
	case *UserInfoBuffer_Logon:
		if value != nil {
			return value.Logon
		}
	case *UserInfoBuffer_LogonHours:
		if value != nil {
			return value.LogonHours
		}
	case *UserInfoBuffer_Account:
		if value != nil {
			return value.Account
		}
	case *UserInfoBuffer_Name:
		if value != nil {
			return value.Name
		}
	case *UserInfoBuffer_AccountName:
		if value != nil {
			return value.AccountName
		}
	case *UserInfoBuffer_FullName:
		if value != nil {
			return value.FullName
		}
	case *UserInfoBuffer_PrimaryGroup:
		if value != nil {
			return value.PrimaryGroup
		}
	case *UserInfoBuffer_Home:
		if value != nil {
			return value.Home
		}
	case *UserInfoBuffer_Script:
		if value != nil {
			return value.Script
		}
	case *UserInfoBuffer_Profile:
		if value != nil {
			return value.Profile
		}
	case *UserInfoBuffer_AdminComment:
		if value != nil {
			return value.AdminComment
		}
	case *UserInfoBuffer_WorkStations:
		if value != nil {
			return value.WorkStations
		}
	case *UserInfoBuffer_Control:
		if value != nil {
			return value.Control
		}
	case *UserInfoBuffer_Expires:
		if value != nil {
			return value.Expires
		}
	case *UserInfoBuffer_Internal1:
		if value != nil {
			return value.Internal1
		}
	case *UserInfoBuffer_Parameters:
		if value != nil {
			return value.Parameters
		}
	case *UserInfoBuffer_All:
		if value != nil {
			return value.All
		}
	case *UserInfoBuffer_Internal4:
		if value != nil {
			return value.Internal4
		}
	case *UserInfoBuffer_Internal5:
		if value != nil {
			return value.Internal5
		}
	case *UserInfoBuffer_Internal4New:
		if value != nil {
			return value.Internal4New
		}
	case *UserInfoBuffer_Internal5New:
		if value != nil {
			return value.Internal5New
		}
	}
	return nil
}

type is_UserInfoBuffer interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_UserInfoBuffer()
}

func (o *UserInfoBuffer) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *UserInfoBuffer_General:
		return uint16(1)
	case *UserInfoBuffer_Preferences:
		return uint16(2)
	case *UserInfoBuffer_Logon:
		return uint16(3)
	case *UserInfoBuffer_LogonHours:
		return uint16(4)
	case *UserInfoBuffer_Account:
		return uint16(5)
	case *UserInfoBuffer_Name:
		return uint16(6)
	case *UserInfoBuffer_AccountName:
		return uint16(7)
	case *UserInfoBuffer_FullName:
		return uint16(8)
	case *UserInfoBuffer_PrimaryGroup:
		return uint16(9)
	case *UserInfoBuffer_Home:
		return uint16(10)
	case *UserInfoBuffer_Script:
		return uint16(11)
	case *UserInfoBuffer_Profile:
		return uint16(12)
	case *UserInfoBuffer_AdminComment:
		return uint16(13)
	case *UserInfoBuffer_WorkStations:
		return uint16(14)
	case *UserInfoBuffer_Control:
		return uint16(16)
	case *UserInfoBuffer_Expires:
		return uint16(17)
	case *UserInfoBuffer_Internal1:
		return uint16(18)
	case *UserInfoBuffer_Parameters:
		return uint16(20)
	case *UserInfoBuffer_All:
		return uint16(21)
	case *UserInfoBuffer_Internal4:
		return uint16(23)
	case *UserInfoBuffer_Internal5:
		return uint16(24)
	case *UserInfoBuffer_Internal4New:
		return uint16(25)
	case *UserInfoBuffer_Internal5New:
		return uint16(26)
	}
	return uint16(0)
}

func (o *UserInfoBuffer) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*UserInfoBuffer_General)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_General{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*UserInfoBuffer_Preferences)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Preferences{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*UserInfoBuffer_Logon)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Logon{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*UserInfoBuffer_LogonHours)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_LogonHours{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5):
		_o, _ := o.Value.(*UserInfoBuffer_Account)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Account{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(6):
		_o, _ := o.Value.(*UserInfoBuffer_Name)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Name{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(7):
		_o, _ := o.Value.(*UserInfoBuffer_AccountName)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_AccountName{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(8):
		_o, _ := o.Value.(*UserInfoBuffer_FullName)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_FullName{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(9):
		_o, _ := o.Value.(*UserInfoBuffer_PrimaryGroup)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_PrimaryGroup{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(10):
		_o, _ := o.Value.(*UserInfoBuffer_Home)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Home{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(11):
		_o, _ := o.Value.(*UserInfoBuffer_Script)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Script{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(12):
		_o, _ := o.Value.(*UserInfoBuffer_Profile)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Profile{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(13):
		_o, _ := o.Value.(*UserInfoBuffer_AdminComment)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_AdminComment{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(14):
		_o, _ := o.Value.(*UserInfoBuffer_WorkStations)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_WorkStations{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(16):
		_o, _ := o.Value.(*UserInfoBuffer_Control)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Control{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(17):
		_o, _ := o.Value.(*UserInfoBuffer_Expires)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Expires{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(18):
		_o, _ := o.Value.(*UserInfoBuffer_Internal1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Internal1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(20):
		_o, _ := o.Value.(*UserInfoBuffer_Parameters)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Parameters{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(21):
		_o, _ := o.Value.(*UserInfoBuffer_All)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_All{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(23):
		_o, _ := o.Value.(*UserInfoBuffer_Internal4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Internal4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(24):
		_o, _ := o.Value.(*UserInfoBuffer_Internal5)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Internal5{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(25):
		_o, _ := o.Value.(*UserInfoBuffer_Internal4New)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Internal4New{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(26):
		_o, _ := o.Value.(*UserInfoBuffer_Internal5New)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer_Internal5New{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *UserInfoBuffer) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &UserInfoBuffer_General{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &UserInfoBuffer_Preferences{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &UserInfoBuffer_Logon{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &UserInfoBuffer_LogonHours{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5):
		o.Value = &UserInfoBuffer_Account{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(6):
		o.Value = &UserInfoBuffer_Name{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(7):
		o.Value = &UserInfoBuffer_AccountName{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(8):
		o.Value = &UserInfoBuffer_FullName{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(9):
		o.Value = &UserInfoBuffer_PrimaryGroup{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(10):
		o.Value = &UserInfoBuffer_Home{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(11):
		o.Value = &UserInfoBuffer_Script{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(12):
		o.Value = &UserInfoBuffer_Profile{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(13):
		o.Value = &UserInfoBuffer_AdminComment{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(14):
		o.Value = &UserInfoBuffer_WorkStations{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(16):
		o.Value = &UserInfoBuffer_Control{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(17):
		o.Value = &UserInfoBuffer_Expires{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(18):
		o.Value = &UserInfoBuffer_Internal1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(20):
		o.Value = &UserInfoBuffer_Parameters{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(21):
		o.Value = &UserInfoBuffer_All{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(23):
		o.Value = &UserInfoBuffer_Internal4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(24):
		o.Value = &UserInfoBuffer_Internal5{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(25):
		o.Value = &UserInfoBuffer_Internal4New{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(26):
		o.Value = &UserInfoBuffer_Internal5New{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// UserInfoBuffer_General structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 1
type UserInfoBuffer_General struct {
	General *UserGeneralInformation `idl:"name:General" json:"general"`
}

func (*UserInfoBuffer_General) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_General) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.General != nil {
		if err := o.General.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserGeneralInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_General) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.General == nil {
		o.General = &UserGeneralInformation{}
	}
	if err := o.General.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Preferences structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 2
type UserInfoBuffer_Preferences struct {
	Preferences *UserPreferencesInformation `idl:"name:Preferences" json:"preferences"`
}

func (*UserInfoBuffer_Preferences) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Preferences) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Preferences != nil {
		if err := o.Preferences.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserPreferencesInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Preferences) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Preferences == nil {
		o.Preferences = &UserPreferencesInformation{}
	}
	if err := o.Preferences.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Logon structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 3
type UserInfoBuffer_Logon struct {
	Logon *UserLogonInformation `idl:"name:Logon" json:"logon"`
}

func (*UserInfoBuffer_Logon) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Logon) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Logon != nil {
		if err := o.Logon.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserLogonInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Logon) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Logon == nil {
		o.Logon = &UserLogonInformation{}
	}
	if err := o.Logon.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_LogonHours structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 4
type UserInfoBuffer_LogonHours struct {
	LogonHours *UserLogonHoursInformation `idl:"name:LogonHours" json:"logon_hours"`
}

func (*UserInfoBuffer_LogonHours) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_LogonHours) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LogonHours != nil {
		if err := o.LogonHours.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserLogonHoursInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_LogonHours) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LogonHours == nil {
		o.LogonHours = &UserLogonHoursInformation{}
	}
	if err := o.LogonHours.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Account structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 5
type UserInfoBuffer_Account struct {
	Account *UserAccountInformation `idl:"name:Account" json:"account"`
}

func (*UserInfoBuffer_Account) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Account) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Account != nil {
		if err := o.Account.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserAccountInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Account) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Account == nil {
		o.Account = &UserAccountInformation{}
	}
	if err := o.Account.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Name structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 6
type UserInfoBuffer_Name struct {
	Name *UserNameInformation `idl:"name:Name" json:"name"`
}

func (*UserInfoBuffer_Name) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Name) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserNameInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Name) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Name == nil {
		o.Name = &UserNameInformation{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_AccountName structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 7
type UserInfoBuffer_AccountName struct {
	AccountName *UserANameInformation `idl:"name:AccountName" json:"account_name"`
}

func (*UserInfoBuffer_AccountName) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_AccountName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AccountName != nil {
		if err := o.AccountName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserANameInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_AccountName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.AccountName == nil {
		o.AccountName = &UserANameInformation{}
	}
	if err := o.AccountName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_FullName structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 8
type UserInfoBuffer_FullName struct {
	FullName *UserFNameInformation `idl:"name:FullName" json:"full_name"`
}

func (*UserInfoBuffer_FullName) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_FullName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.FullName != nil {
		if err := o.FullName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserFNameInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_FullName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.FullName == nil {
		o.FullName = &UserFNameInformation{}
	}
	if err := o.FullName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_PrimaryGroup structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 9
type UserInfoBuffer_PrimaryGroup struct {
	PrimaryGroup *UserPrimaryGroupInformation `idl:"name:PrimaryGroup" json:"primary_group"`
}

func (*UserInfoBuffer_PrimaryGroup) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_PrimaryGroup) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PrimaryGroup != nil {
		if err := o.PrimaryGroup.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserPrimaryGroupInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_PrimaryGroup) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PrimaryGroup == nil {
		o.PrimaryGroup = &UserPrimaryGroupInformation{}
	}
	if err := o.PrimaryGroup.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Home structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 10
type UserInfoBuffer_Home struct {
	Home *UserHomeInformation `idl:"name:Home" json:"home"`
}

func (*UserInfoBuffer_Home) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Home) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Home != nil {
		if err := o.Home.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserHomeInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Home) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Home == nil {
		o.Home = &UserHomeInformation{}
	}
	if err := o.Home.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Script structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 11
type UserInfoBuffer_Script struct {
	Script *UserScriptInformation `idl:"name:Script" json:"script"`
}

func (*UserInfoBuffer_Script) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Script) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Script != nil {
		if err := o.Script.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserScriptInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Script) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Script == nil {
		o.Script = &UserScriptInformation{}
	}
	if err := o.Script.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Profile structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 12
type UserInfoBuffer_Profile struct {
	Profile *UserProfileInformation `idl:"name:Profile" json:"profile"`
}

func (*UserInfoBuffer_Profile) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Profile) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Profile != nil {
		if err := o.Profile.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserProfileInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Profile) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Profile == nil {
		o.Profile = &UserProfileInformation{}
	}
	if err := o.Profile.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_AdminComment structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 13
type UserInfoBuffer_AdminComment struct {
	AdminComment *UserAdminCommentInformation `idl:"name:AdminComment" json:"admin_comment"`
}

func (*UserInfoBuffer_AdminComment) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_AdminComment) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.AdminComment != nil {
		if err := o.AdminComment.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserAdminCommentInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_AdminComment) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.AdminComment == nil {
		o.AdminComment = &UserAdminCommentInformation{}
	}
	if err := o.AdminComment.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_WorkStations structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 14
type UserInfoBuffer_WorkStations struct {
	WorkStations *UserWorkstationsInformation `idl:"name:WorkStations" json:"work_stations"`
}

func (*UserInfoBuffer_WorkStations) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_WorkStations) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.WorkStations != nil {
		if err := o.WorkStations.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserWorkstationsInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_WorkStations) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.WorkStations == nil {
		o.WorkStations = &UserWorkstationsInformation{}
	}
	if err := o.WorkStations.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Control structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 16
type UserInfoBuffer_Control struct {
	Control *UserControlInformation `idl:"name:Control" json:"control"`
}

func (*UserInfoBuffer_Control) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Control) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Control != nil {
		if err := o.Control.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserControlInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Control) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Control == nil {
		o.Control = &UserControlInformation{}
	}
	if err := o.Control.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Expires structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 17
type UserInfoBuffer_Expires struct {
	Expires *UserExpiresInformation `idl:"name:Expires" json:"expires"`
}

func (*UserInfoBuffer_Expires) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Expires) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Expires != nil {
		if err := o.Expires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserExpiresInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Expires) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Expires == nil {
		o.Expires = &UserExpiresInformation{}
	}
	if err := o.Expires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Internal1 structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 18
type UserInfoBuffer_Internal1 struct {
	Internal1 *UserInternal1Information `idl:"name:Internal1" json:"internal1"`
}

func (*UserInfoBuffer_Internal1) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Internal1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Internal1 != nil {
		if err := o.Internal1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserInternal1Information{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Internal1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Internal1 == nil {
		o.Internal1 = &UserInternal1Information{}
	}
	if err := o.Internal1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Parameters structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 20
type UserInfoBuffer_Parameters struct {
	Parameters *UserParametersInformation `idl:"name:Parameters" json:"parameters"`
}

func (*UserInfoBuffer_Parameters) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Parameters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Parameters != nil {
		if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserParametersInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Parameters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Parameters == nil {
		o.Parameters = &UserParametersInformation{}
	}
	if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_All structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 21
type UserInfoBuffer_All struct {
	All *UserAllInformation `idl:"name:All" json:"all"`
}

func (*UserInfoBuffer_All) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_All) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.All != nil {
		if err := o.All.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserAllInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_All) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.All == nil {
		o.All = &UserAllInformation{}
	}
	if err := o.All.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Internal4 structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 23
type UserInfoBuffer_Internal4 struct {
	Internal4 *UserInternal4Information `idl:"name:Internal4" json:"internal4"`
}

func (*UserInfoBuffer_Internal4) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Internal4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Internal4 != nil {
		if err := o.Internal4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserInternal4Information{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Internal4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Internal4 == nil {
		o.Internal4 = &UserInternal4Information{}
	}
	if err := o.Internal4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Internal5 structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 24
type UserInfoBuffer_Internal5 struct {
	Internal5 *UserInternal5Information `idl:"name:Internal5" json:"internal5"`
}

func (*UserInfoBuffer_Internal5) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Internal5) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Internal5 != nil {
		if err := o.Internal5.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserInternal5Information{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Internal5) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Internal5 == nil {
		o.Internal5 = &UserInternal5Information{}
	}
	if err := o.Internal5.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Internal4New structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 25
type UserInfoBuffer_Internal4New struct {
	Internal4New *UserInternal4InformationNew `idl:"name:Internal4New" json:"internal4_new"`
}

func (*UserInfoBuffer_Internal4New) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Internal4New) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Internal4New != nil {
		if err := o.Internal4New.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserInternal4InformationNew{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Internal4New) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Internal4New == nil {
		o.Internal4New = &UserInternal4InformationNew{}
	}
	if err := o.Internal4New.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// UserInfoBuffer_Internal5New structure represents SAMPR_USER_INFO_BUFFER RPC union arm.
//
// It has following labels: 26
type UserInfoBuffer_Internal5New struct {
	Internal5New *UserInternal5InformationNew `idl:"name:Internal5New" json:"internal5_new"`
}

func (*UserInfoBuffer_Internal5New) is_UserInfoBuffer() {}

func (o *UserInfoBuffer_Internal5New) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Internal5New != nil {
		if err := o.Internal5New.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserInternal5InformationNew{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserInfoBuffer_Internal5New) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Internal5New == nil {
		o.Internal5New = &UserInternal5InformationNew{}
	}
	if err := o.Internal5New.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PasswordPolicyValidationType type represents PASSWORD_POLICY_VALIDATION_TYPE RPC enumeration.
//
// The PASSWORD_POLICY_VALIDATION_TYPE enumeration indicates the type of policy validation
// that is being requested.
type PasswordPolicyValidationType uint16

var (
	// SamValidateAuthentication:  Indicates a request to execute the password policy validation
	// performed at logon.
	PasswordPolicyValidationTypeSAMValidateAuthentication PasswordPolicyValidationType = 1
	// SamValidatePasswordChange:  Indicates a request to execute the password policy validation
	// performed during a password change request.
	PasswordPolicyValidationTypeSAMValidatePasswordChange PasswordPolicyValidationType = 2
	// SamValidatePasswordReset:  Indicates a request to execute the password policy validation
	// performed during a password reset.
	PasswordPolicyValidationTypeSAMValidatePasswordReset PasswordPolicyValidationType = 3
)

func (o PasswordPolicyValidationType) String() string {
	switch o {
	case PasswordPolicyValidationTypeSAMValidateAuthentication:
		return "PasswordPolicyValidationTypeSAMValidateAuthentication"
	case PasswordPolicyValidationTypeSAMValidatePasswordChange:
		return "PasswordPolicyValidationTypeSAMValidatePasswordChange"
	case PasswordPolicyValidationTypeSAMValidatePasswordReset:
		return "PasswordPolicyValidationTypeSAMValidatePasswordReset"
	}
	return "Invalid"
}

// SAMValidatePasswordHash structure represents SAM_VALIDATE_PASSWORD_HASH RPC structure.
//
// The SAM_VALIDATE_PASSWORD_HASH structure holds a binary value that represents a cryptographic
// hash.
type SAMValidatePasswordHash struct {
	// Length:  The size, in bytes, of Hash. If zero, Hash MUST be ignored.
	Length uint32 `idl:"name:Length" json:"length"`
	// Hash:  A binary value.
	Hash []byte `idl:"name:Hash;size_is:(Length);pointer:unique" json:"hash"`
}

func (o *SAMValidatePasswordHash) xxx_PreparePayload(ctx context.Context) error {
	if o.Hash != nil && o.Length == 0 {
		o.Length = uint32(len(o.Hash))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidatePasswordHash) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.Hash != nil || o.Length > 0 {
		_ptr_Hash := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Hash {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Hash[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Hash); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Hash, _ptr_Hash); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidatePasswordHash) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_Hash := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			return fmt.Errorf("buffer overflow for size %d of array o.Hash", sizeInfo[0])
		}
		o.Hash = make([]byte, sizeInfo[0])
		for i1 := range o.Hash {
			i1 := i1
			if err := w.ReadData(&o.Hash[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Hash := func(ptr interface{}) { o.Hash = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Hash, _s_Hash, _ptr_Hash); err != nil {
		return err
	}
	return nil
}

// SAMValidatePersistedFields structure represents SAM_VALIDATE_PERSISTED_FIELDS RPC structure.
//
// The SAM_VALIDATE_PERSISTED_FIELDS structure holds various characteristics about password
// state.
type SAMValidatePersistedFields struct {
	// PresentFields:  A bitmask to indicate which of the fields are valid. The following
	// table shows the defined values. If a bit is set, the corresponding field is valid;
	// if a bit is not set, the field is not valid.
	//
	//	+-------------------------------------------------+-----------------------+
	//	|                                                 |                       |
	//	|                      VALUE                      |        MEANING        |
	//	|                                                 |                       |
	//	+-------------------------------------------------+-----------------------+
	//	+-------------------------------------------------+-----------------------+
	//	| SAM_VALIDATE_PASSWORD_LAST_SET 0x00000001       | PasswordLastSet       |
	//	+-------------------------------------------------+-----------------------+
	//	| SAM_VALIDATE_BAD_PASSWORD_TIME 0x00000002       | BadPasswordTime       |
	//	+-------------------------------------------------+-----------------------+
	//	| SAM_VALIDATE_LOCKOUT_TIME 0x00000004            | LockoutTime           |
	//	+-------------------------------------------------+-----------------------+
	//	| SAM_VALIDATE_BAD_PASSWORD_COUNT 0x00000008      | BadPasswordCount      |
	//	+-------------------------------------------------+-----------------------+
	//	| SAM_VALIDATE_PASSWORD_HISTORY_LENGTH 0x00000010 | PasswordHistoryLength |
	//	+-------------------------------------------------+-----------------------+
	//	| SAM_VALIDATE_PASSWORD_HISTORY 0x00000020        | PasswordHistory       |
	//	+-------------------------------------------------+-----------------------+
	PresentFields uint32 `idl:"name:PresentFields" json:"present_fields"`
	// PasswordLastSet:  This field represents the time at which the password was last reset
	// or changed. It uses FILETIME syntax.
	PasswordLastSet *dtyp.LargeInteger `idl:"name:PasswordLastSet" json:"password_last_set"`
	// BadPasswordTime:  This field represents the time at which an invalid password was
	// presented to either a password change request or an authentication request. It uses
	// FILETIME syntax.
	BadPasswordTime *dtyp.LargeInteger `idl:"name:BadPasswordTime" json:"bad_password_time"`
	// LockoutTime:  This field represents the time at which the owner of the password data
	// was locked out. It uses FILETIME syntax.
	LockoutTime *dtyp.LargeInteger `idl:"name:LockoutTime" json:"lockout_time"`
	// BadPasswordCount:  Indicates how many invalid passwords have accumulated (see message
	// processing for details).
	BadPasswordCount uint32 `idl:"name:BadPasswordCount" json:"bad_password_count"`
	// PasswordHistoryLength:  Indicates how many previous passwords are in the PasswordHistory
	// field.
	PasswordHistoryLength uint32 `idl:"name:PasswordHistoryLength" json:"password_history_length"`
	// PasswordHistory:  An array of hash values representing the previous PasswordHistoryLength
	// passwords.
	PasswordHistory []*SAMValidatePasswordHash `idl:"name:PasswordHistory;size_is:(PasswordHistoryLength);pointer:unique" json:"password_history"`
}

func (o *SAMValidatePersistedFields) xxx_PreparePayload(ctx context.Context) error {
	if o.PasswordHistory != nil && o.PasswordHistoryLength == 0 {
		o.PasswordHistoryLength = uint32(len(o.PasswordHistory))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidatePersistedFields) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.PresentFields); err != nil {
		return err
	}
	if o.PasswordLastSet != nil {
		if err := o.PasswordLastSet.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.BadPasswordTime != nil {
		if err := o.BadPasswordTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LockoutTime != nil {
		if err := o.LockoutTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.WriteData(o.PasswordHistoryLength); err != nil {
		return err
	}
	if o.PasswordHistory != nil || o.PasswordHistoryLength > 0 {
		_ptr_PasswordHistory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.PasswordHistoryLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.PasswordHistory {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.PasswordHistory[i1] != nil {
					if err := o.PasswordHistory[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SAMValidatePasswordHash{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.PasswordHistory); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SAMValidatePasswordHash{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PasswordHistory, _ptr_PasswordHistory); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidatePersistedFields) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.PresentFields); err != nil {
		return err
	}
	if o.PasswordLastSet == nil {
		o.PasswordLastSet = &dtyp.LargeInteger{}
	}
	if err := o.PasswordLastSet.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.BadPasswordTime == nil {
		o.BadPasswordTime = &dtyp.LargeInteger{}
	}
	if err := o.BadPasswordTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LockoutTime == nil {
		o.LockoutTime = &dtyp.LargeInteger{}
	}
	if err := o.LockoutTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordHistoryLength); err != nil {
		return err
	}
	_ptr_PasswordHistory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.PasswordHistoryLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.PasswordHistoryLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.PasswordHistory", sizeInfo[0])
		}
		o.PasswordHistory = make([]*SAMValidatePasswordHash, sizeInfo[0])
		for i1 := range o.PasswordHistory {
			i1 := i1
			if o.PasswordHistory[i1] == nil {
				o.PasswordHistory[i1] = &SAMValidatePasswordHash{}
			}
			if err := o.PasswordHistory[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_PasswordHistory := func(ptr interface{}) { o.PasswordHistory = *ptr.(*[]*SAMValidatePasswordHash) }
	if err := w.ReadPointer(&o.PasswordHistory, _s_PasswordHistory, _ptr_PasswordHistory); err != nil {
		return err
	}
	return nil
}

// SAMValidateValidationStatus type represents SAM_VALIDATE_VALIDATION_STATUS RPC enumeration.
//
// The SAM_VALIDATE_VALIDATION_STATUS enumeration defines policy evaluation outcomes.
type SAMValidateValidationStatus uint16

var (
	// SamValidateSuccess:  Password validation succeeded.
	SAMValidateValidationStatusSuccess SAMValidateValidationStatus = 0
	// SamValidatePasswordMustChange:  The password must be changed.
	SAMValidateValidationStatusPasswordMustChange SAMValidateValidationStatus = 1
	// SamValidateAccountLockedOut:  The account is locked out.
	SAMValidateValidationStatusAccountLockedOut SAMValidateValidationStatus = 2
	// SamValidatePasswordExpired:  The password has expired.
	SAMValidateValidationStatusPasswordExpired SAMValidateValidationStatus = 3
	// SamValidatePasswordIncorrect:  The password is incorrect.
	SAMValidateValidationStatusPasswordIncorrect SAMValidateValidationStatus = 4
	// SamValidatePasswordIsInHistory:  The password is in the password history.
	SAMValidateValidationStatusPasswordIsInHistory SAMValidateValidationStatus = 5
	// SamValidatePasswordTooShort:  The password is too short.
	SAMValidateValidationStatusPasswordTooShort SAMValidateValidationStatus = 6
	// SamValidatePasswordTooLong:  The password is too long.
	SAMValidateValidationStatusPasswordTooLong SAMValidateValidationStatus = 7
	// SamValidatePasswordNotComplexEnough:  The password is not complex enough.
	SAMValidateValidationStatusPasswordNotComplexEnough SAMValidateValidationStatus = 8
	// SamValidatePasswordTooRecent:  The password was changed recently.
	SAMValidateValidationStatusPasswordTooRecent SAMValidateValidationStatus = 9
	// SamValidatePasswordFilterError:  The password filter failed to validate the password.
	//
	// See the message processing of SamrValidatePassword (section 3.1.5.13.7) for the semantic
	// meanings of the enumeration values.
	SAMValidateValidationStatusPasswordFilterError SAMValidateValidationStatus = 10
)

func (o SAMValidateValidationStatus) String() string {
	switch o {
	case SAMValidateValidationStatusSuccess:
		return "SAMValidateValidationStatusSuccess"
	case SAMValidateValidationStatusPasswordMustChange:
		return "SAMValidateValidationStatusPasswordMustChange"
	case SAMValidateValidationStatusAccountLockedOut:
		return "SAMValidateValidationStatusAccountLockedOut"
	case SAMValidateValidationStatusPasswordExpired:
		return "SAMValidateValidationStatusPasswordExpired"
	case SAMValidateValidationStatusPasswordIncorrect:
		return "SAMValidateValidationStatusPasswordIncorrect"
	case SAMValidateValidationStatusPasswordIsInHistory:
		return "SAMValidateValidationStatusPasswordIsInHistory"
	case SAMValidateValidationStatusPasswordTooShort:
		return "SAMValidateValidationStatusPasswordTooShort"
	case SAMValidateValidationStatusPasswordTooLong:
		return "SAMValidateValidationStatusPasswordTooLong"
	case SAMValidateValidationStatusPasswordNotComplexEnough:
		return "SAMValidateValidationStatusPasswordNotComplexEnough"
	case SAMValidateValidationStatusPasswordTooRecent:
		return "SAMValidateValidationStatusPasswordTooRecent"
	case SAMValidateValidationStatusPasswordFilterError:
		return "SAMValidateValidationStatusPasswordFilterError"
	}
	return "Invalid"
}

// SAMValidateStandardOutputArg structure represents SAM_VALIDATE_STANDARD_OUTPUT_ARG RPC structure.
//
// The SAM_VALIDATE_STANDARD_OUTPUT_ARG structure holds the output of SamrValidatePassword.
type SAMValidateStandardOutputArg struct {
	// ChangedPersistedFields:  The password state that has changed. See section 2.2.9.2.
	ChangedPersistedFields *SAMValidatePersistedFields `idl:"name:ChangedPersistedFields" json:"changed_persisted_fields"`
	// ValidationStatus:  The result of the policy evaluation. See section 2.2.9.3.
	ValidationStatus SAMValidateValidationStatus `idl:"name:ValidationStatus" json:"validation_status"`
}

func (o *SAMValidateStandardOutputArg) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidateStandardOutputArg) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ChangedPersistedFields != nil {
		if err := o.ChangedPersistedFields.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidatePersistedFields{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.ValidationStatus)); err != nil {
		return err
	}
	return nil
}
func (o *SAMValidateStandardOutputArg) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ChangedPersistedFields == nil {
		o.ChangedPersistedFields = &SAMValidatePersistedFields{}
	}
	if err := o.ChangedPersistedFields.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ValidationStatus)); err != nil {
		return err
	}
	return nil
}

// SAMValidateAuthenticationInputArg structure represents SAM_VALIDATE_AUTHENTICATION_INPUT_ARG RPC structure.
//
// The SAM_VALIDATE_AUTHENTICATION_INPUT_ARG structure holds information about an authentication
// request.
type SAMValidateAuthenticationInputArg struct {
	// InputPersistedFields:  Password state.
	InputPersistedFields *SAMValidatePersistedFields `idl:"name:InputPersistedFields" json:"input_persisted_fields"`
	// PasswordMatched:  A nonzero value indicates that a valid password was presented to
	// the change-password request.
	PasswordMatched uint8 `idl:"name:PasswordMatched" json:"password_matched"`
}

func (o *SAMValidateAuthenticationInputArg) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidateAuthenticationInputArg) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.InputPersistedFields != nil {
		if err := o.InputPersistedFields.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidatePersistedFields{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PasswordMatched); err != nil {
		return err
	}
	return nil
}
func (o *SAMValidateAuthenticationInputArg) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.InputPersistedFields == nil {
		o.InputPersistedFields = &SAMValidatePersistedFields{}
	}
	if err := o.InputPersistedFields.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordMatched); err != nil {
		return err
	}
	return nil
}

// SAMValidatePasswordChangeInputArg structure represents SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG RPC structure.
//
// The SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG structure holds information about a password
// change request.
type SAMValidatePasswordChangeInputArg struct {
	// InputPersistedFields:  Password state. See section 2.2.9.2.
	InputPersistedFields *SAMValidatePersistedFields `idl:"name:InputPersistedFields" json:"input_persisted_fields"`
	// ClearPassword:  The cleartext password of the change-password operation.
	ClearPassword *dtyp.UnicodeString `idl:"name:ClearPassword" json:"clear_password"`
	// UserAccountName:  The application-specific logon name of an account performing the
	// change-password operation.
	UserAccountName *dtyp.UnicodeString `idl:"name:UserAccountName" json:"user_account_name"`
	// HashedPassword:  A binary value containing a hashed form of the value contained in
	// the ClearPassword field. The structure of this binary value is specified in section
	// 2.2.9.1. The hash function used to generate this value is chosen by the client. An
	// example hash function might be MD5 (as specified in [RFC1321]). The server implementation
	// is independent of that choice; that is, through this protocol, the server is exposed
	// to a sequence of bytes formatted per section 2.2.9.1 and is, therefore, not exposed
	// to the hash function chosen by the client. Furthermore, there is no processing by
	// the server that requires knowledge of the specific hash function chosen. Section
	// 2.2.9 contains more information about a scenario in which this field is used.
	HashedPassword *SAMValidatePasswordHash `idl:"name:HashedPassword" json:"hashed_password"`
	// PasswordMatch:  A nonzero value indicates that a valid password was presented to
	// the change-password request.
	PasswordMatch uint8 `idl:"name:PasswordMatch" json:"password_match"`
}

func (o *SAMValidatePasswordChangeInputArg) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidatePasswordChangeInputArg) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.InputPersistedFields != nil {
		if err := o.InputPersistedFields.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidatePersistedFields{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClearPassword != nil {
		if err := o.ClearPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.UserAccountName != nil {
		if err := o.UserAccountName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HashedPassword != nil {
		if err := o.HashedPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidatePasswordHash{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PasswordMatch); err != nil {
		return err
	}
	return nil
}
func (o *SAMValidatePasswordChangeInputArg) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.InputPersistedFields == nil {
		o.InputPersistedFields = &SAMValidatePersistedFields{}
	}
	if err := o.InputPersistedFields.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ClearPassword == nil {
		o.ClearPassword = &dtyp.UnicodeString{}
	}
	if err := o.ClearPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.UserAccountName == nil {
		o.UserAccountName = &dtyp.UnicodeString{}
	}
	if err := o.UserAccountName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HashedPassword == nil {
		o.HashedPassword = &SAMValidatePasswordHash{}
	}
	if err := o.HashedPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordMatch); err != nil {
		return err
	}
	return nil
}

// SAMValidatePasswordResetInputArg structure represents SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG RPC structure.
//
// The SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG structure holds various information about
// a password reset request.
type SAMValidatePasswordResetInputArg struct {
	// InputPersistedFields:  Password state. See section 2.2.9.2.
	InputPersistedFields *SAMValidatePersistedFields `idl:"name:InputPersistedFields" json:"input_persisted_fields"`
	// ClearPassword:  The cleartext password of the reset-password operation.
	ClearPassword *dtyp.UnicodeString `idl:"name:ClearPassword" json:"clear_password"`
	// UserAccountName:  The application-specific logon name of the account performing the
	// reset-password operation.
	UserAccountName *dtyp.UnicodeString `idl:"name:UserAccountName" json:"user_account_name"`
	// HashedPassword:  See the specification for SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG
	// (section 2.2.9.6) for the field with the same name.
	HashedPassword *SAMValidatePasswordHash `idl:"name:HashedPassword" json:"hashed_password"`
	// PasswordMustChangeAtNextLogon:  Nonzero indicates that a password change MUST occur
	// before an authentication request can succeed.
	PasswordMustChangeATNextLogon uint8 `idl:"name:PasswordMustChangeAtNextLogon" json:"password_must_change_at_next_logon"`
	// ClearLockout:  Nonzero indicates that the lockout state is to be reset.
	ClearLockout uint8 `idl:"name:ClearLockout" json:"clear_lockout"`
}

func (o *SAMValidatePasswordResetInputArg) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidatePasswordResetInputArg) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.InputPersistedFields != nil {
		if err := o.InputPersistedFields.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidatePersistedFields{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClearPassword != nil {
		if err := o.ClearPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.UserAccountName != nil {
		if err := o.UserAccountName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HashedPassword != nil {
		if err := o.HashedPassword.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidatePasswordHash{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PasswordMustChangeATNextLogon); err != nil {
		return err
	}
	if err := w.WriteData(o.ClearLockout); err != nil {
		return err
	}
	return nil
}
func (o *SAMValidatePasswordResetInputArg) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.InputPersistedFields == nil {
		o.InputPersistedFields = &SAMValidatePersistedFields{}
	}
	if err := o.InputPersistedFields.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ClearPassword == nil {
		o.ClearPassword = &dtyp.UnicodeString{}
	}
	if err := o.ClearPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.UserAccountName == nil {
		o.UserAccountName = &dtyp.UnicodeString{}
	}
	if err := o.UserAccountName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HashedPassword == nil {
		o.HashedPassword = &SAMValidatePasswordHash{}
	}
	if err := o.HashedPassword.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PasswordMustChangeATNextLogon); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClearLockout); err != nil {
		return err
	}
	return nil
}

// SAMValidateInputArg structure represents SAM_VALIDATE_INPUT_ARG RPC union.
type SAMValidateInputArg struct {
	// Types that are assignable to Value
	//
	// *SAMValidateInputArg_ValidateAuthenticationInput
	// *SAMValidateInputArg_ValidatePasswordChangeInput
	// *SAMValidateInputArg_ValidatePasswordResetInput
	Value is_SAMValidateInputArg `json:"value"`
}

func (o *SAMValidateInputArg) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SAMValidateInputArg_ValidateAuthenticationInput:
		if value != nil {
			return value.ValidateAuthenticationInput
		}
	case *SAMValidateInputArg_ValidatePasswordChangeInput:
		if value != nil {
			return value.ValidatePasswordChangeInput
		}
	case *SAMValidateInputArg_ValidatePasswordResetInput:
		if value != nil {
			return value.ValidatePasswordResetInput
		}
	}
	return nil
}

type is_SAMValidateInputArg interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SAMValidateInputArg()
}

func (o *SAMValidateInputArg) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SAMValidateInputArg_ValidateAuthenticationInput:
		return uint16(1)
	case *SAMValidateInputArg_ValidatePasswordChangeInput:
		return uint16(2)
	case *SAMValidateInputArg_ValidatePasswordResetInput:
		return uint16(3)
	}
	return uint16(0)
}

func (o *SAMValidateInputArg) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*SAMValidateInputArg_ValidateAuthenticationInput)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SAMValidateInputArg_ValidateAuthenticationInput{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*SAMValidateInputArg_ValidatePasswordChangeInput)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SAMValidateInputArg_ValidatePasswordChangeInput{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*SAMValidateInputArg_ValidatePasswordResetInput)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SAMValidateInputArg_ValidatePasswordResetInput{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SAMValidateInputArg) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &SAMValidateInputArg_ValidateAuthenticationInput{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &SAMValidateInputArg_ValidatePasswordChangeInput{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &SAMValidateInputArg_ValidatePasswordResetInput{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SAMValidateInputArg_ValidateAuthenticationInput structure represents SAM_VALIDATE_INPUT_ARG RPC union arm.
//
// It has following labels: 1
type SAMValidateInputArg_ValidateAuthenticationInput struct {
	ValidateAuthenticationInput *SAMValidateAuthenticationInputArg `idl:"name:ValidateAuthenticationInput" json:"validate_authentication_input"`
}

func (*SAMValidateInputArg_ValidateAuthenticationInput) is_SAMValidateInputArg() {}

func (o *SAMValidateInputArg_ValidateAuthenticationInput) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ValidateAuthenticationInput != nil {
		if err := o.ValidateAuthenticationInput.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidateAuthenticationInputArg{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidateInputArg_ValidateAuthenticationInput) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ValidateAuthenticationInput == nil {
		o.ValidateAuthenticationInput = &SAMValidateAuthenticationInputArg{}
	}
	if err := o.ValidateAuthenticationInput.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SAMValidateInputArg_ValidatePasswordChangeInput structure represents SAM_VALIDATE_INPUT_ARG RPC union arm.
//
// It has following labels: 2
type SAMValidateInputArg_ValidatePasswordChangeInput struct {
	ValidatePasswordChangeInput *SAMValidatePasswordChangeInputArg `idl:"name:ValidatePasswordChangeInput" json:"validate_password_change_input"`
}

func (*SAMValidateInputArg_ValidatePasswordChangeInput) is_SAMValidateInputArg() {}

func (o *SAMValidateInputArg_ValidatePasswordChangeInput) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ValidatePasswordChangeInput != nil {
		if err := o.ValidatePasswordChangeInput.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidatePasswordChangeInputArg{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidateInputArg_ValidatePasswordChangeInput) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ValidatePasswordChangeInput == nil {
		o.ValidatePasswordChangeInput = &SAMValidatePasswordChangeInputArg{}
	}
	if err := o.ValidatePasswordChangeInput.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SAMValidateInputArg_ValidatePasswordResetInput structure represents SAM_VALIDATE_INPUT_ARG RPC union arm.
//
// It has following labels: 3
type SAMValidateInputArg_ValidatePasswordResetInput struct {
	ValidatePasswordResetInput *SAMValidatePasswordResetInputArg `idl:"name:ValidatePasswordResetInput" json:"validate_password_reset_input"`
}

func (*SAMValidateInputArg_ValidatePasswordResetInput) is_SAMValidateInputArg() {}

func (o *SAMValidateInputArg_ValidatePasswordResetInput) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ValidatePasswordResetInput != nil {
		if err := o.ValidatePasswordResetInput.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidatePasswordResetInputArg{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidateInputArg_ValidatePasswordResetInput) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ValidatePasswordResetInput == nil {
		o.ValidatePasswordResetInput = &SAMValidatePasswordResetInputArg{}
	}
	if err := o.ValidatePasswordResetInput.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SAMValidateOutputArg structure represents SAM_VALIDATE_OUTPUT_ARG RPC union.
type SAMValidateOutputArg struct {
	// Types that are assignable to Value
	//
	// *SAMValidateOutputArg_ValidateAuthenticationOutput
	// *SAMValidateOutputArg_ValidatePasswordChangeOutput
	// *SAMValidateOutputArg_ValidatePasswordResetOutput
	Value is_SAMValidateOutputArg `json:"value"`
}

func (o *SAMValidateOutputArg) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SAMValidateOutputArg_ValidateAuthenticationOutput:
		if value != nil {
			return value.ValidateAuthenticationOutput
		}
	case *SAMValidateOutputArg_ValidatePasswordChangeOutput:
		if value != nil {
			return value.ValidatePasswordChangeOutput
		}
	case *SAMValidateOutputArg_ValidatePasswordResetOutput:
		if value != nil {
			return value.ValidatePasswordResetOutput
		}
	}
	return nil
}

type is_SAMValidateOutputArg interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SAMValidateOutputArg()
}

func (o *SAMValidateOutputArg) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SAMValidateOutputArg_ValidateAuthenticationOutput:
		return uint16(1)
	case *SAMValidateOutputArg_ValidatePasswordChangeOutput:
		return uint16(2)
	case *SAMValidateOutputArg_ValidatePasswordResetOutput:
		return uint16(3)
	}
	return uint16(0)
}

func (o *SAMValidateOutputArg) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*SAMValidateOutputArg_ValidateAuthenticationOutput)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SAMValidateOutputArg_ValidateAuthenticationOutput{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*SAMValidateOutputArg_ValidatePasswordChangeOutput)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SAMValidateOutputArg_ValidatePasswordChangeOutput{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*SAMValidateOutputArg_ValidatePasswordResetOutput)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SAMValidateOutputArg_ValidatePasswordResetOutput{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SAMValidateOutputArg) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &SAMValidateOutputArg_ValidateAuthenticationOutput{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &SAMValidateOutputArg_ValidatePasswordChangeOutput{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &SAMValidateOutputArg_ValidatePasswordResetOutput{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SAMValidateOutputArg_ValidateAuthenticationOutput structure represents SAM_VALIDATE_OUTPUT_ARG RPC union arm.
//
// It has following labels: 1
type SAMValidateOutputArg_ValidateAuthenticationOutput struct {
	ValidateAuthenticationOutput *SAMValidateStandardOutputArg `idl:"name:ValidateAuthenticationOutput" json:"validate_authentication_output"`
}

func (*SAMValidateOutputArg_ValidateAuthenticationOutput) is_SAMValidateOutputArg() {}

func (o *SAMValidateOutputArg_ValidateAuthenticationOutput) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ValidateAuthenticationOutput != nil {
		if err := o.ValidateAuthenticationOutput.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidateStandardOutputArg{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidateOutputArg_ValidateAuthenticationOutput) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ValidateAuthenticationOutput == nil {
		o.ValidateAuthenticationOutput = &SAMValidateStandardOutputArg{}
	}
	if err := o.ValidateAuthenticationOutput.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SAMValidateOutputArg_ValidatePasswordChangeOutput structure represents SAM_VALIDATE_OUTPUT_ARG RPC union arm.
//
// It has following labels: 2
type SAMValidateOutputArg_ValidatePasswordChangeOutput struct {
	ValidatePasswordChangeOutput *SAMValidateStandardOutputArg `idl:"name:ValidatePasswordChangeOutput" json:"validate_password_change_output"`
}

func (*SAMValidateOutputArg_ValidatePasswordChangeOutput) is_SAMValidateOutputArg() {}

func (o *SAMValidateOutputArg_ValidatePasswordChangeOutput) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ValidatePasswordChangeOutput != nil {
		if err := o.ValidatePasswordChangeOutput.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidateStandardOutputArg{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidateOutputArg_ValidatePasswordChangeOutput) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ValidatePasswordChangeOutput == nil {
		o.ValidatePasswordChangeOutput = &SAMValidateStandardOutputArg{}
	}
	if err := o.ValidatePasswordChangeOutput.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SAMValidateOutputArg_ValidatePasswordResetOutput structure represents SAM_VALIDATE_OUTPUT_ARG RPC union arm.
//
// It has following labels: 3
type SAMValidateOutputArg_ValidatePasswordResetOutput struct {
	ValidatePasswordResetOutput *SAMValidateStandardOutputArg `idl:"name:ValidatePasswordResetOutput" json:"validate_password_reset_output"`
}

func (*SAMValidateOutputArg_ValidatePasswordResetOutput) is_SAMValidateOutputArg() {}

func (o *SAMValidateOutputArg_ValidatePasswordResetOutput) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ValidatePasswordResetOutput != nil {
		if err := o.ValidatePasswordResetOutput.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SAMValidateStandardOutputArg{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SAMValidateOutputArg_ValidatePasswordResetOutput) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ValidatePasswordResetOutput == nil {
		o.ValidatePasswordResetOutput = &SAMValidateStandardOutputArg{}
	}
	if err := o.ValidatePasswordResetOutput.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultSamrClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultSamrClient) Connect(ctx context.Context, in *ConnectRequest, opts ...dcerpc.CallOption) (*ConnectResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) CloseHandle(ctx context.Context, in *CloseHandleRequest, opts ...dcerpc.CallOption) (*CloseHandleResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseHandleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) SetSecurityObject(ctx context.Context, in *SetSecurityObjectRequest, opts ...dcerpc.CallOption) (*SetSecurityObjectResponse, error) {
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

func (o *xxx_DefaultSamrClient) QuerySecurityObject(ctx context.Context, in *QuerySecurityObjectRequest, opts ...dcerpc.CallOption) (*QuerySecurityObjectResponse, error) {
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

func (o *xxx_DefaultSamrClient) LookupDomainInSAMServer(ctx context.Context, in *LookupDomainInSAMServerRequest, opts ...dcerpc.CallOption) (*LookupDomainInSAMServerResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupDomainInSAMServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) EnumerateDomainsInSAMServer(ctx context.Context, in *EnumerateDomainsInSAMServerRequest, opts ...dcerpc.CallOption) (*EnumerateDomainsInSAMServerResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateDomainsInSAMServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) OpenDomain(ctx context.Context, in *OpenDomainRequest, opts ...dcerpc.CallOption) (*OpenDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) QueryInformationDomain(ctx context.Context, in *QueryInformationDomainRequest, opts ...dcerpc.CallOption) (*QueryInformationDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryInformationDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) SetInformationDomain(ctx context.Context, in *SetInformationDomainRequest, opts ...dcerpc.CallOption) (*SetInformationDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInformationDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) CreateGroupInDomain(ctx context.Context, in *CreateGroupInDomainRequest, opts ...dcerpc.CallOption) (*CreateGroupInDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateGroupInDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) EnumerateGroupsInDomain(ctx context.Context, in *EnumerateGroupsInDomainRequest, opts ...dcerpc.CallOption) (*EnumerateGroupsInDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateGroupsInDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) CreateUserInDomain(ctx context.Context, in *CreateUserInDomainRequest, opts ...dcerpc.CallOption) (*CreateUserInDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateUserInDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) EnumerateUsersInDomain(ctx context.Context, in *EnumerateUsersInDomainRequest, opts ...dcerpc.CallOption) (*EnumerateUsersInDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateUsersInDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) CreateAliasInDomain(ctx context.Context, in *CreateAliasInDomainRequest, opts ...dcerpc.CallOption) (*CreateAliasInDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateAliasInDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) EnumerateAliasesInDomain(ctx context.Context, in *EnumerateAliasesInDomainRequest, opts ...dcerpc.CallOption) (*EnumerateAliasesInDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateAliasesInDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) GetAliasMembership(ctx context.Context, in *GetAliasMembershipRequest, opts ...dcerpc.CallOption) (*GetAliasMembershipResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetAliasMembershipResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) LookupNamesInDomain(ctx context.Context, in *LookupNamesInDomainRequest, opts ...dcerpc.CallOption) (*LookupNamesInDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupNamesInDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) LookupIDsInDomain(ctx context.Context, in *LookupIDsInDomainRequest, opts ...dcerpc.CallOption) (*LookupIDsInDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupIDsInDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) OpenGroup(ctx context.Context, in *OpenGroupRequest, opts ...dcerpc.CallOption) (*OpenGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) QueryInformationGroup(ctx context.Context, in *QueryInformationGroupRequest, opts ...dcerpc.CallOption) (*QueryInformationGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryInformationGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) SetInformationGroup(ctx context.Context, in *SetInformationGroupRequest, opts ...dcerpc.CallOption) (*SetInformationGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInformationGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) AddMemberToGroup(ctx context.Context, in *AddMemberToGroupRequest, opts ...dcerpc.CallOption) (*AddMemberToGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddMemberToGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...dcerpc.CallOption) (*DeleteGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) RemoveMemberFromGroup(ctx context.Context, in *RemoveMemberFromGroupRequest, opts ...dcerpc.CallOption) (*RemoveMemberFromGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveMemberFromGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) GetMembersInGroup(ctx context.Context, in *GetMembersInGroupRequest, opts ...dcerpc.CallOption) (*GetMembersInGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetMembersInGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) SetMemberAttributesOfGroup(ctx context.Context, in *SetMemberAttributesOfGroupRequest, opts ...dcerpc.CallOption) (*SetMemberAttributesOfGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetMemberAttributesOfGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) OpenAlias(ctx context.Context, in *OpenAliasRequest, opts ...dcerpc.CallOption) (*OpenAliasResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenAliasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) QueryInformationAlias(ctx context.Context, in *QueryInformationAliasRequest, opts ...dcerpc.CallOption) (*QueryInformationAliasResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryInformationAliasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) SetInformationAlias(ctx context.Context, in *SetInformationAliasRequest, opts ...dcerpc.CallOption) (*SetInformationAliasResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInformationAliasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) DeleteAlias(ctx context.Context, in *DeleteAliasRequest, opts ...dcerpc.CallOption) (*DeleteAliasResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteAliasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) AddMemberToAlias(ctx context.Context, in *AddMemberToAliasRequest, opts ...dcerpc.CallOption) (*AddMemberToAliasResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddMemberToAliasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) RemoveMemberFromAlias(ctx context.Context, in *RemoveMemberFromAliasRequest, opts ...dcerpc.CallOption) (*RemoveMemberFromAliasResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveMemberFromAliasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) GetMembersInAlias(ctx context.Context, in *GetMembersInAliasRequest, opts ...dcerpc.CallOption) (*GetMembersInAliasResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetMembersInAliasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) OpenUser(ctx context.Context, in *OpenUserRequest, opts ...dcerpc.CallOption) (*OpenUserResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...dcerpc.CallOption) (*DeleteUserResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) QueryInformationUser(ctx context.Context, in *QueryInformationUserRequest, opts ...dcerpc.CallOption) (*QueryInformationUserResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryInformationUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) SetInformationUser(ctx context.Context, in *SetInformationUserRequest, opts ...dcerpc.CallOption) (*SetInformationUserResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInformationUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) ChangePasswordUser(ctx context.Context, in *ChangePasswordUserRequest, opts ...dcerpc.CallOption) (*ChangePasswordUserResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ChangePasswordUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) GetGroupsForUser(ctx context.Context, in *GetGroupsForUserRequest, opts ...dcerpc.CallOption) (*GetGroupsForUserResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetGroupsForUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) QueryDisplayInformation(ctx context.Context, in *QueryDisplayInformationRequest, opts ...dcerpc.CallOption) (*QueryDisplayInformationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryDisplayInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) GetDisplayEnumerationIndex(ctx context.Context, in *GetDisplayEnumerationIndexRequest, opts ...dcerpc.CallOption) (*GetDisplayEnumerationIndexResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDisplayEnumerationIndexResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) GetUserDomainPasswordInformation(ctx context.Context, in *GetUserDomainPasswordInformationRequest, opts ...dcerpc.CallOption) (*GetUserDomainPasswordInformationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetUserDomainPasswordInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) RemoveMemberFromForeignDomain(ctx context.Context, in *RemoveMemberFromForeignDomainRequest, opts ...dcerpc.CallOption) (*RemoveMemberFromForeignDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveMemberFromForeignDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) QueryInformationDomain2(ctx context.Context, in *QueryInformationDomain2Request, opts ...dcerpc.CallOption) (*QueryInformationDomain2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryInformationDomain2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) QueryInformationUser2(ctx context.Context, in *QueryInformationUser2Request, opts ...dcerpc.CallOption) (*QueryInformationUser2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryInformationUser2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) QueryDisplayInformation2(ctx context.Context, in *QueryDisplayInformation2Request, opts ...dcerpc.CallOption) (*QueryDisplayInformation2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryDisplayInformation2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) GetDisplayEnumerationIndex2(ctx context.Context, in *GetDisplayEnumerationIndex2Request, opts ...dcerpc.CallOption) (*GetDisplayEnumerationIndex2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDisplayEnumerationIndex2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) CreateUser2InDomain(ctx context.Context, in *CreateUser2InDomainRequest, opts ...dcerpc.CallOption) (*CreateUser2InDomainResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateUser2InDomainResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) QueryDisplayInformation3(ctx context.Context, in *QueryDisplayInformation3Request, opts ...dcerpc.CallOption) (*QueryDisplayInformation3Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryDisplayInformation3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) AddMultipleMembersToAlias(ctx context.Context, in *AddMultipleMembersToAliasRequest, opts ...dcerpc.CallOption) (*AddMultipleMembersToAliasResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddMultipleMembersToAliasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) RemoveMultipleMembersFromAlias(ctx context.Context, in *RemoveMultipleMembersFromAliasRequest, opts ...dcerpc.CallOption) (*RemoveMultipleMembersFromAliasResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveMultipleMembersFromAliasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) OEMChangePasswordUser2(ctx context.Context, in *OEMChangePasswordUser2Request, opts ...dcerpc.CallOption) (*OEMChangePasswordUser2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OEMChangePasswordUser2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) UnicodeChangePasswordUser2(ctx context.Context, in *UnicodeChangePasswordUser2Request, opts ...dcerpc.CallOption) (*UnicodeChangePasswordUser2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UnicodeChangePasswordUser2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) GetDomainPasswordInformation(ctx context.Context, in *GetDomainPasswordInformationRequest, opts ...dcerpc.CallOption) (*GetDomainPasswordInformationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDomainPasswordInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) Connect2(ctx context.Context, in *Connect2Request, opts ...dcerpc.CallOption) (*Connect2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Connect2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) SetInformationUser2(ctx context.Context, in *SetInformationUser2Request, opts ...dcerpc.CallOption) (*SetInformationUser2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInformationUser2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) Connect4(ctx context.Context, in *Connect4Request, opts ...dcerpc.CallOption) (*Connect4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Connect4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) Connect5(ctx context.Context, in *Connect5Request, opts ...dcerpc.CallOption) (*Connect5Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Connect5Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) RIDToSID(ctx context.Context, in *RIDToSIDRequest, opts ...dcerpc.CallOption) (*RIDToSIDResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RIDToSIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) SetDSRMPassword(ctx context.Context, in *SetDSRMPasswordRequest, opts ...dcerpc.CallOption) (*SetDSRMPasswordResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetDSRMPasswordResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) ValidatePassword(ctx context.Context, in *ValidatePasswordRequest, opts ...dcerpc.CallOption) (*ValidatePasswordResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ValidatePasswordResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSamrClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultSamrClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewSamrClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (SamrClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(SamrSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultSamrClient{cc: cc}, nil
}

// xxx_ConnectOperation structure represents the SamrConnect operation
type xxx_ConnectOperation struct {
	ServerName    string  `idl:"name:ServerName;pointer:unique" json:"server_name"`
	Server        *Handle `idl:"name:ServerHandle" json:"server"`
	DesiredAccess uint32  `idl:"name:DesiredAccess" json:"desired_access"`
	Return        int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectOperation) OpNum() int { return 0 }

func (o *xxx_ConnectOperation) OpName() string { return "/samr/v1/SamrConnect" }

func (o *xxx_ConnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PSAMPR_SERVER_NAME}*(1)[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PSAMPR_SERVER_NAME}*(1)[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ServerHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ServerHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &Handle{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectRequest structure represents the SamrConnect operation request
type ConnectRequest struct {
	// ServerName: The first character of the NETBIOS name of the server; this parameter
	// MAY<47> be ignored on receipt.
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// DesiredAccess: An ACCESS_MASK that indicates the access requested for ServerHandle
	// upon output. See section 2.2.1.3 for a listing of possible values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *ConnectRequest) xxx_ToOp(ctx context.Context) *xxx_ConnectOperation {
	if o == nil {
		return &xxx_ConnectOperation{}
	}
	return &xxx_ConnectOperation{
		ServerName:    o.ServerName,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *ConnectRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
}
func (o *ConnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ConnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectResponse structure represents the SamrConnect operation response
type ConnectResponse struct {
	// ServerHandle: An RPC context handle, as specified in section 2.2.7.2.
	Server *Handle `idl:"name:ServerHandle" json:"server"`
	// Return: The SamrConnect return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ConnectResponse) xxx_ToOp(ctx context.Context) *xxx_ConnectOperation {
	if o == nil {
		return &xxx_ConnectOperation{}
	}
	return &xxx_ConnectOperation{
		Server: o.Server,
		Return: o.Return,
	}
}

func (o *ConnectResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Return = op.Return
}
func (o *ConnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ConnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseHandleOperation structure represents the SamrCloseHandle operation
type xxx_CloseHandleOperation struct {
	SAMHandle *Handle `idl:"name:SamHandle" json:"sam_handle"`
	Return    int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseHandleOperation) OpNum() int { return 1 }

func (o *xxx_CloseHandleOperation) OpName() string { return "/samr/v1/SamrCloseHandle" }

func (o *xxx_CloseHandleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseHandleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SamHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SAMHandle != nil {
			if err := o.SAMHandle.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CloseHandleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SamHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SAMHandle == nil {
			o.SAMHandle = &Handle{}
		}
		if err := o.SAMHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseHandleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseHandleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SamHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SAMHandle != nil {
			if err := o.SAMHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseHandleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SamHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SAMHandle == nil {
			o.SAMHandle = &Handle{}
		}
		if err := o.SAMHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseHandleRequest structure represents the SamrCloseHandle operation request
type CloseHandleRequest struct {
	// SamHandle: An RPC context handle, as specified in section 2.2.7.2, representing any
	// context handle returned from this interface.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	SAMHandle *Handle `idl:"name:SamHandle" json:"sam_handle"`
}

func (o *CloseHandleRequest) xxx_ToOp(ctx context.Context) *xxx_CloseHandleOperation {
	if o == nil {
		return &xxx_CloseHandleOperation{}
	}
	return &xxx_CloseHandleOperation{
		SAMHandle: o.SAMHandle,
	}
}

func (o *CloseHandleRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseHandleOperation) {
	if o == nil {
		return
	}
	o.SAMHandle = op.SAMHandle
}
func (o *CloseHandleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseHandleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseHandleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseHandleResponse structure represents the SamrCloseHandle operation response
type CloseHandleResponse struct {
	// SamHandle: An RPC context handle, as specified in section 2.2.7.2, representing any
	// context handle returned from this interface.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	SAMHandle *Handle `idl:"name:SamHandle" json:"sam_handle"`
	// Return: The SamrCloseHandle return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseHandleResponse) xxx_ToOp(ctx context.Context) *xxx_CloseHandleOperation {
	if o == nil {
		return &xxx_CloseHandleOperation{}
	}
	return &xxx_CloseHandleOperation{
		SAMHandle: o.SAMHandle,
		Return:    o.Return,
	}
}

func (o *CloseHandleResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseHandleOperation) {
	if o == nil {
		return
	}
	o.SAMHandle = op.SAMHandle
	o.Return = op.Return
}
func (o *CloseHandleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseHandleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseHandleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSecurityObjectOperation structure represents the SamrSetSecurityObject operation
type xxx_SetSecurityObjectOperation struct {
	Object              *Handle               `idl:"name:ObjectHandle" json:"object"`
	SecurityInformation uint32                `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  *SrSecurityDescriptor `idl:"name:SecurityDescriptor" json:"security_descriptor"`
	Return              int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSecurityObjectOperation) OpNum() int { return 2 }

func (o *xxx_SetSecurityObjectOperation) OpName() string { return "/samr/v1/SamrSetSecurityObject" }

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
	// ObjectHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
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
	// SecurityDescriptor {in} (1:{alias=PSAMPR_SR_SECURITY_DESCRIPTOR}*(1))(2:{alias=SAMPR_SR_SECURITY_DESCRIPTOR}(struct))
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
	// ObjectHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
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
	// SecurityDescriptor {in} (1:{alias=PSAMPR_SR_SECURITY_DESCRIPTOR,pointer=ref}*(1))(2:{alias=SAMPR_SR_SECURITY_DESCRIPTOR}(struct))
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
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSecurityObjectRequest structure represents the SamrSetSecurityObject operation request
type SetSecurityObjectRequest struct {
	// ObjectHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a server, domain, user, group, or alias object.
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	// SecurityInformation: A bit field that indicates the fields of SecurityDescriptor
	// that are requested to be set.
	//
	// The SECURITY_INFORMATION type is defined in [MS-DTYP] section 2.4.7. The following
	// bits are valid; all other bits MUST be zero when sent and ignored on receipt. If
	// none of the bits below are present, the server MUST return STATUS_INVALID_PARAMETER.
	//
	//	+---------------------------------------+-----------------------------------------------------------------------------+
	//	|                                       |                                                                             |
	//	|                 VALUE                 |                                   MEANING                                   |
	//	|                                       |                                                                             |
	//	+---------------------------------------+-----------------------------------------------------------------------------+
	//	+---------------------------------------+-----------------------------------------------------------------------------+
	//	| OWNER_SECURITY_INFORMATION 0x00000001 | Refers to the Owner member of the security descriptor.                      |
	//	+---------------------------------------+-----------------------------------------------------------------------------+
	//	| GROUP_SECURITY_INFORMATION 0x00000002 | Refers to the Group member of the security descriptor.                      |
	//	+---------------------------------------+-----------------------------------------------------------------------------+
	//	| DACL_SECURITY_INFORMATION 0x00000004  | Refers to the DACL of the security descriptor.                              |
	//	+---------------------------------------+-----------------------------------------------------------------------------+
	//	| SACL_SECURITY_INFORMATION 0x00000008  | Refers to the system access control list (SACL) of the security descriptor. |
	//	+---------------------------------------+-----------------------------------------------------------------------------+
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// SecurityDescriptor: A security descriptor expressing access that is specific to the
	// ObjectHandle.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
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

// SetSecurityObjectResponse structure represents the SamrSetSecurityObject operation response
type SetSecurityObjectResponse struct {
	// Return: The SamrSetSecurityObject return value.
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

// xxx_QuerySecurityObjectOperation structure represents the SamrQuerySecurityObject operation
type xxx_QuerySecurityObjectOperation struct {
	Object              *Handle               `idl:"name:ObjectHandle" json:"object"`
	SecurityInformation uint32                `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  *SrSecurityDescriptor `idl:"name:SecurityDescriptor" json:"security_descriptor"`
	Return              int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_QuerySecurityObjectOperation) OpNum() int { return 3 }

func (o *xxx_QuerySecurityObjectOperation) OpName() string { return "/samr/v1/SamrQuerySecurityObject" }

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
	// ObjectHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
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
	// ObjectHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
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
	// SecurityDescriptor {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_SR_SECURITY_DESCRIPTOR}*(1))(3:{alias=SAMPR_SR_SECURITY_DESCRIPTOR}(struct))
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
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySecurityObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SecurityDescriptor {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_SR_SECURITY_DESCRIPTOR,pointer=ref}*(1))(3:{alias=SAMPR_SR_SECURITY_DESCRIPTOR}(struct))
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
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QuerySecurityObjectRequest structure represents the SamrQuerySecurityObject operation request
type QuerySecurityObjectRequest struct {
	// ObjectHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a server, domain, user, group, or alias object.
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	// SecurityInformation: A bit field that specifies which fields of SecurityDescriptor
	// the client is requesting to be returned.
	//
	// The SECURITY_INFORMATION type is defined in [MS-DTYP] section 2.4.7. The following
	// bits are valid; all other bits MUST be zero when sent and ignored on receipt.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                     MEANING                                      |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| OWNER_SECURITY_INFORMATION 0x00000001 | If this bit is set, the client requests that the Owner member be returned. If    |
	//	|                                       | this bit is not set, the client requests that the Owner member not be returned.  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| GROUP_SECURITY_INFORMATION 0x00000002 | If this bit is set, the client requests that the Group member be returned. If    |
	//	|                                       | this bit is not set, the client requests that the Group member not be returned.  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| DACL_SECURITY_INFORMATION 0x00000004  | If this bit is set, the client requests that the DACL be returned. If this bit   |
	//	|                                       | is not set, the client requests that the DACL not be returned.                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| SACL_SECURITY_INFORMATION 0x00000008  | If this bit is set, the client requests that the SACL be returned. If this bit   |
	//	|                                       | is not set, the client requests that the SACL not be returned.                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
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

// QuerySecurityObjectResponse structure represents the SamrQuerySecurityObject operation response
type QuerySecurityObjectResponse struct {
	// SecurityDescriptor: A security descriptor expressing accesses that are specific to
	// the ObjectHandle and the owner and group of the object. [MS-DTYP] section 2.4.6 contains
	// the specification for a valid security descriptor.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	SecurityDescriptor *SrSecurityDescriptor `idl:"name:SecurityDescriptor" json:"security_descriptor"`
	// Return: The SamrQuerySecurityObject return value.
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

// xxx_LookupDomainInSAMServerOperation structure represents the SamrLookupDomainInSamServer operation
type xxx_LookupDomainInSAMServerOperation struct {
	Server   *Handle             `idl:"name:ServerHandle" json:"server"`
	Name     *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	DomainID *dtyp.SID           `idl:"name:DomainId" json:"domain_id"`
	Return   int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupDomainInSAMServerOperation) OpNum() int { return 5 }

func (o *xxx_LookupDomainInSAMServerOperation) OpName() string {
	return "/samr/v1/SamrLookupDomainInSamServer"
}

func (o *xxx_LookupDomainInSAMServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupDomainInSAMServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_LookupDomainInSAMServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &Handle{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
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

func (o *xxx_LookupDomainInSAMServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupDomainInSAMServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// DomainId {out} (1:{pointer=ref}*(2))(2:{alias=PRPC_SID}*(1))(3:{alias=RPC_SID}(struct))
	{
		if o.DomainID != nil {
			_ptr_DomainId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DomainID != nil {
					if err := o.DomainID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainID, _ptr_DomainId); err != nil {
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
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupDomainInSAMServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// DomainId {out} (1:{pointer=ref}*(2))(2:{alias=PRPC_SID,pointer=ref}*(1))(3:{alias=RPC_SID}(struct))
	{
		_ptr_DomainId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DomainID == nil {
				o.DomainID = &dtyp.SID{}
			}
			if err := o.DomainID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DomainId := func(ptr interface{}) { o.DomainID = *ptr.(**dtyp.SID) }
		if err := w.ReadPointer(&o.DomainID, _s_DomainId, _ptr_DomainId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupDomainInSAMServerRequest structure represents the SamrLookupDomainInSamServer operation request
type LookupDomainInSAMServerRequest struct {
	// ServerHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a server object.
	Server *Handle `idl:"name:ServerHandle" json:"server"`
	// Name: A UTF-16 encoded string.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
}

func (o *LookupDomainInSAMServerRequest) xxx_ToOp(ctx context.Context) *xxx_LookupDomainInSAMServerOperation {
	if o == nil {
		return &xxx_LookupDomainInSAMServerOperation{}
	}
	return &xxx_LookupDomainInSAMServerOperation{
		Server: o.Server,
		Name:   o.Name,
	}
}

func (o *LookupDomainInSAMServerRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupDomainInSAMServerOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Name = op.Name
}
func (o *LookupDomainInSAMServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *LookupDomainInSAMServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupDomainInSAMServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupDomainInSAMServerResponse structure represents the SamrLookupDomainInSamServer operation response
type LookupDomainInSAMServerResponse struct {
	// DomainId: A SID value of a domain that corresponds to the Name passed in. The match
	// MUST be exact (no wildcard characters are permitted). See message processing later
	// in this section for more details.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	DomainID *dtyp.SID `idl:"name:DomainId" json:"domain_id"`
	// Return: The SamrLookupDomainInSamServer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupDomainInSAMServerResponse) xxx_ToOp(ctx context.Context) *xxx_LookupDomainInSAMServerOperation {
	if o == nil {
		return &xxx_LookupDomainInSAMServerOperation{}
	}
	return &xxx_LookupDomainInSAMServerOperation{
		DomainID: o.DomainID,
		Return:   o.Return,
	}
}

func (o *LookupDomainInSAMServerResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupDomainInSAMServerOperation) {
	if o == nil {
		return
	}
	o.DomainID = op.DomainID
	o.Return = op.Return
}
func (o *LookupDomainInSAMServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *LookupDomainInSAMServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupDomainInSAMServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateDomainsInSAMServerOperation structure represents the SamrEnumerateDomainsInSamServer operation
type xxx_EnumerateDomainsInSAMServerOperation struct {
	Server                 *Handle            `idl:"name:ServerHandle" json:"server"`
	EnumerationContext     uint32             `idl:"name:EnumerationContext" json:"enumeration_context"`
	Buffer                 *EnumerationBuffer `idl:"name:Buffer" json:"buffer"`
	PreferredMaximumLength uint32             `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
	CountReturned          uint32             `idl:"name:CountReturned" json:"count_returned"`
	Return                 int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateDomainsInSAMServerOperation) OpNum() int { return 6 }

func (o *xxx_EnumerateDomainsInSAMServerOperation) OpName() string {
	return "/samr/v1/SamrEnumerateDomainsInSamServer"
}

func (o *xxx_EnumerateDomainsInSAMServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateDomainsInSAMServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnumerateDomainsInSAMServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &Handle{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnumerateDomainsInSAMServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateDomainsInSAMServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_ENUMERATION_BUFFER}*(1))(3:{alias=SAMPR_ENUMERATION_BUFFER}(struct))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil {
					if err := o.Buffer.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EnumerationBuffer{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// CountReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.CountReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateDomainsInSAMServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_ENUMERATION_BUFFER,pointer=ref}*(1))(3:{alias=SAMPR_ENUMERATION_BUFFER}(struct))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &EnumerationBuffer{}
			}
			if err := o.Buffer.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**EnumerationBuffer) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CountReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.CountReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateDomainsInSAMServerRequest structure represents the SamrEnumerateDomainsInSamServer operation request
type EnumerateDomainsInSAMServerRequest struct {
	// ServerHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a server object.
	Server *Handle `idl:"name:ServerHandle" json:"server"`
	// EnumerationContext: This value is a cookie that the server can use to continue an
	// enumeration on a subsequent call. It is an opaque value to the client. To initiate
	// a new enumeration, the client sets EnumerationContext to zero. Otherwise the client
	// sets EnumerationContext to a value returned by a previous call to the method.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// PreferedMaximumLength: The requested maximum number of bytes to return in Buffer.
	PreferredMaximumLength uint32 `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
}

func (o *EnumerateDomainsInSAMServerRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateDomainsInSAMServerOperation {
	if o == nil {
		return &xxx_EnumerateDomainsInSAMServerOperation{}
	}
	return &xxx_EnumerateDomainsInSAMServerOperation{
		Server:                 o.Server,
		EnumerationContext:     o.EnumerationContext,
		PreferredMaximumLength: o.PreferredMaximumLength,
	}
}

func (o *EnumerateDomainsInSAMServerRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateDomainsInSAMServerOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.EnumerationContext = op.EnumerationContext
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *EnumerateDomainsInSAMServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateDomainsInSAMServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateDomainsInSAMServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateDomainsInSAMServerResponse structure represents the SamrEnumerateDomainsInSamServer operation response
type EnumerateDomainsInSAMServerResponse struct {
	// EnumerationContext: This value is a cookie that the server can use to continue an
	// enumeration on a subsequent call. It is an opaque value to the client. To initiate
	// a new enumeration, the client sets EnumerationContext to zero. Otherwise the client
	// sets EnumerationContext to a value returned by a previous call to the method.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// Buffer: A listing of domain information, as described in section 2.2.7.10.
	Buffer *EnumerationBuffer `idl:"name:Buffer" json:"buffer"`
	// CountReturned: The count of domain elements returned in Buffer.
	//
	// This method asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	CountReturned uint32 `idl:"name:CountReturned" json:"count_returned"`
	// Return: The SamrEnumerateDomainsInSamServer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateDomainsInSAMServerResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateDomainsInSAMServerOperation {
	if o == nil {
		return &xxx_EnumerateDomainsInSAMServerOperation{}
	}
	return &xxx_EnumerateDomainsInSAMServerOperation{
		EnumerationContext: o.EnumerationContext,
		Buffer:             o.Buffer,
		CountReturned:      o.CountReturned,
		Return:             o.Return,
	}
}

func (o *EnumerateDomainsInSAMServerResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateDomainsInSAMServerOperation) {
	if o == nil {
		return
	}
	o.EnumerationContext = op.EnumerationContext
	o.Buffer = op.Buffer
	o.CountReturned = op.CountReturned
	o.Return = op.Return
}
func (o *EnumerateDomainsInSAMServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateDomainsInSAMServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateDomainsInSAMServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenDomainOperation structure represents the SamrOpenDomain operation
type xxx_OpenDomainOperation struct {
	Server        *Handle   `idl:"name:ServerHandle" json:"server"`
	DesiredAccess uint32    `idl:"name:DesiredAccess" json:"desired_access"`
	DomainID      *dtyp.SID `idl:"name:DomainId" json:"domain_id"`
	Domain        *Handle   `idl:"name:DomainHandle" json:"domain"`
	Return        int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenDomainOperation) OpNum() int { return 7 }

func (o *xxx_OpenDomainOperation) OpName() string { return "/samr/v1/SamrOpenDomain" }

func (o *xxx_OpenDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	// DomainId {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.DomainID != nil {
			if err := o.DomainID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &Handle{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	// DomainId {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.DomainID == nil {
			o.DomainID = &dtyp.SID{}
		}
		if err := o.DomainID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// DomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenDomainRequest structure represents the SamrOpenDomain operation request
type OpenDomainRequest struct {
	// ServerHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a server object.
	Server *Handle `idl:"name:ServerHandle" json:"server"`
	// DesiredAccess: An ACCESS_MASK. See section 2.2.1.4 for a list of domain access values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
	// DomainId: A SID value of a domain hosted by the server side of this protocol.
	DomainID *dtyp.SID `idl:"name:DomainId" json:"domain_id"`
}

func (o *OpenDomainRequest) xxx_ToOp(ctx context.Context) *xxx_OpenDomainOperation {
	if o == nil {
		return &xxx_OpenDomainOperation{}
	}
	return &xxx_OpenDomainOperation{
		Server:        o.Server,
		DesiredAccess: o.DesiredAccess,
		DomainID:      o.DomainID,
	}
}

func (o *OpenDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenDomainOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.DesiredAccess = op.DesiredAccess
	o.DomainID = op.DomainID
}
func (o *OpenDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenDomainResponse structure represents the SamrOpenDomain operation response
type OpenDomainResponse struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// Return: The SamrOpenDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenDomainResponse) xxx_ToOp(ctx context.Context) *xxx_OpenDomainOperation {
	if o == nil {
		return &xxx_OpenDomainOperation{}
	}
	return &xxx_OpenDomainOperation{
		Domain: o.Domain,
		Return: o.Return,
	}
}

func (o *OpenDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.Return = op.Return
}
func (o *OpenDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInformationDomainOperation structure represents the SamrQueryInformationDomain operation
type xxx_QueryInformationDomainOperation struct {
	Domain                 *Handle                `idl:"name:DomainHandle" json:"domain"`
	DomainInformationClass DomainInformationClass `idl:"name:DomainInformationClass" json:"domain_information_class"`
	Buffer                 *DomainInfoBuffer      `idl:"name:Buffer;switch_is:DomainInformationClass" json:"buffer"`
	Return                 int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInformationDomainOperation) OpNum() int { return 8 }

func (o *xxx_QueryInformationDomainOperation) OpName() string {
	return "/samr/v1/SamrQueryInformationDomain"
}

func (o *xxx_QueryInformationDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DomainInformationClass {in} (1:{alias=DOMAIN_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.DomainInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DomainInformationClass {in} (1:{alias=DOMAIN_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.DomainInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=PSAMPR_DOMAIN_INFO_BUFFER}*(1))(3:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=SAMPR_DOMAIN_INFO_BUFFER}(union))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swBuffer := uint16(o.DomainInformationClass)
				if o.Buffer != nil {
					if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
						return err
					}
				} else {
					if err := (&DomainInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=PSAMPR_DOMAIN_INFO_BUFFER,pointer=ref}*(1))(3:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=SAMPR_DOMAIN_INFO_BUFFER}(union))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &DomainInfoBuffer{}
			}
			_swBuffer := uint16(o.DomainInformationClass)
			if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**DomainInfoBuffer) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryInformationDomainRequest structure represents the SamrQueryInformationDomain operation request
type QueryInformationDomainRequest struct {
	Domain                 *Handle                `idl:"name:DomainHandle" json:"domain"`
	DomainInformationClass DomainInformationClass `idl:"name:DomainInformationClass" json:"domain_information_class"`
}

func (o *QueryInformationDomainRequest) xxx_ToOp(ctx context.Context) *xxx_QueryInformationDomainOperation {
	if o == nil {
		return &xxx_QueryInformationDomainOperation{}
	}
	return &xxx_QueryInformationDomainOperation{
		Domain:                 o.Domain,
		DomainInformationClass: o.DomainInformationClass,
	}
}

func (o *QueryInformationDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DomainInformationClass = op.DomainInformationClass
}
func (o *QueryInformationDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInformationDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInformationDomainResponse structure represents the SamrQueryInformationDomain operation response
type QueryInformationDomainResponse struct {
	Buffer *DomainInfoBuffer `idl:"name:Buffer;switch_is:DomainInformationClass" json:"buffer"`
	// Return: The SamrQueryInformationDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryInformationDomainResponse) xxx_ToOp(ctx context.Context) *xxx_QueryInformationDomainOperation {
	if o == nil {
		return &xxx_QueryInformationDomainOperation{}
	}
	return &xxx_QueryInformationDomainOperation{
		Buffer: o.Buffer,
		Return: o.Return,
	}
}

func (o *QueryInformationDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationDomainOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *QueryInformationDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInformationDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInformationDomainOperation structure represents the SamrSetInformationDomain operation
type xxx_SetInformationDomainOperation struct {
	Domain                 *Handle                `idl:"name:DomainHandle" json:"domain"`
	DomainInformationClass DomainInformationClass `idl:"name:DomainInformationClass" json:"domain_information_class"`
	DomainInformation      *DomainInfoBuffer      `idl:"name:DomainInformation;switch_is:DomainInformationClass" json:"domain_information"`
	Return                 int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInformationDomainOperation) OpNum() int { return 9 }

func (o *xxx_SetInformationDomainOperation) OpName() string {
	return "/samr/v1/SamrSetInformationDomain"
}

func (o *xxx_SetInformationDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DomainInformationClass {in} (1:{alias=DOMAIN_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.DomainInformationClass)); err != nil {
			return err
		}
	}
	// DomainInformation {in} (1:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=PSAMPR_DOMAIN_INFO_BUFFER}*(1))(2:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=SAMPR_DOMAIN_INFO_BUFFER}(union))
	{
		_swDomainInformation := uint16(o.DomainInformationClass)
		if o.DomainInformation != nil {
			if err := o.DomainInformation.MarshalUnionNDR(ctx, w, _swDomainInformation); err != nil {
				return err
			}
		} else {
			if err := (&DomainInfoBuffer{}).MarshalUnionNDR(ctx, w, _swDomainInformation); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DomainInformationClass {in} (1:{alias=DOMAIN_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.DomainInformationClass)); err != nil {
			return err
		}
	}
	// DomainInformation {in} (1:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=PSAMPR_DOMAIN_INFO_BUFFER,pointer=ref}*(1))(2:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=SAMPR_DOMAIN_INFO_BUFFER}(union))
	{
		if o.DomainInformation == nil {
			o.DomainInformation = &DomainInfoBuffer{}
		}
		_swDomainInformation := uint16(o.DomainInformationClass)
		if err := o.DomainInformation.UnmarshalUnionNDR(ctx, w, _swDomainInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInformationDomainRequest structure represents the SamrSetInformationDomain operation request
type SetInformationDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// DomainInformationClass: An enumeration indicating which attributes to update. See
	// section 2.2.3.16 for a list of possible values.
	DomainInformationClass DomainInformationClass `idl:"name:DomainInformationClass" json:"domain_information_class"`
	// DomainInformation: The requested attributes and values to update. See section 2.2.3.17
	// for structure details.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	DomainInformation *DomainInfoBuffer `idl:"name:DomainInformation;switch_is:DomainInformationClass" json:"domain_information"`
}

func (o *SetInformationDomainRequest) xxx_ToOp(ctx context.Context) *xxx_SetInformationDomainOperation {
	if o == nil {
		return &xxx_SetInformationDomainOperation{}
	}
	return &xxx_SetInformationDomainOperation{
		Domain:                 o.Domain,
		DomainInformationClass: o.DomainInformationClass,
		DomainInformation:      o.DomainInformation,
	}
}

func (o *SetInformationDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInformationDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DomainInformationClass = op.DomainInformationClass
	o.DomainInformation = op.DomainInformation
}
func (o *SetInformationDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetInformationDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInformationDomainResponse structure represents the SamrSetInformationDomain operation response
type SetInformationDomainResponse struct {
	// Return: The SamrSetInformationDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInformationDomainResponse) xxx_ToOp(ctx context.Context) *xxx_SetInformationDomainOperation {
	if o == nil {
		return &xxx_SetInformationDomainOperation{}
	}
	return &xxx_SetInformationDomainOperation{
		Return: o.Return,
	}
}

func (o *SetInformationDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInformationDomainOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetInformationDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetInformationDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateGroupInDomainOperation structure represents the SamrCreateGroupInDomain operation
type xxx_CreateGroupInDomainOperation struct {
	Domain        *Handle             `idl:"name:DomainHandle" json:"domain"`
	Name          *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	DesiredAccess uint32              `idl:"name:DesiredAccess" json:"desired_access"`
	GroupHandle   *Handle             `idl:"name:GroupHandle" json:"group_handle"`
	RelativeID    uint32              `idl:"name:RelativeId" json:"relative_id"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateGroupInDomainOperation) OpNum() int { return 10 }

func (o *xxx_CreateGroupInDomainOperation) OpName() string { return "/samr/v1/SamrCreateGroupInDomain" }

func (o *xxx_CreateGroupInDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupInDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
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
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupInDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
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
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupInDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupInDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// GroupHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle != nil {
			if err := o.GroupHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// RelativeId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RelativeID); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupInDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// GroupHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle == nil {
			o.GroupHandle = &Handle{}
		}
		if err := o.GroupHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// RelativeId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RelativeID); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateGroupInDomainRequest structure represents the SamrCreateGroupInDomain operation request
type CreateGroupInDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// Name: The value to use as the name of the group. Details on how this value maps to
	// the data model are provided later in this section.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// DesiredAccess: The access requested on the GroupHandle on output. See section 2.2.1.5
	// for a listing of possible values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *CreateGroupInDomainRequest) xxx_ToOp(ctx context.Context) *xxx_CreateGroupInDomainOperation {
	if o == nil {
		return &xxx_CreateGroupInDomainOperation{}
	}
	return &xxx_CreateGroupInDomainOperation{
		Domain:        o.Domain,
		Name:          o.Name,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *CreateGroupInDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateGroupInDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.Name = op.Name
	o.DesiredAccess = op.DesiredAccess
}
func (o *CreateGroupInDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateGroupInDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateGroupInDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateGroupInDomainResponse structure represents the SamrCreateGroupInDomain operation response
type CreateGroupInDomainResponse struct {
	// GroupHandle: An RPC context handle, as specified in section 2.2.7.2.
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	// RelativeId: The RID of the newly created group.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	//
	// This method MUST be processed per the specifications in section 3.1.5.4.1, using
	// a group type of GROUP_TYPE_SECURITY_ACCOUNT and using access mask values from section
	// 2.2.1.5.
	RelativeID uint32 `idl:"name:RelativeId" json:"relative_id"`
	// Return: The SamrCreateGroupInDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateGroupInDomainResponse) xxx_ToOp(ctx context.Context) *xxx_CreateGroupInDomainOperation {
	if o == nil {
		return &xxx_CreateGroupInDomainOperation{}
	}
	return &xxx_CreateGroupInDomainOperation{
		GroupHandle: o.GroupHandle,
		RelativeID:  o.RelativeID,
		Return:      o.Return,
	}
}

func (o *CreateGroupInDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateGroupInDomainOperation) {
	if o == nil {
		return
	}
	o.GroupHandle = op.GroupHandle
	o.RelativeID = op.RelativeID
	o.Return = op.Return
}
func (o *CreateGroupInDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateGroupInDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateGroupInDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateGroupsInDomainOperation structure represents the SamrEnumerateGroupsInDomain operation
type xxx_EnumerateGroupsInDomainOperation struct {
	Domain                 *Handle            `idl:"name:DomainHandle" json:"domain"`
	EnumerationContext     uint32             `idl:"name:EnumerationContext" json:"enumeration_context"`
	Buffer                 *EnumerationBuffer `idl:"name:Buffer" json:"buffer"`
	PreferredMaximumLength uint32             `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
	CountReturned          uint32             `idl:"name:CountReturned" json:"count_returned"`
	Return                 int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateGroupsInDomainOperation) OpNum() int { return 11 }

func (o *xxx_EnumerateGroupsInDomainOperation) OpName() string {
	return "/samr/v1/SamrEnumerateGroupsInDomain"
}

func (o *xxx_EnumerateGroupsInDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateGroupsInDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnumerateGroupsInDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnumerateGroupsInDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateGroupsInDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_ENUMERATION_BUFFER}*(1))(3:{alias=SAMPR_ENUMERATION_BUFFER}(struct))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil {
					if err := o.Buffer.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EnumerationBuffer{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// CountReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.CountReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateGroupsInDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_ENUMERATION_BUFFER,pointer=ref}*(1))(3:{alias=SAMPR_ENUMERATION_BUFFER}(struct))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &EnumerationBuffer{}
			}
			if err := o.Buffer.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**EnumerationBuffer) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CountReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.CountReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateGroupsInDomainRequest structure represents the SamrEnumerateGroupsInDomain operation request
type EnumerateGroupsInDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// EnumerationContext: This value is a cookie that the server can use to continue an
	// enumeration on a subsequent call. It is an opaque value to the client. To initiate
	// a new enumeration, the client sets EnumerationContext to zero. Otherwise, the client
	// sets EnumerationContext to a value returned by a previous call to the method.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// PreferedMaximumLength: The requested maximum number of bytes to return in Buffer.
	PreferredMaximumLength uint32 `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
}

func (o *EnumerateGroupsInDomainRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateGroupsInDomainOperation {
	if o == nil {
		return &xxx_EnumerateGroupsInDomainOperation{}
	}
	return &xxx_EnumerateGroupsInDomainOperation{
		Domain:                 o.Domain,
		EnumerationContext:     o.EnumerationContext,
		PreferredMaximumLength: o.PreferredMaximumLength,
	}
}

func (o *EnumerateGroupsInDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateGroupsInDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.EnumerationContext = op.EnumerationContext
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *EnumerateGroupsInDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateGroupsInDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateGroupsInDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateGroupsInDomainResponse structure represents the SamrEnumerateGroupsInDomain operation response
type EnumerateGroupsInDomainResponse struct {
	// EnumerationContext: This value is a cookie that the server can use to continue an
	// enumeration on a subsequent call. It is an opaque value to the client. To initiate
	// a new enumeration, the client sets EnumerationContext to zero. Otherwise, the client
	// sets EnumerationContext to a value returned by a previous call to the method.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// Buffer: A list of group information, as specified in section 2.2.7.10.
	Buffer *EnumerationBuffer `idl:"name:Buffer" json:"buffer"`
	// CountReturned: The count of domain elements returned in Buffer.
	//
	// This method asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	//
	// This method MUST be processed per the specifications in section 3.1.5.2.2 using the
	// following object selection filter:
	//
	// *
	//
	// The *objectClass* attribute value MUST be group or derived from group.
	//
	// *
	//
	// The *groupType* attribute value MUST be one of GROUP_TYPE_SECURITY_UNIVERSAL or GROUP_TYPE_SECURITY_ACCOUNT.
	//
	// *
	//
	// The *objectSid* attribute value MUST have the domain prefix ( 7b2aeb27-92fc-41f6-8437-deb65d950921#gt_9066e9dc-8275-4452-9073-daab5fd427c5
	// ) of the domain referenced by DomainHandle.
	CountReturned uint32 `idl:"name:CountReturned" json:"count_returned"`
	// Return: The SamrEnumerateGroupsInDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateGroupsInDomainResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateGroupsInDomainOperation {
	if o == nil {
		return &xxx_EnumerateGroupsInDomainOperation{}
	}
	return &xxx_EnumerateGroupsInDomainOperation{
		EnumerationContext: o.EnumerationContext,
		Buffer:             o.Buffer,
		CountReturned:      o.CountReturned,
		Return:             o.Return,
	}
}

func (o *EnumerateGroupsInDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateGroupsInDomainOperation) {
	if o == nil {
		return
	}
	o.EnumerationContext = op.EnumerationContext
	o.Buffer = op.Buffer
	o.CountReturned = op.CountReturned
	o.Return = op.Return
}
func (o *EnumerateGroupsInDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateGroupsInDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateGroupsInDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateUserInDomainOperation structure represents the SamrCreateUserInDomain operation
type xxx_CreateUserInDomainOperation struct {
	Domain        *Handle             `idl:"name:DomainHandle" json:"domain"`
	Name          *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	DesiredAccess uint32              `idl:"name:DesiredAccess" json:"desired_access"`
	UserHandle    *Handle             `idl:"name:UserHandle" json:"user_handle"`
	RelativeID    uint32              `idl:"name:RelativeId" json:"relative_id"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateUserInDomainOperation) OpNum() int { return 12 }

func (o *xxx_CreateUserInDomainOperation) OpName() string { return "/samr/v1/SamrCreateUserInDomain" }

func (o *xxx_CreateUserInDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUserInDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
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
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUserInDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
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
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUserInDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUserInDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// UserHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// RelativeId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RelativeID); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUserInDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// UserHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// RelativeId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RelativeID); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateUserInDomainRequest structure represents the SamrCreateUserInDomain operation request
type CreateUserInDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// Name: The value to use as the name of the user. See the message processing shown
	// later in this section for details on how this value maps to the data model.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// DesiredAccess: The access requested on the UserHandle on output. See section 2.2.1.7
	// for a listing of possible values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *CreateUserInDomainRequest) xxx_ToOp(ctx context.Context) *xxx_CreateUserInDomainOperation {
	if o == nil {
		return &xxx_CreateUserInDomainOperation{}
	}
	return &xxx_CreateUserInDomainOperation{
		Domain:        o.Domain,
		Name:          o.Name,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *CreateUserInDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateUserInDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.Name = op.Name
	o.DesiredAccess = op.DesiredAccess
}
func (o *CreateUserInDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateUserInDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateUserInDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateUserInDomainResponse structure represents the SamrCreateUserInDomain operation response
type CreateUserInDomainResponse struct {
	// UserHandle: An RPC context handle, as specified in section 2.2.7.2.
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
	// RelativeId: The RID of the newly created user.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	RelativeID uint32 `idl:"name:RelativeId" json:"relative_id"`
	// Return: The SamrCreateUserInDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateUserInDomainResponse) xxx_ToOp(ctx context.Context) *xxx_CreateUserInDomainOperation {
	if o == nil {
		return &xxx_CreateUserInDomainOperation{}
	}
	return &xxx_CreateUserInDomainOperation{
		UserHandle: o.UserHandle,
		RelativeID: o.RelativeID,
		Return:     o.Return,
	}
}

func (o *CreateUserInDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateUserInDomainOperation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
	o.RelativeID = op.RelativeID
	o.Return = op.Return
}
func (o *CreateUserInDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateUserInDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateUserInDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateUsersInDomainOperation structure represents the SamrEnumerateUsersInDomain operation
type xxx_EnumerateUsersInDomainOperation struct {
	Domain                 *Handle            `idl:"name:DomainHandle" json:"domain"`
	EnumerationContext     uint32             `idl:"name:EnumerationContext" json:"enumeration_context"`
	UserAccountControl     uint32             `idl:"name:UserAccountControl" json:"user_account_control"`
	Buffer                 *EnumerationBuffer `idl:"name:Buffer" json:"buffer"`
	PreferredMaximumLength uint32             `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
	CountReturned          uint32             `idl:"name:CountReturned" json:"count_returned"`
	Return                 int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateUsersInDomainOperation) OpNum() int { return 13 }

func (o *xxx_EnumerateUsersInDomainOperation) OpName() string {
	return "/samr/v1/SamrEnumerateUsersInDomain"
}

func (o *xxx_EnumerateUsersInDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateUsersInDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
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
	// UserAccountControl {in} (1:(uint32))
	{
		if err := w.WriteData(o.UserAccountControl); err != nil {
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

func (o *xxx_EnumerateUsersInDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// UserAccountControl {in} (1:(uint32))
	{
		if err := w.ReadData(&o.UserAccountControl); err != nil {
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

func (o *xxx_EnumerateUsersInDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateUsersInDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_ENUMERATION_BUFFER}*(1))(3:{alias=SAMPR_ENUMERATION_BUFFER}(struct))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil {
					if err := o.Buffer.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EnumerationBuffer{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// CountReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.CountReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateUsersInDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_ENUMERATION_BUFFER,pointer=ref}*(1))(3:{alias=SAMPR_ENUMERATION_BUFFER}(struct))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &EnumerationBuffer{}
			}
			if err := o.Buffer.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**EnumerationBuffer) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CountReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.CountReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateUsersInDomainRequest structure represents the SamrEnumerateUsersInDomain operation request
type EnumerateUsersInDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// EnumerationContext: This value is a cookie that the server can use to continue an
	// enumeration on a subsequent call. It is an opaque value to the client. To initiate
	// a new enumeration the client sets EnumerationContext to zero. Otherwise the client
	// sets EnumerationContext to a value returned by a previous call to the method.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// UserAccountControl: A filter value to be used on the userAccountControl attribute.
	UserAccountControl uint32 `idl:"name:UserAccountControl" json:"user_account_control"`
	// PreferedMaximumLength: The requested maximum number of bytes to return in Buffer.
	PreferredMaximumLength uint32 `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
}

func (o *EnumerateUsersInDomainRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateUsersInDomainOperation {
	if o == nil {
		return &xxx_EnumerateUsersInDomainOperation{}
	}
	return &xxx_EnumerateUsersInDomainOperation{
		Domain:                 o.Domain,
		EnumerationContext:     o.EnumerationContext,
		UserAccountControl:     o.UserAccountControl,
		PreferredMaximumLength: o.PreferredMaximumLength,
	}
}

func (o *EnumerateUsersInDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateUsersInDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.EnumerationContext = op.EnumerationContext
	o.UserAccountControl = op.UserAccountControl
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *EnumerateUsersInDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateUsersInDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateUsersInDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateUsersInDomainResponse structure represents the SamrEnumerateUsersInDomain operation response
type EnumerateUsersInDomainResponse struct {
	// EnumerationContext: This value is a cookie that the server can use to continue an
	// enumeration on a subsequent call. It is an opaque value to the client. To initiate
	// a new enumeration the client sets EnumerationContext to zero. Otherwise the client
	// sets EnumerationContext to a value returned by a previous call to the method.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// Buffer: A list of user information, as specified in section 2.2.7.10.
	Buffer *EnumerationBuffer `idl:"name:Buffer" json:"buffer"`
	// CountReturned: The count of domain elements returned in Buffer.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	//
	// This method MUST be processed per the specifications in section 3.1.5.2.2, using
	// the following object selection filter:
	//
	// *
	//
	// The *objectClass* attribute value MUST be user or derived from user.
	//
	// *
	//
	// The *userAccountControl* attribute value MUST contain all the bits in the method
	// parameter UserAccountControl.
	//
	// *
	//
	// The *objectSid* attribute value MUST have the domain prefix ( 7b2aeb27-92fc-41f6-8437-deb65d950921#gt_9066e9dc-8275-4452-9073-daab5fd427c5
	// ) of the domain referenced by DomainHandle.
	//
	// In addition, all of the following constraints MUST be satisfied before the constraints
	// of section 3.1.5.2.2 are satisfied:
	//
	// *
	//
	// If DomainHandle.Object is a reference to the account domain and the configuration
	// is DC ( 7b2aeb27-92fc-41f6-8437-deb65d950921#gt_76a05049-3531-4abd-aec8-30e19954b4bd
	// ) , the client MUST have the SAM-Enumerate-Entire-Domain control access right ( 7b2aeb27-92fc-41f6-8437-deb65d950921#gt_42f6c9e0-a2b3-4bc3-9b87-fdb902e5505e
	// ) ( [MS-ADTS] ( ../ms-adts/d2435927-0999-4c62-8c6d-13ba31a52e1a ) section 5.1.3.2.1
	// ( ../ms-adts/1522b774-6464-41a3-87a5-1e5633c3fbbb ) ) on the domain's *ntSecurityDescriptor*
	// attribute value.
	//
	// *
	//
	// The server MUST ignore the UF_LOCKOUT and UF_PASSWORD_EXPIRED bits in the UserAccountControl
	// parameter.
	CountReturned uint32 `idl:"name:CountReturned" json:"count_returned"`
	// Return: The SamrEnumerateUsersInDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateUsersInDomainResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateUsersInDomainOperation {
	if o == nil {
		return &xxx_EnumerateUsersInDomainOperation{}
	}
	return &xxx_EnumerateUsersInDomainOperation{
		EnumerationContext: o.EnumerationContext,
		Buffer:             o.Buffer,
		CountReturned:      o.CountReturned,
		Return:             o.Return,
	}
}

func (o *EnumerateUsersInDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateUsersInDomainOperation) {
	if o == nil {
		return
	}
	o.EnumerationContext = op.EnumerationContext
	o.Buffer = op.Buffer
	o.CountReturned = op.CountReturned
	o.Return = op.Return
}
func (o *EnumerateUsersInDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateUsersInDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateUsersInDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateAliasInDomainOperation structure represents the SamrCreateAliasInDomain operation
type xxx_CreateAliasInDomainOperation struct {
	Domain        *Handle             `idl:"name:DomainHandle" json:"domain"`
	AccountName   *dtyp.UnicodeString `idl:"name:AccountName" json:"account_name"`
	DesiredAccess uint32              `idl:"name:DesiredAccess" json:"desired_access"`
	AliasHandle   *Handle             `idl:"name:AliasHandle" json:"alias_handle"`
	RelativeID    uint32              `idl:"name:RelativeId" json:"relative_id"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateAliasInDomainOperation) OpNum() int { return 14 }

func (o *xxx_CreateAliasInDomainOperation) OpName() string { return "/samr/v1/SamrCreateAliasInDomain" }

func (o *xxx_CreateAliasInDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAliasInDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// AccountName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.AccountName != nil {
			if err := o.AccountName.MarshalNDR(ctx, w); err != nil {
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
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAliasInDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// AccountName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.AccountName == nil {
			o.AccountName = &dtyp.UnicodeString{}
		}
		if err := o.AccountName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAliasInDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAliasInDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// AliasHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// RelativeId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RelativeID); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateAliasInDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// RelativeId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RelativeID); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateAliasInDomainRequest structure represents the SamrCreateAliasInDomain operation request
type CreateAliasInDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// AccountName: The value to use as the name of the alias. Details on how this value
	// maps to the data model are provided later in this section.
	AccountName *dtyp.UnicodeString `idl:"name:AccountName" json:"account_name"`
	// DesiredAccess: The access requested on the AliasHandle on output. See section 2.2.1.6
	// for a listing of possible values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *CreateAliasInDomainRequest) xxx_ToOp(ctx context.Context) *xxx_CreateAliasInDomainOperation {
	if o == nil {
		return &xxx_CreateAliasInDomainOperation{}
	}
	return &xxx_CreateAliasInDomainOperation{
		Domain:        o.Domain,
		AccountName:   o.AccountName,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *CreateAliasInDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateAliasInDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.AccountName = op.AccountName
	o.DesiredAccess = op.DesiredAccess
}
func (o *CreateAliasInDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateAliasInDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateAliasInDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateAliasInDomainResponse structure represents the SamrCreateAliasInDomain operation response
type CreateAliasInDomainResponse struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	// RelativeId: The RID of the newly created alias.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	//
	// This method MUST be processed per the specifications in section 3.1.5.4.1, using
	// a group type of GROUP_TYPE_SECURITY_RESOURCE and using access mask values from section
	// 2.2.1.6.
	RelativeID uint32 `idl:"name:RelativeId" json:"relative_id"`
	// Return: The SamrCreateAliasInDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateAliasInDomainResponse) xxx_ToOp(ctx context.Context) *xxx_CreateAliasInDomainOperation {
	if o == nil {
		return &xxx_CreateAliasInDomainOperation{}
	}
	return &xxx_CreateAliasInDomainOperation{
		AliasHandle: o.AliasHandle,
		RelativeID:  o.RelativeID,
		Return:      o.Return,
	}
}

func (o *CreateAliasInDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateAliasInDomainOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
	o.RelativeID = op.RelativeID
	o.Return = op.Return
}
func (o *CreateAliasInDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateAliasInDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateAliasInDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateAliasesInDomainOperation structure represents the SamrEnumerateAliasesInDomain operation
type xxx_EnumerateAliasesInDomainOperation struct {
	Domain                 *Handle            `idl:"name:DomainHandle" json:"domain"`
	EnumerationContext     uint32             `idl:"name:EnumerationContext" json:"enumeration_context"`
	Buffer                 *EnumerationBuffer `idl:"name:Buffer" json:"buffer"`
	PreferredMaximumLength uint32             `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
	CountReturned          uint32             `idl:"name:CountReturned" json:"count_returned"`
	Return                 int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateAliasesInDomainOperation) OpNum() int { return 15 }

func (o *xxx_EnumerateAliasesInDomainOperation) OpName() string {
	return "/samr/v1/SamrEnumerateAliasesInDomain"
}

func (o *xxx_EnumerateAliasesInDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAliasesInDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnumerateAliasesInDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnumerateAliasesInDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAliasesInDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.EnumerationContext); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_ENUMERATION_BUFFER}*(1))(3:{alias=SAMPR_ENUMERATION_BUFFER}(struct))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil {
					if err := o.Buffer.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EnumerationBuffer{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// CountReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.CountReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateAliasesInDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EnumerationContext {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.EnumerationContext); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_ENUMERATION_BUFFER,pointer=ref}*(1))(3:{alias=SAMPR_ENUMERATION_BUFFER}(struct))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &EnumerationBuffer{}
			}
			if err := o.Buffer.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**EnumerationBuffer) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// CountReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.CountReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateAliasesInDomainRequest structure represents the SamrEnumerateAliasesInDomain operation request
type EnumerateAliasesInDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// EnumerationContext: This value is a cookie that the server can use to continue an
	// enumeration on a subsequent call. It is an opaque value to the client. To initiate
	// a new enumeration the client sets EnumerationContext to zero. Otherwise the client
	// sets EnumerationContext to a value returned by a previous call to the method.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// PreferedMaximumLength: The requested maximum number of bytes to return in Buffer.
	PreferredMaximumLength uint32 `idl:"name:PreferedMaximumLength" json:"preferred_maximum_length"`
}

func (o *EnumerateAliasesInDomainRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateAliasesInDomainOperation {
	if o == nil {
		return &xxx_EnumerateAliasesInDomainOperation{}
	}
	return &xxx_EnumerateAliasesInDomainOperation{
		Domain:                 o.Domain,
		EnumerationContext:     o.EnumerationContext,
		PreferredMaximumLength: o.PreferredMaximumLength,
	}
}

func (o *EnumerateAliasesInDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateAliasesInDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.EnumerationContext = op.EnumerationContext
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *EnumerateAliasesInDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateAliasesInDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateAliasesInDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateAliasesInDomainResponse structure represents the SamrEnumerateAliasesInDomain operation response
type EnumerateAliasesInDomainResponse struct {
	// EnumerationContext: This value is a cookie that the server can use to continue an
	// enumeration on a subsequent call. It is an opaque value to the client. To initiate
	// a new enumeration the client sets EnumerationContext to zero. Otherwise the client
	// sets EnumerationContext to a value returned by a previous call to the method.
	EnumerationContext uint32 `idl:"name:EnumerationContext" json:"enumeration_context"`
	// Buffer: A list of alias information, as specified in section 2.2.7.10.
	Buffer *EnumerationBuffer `idl:"name:Buffer" json:"buffer"`
	// CountReturned: The count of domain elements returned in Buffer.
	//
	// This method asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	//
	// This method MUST be processed per the specifications in section 3.1.5.2.2 using the
	// following object selection filter:
	//
	// *
	//
	// The *objectClass* attribute value MUST be group or derived from group.
	//
	// *
	//
	// The *groupType* attribute value MUST be GROUP_TYPE_SECURITY_RESOURCE.
	//
	// *
	//
	// The *objectSid* attribute value MUST have the domain prefix ( 7b2aeb27-92fc-41f6-8437-deb65d950921#gt_9066e9dc-8275-4452-9073-daab5fd427c5
	// ) of the domain referenced by DomainHandle.
	CountReturned uint32 `idl:"name:CountReturned" json:"count_returned"`
	// Return: The SamrEnumerateAliasesInDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateAliasesInDomainResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateAliasesInDomainOperation {
	if o == nil {
		return &xxx_EnumerateAliasesInDomainOperation{}
	}
	return &xxx_EnumerateAliasesInDomainOperation{
		EnumerationContext: o.EnumerationContext,
		Buffer:             o.Buffer,
		CountReturned:      o.CountReturned,
		Return:             o.Return,
	}
}

func (o *EnumerateAliasesInDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateAliasesInDomainOperation) {
	if o == nil {
		return
	}
	o.EnumerationContext = op.EnumerationContext
	o.Buffer = op.Buffer
	o.CountReturned = op.CountReturned
	o.Return = op.Return
}
func (o *EnumerateAliasesInDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateAliasesInDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateAliasesInDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAliasMembershipOperation structure represents the SamrGetAliasMembership operation
type xxx_GetAliasMembershipOperation struct {
	Domain     *Handle      `idl:"name:DomainHandle" json:"domain"`
	SIDArray   *SIDArray    `idl:"name:SidArray" json:"sid_array"`
	Membership *Uint32Array `idl:"name:Membership" json:"membership"`
	Return     int32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAliasMembershipOperation) OpNum() int { return 16 }

func (o *xxx_GetAliasMembershipOperation) OpName() string { return "/samr/v1/SamrGetAliasMembership" }

func (o *xxx_GetAliasMembershipOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAliasMembershipOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SidArray {in} (1:{alias=PSAMPR_PSID_ARRAY}*(1))(2:{alias=SAMPR_PSID_ARRAY}(struct))
	{
		if o.SIDArray != nil {
			if err := o.SIDArray.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SIDArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAliasMembershipOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SidArray {in} (1:{alias=PSAMPR_PSID_ARRAY,pointer=ref}*(1))(2:{alias=SAMPR_PSID_ARRAY}(struct))
	{
		if o.SIDArray == nil {
			o.SIDArray = &SIDArray{}
		}
		if err := o.SIDArray.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAliasMembershipOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAliasMembershipOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Membership {out} (1:{alias=PSAMPR_ULONG_ARRAY}*(1))(2:{alias=SAMPR_ULONG_ARRAY}(struct))
	{
		if o.Membership != nil {
			if err := o.Membership.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Uint32Array{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAliasMembershipOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Membership {out} (1:{alias=PSAMPR_ULONG_ARRAY,pointer=ref}*(1))(2:{alias=SAMPR_ULONG_ARRAY}(struct))
	{
		if o.Membership == nil {
			o.Membership = &Uint32Array{}
		}
		if err := o.Membership.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetAliasMembershipRequest structure represents the SamrGetAliasMembership operation request
type GetAliasMembershipRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// SidArray: A list of SIDs.
	SIDArray *SIDArray `idl:"name:SidArray" json:"sid_array"`
}

func (o *GetAliasMembershipRequest) xxx_ToOp(ctx context.Context) *xxx_GetAliasMembershipOperation {
	if o == nil {
		return &xxx_GetAliasMembershipOperation{}
	}
	return &xxx_GetAliasMembershipOperation{
		Domain:   o.Domain,
		SIDArray: o.SIDArray,
	}
}

func (o *GetAliasMembershipRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAliasMembershipOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.SIDArray = op.SIDArray
}
func (o *GetAliasMembershipRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAliasMembershipRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAliasMembershipOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAliasMembershipResponse structure represents the SamrGetAliasMembership operation response
type GetAliasMembershipResponse struct {
	// Membership: The union of all aliases (represented by RIDs) that all SIDs in SidArray
	// are a member of.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Membership *Uint32Array `idl:"name:Membership" json:"membership"`
	// Return: The SamrGetAliasMembership return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAliasMembershipResponse) xxx_ToOp(ctx context.Context) *xxx_GetAliasMembershipOperation {
	if o == nil {
		return &xxx_GetAliasMembershipOperation{}
	}
	return &xxx_GetAliasMembershipOperation{
		Membership: o.Membership,
		Return:     o.Return,
	}
}

func (o *GetAliasMembershipResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAliasMembershipOperation) {
	if o == nil {
		return
	}
	o.Membership = op.Membership
	o.Return = op.Return
}
func (o *GetAliasMembershipResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAliasMembershipResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAliasMembershipOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupNamesInDomainOperation structure represents the SamrLookupNamesInDomain operation
type xxx_LookupNamesInDomainOperation struct {
	Domain      *Handle               `idl:"name:DomainHandle" json:"domain"`
	Count       uint32                `idl:"name:Count" json:"count"`
	Names       []*dtyp.UnicodeString `idl:"name:Names;size_is:(1000);length_is:(Count)" json:"names"`
	RelativeIDs *Uint32Array          `idl:"name:RelativeIds" json:"relative_ids"`
	Use         *Uint32Array          `idl:"name:Use" json:"use"`
	Return      int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupNamesInDomainOperation) OpNum() int { return 17 }

func (o *xxx_LookupNamesInDomainOperation) OpName() string { return "/samr/v1/SamrLookupNamesInDomain" }

func (o *xxx_LookupNamesInDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	// cannot evaluate expression 1000
	if o.Names != nil && o.Count == 0 {
		o.Count = uint32(len(o.Names))
	}
	if o.Count > uint32(1000) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNamesInDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// Names {in} (1:[dim:0,size_is=1000,length_is=Count])(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		dimSize1 := uint64(1000)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.Count)
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNamesInDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// Names {in} (1:[dim:0,size_is=1000,length_is=Count])(2:{alias=RPC_UNICODE_STRING}(struct))
	{
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
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNamesInDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNamesInDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// RelativeIds {out} (1:{alias=PSAMPR_ULONG_ARRAY}*(1))(2:{alias=SAMPR_ULONG_ARRAY}(struct))
	{
		if o.RelativeIDs != nil {
			if err := o.RelativeIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Uint32Array{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Use {out} (1:{alias=PSAMPR_ULONG_ARRAY}*(1))(2:{alias=SAMPR_ULONG_ARRAY}(struct))
	{
		if o.Use != nil {
			if err := o.Use.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Uint32Array{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNamesInDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// RelativeIds {out} (1:{alias=PSAMPR_ULONG_ARRAY,pointer=ref}*(1))(2:{alias=SAMPR_ULONG_ARRAY}(struct))
	{
		if o.RelativeIDs == nil {
			o.RelativeIDs = &Uint32Array{}
		}
		if err := o.RelativeIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Use {out} (1:{alias=PSAMPR_ULONG_ARRAY,pointer=ref}*(1))(2:{alias=SAMPR_ULONG_ARRAY}(struct))
	{
		if o.Use == nil {
			o.Use = &Uint32Array{}
		}
		if err := o.Use.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupNamesInDomainRequest structure represents the SamrLookupNamesInDomain operation request
type LookupNamesInDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// Count: The number of elements in Names. The maximum value of 1,000 is chosen to limit
	// the amount of memory that the client can force the server to allocate.
	Count uint32 `idl:"name:Count" json:"count"`
	// Names: An array of strings that are to be mapped to RIDs.
	Names []*dtyp.UnicodeString `idl:"name:Names;size_is:(1000);length_is:(Count)" json:"names"`
}

func (o *LookupNamesInDomainRequest) xxx_ToOp(ctx context.Context) *xxx_LookupNamesInDomainOperation {
	if o == nil {
		return &xxx_LookupNamesInDomainOperation{}
	}
	return &xxx_LookupNamesInDomainOperation{
		Domain: o.Domain,
		Count:  o.Count,
		Names:  o.Names,
	}
}

func (o *LookupNamesInDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupNamesInDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.Count = op.Count
	o.Names = op.Names
}
func (o *LookupNamesInDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *LookupNamesInDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNamesInDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupNamesInDomainResponse structure represents the SamrLookupNamesInDomain operation response
type LookupNamesInDomainResponse struct {
	// RelativeIds: An array of RIDs of accounts that correspond to the elements in Names.
	RelativeIDs *Uint32Array `idl:"name:RelativeIds" json:"relative_ids"`
	// Use: An array of SID_NAME_USE enumeration values that describe the type of account
	// for each entry in RelativeIds.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Use *Uint32Array `idl:"name:Use" json:"use"`
	// Return: The SamrLookupNamesInDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupNamesInDomainResponse) xxx_ToOp(ctx context.Context) *xxx_LookupNamesInDomainOperation {
	if o == nil {
		return &xxx_LookupNamesInDomainOperation{}
	}
	return &xxx_LookupNamesInDomainOperation{
		RelativeIDs: o.RelativeIDs,
		Use:         o.Use,
		Return:      o.Return,
	}
}

func (o *LookupNamesInDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupNamesInDomainOperation) {
	if o == nil {
		return
	}
	o.RelativeIDs = op.RelativeIDs
	o.Use = op.Use
	o.Return = op.Return
}
func (o *LookupNamesInDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *LookupNamesInDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNamesInDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupIDsInDomainOperation structure represents the SamrLookupIdsInDomain operation
type xxx_LookupIDsInDomainOperation struct {
	Domain      *Handle               `idl:"name:DomainHandle" json:"domain"`
	Count       uint32                `idl:"name:Count" json:"count"`
	RelativeIDs []uint32              `idl:"name:RelativeIds;size_is:(1000);length_is:(Count)" json:"relative_ids"`
	Names       *ReturnedUstringArray `idl:"name:Names" json:"names"`
	Use         *Uint32Array          `idl:"name:Use" json:"use"`
	Return      int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupIDsInDomainOperation) OpNum() int { return 18 }

func (o *xxx_LookupIDsInDomainOperation) OpName() string { return "/samr/v1/SamrLookupIdsInDomain" }

func (o *xxx_LookupIDsInDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	// cannot evaluate expression 1000
	if o.RelativeIDs != nil && o.Count == 0 {
		o.Count = uint32(len(o.RelativeIDs))
	}
	if o.Count > uint32(1000) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupIDsInDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// RelativeIds {in} (1:{pointer=ref}*(1)[dim:0,size_is=1000,length_is=Count](uint32))
	{
		dimSize1 := uint64(1000)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.Count)
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
		for i1 := range o.RelativeIDs {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.RelativeIDs[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.RelativeIDs); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_LookupIDsInDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// RelativeIds {in} (1:{pointer=ref}*(1)[dim:0,size_is=1000,length_is=Count](uint32))
	{
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
			return fmt.Errorf("buffer overflow for size %d of array o.RelativeIDs", sizeInfo[0])
		}
		o.RelativeIDs = make([]uint32, sizeInfo[0])
		for i1 := range o.RelativeIDs {
			i1 := i1
			if err := w.ReadData(&o.RelativeIDs[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_LookupIDsInDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupIDsInDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Names {out} (1:{alias=PSAMPR_RETURNED_USTRING_ARRAY}*(1))(2:{alias=SAMPR_RETURNED_USTRING_ARRAY}(struct))
	{
		if o.Names != nil {
			if err := o.Names.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ReturnedUstringArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Use {out} (1:{alias=PSAMPR_ULONG_ARRAY}*(1))(2:{alias=SAMPR_ULONG_ARRAY}(struct))
	{
		if o.Use != nil {
			if err := o.Use.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Uint32Array{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupIDsInDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Names {out} (1:{alias=PSAMPR_RETURNED_USTRING_ARRAY,pointer=ref}*(1))(2:{alias=SAMPR_RETURNED_USTRING_ARRAY}(struct))
	{
		if o.Names == nil {
			o.Names = &ReturnedUstringArray{}
		}
		if err := o.Names.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Use {out} (1:{alias=PSAMPR_ULONG_ARRAY,pointer=ref}*(1))(2:{alias=SAMPR_ULONG_ARRAY}(struct))
	{
		if o.Use == nil {
			o.Use = &Uint32Array{}
		}
		if err := o.Use.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupIDsInDomainRequest structure represents the SamrLookupIdsInDomain operation request
type LookupIDsInDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// Count: The number of elements in RelativeIds. The maximum value of 1,000 is chosen
	// to limit the amount of memory that the client can force the server to allocate.
	Count uint32 `idl:"name:Count" json:"count"`
	// RelativeIds: An array of RIDs that are to be mapped to account names.
	RelativeIDs []uint32 `idl:"name:RelativeIds;size_is:(1000);length_is:(Count)" json:"relative_ids"`
}

func (o *LookupIDsInDomainRequest) xxx_ToOp(ctx context.Context) *xxx_LookupIDsInDomainOperation {
	if o == nil {
		return &xxx_LookupIDsInDomainOperation{}
	}
	return &xxx_LookupIDsInDomainOperation{
		Domain:      o.Domain,
		Count:       o.Count,
		RelativeIDs: o.RelativeIDs,
	}
}

func (o *LookupIDsInDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupIDsInDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.Count = op.Count
	o.RelativeIDs = op.RelativeIDs
}
func (o *LookupIDsInDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *LookupIDsInDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupIDsInDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupIDsInDomainResponse structure represents the SamrLookupIdsInDomain operation response
type LookupIDsInDomainResponse struct {
	// Names: A structure containing an array of account names that correspond to the elements
	// in RelativeIds.
	Names *ReturnedUstringArray `idl:"name:Names" json:"names"`
	// Use: A structure containing an array of SID_NAME_USE enumeration values that describe
	// the type of account for each entry in RelativeIds.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Use *Uint32Array `idl:"name:Use" json:"use"`
	// Return: The SamrLookupIdsInDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupIDsInDomainResponse) xxx_ToOp(ctx context.Context) *xxx_LookupIDsInDomainOperation {
	if o == nil {
		return &xxx_LookupIDsInDomainOperation{}
	}
	return &xxx_LookupIDsInDomainOperation{
		Names:  o.Names,
		Use:    o.Use,
		Return: o.Return,
	}
}

func (o *LookupIDsInDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupIDsInDomainOperation) {
	if o == nil {
		return
	}
	o.Names = op.Names
	o.Use = op.Use
	o.Return = op.Return
}
func (o *LookupIDsInDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *LookupIDsInDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupIDsInDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenGroupOperation structure represents the SamrOpenGroup operation
type xxx_OpenGroupOperation struct {
	Domain        *Handle `idl:"name:DomainHandle" json:"domain"`
	DesiredAccess uint32  `idl:"name:DesiredAccess" json:"desired_access"`
	GroupID       uint32  `idl:"name:GroupId" json:"group_id"`
	GroupHandle   *Handle `idl:"name:GroupHandle" json:"group_handle"`
	Return        int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenGroupOperation) OpNum() int { return 19 }

func (o *xxx_OpenGroupOperation) OpName() string { return "/samr/v1/SamrOpenGroup" }

func (o *xxx_OpenGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	// GroupId {in} (1:(uint32))
	{
		if err := w.WriteData(o.GroupID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	// GroupId {in} (1:(uint32))
	{
		if err := w.ReadData(&o.GroupID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// GroupHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle != nil {
			if err := o.GroupHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// GroupHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle == nil {
			o.GroupHandle = &Handle{}
		}
		if err := o.GroupHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenGroupRequest structure represents the SamrOpenGroup operation request
type OpenGroupRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// DesiredAccess: An ACCESS_MASK that indicates the requested access for the returned
	// handle. See section 2.2.1.5 for a list of group access values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
	// GroupId: A RID of a group.
	GroupID uint32 `idl:"name:GroupId" json:"group_id"`
}

func (o *OpenGroupRequest) xxx_ToOp(ctx context.Context) *xxx_OpenGroupOperation {
	if o == nil {
		return &xxx_OpenGroupOperation{}
	}
	return &xxx_OpenGroupOperation{
		Domain:        o.Domain,
		DesiredAccess: o.DesiredAccess,
		GroupID:       o.GroupID,
	}
}

func (o *OpenGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenGroupOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DesiredAccess = op.DesiredAccess
	o.GroupID = op.GroupID
}
func (o *OpenGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenGroupResponse structure represents the SamrOpenGroup operation response
type OpenGroupResponse struct {
	// GroupHandle: An RPC context handle, as specified in section 2.2.7.2.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	// Return: The SamrOpenGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenGroupResponse) xxx_ToOp(ctx context.Context) *xxx_OpenGroupOperation {
	if o == nil {
		return &xxx_OpenGroupOperation{}
	}
	return &xxx_OpenGroupOperation{
		GroupHandle: o.GroupHandle,
		Return:      o.Return,
	}
}

func (o *OpenGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenGroupOperation) {
	if o == nil {
		return
	}
	o.GroupHandle = op.GroupHandle
	o.Return = op.Return
}
func (o *OpenGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInformationGroupOperation structure represents the SamrQueryInformationGroup operation
type xxx_QueryInformationGroupOperation struct {
	GroupHandle           *Handle               `idl:"name:GroupHandle" json:"group_handle"`
	GroupInformationClass GroupInformationClass `idl:"name:GroupInformationClass" json:"group_information_class"`
	Buffer                *GroupInfoBuffer      `idl:"name:Buffer;switch_is:GroupInformationClass" json:"buffer"`
	Return                int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInformationGroupOperation) OpNum() int { return 20 }

func (o *xxx_QueryInformationGroupOperation) OpName() string {
	return "/samr/v1/SamrQueryInformationGroup"
}

func (o *xxx_QueryInformationGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle != nil {
			if err := o.GroupHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// GroupInformationClass {in} (1:{alias=GROUP_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.GroupInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle == nil {
			o.GroupHandle = &Handle{}
		}
		if err := o.GroupHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// GroupInformationClass {in} (1:{alias=GROUP_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.GroupInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=GROUP_INFORMATION_CLASS}(enum), alias=PSAMPR_GROUP_INFO_BUFFER}*(1))(3:{switch_type={alias=GROUP_INFORMATION_CLASS}(enum), alias=SAMPR_GROUP_INFO_BUFFER}(union))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swBuffer := uint16(o.GroupInformationClass)
				if o.Buffer != nil {
					if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
						return err
					}
				} else {
					if err := (&GroupInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=GROUP_INFORMATION_CLASS}(enum), alias=PSAMPR_GROUP_INFO_BUFFER,pointer=ref}*(1))(3:{switch_type={alias=GROUP_INFORMATION_CLASS}(enum), alias=SAMPR_GROUP_INFO_BUFFER}(union))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &GroupInfoBuffer{}
			}
			_swBuffer := uint16(o.GroupInformationClass)
			if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**GroupInfoBuffer) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryInformationGroupRequest structure represents the SamrQueryInformationGroup operation request
type QueryInformationGroupRequest struct {
	// GroupHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a group object.
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	// GroupInformationClass: An enumeration indicating which attributes to return. See
	// section 2.2.4.6 for a listing of possible values.
	GroupInformationClass GroupInformationClass `idl:"name:GroupInformationClass" json:"group_information_class"`
}

func (o *QueryInformationGroupRequest) xxx_ToOp(ctx context.Context) *xxx_QueryInformationGroupOperation {
	if o == nil {
		return &xxx_QueryInformationGroupOperation{}
	}
	return &xxx_QueryInformationGroupOperation{
		GroupHandle:           o.GroupHandle,
		GroupInformationClass: o.GroupInformationClass,
	}
}

func (o *QueryInformationGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationGroupOperation) {
	if o == nil {
		return
	}
	o.GroupHandle = op.GroupHandle
	o.GroupInformationClass = op.GroupInformationClass
}
func (o *QueryInformationGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInformationGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInformationGroupResponse structure represents the SamrQueryInformationGroup operation response
type QueryInformationGroupResponse struct {
	// Buffer: The requested attributes on output. See section 2.2.4.7 for structure details.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Buffer *GroupInfoBuffer `idl:"name:Buffer;switch_is:GroupInformationClass" json:"buffer"`
	// Return: The SamrQueryInformationGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryInformationGroupResponse) xxx_ToOp(ctx context.Context) *xxx_QueryInformationGroupOperation {
	if o == nil {
		return &xxx_QueryInformationGroupOperation{}
	}
	return &xxx_QueryInformationGroupOperation{
		Buffer: o.Buffer,
		Return: o.Return,
	}
}

func (o *QueryInformationGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationGroupOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *QueryInformationGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInformationGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInformationGroupOperation structure represents the SamrSetInformationGroup operation
type xxx_SetInformationGroupOperation struct {
	GroupHandle           *Handle               `idl:"name:GroupHandle" json:"group_handle"`
	GroupInformationClass GroupInformationClass `idl:"name:GroupInformationClass" json:"group_information_class"`
	Buffer                *GroupInfoBuffer      `idl:"name:Buffer;switch_is:GroupInformationClass" json:"buffer"`
	Return                int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInformationGroupOperation) OpNum() int { return 21 }

func (o *xxx_SetInformationGroupOperation) OpName() string { return "/samr/v1/SamrSetInformationGroup" }

func (o *xxx_SetInformationGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle != nil {
			if err := o.GroupHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// GroupInformationClass {in} (1:{alias=GROUP_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.GroupInformationClass)); err != nil {
			return err
		}
	}
	// Buffer {in} (1:{switch_type={alias=GROUP_INFORMATION_CLASS}(enum), alias=PSAMPR_GROUP_INFO_BUFFER}*(1))(2:{switch_type={alias=GROUP_INFORMATION_CLASS}(enum), alias=SAMPR_GROUP_INFO_BUFFER}(union))
	{
		_swBuffer := uint16(o.GroupInformationClass)
		if o.Buffer != nil {
			if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		} else {
			if err := (&GroupInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle == nil {
			o.GroupHandle = &Handle{}
		}
		if err := o.GroupHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// GroupInformationClass {in} (1:{alias=GROUP_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.GroupInformationClass)); err != nil {
			return err
		}
	}
	// Buffer {in} (1:{switch_type={alias=GROUP_INFORMATION_CLASS}(enum), alias=PSAMPR_GROUP_INFO_BUFFER,pointer=ref}*(1))(2:{switch_type={alias=GROUP_INFORMATION_CLASS}(enum), alias=SAMPR_GROUP_INFO_BUFFER}(union))
	{
		if o.Buffer == nil {
			o.Buffer = &GroupInfoBuffer{}
		}
		_swBuffer := uint16(o.GroupInformationClass)
		if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInformationGroupRequest structure represents the SamrSetInformationGroup operation request
type SetInformationGroupRequest struct {
	// GroupHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a group object.
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	// GroupInformationClass: An enumeration indicating which attributes to update. See
	// section 2.2.4.6 for a listing of possible values.
	GroupInformationClass GroupInformationClass `idl:"name:GroupInformationClass" json:"group_information_class"`
	// Buffer: The requested attributes and values to update. See section 2.2.4.7 for structure
	// details.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Buffer *GroupInfoBuffer `idl:"name:Buffer;switch_is:GroupInformationClass" json:"buffer"`
}

func (o *SetInformationGroupRequest) xxx_ToOp(ctx context.Context) *xxx_SetInformationGroupOperation {
	if o == nil {
		return &xxx_SetInformationGroupOperation{}
	}
	return &xxx_SetInformationGroupOperation{
		GroupHandle:           o.GroupHandle,
		GroupInformationClass: o.GroupInformationClass,
		Buffer:                o.Buffer,
	}
}

func (o *SetInformationGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInformationGroupOperation) {
	if o == nil {
		return
	}
	o.GroupHandle = op.GroupHandle
	o.GroupInformationClass = op.GroupInformationClass
	o.Buffer = op.Buffer
}
func (o *SetInformationGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetInformationGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInformationGroupResponse structure represents the SamrSetInformationGroup operation response
type SetInformationGroupResponse struct {
	// Return: The SamrSetInformationGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInformationGroupResponse) xxx_ToOp(ctx context.Context) *xxx_SetInformationGroupOperation {
	if o == nil {
		return &xxx_SetInformationGroupOperation{}
	}
	return &xxx_SetInformationGroupOperation{
		Return: o.Return,
	}
}

func (o *SetInformationGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInformationGroupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetInformationGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetInformationGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddMemberToGroupOperation structure represents the SamrAddMemberToGroup operation
type xxx_AddMemberToGroupOperation struct {
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	MemberID    uint32  `idl:"name:MemberId" json:"member_id"`
	Attributes  uint32  `idl:"name:Attributes" json:"attributes"`
	Return      int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_AddMemberToGroupOperation) OpNum() int { return 22 }

func (o *xxx_AddMemberToGroupOperation) OpName() string { return "/samr/v1/SamrAddMemberToGroup" }

func (o *xxx_AddMemberToGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMemberToGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle != nil {
			if err := o.GroupHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MemberId {in} (1:(uint32))
	{
		if err := w.WriteData(o.MemberID); err != nil {
			return err
		}
	}
	// Attributes {in} (1:(uint32))
	{
		if err := w.WriteData(o.Attributes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMemberToGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle == nil {
			o.GroupHandle = &Handle{}
		}
		if err := o.GroupHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MemberId {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MemberID); err != nil {
			return err
		}
	}
	// Attributes {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Attributes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMemberToGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMemberToGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMemberToGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddMemberToGroupRequest structure represents the SamrAddMemberToGroup operation request
type AddMemberToGroupRequest struct {
	// GroupHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a group object.
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	// MemberId: A RID representing an account to add to the group's membership list.
	MemberID uint32 `idl:"name:MemberId" json:"member_id"`
	// Attributes: The characteristics of the membership relationship. See section 2.2.1.10
	// for legal values and semantics.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Attributes uint32 `idl:"name:Attributes" json:"attributes"`
}

func (o *AddMemberToGroupRequest) xxx_ToOp(ctx context.Context) *xxx_AddMemberToGroupOperation {
	if o == nil {
		return &xxx_AddMemberToGroupOperation{}
	}
	return &xxx_AddMemberToGroupOperation{
		GroupHandle: o.GroupHandle,
		MemberID:    o.MemberID,
		Attributes:  o.Attributes,
	}
}

func (o *AddMemberToGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_AddMemberToGroupOperation) {
	if o == nil {
		return
	}
	o.GroupHandle = op.GroupHandle
	o.MemberID = op.MemberID
	o.Attributes = op.Attributes
}
func (o *AddMemberToGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddMemberToGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddMemberToGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddMemberToGroupResponse structure represents the SamrAddMemberToGroup operation response
type AddMemberToGroupResponse struct {
	// Return: The SamrAddMemberToGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddMemberToGroupResponse) xxx_ToOp(ctx context.Context) *xxx_AddMemberToGroupOperation {
	if o == nil {
		return &xxx_AddMemberToGroupOperation{}
	}
	return &xxx_AddMemberToGroupOperation{
		Return: o.Return,
	}
}

func (o *AddMemberToGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_AddMemberToGroupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddMemberToGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddMemberToGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddMemberToGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteGroupOperation structure represents the SamrDeleteGroup operation
type xxx_DeleteGroupOperation struct {
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	Return      int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteGroupOperation) OpNum() int { return 23 }

func (o *xxx_DeleteGroupOperation) OpName() string { return "/samr/v1/SamrDeleteGroup" }

func (o *xxx_DeleteGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// GroupHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle != nil {
			if err := o.GroupHandle.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_DeleteGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// GroupHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle == nil {
			o.GroupHandle = &Handle{}
		}
		if err := o.GroupHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// GroupHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle != nil {
			if err := o.GroupHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// GroupHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle == nil {
			o.GroupHandle = &Handle{}
		}
		if err := o.GroupHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteGroupRequest structure represents the SamrDeleteGroup operation request
type DeleteGroupRequest struct {
	// GroupHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a group object.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
}

func (o *DeleteGroupRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteGroupOperation {
	if o == nil {
		return &xxx_DeleteGroupOperation{}
	}
	return &xxx_DeleteGroupOperation{
		GroupHandle: o.GroupHandle,
	}
}

func (o *DeleteGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteGroupOperation) {
	if o == nil {
		return
	}
	o.GroupHandle = op.GroupHandle
}
func (o *DeleteGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteGroupResponse structure represents the SamrDeleteGroup operation response
type DeleteGroupResponse struct {
	// GroupHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a group object.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	// Return: The SamrDeleteGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteGroupResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteGroupOperation {
	if o == nil {
		return &xxx_DeleteGroupOperation{}
	}
	return &xxx_DeleteGroupOperation{
		GroupHandle: o.GroupHandle,
		Return:      o.Return,
	}
}

func (o *DeleteGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteGroupOperation) {
	if o == nil {
		return
	}
	o.GroupHandle = op.GroupHandle
	o.Return = op.Return
}
func (o *DeleteGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveMemberFromGroupOperation structure represents the SamrRemoveMemberFromGroup operation
type xxx_RemoveMemberFromGroupOperation struct {
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	MemberID    uint32  `idl:"name:MemberId" json:"member_id"`
	Return      int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveMemberFromGroupOperation) OpNum() int { return 24 }

func (o *xxx_RemoveMemberFromGroupOperation) OpName() string {
	return "/samr/v1/SamrRemoveMemberFromGroup"
}

func (o *xxx_RemoveMemberFromGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle != nil {
			if err := o.GroupHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MemberId {in} (1:(uint32))
	{
		if err := w.WriteData(o.MemberID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle == nil {
			o.GroupHandle = &Handle{}
		}
		if err := o.GroupHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MemberId {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MemberID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveMemberFromGroupRequest structure represents the SamrRemoveMemberFromGroup operation request
type RemoveMemberFromGroupRequest struct {
	// GroupHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a group object.
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	// MemberId: A RID representing an account to remove from the group's membership list.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	MemberID uint32 `idl:"name:MemberId" json:"member_id"`
}

func (o *RemoveMemberFromGroupRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveMemberFromGroupOperation {
	if o == nil {
		return &xxx_RemoveMemberFromGroupOperation{}
	}
	return &xxx_RemoveMemberFromGroupOperation{
		GroupHandle: o.GroupHandle,
		MemberID:    o.MemberID,
	}
}

func (o *RemoveMemberFromGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveMemberFromGroupOperation) {
	if o == nil {
		return
	}
	o.GroupHandle = op.GroupHandle
	o.MemberID = op.MemberID
}
func (o *RemoveMemberFromGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveMemberFromGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMemberFromGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveMemberFromGroupResponse structure represents the SamrRemoveMemberFromGroup operation response
type RemoveMemberFromGroupResponse struct {
	// Return: The SamrRemoveMemberFromGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveMemberFromGroupResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveMemberFromGroupOperation {
	if o == nil {
		return &xxx_RemoveMemberFromGroupOperation{}
	}
	return &xxx_RemoveMemberFromGroupOperation{
		Return: o.Return,
	}
}

func (o *RemoveMemberFromGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveMemberFromGroupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveMemberFromGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveMemberFromGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMemberFromGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMembersInGroupOperation structure represents the SamrGetMembersInGroup operation
type xxx_GetMembersInGroupOperation struct {
	GroupHandle *Handle           `idl:"name:GroupHandle" json:"group_handle"`
	Members     *GetMembersBuffer `idl:"name:Members" json:"members"`
	Return      int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMembersInGroupOperation) OpNum() int { return 25 }

func (o *xxx_GetMembersInGroupOperation) OpName() string { return "/samr/v1/SamrGetMembersInGroup" }

func (o *xxx_GetMembersInGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMembersInGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle != nil {
			if err := o.GroupHandle.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetMembersInGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle == nil {
			o.GroupHandle = &Handle{}
		}
		if err := o.GroupHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMembersInGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMembersInGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Members {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_GET_MEMBERS_BUFFER}*(1))(3:{alias=SAMPR_GET_MEMBERS_BUFFER}(struct))
	{
		if o.Members != nil {
			_ptr_Members := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Members != nil {
					if err := o.Members.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GetMembersBuffer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Members, _ptr_Members); err != nil {
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
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMembersInGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Members {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_GET_MEMBERS_BUFFER,pointer=ref}*(1))(3:{alias=SAMPR_GET_MEMBERS_BUFFER}(struct))
	{
		_ptr_Members := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Members == nil {
				o.Members = &GetMembersBuffer{}
			}
			if err := o.Members.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Members := func(ptr interface{}) { o.Members = *ptr.(**GetMembersBuffer) }
		if err := w.ReadPointer(&o.Members, _s_Members, _ptr_Members); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetMembersInGroupRequest structure represents the SamrGetMembersInGroup operation request
type GetMembersInGroupRequest struct {
	// GroupHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a group object.
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
}

func (o *GetMembersInGroupRequest) xxx_ToOp(ctx context.Context) *xxx_GetMembersInGroupOperation {
	if o == nil {
		return &xxx_GetMembersInGroupOperation{}
	}
	return &xxx_GetMembersInGroupOperation{
		GroupHandle: o.GroupHandle,
	}
}

func (o *GetMembersInGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMembersInGroupOperation) {
	if o == nil {
		return
	}
	o.GroupHandle = op.GroupHandle
}
func (o *GetMembersInGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMembersInGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMembersInGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMembersInGroupResponse structure represents the SamrGetMembersInGroup operation response
type GetMembersInGroupResponse struct {
	// Members: A structure containing an array of RIDs, as well as an array of attribute
	// values.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Members *GetMembersBuffer `idl:"name:Members" json:"members"`
	// Return: The SamrGetMembersInGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMembersInGroupResponse) xxx_ToOp(ctx context.Context) *xxx_GetMembersInGroupOperation {
	if o == nil {
		return &xxx_GetMembersInGroupOperation{}
	}
	return &xxx_GetMembersInGroupOperation{
		Members: o.Members,
		Return:  o.Return,
	}
}

func (o *GetMembersInGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMembersInGroupOperation) {
	if o == nil {
		return
	}
	o.Members = op.Members
	o.Return = op.Return
}
func (o *GetMembersInGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMembersInGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMembersInGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMemberAttributesOfGroupOperation structure represents the SamrSetMemberAttributesOfGroup operation
type xxx_SetMemberAttributesOfGroupOperation struct {
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	MemberID    uint32  `idl:"name:MemberId" json:"member_id"`
	Attributes  uint32  `idl:"name:Attributes" json:"attributes"`
	Return      int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMemberAttributesOfGroupOperation) OpNum() int { return 26 }

func (o *xxx_SetMemberAttributesOfGroupOperation) OpName() string {
	return "/samr/v1/SamrSetMemberAttributesOfGroup"
}

func (o *xxx_SetMemberAttributesOfGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMemberAttributesOfGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle != nil {
			if err := o.GroupHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MemberId {in} (1:(uint32))
	{
		if err := w.WriteData(o.MemberID); err != nil {
			return err
		}
	}
	// Attributes {in} (1:(uint32))
	{
		if err := w.WriteData(o.Attributes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMemberAttributesOfGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// GroupHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.GroupHandle == nil {
			o.GroupHandle = &Handle{}
		}
		if err := o.GroupHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MemberId {in} (1:(uint32))
	{
		if err := w.ReadData(&o.MemberID); err != nil {
			return err
		}
	}
	// Attributes {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Attributes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMemberAttributesOfGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMemberAttributesOfGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMemberAttributesOfGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetMemberAttributesOfGroupRequest structure represents the SamrSetMemberAttributesOfGroup operation request
type SetMemberAttributesOfGroupRequest struct {
	// GroupHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a group object.
	GroupHandle *Handle `idl:"name:GroupHandle" json:"group_handle"`
	// MemberId: A RID that represents a member of a group (which is a user or machine account).
	MemberID uint32 `idl:"name:MemberId" json:"member_id"`
	// Attributes: The characteristics of the membership relationship. For legal values,
	// see section 2.2.1.10.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Attributes uint32 `idl:"name:Attributes" json:"attributes"`
}

func (o *SetMemberAttributesOfGroupRequest) xxx_ToOp(ctx context.Context) *xxx_SetMemberAttributesOfGroupOperation {
	if o == nil {
		return &xxx_SetMemberAttributesOfGroupOperation{}
	}
	return &xxx_SetMemberAttributesOfGroupOperation{
		GroupHandle: o.GroupHandle,
		MemberID:    o.MemberID,
		Attributes:  o.Attributes,
	}
}

func (o *SetMemberAttributesOfGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMemberAttributesOfGroupOperation) {
	if o == nil {
		return
	}
	o.GroupHandle = op.GroupHandle
	o.MemberID = op.MemberID
	o.Attributes = op.Attributes
}
func (o *SetMemberAttributesOfGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetMemberAttributesOfGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMemberAttributesOfGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMemberAttributesOfGroupResponse structure represents the SamrSetMemberAttributesOfGroup operation response
type SetMemberAttributesOfGroupResponse struct {
	// Return: The SamrSetMemberAttributesOfGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMemberAttributesOfGroupResponse) xxx_ToOp(ctx context.Context) *xxx_SetMemberAttributesOfGroupOperation {
	if o == nil {
		return &xxx_SetMemberAttributesOfGroupOperation{}
	}
	return &xxx_SetMemberAttributesOfGroupOperation{
		Return: o.Return,
	}
}

func (o *SetMemberAttributesOfGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMemberAttributesOfGroupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetMemberAttributesOfGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetMemberAttributesOfGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMemberAttributesOfGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenAliasOperation structure represents the SamrOpenAlias operation
type xxx_OpenAliasOperation struct {
	Domain        *Handle `idl:"name:DomainHandle" json:"domain"`
	DesiredAccess uint32  `idl:"name:DesiredAccess" json:"desired_access"`
	AliasID       uint32  `idl:"name:AliasId" json:"alias_id"`
	AliasHandle   *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	Return        int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenAliasOperation) OpNum() int { return 27 }

func (o *xxx_OpenAliasOperation) OpName() string { return "/samr/v1/SamrOpenAlias" }

func (o *xxx_OpenAliasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenAliasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	// AliasId {in} (1:(uint32))
	{
		if err := w.WriteData(o.AliasID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenAliasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	// AliasId {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AliasID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenAliasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenAliasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// AliasHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenAliasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenAliasRequest structure represents the SamrOpenAlias operation request
type OpenAliasRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// DesiredAccess: An ACCESS_MASK that indicates the requested access for the returned
	// handle. See section 2.2.1.6 for a list of alias access values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
	// AliasId: A RID of an alias.
	AliasID uint32 `idl:"name:AliasId" json:"alias_id"`
}

func (o *OpenAliasRequest) xxx_ToOp(ctx context.Context) *xxx_OpenAliasOperation {
	if o == nil {
		return &xxx_OpenAliasOperation{}
	}
	return &xxx_OpenAliasOperation{
		Domain:        o.Domain,
		DesiredAccess: o.DesiredAccess,
		AliasID:       o.AliasID,
	}
}

func (o *OpenAliasRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenAliasOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DesiredAccess = op.DesiredAccess
	o.AliasID = op.AliasID
}
func (o *OpenAliasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenAliasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenAliasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenAliasResponse structure represents the SamrOpenAlias operation response
type OpenAliasResponse struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	// Return: The SamrOpenAlias return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenAliasResponse) xxx_ToOp(ctx context.Context) *xxx_OpenAliasOperation {
	if o == nil {
		return &xxx_OpenAliasOperation{}
	}
	return &xxx_OpenAliasOperation{
		AliasHandle: o.AliasHandle,
		Return:      o.Return,
	}
}

func (o *OpenAliasResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenAliasOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
	o.Return = op.Return
}
func (o *OpenAliasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenAliasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenAliasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInformationAliasOperation structure represents the SamrQueryInformationAlias operation
type xxx_QueryInformationAliasOperation struct {
	AliasHandle           *Handle               `idl:"name:AliasHandle" json:"alias_handle"`
	AliasInformationClass AliasInformationClass `idl:"name:AliasInformationClass" json:"alias_information_class"`
	Buffer                *AliasInfoBuffer      `idl:"name:Buffer;switch_is:AliasInformationClass" json:"buffer"`
	Return                int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInformationAliasOperation) OpNum() int { return 28 }

func (o *xxx_QueryInformationAliasOperation) OpName() string {
	return "/samr/v1/SamrQueryInformationAlias"
}

func (o *xxx_QueryInformationAliasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationAliasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// AliasInformationClass {in} (1:{alias=ALIAS_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.AliasInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationAliasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// AliasInformationClass {in} (1:{alias=ALIAS_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.AliasInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationAliasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationAliasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=ALIAS_INFORMATION_CLASS}(enum), alias=PSAMPR_ALIAS_INFO_BUFFER}*(1))(3:{switch_type={alias=ALIAS_INFORMATION_CLASS}(enum), alias=SAMPR_ALIAS_INFO_BUFFER}(union))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swBuffer := uint16(o.AliasInformationClass)
				if o.Buffer != nil {
					if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
						return err
					}
				} else {
					if err := (&AliasInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationAliasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=ALIAS_INFORMATION_CLASS}(enum), alias=PSAMPR_ALIAS_INFO_BUFFER,pointer=ref}*(1))(3:{switch_type={alias=ALIAS_INFORMATION_CLASS}(enum), alias=SAMPR_ALIAS_INFO_BUFFER}(union))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &AliasInfoBuffer{}
			}
			_swBuffer := uint16(o.AliasInformationClass)
			if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**AliasInfoBuffer) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryInformationAliasRequest structure represents the SamrQueryInformationAlias operation request
type QueryInformationAliasRequest struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// an alias object.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	// AliasInformationClass: An enumeration indicating which attributes to return. See
	// section 2.2.5.5 for a listing of possible values.
	AliasInformationClass AliasInformationClass `idl:"name:AliasInformationClass" json:"alias_information_class"`
}

func (o *QueryInformationAliasRequest) xxx_ToOp(ctx context.Context) *xxx_QueryInformationAliasOperation {
	if o == nil {
		return &xxx_QueryInformationAliasOperation{}
	}
	return &xxx_QueryInformationAliasOperation{
		AliasHandle:           o.AliasHandle,
		AliasInformationClass: o.AliasInformationClass,
	}
}

func (o *QueryInformationAliasRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationAliasOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
	o.AliasInformationClass = op.AliasInformationClass
}
func (o *QueryInformationAliasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInformationAliasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationAliasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInformationAliasResponse structure represents the SamrQueryInformationAlias operation response
type QueryInformationAliasResponse struct {
	// Buffer: The requested attributes on output. See section 2.2.5.6 for structure details.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Buffer *AliasInfoBuffer `idl:"name:Buffer;switch_is:AliasInformationClass" json:"buffer"`
	// Return: The SamrQueryInformationAlias return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryInformationAliasResponse) xxx_ToOp(ctx context.Context) *xxx_QueryInformationAliasOperation {
	if o == nil {
		return &xxx_QueryInformationAliasOperation{}
	}
	return &xxx_QueryInformationAliasOperation{
		Buffer: o.Buffer,
		Return: o.Return,
	}
}

func (o *QueryInformationAliasResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationAliasOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *QueryInformationAliasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInformationAliasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationAliasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInformationAliasOperation structure represents the SamrSetInformationAlias operation
type xxx_SetInformationAliasOperation struct {
	AliasHandle           *Handle               `idl:"name:AliasHandle" json:"alias_handle"`
	AliasInformationClass AliasInformationClass `idl:"name:AliasInformationClass" json:"alias_information_class"`
	Buffer                *AliasInfoBuffer      `idl:"name:Buffer;switch_is:AliasInformationClass" json:"buffer"`
	Return                int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInformationAliasOperation) OpNum() int { return 29 }

func (o *xxx_SetInformationAliasOperation) OpName() string { return "/samr/v1/SamrSetInformationAlias" }

func (o *xxx_SetInformationAliasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationAliasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// AliasInformationClass {in} (1:{alias=ALIAS_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.AliasInformationClass)); err != nil {
			return err
		}
	}
	// Buffer {in} (1:{switch_type={alias=ALIAS_INFORMATION_CLASS}(enum), alias=PSAMPR_ALIAS_INFO_BUFFER}*(1))(2:{switch_type={alias=ALIAS_INFORMATION_CLASS}(enum), alias=SAMPR_ALIAS_INFO_BUFFER}(union))
	{
		_swBuffer := uint16(o.AliasInformationClass)
		if o.Buffer != nil {
			if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		} else {
			if err := (&AliasInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationAliasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// AliasInformationClass {in} (1:{alias=ALIAS_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.AliasInformationClass)); err != nil {
			return err
		}
	}
	// Buffer {in} (1:{switch_type={alias=ALIAS_INFORMATION_CLASS}(enum), alias=PSAMPR_ALIAS_INFO_BUFFER,pointer=ref}*(1))(2:{switch_type={alias=ALIAS_INFORMATION_CLASS}(enum), alias=SAMPR_ALIAS_INFO_BUFFER}(union))
	{
		if o.Buffer == nil {
			o.Buffer = &AliasInfoBuffer{}
		}
		_swBuffer := uint16(o.AliasInformationClass)
		if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationAliasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationAliasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationAliasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInformationAliasRequest structure represents the SamrSetInformationAlias operation request
type SetInformationAliasRequest struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// an alias object.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	// AliasInformationClass: An enumeration indicating which attributes to update. See
	// section 2.2.5.5 for a listing of possible values.
	AliasInformationClass AliasInformationClass `idl:"name:AliasInformationClass" json:"alias_information_class"`
	// Buffer: The requested attributes and values to update. See section 2.2.5.6 for structure
	// details.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Buffer *AliasInfoBuffer `idl:"name:Buffer;switch_is:AliasInformationClass" json:"buffer"`
}

func (o *SetInformationAliasRequest) xxx_ToOp(ctx context.Context) *xxx_SetInformationAliasOperation {
	if o == nil {
		return &xxx_SetInformationAliasOperation{}
	}
	return &xxx_SetInformationAliasOperation{
		AliasHandle:           o.AliasHandle,
		AliasInformationClass: o.AliasInformationClass,
		Buffer:                o.Buffer,
	}
}

func (o *SetInformationAliasRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInformationAliasOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
	o.AliasInformationClass = op.AliasInformationClass
	o.Buffer = op.Buffer
}
func (o *SetInformationAliasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetInformationAliasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationAliasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInformationAliasResponse structure represents the SamrSetInformationAlias operation response
type SetInformationAliasResponse struct {
	// Return: The SamrSetInformationAlias return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInformationAliasResponse) xxx_ToOp(ctx context.Context) *xxx_SetInformationAliasOperation {
	if o == nil {
		return &xxx_SetInformationAliasOperation{}
	}
	return &xxx_SetInformationAliasOperation{
		Return: o.Return,
	}
}

func (o *SetInformationAliasResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInformationAliasOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetInformationAliasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetInformationAliasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationAliasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteAliasOperation structure represents the SamrDeleteAlias operation
type xxx_DeleteAliasOperation struct {
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	Return      int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteAliasOperation) OpNum() int { return 30 }

func (o *xxx_DeleteAliasOperation) OpName() string { return "/samr/v1/SamrDeleteAlias" }

func (o *xxx_DeleteAliasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAliasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AliasHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_DeleteAliasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAliasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAliasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// AliasHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAliasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteAliasRequest structure represents the SamrDeleteAlias operation request
type DeleteAliasRequest struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// an alias object.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
}

func (o *DeleteAliasRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteAliasOperation {
	if o == nil {
		return &xxx_DeleteAliasOperation{}
	}
	return &xxx_DeleteAliasOperation{
		AliasHandle: o.AliasHandle,
	}
}

func (o *DeleteAliasRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteAliasOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
}
func (o *DeleteAliasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteAliasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteAliasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteAliasResponse structure represents the SamrDeleteAlias operation response
type DeleteAliasResponse struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// an alias object.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	// Return: The SamrDeleteAlias return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteAliasResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteAliasOperation {
	if o == nil {
		return &xxx_DeleteAliasOperation{}
	}
	return &xxx_DeleteAliasOperation{
		AliasHandle: o.AliasHandle,
		Return:      o.Return,
	}
}

func (o *DeleteAliasResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteAliasOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
	o.Return = op.Return
}
func (o *DeleteAliasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteAliasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteAliasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddMemberToAliasOperation structure represents the SamrAddMemberToAlias operation
type xxx_AddMemberToAliasOperation struct {
	AliasHandle *Handle   `idl:"name:AliasHandle" json:"alias_handle"`
	MemberID    *dtyp.SID `idl:"name:MemberId" json:"member_id"`
	Return      int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_AddMemberToAliasOperation) OpNum() int { return 31 }

func (o *xxx_AddMemberToAliasOperation) OpName() string { return "/samr/v1/SamrAddMemberToAlias" }

func (o *xxx_AddMemberToAliasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMemberToAliasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MemberId {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.MemberID != nil {
			if err := o.MemberID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_AddMemberToAliasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MemberId {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.MemberID == nil {
			o.MemberID = &dtyp.SID{}
		}
		if err := o.MemberID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMemberToAliasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMemberToAliasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMemberToAliasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddMemberToAliasRequest structure represents the SamrAddMemberToAlias operation request
type AddMemberToAliasRequest struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// an alias object.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	// MemberId: The SID of an account to add to the alias.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	MemberID *dtyp.SID `idl:"name:MemberId" json:"member_id"`
}

func (o *AddMemberToAliasRequest) xxx_ToOp(ctx context.Context) *xxx_AddMemberToAliasOperation {
	if o == nil {
		return &xxx_AddMemberToAliasOperation{}
	}
	return &xxx_AddMemberToAliasOperation{
		AliasHandle: o.AliasHandle,
		MemberID:    o.MemberID,
	}
}

func (o *AddMemberToAliasRequest) xxx_FromOp(ctx context.Context, op *xxx_AddMemberToAliasOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
	o.MemberID = op.MemberID
}
func (o *AddMemberToAliasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddMemberToAliasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddMemberToAliasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddMemberToAliasResponse structure represents the SamrAddMemberToAlias operation response
type AddMemberToAliasResponse struct {
	// Return: The SamrAddMemberToAlias return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddMemberToAliasResponse) xxx_ToOp(ctx context.Context) *xxx_AddMemberToAliasOperation {
	if o == nil {
		return &xxx_AddMemberToAliasOperation{}
	}
	return &xxx_AddMemberToAliasOperation{
		Return: o.Return,
	}
}

func (o *AddMemberToAliasResponse) xxx_FromOp(ctx context.Context, op *xxx_AddMemberToAliasOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddMemberToAliasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddMemberToAliasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddMemberToAliasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveMemberFromAliasOperation structure represents the SamrRemoveMemberFromAlias operation
type xxx_RemoveMemberFromAliasOperation struct {
	AliasHandle *Handle   `idl:"name:AliasHandle" json:"alias_handle"`
	MemberID    *dtyp.SID `idl:"name:MemberId" json:"member_id"`
	Return      int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveMemberFromAliasOperation) OpNum() int { return 32 }

func (o *xxx_RemoveMemberFromAliasOperation) OpName() string {
	return "/samr/v1/SamrRemoveMemberFromAlias"
}

func (o *xxx_RemoveMemberFromAliasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromAliasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MemberId {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.MemberID != nil {
			if err := o.MemberID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RemoveMemberFromAliasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MemberId {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.MemberID == nil {
			o.MemberID = &dtyp.SID{}
		}
		if err := o.MemberID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromAliasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromAliasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromAliasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveMemberFromAliasRequest structure represents the SamrRemoveMemberFromAlias operation request
type RemoveMemberFromAliasRequest struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// an alias object.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	// MemberId: The SID of an account to remove from the alias.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	MemberID *dtyp.SID `idl:"name:MemberId" json:"member_id"`
}

func (o *RemoveMemberFromAliasRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveMemberFromAliasOperation {
	if o == nil {
		return &xxx_RemoveMemberFromAliasOperation{}
	}
	return &xxx_RemoveMemberFromAliasOperation{
		AliasHandle: o.AliasHandle,
		MemberID:    o.MemberID,
	}
}

func (o *RemoveMemberFromAliasRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveMemberFromAliasOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
	o.MemberID = op.MemberID
}
func (o *RemoveMemberFromAliasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveMemberFromAliasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMemberFromAliasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveMemberFromAliasResponse structure represents the SamrRemoveMemberFromAlias operation response
type RemoveMemberFromAliasResponse struct {
	// Return: The SamrRemoveMemberFromAlias return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveMemberFromAliasResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveMemberFromAliasOperation {
	if o == nil {
		return &xxx_RemoveMemberFromAliasOperation{}
	}
	return &xxx_RemoveMemberFromAliasOperation{
		Return: o.Return,
	}
}

func (o *RemoveMemberFromAliasResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveMemberFromAliasOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveMemberFromAliasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveMemberFromAliasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMemberFromAliasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMembersInAliasOperation structure represents the SamrGetMembersInAlias operation
type xxx_GetMembersInAliasOperation struct {
	AliasHandle *Handle      `idl:"name:AliasHandle" json:"alias_handle"`
	Members     *SIDArrayOut `idl:"name:Members" json:"members"`
	Return      int32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMembersInAliasOperation) OpNum() int { return 33 }

func (o *xxx_GetMembersInAliasOperation) OpName() string { return "/samr/v1/SamrGetMembersInAlias" }

func (o *xxx_GetMembersInAliasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMembersInAliasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetMembersInAliasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMembersInAliasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMembersInAliasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Members {out} (1:{alias=PSAMPR_PSID_ARRAY_OUT}*(1))(2:{alias=SAMPR_PSID_ARRAY_OUT}(struct))
	{
		if o.Members != nil {
			if err := o.Members.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SIDArrayOut{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMembersInAliasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Members {out} (1:{alias=PSAMPR_PSID_ARRAY_OUT,pointer=ref}*(1))(2:{alias=SAMPR_PSID_ARRAY_OUT}(struct))
	{
		if o.Members == nil {
			o.Members = &SIDArrayOut{}
		}
		if err := o.Members.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetMembersInAliasRequest structure represents the SamrGetMembersInAlias operation request
type GetMembersInAliasRequest struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// an alias object.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
}

func (o *GetMembersInAliasRequest) xxx_ToOp(ctx context.Context) *xxx_GetMembersInAliasOperation {
	if o == nil {
		return &xxx_GetMembersInAliasOperation{}
	}
	return &xxx_GetMembersInAliasOperation{
		AliasHandle: o.AliasHandle,
	}
}

func (o *GetMembersInAliasRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMembersInAliasOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
}
func (o *GetMembersInAliasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMembersInAliasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMembersInAliasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMembersInAliasResponse structure represents the SamrGetMembersInAlias operation response
type GetMembersInAliasResponse struct {
	// Members: A structure containing an array of SIDs that represent the membership list
	// of the alias referenced by AliasHandle.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Members *SIDArrayOut `idl:"name:Members" json:"members"`
	// Return: The SamrGetMembersInAlias return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMembersInAliasResponse) xxx_ToOp(ctx context.Context) *xxx_GetMembersInAliasOperation {
	if o == nil {
		return &xxx_GetMembersInAliasOperation{}
	}
	return &xxx_GetMembersInAliasOperation{
		Members: o.Members,
		Return:  o.Return,
	}
}

func (o *GetMembersInAliasResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMembersInAliasOperation) {
	if o == nil {
		return
	}
	o.Members = op.Members
	o.Return = op.Return
}
func (o *GetMembersInAliasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMembersInAliasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMembersInAliasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenUserOperation structure represents the SamrOpenUser operation
type xxx_OpenUserOperation struct {
	Domain        *Handle `idl:"name:DomainHandle" json:"domain"`
	DesiredAccess uint32  `idl:"name:DesiredAccess" json:"desired_access"`
	UserID        uint32  `idl:"name:UserId" json:"user_id"`
	UserHandle    *Handle `idl:"name:UserHandle" json:"user_handle"`
	Return        int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenUserOperation) OpNum() int { return 34 }

func (o *xxx_OpenUserOperation) OpName() string { return "/samr/v1/SamrOpenUser" }

func (o *xxx_OpenUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	// UserId {in} (1:(uint32))
	{
		if err := w.WriteData(o.UserID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	// UserId {in} (1:(uint32))
	{
		if err := w.ReadData(&o.UserID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// UserHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// UserHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenUserRequest structure represents the SamrOpenUser operation request
type OpenUserRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// DesiredAccess: An ACCESS_MASK that indicates the requested access for the returned
	// handle. See section 2.2.1.7 for a list of user access values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
	// UserId: A RID of a user account.
	UserID uint32 `idl:"name:UserId" json:"user_id"`
}

func (o *OpenUserRequest) xxx_ToOp(ctx context.Context) *xxx_OpenUserOperation {
	if o == nil {
		return &xxx_OpenUserOperation{}
	}
	return &xxx_OpenUserOperation{
		Domain:        o.Domain,
		DesiredAccess: o.DesiredAccess,
		UserID:        o.UserID,
	}
}

func (o *OpenUserRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenUserOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DesiredAccess = op.DesiredAccess
	o.UserID = op.UserID
}
func (o *OpenUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenUserResponse structure represents the SamrOpenUser operation response
type OpenUserResponse struct {
	// UserHandle: An RPC context handle, as specified in section 2.2.7.2.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
	// Return: The SamrOpenUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenUserResponse) xxx_ToOp(ctx context.Context) *xxx_OpenUserOperation {
	if o == nil {
		return &xxx_OpenUserOperation{}
	}
	return &xxx_OpenUserOperation{
		UserHandle: o.UserHandle,
		Return:     o.Return,
	}
}

func (o *OpenUserResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenUserOperation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
	o.Return = op.Return
}
func (o *OpenUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteUserOperation structure represents the SamrDeleteUser operation
type xxx_DeleteUserOperation struct {
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
	Return     int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteUserOperation) OpNum() int { return 35 }

func (o *xxx_DeleteUserOperation) OpName() string { return "/samr/v1/SamrDeleteUser" }

func (o *xxx_DeleteUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UserHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_DeleteUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UserHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// UserHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// UserHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteUserRequest structure represents the SamrDeleteUser operation request
type DeleteUserRequest struct {
	// UserHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a user object.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
}

func (o *DeleteUserRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteUserOperation {
	if o == nil {
		return &xxx_DeleteUserOperation{}
	}
	return &xxx_DeleteUserOperation{
		UserHandle: o.UserHandle,
	}
}

func (o *DeleteUserRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteUserOperation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
}
func (o *DeleteUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteUserResponse structure represents the SamrDeleteUser operation response
type DeleteUserResponse struct {
	// UserHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a user object.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
	// Return: The SamrDeleteUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteUserResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteUserOperation {
	if o == nil {
		return &xxx_DeleteUserOperation{}
	}
	return &xxx_DeleteUserOperation{
		UserHandle: o.UserHandle,
		Return:     o.Return,
	}
}

func (o *DeleteUserResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteUserOperation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
	o.Return = op.Return
}
func (o *DeleteUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInformationUserOperation structure represents the SamrQueryInformationUser operation
type xxx_QueryInformationUserOperation struct {
	UserHandle           *Handle              `idl:"name:UserHandle" json:"user_handle"`
	UserInformationClass UserInformationClass `idl:"name:UserInformationClass" json:"user_information_class"`
	Buffer               *UserInfoBuffer      `idl:"name:Buffer;switch_is:UserInformationClass" json:"buffer"`
	Return               int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInformationUserOperation) OpNum() int { return 36 }

func (o *xxx_QueryInformationUserOperation) OpName() string {
	return "/samr/v1/SamrQueryInformationUser"
}

func (o *xxx_QueryInformationUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// UserInformationClass {in} (1:{alias=USER_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.UserInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// UserInformationClass {in} (1:{alias=USER_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.UserInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=PSAMPR_USER_INFO_BUFFER}*(1))(3:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=SAMPR_USER_INFO_BUFFER}(union))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swBuffer := uint16(o.UserInformationClass)
				if o.Buffer != nil {
					if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
						return err
					}
				} else {
					if err := (&UserInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=PSAMPR_USER_INFO_BUFFER,pointer=ref}*(1))(3:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=SAMPR_USER_INFO_BUFFER}(union))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &UserInfoBuffer{}
			}
			_swBuffer := uint16(o.UserInformationClass)
			if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**UserInfoBuffer) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryInformationUserRequest structure represents the SamrQueryInformationUser operation request
type QueryInformationUserRequest struct {
	UserHandle           *Handle              `idl:"name:UserHandle" json:"user_handle"`
	UserInformationClass UserInformationClass `idl:"name:UserInformationClass" json:"user_information_class"`
}

func (o *QueryInformationUserRequest) xxx_ToOp(ctx context.Context) *xxx_QueryInformationUserOperation {
	if o == nil {
		return &xxx_QueryInformationUserOperation{}
	}
	return &xxx_QueryInformationUserOperation{
		UserHandle:           o.UserHandle,
		UserInformationClass: o.UserInformationClass,
	}
}

func (o *QueryInformationUserRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationUserOperation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
	o.UserInformationClass = op.UserInformationClass
}
func (o *QueryInformationUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInformationUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInformationUserResponse structure represents the SamrQueryInformationUser operation response
type QueryInformationUserResponse struct {
	Buffer *UserInfoBuffer `idl:"name:Buffer;switch_is:UserInformationClass" json:"buffer"`
	// Return: The SamrQueryInformationUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryInformationUserResponse) xxx_ToOp(ctx context.Context) *xxx_QueryInformationUserOperation {
	if o == nil {
		return &xxx_QueryInformationUserOperation{}
	}
	return &xxx_QueryInformationUserOperation{
		Buffer: o.Buffer,
		Return: o.Return,
	}
}

func (o *QueryInformationUserResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationUserOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *QueryInformationUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInformationUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInformationUserOperation structure represents the SamrSetInformationUser operation
type xxx_SetInformationUserOperation struct {
	UserHandle           *Handle              `idl:"name:UserHandle" json:"user_handle"`
	UserInformationClass UserInformationClass `idl:"name:UserInformationClass" json:"user_information_class"`
	Buffer               *UserInfoBuffer      `idl:"name:Buffer;switch_is:UserInformationClass" json:"buffer"`
	Return               int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInformationUserOperation) OpNum() int { return 37 }

func (o *xxx_SetInformationUserOperation) OpName() string { return "/samr/v1/SamrSetInformationUser" }

func (o *xxx_SetInformationUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// UserInformationClass {in} (1:{alias=USER_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.UserInformationClass)); err != nil {
			return err
		}
	}
	// Buffer {in} (1:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=PSAMPR_USER_INFO_BUFFER}*(1))(2:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=SAMPR_USER_INFO_BUFFER}(union))
	{
		_swBuffer := uint16(o.UserInformationClass)
		if o.Buffer != nil {
			if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// UserInformationClass {in} (1:{alias=USER_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.UserInformationClass)); err != nil {
			return err
		}
	}
	// Buffer {in} (1:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=PSAMPR_USER_INFO_BUFFER,pointer=ref}*(1))(2:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=SAMPR_USER_INFO_BUFFER}(union))
	{
		if o.Buffer == nil {
			o.Buffer = &UserInfoBuffer{}
		}
		_swBuffer := uint16(o.UserInformationClass)
		if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInformationUserRequest structure represents the SamrSetInformationUser operation request
type SetInformationUserRequest struct {
	UserHandle           *Handle              `idl:"name:UserHandle" json:"user_handle"`
	UserInformationClass UserInformationClass `idl:"name:UserInformationClass" json:"user_information_class"`
	Buffer               *UserInfoBuffer      `idl:"name:Buffer;switch_is:UserInformationClass" json:"buffer"`
}

func (o *SetInformationUserRequest) xxx_ToOp(ctx context.Context) *xxx_SetInformationUserOperation {
	if o == nil {
		return &xxx_SetInformationUserOperation{}
	}
	return &xxx_SetInformationUserOperation{
		UserHandle:           o.UserHandle,
		UserInformationClass: o.UserInformationClass,
		Buffer:               o.Buffer,
	}
}

func (o *SetInformationUserRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInformationUserOperation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
	o.UserInformationClass = op.UserInformationClass
	o.Buffer = op.Buffer
}
func (o *SetInformationUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetInformationUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInformationUserResponse structure represents the SamrSetInformationUser operation response
type SetInformationUserResponse struct {
	// Return: The SamrSetInformationUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInformationUserResponse) xxx_ToOp(ctx context.Context) *xxx_SetInformationUserOperation {
	if o == nil {
		return &xxx_SetInformationUserOperation{}
	}
	return &xxx_SetInformationUserOperation{
		Return: o.Return,
	}
}

func (o *SetInformationUserResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInformationUserOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetInformationUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetInformationUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ChangePasswordUserOperation structure represents the SamrChangePasswordUser operation
type xxx_ChangePasswordUserOperation struct {
	UserHandle               *Handle                 `idl:"name:UserHandle" json:"user_handle"`
	LMPresent                uint8                   `idl:"name:LmPresent" json:"lm_present"`
	OldLMEncryptedWithNewLM  *EncryptedLMOWFPassword `idl:"name:OldLmEncryptedWithNewLm;pointer:unique" json:"old_lm_encrypted_with_new_lm"`
	NewLMEncryptedWithOldLM  *EncryptedLMOWFPassword `idl:"name:NewLmEncryptedWithOldLm;pointer:unique" json:"new_lm_encrypted_with_old_lm"`
	NTPresent                uint8                   `idl:"name:NtPresent" json:"nt_present"`
	OldNTEncryptedWithNewNT  *EncryptedNTOWFPassword `idl:"name:OldNtEncryptedWithNewNt;pointer:unique" json:"old_nt_encrypted_with_new_nt"`
	NewNTEncryptedWithOldNT  *EncryptedNTOWFPassword `idl:"name:NewNtEncryptedWithOldNt;pointer:unique" json:"new_nt_encrypted_with_old_nt"`
	NTCrossEncryptionPresent uint8                   `idl:"name:NtCrossEncryptionPresent" json:"nt_cross_encryption_present"`
	NewNTEncryptedWithNewLM  *EncryptedNTOWFPassword `idl:"name:NewNtEncryptedWithNewLm;pointer:unique" json:"new_nt_encrypted_with_new_lm"`
	LMCrossEncryptionPresent uint8                   `idl:"name:LmCrossEncryptionPresent" json:"lm_cross_encryption_present"`
	NewLMEncryptedWithNewNT  *EncryptedLMOWFPassword `idl:"name:NewLmEncryptedWithNewNt;pointer:unique" json:"new_lm_encrypted_with_new_nt"`
	Return                   int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_ChangePasswordUserOperation) OpNum() int { return 38 }

func (o *xxx_ChangePasswordUserOperation) OpName() string { return "/samr/v1/SamrChangePasswordUser" }

func (o *xxx_ChangePasswordUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePasswordUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// LmPresent {in} (1:(uchar))
	{
		if err := w.WriteData(o.LMPresent); err != nil {
			return err
		}
	}
	// OldLmEncryptedWithNewLm {in} (1:{pointer=unique, alias=PENCRYPTED_LM_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_LM_OWF_PASSWORD}(struct))
	{
		if o.OldLMEncryptedWithNewLM != nil {
			_ptr_OldLmEncryptedWithNewLm := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldLMEncryptedWithNewLM != nil {
					if err := o.OldLMEncryptedWithNewLM.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedLMOWFPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldLMEncryptedWithNewLM, _ptr_OldLmEncryptedWithNewLm); err != nil {
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
	// NewLmEncryptedWithOldLm {in} (1:{pointer=unique, alias=PENCRYPTED_LM_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_LM_OWF_PASSWORD}(struct))
	{
		if o.NewLMEncryptedWithOldLM != nil {
			_ptr_NewLmEncryptedWithOldLm := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewLMEncryptedWithOldLM != nil {
					if err := o.NewLMEncryptedWithOldLM.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedLMOWFPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewLMEncryptedWithOldLM, _ptr_NewLmEncryptedWithOldLm); err != nil {
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
	// NtPresent {in} (1:(uchar))
	{
		if err := w.WriteData(o.NTPresent); err != nil {
			return err
		}
	}
	// OldNtEncryptedWithNewNt {in} (1:{pointer=unique, alias=PENCRYPTED_NT_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_NT_OWF_PASSWORD}(struct))
	{
		if o.OldNTEncryptedWithNewNT != nil {
			_ptr_OldNtEncryptedWithNewNt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldNTEncryptedWithNewNT != nil {
					if err := o.OldNTEncryptedWithNewNT.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedNTOWFPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldNTEncryptedWithNewNT, _ptr_OldNtEncryptedWithNewNt); err != nil {
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
	// NewNtEncryptedWithOldNt {in} (1:{pointer=unique, alias=PENCRYPTED_NT_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_NT_OWF_PASSWORD}(struct))
	{
		if o.NewNTEncryptedWithOldNT != nil {
			_ptr_NewNtEncryptedWithOldNt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewNTEncryptedWithOldNT != nil {
					if err := o.NewNTEncryptedWithOldNT.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedNTOWFPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewNTEncryptedWithOldNT, _ptr_NewNtEncryptedWithOldNt); err != nil {
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
	// NtCrossEncryptionPresent {in} (1:(uchar))
	{
		if err := w.WriteData(o.NTCrossEncryptionPresent); err != nil {
			return err
		}
	}
	// NewNtEncryptedWithNewLm {in} (1:{pointer=unique, alias=PENCRYPTED_NT_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_NT_OWF_PASSWORD}(struct))
	{
		if o.NewNTEncryptedWithNewLM != nil {
			_ptr_NewNtEncryptedWithNewLm := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewNTEncryptedWithNewLM != nil {
					if err := o.NewNTEncryptedWithNewLM.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedNTOWFPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewNTEncryptedWithNewLM, _ptr_NewNtEncryptedWithNewLm); err != nil {
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
	// LmCrossEncryptionPresent {in} (1:(uchar))
	{
		if err := w.WriteData(o.LMCrossEncryptionPresent); err != nil {
			return err
		}
	}
	// NewLmEncryptedWithNewNt {in} (1:{pointer=unique, alias=PENCRYPTED_LM_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_LM_OWF_PASSWORD}(struct))
	{
		if o.NewLMEncryptedWithNewNT != nil {
			_ptr_NewLmEncryptedWithNewNt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewLMEncryptedWithNewNT != nil {
					if err := o.NewLMEncryptedWithNewNT.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedLMOWFPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewLMEncryptedWithNewNT, _ptr_NewLmEncryptedWithNewNt); err != nil {
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

func (o *xxx_ChangePasswordUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// LmPresent {in} (1:(uchar))
	{
		if err := w.ReadData(&o.LMPresent); err != nil {
			return err
		}
	}
	// OldLmEncryptedWithNewLm {in} (1:{pointer=unique, alias=PENCRYPTED_LM_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_LM_OWF_PASSWORD}(struct))
	{
		_ptr_OldLmEncryptedWithNewLm := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldLMEncryptedWithNewLM == nil {
				o.OldLMEncryptedWithNewLM = &EncryptedLMOWFPassword{}
			}
			if err := o.OldLMEncryptedWithNewLM.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OldLmEncryptedWithNewLm := func(ptr interface{}) { o.OldLMEncryptedWithNewLM = *ptr.(**EncryptedLMOWFPassword) }
		if err := w.ReadPointer(&o.OldLMEncryptedWithNewLM, _s_OldLmEncryptedWithNewLm, _ptr_OldLmEncryptedWithNewLm); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NewLmEncryptedWithOldLm {in} (1:{pointer=unique, alias=PENCRYPTED_LM_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_LM_OWF_PASSWORD}(struct))
	{
		_ptr_NewLmEncryptedWithOldLm := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewLMEncryptedWithOldLM == nil {
				o.NewLMEncryptedWithOldLM = &EncryptedLMOWFPassword{}
			}
			if err := o.NewLMEncryptedWithOldLM.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_NewLmEncryptedWithOldLm := func(ptr interface{}) { o.NewLMEncryptedWithOldLM = *ptr.(**EncryptedLMOWFPassword) }
		if err := w.ReadPointer(&o.NewLMEncryptedWithOldLM, _s_NewLmEncryptedWithOldLm, _ptr_NewLmEncryptedWithOldLm); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NtPresent {in} (1:(uchar))
	{
		if err := w.ReadData(&o.NTPresent); err != nil {
			return err
		}
	}
	// OldNtEncryptedWithNewNt {in} (1:{pointer=unique, alias=PENCRYPTED_NT_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_NT_OWF_PASSWORD}(struct))
	{
		_ptr_OldNtEncryptedWithNewNt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldNTEncryptedWithNewNT == nil {
				o.OldNTEncryptedWithNewNT = &EncryptedNTOWFPassword{}
			}
			if err := o.OldNTEncryptedWithNewNT.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OldNtEncryptedWithNewNt := func(ptr interface{}) { o.OldNTEncryptedWithNewNT = *ptr.(**EncryptedNTOWFPassword) }
		if err := w.ReadPointer(&o.OldNTEncryptedWithNewNT, _s_OldNtEncryptedWithNewNt, _ptr_OldNtEncryptedWithNewNt); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NewNtEncryptedWithOldNt {in} (1:{pointer=unique, alias=PENCRYPTED_NT_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_NT_OWF_PASSWORD}(struct))
	{
		_ptr_NewNtEncryptedWithOldNt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewNTEncryptedWithOldNT == nil {
				o.NewNTEncryptedWithOldNT = &EncryptedNTOWFPassword{}
			}
			if err := o.NewNTEncryptedWithOldNT.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_NewNtEncryptedWithOldNt := func(ptr interface{}) { o.NewNTEncryptedWithOldNT = *ptr.(**EncryptedNTOWFPassword) }
		if err := w.ReadPointer(&o.NewNTEncryptedWithOldNT, _s_NewNtEncryptedWithOldNt, _ptr_NewNtEncryptedWithOldNt); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NtCrossEncryptionPresent {in} (1:(uchar))
	{
		if err := w.ReadData(&o.NTCrossEncryptionPresent); err != nil {
			return err
		}
	}
	// NewNtEncryptedWithNewLm {in} (1:{pointer=unique, alias=PENCRYPTED_NT_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_NT_OWF_PASSWORD}(struct))
	{
		_ptr_NewNtEncryptedWithNewLm := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewNTEncryptedWithNewLM == nil {
				o.NewNTEncryptedWithNewLM = &EncryptedNTOWFPassword{}
			}
			if err := o.NewNTEncryptedWithNewLM.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_NewNtEncryptedWithNewLm := func(ptr interface{}) { o.NewNTEncryptedWithNewLM = *ptr.(**EncryptedNTOWFPassword) }
		if err := w.ReadPointer(&o.NewNTEncryptedWithNewLM, _s_NewNtEncryptedWithNewLm, _ptr_NewNtEncryptedWithNewLm); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// LmCrossEncryptionPresent {in} (1:(uchar))
	{
		if err := w.ReadData(&o.LMCrossEncryptionPresent); err != nil {
			return err
		}
	}
	// NewLmEncryptedWithNewNt {in} (1:{pointer=unique, alias=PENCRYPTED_LM_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_LM_OWF_PASSWORD}(struct))
	{
		_ptr_NewLmEncryptedWithNewNt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewLMEncryptedWithNewNT == nil {
				o.NewLMEncryptedWithNewNT = &EncryptedLMOWFPassword{}
			}
			if err := o.NewLMEncryptedWithNewNT.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_NewLmEncryptedWithNewNt := func(ptr interface{}) { o.NewLMEncryptedWithNewNT = *ptr.(**EncryptedLMOWFPassword) }
		if err := w.ReadPointer(&o.NewLMEncryptedWithNewNT, _s_NewLmEncryptedWithNewNt, _ptr_NewLmEncryptedWithNewNt); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePasswordUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePasswordUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangePasswordUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ChangePasswordUserRequest structure represents the SamrChangePasswordUser operation request
type ChangePasswordUserRequest struct {
	// UserHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a user object.
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
	// LmPresent: If this parameter is zero, the OldLmEncryptedWithNewLm and NewLmEncryptedWithOldLm
	// fields MUST be ignored by the server; otherwise these fields MUST be processed.
	LMPresent uint8 `idl:"name:LmPresent" json:"lm_present"`
	// OldLmEncryptedWithNewLm: The LM hash of the target user's existing password (as presented
	// by the client) encrypted according to the specification of ENCRYPTED_LM_OWF_PASSWORD
	// (section 2.2.7.3), where the key is the LM hash of the new password for the target
	// user (as presented by the client in the NewLmEncryptedWithOldLm parameter).
	OldLMEncryptedWithNewLM *EncryptedLMOWFPassword `idl:"name:OldLmEncryptedWithNewLm;pointer:unique" json:"old_lm_encrypted_with_new_lm"`
	// NewLmEncryptedWithOldLm: The LM hash of the target user's new password (as presented
	// by the client) encrypted according to the specification of ENCRYPTED_LM_OWF_PASSWORD,
	// where the key is the LM hash of the existing password for the target user (as presented
	// by the client in the OldLmEncryptedWithNewLm parameter).
	NewLMEncryptedWithOldLM *EncryptedLMOWFPassword `idl:"name:NewLmEncryptedWithOldLm;pointer:unique" json:"new_lm_encrypted_with_old_lm"`
	// NtPresent: If this parameter is zero, OldNtEncryptedWithNewNt and NewNtEncryptedWithOldNt
	// MUST be ignored by the server; otherwise these fields MUST be processed.
	NTPresent uint8 `idl:"name:NtPresent" json:"nt_present"`
	// OldNtEncryptedWithNewNt: The NT hash of the target user's existing password (as presented
	// by the client) encrypted according to the specification of ENCRYPTED_NT_OWF_PASSWORD
	// (section 2.2.7.3), where the key is the NT hash of the new password for the target
	// user (as presented by the client).
	OldNTEncryptedWithNewNT *EncryptedNTOWFPassword `idl:"name:OldNtEncryptedWithNewNt;pointer:unique" json:"old_nt_encrypted_with_new_nt"`
	// NewNtEncryptedWithOldNt: The NT hash of the target user's new password (as presented
	// by the client) encrypted according to the specification of ENCRYPTED_NT_OWF_PASSWORD,
	// where the key is the NT hash of the existing password for the target user (as presented
	// by the client).
	NewNTEncryptedWithOldNT *EncryptedNTOWFPassword `idl:"name:NewNtEncryptedWithOldNt;pointer:unique" json:"new_nt_encrypted_with_old_nt"`
	// NtCrossEncryptionPresent: If this parameter is zero, NewNtEncryptedWithNewLm MUST
	// be ignored; otherwise, this field MUST be processed.
	NTCrossEncryptionPresent uint8 `idl:"name:NtCrossEncryptionPresent" json:"nt_cross_encryption_present"`
	// NewNtEncryptedWithNewLm: The NT hash of the target user's new password (as presented
	// by the client) encrypted according to the specification of ENCRYPTED_NT_OWF_PASSWORD,
	// where the key is the LM hash of the new password for the target user (as presented
	// by the client).
	NewNTEncryptedWithNewLM *EncryptedNTOWFPassword `idl:"name:NewNtEncryptedWithNewLm;pointer:unique" json:"new_nt_encrypted_with_new_lm"`
	// LmCrossEncryptionPresent: If this parameter is zero, NewLmEncryptedWithNewNt MUST
	// be ignored; otherwise, this field MUST be processed.
	LMCrossEncryptionPresent uint8 `idl:"name:LmCrossEncryptionPresent" json:"lm_cross_encryption_present"`
	// NewLmEncryptedWithNewNt: The LM hash of the target user's new password (as presented
	// by the client) encrypted according to the specification of ENCRYPTED_LM_OWF_PASSWORD,
	// where the key is the NT hash of the new password for the target user (as presented
	// by the client).
	NewLMEncryptedWithNewNT *EncryptedLMOWFPassword `idl:"name:NewLmEncryptedWithNewNt;pointer:unique" json:"new_lm_encrypted_with_new_nt"`
}

func (o *ChangePasswordUserRequest) xxx_ToOp(ctx context.Context) *xxx_ChangePasswordUserOperation {
	if o == nil {
		return &xxx_ChangePasswordUserOperation{}
	}
	return &xxx_ChangePasswordUserOperation{
		UserHandle:               o.UserHandle,
		LMPresent:                o.LMPresent,
		OldLMEncryptedWithNewLM:  o.OldLMEncryptedWithNewLM,
		NewLMEncryptedWithOldLM:  o.NewLMEncryptedWithOldLM,
		NTPresent:                o.NTPresent,
		OldNTEncryptedWithNewNT:  o.OldNTEncryptedWithNewNT,
		NewNTEncryptedWithOldNT:  o.NewNTEncryptedWithOldNT,
		NTCrossEncryptionPresent: o.NTCrossEncryptionPresent,
		NewNTEncryptedWithNewLM:  o.NewNTEncryptedWithNewLM,
		LMCrossEncryptionPresent: o.LMCrossEncryptionPresent,
		NewLMEncryptedWithNewNT:  o.NewLMEncryptedWithNewNT,
	}
}

func (o *ChangePasswordUserRequest) xxx_FromOp(ctx context.Context, op *xxx_ChangePasswordUserOperation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
	o.LMPresent = op.LMPresent
	o.OldLMEncryptedWithNewLM = op.OldLMEncryptedWithNewLM
	o.NewLMEncryptedWithOldLM = op.NewLMEncryptedWithOldLM
	o.NTPresent = op.NTPresent
	o.OldNTEncryptedWithNewNT = op.OldNTEncryptedWithNewNT
	o.NewNTEncryptedWithOldNT = op.NewNTEncryptedWithOldNT
	o.NTCrossEncryptionPresent = op.NTCrossEncryptionPresent
	o.NewNTEncryptedWithNewLM = op.NewNTEncryptedWithNewLM
	o.LMCrossEncryptionPresent = op.LMCrossEncryptionPresent
	o.NewLMEncryptedWithNewNT = op.NewLMEncryptedWithNewNT
}
func (o *ChangePasswordUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ChangePasswordUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangePasswordUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ChangePasswordUserResponse structure represents the SamrChangePasswordUser operation response
type ChangePasswordUserResponse struct {
	// Return: The SamrChangePasswordUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ChangePasswordUserResponse) xxx_ToOp(ctx context.Context) *xxx_ChangePasswordUserOperation {
	if o == nil {
		return &xxx_ChangePasswordUserOperation{}
	}
	return &xxx_ChangePasswordUserOperation{
		Return: o.Return,
	}
}

func (o *ChangePasswordUserResponse) xxx_FromOp(ctx context.Context, op *xxx_ChangePasswordUserOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ChangePasswordUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ChangePasswordUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangePasswordUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetGroupsForUserOperation structure represents the SamrGetGroupsForUser operation
type xxx_GetGroupsForUserOperation struct {
	UserHandle *Handle          `idl:"name:UserHandle" json:"user_handle"`
	Groups     *GetGroupsBuffer `idl:"name:Groups" json:"groups"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetGroupsForUserOperation) OpNum() int { return 39 }

func (o *xxx_GetGroupsForUserOperation) OpName() string { return "/samr/v1/SamrGetGroupsForUser" }

func (o *xxx_GetGroupsForUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupsForUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetGroupsForUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupsForUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupsForUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Groups {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_GET_GROUPS_BUFFER}*(1))(3:{alias=SAMPR_GET_GROUPS_BUFFER}(struct))
	{
		if o.Groups != nil {
			_ptr_Groups := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Groups != nil {
					if err := o.Groups.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GetGroupsBuffer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Groups, _ptr_Groups); err != nil {
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
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupsForUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Groups {out} (1:{pointer=ref}*(2))(2:{alias=PSAMPR_GET_GROUPS_BUFFER,pointer=ref}*(1))(3:{alias=SAMPR_GET_GROUPS_BUFFER}(struct))
	{
		_ptr_Groups := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Groups == nil {
				o.Groups = &GetGroupsBuffer{}
			}
			if err := o.Groups.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Groups := func(ptr interface{}) { o.Groups = *ptr.(**GetGroupsBuffer) }
		if err := w.ReadPointer(&o.Groups, _s_Groups, _ptr_Groups); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetGroupsForUserRequest structure represents the SamrGetGroupsForUser operation request
type GetGroupsForUserRequest struct {
	// UserHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a user object.
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
}

func (o *GetGroupsForUserRequest) xxx_ToOp(ctx context.Context) *xxx_GetGroupsForUserOperation {
	if o == nil {
		return &xxx_GetGroupsForUserOperation{}
	}
	return &xxx_GetGroupsForUserOperation{
		UserHandle: o.UserHandle,
	}
}

func (o *GetGroupsForUserRequest) xxx_FromOp(ctx context.Context, op *xxx_GetGroupsForUserOperation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
}
func (o *GetGroupsForUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetGroupsForUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGroupsForUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetGroupsForUserResponse structure represents the SamrGetGroupsForUser operation response
type GetGroupsForUserResponse struct {
	// Groups: An array of RIDs of the groups that the user referenced by UserHandle is
	// a member of.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Groups *GetGroupsBuffer `idl:"name:Groups" json:"groups"`
	// Return: The SamrGetGroupsForUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetGroupsForUserResponse) xxx_ToOp(ctx context.Context) *xxx_GetGroupsForUserOperation {
	if o == nil {
		return &xxx_GetGroupsForUserOperation{}
	}
	return &xxx_GetGroupsForUserOperation{
		Groups: o.Groups,
		Return: o.Return,
	}
}

func (o *GetGroupsForUserResponse) xxx_FromOp(ctx context.Context, op *xxx_GetGroupsForUserOperation) {
	if o == nil {
		return
	}
	o.Groups = op.Groups
	o.Return = op.Return
}
func (o *GetGroupsForUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetGroupsForUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGroupsForUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryDisplayInformationOperation structure represents the SamrQueryDisplayInformation operation
type xxx_QueryDisplayInformationOperation struct {
	Domain                  *Handle                  `idl:"name:DomainHandle" json:"domain"`
	DisplayInformationClass DomainDisplayInformation `idl:"name:DisplayInformationClass" json:"display_information_class"`
	Index                   uint32                   `idl:"name:Index" json:"index"`
	EntryCount              uint32                   `idl:"name:EntryCount" json:"entry_count"`
	PreferredMaximumLength  uint32                   `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
	TotalAvailable          uint32                   `idl:"name:TotalAvailable" json:"total_available"`
	TotalReturned           uint32                   `idl:"name:TotalReturned" json:"total_returned"`
	Buffer                  *DisplayInfoBuffer       `idl:"name:Buffer;switch_is:DisplayInformationClass" json:"buffer"`
	Return                  int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryDisplayInformationOperation) OpNum() int { return 40 }

func (o *xxx_QueryDisplayInformationOperation) OpName() string {
	return "/samr/v1/SamrQueryDisplayInformation"
}

func (o *xxx_QueryDisplayInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DisplayInformationClass {in} (1:{alias=DOMAIN_DISPLAY_INFORMATION}(enum))
	{
		if err := w.WriteEnum(uint16(o.DisplayInformationClass)); err != nil {
			return err
		}
	}
	// Index {in} (1:(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// EntryCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.EntryCount); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DisplayInformationClass {in} (1:{alias=DOMAIN_DISPLAY_INFORMATION}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.DisplayInformationClass)); err != nil {
			return err
		}
	}
	// Index {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// EntryCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.EntryCount); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TotalAvailable {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TotalAvailable); err != nil {
			return err
		}
	}
	// TotalReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TotalReturned); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=PSAMPR_DISPLAY_INFO_BUFFER}*(1))(2:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=SAMPR_DISPLAY_INFO_BUFFER}(union))
	{
		_swBuffer := uint16(o.DisplayInformationClass)
		if o.Buffer != nil {
			if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		} else {
			if err := (&DisplayInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TotalAvailable {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TotalAvailable); err != nil {
			return err
		}
	}
	// TotalReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TotalReturned); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=PSAMPR_DISPLAY_INFO_BUFFER,pointer=ref}*(1))(2:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=SAMPR_DISPLAY_INFO_BUFFER}(union))
	{
		if o.Buffer == nil {
			o.Buffer = &DisplayInfoBuffer{}
		}
		_swBuffer := uint16(o.DisplayInformationClass)
		if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryDisplayInformationRequest structure represents the SamrQueryDisplayInformation operation request
type QueryDisplayInformationRequest struct {
	Domain                  *Handle                  `idl:"name:DomainHandle" json:"domain"`
	DisplayInformationClass DomainDisplayInformation `idl:"name:DisplayInformationClass" json:"display_information_class"`
	Index                   uint32                   `idl:"name:Index" json:"index"`
	EntryCount              uint32                   `idl:"name:EntryCount" json:"entry_count"`
	PreferredMaximumLength  uint32                   `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
}

func (o *QueryDisplayInformationRequest) xxx_ToOp(ctx context.Context) *xxx_QueryDisplayInformationOperation {
	if o == nil {
		return &xxx_QueryDisplayInformationOperation{}
	}
	return &xxx_QueryDisplayInformationOperation{
		Domain:                  o.Domain,
		DisplayInformationClass: o.DisplayInformationClass,
		Index:                   o.Index,
		EntryCount:              o.EntryCount,
		PreferredMaximumLength:  o.PreferredMaximumLength,
	}
}

func (o *QueryDisplayInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryDisplayInformationOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DisplayInformationClass = op.DisplayInformationClass
	o.Index = op.Index
	o.EntryCount = op.EntryCount
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *QueryDisplayInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryDisplayInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDisplayInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryDisplayInformationResponse structure represents the SamrQueryDisplayInformation operation response
type QueryDisplayInformationResponse struct {
	TotalAvailable uint32             `idl:"name:TotalAvailable" json:"total_available"`
	TotalReturned  uint32             `idl:"name:TotalReturned" json:"total_returned"`
	Buffer         *DisplayInfoBuffer `idl:"name:Buffer;switch_is:DisplayInformationClass" json:"buffer"`
	// Return: The SamrQueryDisplayInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryDisplayInformationResponse) xxx_ToOp(ctx context.Context) *xxx_QueryDisplayInformationOperation {
	if o == nil {
		return &xxx_QueryDisplayInformationOperation{}
	}
	return &xxx_QueryDisplayInformationOperation{
		TotalAvailable: o.TotalAvailable,
		TotalReturned:  o.TotalReturned,
		Buffer:         o.Buffer,
		Return:         o.Return,
	}
}

func (o *QueryDisplayInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryDisplayInformationOperation) {
	if o == nil {
		return
	}
	o.TotalAvailable = op.TotalAvailable
	o.TotalReturned = op.TotalReturned
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *QueryDisplayInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryDisplayInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDisplayInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDisplayEnumerationIndexOperation structure represents the SamrGetDisplayEnumerationIndex operation
type xxx_GetDisplayEnumerationIndexOperation struct {
	Domain                  *Handle                  `idl:"name:DomainHandle" json:"domain"`
	DisplayInformationClass DomainDisplayInformation `idl:"name:DisplayInformationClass" json:"display_information_class"`
	Prefix                  *dtyp.UnicodeString      `idl:"name:Prefix" json:"prefix"`
	Index                   uint32                   `idl:"name:Index" json:"index"`
	Return                  int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDisplayEnumerationIndexOperation) OpNum() int { return 41 }

func (o *xxx_GetDisplayEnumerationIndexOperation) OpName() string {
	return "/samr/v1/SamrGetDisplayEnumerationIndex"
}

func (o *xxx_GetDisplayEnumerationIndexOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayEnumerationIndexOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DisplayInformationClass {in} (1:{alias=DOMAIN_DISPLAY_INFORMATION}(enum))
	{
		if err := w.WriteEnum(uint16(o.DisplayInformationClass)); err != nil {
			return err
		}
	}
	// Prefix {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.Prefix != nil {
			if err := o.Prefix.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetDisplayEnumerationIndexOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DisplayInformationClass {in} (1:{alias=DOMAIN_DISPLAY_INFORMATION}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.DisplayInformationClass)); err != nil {
			return err
		}
	}
	// Prefix {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.Prefix == nil {
			o.Prefix = &dtyp.UnicodeString{}
		}
		if err := o.Prefix.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayEnumerationIndexOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayEnumerationIndexOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Index {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayEnumerationIndexOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Index {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDisplayEnumerationIndexRequest structure represents the SamrGetDisplayEnumerationIndex operation request
type GetDisplayEnumerationIndexRequest struct {
	Domain                  *Handle                  `idl:"name:DomainHandle" json:"domain"`
	DisplayInformationClass DomainDisplayInformation `idl:"name:DisplayInformationClass" json:"display_information_class"`
	Prefix                  *dtyp.UnicodeString      `idl:"name:Prefix" json:"prefix"`
}

func (o *GetDisplayEnumerationIndexRequest) xxx_ToOp(ctx context.Context) *xxx_GetDisplayEnumerationIndexOperation {
	if o == nil {
		return &xxx_GetDisplayEnumerationIndexOperation{}
	}
	return &xxx_GetDisplayEnumerationIndexOperation{
		Domain:                  o.Domain,
		DisplayInformationClass: o.DisplayInformationClass,
		Prefix:                  o.Prefix,
	}
}

func (o *GetDisplayEnumerationIndexRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayEnumerationIndexOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DisplayInformationClass = op.DisplayInformationClass
	o.Prefix = op.Prefix
}
func (o *GetDisplayEnumerationIndexRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDisplayEnumerationIndexRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayEnumerationIndexOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDisplayEnumerationIndexResponse structure represents the SamrGetDisplayEnumerationIndex operation response
type GetDisplayEnumerationIndexResponse struct {
	Index uint32 `idl:"name:Index" json:"index"`
	// Return: The SamrGetDisplayEnumerationIndex return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDisplayEnumerationIndexResponse) xxx_ToOp(ctx context.Context) *xxx_GetDisplayEnumerationIndexOperation {
	if o == nil {
		return &xxx_GetDisplayEnumerationIndexOperation{}
	}
	return &xxx_GetDisplayEnumerationIndexOperation{
		Index:  o.Index,
		Return: o.Return,
	}
}

func (o *GetDisplayEnumerationIndexResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayEnumerationIndexOperation) {
	if o == nil {
		return
	}
	o.Index = op.Index
	o.Return = op.Return
}
func (o *GetDisplayEnumerationIndexResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDisplayEnumerationIndexResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayEnumerationIndexOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUserDomainPasswordInformationOperation structure represents the SamrGetUserDomainPasswordInformation operation
type xxx_GetUserDomainPasswordInformationOperation struct {
	UserHandle          *Handle                        `idl:"name:UserHandle" json:"user_handle"`
	PasswordInformation *UserDomainPasswordInformation `idl:"name:PasswordInformation" json:"password_information"`
	Return              int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUserDomainPasswordInformationOperation) OpNum() int { return 44 }

func (o *xxx_GetUserDomainPasswordInformationOperation) OpName() string {
	return "/samr/v1/SamrGetUserDomainPasswordInformation"
}

func (o *xxx_GetUserDomainPasswordInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserDomainPasswordInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetUserDomainPasswordInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserDomainPasswordInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserDomainPasswordInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PasswordInformation {out} (1:{alias=PUSER_DOMAIN_PASSWORD_INFORMATION}*(1))(2:{alias=USER_DOMAIN_PASSWORD_INFORMATION}(struct))
	{
		if o.PasswordInformation != nil {
			if err := o.PasswordInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserDomainPasswordInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserDomainPasswordInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PasswordInformation {out} (1:{alias=PUSER_DOMAIN_PASSWORD_INFORMATION,pointer=ref}*(1))(2:{alias=USER_DOMAIN_PASSWORD_INFORMATION}(struct))
	{
		if o.PasswordInformation == nil {
			o.PasswordInformation = &UserDomainPasswordInformation{}
		}
		if err := o.PasswordInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetUserDomainPasswordInformationRequest structure represents the SamrGetUserDomainPasswordInformation operation request
type GetUserDomainPasswordInformationRequest struct {
	// UserHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a user object.
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
}

func (o *GetUserDomainPasswordInformationRequest) xxx_ToOp(ctx context.Context) *xxx_GetUserDomainPasswordInformationOperation {
	if o == nil {
		return &xxx_GetUserDomainPasswordInformationOperation{}
	}
	return &xxx_GetUserDomainPasswordInformationOperation{
		UserHandle: o.UserHandle,
	}
}

func (o *GetUserDomainPasswordInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUserDomainPasswordInformationOperation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
}
func (o *GetUserDomainPasswordInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetUserDomainPasswordInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserDomainPasswordInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUserDomainPasswordInformationResponse structure represents the SamrGetUserDomainPasswordInformation operation response
type GetUserDomainPasswordInformationResponse struct {
	// PasswordInformation: Password policy information from the user's domain.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	PasswordInformation *UserDomainPasswordInformation `idl:"name:PasswordInformation" json:"password_information"`
	// Return: The SamrGetUserDomainPasswordInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUserDomainPasswordInformationResponse) xxx_ToOp(ctx context.Context) *xxx_GetUserDomainPasswordInformationOperation {
	if o == nil {
		return &xxx_GetUserDomainPasswordInformationOperation{}
	}
	return &xxx_GetUserDomainPasswordInformationOperation{
		PasswordInformation: o.PasswordInformation,
		Return:              o.Return,
	}
}

func (o *GetUserDomainPasswordInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUserDomainPasswordInformationOperation) {
	if o == nil {
		return
	}
	o.PasswordInformation = op.PasswordInformation
	o.Return = op.Return
}
func (o *GetUserDomainPasswordInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetUserDomainPasswordInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserDomainPasswordInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveMemberFromForeignDomainOperation structure represents the SamrRemoveMemberFromForeignDomain operation
type xxx_RemoveMemberFromForeignDomainOperation struct {
	Domain    *Handle   `idl:"name:DomainHandle" json:"domain"`
	MemberSID *dtyp.SID `idl:"name:MemberSid" json:"member_sid"`
	Return    int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveMemberFromForeignDomainOperation) OpNum() int { return 45 }

func (o *xxx_RemoveMemberFromForeignDomainOperation) OpName() string {
	return "/samr/v1/SamrRemoveMemberFromForeignDomain"
}

func (o *xxx_RemoveMemberFromForeignDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromForeignDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MemberSid {in} (1:{alias=PRPC_SID}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.MemberSID != nil {
			if err := o.MemberSID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RemoveMemberFromForeignDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MemberSid {in} (1:{alias=PRPC_SID,pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.MemberSID == nil {
			o.MemberSID = &dtyp.SID{}
		}
		if err := o.MemberSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromForeignDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromForeignDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMemberFromForeignDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveMemberFromForeignDomainRequest structure represents the SamrRemoveMemberFromForeignDomain operation request
type RemoveMemberFromForeignDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// MemberSid: The SID to remove from the membership.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	MemberSID *dtyp.SID `idl:"name:MemberSid" json:"member_sid"`
}

func (o *RemoveMemberFromForeignDomainRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveMemberFromForeignDomainOperation {
	if o == nil {
		return &xxx_RemoveMemberFromForeignDomainOperation{}
	}
	return &xxx_RemoveMemberFromForeignDomainOperation{
		Domain:    o.Domain,
		MemberSID: o.MemberSID,
	}
}

func (o *RemoveMemberFromForeignDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveMemberFromForeignDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.MemberSID = op.MemberSID
}
func (o *RemoveMemberFromForeignDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveMemberFromForeignDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMemberFromForeignDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveMemberFromForeignDomainResponse structure represents the SamrRemoveMemberFromForeignDomain operation response
type RemoveMemberFromForeignDomainResponse struct {
	// Return: The SamrRemoveMemberFromForeignDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveMemberFromForeignDomainResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveMemberFromForeignDomainOperation {
	if o == nil {
		return &xxx_RemoveMemberFromForeignDomainOperation{}
	}
	return &xxx_RemoveMemberFromForeignDomainOperation{
		Return: o.Return,
	}
}

func (o *RemoveMemberFromForeignDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveMemberFromForeignDomainOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveMemberFromForeignDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveMemberFromForeignDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMemberFromForeignDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInformationDomain2Operation structure represents the SamrQueryInformationDomain2 operation
type xxx_QueryInformationDomain2Operation struct {
	Domain                 *Handle                `idl:"name:DomainHandle" json:"domain"`
	DomainInformationClass DomainInformationClass `idl:"name:DomainInformationClass" json:"domain_information_class"`
	Buffer                 *DomainInfoBuffer      `idl:"name:Buffer;switch_is:DomainInformationClass" json:"buffer"`
	Return                 int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInformationDomain2Operation) OpNum() int { return 46 }

func (o *xxx_QueryInformationDomain2Operation) OpName() string {
	return "/samr/v1/SamrQueryInformationDomain2"
}

func (o *xxx_QueryInformationDomain2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationDomain2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DomainInformationClass {in} (1:{alias=DOMAIN_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.DomainInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationDomain2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DomainInformationClass {in} (1:{alias=DOMAIN_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.DomainInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationDomain2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationDomain2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=PSAMPR_DOMAIN_INFO_BUFFER}*(1))(3:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=SAMPR_DOMAIN_INFO_BUFFER}(union))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swBuffer := uint16(o.DomainInformationClass)
				if o.Buffer != nil {
					if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
						return err
					}
				} else {
					if err := (&DomainInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationDomain2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=PSAMPR_DOMAIN_INFO_BUFFER,pointer=ref}*(1))(3:{switch_type={alias=DOMAIN_INFORMATION_CLASS}(enum), alias=SAMPR_DOMAIN_INFO_BUFFER}(union))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &DomainInfoBuffer{}
			}
			_swBuffer := uint16(o.DomainInformationClass)
			if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**DomainInfoBuffer) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryInformationDomain2Request structure represents the SamrQueryInformationDomain2 operation request
type QueryInformationDomain2Request struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// DomainInformationClass: An enumeration indicating which attributes to return. See
	// section 2.2.3.16 for a listing of possible values.
	DomainInformationClass DomainInformationClass `idl:"name:DomainInformationClass" json:"domain_information_class"`
}

func (o *QueryInformationDomain2Request) xxx_ToOp(ctx context.Context) *xxx_QueryInformationDomain2Operation {
	if o == nil {
		return &xxx_QueryInformationDomain2Operation{}
	}
	return &xxx_QueryInformationDomain2Operation{
		Domain:                 o.Domain,
		DomainInformationClass: o.DomainInformationClass,
	}
}

func (o *QueryInformationDomain2Request) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationDomain2Operation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DomainInformationClass = op.DomainInformationClass
}
func (o *QueryInformationDomain2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInformationDomain2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationDomain2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInformationDomain2Response structure represents the SamrQueryInformationDomain2 operation response
type QueryInformationDomain2Response struct {
	// Buffer: The requested attributes on output. See section 2.2.3.17 for structure details.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Buffer *DomainInfoBuffer `idl:"name:Buffer;switch_is:DomainInformationClass" json:"buffer"`
	// Return: The SamrQueryInformationDomain2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryInformationDomain2Response) xxx_ToOp(ctx context.Context) *xxx_QueryInformationDomain2Operation {
	if o == nil {
		return &xxx_QueryInformationDomain2Operation{}
	}
	return &xxx_QueryInformationDomain2Operation{
		Buffer: o.Buffer,
		Return: o.Return,
	}
}

func (o *QueryInformationDomain2Response) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationDomain2Operation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *QueryInformationDomain2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInformationDomain2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationDomain2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInformationUser2Operation structure represents the SamrQueryInformationUser2 operation
type xxx_QueryInformationUser2Operation struct {
	UserHandle           *Handle              `idl:"name:UserHandle" json:"user_handle"`
	UserInformationClass UserInformationClass `idl:"name:UserInformationClass" json:"user_information_class"`
	Buffer               *UserInfoBuffer      `idl:"name:Buffer;switch_is:UserInformationClass" json:"buffer"`
	Return               int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInformationUser2Operation) OpNum() int { return 47 }

func (o *xxx_QueryInformationUser2Operation) OpName() string {
	return "/samr/v1/SamrQueryInformationUser2"
}

func (o *xxx_QueryInformationUser2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationUser2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// UserInformationClass {in} (1:{alias=USER_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.UserInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationUser2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// UserInformationClass {in} (1:{alias=USER_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.UserInformationClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationUser2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationUser2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=PSAMPR_USER_INFO_BUFFER}*(1))(3:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=SAMPR_USER_INFO_BUFFER}(union))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swBuffer := uint16(o.UserInformationClass)
				if o.Buffer != nil {
					if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
						return err
					}
				} else {
					if err := (&UserInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInformationUser2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=PSAMPR_USER_INFO_BUFFER,pointer=ref}*(1))(3:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=SAMPR_USER_INFO_BUFFER}(union))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &UserInfoBuffer{}
			}
			_swBuffer := uint16(o.UserInformationClass)
			if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**UserInfoBuffer) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryInformationUser2Request structure represents the SamrQueryInformationUser2 operation request
type QueryInformationUser2Request struct {
	// UserHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a user object.
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
	// UserInformationClass: An enumeration indicating which attributes to return. See section
	// 2.2.6.28 for a list of possible values.
	UserInformationClass UserInformationClass `idl:"name:UserInformationClass" json:"user_information_class"`
}

func (o *QueryInformationUser2Request) xxx_ToOp(ctx context.Context) *xxx_QueryInformationUser2Operation {
	if o == nil {
		return &xxx_QueryInformationUser2Operation{}
	}
	return &xxx_QueryInformationUser2Operation{
		UserHandle:           o.UserHandle,
		UserInformationClass: o.UserInformationClass,
	}
}

func (o *QueryInformationUser2Request) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationUser2Operation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
	o.UserInformationClass = op.UserInformationClass
}
func (o *QueryInformationUser2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInformationUser2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationUser2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInformationUser2Response structure represents the SamrQueryInformationUser2 operation response
type QueryInformationUser2Response struct {
	// Buffer: The requested attributes on output. See section 2.2.6.29 for structure details.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Buffer *UserInfoBuffer `idl:"name:Buffer;switch_is:UserInformationClass" json:"buffer"`
	// Return: The SamrQueryInformationUser2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryInformationUser2Response) xxx_ToOp(ctx context.Context) *xxx_QueryInformationUser2Operation {
	if o == nil {
		return &xxx_QueryInformationUser2Operation{}
	}
	return &xxx_QueryInformationUser2Operation{
		Buffer: o.Buffer,
		Return: o.Return,
	}
}

func (o *QueryInformationUser2Response) xxx_FromOp(ctx context.Context, op *xxx_QueryInformationUser2Operation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *QueryInformationUser2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInformationUser2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInformationUser2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryDisplayInformation2Operation structure represents the SamrQueryDisplayInformation2 operation
type xxx_QueryDisplayInformation2Operation struct {
	Domain                  *Handle                  `idl:"name:DomainHandle" json:"domain"`
	DisplayInformationClass DomainDisplayInformation `idl:"name:DisplayInformationClass" json:"display_information_class"`
	Index                   uint32                   `idl:"name:Index" json:"index"`
	EntryCount              uint32                   `idl:"name:EntryCount" json:"entry_count"`
	PreferredMaximumLength  uint32                   `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
	TotalAvailable          uint32                   `idl:"name:TotalAvailable" json:"total_available"`
	TotalReturned           uint32                   `idl:"name:TotalReturned" json:"total_returned"`
	Buffer                  *DisplayInfoBuffer       `idl:"name:Buffer;switch_is:DisplayInformationClass" json:"buffer"`
	Return                  int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryDisplayInformation2Operation) OpNum() int { return 48 }

func (o *xxx_QueryDisplayInformation2Operation) OpName() string {
	return "/samr/v1/SamrQueryDisplayInformation2"
}

func (o *xxx_QueryDisplayInformation2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformation2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DisplayInformationClass {in} (1:{alias=DOMAIN_DISPLAY_INFORMATION}(enum))
	{
		if err := w.WriteEnum(uint16(o.DisplayInformationClass)); err != nil {
			return err
		}
	}
	// Index {in} (1:(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// EntryCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.EntryCount); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformation2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DisplayInformationClass {in} (1:{alias=DOMAIN_DISPLAY_INFORMATION}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.DisplayInformationClass)); err != nil {
			return err
		}
	}
	// Index {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// EntryCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.EntryCount); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformation2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformation2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TotalAvailable {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TotalAvailable); err != nil {
			return err
		}
	}
	// TotalReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TotalReturned); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=PSAMPR_DISPLAY_INFO_BUFFER}*(1))(2:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=SAMPR_DISPLAY_INFO_BUFFER}(union))
	{
		_swBuffer := uint16(o.DisplayInformationClass)
		if o.Buffer != nil {
			if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		} else {
			if err := (&DisplayInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformation2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TotalAvailable {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TotalAvailable); err != nil {
			return err
		}
	}
	// TotalReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TotalReturned); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=PSAMPR_DISPLAY_INFO_BUFFER,pointer=ref}*(1))(2:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=SAMPR_DISPLAY_INFO_BUFFER}(union))
	{
		if o.Buffer == nil {
			o.Buffer = &DisplayInfoBuffer{}
		}
		_swBuffer := uint16(o.DisplayInformationClass)
		if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryDisplayInformation2Request structure represents the SamrQueryDisplayInformation2 operation request
type QueryDisplayInformation2Request struct {
	Domain                  *Handle                  `idl:"name:DomainHandle" json:"domain"`
	DisplayInformationClass DomainDisplayInformation `idl:"name:DisplayInformationClass" json:"display_information_class"`
	Index                   uint32                   `idl:"name:Index" json:"index"`
	EntryCount              uint32                   `idl:"name:EntryCount" json:"entry_count"`
	PreferredMaximumLength  uint32                   `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
}

func (o *QueryDisplayInformation2Request) xxx_ToOp(ctx context.Context) *xxx_QueryDisplayInformation2Operation {
	if o == nil {
		return &xxx_QueryDisplayInformation2Operation{}
	}
	return &xxx_QueryDisplayInformation2Operation{
		Domain:                  o.Domain,
		DisplayInformationClass: o.DisplayInformationClass,
		Index:                   o.Index,
		EntryCount:              o.EntryCount,
		PreferredMaximumLength:  o.PreferredMaximumLength,
	}
}

func (o *QueryDisplayInformation2Request) xxx_FromOp(ctx context.Context, op *xxx_QueryDisplayInformation2Operation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DisplayInformationClass = op.DisplayInformationClass
	o.Index = op.Index
	o.EntryCount = op.EntryCount
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *QueryDisplayInformation2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryDisplayInformation2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDisplayInformation2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryDisplayInformation2Response structure represents the SamrQueryDisplayInformation2 operation response
type QueryDisplayInformation2Response struct {
	TotalAvailable uint32             `idl:"name:TotalAvailable" json:"total_available"`
	TotalReturned  uint32             `idl:"name:TotalReturned" json:"total_returned"`
	Buffer         *DisplayInfoBuffer `idl:"name:Buffer;switch_is:DisplayInformationClass" json:"buffer"`
	// Return: The SamrQueryDisplayInformation2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryDisplayInformation2Response) xxx_ToOp(ctx context.Context) *xxx_QueryDisplayInformation2Operation {
	if o == nil {
		return &xxx_QueryDisplayInformation2Operation{}
	}
	return &xxx_QueryDisplayInformation2Operation{
		TotalAvailable: o.TotalAvailable,
		TotalReturned:  o.TotalReturned,
		Buffer:         o.Buffer,
		Return:         o.Return,
	}
}

func (o *QueryDisplayInformation2Response) xxx_FromOp(ctx context.Context, op *xxx_QueryDisplayInformation2Operation) {
	if o == nil {
		return
	}
	o.TotalAvailable = op.TotalAvailable
	o.TotalReturned = op.TotalReturned
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *QueryDisplayInformation2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryDisplayInformation2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDisplayInformation2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDisplayEnumerationIndex2Operation structure represents the SamrGetDisplayEnumerationIndex2 operation
type xxx_GetDisplayEnumerationIndex2Operation struct {
	Domain                  *Handle                  `idl:"name:DomainHandle" json:"domain"`
	DisplayInformationClass DomainDisplayInformation `idl:"name:DisplayInformationClass" json:"display_information_class"`
	Prefix                  *dtyp.UnicodeString      `idl:"name:Prefix" json:"prefix"`
	Index                   uint32                   `idl:"name:Index" json:"index"`
	Return                  int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDisplayEnumerationIndex2Operation) OpNum() int { return 49 }

func (o *xxx_GetDisplayEnumerationIndex2Operation) OpName() string {
	return "/samr/v1/SamrGetDisplayEnumerationIndex2"
}

func (o *xxx_GetDisplayEnumerationIndex2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayEnumerationIndex2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DisplayInformationClass {in} (1:{alias=DOMAIN_DISPLAY_INFORMATION}(enum))
	{
		if err := w.WriteEnum(uint16(o.DisplayInformationClass)); err != nil {
			return err
		}
	}
	// Prefix {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.Prefix != nil {
			if err := o.Prefix.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetDisplayEnumerationIndex2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DisplayInformationClass {in} (1:{alias=DOMAIN_DISPLAY_INFORMATION}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.DisplayInformationClass)); err != nil {
			return err
		}
	}
	// Prefix {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.Prefix == nil {
			o.Prefix = &dtyp.UnicodeString{}
		}
		if err := o.Prefix.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayEnumerationIndex2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayEnumerationIndex2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Index {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDisplayEnumerationIndex2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Index {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDisplayEnumerationIndex2Request structure represents the SamrGetDisplayEnumerationIndex2 operation request
type GetDisplayEnumerationIndex2Request struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// DisplayInformationClass: An enumeration indicating which set of objects to return
	// an index into (for a subsequent SamrQueryDisplayInformation3 method call).
	DisplayInformationClass DomainDisplayInformation `idl:"name:DisplayInformationClass" json:"display_information_class"`
	// Prefix: A string matched against the account name to find a starting point for an
	// enumeration. The Prefix parameter enables the client to obtain a listing of an account
	// from SamrQueryDisplayInformation3 such that the accounts are returned in alphabetical
	// order with respect to their account name, starting with the account name that most
	// closely matches Prefix. See details later in this section.
	Prefix *dtyp.UnicodeString `idl:"name:Prefix" json:"prefix"`
}

func (o *GetDisplayEnumerationIndex2Request) xxx_ToOp(ctx context.Context) *xxx_GetDisplayEnumerationIndex2Operation {
	if o == nil {
		return &xxx_GetDisplayEnumerationIndex2Operation{}
	}
	return &xxx_GetDisplayEnumerationIndex2Operation{
		Domain:                  o.Domain,
		DisplayInformationClass: o.DisplayInformationClass,
		Prefix:                  o.Prefix,
	}
}

func (o *GetDisplayEnumerationIndex2Request) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayEnumerationIndex2Operation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DisplayInformationClass = op.DisplayInformationClass
	o.Prefix = op.Prefix
}
func (o *GetDisplayEnumerationIndex2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDisplayEnumerationIndex2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayEnumerationIndex2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDisplayEnumerationIndex2Response structure represents the SamrGetDisplayEnumerationIndex2 operation response
type GetDisplayEnumerationIndex2Response struct {
	// Index: A value to use as input to SamrQueryDisplayInformation3 in order to control
	// the accounts that are returned from that method.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Index uint32 `idl:"name:Index" json:"index"`
	// Return: The SamrGetDisplayEnumerationIndex2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDisplayEnumerationIndex2Response) xxx_ToOp(ctx context.Context) *xxx_GetDisplayEnumerationIndex2Operation {
	if o == nil {
		return &xxx_GetDisplayEnumerationIndex2Operation{}
	}
	return &xxx_GetDisplayEnumerationIndex2Operation{
		Index:  o.Index,
		Return: o.Return,
	}
}

func (o *GetDisplayEnumerationIndex2Response) xxx_FromOp(ctx context.Context, op *xxx_GetDisplayEnumerationIndex2Operation) {
	if o == nil {
		return
	}
	o.Index = op.Index
	o.Return = op.Return
}
func (o *GetDisplayEnumerationIndex2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDisplayEnumerationIndex2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDisplayEnumerationIndex2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateUser2InDomainOperation structure represents the SamrCreateUser2InDomain operation
type xxx_CreateUser2InDomainOperation struct {
	Domain        *Handle             `idl:"name:DomainHandle" json:"domain"`
	Name          *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	AccountType   uint32              `idl:"name:AccountType" json:"account_type"`
	DesiredAccess uint32              `idl:"name:DesiredAccess" json:"desired_access"`
	UserHandle    *Handle             `idl:"name:UserHandle" json:"user_handle"`
	GrantedAccess uint32              `idl:"name:GrantedAccess" json:"granted_access"`
	RelativeID    uint32              `idl:"name:RelativeId" json:"relative_id"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateUser2InDomainOperation) OpNum() int { return 50 }

func (o *xxx_CreateUser2InDomainOperation) OpName() string { return "/samr/v1/SamrCreateUser2InDomain" }

func (o *xxx_CreateUser2InDomainOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUser2InDomainOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
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
	// AccountType {in} (1:(uint32))
	{
		if err := w.WriteData(o.AccountType); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUser2InDomainOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
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
	// AccountType {in} (1:(uint32))
	{
		if err := w.ReadData(&o.AccountType); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUser2InDomainOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUser2InDomainOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// UserHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// GrantedAccess {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.GrantedAccess); err != nil {
			return err
		}
	}
	// RelativeId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RelativeID); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateUser2InDomainOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// UserHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// GrantedAccess {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.GrantedAccess); err != nil {
			return err
		}
	}
	// RelativeId {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RelativeID); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateUser2InDomainRequest structure represents the SamrCreateUser2InDomain operation request
type CreateUser2InDomainRequest struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// Name: The value to use as the name of the user. See the message processing shown
	// later in this section for details on how this value maps to the data model.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// AccountType: A 32-bit value indicating the type of account to create. See the message
	// processing shown later in this section for possible values.
	AccountType uint32 `idl:"name:AccountType" json:"account_type"`
	// DesiredAccess: The access requested on the UserHandle on output. See section 2.2.1.7
	// for a listing of possible values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *CreateUser2InDomainRequest) xxx_ToOp(ctx context.Context) *xxx_CreateUser2InDomainOperation {
	if o == nil {
		return &xxx_CreateUser2InDomainOperation{}
	}
	return &xxx_CreateUser2InDomainOperation{
		Domain:        o.Domain,
		Name:          o.Name,
		AccountType:   o.AccountType,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *CreateUser2InDomainRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateUser2InDomainOperation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.Name = op.Name
	o.AccountType = op.AccountType
	o.DesiredAccess = op.DesiredAccess
}
func (o *CreateUser2InDomainRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateUser2InDomainRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateUser2InDomainOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateUser2InDomainResponse structure represents the SamrCreateUser2InDomain operation response
type CreateUser2InDomainResponse struct {
	// UserHandle: An RPC context handle, as specified in section 2.2.7.2.
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
	// GrantedAccess: The access granted on UserHandle.
	GrantedAccess uint32 `idl:"name:GrantedAccess" json:"granted_access"`
	// RelativeId: The RID of the newly created user.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	RelativeID uint32 `idl:"name:RelativeId" json:"relative_id"`
	// Return: The SamrCreateUser2InDomain return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateUser2InDomainResponse) xxx_ToOp(ctx context.Context) *xxx_CreateUser2InDomainOperation {
	if o == nil {
		return &xxx_CreateUser2InDomainOperation{}
	}
	return &xxx_CreateUser2InDomainOperation{
		UserHandle:    o.UserHandle,
		GrantedAccess: o.GrantedAccess,
		RelativeID:    o.RelativeID,
		Return:        o.Return,
	}
}

func (o *CreateUser2InDomainResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateUser2InDomainOperation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
	o.GrantedAccess = op.GrantedAccess
	o.RelativeID = op.RelativeID
	o.Return = op.Return
}
func (o *CreateUser2InDomainResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateUser2InDomainResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateUser2InDomainOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryDisplayInformation3Operation structure represents the SamrQueryDisplayInformation3 operation
type xxx_QueryDisplayInformation3Operation struct {
	Domain                  *Handle                  `idl:"name:DomainHandle" json:"domain"`
	DisplayInformationClass DomainDisplayInformation `idl:"name:DisplayInformationClass" json:"display_information_class"`
	Index                   uint32                   `idl:"name:Index" json:"index"`
	EntryCount              uint32                   `idl:"name:EntryCount" json:"entry_count"`
	PreferredMaximumLength  uint32                   `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
	TotalAvailable          uint32                   `idl:"name:TotalAvailable" json:"total_available"`
	TotalReturned           uint32                   `idl:"name:TotalReturned" json:"total_returned"`
	Buffer                  *DisplayInfoBuffer       `idl:"name:Buffer;switch_is:DisplayInformationClass" json:"buffer"`
	Return                  int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryDisplayInformation3Operation) OpNum() int { return 51 }

func (o *xxx_QueryDisplayInformation3Operation) OpName() string {
	return "/samr/v1/SamrQueryDisplayInformation3"
}

func (o *xxx_QueryDisplayInformation3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformation3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain != nil {
			if err := o.Domain.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DisplayInformationClass {in} (1:{alias=DOMAIN_DISPLAY_INFORMATION}(enum))
	{
		if err := w.WriteEnum(uint16(o.DisplayInformationClass)); err != nil {
			return err
		}
	}
	// Index {in} (1:(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// EntryCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.EntryCount); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformation3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DomainHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Domain == nil {
			o.Domain = &Handle{}
		}
		if err := o.Domain.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DisplayInformationClass {in} (1:{alias=DOMAIN_DISPLAY_INFORMATION}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.DisplayInformationClass)); err != nil {
			return err
		}
	}
	// Index {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// EntryCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.EntryCount); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformation3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformation3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TotalAvailable {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TotalAvailable); err != nil {
			return err
		}
	}
	// TotalReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TotalReturned); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=PSAMPR_DISPLAY_INFO_BUFFER}*(1))(2:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=SAMPR_DISPLAY_INFO_BUFFER}(union))
	{
		_swBuffer := uint16(o.DisplayInformationClass)
		if o.Buffer != nil {
			if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		} else {
			if err := (&DisplayInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisplayInformation3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TotalAvailable {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TotalAvailable); err != nil {
			return err
		}
	}
	// TotalReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TotalReturned); err != nil {
			return err
		}
	}
	// Buffer {out} (1:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=PSAMPR_DISPLAY_INFO_BUFFER,pointer=ref}*(1))(2:{switch_type={alias=DOMAIN_DISPLAY_INFORMATION}(enum), alias=SAMPR_DISPLAY_INFO_BUFFER}(union))
	{
		if o.Buffer == nil {
			o.Buffer = &DisplayInfoBuffer{}
		}
		_swBuffer := uint16(o.DisplayInformationClass)
		if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryDisplayInformation3Request structure represents the SamrQueryDisplayInformation3 operation request
type QueryDisplayInformation3Request struct {
	// DomainHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a domain object.
	Domain *Handle `idl:"name:DomainHandle" json:"domain"`
	// DisplayInformationClass: An enumeration (see section 2.2.8.12) that indicates the
	// type of accounts, as well as the type of attributes on the accounts, to return via
	// the Buffer parameter.
	DisplayInformationClass DomainDisplayInformation `idl:"name:DisplayInformationClass" json:"display_information_class"`
	// Index: A cursor into an account-name–sorted list of accounts.
	Index uint32 `idl:"name:Index" json:"index"`
	// EntryCount: The number of accounts that the client is requesting on output.
	EntryCount uint32 `idl:"name:EntryCount" json:"entry_count"`
	// PreferredMaximumLength: The requested maximum number of bytes to return in Buffer;
	// this value overrides EntryCount if this value is reached before EntryCount is reached.
	PreferredMaximumLength uint32 `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
}

func (o *QueryDisplayInformation3Request) xxx_ToOp(ctx context.Context) *xxx_QueryDisplayInformation3Operation {
	if o == nil {
		return &xxx_QueryDisplayInformation3Operation{}
	}
	return &xxx_QueryDisplayInformation3Operation{
		Domain:                  o.Domain,
		DisplayInformationClass: o.DisplayInformationClass,
		Index:                   o.Index,
		EntryCount:              o.EntryCount,
		PreferredMaximumLength:  o.PreferredMaximumLength,
	}
}

func (o *QueryDisplayInformation3Request) xxx_FromOp(ctx context.Context, op *xxx_QueryDisplayInformation3Operation) {
	if o == nil {
		return
	}
	o.Domain = op.Domain
	o.DisplayInformationClass = op.DisplayInformationClass
	o.Index = op.Index
	o.EntryCount = op.EntryCount
	o.PreferredMaximumLength = op.PreferredMaximumLength
}
func (o *QueryDisplayInformation3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryDisplayInformation3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDisplayInformation3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryDisplayInformation3Response structure represents the SamrQueryDisplayInformation3 operation response
type QueryDisplayInformation3Response struct {
	// TotalAvailable: The number of bytes required to see a complete listing of accounts
	// specified by the DisplayInformationClass parameter.
	TotalAvailable uint32 `idl:"name:TotalAvailable" json:"total_available"`
	// TotalReturned: The number of bytes returned.<53>
	TotalReturned uint32 `idl:"name:TotalReturned" json:"total_returned"`
	// Buffer: The accounts that are returned.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Buffer *DisplayInfoBuffer `idl:"name:Buffer;switch_is:DisplayInformationClass" json:"buffer"`
	// Return: The SamrQueryDisplayInformation3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryDisplayInformation3Response) xxx_ToOp(ctx context.Context) *xxx_QueryDisplayInformation3Operation {
	if o == nil {
		return &xxx_QueryDisplayInformation3Operation{}
	}
	return &xxx_QueryDisplayInformation3Operation{
		TotalAvailable: o.TotalAvailable,
		TotalReturned:  o.TotalReturned,
		Buffer:         o.Buffer,
		Return:         o.Return,
	}
}

func (o *QueryDisplayInformation3Response) xxx_FromOp(ctx context.Context, op *xxx_QueryDisplayInformation3Operation) {
	if o == nil {
		return
	}
	o.TotalAvailable = op.TotalAvailable
	o.TotalReturned = op.TotalReturned
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *QueryDisplayInformation3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryDisplayInformation3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDisplayInformation3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddMultipleMembersToAliasOperation structure represents the SamrAddMultipleMembersToAlias operation
type xxx_AddMultipleMembersToAliasOperation struct {
	AliasHandle   *Handle   `idl:"name:AliasHandle" json:"alias_handle"`
	MembersBuffer *SIDArray `idl:"name:MembersBuffer" json:"members_buffer"`
	Return        int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_AddMultipleMembersToAliasOperation) OpNum() int { return 52 }

func (o *xxx_AddMultipleMembersToAliasOperation) OpName() string {
	return "/samr/v1/SamrAddMultipleMembersToAlias"
}

func (o *xxx_AddMultipleMembersToAliasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMultipleMembersToAliasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MembersBuffer {in} (1:{alias=PSAMPR_PSID_ARRAY}*(1))(2:{alias=SAMPR_PSID_ARRAY}(struct))
	{
		if o.MembersBuffer != nil {
			if err := o.MembersBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SIDArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMultipleMembersToAliasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MembersBuffer {in} (1:{alias=PSAMPR_PSID_ARRAY,pointer=ref}*(1))(2:{alias=SAMPR_PSID_ARRAY}(struct))
	{
		if o.MembersBuffer == nil {
			o.MembersBuffer = &SIDArray{}
		}
		if err := o.MembersBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMultipleMembersToAliasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMultipleMembersToAliasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMultipleMembersToAliasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddMultipleMembersToAliasRequest structure represents the SamrAddMultipleMembersToAlias operation request
type AddMultipleMembersToAliasRequest struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// an alias object.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	// MembersBuffer: A structure containing a list of SIDs to add as members to the alias.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	MembersBuffer *SIDArray `idl:"name:MembersBuffer" json:"members_buffer"`
}

func (o *AddMultipleMembersToAliasRequest) xxx_ToOp(ctx context.Context) *xxx_AddMultipleMembersToAliasOperation {
	if o == nil {
		return &xxx_AddMultipleMembersToAliasOperation{}
	}
	return &xxx_AddMultipleMembersToAliasOperation{
		AliasHandle:   o.AliasHandle,
		MembersBuffer: o.MembersBuffer,
	}
}

func (o *AddMultipleMembersToAliasRequest) xxx_FromOp(ctx context.Context, op *xxx_AddMultipleMembersToAliasOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
	o.MembersBuffer = op.MembersBuffer
}
func (o *AddMultipleMembersToAliasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddMultipleMembersToAliasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddMultipleMembersToAliasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddMultipleMembersToAliasResponse structure represents the SamrAddMultipleMembersToAlias operation response
type AddMultipleMembersToAliasResponse struct {
	// Return: The SamrAddMultipleMembersToAlias return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddMultipleMembersToAliasResponse) xxx_ToOp(ctx context.Context) *xxx_AddMultipleMembersToAliasOperation {
	if o == nil {
		return &xxx_AddMultipleMembersToAliasOperation{}
	}
	return &xxx_AddMultipleMembersToAliasOperation{
		Return: o.Return,
	}
}

func (o *AddMultipleMembersToAliasResponse) xxx_FromOp(ctx context.Context, op *xxx_AddMultipleMembersToAliasOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddMultipleMembersToAliasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddMultipleMembersToAliasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddMultipleMembersToAliasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveMultipleMembersFromAliasOperation structure represents the SamrRemoveMultipleMembersFromAlias operation
type xxx_RemoveMultipleMembersFromAliasOperation struct {
	AliasHandle   *Handle   `idl:"name:AliasHandle" json:"alias_handle"`
	MembersBuffer *SIDArray `idl:"name:MembersBuffer" json:"members_buffer"`
	Return        int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveMultipleMembersFromAliasOperation) OpNum() int { return 53 }

func (o *xxx_RemoveMultipleMembersFromAliasOperation) OpName() string {
	return "/samr/v1/SamrRemoveMultipleMembersFromAlias"
}

func (o *xxx_RemoveMultipleMembersFromAliasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMultipleMembersFromAliasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle != nil {
			if err := o.AliasHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MembersBuffer {in} (1:{alias=PSAMPR_PSID_ARRAY}*(1))(2:{alias=SAMPR_PSID_ARRAY}(struct))
	{
		if o.MembersBuffer != nil {
			if err := o.MembersBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SIDArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMultipleMembersFromAliasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AliasHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.AliasHandle == nil {
			o.AliasHandle = &Handle{}
		}
		if err := o.AliasHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MembersBuffer {in} (1:{alias=PSAMPR_PSID_ARRAY,pointer=ref}*(1))(2:{alias=SAMPR_PSID_ARRAY}(struct))
	{
		if o.MembersBuffer == nil {
			o.MembersBuffer = &SIDArray{}
		}
		if err := o.MembersBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMultipleMembersFromAliasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMultipleMembersFromAliasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMultipleMembersFromAliasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveMultipleMembersFromAliasRequest structure represents the SamrRemoveMultipleMembersFromAlias operation request
type RemoveMultipleMembersFromAliasRequest struct {
	// AliasHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// an alias object.
	AliasHandle *Handle `idl:"name:AliasHandle" json:"alias_handle"`
	// MembersBuffer: A structure containing a list of SIDs to remove from the alias's membership
	// list.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	MembersBuffer *SIDArray `idl:"name:MembersBuffer" json:"members_buffer"`
}

func (o *RemoveMultipleMembersFromAliasRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveMultipleMembersFromAliasOperation {
	if o == nil {
		return &xxx_RemoveMultipleMembersFromAliasOperation{}
	}
	return &xxx_RemoveMultipleMembersFromAliasOperation{
		AliasHandle:   o.AliasHandle,
		MembersBuffer: o.MembersBuffer,
	}
}

func (o *RemoveMultipleMembersFromAliasRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveMultipleMembersFromAliasOperation) {
	if o == nil {
		return
	}
	o.AliasHandle = op.AliasHandle
	o.MembersBuffer = op.MembersBuffer
}
func (o *RemoveMultipleMembersFromAliasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveMultipleMembersFromAliasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMultipleMembersFromAliasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveMultipleMembersFromAliasResponse structure represents the SamrRemoveMultipleMembersFromAlias operation response
type RemoveMultipleMembersFromAliasResponse struct {
	// Return: The SamrRemoveMultipleMembersFromAlias return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveMultipleMembersFromAliasResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveMultipleMembersFromAliasOperation {
	if o == nil {
		return &xxx_RemoveMultipleMembersFromAliasOperation{}
	}
	return &xxx_RemoveMultipleMembersFromAliasOperation{
		Return: o.Return,
	}
}

func (o *RemoveMultipleMembersFromAliasResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveMultipleMembersFromAliasOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveMultipleMembersFromAliasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveMultipleMembersFromAliasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMultipleMembersFromAliasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OEMChangePasswordUser2Operation structure represents the SamrOemChangePasswordUser2 operation
type xxx_OEMChangePasswordUser2Operation struct {
	ServerName                         *String                 `idl:"name:ServerName;pointer:unique" json:"server_name"`
	UserName                           *String                 `idl:"name:UserName" json:"user_name"`
	NewPasswordEncryptedWithOldLM      *EncryptedUserPassword  `idl:"name:NewPasswordEncryptedWithOldLm;pointer:unique" json:"new_password_encrypted_with_old_lm"`
	OldLMOWFPasswordEncryptedWithNewLM *EncryptedLMOWFPassword `idl:"name:OldLmOwfPasswordEncryptedWithNewLm;pointer:unique" json:"old_lm_owf_password_encrypted_with_new_lm"`
	Return                             int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_OEMChangePasswordUser2Operation) OpNum() int { return 54 }

func (o *xxx_OEMChangePasswordUser2Operation) OpName() string {
	return "/samr/v1/SamrOemChangePasswordUser2"
}

func (o *xxx_OEMChangePasswordUser2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OEMChangePasswordUser2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{pointer=unique, alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.ServerName != nil {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerName != nil {
					if err := o.ServerName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// UserName {in} (1:{alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.UserName != nil {
			if err := o.UserName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// NewPasswordEncryptedWithOldLm {in} (1:{pointer=unique, alias=PSAMPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=SAMPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.NewPasswordEncryptedWithOldLM != nil {
			_ptr_NewPasswordEncryptedWithOldLm := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewPasswordEncryptedWithOldLM != nil {
					if err := o.NewPasswordEncryptedWithOldLM.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewPasswordEncryptedWithOldLM, _ptr_NewPasswordEncryptedWithOldLm); err != nil {
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
	// OldLmOwfPasswordEncryptedWithNewLm {in} (1:{pointer=unique, alias=PENCRYPTED_LM_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_LM_OWF_PASSWORD}(struct))
	{
		if o.OldLMOWFPasswordEncryptedWithNewLM != nil {
			_ptr_OldLmOwfPasswordEncryptedWithNewLm := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldLMOWFPasswordEncryptedWithNewLM != nil {
					if err := o.OldLMOWFPasswordEncryptedWithNewLM.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedLMOWFPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldLMOWFPasswordEncryptedWithNewLM, _ptr_OldLmOwfPasswordEncryptedWithNewLm); err != nil {
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

func (o *xxx_OEMChangePasswordUser2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{pointer=unique, alias=PRPC_STRING}*(1))(2:{alias=RPC_STRING}(struct))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerName == nil {
				o.ServerName = &String{}
			}
			if err := o.ServerName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(**String) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UserName {in} (1:{alias=PRPC_STRING,pointer=ref}*(1))(2:{alias=RPC_STRING}(struct))
	{
		if o.UserName == nil {
			o.UserName = &String{}
		}
		if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NewPasswordEncryptedWithOldLm {in} (1:{pointer=unique, alias=PSAMPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=SAMPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_NewPasswordEncryptedWithOldLm := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewPasswordEncryptedWithOldLM == nil {
				o.NewPasswordEncryptedWithOldLM = &EncryptedUserPassword{}
			}
			if err := o.NewPasswordEncryptedWithOldLM.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_NewPasswordEncryptedWithOldLm := func(ptr interface{}) { o.NewPasswordEncryptedWithOldLM = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.NewPasswordEncryptedWithOldLM, _s_NewPasswordEncryptedWithOldLm, _ptr_NewPasswordEncryptedWithOldLm); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OldLmOwfPasswordEncryptedWithNewLm {in} (1:{pointer=unique, alias=PENCRYPTED_LM_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_LM_OWF_PASSWORD}(struct))
	{
		_ptr_OldLmOwfPasswordEncryptedWithNewLm := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldLMOWFPasswordEncryptedWithNewLM == nil {
				o.OldLMOWFPasswordEncryptedWithNewLM = &EncryptedLMOWFPassword{}
			}
			if err := o.OldLMOWFPasswordEncryptedWithNewLM.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OldLmOwfPasswordEncryptedWithNewLm := func(ptr interface{}) { o.OldLMOWFPasswordEncryptedWithNewLM = *ptr.(**EncryptedLMOWFPassword) }
		if err := w.ReadPointer(&o.OldLMOWFPasswordEncryptedWithNewLM, _s_OldLmOwfPasswordEncryptedWithNewLm, _ptr_OldLmOwfPasswordEncryptedWithNewLm); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OEMChangePasswordUser2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OEMChangePasswordUser2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OEMChangePasswordUser2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OEMChangePasswordUser2Request structure represents the SamrOemChangePasswordUser2 operation request
type OEMChangePasswordUser2Request struct {
	// ServerName: A counted string, encoded in the OEM character set, containing the NETBIOS
	// name of the server; this parameter MAY<65> be ignored by the server.
	ServerName *String `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// UserName: A counted string, encoded in the OEM character set, containing the name
	// of the user whose password is to be changed; see message processing later in this
	// section for details on how this value is used as a database key to locate the account
	// that is the target of this password change operation.
	UserName *String `idl:"name:UserName" json:"user_name"`
	// NewPasswordEncryptedWithOldLm: A cleartext password encrypted according to the specification
	// of SAMPR_ENCRYPTED_USER_PASSWORD (section 2.2.6.21), where the key is the LM hash
	// of the existing password for the target user (as presented by the client). The cleartext
	// password MUST be encoded in an OEM code page character set (as opposed to UTF-16).
	NewPasswordEncryptedWithOldLM *EncryptedUserPassword `idl:"name:NewPasswordEncryptedWithOldLm;pointer:unique" json:"new_password_encrypted_with_old_lm"`
	// OldLmOwfPasswordEncryptedWithNewLm: The LM hash of the target user's existing password
	// (as presented by the client) encrypted according to the specification of ENCRYPTED_LM_OWF_PASSWORD
	// (section 2.2.7.3), where the key is the LM hash of the cleartext password obtained
	// from decrypting NewPasswordEncryptedWithOldLm (see the preceding description for
	// decryption details).
	OldLMOWFPasswordEncryptedWithNewLM *EncryptedLMOWFPassword `idl:"name:OldLmOwfPasswordEncryptedWithNewLm;pointer:unique" json:"old_lm_owf_password_encrypted_with_new_lm"`
}

func (o *OEMChangePasswordUser2Request) xxx_ToOp(ctx context.Context) *xxx_OEMChangePasswordUser2Operation {
	if o == nil {
		return &xxx_OEMChangePasswordUser2Operation{}
	}
	return &xxx_OEMChangePasswordUser2Operation{
		ServerName:                         o.ServerName,
		UserName:                           o.UserName,
		NewPasswordEncryptedWithOldLM:      o.NewPasswordEncryptedWithOldLM,
		OldLMOWFPasswordEncryptedWithNewLM: o.OldLMOWFPasswordEncryptedWithNewLM,
	}
}

func (o *OEMChangePasswordUser2Request) xxx_FromOp(ctx context.Context, op *xxx_OEMChangePasswordUser2Operation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.UserName = op.UserName
	o.NewPasswordEncryptedWithOldLM = op.NewPasswordEncryptedWithOldLM
	o.OldLMOWFPasswordEncryptedWithNewLM = op.OldLMOWFPasswordEncryptedWithNewLM
}
func (o *OEMChangePasswordUser2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OEMChangePasswordUser2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OEMChangePasswordUser2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OEMChangePasswordUser2Response structure represents the SamrOemChangePasswordUser2 operation response
type OEMChangePasswordUser2Response struct {
	// Return: The SamrOemChangePasswordUser2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OEMChangePasswordUser2Response) xxx_ToOp(ctx context.Context) *xxx_OEMChangePasswordUser2Operation {
	if o == nil {
		return &xxx_OEMChangePasswordUser2Operation{}
	}
	return &xxx_OEMChangePasswordUser2Operation{
		Return: o.Return,
	}
}

func (o *OEMChangePasswordUser2Response) xxx_FromOp(ctx context.Context, op *xxx_OEMChangePasswordUser2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *OEMChangePasswordUser2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OEMChangePasswordUser2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OEMChangePasswordUser2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnicodeChangePasswordUser2Operation structure represents the SamrUnicodeChangePasswordUser2 operation
type xxx_UnicodeChangePasswordUser2Operation struct {
	ServerName                         *dtyp.UnicodeString     `idl:"name:ServerName;pointer:unique" json:"server_name"`
	UserName                           *dtyp.UnicodeString     `idl:"name:UserName" json:"user_name"`
	NewPasswordEncryptedWithOldNT      *EncryptedUserPassword  `idl:"name:NewPasswordEncryptedWithOldNt;pointer:unique" json:"new_password_encrypted_with_old_nt"`
	OldNTOWFPasswordEncryptedWithNewNT *EncryptedNTOWFPassword `idl:"name:OldNtOwfPasswordEncryptedWithNewNt;pointer:unique" json:"old_nt_owf_password_encrypted_with_new_nt"`
	LMPresent                          uint8                   `idl:"name:LmPresent" json:"lm_present"`
	NewPasswordEncryptedWithOldLM      *EncryptedUserPassword  `idl:"name:NewPasswordEncryptedWithOldLm;pointer:unique" json:"new_password_encrypted_with_old_lm"`
	OldLMOWFPasswordEncryptedWithNewNT *EncryptedLMOWFPassword `idl:"name:OldLmOwfPasswordEncryptedWithNewNt;pointer:unique" json:"old_lm_owf_password_encrypted_with_new_nt"`
	Return                             int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_UnicodeChangePasswordUser2Operation) OpNum() int { return 55 }

func (o *xxx_UnicodeChangePasswordUser2Operation) OpName() string {
	return "/samr/v1/SamrUnicodeChangePasswordUser2"
}

func (o *xxx_UnicodeChangePasswordUser2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnicodeChangePasswordUser2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{pointer=unique, alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ServerName != nil {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServerName != nil {
					if err := o.ServerName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// UserName {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.UserName != nil {
			if err := o.UserName.MarshalNDR(ctx, w); err != nil {
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
	// NewPasswordEncryptedWithOldNt {in} (1:{pointer=unique, alias=PSAMPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=SAMPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.NewPasswordEncryptedWithOldNT != nil {
			_ptr_NewPasswordEncryptedWithOldNt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewPasswordEncryptedWithOldNT != nil {
					if err := o.NewPasswordEncryptedWithOldNT.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewPasswordEncryptedWithOldNT, _ptr_NewPasswordEncryptedWithOldNt); err != nil {
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
	// OldNtOwfPasswordEncryptedWithNewNt {in} (1:{pointer=unique, alias=PENCRYPTED_NT_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_NT_OWF_PASSWORD}(struct))
	{
		if o.OldNTOWFPasswordEncryptedWithNewNT != nil {
			_ptr_OldNtOwfPasswordEncryptedWithNewNt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldNTOWFPasswordEncryptedWithNewNT != nil {
					if err := o.OldNTOWFPasswordEncryptedWithNewNT.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedNTOWFPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldNTOWFPasswordEncryptedWithNewNT, _ptr_OldNtOwfPasswordEncryptedWithNewNt); err != nil {
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
	// LmPresent {in} (1:(uchar))
	{
		if err := w.WriteData(o.LMPresent); err != nil {
			return err
		}
	}
	// NewPasswordEncryptedWithOldLm {in} (1:{pointer=unique, alias=PSAMPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=SAMPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.NewPasswordEncryptedWithOldLM != nil {
			_ptr_NewPasswordEncryptedWithOldLm := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewPasswordEncryptedWithOldLM != nil {
					if err := o.NewPasswordEncryptedWithOldLM.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewPasswordEncryptedWithOldLM, _ptr_NewPasswordEncryptedWithOldLm); err != nil {
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
	// OldLmOwfPasswordEncryptedWithNewNt {in} (1:{pointer=unique, alias=PENCRYPTED_LM_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_LM_OWF_PASSWORD}(struct))
	{
		if o.OldLMOWFPasswordEncryptedWithNewNT != nil {
			_ptr_OldLmOwfPasswordEncryptedWithNewNt := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldLMOWFPasswordEncryptedWithNewNT != nil {
					if err := o.OldLMOWFPasswordEncryptedWithNewNT.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedLMOWFPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldLMOWFPasswordEncryptedWithNewNT, _ptr_OldLmOwfPasswordEncryptedWithNewNt); err != nil {
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

func (o *xxx_UnicodeChangePasswordUser2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{pointer=unique, alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServerName == nil {
				o.ServerName = &dtyp.UnicodeString{}
			}
			if err := o.ServerName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UserName {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.UserName == nil {
			o.UserName = &dtyp.UnicodeString{}
		}
		if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NewPasswordEncryptedWithOldNt {in} (1:{pointer=unique, alias=PSAMPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=SAMPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_NewPasswordEncryptedWithOldNt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewPasswordEncryptedWithOldNT == nil {
				o.NewPasswordEncryptedWithOldNT = &EncryptedUserPassword{}
			}
			if err := o.NewPasswordEncryptedWithOldNT.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_NewPasswordEncryptedWithOldNt := func(ptr interface{}) { o.NewPasswordEncryptedWithOldNT = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.NewPasswordEncryptedWithOldNT, _s_NewPasswordEncryptedWithOldNt, _ptr_NewPasswordEncryptedWithOldNt); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OldNtOwfPasswordEncryptedWithNewNt {in} (1:{pointer=unique, alias=PENCRYPTED_NT_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_NT_OWF_PASSWORD}(struct))
	{
		_ptr_OldNtOwfPasswordEncryptedWithNewNt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldNTOWFPasswordEncryptedWithNewNT == nil {
				o.OldNTOWFPasswordEncryptedWithNewNT = &EncryptedNTOWFPassword{}
			}
			if err := o.OldNTOWFPasswordEncryptedWithNewNT.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OldNtOwfPasswordEncryptedWithNewNt := func(ptr interface{}) { o.OldNTOWFPasswordEncryptedWithNewNT = *ptr.(**EncryptedNTOWFPassword) }
		if err := w.ReadPointer(&o.OldNTOWFPasswordEncryptedWithNewNT, _s_OldNtOwfPasswordEncryptedWithNewNt, _ptr_OldNtOwfPasswordEncryptedWithNewNt); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// LmPresent {in} (1:(uchar))
	{
		if err := w.ReadData(&o.LMPresent); err != nil {
			return err
		}
	}
	// NewPasswordEncryptedWithOldLm {in} (1:{pointer=unique, alias=PSAMPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=SAMPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_NewPasswordEncryptedWithOldLm := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewPasswordEncryptedWithOldLM == nil {
				o.NewPasswordEncryptedWithOldLM = &EncryptedUserPassword{}
			}
			if err := o.NewPasswordEncryptedWithOldLM.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_NewPasswordEncryptedWithOldLm := func(ptr interface{}) { o.NewPasswordEncryptedWithOldLM = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.NewPasswordEncryptedWithOldLM, _s_NewPasswordEncryptedWithOldLm, _ptr_NewPasswordEncryptedWithOldLm); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OldLmOwfPasswordEncryptedWithNewNt {in} (1:{pointer=unique, alias=PENCRYPTED_LM_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_LM_OWF_PASSWORD}(struct))
	{
		_ptr_OldLmOwfPasswordEncryptedWithNewNt := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldLMOWFPasswordEncryptedWithNewNT == nil {
				o.OldLMOWFPasswordEncryptedWithNewNT = &EncryptedLMOWFPassword{}
			}
			if err := o.OldLMOWFPasswordEncryptedWithNewNT.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_OldLmOwfPasswordEncryptedWithNewNt := func(ptr interface{}) { o.OldLMOWFPasswordEncryptedWithNewNT = *ptr.(**EncryptedLMOWFPassword) }
		if err := w.ReadPointer(&o.OldLMOWFPasswordEncryptedWithNewNT, _s_OldLmOwfPasswordEncryptedWithNewNt, _ptr_OldLmOwfPasswordEncryptedWithNewNt); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnicodeChangePasswordUser2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnicodeChangePasswordUser2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnicodeChangePasswordUser2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UnicodeChangePasswordUser2Request structure represents the SamrUnicodeChangePasswordUser2 operation request
type UnicodeChangePasswordUser2Request struct {
	// ServerName: A null-terminated string containing the NETBIOS name of the server; this
	// parameter MAY<66> be ignored by the server.
	ServerName *dtyp.UnicodeString `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// UserName: The name of the user. See the message processing later in this section
	// for details on how this value is used as a database key to locate the account that
	// is the target of this password change operation.
	UserName *dtyp.UnicodeString `idl:"name:UserName" json:"user_name"`
	// NewPasswordEncryptedWithOldNt: A cleartext password encrypted according to the specification
	// of SAMPR_ENCRYPTED_USER_PASSWORD (section 2.2.6.21), where the key is the NT hash
	// of the existing password for the target user (as presented by the client in the OldNtOwfPasswordEncryptedWithNewNt
	// parameter).
	NewPasswordEncryptedWithOldNT *EncryptedUserPassword `idl:"name:NewPasswordEncryptedWithOldNt;pointer:unique" json:"new_password_encrypted_with_old_nt"`
	// OldNtOwfPasswordEncryptedWithNewNt: The NT hash of the target user's existing password
	// (as presented by the client) encrypted according to the specification of ENCRYPTED_LM_OWF_PASSWORD
	// (section 2.2.7.3), where the key is the NT hash of the cleartext password obtained
	// from decrypting NewPasswordEncryptedWithOldNt.
	OldNTOWFPasswordEncryptedWithNewNT *EncryptedNTOWFPassword `idl:"name:OldNtOwfPasswordEncryptedWithNewNt;pointer:unique" json:"old_nt_owf_password_encrypted_with_new_nt"`
	// LmPresent: If this parameter is zero, NewPasswordEncryptedWithOldLm and OldLmOwfPasswordEncryptedWithNewNt
	// MUST be ignored; otherwise these fields MUST be processed.
	LMPresent uint8 `idl:"name:LmPresent" json:"lm_present"`
	// NewPasswordEncryptedWithOldLm: A cleartext password encrypted according to the specification
	// of SAMPR_ENCRYPTED_USER_PASSWORD, where the key is the LM hash of the existing password
	// for the target user (as presented by the client).
	NewPasswordEncryptedWithOldLM *EncryptedUserPassword `idl:"name:NewPasswordEncryptedWithOldLm;pointer:unique" json:"new_password_encrypted_with_old_lm"`
	// OldLmOwfPasswordEncryptedWithNewNt: The LM hash the target user's existing password
	// (as presented by the client) encrypted according to the specification of ENCRYPTED_LM_OWF_PASSWORD,
	// where the key is the NT hash of the cleartext password obtained from decrypting NewPasswordEncryptedWithOldNt.
	OldLMOWFPasswordEncryptedWithNewNT *EncryptedLMOWFPassword `idl:"name:OldLmOwfPasswordEncryptedWithNewNt;pointer:unique" json:"old_lm_owf_password_encrypted_with_new_nt"`
}

func (o *UnicodeChangePasswordUser2Request) xxx_ToOp(ctx context.Context) *xxx_UnicodeChangePasswordUser2Operation {
	if o == nil {
		return &xxx_UnicodeChangePasswordUser2Operation{}
	}
	return &xxx_UnicodeChangePasswordUser2Operation{
		ServerName:                         o.ServerName,
		UserName:                           o.UserName,
		NewPasswordEncryptedWithOldNT:      o.NewPasswordEncryptedWithOldNT,
		OldNTOWFPasswordEncryptedWithNewNT: o.OldNTOWFPasswordEncryptedWithNewNT,
		LMPresent:                          o.LMPresent,
		NewPasswordEncryptedWithOldLM:      o.NewPasswordEncryptedWithOldLM,
		OldLMOWFPasswordEncryptedWithNewNT: o.OldLMOWFPasswordEncryptedWithNewNT,
	}
}

func (o *UnicodeChangePasswordUser2Request) xxx_FromOp(ctx context.Context, op *xxx_UnicodeChangePasswordUser2Operation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.UserName = op.UserName
	o.NewPasswordEncryptedWithOldNT = op.NewPasswordEncryptedWithOldNT
	o.OldNTOWFPasswordEncryptedWithNewNT = op.OldNTOWFPasswordEncryptedWithNewNT
	o.LMPresent = op.LMPresent
	o.NewPasswordEncryptedWithOldLM = op.NewPasswordEncryptedWithOldLM
	o.OldLMOWFPasswordEncryptedWithNewNT = op.OldLMOWFPasswordEncryptedWithNewNT
}
func (o *UnicodeChangePasswordUser2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UnicodeChangePasswordUser2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnicodeChangePasswordUser2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UnicodeChangePasswordUser2Response structure represents the SamrUnicodeChangePasswordUser2 operation response
type UnicodeChangePasswordUser2Response struct {
	// Return: The SamrUnicodeChangePasswordUser2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UnicodeChangePasswordUser2Response) xxx_ToOp(ctx context.Context) *xxx_UnicodeChangePasswordUser2Operation {
	if o == nil {
		return &xxx_UnicodeChangePasswordUser2Operation{}
	}
	return &xxx_UnicodeChangePasswordUser2Operation{
		Return: o.Return,
	}
}

func (o *UnicodeChangePasswordUser2Response) xxx_FromOp(ctx context.Context, op *xxx_UnicodeChangePasswordUser2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UnicodeChangePasswordUser2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UnicodeChangePasswordUser2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnicodeChangePasswordUser2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDomainPasswordInformationOperation structure represents the SamrGetDomainPasswordInformation operation
type xxx_GetDomainPasswordInformationOperation struct {
	_                   *dtyp.UnicodeString            `idl:"name:Unused;pointer:unique"`
	PasswordInformation *UserDomainPasswordInformation `idl:"name:PasswordInformation" json:"password_information"`
	Return              int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDomainPasswordInformationOperation) OpNum() int { return 56 }

func (o *xxx_GetDomainPasswordInformationOperation) OpName() string {
	return "/samr/v1/SamrGetDomainPasswordInformation"
}

func (o *xxx_GetDomainPasswordInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDomainPasswordInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Unused {in} (1:{pointer=unique, alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		// reserved Unused
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDomainPasswordInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Unused {in} (1:{pointer=unique, alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		// reserved Unused
		var _Unused *dtyp.UnicodeString
		_ptr_Unused := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if _Unused == nil {
				_Unused = &dtyp.UnicodeString{}
			}
			if err := _Unused.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Unused := func(ptr interface{}) { _Unused = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&_Unused, _s_Unused, _ptr_Unused); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDomainPasswordInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDomainPasswordInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PasswordInformation {out} (1:{alias=PUSER_DOMAIN_PASSWORD_INFORMATION}*(1))(2:{alias=USER_DOMAIN_PASSWORD_INFORMATION}(struct))
	{
		if o.PasswordInformation != nil {
			if err := o.PasswordInformation.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UserDomainPasswordInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDomainPasswordInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PasswordInformation {out} (1:{alias=PUSER_DOMAIN_PASSWORD_INFORMATION,pointer=ref}*(1))(2:{alias=USER_DOMAIN_PASSWORD_INFORMATION}(struct))
	{
		if o.PasswordInformation == nil {
			o.PasswordInformation = &UserDomainPasswordInformation{}
		}
		if err := o.PasswordInformation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDomainPasswordInformationRequest structure represents the SamrGetDomainPasswordInformation operation request
type GetDomainPasswordInformationRequest struct {
}

func (o *GetDomainPasswordInformationRequest) xxx_ToOp(ctx context.Context) *xxx_GetDomainPasswordInformationOperation {
	if o == nil {
		return &xxx_GetDomainPasswordInformationOperation{}
	}
	return &xxx_GetDomainPasswordInformationOperation{}
}

func (o *GetDomainPasswordInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDomainPasswordInformationOperation) {
	if o == nil {
		return
	}
}
func (o *GetDomainPasswordInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDomainPasswordInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDomainPasswordInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDomainPasswordInformationResponse structure represents the SamrGetDomainPasswordInformation operation response
type GetDomainPasswordInformationResponse struct {
	// PasswordInformation: Password policy information from the account domain.
	//
	// There is no security enforced for this method beyond the server-wide access check
	// specified in section 3.1.2.1.
	PasswordInformation *UserDomainPasswordInformation `idl:"name:PasswordInformation" json:"password_information"`
	// Return: The SamrGetDomainPasswordInformation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDomainPasswordInformationResponse) xxx_ToOp(ctx context.Context) *xxx_GetDomainPasswordInformationOperation {
	if o == nil {
		return &xxx_GetDomainPasswordInformationOperation{}
	}
	return &xxx_GetDomainPasswordInformationOperation{
		PasswordInformation: o.PasswordInformation,
		Return:              o.Return,
	}
}

func (o *GetDomainPasswordInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDomainPasswordInformationOperation) {
	if o == nil {
		return
	}
	o.PasswordInformation = op.PasswordInformation
	o.Return = op.Return
}
func (o *GetDomainPasswordInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDomainPasswordInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDomainPasswordInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Connect2Operation structure represents the SamrConnect2 operation
type xxx_Connect2Operation struct {
	ServerName    string  `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	Server        *Handle `idl:"name:ServerHandle" json:"server"`
	DesiredAccess uint32  `idl:"name:DesiredAccess" json:"desired_access"`
	Return        int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_Connect2Operation) OpNum() int { return 57 }

func (o *xxx_Connect2Operation) OpName() string { return "/samr/v1/SamrConnect2" }

func (o *xxx_Connect2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=PSAMPR_SERVER_NAME}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=PSAMPR_SERVER_NAME}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ServerHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ServerHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &Handle{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Connect2Request structure represents the SamrConnect2 operation request
type Connect2Request struct {
	// ServerName: The null-terminated NETBIOS name of the server; this parameter MAY<46>
	// be ignored on receipt.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// DesiredAccess: An ACCESS_MASK that indicates the access requested for ServerHandle
	// on output. See section 2.2.1.3 for a listing of possible values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *Connect2Request) xxx_ToOp(ctx context.Context) *xxx_Connect2Operation {
	if o == nil {
		return &xxx_Connect2Operation{}
	}
	return &xxx_Connect2Operation{
		ServerName:    o.ServerName,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *Connect2Request) xxx_FromOp(ctx context.Context, op *xxx_Connect2Operation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
}
func (o *Connect2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Connect2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Connect2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Connect2Response structure represents the SamrConnect2 operation response
type Connect2Response struct {
	// ServerHandle: An RPC context handle, as specified in section 2.2.7.2.
	Server *Handle `idl:"name:ServerHandle" json:"server"`
	// Return: The SamrConnect2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Connect2Response) xxx_ToOp(ctx context.Context) *xxx_Connect2Operation {
	if o == nil {
		return &xxx_Connect2Operation{}
	}
	return &xxx_Connect2Operation{
		Server: o.Server,
		Return: o.Return,
	}
}

func (o *Connect2Response) xxx_FromOp(ctx context.Context, op *xxx_Connect2Operation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Return = op.Return
}
func (o *Connect2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Connect2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Connect2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInformationUser2Operation structure represents the SamrSetInformationUser2 operation
type xxx_SetInformationUser2Operation struct {
	UserHandle           *Handle              `idl:"name:UserHandle" json:"user_handle"`
	UserInformationClass UserInformationClass `idl:"name:UserInformationClass" json:"user_information_class"`
	Buffer               *UserInfoBuffer      `idl:"name:Buffer;switch_is:UserInformationClass" json:"buffer"`
	Return               int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInformationUser2Operation) OpNum() int { return 58 }

func (o *xxx_SetInformationUser2Operation) OpName() string { return "/samr/v1/SamrSetInformationUser2" }

func (o *xxx_SetInformationUser2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationUser2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle != nil {
			if err := o.UserHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// UserInformationClass {in} (1:{alias=USER_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.UserInformationClass)); err != nil {
			return err
		}
	}
	// Buffer {in} (1:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=PSAMPR_USER_INFO_BUFFER}*(1))(2:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=SAMPR_USER_INFO_BUFFER}(union))
	{
		_swBuffer := uint16(o.UserInformationClass)
		if o.Buffer != nil {
			if err := o.Buffer.MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		} else {
			if err := (&UserInfoBuffer{}).MarshalUnionNDR(ctx, w, _swBuffer); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationUser2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UserHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.UserHandle == nil {
			o.UserHandle = &Handle{}
		}
		if err := o.UserHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// UserInformationClass {in} (1:{alias=USER_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.UserInformationClass)); err != nil {
			return err
		}
	}
	// Buffer {in} (1:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=PSAMPR_USER_INFO_BUFFER,pointer=ref}*(1))(2:{switch_type={alias=USER_INFORMATION_CLASS}(enum), alias=SAMPR_USER_INFO_BUFFER}(union))
	{
		if o.Buffer == nil {
			o.Buffer = &UserInfoBuffer{}
		}
		_swBuffer := uint16(o.UserInformationClass)
		if err := o.Buffer.UnmarshalUnionNDR(ctx, w, _swBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationUser2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationUser2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInformationUser2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInformationUser2Request structure represents the SamrSetInformationUser2 operation request
type SetInformationUser2Request struct {
	// UserHandle: An RPC context handle, as specified in section 2.2.7.2, representing
	// a user object.
	UserHandle *Handle `idl:"name:UserHandle" json:"user_handle"`
	// UserInformationClass: An enumeration indicating which attributes to update. See section
	// 2.2.6.28 for a listing of possible values.
	UserInformationClass UserInformationClass `idl:"name:UserInformationClass" json:"user_information_class"`
	// Buffer: The requested attributes and values to update. See section 2.2.6.29 for structure
	// details.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	Buffer *UserInfoBuffer `idl:"name:Buffer;switch_is:UserInformationClass" json:"buffer"`
}

func (o *SetInformationUser2Request) xxx_ToOp(ctx context.Context) *xxx_SetInformationUser2Operation {
	if o == nil {
		return &xxx_SetInformationUser2Operation{}
	}
	return &xxx_SetInformationUser2Operation{
		UserHandle:           o.UserHandle,
		UserInformationClass: o.UserInformationClass,
		Buffer:               o.Buffer,
	}
}

func (o *SetInformationUser2Request) xxx_FromOp(ctx context.Context, op *xxx_SetInformationUser2Operation) {
	if o == nil {
		return
	}
	o.UserHandle = op.UserHandle
	o.UserInformationClass = op.UserInformationClass
	o.Buffer = op.Buffer
}
func (o *SetInformationUser2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetInformationUser2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationUser2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInformationUser2Response structure represents the SamrSetInformationUser2 operation response
type SetInformationUser2Response struct {
	// Return: The SamrSetInformationUser2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInformationUser2Response) xxx_ToOp(ctx context.Context) *xxx_SetInformationUser2Operation {
	if o == nil {
		return &xxx_SetInformationUser2Operation{}
	}
	return &xxx_SetInformationUser2Operation{
		Return: o.Return,
	}
}

func (o *SetInformationUser2Response) xxx_FromOp(ctx context.Context, op *xxx_SetInformationUser2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetInformationUser2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetInformationUser2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInformationUser2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Connect4Operation structure represents the SamrConnect4 operation
type xxx_Connect4Operation struct {
	ServerName     string  `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	Server         *Handle `idl:"name:ServerHandle" json:"server"`
	ClientRevision uint32  `idl:"name:ClientRevision" json:"client_revision"`
	DesiredAccess  uint32  `idl:"name:DesiredAccess" json:"desired_access"`
	Return         int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_Connect4Operation) OpNum() int { return 62 }

func (o *xxx_Connect4Operation) OpName() string { return "/samr/v1/SamrConnect4" }

func (o *xxx_Connect4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=PSAMPR_SERVER_NAME}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.WriteData(o.ClientRevision); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=PSAMPR_SERVER_NAME}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ClientRevision); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ServerHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ServerHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &Handle{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Connect4Request structure represents the SamrConnect4 operation request
type Connect4Request struct {
	// ServerName: The null-terminated NETBIOS name of the server; this parameter MAY<45>
	// be ignored on receipt.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// ClientRevision: Indicates the revision (for this protocol) of the client. The value
	// MUST be set to 2 and MUST be ignored.
	ClientRevision uint32 `idl:"name:ClientRevision" json:"client_revision"`
	// DesiredAccess: An ACCESS_MASK that indicates the access requested for ServerHandle
	// on output. See section 2.2.1.3 for a listing of possible values.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *Connect4Request) xxx_ToOp(ctx context.Context) *xxx_Connect4Operation {
	if o == nil {
		return &xxx_Connect4Operation{}
	}
	return &xxx_Connect4Operation{
		ServerName:     o.ServerName,
		ClientRevision: o.ClientRevision,
		DesiredAccess:  o.DesiredAccess,
	}
}

func (o *Connect4Request) xxx_FromOp(ctx context.Context, op *xxx_Connect4Operation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.ClientRevision = op.ClientRevision
	o.DesiredAccess = op.DesiredAccess
}
func (o *Connect4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Connect4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Connect4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Connect4Response structure represents the SamrConnect4 operation response
type Connect4Response struct {
	// ServerHandle: An RPC context handle, as specified in section 2.2.7.2.
	Server *Handle `idl:"name:ServerHandle" json:"server"`
	// Return: The SamrConnect4 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Connect4Response) xxx_ToOp(ctx context.Context) *xxx_Connect4Operation {
	if o == nil {
		return &xxx_Connect4Operation{}
	}
	return &xxx_Connect4Operation{
		Server: o.Server,
		Return: o.Return,
	}
}

func (o *Connect4Response) xxx_FromOp(ctx context.Context, op *xxx_Connect4Operation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Return = op.Return
}
func (o *Connect4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Connect4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Connect4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Connect5Operation structure represents the SamrConnect5 operation
type xxx_Connect5Operation struct {
	ServerName      string        `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	DesiredAccess   uint32        `idl:"name:DesiredAccess" json:"desired_access"`
	InVersion       uint32        `idl:"name:InVersion" json:"in_version"`
	InRevisionInfo  *RevisionInfo `idl:"name:InRevisionInfo;switch_is:InVersion" json:"in_revision_info"`
	OutVersion      uint32        `idl:"name:OutVersion" json:"out_version"`
	OutRevisionInfo *RevisionInfo `idl:"name:OutRevisionInfo;switch_is:*OutVersion" json:"out_revision_info"`
	Server          *Handle       `idl:"name:ServerHandle" json:"server"`
	Return          int32         `idl:"name:Return" json:"return"`
}

func (o *xxx_Connect5Operation) OpNum() int { return 64 }

func (o *xxx_Connect5Operation) OpName() string { return "/samr/v1/SamrConnect5" }

func (o *xxx_Connect5Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect5Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=PSAMPR_SERVER_NAME}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	// InVersion {in} (1:(uint32))
	{
		if err := w.WriteData(o.InVersion); err != nil {
			return err
		}
	}
	// InRevisionInfo {in} (1:{pointer=ref}*(1))(2:{switch_type={}(uint32), alias=SAMPR_REVISION_INFO}(union))
	{
		_swInRevisionInfo := uint32(o.InVersion)
		if o.InRevisionInfo != nil {
			if err := o.InRevisionInfo.MarshalUnionNDR(ctx, w, _swInRevisionInfo); err != nil {
				return err
			}
		} else {
			if err := (&RevisionInfo{}).MarshalUnionNDR(ctx, w, _swInRevisionInfo); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_Connect5Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=PSAMPR_SERVER_NAME}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	// InVersion {in} (1:(uint32))
	{
		if err := w.ReadData(&o.InVersion); err != nil {
			return err
		}
	}
	// InRevisionInfo {in} (1:{pointer=ref}*(1))(2:{switch_type={}(uint32), alias=SAMPR_REVISION_INFO}(union))
	{
		if o.InRevisionInfo == nil {
			o.InRevisionInfo = &RevisionInfo{}
		}
		_swInRevisionInfo := uint32(o.InVersion)
		if err := o.InRevisionInfo.UnmarshalUnionNDR(ctx, w, _swInRevisionInfo); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect5Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect5Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// OutVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.OutVersion); err != nil {
			return err
		}
	}
	// OutRevisionInfo {out} (1:{pointer=ref}*(1))(2:{switch_type={}(uint32), alias=SAMPR_REVISION_INFO}(union))
	{
		_swOutRevisionInfo := uint32(o.OutVersion)
		if o.OutRevisionInfo != nil {
			if err := o.OutRevisionInfo.MarshalUnionNDR(ctx, w, _swOutRevisionInfo); err != nil {
				return err
			}
		} else {
			if err := (&RevisionInfo{}).MarshalUnionNDR(ctx, w, _swOutRevisionInfo); err != nil {
				return err
			}
		}
	}
	// ServerHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Connect5Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// OutVersion {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.OutVersion); err != nil {
			return err
		}
	}
	// OutRevisionInfo {out} (1:{pointer=ref}*(1))(2:{switch_type={}(uint32), alias=SAMPR_REVISION_INFO}(union))
	{
		if o.OutRevisionInfo == nil {
			o.OutRevisionInfo = &RevisionInfo{}
		}
		_swOutRevisionInfo := uint32(o.OutVersion)
		if err := o.OutRevisionInfo.UnmarshalUnionNDR(ctx, w, _swOutRevisionInfo); err != nil {
			return err
		}
	}
	// ServerHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Server == nil {
			o.Server = &Handle{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Connect5Request structure represents the SamrConnect5 operation request
type Connect5Request struct {
	// ServerName: The null-terminated NETBIOS name of the server; this parameter MAY<44>
	// be ignored on receipt.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// DesiredAccess: An ACCESS_MASK that indicates the access requested for ServerHandle
	// on output. For a listing of possible values, see section 2.2.1.3.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
	// InVersion: Indicates which field of the InRevisionInfo union is used.
	InVersion uint32 `idl:"name:InVersion" json:"in_version"`
	// InRevisionInfo: Revision information. For details, see the definition of the SAMPR_REVISION_INFO_V1
	// structure, which is contained in the SAMPR_REVISION_INFO union.
	InRevisionInfo *RevisionInfo `idl:"name:InRevisionInfo;switch_is:InVersion" json:"in_revision_info"`
}

func (o *Connect5Request) xxx_ToOp(ctx context.Context) *xxx_Connect5Operation {
	if o == nil {
		return &xxx_Connect5Operation{}
	}
	return &xxx_Connect5Operation{
		ServerName:     o.ServerName,
		DesiredAccess:  o.DesiredAccess,
		InVersion:      o.InVersion,
		InRevisionInfo: o.InRevisionInfo,
	}
}

func (o *Connect5Request) xxx_FromOp(ctx context.Context, op *xxx_Connect5Operation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
	o.InVersion = op.InVersion
	o.InRevisionInfo = op.InRevisionInfo
}
func (o *Connect5Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Connect5Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Connect5Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Connect5Response structure represents the SamrConnect5 operation response
type Connect5Response struct {
	// OutVersion: Indicates which field of the OutRevisionInfo union is used.
	OutVersion uint32 `idl:"name:OutVersion" json:"out_version"`
	// OutRevisionInfo: Revision information. For details, see the definition of the SAMPR_REVISION_INFO_V1
	// structure, which is contained in the SAMPR_REVISION_INFO union.
	OutRevisionInfo *RevisionInfo `idl:"name:OutRevisionInfo;switch_is:*OutVersion" json:"out_revision_info"`
	// ServerHandle: An RPC context handle, as specified in section 2.2.7.2.
	Server *Handle `idl:"name:ServerHandle" json:"server"`
	// Return: The SamrConnect5 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Connect5Response) xxx_ToOp(ctx context.Context) *xxx_Connect5Operation {
	if o == nil {
		return &xxx_Connect5Operation{}
	}
	return &xxx_Connect5Operation{
		OutVersion:      o.OutVersion,
		OutRevisionInfo: o.OutRevisionInfo,
		Server:          o.Server,
		Return:          o.Return,
	}
}

func (o *Connect5Response) xxx_FromOp(ctx context.Context, op *xxx_Connect5Operation) {
	if o == nil {
		return
	}
	o.OutVersion = op.OutVersion
	o.OutRevisionInfo = op.OutRevisionInfo
	o.Server = op.Server
	o.Return = op.Return
}
func (o *Connect5Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Connect5Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Connect5Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RIDToSIDOperation structure represents the SamrRidToSid operation
type xxx_RIDToSIDOperation struct {
	Object *Handle   `idl:"name:ObjectHandle" json:"object"`
	RID    uint32    `idl:"name:Rid" json:"rid"`
	SID    *dtyp.SID `idl:"name:Sid" json:"sid"`
	Return int32     `idl:"name:Return" json:"return"`
}

func (o *xxx_RIDToSIDOperation) OpNum() int { return 65 }

func (o *xxx_RIDToSIDOperation) OpName() string { return "/samr/v1/SamrRidToSid" }

func (o *xxx_RIDToSIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RIDToSIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ObjectHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
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
	// Rid {in} (1:(uint32))
	{
		if err := w.WriteData(o.RID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RIDToSIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ObjectHandle {in} (1:{context_handle, alias=SAMPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object == nil {
			o.Object = &Handle{}
		}
		if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Rid {in} (1:(uint32))
	{
		if err := w.ReadData(&o.RID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RIDToSIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RIDToSIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Sid {out} (1:{pointer=ref}*(2))(2:{alias=PRPC_SID}*(1))(3:{alias=RPC_SID}(struct))
	{
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RIDToSIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Sid {out} (1:{pointer=ref}*(2))(2:{alias=PRPC_SID,pointer=ref}*(1))(3:{alias=RPC_SID}(struct))
	{
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
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RIDToSIDRequest structure represents the SamrRidToSid operation request
type RIDToSIDRequest struct {
	// ObjectHandle: An RPC context handle, as specified in section 2.2.7.2. The message
	// processing shown later in this section contains details on which types of ObjectHandle
	// are accepted by the server.
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	// Rid: A RID of an account.
	RID uint32 `idl:"name:Rid" json:"rid"`
}

func (o *RIDToSIDRequest) xxx_ToOp(ctx context.Context) *xxx_RIDToSIDOperation {
	if o == nil {
		return &xxx_RIDToSIDOperation{}
	}
	return &xxx_RIDToSIDOperation{
		Object: o.Object,
		RID:    o.RID,
	}
}

func (o *RIDToSIDRequest) xxx_FromOp(ctx context.Context, op *xxx_RIDToSIDOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
	o.RID = op.RID
}
func (o *RIDToSIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RIDToSIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RIDToSIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RIDToSIDResponse structure represents the SamrRidToSid operation response
type RIDToSIDResponse struct {
	// Sid: The SID of the account referenced by Rid.
	//
	// This protocol asks the RPC runtime, via the strict_context_handle attribute, to reject
	// the use of context handles created by a method of a different RPC interface than
	// this one, as specified in [MS-RPCE] section 3.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
	// Return: The SamrRidToSid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RIDToSIDResponse) xxx_ToOp(ctx context.Context) *xxx_RIDToSIDOperation {
	if o == nil {
		return &xxx_RIDToSIDOperation{}
	}
	return &xxx_RIDToSIDOperation{
		SID:    o.SID,
		Return: o.Return,
	}
}

func (o *RIDToSIDResponse) xxx_FromOp(ctx context.Context, op *xxx_RIDToSIDOperation) {
	if o == nil {
		return
	}
	o.SID = op.SID
	o.Return = op.Return
}
func (o *RIDToSIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RIDToSIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RIDToSIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDSRMPasswordOperation structure represents the SamrSetDSRMPassword operation
type xxx_SetDSRMPasswordOperation struct {
	_                      *dtyp.UnicodeString     `idl:"name:Unused;pointer:unique"`
	UserID                 uint32                  `idl:"name:UserId" json:"user_id"`
	EncryptedNTOWFPassword *EncryptedNTOWFPassword `idl:"name:EncryptedNtOwfPassword;pointer:unique" json:"encrypted_nt_owf_password"`
	Return                 int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDSRMPasswordOperation) OpNum() int { return 66 }

func (o *xxx_SetDSRMPasswordOperation) OpName() string { return "/samr/v1/SamrSetDSRMPassword" }

func (o *xxx_SetDSRMPasswordOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDSRMPasswordOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Unused {in} (1:{pointer=unique, alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		// reserved Unused
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// UserId {in} (1:(uint32))
	{
		if err := w.WriteData(o.UserID); err != nil {
			return err
		}
	}
	// EncryptedNtOwfPassword {in} (1:{pointer=unique, alias=PENCRYPTED_NT_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_NT_OWF_PASSWORD}(struct))
	{
		if o.EncryptedNTOWFPassword != nil {
			_ptr_EncryptedNtOwfPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedNTOWFPassword != nil {
					if err := o.EncryptedNTOWFPassword.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedNTOWFPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedNTOWFPassword, _ptr_EncryptedNtOwfPassword); err != nil {
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

func (o *xxx_SetDSRMPasswordOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Unused {in} (1:{pointer=unique, alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		// reserved Unused
		var _Unused *dtyp.UnicodeString
		_ptr_Unused := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if _Unused == nil {
				_Unused = &dtyp.UnicodeString{}
			}
			if err := _Unused.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Unused := func(ptr interface{}) { _Unused = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&_Unused, _s_Unused, _ptr_Unused); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UserId {in} (1:(uint32))
	{
		if err := w.ReadData(&o.UserID); err != nil {
			return err
		}
	}
	// EncryptedNtOwfPassword {in} (1:{pointer=unique, alias=PENCRYPTED_NT_OWF_PASSWORD}*(1))(2:{alias=ENCRYPTED_NT_OWF_PASSWORD}(struct))
	{
		_ptr_EncryptedNtOwfPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EncryptedNTOWFPassword == nil {
				o.EncryptedNTOWFPassword = &EncryptedNTOWFPassword{}
			}
			if err := o.EncryptedNTOWFPassword.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedNtOwfPassword := func(ptr interface{}) { o.EncryptedNTOWFPassword = *ptr.(**EncryptedNTOWFPassword) }
		if err := w.ReadPointer(&o.EncryptedNTOWFPassword, _s_EncryptedNtOwfPassword, _ptr_EncryptedNtOwfPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDSRMPasswordOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDSRMPasswordOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDSRMPasswordOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetDSRMPasswordRequest structure represents the SamrSetDSRMPassword operation request
type SetDSRMPasswordRequest struct {
	// UserId: A RID of a user account. See the message processing later in this section
	// for details on restrictions on this value.
	UserID uint32 `idl:"name:UserId" json:"user_id"`
	// EncryptedNtOwfPassword: The NT hash of the new password (as presented by the client)
	// encrypted according to the specification of ENCRYPTED_NT_OWF_PASSWORD, where the
	// key is the UserId.
	EncryptedNTOWFPassword *EncryptedNTOWFPassword `idl:"name:EncryptedNtOwfPassword;pointer:unique" json:"encrypted_nt_owf_password"`
}

func (o *SetDSRMPasswordRequest) xxx_ToOp(ctx context.Context) *xxx_SetDSRMPasswordOperation {
	if o == nil {
		return &xxx_SetDSRMPasswordOperation{}
	}
	return &xxx_SetDSRMPasswordOperation{
		UserID:                 o.UserID,
		EncryptedNTOWFPassword: o.EncryptedNTOWFPassword,
	}
}

func (o *SetDSRMPasswordRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDSRMPasswordOperation) {
	if o == nil {
		return
	}
	o.UserID = op.UserID
	o.EncryptedNTOWFPassword = op.EncryptedNTOWFPassword
}
func (o *SetDSRMPasswordRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDSRMPasswordRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDSRMPasswordOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDSRMPasswordResponse structure represents the SamrSetDSRMPassword operation response
type SetDSRMPasswordResponse struct {
	// Return: The SamrSetDSRMPassword return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDSRMPasswordResponse) xxx_ToOp(ctx context.Context) *xxx_SetDSRMPasswordOperation {
	if o == nil {
		return &xxx_SetDSRMPasswordOperation{}
	}
	return &xxx_SetDSRMPasswordOperation{
		Return: o.Return,
	}
}

func (o *SetDSRMPasswordResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDSRMPasswordOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetDSRMPasswordResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDSRMPasswordResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDSRMPasswordOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ValidatePasswordOperation structure represents the SamrValidatePassword operation
type xxx_ValidatePasswordOperation struct {
	ValidationType PasswordPolicyValidationType `idl:"name:ValidationType" json:"validation_type"`
	InputArg       *SAMValidateInputArg         `idl:"name:InputArg;switch_is:ValidationType" json:"input_arg"`
	OutputArg      *SAMValidateOutputArg        `idl:"name:OutputArg;switch_is:ValidationType" json:"output_arg"`
	Return         int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_ValidatePasswordOperation) OpNum() int { return 67 }

func (o *xxx_ValidatePasswordOperation) OpName() string { return "/samr/v1/SamrValidatePassword" }

func (o *xxx_ValidatePasswordOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidatePasswordOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ValidationType {in} (1:{alias=PASSWORD_POLICY_VALIDATION_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.ValidationType)); err != nil {
			return err
		}
	}
	// InputArg {in} (1:{switch_type={alias=PASSWORD_POLICY_VALIDATION_TYPE}(enum), alias=PSAM_VALIDATE_INPUT_ARG}*(1))(2:{switch_type={alias=PASSWORD_POLICY_VALIDATION_TYPE}(enum), alias=SAM_VALIDATE_INPUT_ARG}(union))
	{
		_swInputArg := uint16(o.ValidationType)
		if o.InputArg != nil {
			if err := o.InputArg.MarshalUnionNDR(ctx, w, _swInputArg); err != nil {
				return err
			}
		} else {
			if err := (&SAMValidateInputArg{}).MarshalUnionNDR(ctx, w, _swInputArg); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidatePasswordOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ValidationType {in} (1:{alias=PASSWORD_POLICY_VALIDATION_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ValidationType)); err != nil {
			return err
		}
	}
	// InputArg {in} (1:{switch_type={alias=PASSWORD_POLICY_VALIDATION_TYPE}(enum), alias=PSAM_VALIDATE_INPUT_ARG,pointer=ref}*(1))(2:{switch_type={alias=PASSWORD_POLICY_VALIDATION_TYPE}(enum), alias=SAM_VALIDATE_INPUT_ARG}(union))
	{
		if o.InputArg == nil {
			o.InputArg = &SAMValidateInputArg{}
		}
		_swInputArg := uint16(o.ValidationType)
		if err := o.InputArg.UnmarshalUnionNDR(ctx, w, _swInputArg); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidatePasswordOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidatePasswordOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// OutputArg {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=PASSWORD_POLICY_VALIDATION_TYPE}(enum), alias=PSAM_VALIDATE_OUTPUT_ARG}*(1))(3:{switch_type={alias=PASSWORD_POLICY_VALIDATION_TYPE}(enum), alias=SAM_VALIDATE_OUTPUT_ARG}(union))
	{
		if o.OutputArg != nil {
			_ptr_OutputArg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swOutputArg := uint16(o.ValidationType)
				if o.OutputArg != nil {
					if err := o.OutputArg.MarshalUnionNDR(ctx, w, _swOutputArg); err != nil {
						return err
					}
				} else {
					if err := (&SAMValidateOutputArg{}).MarshalUnionNDR(ctx, w, _swOutputArg); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OutputArg, _ptr_OutputArg); err != nil {
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
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidatePasswordOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// OutputArg {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=PASSWORD_POLICY_VALIDATION_TYPE}(enum), alias=PSAM_VALIDATE_OUTPUT_ARG,pointer=ref}*(1))(3:{switch_type={alias=PASSWORD_POLICY_VALIDATION_TYPE}(enum), alias=SAM_VALIDATE_OUTPUT_ARG}(union))
	{
		_ptr_OutputArg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OutputArg == nil {
				o.OutputArg = &SAMValidateOutputArg{}
			}
			_swOutputArg := uint16(o.ValidationType)
			if err := o.OutputArg.UnmarshalUnionNDR(ctx, w, _swOutputArg); err != nil {
				return err
			}
			return nil
		})
		_s_OutputArg := func(ptr interface{}) { o.OutputArg = *ptr.(**SAMValidateOutputArg) }
		if err := w.ReadPointer(&o.OutputArg, _s_OutputArg, _ptr_OutputArg); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ValidatePasswordRequest structure represents the SamrValidatePassword operation request
type ValidatePasswordRequest struct {
	// ValidationType: The password policy validation requested.
	ValidationType PasswordPolicyValidationType `idl:"name:ValidationType" json:"validation_type"`
	// InputArg: The password-related material to validate.
	InputArg *SAMValidateInputArg `idl:"name:InputArg;switch_is:ValidationType" json:"input_arg"`
}

func (o *ValidatePasswordRequest) xxx_ToOp(ctx context.Context) *xxx_ValidatePasswordOperation {
	if o == nil {
		return &xxx_ValidatePasswordOperation{}
	}
	return &xxx_ValidatePasswordOperation{
		ValidationType: o.ValidationType,
		InputArg:       o.InputArg,
	}
}

func (o *ValidatePasswordRequest) xxx_FromOp(ctx context.Context, op *xxx_ValidatePasswordOperation) {
	if o == nil {
		return
	}
	o.ValidationType = op.ValidationType
	o.InputArg = op.InputArg
}
func (o *ValidatePasswordRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ValidatePasswordRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ValidatePasswordOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ValidatePasswordResponse structure represents the SamrValidatePassword operation response
type ValidatePasswordResponse struct {
	// OutputArg: The result of the validation.
	OutputArg *SAMValidateOutputArg `idl:"name:OutputArg;switch_is:ValidationType" json:"output_arg"`
	// Return: The SamrValidatePassword return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ValidatePasswordResponse) xxx_ToOp(ctx context.Context) *xxx_ValidatePasswordOperation {
	if o == nil {
		return &xxx_ValidatePasswordOperation{}
	}
	return &xxx_ValidatePasswordOperation{
		OutputArg: o.OutputArg,
		Return:    o.Return,
	}
}

func (o *ValidatePasswordResponse) xxx_FromOp(ctx context.Context, op *xxx_ValidatePasswordOperation) {
	if o == nil {
		return
	}
	o.OutputArg = op.OutputArg
	o.Return = op.Return
}
func (o *ValidatePasswordResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ValidatePasswordResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ValidatePasswordOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
