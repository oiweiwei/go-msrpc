// The pac package implements the PAC client protocol.
//
// # Introduction
//
// The Privilege Attribute Certificate (PAC) Data Structure is used by authentication
// protocols that verify identities to transport authorization information, which controls
// access to resources. The Kerberos protocol [RFC4120] does not provide authorization.
// The Privilege Attribute Certificate (PAC) was created to provide this authorization
// data for Kerberos Protocol Extensions [MS-KILE]. Into the PAC structure [MS-KILE]
// encodes authorization information, which consists of group memberships, additional
// credential information, profile and policy information, and supporting security metadata.<1>
//
// # Overview
//
// The PAC is a structure that conveys authorization-related information provided by
// domain controllers (DCs). The PAC is used by authentication protocols that verify
// identities to transport authorization information, which controls access to resources.
// Once authentication has been accomplished, the next task is to decide if a particular
// request is authorized. Management of network systems often models broad authorization
// decisions through groups; for example, all engineers who can access a specific printer
// or all sales personnel who can access a certain web server. Making group information
// consistently available to several services allows for simpler management.
//
// The Kerberos protocol is one of the most commonly used authentication mechanisms.
// However, the Kerberos protocol [RFC4120] does not provide authorization; "kerberized"
// applications are expected to manage their own authorization, typically through names.
// Specifically, the Kerberos protocol does not define any explicit group membership
// or logon policy information to be carried in the Kerberos tickets. It leaves that
// for Kerberos extensions to provide a mechanism to convey authorization information
// by encapsulating this information within an AuthorizationData structure ([RFC4120]
// section 5.2.6). The PAC was created to provide this authorization data for Kerberos
// Protocol Extensions [MS-KILE].
//
// [MS-KILE] requires that the PAC information be encoded within an AuthorizationData
// element ([RFC4120] section 5.2.6) which consists of group memberships, additional
// credential information, profile and policy information, and supporting security metadata.
// [MS-KILE] also requires that the PAC information be enclosed in an AD-IF-RELEVANT
// AuthorizationData element, since this information is noncritical authorization data.
// This clearly indicates to the receiver that this data can be ignored if the receiver
// does consume the information in the PAC.
//
// Examples of information that can be provided by a DC include:
//
// * Authorization data such as security identifiers (SIDs) ( f2ef15b6-1e9b-48b5-bf0b-019f061d41c8#gt_83f2020d-0804-4840-a5ac-e06439d50f8d
// ) and relative identifiers (RIDs) ( f2ef15b6-1e9b-48b5-bf0b-019f061d41c8#gt_df3d0b61-56cd-4dac-9402-982f1fedc41c
// ).
//
// * User profile information such as a home directory or logon script.
//
// * Password credentials, used during smart card authentication, for password based
// authentication protocols to use at a later time.
//
// * Service for User (S4U) ( f2ef15b6-1e9b-48b5-bf0b-019f061d41c8#gt_083a5403-f654-4db6-b17e-9c10dc5cd420
// ) protocol [MS-SFU] ( ../ms-sfu/3bff5864-8135-400e-bdd9-33b552051d94 ) data.
package pac

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	claims "github.com/oiweiwei/go-msrpc/msrpc/adts/claims/claims/v1"
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
	_ = dtyp.GoPackage
	_ = claims.GoPackage
)

var (
	// import guard
	GoPackage = "pac"
)

// PACInfoBufferTypeLogonInfo represents the PAC_INFO_BUFFER_TYPE_LOGON_INFO RPC constant
const PACInfoBufferTypeLogonInfo = 0x00000001

// PACInfoBufferTypeCredentialsInfo represents the PAC_INFO_BUFFER_TYPE_CREDENTIALS_INFO RPC constant
const PACInfoBufferTypeCredentialsInfo = 0x00000002

// PACInfoBufferTypeServerChecksum represents the PAC_INFO_BUFFER_TYPE_SERVER_CHECKSUM RPC constant
const PACInfoBufferTypeServerChecksum = 0x00000006

// PACInfoBufferTypeKDCChecksum represents the PAC_INFO_BUFFER_TYPE_KDC_CHECKSUM RPC constant
const PACInfoBufferTypeKDCChecksum = 0x00000007

// PACInfoBufferTypeClientNameAndTicketInfo represents the PAC_INFO_BUFFER_TYPE_CLIENT_NAME_AND_TICKET_INFO RPC constant
const PACInfoBufferTypeClientNameAndTicketInfo = 0x0000000A

// PACInfoBufferTypeConstrainedDelegationInfo represents the PAC_INFO_BUFFER_TYPE_CONSTRAINED_DELEGATION_INFO RPC constant
const PACInfoBufferTypeConstrainedDelegationInfo = 0x0000000B

// PACInfoBufferTypeUPNAndDNSInfo represents the PAC_INFO_BUFFER_TYPE_UPN_AND_DNS_INFO RPC constant
const PACInfoBufferTypeUPNAndDNSInfo = 0x0000000C

// PACInfoBufferTypeClientClaimsInfo represents the PAC_INFO_BUFFER_TYPE_CLIENT_CLAIMS_INFO RPC constant
const PACInfoBufferTypeClientClaimsInfo = 0x0000000D

// PACInfoBufferTypeDeviceInfo represents the PAC_INFO_BUFFER_TYPE_DEVICE_INFO RPC constant
const PACInfoBufferTypeDeviceInfo = 0x0000000E

// PACInfoBufferTypeDeviceClaimsInfo represents the PAC_INFO_BUFFER_TYPE_DEVICE_CLAIMS_INFO RPC constant
const PACInfoBufferTypeDeviceClaimsInfo = 0x0000000F

// PACInfoBufferTypeTicketChecksum represents the PAC_INFO_BUFFER_TYPE_TICKET_CHECKSUM RPC constant
const PACInfoBufferTypeTicketChecksum = 0x00000010

// PACInfoBufferTypeAttributes represents the PAC_INFO_BUFFER_TYPE_ATTRIBUTES RPC constant
const PACInfoBufferTypeAttributes = 0x00000011

// PACInfoBufferTypeRequestorSID represents the PAC_INFO_BUFFER_TYPE_REQUESTOR_SID RPC constant
const PACInfoBufferTypeRequestorSID = 0x00000012

// PACInfoBufferTypeExtendedKDCChecksum represents the PAC_INFO_BUFFER_TYPE_EXTENDED_KDC_CHECKSUM RPC constant
const PACInfoBufferTypeExtendedKDCChecksum = 0x00000013

// PACInfoBufferTypeRequestorGUID represents the PAC_INFO_BUFFER_TYPE_REQUESTOR_GUID RPC constant
const PACInfoBufferTypeRequestorGUID = 0x00000014

// EncryptionTypeNone represents the ENCRYPTION_TYPE_NONE RPC constant
const EncryptionTypeNone = 0x00000000

// EncryptionTypeDESCBCCRC represents the ENCRYPTION_TYPE_DES_CBC_CRC RPC constant
const EncryptionTypeDESCBCCRC = 0x00000001

// EncryptionTypeDESCBCMD5 represents the ENCRYPTION_TYPE_DES_CBC_MD5 RPC constant
const EncryptionTypeDESCBCMD5 = 0x00000003

// EncryptionTypeAES128CTSHMACSHA196 represents the ENCRYPTION_TYPE_AES128_CTS_HMAC_SHA1_96 RPC constant
const EncryptionTypeAES128CTSHMACSHA196 = 0x00000011

// EncryptionTypeAES256CTSHMACSHA196 represents the ENCRYPTION_TYPE_AES256_CTS_HMAC_SHA1_96 RPC constant
const EncryptionTypeAES256CTSHMACSHA196 = 0x00000012

// EncryptionTypeRC4HMAC represents the ENCRYPTION_TYPE_RC4_HMAC RPC constant
const EncryptionTypeRC4HMAC = 0x00000017

// SignatureTypeKerberosChecksumHMACMD5 represents the SIGNATURE_TYPE_KERB_CHECKSUM_HMAC_MD5 RPC constant
const SignatureTypeKerberosChecksumHMACMD5 = 0xFFFFFF76

// SignatureTypeHMACSHA196AES128 represents the SIGNATURE_TYPE_HMAC_SHA1_96_AES128 RPC constant
const SignatureTypeHMACSHA196AES128 = 0x0000000F

// SignatureTypeHMACSHA196AES256 represents the SIGNATURE_TYPE_HMAC_SHA1_96_AES256 RPC constant
const SignatureTypeHMACSHA196AES256 = 0x00000010

// PACWasRequested represents the PAC_WAS_REQUESTED RPC constant
const PACWasRequested = 0x00000001

// PACWasGivenImplicitly represents the PAC_WAS_GIVEN_IMPLICITLY RPC constant
const PACWasGivenImplicitly = 0x00000002

// KerberosSIDAndAttributes structure represents KERB_SID_AND_ATTRIBUTES RPC structure.
//
// The KERB_SID_AND_ATTRIBUTES structure represents a SID and its attributes for use
// in authentication. It is sent within the KERB_VALIDATION_INFO (section 2.5) structure
// and used to include additional information about the group that the SID references.
//
// The KERB_SID_AND_ATTRIBUTES structure is defined as follows.
type KerberosSIDAndAttributes struct {
	// Sid: A pointer to an RPC_SID structure ([MS-DTYP] section 2.4.2.3).
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
	// Attributes: A set of bit flags that describe attributes of the SID in the Sid field.
	//
	// Attributes can contain one or more of the following bits.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | E | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | D | C | B | A |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// Where the bits are defined as:
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                   DESCRIPTION                                    |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	| A     | This setting means that the group is mandatory for the user and cannot be        |
	//	|       | disabled. Corresponds to SE_GROUP_MANDATORY. For more information, see [SIDATT]. |
	//	+-------+----------------------------------------------------------------------------------+
	//	| B     | This setting means that the group is marked as enabled by default. Corresponds   |
	//	|       | to SE_GROUP_ENABLED_BY_DEFAULT. For more information, see [SIDATT].              |
	//	+-------+----------------------------------------------------------------------------------+
	//	| C     | This setting means that the group is enabled for use. Corresponds to             |
	//	|       | SE_GROUP_ENABLED. For more information, see [SIDATT].                            |
	//	+-------+----------------------------------------------------------------------------------+
	//	| D     | This setting means that the group can be assigned as an owner of a resource.     |
	//	|       | Corresponds to SE_GROUP_OWNER. For more information, see [SIDATT].               |
	//	+-------+----------------------------------------------------------------------------------+
	//	| E     | This setting means that the group is a domain-local or resource group.           |
	//	|       | Corresponds to SE_GROUP_RESOURCE. For more information, see [SIDATT].            |
	//	+-------+----------------------------------------------------------------------------------+
	Attributes uint32 `idl:"name:Attributes" json:"attributes"`
}

