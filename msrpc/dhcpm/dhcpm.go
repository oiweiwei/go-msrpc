// The dhcpm package implements the DHCPM client protocol.
//
// # Introduction
//
// The Dynamic Host Configuration Protocol (DHCP) Server Management Protocol (DHCPM)
// defines remote procedure call (RPC) interfaces that provide methods for remotely
// accessing and administering the DHCP server. This RPC-based client/server protocol
// is used to configure, manage, and monitor a DHCP server.
//
// An application implementing this protocol can remotely administer the DHCP server.
// This protocol enables service monitoring as well as creating, updating, and deleting
// DHCP scopes and associated configuration options; retrieving and setting DHCP server
// bindings; and retrieving and creating DHCP client lease records.
//
// # Overview
//
// The Dynamic Host Configuration Protocol (DHCP) Server Management Protocol is a client/server
// protocol that is used to remotely configure, manage, and monitor the DHCP server.
// This protocol allows a client to view and update the server configuration settings
// as well as to create, modify, and delete DHCP client lease records. The protocol
// allows a client to access and modify DHCP server settings, enumerate and modify DHCP
// server configuration (DHCP scopes, reservations, exclusions, option definition, and
// option values), and monitor DHCP client lease records.
//
// The DHCP Server Management Protocol (DHCPM) is a stateless protocol with no state
// shared across RPC method calls. Each RPC method call contains one complete request.
// Output from one method call can be used as an input to another call, but the protocol
// does not provide for locking of the DHCP server configuration or state data across
// method calls. For example, a client enumerates DHCP subnets with one call and then
// retrieves the properties of one or more DHCP subnets with another call. However,
// the protocol does not guarantee that the specified subnet has not been deleted by
// another client between the two method calls.
package dhcpm

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
	GoPackage = "dhcpm"
)

// DateTimeZeroHigh represents the DHCP_DATE_TIME_ZERO_HIGH RPC constant
var DateTimeZeroHigh = 0

// DateTimeZeroLow represents the DHCP_DATE_TIME_ZERO_LOW RPC constant
var DateTimeZeroLow = 0

// DateTimeInfinitHigh represents the DHCP_DATE_TIME_INFINIT_HIGH RPC constant
var DateTimeInfinitHigh = 2147483647

// DateTimeInfinitLow represents the DHCP_DATE_TIME_INFINIT_LOW RPC constant
var DateTimeInfinitLow = 4294967295

// DebugAddress represents the DEBUG_ADDRESS RPC constant
var DebugAddress = 1

// DebugClient represents the DEBUG_CLIENT RPC constant
var DebugClient = 2

// DebugParameters represents the DEBUG_PARAMETERS RPC constant
var DebugParameters = 4

// DebugOptions represents the DEBUG_OPTIONS RPC constant
var DebugOptions = 8

// DebugErrors represents the DEBUG_ERRORS RPC constant
var DebugErrors = 16

// DebugStoc represents the DEBUG_STOC RPC constant
var DebugStoc = 32

// DebugInit represents the DEBUG_INIT RPC constant
var DebugInit = 64

// DebugScavenger represents the DEBUG_SCAVENGER RPC constant
var DebugScavenger = 128

// DebugTimestamp represents the DEBUG_TIMESTAMP RPC constant
var DebugTimestamp = 256

// DebugAPIs represents the DEBUG_APIS RPC constant
var DebugAPIs = 512

// DebugRegistry represents the DEBUG_REGISTRY RPC constant
var DebugRegistry = 1024

// DebugJet represents the DEBUG_JET RPC constant
var DebugJet = 2048

// DebugThreadPool represents the DEBUG_THREADPOOL RPC constant
var DebugThreadPool = 4096

// DebugAuditlog represents the DEBUG_AUDITLOG RPC constant
var DebugAuditlog = 8192

// DebugQuarantine represents the DEBUG_QUARANTINE RPC constant
var DebugQuarantine = 16384

// DebugMisc represents the DEBUG_MISC RPC constant
var DebugMisc = 32768

// DebugMessage represents the DEBUG_MESSAGE RPC constant
var DebugMessage = 65536

// DebugAPIVerbose represents the DEBUG_API_VERBOSE RPC constant
var DebugAPIVerbose = 131072

// DebugDNS represents the DEBUG_DNS RPC constant
var DebugDNS = 262144

// DebugMstoc represents the DEBUG_MSTOC RPC constant
var DebugMstoc = 524288

// DebugTrack represents the DEBUG_TRACK RPC constant
var DebugTrack = 1048576

// DebugRogue represents the DEBUG_ROGUE RPC constant
var DebugRogue = 2097152

// DebugPnP represents the DEBUG_PNP RPC constant
var DebugPnP = 4194304

// DebugPerf represents the DEBUG_PERF RPC constant
var DebugPerf = 16777216

// DebugAlloc represents the DEBUG_ALLOC RPC constant
var DebugAlloc = 33554432

// DebugPing represents the DEBUG_PING RPC constant
var DebugPing = 67108864

// DebugThread represents the DEBUG_THREAD RPC constant
var DebugThread = 134217728

// DebugTrace represents the DEBUG_TRACE RPC constant
var DebugTrace = 268435456

// DebugTraceCalls represents the DEBUG_TRACE_CALLS RPC constant
var DebugTraceCalls = 536870912

// DebugStartupBrk represents the DEBUG_STARTUP_BRK RPC constant
var DebugStartupBrk = 1073741824

// DebugLogInFile represents the DEBUG_LOG_IN_FILE RPC constant
var DebugLogInFile = 2147483648

// ServerUseRPCOverTCPIP represents the DHCP_SERVER_USE_RPC_OVER_TCPIP RPC constant
var ServerUseRPCOverTCPIP = 1

// ServerUseRPCOverNP represents the DHCP_SERVER_USE_RPC_OVER_NP RPC constant
var ServerUseRPCOverNP = 2

// ServerUseRPCOverLPC represents the DHCP_SERVER_USE_RPC_OVER_LPC RPC constant
var ServerUseRPCOverLPC = 4

// ServerUseRPCOverAll represents the DHCP_SERVER_USE_RPC_OVER_ALL RPC constant
var ServerUseRPCOverAll = 7

// DatabaseCleanupInterval represents the DHCP_DATABASE_CLEANUP_INTERVAL RPC constant
var DatabaseCleanupInterval = 10800000

// DefaultBackupInterval represents the DEFAULT_BACKUP_INTERVAL RPC constant
var DefaultBackupInterval = 900000

// SetAPIProtocolSupport represents the Set_APIProtocolSupport RPC constant
var SetAPIProtocolSupport = 1

// SetDatabaseName represents the Set_DatabaseName RPC constant
var SetDatabaseName = 2

// SetDatabasePath represents the Set_DatabasePath RPC constant
var SetDatabasePath = 4

// SetBackupPath represents the Set_BackupPath RPC constant
var SetBackupPath = 8

// SetBackupInterval represents the Set_BackupInterval RPC constant
var SetBackupInterval = 16

// SetDatabaseLoggingFlag represents the Set_DatabaseLoggingFlag RPC constant
var SetDatabaseLoggingFlag = 32

// SetRestoreFlag represents the Set_RestoreFlag RPC constant
var SetRestoreFlag = 64

// SetDatabaseCleanupInterval represents the Set_DatabaseCleanupInterval RPC constant
var SetDatabaseCleanupInterval = 128

// SetDebugFlag represents the Set_DebugFlag RPC constant
var SetDebugFlag = 256

// SetPingRetries represents the Set_PingRetries RPC constant
var SetPingRetries = 512

// SetBootFileTable represents the Set_BootFileTable RPC constant
var SetBootFileTable = 1024

// SetAuditLogState represents the Set_AuditLogState RPC constant
var SetAuditLogState = 2048

// SetQuarantineOn represents the Set_QuarantineON RPC constant
var SetQuarantineOn = 4096

// SetQuarantineDefFail represents the Set_QuarantineDefFail RPC constant
var SetQuarantineDefFail = 8192

// ClientTypeUnspecified represents the CLIENT_TYPE_UNSPECIFIED RPC constant
var ClientTypeUnspecified = 0

// ClientTypeDHCP represents the CLIENT_TYPE_DHCP RPC constant
var ClientTypeDHCP = 1

// ClientTypeBOOTP represents the CLIENT_TYPE_BOOTP RPC constant
var ClientTypeBOOTP = 2

// ClientTypeBoth represents the CLIENT_TYPE_BOTH RPC constant
var ClientTypeBoth = 3

// MaxDetectConflictRetries represents the MAX_DETECT_CONFLICT_RETRIES RPC constant
var MaxDetectConflictRetries = 5

// MinDetectConflictRetries represents the MIN_DETECT_CONFLICT_RETRIES RPC constant
var MinDetectConflictRetries = 0

// ClientTypeReservationFlag represents the CLIENT_TYPE_RESERVATION_FLAG RPC constant
var ClientTypeReservationFlag = 4

// ClientTypeNone represents the CLIENT_TYPE_NONE RPC constant
var ClientTypeNone = 100

// AddressStateOffered represents the ADDRESS_STATE_OFFERED RPC constant
var AddressStateOffered = 0

// AddressStateActive represents the ADDRESS_STATE_ACTIVE RPC constant
var AddressStateActive = 1

// AddressStateDeclined represents the ADDRESS_STATE_DECLINED RPC constant
var AddressStateDeclined = 2

// AddressStateDoom represents the ADDRESS_STATE_DOOM RPC constant
var AddressStateDoom = 3

// AttributeTypeBool represents the DHCP_ATTRIB_TYPE_BOOL RPC constant
var AttributeTypeBool = 1

// AttributeTypeUint32 represents the DHCP_ATTRIB_TYPE_ULONG RPC constant
var AttributeTypeUint32 = 2

// EndpointFlagCantModify represents the DHCP_ENDPOINT_FLAG_CANT_MODIFY RPC constant
var EndpointFlagCantModify = 1

// MaxPatternLength represents the MAX_PATTERN_LENGTH RPC constant
var MaxPatternLength = 255

// MACAddressLength represents the MAC_ADDRESS_LENGTH RPC constant
var MACAddressLength = 6

// HwTypeEthernet10Mb represents the HWTYPE_ETHERNET_10MB RPC constant
var HwTypeEthernet10Mb = 1

// MaxFreeAddressesRequested represents the DHCP_MAX_FREE_ADDRESSES_REQUESTED RPC constant
var MaxFreeAddressesRequested = 1024

// BinaryData structure represents DHCP_BINARY_DATA RPC structure.
//
// The DHCP_BINARY_DATA structure defines a buffer containing binary data. This data
// structure is generally used by other data structures, such as DHCP_OPTION_DATA_ELEMENT
// (section 2.2.1.2.23).
type BinaryData struct {
	// DataLength:   This is a DWORD that contains the number of bytes of data stored in
	// the Data member buffer. There is no restriction imposed by this protocol on the length
	// of the data.<11>
	DataLength uint32 `idl:"name:DataLength" json:"data_length"`
	// Data:   This is a pointer to BYTE, pointing to an array of bytes of length specified
	// by the DataLength member.
	Data []byte `idl:"name:Data;size_is:(DataLength)" json:"data"`
}

func (o *BinaryData) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BinaryData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if o.Data != nil || o.DataLength > 0 {
		_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Data {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Data[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_Data); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *BinaryData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_Data, _ptr_Data); err != nil {
		return err
	}
	return nil
}

// ClientUID structure represents DHCP_CLIENT_UID RPC structure.
type ClientUID BinaryData

func (o *ClientUID) BinaryData() *BinaryData { return (*BinaryData)(o) }

func (o *ClientUID) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.DataLength); err != nil {
		return err
	}
	if o.Data != nil || o.DataLength > 0 {
		_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Data {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Data[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_Data); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataLength); err != nil {
		return err
	}
	_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_Data, _ptr_Data); err != nil {
		return err
	}
	return nil
}

// SubnetState type represents DHCP_SUBNET_STATE RPC enumeration.
//
// The DHCP_SUBNET_STATE enumeration is a DWORD value that specifies the set of possible
// states for a subnet configured on a DHCPv4 server.
type SubnetState uint16

var (
	// DhcpSubnetEnabled: The subnet is enabled; the DHCP server assigns IP addresses, extends
	// IP address leases, and releases unused IP addresses for DHCP clients on this subnet.
	SubnetStateEnabled SubnetState = 0
	// DhcpSubnetDisabled: The subnet is disabled; the DHCP server does not assign IP addresses
	// or extend IP address leases for DHCP clients on this subnet. However, the DHCP server
	// still releases unused IP addresses for DHCP clients on this subnet.
	SubnetStateDisabled SubnetState = 1
	// DhcpSubnetEnabledSwitched: The subnet is enabled; the DHCP server assigns IP addresses,
	// extends IP address leases, and releases unused IP addresses for DHCP clients on this
	// subnet. In addition, the default gateway for the DHCP client is set to on-link route.
	SubnetStateEnabledSwitched SubnetState = 2
	// DhcpSubnetDisabledSwitched: The subnet is disabled; the DHCP server does not distribute
	// addresses or extend leases within the subnet range to clients. However, the DHCP
	// server still releases addresses within the subnet range. The system behavior in the
	// DhcpSubnetDisabledSwitched state is identical to the state described in DhcpSubnetDisabled.
	// Any software that uses the DHCPM API can use the DhcpSubnetDisabledSwitched state
	// to remember that a particular scope needs to be put into the DhcpSubnetEnabledSwitched
	// state when enabled.
	SubnetStateDisabledSwitched SubnetState = 3
	// DhcpSubnetInvalidState: The subnet is not valid, and hence no address will be distributed
	// or extended.
	SubnetStateInvalidState SubnetState = 4
)

func (o SubnetState) String() string {
	switch o {
	case SubnetStateEnabled:
		return "SubnetStateEnabled"
	case SubnetStateDisabled:
		return "SubnetStateDisabled"
	case SubnetStateEnabledSwitched:
		return "SubnetStateEnabledSwitched"
	case SubnetStateDisabledSwitched:
		return "SubnetStateDisabledSwitched"
	case SubnetStateInvalidState:
		return "SubnetStateInvalidState"
	}
	return "Invalid"
}

// SubnetElementType type represents DHCP_SUBNET_ELEMENT_TYPE RPC enumeration.
//
// The DHCP_SUBNET_ELEMENT_TYPE enumeration defines the type of a configuration parameter
// for a DHCPv4 scope configured on the DHCP server. This value is used in conjunction
// with other data types to specify the configuration parameters for a DHCPv4 scope
// by the RPC methods defined in this specification.
type SubnetElementType uint16

var (
	// DhcpIpRanges: The configuration parameter is the IP range of a DHCPv4 scope configured
	// on the DHCP server.
	SubnetElementTypeIPRanges SubnetElementType = 0
	// DhcpSecondaryHosts: This enumeration type is unused. If this value is passed as a
	// parameter to a method, it will return ERROR_CALL_NOT_IMPLEMENTED or ERROR_NOT_SUPPORTED,
	// as specified in the processing rules of methods that use the DHCP_SUBNET_ELEMENT_TYPE
	// enumeration.
	SubnetElementTypeSecondaryHosts SubnetElementType = 1
	// DhcpReservedIps: The configuration parameter is a reservation for a DHCPv4 client
	// in a DHCPv4 scope element configured on the DHCP server.
	SubnetElementTypeReservedIPs SubnetElementType = 2
	// DhcpExcludedIpRanges: The configuration parameter is the exclusion range of a DHCPv4
	// scope configured on the DHCPv4 server.
	SubnetElementTypeExcludedIPRanges SubnetElementType = 3
	// DhcpIpUsedClusters: This enumeration type is unused, and the DHCP server returns
	// ERROR_INVALID_PARAMETER when specified.
	SubnetElementTypeIPUsedClusters SubnetElementType = 4
	// DhcpIpRangesDhcpOnly: The configuration parameter is an IP range of a DHCPv4 scope
	// configured on the DHCPv4 server, which MUST be used only for assignment of addresses
	// to DHCPv4 clients on the subnet. The IP addresses from this range MUST NOT be assigned
	// to bootstrap protocol (BOOTP) clients ([RFC2132]).
	SubnetElementTypeIPRangesDHCPOnly SubnetElementType = 5
	// DhcpIpRangesDhcpBootp: The configuration parameter is an IP range of a DHCPv4 scope
	// configured on the DHCPv4 server, which can be used for assignment of addresses to
	// both DHCPv4 and BOOTP.
	SubnetElementTypeIPRangesDHCPBOOTP SubnetElementType = 6
	// DhcpIpRangesBootpOnly: The configuration parameter is an IP range of a DHCPv4 scope
	// configured on the DHCPv4 server, which MUST be used only for assignment of IPv4 addresses
	// to BOOTP clients.
	SubnetElementTypeIPRangesBOOTPOnly SubnetElementType = 7
)

func (o SubnetElementType) String() string {
	switch o {
	case SubnetElementTypeIPRanges:
		return "SubnetElementTypeIPRanges"
	case SubnetElementTypeSecondaryHosts:
		return "SubnetElementTypeSecondaryHosts"
	case SubnetElementTypeReservedIPs:
		return "SubnetElementTypeReservedIPs"
	case SubnetElementTypeExcludedIPRanges:
		return "SubnetElementTypeExcludedIPRanges"
	case SubnetElementTypeIPUsedClusters:
		return "SubnetElementTypeIPUsedClusters"
	case SubnetElementTypeIPRangesDHCPOnly:
		return "SubnetElementTypeIPRangesDHCPOnly"
	case SubnetElementTypeIPRangesDHCPBOOTP:
		return "SubnetElementTypeIPRangesDHCPBOOTP"
	case SubnetElementTypeIPRangesBOOTPOnly:
		return "SubnetElementTypeIPRangesBOOTPOnly"
	}
	return "Invalid"
}

// ForceFlag type represents DHCP_FORCE_FLAG RPC enumeration.
//
// The DHCP_FORCE_FLAG enumeration defines the type of deletion operation being requested
// by an RPC method specified by this protocol. This value is used with the RPC method
// R_DhcpDeleteSubnetV6 (section 3.2.4.63).
type ForceFlag uint16

var (
	// DhcpFullForce: The DHCP server deletes all the active DHCP client lease records for
	// the specified subnet and then deletes all the configurations associated with that
	// subnet.
	ForceFlagFullForce ForceFlag = 0
	// DhcpNoForce: The DHCP server deletes all the configuration associated with the specified
	// subnet, but only if there are no active DHCP client lease records for the specified
	// subnet. If there are any active DHCP client lease records for the specified subnet,
	// then nothing is deleted.
	ForceFlagNoForce ForceFlag = 1
	// DhcpFailoverForce: The DHCP server deletes all the active DHCP client lease records
	// for the specified subnet but does not delete the Dynamic DNS updates.<6>
	ForceFlagFailoverForce ForceFlag = 2
)

func (o ForceFlag) String() string {
	switch o {
	case ForceFlagFullForce:
		return "ForceFlagFullForce"
	case ForceFlagNoForce:
		return "ForceFlagNoForce"
	case ForceFlagFailoverForce:
		return "ForceFlagFailoverForce"
	}
	return "Invalid"
}

// OptionType type represents DHCP_OPTION_TYPE RPC enumeration.
//
// The DHCP_OPTION_TYPE enumeration specifies whether the option value for a specific
// standard or vendor-specific option is single-valued or multivalued. The following
// structure specifies the values defined for this.
type OptionType uint16

var (
	// DhcpUnaryElementTypeOption: The option value is single-valued.
	OptionTypeUnaryElementTypeOption OptionType = 0
	// DhcpArrayTypeOption: The option value is multivalued.
	OptionTypeArrayTypeOption OptionType = 1
)

func (o OptionType) String() string {
	switch o {
	case OptionTypeUnaryElementTypeOption:
		return "OptionTypeUnaryElementTypeOption"
	case OptionTypeArrayTypeOption:
		return "OptionTypeArrayTypeOption"
	}
	return "Invalid"
}

// OptionDataType type represents DHCP_OPTION_DATA_TYPE RPC enumeration.
//
// The DHCP_OPTION_DATA_TYPE enumeration defines the format types for DHCP option values
// and is used in the DHCP_OPTION_DATA_ELEMENT (section 2.2.1.2.23) structure. The DHCPM
// RPC methods can create the option definition on the DHCP server, which can contain
// option value in various formats.
type OptionDataType uint16

var (
	// DhcpByteOption: The option value is of type BYTE.
	OptionDataTypeByteOption OptionDataType = 0
	// DhcpWordOption: The option value is of type WORD.
	OptionDataTypeWordOption OptionDataType = 1
	// DhcpDWordOption: The option value is of type DWORD [MS-DTYP] section 2.2.9.
	OptionDataTypeDwordOption OptionDataType = 2
	// DhcpDWordDWordOption: The option value is of type DWORD_DWORD (section 2.2.1.2.22).
	OptionDataTypeDwordDwordOption OptionDataType = 3
	// DhcpIpAddressOption: The option value is of type DHCP_IP_ADDRESS (section 2.2.1.2.1).
	OptionDataTypeIPAddressOption OptionDataType = 4
	// DhcpStringDataOption: The option value is a pointer, of type LPWSTR, to a null-terminated
	// UNICODE string.
	OptionDataTypeStringDataOption OptionDataType = 5
	// DhcpBinaryDataOption: The option value is of type DHCP_BINARY_DATA (section 2.2.1.2.9).
	OptionDataTypeBinaryDataOption OptionDataType = 6
	// DhcpEncapsulatedDataOption: The option value is encapsulated and of type DHCP_BINARY_DATA.
	OptionDataTypeEncapsulatedDataOption OptionDataType = 7
	// DhcpIpv6AddressOption: The option value is an IPv6 address represented as a pointer,
	// of type LPWSTR, to a null-terminated Unicode string.
	OptionDataTypeIPv6AddressOption OptionDataType = 8
)

func (o OptionDataType) String() string {
	switch o {
	case OptionDataTypeByteOption:
		return "OptionDataTypeByteOption"
	case OptionDataTypeWordOption:
		return "OptionDataTypeWordOption"
	case OptionDataTypeDwordOption:
		return "OptionDataTypeDwordOption"
	case OptionDataTypeDwordDwordOption:
		return "OptionDataTypeDwordDwordOption"
	case OptionDataTypeIPAddressOption:
		return "OptionDataTypeIPAddressOption"
	case OptionDataTypeStringDataOption:
		return "OptionDataTypeStringDataOption"
	case OptionDataTypeBinaryDataOption:
		return "OptionDataTypeBinaryDataOption"
	case OptionDataTypeEncapsulatedDataOption:
		return "OptionDataTypeEncapsulatedDataOption"
	case OptionDataTypeIPv6AddressOption:
		return "OptionDataTypeIPv6AddressOption"
	}
	return "Invalid"
}

// OptionScopeType type represents DHCP_OPTION_SCOPE_TYPE RPC enumeration.
//
// The DHCP_OPTION_SCOPE_TYPE enumeration defines the type of DHCPv4 options being referred
// to by an RPC method in the DHCPM. The DHCP server allows for configuration of standard
// and vendor-specific options at various levels, such as the default level, server
// level, or scope level, or for a specific reservation. This value is used in conjunction
// with union DHCP_OPTION_SCOPE_UNION, as defined in the DHCP_OPTION_SCOPE_INFO (section
// 2.2.1.2.41) structure, to specify option values in the RPC methods defined by this
// protocol.
type OptionScopeType uint16

var (
	// DhcpDefaultOptions: Option is defined at the default level. The option definition
	// is created or modified on the DHCPv4 server and the default value of the option is
	// stored.
	OptionScopeTypeDefaultOptions OptionScopeType = 0
	// DhcpGlobalOptions: Option is defined at the server level. The option value is added
	// or modified at the DHCPv4 server, which is valid for all scopes in that server.
	OptionScopeTypeGlobalOptions OptionScopeType = 1
	// DhcpSubnetOptions: Option is defined at the scope level. The option value is added
	// or modified at the scope and is valid for that specific scope.
	OptionScopeTypeSubnetOptions OptionScopeType = 2
	// DhcpReservedOptions: Option is defined for a specific IP address reservation. The
	// option value is added or modified for a specific IP reservation in a scope.
	OptionScopeTypeReservedOptions OptionScopeType = 3
	// DhcpMScopeOptions: Option is defined for a multicast scope. The option value is added
	// or modified for a multicast scope.
	OptionScopeTypeMScopeOptions OptionScopeType = 4
)

func (o OptionScopeType) String() string {
	switch o {
	case OptionScopeTypeDefaultOptions:
		return "OptionScopeTypeDefaultOptions"
	case OptionScopeTypeGlobalOptions:
		return "OptionScopeTypeGlobalOptions"
	case OptionScopeTypeSubnetOptions:
		return "OptionScopeTypeSubnetOptions"
	case OptionScopeTypeReservedOptions:
		return "OptionScopeTypeReservedOptions"
	case OptionScopeTypeMScopeOptions:
		return "OptionScopeTypeMScopeOptions"
	}
	return "Invalid"
}

// SearchInfoType type represents DHCP_SEARCH_INFO_TYPE RPC enumeration.
//
// The DHCP_SEARCH_INFO_TYPE enumeration defines the type of search that can be performed
// on the DHCPv4 server to query specific DHCP client records. DHCPM uses this value
// in conjunction with DHCP_SEARCH_INFO (section 2.2.1.2.18) to query specific DHCPv4
// client address records.
type SearchInfoType uint16

var (
	// DhcpClientIpAddress: The DHCPv4 client IP address MUST be used for querying the DHCPv4
	// client lease records from the database on the DHCPv4 server.
	SearchInfoTypeClientIPAddress SearchInfoType = 0
	// DhcpClientHardwareAddress: The DHCPv4 client unique ID (section 2.2.1.2.5.2) MUST
	// be used for querying the DHCPv4 client lease records from the database on the DHCPv4
	// server.
	SearchInfoTypeClientHardwareAddress SearchInfoType = 1
	// DhcpClientName: The null-terminated Unicode string containing the name of the DHCPv4
	// client MUST be used for querying the DHCPv4 client lease records on the DHCPv4 server.
	// There is no restriction on the length of this UNICODE string.
	SearchInfoTypeClientName SearchInfoType = 2
)

func (o SearchInfoType) String() string {
	switch o {
	case SearchInfoTypeClientIPAddress:
		return "SearchInfoTypeClientIPAddress"
	case SearchInfoTypeClientHardwareAddress:
		return "SearchInfoTypeClientHardwareAddress"
	case SearchInfoTypeClientName:
		return "SearchInfoTypeClientName"
	}
	return "Invalid"
}

// ScanFlag type represents DHCP_SCAN_FLAG RPC enumeration.
//
// The DHCP_SCAN_FLAG enumeration defines whether an inconsistent IP address needs to
// be fixed in the DHCPv4 client Lease records or the bitmask representation in memory
// (section 3.1.1.4). This enumeration is used in the DHCP_SCAN_ITEM (section 2.2.1.2.73)
// structure.
type ScanFlag uint16

var (
	// DhcpRegistryFix: The DHCPv4 server sets this value in DHCP_SCAN_ITEM when a DHCP
	// client IPv4 address is found in the DHCPv4 client Lease records but not in the bitmask
	// representation in memory (section 3.1.1.4).
	ScanFlagRegistryFix ScanFlag = 0
	// DhcpDatabaseFix: The DHCPv4 server sets this value in DHCP_SCAN_ITEM when the DHCP
	// client IPv4 address is found in the bitmask representation in memory (section 3.1.1.4)
	// but not in the DHCPv4 client Lease records.
	ScanFlagDatabaseFix ScanFlag = 1
)

func (o ScanFlag) String() string {
	switch o {
	case ScanFlagRegistryFix:
		return "ScanFlagRegistryFix"
	case ScanFlagDatabaseFix:
		return "ScanFlagDatabaseFix"
	}
	return "Invalid"
}

// QuarantineStatus type represents QuarantineStatus RPC enumeration.
//
// The QuarantineStatus enumeration defines the Network Access Protection (NAP) state
// of the DHCP client.<7>
type QuarantineStatus uint16

var (
	// NOQUARANTINE: The DHCP client is compliant with the health policies defined by the
	// administrator and has normal access to the network.
	QuarantineStatusNoQuarantine QuarantineStatus = 0
	// RESTRICTEDACCESS: The DHCP client is not compliant with the health policies defined
	// by the administrator and is being quarantined with restricted access to the network.
	QuarantineStatusRestrictedAccess QuarantineStatus = 1
	// DROPPACKET: The DHCP client is not compliant with the health policies defined by
	// the administrator and is being denied access to the network. The DHCP server does
	// not grant an IP address lease to this client.
	QuarantineStatusDropPacket QuarantineStatus = 2
	// PROBATION: The DHCP client is not compliant with the health policies defined by the
	// administrator and is being granted normal access to the network for a limited time.
	QuarantineStatusProbation QuarantineStatus = 3
	// EXEMPT: The DHCP client is exempt from compliance with the health policies defined
	// by the administrator and is granted normal access to the network.
	QuarantineStatusExempt QuarantineStatus = 4
	// DEFAULTQUARSETTING: The DHCP client is put into the default quarantine state configured
	// on the DHCP NAP server. When a network policy server (NPS) is unavailable, the DHCP
	// client can be put in any of the states NOQUARANTINE, RESTRICTEDACCESS, or DROPPACKET,
	// depending on the default setting on the DHCP NAP server.
	QuarantineStatusDefaultQuarantineSetting QuarantineStatus = 5
	// NOQUARINFO: No quarantine.
	QuarantineStatusNoQuarantineInfo QuarantineStatus = 6
)

func (o QuarantineStatus) String() string {
	switch o {
	case QuarantineStatusNoQuarantine:
		return "QuarantineStatusNoQuarantine"
	case QuarantineStatusRestrictedAccess:
		return "QuarantineStatusRestrictedAccess"
	case QuarantineStatusDropPacket:
		return "QuarantineStatusDropPacket"
	case QuarantineStatusProbation:
		return "QuarantineStatusProbation"
	case QuarantineStatusExempt:
		return "QuarantineStatusExempt"
	case QuarantineStatusDefaultQuarantineSetting:
		return "QuarantineStatusDefaultQuarantineSetting"
	case QuarantineStatusNoQuarantineInfo:
		return "QuarantineStatusNoQuarantineInfo"
	}
	return "Invalid"
}

// HostInfo structure represents DHCP_HOST_INFO RPC structure.
//
// The DHCP_HOST_INFO structure provides information on the DHCPv4 server. This structure
// is used in DHCP_CLIENT_INFO_V4 (section 2.2.1.2.14) and DHCP_CLIENT_INFO_VQ (section
// 2.2.1.2.19).
type HostInfo struct {
	// IpAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing the IPv4
	// address of the DHCPv4 server.
	IPAddress uint32 `idl:"name:IpAddress" json:"ip_address"`
	// NetBiosName:  A pointer to a null-terminated Unicode string that points to the NetBIOS
	// name of the DHCPv4 server.
	NetBIOSName string `idl:"name:NetBiosName" json:"net_bios_name"`
	// HostName:  A pointer to a null-terminated Unicode string that points to the name
	// of the DHCPv4 server. Currently not used in any set method. If used in a get method,
	// the value returned is NULL.
	HostName string `idl:"name:HostName" json:"host_name"`
}

func (o *HostInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *HostInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddress); err != nil {
		return err
	}
	if o.NetBIOSName != "" {
		_ptr_NetBiosName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.NetBIOSName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.NetBIOSName, _ptr_NetBiosName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.HostName != "" {
		_ptr_HostName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.HostName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.HostName, _ptr_HostName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *HostInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddress); err != nil {
		return err
	}
	_ptr_NetBiosName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.NetBIOSName); err != nil {
			return err
		}
		return nil
	})
	_s_NetBiosName := func(ptr interface{}) { o.NetBIOSName = *ptr.(*string) }
	if err := w.ReadPointer(&o.NetBIOSName, _s_NetBiosName, _ptr_NetBiosName); err != nil {
		return err
	}
	_ptr_HostName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.HostName); err != nil {
			return err
		}
		return nil
	})
	_s_HostName := func(ptr interface{}) { o.HostName = *ptr.(*string) }
	if err := w.ReadPointer(&o.HostName, _s_HostName, _ptr_HostName); err != nil {
		return err
	}
	return nil
}

// SubnetInfo structure represents DHCP_SUBNET_INFO RPC structure.
//
// The DHCP_SUBNET_INFO structure defines the information about an IPv4 subnet. This
// structure is used in the R_DhcpCreateSubnet (section 3.1.4.1) method.
type SubnetInfo struct {
	// SubnetAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), specifying the
	// IPv4 subnet ID.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// SubnetMask:  This is of type DHCP_IP_MASK (section 2.2.1.2.2), specifying the subnet
	// IPv4 mask.
	SubnetMask uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
	// SubnetName:  A pointer of type LPWSTR to a null-terminated Unicode string that points
	// to the name of the subnet. There is no restriction on the length of this Unicode
	// string.
	SubnetName string `idl:"name:SubnetName" json:"subnet_name"`
	// SubnetComment:  A pointer of type LPWSTR to a null-terminated Unicode string that
	// points an optional comment specific to this subnet. There is no restriction on the
	// length of this Unicode string.
	SubnetComment string `idl:"name:SubnetComment" json:"subnet_comment"`
	// PrimaryHost:  This is of type DHCP_HOST_INFO (section 2.2.1.2.7) structure, containing
	// information about the DHCPv4 server servicing this IPv4 subnet.
	PrimaryHost *HostInfo `idl:"name:PrimaryHost" json:"primary_host"`
	// SubnetState:  This is an enumeration of type DHCP_SUBNET_STATE (section 2.2.1.1.2),
	// indicating the current state of this IPv4 subnet.
	SubnetState SubnetState `idl:"name:SubnetState" json:"subnet_state"`
}

func (o *SubnetInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	if o.SubnetName != "" {
		_ptr_SubnetName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SubnetName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SubnetName, _ptr_SubnetName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SubnetComment != "" {
		_ptr_SubnetComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SubnetComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SubnetComment, _ptr_SubnetComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PrimaryHost != nil {
		if err := o.PrimaryHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.SubnetState)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *SubnetInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	_ptr_SubnetName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SubnetName); err != nil {
			return err
		}
		return nil
	})
	_s_SubnetName := func(ptr interface{}) { o.SubnetName = *ptr.(*string) }
	if err := w.ReadPointer(&o.SubnetName, _s_SubnetName, _ptr_SubnetName); err != nil {
		return err
	}
	_ptr_SubnetComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SubnetComment); err != nil {
			return err
		}
		return nil
	})
	_s_SubnetComment := func(ptr interface{}) { o.SubnetComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.SubnetComment, _s_SubnetComment, _ptr_SubnetComment); err != nil {
		return err
	}
	if o.PrimaryHost == nil {
		o.PrimaryHost = &HostInfo{}
	}
	if err := o.PrimaryHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.SubnetState)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// IPArray structure represents DHCP_IP_ARRAY RPC structure.
//
// The DHCP_IP_ARRAY structure defines the array of type DHCP_IP_ADDRESS (section 2.2.1.2.1),
// typed as a DWORD. This structure is used in the R_DhcpEnumSubnets (section 3.1.4.4)
// method.
type IPArray struct {
	// NumElements:  This is of type DWORD, containing the number of IPv4 addresses in the
	// subsequent field, the Elements member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This is a pointer to an array of DHCP_IP_ADDRESS DWORD types of length
	// NumElements containing the IPv4 addresses of the subnets.
	Elements []uint32 `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *IPArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Elements[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]uint32, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if err := w.ReadData(&o.Elements[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// IPRange structure represents DHCP_IP_RANGE RPC structure.
//
// The DHCP_IP_RANGE structure defines the IPv4 range for an IPv4 scope. This is used
// in structure DHCP_SUBNET_ELEMENT_DATA (section 2.2.1.2.33).
type IPRange struct {
	// StartAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing the
	// first IPv4 address in the IPv4 range.
	StartAddress uint32 `idl:"name:StartAddress" json:"start_address"`
	// EndAddress:  This is of type DHCP_IP_ADDRESS, containing the last IPv4 address in
	// the IPv4 range.
	EndAddress uint32 `idl:"name:EndAddress" json:"end_address"`
}

func (o *IPRange) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.StartAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.EndAddress); err != nil {
		return err
	}
	return nil
}
func (o *IPRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.StartAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.EndAddress); err != nil {
		return err
	}
	return nil
}

// IPReservation structure represents DHCP_IP_RESERVATION RPC structure.
//
// The DHCP_IP_RESERVATION structure defines an IPv4 reservation for a DHCPv4 client.
// This is used in structure DHCP_SUBNET_ELEMENT_DATA (section 2.2.1.2.33).
type IPReservation struct {
	// ReservedIpAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1) that contains
	// the IPv4 address of the DHCPv4 client for which a reservation is created.
	ReservedIPAddress uint32 `idl:"name:ReservedIpAddress" json:"reserved_ip_address"`
	// ReservedForClient:  This is a pointer of type DHCP_CLIENT_UID (section 2.2.1.2.5)
	// that represents the DHCPv4 client identifier (section 2.2.1.2.5.1).
	ReservedForClient *ClientUID `idl:"name:ReservedForClient" json:"reserved_for_client"`
}

func (o *IPReservation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPReservation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReservedIPAddress); err != nil {
		return err
	}
	if o.ReservedForClient != nil {
		_ptr_ReservedForClient := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ReservedForClient != nil {
				if err := o.ReservedForClient.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReservedForClient, _ptr_ReservedForClient); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPReservation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReservedIPAddress); err != nil {
		return err
	}
	_ptr_ReservedForClient := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ReservedForClient == nil {
			o.ReservedForClient = &ClientUID{}
		}
		if err := o.ReservedForClient.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ReservedForClient := func(ptr interface{}) { o.ReservedForClient = *ptr.(**ClientUID) }
	if err := w.ReadPointer(&o.ReservedForClient, _s_ReservedForClient, _ptr_ReservedForClient); err != nil {
		return err
	}
	return nil
}

// IPCluster structure represents DHCP_IP_CLUSTER RPC structure.
//
// The DHCP_IP_CLUSTER structure is not used.
type IPCluster struct {
	// ClusterAddress:  This is of type DHCP_IP_ADDRESS, a DWORD.
	ClusterAddress uint32 `idl:"name:ClusterAddress" json:"cluster_address"`
	// ClusterMask:  This is of type DWORD.
	ClusterMask uint32 `idl:"name:ClusterMask" json:"cluster_mask"`
}

func (o *IPCluster) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPCluster) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ClusterAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.ClusterMask); err != nil {
		return err
	}
	return nil
}
func (o *IPCluster) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClusterAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClusterMask); err != nil {
		return err
	}
	return nil
}

// BOOTPIPRange structure represents DHCP_BOOTP_IP_RANGE RPC structure.
//
// The DHCP_BOOTP_IP_RANGE structure defines a suite of IPv4 addresses that can be leased
// to BOOTP-specific clients ([RFC2132]). This structure is an extension of the DHCP_IP_RANGE
// (section 2.2.1.2.31) structure with some additional information for BOOTP-specific
// clients.
type BOOTPIPRange struct {
	// StartAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing the
	// first IPv4 address in the IPv4 range defined on the DHCPv4 server for BOOTP clients.
	StartAddress uint32 `idl:"name:StartAddress" json:"start_address"`
	// EndAddress:  This is of type DHCP_IP_ADDRESS, containing the last IPv4 address in
	// the IPv4 range defined on the DHCPv4 server for BOOTP clients.
	EndAddress uint32 `idl:"name:EndAddress" json:"end_address"`
	// BootpAllocated:  This is of type ULONG, specifying the number of BOOTP clients that
	// have been served from this IPv4 range.
	BOOTPAllocated uint32 `idl:"name:BootpAllocated" json:"bootp_allocated"`
	// MaxBootpAllowed:  This is of type ULONG, specifying the maximum count of BOOTP clients
	// in this IPv4 range that the DHCPv4 server is allowed to serve.
	MaxBOOTPAllowed uint32 `idl:"name:MaxBootpAllowed" json:"max_bootp_allowed"`
}

func (o *BOOTPIPRange) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BOOTPIPRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.StartAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.EndAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.BOOTPAllocated); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxBOOTPAllowed); err != nil {
		return err
	}
	return nil
}
func (o *BOOTPIPRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.StartAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.EndAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.BOOTPAllocated); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxBOOTPAllowed); err != nil {
		return err
	}
	return nil
}

// SubnetElementData structure represents DHCP_SUBNET_ELEMENT_DATA RPC structure.
//
// The DHCP_SUBNET_ELEMENT_DATA structure defines the elements of an IPv4 reservation,
// IPv4 exclusion range, or IPv4 range for the subnet. This structure is used in methods
// R_DhcpAddSubnetElement (section 3.1.4.5) and R_DhcpRemoveSubnetElement (section 3.1.4.7).
type SubnetElementData struct {
	// ElementType:  This is of type DHCP_SUBNET_ELEMENT_TYPE (section 2.2.1.1.7) enumeration,
	// defining the set of possible subnet element types. This value defines which of the
	// values is chosen from the subsequent union Element member.
	ElementType SubnetElementType `idl:"name:ElementType" json:"element_type"`
	// Element:  Element is a union of subnet elements. The value of the union is dependent
	// on the previous field the ElementType member.
	Element *SubnetElementData_Element `idl:"name:Element;switch_is:(((ElementType 7 <=) (5 ElementType <=) &&) 0 ElementType ?:)" json:"element"`
}

func (o *SubnetElementData) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ElementType)); err != nil {
		return err
	}
	_exprElementType := uint16(0)
	if (o.ElementType <= 7) && (5 <= o.ElementType) {
		_exprElementType = uint16(0)
	} else {
		_exprElementType = uint16(o.ElementType)
	}
	_swElement := uint16(_exprElementType)
	if o.Element != nil {
		if err := o.Element.MarshalUnionNDR(ctx, w, _swElement); err != nil {
			return err
		}
	} else {
		if err := (&SubnetElementData_Element{}).MarshalUnionNDR(ctx, w, _swElement); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ElementType)); err != nil {
		return err
	}
	if o.Element == nil {
		o.Element = &SubnetElementData_Element{}
	}
	_exprElementType := uint16(0)
	if (o.ElementType <= 7) && (5 <= o.ElementType) {
		_exprElementType = uint16(0)
	} else {
		_exprElementType = uint16(o.ElementType)
	}
	_swElement := uint16(_exprElementType)
	if err := o.Element.UnmarshalUnionNDR(ctx, w, _swElement); err != nil {
		return err
	}
	return nil
}

// SubnetElementData_Element structure represents DHCP_SUBNET_ELEMENT_DATA union anonymous member.
//
// The DHCP_SUBNET_ELEMENT_DATA structure defines the elements of an IPv4 reservation,
// IPv4 exclusion range, or IPv4 range for the subnet. This structure is used in methods
// R_DhcpAddSubnetElement (section 3.1.4.5) and R_DhcpRemoveSubnetElement (section 3.1.4.7).
type SubnetElementData_Element struct {
	// Types that are assignable to Value
	//
	// *SubnetElementData_IPRange
	// *SubnetElementData_SecondaryHost
	// *SubnetElementData_ReservedIP
	// *SubnetElementData_ExcludeIPRange
	// *SubnetElementData_IPUsedCluster
	Value is_SubnetElementData_Element `json:"value"`
}

func (o *SubnetElementData_Element) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SubnetElementData_IPRange:
		if value != nil {
			return value.IPRange
		}
	case *SubnetElementData_SecondaryHost:
		if value != nil {
			return value.SecondaryHost
		}
	case *SubnetElementData_ReservedIP:
		if value != nil {
			return value.ReservedIP
		}
	case *SubnetElementData_ExcludeIPRange:
		if value != nil {
			return value.ExcludeIPRange
		}
	case *SubnetElementData_IPUsedCluster:
		if value != nil {
			return value.IPUsedCluster
		}
	}
	return nil
}

type is_SubnetElementData_Element interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SubnetElementData_Element()
}

func (o *SubnetElementData_Element) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SubnetElementData_IPRange:
		return uint16(0)
	case *SubnetElementData_SecondaryHost:
		return uint16(1)
	case *SubnetElementData_ReservedIP:
		return uint16(2)
	case *SubnetElementData_ExcludeIPRange:
		return uint16(3)
	case *SubnetElementData_IPUsedCluster:
		return uint16(4)
	}
	return uint16(0)
}

func (o *SubnetElementData_Element) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*SubnetElementData_IPRange)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementData_IPRange{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*SubnetElementData_SecondaryHost)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementData_SecondaryHost{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*SubnetElementData_ReservedIP)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementData_ReservedIP{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*SubnetElementData_ExcludeIPRange)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementData_ExcludeIPRange{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*SubnetElementData_IPUsedCluster)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementData_IPUsedCluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SubnetElementData_Element) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &SubnetElementData_IPRange{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &SubnetElementData_SecondaryHost{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &SubnetElementData_ReservedIP{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &SubnetElementData_ExcludeIPRange{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &SubnetElementData_IPUsedCluster{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SubnetElementData_IPRange structure represents SubnetElementData_Element RPC union arm.
//
// It has following labels: 0
type SubnetElementData_IPRange struct {
	// IpRange:  This is of type DHCP_IP_RANGE (section 2.2.1.2.31), containing the IPv4
	// range for the IPv4 subnet. This contains the range for the following valid enumeration
	// values.
	//
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	|                          |                                                                                  |
	//	| DHCP SUBNET ELEMENT TYPE |                                     MEANING                                      |
	//	|                          |                                                                                  |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRanges 0           | The configuration parameter is the IP range of a DHCPv4 scope configured on the  |
	//	|                          | DHCP server.                                                                     |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRangesDhcpOnly 5   | The configuration parameter is an IP range of a DHCPv4 scope configured on the   |
	//	|                          | DHCP server that MUST be used only for assignment of addresses to DHCP clients   |
	//	|                          | on the subnet. The IP addresses from this range MUST NOT be assigned to BOOTP    |
	//	|                          | clients ([RFC2132]).                                                             |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRangesDhcpBootp 6  | The configuration parameter is an IP range of a DHCPv4 scope configured on the   |
	//	|                          | DHCP server that can be used for assignment of addresses to both DHCP and BOOTP. |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRangesBootpOnly 7  | The configuration parameter is an IP range of a DHCPv4 scope configured on the   |
	//	|                          | DHCP server that MUST be used only for assignment of IPv4 addresses to BOOTP     |
	//	|                          | clients.                                                                         |
	//	+--------------------------+----------------------------------------------------------------------------------+
	IPRange *IPRange `idl:"name:IpRange" json:"ip_range"`
}

func (*SubnetElementData_IPRange) is_SubnetElementData_Element() {}

func (o *SubnetElementData_IPRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IPRange != nil {
		_ptr_IpRange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPRange != nil {
				if err := o.IPRange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPRange{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPRange, _ptr_IpRange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementData_IPRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_IpRange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPRange == nil {
			o.IPRange = &IPRange{}
		}
		if err := o.IPRange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IpRange := func(ptr interface{}) { o.IPRange = *ptr.(**IPRange) }
	if err := w.ReadPointer(&o.IPRange, _s_IpRange, _ptr_IpRange); err != nil {
		return err
	}
	return nil
}

// SubnetElementData_SecondaryHost structure represents SubnetElementData_Element RPC union arm.
//
// It has following labels: 1
type SubnetElementData_SecondaryHost struct {
	// SecondaryHost:  This is of type DHCP_HOST_INFO (section 2.2.1.2.7) and is not used.
	// If the ElementType value mandates that the SecondaryHost element is to be used in
	// any method, that method will return ERROR_CALL_NOT_IMPLEMENTED or ERROR_NOT_SUPPORTED,
	// as specified in the processing rules of methods that use the DHCP_SUBNET_ELEMENT_DATA
	// structure.
	SecondaryHost *HostInfo `idl:"name:SecondaryHost" json:"secondary_host"`
}

func (*SubnetElementData_SecondaryHost) is_SubnetElementData_Element() {}

func (o *SubnetElementData_SecondaryHost) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SecondaryHost != nil {
		_ptr_SecondaryHost := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SecondaryHost != nil {
				if err := o.SecondaryHost.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecondaryHost, _ptr_SecondaryHost); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementData_SecondaryHost) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SecondaryHost := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SecondaryHost == nil {
			o.SecondaryHost = &HostInfo{}
		}
		if err := o.SecondaryHost.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SecondaryHost := func(ptr interface{}) { o.SecondaryHost = *ptr.(**HostInfo) }
	if err := w.ReadPointer(&o.SecondaryHost, _s_SecondaryHost, _ptr_SecondaryHost); err != nil {
		return err
	}
	return nil
}

// SubnetElementData_ReservedIP structure represents SubnetElementData_Element RPC union arm.
//
// It has following labels: 2
type SubnetElementData_ReservedIP struct {
	// ReservedIp:  This is of type DHCP_IP_RESERVATION (section 2.2.1.2.10), containing
	// the IPv4 reservation.
	ReservedIP *IPReservation `idl:"name:ReservedIp" json:"reserved_ip"`
}

func (*SubnetElementData_ReservedIP) is_SubnetElementData_Element() {}

func (o *SubnetElementData_ReservedIP) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ReservedIP != nil {
		_ptr_ReservedIp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ReservedIP != nil {
				if err := o.ReservedIP.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPReservation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReservedIP, _ptr_ReservedIp); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementData_ReservedIP) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ReservedIp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ReservedIP == nil {
			o.ReservedIP = &IPReservation{}
		}
		if err := o.ReservedIP.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ReservedIp := func(ptr interface{}) { o.ReservedIP = *ptr.(**IPReservation) }
	if err := w.ReadPointer(&o.ReservedIP, _s_ReservedIp, _ptr_ReservedIp); err != nil {
		return err
	}
	return nil
}

// SubnetElementData_ExcludeIPRange structure represents SubnetElementData_Element RPC union arm.
//
// It has following labels: 3
type SubnetElementData_ExcludeIPRange struct {
	// ExcludeIpRange:  This is of type DHCP_IP_RANGE, containing the IPv4 exclusion range.
	ExcludeIPRange *IPRange `idl:"name:ExcludeIpRange" json:"exclude_ip_range"`
}

func (*SubnetElementData_ExcludeIPRange) is_SubnetElementData_Element() {}

func (o *SubnetElementData_ExcludeIPRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ExcludeIPRange != nil {
		_ptr_ExcludeIpRange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ExcludeIPRange != nil {
				if err := o.ExcludeIPRange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPRange{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ExcludeIPRange, _ptr_ExcludeIpRange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementData_ExcludeIPRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ExcludeIpRange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ExcludeIPRange == nil {
			o.ExcludeIPRange = &IPRange{}
		}
		if err := o.ExcludeIPRange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ExcludeIpRange := func(ptr interface{}) { o.ExcludeIPRange = *ptr.(**IPRange) }
	if err := w.ReadPointer(&o.ExcludeIPRange, _s_ExcludeIpRange, _ptr_ExcludeIpRange); err != nil {
		return err
	}
	return nil
}

// SubnetElementData_IPUsedCluster structure represents SubnetElementData_Element RPC union arm.
//
// It has following labels: 4
type SubnetElementData_IPUsedCluster struct {
	// IpUsedCluster:  This is of type DHCP_IP_CLUSTER (section 2.2.1.2.88) and is not used.
	// If the ElementType member mandates this element to be used in any method, the method
	// will return ERROR_INVALID_PARAMETER.
	IPUsedCluster *IPCluster `idl:"name:IpUsedCluster" json:"ip_used_cluster"`
}

func (*SubnetElementData_IPUsedCluster) is_SubnetElementData_Element() {}

func (o *SubnetElementData_IPUsedCluster) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IPUsedCluster != nil {
		_ptr_IpUsedCluster := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPUsedCluster != nil {
				if err := o.IPUsedCluster.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPCluster{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPUsedCluster, _ptr_IpUsedCluster); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementData_IPUsedCluster) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_IpUsedCluster := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPUsedCluster == nil {
			o.IPUsedCluster = &IPCluster{}
		}
		if err := o.IPUsedCluster.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IpUsedCluster := func(ptr interface{}) { o.IPUsedCluster = *ptr.(**IPCluster) }
	if err := w.ReadPointer(&o.IPUsedCluster, _s_IpUsedCluster, _ptr_IpUsedCluster); err != nil {
		return err
	}
	return nil
}

// SubnetElementInfoArray structure represents DHCP_SUBNET_ELEMENT_INFO_ARRAY RPC structure.
//
// The DHCP_SUBNET_ELEMENT_INFO_ARRAY structure defines an array of DHCP_SUBNET_ELEMENT_DATA
// (section 2.2.1.2.33) structures. The first parameter contains the number of subnet
// elements (IPv4 reservation, IPv4 exclusion range, and IPv4 range), and the second
// parameter points to the array of length NumElements containing DHCP subnet elements.
// This structure is used in the R_DhcpEnumSubnetElements (section 3.1.4.6) method.
type SubnetElementInfoArray struct {
	// NumElements:  This is of type DWORD, containing the number of subnet elements in
	// the subsequent field the Elements member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This is a pointer to an array of DHCP_SUBNET_ELEMENT_DATA structures of
	// length NumElements containing IPv4 subnet elements.
	Elements []*SubnetElementData `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *SubnetElementInfoArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementInfoArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SubnetElementData{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SubnetElementData{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementInfoArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*SubnetElementData, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &SubnetElementData{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*SubnetElementData) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// DwordDword structure represents DWORD_DWORD RPC structure.
//
// The DWORD_DWORD structure defines a 64-bit integer value. This is used in DHCP_OPTION_DATA_ELEMENT
// (section 2.2.1.2.23).
type DwordDword struct {
	// DWord1:  This is of type DWORD, specifying the upper 32 bits of the value.
	Dword1 uint32 `idl:"name:DWord1" json:"dword1"`
	// DWord2:  This is of type DWORD, specifying the lower 32 bits of the value.
	Dword2 uint32 `idl:"name:DWord2" json:"dword2"`
}

func (o *DwordDword) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DwordDword) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Dword1); err != nil {
		return err
	}
	if err := w.WriteData(o.Dword2); err != nil {
		return err
	}
	return nil
}
func (o *DwordDword) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Dword1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Dword2); err != nil {
		return err
	}
	return nil
}

// OptionDataElement structure represents DHCP_OPTION_DATA_ELEMENT RPC structure.
//
// The DHCP_OPTION_DATA_ELEMENT structure contains the type of the option and its data
// value. This is used within a DHCP_OPTION_DATA (section 2.2.1.2.24) structure.
type OptionDataElement struct {
	// OptionType:  This is of type DHCP_OPTION_DATA_TYPE (section 2.2.1.1.10) enumeration
	// value, indicating the option value that is present in the subsequent field, Element.
	OptionType OptionDataType `idl:"name:OptionType" json:"option_type"`
	// Element:  This is a union that can contain one of the following values chosen based
	// on the value of OptionType.
	Element *OptionDataElement_Element `idl:"name:Element;switch_is:OptionType" json:"element"`
}

func (o *OptionDataElement) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionDataElement) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.OptionType)); err != nil {
		return err
	}
	_swElement := uint16(o.OptionType)
	if o.Element != nil {
		if err := o.Element.MarshalUnionNDR(ctx, w, _swElement); err != nil {
			return err
		}
	} else {
		if err := (&OptionDataElement_Element{}).MarshalUnionNDR(ctx, w, _swElement); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionDataElement) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.OptionType)); err != nil {
		return err
	}
	if o.Element == nil {
		o.Element = &OptionDataElement_Element{}
	}
	_swElement := uint16(o.OptionType)
	if err := o.Element.UnmarshalUnionNDR(ctx, w, _swElement); err != nil {
		return err
	}
	return nil
}

// OptionDataElement_Element structure represents DHCP_OPTION_DATA_ELEMENT union anonymous member.
//
// The DHCP_OPTION_DATA_ELEMENT structure contains the type of the option and its data
// value. This is used within a DHCP_OPTION_DATA (section 2.2.1.2.24) structure.
type OptionDataElement_Element struct {
	// Types that are assignable to Value
	//
	// *OptionDataElement_ByteOption
	// *OptionDataElement_WordOption
	// *OptionDataElement_DwordOption
	// *OptionDataElement_DwordDwordOption
	// *OptionDataElement_IPAddressOption
	// *OptionDataElement_StringDataOption
	// *OptionDataElement_BinaryDataOption
	// *OptionDataElement_EncapsulatedDataOption
	// *OptionDataElement_IPv6AddressDataOption
	Value is_OptionDataElement_Element `json:"value"`
}

func (o *OptionDataElement_Element) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *OptionDataElement_ByteOption:
		if value != nil {
			return value.ByteOption
		}
	case *OptionDataElement_WordOption:
		if value != nil {
			return value.WordOption
		}
	case *OptionDataElement_DwordOption:
		if value != nil {
			return value.DwordOption
		}
	case *OptionDataElement_DwordDwordOption:
		if value != nil {
			return value.DwordDwordOption
		}
	case *OptionDataElement_IPAddressOption:
		if value != nil {
			return value.IPAddressOption
		}
	case *OptionDataElement_StringDataOption:
		if value != nil {
			return value.StringDataOption
		}
	case *OptionDataElement_BinaryDataOption:
		if value != nil {
			return value.BinaryDataOption
		}
	case *OptionDataElement_EncapsulatedDataOption:
		if value != nil {
			return value.EncapsulatedDataOption
		}
	case *OptionDataElement_IPv6AddressDataOption:
		if value != nil {
			return value.IPv6AddressDataOption
		}
	}
	return nil
}

type is_OptionDataElement_Element interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_OptionDataElement_Element()
}

func (o *OptionDataElement_Element) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *OptionDataElement_ByteOption:
		return uint16(0)
	case *OptionDataElement_WordOption:
		return uint16(1)
	case *OptionDataElement_DwordOption:
		return uint16(2)
	case *OptionDataElement_DwordDwordOption:
		return uint16(3)
	case *OptionDataElement_IPAddressOption:
		return uint16(4)
	case *OptionDataElement_StringDataOption:
		return uint16(5)
	case *OptionDataElement_BinaryDataOption:
		return uint16(6)
	case *OptionDataElement_EncapsulatedDataOption:
		return uint16(7)
	case *OptionDataElement_IPv6AddressDataOption:
		return uint16(8)
	}
	return uint16(0)
}

func (o *OptionDataElement_Element) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*OptionDataElement_ByteOption)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionDataElement_ByteOption{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*OptionDataElement_WordOption)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionDataElement_WordOption{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*OptionDataElement_DwordOption)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionDataElement_DwordOption{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*OptionDataElement_DwordDwordOption)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionDataElement_DwordDwordOption{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*OptionDataElement_IPAddressOption)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionDataElement_IPAddressOption{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(5):
		_o, _ := o.Value.(*OptionDataElement_StringDataOption)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionDataElement_StringDataOption{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(6):
		_o, _ := o.Value.(*OptionDataElement_BinaryDataOption)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionDataElement_BinaryDataOption{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(7):
		_o, _ := o.Value.(*OptionDataElement_EncapsulatedDataOption)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionDataElement_EncapsulatedDataOption{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(8):
		_o, _ := o.Value.(*OptionDataElement_IPv6AddressDataOption)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionDataElement_IPv6AddressDataOption{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *OptionDataElement_Element) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &OptionDataElement_ByteOption{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &OptionDataElement_WordOption{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &OptionDataElement_DwordOption{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &OptionDataElement_DwordDwordOption{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &OptionDataElement_IPAddressOption{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(5):
		o.Value = &OptionDataElement_StringDataOption{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(6):
		o.Value = &OptionDataElement_BinaryDataOption{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(7):
		o.Value = &OptionDataElement_EncapsulatedDataOption{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(8):
		o.Value = &OptionDataElement_IPv6AddressDataOption{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// OptionDataElement_ByteOption structure represents OptionDataElement_Element RPC union arm.
//
// It has following labels: 0
type OptionDataElement_ByteOption struct {
	// ByteOption:  Specifies the data as a BYTE value. This field is present if the OptionType
	// is DhcpByteOption.
	ByteOption uint8 `idl:"name:ByteOption" json:"byte_option"`
}

func (*OptionDataElement_ByteOption) is_OptionDataElement_Element() {}

func (o *OptionDataElement_ByteOption) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.ByteOption); err != nil {
		return err
	}
	return nil
}
func (o *OptionDataElement_ByteOption) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.ByteOption); err != nil {
		return err
	}
	return nil
}

// OptionDataElement_WordOption structure represents OptionDataElement_Element RPC union arm.
//
// It has following labels: 1
type OptionDataElement_WordOption struct {
	// WordOption:  Specifies the data as a WORD value. This field is present if the OptionType
	// is DhcpWordOption.
	WordOption uint16 `idl:"name:WordOption" json:"word_option"`
}

func (*OptionDataElement_WordOption) is_OptionDataElement_Element() {}

func (o *OptionDataElement_WordOption) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.WordOption); err != nil {
		return err
	}
	return nil
}
func (o *OptionDataElement_WordOption) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.WordOption); err != nil {
		return err
	}
	return nil
}

// OptionDataElement_DwordOption structure represents OptionDataElement_Element RPC union arm.
//
// It has following labels: 2
type OptionDataElement_DwordOption struct {
	// DWordOption:  Specifies the data as a DWORD value. This field is present if the OptionType
	// is DhcpDWordOption.
	DwordOption uint32 `idl:"name:DWordOption" json:"dword_option"`
}

func (*OptionDataElement_DwordOption) is_OptionDataElement_Element() {}

func (o *OptionDataElement_DwordOption) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.DwordOption); err != nil {
		return err
	}
	return nil
}
func (o *OptionDataElement_DwordOption) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.DwordOption); err != nil {
		return err
	}
	return nil
}

// OptionDataElement_DwordDwordOption structure represents OptionDataElement_Element RPC union arm.
//
// It has following labels: 3
type OptionDataElement_DwordDwordOption struct {
	// DWordDWordOption:  Specifies the data as a DWORD_DWORD (section 2.2.1.2.22) value.
	// This field is present if the OptionType is DhcpDWordDWordOption.
	DwordDwordOption *DwordDword `idl:"name:DWordDWordOption" json:"dword_dword_option"`
}

func (*OptionDataElement_DwordDwordOption) is_OptionDataElement_Element() {}

func (o *OptionDataElement_DwordDwordOption) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DwordDwordOption != nil {
		if err := o.DwordDwordOption.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DwordDword{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionDataElement_DwordDwordOption) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DwordDwordOption == nil {
		o.DwordDwordOption = &DwordDword{}
	}
	if err := o.DwordDwordOption.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// OptionDataElement_IPAddressOption structure represents OptionDataElement_Element RPC union arm.
//
// It has following labels: 4
type OptionDataElement_IPAddressOption struct {
	// IpAddressOption:  Specifies the data as a DHCP_IP_ADDRESS (section 2.2.1.2.1) value.
	// This field is present if the OptionType is IpAddressOption.
	IPAddressOption uint32 `idl:"name:IpAddressOption" json:"ip_address_option"`
}

func (*OptionDataElement_IPAddressOption) is_OptionDataElement_Element() {}

func (o *OptionDataElement_IPAddressOption) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.IPAddressOption); err != nil {
		return err
	}
	return nil
}
func (o *OptionDataElement_IPAddressOption) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.IPAddressOption); err != nil {
		return err
	}
	return nil
}

// OptionDataElement_StringDataOption structure represents OptionDataElement_Element RPC union arm.
//
// It has following labels: 5
type OptionDataElement_StringDataOption struct {
	// StringDataOption:  Specifies the data as LPWSTR, a pointer to a Unicode string value.
	// This field is present if the OptionType is DhcpStringDataOption.
	StringDataOption string `idl:"name:StringDataOption" json:"string_data_option"`
}

func (*OptionDataElement_StringDataOption) is_OptionDataElement_Element() {}

func (o *OptionDataElement_StringDataOption) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StringDataOption != "" {
		_ptr_StringDataOption := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.StringDataOption); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.StringDataOption, _ptr_StringDataOption); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionDataElement_StringDataOption) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_StringDataOption := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.StringDataOption); err != nil {
			return err
		}
		return nil
	})
	_s_StringDataOption := func(ptr interface{}) { o.StringDataOption = *ptr.(*string) }
	if err := w.ReadPointer(&o.StringDataOption, _s_StringDataOption, _ptr_StringDataOption); err != nil {
		return err
	}
	return nil
}

// OptionDataElement_BinaryDataOption structure represents OptionDataElement_Element RPC union arm.
//
// It has following labels: 6
type OptionDataElement_BinaryDataOption struct {
	// BinaryDataOption:  Specifies the data as a DHCP_BINARY_DATA (section 2.2.1.2.9) structure.
	// This field is present if the OptionType is DhcpBinaryDataOption.
	BinaryDataOption *BinaryData `idl:"name:BinaryDataOption" json:"binary_data_option"`
}

func (*OptionDataElement_BinaryDataOption) is_OptionDataElement_Element() {}

func (o *OptionDataElement_BinaryDataOption) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.BinaryDataOption != nil {
		if err := o.BinaryDataOption.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&BinaryData{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionDataElement_BinaryDataOption) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.BinaryDataOption == nil {
		o.BinaryDataOption = &BinaryData{}
	}
	if err := o.BinaryDataOption.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// OptionDataElement_EncapsulatedDataOption structure represents OptionDataElement_Element RPC union arm.
//
// It has following labels: 7
type OptionDataElement_EncapsulatedDataOption struct {
	// EncapsulatedDataOption:  Specifies the data as encapsulated within a DHCP_BINARY_DATA
	// structure. The application MUST recognize the format of the opaque data capsule in
	// order to read it from the Data field of DHCP_BINARY_DATA. This field is present if
	// the OptionType is DhcpEncapsulatedDataOption.
	EncapsulatedDataOption *BinaryData `idl:"name:EncapsulatedDataOption" json:"encapsulated_data_option"`
}

func (*OptionDataElement_EncapsulatedDataOption) is_OptionDataElement_Element() {}

func (o *OptionDataElement_EncapsulatedDataOption) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.EncapsulatedDataOption != nil {
		if err := o.EncapsulatedDataOption.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&BinaryData{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionDataElement_EncapsulatedDataOption) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.EncapsulatedDataOption == nil {
		o.EncapsulatedDataOption = &BinaryData{}
	}
	if err := o.EncapsulatedDataOption.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// OptionDataElement_IPv6AddressDataOption structure represents OptionDataElement_Element RPC union arm.
//
// It has following labels: 8
type OptionDataElement_IPv6AddressDataOption struct {
	// Ipv6AddressDataOption:  Specifies the data as LPWSTR, a pointer to a Unicode string
	// value. This field is present if the OptionType is DhcpIpv6AddressOption.
	IPv6AddressDataOption string `idl:"name:Ipv6AddressDataOption" json:"ipv6_address_data_option"`
}

func (*OptionDataElement_IPv6AddressDataOption) is_OptionDataElement_Element() {}

func (o *OptionDataElement_IPv6AddressDataOption) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IPv6AddressDataOption != "" {
		_ptr_Ipv6AddressDataOption := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.IPv6AddressDataOption); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.IPv6AddressDataOption, _ptr_Ipv6AddressDataOption); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionDataElement_IPv6AddressDataOption) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Ipv6AddressDataOption := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.IPv6AddressDataOption); err != nil {
			return err
		}
		return nil
	})
	_s_Ipv6AddressDataOption := func(ptr interface{}) { o.IPv6AddressDataOption = *ptr.(*string) }
	if err := w.ReadPointer(&o.IPv6AddressDataOption, _s_Ipv6AddressDataOption, _ptr_Ipv6AddressDataOption); err != nil {
		return err
	}
	return nil
}

// OptionData structure represents DHCP_OPTION_DATA RPC structure.
//
// The DHCP_OPTION_DATA structure defines an array of DHCP_OPTION_DATA_ELEMENT (section
// 2.2.1.2.23) structures. This structure is a data container for one or more data elements
// associated with a DHCP option. This structure is used in the DHCP_OPTION_VALUE (section
// 2.2.1.2.42) structure.
type OptionData struct {
	// NumElements:  This is of type DWORD, specifying the number of data elements in the
	// specific DHCP option, which is also the number of option data elements listed in
	// the Elements array member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This is a pointer of type DHCP_OPTION_DATA_ELEMENT structure that points
	// to the array of length NumElements containing the data elements associated with a
	// specific option.
	Elements []*OptionDataElement `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *OptionData) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&OptionDataElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&OptionDataElement{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*OptionDataElement, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &OptionDataElement{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*OptionDataElement) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// Option structure represents DHCP_OPTION RPC structure.
//
// The DHCP_OPTION structure contains the information for an option definition created
// on the DHCP server. This structure is used in the LPDHCP_OPTION_ARRAY (section 2.2.1.2.26)
// structure.
type Option struct {
	// OptionID:  This is of type DHCP_OPTION_ID (section 2.2.1.2.3), containing a value
	// that uniquely identifies the option.
	OptionID uint32 `idl:"name:OptionID" json:"option_id"`
	// OptionName:  A pointer of type LPWSTR to a null-terminated Unicode string that specifies
	// the option name of the option. There is no restriction on the length of this Unicode
	// string.
	OptionName string `idl:"name:OptionName" json:"option_name"`
	// OptionComment:  A pointer of type LPWSTR to a null-terminated Unicode string that
	// specifies a comment for the option. This is an optional parameter. There is no restriction
	// on the length of this Unicode string.
	OptionComment string `idl:"name:OptionComment" json:"option_comment"`
	// DefaultValue:  This is of type DHCP_OPTION_DATA (section 2.2.1.2.24), containing
	// the default value for the option. This also defines the data type used to store the
	// value of the option.
	DefaultValue *OptionData `idl:"name:DefaultValue" json:"default_value"`
	// OptionType:  This is of type DHCP_OPTION_TYPE (section 2.2.1.1.6), indicating whether
	// the default value is a unary item or an array of elements.
	OptionType OptionType `idl:"name:OptionType" json:"option_type"`
}

func (o *Option) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Option) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.OptionID); err != nil {
		return err
	}
	if o.OptionName != "" {
		_ptr_OptionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.OptionName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.OptionName, _ptr_OptionName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.OptionComment != "" {
		_ptr_OptionComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.OptionComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.OptionComment, _ptr_OptionComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DefaultValue != nil {
		if err := o.DefaultValue.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OptionData{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.OptionType)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *Option) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.OptionID); err != nil {
		return err
	}
	_ptr_OptionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.OptionName); err != nil {
			return err
		}
		return nil
	})
	_s_OptionName := func(ptr interface{}) { o.OptionName = *ptr.(*string) }
	if err := w.ReadPointer(&o.OptionName, _s_OptionName, _ptr_OptionName); err != nil {
		return err
	}
	_ptr_OptionComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.OptionComment); err != nil {
			return err
		}
		return nil
	})
	_s_OptionComment := func(ptr interface{}) { o.OptionComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.OptionComment, _s_OptionComment, _ptr_OptionComment); err != nil {
		return err
	}
	if o.DefaultValue == nil {
		o.DefaultValue = &OptionData{}
	}
	if err := o.DefaultValue.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.OptionType)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ReservedScope structure represents DHCP_RESERVED_SCOPE RPC structure.
//
// The DHCP_RESERVED_SCOPE structure defines an IPv4 reservation. This structure is
// used in the DHCP_OPTION_SCOPE_INFO (section 2.2.1.2.41) structure.
type ReservedScope struct {
	// ReservedIpAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing
	// the IPv4 address of the reservation.
	ReservedIPAddress uint32 `idl:"name:ReservedIpAddress" json:"reserved_ip_address"`
	// ReservedIpSubnetAddress:  This is of type DHCP_IP_ADDRESS, containing the IPv4 address
	// of the subnet ID.
	ReservedIPSubnetAddress uint32 `idl:"name:ReservedIpSubnetAddress" json:"reserved_ip_subnet_address"`
}

func (o *ReservedScope) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReservedScope) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ReservedIPAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.ReservedIPSubnetAddress); err != nil {
		return err
	}
	return nil
}
func (o *ReservedScope) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReservedIPAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReservedIPSubnetAddress); err != nil {
		return err
	}
	return nil
}

// OptionScopeInfo structure represents DHCP_OPTION_SCOPE_INFO RPC structure.
//
// The DHCP_OPTION_SCOPE_INFO structure defines the information about the option. The
// information consists of the option type and the level of the option (server level,
// scope level, or reservation level).
type OptionScopeInfo struct {
	// ScopeType:  This is of type DHCP_OPTION_SCOPE_TYPE (section 2.2.1.1.4) enumeration,
	// defining the scope of the DHCP option and indicating which of the following fields
	// in the union is used.
	ScopeType OptionScopeType `idl:"name:ScopeType" json:"scope_type"`
	// ScopeInfo:  This is a union from which one of the following fields is used based
	// on the value of the ScopeType member:
	ScopeInfo *OptionScopeInfo_ScopeInfo `idl:"name:ScopeInfo;switch_is:ScopeType" json:"scope_info"`
}

func (o *OptionScopeInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionScopeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ScopeType)); err != nil {
		return err
	}
	_swScopeInfo := uint16(o.ScopeType)
	if o.ScopeInfo != nil {
		if err := o.ScopeInfo.MarshalUnionNDR(ctx, w, _swScopeInfo); err != nil {
			return err
		}
	} else {
		if err := (&OptionScopeInfo_ScopeInfo{}).MarshalUnionNDR(ctx, w, _swScopeInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionScopeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ScopeType)); err != nil {
		return err
	}
	if o.ScopeInfo == nil {
		o.ScopeInfo = &OptionScopeInfo_ScopeInfo{}
	}
	_swScopeInfo := uint16(o.ScopeType)
	if err := o.ScopeInfo.UnmarshalUnionNDR(ctx, w, _swScopeInfo); err != nil {
		return err
	}
	return nil
}

// OptionScopeInfo_ScopeInfo structure represents DHCP_OPTION_SCOPE_INFO union anonymous member.
//
// The DHCP_OPTION_SCOPE_INFO structure defines the information about the option. The
// information consists of the option type and the level of the option (server level,
// scope level, or reservation level).
type OptionScopeInfo_ScopeInfo struct {
	// Types that are assignable to Value
	//
	// *OptionScopeInfo_DefaultOptions
	// *OptionScopeInfo_GlobalOptions
	// *OptionScopeInfo_SubnetScopeInfo
	// *OptionScopeInfo_ReservedScopeInfo
	// *OptionScopeInfo_MScopeInfo
	Value is_OptionScopeInfo_ScopeInfo `json:"value"`
}

func (o *OptionScopeInfo_ScopeInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *OptionScopeInfo_SubnetScopeInfo:
		if value != nil {
			return value.SubnetScopeInfo
		}
	case *OptionScopeInfo_ReservedScopeInfo:
		if value != nil {
			return value.ReservedScopeInfo
		}
	case *OptionScopeInfo_MScopeInfo:
		if value != nil {
			return value.MScopeInfo
		}
	}
	return nil
}

type is_OptionScopeInfo_ScopeInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_OptionScopeInfo_ScopeInfo()
}

func (o *OptionScopeInfo_ScopeInfo) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *OptionScopeInfo_DefaultOptions:
		return uint16(0)
	case *OptionScopeInfo_GlobalOptions:
		return uint16(1)
	case *OptionScopeInfo_SubnetScopeInfo:
		return uint16(2)
	case *OptionScopeInfo_ReservedScopeInfo:
		return uint16(3)
	case *OptionScopeInfo_MScopeInfo:
		return uint16(4)
	}
	return uint16(0)
}

func (o *OptionScopeInfo_ScopeInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
	case uint16(1):
	case uint16(2):
		_o, _ := o.Value.(*OptionScopeInfo_SubnetScopeInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionScopeInfo_SubnetScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*OptionScopeInfo_ReservedScopeInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionScopeInfo_ReservedScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*OptionScopeInfo_MScopeInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionScopeInfo_MScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *OptionScopeInfo_ScopeInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &OptionScopeInfo_DefaultOptions{}
	case uint16(1):
		o.Value = &OptionScopeInfo_GlobalOptions{}
	case uint16(2):
		o.Value = &OptionScopeInfo_SubnetScopeInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &OptionScopeInfo_ReservedScopeInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &OptionScopeInfo_MScopeInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// OptionScopeInfo_DefaultOptions structure represents OptionScopeInfo_ScopeInfo RPC union arm.
//
// It has following labels: 0
type OptionScopeInfo_DefaultOptions struct {
}

func (*OptionScopeInfo_DefaultOptions) is_OptionScopeInfo_ScopeInfo() {}

func (o *OptionScopeInfo_DefaultOptions) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *OptionScopeInfo_DefaultOptions) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// OptionScopeInfo_GlobalOptions structure represents OptionScopeInfo_ScopeInfo RPC union arm.
//
// It has following labels: 1
type OptionScopeInfo_GlobalOptions struct {
}

func (*OptionScopeInfo_GlobalOptions) is_OptionScopeInfo_ScopeInfo() {}

func (o *OptionScopeInfo_GlobalOptions) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *OptionScopeInfo_GlobalOptions) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// OptionScopeInfo_SubnetScopeInfo structure represents OptionScopeInfo_ScopeInfo RPC union arm.
//
// It has following labels: 2
type OptionScopeInfo_SubnetScopeInfo struct {
	// SubnetScopeInfo:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), containing
	// the IPv4 subnet ID as a DWORD.
	SubnetScopeInfo uint32 `idl:"name:SubnetScopeInfo" json:"subnet_scope_info"`
}

func (*OptionScopeInfo_SubnetScopeInfo) is_OptionScopeInfo_ScopeInfo() {}

func (o *OptionScopeInfo_SubnetScopeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.SubnetScopeInfo); err != nil {
		return err
	}
	return nil
}
func (o *OptionScopeInfo_SubnetScopeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.SubnetScopeInfo); err != nil {
		return err
	}
	return nil
}

// OptionScopeInfo_ReservedScopeInfo structure represents OptionScopeInfo_ScopeInfo RPC union arm.
//
// It has following labels: 3
type OptionScopeInfo_ReservedScopeInfo struct {
	// ReservedScopeInfo:  This is a DHCP_RESERVED_SCOPE (section 2.2.1.2.40) structure
	// that contains an IPv4 reservation and its corresponding IPv4 subnet ID.
	ReservedScopeInfo *ReservedScope `idl:"name:ReservedScopeInfo" json:"reserved_scope_info"`
}

func (*OptionScopeInfo_ReservedScopeInfo) is_OptionScopeInfo_ScopeInfo() {}

func (o *OptionScopeInfo_ReservedScopeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ReservedScopeInfo != nil {
		if err := o.ReservedScopeInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReservedScope{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionScopeInfo_ReservedScopeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ReservedScopeInfo == nil {
		o.ReservedScopeInfo = &ReservedScope{}
	}
	if err := o.ReservedScopeInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// OptionScopeInfo_MScopeInfo structure represents OptionScopeInfo_ScopeInfo RPC union arm.
//
// It has following labels: 4
type OptionScopeInfo_MScopeInfo struct {
	// MScopeInfo:  This is a pointer to a null-terminated Unicode string that contains
	// the multicast scope name.
	MScopeInfo string `idl:"name:MScopeInfo" json:"mscope_info"`
}

func (*OptionScopeInfo_MScopeInfo) is_OptionScopeInfo_ScopeInfo() {}

func (o *OptionScopeInfo_MScopeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MScopeInfo != "" {
		_ptr_MScopeInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.MScopeInfo); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.MScopeInfo, _ptr_MScopeInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionScopeInfo_MScopeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_MScopeInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.MScopeInfo); err != nil {
			return err
		}
		return nil
	})
	_s_MScopeInfo := func(ptr interface{}) { o.MScopeInfo = *ptr.(*string) }
	if err := w.ReadPointer(&o.MScopeInfo, _s_MScopeInfo, _ptr_MScopeInfo); err != nil {
		return err
	}
	return nil
}

// OptionValue structure represents DHCP_OPTION_VALUE RPC structure.
//
// The DHCP_OPTION_VALUE structure contains the option identifier and its option value.
// This structure is used in the DHCP_OPTION_VALUE_ARRAY (section 2.2.1.2.43) structure.
type OptionValue struct {
	// OptionID:  This is of type DHCP_OPTION_ID (section 2.2.1.2.3), containing the identifier
	// for a specific option.
	OptionID uint32 `idl:"name:OptionID" json:"option_id"`
	// Value:  This is of type DHCP_OPTION_DATA (section 2.2.1.2.24), containing the option
	// value for an option.
	Value *OptionData `idl:"name:Value" json:"value"`
}

func (o *OptionValue) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.OptionID); err != nil {
		return err
	}
	if o.Value != nil {
		if err := o.Value.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OptionData{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.OptionID); err != nil {
		return err
	}
	if o.Value == nil {
		o.Value = &OptionData{}
	}
	if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// OptionValueArray structure represents DHCP_OPTION_VALUE_ARRAY RPC structure.
//
// The DHCP_OPTION_VALUE_ARRAY structure defines an array of DHCP_OPTION_VALUE (section
// 2.2.1.2.42) structures. This structure is used in the DHCP_ALL_OPTION_VALUES (section
// 2.2.1.2.44) structure. The first member contains the number of option values, and
// the second member points to the array of length NumElements containing option values.
type OptionValueArray struct {
	// NumElements:  This is a DWORD that specifies the number of option values in the subsequent
	// field the Values member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Values:  This is a pointer to an array of DHCP_OPTION_VALUE structures of length
	// NumElements that contains values.
	Values []*OptionValue `idl:"name:Values;size_is:(NumElements)" json:"values"`
}

func (o *OptionValueArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Values != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Values))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionValueArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Values != nil || o.ElementsLength > 0 {
		_ptr_Values := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Values {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Values[i1] != nil {
					if err := o.Values[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&OptionValue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Values); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&OptionValue{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Values, _ptr_Values); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionValueArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Values := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Values", sizeInfo[0])
		}
		o.Values = make([]*OptionValue, sizeInfo[0])
		for i1 := range o.Values {
			i1 := i1
			if o.Values[i1] == nil {
				o.Values[i1] = &OptionValue{}
			}
			if err := o.Values[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Values := func(ptr interface{}) { o.Values = *ptr.(*[]*OptionValue) }
	if err := w.ReadPointer(&o.Values, _s_Values, _ptr_Values); err != nil {
		return err
	}
	return nil
}

// DateTime structure represents DATE_TIME RPC structure.
//
// The DATE_TIME structure contains a 64-bit value interpreted as an unsigned number
// that represents the number of 100-nanosecond intervals since January 1, 1601 (UTC).
type DateTime struct {
	// dwLowDateTime:  This is of type DWORD, containing the lower 32 bits of the time value.
	LowDateTime uint32 `idl:"name:dwLowDateTime" json:"low_date_time"`
	// dwHighDateTime:  This is of type DWORD, containing the upper 32 bits of the time
	// value.
	HighDateTime uint32 `idl:"name:dwHighDateTime" json:"high_date_time"`
}

func (o *DateTime) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DateTime) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LowDateTime); err != nil {
		return err
	}
	if err := w.WriteData(o.HighDateTime); err != nil {
		return err
	}
	return nil
}
func (o *DateTime) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LowDateTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.HighDateTime); err != nil {
		return err
	}
	return nil
}

// ClientInfo structure represents DHCP_CLIENT_INFO RPC structure.
//
// The DHCP_CLIENT_INFO structure defines information about the DHCPv4 client that is
// used by the R_DhcpGetClientInfo (section 3.1.4.19) method.
type ClientInfo struct {
	// ClientIpAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1) and contains
	// the DHCPv4 client's IPv4 address.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// SubnetMask:  This is of type DHCP_IP_MASK (section 2.2.1.2.2) and contains the DHCPv4
	// client's IPv4 subnet mask address.
	SubnetMask uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
	// ClientHardwareAddress:  This is of type DHCP_CLIENT_UID (section 2.2.1.2.5) that
	// represents a DHCPv4 client-identifier (section 2.2.1.2.5.1) or a DHCPv4 client unique
	// ID (section 2.2.1.2.5.2). Methods that accept DHCP_CLIENT_INFO as a parameter specify
	// which representations are acceptable.
	ClientHardwareAddress *ClientUID `idl:"name:ClientHardwareAddress" json:"client_hardware_address"`
	// ClientName:  A pointer to a null-terminated Unicode string that represents the DHCPv4
	// client's internet host name. There is no restriction on the length of this Unicode
	// string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
	// ClientComment:  A pointer to a null-terminated Unicode string that represents a description
	// of the DHCPv4 client. There is no restriction on the length of this Unicode string.
	ClientComment string `idl:"name:ClientComment" json:"client_comment"`
	// ClientLeaseExpires:  This is of type DATE_TIME (section 2.2.1.2.11) that contains
	// the lease expiry time for the DHCPv4 client. This is Coordinated Universal Time (UTC).
	ClientLeaseExpires *DateTime `idl:"name:ClientLeaseExpires" json:"client_lease_expires"`
	// OwnerHost:  This is of type DHCP_HOST_INFO (section 2.2.1.2.7) that contains information
	// about the DHCPv4 server machine that has provided a lease to the DHCPv4 client.
	OwnerHost *HostInfo `idl:"name:OwnerHost" json:"owner_host"`
}

func (o *ClientInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress != nil {
		if err := o.ClientHardwareAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientComment != "" {
		_ptr_ClientComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientComment, _ptr_ClientComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientLeaseExpires != nil {
		if err := o.ClientLeaseExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OwnerHost != nil {
		if err := o.OwnerHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress == nil {
		o.ClientHardwareAddress = &ClientUID{}
	}
	if err := o.ClientHardwareAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	_ptr_ClientComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClientComment := func(ptr interface{}) { o.ClientComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientComment, _s_ClientComment, _ptr_ClientComment); err != nil {
		return err
	}
	if o.ClientLeaseExpires == nil {
		o.ClientLeaseExpires = &DateTime{}
	}
	if err := o.ClientLeaseExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OwnerHost == nil {
		o.OwnerHost = &HostInfo{}
	}
	if err := o.OwnerHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SearchInfo structure represents DHCP_SEARCH_INFO RPC structure.
//
// The DHCP_SEARCH_INFO structure defines the DHCPv4 client information search type
// defined by SearchType, along with the data used within that search. This structure,
// used in the R_DhcpGetClientInfo (section 3.1.4.19) method, is used to search a specific
// DHCPv4 client.
type SearchInfo struct {
	// SearchType:  This is an enumeration of type DHCP_SEARCH_INFO_TYPE (section 2.2.1.1.3)
	// that contains the data type based on which the search is performed for a specific
	// DHCPv4 client record held by the DHCPv4 server.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| DhcpClientIpAddress 0       | The DHCPv4 client IPv4 address (section 2.2.1.2.1), specified in a subsequent    |
	//	|                             | field, is used to search for the DHCPv4 client lease record in the DHCPv4 server |
	//	|                             | database.                                                                        |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| DhcpClientHardwareAddress 1 | The DHCPv4 client unique ID (section 2.2.1.2.5.2), specified in a subsequent     |
	//	|                             | field, is used to search for the DHCPv4 client lease record in the DHCPv4 server |
	//	|                             | database.                                                                        |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| DhcpClientName 2            | A pointer to a null-terminated Unicode string that contains the name of the      |
	//	|                             | DHCPv4 client. It is used to search for the DHCPv4 client lease record in the    |
	//	|                             | DHCPv4 server database.                                                          |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	SearchType SearchInfoType `idl:"name:SearchType" json:"search_type"`
	// SearchInfo:  A union that can contain one of the following values chosen based on
	// the value of SearchType.
	SearchInfo *SearchInfo_SearchInfo `idl:"name:SearchInfo;switch_is:SearchType" json:"search_info"`
}

func (o *SearchInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SearchInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.SearchType)); err != nil {
		return err
	}
	_swSearchInfo := uint16(o.SearchType)
	if o.SearchInfo != nil {
		if err := o.SearchInfo.MarshalUnionNDR(ctx, w, _swSearchInfo); err != nil {
			return err
		}
	} else {
		if err := (&SearchInfo_SearchInfo{}).MarshalUnionNDR(ctx, w, _swSearchInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *SearchInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.SearchType)); err != nil {
		return err
	}
	if o.SearchInfo == nil {
		o.SearchInfo = &SearchInfo_SearchInfo{}
	}
	_swSearchInfo := uint16(o.SearchType)
	if err := o.SearchInfo.UnmarshalUnionNDR(ctx, w, _swSearchInfo); err != nil {
		return err
	}
	return nil
}

// SearchInfo_SearchInfo structure represents DHCP_SEARCH_INFO union anonymous member.
//
// The DHCP_SEARCH_INFO structure defines the DHCPv4 client information search type
// defined by SearchType, along with the data used within that search. This structure,
// used in the R_DhcpGetClientInfo (section 3.1.4.19) method, is used to search a specific
// DHCPv4 client.
type SearchInfo_SearchInfo struct {
	// Types that are assignable to Value
	//
	// *SearchInfo_ClientIPAddress
	// *SearchInfo_ClientHardwareAddress
	// *SearchInfo_ClientName
	Value is_SearchInfo_SearchInfo `json:"value"`
}

func (o *SearchInfo_SearchInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SearchInfo_ClientIPAddress:
		if value != nil {
			return value.ClientIPAddress
		}
	case *SearchInfo_ClientHardwareAddress:
		if value != nil {
			return value.ClientHardwareAddress
		}
	case *SearchInfo_ClientName:
		if value != nil {
			return value.ClientName
		}
	}
	return nil
}

type is_SearchInfo_SearchInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SearchInfo_SearchInfo()
}

func (o *SearchInfo_SearchInfo) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SearchInfo_ClientIPAddress:
		return uint16(0)
	case *SearchInfo_ClientHardwareAddress:
		return uint16(1)
	case *SearchInfo_ClientName:
		return uint16(2)
	}
	return uint16(0)
}

func (o *SearchInfo_SearchInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*SearchInfo_ClientIPAddress)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SearchInfo_ClientIPAddress{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*SearchInfo_ClientHardwareAddress)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SearchInfo_ClientHardwareAddress{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*SearchInfo_ClientName)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SearchInfo_ClientName{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SearchInfo_SearchInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &SearchInfo_ClientIPAddress{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &SearchInfo_ClientHardwareAddress{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &SearchInfo_ClientName{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SearchInfo_ClientIPAddress structure represents SearchInfo_SearchInfo RPC union arm.
//
// It has following labels: 0
type SearchInfo_ClientIPAddress struct {
	// ClientIpAddress:   A pointer to a DHCP_IP_ADDRESS (section 2.2.1.2.1) structure that
	// is used to search for the DHCPv4 client lease record in the DHCPv4 server database.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
}

func (*SearchInfo_ClientIPAddress) is_SearchInfo_SearchInfo() {}

func (o *SearchInfo_ClientIPAddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.ClientIPAddress); err != nil {
		return err
	}
	return nil
}
func (o *SearchInfo_ClientIPAddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.ClientIPAddress); err != nil {
		return err
	}
	return nil
}

// SearchInfo_ClientHardwareAddress structure represents SearchInfo_SearchInfo RPC union arm.
//
// It has following labels: 1
type SearchInfo_ClientHardwareAddress struct {
	// ClientHardwareAddress:  A pointer to a DHCP_CLIENT_UID (section 2.2.1.2.5) structure
	// that represents the unique ID of a DHCPv4 client (section 2.2.1.2.5.2). It is used
	// to search for the DHCPv4 client lease record in the DHCPv4 server database.
	ClientHardwareAddress *ClientUID `idl:"name:ClientHardwareAddress" json:"client_hardware_address"`
}

func (*SearchInfo_ClientHardwareAddress) is_SearchInfo_SearchInfo() {}

func (o *SearchInfo_ClientHardwareAddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ClientHardwareAddress != nil {
		if err := o.ClientHardwareAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SearchInfo_ClientHardwareAddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ClientHardwareAddress == nil {
		o.ClientHardwareAddress = &ClientUID{}
	}
	if err := o.ClientHardwareAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SearchInfo_ClientName structure represents SearchInfo_SearchInfo RPC union arm.
//
// It has following labels: 2
type SearchInfo_ClientName struct {
	// ClientName:  A pointer to a null-terminated Unicode string, of type LPWSTR, that
	// contains the name of the DHCPv4 client. It is used to search for the DHCPv4 client
	// lease record in the DHCPv4 server database. There is no restriction on the length
	// of this Unicode string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
}

func (*SearchInfo_ClientName) is_SearchInfo_SearchInfo() {}

func (o *SearchInfo_ClientName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SearchInfo_ClientName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	return nil
}

// ClientInfoArray structure represents DHCP_CLIENT_INFO_ARRAY RPC structure.
//
// The DHCP_CLIENT_INFO_ARRAY structure defines an array of DHCP_CLIENT_INFO (section
// 2.2.1.2.12) structures.
//
// This structure is used by methods that retrieve information for more than one DHCPv4
// client.
type ClientInfoArray struct {
	// NumElements:  This is of type DWORD, containing the number of DHCPv4 clients in the
	// subsequent Clients member field. There are no inherent restrictions on the NumElements
	// member. Methods that retrieve DHCPv4 client information using the DHCP_CLIENT_INFO_ARRAY
	// structure can limit the maximum value of the NumElements member. For example, R_DhcpEnumSubnetClients
	// (section 3.1.4.21) restricts the number of elements based on input parameters and
	// the size, in addition to the number, of DHCPv4 client lease records available for
	// retrieval.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Clients:  This is a pointer of type LPDHCP_CLIENT_INFO (section 2.2.1.2.12) that
	// points to the array of length NumElements containing the DHCPv4 client's information.
	Clients []*ClientInfo `idl:"name:Clients;size_is:(NumElements)" json:"clients"`
}

func (o *ClientInfoArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Clients != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Clients))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Clients != nil || o.ElementsLength > 0 {
		_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Clients {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Clients[i1] != nil {
					_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Clients[i1] != nil {
							if err := o.Clients[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ClientInfo{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Clients[i1], _ptr_Clients); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Clients); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Clients, _ptr_Clients); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Clients", sizeInfo[0])
		}
		o.Clients = make([]*ClientInfo, sizeInfo[0])
		for i1 := range o.Clients {
			i1 := i1
			_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Clients[i1] == nil {
					o.Clients[i1] = &ClientInfo{}
				}
				if err := o.Clients[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Clients := func(ptr interface{}) { o.Clients[i1] = *ptr.(**ClientInfo) }
			if err := w.ReadPointer(&o.Clients[i1], _s_Clients, _ptr_Clients); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Clients := func(ptr interface{}) { o.Clients = *ptr.(*[]*ClientInfo) }
	if err := w.ReadPointer(&o.Clients, _s_Clients, _ptr_Clients); err != nil {
		return err
	}
	return nil
}

// OptionList structure represents DHCP_OPTION_LIST RPC structure.
//
// The DHCP_OPTION_LIST structure defines an array of option values. The first member
// contains the number of options present, and the second member contains a pointer
// to the array of length NumOptions given to a specific DHCPv6 client. This structure
// is used by the methods that retrieve options that are given to the specified DHCPv6
// client on request.
type OptionList struct {
	// NumOptions:  This is of type DWORD, containing the number of DHCPv6 options in the
	// subsequent field the Options member.
	OptionsLength uint32 `idl:"name:NumOptions" json:"options_length"`
	// Options:  This is a pointer to an array of DHCP_OPTION_VALUE (section 2.2.1.2.42)
	// structures and of length NumOptions containing DHCPv6 option values.
	Options []*OptionValue `idl:"name:Options;size_is:(NumOptions)" json:"options"`
}

func (o *OptionList) xxx_PreparePayload(ctx context.Context) error {
	if o.Options != nil && o.OptionsLength == 0 {
		o.OptionsLength = uint32(len(o.Options))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.OptionsLength); err != nil {
		return err
	}
	if o.Options != nil || o.OptionsLength > 0 {
		_ptr_Options := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.OptionsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Options {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Options[i1] != nil {
					if err := o.Options[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&OptionValue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Options); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&OptionValue{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Options, _ptr_Options); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.OptionsLength); err != nil {
		return err
	}
	_ptr_Options := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.OptionsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.OptionsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Options", sizeInfo[0])
		}
		o.Options = make([]*OptionValue, sizeInfo[0])
		for i1 := range o.Options {
			i1 := i1
			if o.Options[i1] == nil {
				o.Options[i1] = &OptionValue{}
			}
			if err := o.Options[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Options := func(ptr interface{}) { o.Options = *ptr.(*[]*OptionValue) }
	if err := w.ReadPointer(&o.Options, _s_Options, _ptr_Options); err != nil {
		return err
	}
	return nil
}

// ScopeMIBInfo structure represents SCOPE_MIB_INFO RPC structure.
//
// The SCOPE_MIB_INFO structure defines a structure that contains the address counters
// for a specific IPv4 subnet. The numbers of free, used, and offered IPv4 addresses
// are stored in this structure. This structure is used in the DHCP_MIB_INFO (section
// 2.2.1.2.48) structure.
type ScopeMIBInfo struct {
	// Subnet:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), a DWORD specifying
	// the IPv4 subnet ID for the scope.
	Subnet uint32 `idl:"name:Subnet" json:"subnet"`
	// NumAddressesInuse:  This is of type DWORD, containing the number of IPv4 addresses
	// leased out to DHCPv4 clients for a given IPv4 subnet.
	AddressesInUseLength uint32 `idl:"name:NumAddressesInuse" json:"addresses_inuse_length"`
	// NumAddressesFree:  This is of type DWORD, containing the number of IPv4 addresses
	// that are free and can be leased out to DHCPv4 clients in a specific IPv4 subnet.
	AddressesFreeLength uint32 `idl:"name:NumAddressesFree" json:"addresses_free_length"`
	// NumPendingOffers:  This is of type DWORD, containing the number of IPv4 addresses
	// that have been offered to DHCPv4 clients in a specific IPv4 subnet but that the DHCP
	// client has not yet confirmed.
	PendingOffersLength uint32 `idl:"name:NumPendingOffers" json:"pending_offers_length"`
}

func (o *ScopeMIBInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ScopeMIBInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Subnet); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressesInUseLength); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressesFreeLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PendingOffersLength); err != nil {
		return err
	}
	return nil
}
func (o *ScopeMIBInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Subnet); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressesInUseLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressesFreeLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PendingOffersLength); err != nil {
		return err
	}
	return nil
}

// MIBInfo structure represents DHCP_MIB_INFO RPC structure.
//
// The DHCP_MIB_INFO structure contains counter values for the DHCPv4 server. This structure
// is used by RPC methods like R_DhcpGetMibInfo (section 3.1.4.23) to find the DHCPv4
// server statistics.
type MIBInfo struct {
	// Discovers:  This is of type DWORD; it contains the number of DHCPDISCOVER messages
	// [RFC2131] received by the DHCPv4 server from the DHCPv4 clients since the DHCPv4
	// server was last started. This is used for statistical analysis by the DHCPv4 server.
	Discovers uint32 `idl:"name:Discovers" json:"discovers"`
	// Offers:  This is of type DWORD, containing the number of DHCPOFFER messages sent
	// by the DHCPv4 server to the DHCPv4 client that the DHCPv4 client has not confirmed
	// since the DHCP server was last started. This is used for statistical analysis by
	// the DHCPv4 server.
	Offers uint32 `idl:"name:Offers" json:"offers"`
	// Requests:  This is of type DWORD, containing the number of DHCPREQUEST messages received
	// by the DHCPv4 server from the DHCPv4 clients since the DHCP server was last started.
	// This is used for statistical analysis by the DHCPv4 server.
	Requests uint32 `idl:"name:Requests" json:"requests"`
	// Acks:  This is of type DWORD, containing the number of DHCPACK messages sent by the
	// DHCPv4 server to DHCPv4 clients since the DHCPv4 server was last started. This is
	// used for statistical analysis by the DHCPv4 server.
	ACKs uint32 `idl:"name:Acks" json:"acks"`
	// Naks:  This is of type DWORD, containing the number of DHCPNAK messages sent by the
	// DHCPv4 server to DHCPv4 clients since the DHCP server was last started. This is used
	// for statistical analysis by the DHCPv4 server.
	NAKs uint32 `idl:"name:Naks" json:"naks"`
	// Declines:  This is of type DWORD, containing the number of DHCPDECLINE messages received
	// by the DHCPv4 server from the DHCPv4 client since the DHCP server was last started.
	// This is used for statistical analysis by the DHCPv4 server.
	Declines uint32 `idl:"name:Declines" json:"declines"`
	// Releases:  This is of type DWORD, containing the number of DHCPRELEASE messages received
	// by the DHCPv4 server from the DHCPv4 client since the DHCP server was last started.
	// This is used for statistical analysis by the DHCPv4 server.
	Releases uint32 `idl:"name:Releases" json:"releases"`
	// ServerStartTime:  This is of type DATE_TIME (section 2.2.1.2.11), containing the
	// start time of the DHCPv4 server.
	ServerStartTime *DateTime `idl:"name:ServerStartTime" json:"server_start_time"`
	// Scopes:  This is of type DWORD, containing the number of IPv4 scopes configured on
	// the current DHCPv4 server. This is used for statistical analysis by the DHCPv4 server.
	// This field defines the number of DHCPv4 scopes in the subsequent field, ScopeInfo.
	Scopes uint32 `idl:"name:Scopes" json:"scopes"`
	// ScopeInfo:  This is a pointer to an array of SCOPE_MIB_INFO (section 2.2.1.2.47)
	// structures of length Scopes that contains the information about the IPv4 scopes configured
	// on the DHCPv4 server.
	ScopeInfo []*ScopeMIBInfo `idl:"name:ScopeInfo;size_is:(Scopes)" json:"scope_info"`
}

func (o *MIBInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.ScopeInfo != nil && o.Scopes == 0 {
		o.Scopes = uint32(len(o.ScopeInfo))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Discovers); err != nil {
		return err
	}
	if err := w.WriteData(o.Offers); err != nil {
		return err
	}
	if err := w.WriteData(o.Requests); err != nil {
		return err
	}
	if err := w.WriteData(o.ACKs); err != nil {
		return err
	}
	if err := w.WriteData(o.NAKs); err != nil {
		return err
	}
	if err := w.WriteData(o.Declines); err != nil {
		return err
	}
	if err := w.WriteData(o.Releases); err != nil {
		return err
	}
	if o.ServerStartTime != nil {
		if err := o.ServerStartTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Scopes); err != nil {
		return err
	}
	if o.ScopeInfo != nil || o.Scopes > 0 {
		_ptr_ScopeInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Scopes)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ScopeInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ScopeInfo[i1] != nil {
					if err := o.ScopeInfo[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ScopeMIBInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ScopeInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ScopeMIBInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeInfo, _ptr_ScopeInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Discovers); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offers); err != nil {
		return err
	}
	if err := w.ReadData(&o.Requests); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACKs); err != nil {
		return err
	}
	if err := w.ReadData(&o.NAKs); err != nil {
		return err
	}
	if err := w.ReadData(&o.Declines); err != nil {
		return err
	}
	if err := w.ReadData(&o.Releases); err != nil {
		return err
	}
	if o.ServerStartTime == nil {
		o.ServerStartTime = &DateTime{}
	}
	if err := o.ServerStartTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Scopes); err != nil {
		return err
	}
	_ptr_ScopeInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Scopes > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Scopes)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ScopeInfo", sizeInfo[0])
		}
		o.ScopeInfo = make([]*ScopeMIBInfo, sizeInfo[0])
		for i1 := range o.ScopeInfo {
			i1 := i1
			if o.ScopeInfo[i1] == nil {
				o.ScopeInfo[i1] = &ScopeMIBInfo{}
			}
			if err := o.ScopeInfo[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ScopeInfo := func(ptr interface{}) { o.ScopeInfo = *ptr.(*[]*ScopeMIBInfo) }
	if err := w.ReadPointer(&o.ScopeInfo, _s_ScopeInfo, _ptr_ScopeInfo); err != nil {
		return err
	}
	return nil
}

// OptionArray structure represents DHCP_OPTION_ARRAY RPC structure.
//
// The DHCP_OPTION_ARRAY structure contains an array of the DHCP server option definition.
// This structure is used in the DHCP_ALL_OPTIONS (section 2.2.1.2.27) structure.
type OptionArray struct {
	// NumElements:  This is of type DWORD, containing the number of option definitions
	// in the subsequent field, the Options member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Options:  This is a pointer of type DHCP_OPTION (section 2.2.1.2.25) that points
	// to an array of length NumElements containing DHCP server option definitions.
	Options []*Option `idl:"name:Options;size_is:(NumElements)" json:"options"`
}

func (o *OptionArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Options != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Options))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Options != nil || o.ElementsLength > 0 {
		_ptr_Options := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Options {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Options[i1] != nil {
					if err := o.Options[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Option{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Options); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Option{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Options, _ptr_Options); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Options := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Options", sizeInfo[0])
		}
		o.Options = make([]*Option, sizeInfo[0])
		for i1 := range o.Options {
			i1 := i1
			if o.Options[i1] == nil {
				o.Options[i1] = &Option{}
			}
			if err := o.Options[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Options := func(ptr interface{}) { o.Options = *ptr.(*[]*Option) }
	if err := w.ReadPointer(&o.Options, _s_Options, _ptr_Options); err != nil {
		return err
	}
	return nil
}

// ServerConfigInfo structure represents DHCP_SERVER_CONFIG_INFO RPC structure.
//
// The DHCP_SERVER_CONFIG_INFO structure contains settings for the DHCP server. This
// structure is used in the R_DhcpServerSetConfig (section 3.1.4.26) method.
type ServerConfigInfo struct {
	// APIProtocolSupport:  This is of type DWORD, defining the type of RPC protocol supported
	// by the DHCP server. The following type MUST be supported.
	//
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	|                                           |                                                               |
	//	|                   VALUE                   |                            MEANING                            |
	//	|                                           |                                                               |
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_TCPIP 0x00000001 | RPC protocol over TCP is used by the DHCP server to register. |
	//	+-------------------------------------------+---------------------------------------------------------------+
	//
	// The following types MAY<12> be supported.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_NP 0x00000002  | RPC protocol over named pipes is used by the DHCP server to register.            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_LPC 0x00000004 | RPC protocol over local procedure call (LPC) is used by the DHCP server to       |
	//	|                                         | register.                                                                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_ALL 0x00000007 | The DHCP server supports all of the preceding protocols.                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	APIProtocolSupport uint32 `idl:"name:APIProtocolSupport" json:"api_protocol_support"`
	// DatabaseName:  A pointer of type LPWSTR to a null-terminated Unicode string that
	// represents the DHCP server database name which is used by the DHCP server for persistent
	// storage. There is no restriction on the length of this Unicode string. This field
	// MUST be convertible to an OEM or ANSI character string.
	DatabaseName string `idl:"name:DatabaseName" json:"database_name"`
	// DatabasePath:  A pointer of type LPWSTR to a null-terminated Unicode string that
	// contains the absolute path, where the DHCP server database is stored. The maximum
	// number of characters allowed in this field is 248, including the terminating null
	// character. This field MUST be convertible to an OEM or ANSI character string.
	DatabasePath string `idl:"name:DatabasePath" json:"database_path"`
	// BackupPath:  A pointer of type LPWSTR to a null-terminated Unicode string that contains
	// the absolute path for backup storage that is used by the DHCP server for backup.
	// The maximum number of characters allowed in this field is 248, including the terminating
	// null character. This field MUST be convertible to an OEM or ANSI character string.
	BackupPath string `idl:"name:BackupPath" json:"backup_path"`
	// BackupInterval:  This is of type DWORD, containing the interval (specified in minutes)
	// between backups of the current DHCP server database.
	BackupInterval uint32 `idl:"name:BackupInterval" json:"backup_interval"`
	// DatabaseLoggingFlag:   This is of type DWORD (used as a BOOL flag), indicating the
	// transaction logging mode of the DHCP server. The value 1 indicates that transaction
	// logging mode is enabled for the DHCP server, and zero indicates that transaction
	// logging mode is disabled for the DHCP server.
	DatabaseLoggingFlag uint32 `idl:"name:DatabaseLoggingFlag" json:"database_logging_flag"`
	// RestoreFlag:  This is of type DWORD (used as a BOOL flag), and if this setting is
	// TRUE, the DHCP server loads the DHCP server database from the backup database on
	// DHCP server startup. The default value of this flag is FALSE.
	RestoreFlag uint32 `idl:"name:RestoreFlag" json:"restore_flag"`
	// DatabaseCleanupInterval:  This is of type DWORD and specifies the maximum time interval
	// in minutes that DOOMED IPv4 DHCP client lease records can persist before being deleted
	// from the DHCP server database.
	DatabaseCleanupInterval uint32 `idl:"name:DatabaseCleanupInterval" json:"database_cleanup_interval"`
	// DebugFlag:  A flag that specifies the level of logging done by the DHCP server. The
	// following table defines the set values that can be used. Specifying 0xFFFFFFFF enables
	// all types of logging.
	//
	// LOW WORD bitmask (0x0000FFFF) for low-frequency debug output.
	//
	//	+-----------------------------+----------------------------------------------------------------------+
	//	|                             |                                                                      |
	//	|            VALUE            |                               MEANING                                |
	//	|                             |                                                                      |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_ADDRESS 0x00000001    | Enable IP address-related logging.                                   |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_CLIENT 0x00000002     | Enable DHCP-client-API-related logging.                              |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_PARAMETERS 0x00000004 | Enable DHCP-server-parameters-related logging.                       |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_OPTIONS 0x00000008    | Enable DHCP-options-related logging.                                 |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_ERRORS 0x00000010     | Enable DHCP-errors-related logging.                                  |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_STOC 0x00000020       | Enable DHCPv4 and DCHPv6-protocol-errors-related logging.            |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_INIT 0x00000040       | Enable DHCP-server-initialization-related logging.                   |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_SCAVENGER 0x00000080  | Enable scavenger's-error-related logging.                            |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_TIMESTAMP 0x00000100  | Enable timing-errors-related logging.                                |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_APIS 0x00000200       | Enable DHCP-APIs-related logging.                                    |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_REGISTRY 0x00000400   | Enable the logging of errors caused by registry setting operations.  |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_JET 0x00000800        | Enable the logging of the DHCP server database errors.               |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_THREADPOOL 0x00001000 | Enable the logging related to executing thread pool operations.      |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_AUDITLOG 0x00002000   | Enable the logging related to errors caused by audit log operations. |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_QUARANTINE 0x00004000 | Enable the logging of errors caused by quarantine errors.            |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_MISC 0x00008000       | Enable the logging caused by miscellaneous errors.                   |
	//	+-----------------------------+----------------------------------------------------------------------+
	//
	// HIGH WORD bitmask (0xFFFF0000) for high-frequency debug output, that is, more verbose.
	//
	//	+------------------------------+-------------------------------------------------------------------------+
	//	|                              |                                                                         |
	//	|            VALUE             |                                 MEANING                                 |
	//	|                              |                                                                         |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_MESSAGE 0x00010000     | Enable the logging related to debug messages.                           |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_API_VERBOSE 0x00020000 | Enable the logging related to DHCP API verbose errors.                  |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_DNS 0x00040000         | Enable the logging related to Domain Name System (DNS) messages.        |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_MSTOC 0x00080000       | Enable the logging related to multicast protocol layer errors.          |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_TRACK 0x00100000       | Enable the logging tracking specific problems.                          |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_ROGUE 0x00200000       | Enable the logging related to a rogue DHCP server.                      |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_PNP 0x00400000         | Enable the logging related to PNP interface errors.                     |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_PERF 0x01000000        | Enable the logging of performance-related messages.                     |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_ALLOC 0x02000000       | Enable the logging of messages related to allocation and de-allocation. |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_PING 0x04000000        | Enable the logging of synchronous-ping-related messages.                |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_THREAD 0x08000000      | Enable the logging of thread-related messages.                          |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_TRACE 0x10000000       | Enable the logging for tracing through code messages.                   |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_TRACE_CALLS 0x20000000 | Enable the logging for tracing through piles of code.                   |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_STARTUP_BRK 0x40000000 | Enable the logging related to debugger break during setup messages.     |
	//	+------------------------------+-------------------------------------------------------------------------+
	//	| DEBUG_LOG_IN_FILE 0x80000000 | Enable the logging of debug output in a file.                           |
	//	+------------------------------+-------------------------------------------------------------------------+
	DebugFlag uint32 `idl:"name:DebugFlag" json:"debug_flag"`
}

func (o *ServerConfigInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerConfigInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.APIProtocolSupport); err != nil {
		return err
	}
	if o.DatabaseName != "" {
		_ptr_DatabaseName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DatabaseName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DatabaseName, _ptr_DatabaseName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DatabasePath != "" {
		_ptr_DatabasePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DatabasePath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DatabasePath, _ptr_DatabasePath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.BackupPath != "" {
		_ptr_BackupPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.BackupPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.BackupPath, _ptr_BackupPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BackupInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DatabaseLoggingFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.RestoreFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.DatabaseCleanupInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DebugFlag); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ServerConfigInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.APIProtocolSupport); err != nil {
		return err
	}
	_ptr_DatabaseName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DatabaseName); err != nil {
			return err
		}
		return nil
	})
	_s_DatabaseName := func(ptr interface{}) { o.DatabaseName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DatabaseName, _s_DatabaseName, _ptr_DatabaseName); err != nil {
		return err
	}
	_ptr_DatabasePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DatabasePath); err != nil {
			return err
		}
		return nil
	})
	_s_DatabasePath := func(ptr interface{}) { o.DatabasePath = *ptr.(*string) }
	if err := w.ReadPointer(&o.DatabasePath, _s_DatabasePath, _ptr_DatabasePath); err != nil {
		return err
	}
	_ptr_BackupPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.BackupPath); err != nil {
			return err
		}
		return nil
	})
	_s_BackupPath := func(ptr interface{}) { o.BackupPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.BackupPath, _s_BackupPath, _ptr_BackupPath); err != nil {
		return err
	}
	if err := w.ReadData(&o.BackupInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DatabaseLoggingFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.RestoreFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.DatabaseCleanupInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DebugFlag); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ScanItem structure represents DHCP_SCAN_ITEM RPC structure.
//
// The DHCP_SCAN_ITEM structure defines the type of fix that is required for DHCPv4
// client lease records that are missing in the bitmask representation in memory (section
// 3.1.1.4) or vice versa.
type ScanItem struct {
	// IpAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), a DWORD containing
	// the IPv4 address of the missing DHCPv4 client entry in one of the stores.
	IPAddress uint32 `idl:"name:IpAddress" json:"ip_address"`
	// ScanFlag:  This is of type DHCP_SCAN_FLAG (section 2.2.1.1.13) enumeration, which
	// contains an enumerated value of 0 to fix the bitmask representation (section 3.1.1.4)
	// and an enumerated value of 1 to fix the DHCPv4 client Lease records.
	ScanFlag ScanFlag `idl:"name:ScanFlag" json:"scan_flag"`
}

func (o *ScanItem) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ScanItem) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.IPAddress); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ScanFlag)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *ScanItem) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPAddress); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ScanFlag)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// ScanList structure represents DHCP_SCAN_LIST RPC structure.
//
// The DHCP_SCAN_LIST structure defines an array of DHCP_SCAN_ITEM (section 2.2.1.2.73)
// structures that contains all the entries on the DHCP server that require a fix. This
// structure is used in the R_DhcpScanDatabase (section 3.1.4.28) method.
type ScanList struct {
	// NumScanItems:  This is of type DWORD, containing the number of DHCPv4 client lease
	// entries and/or bits in the bitmask representation in memory (section 3.1.1.4) that
	// require a fix in the subsequent field in the ScanItems member.
	ScanItemsLength uint32 `idl:"name:NumScanItems" json:"scan_items_length"`
	// ScanItems:  This is a pointer to an array of DHCP_SCAN_ITEM (section 2.2.1.2.73)
	// structures of length NumScanItems that contains the DHCPv4 client IPv4 addresses
	// that require a fix.
	ScanItems []*ScanItem `idl:"name:ScanItems;size_is:(NumScanItems)" json:"scan_items"`
}

func (o *ScanList) xxx_PreparePayload(ctx context.Context) error {
	if o.ScanItems != nil && o.ScanItemsLength == 0 {
		o.ScanItemsLength = uint32(len(o.ScanItems))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ScanList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ScanItemsLength); err != nil {
		return err
	}
	if o.ScanItems != nil || o.ScanItemsLength > 0 {
		_ptr_ScanItems := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ScanItemsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ScanItems {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ScanItems[i1] != nil {
					if err := o.ScanItems[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ScanItem{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ScanItems); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ScanItem{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScanItems, _ptr_ScanItems); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ScanList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScanItemsLength); err != nil {
		return err
	}
	_ptr_ScanItems := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ScanItemsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ScanItemsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ScanItems", sizeInfo[0])
		}
		o.ScanItems = make([]*ScanItem, sizeInfo[0])
		for i1 := range o.ScanItems {
			i1 := i1
			if o.ScanItems[i1] == nil {
				o.ScanItems[i1] = &ScanItem{}
			}
			if err := o.ScanItems[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ScanItems := func(ptr interface{}) { o.ScanItems = *ptr.(*[]*ScanItem) }
	if err := w.ReadPointer(&o.ScanItems, _s_ScanItems, _ptr_ScanItems); err != nil {
		return err
	}
	return nil
}

// IPReservationV4 structure represents DHCP_IP_RESERVATION_V4 RPC structure.
//
// The DHCP_IP_RESERVATION_V4 structure defines an IPv4 reservation for a DHCP client.
// This structure is an extension of DHCP_IP_RESERVATION (section 2.2.1.2.10) structure
// by including the type of client (DHCP, BOOTP or both) ([RFC2132]) holding this IPv4
// reservation. This structure is used in the DHCP_SUBNET_ELEMENT_DATA_V4 (section 2.2.1.2.35)
// structure.
type IPReservationV4 struct {
	// ReservedIpAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1) that contains
	// the IPv4 address of client (DHCP or BOOTP) for which a reservation was created.
	ReservedIPAddress uint32 `idl:"name:ReservedIpAddress" json:"reserved_ip_address"`
	// ReservedForClient:  This is a pointer of type DHCP_CLIENT_UID (section 2.2.1.2.5)
	// that represents the DHCPv4 client-identifier (section 2.2.1.2.5.1).
	ReservedForClient *ClientUID `idl:"name:ReservedForClient" json:"reserved_for_client"`
	// bAllowedClientTypes:  This is of type BYTE that specifies the type of client holding
	// this reservation.
	//
	//	+------------------------+----------------------------------------------------+
	//	|                        |                                                    |
	//	|         VALUE          |                      MEANING                       |
	//	|                        |                                                    |
	//	+------------------------+----------------------------------------------------+
	//	+------------------------+----------------------------------------------------+
	//	| CLIENT_TYPE_DHCP 0x01  | The IPv4 reservation is for a DHCPv4 client.       |
	//	+------------------------+----------------------------------------------------+
	//	| CLIENT_TYPE_BOOTP 0x02 | The IPv4 reservation is for a BOOTP client.        |
	//	+------------------------+----------------------------------------------------+
	//	| CLIENT_TYPE_BOTH 0x03  | The IPv4 reservation is for both kinds of clients. |
	//	+------------------------+----------------------------------------------------+
	AllowedClientTypes uint8 `idl:"name:bAllowedClientTypes" json:"allowed_client_types"`
}

func (o *IPReservationV4) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPReservationV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReservedIPAddress); err != nil {
		return err
	}
	if o.ReservedForClient != nil {
		_ptr_ReservedForClient := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ReservedForClient != nil {
				if err := o.ReservedForClient.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReservedForClient, _ptr_ReservedForClient); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AllowedClientTypes); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *IPReservationV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReservedIPAddress); err != nil {
		return err
	}
	_ptr_ReservedForClient := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ReservedForClient == nil {
			o.ReservedForClient = &ClientUID{}
		}
		if err := o.ReservedForClient.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ReservedForClient := func(ptr interface{}) { o.ReservedForClient = *ptr.(**ClientUID) }
	if err := w.ReadPointer(&o.ReservedForClient, _s_ReservedForClient, _ptr_ReservedForClient); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowedClientTypes); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV4 structure represents DHCP_SUBNET_ELEMENT_DATA_V4 RPC structure.
//
// The DHCP_SUBNET_ELEMENT_DATA_V4 structure defines the IPv4 reservation, IPv4 exclusion
// range, or IPv4 range elements for the subnet. This structure is an extension of the
// DHCP_SUBNET_ELEMENT_DATA (section 2.2.1.2.33) structure and is used in the method
// R_DhcpAddSubnetElementV4 (section 3.1.4.30).
type SubnetElementDataV4 struct {
	// ElementType:  This is of type DHCP_SUBNET_ELEMENT_TYPE (section 2.2.1.1.7) enumeration,
	// defining the set of possible IPv4 subnet element types. This value defines which
	// of the values is chosen from the subsequent union, the Element member.
	ElementType SubnetElementType `idl:"name:ElementType" json:"element_type"`
	// Element:  Element is a union of different types of IPv4 subnet elements. The value
	// of the union is dependent on the previous field, ElementType.
	Element *SubnetElementDataV4_Element `idl:"name:Element;switch_is:(((ElementType 7 <=) (5 ElementType <=) &&) 0 ElementType ?:)" json:"element"`
}

func (o *SubnetElementDataV4) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ElementType)); err != nil {
		return err
	}
	_exprElementType := uint16(0)
	if (o.ElementType <= 7) && (5 <= o.ElementType) {
		_exprElementType = uint16(0)
	} else {
		_exprElementType = uint16(o.ElementType)
	}
	_swElement := uint16(_exprElementType)
	if o.Element != nil {
		if err := o.Element.MarshalUnionNDR(ctx, w, _swElement); err != nil {
			return err
		}
	} else {
		if err := (&SubnetElementDataV4_Element{}).MarshalUnionNDR(ctx, w, _swElement); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ElementType)); err != nil {
		return err
	}
	if o.Element == nil {
		o.Element = &SubnetElementDataV4_Element{}
	}
	_exprElementType := uint16(0)
	if (o.ElementType <= 7) && (5 <= o.ElementType) {
		_exprElementType = uint16(0)
	} else {
		_exprElementType = uint16(o.ElementType)
	}
	_swElement := uint16(_exprElementType)
	if err := o.Element.UnmarshalUnionNDR(ctx, w, _swElement); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV4_Element structure represents DHCP_SUBNET_ELEMENT_DATA_V4 union anonymous member.
//
// The DHCP_SUBNET_ELEMENT_DATA_V4 structure defines the IPv4 reservation, IPv4 exclusion
// range, or IPv4 range elements for the subnet. This structure is an extension of the
// DHCP_SUBNET_ELEMENT_DATA (section 2.2.1.2.33) structure and is used in the method
// R_DhcpAddSubnetElementV4 (section 3.1.4.30).
type SubnetElementDataV4_Element struct {
	// Types that are assignable to Value
	//
	// *SubnetElementDataV4_IPRange
	// *SubnetElementDataV4_SecondaryHost
	// *SubnetElementDataV4_ReservedIP
	// *SubnetElementDataV4_ExcludeIPRange
	// *SubnetElementDataV4_IPUsedCluster
	Value is_SubnetElementDataV4_Element `json:"value"`
}

func (o *SubnetElementDataV4_Element) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SubnetElementDataV4_IPRange:
		if value != nil {
			return value.IPRange
		}
	case *SubnetElementDataV4_SecondaryHost:
		if value != nil {
			return value.SecondaryHost
		}
	case *SubnetElementDataV4_ReservedIP:
		if value != nil {
			return value.ReservedIP
		}
	case *SubnetElementDataV4_ExcludeIPRange:
		if value != nil {
			return value.ExcludeIPRange
		}
	case *SubnetElementDataV4_IPUsedCluster:
		if value != nil {
			return value.IPUsedCluster
		}
	}
	return nil
}

type is_SubnetElementDataV4_Element interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SubnetElementDataV4_Element()
}

func (o *SubnetElementDataV4_Element) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SubnetElementDataV4_IPRange:
		return uint16(0)
	case *SubnetElementDataV4_SecondaryHost:
		return uint16(1)
	case *SubnetElementDataV4_ReservedIP:
		return uint16(2)
	case *SubnetElementDataV4_ExcludeIPRange:
		return uint16(3)
	case *SubnetElementDataV4_IPUsedCluster:
		return uint16(4)
	}
	return uint16(0)
}

func (o *SubnetElementDataV4_Element) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*SubnetElementDataV4_IPRange)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV4_IPRange{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*SubnetElementDataV4_SecondaryHost)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV4_SecondaryHost{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*SubnetElementDataV4_ReservedIP)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV4_ReservedIP{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*SubnetElementDataV4_ExcludeIPRange)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV4_ExcludeIPRange{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*SubnetElementDataV4_IPUsedCluster)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV4_IPUsedCluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SubnetElementDataV4_Element) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &SubnetElementDataV4_IPRange{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &SubnetElementDataV4_SecondaryHost{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &SubnetElementDataV4_ReservedIP{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &SubnetElementDataV4_ExcludeIPRange{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &SubnetElementDataV4_IPUsedCluster{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SubnetElementDataV4_IPRange structure represents SubnetElementDataV4_Element RPC union arm.
//
// It has following labels: 0
type SubnetElementDataV4_IPRange struct {
	// IpRange:  This is of type DHCP_IP_RANGE (section 2.2.1.2.31) structure, containing
	// the IPv4 range for the IPv4 subnet. This contains the range for the following valid
	// enumeration values.
	//
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	|                          |                                                                                  |
	//	| DHCP SUBNET ELEMENT TYPE |                                     MEANING                                      |
	//	|                          |                                                                                  |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRanges 0           | The configuration parameter is the IP range of a DHCPv4 scope configured on the  |
	//	|                          | DHCPv4 server.                                                                   |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRangesDhcpOnly 5   | The configuration parameter is an IP range of a DHCPv4 scope configured on       |
	//	|                          | the DHCPv4 server that MUST be used only for assignment of addresses to DHCPv4   |
	//	|                          | clients on the subnet. The IP addresses from this range MUST NOT be assigned to  |
	//	|                          | BOOTP clients ([RFC2132]).                                                       |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRangesDhcpBootp 6  | The configuration parameter is an IP range of a DHCPv4 scope configured on the   |
	//	|                          | DHCPv4 server that can be used for assignment of addresses to both DHCPv4 and    |
	//	|                          | BOOTP.                                                                           |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRangesBootpOnly 7  | The configuration parameter is an IP range of a DHCPv4 scope configured on the   |
	//	|                          | DHCPv4 server that MUST be used only for assignment of IPv4 addresses to BOOTP   |
	//	|                          | clients.                                                                         |
	//	+--------------------------+----------------------------------------------------------------------------------+
	IPRange *IPRange `idl:"name:IpRange" json:"ip_range"`
}

func (*SubnetElementDataV4_IPRange) is_SubnetElementDataV4_Element() {}

func (o *SubnetElementDataV4_IPRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IPRange != nil {
		_ptr_IpRange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPRange != nil {
				if err := o.IPRange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPRange{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPRange, _ptr_IpRange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV4_IPRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_IpRange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPRange == nil {
			o.IPRange = &IPRange{}
		}
		if err := o.IPRange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IpRange := func(ptr interface{}) { o.IPRange = *ptr.(**IPRange) }
	if err := w.ReadPointer(&o.IPRange, _s_IpRange, _ptr_IpRange); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV4_SecondaryHost structure represents SubnetElementDataV4_Element RPC union arm.
//
// It has following labels: 1
type SubnetElementDataV4_SecondaryHost struct {
	// SecondaryHost:  This is of type DHCP_HOST_INFO (section 2.2.1.2.7) structure and
	// is not used. If the ElementType value mandates that the SecondaryHost element is
	// to be used in any method, the method will return ERROR_CALL_NOT_IMPLEMENTED or ERROR_NOT_SUPPORTED,
	// as specified in the processing rules of methods that use the DHCP_SUBNET_ELEMENT_DATA_V4
	// structure.
	SecondaryHost *HostInfo `idl:"name:SecondaryHost" json:"secondary_host"`
}

func (*SubnetElementDataV4_SecondaryHost) is_SubnetElementDataV4_Element() {}

func (o *SubnetElementDataV4_SecondaryHost) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SecondaryHost != nil {
		_ptr_SecondaryHost := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SecondaryHost != nil {
				if err := o.SecondaryHost.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecondaryHost, _ptr_SecondaryHost); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV4_SecondaryHost) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SecondaryHost := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SecondaryHost == nil {
			o.SecondaryHost = &HostInfo{}
		}
		if err := o.SecondaryHost.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SecondaryHost := func(ptr interface{}) { o.SecondaryHost = *ptr.(**HostInfo) }
	if err := w.ReadPointer(&o.SecondaryHost, _s_SecondaryHost, _ptr_SecondaryHost); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV4_ReservedIP structure represents SubnetElementDataV4_Element RPC union arm.
//
// It has following labels: 2
type SubnetElementDataV4_ReservedIP struct {
	// ReservedIp:  This is of type DHCP_IP_RESERVATION_V4 (section 2.2.1.2.32) structure,
	// containing the IPv4 reservation.
	ReservedIP *IPReservationV4 `idl:"name:ReservedIp" json:"reserved_ip"`
}

func (*SubnetElementDataV4_ReservedIP) is_SubnetElementDataV4_Element() {}

func (o *SubnetElementDataV4_ReservedIP) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ReservedIP != nil {
		_ptr_ReservedIp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ReservedIP != nil {
				if err := o.ReservedIP.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPReservationV4{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReservedIP, _ptr_ReservedIp); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV4_ReservedIP) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ReservedIp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ReservedIP == nil {
			o.ReservedIP = &IPReservationV4{}
		}
		if err := o.ReservedIP.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ReservedIp := func(ptr interface{}) { o.ReservedIP = *ptr.(**IPReservationV4) }
	if err := w.ReadPointer(&o.ReservedIP, _s_ReservedIp, _ptr_ReservedIp); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV4_ExcludeIPRange structure represents SubnetElementDataV4_Element RPC union arm.
//
// It has following labels: 3
type SubnetElementDataV4_ExcludeIPRange struct {
	// ExcludeIpRange:  This is of type DHCP_IP_RANGE structure, containing the IPv4 exclusion
	// range.
	ExcludeIPRange *IPRange `idl:"name:ExcludeIpRange" json:"exclude_ip_range"`
}

func (*SubnetElementDataV4_ExcludeIPRange) is_SubnetElementDataV4_Element() {}

func (o *SubnetElementDataV4_ExcludeIPRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ExcludeIPRange != nil {
		_ptr_ExcludeIpRange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ExcludeIPRange != nil {
				if err := o.ExcludeIPRange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPRange{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ExcludeIPRange, _ptr_ExcludeIpRange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV4_ExcludeIPRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ExcludeIpRange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ExcludeIPRange == nil {
			o.ExcludeIPRange = &IPRange{}
		}
		if err := o.ExcludeIPRange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ExcludeIpRange := func(ptr interface{}) { o.ExcludeIPRange = *ptr.(**IPRange) }
	if err := w.ReadPointer(&o.ExcludeIPRange, _s_ExcludeIpRange, _ptr_ExcludeIpRange); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV4_IPUsedCluster structure represents SubnetElementDataV4_Element RPC union arm.
//
// It has following labels: 4
type SubnetElementDataV4_IPUsedCluster struct {
	// IpUsedCluster:  This is of type DHCP_IP_CLUSTER (section 2.2.1.2.88) structure and
	// is not used. If the ElementType member mandates this element to be used in any method,
	// the method will return ERROR_INVALID_PARAMETER.
	IPUsedCluster *IPCluster `idl:"name:IpUsedCluster" json:"ip_used_cluster"`
}

func (*SubnetElementDataV4_IPUsedCluster) is_SubnetElementDataV4_Element() {}

func (o *SubnetElementDataV4_IPUsedCluster) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IPUsedCluster != nil {
		_ptr_IpUsedCluster := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPUsedCluster != nil {
				if err := o.IPUsedCluster.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPCluster{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPUsedCluster, _ptr_IpUsedCluster); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV4_IPUsedCluster) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_IpUsedCluster := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPUsedCluster == nil {
			o.IPUsedCluster = &IPCluster{}
		}
		if err := o.IPUsedCluster.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IpUsedCluster := func(ptr interface{}) { o.IPUsedCluster = *ptr.(**IPCluster) }
	if err := w.ReadPointer(&o.IPUsedCluster, _s_IpUsedCluster, _ptr_IpUsedCluster); err != nil {
		return err
	}
	return nil
}

// SubnetElementInfoArrayV4 structure represents DHCP_SUBNET_ELEMENT_INFO_ARRAY_V4 RPC structure.
//
// The DHCP_SUBNET_ELEMENT_INFO_ARRAY_V4 structure defines an array of DHCP_SUBNET_ELEMENT_DATA_V4
// (section 2.2.1.2.35) structures.
//
// This structure is used in the R_DhcpEnumSubnetElementsV4 (section 3.1.4.31) method.
// The first member contains the number of subnet elements (IPv4 reservation, IPv4 exclusion
// range, or IPv4 range), and the second member points to the array of length NumElements
// containing DHCPv4 subnet elements.
type SubnetElementInfoArrayV4 struct {
	// NumElements:  This is of type DWORD, containing the number of subnet elements in
	// the subsequent field, the Elements member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This is a pointer to an array of DHCP_SUBNET_ELEMENT_DATA_V4 structures
	// of length NumElements, containing subnet elements.
	Elements []*SubnetElementDataV4 `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *SubnetElementInfoArrayV4) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementInfoArrayV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SubnetElementDataV4{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SubnetElementDataV4{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementInfoArrayV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*SubnetElementDataV4, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &SubnetElementDataV4{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*SubnetElementDataV4) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// ClientInfoV4 structure represents DHCP_CLIENT_INFO_V4 RPC structure.
//
// The DHCP_CLIENT_INFO_V4 structure defines information about the DHCPv4 client that
// is used by the R_DhcpGetClientInfoV4 (section 3.1.4.35) method.
//
// This structure augments the DHCP_CLIENT_INFO (section 2.2.1.2.12) structure by including
// the additional element bClientType.
type ClientInfoV4 struct {
	// ClientIpAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), a DWORD that
	// contains the DHCPv4 client's IPv4 address.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// SubnetMask:  This is of type DHCP_IP_MASK (section 2.2.1.2.2), a DWORD that contains
	// the DHCPv4 client's IPv4 subnet mask address.
	SubnetMask uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
	// ClientHardwareAddress:  This is of type DHCP_CLIENT_UID (section 2.2.1.2.5), a structure
	// that represents a DHCPv4 client-identifier (section 2.2.1.2.5.1) or a DHCPv4 client
	// unique ID (section 2.2.1.2.5.2). Methods that accept DHCP_CLIENT_INFO_V4 as a parameter
	// specify which representations are acceptable.
	ClientHardwareAddress *ClientUID `idl:"name:ClientHardwareAddress" json:"client_hardware_address"`
	// ClientName:  A pointer to a null-terminated Unicode string that represents the DHCPv4
	// client's internet host name. There is no restriction on the length of this Unicode
	// string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
	// ClientComment:  A pointer to a null-terminated Unicode string that represents a description
	// of the DHCPv4 client. There is no restriction on the length of this Unicode string.
	ClientComment string `idl:"name:ClientComment" json:"client_comment"`
	// ClientLeaseExpires:  This is of type DATE_TIME (section 2.2.1.2.11), a structure
	// that contains the lease expiry time for the DHCPv4 client. This is UTC time.
	ClientLeaseExpires *DateTime `idl:"name:ClientLeaseExpires" json:"client_lease_expires"`
	// OwnerHost:  This is of type DHCP_HOST_INFO (section 2.2.1.2.7), a structure that
	// contains information on the DHCPv4 server machine that has provided a lease to the
	// DHCPv4 client.
	OwnerHost *HostInfo `idl:"name:OwnerHost" json:"owner_host"`
	// bClientType:  This is of type BYTE that identifies the type of the DHCPv4 client.
	// Possible values for this field are provided in the following table.
	//
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	|                                   |                                                                       |
	//	|               VALUE               |                                MEANING                                |
	//	|                                   |                                                                       |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_UNSPECIFIED 0x00      | A DHCPv4 client other than ones defined in this table.                |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_DHCP 0x01             | The DHCPv4 client supports the DHCPv4 protocol.                       |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOOTP 0x02            | The DHCPv4 client supports the BOOTP protocol ([RFC2132]).            |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOTH 0x03             | The DHCPv4 client identifies both the DHCPv4 and the BOOTP protocols. |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_RESERVATION_FLAG 0x04 | There is an IPv4 reservation created for the DHCPv4 client.           |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_NONE 0x64             | Backward compatibility for manual addressing.                         |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	ClientType uint8 `idl:"name:bClientType" json:"client_type"`
}

func (o *ClientInfoV4) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress != nil {
		if err := o.ClientHardwareAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientComment != "" {
		_ptr_ClientComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientComment, _ptr_ClientComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientLeaseExpires != nil {
		if err := o.ClientLeaseExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OwnerHost != nil {
		if err := o.OwnerHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ClientType); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ClientInfoV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress == nil {
		o.ClientHardwareAddress = &ClientUID{}
	}
	if err := o.ClientHardwareAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	_ptr_ClientComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClientComment := func(ptr interface{}) { o.ClientComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientComment, _s_ClientComment, _ptr_ClientComment); err != nil {
		return err
	}
	if o.ClientLeaseExpires == nil {
		o.ClientLeaseExpires = &DateTime{}
	}
	if err := o.ClientLeaseExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OwnerHost == nil {
		o.OwnerHost = &HostInfo{}
	}
	if err := o.OwnerHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientType); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ClientInfoArrayV4 structure represents DHCP_CLIENT_INFO_ARRAY_V4 RPC structure.
//
// The DHCP_CLIENT_INFO_ARRAY_V4 structure defines an array of DHCP_CLIENT_INFO_V4 (section
// 2.2.1.2.14) structures.
//
// This structure is used by methods, such as R_DhcpEnumSubnetClientsV4 (section 3.1.4.36),
// that retrieve information for more than one DHCP client.
type ClientInfoArrayV4 struct {
	// NumElements:   This is of type DWORD, containing the number of DHCPv4 client-specific
	// subnets, which is also the number of DHCPv4 clients in the Clients member. There
	// are no inherent restrictions on the NumElements member. Methods that retrieve DHCPv4
	// client information using the DHCP_CLIENT_INFO_ARRAY_V4 structure can limit the maximum
	// value of the NumElements member. For example, R_DhcpEnumSubnetClientsV4 restricts
	// the number of elements based on input parameters and the size, as well as number,
	// of DHCPv4 client lease records available for retrieval.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Clients:  This is a pointer of type LPDHCP_CLIENT_INFO_V4 that points to the array
	// of length NumElements containing the DHCPv4 client information.
	Clients []*ClientInfoV4 `idl:"name:Clients;size_is:(NumElements)" json:"clients"`
}

func (o *ClientInfoArrayV4) xxx_PreparePayload(ctx context.Context) error {
	if o.Clients != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Clients))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoArrayV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Clients != nil || o.ElementsLength > 0 {
		_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Clients {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Clients[i1] != nil {
					_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Clients[i1] != nil {
							if err := o.Clients[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ClientInfoV4{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Clients[i1], _ptr_Clients); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Clients); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Clients, _ptr_Clients); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoArrayV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Clients", sizeInfo[0])
		}
		o.Clients = make([]*ClientInfoV4, sizeInfo[0])
		for i1 := range o.Clients {
			i1 := i1
			_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Clients[i1] == nil {
					o.Clients[i1] = &ClientInfoV4{}
				}
				if err := o.Clients[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Clients := func(ptr interface{}) { o.Clients[i1] = *ptr.(**ClientInfoV4) }
			if err := w.ReadPointer(&o.Clients[i1], _s_Clients, _ptr_Clients); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Clients := func(ptr interface{}) { o.Clients = *ptr.(*[]*ClientInfoV4) }
	if err := w.ReadPointer(&o.Clients, _s_Clients, _ptr_Clients); err != nil {
		return err
	}
	return nil
}

// SuperScopeTableEntry structure represents DHCP_SUPER_SCOPE_TABLE_ENTRY RPC structure.
//
// The DHCP_SUPER_SCOPE_TABLE_ENTRY structure defines the superscope-specific subnet
// information. This structure is used in the DHCP_SUPER_SCOPE_TABLE (section 2.2.1.2.86)
// structure.
type SuperScopeTableEntry struct {
	// SubnetAddress:  This is of type DHCP_IP_ADDRESS, a DWORD containing the IPv4 subnet
	// ID.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// SuperScopeNumber:  This is of type DWORD, containing the unique identifier of the
	// superscope.
	SuperScopeNumber uint32 `idl:"name:SuperScopeNumber" json:"super_scope_number"`
	// NextInSuperScope:  This is of type DWORD, specifying the index of the next subnet
	// ID in the superscope.
	NextInSuperScope uint32 `idl:"name:NextInSuperScope" json:"next_in_super_scope"`
	// SuperScopeName:  This is a pointer, of type LPWSTR, to a null-terminated Unicode
	// string that contains the superscope name. There is no restriction on the length of
	// this Unicode string.
	SuperScopeName string `idl:"name:SuperScopeName" json:"super_scope_name"`
}

func (o *SuperScopeTableEntry) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SuperScopeTableEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SuperScopeNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.NextInSuperScope); err != nil {
		return err
	}
	if o.SuperScopeName != "" {
		_ptr_SuperScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SuperScopeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SuperScopeName, _ptr_SuperScopeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SuperScopeTableEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SuperScopeNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.NextInSuperScope); err != nil {
		return err
	}
	_ptr_SuperScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SuperScopeName); err != nil {
			return err
		}
		return nil
	})
	_s_SuperScopeName := func(ptr interface{}) { o.SuperScopeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.SuperScopeName, _s_SuperScopeName, _ptr_SuperScopeName); err != nil {
		return err
	}
	return nil
}

// SuperScopeTable structure represents DHCP_SUPER_SCOPE_TABLE RPC structure.
//
// The DHCP_SUPER_SCOPE_TABLE structure defines an array of DHCP_SUPER_SCOPE_TABLE_ENTRY
// (section 2.2.1.2.85) structures. This contains information about more than one subnet
// within a superscope. The first member contains the number of IPv4 subnets present,
// and the second member points to the array of length cEntries containing all subnet
// information. This structure is used in the R_DhcpGetSuperScopeInfoV4 (section 3.1.4.38)
// method.
type SuperScopeTable struct {
	// cEntries:  This is of type DWORD, containing the number of superscope entries in
	// the subsequent field the pEntries member.
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	// pEntries:  This is a pointer of type DHCP_SUPER_SCOPE_TABLE_ENTRY (section 2.2.1.2.85)
	// structure that points to an array of length cEntries containing superscope-specific
	// subnet information.
	Entries []*SuperScopeTableEntry `idl:"name:pEntries;size_is:(cEntries)" json:"entries"`
}

func (o *SuperScopeTable) xxx_PreparePayload(ctx context.Context) error {
	if o.Entries != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.Entries))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SuperScopeTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesCount); err != nil {
		return err
	}
	if o.Entries != nil || o.EntriesCount > 0 {
		_ptr_pEntries := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesCount)
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
					if err := (&SuperScopeTableEntry{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Entries); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SuperScopeTableEntry{}).MarshalNDR(ctx, w); err != nil {
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
func (o *SuperScopeTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesCount); err != nil {
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
		if o.EntriesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Entries", sizeInfo[0])
		}
		o.Entries = make([]*SuperScopeTableEntry, sizeInfo[0])
		for i1 := range o.Entries {
			i1 := i1
			if o.Entries[i1] == nil {
				o.Entries[i1] = &SuperScopeTableEntry{}
			}
			if err := o.Entries[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pEntries := func(ptr interface{}) { o.Entries = *ptr.(*[]*SuperScopeTableEntry) }
	if err := w.ReadPointer(&o.Entries, _s_pEntries, _ptr_pEntries); err != nil {
		return err
	}
	return nil
}

// ServerConfigInfoV4 structure represents DHCP_SERVER_CONFIG_INFO_V4 RPC structure.
//
// The DHCP_SERVER_CONFIG_INFO_V4 structure defines DHCP server settings. This structure
// is an extension of DHCP_SERVER_CONFIG_INFO (section 2.2.1.2.53) structure and used
// in the R_DhcpServerSetConfigV4 (section 3.1.4.40) method.
type ServerConfigInfoV4 struct {
	// APIProtocolSupport:  This is of type DWORD, defining the type of RPC protocol used
	// by the DHCP server to register with RPC. Following is the set of supported types,
	// which can be bitwise OR'd to produce valid values. The following type MUST be supported.
	//
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	|                                           |                                                               |
	//	|                   VALUE                   |                            MEANING                            |
	//	|                                           |                                                               |
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_TCPIP 0x00000001 | RPC protocol over TCP is used by the DHCP server to register. |
	//	+-------------------------------------------+---------------------------------------------------------------+
	//
	// The following types MAY<13> be supported.
	//
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	|                                         |                                                                       |
	//	|                  VALUE                  |                                MEANING                                |
	//	|                                         |                                                                       |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_NP 0x00000002  | RPC protocol over named pipes is used by the DHCP server to register. |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_LPC 0x00000004 | RPC protocol over LPC is used by the DHCP server to register.         |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_ALL 0x00000007 | The DHCP server supports all the preceding protocols.                 |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	APIProtocolSupport uint32 `idl:"name:APIProtocolSupport" json:"api_protocol_support"`
	// DatabaseName:  A pointer of type LPWSTR to a null-terminated Unicode string that
	// represents the DHCP server database name that is used by the DHCP server for persistent
	// storage. There is no restriction on the length of this Unicode string. This field
	// MUST be convertible to an OEM or ANSI character string.
	DatabaseName string `idl:"name:DatabaseName" json:"database_name"`
	// DatabasePath:  A pointer of type LPWSTR to a null-terminated Unicode string that
	// contains the absolute path, where the DHCP server database is stored. The maximum
	// number of characters allowed in this field is 248, including the terminating null
	// character. This field MUST be convertible to an OEM or ANSI character string.
	DatabasePath string `idl:"name:DatabasePath" json:"database_path"`
	// BackupPath:  A pointer of type LPWSTR to a null-terminated Unicode string that contains
	// the absolute path for backup storage that is used by the DHCP server for backup.
	// The maximum number of characters allowed in this field is 248, including the terminating
	// null character. This field MUST be convertible to an OEM or ANSI character string.
	BackupPath string `idl:"name:BackupPath" json:"backup_path"`
	// BackupInterval:  This is of type DWORD, specifying the interval in minutes between
	// backups of the DHCP server database.
	BackupInterval uint32 `idl:"name:BackupInterval" json:"backup_interval"`
	// DatabaseLoggingFlag:  This is of type DWORD (used as a BOOL flag), indicating the
	// transaction logging mode of the DHCP server. The value 1 indicates that the transaction
	// log is enabled for the DHCP server, and zero indicates that the transaction log is
	// disabled for the DHCP server.
	DatabaseLoggingFlag uint32 `idl:"name:DatabaseLoggingFlag" json:"database_logging_flag"`
	// RestoreFlag:  This is of type DWORD (used as a BOOL flag), and if this setting is
	// TRUE, the DHCP server loads the DHCP database from the backup database on DHCP server
	// startup. The default value of this flag is FALSE.
	RestoreFlag uint32 `idl:"name:RestoreFlag" json:"restore_flag"`
	// DatabaseCleanupInterval:  This is of type DWORD and specifies the maximum time interval,
	// in minutes, for which DOOMED IPv4 DHCP client records are allowed to persist within
	// the DHCP server database.
	DatabaseCleanupInterval uint32 `idl:"name:DatabaseCleanupInterval" json:"database_cleanup_interval"`
	// DebugFlag:  A flag that specifies the level of logging done by the DHCP server. The
	// following table defines the set values that can be used. Specifying 0xFFFFFFFF enables
	// all types of logging.
	//
	// LOW WORD bitmask (0x0000FFFF) for low-frequency debug output.
	//
	//	+-----------------------------+----------------------------------------------------------------------+
	//	|                             |                                                                      |
	//	|            VALUE            |                               MEANING                                |
	//	|                             |                                                                      |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_ADDRESS 0x00000001    | Enable IP-address-related logging.                                   |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_CLIENT 0x00000002     | Enable DHCP-client-API-related logging.                              |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_PARAMETERS 0x00000004 | Enable DHCP-server-parameters-related logging.                       |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_OPTIONS 0x00000008    | Enable DHCP-options-related logging.                                 |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_ERRORS 0x00000010     | Enable DHCP-errors-related logging.                                  |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_STOC 0x00000020       | Enable DHCPv4 and DCHPv6-protocol-errors-related logging.            |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_INIT 0x00000040       | Enable DHCP-server-initialization-related logging.                   |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_SCAVENGER 0x00000080  | Enable scavenger's-error-related logging.                            |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_TIMESTAMP 0x00000100  | Enable timing-errors-related logging.                                |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_APIS 0x00000200       | Enable DHCP-APIs-related logging.                                    |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_REGISTRY 0x00000400   | Enable the logging of errors caused by registry setting operations.  |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_JET 0x00000800        | Enable the logging of the DHCP server database errors.               |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_THREADPOOL 0x00001000 | Enable the logging related to executing thread pool operations.      |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_AUDITLOG 0x00002000   | Enable the logging related to errors caused by audit log operations. |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_QUARANTINE 0x00004000 | Enable the logging of errors caused by quarantine errors.            |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_MISC 0x00008000       | Enable the logging caused by miscellaneous errors.                   |
	//	+-----------------------------+----------------------------------------------------------------------+
	//
	// HIGH WORD bitmask (0xFFFF0000) for high-frequency debug output, that is, more verbose.
	//
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	|                              |                                                                             |
	//	|            VALUE             |                                   MEANING                                   |
	//	|                              |                                                                             |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_MESSAGE 0x00010000     | Enable the logging related to debug messages.                               |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_API_VERBOSE 0x00020000 | Enable the logging related to DHCP API verbose errors.                      |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_DNS 0x00040000         | Enable the logging related to DNS messages.                                 |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_MSTOC 0x00080000       | Enable the logging related to multicast protocol layer errors.              |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_TRACK 0x00100000       | Enable the logging tracking specific problems.                              |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_ROGUE 0x00200000       | Enable the logging related to a rogue DHCP server.                          |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_PNP 0x00400000         | Enable the logging related to PNP interface errors.                         |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_PERF 0x01000000        | Enable the logging of performance-related messages.                         |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_ALLOC 0x02000000       | Enable the logging of allocation-related and deallocation-related messages. |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_PING 0x04000000        | Enable the logging of synchronous ping–related messages.                    |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_THREAD 0x08000000      | Enable the logging of thread-related messages.                              |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_TRACE 0x10000000       | Enable the logging for tracing through code messages.                       |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_TRACE_CALLS 0x20000000 | Enable the logging for tracing through piles of code.                       |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_STARTUP_BRK 0x40000000 | Enable the logging related to debugger break during setup messages.         |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_LOG_IN_FILE 0x80000000 | Enable the logging of debug output in a file.                               |
	//	+------------------------------+-----------------------------------------------------------------------------+
	DebugFlag uint32 `idl:"name:DebugFlag" json:"debug_flag"`
	// dwPingRetries:  This is of type DWORD, specifying the number of retries that the
	// DHCP server can make to verify whether a specific address is already in use by any
	// client by issuing a ping before issuing any address to the DHCP client (valid range:
	// 0–5, inclusive).
	PingRetries uint32 `idl:"name:dwPingRetries" json:"ping_retries"`
	// cbBootTableString:  This is of type DWORD, containing the size of the BOOT TABLE
	// given to the DHCP client. The maximum value of this field is 1048576.
	BootTableStringLength uint32 `idl:"name:cbBootTableString" json:"boot_table_string_length"`
	// wszBootTableString:  A pointer of type WCHAR* to a null-terminated Unicode string
	// that contains the absolute path of the BOOTP TABLE given to the BOOTP client. The
	// size of this string is limited to 1 MB.
	BootTableString string `idl:"name:wszBootTableString;size_is:(cbBootTableString)" json:"boot_table_string"`
	// fAuditLog:  This is a BOOL that represents whether an audit log needs to be written
	// by the DHCP server. The value of this member defaults to TRUE, which indicates that
	// the server writes an audit log.
	AuditLog bool `idl:"name:fAuditLog" json:"audit_log"`
}

func (o *ServerConfigInfoV4) xxx_PreparePayload(ctx context.Context) error {
	if o.BootTableString != "" && o.BootTableStringLength == 0 {
		o.BootTableStringLength = uint32(ndr.UTF16Len(o.BootTableString))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerConfigInfoV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.APIProtocolSupport); err != nil {
		return err
	}
	if o.DatabaseName != "" {
		_ptr_DatabaseName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DatabaseName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DatabaseName, _ptr_DatabaseName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DatabasePath != "" {
		_ptr_DatabasePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DatabasePath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DatabasePath, _ptr_DatabasePath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.BackupPath != "" {
		_ptr_BackupPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.BackupPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.BackupPath, _ptr_BackupPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BackupInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DatabaseLoggingFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.RestoreFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.DatabaseCleanupInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DebugFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.PingRetries); err != nil {
		return err
	}
	if err := w.WriteData(o.BootTableStringLength); err != nil {
		return err
	}
	if o.BootTableString != "" || o.BootTableStringLength > 0 {
		_ptr_wszBootTableString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BootTableStringLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_BootTableString_buf := utf16.Encode([]rune(o.BootTableString))
			if uint64(len(_BootTableString_buf)) > sizeInfo[0] {
				_BootTableString_buf = _BootTableString_buf[:sizeInfo[0]]
			}
			for i1 := range _BootTableString_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_BootTableString_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_BootTableString_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.BootTableString, _ptr_wszBootTableString); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.AuditLog {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ServerConfigInfoV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.APIProtocolSupport); err != nil {
		return err
	}
	_ptr_DatabaseName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DatabaseName); err != nil {
			return err
		}
		return nil
	})
	_s_DatabaseName := func(ptr interface{}) { o.DatabaseName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DatabaseName, _s_DatabaseName, _ptr_DatabaseName); err != nil {
		return err
	}
	_ptr_DatabasePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DatabasePath); err != nil {
			return err
		}
		return nil
	})
	_s_DatabasePath := func(ptr interface{}) { o.DatabasePath = *ptr.(*string) }
	if err := w.ReadPointer(&o.DatabasePath, _s_DatabasePath, _ptr_DatabasePath); err != nil {
		return err
	}
	_ptr_BackupPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.BackupPath); err != nil {
			return err
		}
		return nil
	})
	_s_BackupPath := func(ptr interface{}) { o.BackupPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.BackupPath, _s_BackupPath, _ptr_BackupPath); err != nil {
		return err
	}
	if err := w.ReadData(&o.BackupInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DatabaseLoggingFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.RestoreFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.DatabaseCleanupInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DebugFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.PingRetries); err != nil {
		return err
	}
	if err := w.ReadData(&o.BootTableStringLength); err != nil {
		return err
	}
	_ptr_wszBootTableString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BootTableStringLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BootTableStringLength)
		}
		var _BootTableString_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _BootTableString_buf", sizeInfo[0])
		}
		_BootTableString_buf = make([]uint16, sizeInfo[0])
		for i1 := range _BootTableString_buf {
			i1 := i1
			if err := w.ReadData(&_BootTableString_buf[i1]); err != nil {
				return err
			}
		}
		o.BootTableString = strings.TrimRight(string(utf16.Decode(_BootTableString_buf)), ndr.ZeroString)
		return nil
	})
	_s_wszBootTableString := func(ptr interface{}) { o.BootTableString = *ptr.(*string) }
	if err := w.ReadPointer(&o.BootTableString, _s_wszBootTableString, _ptr_wszBootTableString); err != nil {
		return err
	}
	var _bAuditLog int32
	if err := w.ReadData(&_bAuditLog); err != nil {
		return err
	}
	o.AuditLog = _bAuditLog != 0
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ServerConfigInfoVQ structure represents DHCP_SERVER_CONFIG_INFO_VQ RPC structure.
//
// The DHCP_SERVER_CONFIG_INFO_VQ structure defines settings for the DHCP server. This
// structure is an extension of the DHCP_SERVER_CONFIG_INFO_V4 (section 2.2.1.2.54)
// structure and is used in the R_DhcpServerSetConfigVQ (section 3.1.4.42) method.
type ServerConfigInfoVQ struct {
	// APIProtocolSupport: This is of type DWORD, defining the type of RPC protocol used
	// by the DHCP server to register with RPC. The following type MUST be supported.
	//
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	|                                           |                                                               |
	//	|                   VALUE                   |                            MEANING                            |
	//	|                                           |                                                               |
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_TCPIP 0x00000001 | RPC protocol over TCP is used by the DHCP server to register. |
	//	+-------------------------------------------+---------------------------------------------------------------+
	//
	// The following types MAY<14> be supported.
	//
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	|                                         |                                                                       |
	//	|                  VALUE                  |                                MEANING                                |
	//	|                                         |                                                                       |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_NP 0x00000002  | RPC protocol over named pipes is used by the DHCP server to register. |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_LPC 0x00000004 | RPC protocol over LPC is used by the DHCP server to register.         |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	//	| DHCP_SERVER_USE_RPC_OVER_ALL 0x00000007 | The DHCP server supports all of the preceding protocols.              |
	//	+-----------------------------------------+-----------------------------------------------------------------------+
	APIProtocolSupport uint32 `idl:"name:APIProtocolSupport" json:"api_protocol_support"`
	// DatabaseName: A pointer of type LPWSTR to a null-terminated Unicode string that represents
	// the database name that is used by the DHCP server for persistent storage. There is
	// no restriction on the length of this Unicode string. This field MUST be convertible
	// to an OEM or ANSI character string.
	DatabaseName string `idl:"name:DatabaseName" json:"database_name"`
	// DatabasePath: A pointer of type LPWSTR to a null-terminated Unicode string that contains
	// the absolute path where the DHCP server database is stored. The maximum number of
	// characters allowed in this field is 248, including the terminating null character.
	// This field MUST be convertible to an OEM or ANSI character string.
	DatabasePath string `idl:"name:DatabasePath" json:"database_path"`
	// BackupPath: A pointer of type LPWSTR to a null-terminated Unicode string that contains
	// the absolute path for the storage that is used by the DHCP server for backup. The
	// maximum number of characters allowed in this field is 248, including the terminating
	// null character. This field MUST be convertible to an OEM or ANSI character string.
	BackupPath string `idl:"name:BackupPath" json:"backup_path"`
	// BackupInterval: This is of type DWORD, containing the interval in minutes between
	// backups of the DHCP server database.
	BackupInterval uint32 `idl:"name:BackupInterval" json:"backup_interval"`
	// DatabaseLoggingFlag: This is of type DWORD (used as a BOOL flag), indicating the
	// transaction logging mode of the DHCP server. The value 1 indicates that transaction
	// logging mode is enabled for the DHCP server, and zero indicates that transaction
	// logging mode is disabled for the DHCP server.
	DatabaseLoggingFlag uint32 `idl:"name:DatabaseLoggingFlag" json:"database_logging_flag"`
	// RestoreFlag: This is of type DWORD (used as a BOOL flag), and if this setting is
	// TRUE, the DHCP server loads the DHCP database from the backup database on DHCP server
	// startup. The default value of this flag is FALSE.
	RestoreFlag uint32 `idl:"name:RestoreFlag" json:"restore_flag"`
	// DatabaseCleanupInterval: This is of type DWORD, and it specifies the time interval
	// in minutes over which the scavenger deletes the DOOMED IPv4 DHCP client lease records
	// from the DHCP server database.
	DatabaseCleanupInterval uint32 `idl:"name:DatabaseCleanupInterval" json:"database_cleanup_interval"`
	// DebugFlag: A flag that specifies the level of logging done by the DHCP server. The
	// following table defines the set values that can be used. Specifying 0xFFFFFFFF enables
	// all types of logging.
	//
	// LOW WORD bitmask (0x0000FFFF) for low-frequency debug output.
	//
	//	+-----------------------------+----------------------------------------------------------------------+
	//	|                             |                                                                      |
	//	|            VALUE            |                               MEANING                                |
	//	|                             |                                                                      |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_ADDRESS 0x00000001    | Enable IP-address-related logging.                                   |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_CLIENT 0x00000002     | Enable DHCP-client-API-related logging.                              |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_PARAMETERS 0x00000004 | Enable DHCP-server-parameters-related logging.                       |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_OPTIONS 0x00000008    | Enable DHCP-options-related logging.                                 |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_ERRORS 0x00000010     | Enable DHCP-errors-related logging.                                  |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_STOC 0x00000020       | Enable DHCPv4 and DCHPv6-protocol-errors-related logging.            |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_INIT 0x00000040       | Enable DHCP-server-initialization-related logging.                   |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_SCAVENGER 0x00000080  | Enable scavenger's-error-related logging.                            |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_TIMESTAMP 0x00000100  | Enable timing-errors-related logging.                                |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_APIS 0x00000200       | Enable DHCP-APIs-related logging.                                    |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_REGISTRY 0x00000400   | Enable the logging of errors caused by registry setting operations.  |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_JET 0x00000800        | Enable the logging of the DHCP server database errors.               |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_THREADPOOL 0x00001000 | Enable the logging related to executing thread pool operations.      |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_AUDITLOG 0x00002000   | Enable the logging related to errors caused by audit log operations. |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_QUARANTINE 0x00004000 | Enable the logging of errors caused by quarantine errors.            |
	//	+-----------------------------+----------------------------------------------------------------------+
	//	| DEBUG_MISC 0x00008000       | Enable the logging caused by miscellaneous errors.                   |
	//	+-----------------------------+----------------------------------------------------------------------+
	//
	// HIGH WORD bitmask (0xFFFF0000) for high-frequency debug output, that is, more verbose.
	//
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	|                              |                                                                             |
	//	|            VALUE             |                                   MEANING                                   |
	//	|                              |                                                                             |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_MESSAGE 0x00010000     | Enable the logging related to debug messages.                               |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_API_VERBOSE 0x00020000 | Enable the logging related to DHCP API verbose errors.                      |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_DNS 0x00040000         | Enable the logging related to DNS messages.                                 |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_MSTOC 0x00080000       | Enable the logging related to multicast protocol layer errors.              |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_TRACK 0x00100000       | Enable the logging tracking specific problems.                              |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_ROGUE 0x00200000       | Enable the logging related to a ROGUE DHCP server.                          |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_PNP 0x00400000         | Enable the logging related to PNP interface errors.                         |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_PERF 0x01000000        | Enable the logging of performance-related messages.                         |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_ALLOC 0x02000000       | Enable the logging of allocation-related and deallocation-related messages. |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_PING 0x04000000        | Enable the logging of synchronous ping–related messages.                    |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_THREAD 0x08000000      | Enable the logging of thread-related messages.                              |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_TRACE 0x10000000       | Enable the logging for tracing through code messages.                       |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_TRACE_CALLS 0x20000000 | Enable the logging for tracing through piles of code.                       |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_STARTUP_BRK 0x40000000 | Enable the logging related to debugger break during setup messages.         |
	//	+------------------------------+-----------------------------------------------------------------------------+
	//	| DEBUG_LOG_IN_FILE 0x80000000 | Enable the logging of debug output in a file.                               |
	//	+------------------------------+-----------------------------------------------------------------------------+
	DebugFlag uint32 `idl:"name:DebugFlag" json:"debug_flag"`
	// dwPingRetries: This is of type DWORD, specifying the number of retries that the DHCP
	// server can verify whether a specific address is already in use by any client by issuing
	// a ping before issuing any address to DHCP client (valid range: 0–5, inclusive).
	PingRetries uint32 `idl:"name:dwPingRetries" json:"ping_retries"`
	// cbBootTableString: This is of type DWORD, containing the size of the BOOT TABLE given
	// to the DHCP client. The maximum value of this field is 1048576.
	BootTableStringLength uint32 `idl:"name:cbBootTableString" json:"boot_table_string_length"`
	// wszBootTableString: A pointer of type WCHAR* to a null-terminated Unicode string
	// that contains the absolute path of the BOOT TABLE given to the DHCP client. The size
	// of this string is limited to 1 MB.
	BootTableString string `idl:"name:wszBootTableString;size_is:(cbBootTableString)" json:"boot_table_string"`
	// fAuditLog: This is of type BOOL, representing whether an audit log needs to be written
	// by the DHCP server. This member defaults to a value of TRUE, which indicates that
	// the DHCP server writes an audit log.
	AuditLog bool `idl:"name:fAuditLog" json:"audit_log"`
	// QuarantineOn: This is of type BOOL (a global flag), indicating whether quarantine
	// is on/off on the DHCP server. This member defaults to a value of FALSE, which indicates
	// that quarantine is off on the DHCP server.
	QuarantineOn bool `idl:"name:QuarantineOn" json:"quarantine_on"`
	// QuarDefFail: This is of type DWORD and determines the default policy for a DHCP NAP
	// server when an NPS server is not reachable. The range of permissible values is a
	// subset of those described in QuarantineStatus (section 2.2.1.1.11). Valid values
	// are NOQUARANTINE, RESTRICTEDACCESS, and DROPPACKET. This member defaults to a value
	// of NOQUARANTINE.
	QuarantineDefaultFail uint32 `idl:"name:QuarDefFail" json:"quarantine_default_fail"`
	// QuarRuntimeStatus: This flag determines whether NAP is enabled on the DHCP server
	// (Scope).
	QuarantineRuntimeStatus bool `idl:"name:QuarRuntimeStatus" json:"quarantine_runtime_status"`
}

func (o *ServerConfigInfoVQ) xxx_PreparePayload(ctx context.Context) error {
	if o.BootTableString != "" && o.BootTableStringLength == 0 {
		o.BootTableStringLength = uint32(ndr.UTF16Len(o.BootTableString))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerConfigInfoVQ) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.APIProtocolSupport); err != nil {
		return err
	}
	if o.DatabaseName != "" {
		_ptr_DatabaseName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DatabaseName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DatabaseName, _ptr_DatabaseName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DatabasePath != "" {
		_ptr_DatabasePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DatabasePath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DatabasePath, _ptr_DatabasePath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.BackupPath != "" {
		_ptr_BackupPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.BackupPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.BackupPath, _ptr_BackupPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BackupInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DatabaseLoggingFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.RestoreFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.DatabaseCleanupInterval); err != nil {
		return err
	}
	if err := w.WriteData(o.DebugFlag); err != nil {
		return err
	}
	if err := w.WriteData(o.PingRetries); err != nil {
		return err
	}
	if err := w.WriteData(o.BootTableStringLength); err != nil {
		return err
	}
	if o.BootTableString != "" || o.BootTableStringLength > 0 {
		_ptr_wszBootTableString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BootTableStringLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			_BootTableString_buf := utf16.Encode([]rune(o.BootTableString))
			if uint64(len(_BootTableString_buf)) > sizeInfo[0] {
				_BootTableString_buf = _BootTableString_buf[:sizeInfo[0]]
			}
			for i1 := range _BootTableString_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_BootTableString_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_BootTableString_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.BootTableString, _ptr_wszBootTableString); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.AuditLog {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.QuarantineOn {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.QuarantineDefaultFail); err != nil {
		return err
	}
	if !o.QuarantineRuntimeStatus {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ServerConfigInfoVQ) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.APIProtocolSupport); err != nil {
		return err
	}
	_ptr_DatabaseName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DatabaseName); err != nil {
			return err
		}
		return nil
	})
	_s_DatabaseName := func(ptr interface{}) { o.DatabaseName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DatabaseName, _s_DatabaseName, _ptr_DatabaseName); err != nil {
		return err
	}
	_ptr_DatabasePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DatabasePath); err != nil {
			return err
		}
		return nil
	})
	_s_DatabasePath := func(ptr interface{}) { o.DatabasePath = *ptr.(*string) }
	if err := w.ReadPointer(&o.DatabasePath, _s_DatabasePath, _ptr_DatabasePath); err != nil {
		return err
	}
	_ptr_BackupPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.BackupPath); err != nil {
			return err
		}
		return nil
	})
	_s_BackupPath := func(ptr interface{}) { o.BackupPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.BackupPath, _s_BackupPath, _ptr_BackupPath); err != nil {
		return err
	}
	if err := w.ReadData(&o.BackupInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DatabaseLoggingFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.RestoreFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.DatabaseCleanupInterval); err != nil {
		return err
	}
	if err := w.ReadData(&o.DebugFlag); err != nil {
		return err
	}
	if err := w.ReadData(&o.PingRetries); err != nil {
		return err
	}
	if err := w.ReadData(&o.BootTableStringLength); err != nil {
		return err
	}
	_ptr_wszBootTableString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BootTableStringLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BootTableStringLength)
		}
		var _BootTableString_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _BootTableString_buf", sizeInfo[0])
		}
		_BootTableString_buf = make([]uint16, sizeInfo[0])
		for i1 := range _BootTableString_buf {
			i1 := i1
			if err := w.ReadData(&_BootTableString_buf[i1]); err != nil {
				return err
			}
		}
		o.BootTableString = strings.TrimRight(string(utf16.Decode(_BootTableString_buf)), ndr.ZeroString)
		return nil
	})
	_s_wszBootTableString := func(ptr interface{}) { o.BootTableString = *ptr.(*string) }
	if err := w.ReadPointer(&o.BootTableString, _s_wszBootTableString, _ptr_wszBootTableString); err != nil {
		return err
	}
	var _bAuditLog int32
	if err := w.ReadData(&_bAuditLog); err != nil {
		return err
	}
	o.AuditLog = _bAuditLog != 0
	var _bQuarantineOn int32
	if err := w.ReadData(&_bQuarantineOn); err != nil {
		return err
	}
	o.QuarantineOn = _bQuarantineOn != 0
	if err := w.ReadData(&o.QuarantineDefaultFail); err != nil {
		return err
	}
	var _bQuarantineRuntimeStatus int32
	if err := w.ReadData(&_bQuarantineRuntimeStatus); err != nil {
		return err
	}
	o.QuarantineRuntimeStatus = _bQuarantineRuntimeStatus != 0
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ScopeMIBInfoVQ structure represents SCOPE_MIB_INFO_VQ RPC structure.
//
// The SCOPE_MIB_INFO_VQ structure contains the address counters for a specific IPv4
// subnet. The numbers of free, used, and offered IPv4 address are stored in this structure.
// This structure is an extension of the SCOPE_MIB_INFO (section 2.2.1.2.47) structure
// and is used in the DHCP_MIB_INFO_VQ (section 2.2.1.2.50) structure.
type ScopeMIBInfoVQ struct {
	// Subnet:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), a DWORD containing
	// the IPv4 subnet ID for the scope.
	Subnet uint32 `idl:"name:Subnet" json:"subnet"`
	// NumAddressesInuse:  This is of type DWORD, containing the number of IPv4 addresses
	// leased to DHCPv4 clients on a given IPv4 subnet.
	AddressesInUseLength uint32 `idl:"name:NumAddressesInuse" json:"addresses_inuse_length"`
	// NumAddressesFree:  This is of type DWORD, containing the number of IPv4 addresses
	// that are free and can be leased to DHCPv4 clients on a given IPv4 subnet.
	AddressesFreeLength uint32 `idl:"name:NumAddressesFree" json:"addresses_free_length"`
	// NumPendingOffers:  This is of type DWORD, containing the number of IPv4 addresses
	// that are offered to DHCPv4 clients on a given IPv4 subnet but which the DHCPv4 client
	// has not confirmed.
	PendingOffersLength uint32 `idl:"name:NumPendingOffers" json:"pending_offers_length"`
	// QtnNumLeases:  This field MUST be set to zero when sent and ignored on receipt. This
	// is treated as an error if it is nonzero in an RPC method that queries DHCPv4 server
	// configuration.
	QTNNumLeases uint32 `idl:"name:QtnNumLeases" json:"qtn_num_leases"`
	// QtnPctQtnLeases:  This field MUST be set to zero when sent and ignored on receipt.
	// This is treated as an error if it is nonzero in an RPC method that queries DHCPv4
	// server configuration.
	QTNPctQTNLeases uint32 `idl:"name:QtnPctQtnLeases" json:"qtn_pct_qtn_leases"`
	// QtnProbationLeases:  This field MUST be set to zero when sent and ignored on receipt.
	// This is treated as an error if it is nonzero in an RPC method that queries DHCPv4
	// server configuration.
	QTNProbationLeases uint32 `idl:"name:QtnProbationLeases" json:"qtn_probation_leases"`
	// QtnNonQtnLeases:  This field MUST be set to zero when sent and ignored on receipt.
	// This is treated as an error if it is nonzero in an RPC method that queries DHCPv4
	// server configuration.
	QTNNonQTNLeases uint32 `idl:"name:QtnNonQtnLeases" json:"qtn_non_qtn_leases"`
	// QtnExemptLeases:  This field MUST be set to zero when sent and ignored on receipt.
	// This is treated as an error if it is nonzero in an RPC method that queries DHCPv4
	// server configuration.
	QTNExemptLeases uint32 `idl:"name:QtnExemptLeases" json:"qtn_exempt_leases"`
	// QtnCapableClients:  This field MUST be set to zero when sent and ignored on receipt.
	// This is treated as an error if it is nonzero in an RPC method that queries DHCPv4
	// server configuration.
	QTNCapableClients uint32 `idl:"name:QtnCapableClients" json:"qtn_capable_clients"`
}

func (o *ScopeMIBInfoVQ) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ScopeMIBInfoVQ) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Subnet); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressesInUseLength); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressesFreeLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PendingOffersLength); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNNumLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNPctQTNLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNProbationLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNNonQTNLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNExemptLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNCapableClients); err != nil {
		return err
	}
	return nil
}
func (o *ScopeMIBInfoVQ) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Subnet); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressesInUseLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressesFreeLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PendingOffersLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNNumLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNPctQTNLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNProbationLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNNonQTNLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNExemptLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNCapableClients); err != nil {
		return err
	}
	return nil
}

// MIBInfoVQ structure represents DHCP_MIB_INFO_VQ RPC structure.
//
// The DHCP_MIB_INFO_VQ structure contains the counter values for the DHCP server. This
// structure is an extension of the DHCP_MIB_INFO (section 2.2.1.2.48) structure. This
// structure is used by an RPC method like R_DhcpGetMibInfoVQ (section 3.1.4.44) to
// return DHCP server statistics.
type MIBInfoVQ struct {
	// Discovers:  This is of type DWORD, containing the number of DHCPDISCOVER messages
	// [RFC2131] received by the DHCPv4 server from the DHCPv4 clients since the DHCP server
	// was last started. This is used for statistical analysis by the DHCPv4 server.
	Discovers uint32 `idl:"name:Discovers" json:"discovers"`
	// Offers:  This is of type DWORD, containing the number of DHCPOFFER messages sent
	// by the DHCPv4 server to the DHCPv4 client that the DHCPv4 client has not confirmed
	// since the DHCP server was last started. This is used for statistical analysis by
	// the DHCPv4 server.
	Offers uint32 `idl:"name:Offers" json:"offers"`
	// Requests:   This is of type DWORD, containing the number of DHCPREQUEST messages
	// received by the DHCPv4 server from the DHCPv4 clients since the DHCP server was last
	// started. This is used for statistical analysis by the DHCPv4 server.
	Requests uint32 `idl:"name:Requests" json:"requests"`
	// Acks:  This is of type DWORD, containing the number of DHCPACK messages sent by the
	// DHCPv4 server to the DHCPv4 client since the DHCP server was last started. This is
	// used for statistical analysis by the DHCPv4 server.
	ACKs uint32 `idl:"name:Acks" json:"acks"`
	// Naks:  This is of type DWORD, containing the number of DHCPNAK messages sent by the
	// DHCPv4 server to DHCPv4 clients since the DHCP server was last started. This is used
	// for statistical analysis by the DHCPv4 server.
	NAKs uint32 `idl:"name:Naks" json:"naks"`
	// Declines:   This is of type DWORD, containing the number of DHCPDECLINE messages
	// received by the DHCPv4 server from DHCPv4 clients since the DHCP server was last
	// started. This is used for statistical analysis by the DHCPv4 server.
	Declines uint32 `idl:"name:Declines" json:"declines"`
	// Releases:  This is of type DWORD, containing the number of DHCPRELEASE messages received
	// by the DHCPv4 server from DHCP clients since the DHCP server was last started. This
	// is used for statistical analysis by the DHCPv4 server.
	Releases uint32 `idl:"name:Releases" json:"releases"`
	// ServerStartTime:  This is of type DATE_TIME (section 2.2.1.2.11), containing the
	// start time of the DHCPv4 server.
	ServerStartTime *DateTime `idl:"name:ServerStartTime" json:"server_start_time"`
	// QtnNumLeases:  This is an unused field; it MUST be initialized to zero in an RPC
	// method that modifies the DHCPv4 server configuration and treated as an error if it
	// is nonzero in an RPC method that queries DHCPv4 server configuration.
	QTNNumLeases uint32 `idl:"name:QtnNumLeases" json:"qtn_num_leases"`
	// QtnPctQtnLeases:  This is an unused field; it MUST be initialized to zero in an RPC
	// method that modifies the DHCPv4 server configuration and treated as an error if it
	// is nonzero in an RPC method that queries DHCPv4 server configuration.
	QTNPctQTNLeases uint32 `idl:"name:QtnPctQtnLeases" json:"qtn_pct_qtn_leases"`
	// QtnProbationLeases:  This is an unused field; it MUST be initialized to zero in an
	// RPC method that modifies the DHCPv4 server configuration and treated as an error
	// if nonzero in an RPC method that queries DHCPv4 server configuration.
	QTNProbationLeases uint32 `idl:"name:QtnProbationLeases" json:"qtn_probation_leases"`
	// QtnNonQtnLeases:  This is an unused field; it MUST be initialized to zero in an RPC
	// method that modifies the DHCPv4 server configuration and treated as an error if it
	// is nonzero in an RPC method that queries DHCPv4 server configuration.
	QTNNonQTNLeases uint32 `idl:"name:QtnNonQtnLeases" json:"qtn_non_qtn_leases"`
	// QtnExemptLeases:  This is an unused field; it MUST be initialized to zero in an RPC
	// method that modifies the DHCPv4 server configuration and treated as an error if it
	// is nonzero in an RPC method that queries DHCPv4 server configuration.
	QTNExemptLeases uint32 `idl:"name:QtnExemptLeases" json:"qtn_exempt_leases"`
	// QtnCapableClients:  This is an unused field; it MUST be initialized to zero in an
	// RPC method that modifies the DHCPv4 server configuration and treated as an error
	// if nonzero in an RPC method that queries DHCPv4 server configuration.
	QTNCapableClients uint32 `idl:"name:QtnCapableClients" json:"qtn_capable_clients"`
	// QtnIASErrors:  This is an unused field; it MUST be initialized to zero in an RPC
	// method that modifies the DHCPv4 server configuration and treated as an error if nonzero
	// in an RPC method that queries DHCPv4 server configuration.
	QTNIASErrors uint32 `idl:"name:QtnIASErrors" json:"qtn_ias_errors"`
	// Scopes:  This is of type DWORD, containing the number of DHCPv4 scopes configured
	// on the current DHCPv4 server. This is used for statistical analysis by the DHCPv4
	// server. This field defines the number of DHCPv4 scopes in the subsequent field ScopeInfo.
	Scopes uint32 `idl:"name:Scopes" json:"scopes"`
	// ScopeInfo:  This is a pointer to an array SCOPE_MIB_INFO_VQ (section 2.2.1.2.49)
	// of length Scopes that contains the information about the IPv4 scopes configured on
	// DHCPv4 server.
	ScopeInfo []*ScopeMIBInfoVQ `idl:"name:ScopeInfo;size_is:(Scopes)" json:"scope_info"`
}

func (o *MIBInfoVQ) xxx_PreparePayload(ctx context.Context) error {
	if o.ScopeInfo != nil && o.Scopes == 0 {
		o.Scopes = uint32(len(o.ScopeInfo))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBInfoVQ) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Discovers); err != nil {
		return err
	}
	if err := w.WriteData(o.Offers); err != nil {
		return err
	}
	if err := w.WriteData(o.Requests); err != nil {
		return err
	}
	if err := w.WriteData(o.ACKs); err != nil {
		return err
	}
	if err := w.WriteData(o.NAKs); err != nil {
		return err
	}
	if err := w.WriteData(o.Declines); err != nil {
		return err
	}
	if err := w.WriteData(o.Releases); err != nil {
		return err
	}
	if o.ServerStartTime != nil {
		if err := o.ServerStartTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.QTNNumLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNPctQTNLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNProbationLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNNonQTNLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNExemptLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNCapableClients); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNIASErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.Scopes); err != nil {
		return err
	}
	if o.ScopeInfo != nil || o.Scopes > 0 {
		_ptr_ScopeInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Scopes)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ScopeInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ScopeInfo[i1] != nil {
					if err := o.ScopeInfo[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ScopeMIBInfoVQ{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ScopeInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ScopeMIBInfoVQ{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeInfo, _ptr_ScopeInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBInfoVQ) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Discovers); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offers); err != nil {
		return err
	}
	if err := w.ReadData(&o.Requests); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACKs); err != nil {
		return err
	}
	if err := w.ReadData(&o.NAKs); err != nil {
		return err
	}
	if err := w.ReadData(&o.Declines); err != nil {
		return err
	}
	if err := w.ReadData(&o.Releases); err != nil {
		return err
	}
	if o.ServerStartTime == nil {
		o.ServerStartTime = &DateTime{}
	}
	if err := o.ServerStartTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNNumLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNPctQTNLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNProbationLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNNonQTNLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNExemptLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNCapableClients); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNIASErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.Scopes); err != nil {
		return err
	}
	_ptr_ScopeInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Scopes > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Scopes)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ScopeInfo", sizeInfo[0])
		}
		o.ScopeInfo = make([]*ScopeMIBInfoVQ, sizeInfo[0])
		for i1 := range o.ScopeInfo {
			i1 := i1
			if o.ScopeInfo[i1] == nil {
				o.ScopeInfo[i1] = &ScopeMIBInfoVQ{}
			}
			if err := o.ScopeInfo[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ScopeInfo := func(ptr interface{}) { o.ScopeInfo = *ptr.(*[]*ScopeMIBInfoVQ) }
	if err := w.ReadPointer(&o.ScopeInfo, _s_ScopeInfo, _ptr_ScopeInfo); err != nil {
		return err
	}
	return nil
}

// ClientInfoVQ structure represents DHCP_CLIENT_INFO_VQ RPC structure.
//
// The DHCP_CLIENT_INFO_VQ structure defines information about the DHCPv4 client. This
// structure is used in the R_DhcpGetClientInfoVQ (section 3.1.4.47) method.
//
// DHCP_CLIENT_INFO_VQ augments the DHCP_CLIENT_INFO_V5 (section 2.2.1.2.16) structure
// by including information related to the NAP settings of the DHCPv4 client.
type ClientInfoVQ struct {
	// ClientIpAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), a DWORD that
	// contains the DHCPv4 client's IPv4 address.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// SubnetMask:  This is of type DHCP_IP_MASK (section 2.2.1.2.2), a DWORD that contains
	// the DHCPv4 client's IPv4 subnet mask address.
	SubnetMask uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
	// ClientHardwareAddress:  This is of type DHCP_CLIENT_UID (section 2.2.1.2.5), a structure
	// that represents a DHCPv4 client-identifier (section 2.2.1.2.5.1) or a DHCPv4 client
	// unique ID (section 2.2.1.2.5.2). Methods that accept DHCP_CLIENT_INFO_VQ as a parameter
	// specify which representations are acceptable.
	ClientHardwareAddress *ClientUID `idl:"name:ClientHardwareAddress" json:"client_hardware_address"`
	// ClientName:  A pointer to a null-terminated Unicode string that represents the DHCPv4
	// client's internet host name. There is no restriction on the length of this Unicode
	// string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
	// ClientComment:  A pointer to a null-terminated Unicode string that represents the
	// description given to the DHCPv4 client. There is no restriction on the length of
	// this Unicode string.
	ClientComment string `idl:"name:ClientComment" json:"client_comment"`
	// ClientLeaseExpires:  This is of type DATE_TIME (section 2.2.1.2.11), a structure
	// that contains the lease expiry time for the DHCPv4 client. This is UTC time represented
	// in the FILETIME format.
	ClientLeaseExpires *DateTime `idl:"name:ClientLeaseExpires" json:"client_lease_expires"`
	// OwnerHost:  This is of type DHCP_HOST_INFO (section 2.2.1.2.7), a structure that
	// contains information about the DHCPv4 server machine that has provided a lease to
	// the DHCPv4 client.
	OwnerHost *HostInfo `idl:"name:OwnerHost" json:"owner_host"`
	// bClientType:  This is of type BYTE that identifies the type of the DHCPv4 client.
	// The possible values are shown in the table that follows.
	//
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	|                                   |                                                                       |
	//	|               VALUE               |                                MEANING                                |
	//	|                                   |                                                                       |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_UNSPECIFIED 0x00      | A DHCPv4 client other than ones defined in this table.                |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_DHCP 0x01             | The DHCPv4 client supports the DHCP protocol.                         |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOOTP 0x02            | The DHCPv4 client supports the BOOTP protocol ([RFC2132]).            |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOTH 0x03             | The DHCPv4 client identifies both the DHCPv4 and the BOOTP protocols. |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_RESERVATION_FLAG 0x04 | There is an IPv4 reservation created for the DHCPv4 client.           |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_NONE 0x64             | Backward compatibility for manual addressing.                         |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	ClientType uint8 `idl:"name:bClientType" json:"client_type"`
	// AddressState:  This is of type BYTE, as shown by the following set of bits. The AddressState
	// member represents the state of the IPv4 address given to the DHCPv4 client.
	//
	//	+-------+------+------+------+------+------+------+------+
	//	|  BIT  |      |      |      |      |      |      |      |
	//	|   7   | BIT6 | BIT5 | BIT4 | BIT3 | BIT2 | BIT1 | BIT0 |
	//	|       |      |      |      |      |      |      |      |
	//	+-------+------+------+------+------+------+------+------+
	//	+-------+------+------+------+------+------+------+------+
	//
	// The following tables show the various bit representation values and their meanings.
	//
	// BIT 0 and BIT 1 signify the state of the leased IPv4 address, as shown in the table
	// that follows.
	//
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	|                            |                                                                                  |
	//	|           VALUE            |                                     MEANING                                      |
	//	|                            |                                                                                  |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_OFFERED 0x0  | The DHCPv4 client is offered this IPv4 address.                                  |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_ACTIVE 0x1   | The IPv4 address is active and has an active DHCPv4 client lease record.         |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DECLINED 0x2 | The IPv4 address request is declined by the DHCPv4 client; hence, it is a bad    |
	//	|                            | IPv4 address.                                                                    |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DOOM 0x3     | The IPv4 address is in DOOMED state and is due to be deleted.                    |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//
	// BIT 2 and BIT 3 signify the Name Protection (section 3.3.3) related information of
	// the leased IPv4 address, as shown in the table that follows.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                           |                                                                                  |
	//	|                   VALUE                   |                                     MEANING                                      |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_NO_DHCID 0x0                  | The address is leased to the DHCPv4 client without DHCID (sections 3 and 3.5 of  |
	//	|                                           | [RFC4701]).                                                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DHCID_NO_CLIENTIDOPTION 0x1   | The address is leased to the DHCPv4 client with DHCID as specified in section    |
	//	|                                           | 3.5.3 of [RFC4701].                                                              |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DHCID_WITH_CLIENTIDOPTION 0x2 | The address is leased to the DHCPv4 client with DHCID as specified in section    |
	//	|                                           | 3.5.2 of [RFC4701].                                                              |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DHCID_WITH_DUID 0x3           | The address is leased to the DHCPv4 client with DHCID as specified in section    |
	//	|                                           | 3.5.1 of [RFC4701].                                                              |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// BIT 4, BIT 5, BIT 6, and BIT 7 specify DNS-related information as shown in the table
	// that follows.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_CLEANUP 0x1      | The DNS update for the DHCPv4 client lease record needs to be deleted from the   |
	//	|                              | DNS server when the lease is deleted.                                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_BOTH_REC 0x2     | The DNS update needs to be sent for both A and PTR resource records ([RFC1034]   |
	//	|                              | section 3.6).                                                                    |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_UNREGISTERED 0x4 | The DNS update is not completed for the lease record.                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DELETED 0x8      | The address lease is expired, but the DNS updates for the lease record have not  |
	//	|                              | been deleted from the DNS server.                                                |
	//	+------------------------------+----------------------------------------------------------------------------------+
	AddressState uint8 `idl:"name:AddressState" json:"address_state"`
	// Status:  This is of type QuarantineStatus (section 2.2.1.1.11), and enumeration that
	// contains the health status of the DHCPv4 client, as validated at the NAP server.
	// Possible values validated by the NAP server are NOQUARANTINE, RESTRICTEDACCESS, DROPPACKET,
	// and PROBATION.
	Status QuarantineStatus `idl:"name:Status" json:"status"`
	// ProbationEnds:  This is of type DATE_TIME, a structure that contains the end time
	// of the probation if the DHCPv4 client is on probation. For this time period, the
	// DHCPv4 client has full access to the network.
	ProbationEnds *DateTime `idl:"name:ProbationEnds" json:"probation_ends"`
	// QuarantineCapable:  This is of type BOOL that takes on the values shown in the table
	// that follows.
	//
	//	+---------+------------------------------------------------------+
	//	|         |                                                      |
	//	|  VALUE  |                       MEANING                        |
	//	|         |                                                      |
	//	+---------+------------------------------------------------------+
	//	+---------+------------------------------------------------------+
	//	| TRUE 1  | The DHCPv4 client machine is quarantine-enabled.     |
	//	+---------+------------------------------------------------------+
	//	| FALSE 0 | The DHCPv4 client machine is not quarantine-enabled. |
	//	+---------+------------------------------------------------------+
	QuarantineCapable bool `idl:"name:QuarantineCapable" json:"quarantine_capable"`
}

func (o *ClientInfoVQ) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoVQ) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress != nil {
		if err := o.ClientHardwareAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientComment != "" {
		_ptr_ClientComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientComment, _ptr_ClientComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientLeaseExpires != nil {
		if err := o.ClientLeaseExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OwnerHost != nil {
		if err := o.OwnerHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ClientType); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressState); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if o.ProbationEnds != nil {
		if err := o.ProbationEnds.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.QuarantineCapable {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ClientInfoVQ) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress == nil {
		o.ClientHardwareAddress = &ClientUID{}
	}
	if err := o.ClientHardwareAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	_ptr_ClientComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClientComment := func(ptr interface{}) { o.ClientComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientComment, _s_ClientComment, _ptr_ClientComment); err != nil {
		return err
	}
	if o.ClientLeaseExpires == nil {
		o.ClientLeaseExpires = &DateTime{}
	}
	if err := o.ClientLeaseExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OwnerHost == nil {
		o.OwnerHost = &HostInfo{}
	}
	if err := o.OwnerHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressState); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	if o.ProbationEnds == nil {
		o.ProbationEnds = &DateTime{}
	}
	if err := o.ProbationEnds.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bQuarantineCapable int32
	if err := w.ReadData(&_bQuarantineCapable); err != nil {
		return err
	}
	o.QuarantineCapable = _bQuarantineCapable != 0
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ClientInfoArrayVQ structure represents DHCP_CLIENT_INFO_ARRAY_VQ RPC structure.
//
// The DHCP_CLIENT_INFO_ARRAY_VQ structure defines an array of DHCP_CLIENT_INFO_VQ (section
// 2.2.1.2.19) structures. This structure is used by methods, such as R_DhcpEnumSubnetClientsVQ
// (section 3.1.4.48), that retrieve information for more than one DHCPv4 client.
type ClientInfoArrayVQ struct {
	// NumElements:   This is of type DWORD, containing the number of clients in the specific
	// IPv4 subnet, which is also the number of entries in the Clients member element.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Clients:  This is a pointer of type DHCP_CLIENT_INFO_VQ that points to the array
	// of length NumElements containing the DHCP client information.
	Clients []*ClientInfoVQ `idl:"name:Clients;size_is:(NumElements)" json:"clients"`
}

func (o *ClientInfoArrayVQ) xxx_PreparePayload(ctx context.Context) error {
	if o.Clients != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Clients))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoArrayVQ) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Clients != nil || o.ElementsLength > 0 {
		_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Clients {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Clients[i1] != nil {
					_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Clients[i1] != nil {
							if err := o.Clients[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ClientInfoVQ{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Clients[i1], _ptr_Clients); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Clients); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Clients, _ptr_Clients); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoArrayVQ) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Clients", sizeInfo[0])
		}
		o.Clients = make([]*ClientInfoVQ, sizeInfo[0])
		for i1 := range o.Clients {
			i1 := i1
			_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Clients[i1] == nil {
					o.Clients[i1] = &ClientInfoVQ{}
				}
				if err := o.Clients[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Clients := func(ptr interface{}) { o.Clients[i1] = *ptr.(**ClientInfoVQ) }
			if err := w.ReadPointer(&o.Clients[i1], _s_Clients, _ptr_Clients); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Clients := func(ptr interface{}) { o.Clients = *ptr.(*[]*ClientInfoVQ) }
	if err := w.ReadPointer(&o.Clients, _s_Clients, _ptr_Clients); err != nil {
		return err
	}
	return nil
}

// SubnetInfoVQ structure represents DHCP_SUBNET_INFO_VQ RPC structure.
//
// The DHCP_SUBNET_INFO_VQ structure contains the information about an IPv4 subnet.
// This structure is an extension of the DHCP_SUBNET_INFO (section 2.2.1.2.8) structure,
// adding information on NAP state for the IPv4 subnet. This structure is used in the
// R_DhcpCreateSubnetVQ (section 3.1.4.49) method.
type SubnetInfoVQ struct {
	// SubnetAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), a DWORD specifying
	// the IPv4 subnet ID.
	SubnetAddress uint32 `idl:"name:SubnetAddress" json:"subnet_address"`
	// SubnetMask:  This is of type DHCP_IP_MASK (section 2.2.1.2.2), a DWORD specifying
	// the IPv4 subnet mask.
	SubnetMask uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
	// SubnetName:  A pointer of type LPWSTR to a null-terminated Unicode string that points
	// to the name of this IPv4 subnet. There is no restriction on the length of this Unicode
	// string.
	SubnetName string `idl:"name:SubnetName" json:"subnet_name"`
	// SubnetComment:  A pointer of type LPWSTR to a null-terminated Unicode string that
	// points to an optional comment specific to this IPv4 subnet. There is no restriction
	// on the length of this Unicode string.
	SubnetComment string `idl:"name:SubnetComment" json:"subnet_comment"`
	// PrimaryHost:  This is of type DHCP_HOST_INFO (section 2.2.1.2.7) structure that contains
	// information about the DHCPv4 server servicing this IPv4 subnet.
	PrimaryHost *HostInfo `idl:"name:PrimaryHost" json:"primary_host"`
	// SubnetState:  This is of type DHCP_SUBNET_STATE (section 2.2.1.1.2) enumeration that
	// indicates the current state of this IPv4 subnet.
	SubnetState SubnetState `idl:"name:SubnetState" json:"subnet_state"`
	// QuarantineOn:  The information relating to the NAP state of this IPv4 subnet.
	QuarantineOn uint32 `idl:"name:QuarantineOn" json:"quarantine_on"`
	// Reserved1:  This is of type DWORD. Currently it is not used, and any value set to
	// this field will not affect the behavior of the method that uses this structure. The
	// value returned in this parameter from the server is ignored.
	_ uint32 `idl:"name:Reserved1"`
	// Reserved2:  This is of type DWORD. Currently it is not used, and any value set to
	// this field will not affect the behavior of the method that uses this structure. The
	// value returned in this parameter from the server is ignored.
	_ uint32 `idl:"name:Reserved2"`
	// Reserved3:  This is of type INT64. Currently it is not used, and any value set to
	// this field will not affect the behavior of the method that uses this structure. The
	// value returned in this parameter from the server is ignored.
	_ int64 `idl:"name:Reserved3"`
	// Reserved4:  This is of type INT64. Currently it is not used, and any value set to
	// this field will not affect the behavior of the method that uses this structure. The
	// value returned in this parameter from the server is ignored.
	_ int64 `idl:"name:Reserved4"`
}

func (o *SubnetInfoVQ) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetInfoVQ) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	if o.SubnetName != "" {
		_ptr_SubnetName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SubnetName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SubnetName, _ptr_SubnetName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SubnetComment != "" {
		_ptr_SubnetComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SubnetComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SubnetComment, _ptr_SubnetComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PrimaryHost != nil {
		if err := o.PrimaryHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.SubnetState)); err != nil {
		return err
	}
	if err := w.WriteData(o.QuarantineOn); err != nil {
		return err
	}
	// reserved Reserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved Reserved2
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved Reserved3
	if err := w.WriteData(int64(0)); err != nil {
		return err
	}
	// reserved Reserved4
	if err := w.WriteData(int64(0)); err != nil {
		return err
	}
	return nil
}
func (o *SubnetInfoVQ) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	_ptr_SubnetName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SubnetName); err != nil {
			return err
		}
		return nil
	})
	_s_SubnetName := func(ptr interface{}) { o.SubnetName = *ptr.(*string) }
	if err := w.ReadPointer(&o.SubnetName, _s_SubnetName, _ptr_SubnetName); err != nil {
		return err
	}
	_ptr_SubnetComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SubnetComment); err != nil {
			return err
		}
		return nil
	})
	_s_SubnetComment := func(ptr interface{}) { o.SubnetComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.SubnetComment, _s_SubnetComment, _ptr_SubnetComment); err != nil {
		return err
	}
	if o.PrimaryHost == nil {
		o.PrimaryHost = &HostInfo{}
	}
	if err := o.PrimaryHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.SubnetState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.QuarantineOn); err != nil {
		return err
	}
	// reserved Reserved1
	var _Reserved1 uint32
	if err := w.ReadData(&_Reserved1); err != nil {
		return err
	}
	// reserved Reserved2
	var _Reserved2 uint32
	if err := w.ReadData(&_Reserved2); err != nil {
		return err
	}
	// reserved Reserved3
	var _Reserved3 int64
	if err := w.ReadData(&_Reserved3); err != nil {
		return err
	}
	// reserved Reserved4
	var _Reserved4 int64
	if err := w.ReadData(&_Reserved4); err != nil {
		return err
	}
	return nil
}

// OptionScopeType6 type represents DHCP_OPTION_SCOPE_TYPE6 RPC enumeration.
//
// The DHCP_OPTION_SCOPE_TYPE6 enumeration defines the type of DHCPv6 options being
// referred to by an RPC method in the DHCPM. The DHCP server allows for configuration
// of standard and vendor-specific options at various levels, such as the default level,
// server level, or scope level, or for a specific reservation. This value is used in
// conjunction with the DHCP_OPTION_SCOPE_INFO6 (section 2.2.1.2.30) structure, to specify
// option values in the RPC methods defined by this protocol.
type OptionScopeType6 uint16

var (
	// DhcpDefaultOptions6: Option is defined at the default level. The option definition
	// is created or modified on the DHCPv6 server and the default value of the option is
	// stored.
	OptionScopeType6DefaultOptions6 OptionScopeType6 = 0
	// DhcpScopeOptions6: Option is defined at the scope level. The option value is added
	// or modified at the scope and is valid for that specific scope.
	OptionScopeType6Options6 OptionScopeType6 = 1
	// DhcpReservedOptions6: Option is defined for a specific IP address reservation. The
	// option value is added or modified for a particular IP reservation in a scope.
	OptionScopeType6ReservedOptions6 OptionScopeType6 = 2
	// DhcpGlobalOptions6: Option is defined at the global level. The option value is added
	// or modified at the DHCPv6 server, which is valid for all scopes in that server.
	OptionScopeType6GlobalOptions6 OptionScopeType6 = 3
)

func (o OptionScopeType6) String() string {
	switch o {
	case OptionScopeType6DefaultOptions6:
		return "OptionScopeType6DefaultOptions6"
	case OptionScopeType6Options6:
		return "OptionScopeType6Options6"
	case OptionScopeType6ReservedOptions6:
		return "OptionScopeType6ReservedOptions6"
	case OptionScopeType6GlobalOptions6:
		return "OptionScopeType6GlobalOptions6"
	}
	return "Invalid"
}

// SubnetElementTypeV6 type represents DHCP_SUBNET_ELEMENT_TYPE_V6 RPC enumeration.
//
// The DHCP_SUBNET_ELEMENT_TYPE_V6 enumeration defines the type of a configuration parameter
// for a DHCPv6 scope configured on the DHCP server. This value is used in conjunction
// with other data types to specify the configuration parameters for a DHCPv6 scope
// by the RPC methods defined in this specification.
type SubnetElementTypeV6 uint16

var (
	// Dhcpv6IpRanges: The configuration parameter is not used, and it MUST NOT be used
	// by an RPC method defined in this specification. If this is used in any of the methods,
	// the method would return ERROR_INVALID_PARAMETER, except for R_DhcpAddSubnetElementV6
	// and R_DhcpRemoveSubnetElementV6, which return ERROR_SUCCESS.
	SubnetElementTypeV6IPRangesV6 SubnetElementTypeV6 = 0
	// Dhcpv6ReservedIps: The configuration parameter is a reservation for a DHCPv6 client
	// in a DHCPv6 scope element configured on the DHCP server.
	SubnetElementTypeV6ReservedIPsV6 SubnetElementTypeV6 = 1
	// Dhcpv6ExcludedIpRanges: The configuration parameter is the exclusion range of a DHCPv6
	// subnet configured on the DHCPv6 server.
	SubnetElementTypeV6ExcludedIPRangesV6 SubnetElementTypeV6 = 2
)

func (o SubnetElementTypeV6) String() string {
	switch o {
	case SubnetElementTypeV6IPRangesV6:
		return "SubnetElementTypeV6IPRangesV6"
	case SubnetElementTypeV6ReservedIPsV6:
		return "SubnetElementTypeV6ReservedIPsV6"
	case SubnetElementTypeV6ExcludedIPRangesV6:
		return "SubnetElementTypeV6ExcludedIPRangesV6"
	}
	return "Invalid"
}

// SearchInfoTypeV6 type represents DHCP_SEARCH_INFO_TYPE_V6 RPC enumeration.
//
// The DHCP_SEARCH_INFO_TYPE_V6 enumeration defines the field over which the search
// can be performed for a specific IPv6 DHCPv6 client lease record in the DHCPv6 server
// database. This enumeration is used in structure DHCP_SEARCH_INFO_V6 (section 2.2.1.2.69).
type SearchInfoTypeV6 uint16

var (
	// Dhcpv6ClientIpAddress: Use DHCPv6 client IPv6 address for searching the DHCPv6 IPv6
	// client lease record in the DHCP server.
	SearchInfoTypeV6ClientIPAddressV6 SearchInfoTypeV6 = 0
	// Dhcpv6ClientDUID: Use DHCPv6 client DUID (as specified in [RFC3315]) for searching
	// the DHCP IPv6 client lease record in the DHCPv6 server.
	SearchInfoTypeV6ClientDUIDV6 SearchInfoTypeV6 = 1
	// Dhcpv6ClientName: Use a null-terminated Unicode string that contains the name of
	// the DHCPv6 IPv6 client for searching for the DHCPv6 client lease record in the DHCPv6
	// server database.
	SearchInfoTypeV6ClientNameV6 SearchInfoTypeV6 = 2
)

func (o SearchInfoTypeV6) String() string {
	switch o {
	case SearchInfoTypeV6ClientIPAddressV6:
		return "SearchInfoTypeV6ClientIPAddressV6"
	case SearchInfoTypeV6ClientDUIDV6:
		return "SearchInfoTypeV6ClientDUIDV6"
	case SearchInfoTypeV6ClientNameV6:
		return "SearchInfoTypeV6ClientNameV6"
	}
	return "Invalid"
}

// IPv6Address structure represents DHCP_IPV6_ADDRESS RPC structure.
//
// The DHCP_IPV6_ADDRESS structure contains the IPv6 address. This is used in the DHCP_OPTION_SCOPE_INFO6
// (section 2.2.1.2.30) structure.
type IPv6Address struct {
	// HighOrderBits:  This is of type ULONGLONG, containing the higher 64 bits of the IPv6
	// address.
	HighOrderBits uint64 `idl:"name:HighOrderBits" json:"high_order_bits"`
	// LowOrderBits:  This is of type ULONGLONG, containing the lower 64 bits of the IPv6
	// address.
	LowOrderBits uint64 `idl:"name:LowOrderBits" json:"low_order_bits"`
}

func (o *IPv6Address) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPv6Address) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.HighOrderBits); err != nil {
		return err
	}
	if err := w.WriteData(o.LowOrderBits); err != nil {
		return err
	}
	return nil
}
func (o *IPv6Address) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.HighOrderBits); err != nil {
		return err
	}
	if err := w.ReadData(&o.LowOrderBits); err != nil {
		return err
	}
	return nil
}

// ClientInfoV5 structure represents DHCP_CLIENT_INFO_V5 RPC structure.
//
// The DHCP_CLIENT_INFO_V5 structure defines information about the DHCPv4 client. It
// augments the DHCP_CLIENT_INFO_V4 (section 2.2.1.2.14) structure by including the
// additional element AddressState. This structure is used in the DHCP_CLIENT_INFO_ARRAY_V5
// structure.
type ClientInfoV5 struct {
	// ClientIpAddress:   This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), a DWORD that
	// contains the DHCPv4 client's IPv4 address.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// SubnetMask:  This is of type DHCP_IP_MASK (section 2.2.1.2.2), a DWORD that contains
	// the DHCPv4 client's IPv4 subnet mask address.
	SubnetMask uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
	// ClientHardwareAddress:  This is of type DHCP_CLIENT_UID (section 2.2.1.2.5), a structure
	// that represents a DHCPv4 client-identifier (section 2.2.1.2.5.1) or a DHCPv4 client
	// unique ID (section 2.2.1.2.5.2). Methods that accept DHCP_CLIENT_INFO_V5 as a parameter
	// specify which representations are acceptable.
	ClientHardwareAddress *ClientUID `idl:"name:ClientHardwareAddress" json:"client_hardware_address"`
	// ClientName:  A pointer to a null-terminated Unicode string that represents the DHCPv4
	// client's internet host name. There is no restriction on the length of this Unicode
	// string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
	// ClientComment:  A pointer to a null-terminated Unicode string that represents the
	// description given to the DHCPv4 client. There is no restriction on the length of
	// this Unicode string.
	ClientComment string `idl:"name:ClientComment" json:"client_comment"`
	// ClientLeaseExpires:  This is of type DATE_TIME (section 2.2.1.2.11), a structure
	// that contains the lease expiry time for the DHCPv4 client. This is UTC time.
	ClientLeaseExpires *DateTime `idl:"name:ClientLeaseExpires" json:"client_lease_expires"`
	// OwnerHost:  This of type DHCP_HOST_INFO (section 2.2.1.2.7), a structure that contains
	// information about the DHCPv4 server machine that has provided a lease to the DHCPv4
	// client.
	OwnerHost *HostInfo `idl:"name:OwnerHost" json:"owner_host"`
	// bClientType:  This is of type BYTE that identifies the type of the DHCPv4 client.
	// Possible values for this field are provided in the following table.
	//
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	|                                   |                                                                       |
	//	|               VALUE               |                                MEANING                                |
	//	|                                   |                                                                       |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_UNSPECIFIED 0x00      | A DHCPv4 client other than ones defined in this table.                |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_DHCP 0x01             | The DHCPv4 client supports the DHCP protocol.                         |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOOTP 0x02            | The DHCPv4 client supports the BOOTP protocols ([RFC2132]).           |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOTH 0x03             | The DHCPv4 client identifies both the DHCPv4 and the BOOTP protocols. |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_RESERVATION_FLAG 0x04 | There is an IPv4 reservation created for the DHCPv4 client.           |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_NONE 0x64             | Backward compatibility for manual addressing.                         |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	ClientType uint8 `idl:"name:bClientType" json:"client_type"`
	// AddressState:  This is of type BYTE, as shown by the following set of bits. The AddressState
	// member represents the state of the IPv4 address given to the DHCPv4 client.
	//
	//	+-------+------+------+------+------+------+------+------+
	//	|  BIT  |      |      |      |      |      |      |      |
	//	|   7   | BIT6 | BIT5 | BIT4 | BIT3 | BIT2 | BIT1 | BIT0 |
	//	|       |      |      |      |      |      |      |      |
	//	+-------+------+------+------+------+------+------+------+
	//	+-------+------+------+------+------+------+------+------+
	//
	// The following tables show the various bit representation values and their meanings.
	//
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	|                            |                                                                                  |
	//	|           VALUE            |                                     MEANING                                      |
	//	|                            |                                                                                  |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_OFFERED 0x0  | The DHCPv4 client is offered this IPv4 address.                                  |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_ACTIVE 0x1   | The IPv4 address is active and has an active DHCPv4 client lease record.         |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DECLINED 0x2 | The IPv4 address request is declined by the DHCPv4 client; hence it is a bad     |
	//	|                            | IPv4 address.                                                                    |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DOOM 0x3     | The IPv4 address is in DOOMED state and is due to be deleted.                    |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//
	// BIT 2 and BIT 3 signify the Name Protection (section 3.3.3) related information of
	// the leased IPv4 address, as shown in the table that follows.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                           |                                                                                  |
	//	|                   VALUE                   |                                     MEANING                                      |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_NO_DHCID 0x0                  | The address is leased to the DHCPv4 client without DHCID (sections 3 and 3.5 of  |
	//	|                                           | [RFC4701]).                                                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DHCID_NO_CLIENTIDOPTION 0x1   | The address is leased to the DHCPv4 client with DHCID as specified in section    |
	//	|                                           | 3.5.3 of [RFC4701].                                                              |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DHCID_WITH_CLIENTIDOPTION 0x2 | The address is leased to the DHCPv4 client with DHCID as specified in section    |
	//	|                                           | 3.5.2 of [RFC4701].                                                              |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DHCID_WITH_DUID 0x3           | The address is leased to the DHCPv4 client with DHCID as specified in section    |
	//	|                                           | 3.5.1 of [RFC4701].                                                              |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// BIT 4, BIT 5, BIT 6, and BIT 7 specify DNS-related information as shown in the table
	// that follows.
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|                               |                                                                                  |
	//	|             VALUE             |                                     MEANING                                      |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_CLEANUP 0x10      | The DNS update for the DHCPv4 client lease record needs to be deleted from the   |
	//	|                               | DNS server when the lease is deleted.                                            |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_BOTH_REC 0x20     | The DNS update needs to be sent for both A and PTR resource records ([RFC1034]   |
	//	|                               | section 3.6).                                                                    |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_UNREGISTERED 0x40 | The DNS update is not completed for the lease record.                            |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DELETED 0x80      | The address lease is expired, but the DNS updates for the lease record have not  |
	//	|                               | been deleted from the DNS server.                                                |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	AddressState uint8 `idl:"name:AddressState" json:"address_state"`
}

func (o *ClientInfoV5) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoV5) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress != nil {
		if err := o.ClientHardwareAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientComment != "" {
		_ptr_ClientComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientComment, _ptr_ClientComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientLeaseExpires != nil {
		if err := o.ClientLeaseExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OwnerHost != nil {
		if err := o.OwnerHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ClientType); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressState); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ClientInfoV5) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress == nil {
		o.ClientHardwareAddress = &ClientUID{}
	}
	if err := o.ClientHardwareAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	_ptr_ClientComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClientComment := func(ptr interface{}) { o.ClientComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientComment, _s_ClientComment, _ptr_ClientComment); err != nil {
		return err
	}
	if o.ClientLeaseExpires == nil {
		o.ClientLeaseExpires = &DateTime{}
	}
	if err := o.ClientLeaseExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OwnerHost == nil {
		o.OwnerHost = &HostInfo{}
	}
	if err := o.OwnerHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressState); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ClientInfoArrayV5 structure represents DHCP_CLIENT_INFO_ARRAY_V5 RPC structure.
//
// The DHCP_CLIENT_INFO_ARRAY_V5 structure defines the array of DHCP_CLIENT_INFO_V5
// (section 2.2.1.2.16) structures.
//
// This structure is used by methods, such as R_DhcpEnumSubnetClientsV5 (section 3.2.4.1),
// that retrieve information for more than one DHCPv4 client.
type ClientInfoArrayV5 struct {
	// NumElements:   This is of type DWORD, containing the number of DHCPv4 client–specific
	// subnets, which is also the number of DHCPv4 clients in the Clients member element.
	// There are no inherent restrictions on the NumElements member. Methods that retrieve
	// DHCPv4 client information using the DHCP_CLIENT_INFO_ARRAY_V5 structure can limit
	// the maximum value of the NumElements member. For example, R_DhcpEnumSubnetClientsV5
	// restricts the number of elements based on input parameters and the size, as well
	// as the number, of DHCPv4 client lease records available for retrieval.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Clients:  This is a pointer of type DHCP_CLIENT_INFO_V5 that points to the array
	// of length NumElements containing the DHCPv4 client's information.
	Clients []*ClientInfoV5 `idl:"name:Clients;size_is:(NumElements)" json:"clients"`
}

func (o *ClientInfoArrayV5) xxx_PreparePayload(ctx context.Context) error {
	if o.Clients != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Clients))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoArrayV5) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Clients != nil || o.ElementsLength > 0 {
		_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Clients {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Clients[i1] != nil {
					_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Clients[i1] != nil {
							if err := o.Clients[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ClientInfoV5{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Clients[i1], _ptr_Clients); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Clients); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Clients, _ptr_Clients); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoArrayV5) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Clients", sizeInfo[0])
		}
		o.Clients = make([]*ClientInfoV5, sizeInfo[0])
		for i1 := range o.Clients {
			i1 := i1
			_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Clients[i1] == nil {
					o.Clients[i1] = &ClientInfoV5{}
				}
				if err := o.Clients[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Clients := func(ptr interface{}) { o.Clients[i1] = *ptr.(**ClientInfoV5) }
			if err := w.ReadPointer(&o.Clients[i1], _s_Clients, _ptr_Clients); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Clients := func(ptr interface{}) { o.Clients = *ptr.(*[]*ClientInfoV5) }
	if err := w.ReadPointer(&o.Clients, _s_Clients, _ptr_Clients); err != nil {
		return err
	}
	return nil
}

// MScopeInfo structure represents DHCP_MSCOPE_INFO RPC structure.
//
// The DHCP_MSCOPE_INFO structure defines the multicast scope information for a specific
// multicast subnet. This structure is used in the R_DhcpSetMScopeInfo (section 3.2.4.2)
// method.
type MScopeInfo struct {
	// MScopeName:  This is of type LPWSTR, containing a null-terminated Unicode string
	// that represents the multicast scope name. There is no restriction on the length of
	// this Unicode string.
	MScopeName string `idl:"name:MScopeName" json:"mscope_name"`
	// MScopeComment:  This is of type LPWSTR, containing a null-terminated Unicode string
	// that represents the description given to multicast scope. There is no restriction
	// on the length of this Unicode string.
	MScopeComment string `idl:"name:MScopeComment" json:"mscope_comment"`
	// MScopeId:  This is of type DWORD, containing the unique identification of the multicast
	// scope defined on the MADCAP server.
	MScopeID uint32 `idl:"name:MScopeId" json:"mscope_id"`
	// MScopeAddressPolicy:  This is of type DWORD. This MUST be set to zero when sent and
	// ignored on receipt.
	MScopeAddressPolicy uint32 `idl:"name:MScopeAddressPolicy" json:"mscope_address_policy"`
	// PrimaryHost:  This of type DHCP_HOST_INFO (section 2.2.1.2.7), a structure containing
	// information about the MADCAP server servicing this multicast scope.
	PrimaryHost *HostInfo `idl:"name:PrimaryHost" json:"primary_host"`
	// MScopeState:  This is of type DHCP_SUBNET_STATE (section 2.2.1.1.2), a structure
	// containing the state of the multicast subnet.
	MScopeState SubnetState `idl:"name:MScopeState" json:"mscope_state"`
	// MScopeFlags:  This is of type DWORD. Currently it is not used, and any value set
	// to this member will not affect the behavior of the method that uses this structure.
	MScopeFlags uint32 `idl:"name:MScopeFlags" json:"mscope_flags"`
	// ExpiryTime:  This is of type DATE_TIME (section 2.2.1.2.11), a structure specifying
	// the multicast scope lifetime.
	ExpiryTime *DateTime `idl:"name:ExpiryTime" json:"expiry_time"`
	// LangTag:  This is of type LPWSTR, containing a null-terminated Unicode string that
	// represents the multicast scope language (default is LOCALE_SYSTEM_DEFAULT). There
	// is no restriction on the length of this Unicode string.
	LangTag string `idl:"name:LangTag" json:"lang_tag"`
	// TTL:  This is of type BYTE, containing the Time-to-Live (TTL) value for the multicast
	// scope. The valid range for this field is between 1 and 255, with a default of 32.
	TTL uint8 `idl:"name:TTL" json:"ttl"`
}

func (o *MScopeInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MScopeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.MScopeName != "" {
		_ptr_MScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.MScopeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.MScopeName, _ptr_MScopeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.MScopeComment != "" {
		_ptr_MScopeComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.MScopeComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.MScopeComment, _ptr_MScopeComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MScopeID); err != nil {
		return err
	}
	if err := w.WriteData(o.MScopeAddressPolicy); err != nil {
		return err
	}
	if o.PrimaryHost != nil {
		if err := o.PrimaryHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.MScopeState)); err != nil {
		return err
	}
	if err := w.WriteData(o.MScopeFlags); err != nil {
		return err
	}
	if o.ExpiryTime != nil {
		if err := o.ExpiryTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LangTag != "" {
		_ptr_LangTag := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LangTag); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LangTag, _ptr_LangTag); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TTL); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *MScopeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_MScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.MScopeName); err != nil {
			return err
		}
		return nil
	})
	_s_MScopeName := func(ptr interface{}) { o.MScopeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.MScopeName, _s_MScopeName, _ptr_MScopeName); err != nil {
		return err
	}
	_ptr_MScopeComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.MScopeComment); err != nil {
			return err
		}
		return nil
	})
	_s_MScopeComment := func(ptr interface{}) { o.MScopeComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.MScopeComment, _s_MScopeComment, _ptr_MScopeComment); err != nil {
		return err
	}
	if err := w.ReadData(&o.MScopeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.MScopeAddressPolicy); err != nil {
		return err
	}
	if o.PrimaryHost == nil {
		o.PrimaryHost = &HostInfo{}
	}
	if err := o.PrimaryHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.MScopeState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.MScopeFlags); err != nil {
		return err
	}
	if o.ExpiryTime == nil {
		o.ExpiryTime = &DateTime{}
	}
	if err := o.ExpiryTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_LangTag := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LangTag); err != nil {
			return err
		}
		return nil
	})
	_s_LangTag := func(ptr interface{}) { o.LangTag = *ptr.(*string) }
	if err := w.ReadPointer(&o.LangTag, _s_LangTag, _ptr_LangTag); err != nil {
		return err
	}
	if err := w.ReadData(&o.TTL); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// MScopeTable structure represents DHCP_MSCOPE_TABLE RPC structure.
//
// The DHCP_MSCOPE_TABLE structure contains an array of multicast scope names managed
// by the MADCAP server. This structure is used in the R_DhcpEnumMScopes (section 3.2.4.4)
// method.
type MScopeTable struct {
	// NumElements:  This is of type DWORD, containing the number of multicast scope names
	// in the subsequent field the pMScopeNames member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// pMScopeNames:  This is a pointer of type LPWSTR that points to an array of null-terminated
	// Unicode strings that refers to the multicast scope names. There is no restriction
	// on the size of this field.
	MScopeNames []string `idl:"name:pMScopeNames;size_is:(NumElements)" json:"mscope_names"`
}

func (o *MScopeTable) xxx_PreparePayload(ctx context.Context) error {
	if o.MScopeNames != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.MScopeNames))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MScopeTable) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.MScopeNames != nil || o.ElementsLength > 0 {
		_ptr_pMScopeNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.MScopeNames {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.MScopeNames[i1] != "" {
					_ptr_pMScopeNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.MScopeNames[i1]); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.MScopeNames[i1], _ptr_pMScopeNames); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.MScopeNames); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.MScopeNames, _ptr_pMScopeNames); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MScopeTable) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_pMScopeNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.MScopeNames", sizeInfo[0])
		}
		o.MScopeNames = make([]string, sizeInfo[0])
		for i1 := range o.MScopeNames {
			i1 := i1
			_ptr_pMScopeNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.MScopeNames[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_pMScopeNames := func(ptr interface{}) { o.MScopeNames[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.MScopeNames[i1], _s_pMScopeNames, _ptr_pMScopeNames); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pMScopeNames := func(ptr interface{}) { o.MScopeNames = *ptr.(*[]string) }
	if err := w.ReadPointer(&o.MScopeNames, _s_pMScopeNames, _ptr_pMScopeNames); err != nil {
		return err
	}
	return nil
}

// MADCAPClientInfo structure represents DHCP_MCLIENT_INFO RPC structure.
//
// The DHCP_MCLIENT_INFO structure defines information about the MADCAP client that
// is used by the method R_DhcpGetMClientInfo (section 3.2.4.12).
type MADCAPClientInfo struct {
	// ClientIpAddress:   This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1) that contains
	// the MADCAP client's IPv4 address.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// MScopeId:   This is of type DWORD that specifies the unique identifier of the multicast
	// scope from which the MADCAP client receives an IPv4 multicast address.
	MScopeID uint32 `idl:"name:MScopeId" json:"mscope_id"`
	// ClientId:  This is of type DHCP_CLIENT_UID (section 2.2.1.2.5) that represents a
	// MADCAP lease identifier (section 2.2.1.2.5.4).
	ClientID *ClientUID `idl:"name:ClientId" json:"client_id"`
	// ClientName:  A pointer to a null-terminated Unicode string that represents the MADCAP
	// client's internet host name. There is no restriction on the length of this Unicode
	// string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
	// ClientLeaseStarts:  This is of type DATE_TIME (section 2.2.1.2.11) that contains
	// the lease start date and time for the MADCAP client. This is UTC time.
	ClientLeaseStarts *DateTime `idl:"name:ClientLeaseStarts" json:"client_lease_starts"`
	// ClientLeaseEnds:  This is of type DATE_TIME that contains the lease expiry time for
	// the MADCAP client. This is UTC time.
	ClientLeaseEnds *DateTime `idl:"name:ClientLeaseEnds" json:"client_lease_ends"`
	// OwnerHost:  This is of type DHCP_HOST_INFO (section 2.2.1.2.7) that contains information
	// about the MADCAP server machine that has provided a lease to the MADCAP client.
	OwnerHost *HostInfo `idl:"name:OwnerHost" json:"owner_host"`
	// AddressFlags:  This is of type DWORD. This MUST be set to zero when sent and ignored
	// on receipt. It MUST be specified as zero in the RPC method that modifies the MADCAP
	// server configuration. It MUST be treated as an error if the value is nonzero in the
	// RPC method that queries the MADCAP server configuration.
	AddressFlags uint32 `idl:"name:AddressFlags" json:"address_flags"`
	// AddressState:  This is of type BYTE that represents the state of the IPv4 address
	// given to the MADCAP client. The following table represents the different values and
	// their meanings.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_OFFERED 0x00000000  | The MADCAP client has been offered this IPv4 address.                            |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_ACTIVE 0x00000001   | The IPv4 address is active and has an active MADCAP client lease record.         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DECLINED 0x00000002 | The IPv4 address request was declined by the MADCAP client; hence it is a bad    |
	//	|                                   | IPv4 address.                                                                    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DOOM 0x00000003     | The IPv4 address is in DOOMED state and is due to be deleted.                    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	AddressState uint8 `idl:"name:AddressState" json:"address_state"`
}

func (o *MADCAPClientInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MADCAPClientInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.MScopeID); err != nil {
		return err
	}
	if o.ClientID != nil {
		if err := o.ClientID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientLeaseStarts != nil {
		if err := o.ClientLeaseStarts.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientLeaseEnds != nil {
		if err := o.ClientLeaseEnds.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OwnerHost != nil {
		if err := o.OwnerHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AddressFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressState); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *MADCAPClientInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.MScopeID); err != nil {
		return err
	}
	if o.ClientID == nil {
		o.ClientID = &ClientUID{}
	}
	if err := o.ClientID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	if o.ClientLeaseStarts == nil {
		o.ClientLeaseStarts = &DateTime{}
	}
	if err := o.ClientLeaseStarts.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ClientLeaseEnds == nil {
		o.ClientLeaseEnds = &DateTime{}
	}
	if err := o.ClientLeaseEnds.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OwnerHost == nil {
		o.OwnerHost = &HostInfo{}
	}
	if err := o.OwnerHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressState); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// MADCAPClientInfoArray structure represents DHCP_MCLIENT_INFO_ARRAY RPC structure.
//
// The DHCP_MCLIENT_INFO_ARRAY structure defines an array of DHCP_MCLIENT_INFO (section
// 2.2.1.2.21) structures. This structure is used by the methods that retrieve information
// for more than one MADCAP client. The first member contains the number of MADCAP clients
// present in the specific IPv4 multicast subnet, and the second member points to the
// array of length NumElements containing the MADCAP client's information.
type MADCAPClientInfoArray struct {
	// NumElements:  This is of type DWORD, specifying the number of MADCAP clients in subsequent
	// field the Clients member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Clients:  This is pointer of type DHCP_MCLIENT_INFO (section 2.2.1.2.21), a structure
	// that points to an array of length NumElements containing MADCAP client information.
	Clients []*MADCAPClientInfo `idl:"name:Clients;size_is:(NumElements)" json:"clients"`
}

func (o *MADCAPClientInfoArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Clients != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Clients))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MADCAPClientInfoArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Clients != nil || o.ElementsLength > 0 {
		_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Clients {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Clients[i1] != nil {
					_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Clients[i1] != nil {
							if err := o.Clients[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&MADCAPClientInfo{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Clients[i1], _ptr_Clients); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Clients); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Clients, _ptr_Clients); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MADCAPClientInfoArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Clients", sizeInfo[0])
		}
		o.Clients = make([]*MADCAPClientInfo, sizeInfo[0])
		for i1 := range o.Clients {
			i1 := i1
			_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Clients[i1] == nil {
					o.Clients[i1] = &MADCAPClientInfo{}
				}
				if err := o.Clients[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Clients := func(ptr interface{}) { o.Clients[i1] = *ptr.(**MADCAPClientInfo) }
			if err := w.ReadPointer(&o.Clients[i1], _s_Clients, _ptr_Clients); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Clients := func(ptr interface{}) { o.Clients = *ptr.(*[]*MADCAPClientInfo) }
	if err := w.ReadPointer(&o.Clients, _s_Clients, _ptr_Clients); err != nil {
		return err
	}
	return nil
}

// ReservedScope6 structure represents DHCP_RESERVED_SCOPE6 RPC structure.
//
// The DHCP_RESERVED_SCOPE6 structure defines an IPv6 reservation. This is used in DHCP_OPTION_SCOPE_INFO6
// (section 2.2.1.2.30) structure.
type ReservedScope6 struct {
	// ReservedIpAddress:  This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28), containing
	// theIPv6 address of an IPv6 reservation.
	ReservedIPAddress *IPv6Address `idl:"name:ReservedIpAddress" json:"reserved_ip_address"`
	// ReservedIpSubnetAddress:   This is of type DHCP_IPV6_ADDRESS, containing the IPv6
	// prefix ID of the subnet.
	ReservedIPSubnetAddress *IPv6Address `idl:"name:ReservedIpSubnetAddress" json:"reserved_ip_subnet_address"`
}

func (o *ReservedScope6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReservedScope6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ReservedIPAddress != nil {
		if err := o.ReservedIPAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ReservedIPSubnetAddress != nil {
		if err := o.ReservedIPSubnetAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReservedScope6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ReservedIPAddress == nil {
		o.ReservedIPAddress = &IPv6Address{}
	}
	if err := o.ReservedIPAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ReservedIPSubnetAddress == nil {
		o.ReservedIPSubnetAddress = &IPv6Address{}
	}
	if err := o.ReservedIPSubnetAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// OptionScopeInfo6 structure represents DHCP_OPTION_SCOPE_INFO6 RPC structure.
//
// The DHCP_OPTION_SCOPE_INFO6 structure contains information about the option. The
// information includes the type of the option and the level of the option (server level,
// scope level, or reservation level).
type OptionScopeInfo6 struct {
	// ScopeType:  This is of type DHCP_OPTION_SCOPE_TYPE6 (section 2.2.1.1.5) enumeration,
	// defining the scope type of the associated DHCP options, and indicates which of the
	// following fields in the union is used.
	ScopeType OptionScopeType6 `idl:"name:ScopeType" json:"scope_type"`
	// ScopeInfo:  This is a union that can contain one of the following values chosen based
	// on the value of ScopeType.
	ScopeInfo *OptionScopeInfo6_ScopeInfo `idl:"name:ScopeInfo;switch_is:ScopeType" json:"scope_info"`
}

func (o *OptionScopeInfo6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionScopeInfo6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ScopeType)); err != nil {
		return err
	}
	_swScopeInfo := uint16(o.ScopeType)
	if o.ScopeInfo != nil {
		if err := o.ScopeInfo.MarshalUnionNDR(ctx, w, _swScopeInfo); err != nil {
			return err
		}
	} else {
		if err := (&OptionScopeInfo6_ScopeInfo{}).MarshalUnionNDR(ctx, w, _swScopeInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionScopeInfo6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ScopeType)); err != nil {
		return err
	}
	if o.ScopeInfo == nil {
		o.ScopeInfo = &OptionScopeInfo6_ScopeInfo{}
	}
	_swScopeInfo := uint16(o.ScopeType)
	if err := o.ScopeInfo.UnmarshalUnionNDR(ctx, w, _swScopeInfo); err != nil {
		return err
	}
	return nil
}

// OptionScopeInfo6_ScopeInfo structure represents DHCP_OPTION_SCOPE_INFO6 union anonymous member.
//
// The DHCP_OPTION_SCOPE_INFO6 structure contains information about the option. The
// information includes the type of the option and the level of the option (server level,
// scope level, or reservation level).
type OptionScopeInfo6_ScopeInfo struct {
	// Types that are assignable to Value
	//
	// *OptionScopeInfo6_DefaultOptions
	// *OptionScopeInfo6_SubnetScopeInfo
	// *OptionScopeInfo6_ReservedScopeInfo
	// *OptionScopeInfo6_ReservedOptions
	Value is_OptionScopeInfo6_ScopeInfo `json:"value"`
}

func (o *OptionScopeInfo6_ScopeInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *OptionScopeInfo6_SubnetScopeInfo:
		if value != nil {
			return value.SubnetScopeInfo
		}
	case *OptionScopeInfo6_ReservedScopeInfo:
		if value != nil {
			return value.ReservedScopeInfo
		}
	}
	return nil
}

type is_OptionScopeInfo6_ScopeInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_OptionScopeInfo6_ScopeInfo()
}

func (o *OptionScopeInfo6_ScopeInfo) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *OptionScopeInfo6_DefaultOptions:
		return uint16(0)
	case *OptionScopeInfo6_SubnetScopeInfo:
		return uint16(1)
	case *OptionScopeInfo6_ReservedScopeInfo:
		return uint16(2)
	case *OptionScopeInfo6_ReservedOptions:
		return uint16(3)
	}
	return uint16(0)
}

func (o *OptionScopeInfo6_ScopeInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
	case uint16(1):
		_o, _ := o.Value.(*OptionScopeInfo6_SubnetScopeInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionScopeInfo6_SubnetScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*OptionScopeInfo6_ReservedScopeInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OptionScopeInfo6_ReservedScopeInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *OptionScopeInfo6_ScopeInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &OptionScopeInfo6_DefaultOptions{}
	case uint16(1):
		o.Value = &OptionScopeInfo6_SubnetScopeInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &OptionScopeInfo6_ReservedScopeInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &OptionScopeInfo6_ReservedOptions{}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// OptionScopeInfo6_DefaultOptions structure represents OptionScopeInfo6_ScopeInfo RPC union arm.
//
// It has following labels: 0
type OptionScopeInfo6_DefaultOptions struct {
}

func (*OptionScopeInfo6_DefaultOptions) is_OptionScopeInfo6_ScopeInfo() {}

func (o *OptionScopeInfo6_DefaultOptions) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *OptionScopeInfo6_DefaultOptions) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// OptionScopeInfo6_SubnetScopeInfo structure represents OptionScopeInfo6_ScopeInfo RPC union arm.
//
// It has following labels: 1
type OptionScopeInfo6_SubnetScopeInfo struct {
	// SubnetScopeInfo:  This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28) structure,
	// containing the IPv6 prefix ID of the subnet for which the option value is to be set.
	SubnetScopeInfo *IPv6Address `idl:"name:SubnetScopeInfo" json:"subnet_scope_info"`
}

func (*OptionScopeInfo6_SubnetScopeInfo) is_OptionScopeInfo6_ScopeInfo() {}

func (o *OptionScopeInfo6_SubnetScopeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SubnetScopeInfo != nil {
		if err := o.SubnetScopeInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionScopeInfo6_SubnetScopeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.SubnetScopeInfo == nil {
		o.SubnetScopeInfo = &IPv6Address{}
	}
	if err := o.SubnetScopeInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// OptionScopeInfo6_ReservedScopeInfo structure represents OptionScopeInfo6_ScopeInfo RPC union arm.
//
// It has following labels: 2
type OptionScopeInfo6_ReservedScopeInfo struct {
	// ReservedScopeInfo:  This is of type DHCP_RESERVED_SCOPE6 (section 2.2.1.2.29) structure,
	// containing the IPv6 address of the reservation and the IPv6 prefix ID for which the
	// option value is to be set.
	ReservedScopeInfo *ReservedScope6 `idl:"name:ReservedScopeInfo" json:"reserved_scope_info"`
}

func (*OptionScopeInfo6_ReservedScopeInfo) is_OptionScopeInfo6_ScopeInfo() {}

func (o *OptionScopeInfo6_ReservedScopeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ReservedScopeInfo != nil {
		if err := o.ReservedScopeInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ReservedScope6{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *OptionScopeInfo6_ReservedScopeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ReservedScopeInfo == nil {
		o.ReservedScopeInfo = &ReservedScope6{}
	}
	if err := o.ReservedScopeInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// OptionScopeInfo6_ReservedOptions structure represents OptionScopeInfo6_ScopeInfo RPC union arm.
//
// It has following labels: 3
type OptionScopeInfo6_ReservedOptions struct {
}

func (*OptionScopeInfo6_ReservedOptions) is_OptionScopeInfo6_ScopeInfo() {}

func (o *OptionScopeInfo6_ReservedOptions) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *OptionScopeInfo6_ReservedOptions) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// ClassInfo structure represents DHCP_CLASS_INFO RPC structure.
//
// The DHCP_CLASS_INFO structure contains the information for a specific user class
// or vendor class. This structure is used in the R_DhcpCreateClass (section 3.2.4.25)
// method.
type ClassInfo struct {
	// ClassName:  A pointer, of type LPWSTR, to a null-terminated Unicode string that contains
	// the class name. There is no restriction on the length of this Unicode string.
	ClassName string `idl:"name:ClassName" json:"class_name"`
	// ClassComment:  A pointer, of  type LPWSTR, to a null-terminated Unicode string that
	// contains the comment for the class. There is no restriction on the length of this
	// Unicode string.
	ClassComment string `idl:"name:ClassComment" json:"class_comment"`
	// ClassDataLength:  This is of type DWORD, containing the length of data as pointed
	// to by the ClassData member.
	ClassDataLength uint32 `idl:"name:ClassDataLength" json:"class_data_length"`
	// IsVendor:  This is of type BOOL and specifies whether the class is user class or
	// vendor class.
	//
	//	+------------+------------------------------------+
	//	|            |                                    |
	//	|   VALUE    |              MEANING               |
	//	|            |                                    |
	//	+------------+------------------------------------+
	//	+------------+------------------------------------+
	//	| 0x00000000 | Class specified is a user class.   |
	//	+------------+------------------------------------+
	//	| 0x00000001 | Class specified is a vendor class. |
	//	+------------+------------------------------------+
	IsVendor bool `idl:"name:IsVendor" json:"is_vendor"`
	// Flags:  This is of type DWORD. Currently it is not used, and any value set to this
	// member will not affect the behavior of the method that uses this structure.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ClassData:  This is a pointer of type BYTE that points to an array of bytes of length
	// specified by the ClassDataLength member. This contains data regarding a user class
	// or a vendor class.
	ClassData []byte `idl:"name:ClassData;size_is:(ClassDataLength)" json:"class_data"`
}

func (o *ClassInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.ClassData != nil && o.ClassDataLength == 0 {
		o.ClassDataLength = uint32(len(o.ClassData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ClassName != "" {
		_ptr_ClassName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClassName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClassName, _ptr_ClassName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClassComment != "" {
		_ptr_ClassComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClassComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClassComment, _ptr_ClassComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ClassDataLength); err != nil {
		return err
	}
	if !o.IsVendor {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.ClassData != nil || o.ClassDataLength > 0 {
		_ptr_ClassData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ClassDataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ClassData {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.ClassData[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.ClassData); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ClassData, _ptr_ClassData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_ClassName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClassName); err != nil {
			return err
		}
		return nil
	})
	_s_ClassName := func(ptr interface{}) { o.ClassName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClassName, _s_ClassName, _ptr_ClassName); err != nil {
		return err
	}
	_ptr_ClassComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClassComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClassComment := func(ptr interface{}) { o.ClassComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClassComment, _s_ClassComment, _ptr_ClassComment); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClassDataLength); err != nil {
		return err
	}
	var _bIsVendor int32
	if err := w.ReadData(&_bIsVendor); err != nil {
		return err
	}
	o.IsVendor = _bIsVendor != 0
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_ClassData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ClassDataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ClassDataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ClassData", sizeInfo[0])
		}
		o.ClassData = make([]byte, sizeInfo[0])
		for i1 := range o.ClassData {
			i1 := i1
			if err := w.ReadData(&o.ClassData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ClassData := func(ptr interface{}) { o.ClassData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ClassData, _s_ClassData, _ptr_ClassData); err != nil {
		return err
	}
	return nil
}

// ClassInfoArray structure represents DHCP_CLASS_INFO_ARRAY RPC structure.
//
// The DHCP_CLASS_INFO_ARRAY structure defines an array of DHCP_CLASS_INFO (section
// 2.2.1.2.75) structures. This structure is used by the methods that retrieve more
// than one class of information, such as the R_DhcpEnumClasses (section 3.2.4.29) method.
// The first member contains the number of classes present for the DHCPv4 server, and
// the second member contains a pointer to the array of length NumElements containing
// class information.
type ClassInfoArray struct {
	// NumElements:  This is of type DWORD, containing the count of classes in the subsequent
	// field the Classes member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Classes:  This is a pointer of type DHCP_CLASS_INFO (section 2.2.1.2.75) structure
	// that points to the array of length NumElements containing class information.
	Classes []*ClassInfo `idl:"name:Classes;size_is:(NumElements)" json:"classes"`
}

func (o *ClassInfoArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Classes != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Classes))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassInfoArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Classes != nil || o.ElementsLength > 0 {
		_ptr_Classes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Classes {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Classes[i1] != nil {
					if err := o.Classes[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ClassInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Classes); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ClassInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Classes, _ptr_Classes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassInfoArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Classes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Classes", sizeInfo[0])
		}
		o.Classes = make([]*ClassInfo, sizeInfo[0])
		for i1 := range o.Classes {
			i1 := i1
			if o.Classes[i1] == nil {
				o.Classes[i1] = &ClassInfo{}
			}
			if err := o.Classes[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Classes := func(ptr interface{}) { o.Classes = *ptr.(*[]*ClassInfo) }
	if err := w.ReadPointer(&o.Classes, _s_Classes, _ptr_Classes); err != nil {
		return err
	}
	return nil
}

// AllOptions structure represents DHCP_ALL_OPTIONS RPC structure.
//
// The DHCP_ALL_OPTIONS structure contains all the option definitions created on the
// DHCP server. This includes the vendor-specific option definition as well as the default
// vendor option definition. This structure is used in the R_DhcpGetAllOptions (section
// 3.2.4.30) method.
type AllOptions struct {
	// Flags:  This is of type DWORD. This MUST be set to zero when sent and ignored on
	// receipt.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// NonVendorOptions:  This is a pointer of type DHCP_OPTION_ARRAY (section 2.2.1.2.26)
	// structure that points to the location that contains all non-vendor-specific options
	// created on the DHCP server.
	NonVendorOptions *OptionArray `idl:"name:NonVendorOptions" json:"non_vendor_options"`
	// NumVendorOptions:  This is of type DWORD, containing the number of vendor-specific
	// options created on the DHCP server. This field specifies the number of vendor-specific
	// options defined in the subsequent field, the VendorOptions member.
	VendorOptionsLength uint32 `idl:"name:NumVendorOptions" json:"vendor_options_length"`
	// VendorOptions:  This structure defines the vendor-specific options.
	VendorOptions []*AllOptions_VendorOptions `idl:"name:VendorOptions;size_is:(NumVendorOptions)" json:"vendor_options"`
}

func (o *AllOptions) xxx_PreparePayload(ctx context.Context) error {
	if o.VendorOptions != nil && o.VendorOptionsLength == 0 {
		o.VendorOptionsLength = uint32(len(o.VendorOptions))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptions) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.NonVendorOptions != nil {
		_ptr_NonVendorOptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.NonVendorOptions != nil {
				if err := o.NonVendorOptions.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&OptionArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NonVendorOptions, _ptr_NonVendorOptions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VendorOptionsLength); err != nil {
		return err
	}
	if o.VendorOptions != nil || o.VendorOptionsLength > 0 {
		_ptr_VendorOptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.VendorOptionsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.VendorOptions {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.VendorOptions[i1] != nil {
					if err := o.VendorOptions[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&AllOptions_VendorOptions{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.VendorOptions); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&AllOptions_VendorOptions{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VendorOptions, _ptr_VendorOptions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptions) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_NonVendorOptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.NonVendorOptions == nil {
			o.NonVendorOptions = &OptionArray{}
		}
		if err := o.NonVendorOptions.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_NonVendorOptions := func(ptr interface{}) { o.NonVendorOptions = *ptr.(**OptionArray) }
	if err := w.ReadPointer(&o.NonVendorOptions, _s_NonVendorOptions, _ptr_NonVendorOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.VendorOptionsLength); err != nil {
		return err
	}
	_ptr_VendorOptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.VendorOptionsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.VendorOptionsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.VendorOptions", sizeInfo[0])
		}
		o.VendorOptions = make([]*AllOptions_VendorOptions, sizeInfo[0])
		for i1 := range o.VendorOptions {
			i1 := i1
			if o.VendorOptions[i1] == nil {
				o.VendorOptions[i1] = &AllOptions_VendorOptions{}
			}
			if err := o.VendorOptions[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_VendorOptions := func(ptr interface{}) { o.VendorOptions = *ptr.(*[]*AllOptions_VendorOptions) }
	if err := w.ReadPointer(&o.VendorOptions, _s_VendorOptions, _ptr_VendorOptions); err != nil {
		return err
	}
	return nil
}

// AllOptions_VendorOptions structure represents DHCP_ALL_OPTIONS structure anonymous member.
//
// The DHCP_ALL_OPTIONS structure contains all the option definitions created on the
// DHCP server. This includes the vendor-specific option definition as well as the default
// vendor option definition. This structure is used in the R_DhcpGetAllOptions (section
// 3.2.4.30) method.
type AllOptions_VendorOptions struct {
	// Option:  This is of type DHCP_OPTION (section 2.2.1.2.25), containing the option
	// definition for the specific vendor class and user class.
	Option *Option `idl:"name:Option" json:"option"`
	// VendorName:  A pointer to a null-terminated Unicode string that specifies the name
	// of a vendor class for a specific option definition. There is no restriction on the
	// length of this Unicode string.
	VendorName string `idl:"name:VendorName" json:"vendor_name"`
	// ClassName:  A pointer to a null-terminated Unicode string that specifies the name
	// of a user class for a specific user class. There is no restriction on the length
	// of this Unicode string.
	ClassName string `idl:"name:ClassName" json:"class_name"`
}

func (o *AllOptions_VendorOptions) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptions_VendorOptions) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Option != nil {
		if err := o.Option.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Option{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.VendorName != "" {
		_ptr_VendorName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.VendorName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VendorName, _ptr_VendorName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClassName != "" {
		_ptr_ClassName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClassName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClassName, _ptr_ClassName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptions_VendorOptions) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Option == nil {
		o.Option = &Option{}
	}
	if err := o.Option.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_VendorName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.VendorName); err != nil {
			return err
		}
		return nil
	})
	_s_VendorName := func(ptr interface{}) { o.VendorName = *ptr.(*string) }
	if err := w.ReadPointer(&o.VendorName, _s_VendorName, _ptr_VendorName); err != nil {
		return err
	}
	_ptr_ClassName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClassName); err != nil {
			return err
		}
		return nil
	})
	_s_ClassName := func(ptr interface{}) { o.ClassName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClassName, _s_ClassName, _ptr_ClassName); err != nil {
		return err
	}
	return nil
}

// AllOptionValues structure represents DHCP_ALL_OPTION_VALUES RPC structure.
//
// The DHCP_ALL_OPTION_VALUES structure contains all option values for a specified user
// class and vendor class. This structure is used in the R_DhcpGetAllOptionValuesV6
// (section 3.2.4.57) method.
type AllOptionValues struct {
	// Flags:  This is an unused field, and it MUST be initialized to 0 in an RPC method
	// that modifies the DHCP server configuration. This MUST be treated as an error if
	// it is nonzero in an RPC method that queries DHCP server configuration.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// NumElements:  This is of type DWORD, containing the number of options in the subsequent
	// field, the Options structure member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Options:  This structure defines the user and vendor-specific options
	Options []*AllOptionValues_Options `idl:"name:Options;size_is:(NumElements)" json:"options"`
}

func (o *AllOptionValues) xxx_PreparePayload(ctx context.Context) error {
	if o.Options != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Options))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptionValues) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Options != nil || o.ElementsLength > 0 {
		_ptr_Options := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Options {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Options[i1] != nil {
					if err := o.Options[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&AllOptionValues_Options{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Options); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&AllOptionValues_Options{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Options, _ptr_Options); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptionValues) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Options := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Options", sizeInfo[0])
		}
		o.Options = make([]*AllOptionValues_Options, sizeInfo[0])
		for i1 := range o.Options {
			i1 := i1
			if o.Options[i1] == nil {
				o.Options[i1] = &AllOptionValues_Options{}
			}
			if err := o.Options[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Options := func(ptr interface{}) { o.Options = *ptr.(*[]*AllOptionValues_Options) }
	if err := w.ReadPointer(&o.Options, _s_Options, _ptr_Options); err != nil {
		return err
	}
	return nil
}

// AllOptionValues_Options structure represents DHCP_ALL_OPTION_VALUES structure anonymous member.
//
// The DHCP_ALL_OPTION_VALUES structure contains all option values for a specified user
// class and vendor class. This structure is used in the R_DhcpGetAllOptionValuesV6
// (section 3.2.4.57) method.
type AllOptionValues_Options struct {
	// ClassName:  A pointer to a null-terminated Unicode string that contains the name
	// of the user class.
	ClassName string `idl:"name:ClassName" json:"class_name"`
	// VendorName:  A pointer to a null-terminated Unicode string that contains the name
	// of the vendor class.
	VendorName string `idl:"name:VendorName" json:"vendor_name"`
	// IsVendor:  This is of type BOOL that specifies whether this option set is specific
	// to a vendor class or default vendor class.
	IsVendor bool `idl:"name:IsVendor" json:"is_vendor"`
	// OptionsArray:  This is a pointer to an array of DHCP_OPTION_VALUE_ARRAY (section
	// 2.2.1.2.43) structures that points to an array of all the options for a specified
	// user class and vendor class.
	OptionsArray *OptionValueArray `idl:"name:OptionsArray" json:"options_array"`
}

func (o *AllOptionValues_Options) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptionValues_Options) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ClassName != "" {
		_ptr_ClassName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClassName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClassName, _ptr_ClassName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.VendorName != "" {
		_ptr_VendorName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.VendorName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VendorName, _ptr_VendorName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.IsVendor {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.OptionsArray != nil {
		_ptr_OptionsArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.OptionsArray != nil {
				if err := o.OptionsArray.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&OptionValueArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OptionsArray, _ptr_OptionsArray); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptionValues_Options) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_ClassName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClassName); err != nil {
			return err
		}
		return nil
	})
	_s_ClassName := func(ptr interface{}) { o.ClassName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClassName, _s_ClassName, _ptr_ClassName); err != nil {
		return err
	}
	_ptr_VendorName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.VendorName); err != nil {
			return err
		}
		return nil
	})
	_s_VendorName := func(ptr interface{}) { o.VendorName = *ptr.(*string) }
	if err := w.ReadPointer(&o.VendorName, _s_VendorName, _ptr_VendorName); err != nil {
		return err
	}
	var _bIsVendor int32
	if err := w.ReadData(&_bIsVendor); err != nil {
		return err
	}
	o.IsVendor = _bIsVendor != 0
	_ptr_OptionsArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.OptionsArray == nil {
			o.OptionsArray = &OptionValueArray{}
		}
		if err := o.OptionsArray.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_OptionsArray := func(ptr interface{}) { o.OptionsArray = *ptr.(**OptionValueArray) }
	if err := w.ReadPointer(&o.OptionsArray, _s_OptionsArray, _ptr_OptionsArray); err != nil {
		return err
	}
	return nil
}

// MScopeMIBInfo structure represents MSCOPE_MIB_INFO RPC structure.
//
// The MSCOPE_MIB_INFO structure defines the address counters for a specific multicast
// scope. The number of free, used, and offered addresses are stored in this structure.
// This structure is used in the DHCP_MCAST_MIB_INFO (section 2.2.1.2.52) structure.
type MScopeMIBInfo struct {
	// MScopeId:  This is of type DWORD, containing the unique identification of the multicast
	// scope defined on the DHCP server.
	MScopeID uint32 `idl:"name:MScopeId" json:"mscope_id"`
	// MScopeName:  This is of type LPWSTR, containing a null-terminated Unicode string
	// that points to the multicast scope name. There is no restriction on the length of
	// this Unicode string.
	MScopeName string `idl:"name:MScopeName" json:"mscope_name"`
	// NumAddressesInuse:  This is of type DWORD, containing the number of IPv4 multicast
	// addresses that are leased out to MADCAP clients from a given multicast scope.
	AddressesInUseLength uint32 `idl:"name:NumAddressesInuse" json:"addresses_inuse_length"`
	// NumAddressesFree:  This is of type DWORD, containing the number of IPv4 multicast
	// addresses that are free and can be leased out to MADCAP clients from a given multicast
	// scope.
	AddressesFreeLength uint32 `idl:"name:NumAddressesFree" json:"addresses_free_length"`
	// NumPendingOffers:  This is of type DWORD, containing the number of IPv4 multicast
	// addresses that are offered to MADCAP clients in a specific IPv4 subnet but that the
	// MADCAP client has not confirmed.
	PendingOffersLength uint32 `idl:"name:NumPendingOffers" json:"pending_offers_length"`
}

func (o *MScopeMIBInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MScopeMIBInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MScopeID); err != nil {
		return err
	}
	if o.MScopeName != "" {
		_ptr_MScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.MScopeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.MScopeName, _ptr_MScopeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AddressesInUseLength); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressesFreeLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PendingOffersLength); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *MScopeMIBInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MScopeID); err != nil {
		return err
	}
	_ptr_MScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.MScopeName); err != nil {
			return err
		}
		return nil
	})
	_s_MScopeName := func(ptr interface{}) { o.MScopeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.MScopeName, _s_MScopeName, _ptr_MScopeName); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressesInUseLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressesFreeLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PendingOffersLength); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// MCastMIBInfo structure represents DHCP_MCAST_MIB_INFO RPC structure.
//
// The DHCP_MCAST_MIB_INFO structure contains counter values for all multicast scopes
// defined on the MADCAP server. This structure is used in R_DhcpGetMCastMibInfo (section
// 3.2.4.32) method which retrieves statistics on multicast scopes defined on the MADCAP
// server.
type MCastMIBInfo struct {
	// Discovers:  This is of type DWORD, containing the number of DHCPDISCOVER messages
	// [RFC2131] received by the MADCAP server from MADCAP clients.
	Discovers uint32 `idl:"name:Discovers" json:"discovers"`
	// Offers:  This is of type DWORD, containing the number of DHCPOFFER messages sent
	// by the MADCAP server to the MADCAP client.
	Offers uint32 `idl:"name:Offers" json:"offers"`
	// Requests:  This is of type DWORD, containing the number of DHCPREQUEST messages received
	// by the MADCAP server from MADCAP clients.
	Requests uint32 `idl:"name:Requests" json:"requests"`
	// Renews:  This is of type DWORD, containing the number of DHCPRENEW messages received
	// by the MADCAP server from MADCAP clients.
	Renews uint32 `idl:"name:Renews" json:"renews"`
	// Acks:  This is of type DWORD, containing the number of DHCPACK messages sent by the
	// MADCAP server to the MADCAP client.
	ACKs uint32 `idl:"name:Acks" json:"acks"`
	// Naks:  This is of type DWORD, containing the number of DHCPNAK messages sent by the
	// MADCAP server to MADCAP clients.
	NAKs uint32 `idl:"name:Naks" json:"naks"`
	// Releases:  This is of type DWORD, containing the number of DHCPRELEASE messages received
	// by the MADCAP server from the MADCAP client.
	Releases uint32 `idl:"name:Releases" json:"releases"`
	// Informs:  This is of type DWORD, containing the number of DHCPINFORM messages received
	// by the MADCAP server from the MADCAP client.
	Informs uint32 `idl:"name:Informs" json:"informs"`
	// ServerStartTime:  This is of type DATE_TIME (section 2.2.1.2.11), containing the
	// start time of the MADCAP server.
	ServerStartTime *DateTime `idl:"name:ServerStartTime" json:"server_start_time"`
	// Scopes:  This is of type DWORD, containing the number of IPv4 multicast scopes configured
	// on the current MADCAP server. This field defines the number of IPv4 multicast scopes
	// in the subsequent field ScopeInfo.
	Scopes uint32 `idl:"name:Scopes" json:"scopes"`
	// ScopeInfo:  This is a pointer to an array of MSCOPE_MIB_INFO (section 2.2.1.2.51)
	// structures of length Scopes that contains information about the IPv4 scopes configured
	// on the MADCAP server.
	ScopeInfo []*MScopeMIBInfo `idl:"name:ScopeInfo;size_is:(Scopes)" json:"scope_info"`
}

func (o *MCastMIBInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.ScopeInfo != nil && o.Scopes == 0 {
		o.Scopes = uint32(len(o.ScopeInfo))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MCastMIBInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Discovers); err != nil {
		return err
	}
	if err := w.WriteData(o.Offers); err != nil {
		return err
	}
	if err := w.WriteData(o.Requests); err != nil {
		return err
	}
	if err := w.WriteData(o.Renews); err != nil {
		return err
	}
	if err := w.WriteData(o.ACKs); err != nil {
		return err
	}
	if err := w.WriteData(o.NAKs); err != nil {
		return err
	}
	if err := w.WriteData(o.Releases); err != nil {
		return err
	}
	if err := w.WriteData(o.Informs); err != nil {
		return err
	}
	if o.ServerStartTime != nil {
		if err := o.ServerStartTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Scopes); err != nil {
		return err
	}
	if o.ScopeInfo != nil || o.Scopes > 0 {
		_ptr_ScopeInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Scopes)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ScopeInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ScopeInfo[i1] != nil {
					if err := o.ScopeInfo[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&MScopeMIBInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ScopeInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&MScopeMIBInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeInfo, _ptr_ScopeInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MCastMIBInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Discovers); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offers); err != nil {
		return err
	}
	if err := w.ReadData(&o.Requests); err != nil {
		return err
	}
	if err := w.ReadData(&o.Renews); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACKs); err != nil {
		return err
	}
	if err := w.ReadData(&o.NAKs); err != nil {
		return err
	}
	if err := w.ReadData(&o.Releases); err != nil {
		return err
	}
	if err := w.ReadData(&o.Informs); err != nil {
		return err
	}
	if o.ServerStartTime == nil {
		o.ServerStartTime = &DateTime{}
	}
	if err := o.ServerStartTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Scopes); err != nil {
		return err
	}
	_ptr_ScopeInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Scopes > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Scopes)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ScopeInfo", sizeInfo[0])
		}
		o.ScopeInfo = make([]*MScopeMIBInfo, sizeInfo[0])
		for i1 := range o.ScopeInfo {
			i1 := i1
			if o.ScopeInfo[i1] == nil {
				o.ScopeInfo[i1] = &MScopeMIBInfo{}
			}
			if err := o.ScopeInfo[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ScopeInfo := func(ptr interface{}) { o.ScopeInfo = *ptr.(*[]*MScopeMIBInfo) }
	if err := w.ReadPointer(&o.ScopeInfo, _s_ScopeInfo, _ptr_ScopeInfo); err != nil {
		return err
	}
	return nil
}

// Attribute structure represents DHCP_ATTRIB RPC structure.
//
// The DHCP_ATTRIB structure contains the attribute and its values for the DHCPv4 server.
// This structure is used in the R_DhcpServerQueryAttribute (section 3.2.4.35) method.
type Attribute struct {
	// DhcpAttribId:  This is of type DHCP_ATTRIB_ID (section 2.2.1.1.1), a ULONG value
	// specifying the attribute.
	AttributeID uint32 `idl:"name:DhcpAttribId" json:"attribute_id"`
	// DhcpAttribType:  This is of type ULONG. The value specifies the type of the attribute's
	// data and which one of the values is chosen from the subsequent union.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| DHCP_ATTRIB_TYPE_BOOL 0x00000001  | The attribute value is of type BOOL, and DhcpAttribBool is chosen from the       |
	//	|                                   | following union.                                                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| DHCP_ATTRIB_TYPE_ULONG 0x00000002 | The attribute value is of type ULONG, and DhcpAttribUlong is chosen from the     |
	//	|                                   | following union.                                                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	AttributeType uint32               `idl:"name:DhcpAttribType" json:"attribute_type"`
	Attribute     *Attribute_Attribute `idl:"name:Attribute;switch_is:DhcpAttribType" json:"attribute"`
}

func (o *Attribute) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Attribute) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeID); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeType); err != nil {
		return err
	}
	_swAttribute := uint32(o.AttributeType)
	if o.Attribute != nil {
		if err := o.Attribute.MarshalUnionNDR(ctx, w, _swAttribute); err != nil {
			return err
		}
	} else {
		if err := (&Attribute_Attribute{}).MarshalUnionNDR(ctx, w, _swAttribute); err != nil {
			return err
		}
	}
	return nil
}
func (o *Attribute) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeType); err != nil {
		return err
	}
	if o.Attribute == nil {
		o.Attribute = &Attribute_Attribute{}
	}
	_swAttribute := uint32(o.AttributeType)
	if err := o.Attribute.UnmarshalUnionNDR(ctx, w, _swAttribute); err != nil {
		return err
	}
	return nil
}

// Attribute_Attribute structure represents DHCP_ATTRIB union anonymous member.
//
// The DHCP_ATTRIB structure contains the attribute and its values for the DHCPv4 server.
// This structure is used in the R_DhcpServerQueryAttribute (section 3.2.4.35) method.
type Attribute_Attribute struct {
	// Types that are assignable to Value
	//
	// *Attribute_Bool
	// *Attribute_Uint32
	Value is_Attribute_Attribute `json:"value"`
}

func (o *Attribute_Attribute) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Attribute_Bool:
		if value != nil {
			return value.AttributeBool
		}
	case *Attribute_Uint32:
		if value != nil {
			return value.AttributeUint32
		}
	}
	return nil
}

type is_Attribute_Attribute interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Attribute_Attribute()
}

func (o *Attribute_Attribute) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Attribute_Bool:
		return uint32(1)
	case *Attribute_Uint32:
		return uint32(2)
	}
	return uint32(0)
}

func (o *Attribute_Attribute) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*Attribute_Bool)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Attribute_Bool{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*Attribute_Uint32)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Attribute_Uint32{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *Attribute_Attribute) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(4); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &Attribute_Bool{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &Attribute_Uint32{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// Attribute_Bool structure represents Attribute_Attribute RPC union arm.
//
// It has following labels: 1
type Attribute_Bool struct {
	// DhcpAttribBool:  This is of type BOOL and contains the value of the attribute. This
	// is chosen from the union if DhcpAttribType contains DHCP_ATTRIB_TYPE_BOOL.
	AttributeBool bool `idl:"name:DhcpAttribBool" json:"attribute_bool"`
}

func (*Attribute_Bool) is_Attribute_Attribute() {}

func (o *Attribute_Bool) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if !o.AttributeBool {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Attribute_Bool) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	var _bAttributeBool int32
	if err := w.ReadData(&_bAttributeBool); err != nil {
		return err
	}
	o.AttributeBool = _bAttributeBool != 0
	return nil
}

// Attribute_Uint32 structure represents Attribute_Attribute RPC union arm.
//
// It has following labels: 2
type Attribute_Uint32 struct {
	// DhcpAttribUlong:  This is of type ULONG and contains the value of the attribute.
	// This is chosen from the union if DhcpAttribType  contains DHCP_ATTRIB_TYPE_ULONG.
	AttributeUint32 uint32 `idl:"name:DhcpAttribUlong" json:"attribute_uint32"`
}

func (*Attribute_Uint32) is_Attribute_Attribute() {}

func (o *Attribute_Uint32) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.AttributeUint32); err != nil {
		return err
	}
	return nil
}
func (o *Attribute_Uint32) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.AttributeUint32); err != nil {
		return err
	}
	return nil
}

// AttributeArray structure represents DHCP_ATTRIB_ARRAY RPC structure.
//
// The DHCP_ATTRIB_ARRAY structure defines an array of DHCP_ATTRIB (section 2.2.1.2.78)
// structures. This structure is used by methods that retrieve more than one attribute,
// such as the R_DhcpServerQueryAttributes (section 3.2.4.36) method. This structure
// defines an array of length NumElements that contains attributes and their values.
type AttributeArray struct {
	// NumElements:  This is of type ULONG, containing the number of attributes in the subsequent
	// field the DhcpAttribs member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// DhcpAttribs:  This is a pointer to an array of type DHCP_ATTRIB (section 2.2.1.2.78)
	// structure and of length NumElements that contains the attributes and its values.
	Attributes []*Attribute `idl:"name:DhcpAttribs;size_is:(NumElements)" json:"attributes"`
}

func (o *AttributeArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Attributes != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Attributes))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AttributeArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Attributes != nil || o.ElementsLength > 0 {
		_ptr_DhcpAttribs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
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
				if o.Attributes[i1] != nil {
					if err := o.Attributes[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Attribute{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Attributes); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Attribute{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Attributes, _ptr_DhcpAttribs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AttributeArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_DhcpAttribs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Attributes", sizeInfo[0])
		}
		o.Attributes = make([]*Attribute, sizeInfo[0])
		for i1 := range o.Attributes {
			i1 := i1
			if o.Attributes[i1] == nil {
				o.Attributes[i1] = &Attribute{}
			}
			if err := o.Attributes[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_DhcpAttribs := func(ptr interface{}) { o.Attributes = *ptr.(*[]*Attribute) }
	if err := w.ReadPointer(&o.Attributes, _s_DhcpAttribs, _ptr_DhcpAttribs); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV5 structure represents DHCP_SUBNET_ELEMENT_DATA_V5 RPC structure.
//
// The DHCP_SUBNET_ELEMENT_DATA_V5 structure defines the element IPv4 reservation, IPv4
// exclusion range, or IPv4 range for the subnet. This structure is an extension of
// the DHCP_SUBNET_ELEMENT_DATA_V4 (section 2.2.1.2.35) structure and is used in the
// R_DhcpAddSubnetElementV5 (section 3.2.4.38) method.
type SubnetElementDataV5 struct {
	// ElementType:  This is of type DHCP_SUBNET_ELEMENT_TYPE (section 2.2.1.1.7) enumeration,
	// defining the set of possible IPv4 subnet element types. This value defines which
	// of the values is chosen from the subsequent union, the Element member.
	ElementType SubnetElementType `idl:"name:ElementType" json:"element_type"`
	// Element:  Element is a union of different types of IPv4 subnet elements. The value
	// of the union is dependent on the previous field, the ElementType member.
	Element *SubnetElementDataV5_Element `idl:"name:Element;switch_is:(((ElementType 7 <=) (5 ElementType <=) &&) 0 ElementType ?:)" json:"element"`
}

func (o *SubnetElementDataV5) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV5) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ElementType)); err != nil {
		return err
	}
	_exprElementType := uint16(0)
	if (o.ElementType <= 7) && (5 <= o.ElementType) {
		_exprElementType = uint16(0)
	} else {
		_exprElementType = uint16(o.ElementType)
	}
	_swElement := uint16(_exprElementType)
	if o.Element != nil {
		if err := o.Element.MarshalUnionNDR(ctx, w, _swElement); err != nil {
			return err
		}
	} else {
		if err := (&SubnetElementDataV5_Element{}).MarshalUnionNDR(ctx, w, _swElement); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV5) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ElementType)); err != nil {
		return err
	}
	if o.Element == nil {
		o.Element = &SubnetElementDataV5_Element{}
	}
	_exprElementType := uint16(0)
	if (o.ElementType <= 7) && (5 <= o.ElementType) {
		_exprElementType = uint16(0)
	} else {
		_exprElementType = uint16(o.ElementType)
	}
	_swElement := uint16(_exprElementType)
	if err := o.Element.UnmarshalUnionNDR(ctx, w, _swElement); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV5_Element structure represents DHCP_SUBNET_ELEMENT_DATA_V5 union anonymous member.
//
// The DHCP_SUBNET_ELEMENT_DATA_V5 structure defines the element IPv4 reservation, IPv4
// exclusion range, or IPv4 range for the subnet. This structure is an extension of
// the DHCP_SUBNET_ELEMENT_DATA_V4 (section 2.2.1.2.35) structure and is used in the
// R_DhcpAddSubnetElementV5 (section 3.2.4.38) method.
type SubnetElementDataV5_Element struct {
	// Types that are assignable to Value
	//
	// *SubnetElementDataV5_IPRange
	// *SubnetElementDataV5_SecondaryHost
	// *SubnetElementDataV5_ReservedIP
	// *SubnetElementDataV5_ExcludeIPRange
	// *SubnetElementDataV5_IPUsedCluster
	Value is_SubnetElementDataV5_Element `json:"value"`
}

func (o *SubnetElementDataV5_Element) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SubnetElementDataV5_IPRange:
		if value != nil {
			return value.IPRange
		}
	case *SubnetElementDataV5_SecondaryHost:
		if value != nil {
			return value.SecondaryHost
		}
	case *SubnetElementDataV5_ReservedIP:
		if value != nil {
			return value.ReservedIP
		}
	case *SubnetElementDataV5_ExcludeIPRange:
		if value != nil {
			return value.ExcludeIPRange
		}
	case *SubnetElementDataV5_IPUsedCluster:
		if value != nil {
			return value.IPUsedCluster
		}
	}
	return nil
}

type is_SubnetElementDataV5_Element interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SubnetElementDataV5_Element()
}

func (o *SubnetElementDataV5_Element) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SubnetElementDataV5_IPRange:
		return uint16(0)
	case *SubnetElementDataV5_SecondaryHost:
		return uint16(1)
	case *SubnetElementDataV5_ReservedIP:
		return uint16(2)
	case *SubnetElementDataV5_ExcludeIPRange:
		return uint16(3)
	case *SubnetElementDataV5_IPUsedCluster:
		return uint16(4)
	}
	return uint16(0)
}

func (o *SubnetElementDataV5_Element) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*SubnetElementDataV5_IPRange)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV5_IPRange{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*SubnetElementDataV5_SecondaryHost)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV5_SecondaryHost{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*SubnetElementDataV5_ReservedIP)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV5_ReservedIP{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*SubnetElementDataV5_ExcludeIPRange)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV5_ExcludeIPRange{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*SubnetElementDataV5_IPUsedCluster)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV5_IPUsedCluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SubnetElementDataV5_Element) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &SubnetElementDataV5_IPRange{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &SubnetElementDataV5_SecondaryHost{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &SubnetElementDataV5_ReservedIP{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &SubnetElementDataV5_ExcludeIPRange{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &SubnetElementDataV5_IPUsedCluster{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SubnetElementDataV5_IPRange structure represents SubnetElementDataV5_Element RPC union arm.
//
// It has following labels: 0
type SubnetElementDataV5_IPRange struct {
	// IpRange:  This is of type DHCP_BOOTP_IP_RANGE (section 2.2.1.2.37) structure, containing
	// the IPv4 range for the IPv4 subnet. This contains the range for the following valid
	// enumeration values.
	//
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	|                          |                                                                                  |
	//	| DHCP SUBNET ELEMENT TYPE |                                     MEANING                                      |
	//	|                          |                                                                                  |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRanges 0           | The configuration parameter is the IP range of a DHCPv4 scope configured on the  |
	//	|                          | DHCPv4 server.                                                                   |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRangesDhcpOnly 5   | The configuration parameter is an IP range of a DHCPv4 scope configured on       |
	//	|                          | the DHCPv4 server that MUST be used only for assignment of addresses to DHCPv4   |
	//	|                          | clients on the subnet. The IP addresses from this range MUST NOT be assigned to  |
	//	|                          | BOOTP clients ([RFC2132]).                                                       |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRangesDhcpBootp 6  | The configuration parameter is an IP range of a DHCPv4 scope configured on the   |
	//	|                          | DHCPv4 server that can be used for assignment of addresses to both DHCPv4 and    |
	//	|                          | BOOTP.                                                                           |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| DhcpIpRangesBootpOnly 7  | The configuration parameter is an IP range of a DHCPv4 scope configured on the   |
	//	|                          | DHCPv4 server that MUST be used only for assignment of IPv4 addresses to BOOTP   |
	//	|                          | clients.                                                                         |
	//	+--------------------------+----------------------------------------------------------------------------------+
	IPRange *BOOTPIPRange `idl:"name:IpRange" json:"ip_range"`
}

func (*SubnetElementDataV5_IPRange) is_SubnetElementDataV5_Element() {}

func (o *SubnetElementDataV5_IPRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IPRange != nil {
		_ptr_IpRange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPRange != nil {
				if err := o.IPRange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&BOOTPIPRange{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPRange, _ptr_IpRange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV5_IPRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_IpRange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPRange == nil {
			o.IPRange = &BOOTPIPRange{}
		}
		if err := o.IPRange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IpRange := func(ptr interface{}) { o.IPRange = *ptr.(**BOOTPIPRange) }
	if err := w.ReadPointer(&o.IPRange, _s_IpRange, _ptr_IpRange); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV5_SecondaryHost structure represents SubnetElementDataV5_Element RPC union arm.
//
// It has following labels: 1
type SubnetElementDataV5_SecondaryHost struct {
	// SecondaryHost:  This is of type DHCP_HOST_INFO (section 2.2.1.2.7) structure and
	// is not used. If the ElementType value mandates that the SecondaryHost element is
	// to be used in any method, the method will return ERROR_CALL_NOT_IMPLEMENTED or ERROR_NOT_SUPPORTED,
	// as specified in the processing rules of methods that use the DHCP_SUBNET_ELEMENT_DATA_V5
	// structure.
	SecondaryHost *HostInfo `idl:"name:SecondaryHost" json:"secondary_host"`
}

func (*SubnetElementDataV5_SecondaryHost) is_SubnetElementDataV5_Element() {}

func (o *SubnetElementDataV5_SecondaryHost) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SecondaryHost != nil {
		_ptr_SecondaryHost := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SecondaryHost != nil {
				if err := o.SecondaryHost.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecondaryHost, _ptr_SecondaryHost); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV5_SecondaryHost) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_SecondaryHost := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SecondaryHost == nil {
			o.SecondaryHost = &HostInfo{}
		}
		if err := o.SecondaryHost.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SecondaryHost := func(ptr interface{}) { o.SecondaryHost = *ptr.(**HostInfo) }
	if err := w.ReadPointer(&o.SecondaryHost, _s_SecondaryHost, _ptr_SecondaryHost); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV5_ReservedIP structure represents SubnetElementDataV5_Element RPC union arm.
//
// It has following labels: 2
type SubnetElementDataV5_ReservedIP struct {
	// ReservedIp:  This is of type DHCP_IP_RESERVATION_V4 (section 2.2.1.2.32) structure,
	// containing the IPv4 reservation.
	ReservedIP *IPReservationV4 `idl:"name:ReservedIp" json:"reserved_ip"`
}

func (*SubnetElementDataV5_ReservedIP) is_SubnetElementDataV5_Element() {}

func (o *SubnetElementDataV5_ReservedIP) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ReservedIP != nil {
		_ptr_ReservedIp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ReservedIP != nil {
				if err := o.ReservedIP.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPReservationV4{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReservedIP, _ptr_ReservedIp); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV5_ReservedIP) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ReservedIp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ReservedIP == nil {
			o.ReservedIP = &IPReservationV4{}
		}
		if err := o.ReservedIP.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ReservedIp := func(ptr interface{}) { o.ReservedIP = *ptr.(**IPReservationV4) }
	if err := w.ReadPointer(&o.ReservedIP, _s_ReservedIp, _ptr_ReservedIp); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV5_ExcludeIPRange structure represents SubnetElementDataV5_Element RPC union arm.
//
// It has following labels: 3
type SubnetElementDataV5_ExcludeIPRange struct {
	// ExcludeIpRange:  This is of type DHCP_IP_RANGE, containing the IPv4 exclusion range.
	ExcludeIPRange *IPRange `idl:"name:ExcludeIpRange" json:"exclude_ip_range"`
}

func (*SubnetElementDataV5_ExcludeIPRange) is_SubnetElementDataV5_Element() {}

func (o *SubnetElementDataV5_ExcludeIPRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ExcludeIPRange != nil {
		_ptr_ExcludeIpRange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ExcludeIPRange != nil {
				if err := o.ExcludeIPRange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPRange{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ExcludeIPRange, _ptr_ExcludeIpRange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV5_ExcludeIPRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ExcludeIpRange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ExcludeIPRange == nil {
			o.ExcludeIPRange = &IPRange{}
		}
		if err := o.ExcludeIPRange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ExcludeIpRange := func(ptr interface{}) { o.ExcludeIPRange = *ptr.(**IPRange) }
	if err := w.ReadPointer(&o.ExcludeIPRange, _s_ExcludeIpRange, _ptr_ExcludeIpRange); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV5_IPUsedCluster structure represents SubnetElementDataV5_Element RPC union arm.
//
// It has following labels: 4
type SubnetElementDataV5_IPUsedCluster struct {
	// IpUsedCluster:  This is of type DHCP_IP_CLUSTER (section 2.2.1.2.88) structure and
	// is not used. If the ElementType member mandates this element to be used in any method,
	// the method will return ERROR_INVALID_PARAMETER.
	IPUsedCluster *IPCluster `idl:"name:IpUsedCluster" json:"ip_used_cluster"`
}

func (*SubnetElementDataV5_IPUsedCluster) is_SubnetElementDataV5_Element() {}

func (o *SubnetElementDataV5_IPUsedCluster) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IPUsedCluster != nil {
		_ptr_IpUsedCluster := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPUsedCluster != nil {
				if err := o.IPUsedCluster.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPCluster{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPUsedCluster, _ptr_IpUsedCluster); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV5_IPUsedCluster) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_IpUsedCluster := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPUsedCluster == nil {
			o.IPUsedCluster = &IPCluster{}
		}
		if err := o.IPUsedCluster.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IpUsedCluster := func(ptr interface{}) { o.IPUsedCluster = *ptr.(**IPCluster) }
	if err := w.ReadPointer(&o.IPUsedCluster, _s_IpUsedCluster, _ptr_IpUsedCluster); err != nil {
		return err
	}
	return nil
}

// SubnetElementInfoArrayV5 structure represents DHCP_SUBNET_ELEMENT_INFO_ARRAY_V5 RPC structure.
//
// The DHCP_SUBNET_ELEMENT_INFO_ARRAY_V5 structure defines an array of DHCP_SUBNET_ELEMENT_DATA_V5
// (section 2.2.1.2.38) structures. The first member contains the number of subnet elements
// (IPv4 reservation, IPv4 exclusion range, and IPv4 range), and the second member points
// to the array of length NumElements containing IPv4 subnet elements. This structure
// is an extension of the DHCP_SUBNET_ELEMENT_INFO_ARRAY_V4 (section 2.2.1.2.36) structure
// and is used in the R_DhcpEnumSubnetElementsV5 (section 3.2.4.39) method.
type SubnetElementInfoArrayV5 struct {
	// NumElements:  This is of type DWORD, containing the number of subnet elements in
	// the subsequent field, the Elements member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This is a pointer to an array of DHCP_SUBNET_ELEMENT_DATA_V5 structures
	// of length NumElements containing IPv4 subnet elements.
	Elements []*SubnetElementDataV5 `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *SubnetElementInfoArrayV5) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementInfoArrayV5) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SubnetElementDataV5{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SubnetElementDataV5{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementInfoArrayV5) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*SubnetElementDataV5, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &SubnetElementDataV5{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*SubnetElementDataV5) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// BindElement structure represents DHCP_BIND_ELEMENT RPC structure.
//
// The DHCP_BIND_ELEMENT structure defines an IPv4 interface binding for the DHCP server
// over which it receives DHCP packets. This structure is used in the DHCP_BIND_ELEMENT_ARRAY
// (section 2.2.1.2.81) structure.
type BindElement struct {
	// Flags:  This is of type ULONG, specifying a set of bit flags indicating properties
	// of the interface binding.
	//
	//	+-------------------------------------------+-----------------------------------+
	//	|                                           |                                   |
	//	|                   VALUE                   |              MEANING              |
	//	|                                           |                                   |
	//	+-------------------------------------------+-----------------------------------+
	//	+-------------------------------------------+-----------------------------------+
	//	| DHCP_ENDPOINT_FLAG_CANT_MODIFY 0x00000001 | The endpoints cannot be modified. |
	//	+-------------------------------------------+-----------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// fBoundToDHCPServer:  This is of type BOOL and specifies whether this binding is set
	// on the DHCP server.
	//
	//	+------------------+------------------------------------------------------------------+
	//	|                  |                                                                  |
	//	|      VALUE       |                             MEANING                              |
	//	|                  |                                                                  |
	//	+------------------+------------------------------------------------------------------+
	//	+------------------+------------------------------------------------------------------+
	//	| FALSE 0x00000000 | It specifies that the interface is not bound to the DHCP server. |
	//	+------------------+------------------------------------------------------------------+
	//	| TRUE 0x00000001  | It specifies that the interface is bound to the DHCP server.     |
	//	+------------------+------------------------------------------------------------------+
	BoundToDHCPServer bool `idl:"name:fBoundToDHCPServer" json:"bound_to_dhcp_server"`
	// AdapterPrimaryAddress:  This is of type DHCP_IP_ADDRESS, a DWORD specifying the IPv4
	// address assigned to the interface over which the DHCP server is receiving DHCP packets.
	AdapterPrimaryAddress uint32 `idl:"name:AdapterPrimaryAddress" json:"adapter_primary_address"`
	// AdapterSubnetAddress:  This is of type DHCP_IP_ADDRESS, a DWORD specifying the subnet
	// ID from which this interface is receiving DHCP packets.
	AdapterSubnetAddress uint32 `idl:"name:AdapterSubnetAddress" json:"adapter_subnet_address"`
	// IfDescription:  A pointer, of type LPWSTR, to a null-terminated Unicode string that
	// specifies the name assigned to this interface. The maximum number of characters allowed
	// in this string is 256, excluding the terminating null character.
	InterfaceDescription string `idl:"name:IfDescription" json:"interface_description"`
	// IfIdSize:  This is of type ULONG, and it contains the size of the interface GUID
	// ([MS-DTYP] section 2.3.4) stored in the IfId member.
	InterfaceIDSize uint32 `idl:"name:IfIdSize" json:"interface_id_size"`
	// IfId:  This is a pointer to a BYTE that contains the interface GUID ([MS-DTYP] section
	// 2.3.4) assigned to this interface.
	InterfaceID []byte `idl:"name:IfId;size_is:(IfIdSize)" json:"interface_id"`
}

func (o *BindElement) xxx_PreparePayload(ctx context.Context) error {
	if o.InterfaceID != nil && o.InterfaceIDSize == 0 {
		o.InterfaceIDSize = uint32(len(o.InterfaceID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindElement) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if !o.BoundToDHCPServer {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AdapterPrimaryAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.AdapterSubnetAddress); err != nil {
		return err
	}
	if o.InterfaceDescription != "" {
		_ptr_IfDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.InterfaceDescription); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.InterfaceDescription, _ptr_IfDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InterfaceIDSize); err != nil {
		return err
	}
	if o.InterfaceID != nil || o.InterfaceIDSize > 0 {
		_ptr_IfId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InterfaceIDSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.InterfaceID {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.InterfaceID[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.InterfaceID); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InterfaceID, _ptr_IfId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindElement) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	var _bBoundToDHCPServer int32
	if err := w.ReadData(&_bBoundToDHCPServer); err != nil {
		return err
	}
	o.BoundToDHCPServer = _bBoundToDHCPServer != 0
	if err := w.ReadData(&o.AdapterPrimaryAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdapterSubnetAddress); err != nil {
		return err
	}
	_ptr_IfDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.InterfaceDescription); err != nil {
			return err
		}
		return nil
	})
	_s_IfDescription := func(ptr interface{}) { o.InterfaceDescription = *ptr.(*string) }
	if err := w.ReadPointer(&o.InterfaceDescription, _s_IfDescription, _ptr_IfDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIDSize); err != nil {
		return err
	}
	_ptr_IfId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InterfaceIDSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InterfaceIDSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.InterfaceID", sizeInfo[0])
		}
		o.InterfaceID = make([]byte, sizeInfo[0])
		for i1 := range o.InterfaceID {
			i1 := i1
			if err := w.ReadData(&o.InterfaceID[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_IfId := func(ptr interface{}) { o.InterfaceID = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.InterfaceID, _s_IfId, _ptr_IfId); err != nil {
		return err
	}
	return nil
}

// BindElementArray structure represents DHCP_BIND_ELEMENT_ARRAY RPC structure.
//
// The DHCP_BIND_ELEMENT_ARRAY structure defines an array of DHCP_BIND_ELEMENT (section
// 2.2.1.2.80) structures. This contains an array of IPv4 interface bindings over which
// the DHCP server receives DHCP packets. The first member contains the number of IPv4
// interface bindings present, and the second member points to the array of interface
// bindings over which the DHCP server is receiving DHCP packets.
type BindElementArray struct {
	// NumElements:  This is of type DWORD and specifies the number of interface bindings
	// listed in the subsequent field the Elements member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This is a pointer to an array of type DHCP_BIND_ELEMENT (section 2.2.1.2.80)
	// structure and of length NumElements that contains information about the interface
	// bindings of a DHCP server.
	Elements []*BindElement `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *BindElementArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindElementArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&BindElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&BindElement{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindElementArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*BindElement, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &BindElement{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*BindElement) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// ServerSpecificStrings structure represents DHCP_SERVER_SPECIFIC_STRINGS RPC structure.
//
// The DHCP_SERVER_SPECIFIC_STRINGS structure defines the names of the default user
// class and vendor class.
type ServerSpecificStrings struct {
	// DefaultVendorClassName:  A pointer, of type LPWSTR, to a null-terminated Unicode
	// string that contains the name of the default vendor class. The maximum number of
	// characters allowed in this string is 255, which includes the terminating null character.
	DefaultVendorClassName string `idl:"name:DefaultVendorClassName" json:"default_vendor_class_name"`
	// DefaultUserClassName:  A pointer, of type LPWSTR, to a null-terminated Unicode string
	// that contains the name of the default user class.  The maximum number of characters
	// allowed in this string is 255, which includes the terminating null character.
	DefaultUserClassName string `idl:"name:DefaultUserClassName" json:"default_user_class_name"`
}

func (o *ServerSpecificStrings) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerSpecificStrings) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.DefaultVendorClassName != "" {
		_ptr_DefaultVendorClassName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DefaultVendorClassName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DefaultVendorClassName, _ptr_DefaultVendorClassName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DefaultUserClassName != "" {
		_ptr_DefaultUserClassName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DefaultUserClassName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DefaultUserClassName, _ptr_DefaultUserClassName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerSpecificStrings) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_DefaultVendorClassName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DefaultVendorClassName); err != nil {
			return err
		}
		return nil
	})
	_s_DefaultVendorClassName := func(ptr interface{}) { o.DefaultVendorClassName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DefaultVendorClassName, _s_DefaultVendorClassName, _ptr_DefaultVendorClassName); err != nil {
		return err
	}
	_ptr_DefaultUserClassName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DefaultUserClassName); err != nil {
			return err
		}
		return nil
	})
	_s_DefaultUserClassName := func(ptr interface{}) { o.DefaultUserClassName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DefaultUserClassName, _s_DefaultUserClassName, _ptr_DefaultUserClassName); err != nil {
		return err
	}
	return nil
}

// ScopeMIBInfoV5 structure represents SCOPE_MIB_INFO_V5 RPC structure.
//
// The SCOPE_MIB_INFO_V5 structure defines a structure that contains the address counters
// for a specific IPv4 subnet. The numbers of free, used, and offered IPv4 addresses
// are stored in this structure. This structure is used in the DHCP_MIB_INFO_V5 (section
// 2.2.1.2.95) structure.
type ScopeMIBInfoV5 struct {
	// Subnet:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), a DWORD specifying
	// the IPv4 subnet ID for the scope.
	Subnet uint32 `idl:"name:Subnet" json:"subnet"`
	// NumAddressesInuse:  This is of type DWORD, containing the number of IPv4 addresses
	// leased out to DHCPv4 clients for a given IPv4 subnet.
	AddressesInUseLength uint32 `idl:"name:NumAddressesInuse" json:"addresses_inuse_length"`
	// NumAddressesFree:  This is of type DWORD, containing the number of IPv4 addresses
	// that are free and can be leased out to DHCPv4 clients in a given IPv4 subnet.
	AddressesFreeLength uint32 `idl:"name:NumAddressesFree" json:"addresses_free_length"`
	// NumPendingOffers:  This is of type DWORD, containing the number of IPv4 addresses
	// that have been offered to DHCPv4 clients in a given IPv4 subnet but that the DHCP
	// client has not yet confirmed.
	PendingOffersLength uint32 `idl:"name:NumPendingOffers" json:"pending_offers_length"`
}

func (o *ScopeMIBInfoV5) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ScopeMIBInfoV5) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Subnet); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressesInUseLength); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressesFreeLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PendingOffersLength); err != nil {
		return err
	}
	return nil
}
func (o *ScopeMIBInfoV5) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Subnet); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressesInUseLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressesFreeLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PendingOffersLength); err != nil {
		return err
	}
	return nil
}

// MIBInfoV5 structure represents DHCP_MIB_INFO_V5 RPC structure.
//
// The DHCP_MIB_INFO_V5 structure contains the counter values for the DHCP Server service.
type MIBInfoV5 struct {
	// Discovers:  This member is a type DWORD that contains the number of DHCPDISCOVER
	// messages [RFC2131] received by the DHCP server from the DHCP clients since the DHCP
	// server was last started. This is used for statistical analysis by the DHCP server.
	Discovers uint32 `idl:"name:Discovers" json:"discovers"`
	// Offers:  This member is a type DWORD that contains the number of DHCPOFFER messages
	// sent by the DHCP server to the DHCP clients for which the DHCP server has not confirmed
	// since the DHCP server was last started. This is used for statistical analysis by
	// the DHCP server.
	Offers uint32 `idl:"name:Offers" json:"offers"`
	// Requests:  This member is a type DWORD that contains the number of DHCPREQUEST messages
	// received by the DHCP server from the DHCP clients since the DHCP server was last
	// started. This is used for statistical analysis by the DHCP server.
	Requests uint32 `idl:"name:Requests" json:"requests"`
	// Acks:  This member is a type DWORD that contains the number of DHCPACK messages sent
	// by the DHCP Server to the DHCP clients since the DHCP server was last started. This
	// is used for statistical analysis by the DHCP server.
	ACKs uint32 `idl:"name:Acks" json:"acks"`
	// Naks:  This member is a type DWORD that contains the number of DHCPNAK messages sent
	// by the DHCP server to DHCP clients since the DHCP server was last started. This is
	// used for statistical analysis by the DHCP server.
	NAKs uint32 `idl:"name:Naks" json:"naks"`
	// Declines:  This member is a type DWORD that contains the number of DHCPDECLINE messages
	// received by the DHCP server from the DHCP client since the DHCP server was last started.
	// This is used for statistical analysis by the DHCP server.
	Declines uint32 `idl:"name:Declines" json:"declines"`
	// Releases:  This member is a type DWORD that contains the number of DHCPRELEASE messages
	// received by the DHCP server from the DHCP client since the DHCP server was last started.
	// This is used for statistical analysis by the DHCP server.
	Releases uint32 `idl:"name:Releases" json:"releases"`
	// ServerStartTime:  This member is a type DATE_TIME (section 2.2.1.2.11) structure
	// that contains the start time of the DHCP server.
	ServerStartTime *DateTime `idl:"name:ServerStartTime" json:"server_start_time"`
	// QtnNumLeases:   This member is a type DWORD that MUST be set to zero when sent and
	// ignored on receipt. This is treated as an error if it is nonzero in an RPC method
	// which queries DHCP server configuration.
	//
	//	+------------+---------+
	//	|            |         |
	//	|   VALUE    | MEANING |
	//	|            |         |
	//	+------------+---------+
	//	+------------+---------+
	//	| 0x00000000 | Sending |
	//	+------------+---------+
	QTNNumLeases uint32 `idl:"name:QtnNumLeases" json:"qtn_num_leases"`
	// QtnPctQtnLeases:   This member is a type DWORD that MUST be set to zero when sent
	// and ignored on receipt. This is treated as an error if it is nonzero in an RPC method
	// which queries DHCP server configuration.
	//
	//	+------------+---------+
	//	|            |         |
	//	|   VALUE    | MEANING |
	//	|            |         |
	//	+------------+---------+
	//	+------------+---------+
	//	| 0x00000000 | Sending |
	//	+------------+---------+
	QTNPctQTNLeases uint32 `idl:"name:QtnPctQtnLeases" json:"qtn_pct_qtn_leases"`
	// QtnProbationLeases:  This member is a type DWORD that MUST be set to zero when sent
	// and ignored on receipt. This is treated as an error if it is nonzero in an RPC method
	// which queries DHCP server configuration.
	//
	//	+------------+---------+
	//	|            |         |
	//	|   VALUE    | MEANING |
	//	|            |         |
	//	+------------+---------+
	//	+------------+---------+
	//	| 0x00000000 | Sending |
	//	+------------+---------+
	QTNProbationLeases uint32 `idl:"name:QtnProbationLeases" json:"qtn_probation_leases"`
	// QtnNonQtnLeases:   This member is a type DWORD that MUST be set to zero when sent
	// and ignored on receipt. This is treated as an error if it is nonzero in an RPC method
	// which queries DHCP server configuration.
	//
	//	+------------+---------+
	//	|            |         |
	//	|   VALUE    | MEANING |
	//	|            |         |
	//	+------------+---------+
	//	+------------+---------+
	//	| 0x00000000 | Sending |
	//	+------------+---------+
	QTNNonQTNLeases uint32 `idl:"name:QtnNonQtnLeases" json:"qtn_non_qtn_leases"`
	// QtnExemptLeases:   This member is a type DWORD that MUST be set to zero when sent
	// and ignored on receipt. This is treated as an error if it is nonzero in an RPC method
	// which queries DHCP server configuration.
	//
	//	+------------+---------+
	//	|            |         |
	//	|   VALUE    | MEANING |
	//	|            |         |
	//	+------------+---------+
	//	+------------+---------+
	//	| 0x00000000 | Sending |
	//	+------------+---------+
	QTNExemptLeases uint32 `idl:"name:QtnExemptLeases" json:"qtn_exempt_leases"`
	// QtnCapableClients:  This member is a type DWORD that MUST be set to zero when sent
	// and ignored on receipt. This is treated as an error if it is nonzero in an RPC method
	// which queries DHCP server configuration.
	//
	//	+------------+---------+
	//	|            |         |
	//	|   VALUE    | MEANING |
	//	|            |         |
	//	+------------+---------+
	//	+------------+---------+
	//	| 0x00000000 | Sending |
	//	+------------+---------+
	QTNCapableClients uint32 `idl:"name:QtnCapableClients" json:"qtn_capable_clients"`
	// QtnIASErrors:  This member is a type DWORD that MUST be set to zero when sent and
	// ignored on receipt. This is treated as an error if it is nonzero in an RPC method
	// which queries DHCP server configuration.
	//
	//	+------------+---------+
	//	|            |         |
	//	|   VALUE    | MEANING |
	//	|            |         |
	//	+------------+---------+
	//	+------------+---------+
	//	| 0x00000000 | Sending |
	//	+------------+---------+
	QTNIASErrors uint32 `idl:"name:QtnIASErrors" json:"qtn_ias_errors"`
	// DelayedOffers:  This member is a type DWORD that contains the number of delayed DHCPOFFER
	// messages sent by the DHCP server to the DHCP clients. This is used for statistical
	// analysis by the DHCP server.
	DelayedOffers uint32 `idl:"name:DelayedOffers" json:"delayed_offers"`
	// ScopesWithDelayedOffers:  This member is a type DWORD that contains the number of
	// scopes which are configured with subnet delay for DHCPOFFER messages.
	ScopesWithDelayedOffers uint32 `idl:"name:ScopesWithDelayedOffers" json:"scopes_with_delayed_offers"`
	// Scopes:  This member is a type DWORD which contains the number of IPv4 scopes configured
	// on the current DHCP server. This is used for statistical analysis by the DHCP server.
	// This field defines the number of DHCP scopes in the subsequent field, the ScopeInfo
	// member.
	Scopes uint32 `idl:"name:Scopes" json:"scopes"`
	// ScopeInfo:  This member is a pointer to an array of SCOPE_MIB_INFO (section 2.2.1.2.47)
	// structures of length Scopes that contains the information about the IPv4 scopes configured
	// on the DHCP server.
	ScopeInfo []*ScopeMIBInfoV5 `idl:"name:ScopeInfo;size_is:(Scopes)" json:"scope_info"`
}

func (o *MIBInfoV5) xxx_PreparePayload(ctx context.Context) error {
	if o.ScopeInfo != nil && o.Scopes == 0 {
		o.Scopes = uint32(len(o.ScopeInfo))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBInfoV5) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Discovers); err != nil {
		return err
	}
	if err := w.WriteData(o.Offers); err != nil {
		return err
	}
	if err := w.WriteData(o.Requests); err != nil {
		return err
	}
	if err := w.WriteData(o.ACKs); err != nil {
		return err
	}
	if err := w.WriteData(o.NAKs); err != nil {
		return err
	}
	if err := w.WriteData(o.Declines); err != nil {
		return err
	}
	if err := w.WriteData(o.Releases); err != nil {
		return err
	}
	if o.ServerStartTime != nil {
		if err := o.ServerStartTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.QTNNumLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNPctQTNLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNProbationLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNNonQTNLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNExemptLeases); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNCapableClients); err != nil {
		return err
	}
	if err := w.WriteData(o.QTNIASErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.DelayedOffers); err != nil {
		return err
	}
	if err := w.WriteData(o.ScopesWithDelayedOffers); err != nil {
		return err
	}
	if err := w.WriteData(o.Scopes); err != nil {
		return err
	}
	if o.ScopeInfo != nil || o.Scopes > 0 {
		_ptr_ScopeInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Scopes)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ScopeInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ScopeInfo[i1] != nil {
					if err := o.ScopeInfo[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ScopeMIBInfoV5{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ScopeInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ScopeMIBInfoV5{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeInfo, _ptr_ScopeInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBInfoV5) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Discovers); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offers); err != nil {
		return err
	}
	if err := w.ReadData(&o.Requests); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACKs); err != nil {
		return err
	}
	if err := w.ReadData(&o.NAKs); err != nil {
		return err
	}
	if err := w.ReadData(&o.Declines); err != nil {
		return err
	}
	if err := w.ReadData(&o.Releases); err != nil {
		return err
	}
	if o.ServerStartTime == nil {
		o.ServerStartTime = &DateTime{}
	}
	if err := o.ServerStartTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNNumLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNPctQTNLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNProbationLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNNonQTNLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNExemptLeases); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNCapableClients); err != nil {
		return err
	}
	if err := w.ReadData(&o.QTNIASErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.DelayedOffers); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScopesWithDelayedOffers); err != nil {
		return err
	}
	if err := w.ReadData(&o.Scopes); err != nil {
		return err
	}
	_ptr_ScopeInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Scopes > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Scopes)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ScopeInfo", sizeInfo[0])
		}
		o.ScopeInfo = make([]*ScopeMIBInfoV5, sizeInfo[0])
		for i1 := range o.ScopeInfo {
			i1 := i1
			if o.ScopeInfo[i1] == nil {
				o.ScopeInfo[i1] = &ScopeMIBInfoV5{}
			}
			if err := o.ScopeInfo[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ScopeInfo := func(ptr interface{}) { o.ScopeInfo = *ptr.(*[]*ScopeMIBInfoV5) }
	if err := w.ReadPointer(&o.ScopeInfo, _s_ScopeInfo, _ptr_ScopeInfo); err != nil {
		return err
	}
	return nil
}

// FilterListType type represents DHCP_FILTER_LIST_TYPE RPC enumeration.
//
// The DHCP_FILTER_LIST_TYPE enumeration defines the type of filter list to which the
// link-layer filter is to be added.
type FilterListType uint16

var (
	// Deny: Add the link-layer filter to the deny list.
	FilterListTypeDeny FilterListType = 0
	// Allow: Add the link-layer filter to the allow list.
	FilterListTypeAllow FilterListType = 1
)

func (o FilterListType) String() string {
	switch o {
	case FilterListTypeDeny:
		return "FilterListTypeDeny"
	case FilterListTypeAllow:
		return "FilterListTypeAllow"
	}
	return "Invalid"
}

// AddrPattern structure represents DHCP_ADDR_PATTERN RPC structure.
//
// The DHCP_ADDR_PATTERN structure contains the information regarding the link-layer
// address/pattern.
type AddrPattern struct {
	// MatchHWType:  This is of type BOOL. Setting the field to FALSE will cause the filtering
	// to disregard the hardware type field and a TRUE value will match the hardware type
	// field.
	MatchHwType bool `idl:"name:MatchHWType" json:"match_hw_type"`
	// HWType:  This is of type BYTE and specifies the hardware type of the address, specified
	// in the pattern. For the list of valid options, see [RFC1700].
	HwType uint8 `idl:"name:HWType" json:"hw_type"`
	// IsWildcard:  This is of type BOOL and specifies whether the current pattern represents
	// a wildcard pattern.
	//
	//	+---------+------------------------------------+
	//	|         |                                    |
	//	|  BOOL   |              MEANING               |
	//	|         |                                    |
	//	+---------+------------------------------------+
	//	+---------+------------------------------------+
	//	| TRUE 1  | The pattern is a wildcard pattern. |
	//	+---------+------------------------------------+
	//	| FALSE 0 | The pattern is a hardware address. |
	//	+---------+------------------------------------+
	IsWildcard bool `idl:"name:IsWildcard" json:"is_wildcard"`
	// Length:  This is of type BYTE and specifies the length of the pattern.
	Length uint8 `idl:"name:Length" json:"length"`
	// Pattern:  This is a pointer to a type BYTE and contains the address/pattern.
	Pattern []byte `idl:"name:Pattern" json:"pattern"`
}

func (o *AddrPattern) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AddrPattern) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.MatchHwType {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.HwType); err != nil {
		return err
	}
	if !o.IsWildcard {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	for i1 := range o.Pattern {
		i1 := i1
		if uint64(i1) >= 255 {
			break
		}
		if err := w.WriteData(o.Pattern[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Pattern); uint64(i1) < 255; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *AddrPattern) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bMatchHwType int32
	if err := w.ReadData(&_bMatchHwType); err != nil {
		return err
	}
	o.MatchHwType = _bMatchHwType != 0
	if err := w.ReadData(&o.HwType); err != nil {
		return err
	}
	var _bIsWildcard int32
	if err := w.ReadData(&_bIsWildcard); err != nil {
		return err
	}
	o.IsWildcard = _bIsWildcard != 0
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	o.Pattern = make([]byte, 255)
	for i1 := range o.Pattern {
		i1 := i1
		if err := w.ReadData(&o.Pattern[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// FilterAddInfo structure represents DHCP_FILTER_ADD_INFO RPC structure.
//
// The DHCP_FILTER_ADD_INFO structure contains information regarding the link-layer
// filter to be added to the allow and deny filter list.
type FilterAddInfo struct {
	// AddrPatt:  This is of type DHCP_ADDR_PATTERN (section 2.2.1.2.89) structure and contains
	// the address/pattern-related information of the link-layer filter.
	AddrPattern *AddrPattern `idl:"name:AddrPatt" json:"addr_pattern"`
	// Comment:  This is a pointer, of type LPWSTR, to a null-terminated Unicode string
	// that contains the comment associated with the address/pattern. The maximum number
	// of characters allowed in this string is 128, which includes the terminating null
	// character.
	Comment string `idl:"name:Comment" json:"comment"`
	// ListType:  This is of type DHCP_FILTER_LIST_TYPE (section 2.2.1.1.17) enumeration,
	// which specifies the list type to which the filter is to be added.
	ListType FilterListType `idl:"name:ListType" json:"list_type"`
}

func (o *FilterAddInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FilterAddInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.AddrPattern != nil {
		if err := o.AddrPattern.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AddrPattern{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.ListType)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *FilterAddInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.AddrPattern == nil {
		o.AddrPattern = &AddrPattern{}
	}
	if err := o.AddrPattern.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ListType)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// FilterGlobalInfo structure represents DHCP_FILTER_GLOBAL_INFO RPC structure.
//
// The DHCP_FILTER_GLOBAL_INFO structure contains information regarding enabling/disabling
// the allow and deny filter lists.
type FilterGlobalInfo struct {
	// EnforceAllowList:  This is of type BOOL and specifies whether the allow list is enabled
	// or disabled.
	//
	//	+---------+-----------------------------+
	//	|         |                             |
	//	|  BOOL   |           MEANING           |
	//	|         |                             |
	//	+---------+-----------------------------+
	//	+---------+-----------------------------+
	//	| TRUE 1  | The allow list is enabled.  |
	//	+---------+-----------------------------+
	//	| FALSE 0 | The allow list is disabled. |
	//	+---------+-----------------------------+
	EnforceAllowList bool `idl:"name:EnforceAllowList" json:"enforce_allow_list"`
	// EnforceDenyList:  This is of type BOOL and specifies whether the deny list is enabled
	// or disabled.
	//
	//	+---------+----------------------------+
	//	|         |                            |
	//	|  BOOL   |          MEANING           |
	//	|         |                            |
	//	+---------+----------------------------+
	//	+---------+----------------------------+
	//	| TRUE 1  | The deny list is enabled.  |
	//	+---------+----------------------------+
	//	| FALSE 0 | The deny list is disabled. |
	//	+---------+----------------------------+
	EnforceDenyList bool `idl:"name:EnforceDenyList" json:"enforce_deny_list"`
}

func (o *FilterGlobalInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FilterGlobalInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.EnforceAllowList {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.EnforceDenyList {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *FilterGlobalInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bEnforceAllowList int32
	if err := w.ReadData(&_bEnforceAllowList); err != nil {
		return err
	}
	o.EnforceAllowList = _bEnforceAllowList != 0
	var _bEnforceDenyList int32
	if err := w.ReadData(&_bEnforceDenyList); err != nil {
		return err
	}
	o.EnforceDenyList = _bEnforceDenyList != 0
	return nil
}

// FilterRecord structure represents DHCP_FILTER_RECORD RPC structure.
//
// The DHCP_FILTER_RECORD structure contains information regarding a link-layer filter
// record.
type FilterRecord struct {
	// AddrPatt:  This is of type DHCP_ADDR_PATTERN (section 2.2.1.2.89) structure and contains
	// the address/pattern related information of the link-layer filter.
	AddrPattern *AddrPattern `idl:"name:AddrPatt" json:"addr_pattern"`
	// Comment:  This is a pointer, of type LPWSTR, to a null-terminated Unicode string
	// that contains the comment associated with the address/pattern. The maximum number
	// of characters allowed in this string is 128, which includes the terminating null
	// character.
	Comment string `idl:"name:Comment" json:"comment"`
}

func (o *FilterRecord) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FilterRecord) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.AddrPattern != nil {
		if err := o.AddrPattern.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AddrPattern{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *FilterRecord) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.AddrPattern == nil {
		o.AddrPattern = &AddrPattern{}
	}
	if err := o.AddrPattern.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	return nil
}

// FilterEnumInfo structure represents DHCP_FILTER_ENUM_INFO RPC structure.
//
// The DHCP_FILTER_ENUM_INFO structure contains information regarding the number of
// link-layer filter records.
type FilterEnumInfo struct {
	// NumElements:  This is of type DWORD, which specifies the number of link-layer filter
	// records contained in the array specified by the pEnumRecords member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// pEnumRecords:  This is a pointer to an array of DHCP_FILTER_RECORD (section 2.2.1.2.92)
	// structures that contains link-layer filter records.
	EnumRecords []*FilterRecord `idl:"name:pEnumRecords;size_is:(NumElements)" json:"enum_records"`
}

func (o *FilterEnumInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.EnumRecords != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.EnumRecords))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FilterEnumInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.EnumRecords != nil || o.ElementsLength > 0 {
		_ptr_pEnumRecords := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.EnumRecords {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.EnumRecords[i1] != nil {
					if err := o.EnumRecords[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&FilterRecord{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.EnumRecords); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&FilterRecord{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EnumRecords, _ptr_pEnumRecords); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *FilterEnumInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_pEnumRecords := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.EnumRecords", sizeInfo[0])
		}
		o.EnumRecords = make([]*FilterRecord, sizeInfo[0])
		for i1 := range o.EnumRecords {
			i1 := i1
			if o.EnumRecords[i1] == nil {
				o.EnumRecords[i1] = &FilterRecord{}
			}
			if err := o.EnumRecords[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pEnumRecords := func(ptr interface{}) { o.EnumRecords = *ptr.(*[]*FilterRecord) }
	if err := w.ReadPointer(&o.EnumRecords, _s_pEnumRecords, _ptr_pEnumRecords); err != nil {
		return err
	}
	return nil
}

// SubnetInfoV6 structure represents DHCP_SUBNET_INFO_V6 RPC structure.
//
// The DHCP_SUBNET_INFO_V6 structure contains information about an IPv6 subnet. This
// structure is used in the R_DhcpCreateSubnetV6 (section 3.2.4.58) method.
type SubnetInfoV6 struct {
	// SubnetAddress:   This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28) structure,
	// specifying the IPv6 prefix.
	SubnetAddress *IPv6Address `idl:"name:SubnetAddress" json:"subnet_address"`
	// Prefix:  This is of type ULONG, specifying the prefix length of the IPv6 prefix.
	Prefix uint32 `idl:"name:Prefix" json:"prefix"`
	// Preference:  This is of type USHORT, specifying the preference for the IPv6 prefix
	// specified by SubnetAddress.
	Preference uint16 `idl:"name:Preference" json:"preference"`
	// SubnetName:  A pointer, of type LPWSTR, to a null-terminated Unicode string that
	// contains the name of the IPv6 prefix. There is no restriction on the length of this
	// Unicode string.
	SubnetName string `idl:"name:SubnetName" json:"subnet_name"`
	// SubnetComment:  A pointer, of type LPWSTR, to a null-terminated Unicode string that
	// contains an optional comment for the IPv6 prefix. There is no restriction on the
	// length of this Unicode string.
	SubnetComment string `idl:"name:SubnetComment" json:"subnet_comment"`
	// State:  This is of type DHCP_SUBNET_STATE (section 2.2.1.1.2) enumeration that indicates
	// the current state of the IPv6 prefix.
	State uint32 `idl:"name:State" json:"state"`
	// ScopeId:  This is of type DWORD and is the unique identifier for that IPv6 prefix.
	// This value is generated by the DHCPv6 server.
	ScopeID uint32 `idl:"name:ScopeId" json:"scope_id"`
}

func (o *SubnetInfoV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetInfoV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.SubnetAddress != nil {
		if err := o.SubnetAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Prefix); err != nil {
		return err
	}
	if err := w.WriteData(o.Preference); err != nil {
		return err
	}
	if o.SubnetName != "" {
		_ptr_SubnetName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SubnetName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SubnetName, _ptr_SubnetName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SubnetComment != "" {
		_ptr_SubnetComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SubnetComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SubnetComment, _ptr_SubnetComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.ScopeID); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *SubnetInfoV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.SubnetAddress == nil {
		o.SubnetAddress = &IPv6Address{}
	}
	if err := o.SubnetAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Prefix); err != nil {
		return err
	}
	if err := w.ReadData(&o.Preference); err != nil {
		return err
	}
	_ptr_SubnetName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SubnetName); err != nil {
			return err
		}
		return nil
	})
	_s_SubnetName := func(ptr interface{}) { o.SubnetName = *ptr.(*string) }
	if err := w.ReadPointer(&o.SubnetName, _s_SubnetName, _ptr_SubnetName); err != nil {
		return err
	}
	_ptr_SubnetComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SubnetComment); err != nil {
			return err
		}
		return nil
	})
	_s_SubnetComment := func(ptr interface{}) { o.SubnetComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.SubnetComment, _s_SubnetComment, _ptr_SubnetComment); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScopeID); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// IPArrayV6 structure represents DHCPV6_IP_ARRAY RPC structure.
//
// The DHCPV6_IP_ARRAY structure defines an array of type DHCP_IPV6_ADDRESS (section
// 2.2.1.2.28) structure that contains IPv6 prefixes. This is used in the  R_Dhcp_EnumSubnetsV6
// (section 3.2.4.59) method.
type IPArrayV6 struct {
	// NumElements:  This is of type DWORD, containing the number of IPv6 addresses in the
	// subsequent field the Elements member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This is a pointer to an array of DHCP_IPV6_ADDRESS structures of length
	// NumElements containing IPv6 addresses of the prefixes.
	Elements []*IPv6Address `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *IPArrayV6) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPArrayV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPArrayV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*IPv6Address, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &IPv6Address{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*IPv6Address) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// IPRangeV6 structure represents DHCP_IP_RANGE_V6 RPC structure.
//
// The DHCP_IP_RANGE_V6 structure defines the IPv6 range for an IPv6 subnet. This is
// used in the DHCP_SUBNET_ELEMENT_DATA_V6 (section 2.2.1.2.60) structure.
type IPRangeV6 struct {
	// StartAddress:   This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28) structure,
	// containing the first IPv6 address in the IPv6 range.
	StartAddress *IPv6Address `idl:"name:StartAddress" json:"start_address"`
	// EndAddress:  This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28) structure, containing
	// the last IPv6 address in the IPv6 range.
	EndAddress *IPv6Address `idl:"name:EndAddress" json:"end_address"`
}

func (o *IPRangeV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPRangeV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.StartAddress != nil {
		if err := o.StartAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.EndAddress != nil {
		if err := o.EndAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPRangeV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.StartAddress == nil {
		o.StartAddress = &IPv6Address{}
	}
	if err := o.StartAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.EndAddress == nil {
		o.EndAddress = &IPv6Address{}
	}
	if err := o.EndAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IPReservationV6 structure represents DHCP_IP_RESERVATION_V6 RPC structure.
//
// The DHCP_IP_RESERVATION_V6 structure defines an IPv6 reservation for a DHCPv6 client
// in a specific IPv6 prefix. This structure is used in the DHCP_SUBNET_ELEMENT_DATA_V6
// (section 2.2.1.2.60) structure.
type IPReservationV6 struct {
	// ReservedIpAddress:   This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28) structure
	// that contains the IPv6 address of the DHCPv6 client for which an IPv6 reservation
	// is created.
	ReservedIPAddress *IPv6Address `idl:"name:ReservedIpAddress" json:"reserved_ip_address"`
	// ReservedForClient:  This is a pointer of type DHCP_CLIENT_UID (section 2.2.1.2.5)
	// structure that represents the DHCPv6 client-identifier (section 2.2.1.2.5.3).
	ReservedForClient *ClientUID `idl:"name:ReservedForClient" json:"reserved_for_client"`
	// InterfaceId:   This is of type DWORD that specifies the interface identifier for
	// which the IPv6 reservation is created.
	InterfaceID uint32 `idl:"name:InterfaceId" json:"interface_id"`
}

func (o *IPReservationV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPReservationV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ReservedIPAddress != nil {
		if err := o.ReservedIPAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ReservedForClient != nil {
		_ptr_ReservedForClient := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ReservedForClient != nil {
				if err := o.ReservedForClient.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReservedForClient, _ptr_ReservedForClient); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InterfaceID); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *IPReservationV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ReservedIPAddress == nil {
		o.ReservedIPAddress = &IPv6Address{}
	}
	if err := o.ReservedIPAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ReservedForClient := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ReservedForClient == nil {
			o.ReservedForClient = &ClientUID{}
		}
		if err := o.ReservedForClient.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ReservedForClient := func(ptr interface{}) { o.ReservedForClient = *ptr.(**ClientUID) }
	if err := w.ReadPointer(&o.ReservedForClient, _s_ReservedForClient, _ptr_ReservedForClient); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceID); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV6 structure represents DHCP_SUBNET_ELEMENT_DATA_V6 RPC structure.
//
// The DHCP_SUBNET_ELEMENT_DATA_V6 structure defines the elements of the IPv6 prefix,
// such as IPv6 reservation, IPv6 exclusion range, or IPv6 range. This is used in the
// R_DhcpAddSubnetElementV6 (section 3.2.4.60) method.
type SubnetElementDataV6 struct {
	// ElementType:  ElementType is of type DHCP_SUBNET_ELEMENT_TYPE_V6 (section 2.2.1.1.8)
	// structure defining the set of possible prefix element types. This value defines which
	// of the values is chosen from the subsequent union the Element member.
	ElementType SubnetElementTypeV6 `idl:"name:ElementType" json:"element_type"`
	// Element:  Element is a union of different types of IPv6 prefix elements. The value
	// of the union is dependent on the previous field the ElementType member.
	Element *SubnetElementDataV6_Element `idl:"name:Element;switch_is:(((ElementType 7 <=) (5 ElementType <=) &&) 0 ElementType ?:)" json:"element"`
}

func (o *SubnetElementDataV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ElementType)); err != nil {
		return err
	}
	_exprElementType := uint16(0)
	if (o.ElementType <= 7) && (5 <= o.ElementType) {
		_exprElementType = uint16(0)
	} else {
		_exprElementType = uint16(o.ElementType)
	}
	_swElement := uint16(_exprElementType)
	if o.Element != nil {
		if err := o.Element.MarshalUnionNDR(ctx, w, _swElement); err != nil {
			return err
		}
	} else {
		if err := (&SubnetElementDataV6_Element{}).MarshalUnionNDR(ctx, w, _swElement); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ElementType)); err != nil {
		return err
	}
	if o.Element == nil {
		o.Element = &SubnetElementDataV6_Element{}
	}
	_exprElementType := uint16(0)
	if (o.ElementType <= 7) && (5 <= o.ElementType) {
		_exprElementType = uint16(0)
	} else {
		_exprElementType = uint16(o.ElementType)
	}
	_swElement := uint16(_exprElementType)
	if err := o.Element.UnmarshalUnionNDR(ctx, w, _swElement); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV6_Element structure represents DHCP_SUBNET_ELEMENT_DATA_V6 union anonymous member.
//
// The DHCP_SUBNET_ELEMENT_DATA_V6 structure defines the elements of the IPv6 prefix,
// such as IPv6 reservation, IPv6 exclusion range, or IPv6 range. This is used in the
// R_DhcpAddSubnetElementV6 (section 3.2.4.60) method.
type SubnetElementDataV6_Element struct {
	// Types that are assignable to Value
	//
	// *SubnetElementDataV6_IPRange
	// *SubnetElementDataV6_ReservedIP
	// *SubnetElementDataV6_ExcludeIPRange
	Value is_SubnetElementDataV6_Element `json:"value"`
}

func (o *SubnetElementDataV6_Element) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SubnetElementDataV6_IPRange:
		if value != nil {
			return value.IPRange
		}
	case *SubnetElementDataV6_ReservedIP:
		if value != nil {
			return value.ReservedIP
		}
	case *SubnetElementDataV6_ExcludeIPRange:
		if value != nil {
			return value.ExcludeIPRange
		}
	}
	return nil
}

type is_SubnetElementDataV6_Element interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SubnetElementDataV6_Element()
}

func (o *SubnetElementDataV6_Element) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SubnetElementDataV6_IPRange:
		return uint16(0)
	case *SubnetElementDataV6_ReservedIP:
		return uint16(1)
	case *SubnetElementDataV6_ExcludeIPRange:
		return uint16(2)
	}
	return uint16(0)
}

func (o *SubnetElementDataV6_Element) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*SubnetElementDataV6_IPRange)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV6_IPRange{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*SubnetElementDataV6_ReservedIP)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV6_ReservedIP{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*SubnetElementDataV6_ExcludeIPRange)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SubnetElementDataV6_ExcludeIPRange{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SubnetElementDataV6_Element) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &SubnetElementDataV6_IPRange{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &SubnetElementDataV6_ReservedIP{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &SubnetElementDataV6_ExcludeIPRange{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SubnetElementDataV6_IPRange structure represents SubnetElementDataV6_Element RPC union arm.
//
// It has following labels: 0
type SubnetElementDataV6_IPRange struct {
	// IpRange:  This is a pointer to a DHCP_IP_RANGE_V6 (section 2.2.1.2.59) structure
	// that contains the IPv6 range for this IPv6 prefix.
	IPRange *IPRangeV6 `idl:"name:IpRange" json:"ip_range"`
}

func (*SubnetElementDataV6_IPRange) is_SubnetElementDataV6_Element() {}

func (o *SubnetElementDataV6_IPRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IPRange != nil {
		_ptr_IpRange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IPRange != nil {
				if err := o.IPRange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPRangeV6{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IPRange, _ptr_IpRange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV6_IPRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_IpRange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IPRange == nil {
			o.IPRange = &IPRangeV6{}
		}
		if err := o.IPRange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_IpRange := func(ptr interface{}) { o.IPRange = *ptr.(**IPRangeV6) }
	if err := w.ReadPointer(&o.IPRange, _s_IpRange, _ptr_IpRange); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV6_ReservedIP structure represents SubnetElementDataV6_Element RPC union arm.
//
// It has following labels: 1
type SubnetElementDataV6_ReservedIP struct {
	// ReservedIp:  This is a pointer to a DHCP_IP_RESERVATION_V6 (section 2.2.1.2.58) structure
	// that contains the information on IPv6 reservations.
	ReservedIP *IPReservationV6 `idl:"name:ReservedIp" json:"reserved_ip"`
}

func (*SubnetElementDataV6_ReservedIP) is_SubnetElementDataV6_Element() {}

func (o *SubnetElementDataV6_ReservedIP) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ReservedIP != nil {
		_ptr_ReservedIp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ReservedIP != nil {
				if err := o.ReservedIP.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPReservationV6{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ReservedIP, _ptr_ReservedIp); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV6_ReservedIP) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ReservedIp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ReservedIP == nil {
			o.ReservedIP = &IPReservationV6{}
		}
		if err := o.ReservedIP.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ReservedIp := func(ptr interface{}) { o.ReservedIP = *ptr.(**IPReservationV6) }
	if err := w.ReadPointer(&o.ReservedIP, _s_ReservedIp, _ptr_ReservedIp); err != nil {
		return err
	}
	return nil
}

// SubnetElementDataV6_ExcludeIPRange structure represents SubnetElementDataV6_Element RPC union arm.
//
// It has following labels: 2
type SubnetElementDataV6_ExcludeIPRange struct {
	// ExcludeIpRange:  This is a pointer to a DHCP_IP_RANGE_V6 (section 2.2.1.2.59) structure
	// that contains information about IPv6 exclusion ranges.
	ExcludeIPRange *IPRangeV6 `idl:"name:ExcludeIpRange" json:"exclude_ip_range"`
}

func (*SubnetElementDataV6_ExcludeIPRange) is_SubnetElementDataV6_Element() {}

func (o *SubnetElementDataV6_ExcludeIPRange) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ExcludeIPRange != nil {
		_ptr_ExcludeIpRange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ExcludeIPRange != nil {
				if err := o.ExcludeIPRange.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPRangeV6{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ExcludeIPRange, _ptr_ExcludeIpRange); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementDataV6_ExcludeIPRange) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ExcludeIpRange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ExcludeIPRange == nil {
			o.ExcludeIPRange = &IPRangeV6{}
		}
		if err := o.ExcludeIPRange.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ExcludeIpRange := func(ptr interface{}) { o.ExcludeIPRange = *ptr.(**IPRangeV6) }
	if err := w.ReadPointer(&o.ExcludeIPRange, _s_ExcludeIpRange, _ptr_ExcludeIpRange); err != nil {
		return err
	}
	return nil
}

// SubnetElementInfoArrayV6 structure represents DHCP_SUBNET_ELEMENT_INFO_ARRAY_V6 RPC structure.
//
// The DHCP_SUBNET_ELEMENT_INFO_ARRAY_V6 structure defines an array of DHCP_SUBNET_ELEMENT_DATA_V6
// (section 2.2.1.2.60) structures of IPv6 prefix elements. The first member contains
// the number of IPv6 prefix elements (such as IPv6 reservation, IPv6 exclusion range,
// and IPv6 range), and the second member points to the array of length NumElements
// containing DHCPv6 IPv6 prefix elements. This structure is used in the R_DhcpEnumSubnetElementsV6
// (section 3.2.4.61) method.
type SubnetElementInfoArrayV6 struct {
	// NumElements:  This is of type DWORD, containing the number of IPv6 subnet elements
	// in the subsequent field the Elements member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This is a pointer to an array of DHCP_SUBNET_ELEMENT_DATA_V6 (section
	// 2.2.1.2.60) structures of length NumElements containing IPv6 prefix elements.
	Elements []*SubnetElementDataV6 `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *SubnetElementInfoArrayV6) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementInfoArrayV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SubnetElementDataV6{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SubnetElementDataV6{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SubnetElementInfoArrayV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*SubnetElementDataV6, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &SubnetElementDataV6{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*SubnetElementDataV6) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// ResumeIPv6Handle structure represents DHCP_RESUME_IPV6_HANDLE RPC structure.
type ResumeIPv6Handle IPv6Address

func (o *ResumeIPv6Handle) IPv6Address() *IPv6Address { return (*IPv6Address)(o) }

func (o *ResumeIPv6Handle) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ResumeIPv6Handle) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.HighOrderBits); err != nil {
		return err
	}
	if err := w.WriteData(o.LowOrderBits); err != nil {
		return err
	}
	return nil
}
func (o *ResumeIPv6Handle) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.HighOrderBits); err != nil {
		return err
	}
	if err := w.ReadData(&o.LowOrderBits); err != nil {
		return err
	}
	return nil
}

// HostInfoV6 structure represents DHCP_HOST_INFO_V6 RPC structure.
//
// The DHCP_HOST_INFO_V6 structure contains information on the DHCPv6 server. This structure
// is used in the DHCP_CLIENT_INFO_V6 (section 2.2.1.2.64) structure.
type HostInfoV6 struct {
	// IpAddress:   This is of type DHCP_IPV6_ADDRESS structure, containing the IPv6 address
	// of the DHCPv6 server.
	IPAddress *IPv6Address `idl:"name:IpAddress" json:"ip_address"`
	// NetBiosName:  A pointer to a null-terminated Unicode string. Currently not used in
	// any setting method, and if used in get method, the value returned is NULL.
	NetBIOSName string `idl:"name:NetBiosName" json:"net_bios_name"`
	// HostName:  A pointer to a null-terminated Unicode string. Currently not used in any
	// setting method, and if used in get method, the value returned is NULL.
	HostName string `idl:"name:HostName" json:"host_name"`
}

func (o *HostInfoV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *HostInfoV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.IPAddress != nil {
		if err := o.IPAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.NetBIOSName != "" {
		_ptr_NetBiosName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.NetBIOSName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.NetBIOSName, _ptr_NetBiosName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.HostName != "" {
		_ptr_HostName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.HostName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.HostName, _ptr_HostName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *HostInfoV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.IPAddress == nil {
		o.IPAddress = &IPv6Address{}
	}
	if err := o.IPAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_NetBiosName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.NetBIOSName); err != nil {
			return err
		}
		return nil
	})
	_s_NetBiosName := func(ptr interface{}) { o.NetBIOSName = *ptr.(*string) }
	if err := w.ReadPointer(&o.NetBIOSName, _s_NetBiosName, _ptr_NetBiosName); err != nil {
		return err
	}
	_ptr_HostName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.HostName); err != nil {
			return err
		}
		return nil
	})
	_s_HostName := func(ptr interface{}) { o.HostName = *ptr.(*string) }
	if err := w.ReadPointer(&o.HostName, _s_HostName, _ptr_HostName); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ClientInfoV6 structure represents DHCP_CLIENT_INFO_V6 RPC structure.
//
// The DHCP_CLIENT_INFO_V6 structure contains information on DHCPv6 client. This structure
// is used in the R_DhcpGetClientInfoV6 (section 3.2.4.73) method.
type ClientInfoV6 struct {
	// ClientIpAddress:  This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28), a structure
	// that contains the DHCPv6 client's IPv6 address.
	ClientIPAddress *IPv6Address `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// ClientDUID:  This is of type DHCP_CLIENT_UID (section 2.2.1.2.5), a structure that
	// represents the DHCPv6 client-identifier (section 2.2.1.2.5.3).
	ClientDUID *ClientUID `idl:"name:ClientDUID" json:"client_duid"`
	// AddressType:  This is of type DWORD that specifies the type of IPv6 address.
	//
	//	+------------------------------+--------------------------------------+
	//	|                              |                                      |
	//	|            VALUE             |               MEANING                |
	//	|                              |                                      |
	//	+------------------------------+--------------------------------------+
	//	+------------------------------+--------------------------------------+
	//	| ADDRESS_TYPE_IANA 0x00000000 | Indicates an IANA address. [RFC3315] |
	//	+------------------------------+--------------------------------------+
	//	| ADDRESS_TYPE_IATA 0x00000001 | Indicates an IATA address. [RFC3315] |
	//	+------------------------------+--------------------------------------+
	AddressType uint32 `idl:"name:AddressType" json:"address_type"`
	// IAID:  This is of type DWORD that specifies the interface identifier of the DHCPv6
	// client interface.
	IAID uint32 `idl:"name:IAID" json:"iaid"`
	// ClientName:  A pointer to a null-terminated Unicode string that contains the name
	// of the DHCPv6 client. There is no restriction on the length of this Unicode string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
	// ClientComment:  A pointer to a null-terminated Unicode string that contains a comment
	// relating to the DHCPv6 client. There is no restriction on the length of this Unicode
	// string.
	ClientComment string `idl:"name:ClientComment" json:"client_comment"`
	// ClientValidLeaseExpires:  This is of type DATE_TIME (section 2.2.1.2.11), a structure
	// that contains the valid lifetime of the DHCPv6 client lease.
	ClientValidLeaseExpires *DateTime `idl:"name:ClientValidLeaseExpires" json:"client_valid_lease_expires"`
	// ClientPrefLeaseExpires:  This is of type DATE_TIME, a structure that contains the
	// preferred lifetime of the DHCPv6 client lease.
	ClientPrefLeaseExpires *DateTime `idl:"name:ClientPrefLeaseExpires" json:"client_pref_lease_expires"`
	// OwnerHost:  This is of type DHCP_HOST_INFO_V6 (section 2.2.1.2.63), a structure that
	// contains information about the DHCPv6 server machine that has given this IPv6 lease
	// to this DHCPv6 client.
	OwnerHost *HostInfoV6 `idl:"name:OwnerHost" json:"owner_host"`
}

func (o *ClientInfoV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ClientIPAddress != nil {
		if err := o.ClientIPAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientDUID != nil {
		if err := o.ClientDUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AddressType); err != nil {
		return err
	}
	if err := w.WriteData(o.IAID); err != nil {
		return err
	}
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientComment != "" {
		_ptr_ClientComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientComment, _ptr_ClientComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientValidLeaseExpires != nil {
		if err := o.ClientValidLeaseExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientPrefLeaseExpires != nil {
		if err := o.ClientPrefLeaseExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OwnerHost != nil {
		if err := o.OwnerHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfoV6{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ClientIPAddress == nil {
		o.ClientIPAddress = &IPv6Address{}
	}
	if err := o.ClientIPAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ClientDUID == nil {
		o.ClientDUID = &ClientUID{}
	}
	if err := o.ClientDUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressType); err != nil {
		return err
	}
	if err := w.ReadData(&o.IAID); err != nil {
		return err
	}
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	_ptr_ClientComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClientComment := func(ptr interface{}) { o.ClientComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientComment, _s_ClientComment, _ptr_ClientComment); err != nil {
		return err
	}
	if o.ClientValidLeaseExpires == nil {
		o.ClientValidLeaseExpires = &DateTime{}
	}
	if err := o.ClientValidLeaseExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ClientPrefLeaseExpires == nil {
		o.ClientPrefLeaseExpires = &DateTime{}
	}
	if err := o.ClientPrefLeaseExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OwnerHost == nil {
		o.OwnerHost = &HostInfoV6{}
	}
	if err := o.OwnerHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ClientInfoArrayV6 structure represents DHCP_CLIENT_INFO_ARRAY_V6 RPC structure.
//
// The DHCP_CLIENT_INFO_ARRAY_V6 structure defines an array of DHCP_CLIENT_INFO_V6 (section
// 2.2.1.2.64) structures. The first member contains the number of DHCPv6 clients present
// in the specific prefix, and the second member contains a pointer to the array of
// length NumElements containing the DHCPv6 client's information. This structure is
// used by the methods that retrieve more than one DHCPv6 client's information.
type ClientInfoArrayV6 struct {
	// NumElements:  This is of type DWORD, containing the number of DHCPv6 clients in the
	// subsequent field the Clients member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Clients:  This is a pointer of type DHCP_CLIENT_INFO_V6 (section 2.2.1.2.64) structure
	// that points to the array of length NumElements containing the DHCPv6 client's information.
	Clients []*ClientInfoV6 `idl:"name:Clients;size_is:(NumElements)" json:"clients"`
}

func (o *ClientInfoArrayV6) xxx_PreparePayload(ctx context.Context) error {
	if o.Clients != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Clients))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoArrayV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Clients != nil || o.ElementsLength > 0 {
		_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Clients {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Clients[i1] != nil {
					_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Clients[i1] != nil {
							if err := o.Clients[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ClientInfoV6{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Clients[i1], _ptr_Clients); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Clients); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Clients, _ptr_Clients); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoArrayV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Clients", sizeInfo[0])
		}
		o.Clients = make([]*ClientInfoV6, sizeInfo[0])
		for i1 := range o.Clients {
			i1 := i1
			_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Clients[i1] == nil {
					o.Clients[i1] = &ClientInfoV6{}
				}
				if err := o.Clients[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Clients := func(ptr interface{}) { o.Clients[i1] = *ptr.(**ClientInfoV6) }
			if err := w.ReadPointer(&o.Clients[i1], _s_Clients, _ptr_Clients); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Clients := func(ptr interface{}) { o.Clients = *ptr.(*[]*ClientInfoV6) }
	if err := w.ReadPointer(&o.Clients, _s_Clients, _ptr_Clients); err != nil {
		return err
	}
	return nil
}

// ServerConfigInfoV6 structure represents DHCP_SERVER_CONFIG_INFO_V6 RPC structure.
//
// The DHCP_SERVER_CONFIG_INFO_V6 structure defines the settings for the DHCPv6 server.
// This structure is used in the R_DhcpServerSetConfigV6 (section 3.2.4.66) method.
type ServerConfigInfoV6 struct {
	// UnicastFlag:  This is of type BOOL, specifying whether the DHCPv6 client is allowed
	// to send unicast messages [RFC3315] to the server.
	UnicastFlag bool `idl:"name:UnicastFlag" json:"unicast_flag"`
	// RapidCommitFlag:  This is of type BOOL, specifying that server is to skip the AR
	// of the SARR [RFC3315] sequence in leasing a DHCPv6 client.
	RapidCommitFlag bool `idl:"name:RapidCommitFlag" json:"rapid_commit_flag"`
	// PreferredLifetime:  This is of type DWORD, specifying the preferred lifetime in seconds
	// for IANA addresses. [RFC3315]
	PreferredLifetime uint32 `idl:"name:PreferredLifetime" json:"preferred_lifetime"`
	// ValidLifetime:  This is of type DWORD, specifying the valid lifetime in seconds for
	// IANA addresses. [RFC3315]
	ValidLifetime uint32 `idl:"name:ValidLifetime" json:"valid_lifetime"`
	// T1:  This is of type DWORD, specifying the value for time T1 in seconds. [RFC3315]
	T1 uint32 `idl:"name:T1" json:"t1"`
	// T2:  This is of type DWORD, specifying value for time T2 in seconds. [RFC3315]
	T2 uint32 `idl:"name:T2" json:"t2"`
	// PreferredLifetimeIATA:  This is of type DWORD. Currently this is not implemented
	// and if used in setting the value through the method with any value, the method will
	// return ERROR_SUCCESS without any processing. If used in a method to retrieve, the
	// value returned is 86400 (1 day).<15>
	PreferredLifetimeIATA uint32 `idl:"name:PreferredLifetimeIATA" json:"preferred_lifetime_iata"`
	// ValidLifetimeIATA:  This is of type DWORD. Currently this is not implemented and
	// if used in setting the value through the method with any value, the method will return
	// ERROR_SUCCESS without any processing. If used in a method to retrieve, the value
	// returned is 259200 (3 days).<16>
	ValidLifetimeIATA uint32 `idl:"name:ValidLifetimeIATA" json:"valid_lifetime_iata"`
	// fAuditLog:  This is of type BOOL, specifying whether audit logs are enabled or disabled.
	// The field defaults to true to indicate that the audit logs are enabled.
	AuditLog bool `idl:"name:fAuditLog" json:"audit_log"`
}

func (o *ServerConfigInfoV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerConfigInfoV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.UnicastFlag {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.RapidCommitFlag {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PreferredLifetime); err != nil {
		return err
	}
	if err := w.WriteData(o.ValidLifetime); err != nil {
		return err
	}
	if err := w.WriteData(o.T1); err != nil {
		return err
	}
	if err := w.WriteData(o.T2); err != nil {
		return err
	}
	if err := w.WriteData(o.PreferredLifetimeIATA); err != nil {
		return err
	}
	if err := w.WriteData(o.ValidLifetimeIATA); err != nil {
		return err
	}
	if !o.AuditLog {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerConfigInfoV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bUnicastFlag int32
	if err := w.ReadData(&_bUnicastFlag); err != nil {
		return err
	}
	o.UnicastFlag = _bUnicastFlag != 0
	var _bRapidCommitFlag int32
	if err := w.ReadData(&_bRapidCommitFlag); err != nil {
		return err
	}
	o.RapidCommitFlag = _bRapidCommitFlag != 0
	if err := w.ReadData(&o.PreferredLifetime); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValidLifetime); err != nil {
		return err
	}
	if err := w.ReadData(&o.T1); err != nil {
		return err
	}
	if err := w.ReadData(&o.T2); err != nil {
		return err
	}
	if err := w.ReadData(&o.PreferredLifetimeIATA); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValidLifetimeIATA); err != nil {
		return err
	}
	var _bAuditLog int32
	if err := w.ReadData(&_bAuditLog); err != nil {
		return err
	}
	o.AuditLog = _bAuditLog != 0
	return nil
}

// ScopeMIBInfoV6 structure represents SCOPE_MIB_INFO_V6 RPC structure.
//
// The SCOPE_MIB_INFO_V6 structure defines a structure that contains the address counters
// for a specific IPv6 prefix. The numbers of free, used, and offered IPv6 addresses
// are stored in this structure. This structure is used in the DHCP_MIB_INFO_V6 (section
// 2.2.1.2.68) structure.
type ScopeMIBInfoV6 struct {
	// Subnet:  This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28), a structure specifying
	// the IPv6 prefix for the scope.
	Subnet *IPv6Address `idl:"name:Subnet" json:"subnet"`
	// NumAddressesInuse:  This is of type ULONGLONG, containing the number of IPv6 addresses
	// that have been leased to the DHCPv6 clients from the current prefix.
	AddressesInUseLength uint64 `idl:"name:NumAddressesInuse" json:"addresses_inuse_length"`
	// NumAddressesFree:  This is of type ULONGLONG, containing the number of addresses
	// that are free and can be leased out to DHCPv6 clients in the current subnet.
	AddressesFreeLength uint64 `idl:"name:NumAddressesFree" json:"addresses_free_length"`
	// NumPendingAdvertises:  This is of type ULONGLONG, containing the number of IPv6 addresses
	// that are advertised to the DHCPv6 clients from the prefix but that have not yet been
	// confirmed by the DHCPv6 client.
	PendingAdvertisesLength uint64 `idl:"name:NumPendingAdvertises" json:"pending_advertises_length"`
}

func (o *ScopeMIBInfoV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ScopeMIBInfoV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Subnet != nil {
		if err := o.Subnet.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AddressesInUseLength); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressesFreeLength); err != nil {
		return err
	}
	if err := w.WriteData(o.PendingAdvertisesLength); err != nil {
		return err
	}
	return nil
}
func (o *ScopeMIBInfoV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.Subnet == nil {
		o.Subnet = &IPv6Address{}
	}
	if err := o.Subnet.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressesInUseLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressesFreeLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.PendingAdvertisesLength); err != nil {
		return err
	}
	return nil
}

// MIBInfoV6 structure represents DHCP_MIB_INFO_V6 RPC structure.
//
// The DHCP_MIB_INFO_V6 structure contains the DHCPv6 server counter values for the
// DHCPv6 server. This structure is used in the R_DhcpGetMibInfoV6 (section 3.2.4.69)
// method to find DHCPv6 server statistics.
type MIBInfoV6 struct {
	// Solicits:  This is of type DWORD and contains the number of DHCPSOLICIT message received
	// by the DHCPv6 server from DHCPv6 clients since the DHCP server was last started.
	// This is used for statistical analysis by the DHCPv6 server.
	Solicits uint32 `idl:"name:Solicits" json:"solicits"`
	// Advertises:  This is of type DWORD and contains the number of DHCPADVERTISE message
	// sent by DHCPv6 server to DHCPv6 clients since the DHCP server was last started. This
	// is used for statistical analysis by the DHCPv6 server.
	Advertises uint32 `idl:"name:Advertises" json:"advertises"`
	// Requests:  This is of type DWORD and contains the number of DHCPREQUEST messages
	// received by the DHCPv6 server from DHCPv6 clients since the DHCP server was last
	// started. This is used for statistical analysis by the DHCPv6 server.
	Requests uint32 `idl:"name:Requests" json:"requests"`
	// Renews:  This is of type DWORD and contains the number of DHCPRENEW message received
	// by the DHCPv6 server from DHCPv6 clients since the DHCP server was last started.
	// This is used for statistical analysis by the DHCPv6 server.
	Renews uint32 `idl:"name:Renews" json:"renews"`
	// Rebinds:  This is of type DWORD and contains the number of DHCPREBIND messages received
	// by the DHCPv6 server from DHCPv6 clients since the DHCP server was last started.
	// This is used for statistical analysis by the DHCPv6 server.
	Rebinds uint32 `idl:"name:Rebinds" json:"rebinds"`
	// Replies:  This is of type DWORD and contains the number of DHCPREPLY messages sent
	// by the DHCPv6 server to DHCP clients since the DHCP server was last started. This
	// is used for statistical analysis by the DHCPv6 server.
	Replies uint32 `idl:"name:Replies" json:"replies"`
	// Confirms:  This is of type DWORD and contains the number of DHCPCONFIRM messages
	// received by the DHCPv6 server from DHCPv6 clients since the DHCP server was last
	// started. This is used for statistical analysis by the DHCPv6 server.
	Confirms uint32 `idl:"name:Confirms" json:"confirms"`
	// Declines:  This is of type DWORD and contains the number of DHCPDECLINES messages
	// received by the DHCPv6 server from DHCPv6 clients since the DHCP server was last
	// started. This is used for statistical analysis by the DHCPv6 server.
	Declines uint32 `idl:"name:Declines" json:"declines"`
	// Releases:  This is of type DWORD and contains the number of DHCPRELEASE messages
	// received by the DHCPv6 server from DHCPv6 clients since the DHCP server was last
	// started. This is used for statistical analysis by the DHCPv6 server.
	Releases uint32 `idl:"name:Releases" json:"releases"`
	// Informs:  This is of type DWORD and contains the number of DHCPINFORM messages received
	// by the DHCPv6 server from DHCPv6 clients since the DHCP server was last started.
	// This is used for statistical analysis by the DHCPv6 server.
	Informs uint32 `idl:"name:Informs" json:"informs"`
	// ServerStartTime:  This is of type DATE_TIME (section 2.2.1.2.11), a structure containing
	// the start time of the DHCPv6 server.
	ServerStartTime *DateTime `idl:"name:ServerStartTime" json:"server_start_time"`
	// Scopes:  This is of type DWORD, containing the number of IPv6 scopes configured on
	// the current DHCPv6 server. This is used for statistical analysis by the DHCPv6 server.
	// This field defines the number of DHCPv6 scopes in the subsequent field the ScopeInfo
	// member.
	Scopes uint32 `idl:"name:Scopes" json:"scopes"`
	// ScopeInfo:  This is a pointer to an array of SCOPE_MIB_INFO_V6 (section 2.2.1.2.67)
	// structures that points to an array of length Scopes, containing the information about
	// the IPv6 scopes configured on the DHCPv6 server.
	ScopeInfo []*ScopeMIBInfoV6 `idl:"name:ScopeInfo;size_is:(Scopes)" json:"scope_info"`
}

func (o *MIBInfoV6) xxx_PreparePayload(ctx context.Context) error {
	if o.ScopeInfo != nil && o.Scopes == 0 {
		o.Scopes = uint32(len(o.ScopeInfo))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBInfoV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Solicits); err != nil {
		return err
	}
	if err := w.WriteData(o.Advertises); err != nil {
		return err
	}
	if err := w.WriteData(o.Requests); err != nil {
		return err
	}
	if err := w.WriteData(o.Renews); err != nil {
		return err
	}
	if err := w.WriteData(o.Rebinds); err != nil {
		return err
	}
	if err := w.WriteData(o.Replies); err != nil {
		return err
	}
	if err := w.WriteData(o.Confirms); err != nil {
		return err
	}
	if err := w.WriteData(o.Declines); err != nil {
		return err
	}
	if err := w.WriteData(o.Releases); err != nil {
		return err
	}
	if err := w.WriteData(o.Informs); err != nil {
		return err
	}
	if o.ServerStartTime != nil {
		if err := o.ServerStartTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Scopes); err != nil {
		return err
	}
	if o.ScopeInfo != nil || o.Scopes > 0 {
		_ptr_ScopeInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Scopes)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ScopeInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ScopeInfo[i1] != nil {
					if err := o.ScopeInfo[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ScopeMIBInfoV6{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ScopeInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ScopeMIBInfoV6{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeInfo, _ptr_ScopeInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *MIBInfoV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Solicits); err != nil {
		return err
	}
	if err := w.ReadData(&o.Advertises); err != nil {
		return err
	}
	if err := w.ReadData(&o.Requests); err != nil {
		return err
	}
	if err := w.ReadData(&o.Renews); err != nil {
		return err
	}
	if err := w.ReadData(&o.Rebinds); err != nil {
		return err
	}
	if err := w.ReadData(&o.Replies); err != nil {
		return err
	}
	if err := w.ReadData(&o.Confirms); err != nil {
		return err
	}
	if err := w.ReadData(&o.Declines); err != nil {
		return err
	}
	if err := w.ReadData(&o.Releases); err != nil {
		return err
	}
	if err := w.ReadData(&o.Informs); err != nil {
		return err
	}
	if o.ServerStartTime == nil {
		o.ServerStartTime = &DateTime{}
	}
	if err := o.ServerStartTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Scopes); err != nil {
		return err
	}
	_ptr_ScopeInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Scopes > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Scopes)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ScopeInfo", sizeInfo[0])
		}
		o.ScopeInfo = make([]*ScopeMIBInfoV6, sizeInfo[0])
		for i1 := range o.ScopeInfo {
			i1 := i1
			if o.ScopeInfo[i1] == nil {
				o.ScopeInfo[i1] = &ScopeMIBInfoV6{}
			}
			if err := o.ScopeInfo[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ScopeInfo := func(ptr interface{}) { o.ScopeInfo = *ptr.(*[]*ScopeMIBInfoV6) }
	if err := w.ReadPointer(&o.ScopeInfo, _s_ScopeInfo, _ptr_ScopeInfo); err != nil {
		return err
	}
	return nil
}

// BindElementV6 structure represents DHCPV6_BIND_ELEMENT RPC structure.
//
// The DHCPV6_BIND_ELEMENT structure defines an IPv6 interface binding for the DHCP
// server over which it receives DHCPv6 packets. This structure is used in DHCPV6_BIND_ELEMENT_ARRAY
// (section 2.2.1.2.83) structure.
type BindElementV6 struct {
	// Flags:  This is of type ULONG, specifying a set of bit flags indicating properties
	// of the interface binding.
	//
	//	+-------------------------------------------+-----------------------------------+
	//	|                                           |                                   |
	//	|                   VALUE                   |              MEANING              |
	//	|                                           |                                   |
	//	+-------------------------------------------+-----------------------------------+
	//	+-------------------------------------------+-----------------------------------+
	//	| DHCP_ENDPOINT_FLAG_CANT_MODIFY 0x00000001 | The endpoints cannot be modified. |
	//	+-------------------------------------------+-----------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// fBoundToDHCPServer:  This is of type BOOL, specifying whether the interface is bound
	// to the DHCP server.
	//
	//	+------------------+------------------------------------------------------------------+
	//	|                  |                                                                  |
	//	|      VALUE       |                             MEANING                              |
	//	|                  |                                                                  |
	//	+------------------+------------------------------------------------------------------+
	//	+------------------+------------------------------------------------------------------+
	//	| FALSE 0x00000000 | It specifies that the interface is not bound to the DHCP server. |
	//	+------------------+------------------------------------------------------------------+
	//	| TRUE 0x00000001  | It specifies that the interface is bound to the DHCP server.     |
	//	+------------------+------------------------------------------------------------------+
	BoundToDHCPServer bool `idl:"name:fBoundToDHCPServer" json:"bound_to_dhcp_server"`
	// AdapterPrimaryAddress:  This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28) structure,
	// specifying the IPv6 address assigned to the interface over which the DHCP server
	// is receiving DHCPv6 packets.
	AdapterPrimaryAddress *IPv6Address `idl:"name:AdapterPrimaryAddress" json:"adapter_primary_address"`
	// AdapterSubnetAddress:  This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28) structure,
	// specifying the IPv6 prefix ID of the subnet from which this interface is receiving
	// DHCPv6 packets.
	AdapterSubnetAddress *IPv6Address `idl:"name:AdapterSubnetAddress" json:"adapter_subnet_address"`
	// IfDescription:  A pointer, of type LPWSTR, to a null-terminated Unicode string that
	// specifies the name assigned to this interface. The maximum number of characters allowed
	// in this string is 256, excluding the terminating null character.
	InterfaceDescription string `idl:"name:IfDescription" json:"interface_description"`
	// IpV6IfIndex:  This is of type DWORD, containing the IPv6 interface index of the current
	// interface.
	IPv6InterfaceIndex uint32 `idl:"name:IpV6IfIndex" json:"ipv6_interface_index"`
	// IfIdSize:  This is of type ULONG, containing the size of the interface GUID stored
	// in the  IfId member.
	InterfaceIDSize uint32 `idl:"name:IfIdSize" json:"interface_id_size"`
	// IfId:  This is a pointer to a BYTE, containing the interface GUID assigned to this
	// interface.
	InterfaceID []byte `idl:"name:IfId;size_is:(IfIdSize)" json:"interface_id"`
}

func (o *BindElementV6) xxx_PreparePayload(ctx context.Context) error {
	if o.InterfaceID != nil && o.InterfaceIDSize == 0 {
		o.InterfaceIDSize = uint32(len(o.InterfaceID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindElementV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if !o.BoundToDHCPServer {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.AdapterPrimaryAddress != nil {
		if err := o.AdapterPrimaryAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.AdapterSubnetAddress != nil {
		if err := o.AdapterSubnetAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.InterfaceDescription != "" {
		_ptr_IfDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.InterfaceDescription); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.InterfaceDescription, _ptr_IfDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.IPv6InterfaceIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfaceIDSize); err != nil {
		return err
	}
	if o.InterfaceID != nil || o.InterfaceIDSize > 0 {
		_ptr_IfId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InterfaceIDSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.InterfaceID {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.InterfaceID[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.InterfaceID); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InterfaceID, _ptr_IfId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *BindElementV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	var _bBoundToDHCPServer int32
	if err := w.ReadData(&_bBoundToDHCPServer); err != nil {
		return err
	}
	o.BoundToDHCPServer = _bBoundToDHCPServer != 0
	if o.AdapterPrimaryAddress == nil {
		o.AdapterPrimaryAddress = &IPv6Address{}
	}
	if err := o.AdapterPrimaryAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.AdapterSubnetAddress == nil {
		o.AdapterSubnetAddress = &IPv6Address{}
	}
	if err := o.AdapterSubnetAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_IfDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.InterfaceDescription); err != nil {
			return err
		}
		return nil
	})
	_s_IfDescription := func(ptr interface{}) { o.InterfaceDescription = *ptr.(*string) }
	if err := w.ReadPointer(&o.InterfaceDescription, _s_IfDescription, _ptr_IfDescription); err != nil {
		return err
	}
	if err := w.ReadData(&o.IPv6InterfaceIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfaceIDSize); err != nil {
		return err
	}
	_ptr_IfId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InterfaceIDSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InterfaceIDSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.InterfaceID", sizeInfo[0])
		}
		o.InterfaceID = make([]byte, sizeInfo[0])
		for i1 := range o.InterfaceID {
			i1 := i1
			if err := w.ReadData(&o.InterfaceID[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_IfId := func(ptr interface{}) { o.InterfaceID = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.InterfaceID, _s_IfId, _ptr_IfId); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// BindElementArrayV6 structure represents DHCPV6_BIND_ELEMENT_ARRAY RPC structure.
//
// The DHCPV6_BIND_ELEMENT_ARRAY structure defines an array of DHCPV6_BIND_ELEMENT (section
// 2.2.1.2.82) structures. This contains an array of IPv6 interface binding over which
// the DHCPv6 server receives DHCPv6 packets. The first member contains the number of
// IPv6 interface bindings present in the specific subnet, and the second member points
// to the array of interface bindings over which the DHCPv6 server is receiving DHCPv6
// packets.
type BindElementArrayV6 struct {
	// NumElements:  This is of type DWORD and specifies the number of IPv6 interface binding
	// listed in subsequent field the Elements member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This is a pointer to an array of type DHCPV6_BIND_ELEMENT (section 2.2.1.2.82)
	// structure and length NumElements that contains information for interface bindings
	// for a DHCPv6 server.
	Elements []*BindElementV6 `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *BindElementArrayV6) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindElementArrayV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&BindElementV6{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&BindElementV6{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindElementArrayV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*BindElementV6, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &BindElementV6{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*BindElementV6) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// SearchInfoV6 structure represents DHCP_SEARCH_INFO_V6 RPC structure.
//
// The DHCP_SEARCH_INFO_V6 structure contains the DHCPv6 client information search type
// defined by the SearchType member, along with the data supporting that search. With
// this structure, a search is performed for a specific DHCPv6 client. This structure
// is used in the R_DhcpGetClientInfoV6 (section 3.2.4.73) method.
type SearchInfoV6 struct {
	// SearchType:  This is an enumeration value of type DHCP_SEARCH_INFO_TYPE_V6 (section
	// 2.2.1.1.12) enumeration that contains the data type, based on which the search is
	// performed, for a specific DHCPv6 client record on the DHCPv6 server.
	SearchType SearchInfoTypeV6 `idl:"name:SearchType" json:"search_type"`
	// SearchInfo:  This is a union that can contain one of the following values chosen
	// based on the value of the SearchType member.
	SearchInfo *SearchInfoV6_SearchInfo `idl:"name:SearchInfo;switch_is:SearchType" json:"search_info"`
}

func (o *SearchInfoV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SearchInfoV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.SearchType)); err != nil {
		return err
	}
	_swSearchInfo := uint16(o.SearchType)
	if o.SearchInfo != nil {
		if err := o.SearchInfo.MarshalUnionNDR(ctx, w, _swSearchInfo); err != nil {
			return err
		}
	} else {
		if err := (&SearchInfoV6_SearchInfo{}).MarshalUnionNDR(ctx, w, _swSearchInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *SearchInfoV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.SearchType)); err != nil {
		return err
	}
	if o.SearchInfo == nil {
		o.SearchInfo = &SearchInfoV6_SearchInfo{}
	}
	_swSearchInfo := uint16(o.SearchType)
	if err := o.SearchInfo.UnmarshalUnionNDR(ctx, w, _swSearchInfo); err != nil {
		return err
	}
	return nil
}

// SearchInfoV6_SearchInfo structure represents DHCP_SEARCH_INFO_V6 union anonymous member.
//
// The DHCP_SEARCH_INFO_V6 structure contains the DHCPv6 client information search type
// defined by the SearchType member, along with the data supporting that search. With
// this structure, a search is performed for a specific DHCPv6 client. This structure
// is used in the R_DhcpGetClientInfoV6 (section 3.2.4.73) method.
type SearchInfoV6_SearchInfo struct {
	// Types that are assignable to Value
	//
	// *SearchInfoV6_ClientIPAddress
	// *SearchInfoV6_ClientDUID
	// *SearchInfoV6_ClientName
	Value is_SearchInfoV6_SearchInfo `json:"value"`
}

func (o *SearchInfoV6_SearchInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SearchInfoV6_ClientIPAddress:
		if value != nil {
			return value.ClientIPAddress
		}
	case *SearchInfoV6_ClientDUID:
		if value != nil {
			return value.ClientDUID
		}
	case *SearchInfoV6_ClientName:
		if value != nil {
			return value.ClientName
		}
	}
	return nil
}

type is_SearchInfoV6_SearchInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SearchInfoV6_SearchInfo()
}

func (o *SearchInfoV6_SearchInfo) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SearchInfoV6_ClientIPAddress:
		return uint16(0)
	case *SearchInfoV6_ClientDUID:
		return uint16(1)
	case *SearchInfoV6_ClientName:
		return uint16(2)
	}
	return uint16(0)
}

func (o *SearchInfoV6_SearchInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*SearchInfoV6_ClientIPAddress)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SearchInfoV6_ClientIPAddress{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*SearchInfoV6_ClientDUID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SearchInfoV6_ClientDUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*SearchInfoV6_ClientName)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SearchInfoV6_ClientName{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SearchInfoV6_SearchInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &SearchInfoV6_ClientIPAddress{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &SearchInfoV6_ClientDUID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &SearchInfoV6_ClientName{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SearchInfoV6_ClientIPAddress structure represents SearchInfoV6_SearchInfo RPC union arm.
//
// It has following labels: 0
type SearchInfoV6_ClientIPAddress struct {
	// ClientIpAddress:  This is of type DHCP_IPV6_ADDRESS (section 2.2.1.2.28), a structure
	// that contains the IPv6 address of the DHCPv6 client lease record. It is used for
	// searching in the DHCPv6 server database.
	ClientIPAddress *IPv6Address `idl:"name:ClientIpAddress" json:"client_ip_address"`
}

func (*SearchInfoV6_ClientIPAddress) is_SearchInfoV6_SearchInfo() {}

func (o *SearchInfoV6_ClientIPAddress) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ClientIPAddress != nil {
		if err := o.ClientIPAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SearchInfoV6_ClientIPAddress) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ClientIPAddress == nil {
		o.ClientIPAddress = &IPv6Address{}
	}
	if err := o.ClientIPAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SearchInfoV6_ClientDUID structure represents SearchInfoV6_SearchInfo RPC union arm.
//
// It has following labels: 1
type SearchInfoV6_ClientDUID struct {
	// ClientDUID:  This is of type DHCP_CLIENT_UID (section 2.2.1.2.5), a structure that
	// represents the DHCPv6 client-identifier (section 2.2.1.2.5.3). It is used to search
	// for the DHCPv6 client lease record in the DHCPv6 server database.
	ClientDUID *ClientUID `idl:"name:ClientDUID" json:"client_duid"`
}

func (*SearchInfoV6_ClientDUID) is_SearchInfoV6_SearchInfo() {}

func (o *SearchInfoV6_ClientDUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ClientDUID != nil {
		if err := o.ClientDUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SearchInfoV6_ClientDUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ClientDUID == nil {
		o.ClientDUID = &ClientUID{}
	}
	if err := o.ClientDUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SearchInfoV6_ClientName structure represents SearchInfoV6_SearchInfo RPC union arm.
//
// It has following labels: 2
type SearchInfoV6_ClientName struct {
	// ClientName:  A pointer to a null-terminated Unicode string that contains the name
	// of the DHCPv6 client. It is used to search for the DHCPv6 client lease record in
	// the DHCPv6 server database. There is no restriction on the length of this Unicode
	// string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
}

func (*SearchInfoV6_ClientName) is_SearchInfoV6_SearchInfo() {}

func (o *SearchInfoV6_ClientName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SearchInfoV6_ClientName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	return nil
}

// ClassInfoV6 structure represents DHCP_CLASS_INFO_V6 RPC structure.
//
// The DHCP_CLASS_INFO_V6 structure contains the information for a specific user class
// or vendor class. This structure is used in the R_DhcpCreateClassV6 (section 3.2.4.75)
// method.
type ClassInfoV6 struct {
	// ClassName:  A pointer, of type LPWSTR, to a null-terminated Unicode string that contains
	// the class name. There is no restriction on the length of this Unicode string.
	ClassName string `idl:"name:ClassName" json:"class_name"`
	// ClassComment:  A pointer, of type LPWSTR, to a null-terminated Unicode string that
	// contains the comment for the class. There is no restriction on the length of this
	// Unicode string.
	ClassComment string `idl:"name:ClassComment" json:"class_comment"`
	// ClassDataLength:  This is of type DWORD, containing the length of data as pointed
	// to by the ClassData member.
	ClassDataLength uint32 `idl:"name:ClassDataLength" json:"class_data_length"`
	// IsVendor:  This is of type BOOL and specifies whether the current class is vendor
	// class or user class.
	//
	//	+------------+------------------------------------+
	//	|            |                                    |
	//	|   VALUE    |              MEANING               |
	//	|            |                                    |
	//	+------------+------------------------------------+
	//	+------------+------------------------------------+
	//	| 0x00000000 | Class specified is a user class.   |
	//	+------------+------------------------------------+
	//	| 0x00000001 | Class specified is a vendor class. |
	//	+------------+------------------------------------+
	IsVendor bool `idl:"name:IsVendor" json:"is_vendor"`
	// EnterpriseNumber:  This is of type DWORD, containing the vendor class identifier.
	// It is default 0 for user class.
	EnterpriseNumber uint32 `idl:"name:EnterpriseNumber" json:"enterprise_number"`
	// Flags:  This is of type DWORD. Currently it is not used, and any value set to this
	// parameter will not affect the behavior of the method that uses this structure.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ClassData:  This is a pointer of type BYTE that points to an array of bytes of length
	// specified by the ClassDataLength member. This contains data regarding a user class
	// or a vendor class.
	ClassData []byte `idl:"name:ClassData;size_is:(ClassDataLength)" json:"class_data"`
}

func (o *ClassInfoV6) xxx_PreparePayload(ctx context.Context) error {
	if o.ClassData != nil && o.ClassDataLength == 0 {
		o.ClassDataLength = uint32(len(o.ClassData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassInfoV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ClassName != "" {
		_ptr_ClassName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClassName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClassName, _ptr_ClassName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClassComment != "" {
		_ptr_ClassComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClassComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClassComment, _ptr_ClassComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ClassDataLength); err != nil {
		return err
	}
	if !o.IsVendor {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.EnterpriseNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.ClassData != nil || o.ClassDataLength > 0 {
		_ptr_ClassData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ClassDataLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ClassData {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.ClassData[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.ClassData); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ClassData, _ptr_ClassData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassInfoV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_ClassName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClassName); err != nil {
			return err
		}
		return nil
	})
	_s_ClassName := func(ptr interface{}) { o.ClassName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClassName, _s_ClassName, _ptr_ClassName); err != nil {
		return err
	}
	_ptr_ClassComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClassComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClassComment := func(ptr interface{}) { o.ClassComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClassComment, _s_ClassComment, _ptr_ClassComment); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClassDataLength); err != nil {
		return err
	}
	var _bIsVendor int32
	if err := w.ReadData(&_bIsVendor); err != nil {
		return err
	}
	o.IsVendor = _bIsVendor != 0
	if err := w.ReadData(&o.EnterpriseNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_ClassData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ClassDataLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ClassDataLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ClassData", sizeInfo[0])
		}
		o.ClassData = make([]byte, sizeInfo[0])
		for i1 := range o.ClassData {
			i1 := i1
			if err := w.ReadData(&o.ClassData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ClassData := func(ptr interface{}) { o.ClassData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ClassData, _s_ClassData, _ptr_ClassData); err != nil {
		return err
	}
	return nil
}

// ClassInfoArrayV6 structure represents DHCP_CLASS_INFO_ARRAY_V6 RPC structure.
//
// The DHCP_CLASS_INFO_ARRAY_V6 structure contains a list of information regarding a
// user class or a vendor class.
type ClassInfoArrayV6 struct {
	// NumElements:  This is of type DWORD, specifying the number of classes whose information
	// is contained in the array specified by the Classes member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Classes:  A pointer to an array of DHCP_CLASS_INFO_V6 (section 2.2.1.2.70) structures
	// that contains information regarding the various user classes and vendor classes.
	Classes []*ClassInfoV6 `idl:"name:Classes;size_is:(NumElements)" json:"classes"`
}

func (o *ClassInfoArrayV6) xxx_PreparePayload(ctx context.Context) error {
	if o.Classes != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Classes))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassInfoArrayV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Classes != nil || o.ElementsLength > 0 {
		_ptr_Classes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Classes {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Classes[i1] != nil {
					if err := o.Classes[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ClassInfoV6{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Classes); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ClassInfoV6{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Classes, _ptr_Classes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassInfoArrayV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Classes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Classes", sizeInfo[0])
		}
		o.Classes = make([]*ClassInfoV6, sizeInfo[0])
		for i1 := range o.Classes {
			i1 := i1
			if o.Classes[i1] == nil {
				o.Classes[i1] = &ClassInfoV6{}
			}
			if err := o.Classes[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Classes := func(ptr interface{}) { o.Classes = *ptr.(*[]*ClassInfoV6) }
	if err := w.ReadPointer(&o.Classes, _s_Classes, _ptr_Classes); err != nil {
		return err
	}
	return nil
}

// ClientFilterStatusInfo structure represents DHCP_CLIENT_FILTER_STATUS_INFO RPC structure.
//
// The DHCP_CLIENT_FILTER_STATUS_INFO structure defines information about the DHCPv4
// client, including filter status information. The DHCP_CLIENT_FILTER_STATUS_INFO structure
// augments the DHCP_CLIENT_INFO_VQ (section 2.2.1.2.19) structure by including information
// related to the filters applicable to a DHCPv4 client.
type ClientFilterStatusInfo struct {
	// ClientIpAddress:  This is of type DHCP_IP_ADDRESS (section 2.2.1.2.1), a DWORD that
	// contains the DHCPv4 client's IPv4 address.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// SubnetMask:  This is of type DHCP_IP_MASK (section 2.2.1.2.2), a DWORD that contains
	// the DHCPv4 client's IPv4 Subnet mask address.
	SubnetMask uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
	// ClientHardwareAddress:  This is of type DHCP_CLIENT_UID (section 2.2.1.2.5), a structure
	// that represents the DHCPv4 client-identifier (section 2.2.1.2.5.1).
	ClientHardwareAddress *ClientUID `idl:"name:ClientHardwareAddress" json:"client_hardware_address"`
	// ClientName:  A pointer, of type LPWSTR, to a null-terminated Unicode string that
	// represents the DHCPv4 client's Internet host name. There is no restriction on the
	// length of this Unicode string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
	// ClientComment:  A pointer, of type LPWSTR, to a null-terminated Unicode string that
	// represents the description given to the DHCPv4 client. The maximum number of characters
	// allowed in this string is 128, including the terminating null character.
	ClientComment string `idl:"name:ClientComment" json:"client_comment"`
	// ClientLeaseExpires:  This is of type DATE_TIME (section 2.2.1.2.11), a structure
	// that contains the lease expiry time for the DHCPv4 client. This is UTC time.
	ClientLeaseExpires *DateTime `idl:"name:ClientLeaseExpires" json:"client_lease_expires"`
	// OwnerHost:  This of type DHCP_HOST_INFO (section 2.2.1.2.7), a structure that contains
	// information about the DHCPv4 server machine that has provided a lease to the DHCPv4
	// client.
	OwnerHost *HostInfo `idl:"name:OwnerHost" json:"owner_host"`
	// bClientType:  This is of type BYTE that identifies the type of the DHCPv4 client.
	// The possible values are shown in the following table.
	//
	//	+-----------------------------------+----------------------------------------------------------------+
	//	|                                   |                                                                |
	//	|               VALUE               |                            MEANING                             |
	//	|                                   |                                                                |
	//	+-----------------------------------+----------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------+
	//	| CLIENT_TYPE_UNSPECIFIED 0x00      | A DHCPv4 client other than ones defined in this table.         |
	//	+-----------------------------------+----------------------------------------------------------------+
	//	| CLIENT_TYPE_DCHP 0x01             | The DHCPv4 client supports the DHCP protocol.                  |
	//	+-----------------------------------+----------------------------------------------------------------+
	//	| CLIENT_TYPE_BOOTP 0x02            | The DHCPv4 client supports the BOOTP protocol ([RFC2132]).     |
	//	+-----------------------------------+----------------------------------------------------------------+
	//	| CLIENT_TYPE_BOTH 0x03             | The DHCPv4 client identifies both the DHCP and BOOTP protocol. |
	//	+-----------------------------------+----------------------------------------------------------------+
	//	| CLIENT_TYPE_RESERVATION_FLAG 0x04 | There is an IPv4 reservation created for the DHCPv4 client.    |
	//	+-----------------------------------+----------------------------------------------------------------+
	//	| CLIENT_TYPE_NONE 0x05             | Backward compatibility for manual addressing.                  |
	//	+-----------------------------------+----------------------------------------------------------------+
	ClientType uint8 `idl:"name:bClientType" json:"client_type"`
	// AddressState:  This is of type BYTE that represents the state of the IPv4 address
	// given to the DHCPv4 client. The following table represents the different values and
	// their meanings.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_OFFERED 0x00  | The DHCPv4 client has been offered this IPv4 address.                            |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_ACTIVE 0x01   | The IPv4 address is active and has an active DHCPv4 client lease record.         |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DECLINED 0x02 | The IPv4 address request was declined by the DHCPv4 client; hence it is a bad    |
	//	|                             | IPv4 address.                                                                    |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DOOM 0x03     | The IPv4 address is in DOOMED state and is due to be deleted.                    |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	AddressState uint8 `idl:"name:AddressState" json:"address_state"`
	// Status:  This is of type QuarantineStatus (section 2.2.1.1.11), an enumeration that
	// contains the health status of the DHCPv4 client, as validated at the NAP server.
	Status QuarantineStatus `idl:"name:Status" json:"status"`
	// ProbationEnds:  This is of type DATE_TIME, a structure that contains the end time
	// of the probation if the DHCPv4 client is on probation. For this time period, the
	// DHCPv4 client has full access to the network.
	ProbationEnds *DateTime `idl:"name:ProbationEnds" json:"probation_ends"`
	// QuarantineCapable:  This is of type BOOL that can take on the following values.
	//
	//	+---------+------------------------------------------------------+
	//	|         |                                                      |
	//	|  VALUE  |                       MEANING                        |
	//	|         |                                                      |
	//	+---------+------------------------------------------------------+
	//	+---------+------------------------------------------------------+
	//	| TRUE 1  | The DHCPv4 client machine is quarantine-enabled.     |
	//	+---------+------------------------------------------------------+
	//	| FALSE 0 | The DHCPv4 client machine is not quarantine-enabled. |
	//	+---------+------------------------------------------------------+
	QuarantineCapable bool `idl:"name:QuarantineCapable" json:"quarantine_capable"`
	// FilterStatus:  This is of type DWORD that specifies the status of the link-layer
	// filter.
	//
	//	+-------------------------------------------------------+-----------------------------------------------------------------------+
	//	|                                                       |                                                                       |
	//	|                         VALUE                         |                                MEANING                                |
	//	|                                                       |                                                                       |
	//	+-------------------------------------------------------+-----------------------------------------------------------------------+
	//	+-------------------------------------------------------+-----------------------------------------------------------------------+
	//	| FILTER_STATUS_NONE 0x00000001                         | The DHCPv4 client MAC address does not match any filter.              |
	//	+-------------------------------------------------------+-----------------------------------------------------------------------+
	//	| FILTER_STATUS_FULL_MATCH_IN_ALLOW_LIST 0x00000002     | The DHCv4P client MAC address fully matches an allow list filter.     |
	//	+-------------------------------------------------------+-----------------------------------------------------------------------+
	//	| FILTER_STATUS_FULL_MATCH_IN_DENY_LIST 0x00000004      | The DHCPv4 client MAC address fully matches a deny list filter.       |
	//	+-------------------------------------------------------+-----------------------------------------------------------------------+
	//	| FILTER_STATUS_WILDCARD_MATCH_IN_ALLOW_LIST 0x00000008 | The DHCPv4 client MAC address has a wildcard match in the allow list. |
	//	+-------------------------------------------------------+-----------------------------------------------------------------------+
	//	| FILTER_STATUS_WILDCARD_MATCH_IN_DENY_LIST 0x00000010  | The DHCPv4 client MAC address has a wildcard match in the deny list.  |
	//	+-------------------------------------------------------+-----------------------------------------------------------------------+
	FilterStatus uint32 `idl:"name:FilterStatus" json:"filter_status"`
}

func (o *ClientFilterStatusInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientFilterStatusInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress != nil {
		if err := o.ClientHardwareAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientComment != "" {
		_ptr_ClientComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientComment, _ptr_ClientComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientLeaseExpires != nil {
		if err := o.ClientLeaseExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OwnerHost != nil {
		if err := o.OwnerHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ClientType); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressState); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if o.ProbationEnds != nil {
		if err := o.ProbationEnds.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.QuarantineCapable {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.FilterStatus); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ClientFilterStatusInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress == nil {
		o.ClientHardwareAddress = &ClientUID{}
	}
	if err := o.ClientHardwareAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	_ptr_ClientComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClientComment := func(ptr interface{}) { o.ClientComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientComment, _s_ClientComment, _ptr_ClientComment); err != nil {
		return err
	}
	if o.ClientLeaseExpires == nil {
		o.ClientLeaseExpires = &DateTime{}
	}
	if err := o.ClientLeaseExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OwnerHost == nil {
		o.OwnerHost = &HostInfo{}
	}
	if err := o.OwnerHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressState); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	if o.ProbationEnds == nil {
		o.ProbationEnds = &DateTime{}
	}
	if err := o.ProbationEnds.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bQuarantineCapable int32
	if err := w.ReadData(&_bQuarantineCapable); err != nil {
		return err
	}
	o.QuarantineCapable = _bQuarantineCapable != 0
	if err := w.ReadData(&o.FilterStatus); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ClientFilterStatusInfoArray structure represents DHCP_CLIENT_FILTER_STATUS_INFO_ARRAY RPC structure.
//
// The DHCP_CLIENT_FILTER_STATUS_INFO_ARRAY structure defines an array of DHCP_CLIENT_FILTER_STATUS_INFO
// (section 2.2.1.2.96) structures that contains a list of DHCPv4 client information.
// This structure is used by methods such as R_DhcpEnumSubnetClientsFilterStatusInfo
// (section 3.2.4.89) that retrieve information for more than one DHCPv4 client.
type ClientFilterStatusInfoArray struct {
	// NumElements:  This member is of type DWORD that contains the number of DHCPv4 clients
	// in the subsequent field the Clients member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Clients:  This member is a pointer of type DHCP_CLIENT_FILTER_STATUS_INFO (section
	// 2.2.1.2.96) structure that points to the array of length NumElements containing the
	// DHCPv4 client's information.
	Clients []*ClientFilterStatusInfo `idl:"name:Clients;size_is:(NumElements)" json:"clients"`
}

func (o *ClientFilterStatusInfoArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Clients != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Clients))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientFilterStatusInfoArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Clients != nil || o.ElementsLength > 0 {
		_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Clients {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Clients[i1] != nil {
					_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Clients[i1] != nil {
							if err := o.Clients[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ClientFilterStatusInfo{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Clients[i1], _ptr_Clients); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Clients); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Clients, _ptr_Clients); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientFilterStatusInfoArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Clients", sizeInfo[0])
		}
		o.Clients = make([]*ClientFilterStatusInfo, sizeInfo[0])
		for i1 := range o.Clients {
			i1 := i1
			_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Clients[i1] == nil {
					o.Clients[i1] = &ClientFilterStatusInfo{}
				}
				if err := o.Clients[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Clients := func(ptr interface{}) { o.Clients[i1] = *ptr.(**ClientFilterStatusInfo) }
			if err := w.ReadPointer(&o.Clients[i1], _s_Clients, _ptr_Clients); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Clients := func(ptr interface{}) { o.Clients = *ptr.(*[]*ClientFilterStatusInfo) }
	if err := w.ReadPointer(&o.Clients, _s_Clients, _ptr_Clients); err != nil {
		return err
	}
	return nil
}

// FailoverMode type represents DHCP_FAILOVER_MODE RPC enumeration.
//
// The DHCP_FAILOVER_MODE enumeration defines a set of possible modes of operation for
// a failover relationship configured on a DHCPv4 server.
type FailoverMode uint16

var (
	// LoadBalance:  Configures a DHCPv4 server failover relationship in a loadbalance
	// mode.
	FailoverModeLoadBalance FailoverMode = 0
	// HotStandby:  Configures a DHCPv4 server failover relationship in a hotstandby mode.
	FailoverModeHotStandby FailoverMode = 1
)

func (o FailoverMode) String() string {
	switch o {
	case FailoverModeLoadBalance:
		return "FailoverModeLoadBalance"
	case FailoverModeHotStandby:
		return "FailoverModeHotStandby"
	}
	return "Invalid"
}

// FailoverServer type represents DHCP_FAILOVER_SERVER RPC enumeration.
//
// The DHCP_FAILOVER_SERVER enumeration defines a set of possible values for a DHCPv4
// server in a failover relationship.
type FailoverServer uint16

var (
	// PrimaryServer:  The server is a primary server in the failover relationship.
	FailoverServerPrimaryServer FailoverServer = 0
	// SecondaryServer:  The server is a secondary server in the failover relationship.
	FailoverServerSecondaryServer FailoverServer = 1
)

func (o FailoverServer) String() string {
	switch o {
	case FailoverServerPrimaryServer:
		return "FailoverServerPrimaryServer"
	case FailoverServerSecondaryServer:
		return "FailoverServerSecondaryServer"
	}
	return "Invalid"
}

// FSMState type represents FSM_STATE RPC enumeration.
//
// The FSM_STATE enumeration defines a set of possible values representing various failover
// relationship states on a DHCPv4 server. For additional information about server state
// transitions, see [IETF-DHCPFOP-12], section 9.2.
type FSMState uint16

var (
	// NO_STATE: This value means that no state is configured for the DHCPv4 failover relationship.
	FSMStateNoState FSMState = 0
	// INIT: This value means that the failover relationship on the DHCPv4 server is in
	// the initializing state.
	FSMStateInit FSMState = 1
	// STARTUP: This value means that each server participating in the failover relationship
	// moves into the STARTUP state after initializing itself. The STARTUP state enables
	// a server to probe its partner server, before starting DHCP client service.
	FSMStateStartup FSMState = 2
	// NORMAL: This value means that each server services DHCPDISCOVER messages [RFC2131]
	// and all other DHCP requests, other than DHCPREQUEST/RENEWAL or DHCPREQUEST/REBINDING
	// requests from the client set, as defined by the load balancing algorithm specified
	// in [RFC3074]. Each server services DHCPREQUEST/RENEWAL or DHCPDISCOVER/REBINDING
	// requests from any client.
	FSMStateNormal FSMState = 3
	// COMMUNICATION_INT: This value means that each server in a failover relationship is
	// operating independently in the COMMUNICATION_INT state, but neither assumes that
	// its partner is not operating. The partner server might be operating and simply unable
	// to communicate with this server, or it might not be operating at all.
	FSMStateCommunicationInt FSMState = 4
	// PARTNER_DOWN: This value means that when operating in the PARTNER_DOWN state, a server
	// assumes that its partner is not currently operating.
	FSMStatePartnerDown FSMState = 5
	// POTENTIAL_CONFLICT: This value indicates that the failover relationship between two
	// DHCP servers is attempting to re-establish itself.
	FSMStatePotentialConflict FSMState = 6
	// CONFLICT_DONE: This value indicates that during the process where two servers in
	// a failover relationship attempt reintegration with each other, the primary server
	// has received all updates from the secondary server.
	FSMStateConflictDone FSMState = 7
	// RESOLUTION_INT: This value indicates that the two servers in a failover relationship
	// were attempting reintegration with each other in the POTENTIAL_CONFLICT state, but
	// communications failed prior to completion of the reintegration.
	FSMStateResolutionInt FSMState = 8
	// RECOVER: This value indicates that a server in a failover relationship has no information
	// in its stable storage facility or that it is reintegrating with a server in the PARTNER_DOWN
	// state after it has been down.
	FSMStateRecover FSMState = 9
	// RECOVER_WAIT: This value means that the DHCPv4 server waits for a time period equal
	// to maximum client lead time (MCLT) before moving to the RECOVER_DONE state.
	FSMStateRecoverWait FSMState = 10
	// RECOVER_DONE: This value enables an interlocked transition of one server from the
	// RECOVER state and another server from the PARTNER_DOWN or COMMUNICATION-INT state
	// to the NORMAL state.
	FSMStateRecoverDone FSMState = 11
)

func (o FSMState) String() string {
	switch o {
	case FSMStateNoState:
		return "FSMStateNoState"
	case FSMStateInit:
		return "FSMStateInit"
	case FSMStateStartup:
		return "FSMStateStartup"
	case FSMStateNormal:
		return "FSMStateNormal"
	case FSMStateCommunicationInt:
		return "FSMStateCommunicationInt"
	case FSMStatePartnerDown:
		return "FSMStatePartnerDown"
	case FSMStatePotentialConflict:
		return "FSMStatePotentialConflict"
	case FSMStateConflictDone:
		return "FSMStateConflictDone"
	case FSMStateResolutionInt:
		return "FSMStateResolutionInt"
	case FSMStateRecover:
		return "FSMStateRecover"
	case FSMStateRecoverWait:
		return "FSMStateRecoverWait"
	case FSMStateRecoverDone:
		return "FSMStateRecoverDone"
	}
	return "Invalid"
}

// FailoverRelationship structure represents DHCP_FAILOVER_RELATIONSHIP RPC structure.
//
// The DHCP_FAILOVER_RELATIONSHIP structure defines the information about a DHCPv4 server
// failover relationship.
type FailoverRelationship struct {
	// primaryServer:  This member is of type DHCP_IP_ADDRESS structure (section 2.2.1.2.1)
	// and specifies the IPv4 address of the primary server in the failover relationship.
	PrimaryServer uint32 `idl:"name:primaryServer" json:"primary_server"`
	// secondaryServer:  This member is of type DHCP_IP_ADDRESS structure and specifies
	// the IPv4 address of the secondary server in the failover relationship.
	SecondaryServer uint32 `idl:"name:secondaryServer" json:"secondary_server"`
	// mode:  This member is of type DHCP_FAILOVER_MODE enumeration (section 2.2.1.1.18)
	// and specifies the mode of the failover relationship.
	Mode FailoverMode `idl:"name:mode" json:"mode"`
	// serverType:  This member is of type DHCP_FAILOVER_SERVER enumeration (section 2.2.1.1.19)
	// and specifies the type of failover server.
	ServerType FailoverServer `idl:"name:serverType" json:"server_type"`
	// state:  This member is of type FSM_STATE enumeration (section 2.2.1.1.20) and specifies
	// the state of the failover relationship.
	State FSMState `idl:"name:state" json:"state"`
	// prevState:  This member is of type FSM_STATE enumeration and specifies the previous
	// state of the failover relationship.
	PrevState FSMState `idl:"name:prevState" json:"prev_state"`
	// mclt:  This member is of type DWORD and defines the maximum client lead time (MCLT)
	// of the failover relationship, in seconds.
	MCLT uint32 `idl:"name:mclt" json:"mclt"`
	// safePeriod:  This member is of type DWORD and specifies a safe period time in seconds,
	// that the DHCPv4 server will wait before transitioning the server from the COMMUNICATION-INT
	// state to PARTNER-DOWN state, as described in [IETF-DHCPFOP-12], section 10.
	SafePeriod uint32 `idl:"name:safePeriod" json:"safe_period"`
	// relationshipName:  This member is a pointer of type LPWSTR that points to a null-terminated
	// Unicode string containing the name of the failover relationship that uniquely identifies
	// a failover relationship. There is no restriction on the length of this Unicode string.
	RelationshipName string `idl:"name:relationshipName" json:"relationship_name"`
	// primaryServerName:  This member is a pointer of type LPWSTR that points to a null-terminated
	// Unicode string containing the host name of the primary server in the failover relationship.
	// There is no restriction on the length of this Unicode string.
	PrimaryServerName string `idl:"name:primaryServerName" json:"primary_server_name"`
	// secondaryServerName:  This member is a pointer of type LPWSTR that points to a null-terminated
	// Unicode string containing the host name of the secondary server in the failover relationship.
	// There is no restriction on the length of this Unicode string.
	SecondaryServerName string `idl:"name:secondaryServerName" json:"secondary_server_name"`
	// pScopes:  This member is a pointer of type LPDHCP_IP_ARRAY (section 2.2.1.2.46),
	// which contains the list of IPv4 subnet addresses that are part of the failover relationship.
	Scopes *IPArray `idl:"name:pScopes" json:"scopes"`
	// percentage:  This member is of type BYTE and indicates the ratio of the DHCPv4 client
	// load shared between a primary and secondary server in the failover relationship.
	Percentage uint8 `idl:"name:percentage" json:"percentage"`
	// pSharedSecret:  This member is a pointer of type LPWSTR that points to a null-terminated
	// Unicode string containing the shared secret key associated with this failover relationship.
	// There is no restriction on the length of this string.
	SharedSecret string `idl:"name:pSharedSecret" json:"shared_secret"`
}

func (o *FailoverRelationship) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FailoverRelationship) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PrimaryServer); err != nil {
		return err
	}
	if err := w.WriteData(o.SecondaryServer); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Mode)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ServerType)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.State)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.PrevState)); err != nil {
		return err
	}
	if err := w.WriteData(o.MCLT); err != nil {
		return err
	}
	if err := w.WriteData(o.SafePeriod); err != nil {
		return err
	}
	if o.RelationshipName != "" {
		_ptr_relationshipName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.RelationshipName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.RelationshipName, _ptr_relationshipName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PrimaryServerName != "" {
		_ptr_primaryServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PrimaryServerName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PrimaryServerName, _ptr_primaryServerName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SecondaryServerName != "" {
		_ptr_secondaryServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SecondaryServerName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SecondaryServerName, _ptr_secondaryServerName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Scopes != nil {
		_ptr_pScopes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Scopes != nil {
				if err := o.Scopes.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Scopes, _ptr_pScopes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Percentage); err != nil {
		return err
	}
	if o.SharedSecret != "" {
		_ptr_pSharedSecret := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SharedSecret); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SharedSecret, _ptr_pSharedSecret); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *FailoverRelationship) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrimaryServer); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecondaryServer); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Mode)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ServerType)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.State)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.PrevState)); err != nil {
		return err
	}
	if err := w.ReadData(&o.MCLT); err != nil {
		return err
	}
	if err := w.ReadData(&o.SafePeriod); err != nil {
		return err
	}
	_ptr_relationshipName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.RelationshipName); err != nil {
			return err
		}
		return nil
	})
	_s_relationshipName := func(ptr interface{}) { o.RelationshipName = *ptr.(*string) }
	if err := w.ReadPointer(&o.RelationshipName, _s_relationshipName, _ptr_relationshipName); err != nil {
		return err
	}
	_ptr_primaryServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PrimaryServerName); err != nil {
			return err
		}
		return nil
	})
	_s_primaryServerName := func(ptr interface{}) { o.PrimaryServerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PrimaryServerName, _s_primaryServerName, _ptr_primaryServerName); err != nil {
		return err
	}
	_ptr_secondaryServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SecondaryServerName); err != nil {
			return err
		}
		return nil
	})
	_s_secondaryServerName := func(ptr interface{}) { o.SecondaryServerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.SecondaryServerName, _s_secondaryServerName, _ptr_secondaryServerName); err != nil {
		return err
	}
	_ptr_pScopes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Scopes == nil {
			o.Scopes = &IPArray{}
		}
		if err := o.Scopes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pScopes := func(ptr interface{}) { o.Scopes = *ptr.(**IPArray) }
	if err := w.ReadPointer(&o.Scopes, _s_pScopes, _ptr_pScopes); err != nil {
		return err
	}
	if err := w.ReadData(&o.Percentage); err != nil {
		return err
	}
	_ptr_pSharedSecret := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SharedSecret); err != nil {
			return err
		}
		return nil
	})
	_s_pSharedSecret := func(ptr interface{}) { o.SharedSecret = *ptr.(*string) }
	if err := w.ReadPointer(&o.SharedSecret, _s_pSharedSecret, _ptr_pSharedSecret); err != nil {
		return err
	}
	return nil
}

// FailoverRelationshipArray structure represents DHCP_FAILOVER_RELATIONSHIP_ARRAY RPC structure.
//
// The DHCP_FAILOVER_RELATIONSHIP_ARRAY structure defines an array of DHCP_FAILOVER_RELATIONSHIP
// (section 2.2.1.2.98) structures. This structure is used in the R_DhcpV4FailoverEnumRelationship
// method.
type FailoverRelationshipArray struct {
	// numElements:  This member is of type DWORD and contains the number of DHCP_FAILOVER_RELATIONSHIP
	// elements specified in the subsequent pRelationships field.
	ElementsLength uint32 `idl:"name:numElements" json:"elements_length"`
	// pRelationships:  This member is a pointer to an array of DHCP_FAILOVER_RELATIONSHIP
	// structures of length numElements and contains failover relationship information.
	Relationships []*FailoverRelationship `idl:"name:pRelationships;size_is:(numElements)" json:"relationships"`
}

func (o *FailoverRelationshipArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Relationships != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Relationships))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FailoverRelationshipArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Relationships != nil || o.ElementsLength > 0 {
		_ptr_pRelationships := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Relationships {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Relationships[i1] != nil {
					if err := o.Relationships[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&FailoverRelationship{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Relationships); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&FailoverRelationship{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Relationships, _ptr_pRelationships); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *FailoverRelationshipArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_pRelationships := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Relationships", sizeInfo[0])
		}
		o.Relationships = make([]*FailoverRelationship, sizeInfo[0])
		for i1 := range o.Relationships {
			i1 := i1
			if o.Relationships[i1] == nil {
				o.Relationships[i1] = &FailoverRelationship{}
			}
			if err := o.Relationships[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pRelationships := func(ptr interface{}) { o.Relationships = *ptr.(*[]*FailoverRelationship) }
	if err := w.ReadPointer(&o.Relationships, _s_pRelationships, _ptr_pRelationships); err != nil {
		return err
	}
	return nil
}

// FailoverStatistics structure represents DHCP_FAILOVER_STATISTICS RPC structure.
//
// The DHCP_FAILOVER_STATISTICS structure defines the statistical information for an
// IPv4 subnet configured for a failover relationship.
type FailoverStatistics struct {
	// numAddr:  This member is of type DWORD and contains the total number of IPv4 addresses
	// that can be leased out to DHCPv4 clients on an IPv4 subnet that is part of a failover
	// relationship.
	AddrLength uint32 `idl:"name:numAddr" json:"addr_length"`
	// addrFree:  This member is of type DWORD and contains the total number of IPv4 addresses
	// that are free and can be leased to DHCPv4 clients on an IPv4 subnet that is part
	// of a failover relationship.
	AddrFree uint32 `idl:"name:addrFree" json:"addr_free"`
	// addrInUse:  This member is of type DWORD and contains the total number of IPv4 addresses
	// leased to DHCPv4 clients on an IPv4 subnet that is part of a failover relationship.
	AddrInUse uint32 `idl:"name:addrInUse" json:"addr_in_use"`
	// partnerAddrFree:  This member is of type DWORD and contains the total number of IPv4
	// addresses that are free and can be leased to DHCPv4 clients on an IPv4 subnet that
	// is part of a failover relationship on the partner server.
	PartnerAddrFree uint32 `idl:"name:partnerAddrFree" json:"partner_addr_free"`
	// thisAddrFree:  This member is of type DWORD and contains the total number of IPv4
	// addresses that are free and can be leased to DHCPv4 clients on an IPv4 subnet that
	// is part of a failover relationship on the local DHCP server.
	ThisAddrFree uint32 `idl:"name:thisAddrFree" json:"this_addr_free"`
	// partnerAddrInUse:  This member is of type DWORD and contains the total number of
	// IPv4 addresses leased to DHCPv4 clients on an IPv4 subnet that is part of a failover
	// relationship on the partner server.
	PartnerAddrInUse uint32 `idl:"name:partnerAddrInUse" json:"partner_addr_in_use"`
	// thisAddrInUse:  This member is of type DWORD and contains the total number of IPv4
	// addresses leased to DHCPv4 clients on an IPv4 subnet that is part of a failover relationship
	// on the local DHCP server.
	ThisAddrInUse uint32 `idl:"name:thisAddrInUse" json:"this_addr_in_use"`
}

func (o *FailoverStatistics) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FailoverStatistics) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrLength); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrFree); err != nil {
		return err
	}
	if err := w.WriteData(o.AddrInUse); err != nil {
		return err
	}
	if err := w.WriteData(o.PartnerAddrFree); err != nil {
		return err
	}
	if err := w.WriteData(o.ThisAddrFree); err != nil {
		return err
	}
	if err := w.WriteData(o.PartnerAddrInUse); err != nil {
		return err
	}
	if err := w.WriteData(o.ThisAddrInUse); err != nil {
		return err
	}
	return nil
}
func (o *FailoverStatistics) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrFree); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddrInUse); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartnerAddrFree); err != nil {
		return err
	}
	if err := w.ReadData(&o.ThisAddrFree); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartnerAddrInUse); err != nil {
		return err
	}
	if err := w.ReadData(&o.ThisAddrInUse); err != nil {
		return err
	}
	return nil
}

// FailoverClientInfoV4 structure represents DHCPV4_FAILOVER_CLIENT_INFO RPC structure.
//
// The DHCPV4_FAILOVER_CLIENT_INFO structure defines information about a DHCPv4 client
// leased out by an IPv4 subnet that is a part of failover relationship. This structure
// augments the DHCP_CLIENT_INFO_VQ structure (section 2.2.1.2.19) by including information
// related to DHCP failover and policy-related information. This structure is used by
// the R_DhcpV4FailoverGetClientInfo method specified in section 3.2.4.99.
type FailoverClientInfoV4 struct {
	// ClientIpAddress:  This member is a structure of type DHCP_IP_ADDRESS (section 2.2.1.2.1),
	// which is a DWORD containing the DHCPv4 client's IPv4 address.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// SubnetMask:  This member is a structure of type DHCP_IP_MASK (section 2.2.1.2.2),
	// and is a DWORD containing the DHCPv4 client's IPv4 subnet mask address.
	SubnetMask uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
	// ClientHardwareAddress:  This member is a structure of type DHCP_CLIENT_UID (section
	// 2.2.1.2.5) that represents a DHCPv4 client-identifier (section 2.2.1.2.5.1) or a
	// DHCPv4 client unique ID (section 2.2.1.2.5.2). Methods that accept DHCP_CLIENT_INFO_VQ
	// (section 2.2.1.2.19) as a parameter specify which representations are acceptable.
	ClientHardwareAddress *ClientUID `idl:"name:ClientHardwareAddress" json:"client_hardware_address"`
	// ClientName:  This member is a pointer to a null-terminated Unicode string that represents
	// the DHCPv4 client's internet host name. There is no restriction on the length of
	// this Unicode string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
	// ClientComment:  This member is a pointer to a null-terminated Unicode string that
	// represents the description given to the DHCPv4 client. There is no restriction on
	// the length of this Unicode string.
	ClientComment string `idl:"name:ClientComment" json:"client_comment"`
	// ClientLeaseExpires:  This member is a structure of type DATE_TIME (section 2.2.1.2.11)
	// and contains the lease expiry time for the DHCPv4 client. This is UTC time represented
	// in FILETIME format.
	ClientLeaseExpires *DateTime `idl:"name:ClientLeaseExpires" json:"client_lease_expires"`
	// OwnerHost:  This member is a structure of type DHCP_HOST_INFO (section 2.2.1.2.7)
	// that contains information about the DHCPv4 Server machine that has provided a lease
	// to the DHCPv4 client.
	OwnerHost *HostInfo `idl:"name:OwnerHost" json:"owner_host"`
	// bClientType:  This member is of type BYTE and identifies the type of the DHCPv4 client.
	// Possible values are specified in the following table.
	//
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	|                                   |                                                                       |
	//	|               VALUE               |                                MEANING                                |
	//	|                                   |                                                                       |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_UNSPECIFIED 0x00      | A DHCPv4 client other than ones defined in this table.                |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_DHCP 0x01             | The DHCPv4 client supports the DHCP protocol.                         |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOOTP 0x02            | The DHCPv4 client supports the BOOTP protocol ([RFC2132]).            |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOTH 0x03             | The DHCPv4 client identifies both the DHCPv4 and the BOOTP protocols. |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_RESERVATION_FLAG 0x04 | There is an IPv4 reservation created for the DHCPv4 client.           |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	//	| CLIENT_TYPE_NONE 0x64             | Backward compatibility for manual addressing.                         |
	//	+-----------------------------------+-----------------------------------------------------------------------+
	ClientType uint8 `idl:"name:bClientType" json:"client_type"`
	// AddressState:  This member is of type BYTE, as shown by the following set of bits.
	// This member represents the state of the IPv4 address given to the DHCPv4 client.
	//
	//	+-------+-------+-------+-------+-------+-------+-------+-------+
	//	|  BIT  |  BIT  |  BIT  |  BIT  |  BIT  |  BIT  |  BIT  |  BIT  |
	//	|   7   |   6   |   5   |   4   |   3   |   2   |   1   |   0   |
	//	+-------+-------+-------+-------+-------+-------+-------+-------+
	//	+-------+-------+-------+-------+-------+-------+-------+-------+
	//
	// The following tables show various bit representation values and their meanings.
	//
	// BIT 0 and BIT 1 signify the state of the leased IPv4 address, as shown in the following
	// table.
	//
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	|                            |                                                                                  |
	//	|           VALUE            |                                     MEANING                                      |
	//	|                            |                                                                                  |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_OFFERED 0x0  | The DHCPv4 client is offered this IPv4 address.                                  |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_ACTIVE 0x1   | The IPv4 address is active and has an active DHCPv4 client lease record.         |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DECLINED 0x2 | The IPv4 address request is declined by the DHCPv4 client; hence, it is a bad    |
	//	|                            | IPv4 address.                                                                    |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DOOM 0x3     | The IPv4 address is in DOOMED state and is due to be deleted.                    |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//
	// BIT 2 and BIT 3 signify information related to Name Protection (section 3.3.3) for
	// the leased IPv4 address, as shown in the following table.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                           |                                                                                  |
	//	|                   VALUE                   |                                     MEANING                                      |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_NO_DHCID 0x0                  | The address is leased to the DHCPv4 client without DHCID (sections 3 and 3.5 of  |
	//	|                                           | [RFC4701]).                                                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DHCID_NO_CLIENTIDOPTION 0x1   | The address is leased to the DHCPv4 client with DHCID.                           |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DHCID_WITH_CLIENTIDOPTION 0x2 | The address is leased to the DHCPv4 client with DHCID.                           |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DHCID_WITH_DUID 0x3           | The address is leased to the DHCPv4 client with DHCID.                           |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// BIT 4, BIT 5, BIT 6, and BIT 7 specify information related to DNS, as shown in the
	// following table.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_CLEANUP 0x1      | The DNS update for the DHCPv4 client lease record needs to be deleted from the   |
	//	|                              | DNS server when the lease is deleted.                                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_BOTH_REC 0x2     | The DNS update needs to be sent for both A and PTR resource records ([RFC1034]   |
	//	|                              | section 3.6).                                                                    |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_UNREGISTERED 0x4 | The DNS update is not completed for the lease record.                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_BIT_DELETED 0x8      | The address lease is expired, but the DNS updates for the lease record have not  |
	//	|                              | been deleted from the DNS server.                                                |
	//	+------------------------------+----------------------------------------------------------------------------------+
	AddressState uint8 `idl:"name:AddressState" json:"address_state"`
	// Status:  This member is an enumeration of type QuarantineStatus (section 2.2.1.1.11)
	// that contains the health status of the DHCPv4 client, as validated by the NAP server.
	// The possible enumeration values that are validated by the NAP server are NOQUARANTINE,
	// RESTRICTEDACCESS, DROPPACKET, and PROBATION.
	Status QuarantineStatus `idl:"name:Status" json:"status"`
	// ProbationEnds:  This member is a structure of type DATE_TIME (section 2.2.1.2.11)
	// that contains probation end time, assuming the DHCPv4 client is on probation. For
	// this time period, the DHCPv4 client has full access to the network.
	ProbationEnds *DateTime `idl:"name:ProbationEnds" json:"probation_ends"`
	// QuarantineCapable:  This member is of type BOOL and takes on one of the meanings
	// associated with the values in the following table.
	//
	//	+---------+------------------------------------------------------+
	//	|         |                                                      |
	//	|  VALUE  |                       MEANING                        |
	//	|         |                                                      |
	//	+---------+------------------------------------------------------+
	//	+---------+------------------------------------------------------+
	//	| TRUE 1  | The DHCPv4 client machine is quarantine-enabled.     |
	//	+---------+------------------------------------------------------+
	//	| FALSE 0 | The DHCPv4 client machine is not quarantine-enabled. |
	//	+---------+------------------------------------------------------+
	QuarantineCapable bool `idl:"name:QuarantineCapable" json:"quarantine_capable"`
	// SentPotExpTime:  This member is of type DWORD and contains the time sent to the partner
	// server as potential-expiration-time. The time is specified in seconds elapsed since
	// midnight, January 1, 1970, Coordinated Universal Time (UTC).
	SentPotentialExpirationTime uint32 `idl:"name:SentPotExpTime" json:"sent_potential_expiration_time"`
	// AckPotExpTime:  This member is of type DWORD and contains the time that the partner
	// server has acknowledged as potential-expiration-time. The time is specified in seconds
	// elapsed since midnight, January 1, 1970, UTC.
	AckPotentialExpirationTime uint32 `idl:"name:AckPotExpTime" json:"ack_potential_expiration_time"`
	// RecvPotExpTime:  This member is of type DWORD and contains the time that the server
	// has received as a potential-expiration-time from its partner server. The time is
	// specified in seconds elapsed since midnight, January 1, 1970, UTC.
	RecvPotentialExpirationTime uint32 `idl:"name:RecvPotExpTime" json:"recv_potential_expiration_time"`
	// StartTime:  This member is of type DWORD and contains the time at which the client
	// lease first went into the current state. The time is specified in seconds elapsed
	// since midnight, January 1, 1970, UTC.
	StartTime uint32 `idl:"name:StartTime" json:"start_time"`
	// CltLastTransTime:  This member is of type DWORD and contains the time for client-last-transaction-time.
	// The time is specified in seconds elapsed since midnight, January 1, 1970, UTC.
	ClientLastTransTime uint32 `idl:"name:CltLastTransTime" json:"client_last_trans_time"`
	// LastBndUpdTime:  This member is of type DWORD and contains the time when the partner
	// server has updated the DHCPv4 client lease. The time is specified in seconds elapsed
	// since midnight, January 1, 1970, UTC.
	LastBindingUpdateTime uint32 `idl:"name:LastBndUpdTime" json:"last_binding_update_time"`
	// bndMsgStatus:  This member is of type DWORD and MUST be ignored.
	BindingMessageStatus uint32 `idl:"name:bndMsgStatus" json:"binding_message_status"`
	// PolicyName:  This member is a pointer to a null-terminated Unicode string that identifies
	// the policy that determined the ClientIpAddress in the lease. The length of the Policy
	// Name member is restricted to 64 characters.
	PolicyName string `idl:"name:PolicyName" json:"policy_name"`
	// flags:  This member is of type BYTE and MUST be ignored.
	Flags uint8 `idl:"name:flags" json:"flags"`
}

func (o *FailoverClientInfoV4) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FailoverClientInfoV4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress != nil {
		if err := o.ClientHardwareAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientComment != "" {
		_ptr_ClientComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientComment, _ptr_ClientComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientLeaseExpires != nil {
		if err := o.ClientLeaseExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OwnerHost != nil {
		if err := o.OwnerHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ClientType); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressState); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if o.ProbationEnds != nil {
		if err := o.ProbationEnds.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.QuarantineCapable {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SentPotentialExpirationTime); err != nil {
		return err
	}
	if err := w.WriteData(o.AckPotentialExpirationTime); err != nil {
		return err
	}
	if err := w.WriteData(o.RecvPotentialExpirationTime); err != nil {
		return err
	}
	if err := w.WriteData(o.StartTime); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientLastTransTime); err != nil {
		return err
	}
	if err := w.WriteData(o.LastBindingUpdateTime); err != nil {
		return err
	}
	if err := w.WriteData(o.BindingMessageStatus); err != nil {
		return err
	}
	if o.PolicyName != "" {
		_ptr_PolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PolicyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyName, _ptr_PolicyName); err != nil {
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
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *FailoverClientInfoV4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress == nil {
		o.ClientHardwareAddress = &ClientUID{}
	}
	if err := o.ClientHardwareAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	_ptr_ClientComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClientComment := func(ptr interface{}) { o.ClientComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientComment, _s_ClientComment, _ptr_ClientComment); err != nil {
		return err
	}
	if o.ClientLeaseExpires == nil {
		o.ClientLeaseExpires = &DateTime{}
	}
	if err := o.ClientLeaseExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OwnerHost == nil {
		o.OwnerHost = &HostInfo{}
	}
	if err := o.OwnerHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressState); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	if o.ProbationEnds == nil {
		o.ProbationEnds = &DateTime{}
	}
	if err := o.ProbationEnds.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bQuarantineCapable int32
	if err := w.ReadData(&_bQuarantineCapable); err != nil {
		return err
	}
	o.QuarantineCapable = _bQuarantineCapable != 0
	if err := w.ReadData(&o.SentPotentialExpirationTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.AckPotentialExpirationTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.RecvPotentialExpirationTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.StartTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientLastTransTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.LastBindingUpdateTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.BindingMessageStatus); err != nil {
		return err
	}
	_ptr_PolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PolicyName); err != nil {
			return err
		}
		return nil
	})
	_s_PolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PolicyName, _s_PolicyName, _ptr_PolicyName); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// IPReservationInfo structure represents DHCP_IP_RESERVATION_INFO RPC structure.
//
// The DHCP_IP_RESERVATION_INFO structure defines an IPv4 reservation for a DHCPv4 client.
// This structure is an extension of the DHCP_IP_RESERVATION_V4 structure (section 2.2.1.2.32),
// which is extended by including the reservation client name and description.
type IPReservationInfo struct {
	// ReservedIpAddress:  This member is a structure of type DHCP_IP_ADDRESS (section 2.2.1.2.1)
	// that contains the IPv4 address of the client (DHCP or BOOTP) ([RFC2132]), for which
	// a reservation was created.
	ReservedIPAddress uint32 `idl:"name:ReservedIpAddress" json:"reserved_ip_address"`
	// ReservedForClient:  This member is a pointer of type DHCP_CLIENT_UID structure (section
	// 2.2.1.2.5) that represents the DHCPv4 client-identifier (section 2.2.1.2.5.1).
	ReservedForClient *ClientUID `idl:"name:ReservedForClient" json:"reserved_for_client"`
	// ReservedClientName:  This member is a pointer to a null-terminated Unicode string
	// that represents the host name of the DHCPv4 reservation client. There is no restriction
	// on the length of this Unicode string.
	ReservedClientName string `idl:"name:ReservedClientName" json:"reserved_client_name"`
	// ReservedClientDesc:  This member is a pointer to a null-terminated Unicode string
	// that represents the description of the DHCPv4 reservation client. There is no restriction
	// on the length of this Unicode string.
	ReservedClientDesc string `idl:"name:ReservedClientDesc" json:"reserved_client_desc"`
	// bAllowedClientTypes:  This member is of type BYTE and specifies the type of client
	// holding this reservation, as indicated in the following table.
	//
	//	+------------------------+---------------------------------------------------------+
	//	|                        |                                                         |
	//	|         VALUE          |                         MEANING                         |
	//	|                        |                                                         |
	//	+------------------------+---------------------------------------------------------+
	//	+------------------------+---------------------------------------------------------+
	//	| CLIENT_TYPE_DHCP 0x01  | The IPv4 reservation is for a DHCPv4 client.            |
	//	+------------------------+---------------------------------------------------------+
	//	| CLIENT_TYPE_BOOTP 0x02 | The IPv4 reservation is for a BOOTP client ([RFC2132]). |
	//	+------------------------+---------------------------------------------------------+
	//	| CLIENT_TYPE_BOTH 0x03  | The IPv4 reservation is for both kinds of clients.      |
	//	+------------------------+---------------------------------------------------------+
	AllowedClientTypes uint8 `idl:"name:bAllowedClientTypes" json:"allowed_client_types"`
	// fOptionsPresent:  This member is of type BYTE and specifies whether there are any
	// DHCPv4 options configured on the reservation, as indicated in the following table.
	//
	//	+------------+-----------------------------------------------------+
	//	|            |                                                     |
	//	|   VALUE    |                       MEANING                       |
	//	|            |                                                     |
	//	+------------+-----------------------------------------------------+
	//	+------------+-----------------------------------------------------+
	//	| 0x00000000 | No option values are configured on the reservation. |
	//	+------------+-----------------------------------------------------+
	//	| 0x00000001 | Option values are configured on the reservation.    |
	//	+------------+-----------------------------------------------------+
	OptionsPresent uint8 `idl:"name:fOptionsPresent" json:"options_present"`
}

func (o *IPReservationInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPReservationInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ReservedIPAddress); err != nil {
		return err
	}
	if o.ReservedForClient != nil {
		if err := o.ReservedForClient.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ReservedClientName != "" {
		_ptr_ReservedClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ReservedClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ReservedClientName, _ptr_ReservedClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ReservedClientDesc != "" {
		_ptr_ReservedClientDesc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ReservedClientDesc); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ReservedClientDesc, _ptr_ReservedClientDesc); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AllowedClientTypes); err != nil {
		return err
	}
	if err := w.WriteData(o.OptionsPresent); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *IPReservationInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReservedIPAddress); err != nil {
		return err
	}
	if o.ReservedForClient == nil {
		o.ReservedForClient = &ClientUID{}
	}
	if err := o.ReservedForClient.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ReservedClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ReservedClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ReservedClientName := func(ptr interface{}) { o.ReservedClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ReservedClientName, _s_ReservedClientName, _ptr_ReservedClientName); err != nil {
		return err
	}
	_ptr_ReservedClientDesc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ReservedClientDesc); err != nil {
			return err
		}
		return nil
	})
	_s_ReservedClientDesc := func(ptr interface{}) { o.ReservedClientDesc = *ptr.(*string) }
	if err := w.ReadPointer(&o.ReservedClientDesc, _s_ReservedClientDesc, _ptr_ReservedClientDesc); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllowedClientTypes); err != nil {
		return err
	}
	if err := w.ReadData(&o.OptionsPresent); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ReservationInfoArray structure represents DHCP_RESERVATION_INFO_ARRAY RPC structure.
//
// The DHCP_RESERVATION_INFO_ARRAY structure defines an array of DHCP_IP RESERVATION_INFO
// (section 2.2.1.2.102) structures. This structure is used by the method R_DhcpV4EnumSubnetReservations
// (section 3.2.4.120).
type ReservationInfoArray struct {
	// NumElements:  This member is of type DWORD and contains the number of DHCP_IP_RESERVATION_INFO
	// elements specified by the subsequent Elements member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This member is a pointer to an array of DHCP_IP_RESERVATION_INFO structures
	// of length NumElements, and contains DHCPv4 reservation information.
	Elements []*IPReservationInfo `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *ReservationInfoArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReservationInfoArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Elements[i1] != nil {
							if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&IPReservationInfo{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Elements[i1], _ptr_Elements); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReservationInfoArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*IPReservationInfo, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Elements[i1] == nil {
					o.Elements[i1] = &IPReservationInfo{}
				}
				if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Elements := func(ptr interface{}) { o.Elements[i1] = *ptr.(**IPReservationInfo) }
			if err := w.ReadPointer(&o.Elements[i1], _s_Elements, _ptr_Elements); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*IPReservationInfo) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// AllOptionValuesPB structure represents DHCP_ALL_OPTION_VALUES_PB RPC structure.
//
// The DHCP_ALL_OPTION_VALUES_PB structure contains all the option values set for a
// specific policy.
type AllOptionValuesPB struct {
	// Flags:  This member is of type DWORD. It is an unused field that MUST be initialized
	// to 0 in any RPC method that modifies the DHCP server configuration. This MUST be
	// treated as an error if it is nonzero in an RPC method that queries the DHCP server
	// configuration.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// NumElements:  This member is of type DWORD and specifies the number of option structures
	// in the DHCP_ALL_OPTION_VALUES_PB structure.
	ElementsLength uint32                       `idl:"name:NumElements" json:"elements_length"`
	Options        []*AllOptionValuesPB_Options `idl:"name:Options;size_is:(NumElements)" json:"options"`
}

func (o *AllOptionValuesPB) xxx_PreparePayload(ctx context.Context) error {
	if o.Options != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Options))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptionValuesPB) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Options != nil || o.ElementsLength > 0 {
		_ptr_Options := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Options {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Options[i1] != nil {
					if err := o.Options[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&AllOptionValuesPB_Options{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Options); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&AllOptionValuesPB_Options{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Options, _ptr_Options); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptionValuesPB) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Options := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Options", sizeInfo[0])
		}
		o.Options = make([]*AllOptionValuesPB_Options, sizeInfo[0])
		for i1 := range o.Options {
			i1 := i1
			if o.Options[i1] == nil {
				o.Options[i1] = &AllOptionValuesPB_Options{}
			}
			if err := o.Options[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Options := func(ptr interface{}) { o.Options = *ptr.(*[]*AllOptionValuesPB_Options) }
	if err := w.ReadPointer(&o.Options, _s_Options, _ptr_Options); err != nil {
		return err
	}
	return nil
}

// AllOptionValuesPB_Options structure represents DHCP_ALL_OPTION_VALUES_PB structure anonymous member.
//
// The DHCP_ALL_OPTION_VALUES_PB structure contains all the option values set for a
// specific policy.
type AllOptionValuesPB_Options struct {
	// PolicyName:  This member is a pointer of type LPWSTR and contains the null-terminated
	// Unicode string identifying the name of the policy. The name of the policy is restricted
	// to 64 characters.
	PolicyName string `idl:"name:PolicyName" json:"policy_name"`
	// VendorName:  This member is a pointer of type LPWSTR and contains the vendor class
	// name. This field is unused.
	VendorName string `idl:"name:VendorName" json:"vendor_name"`
	// IsVendor:  This member is of type BOOL and specifies whether this is a vendor class
	// option.
	IsVendor bool `idl:"name:IsVendor" json:"is_vendor"`
	// OptionsArray:  This member is a pointer to the DHCP_OPTION_VALUE_ARRAY structure
	// (section 2.2.1.2.43) and contains the option values set for the policy.
	OptionsArray *OptionValueArray `idl:"name:OptionsArray" json:"options_array"`
}

func (o *AllOptionValuesPB_Options) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptionValuesPB_Options) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.PolicyName != "" {
		_ptr_PolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PolicyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyName, _ptr_PolicyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.VendorName != "" {
		_ptr_VendorName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.VendorName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VendorName, _ptr_VendorName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.IsVendor {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.OptionsArray != nil {
		_ptr_OptionsArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.OptionsArray != nil {
				if err := o.OptionsArray.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&OptionValueArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OptionsArray, _ptr_OptionsArray); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AllOptionValuesPB_Options) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_PolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PolicyName); err != nil {
			return err
		}
		return nil
	})
	_s_PolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PolicyName, _s_PolicyName, _ptr_PolicyName); err != nil {
		return err
	}
	_ptr_VendorName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.VendorName); err != nil {
			return err
		}
		return nil
	})
	_s_VendorName := func(ptr interface{}) { o.VendorName = *ptr.(*string) }
	if err := w.ReadPointer(&o.VendorName, _s_VendorName, _ptr_VendorName); err != nil {
		return err
	}
	var _bIsVendor int32
	if err := w.ReadData(&_bIsVendor); err != nil {
		return err
	}
	o.IsVendor = _bIsVendor != 0
	_ptr_OptionsArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.OptionsArray == nil {
			o.OptionsArray = &OptionValueArray{}
		}
		if err := o.OptionsArray.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_OptionsArray := func(ptr interface{}) { o.OptionsArray = *ptr.(**OptionValueArray) }
	if err := w.ReadPointer(&o.OptionsArray, _s_OptionsArray, _ptr_OptionsArray); err != nil {
		return err
	}
	return nil
}

// PolicyAttributeType type represents DHCP_POL_ATTR_TYPE RPC enumeration.
//
// The DHCP_POL_ATTR_TYPE enumeration specifies the attribute type for a condition of
// a policy.<10>
type PolicyAttributeType uint16

var (
	// DhcpAttrHWAddr: Specifies that the condition is based on hardware address.
	PolicyAttributeTypeHwAddr PolicyAttributeType = 0
	// DhcpAttrOption: Specifies that the condition is based on a DHCP option.
	PolicyAttributeTypeOption PolicyAttributeType = 1
	// DhcpAttrSubOption: Specifies that the condition is based on a DHCP suboption.
	PolicyAttributeTypeSubOption PolicyAttributeType = 2
	// DhcpAttrFqdn: Specifies that the condition is based on the fully qualified domain
	// name of the client.
	PolicyAttributeTypeFQDN PolicyAttributeType = 3
	// DhcpAttrFqdnSingleLabel: Specifies whether the condition is based on the DHCP client
	// being a single-label machine.
	PolicyAttributeTypeFQDNSingleLabel PolicyAttributeType = 4
)

func (o PolicyAttributeType) String() string {
	switch o {
	case PolicyAttributeTypeHwAddr:
		return "PolicyAttributeTypeHwAddr"
	case PolicyAttributeTypeOption:
		return "PolicyAttributeTypeOption"
	case PolicyAttributeTypeSubOption:
		return "PolicyAttributeTypeSubOption"
	case PolicyAttributeTypeFQDN:
		return "PolicyAttributeTypeFQDN"
	case PolicyAttributeTypeFQDNSingleLabel:
		return "PolicyAttributeTypeFQDNSingleLabel"
	}
	return "Invalid"
}

// PolicyComparator type represents DHCP_POL_COMPARATOR RPC enumeration.
//
// The DHCP_POL_COMPARATOR enumeration defines the different comparison operators used
// within a condition of a policy.<9>
type PolicyComparator uint16

var (
	// DhcpCompEqual: Specifies the equal operator for a condition.
	PolicyComparatorEqual PolicyComparator = 0
	// DhcpCompNotEqual: Specifies the not-equal operator for a condition.
	PolicyComparatorNotEqual PolicyComparator = 1
	// DhcpCompBeginsWith: Specifies the begins-with operator for a condition.
	PolicyComparatorBeginsWith PolicyComparator = 2
	// DhcpCompNotBeginWith: Specifies the does-not-begin-with operator for a condition.
	PolicyComparatorNotBeginWith PolicyComparator = 3
	// DhcpCompEndsWith: Specifies the ends-with operator for a condition.
	PolicyComparatorEndsWith PolicyComparator = 4
	// DhcpCompNotEndWith: Specifies the does-not-end-with operator for a condition.
	PolicyComparatorNotEndWith PolicyComparator = 5
)

func (o PolicyComparator) String() string {
	switch o {
	case PolicyComparatorEqual:
		return "PolicyComparatorEqual"
	case PolicyComparatorNotEqual:
		return "PolicyComparatorNotEqual"
	case PolicyComparatorBeginsWith:
		return "PolicyComparatorBeginsWith"
	case PolicyComparatorNotBeginWith:
		return "PolicyComparatorNotBeginWith"
	case PolicyComparatorEndsWith:
		return "PolicyComparatorEndsWith"
	case PolicyComparatorNotEndWith:
		return "PolicyComparatorNotEndWith"
	}
	return "Invalid"
}

// PolicyLogicOperator type represents DHCP_POL_LOGIC_OPER RPC enumeration.
//
// The DHCP_POL_LOGIC_OPER enumeration contains the logical operator used to combine
// the conditions of a policy.
type PolicyLogicOperator uint16

var (
	// DhcpLogicalOr: Specifies the OR logical operator.
	PolicyLogicOperatorLogicalOr PolicyLogicOperator = 0
	// DhcpLogicalAnd: Specifies the AND logical operator.
	PolicyLogicOperatorLogicalAnd PolicyLogicOperator = 1
)

func (o PolicyLogicOperator) String() string {
	switch o {
	case PolicyLogicOperatorLogicalOr:
		return "PolicyLogicOperatorLogicalOr"
	case PolicyLogicOperatorLogicalAnd:
		return "PolicyLogicOperatorLogicalAnd"
	}
	return "Invalid"
}

// PolicyFieldsToUpdate type represents DHCP_POLICY_FIELDS_TO_UPDATE RPC enumeration.
//
// The DHCP_POLICY_FIELDS_TO_UPDATE enumeration defines the policy fields to be updated
// during a set operation using the R_Dhcpv4SetPolicy method specified in section 3.2.4.111.
type PolicyFieldsToUpdate uint16

var (
	// DhcpUpdatePolicyName:  Updates the name of the policy.
	PolicyFieldsToUpdatePolicyName PolicyFieldsToUpdate = 1
	// DhcpUpdatePolicyOrder: Updates the processing order of the policy.
	PolicyFieldsToUpdatePolicyOrder PolicyFieldsToUpdate = 2
	// DhcpUpdatePolicyExpr: Updates the expressions and conditions of the policy.
	PolicyFieldsToUpdatePolicyExpr PolicyFieldsToUpdate = 4
	// DhcpUpdatePolicyRanges: Updates the IP ranges of the policy.
	PolicyFieldsToUpdatePolicyRanges PolicyFieldsToUpdate = 8
	// DhcpUpdatePolicyDescr: Updates the description of the policy.
	PolicyFieldsToUpdatePolicyDescription PolicyFieldsToUpdate = 16
	// DhcpUpdatePolicyStatus: Updates the state (enabled/disabled) of the policy.
	PolicyFieldsToUpdatePolicyStatus PolicyFieldsToUpdate = 32
	// DhcpUpdatePolicyDnsSuffix: Updates the DNS suffix for the policy.<8>
	PolicyFieldsToUpdatePolicyDNSSuffix PolicyFieldsToUpdate = 64
)

func (o PolicyFieldsToUpdate) String() string {
	switch o {
	case PolicyFieldsToUpdatePolicyName:
		return "PolicyFieldsToUpdatePolicyName"
	case PolicyFieldsToUpdatePolicyOrder:
		return "PolicyFieldsToUpdatePolicyOrder"
	case PolicyFieldsToUpdatePolicyExpr:
		return "PolicyFieldsToUpdatePolicyExpr"
	case PolicyFieldsToUpdatePolicyRanges:
		return "PolicyFieldsToUpdatePolicyRanges"
	case PolicyFieldsToUpdatePolicyDescription:
		return "PolicyFieldsToUpdatePolicyDescription"
	case PolicyFieldsToUpdatePolicyStatus:
		return "PolicyFieldsToUpdatePolicyStatus"
	case PolicyFieldsToUpdatePolicyDNSSuffix:
		return "PolicyFieldsToUpdatePolicyDNSSuffix"
	}
	return "Invalid"
}

// PolicyCondition structure represents DHCP_POL_COND RPC structure.
//
// The DHCP_POL_COND structure specifies an individual condition of a policy.
type PolicyCondition struct {
	// ParentExpr:  This member is of type DWORD and contains the index of the parent expression
	// in the DHCP_POL_EXPR_ARRAY (section 2.2.1.2.108) structure of the same policy.
	ParentExpr uint32 `idl:"name:ParentExpr" json:"parent_expr"`
	// Type:  This member is of type DHCP_POL_ATTR_TYPE enumeration (section 2.2.1.1.23)
	// and identifies whether the condition is specified for an option, suboption, or hardware
	// address.
	Type PolicyAttributeType `idl:"name:Type" json:"type"`
	// OptionID:  This member is of type DWORD and contains the identifier for the DHCP
	// option if the Type member contains the DhcpAttrOption enumeration value.
	OptionID uint32 `idl:"name:OptionID" json:"option_id"`
	// SubOptionID:  This member is of type DWORD and contains the identifier for the DHCP
	// suboption of the option contained in the OptionID member, providing that the Type
	// member contains the DhcpAttrSubOption enumeration value.
	SubOptionID uint32 `idl:"name:SubOptionID" json:"sub_option_id"`
	// VendorName:  This member is a pointer of type LPWSTR that points to a NULL terminated
	// Unicode string containing the name of a vendor class. This member identifies the
	// vendor class to which the OptionID or SubOptionID belongs, in case of a vendor-specific
	// option/suboption being specified in the condition. This field is currently unused.
	VendorName string `idl:"name:VendorName" json:"vendor_name"`
	// Operator:  This member is of type DHCP_POL_COMPARATOR enumeration (section 2.2.1.1.22)
	// and specifies the comparison operator for the condition.
	Operator PolicyComparator `idl:"name:Operator" json:"operator"`
	// Value:  This member is of type LPBYTE and points to an array of bytes containing
	// the value to be used for the comparison.
	Value []byte `idl:"name:Value;size_is:(ValueLength)" json:"value"`
	// ValueLength:  This member is of type DWORD and specifies the length of the Value
	// member.
	ValueLength uint32 `idl:"name:ValueLength" json:"value_length"`
}

func (o *PolicyCondition) xxx_PreparePayload(ctx context.Context) error {
	if o.Value != nil && o.ValueLength == 0 {
		o.ValueLength = uint32(len(o.Value))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyCondition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ParentExpr); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(o.OptionID); err != nil {
		return err
	}
	if err := w.WriteData(o.SubOptionID); err != nil {
		return err
	}
	if o.VendorName != "" {
		_ptr_VendorName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.VendorName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VendorName, _ptr_VendorName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.Operator)); err != nil {
		return err
	}
	if o.Value != nil || o.ValueLength > 0 {
		_ptr_Value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValueLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Value {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Value[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Value); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Value, _ptr_Value); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ValueLength); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *PolicyCondition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ParentExpr); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData(&o.OptionID); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubOptionID); err != nil {
		return err
	}
	_ptr_VendorName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.VendorName); err != nil {
			return err
		}
		return nil
	})
	_s_VendorName := func(ptr interface{}) { o.VendorName = *ptr.(*string) }
	if err := w.ReadPointer(&o.VendorName, _s_VendorName, _ptr_VendorName); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Operator)); err != nil {
		return err
	}
	_ptr_Value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			return fmt.Errorf("buffer overflow for size %d of array o.Value", sizeInfo[0])
		}
		o.Value = make([]byte, sizeInfo[0])
		for i1 := range o.Value {
			i1 := i1
			if err := w.ReadData(&o.Value[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Value := func(ptr interface{}) { o.Value = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Value, _s_Value, _ptr_Value); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueLength); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// PolicyConditionArray structure represents DHCP_POL_COND_ARRAY RPC structure.
//
// The DHCP_POL_COND_ARRAY structure specifies an array of conditions of a policy.
type PolicyConditionArray struct {
	// NumElements:  This member is of type DWORD and specifies the number of conditions
	// in the array.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This member is a pointer of type LPDHCP_POL_COND (section 2.2.1.2.105)
	// that points to an array of DHCP_POL_COND structure (section 2.2.1.2.105) elements.
	Elements []*PolicyCondition `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *PolicyConditionArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyConditionArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PolicyCondition{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&PolicyCondition{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyConditionArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*PolicyCondition, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &PolicyCondition{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*PolicyCondition) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// PolicyExpr structure represents DHCP_POL_EXPR RPC structure.
//
// The DHCP_POL_EXPR structure specifies an individual expression of a policy.
type PolicyExpr struct {
	// ParentExpr:  This member is of type DWORD and contains the index of the parent expression
	// in the DHCP_POL_EXPR_ARRAY structure (section 2.2.1.2.108) associated with the policy.
	ParentExpr uint32 `idl:"name:ParentExpr" json:"parent_expr"`
	// Operator:  This member is of type DHCP_POL_LOGIC_OPER enumeration (section 2.2.1.1.24)
	// and specifies the logical operator of this expression.
	Operator PolicyLogicOperator `idl:"name:Operator" json:"operator"`
}

func (o *PolicyExpr) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyExpr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ParentExpr); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Operator)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *PolicyExpr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ParentExpr); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Operator)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// PolicyExprArray structure represents DHCP_POL_EXPR_ARRAY RPC structure.
//
// The DHCP_POL_EXPR_ARRAY structure specifies the array of expressions of a policy.
type PolicyExprArray struct {
	// NumElements:  This member is of type DWORD and contains the number of expression
	// elements in the array.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This member is a pointer of type LPDHCP_POL_EXPR (section 2.2.1.2.107)
	// that points to an array of DHCP_POL_EXPR (section 2.2.1.2.107) elements.
	Elements []*PolicyExpr `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *PolicyExprArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyExprArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PolicyExpr{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&PolicyExpr{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyExprArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*PolicyExpr, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &PolicyExpr{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*PolicyExpr) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// IPRangeArray structure represents DHCP_IP_RANGE_ARRAY RPC structure.
//
// The DHCP_IP_RANGE_ARRAY structure specifies an array of IP address ranges.
type IPRangeArray struct {
	// NumElements:  This member contains the number of IP address range elements in the
	// array.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This member is a pointer to an array of type DHCP_IP_RANGE structure (section
	// 2.2.1.2.31) elements.
	Elements []*IPRange `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *IPRangeArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPRangeArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&IPRange{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&IPRange{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPRangeArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*IPRange, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &IPRange{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*IPRange) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// Policy structure represents DHCP_POLICY RPC structure.
//
// The DHCP_POLICY structure contains information for a policy used to filter client
// requests.
type Policy struct {
	// PolicyName:  This member is a pointer of type LPWSTR that points to a null-terminated
	// Unicode string identifying the name of the policy. The name of the policy is restricted
	// to 64 characters.
	PolicyName string `idl:"name:PolicyName" json:"policy_name"`
	// IsGlobalPolicy:  This member is of type BOOL and indicates whether this is a server
	// level policy.
	IsGlobalPolicy bool `idl:"name:IsGlobalPolicy" json:"is_global_policy"`
	// Subnet:  This member is of type DHCP_IP_ADDRESS structure (section 2.2.1.2.1) and
	// identifies the IPv4 subnet to which the policy belongs, if this is a scope level
	// policy. The value of this member will be 0 for a server level policy.
	Subnet uint32 `idl:"name:Subnet" json:"subnet"`
	// ProcessingOrder:  This member is of type DWORD and identifies the relative order
	// in which the DHCPv4 server will process the policy.
	ProcessingOrder uint32 `idl:"name:ProcessingOrder" json:"processing_order"`
	// Conditions:  This member is a pointer of type LPDHCP_POL_COND_ARRAY (section 2.2.1.2.106)
	// that contains the array of conditions for the policy.
	Conditions *PolicyConditionArray `idl:"name:Conditions" json:"conditions"`
	// Expressions:  This member is a pointer of type LPDHCP_POL_EXPR_ARRAY (section 2.2.1.2.108)
	// that contains the array of expressions for the policy.
	Expressions *PolicyExprArray `idl:"name:Expressions" json:"expressions"`
	// Ranges:  This member is a pointer of type LPDHCP_IP_RANGE_ARRAY (section 2.2.1.2.104)
	// which points to an array of DHCP_IP_RANGE structures (section 2.2.1.2.31) that represent
	// the policy IP ranges.
	Ranges *IPRangeArray `idl:"name:Ranges" json:"ranges"`
	// Description:  This member is a pointer of type LPWSTR and contains the null-terminated
	// Unicode string with the description of the policy. This description string is restricted
	// to 255 characters.
	Description string `idl:"name:Description" json:"description"`
	// Enabled:  This member is a flag of type BOOL that indicates whether the policy is
	// in the enabled or disabled state.
	Enabled bool `idl:"name:Enabled" json:"enabled"`
}

func (o *Policy) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Policy) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.PolicyName != "" {
		_ptr_PolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PolicyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyName, _ptr_PolicyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.IsGlobalPolicy {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Subnet); err != nil {
		return err
	}
	if err := w.WriteData(o.ProcessingOrder); err != nil {
		return err
	}
	if o.Conditions != nil {
		_ptr_Conditions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Conditions != nil {
				if err := o.Conditions.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PolicyConditionArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Conditions, _ptr_Conditions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Expressions != nil {
		_ptr_Expressions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Expressions != nil {
				if err := o.Expressions.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PolicyExprArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Expressions, _ptr_Expressions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Ranges != nil {
		_ptr_Ranges := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Ranges != nil {
				if err := o.Ranges.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPRangeArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Ranges, _ptr_Ranges); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_Description := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_Description); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.Enabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *Policy) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_PolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PolicyName); err != nil {
			return err
		}
		return nil
	})
	_s_PolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PolicyName, _s_PolicyName, _ptr_PolicyName); err != nil {
		return err
	}
	var _bIsGlobalPolicy int32
	if err := w.ReadData(&_bIsGlobalPolicy); err != nil {
		return err
	}
	o.IsGlobalPolicy = _bIsGlobalPolicy != 0
	if err := w.ReadData(&o.Subnet); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProcessingOrder); err != nil {
		return err
	}
	_ptr_Conditions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Conditions == nil {
			o.Conditions = &PolicyConditionArray{}
		}
		if err := o.Conditions.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Conditions := func(ptr interface{}) { o.Conditions = *ptr.(**PolicyConditionArray) }
	if err := w.ReadPointer(&o.Conditions, _s_Conditions, _ptr_Conditions); err != nil {
		return err
	}
	_ptr_Expressions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Expressions == nil {
			o.Expressions = &PolicyExprArray{}
		}
		if err := o.Expressions.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Expressions := func(ptr interface{}) { o.Expressions = *ptr.(**PolicyExprArray) }
	if err := w.ReadPointer(&o.Expressions, _s_Expressions, _ptr_Expressions); err != nil {
		return err
	}
	_ptr_Ranges := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Ranges == nil {
			o.Ranges = &IPRangeArray{}
		}
		if err := o.Ranges.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Ranges := func(ptr interface{}) { o.Ranges = *ptr.(**IPRangeArray) }
	if err := w.ReadPointer(&o.Ranges, _s_Ranges, _ptr_Ranges); err != nil {
		return err
	}
	_ptr_Description := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_Description := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_Description, _ptr_Description); err != nil {
		return err
	}
	var _bEnabled int32
	if err := w.ReadData(&_bEnabled); err != nil {
		return err
	}
	o.Enabled = _bEnabled != 0
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// PolicyArray structure represents DHCP_POLICY_ARRAY RPC structure.
//
// The DHCP_POLICY_ARRAY structure contains a list of policy elements.
type PolicyArray struct {
	// NumElements:  This member contains the number of policies in the array.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  This member is a pointer of type DHCP_POLICY (section 2.2.1.2.110) that
	// points to an array of length NumElements.
	Elements []*Policy `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *PolicyArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Policy{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Policy{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*Policy, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &Policy{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*Policy) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// StatelessParamsV6 structure represents DHCPV6_STATELESS_PARAMS RPC structure.
//
// The DHCPV6_STATELESS_PARAMS structure contains stateless settings for a DHCPv6 server.
// This structure is used with the R_DhcpV6SetStatelessStoreParams (section 3.2.4.117)
// and R_DhcpV6GetStatelessStoreParams (section 3.2.4.118) methods.
type StatelessParamsV6 struct {
	// Status:  This member indicates whether a stateless client inventory needs to be maintained
	// by the DHCPv6 server. The value of this member defaults to FALSE, indicating that
	// the server does not need to maintain a stateless client inventory.
	Status bool `idl:"name:Status" json:"status"`
	// PurgeInterval:  This member specifies the maximum time interval, in hours, that stateless
	// IPv6 DHCP client lease records will persist before being deleted from the DHCP server
	// database.
	PurgeInterval uint32 `idl:"name:PurgeInterval" json:"purge_interval"`
}

func (o *StatelessParamsV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StatelessParamsV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if !o.Status {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PurgeInterval); err != nil {
		return err
	}
	return nil
}
func (o *StatelessParamsV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	var _bStatus int32
	if err := w.ReadData(&_bStatus); err != nil {
		return err
	}
	o.Status = _bStatus != 0
	if err := w.ReadData(&o.PurgeInterval); err != nil {
		return err
	}
	return nil
}

// StatelessScopeStatsV6 structure represents DHCPV6_STATELESS_SCOPE_STATS RPC structure.
//
// The DHCPV6_STATELESS_SCOPE_STATS structure contains the address counters for a specific
// IPv6 stateless subnet. The number of stateless IPv6 clients added and removed from
// the stateless client inventory is stored in this structure. This structure is used
// in the DHCPV6_STATELESS_STATS (section 2.2.1.2.114) structure.
type StatelessScopeStatsV6 struct {
	// SubnetAddress:  The IPv6 prefix for the scope.
	SubnetAddress *IPv6Address `idl:"name:SubnetAddress" json:"subnet_address"`
	// NumStatelessClientsAdded:  The number of IPv6 stateless clients that have been added
	// to the DHCPv6 stateless client inventory for the IPv6 prefix stored in SubnetAddress.
	StatelessClientsAddedLength uint64 `idl:"name:NumStatelessClientsAdded" json:"stateless_clients_added_length"`
	// NumStatelessClientsRemoved:  The number of IPv6 stateless clients that have been
	// removed from the DHCPv6 stateless client inventory for the IPv6 prefix stored in
	// SubnetAddress.
	StatelessClientsRemovedLength uint64 `idl:"name:NumStatelessClientsRemoved" json:"stateless_clients_removed_length"`
}

func (o *StatelessScopeStatsV6) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StatelessScopeStatsV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.SubnetAddress != nil {
		if err := o.SubnetAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPv6Address{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.StatelessClientsAddedLength); err != nil {
		return err
	}
	if err := w.WriteData(o.StatelessClientsRemovedLength); err != nil {
		return err
	}
	return nil
}
func (o *StatelessScopeStatsV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.SubnetAddress == nil {
		o.SubnetAddress = &IPv6Address{}
	}
	if err := o.SubnetAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.StatelessClientsAddedLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.StatelessClientsRemovedLength); err != nil {
		return err
	}
	return nil
}

// StatelessStatsV6 structure represents DHCPV6_STATELESS_STATS RPC structure.
//
// The DHCPV6_STATELESS_STATS structure represents an array of DHCPV6_STATELESS_SCOPE_STATS
// (section 2.2.1.2.113) structures. This structure is used with the R_DhcpV6GetStatelessStatistics
// (section 3.2.4.119) method. The server uses this array for statistical analysis.
type StatelessStatsV6 struct {
	// NumScopes:  The number of elements in the ScopeStats member.
	ScopesLength uint32 `idl:"name:NumScopes" json:"scopes_length"`
	// ScopeStats:  A pointer to an array of DHCPV6_STATELESS_SCOPE_STATS (section 2.2.1.2.113)
	// structures, each one representing an IPv6 stateless prefix serviced by the current
	// DHCPv6 server.
	ScopeStats []*StatelessScopeStatsV6 `idl:"name:ScopeStats;size_is:(NumScopes)" json:"scope_stats"`
}

func (o *StatelessStatsV6) xxx_PreparePayload(ctx context.Context) error {
	if o.ScopeStats != nil && o.ScopesLength == 0 {
		o.ScopesLength = uint32(len(o.ScopeStats))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StatelessStatsV6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ScopesLength); err != nil {
		return err
	}
	if o.ScopeStats != nil || o.ScopesLength > 0 {
		_ptr_ScopeStats := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ScopesLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ScopeStats {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ScopeStats[i1] != nil {
					if err := o.ScopeStats[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&StatelessScopeStatsV6{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ScopeStats); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&StatelessScopeStatsV6{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ScopeStats, _ptr_ScopeStats); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *StatelessStatsV6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ScopesLength); err != nil {
		return err
	}
	_ptr_ScopeStats := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ScopesLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ScopesLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ScopeStats", sizeInfo[0])
		}
		o.ScopeStats = make([]*StatelessScopeStatsV6, sizeInfo[0])
		for i1 := range o.ScopeStats {
			i1 := i1
			if o.ScopeStats[i1] == nil {
				o.ScopeStats[i1] = &StatelessScopeStatsV6{}
			}
			if err := o.ScopeStats[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ScopeStats := func(ptr interface{}) { o.ScopeStats = *ptr.(*[]*StatelessScopeStatsV6) }
	if err := w.ReadPointer(&o.ScopeStats, _s_ScopeStats, _ptr_ScopeStats); err != nil {
		return err
	}
	return nil
}

// ClientInfoPB structure represents DHCP_CLIENT_INFO_PB RPC structure.
//
// The DHCP_CLIENT_INFO_PB structure encapsulates information about a DHCPv4 client,
// including filter status information and the policy, if any, that resulted in the
// client's specific IPv4 address assignment. This structure augments the DHCP_CLIENT_FILTER_STATUS_INFO
// (section 2.2.1.2.96) structure by including the PolicyName member.
type ClientInfoPB struct {
	// ClientIpAddress:  The client's IPv4 address.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// SubnetMask:  The client's IPv4 subnet mask.
	SubnetMask uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
	// ClientHardwareAddress:  A DHCP_CLIENT_UID structure containing the client's DHCPv4
	// client-identifier.
	ClientHardwareAddress *ClientUID `idl:"name:ClientHardwareAddress" json:"client_hardware_address"`
	// ClientName:  A pointer to a null-terminated Unicode string representing the client's
	// DHCPv4 internet host name. There is no restriction on the length of this string.
	ClientName string `idl:"name:ClientName" json:"client_name"`
	// ClientComment:  A pointer to a null-terminated Unicode string containing a description
	// given to the DHCPv4 client. This string is limited to 128 characters, including the
	// terminating null character.
	ClientComment string `idl:"name:ClientComment" json:"client_comment"`
	// ClientLeaseExpires:  The time at which the client's lease on its assigned IPv4 address
	// expires.
	ClientLeaseExpires *DateTime `idl:"name:ClientLeaseExpires" json:"client_lease_expires"`
	// OwnerHost:  A DHCP_HOST_INFO structure that contains information about the DHCPv4
	// server that assigned the client's IPv4 address.
	OwnerHost *HostInfo `idl:"name:OwnerHost" json:"owner_host"`
	// bClientType:  The type of the DHCPv4 client. This member MUST have one of the following
	// values.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_UNSPECIFIED 0x00      | Used for DHCPv4 clients that are not any of the other types in this table.       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_DHCP 0x01             | The DHCPv4 client supports the DHCP protocol.                                    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOOTP 0x02            | The DHCPv4 client supports the BOOTP protocol ([RFC2132]).                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOTH 0x03             | The DHCPv4 client supports the DHCP and BOOTP protocols.                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_RESERVATION_FLAG 0x04 | An IPv4 registration has been created for the DHCPv4 client.                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_NONE 0x05             | The client uses manual addressing. This value supports backward compatibility    |
	//	|                                   | with clients that do not use dynamic IP address assignment.                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	ClientType uint8 `idl:"name:bClientType" json:"client_type"`
	// AddressState:  The state of the IPv4 address given to the DHCPv4 client. This member
	// MUST be set to one of the following values.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                   DESCRIPTION                                    |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_OFFERED 0x00  | The address has been offered to the client.                                      |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_ACTIVE 0x01   | The address is active and has an active DHCPv4 client lease record.              |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DECLINED 0x02 | The IPv4 address request was declined by the DHCPv4 client. The server will not  |
	//	|                             | issue this IPv4 address to other clients for a period of time.                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DOOM 0x03     | The IPv4 address is in the DOOMED state prior to being deleted.                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	AddressState uint8 `idl:"name:AddressState" json:"address_state"`
	// Status:  A QuarantineStatus (section 2.2.1.1.11) value representing the DHCPv4 client's
	// health status as validated by the NAP server.
	Status QuarantineStatus `idl:"name:Status" json:"status"`
	// ProbationEnds:  If the client is on probation, this member contains the time at which
	// the probation ends. Up to that time, the client has full access to the network.
	ProbationEnds *DateTime `idl:"name:ProbationEnds" json:"probation_ends"`
	// QuarantineCapable:  Whether the client machine is quarantine-enabled. A value of
	// TRUE indicates that the client machine is quarantine-enabled, whereas FALSE indicates
	// that it is not.
	QuarantineCapable bool `idl:"name:QuarantineCapable" json:"quarantine_capable"`
	// FilterStatus:  The status of the link-layer filter. This member MUST be set to one
	// of the following values.
	//
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	|                                                       |                                                                               |
	//	|                         VALUE                         |                                  DESCRIPTION                                  |
	//	|                                                       |                                                                               |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	| FILTER_STATUS_NONE 0x00000001                         | The DHCPv4 client's MAC address does not match any filter.                    |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	| FILTER_STATUS_FULL_MATCH_IN_ALLOW_LIST 0x00000002     | The DHCPv4 client's MAC address fully matches an allow-list filter.           |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	| FILTER_STATUS_FULL_MATCH_IN_DENY_LIST 0x00000004      | The DHCPv4 client's MAC address fully matches a deny-list filter.             |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	| FILTER_STATUS_WILDCARD_MATCH_IN_ALLOW_LIST 0x00000008 | The DHCPv4 client's MAC address has a wildcard match to an allow-list filter. |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	| FILTER_STATUS_WILDCARD_MATCH_IN_DENY_LIST 0x00000010  | The DHCPv4 client's MAC address has a wildcard match to a deny-list filter.   |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	FilterStatus uint32 `idl:"name:FilterStatus" json:"filter_status"`
	// PolicyName:  A pointer to a null-terminated Unicode string containing the name of
	// the scope-level policy, if any, that resulted in the current IPv4 address being assigned
	// to the client. This string is limited to 64 characters, including the terminating
	// null character.
	PolicyName string `idl:"name:PolicyName" json:"policy_name"`
}

func (o *ClientInfoPB) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoPB) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress != nil {
		if err := o.ClientHardwareAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientComment != "" {
		_ptr_ClientComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientComment, _ptr_ClientComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientLeaseExpires != nil {
		if err := o.ClientLeaseExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OwnerHost != nil {
		if err := o.OwnerHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ClientType); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressState); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if o.ProbationEnds != nil {
		if err := o.ProbationEnds.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.QuarantineCapable {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.FilterStatus); err != nil {
		return err
	}
	if o.PolicyName != "" {
		_ptr_PolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PolicyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyName, _ptr_PolicyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoPB) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress == nil {
		o.ClientHardwareAddress = &ClientUID{}
	}
	if err := o.ClientHardwareAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	_ptr_ClientComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClientComment := func(ptr interface{}) { o.ClientComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientComment, _s_ClientComment, _ptr_ClientComment); err != nil {
		return err
	}
	if o.ClientLeaseExpires == nil {
		o.ClientLeaseExpires = &DateTime{}
	}
	if err := o.ClientLeaseExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OwnerHost == nil {
		o.OwnerHost = &HostInfo{}
	}
	if err := o.OwnerHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressState); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	if o.ProbationEnds == nil {
		o.ProbationEnds = &DateTime{}
	}
	if err := o.ProbationEnds.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bQuarantineCapable int32
	if err := w.ReadData(&_bQuarantineCapable); err != nil {
		return err
	}
	o.QuarantineCapable = _bQuarantineCapable != 0
	if err := w.ReadData(&o.FilterStatus); err != nil {
		return err
	}
	_ptr_PolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PolicyName); err != nil {
			return err
		}
		return nil
	})
	_s_PolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PolicyName, _s_PolicyName, _ptr_PolicyName); err != nil {
		return err
	}
	return nil
}

// ClientInfoPBArray structure represents DHCP_CLIENT_INFO_PB_ARRAY RPC structure.
//
// The DHCP_CLIENT_INFO_PB_ARRAY structure encapsulates an array of DHCP_CLIENT_INFO_PB
// (section 2.2.1.2.115) structures.
type ClientInfoPBArray struct {
	// NumElements:  The number of elements in the Clients member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Clients:  An array of pointers to DHCP_CLIENT_INFO_PB structures.
	Clients []*ClientInfoPB `idl:"name:Clients;size_is:(NumElements)" json:"clients"`
}

func (o *ClientInfoPBArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Clients != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Clients))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoPBArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Clients != nil || o.ElementsLength > 0 {
		_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Clients {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Clients[i1] != nil {
					_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Clients[i1] != nil {
							if err := o.Clients[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ClientInfoPB{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Clients[i1], _ptr_Clients); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Clients); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Clients, _ptr_Clients); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoPBArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Clients", sizeInfo[0])
		}
		o.Clients = make([]*ClientInfoPB, sizeInfo[0])
		for i1 := range o.Clients {
			i1 := i1
			_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Clients[i1] == nil {
					o.Clients[i1] = &ClientInfoPB{}
				}
				if err := o.Clients[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Clients := func(ptr interface{}) { o.Clients[i1] = *ptr.(**ClientInfoPB) }
			if err := w.ReadPointer(&o.Clients[i1], _s_Clients, _ptr_Clients); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Clients := func(ptr interface{}) { o.Clients = *ptr.(*[]*ClientInfoPB) }
	if err := w.ReadPointer(&o.Clients, _s_Clients, _ptr_Clients); err != nil {
		return err
	}
	return nil
}

// PropertyType type represents DHCP_PROPERTY_TYPE RPC enumeration.
//
// The DHCP_PROPERTY_TYPE enumeration defines the data types for DHCP property values
// and is used in the DHCP_PROPERTY (section 2.2.1.2.117) structure.
type PropertyType uint16

var (
	// DhcpPropTypeByte: The property value is of type BYTE.
	PropertyTypeByte PropertyType = 0
	// DhcpPropTypeWord: The property value is of type WORD.
	PropertyTypeWord PropertyType = 1
	// DhcpPropTypeDword: The property value is of type DWORD.
	PropertyTypeDword PropertyType = 2
	// DhcpPropTypeString: The property value is a pointer of type LPWSTR that points to
	// a Unicode string that includes the terminating null character.
	PropertyTypeString PropertyType = 3
	// DhcpPropTypeBinary: The property value is of type DHCP_BINARY_DATA (section 2.2.1.2.9).
	PropertyTypeBinary PropertyType = 4
)

func (o PropertyType) String() string {
	switch o {
	case PropertyTypeByte:
		return "PropertyTypeByte"
	case PropertyTypeWord:
		return "PropertyTypeWord"
	case PropertyTypeDword:
		return "PropertyTypeDword"
	case PropertyTypeString:
		return "PropertyTypeString"
	case PropertyTypeBinary:
		return "PropertyTypeBinary"
	}
	return "Invalid"
}

// PropertyID type represents DHCP_PROPERTY_ID RPC enumeration.
//
// The DHCP_PROPERTY_ID enumeration defines the property identifier for a DHCP_PROPERTY
// (section 2.2.1.2.117) structure. It is used to uniquely identify a specified DHCP
// property.
type PropertyID uint16

var (
	// DhcpPropIdPolicyDnsSuffix: Identifies the DNS suffix of a policy. It is of property
	// type DhcpPropTypeString, as specified in DHCP_PROPERTY_TYPE (section 2.2.1.1.26).
	PropertyIDPolicyDNSSuffix PropertyID = 0
	// DhcpPropIdClientAddressStateEx: Identifies the extended address state flags of a
	// lease table entry. It is of property type DhcpPropTypeDword, as specified in DHCP_PROPERTY_TYPE
	// (section 2.2.1.1.26).
	PropertyIDClientAddressStateEx PropertyID = 1
)

func (o PropertyID) String() string {
	switch o {
	case PropertyIDPolicyDNSSuffix:
		return "PropertyIDPolicyDNSSuffix"
	case PropertyIDClientAddressStateEx:
		return "PropertyIDClientAddressStateEx"
	}
	return "Invalid"
}

// Property structure represents DHCP_PROPERTY RPC structure.
//
// The DHCP_PROPERTY structure contains the type of the property, the property identifier,
// and the property data value. The DHCP_PROPERTY identifies a DHCP property and is
// used by the DHCP_CLIENT_INFO_EX (section 2.2.1.2.119) and DHCP_POLICY_EX (section
// 2.2.1.2.121) structures, which allow a list of properties to be associated with them.
type Property struct {
	// ID:  An enumeration of type DHCP_PROPERTY_ID (section 2.2.1.1.27) that indicates
	// the property identifier for the data value contained in the Value field.
	ID PropertyID `idl:"name:ID" json:"id"`
	// Type:  An enumeration of type DHCP_PROPERTY_TYPE (section 2.2.1.1.26) that indicates
	// the property type for the data value contained in the Value field.
	Type PropertyType `idl:"name:Type" json:"type"`
	// Value:  Specifies the property data using one of the following values based on the
	// value of the Type field.
	Value *Property_Value `idl:"name:Value;switch_is:Type" json:"value"`
}

func (o *Property) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Property) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ID)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
		return err
	}
	_swValue := uint16(o.Type)
	if o.Value != nil {
		if err := o.Value.MarshalUnionNDR(ctx, w, _swValue); err != nil {
			return err
		}
	} else {
		if err := (&Property_Value{}).MarshalUnionNDR(ctx, w, _swValue); err != nil {
			return err
		}
	}
	return nil
}
func (o *Property) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ID)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
		return err
	}
	if o.Value == nil {
		o.Value = &Property_Value{}
	}
	_swValue := uint16(o.Type)
	if err := o.Value.UnmarshalUnionNDR(ctx, w, _swValue); err != nil {
		return err
	}
	return nil
}

// Property_Value structure represents DHCP_PROPERTY union anonymous member.
//
// The DHCP_PROPERTY structure contains the type of the property, the property identifier,
// and the property data value. The DHCP_PROPERTY identifies a DHCP property and is
// used by the DHCP_CLIENT_INFO_EX (section 2.2.1.2.119) and DHCP_POLICY_EX (section
// 2.2.1.2.121) structures, which allow a list of properties to be associated with them.
type Property_Value struct {
	// Types that are assignable to Value
	//
	// *Property_ByteValue
	// *Property_WordValue
	// *Property_DwordValue
	// *Property_StringValue
	// *Property_BinaryValue
	Value is_Property_Value `json:"value"`
}

func (o *Property_Value) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Property_ByteValue:
		if value != nil {
			return value.ByteValue
		}
	case *Property_WordValue:
		if value != nil {
			return value.WordValue
		}
	case *Property_DwordValue:
		if value != nil {
			return value.DwordValue
		}
	case *Property_StringValue:
		if value != nil {
			return value.StringValue
		}
	case *Property_BinaryValue:
		if value != nil {
			return value.BinaryValue
		}
	}
	return nil
}

type is_Property_Value interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Property_Value()
}

func (o *Property_Value) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Property_ByteValue:
		return uint16(0)
	case *Property_WordValue:
		return uint16(1)
	case *Property_DwordValue:
		return uint16(2)
	case *Property_StringValue:
		return uint16(3)
	case *Property_BinaryValue:
		return uint16(4)
	}
	return uint16(0)
}

func (o *Property_Value) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*Property_ByteValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Property_ByteValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Property_WordValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Property_WordValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*Property_DwordValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Property_DwordValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*Property_StringValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Property_StringValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4):
		_o, _ := o.Value.(*Property_BinaryValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Property_BinaryValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *Property_Value) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &Property_ByteValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Property_WordValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &Property_DwordValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &Property_StringValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4):
		o.Value = &Property_BinaryValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// Property_ByteValue structure represents Property_Value RPC union arm.
//
// It has following labels: 0
type Property_ByteValue struct {
	// ByteValue:  Specifies the data as a BYTE value. This field is present if the Type
	// field is set to DhcpPropertyTypeByte.
	ByteValue uint8 `idl:"name:ByteValue" json:"byte_value"`
}

func (*Property_ByteValue) is_Property_Value() {}

func (o *Property_ByteValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.ByteValue); err != nil {
		return err
	}
	return nil
}
func (o *Property_ByteValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.ByteValue); err != nil {
		return err
	}
	return nil
}

// Property_WordValue structure represents Property_Value RPC union arm.
//
// It has following labels: 1
type Property_WordValue struct {
	// WordValue:  Specifies the data as a WORD value. This field is present if the Type
	// field is set to DhcpPropertyTypeWord.
	WordValue uint16 `idl:"name:WordValue" json:"word_value"`
}

func (*Property_WordValue) is_Property_Value() {}

func (o *Property_WordValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.WordValue); err != nil {
		return err
	}
	return nil
}
func (o *Property_WordValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.WordValue); err != nil {
		return err
	}
	return nil
}

// Property_DwordValue structure represents Property_Value RPC union arm.
//
// It has following labels: 2
type Property_DwordValue struct {
	// DWordValue:  Specifies the data as a DWORD value. This field is present if the Type
	// field is set to DhcpPropertyTypeDWord.
	DwordValue uint32 `idl:"name:DWordValue" json:"dword_value"`
}

func (*Property_DwordValue) is_Property_Value() {}

func (o *Property_DwordValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.DwordValue); err != nil {
		return err
	}
	return nil
}
func (o *Property_DwordValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.DwordValue); err != nil {
		return err
	}
	return nil
}

// Property_StringValue structure represents Property_Value RPC union arm.
//
// It has following labels: 3
type Property_StringValue struct {
	// StringValue:  Specifies the data as a LPWSTR pointer to a Unicode string value. This
	// field is present if the Type field is set to DhcpPropertyTypeString.
	StringValue string `idl:"name:StringValue" json:"string_value"`
}

func (*Property_StringValue) is_Property_Value() {}

func (o *Property_StringValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StringValue != "" {
		_ptr_StringValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.StringValue); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.StringValue, _ptr_StringValue); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Property_StringValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_StringValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.StringValue); err != nil {
			return err
		}
		return nil
	})
	_s_StringValue := func(ptr interface{}) { o.StringValue = *ptr.(*string) }
	if err := w.ReadPointer(&o.StringValue, _s_StringValue, _ptr_StringValue); err != nil {
		return err
	}
	return nil
}

// Property_BinaryValue structure represents Property_Value RPC union arm.
//
// It has following labels: 4
type Property_BinaryValue struct {
	// BinaryValue:  Specifies the data as a DHCP_BINARY_DATA (section 2.2.1.2.9) structure.
	// This field is present if the Type field is set to DhcpPropertyTypeBinary.
	BinaryValue *BinaryData `idl:"name:BinaryValue" json:"binary_value"`
}

func (*Property_BinaryValue) is_Property_Value() {}

func (o *Property_BinaryValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.BinaryValue != nil {
		if err := o.BinaryValue.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&BinaryData{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Property_BinaryValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.BinaryValue == nil {
		o.BinaryValue = &BinaryData{}
	}
	if err := o.BinaryValue.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyArray structure represents DHCP_PROPERTY_ARRAY RPC structure.
//
// The DHCP_PROPERTY_ARRAY structure defines an array of DHCP_PROPERTY (section 2.2.1.2.117)
// structures. This structure is a data container for one or more data elements associated
// with a DHCP property array.
type PropertyArray struct {
	// NumElements:  Specifies the number of DHCP Property elements.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  Contains the DHCP Property elements.
	Elements []*Property `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *PropertyArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Property{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Property{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*Property, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &Property{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*Property) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}

// ClientInfoEx structure represents DHCP_CLIENT_INFO_EX RPC structure.
//
// The DHCP_CLIENT_INFO_EX structure encapsulates information about a DHCPv4 client,
// including filter status information and the policy, if any, that resulted in the
// client's specific IPv4 address assignment. This structure augments the DHCP_CLIENT_INFO
// (section 2.2.1.2.12) structure by including a list of DHCP properties represented
// by the field Properties.
type ClientInfoEx struct {
	// ClientIpAddress:  The client's IPv4 address.
	ClientIPAddress uint32 `idl:"name:ClientIpAddress" json:"client_ip_address"`
	// SubnetMask:  The client's IPv4 subnet mask.
	SubnetMask uint32 `idl:"name:SubnetMask" json:"subnet_mask"`
	// ClientHardwareAddress:  A DHCP_CLIENT_UID (section 2.2.1.2.5) structure containing
	// the client'sDHCPv4 client-identifier.
	ClientHardwareAddress *ClientUID `idl:"name:ClientHardwareAddress" json:"client_hardware_address"`
	// ClientName:  A pointer to a null-terminated Unicode string representing the client's
	// DHCPv4 internet host name. There is no restriction on the string length.
	ClientName string `idl:"name:ClientName" json:"client_name"`
	// ClientComment:  A pointer to a null-terminated Unicode string representing the client's
	// DHCPv4 internet host name. There is no restriction on the string length.
	ClientComment string `idl:"name:ClientComment" json:"client_comment"`
	// ClientLeaseExpires:  The time at which the client's lease on its assigned IPv4 address
	// expires.
	ClientLeaseExpires *DateTime `idl:"name:ClientLeaseExpires" json:"client_lease_expires"`
	// OwnerHost:  A DHCP_HOST_INFO (section 2.2.1.2.7) structure that contains information
	// about the DHCPv4 server that assigned the client's IPv4 address.
	OwnerHost *HostInfo `idl:"name:OwnerHost" json:"owner_host"`
	// bClientType:  The type of the DHCPv4 client. This member MUST have one of the following
	// values.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_UNSPECIFIED 0x00      | Used for DHCPv4 clients that are not any other type as defined in this table.    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_DHCP 0x01             | The DHCPv4 client supports the DHCP protocol.                                    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOOTP 0x02            | The DHCPv4 client supports the BOOTP protocol ([RFC2132]).                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_BOTH 0x03             | The DHCPv4 client supports the DHCP and BOOTP protocols.                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_RESERVATION_FLAG 0x04 | An IPv4 registration has been created for the DHCPv4 client.                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| CLIENT_TYPE_NONE 0x05             | The client uses manual addressing. This value supports backward compatibility    |
	//	|                                   | with clients that do not use dynamic IP address assignment.                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	ClientType uint8 `idl:"name:bClientType" json:"client_type"`
	// AddressState:  The state of the IPv4 address for the DHCPv4 client. This member MUST
	// be set to one of the following values.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                   DESCRIPTION                                    |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_OFFERED 0x00  | The address has been offered to the client.                                      |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_ACTIVE 0x01   | The address is active and has an active DHCPv4 client lease record.              |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DECLINED 0x02 | The IPv4 address request was declined by the DHCPv4 client. The server will not  |
	//	|                             | issue this IPv4 address to other clients for a period of time.                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| ADDRESS_STATE_DOOM 0x03     | The IPv4 address is in the DOOMED state prior to being deleted.                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	AddressState uint8 `idl:"name:AddressState" json:"address_state"`
	// Status:  A QuarantineStatus (section 2.2.1.1.11) value representing the DHCPv4 client's
	// health status as validated by the NAP server.
	Status QuarantineStatus `idl:"name:Status" json:"status"`
	// ProbationEnds:  If the client is on probation, this member contains the time at which
	// the probation ends. Up to the time at which probation ends, the client has full access
	// to the network.
	ProbationEnds *DateTime `idl:"name:ProbationEnds" json:"probation_ends"`
	// QuarantineCapable:  Indicates whether the client machine is quarantine-enabled. A
	// value of TRUE indicates that the client machine is quarantine-enabled, whereas FALSE
	// indicates that it is not.
	QuarantineCapable bool `idl:"name:QuarantineCapable" json:"quarantine_capable"`
	// FilterStatus:  The status of the link-layer filter. This member MUST be set to one
	// of the following values.
	//
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	|                                                       |                                                                               |
	//	|                         VALUE                         |                                  DESCRIPTION                                  |
	//	|                                                       |                                                                               |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	| FILTER_STATUS_NONE 0x00000001                         | The DHCPv4 client's MAC address does not match any filter.                    |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	| FILTER_STATUS_FULL_MATCH_IN_ALLOW_LIST 0x00000002     | The DHCPv4 client's MAC address fully matches an allow-list filter.           |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	| FILTER_STATUS_FULL_MATCH_IN_DENY_LIST 0x00000004      | The DHCPv4 client's MAC address fully matches a deny-list filter.             |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	| FILTER_STATUS_WILDCARD_MATCH_IN_ALLOW_LIST 0x00000008 | The DHCPv4 client's MAC address has a wildcard match to an allow-list filter. |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	//	| FILTER_STATUS_WILDCARD_MATCH_IN_DENY_LIST 0x00000010  | The DHCPv4 client's MAC address has a wildcard match to a deny-list filter.   |
	//	+-------------------------------------------------------+-------------------------------------------------------------------------------+
	FilterStatus uint32 `idl:"name:FilterStatus" json:"filter_status"`
	// PolicyName:  A pointer to a null-terminated Unicode string containing the name of
	// the scope-level policy, if any, that resulted in the current IPv4 address being assigned
	// to the client. The string is limited to 64 characters, including the terminating
	// null character.
	PolicyName string `idl:"name:PolicyName" json:"policy_name"`
	// Properties:  A list of properties that is associated with the given client. See the
	// following list for properties allowed for the DHCPv4 client. Properties not identified
	// are ignored.
	Properties *PropertyArray `idl:"name:Properties" json:"properties"`
}

func (o *ClientInfoEx) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.WriteData(o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress != nil {
		if err := o.ClientHardwareAddress.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClientUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClientName != "" {
		_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientComment != "" {
		_ptr_ClientComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ClientComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ClientComment, _ptr_ClientComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ClientLeaseExpires != nil {
		if err := o.ClientLeaseExpires.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.OwnerHost != nil {
		if err := o.OwnerHost.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HostInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ClientType); err != nil {
		return err
	}
	if err := w.WriteData(o.AddressState); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Status)); err != nil {
		return err
	}
	if o.ProbationEnds != nil {
		if err := o.ProbationEnds.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DateTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.QuarantineCapable {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.FilterStatus); err != nil {
		return err
	}
	if o.PolicyName != "" {
		_ptr_PolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PolicyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyName, _ptr_PolicyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Properties != nil {
		_ptr_Properties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Properties != nil {
				if err := o.Properties.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PropertyArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Properties, _ptr_Properties); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientIPAddress); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubnetMask); err != nil {
		return err
	}
	if o.ClientHardwareAddress == nil {
		o.ClientHardwareAddress = &ClientUID{}
	}
	if err := o.ClientHardwareAddress.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
			return err
		}
		return nil
	})
	_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
		return err
	}
	_ptr_ClientComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComment); err != nil {
			return err
		}
		return nil
	})
	_s_ClientComment := func(ptr interface{}) { o.ClientComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ClientComment, _s_ClientComment, _ptr_ClientComment); err != nil {
		return err
	}
	if o.ClientLeaseExpires == nil {
		o.ClientLeaseExpires = &DateTime{}
	}
	if err := o.ClientLeaseExpires.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.OwnerHost == nil {
		o.OwnerHost = &HostInfo{}
	}
	if err := o.OwnerHost.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientType); err != nil {
		return err
	}
	if err := w.ReadData(&o.AddressState); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Status)); err != nil {
		return err
	}
	if o.ProbationEnds == nil {
		o.ProbationEnds = &DateTime{}
	}
	if err := o.ProbationEnds.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bQuarantineCapable int32
	if err := w.ReadData(&_bQuarantineCapable); err != nil {
		return err
	}
	o.QuarantineCapable = _bQuarantineCapable != 0
	if err := w.ReadData(&o.FilterStatus); err != nil {
		return err
	}
	_ptr_PolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PolicyName); err != nil {
			return err
		}
		return nil
	})
	_s_PolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PolicyName, _s_PolicyName, _ptr_PolicyName); err != nil {
		return err
	}
	_ptr_Properties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Properties == nil {
			o.Properties = &PropertyArray{}
		}
		if err := o.Properties.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Properties := func(ptr interface{}) { o.Properties = *ptr.(**PropertyArray) }
	if err := w.ReadPointer(&o.Properties, _s_Properties, _ptr_Properties); err != nil {
		return err
	}
	return nil
}

// ClientInfoExArray structure represents DHCP_CLIENT_INFO_EX_ARRAY RPC structure.
//
// The DHCP_CLIENT_INFO_EX_ARRAY structure encapsulates an array of DHCP_CLIENT_INFO_EX
// (section 2.2.1.2.119) structures.
type ClientInfoExArray struct {
	// NumElements:  The number of elements in the Clients member.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Clients:  An array of pointers to DHCP_CLIENT_INFO_EX (section 2.2.1.2.119) structures.
	Clients []*ClientInfoEx `idl:"name:Clients;size_is:(NumElements)" json:"clients"`
}

func (o *ClientInfoExArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Clients != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Clients))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoExArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Clients != nil || o.ElementsLength > 0 {
		_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Clients {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Clients[i1] != nil {
					_ptr_Clients := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Clients[i1] != nil {
							if err := o.Clients[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ClientInfoEx{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Clients[i1], _ptr_Clients); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Clients); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Clients, _ptr_Clients); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClientInfoExArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Clients", sizeInfo[0])
		}
		o.Clients = make([]*ClientInfoEx, sizeInfo[0])
		for i1 := range o.Clients {
			i1 := i1
			_ptr_Clients := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Clients[i1] == nil {
					o.Clients[i1] = &ClientInfoEx{}
				}
				if err := o.Clients[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_Clients := func(ptr interface{}) { o.Clients[i1] = *ptr.(**ClientInfoEx) }
			if err := w.ReadPointer(&o.Clients[i1], _s_Clients, _ptr_Clients); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Clients := func(ptr interface{}) { o.Clients = *ptr.(*[]*ClientInfoEx) }
	if err := w.ReadPointer(&o.Clients, _s_Clients, _ptr_Clients); err != nil {
		return err
	}
	return nil
}

// PolicyEx structure represents DHCP_POLICY_EX RPC structure.
//
// The DHCP_POLICY_EX structure contains information for a policy that is used to filter
// client requests. This structure augments the DHCP_POLICY (section 2.2.1.2.110) structure
// by including a list of properties represented by the field Properties.
type PolicyEx struct {
	// PolicyName:  A pointer of type LPWSTR that points to a null-terminated Unicode string
	// identifying the name of the policy. The name of the policy is restricted to 64 characters.
	PolicyName string `idl:"name:PolicyName" json:"policy_name"`
	// IsGlobalPolicy:  Indicates whether this is a server-level policy.
	IsGlobalPolicy bool `idl:"name:IsGlobalPolicy" json:"is_global_policy"`
	// Subnet:  This member is of type DHCP_IP_ADDRESS (section 2.2.1.2.1) structure and
	// identifies the IPv4 subnet to which the policy belongs, if this is a scope-level
	// policy. The value of this member will be 0 for a server-level policy.
	Subnet uint32 `idl:"name:Subnet" json:"subnet"`
	// ProcessingOrder:  Identifies the relative order in which the DHCPv4 server will process
	// the policy.
	ProcessingOrder uint32 `idl:"name:ProcessingOrder" json:"processing_order"`
	// Conditions:  A pointer of type LPDHCP_POL_COND_ARRAY (section 2.2.1.2.106) that contains
	// the array of conditions for the policy.
	Conditions *PolicyConditionArray `idl:"name:Conditions" json:"conditions"`
	// Expressions:  A pointer of type LPDHCP_POL_EXPR_ARRAY (section 2.2.1.2.108) that
	// contains the array of expressions for the policy.
	Expressions *PolicyExprArray `idl:"name:Expressions" json:"expressions"`
	// Ranges:  A pointer of type LPDHCP_IP_RANGE_ARRAY (section 2.2.1.2.104) that points
	// to an array of DHCP_IP_RANGE (section 2.2.1.2.31) structures that represent the policy
	// IP ranges.
	Ranges *IPRangeArray `idl:"name:Ranges" json:"ranges"`
	// Description:  A pointer of type LPWSTR that contains the null-terminated Unicode
	// string containing the description of the policy. The string is restricted to 255
	// characters.
	Description string `idl:"name:Description" json:"description"`
	// Enabled:  Indicates whether the policy is in the enabled (TRUE) or disabled (FALSE)
	// state.
	Enabled bool `idl:"name:Enabled" json:"enabled"`
	// Properties:  A list of properties that is associated with the given client. See the
	// following list for allowed properties. Properties not identified are ignored.
	Properties *PropertyArray `idl:"name:Properties" json:"properties"`
}

func (o *PolicyEx) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.PolicyName != "" {
		_ptr_PolicyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.PolicyName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyName, _ptr_PolicyName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.IsGlobalPolicy {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Subnet); err != nil {
		return err
	}
	if err := w.WriteData(o.ProcessingOrder); err != nil {
		return err
	}
	if o.Conditions != nil {
		_ptr_Conditions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Conditions != nil {
				if err := o.Conditions.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PolicyConditionArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Conditions, _ptr_Conditions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Expressions != nil {
		_ptr_Expressions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Expressions != nil {
				if err := o.Expressions.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PolicyExprArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Expressions, _ptr_Expressions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Ranges != nil {
		_ptr_Ranges := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Ranges != nil {
				if err := o.Ranges.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPRangeArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Ranges, _ptr_Ranges); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != "" {
		_ptr_Description := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Description); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_Description); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.Enabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.Properties != nil {
		_ptr_Properties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Properties != nil {
				if err := o.Properties.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&PropertyArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Properties, _ptr_Properties); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_PolicyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.PolicyName); err != nil {
			return err
		}
		return nil
	})
	_s_PolicyName := func(ptr interface{}) { o.PolicyName = *ptr.(*string) }
	if err := w.ReadPointer(&o.PolicyName, _s_PolicyName, _ptr_PolicyName); err != nil {
		return err
	}
	var _bIsGlobalPolicy int32
	if err := w.ReadData(&_bIsGlobalPolicy); err != nil {
		return err
	}
	o.IsGlobalPolicy = _bIsGlobalPolicy != 0
	if err := w.ReadData(&o.Subnet); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProcessingOrder); err != nil {
		return err
	}
	_ptr_Conditions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Conditions == nil {
			o.Conditions = &PolicyConditionArray{}
		}
		if err := o.Conditions.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Conditions := func(ptr interface{}) { o.Conditions = *ptr.(**PolicyConditionArray) }
	if err := w.ReadPointer(&o.Conditions, _s_Conditions, _ptr_Conditions); err != nil {
		return err
	}
	_ptr_Expressions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Expressions == nil {
			o.Expressions = &PolicyExprArray{}
		}
		if err := o.Expressions.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Expressions := func(ptr interface{}) { o.Expressions = *ptr.(**PolicyExprArray) }
	if err := w.ReadPointer(&o.Expressions, _s_Expressions, _ptr_Expressions); err != nil {
		return err
	}
	_ptr_Ranges := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Ranges == nil {
			o.Ranges = &IPRangeArray{}
		}
		if err := o.Ranges.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Ranges := func(ptr interface{}) { o.Ranges = *ptr.(**IPRangeArray) }
	if err := w.ReadPointer(&o.Ranges, _s_Ranges, _ptr_Ranges); err != nil {
		return err
	}
	_ptr_Description := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Description); err != nil {
			return err
		}
		return nil
	})
	_s_Description := func(ptr interface{}) { o.Description = *ptr.(*string) }
	if err := w.ReadPointer(&o.Description, _s_Description, _ptr_Description); err != nil {
		return err
	}
	var _bEnabled int32
	if err := w.ReadData(&_bEnabled); err != nil {
		return err
	}
	o.Enabled = _bEnabled != 0
	_ptr_Properties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Properties == nil {
			o.Properties = &PropertyArray{}
		}
		if err := o.Properties.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Properties := func(ptr interface{}) { o.Properties = *ptr.(**PropertyArray) }
	if err := w.ReadPointer(&o.Properties, _s_Properties, _ptr_Properties); err != nil {
		return err
	}
	return nil
}

// PolicyExArray structure represents DHCP_POLICY_EX_ARRAY RPC structure.
//
// The DHCP_POLICY_EX_ARRAY structure contains a list of policy elements.
type PolicyExArray struct {
	// NumElements:  Specifies the number of policies in the array.
	ElementsLength uint32 `idl:"name:NumElements" json:"elements_length"`
	// Elements:  A pointer of type DHCP_POLICY_EX (section 2.2.1.2.121) that points to
	// an array with length as specified in the NumElements member.
	Elements []*PolicyEx `idl:"name:Elements;size_is:(NumElements)" json:"elements"`
}

func (o *PolicyExArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elements != nil && o.ElementsLength == 0 {
		o.ElementsLength = uint32(len(o.Elements))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyExArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if o.Elements != nil || o.ElementsLength > 0 {
		_ptr_Elements := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElementsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elements {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elements[i1] != nil {
					if err := o.Elements[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PolicyEx{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elements); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&PolicyEx{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elements, _ptr_Elements); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PolicyExArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	_ptr_Elements := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElementsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElementsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elements", sizeInfo[0])
		}
		o.Elements = make([]*PolicyEx, sizeInfo[0])
		for i1 := range o.Elements {
			i1 := i1
			if o.Elements[i1] == nil {
				o.Elements[i1] = &PolicyEx{}
			}
			if err := o.Elements[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Elements := func(ptr interface{}) { o.Elements = *ptr.(*[]*PolicyEx) }
	if err := w.ReadPointer(&o.Elements, _s_Elements, _ptr_Elements); err != nil {
		return err
	}
	return nil
}
