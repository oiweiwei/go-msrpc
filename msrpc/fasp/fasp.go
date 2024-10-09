// The fasp package implements the FASP client protocol.
//
// # Introduction
//
// The Firewall and Advanced Security Protocol describes managing security policies
// on remote computers. The specific policies that this protocol manages are those of
// the firewall and advanced security components. The protocol allows the same functionality
// that is available locally; it can add, modify, delete, and enumerate policies. It
// can also enumerate security associations that can be generated between hosts after
// this policy is enforced.
//
// # Overview
//
// A host firewall is a software component that runs on host computers. It provides
// a layer of defense that can add depth to the collection of security measures when
// combined with other security measures, such as edge firewalls. Any threats that manage
// to get through the edge firewall, or those that are launched from within a corporate
// network, can still be defended against when host firewalls are used. Host firewalls
// are also useful in consumer scenarios in which there is, typically, no edge firewall
// to protect the home network.
//
// Internet Protocol Security (IPsec) is a host-based, policy-driven security solution
// for protecting the host from all network access. IPsec focuses on connection security,
// which includes authentication, integrity protection, and confidentiality (encryption)
// of communication.
//
// Because both IPsec and firewalls are host-based policy security technologies that
// operate in the network stack, they are managed together to avoid conflicts. Furthermore,
// firewall and connection security (IPsec) can interact, providing deeper and more
// effective filtering capabilities based on identities that are negotiated by IPsec
// as well as other IPsec state information. This document refers to this combined security
// solution as the firewall and advanced security components.
//
// Firewall and advanced security components can be governed by policy that is received
// from local users or from network-wide policy that is distributed by an administrator,
// or both. There is a need in managed environments for a network administrator to be
// able to monitor the policies in effect on hosts, assuming that hosts might have received
// policies from both sources.
//
// Network-wide policies are usually distributed by using Group Policy Objects (GPOs)
// that live on active directories of domains. However, some workgroups or networks
// might not have a domain infrastructure. Even in non-domain joined environments, the
// network administrator needs to be able to remotely manage the advanced firewall and
// IPsec policy of a host.
//
// Lastly, the network administrator might also be required to diagnose problems on
// the remote hosts. A common technique is to create temporary changes and then see
// if the changes fix the problem. This is the third scenario that warrants the capability
// to remotely administer host policies.
//
// The Firewall and Advanced Security Protocol is designed and used to address the three
// needs previously mentioned. That is, its purpose is to monitor and manage remote
// host policies. It can manage all the policies that an administrator can manage locally.
// It can also monitor the specific policies coming from the different sources or monitor
// them aggregated, that is, all together, to understand and predict expected behavior.
// Lastly, it can make temporary modifications on the remote host policy to test online
// fixes and see whether they are effective.
//
// The Firewall and Advanced Security Protocol is a client/server, RPC-based protocol.
// It consists of data types and methods. The data types are used to represent the different
// types of policy components that compose policy objects and policy configuration options.
// The methods are operations that are used to read and manage the different available
// policies. Therefore, the user can make method calls that pass new policy objects
// to be added to the policy, delete from the policy, or modify an existing object within
// the policy. The user can also call methods to retrieve all the policy objects of
// interest. The following illustration shows read and write operations and their message
// sequences.
package fasp

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

var (
	// import guard
	GoPackage = "fasp"
)

// ICMPCodeAny represents the FW_ICMP_CODE_ANY RPC constant
var ICMPCodeAny = 256

// IPProtocolAny represents the FW_IP_PROTOCOL_ANY RPC constant
var IPProtocolAny = 256

// ProfileConfigLogFileSizeMin represents the FW_PROFILE_CONFIG_LOG_FILE_SIZE_MIN RPC constant
var ProfileConfigLogFileSizeMin = 1

// ProfileConfigLogFileSizeMax represents the FW_PROFILE_CONFIG_LOG_FILE_SIZE_MAX RPC constant
var ProfileConfigLogFileSizeMax = 32767

// GlobalConfigCRLCheckMax represents the FW_GLOBAL_CONFIG_CRL_CHECK_MAX RPC constant
var GlobalConfigCRLCheckMax = 2

// GlobalConfigSAIdleTimeMax represents the FW_GLOBAL_CONFIG_SA_IDLE_TIME_MAX RPC constant
var GlobalConfigSAIdleTimeMax = 3600

// GlobalConfigSAIdleTimeMin represents the FW_GLOBAL_CONFIG_SA_IDLE_TIME_MIN RPC constant
var GlobalConfigSAIdleTimeMin = 300

// GlobalConfigPacketQueueValidationMask represents the FW_GLOBAL_CONFIG_PACKET_QUEUE_VALIDATION_MASK RPC constant
var GlobalConfigPacketQueueValidationMask = 3

// StoreType type represents FW_STORE_TYPE RPC enumeration.
//
// This data type defines enumerations used to identify store types.
type StoreType uint16

var (
	// FW_STORE_TYPE_INVALID:  This value is invalid and MUST NOT be used. It is defined
	// for simplicity in writing IDL definitions and code. This symbolic constant has a
	// value of zero.
	StoreTypeInvalid StoreType = 0
	// FW_STORE_TYPE_GP_RSOP:  This value identifies the store that contains all the policies
	// from the different Group Policy Objects (GPOs) that contain the networkwide policy.
	// This store is persisted in the registry. It is downloaded by the Group Policy component
	// (for more information, see [MS-GPREG]) and read by the firewall and advanced security
	// components; therefore, it is a read-only store. This symbolic constant has a value
	// of 1.
	StoreTypeGPRSOP StoreType = 1
	// FW_STORE_TYPE_LOCAL:  This value identifies the store that contains the local host
	// policy. This store is persisted in the registry by the firewall and advanced security
	// components; therefore, it is a read/write store. This symbolic constant has a value
	// of 2.
	StoreTypeLocal StoreType = 2
	// FW_STORE_TYPE_NOT_USED_VALUE_3:  This store is currently not used over the wire.
	// This symbolic constant has a value of 3.
	StoreTypeNotUsedValue3 StoreType = 3
	// FW_STORE_TYPE_NOT_USED_VALUE_4:  This store is currently not used over the wire.
	// This symbolic constant has a value of 4.
	StoreTypeNotUsedValue4 StoreType = 4
	// FW_STORE_TYPE_DYNAMIC:  This value identifies the store that contains the effective
	// policy, that is, the aggregated and merged policy from all policy sources. Policy
	// objects can be added and modified on this store, but they are not persisted and will
	// be lost the next time the firewall and advanced security components initialize. Policy
	// objects on this store can be modified only if they were originally added to this
	// store. This symbolic constant has a value of 5.
	StoreTypeDynamic StoreType = 5
	// FW_STORE_TYPE_GPO:  This value is not used on the wire. This symbolic constant has
	// a value of 6.
	StoreTypeGPO StoreType = 6
	// FW_STORE_TYPE_DEFAULTS:  This value identifies the store that contains the defaults
	// that the host operating system had out-of-box. This store is persisted in the registry.
	// It is written by the host operating system setup. It is read by the firewall and
	// advanced security components when it is instructed to go back to the default out-of-box
	// configuration; hence it is a read-only store. This symbolic constant has a value
	// of 7.
	StoreTypeDefaults StoreType = 7
	// FW_STORE_TYPE_NOT_USED_VALUE_8:  This store is currently not used over the wire.
	// This symbolic constant has a value of 8.
	StoreTypeNotUsedValue8 StoreType = 8
	// FW_STORE_TYPE_NOT_USED_VALUE_9:  This store is currently not used over the wire.
	// This symbolic constant has a value of 9.
	StoreTypeNotUsedValue9 StoreType = 9
	// FW_STORE_TYPE_NOT_USED_VALUE_10:  This store is currently not used over the wire.
	// This symbolic constant has a value of 10.
	StoreTypeNotUsedValue10 StoreType = 10
	// FW_STORE_TYPE_NOT_USED_VALUE_11:  This store is currently not used over the wire.
	// This symbolic constant has a value of 11.
	StoreTypeNotUsedValue11 StoreType = 11
	// FW_STORE_TYPE_MAX:  This value and values that exceed this value are not valid and
	// MUST NOT be used. This symbolic constant is defined for simplicity in writing IDL
	// definitions and code. It has a value of 8.
	StoreTypeMax StoreType = 12
)

func (o StoreType) String() string {
	switch o {
	case StoreTypeInvalid:
		return "StoreTypeInvalid"
	case StoreTypeGPRSOP:
		return "StoreTypeGPRSOP"
	case StoreTypeLocal:
		return "StoreTypeLocal"
	case StoreTypeNotUsedValue3:
		return "StoreTypeNotUsedValue3"
	case StoreTypeNotUsedValue4:
		return "StoreTypeNotUsedValue4"
	case StoreTypeDynamic:
		return "StoreTypeDynamic"
	case StoreTypeGPO:
		return "StoreTypeGPO"
	case StoreTypeDefaults:
		return "StoreTypeDefaults"
	case StoreTypeNotUsedValue8:
		return "StoreTypeNotUsedValue8"
	case StoreTypeNotUsedValue9:
		return "StoreTypeNotUsedValue9"
	case StoreTypeNotUsedValue10:
		return "StoreTypeNotUsedValue10"
	case StoreTypeNotUsedValue11:
		return "StoreTypeNotUsedValue11"
	case StoreTypeMax:
		return "StoreTypeMax"
	}
	return "Invalid"
}

// TransactionalState type represents FW_TRANSACTIONAL_STATE RPC enumeration.
type TransactionalState uint16

var (
	TransactionalStateNone    TransactionalState = 0
	TransactionalStateNoFlush TransactionalState = 1
	TransactionalStateMax     TransactionalState = 2
)

func (o TransactionalState) String() string {
	switch o {
	case TransactionalStateNone:
		return "TransactionalStateNone"
	case TransactionalStateNoFlush:
		return "TransactionalStateNoFlush"
	case TransactionalStateMax:
		return "TransactionalStateMax"
	}
	return "Invalid"
}

// ProfileType type represents FW_PROFILE_TYPE RPC enumeration.
//
// This data type defines the enumerations that are used to identify profile types.
// The enumeration values are bitmasks. Implementations MUST support using a single
// bitmask value and MUST support a combination of bitmask values. Valid combinations
// of bitmask values are all possible combinations using FW_PROFILE_TYPE_DOMAIN, FW_PROFILE_TYPE_PRIVATE,
// FW_PROFILE_TYPE_PUBLIC, and FW_PROFILE_TYPE_ALL. A profile is a set of networks to
// which a firewall policy might apply.
type ProfileType uint32

var (
	// FW_PROFILE_TYPE_INVALID:  This value is invalid and MUST NOT be used. It is defined
	// for simplicity in writing IDL definitions and code.
	ProfileTypeInvalid ProfileType = 0
	// FW_PROFILE_TYPE_DOMAIN:  This value represents the profile for networks that are
	// connected to domains.
	ProfileTypeDomain ProfileType = 1
	// FW_PROFILE_TYPE_STANDARD:  This value represents the standard profile for networks.
	// These networks are classified as private by the administrators in the server host.
	// The classification happens the first time the host connects to the network. Usually
	// these networks are behind Network Address Translation (NAT) devices, routers, and
	// other edge devices, and they are in a private location, such as a home or an office.
	ProfileTypeStandard ProfileType = 2
	// FW_PROFILE_TYPE_PRIVATE:  This value represents the profile for private networks,
	// which is represented by the same value as that used for FW_PROFILE_TYPE_STANDARD.
	ProfileTypePrivate ProfileType = 2
	// FW_PROFILE_TYPE_PUBLIC:  This value represents the profile for public networks.
	// These networks are classified as public by the administrators in the server host.
	// The classification happens the first time the host connects to the network. Usually
	// these networks are those at airports, coffee shops, and other public places where
	// the peers in the network or the network administrator are not trusted.
	ProfileTypePublic ProfileType = 4
	// FW_PROFILE_TYPE_ALL:  This value represents all these network sets and any future
	// network sets.
	ProfileTypeAll ProfileType = 2147483647
	// FW_PROFILE_TYPE_CURRENT:  This value represents the current profiles to which the
	// firewall and advanced security components determine the host is connected at the
	// moment of the call. This value can be specified only in method calls, and it cannot
	// be combined with other flags.
	ProfileTypeCurrent ProfileType = 2147483648
	// FW_PROFILE_TYPE_NONE:  This value represents no profile and is invalid. It is defined
	// for simplicity in writing IDL definitions and code. This and greater values MUST
	// NOT be used.
	ProfileTypeNone ProfileType = 2147483649
)

func (o ProfileType) String() string {
	switch o {
	case ProfileTypeInvalid:
		return "ProfileTypeInvalid"
	case ProfileTypeDomain:
		return "ProfileTypeDomain"
	case ProfileTypeStandard:
		return "ProfileTypeStandard"
	case ProfileTypePrivate:
		return "ProfileTypePrivate"
	case ProfileTypePublic:
		return "ProfileTypePublic"
	case ProfileTypeAll:
		return "ProfileTypeAll"
	case ProfileTypeCurrent:
		return "ProfileTypeCurrent"
	case ProfileTypeNone:
		return "ProfileTypeNone"
	}
	return "Invalid"
}

// PolicyAccessRight type represents FW_POLICY_ACCESS_RIGHT RPC enumeration.
//
// This enumeration defines access rights for the policy elements that can be accessed
// using the Firewall and Advanced Security Protocol. The values are not bitmasks and
// SHOULD NOT be used in bitwise OR operations.
type PolicyAccessRight uint16

var (
	// FW_POLICY_ACCESS_RIGHT_INVALID:  This value is invalid and MUST NOT be used. It
	// is defined for simplicity in writing IDL definitions and code. This symbolic constant
	// has a value of zero.
	PolicyAccessRightInvalid PolicyAccessRight = 0
	// FW_POLICY_ACCESS_RIGHT_READ:  This value represents a read-only access right. This
	// symbolic constant has a value of 1.
	PolicyAccessRightRead PolicyAccessRight = 1
	// FW_POLICY_ACCESS_RIGHT_READ_WRITE:  This value represents a read and write access
	// right. This symbolic constant has a value of 2.
	PolicyAccessRightReadWrite PolicyAccessRight = 2
	// FW_POLICY_ACCESS_RIGHT_MAX:  This value and values that exceed this value are not
	// valid and MUST NOT be used. This symbolic constant is defined for simplicity in writing
	// IDL definitions and code. It has a value of 3.
	PolicyAccessRightMax PolicyAccessRight = 3
)

func (o PolicyAccessRight) String() string {
	switch o {
	case PolicyAccessRightInvalid:
		return "PolicyAccessRightInvalid"
	case PolicyAccessRightRead:
		return "PolicyAccessRightRead"
	case PolicyAccessRightReadWrite:
		return "PolicyAccessRightReadWrite"
	case PolicyAccessRightMax:
		return "PolicyAccessRightMax"
	}
	return "Invalid"
}

// PolicyStoreFlags type represents FW_POLICY_STORE_FLAGS RPC enumeration.
type PolicyStoreFlags uint16

var (
	PolicyStoreFlagsNone                         PolicyStoreFlags = 0
	PolicyStoreFlagsDeleteDynamicRulesAfterClose PolicyStoreFlags = 1
	PolicyStoreFlagsOpenGPCache                  PolicyStoreFlags = 2
	PolicyStoreFlagsUseGPCache                   PolicyStoreFlags = 4
	PolicyStoreFlagsSaveGPCache                  PolicyStoreFlags = 8
	PolicyStoreFlagsNotUsedValue16               PolicyStoreFlags = 16
	PolicyStoreFlagsMax                          PolicyStoreFlags = 32
)

func (o PolicyStoreFlags) String() string {
	switch o {
	case PolicyStoreFlagsNone:
		return "PolicyStoreFlagsNone"
	case PolicyStoreFlagsDeleteDynamicRulesAfterClose:
		return "PolicyStoreFlagsDeleteDynamicRulesAfterClose"
	case PolicyStoreFlagsOpenGPCache:
		return "PolicyStoreFlagsOpenGPCache"
	case PolicyStoreFlagsUseGPCache:
		return "PolicyStoreFlagsUseGPCache"
	case PolicyStoreFlagsSaveGPCache:
		return "PolicyStoreFlagsSaveGPCache"
	case PolicyStoreFlagsNotUsedValue16:
		return "PolicyStoreFlagsNotUsedValue16"
	case PolicyStoreFlagsMax:
		return "PolicyStoreFlagsMax"
	}
	return "Invalid"
}

// IPv4Subnet structure represents FW_IPV4_SUBNET RPC structure.
//
// This structure defines IPv4 subnets. It is used in policy rules.
type IPv4Subnet struct {
	// dwAddress:   This field represents the IPv4 address.
	Address uint32 `idl:"name:dwAddress" json:"address"`
	// dwSubNetMask:  This field contains the subnet mask in host network order. If it contains
	// ones, they MUST be contiguous and shifted to the most significant bits.
	SubnetMask uint32 `idl:"name:dwSubNetMask" json:"subnet_mask"`
}

func (o *IPv4Subnet) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv4Subnet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Address); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	return nil
}
func (o *IPv4Subnet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Address); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	return nil
}

// IPv4SubnetList structure represents FW_IPV4_SUBNET_LIST RPC structure.
//
// This structure is used to contain a number of FW_IPV4_SUBNET elements.
type IPv4SubnetList struct {
	// dwNumEntries:  This field specifies the number of subnets that the structure contains.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// pSubNets:   A pointer to an array of FW_IPV4_SUBNET elements. The number of elements
	// is given by dwNumEntries.
	Subnets []*IPv4Subnet `idl:"name:pSubNets;size_is:(dwNumEntries)" json:"subnets"`
}

func (o *IPv4SubnetList) xxx_PreparePayload(ctx context.Context) error {
	if o.Subnets != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Subnets))
	}
	if o.EntriesLength > uint32(10000) {
		return fmt.Errorf("EntriesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv4SubnetList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.Subnets != nil || o.EntriesLength > 0 {
		_ptr_pSubNets := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Subnets {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Subnets[i1] != nil {
					if err := o.Subnets[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&IPv4Subnet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Subnets); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&IPv4Subnet{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Subnets, _ptr_pSubNets); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv4SubnetList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_pSubNets := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Subnets", sizeInfo[0])
		}
		o.Subnets = make([]*IPv4Subnet, sizeInfo[0])
		for i1 := range o.Subnets {
			i1 := i1
			if o.Subnets[i1] == nil {
				o.Subnets[i1] = &IPv4Subnet{}
			}
			if err := o.Subnets[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pSubNets := func(ptr interface{}) { o.Subnets = *ptr.(*[]*IPv4Subnet) }
	if err := w.ReadPointer(&o.Subnets, _s_pSubNets, _ptr_pSubNets); err != nil {
		return err
	}
	return nil
}

// IPv6Subnet structure represents FW_IPV6_SUBNET RPC structure.
//
// This structure represents an IPv6 subnet.
type IPv6Subnet struct {
	// Address:  This field contains a 16-octet IPv6 address.
	Address []byte `idl:"name:Address" json:"address"`
	// dwNumPrefixBits:  This field contains the number of more-significant bits that represent
	// the IPv6 subnet.
	PrefixBitsLength uint32 `idl:"name:dwNumPrefixBits" json:"prefix_bits_length"`
}

func (o *IPv6Subnet) xxx_PreparePayload(ctx context.Context) error {
	if o.PrefixBitsLength > uint32(128) {
		return fmt.Errorf("PrefixBitsLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv6Subnet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PrefixBitsLength); err != nil {
		return err
	}
	return nil
}
func (o *IPv6Subnet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.Address = make([]byte, 16)
	for i1 := range o.Address {
		i1 := i1
		if err := w.ReadData(&o.Address[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.PrefixBitsLength); err != nil {
		return err
	}
	return nil
}

// IPv6SubnetList structure represents FW_IPV6_SUBNET_LIST RPC structure.
//
// This structure is used to contain a number of FW_IPV6_SUBNET elements.
type IPv6SubnetList struct {
	// dwNumEntries:  This field specifies the number of subnets that the structure contains.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// pSubNets:  A pointer to an array of FW_IPV6_SUBNET elements. The number of elements
	// is given by dwNumEntries.
	Subnets []*IPv6Subnet `idl:"name:pSubNets;size_is:(dwNumEntries)" json:"subnets"`
}

func (o *IPv6SubnetList) xxx_PreparePayload(ctx context.Context) error {
	if o.Subnets != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Subnets))
	}
	if o.EntriesLength > uint32(10000) {
		return fmt.Errorf("EntriesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv6SubnetList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.Subnets != nil || o.EntriesLength > 0 {
		_ptr_pSubNets := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Subnets {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Subnets[i1] != nil {
					if err := o.Subnets[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&IPv6Subnet{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Subnets); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&IPv6Subnet{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Subnets, _ptr_pSubNets); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv6SubnetList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_pSubNets := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Subnets", sizeInfo[0])
		}
		o.Subnets = make([]*IPv6Subnet, sizeInfo[0])
		for i1 := range o.Subnets {
			i1 := i1
			if o.Subnets[i1] == nil {
				o.Subnets[i1] = &IPv6Subnet{}
			}
			if err := o.Subnets[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pSubNets := func(ptr interface{}) { o.Subnets = *ptr.(*[]*IPv6Subnet) }
	if err := w.ReadPointer(&o.Subnets, _s_pSubNets, _ptr_pSubNets); err != nil {
		return err
	}
	return nil
}

// IPv4AddressRange structure represents FW_IPV4_ADDRESS_RANGE RPC structure.
//
// This structure represents a range of IPv4 addresses within the IPv4 address space.
type IPv4AddressRange struct {
	// dwBegin:  The first IPv4 address of the range in the IPv4 address space defined by
	// this structure. The address is included in the range.
	Begin uint32 `idl:"name:dwBegin" json:"begin"`
	// dwEnd:  The last IPv4 address of the range in the IPv4 address space defined by this
	// structure. The address is included in the range.
	//
	// Valid FW_IPV4_ADDRESS_RANGE structures MUST have a dwBegin value less than or equal
	// to the dwEnd value. Structures with dwBegin equal to dwEnd represent a single IPv4
	// address.
	End uint32 `idl:"name:dwEnd" json:"end"`
}

func (o *IPv4AddressRange) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv4AddressRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Begin); err != nil {
		return err
	}
	if err := w.WriteData(o.End); err != nil {
		return err
	}
	return nil
}
func (o *IPv4AddressRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Begin); err != nil {
		return err
	}
	if err := w.ReadData(&o.End); err != nil {
		return err
	}
	return nil
}

// IPv6AddressRange structure represents FW_IPV6_ADDRESS_RANGE RPC structure.
//
// This structure represents a range of IPv6 addresses within the IPv6 address space.
type IPv6AddressRange struct {
	// Begin:  A 16-octet array containing the first IPv6 address of the range in the IPv6
	// address range defined by this structure.
	Begin []byte `idl:"name:Begin" json:"begin"`
	// End:  A 16-octet array containing the last IPv6 address of the range in the IPv6
	// address range defined by this structure.
	//
	// Valid FW_IPV6_ADDRESS_RANGE structures MUST have a Begin value less than or equal
	// to the End value. Structures with Begin equal to End represent a single IPv6 address.
	// Begin and End MUST NOT contain either an unspecified or a loopback address.
	//
	// Begin and End are in network order.
	End []byte `idl:"name:End" json:"end"`
}

func (o *IPv6AddressRange) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv6AddressRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Begin {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Begin[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Begin); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.End {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.End[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.End); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv6AddressRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Begin = make([]byte, 16)
	for i1 := range o.Begin {
		i1 := i1
		if err := w.ReadData(&o.Begin[i1]); err != nil {
			return err
		}
	}
	o.End = make([]byte, 16)
	for i1 := range o.End {
		i1 := i1
		if err := w.ReadData(&o.End[i1]); err != nil {
			return err
		}
	}
	return nil
}

// IPv4RangeList structure represents FW_IPV4_RANGE_LIST RPC structure.
//
// This structure is used to contain a number of FW_IPV4_ADDRESS_RANGE elements.
type IPv4RangeList struct {
	// dwNumEntries:  This field specifies the number of IPv4 address ranges that the structure
	// contains.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// pRanges:  A pointer to an array of FW_IPV4_ADDRESS_RANGE elements. The number of
	// elements is given by dwNumEntries.
	Ranges []*IPv4AddressRange `idl:"name:pRanges;size_is:(dwNumEntries)" json:"ranges"`
}

func (o *IPv4RangeList) xxx_PreparePayload(ctx context.Context) error {
	if o.Ranges != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Ranges))
	}
	if o.EntriesLength > uint32(10000) {
		return fmt.Errorf("EntriesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv4RangeList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.Ranges != nil || o.EntriesLength > 0 {
		_ptr_pRanges := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Ranges {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Ranges[i1] != nil {
					if err := o.Ranges[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&IPv4AddressRange{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Ranges); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&IPv4AddressRange{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Ranges, _ptr_pRanges); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv4RangeList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_pRanges := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Ranges", sizeInfo[0])
		}
		o.Ranges = make([]*IPv4AddressRange, sizeInfo[0])
		for i1 := range o.Ranges {
			i1 := i1
			if o.Ranges[i1] == nil {
				o.Ranges[i1] = &IPv4AddressRange{}
			}
			if err := o.Ranges[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pRanges := func(ptr interface{}) { o.Ranges = *ptr.(*[]*IPv4AddressRange) }
	if err := w.ReadPointer(&o.Ranges, _s_pRanges, _ptr_pRanges); err != nil {
		return err
	}
	return nil
}

// IPv6RangeList structure represents FW_IPV6_RANGE_LIST RPC structure.
//
// This structure is used to contain a number of FW_IPV6_ADDRESS_RANGE elements.
type IPv6RangeList struct {
	// dwNumEntries:  This field specifies the number of IPv6 address ranges that the structure
	// contains.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// pRanges:  A pointer to an array of FW_IPV6_ADDRESS_RANGE elements. The number of
	// elements is given by dwNumEntries.
	Ranges []*IPv6AddressRange `idl:"name:pRanges;size_is:(dwNumEntries)" json:"ranges"`
}

func (o *IPv6RangeList) xxx_PreparePayload(ctx context.Context) error {
	if o.Ranges != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Ranges))
	}
	if o.EntriesLength > uint32(10000) {
		return fmt.Errorf("EntriesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv6RangeList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.Ranges != nil || o.EntriesLength > 0 {
		_ptr_pRanges := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Ranges {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Ranges[i1] != nil {
					if err := o.Ranges[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&IPv6AddressRange{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Ranges); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&IPv6AddressRange{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Ranges, _ptr_pRanges); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv6RangeList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_pRanges := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Ranges", sizeInfo[0])
		}
		o.Ranges = make([]*IPv6AddressRange, sizeInfo[0])
		for i1 := range o.Ranges {
			i1 := i1
			if o.Ranges[i1] == nil {
				o.Ranges[i1] = &IPv6AddressRange{}
			}
			if err := o.Ranges[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pRanges := func(ptr interface{}) { o.Ranges = *ptr.(*[]*IPv6AddressRange) }
	if err := w.ReadPointer(&o.Ranges, _s_pRanges, _ptr_pRanges); err != nil {
		return err
	}
	return nil
}

// PortRange structure represents FW_PORT_RANGE RPC structure.
//
// This structure represents a range of ports. Ports are 16-bit unsigned values used
// in TCP and UDP protocols.
type PortRange struct {
	// wBegin:  This field specifies the first port included in the range defined.
	Begin uint16 `idl:"name:wBegin" json:"begin"`
	// wEnd:  This field specifies the last port included in the range defined.
	//
	// Valid FW_PORT_RANGE structures MUST have a wBegin value less than or equal to the
	// wEnd value. In this protocol, wBegin is equal to wEnd.
	End uint16 `idl:"name:wEnd" json:"end"`
}

func (o *PortRange) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PortRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.Begin); err != nil {
		return err
	}
	if err := w.WriteData(o.End); err != nil {
		return err
	}
	return nil
}
func (o *PortRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Begin); err != nil {
		return err
	}
	if err := w.ReadData(&o.End); err != nil {
		return err
	}
	return nil
}

// PortRangeList structure represents FW_PORT_RANGE_LIST RPC structure.
//
// This structure is used to contain a number of FW_PORT_RANGE elements.
type PortRangeList struct {
	// dwNumEntries:  This field specifies the number of port ranges that the structure
	// contains.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// pPorts:  A pointer to an array of FW_PORT_RANGE elements. The number of elements
	// is given as dwNumEntries.
	Ports []*PortRange `idl:"name:pPorts;size_is:(dwNumEntries)" json:"ports"`
}

func (o *PortRangeList) xxx_PreparePayload(ctx context.Context) error {
	if o.Ports != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Ports))
	}
	if o.EntriesLength > uint32(10000) {
		return fmt.Errorf("EntriesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PortRangeList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.Ports != nil || o.EntriesLength > 0 {
		_ptr_pPorts := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Ports {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Ports[i1] != nil {
					if err := o.Ports[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PortRange{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Ports); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&PortRange{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Ports, _ptr_pPorts); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PortRangeList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_pPorts := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Ports", sizeInfo[0])
		}
		o.Ports = make([]*PortRange, sizeInfo[0])
		for i1 := range o.Ports {
			i1 := i1
			if o.Ports[i1] == nil {
				o.Ports[i1] = &PortRange{}
			}
			if err := o.Ports[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pPorts := func(ptr interface{}) { o.Ports = *ptr.(*[]*PortRange) }
	if err := w.ReadPointer(&o.Ports, _s_pPorts, _ptr_pPorts); err != nil {
		return err
	}
	return nil
}

// PortKeyword type represents FW_PORT_KEYWORD RPC enumeration.
//
// This enumeration identifies (with bitmask flags) the ports used by specific well-known
// protocols.<2> The ports corresponding to these keywords change dynamically and are
// tracked by the PortsInUse object (see section 3.1.1). All the flags supported by
// a given schema version can be combined, except for the restrictions placed on the
// wPortKeywords field as stated in FW_RULE (section 2.2.37) and FW_CS_RULE (section
// 2.2.55).
type PortKeyword uint16

var (
	// FW_PORT_KEYWORD_NONE:  Specifies that no port keywords are used.
	PortKeywordNone PortKeyword = 0
	// FW_PORT_KEYWORD_DYNAMIC_RPC_PORTS:  Represents all ports in the PortsInUse collection
	// where IsDynamicRPC is true.
	PortKeywordDynamicRPCPorts PortKeyword = 1
	// FW_PORT_KEYWORD_RPC_EP:  Represents all ports in the PortsInUse collection where
	// IsRPCEndpointMapper is true.
	PortKeywordRPCEP PortKeyword = 2
	// FW_PORT_KEYWORD_TEREDO_PORT:  Represents all ports in the PortsInUse collection
	// where IsTeredo is true.
	PortKeywordTeredoPort PortKeyword = 4
	// FW_PORT_KEYWORD_IP_TLS_IN:  Represents all ports in the PortsInUse collection where
	// IsIPTLSIn is true. For schema versions 0x0200 and 0x0201, this value is invalid and
	// MUST NOT be used. This symbolic constant has a value of 0x08.
	PortKeywordIPTLSIn PortKeyword = 8
	// FW_PORT_KEYWORD_IP_TLS_OUT:  Represents all ports in the PortsInUse collection where
	// IsIPTLSOut is true. For schema versions 0x0200 and 0x0201, this value is invalid
	// and MUST NOT be used. This symbolic constant has a value of 0x10.
	PortKeywordIPTLSOut PortKeyword = 16
	// FW_PORT_KEYWORD_DHCP:  Represents all ports in the PortsInUse collection where IsDHCPClient
	// is true. For schema versions 0x0200, 0x0201, and 0x020A, this value is invalid and
	// MUST NOT be used. This symbolic constant has a value of 0x20.
	PortKeywordDHCP PortKeyword = 32
	// FW_PORT_KEYWORD_PLAYTO_DISCOVERY:  Represents all ports in the PortsInUse collection
	// where IsPlayToDiscovery is true. For schema versions 0x0200, 0x0201, and 0x020A,
	// this value is invalid and MUST NOT be used. This symbolic constant has a value of
	// 0x40.
	PortKeywordPlayToDiscovery PortKeyword = 64
	// FW_PORT_KEYWORD_MDNS:  Represents all ports in the PortsInUse collection where IsMDNS
	// is true. For schema versions 0x0200, 0x0201, 0x020A, and 0x0214, this value is invalid
	// and MUST NOT be used. This symbolic constant has a value of 0x80.
	PortKeywordMDNS PortKeyword = 128
	// FW_PORT_KEYWORD_CORTANA_OUT:  Represents all ports in the PortsInUse collection
	// where IsCortanaOut is true. For schema versions 0x0200, 0x0201, 0x020A, 0x0214, 0x0216,
	// 0x0218, and 0x0219, this value is invalid and MUST NOT be used. This symbolic constant
	// has a value of 0x100.
	PortKeywordCortanaOut PortKeyword = 256
	// FW_PORT_KEYWORD_PROXIMAL_TCP_CDP:  Represents all ports in the PortsInUse collection
	// where IsProximalTCPCDP is true. For schema versions 0x0200, 0x0201, 0x020A, 0x0214,
	// 0x0216, 0x0218, 0x0219, 0x021A, 0x021B, and 0x021C, this value is invalid and MUST
	// NOT be used. This symbolic constant has a value of 0x200.
	PortKeywordProximalTCPCDP PortKeyword = 512
	// FW_PORT_KEYWORD_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 0x400.<3>
	PortKeywordMax PortKeyword = 1024
	// FW_PORT_KEYWORD_MAX_V2_1:  This value and values that exceed this value are not
	// valid and MUST NOT be used by servers and clients with schema version 0x0201 and
	// earlier. It is defined for simplicity in writing IDL definitions and code. This symbolic
	// constant has a value of 0x08.
	PortKeywordMaxV21 PortKeyword = 8
	// FW_PORT_KEYWORD_MAX_V2_10:  This value and values that exceed this value are not
	// valid and MUST NOT be used by servers and clients with schema version 0x020A and
	// earlier. It is defined for simplicity in writing IDL definitions and code. This symbolic
	// constant has a value of 0x20.
	PortKeywordMaxV210 PortKeyword = 32
	// FW_PORT_KEYWORD_MAX_V2_20:  This value and values that exceed this value are not
	// valid and MUST NOT be used by servers and clients with schema version 0x0214 and
	// earlier. It is defined for simplicity in writing IDL definitions and code. This symbolic
	// constant has a value of 0x80.
	PortKeywordMaxV220 PortKeyword = 128
	// FW_PORT_KEYWORD_MAX_V2_24:  This value and values that exceed this value are not
	// valid and MUST NOT be used by servers and clients with schema version 0x0219 and
	// earlier. It is defined for simplicity in writing IDL definitions and code. This symbolic
	// constant has a value of 0x100.
	PortKeywordMaxV224 PortKeyword = 256
	// FW_PORT_KEYWORD_MAX_V2_25:  This value and values that exceed this value are not
	// valid and MUST NOT be used by servers and clients with schema version 0x021B and
	// earlier. It is defined for simplicity in writing IDL definitions and code. This symbolic
	// constant has a value of 0x200.
	PortKeywordMaxV225 PortKeyword = 512
)

func (o PortKeyword) String() string {
	switch o {
	case PortKeywordNone:
		return "PortKeywordNone"
	case PortKeywordDynamicRPCPorts:
		return "PortKeywordDynamicRPCPorts"
	case PortKeywordRPCEP:
		return "PortKeywordRPCEP"
	case PortKeywordTeredoPort:
		return "PortKeywordTeredoPort"
	case PortKeywordIPTLSIn:
		return "PortKeywordIPTLSIn"
	case PortKeywordIPTLSOut:
		return "PortKeywordIPTLSOut"
	case PortKeywordDHCP:
		return "PortKeywordDHCP"
	case PortKeywordPlayToDiscovery:
		return "PortKeywordPlayToDiscovery"
	case PortKeywordMDNS:
		return "PortKeywordMDNS"
	case PortKeywordCortanaOut:
		return "PortKeywordCortanaOut"
	case PortKeywordProximalTCPCDP:
		return "PortKeywordProximalTCPCDP"
	case PortKeywordMax:
		return "PortKeywordMax"
	case PortKeywordMaxV21:
		return "PortKeywordMaxV21"
	case PortKeywordMaxV210:
		return "PortKeywordMaxV210"
	case PortKeywordMaxV220:
		return "PortKeywordMaxV220"
	case PortKeywordMaxV224:
		return "PortKeywordMaxV224"
	case PortKeywordMaxV225:
		return "PortKeywordMaxV225"
	}
	return "Invalid"
}

// Ports structure represents FW_PORTS RPC structure.
//
// This structure contains the ports represented statically through FW_PORT_RANGE structures
// or symbolically through FW_PORT_KEYWORD enumeration values.
type Ports struct {
	// wPortKeywords:  This field is a combination of FW_PORT_KEYWORD values.
	PortKeywords uint16 `idl:"name:wPortKeywords" json:"port_keywords"`
	// Ports:  This field is a list of specifically defined ports.
	Ports *PortRangeList `idl:"name:Ports" json:"ports"`
}

func (o *Ports) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Ports) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PortKeywords); err != nil {
		return err
	}
	if o.Ports != nil {
		if err := o.Ports.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PortRangeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Ports) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PortKeywords); err != nil {
		return err
	}
	if o.Ports == nil {
		o.Ports = &PortRangeList{}
	}
	if err := o.Ports.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ICMPTypeCode structure represents FW_ICMP_TYPE_CODE RPC structure.
//
// This data type defines ICMP (internet control message protocol with protocol numbers
// assigned in [IANA-PROTO-NUM]) message types and codes. It specifies an ICMP type
// and either its specific code or all codes for that type.
type ICMPTypeCode struct {
	// bType:  This field specifies the ICMP type.
	Type uint8 `idl:"name:bType" json:"type"`
	// wCode:  This field specifies the ICMP code.
	//
	// The wCode field MUST contain values between 0x0000 and 0x0100. When wCode contains
	// 0x100, it expresses any ICMP code belonging to the corresponding ICMP type. When
	// wCode contains values in the range 0 to 0x00FF, it expresses a specific ICMP code.
	//
	// All valid ICMP type and code combinations are valid, even those not currently assigned
	// for a specific use.
	Code uint16 `idl:"name:wCode" json:"code"`
}

func (o *ICMPTypeCode) xxx_PreparePayload(ctx context.Context) error {
	if o.Code > uint16(256) {
		return fmt.Errorf("Code is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ICMPTypeCode) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Code); err != nil {
		return err
	}
	return nil
}
func (o *ICMPTypeCode) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Code); err != nil {
		return err
	}
	return nil
}

// ICMPTypeCodeList structure represents FW_ICMP_TYPE_CODE_LIST RPC structure.
//
// This structure is used to contain a number of FW_ICMP_TYPE_CODE elements.
type ICMPTypeCodeList struct {
	// dwNumEntries:  This field specifies the number of FW_ICMP_TYPE_CODE elements that
	// the structure contains.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// pEntries:  A pointer to an array of FW_ICMP_TYPE_CODE elements. The number of elements
	// is given by dwNumEntries.
	Entries []*ICMPTypeCode `idl:"name:pEntries;size_is:(dwNumEntries)" json:"entries"`
}

func (o *ICMPTypeCodeList) xxx_PreparePayload(ctx context.Context) error {
	if o.Entries != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Entries))
	}
	if o.EntriesLength > uint32(10000) {
		return fmt.Errorf("EntriesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ICMPTypeCodeList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.Entries != nil || o.EntriesLength > 0 {
		_ptr_pEntries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
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
					if err := o.Entries[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ICMPTypeCode{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ICMPTypeCode{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Entries, _ptr_pEntries); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ICMPTypeCodeList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_pEntries := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*ICMPTypeCode, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			if o.Entries[i1] == nil {
				o.Entries[i1] = &ICMPTypeCode{}
			}
			if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pEntries := func(ptr interface{}) { o.Entries = *ptr.(*[]*ICMPTypeCode) }
	if err := w.ReadPointer(&o.Entries, _s_pEntries, _ptr_pEntries); err != nil {
		return err
	}
	return nil
}

// InterfaceLUIDs structure represents FW_INTERFACE_LUIDS RPC structure.
//
// This structure is used to contain locally unique identifier (LUID) values that uniquely
// represent single network adapters (NICs) within a specific computer.
type InterfaceLUIDs struct {
	// dwNumLUIDs:  This field specifies the number of interface LUIDs that the structure
	// contains.
	LUIDsLength uint32 `idl:"name:dwNumLUIDs" json:"luids_length"`
	// pLUIDs:  A pointer to an array of GUID elements. The number of elements is given
	// by dwNumLUIDs. The GUID data type is specified in [MS-DTYP].
	LUIDs []*dtyp.GUID `idl:"name:pLUIDs;size_is:(dwNumLUIDs)" json:"luids"`
}

func (o *InterfaceLUIDs) xxx_PreparePayload(ctx context.Context) error {
	if o.LUIDs != nil && o.LUIDsLength == 0 {
		o.LUIDsLength = uint32(len(o.LUIDs))
	}
	if o.LUIDsLength > uint32(10000) {
		return fmt.Errorf("LUIDsLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceLUIDs) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.LUIDsLength); err != nil {
		return err
	}
	if o.LUIDs != nil || o.LUIDsLength > 0 {
		_ptr_pLUIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.LUIDsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.LUIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.LUIDs[i1] != nil {
					if err := o.LUIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.LUIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LUIDs, _ptr_pLUIDs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceLUIDs) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.LUIDsLength); err != nil {
		return err
	}
	_ptr_pLUIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.LUIDsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.LUIDsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.LUIDs", sizeInfo[0])
		}
		o.LUIDs = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.LUIDs {
			i1 := i1
			if o.LUIDs[i1] == nil {
				o.LUIDs[i1] = &dtyp.GUID{}
			}
			if err := o.LUIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pLUIDs := func(ptr interface{}) { o.LUIDs = *ptr.(*[]*dtyp.GUID) }
	if err := w.ReadPointer(&o.LUIDs, _s_pLUIDs, _ptr_pLUIDs); err != nil {
		return err
	}
	return nil
}

// Direction type represents FW_DIRECTION RPC enumeration.
//
// This enumeration represents the direction of network traffic flow.
type Direction uint16

var (
	// FW_DIR_INVALID:  This is an invalid value, and it MUST NOT be used. It is defined
	// for simplicity in writing IDL definitions and code. This symbolic constant has a
	// value of zero.
	DirectionInvalid Direction = 0
	// FW_DIR_IN:  Specifies an inbound network traffic flow. These are flows that are
	// initiated by a remote machine toward the local machine. This symbolic constant has
	// a value of 1.
	DirectionIn Direction = 1
	// FW_DIR_OUT:  Specifies an outbound network traffic flow. These are flows that are
	// initiated by the local machine toward a remote machine. This symbolic constant has
	// a value of 2.
	DirectionOut Direction = 2
	// FW_DIR_MAX:  This value and values that exceed this value are not valid and MUST
	// NOT be used. This symbolic constant is defined for simplicity in writing IDL definitions
	// and code. It has a value of 3.
	DirectionMax Direction = 3
)

func (o Direction) String() string {
	switch o {
	case DirectionInvalid:
		return "DirectionInvalid"
	case DirectionIn:
		return "DirectionIn"
	case DirectionOut:
		return "DirectionOut"
	case DirectionMax:
		return "DirectionMax"
	}
	return "Invalid"
}

// InterfaceType type represents FW_INTERFACE_TYPE RPC enumeration.
//
// This enumeration is used to represent types of network adapters (NICs) in a specific
// machine. Each type might have one or more network adapters.
type InterfaceType uint16

var (
	// FW_INTERFACE_TYPE_ALL:  Represents all types of network adapters (NICs). The following
	// types fall into this type.
	InterfaceTypeAll InterfaceType = 0
	// FW_INTERFACE_TYPE_LAN:  Represents network adapters (NICs) that use wired network
	// physical layers such as Ethernet.
	InterfaceTypeLAN InterfaceType = 1
	// FW_INTERFACE_TYPE_WIRELESS:  Represents network adapters that use the wireless 802
	// network physical layer.
	InterfaceTypeWireless InterfaceType = 2
	// FW_INTERFACE_TYPE_REMOTE_ACCESS:  Represents network adapters that use VPN connections.
	InterfaceTypeRemoteAccess InterfaceType = 4
	InterfaceTypeMobileBband  InterfaceType = 8
	// FW_INTERFACE_TYPE_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 0x0008.
	InterfaceTypeMax     InterfaceType = 16
	InterfaceTypeMaxV223 InterfaceType = 8
)

func (o InterfaceType) String() string {
	switch o {
	case InterfaceTypeAll:
		return "InterfaceTypeAll"
	case InterfaceTypeLAN:
		return "InterfaceTypeLAN"
	case InterfaceTypeWireless:
		return "InterfaceTypeWireless"
	case InterfaceTypeRemoteAccess:
		return "InterfaceTypeRemoteAccess"
	case InterfaceTypeMobileBband:
		return "InterfaceTypeMobileBband"
	case InterfaceTypeMax:
		return "InterfaceTypeMax"
	case InterfaceTypeMaxV223:
		return "InterfaceTypeMaxV223"
	}
	return "Invalid"
}

// AddressKeyword type represents FW_ADDRESS_KEYWORD RPC enumeration.
//
// This enumeration is used to represent specific address types. As specified in the
// following descriptions, these address types can change dynamically.
type AddressKeyword uint16

var (
	// FW_ADDRESS_KEYWORD_NONE:  Specifies that no specific keyword is used.
	AddressKeywordNone AddressKeyword = 0
	// FW_ADDRESS_KEYWORD_LOCAL_SUBNET:  Represents the collection of addresses that are
	// currently within the local subnet of the computer.
	AddressKeywordLocalSubnet AddressKeyword = 1
	// FW_ADDRESS_KEYWORD_DNS:  Represents the collection of addresses of the current DNS
	// servers.
	AddressKeywordDNS AddressKeyword = 2
	// FW_ADDRESS_KEYWORD_DHCP:  Represents the collection of addresses of the current
	// DHCP servers.
	AddressKeywordDHCP AddressKeyword = 4
	// FW_ADDRESS_KEYWORD_WINS:  Represents the collection of addresses of the current
	// WINS servers.
	AddressKeywordWINS AddressKeyword = 8
	// FW_ADDRESS_KEYWORD_DEFAULT_GATEWAY: Represents the collection of addresses of the
	// current gateway servers.
	AddressKeywordDefaultGateway AddressKeyword = 16
	// FW_ADDRESS_KEYWORD_INTRANET:  Represents the collection of addresses that are currently
	// within the local intranet of the computer. For schema versions 0x0200, 0x0201, and
	// 0x020A, this value is invalid and MUST NOT be used.
	AddressKeywordIntranet AddressKeyword = 32
	// FW_ADDRESS_KEYWORD_INTERNET:  Represents the collection of addresses that are currently
	// not within the local intranet or remote intranet of the computer. For schema versions
	// 0x0200, 0x0201, and 0x020A, this value is invalid and MUST NOT be used.
	AddressKeywordInternet AddressKeyword = 64
	// FW_ADDRESS_KEYWORD_PLAYTO_RENDERERS: Represents the collection of addresses of the
	// current Digital Media Renderer devices as defined in [MS-DLNHND] section 3.3. For
	// schema versions 0x0200, 0x0201, and 0x020A, this value is invalid and MUST NOT be
	// used.
	AddressKeywordPlayToRenderers AddressKeyword = 128
	// FW_ADDRESS_KEYWORD_REMOTE_INTRANET: Represents the collection of addresses that are
	// currently within the remote intranet of the computer. For schema versions 0x0200,
	// 0x0201, and 0x020A, this value is invalid and MUST NOT be used.
	AddressKeywordRemoteIntranet AddressKeyword = 256
	// FW_ADDRESS_KEYWORD_CAPTIVE_PORTAL:  Represents the collection of addresses of the
	// current captive portal. For schema versions 0x021D and earlier, this value is invalid
	// and MUST NOT be used.
	AddressKeywordCaptivePortal AddressKeyword = 512
	// FW_ADDRESS_KEYWORD_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 0x0400.
	AddressKeywordMax AddressKeyword = 1024
	// FW_ADDRESS_KEYWORD_MAX_V2_10:  This value and values that exceed this value are
	// not valid and MUST NOT be used by servers and clients with schema version 0x020A
	// and earlier. It is defined for simplicity in writing IDL definitions and code. This
	// symbolic constant has a value of 0x0020.
	AddressKeywordMaxV210 AddressKeyword = 32
	// FW_ADDRESS_KEYWORD_MAX_V2_29:  This value and values that exceed this value are
	// not valid and MUST NOT be used by servers and clients with schema version 0x021D
	// and earlier. It is defined for simplicity in writing IDL definitions and code. This
	// symbolic constant has a value of 0x0200.
	AddressKeywordMaxV229 AddressKeyword = 512
)

func (o AddressKeyword) String() string {
	switch o {
	case AddressKeywordNone:
		return "AddressKeywordNone"
	case AddressKeywordLocalSubnet:
		return "AddressKeywordLocalSubnet"
	case AddressKeywordDNS:
		return "AddressKeywordDNS"
	case AddressKeywordDHCP:
		return "AddressKeywordDHCP"
	case AddressKeywordWINS:
		return "AddressKeywordWINS"
	case AddressKeywordDefaultGateway:
		return "AddressKeywordDefaultGateway"
	case AddressKeywordIntranet:
		return "AddressKeywordIntranet"
	case AddressKeywordInternet:
		return "AddressKeywordInternet"
	case AddressKeywordPlayToRenderers:
		return "AddressKeywordPlayToRenderers"
	case AddressKeywordRemoteIntranet:
		return "AddressKeywordRemoteIntranet"
	case AddressKeywordCaptivePortal:
		return "AddressKeywordCaptivePortal"
	case AddressKeywordMax:
		return "AddressKeywordMax"
	case AddressKeywordMaxV210:
		return "AddressKeywordMaxV210"
	case AddressKeywordMaxV229:
		return "AddressKeywordMaxV229"
	}
	return "Invalid"
}

// Addresses structure represents FW_ADDRESSES RPC structure.
//
// This structure contains a list of address structures. Static and symbolic representations
// are supported, but a structure can contain only one representation type. The address
// structure representations follow:
//
// # Static Representation
//
// * FW_IPV4_SUBNET_LIST ( a230f0e4-8005-4983-89e3-e417b3354ae5 )
//
// * FW_IPV4_RANGE_LIST ( de836948-bf68-424d-906d-406723bd0deb )
//
// * FW_IPV6_SUBNET_LIST ( a8d05086-bc48-453e-aebf-aa11fb95b9bd )
//
// * FW_IPV6_RANGE_LIST ( 062373f2-c3dc-49a9-8b3f-960d2a88332d )
//
// # Symbolic Representation
//
// * FW_ADDRESS_KEYWORD ( d69ec3fe-8507-4524-bdcc-813cbb3bf85f ) enumeration values
//
// The FW_ADDRESSES definition follows:
type Addresses struct {
	// dwV4AddressKeywords:  A combination of FW_ADDRESS_KEYWORD flags. Addresses in this
	// field are specified from the IPv4 address space.
	V4AddressKeywords uint32 `idl:"name:dwV4AddressKeywords" json:"v4_address_keywords"`
	// dwV6AddressKeywords:  A combination of FW_ADDRESS_KEYWORD flags. Addresses in this
	// field are specified from the IPv6 address space.
	V6AddressKeywords uint32 `idl:"name:dwV6AddressKeywords" json:"v6_address_keywords"`
	// V4SubNets:  A list of specifically defined IPv4 address subnets.
	SubnetsV4 *IPv4SubnetList `idl:"name:V4SubNets" json:"subnets_v4"`
	// V4Ranges:  A list of specifically defined IPv4 address ranges.
	RangesV4 *IPv4RangeList `idl:"name:V4Ranges" json:"ranges_v4"`
	// V6SubNets:  A list of specifically defined IPv6 address subnets.
	SubnetsV6 *IPv6SubnetList `idl:"name:V6SubNets" json:"subnets_v6"`
	// V6Ranges:  A list of specifically defined IPv6 address ranges.
	RangesV6 *IPv6RangeList `idl:"name:V6Ranges" json:"ranges_v6"`
}

func (o *Addresses) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Addresses) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.V4AddressKeywords); err != nil {
		return err
	}
	if err := w.WriteData(o.V6AddressKeywords); err != nil {
		return err
	}
	if o.SubnetsV4 != nil {
		if err := o.SubnetsV4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv4SubnetList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RangesV4 != nil {
		if err := o.RangesV4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv4RangeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SubnetsV6 != nil {
		if err := o.SubnetsV6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6SubnetList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RangesV6 != nil {
		if err := o.RangesV6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6RangeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Addresses) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.V4AddressKeywords); err != nil {
		return err
	}
	if err := w.ReadData(&o.V6AddressKeywords); err != nil {
		return err
	}
	if o.SubnetsV4 == nil {
		o.SubnetsV4 = &IPv4SubnetList{}
	}
	if err := o.SubnetsV4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RangesV4 == nil {
		o.RangesV4 = &IPv4RangeList{}
	}
	if err := o.RangesV4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SubnetsV6 == nil {
		o.SubnetsV6 = &IPv6SubnetList{}
	}
	if err := o.SubnetsV6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RangesV6 == nil {
		o.RangesV6 = &IPv6RangeList{}
	}
	if err := o.RangesV6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TrustTupleKeyword type represents FW_TRUST_TUPLE_KEYWORD RPC enumeration.
//
// This enumeration represents flags that are used to identify trust tuples.<24> The
// traffic corresponding to these keywords changes dynamically and is tracked by the
// TrustTuples object (section 3.1.1). All the flags supported by a given schema version
// can be combined.
type TrustTupleKeyword uint16

var (
	// FW_TRUST_TUPLE_KEYWORD_NONE:  This value means that none of the following flags
	// are set. It is defined for simplicity in writing IDL definitions and code.
	TrustTupleKeywordNone TrustTupleKeyword = 0
	// FW_TRUST_TUPLE_KEYWORD_PROXIMITY:  Represents all traffic matching a trust tuple
	// in theTrustTuples collection where IsProximity is true.
	TrustTupleKeywordProximity TrustTupleKeyword = 1
	// FW_TRUST_TUPLE_KEYWORD_PROXIMITY_SHARING:  Represents all traffic matching a trust
	// tuple in the TrustTuples collection where IsProximitySharing is true.
	TrustTupleKeywordProximitySharing TrustTupleKeyword = 2
	// FW_TRUST_TUPLE_KEYWORD_WFD_PRINT:  Represents all traffic matching a trust tuple
	// in the TrustTuples collection where IsWFDPrint is true.
	TrustTupleKeywordWFDPrint TrustTupleKeyword = 4
	// FW_TRUST_TUPLE_KEYWORD_WFD_DISPLAY:  Represents all traffic matching a trust tuple
	// in the TrustTuples collection where IsWFDDevices is true.
	TrustTupleKeywordWFDDisplay TrustTupleKeyword = 8
	// FW_TRUST_TUPLE_KEYWORD_WFD_DEVICES:  Represents all traffic matching a trust tuple
	// in the TrustTuples collection where IsWFDDevices is true.
	TrustTupleKeywordWFDDevices TrustTupleKeyword = 16
	// FW_TRUST_TUPLE_KEYWORD_WFD_KM_DRIVER:  Represents all traffic matching a trust tuple
	// in the TrustTuples collection, where IsWFDMaUsbWirelessDocking is true.
	TrustTupleKeywordWFDKMDriver TrustTupleKeyword = 32
	// FW_TRUST_TUPLE_KEYWORD_UPNP:  Represents all traffic that matches a tuple in the
	// TrustTuples collection, where IsUpnP-Secure-Sockets-with-Teredo is true.
	TrustTupleKeywordUpnp TrustTupleKeyword = 64
	// FW_TRUST_TUPLE_KEYWORD_WFD_CDP:  Represents all traffic that matches a tuple in
	// the TrustTuples collection, where IsWFDCDPSvc is true.
	TrustTupleKeywordWfdcdp TrustTupleKeyword = 128
	// FW_TRUST_TUPLE_KEYWORD_MAX:  This value and values that exceed this value are not
	// valid and MUST NOT be used. It is defined for simplicity in writing IDL definitions
	// and code. This symbolic constant has a value of 0x0100.<25> <26>
	TrustTupleKeywordMax TrustTupleKeyword = 256
	// FW_TRUST_TUPLE_KEYWORD_MAX_V2_20:  This value and values that exceed this value
	// are not valid and MUST NOT be used by servers and clients with schema version 0x0214
	// and earlier. It is defined for simplicity in writing IDL definitions and code. This
	// symbolic constant has a value of 0x0004.
	TrustTupleKeywordMaxV220 TrustTupleKeyword = 4
	// FW_TRUST_TUPLE_KEYWORD_MAX_V2_26: This value and values that exceed this value are
	// not valid and MUST NOT be used by servers and clients with schema version 0x021A
	// and earlier. It is defined for simplicity in writing IDL definitions and code. This
	// symbolic constant has a value of 0x0020.
	TrustTupleKeywordMaxV226 TrustTupleKeyword = 32
	// FW_TRUST_TUPLE_KEYWORD_MAX_V2_27: This value and values that exceed this value are
	// not valid and MUST NOT be used by servers and clients with schema version 0x021B
	// and earlier. It is defined for simplicity in writing IDL definitions and code. This
	// symbolic constant has a value of 0x0080.
	TrustTupleKeywordMaxV227 TrustTupleKeyword = 128
)

func (o TrustTupleKeyword) String() string {
	switch o {
	case TrustTupleKeywordNone:
		return "TrustTupleKeywordNone"
	case TrustTupleKeywordProximity:
		return "TrustTupleKeywordProximity"
	case TrustTupleKeywordProximitySharing:
		return "TrustTupleKeywordProximitySharing"
	case TrustTupleKeywordWFDPrint:
		return "TrustTupleKeywordWFDPrint"
	case TrustTupleKeywordWFDDisplay:
		return "TrustTupleKeywordWFDDisplay"
	case TrustTupleKeywordWFDDevices:
		return "TrustTupleKeywordWFDDevices"
	case TrustTupleKeywordWFDKMDriver:
		return "TrustTupleKeywordWFDKMDriver"
	case TrustTupleKeywordUpnp:
		return "TrustTupleKeywordUpnp"
	case TrustTupleKeywordWfdcdp:
		return "TrustTupleKeywordWfdcdp"
	case TrustTupleKeywordMax:
		return "TrustTupleKeywordMax"
	case TrustTupleKeywordMaxV220:
		return "TrustTupleKeywordMaxV220"
	case TrustTupleKeywordMaxV226:
		return "TrustTupleKeywordMaxV226"
	case TrustTupleKeywordMaxV227:
		return "TrustTupleKeywordMaxV227"
	}
	return "Invalid"
}

// RuleStatus type represents FW_RULE_STATUS RPC enumeration.
//
// This enumeration represents status codes that identify the error states of a policy
// object, including successful states. If an object is in an erroneous state, the enumeration
// value represents a reason for the error.
type RuleStatus uint32

var (
	// FW_RULE_STATUS_OK:  The rule was parsed successfully from the store, is correctly
	// constructed, and has no issue.
	RuleStatusOK RuleStatus = 65536
	// FW_RULE_STATUS_PARTIALLY_IGNORED:  The rule has fields that the service can successfully
	// ignore. The ignored fields can be present only if the policy (such as the Group Policy)
	// was written by future firewall and advanced security components that support a higher
	// schema version. Therefore, this error occurs only if the version of the rule is higher;
	// specifically, a higher minor version means that part of the rule might not be understandable.
	// Because the host firewall component does not understand these new fields, it cannot
	// meaningfully specify what was ignored in the rule.
	RuleStatusPartiallyIgnored RuleStatus = 131072
	// FW_RULE_STATUS_IGNORED:  The rule has a higher major version that the service MUST
	// ignore. Higher major schema versions specify that nothing in the rule is understandable
	// to lower major version components.
	RuleStatusIgnored RuleStatus = 262144
	// FW_RULE_STATUS_PARSING_ERROR:  The rule did not parse correctly.
	RuleStatusParsingError RuleStatus = 524288
	// FW_RULE_STATUS_PARSING_ERROR_NAME:  The name contains characters that are not valid
	// or the length is not valid.
	RuleStatusParsingErrorName RuleStatus = 524289
	// FW_RULE_STATUS_PARSING_ERROR_DESC:  The description contains characters that are
	// not valid or the length is not valid.
	RuleStatusParsingErrorDesc RuleStatus = 524290
	// FW_RULE_STATUS_PARSING_ERROR_APP:  The application contains characters that are
	// not valid or the length is not valid.
	RuleStatusParsingErrorApp RuleStatus = 524291
	// FW_RULE_STATUS_PARSING_ERROR_SVC:  The service contains characters that are not
	// valid or the length is not valid.
	RuleStatusParsingErrorService RuleStatus = 524292
	// FW_RULE_STATUS_PARSING_ERROR_RMA:  The remote machine authentication contains characters
	// that are not valid or the length is not valid.
	RuleStatusParsingErrorRMA RuleStatus = 524293
	// FW_RULE_STATUS_PARSING_ERROR_RUA:  The remote user authentication contains characters
	// that are not valid or the length is not valid.
	RuleStatusParsingErrorRUA RuleStatus = 524294
	// FW_RULE_STATUS_PARSING_ERROR_EMBD:  The embedded context contains characters that
	// are not valid or the length is not valid.
	RuleStatusParsingErrorEmbedded RuleStatus = 524295
	// FW_RULE_STATUS_PARSING_ERROR_RULE_ID:  The rule ID contains characters that are
	// not valid or the length is not valid.
	RuleStatusParsingErrorRuleID RuleStatus = 524296
	// FW_RULE_STATUS_PARSING_ERROR_PHASE1_AUTH:  The Phase1 authentication set ID contains
	// characters that are not valid or the length is not valid.
	RuleStatusParsingErrorPhase1Auth RuleStatus = 524297
	// FW_RULE_STATUS_PARSING_ERROR_PHASE2_CRYPTO:  The Phase2 cryptographic set ID contains
	// characters that are not valid or the length is not valid.
	RuleStatusParsingErrorPhase2Crypto RuleStatus = 524298
	// FW_RULE_STATUS_PARSING_ERROR_PHASE2_AUTH:  The Phase2 authentication set ID contains
	// characters that are not valid or the length is not valid.
	RuleStatusParsingErrorPhase2Auth RuleStatus = 524299
	// FW_RULE_STATUS_PARSING_ERROR_RESOLVE_APP:  The application name cannot be resolved.
	RuleStatusParsingErrorResolveApp RuleStatus = 524300
	// FW_RULE_STATUS_PARSING_ERROR_MAINMODE_ID:  This error is unused and not returned
	// by the system.
	RuleStatusParsingErrorMainModuleID RuleStatus = 524301
	// FW_RULE_STATUS_PARSING_ERROR_PHASE1_CRYPTO:  The Phase1 cryptographic set ID contains
	// characters that are not valid or the length is not valid.
	RuleStatusParsingErrorPhase1Crypto RuleStatus = 524302
	// FW_RULE_STATUS_PARSING_ERROR_REMOTE_ENDPOINTS:  The remote tunnel endpoints contain
	// characters that are not valid, or the length is not valid.
	RuleStatusParsingErrorRemoteEndpoints RuleStatus = 524303
	// FW_RULE_STATUS_PARSING_ERROR_REMOTE_ENDPOINT_FQDN: The remote tunnel endpoint fully
	// qualified domain name (FQDN) contains characters that are not valid, or the length
	// is not valid.
	RuleStatusParsingErrorRemoteEndpointFQDN RuleStatus = 524304
	// FW_RULE_STATUS_PARSING_ERROR_KEY_MODULE:  The keying modules contain characters
	// that are not valid, or the length is not valid.
	RuleStatusParsingErrorKeyModule RuleStatus = 524305
	// FW_RULE_STATUS_PARSING_ERROR_LUA:  The local user authorization list contains characters
	// that are not valid or the length is not valid.
	RuleStatusParsingErrorLUA RuleStatus = 524306
	// FW_RULE_STATUS_PARSING_ERROR_FWD_LIFETIME:  The forward path security association
	// (SA) lifetime contains characters that are not valid or the length is not valid.
	RuleStatusParsingErrorFwdLifetime RuleStatus = 524307
	// FW_RULE_STATUS_PARSING_ERROR_TRANSPORT_MACHINE_AUTHZ_SDDL: The IPsec transport mode
	// machine authorization SDDL string contains characters that are not valid, or the
	// length is not valid.
	RuleStatusParsingErrorTransportMachineAuthzSDDL RuleStatus = 524308
	// FW_RULE_STATUS_PARSING_ERROR_TRANSPORT_USER_AUTHZ_SDDL: The IPsec transport mode
	// user authorization SDDL string contains characters that are not valid, or the length
	// is not valid.
	RuleStatusParsingErrorTransportUserAuthzSDDL RuleStatus = 524309
	// FW_RULE_STATUS_PARSING_ERROR_NETNAMES_STRING: A string for the network name structure
	// is invalid.
	RuleStatusParsingErrorNetNamesString RuleStatus = 524310
	// FW_RULE_STATUS_PARSING_ERROR_SECURITY_REALM_ID_STRING: A string for the security
	// realm ID is invalid.
	RuleStatusParsingErrorSecurityRealmIDString RuleStatus = 524311
	// FW_RULE_STATUS_PARSING_ERROR_FQBN_STRING: A string for the fully qualified binary
	// name (FQBN) is invalid; also see [MSDN-FQBN].
	RuleStatusParsingErrorFQBNString RuleStatus = 524312
	// FW_RULE_STATUS_SEMANTIC_ERROR:  There is a semantic error when considering the fields
	// of the rule in conjunction with other policy objects.
	RuleStatusSemanticError RuleStatus = 1048576
	// FW_RULE_STATUS_SEMANTIC_ERROR_RULE_ID:  Semantic error: The rule ID is not specified.
	RuleStatusSemanticErrorRuleID RuleStatus = 1048592
	// FW_RULE_STATUS_SEMANTIC_ERROR_PORTS:  Semantic error: Mismatch in the number of
	// ports and port buffers.
	RuleStatusSemanticErrorPorts RuleStatus = 1048608
	// FW_RULE_STATUS_SEMANTIC_ERROR_PORT_KEYW:  Semantic error: The port keyword is not
	// valid.
	RuleStatusSemanticErrorPortKeyword RuleStatus = 1048609
	// FW_RULE_STATUS_SEMANTIC_ERROR_PORT_RANGE:  Semantic error: End != Begin or port
	// = 0.
	RuleStatusSemanticErrorPortRange RuleStatus = 1048610
	// FW_RULE_STATUS_SEMANTIC_ERROR_PORTRANGE_RESTRICTION: Semantic error: A port range
	// has been specified for a connection security rule, but the action is not Do Not Secure.
	RuleStatusSemanticErrorPortrangeRestriction RuleStatus = 1048611
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_V4_SUBNETS:  Semantic error: Mismatch in the
	// number of v4 subnets and subnet buffers.
	RuleStatusSemanticErrorAddrV4Subnets RuleStatus = 1048640
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_V6_SUBNETS:  Semantic error: Mismatch in the
	// number of v6 subnets and subnet buffers.
	RuleStatusSemanticErrorAddrV6Subnets RuleStatus = 1048641
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_V4_RANGES:  Semantic error: Mismatch in the number
	// of v4 ranges and range buffers.
	RuleStatusSemanticErrorAddrV4Ranges RuleStatus = 1048642
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_V6_RANGES:  Semantic error: Mismatch in the number
	// of v6 ranges and range buffers.
	RuleStatusSemanticErrorAddrV6Ranges RuleStatus = 1048643
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_RANGE:  Semantic error: End < Begin.
	RuleStatusSemanticErrorAddrRange RuleStatus = 1048644
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_MASK:  Semantic error: The mask specified on
	// a v4 subnet is not valid.
	RuleStatusSemanticErrorAddrMask RuleStatus = 1048645
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_PREFIX:  Semantic error: The prefix specified
	// on a v6 subnet is not valid.
	RuleStatusSemanticErrorAddrPrefix RuleStatus = 1048646
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_KEYW:  Semantic error: The specified keyword
	// is not valid.
	RuleStatusSemanticErrorAddrKeyword RuleStatus = 1048647
	// FW_RULE_STATUS_SEMANTIC_ERROR_LADDR_PROP:  Semantic error: A property on local addresses
	// does not belong to the LocalAddress.
	RuleStatusSemanticErrorLocalAddrProperty RuleStatus = 1048648
	// FW_RULE_STATUS_SEMANTIC_ERROR_RADDR_PROP:  Semantic error: A property on remote
	// addresses does not belong to the RemoteAddress.
	RuleStatusSemanticErrorRemoteAddrProperty RuleStatus = 1048649
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_V6:  Semantic error: An unspecified or loopback
	// IPv6 address was specified.
	RuleStatusSemanticErrorAddrV6 RuleStatus = 1048650
	// FW_RULE_STATUS_SEMANTIC_ERROR_LADDR_INTF:  Semantic error: A local address cannot
	// be used together with either an interface or an interface type.
	RuleStatusSemanticErrorLocalAddrInterface RuleStatus = 1048651
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_V4:  Semantic error: An unspecified or loopback
	// IPv4 address was specified.
	RuleStatusSemanticErrorAddrV4 RuleStatus = 1048652
	// FW_RULE_STATUS_SEMANTIC_ERROR_TUNNEL_ENDPOINT_ADDR: Semantic error: An endpoint "any"
	// cannot be specified for a tunnel mode rule.
	RuleStatusSemanticErrorTunnelEndpointAddr RuleStatus = 1048653
	// FW_RULE_STATUS_SEMANTIC_ERROR_DTE_VER:  Semantic error: An incorrect schema version
	// was specified for using dynamic tunnel endpoints.
	RuleStatusSemanticErrorDTEVer RuleStatus = 1048654
	// FW_RULE_STATUS_SEMANTIC_ERROR_DTE_MISMATCH_ADDR:  Semantic error: The v4 and v6
	// tunnel endpoints are neither local nor remote endpoints.
	RuleStatusSemanticErrorDTEMismatchAddr RuleStatus = 1048655
	// FW_RULE_STATUS_SEMANTIC_ERROR_PROFILE:  Semantic error: The profile type is not
	// valid.
	RuleStatusSemanticErrorProfile RuleStatus = 1048656
	// FW_RULE_STATUS_SEMANTIC_ERROR_ICMP:  Semantic error: Mismatch in the number of ICMPs
	// and ICMP buffers.
	RuleStatusSemanticErrorICMP RuleStatus = 1048672
	// FW_RULE_STATUS_SEMANTIC_ERROR_ICMP_CODE:  Semantic error: The specified ICMP code
	// is not valid.
	RuleStatusSemanticErrorICMPCode RuleStatus = 1048673
	// FW_RULE_STATUS_SEMANTIC_ERROR_IF_ID:  Semantic error: Mismatch in the number of
	// interfaces and interface buffers.
	RuleStatusSemanticErrorInterfaceID RuleStatus = 1048688
	// FW_RULE_STATUS_SEMANTIC_ERROR_IF_TYPE:  Semantic error: The specified interface
	// type is not valid.
	RuleStatusSemanticErrorInterfaceType RuleStatus = 1048689
	// FW_RULE_STATUS_SEMANTIC_ERROR_ACTION:  Semantic error: The specified action is not
	// valid.
	RuleStatusSemanticErrorAction RuleStatus = 1048704
	// FW_RULE_STATUS_SEMANTIC_ERROR_ALLOW_BYPASS:  Semantic error: An allow-bypass action
	// is specified, but the rule does not meet allow-bypass criteria (such as, the direction
	// is inbound, authenticate/encrypt flags are set, or remote machine authentication
	// is set).
	RuleStatusSemanticErrorAllowBypass RuleStatus = 1048705
	// FW_RULE_STATUS_SEMANTIC_ERROR_DO_NOT_SECURE:  Semantic error: A DO_NOT_SECURE action
	// is specified together with authentication or cryptographic sets.
	RuleStatusSemanticErrorDoNotSecure RuleStatus = 1048706
	// FW_RULE_STATUS_SEMANTIC_ERROR_ACTION_BLOCK_IS_ENCRYPTED_SECURE: Semantic error: A
	// block action was specified together with a require security or a require encryption
	// action.
	RuleStatusSemanticErrorActionBlockIsEncryptedSecure              RuleStatus = 1048707
	RuleStatusSemanticErrorIncompatibleFlagOrActionWithSecurityRealm RuleStatus = 1048708
	// FW_RULE_STATUS_SEMANTIC_ERROR_DIR:  Semantic error: The specified direction is not
	// valid.
	RuleStatusSemanticErrorDir RuleStatus = 1048720
	// FW_RULE_STATUS_SEMANTIC_ERROR_PROT:  Semantic error: The specified protocol is not
	// valid.
	RuleStatusSemanticErrorProtocol RuleStatus = 1048736
	// FW_RULE_STATUS_SEMANTIC_ERROR_PROT_PROP:  Semantic error: The protocol and protocol-dependent
	// fields do not match.
	RuleStatusSemanticErrorProtocolProperty RuleStatus = 1048737
	// FW_RULE_STATUS_SEMANTIC_ERROR_DEFER_EDGE_PROP:  Semantic error: A Dynamic edge flag
	// (either defer to app or defer to user) is set without having an edge flag set.
	RuleStatusSemanticErrorDeferEdgeProperty RuleStatus = 1048738
	// FW_RULE_STATUS_SEMANTIC_ERROR_ALLOW_BYPASS_OUTBOUND: Semantic error: An outbound
	// allow-bypass action is specified, but the rule does not meet allow-bypass criteria
	// (authenticate/encrypt flags set).
	RuleStatusSemanticErrorAllowBypassOutbound RuleStatus = 1048739
	// FW_RULE_STATUS_SEMANTIC_ERROR_DEFER_USER_INVALID_RULE: The rule does not allow the
	// defer user property to be set.
	RuleStatusSemanticErrorDeferUserInvalidRule RuleStatus = 1048740
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS:  Semantic error: The specified flags are not
	// valid.
	RuleStatusSemanticErrorFlags RuleStatus = 1048752
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_AUTO_AUTH:  Semantic error: The autogenerate
	// flag is set, but no authentication flags are set.
	RuleStatusSemanticErrorFlagsAutoAuth RuleStatus = 1048753
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_AUTO_BLOCK:  Semantic error: The autogenerate
	// flag is set, but the action is block.
	RuleStatusSemanticErrorFlagsAutoBlock RuleStatus = 1048754
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_AUTO_DYN_RPC: Semantic error: The autogenerate
	// flag is set together with the dynamic RPC flag.
	RuleStatusSemanticErrorFlagsAutoDynRPC RuleStatus = 1048755
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_AUTHENTICATE_ENCRYPT: Semantic error: The authenticate
	// and authenticate-encrypt flags are both specified.
	RuleStatusSemanticErrorFlagsAuthenticateEncrypt RuleStatus = 1048756
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_AUTH_WITH_ENC_NEGOTIATE_VER: Semantic error:
	// The schema version is not compliant with the Authenticate with Encryption flag.
	RuleStatusSemanticErrorFlagsAuthWithEncNegotiateVer RuleStatus = 1048757
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_AUTH_WITH_ENC_NEGOTIATE: Semantic error: The
	// Authenticate with Encryption Negotiate flag is specified but the basic Authenticate
	// with Encryption flag is not set.
	RuleStatusSemanticErrorFlagsAuthWithEncNegotiate RuleStatus = 1048758
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_ESP_NO_ENCAP_VER: Semantic error: The schema
	// version is not compliant with the Authenticate with No Encapsulation flag.
	RuleStatusSemanticErrorFlagsESPNoEncapVer RuleStatus = 1048759
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_ESP_NO_ENCAP: Semantic error: The Authenticate
	// with No Encapsulation flag is specified but the basic Authenticate flag is not set.
	RuleStatusSemanticErrorFlagsESPNoEncap RuleStatus = 1048760
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_TUNNEL_AUTH_MODES_VER: Semantic error: The schema
	// version is not compliant with the tunnel authentication modes.
	RuleStatusSemanticErrorFlagsTunnelAuthModesVer RuleStatus = 1048761
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_TUNNEL_AUTH_MODES: Semantic error: The tunnel
	// authentication modes are specified by a lower-version client.
	RuleStatusSemanticErrorFlagsTunnelAuthModes RuleStatus = 1048762
	RuleStatusSemanticErrorFlagsIPHTTPSVer      RuleStatus = 1048763
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_IP_TLS_VER:  Semantic error: The schema version
	// is not compliant with the IP_TLS flag.
	RuleStatusSemanticErrorFlagsIPTLSVer RuleStatus = 1048763
	RuleStatusSemanticErrorPortrangeVer  RuleStatus = 1048764
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_ADDRS_TRAVERSE_DEFER_VER: Semantic error: The
	// schema version is not compliant with the FW_RULE_FLAGS_ROUTEABLE_ADDRS_TRAVERSE_DEFER_APP
	// flag. For more information, see 2.2.35.
	RuleStatusSemanticErrorFlagsAddrsTraverseDeferVer RuleStatus = 1048765
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_AUTH_WITH_ENC_NEGOTIATE_OUTBOUND: Semantic error:
	// The Authenticate with Encryption Negotiate flag is set for the outbound rule.
	RuleStatusSemanticErrorFlagsAuthWithEncNegotiateOutbound RuleStatus = 1048766
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_AUTHENTICATE_WITH_OUTBOUND_BYPASS_VER: Semantic
	// error: The Outbound Authenticated bypass is not supported on this version.
	RuleStatusSemanticErrorFlagsAuthenticateWithOutboundBypassVer RuleStatus = 1048767
	// FW_RULE_STATUS_SEMANTIC_ERROR_REMOTE_AUTH_LIST:  Semantic error: An authorized remote
	// machine or user list is specified, but the authenticate/encryption flags were not
	// set.
	RuleStatusSemanticErrorRemoteAuthList RuleStatus = 1048768
	// FW_RULE_STATUS_SEMANTIC_ERROR_REMOTE_USER_LIST:  Semantic error: An authorized remote
	// user list is specified on an outbound direction.
	RuleStatusSemanticErrorRemoteUserList RuleStatus = 1048769
	// FW_RULE_STATUS_SEMANTIC_ERROR_LOCAL_USER_LIST:  Semantic error: The authorized local
	// user list is specified, but a local service has also been specified.
	RuleStatusSemanticErrorLocalUserList RuleStatus = 1048770
	// FW_RULE_STATUS_SEMANTIC_ERROR_LUA_VER:  Semantic error: The schema version is not
	// compliant with the authorized local user list.
	RuleStatusSemanticErrorLUAVer RuleStatus = 1048771
	// FW_RULE_STATUS_SEMANTIC_ERROR_LOCAL_USER_OWNER:  Semantic error: The local user
	// owner is specified, but a local service has also been specified.
	RuleStatusSemanticErrorLocalUserOwner RuleStatus = 1048772
	// FW_RULE_STATUS_SEMANTIC_ERROR_LOCAL_USER_OWNER_VER: Semantic error: The schema version
	// is not compliant with the local user owner.
	RuleStatusSemanticErrorLocalUserOwnerVer   RuleStatus = 1048773
	RuleStatusSemanticErrorLUAConditionalVer   RuleStatus = 1048774
	RuleStatusSemanticErrorFlagsSystemOSGameOS RuleStatus = 1048775
	RuleStatusSemanticErrorFlagsCortanaVer     RuleStatus = 1048776
	RuleStatusSemanticErrorFlagsRemotename     RuleStatus = 1048777
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_ALLOW_PROFILE_CROSSING_VER: Semantic error: The
	// schema version is not compliant with profile crossing.
	RuleStatusSemanticErrorFlagsAllowProfileCrossingVer RuleStatus = 1048784
	RuleStatusSemanticErrorFlagsLocalOnlyMappedVer      RuleStatus = 1048785
	// FW_RULE_STATUS_SEMANTIC_ERROR_PLATFORM:  Semantic error: The number of valid operating
	// system platforms and the list of valid operating system platforms do not match.
	RuleStatusSemanticErrorPlatform RuleStatus = 1048800
	// FW_RULE_STATUS_SEMANTIC_ERROR_PLATFORM_OP_VER:  Semantic error: Schema version not
	// compliant with the platform operator used.
	RuleStatusSemanticErrorPlatformOperationVer RuleStatus = 1048801
	// FW_RULE_STATUS_SEMANTIC_ERROR_PLATFORM_OP:  Semantic error: Invalid platform operator
	// used.
	RuleStatusSemanticErrorPlatformOperation RuleStatus = 1048802
	// FW_RULE_STATUS_SEMANTIC_ERROR_DTE_NOANY_ADDR:  Semantic error: DTE is specified
	// but all tunnel endpoints are specified.
	RuleStatusSemanticErrorDTENoanyAddr            RuleStatus = 1048816
	RuleStatusSemanticErrorTunnelExemptWithGateway RuleStatus = 1048817
	RuleStatusSemanticErrorTunnelExemptVer         RuleStatus = 1048818
	// FW_RULE_STATUS_SEMANTIC_ERROR_ADDR_KEYWORD_VER:  Semantic error: The schema version
	// is not compliant with one or more address keywords.
	RuleStatusSemanticErrorAddrKeywordVer RuleStatus = 1048819
	// FW_RULE_STATUS_SEMANTIC_ERROR_KEY_MODULE_VER:  Semantic error: The schema version
	// is not compliant with the keying modules.
	RuleStatusSemanticErrorKeyModuleVer RuleStatus = 1048820
	// FW_RULE_STATUS_SEMANTIC_ERROR_APP_CONTAINER_PACKAGE_ID: Semantic error: The application
	// container package ID is not a valid security identifier (SID).
	RuleStatusSemanticErrorAppContainerPackageID RuleStatus = 1048832
	// FW_RULE_STATUS_SEMANTIC_ERROR_APP_CONTAINER_PACKAGE_ID_VER: Semantic error: The schema
	// version is not compliant with application containers.
	RuleStatusSemanticErrorAppContainerPackageIDVer RuleStatus = 1048833
	// FW_RULE_STATUS_SEMANTIC_ERROR_TRUST_TUPLE_KEYWORD_INCOMPATIBLE: Semantic error: Trust
	// tuple keywords are specified, but specific addresses or ports have also been specified.
	RuleStatusSemanticErrorTrustTupleKeywordIncompatible RuleStatus = 1049088
	// FW_RULE_STATUS_SEMANTIC_ERROR_TRUST_TUPLE_KEYWORD_INVALID: Semantic error: One or
	// more trust tuple keywords is invalid.
	RuleStatusSemanticErrorTrustTupleKeywordInvalid RuleStatus = 1049089
	// FW_RULE_STATUS_SEMANTIC_ERROR_TRUST_TUPLE_KEYWORD_VER: Semantic error: The schema
	// version is not compliant with the trust tuple keywords.
	RuleStatusSemanticErrorTrustTupleKeywordVer RuleStatus = 1049090
	RuleStatusSemanticErrorInterfaceTypesVer    RuleStatus = 1049345
	RuleStatusSemanticErrorNetNamesVer          RuleStatus = 1049601
	RuleStatusSemanticErrorSecurityRealmIDVer   RuleStatus = 1049602
	RuleStatusSemanticErrorSystemOSGameOSVer    RuleStatus = 1049603
	RuleStatusSemanticErrorDevModeVer           RuleStatus = 1049604
	RuleStatusSemanticErrorRemoteServernameVer  RuleStatus = 1049605
	RuleStatusSemanticErrorFQBNVer              RuleStatus = 1049606
	RuleStatusSemanticErrorCompartmentIDVer     RuleStatus = 1049607
	RuleStatusSemanticErrorCalloutAndAuditVer   RuleStatus = 1049608
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_AUTH_SET_ID: Semantic error: Phase1 authentication
	// set ID is not specified.
	RuleStatusSemanticErrorPhase1AuthSetID RuleStatus = 1049856
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_CRYPTO_SET_ID: Semantic error: Phase2 cryptographic
	// set ID is not specified.
	RuleStatusSemanticErrorPhase2CryptoSetID RuleStatus = 1049872
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_SET_ID: Semantic error: Phase1 cryptographic
	// set ID is not specified.
	RuleStatusSemanticErrorPhase1CryptoSetID RuleStatus = 1049873
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_KEY_MANAGER_DICTATE_VER: Semantic error: The
	// schema version is not compliant with the Key Manager Dictation flag.
	RuleStatusSemanticErrorFlagsKeyManagerDictateVer RuleStatus = 1049874
	// FW_RULE_STATUS_SEMANTIC_ERROR_FLAGS_KEY_MANAGER_NOTIFY_VER: Semantic error: The schema
	// version is not compliant with the Key Manager Notification flag.
	RuleStatusSemanticErrorFlagsKeyManagerNotifyVer RuleStatus = 1049875
	// FW_RULE_STATUS_SEMANTIC_ERROR_TRANSPORT_MACHINE_AUTHZ_VER: Semantic error: The schema
	// version is not compliant with IPsec transport mode machine authorization lists.
	RuleStatusSemanticErrorTransportMachineAuthzVer RuleStatus = 1049876
	// FW_RULE_STATUS_SEMANTIC_ERROR_TRANSPORT_USER_AUTHZ_VER: Semantic error: The schema
	// version is not compliant with IPsec transport mode user authorization lists.
	RuleStatusSemanticErrorTransportUserAuthzVer RuleStatus = 1049877
	// FW_RULE_STATUS_SEMANTIC_ERROR_TRANSPORT_MACHINE_AUTHZ_ON_TUNNEL: Semantic error:
	// An IPsec transport mode machine authorization list is specified on a tunnel mode
	// rule.
	RuleStatusSemanticErrorTransportMachineAuthzOnTunnel RuleStatus = 1049878
	// FW_RULE_STATUS_SEMANTIC_ERROR_TRANSPORT_USER_AUTHZ_ON_TUNNEL: Semantic error: An
	// IPsec transport mode user authorization list is specified on a tunnel mode rule.
	RuleStatusSemanticErrorTransportUserAuthzOnTunnel RuleStatus = 1049879
	// FW_RULE_STATUS_SEMANTIC_ERROR_PER_RULE_AND_GLOBAL_AUTHZ: Semantic error: The FW_CS_RULE_FLAGS_APPLY_AUTHZ
	// flag is set, but a per-rule authorization list is also specified.
	RuleStatusSemanticErrorPerRuleAndGlobalAuthz RuleStatus = 1049880
	RuleStatusSemanticErrorFlagsSecurityRealm    RuleStatus = 1049881
	// FW_RULE_STATUS_SEMANTIC_ERROR_SET_ID:  Semantic error: The set ID is not specified.
	RuleStatusSemanticErrorSetID RuleStatus = 1052672
	// FW_RULE_STATUS_SEMANTIC_ERROR_IPSEC_PHASE:  Semantic error: The specified phase
	// is not valid.
	RuleStatusSemanticErrorIIIPsecPhase RuleStatus = 1052688
	// FW_RULE_STATUS_SEMANTIC_ERROR_EMPTY_SUITES:  Semantic error: No suites are specified
	// in the set.
	RuleStatusSemanticErrorEmptySuites RuleStatus = 1052704
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_AUTH_METHOD: Semantic error: The Phase1 authentication
	// method is not valid.
	RuleStatusSemanticErrorPhase1AuthMethod RuleStatus = 1052720
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_AUTH_METHOD: Semantic error: The Phase2 authentication
	// method is not valid.
	RuleStatusSemanticErrorPhase2AuthMethod RuleStatus = 1052721
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_METHOD_ANONYMOUS: Semantic error: Anonymous authentication
	// is specified as the only authentication proposal (or authentication proposal suite).
	RuleStatusSemanticErrorAuthMethodAnonymous RuleStatus = 1052722
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_METHOD_DUPLICATE: Semantic error: Duplicate authentication
	// methods are specified but not supported.
	RuleStatusSemanticErrorAuthMethodDuplicate RuleStatus = 1052723
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_METHOD_VER:  Semantic error: Suite specifies
	// authentication method that is not compliant with its schema version.
	RuleStatusSemanticErrorAuthMethodVer RuleStatus = 1052724
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_SUITE_FLAGS:  Semantic error: The specified authentication
	// suite flags are not valid.
	RuleStatusSemanticErrorAuthSuiteFlags RuleStatus = 1052736
	// FW_RULE_STATUS_SEMANTIC_ERROR_HEALTH_CERT:  Semantic error: The machine certificate
	// MUST be a health certificate for Phase2 authentication.
	RuleStatusSemanticErrorHealthCert RuleStatus = 1052737
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_SIGNCERT_VER:  Semantic error: The suite specifies
	// signing that is not compliant with its schema version.
	RuleStatusSemanticErrorAuthSignCertVer RuleStatus = 1052738
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_INTERMEDIATE_CA_VER: Semantic error: Specifies
	// an intermediate certificate authority (CA) that is not compliant with its schema
	// version.
	RuleStatusSemanticErrorAuthIntermediateCAVer RuleStatus = 1052739
	// FW_RULE_STATUS_SEMANTIC_ERROR_MACHINE_SHKEY:  Semantic error: The machine shared
	// key is either missing or not valid.
	RuleStatusSemanticErrorMachineSharedKey RuleStatus = 1052752
	// FW_RULE_STATUS_SEMANTIC_ERROR_CA_NAME:  Semantic error: The CA name is either missing
	// or not valid.
	RuleStatusSemanticErrorCAName RuleStatus = 1052768
	// FW_RULE_STATUS_SEMANTIC_ERROR_MIXED_CERTS:  Semantic error: Health certificates
	// (CERTS) cannot be specified together with regular certificates.
	RuleStatusSemanticErrorMixedCerts RuleStatus = 1052769
	// FW_RULE_STATUS_SEMANTIC_ERROR_NON_CONTIGUOUS_CERTS: Semantic error: Certificates
	// that have a specific signing algorithm are not contiguous.
	RuleStatusSemanticErrorNonContiguousCerts RuleStatus = 1052770
	// FW_RULE_STATUS_SEMANTIC_ERROR_MIXED_CA_TYPE_IN_BLOCK: Semantic error: Both root and
	// intermediate CA types cannot be present in the same signing algorithm block.
	RuleStatusSemanticErrorMixedCATypeInBlock RuleStatus = 1052771
	// FW_RULE_STATUS_SEMANTIC_ERROR_MACHINE_USER_AUTH:  Semantic error: Both machine and
	// user authentications are specified.
	RuleStatusSemanticErrorMachineUserAuth RuleStatus = 1052784
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_CERT_CRITERIA_VER: The suite specifies certificate
	// criteria but the schema version does not allow certificate criteria to be present.
	// Certificate criteria are supported only in schemas with version number 2.20 and greater.
	RuleStatusSemanticErrorAuthCertCriteriaVer RuleStatus = 1052785
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_CERT_CRITERIA_VER_MISMATCH: The version specified
	// for the criteria structure is different from the auth set version.
	RuleStatusSemanticErrorAuthCertCriteriaVerMismatch RuleStatus = 1052786
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_CERT_CRITERIA_RENEWAL_HASH: Cert criteria were
	// specified for a non-cert authentication method.
	RuleStatusSemanticErrorAuthCertCriteriaRenewalHash RuleStatus = 1052787
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_CERT_CRITERIA_INVALID_HASH: An invalid hash was
	// specified in the criteria. A valid hash is a string of hex characters (40 characters
	// in length).
	RuleStatusSemanticErrorAuthCertCriteriaInvalidHash RuleStatus = 1052788
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_CERT_CRITERIA_INVALID_EKU: An invalid EKU was
	// specified. Validity checking of an EKU involves checking that the EKU is composed
	// of characters representing 0 to 9 and ".".
	RuleStatusSemanticErrorAuthCertCriteriaInvalidEKU RuleStatus = 1052789
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_CERT_CRITERIA_INVALID_NAME_TYPE: A name type greater
	// than FW_CERT_CRITERIA_NAME_MAX was specified.
	RuleStatusSemanticErrorAuthCertCriteriaInvalidNameType RuleStatus = 1052790
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_CERT_CRITERIA_INVALID_NAME: A name type was specified but either a NULL name is also specified, or the number of characters in the name is greater than FW_MAX_RULE_STRING_LEN(10000), or the name string contains the "|" character.
	RuleStatusSemanticErrorAuthCertCriteriaInvalidName RuleStatus = 1052791
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_CERT_CRITERIA_INVALID_CRITERIA_TYPE: The criteria
	// type specified is greater than FW_CERT_CRITERIA_TYPE_MAX.
	RuleStatusSemanticErrorAuthCertCriteriaInvalidCriteriaType RuleStatus = 1052792
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_CERT_CRITERIA_MISSING_CRITERIA: The specified
	// suites are missing either selection or validation criteria.
	RuleStatusSemanticErrorAuthCertCriteriaMissingCriteria RuleStatus = 1052793
	// FW_RULE_STATUS_SEMANTIC_ERROR_PROXY_SERVER:  Semantic error: The Kerberos proxy
	// server name is not a valid fully qualified domain name (FQDN).
	RuleStatusSemanticErrorProxyServer RuleStatus = 1052800
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_PROXY_SERVER_VER: Semantic error: The schema version
	// is not compliant with Kerberos proxy servers.
	RuleStatusSemanticErrorAuthProxyServerVer RuleStatus = 1052801
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_NON_DEFAULT_ID: Semantic error: The ID
	// for the Phase1 cryptographic set is not the default.
	RuleStatusSemanticErrorPhase1CryptoNonDefaultID RuleStatus = 1069056
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_FLAGS: Semantic error: The Phase1 cryptographic
	// set flags are not valid.
	RuleStatusSemanticErrorPhase1CryptoFlags RuleStatus = 1069057
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_TIMEOUT_MINUTES: Semantic error: The
	// Phase1 cryptographic set time-out minutes are not valid.
	RuleStatusSemanticErrorPhase1CryptoTimeoutMinutes RuleStatus = 1069058
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_TIMEOUT_SESSIONS: Semantic error: The
	// time-out sessions for the Phase1 cryptographic set are not valid.
	RuleStatusSemanticErrorPhase1CryptoTimeoutSessions RuleStatus = 1069059
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_KEY_EXCHANGE: Semantic error: The key
	// exchange for the Phase1 cryptographic set is not valid.
	RuleStatusSemanticErrorPhase1CryptoKeyExchange RuleStatus = 1069060
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_ENCRYPTION: Semantic error: The Phase1
	// cryptographic set encryption is not valid.
	RuleStatusSemanticErrorPhase1CryptoEncryption RuleStatus = 1069061
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_HASH: Semantic error: The Phase1 cryptographic
	// set hash is not valid.
	RuleStatusSemanticErrorPhase1CryptoHash RuleStatus = 1069062
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_ENCRYPTION_VER: Semantic error: The Phase1
	// cryptographic set encryption is not schema-version compliant.
	RuleStatusSemanticErrorPhase1CryptoEncryptionVer RuleStatus = 1069063
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_HASH_VER: Semantic error: The Phase1
	// cryptographic set hash is not schema version compliant.
	RuleStatusSemanticErrorPhase1CryptoHashVer RuleStatus = 1069064
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE1_CRYPTO_KEY_EXCH_VER: Semantic error: The schema
	// version is not compliant with one or more of the specified main mode key exchange
	// algorithms.
	RuleStatusSemanticErrorPhase1CryptoKeyExchVer RuleStatus = 1069065
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_CRYPTO_PFS:  Semantic error: The Phase2 cryptographic
	// set perfect forward secrecy (PFS) is not valid.
	RuleStatusSemanticErrorPhase2CryptoPFS RuleStatus = 1069088
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_CRYPTO_PROTOCOL: Semantic error: The Phase2
	// cryptographic set protocol is not valid.
	RuleStatusSemanticErrorPhase2CryptoProtocol RuleStatus = 1069089
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_CRYPTO_ENCRYPTION: Semantic error: The Phase2
	// cryptographic set encryption is not valid.
	RuleStatusSemanticErrorPhase2CryptoEncryption RuleStatus = 1069090
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_CRYPTO_HASH: Semantic error: The Phase2 cryptographic
	// set hash is not valid.
	RuleStatusSemanticErrorPhase2CryptoHash RuleStatus = 1069091
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_CRYPTO_TIMEOUT_MINUTES: Semantic error: The
	// Phase2 cryptographic set time-out minutes are not valid.
	RuleStatusSemanticErrorPhase2CryptoTimeoutMinutes RuleStatus = 1069092
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_CRYPTO_TIMEOUT_KBYTES: Semantic error: The Phase2
	// cryptographic set time-out kilobytes are not valid.
	RuleStatusSemanticErrorPhase2CryptoTimeoutKbytes RuleStatus = 1069093
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_CRYPTO_ENCRYPTION_VER: Semantic error: The Phase2
	// cryptographic set encryption is not schema-version compliant.
	RuleStatusSemanticErrorPhase2CryptoEncryptionVer RuleStatus = 1069094
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_CRYPTO_HASH_VER: The Phase2 cryptographic set
	// hash is not schema-version compliant.
	RuleStatusSemanticErrorPhase2CryptoHashVer RuleStatus = 1069095
	// FW_RULE_STATUS_SEMANTIC_ERROR_PHASE2_CRYPTO_PFS_VER: Semantic error: The schema version
	// is not compliant with the specified Phase2 perfect forward secrecy (PFS) option.
	RuleStatusSemanticErrorPhase2CryptoPFSVer RuleStatus = 1069096
	// FW_RULE_STATUS_SEMANTIC_ERROR_CRYPTO_ENCR_HASH:  Semantic error: Neither the encryption
	// nor the hash is specified.
	RuleStatusSemanticErrorCryptoEncrHash RuleStatus = 1069120
	// FW_RULE_STATUS_SEMANTIC_ERROR_CRYPTO_ENCR_HASH_COMPAT: Semantic error: The encryption
	// and hash use incompatible algorithms.
	RuleStatusSemanticErrorCryptoEncrHashCompatibility RuleStatus = 1069121
	// FW_RULE_STATUS_SEMANTIC_ERROR_SCHEMA_VERSION:  Semantic error: The specified schema
	// version is lower than the lowest supported version.
	RuleStatusSemanticErrorSchemaVersion RuleStatus = 1069136
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_OR_AND_CONDITIONS: Semantic error: A mismatch
	// exists in the number of OR'd terms and term arrays.
	RuleStatusSemanticErrorQueryOrAndConditions RuleStatus = 1073152
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_AND_CONDITIONS: Semantic error: A mismatch exists
	// in the number of AND'd conditions and condition arrays.
	RuleStatusSemanticErrorQueryAndConditions RuleStatus = 1073153
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_CONDITION_KEY: Semantic error: The condition
	// match key is not valid.
	RuleStatusSemanticErrorQueryConditionKey RuleStatus = 1073154
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_CONDITION_MATCH_TYPE: Semantic error: The condition
	// match type is not valid.
	RuleStatusSemanticErrorQueryConditionMatchType RuleStatus = 1073155
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_CONDITION_DATA_TYPE: Semantic error: The condition
	// data type is not valid.
	RuleStatusSemanticErrorQueryConditionDataType RuleStatus = 1073156
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_CONDITION_KEY_AND_DATA_TYPE: Semantic error:
	// The key and data type combination is not valid.
	RuleStatusSemanticErrorQueryConditionKeyAndDataType RuleStatus = 1073157
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_KEYS_PROTOCOL_PORT: Semantic error: A port condition
	// is present without a protocol condition.
	RuleStatusSemanticErrorQueryKeysProtocolPort RuleStatus = 1073158
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_KEY_PROFILE:  Semantic error: The profile key
	// is unavailable for the queried object type.
	RuleStatusSemanticErrorQueryKeyProfile RuleStatus = 1073159
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_KEY_STATUS:  Semantic error: The status key
	// is unavailable for the queried object type.
	RuleStatusSemanticErrorQueryKeyStatus RuleStatus = 1073160
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_KEY_FILTERID: Semantic error: The FilterID key
	// is unavailable for the queried object type.
	RuleStatusSemanticErrorQueryKeyFilterID RuleStatus = 1073161
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_KEY_APP_PATH: Semantic error: The application
	// key is unavailable for the queried object type.
	RuleStatusSemanticErrorQueryKeyAppPath RuleStatus = 1073168
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_KEY_PROTOCOL: Semantic error: The protocol key
	// is unavailable for the queried object type.
	RuleStatusSemanticErrorQueryKeyProtocol RuleStatus = 1073169
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_KEY_LOCAL_PORT: Semantic error: The local port
	// key is unavailable for the queried object type.
	RuleStatusSemanticErrorQueryKeyLocalPort RuleStatus = 1073170
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_KEY_REMOTE_PORT: Semantic error: The remote port
	// key is unavailable for the queried object type.
	RuleStatusSemanticErrorQueryKeyRemotePort RuleStatus = 1073171
	// FW_RULE_STATUS_SEMANTIC_ERROR_QUERY_KEY_SVC_NAME: Semantic error: The service name
	// key is unavailable for the queried object type.
	RuleStatusSemanticErrorQueryKeyServiceName RuleStatus = 1073173
	// FW_RULE_STATUS_SEMANTIC_ERROR_REQUIRE_IN_CLEAR_OUT_ON_TRANSPORT: Semantic error:
	// "Require in clear out" tunnel authentication mode cannot be set on transport mode
	// rules.
	RuleStatusSemanticErrorRequireInClearOutOnTransport           RuleStatus = 1077248
	RuleStatusSemanticErrorBypassTunnelInterfaceSecureOnTransport RuleStatus = 1077249
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_NOENCAP_ON_TUNNEL: Semantic error: Cannot set
	// FW_CRYPTO_PROTOCOL_AUTH_NO_ENCAP (see section 2.2.69) on a tunnel mode rule.
	RuleStatusSemanticErrorAuthNoEncapOnTunnel RuleStatus = 1077250
	// FW_RULE_STATUS_SEMANTIC_ERROR_AUTH_NOENCAP_ON_PSK: Semantic error: Cannot mix FW_CRYPTO_PROTOCOL_AUTH_NO_ENCAP
	// (see section 2.2.69) protocol with Preshared key authentication methods.
	RuleStatusSemanticErrorAuthNoEncapOnPSK RuleStatus = 1077251
	// FW_RULE_STATUS_RUNTIME_ERROR:  There is a runtime error when the object is considered
	// with other policy objects.
	RuleStatusRuntimeError RuleStatus = 2097152
	// FW_RULE_STATUS_RUNTIME_ERROR_PHASE1_AUTH_NOT_FOUND: A Phase1 authentication set is
	// not found.
	RuleStatusRuntimeErrorPhase1AuthNotFound RuleStatus = 2097153
	// FW_RULE_STATUS_RUNTIME_ERROR_PHASE2_AUTH_NOT_FOUND: A Phase2 authentication set is
	// not found.
	RuleStatusRuntimeErrorPhase2AuthNotFound RuleStatus = 2097154
	// FW_RULE_STATUS_RUNTIME_ERROR_PHASE2_CRYPTO_NOT_FOUND: A Phase2 cryptographic set
	// is not found.
	RuleStatusRuntimeErrorPhase2CryptoNotFound RuleStatus = 2097155
	// FW_RULE_STATUS_RUNTIME_ERROR_AUTH_MCHN_SHKEY_MISMATCH: A Phase2 authentication set
	// cannot be specified when the Phase1 authentication set contains a pre-shared key
	// as an authentication method.
	RuleStatusRuntimeErrorAuthMchnSharedKeyMismatch RuleStatus = 2097156
	// FW_RULE_STATUS_RUNTIME_ERROR_PHASE1_CRYPTO_NOT_FOUND: A Phase1 cryptographic set
	// is not found.
	RuleStatusRuntimeErrorPhase1CryptoNotFound RuleStatus = 2097157
	// FW_RULE_STATUS_RUNTIME_ERROR_AUTH_NOENCAP_ON_TUNNEL: Semantic error: Cannot set FW_CRYPTO_PROTOCOL_AUTH_NO_ENCAP
	// (see section 2.2.69) on a tunnel mode rule.
	RuleStatusRuntimeErrorAuthNoEncapOnTunnel RuleStatus = 2097158
	// FW_RULE_STATUS_RUNTIME_ERROR_AUTH_NOENCAP_ON_PSK: Semantic error: Cannot mix FW_CRYPTO_PROTOCOL_AUTH_NO_ENCAP
	// (see section 2.2.69) protocol with Preshared key authentication methods.
	RuleStatusRuntimeErrorAuthNoEncapOnPSK RuleStatus = 2097159
	// FW_RULE_STATUS_RUNTIME_ERROR_KEY_MODULE_AUTH_MISMATCH: Semantic error: The key module
	// in the rule is incompatible with the authentication methods specified in the associated
	// authentication sets.
	RuleStatusRuntimeErrorKeyModuleAuthMismatch RuleStatus = 2097160
	// FW_RULE_STATUS_ERROR:  An error of any kind occurred. This symbolic constant has
	// a value of 0x00380000.
	RuleStatusError RuleStatus = 3670016
	// FW_RULE_STATUS_ALL:  The status of all (it is used to enumerate all the rules, regardless
	// of the status).
	RuleStatusAll RuleStatus = 4294901760
)

func (o RuleStatus) String() string {
	switch o {
	case RuleStatusOK:
		return "RuleStatusOK"
	case RuleStatusPartiallyIgnored:
		return "RuleStatusPartiallyIgnored"
	case RuleStatusIgnored:
		return "RuleStatusIgnored"
	case RuleStatusParsingError:
		return "RuleStatusParsingError"
	case RuleStatusParsingErrorName:
		return "RuleStatusParsingErrorName"
	case RuleStatusParsingErrorDesc:
		return "RuleStatusParsingErrorDesc"
	case RuleStatusParsingErrorApp:
		return "RuleStatusParsingErrorApp"
	case RuleStatusParsingErrorService:
		return "RuleStatusParsingErrorService"
	case RuleStatusParsingErrorRMA:
		return "RuleStatusParsingErrorRMA"
	case RuleStatusParsingErrorRUA:
		return "RuleStatusParsingErrorRUA"
	case RuleStatusParsingErrorEmbedded:
		return "RuleStatusParsingErrorEmbedded"
	case RuleStatusParsingErrorRuleID:
		return "RuleStatusParsingErrorRuleID"
	case RuleStatusParsingErrorPhase1Auth:
		return "RuleStatusParsingErrorPhase1Auth"
	case RuleStatusParsingErrorPhase2Crypto:
		return "RuleStatusParsingErrorPhase2Crypto"
	case RuleStatusParsingErrorPhase2Auth:
		return "RuleStatusParsingErrorPhase2Auth"
	case RuleStatusParsingErrorResolveApp:
		return "RuleStatusParsingErrorResolveApp"
	case RuleStatusParsingErrorMainModuleID:
		return "RuleStatusParsingErrorMainModuleID"
	case RuleStatusParsingErrorPhase1Crypto:
		return "RuleStatusParsingErrorPhase1Crypto"
	case RuleStatusParsingErrorRemoteEndpoints:
		return "RuleStatusParsingErrorRemoteEndpoints"
	case RuleStatusParsingErrorRemoteEndpointFQDN:
		return "RuleStatusParsingErrorRemoteEndpointFQDN"
	case RuleStatusParsingErrorKeyModule:
		return "RuleStatusParsingErrorKeyModule"
	case RuleStatusParsingErrorLUA:
		return "RuleStatusParsingErrorLUA"
	case RuleStatusParsingErrorFwdLifetime:
		return "RuleStatusParsingErrorFwdLifetime"
	case RuleStatusParsingErrorTransportMachineAuthzSDDL:
		return "RuleStatusParsingErrorTransportMachineAuthzSDDL"
	case RuleStatusParsingErrorTransportUserAuthzSDDL:
		return "RuleStatusParsingErrorTransportUserAuthzSDDL"
	case RuleStatusParsingErrorNetNamesString:
		return "RuleStatusParsingErrorNetNamesString"
	case RuleStatusParsingErrorSecurityRealmIDString:
		return "RuleStatusParsingErrorSecurityRealmIDString"
	case RuleStatusParsingErrorFQBNString:
		return "RuleStatusParsingErrorFQBNString"
	case RuleStatusSemanticError:
		return "RuleStatusSemanticError"
	case RuleStatusSemanticErrorRuleID:
		return "RuleStatusSemanticErrorRuleID"
	case RuleStatusSemanticErrorPorts:
		return "RuleStatusSemanticErrorPorts"
	case RuleStatusSemanticErrorPortKeyword:
		return "RuleStatusSemanticErrorPortKeyword"
	case RuleStatusSemanticErrorPortRange:
		return "RuleStatusSemanticErrorPortRange"
	case RuleStatusSemanticErrorPortrangeRestriction:
		return "RuleStatusSemanticErrorPortrangeRestriction"
	case RuleStatusSemanticErrorAddrV4Subnets:
		return "RuleStatusSemanticErrorAddrV4Subnets"
	case RuleStatusSemanticErrorAddrV6Subnets:
		return "RuleStatusSemanticErrorAddrV6Subnets"
	case RuleStatusSemanticErrorAddrV4Ranges:
		return "RuleStatusSemanticErrorAddrV4Ranges"
	case RuleStatusSemanticErrorAddrV6Ranges:
		return "RuleStatusSemanticErrorAddrV6Ranges"
	case RuleStatusSemanticErrorAddrRange:
		return "RuleStatusSemanticErrorAddrRange"
	case RuleStatusSemanticErrorAddrMask:
		return "RuleStatusSemanticErrorAddrMask"
	case RuleStatusSemanticErrorAddrPrefix:
		return "RuleStatusSemanticErrorAddrPrefix"
	case RuleStatusSemanticErrorAddrKeyword:
		return "RuleStatusSemanticErrorAddrKeyword"
	case RuleStatusSemanticErrorLocalAddrProperty:
		return "RuleStatusSemanticErrorLocalAddrProperty"
	case RuleStatusSemanticErrorRemoteAddrProperty:
		return "RuleStatusSemanticErrorRemoteAddrProperty"
	case RuleStatusSemanticErrorAddrV6:
		return "RuleStatusSemanticErrorAddrV6"
	case RuleStatusSemanticErrorLocalAddrInterface:
		return "RuleStatusSemanticErrorLocalAddrInterface"
	case RuleStatusSemanticErrorAddrV4:
		return "RuleStatusSemanticErrorAddrV4"
	case RuleStatusSemanticErrorTunnelEndpointAddr:
		return "RuleStatusSemanticErrorTunnelEndpointAddr"
	case RuleStatusSemanticErrorDTEVer:
		return "RuleStatusSemanticErrorDTEVer"
	case RuleStatusSemanticErrorDTEMismatchAddr:
		return "RuleStatusSemanticErrorDTEMismatchAddr"
	case RuleStatusSemanticErrorProfile:
		return "RuleStatusSemanticErrorProfile"
	case RuleStatusSemanticErrorICMP:
		return "RuleStatusSemanticErrorICMP"
	case RuleStatusSemanticErrorICMPCode:
		return "RuleStatusSemanticErrorICMPCode"
	case RuleStatusSemanticErrorInterfaceID:
		return "RuleStatusSemanticErrorInterfaceID"
	case RuleStatusSemanticErrorInterfaceType:
		return "RuleStatusSemanticErrorInterfaceType"
	case RuleStatusSemanticErrorAction:
		return "RuleStatusSemanticErrorAction"
	case RuleStatusSemanticErrorAllowBypass:
		return "RuleStatusSemanticErrorAllowBypass"
	case RuleStatusSemanticErrorDoNotSecure:
		return "RuleStatusSemanticErrorDoNotSecure"
	case RuleStatusSemanticErrorActionBlockIsEncryptedSecure:
		return "RuleStatusSemanticErrorActionBlockIsEncryptedSecure"
	case RuleStatusSemanticErrorIncompatibleFlagOrActionWithSecurityRealm:
		return "RuleStatusSemanticErrorIncompatibleFlagOrActionWithSecurityRealm"
	case RuleStatusSemanticErrorDir:
		return "RuleStatusSemanticErrorDir"
	case RuleStatusSemanticErrorProtocol:
		return "RuleStatusSemanticErrorProtocol"
	case RuleStatusSemanticErrorProtocolProperty:
		return "RuleStatusSemanticErrorProtocolProperty"
	case RuleStatusSemanticErrorDeferEdgeProperty:
		return "RuleStatusSemanticErrorDeferEdgeProperty"
	case RuleStatusSemanticErrorAllowBypassOutbound:
		return "RuleStatusSemanticErrorAllowBypassOutbound"
	case RuleStatusSemanticErrorDeferUserInvalidRule:
		return "RuleStatusSemanticErrorDeferUserInvalidRule"
	case RuleStatusSemanticErrorFlags:
		return "RuleStatusSemanticErrorFlags"
	case RuleStatusSemanticErrorFlagsAutoAuth:
		return "RuleStatusSemanticErrorFlagsAutoAuth"
	case RuleStatusSemanticErrorFlagsAutoBlock:
		return "RuleStatusSemanticErrorFlagsAutoBlock"
	case RuleStatusSemanticErrorFlagsAutoDynRPC:
		return "RuleStatusSemanticErrorFlagsAutoDynRPC"
	case RuleStatusSemanticErrorFlagsAuthenticateEncrypt:
		return "RuleStatusSemanticErrorFlagsAuthenticateEncrypt"
	case RuleStatusSemanticErrorFlagsAuthWithEncNegotiateVer:
		return "RuleStatusSemanticErrorFlagsAuthWithEncNegotiateVer"
	case RuleStatusSemanticErrorFlagsAuthWithEncNegotiate:
		return "RuleStatusSemanticErrorFlagsAuthWithEncNegotiate"
	case RuleStatusSemanticErrorFlagsESPNoEncapVer:
		return "RuleStatusSemanticErrorFlagsESPNoEncapVer"
	case RuleStatusSemanticErrorFlagsESPNoEncap:
		return "RuleStatusSemanticErrorFlagsESPNoEncap"
	case RuleStatusSemanticErrorFlagsTunnelAuthModesVer:
		return "RuleStatusSemanticErrorFlagsTunnelAuthModesVer"
	case RuleStatusSemanticErrorFlagsTunnelAuthModes:
		return "RuleStatusSemanticErrorFlagsTunnelAuthModes"
	case RuleStatusSemanticErrorFlagsIPHTTPSVer:
		return "RuleStatusSemanticErrorFlagsIPHTTPSVer"
	case RuleStatusSemanticErrorFlagsIPTLSVer:
		return "RuleStatusSemanticErrorFlagsIPTLSVer"
	case RuleStatusSemanticErrorPortrangeVer:
		return "RuleStatusSemanticErrorPortrangeVer"
	case RuleStatusSemanticErrorFlagsAddrsTraverseDeferVer:
		return "RuleStatusSemanticErrorFlagsAddrsTraverseDeferVer"
	case RuleStatusSemanticErrorFlagsAuthWithEncNegotiateOutbound:
		return "RuleStatusSemanticErrorFlagsAuthWithEncNegotiateOutbound"
	case RuleStatusSemanticErrorFlagsAuthenticateWithOutboundBypassVer:
		return "RuleStatusSemanticErrorFlagsAuthenticateWithOutboundBypassVer"
	case RuleStatusSemanticErrorRemoteAuthList:
		return "RuleStatusSemanticErrorRemoteAuthList"
	case RuleStatusSemanticErrorRemoteUserList:
		return "RuleStatusSemanticErrorRemoteUserList"
	case RuleStatusSemanticErrorLocalUserList:
		return "RuleStatusSemanticErrorLocalUserList"
	case RuleStatusSemanticErrorLUAVer:
		return "RuleStatusSemanticErrorLUAVer"
	case RuleStatusSemanticErrorLocalUserOwner:
		return "RuleStatusSemanticErrorLocalUserOwner"
	case RuleStatusSemanticErrorLocalUserOwnerVer:
		return "RuleStatusSemanticErrorLocalUserOwnerVer"
	case RuleStatusSemanticErrorLUAConditionalVer:
		return "RuleStatusSemanticErrorLUAConditionalVer"
	case RuleStatusSemanticErrorFlagsSystemOSGameOS:
		return "RuleStatusSemanticErrorFlagsSystemOSGameOS"
	case RuleStatusSemanticErrorFlagsCortanaVer:
		return "RuleStatusSemanticErrorFlagsCortanaVer"
	case RuleStatusSemanticErrorFlagsRemotename:
		return "RuleStatusSemanticErrorFlagsRemotename"
	case RuleStatusSemanticErrorFlagsAllowProfileCrossingVer:
		return "RuleStatusSemanticErrorFlagsAllowProfileCrossingVer"
	case RuleStatusSemanticErrorFlagsLocalOnlyMappedVer:
		return "RuleStatusSemanticErrorFlagsLocalOnlyMappedVer"
	case RuleStatusSemanticErrorPlatform:
		return "RuleStatusSemanticErrorPlatform"
	case RuleStatusSemanticErrorPlatformOperationVer:
		return "RuleStatusSemanticErrorPlatformOperationVer"
	case RuleStatusSemanticErrorPlatformOperation:
		return "RuleStatusSemanticErrorPlatformOperation"
	case RuleStatusSemanticErrorDTENoanyAddr:
		return "RuleStatusSemanticErrorDTENoanyAddr"
	case RuleStatusSemanticErrorTunnelExemptWithGateway:
		return "RuleStatusSemanticErrorTunnelExemptWithGateway"
	case RuleStatusSemanticErrorTunnelExemptVer:
		return "RuleStatusSemanticErrorTunnelExemptVer"
	case RuleStatusSemanticErrorAddrKeywordVer:
		return "RuleStatusSemanticErrorAddrKeywordVer"
	case RuleStatusSemanticErrorKeyModuleVer:
		return "RuleStatusSemanticErrorKeyModuleVer"
	case RuleStatusSemanticErrorAppContainerPackageID:
		return "RuleStatusSemanticErrorAppContainerPackageID"
	case RuleStatusSemanticErrorAppContainerPackageIDVer:
		return "RuleStatusSemanticErrorAppContainerPackageIDVer"
	case RuleStatusSemanticErrorTrustTupleKeywordIncompatible:
		return "RuleStatusSemanticErrorTrustTupleKeywordIncompatible"
	case RuleStatusSemanticErrorTrustTupleKeywordInvalid:
		return "RuleStatusSemanticErrorTrustTupleKeywordInvalid"
	case RuleStatusSemanticErrorTrustTupleKeywordVer:
		return "RuleStatusSemanticErrorTrustTupleKeywordVer"
	case RuleStatusSemanticErrorInterfaceTypesVer:
		return "RuleStatusSemanticErrorInterfaceTypesVer"
	case RuleStatusSemanticErrorNetNamesVer:
		return "RuleStatusSemanticErrorNetNamesVer"
	case RuleStatusSemanticErrorSecurityRealmIDVer:
		return "RuleStatusSemanticErrorSecurityRealmIDVer"
	case RuleStatusSemanticErrorSystemOSGameOSVer:
		return "RuleStatusSemanticErrorSystemOSGameOSVer"
	case RuleStatusSemanticErrorDevModeVer:
		return "RuleStatusSemanticErrorDevModeVer"
	case RuleStatusSemanticErrorRemoteServernameVer:
		return "RuleStatusSemanticErrorRemoteServernameVer"
	case RuleStatusSemanticErrorFQBNVer:
		return "RuleStatusSemanticErrorFQBNVer"
	case RuleStatusSemanticErrorCompartmentIDVer:
		return "RuleStatusSemanticErrorCompartmentIDVer"
	case RuleStatusSemanticErrorCalloutAndAuditVer:
		return "RuleStatusSemanticErrorCalloutAndAuditVer"
	case RuleStatusSemanticErrorPhase1AuthSetID:
		return "RuleStatusSemanticErrorPhase1AuthSetID"
	case RuleStatusSemanticErrorPhase2CryptoSetID:
		return "RuleStatusSemanticErrorPhase2CryptoSetID"
	case RuleStatusSemanticErrorPhase1CryptoSetID:
		return "RuleStatusSemanticErrorPhase1CryptoSetID"
	case RuleStatusSemanticErrorFlagsKeyManagerDictateVer:
		return "RuleStatusSemanticErrorFlagsKeyManagerDictateVer"
	case RuleStatusSemanticErrorFlagsKeyManagerNotifyVer:
		return "RuleStatusSemanticErrorFlagsKeyManagerNotifyVer"
	case RuleStatusSemanticErrorTransportMachineAuthzVer:
		return "RuleStatusSemanticErrorTransportMachineAuthzVer"
	case RuleStatusSemanticErrorTransportUserAuthzVer:
		return "RuleStatusSemanticErrorTransportUserAuthzVer"
	case RuleStatusSemanticErrorTransportMachineAuthzOnTunnel:
		return "RuleStatusSemanticErrorTransportMachineAuthzOnTunnel"
	case RuleStatusSemanticErrorTransportUserAuthzOnTunnel:
		return "RuleStatusSemanticErrorTransportUserAuthzOnTunnel"
	case RuleStatusSemanticErrorPerRuleAndGlobalAuthz:
		return "RuleStatusSemanticErrorPerRuleAndGlobalAuthz"
	case RuleStatusSemanticErrorFlagsSecurityRealm:
		return "RuleStatusSemanticErrorFlagsSecurityRealm"
	case RuleStatusSemanticErrorSetID:
		return "RuleStatusSemanticErrorSetID"
	case RuleStatusSemanticErrorIIIPsecPhase:
		return "RuleStatusSemanticErrorIIIPsecPhase"
	case RuleStatusSemanticErrorEmptySuites:
		return "RuleStatusSemanticErrorEmptySuites"
	case RuleStatusSemanticErrorPhase1AuthMethod:
		return "RuleStatusSemanticErrorPhase1AuthMethod"
	case RuleStatusSemanticErrorPhase2AuthMethod:
		return "RuleStatusSemanticErrorPhase2AuthMethod"
	case RuleStatusSemanticErrorAuthMethodAnonymous:
		return "RuleStatusSemanticErrorAuthMethodAnonymous"
	case RuleStatusSemanticErrorAuthMethodDuplicate:
		return "RuleStatusSemanticErrorAuthMethodDuplicate"
	case RuleStatusSemanticErrorAuthMethodVer:
		return "RuleStatusSemanticErrorAuthMethodVer"
	case RuleStatusSemanticErrorAuthSuiteFlags:
		return "RuleStatusSemanticErrorAuthSuiteFlags"
	case RuleStatusSemanticErrorHealthCert:
		return "RuleStatusSemanticErrorHealthCert"
	case RuleStatusSemanticErrorAuthSignCertVer:
		return "RuleStatusSemanticErrorAuthSignCertVer"
	case RuleStatusSemanticErrorAuthIntermediateCAVer:
		return "RuleStatusSemanticErrorAuthIntermediateCAVer"
	case RuleStatusSemanticErrorMachineSharedKey:
		return "RuleStatusSemanticErrorMachineSharedKey"
	case RuleStatusSemanticErrorCAName:
		return "RuleStatusSemanticErrorCAName"
	case RuleStatusSemanticErrorMixedCerts:
		return "RuleStatusSemanticErrorMixedCerts"
	case RuleStatusSemanticErrorNonContiguousCerts:
		return "RuleStatusSemanticErrorNonContiguousCerts"
	case RuleStatusSemanticErrorMixedCATypeInBlock:
		return "RuleStatusSemanticErrorMixedCATypeInBlock"
	case RuleStatusSemanticErrorMachineUserAuth:
		return "RuleStatusSemanticErrorMachineUserAuth"
	case RuleStatusSemanticErrorAuthCertCriteriaVer:
		return "RuleStatusSemanticErrorAuthCertCriteriaVer"
	case RuleStatusSemanticErrorAuthCertCriteriaVerMismatch:
		return "RuleStatusSemanticErrorAuthCertCriteriaVerMismatch"
	case RuleStatusSemanticErrorAuthCertCriteriaRenewalHash:
		return "RuleStatusSemanticErrorAuthCertCriteriaRenewalHash"
	case RuleStatusSemanticErrorAuthCertCriteriaInvalidHash:
		return "RuleStatusSemanticErrorAuthCertCriteriaInvalidHash"
	case RuleStatusSemanticErrorAuthCertCriteriaInvalidEKU:
		return "RuleStatusSemanticErrorAuthCertCriteriaInvalidEKU"
	case RuleStatusSemanticErrorAuthCertCriteriaInvalidNameType:
		return "RuleStatusSemanticErrorAuthCertCriteriaInvalidNameType"
	case RuleStatusSemanticErrorAuthCertCriteriaInvalidName:
		return "RuleStatusSemanticErrorAuthCertCriteriaInvalidName"
	case RuleStatusSemanticErrorAuthCertCriteriaInvalidCriteriaType:
		return "RuleStatusSemanticErrorAuthCertCriteriaInvalidCriteriaType"
	case RuleStatusSemanticErrorAuthCertCriteriaMissingCriteria:
		return "RuleStatusSemanticErrorAuthCertCriteriaMissingCriteria"
	case RuleStatusSemanticErrorProxyServer:
		return "RuleStatusSemanticErrorProxyServer"
	case RuleStatusSemanticErrorAuthProxyServerVer:
		return "RuleStatusSemanticErrorAuthProxyServerVer"
	case RuleStatusSemanticErrorPhase1CryptoNonDefaultID:
		return "RuleStatusSemanticErrorPhase1CryptoNonDefaultID"
	case RuleStatusSemanticErrorPhase1CryptoFlags:
		return "RuleStatusSemanticErrorPhase1CryptoFlags"
	case RuleStatusSemanticErrorPhase1CryptoTimeoutMinutes:
		return "RuleStatusSemanticErrorPhase1CryptoTimeoutMinutes"
	case RuleStatusSemanticErrorPhase1CryptoTimeoutSessions:
		return "RuleStatusSemanticErrorPhase1CryptoTimeoutSessions"
	case RuleStatusSemanticErrorPhase1CryptoKeyExchange:
		return "RuleStatusSemanticErrorPhase1CryptoKeyExchange"
	case RuleStatusSemanticErrorPhase1CryptoEncryption:
		return "RuleStatusSemanticErrorPhase1CryptoEncryption"
	case RuleStatusSemanticErrorPhase1CryptoHash:
		return "RuleStatusSemanticErrorPhase1CryptoHash"
	case RuleStatusSemanticErrorPhase1CryptoEncryptionVer:
		return "RuleStatusSemanticErrorPhase1CryptoEncryptionVer"
	case RuleStatusSemanticErrorPhase1CryptoHashVer:
		return "RuleStatusSemanticErrorPhase1CryptoHashVer"
	case RuleStatusSemanticErrorPhase1CryptoKeyExchVer:
		return "RuleStatusSemanticErrorPhase1CryptoKeyExchVer"
	case RuleStatusSemanticErrorPhase2CryptoPFS:
		return "RuleStatusSemanticErrorPhase2CryptoPFS"
	case RuleStatusSemanticErrorPhase2CryptoProtocol:
		return "RuleStatusSemanticErrorPhase2CryptoProtocol"
	case RuleStatusSemanticErrorPhase2CryptoEncryption:
		return "RuleStatusSemanticErrorPhase2CryptoEncryption"
	case RuleStatusSemanticErrorPhase2CryptoHash:
		return "RuleStatusSemanticErrorPhase2CryptoHash"
	case RuleStatusSemanticErrorPhase2CryptoTimeoutMinutes:
		return "RuleStatusSemanticErrorPhase2CryptoTimeoutMinutes"
	case RuleStatusSemanticErrorPhase2CryptoTimeoutKbytes:
		return "RuleStatusSemanticErrorPhase2CryptoTimeoutKbytes"
	case RuleStatusSemanticErrorPhase2CryptoEncryptionVer:
		return "RuleStatusSemanticErrorPhase2CryptoEncryptionVer"
	case RuleStatusSemanticErrorPhase2CryptoHashVer:
		return "RuleStatusSemanticErrorPhase2CryptoHashVer"
	case RuleStatusSemanticErrorPhase2CryptoPFSVer:
		return "RuleStatusSemanticErrorPhase2CryptoPFSVer"
	case RuleStatusSemanticErrorCryptoEncrHash:
		return "RuleStatusSemanticErrorCryptoEncrHash"
	case RuleStatusSemanticErrorCryptoEncrHashCompatibility:
		return "RuleStatusSemanticErrorCryptoEncrHashCompatibility"
	case RuleStatusSemanticErrorSchemaVersion:
		return "RuleStatusSemanticErrorSchemaVersion"
	case RuleStatusSemanticErrorQueryOrAndConditions:
		return "RuleStatusSemanticErrorQueryOrAndConditions"
	case RuleStatusSemanticErrorQueryAndConditions:
		return "RuleStatusSemanticErrorQueryAndConditions"
	case RuleStatusSemanticErrorQueryConditionKey:
		return "RuleStatusSemanticErrorQueryConditionKey"
	case RuleStatusSemanticErrorQueryConditionMatchType:
		return "RuleStatusSemanticErrorQueryConditionMatchType"
	case RuleStatusSemanticErrorQueryConditionDataType:
		return "RuleStatusSemanticErrorQueryConditionDataType"
	case RuleStatusSemanticErrorQueryConditionKeyAndDataType:
		return "RuleStatusSemanticErrorQueryConditionKeyAndDataType"
	case RuleStatusSemanticErrorQueryKeysProtocolPort:
		return "RuleStatusSemanticErrorQueryKeysProtocolPort"
	case RuleStatusSemanticErrorQueryKeyProfile:
		return "RuleStatusSemanticErrorQueryKeyProfile"
	case RuleStatusSemanticErrorQueryKeyStatus:
		return "RuleStatusSemanticErrorQueryKeyStatus"
	case RuleStatusSemanticErrorQueryKeyFilterID:
		return "RuleStatusSemanticErrorQueryKeyFilterID"
	case RuleStatusSemanticErrorQueryKeyAppPath:
		return "RuleStatusSemanticErrorQueryKeyAppPath"
	case RuleStatusSemanticErrorQueryKeyProtocol:
		return "RuleStatusSemanticErrorQueryKeyProtocol"
	case RuleStatusSemanticErrorQueryKeyLocalPort:
		return "RuleStatusSemanticErrorQueryKeyLocalPort"
	case RuleStatusSemanticErrorQueryKeyRemotePort:
		return "RuleStatusSemanticErrorQueryKeyRemotePort"
	case RuleStatusSemanticErrorQueryKeyServiceName:
		return "RuleStatusSemanticErrorQueryKeyServiceName"
	case RuleStatusSemanticErrorRequireInClearOutOnTransport:
		return "RuleStatusSemanticErrorRequireInClearOutOnTransport"
	case RuleStatusSemanticErrorBypassTunnelInterfaceSecureOnTransport:
		return "RuleStatusSemanticErrorBypassTunnelInterfaceSecureOnTransport"
	case RuleStatusSemanticErrorAuthNoEncapOnTunnel:
		return "RuleStatusSemanticErrorAuthNoEncapOnTunnel"
	case RuleStatusSemanticErrorAuthNoEncapOnPSK:
		return "RuleStatusSemanticErrorAuthNoEncapOnPSK"
	case RuleStatusRuntimeError:
		return "RuleStatusRuntimeError"
	case RuleStatusRuntimeErrorPhase1AuthNotFound:
		return "RuleStatusRuntimeErrorPhase1AuthNotFound"
	case RuleStatusRuntimeErrorPhase2AuthNotFound:
		return "RuleStatusRuntimeErrorPhase2AuthNotFound"
	case RuleStatusRuntimeErrorPhase2CryptoNotFound:
		return "RuleStatusRuntimeErrorPhase2CryptoNotFound"
	case RuleStatusRuntimeErrorAuthMchnSharedKeyMismatch:
		return "RuleStatusRuntimeErrorAuthMchnSharedKeyMismatch"
	case RuleStatusRuntimeErrorPhase1CryptoNotFound:
		return "RuleStatusRuntimeErrorPhase1CryptoNotFound"
	case RuleStatusRuntimeErrorAuthNoEncapOnTunnel:
		return "RuleStatusRuntimeErrorAuthNoEncapOnTunnel"
	case RuleStatusRuntimeErrorAuthNoEncapOnPSK:
		return "RuleStatusRuntimeErrorAuthNoEncapOnPSK"
	case RuleStatusRuntimeErrorKeyModuleAuthMismatch:
		return "RuleStatusRuntimeErrorKeyModuleAuthMismatch"
	case RuleStatusError:
		return "RuleStatusError"
	case RuleStatusAll:
		return "RuleStatusAll"
	}
	return "Invalid"
}

// RuleStatusClass type represents FW_RULE_STATUS_CLASS RPC enumeration.
//
// This enumeration defines classes of status codes.
type RuleStatusClass uint32

var (
	// FW_RULE_STATUS_CLASS_OK:  The rule is correctly constructed and has no issue. This
	// symbolic constant has a value of 0x00010000.
	RuleStatusClassOK RuleStatusClass = 65536
	// FW_RULE_STATUS_CLASS_PARTIALLY_IGNORED:  The rule has fields that the service can
	// successfully ignore. This symbolic constant has a value of 0x00020000.
	RuleStatusClassPartiallyIgnored RuleStatusClass = 131072
	// FW_RULE_STATUS_CLASS_IGNORED:  The rule has a higher version that the service MUST
	// ignore. This symbolic constant has a value of 0x00040000.
	RuleStatusClassIgnored RuleStatusClass = 262144
	// FW_RULE_STATUS_CLASS_PARSING_ERROR:  The rule failed to be parsed correctly. This
	// symbolic constant has a value of 0x00080000.
	RuleStatusClassParsingError RuleStatusClass = 524288
	// FW_RULE_STATUS_CLASS_SEMANTIC_ERROR:  There is a semantic error when considering
	// the fields of the rule in conjunction. This symbolic constant has a value of 0x00100000.
	RuleStatusClassSemanticError RuleStatusClass = 1048576
	// FW_RULE_STATUS_CLASS_RUNTIME_ERROR:  There is a runtime error when the object is
	// considered in conjunction with other policy objects. This symbolic constant has a
	// value of 0x00200000.
	RuleStatusClassRuntimeError RuleStatusClass = 2097152
	// FW_RULE_STATUS_CLASS_ERROR:  An error occurred. This symbolic constant has a value
	// of 0x00380000.
	RuleStatusClassError RuleStatusClass = 3670016
	// FW_RULE_STATUS_CLASS_ALL:  The status of all (used to enumerate ALL the rules, regardless
	// of the status). This symbolic constant has a value of 0xFFFF0000.
	RuleStatusClassAll RuleStatusClass = 4294901760
)

func (o RuleStatusClass) String() string {
	switch o {
	case RuleStatusClassOK:
		return "RuleStatusClassOK"
	case RuleStatusClassPartiallyIgnored:
		return "RuleStatusClassPartiallyIgnored"
	case RuleStatusClassIgnored:
		return "RuleStatusClassIgnored"
	case RuleStatusClassParsingError:
		return "RuleStatusClassParsingError"
	case RuleStatusClassSemanticError:
		return "RuleStatusClassSemanticError"
	case RuleStatusClassRuntimeError:
		return "RuleStatusClassRuntimeError"
	case RuleStatusClassError:
		return "RuleStatusClassError"
	case RuleStatusClassAll:
		return "RuleStatusClassAll"
	}
	return "Invalid"
}

// ObjectCtrlFlag type represents FW_OBJECT_CTRL_FLAG RPC enumeration.
//
// This enumeration is used to indicate the RPC protocol when elements in structures
// are included.
type ObjectCtrlFlag uint16

var (
	// FW_OBJECT_CTRL_FLAG_INCLUDE_METADATA:  This flag indicates that the structure where
	// this flag is specified contains metadata information.
	ObjectCtrlFlagIncludeMetadata ObjectCtrlFlag = 1
)

func (o ObjectCtrlFlag) String() string {
	switch o {
	case ObjectCtrlFlagIncludeMetadata:
		return "ObjectCtrlFlagIncludeMetadata"
	}
	return "Invalid"
}

// EnforcementState type represents FW_ENFORCEMENT_STATE RPC enumeration.
//
// This enumeration is part of the metadata information. It provides information about
// whether or not the policy expressed by an object is currently being enforced by the
// server.
type EnforcementState uint16

var (
	// FW_ENFORCEMENT_STATE_INVALID:  This value is invalid and MUST NOT be used by the
	// server. It is defined for simplicity in writing IDL definitions and code. This symbolic
	// constant has a value of 0.
	EnforcementStateInvalid EnforcementState = 0
	// FW_ENFORCEMENT_STATE_FULL:  The object is being enforced. This symbolic constant
	// has a value of 1.
	EnforcementStateFull EnforcementState = 1
	// FW_ENFORCEMENT_STATE_WF_OFF_IN_PROFILE:  The object is not being enforced because
	// the firewall and advanced security component is not active in a profile where the
	// object is meant to be applied. This symbolic constant has a value of 2.
	EnforcementStateWfOffInProfile EnforcementState = 2
	// FW_ENFORCEMENT_STATE_CATEGORY_OFF:  The object is not being enforced because a third-party
	// software component registered with the firewall and advanced security component to
	// own the functionality that the object is meant to perform. This symbolic constant
	// has a value of 3.
	EnforcementStateCategoryOff EnforcementState = 3
	// FW_ENFORCEMENT_STATE_DISABLED_OBJECT:  The object is not being enforced because
	// the object is disabled. This symbolic constant has a value of 4.
	EnforcementStateDisabledObject EnforcementState = 4
	// FW_ENFORCEMENT_STATE_INACTIVE_PROFILE:  The object is not being enforced because
	// at least one of the profiles that the object is meant to be applied to is not currently
	// active. This symbolic constant has a value of 5.
	EnforcementStateInactiveProfile EnforcementState = 5
	// FW_ENFORCEMENT_STATE_LOCAL_ADDRESS_RESOLUTION_EMPTY: The object is not being enforced
	// because the local address condition of the object contains a keyword that resolves
	// to an empty set. This symbolic constant has a value of 6.
	EnforcementStateLocalAddressResolutionEmpty EnforcementState = 6
	// FW_ENFORCEMENT_STATE_REMOTE_ADDRESS_RESOLUTION_EMPTY: The object is not being enforced
	// because the remote address condition of the object contains a keyword that resolves
	// to an empty set. This symbolic constant has a value of 7.
	EnforcementStateRemoteAddressResolutionEmpty EnforcementState = 7
	// FW_ENFORCEMENT_STATE_LOCAL_PORT_RESOLUTION_EMPTY: The object is not being enforced
	// because the local port condition of the object contains a keyword that resolves to
	// an empty set. This symbolic constant has a value of 8.
	EnforcementStateLocalPortResolutionEmpty EnforcementState = 8
	// FW_ENFORCEMENT_STATE_REMOTE_PORT_RESOLUTION_EMPTY: The object is not being enforced
	// because the remote port condition of the object contains a keyword that resolves
	// to an empty set. This symbolic constant has a value of 9.
	EnforcementStateRemotePortResolutionEmpty EnforcementState = 9
	// FW_ENFORCEMENT_STATE_INTERFACE_RESOLUTION_EMPTY:  The object is not being enforced
	// because the interface condition of the object contains a keyword that resolves to
	// an empty set. This symbolic constant has a value of 10.
	EnforcementStateInterfaceResolutionEmpty EnforcementState = 10
	// FW_ENFORCEMENT_STATE_APPLICATION_RESOLUTION_EMPTY: The object is not being enforced
	// because the application condition of the object contains a path that could not resolve
	// to a valid file system path. This symbolic constant has a value of 11.
	EnforcementStateApplicationResolutionEmpty EnforcementState = 11
	// FW_ENFORCEMENT_STATE_REMOTE_MACHINE_EMPTY:  The object is not being enforced because
	// the remote machine condition of the object contains an SDDL with a security identifier
	// (SID) that is not currently available on the host. This symbolic constant has a value
	// of 12.
	EnforcementStateRemoteMachineEmpty EnforcementState = 12
	// FW_ENFORCEMENT_STATE_REMOTE_USER_EMPTY:  The object is not being enforced because
	// the remote user condition of the object contains an SDDL with a SID that is not currently
	// available on the host. This symbolic constant has a value of 13.
	EnforcementStateRemoteUserEmpty EnforcementState = 13
	// FW_ENFORCEMENT_STATE_LOCAL_GLOBAL_OPEN_PORTS_DISALLOWED: The object is not being
	// enforced because the FW_PROFILE_CONFIG_AUTH_APPS_ALLOW_USER_PREF_MERGE configuration
	// option (see section 2.2.38 for more details) from a profile that the object applied
	// to, disallowed its use. This symbolic constant has a value of 14.
	EnforcementStateLocalGlobalOpenPortsDisallowed EnforcementState = 14
	// FW_ENFORCEMENT_STATE_LOCAL_AUTHORIZED_APPLICATIONS_DISALLOWED: The object is not
	// being enforced because the FW_PROFILE_CONFIG_GLOBAL_PORTS_ALLOW_USER_PREF_MERGE configuration
	// option (see section 2.2.38 for more details) from a profile that the object applied
	// to, disallowed its use. This symbolic constant has a value of 15.
	EnforcementStateLocalAuthorizedApplicationsDisallowed EnforcementState = 15
	// FW_ENFORCEMENT_STATE_LOCAL_FIREWALL_RULES_DISALLOWED: The object is not being enforced
	// because the FW_PROFILE_CONFIG_ALLOW_LOCAL_POLICY_MERGE configuration option (see
	// section 2.2.38 for more details) from a profile that the object applied to, disallowed
	// its use. This symbolic constant has a value of 16.
	EnforcementStateLocalFirewallRulesDisallowed EnforcementState = 16
	// FW_ENFORCEMENT_STATE_LOCAL_CONSEC_RULES_DISALLOWED: The object is not being enforced
	// because the FW_PROFILE_CONFIG_ALLOW_LOCAL_IPSEC_POLICY_MERGE configuration option
	// (see section 2.2.38 for more details) from a profile that the object applied to,
	// disallowed its use. This symbolic constant has a value of 17.
	EnforcementStateLocalConsecRulesDisallowed EnforcementState = 17
	// FW_ENFORCEMENT_STATE_MISMATCHED_PLATFORM:  The object is not being enforced because
	// the platform validity condition does not match the current platform of the host.
	// This symbolic constant has a value of 18.
	EnforcementStateMismatchedPlatform EnforcementState = 18
	// FW_ENFORCEMENT_STATE_OPTIMIZED_OUT:  The object is not being enforced because the
	// firewall and advanced security component determined that the object-implemented functionality
	// is irrelevant (would not change or affect what traffic is allowed or permitted) at
	// the current time. Therefore, the component optimized out the irrelevant functionality
	// and ignored it. This is a pure optimization. This symbolic constant has a value of
	// 19.
	EnforcementStateOptimizedOut EnforcementState = 19
	// FW_ENFORCEMENT_STATE_LOCAL_USER_EMPTY:  The object is not being enforced, because
	// the local user condition of the object contains an SDDL with a SID that is not currently
	// available on the host. For schema versions 0x0200, 0x0201, and 0x020A, this value
	// is invalid and MUST NOT be used. This symbolic constant has a value of 20.
	EnforcementStateLocalUserEmpty EnforcementState = 20
	// FW_ENFORCEMENT_STATE_TRANSPORT_MACHINE_SD_EMPTY:  The object is not being enforced
	// because the IPsec transport mode machine authorization list contains an SDDL with
	// a SID that is not currently available on the host. For schema versions 0x0200, 0x0201,
	// and 0x020A, this value is invalid and MUST NOT be used. This symbolic constant has
	// a value of 21.
	EnforcementStateTransportMachineSDEmpty EnforcementState = 21
	// FW_ENFORCEMENT_STATE_TRANSPORT_USER_SD_EMPTY:  The object is not being enforced,
	// because the IPsec transport mode user authorization list contains an SDDL with a
	// SID that is not currently available on the host. For schema versions 0x0200, 0x0201,
	// and 0x020A, this value is invalid and MUST NOT be used. This symbolic constant has
	// a value of 22.
	EnforcementStateTransportUserSDEmpty EnforcementState = 22
	// FW_ENFORCEMENT_STATE_TUPLE_RESOLUTION_EMPTY:  The object is not being enforced,
	// because the trust tuple keywords resolve to an empty set. For schema versions 0x0200,
	// 0x0201, and 0x020A, this value is invalid and MUST NOT be used. This symbolic constant
	// has a value of 23.
	EnforcementStateTupleResolutionEmpty       EnforcementState = 23
	EnforcementStateNetworkNameResolutionEmpty EnforcementState = 24
	// FW_ENFORCEMENT_STATE_MAX:  This value and values that exceed this value are not
	// valid and MUST NOT be used. It is defined for simplicity in writing IDL definitions
	// and code. This symbolic constant has a value of 25.
	EnforcementStateMax EnforcementState = 25
)

func (o EnforcementState) String() string {
	switch o {
	case EnforcementStateInvalid:
		return "EnforcementStateInvalid"
	case EnforcementStateFull:
		return "EnforcementStateFull"
	case EnforcementStateWfOffInProfile:
		return "EnforcementStateWfOffInProfile"
	case EnforcementStateCategoryOff:
		return "EnforcementStateCategoryOff"
	case EnforcementStateDisabledObject:
		return "EnforcementStateDisabledObject"
	case EnforcementStateInactiveProfile:
		return "EnforcementStateInactiveProfile"
	case EnforcementStateLocalAddressResolutionEmpty:
		return "EnforcementStateLocalAddressResolutionEmpty"
	case EnforcementStateRemoteAddressResolutionEmpty:
		return "EnforcementStateRemoteAddressResolutionEmpty"
	case EnforcementStateLocalPortResolutionEmpty:
		return "EnforcementStateLocalPortResolutionEmpty"
	case EnforcementStateRemotePortResolutionEmpty:
		return "EnforcementStateRemotePortResolutionEmpty"
	case EnforcementStateInterfaceResolutionEmpty:
		return "EnforcementStateInterfaceResolutionEmpty"
	case EnforcementStateApplicationResolutionEmpty:
		return "EnforcementStateApplicationResolutionEmpty"
	case EnforcementStateRemoteMachineEmpty:
		return "EnforcementStateRemoteMachineEmpty"
	case EnforcementStateRemoteUserEmpty:
		return "EnforcementStateRemoteUserEmpty"
	case EnforcementStateLocalGlobalOpenPortsDisallowed:
		return "EnforcementStateLocalGlobalOpenPortsDisallowed"
	case EnforcementStateLocalAuthorizedApplicationsDisallowed:
		return "EnforcementStateLocalAuthorizedApplicationsDisallowed"
	case EnforcementStateLocalFirewallRulesDisallowed:
		return "EnforcementStateLocalFirewallRulesDisallowed"
	case EnforcementStateLocalConsecRulesDisallowed:
		return "EnforcementStateLocalConsecRulesDisallowed"
	case EnforcementStateMismatchedPlatform:
		return "EnforcementStateMismatchedPlatform"
	case EnforcementStateOptimizedOut:
		return "EnforcementStateOptimizedOut"
	case EnforcementStateLocalUserEmpty:
		return "EnforcementStateLocalUserEmpty"
	case EnforcementStateTransportMachineSDEmpty:
		return "EnforcementStateTransportMachineSDEmpty"
	case EnforcementStateTransportUserSDEmpty:
		return "EnforcementStateTransportUserSDEmpty"
	case EnforcementStateTupleResolutionEmpty:
		return "EnforcementStateTupleResolutionEmpty"
	case EnforcementStateNetworkNameResolutionEmpty:
		return "EnforcementStateNetworkNameResolutionEmpty"
	case EnforcementStateMax:
		return "EnforcementStateMax"
	}
	return "Invalid"
}

// ObjectMetadata structure represents FW_OBJECT_METADATA RPC structure.
//
// This structure contains the metadata that is associated with a specific policy object.
type ObjectMetadata struct {
	// qwFilterContextID:  This field is not used across the wires.
	FilterContextID uint64 `idl:"name:qwFilterContextID" json:"filter_context_id"`
	// dwNumEntries:  A field that specifies the number of metadata hints (FW_ENFORCEMENT_STATEs)
	// that the structure contains.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// pEnforcementStates:  A pointer to an array of FW_ENFORCEMENT_STATE elements. The
	// number of elements is given by dwNumEntries.
	EnforcementStates []EnforcementState `idl:"name:pEnforcementStates;size_is:(dwNumEntries)" json:"enforcement_states"`
}

func (o *ObjectMetadata) xxx_PreparePayload(ctx context.Context) error {
	if o.EnforcementStates != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.EnforcementStates))
	}
	if o.EntriesLength > uint32(100) {
		return fmt.Errorf("EntriesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectMetadata) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.FilterContextID); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.EnforcementStates != nil || o.EntriesLength > 0 {
		_ptr_pEnforcementStates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.EnforcementStates {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(uint16(o.EnforcementStates[i1])); err != nil {
					return err
				}
			}
			for i1 := len(o.EnforcementStates); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EnforcementStates, _ptr_pEnforcementStates); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectMetadata) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.FilterContextID); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_pEnforcementStates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.EnforcementStates", sizeInfo[0])
		}
		o.EnforcementStates = make([]EnforcementState, sizeInfo[0])
		for i1 := range o.EnforcementStates {
			i1 := i1
			if err := w.ReadData((*uint16)(&o.EnforcementStates[i1])); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pEnforcementStates := func(ptr interface{}) { o.EnforcementStates = *ptr.(*[]EnforcementState) }
	if err := w.ReadPointer(&o.EnforcementStates, _s_pEnforcementStates, _ptr_pEnforcementStates); err != nil {
		return err
	}
	return nil
}

// OSPlatformOperation type represents FW_OS_PLATFORM_OP RPC enumeration.
//
// This enumeration describes the operations used in the FW_OS_PLATFORM structure to
// determine if an object should be applied to a specified operating system platform.
type OSPlatformOperation uint16

var (
	// FW_OS_PLATFORM_OP_EQ:  The operating system platform MUST be the same as the one
	// specified. This is satisfied when the following occurs:
	//
	// If ( ((bPlatform & 0x7) == platform type) && (bMajorVersion == major version) &&
	// (bMinorVersion == minor version) ).
	OSPlatformOperationEQ OSPlatformOperation = 0
	// FW_OS_PLATFORM_OP_GTEQ:  The operating system MUST be greater than or equal to the
	// one specified. This is satisfied when any of the following occur:
	//
	// If (bPlatform & 0x7) > platform type
	//
	// If (((bPlatform & 0x7) == platform type) && (bMajorVersion > major version))
	//
	// If (((bPlatform & 0x7) == platform type) && (bMajorVersion == major version) && (bMinorVersion
	// >= minor version))
	OSPlatformOperationGTEQ OSPlatformOperation = 1
	// FW_OS_PLATFORM_OP_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 2.
	OSPlatformOperationMax       OSPlatformOperation = 2
	OSPlatformOperationFieldSize OSPlatformOperation = 5
	OSPlatformOperationFieldMask OSPlatformOperation = 248
)

func (o OSPlatformOperation) String() string {
	switch o {
	case OSPlatformOperationEQ:
		return "OSPlatformOperationEQ"
	case OSPlatformOperationGTEQ:
		return "OSPlatformOperationGTEQ"
	case OSPlatformOperationMax:
		return "OSPlatformOperationMax"
	case OSPlatformOperationFieldSize:
		return "OSPlatformOperationFieldSize"
	case OSPlatformOperationFieldMask:
		return "OSPlatformOperationFieldMask"
	}
	return "Invalid"
}

// OSPlatform structure represents FW_OS_PLATFORM RPC structure.
//
// This structure describes a set of operating system platforms. The fields in this
// data type correspond to the fields of the OSVERSIONINFOEX data type (for more information,
// see [MSDN-OSVERSIONINFOEX]). There are no constraints on the values allowed for the
// platform type, major version, or minor version. The set can include values that do
// not correspond to any existing operating system platform.
type OSPlatform struct {
	// bPlatform:  The three least significant bits identify the platform type. This corresponds
	// to the dwPlatformId field in MSDN. The five most significant bits contain a value
	// from the FW_OS_PLATFORM_OP enumeration.
	Platform uint8 `idl:"name:bPlatform" json:"platform"`
	// bMajorVersion:  Specifies the major version number for the OS. This corresponds to
	// the dwMajorVersion field in MSDN.
	MajorVersion uint8 `idl:"name:bMajorVersion" json:"major_version"`
	// bMinorVersion:  Specifies the minor version number for the OS. This corresponds to
	// the dwMinorVersion field in MSDN.
	MinorVersion uint8 `idl:"name:bMinorVersion" json:"minor_version"`
	// Reserved:  Not used. Reserved for future use.
	_ uint8 `idl:"name:Reserved"`
}

func (o *OSPlatform) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OSPlatform) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteData(o.Platform); err != nil {
		return err
	}
	if err := w.WriteData(o.MajorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.MinorVersion); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint8(0)); err != nil {
		return err
	}
	return nil
}
func (o *OSPlatform) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Platform); err != nil {
		return err
	}
	if err := w.ReadData(&o.MajorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinorVersion); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint8
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	return nil
}

// OSPlatformList structure represents FW_OS_PLATFORM_LIST RPC structure.
//
// This structure contains an array of FW_OS_PLATFORM elements. The structure describes
// a set of operating system platforms. This set is the union of the sets identified
// by each FW_OS_PLATFORM element.
type OSPlatformList struct {
	// dwNumEntries:  This field specifies the number of OS platforms that the structure
	// contains.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// pPlatforms:  A pointer to an array of dwNumEntries contiguous FW_OS_PLATFORM elements.
	Platforms []*OSPlatform `idl:"name:pPlatforms;size_is:(dwNumEntries)" json:"platforms"`
}

func (o *OSPlatformList) xxx_PreparePayload(ctx context.Context) error {
	if o.Platforms != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Platforms))
	}
	if o.EntriesLength > uint32(10000) {
		return fmt.Errorf("EntriesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OSPlatformList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.Platforms != nil || o.EntriesLength > 0 {
		_ptr_pPlatforms := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Platforms {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Platforms[i1] != nil {
					if err := o.Platforms[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&OSPlatform{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Platforms); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&OSPlatform{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Platforms, _ptr_pPlatforms); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *OSPlatformList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_pPlatforms := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Platforms", sizeInfo[0])
		}
		o.Platforms = make([]*OSPlatform, sizeInfo[0])
		for i1 := range o.Platforms {
			i1 := i1
			if o.Platforms[i1] == nil {
				o.Platforms[i1] = &OSPlatform{}
			}
			if err := o.Platforms[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pPlatforms := func(ptr interface{}) { o.Platforms = *ptr.(*[]*OSPlatform) }
	if err := w.ReadPointer(&o.Platforms, _s_pPlatforms, _ptr_pPlatforms); err != nil {
		return err
	}
	return nil
}

// NetworkNames structure represents FW_NETWORK_NAMES RPC structure.
//
// The FW_NETWORK_NAMES structure represents a firewall rule that is used by the 2.24
// binary version servers and clients (see sections 2.2.41 and 2.2.42).
type NetworkNames struct {
	// dwNumEntries:  Specifies the number of network names in the wszNames array.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// wszNames:  An array of pointers to null-terminated Unicode strings representing the network DNS suffix as specified in the network interface DNS suffix. Each pointer string MUST NOT be NULL , the string MUST NOT contain the pipe (|) character, MUST be a string at least 1 character long, and MUST NOT be greater than or equal to 255 characters.
	Names []string `idl:"name:wszNames;size_is:(dwNumEntries, );string;pointer:unique" json:"names"`
}

func (o *NetworkNames) xxx_PreparePayload(ctx context.Context) error {
	if o.Names != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.Names))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NetworkNames) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.Names != nil || o.EntriesLength > 0 {
		_ptr_wszNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
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
				if o.Names[i1] != "" {
					_ptr_wszNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.Names[i1]); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.Names[i1], _ptr_wszNames); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Names, _ptr_wszNames); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NetworkNames) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_wszNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]string, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			_ptr_wszNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.Names[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_wszNames := func(ptr interface{}) { o.Names[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.Names[i1], _s_wszNames, _ptr_wszNames); err != nil {
				return err
			}
		}
		return nil
	})
	_s_wszNames := func(ptr interface{}) { o.Names = *ptr.(*[]string) }
	if err := w.ReadPointer(&o.Names, _s_wszNames, _ptr_wszNames); err != nil {
		return err
	}
	return nil
}

// RuleOriginType type represents FW_RULE_ORIGIN_TYPE RPC enumeration.
//
// This enumeration represents where the policy object is stored and from where it originates.
type RuleOriginType uint16

var (
	// FW_RULE_ORIGIN_INVALID:  On enumeration, this value is invalid, and MUST NOT be
	// used by the server. It is defined for simplicity in writing IDL definitions and code.
	// However, the server ignores the fields of this data type on input, and hence it is
	// valid for filling rules. This symbolic constant has a value of 0.
	RuleOriginTypeInvalid RuleOriginType = 0
	// FW_RULE_ORIGIN_LOCAL:  Specifies that the policy object originates from the local
	// store. This symbolic constant has a value of 1.
	RuleOriginTypeLocal RuleOriginType = 1
	// FW_RULE_ORIGIN_GP:  Specifies that the policy object originates from the GP store.
	// This symbolic constant has a value of 2.
	RuleOriginTypeGP RuleOriginType = 2
	// FW_RULE_ORIGIN_DYNAMIC:  Specifies that the policy object originates from the dynamic
	// store. This symbolic constant has a value of 3.
	RuleOriginTypeDynamic RuleOriginType = 3
	// FW_RULE_ORIGIN_AUTOGEN:  Not used. This symbolic constant has a value of 4.
	RuleOriginTypeAutogen RuleOriginType = 4
	// FW_RULE_ORIGIN_HARDCODED:  Specifies that the policy object originates from the
	// firewall and advanced security component hard-coded values and is used due to lack
	// of user settings. These values are not configurable and are not addressed in this
	// protocol specification. Specific implementations of firewall and advanced security
	// components can choose what hard-coded values to use when no other user settings are
	// available. The only policy objects in this protocol specification that can have this
	// FW_RULE_ORIGIN_HARDCODED value assigned are authentication sets and cryptographic
	// sets, which are defined in sections 2.2.65 and 2.2.74, respectively.<5> This symbolic
	// constant has a value of 5.
	RuleOriginTypeHardcoded RuleOriginType = 5
	// FW_RULE_ORIGIN_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 6.
	RuleOriginTypeMax RuleOriginType = 6
)

func (o RuleOriginType) String() string {
	switch o {
	case RuleOriginTypeInvalid:
		return "RuleOriginTypeInvalid"
	case RuleOriginTypeLocal:
		return "RuleOriginTypeLocal"
	case RuleOriginTypeGP:
		return "RuleOriginTypeGP"
	case RuleOriginTypeDynamic:
		return "RuleOriginTypeDynamic"
	case RuleOriginTypeAutogen:
		return "RuleOriginTypeAutogen"
	case RuleOriginTypeHardcoded:
		return "RuleOriginTypeHardcoded"
	case RuleOriginTypeMax:
		return "RuleOriginTypeMax"
	}
	return "Invalid"
}

// EnumRulesFlags type represents FW_ENUM_RULES_FLAGS RPC enumeration.
//
// This enumeration defines flag values that can be used in the enumeration methods
// that are defined in RRPC_FWEnumFirewallRules, RRPC_FWEnumConnectionSecurityRules,
// RRPC_FWEnumAuthenticationSets, and RRPC_FWEnumCryptoSets.
type EnumRulesFlags uint16

var (
	// FW_ENUM_RULES_FLAG_NONE:  This value signifies that no specific flag is used. It
	// is defined for IDL definitions and code to add readability, instead of using the
	// number 0. This symbolic constant has a value 0x0000.
	EnumRulesFlagsNone EnumRulesFlags = 0
	// FW_ENUM_RULES_FLAG_RESOLVE_NAME:  Resolves rule description strings to user-friendly,
	// localizable strings if they are in the following format: @file.dll,-<resID>. resID
	// refers to the resource ID in the indirect string. Please see [MSDN-SHLoadIndirectString]
	// for further documentation on the string format. This symbolic constant has a value
	// 0x0001.
	EnumRulesFlagsResolveName EnumRulesFlags = 1
	// FW_ENUM_RULES_FLAG_RESOLVE_DESCRIPTION:  Resolves rule description strings to user-friendly,
	// localizable strings if they are in the following format: @file.dll,-<resID>. resID
	// refers to the resource ID in the indirect string. Please see [MSDN-SHLoadIndirectString]
	// for further documentation on the string format. This symbolic constant has a value
	// 0x0002.
	EnumRulesFlagsResolveDescription EnumRulesFlags = 2
	// FW_ENUM_RULES_FLAG_RESOLVE_APPLICATION:  If this flag is set, the server MUST inspect
	// the wszLocalApplication field of each FW_RULE structure and replace all environment
	// variables in the string with their corresponding values. See [MSDN-ExpandEnvironmentStrings]
	// for more details about environment-variable strings. This symbolic constant has a
	// value 0x0004.
	EnumRulesFlagsResolveApplication EnumRulesFlags = 4
	// FW_ENUM_RULES_FLAG_RESOLVE_KEYWORD:  Resolves keywords in addresses and ports to
	// the actual addresses and ports (dynamic store only). This symbolic constant has a
	// value 0x0008.
	EnumRulesFlagsResolveKeyword EnumRulesFlags = 8
	// FW_ENUM_RULES_FLAG_RESOLVE_GPO_NAME:  Resolves the GPO name for the GP_RSOP rules.
	// This symbolic constant has a value 0x0010.
	EnumRulesFlagsResolveGPOName EnumRulesFlags = 16
	// FW_ENUM_RULES_FLAG_EFFECTIVE:  If this flag is set, the server MUST only return
	// objects where at least one FW_ENFORCEMENT_STATE entry in the object's metadata is
	// equal to FW_ENFORCEMENT_STATE_FULL. This flag is available for the dynamic store
	// only. This symbolic constant has a value 0x0020.
	EnumRulesFlagsEffective EnumRulesFlags = 32
	// FW_ENUM_RULES_FLAG_INCLUDE_METADATA:  Includes the metadata object information,
	// represented by the FW_OBJECT_METADATA structure, in the enumerated objects. This
	// symbolic constant has a value 0x0040.
	EnumRulesFlagsIncludeMetadata EnumRulesFlags = 64
	// FW_ENUM_RULES_FLAG_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value 0x0080.
	EnumRulesFlagsMax EnumRulesFlags = 128
)

func (o EnumRulesFlags) String() string {
	switch o {
	case EnumRulesFlagsNone:
		return "EnumRulesFlagsNone"
	case EnumRulesFlagsResolveName:
		return "EnumRulesFlagsResolveName"
	case EnumRulesFlagsResolveDescription:
		return "EnumRulesFlagsResolveDescription"
	case EnumRulesFlagsResolveApplication:
		return "EnumRulesFlagsResolveApplication"
	case EnumRulesFlagsResolveKeyword:
		return "EnumRulesFlagsResolveKeyword"
	case EnumRulesFlagsResolveGPOName:
		return "EnumRulesFlagsResolveGPOName"
	case EnumRulesFlagsEffective:
		return "EnumRulesFlagsEffective"
	case EnumRulesFlagsIncludeMetadata:
		return "EnumRulesFlagsIncludeMetadata"
	case EnumRulesFlagsMax:
		return "EnumRulesFlagsMax"
	}
	return "Invalid"
}

// RuleAction type represents FW_RULE_ACTION RPC enumeration.
//
// This enumeration describes the possible actions on firewall rules.
type RuleAction uint16

var (
	// FW_RULE_ACTION_INVALID:  This value is invalid and MUST NOT be used. It is defined
	// for simplicity in writing IDL definitions and code. This symbolic constant has a
	// value of 0.
	RuleActionInvalid RuleAction = 0
	// FW_RULE_ACTION_ALLOW_BYPASS:  Rules with this action allow traffic but are applicable
	// only to rules that at least specify the FW_RULE_FLAGS_AUTHENTICATE flag. This symbolic
	// constant has a value of 1.
	RuleActionAllowBypass RuleAction = 1
	// FW_RULE_ACTION_BLOCK:  Rules with this action block traffic. This symbolic constant
	// has a value of 2.
	RuleActionBlock RuleAction = 2
	// FW_RULE_ACTION_ALLOW:  Rules with this action allow traffic. This symbolic constant
	// has a value of 3.
	RuleActionAllow RuleAction = 3
	// FW_RULE_ACTION_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 4.
	//
	// If conflicting rules match the same network traffic, the actions determine the order
	// of precedence. Allow-bypass rules take precedence over block rules, and block rules
	// take precedence over allow rules.
	RuleActionMax RuleAction = 4
)

func (o RuleAction) String() string {
	switch o {
	case RuleActionInvalid:
		return "RuleActionInvalid"
	case RuleActionAllowBypass:
		return "RuleActionAllowBypass"
	case RuleActionBlock:
		return "RuleActionBlock"
	case RuleActionAllow:
		return "RuleActionAllow"
	case RuleActionMax:
		return "RuleActionMax"
	}
	return "Invalid"
}

// RuleFlags type represents FW_RULE_FLAGS RPC enumeration.
//
// This enumeration represents flags that can be specified in firewall rules of section
// 2.2.37.
type RuleFlags uint16

var (
	// FW_RULE_FLAGS_NONE:  This value means that none of the following flags are set.
	// It is defined for simplicity in writing IDL definitions and code.
	RuleFlagsNone RuleFlags = 0
	// FW_RULE_FLAGS_ACTIVE:  The rule is enabled if this flag is set; otherwise, it is
	// disabled.
	RuleFlagsActive RuleFlags = 1
	// FW_RULE_FLAGS_AUTHENTICATE:  This flag MUST be set only on rules that have the allow
	// actions. If set, traffic that matches the rule is allowed only if it has been authenticated
	// by IPsec; otherwise, traffic is blocked.
	RuleFlagsAuthenticate RuleFlags = 2
	// FW_RULE_FLAGS_AUTHENTICATE_WITH_ENCRYPTION:  This flag is similar to the FW_RULE_FLAGS_AUTHENTICATE
	// flag; however, traffic MUST also be encrypted.
	RuleFlagsAuthenticateWithEncryption RuleFlags = 4
	// FW_RULE_FLAGS_ROUTEABLE_ADDRS_TRAVERSE:  This flag MUST be set only on inbound rules.
	// This flag allows the matching traffic to traverse a NAT edge device and be allowed
	// in the host computer.
	RuleFlagsRouteableAddrsTraverse RuleFlags = 8
	// FW_RULE_FLAGS_LOOSE_SOURCE_MAPPED:  This flag allows responses from a remote IP
	// address that is different from the one to which the outbound matched traffic originally
	// went.
	RuleFlagsLooseSourceMapped RuleFlags = 16
	// FW_RULE_FLAGS_MAX_V2_1:  This value and values that exceed this value are not valid
	// and MUST NOT be used by servers and clients with schema version 0x0201 and earlier.
	// It is defined for simplicity in writing IDL definitions and code. This symbolic constant
	// has a value of 0x0020.
	RuleFlagsMaxV21 RuleFlags = 32
	// FW_RULE_FLAGS_AUTH_WITH_NO_ENCAPSULATION:  This flag MUST be set only on rules that
	// have the FW_RULE_FLAGS_AUTHENTICATE flag set. If set, traffic that matches the rule
	// is allowed if IKE or AuthIP authentication was successful; however, this flag does
	// not necessarily require that traffic be protected by IPsec encapsulations. For schema
	// versions 0x0200 and 0x0201, this value is invalid and MUST NOT be used.
	RuleFlagsAuthWithNoEncapsulation RuleFlags = 32
	// FW_RULE_FLAGS_MAX_V2_9:  This value and values that exceed this value are not valid
	// and MUST NOT be used by servers and clients with schema version 0x0209 and earlier.
	// It is defined for simplicity in writing IDL definitions and code. This symbolic constant
	// has a value of 0x0040.
	RuleFlagsMaxV29 RuleFlags = 64
	// FW_RULE_FLAGS_AUTH_WITH_ENC_NEGOTIATE:  This flag MUST be set only on inbound rules
	// that have the FW_RULE_FLAGS_AUTHENTICATE_WITH_ENCRYPTION flag set. If set and if
	// the first packet that arrives is unencrypted but authenticated by IPsec, the packet
	// is allowed, and an IKE or AuthIP negotiation is started to negotiate encryption settings
	// and encrypt subsequent packets. [MS-AIPS] section 3.2.4 specifies negotiation initiation
	// behavior for hosts that support both IKE and AuthIP negotiation. If the negotiation
	// fails, the connection is dropped. For schema versions 0x0200 and 0x0201, this value
	// is invalid and MUST NOT be used.
	RuleFlagsAuthWithEncNegotiate RuleFlags = 64
	// FW_RULE_FLAGS_ROUTEABLE_ADDRS_TRAVERSE_DEFER_APP: This flag MUST be set only on inbound
	// rules. This flag allows the matching traffic to traverse a NAT edge device and be
	// allowed in the host computer, if and only if a matching PortInUse object is found
	// in the PortsInUse collection with NATTraversalRequested set to true (see section
	// 3.1.1). For schema versions 0x0200 and 0x0201, this value is invalid and MUST NOT
	// be used.
	RuleFlagsRouteableAddrsTraverseDeferApp RuleFlags = 128
	// FW_RULE_FLAGS_ROUTEABLE_ADDRS_TRAVERSE_DEFER_USER: This flag MUST be set only on
	// inbound rules. Whenever an application tries to listen for traffic that matches this
	// rule, the operating system asks the administrator of the host whether it should allow
	// this traffic to traverse the NAT. For schema versions 0x0200 and 0x0201, this value
	// is invalid and MUST NOT be used.
	RuleFlagsRouteableAddrsTraverseDeferUser RuleFlags = 256
	// FW_RULE_FLAGS_AUTHENTICATE_BYPASS_OUTBOUND:  This flag MUST be set only on outbound
	// rules that have an allow action with either the FW_RULE_FLAGS_AUTHENTICATE or the
	// FW_RULE_FLAGS_AUTHENTICATE_WITH_ENCRYPTION flag set. If set, this rule is evaluated
	// before block rules, making it equivalent to a rule with an FW_RULE_ACTION_ALLOW_BYPASS,
	// but for outbound. For schema versions 0x0200 and 0x0201, this value is invalid and
	// MUST NOT be used.
	RuleFlagsAuthenticateBypassOutbound RuleFlags = 512
	// FW_RULE_FLAGS_MAX_V2_10:  This value and values that exceed this value are not valid
	// and MUST NOT be used by servers and clients with schema version 0x020A and earlier.
	// It is defined for simplicity in writing IDL definitions and code. This symbolic constant
	// has a value of 0x0400.
	RuleFlagsMaxV210 RuleFlags = 1024
	// FW_RULE_FLAGS_ALLOW_PROFILE_CROSSING:  This flag allows responses from a network
	// with a different profile type than the network to which the outbound traffic was
	// originally sent. This flag MUST be ignored on rules with an action of FW_RULE_ACTION_BLOCK.
	// For schema versions 0x0200, 0x0201, and 0x020A, this value is invalid and MUST NOT
	// be used.
	RuleFlagsAllowProfileCrossing RuleFlags = 1024
	// FW_RULE_FLAGS_LOCAL_ONLY_MAPPED:  If this flag is set on a rule, the remote address
	// and remote port conditions are ignored when determining whether a network traffic
	// flow matches the rule. This flag MUST be ignored on rules with an action of FW_RULE_ACTION_BLOCK.
	// For schema versions 0x0200, 0x0201, and 0x020A, this value is invalid and MUST NOT
	// be used.
	RuleFlagsLocalOnlyMapped RuleFlags = 2048
	// FW_RULE_FLAGS_MAX_V2_20:  This value and values that exceed this value are not valid
	// and MUST NOT be used by servers and clients with schema version 0x0214 and earlier.
	// It is defined for simplicity in writing IDL definitions and code. This symbolic constant
	// has a value of 0x1000.
	RuleFlagsMaxV220 RuleFlags = 4096
	// FW_RULE_FLAGS_LUA_CONDITIONAL_ACE:  This flag MUST be set if and only if the wszLocalUserAuthorizationList
	// field of the FW_RULE2_24 structure (section 2.2.104) is to include conditional ACEs.
	// For schema versions 0x0200, 0x0201, 0x020A, 0x0214, and 0x0216, this value is invalid
	// and MUST NOT be used.
	RuleFlagsLUAConditionalACE RuleFlags = 4096
	// FW_RULE_FLAGS_BIND_TO_INTERFACE: This flag is not used.
	RuleFlagsBindToInterface RuleFlags = 8192
	// FW_RULE_FLAGS_MAX:  This value and values that exceed this value are not valid and
	// MUST NOT be used. It is defined for simplicity in writing IDL definitions and code.
	// This symbolic constant has a value of 0x4000.
	RuleFlagsMax RuleFlags = 16384
)

func (o RuleFlags) String() string {
	switch o {
	case RuleFlagsNone:
		return "RuleFlagsNone"
	case RuleFlagsActive:
		return "RuleFlagsActive"
	case RuleFlagsAuthenticate:
		return "RuleFlagsAuthenticate"
	case RuleFlagsAuthenticateWithEncryption:
		return "RuleFlagsAuthenticateWithEncryption"
	case RuleFlagsRouteableAddrsTraverse:
		return "RuleFlagsRouteableAddrsTraverse"
	case RuleFlagsLooseSourceMapped:
		return "RuleFlagsLooseSourceMapped"
	case RuleFlagsMaxV21:
		return "RuleFlagsMaxV21"
	case RuleFlagsAuthWithNoEncapsulation:
		return "RuleFlagsAuthWithNoEncapsulation"
	case RuleFlagsMaxV29:
		return "RuleFlagsMaxV29"
	case RuleFlagsAuthWithEncNegotiate:
		return "RuleFlagsAuthWithEncNegotiate"
	case RuleFlagsRouteableAddrsTraverseDeferApp:
		return "RuleFlagsRouteableAddrsTraverseDeferApp"
	case RuleFlagsRouteableAddrsTraverseDeferUser:
		return "RuleFlagsRouteableAddrsTraverseDeferUser"
	case RuleFlagsAuthenticateBypassOutbound:
		return "RuleFlagsAuthenticateBypassOutbound"
	case RuleFlagsMaxV210:
		return "RuleFlagsMaxV210"
	case RuleFlagsAllowProfileCrossing:
		return "RuleFlagsAllowProfileCrossing"
	case RuleFlagsLocalOnlyMapped:
		return "RuleFlagsLocalOnlyMapped"
	case RuleFlagsMaxV220:
		return "RuleFlagsMaxV220"
	case RuleFlagsLUAConditionalACE:
		return "RuleFlagsLUAConditionalACE"
	case RuleFlagsBindToInterface:
		return "RuleFlagsBindToInterface"
	case RuleFlagsMax:
		return "RuleFlagsMax"
	}
	return "Invalid"
}

// RuleFlags2 type represents FW_RULE_FLAGS2 RPC enumeration.
type RuleFlags2 uint16

var (
	RuleFlags2None             RuleFlags2 = 0
	RuleFlags2SystemOSOnly     RuleFlags2 = 1
	RuleFlags2GameOSOnly       RuleFlags2 = 2
	RuleFlags2DevMode          RuleFlags2 = 4
	RuleFlags2MaxV226          RuleFlags2 = 8
	RuleFlags2NotUsedValue8    RuleFlags2 = 8
	RuleFlags2NotUsedValue16   RuleFlags2 = 16
	RuleFlags2NotUsedValue32   RuleFlags2 = 32
	RuleFlags2NotUsedValue64   RuleFlags2 = 64
	RuleFlags2CalloutAndAudit  RuleFlags2 = 128
	RuleFlags2NotUsedValue256  RuleFlags2 = 256
	RuleFlags2NotUsedValue512  RuleFlags2 = 512
	RuleFlags2NotUsedValue1024 RuleFlags2 = 1024
	RuleFlags2Max              RuleFlags2 = 2048
)

func (o RuleFlags2) String() string {
	switch o {
	case RuleFlags2None:
		return "RuleFlags2None"
	case RuleFlags2SystemOSOnly:
		return "RuleFlags2SystemOSOnly"
	case RuleFlags2GameOSOnly:
		return "RuleFlags2GameOSOnly"
	case RuleFlags2DevMode:
		return "RuleFlags2DevMode"
	case RuleFlags2MaxV226:
		return "RuleFlags2MaxV226"
	case RuleFlags2NotUsedValue8:
		return "RuleFlags2NotUsedValue8"
	case RuleFlags2NotUsedValue16:
		return "RuleFlags2NotUsedValue16"
	case RuleFlags2NotUsedValue32:
		return "RuleFlags2NotUsedValue32"
	case RuleFlags2NotUsedValue64:
		return "RuleFlags2NotUsedValue64"
	case RuleFlags2CalloutAndAudit:
		return "RuleFlags2CalloutAndAudit"
	case RuleFlags2NotUsedValue256:
		return "RuleFlags2NotUsedValue256"
	case RuleFlags2NotUsedValue512:
		return "RuleFlags2NotUsedValue512"
	case RuleFlags2NotUsedValue1024:
		return "RuleFlags2NotUsedValue1024"
	case RuleFlags2Max:
		return "RuleFlags2Max"
	}
	return "Invalid"
}

// Rule20 structure represents FW_RULE2_0 RPC structure.
//
// This structure represents a firewall rule that is used by the 2.0 binary version
// servers and clients (see section 2.2.42). The fields of this structure are identical
// to the FW_RULE structure and its meanings are covered in section 2.2.37.
type Rule20 struct {
	Next                           *Rule20                `idl:"name:pNext" json:"next"`
	SchemaVersion                  uint16                 `idl:"name:wSchemaVersion" json:"schema_version"`
	RuleID                         string                 `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	Name                           string                 `idl:"name:wszName;string" json:"name"`
	Description                    string                 `idl:"name:wszDescription;string" json:"description"`
	Profiles                       uint32                 `idl:"name:dwProfiles" json:"profiles"`
	Direction                      Direction              `idl:"name:Direction" json:"direction"`
	IPProtocol                     uint16                 `idl:"name:wIpProtocol" json:"ip_protocol"`
	IPProtocolData                 *Rule20_IPProtocolData `idl:"name:IpProtocolData;switch_is:wIpProtocol" json:"ip_protocol_data"`
	LocalAddresses                 *Addresses             `idl:"name:LocalAddresses" json:"local_addresses"`
	RemoteAddresses                *Addresses             `idl:"name:RemoteAddresses" json:"remote_addresses"`
	LocalInterfaceIDs              *InterfaceLUIDs        `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	LocalInterfaceTypes            uint32                 `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	LocalApplication               string                 `idl:"name:wszLocalApplication;string" json:"local_application"`
	LocalService                   string                 `idl:"name:wszLocalService;string" json:"local_service"`
	Action                         RuleAction             `idl:"name:Action" json:"action"`
	Flags                          uint16                 `idl:"name:wFlags" json:"flags"`
	RemoteMachineAuthorizationList string                 `idl:"name:wszRemoteMachineAuthorizationList;string" json:"remote_machine_authorization_list"`
	RemoteUserAuthorizationList    string                 `idl:"name:wszRemoteUserAuthorizationList;string" json:"remote_user_authorization_list"`
	EmbeddedContext                string                 `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	PlatformValidityList           *OSPlatformList        `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	Status                         RuleStatus             `idl:"name:Status" json:"status"`
	Origin                         RuleOriginType         `idl:"name:Origin" json:"origin"`
	GPOName                        string                 `idl:"name:wszGPOName;string" json:"gpo_name"`
	MetadataReserved               uint32                 `idl:"name:MetaDataReserved" json:"metadata_reserved"`
}

func (o *Rule20) xxx_PreparePayload(ctx context.Context) error {
	if len(o.RuleID) > int(10001) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.Direction > Direction(2) {
		return fmt.Errorf("Direction is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.LocalApplication) > int(10001) {
		return fmt.Errorf("LocalApplication is out of range")
	}
	if len(o.LocalService) > int(10001) {
		return fmt.Errorf("LocalService is out of range")
	}
	if o.Action > RuleAction(4) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.RemoteMachineAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteMachineAuthorizationList is out of range")
	}
	if len(o.RemoteUserAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteUserAuthorizationList is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(6) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule20) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Rule20{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Direction)); err != nil {
		return err
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if o.IPProtocolData != nil {
		if err := o.IPProtocolData.MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	} else {
		if err := (&Rule20_IPProtocolData{}).MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	}
	if o.LocalAddresses != nil {
		if err := o.LocalAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemoteAddresses != nil {
		if err := o.RemoteAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if o.LocalApplication != "" {
		_ptr_wszLocalApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalApplication); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalApplication, _ptr_wszLocalApplication); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalService != "" {
		_ptr_wszLocalService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalService); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalService, _ptr_wszLocalService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.RemoteMachineAuthorizationList != "" {
		_ptr_wszRemoteMachineAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteMachineAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteUserAuthorizationList != "" {
		_ptr_wszRemoteUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	return nil
}
func (o *Rule20) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &Rule20{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**Rule20) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Direction)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	if o.IPProtocolData == nil {
		o.IPProtocolData = &Rule20_IPProtocolData{}
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if err := o.IPProtocolData.UnmarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
		return err
	}
	if o.LocalAddresses == nil {
		o.LocalAddresses = &Addresses{}
	}
	if err := o.LocalAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemoteAddresses == nil {
		o.RemoteAddresses = &Addresses{}
	}
	if err := o.RemoteAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	_ptr_wszLocalApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalApplication); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalApplication := func(ptr interface{}) { o.LocalApplication = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalApplication, _s_wszLocalApplication, _ptr_wszLocalApplication); err != nil {
		return err
	}
	_ptr_wszLocalService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalService); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalService := func(ptr interface{}) { o.LocalService = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalService, _s_wszLocalService, _ptr_wszLocalService); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszRemoteMachineAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteMachineAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteMachineAuthorizationList := func(ptr interface{}) { o.RemoteMachineAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteMachineAuthorizationList, _s_wszRemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
		return err
	}
	_ptr_wszRemoteUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteUserAuthorizationList := func(ptr interface{}) { o.RemoteUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteUserAuthorizationList, _s_wszRemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	return nil
}

// Rule20_IPProtocolData structure represents FW_RULE2_0 union anonymous member.
//
// This structure represents a firewall rule that is used by the 2.0 binary version
// servers and clients (see section 2.2.42). The fields of this structure are identical
// to the FW_RULE structure and its meanings are covered in section 2.2.37.
type Rule20_IPProtocolData struct {
	// Types that are assignable to Value
	//
	// *Rule20_IPProtocolData_Ports
	// *Rule20_IPProtocolData_TypeCodeListV4
	// *Rule20_IPProtocolData_TypeCodeListV6
	Value is_Rule20_IPProtocolData `json:"value"`
}

func (o *Rule20_IPProtocolData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Rule20_IPProtocolData_Ports:
		if value != nil {
			return value
		}
	case *Rule20_IPProtocolData_TypeCodeListV4:
		if value != nil {
			return value.TypeCodeListV4
		}
	case *Rule20_IPProtocolData_TypeCodeListV6:
		if value != nil {
			return value.TypeCodeListV6
		}
	}
	return nil
}

type is_Rule20_IPProtocolData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Rule20_IPProtocolData()
}

func (o *Rule20_IPProtocolData) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Rule20_IPProtocolData_Ports:
		switch sw {
		case uint16(6),
			uint16(17):
			return sw
		}
		return uint16(6)
	case *Rule20_IPProtocolData_TypeCodeListV4:
		return uint16(1)
	case *Rule20_IPProtocolData_TypeCodeListV6:
		return uint16(58)
	}
	return uint16(0)
}

func (o *Rule20_IPProtocolData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		_o, _ := o.Value.(*Rule20_IPProtocolData_Ports)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule20_IPProtocolData_Ports{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Rule20_IPProtocolData_TypeCodeListV4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule20_IPProtocolData_TypeCodeListV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(58):
		_o, _ := o.Value.(*Rule20_IPProtocolData_TypeCodeListV6)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule20_IPProtocolData_TypeCodeListV6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *Rule20_IPProtocolData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		o.Value = &Rule20_IPProtocolData_Ports{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Rule20_IPProtocolData_TypeCodeListV4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(58):
		o.Value = &Rule20_IPProtocolData_TypeCodeListV6{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// Rule20_IPProtocolData_Ports structure represents Rule20_IPProtocolData RPC union arm.
//
// It has following labels: 6, 17
type Rule20_IPProtocolData_Ports struct {
	LocalPorts  *Ports `idl:"name:LocalPorts" json:"local_ports"`
	RemotePorts *Ports `idl:"name:RemotePorts" json:"remote_ports"`
}

func (*Rule20_IPProtocolData_Ports) is_Rule20_IPProtocolData() {}

func (o *Rule20_IPProtocolData_Ports) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LocalPorts != nil {
		if err := o.LocalPorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemotePorts != nil {
		if err := o.RemotePorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule20_IPProtocolData_Ports) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LocalPorts == nil {
		o.LocalPorts = &Ports{}
	}
	if err := o.LocalPorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemotePorts == nil {
		o.RemotePorts = &Ports{}
	}
	if err := o.RemotePorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule20_IPProtocolData_TypeCodeListV4 structure represents Rule20_IPProtocolData RPC union arm.
//
// It has following labels: 1
type Rule20_IPProtocolData_TypeCodeListV4 struct {
	TypeCodeListV4 *ICMPTypeCodeList `idl:"name:V4TypeCodeList" json:"type_code_list_v4"`
}

func (*Rule20_IPProtocolData_TypeCodeListV4) is_Rule20_IPProtocolData() {}

func (o *Rule20_IPProtocolData_TypeCodeListV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV4 != nil {
		if err := o.TypeCodeListV4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule20_IPProtocolData_TypeCodeListV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV4 == nil {
		o.TypeCodeListV4 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule20_IPProtocolData_TypeCodeListV6 structure represents Rule20_IPProtocolData RPC union arm.
//
// It has following labels: 58
type Rule20_IPProtocolData_TypeCodeListV6 struct {
	TypeCodeListV6 *ICMPTypeCodeList `idl:"name:V6TypeCodeList" json:"type_code_list_v6"`
}

func (*Rule20_IPProtocolData_TypeCodeListV6) is_Rule20_IPProtocolData() {}

func (o *Rule20_IPProtocolData_TypeCodeListV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV6 != nil {
		if err := o.TypeCodeListV6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule20_IPProtocolData_TypeCodeListV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV6 == nil {
		o.TypeCodeListV6 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule210 structure represents FW_RULE2_10 RPC structure.
type Rule210 struct {
	Next                           *Rule210                `idl:"name:pNext" json:"next"`
	SchemaVersion                  uint16                  `idl:"name:wSchemaVersion" json:"schema_version"`
	RuleID                         string                  `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	Name                           string                  `idl:"name:wszName;string" json:"name"`
	Description                    string                  `idl:"name:wszDescription;string" json:"description"`
	Profiles                       uint32                  `idl:"name:dwProfiles" json:"profiles"`
	Direction                      Direction               `idl:"name:Direction" json:"direction"`
	IPProtocol                     uint16                  `idl:"name:wIpProtocol" json:"ip_protocol"`
	IPProtocolData                 *Rule210_IPProtocolData `idl:"name:IpProtocolData;switch_is:wIpProtocol" json:"ip_protocol_data"`
	LocalAddresses                 *Addresses              `idl:"name:LocalAddresses" json:"local_addresses"`
	RemoteAddresses                *Addresses              `idl:"name:RemoteAddresses" json:"remote_addresses"`
	LocalInterfaceIDs              *InterfaceLUIDs         `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	LocalInterfaceTypes            uint32                  `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	LocalApplication               string                  `idl:"name:wszLocalApplication;string" json:"local_application"`
	LocalService                   string                  `idl:"name:wszLocalService;string" json:"local_service"`
	Action                         RuleAction              `idl:"name:Action" json:"action"`
	Flags                          uint16                  `idl:"name:wFlags" json:"flags"`
	RemoteMachineAuthorizationList string                  `idl:"name:wszRemoteMachineAuthorizationList;string" json:"remote_machine_authorization_list"`
	RemoteUserAuthorizationList    string                  `idl:"name:wszRemoteUserAuthorizationList;string" json:"remote_user_authorization_list"`
	EmbeddedContext                string                  `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	PlatformValidityList           *OSPlatformList         `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	Status                         RuleStatus              `idl:"name:Status" json:"status"`
	Origin                         RuleOriginType          `idl:"name:Origin" json:"origin"`
	GPOName                        string                  `idl:"name:wszGPOName;string" json:"gpo_name"`
	MetadataReserved               uint32                  `idl:"name:MetaDataReserved" json:"metadata_reserved"`
	Metadata                       []*ObjectMetadata       `idl:"name:pMetaData;size_is:(((MetaDataReserved&1)?1:0))" json:"metadata"`
}

func (o *Rule210) xxx_PreparePayload(ctx context.Context) error {
	if o.Metadata != nil && o.MetadataReserved == 0 {
		o.MetadataReserved = uint32(len(o.Metadata))
	}
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.Direction > Direction(2) {
		return fmt.Errorf("Direction is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.LocalApplication) > int(10001) {
		return fmt.Errorf("LocalApplication is out of range")
	}
	if len(o.LocalService) > int(10001) {
		return fmt.Errorf("LocalService is out of range")
	}
	if o.Action > RuleAction(4) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.RemoteMachineAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteMachineAuthorizationList is out of range")
	}
	if len(o.RemoteUserAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteUserAuthorizationList is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(6) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule210) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Rule210{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Direction)); err != nil {
		return err
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if o.IPProtocolData != nil {
		if err := o.IPProtocolData.MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	} else {
		if err := (&Rule210_IPProtocolData{}).MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	}
	if o.LocalAddresses != nil {
		if err := o.LocalAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemoteAddresses != nil {
		if err := o.RemoteAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if o.LocalApplication != "" {
		_ptr_wszLocalApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalApplication); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalApplication, _ptr_wszLocalApplication); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalService != "" {
		_ptr_wszLocalService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalService); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalService, _ptr_wszLocalService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.RemoteMachineAuthorizationList != "" {
		_ptr_wszRemoteMachineAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteMachineAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteUserAuthorizationList != "" {
		_ptr_wszRemoteUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	_exprMetaDataReserved := uint32(0)
	if (o.MetadataReserved & 1) != 0 {
		_exprMetaDataReserved = uint32(1)
	} else {
		_exprMetaDataReserved = uint32(0)
	}
	if o.Metadata != nil || _exprMetaDataReserved > 0 {
		_ptr_pMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprMetaDataReserved := uint64(0)
			if (o.MetadataReserved & 1) != 0 {
				_exprMetaDataReserved = uint64(1)
			} else {
				_exprMetaDataReserved = uint64(0)
			}
			dimSize1 := uint64(_exprMetaDataReserved)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Metadata {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Metadata[i1] != nil {
					if err := o.Metadata[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Metadata); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Metadata, _ptr_pMetaData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule210) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &Rule210{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**Rule210) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Direction)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	if o.IPProtocolData == nil {
		o.IPProtocolData = &Rule210_IPProtocolData{}
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if err := o.IPProtocolData.UnmarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
		return err
	}
	if o.LocalAddresses == nil {
		o.LocalAddresses = &Addresses{}
	}
	if err := o.LocalAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemoteAddresses == nil {
		o.RemoteAddresses = &Addresses{}
	}
	if err := o.RemoteAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	_ptr_wszLocalApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalApplication); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalApplication := func(ptr interface{}) { o.LocalApplication = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalApplication, _s_wszLocalApplication, _ptr_wszLocalApplication); err != nil {
		return err
	}
	_ptr_wszLocalService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalService); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalService := func(ptr interface{}) { o.LocalService = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalService, _s_wszLocalService, _ptr_wszLocalService); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszRemoteMachineAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteMachineAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteMachineAuthorizationList := func(ptr interface{}) { o.RemoteMachineAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteMachineAuthorizationList, _s_wszRemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
		return err
	}
	_ptr_wszRemoteUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteUserAuthorizationList := func(ptr interface{}) { o.RemoteUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteUserAuthorizationList, _s_wszRemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	_ptr_pMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprpMetaData := uint64(0)
		if (o.MetadataReserved & 1) != 0 {
			_exprpMetaData = uint64(1)
		} else {
			_exprpMetaData = uint64(0)
		}
		if _exprpMetaData > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprpMetaData)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Metadata", sizeInfo[0])
		}
		o.Metadata = make([]*ObjectMetadata, sizeInfo[0])
		for i1 := range o.Metadata {
			i1 := i1
			if o.Metadata[i1] == nil {
				o.Metadata[i1] = &ObjectMetadata{}
			}
			if err := o.Metadata[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMetaData := func(ptr interface{}) { o.Metadata = *ptr.(*[]*ObjectMetadata) }
	if err := w.ReadPointer(&o.Metadata, _s_pMetaData, _ptr_pMetaData); err != nil {
		return err
	}
	return nil
}

// Rule210_IPProtocolData structure represents FW_RULE2_10 union anonymous member.
type Rule210_IPProtocolData struct {
	// Types that are assignable to Value
	//
	// *Rule210_IPProtocolData_Ports
	// *Rule210_IPProtocolData_TypeCodeListV4
	// *Rule210_IPProtocolData_TypeCodeListV6
	Value is_Rule210_IPProtocolData `json:"value"`
}

func (o *Rule210_IPProtocolData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Rule210_IPProtocolData_Ports:
		if value != nil {
			return value
		}
	case *Rule210_IPProtocolData_TypeCodeListV4:
		if value != nil {
			return value.TypeCodeListV4
		}
	case *Rule210_IPProtocolData_TypeCodeListV6:
		if value != nil {
			return value.TypeCodeListV6
		}
	}
	return nil
}

type is_Rule210_IPProtocolData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Rule210_IPProtocolData()
}

func (o *Rule210_IPProtocolData) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Rule210_IPProtocolData_Ports:
		switch sw {
		case uint16(6),
			uint16(17):
			return sw
		}
		return uint16(6)
	case *Rule210_IPProtocolData_TypeCodeListV4:
		return uint16(1)
	case *Rule210_IPProtocolData_TypeCodeListV6:
		return uint16(58)
	}
	return uint16(0)
}

func (o *Rule210_IPProtocolData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		_o, _ := o.Value.(*Rule210_IPProtocolData_Ports)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule210_IPProtocolData_Ports{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Rule210_IPProtocolData_TypeCodeListV4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule210_IPProtocolData_TypeCodeListV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(58):
		_o, _ := o.Value.(*Rule210_IPProtocolData_TypeCodeListV6)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule210_IPProtocolData_TypeCodeListV6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *Rule210_IPProtocolData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		o.Value = &Rule210_IPProtocolData_Ports{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Rule210_IPProtocolData_TypeCodeListV4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(58):
		o.Value = &Rule210_IPProtocolData_TypeCodeListV6{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// Rule210_IPProtocolData_Ports structure represents Rule210_IPProtocolData RPC union arm.
//
// It has following labels: 6, 17
type Rule210_IPProtocolData_Ports struct {
	LocalPorts  *Ports `idl:"name:LocalPorts" json:"local_ports"`
	RemotePorts *Ports `idl:"name:RemotePorts" json:"remote_ports"`
}

func (*Rule210_IPProtocolData_Ports) is_Rule210_IPProtocolData() {}

func (o *Rule210_IPProtocolData_Ports) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LocalPorts != nil {
		if err := o.LocalPorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemotePorts != nil {
		if err := o.RemotePorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule210_IPProtocolData_Ports) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LocalPorts == nil {
		o.LocalPorts = &Ports{}
	}
	if err := o.LocalPorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemotePorts == nil {
		o.RemotePorts = &Ports{}
	}
	if err := o.RemotePorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule210_IPProtocolData_TypeCodeListV4 structure represents Rule210_IPProtocolData RPC union arm.
//
// It has following labels: 1
type Rule210_IPProtocolData_TypeCodeListV4 struct {
	TypeCodeListV4 *ICMPTypeCodeList `idl:"name:V4TypeCodeList" json:"type_code_list_v4"`
}

func (*Rule210_IPProtocolData_TypeCodeListV4) is_Rule210_IPProtocolData() {}

func (o *Rule210_IPProtocolData_TypeCodeListV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV4 != nil {
		if err := o.TypeCodeListV4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule210_IPProtocolData_TypeCodeListV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV4 == nil {
		o.TypeCodeListV4 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule210_IPProtocolData_TypeCodeListV6 structure represents Rule210_IPProtocolData RPC union arm.
//
// It has following labels: 58
type Rule210_IPProtocolData_TypeCodeListV6 struct {
	TypeCodeListV6 *ICMPTypeCodeList `idl:"name:V6TypeCodeList" json:"type_code_list_v6"`
}

func (*Rule210_IPProtocolData_TypeCodeListV6) is_Rule210_IPProtocolData() {}

func (o *Rule210_IPProtocolData_TypeCodeListV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV6 != nil {
		if err := o.TypeCodeListV6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule210_IPProtocolData_TypeCodeListV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV6 == nil {
		o.TypeCodeListV6 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule220 structure represents FW_RULE2_20 RPC structure.
type Rule220 struct {
	Next                           *Rule220                `idl:"name:pNext" json:"next"`
	SchemaVersion                  uint16                  `idl:"name:wSchemaVersion" json:"schema_version"`
	RuleID                         string                  `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	Name                           string                  `idl:"name:wszName;string" json:"name"`
	Description                    string                  `idl:"name:wszDescription;string" json:"description"`
	Profiles                       uint32                  `idl:"name:dwProfiles" json:"profiles"`
	Direction                      Direction               `idl:"name:Direction" json:"direction"`
	IPProtocol                     uint16                  `idl:"name:wIpProtocol" json:"ip_protocol"`
	IPProtocolData                 *Rule220_IPProtocolData `idl:"name:IpProtocolData;switch_is:wIpProtocol" json:"ip_protocol_data"`
	LocalAddresses                 *Addresses              `idl:"name:LocalAddresses" json:"local_addresses"`
	RemoteAddresses                *Addresses              `idl:"name:RemoteAddresses" json:"remote_addresses"`
	LocalInterfaceIDs              *InterfaceLUIDs         `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	LocalInterfaceTypes            uint32                  `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	LocalApplication               string                  `idl:"name:wszLocalApplication;string" json:"local_application"`
	LocalService                   string                  `idl:"name:wszLocalService;string" json:"local_service"`
	Action                         RuleAction              `idl:"name:Action" json:"action"`
	Flags                          uint16                  `idl:"name:wFlags" json:"flags"`
	RemoteMachineAuthorizationList string                  `idl:"name:wszRemoteMachineAuthorizationList;string" json:"remote_machine_authorization_list"`
	RemoteUserAuthorizationList    string                  `idl:"name:wszRemoteUserAuthorizationList;string" json:"remote_user_authorization_list"`
	EmbeddedContext                string                  `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	PlatformValidityList           *OSPlatformList         `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	Status                         RuleStatus              `idl:"name:Status" json:"status"`
	Origin                         RuleOriginType          `idl:"name:Origin" json:"origin"`
	GPOName                        string                  `idl:"name:wszGPOName;string" json:"gpo_name"`
	MetadataReserved               uint32                  `idl:"name:MetaDataReserved" json:"metadata_reserved"`
	Metadata                       []*ObjectMetadata       `idl:"name:pMetaData;size_is:(((MetaDataReserved&1)?1:0))" json:"metadata"`
	LocalUserAuthorizationList     string                  `idl:"name:wszLocalUserAuthorizationList;string" json:"local_user_authorization_list"`
	PackageID                      string                  `idl:"name:wszPackageId;string" json:"package_id"`
	LocalUserOwner                 string                  `idl:"name:wszLocalUserOwner;string" json:"local_user_owner"`
	TrustTupleKeywords             uint32                  `idl:"name:dwTrustTupleKeywords" json:"trust_tuple_keywords"`
}

func (o *Rule220) xxx_PreparePayload(ctx context.Context) error {
	if o.Metadata != nil && o.MetadataReserved == 0 {
		o.MetadataReserved = uint32(len(o.Metadata))
	}
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.Direction > Direction(2) {
		return fmt.Errorf("Direction is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.LocalApplication) > int(10001) {
		return fmt.Errorf("LocalApplication is out of range")
	}
	if len(o.LocalService) > int(10001) {
		return fmt.Errorf("LocalService is out of range")
	}
	if o.Action > RuleAction(4) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.RemoteMachineAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteMachineAuthorizationList is out of range")
	}
	if len(o.RemoteUserAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteUserAuthorizationList is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(6) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if len(o.LocalUserAuthorizationList) > int(10001) {
		return fmt.Errorf("LocalUserAuthorizationList is out of range")
	}
	if len(o.PackageID) > int(10001) {
		return fmt.Errorf("PackageID is out of range")
	}
	if len(o.LocalUserOwner) > int(10001) {
		return fmt.Errorf("LocalUserOwner is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule220) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Rule220{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Direction)); err != nil {
		return err
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if o.IPProtocolData != nil {
		if err := o.IPProtocolData.MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	} else {
		if err := (&Rule220_IPProtocolData{}).MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	}
	if o.LocalAddresses != nil {
		if err := o.LocalAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemoteAddresses != nil {
		if err := o.RemoteAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if o.LocalApplication != "" {
		_ptr_wszLocalApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalApplication); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalApplication, _ptr_wszLocalApplication); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalService != "" {
		_ptr_wszLocalService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalService); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalService, _ptr_wszLocalService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.RemoteMachineAuthorizationList != "" {
		_ptr_wszRemoteMachineAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteMachineAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteUserAuthorizationList != "" {
		_ptr_wszRemoteUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	_exprMetaDataReserved := uint32(0)
	if (o.MetadataReserved & 1) != 0 {
		_exprMetaDataReserved = uint32(1)
	} else {
		_exprMetaDataReserved = uint32(0)
	}
	if o.Metadata != nil || _exprMetaDataReserved > 0 {
		_ptr_pMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprMetaDataReserved := uint64(0)
			if (o.MetadataReserved & 1) != 0 {
				_exprMetaDataReserved = uint64(1)
			} else {
				_exprMetaDataReserved = uint64(0)
			}
			dimSize1 := uint64(_exprMetaDataReserved)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Metadata {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Metadata[i1] != nil {
					if err := o.Metadata[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Metadata); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Metadata, _ptr_pMetaData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserAuthorizationList != "" {
		_ptr_wszLocalUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PackageID != "" {
		_ptr_wszPackageId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PackageID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PackageID, _ptr_wszPackageId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserOwner != "" {
		_ptr_wszLocalUserOwner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserOwner); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TrustTupleKeywords); err != nil {
		return err
	}
	return nil
}
func (o *Rule220) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &Rule220{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**Rule220) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Direction)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	if o.IPProtocolData == nil {
		o.IPProtocolData = &Rule220_IPProtocolData{}
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if err := o.IPProtocolData.UnmarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
		return err
	}
	if o.LocalAddresses == nil {
		o.LocalAddresses = &Addresses{}
	}
	if err := o.LocalAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemoteAddresses == nil {
		o.RemoteAddresses = &Addresses{}
	}
	if err := o.RemoteAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	_ptr_wszLocalApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalApplication); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalApplication := func(ptr interface{}) { o.LocalApplication = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalApplication, _s_wszLocalApplication, _ptr_wszLocalApplication); err != nil {
		return err
	}
	_ptr_wszLocalService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalService); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalService := func(ptr interface{}) { o.LocalService = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalService, _s_wszLocalService, _ptr_wszLocalService); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszRemoteMachineAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteMachineAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteMachineAuthorizationList := func(ptr interface{}) { o.RemoteMachineAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteMachineAuthorizationList, _s_wszRemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
		return err
	}
	_ptr_wszRemoteUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteUserAuthorizationList := func(ptr interface{}) { o.RemoteUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteUserAuthorizationList, _s_wszRemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	_ptr_pMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprpMetaData := uint64(0)
		if (o.MetadataReserved & 1) != 0 {
			_exprpMetaData = uint64(1)
		} else {
			_exprpMetaData = uint64(0)
		}
		if _exprpMetaData > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprpMetaData)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Metadata", sizeInfo[0])
		}
		o.Metadata = make([]*ObjectMetadata, sizeInfo[0])
		for i1 := range o.Metadata {
			i1 := i1
			if o.Metadata[i1] == nil {
				o.Metadata[i1] = &ObjectMetadata{}
			}
			if err := o.Metadata[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMetaData := func(ptr interface{}) { o.Metadata = *ptr.(*[]*ObjectMetadata) }
	if err := w.ReadPointer(&o.Metadata, _s_pMetaData, _ptr_pMetaData); err != nil {
		return err
	}
	_ptr_wszLocalUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserAuthorizationList := func(ptr interface{}) { o.LocalUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserAuthorizationList, _s_wszLocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszPackageId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PackageID); err != nil {
			return err
		}
		return nil
	})
	_s_wszPackageId := func(ptr interface{}) { o.PackageID = *ptr.(*string) }
	if err := w.ReadPointer(&o.PackageID, _s_wszPackageId, _ptr_wszPackageId); err != nil {
		return err
	}
	_ptr_wszLocalUserOwner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserOwner); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserOwner := func(ptr interface{}) { o.LocalUserOwner = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserOwner, _s_wszLocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustTupleKeywords); err != nil {
		return err
	}
	return nil
}

// Rule220_IPProtocolData structure represents FW_RULE2_20 union anonymous member.
type Rule220_IPProtocolData struct {
	// Types that are assignable to Value
	//
	// *Rule220_IPProtocolData_Ports
	// *Rule220_IPProtocolData_TypeCodeListV4
	// *Rule220_IPProtocolData_TypeCodeListV6
	Value is_Rule220_IPProtocolData `json:"value"`
}

func (o *Rule220_IPProtocolData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Rule220_IPProtocolData_Ports:
		if value != nil {
			return value
		}
	case *Rule220_IPProtocolData_TypeCodeListV4:
		if value != nil {
			return value.TypeCodeListV4
		}
	case *Rule220_IPProtocolData_TypeCodeListV6:
		if value != nil {
			return value.TypeCodeListV6
		}
	}
	return nil
}

type is_Rule220_IPProtocolData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Rule220_IPProtocolData()
}

func (o *Rule220_IPProtocolData) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Rule220_IPProtocolData_Ports:
		switch sw {
		case uint16(6),
			uint16(17):
			return sw
		}
		return uint16(6)
	case *Rule220_IPProtocolData_TypeCodeListV4:
		return uint16(1)
	case *Rule220_IPProtocolData_TypeCodeListV6:
		return uint16(58)
	}
	return uint16(0)
}

func (o *Rule220_IPProtocolData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		_o, _ := o.Value.(*Rule220_IPProtocolData_Ports)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule220_IPProtocolData_Ports{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Rule220_IPProtocolData_TypeCodeListV4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule220_IPProtocolData_TypeCodeListV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(58):
		_o, _ := o.Value.(*Rule220_IPProtocolData_TypeCodeListV6)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule220_IPProtocolData_TypeCodeListV6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *Rule220_IPProtocolData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		o.Value = &Rule220_IPProtocolData_Ports{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Rule220_IPProtocolData_TypeCodeListV4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(58):
		o.Value = &Rule220_IPProtocolData_TypeCodeListV6{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// Rule220_IPProtocolData_Ports structure represents Rule220_IPProtocolData RPC union arm.
//
// It has following labels: 6, 17
type Rule220_IPProtocolData_Ports struct {
	LocalPorts  *Ports `idl:"name:LocalPorts" json:"local_ports"`
	RemotePorts *Ports `idl:"name:RemotePorts" json:"remote_ports"`
}

func (*Rule220_IPProtocolData_Ports) is_Rule220_IPProtocolData() {}

func (o *Rule220_IPProtocolData_Ports) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LocalPorts != nil {
		if err := o.LocalPorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemotePorts != nil {
		if err := o.RemotePorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule220_IPProtocolData_Ports) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LocalPorts == nil {
		o.LocalPorts = &Ports{}
	}
	if err := o.LocalPorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemotePorts == nil {
		o.RemotePorts = &Ports{}
	}
	if err := o.RemotePorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule220_IPProtocolData_TypeCodeListV4 structure represents Rule220_IPProtocolData RPC union arm.
//
// It has following labels: 1
type Rule220_IPProtocolData_TypeCodeListV4 struct {
	TypeCodeListV4 *ICMPTypeCodeList `idl:"name:V4TypeCodeList" json:"type_code_list_v4"`
}

func (*Rule220_IPProtocolData_TypeCodeListV4) is_Rule220_IPProtocolData() {}

func (o *Rule220_IPProtocolData_TypeCodeListV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV4 != nil {
		if err := o.TypeCodeListV4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule220_IPProtocolData_TypeCodeListV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV4 == nil {
		o.TypeCodeListV4 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule220_IPProtocolData_TypeCodeListV6 structure represents Rule220_IPProtocolData RPC union arm.
//
// It has following labels: 58
type Rule220_IPProtocolData_TypeCodeListV6 struct {
	TypeCodeListV6 *ICMPTypeCodeList `idl:"name:V6TypeCodeList" json:"type_code_list_v6"`
}

func (*Rule220_IPProtocolData_TypeCodeListV6) is_Rule220_IPProtocolData() {}

func (o *Rule220_IPProtocolData_TypeCodeListV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV6 != nil {
		if err := o.TypeCodeListV6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule220_IPProtocolData_TypeCodeListV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV6 == nil {
		o.TypeCodeListV6 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule224 structure represents FW_RULE2_24 RPC structure.
type Rule224 struct {
	Next                           *Rule224                `idl:"name:pNext" json:"next"`
	SchemaVersion                  uint16                  `idl:"name:wSchemaVersion" json:"schema_version"`
	RuleID                         string                  `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	Name                           string                  `idl:"name:wszName;string" json:"name"`
	Description                    string                  `idl:"name:wszDescription;string" json:"description"`
	Profiles                       uint32                  `idl:"name:dwProfiles" json:"profiles"`
	Direction                      Direction               `idl:"name:Direction" json:"direction"`
	IPProtocol                     uint16                  `idl:"name:wIpProtocol" json:"ip_protocol"`
	IPProtocolData                 *Rule224_IPProtocolData `idl:"name:IpProtocolData;switch_is:wIpProtocol" json:"ip_protocol_data"`
	LocalAddresses                 *Addresses              `idl:"name:LocalAddresses" json:"local_addresses"`
	RemoteAddresses                *Addresses              `idl:"name:RemoteAddresses" json:"remote_addresses"`
	LocalInterfaceIDs              *InterfaceLUIDs         `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	LocalInterfaceTypes            uint32                  `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	LocalApplication               string                  `idl:"name:wszLocalApplication;string" json:"local_application"`
	LocalService                   string                  `idl:"name:wszLocalService;string" json:"local_service"`
	Action                         RuleAction              `idl:"name:Action" json:"action"`
	Flags                          uint16                  `idl:"name:wFlags" json:"flags"`
	RemoteMachineAuthorizationList string                  `idl:"name:wszRemoteMachineAuthorizationList;string" json:"remote_machine_authorization_list"`
	RemoteUserAuthorizationList    string                  `idl:"name:wszRemoteUserAuthorizationList;string" json:"remote_user_authorization_list"`
	EmbeddedContext                string                  `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	PlatformValidityList           *OSPlatformList         `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	Status                         RuleStatus              `idl:"name:Status" json:"status"`
	Origin                         RuleOriginType          `idl:"name:Origin" json:"origin"`
	GPOName                        string                  `idl:"name:wszGPOName;string" json:"gpo_name"`
	MetadataReserved               uint32                  `idl:"name:MetaDataReserved" json:"metadata_reserved"`
	Metadata                       []*ObjectMetadata       `idl:"name:pMetaData;size_is:(((MetaDataReserved&1)?1:0))" json:"metadata"`
	LocalUserAuthorizationList     string                  `idl:"name:wszLocalUserAuthorizationList;string" json:"local_user_authorization_list"`
	PackageID                      string                  `idl:"name:wszPackageId;string" json:"package_id"`
	LocalUserOwner                 string                  `idl:"name:wszLocalUserOwner;string" json:"local_user_owner"`
	TrustTupleKeywords             uint32                  `idl:"name:dwTrustTupleKeywords" json:"trust_tuple_keywords"`
	OnNetworkNames                 *NetworkNames           `idl:"name:OnNetworkNames" json:"on_network_names"`
	SecurityRealmID                string                  `idl:"name:wszSecurityRealmId;string" json:"security_realm_id"`
}

func (o *Rule224) xxx_PreparePayload(ctx context.Context) error {
	if o.Metadata != nil && o.MetadataReserved == 0 {
		o.MetadataReserved = uint32(len(o.Metadata))
	}
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.Direction > Direction(2) {
		return fmt.Errorf("Direction is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.LocalApplication) > int(10001) {
		return fmt.Errorf("LocalApplication is out of range")
	}
	if len(o.LocalService) > int(10001) {
		return fmt.Errorf("LocalService is out of range")
	}
	if o.Action > RuleAction(4) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.RemoteMachineAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteMachineAuthorizationList is out of range")
	}
	if len(o.RemoteUserAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteUserAuthorizationList is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(6) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if len(o.LocalUserAuthorizationList) > int(10001) {
		return fmt.Errorf("LocalUserAuthorizationList is out of range")
	}
	if len(o.PackageID) > int(10001) {
		return fmt.Errorf("PackageID is out of range")
	}
	if len(o.LocalUserOwner) > int(10001) {
		return fmt.Errorf("LocalUserOwner is out of range")
	}
	if len(o.SecurityRealmID) > int(10001) {
		return fmt.Errorf("SecurityRealmID is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule224) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Rule224{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Direction)); err != nil {
		return err
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if o.IPProtocolData != nil {
		if err := o.IPProtocolData.MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	} else {
		if err := (&Rule224_IPProtocolData{}).MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	}
	if o.LocalAddresses != nil {
		if err := o.LocalAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemoteAddresses != nil {
		if err := o.RemoteAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if o.LocalApplication != "" {
		_ptr_wszLocalApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalApplication); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalApplication, _ptr_wszLocalApplication); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalService != "" {
		_ptr_wszLocalService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalService); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalService, _ptr_wszLocalService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.RemoteMachineAuthorizationList != "" {
		_ptr_wszRemoteMachineAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteMachineAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteUserAuthorizationList != "" {
		_ptr_wszRemoteUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	_exprMetaDataReserved := uint32(0)
	if (o.MetadataReserved & 1) != 0 {
		_exprMetaDataReserved = uint32(1)
	} else {
		_exprMetaDataReserved = uint32(0)
	}
	if o.Metadata != nil || _exprMetaDataReserved > 0 {
		_ptr_pMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprMetaDataReserved := uint64(0)
			if (o.MetadataReserved & 1) != 0 {
				_exprMetaDataReserved = uint64(1)
			} else {
				_exprMetaDataReserved = uint64(0)
			}
			dimSize1 := uint64(_exprMetaDataReserved)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Metadata {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Metadata[i1] != nil {
					if err := o.Metadata[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Metadata); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Metadata, _ptr_pMetaData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserAuthorizationList != "" {
		_ptr_wszLocalUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PackageID != "" {
		_ptr_wszPackageId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PackageID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PackageID, _ptr_wszPackageId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserOwner != "" {
		_ptr_wszLocalUserOwner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserOwner); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TrustTupleKeywords); err != nil {
		return err
	}
	if o.OnNetworkNames != nil {
		if err := o.OnNetworkNames.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NetworkNames{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SecurityRealmID != "" {
		_ptr_wszSecurityRealmId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SecurityRealmID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityRealmID, _ptr_wszSecurityRealmId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule224) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &Rule224{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**Rule224) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Direction)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	if o.IPProtocolData == nil {
		o.IPProtocolData = &Rule224_IPProtocolData{}
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if err := o.IPProtocolData.UnmarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
		return err
	}
	if o.LocalAddresses == nil {
		o.LocalAddresses = &Addresses{}
	}
	if err := o.LocalAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemoteAddresses == nil {
		o.RemoteAddresses = &Addresses{}
	}
	if err := o.RemoteAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	_ptr_wszLocalApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalApplication); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalApplication := func(ptr interface{}) { o.LocalApplication = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalApplication, _s_wszLocalApplication, _ptr_wszLocalApplication); err != nil {
		return err
	}
	_ptr_wszLocalService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalService); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalService := func(ptr interface{}) { o.LocalService = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalService, _s_wszLocalService, _ptr_wszLocalService); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszRemoteMachineAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteMachineAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteMachineAuthorizationList := func(ptr interface{}) { o.RemoteMachineAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteMachineAuthorizationList, _s_wszRemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
		return err
	}
	_ptr_wszRemoteUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteUserAuthorizationList := func(ptr interface{}) { o.RemoteUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteUserAuthorizationList, _s_wszRemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	_ptr_pMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprpMetaData := uint64(0)
		if (o.MetadataReserved & 1) != 0 {
			_exprpMetaData = uint64(1)
		} else {
			_exprpMetaData = uint64(0)
		}
		if _exprpMetaData > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprpMetaData)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Metadata", sizeInfo[0])
		}
		o.Metadata = make([]*ObjectMetadata, sizeInfo[0])
		for i1 := range o.Metadata {
			i1 := i1
			if o.Metadata[i1] == nil {
				o.Metadata[i1] = &ObjectMetadata{}
			}
			if err := o.Metadata[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMetaData := func(ptr interface{}) { o.Metadata = *ptr.(*[]*ObjectMetadata) }
	if err := w.ReadPointer(&o.Metadata, _s_pMetaData, _ptr_pMetaData); err != nil {
		return err
	}
	_ptr_wszLocalUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserAuthorizationList := func(ptr interface{}) { o.LocalUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserAuthorizationList, _s_wszLocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszPackageId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PackageID); err != nil {
			return err
		}
		return nil
	})
	_s_wszPackageId := func(ptr interface{}) { o.PackageID = *ptr.(*string) }
	if err := w.ReadPointer(&o.PackageID, _s_wszPackageId, _ptr_wszPackageId); err != nil {
		return err
	}
	_ptr_wszLocalUserOwner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserOwner); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserOwner := func(ptr interface{}) { o.LocalUserOwner = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserOwner, _s_wszLocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustTupleKeywords); err != nil {
		return err
	}
	if o.OnNetworkNames == nil {
		o.OnNetworkNames = &NetworkNames{}
	}
	if err := o.OnNetworkNames.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_wszSecurityRealmId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SecurityRealmID); err != nil {
			return err
		}
		return nil
	})
	_s_wszSecurityRealmId := func(ptr interface{}) { o.SecurityRealmID = *ptr.(*string) }
	if err := w.ReadPointer(&o.SecurityRealmID, _s_wszSecurityRealmId, _ptr_wszSecurityRealmId); err != nil {
		return err
	}
	return nil
}

// Rule224_IPProtocolData structure represents FW_RULE2_24 union anonymous member.
type Rule224_IPProtocolData struct {
	// Types that are assignable to Value
	//
	// *Rule224_IPProtocolData_Ports
	// *Rule224_IPProtocolData_TypeCodeListV4
	// *Rule224_IPProtocolData_TypeCodeListV6
	Value is_Rule224_IPProtocolData `json:"value"`
}

func (o *Rule224_IPProtocolData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Rule224_IPProtocolData_Ports:
		if value != nil {
			return value
		}
	case *Rule224_IPProtocolData_TypeCodeListV4:
		if value != nil {
			return value.TypeCodeListV4
		}
	case *Rule224_IPProtocolData_TypeCodeListV6:
		if value != nil {
			return value.TypeCodeListV6
		}
	}
	return nil
}

type is_Rule224_IPProtocolData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Rule224_IPProtocolData()
}

func (o *Rule224_IPProtocolData) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Rule224_IPProtocolData_Ports:
		switch sw {
		case uint16(6),
			uint16(17):
			return sw
		}
		return uint16(6)
	case *Rule224_IPProtocolData_TypeCodeListV4:
		return uint16(1)
	case *Rule224_IPProtocolData_TypeCodeListV6:
		return uint16(58)
	}
	return uint16(0)
}

func (o *Rule224_IPProtocolData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		_o, _ := o.Value.(*Rule224_IPProtocolData_Ports)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule224_IPProtocolData_Ports{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Rule224_IPProtocolData_TypeCodeListV4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule224_IPProtocolData_TypeCodeListV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(58):
		_o, _ := o.Value.(*Rule224_IPProtocolData_TypeCodeListV6)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule224_IPProtocolData_TypeCodeListV6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *Rule224_IPProtocolData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		o.Value = &Rule224_IPProtocolData_Ports{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Rule224_IPProtocolData_TypeCodeListV4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(58):
		o.Value = &Rule224_IPProtocolData_TypeCodeListV6{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// Rule224_IPProtocolData_Ports structure represents Rule224_IPProtocolData RPC union arm.
//
// It has following labels: 6, 17
type Rule224_IPProtocolData_Ports struct {
	LocalPorts  *Ports `idl:"name:LocalPorts" json:"local_ports"`
	RemotePorts *Ports `idl:"name:RemotePorts" json:"remote_ports"`
}

func (*Rule224_IPProtocolData_Ports) is_Rule224_IPProtocolData() {}

func (o *Rule224_IPProtocolData_Ports) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LocalPorts != nil {
		if err := o.LocalPorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemotePorts != nil {
		if err := o.RemotePorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule224_IPProtocolData_Ports) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LocalPorts == nil {
		o.LocalPorts = &Ports{}
	}
	if err := o.LocalPorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemotePorts == nil {
		o.RemotePorts = &Ports{}
	}
	if err := o.RemotePorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule224_IPProtocolData_TypeCodeListV4 structure represents Rule224_IPProtocolData RPC union arm.
//
// It has following labels: 1
type Rule224_IPProtocolData_TypeCodeListV4 struct {
	TypeCodeListV4 *ICMPTypeCodeList `idl:"name:V4TypeCodeList" json:"type_code_list_v4"`
}

func (*Rule224_IPProtocolData_TypeCodeListV4) is_Rule224_IPProtocolData() {}

func (o *Rule224_IPProtocolData_TypeCodeListV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV4 != nil {
		if err := o.TypeCodeListV4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule224_IPProtocolData_TypeCodeListV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV4 == nil {
		o.TypeCodeListV4 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule224_IPProtocolData_TypeCodeListV6 structure represents Rule224_IPProtocolData RPC union arm.
//
// It has following labels: 58
type Rule224_IPProtocolData_TypeCodeListV6 struct {
	TypeCodeListV6 *ICMPTypeCodeList `idl:"name:V6TypeCodeList" json:"type_code_list_v6"`
}

func (*Rule224_IPProtocolData_TypeCodeListV6) is_Rule224_IPProtocolData() {}

func (o *Rule224_IPProtocolData_TypeCodeListV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV6 != nil {
		if err := o.TypeCodeListV6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule224_IPProtocolData_TypeCodeListV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV6 == nil {
		o.TypeCodeListV6 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule225 structure represents FW_RULE2_25 RPC structure.
type Rule225 struct {
	Next                           *Rule225                `idl:"name:pNext" json:"next"`
	SchemaVersion                  uint16                  `idl:"name:wSchemaVersion" json:"schema_version"`
	RuleID                         string                  `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	Name                           string                  `idl:"name:wszName;string" json:"name"`
	Description                    string                  `idl:"name:wszDescription;string" json:"description"`
	Profiles                       uint32                  `idl:"name:dwProfiles" json:"profiles"`
	Direction                      Direction               `idl:"name:Direction" json:"direction"`
	IPProtocol                     uint16                  `idl:"name:wIpProtocol" json:"ip_protocol"`
	IPProtocolData                 *Rule225_IPProtocolData `idl:"name:IpProtocolData;switch_is:wIpProtocol" json:"ip_protocol_data"`
	LocalAddresses                 *Addresses              `idl:"name:LocalAddresses" json:"local_addresses"`
	RemoteAddresses                *Addresses              `idl:"name:RemoteAddresses" json:"remote_addresses"`
	LocalInterfaceIDs              *InterfaceLUIDs         `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	LocalInterfaceTypes            uint32                  `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	LocalApplication               string                  `idl:"name:wszLocalApplication;string" json:"local_application"`
	LocalService                   string                  `idl:"name:wszLocalService;string" json:"local_service"`
	Action                         RuleAction              `idl:"name:Action" json:"action"`
	Flags                          uint16                  `idl:"name:wFlags" json:"flags"`
	RemoteMachineAuthorizationList string                  `idl:"name:wszRemoteMachineAuthorizationList;string" json:"remote_machine_authorization_list"`
	RemoteUserAuthorizationList    string                  `idl:"name:wszRemoteUserAuthorizationList;string" json:"remote_user_authorization_list"`
	EmbeddedContext                string                  `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	PlatformValidityList           *OSPlatformList         `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	Status                         RuleStatus              `idl:"name:Status" json:"status"`
	Origin                         RuleOriginType          `idl:"name:Origin" json:"origin"`
	GPOName                        string                  `idl:"name:wszGPOName;string" json:"gpo_name"`
	MetadataReserved               uint32                  `idl:"name:MetaDataReserved" json:"metadata_reserved"`
	Metadata                       []*ObjectMetadata       `idl:"name:pMetaData;size_is:(((MetaDataReserved&1)?1:0))" json:"metadata"`
	LocalUserAuthorizationList     string                  `idl:"name:wszLocalUserAuthorizationList;string" json:"local_user_authorization_list"`
	PackageID                      string                  `idl:"name:wszPackageId;string" json:"package_id"`
	LocalUserOwner                 string                  `idl:"name:wszLocalUserOwner;string" json:"local_user_owner"`
	TrustTupleKeywords             uint32                  `idl:"name:dwTrustTupleKeywords" json:"trust_tuple_keywords"`
	OnNetworkNames                 *NetworkNames           `idl:"name:OnNetworkNames" json:"on_network_names"`
	SecurityRealmID                string                  `idl:"name:wszSecurityRealmId;string" json:"security_realm_id"`
	Flags2                         uint16                  `idl:"name:wFlags2" json:"flags2"`
}

func (o *Rule225) xxx_PreparePayload(ctx context.Context) error {
	if o.Metadata != nil && o.MetadataReserved == 0 {
		o.MetadataReserved = uint32(len(o.Metadata))
	}
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.Direction > Direction(2) {
		return fmt.Errorf("Direction is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.LocalApplication) > int(10001) {
		return fmt.Errorf("LocalApplication is out of range")
	}
	if len(o.LocalService) > int(10001) {
		return fmt.Errorf("LocalService is out of range")
	}
	if o.Action > RuleAction(4) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.RemoteMachineAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteMachineAuthorizationList is out of range")
	}
	if len(o.RemoteUserAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteUserAuthorizationList is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(6) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if len(o.LocalUserAuthorizationList) > int(10001) {
		return fmt.Errorf("LocalUserAuthorizationList is out of range")
	}
	if len(o.PackageID) > int(10001) {
		return fmt.Errorf("PackageID is out of range")
	}
	if len(o.LocalUserOwner) > int(10001) {
		return fmt.Errorf("LocalUserOwner is out of range")
	}
	if len(o.SecurityRealmID) > int(10001) {
		return fmt.Errorf("SecurityRealmID is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule225) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Rule225{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Direction)); err != nil {
		return err
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if o.IPProtocolData != nil {
		if err := o.IPProtocolData.MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	} else {
		if err := (&Rule225_IPProtocolData{}).MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	}
	if o.LocalAddresses != nil {
		if err := o.LocalAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemoteAddresses != nil {
		if err := o.RemoteAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if o.LocalApplication != "" {
		_ptr_wszLocalApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalApplication); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalApplication, _ptr_wszLocalApplication); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalService != "" {
		_ptr_wszLocalService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalService); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalService, _ptr_wszLocalService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.RemoteMachineAuthorizationList != "" {
		_ptr_wszRemoteMachineAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteMachineAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteUserAuthorizationList != "" {
		_ptr_wszRemoteUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	_exprMetaDataReserved := uint32(0)
	if (o.MetadataReserved & 1) != 0 {
		_exprMetaDataReserved = uint32(1)
	} else {
		_exprMetaDataReserved = uint32(0)
	}
	if o.Metadata != nil || _exprMetaDataReserved > 0 {
		_ptr_pMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprMetaDataReserved := uint64(0)
			if (o.MetadataReserved & 1) != 0 {
				_exprMetaDataReserved = uint64(1)
			} else {
				_exprMetaDataReserved = uint64(0)
			}
			dimSize1 := uint64(_exprMetaDataReserved)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Metadata {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Metadata[i1] != nil {
					if err := o.Metadata[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Metadata); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Metadata, _ptr_pMetaData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserAuthorizationList != "" {
		_ptr_wszLocalUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PackageID != "" {
		_ptr_wszPackageId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PackageID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PackageID, _ptr_wszPackageId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserOwner != "" {
		_ptr_wszLocalUserOwner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserOwner); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TrustTupleKeywords); err != nil {
		return err
	}
	if o.OnNetworkNames != nil {
		if err := o.OnNetworkNames.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NetworkNames{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SecurityRealmID != "" {
		_ptr_wszSecurityRealmId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SecurityRealmID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityRealmID, _ptr_wszSecurityRealmId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags2); err != nil {
		return err
	}
	return nil
}
func (o *Rule225) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &Rule225{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**Rule225) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Direction)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	if o.IPProtocolData == nil {
		o.IPProtocolData = &Rule225_IPProtocolData{}
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if err := o.IPProtocolData.UnmarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
		return err
	}
	if o.LocalAddresses == nil {
		o.LocalAddresses = &Addresses{}
	}
	if err := o.LocalAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemoteAddresses == nil {
		o.RemoteAddresses = &Addresses{}
	}
	if err := o.RemoteAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	_ptr_wszLocalApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalApplication); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalApplication := func(ptr interface{}) { o.LocalApplication = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalApplication, _s_wszLocalApplication, _ptr_wszLocalApplication); err != nil {
		return err
	}
	_ptr_wszLocalService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalService); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalService := func(ptr interface{}) { o.LocalService = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalService, _s_wszLocalService, _ptr_wszLocalService); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszRemoteMachineAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteMachineAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteMachineAuthorizationList := func(ptr interface{}) { o.RemoteMachineAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteMachineAuthorizationList, _s_wszRemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
		return err
	}
	_ptr_wszRemoteUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteUserAuthorizationList := func(ptr interface{}) { o.RemoteUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteUserAuthorizationList, _s_wszRemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	_ptr_pMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprpMetaData := uint64(0)
		if (o.MetadataReserved & 1) != 0 {
			_exprpMetaData = uint64(1)
		} else {
			_exprpMetaData = uint64(0)
		}
		if _exprpMetaData > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprpMetaData)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Metadata", sizeInfo[0])
		}
		o.Metadata = make([]*ObjectMetadata, sizeInfo[0])
		for i1 := range o.Metadata {
			i1 := i1
			if o.Metadata[i1] == nil {
				o.Metadata[i1] = &ObjectMetadata{}
			}
			if err := o.Metadata[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMetaData := func(ptr interface{}) { o.Metadata = *ptr.(*[]*ObjectMetadata) }
	if err := w.ReadPointer(&o.Metadata, _s_pMetaData, _ptr_pMetaData); err != nil {
		return err
	}
	_ptr_wszLocalUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserAuthorizationList := func(ptr interface{}) { o.LocalUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserAuthorizationList, _s_wszLocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszPackageId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PackageID); err != nil {
			return err
		}
		return nil
	})
	_s_wszPackageId := func(ptr interface{}) { o.PackageID = *ptr.(*string) }
	if err := w.ReadPointer(&o.PackageID, _s_wszPackageId, _ptr_wszPackageId); err != nil {
		return err
	}
	_ptr_wszLocalUserOwner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserOwner); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserOwner := func(ptr interface{}) { o.LocalUserOwner = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserOwner, _s_wszLocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustTupleKeywords); err != nil {
		return err
	}
	if o.OnNetworkNames == nil {
		o.OnNetworkNames = &NetworkNames{}
	}
	if err := o.OnNetworkNames.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_wszSecurityRealmId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SecurityRealmID); err != nil {
			return err
		}
		return nil
	})
	_s_wszSecurityRealmId := func(ptr interface{}) { o.SecurityRealmID = *ptr.(*string) }
	if err := w.ReadPointer(&o.SecurityRealmID, _s_wszSecurityRealmId, _ptr_wszSecurityRealmId); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags2); err != nil {
		return err
	}
	return nil
}

// Rule225_IPProtocolData structure represents FW_RULE2_25 union anonymous member.
type Rule225_IPProtocolData struct {
	// Types that are assignable to Value
	//
	// *Rule225_IPProtocolData_Ports
	// *Rule225_IPProtocolData_TypeCodeListV4
	// *Rule225_IPProtocolData_TypeCodeListV6
	Value is_Rule225_IPProtocolData `json:"value"`
}

func (o *Rule225_IPProtocolData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Rule225_IPProtocolData_Ports:
		if value != nil {
			return value
		}
	case *Rule225_IPProtocolData_TypeCodeListV4:
		if value != nil {
			return value.TypeCodeListV4
		}
	case *Rule225_IPProtocolData_TypeCodeListV6:
		if value != nil {
			return value.TypeCodeListV6
		}
	}
	return nil
}

type is_Rule225_IPProtocolData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Rule225_IPProtocolData()
}

func (o *Rule225_IPProtocolData) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Rule225_IPProtocolData_Ports:
		switch sw {
		case uint16(6),
			uint16(17):
			return sw
		}
		return uint16(6)
	case *Rule225_IPProtocolData_TypeCodeListV4:
		return uint16(1)
	case *Rule225_IPProtocolData_TypeCodeListV6:
		return uint16(58)
	}
	return uint16(0)
}

func (o *Rule225_IPProtocolData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		_o, _ := o.Value.(*Rule225_IPProtocolData_Ports)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule225_IPProtocolData_Ports{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Rule225_IPProtocolData_TypeCodeListV4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule225_IPProtocolData_TypeCodeListV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(58):
		_o, _ := o.Value.(*Rule225_IPProtocolData_TypeCodeListV6)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule225_IPProtocolData_TypeCodeListV6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *Rule225_IPProtocolData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		o.Value = &Rule225_IPProtocolData_Ports{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Rule225_IPProtocolData_TypeCodeListV4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(58):
		o.Value = &Rule225_IPProtocolData_TypeCodeListV6{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// Rule225_IPProtocolData_Ports structure represents Rule225_IPProtocolData RPC union arm.
//
// It has following labels: 6, 17
type Rule225_IPProtocolData_Ports struct {
	LocalPorts  *Ports `idl:"name:LocalPorts" json:"local_ports"`
	RemotePorts *Ports `idl:"name:RemotePorts" json:"remote_ports"`
}

func (*Rule225_IPProtocolData_Ports) is_Rule225_IPProtocolData() {}

func (o *Rule225_IPProtocolData_Ports) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LocalPorts != nil {
		if err := o.LocalPorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemotePorts != nil {
		if err := o.RemotePorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule225_IPProtocolData_Ports) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LocalPorts == nil {
		o.LocalPorts = &Ports{}
	}
	if err := o.LocalPorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemotePorts == nil {
		o.RemotePorts = &Ports{}
	}
	if err := o.RemotePorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule225_IPProtocolData_TypeCodeListV4 structure represents Rule225_IPProtocolData RPC union arm.
//
// It has following labels: 1
type Rule225_IPProtocolData_TypeCodeListV4 struct {
	TypeCodeListV4 *ICMPTypeCodeList `idl:"name:V4TypeCodeList" json:"type_code_list_v4"`
}

func (*Rule225_IPProtocolData_TypeCodeListV4) is_Rule225_IPProtocolData() {}

func (o *Rule225_IPProtocolData_TypeCodeListV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV4 != nil {
		if err := o.TypeCodeListV4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule225_IPProtocolData_TypeCodeListV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV4 == nil {
		o.TypeCodeListV4 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule225_IPProtocolData_TypeCodeListV6 structure represents Rule225_IPProtocolData RPC union arm.
//
// It has following labels: 58
type Rule225_IPProtocolData_TypeCodeListV6 struct {
	TypeCodeListV6 *ICMPTypeCodeList `idl:"name:V6TypeCodeList" json:"type_code_list_v6"`
}

func (*Rule225_IPProtocolData_TypeCodeListV6) is_Rule225_IPProtocolData() {}

func (o *Rule225_IPProtocolData_TypeCodeListV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV6 != nil {
		if err := o.TypeCodeListV6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule225_IPProtocolData_TypeCodeListV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV6 == nil {
		o.TypeCodeListV6 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule226 structure represents FW_RULE2_26 RPC structure.
type Rule226 struct {
	Next                           *Rule226                `idl:"name:pNext" json:"next"`
	SchemaVersion                  uint16                  `idl:"name:wSchemaVersion" json:"schema_version"`
	RuleID                         string                  `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	Name                           string                  `idl:"name:wszName;string" json:"name"`
	Description                    string                  `idl:"name:wszDescription;string" json:"description"`
	Profiles                       uint32                  `idl:"name:dwProfiles" json:"profiles"`
	Direction                      Direction               `idl:"name:Direction" json:"direction"`
	IPProtocol                     uint16                  `idl:"name:wIpProtocol" json:"ip_protocol"`
	IPProtocolData                 *Rule226_IPProtocolData `idl:"name:IpProtocolData;switch_is:wIpProtocol" json:"ip_protocol_data"`
	LocalAddresses                 *Addresses              `idl:"name:LocalAddresses" json:"local_addresses"`
	RemoteAddresses                *Addresses              `idl:"name:RemoteAddresses" json:"remote_addresses"`
	LocalInterfaceIDs              *InterfaceLUIDs         `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	LocalInterfaceTypes            uint32                  `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	LocalApplication               string                  `idl:"name:wszLocalApplication;string" json:"local_application"`
	LocalService                   string                  `idl:"name:wszLocalService;string" json:"local_service"`
	Action                         RuleAction              `idl:"name:Action" json:"action"`
	Flags                          uint16                  `idl:"name:wFlags" json:"flags"`
	RemoteMachineAuthorizationList string                  `idl:"name:wszRemoteMachineAuthorizationList;string" json:"remote_machine_authorization_list"`
	RemoteUserAuthorizationList    string                  `idl:"name:wszRemoteUserAuthorizationList" json:"remote_user_authorization_list"`
	EmbeddedContext                string                  `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	PlatformValidityList           *OSPlatformList         `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	Status                         RuleStatus              `idl:"name:Status" json:"status"`
	Origin                         RuleOriginType          `idl:"name:Origin" json:"origin"`
	GPOName                        string                  `idl:"name:wszGPOName;string" json:"gpo_name"`
	MetadataReserved               uint32                  `idl:"name:MetaDataReserved" json:"metadata_reserved"`
	Metadata                       []*ObjectMetadata       `idl:"name:pMetaData;size_is:(((MetaDataReserved&1)?1:0))" json:"metadata"`
	LocalUserAuthorizationList     string                  `idl:"name:wszLocalUserAuthorizationList;string" json:"local_user_authorization_list"`
	PackageID                      string                  `idl:"name:wszPackageId;string" json:"package_id"`
	LocalUserOwner                 string                  `idl:"name:wszLocalUserOwner;string" json:"local_user_owner"`
	TrustTupleKeywords             uint32                  `idl:"name:dwTrustTupleKeywords" json:"trust_tuple_keywords"`
	OnNetworkNames                 *NetworkNames           `idl:"name:OnNetworkNames" json:"on_network_names"`
	SecurityRealmID                string                  `idl:"name:wszSecurityRealmId;string" json:"security_realm_id"`
	Flags2                         uint16                  `idl:"name:wFlags2" json:"flags2"`
	RemoteOutServerNames           *NetworkNames           `idl:"name:RemoteOutServerNames" json:"remote_out_server_names"`
}

func (o *Rule226) xxx_PreparePayload(ctx context.Context) error {
	if o.Metadata != nil && o.MetadataReserved == 0 {
		o.MetadataReserved = uint32(len(o.Metadata))
	}
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.Direction > Direction(2) {
		return fmt.Errorf("Direction is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.LocalApplication) > int(10001) {
		return fmt.Errorf("LocalApplication is out of range")
	}
	if len(o.LocalService) > int(10001) {
		return fmt.Errorf("LocalService is out of range")
	}
	if o.Action > RuleAction(4) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.RemoteMachineAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteMachineAuthorizationList is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(6) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if len(o.LocalUserAuthorizationList) > int(10001) {
		return fmt.Errorf("LocalUserAuthorizationList is out of range")
	}
	if len(o.PackageID) > int(10001) {
		return fmt.Errorf("PackageID is out of range")
	}
	if len(o.LocalUserOwner) > int(10001) {
		return fmt.Errorf("LocalUserOwner is out of range")
	}
	if len(o.SecurityRealmID) > int(10001) {
		return fmt.Errorf("SecurityRealmID is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule226) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Rule226{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Direction)); err != nil {
		return err
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if o.IPProtocolData != nil {
		if err := o.IPProtocolData.MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	} else {
		if err := (&Rule226_IPProtocolData{}).MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	}
	if o.LocalAddresses != nil {
		if err := o.LocalAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemoteAddresses != nil {
		if err := o.RemoteAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if o.LocalApplication != "" {
		_ptr_wszLocalApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalApplication); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalApplication, _ptr_wszLocalApplication); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalService != "" {
		_ptr_wszLocalService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalService); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalService, _ptr_wszLocalService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.RemoteMachineAuthorizationList != "" {
		_ptr_wszRemoteMachineAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteMachineAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteUserAuthorizationList != "" {
		_ptr_wszRemoteUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.RemoteUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	_exprMetaDataReserved := uint32(0)
	if (o.MetadataReserved & 1) != 0 {
		_exprMetaDataReserved = uint32(1)
	} else {
		_exprMetaDataReserved = uint32(0)
	}
	if o.Metadata != nil || _exprMetaDataReserved > 0 {
		_ptr_pMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprMetaDataReserved := uint64(0)
			if (o.MetadataReserved & 1) != 0 {
				_exprMetaDataReserved = uint64(1)
			} else {
				_exprMetaDataReserved = uint64(0)
			}
			dimSize1 := uint64(_exprMetaDataReserved)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Metadata {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Metadata[i1] != nil {
					if err := o.Metadata[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Metadata); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Metadata, _ptr_pMetaData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserAuthorizationList != "" {
		_ptr_wszLocalUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PackageID != "" {
		_ptr_wszPackageId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PackageID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PackageID, _ptr_wszPackageId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserOwner != "" {
		_ptr_wszLocalUserOwner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserOwner); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TrustTupleKeywords); err != nil {
		return err
	}
	if o.OnNetworkNames != nil {
		if err := o.OnNetworkNames.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NetworkNames{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SecurityRealmID != "" {
		_ptr_wszSecurityRealmId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SecurityRealmID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityRealmID, _ptr_wszSecurityRealmId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags2); err != nil {
		return err
	}
	if o.RemoteOutServerNames != nil {
		if err := o.RemoteOutServerNames.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NetworkNames{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule226) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &Rule226{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**Rule226) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Direction)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	if o.IPProtocolData == nil {
		o.IPProtocolData = &Rule226_IPProtocolData{}
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if err := o.IPProtocolData.UnmarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
		return err
	}
	if o.LocalAddresses == nil {
		o.LocalAddresses = &Addresses{}
	}
	if err := o.LocalAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemoteAddresses == nil {
		o.RemoteAddresses = &Addresses{}
	}
	if err := o.RemoteAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	_ptr_wszLocalApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalApplication); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalApplication := func(ptr interface{}) { o.LocalApplication = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalApplication, _s_wszLocalApplication, _ptr_wszLocalApplication); err != nil {
		return err
	}
	_ptr_wszLocalService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalService); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalService := func(ptr interface{}) { o.LocalService = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalService, _s_wszLocalService, _ptr_wszLocalService); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszRemoteMachineAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteMachineAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteMachineAuthorizationList := func(ptr interface{}) { o.RemoteMachineAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteMachineAuthorizationList, _s_wszRemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
		return err
	}
	_ptr_wszRemoteUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.RemoteUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteUserAuthorizationList := func(ptr interface{}) { o.RemoteUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteUserAuthorizationList, _s_wszRemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	_ptr_pMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprpMetaData := uint64(0)
		if (o.MetadataReserved & 1) != 0 {
			_exprpMetaData = uint64(1)
		} else {
			_exprpMetaData = uint64(0)
		}
		if _exprpMetaData > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprpMetaData)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Metadata", sizeInfo[0])
		}
		o.Metadata = make([]*ObjectMetadata, sizeInfo[0])
		for i1 := range o.Metadata {
			i1 := i1
			if o.Metadata[i1] == nil {
				o.Metadata[i1] = &ObjectMetadata{}
			}
			if err := o.Metadata[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMetaData := func(ptr interface{}) { o.Metadata = *ptr.(*[]*ObjectMetadata) }
	if err := w.ReadPointer(&o.Metadata, _s_pMetaData, _ptr_pMetaData); err != nil {
		return err
	}
	_ptr_wszLocalUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserAuthorizationList := func(ptr interface{}) { o.LocalUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserAuthorizationList, _s_wszLocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszPackageId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PackageID); err != nil {
			return err
		}
		return nil
	})
	_s_wszPackageId := func(ptr interface{}) { o.PackageID = *ptr.(*string) }
	if err := w.ReadPointer(&o.PackageID, _s_wszPackageId, _ptr_wszPackageId); err != nil {
		return err
	}
	_ptr_wszLocalUserOwner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserOwner); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserOwner := func(ptr interface{}) { o.LocalUserOwner = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserOwner, _s_wszLocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustTupleKeywords); err != nil {
		return err
	}
	if o.OnNetworkNames == nil {
		o.OnNetworkNames = &NetworkNames{}
	}
	if err := o.OnNetworkNames.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_wszSecurityRealmId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SecurityRealmID); err != nil {
			return err
		}
		return nil
	})
	_s_wszSecurityRealmId := func(ptr interface{}) { o.SecurityRealmID = *ptr.(*string) }
	if err := w.ReadPointer(&o.SecurityRealmID, _s_wszSecurityRealmId, _ptr_wszSecurityRealmId); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags2); err != nil {
		return err
	}
	if o.RemoteOutServerNames == nil {
		o.RemoteOutServerNames = &NetworkNames{}
	}
	if err := o.RemoteOutServerNames.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule226_IPProtocolData structure represents FW_RULE2_26 union anonymous member.
type Rule226_IPProtocolData struct {
	// Types that are assignable to Value
	//
	// *Rule226_IPProtocolData_Ports
	// *Rule226_IPProtocolData_TypeCodeListV4
	// *Rule226_IPProtocolData_TypeCodeListV6
	Value is_Rule226_IPProtocolData `json:"value"`
}

func (o *Rule226_IPProtocolData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Rule226_IPProtocolData_Ports:
		if value != nil {
			return value
		}
	case *Rule226_IPProtocolData_TypeCodeListV4:
		if value != nil {
			return value.TypeCodeListV4
		}
	case *Rule226_IPProtocolData_TypeCodeListV6:
		if value != nil {
			return value.TypeCodeListV6
		}
	}
	return nil
}

type is_Rule226_IPProtocolData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Rule226_IPProtocolData()
}

func (o *Rule226_IPProtocolData) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Rule226_IPProtocolData_Ports:
		switch sw {
		case uint16(6),
			uint16(17):
			return sw
		}
		return uint16(6)
	case *Rule226_IPProtocolData_TypeCodeListV4:
		return uint16(1)
	case *Rule226_IPProtocolData_TypeCodeListV6:
		return uint16(58)
	}
	return uint16(0)
}

func (o *Rule226_IPProtocolData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		_o, _ := o.Value.(*Rule226_IPProtocolData_Ports)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule226_IPProtocolData_Ports{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Rule226_IPProtocolData_TypeCodeListV4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule226_IPProtocolData_TypeCodeListV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(58):
		_o, _ := o.Value.(*Rule226_IPProtocolData_TypeCodeListV6)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule226_IPProtocolData_TypeCodeListV6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *Rule226_IPProtocolData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		o.Value = &Rule226_IPProtocolData_Ports{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Rule226_IPProtocolData_TypeCodeListV4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(58):
		o.Value = &Rule226_IPProtocolData_TypeCodeListV6{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// Rule226_IPProtocolData_Ports structure represents Rule226_IPProtocolData RPC union arm.
//
// It has following labels: 6, 17
type Rule226_IPProtocolData_Ports struct {
	LocalPorts  *Ports `idl:"name:LocalPorts" json:"local_ports"`
	RemotePorts *Ports `idl:"name:RemotePorts" json:"remote_ports"`
}

func (*Rule226_IPProtocolData_Ports) is_Rule226_IPProtocolData() {}

func (o *Rule226_IPProtocolData_Ports) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LocalPorts != nil {
		if err := o.LocalPorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemotePorts != nil {
		if err := o.RemotePorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule226_IPProtocolData_Ports) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LocalPorts == nil {
		o.LocalPorts = &Ports{}
	}
	if err := o.LocalPorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemotePorts == nil {
		o.RemotePorts = &Ports{}
	}
	if err := o.RemotePorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule226_IPProtocolData_TypeCodeListV4 structure represents Rule226_IPProtocolData RPC union arm.
//
// It has following labels: 1
type Rule226_IPProtocolData_TypeCodeListV4 struct {
	TypeCodeListV4 *ICMPTypeCodeList `idl:"name:V4TypeCodeList" json:"type_code_list_v4"`
}

func (*Rule226_IPProtocolData_TypeCodeListV4) is_Rule226_IPProtocolData() {}

func (o *Rule226_IPProtocolData_TypeCodeListV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV4 != nil {
		if err := o.TypeCodeListV4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule226_IPProtocolData_TypeCodeListV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV4 == nil {
		o.TypeCodeListV4 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule226_IPProtocolData_TypeCodeListV6 structure represents Rule226_IPProtocolData RPC union arm.
//
// It has following labels: 58
type Rule226_IPProtocolData_TypeCodeListV6 struct {
	TypeCodeListV6 *ICMPTypeCodeList `idl:"name:V6TypeCodeList" json:"type_code_list_v6"`
}

func (*Rule226_IPProtocolData_TypeCodeListV6) is_Rule226_IPProtocolData() {}

func (o *Rule226_IPProtocolData_TypeCodeListV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV6 != nil {
		if err := o.TypeCodeListV6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule226_IPProtocolData_TypeCodeListV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV6 == nil {
		o.TypeCodeListV6 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule227 structure represents FW_RULE2_27 RPC structure.
type Rule227 struct {
	Next                           *Rule227                `idl:"name:pNext" json:"next"`
	SchemaVersion                  uint16                  `idl:"name:wSchemaVersion" json:"schema_version"`
	RuleID                         string                  `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	Name                           string                  `idl:"name:wszName;string" json:"name"`
	Description                    string                  `idl:"name:wszDescription;string" json:"description"`
	Profiles                       uint32                  `idl:"name:dwProfiles" json:"profiles"`
	Direction                      Direction               `idl:"name:Direction" json:"direction"`
	IPProtocol                     uint16                  `idl:"name:wIpProtocol" json:"ip_protocol"`
	IPProtocolData                 *Rule227_IPProtocolData `idl:"name:IpProtocolData;switch_is:wIpProtocol" json:"ip_protocol_data"`
	LocalAddresses                 *Addresses              `idl:"name:LocalAddresses" json:"local_addresses"`
	RemoteAddresses                *Addresses              `idl:"name:RemoteAddresses" json:"remote_addresses"`
	LocalInterfaceIDs              *InterfaceLUIDs         `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	LocalInterfaceTypes            uint32                  `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	LocalApplication               string                  `idl:"name:wszLocalApplication;string" json:"local_application"`
	LocalService                   string                  `idl:"name:wszLocalService;string" json:"local_service"`
	Action                         RuleAction              `idl:"name:Action" json:"action"`
	Flags                          uint16                  `idl:"name:wFlags" json:"flags"`
	RemoteMachineAuthorizationList string                  `idl:"name:wszRemoteMachineAuthorizationList;string" json:"remote_machine_authorization_list"`
	RemoteUserAuthorizationList    string                  `idl:"name:wszRemoteUserAuthorizationList;string" json:"remote_user_authorization_list"`
	EmbeddedContext                string                  `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	PlatformValidityList           *OSPlatformList         `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	Status                         RuleStatus              `idl:"name:Status" json:"status"`
	Origin                         RuleOriginType          `idl:"name:Origin" json:"origin"`
	GPOName                        string                  `idl:"name:wszGPOName;string" json:"gpo_name"`
	MetadataReserved               uint32                  `idl:"name:MetaDataReserved" json:"metadata_reserved"`
	Metadata                       []*ObjectMetadata       `idl:"name:pMetaData;size_is:(((MetaDataReserved&1)?1:0))" json:"metadata"`
	LocalUserAuthorizationList     string                  `idl:"name:wszLocalUserAuthorizationList;string" json:"local_user_authorization_list"`
	PackageID                      string                  `idl:"name:wszPackageId;string" json:"package_id"`
	LocalUserOwner                 string                  `idl:"name:wszLocalUserOwner;string" json:"local_user_owner"`
	TrustTupleKeywords             uint32                  `idl:"name:dwTrustTupleKeywords" json:"trust_tuple_keywords"`
	OnNetworkNames                 *NetworkNames           `idl:"name:OnNetworkNames" json:"on_network_names"`
	SecurityRealmID                string                  `idl:"name:wszSecurityRealmId;string" json:"security_realm_id"`
	Flags2                         uint16                  `idl:"name:wFlags2" json:"flags2"`
	RemoteOutServerNames           *NetworkNames           `idl:"name:RemoteOutServerNames" json:"remote_out_server_names"`
	FQBN                           string                  `idl:"name:wszFqbn;string" json:"fqbn"`
	CompartmentID                  uint32                  `idl:"name:compartmentId" json:"compartment_id"`
}

func (o *Rule227) xxx_PreparePayload(ctx context.Context) error {
	if o.Metadata != nil && o.MetadataReserved == 0 {
		o.MetadataReserved = uint32(len(o.Metadata))
	}
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.Direction > Direction(2) {
		return fmt.Errorf("Direction is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.LocalApplication) > int(10001) {
		return fmt.Errorf("LocalApplication is out of range")
	}
	if len(o.LocalService) > int(10001) {
		return fmt.Errorf("LocalService is out of range")
	}
	if o.Action > RuleAction(4) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.RemoteMachineAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteMachineAuthorizationList is out of range")
	}
	if len(o.RemoteUserAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteUserAuthorizationList is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(6) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if len(o.LocalUserAuthorizationList) > int(10001) {
		return fmt.Errorf("LocalUserAuthorizationList is out of range")
	}
	if len(o.PackageID) > int(10001) {
		return fmt.Errorf("PackageID is out of range")
	}
	if len(o.LocalUserOwner) > int(10001) {
		return fmt.Errorf("LocalUserOwner is out of range")
	}
	if len(o.SecurityRealmID) > int(10001) {
		return fmt.Errorf("SecurityRealmID is out of range")
	}
	if len(o.FQBN) > int(10001) {
		return fmt.Errorf("FQBN is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule227) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Rule227{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Direction)); err != nil {
		return err
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if o.IPProtocolData != nil {
		if err := o.IPProtocolData.MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	} else {
		if err := (&Rule227_IPProtocolData{}).MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	}
	if o.LocalAddresses != nil {
		if err := o.LocalAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemoteAddresses != nil {
		if err := o.RemoteAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if o.LocalApplication != "" {
		_ptr_wszLocalApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalApplication); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalApplication, _ptr_wszLocalApplication); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalService != "" {
		_ptr_wszLocalService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalService); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalService, _ptr_wszLocalService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.RemoteMachineAuthorizationList != "" {
		_ptr_wszRemoteMachineAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteMachineAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteUserAuthorizationList != "" {
		_ptr_wszRemoteUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	_exprMetaDataReserved := uint32(0)
	if (o.MetadataReserved & 1) != 0 {
		_exprMetaDataReserved = uint32(1)
	} else {
		_exprMetaDataReserved = uint32(0)
	}
	if o.Metadata != nil || _exprMetaDataReserved > 0 {
		_ptr_pMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprMetaDataReserved := uint64(0)
			if (o.MetadataReserved & 1) != 0 {
				_exprMetaDataReserved = uint64(1)
			} else {
				_exprMetaDataReserved = uint64(0)
			}
			dimSize1 := uint64(_exprMetaDataReserved)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Metadata {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Metadata[i1] != nil {
					if err := o.Metadata[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Metadata); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Metadata, _ptr_pMetaData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserAuthorizationList != "" {
		_ptr_wszLocalUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PackageID != "" {
		_ptr_wszPackageId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PackageID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PackageID, _ptr_wszPackageId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserOwner != "" {
		_ptr_wszLocalUserOwner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserOwner); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TrustTupleKeywords); err != nil {
		return err
	}
	if o.OnNetworkNames != nil {
		if err := o.OnNetworkNames.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NetworkNames{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SecurityRealmID != "" {
		_ptr_wszSecurityRealmId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SecurityRealmID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityRealmID, _ptr_wszSecurityRealmId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags2); err != nil {
		return err
	}
	if o.RemoteOutServerNames != nil {
		if err := o.RemoteOutServerNames.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NetworkNames{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FQBN != "" {
		_ptr_wszFqbn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FQBN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FQBN, _ptr_wszFqbn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CompartmentID); err != nil {
		return err
	}
	return nil
}
func (o *Rule227) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &Rule227{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**Rule227) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Direction)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	if o.IPProtocolData == nil {
		o.IPProtocolData = &Rule227_IPProtocolData{}
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if err := o.IPProtocolData.UnmarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
		return err
	}
	if o.LocalAddresses == nil {
		o.LocalAddresses = &Addresses{}
	}
	if err := o.LocalAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemoteAddresses == nil {
		o.RemoteAddresses = &Addresses{}
	}
	if err := o.RemoteAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	_ptr_wszLocalApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalApplication); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalApplication := func(ptr interface{}) { o.LocalApplication = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalApplication, _s_wszLocalApplication, _ptr_wszLocalApplication); err != nil {
		return err
	}
	_ptr_wszLocalService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalService); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalService := func(ptr interface{}) { o.LocalService = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalService, _s_wszLocalService, _ptr_wszLocalService); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszRemoteMachineAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteMachineAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteMachineAuthorizationList := func(ptr interface{}) { o.RemoteMachineAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteMachineAuthorizationList, _s_wszRemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
		return err
	}
	_ptr_wszRemoteUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteUserAuthorizationList := func(ptr interface{}) { o.RemoteUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteUserAuthorizationList, _s_wszRemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	_ptr_pMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprpMetaData := uint64(0)
		if (o.MetadataReserved & 1) != 0 {
			_exprpMetaData = uint64(1)
		} else {
			_exprpMetaData = uint64(0)
		}
		if _exprpMetaData > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprpMetaData)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Metadata", sizeInfo[0])
		}
		o.Metadata = make([]*ObjectMetadata, sizeInfo[0])
		for i1 := range o.Metadata {
			i1 := i1
			if o.Metadata[i1] == nil {
				o.Metadata[i1] = &ObjectMetadata{}
			}
			if err := o.Metadata[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMetaData := func(ptr interface{}) { o.Metadata = *ptr.(*[]*ObjectMetadata) }
	if err := w.ReadPointer(&o.Metadata, _s_pMetaData, _ptr_pMetaData); err != nil {
		return err
	}
	_ptr_wszLocalUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserAuthorizationList := func(ptr interface{}) { o.LocalUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserAuthorizationList, _s_wszLocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszPackageId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PackageID); err != nil {
			return err
		}
		return nil
	})
	_s_wszPackageId := func(ptr interface{}) { o.PackageID = *ptr.(*string) }
	if err := w.ReadPointer(&o.PackageID, _s_wszPackageId, _ptr_wszPackageId); err != nil {
		return err
	}
	_ptr_wszLocalUserOwner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserOwner); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserOwner := func(ptr interface{}) { o.LocalUserOwner = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserOwner, _s_wszLocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustTupleKeywords); err != nil {
		return err
	}
	if o.OnNetworkNames == nil {
		o.OnNetworkNames = &NetworkNames{}
	}
	if err := o.OnNetworkNames.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_wszSecurityRealmId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SecurityRealmID); err != nil {
			return err
		}
		return nil
	})
	_s_wszSecurityRealmId := func(ptr interface{}) { o.SecurityRealmID = *ptr.(*string) }
	if err := w.ReadPointer(&o.SecurityRealmID, _s_wszSecurityRealmId, _ptr_wszSecurityRealmId); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags2); err != nil {
		return err
	}
	if o.RemoteOutServerNames == nil {
		o.RemoteOutServerNames = &NetworkNames{}
	}
	if err := o.RemoteOutServerNames.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_wszFqbn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FQBN); err != nil {
			return err
		}
		return nil
	})
	_s_wszFqbn := func(ptr interface{}) { o.FQBN = *ptr.(*string) }
	if err := w.ReadPointer(&o.FQBN, _s_wszFqbn, _ptr_wszFqbn); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompartmentID); err != nil {
		return err
	}
	return nil
}

// Rule227_IPProtocolData structure represents FW_RULE2_27 union anonymous member.
type Rule227_IPProtocolData struct {
	// Types that are assignable to Value
	//
	// *Rule227_IPProtocolData_Ports
	// *Rule227_IPProtocolData_TypeCodeListV4
	// *Rule227_IPProtocolData_TypeCodeListV6
	Value is_Rule227_IPProtocolData `json:"value"`
}

func (o *Rule227_IPProtocolData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Rule227_IPProtocolData_Ports:
		if value != nil {
			return value
		}
	case *Rule227_IPProtocolData_TypeCodeListV4:
		if value != nil {
			return value.TypeCodeListV4
		}
	case *Rule227_IPProtocolData_TypeCodeListV6:
		if value != nil {
			return value.TypeCodeListV6
		}
	}
	return nil
}

type is_Rule227_IPProtocolData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Rule227_IPProtocolData()
}

func (o *Rule227_IPProtocolData) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Rule227_IPProtocolData_Ports:
		switch sw {
		case uint16(6),
			uint16(17):
			return sw
		}
		return uint16(6)
	case *Rule227_IPProtocolData_TypeCodeListV4:
		return uint16(1)
	case *Rule227_IPProtocolData_TypeCodeListV6:
		return uint16(58)
	}
	return uint16(0)
}

func (o *Rule227_IPProtocolData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		_o, _ := o.Value.(*Rule227_IPProtocolData_Ports)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule227_IPProtocolData_Ports{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Rule227_IPProtocolData_TypeCodeListV4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule227_IPProtocolData_TypeCodeListV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(58):
		_o, _ := o.Value.(*Rule227_IPProtocolData_TypeCodeListV6)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule227_IPProtocolData_TypeCodeListV6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *Rule227_IPProtocolData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		o.Value = &Rule227_IPProtocolData_Ports{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Rule227_IPProtocolData_TypeCodeListV4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(58):
		o.Value = &Rule227_IPProtocolData_TypeCodeListV6{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// Rule227_IPProtocolData_Ports structure represents Rule227_IPProtocolData RPC union arm.
//
// It has following labels: 6, 17
type Rule227_IPProtocolData_Ports struct {
	LocalPorts  *Ports `idl:"name:LocalPorts" json:"local_ports"`
	RemotePorts *Ports `idl:"name:RemotePorts" json:"remote_ports"`
}

func (*Rule227_IPProtocolData_Ports) is_Rule227_IPProtocolData() {}

func (o *Rule227_IPProtocolData_Ports) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LocalPorts != nil {
		if err := o.LocalPorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemotePorts != nil {
		if err := o.RemotePorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule227_IPProtocolData_Ports) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LocalPorts == nil {
		o.LocalPorts = &Ports{}
	}
	if err := o.LocalPorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemotePorts == nil {
		o.RemotePorts = &Ports{}
	}
	if err := o.RemotePorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule227_IPProtocolData_TypeCodeListV4 structure represents Rule227_IPProtocolData RPC union arm.
//
// It has following labels: 1
type Rule227_IPProtocolData_TypeCodeListV4 struct {
	TypeCodeListV4 *ICMPTypeCodeList `idl:"name:V4TypeCodeList" json:"type_code_list_v4"`
}

func (*Rule227_IPProtocolData_TypeCodeListV4) is_Rule227_IPProtocolData() {}

func (o *Rule227_IPProtocolData_TypeCodeListV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV4 != nil {
		if err := o.TypeCodeListV4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule227_IPProtocolData_TypeCodeListV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV4 == nil {
		o.TypeCodeListV4 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule227_IPProtocolData_TypeCodeListV6 structure represents Rule227_IPProtocolData RPC union arm.
//
// It has following labels: 58
type Rule227_IPProtocolData_TypeCodeListV6 struct {
	TypeCodeListV6 *ICMPTypeCodeList `idl:"name:V6TypeCodeList" json:"type_code_list_v6"`
}

func (*Rule227_IPProtocolData_TypeCodeListV6) is_Rule227_IPProtocolData() {}

func (o *Rule227_IPProtocolData_TypeCodeListV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV6 != nil {
		if err := o.TypeCodeListV6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule227_IPProtocolData_TypeCodeListV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV6 == nil {
		o.TypeCodeListV6 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DynamicKeywordAddressIDList structure represents FW_DYNAMIC_KEYWORD_ADDRESS_ID_LIST RPC structure.
type DynamicKeywordAddressIDList struct {
	IDsLength uint32   `idl:"name:dwNumIds" json:"ids_length"`
	IDs       []uint32 `idl:"name:ids;size_is:(dwNumIds)" json:"ids"`
}

func (o *DynamicKeywordAddressIDList) xxx_PreparePayload(ctx context.Context) error {
	if o.IDs != nil && o.IDsLength == 0 {
		o.IDsLength = uint32(len(o.IDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DynamicKeywordAddressIDList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.IDsLength); err != nil {
		return err
	}
	if o.IDs != nil || o.IDsLength > 0 {
		_ptr_ids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.IDsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.IDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.IDs[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.IDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IDs, _ptr_ids); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DynamicKeywordAddressIDList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.IDsLength); err != nil {
		return err
	}
	_ptr_ids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.IDsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.IDsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.IDs", sizeInfo[0])
		}
		o.IDs = make([]uint32, sizeInfo[0])
		for i1 := range o.IDs {
			i1 := i1
			if err := w.ReadData(&o.IDs[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ids := func(ptr interface{}) { o.IDs = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.IDs, _s_ids, _ptr_ids); err != nil {
		return err
	}
	return nil
}

// Rule structure represents FW_RULE RPC structure.
//
// This structure is used to represent a firewall rule.
type Rule struct {
	// pNext:  A pointer to the next FW_RULE in the list.
	Next *Rule `idl:"name:pNext" json:"next"`
	// wSchemaVersion:  Specifies the version of the rule.
	SchemaVersion uint16 `idl:"name:wSchemaVersion" json:"schema_version"`
	// wszRuleId:  A pointer to a Unicode string that uniquely identifies the rule.
	RuleID string `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	// wszName:  A pointer to a Unicode string that provides a friendly name for the rule.
	Name string `idl:"name:wszName;string" json:"name"`
	// wszDescription:  A pointer to a Unicode string that provides a friendly description
	// for the rule.
	Description string `idl:"name:wszDescription;string" json:"description"`
	// dwProfiles:  A bitmask of the FW_PROFILE_TYPE flags. It is a condition that matches
	// traffic on the specified profiles.
	Profiles uint32 `idl:"name:dwProfiles" json:"profiles"`
	// Direction:  Specifies the direction of the traffic that the rule matches.
	Direction Direction `idl:"name:Direction" json:"direction"`
	// wIpProtocol:  A condition that specifies the protocol of the traffic that the rule
	// matches. If the value is within the range 0 to 255, the value describes a protocol
	// in IETF IANA numbers (for more information, see [IANA-PROTO-NUM]). If the value is
	// 256, the rule matches any protocol.
	IPProtocol     uint16               `idl:"name:wIpProtocol" json:"ip_protocol"`
	IPProtocolData *Rule_IPProtocolData `idl:"name:IpProtocolData;switch_is:wIpProtocol" json:"ip_protocol_data"`
	// LocalAddresses:  A condition that specifies the addresses of the local host of the
	// traffic that the rule matches. An empty LocalAddresses structure means that this
	// condition is not applied.
	LocalAddresses *Addresses `idl:"name:LocalAddresses" json:"local_addresses"`
	// RemoteAddresses:  A condition that specifies the addresses of the remote host of
	// the traffic that the rule matches. An empty RemoteAddresses structure means that
	// this condition is not applied.
	RemoteAddresses *Addresses `idl:"name:RemoteAddresses" json:"remote_addresses"`
	// LocalInterfaceIds:  A condition that specifies the list of specific network interfaces
	// used by the traffic that the rule matches. A LocalInterfaceIds field with no interface
	// GUID specified means that the rule applies to all interfaces; that is, the condition
	// is not applied.
	LocalInterfaceIDs *InterfaceLUIDs `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	// dwLocalInterfaceTypes:  A bitmask of FW_INTERFACE_TYPE. It is a condition that restricts
	// the interface types that are used by the traffic that the rule matches. 0x00000000
	// means that the condition matches all interface types.
	LocalInterfaceTypes uint32 `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	// wszLocalApplication:  A pointer to a Unicode string. It is a condition that specifies
	// a file path name to the executable that uses the traffic that the rule matches. A
	// null in this field means that the rule applies to all processes in the host.
	LocalApplication string `idl:"name:wszLocalApplication;string" json:"local_application"`
	// wszLocalService:  A pointer to a Unicode string. It is a condition that specifies
	// the service name of the service that uses the traffic that the rule matches. An L"*"
	// string in this field means that the rule applies to all services in the system. A
	// null in this field means that the rule applies to all processes.
	LocalService string `idl:"name:wszLocalService;string" json:"local_service"`
	// Action:  The action that the rule will take for the traffic matches.
	Action RuleAction `idl:"name:Action" json:"action"`
	// wFlags:  Bit flags from FW_RULE_FLAGS.
	Flags uint16 `idl:"name:wFlags" json:"flags"`
	// wszRemoteMachineAuthorizationList:  A pointer to a Unicode string. A condition that
	// specifies the remote machines sending or receiving the traffic that the rule matches.
	// The string is in SDDL format ([MS-DTYP] section 2.5.1).
	RemoteMachineAuthorizationList string `idl:"name:wszRemoteMachineAuthorizationList;string" json:"remote_machine_authorization_list"`
	// wszRemoteUserAuthorizationList:  A pointer to a Unicode string. A condition that
	// specifies the remote users accepting or receiving the traffic that the rule matches.
	// The string is in SDDL format ([MS-DTYP] section 2.5.1).
	RemoteUserAuthorizationList string `idl:"name:wszRemoteUserAuthorizationList;string" json:"remote_user_authorization_list"`
	// wszEmbeddedContext:  A pointer to a Unicode string. It specifies a group name for
	// this rule. Other components in the system use this string to enable or disable groups
	// of rules by verifying that they all have the same group name.
	EmbeddedContext string `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	// PlatformValidityList:  A condition in a rule that determines whether or not the rule
	// is enforced by the local computer based on the local computer's platform information.
	// The rule is enforced only if the local computer's operating system platform is an
	// element of the set described by PlatformValidityList.<6>
	PlatformValidityList *OSPlatformList `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	// Status:  The status code of the rule, as specified by the FW_RULE_STATUS enumeration.
	// This field is filled out when the structure is returned as output. On input, this
	// field MUST be set to FW_RULE_STATUS_OK.
	Status RuleStatus `idl:"name:Status" json:"status"`
	// Origin:  The rule origin, as specified in the FW_RULE_ORIGIN_TYPE enumeration. It
	// MUST be filled on enumerated rules and ignored on input.
	Origin RuleOriginType `idl:"name:Origin" json:"origin"`
	// wszGPOName:  A pointer to a Unicode string containing the displayName of the GPO
	// containing this object. When adding a new object, this field is not used. The client
	// SHOULD set the value to NULL, and the server MUST ignore the value. When enumerating
	// an existing object, if the client does not set the FW_ENUM_RULES_FLAG_RESOLVE_GPO_NAME
	// flag, the server MUST set the value to NULL. Otherwise, the server MUST set the value
	// to the displayName of the GPO containing the object or NULL if the object is not
	// contained within a GPO. For details about how the server initializes an object from
	// a GPO, see section 3.1.3. For details about how the displayName of a GPO is stored,
	// see [MS-GPOL] section 2.3.
	GPOName          string `idl:"name:wszGPOName;string" json:"gpo_name"`
	MetadataReserved uint32 `idl:"name:MetaDataReserved" json:"metadata_reserved"`
	// pMetaData:  A pointer to an FW_OBJECT_METADATA structure that contains specific metadata
	// about the current state of the firewall rule.
	Metadata []*ObjectMetadata `idl:"name:pMetaData;size_is:(((MetaDataReserved&1)?1:0))" json:"metadata"`
	// wszLocalUserAuthorizationList:  A pointer to a Unicode string in SDDL format ([MS-DTYP]
	// section 2.5.1). It is a condition that specifies the local users accepting or receiving
	// the traffic that the rule matches.
	LocalUserAuthorizationList string `idl:"name:wszLocalUserAuthorizationList;string" json:"local_user_authorization_list"`
	// wszPackageId:  A pointer to a Unicode string in SID string format ([MS-DTYP] section
	// 2.4.2.1). It is a condition that specifies the application SID of the process that
	// uses the traffic that the rule matches. A null in this field means that the rule
	// applies to all processes in the host.
	PackageID string `idl:"name:wszPackageId;string" json:"package_id"`
	// wszLocalUserOwner:  A pointer to a Unicode string in SID string format. The SID specifies
	// the security principal that owns the rule.
	LocalUserOwner string `idl:"name:wszLocalUserOwner;string" json:"local_user_owner"`
	// dwTrustTupleKeywords:  A bitmask of the FW_TRUST_TUPLE_KEYWORD flags. It is a condition
	// that matches traffic associated with the specified trust tuples.
	TrustTupleKeywords uint32 `idl:"name:dwTrustTupleKeywords" json:"trust_tuple_keywords"`
	// OnNetworkNames: Specifies the networks, identified by name, in which the rule must
	// be enforced.
	OnNetworkNames *NetworkNames `idl:"name:OnNetworkNames" json:"on_network_names"`
	// wszSecurityRealmId: A pointer to a Unicode string in SID string format. The SID specifies
	// the Security Realm ID, which identifies a security realm that this firewall rule
	// is associated with. Any application that matches this rule will be subject to the
	// IPsec polices for this security realm.
	SecurityRealmID string `idl:"name:wszSecurityRealmId;string" json:"security_realm_id"`
	// wFlags2: Bit flags from FW_RULE_FLAGS2 (section 2.2.103).
	Flags2 uint16 `idl:"name:wFlags2" json:"flags2"`
	// RemoteOutServerNames: This value is not used over the wire.
	RemoteOutServerNames *NetworkNames `idl:"name:RemoteOutServerNames" json:"remote_out_server_names"`
	// wszFqbn: A string that is formatted as an fully qualified binary name (FQBN); also
	// see [MSDN-FQBN].
	FQBN string `idl:"name:wszFqbn;string" json:"fqbn"`
	// compartmentId: The ID of the compartment or Windows Server Container.
	CompartmentID uint32 `idl:"name:compartmentId" json:"compartment_id"`
	// providerContextKey: This value MUST the all-zeroes GUID.
	ProviderContextKey *dtyp.GUID `idl:"name:providerContextKey" json:"provider_context_key"`
	// RemoteDynamicKeywordAddresses: A type that specifies the dynamic keyword address
	// Ids to be used for the remote host of the traffic matched by this rule.
	RemoteDynamicKeywordAddresses *DynamicKeywordAddressIDList `idl:"name:RemoteDynamicKeywordAddresses" json:"remote_dynamic_keyword_addresses"`
}

func (o *Rule) xxx_PreparePayload(ctx context.Context) error {
	if o.Metadata != nil && o.MetadataReserved == 0 {
		o.MetadataReserved = uint32(len(o.Metadata))
	}
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.Direction > Direction(2) {
		return fmt.Errorf("Direction is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.LocalApplication) > int(10001) {
		return fmt.Errorf("LocalApplication is out of range")
	}
	if len(o.LocalService) > int(10001) {
		return fmt.Errorf("LocalService is out of range")
	}
	if o.Action > RuleAction(4) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.RemoteMachineAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteMachineAuthorizationList is out of range")
	}
	if len(o.RemoteUserAuthorizationList) > int(10001) {
		return fmt.Errorf("RemoteUserAuthorizationList is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(6) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if len(o.LocalUserAuthorizationList) > int(10001) {
		return fmt.Errorf("LocalUserAuthorizationList is out of range")
	}
	if len(o.PackageID) > int(10001) {
		return fmt.Errorf("PackageID is out of range")
	}
	if len(o.LocalUserOwner) > int(10001) {
		return fmt.Errorf("LocalUserOwner is out of range")
	}
	if len(o.SecurityRealmID) > int(10001) {
		return fmt.Errorf("SecurityRealmID is out of range")
	}
	if len(o.FQBN) > int(10001) {
		return fmt.Errorf("FQBN is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Rule{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Direction)); err != nil {
		return err
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if o.IPProtocolData != nil {
		if err := o.IPProtocolData.MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	} else {
		if err := (&Rule_IPProtocolData{}).MarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
			return err
		}
	}
	if o.LocalAddresses != nil {
		if err := o.LocalAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemoteAddresses != nil {
		if err := o.RemoteAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if o.LocalApplication != "" {
		_ptr_wszLocalApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalApplication); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalApplication, _ptr_wszLocalApplication); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalService != "" {
		_ptr_wszLocalService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalService); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalService, _ptr_wszLocalService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.RemoteMachineAuthorizationList != "" {
		_ptr_wszRemoteMachineAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteMachineAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteUserAuthorizationList != "" {
		_ptr_wszRemoteUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	_exprMetaDataReserved := uint32(0)
	if (o.MetadataReserved & 1) != 0 {
		_exprMetaDataReserved = uint32(1)
	} else {
		_exprMetaDataReserved = uint32(0)
	}
	if o.Metadata != nil || _exprMetaDataReserved > 0 {
		_ptr_pMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprMetaDataReserved := uint64(0)
			if (o.MetadataReserved & 1) != 0 {
				_exprMetaDataReserved = uint64(1)
			} else {
				_exprMetaDataReserved = uint64(0)
			}
			dimSize1 := uint64(_exprMetaDataReserved)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Metadata {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Metadata[i1] != nil {
					if err := o.Metadata[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Metadata); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Metadata, _ptr_pMetaData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserAuthorizationList != "" {
		_ptr_wszLocalUserAuthorizationList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserAuthorizationList); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PackageID != "" {
		_ptr_wszPackageId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PackageID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PackageID, _ptr_wszPackageId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LocalUserOwner != "" {
		_ptr_wszLocalUserOwner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LocalUserOwner); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TrustTupleKeywords); err != nil {
		return err
	}
	if o.OnNetworkNames != nil {
		if err := o.OnNetworkNames.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NetworkNames{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SecurityRealmID != "" {
		_ptr_wszSecurityRealmId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SecurityRealmID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityRealmID, _ptr_wszSecurityRealmId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags2); err != nil {
		return err
	}
	if o.RemoteOutServerNames != nil {
		if err := o.RemoteOutServerNames.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NetworkNames{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FQBN != "" {
		_ptr_wszFqbn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FQBN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FQBN, _ptr_wszFqbn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CompartmentID); err != nil {
		return err
	}
	if o.ProviderContextKey != nil {
		if err := o.ProviderContextKey.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemoteDynamicKeywordAddresses != nil {
		if err := o.RemoteDynamicKeywordAddresses.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DynamicKeywordAddressIDList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &Rule{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**Rule) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Direction)); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	if o.IPProtocolData == nil {
		o.IPProtocolData = &Rule_IPProtocolData{}
	}
	_swIPProtocolData := uint16(o.IPProtocol)
	if err := o.IPProtocolData.UnmarshalUnionNDR(ctx, w, _swIPProtocolData); err != nil {
		return err
	}
	if o.LocalAddresses == nil {
		o.LocalAddresses = &Addresses{}
	}
	if err := o.LocalAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemoteAddresses == nil {
		o.RemoteAddresses = &Addresses{}
	}
	if err := o.RemoteAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	_ptr_wszLocalApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalApplication); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalApplication := func(ptr interface{}) { o.LocalApplication = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalApplication, _s_wszLocalApplication, _ptr_wszLocalApplication); err != nil {
		return err
	}
	_ptr_wszLocalService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalService); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalService := func(ptr interface{}) { o.LocalService = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalService, _s_wszLocalService, _ptr_wszLocalService); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszRemoteMachineAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteMachineAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteMachineAuthorizationList := func(ptr interface{}) { o.RemoteMachineAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteMachineAuthorizationList, _s_wszRemoteMachineAuthorizationList, _ptr_wszRemoteMachineAuthorizationList); err != nil {
		return err
	}
	_ptr_wszRemoteUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteUserAuthorizationList := func(ptr interface{}) { o.RemoteUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteUserAuthorizationList, _s_wszRemoteUserAuthorizationList, _ptr_wszRemoteUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	_ptr_pMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprpMetaData := uint64(0)
		if (o.MetadataReserved & 1) != 0 {
			_exprpMetaData = uint64(1)
		} else {
			_exprpMetaData = uint64(0)
		}
		if _exprpMetaData > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprpMetaData)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Metadata", sizeInfo[0])
		}
		o.Metadata = make([]*ObjectMetadata, sizeInfo[0])
		for i1 := range o.Metadata {
			i1 := i1
			if o.Metadata[i1] == nil {
				o.Metadata[i1] = &ObjectMetadata{}
			}
			if err := o.Metadata[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMetaData := func(ptr interface{}) { o.Metadata = *ptr.(*[]*ObjectMetadata) }
	if err := w.ReadPointer(&o.Metadata, _s_pMetaData, _ptr_pMetaData); err != nil {
		return err
	}
	_ptr_wszLocalUserAuthorizationList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserAuthorizationList); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserAuthorizationList := func(ptr interface{}) { o.LocalUserAuthorizationList = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserAuthorizationList, _s_wszLocalUserAuthorizationList, _ptr_wszLocalUserAuthorizationList); err != nil {
		return err
	}
	_ptr_wszPackageId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PackageID); err != nil {
			return err
		}
		return nil
	})
	_s_wszPackageId := func(ptr interface{}) { o.PackageID = *ptr.(*string) }
	if err := w.ReadPointer(&o.PackageID, _s_wszPackageId, _ptr_wszPackageId); err != nil {
		return err
	}
	_ptr_wszLocalUserOwner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LocalUserOwner); err != nil {
			return err
		}
		return nil
	})
	_s_wszLocalUserOwner := func(ptr interface{}) { o.LocalUserOwner = *ptr.(*string) }
	if err := w.ReadPointer(&o.LocalUserOwner, _s_wszLocalUserOwner, _ptr_wszLocalUserOwner); err != nil {
		return err
	}
	if err := w.ReadData(&o.TrustTupleKeywords); err != nil {
		return err
	}
	if o.OnNetworkNames == nil {
		o.OnNetworkNames = &NetworkNames{}
	}
	if err := o.OnNetworkNames.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_wszSecurityRealmId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SecurityRealmID); err != nil {
			return err
		}
		return nil
	})
	_s_wszSecurityRealmId := func(ptr interface{}) { o.SecurityRealmID = *ptr.(*string) }
	if err := w.ReadPointer(&o.SecurityRealmID, _s_wszSecurityRealmId, _ptr_wszSecurityRealmId); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags2); err != nil {
		return err
	}
	if o.RemoteOutServerNames == nil {
		o.RemoteOutServerNames = &NetworkNames{}
	}
	if err := o.RemoteOutServerNames.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_wszFqbn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FQBN); err != nil {
			return err
		}
		return nil
	})
	_s_wszFqbn := func(ptr interface{}) { o.FQBN = *ptr.(*string) }
	if err := w.ReadPointer(&o.FQBN, _s_wszFqbn, _ptr_wszFqbn); err != nil {
		return err
	}
	if err := w.ReadData(&o.CompartmentID); err != nil {
		return err
	}
	if o.ProviderContextKey == nil {
		o.ProviderContextKey = &dtyp.GUID{}
	}
	if err := o.ProviderContextKey.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemoteDynamicKeywordAddresses == nil {
		o.RemoteDynamicKeywordAddresses = &DynamicKeywordAddressIDList{}
	}
	if err := o.RemoteDynamicKeywordAddresses.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule_IPProtocolData structure represents FW_RULE union anonymous member.
//
// This structure is used to represent a firewall rule.
type Rule_IPProtocolData struct {
	// Types that are assignable to Value
	//
	// *Rule_IPProtocolData_Ports
	// *Rule_IPProtocolData_TypeCodeListV4
	// *Rule_IPProtocolData_TypeCodeListV6
	Value is_Rule_IPProtocolData `json:"value"`
}

func (o *Rule_IPProtocolData) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Rule_IPProtocolData_Ports:
		if value != nil {
			return value
		}
	case *Rule_IPProtocolData_TypeCodeListV4:
		if value != nil {
			return value.TypeCodeListV4
		}
	case *Rule_IPProtocolData_TypeCodeListV6:
		if value != nil {
			return value.TypeCodeListV6
		}
	}
	return nil
}

type is_Rule_IPProtocolData interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Rule_IPProtocolData()
}

func (o *Rule_IPProtocolData) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Rule_IPProtocolData_Ports:
		switch sw {
		case uint16(6),
			uint16(17):
			return sw
		}
		return uint16(6)
	case *Rule_IPProtocolData_TypeCodeListV4:
		return uint16(1)
	case *Rule_IPProtocolData_TypeCodeListV6:
		return uint16(58)
	}
	return uint16(0)
}

func (o *Rule_IPProtocolData) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		_o, _ := o.Value.(*Rule_IPProtocolData_Ports)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule_IPProtocolData_Ports{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Rule_IPProtocolData_TypeCodeListV4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule_IPProtocolData_TypeCodeListV4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(58):
		_o, _ := o.Value.(*Rule_IPProtocolData_TypeCodeListV6)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Rule_IPProtocolData_TypeCodeListV6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *Rule_IPProtocolData) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(6),
		uint16(17):
		o.Value = &Rule_IPProtocolData_Ports{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Rule_IPProtocolData_TypeCodeListV4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(58):
		o.Value = &Rule_IPProtocolData_TypeCodeListV6{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// Rule_IPProtocolData_Ports structure represents Rule_IPProtocolData RPC union arm.
//
// It has following labels: 6, 17
type Rule_IPProtocolData_Ports struct {
	// LocalPorts:  A condition that specifies the local host ports of the TCP or UDP traffic
	// that the rule matches.
	LocalPorts *Ports `idl:"name:LocalPorts" json:"local_ports"`
	// RemotePorts:  A condition that specifies the remote host ports of the TCP or UDP
	// traffic that the rule matches.
	RemotePorts *Ports `idl:"name:RemotePorts" json:"remote_ports"`
}

func (*Rule_IPProtocolData_Ports) is_Rule_IPProtocolData() {}

func (o *Rule_IPProtocolData_Ports) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LocalPorts != nil {
		if err := o.LocalPorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.RemotePorts != nil {
		if err := o.RemotePorts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule_IPProtocolData_Ports) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LocalPorts == nil {
		o.LocalPorts = &Ports{}
	}
	if err := o.LocalPorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.RemotePorts == nil {
		o.RemotePorts = &Ports{}
	}
	if err := o.RemotePorts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule_IPProtocolData_TypeCodeListV4 structure represents Rule_IPProtocolData RPC union arm.
//
// It has following labels: 1
type Rule_IPProtocolData_TypeCodeListV4 struct {
	// V4TypeCodeList:  A condition that specifies the list of ICMP types of the traffic
	// that the rule matches. This field applies only when wIpProtocol specifies ICMP v4.
	TypeCodeListV4 *ICMPTypeCodeList `idl:"name:V4TypeCodeList" json:"type_code_list_v4"`
}

func (*Rule_IPProtocolData_TypeCodeListV4) is_Rule_IPProtocolData() {}

func (o *Rule_IPProtocolData_TypeCodeListV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV4 != nil {
		if err := o.TypeCodeListV4.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule_IPProtocolData_TypeCodeListV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV4 == nil {
		o.TypeCodeListV4 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV4.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Rule_IPProtocolData_TypeCodeListV6 structure represents Rule_IPProtocolData RPC union arm.
//
// It has following labels: 58
type Rule_IPProtocolData_TypeCodeListV6 struct {
	// V6TypeCodeList:  A condition that specifies the list of ICMP types of the traffic
	// that the rule matches. This field applies only when wIpProtocol specifies ICMP v6.
	TypeCodeListV6 *ICMPTypeCodeList `idl:"name:V6TypeCodeList" json:"type_code_list_v6"`
}

func (*Rule_IPProtocolData_TypeCodeListV6) is_Rule_IPProtocolData() {}

func (o *Rule_IPProtocolData_TypeCodeListV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeCodeListV6 != nil {
		if err := o.TypeCodeListV6.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ICMPTypeCodeList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Rule_IPProtocolData_TypeCodeListV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.TypeCodeListV6 == nil {
		o.TypeCodeListV6 = &ICMPTypeCodeList{}
	}
	if err := o.TypeCodeListV6.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ProfileConfig type represents FW_PROFILE_CONFIG RPC enumeration.
//
// This enumeration identifies each of the per-profile configuration options supported
// by this protocol. Each configuration option has a merge law that is used to determine
// how to merge the values of these options across stores.
type ProfileConfig uint16

var (
	// FW_PROFILE_CONFIG_INVALID:  This value is invalid and MUST NOT be used. It is defined
	// for simplicity in writing IDL definitions and code. This symbolic constant has a
	// value of 0.
	ProfileConfigInvalid ProfileConfig = 0
	// FW_PROFILE_CONFIG_ENABLE_FW:  This value is an on/off switch for the firewall and
	// advanced security enforcement. It is a DWORD type value; 0x00000000 is off; 0x00000001
	// is on. If this value is off, the server MUST NOT block any network traffic, regardless
	// of other policy settings. The merge law for this option is to let the value of the
	// GroupPolicyRSoPStore win if it is configured; otherwise, the local store value is
	// used. This symbolic constant has a value of 1.
	ProfileConfigEnableFw ProfileConfig = 1
	// FW_PROFILE_CONFIG_DISABLE_STEALTH_MODE:  This value is a DWORD used as an on/off
	// switch. When this option is off, the server operates in stealth mode. The firewall
	// rules used to enforce stealth mode are implementation-specific.<9> The merge law
	// for this option is to let the value of the GroupPolicyRSoPStore win if it is configured;
	// otherwise, the local store value is used. This symbolic constant has a value of 2.
	ProfileConfigDisableStealthMode ProfileConfig = 2
	// FW_PROFILE_CONFIG_SHIELDED:  This value is a DWORD used as an on/off switch. If
	// this value is on and FW_PROFILE_CONFIG_ENABLE_FW is on, the server MUST block all
	// incoming traffic regardless of other policy settings. The merge law for this option
	// is to let "on" values win. This symbolic constant has a value of 3.
	ProfileConfigShielded ProfileConfig = 3
	// FW_PROFILE_CONFIG_DISABLE_UNICAST_RESPONSES_TO_MULTICAST_BROADCAST: This value is
	// a DWORD used as an on/off switch. If it is on, unicast responses to multicast broadcast
	// traffic is blocked. The merge law for this option is to let the value of the GroupPolicyRSoPStore
	// win if it is configured; otherwise, the local store value is used. This symbolic
	// constant has a value of 4.
	ProfileConfigDisableUnicastResponsesToMulticastBroadcast ProfileConfig = 4
	// FW_PROFILE_CONFIG_LOG_DROPPED_PACKETS:  This value is a DWORD used as an on/off
	// switch. If this value is on, the firewall logs all the dropped packets. The merge
	// law for this option is to let "on" values win. This symbolic constant has a value
	// of 5.
	ProfileConfigLogDroppedPackets ProfileConfig = 5
	// FW_PROFILE_CONFIG_LOG_SUCCESS_CONNECTIONS:  This value is a DWORD used as an on/off
	// switch. If this value is on, the firewall logs all successful inbound connections.
	// The merge law for this option is to let "on" values win. This symbolic constant has
	// a value of 6.
	ProfileConfigLogSuccessConnections ProfileConfig = 6
	// FW_PROFILE_CONFIG_LOG_IGNORED_RULES:  This value is a DWORD used as an on/off switch.
	// The server MAY use this value in an implementation-specific way to control logging
	// of events if a rule is not enforced for any reason. The merge law for this option
	// is to let "on" values win. This symbolic constant has a value of 7.<10>
	ProfileConfigLogIgnoredRules ProfileConfig = 7
	// FW_PROFILE_CONFIG_LOG_MAX_FILE_SIZE:  This value is a DWORD and specifies the size,
	// in kilobytes, of the log where dropped packets and successful connections are logged.
	// The merge law for this option is to let the value of the GroupPolicyRSoPStore win
	// if it is configured; otherwise, the local store value is used. This symbolic constant
	// has a value of 8.
	ProfileConfigLogMaxFileSize ProfileConfig = 8
	// FW_PROFILE_CONFIG_LOG_FILE_PATH:  This configuration value is a string that represents
	// a file path to the log for when the firewall logs dropped packets and successful
	// connections. The merge law for this option is to let the value of the GroupPolicyRSoPStore
	// win if it is configured; otherwise, the local store value is used. This symbolic
	// constant has a value of 9.
	ProfileConfigLogFilePath ProfileConfig = 9
	// FW_PROFILE_CONFIG_DISABLE_INBOUND_NOTIFICATIONS:  This value is a DWORD used as
	// an on/off switch. If this value is off, the firewall MAY display a notification to
	// the user when an application is blocked from listening on a port.<11>  If this value
	// is on, the firewall MUST NOT display such a notification. The merge law for this
	// option is to let the value of the GroupPolicyRSoPStore win if it is configured; otherwise,
	// the local store value is used. This symbolic constant has a value of 10.
	ProfileConfigDisableInboundNotifications ProfileConfig = 10
	// FW_PROFILE_CONFIG_AUTH_APPS_ALLOW_USER_PREF_MERGE: This value is a DWORD used as
	// an on/off switch. If this value is off, authorized application firewall rules in
	// the local store are ignored and not enforced. The merge law for this option is to
	// let the value of the GroupPolicyRSoPStore win if it is configured; otherwise, the
	// local store value is used. This symbolic constant has a value of 11.
	//
	// The authorized application firewall rules consist of the FW_RULE objects where all
	// of the following are true:
	//
	// wszLocalApplication is not NULL
	//
	// wszLocalService == NULL
	//
	//	(wIpProtocol == 6) || (wIpProtocol == 17)
	//
	// LocalPorts.Ports.dwNumEntries == 0
	//
	// LocalPorts.wPortKeywords == FW_PORT_KEYWORD_NONE
	ProfileConfigAuthAppsAllowUserPrefMerge ProfileConfig = 11
	// FW_PROFILE_CONFIG_GLOBAL_PORTS_ALLOW_USER_PREF_MERGE: This value is a DWORD used
	// as an on/off switch. If this value is off, global port firewall rules in the local
	// store are ignored and not enforced. The setting only has meaning if it is set or
	// enumerated in the Group Policy store or if it is enumerated from the GroupPolicyRSoPStore.
	// The merge law for this option is to let the value GroupPolicyRSoPStore win if it
	// is configured; otherwise, the local store value is used. This symbolic constant has
	// a value of 12.
	//
	// The global port firewall rules consist of the FW_RULE objects where all of the following
	// are true:
	//
	// wszLocalApplication == NULL
	//
	// wszLocalService == NULL
	//
	//	(wIpProtocol == 6) || (wIpProtocol == 17)
	//
	// LocalPorts.Ports.dwNumEntries == 1
	//
	// LocalPorts.wPortKeywords == FW_PORT_KEYWORD_NONE
	ProfileConfigGlobalPortsAllowUserPrefMerge ProfileConfig = 12
	// FW_PROFILE_CONFIG_ALLOW_LOCAL_POLICY_MERGE:  This value is a DWORD used as an on/off
	// switch. If this value is off, firewall rules from the local store are ignored and
	// not enforced. The merge law for this option is to always use the value of the GroupPolicyRSoPStore.
	// This value is valid for all schema versions. This symbolic constant has a value of
	// 13.
	ProfileConfigAllowLocalPolicyMerge ProfileConfig = 13
	// FW_PROFILE_CONFIG_ALLOW_LOCAL_IPSEC_POLICY_MERGE: This value is a DWORD; it is an
	// on/off switch. If this value is off, connection security rules from the local store
	// are ignored and not enforced, regardless of the schema version and connection security
	// rule version. The merge law for this option is to always use the value of the GroupPolicyRSoPStore.
	// This symbolic constant has a value of 14.
	ProfileConfigAllowLocalIPsecPolicyMerge ProfileConfig = 14
	// FW_PROFILE_CONFIG_DISABLED_INTERFACES:  This value is an FW_INTERFACE_LUIDS structure
	// that represents the network adapters where the firewall (only the firewall rules
	// and actions) is off. The merge law for this option is to let the value of the GroupPolicyRSoPStore
	// win if it is configured; otherwise, the local store value is used. This symbolic
	// constant has a value of 15.
	ProfileConfigDisabledInterfaces ProfileConfig = 15
	// FW_PROFILE_CONFIG_DEFAULT_OUTBOUND_ACTION:  This value is the action that the firewall
	// does by default (and evaluates at the very end) on outbound connections. The allow
	// action is represented by 0x00000000; 0x00000001 represents a block action. The merge
	// law for this option is to let the value of the GroupPolicyRSoPStore win if it is
	// configured; otherwise, the local store value is used. This symbolic constant has
	// a value of 16.
	ProfileConfigDefaultOutboundAction ProfileConfig = 16
	// FW_PROFILE_CONFIG_DEFAULT_INBOUND_ACTION:  This value is the action that the firewall
	// does by default (and evaluates at the very end) on inbound connections. The allow
	// action is represented by 0x00000000; 0x00000001 represents a block action. The merge
	// law for this option is to let the value of the GroupPolicyRSoPStore.win if it is
	// configured; otherwise, the local store value is used. This symbolic constant has
	// a value of 17.
	ProfileConfigDefaultInboundAction ProfileConfig = 17
	// FW_PROFILE_CONFIG_DISABLE_STEALTH_MODE_IPSEC_SECURED_PACKET_EXEMPTION: This value
	// is a DWORD used as an on/off switch. This option is ignored if FW_PROFILE_CONFIG_DISABLE_STEALTH_MODE
	// is on. Otherwise, when this option is on, the firewall's stealth mode rules MUST
	// NOT prevent the host computer from responding to unsolicited network traffic if that
	// traffic is secured by IPsec. The merge law for this option is to let the value of
	// the GroupPolicyRSoPStore win if it is configured; otherwise, the local store value
	// is used. For schema versions 0x0200, 0x0201, and 0x020A, this value is invalid and
	// MUST NOT be used. This symbolic constant has a value of 18.
	ProfileConfigDisableStealthModeIPsecSecuredPacketExemption ProfileConfig = 18
	// FW_PROFILE_CONFIG_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 19.
	ProfileConfigMax ProfileConfig = 19
)

func (o ProfileConfig) String() string {
	switch o {
	case ProfileConfigInvalid:
		return "ProfileConfigInvalid"
	case ProfileConfigEnableFw:
		return "ProfileConfigEnableFw"
	case ProfileConfigDisableStealthMode:
		return "ProfileConfigDisableStealthMode"
	case ProfileConfigShielded:
		return "ProfileConfigShielded"
	case ProfileConfigDisableUnicastResponsesToMulticastBroadcast:
		return "ProfileConfigDisableUnicastResponsesToMulticastBroadcast"
	case ProfileConfigLogDroppedPackets:
		return "ProfileConfigLogDroppedPackets"
	case ProfileConfigLogSuccessConnections:
		return "ProfileConfigLogSuccessConnections"
	case ProfileConfigLogIgnoredRules:
		return "ProfileConfigLogIgnoredRules"
	case ProfileConfigLogMaxFileSize:
		return "ProfileConfigLogMaxFileSize"
	case ProfileConfigLogFilePath:
		return "ProfileConfigLogFilePath"
	case ProfileConfigDisableInboundNotifications:
		return "ProfileConfigDisableInboundNotifications"
	case ProfileConfigAuthAppsAllowUserPrefMerge:
		return "ProfileConfigAuthAppsAllowUserPrefMerge"
	case ProfileConfigGlobalPortsAllowUserPrefMerge:
		return "ProfileConfigGlobalPortsAllowUserPrefMerge"
	case ProfileConfigAllowLocalPolicyMerge:
		return "ProfileConfigAllowLocalPolicyMerge"
	case ProfileConfigAllowLocalIPsecPolicyMerge:
		return "ProfileConfigAllowLocalIPsecPolicyMerge"
	case ProfileConfigDisabledInterfaces:
		return "ProfileConfigDisabledInterfaces"
	case ProfileConfigDefaultOutboundAction:
		return "ProfileConfigDefaultOutboundAction"
	case ProfileConfigDefaultInboundAction:
		return "ProfileConfigDefaultInboundAction"
	case ProfileConfigDisableStealthModeIPsecSecuredPacketExemption:
		return "ProfileConfigDisableStealthModeIPsecSecuredPacketExemption"
	case ProfileConfigMax:
		return "ProfileConfigMax"
	}
	return "Invalid"
}

// GlobalConfigIPsecExemptValues type represents FW_GLOBAL_CONFIG_IPSEC_EXEMPT_VALUES RPC enumeration.
//
// This enumeration identifies specific traffic to be exempted from performing IPsec.
type GlobalConfigIPsecExemptValues uint16

var (
	// FW_GLOBAL_CONFIG_IPSEC_EXEMPT_NONE:  No IPsec exemptions.
	GlobalConfigIPsecExemptValuesNone GlobalConfigIPsecExemptValues = 0
	// FW_GLOBAL_CONFIG_IPSEC_EXEMPT_NEIGHBOR_DISC: Exempt neighbor discover IPv6 ICMP type-codes
	// from IPsec.
	GlobalConfigIPsecExemptValuesNeighborDisc GlobalConfigIPsecExemptValues = 1
	// FW_GLOBAL_CONFIG_IPSEC_EXEMPT_ICMP:  Exempt ICMP from IPsec.
	GlobalConfigIPsecExemptValuesICMP GlobalConfigIPsecExemptValues = 2
	// FW_GLOBAL_CONFIG_IPSEC_EXEMPT_ROUTER_DISC: Exempt router discover IPv6 ICMP type-codes
	// from IPsec.
	GlobalConfigIPsecExemptValuesRouterDisc      GlobalConfigIPsecExemptValues = 4
	GlobalConfigIPsecExemptValuesNeighborDiscRFC GlobalConfigIPsecExemptValues = 5
	// FW_GLOBAL_CONFIG_IPSEC_EXEMPT_DHCP:  Exempt both IPv4 and IPv6 DHCP traffic from
	// IPsec.
	GlobalConfigIPsecExemptValuesDHCP GlobalConfigIPsecExemptValues = 8
	// FW_GLOBAL_CONFIG_IPSEC_EXEMPT_MAX:  This value and values that exceed this value
	// are not valid and MUST NOT be used. It is defined for simplicity in writing IDL definitions
	// and code. This symbolic constant has a value of 0x0010.
	GlobalConfigIPsecExemptValuesMax GlobalConfigIPsecExemptValues = 16
)

func (o GlobalConfigIPsecExemptValues) String() string {
	switch o {
	case GlobalConfigIPsecExemptValuesNone:
		return "GlobalConfigIPsecExemptValuesNone"
	case GlobalConfigIPsecExemptValuesNeighborDisc:
		return "GlobalConfigIPsecExemptValuesNeighborDisc"
	case GlobalConfigIPsecExemptValuesICMP:
		return "GlobalConfigIPsecExemptValuesICMP"
	case GlobalConfigIPsecExemptValuesRouterDisc:
		return "GlobalConfigIPsecExemptValuesRouterDisc"
	case GlobalConfigIPsecExemptValuesNeighborDiscRFC:
		return "GlobalConfigIPsecExemptValuesNeighborDiscRFC"
	case GlobalConfigIPsecExemptValuesDHCP:
		return "GlobalConfigIPsecExemptValuesDHCP"
	case GlobalConfigIPsecExemptValuesMax:
		return "GlobalConfigIPsecExemptValuesMax"
	}
	return "Invalid"
}

// GlobalConfigPresharedKeyEncodingValues type represents FW_GLOBAL_CONFIG_PRESHARED_KEY_ENCODING_VALUES RPC enumeration.
//
// This enumeration is used to describe how preshared keys are encoded before being
// used.
type GlobalConfigPresharedKeyEncodingValues uint16

var (
	// FW_GLOBAL_CONFIG_PRESHARED_KEY_ENCODING_NONE:  Preshared key is not encoded. Instead,
	// it is kept in its wide-character format. This symbolic constant has a value of 0.
	GlobalConfigPresharedKeyEncodingValuesNone GlobalConfigPresharedKeyEncodingValues = 0
	// FW_GLOBAL_CONFIG_PRESHARED_KEY_ENCODING_UTF_8:  Encode the preshared key using UTF-8.
	// This symbolic constant has a value of 1.
	GlobalConfigPresharedKeyEncodingValuesUTF8 GlobalConfigPresharedKeyEncodingValues = 1
	// FW_GLOBAL_CONFIG_PRESHARED_KEY_ENCODING_MAX:  This value and values that exceed
	// this value are not valid and MUST NOT be used. It is defined for simplicity in writing
	// IDL definitions and code. This symbolic constant has a value of 2.
	GlobalConfigPresharedKeyEncodingValuesMax GlobalConfigPresharedKeyEncodingValues = 2
)

func (o GlobalConfigPresharedKeyEncodingValues) String() string {
	switch o {
	case GlobalConfigPresharedKeyEncodingValuesNone:
		return "GlobalConfigPresharedKeyEncodingValuesNone"
	case GlobalConfigPresharedKeyEncodingValuesUTF8:
		return "GlobalConfigPresharedKeyEncodingValuesUTF8"
	case GlobalConfigPresharedKeyEncodingValuesMax:
		return "GlobalConfigPresharedKeyEncodingValuesMax"
	}
	return "Invalid"
}

// GlobalConfigIPsecThroughNATValues type represents FW_GLOBAL_CONFIG_IPSEC_THROUGH_NAT_VALUES RPC enumeration.
//
// This enumeration is used to describe when IPsec security associations can be established
// across NAT devices.
type GlobalConfigIPsecThroughNATValues uint16

var (
	// FW_GLOBAL_CONFIG_IPSEC_THROUGH_NAT_NEVER:  IPsec does not cross NAT boundaries.
	// This symbolic constant has a value of 0.
	GlobalConfigIPsecThroughNATValuesNever GlobalConfigIPsecThroughNATValues = 0
	// FW_GLOBAL_CONFIG_IPSEC_THROUGH_NAT_SERVER_BEHIND_NAT: IPsec security associations
	// can be established when the server is across NAT boundaries. This symbolic constant
	// has a value of 1.
	GlobalConfigIPsecThroughNATValuesServerBehindNAT GlobalConfigIPsecThroughNATValues = 1
	// FW_GLOBAL_CONFIG_IPSEC_THROUGH_NAT_SERVER_AND_CLIENT_BEHIND_NAT: IPsec security associations
	// can be established when the server and client are across NAT boundaries. This symbolic
	// constant has a value of 2.
	GlobalConfigIPsecThroughNATValuesServerAndClientBehindNAT GlobalConfigIPsecThroughNATValues = 2
	// FW_GLOBAL_CONFIG_IPSEC_THROUGH_NAT_MAX:  This value and values that exceed this
	// value are not valid and MUST NOT be used. It is defined for simplicity in writing
	// IDL definitions and code. This symbolic constant has a value of 3.
	GlobalConfigIPsecThroughNATValuesMax GlobalConfigIPsecThroughNATValues = 3
)

func (o GlobalConfigIPsecThroughNATValues) String() string {
	switch o {
	case GlobalConfigIPsecThroughNATValuesNever:
		return "GlobalConfigIPsecThroughNATValuesNever"
	case GlobalConfigIPsecThroughNATValuesServerBehindNAT:
		return "GlobalConfigIPsecThroughNATValuesServerBehindNAT"
	case GlobalConfigIPsecThroughNATValuesServerAndClientBehindNAT:
		return "GlobalConfigIPsecThroughNATValuesServerAndClientBehindNAT"
	case GlobalConfigIPsecThroughNATValuesMax:
		return "GlobalConfigIPsecThroughNATValuesMax"
	}
	return "Invalid"
}

// GlobalConfigEnablePacketQueueFlags type represents FW_GLOBAL_CONFIG_ENABLE_PACKET_QUEUE_FLAGS RPC enumeration.
type GlobalConfigEnablePacketQueueFlags uint16

var (
	GlobalConfigEnablePacketQueueFlagsPacketQueueNone    GlobalConfigEnablePacketQueueFlags = 0
	GlobalConfigEnablePacketQueueFlagsPacketQueueInbound GlobalConfigEnablePacketQueueFlags = 1
	GlobalConfigEnablePacketQueueFlagsPacketQueueForward GlobalConfigEnablePacketQueueFlags = 2
	GlobalConfigEnablePacketQueueFlagsPacketQueueMax     GlobalConfigEnablePacketQueueFlags = 3
)

func (o GlobalConfigEnablePacketQueueFlags) String() string {
	switch o {
	case GlobalConfigEnablePacketQueueFlagsPacketQueueNone:
		return "GlobalConfigEnablePacketQueueFlagsPacketQueueNone"
	case GlobalConfigEnablePacketQueueFlagsPacketQueueInbound:
		return "GlobalConfigEnablePacketQueueFlagsPacketQueueInbound"
	case GlobalConfigEnablePacketQueueFlagsPacketQueueForward:
		return "GlobalConfigEnablePacketQueueFlagsPacketQueueForward"
	case GlobalConfigEnablePacketQueueFlagsPacketQueueMax:
		return "GlobalConfigEnablePacketQueueFlagsPacketQueueMax"
	}
	return "Invalid"
}

// GlobalConfig type represents FW_GLOBAL_CONFIG RPC enumeration.
//
// This enumeration identifies the global policy configuration options. Each configuration
// option has a merge law that is used to determine how to merge the values of these
// options across stores.
type GlobalConfig uint16

var (
	// FW_GLOBAL_CONFIG_INVALID:  This value MUST NOT be used. It is defined for simplicity
	// in writing IDL definitions and code. This symbolic constant has a value of 0.
	GlobalConfigInvalid GlobalConfig = 0
	// FW_GLOBAL_CONFIG_POLICY_VERSION_SUPPORTED:  This value is a DWORD containing the
	// maximum policy version that the server host can accept. The version number is two
	// octets in size. The lowest-order octet is the minor version; the second-to-lowest
	// octet is the major version. This value is not merged and is always a fixed value
	// for a particular firewall and advanced security components software build. This symbolic
	// constant has a value of 1.<12>
	//
	// In section 2 structures and section 3 methods of this document, some section titles
	// are appended with a schema version to which the particular structure or method applies.
	// For example, the appended characters “2_20” in the “FW_RULE2_20” structure
	// section name indicates that this structure applies to and corresponds with policy
	// version 0x0214, as described just ahead.
	//
	// Policy version numbers vary depending on the version of this protocol to which they
	// apply. They also approximately correspond with regular product updates. Policy versions
	// in this protocol are expressed in hexadecimal notation that presently consist of
	// the following versions: 0x0200, 0x0201, 0x020A, 0x0214, 0x0216, 0x0218, 0x0219, 0x021A,
	// 0x021B, 0x021C, 0x021D, 0x021E, 0x021F, and 0x0220. These policy versions are mapped
	// to specific release versions of this protocol.
	GlobalConfigPolicyVersionSupported GlobalConfig = 1
	// FW_GLOBAL_CONFIG_CURRENT_PROFILE:  This value is a DWORD and contains a bitmask
	// of the current enforced profiles that are maintained by the server firewall host.
	// See FW_PROFILE_TYPE (section 2.2.2) for the bitmasks that are used to identify profile
	// types. This value is available only in the dynamic store; therefore, it is not merged
	// and has no merge law. This symbolic constant has a value of 2.
	GlobalConfigCurrentProfile GlobalConfig = 2
	// FW_GLOBAL_CONFIG_DISABLE_STATEFUL_FTP:  This value is an on/off switch. If off,
	// the firewall performs stateful File Transfer Protocol (FTP) filtering to allow secondary
	// connections. The value is a DWORD; 0x00000000 means off; 0x00000001 means on. The
	// merge law for this option is to let "on" values win. This symbolic constant has a
	// value of 3.
	GlobalConfigDisableStatefulFTP GlobalConfig = 3
	// FW_GLOBAL_CONFIG_DISABLE_STATEFUL_PPTP:  This value is an on/off switch. If off,
	// the firewall performs stateful Point-to-Point Tunneling Protocol (PPTP) analysis.
	// The value is a DWORD; 0x00000000 means off; 0x00000001 means on. The merge law for
	// this option is to let "on" values win. This symbolic constant has a value of 4.
	GlobalConfigDisableStatefulPPTP GlobalConfig = 4
	// FW_GLOBAL_CONFIG_SA_IDLE_TIME:  This value configures the security association idle
	// time, in seconds. Security associations are deleted after network traffic is not
	// seen for this specified period of time. The value is a DWORD and MUST be a value
	// in the range of 300 to 3,600 inclusive. The merge law for this option is to let the
	// value of the GroupPolicyRSoPStore win if it is configured; otherwise, use the local
	// store value. This symbolic constant has a value of 5.
	GlobalConfigSAIdleTime GlobalConfig = 5
	// FW_GLOBAL_CONFIG_PRESHARED_KEY_ENCODING:  This configuration value specifies the
	// preshared key encoding that is used. The value is a DWORD and MUST be a valid value
	// from the FW_GLOBAL_CONFIG_PRESHARED_KEY_ENCODING_VALUES enumeration. The merge law
	// for this option is to let the value of the GroupPolicyRSoPStore win if it is configured;
	// otherwise, use the local store value. This symbolic constant has a value of 6.
	GlobalConfigPresharedKeyEncoding GlobalConfig = 6
	// FW_GLOBAL_CONFIG_IPSEC_EXEMPT:  This configuration value configures IPsec exceptions.
	// The value is a DWORD and MUST be a combination of the valid flags that are defined
	// in FW_GLOBAL_CONFIG_IPSEC_EXEMPT_VALUES; therefore, the maximum value MUST always
	// be FW_GLOBAL_CONFIG_IPSEC_EXEMPT_MAX-1 for servers supporting a schema version of
	// 0x0201 and FW_GLOBAL_CONFIG_IPSEC_EXEMPT_MAX_V2_0-1 for servers supporting a schema
	// version of 0x0200. If the maximum value is exceeded when the method RRPC_FWSetGlobalConfig
	// (Opnum 4) is called, the method returns ERROR_INVALID_PARAMETER. This error code
	// is returned if no other preceding error is discovered. The merge law for this option
	// is to let the value of the GroupPolicyRSoPStore win if it is configured; otherwise,
	// use the local store value. This symbolic constant has a value of 7.
	GlobalConfigIPsecExempt GlobalConfig = 7
	// FW_GLOBAL_CONFIG_CRL_CHECK:  This value specifies how certificate revocation list
	// (CRL) verification is enforced. The value is a DWORD and MUST be 0, 1, or 2. A value
	// of 0 disables CRL checking. A value of 1 specifies that CRL checking is attempted
	// and that certificate validation fails only if the certificate is revoked. Other failures
	// that are encountered during CRL checking (such as the revocation URL being unreachable)
	// do not cause certificate validation to fail. A value of 2 means that checking is
	// required and that certificate validation fails if any error is encountered during
	// CRL processing. The merge law for this option is to let the value of the GroupPolicyRSoPStore
	// win if it is configured; otherwise, use the local store value. This symbolic constant
	// has a value of 8.
	GlobalConfigCRLCheck GlobalConfig = 8
	// FW_GLOBAL_CONFIG_IPSEC_THROUGH_NAT:  This value is configured when an IPsec security
	// association can be established with a computer across NAT devices. The value is of
	// type FW_GLOBAL_CONFIG_IPSEC_THROUGH_NAT_VALUES and MUST contain valid values of the
	// same enumeration type. The merge law for this option is to let the value of the GroupPolicyRSoPStore
	// win if it is configured; otherwise, use the local store value. This symbolic constant
	// has a value of 9.
	GlobalConfigIPsecThroughNAT GlobalConfig = 9
	// FW_GLOBAL_CONFIG_POLICY_VERSION:  This value contains the policy version of the
	// policy store being managed. This value is not merged and therefore, has no merge
	// law. This symbolic constant has a value of 10.
	GlobalConfigPolicyVersion GlobalConfig = 10
	// FW_GLOBAL_CONFIG_BINARY_VERSION_SUPPORTED:  This value contains the binary version
	// of the structures and data types that are supported by the server. This value is
	// not merged. In addition, this value is always a fixed value for a specific firewall
	// and advanced security component's software build. This symbolic constant has a value
	// of 11. This value identifies a policy configuration option that is supported only
	// on servers that have a schema version of 0x0201.
	GlobalConfigBinaryVersionSupported GlobalConfig = 11
	// FW_GLOBAL_CONFIG_IPSEC_TUNNEL_REMOTE_MACHINE_AUTHORIZATION_LIST: This value represents
	// a list of remote machines that are allowed to send and receive traffic through the
	// tunnels which request this access check. Machines in the list are allowed through
	// the tunnels. Machines not in the list are denied through the tunnels. The list is
	// specified as a security descriptor which specifies which SIDs ([MS-DTYP] section
	// 2.4.2.1) of the remote machines. The value is a Unicode string in Security Descriptor
	// Definition Language (SDDL) format ([MS-DTYP] section 2.5.1). This symbolic constant
	// has a value of 12.
	GlobalConfigIPsecTunnelRemoteMachineAuthorizationList GlobalConfig = 12
	// FW_GLOBAL_CONFIG_IPSEC_TUNNEL_REMOTE_USER_AUTHORIZATION_LIST: This value represents
	// a list of remote users who are allowed to send and receive traffic through the tunnels
	// which request this access check. Users in the list are allowed through the tunnels.
	// Users not in the list are denied through the tunnels. The list is specified as a
	// security descriptor which specifies which SIDs ([MS-DTYP] section 2.4.2.1) of the
	// remote users. The value is a Unicode string in SDDL format ([MS-DTYP] section 2.5.1).
	// This symbolic constant has a value of 13.
	GlobalConfigIPsecTunnelRemoteUserAuthorizationList GlobalConfig = 13
	// FW_GLOBAL_CONFIG_OPPORTUNISTICALLY_MATCH_AUTH_SET_PER_KM: This value is a DWORD used
	// as an on/off switch. When this option is off, keying modules MUST ignore the entire
	// authentication set if they do not support all of the authentication suites specified
	// in the set. When this option is on, keying modules MUST ignore only the authentication
	// suites that they don’t support. For schema versions 0x0200, 0x0201, and 0x020A,
	// this value is invalid and MUST NOT be used. This symbolic constant has a value of
	// 14.
	GlobalConfigOpportunisticallyMatchAuthSetPerKM GlobalConfig = 14
	// FW_GLOBAL_CONFIG_IPSEC_TRANSPORT_REMOTE_MACHINE_AUTHORIZATION_LIST: This value is
	// a Unicode string in Security Descriptor Definition Language (SDDL) format ([MS-DTYP]
	// section 2.5.1). The security descriptor describes which remote machines are allowed
	// to send and receive traffic secured by transport mode connection security rules which
	// request this access check. Machines granted access by the security descriptor are
	// allowed to send and receive traffic. Machines denied access by the security descriptor
	// are blocked from sending and receiving traffic. For schema versions 0x0200, 0x0201,
	// and 0x020A, this value is invalid and MUST NOT be used. This symbolic constant has
	// a value of 15.
	GlobalConfigIPsecTransportRemoteMachineAuthorizationList GlobalConfig = 15
	// FW_GLOBAL_CONFIG_IPSEC_TRANSPORT_REMOTE_USER_AUTHORIZATION_LIST: This value is a
	// Unicode string in Security Descriptor Definition Language (SDDL) format. The security
	// descriptor describes which remote users are allowed to send and receive traffic secured
	// by transport mode connection security rules which request this access check. Users
	// granted access by the security descriptor are allowed to send and receive traffic.
	// Users denied access by the security descriptor are blocked from sending and receiving
	// traffic. For schema versions 0x0200, 0x0201, and 0x020A, this value is invalid and
	// MUST NOT be used. This symbolic constant has a value of 16.
	GlobalConfigIPsecTransportRemoteUserAuthorizationList GlobalConfig = 16
	// FW_GLOBAL_CONFIG_ENABLE_PACKET_QUEUE:  This value specifies how scaling for the
	// software on the receive side is enabled for both the encrypted receive and clear
	// text forward path for the IPsec tunnel gateway scenario (as configured by FW_CS_RULE
	// (section 2.2.55)). Use of this option also ensures that the packet order is preserved.
	// The data type for this option value is a DWORD and is a combination of flags. A value
	// of 0x00 indicates that all queuing is to be disabled. A value of 0x01 specifies that
	// inbound encrypted packets are to be queued. A value of 0x02 specifies that packets
	// are to be queued after decryption is performed for forwarding. This symbolic constant
	// has a value of 17.
	GlobalConfigEnablePacketQueue GlobalConfig = 17
	// FW_GLOBAL_CONFIG_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. This symbolic constant is defined for simplicity in writing
	// IDL definitions and code. It has a value of 18.
	GlobalConfigMax GlobalConfig = 18
)

func (o GlobalConfig) String() string {
	switch o {
	case GlobalConfigInvalid:
		return "GlobalConfigInvalid"
	case GlobalConfigPolicyVersionSupported:
		return "GlobalConfigPolicyVersionSupported"
	case GlobalConfigCurrentProfile:
		return "GlobalConfigCurrentProfile"
	case GlobalConfigDisableStatefulFTP:
		return "GlobalConfigDisableStatefulFTP"
	case GlobalConfigDisableStatefulPPTP:
		return "GlobalConfigDisableStatefulPPTP"
	case GlobalConfigSAIdleTime:
		return "GlobalConfigSAIdleTime"
	case GlobalConfigPresharedKeyEncoding:
		return "GlobalConfigPresharedKeyEncoding"
	case GlobalConfigIPsecExempt:
		return "GlobalConfigIPsecExempt"
	case GlobalConfigCRLCheck:
		return "GlobalConfigCRLCheck"
	case GlobalConfigIPsecThroughNAT:
		return "GlobalConfigIPsecThroughNAT"
	case GlobalConfigPolicyVersion:
		return "GlobalConfigPolicyVersion"
	case GlobalConfigBinaryVersionSupported:
		return "GlobalConfigBinaryVersionSupported"
	case GlobalConfigIPsecTunnelRemoteMachineAuthorizationList:
		return "GlobalConfigIPsecTunnelRemoteMachineAuthorizationList"
	case GlobalConfigIPsecTunnelRemoteUserAuthorizationList:
		return "GlobalConfigIPsecTunnelRemoteUserAuthorizationList"
	case GlobalConfigOpportunisticallyMatchAuthSetPerKM:
		return "GlobalConfigOpportunisticallyMatchAuthSetPerKM"
	case GlobalConfigIPsecTransportRemoteMachineAuthorizationList:
		return "GlobalConfigIPsecTransportRemoteMachineAuthorizationList"
	case GlobalConfigIPsecTransportRemoteUserAuthorizationList:
		return "GlobalConfigIPsecTransportRemoteUserAuthorizationList"
	case GlobalConfigEnablePacketQueue:
		return "GlobalConfigEnablePacketQueue"
	case GlobalConfigMax:
		return "GlobalConfigMax"
	}
	return "Invalid"
}

// ConfigFlags type represents FW_CONFIG_FLAGS RPC enumeration.
//
// This enumeration identifies flags that can be set on the RRPC_FWGetConfig (Opnum
// 10) and RRPC_FWGetGlobalConfig (Opnum 3) methods.
type ConfigFlags uint16

var (
	// FW_CONFIG_FLAG_RETURN_DEFAULT_IF_NOT_FOUND:  If this flag is specified, and if the
	// RRPC_FWGetConfig (Opnum 10) method or the RRPC_FWGetGlobalConfig (Opnum 3) method
	// fails to find the configuration value in the policy store, then the call will succeed
	// and return the default value used by the firewall service. If this flag is not specified,
	// these methods will fail with ERROR_FILE_NOT_FOUND. The default set of values returned
	// by these two calls is a firewall and advanced security component implementation-specific<13>
	// decision, and is outside the scope of this protocol specification.
	ConfigFlagsReturnDefaultInterfaceNotFound ConfigFlags = 1
)

func (o ConfigFlags) String() string {
	switch o {
	case ConfigFlagsReturnDefaultInterfaceNotFound:
		return "ConfigFlagsReturnDefaultInterfaceNotFound"
	}
	return "Invalid"
}

// Network structure represents FW_NETWORK RPC structure.
//
// This structure represents a network that is associated with a firewall profile. It
// is used for display purposes in user interfaces.
type Network struct {
	// pszName:  A pointer to a Unicode string that represents the name of the network.
	Name string `idl:"name:pszName;string;pointer:unique" json:"name"`
	// ProfileType:  The profile type that is associated with the network. The type MUST
	// be one of the FW_PROFILE_TYPE flags, except FW_PROFILE_TYPE_ALL.
	ProfileType ProfileType `idl:"name:ProfileType" json:"profile_type"`
}

func (o *Network) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Network) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_pszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.ProfileType)); err != nil {
		return err
	}
	return nil
}
func (o *Network) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_pszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_pszName, _ptr_pszName); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.ProfileType)); err != nil {
		return err
	}
	return nil
}

// Adapter structure represents FW_ADAPTER RPC structure.
//
// This structure represents a network interface in the host. It is used for display
// purposes in the user interface when configuring the FW_PROFILE_CONFIG_DISABLED_INTERFACES
// (section 2.2.38) configuration option.
type Adapter struct {
	// pszFriendlyName:  A pointer to a Unicode string that presents the friendly name that
	// is associated with the network interface.
	FriendlyName string `idl:"name:pszFriendlyName;string;pointer:unique" json:"friendly_name"`
	// Guid:  A GUID that uniquely identifies the interface in the host system.
	GUID *dtyp.GUID `idl:"name:Guid" json:"guid"`
}

func (o *Adapter) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Adapter) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.FriendlyName != "" {
		_ptr_pszFriendlyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FriendlyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FriendlyName, _ptr_pszFriendlyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Adapter) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pszFriendlyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FriendlyName); err != nil {
			return err
		}
		return nil
	})
	_s_pszFriendlyName := func(ptr interface{}) { o.FriendlyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FriendlyName, _s_pszFriendlyName, _ptr_pszFriendlyName); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DiagApp structure represents FW_DIAG_APP RPC structure.
//
// This structure is not used on the wire.
type DiagApp struct {
	AppPath string `idl:"name:pszAppPath;string;pointer:unique" json:"app_path"`
}

func (o *DiagApp) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiagApp) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.AppPath != "" {
		_ptr_pszAppPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.AppPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.AppPath, _ptr_pszAppPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DiagApp) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_pszAppPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.AppPath); err != nil {
			return err
		}
		return nil
	})
	_s_pszAppPath := func(ptr interface{}) { o.AppPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.AppPath, _s_pszAppPath, _ptr_pszAppPath); err != nil {
		return err
	}
	return nil
}

// RuleCategory type represents FW_RULE_CATEGORY RPC enumeration.
//
// This enumeration represents the classes of functionality that a third-party software
// component can register for, take ownership of, and commit to implement. The implementation
// of such functionality by the firewall and advanced security component, or by the
// third-party software component, are implementation-specific decisions. This enumeration
// is only used to present the state of the registrations.
type RuleCategory uint32

var (
	// FW_RULE_CATEGORY_BOOT:  This category of functionality represents the policy that
	// is used while the system is starting up and the firewall and advance security component
	// is not yet running. This symbolic constant has a value of 0.
	RuleCategoryBoot RuleCategory = 0
	// FW_RULE_CATEGORY_STEALTH:  This category of functionality represents the policy
	// that is used to make the system appear invisible when it is connected to a network.
	// For example, this functionality helps prevent attackers from discovering the host
	// and the ports that open to the host. This symbolic constant has a value of 1.
	RuleCategoryStealth RuleCategory = 1
	// FW_RULE_CATEGORY_FIREWALL:  This category of functionality represents functions
	// that are performed by firewall objects while they are present on the FW_STORE_TYPE_LOCAL,
	// FW_STORE_TYPE_DYNAMIC, and FW_STORE_TYPE_GP_RSOP policy stores (see section 2.2.1).
	// This symbolic constant has a value of 2.
	RuleCategoryFirewall RuleCategory = 2
	// FW_RULE_CATEGORY_CONSEC:  This category of functionality represents functions that
	// are performed by the connection security objects. This symbolic constant has a value
	// of 3.
	RuleCategoryConsec RuleCategory = 3
	// FW_RULE_CATEGORY_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 4.
	RuleCategoryMax RuleCategory = 4
)

func (o RuleCategory) String() string {
	switch o {
	case RuleCategoryBoot:
		return "RuleCategoryBoot"
	case RuleCategoryStealth:
		return "RuleCategoryStealth"
	case RuleCategoryFirewall:
		return "RuleCategoryFirewall"
	case RuleCategoryConsec:
		return "RuleCategoryConsec"
	case RuleCategoryMax:
		return "RuleCategoryMax"
	}
	return "Invalid"
}

// Product structure represents FW_PRODUCT RPC structure.
//
// This structure represents a third-party software component that registers with the
// firewall and advanced security component to implement some of the categories.
type Product struct {
	// dwFlags:  This field is not used.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// dwNumRuleCategories:  The number of rule categories with which the third-party software
	// component registered.
	RuleCategoriesLength uint32 `idl:"name:dwNumRuleCategories" json:"rule_categories_length"`
	// pRuleCategories:  A pointer to an array of dwNumRuleCategories that are contiguous
	// FW_RULE_CATEGORY elements.
	RuleCategories []RuleCategory `idl:"name:pRuleCategories;size_is:(dwNumRuleCategories);pointer:unique" json:"rule_categories"`
	// pszDisplayName:  A pointer to a Unicode string. The string represents the name of
	// the third-party software component.
	DisplayName string `idl:"name:pszDisplayName;string;pointer:ref" json:"display_name"`
	// pszPathToSignedProductExe:  A pointer to a Unicode string. The string represents
	// the file path to the binary executable of the third-party software component.
	PathToSignedProductExe string `idl:"name:pszPathToSignedProductExe;string;pointer:unique" json:"path_to_signed_product_exe"`
}

func (o *Product) xxx_PreparePayload(ctx context.Context) error {
	if o.RuleCategories != nil && o.RuleCategoriesLength == 0 {
		o.RuleCategoriesLength = uint32(len(o.RuleCategories))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Product) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.RuleCategoriesLength); err != nil {
		return err
	}
	if o.RuleCategories != nil || o.RuleCategoriesLength > 0 {
		_ptr_pRuleCategories := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RuleCategoriesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.RuleCategories {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(uint32(o.RuleCategories[i1])); err != nil {
					return err
				}
			}
			for i1 := len(o.RuleCategories); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleCategories, _ptr_pRuleCategories); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DisplayName != "" {
		_ptr_pszDisplayName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DisplayName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DisplayName, _ptr_pszDisplayName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PathToSignedProductExe != "" {
		_ptr_pszPathToSignedProductExe := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PathToSignedProductExe); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PathToSignedProductExe, _ptr_pszPathToSignedProductExe); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Product) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.RuleCategoriesLength); err != nil {
		return err
	}
	_ptr_pRuleCategories := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RuleCategoriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RuleCategoriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.RuleCategories", sizeInfo[0])
		}
		o.RuleCategories = make([]RuleCategory, sizeInfo[0])
		for i1 := range o.RuleCategories {
			i1 := i1
			if err := w.ReadData((*uint32)(&o.RuleCategories[i1])); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pRuleCategories := func(ptr interface{}) { o.RuleCategories = *ptr.(*[]RuleCategory) }
	if err := w.ReadPointer(&o.RuleCategories, _s_pRuleCategories, _ptr_pRuleCategories); err != nil {
		return err
	}
	_ptr_pszDisplayName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DisplayName); err != nil {
			return err
		}
		return nil
	})
	_s_pszDisplayName := func(ptr interface{}) { o.DisplayName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DisplayName, _s_pszDisplayName, _ptr_pszDisplayName); err != nil {
		return err
	}
	_ptr_pszPathToSignedProductExe := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathToSignedProductExe); err != nil {
			return err
		}
		return nil
	})
	_s_pszPathToSignedProductExe := func(ptr interface{}) { o.PathToSignedProductExe = *ptr.(*string) }
	if err := w.ReadPointer(&o.PathToSignedProductExe, _s_pszPathToSignedProductExe, _ptr_pszPathToSignedProductExe); err != nil {
		return err
	}
	return nil
}

// IPVersion type represents FW_IP_VERSION RPC enumeration.
//
// This enumeration is used to represent the two current IP protocol versions in use:
// IP version 4 and IP version 6.
type IPVersion uint16

var (
	// FW_IP_VERSION_INVALID:  This value MUST NOT be used. It is defined for simplicity
	// in writing IDL definitions and code. This symbolic constant has a value of 0.
	IPVersionInvalid IPVersion = 0
	// FW_IP_VERSION_V4:  This value represents IPv4. This symbolic constant has a value
	// of 1.
	IPVersionV4 IPVersion = 1
	// FW_IP_VERSION_V6:  This value represents the IPv6. This symbolic constant has a
	// value of 2.
	IPVersionV6 IPVersion = 2
	// FW_IP_VERSION_MAX:  This value and values that exceed this value are not valid and
	// MUST NOT be used. It is defined for simplicity in writing IDL definitions and code.
	// This symbolic constant has a value of 3.
	IPVersionMax IPVersion = 3
)

func (o IPVersion) String() string {
	switch o {
	case IPVersionInvalid:
		return "IPVersionInvalid"
	case IPVersionV4:
		return "IPVersionV4"
	case IPVersionV6:
		return "IPVersionV6"
	case IPVersionMax:
		return "IPVersionMax"
	}
	return "Invalid"
}

// IPsecPhase type represents FW_IPSEC_PHASE RPC enumeration.
//
// This enumeration is used to identify the IPsec phase of negotiations.
type IPsecPhase uint16

var (
	// FW_IPSEC_PHASE_INVALID:  This value MUST NOT be used. It is defined for simplicity
	// in writing IDL definitions and code. This symbolic constant has a value of 0.
	IPsecPhaseInvalid IPsecPhase = 0
	// FW_IPSEC_PHASE_1:  This value represents the IPsec first phase of negotiations,
	// also called main mode. This symbolic constant has a value of 1.
	IPsecPhasePhase1 IPsecPhase = 1
	// FW_IPSEC_PHASE_2:  This value represents the IPsec second phase of negotiations.
	// A phase 2 authentication is the second authentication and can mean extended mode
	// or quick mode. On auth sets, phase 2 authentication refers to extended mode (specified
	// in [MS-AIPS] sections 3.6 and 3.7); and on crypto sets, phase 2 refers to quick mode
	// (specified in [MS-AIPS] sections 3.4 and 3.5). This symbolic constant has a value
	// of 2.
	IPsecPhasePhase2 IPsecPhase = 2
	// FW_IPSEC_PHASE_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 3.
	IPsecPhaseMax IPsecPhase = 3
)

func (o IPsecPhase) String() string {
	switch o {
	case IPsecPhaseInvalid:
		return "IPsecPhaseInvalid"
	case IPsecPhasePhase1:
		return "IPsecPhasePhase1"
	case IPsecPhasePhase2:
		return "IPsecPhasePhase2"
	case IPsecPhaseMax:
		return "IPsecPhaseMax"
	}
	return "Invalid"
}

// CSRuleFlags type represents FW_CS_RULE_FLAGS RPC enumeration.
//
// This enumeration describes flag values for connection security rules.
type CSRuleFlags uint16

var (
	// FW_CS_RULE_FLAGS_NONE:  This value means that none of the following flags are set.
	// This value is defined for simplicity in writing IDL definitions and code.
	CSRuleFlagsNone CSRuleFlags = 0
	// FW_CS_RULE_FLAGS_ACTIVE:  If this flag is set, the rule is enabled; otherwise, the
	// rule is disabled.
	CSRuleFlagsActive CSRuleFlags = 1
	// FW_CS_RULE_FLAGS_DTM:  If this flag is set, the rule is a dynamic tunnel mode rule.
	CSRuleFlagsDTM CSRuleFlags = 2
	// FW_CS_RULE_FLAGS_TUNNEL_BYPASS_IF_ENCRYPTED:  This flag MUST only be set on tunnel
	// mode rules. If this flag is set and traffic is already arriving encrypted, it is
	// exempted from the tunnel.
	CSRuleFlagsTunnelBypassInterfaceEncrypted CSRuleFlags = 8
	// FW_CS_RULE_FLAGS_OUTBOUND_CLEAR:  This flag MUST only be set on tunnel mode rules.
	// If set, when outbound traffic matches the rule, it leaves unprotected, but inbound
	// traffic MUST arrive through the tunnel.
	CSRuleFlagsOutboundClear CSRuleFlags = 16
	// FW_CS_RULE_FLAGS_APPLY_AUTHZ:  This flag MUST only be set on tunnel mode rules.
	// If this flag is set, the authenticated peers of the traffic MUST match the SDDLs
	// that are specified in FW_GLOBAL_CONFIG_IPSEC_TUNNEL_REMOTE_MACHINE_AUTHORIZATION_LIST
	// and FW_GLOBAL_CONFIG_IPSEC_TUNNEL_REMOTE_USER_AUTHORIZATION_LIST.
	CSRuleFlagsApplyAuthz CSRuleFlags = 32
	// FW_CS_RULE_FLAGS_KEY_MANAGER_ALLOW_DICTATE_KEY:  If this flag is set, external key
	// managers are permitted to dictate the cryptographic keys used. For schema versions
	// 0x0200, 0x0201, and 0x020A, this value is invalid and MUST NOT be used.
	CSRuleFlagsKeyManagerAllowDictateKey CSRuleFlags = 64
	// FW_CS_RULE_FLAGS_KEY_MANAGER_ALLOW_NOTIFY_KEY:  If this flag is set, external key
	// managers are notified of the cryptographic keys used. For schema versions 0x0200,
	// 0x0201, and 0x020A, this value is invalid and MUST NOT be used.
	CSRuleFlagsKeyManagerAllowNotifyKey CSRuleFlags = 128
	// FW_CS_RULE_FLAGS_SECURITY_REALM: If this flag is set, the connection security rule
	// is associated with a security realm. The wszRuleId of the connection security rule
	// is the same as the IPsec Security Realm ID that it is associated with. For schema
	// versions 0x0200, 0x0201, 0x20A, and 0x0214, this value is invalid and MUST NOT be
	// used.
	CSRuleFlagsSecurityRealm CSRuleFlags = 256
	// FW_CS_RULE_FLAGS_MAX:  This value and values that exceed this value are not valid
	// for all schema versions and MUST NOT be used. It is only defined for simplicity in
	// writing IDL definitions and code. This symbolic constant has a value of 0x200.
	CSRuleFlagsMax    CSRuleFlags = 512
	CSRuleFlagsMaxV21 CSRuleFlags = 2
	CSRuleFlagsMaxV28 CSRuleFlags = 4
	// FW_CS_RULE_FLAGS_MAX_V2_10:  This value and values that exceed this value are not
	// valid and MUST NOT be used by servers and clients with schema version 0x020A and
	// earlier. It is defined for simplicity in writing IDL definitions and code. This symbolic
	// constant has a value of 0x40.
	CSRuleFlagsMaxV210 CSRuleFlags = 64
	// FW_CS_RULE_FLAGS_MAX_V2_20:  This value and values that exceed this value are not
	// valid and MUST NOT be used by servers and clients with schema version 0x0214 and
	// earlier. It is defined for simplicity in writing IDL definitions and code. This symbolic
	// constant has a value of 0x100.
	CSRuleFlagsMaxV220 CSRuleFlags = 256
)

func (o CSRuleFlags) String() string {
	switch o {
	case CSRuleFlagsNone:
		return "CSRuleFlagsNone"
	case CSRuleFlagsActive:
		return "CSRuleFlagsActive"
	case CSRuleFlagsDTM:
		return "CSRuleFlagsDTM"
	case CSRuleFlagsTunnelBypassInterfaceEncrypted:
		return "CSRuleFlagsTunnelBypassInterfaceEncrypted"
	case CSRuleFlagsOutboundClear:
		return "CSRuleFlagsOutboundClear"
	case CSRuleFlagsApplyAuthz:
		return "CSRuleFlagsApplyAuthz"
	case CSRuleFlagsKeyManagerAllowDictateKey:
		return "CSRuleFlagsKeyManagerAllowDictateKey"
	case CSRuleFlagsKeyManagerAllowNotifyKey:
		return "CSRuleFlagsKeyManagerAllowNotifyKey"
	case CSRuleFlagsSecurityRealm:
		return "CSRuleFlagsSecurityRealm"
	case CSRuleFlagsMax:
		return "CSRuleFlagsMax"
	case CSRuleFlagsMaxV21:
		return "CSRuleFlagsMaxV21"
	case CSRuleFlagsMaxV28:
		return "CSRuleFlagsMaxV28"
	case CSRuleFlagsMaxV210:
		return "CSRuleFlagsMaxV210"
	case CSRuleFlagsMaxV220:
		return "CSRuleFlagsMaxV220"
	}
	return "Invalid"
}

// CSRuleAction type represents FW_CS_RULE_ACTION RPC enumeration.
//
// This enumeration identifies the possible actions a connection security rule (section
// 2.2.55) can have.
type CSRuleAction uint16

var (
	// FW_CS_RULE_ACTION_INVALID:  This value MUST NOT be used. It is defined for simplicity
	// in writing IDL definitions and code. This symbolic constant has a value of 0.
	CSRuleActionInvalid CSRuleAction = 0
	// FW_CS_RULE_ACTION_SECURE_SERVER:  This action requires inbound traffic to be IPsec
	// traffic and attempts to secure outbound traffic with IPsec. This symbolic constant
	// has a value of 1.
	CSRuleActionSecureServer CSRuleAction = 1
	// FW_CS_RULE_ACTION_BOUNDARY:  This action attempts to secure inbound and outbound
	// traffic with IPsec. If the action fails to secure the traffic, the traffic still
	// flows on the clear. This symbolic constant has a value of 2.
	CSRuleActionBoundary CSRuleAction = 2
	// FW_CS_RULE_ACTION_SECURE:  This action requires inbound and outbound traffic to
	// be secured by IPsec. This symbolic constant has a value of 3.
	CSRuleActionSecure CSRuleAction = 3
	// FW_CS_RULE_ACTION_DO_NOT_SECURE:  This action exempts the traffic from being secured
	// by IPsec. This symbolic constant has a value of 4.
	CSRuleActionDoNotSecure CSRuleAction = 4
	// FW_CS_RULE_ACTION_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 5.
	CSRuleActionMax CSRuleAction = 5
)

func (o CSRuleAction) String() string {
	switch o {
	case CSRuleActionInvalid:
		return "CSRuleActionInvalid"
	case CSRuleActionSecureServer:
		return "CSRuleActionSecureServer"
	case CSRuleActionBoundary:
		return "CSRuleActionBoundary"
	case CSRuleActionSecure:
		return "CSRuleActionSecure"
	case CSRuleActionDoNotSecure:
		return "CSRuleActionDoNotSecure"
	case CSRuleActionMax:
		return "CSRuleActionMax"
	}
	return "Invalid"
}

// CSRule20 structure represents FW_CS_RULE2_0 RPC structure.
//
// This structure describes a connection security rule that is used by the 2.0 binary
// version for servers and clients (see sections 2.2.42 and 2.2.38). The fields of this
// structure are identical to the FW_CS_RULE structure and their meanings are covered
// in section 2.2.55.
type CSRule20 struct {
	Next                   *CSRule20       `idl:"name:pNext" json:"next"`
	SchemaVersion          uint16          `idl:"name:wSchemaVersion" json:"schema_version"`
	RuleID                 string          `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	Name                   string          `idl:"name:wszName;string" json:"name"`
	Description            string          `idl:"name:wszDescription;string" json:"description"`
	Profiles               uint32          `idl:"name:dwProfiles" json:"profiles"`
	Endpoint1              *Addresses      `idl:"name:Endpoint1" json:"endpoint1"`
	Endpoint2              *Addresses      `idl:"name:Endpoint2" json:"endpoint2"`
	LocalInterfaceIDs      *InterfaceLUIDs `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	LocalInterfaceTypes    uint32          `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	LocalTunnelEndpointV4  uint32          `idl:"name:dwLocalTunnelEndpointV4" json:"local_tunnel_endpoint_v4"`
	LocalTunnelEndpointV6  []byte          `idl:"name:LocalTunnelEndpointV6" json:"local_tunnel_endpoint_v6"`
	RemoteTunnelEndpointV4 uint32          `idl:"name:dwRemoteTunnelEndpointV4" json:"remote_tunnel_endpoint_v4"`
	RemoteTunnelEndpointV6 []byte          `idl:"name:RemoteTunnelEndpointV6" json:"remote_tunnel_endpoint_v6"`
	Endpoint1Ports         *Ports          `idl:"name:Endpoint1Ports" json:"endpoint1_ports"`
	Endpoint2Ports         *Ports          `idl:"name:Endpoint2Ports" json:"endpoint2_ports"`
	IPProtocol             uint16          `idl:"name:wIpProtocol" json:"ip_protocol"`
	Phase1AuthSet          string          `idl:"name:wszPhase1AuthSet;string" json:"phase1_auth_set"`
	Phase2CryptoSet        string          `idl:"name:wszPhase2CryptoSet;string" json:"phase2_crypto_set"`
	Phase2AuthSet          string          `idl:"name:wszPhase2AuthSet;string" json:"phase2_auth_set"`
	Action                 CSRuleAction    `idl:"name:Action" json:"action"`
	Flags                  uint16          `idl:"name:wFlags" json:"flags"`
	EmbeddedContext        string          `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	PlatformValidityList   *OSPlatformList `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	Origin                 RuleOriginType  `idl:"name:Origin" json:"origin"`
	GPOName                string          `idl:"name:wszGPOName;string" json:"gpo_name"`
	Status                 RuleStatus      `idl:"name:Status" json:"status"`
}

func (o *CSRule20) xxx_PreparePayload(ctx context.Context) error {
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.Phase1AuthSet) > int(255) {
		return fmt.Errorf("Phase1AuthSet is out of range")
	}
	if len(o.Phase2CryptoSet) > int(255) {
		return fmt.Errorf("Phase2CryptoSet is out of range")
	}
	if len(o.Phase2AuthSet) > int(255) {
		return fmt.Errorf("Phase2AuthSet is out of range")
	}
	if o.Action < CSRuleAction(1) || o.Action > CSRuleAction(5) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(5) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CSRule20) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CSRule20{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if o.Endpoint1 != nil {
		if err := o.Endpoint1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Endpoint2 != nil {
		if err := o.Endpoint2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalTunnelEndpointV4); err != nil {
		return err
	}
	for i1 := range o.LocalTunnelEndpointV6 {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.LocalTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LocalTunnelEndpointV6); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RemoteTunnelEndpointV4); err != nil {
		return err
	}
	for i1 := range o.RemoteTunnelEndpointV6 {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RemoteTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteTunnelEndpointV6); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if o.Endpoint1Ports != nil {
		if err := o.Endpoint1Ports.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Endpoint2Ports != nil {
		if err := o.Endpoint2Ports.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	if o.Phase1AuthSet != "" {
		_ptr_wszPhase1AuthSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase1AuthSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase1AuthSet, _ptr_wszPhase1AuthSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Phase2CryptoSet != "" {
		_ptr_wszPhase2CryptoSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase2CryptoSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase2CryptoSet, _ptr_wszPhase2CryptoSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Phase2AuthSet != "" {
		_ptr_wszPhase2AuthSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase2AuthSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase2AuthSet, _ptr_wszPhase2AuthSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	return nil
}
func (o *CSRule20) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &CSRule20{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**CSRule20) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if o.Endpoint1 == nil {
		o.Endpoint1 = &Addresses{}
	}
	if err := o.Endpoint1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Endpoint2 == nil {
		o.Endpoint2 = &Addresses{}
	}
	if err := o.Endpoint2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalTunnelEndpointV4); err != nil {
		return err
	}
	o.LocalTunnelEndpointV6 = make([]byte, 16)
	for i1 := range o.LocalTunnelEndpointV6 {
		i1 := i1
		if err := w.ReadData(&o.LocalTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.RemoteTunnelEndpointV4); err != nil {
		return err
	}
	o.RemoteTunnelEndpointV6 = make([]byte, 16)
	for i1 := range o.RemoteTunnelEndpointV6 {
		i1 := i1
		if err := w.ReadData(&o.RemoteTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	if o.Endpoint1Ports == nil {
		o.Endpoint1Ports = &Ports{}
	}
	if err := o.Endpoint1Ports.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Endpoint2Ports == nil {
		o.Endpoint2Ports = &Ports{}
	}
	if err := o.Endpoint2Ports.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	_ptr_wszPhase1AuthSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase1AuthSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase1AuthSet := func(ptr interface{}) { o.Phase1AuthSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase1AuthSet, _s_wszPhase1AuthSet, _ptr_wszPhase1AuthSet); err != nil {
		return err
	}
	_ptr_wszPhase2CryptoSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase2CryptoSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase2CryptoSet := func(ptr interface{}) { o.Phase2CryptoSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase2CryptoSet, _s_wszPhase2CryptoSet, _ptr_wszPhase2CryptoSet); err != nil {
		return err
	}
	_ptr_wszPhase2AuthSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase2AuthSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase2AuthSet := func(ptr interface{}) { o.Phase2AuthSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase2AuthSet, _s_wszPhase2AuthSet, _ptr_wszPhase2AuthSet); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	return nil
}

// KeyModule type represents FW_KEY_MODULE RPC enumeration.
//
// This enumeration defines the possible keying modules that the policy rule applies
// to.
type KeyModule uint16

var (
	// FW_KEY_MODULE_DEFAULT:  This value represents the default keying modules. The default
	// keying modules are implementation-specific.<23>
	KeyModuleDefault KeyModule = 0
	// FW_KEY_MODULE_IKEv1:  This value represents a keying module implementing the Internet
	// Key Exchange (IKE) protocol as specified in [RFC2409].
	KeyModuleIKEv1 KeyModule = 1
	// FW_KEY_MODULE_AUTHIP:  This value represents a keying module implementing the Authenticated
	// Internet protocol as specified in [MS-AIPS].
	KeyModuleAuthIP KeyModule = 2
	// FW_KEY_MODULE_IKEv2:  This value represents a keying module implementing the Internet
	// Key Exchange (IKEv2) protocol as specified in [RFC4306].
	KeyModuleIKEv2 KeyModule = 4
	// FW_KEY_MODULE_MAX:  This value and values that exceed this value are not valid and
	// MUST NOT be used. It is defined to provide for simplicity when writing IDL definitions
	// and code. This symbolic constant has a value of 4.
	KeyModuleMax KeyModule = 8
)

func (o KeyModule) String() string {
	switch o {
	case KeyModuleDefault:
		return "KeyModuleDefault"
	case KeyModuleIKEv1:
		return "KeyModuleIKEv1"
	case KeyModuleAuthIP:
		return "KeyModuleAuthIP"
	case KeyModuleIKEv2:
		return "KeyModuleIKEv2"
	case KeyModuleMax:
		return "KeyModuleMax"
	}
	return "Invalid"
}

// CSRule210 structure represents FW_CS_RULE2_10 RPC structure.
type CSRule210 struct {
	Next                   *CSRule210        `idl:"name:pNext" json:"next"`
	SchemaVersion          uint16            `idl:"name:wSchemaVersion" json:"schema_version"`
	RuleID                 string            `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	Name                   string            `idl:"name:wszName;string" json:"name"`
	Description            string            `idl:"name:wszDescription;string" json:"description"`
	Profiles               uint32            `idl:"name:dwProfiles" json:"profiles"`
	Endpoint1              *Addresses        `idl:"name:Endpoint1" json:"endpoint1"`
	Endpoint2              *Addresses        `idl:"name:Endpoint2" json:"endpoint2"`
	LocalInterfaceIDs      *InterfaceLUIDs   `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	LocalInterfaceTypes    uint32            `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	LocalTunnelEndpointV4  uint32            `idl:"name:dwLocalTunnelEndpointV4" json:"local_tunnel_endpoint_v4"`
	LocalTunnelEndpointV6  []byte            `idl:"name:LocalTunnelEndpointV6" json:"local_tunnel_endpoint_v6"`
	RemoteTunnelEndpointV4 uint32            `idl:"name:dwRemoteTunnelEndpointV4" json:"remote_tunnel_endpoint_v4"`
	RemoteTunnelEndpointV6 []byte            `idl:"name:RemoteTunnelEndpointV6" json:"remote_tunnel_endpoint_v6"`
	Endpoint1Ports         *Ports            `idl:"name:Endpoint1Ports" json:"endpoint1_ports"`
	Endpoint2Ports         *Ports            `idl:"name:Endpoint2Ports" json:"endpoint2_ports"`
	IPProtocol             uint16            `idl:"name:wIpProtocol" json:"ip_protocol"`
	Phase1AuthSet          string            `idl:"name:wszPhase1AuthSet;string" json:"phase1_auth_set"`
	Phase2CryptoSet        string            `idl:"name:wszPhase2CryptoSet;string" json:"phase2_crypto_set"`
	Phase2AuthSet          string            `idl:"name:wszPhase2AuthSet;string" json:"phase2_auth_set"`
	Action                 CSRuleAction      `idl:"name:Action" json:"action"`
	Flags                  uint16            `idl:"name:wFlags" json:"flags"`
	EmbeddedContext        string            `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	PlatformValidityList   *OSPlatformList   `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	Origin                 RuleOriginType    `idl:"name:Origin" json:"origin"`
	GPOName                string            `idl:"name:wszGPOName;string" json:"gpo_name"`
	Status                 RuleStatus        `idl:"name:Status" json:"status"`
	MMParentRuleID         string            `idl:"name:wszMMParentRuleId;string" json:"mm_parent_rule_id"`
	MetadataReserved       uint32            `idl:"name:MetaDataReserved" json:"metadata_reserved"`
	Metadata               []*ObjectMetadata `idl:"name:pMetaData;size_is:(((MetaDataReserved&1)?1:0))" json:"metadata"`
}

func (o *CSRule210) xxx_PreparePayload(ctx context.Context) error {
	if o.Metadata != nil && o.MetadataReserved == 0 {
		o.MetadataReserved = uint32(len(o.Metadata))
	}
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.Phase1AuthSet) > int(255) {
		return fmt.Errorf("Phase1AuthSet is out of range")
	}
	if len(o.Phase2CryptoSet) > int(255) {
		return fmt.Errorf("Phase2CryptoSet is out of range")
	}
	if len(o.Phase2AuthSet) > int(255) {
		return fmt.Errorf("Phase2AuthSet is out of range")
	}
	if o.Action < CSRuleAction(1) || o.Action > CSRuleAction(5) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(5) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if len(o.MMParentRuleID) > int(512) {
		return fmt.Errorf("MMParentRuleID is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CSRule210) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CSRule210{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if o.Endpoint1 != nil {
		if err := o.Endpoint1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Endpoint2 != nil {
		if err := o.Endpoint2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalTunnelEndpointV4); err != nil {
		return err
	}
	for i1 := range o.LocalTunnelEndpointV6 {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.LocalTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LocalTunnelEndpointV6); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RemoteTunnelEndpointV4); err != nil {
		return err
	}
	for i1 := range o.RemoteTunnelEndpointV6 {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RemoteTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteTunnelEndpointV6); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if o.Endpoint1Ports != nil {
		if err := o.Endpoint1Ports.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Endpoint2Ports != nil {
		if err := o.Endpoint2Ports.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	if o.Phase1AuthSet != "" {
		_ptr_wszPhase1AuthSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase1AuthSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase1AuthSet, _ptr_wszPhase1AuthSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Phase2CryptoSet != "" {
		_ptr_wszPhase2CryptoSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase2CryptoSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase2CryptoSet, _ptr_wszPhase2CryptoSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Phase2AuthSet != "" {
		_ptr_wszPhase2AuthSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase2AuthSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase2AuthSet, _ptr_wszPhase2AuthSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if o.MMParentRuleID != "" {
		_ptr_wszMMParentRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.MMParentRuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.MMParentRuleID, _ptr_wszMMParentRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	_exprMetaDataReserved := uint32(0)
	if (o.MetadataReserved & 1) != 0 {
		_exprMetaDataReserved = uint32(1)
	} else {
		_exprMetaDataReserved = uint32(0)
	}
	if o.Metadata != nil || _exprMetaDataReserved > 0 {
		_ptr_pMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprMetaDataReserved := uint64(0)
			if (o.MetadataReserved & 1) != 0 {
				_exprMetaDataReserved = uint64(1)
			} else {
				_exprMetaDataReserved = uint64(0)
			}
			dimSize1 := uint64(_exprMetaDataReserved)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Metadata {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Metadata[i1] != nil {
					if err := o.Metadata[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Metadata); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Metadata, _ptr_pMetaData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CSRule210) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &CSRule210{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**CSRule210) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if o.Endpoint1 == nil {
		o.Endpoint1 = &Addresses{}
	}
	if err := o.Endpoint1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Endpoint2 == nil {
		o.Endpoint2 = &Addresses{}
	}
	if err := o.Endpoint2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalTunnelEndpointV4); err != nil {
		return err
	}
	o.LocalTunnelEndpointV6 = make([]byte, 16)
	for i1 := range o.LocalTunnelEndpointV6 {
		i1 := i1
		if err := w.ReadData(&o.LocalTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.RemoteTunnelEndpointV4); err != nil {
		return err
	}
	o.RemoteTunnelEndpointV6 = make([]byte, 16)
	for i1 := range o.RemoteTunnelEndpointV6 {
		i1 := i1
		if err := w.ReadData(&o.RemoteTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	if o.Endpoint1Ports == nil {
		o.Endpoint1Ports = &Ports{}
	}
	if err := o.Endpoint1Ports.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Endpoint2Ports == nil {
		o.Endpoint2Ports = &Ports{}
	}
	if err := o.Endpoint2Ports.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	_ptr_wszPhase1AuthSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase1AuthSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase1AuthSet := func(ptr interface{}) { o.Phase1AuthSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase1AuthSet, _s_wszPhase1AuthSet, _ptr_wszPhase1AuthSet); err != nil {
		return err
	}
	_ptr_wszPhase2CryptoSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase2CryptoSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase2CryptoSet := func(ptr interface{}) { o.Phase2CryptoSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase2CryptoSet, _s_wszPhase2CryptoSet, _ptr_wszPhase2CryptoSet); err != nil {
		return err
	}
	_ptr_wszPhase2AuthSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase2AuthSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase2AuthSet := func(ptr interface{}) { o.Phase2AuthSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase2AuthSet, _s_wszPhase2AuthSet, _ptr_wszPhase2AuthSet); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	_ptr_wszMMParentRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.MMParentRuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszMMParentRuleId := func(ptr interface{}) { o.MMParentRuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.MMParentRuleID, _s_wszMMParentRuleId, _ptr_wszMMParentRuleId); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	_ptr_pMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprpMetaData := uint64(0)
		if (o.MetadataReserved & 1) != 0 {
			_exprpMetaData = uint64(1)
		} else {
			_exprpMetaData = uint64(0)
		}
		if _exprpMetaData > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprpMetaData)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Metadata", sizeInfo[0])
		}
		o.Metadata = make([]*ObjectMetadata, sizeInfo[0])
		for i1 := range o.Metadata {
			i1 := i1
			if o.Metadata[i1] == nil {
				o.Metadata[i1] = &ObjectMetadata{}
			}
			if err := o.Metadata[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMetaData := func(ptr interface{}) { o.Metadata = *ptr.(*[]*ObjectMetadata) }
	if err := w.ReadPointer(&o.Metadata, _s_pMetaData, _ptr_pMetaData); err != nil {
		return err
	}
	return nil
}

// CSRule structure represents FW_CS_RULE RPC structure.
//
// This structure describes a connection security rule.
type CSRule struct {
	// pNext:  A pointer to the next FW_CS_RULE in the list.
	Next *CSRule `idl:"name:pNext" json:"next"`
	// wSchemaVersion:  Specifies the version of the rule.
	SchemaVersion uint16 `idl:"name:wSchemaVersion" json:"schema_version"`
	// wszRuleId:  A pointer to a Unicode string that uniquely identifies the rule.
	RuleID string `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	// wszName:  A pointer to a Unicode string that provides a friendly name for the rule.
	Name string `idl:"name:wszName;string" json:"name"`
	// wszDescription:  A pointer to a Unicode string that provides a friendly description
	// for the rule.
	Description string `idl:"name:wszDescription;string" json:"description"`
	// dwProfiles:  A bitmask of the FW_PROFILE_TYPE flags. It is a condition that matches
	// traffic on the specified profiles.
	Profiles uint32 `idl:"name:dwProfiles" json:"profiles"`
	// Endpoint1:  A condition that specifies the addresses of the first host of the traffic
	// that the rule matches. An empty EndPoint1 structure means that this condition is
	// not applied (any match).
	Endpoint1 *Addresses `idl:"name:Endpoint1" json:"endpoint1"`
	// Endpoint2:  A condition that specifies the addresses of the second host of the traffic
	// that the rule matches. An empty EndPoint2 structure means that this condition is
	// not applied (any match).
	Endpoint2 *Addresses `idl:"name:Endpoint2" json:"endpoint2"`
	// LocalInterfaceIds:  A condition that specifies the list of specific network interfaces
	// that are used by the traffic that the rule matches. If the LocalInterfaceIds field
	// does not specify an interface GUID, the rule applies to all interfaces; that is,
	// the condition is not applied.
	LocalInterfaceIDs *InterfaceLUIDs `idl:"name:LocalInterfaceIds" json:"local_interface_ids"`
	// dwLocalInterfaceTypes:  A bitmask of FW_INTERFACE_TYPE. It is a condition that restricts
	// the interface types used by the traffic that the rule matches. A value of 0x00000000
	// means the condition matches all interface types.
	LocalInterfaceTypes uint32 `idl:"name:dwLocalInterfaceTypes" json:"local_interface_types"`
	// dwLocalTunnelEndpointV4:  This field specifies the IPv4 address of the endpoint that
	// the host machines use as their local endpoint when IPsec operates in tunnel mode.
	LocalTunnelEndpointV4 uint32 `idl:"name:dwLocalTunnelEndpointV4" json:"local_tunnel_endpoint_v4"`
	// LocalTunnelEndpointV6:  This field specifies the IPv6 address of the endpoint that
	// the host machines use as their local endpoint when IPsec operates in tunnel mode.
	LocalTunnelEndpointV6 []byte `idl:"name:LocalTunnelEndpointV6" json:"local_tunnel_endpoint_v6"`
	// dwRemoteTunnelEndpointV4:  This field specifies the IPv4 address of the endpoint
	// that the host machines use as their remote endpoint when IPsec operates in tunnel
	// mode.
	RemoteTunnelEndpointV4 uint32 `idl:"name:dwRemoteTunnelEndpointV4" json:"remote_tunnel_endpoint_v4"`
	// RemoteTunnelEndpointV6:  This field specifies the IPv6 address of the endpoint that
	// the host machines use as their remote endpoint when IPsec operates in tunnel mode.
	RemoteTunnelEndpointV6 []byte `idl:"name:RemoteTunnelEndpointV6" json:"remote_tunnel_endpoint_v6"`
	// Endpoint1Ports:  A condition that specifies the first host's ports of the TCP or
	// UDP traffic that the rule matches.
	Endpoint1Ports *Ports `idl:"name:Endpoint1Ports" json:"endpoint1_ports"`
	// Endpoint2Ports:  A condition that specifies the second host's ports of the TCP or
	// UDP traffic that the rule matches.
	Endpoint2Ports *Ports `idl:"name:Endpoint2Ports" json:"endpoint2_ports"`
	// wIpProtocol:  A condition that specifies the protocol of the traffic that the rule
	// matches. If the value is in the range of 0 to 255, the value describes a protocol
	// as in IETF IANA numbers (for more information, see [IANA-PROTO-NUM]). If the value
	// is 256, the rule matches any protocol.
	IPProtocol uint16 `idl:"name:wIpProtocol" json:"ip_protocol"`
	// wszPhase1AuthSet:  A Unicode string that represents the set identifier for the Phase1
	// authentication policy objects.
	Phase1AuthSet string `idl:"name:wszPhase1AuthSet;string" json:"phase1_auth_set"`
	// wszPhase2CryptoSet:  A Unicode string that represents the set identifier for the
	// Phase2 cryptographic policy objects.
	Phase2CryptoSet string `idl:"name:wszPhase2CryptoSet;string" json:"phase2_crypto_set"`
	// wszPhase2AuthSet:  A Unicode string that represents the set identifier of the Phase2
	// authentication policy objects. If this field is NULL, no second authentication is
	// performed.
	Phase2AuthSet string `idl:"name:wszPhase2AuthSet;string" json:"phase2_auth_set"`
	// Action:  The connection security action that the rule takes for the traffic matches.
	// This field MUST contain a valid value from the FW_CS_RULE_ACTION enumeration.
	Action CSRuleAction `idl:"name:Action" json:"action"`
	// wFlags:  A bit flag or flags from FW_CS_RULE_FLAGS.
	Flags uint16 `idl:"name:wFlags" json:"flags"`
	// wszEmbeddedContext:  A pointer to a Unicode string. It specifies a group name for
	// this rule. Other components in the system use this string to enable or disable a
	// group of rules by verifying that all rules have the same group name.
	EmbeddedContext string `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	// PlatformValidityList:  A condition in a rule that determines whether or not the rule
	// is enforced by the local computer based on the local computer's platform information.
	// The rule is enforced only if the local computer's operating system platform is an
	// element of the set described by PlatformValidityList.<14>
	PlatformValidityList *OSPlatformList `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	// Origin:  This field is the rule origin, as specified in the FW_RULE_ORIGIN_TYPE enumeration.
	// It MUST be filled on enumerated rules and ignored on input.
	Origin RuleOriginType `idl:"name:Origin" json:"origin"`
	// wszGPOName:  A pointer to a Unicode string containing the displayName of the GPO
	// containing this object. When adding a new object, this field is not used. The client
	// SHOULD set the value to NULL, and the server MUST ignore the value. When enumerating
	// an existing object, if the client does not set the FW_ENUM_RULES_FLAG_RESOLVE_GPO_NAME
	// flag, the server MUST set the value to NULL. Otherwise, the server MUST set the value
	// to the displayName of the GPO containing the object or NULL if the object is not
	// contained within a GPO. For details about how the server initializes an object from
	// a GPO, see section 3.1.3. For details about how the displayName of a GPO is stored,
	// see [MS-GPOL] section 2.3.
	GPOName string `idl:"name:wszGPOName;string" json:"gpo_name"`
	// Status:  The status code of the rule, as specified by the FW_RULE_STATUS enumeration.
	// This field is filled out when the structure is returned as output. On input, this
	// field MUST be set to FW_RULE_STATUS_OK.
	Status RuleStatus `idl:"name:Status" json:"status"`
	// wszMMParentRuleId:  This field is not used.
	MMParentRuleID   string `idl:"name:wszMMParentRuleId;string" json:"mm_parent_rule_id"`
	MetadataReserved uint32 `idl:"name:MetaDataReserved" json:"metadata_reserved"`
	// pMetaData:  A pointer to an FW_OBJECT_METADATA structure that contains specific metadata
	// about the current state of the connection security rule.
	Metadata []*ObjectMetadata `idl:"name:pMetaData;size_is:(((MetaDataReserved&1)?1:0))" json:"metadata"`
	// wszRemoteTunnelEndpointFqdn:  A pointer to a Unicode string containing the fully
	// qualified domain name (FQDN) of the endpoints that the host machines use as their
	// remote endpoint when IPsec operates in tunnel mode.
	RemoteTunnelEndpointFQDN string `idl:"name:wszRemoteTunnelEndpointFqdn;string" json:"remote_tunnel_endpoint_fqdn"`
	// RemoteTunnelEndpoints:  This field specifies the IPv4 and IPv6 addresses of the endpoints
	// that the host machines use as their remote endpoint when IPsec operates in tunnel
	// mode.
	RemoteTunnelEndpoints *Addresses `idl:"name:RemoteTunnelEndpoints" json:"remote_tunnel_endpoints"`
	// dwKeyModules:  A bitmask of the FW_KEY_MODULE flags. It specifies the key modules
	// used to establish the cryptographic keys used by IPsec.
	KeyModules uint32 `idl:"name:dwKeyModules" json:"key_modules"`
	// FwdPathSALifetime:  This value is the lifetime in seconds before a Phase2 established
	// key is renegotiated if the key is used to secure traffic forwarded from one interface
	// to another on the same host machine.
	FwdPathSALifetime uint32 `idl:"name:FwdPathSALifetime" json:"fwd_path_sa_lifetime"`
	// wszTransportMachineAuthzSDDL:  A pointer to a Unicode string in Security Descriptor
	// Definition Language (SDDL) format ([MS-DTYP] section 2.2.36). The security descriptor
	// describes which remote machines are allowed to send and receive traffic. Machines
	// granted access by the security descriptor are allowed to send and receive traffic.
	// Machines denied access by the security descriptor are blocked from sending and receiving
	// traffic. This field MUST be null for tunnel mode rules.
	TransportMachineAuthzSDDL string `idl:"name:wszTransportMachineAuthzSDDL;string" json:"transport_machine_authz_sddl"`
	// wszTransportUserAuthzSDDL:  A pointer to a Unicode string in Security Descriptor
	// Definition Language (SDDL) format ([MS-DTYP] section 2.2.36). The security descriptor
	// describes which remote users are allowed to send and receive traffic. Users granted
	// access by the security descriptor are allowed to send and receive traffic. Users
	// denied access by the security descriptor are blocked from sending and receiving traffic.
	// This field MUST be null for tunnel mode rules.
	TransportUserAuthzSDDL string `idl:"name:wszTransportUserAuthzSDDL;string" json:"transport_user_authz_sddl"`
}

func (o *CSRule) xxx_PreparePayload(ctx context.Context) error {
	if o.Metadata != nil && o.MetadataReserved == 0 {
		o.MetadataReserved = uint32(len(o.Metadata))
	}
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if o.IPProtocol > uint16(256) {
		return fmt.Errorf("IPProtocol is out of range")
	}
	if len(o.Phase1AuthSet) > int(255) {
		return fmt.Errorf("Phase1AuthSet is out of range")
	}
	if len(o.Phase2CryptoSet) > int(255) {
		return fmt.Errorf("Phase2CryptoSet is out of range")
	}
	if len(o.Phase2AuthSet) > int(255) {
		return fmt.Errorf("Phase2AuthSet is out of range")
	}
	if o.Action < CSRuleAction(1) || o.Action > CSRuleAction(5) {
		return fmt.Errorf("Action is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(5) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if len(o.MMParentRuleID) > int(512) {
		return fmt.Errorf("MMParentRuleID is out of range")
	}
	if len(o.RemoteTunnelEndpointFQDN) > int(512) {
		return fmt.Errorf("RemoteTunnelEndpointFQDN is out of range")
	}
	if len(o.TransportMachineAuthzSDDL) > int(10001) {
		return fmt.Errorf("TransportMachineAuthzSDDL is out of range")
	}
	if len(o.TransportUserAuthzSDDL) > int(10001) {
		return fmt.Errorf("TransportUserAuthzSDDL is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CSRule) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CSRule{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if o.Endpoint1 != nil {
		if err := o.Endpoint1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Endpoint2 != nil {
		if err := o.Endpoint2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LocalInterfaceIDs != nil {
		if err := o.LocalInterfaceIDs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalInterfaceTypes); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalTunnelEndpointV4); err != nil {
		return err
	}
	for i1 := range o.LocalTunnelEndpointV6 {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.LocalTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LocalTunnelEndpointV6); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.RemoteTunnelEndpointV4); err != nil {
		return err
	}
	for i1 := range o.RemoteTunnelEndpointV6 {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RemoteTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RemoteTunnelEndpointV6); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if o.Endpoint1Ports != nil {
		if err := o.Endpoint1Ports.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Endpoint2Ports != nil {
		if err := o.Endpoint2Ports.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Ports{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	if o.Phase1AuthSet != "" {
		_ptr_wszPhase1AuthSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase1AuthSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase1AuthSet, _ptr_wszPhase1AuthSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Phase2CryptoSet != "" {
		_ptr_wszPhase2CryptoSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase2CryptoSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase2CryptoSet, _ptr_wszPhase2CryptoSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Phase2AuthSet != "" {
		_ptr_wszPhase2AuthSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase2AuthSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase2AuthSet, _ptr_wszPhase2AuthSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Action)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if o.MMParentRuleID != "" {
		_ptr_wszMMParentRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.MMParentRuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.MMParentRuleID, _ptr_wszMMParentRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	_exprMetaDataReserved := uint32(0)
	if (o.MetadataReserved & 1) != 0 {
		_exprMetaDataReserved = uint32(1)
	} else {
		_exprMetaDataReserved = uint32(0)
	}
	if o.Metadata != nil || _exprMetaDataReserved > 0 {
		_ptr_pMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprMetaDataReserved := uint64(0)
			if (o.MetadataReserved & 1) != 0 {
				_exprMetaDataReserved = uint64(1)
			} else {
				_exprMetaDataReserved = uint64(0)
			}
			dimSize1 := uint64(_exprMetaDataReserved)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Metadata {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Metadata[i1] != nil {
					if err := o.Metadata[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Metadata); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Metadata, _ptr_pMetaData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteTunnelEndpointFQDN != "" {
		_ptr_wszRemoteTunnelEndpointFqdn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RemoteTunnelEndpointFQDN); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteTunnelEndpointFQDN, _ptr_wszRemoteTunnelEndpointFqdn); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.RemoteTunnelEndpoints != nil {
		if err := o.RemoteTunnelEndpoints.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.KeyModules); err != nil {
		return err
	}
	if err := w.WriteData(o.FwdPathSALifetime); err != nil {
		return err
	}
	if o.TransportMachineAuthzSDDL != "" {
		_ptr_wszTransportMachineAuthzSDDL := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.TransportMachineAuthzSDDL); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TransportMachineAuthzSDDL, _ptr_wszTransportMachineAuthzSDDL); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.TransportUserAuthzSDDL != "" {
		_ptr_wszTransportUserAuthzSDDL := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.TransportUserAuthzSDDL); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TransportUserAuthzSDDL, _ptr_wszTransportUserAuthzSDDL); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CSRule) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &CSRule{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**CSRule) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if o.Endpoint1 == nil {
		o.Endpoint1 = &Addresses{}
	}
	if err := o.Endpoint1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Endpoint2 == nil {
		o.Endpoint2 = &Addresses{}
	}
	if err := o.Endpoint2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LocalInterfaceIDs == nil {
		o.LocalInterfaceIDs = &InterfaceLUIDs{}
	}
	if err := o.LocalInterfaceIDs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalInterfaceTypes); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalTunnelEndpointV4); err != nil {
		return err
	}
	o.LocalTunnelEndpointV6 = make([]byte, 16)
	for i1 := range o.LocalTunnelEndpointV6 {
		i1 := i1
		if err := w.ReadData(&o.LocalTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.RemoteTunnelEndpointV4); err != nil {
		return err
	}
	o.RemoteTunnelEndpointV6 = make([]byte, 16)
	for i1 := range o.RemoteTunnelEndpointV6 {
		i1 := i1
		if err := w.ReadData(&o.RemoteTunnelEndpointV6[i1]); err != nil {
			return err
		}
	}
	if o.Endpoint1Ports == nil {
		o.Endpoint1Ports = &Ports{}
	}
	if err := o.Endpoint1Ports.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Endpoint2Ports == nil {
		o.Endpoint2Ports = &Ports{}
	}
	if err := o.Endpoint2Ports.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	_ptr_wszPhase1AuthSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase1AuthSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase1AuthSet := func(ptr interface{}) { o.Phase1AuthSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase1AuthSet, _s_wszPhase1AuthSet, _ptr_wszPhase1AuthSet); err != nil {
		return err
	}
	_ptr_wszPhase2CryptoSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase2CryptoSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase2CryptoSet := func(ptr interface{}) { o.Phase2CryptoSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase2CryptoSet, _s_wszPhase2CryptoSet, _ptr_wszPhase2CryptoSet); err != nil {
		return err
	}
	_ptr_wszPhase2AuthSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase2AuthSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase2AuthSet := func(ptr interface{}) { o.Phase2AuthSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase2AuthSet, _s_wszPhase2AuthSet, _ptr_wszPhase2AuthSet); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Action)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	_ptr_wszMMParentRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.MMParentRuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszMMParentRuleId := func(ptr interface{}) { o.MMParentRuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.MMParentRuleID, _s_wszMMParentRuleId, _ptr_wszMMParentRuleId); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	_ptr_pMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprpMetaData := uint64(0)
		if (o.MetadataReserved & 1) != 0 {
			_exprpMetaData = uint64(1)
		} else {
			_exprpMetaData = uint64(0)
		}
		if _exprpMetaData > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprpMetaData)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Metadata", sizeInfo[0])
		}
		o.Metadata = make([]*ObjectMetadata, sizeInfo[0])
		for i1 := range o.Metadata {
			i1 := i1
			if o.Metadata[i1] == nil {
				o.Metadata[i1] = &ObjectMetadata{}
			}
			if err := o.Metadata[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMetaData := func(ptr interface{}) { o.Metadata = *ptr.(*[]*ObjectMetadata) }
	if err := w.ReadPointer(&o.Metadata, _s_pMetaData, _ptr_pMetaData); err != nil {
		return err
	}
	_ptr_wszRemoteTunnelEndpointFqdn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RemoteTunnelEndpointFQDN); err != nil {
			return err
		}
		return nil
	})
	_s_wszRemoteTunnelEndpointFqdn := func(ptr interface{}) { o.RemoteTunnelEndpointFQDN = *ptr.(*string) }
	if err := w.ReadPointer(&o.RemoteTunnelEndpointFQDN, _s_wszRemoteTunnelEndpointFqdn, _ptr_wszRemoteTunnelEndpointFqdn); err != nil {
		return err
	}
	if o.RemoteTunnelEndpoints == nil {
		o.RemoteTunnelEndpoints = &Addresses{}
	}
	if err := o.RemoteTunnelEndpoints.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeyModules); err != nil {
		return err
	}
	if err := w.ReadData(&o.FwdPathSALifetime); err != nil {
		return err
	}
	_ptr_wszTransportMachineAuthzSDDL := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.TransportMachineAuthzSDDL); err != nil {
			return err
		}
		return nil
	})
	_s_wszTransportMachineAuthzSDDL := func(ptr interface{}) { o.TransportMachineAuthzSDDL = *ptr.(*string) }
	if err := w.ReadPointer(&o.TransportMachineAuthzSDDL, _s_wszTransportMachineAuthzSDDL, _ptr_wszTransportMachineAuthzSDDL); err != nil {
		return err
	}
	_ptr_wszTransportUserAuthzSDDL := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.TransportUserAuthzSDDL); err != nil {
			return err
		}
		return nil
	})
	_s_wszTransportUserAuthzSDDL := func(ptr interface{}) { o.TransportUserAuthzSDDL = *ptr.(*string) }
	if err := w.ReadPointer(&o.TransportUserAuthzSDDL, _s_wszTransportUserAuthzSDDL, _ptr_wszTransportUserAuthzSDDL); err != nil {
		return err
	}
	return nil
}

// AuthMethod type represents FW_AUTH_METHOD RPC enumeration.
//
// This enumeration defines the different authentication methods that are used for authentication.
// The IpSecPhase field of the FW_AUTH_SET containing the FW_AUTH_SUITE determines which
// authentication methods are valid for a particular authentication suite.
type AuthMethod uint16

var (
	// FW_AUTH_METHOD_INVALID:  This value MUST NOT be used. It is defined for simplicity
	// in writing IDL definitions and code. This symbolic constant has a value of 0.
	AuthMethodInvalid AuthMethod = 0
	// FW_AUTH_METHOD_ANONYMOUS:  This method does not require identity to authenticate.
	// It is equal to no authentication. This method can be used for both FW_IPSEC_PHASE_1
	// or FW_IPSEC_PHASE_2. This symbolic constant has a value of 1.
	AuthMethodAnonymous AuthMethod = 1
	// FW_AUTH_METHOD_MACHINE_KERB:  This method authenticates the identity of machines
	// by using Kerberos Protocol Extensions as specified in [MS-KILE]. This method MUST
	// be used only on FW_IPSEC_PHASE_1. This symbolic constant has a value of 2.
	AuthMethodMachineKerberos AuthMethod = 2
	// FW_AUTH_METHOD_MACHINE_SHKEY:  This method uses a previous manually shared key to
	// authenticate machine identities. This method MUST be used only on FW_IPSEC_PHASE_1.
	// This symbolic constant has a value of 3.
	AuthMethodMachineSharedKey AuthMethod = 3
	// FW_AUTH_METHOD_MACHINE_NTLM:  This method authenticates the identity of machines
	// by using the NTLM Authentication Protocol as specified in [MS-NLMP]. This method
	// MUST be used only on FW_IPSEC_PHASE_1. This symbolic constant has a value of 4.
	AuthMethodMachineNTLM AuthMethod = 4
	// FW_AUTH_METHOD_MACHINE_CERT:  This method authenticates the identity of a machine
	// by using machine certificates. This method can be used for both FW_IPSEC_PHASE_1
	// or FW_IPSEC_PHASE_2. This symbolic constant has a value of 5.
	AuthMethodMachineCert AuthMethod = 5
	// FW_AUTH_METHOD_USER_KERB:  This method authenticates user identities by using the
	// Kerberos Protocol Extensions. This method MUST be used only on FW_IPSEC_PHASE_2.
	// This symbolic constant has a value of 6.
	AuthMethodUserKerberos AuthMethod = 6
	// FW_AUTH_METHOD_USER_CERT:  This method authenticates user identities by using user
	// certificates. This method MUST be used only on FW_IPSEC_PHASE_2. This symbolic constant
	// has a value of 7.
	AuthMethodUserCert AuthMethod = 7
	// FW_AUTH_METHOD_USER_NTLM:  This method authenticates user identities by using the
	// NTLM Authentication Protocol. This method MUST be used only on FW_IPSEC_PHASE_2.
	// This symbolic constant has a value of 8.
	AuthMethodUserNTLM AuthMethod = 8
	// FW_AUTH_METHOD_MACHINE_RESERVED:  This value is invalid and MUST NOT be used. This
	// symbolic constant has a value of 9.
	AuthMethodMachineReserved AuthMethod = 9
	// FW_AUTH_METHOD_USER_RESERVED:  This value is invalid and MUST NOT be used. This
	// symbolic constant has a value of 10.
	AuthMethodUserReserved AuthMethod = 10
	// FW_AUTH_METHOD_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 11.
	AuthMethodMax AuthMethod = 11
	// FW_AUTH_METHOD_MAX_2_10:  This value and values that exceed this value are not valid
	// and MUST NOT be used by servers and clients with schema version 0x020A and earlier.
	// It is defined for simplicity in writing IDL definitions and code. This symbolic constant
	// has a value of 9.
	AuthMethodMax210 AuthMethod = 9
)

func (o AuthMethod) String() string {
	switch o {
	case AuthMethodInvalid:
		return "AuthMethodInvalid"
	case AuthMethodAnonymous:
		return "AuthMethodAnonymous"
	case AuthMethodMachineKerberos:
		return "AuthMethodMachineKerberos"
	case AuthMethodMachineSharedKey:
		return "AuthMethodMachineSharedKey"
	case AuthMethodMachineNTLM:
		return "AuthMethodMachineNTLM"
	case AuthMethodMachineCert:
		return "AuthMethodMachineCert"
	case AuthMethodUserKerberos:
		return "AuthMethodUserKerberos"
	case AuthMethodUserCert:
		return "AuthMethodUserCert"
	case AuthMethodUserNTLM:
		return "AuthMethodUserNTLM"
	case AuthMethodMachineReserved:
		return "AuthMethodMachineReserved"
	case AuthMethodUserReserved:
		return "AuthMethodUserReserved"
	case AuthMethodMax:
		return "AuthMethodMax"
	case AuthMethodMax210:
		return "AuthMethodMax210"
	}
	return "Invalid"
}

// AuthSuiteFlags type represents FW_AUTH_SUITE_FLAGS RPC enumeration.
//
// This enumeration describes bitmask flags that can be set on authentication proposals.
type AuthSuiteFlags uint16

var (
	// FW_AUTH_SUITE_FLAGS_NONE:  This value means that none of the following flags are
	// set. This value is defined for simplicity in writing IDL definitions and code.
	AuthSuiteFlagsNone AuthSuiteFlags = 0
	// FW_AUTH_SUITE_FLAGS_CERT_EXCLUDE_CA_NAME:  If this flag is set, certificate authority
	// (CA) names are excluded. This flag MUST be set only on first authentications.
	AuthSuiteFlagsCertExcludeCAName AuthSuiteFlags = 1
	// FW_AUTH_SUITE_FLAGS_HEALTH_CERT:  This flag specifies that the certificate in use
	// is a health certificate. On second authentications, if the authentication method
	// is using a machine certificate, this flag MUST be specified. Also on second authentications,
	// if the authentication method is using a user certificate, this flag MUST NOT be specified.
	AuthSuiteFlagsHealthCert AuthSuiteFlags = 2
	// FW_AUTH_SUITE_FLAGS_PERFORM_CERT_ACCOUNT_MAPPING: This flag specifies that the certificate
	// that is used maps to an account.
	AuthSuiteFlagsPerformCertAccountMapping AuthSuiteFlags = 4
	// FW_AUTH_SUITE_FLAGS_CERT_SIGNING_ECDSA256:  This flag specifies that the default
	// certificate signing algorithm of RSA MUST be replaced by the Elliptic Curve Digital
	// Signature Algorithm (ECDSA) using curves with a 256-bit prime moduli.
	AuthSuiteFlagsCertSigningECDSA256 AuthSuiteFlags = 8
	// FW_AUTH_SUITE_FLAGS_CERT_SIGNING_ECDSA384:  This flag specifies that the default
	// certificate signing algorithm of RSA MUST be replaced by the Elliptic Curve Digital
	// Signature Algorithm using curves with a 384-bit prime moduli.
	AuthSuiteFlagsCertSigningECDSA384 AuthSuiteFlags = 16
	// FW_AUTH_SUITE_FLAGS_MAX_V2_1:  This value and values that exceed this value are
	// not valid and MUST NOT be used by servers and clients with schema version 0x0201
	// and earlier. It is defined for simplicity in writing IDL definitions and code. This
	// symbolic constant has a value of 0x0020.
	AuthSuiteFlagsMaxV21 AuthSuiteFlags = 32
	// FW_AUTH_SUITE_FLAGS_INTERMEDIATE_CA:  This flag specifies that the certificate used
	// is not from a root certificate authority but from an intermediate authority in the
	// chain.
	AuthSuiteFlagsIntermediateCA AuthSuiteFlags = 32
	AuthSuiteFlagsMaxV210        AuthSuiteFlags = 64
	// FW_AUTH_SUITE_FLAGS_ALLOW_PROXY:  This flag specifies that the host machine MUST
	// use a proxy server to communicate with the Key Distribution Center (KDC) when performing
	// Kerberos authentication. For schema versions 0x0200, 0x0201, and 0x020A, this value
	// is invalid and MUST NOT be used.
	AuthSuiteFlagsAllowProxy AuthSuiteFlags = 64
	// FW_AUTH_SUITE_FLAGS_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 0x0080.
	AuthSuiteFlagsMax AuthSuiteFlags = 128
)

func (o AuthSuiteFlags) String() string {
	switch o {
	case AuthSuiteFlagsNone:
		return "AuthSuiteFlagsNone"
	case AuthSuiteFlagsCertExcludeCAName:
		return "AuthSuiteFlagsCertExcludeCAName"
	case AuthSuiteFlagsHealthCert:
		return "AuthSuiteFlagsHealthCert"
	case AuthSuiteFlagsPerformCertAccountMapping:
		return "AuthSuiteFlagsPerformCertAccountMapping"
	case AuthSuiteFlagsCertSigningECDSA256:
		return "AuthSuiteFlagsCertSigningECDSA256"
	case AuthSuiteFlagsCertSigningECDSA384:
		return "AuthSuiteFlagsCertSigningECDSA384"
	case AuthSuiteFlagsMaxV21:
		return "AuthSuiteFlagsMaxV21"
	case AuthSuiteFlagsIntermediateCA:
		return "AuthSuiteFlagsIntermediateCA"
	case AuthSuiteFlagsMaxV210:
		return "AuthSuiteFlagsMaxV210"
	case AuthSuiteFlagsAllowProxy:
		return "AuthSuiteFlagsAllowProxy"
	case AuthSuiteFlagsMax:
		return "AuthSuiteFlagsMax"
	}
	return "Invalid"
}

// AuthSuite210 structure represents FW_AUTH_SUITE2_10 RPC structure.
//
// This structure describes an IPsec authentication suite. An authentication suite is
// a proposal of a set of algorithms and parameters that specify the authentication
// method to be used. It also includes some modifiers and parameters for the authentication
// method.
type AuthSuite210 struct {
	// Method:  This field is of type FW_AUTH_METHOD. It specifies the authentication method
	// that is suggested by this proposal suite.
	Method AuthMethod `idl:"name:Method" json:"method"`
	// wFlags:  This flag is a combination of flags from FW_AUTH_SUITE_FLAGS.
	Flags     uint16                  `idl:"name:wFlags" json:"flags"`
	AuthSuite *AuthSuite210_AuthSuite `idl:"name:AuthSuite;switch_is:Method" json:"auth_suite"`
}

func (o *AuthSuite210) xxx_PreparePayload(ctx context.Context) error {
	if o.Method < AuthMethod(1) || o.Method > AuthMethod(11) {
		return fmt.Errorf("Method is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite210) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Method)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	_swAuthSuite := uint16(o.Method)
	if o.AuthSuite != nil {
		if err := o.AuthSuite.MarshalUnionNDR(ctx, w, _swAuthSuite); err != nil {
			return err
		}
	} else {
		if err := (&AuthSuite210_AuthSuite{}).MarshalUnionNDR(ctx, w, _swAuthSuite); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite210) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Method)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if o.AuthSuite == nil {
		o.AuthSuite = &AuthSuite210_AuthSuite{}
	}
	_swAuthSuite := uint16(o.Method)
	if err := o.AuthSuite.UnmarshalUnionNDR(ctx, w, _swAuthSuite); err != nil {
		return err
	}
	return nil
}

// AuthSuite210_AuthSuite structure represents FW_AUTH_SUITE2_10 union anonymous member.
//
// This structure describes an IPsec authentication suite. An authentication suite is
// a proposal of a set of algorithms and parameters that specify the authentication
// method to be used. It also includes some modifiers and parameters for the authentication
// method.
type AuthSuite210_AuthSuite struct {
	// Types that are assignable to Value
	//
	// *AuthSuite210_Cert
	// *AuthSuite210_SharedKey
	Value is_AuthSuite210_AuthSuite `json:"value"`
}

func (o *AuthSuite210_AuthSuite) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *AuthSuite210_Cert:
		if value != nil {
			return value.Cert
		}
	case *AuthSuite210_SharedKey:
		if value != nil {
			return value.SharedKey
		}
	}
	return nil
}

type is_AuthSuite210_AuthSuite interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_AuthSuite210_AuthSuite()
}

func (o *AuthSuite210_AuthSuite) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *AuthSuite210_Cert:
		switch sw {
		case uint16(5),
			uint16(7):
			return sw
		}
		return uint16(5)
	case *AuthSuite210_SharedKey:
		return uint16(3)
	}
	return uint16(0)
}

func (o *AuthSuite210_AuthSuite) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(5),
		uint16(7):
		_o, _ := o.Value.(*AuthSuite210_Cert)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AuthSuite210_Cert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*AuthSuite210_SharedKey)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AuthSuite210_SharedKey{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *AuthSuite210_AuthSuite) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(5),
		uint16(7):
		o.Value = &AuthSuite210_Cert{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &AuthSuite210_SharedKey{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// AuthSuite210_Cert structure represents AuthSuite210_AuthSuite RPC union arm.
//
// It has following labels: 5, 7
type AuthSuite210_Cert struct {
	Cert *AuthSuite210_AuthSuite_Cert `idl:"name:Cert" json:"cert"`
}

func (*AuthSuite210_Cert) is_AuthSuite210_AuthSuite() {}

func (o *AuthSuite210_Cert) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Cert != nil {
		if err := o.Cert.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthSuite210_AuthSuite_Cert{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite210_Cert) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Cert == nil {
		o.Cert = &AuthSuite210_AuthSuite_Cert{}
	}
	if err := o.Cert.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AuthSuite210_AuthSuite_Cert structure represents FW_AUTH_SUITE2_10 structure anonymous member.
//
// This structure describes an IPsec authentication suite. An authentication suite is
// a proposal of a set of algorithms and parameters that specify the authentication
// method to be used. It also includes some modifiers and parameters for the authentication
// method.
type AuthSuite210_AuthSuite_Cert struct {
	// wszCAName:  A pointer to a Unicode string. This string represents the name of the
	// certificate authority to be used to authenticate when using machine or user certificate
	// methods.
	CAName string `idl:"name:wszCAName;string;pointer:ref" json:"ca_name"`
}

func (o *AuthSuite210_AuthSuite_Cert) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite210_AuthSuite_Cert) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.CAName != "" {
		_ptr_wszCAName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.CAName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.CAName, _ptr_wszCAName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite210_AuthSuite_Cert) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_wszCAName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.CAName); err != nil {
			return err
		}
		return nil
	})
	_s_wszCAName := func(ptr interface{}) { o.CAName = *ptr.(*string) }
	if err := w.ReadPointer(&o.CAName, _s_wszCAName, _ptr_wszCAName); err != nil {
		return err
	}
	return nil
}

// AuthSuite210_SharedKey structure represents AuthSuite210_AuthSuite RPC union arm.
//
// It has following labels: 3
type AuthSuite210_SharedKey struct {
	SharedKey *AuthSuite210_AuthSuite_SharedKey `idl:"name:SHKey" json:"shared_key"`
}

func (*AuthSuite210_SharedKey) is_AuthSuite210_AuthSuite() {}

func (o *AuthSuite210_SharedKey) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SharedKey != nil {
		if err := o.SharedKey.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthSuite210_AuthSuite_SharedKey{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite210_SharedKey) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.SharedKey == nil {
		o.SharedKey = &AuthSuite210_AuthSuite_SharedKey{}
	}
	if err := o.SharedKey.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AuthSuite210_AuthSuite_SharedKey structure represents FW_AUTH_SUITE2_10 structure anonymous member.
//
// This structure describes an IPsec authentication suite. An authentication suite is
// a proposal of a set of algorithms and parameters that specify the authentication
// method to be used. It also includes some modifiers and parameters for the authentication
// method.
type AuthSuite210_AuthSuite_SharedKey struct {
	// wszSHKey:  A pointer to a Unicode string. This string is the previous, manually shared
	// secret that is used to authenticate when using preshared key methods.
	//
	//	If the method is machine certificate or user certificate, the wszCAName string MUST NOT be NULL, MUST be at least 1 character long, MUST NOT be greater than or equal to 10,000 characters, MUST NOT contain the pipe(|) character, and MUST be a CERT_X500_NAME_STR string type name encoded with X509_ASN_ENCODING. If the method is SHKEY, the wszSHKey string MUST NOT be NULL, MUST be at least 1 character long, MUST NOT be greater than or equal to 10,000 characters, and MUST NOT contain the pipe (|) character.
	SharedKey string `idl:"name:wszSHKey;string;pointer:ref" json:"shared_key"`
}

func (o *AuthSuite210_AuthSuite_SharedKey) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite210_AuthSuite_SharedKey) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.SharedKey != "" {
		_ptr_wszSHKey := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SharedKey); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SharedKey, _ptr_wszSHKey); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite210_AuthSuite_SharedKey) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_wszSHKey := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SharedKey); err != nil {
			return err
		}
		return nil
	})
	_s_wszSHKey := func(ptr interface{}) { o.SharedKey = *ptr.(*string) }
	if err := w.ReadPointer(&o.SharedKey, _s_wszSHKey, _ptr_wszSHKey); err != nil {
		return err
	}
	return nil
}

// CertCriteriaNameType type represents FW_CERT_CRITERIA_NAME_TYPE RPC enumeration.
//
// This enumeration defines the type of name to match in the certificate for a given
// criterion.
type CertCriteriaNameType uint16

var (
	// FW_CERT_CRITERIA_NAME_NONE:  Do not perform any name matching.
	CertCriteriaNameTypeNone CertCriteriaNameType = 0
	// FW_CERT_CRITERIA_NAME_DNS:  Match the DNS name in the Subject Alternative Name of
	// the certificate.
	CertCriteriaNameTypeDNS CertCriteriaNameType = 1
	// FW_CERT_CRITERIA_NAME_UPN:  Match the UPN name in the Subject Alternative Name of
	// the certificate.
	CertCriteriaNameTypeUPN CertCriteriaNameType = 2
	// FW_CERT_CRITERIA_NAME_RFC822:  Match the RFC822 name in the Subject Alternative
	// Name of the certificate.
	CertCriteriaNameTypeRFC822 CertCriteriaNameType = 3
	// FW_CERT_CRITERIA_NAME_CN:  Match the CN relative distinguished names (RDNs) in the
	// Subject DN of the certificate.
	CertCriteriaNameTypeCN CertCriteriaNameType = 4
	// FW_CERT_CRITERIA_NAME_OU:  Match the OU RDNs in the Subject DN of the certificate.
	CertCriteriaNameTypeOU CertCriteriaNameType = 5
	// FW_CERT_CRITERIA_NAME_O:  Match the O RDNs in the Subject DN of the certificate.
	CertCriteriaNameTypeO CertCriteriaNameType = 6
	// FW_CERT_CRITERIA_NAME_DC:  Match the DC RDNs in the Subject DN of the certificate.
	CertCriteriaNameTypeDC CertCriteriaNameType = 7
	// FW_CERT_CRITERIA_NAME_MAX:  To be valid, a value of this type MUST be less than
	// this constant.
	CertCriteriaNameTypeMax CertCriteriaNameType = 8
)

func (o CertCriteriaNameType) String() string {
	switch o {
	case CertCriteriaNameTypeNone:
		return "CertCriteriaNameTypeNone"
	case CertCriteriaNameTypeDNS:
		return "CertCriteriaNameTypeDNS"
	case CertCriteriaNameTypeUPN:
		return "CertCriteriaNameTypeUPN"
	case CertCriteriaNameTypeRFC822:
		return "CertCriteriaNameTypeRFC822"
	case CertCriteriaNameTypeCN:
		return "CertCriteriaNameTypeCN"
	case CertCriteriaNameTypeOU:
		return "CertCriteriaNameTypeOU"
	case CertCriteriaNameTypeO:
		return "CertCriteriaNameTypeO"
	case CertCriteriaNameTypeDC:
		return "CertCriteriaNameTypeDC"
	case CertCriteriaNameTypeMax:
		return "CertCriteriaNameTypeMax"
	}
	return "Invalid"
}

// CertCriteriaType type represents FW_CERT_CRITERIA_TYPE RPC enumeration.
//
// The FW_CERT_CRITERIA_TYPE enumeration defines whether the criteria are to be used
// for selection, validation, or both.
type CertCriteriaType uint16

var (
	// FW_CERT_CRITERIA_TYPE_BOTH:  Indicates that the criteria are to be used for both
	// certificate selection and validation.
	CertCriteriaTypeBoth CertCriteriaType = 0
	// FW_CERT_CRITERIA_TYPE_SELECTION:  Indicates that the criteria are meant for local
	// certificate selection.
	CertCriteriaTypeSelection CertCriteriaType = 1
	// FW_CERT_CRITERIA_TYPE_VALIDATION:  Indicates that the criteria are meant for validation
	// of a peer certificate.
	CertCriteriaTypeValidation CertCriteriaType = 2
	// FW_CERT_CRITERIA_TYPE_MAX:  To be valid, a value of this type MUST be less than
	// this constant.
	CertCriteriaTypeMax CertCriteriaType = 3
)

func (o CertCriteriaType) String() string {
	switch o {
	case CertCriteriaTypeBoth:
		return "CertCriteriaTypeBoth"
	case CertCriteriaTypeSelection:
		return "CertCriteriaTypeSelection"
	case CertCriteriaTypeValidation:
		return "CertCriteriaTypeValidation"
	case CertCriteriaTypeMax:
		return "CertCriteriaTypeMax"
	}
	return "Invalid"
}

// AuthCertCriteriaFlags type represents FW_AUTH_CERT_CRITERIA_FLAGS RPC enumeration.
type AuthCertCriteriaFlags uint16

var (
	AuthCertCriteriaFlagsNone          AuthCertCriteriaFlags = 0
	AuthCertCriteriaFlagsFollowRenewal AuthCertCriteriaFlags = 1
	AuthCertCriteriaFlagsMax           AuthCertCriteriaFlags = 2
)

func (o AuthCertCriteriaFlags) String() string {
	switch o {
	case AuthCertCriteriaFlagsNone:
		return "AuthCertCriteriaFlagsNone"
	case AuthCertCriteriaFlagsFollowRenewal:
		return "AuthCertCriteriaFlagsFollowRenewal"
	case AuthCertCriteriaFlagsMax:
		return "AuthCertCriteriaFlagsMax"
	}
	return "Invalid"
}

// CertCriteria structure represents FW_CERT_CRITERIA RPC structure.
//
// This structure contains fields that are used when selecting a local certificate and
// validating a remote peer's certificate during certificate authentication.
type CertCriteria struct {
	// wSchemaVersion:  Specifies the version of the criteria structure.
	SchemaVersion uint16 `idl:"name:wSchemaVersion" json:"schema_version"`
	// wFlags:  A WORD containing bit flags, whose value is defined in FW_CERT_CRITERIA_FLAGS.
	// The flag FW_AUTH_CERT_CRITERIA_FLAGS_FOLLOW_RENEWAL MUST NOT be set if the field
	// wszHash is null. If specified, the flag FW_AUTH_CERT_CRITERIA_FLAGS_FOLLOW_RENEWAL
	// MUST NOT be used if CertCriteriaType is equal to FW_CERT_CRITERIA_TYPE_VALIDATION.
	Flags uint16 `idl:"name:wFlags" json:"flags"`
	// CertCriteriaType:  Specifies the type of criteria used, as among those specified
	// in the FW_CERT_CRITERIA_TYPE enumeration. This value MUST be less than FW_CERT_CRITERIA_TYPE_MAX.
	CertCriteriaType CertCriteriaType `idl:"name:CertCriteriaType" json:"cert_criteria_type"`
	// NameType:  Specifies the type of name, as among those specified in the FW_CERT_CRITERIA_NAME_TYPE
	// enumeration. This value MUST be less than FW_CERT_CRITERIA_NAME_MAX. If the value
	// is not equal to FW_CERT_CRITERIA_NAME_NONE, then the value for wszName MUST be specified.
	NameType CertCriteriaNameType `idl:"name:NameType" json:"name_type"`
	// wszName:  A Unicode string that specifies a name corresponding to the NameType specified. The length of this Unicode string MUST be less than 10,000 characters. The name MUST not contain the pipe "|" character.
	Name string `idl:"name:wszName;string;pointer:unique" json:"name"`
	// dwNumEku:  Specifies the number of EKU element entries in the ppEku array.
	EKULength uint32 `idl:"name:dwNumEku" json:"eku_length"`
	// ppEku:  Pointer to an array of pointers to null-terminated strings. Each string in
	// the array MUST contain only characters in the range "0" to "9" or the "." character.
	// The number of elements in the array MUST be equal to the value of the dwNumEku field.
	EKU []string `idl:"name:ppEku;size_is:(dwNumEku);pointer:unique" json:"eku"`
	// wszHash:  A Unicode string that specifies the hash of the certificate. The number
	// of characters in the string MUST be equal to 40. Each character in the string MUST
	// be in one of the following ranges: "0" to "9", "a" to "f", or "A" to "F".
	Hash string `idl:"name:wszHash;string;pointer:unique" json:"hash"`
}

func (o *CertCriteria) xxx_PreparePayload(ctx context.Context) error {
	if o.EKU != nil && o.EKULength == 0 {
		o.EKULength = uint32(len(o.EKU))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertCriteria) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.CertCriteriaType)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.NameType)); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.EKULength); err != nil {
		return err
	}
	if o.EKU != nil || o.EKULength > 0 {
		_ptr_ppEku := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EKULength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.EKU {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.EKU[i1] != "" {
					_ptr_ppEku := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteCharString(ctx, w, o.EKU[i1]); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.EKU[i1], _ptr_ppEku); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.EKU); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EKU, _ptr_ppEku); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Hash != "" {
		_ptr_wszHash := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Hash); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Hash, _ptr_wszHash); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertCriteria) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.CertCriteriaType)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.NameType)); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	if err := w.ReadData(&o.EKULength); err != nil {
		return err
	}
	_ptr_ppEku := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EKULength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EKULength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.EKU", sizeInfo[0])
		}
		o.EKU = make([]string, sizeInfo[0])
		for i1 := range o.EKU {
			i1 := i1
			_ptr_ppEku := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadCharString(ctx, w, &o.EKU[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_ppEku := func(ptr interface{}) { o.EKU[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.EKU[i1], _s_ppEku, _ptr_ppEku); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ppEku := func(ptr interface{}) { o.EKU = *ptr.(*[]string) }
	if err := w.ReadPointer(&o.EKU, _s_ppEku, _ptr_ppEku); err != nil {
		return err
	}
	_ptr_wszHash := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Hash); err != nil {
			return err
		}
		return nil
	})
	_s_wszHash := func(ptr interface{}) { o.Hash = *ptr.(*string) }
	if err := w.ReadPointer(&o.Hash, _s_wszHash, _ptr_wszHash); err != nil {
		return err
	}
	return nil
}

// AuthSuite structure represents FW_AUTH_SUITE RPC structure.
//
// This structure specifies an IPsec authentication suite and includes certification
// selection criteria. An authentication suite is a proposal of a set of algorithms
// and parameters that specify the authentication method to be used.
type AuthSuite struct {
	// Method:  This field is of type FW_AUTH_METHOD. It specifies the authentication method
	// that is suggested by this proposal suite.
	Method AuthMethod `idl:"name:Method" json:"method"`
	// wFlags:  This flag is a combination of flags from FW_AUTH_SUITE_FLAGS.
	Flags     uint16               `idl:"name:wFlags" json:"flags"`
	AuthSuite *AuthSuite_AuthSuite `idl:"name:AuthSuite;switch_is:Method" json:"auth_suite"`
}

func (o *AuthSuite) xxx_PreparePayload(ctx context.Context) error {
	if o.Method < AuthMethod(1) || o.Method > AuthMethod(11) {
		return fmt.Errorf("Method is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Method)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	_swAuthSuite := uint16(o.Method)
	if o.AuthSuite != nil {
		if err := o.AuthSuite.MarshalUnionNDR(ctx, w, _swAuthSuite); err != nil {
			return err
		}
	} else {
		if err := (&AuthSuite_AuthSuite{}).MarshalUnionNDR(ctx, w, _swAuthSuite); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Method)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if o.AuthSuite == nil {
		o.AuthSuite = &AuthSuite_AuthSuite{}
	}
	_swAuthSuite := uint16(o.Method)
	if err := o.AuthSuite.UnmarshalUnionNDR(ctx, w, _swAuthSuite); err != nil {
		return err
	}
	return nil
}

// AuthSuite_AuthSuite structure represents FW_AUTH_SUITE union anonymous member.
//
// This structure specifies an IPsec authentication suite and includes certification
// selection criteria. An authentication suite is a proposal of a set of algorithms
// and parameters that specify the authentication method to be used.
type AuthSuite_AuthSuite struct {
	// Types that are assignable to Value
	//
	// *AuthSuite_Cert
	// *AuthSuite_SharedKey
	// *AuthSuite_ProxyServer
	Value is_AuthSuite_AuthSuite `json:"value"`
}

func (o *AuthSuite_AuthSuite) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *AuthSuite_Cert:
		if value != nil {
			return value.Cert
		}
	case *AuthSuite_SharedKey:
		if value != nil {
			return value.SharedKey
		}
	case *AuthSuite_ProxyServer:
		if value != nil {
			return value.ProxyServer
		}
	}
	return nil
}

type is_AuthSuite_AuthSuite interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_AuthSuite_AuthSuite()
}

func (o *AuthSuite_AuthSuite) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *AuthSuite_Cert:
		switch sw {
		case uint16(5),
			uint16(7):
			return sw
		}
		return uint16(5)
	case *AuthSuite_SharedKey:
		return uint16(3)
	case *AuthSuite_ProxyServer:
		switch sw {
		case uint16(2),
			uint16(6):
			return sw
		}
		return uint16(2)
	}
	return uint16(0)
}

func (o *AuthSuite_AuthSuite) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(5),
		uint16(7):
		_o, _ := o.Value.(*AuthSuite_Cert)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AuthSuite_Cert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*AuthSuite_SharedKey)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AuthSuite_SharedKey{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2),
		uint16(6):
		_o, _ := o.Value.(*AuthSuite_ProxyServer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AuthSuite_ProxyServer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *AuthSuite_AuthSuite) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(5),
		uint16(7):
		o.Value = &AuthSuite_Cert{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &AuthSuite_SharedKey{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2),
		uint16(6):
		o.Value = &AuthSuite_ProxyServer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// AuthSuite_Cert structure represents AuthSuite_AuthSuite RPC union arm.
//
// It has following labels: 5, 7
type AuthSuite_Cert struct {
	Cert *AuthSuite_AuthSuite_Cert `idl:"name:Cert" json:"cert"`
}

func (*AuthSuite_Cert) is_AuthSuite_AuthSuite() {}

func (o *AuthSuite_Cert) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Cert != nil {
		if err := o.Cert.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthSuite_AuthSuite_Cert{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite_Cert) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Cert == nil {
		o.Cert = &AuthSuite_AuthSuite_Cert{}
	}
	if err := o.Cert.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AuthSuite_AuthSuite_Cert structure represents FW_AUTH_SUITE structure anonymous member.
//
// This structure specifies an IPsec authentication suite and includes certification
// selection criteria. An authentication suite is a proposal of a set of algorithms
// and parameters that specify the authentication method to be used.
type AuthSuite_AuthSuite_Cert struct {
	// wszCAName:  A pointer to a Unicode string. This string represents the name of the
	// certificate authority to be used to authenticate when using machine or user certificate
	// methods.
	CAName string `idl:"name:wszCAName;string;pointer:ref" json:"ca_name"`
	// pCertCriteria:  A pointer to a structure of type PFW_CERT_CRITERIA. This field MUST
	// NOT be present unless the Method field has the value FW_AUTH_METHOD_MACHINE_CERT
	// or FW_AUTH_METHOD_USER_CERT.
	CertCriteria *CertCriteria `idl:"name:pCertCriteria;pointer:unique" json:"cert_criteria"`
}

func (o *AuthSuite_AuthSuite_Cert) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite_AuthSuite_Cert) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.CAName != "" {
		_ptr_wszCAName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.CAName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.CAName, _ptr_wszCAName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.CertCriteria != nil {
		_ptr_pCertCriteria := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CertCriteria != nil {
				if err := o.CertCriteria.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CertCriteria{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CertCriteria, _ptr_pCertCriteria); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite_AuthSuite_Cert) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_wszCAName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.CAName); err != nil {
			return err
		}
		return nil
	})
	_s_wszCAName := func(ptr interface{}) { o.CAName = *ptr.(*string) }
	if err := w.ReadPointer(&o.CAName, _s_wszCAName, _ptr_wszCAName); err != nil {
		return err
	}
	_ptr_pCertCriteria := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CertCriteria == nil {
			o.CertCriteria = &CertCriteria{}
		}
		if err := o.CertCriteria.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pCertCriteria := func(ptr interface{}) { o.CertCriteria = *ptr.(**CertCriteria) }
	if err := w.ReadPointer(&o.CertCriteria, _s_pCertCriteria, _ptr_pCertCriteria); err != nil {
		return err
	}
	return nil
}

// AuthSuite_SharedKey structure represents AuthSuite_AuthSuite RPC union arm.
//
// It has following labels: 3
type AuthSuite_SharedKey struct {
	SharedKey *AuthSuite_AuthSuite_SharedKey `idl:"name:SHKey" json:"shared_key"`
}

func (*AuthSuite_SharedKey) is_AuthSuite_AuthSuite() {}

func (o *AuthSuite_SharedKey) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SharedKey != nil {
		if err := o.SharedKey.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthSuite_AuthSuite_SharedKey{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite_SharedKey) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.SharedKey == nil {
		o.SharedKey = &AuthSuite_AuthSuite_SharedKey{}
	}
	if err := o.SharedKey.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AuthSuite_AuthSuite_SharedKey structure represents FW_AUTH_SUITE structure anonymous member.
//
// This structure specifies an IPsec authentication suite and includes certification
// selection criteria. An authentication suite is a proposal of a set of algorithms
// and parameters that specify the authentication method to be used.
type AuthSuite_AuthSuite_SharedKey struct {
	// wszSHKey:  A pointer to a Unicode string. This string is the previous, manually shared
	// secret that is used to authenticate when using preshared key methods.
	SharedKey string `idl:"name:wszSHKey;string;pointer:ref" json:"shared_key"`
}

func (o *AuthSuite_AuthSuite_SharedKey) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite_AuthSuite_SharedKey) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.SharedKey != "" {
		_ptr_wszSHKey := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SharedKey); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SharedKey, _ptr_wszSHKey); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite_AuthSuite_SharedKey) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_wszSHKey := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SharedKey); err != nil {
			return err
		}
		return nil
	})
	_s_wszSHKey := func(ptr interface{}) { o.SharedKey = *ptr.(*string) }
	if err := w.ReadPointer(&o.SharedKey, _s_wszSHKey, _ptr_wszSHKey); err != nil {
		return err
	}
	return nil
}

// AuthSuite_ProxyServer structure represents AuthSuite_AuthSuite RPC union arm.
//
// It has following labels: 2, 6
type AuthSuite_ProxyServer struct {
	ProxyServer *AuthSuite_AuthSuite_ProxyServer `idl:"name:ProxyServer" json:"proxy_server"`
}

func (*AuthSuite_ProxyServer) is_AuthSuite_AuthSuite() {}

func (o *AuthSuite_ProxyServer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ProxyServer != nil {
		if err := o.ProxyServer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthSuite_AuthSuite_ProxyServer{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite_ProxyServer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ProxyServer == nil {
		o.ProxyServer = &AuthSuite_AuthSuite_ProxyServer{}
	}
	if err := o.ProxyServer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AuthSuite_AuthSuite_ProxyServer structure represents FW_AUTH_SUITE structure anonymous member.
//
// This structure specifies an IPsec authentication suite and includes certification
// selection criteria. An authentication suite is a proposal of a set of algorithms
// and parameters that specify the authentication method to be used.
type AuthSuite_AuthSuite_ProxyServer struct {
	// wszProxyServer:  A pointer to a Unicode string specifying the fully qualified domain
	// name (FQDN) of the Kerberos proxy server. This field MUST be set if and only if the
	// FW_AUTH_SUITE_FLAGS_ALLOW_PROXY flag is set.
	//
	//	If the method is machine certificate or user certificate, the wszCAName string MUST NOT be NULL, MUST be at least 1 character long, MUST NOT be greater than or equal to 10,000 characters, MUST NOT contain the pipe(|) character, and MUST be a valid Name as defined in [X501] section 9.2. If the method is SHKEY, the wszSHKey string MUST NOT be NULL, MUST be at least 1 character long, MUST NOT be greater than or equal to 10,000 characters, and MUST NOT contain the pipe (|) character.
	//
	// If the Method is not FW_AUTH_METHOD_MACHINE_CERT or FW_AUTH_METHOD_USER_CERT then
	// the pCertCriteria field MUST be NULL.
	ProxyServer string `idl:"name:wszProxyServer;string;pointer:unique" json:"proxy_server"`
}

func (o *AuthSuite_AuthSuite_ProxyServer) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite_AuthSuite_ProxyServer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.ProxyServer != "" {
		_ptr_wszProxyServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ProxyServer); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ProxyServer, _ptr_wszProxyServer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSuite_AuthSuite_ProxyServer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_wszProxyServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ProxyServer); err != nil {
			return err
		}
		return nil
	})
	_s_wszProxyServer := func(ptr interface{}) { o.ProxyServer = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProxyServer, _s_wszProxyServer, _ptr_wszProxyServer); err != nil {
		return err
	}
	return nil
}

// AuthSetFlags type represents FW_AUTH_SET_FLAGS RPC enumeration.
//
// This enumeration represents flags that can be specified in authentication sets of
// section 2.2.65.
type AuthSetFlags uint16

var (
	// FW_AUTH_SET_FLAGS_NONE:  This value means that none of the following flags are set.
	// It is defined for simplicity in writing IDL definitions and code.
	AuthSetFlagsNone AuthSetFlags = 0
	// FW_AUTH_SET_FLAGS_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 2.
	AuthSetFlagsMax AuthSetFlags = 1
)

func (o AuthSetFlags) String() string {
	switch o {
	case AuthSetFlagsNone:
		return "AuthSetFlagsNone"
	case AuthSetFlagsMax:
		return "AuthSetFlagsMax"
	}
	return "Invalid"
}

// AuthSet210 structure represents FW_AUTH_SET2_10 RPC structure.
//
// This structure contains a list of FW_AUTH_SUITE2_10 elements that are ordered from
// highest to lowest preference and are negotiated with remote peers to establish authentication
// algorithms.
type AuthSet210 struct {
	// pNext:  A pointer to the next FW_AUTH_SET2_10 in the list.
	Next *AuthSet210 `idl:"name:pNext" json:"next"`
	// wSchemaVersion:  Specifies the version of the set.
	SchemaVersion uint16 `idl:"name:wSchemaVersion" json:"schema_version"`
	// IpSecPhase:  This field is of type FW_IPSEC_PHASE, and it specifies if this authentication
	// set applies for first or second authentications.
	IPsecPhase IPsecPhase `idl:"name:IpSecPhase" json:"ipsec_phase"`
	// wszSetId:  A pointer to a Unicode string that uniquely identifies the set. The default
	// set for this policy object is identified with the "{E5A5D32A-4BCE-4e4d-B07F-4AB1BA7E5FE3}"
	// string for Phase1 and the "{E5A5D32A-4BCE-4e4d-B07F-4AB1BA7E5FE4}" string for Phase2.
	// Default sets are merged across policy stores, and only one is enforced according
	// to predefined merge logic rules.
	SetID string `idl:"name:wszSetId;string;pointer:ref" json:"set_id"`
	// wszName:  A pointer to a Unicode string that provides a friendly name for the set.
	Name string `idl:"name:wszName;string" json:"name"`
	// wszDescription:  A pointer to a Unicode string that provides a friendly description
	// for the set.
	Description string `idl:"name:wszDescription;string" json:"description"`
	// wszEmbeddedContext:  A pointer to a Unicode string that provides a way for applications
	// to store relevant application-specific context that is related to the set.
	EmbeddedContext string `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	// dwNumSuites:  Specifies the number of authentication suites that the structure contains.
	SuitesLength uint32 `idl:"name:dwNumSuites" json:"suites_length"`
	// pSuites:  A pointer to an array of FW_AUTH_SUITE elements. The number of elements
	// is given by dwNumSuites.
	Suites []*AuthSuite210 `idl:"name:pSuites;size_is:(dwNumSuites)" json:"suites"`
	// Origin:  This field is the set origin, as specified in the FW_RULE_ORIGIN_TYPE enumeration.
	// It MUST be filled on enumerated rules and ignored on input.
	Origin RuleOriginType `idl:"name:Origin" json:"origin"`
	// wszGPOName:  A Unicode string that represents the name of the originating GPO. It
	// MUST be set if the origin is Group Policy; otherwise, it MUST be NULL.
	GPOName string `idl:"name:wszGPOName;string" json:"gpo_name"`
	// Status:  A status code of the set, as specified by the FW_RULE_STATUS enumeration.
	// This field is filled out when the structure is returned as output. On input, this
	// field MUST be set to FW_RULE_STATUS_OK.
	Status RuleStatus `idl:"name:Status" json:"status"`
	// dwAuthSetFlags:  A reserved value and not currently used. It MUST be set to 0.
	AuthSetFlags uint32 `idl:"name:dwAuthSetFlags" json:"auth_set_flags"`
}

func (o *AuthSet210) xxx_PreparePayload(ctx context.Context) error {
	if o.Suites != nil && o.SuitesLength == 0 {
		o.SuitesLength = uint32(len(o.Suites))
	}
	if o.IPsecPhase < IPsecPhase(1) || o.IPsecPhase > IPsecPhase(2) {
		return fmt.Errorf("IPsecPhase is out of range")
	}
	if len(o.SetID) > int(255) {
		return fmt.Errorf("SetID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.SuitesLength > uint32(10000) {
		return fmt.Errorf("SuitesLength is out of range")
	}
	if o.Origin > RuleOriginType(5) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSet210) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AuthSet210{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.IPsecPhase)); err != nil {
		return err
	}
	if o.SetID != "" {
		_ptr_wszSetId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SetID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SetID, _ptr_wszSetId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SuitesLength); err != nil {
		return err
	}
	if o.Suites != nil || o.SuitesLength > 0 {
		_ptr_pSuites := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SuitesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Suites {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Suites[i1] != nil {
					if err := o.Suites[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&AuthSuite210{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Suites); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&AuthSuite210{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Suites, _ptr_pSuites); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthSetFlags); err != nil {
		return err
	}
	return nil
}
func (o *AuthSet210) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &AuthSet210{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**AuthSet210) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.IPsecPhase)); err != nil {
		return err
	}
	_ptr_wszSetId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SetID); err != nil {
			return err
		}
		return nil
	})
	_s_wszSetId := func(ptr interface{}) { o.SetID = *ptr.(*string) }
	if err := w.ReadPointer(&o.SetID, _s_wszSetId, _ptr_wszSetId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SuitesLength); err != nil {
		return err
	}
	_ptr_pSuites := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SuitesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SuitesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Suites", sizeInfo[0])
		}
		o.Suites = make([]*AuthSuite210, sizeInfo[0])
		for i1 := range o.Suites {
			i1 := i1
			if o.Suites[i1] == nil {
				o.Suites[i1] = &AuthSuite210{}
			}
			if err := o.Suites[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pSuites := func(ptr interface{}) { o.Suites = *ptr.(*[]*AuthSuite210) }
	if err := w.ReadPointer(&o.Suites, _s_pSuites, _ptr_pSuites); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthSetFlags); err != nil {
		return err
	}
	return nil
}

// AuthSet structure represents FW_AUTH_SET RPC structure.
//
// This structure contains a list of FW_AUTH_SUITE elements that are ordered from highest
// to lowest preference and are negotiated with remote peers to establish authentication
// algorithms.
type AuthSet struct {
	// pNext:  A pointer to the next FW_AUTH_SET in the list.
	Next *AuthSet `idl:"name:pNext" json:"next"`
	// wSchemaVersion:  Specifies the version of the set.
	SchemaVersion uint16 `idl:"name:wSchemaVersion" json:"schema_version"`
	// IpSecPhase:  This field is of type FW_IPSEC_PHASE, and it specifies if this authentication
	// set applies for first or second authentications.
	IPsecPhase IPsecPhase `idl:"name:IpSecPhase" json:"ipsec_phase"`
	// wszSetId:  A pointer to a Unicode string that uniquely identifies the set. The primary
	// set for this policy object is identified with the "{E5A5D32A-4BCE-4e4d-B07F-4AB1BA7E5FE3}"
	// string for Phase1 and the "{E5A5D32A-4BCE-4e4d-B07F-4AB1BA7E5FE4}" string for Phase2.
	SetID string `idl:"name:wszSetId;string;pointer:ref" json:"set_id"`
	// wszName:  A pointer to a Unicode string that provides a friendly name for the set.
	Name string `idl:"name:wszName;string" json:"name"`
	// wszDescription:  A pointer to a Unicode string that provides a friendly description
	// for the set.
	Description string `idl:"name:wszDescription;string" json:"description"`
	// wszEmbeddedContext:  A pointer to a Unicode string that provides a way for applications
	// to store relevant application-specific context that is related to the set.
	EmbeddedContext string `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	// dwNumSuites:  Specifies the number of authentication suites that the structure contains.
	SuitesLength uint32 `idl:"name:dwNumSuites" json:"suites_length"`
	// pSuites:  A pointer to an array of FW_AUTH_SUITE elements. The number of elements
	// is given by dwNumSuites.
	Suites []*AuthSuite `idl:"name:pSuites;size_is:(dwNumSuites)" json:"suites"`
	// Origin:  This field is the set origin, as specified in the FW_RULE_ORIGIN_TYPE enumeration.
	// It MUST be filled on enumerated rules and ignored on input.
	Origin RuleOriginType `idl:"name:Origin" json:"origin"`
	// wszGPOName:  A pointer to a Unicode string containing the displayName of the GPO
	// containing this object. When adding a new object, this field is not used. The client
	// SHOULD set the value to NULL, and the server MUST ignore the value. When enumerating
	// an existing object, if the client does not set the FW_ENUM_RULES_FLAG_RESOLVE_GPO_NAME
	// flag, the server MUST set the value to NULL. Otherwise, the server MUST set the value
	// to the displayName of the GPO containing the object or NULL if the object is not
	// contained within a GPO. For details about how the server initializes an object from
	// a GPO, see section 3.1.3. For details about how the displayName of a GPO is stored,
	// see [MS-GPOL] section 2.3.
	GPOName string `idl:"name:wszGPOName;string" json:"gpo_name"`
	// Status:  The status code of the set which MUST be one of the values defined in the
	// FW_RULE_STATUS enumeration. This field's value is assigned when the structure is
	// returned as output. When first sent, this field MUST be set to FW_RULE_STATUS_OK.
	Status RuleStatus `idl:"name:Status" json:"status"`
	// dwAuthSetFlags:  Bit flags from FW_AUTH_SET_FLAGS.
	AuthSetFlags uint32 `idl:"name:dwAuthSetFlags" json:"auth_set_flags"`
}

func (o *AuthSet) xxx_PreparePayload(ctx context.Context) error {
	if o.Suites != nil && o.SuitesLength == 0 {
		o.SuitesLength = uint32(len(o.Suites))
	}
	if o.IPsecPhase < IPsecPhase(1) || o.IPsecPhase > IPsecPhase(2) {
		return fmt.Errorf("IPsecPhase is out of range")
	}
	if len(o.SetID) > int(255) {
		return fmt.Errorf("SetID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.SuitesLength > uint32(10000) {
		return fmt.Errorf("SuitesLength is out of range")
	}
	if o.Origin > RuleOriginType(5) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AuthSet{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.IPsecPhase)); err != nil {
		return err
	}
	if o.SetID != "" {
		_ptr_wszSetId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SetID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SetID, _ptr_wszSetId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SuitesLength); err != nil {
		return err
	}
	if o.Suites != nil || o.SuitesLength > 0 {
		_ptr_pSuites := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SuitesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Suites {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Suites[i1] != nil {
					if err := o.Suites[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&AuthSuite{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Suites); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&AuthSuite{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Suites, _ptr_pSuites); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthSetFlags); err != nil {
		return err
	}
	return nil
}
func (o *AuthSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &AuthSet{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**AuthSet) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.IPsecPhase)); err != nil {
		return err
	}
	_ptr_wszSetId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SetID); err != nil {
			return err
		}
		return nil
	})
	_s_wszSetId := func(ptr interface{}) { o.SetID = *ptr.(*string) }
	if err := w.ReadPointer(&o.SetID, _s_wszSetId, _ptr_wszSetId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SuitesLength); err != nil {
		return err
	}
	_ptr_pSuites := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SuitesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SuitesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Suites", sizeInfo[0])
		}
		o.Suites = make([]*AuthSuite, sizeInfo[0])
		for i1 := range o.Suites {
			i1 := i1
			if o.Suites[i1] == nil {
				o.Suites[i1] = &AuthSuite{}
			}
			if err := o.Suites[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pSuites := func(ptr interface{}) { o.Suites = *ptr.(*[]*AuthSuite) }
	if err := w.ReadPointer(&o.Suites, _s_pSuites, _ptr_pSuites); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthSetFlags); err != nil {
		return err
	}
	return nil
}

// CryptoKeyExchangeType type represents FW_CRYPTO_KEY_EXCHANGE_TYPE RPC enumeration.
//
// This enumeration is used to identify supported key exchange algorithms.
type CryptoKeyExchangeType uint16

var (
	// FW_CRYPTO_KEY_EXCHANGE_NONE:  This value means that there are no key exchange algorithms
	// defined. When enumerating SAs, this value MAY be returned. It MUST NOT be used for
	// other cases. This symbolic constant has a value of 0.
	CryptoKeyExchangeTypeNone CryptoKeyExchangeType = 0
	// FW_CRYPTO_KEY_EXCHANGE_DH1:  Do key exchange with Diffie-Hellman group 1. This symbolic
	// constant has a value of 1.
	CryptoKeyExchangeTypeDH1 CryptoKeyExchangeType = 1
	// FW_CRYPTO_KEY_EXCHANGE_DH2:  Do key exchange with Diffie-Hellman group 2. This symbolic
	// constant has a value of 2.
	CryptoKeyExchangeTypeDH2 CryptoKeyExchangeType = 2
	// FW_CRYPTO_KEY_EXCHANGE_ECDH256:  Do key exchange with elliptic curve Diffie-Hellman
	// 256. This symbolic constant has a value of 3.
	CryptoKeyExchangeTypeECDH256 CryptoKeyExchangeType = 3
	// FW_CRYPTO_KEY_EXCHANGE_ECDH384:  Do key exchange with elliptic curve Diffie-Hellman
	// 384. This symbolic constant has a value of 4.
	CryptoKeyExchangeTypeECDH384 CryptoKeyExchangeType = 4
	// FW_CRYPTO_KEY_EXCHANGE_DH14 = FW_CRYPTO_KEY_EXCHANGE_DH2048:  Do key exchange with
	// Diffie-Hellman group 14. This group was called Diffie-Hellman group 2048 when it
	// was introduced. The name has been changed to match standard terminology. This symbolic
	// constant has a value of 5.
	CryptoKeyExchangeTypeDH2048 CryptoKeyExchangeType = 5
	// FW_CRYPTO_KEY_EXCHANGE_DH24:  Do key exchange with Diffie-Hellman group 24. For
	// schema versions 0x0200, 0x0201, and 0x020A, this value is invalid and MUST NOT be
	// used. This symbolic constant has a value of 6.
	CryptoKeyExchangeTypeDH24 CryptoKeyExchangeType = 6
	// FW_CRYPTO_KEY_EXCHANGE_MAX:  This value and values that exceed this value are not
	// valid and MUST NOT be used. It is defined for simplicity in writing IDL definitions
	// and code. This symbolic constant has a value of 7.
	CryptoKeyExchangeTypeMax CryptoKeyExchangeType = 7
	// FW_CRYPTO_KEY_EXCHANGE_DH14:  Do key exchange with Diffie-Hellman group 14. This
	// symbolic constant has a value of 5.
	CryptoKeyExchangeTypeDH14    CryptoKeyExchangeType = 5
	CryptoKeyExchangeTypeMaxV210 CryptoKeyExchangeType = 6
)

func (o CryptoKeyExchangeType) String() string {
	switch o {
	case CryptoKeyExchangeTypeNone:
		return "CryptoKeyExchangeTypeNone"
	case CryptoKeyExchangeTypeDH1:
		return "CryptoKeyExchangeTypeDH1"
	case CryptoKeyExchangeTypeDH2:
		return "CryptoKeyExchangeTypeDH2"
	case CryptoKeyExchangeTypeECDH256:
		return "CryptoKeyExchangeTypeECDH256"
	case CryptoKeyExchangeTypeECDH384:
		return "CryptoKeyExchangeTypeECDH384"
	case CryptoKeyExchangeTypeDH2048:
		return "CryptoKeyExchangeTypeDH2048"
	case CryptoKeyExchangeTypeDH24:
		return "CryptoKeyExchangeTypeDH24"
	case CryptoKeyExchangeTypeMax:
		return "CryptoKeyExchangeTypeMax"
	case CryptoKeyExchangeTypeDH14:
		return "CryptoKeyExchangeTypeDH14"
	case CryptoKeyExchangeTypeMaxV210:
		return "CryptoKeyExchangeTypeMaxV210"
	}
	return "Invalid"
}

// CryptoEncryptionType type represents FW_CRYPTO_ENCRYPTION_TYPE RPC enumeration.
//
// This enumeration is used to identify supported encryption algorithms.
type CryptoEncryptionType uint16

var (
	// FW_CRYPTO_ENCRYPTION_NONE:  This value MUST be used only when no encryption is to
	// be performed. This is a valid value. This symbolic constant has a value of 0.
	CryptoEncryptionTypeNone CryptoEncryptionType = 0
	// FW_CRYPTO_ENCRYPTION_DES:  Uses the DES algorithm for encryption. This symbolic
	// constant has a value of 1.
	CryptoEncryptionTypeDES  CryptoEncryptionType = 1
	CryptoEncryptionType3DES CryptoEncryptionType = 2
	// FW_CRYPTO_ENCRYPTION_AES128:  Uses the AES algorithm with a 128-bit key size for
	// encryption. This symbolic constant has a value of 3.
	CryptoEncryptionTypeAES128 CryptoEncryptionType = 3
	// FW_CRYPTO_ENCRYPTION_AES192:  Uses the AES algorithm with a 192-bit key size for
	// encryption. This symbolic constant has a value of 4.
	CryptoEncryptionTypeAES192 CryptoEncryptionType = 4
	// FW_CRYPTO_ENCRYPTION_AES256:  Uses the AES algorithm with a 256-bit key size for
	// encryption. This symbolic constant has a value of 5.
	CryptoEncryptionTypeAES256 CryptoEncryptionType = 5
	// FW_CRYPTO_ENCRYPTION_AES_GCM128:  Uses the AESGCM algorithm with a 128-bit key size
	// for encryption. This symbolic constant has a value of 6.
	CryptoEncryptionTypeAESGCM128 CryptoEncryptionType = 6
	// FW_CRYPTO_ENCRYPTION_AES_GCM192:  Uses the AESGCM algorithm with a 192-bit key size
	// for encryption. This symbolic constant has a value of 7.
	CryptoEncryptionTypeAESGCM192 CryptoEncryptionType = 7
	// FW_CRYPTO_ENCRYPTION_AES_GCM256:  Uses the AESGCM algorithm with a 256-bit key size
	// for encryption. This symbolic constant has a value of 8.
	CryptoEncryptionTypeAESGCM256 CryptoEncryptionType = 8
	// FW_CRYPTO_ENCRYPTION_MAX:  This value and values that exceed this value are not
	// valid and MUST NOT be used. It is defined for simplicity in writing IDL definitions
	// and code. This symbolic constant has a value of 9.
	CryptoEncryptionTypeMax CryptoEncryptionType = 9
	// FW_CRYPTO_ENCRYPTION_MAX_V2_0:  For schema version 0x0200, this value and values
	// that exceed this value are not valid and MUST NOT be used by servers and clients
	// with schema version 0x0200 and earlier. It is defined for simplicity in writing IDL
	// definitions and describing semantic checks against policy schema versions of 0x0200.
	// This symbolic constant has a value of 6.
	CryptoEncryptionTypeMaxV20 CryptoEncryptionType = 6
)

func (o CryptoEncryptionType) String() string {
	switch o {
	case CryptoEncryptionTypeNone:
		return "CryptoEncryptionTypeNone"
	case CryptoEncryptionTypeDES:
		return "CryptoEncryptionTypeDES"
	case CryptoEncryptionType3DES:
		return "CryptoEncryptionType3DES"
	case CryptoEncryptionTypeAES128:
		return "CryptoEncryptionTypeAES128"
	case CryptoEncryptionTypeAES192:
		return "CryptoEncryptionTypeAES192"
	case CryptoEncryptionTypeAES256:
		return "CryptoEncryptionTypeAES256"
	case CryptoEncryptionTypeAESGCM128:
		return "CryptoEncryptionTypeAESGCM128"
	case CryptoEncryptionTypeAESGCM192:
		return "CryptoEncryptionTypeAESGCM192"
	case CryptoEncryptionTypeAESGCM256:
		return "CryptoEncryptionTypeAESGCM256"
	case CryptoEncryptionTypeMax:
		return "CryptoEncryptionTypeMax"
	case CryptoEncryptionTypeMaxV20:
		return "CryptoEncryptionTypeMaxV20"
	}
	return "Invalid"
}

// CryptoHashType type represents FW_CRYPTO_HASH_TYPE RPC enumeration.
//
// This enumeration is used to identify the different hashing (integrity protection)
// algorithms supported.
type CryptoHashType uint16

var (
	// FW_CRYPTO_HASH_NONE:  This value MUST be used only when no hashing is to be performed.
	// This is a valid value. This symbolic constant has a value of 0.
	CryptoHashTypeNone CryptoHashType = 0
	// FW_CRYPTO_HASH_MD5:  Use the MD5 algorithm for hashing (integrity protection). This
	// symbolic constant has a value of 1.
	CryptoHashTypeMD5 CryptoHashType = 1
	// FW_CRYPTO_HASH_SHA1:  Use the SHA1 algorithm for hashing (integrity protection).
	// This symbolic constant has a value of 2.
	CryptoHashTypeSHA1 CryptoHashType = 2
	// FW_CRYPTO_HASH_SHA256:  Use the SHA256 algorithm for hashing (integrity protection).
	// This symbolic constant has a value of 3.
	CryptoHashTypeSHA256 CryptoHashType = 3
	// FW_CRYPTO_HASH_SHA384:  Use the SHA384 algorithm for hashing (integrity protection).
	// This symbolic constant has a value of 4.
	CryptoHashTypeSHA384 CryptoHashType = 4
	// FW_CRYPTO_HASH_AES_GMAC128:  Use the AESGMAC128 algorithm for hashing (integrity
	// protection). This symbolic constant has a value of 5.
	CryptoHashTypeAESGMAC128 CryptoHashType = 5
	// FW_CRYPTO_HASH_AES_GMAC192:  Use the AESGMAC192 algorithm for hashing (integrity
	// protection). This symbolic constant has a value of 6.
	CryptoHashTypeAESGMAC192 CryptoHashType = 6
	// FW_CRYPTO_HASH_AES_GMAC256:  Use the AESGMAC256 algorithm for hashing (integrity
	// protection). This symbolic constant has a value of 7.
	CryptoHashTypeAESGMAC256 CryptoHashType = 7
	// FW_CRYPTO_HASH_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 8.
	CryptoHashTypeMax CryptoHashType = 8
	// FW_CRYPTO_HASH_MAX_V2_0:  This value and values that exceed this value are not valid
	// and MUST NOT be used by servers and clients with schema version 0x0200 and earlier.
	// It is defined for simplicity in writing IDL definitions and describing semantic checks
	// against policy schema versions of 0x0200. This symbolic constant has a value of 3.
	CryptoHashTypeMaxV20 CryptoHashType = 3
)

func (o CryptoHashType) String() string {
	switch o {
	case CryptoHashTypeNone:
		return "CryptoHashTypeNone"
	case CryptoHashTypeMD5:
		return "CryptoHashTypeMD5"
	case CryptoHashTypeSHA1:
		return "CryptoHashTypeSHA1"
	case CryptoHashTypeSHA256:
		return "CryptoHashTypeSHA256"
	case CryptoHashTypeSHA384:
		return "CryptoHashTypeSHA384"
	case CryptoHashTypeAESGMAC128:
		return "CryptoHashTypeAESGMAC128"
	case CryptoHashTypeAESGMAC192:
		return "CryptoHashTypeAESGMAC192"
	case CryptoHashTypeAESGMAC256:
		return "CryptoHashTypeAESGMAC256"
	case CryptoHashTypeMax:
		return "CryptoHashTypeMax"
	case CryptoHashTypeMaxV20:
		return "CryptoHashTypeMaxV20"
	}
	return "Invalid"
}

// CryptoProtocolType type represents FW_CRYPTO_PROTOCOL_TYPE RPC enumeration.
//
// This enumeration is used to identify the different combinations of supported IPsec
// enforcement protocols.
type CryptoProtocolType uint16

var (
	// FW_CRYPTO_PROTOCOL_INVALID:  This value MUST NOT be used. It is defined for simplicity
	// in writing IDL definitions and code. This symbolic constant has a value of 0.
	CryptoProtocolTypeInvalid CryptoProtocolType = 0
	// FW_CRYPTO_PROTOCOL_AH:  Uses the authentication header (AH) to enforce IPsec. This
	// symbolic constant has a value of 1.
	CryptoProtocolTypeAH CryptoProtocolType = 1
	// FW_CRYPTO_PROTOCOL_ESP:  Uses the ESP protocol header. This symbolic constant has
	// a value of 2.
	CryptoProtocolTypeESP CryptoProtocolType = 2
	// FW_CRYPTO_PROTOCOL_BOTH:  Uses both the AH and ESP protocol headers. This symbolic
	// constant has a value of 3.
	CryptoProtocolTypeBoth CryptoProtocolType = 3
	// FW_CRYPTO_PROTOCOL_AUTH_NO_ENCAP:  Uses no encapsulation. This sends the first packet
	// twice: once by using an ESP header and again without any header; subsequent packets
	// have no additional headers. This symbolic constant has a value of 4.
	CryptoProtocolTypeAuthNoEncap CryptoProtocolType = 4
	// FW_CRYPTO_PROTOCOL_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 5.
	CryptoProtocolTypeMax CryptoProtocolType = 5
	// FW_CRYPTO_PROTOCOL_MAX_2_1:  This value and values that exceed this value are not
	// valid and MUST NOT be used by servers and clients with schema version 0x0201 and
	// earlier. It is defined for simplicity in writing IDL definitions and code. This symbolic
	// constant has a value of 4.
	CryptoProtocolTypeMax21 CryptoProtocolType = 4
)

func (o CryptoProtocolType) String() string {
	switch o {
	case CryptoProtocolTypeInvalid:
		return "CryptoProtocolTypeInvalid"
	case CryptoProtocolTypeAH:
		return "CryptoProtocolTypeAH"
	case CryptoProtocolTypeESP:
		return "CryptoProtocolTypeESP"
	case CryptoProtocolTypeBoth:
		return "CryptoProtocolTypeBoth"
	case CryptoProtocolTypeAuthNoEncap:
		return "CryptoProtocolTypeAuthNoEncap"
	case CryptoProtocolTypeMax:
		return "CryptoProtocolTypeMax"
	case CryptoProtocolTypeMax21:
		return "CryptoProtocolTypeMax21"
	}
	return "Invalid"
}

// CryptoSetFlags type represents FW_CRYPTO_SET_FLAGS RPC enumeration.
//
// This enumeration represents flags that can be specified in crypto sets of section
// 2.2.74.
type CryptoSetFlags uint16

var (
	// FW_CRYPTO_SET_FLAGS_NONE:  This value means that none of the following flags are
	// set. It is defined for simplicity in writing IDL definitions and code.
	CryptoSetFlagsNone CryptoSetFlags = 0
	// FW_CRYPTO_SET_FLAGS_MAX:  This value and values that exceed this value are not valid
	// and MUST NOT be used. It is defined for simplicity in writing IDL definitions and
	// code. This symbolic constant has a value of 2.
	CryptoSetFlagsMax CryptoSetFlags = 1
)

func (o CryptoSetFlags) String() string {
	switch o {
	case CryptoSetFlagsNone:
		return "CryptoSetFlagsNone"
	case CryptoSetFlagsMax:
		return "CryptoSetFlagsMax"
	}
	return "Invalid"
}

// Phase1CryptoSuite structure represents FW_PHASE1_CRYPTO_SUITE RPC structure.
//
// This structure describes an IPsec Phase 1 (or main mode) cryptographic suite. A cryptographic
// suite is a proposal of a set of algorithms and parameters that specify how different
// types of enforcement and protection are suggested to be performed.
type Phase1CryptoSuite struct {
	// KeyExchange:  This field is of type FW_CRYPTO_KEY_EXCHANGE_TYPE. It specifies the
	// key exchange algorithm for this suite proposal.
	KeyExchange CryptoKeyExchangeType `idl:"name:KeyExchange" json:"key_exchange"`
	// Encryption:  This field is of type FW_CRYPTO_ENCRYPTION_TYPE. It specifies the encryption
	// algorithm for this suite proposal.
	Encryption CryptoEncryptionType `idl:"name:Encryption" json:"encryption"`
	// Hash:  This field is of type FW_CRYPTO_HASH_TYPE. It specifies the hash (integrity
	// protection) algorithm for this suite proposal.
	Hash CryptoHashType `idl:"name:Hash" json:"hash"`
	// dwP1CryptoSuiteFlags:  This is a reserved value and is not used. It MUST be set to
	// 0x00000000.
	P1CryptoSuiteFlags uint32 `idl:"name:dwP1CryptoSuiteFlags" json:"p1_crypto_suite_flags"`
}

func (o *Phase1CryptoSuite) xxx_PreparePayload(ctx context.Context) error {
	if o.KeyExchange > CryptoKeyExchangeType(6) {
		return fmt.Errorf("KeyExchange is out of range")
	}
	if o.Encryption < CryptoEncryptionType(1) || o.Encryption > CryptoEncryptionType(8) {
		return fmt.Errorf("Encryption is out of range")
	}
	if o.Hash < CryptoHashType(1) || o.Hash > CryptoHashType(7) {
		return fmt.Errorf("Hash is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Phase1CryptoSuite) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.KeyExchange)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Encryption)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Hash)); err != nil {
		return err
	}
	if err := w.WriteData(o.P1CryptoSuiteFlags); err != nil {
		return err
	}
	return nil
}
func (o *Phase1CryptoSuite) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.KeyExchange)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Encryption)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Hash)); err != nil {
		return err
	}
	if err := w.ReadData(&o.P1CryptoSuiteFlags); err != nil {
		return err
	}
	return nil
}

// Phase2CryptoSuite structure represents FW_PHASE2_CRYPTO_SUITE RPC structure.
//
// This structure describes an IPsec Phase 2 (or quick mode) cryptographic suite. A
// cryptographic suite is a proposal of a set of algorithms and parameters that specify
// how different types of enforcement and protection are suggested to be performed.
// It also suggests timeouts for which a key is valid and at which re-keying operations
// should be performed.
type Phase2CryptoSuite struct {
	// Protocol:  This field is of type FW_CRYPTO_PROTOCOL_TYPE, and it specifies the IPsec
	// enforcement protocol combination suggested for this suite.
	Protocol CryptoProtocolType `idl:"name:Protocol" json:"protocol"`
	// AhHash:  This field is of type FW_CRYPTO_HASH_TYPE. It specifies the hash (integrity
	// protection) algorithm for this suite proposal when using the authentication header
	// protocol.
	AHHash CryptoHashType `idl:"name:AhHash" json:"ah_hash"`
	// EspHash:  This field is of type FW_CRYPTO_HASH_TYPE. It specifies the hash (integrity
	// protection) algorithm for this suite proposal when using the ESP protocol.
	ESPHash CryptoHashType `idl:"name:EspHash" json:"esp_hash"`
	// Encryption:  This field is of type FW_CRYPTO_ENCRYPTION_TYPE. It specifies the encryption
	// algorithm for this suite proposal.
	Encryption CryptoEncryptionType `idl:"name:Encryption" json:"encryption"`
	// dwTimeoutMinutes:  This is the timeout or lifetime of the key used in this proposal
	// defined in minutes.
	TimeoutMinutes uint32 `idl:"name:dwTimeoutMinutes" json:"timeout_minutes"`
	// dwTimeoutKBytes:  This is the timeout or lifetime of the key used in this proposal
	// defined in kilobytes processed with this configuration.
	TimeoutKBytes uint32 `idl:"name:dwTimeoutKBytes" json:"timeout_k_bytes"`
	// dwP2CryptoSuiteFlags:  This field is reserved and is not used. It MUST be set to
	// 0x00000000.
	//
	// The following are semantic validation checks that Phase 2 cryptographic suites MUST
	// pass:
	//
	// * The *dwTimeoutMinutes* field MUST be greater than or equal to 5 and less than or
	// equal to 2,879.
	//
	// * The *dwTimeoutKBytes* field MUST be greater than or equal to 20,480 and less than
	// or equal to 2,147,483,647.
	//
	// * If the *Protocol* field is FW_CRYPTO_PROTOCOL_AH or FW_CRYPTO_PROTOCOL_BOTH, the
	// *AhHash* field MUST NOT be equal to FW_CRYPTO_HASH_NONE.
	//
	// * If the *Protocol* field is FW_CRYPTO_PROTOCOL_BOTH, the *AhHash* field MUST be
	// equal to the *EspHash* field.
	//
	// * If the *Protocol* field is FW_CRYPTO_PROTOCOL_BOTH or FW_CRYPTO_PROTOCOL_ESP, *EspHash*
	// MUST NOT be set to FW_CRYPTO_HASH_NONE or *Encryption* MUST NOT be set to FW_CRYPTO_ENCRYPTION_NONE,
	// but not both.
	P2CryptoSuiteFlags uint32 `idl:"name:dwP2CryptoSuiteFlags" json:"p2_crypto_suite_flags"`
}

func (o *Phase2CryptoSuite) xxx_PreparePayload(ctx context.Context) error {
	if o.Protocol < CryptoProtocolType(1) || o.Protocol > CryptoProtocolType(4) {
		return fmt.Errorf("Protocol is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Phase2CryptoSuite) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Protocol)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.AHHash)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.ESPHash)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Encryption)); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeoutMinutes); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeoutKBytes); err != nil {
		return err
	}
	if err := w.WriteData(o.P2CryptoSuiteFlags); err != nil {
		return err
	}
	return nil
}
func (o *Phase2CryptoSuite) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Protocol)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.AHHash)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.ESPHash)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Encryption)); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeoutMinutes); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeoutKBytes); err != nil {
		return err
	}
	if err := w.ReadData(&o.P2CryptoSuiteFlags); err != nil {
		return err
	}
	return nil
}

// Phase1CryptoFlags type represents FW_PHASE1_CRYPTO_FLAGS RPC enumeration.
//
// This enumeration is used to identify the different cryptographic flags that are supported.
type Phase1CryptoFlags uint16

var (
	// FW_PHASE1_CRYPTO_FLAGS_NONE:  This value represents no flag. It is used when none
	// of the behaviors that are represented by the defined flags in the enumeration are
	// intended. This symbolic constant has a value of 0x00.
	Phase1CryptoFlagsNone Phase1CryptoFlags = 0
	// FW_PHASE1_CRYPTO_FLAGS_DO_NOT_SKIP_DH:  This flag ensures that Authenticated IP
	// (AuthIP), as specified in [MS-AIPS], always performs a DH key exchange. (AuthIP can
	// avoid this exchange because the protocol already contains enough key material information
	// to protect the negotiation. Hence, by skipping DH, round trips and the computational
	// cost of DH are avoided.) This symbolic constant has a value of 0x01.
	Phase1CryptoFlagsDoNotSkipDH Phase1CryptoFlags = 1
	// FW_PHASE1_CRYPTO_FLAGS_MAX:  This value and values that exceed this value are not
	// valid and MUST NOT be used. It is defined for simplicity in writing IDL definitions
	// and code. This symbolic constant has a value of 0x02.
	Phase1CryptoFlagsMax Phase1CryptoFlags = 2
)

func (o Phase1CryptoFlags) String() string {
	switch o {
	case Phase1CryptoFlagsNone:
		return "Phase1CryptoFlagsNone"
	case Phase1CryptoFlagsDoNotSkipDH:
		return "Phase1CryptoFlagsDoNotSkipDH"
	case Phase1CryptoFlagsMax:
		return "Phase1CryptoFlagsMax"
	}
	return "Invalid"
}

// Phase2CryptoPFS type represents FW_PHASE2_CRYPTO_PFS RPC enumeration.
//
// This enumeration is used to identify the different perfect forward secrecy (PFS)
// options supported.
type Phase2CryptoPFS uint16

var (
	// FW_PHASE2_CRYPTO_PFS_INVALID:  This value MUST NOT be used. It is defined for simplicity
	// in writing IDL definitions and code. This symbolic constant has a value of 0.
	Phase2CryptoPFSInvalid Phase2CryptoPFS = 0
	// FW_PHASE2_CRYPTO_PFS_DISABLE:  Do not renegotiate; instead, reuse the keying material
	// negotiated in Phase 1 (main mode). This symbolic constant has a value of 1.
	Phase2CryptoPFSDisable Phase2CryptoPFS = 1
	// FW_PHASE2_CRYPTO_PFS_PHASE1:  Use Phase 1 key exchange to negotiate a Phase 2 (quick
	// mode) key for every Phase 2 negotiation. This symbolic constant has a value of 2.
	Phase2CryptoPFSPhase1 Phase2CryptoPFS = 2
	// FW_PHASE2_CRYPTO_PFS_DH1:  Use DH1 key exchange to negotiate a Phase 2 (quick mode)
	// key for every Phase 2 negotiation. This symbolic constant has a value of 3.
	Phase2CryptoPFSDH1 Phase2CryptoPFS = 3
	// FW_PHASE2_CRYPTO_PFS_DH2:  Use DH2 key exchange to negotiate a Phase 2 (quick mode)
	// key for every Phase 2 negotiation. This symbolic constant has a value of 4.
	Phase2CryptoPFSDH2 Phase2CryptoPFS = 4
	// FW_PHASE2_CRYPTO_PFS_DH2048:  Use DH2048 key exchange to negotiate a Phase 2 (quick
	// mode) key for every Phase 2 negotiation. This symbolic constant has a value of 5.
	Phase2CryptoPFSDH2048 Phase2CryptoPFS = 5
	// FW_PHASE2_CRYPTO_PFS_ECDH256:  Use ECDH256 key exchange to negotiate a Phase 2 (quick
	// mode) key for every Phase 2 negotiation. This symbolic constant has a value of 6.
	Phase2CryptoPFSECDH256 Phase2CryptoPFS = 6
	// FW_PHASE2_CRYPTO_PFS_ECDH384:  Use ECDH384 key exchange to negotiate a Phase 2 (quick
	// mode) key for every Phase 2 negotiation. This symbolic constant has a value of 7.
	Phase2CryptoPFSECDH384 Phase2CryptoPFS = 7
	// FW_PHASE2_CRYPTO_PFS_DH24:  Use DH24 key exchange to negotiate a Phase 2 (quick
	// mode) key for every Phase 2 negotiation. For schema versions 0x0200, 0x0201, and
	// 0x020A, this value is invalid and MUST NOT be used. This symbolic constant has a
	// value of 8.
	Phase2CryptoPFSDH24 Phase2CryptoPFS = 8
	// FW_PHASE2_CRYPTO_PFS_MAX:  This value and values that exceed this value are not
	// valid and MUST NOT be used. It is defined for simplicity in writing IDL definitions
	// and code. This symbolic constant has a value of 9.
	Phase2CryptoPFSMax     Phase2CryptoPFS = 9
	Phase2CryptoPFSMaxV210 Phase2CryptoPFS = 8
)

func (o Phase2CryptoPFS) String() string {
	switch o {
	case Phase2CryptoPFSInvalid:
		return "Phase2CryptoPFSInvalid"
	case Phase2CryptoPFSDisable:
		return "Phase2CryptoPFSDisable"
	case Phase2CryptoPFSPhase1:
		return "Phase2CryptoPFSPhase1"
	case Phase2CryptoPFSDH1:
		return "Phase2CryptoPFSDH1"
	case Phase2CryptoPFSDH2:
		return "Phase2CryptoPFSDH2"
	case Phase2CryptoPFSDH2048:
		return "Phase2CryptoPFSDH2048"
	case Phase2CryptoPFSECDH256:
		return "Phase2CryptoPFSECDH256"
	case Phase2CryptoPFSECDH384:
		return "Phase2CryptoPFSECDH384"
	case Phase2CryptoPFSDH24:
		return "Phase2CryptoPFSDH24"
	case Phase2CryptoPFSMax:
		return "Phase2CryptoPFSMax"
	case Phase2CryptoPFSMaxV210:
		return "Phase2CryptoPFSMaxV210"
	}
	return "Invalid"
}

// CryptoSet structure represents FW_CRYPTO_SET RPC structure.
//
// This structure contains a list of cryptographic suite elements that are ordered from
// highest to lowest preference and are negotiated with remote peers to establish cryptographic
// protection algorithms.
type CryptoSet struct {
	// pNext:  A pointer to the next FW_CRYPTO_SET in the list.
	Next *CryptoSet `idl:"name:pNext" json:"next"`
	// wSchemaVersion:  Specifies the version of the set.
	SchemaVersion uint16 `idl:"name:wSchemaVersion" json:"schema_version"`
	// IpSecPhase:  This field is of type FW_IPSEC_PHASE, and it specifies if this cryptographic
	// set applies for Phase1 (main mode) or Phase2 (quick mode).
	IPsecPhase IPsecPhase `idl:"name:IpSecPhase" json:"ipsec_phase"`
	// wszSetId:  A pointer to a Unicode string that uniquely identifies the set. The primary
	// set for this policy object is identified with the "{E5A5D32A-4BCE-4e4d-B07F-4AB1BA7E5FE1}"
	// string for Phase1 and with the "{E5A5D32A-4BCE-4e4d-B07F-4AB1BA7E5FE2}" string for
	// Phase2.
	SetID string `idl:"name:wszSetId;string;pointer:ref" json:"set_id"`
	// wszName:  A pointer to a Unicode string that provides a friendly name for the set.
	Name string `idl:"name:wszName;string" json:"name"`
	// wszDescription:  A pointer to a Unicode string that provides a friendly description
	// for the set.
	Description string `idl:"name:wszDescription;string" json:"description"`
	// wszEmbeddedContext:  A pointer to a Unicode string. A client implementation MAY use
	// this field to store implementation-specific client context. The server MUST NOT interpret
	// the value of this string. The server MUST preserve the value of this string unmodified.
	EmbeddedContext string               `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	CryptoSet       *CryptoSet_CryptoSet `idl:"name:CryptoSet;switch_is:IpSecPhase" json:"crypto_set"`
	// Origin:  This field is the set origin, as specified in the FW_RULE_ORIGIN_TYPE enumeration.
	// It MUST be filled on enumerated rules and ignored on input.
	Origin RuleOriginType `idl:"name:Origin" json:"origin"`
	// wszGPOName:  A pointer to a Unicode string containing the displayName of the GPO
	// containing this object. When adding a new object, this field is not used. The client
	// SHOULD set the value to NULL, and the server MUST ignore the value. When enumerating
	// an existing object, if the client does not set the FW_ENUM_RULES_FLAG_RESOLVE_GPO_NAME
	// flag, the server MUST set the value to NULL. Otherwise, the server MUST set the value
	// to the displayName of the GPO containing the object or NULL if the object is not
	// contained within a GPO. For details about how the server initializes an object from
	// a GPO, see section 3.1.3. For details about how the displayName of a GPO is stored,
	// see [MS-GPOL] section 2.3.
	GPOName string `idl:"name:wszGPOName;string" json:"gpo_name"`
	// Status:  The status code of the set, as specified by the FW_RULE_STATUS enumeration.
	// This field is filled out when the structure is returned as output. On input, this
	// field MUST be set to FW_RULE_STATUS_OK.
	Status RuleStatus `idl:"name:Status" json:"status"`
	// dwCryptoSetFlags:  Bit flags from FW_CRYPTO_SET_FLAGS.
	CryptoSetFlags uint32 `idl:"name:dwCryptoSetFlags" json:"crypto_set_flags"`
}

func (o *CryptoSet) xxx_PreparePayload(ctx context.Context) error {
	if o.IPsecPhase < IPsecPhase(1) || o.IPsecPhase > IPsecPhase(2) {
		return fmt.Errorf("IPsecPhase is out of range")
	}
	if len(o.SetID) > int(255) {
		return fmt.Errorf("SetID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(5) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CryptoSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CryptoSet{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.IPsecPhase)); err != nil {
		return err
	}
	if o.SetID != "" {
		_ptr_wszSetId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SetID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SetID, _ptr_wszSetId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	_swCryptoSet := uint16(o.IPsecPhase)
	if o.CryptoSet != nil {
		if err := o.CryptoSet.MarshalUnionNDR(ctx, w, _swCryptoSet); err != nil {
			return err
		}
	} else {
		if err := (&CryptoSet_CryptoSet{}).MarshalUnionNDR(ctx, w, _swCryptoSet); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(o.CryptoSetFlags); err != nil {
		return err
	}
	return nil
}
func (o *CryptoSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &CryptoSet{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**CryptoSet) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.IPsecPhase)); err != nil {
		return err
	}
	_ptr_wszSetId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SetID); err != nil {
			return err
		}
		return nil
	})
	_s_wszSetId := func(ptr interface{}) { o.SetID = *ptr.(*string) }
	if err := w.ReadPointer(&o.SetID, _s_wszSetId, _ptr_wszSetId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.CryptoSet == nil {
		o.CryptoSet = &CryptoSet_CryptoSet{}
	}
	_swCryptoSet := uint16(o.IPsecPhase)
	if err := o.CryptoSet.UnmarshalUnionNDR(ctx, w, _swCryptoSet); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData(&o.CryptoSetFlags); err != nil {
		return err
	}
	return nil
}

// CryptoSet_CryptoSet structure represents FW_CRYPTO_SET union anonymous member.
//
// This structure contains a list of cryptographic suite elements that are ordered from
// highest to lowest preference and are negotiated with remote peers to establish cryptographic
// protection algorithms.
type CryptoSet_CryptoSet struct {
	// Types that are assignable to Value
	//
	// *CryptoSet_Phase1
	// *CryptoSet_Phase2
	Value is_CryptoSet_CryptoSet `json:"value"`
}

func (o *CryptoSet_CryptoSet) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *CryptoSet_Phase1:
		if value != nil {
			return value.Phase1
		}
	case *CryptoSet_Phase2:
		if value != nil {
			return value.Phase2
		}
	}
	return nil
}

type is_CryptoSet_CryptoSet interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_CryptoSet_CryptoSet()
}

func (o *CryptoSet_CryptoSet) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *CryptoSet_Phase1:
		return uint16(1)
	case *CryptoSet_Phase2:
		return uint16(2)
	}
	return uint16(0)
}

func (o *CryptoSet_CryptoSet) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*CryptoSet_Phase1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&CryptoSet_Phase1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*CryptoSet_Phase2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&CryptoSet_Phase2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *CryptoSet_CryptoSet) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &CryptoSet_Phase1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &CryptoSet_Phase2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// CryptoSet_Phase1 structure represents CryptoSet_CryptoSet RPC union arm.
//
// It has following labels: 1
type CryptoSet_Phase1 struct {
	Phase1 *CryptoSet_CryptoSet_Phase1 `idl:"name:Phase1" json:"phase1"`
}

func (*CryptoSet_Phase1) is_CryptoSet_CryptoSet() {}

func (o *CryptoSet_Phase1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Phase1 != nil {
		if err := o.Phase1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CryptoSet_CryptoSet_Phase1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *CryptoSet_Phase1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Phase1 == nil {
		o.Phase1 = &CryptoSet_CryptoSet_Phase1{}
	}
	if err := o.Phase1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// CryptoSet_CryptoSet_Phase1 structure represents FW_CRYPTO_SET structure anonymous member.
//
// This structure contains a list of cryptographic suite elements that are ordered from
// highest to lowest preference and are negotiated with remote peers to establish cryptographic
// protection algorithms.
type CryptoSet_CryptoSet_Phase1 struct {
	// wFlags:  This field is a combination of the FW_PHASE1_CRYPTO_FLAGS enumeration bit
	// flags.
	Flags uint16 `idl:"name:wFlags" json:"flags"`
	// dwNumPhase1Suites:  Specifies the number of Phase1 suites that the structure contains.
	Phase1SuitesLength uint32 `idl:"name:dwNumPhase1Suites" json:"phase1_suites_length"`
	// pPhase1Suites:  A pointer to an array of dwNumPhase1Suites contiguous FW_PHASE1_CRYPTO_SUITE
	// elements.
	Phase1Suites    []*Phase1CryptoSuite `idl:"name:pPhase1Suites;size_is:(dwNumPhase1Suites)" json:"phase1_suites"`
	TimeoutMinutes  uint32               `idl:"name:dwTimeOutMinutes" json:"timeout_minutes"`
	TimeoutSessions uint32               `idl:"name:dwTimeOutSessions" json:"timeout_sessions"`
}

func (o *CryptoSet_CryptoSet_Phase1) xxx_PreparePayload(ctx context.Context) error {
	if o.Phase1Suites != nil && o.Phase1SuitesLength == 0 {
		o.Phase1SuitesLength = uint32(len(o.Phase1Suites))
	}
	if o.Phase1SuitesLength > uint32(10000) {
		return fmt.Errorf("Phase1SuitesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CryptoSet_CryptoSet_Phase1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Phase1SuitesLength); err != nil {
		return err
	}
	if o.Phase1Suites != nil || o.Phase1SuitesLength > 0 {
		_ptr_pPhase1Suites := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Phase1SuitesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Phase1Suites {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Phase1Suites[i1] != nil {
					if err := o.Phase1Suites[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Phase1CryptoSuite{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Phase1Suites); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Phase1CryptoSuite{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase1Suites, _ptr_pPhase1Suites); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TimeoutMinutes); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeoutSessions); err != nil {
		return err
	}
	return nil
}
func (o *CryptoSet_CryptoSet_Phase1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Phase1SuitesLength); err != nil {
		return err
	}
	_ptr_pPhase1Suites := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Phase1SuitesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Phase1SuitesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Phase1Suites", sizeInfo[0])
		}
		o.Phase1Suites = make([]*Phase1CryptoSuite, sizeInfo[0])
		for i1 := range o.Phase1Suites {
			i1 := i1
			if o.Phase1Suites[i1] == nil {
				o.Phase1Suites[i1] = &Phase1CryptoSuite{}
			}
			if err := o.Phase1Suites[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pPhase1Suites := func(ptr interface{}) { o.Phase1Suites = *ptr.(*[]*Phase1CryptoSuite) }
	if err := w.ReadPointer(&o.Phase1Suites, _s_pPhase1Suites, _ptr_pPhase1Suites); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeoutMinutes); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeoutSessions); err != nil {
		return err
	}
	return nil
}

// CryptoSet_Phase2 structure represents CryptoSet_CryptoSet RPC union arm.
//
// It has following labels: 2
type CryptoSet_Phase2 struct {
	Phase2 *CryptoSet_CryptoSet_Phase2 `idl:"name:Phase2" json:"phase2"`
}

func (*CryptoSet_Phase2) is_CryptoSet_CryptoSet() {}

func (o *CryptoSet_Phase2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Phase2 != nil {
		if err := o.Phase2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CryptoSet_CryptoSet_Phase2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *CryptoSet_Phase2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Phase2 == nil {
		o.Phase2 = &CryptoSet_CryptoSet_Phase2{}
	}
	if err := o.Phase2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// CryptoSet_CryptoSet_Phase2 structure represents FW_CRYPTO_SET structure anonymous member.
//
// This structure contains a list of cryptographic suite elements that are ordered from
// highest to lowest preference and are negotiated with remote peers to establish cryptographic
// protection algorithms.
type CryptoSet_CryptoSet_Phase2 struct {
	// Pfs:  This field MUST contain a valid value of those in the FW_PHASE2_CRYPTO_PFS
	// enumeration. It describes the perfect forward secrecy used for quick mode cryptographic
	// operations.
	PFS Phase2CryptoPFS `idl:"name:Pfs" json:"pfs"`
	// dwNumPhase2Suites:  Specifies the number of Phase2 suites that the structure contains.
	Phase2SuitesLength uint32 `idl:"name:dwNumPhase2Suites" json:"phase2_suites_length"`
	// pPhase2Suites:  A pointer to an array of FW_PHASE2_CRYPTO_SUITE elements. The number
	// of elements is given by dwNumPhase2Suites.
	Phase2Suites []*Phase2CryptoSuite `idl:"name:pPhase2Suites;size_is:(dwNumPhase2Suites)" json:"phase2_suites"`
}

func (o *CryptoSet_CryptoSet_Phase2) xxx_PreparePayload(ctx context.Context) error {
	if o.Phase2Suites != nil && o.Phase2SuitesLength == 0 {
		o.Phase2SuitesLength = uint32(len(o.Phase2Suites))
	}
	if o.Phase2SuitesLength > uint32(10000) {
		return fmt.Errorf("Phase2SuitesLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CryptoSet_CryptoSet_Phase2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.PFS)); err != nil {
		return err
	}
	if err := w.WriteData(o.Phase2SuitesLength); err != nil {
		return err
	}
	if o.Phase2Suites != nil || o.Phase2SuitesLength > 0 {
		_ptr_pPhase2Suites := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Phase2SuitesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Phase2Suites {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Phase2Suites[i1] != nil {
					if err := o.Phase2Suites[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Phase2CryptoSuite{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Phase2Suites); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Phase2CryptoSuite{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase2Suites, _ptr_pPhase2Suites); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CryptoSet_CryptoSet_Phase2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.PFS)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Phase2SuitesLength); err != nil {
		return err
	}
	_ptr_pPhase2Suites := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Phase2SuitesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Phase2SuitesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Phase2Suites", sizeInfo[0])
		}
		o.Phase2Suites = make([]*Phase2CryptoSuite, sizeInfo[0])
		for i1 := range o.Phase2Suites {
			i1 := i1
			if o.Phase2Suites[i1] == nil {
				o.Phase2Suites[i1] = &Phase2CryptoSuite{}
			}
			if err := o.Phase2Suites[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pPhase2Suites := func(ptr interface{}) { o.Phase2Suites = *ptr.(*[]*Phase2CryptoSuite) }
	if err := w.ReadPointer(&o.Phase2Suites, _s_pPhase2Suites, _ptr_pPhase2Suites); err != nil {
		return err
	}
	return nil
}

// ByteBlob structure represents FW_BYTE_BLOB RPC structure.
//
// This structure contains a memory section. The format of the memory is defined by
// the context where it is used; for example, see the SubjectName field of the FW_CERT_INFO
// structure.
type ByteBlob struct {
	// dwSize:  This field specifies the size in octets of the Blob field.
	Size uint32 `idl:"name:dwSize" json:"size"`
	// Blob:  A pointer to an array of dwSize octets.
	Blob []byte `idl:"name:Blob;size_is:(dwSize)" json:"blob"`
}

func (o *ByteBlob) xxx_PreparePayload(ctx context.Context) error {
	if o.Blob != nil && o.Size == 0 {
		o.Size = uint32(len(o.Blob))
	}
	if o.Size > uint32(10000) {
		return fmt.Errorf("Size is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ByteBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Blob != nil || o.Size > 0 {
		_ptr_Blob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Blob {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Blob[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Blob); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Blob, _ptr_Blob); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ByteBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_Blob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Blob", sizeInfo[0])
		}
		o.Blob = make([]byte, sizeInfo[0])
		for i1 := range o.Blob {
			i1 := i1
			if err := w.ReadData(&o.Blob[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Blob := func(ptr interface{}) { o.Blob = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Blob, _s_Blob, _ptr_Blob); err != nil {
		return err
	}
	return nil
}

// CookiePair structure represents FW_COOKIE_PAIR RPC structure.
//
// This structure holds random numbers generated out of IPsec negotiations.
type CookiePair struct {
	// Initiator:  A random number that maps to the negotiated state that is a security
	// association of the machine that initiated communication and, hence, initiated IKE/AuthIP
	// (for more information, see [RFC2409]) as specified in [MS-IKEE] and [MS-AIPS] traffic.
	Initiator uint64 `idl:"name:Initiator" json:"initiator"`
	// Responder:  A random number that maps to the negotiated state that is a security
	// association of the machine that responded to the communication and, hence, responded
	// to the IKE/AuthIP traffic.
	Responder uint64 `idl:"name:Responder" json:"responder"`
}

func (o *CookiePair) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CookiePair) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Initiator); err != nil {
		return err
	}
	if err := w.WriteData(o.Responder); err != nil {
		return err
	}
	return nil
}
func (o *CookiePair) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Initiator); err != nil {
		return err
	}
	if err := w.ReadData(&o.Responder); err != nil {
		return err
	}
	return nil
}

// Phase1KeyModuleType type represents FW_PHASE1_KEY_MODULE_TYPE RPC enumeration.
//
// This enumeration identifies the different IPsec Key Exchange negotiation protocols
// that can be used.
type Phase1KeyModuleType uint16

var (
	// FW_PHASE1_KEY_MODULE_INVALID:  The FW_PHASE1_KEY_MODULE_INVALID constant MUST NOT
	// be used. It is defined for simplicity in writing IDL definitions and code. This symbolic
	// constant has a value of 0.
	Phase1KeyModuleTypeInvalid Phase1KeyModuleType = 0
	// FW_PHASE1_KEY_MODULE_IKE:  The keying protocol was IKE. This symbolic constant has
	// a value of 1.
	Phase1KeyModuleTypeIKE Phase1KeyModuleType = 1
	// FW_PHASE1_KEY_MODULE_AUTH_IP:  The keying protocol was AuthIP. This symbolic constant
	// has a value of 2.
	Phase1KeyModuleTypeAuthIP Phase1KeyModuleType = 2
	// FW_PHASE1_KEY_MODULE_MAX:  This value and values that exceed this value are not
	// valid and MUST NOT be used. It is defined for simplicity in writing IDL definitions
	// and code. This symbolic constant has a value of 3.
	Phase1KeyModuleTypeMax Phase1KeyModuleType = 3
)

func (o Phase1KeyModuleType) String() string {
	switch o {
	case Phase1KeyModuleTypeInvalid:
		return "Phase1KeyModuleTypeInvalid"
	case Phase1KeyModuleTypeIKE:
		return "Phase1KeyModuleTypeIKE"
	case Phase1KeyModuleTypeAuthIP:
		return "Phase1KeyModuleTypeAuthIP"
	case Phase1KeyModuleTypeMax:
		return "Phase1KeyModuleTypeMax"
	}
	return "Invalid"
}

// CertInfo structure represents FW_CERT_INFO RPC structure.
//
// This structure represents information on the certificate used in the certificate-based
// authentication mechanisms.
type CertInfo struct {
	// SubjectName:  The subject name of the certificate represented as a FW_BYTE_BLOB type.
	// This BLOB is an ASN.1-encoded sequence of RDN attributes.
	SubjectName *ByteBlob `idl:"name:SubjectName" json:"subject_name"`
	// dwCertFlags:  This field can be a combination of bit flags from FW_AUTH_SUITE_FLAGS.
	// This field MUST use only health certificate or certificate to account mapping flags,
	// which represent certificate characteristics.
	CertFlags uint32 `idl:"name:dwCertFlags" json:"cert_flags"`
}

func (o *CertInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.CertFlags > uint32(127) {
		return fmt.Errorf("CertFlags is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.SubjectName != nil {
		if err := o.SubjectName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ByteBlob{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CertFlags); err != nil {
		return err
	}
	return nil
}
func (o *CertInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.SubjectName == nil {
		o.SubjectName = &ByteBlob{}
	}
	if err := o.SubjectName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.CertFlags); err != nil {
		return err
	}
	return nil
}

// AuthInfo structure represents FW_AUTH_INFO RPC structure.
//
// This structure contains information on the local and remote hosts that resulted from
// the authentication methods performed between them.
type AuthInfo struct {
	// AuthMethod:  This field contains the authentication method used to establish the
	// identities of the endpoints and is stored in the security association. The field
	// can take valid values from the FW_AUTH_METHOD enumeration.
	AuthMethod AuthMethod         `idl:"name:AuthMethod" json:"auth_method"`
	AuthInfo   *AuthInfo_AuthInfo `idl:"name:AuthInfo;switch_is:AuthMethod" json:"auth_info"`
	// dwAuthInfoFlags:  Reserved value and not currently used. It MUST be set to 0.
	AuthInfoFlags uint32 `idl:"name:dwAuthInfoFlags" json:"auth_info_flags"`
}

func (o *AuthInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.AuthMethod < AuthMethod(1) || o.AuthMethod > AuthMethod(11) {
		return fmt.Errorf("AuthMethod is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.AuthMethod)); err != nil {
		return err
	}
	_swAuthInfo := uint16(o.AuthMethod)
	if o.AuthInfo != nil {
		if err := o.AuthInfo.MarshalUnionNDR(ctx, w, _swAuthInfo); err != nil {
			return err
		}
	} else {
		if err := (&AuthInfo_AuthInfo{}).MarshalUnionNDR(ctx, w, _swAuthInfo); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AuthInfoFlags); err != nil {
		return err
	}
	return nil
}
func (o *AuthInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.AuthMethod)); err != nil {
		return err
	}
	if o.AuthInfo == nil {
		o.AuthInfo = &AuthInfo_AuthInfo{}
	}
	_swAuthInfo := uint16(o.AuthMethod)
	if err := o.AuthInfo.UnmarshalUnionNDR(ctx, w, _swAuthInfo); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthInfoFlags); err != nil {
		return err
	}
	return nil
}

// AuthInfo_AuthInfo structure represents FW_AUTH_INFO union anonymous member.
//
// This structure contains information on the local and remote hosts that resulted from
// the authentication methods performed between them.
type AuthInfo_AuthInfo struct {
	// Types that are assignable to Value
	//
	// *AuthInfo_Cert
	// *AuthInfo_Kerberos
	Value is_AuthInfo_AuthInfo `json:"value"`
}

func (o *AuthInfo_AuthInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *AuthInfo_Cert:
		if value != nil {
			return value.Cert
		}
	case *AuthInfo_Kerberos:
		if value != nil {
			return value.Kerberos
		}
	}
	return nil
}

type is_AuthInfo_AuthInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_AuthInfo_AuthInfo()
}

func (o *AuthInfo_AuthInfo) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *AuthInfo_Cert:
		switch sw {
		case uint16(5),
			uint16(7):
			return sw
		}
		return uint16(5)
	case *AuthInfo_Kerberos:
		switch sw {
		case uint16(2),
			uint16(6),
			uint16(9),
			uint16(10):
			return sw
		}
		return uint16(2)
	}
	return uint16(0)
}

func (o *AuthInfo_AuthInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(5),
		uint16(7):
		_o, _ := o.Value.(*AuthInfo_Cert)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AuthInfo_Cert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2),
		uint16(6),
		uint16(9),
		uint16(10):
		_o, _ := o.Value.(*AuthInfo_Kerberos)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AuthInfo_Kerberos{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *AuthInfo_AuthInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(5),
		uint16(7):
		o.Value = &AuthInfo_Cert{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2),
		uint16(6),
		uint16(9),
		uint16(10):
		o.Value = &AuthInfo_Kerberos{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// AuthInfo_Cert structure represents AuthInfo_AuthInfo RPC union arm.
//
// It has following labels: 5, 7
type AuthInfo_Cert struct {
	Cert *AuthInfo_AuthInfo_Cert `idl:"name:Cert" json:"cert"`
}

func (*AuthInfo_Cert) is_AuthInfo_AuthInfo() {}

func (o *AuthInfo_Cert) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Cert != nil {
		if err := o.Cert.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthInfo_AuthInfo_Cert{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthInfo_Cert) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Cert == nil {
		o.Cert = &AuthInfo_AuthInfo_Cert{}
	}
	if err := o.Cert.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AuthInfo_AuthInfo_Cert structure represents FW_AUTH_INFO structure anonymous member.
//
// This structure contains information on the local and remote hosts that resulted from
// the authentication methods performed between them.
type AuthInfo_AuthInfo_Cert struct {
	// MyCert:  This field contains the subject name and certification flags (health, account
	// mapping, exclude CA) from the certificate of the local host that was used in the
	// authentication process when a certificate-based authentication method is used.
	MyCert *CertInfo `idl:"name:MyCert" json:"my_cert"`
	// PeerCert:  This field contains the subject name and certification flags (health,
	// account mapping, exclude CA) from the certificate of the remote host that was used
	// in the authentication process when a certificate-based authentication method is used.
	PeerCert *CertInfo `idl:"name:PeerCert" json:"peer_cert"`
}

func (o *AuthInfo_AuthInfo_Cert) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthInfo_AuthInfo_Cert) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.MyCert != nil {
		if err := o.MyCert.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CertInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PeerCert != nil {
		if err := o.PeerCert.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CertInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthInfo_AuthInfo_Cert) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.MyCert == nil {
		o.MyCert = &CertInfo{}
	}
	if err := o.MyCert.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PeerCert == nil {
		o.PeerCert = &CertInfo{}
	}
	if err := o.PeerCert.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AuthInfo_Kerberos structure represents AuthInfo_AuthInfo RPC union arm.
//
// It has following labels: 2, 6, 9, 10
type AuthInfo_Kerberos struct {
	Kerberos *AuthInfo_AuthInfo_Kerberos `idl:"name:Kerb" json:"kerberos"`
}

func (*AuthInfo_Kerberos) is_AuthInfo_AuthInfo() {}

func (o *AuthInfo_Kerberos) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Kerberos != nil {
		if err := o.Kerberos.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AuthInfo_AuthInfo_Kerberos{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthInfo_Kerberos) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Kerberos == nil {
		o.Kerberos = &AuthInfo_AuthInfo_Kerberos{}
	}
	if err := o.Kerberos.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// AuthInfo_AuthInfo_Kerberos structure represents FW_AUTH_INFO structure anonymous member.
//
// This structure contains information on the local and remote hosts that resulted from
// the authentication methods performed between them.
type AuthInfo_AuthInfo_Kerberos struct {
	// wszMyId:  A pointer to a Unicode string representing the identity of the local host
	// when a Kerberos-based authentication method, as specified in [MS-KILE], is used.
	MyID string `idl:"name:wszMyId;string" json:"my_id"`
	// wszPeerId:  A pointer to a Unicode string representing the identity of the remote
	// host when a Kerberos-based authentication method, as specified in [MS-KILE], is used.
	PeerID string `idl:"name:wszPeerId;string" json:"peer_id"`
}

func (o *AuthInfo_AuthInfo_Kerberos) xxx_PreparePayload(ctx context.Context) error {
	if len(o.MyID) > int(10001) {
		return fmt.Errorf("MyID is out of range")
	}
	if len(o.PeerID) > int(10001) {
		return fmt.Errorf("PeerID is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthInfo_AuthInfo_Kerberos) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.MyID != "" {
		_ptr_wszMyId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.MyID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.MyID, _ptr_wszMyId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PeerID != "" {
		_ptr_wszPeerId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PeerID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PeerID, _ptr_wszPeerId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AuthInfo_AuthInfo_Kerberos) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_wszMyId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.MyID); err != nil {
			return err
		}
		return nil
	})
	_s_wszMyId := func(ptr interface{}) { o.MyID = *ptr.(*string) }
	if err := w.ReadPointer(&o.MyID, _s_wszMyId, _ptr_wszMyId); err != nil {
		return err
	}
	_ptr_wszPeerId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PeerID); err != nil {
			return err
		}
		return nil
	})
	_s_wszPeerId := func(ptr interface{}) { o.PeerID = *ptr.(*string) }
	if err := w.ReadPointer(&o.PeerID, _s_wszPeerId, _ptr_wszPeerId); err != nil {
		return err
	}
	return nil
}

// Endpoints structure represents FW_ENDPOINTS RPC structure.
//
// This structure represents the two endpoints, source and destination, that participate
// in IP communication.
type Endpoints struct {
	// IpVersion:  This field specifies the Internet Protocol version used. This field MUST
	// contain a valid value from the FW_IP_VERSION enumeration.
	IPVersion IPVersion `idl:"name:IpVersion" json:"ip_version"`
	// dwSourceV4Address:  This field is the IPv4 address of the source endpoint.
	SourceV4Address uint32 `idl:"name:dwSourceV4Address" json:"source_v4_address"`
	// dwDestinationV4Address:  This field is the IPv4 address of the destination endpoint.
	DestinationV4Address uint32 `idl:"name:dwDestinationV4Address" json:"destination_v4_address"`
	// SourceV6Address:  This field is a 16-octet array that represents the IPv6 address
	// of the source endpoint.
	SourceV6Address []byte `idl:"name:SourceV6Address" json:"source_v6_address"`
	// DestinationV6Address:  This field is a 16-octet array that represents the IPv6 address
	// of the destination endpoint.
	//
	// The v4 versions or the v6 versions of the fields are used depending on the IpVersion
	// field value.
	DestinationV6Address []byte `idl:"name:DestinationV6Address" json:"destination_v6_address"`
}

func (o *Endpoints) xxx_PreparePayload(ctx context.Context) error {
	if o.IPVersion < IPVersion(1) || o.IPVersion > IPVersion(2) {
		return fmt.Errorf("IPVersion is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Endpoints) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.IPVersion)); err != nil {
		return err
	}
	if err := w.WriteData(o.SourceV4Address); err != nil {
		return err
	}
	if err := w.WriteData(o.DestinationV4Address); err != nil {
		return err
	}
	for i1 := range o.SourceV6Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.SourceV6Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SourceV6Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DestinationV6Address {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.DestinationV6Address[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DestinationV6Address); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Endpoints) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.IPVersion)); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceV4Address); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationV4Address); err != nil {
		return err
	}
	o.SourceV6Address = make([]byte, 16)
	for i1 := range o.SourceV6Address {
		i1 := i1
		if err := w.ReadData(&o.SourceV6Address[i1]); err != nil {
			return err
		}
	}
	o.DestinationV6Address = make([]byte, 16)
	for i1 := range o.DestinationV6Address {
		i1 := i1
		if err := w.ReadData(&o.DestinationV6Address[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Phase1SADetails structure represents FW_PHASE1_SA_DETAILS RPC structure.
//
// This structure represents a security association that is established after the main
// mode negotiations take place; it contains the selected algorithms to enforce IPsec
// and the methods and results of the authentication process.
type Phase1SADetails struct {
	// SaId:  A 64-bit integer that uniquely identifies the security association.
	SAID uint64 `idl:"name:SaId" json:"sa_id"`
	// KeyModuleType:  The keying protocol used, IKE or AuthIP. The field MUST contain only
	// a value from the FW_PHASE1_KEY_MODULE_TYPE enumeration.
	KeyModuleType Phase1KeyModuleType `idl:"name:KeyModuleType" json:"key_module_type"`
	// Endpoints:  This field contains IP address information of the two endpoints that
	// established this security association. An address of zero means the security association
	// applies to any endpoint.
	Endpoints *Endpoints `idl:"name:Endpoints" json:"endpoints"`
	// SelectedProposal:  This is the Phase1 cryptographic suite that was selected by the
	// negotiation of the keying protocol.
	SelectedProposal *Phase1CryptoSuite `idl:"name:SelectedProposal" json:"selected_proposal"`
	// dwProposalLifetimeKBytes:  Currently not supported.
	ProposalLifetimeKBytes uint32 `idl:"name:dwProposalLifetimeKBytes" json:"proposal_lifetime_k_bytes"`
	// dwProposalLifetimeMinutes:  This field specifies the lifetime in minutes of this
	// security association before a rekey MUST happen.
	ProposalLifetimeMinutes uint32 `idl:"name:dwProposalLifetimeMinutes" json:"proposal_lifetime_minutes"`
	// dwProposalMaxNumPhase2:  This field specifies the number of Phase2 (quick mode) negotiations
	// (rekeys) that can happen before this security association MUST be renegotiated.
	ProposalMaxNumPhase2 uint32 `idl:"name:dwProposalMaxNumPhase2" json:"proposal_max_num_phase2"`
	// CookiePair:  This value is used for diagnostics.
	CookiePair *CookiePair `idl:"name:CookiePair" json:"cookie_pair"`
	// pFirstAuth:  A pointer to an FW_AUTH_INFO structure that contains the information
	// that resulted from the method negotiated and used for first authentication. This
	// pointer MUST NOT be null.
	FirstAuth *AuthInfo `idl:"name:pFirstAuth" json:"first_auth"`
	// pSecondAuth:  A pointer to an FW_AUTH_INFO structure that contains the information
	// that resulted from the method negotiated and used for second authentication. If the
	// field is NULL, the second authentication was not performed.
	SecondAuth *AuthInfo `idl:"name:pSecondAuth" json:"second_auth"`
	// dwP1SaFlags:  Reserved value and not currently used. It MUST be set to 0.
	P1SAFlags uint32 `idl:"name:dwP1SaFlags" json:"p1_sa_flags"`
}

func (o *Phase1SADetails) xxx_PreparePayload(ctx context.Context) error {
	if o.KeyModuleType < Phase1KeyModuleType(1) || o.KeyModuleType > Phase1KeyModuleType(2) {
		return fmt.Errorf("KeyModuleType is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Phase1SADetails) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.SAID); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.KeyModuleType)); err != nil {
		return err
	}
	if o.Endpoints != nil {
		if err := o.Endpoints.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Endpoints{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SelectedProposal != nil {
		if err := o.SelectedProposal.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Phase1CryptoSuite{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProposalLifetimeKBytes); err != nil {
		return err
	}
	if err := w.WriteData(o.ProposalLifetimeMinutes); err != nil {
		return err
	}
	if err := w.WriteData(o.ProposalMaxNumPhase2); err != nil {
		return err
	}
	if o.CookiePair != nil {
		if err := o.CookiePair.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CookiePair{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FirstAuth != nil {
		_ptr_pFirstAuth := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.FirstAuth != nil {
				if err := o.FirstAuth.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AuthInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.FirstAuth, _ptr_pFirstAuth); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SecondAuth != nil {
		_ptr_pSecondAuth := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SecondAuth != nil {
				if err := o.SecondAuth.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&AuthInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecondAuth, _ptr_pSecondAuth); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.P1SAFlags); err != nil {
		return err
	}
	return nil
}
func (o *Phase1SADetails) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.SAID); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.KeyModuleType)); err != nil {
		return err
	}
	if o.Endpoints == nil {
		o.Endpoints = &Endpoints{}
	}
	if err := o.Endpoints.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SelectedProposal == nil {
		o.SelectedProposal = &Phase1CryptoSuite{}
	}
	if err := o.SelectedProposal.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProposalLifetimeKBytes); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProposalLifetimeMinutes); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProposalMaxNumPhase2); err != nil {
		return err
	}
	if o.CookiePair == nil {
		o.CookiePair = &CookiePair{}
	}
	if err := o.CookiePair.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pFirstAuth := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.FirstAuth == nil {
			o.FirstAuth = &AuthInfo{}
		}
		if err := o.FirstAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pFirstAuth := func(ptr interface{}) { o.FirstAuth = *ptr.(**AuthInfo) }
	if err := w.ReadPointer(&o.FirstAuth, _s_pFirstAuth, _ptr_pFirstAuth); err != nil {
		return err
	}
	_ptr_pSecondAuth := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SecondAuth == nil {
			o.SecondAuth = &AuthInfo{}
		}
		if err := o.SecondAuth.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pSecondAuth := func(ptr interface{}) { o.SecondAuth = *ptr.(**AuthInfo) }
	if err := w.ReadPointer(&o.SecondAuth, _s_pSecondAuth, _ptr_pSecondAuth); err != nil {
		return err
	}
	if err := w.ReadData(&o.P1SAFlags); err != nil {
		return err
	}
	return nil
}

// Phase2TrafficType type represents FW_PHASE2_TRAFFIC_TYPE RPC enumeration.
//
// This enumeration identifies the two types of traffic enforcement modes that IPsec
// supports. It is defined in the IDL for future use.
type Phase2TrafficType uint16

var (
	// FW_PHASE2_TRAFFIC_TYPE_INVALID:  This value MUST NOT be used. It is defined for
	// simplicity in writing IDL definitions and code. This symbolic constant has a value
	// of 0.
	Phase2TrafficTypeInvalid Phase2TrafficType = 0
	// FW_PHASE2_TRAFFIC_TYPE_TRANSPORT:  This value represents IPsec transport mode, which
	// happens directly between two endpoints. This symbolic constant has a value of 1.
	Phase2TrafficTypeTransport Phase2TrafficType = 1
	// FW_PHASE2_TRAFFIC_TYPE_TUNNEL:  This value represents IPsec tunnel mode, which uses
	// two other endpoints to tunnel through them when the original endpoints communicate.
	// This symbolic constant has a value of 2.
	Phase2TrafficTypeTunnel Phase2TrafficType = 2
	// FW_PHASE2_TRAFFIC_TYPE_MAX:  This value and values that exceed this value are not
	// valid and MUST NOT be used. It is defined for simplicity in writing IDL definitions
	// and code. This symbolic constant has a value of 3.
	Phase2TrafficTypeMax Phase2TrafficType = 3
)

func (o Phase2TrafficType) String() string {
	switch o {
	case Phase2TrafficTypeInvalid:
		return "Phase2TrafficTypeInvalid"
	case Phase2TrafficTypeTransport:
		return "Phase2TrafficTypeTransport"
	case Phase2TrafficTypeTunnel:
		return "Phase2TrafficTypeTunnel"
	case Phase2TrafficTypeMax:
		return "Phase2TrafficTypeMax"
	}
	return "Invalid"
}

// Phase2SADetails structure represents FW_PHASE2_SA_DETAILS RPC structure.
//
// This structure represents a security association that is established after the quick
// mode negotiations take place; it contains the selected algorithms to enforce IPsec.
type Phase2SADetails struct {
	// SaId:  A 64-bit integer number that uniquely identifies the security association.
	SAID uint64 `idl:"name:SaId" json:"sa_id"`
	// Direction:  This field specifies the direction of the traffic this security association
	// is securing.
	Direction Direction `idl:"name:Direction" json:"direction"`
	// Endpoints:  This field contains IP address information of the two endpoints that
	// established this security association. An address of zero means the security association
	// applies to any endpoint.
	Endpoints *Endpoints `idl:"name:Endpoints" json:"endpoints"`
	// wLocalPort:  This field specifies the port of the local endpoint that is used in
	// the traffic secured by this security association. A value of 0 specifies any port.
	LocalPort uint16 `idl:"name:wLocalPort" json:"local_port"`
	// wRemotePort:  This field specifies the port of the remote endpoint that is used in
	// the traffic secured by this security association. A value of 0 specifies any port.
	RemotePort uint16 `idl:"name:wRemotePort" json:"remote_port"`
	// wIpProtocol:  This field specifies the protocol of the traffic secured by this security
	// association. If the value is within the range 0 to 255, the value describes a protocol
	// as in IETF IANA numbers (for more information, see [IANA-PROTO-NUM]). If the value
	// is 256, the rule matches ANY protocol.
	IPProtocol uint16 `idl:"name:wIpProtocol" json:"ip_protocol"`
	// SelectedProposal:  This field contains the Phase2 cryptographic suite selected by
	// the negotiation that is used by this security association to enforce IPsec.
	SelectedProposal *Phase2CryptoSuite `idl:"name:SelectedProposal" json:"selected_proposal"`
	// Pfs:  This field specifies the perfect forward secrecy used by this security association.
	PFS Phase2CryptoPFS `idl:"name:Pfs" json:"pfs"`
	// TransportFilterId:  This GUID MAY contain additional implementation-specific<21>
	// information about the security association. The client MUST ignore this value.
	TransportFilterID *dtyp.GUID `idl:"name:TransportFilterId" json:"transport_filter_id"`
	// dwP2SaFlags:  Reserved value and not currently used. It MUST be set to 0.
	P2SAFlags uint32 `idl:"name:dwP2SaFlags" json:"p2_sa_flags"`
}

func (o *Phase2SADetails) xxx_PreparePayload(ctx context.Context) error {
	if o.Direction < Direction(1) || o.Direction > Direction(2) {
		return fmt.Errorf("Direction is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Phase2SADetails) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.SAID); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Direction)); err != nil {
		return err
	}
	if o.Endpoints != nil {
		if err := o.Endpoints.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Endpoints{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocalPort); err != nil {
		return err
	}
	if err := w.WriteData(o.RemotePort); err != nil {
		return err
	}
	if err := w.WriteData(o.IPProtocol); err != nil {
		return err
	}
	if o.SelectedProposal != nil {
		if err := o.SelectedProposal.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Phase2CryptoSuite{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.PFS)); err != nil {
		return err
	}
	if o.TransportFilterID != nil {
		if err := o.TransportFilterID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.P2SAFlags); err != nil {
		return err
	}
	return nil
}
func (o *Phase2SADetails) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.SAID); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Direction)); err != nil {
		return err
	}
	if o.Endpoints == nil {
		o.Endpoints = &Endpoints{}
	}
	if err := o.Endpoints.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalPort); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemotePort); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPProtocol); err != nil {
		return err
	}
	if o.SelectedProposal == nil {
		o.SelectedProposal = &Phase2CryptoSuite{}
	}
	if err := o.SelectedProposal.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.PFS)); err != nil {
		return err
	}
	if o.TransportFilterID == nil {
		o.TransportFilterID = &dtyp.GUID{}
	}
	if err := o.TransportFilterID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.P2SAFlags); err != nil {
		return err
	}
	return nil
}

// ProfileConfigValue structure represents FW_PROFILE_CONFIG_VALUE RPC union.
type ProfileConfigValue struct {
	// Types that are assignable to Value
	//
	// *ProfileConfigValue_String
	// *ProfileConfigValue_DisabledInterfaces
	// *ProfileConfigValue_Value
	Value is_ProfileConfigValue `json:"value"`
}

func (o *ProfileConfigValue) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ProfileConfigValue_String:
		if value != nil {
			return value.String
		}
	case *ProfileConfigValue_DisabledInterfaces:
		if value != nil {
			return value.DisabledInterfaces
		}
	case *ProfileConfigValue_Value:
		if value != nil {
			return value.Value
		}
	}
	return nil
}

type is_ProfileConfigValue interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ProfileConfigValue()
}

func (o *ProfileConfigValue) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ProfileConfigValue_String:
		return uint16(9)
	case *ProfileConfigValue_DisabledInterfaces:
		return uint16(15)
	case *ProfileConfigValue_Value:
		switch sw {
		case uint16(1),
			uint16(2),
			uint16(3),
			uint16(4),
			uint16(5),
			uint16(6),
			uint16(7),
			uint16(8),
			uint16(10),
			uint16(11),
			uint16(12),
			uint16(13),
			uint16(14),
			uint16(16),
			uint16(17),
			uint16(18):
			return sw
		}
		return uint16(1)
	}
	return uint16(0)
}

func (o *ProfileConfigValue) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(9):
		_o, _ := o.Value.(*ProfileConfigValue_String)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProfileConfigValue_String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(15):
		_o, _ := o.Value.(*ProfileConfigValue_DisabledInterfaces)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ProfileConfigValue_DisabledInterfaces{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1),
		uint16(2),
		uint16(3),
		uint16(4),
		uint16(5),
		uint16(6),
		uint16(7),
		uint16(8),
		uint16(10),
		uint16(11),
		uint16(12),
		uint16(13),
		uint16(14),
		uint16(16),
		uint16(17),
		uint16(18):
		_o, _ := o.Value.(*ProfileConfigValue_Value)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ProfileConfigValue) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(9):
		o.Value = &ProfileConfigValue_String{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(15):
		o.Value = &ProfileConfigValue_DisabledInterfaces{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1),
		uint16(2),
		uint16(3),
		uint16(4),
		uint16(5),
		uint16(6),
		uint16(7),
		uint16(8),
		uint16(10),
		uint16(11),
		uint16(12),
		uint16(13),
		uint16(14),
		uint16(16),
		uint16(17),
		uint16(18):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &ProfileConfigValue_Value{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**ProfileConfigValue_Value) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ProfileConfigValue_String structure represents FW_PROFILE_CONFIG_VALUE RPC union arm.
//
// It has following labels: 9
type ProfileConfigValue_String struct {
	String string `idl:"name:wszStr;string" json:"string"`
}

func (*ProfileConfigValue_String) is_ProfileConfigValue() {}

func (o *ProfileConfigValue_String) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.String != "" {
		_ptr_wszStr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.String); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_wszStr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProfileConfigValue_String) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_wszStr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.String); err != nil {
			return err
		}
		return nil
	})
	_s_wszStr := func(ptr interface{}) { o.String = *ptr.(*string) }
	if err := w.ReadPointer(&o.String, _s_wszStr, _ptr_wszStr); err != nil {
		return err
	}
	return nil
}

// ProfileConfigValue_DisabledInterfaces structure represents FW_PROFILE_CONFIG_VALUE RPC union arm.
//
// It has following labels: 15
type ProfileConfigValue_DisabledInterfaces struct {
	DisabledInterfaces *InterfaceLUIDs `idl:"name:pDisabledInterfaces" json:"disabled_interfaces"`
}

func (*ProfileConfigValue_DisabledInterfaces) is_ProfileConfigValue() {}

func (o *ProfileConfigValue_DisabledInterfaces) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DisabledInterfaces != nil {
		_ptr_pDisabledInterfaces := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DisabledInterfaces != nil {
				if err := o.DisabledInterfaces.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&InterfaceLUIDs{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DisabledInterfaces, _ptr_pDisabledInterfaces); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ProfileConfigValue_DisabledInterfaces) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pDisabledInterfaces := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DisabledInterfaces == nil {
			o.DisabledInterfaces = &InterfaceLUIDs{}
		}
		if err := o.DisabledInterfaces.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pDisabledInterfaces := func(ptr interface{}) { o.DisabledInterfaces = *ptr.(**InterfaceLUIDs) }
	if err := w.ReadPointer(&o.DisabledInterfaces, _s_pDisabledInterfaces, _ptr_pDisabledInterfaces); err != nil {
		return err
	}
	return nil
}

// ProfileConfigValue_Value structure represents FW_PROFILE_CONFIG_VALUE RPC union arm.
//
// It has following labels: 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 16, 17, 18
type ProfileConfigValue_Value struct {
	Value uint32 `idl:"name:pdwVal" json:"value"`
}

func (*ProfileConfigValue_Value) is_ProfileConfigValue() {}

func (o *ProfileConfigValue_Value) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Value); err != nil {
		return err
	}
	return nil
}
func (o *ProfileConfigValue_Value) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Value); err != nil {
		return err
	}
	return nil
}

// MMRule structure represents FW_MM_RULE RPC structure.
//
// This structure is used to represent a main mode rule.
type MMRule struct {
	// pNext:  A pointer to the next FW_MM_RULE in the list.
	Next *MMRule `idl:"name:pNext" json:"next"`
	// wSchemaVersion:  Specifies the version of the rule.
	SchemaVersion uint16 `idl:"name:wSchemaVersion" json:"schema_version"`
	// wszRuleId:  A pointer to a Unicode string that uniquely identifies the rule.
	RuleID string `idl:"name:wszRuleId;string;pointer:ref" json:"rule_id"`
	// wszName:  A pointer to a Unicode string that provides a friendly name for the rule.
	Name string `idl:"name:wszName;string" json:"name"`
	// wszDescription:  A pointer to a Unicode string that provides a friendly description
	// for the rule.
	Description string `idl:"name:wszDescription;string" json:"description"`
	// dwProfiles:  A bitmask of the FW_PROFILE_TYPE flags. It is a condition that matches
	// traffic on the specified profiles.
	Profiles uint32 `idl:"name:dwProfiles" json:"profiles"`
	// Endpoint1:  A condition that specifies the addresses of the first host of the traffic
	// that the rule matches. An empty EndPoint1 structure means this condition is not applied
	// (no match).
	Endpoint1 *Addresses `idl:"name:Endpoint1" json:"endpoint1"`
	// Endpoint2:  A condition that specifies the addresses of the second host of the traffic
	// that the rule matches. An empty EndPoint2 structure means this condition is not applied
	// (no match).
	Endpoint2 *Addresses `idl:"name:Endpoint2" json:"endpoint2"`
	// wszPhase1AuthSet:  A Unicode string that represents the set identifier of a Phase1
	// authentication sets policy objects.
	Phase1AuthSet   string `idl:"name:wszPhase1AuthSet;string" json:"phase1_auth_set"`
	Phase1CryptoSet string `idl:"name:wszPhase1CryptoSet;string" json:"phase1_crypto_set"`
	// wFlags:  Bit flags from FW_CS_RULE_FLAGS.
	Flags uint16 `idl:"name:wFlags" json:"flags"`
	// wszEmbeddedContext:  A pointer to a Unicode string that specifies a group name for
	// this rule. Other components in the system use this string to enable or disable a
	// group of rules by verifying that all rules have the same group name.
	EmbeddedContext string `idl:"name:wszEmbeddedContext;string" json:"embedded_context"`
	// PlatformValidityList:  A condition in a rule that determines whether or not the rule
	// is enforced by the local computer based on the local computer's platform information.
	// The rule is enforced only if the local computer's operating system platform is an
	// element of the set described by PlatformValidityList.<22>
	PlatformValidityList *OSPlatformList `idl:"name:PlatformValidityList" json:"platform_validity_list"`
	// Origin:  This field is the rule origin, as specified in the FW_RULE_ORIGIN_TYPE enumeration.
	// It MUST be filled on enumerated rules and ignored on input.
	Origin RuleOriginType `idl:"name:Origin" json:"origin"`
	// wszGPOName:  A pointer to a Unicode string containing the displayName of the GPO
	// containing this object. When adding a new object, this field is not used. The client
	// SHOULD set the value to NULL, and the server MUST ignore the value. When enumerating
	// an existing object, if the client does not set the FW_ENUM_RULES_FLAG_RESOLVE_GPO_NAME
	// flag, the server MUST set the value to NULL. Otherwise, the server MUST set the value
	// to the displayName of the GPO containing the object or NULL if the object is not
	// contained within a GPO. For details about how the server initializes an object from
	// a GPO, see section 3.1.3. For details about how the displayName of a GPO is stored,
	// see [MS-GPOL] section 2.3.
	GPOName string `idl:"name:wszGPOName;string" json:"gpo_name"`
	// Status:  The status code of the rule, as specified by the FW_RULE_STATUS enumeration.
	// This field is filled out when the structure is returned as output. On input, this
	// field MUST be set to FW_RULE_STATUS_OK.
	Status           RuleStatus `idl:"name:Status" json:"status"`
	MetadataReserved uint32     `idl:"name:MetaDataReserved" json:"metadata_reserved"`
	// pMetaData:  A pointer to an FW_OBJECT_METADATA structure that contains specific metadata
	// about the current state of the connection security rule.
	Metadata []*ObjectMetadata `idl:"name:pMetaData;size_is:(((MetaDataReserved&1)?1:0))" json:"metadata"`
}

func (o *MMRule) xxx_PreparePayload(ctx context.Context) error {
	if o.Metadata != nil && o.MetadataReserved == 0 {
		o.MetadataReserved = uint32(len(o.Metadata))
	}
	if len(o.RuleID) > int(512) {
		return fmt.Errorf("RuleID is out of range")
	}
	if len(o.Name) > int(10001) {
		return fmt.Errorf("Name is out of range")
	}
	if len(o.Description) > int(10001) {
		return fmt.Errorf("Description is out of range")
	}
	if len(o.Phase1AuthSet) > int(255) {
		return fmt.Errorf("Phase1AuthSet is out of range")
	}
	if len(o.Phase1CryptoSet) > int(255) {
		return fmt.Errorf("Phase1CryptoSet is out of range")
	}
	if len(o.EmbeddedContext) > int(10001) {
		return fmt.Errorf("EmbeddedContext is out of range")
	}
	if o.Origin > RuleOriginType(5) {
		return fmt.Errorf("Origin is out of range")
	}
	if len(o.GPOName) > int(10001) {
		return fmt.Errorf("GPOName is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MMRule) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Next != nil {
		_ptr_pNext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Next != nil {
				if err := o.Next.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&MMRule{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Next, _ptr_pNext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if o.RuleID != "" {
		_ptr_wszRuleId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RuleID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RuleID, _ptr_wszRuleId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Name != "" {
		_ptr_wszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_wszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_wszDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_wszDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Profiles); err != nil {
		return err
	}
	if o.Endpoint1 != nil {
		if err := o.Endpoint1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Endpoint2 != nil {
		if err := o.Endpoint2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Addresses{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Phase1AuthSet != "" {
		_ptr_wszPhase1AuthSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase1AuthSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase1AuthSet, _ptr_wszPhase1AuthSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Phase1CryptoSet != "" {
		_ptr_wszPhase1CryptoSet := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Phase1CryptoSet); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Phase1CryptoSet, _ptr_wszPhase1CryptoSet); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.EmbeddedContext != "" {
		_ptr_wszEmbeddedContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EmbeddedContext); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PlatformValidityList != nil {
		if err := o.PlatformValidityList.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OSPlatformList{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint16(o.Origin)); err != nil {
		return err
	}
	if o.GPOName != "" {
		_ptr_wszGPOName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.GPOName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.GPOName, _ptr_wszGPOName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteData(o.MetadataReserved); err != nil {
		return err
	}
	_exprMetaDataReserved := uint32(0)
	if (o.MetadataReserved & 1) != 0 {
		_exprMetaDataReserved = uint32(1)
	} else {
		_exprMetaDataReserved = uint32(0)
	}
	if o.Metadata != nil || _exprMetaDataReserved > 0 {
		_ptr_pMetaData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_exprMetaDataReserved := uint64(0)
			if (o.MetadataReserved & 1) != 0 {
				_exprMetaDataReserved = uint64(1)
			} else {
				_exprMetaDataReserved = uint64(0)
			}
			dimSize1 := uint64(_exprMetaDataReserved)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Metadata {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Metadata[i1] != nil {
					if err := o.Metadata[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Metadata); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ObjectMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Metadata, _ptr_pMetaData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MMRule) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pNext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Next == nil {
			o.Next = &MMRule{}
		}
		if err := o.Next.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pNext := func(ptr interface{}) { o.Next = *ptr.(**MMRule) }
	if err := w.ReadPointer(&o.Next, _s_pNext, _ptr_pNext); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	_ptr_wszRuleId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RuleID); err != nil {
			return err
		}
		return nil
	})
	_s_wszRuleId := func(ptr interface{}) { o.RuleID = *ptr.(*string) }
	if err := w.ReadPointer(&o.RuleID, _s_wszRuleId, _ptr_wszRuleId); err != nil {
		return err
	}
	_ptr_wszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_wszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_wszName, _ptr_wszName); err != nil {
		return err
	}
	_ptr_wszDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_wszDescription := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_wszDescription, _ptr_wszDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.Profiles); err != nil {
		return err
	}
	if o.Endpoint1 == nil {
		o.Endpoint1 = &Addresses{}
	}
	if err := o.Endpoint1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Endpoint2 == nil {
		o.Endpoint2 = &Addresses{}
	}
	if err := o.Endpoint2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_wszPhase1AuthSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase1AuthSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase1AuthSet := func(ptr interface{}) { o.Phase1AuthSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase1AuthSet, _s_wszPhase1AuthSet, _ptr_wszPhase1AuthSet); err != nil {
		return err
	}
	_ptr_wszPhase1CryptoSet := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phase1CryptoSet); err != nil {
			return err
		}
		return nil
	})
	_s_wszPhase1CryptoSet := func(ptr interface{}) { o.Phase1CryptoSet = *ptr.(*string) }
	if err := w.ReadPointer(&o.Phase1CryptoSet, _s_wszPhase1CryptoSet, _ptr_wszPhase1CryptoSet); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_wszEmbeddedContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EmbeddedContext); err != nil {
			return err
		}
		return nil
	})
	_s_wszEmbeddedContext := func(ptr interface{}) { o.EmbeddedContext = *ptr.(*string) }
	if err := w.ReadPointer(&o.EmbeddedContext, _s_wszEmbeddedContext, _ptr_wszEmbeddedContext); err != nil {
		return err
	}
	if o.PlatformValidityList == nil {
		o.PlatformValidityList = &OSPlatformList{}
	}
	if err := o.PlatformValidityList.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Origin)); err != nil {
		return err
	}
	_ptr_wszGPOName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.GPOName); err != nil {
			return err
		}
		return nil
	})
	_s_wszGPOName := func(ptr interface{}) { o.GPOName = *ptr.(*string) }
	if err := w.ReadPointer(&o.GPOName, _s_wszGPOName, _ptr_wszGPOName); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataReserved); err != nil {
		return err
	}
	_ptr_pMetaData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		_exprpMetaData := uint64(0)
		if (o.MetadataReserved & 1) != 0 {
			_exprpMetaData = uint64(1)
		} else {
			_exprpMetaData = uint64(0)
		}
		if _exprpMetaData > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(_exprpMetaData)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Metadata", sizeInfo[0])
		}
		o.Metadata = make([]*ObjectMetadata, sizeInfo[0])
		for i1 := range o.Metadata {
			i1 := i1
			if o.Metadata[i1] == nil {
				o.Metadata[i1] = &ObjectMetadata{}
			}
			if err := o.Metadata[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMetaData := func(ptr interface{}) { o.Metadata = *ptr.(*[]*ObjectMetadata) }
	if err := w.ReadPointer(&o.Metadata, _s_pMetaData, _ptr_pMetaData); err != nil {
		return err
	}
	return nil
}

// MatchKey type represents FW_MATCH_KEY RPC enumeration.
//
// This enumeration describes the keys that a query is allowed to match.
type MatchKey uint16

var (
	// FW_MATCH_KEY_PROFILE:  This key matches the profile conditions of the queried object.
	// This symbolic constant has a value of 0.
	MatchKeyProfile MatchKey = 0
	// FW_MATCH_KEY_STATUS:  This key matches the status conditions of the queried object.
	// This symbolic constant has a value of 1.
	MatchKeyStatus MatchKey = 1
	// FW_MATCH_KEY_OBJECTID:  This key matches the object ID (rule ID or set ID) of the
	// queried object. This symbolic constant has a value of 2.
	MatchKeyObjectID MatchKey = 2
	// FW_MATCH_KEY_FILTERID:  This value is not used on the wire. This symbolic constant
	// has a value of 3.
	MatchKeyFilterID MatchKey = 3
	// FW_MATCH_KEY_APP_PATH:  This key matches the application condition of the queried
	// object. This symbolic constant has a value of 4.
	MatchKeyAppPath MatchKey = 4
	// FW_MATCH_KEY_PROTOCOL:  This key matches the protocol condition of the queried object.
	// This symbolic constant has a value of 5.
	MatchKeyProtocol MatchKey = 5
	// FW_MATCH_KEY_LOCAL_PORT:  This key matches the TCP or UDP local port condition of
	// the queried object. This symbolic constant has a value of 6.
	MatchKeyLocalPort MatchKey = 6
	// FW_MATCH_KEY_REMOTE_PORT:  This key matches the TCP or UDP remote port condition
	// of the queried object. This symbolic constant has a value of 7.
	MatchKeyRemotePort MatchKey = 7
	// FW_MATCH_KEY_GROUP:  This key matches the group name (the Embedded context field)
	// of the queried object. This symbolic constant has a value of 8.
	MatchKeyGroup MatchKey = 8
	// FW_MATCH_KEY_SVC_NAME:  This key matches the service name condition of the queried
	// object. This symbolic constant has a value of 9.
	MatchKeyServiceName MatchKey = 9
	// FW_MATCH_KEY_DIRECTION:  This key matches the direction condition of the queried
	// object. This symbolic constant has a value of 10.
	MatchKeyDirection MatchKey = 10
	// FW_MATCH_KEY_LOCAL_USER_OWNER:  This key matches the local user owner condition
	// of the queried object. For schema versions 0x0200, 0x0201, and 0x020A, this value
	// is invalid and MUST NOT be used. This symbolic constant has a value of 11.
	MatchKeyLocalUserOwner MatchKey = 11
	// FW_MATCH_KEY_PACKAGE_ID:  This key matches the package ID condition of the queried
	// object. For schema versions 0x0200, 0x0201, and 0x020A, this value is invalid and
	// MUST NOT be used. This symbolic constant has a value of 12.
	MatchKeyPackageID MatchKey = 12
	// FW_MATCH_KEY_FQBN:  This key matches the fully qualified binary name (FQBN) condition
	// of the queried object. For schema versions 0x0200 through 0x021A, this value is invalid
	// and MUST NOT be used. This symbolic constant has a value of 13.
	MatchKeyFQBN MatchKey = 13
	// FW_MATCH_KEY_COMPARTMENT_ID: This key matches the compartment ID condition of the
	// queried object. For schema versions 0x0200 through 0x021A, this value is invalid
	// and MUST NOT be used. This symbolic constant has a value of 14.
	MatchKeyCompartmentID MatchKey = 14
	// FW_MATCH_KEY_MAX:  This value and values that exceed this value are not valid and
	// MUST NOT be used. It is defined for simplicity in writing IDL definitions and code.
	// This symbolic constant has a value of 15.
	MatchKeyMax MatchKey = 15
)

func (o MatchKey) String() string {
	switch o {
	case MatchKeyProfile:
		return "MatchKeyProfile"
	case MatchKeyStatus:
		return "MatchKeyStatus"
	case MatchKeyObjectID:
		return "MatchKeyObjectID"
	case MatchKeyFilterID:
		return "MatchKeyFilterID"
	case MatchKeyAppPath:
		return "MatchKeyAppPath"
	case MatchKeyProtocol:
		return "MatchKeyProtocol"
	case MatchKeyLocalPort:
		return "MatchKeyLocalPort"
	case MatchKeyRemotePort:
		return "MatchKeyRemotePort"
	case MatchKeyGroup:
		return "MatchKeyGroup"
	case MatchKeyServiceName:
		return "MatchKeyServiceName"
	case MatchKeyDirection:
		return "MatchKeyDirection"
	case MatchKeyLocalUserOwner:
		return "MatchKeyLocalUserOwner"
	case MatchKeyPackageID:
		return "MatchKeyPackageID"
	case MatchKeyFQBN:
		return "MatchKeyFQBN"
	case MatchKeyCompartmentID:
		return "MatchKeyCompartmentID"
	case MatchKeyMax:
		return "MatchKeyMax"
	}
	return "Invalid"
}

// DataType type represents FW_DATA_TYPE RPC enumeration.
//
// This enumeration describes the data types that this protocol uses in generic structures.
// It is currently used only in section 2.2.89.
type DataType uint16

var (
	// FW_DATA_TYPE_EMPTY:  The value SHOULD be empty and not used. This symbolic constant
	// has a value of zero.
	DataTypeEmpty DataType = 0
	// FW_DATA_TYPE_UINT8:  This data type is a UINT8, which is an 8-bit unsigned integer.
	// This symbolic constant has a value of 1.
	DataTypeUint8 DataType = 1
	// FW_DATA_TYPE_UINT16:  This data type is a UINT16, which is a 16-bit unsigned integer.
	// This symbolic constant has a value of 2.
	DataTypeUint16 DataType = 2
	// FW_DATA_TYPE_UINT32:  This data type is a UINT32, which is a 32-bit unsigned integer.
	// This symbolic constant has a value of 3.
	DataTypeUint32 DataType = 3
	// FW_DATA_TYPE_UINT64:  This data type is a UINT64, which is a 64-bit unsigned integer.
	// This symbolic constant has a value of 4.
	DataTypeUint64 DataType = 4
	// FW_DATA_TYPE_UNICODE_STRING:  This data type is a Unicode string. This symbolic
	// constant has a value of 5.
	DataTypeUnicodeString DataType = 5
)

func (o DataType) String() string {
	switch o {
	case DataTypeEmpty:
		return "DataTypeEmpty"
	case DataTypeUint8:
		return "DataTypeUint8"
	case DataTypeUint16:
		return "DataTypeUint16"
	case DataTypeUint32:
		return "DataTypeUint32"
	case DataTypeUint64:
		return "DataTypeUint64"
	case DataTypeUnicodeString:
		return "DataTypeUnicodeString"
	}
	return "Invalid"
}

// MatchValue structure represents FW_MATCH_VALUE RPC structure.
//
// This structure is used to generically store different data types.
type MatchValue struct {
	// type:  This field identifies the data type that is stored in the structure.
	Type       DataType               `idl:"name:type" json:"type"`
	MatchValue *MatchValue_MatchValue `idl:"name:MatchValue;switch_is:type" json:"match_value"`
}

func (o *MatchValue) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MatchValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.Type)); err != nil {
		return err
	}
	_swMatchValue := uint16(o.Type)
	if o.MatchValue != nil {
		if err := o.MatchValue.MarshalUnionNDR(ctx, w, _swMatchValue); err != nil {
			return err
		}
	} else {
		if err := (&MatchValue_MatchValue{}).MarshalUnionNDR(ctx, w, _swMatchValue); err != nil {
			return err
		}
	}
	return nil
}
func (o *MatchValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.Type)); err != nil {
		return err
	}
	if o.MatchValue == nil {
		o.MatchValue = &MatchValue_MatchValue{}
	}
	_swMatchValue := uint16(o.Type)
	if err := o.MatchValue.UnmarshalUnionNDR(ctx, w, _swMatchValue); err != nil {
		return err
	}
	return nil
}

// MatchValue_MatchValue structure represents FW_MATCH_VALUE union anonymous member.
//
// This structure is used to generically store different data types.
type MatchValue_MatchValue struct {
	// Types that are assignable to Value
	//
	// *MatchValue_Int8
	// *MatchValue_Int16
	// *MatchValue_Uint32
	// *MatchValue_Uint64
	// *MatchValue_UnicodeString
	// *MatchValue_DataTypeEmpty
	Value is_MatchValue_MatchValue `json:"value"`
}

func (o *MatchValue_MatchValue) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *MatchValue_Int8:
		if value != nil {
			return value.Int8
		}
	case *MatchValue_Int16:
		if value != nil {
			return value.Int16
		}
	case *MatchValue_Uint32:
		if value != nil {
			return value.Uint32
		}
	case *MatchValue_Uint64:
		if value != nil {
			return value.Uint64
		}
	case *MatchValue_UnicodeString:
		if value != nil {
			return value.UnicodeString
		}
	}
	return nil
}

type is_MatchValue_MatchValue interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_MatchValue_MatchValue()
}

func (o *MatchValue_MatchValue) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *MatchValue_Int8:
		return uint16(1)
	case *MatchValue_Int16:
		return uint16(2)
	case *MatchValue_Uint32:
		return uint16(3)
	case *MatchValue_Uint64:
		return uint16(4)
	case *MatchValue_UnicodeString:
		return uint16(5)
	case *MatchValue_DataTypeEmpty:
		return uint16(0)
	}
	return uint16(0)
}

func (o *MatchValue_MatchValue) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*MatchValue_Int8)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MatchValue_Int8{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*MatchValue_Int16)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MatchValue_Int16{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*MatchValue_Uint32)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MatchValue_Uint32{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*MatchValue_Uint64)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MatchValue_Uint64{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5):
		_o, _ := o.Value.(*MatchValue_UnicodeString)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&MatchValue_UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(0):
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *MatchValue_MatchValue) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &MatchValue_Int8{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &MatchValue_Int16{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &MatchValue_Uint32{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &MatchValue_Uint64{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5):
		o.Value = &MatchValue_UnicodeString{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(0):
		o.Value = &MatchValue_DataTypeEmpty{}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// MatchValue_Int8 structure represents MatchValue_MatchValue RPC union arm.
//
// It has following labels: 1
type MatchValue_Int8 struct {
	// uInt8:  This field contains an 8-bit unsigned integer.
	Int8 uint8 `idl:"name:uInt8" json:"int8"`
}

func (*MatchValue_Int8) is_MatchValue_MatchValue() {}

func (o *MatchValue_Int8) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Int8); err != nil {
		return err
	}
	return nil
}
func (o *MatchValue_Int8) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Int8); err != nil {
		return err
	}
	return nil
}

// MatchValue_Int16 structure represents MatchValue_MatchValue RPC union arm.
//
// It has following labels: 2
type MatchValue_Int16 struct {
	// uInt16:  This field contains a 16-bit unsigned integer.
	Int16 uint16 `idl:"name:uInt16" json:"int16"`
}

func (*MatchValue_Int16) is_MatchValue_MatchValue() {}

func (o *MatchValue_Int16) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Int16); err != nil {
		return err
	}
	return nil
}
func (o *MatchValue_Int16) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Int16); err != nil {
		return err
	}
	return nil
}

// MatchValue_Uint32 structure represents MatchValue_MatchValue RPC union arm.
//
// It has following labels: 3
type MatchValue_Uint32 struct {
	// uInt32:  This field contains a 32-bit unsigned integer.
	Uint32 uint32 `idl:"name:uInt32" json:"uint32"`
}

func (*MatchValue_Uint32) is_MatchValue_MatchValue() {}

func (o *MatchValue_Uint32) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Uint32); err != nil {
		return err
	}
	return nil
}
func (o *MatchValue_Uint32) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Uint32); err != nil {
		return err
	}
	return nil
}

// MatchValue_Uint64 structure represents MatchValue_MatchValue RPC union arm.
//
// It has following labels: 4
type MatchValue_Uint64 struct {
	// uInt64:  This field contains a 64-bit unsigned integer.
	Uint64 uint64 `idl:"name:uInt64" json:"uint64"`
}

func (*MatchValue_Uint64) is_MatchValue_MatchValue() {}

func (o *MatchValue_Uint64) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Uint64); err != nil {
		return err
	}
	return nil
}
func (o *MatchValue_Uint64) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Uint64); err != nil {
		return err
	}
	return nil
}

// MatchValue_UnicodeString structure represents MatchValue_MatchValue RPC union arm.
//
// It has following labels: 5
type MatchValue_UnicodeString struct {
	UnicodeString *MatchValue_MatchValue_UnicodeString `idl:"name:UnicodeString" json:"unicode_string"`
}

func (*MatchValue_UnicodeString) is_MatchValue_MatchValue() {}

func (o *MatchValue_UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UnicodeString != nil {
		if err := o.UnicodeString.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MatchValue_MatchValue_UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *MatchValue_UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UnicodeString == nil {
		o.UnicodeString = &MatchValue_MatchValue_UnicodeString{}
	}
	if err := o.UnicodeString.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MatchValue_MatchValue_UnicodeString structure represents FW_MATCH_VALUE structure anonymous member.
//
// This structure is used to generically store different data types.
type MatchValue_MatchValue_UnicodeString struct {
	// wszString:  This field contains a pointer to a Unicode string.
	String string `idl:"name:wszString;string" json:"string"`
}

func (o *MatchValue_MatchValue_UnicodeString) xxx_PreparePayload(ctx context.Context) error {
	if len(o.String) > int(10001) {
		return fmt.Errorf("String is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MatchValue_MatchValue_UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.String != "" {
		_ptr_wszString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.String); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_wszString); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MatchValue_MatchValue_UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_wszString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.String); err != nil {
			return err
		}
		return nil
	})
	_s_wszString := func(ptr interface{}) { o.String = *ptr.(*string) }
	if err := w.ReadPointer(&o.String, _s_wszString, _ptr_wszString); err != nil {
		return err
	}
	return nil
}

// MatchValue_DataTypeEmpty structure represents MatchValue_MatchValue RPC union arm.
//
// It has following labels: 0
type MatchValue_DataTypeEmpty struct {
}

func (*MatchValue_DataTypeEmpty) is_MatchValue_MatchValue() {}

func (o *MatchValue_DataTypeEmpty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *MatchValue_DataTypeEmpty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// MatchType type represents FW_MATCH_TYPE RPC enumeration.
//
// This enumeration specifies how a match key is matched against an object.
type MatchType uint16

var (
	// FW_MATCH_TYPE_TRAFFIC_MATCH:  The match operation evaluates to TRUE for all objects
	// that match the network traffic that is represented by the value matched against.
	// This symbolic constant has a value of 0.
	MatchTypeTrafficMatch MatchType = 0
	// FW_MATCH_TYPE_EQUAL:  The match operation evaluates to TRUE for all objects that
	// have a value equal to the one matched against. This symbolic constant has a value
	// of 1.
	MatchTypeEqual MatchType = 1
	// FW_MATCH_TYPE_MAX:  This value and values that exceed this value are not valid and
	// MUST NOT be used. It is defined for simplicity in writing IDL definitions and code.
	// This symbolic constant has a value of 2.
	MatchTypeMax MatchType = 2
)

func (o MatchType) String() string {
	switch o {
	case MatchTypeTrafficMatch:
		return "MatchTypeTrafficMatch"
	case MatchTypeEqual:
		return "MatchTypeEqual"
	case MatchTypeMax:
		return "MatchTypeMax"
	}
	return "Invalid"
}

// QueryCondition structure represents FW_QUERY_CONDITION RPC structure.
//
// This structure specifies a condition of a query. A condition can evaluate to TRUE
// or FALSE. It contains a match key that identifies what to match, a match value that
// identifies what to match with, and a match type that identifies how to match.
type QueryCondition struct {
	// matchKey:  This field identifies what information to match.
	MatchKey MatchKey `idl:"name:matchKey" json:"match_key"`
	// matchType:  This field identifies how to perform the match operation.
	MatchType MatchType `idl:"name:matchType" json:"match_type"`
	// matchValue:  This field identifies what to match with.
	//
	// A query condition structure MUST pass the following semantics checks:
	//
	// * The *matchKey* field MUST have a valid *FW_MATCH_KEY* value that is less than FW_MATCH_KEY_MAX,
	// MUST be a string of 1 or more characters, and MUST NOT be greater than or equal to
	// 255 characters.
	//
	// * The *matchType* field MUST have a valid *FW_MATCH_TYPE* value that is less than
	// FW_MATCH_KEY_MAX.
	//
	// * If the *matchType* field is equal to FW_MATCH_TYPE_EQUAL, the *matchKey* field
	// MUST be either FW_MATCH_KEY_GROUP or FW_MATCH_KEY_DIRECTION.
	//
	// * If the *matchKey* field is equal to FW_MATCH_KEY_PROFILE or FW_MATCH_KEY_STATUS,
	// the *matchValue* MUST have its type field equal to FW_DATA_TYPE_UINT32.
	//
	// * If the *matchKey* field is equal to FW_MATCH_KEY_FILTERID, the *matchValue* MUST
	// have its type field equal to FW_DATA_TYPE_UINT64.
	//
	// * If the *matchKey* field is equal to FW_MATCH_KEY_PROTOCOL, FW_MATCH_KEY_LOCAL_PORT,
	// or FW_MATCH_KEY_REMOTE_PORT; then the *matchValue* MUST have its type field equal
	// to FW_DATA_TYPE_UINT16.
	//
	// * If the *matchKey* field is equal to FW_MATCH_KEY_OBJECTID, FW_MATCH_KEY_APP_PATH,
	// FW_MATCH_KEY_GROUP, or FW_MATCH_KEY_SVC_NAME; then the *matchValue* MUST have its
	// type field equal to FW_DATA_TYPE_UNICODE_STRING.
	MatchValue *MatchValue `idl:"name:matchValue" json:"match_value"`
}

func (o *QueryCondition) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueryCondition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.MatchKey)); err != nil {
		return err
	}
	if err := w.WriteData(uint16(o.MatchType)); err != nil {
		return err
	}
	if o.MatchValue != nil {
		if err := o.MatchValue.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MatchValue{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueryCondition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.MatchKey)); err != nil {
		return err
	}
	if err := w.ReadData((*uint16)(&o.MatchType)); err != nil {
		return err
	}
	if o.MatchValue == nil {
		o.MatchValue = &MatchValue{}
	}
	if err := o.MatchValue.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// QueryConditions structure represents FW_QUERY_CONDITIONS RPC structure.
//
// This structure is used to contain a number of FW_QUERY_CONDITION elements. This structure
// can evaluate to either TRUE or FALSE. It evaluates to TRUE if all query condition
// elements evaluate to TRUE; otherwise, it evaluates to FALSE.
type QueryConditions struct {
	// dwNumEntries:  Specifies the number of query conditions that the structure contains.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// AndedConditions:  A pointer to an array of FW_QUERY_CONDITIONS elements, which are
	// to be logically AND'd together by the server.
	//
	// A query condition structure MUST pass the following semantic checks:
	//
	// * If the *dwNumEntries* field is zero, the *AndedConditions* field MUST be NULL;
	// and if the *dwNumEntries* field is not zero, the *AndedConditions* field MUST NOT
	// be NULL.
	//
	// * If the *AndedConditions* field array has a FW_QUERY_CONDITION element with the
	// *matchKey* field equal to FW_MATCH_KEY_LOCAL_PORT or FW_MATCH_KEY_REMOTE_PORT at
	// position N of the array, the array MUST have another element whose *matchKey* field
	// is equal to FW_MATCH_KEY_PROTOCOL at position M, where M < N.
	//
	// * All elements of the *AndedConditions* array MUST have valid FW_QUERY_CONDITION
	// structures.
	AndedConditions []*QueryCondition `idl:"name:AndedConditions;size_is:(dwNumEntries)" json:"anded_conditions"`
}

func (o *QueryConditions) xxx_PreparePayload(ctx context.Context) error {
	if o.AndedConditions != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.AndedConditions))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueryConditions) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.AndedConditions != nil || o.EntriesLength > 0 {
		_ptr_AndedConditions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.AndedConditions {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.AndedConditions[i1] != nil {
					if err := o.AndedConditions[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&QueryCondition{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.AndedConditions); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&QueryCondition{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AndedConditions, _ptr_AndedConditions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueryConditions) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_AndedConditions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AndedConditions", sizeInfo[0])
		}
		o.AndedConditions = make([]*QueryCondition, sizeInfo[0])
		for i1 := range o.AndedConditions {
			i1 := i1
			if o.AndedConditions[i1] == nil {
				o.AndedConditions[i1] = &QueryCondition{}
			}
			if err := o.AndedConditions[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_AndedConditions := func(ptr interface{}) { o.AndedConditions = *ptr.(*[]*QueryCondition) }
	if err := w.ReadPointer(&o.AndedConditions, _s_AndedConditions, _ptr_AndedConditions); err != nil {
		return err
	}
	return nil
}

// Query structure represents FW_QUERY RPC structure.
//
// This structure is used to query objects from the store. The structure contains a
// number of FW_QUERY_CONDITIONS elements. This structure can evaluate to either TRUE
// or FALSE. It evaluates to TRUE if at least one of the query conditions containers
// evaluates to TRUE; otherwise, if all evaluate to FALSE, it evaluates to FALSE.
type Query struct {
	// wSchemaVersion:  The schema version of the query object. The version MUST be at least
	// 0x00020A.
	SchemaVersion uint16 `idl:"name:wSchemaVersion" json:"schema_version"`
	// dwNumEntries:  This field specifies the number of query conditions containers that
	// the structure contains.
	EntriesLength uint32 `idl:"name:dwNumEntries" json:"entries_length"`
	// ORConditions:  A pointer to an array of FW_QUERY_CONDITIONS elements, which are all
	// logically OR'd together. The number of elements is given by dwNumEntries.
	OrConditions []*QueryConditions `idl:"name:ORConditions;size_is:(dwNumEntries)" json:"or_conditions"`
	// Status:  The status code of the query, as specified by the FW_RULE_STATUS enumeration.
	// This field is filled out when the structure is returned as output. On input, this
	// field SHOULD be set to FW_RULE_STATUS_OK.
	Status RuleStatus `idl:"name:Status" json:"status"`
}

func (o *Query) xxx_PreparePayload(ctx context.Context) error {
	if o.OrConditions != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint32(len(o.OrConditions))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Query) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SchemaVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if o.OrConditions != nil || o.EntriesLength > 0 {
		_ptr_ORConditions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.OrConditions {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.OrConditions[i1] != nil {
					if err := o.OrConditions[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&QueryConditions{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.OrConditions); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&QueryConditions{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OrConditions, _ptr_ORConditions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(uint32(o.Status)); err != nil {
		return err
	}
	return nil
}
func (o *Query) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SchemaVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	_ptr_ORConditions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OrConditions", sizeInfo[0])
		}
		o.OrConditions = make([]*QueryConditions, sizeInfo[0])
		for i1 := range o.OrConditions {
			i1 := i1
			if o.OrConditions[i1] == nil {
				o.OrConditions[i1] = &QueryConditions{}
			}
			if err := o.OrConditions[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ORConditions := func(ptr interface{}) { o.OrConditions = *ptr.(*[]*QueryConditions) }
	if err := w.ReadPointer(&o.OrConditions, _s_ORConditions, _ptr_ORConditions); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.Status)); err != nil {
		return err
	}
	return nil
}