func (o *KerberosSIDAndAttributes) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosSIDAndAttributes) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
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
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	return nil
}
func (o *KerberosSIDAndAttributes) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
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
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	return nil
}

// GroupMembership structure represents GROUP_MEMBERSHIP RPC structure.
//
// The GROUP_MEMBERSHIP structure identifies a group to which an account belongs. It
// is sent within the KERB_VALIDATION_INFO (section 2.5) structure.
//
// The GROUP_MEMBERSHIP structure is defined as follows.
type GroupMembership struct {
	// RelativeId: A 32-bit unsigned integer that contains the RID of a particular group.
	RelativeID uint32 `idl:"name:RelativeId" json:"relative_id"`
	// Attributes: A 32-bit unsigned integer value that contains the group membership attributes
	// set for the RID contained in RelativeId. The possible values for the Attributes flags
	// are identical to those specified in KERB_SID_AND_ATTRIBUTES (section 2.2.1).
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

// DomainGroupMembership structure represents DOMAIN_GROUP_MEMBERSHIP RPC structure.
//
// The DOMAIN_GROUP_MEMBERSHIP structure identifies a domain and groups to which an
// account belongs. It is sent within the PAC_DEVICE_INFO (section 2.12) structure.<2>
//
// The DOMAIN_GROUP_MEMBERSHIP structure is defined as follows.
type DomainGroupMembership struct {
	// DomainId: A SID structure that contains the SID for the domain. This member is used
	// in conjunction with the GroupIds members to create group SIDs for the device.
	DomainID *dtyp.SID `idl:"name:DomainId" json:"domain_id"`
	// GroupCount: A 32-bit unsigned integer that contains the number of groups within the
	// domain to which the account belongs.
	GroupCount uint32 `idl:"name:GroupCount" json:"group_count"`
	// GroupIds: A pointer to a list of GROUP_MEMBERSHIP structures that contain the groups
	// to which the account belongs in the domain. The number of groups in this list MUST
	// be equal to GroupCount.
	GroupIDs []*GroupMembership `idl:"name:GroupIds;size_is:(GroupCount)" json:"group_ids"`
}

func (o *DomainGroupMembership) xxx_PreparePayload(ctx context.Context) error {
	if o.GroupIDs != nil && o.GroupCount == 0 {
		o.GroupCount = uint32(len(o.GroupIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainGroupMembership) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
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
	if err := w.WriteData(o.GroupCount); err != nil {
		return err
	}
	if o.GroupIDs != nil || o.GroupCount > 0 {
		_ptr_GroupIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.GroupCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.GroupIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.GroupIDs[i1] != nil {
					if err := o.GroupIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.GroupIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GroupIDs, _ptr_GroupIds); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainGroupMembership) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
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
	if err := w.ReadData(&o.GroupCount); err != nil {
		return err
	}
	_ptr_GroupIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.GroupCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.GroupCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.GroupIDs", sizeInfo[0])
		}
		o.GroupIDs = make([]*GroupMembership, sizeInfo[0])
		for i1 := range o.GroupIDs {
			i1 := i1
			if o.GroupIDs[i1] == nil {
				o.GroupIDs[i1] = &GroupMembership{}
			}
			if err := o.GroupIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_GroupIds := func(ptr interface{}) { o.GroupIDs = *ptr.(*[]*GroupMembership) }
	if err := w.ReadPointer(&o.GroupIDs, _s_GroupIds, _ptr_GroupIds); err != nil {
		return err
	}
	return nil
}

// PACType structure represents PACTYPE RPC structure.
//
// The PACTYPE structure is the topmost structure of the PAC and specifies the number
// of elements in the PAC_INFO_BUFFER (section 2.4) array. The PACTYPE structure serves
// as the header for the complete PAC data.
//
// The PACTYPE structure is defined as follows.
//
// The format of the PACTYPE structure is defined as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cBuffers                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Version                                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Buffers (variable)                                                                                                            |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type PACType struct {
	// cBuffers (4 bytes): A 32-bit unsigned integer in little-endian format that defines
	// the number of entries in the Buffers array.
	BuffersCount uint32 `idl:"name:cBuffers" json:"buffers_count"`
	// Version (4 bytes): A 32-bit unsigned integer in little-endian format that defines
	// the PAC version; MUST be 0x00000000.
	Version uint32 `idl:"name:Version" json:"version"`
	// Buffers (variable): An array of PAC_INFO_BUFFER structures (section 2.4).
	//
	// The actual contents of the PAC are placed serially after the variable set of PAC_INFO_BUFFER
	// structures. The contents are individually serialized PAC elements. All PAC elements
	// MUST be placed on an 8-byte boundary.
	Buffers []*PACInfoBuffer `idl:"name:Buffers;size_is:(cBuffers)" json:"buffers"`
}

func (o *PACType) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffers != nil && o.BuffersCount == 0 {
		o.BuffersCount = uint32(len(o.Buffers))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.BuffersCount); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if o.Buffers != nil || o.BuffersCount > 0 {
		_ptr_Buffers := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BuffersCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffers {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Buffers[i1] != nil {
					if err := o.Buffers[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PACInfoBuffer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffers); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&PACInfoBuffer{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffers, _ptr_Buffers); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.BuffersCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	_ptr_Buffers := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BuffersCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BuffersCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffers", sizeInfo[0])
		}
		o.Buffers = make([]*PACInfoBuffer, sizeInfo[0])
		for i1 := range o.Buffers {
			i1 := i1
			if o.Buffers[i1] == nil {
				o.Buffers[i1] = &PACInfoBuffer{}
			}
			if err := o.Buffers[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffers := func(ptr interface{}) { o.Buffers = *ptr.(*[]*PACInfoBuffer) }
	if err := w.ReadPointer(&o.Buffers, _s_Buffers, _ptr_Buffers); err != nil {
		return err
	}
	return nil
}

// PACInfoBuffer structure represents PAC_INFO_BUFFER RPC structure.
//
// Following the PACTYPE (section 2.3) structure is an array of PAC_INFO_BUFFER structures
// each of which defines the type and byte offset to a buffer of the PAC. The PAC_INFO_BUFFER
// array has no defined ordering. Therefore, the order of the PAC_INFO_BUFFER buffers
// has no significance. However, once the Key Distribution Center (KDC) and server signatures
// are generated, the ordering of the buffers MUST NOT change, or signature verification
// of the PAC contents will fail.
//
// The PAC_INFO_BUFFER structure is defined as follows.
//
// The format of the PAC_INFO_BUFFER structure is defined as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ulType                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cbBufferSize                                                                                                                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Offset                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type PACInfoBuffer struct {
	// ulType (4 bytes): A 32-bit unsigned integer in little-endian format that describes
	// the type of data present in the buffer contained at Offset. Types that are not understood
	// MUST be ignored.
	//
	//	+-----------------+----------------------------------------------------------------------------------+
	//	|                 |                                                                                  |
	//	|      VALUE      |                                     MEANING                                      |
	//	|                 |                                                                                  |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 (1)  | Logon information (section 2.5). PAC structures MUST contain one buffer of this  |
	//	|                 | type. Additional logon information buffers MUST be ignored.                      |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 (2)  | Credentials information (section 2.6). PAC structures SHOULD NOT contain more    |
	//	|                 | than one buffer of this type, based on constraints specified in section 2.6.     |
	//	|                 | Second or subsequent credentials information buffers MUST be ignored on receipt. |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 (6)  | Server checksum (section 2.8). PAC structures MUST contain one buffer of this    |
	//	|                 | type for Kerberos ticket-granting service (TGS) requests or Kerberos application |
	//	|                 | protocol (AP) requests, and none otherwise. Additional logon server checksum     |
	//	|                 | buffers MUST be ignored.                                                         |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000007 (7)  | KDC (privilege server) checksum (section 2.8). PAC structures MUST contain       |
	//	|                 | one buffer of this type for Kerberos ticket-granting service (TGS) requests or   |
	//	|                 | Kerberos application protocol (AP) requests, and none otherwise. Additional KDC  |
	//	|                 | checksum buffers MUST be ignored.                                                |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x0000000A (10) | Client name and ticket information (section 2.7). PAC structures MUST contain    |
	//	|                 | one buffer of this type. Additional client and ticket information buffers MUST   |
	//	|                 | be ignored.                                                                      |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x0000000B (11) | Constrained delegation information (section 2.9). PAC structures MUST contain    |
	//	|                 | one buffer of this type for Service for User to Proxy (S4U2proxy) [MS-SFU]       |
	//	|                 | requests and none otherwise. Additional constrained delegation information       |
	//	|                 | buffers MUST be ignored.                                                         |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x0000000C (12) | User principal name (UPN) and Domain Name System (DNS) information (section      |
	//	|                 | 2.10). PAC structures SHOULD NOT<3> contain more than one buffer of this type.   |
	//	|                 | Second or subsequent UPN and DNS information buffers MUST be ignored on receipt. |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x0000000D (13) | Client claims information (section 2.11). PAC structures SHOULD NOT<4> contain   |
	//	|                 | more than one buffer of this type. Additional client claims information buffers  |
	//	|                 | MUST be ignored.                                                                 |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x0000000E (14) | Device information (section 2.12). PAC structures SHOULD NOT<5> contain more     |
	//	|                 | than one buffer of this type. Additional device information buffers MUST be      |
	//	|                 | ignored.                                                                         |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x0000000F (15) | Device claims information (section 2.13). PAC structures SHOULD NOT<6> contain   |
	//	|                 | more than one buffer of this type. Additional device claims information buffers  |
	//	|                 | MUST be ignored.                                                                 |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000010 (16) | Ticket checksum (section 2.8). PAC structures MUST contain one buffer of this    |
	//	|                 | type for Kerberos ticket-granting service (TGS) requests, and none otherwise.    |
	//	|                 | Additional ticket checksum buffers MUST be ignored.<7>                           |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000011 (17) | PAC Attributes indicates that the buffer contains attribute bits for the PAC     |
	//	|                 | (section 2.14). PAC structures SHOULD NOT contain more than one buffer of this   |
	//	|                 | type. Additional attribute buffers MUST be ignored.<8>                           |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000012 (18) | PAC Requestor indicates that the buffer contains the SID of principal that       |
	//	|                 | requested the PAC (section 2.15). PAC structures MUST contain one buffer of this |
	//	|                 | type.<9>                                                                         |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 (19) | Extended KDC (privilege server) checksum (section 2.8). PAC structures MUST      |
	//	|                 | contain one buffer of this type for Kerberos ticket-granting service (TGS)       |
	//	|                 | requests, and none otherwise. Additional Extended KDC checksum buffers MUST be   |
	//	|                 | ignored.<10>                                                                     |
	//	+-----------------+----------------------------------------------------------------------------------+
	Type uint32 `idl:"name:ulType" json:"type"`
	// cbBufferSize (4 bytes): A 32-bit unsigned integer in little-endian format that contains
	// the size, in bytes, of the buffer in the PAC located at Offset.
	BufferLength uint32 `idl:"name:cbBufferSize" json:"buffer_length"`
	// Offset (8 bytes): A 64-bit unsigned integer in little-endian format that contains
	// the offset to the beginning of the buffer, in bytes, from the beginning of the PACTYPE
	// structure (section 2.3). The data offset MUST be a multiple of eight. The following
	// sections specify the format of each type of element.
	Offset uint64 `idl:"name:Offset" json:"offset"`
}

func (o *PACInfoBuffer) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACInfoBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Offset); err != nil {
		return err
	}
	return nil
}
func (o *PACInfoBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offset); err != nil {
		return err
	}
	return nil
}

// CypherBlock structure represents CYPHER_BLOCK RPC structure.
type CypherBlock struct {
	Data []byte `idl:"name:data" json:"data"`
}

func (o *CypherBlock) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CypherBlock) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *CypherBlock) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Data = make([]byte, 8)
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// UserSessionKey structure represents USER_SESSION_KEY RPC structure.
type UserSessionKey struct {
	Data []*CypherBlock `idl:"name:data" json:"data"`
}

func (o *UserSessionKey) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserSessionKey) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= 2 {
			break
		}
		if o.Data[i1] != nil {
			if err := o.Data[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&CypherBlock{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Data); uint64(i1) < 2; i1++ {
		if err := (&CypherBlock{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserSessionKey) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Data = make([]*CypherBlock, 2)
	for i1 := range o.Data {
		i1 := i1
		if o.Data[i1] == nil {
			o.Data[i1] = &CypherBlock{}
		}
		if err := o.Data[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// KerberosValidationInfo structure represents KERB_VALIDATION_INFO RPC structure.
//
// The KERB_VALIDATION_INFO structure defines the user's logon and authorization information
// provided by the DC. A pointer to the KERB_VALIDATION_INFO structure is serialized
// into an array of bytes and then placed after the Buffers array of the topmost PACTYPE
// structure (section 2.3), at the offset specified in the Offset field of the corresponding
// PAC_INFO_BUFFER structure (section 2.4) in the Buffers array. The ulType field of
// the corresponding PAC_INFO_BUFFER structure is set to 0x00000001.
//
// The KERB_VALIDATION_INFO structure is a subset of the NETLOGON_VALIDATION_SAM_INFO4
// structure ([MS-NRPC] section 2.2.1.4.13). It is a subset due to historical reasons
// and to the use of Active Directory to generate this information. NTLM uses the NETLOGON_VALIDATION_SAM_INFO4
// structure in the context of the server to domain controller exchange, as defined
// in [MS-APDS] section 3.1. Consequently, the KERB_VALIDATION_INFO structure includes
// NTLM-specific fields. Fields that are common to the KERB_VALIDATION_INFO and the
// NETLOGON_VALIDATION_SAM_INFO4 structures, and which are specific to the NTLM authentication
// operation, are not used with [MS-KILE] authentication. The KERB_VALIDATION_INFO structure
// is marshaled by RPC [MS-RPCE].
//
// The KERB_VALIDATION_INFO structure is defined as follows.
type KerberosValidationInfo struct {
	// LogonTime: A FILETIME structure that contains the user account's lastLogon attribute
	// ([MS-ADA1] section 2.351) value.
	LogonTime *dtyp.Filetime `idl:"name:LogonTime" json:"logon_time"`
	// LogoffTime: A FILETIME structure that contains the time the client's logon session
	// is set to expire. If the session is set not to expire, the dwHighDateTime member
	// is set to 0x7FFFFFFF and the dwLowDateTime member set to 0xFFFFFFFF. A recipient
	// of the PAC SHOULD<11> use this value as an indicator of when to warn the user that
	// the allowed time is due to expire.
	LogoffTime *dtyp.Filetime `idl:"name:LogoffTime" json:"logoff_time"`
	// KickOffTime: A FILETIME structure that contains LogoffTime minus the user account's
	// forceLogoff attribute ([MS-ADA1] section 2.233) value. If the client is not to be
	// forcibly logged off, the dwHighDateTime member is set to 0x7FFFFFFF and the dwLowDateTime
	// member set to 0xFFFFFFFF. The Kerberos service ticket end time is a replacement for
	// KickOffTime. The service ticket lifetime SHOULD NOT<12> be set longer than the KickOffTime
	// of an account. A recipient of the PAC uses this value as the indicator of when the
	// client is to be forcibly disconnected.
	KickOffTime *dtyp.Filetime `idl:"name:KickOffTime" json:"kick_off_time"`
	// PasswordLastSet:  A FILETIME structure that contains the user account's pwdLastSet
	// attribute ([MS-ADA3] section 2.175) value. If the password was never set, this structure
	// MUST have the dwHighDateTime member set to 0x00000000 and the dwLowDateTime member
	// set to 0x00000000.
	PasswordLastSet *dtyp.Filetime `idl:"name:PasswordLastSet" json:"password_last_set"`
	// PasswordCanChange: A FILETIME structure that contains the time at which the client's
	// password is allowed to change. If there is no restriction on when the client can
	// change the password, this member MUST be set to zero.
	PasswordCanChange *dtyp.Filetime `idl:"name:PasswordCanChange" json:"password_can_change"`
	// PasswordMustChange: A FILETIME structure that contains the time at which the client's
	// password expires. If the password will not expire, this structure MUST have the dwHighDateTime
	// member set to 0x7FFFFFFF and the dwLowDateTime member set to 0xFFFFFFFF.
	PasswordMustChange *dtyp.Filetime `idl:"name:PasswordMustChange" json:"password_must_change"`
	// EffectiveName: An RPC_UNICODE_STRING structure that contains the user account's samAccountName
	// attribute ([MS-ADA3] section 2.222) value.
	EffectiveName *dtyp.UnicodeString `idl:"name:EffectiveName" json:"effective_name"`
	// FullName: An RPC_UNICODE_STRING structure that contains the user account's full name
	// for interactive logon and is set to zero for network logon. If FullName is omitted,
	// this member MUST contain an RPC_UNICODE_STRING structure with the Length member set
	// to zero.
	FullName *dtyp.UnicodeString `idl:"name:FullName" json:"full_name"`
	// LogonScript: An RPC_UNICODE_STRING structure that contains the user account's scriptPath
	// attribute ([MS-ADA3] section 2.232) value for interactive logon and is set to zero
	// for network logon. If no LogonScript is configured for the user, this member MUST
	// contain an RPC_UNICODE_STRING structure with the Length member set to zero.
	LogonScript *dtyp.UnicodeString `idl:"name:LogonScript" json:"logon_script"`
	// ProfilePath: An RPC_UNICODE_STRING structure that contains the user account's profilePath
	// attribute ([MS-ADA3] section 2.167) value for interactive logon and is set to zero
	// for network logon. If no ProfilePath is configured for the user, this member MUST
	// contain an RPC_UNICODE_STRING structure with the Length member set to zero.
	ProfilePath *dtyp.UnicodeString `idl:"name:ProfilePath" json:"profile_path"`
	// HomeDirectory: An RPC_UNICODE_STRING structure that contains the user account's HomeDirectory
	// attribute ([MS-ADA1] section 2.295) value for interactive logon and is set to zero
	// for network logon. If no HomeDirectory is configured for the user, this member MUST
	// contain an RPC_UNICODE_STRING structure with the Length member set to zero.
	HomeDirectory *dtyp.UnicodeString `idl:"name:HomeDirectory" json:"home_directory"`
	// HomeDirectoryDrive: An RPC_UNICODE_STRING structure that contains the user account's
	// HomeDrive attribute ([MS-ADA1] section 2.296) value for interactive logon and is
	// set to zero for network logon. This member MUST be populated if HomeDirectory contains
	// a UNC path. If no HomeDirectoryDrive is configured for the user, this member MUST
	// contain an RPC_UNICODE_STRING structure with the Length member set to zero.
	HomeDirectoryDrive *dtyp.UnicodeString `idl:"name:HomeDirectoryDrive" json:"home_directory_drive"`
	// LogonCount: A 16-bit unsigned integer that contains the user account's LogonCount
	// attribute ([MS-ADA1] section 2.375) value.
	LogonCount uint16 `idl:"name:LogonCount" json:"logon_count"`
	// BadPasswordCount: A 16-bit unsigned integer that contains the user account's badPwdCount
	// attribute ([MS-ADA1] section 2.83) value for interactive logon and is set to zero
	// for network logon.
	BadPasswordCount uint16 `idl:"name:BadPasswordCount" json:"bad_password_count"`
	// UserId: A 32-bit unsigned integer that contains the RID of the account. If the UserId
	// member equals 0x00000000, the first group SID in this member is the SID for this
	// account.
	UserID uint32 `idl:"name:UserId" json:"user_id"`
	// PrimaryGroupId: A 32-bit unsigned integer that contains the RID for the primary group
	// to which this account belongs.
	PrimaryGroupID uint32 `idl:"name:PrimaryGroupId" json:"primary_group_id"`
	// GroupCount: A 32-bit unsigned integer that contains the number of groups within the
	// account domain to which the account belongs.
	GroupCount uint32 `idl:"name:GroupCount" json:"group_count"`
	// GroupIds: A pointer to a list of GROUP_MEMBERSHIP (section 2.2.2) structures that
	// contains the groups to which the account belongs in the account domain. The number
	// of groups in this list MUST be equal to GroupCount.
	GroupIDs []*GroupMembership `idl:"name:GroupIds;size_is:(GroupCount)" json:"group_ids"`
	// UserFlags: A 32-bit unsigned integer that contains a set of bit flags that describe
	// the user's logon information.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | L | K | J | I | H | G | F | E | D | 0 | C | 0 | B | A |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// The following flags are set only when this structure is created as the result of
	// an NTLM authentication, as specified in [MS-NLMP]. These flags MUST be zero for any
	// other authentication protocol, such as [MS-KILE] authentication.
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                   DESCRIPTION                                    |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	| A     | Authentication was done via the GUEST account; no password was used.             |
	//	+-------+----------------------------------------------------------------------------------+
	//	| B     | No encryption is available.                                                      |
	//	+-------+----------------------------------------------------------------------------------+
	//	| C     | LAN Manager key was used for authentication.                                     |
	//	+-------+----------------------------------------------------------------------------------+
	//	| E     | Sub-authentication used; session key came from the sub-authentication package.   |
	//	+-------+----------------------------------------------------------------------------------+
	//	| F     | Indicates that the account is a machine account.                                 |
	//	+-------+----------------------------------------------------------------------------------+
	//	| G     | Indicates that the domain controller understands NTLMv2.                         |
	//	+-------+----------------------------------------------------------------------------------+
	//	| I     | Indicates that ProfilePath is populated.                                         |
	//	+-------+----------------------------------------------------------------------------------+
	//	| J     | The NTLMv2 response from the NtChallengeResponseFields ([MS-NLMP] section        |
	//	|       | 2.2.1.3) was used for authentication and session key generation.                 |
	//	+-------+----------------------------------------------------------------------------------+
	//	| K     | The LMv2 response from the LmChallengeResponseFields ([MS-NLMP] section 2.2.1.3) |
	//	|       | was used for authentication and session key generation.                          |
	//	+-------+----------------------------------------------------------------------------------+
	//	| L     | The LMv2 response from the LmChallengeResponseFields ([MS-NLMP] section          |
	//	|       | 2.2.1.3) was used for authentication and the NTLMv2 response from the            |
	//	|       | NtChallengeResponseFields ([MS-NLMP] section 2.2.1.3) was used session key       |
	//	|       | generation.                                                                      |
	//	+-------+----------------------------------------------------------------------------------+
	//
	// The following flags are valid for [MS-KILE] authentications; settings depend on the
	// configuration of the user and groups referenced in the PAC.
	//
	//	+-------+-------------------------------------------------------------------------------+
	//	|       |                                                                               |
	//	| VALUE |                                  DESCRIPTION                                  |
	//	|       |                                                                               |
	//	+-------+-------------------------------------------------------------------------------+
	//	+-------+-------------------------------------------------------------------------------+
	//	| D     | Indicates that the ExtraSids field is populated and contains additional SIDs. |
	//	+-------+-------------------------------------------------------------------------------+
	//	| H     | Indicates that the ResourceGroupIds field is populated.                       |
	//	+-------+-------------------------------------------------------------------------------+
	UserFlags uint32 `idl:"name:UserFlags" json:"user_flags"`
	// UserSessionKey: A session key that is used for cryptographic operations on a session.
	// This field is valid only when authentication is performed using NTLM. For any other
	// protocol, this field MUST be zero.
	UserSessionKey *UserSessionKey `idl:"name:UserSessionKey" json:"user_session_key"`
	// LogonServer: An RPC_UNICODE_STRING structure that contains the NetBIOS name of the
	// Kerberos KDC that performed the authentication server (AS) ticket request.
	LogonServer *dtyp.UnicodeString `idl:"name:LogonServer" json:"logon_server"`
	// LogonDomainName: An RPC_UNICODE_STRING structure that contains the NetBIOS name of
	// the domain to which this account belongs.
	LogonDomainName *dtyp.UnicodeString `idl:"name:LogonDomainName" json:"logon_domain_name"`
	// LogonDomainId: An RPC_SID structure ([MS-DTYP] section 2.4.2.3) that contains the
	// SID for the domain specified in LogonDomainName. This member is used in conjunction
	// with the UserId, PrimaryGroupId, and GroupIds members to create the user and group
	// SIDs for the client.
	LogonDomainID *dtyp.SID `idl:"name:LogonDomainId" json:"logon_domain_id"`
	// Reserved1: A two-element array of unsigned 32-bit integers. This member is reserved,
	// and each element of the array MUST be zero when sent and MUST be ignored on receipt.
	_ []uint32 `idl:"name:Reserved1"`
	// UserAccountControl: A 32-bit unsigned integer that contains a set of bit flags that
	// represent information about this account. This field carries the UserAccountControl
	// information from the corresponding Security Account Manager field, as specified in
	// [MS-SAMR].
	UserAccountControl uint32 `idl:"name:UserAccountControl" json:"user_account_control"`
	// Reserved3: A 32-bit integer. This member is reserved, and MUST be zero when sent
	// and MUST be ignored on receipt.
	_ []uint32 `idl:"name:Reserved3"`
	// SidCount: A 32-bit unsigned integer that contains the total number of SIDs present
	// in the ExtraSids member. If this member is not zero then the D bit MUST be set in
	// the UserFlags member.
	SIDCount uint32 `idl:"name:SidCount" json:"sid_count"`
	// ExtraSids: A pointer to a list of KERB_SID_AND_ATTRIBUTES (section 2.2.1) structures
	// that contain a list of SIDs corresponding to groups in domains other than the account
	// domain to which the principal belongs. This member is not NULL only if the D bit
	// has been set in the UserFlags member. If the UserId member equals 0x00000000, the
	// first group SID in this member is the SID for this account.
	ExtraSIDs []*KerberosSIDAndAttributes `idl:"name:ExtraSids;size_is:(SidCount)" json:"extra_sids"`
	// ResourceGroupDomainSid: An RPC_SID structure that contains the SID of the domain
	// for the server whose resources the client is authenticating to. This member is used
	// in conjunction with the ResourceGroupIds member to create the group SIDs for the
	// user. If this member is populated, then the H bit MUST be set in the UserFlags member.
	ResourceGroupDomainSID *dtyp.SID `idl:"name:ResourceGroupDomainSid" json:"resource_group_domain_sid"`
	// ResourceGroupCount: A 32-bit unsigned integer that contains the number of resource
	// group identifiers stored in ResourceGroupIds. If this member is not zero, then the
	// H bit MUST be set in the UserFlags member.
	ResourceGroupCount uint32 `idl:"name:ResourceGroupCount" json:"resource_group_count"`
	// ResourceGroupIds: A pointer to a list of GROUP_MEMBERSHIP structures that contain
	// the RIDs and attributes of the account's groups in the resource domain. If this member
	// is not NULL, then the H bit MUST be set in the UserFlags member.
	ResourceGroupIDs []*GroupMembership `idl:"name:ResourceGroupIds;size_is:(ResourceGroupCount)" json:"resource_group_ids"`
}

func (o *KerberosValidationInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.GroupIDs != nil && o.GroupCount == 0 {
		o.GroupCount = uint32(len(o.GroupIDs))
	}
	if o.ExtraSIDs != nil && o.SIDCount == 0 {
		o.SIDCount = uint32(len(o.ExtraSIDs))
	}
	if o.ResourceGroupIDs != nil && o.ResourceGroupCount == 0 {
		o.ResourceGroupCount = uint32(len(o.ResourceGroupIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosValidationInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LogonTime != nil {
		if err := o.LogonTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogoffTime != nil {
		if err := o.LogoffTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.KickOffTime != nil {
		if err := o.KickOffTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordLastSet != nil {
		if err := o.PasswordLastSet.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordCanChange != nil {
		if err := o.PasswordCanChange.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordMustChange != nil {
		if err := o.PasswordMustChange.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.EffectiveName != nil {
		if err := o.EffectiveName.MarshalNDR(ctx, w); err != nil {
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
	if o.LogonScript != nil {
		if err := o.LogonScript.MarshalNDR(ctx, w); err != nil {
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
	if err := w.WriteData(o.LogonCount); err != nil {
		return err
	}
	if err := w.WriteData(o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.WriteData(o.UserID); err != nil {
		return err
	}
	if err := w.WriteData(o.PrimaryGroupID); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupCount); err != nil {
		return err
	}
	if o.GroupIDs != nil || o.GroupCount > 0 {
		_ptr_GroupIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.GroupCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.GroupIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.GroupIDs[i1] != nil {
					if err := o.GroupIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.GroupIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GroupIDs, _ptr_GroupIds); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UserFlags); err != nil {
		return err
	}
	if o.UserSessionKey != nil {
		if err := o.UserSessionKey.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserSessionKey{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogonServer != nil {
		if err := o.LogonServer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogonDomainName != nil {
		if err := o.LogonDomainName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogonDomainID != nil {
		_ptr_LogonDomainId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LogonDomainID != nil {
				if err := o.LogonDomainID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LogonDomainID, _ptr_LogonDomainId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved Reserved1
	for i1 := 0; uint64(i1) < 2; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UserAccountControl); err != nil {
		return err
	}
	// reserved Reserved3
	for i1 := 0; uint64(i1) < 7; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SIDCount); err != nil {
		return err
	}
	if o.ExtraSIDs != nil || o.SIDCount > 0 {
		_ptr_ExtraSids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SIDCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ExtraSIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ExtraSIDs[i1] != nil {
					if err := o.ExtraSIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&KerberosSIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ExtraSIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&KerberosSIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ExtraSIDs, _ptr_ExtraSids); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ResourceGroupDomainSID != nil {
		_ptr_ResourceGroupDomainSid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ResourceGroupDomainSID != nil {
				if err := o.ResourceGroupDomainSID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResourceGroupDomainSID, _ptr_ResourceGroupDomainSid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ResourceGroupCount); err != nil {
		return err
	}
	if o.ResourceGroupIDs != nil || o.ResourceGroupCount > 0 {
		_ptr_ResourceGroupIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ResourceGroupCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ResourceGroupIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ResourceGroupIDs[i1] != nil {
					if err := o.ResourceGroupIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ResourceGroupIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResourceGroupIDs, _ptr_ResourceGroupIds); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosValidationInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LogonTime == nil {
		o.LogonTime = &dtyp.Filetime{}
	}
	if err := o.LogonTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogoffTime == nil {
		o.LogoffTime = &dtyp.Filetime{}
	}
	if err := o.LogoffTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.KickOffTime == nil {
		o.KickOffTime = &dtyp.Filetime{}
	}
	if err := o.KickOffTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordLastSet == nil {
		o.PasswordLastSet = &dtyp.Filetime{}
	}
	if err := o.PasswordLastSet.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordCanChange == nil {
		o.PasswordCanChange = &dtyp.Filetime{}
	}
	if err := o.PasswordCanChange.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordMustChange == nil {
		o.PasswordMustChange = &dtyp.Filetime{}
	}
	if err := o.PasswordMustChange.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.EffectiveName == nil {
		o.EffectiveName = &dtyp.UnicodeString{}
	}
	if err := o.EffectiveName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.FullName == nil {
		o.FullName = &dtyp.UnicodeString{}
	}
	if err := o.FullName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogonScript == nil {
		o.LogonScript = &dtyp.UnicodeString{}
	}
	if err := o.LogonScript.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ProfilePath == nil {
		o.ProfilePath = &dtyp.UnicodeString{}
	}
	if err := o.ProfilePath.UnmarshalNDR(ctx, w); err != nil {
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
	if err := w.ReadData(&o.LogonCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrimaryGroupID); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupCount); err != nil {
		return err
	}
	_ptr_GroupIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.GroupCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.GroupCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.GroupIDs", sizeInfo[0])
		}
		o.GroupIDs = make([]*GroupMembership, sizeInfo[0])
		for i1 := range o.GroupIDs {
			i1 := i1
			if o.GroupIDs[i1] == nil {
				o.GroupIDs[i1] = &GroupMembership{}
			}
			if err := o.GroupIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_GroupIds := func(ptr interface{}) { o.GroupIDs = *ptr.(*[]*GroupMembership) }
	if err := w.ReadPointer(&o.GroupIDs, _s_GroupIds, _ptr_GroupIds); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserFlags); err != nil {
		return err
	}
	if o.UserSessionKey == nil {
		o.UserSessionKey = &UserSessionKey{}
	}
	if err := o.UserSessionKey.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogonServer == nil {
		o.LogonServer = &dtyp.UnicodeString{}
	}
	if err := o.LogonServer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogonDomainName == nil {
		o.LogonDomainName = &dtyp.UnicodeString{}
	}
	if err := o.LogonDomainName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_LogonDomainId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LogonDomainID == nil {
			o.LogonDomainID = &dtyp.SID{}
		}
		if err := o.LogonDomainID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_LogonDomainId := func(ptr interface{}) { o.LogonDomainID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.LogonDomainID, _s_LogonDomainId, _ptr_LogonDomainId); err != nil {
		return err
	}
	// reserved Reserved1
	var _Reserved1 []uint32
	_Reserved1 = make([]uint32, 2)
	for i1 := range _Reserved1 {
		i1 := i1
		if err := w.ReadData(&_Reserved1[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.UserAccountControl); err != nil {
		return err
	}
	// reserved Reserved3
	var _Reserved3 []uint32
	_Reserved3 = make([]uint32, 7)
	for i1 := range _Reserved3 {
		i1 := i1
		if err := w.ReadData(&_Reserved3[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.SIDCount); err != nil {
		return err
	}
	_ptr_ExtraSids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SIDCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SIDCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ExtraSIDs", sizeInfo[0])
		}
		o.ExtraSIDs = make([]*KerberosSIDAndAttributes, sizeInfo[0])
		for i1 := range o.ExtraSIDs {
			i1 := i1
			if o.ExtraSIDs[i1] == nil {
				o.ExtraSIDs[i1] = &KerberosSIDAndAttributes{}
			}
			if err := o.ExtraSIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ExtraSids := func(ptr interface{}) { o.ExtraSIDs = *ptr.(*[]*KerberosSIDAndAttributes) }
	if err := w.ReadPointer(&o.ExtraSIDs, _s_ExtraSids, _ptr_ExtraSids); err != nil {
		return err
	}
	_ptr_ResourceGroupDomainSid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ResourceGroupDomainSID == nil {
			o.ResourceGroupDomainSID = &dtyp.SID{}
		}
		if err := o.ResourceGroupDomainSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ResourceGroupDomainSid := func(ptr interface{}) { o.ResourceGroupDomainSID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.ResourceGroupDomainSID, _s_ResourceGroupDomainSid, _ptr_ResourceGroupDomainSid); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResourceGroupCount); err != nil {
		return err
	}
	_ptr_ResourceGroupIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ResourceGroupCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ResourceGroupCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ResourceGroupIDs", sizeInfo[0])
		}
		o.ResourceGroupIDs = make([]*GroupMembership, sizeInfo[0])
		for i1 := range o.ResourceGroupIDs {
			i1 := i1
			if o.ResourceGroupIDs[i1] == nil {
				o.ResourceGroupIDs[i1] = &GroupMembership{}
			}
			if err := o.ResourceGroupIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ResourceGroupIds := func(ptr interface{}) { o.ResourceGroupIDs = *ptr.(*[]*GroupMembership) }
	if err := w.ReadPointer(&o.ResourceGroupIDs, _s_ResourceGroupIds, _ptr_ResourceGroupIds); err != nil {
		return err
	}
	return nil
}

// PACCredentialInfo structure represents PAC_CREDENTIAL_INFO RPC structure.
//
// The PAC_CREDENTIAL_INFO structure serves as the header for the credential information.
// The PAC_CREDENTIAL_INFO header indicates the encryption algorithm that was used to
// encrypt the data that follows it. The data that follows is an encrypted, IDL-serialized
// PAC_CREDENTIAL_DATA structure that contains the user's actual credentials. Note that
// this structure cannot be used by protocols other than the [MS-KILE] protocol; the
// encryption method relies on the encryption key currently in use by the Kerberos AS-REQ
// ([RFC4120] section 3.1 and [MS-KILE]) message.<13>
//
// A PAC_CREDENTIAL_INFO structure contains the user's encrypted  credentials. The
// Key Usage Number [RFC4120] used in the encryption is KERB_NON_KERB_SALT [16] [MS-KILE]
// section 3.1.5.9. The encryption key used is the AS reply key. The PAC credentials
// buffer is included only when PKINIT [RFC4556] is used. Therefore, the AS reply key
// is derived based on PKINIT.
//
// The PAC_CREDENTIAL_INFO structure is defined as follows.
//
// The format of the PAC_CREDENTIAL_INFO structure is defined as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Version                                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| EncryptionType                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| SerializedData (variable)                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type PACCredentialInfo struct {
	// Version (4 bytes): A 32-bit unsigned integer in little-endian format that defines
	// the version. MUST be 0x00000000.
	Version uint32 `idl:"name:Version" json:"version"`
	// EncryptionType (4 bytes): A 32-bit unsigned integer in little-endian format that
	// indicates the Kerberos encryption type used to encode the SerializedData array. This
	// value MUST be one of the following encryption types, which are a subset of the possible
	// encryption types supported in Kerberos authentication (as specified in [RFC4120],
	// [RFC4757], and [RFC4556]). Note that the Key Usage Number ([RFC4120] sections 4 and
	// 7.5.1) is KERB_NON_KERB_SALT [16] [MS-KILE] section 3.1.5.9.<14>
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | Data Encryption Standard (DES) in cipher block chaining (CBC) mode with cyclic   |
	//	|            | redundancy check (CRC).                                                          |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000003 | DES in CBC mode with MD5.                                                        |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000011 | AES128_CTS_HMAC_SHA1_96 (128-bit encryption key in clear to send (CTS)           |
	//	|            | encryption mode with integrity check algorithm HMAC_SHA1_96).<15>                |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000012 | AES256_CTS_HMAC_SHA1_96 (256-bit encryption key in CTS encryption mode with      |
	//	|            | integrity check algorithm HMAC_SHA1_96).<16>                                     |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000017 | RC4 with hashed message authentication code (HMAC) key.                          |
	//	+------------+----------------------------------------------------------------------------------+
	EncryptionType uint32 `idl:"name:EncryptionType" json:"encryption_type"`
	// SerializedData (variable): A variable length PAC_CREDENTIAL_DATA structure that contains
	// credentials encrypted using the mechanism specified by the EncryptionType field.
	// The byte array of encrypted data is computed according to the procedures specified
	// in [RFC3961].
	SerializedData []byte `idl:"name:SerializedData" json:"serialized_data"`
}

func (o *PACCredentialInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACCredentialInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.EncryptionType); err != nil {
		return err
	}
	if o.SerializedData != nil {
		_ptr_SerializedData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.SerializedData {
				i1 := i1
				if err := w.WriteData(o.SerializedData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SerializedData, _ptr_SerializedData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACCredentialInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionType); err != nil {
		return err
	}
	_ptr_SerializedData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.SerializedData = append(o.SerializedData, uint8(0))
			if err := w.ReadData(&o.SerializedData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_SerializedData := func(ptr interface{}) { o.SerializedData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SerializedData, _s_SerializedData, _ptr_SerializedData); err != nil {
		return err
	}
	return nil
}

// PACCredentialData structure represents PAC_CREDENTIAL_DATA RPC structure.
//
// The PAC_CREDENTIAL_DATA structure defines an array of security package-specific credentials
// that are provided to the Kerberos client. The PAC_CREDENTIAL_DATA structure is marshaled
// by RPC [MS-RPCE].
//
// The PAC_CREDENTIAL_DATA structure is defined as follows.
type PACCredentialData struct {
	// CredentialCount: A 32-bit unsigned integer that defines the number of elements in
	// the Credentials member.
	CredentialCount uint32 `idl:"name:CredentialCount" json:"credential_count"`
	// Credentials: An array of SECPKG_SUPPLEMENTAL_CRED (section 2.6.3) structures that
	// define the supplemental credentials.
	Credentials []*SecurityPackageSupplementalCred `idl:"name:Credentials;size_is:(CredentialCount)" json:"credentials"`
}

func (o *PACCredentialData) xxx_PreparePayload(ctx context.Context) error {
	if o.Credentials != nil && o.CredentialCount == 0 {
		o.CredentialCount = uint32(len(o.Credentials))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *PACCredentialData) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.CredentialCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PACCredentialData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.CredentialCount); err != nil {
		return err
	}
	for i1 := range o.Credentials {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Credentials[i1] != nil {
			if err := o.Credentials[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityPackageSupplementalCred{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Credentials); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&SecurityPackageSupplementalCred{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACCredentialData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.CredentialCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.CredentialCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.CredentialCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Credentials", sizeInfo[0])
	}
	o.Credentials = make([]*SecurityPackageSupplementalCred, sizeInfo[0])
	for i1 := range o.Credentials {
		i1 := i1
		if o.Credentials[i1] == nil {
			o.Credentials[i1] = &SecurityPackageSupplementalCred{}
		}
		if err := o.Credentials[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// SecurityPackageSupplementalCred structure represents SECPKG_SUPPLEMENTAL_CRED RPC structure.
//
// The SECPKG_SUPPLEMENTAL_CRED structure defines the name of the security package that
// requires supplemental credentials and the credential buffer for that package. The
// SECPKG_SUPPLEMENTAL_CRED structure is marshaled by RPC [MS-RPCE].
//
// The SECPKG_SUPPLEMENTAL_CRED structure is defined as follows.
type SecurityPackageSupplementalCred struct {
	// PackageName: A RPC_UNICODE_STRING structure that MUST store the name of the security
	// protocol for which the supplemental credentials are being presented.<17>
	PackageName *dtyp.UnicodeString `idl:"name:PackageName" json:"package_name"`
	// CredentialSize: A 32-bit unsigned integer that MUST specify the length, in bytes,
	// of the data in the Credentials member.
	CredentialSize uint32 `idl:"name:CredentialSize" json:"credential_size"`
	// Credentials: A pointer that MUST reference the serialized credentials being presented
	// to the security protocol named in PackageName.
	Credentials                []byte                      `idl:"name:Credentials;size_is:(CredentialSize)" json:"credentials"`
	NTLMSupplementalCredential *NTLMSupplementalCredential `idl:"name:NtlmSupplementalCredential" json:"ntlm_supplemental_credential"`
}

func (o *SecurityPackageSupplementalCred) xxx_PreparePayload(ctx context.Context) error {
	if o.Credentials != nil && o.CredentialSize == 0 {
		o.CredentialSize = uint32(len(o.Credentials))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityPackageSupplementalCred) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.PackageName != nil {
		if err := o.PackageName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CredentialSize); err != nil {
		return err
	}
	if o.Credentials != nil || o.CredentialSize > 0 {
		_ptr_Credentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.CredentialSize)
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
				if err := w.WriteData(o.Credentials[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Credentials); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
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
	return nil
}
func (o *SecurityPackageSupplementalCred) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.PackageName == nil {
		o.PackageName = &dtyp.UnicodeString{}
	}
	if err := o.PackageName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.CredentialSize); err != nil {
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
		if o.CredentialSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.CredentialSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Credentials", sizeInfo[0])
		}
		o.Credentials = make([]byte, sizeInfo[0])
		for i1 := range o.Credentials {
			i1 := i1
			if err := w.ReadData(&o.Credentials[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Credentials := func(ptr interface{}) { o.Credentials = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Credentials, _s_Credentials, _ptr_Credentials); err != nil {
		return err
	}
	return nil
}

// NTLMSupplementalCredential structure represents NTLM_SUPPLEMENTAL_CREDENTIAL RPC structure.
//
// The NTLM_SUPPLEMENTAL_CREDENTIAL structure is used to encode the credentials that
// the NTLM security protocol uses, specifically the LAN Manager hash (LM OWF) and the
// NT hash (NT OWF). Generating the hashes encoded in this structure is not addressed
// in the PAC structure specification. Details on how the hashes are created are as
// specified in [MS-NLMP]. The PAC buffer type is included only when PKINIT [MS-PKCA]
// is used to authenticate the user. The NTLM_SUPPLEMENTAL_CREDENTIAL structure is marshaled
// by RPC [MS-RPCE].
//
// The NTLM_SUPPLEMENTAL_CREDENTIAL structure is defined as follows.
type NTLMSupplementalCredential struct {
	// Version: A 32-bit unsigned integer that defines the credential version. This field
	// MUST be 0x00000000.
	Version uint32 `idl:"name:Version" json:"version"`
	// Flags: A 32-bit unsigned integer containing flags that define the credential options.
	// Flags MUST contain at least one of the following values.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | N | L |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// Where the bits are defined as:
	//
	//	+-------+--------------------------------------------------------+
	//	|       |                                                        |
	//	| VALUE |                      DESCRIPTION                       |
	//	|       |                                                        |
	//	+-------+--------------------------------------------------------+
	//	+-------+--------------------------------------------------------+
	//	| L     | Indicates that the LM OWF member is present and valid. |
	//	+-------+--------------------------------------------------------+
	//	| N     | Indicates that the NT OWF member is present and valid. |
	//	+-------+--------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// LmPassword: A 16-element array of unsigned 8-bit integers that define the LM OWF.
	// The LmPassword member MUST be ignored if the L flag is not set in the Flags member.
	LMPassword []byte `idl:"name:LmPassword" json:"lm_password"`
	// NtPassword: A 16-element array of unsigned 8-bit integers that define the NT OWF.
	// The NtPassword member MUST be ignored if the N flag is not set in the Flags member.
	NTPassword []byte `idl:"name:NtPassword" json:"nt_password"`
}

func (o *NTLMSupplementalCredential) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTLMSupplementalCredential) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	for i1 := range o.LMPassword {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.LMPassword[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LMPassword); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.NTPassword {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.NTPassword[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.NTPassword); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTLMSupplementalCredential) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	o.LMPassword = make([]byte, 16)
	for i1 := range o.LMPassword {
		i1 := i1
		if err := w.ReadData(&o.LMPassword[i1]); err != nil {
			return err
		}
	}
	o.NTPassword = make([]byte, 16)
	for i1 := range o.NTPassword {
		i1 := i1
		if err := w.ReadData(&o.NTPassword[i1]); err != nil {
			return err
		}
	}
	return nil
}

// PACClientInfo structure represents PAC_CLIENT_INFO RPC structure.
//
// The PAC_CLIENT_INFO structure is a variable length buffer of the PAC that contains
// the client's name and authentication time. It is used to verify that the PAC corresponds
// to the client of the ticket. The PAC_CLIENT_INFO structure is placed directly after
// the Buffers array of the topmost PACTYPE structure (section 2.3), at the offset specified
// in the Offset field of the corresponding PAC_INFO_BUFFER structure (section 2.4)
// in the Buffers array. The ulType field of the corresponding PAC_INFO_BUFFER is set
// to 0x0000000A.
//
// The PAC_CLIENT_INFO structure is defined as follows.
//
// The format of the PAC_CLIENT_INFO structure is defined as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ClientId                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| NameLength                                                    | Name (variable)                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type PACClientInfo struct {
	// ClientId (8 bytes): A FILETIME structure in little-endian format that contains the
	// Kerberos initial ticket-granting ticket (TGT) authentication time, as specified in
	// [RFC4120] section 5.3.
	ClientID *dtyp.Filetime `idl:"name:ClientId" json:"client_id"`
	// NameLength (2 bytes): An unsigned 16-bit integer in little-endian format that specifies
	// the length, in bytes, of the Name field.
	NameLength uint16 `idl:"name:NameLength" json:"name_length"`
	// Name (variable): An array of 16-bit Unicode characters in little-endian format that
	// contains the client's account name.
	Name string `idl:"name:Name;size_is:((NameLength/2))" json:"name"`
}

func (o *PACClientInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.Name != "" && o.NameLength == 0 {
		o.NameLength = uint16((len(o.Name) * 2))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACClientInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ClientID != nil {
		if err := o.ClientID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NameLength); err != nil {
		return err
	}
	if o.Name != "" || (o.NameLength/2) > 0 {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64((o.NameLength / 2))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_Name_buf := utf16.Encode([]rune(o.Name))
			if uint64(len(_Name_buf)) > sizeInfo[0] {
				_Name_buf = _Name_buf[:sizeInfo[0]]
			}
			for i1 := range _Name_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Name_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Name_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
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
	return nil
}
func (o *PACClientInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.ClientID == nil {
		o.ClientID = &dtyp.Filetime{}
	}
	if err := o.ClientID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameLength); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _Name_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Name_buf", sizeInfo[0])
		}
		_Name_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Name_buf {
			i1 := i1
			if err := w.ReadData(&_Name_buf[i1]); err != nil {
				return err
			}
		}
		o.Name = strings.TrimRight(string(utf16.Decode(_Name_buf)), ndr.ZeroString)
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// PACSignatureData structure represents PAC_SIGNATURE_DATA RPC structure.
//
// Two PAC_SIGNATURE_DATA structures are appended to the PAC which stores the server
// and KDC signatures. These structures are placed after the Buffers array of the topmost
// PACTYPE structure (section 2.3), at the offsets specified in the Offset fields in
// each of the corresponding PAC_INFO_BUFFER structures (section 2.4) in the Buffers
// array. The ulType field of the PAC_INFO_BUFFER corresponding to the server signature
// contains the value 0x00000006 and the ulType field of the PAC_INFO_BUFFER corresponding
// to the KDC signature contains the value 0x00000007. PAC signatures can be generated
// only when the PAC is used by the [MS-KILE] protocol because the keys used to create
// and verify the signatures are the keys known to the KDC. No other protocol can use
// these PAC signatures.
//
// The PAC_SIGNATURE_DATA structure is defined as follows.
//
// The format of the PAC_SIGNATURE_DATA structures is defined as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| SignatureType                                                                                                                 |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Signature (variable)                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| RODCIdentifier                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type PACSignatureData struct {
	// SignatureType (4 bytes): A 32-bit unsigned integer value in little-endian format
	// that defines the cryptographic system used to calculate the checksum. This MUST be
	// one of the values defined in the following table. The corresponding sizes of the
	// signatures are also given. The key used with the cryptographic system corresponds
	// to the value of the ulType field of the outer PAC_INFO_BUFFER (section 2.4) structure.
	// The value 0x00000006 specifies the server's key, and the value 0x00000007 specifies
	// the KDC's key.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| KERB_CHECKSUM_HMAC_MD5 0xFFFFFF76 | As specified in [RFC4120] and [RFC4757] section 4. Signature size is 16 bytes.   |
	//	|                                   | Decimal value is -138.                                                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| HMAC_SHA1_96_AES128 0x0000000F    | As specified in [RFC3962] section 7. Signature size is 12 bytes. Decimal value   |
	//	|                                   | is 15.                                                                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| HMAC_SHA1_96_AES256 0x00000010    | As specified in [RFC3962] section 7. Signature size is 12 bytes. Decimal value   |
	//	|                                   | is 16.                                                                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	SignatureType uint32 `idl:"name:SignatureType" json:"signature_type"`
	// Signature (variable): An array of 8-bit unsigned characters that contains the checksum.
	// The KERB_CHECKSUM_HMAC_MD5 checksum (defined in the preceding table) is 16 bytes
	// in length. The size of the signature is determined by the value of the SignatureType
	// field, as indicated in the preceding table.
	Signature []byte `idl:"name:Signature;size_is:(((SignatureType==4294967158)?16:12))" json:"signature"`
}

func (o *PACSignatureData) xxx_PreparePayload(ctx context.Context) error {
	if o.Signature != nil && o.SignatureType == 0 {
		o.SignatureType = uint32(len(o.Signature))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACSignatureData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SignatureType); err != nil {
		return err
	}
	_exprSignatureType := uint32(0)
	if o.SignatureType == 4294967158 {
		_exprSignatureType = uint32(16)
	} else {
		_exprSignatureType = uint32(12)
	}
	if o.Signature != nil || _exprSignatureType > 0 {
		_ptr_Signature := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprSignatureType := uint64(0)
			if o.SignatureType == 4294967158 {
				_exprSignatureType = uint64(16)
			} else {
				_exprSignatureType = uint64(12)
			}
			dimSize1 := uint64(_exprSignatureType)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Signature {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Signature[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Signature); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Signature, _ptr_Signature); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACSignatureData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SignatureType); err != nil {
		return err
	}
	_ptr_Signature := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprSignature := uint64(0)
		if o.SignatureType == 4294967158 {
			_exprSignature = uint64(16)
		} else {
			_exprSignature = uint64(12)
		}
		if _exprSignature > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprSignature)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Signature", sizeInfo[0])
		}
		o.Signature = make([]byte, sizeInfo[0])
		for i1 := range o.Signature {
			i1 := i1
			if err := w.ReadData(&o.Signature[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Signature := func(ptr interface{}) { o.Signature = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Signature, _s_Signature, _ptr_Signature); err != nil {
		return err
	}
	return nil
}

// S4UDelegationInfo structure represents S4U_DELEGATION_INFO RPC structure.
//
// The S4U_DELEGATION_INFO structure is used for constrained delegation information.<22>
// It lists the services that have been delegated through this Kerberos client and subsequent
// services or servers. The list is used only in a Service for User to Proxy (S4U2proxy)
// [MS-SFU] request. This feature could be used multiple times in succession from service
// to service, which is useful for auditing purposes. The S4U_DELEGATION_INFO structure
// is marshaled by RPC [MS-RPCE].
//
// The S4U_DELEGATION_INFO structure is defined as follows.
type S4UDelegationInfo struct {
	// S4U2proxyTarget: An RPC_UNICODE_STRING structure that MUST contain the name of the
	// principal to whom the application can forward the ticket.
	S4u2proxyTarget *dtyp.UnicodeString `idl:"name:S4U2proxyTarget" json:"s4u2proxy_target"`
	// TransitedListSize: MUST be the number of elements in the S4UTransitedServices array.
	TransitedListSize uint32 `idl:"name:TransitedListSize" json:"transited_list_size"`
	// S4UTransitedServices: MUST contain the list of all services that have been delegated
	// through by this client and subsequent services or servers.
	S4UTransitedServices []*dtyp.UnicodeString `idl:"name:S4UTransitedServices;size_is:(TransitedListSize)" json:"s4u_transited_services"`
}

func (o *S4UDelegationInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.S4UTransitedServices != nil && o.TransitedListSize == 0 {
		o.TransitedListSize = uint32(len(o.S4UTransitedServices))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *S4UDelegationInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.S4u2proxyTarget != nil {
		if err := o.S4u2proxyTarget.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TransitedListSize); err != nil {
		return err
	}
	if o.S4UTransitedServices != nil || o.TransitedListSize > 0 {
		_ptr_S4UTransitedServices := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.TransitedListSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.S4UTransitedServices {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.S4UTransitedServices[i1] != nil {
					if err := o.S4UTransitedServices[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.S4UTransitedServices); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.S4UTransitedServices, _ptr_S4UTransitedServices); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *S4UDelegationInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.S4u2proxyTarget == nil {
		o.S4u2proxyTarget = &dtyp.UnicodeString{}
	}
	if err := o.S4u2proxyTarget.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.TransitedListSize); err != nil {
		return err
	}
	_ptr_S4UTransitedServices := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.TransitedListSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.TransitedListSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.S4UTransitedServices", sizeInfo[0])
		}
		o.S4UTransitedServices = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.S4UTransitedServices {
			i1 := i1
			if o.S4UTransitedServices[i1] == nil {
				o.S4UTransitedServices[i1] = &dtyp.UnicodeString{}
			}
			if err := o.S4UTransitedServices[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_S4UTransitedServices := func(ptr interface{}) { o.S4UTransitedServices = *ptr.(*[]*dtyp.UnicodeString) }
	if err := w.ReadPointer(&o.S4UTransitedServices, _s_S4UTransitedServices, _ptr_S4UTransitedServices); err != nil {
		return err
	}
	return nil
}

// UPNDNSInfo structure represents UPN_DNS_INFO RPC structure.
//
// The UPN_DNS_INFO structure contains the client's UPN, fully qualified domain name
// (FQDN), SAM name (optional), and SID (optional). It is used to provide the UPN, FQDN,
// SAM name, and SID that corresponds to the client of the ticket. The UPN_DNS_INFO
// structure is placed directly after the Buffers array of the topmost PACTYPE structure
// (section 2.3), at the offset specified in the Offset field of the corresponding PAC_INFO_BUFFER
// structure (section 2.4) in the Buffers array. The ulType field of the corresponding
// PAC_INFO_BUFFER is set to 0x0000000C.<23>
//
// The UPN_DNS_INFO structure is defined as follows.
//
// The format of the UPN_DNS_INFO structure is defined as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| UpnLength                                                     | UpnOffset                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| DnsDomainNameLength                                           | DnsDomainNameOffset                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Flags                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| SamNameLength                                                 | SamNameOffset                                                 |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| SidLength                                                     | SidOffset                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type UPNDNSInfo struct {
	// UpnLength (2 bytes): An unsigned 16-bit integer in little-endian format that specifies
	// the length, in bytes, of the UPN information. The location of the UPN information
	// is described later in this section.
	UPNLength uint16 `idl:"name:UpnLength" json:"upn_length"`
	// UpnOffset (2 bytes): An unsigned 16-bit integer in little-endian format that contains
	// the offset to the beginning of the UPN information, in bytes, from the beginning
	// of the UPN_DNS_INFO structure (section 2.10).
	UPNOffset uint16 `idl:"name:UpnOffset" json:"upn_offset"`
	// DnsDomainNameLength (2 bytes): An unsigned 16-bit integer in little-endian format
	// that specifies the length, in bytes, of the DNS information. The location of the
	// DNS information is described later in this section.
	DNSDomainNameLength uint16 `idl:"name:DnsDomainNameLength" json:"dns_domain_name_length"`
	// DnsDomainNameOffset (2 bytes): An unsigned 16-bit integer in little-endian format
	// that contains the offset to the beginning of the DNS information, in bytes, from
	// the beginning of the UPN_DNS_INFO structure.
	DNSDomainNameOffset uint16 `idl:"name:DnsDomainNameOffset" json:"dns_domain_name_offset"`
	// Flags (4 bytes): A set of bit flags in little-endian format. A flag is TRUE (or set)
	// if its value is equal to 1. The value is constructed from zero or more bit flags
	// from the following table:
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | S | U |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// Where the bits are defined as:
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                   DESCRIPTION                                    |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	| U     | The user account object does not have the userPrincipalName attribute ([MS-ADA3] |
	//	|       | section 2.349) set. A UPN constructed by concatenating the user name with the    |
	//	|       | DNS domain name of the account domain is provided.                               |
	//	+-------+----------------------------------------------------------------------------------+
	//	| S     | The UPN_DNS_INFO structure has been extended with the user account’s SAM Name    |
	//	|       | and SID.                                                                         |
	//	+-------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	Raw   []byte `idl:"name:Raw" json:"raw"`
	// SamNameLength (2 bytes): An unsigned 16-bit integer in little-endian format that
	// specifies the length, in bytes, of the SAM name. The location of the SAM name is
	// described later in this section. This field is only present if the S flag bit is
	// set.
	SAMNameLength uint16 `idl:"name:SamNameLength" json:"sam_name_length"`
	// SamNameOffset (2 bytes): An unsigned 16-bit integer in little-endian format that
	// contains the offset to the beginning of the SAM name, in bytes, from the beginning
	// of the UPN_DNS_INFO structure. This field is only present if the S flag bit is set.
	SAMNameOffset uint16 `idl:"name:SamNameOffset" json:"sam_name_offset"`
	// SidLength (2 bytes): An unsigned 16-bit integer in little-endian format that specifies
	// the length, in bytes, of the client’s SID. The location of the SID is described
	// later in this section. This field is only present if the S flag bit is set.
	SIDLength uint16 `idl:"name:SidLength" json:"sid_length"`
	// SidOffset (2 byte): An unsigned 16-bit integer in little-endian format that contains
	// the offset to the beginning of the client’s SID, in bytes, from the beginning of
	// the UPN_DNS_INFO structure. This field is only present if the S flag bit is set.
	//
	// The actual DNS and UPN information (and, if the S flag bit is set, the SAM name and
	// SID) is placed after the UPN_DNS_INFO structure following the header and starting
	// with the corresponding offset in a consecutive buffer. The UPN, FQDN, and SAM name
	// are encoded using a two-byte UTF16 scheme, in little-endian order.
	SIDOffset     uint16    `idl:"name:SidOffset" json:"sid_offset"`
	UPN           string    `idl:"name:Upn" json:"upn"`
	DNSDomainName string    `idl:"name:DnsDomainName" json:"dns_domain_name"`
	SAMName       string    `idl:"name:SamName" json:"sam_name"`
	SID           *dtyp.SID `idl:"name:Sid" json:"sid"`
}

func (o *UPNDNSInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UPNDNSInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.UPNLength); err != nil {
		return err
	}
	if err := w.WriteData(o.UPNOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.DNSDomainNameLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DNSDomainNameOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.Raw != nil {
		_ptr_Raw := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.Raw {
				i1 := i1
				if err := w.WriteData(o.Raw[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Raw, _ptr_Raw); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UPNDNSInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.UPNLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.UPNOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.DNSDomainNameLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DNSDomainNameOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_Raw := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.Raw = append(o.Raw, uint8(0))
			if err := w.ReadData(&o.Raw[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Raw := func(ptr interface{}) { o.Raw = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Raw, _s_Raw, _ptr_Raw); err != nil {
		return err
	}
	return nil
}

// PACClientClaimsInfo structure represents PAC_CLIENT_CLAIMS_INFO RPC structure.
//
// The PAC_CLIENT_CLAIMS_ INFO structure is a variable length buffer of the PAC that
// SHOULD<24> contain the client's marshaled claims blob. The PAC_CLIENT_CLAIMS_ INFO
// structure is placed directly after the Buffers array of the topmost PACTYPE structure
// (section 2.3), at the offset specified in the Offset field of the corresponding PAC_INFO_BUFFER
// structure (section 2.4) in the Buffers array. The ulType field of the corresponding
// PAC_INFO_BUFFER is set to 0x0000000D.
//
// The PAC_CLIENT_CLAIMS_ INFO structure is defined as follows.
//
// The format of the PAC_CLIENT_CLAIMS_ INFO structure is defined as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Claims (variable)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type PACClientClaimsInfo struct {
	// Claims (variable): A variable-length CLAIMS_SET_METADATA structure ([MS-ADTS] section
	// 2.2.18.8) that contains claims.
	Claims *claims.ClaimsSetMetadata `idl:"name:Claims" json:"claims"`
}

func (o *PACClientClaimsInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACClientClaimsInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Claims != nil {
		_ptr_Claims := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Claims != nil {
				if err := o.Claims.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&claims.ClaimsSetMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Claims, _ptr_Claims); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACClientClaimsInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Claims := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Claims == nil {
			o.Claims = &claims.ClaimsSetMetadata{}
		}
		if err := o.Claims.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Claims := func(ptr interface{}) { o.Claims = *ptr.(**claims.ClaimsSetMetadata) }
	if err := w.ReadPointer(&o.Claims, _s_Claims, _ptr_Claims); err != nil {
		return err
	}
	return nil
}

// PACDeviceInfo structure represents PAC_DEVICE_INFO RPC structure.
//
// The PAC_DEVICE_INFO structure is a variable length buffer of the PAC that SHOULD<25>
// contain the device's logon and authorization information provided by the DC. A pointer
// to the PAC_DEVICE_INFO structure is serialized into an array of bytes and placed
// directly after the Buffers array of the topmost PACTYPE structure (section 2.3),
// at the offset specified in the Offset field of the corresponding PAC_INFO_BUFFER
// structure (section 2.4) in the Buffers array. The ulType field of the corresponding
// PAC_INFO_BUFFER is set to 0x0000000E.
//
// The PAC_DEVICE_INFO structure is defined as follows.
type PACDeviceInfo struct {
	// UserId: A 32-bit unsigned integer that contains the RID of the account. If the UserId
	// member equals 0x00000000, the first group SID in this member is the SID for this
	// account.
	UserID uint32 `idl:"name:UserId" json:"user_id"`
	// PrimaryGroupId: A 32-bit unsigned integer that contains the RID for the primary group
	// to which this account belongs.
	PrimaryGroupID uint32 `idl:"name:PrimaryGroupId" json:"primary_group_id"`
	// AccountDomainId: A SID structure that contains the SID for the domain of the account.
	// This member is used in conjunction with the UserId, and GroupIds members to create
	// the user and group SIDs for the client.
	AccountDomainID *dtyp.SID `idl:"name:AccountDomainId" json:"account_domain_id"`
	// AccountGroupCount: A 32-bit unsigned integer that contains the number of groups within
	// the account domain to which the account belongs.
	AccountGroupCount uint32 `idl:"name:AccountGroupCount" json:"account_group_count"`
	// AccountGroupIds: A pointer to a list of GROUP_MEMBERSHIP (section 2.2.2) structures
	// that contains the groups to which the account belongs in the account domain. The
	// number of groups in this list MUST be equal to GroupCount.
	AccountGroupIDs []*GroupMembership `idl:"name:AccountGroupIds;size_is:(AccountGroupCount)" json:"account_group_ids"`
	// SidCount: A 32-bit unsigned integer that contains the total number of SIDs present
	// in the ExtraSids member.
	SIDCount uint32 `idl:"name:SidCount" json:"sid_count"`
	// ExtraSids: A pointer to a list of KERB_SID_AND_ATTRIBUTES structures (section 2.2.1)
	// that contain a list of SIDs corresponding to groups not in domains. If the UserId
	// member equals 0x00000000, the first group SID in this member is the SID for this
	// account.
	ExtraSIDs []*KerberosSIDAndAttributes `idl:"name:ExtraSids;size_is:(SidCount)" json:"extra_sids"`
	// DomainGroupCount: A 32-bit unsigned integer that contains the number of domains with
	// groups to which the account belongs.
	DomainGroupCount uint32 `idl:"name:DomainGroupCount" json:"domain_group_count"`
	// DomainGroup: A pointer to a list of DOMAIN_GROUP_MEMBERSHIP structures (section 2.2.3)
	// that contains the domains to which the account belongs to a group. The number of
	// sets in this list MUST be equal to DomainCount.
	DomainGroup []*DomainGroupMembership `idl:"name:DomainGroup;size_is:(DomainGroupCount)" json:"domain_group"`
}

func (o *PACDeviceInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.AccountGroupIDs != nil && o.AccountGroupCount == 0 {
		o.AccountGroupCount = uint32(len(o.AccountGroupIDs))
	}
	if o.ExtraSIDs != nil && o.SIDCount == 0 {
		o.SIDCount = uint32(len(o.ExtraSIDs))
	}
	if o.DomainGroup != nil && o.DomainGroupCount == 0 {
		o.DomainGroupCount = uint32(len(o.DomainGroup))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACDeviceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.UserID); err != nil {
		return err
	}
	if err := w.WriteData(o.PrimaryGroupID); err != nil {
		return err
	}
	if o.AccountDomainID != nil {
		_ptr_AccountDomainId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AccountDomainID != nil {
				if err := o.AccountDomainID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccountDomainID, _ptr_AccountDomainId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AccountGroupCount); err != nil {
		return err
	}
	if o.AccountGroupIDs != nil || o.AccountGroupCount > 0 {
		_ptr_AccountGroupIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AccountGroupCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.AccountGroupIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.AccountGroupIDs[i1] != nil {
					if err := o.AccountGroupIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.AccountGroupIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccountGroupIDs, _ptr_AccountGroupIds); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SIDCount); err != nil {
		return err
	}
	if o.ExtraSIDs != nil || o.SIDCount > 0 {
		_ptr_ExtraSids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SIDCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ExtraSIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ExtraSIDs[i1] != nil {
					if err := o.ExtraSIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&KerberosSIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ExtraSIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&KerberosSIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ExtraSIDs, _ptr_ExtraSids); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DomainGroupCount); err != nil {
		return err
	}
	if o.DomainGroup != nil || o.DomainGroupCount > 0 {
		_ptr_DomainGroup := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DomainGroupCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.DomainGroup {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.DomainGroup[i1] != nil {
					if err := o.DomainGroup[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&DomainGroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.DomainGroup); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&DomainGroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainGroup, _ptr_DomainGroup); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACDeviceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrimaryGroupID); err != nil {
		return err
	}
	_ptr_AccountDomainId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AccountDomainID == nil {
			o.AccountDomainID = &dtyp.SID{}
		}
		if err := o.AccountDomainID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AccountDomainId := func(ptr interface{}) { o.AccountDomainID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.AccountDomainID, _s_AccountDomainId, _ptr_AccountDomainId); err != nil {
		return err
	}
	if err := w.ReadData(&o.AccountGroupCount); err != nil {
		return err
	}
	_ptr_AccountGroupIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AccountGroupCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AccountGroupCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AccountGroupIDs", sizeInfo[0])
		}
		o.AccountGroupIDs = make([]*GroupMembership, sizeInfo[0])
		for i1 := range o.AccountGroupIDs {
			i1 := i1
			if o.AccountGroupIDs[i1] == nil {
				o.AccountGroupIDs[i1] = &GroupMembership{}
			}
			if err := o.AccountGroupIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_AccountGroupIds := func(ptr interface{}) { o.AccountGroupIDs = *ptr.(*[]*GroupMembership) }
	if err := w.ReadPointer(&o.AccountGroupIDs, _s_AccountGroupIds, _ptr_AccountGroupIds); err != nil {
		return err
	}
	if err := w.ReadData(&o.SIDCount); err != nil {
		return err
	}
	_ptr_ExtraSids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SIDCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SIDCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ExtraSIDs", sizeInfo[0])
		}
		o.ExtraSIDs = make([]*KerberosSIDAndAttributes, sizeInfo[0])
		for i1 := range o.ExtraSIDs {
			i1 := i1
			if o.ExtraSIDs[i1] == nil {
				o.ExtraSIDs[i1] = &KerberosSIDAndAttributes{}
			}
			if err := o.ExtraSIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ExtraSids := func(ptr interface{}) { o.ExtraSIDs = *ptr.(*[]*KerberosSIDAndAttributes) }
	if err := w.ReadPointer(&o.ExtraSIDs, _s_ExtraSids, _ptr_ExtraSids); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainGroupCount); err != nil {
		return err
	}
	_ptr_DomainGroup := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DomainGroupCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DomainGroupCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DomainGroup", sizeInfo[0])
		}
		o.DomainGroup = make([]*DomainGroupMembership, sizeInfo[0])
		for i1 := range o.DomainGroup {
			i1 := i1
			if o.DomainGroup[i1] == nil {
				o.DomainGroup[i1] = &DomainGroupMembership{}
			}
			if err := o.DomainGroup[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_DomainGroup := func(ptr interface{}) { o.DomainGroup = *ptr.(*[]*DomainGroupMembership) }
	if err := w.ReadPointer(&o.DomainGroup, _s_DomainGroup, _ptr_DomainGroup); err != nil {
		return err
	}
	return nil
}

// PACDeviceClaimsInfo structure represents PAC_DEVICE_CLAIMS_INFO RPC structure.
//
// The PAC_DEVICE_CLAIMS_ INFO structure is a variable length buffer of the PAC that
// SHOULD<26> contain the client's marshaled claims blob. The PAC_DEVICE_CLAIMS_ INFO
// structure is placed directly after the Buffers array of the topmost PACTYPE structure
// (section 2.3), at the offset specified in the Offset field of the corresponding PAC_INFO_BUFFER
// structure (section 2.4) in the Buffers array. The ulType field of the corresponding
// PAC_INFO_BUFFER is set to 0x0000000F.
//
// The PAC_DEVICE_CLAIMS_INFO structure is defined as follows.
//
// The format of the PAC_DEVICE_CLAIMS_ INFO structure is defined as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Claims (variable)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type PACDeviceClaimsInfo struct {
	// Claims (variable): A variable-length CLAIMS_SET_METADATA structure ([MS-ADTS] section
	// 2.2.18.8) that contains claims.
	Claims *claims.ClaimsSetMetadata `idl:"name:Claims" json:"claims"`
}

func (o *PACDeviceClaimsInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACDeviceClaimsInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Claims != nil {
		_ptr_Claims := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Claims != nil {
				if err := o.Claims.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&claims.ClaimsSetMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Claims, _ptr_Claims); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACDeviceClaimsInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Claims := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Claims == nil {
			o.Claims = &claims.ClaimsSetMetadata{}
		}
		if err := o.Claims.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Claims := func(ptr interface{}) { o.Claims = *ptr.(**claims.ClaimsSetMetadata) }
	if err := w.ReadPointer(&o.Claims, _s_Claims, _ptr_Claims); err != nil {
		return err
	}
	return nil
}

// PACAttributesInfo structure represents PAC_ATTRIBUTES_INFO RPC structure.
type PACAttributesInfo struct {
	FlagsLength uint32 `idl:"name:FlagsLength" json:"flags_length"`
	Flags       uint32 `idl:"name:Flags" json:"flags"`
}

func (o *PACAttributesInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACAttributesInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.FlagsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *PACAttributesInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.FlagsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}
